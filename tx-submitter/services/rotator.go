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

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
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
	indexer         *event.EventIndexer
}

func NewRotator(l2SeqencerAddr common.Address, indexer *event.EventIndexer, rollupEpoch uint64) *Rotator {
	return &Rotator{
		l2SequencerAddr: l2SeqencerAddr,
		indexer:         indexer,
		epoch:           big.NewInt(int64(rollupEpoch)),
	}
}

func (r *Rotator) StartEventIndexer() {
	go r.indexer.Index()
}

// UpdateState updates the state of the rotator
// updated by event listener in the future
func (r *Rotator) UpdateState(clients []iface.L2Client, l1Staking iface.IL1Staking) error {
	epochUpdateTime := new(big.Int)
	// sequencer set update time
	sequcerUpdateTime, err := GetSequencerSetUpdateTime(r.l2SequencerAddr, clients)
	if err != nil {
		log.Error("failed to get sequencer set update time", "err", err)
		return fmt.Errorf("GetCurrentSubmitter: failed to get sequencer set update time: %w", err)
	}

	storage := r.indexer.GetStorage()
	err = storage.Load()
	if err != nil {
		log.Error("failed to load storage", "err", err)
		return fmt.Errorf("GetCurrentSubmitter: failed to load storage: %w", err)
	}
	// if index not complete
	if storage.BlockProcessed() == 0 {
		return errors.New("wait event index service to complete")
	}

	r.startTime = utils.MaxOfThreeBig(epochUpdateTime, sequcerUpdateTime, big.NewInt(int64(storage.BlockTime())))

	// get current sequencer set
	seqSet, err := QuerySequencerSet(r.l2SequencerAddr, clients)
	if err != nil {
		log.Error("failed to get sequencer set", "err", err)
		return fmt.Errorf("UpdateState: failed to get sequencer set: %w", err)
	}
	r.SetSequencerSet(seqSet)

	// get l1staking active staker set
	stakers, err := l1Staking.GetActiveStakers(nil)
	if err != nil {
		log.Error("failed to get active stakers", "err", err)
		return fmt.Errorf("UpdateState: failed to get active stakers: %w", err)
	}
	submitterSet := utils.IntersectionOfAddresses(r.GetSequencerSet(), stakers)
	r.SetSubmitterSet(submitterSet)
	// rotator info
	log.Info(
		"rotator state updated",
		"epoch", r.epoch,
		"start_time", utils.FormatTime(r.startTime),
		"start_timestamp", r.startTime,
		"epoch_update_time", utils.FormatTime(epochUpdateTime),
		"epoch_update_timestamp", epochUpdateTime,
		"seq_update_time", utils.FormatTime(sequcerUpdateTime),
		"seq_update_timestamp", sequcerUpdateTime,
		"indexed_latest_block", storage.BlockProcessed(),
		"indexed_event_time", utils.FormatTime(big.NewInt(int64(storage.BlockTime()))),
		"indexed_event_timestamp", storage.BlockTime(),
	)

	return nil
}

// GetCurrentSubmitter returns the current sequencer that should be submitting
func (r *Rotator) CurrentSubmitter(clients []iface.L2Client, l1Staking iface.IL1Staking) (*common.Address, int64, error) {

	err := r.UpdateState(clients, l1Staking)
	if err != nil {
		return nil, 0, fmt.Errorf("update state err: %w", err)
	}

	if len(r.GetSubmitterSet()) == 0 {
		return nil, 0, fmt.Errorf("GetCurrentSubmitter: sequencer set is empty")
	}

	if r.epoch.Int64() == 0 {
		return nil, 0, fmt.Errorf("GetCurrentSubmitter: epoch is 0")
	}

	sec := time.Now().Unix() - r.startTime.Int64()
	seqIdx := sec / r.epoch.Int64() % int64(len(r.GetSequencerSet()))

	return &r.GetSubmitterSet()[seqIdx], seqIdx, nil

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
