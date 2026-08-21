package mock

import (
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
)

// MockSubmitter is a controllable implementation of iface.ISubmitter.
type MockSubmitter struct {
	active map[common.Address]bool
	err    error
}

func NewMockSubmitter() *MockSubmitter {
	return &MockSubmitter{active: make(map[common.Address]bool)}
}

func (m *MockSubmitter) IsActive(_ *bind.CallOpts, addr common.Address) (bool, error) {
	if m.err != nil {
		return false, m.err
	}
	return m.active[addr], nil
}

func (m *MockSubmitter) SetActive(addr common.Address, active bool) {
	if m.active == nil {
		m.active = make(map[common.Address]bool)
	}
	m.active[addr] = active
}

func (m *MockSubmitter) SetError(err error) {
	m.err = err
}
