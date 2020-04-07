package execution_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/dapperlabs/flow-go/engine"
	execTestutil "github.com/dapperlabs/flow-go/engine/execution/testutil"
	"github.com/dapperlabs/flow-go/engine/testutil"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/model/messages"
	network "github.com/dapperlabs/flow-go/network/mock"
	"github.com/dapperlabs/flow-go/network/stub"
	"github.com/dapperlabs/flow-go/utils/unittest"
)

func TestSyncFlow(t *testing.T) {
	hub := stub.NewNetworkHub()

	colID := unittest.IdentityFixture(unittest.WithRole(flow.RoleCollection))
	conID := unittest.IdentityFixture(unittest.WithRole(flow.RoleConsensus))
	exe1ID := unittest.IdentityFixture(unittest.WithRole(flow.RoleExecution))
	exe2ID := unittest.IdentityFixture(unittest.WithRole(flow.RoleExecution))
	verID := unittest.IdentityFixture(unittest.WithRole(flow.RoleVerification))

	identities := flow.IdentityList{colID, conID, exe1ID, exe2ID, verID}

	// exe1 will sync from exe2 after exe2 executes transactions
	exeNode1 := testutil.ExecutionNode(t, hub, exe1ID, identities, 0)
	exeNode2 := testutil.ExecutionNode(t, hub, exe2ID, identities, 21)
	defer exeNode1.Done()
	defer exeNode2.Done()

	collectionNode := testutil.GenericNode(t, hub, colID, identities)
	verificationNode := testutil.GenericNode(t, hub, verID, identities)
	consensusNode := testutil.GenericNode(t, hub, conID, identities)

	genesis := flow.Genesis(identities)

	tx1 := execTestutil.DeployCounterContractTransaction()
	tx2 := execTestutil.CreateCounterTransaction()
	tx4 := execTestutil.AddToCounterTransaction()

	col1 := flow.Collection{Transactions: []*flow.TransactionBody{&tx1}}
	col2 := flow.Collection{Transactions: []*flow.TransactionBody{&tx2}}
	col4 := flow.Collection{Transactions: []*flow.TransactionBody{&tx4}}

	//collections := map[flow.Identifier]flow.Collection{
	//	col1.ID(): col1,
	//	col2.ID(): col2,
	//	col3.ID(): col3,
	//}

	//Create three blocks, with one tx each
	block1 := &flow.Block{
		Header: flow.Header{
			ParentID: genesis.ID(),
			View:     42,
			Height:   1,
		},
		Payload: flow.Payload{
			Guarantees: []*flow.CollectionGuarantee{
				{
					CollectionID: col1.ID(),
				},
			},
		},
	}

	block2 := &flow.Block{
		Header: flow.Header{
			ParentID: block1.ID(),
			View:     44,
			Height:   2,
		},
		Payload: flow.Payload{
			Guarantees: []*flow.CollectionGuarantee{
				{
					CollectionID: col2.ID(),
				},
			},
		},
	}

	block3 := &flow.Block{
		Header: flow.Header{
			ParentID: block2.ID(),
			View:     45,
			Height:   3,
		},
		Payload: flow.Payload{
		},
	}

	block4 := &flow.Block{
		Header: flow.Header{
			ParentID: block3.ID(),
			View:     46,
			Height:   4,
		},
		Payload: flow.Payload{
			Guarantees: []*flow.CollectionGuarantee{
				{
					CollectionID: col4.ID(),
				},
			},
		},
	}

	fmt.Printf("block1 ID %x parent %x\n", block1.ID(), block1.ParentID)
	fmt.Printf("block2 ID %x parent %x\n", block2.ID(), block2.ParentID)

	collectionEngine := new(network.Engine)
	colConduit, _ := collectionNode.Net.Register(engine.CollectionProvider, collectionEngine)

	collectionEngine.On("Submit", exe2ID.NodeID, mock.MatchedBy(func(r *messages.CollectionRequest) bool { return r.ID == col1.ID() })).Run(func(args mock.Arguments) {
		colConduit.Submit(&messages.CollectionResponse{Collection: col1,}, exe2ID.NodeID)
	}).Return(nil)
	collectionEngine.On("Submit", exe2ID.NodeID, mock.MatchedBy(func(r *messages.CollectionRequest) bool { return r.ID == col2.ID() })).Run(func(args mock.Arguments) {
		colConduit.Submit(&messages.CollectionResponse{Collection: col2,}, exe2ID.NodeID)
	}).Return(nil)
	collectionEngine.On("Submit", exe1ID.NodeID, &messages.CollectionRequest{ID: col4.ID()}).Run(func(args mock.Arguments) {
		colConduit.Submit(&messages.CollectionResponse{Collection: col4,}, exe1ID.NodeID)
	}).Return(nil)

	var receipt *flow.ExecutionReceipt


	executedBlocks := map[flow.Identifier]bool{}

	verificationEngine := new(network.Engine)
	_, _ = verificationNode.Net.Register(engine.ExecutionReceiptProvider, verificationEngine)
	verificationEngine.On("Submit", exe2ID.NodeID, mock.Anything).
		Run(func(args mock.Arguments) {
			receipt, _ = args[1].(*flow.ExecutionReceipt)
			executedBlocks[receipt.ExecutionResult.BlockID] = true
		}).
		Return(nil)

	consensusEngine := new(network.Engine)
	_, _ = consensusNode.Net.Register(engine.ExecutionReceiptProvider, consensusEngine)
	// should be both exe2 and exe1
	consensusEngine.On("Submit", exe2ID.NodeID, mock.Anything).Return(nil)

	// submit block from consensus node
	exeNode2.IngestionEngine.Submit(conID.NodeID, block1)
	//time.Sleep(200 * time.Millisecond)
	exeNode2.IngestionEngine.Submit(conID.NodeID, block2)

	// wait for block2 to be executed on execNode2
	hub.Eventually(t, func() bool {
		return executedBlocks[block2.ID()]
	})

	// make sure exe1 didn't get any blocks
	exeNode1.AssertHighestExecutedBlock(t, &genesis.Header)
	exeNode2.AssertHighestExecutedBlock(t, &block2.Header)

	// submit block3 and block4 to exe1 which should trigger sync
	exeNode1.IngestionEngine.Submit(conID.NodeID, block3)
	exeNode1.IngestionEngine.Submit(conID.NodeID, block4)

	//verificationEngine.On("Submit", exe1ID.NodeID, mock.MatchedBy(func(r *flow.ExecutionReceipt) bool { return r.ExecutionResult.BlockID == block3.ID() })).Return(nil)
	//consensusEngine.On("Submit", exe1ID.NodeID, mock.MatchedBy(func(r *flow.ExecutionReceipt) bool { return r.ExecutionResult.BlockID == block3.ID() })).Return(nil)

	// wait for block2 to be executed on execNode2
	hub.Eventually(t, func() bool {
		return executedBlocks[block4.ID()] && executedBlocks[block3.ID()]
	})

	collectionEngine.AssertExpectations(t)
	verificationEngine.AssertExpectations(t)
}
