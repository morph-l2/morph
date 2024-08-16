package services

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
)

type Rotator struct {
	// used to record the start time of rotation
	// updated when the sequencer set or epoch is updated
	startTime    *big.Int // timestamp
	sequencerSet []common.Address
	epoch        *big.Int   // epoch for rotation
	mu           sync.Mutex // lock

	// addrs
	l2SequencerAddr common.Address
	l2GovAddr       common.Address
}

func NewRotator(l2SeqencerAddr, l2GovAddr common.Address) *Rotator {
	return &Rotator{
		l2SequencerAddr: l2SeqencerAddr,
		l2GovAddr:       l2GovAddr,
	}
}

// UpdateState updates the state of the rotator
// updated by event listener in the future
func (r *Rotator) UpdateState(clients []iface.L2Client) error {

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

	// start time
	if epochUpdateTime.Cmp(sequcerUpdateTime) > 0 {
		r.SetStartTime(epochUpdateTime)
	} else {
		r.SetStartTime(sequcerUpdateTime)
	}

	// get current sequencer set
	seqSet, err := GetSequencerSet(r.l2SequencerAddr, clients)
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
	r.SetEpoch(epoch)

	return nil
}

// GetCurrentSubmitter returns the current sequencer that should be submitting
func (r *Rotator) CurrentSubmitter(clients []iface.L2Client) (*common.Address, error) {

	err := r.UpdateState(clients)
	if err != nil {
		return nil, fmt.Errorf("update state err: %w", err)
	}

	if len(r.GetSequencerSet()) == 0 {
		return nil, fmt.Errorf("GetCurrentSubmitter: sequencer set is empty")
	}

	if r.GetEpoch().Int64() == 0 {
		return nil, fmt.Errorf("GetCurrentSubmitter: epoch is 0")
	}

	sec := time.Now().Unix() - r.GetStartTime().Int64()
	seqIdx := sec / r.GetEpoch().Int64() % int64(len(r.GetSequencerSet()))

	return &r.GetSequencerSet()[seqIdx], nil

}

func (r *Rotator) SetStartTime(newTime *big.Int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.startTime = newTime
}

func (r *Rotator) GetStartTime() *big.Int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.startTime
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

func (r *Rotator) SetEpoch(newEpoch *big.Int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.epoch = newEpoch
}

func (r *Rotator) GetEpoch() *big.Int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.epoch
}
