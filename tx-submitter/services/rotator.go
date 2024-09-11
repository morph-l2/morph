package services

import (
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"morph-l2/tx-submitter/event"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/utils"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/log"
)

type Rotator struct {
	// used to record the start time of rotation
	// updated when the sequencerSet, stakingSet or epoch is updated
	startTime       *big.Int // timestamp
	submitterSet    []common.Address
	sequencerSet    []common.Address
	epoch           *big.Int   // epoch for rotation
	mu              sync.Mutex // lock
	l2SequencerAddr common.Address
	l2GovAddr       common.Address
	indexer         *event.EventIndexer
}

func NewRotator(l2SeqencerAddr, l2GovAddr common.Address, indexer *event.EventIndexer) *Rotator {
	return &Rotator{
		l2SequencerAddr: l2SeqencerAddr,
		l2GovAddr:       l2GovAddr,
		indexer:         indexer,
	}
}

func (r *Rotator) StartEventIndexer() {
	go r.indexer.Index()
}

// UpdateState updates the state of the rotator
// updated by event listener in the future
func (r *Rotator) UpdateState(clients []iface.L2Client, l1Staking iface.IL1Staking) error {

	epochUpdateTime, err := GetEpochUpdateTime(r.l2GovAddr, clients)
	if err != nil {
		log.Error("failed to get epoch update time", "err", err)
		return fmt.Errorf("GetCurrentSubmitter: failed to get epoch update time: %w", err)
	}
	// sequencer set update time
	sequcerUpdateTime, err := GetSequencerSetUpdateTime(r.l2SequencerAddr, clients)
	if err != nil {
		log.Error("failed to get sequencer set update time", "err", err)
		return fmt.Errorf("GetCurrentSubmitter: failed to get sequencer set update time: %w", err)
	}

	storage := event.NewEventInfoStorage(r.indexer.GetStorePath())
	err = storage.Load()
	if err != nil {
		log.Error("failed to load storage", "err", err)
		return fmt.Errorf("GetCurrentSubmitter: failed to load storage: %w", err)
	}
	// if index not complete
	if storage.BlockProcessed == 0 {
		return errors.New("wait event index service to complete")
	}

	r.startTime = utils.MaxOfThreeBig(epochUpdateTime, sequcerUpdateTime, big.NewInt(int64(storage.BlockTime)))

	// get current sequencer set
	seqSet, err := QuerySequencerSet(r.l2SequencerAddr, clients)
	if err != nil {
		log.Error("failed to get sequencer set", "err", err)
		return fmt.Errorf("UpdateState: failed to get sequencer set: %w", err)
	}
	r.SetSequencerSet(seqSet)
	// get current epoch
	epoch, err := GetEpoch(r.l2GovAddr, clients)
	if err != nil {
		log.Error("failed to get epoch", "err", err)
		return err
	}
	r.epoch = epoch

	// get l1staking active staker set
	stakers, err := l1Staking.GetActiveStakers(nil)
	if err != nil {
		log.Error("failed to get active stakers", "err", err)
		return fmt.Errorf("UpdateState: failed to get active stakers: %w", err)
	}
	submitterSet := utils.IntersectionOfAddresses(r.GetSequencerSet(), stakers)
	r.SetSubmitterSet(submitterSet)

	return nil
}

// GetCurrentSubmitter returns the current sequencer that should be submitting
func (r *Rotator) CurrentSubmitter(clients []iface.L2Client, l1Staking iface.IL1Staking) (*common.Address, error) {

	err := r.UpdateState(clients, l1Staking)
	if err != nil {
		return nil, fmt.Errorf("update state err: %w", err)
	}

	if len(r.GetSubmitterSet()) == 0 {
		return nil, fmt.Errorf("GetCurrentSubmitter: sequencer set is empty")
	}

	if r.epoch.Int64() == 0 {
		return nil, fmt.Errorf("GetCurrentSubmitter: epoch is 0")
	}

	sec := time.Now().Unix() - r.startTime.Int64()
	seqIdx := sec / r.epoch.Int64() % int64(len(r.GetSequencerSet()))

	return &r.GetSubmitterSet()[seqIdx], nil

}

func (r *Rotator) SetSubmitterSet(newSet []common.Address) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.submitterSet = newSet
}
func (r *Rotator) GetSubmitterSet() []common.Address {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.submitterSet
}

func (r *Rotator) SetSequencerSet(newSet []common.Address) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sequencerSet = newSet
}

func (r *Rotator) GetSequencerSet() []common.Address {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.sequencerSet

}
