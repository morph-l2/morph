package types

import (
	"sync"

	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
)

type BatchCacheLegacy struct {
	m          sync.RWMutex
	batchCache map[uint64]*eth.RPCRollupBatch
	fetcher    iface.BatchFetcher
}

// NewBatchCacheLegacy creates a new batch cache instance
func NewBatchCacheLegacy(fetcher iface.BatchFetcher) *BatchCacheLegacy {
	return &BatchCacheLegacy{
		batchCache: make(map[uint64]*eth.RPCRollupBatch),
		fetcher:    fetcher,
	}
}

// Get retrieves a batch from the cache by its index
// If not found in cache, tries to fetch from node
func (b *BatchCacheLegacy) Get(batchIndex uint64) (*eth.RPCRollupBatch, bool) {
	// First try to get from cache
	b.m.RLock()
	batch, ok := b.batchCache[batchIndex]
	b.m.RUnlock()

	if ok {
		return batch, true
	}

	// If not in cache, try to fetch from node
	if b.fetcher != nil {
		fetchedBatch, err := b.fetcher.GetRollupBatchByIndex(batchIndex)
		if err != nil {
			log.Warn("Failed to fetch batch from node",
				"index", batchIndex,
				"error", err)
			return nil, false
		}

		if fetchedBatch != nil {
			b.m.Lock()
			b.batchCache[batchIndex] = fetchedBatch
			b.m.Unlock()

			return fetchedBatch, true
		}
	}

	return nil, false
}

func (b *BatchCacheLegacy) Set(batchIndex uint64, batch *eth.RPCRollupBatch) {
	if batch == nil {
		log.Debug("Refusing to cache nil batch", "batch_index", batchIndex)
		return
	}

	b.m.Lock()
	defer b.m.Unlock()

	b.batchCache[batchIndex] = batch
}

func (b *BatchCacheLegacy) Delete(batchIndex uint64) {
	b.m.Lock()
	defer b.m.Unlock()

	delete(b.batchCache, batchIndex)
}

// Clear removes all entries from the batch cache
func (bc *BatchCacheLegacy) Clear() {
	bc.m.Lock()
	defer bc.m.Unlock()
	bc.batchCache = make(map[uint64]*eth.RPCRollupBatch)
}
