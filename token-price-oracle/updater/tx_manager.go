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
// Before sending, it checks if there are any pending transactions by comparing nonces
func (m *TxManager) SendTransaction(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.l2Client.IsExternalSign() {
		return m.sendWithExternalSign(ctx, txFunc)
	}
	return m.sendWithLocalSign(ctx, txFunc)
}

// sendWithLocalSign sends transaction using local private key signing
func (m *TxManager) sendWithLocalSign(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
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

	// Apply gas caps if configured (same logic as external sign)
	if err := m.applyGasCaps(ctx, auth); err != nil {
		return nil, fmt.Errorf("failed to apply gas caps: %w", err)
	}

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

	log.Info("Transaction sent (local sign)",
		"tx_hash", tx.Hash().Hex(),
		"nonce", tx.Nonce(),
		"gas_limit", tx.Gas())

	// Wait for transaction to be mined with timeout and retry logic
	receipt, err := m.waitForReceipt(ctx, tx.Hash(), 60*time.Second, 2*time.Second)
	if err != nil {
		log.Error("Failed to wait for transaction receipt",
			"tx_hash", tx.Hash().Hex(),
			"error", err)
		return nil, err
	}

	return receipt, nil
}

// sendWithExternalSign sends transaction using external signing service
func (m *TxManager) sendWithExternalSign(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	signer := m.l2Client.GetSigner()
	if signer == nil {
		return nil, fmt.Errorf("external signer is not initialized")
	}

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

	// Get transaction options (returns a copy) with NoSend=true to get calldata
	auth := m.l2Client.GetOpts()
	auth.Context = ctx
	auth.NoSend = true
	auth.GasLimit = 0

	// Call txFunc to get the transaction (this gives us the calldata and to address)
	tx, err := txFunc(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction: %w", err)
	}

	// Get the target contract address and calldata from the unsigned tx
	if tx.To() == nil {
		return nil, fmt.Errorf("contract creation transactions are not supported")
	}
	toAddr := *tx.To()
	callData := tx.Data()

	log.Info("Building external sign transaction",
		"to", toAddr.Hex(),
		"calldata_len", len(callData))

	// Create and sign transaction using external signer
	signedTx, err := signer.CreateAndSignTx(ctx, m.l2Client, toAddr, callData)
	if err != nil {
		return nil, fmt.Errorf("failed to create and sign transaction: %w", err)
	}

	// Send the signed transaction
	err = m.l2Client.GetClient().SendTransaction(ctx, signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send signed transaction: %w", err)
	}

	log.Info("Transaction sent (external sign)",
		"tx_hash", signedTx.Hash().Hex(),
		"gas_limit", signedTx.Gas())

	// Wait for transaction to be mined with custom timeout and retry logic
	receipt, err := m.waitForReceipt(ctx, signedTx.Hash(), 60*time.Second, 2*time.Second)
	if err != nil {
		log.Error("Failed to wait for transaction receipt",
			"tx_hash", signedTx.Hash().Hex(),
			"error", err)
		return nil, err
	}
	return receipt, nil
}

// applyGasCaps applies configured gas caps as upper limits to dynamic gas prices
// This ensures consistent behavior between local sign and external sign
func (m *TxManager) applyGasCaps(ctx context.Context, auth *bind.TransactOpts) error {
	caps, err := client.CalculateGasCaps(ctx, m.l2Client)
	if err != nil {
		return err
	}

	auth.GasTipCap = caps.TipCap
	auth.GasFeeCap = caps.FeeCap
	return nil
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
