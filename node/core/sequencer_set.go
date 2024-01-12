package node

import (
	"fmt"
)

type SequencerSetInfo struct {
	version      uint64                           // version in sequencer contract
	startHeight  uint64                           // start height from which the sequencers begin to participate in consensus process
	sequencerSet map[[tmKeySize]byte]sequencerKey // tendermint pk -> bls pk
}

func (ssi SequencerSetInfo) String() string {
	var sequencerKeys []string
	for k := range ssi.sequencerSet {
		sequencerKeys = append(sequencerKeys, fmt.Sprintf("%x", k))
	}
	return fmt.Sprintf("version: %d, startHeight: %d, sequencers: %v", ssi.version, ssi.startHeight, sequencerKeys)
}
