package hotstuff

import (
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/state/protocol"
)

// Committee accounts for the fact that we might have multiple HotStuff instances
// (collector committees and main consensus committee). Each hostuff instance is supposed to
// have a dedicated Committee state.
// A Committee provides subset of the protocol.State, which is restricted to exactly those
// nodes that participate in the current HotStuff instance: the state of all legitimate HotStuff
// participants for the specified block. Legitimate HotStuff participants have NON-ZERO STAKE.
//
// The intended use case is to support collectors running HotStuff within Flow. Specifically,
// the collectors produced their own blocks, independently of the Consensus Nodes (aka the main consensus).
// Given a collector block, some logic is required to find the main consensus block
// for determining the valid collector-HotStuff participants.
type Committee interface {

	// Identity returns the full Identity for specified HotStuff participant.
	// The node must be a legitimate HotStuff participant with NON-ZERO STAKE at the specified block.
	// ERROR conditions:
	//  * model.InvalidSignerError if participantID does NOT correspond to a _staked_ HotStuff participant at the specified block.
	Identity(blockID flow.Identifier, participantID flow.Identifier) (*flow.Identity, error)

	// Identities returns a IdentityList with legitimate HotStuff participants for the specified block.
	// The returned list of HotStuff participants
	//   * contains nodes that are allowed to sign the specified block (legitimate participants with NON-ZERO STAKE)
	//   * is ordered in the canonical order
	//   * contains no duplicates.
	Identities(blockID flow.Identifier) (flow.IdentityList, error)

	// LeaderForView returns the identity of the leader for a given view.
	// CAUTION: per liveness requirement of HotStuff, the leader must be fork-independent.
	//          Therefore, a node retains its proposer view slots even if it is slashed.
	//          Its proposal is simply considered invalid, as it is not from a legitimate participant.
	// Returns the following expected errors for invalid inputs:
	//  * epoch containing the requested view has not been set up (protocol.ErrNextEpochNotSetup)
	//  * epoch is too far in the past (leader.InvalidViewError)
	LeaderForView(view uint64) (flow.Identifier, error)

	// Self returns our own node identifier.
	// TODO: ultimately, the own identity of the node is necessary for signing.
	//       Ideally, we would move the method for checking whether an Identifier refers to this node to the signer.
	//       This would require some refactoring of EventHandler (postponed to later)
	Self() flow.Identifier

	// DKG returns the DKG info for the given block.
	DKG(blockID flow.Identifier) (DKG, error)
}

type DKG interface {
	protocol.DKG
}

// ComputeStakeThresholdForBuildingQC returns the stake that is minimally required for building a QC
func ComputeStakeThresholdForBuildingQC(totalStake uint64) uint64 {
	// Given totalStake, we need the smallest integer t such that 2 * totalStake / 3 < t
	// Formally, the minimally required stake is: 2 * Floor(totalStake/3) + max(1, totalStake mod 3)
	floorOneThird := totalStake / 3 // integer division, includes floor
	res := 2 * floorOneThird
	divRemainder := totalStake % 3
	if divRemainder <= 1 {
		res = res + 1
	} else {
		res += divRemainder
	}
	return res
}
