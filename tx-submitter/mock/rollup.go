package mock

import (
	"errors"
	"math/big"

	"github.com/morph-l2/go-ethereum/core/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"

	"morph-l2/bindings/bindings"
)

// MockRollup implements the IRollup interface for testing
type MockRollup struct {
	lastCommittedBatchIndex *big.Int
	lastFinalizedBatchIndex *big.Int
	legacyCutoverBatchIndex *big.Int
	insideChallengeWindow   bool
	batchExists             bool
	storedBlobHash          [32]byte
	finalizeTx              *types.Transaction
}

// NewMockRollup creates a new instance of MockRollup
func NewMockRollup() *MockRollup {
	return &MockRollup{
		lastCommittedBatchIndex: big.NewInt(0),
		lastFinalizedBatchIndex: big.NewInt(0),
		legacyCutoverBatchIndex: big.NewInt(0),
		insideChallengeWindow:   false,
		batchExists:             false,
	}
}

// LastCommittedBatchIndex implements IRollup
func (m *MockRollup) LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	return m.lastCommittedBatchIndex, nil
}

// LastFinalizedBatchIndex implements IRollup
func (m *MockRollup) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	return m.lastFinalizedBatchIndex, nil
}

// LegacyCutoverBatchIndex implements IRollup.
func (m *MockRollup) LegacyCutoverBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	return m.legacyCutoverBatchIndex, nil
}

// FinalizeBatch implements IRollup
func (m *MockRollup) FinalizeBatch(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return m.finalizeTx, nil
}

// BatchInsideChallengeWindow implements IRollup
func (m *MockRollup) BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	return m.insideChallengeWindow, nil
}

// BatchExist implements IRollup
func (m *MockRollup) BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	return m.batchExists, nil
}

// CommittedBatches implements IRollup
func (m *MockRollup) CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	return [32]byte{}, nil
}

// BatchBlobVersionedHashes implements IRollup (no stored hash by default)
func (m *MockRollup) BatchBlobVersionedHashes(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	return m.storedBlobHash, nil
}

// BatchDataStore implements IRollup
func (m *MockRollup) BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	return struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}{
		OriginTimestamp:   big.NewInt(0),
		FinalizeTimestamp: big.NewInt(0),
		BlockNumber:       big.NewInt(0),
	}, nil
}

// FilterCommitBatch implements IRollup
func (m *MockRollup) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupCommitBatchIterator, error) {
	return nil, errors.New("FilterCommitBatch not implemented in mock")
}

// FilterFinalizeBatch implements IRollup
func (m *MockRollup) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupFinalizeBatchIterator, error) {
	return nil, nil
}

// SetLastCommittedBatchIndex sets the mock value for LastCommittedBatchIndex
func (m *MockRollup) SetLastCommittedBatchIndex(index *big.Int) {
	m.lastCommittedBatchIndex = index
}

// SetLastFinalizedBatchIndex sets the mock value for LastFinalizedBatchIndex
func (m *MockRollup) SetLastFinalizedBatchIndex(index *big.Int) {
	m.lastFinalizedBatchIndex = index
}

// SetLegacyCutoverBatchIndex sets the immutable pre-upgrade batch boundary.
func (m *MockRollup) SetLegacyCutoverBatchIndex(index *big.Int) {
	m.legacyCutoverBatchIndex = index
}

// SetBatchInsideChallengeWindow sets the mock value for BatchInsideChallengeWindow
func (m *MockRollup) SetBatchInsideChallengeWindow(inside bool) {
	m.insideChallengeWindow = inside
}

// SetBatchExists sets the mock value for BatchExist
func (m *MockRollup) SetBatchExists(exists bool) {
	m.batchExists = exists
}

// SetStoredBlobHash controls the value returned by BatchBlobVersionedHashes.
func (m *MockRollup) SetStoredBlobHash(hash [32]byte) {
	m.storedBlobHash = hash
}

// SetFinalizeTx sets the mock value for FinalizeBatch transaction
func (m *MockRollup) SetFinalizeTx(tx *types.Transaction) {
	m.finalizeTx = tx
}
