package streamer

import (
	"github.com/onflow/flow-go/model/flow"
	"github.com/rs/zerolog"
)

type LogConsumer struct {
	log zerolog.Logger
}

var _LogStreamerClient = (*LogConsumer)(nil)

func NewLogConsumer() *LogConsumer {

	log := zerolog.Nop()
	lc := &LogConsumer{
		log: log,
	}
	return lc
}

func (lc *LogConsumer) onTest(b *flow.Block) {
	lc.log.Debug().Msg("on test")
}
