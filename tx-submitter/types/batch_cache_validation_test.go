package types

import (
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/stretchr/testify/assert"
)

func TestBatchValidation(t *testing.T) {
	t.Run("Get - Valid batch with signatures is cached", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Create valid batch with signatures
		validBatch := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signer:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
					Signature: []byte("test-signature"),
				},
			},
		}

		mockFetcher.On("GetRollupBatchByIndex", uint64(1)).Return(validBatch, nil).Once()

		// Get should return the batch and cache it
		batch, ok := cache.Get(1)
		assert.True(t, ok)
		assert.Equal(t, validBatch, batch)
		assert.Equal(t, 1, len(batch.Signatures))

		// Second get should use cache without calling fetcher
		batch, ok = cache.Get(1)
		assert.True(t, ok)
		assert.Equal(t, validBatch, batch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Get - Invalid batch without signatures is not cached", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Create invalid batch without signatures
		invalidBatch := &eth.RPCRollupBatch{
			Version:    1,
			Signatures: []eth.RPCBatchSignature{}, // Empty signatures
		}

		mockFetcher.On("GetRollupBatchByIndex", uint64(2)).Return(invalidBatch, nil).Once()
		mockFetcher.On("GetRollupBatchByIndex", uint64(2)).Return(invalidBatch, nil).Once() // Second call because not cached

		// Get should return the batch but not cache it
		batch, ok := cache.Get(2)
		assert.True(t, ok) // Still returns true because batch was found, just not cached
		assert.Equal(t, invalidBatch, batch)
		assert.Equal(t, 0, len(batch.Signatures))

		// Second get should call fetcher again since it wasn't cached
		batch, ok = cache.Get(2)
		assert.True(t, ok)
		assert.Equal(t, invalidBatch, batch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Set - Valid batch with signatures is stored", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Create valid batch with signatures
		validBatch := &eth.RPCRollupBatch{
			Version: 1,
			Signatures: []eth.RPCBatchSignature{
				{
					Signer:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
					Signature: []byte("test-signature"),
				},
			},
		}

		// Set should store the batch
		cache.Set(3, validBatch)

		// Get should retrieve from cache
		batch, ok := cache.Get(3)
		assert.True(t, ok)
		assert.Equal(t, validBatch, batch)
	})

	t.Run("Set - Invalid batch without signatures is not stored", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Create invalid batch without signatures
		invalidBatch := &eth.RPCRollupBatch{
			Version:    1,
			Signatures: []eth.RPCBatchSignature{}, // Empty signatures
		}

		// Set should not store the batch
		cache.Set(4, invalidBatch)

		// Setup mock for fetching since batch shouldn't be in cache
		mockFetcher.On("GetRollupBatchByIndex", uint64(4)).Return(nil, assert.AnError).Once()

		// Get should try to fetch from node and fail
		batch, ok := cache.Get(4)
		assert.False(t, ok)
		assert.Nil(t, batch)

		mockFetcher.AssertExpectations(t)
	})

	t.Run("Set - Nil batch is not stored", func(t *testing.T) {
		mockFetcher := new(MockBatchFetcher)
		cache := NewBatchCache(mockFetcher)

		// Set with nil batch should not store anything
		cache.Set(5, nil)

		// Setup mock for fetching since nothing should be in cache
		mockFetcher.On("GetRollupBatchByIndex", uint64(5)).Return(nil, assert.AnError).Once()

		// Get should try to fetch from node and fail
		batch, ok := cache.Get(5)
		assert.False(t, ok)
		assert.Nil(t, batch)

		mockFetcher.AssertExpectations(t)
	})
}
