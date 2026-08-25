package types

import (
	"testing"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/require"
)

type fixedBatchFetcher struct {
	batch *eth.RPCRollupBatch
}

func (f fixedBatchFetcher) GetRollupBatchByIndex(uint64) (*eth.RPCRollupBatch, error) {
	return f.batch, nil
}

func TestBatchCacheLegacyCachesBatchWithoutSignatures(t *testing.T) {
	want := &eth.RPCRollupBatch{Version: 1}
	cache := NewBatchCacheLegacy(fixedBatchFetcher{batch: want})

	got, ok := cache.Get(3)
	require.True(t, ok)
	require.Same(t, want, got)

	cache.fetcher = nil
	got, ok = cache.Get(3)
	require.True(t, ok)
	require.Same(t, want, got)
}
