package sema

import (
	"github.com/dapperlabs/flow-go/pkg/language/runtime/ast"
	"github.com/dapperlabs/flow-go/pkg/language/runtime/common"
	"github.com/dapperlabs/flow-go/pkg/language/runtime/errors"
	"github.com/dapperlabs/flow-go/pkg/language/runtime/sema/self_field_analyzer"
)

func (checker *Checker) VisitCompositeDeclaration(declaration *ast.CompositeDeclaration) ast.Repr {

	compositeType := checker.Elaboration.CompositeDeclarationTypes[declaration]

	// TODO: also check nested composite members

	// TODO: also check nested composite members' identifiers

	// TODO: also check nested composite fields' type annotations

	checker.checkMemberIdentifiers(
		declaration.Members.Fields,
		declaration.Members.Functions,
	)

	checker.checkInitializers(
		declaration.Members.Initializers,
		declaration.Members.Fields,
		compositeType,
		declaration.DeclarationKind(),
		declaration.Identifier.Identifier,
		compositeType.ConstructorParameterTypeAnnotations,
		initializerKindComposite,
	)

	checker.checkFieldsInitialized(declaration, compositeType)

	checker.checkCompositeFunctions(declaration.Members.Functions, compositeType)

	// check composite conforms to interfaces.
	// NOTE: perform after completing composite type (e.g. setting constructor parameter types)

	for i, interfaceType := range compositeType.Conformances {
		conformance := declaration.Conformances[i]

		checker.checkCompositeConformance(
			compositeType,
			interfaceType,
			declaration.Identifier.Pos,
			conformance.Identifier,
		)
	}

	// TODO: support non-structure composites, such as contracts and resources

	if declaration.CompositeKind != common.CompositeKindStructure {
		checker.report(
			&UnsupportedDeclarationError{
				DeclarationKind: declaration.DeclarationKind(),
				StartPos:        declaration.Identifier.StartPosition(),
				EndPos:          declaration.Identifier.EndPosition(),
			},
		)
	}

	// TODO: support nested declarations for contracts and contract interfaces

	// report error for first nested composite declaration, if any
	if len(declaration.Members.CompositeDeclarations) > 0 {
		firstNestedCompositeDeclaration := declaration.Members.CompositeDeclarations[0]

		checker.report(
			&UnsupportedDeclarationError{
				DeclarationKind: firstNestedCompositeDeclaration.DeclarationKind(),
				StartPos:        firstNestedCompositeDeclaration.Identifier.StartPosition(),
				EndPos:          firstNestedCompositeDeclaration.Identifier.EndPosition(),
			},
		)
	}

	return nil
}

func (checker *Checker) declareCompositeDeclaration(declaration *ast.CompositeDeclaration) {

	// NOTE: fields and functions might already refer to declaration itself.
	// insert a dummy type for now, so lookup succeeds during conversion,
	// then fix up the type reference

	compositeType := &CompositeType{}

	identifier := declaration.Identifier

	err := checker.typeActivations.Declare(identifier, compositeType)
	checker.report(err)
	checker.recordVariableDeclarationOccurrence(
		identifier.Identifier,
		&Variable{
			Kind:       declaration.DeclarationKind(),
			IsConstant: true,
			Type:       compositeType,
			Pos:        &identifier.Pos,
		},
	)

	conformances := checker.conformances(declaration)

	members, origins := checker.membersAndOrigins(
		declaration.Members.Fields,
		declaration.Members.Functions,
		true,
	)

	*compositeType = CompositeType{
		Kind:         declaration.CompositeKind,
		Identifier:   identifier.Identifier,
		Members:      members,
		Conformances: conformances,
	}

	checker.memberOrigins[compositeType] = origins

	// TODO: support multiple overloaded initializers

	var parameterTypeAnnotations []*TypeAnnotation
	initializerCount := len(declaration.Members.Initializers)
	if initializerCount > 0 {
		firstInitializer := declaration.Members.Initializers[0]
		parameterTypeAnnotations = checker.parameterTypeAnnotations(firstInitializer.Parameters)

		if initializerCount > 1 {
			secondInitializer := declaration.Members.Initializers[1]

			checker.report(
				&UnsupportedOverloadingError{
					DeclarationKind: common.DeclarationKindInitializer,
					StartPos:        secondInitializer.StartPosition(),
					EndPos:          secondInitializer.EndPosition(),
				},
			)
		}
	}

	compositeType.ConstructorParameterTypeAnnotations = parameterTypeAnnotations

	checker.Elaboration.CompositeDeclarationTypes[declaration] = compositeType

	// declare constructor

	checker.declareCompositeConstructor(declaration, compositeType, parameterTypeAnnotations)
}

func (checker *Checker) conformances(declaration *ast.CompositeDeclaration) []*InterfaceType {

	var interfaceTypes []*InterfaceType
	seenConformances := map[string]bool{}

	compositeIdentifier := declaration.Identifier.Identifier

	for _, conformance := range declaration.Conformances {
		convertedType := checker.ConvertType(conformance)

		if interfaceType, ok := convertedType.(*InterfaceType); ok {
			interfaceTypes = append(interfaceTypes, interfaceType)

		} else if !IsInvalidType(convertedType) {
			checker.report(
				&InvalidConformanceError{
					Type: convertedType,
					Pos:  conformance.Pos,
				},
			)
		}

		conformanceIdentifier := conformance.Identifier.Identifier

		if seenConformances[conformanceIdentifier] {
			checker.report(
				&DuplicateConformanceError{
					CompositeIdentifier: compositeIdentifier,
					Conformance:         conformance,
				},
			)

		}
		seenConformances[conformanceIdentifier] = true
	}
	return interfaceTypes
}

func (checker *Checker) checkCompositeConformance(
	compositeType *CompositeType,
	interfaceType *InterfaceType,
	compositeIdentifierPos ast.Position,
	interfaceIdentifier ast.Identifier,
) {
	var missingMembers []*Member
	var memberMismatches []MemberMismatch
	var initializerMismatch *InitializerMismatch

	// ensure the composite kinds match, e.g. a structure shouldn't be able
	// to conform to a resource interface

	if interfaceType.CompositeKind != compositeType.Kind {
		checker.report(
			&CompositeKindMismatchError{
				ExpectedKind: compositeType.Kind,
				ActualKind:   interfaceType.CompositeKind,
				StartPos:     interfaceIdentifier.StartPosition(),
				EndPos:       interfaceIdentifier.EndPosition(),
			},
		)
	}

	if interfaceType.InitializerParameterTypeAnnotations != nil {

		initializerType := &FunctionType{
			ParameterTypeAnnotations: compositeType.ConstructorParameterTypeAnnotations,
			ReturnTypeAnnotation:     NewTypeAnnotation(&VoidType{}),
		}
		interfaceInitializerType := &FunctionType{
			ParameterTypeAnnotations: interfaceType.InitializerParameterTypeAnnotations,
			ReturnTypeAnnotation:     NewTypeAnnotation(&VoidType{}),
		}

		// TODO: subtype?
		if !initializerType.Equal(interfaceInitializerType) {
			initializerMismatch = &InitializerMismatch{
				CompositeParameterTypes: compositeType.ConstructorParameterTypeAnnotations,
				InterfaceParameterTypes: interfaceType.InitializerParameterTypeAnnotations,
			}
		}
	}

	for name, interfaceMember := range interfaceType.Members {

		compositeMember, ok := compositeType.Members[name]
		if !ok {
			missingMembers = append(missingMembers, interfaceMember)
			continue
		}

		if !checker.memberSatisfied(compositeMember, interfaceMember) {
			memberMismatches = append(memberMismatches,
				MemberMismatch{
					CompositeMember: compositeMember,
					InterfaceMember: interfaceMember,
				},
			)
		}
	}

	if len(missingMembers) > 0 ||
		len(memberMismatches) > 0 ||
		initializerMismatch != nil {

		checker.report(
			&ConformanceError{
				CompositeType:       compositeType,
				InterfaceType:       interfaceType,
				Pos:                 compositeIdentifierPos,
				InitializerMismatch: initializerMismatch,
				MissingMembers:      missingMembers,
				MemberMismatches:    memberMismatches,
			},
		)
	}
}

func (checker *Checker) memberSatisfied(compositeMember, interfaceMember *Member) bool {
	// TODO: subtype?
	if !compositeMember.Type.Equal(interfaceMember.Type) {
		return false
	}

	if interfaceMember.VariableKind != ast.VariableKindNotSpecified &&
		compositeMember.VariableKind != interfaceMember.VariableKind {

		return false
	}

	return true
}

func (checker *Checker) checkFieldsInitialized(
	declaration *ast.CompositeDeclaration,
	compositeType *CompositeType,
) {
	for _, initializer := range declaration.Members.Initializers {
		unassigned, errs := self_field_analyzer.CheckSelfFieldInitializations(
			declaration.Members.Fields,
			initializer.FunctionBlock,
		)

		for _, field := range unassigned {
			checker.report(
				&FieldUninitializedError{
					Name:          field.Identifier.Identifier,
					Pos:           field.Identifier.Pos,
					CompositeType: compositeType,
					Initializer:   initializer,
				},
			)
		}

		checker.report(errs...)
	}
}

func (checker *Checker) declareCompositeConstructor(
	compositeDeclaration *ast.CompositeDeclaration,
	compositeType *CompositeType,
	parameterTypeAnnotations []*TypeAnnotation,
) {
	functionType := &ConstructorFunctionType{
		&FunctionType{
			ReturnTypeAnnotation: NewTypeAnnotation(
				compositeType,
			),
		},
	}

	var argumentLabels []string

	// TODO: support multiple overloaded initializers

	if len(compositeDeclaration.Members.Initializers) > 0 {
		firstInitializer := compositeDeclaration.Members.Initializers[0]

		argumentLabels = firstInitializer.Parameters.ArgumentLabels()

		functionType = &ConstructorFunctionType{
			FunctionType: &FunctionType{
				ParameterTypeAnnotations: parameterTypeAnnotations,
				ReturnTypeAnnotation:     NewTypeAnnotation(compositeType),
			},
		}

		checker.Elaboration.InitializerFunctionTypes[firstInitializer] = functionType
	}

	_, err := checker.valueActivations.DeclareFunction(
		compositeDeclaration.Identifier,
		functionType,
		argumentLabels,
	)
	checker.report(err)
}

func (checker *Checker) membersAndOrigins(
	fields []*ast.FieldDeclaration,
	functions []*ast.FunctionDeclaration,
	requireVariableKind bool,
) (
	members map[string]*Member,
	origins map[string]*Origin,
) {
	memberCount := len(fields) + len(functions)
	members = make(map[string]*Member, memberCount)
	origins = make(map[string]*Origin, memberCount)

	// declare a member for each field
	for _, field := range fields {
		fieldTypeAnnotation := checker.ConvertTypeAnnotation(field.TypeAnnotation)

		fieldType := fieldTypeAnnotation.Type

		checker.checkTypeAnnotation(fieldTypeAnnotation, field.TypeAnnotation.StartPos)

		identifier := field.Identifier.Identifier

		members[identifier] = &Member{
			Type:          fieldType,
			VariableKind:  field.VariableKind,
			IsInitialized: false,
		}

		origins[identifier] =
			checker.recordFieldDeclarationOrigin(field, fieldType)

		if requireVariableKind &&
			field.VariableKind == ast.VariableKindNotSpecified {

			checker.report(
				&InvalidVariableKindError{
					Kind:     field.VariableKind,
					StartPos: field.Identifier.Pos,
					EndPos:   field.Identifier.Pos,
				},
			)
		}
	}

	// declare a member for each function
	for _, function := range functions {
		functionType := checker.functionType(function.Parameters, function.ReturnTypeAnnotation)

		argumentLabels := function.Parameters.ArgumentLabels()

		identifier := function.Identifier.Identifier

		members[identifier] = &Member{
			Type:           functionType,
			VariableKind:   ast.VariableKindConstant,
			IsInitialized:  true,
			ArgumentLabels: argumentLabels,
		}

		origins[identifier] =
			checker.recordFunctionDeclarationOrigin(function, functionType)
	}

	return members, origins
}

func (checker *Checker) checkInitializers(
	initializers []*ast.InitializerDeclaration,
	fields []*ast.FieldDeclaration,
	containerType Type,
	containerDeclarationKind common.DeclarationKind,
	typeIdentifier string,
	initializerParameterTypeAnnotations []*TypeAnnotation,
	initializerKind initializerKind,
) {
	count := len(initializers)

	if count == 0 {
		checker.checkNoInitializerNoFields(fields, initializerKind, typeIdentifier)
		return
	}

	// TODO: check all initializers:
	//  parameter initializerParameterTypeAnnotations needs to be a slice

	initializer := initializers[0]
	checker.checkInitializer(
		initializer,
		fields,
		containerType,
		containerDeclarationKind,
		typeIdentifier,
		initializerParameterTypeAnnotations,
		initializerKind,
	)
}

// checkNoInitializerNoFields checks that if there are no initializers
// there are also no fields – otherwise the fields will be uninitialized.
// In interfaces this is allowed.
//
func (checker *Checker) checkNoInitializerNoFields(
	fields []*ast.FieldDeclaration,
	initializerKind initializerKind,
	typeIdentifier string,
) {
	if len(fields) == 0 || initializerKind == initializerKindInterface {
		return
	}

	// report error for first field
	firstField := fields[0]

	checker.report(
		&MissingInitializerError{
			TypeIdentifier: typeIdentifier,
			FirstFieldName: firstField.Identifier.Identifier,
			FirstFieldPos:  firstField.Identifier.Pos,
		},
	)
}

func (checker *Checker) checkInitializer(
	initializer *ast.InitializerDeclaration,
	fields []*ast.FieldDeclaration,
	containerType Type,
	containerDeclarationKind common.DeclarationKind,
	typeIdentifier string,
	initializerParameterTypeAnnotations []*TypeAnnotation,
	initializerKind initializerKind,
) {
	// NOTE: new activation, so `self`
	// is only visible inside initializer

	checker.valueActivations.Enter()
	defer checker.valueActivations.Leave()

	checker.declareSelfValue(containerType)

	// check the initializer is named properly
	identifier := initializer.Identifier.Identifier
	if identifier != InitializerIdentifier {
		checker.report(
			&InvalidInitializerNameError{
				Name: identifier,
				Pos:  initializer.StartPos,
			},
		)
	}

	functionType := &FunctionType{
		ParameterTypeAnnotations: initializerParameterTypeAnnotations,
		ReturnTypeAnnotation:     NewTypeAnnotation(&VoidType{}),
	}

	checker.checkFunction(
		initializer.Parameters,
		ast.Position{},
		functionType,
		initializer.FunctionBlock,
		true,
	)

	if initializerKind == initializerKindInterface &&
		initializer.FunctionBlock != nil {

		checker.checkInterfaceFunctionBlock(
			initializer.FunctionBlock,
			containerDeclarationKind,
			common.DeclarationKindInitializer,
		)
	}
}

func (checker *Checker) checkCompositeFunctions(
	functions []*ast.FunctionDeclaration,
	selfType *CompositeType,
) {
	for _, function := range functions {
		// NOTE: new activation, as function declarations
		// shouldn't be visible in other function declarations,
		// and `self` is is only visible inside function

		checker.valueActivations.WithScope(func() {

			checker.declareSelfValue(selfType)

			function.Accept(checker)
		})
	}
}

func (checker *Checker) declareSelfValue(selfType Type) {

	// NOTE: declare `self` one depth lower ("inside" function),
	// so it can't be re-declared by the function's parameters

	depth := checker.valueActivations.Depth() + 1

	self := &Variable{
		Kind:       common.DeclarationKindSelf,
		Type:       selfType,
		IsConstant: true,
		Depth:      depth,
		Pos:        nil,
	}
	checker.valueActivations.Set(SelfIdentifier, self)
	checker.recordVariableDeclarationOccurrence(SelfIdentifier, self)
}

// checkMemberIdentifiers checks the fields and functions are unique and aren't named `init`
//
func (checker *Checker) checkMemberIdentifiers(
	fields []*ast.FieldDeclaration,
	functions []*ast.FunctionDeclaration,
) {

	positions := map[string]ast.Position{}

	for _, field := range fields {
		checker.checkMemberIdentifier(
			field.Identifier,
			common.DeclarationKindField,
			positions,
		)
	}

	for _, function := range functions {
		checker.checkMemberIdentifier(
			function.Identifier,
			common.DeclarationKindFunction,
			positions,
		)
	}
}

func (checker *Checker) checkMemberIdentifier(
	identifier ast.Identifier,
	kind common.DeclarationKind,
	positions map[string]ast.Position,
) {
	name := identifier.Identifier
	pos := identifier.Pos

	if name == InitializerIdentifier {
		checker.report(
			&InvalidNameError{
				Name: name,
				Pos:  pos,
			},
		)
	}

	if previousPos, ok := positions[name]; ok {
		checker.report(
			&RedeclarationError{
				Name:        name,
				Pos:         pos,
				Kind:        kind,
				PreviousPos: &previousPos,
			},
		)
	} else {
		positions[name] = pos
	}
}

func (checker *Checker) VisitFieldDeclaration(field *ast.FieldDeclaration) ast.Repr {

	// NOTE: field type is already checked when determining composite function in `compositeType`

	panic(&errors.UnreachableError{})
}

func (checker *Checker) VisitInitializerDeclaration(initializer *ast.InitializerDeclaration) ast.Repr {

	// NOTE: already checked in `checkInitializer`

	panic(&errors.UnreachableError{})
}
