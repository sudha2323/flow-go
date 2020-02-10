package messages

import (
	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/model/flow"
)

type BlockProposal struct {
	Header  *flow.Header
	Payload *flow.Payload
}

type BlockVote struct {
	View      uint64
	BlockID   flow.Identifier
	Index     uint32
	Signature crypto.Signature
}
