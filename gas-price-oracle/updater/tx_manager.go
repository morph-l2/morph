package updater

import (
	"context"
	"sync"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/morph/gas-price-oracle/client"
	"github.com/sirupsen/logrus"
)

// TxManager manages transaction sending to avoid nonce conflicts
type TxManager struct {
	l2Client *client.L2Client
	mu       sync.Mutex
	log      *logrus.Entry
}

// NewTxManager creates a new transaction manager
func NewTxManager(l2Client *client.L2Client) *TxManager {
	return &TxManager{
		l2Client: l2Client,
		log:      logrus.WithField("component", "tx_manager"),
	}
}

// SendTransaction sends a transaction in a thread-safe manner
// It ensures only one transaction is sent at a time to avoid nonce conflicts
func (m *TxManager) SendTransaction(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Get transaction options
	auth := m.l2Client.GetAuth()
	auth.Context = ctx
	auth.GasLimit = 100000 // Default gas limit

	// Execute transaction function
	tx, err := txFunc(auth)
	if err != nil {
		return nil, err
	}

	m.log.WithField("tx_hash", tx.Hash().Hex()).Debug("Transaction sent")

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(ctx, m.l2Client.GetClient(), tx)
	if err != nil {
		return nil, err
	}

	if receipt.Status == 0 {
		m.log.WithField("tx_hash", tx.Hash().Hex()).Error("Transaction failed")
	} else {
		m.log.WithFields(logrus.Fields{
			"tx_hash":  tx.Hash().Hex(),
			"gas_used": receipt.GasUsed,
		}).Debug("Transaction confirmed")
	}

	return receipt, nil
}
