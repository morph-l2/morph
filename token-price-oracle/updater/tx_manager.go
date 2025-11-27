package updater

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
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
	
	// First, estimate gas with GasLimit = 0
	auth.GasLimit = 0
	auth.NoSend = true
	tx, err := txFunc(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}
	
	// Use 1.5x of estimated gas as the actual gas limit
	estimatedGas := tx.Gas()
	auth.GasLimit = estimatedGas * 3 / 2
	log.Info("Gas estimation completed", "estimated", estimatedGas, "actual_limit", auth.GasLimit)
	
	// Now send the actual transaction
	auth.NoSend = false
	tx, err = txFunc(auth)
	if err != nil {
		return nil, err
	}

	log.Info("Transaction sent",
		"tx_hash", tx.Hash().Hex(),
		"gas_limit", tx.Gas())

	// Wait for transaction to be mined with custom timeout and retry logic
	receipt, err := m.waitForReceipt(ctx, tx.Hash(), 60*time.Second, 2*time.Second)
	if err != nil {
		log.Error("Failed to wait for transaction receipt",
			"tx_hash", tx.Hash().Hex(),
			"error", err)
		return nil, err
	}
	return receipt, nil
}

// waitForReceipt waits for a transaction receipt with timeout and custom polling interval
func (m *TxManager) waitForReceipt(ctx context.Context, txHash common.Hash, timeout, pollInterval time.Duration) (*types.Receipt, error) {
	deadline := time.Now().Add(timeout)
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	log.Debug("Waiting for transaction receipt",
		"tx_hash", txHash.Hex(),
		"timeout", timeout,
		"poll_interval", pollInterval)

	for {
		// Check if we've exceeded the timeout
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("timeout waiting for transaction %s after %v", txHash.Hex(), timeout)
		}

		// Try to get the receipt
		receipt, err := m.l2Client.GetClient().TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil {
			log.Debug("Receipt received",
				"tx_hash", txHash.Hex(),
				"status", receipt.Status,
				"gas_used", receipt.GasUsed,
				"block_number", receipt.BlockNumber)
			return receipt, nil
		}

		if err != nil {
			log.Trace("Receipt retrieval failed, will retry",
				"tx_hash", txHash.Hex(),
				"error", err)
		}

		// Wait for next poll or context cancellation
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while waiting for transaction %s: %w", txHash.Hex(), ctx.Err())
		case <-ticker.C:
			// Continue to next iteration
		}
	}
}
