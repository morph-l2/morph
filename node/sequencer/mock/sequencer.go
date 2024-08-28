package mock

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/log"

	node "morph-l2/node/core"
)

type Sequencer struct {
	engine *node.Executor
	exitCh chan struct{}
	stop   chan struct{}

	currentBlock int64
}

func NewSequencer(executor *node.Executor) (*Sequencer, error) {
	currentBlock, err := executor.L2Client().BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	return &Sequencer{
		engine:       executor,
		exitCh:       make(chan struct{}),
		stop:         make(chan struct{}),
		currentBlock: int64(currentBlock),
	}, nil
}

func (s *Sequencer) Start() {
	log.Info("starting simulated sequencer......")
	blockTicker := time.NewTicker(time.Second * 3)
	defer blockTicker.Stop()
	for {
		select {
		case <-blockTicker.C:
			log.Info("start to build new block", "block number", s.currentBlock+1)
			l2Data, err := s.engine.L2Client().AssembleL2Block(context.Background(), big.NewInt(s.currentBlock+1), nil)
			if err != nil {
				log.Error("error assembling block", "error", err)
				continue
			}
			if l2Data == nil || l2Data.Number == 0 {
				log.Info("Not now: no txs found")
				continue
			}
			pass, err := s.engine.L2Client().ValidateL2Block(context.Background(), l2Data)
			if err != nil {
				log.Error("error validating block", "error", err)
				continue
			}
			if !pass { // would NOT reach here
				log.Error("validating failed")
				continue
			}
			if err := s.engine.L2Client().NewL2Block(context.Background(), l2Data, nil); err != nil {
				log.Error("error occurs when creating l2 block", "error", err)
				continue
			}
			s.currentBlock++
			log.Info("successfully built block", "block number", s.currentBlock)
		case <-s.exitCh:
			close(s.stop)
			return
		}
	}

}

func (s *Sequencer) Stop() {
	log.Info("sequencer is stopping")
	close(s.exitCh)
	<-s.stop
	log.Info("sequencer is stopped")
}
