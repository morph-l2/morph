package mock

import (
	"math/big"

	"github.com/morph-l2/go-ethereum/core/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"

	"morph-l2/bindings/bindings"
)

// MockRollup implements the IRollup interface for testing
type MockRollup struct {
	lastCommittedBatchIndex *big.Int
	lastFinalizedBatchIndex *big.Int
	insideChallengeWindow   bool
	batchExists             bool
	batchTx                 *types.Transaction
	finalizeTx              *types.Transaction
}

// NewMockRollup creates a new instance of MockRollup
func NewMockRollup() *MockRollup {
	return &MockRollup{
		lastCommittedBatchIndex: big.NewInt(0),
		lastFinalizedBatchIndex: big.NewInt(0),
		insideChallengeWindow:   false,
		batchExists:             false,
	}
}

// LastCommittedBatchIndex implements IRollup
func (m *MockRollup) LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	return m.lastCommittedBatchIndex, nil
}

// CommitBatch implements IRollup
func (m *MockRollup) CommitBatch(opts *bind.TransactOpts, batchDataInput bindings.IRollupBatchDataInput, batchSignatureInput bindings.IRollupBatchSignatureInput) (*types.Transaction, error) {
	return m.batchTx, nil
}

// LastFinalizedBatchIndex implements IRollup
func (m *MockRollup) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	return m.lastFinalizedBatchIndex, nil
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

// SetLastCommittedBatchIndex sets the mock value for LastCommittedBatchIndex
func (m *MockRollup) SetLastCommittedBatchIndex(index *big.Int) {
	m.lastCommittedBatchIndex = index
}

// SetLastFinalizedBatchIndex sets the mock value for LastFinalizedBatchIndex
func (m *MockRollup) SetLastFinalizedBatchIndex(index *big.Int) {
	m.lastFinalizedBatchIndex = index
}

// SetBatchInsideChallengeWindow sets the mock value for BatchInsideChallengeWindow
func (m *MockRollup) SetBatchInsideChallengeWindow(inside bool) {
	m.insideChallengeWindow = inside
}

// SetBatchExists sets the mock value for BatchExist
func (m *MockRollup) SetBatchExists(exists bool) {
	m.batchExists = exists
}

// SetFinalizeTx sets the mock value for FinalizeBatch transaction
func (m *MockRollup) SetFinalizeTx(tx *types.Transaction) {
	m.finalizeTx = tx
}

// SetBatchTx sets the mock value for CommitBatch transaction
func (m *MockRollup) SetBatchTx(tx *types.Transaction) {
	m.batchTx = tx
}
