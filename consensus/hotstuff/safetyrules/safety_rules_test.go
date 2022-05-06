package safetyrules

import (
	"errors"
	"github.com/onflow/flow-go/consensus/hotstuff"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/consensus/hotstuff/helper"
	"github.com/onflow/flow-go/consensus/hotstuff/mocks"
	"github.com/onflow/flow-go/utils/unittest"
)

func TestSafetyRules(t *testing.T) {
	//unittest.SkipUnless(t, unittest.TEST_TODO, "COMMITTEE_BY_VIEW - updating in next pr")
	//t.Run("should vote for bootstrapBlock", TestProduceVote_Ok)
	//t.Run("should not vote for unsafe bootstrapBlock", testUnsafe)
	//t.Run("should not vote for bootstrapBlock with its view below the current view", testBelowVote)
	//t.Run("should not vote for bootstrapBlock with its view above the current view", testAboveVote)
	//t.Run("should not vote for bootstrapBlock with the same view as the last voted view", testEqualLastVotedView)
	//t.Run("should not vote for bootstrapBlock with its view below the last voted view", testBelowLastVotedView)
	//t.Run("should not vote for the same view again", testVotingAgain)
	//t.Run("should not vote while not a committee member", testVotingWhileNonCommitteeMember)
	suite.Run(t, new(SafetyRulesTestSuite))
}

// SafetyRulesTestSuite is a test suite for testing SafetyRules related functionality.
// SafetyRulesTestSuite setups mocks for injected modules and creates hotstuff.SafetyData
// based on next configuration:
// R <- B[QC_R] <- P[QC_B]
// B.View + 1 = S.View + 2
// B - bootstrapped block, we are creating SafetyRules at block B
// Based on this HighestAcknowledgedView = B.View and
type SafetyRulesTestSuite struct {
	suite.Suite

	bootstrapBlock *model.Block
	proposal       *model.Proposal
	ourIdentity    *flow.Identity
	signer         *mocks.Signer
	persister      *mocks.Persister
	committee      *mocks.DynamicCommittee
	safetyData     *hotstuff.SafetyData
	safety         *SafetyRules
}

func (s *SafetyRulesTestSuite) SetupTest() {
	s.ourIdentity = unittest.IdentityFixture()
	s.signer = &mocks.Signer{}
	s.persister = &mocks.Persister{}
	s.committee = &mocks.DynamicCommittee{}
	s.committee.On("Self").Return(s.ourIdentity.NodeID).Maybe()
	s.committee.On("IdentityByBlock", mock.Anything, s.ourIdentity.NodeID).Return(s.ourIdentity, nil).Maybe()

	// bootstrap at random bootstrapBlock
	s.bootstrapBlock = helper.MakeBlock(helper.WithBlockView(100))
	s.proposal = helper.MakeProposal(
		helper.WithBlock(
			helper.MakeBlock(
				helper.WithParentBlock(s.bootstrapBlock),
				helper.WithBlockView(s.bootstrapBlock.View+1))))

	s.safetyData = &hotstuff.SafetyData{
		LockedOneChainView:      s.bootstrapBlock.View,
		HighestAcknowledgedView: s.bootstrapBlock.View,
		LastTimeout:             nil,
	}

	s.safety = New(s.signer, s.persister, s.committee, s.safetyData)
}

// TestShouldVote test basic happy path scenario where we vote for first block after bootstrap
func (s *SafetyRulesTestSuite) TestShouldVote() {
	expectedSafetyData := &hotstuff.SafetyData{
		LockedOneChainView:      s.proposal.Block.QC.View,
		HighestAcknowledgedView: s.proposal.Block.View,
		LastTimeout:             nil,
	}

	expectedVote := makeVote(s.proposal.Block)
	s.signer.On("CreateVote", s.proposal.Block).Return(expectedVote, nil).Once()
	s.persister.On("PutSafetyData", expectedSafetyData).Return(nil).Once()

	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), vote)
	require.Equal(s.T(), expectedVote, vote)

	s.persister.AssertCalled(s.T(), "PutSafetyData", expectedSafetyData)
}

// TestInvalidCurrentView tests that no vote is created if proposal is for invalid view
func (s *SafetyRulesTestSuite) TestInvalidCurrentView() {
	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View+1)
	require.Nil(s.T(), vote)
	require.Error(s.T(), err)
	s.persister.AssertNotCalled(s.T(), "PutSafetyData")
}

// TestNodeEjected tests that no vote is created if voter is ejected
func (s *SafetyRulesTestSuite) TestNodeEjected() {
	*s.committee = mocks.DynamicCommittee{}
	s.committee.On("Self").Return(s.ourIdentity.NodeID)
	s.committee.On("IdentityByBlock", s.proposal.Block.BlockID, s.ourIdentity.NodeID).Return(nil, model.NewInvalidSignerErrorf("node-ejected")).Once()

	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View)
	require.Nil(s.T(), vote)
	require.True(s.T(), model.IsNoVoteError(err))
	s.persister.AssertNotCalled(s.T(), "PutSafetyData")
}

// TestInvalidVoterIdentity tests that no vote is created if there was an exception retrieving voter identity
func (s *SafetyRulesTestSuite) TestInvalidVoterIdentity() {
	*s.committee = mocks.DynamicCommittee{}
	s.committee.On("Self").Return(s.ourIdentity.NodeID)
	exception := errors.New("invalid-signer-identity")
	s.committee.On("IdentityByBlock", s.proposal.Block.BlockID, s.ourIdentity.NodeID).Return(nil, exception).Once()

	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View)
	require.Nil(s.T(), vote)
	require.ErrorAs(s.T(), err, &exception)
	s.persister.AssertNotCalled(s.T(), "PutSaf etyData")
}

// TestCreateVoteException tests that no vote is created if vote creation raised an exception
func (s *SafetyRulesTestSuite) TestCreateVoteException() {
	exception := errors.New("create-vote-exception")
	s.signer.On("CreateVote", s.proposal.Block).Return(nil, exception).Once()
	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View)
	require.Nil(s.T(), vote)
	require.ErrorAs(s.T(), err, &exception)
	s.persister.AssertNotCalled(s.T(), "PutSafetyData")
}

// TestPersistStateException tests that no vote is created if persisting state failed
func (s *SafetyRulesTestSuite) TestPersistStateException() {
	exception := errors.New("persister-exception")
	s.persister.On("PutSafetyData", mock.Anything).Return(exception)

	vote := makeVote(s.proposal.Block)
	s.signer.On("CreateVote", s.proposal.Block).Return(vote, nil).Once()
	vote, err := s.safety.ProduceVote(s.proposal, s.proposal.Block.View)
	require.Nil(s.T(), vote)
	require.ErrorAs(s.T(), err, &exception)
}

// TestVotingOnUnsafeProposal tests different scenarios where we try to vote on unsafe blocks
func (s *SafetyRulesTestSuite) TestVotingOnUnsafeProposal() {
	s.Run("invalid-block-view", func() {
		// create block with block.View == block.QC.View
		proposal := helper.MakeProposal(
			helper.WithBlock(
				helper.MakeBlock(
					helper.WithParentBlock(s.bootstrapBlock),
					helper.WithBlockView(s.bootstrapBlock.View))))
		vote, err := s.safety.ProduceVote(proposal, proposal.Block.View)
		require.Error(s.T(), err)
		require.Nil(s.T(), vote)
	})
	s.Run("view-already-acknowledged", func() {
		// create block with view <= HighestAcknowledgedView
		proposal := helper.MakeProposal(
			helper.WithBlock(
				helper.MakeBlock(
					func(block *model.Block) {
						block.QC = helper.MakeQC(helper.WithQCView(s.safetyData.HighestAcknowledgedView - 1))
					},
					helper.WithBlockView(s.safetyData.HighestAcknowledgedView))))
		vote, err := s.safety.ProduceVote(proposal, proposal.Block.View)
		require.True(s.T(), model.IsNoVoteError(err))
		require.Nil(s.T(), vote)
	})
	s.Run("no-last-view-tc", func() {
		// create block where Block.View != Block.QC.View+1 and LastViewTC = nil
		proposal := helper.MakeProposal(
			helper.WithBlock(
				helper.MakeBlock(
					helper.WithParentBlock(s.bootstrapBlock),
					helper.WithBlockView(s.bootstrapBlock.View+2))))
		vote, err := s.safety.ProduceVote(proposal, proposal.Block.View)
		require.Error(s.T(), err)
		require.Nil(s.T(), vote)
	})
	s.Run("last-view-tc-invalid-view", func() {
		// create block where Block.View != Block.QC.View+1 and
		// Block.View != LastViewTC.View+1
		proposal := helper.MakeProposal(
			helper.WithBlock(
				helper.MakeBlock(
					helper.WithParentBlock(s.bootstrapBlock),
					helper.WithBlockView(s.bootstrapBlock.View+2))),
			helper.WithLastViewTC(
				helper.MakeTC(
					helper.WithTCView(s.bootstrapBlock.View))))
		vote, err := s.safety.ProduceVote(proposal, proposal.Block.View)
		require.Error(s.T(), err)
		require.Nil(s.T(), vote)
	})
	s.Run("last-view-tc-invalid-highest-qc", func() {
		// create block where Block.View != Block.QC.View+1 and
		// Block.View == LastViewTC.View+1 and Block.QC.View < LastViewTC.TOHighestQC.View
		// in this case block is not safe to extend since proposal is built on top of QC, which is lower
		// than QC presented in LastViewTC.
		TOHighestQC := helper.MakeQC(helper.WithQCView(s.bootstrapBlock.View + 1))
		proposal := helper.MakeProposal(
			helper.WithBlock(
				helper.MakeBlock(
					helper.WithParentBlock(s.bootstrapBlock),
					helper.WithBlockView(s.bootstrapBlock.View+2))),
			helper.WithLastViewTC(
				helper.MakeTC(
					helper.WithTCView(s.bootstrapBlock.View+1),
					helper.WithTCHighestQC(TOHighestQC))))
		vote, err := s.safety.ProduceVote(proposal, proposal.Block.View)
		require.Error(s.T(), err)
		require.Nil(s.T(), vote)
	})

	s.signer.AssertNotCalled(s.T(), "CreateVote")
	s.signer.AssertNotCalled(s.T(), "PutSafetyData")
}

//func createVoter(blockView uint64, lastVotedView uint64, isBlockSafe, isCommitteeMember bool) (*model.Proposal, *model.Vote, *SafetyRules) {
//	bootstrapBlock := helper.MakeBlock(helper.WithBlockView(blockView))
//	expectVote := makeVote(bootstrapBlock)
//
//	forks := &mocks.ForksReader{}
//	forks.On("IsSafeBlock", bootstrapBlock).Return(isBlockSafe)
//
//	persist := &mocks.Persister{}
//	persist.On("PutVoted", mock.Anything).Return(nil)
//
//	signer := &mocks.Signer{}
//	signer.On("CreateVote", mock.Anything).Return(expectVote, nil)
//
//	committee := &mocks.DynamicCommittee{}
//	me := unittest.IdentityFixture()
//	committee.On("Self").Return(me.NodeID, nil)
//	if isCommitteeMember {
//		committee.On("Identity", mock.Anything, me.NodeID).Return(me, nil)
//	} else {
//		committee.On("Identity", mock.Anything, me.NodeID).Return(nil, model.NewInvalidSignerErrorf(""))
//	}
//
//	voter := New(signer, forks, persist, committee, lastVotedView)
//	return &model.Proposal{
//		Block:   bootstrapBlock,
//		SigData: nil,
//	}, expectVote, voter
//}
//
//func TestProduceVote_Ok(t *testing.T) {
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(2), true, true
//
//	// create voter
//	bootstrapBlock, expectVote, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	// produce vote
//	vote, err := voter.ProduceVote(bootstrapBlock, curView)
//
//	require.NoError(t, err)
//	require.Equal(t, vote, expectVote)
//}
//
//func testUnsafe(t *testing.T) {
//	// create unsafe bootstrapBlock
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(2), false, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "not safe")
//	require.True(t, model.IsNoVoteError(err))
//}
//
//func testBelowVote(t *testing.T) {
//	// curView < blockView
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(2), uint64(2), true, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "expecting bootstrapBlock for current view")
//	require.False(t, model.IsNoVoteError(err))
//}
//
//func testAboveVote(t *testing.T) {
//	// curView > blockView
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(4), uint64(2), true, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "expecting bootstrapBlock for current view")
//	require.False(t, model.IsNoVoteError(err))
//}
//
//func testEqualLastVotedView(t *testing.T) {
//	// curView == lastVotedView
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(3), true, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "must be larger than the last voted view")
//	require.False(t, model.IsNoVoteError(err))
//}
//
//func testBelowLastVotedView(t *testing.T) {
//	// curView < lastVotedView
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(4), true, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "must be larger than the last voted view")
//	require.False(t, model.IsNoVoteError(err))
//}
//
//func testVotingAgain(t *testing.T) {
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(2), true, true
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	// produce vote
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//	require.NoError(t, err)
//
//	// produce vote again for the same view
//	_, err = voter.ProduceVote(bootstrapBlock, curView)
//	require.Error(t, err)
//	require.Contains(t, err.Error(), "must be larger than the last voted view")
//	require.False(t, model.IsNoVoteError(err))
//}
//
//func testVotingWhileNonCommitteeMember(t *testing.T) {
//	blockView, curView, lastVotedView, isBlockSafe, isCommitteeMember := uint64(3), uint64(3), uint64(2), true, false
//
//	// create voter
//	bootstrapBlock, _, voter := createVoter(blockView, lastVotedView, isBlockSafe, isCommitteeMember)
//
//	// produce vote
//	_, err := voter.ProduceVote(bootstrapBlock, curView)
//
//	require.Error(t, err)
//	require.True(t, model.IsNoVoteError(err))
//}
//
func makeVote(block *model.Block) *model.Vote {
	return &model.Vote{
		BlockID: block.BlockID,
		View:    block.View,
		SigData: nil, // signature doesn't matter in this test case
	}
}
