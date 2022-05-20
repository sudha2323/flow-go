package streamer

import (
	"github.com/onflow/flow-go/module/component"
	"github.com/onflow/flow-go/module/state_synchronization"
)

type ExecutionDataStreamer interface {
	component.Component

	// RegisterExecutionDataConsumer registers a callback that will be called with the
	// Execution Data for every executed block. The callback will be called in the order
	// that blocks are sealed.
	RegisterExecutionDataConsumer(ExecutionDataConsumer)
}

type ExecutionDataConsumer = func(*state_synchronization.ExecutionData)
