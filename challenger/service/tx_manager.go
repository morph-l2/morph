package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/challenger/client"
)

// TxManager serializes challenge submissions so two requests cannot spend the same nonce.
type TxManager struct {
	l1             *client.L1Client
	receiptTimeout time.Duration
	mu             sync.Mutex
}

// NewTxManager creates a transaction manager.
func NewTxManager(l1 *client.L1Client, receiptTimeout time.Duration) *TxManager {
	return &TxManager{l1: l1, receiptTimeout: receiptTimeout}
}

// SendTransaction builds, signs, and broadcasts one transaction.
// txFunc must set opts.Value before invoking the contract binding.
// methodSig is the canonical ABI signature passed to the remote signer.
func (m *TxManager) SendTransaction(
	ctx context.Context,
	txFunc func(*bind.TransactOpts) (*types.Transaction, error),
	methodSig string,
) (*types.Receipt, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.l1.IsExternalSign() {
		return m.sendWithExternalSign(ctx, txFunc, methodSig)
	}
	return m.sendWithLocalSign(ctx, txFunc)
}

func (m *TxManager) rejectIfPending(ctx context.Context, from common.Address) error {
	confirmed, err := m.l1.GetClient().NonceAt(ctx, from, nil)
	if err != nil {
		return fmt.Errorf("failed to get confirmed nonce: %w", err)
	}
	pending, err := m.l1.GetClient().PendingNonceAt(ctx, from)
	if err != nil {
		return fmt.Errorf("failed to get pending nonce: %w", err)
	}
	if pending > confirmed {
		log.Warn("Found pending transactions, skipping challenge",
			"address", from.Hex(),
			"confirmed_nonce", confirmed,
			"pending_nonce", pending)
		return &BusyError{Confirmed: confirmed, Pending: pending}
	}
	return nil
}

func (m *TxManager) applyGasCaps(ctx context.Context, auth *bind.TransactOpts) error {
	caps, err := client.CalculateGasCaps(ctx, m.l1.GetClient(), m.l1.GetMaxGasTipCap(), m.l1.GetMaxGasFeeCap())
	if err != nil {
		return err
	}
	auth.GasTipCap = caps.TipCap
	auth.GasFeeCap = caps.FeeCap
	return nil
}

func (m *TxManager) sendWithLocalSign(ctx context.Context, txFunc func(*bind.TransactOpts) (*types.Transaction, error)) (*types.Receipt, error) {
	from := m.l1.WalletAddress()
	if err := m.rejectIfPending(ctx, from); err != nil {
		return nil, err
	}

	auth := m.l1.GetOpts()
	auth.Context = ctx
	if err := m.applyGasCaps(ctx, auth); err != nil {
		return nil, fmt.Errorf("failed to apply gas caps: %w", err)
	}

	auth.GasLimit = 0
	auth.NoSend = true
	tx, err := txFunc(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas: %w", err)
	}
	auth.GasLimit = tx.Gas() * 3 / 2
	auth.NoSend = false
	tx, err = txFunc(auth)
	if err != nil {
		return nil, err
	}
	log.Info("Transaction sent (local sign)", "tx_hash", tx.Hash().Hex(), "nonce", tx.Nonce(), "value", tx.Value())
	return m.waitForReceipt(ctx, tx.Hash())
}

func (m *TxManager) sendWithExternalSign(
	ctx context.Context,
	txFunc func(*bind.TransactOpts) (*types.Transaction, error),
	methodSig string,
) (*types.Receipt, error) {
	signer := m.l1.GetSigner()
	if signer == nil {
		return nil, fmt.Errorf("external signer is not initialized")
	}
	from := m.l1.WalletAddress()
	if err := m.rejectIfPending(ctx, from); err != nil {
		return nil, err
	}

	auth := m.l1.GetOpts()
	auth.Context = ctx
	auth.NoSend = true
	auth.GasLimit = 0
	tx, err := txFunc(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction: %w", err)
	}
	if tx.To() == nil {
		return nil, fmt.Errorf("contract creation transactions are not supported")
	}

	signedTx, err := signer.CreateAndSignTx(
		ctx,
		m.l1.GetClient(),
		*tx.To(),
		tx.Value(),
		tx.Data(),
		methodSig,
		m.l1.GetMaxGasTipCap(),
		m.l1.GetMaxGasFeeCap(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create and sign transaction: %w", err)
	}
	if err := m.l1.GetClient().SendTransaction(ctx, signedTx); err != nil {
		return nil, fmt.Errorf("failed to send signed transaction %s: %w", signedTx.Hash().Hex(), err)
	}
	log.Info("Transaction sent (external sign)", "tx_hash", signedTx.Hash().Hex(), "value", signedTx.Value())
	return m.waitForReceipt(ctx, signedTx.Hash())
}

func (m *TxManager) waitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	deadline := time.Now().Add(m.receiptTimeout)
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		if time.Now().After(deadline) {
			return nil, &TxError{
				TxHash:  txHash.Hex(),
				Pending: true,
				Err:     fmt.Errorf("timeout waiting for transaction %s after %s", txHash.Hex(), m.receiptTimeout),
			}
		}
		receipt, err := m.l1.GetClient().TransactionReceipt(ctx, txHash)
		if err == nil && receipt != nil {
			if receipt.Status == types.ReceiptStatusFailed {
				return receipt, &TxError{
					TxHash:   txHash.Hex(),
					Reverted: true,
					Err:      fmt.Errorf("transaction %s reverted", txHash.Hex()),
				}
			}
			return receipt, nil
		}
		select {
		case <-ctx.Done():
			return nil, &TxError{
				TxHash:  txHash.Hex(),
				Pending: true,
				Err:     fmt.Errorf("context cancelled while waiting for transaction %s: %w", txHash.Hex(), ctx.Err()),
			}
		case <-ticker.C:
		}
	}
}
