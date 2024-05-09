package services

import (
	"sync"

	"github.com/morph-l2/tx-submitter/iface"
	"github.com/scroll-tech/go-ethereum/common"
)

type Rotator struct {
	// used to record the start time of rotation
	// updated when the sequencer set or epoch is updated
	startTime    uint64 // timestamp
	sequencerSet []common.Address
	epoch        uint64
	mu           sync.Mutex // 互斥锁
}

func NewSubmitterRotator() *Rotator {
	return &Rotator{}
}

func (r *Rotator) CurrentSubmitter(clients []iface.L2Client) (*common.Address, error) {
	// check

	GetEpochUpdateTime()

	// get current sequencer set
	// get current epoch
	// get current time
	// calc the index of the sequencer in the sequencer set
	return nil, nil

}

func (r *Rotator) SetStartTime(newTime uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.startTime = newTime
}

func (r *Rotator) GetStartTime() uint64 {
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
