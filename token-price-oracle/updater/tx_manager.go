package updater

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum"
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
// Before sending, it checks if there are any pending transactions by comparing nonces
func (m *TxManager) SendTransaction(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	fromAddr := m.l2Client.WalletAddress()

	// Check if there are pending transactions by comparing nonces
	confirmedNonce, err := m.l2Client.GetClient().NonceAt(ctx, fromAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get confirmed nonce: %w", err)
	}

	pendingNonce, err := m.l2Client.GetClient().PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get pending nonce: %w", err)
	}

	if pendingNonce > confirmedNonce {
		// There are pending transactions, don't send new one
		log.Warn("Found pending transactions, skipping this round",
			"address", fromAddr.Hex(),
			"confirmed_nonce", confirmedNonce,
			"pending_nonce", pendingNonce,
			"pending_count", pendingNonce-confirmedNonce)
		return nil, fmt.Errorf("pending transactions exist (confirmed: %d, pending: %d)", confirmedNonce, pendingNonce)
	}

	log.Info("No pending transactions, proceeding to send",
		"address", fromAddr.Hex(),
		"nonce", confirmedNonce)

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
		"nonce", tx.Nonce(),
		"gas_limit", tx.Gas())

	// Wait for transaction to be mined - will keep waiting until confirmed or dropped
	receipt, err := m.waitForReceipt(ctx, tx.Hash(), 2*time.Second)
	if err != nil {
		log.Error("Failed to wait for transaction receipt",
			"tx_hash", tx.Hash().Hex(),
			"error", err)
		return nil, err
	}

	return receipt, nil
}

// waitForReceipt waits for a transaction receipt indefinitely until:
// 1. Receipt is received (transaction confirmed)
// 2. Transaction is not found (dropped from pool) - exits immediately
// Network errors will cause retry, NOT exit
func (m *TxManager) waitForReceipt(ctx context.Context, txHash common.Hash, pollInterval time.Duration) (*types.Receipt, error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	startTime := time.Now()

	log.Info("Waiting for transaction receipt (will wait indefinitely)",
		"tx_hash", txHash.Hex(),
		"poll_interval", pollInterval)

	for {
		// Check context cancellation first
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while waiting for transaction %s (waited %v): %w",
				txHash.Hex(), time.Since(startTime), ctx.Err())
		default:
		}

		// Try to get the receipt first
		receipt, err := m.l2Client.GetClient().TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil {
			log.Info("Receipt received",
				"tx_hash", txHash.Hex(),
				"status", receipt.Status,
				"gas_used", receipt.GasUsed,
				"block_number", receipt.BlockNumber,
				"waited", time.Since(startTime))
			return receipt, nil
		}

		// No receipt yet, check if transaction is still in the pool
		tx, isPending, err := m.l2Client.GetClient().TransactionByHash(ctx, txHash)

		if err != nil {
			// Check if it's a "not found" error - transaction dropped
			if errors.Is(err, ethereum.NotFound) {
				log.Error("Transaction not found, dropped from pool",
					"tx_hash", txHash.Hex(),
					"waited", time.Since(startTime))
				return nil, fmt.Errorf("transaction %s dropped from pool (not found)", txHash.Hex())
			}

			// Other errors (network, etc.) - just log and retry
			log.Warn("Transaction query failed, will retry",
				"tx_hash", txHash.Hex(),
				"waited", time.Since(startTime),
				"error", err)
		} else if tx == nil {
			// tx is nil but no error - treat as not found
			log.Error("Transaction returned nil, dropped from pool",
				"tx_hash", txHash.Hex(),
				"waited", time.Since(startTime))
			return nil, fmt.Errorf("transaction %s dropped from pool (returned nil)", txHash.Hex())
		} else {
			// Transaction found, log progress every minute
			elapsed := time.Since(startTime)
			if int(elapsed.Seconds()) > 0 && int(elapsed.Seconds())%60 == 0 {
				log.Info("Still waiting for transaction receipt",
					"tx_hash", txHash.Hex(),
					"is_pending", isPending,
					"waited", elapsed)
			}
		}

		// Wait for next poll
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context cancelled while waiting for transaction %s (waited %v): %w",
				txHash.Hex(), time.Since(startTime), ctx.Err())
		case <-ticker.C:
			// Continue to next iteration
		}
	}
}
