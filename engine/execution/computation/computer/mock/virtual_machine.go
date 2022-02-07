// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	fvm "github.com/onflow/flow-go/fvm"
	mock "github.com/stretchr/testify/mock"

	programs "github.com/onflow/flow-go/fvm/programs"

	state "github.com/onflow/flow-go/fvm/state"
)

// VirtualMachine is an autogenerated mock type for the VirtualMachine type
type VirtualMachine struct {
	mock.Mock
}

// Run provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m VirtualMachine) Run(_a0 fvm.Context, _a1 fvm.Procedure, _a2 state.View, _a3 programs.Programs) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(fvm.Context, fvm.Procedure, state.View, programs.Programs) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
