package batch

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"morph-l2/tx-submitter/db"

	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/log"
)

const (
	// Key prefixes for LevelDB storage
	SealedBatchKeyPrefix  = "sealed_batch_"
	SealedBatchIndicesKey = "sealed_batch_indices"
)

// BatchStorage handles persistence of sealed batches using JSON encoding
type BatchStorage struct {
	db db.Database
	mu sync.RWMutex
}

// NewBatchStorage creates a new BatchStorage instance
func NewBatchStorage(db db.Database) *BatchStorage {
	return &BatchStorage{
		db: db,
	}
}

// StoreSealedBatch stores a single sealed batch to LevelDB
// Uses JSON encoding for serialization
func (s *BatchStorage) StoreSealedBatch(batchIndex uint64, batch *eth.RPCRollupBatch) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Serialize batch to JSON
	encoded, err := json.Marshal(batch)
	if err != nil {
		return fmt.Errorf("failed to marshal sealed batch %d: %w", batchIndex, err)
	}

	// Store batch data
	key := encodeBatchKey(batchIndex)
	if err := s.db.PutBytes(key, encoded); err != nil {
		return fmt.Errorf("failed to store sealed batch %d: %w", batchIndex, err)
	}

	// Update indices list
	if err := s.updateBatchIndices(batchIndex, true); err != nil {
		log.Warn("Failed to update batch indices", "batch_index", batchIndex, "error", err)
		// Don't fail the operation if indices update fails
	}

	return nil
}

// LoadSealedBatch loads a single sealed batch from LevelDB
func (s *BatchStorage) LoadSealedBatch(batchIndex uint64) (*eth.RPCRollupBatch, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	key := encodeBatchKey(batchIndex)
	encoded, err := s.db.GetBytes(key)
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			return nil, fmt.Errorf("sealed batch %d not found", batchIndex)
		}
		return nil, fmt.Errorf("failed to get sealed batch %d: %w", batchIndex, err)
	}

	// Deserialize from JSON
	var batch eth.RPCRollupBatch
	if err := json.Unmarshal(encoded, &batch); err != nil {
		return nil, fmt.Errorf("failed to unmarshal sealed batch %d: %w", batchIndex, err)
	}

	return &batch, nil
}

// LoadAllSealedBatches loads all sealed batches from LevelDB
// Returns a map of batchIndex -> RPCRollupBatch
func (s *BatchStorage) LoadAllSealedBatches() (map[uint64]*eth.RPCRollupBatch, []uint64, error) {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			// No batches stored yet
			return make(map[uint64]*eth.RPCRollupBatch), nil, nil
		}
		return nil, nil, fmt.Errorf("failed to load batch indices: %w", err)
	}

	// Load each batch (without holding the lock to avoid deadlock)
	batches := make(map[uint64]*eth.RPCRollupBatch, len(indices))
	for _, idx := range indices {
		batch, err := s.LoadSealedBatch(idx)
		if err != nil {
			log.Warn("Failed to load sealed batch, skipping",
				"batch_index", idx, "error", err)
			continue
		}
		batches[idx] = batch
	}

	return batches, indices, nil
}

// DeleteSealedBatch removes a sealed batch from LevelDB
func (s *BatchStorage) DeleteSealedBatch(batchIndex uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	key := encodeBatchKey(batchIndex)
	if err := s.db.Delete(key); err != nil {
		return fmt.Errorf("failed to delete sealed batch %d: %w", batchIndex, err)
	}

	// Update indices list
	if err := s.updateBatchIndices(batchIndex, false); err != nil {
		log.Warn("Failed to update batch indices after deletion",
			"batch_index", batchIndex, "error", err)
		// Don't fail the operation if indices update fails
	}

	return nil
}

func (s *BatchStorage) DeleteAllSealedBatches() error {
	s.mu.RLock()
	// Load batch indices
	indices, err := s.loadBatchIndices()
	s.mu.RUnlock()
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			// No batches stored yet
			return nil
		}
		return fmt.Errorf("failed to load batch indices: %w", err)
	}

	for _, idx := range indices {
		err = s.DeleteSealedBatch(idx)
		if err != nil {
			log.Error("Failed to delete sealed batch",
				"batch_index", idx, "error", err)
			return err
		}
	}

	return nil
}

// encodeBatchKey encodes batch index to a byte key
func encodeBatchKey(batchIndex uint64) []byte {
	key := make([]byte, len(SealedBatchKeyPrefix)+8)
	copy(key, SealedBatchKeyPrefix)
	binary.BigEndian.PutUint64(key[len(SealedBatchKeyPrefix):], batchIndex)
	return key
}

// updateBatchIndices updates the list of stored batch indices
// add: true to add index, false to remove
func (s *BatchStorage) updateBatchIndices(batchIndex uint64, add bool) error {
	indices, err := s.loadBatchIndices()
	if err != nil {
		if errors.Is(err, db.ErrKeyNotFound) {
			indices = []uint64{}
		} else {
			return err
		}
	}

	if add {
		// Add index if not already present
		found := false
		for _, idx := range indices {
			if idx == batchIndex {
				found = true
				break
			}
		}
		if !found {
			indices = append(indices, batchIndex)
		}
	} else {
		// Remove index
		newIndices := make([]uint64, 0, len(indices))
		for _, idx := range indices {
			if idx != batchIndex {
				newIndices = append(newIndices, idx)
			}
		}
		indices = newIndices
	}

	return s.saveBatchIndices(indices)
}

// loadBatchIndices loads the list of stored batch indices
func (s *BatchStorage) loadBatchIndices() ([]uint64, error) {
	encoded, err := s.db.GetBytes([]byte(SealedBatchIndicesKey))
	if err != nil {
		return nil, err
	}

	var indices []uint64
	if err := json.Unmarshal(encoded, &indices); err != nil {
		return nil, fmt.Errorf("failed to unmarshal batch indices: %w", err)
	}

	return indices, nil
}

// saveBatchIndices saves the list of batch indices
func (s *BatchStorage) saveBatchIndices(indices []uint64) error {
	encoded, err := json.Marshal(indices)
	if err != nil {
		return fmt.Errorf("failed to marshal batch indices: %w", err)
	}

	return s.db.PutBytes([]byte(SealedBatchIndicesKey), encoded)
}
