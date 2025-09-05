package types

import (
	"sync"
	"testing"

	"github.com/morph-l2/go-ethereum/eth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBatchFetcher implements the BatchFetcher interface for testing
type MockBatchFetcher struct {
	mock.Mock
}

func (m *MockBatchFetcher) GetRollupBatchByIndex(index uint64) (*eth.RPCRollupBatch, error) {
	args := m.Called(index)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*eth.RPCRollupBatch), args.Error(1)
}

func TestBatchCache(t *testing.T) {
	t.Run("Get non-existent batch - fetch from node", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		expectedBatch := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature"),
				},
			},
		}
		mockFetcher.On("GetRollupBatchByIndex", uint64(1)).Return(expectedBatch, nil)

		batch, ok := cache.Get(1)
		assert.True(t, ok)
		assert.Equal(t, expectedBatch, batch)

		mockFetcher.AssertExpectations(t)

		// Second get should use cache
		batch, ok = cache.Get(1)
		assert.True(t, ok)
		assert.Equal(t, expectedBatch, batch)
	})

	t.Run("Get non-existent batch - fetch fails", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		mockFetcher.On("GetRollupBatchByIndex", uint64(2)).Return(nil, assert.AnError).Once()

		batch, ok := cache.Get(2)
		assert.False(t, ok)
		assert.Nil(t, batch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Set and Get batch", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		batch := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature"),
				},
			},
		}

		// Add this line to set up the mock expectation
		mockFetcher.On("GetRollupBatchByIndex", uint64(3)).Return(batch, nil).Maybe()

		cache.Set(3, batch)

		gotBatch, ok := cache.Get(3)
		assert.True(t, ok)
		assert.Equal(t, batch, gotBatch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Delete batch", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		batch := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature"),
				},
			},
		}

		cache.Set(4, batch)
		gotBatch, ok := cache.Get(4)
		assert.True(t, ok)
		assert.Equal(t, batch, gotBatch)

		cache.Delete(4)

		// Setup mock for fetching after delete
		mockFetcher.On("GetRollupBatchByIndex", uint64(4)).Return(nil, assert.AnError).Once()

		gotBatch, ok = cache.Get(4)
		assert.False(t, ok)
		assert.Nil(t, gotBatch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Clear cache", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		batch1 := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature1"),
				},
			},
		}
		batch2 := &eth.RPCRollupBatch{
			Version: 2,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature2"),
				},
			},
		}

		cache.Set(5, batch1)
		cache.Set(6, batch2)

		cache.Clear()

		// Setup mocks for fetching after clear
		mockFetcher.On("GetRollupBatchByIndex", uint64(5)).Return(nil, assert.AnError).Once()
		mockFetcher.On("GetRollupBatchByIndex", uint64(6)).Return(nil, assert.AnError).Once()

		gotBatch, ok := cache.Get(5)
		assert.False(t, ok)
		assert.Nil(t, gotBatch)

		gotBatch, ok = cache.Get(6)
		assert.False(t, ok)
		assert.Nil(t, gotBatch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Concurrent access", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Pre-set a batch to avoid nil pointer in concurrent access
		testBatch := &eth.RPCRollupBatch{
			Version: 7,
			Signatures: []eth.RPCBatchSignature{
				{
					Signature: []byte("signature"),
				},
			},
		}
		cache.Set(7, testBatch)

		// Setup mock expectation to allow any number of calls
		mockFetcher.On("GetRollupBatchByIndex", uint64(7)).Return(testBatch, nil).Maybe()

		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				batch, ok := cache.Get(7)
				if ok && batch != nil {
					cache.Set(7, batch)
				}
			}()
		}

		wg.Wait()

		// Final validation of cache state
		batch, ok := cache.Get(7)
		assert.True(t, ok)
		assert.NotNil(t, batch)
		assert.Equal(t, testBatch.Version, batch.Version)
	})
}
