package types

import (
	"sync"

	"morph-l2/tx-submitter/iface"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
)

type BatchCache struct {
	m          sync.RWMutex
	batchCache map[uint64]*eth.RPCRollupBatch
	fetcher    iface.BatchFetcher
}

// NewBatchCache creates a new batch cache instance
func NewBatchCache(fetcher iface.BatchFetcher) *BatchCache {
	return &BatchCache{
		batchCache: make(map[uint64]*eth.RPCRollupBatch),
		fetcher:    fetcher,
	}
}

// Get retrieves a batch from the cache by its index
// If not found in cache, tries to fetch from node
func (b *BatchCache) Get(batchIndex uint64) (*eth.RPCRollupBatch, bool) {
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

		// Validate batch before caching - batch must exist and have signatures
		if fetchedBatch != nil && len(fetchedBatch.Signatures) > 0 {
			// Store valid batch in cache for future use
			b.m.Lock()
			b.batchCache[batchIndex] = fetchedBatch
			b.m.Unlock()

			return fetchedBatch, true
		} else if fetchedBatch != nil {
			// Batch exists but doesn't have signatures, don't cache it
			log.Debug("Batch validation failed - no signatures",
				"batch_index", batchIndex,
				"found", fetchedBatch != nil,
				"has_signatures", len(fetchedBatch.Signatures) > 0)
			return fetchedBatch, true
		}
	}

	return nil, false
}

func (b *BatchCache) Set(batchIndex uint64, batch *eth.RPCRollupBatch) {
	// Validate batch before caching - batch must exist and have signatures
	if batch == nil || len(batch.Signatures) == 0 {
		log.Debug("Refusing to cache invalid batch",
			"batch_index", batchIndex,
			"exists", batch != nil,
			"has_signatures", batch != nil && len(batch.Signatures) > 0)
		return
	}

	b.m.Lock()
	defer b.m.Unlock()

	b.batchCache[batchIndex] = batch
}

func (b *BatchCache) Delete(batchIndex uint64) {
	b.m.Lock()
	defer b.m.Unlock()

	delete(b.batchCache, batchIndex)
}

// Clear removes all entries from the batch cache
func (bc *BatchCache) Clear() {
	bc.m.Lock()
	defer bc.m.Unlock()
	bc.batchCache = make(map[uint64]*eth.RPCRollupBatch)
}
