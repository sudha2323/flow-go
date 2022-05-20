package streamer

import (
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/component"
)

type BlockStreamer interface {
	component.Component

	// RegisterBlockConsumer registers a callback that will be called for every new
	// finalized block. The callback will be called in the order of finalization.
	RegisterBlockConsumer(FinalizedBlockConsumer)
}

type FinalizedBlockConsumer = func(*flow.Block)
