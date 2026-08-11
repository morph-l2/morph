package mock

import (
	"sync"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
)

// MockSubmitter implements the minimal submitter qualification view used by
// the transaction submitter.
type MockSubmitter struct {
	mu        sync.RWMutex
	active    map[common.Address]bool
	defaultOn bool
	lastQuery common.Address
	err       error
}

func NewMockSubmitter() *MockSubmitter {
	return &MockSubmitter{
		active:    make(map[common.Address]bool),
		defaultOn: true,
	}
}

func (m *MockSubmitter) IsActive(_ *bind.CallOpts, addr common.Address) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.lastQuery = addr
	if m.err != nil {
		return false, m.err
	}
	active, configured := m.active[addr]
	if !configured {
		return m.defaultOn, nil
	}
	return active, nil
}

func (m *MockSubmitter) SetActive(addr common.Address, active bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.active[addr] = active
}

func (m *MockSubmitter) SetDefaultActive(active bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.defaultOn = active
}

func (m *MockSubmitter) SetError(err error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.err = err
}

func (m *MockSubmitter) LastQuery() common.Address {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.lastQuery
}
