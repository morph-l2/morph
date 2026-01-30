package batch

import (
	"testing"

	"github.com/stretchr/testify/require"
	"morph-l2/bindings/bindings"
)

func init() {
	var err error
	rollupContract, err = bindings.NewRollup(rollupAddr, l1Client)
	if err != nil {
		panic(err)
	}
	sequencerContract, err = bindings.NewSequencer(sequencerAddr, l2Client)
	if err != nil {
		panic(err)
	}
	l2MessagePasserContract, err = bindings.NewL2ToL1MessagePasser(l2MessagePasserAddr, l2Client)
	if err != nil {
		panic(err)
	}
	govContract, err = bindings.NewGov(govAddr, l2Client)
	if err != nil {
		panic(err)
	}
}

func TestBatchCacheInit(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, l2Client, rollupContract, sequencerContract, l2MessagePasserContract, govContract)
	err := cache.InitFromRollup()
	require.NoError(t, err)
}

func TestBatchCacheInitByBlockRange(t *testing.T) {
	cache := NewBatchCache(nil, l1Client, l2Client, rollupContract, sequencerContract, l2MessagePasserContract, govContract)
	err := cache.InitFromRollupByRange()
	require.NoError(t, err)
}
