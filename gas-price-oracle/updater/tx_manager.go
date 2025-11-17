package updater

import (
	"context"
	"github.com/morph-l2/go-ethereum/log"
	"sync"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"morph-l2/gas-price-oracle/client"
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

	// Get transaction options
	auth := m.l2Client.GetOpts()
	auth.Context = ctx
	auth.GasLimit = 100000 // Default gas limit

	// Execute transaction function
	tx, err := txFunc(auth)
	if err != nil {
		return nil, err
	}

	log.Info("Transaction sent", "tx_hash", tx.Hash().Hex())

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.l2Client.GetClient(), tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == 0 {
		log.Error("Transaction failed", "tx_hash", tx.Hash().Hex())
	} else {
		log.Debug("Transaction confirmed", "tx_hash", tx.Hash().Hex(), "gas_used", receipt.GasUsed)
	}

	return receipt, nil
}
