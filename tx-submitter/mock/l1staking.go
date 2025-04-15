package mock

import (
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
)

// MockL1Staking is a mock implementation of IL1Staking
type MockL1Staking struct {
	activeStakers []common.Address
}

// NewMockL1Staking creates a new mock L1Staking instance
func NewMockL1Staking() *MockL1Staking {
	return &MockL1Staking{
		activeStakers: []common.Address{},
	}
}

// IsStaker implements IL1Staking
func (m *MockL1Staking) IsStaker(opts *bind.CallOpts, addr common.Address) (bool, error) {
	for _, staker := range m.activeStakers {
		if staker == addr {
			return true, nil
		}
	}
	return false, nil
}

// GetStakersBitmap implements IL1Staking
func (m *MockL1Staking) GetStakersBitmap(opts *bind.CallOpts, _stakers []common.Address) (*big.Int, error) {
	bitmap := big.NewInt(0)
	for i, staker := range _stakers {
		for _, activeStaker := range m.activeStakers {
			if staker == activeStaker {
				bitmap.SetBit(bitmap, i, 1)
			}
		}
	}
	return bitmap, nil
}

// GetActiveStakers implements IL1Staking
func (m *MockL1Staking) GetActiveStakers(opts *bind.CallOpts) ([]common.Address, error) {
	return m.activeStakers, nil
}

// GetStakers implements IL1Staking
func (m *MockL1Staking) GetStakers(opts *bind.CallOpts) ([255]common.Address, error) {
	var result [255]common.Address
	for i := 0; i < len(m.activeStakers) && i < 255; i++ {
		result[i] = m.activeStakers[i]
	}
	return result, nil
}

// SetActiveStakers sets the active stakers for testing
func (m *MockL1Staking) SetActiveStakers(stakers []common.Address) {
	m.activeStakers = stakers
}
