package updater

import (
	"context"
	"github.com/morph-l2/go-ethereum/log"
	"sync"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"morph-l2/token-price-oracle/client"
)

// TxManager manages transaction sending to avoid nonce conflicts
type TxManager struct {
	l2Client *client.L2Client
	mu       sync.Mutex
}

// NewTxManager creates a new transaction manager
func NewTxManager(l2Client *client.L2Client) *TxManager {
	return &TxManager{
		l2Client: l2Client,
	}
}

// SendTransaction sends a transaction in a thread-safe manner
// It ensures only one transaction is sent at a time to avoid nonce conflicts
func (m *TxManager) SendTransaction(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Get transaction options (returns a copy)
	auth := m.l2Client.GetOpts()
	auth.Context = ctx
	// GasLimit is set to 0 to enable automatic gas estimation
	// The go-ethereum library will estimate gas if GasLimit is 0
	auth.GasLimit = 0

	// Execute transaction function
	tx, err := txFunc(auth)
	if err != nil {
		return nil, err
	}

	log.Info("Transaction sent", 
		"tx_hash", tx.Hash().Hex(),
		"gas_limit", tx.Gas())

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.l2Client.GetClient(), tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == 0 {
		log.Error("Transaction failed", 
			"tx_hash", tx.Hash().Hex(),
			"gas_used", receipt.GasUsed)
	} else {
		log.Info("Transaction confirmed", 
			"tx_hash", tx.Hash().Hex(), 
			"gas_used", receipt.GasUsed,
			"gas_limit", tx.Gas())
	}

	return receipt, nil
}
