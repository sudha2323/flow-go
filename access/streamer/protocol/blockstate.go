package protocol

import (
	"sync"

	"github.com/onflow/flow-go/access/streamer"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/irrecoverable"
)

type BlockStateDistributor struct {
	address                 string
	startHeight             uint64
	blockFinalizedConsumers []streamer.FinalizedBlockConsumer
	lock                    sync.RWMutex
	unit                    *engine.Unit
}

// Done implements streamer.BlockStreamer
func (e *BlockStateDistributor) Done() <-chan struct{} {
	return e.unit.Done()
}

// Ready implements streamer.BlockStreamer
func (e *BlockStateDistributor) Ready() <-chan struct{} {
	//panic("unimplemented")
	return e.unit.Ready()
}

// Start implements streamer.BlockStreamer
func (*BlockStateDistributor) Start(irrecoverable.SignalerContext) {
	panic("unimplemented")
}

var _ streamer.BlockStreamer = (*BlockStateDistributor)(nil)

func NewBlockStateDistributor(serverAddr string, startHeight uint64) *BlockStateDistributor {
	return &BlockStateDistributor{
		address: serverAddr, startHeight: startHeight,
		blockFinalizedConsumers: make([]streamer.FinalizedBlockConsumer, 0),
		lock:                    sync.RWMutex{},
		unit:                    engine.NewUnit(),
	}
}

func (p *BlockStateDistributor) RegisterBlockConsumer(consumer streamer.FinalizedBlockConsumer) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.blockFinalizedConsumers = append(p.blockFinalizedConsumers, consumer)
}

func OnFinalizedBlock(p *BlockStateDistributor, block *flow.Block) {
	p.lock.RLock()
	defer p.lock.RUnlock()
	for _, consumer := range p.blockFinalizedConsumers {
		consumer(block)
	}
}
