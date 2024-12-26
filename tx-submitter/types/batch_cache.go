package types

import (
	"sync"

	"github.com/morph-l2/go-ethereum/eth"
)

type BatchCache struct {
	batchCache map[uint64]*eth.RPCRollupBatch
	m          sync.Mutex
}

func NewBatchCache() *BatchCache {
	return &BatchCache{
		batchCache: make(map[uint64]*eth.RPCRollupBatch),
	}
}

func (b *BatchCache) Get(batchIndex uint64) (*eth.RPCRollupBatch, bool) {
	b.m.Lock()
	defer b.m.Unlock()

	batch, ok := b.batchCache[batchIndex]
	return batch, ok
}
func (b *BatchCache) Set(batchIndex uint64, batch *eth.RPCRollupBatch) {
	b.m.Lock()
	defer b.m.Unlock()

	b.batchCache[batchIndex] = batch
}

func (b *BatchCache) Delete(batchIndex uint64) {
	b.m.Lock()
	defer b.m.Unlock()

	delete(b.batchCache, batchIndex)
}
