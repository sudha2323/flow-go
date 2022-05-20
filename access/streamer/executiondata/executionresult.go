package executiondata

import (
	"sync"

	"github.com/onflow/flow-go/access/streamer"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/state_synchronization"
)

type ExecutionDataDistributor struct {
	address                string
	blockID                []byte
	executionDataConsumers []streamer.ExecutionDataConsumer
	lock                   sync.RWMutex
	unit                   *engine.Unit
}

// Done implements streamer.BlockStreamer
func (e *ExecutionDataDistributor) Done() <-chan struct{} {
	return e.unit.Done()
}

// Ready implements streamer.BlockStreamer
func (e *ExecutionDataDistributor) Ready() <-chan struct{} {
	//panic("unimplemented")
	return e.unit.Ready()
}

// Start implements streamer.BlockStreamer
func (*ExecutionDataDistributor) Start(irrecoverable.SignalerContext) {
	panic("unimplemented")
}

var _ streamer.ExecutionDataStreamer = (*ExecutionDataDistributor)(nil)

func NewExecutionDataDistributor(serverAddr string, BlockID []byte) *ExecutionDataDistributor {
	return &ExecutionDataDistributor{
		address: serverAddr, blockID: BlockID,
		executionDataConsumers: make([]streamer.ExecutionDataConsumer, 0),
		lock:                   sync.RWMutex{},
		unit:                   engine.NewUnit(),
	}
}

func (p *ExecutionDataDistributor) RegisterExecutionDataConsumer(consumer streamer.ExecutionDataConsumer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.executionDataConsumers = append(p.executionDataConsumers, consumer)
}

func (p *ExecutionDataDistributor) OnExecutedBlock(data *state_synchronization.ExecutionData) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	for _, consumer := range p.executionDataConsumers {
		consumer(data)
	}
}
