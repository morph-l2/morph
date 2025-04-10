package mock

import (
	"github.com/morph-l2/go-ethereum/core/types"
)

// MockJournal is a mock implementation of the journal interface for testing
type MockJournal struct {
	txs []*types.Transaction
}

func NewMockJournal() *MockJournal {
	return &MockJournal{
		txs: make([]*types.Transaction, 0),
	}
}

func (j *MockJournal) Init() error {
	return nil
}

func (j *MockJournal) AppendTx(tx *types.Transaction) error {
	j.txs = append(j.txs, tx)
	return nil
}

func (j *MockJournal) ParseAllTxs() ([]*types.Transaction, error) {
	return j.txs, nil
}

func (j *MockJournal) Clean() error {
	j.txs = make([]*types.Transaction, 0)
	return nil
}
