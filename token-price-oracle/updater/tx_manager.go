package updater

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
	"morph-l2/token-price-oracle/client"
)

const (
	// GasPriceBumpPercent is the percentage to increase gas price for replacement tx
	// EIP-1559 requires at least 10% bump, we use 15% to be safe
	GasPriceBumpPercent = 15
	// MaxGasPriceBumpMultiplier limits how much we can bump gas price (e.g., 3x original)
	MaxGasPriceBumpMultiplier = 3
)

// PendingTxInfo stores information about a pending transaction
type PendingTxInfo struct {
	TxHash    common.Hash
	Nonce     uint64
	GasFeeCap *big.Int
	GasTipCap *big.Int
	SentAt    time.Time
}

// TxManager manages transaction sending to avoid nonce conflicts
type TxManager struct {
	l2Client  *client.L2Client
	mu        sync.Mutex
	pendingTx *PendingTxInfo // Track the last pending transaction
}

// NewTxManager creates a new transaction manager
func NewTxManager(l2Client *client.L2Client) *TxManager {
	return &TxManager{
		l2Client: l2Client,
	}
}

// SendTransaction sends a transaction in a thread-safe manner
// It ensures only one transaction is sent at a time to avoid nonce conflicts
// If there's a pending transaction, it will wait for it or replace it
func (m *TxManager) SendTransaction(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if there's a pending transaction that needs to be handled
	if m.pendingTx != nil {
		receipt, err := m.handlePendingTx(ctx)
		if err != nil {
			log.Warn("Failed to handle pending transaction, will try to replace",
				"pending_tx", m.pendingTx.TxHash.Hex(),
				"error", err)
			// Continue to send new transaction which may replace the pending one
		} else if receipt != nil {
			log.Info("Previous pending transaction confirmed",
				"tx_hash", m.pendingTx.TxHash.Hex(),
				"status", receipt.Status)
			m.pendingTx = nil
			// Previous tx confirmed, continue to send new transaction
		}
	}

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

	// Check if we need to replace a pending transaction (same nonce)
	if m.pendingTx != nil {
		// Get current nonce from network
		fromAddr := m.l2Client.WalletAddress()
		pendingNonce, err := m.l2Client.GetClient().PendingNonceAt(ctx, fromAddr)
		if err != nil {
			log.Warn("Failed to get pending nonce", "error", err)
		} else if pendingNonce <= m.pendingTx.Nonce {
			// There's still a pending tx with this nonce, need to replace it
			log.Info("Replacing pending transaction",
				"old_tx", m.pendingTx.TxHash.Hex(),
				"old_nonce", m.pendingTx.Nonce,
				"pending_nonce", pendingNonce)

			// Bump gas price for replacement
			auth.Nonce = big.NewInt(int64(m.pendingTx.Nonce))
			auth.GasFeeCap, auth.GasTipCap = m.bumpGasPrice(m.pendingTx.GasFeeCap, m.pendingTx.GasTipCap)

			log.Info("Gas price bumped for replacement",
				"old_fee_cap", m.pendingTx.GasFeeCap,
				"new_fee_cap", auth.GasFeeCap,
				"old_tip_cap", m.pendingTx.GasTipCap,
				"new_tip_cap", auth.GasTipCap)
		}
	}

	// Now send the actual transaction
	auth.NoSend = false
	tx, err = txFunc(auth)
	if err != nil {
		return nil, err
	}

	// Store pending transaction info
	m.pendingTx = &PendingTxInfo{
		TxHash:    tx.Hash(),
		Nonce:     tx.Nonce(),
		GasFeeCap: tx.GasFeeCap(),
		GasTipCap: tx.GasTipCap(),
		SentAt:    time.Now(),
	}

	log.Info("Transaction sent",
		"tx_hash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_limit", tx.Gas(),
		"gas_fee_cap", tx.GasFeeCap(),
		"gas_tip_cap", tx.GasTipCap())

	// Wait for transaction to be mined with custom timeout and retry logic
	receipt, err := m.waitForReceipt(ctx, tx.Hash(), 60*time.Second, 2*time.Second)
	if err != nil {
		log.Error("Failed to wait for transaction receipt",
			"tx_hash", tx.Hash().Hex(),
			"error", err)
		// Don't clear pendingTx here - let next round handle it
		return nil, err
	}

	// Transaction confirmed, clear pending tx
	m.pendingTx = nil
	return receipt, nil
}

// handlePendingTx checks if the pending transaction has been confirmed
func (m *TxManager) handlePendingTx(ctx context.Context) (*types.Receipt, error) {
	if m.pendingTx == nil {
		return nil, nil
	}

	// Try to get receipt for pending tx
	receipt, err := m.l2Client.GetClient().TransactionReceipt(ctx, m.pendingTx.TxHash)
	if err == nil && receipt != nil {
		return receipt, nil
	}

	// Check if the nonce has been used (tx might have been replaced)
	fromAddr := m.l2Client.WalletAddress()
	confirmedNonce, err := m.l2Client.GetClient().NonceAt(ctx, fromAddr, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get confirmed nonce: %w", err)
	}

	if confirmedNonce > m.pendingTx.Nonce {
		// Nonce has been used, the tx (or a replacement) was confirmed
		log.Info("Pending nonce has been confirmed (possibly replaced)",
			"pending_nonce", m.pendingTx.Nonce,
			"confirmed_nonce", confirmedNonce)
		m.pendingTx = nil
		return nil, nil
	}

	// Transaction still pending
	return nil, fmt.Errorf("transaction %s still pending (nonce: %d)", m.pendingTx.TxHash.Hex(), m.pendingTx.Nonce)
}

// bumpGasPrice increases gas price by GasPriceBumpPercent, capped at MaxGasPriceBumpMultiplier
func (m *TxManager) bumpGasPrice(oldFeeCap, oldTipCap *big.Int) (*big.Int, *big.Int) {
	// Calculate bump: oldPrice * (100 + GasPriceBumpPercent) / 100
	bumpMultiplier := big.NewInt(100 + GasPriceBumpPercent)
	hundred := big.NewInt(100)

	newFeeCap := new(big.Int).Mul(oldFeeCap, bumpMultiplier)
	newFeeCap.Div(newFeeCap, hundred)

	newTipCap := new(big.Int).Mul(oldTipCap, bumpMultiplier)
	newTipCap.Div(newTipCap, hundred)

	// Cap at MaxGasPriceBumpMultiplier times original
	maxFeeCap := new(big.Int).Mul(oldFeeCap, big.NewInt(MaxGasPriceBumpMultiplier))
	maxTipCap := new(big.Int).Mul(oldTipCap, big.NewInt(MaxGasPriceBumpMultiplier))

	if newFeeCap.Cmp(maxFeeCap) > 0 {
		log.Warn("Gas fee cap bump capped at max multiplier",
			"calculated", newFeeCap,
			"capped", maxFeeCap)
		newFeeCap = maxFeeCap
	}

	if newTipCap.Cmp(maxTipCap) > 0 {
		log.Warn("Gas tip cap bump capped at max multiplier",
			"calculated", newTipCap,
			"capped", maxTipCap)
		newTipCap = maxTipCap
	}

	return newFeeCap, newTipCap
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
