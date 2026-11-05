package services

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"fmt"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"

	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/localpool"
)

// ExternalSignConfig carries the app-level credentials of the external sign
// service. They are shared by every account; only the signing address differs.
type ExternalSignConfig struct {
	Appid   string
	RsaPriv *rsa.PrivateKey
	Chain   string
	Url     string
}

// Account bundles everything that is private to a single L1 sender: signing
// material, nonce sequence (pending pool) and journal. Global state such as
// leveldb, batch cache and contract bindings stays on Rollup.
type Account struct {
	name       string
	privKey    *ecdsa.PrivateKey // local signing mode
	extAddress common.Address    // external signing mode: address to sign for
	useExtSign bool
	extSigner  *externalsign.ExternalSign
	extURL     string
	signer     ethtypes.Signer
	pendingTxs *PendingTxs
	journal    *localpool.Journal
}

// NewLocalAccount creates an account that signs with a local private key.
func NewLocalAccount(name string, privKey *ecdsa.PrivateKey, signer ethtypes.Signer, journal *localpool.Journal) *Account {
	return &Account{name: name, privKey: privKey, signer: signer, journal: journal}
}

// NewExternalAccount creates an account that signs through the external sign
// service on behalf of addr, reusing the shared app credentials.
func NewExternalAccount(name string, addr common.Address, cfg ExternalSignConfig, signer ethtypes.Signer, journal *localpool.Journal) *Account {
	return &Account{
		name:       name,
		extAddress: addr,
		useExtSign: true,
		extSigner:  externalsign.NewExternalSign(cfg.Appid, cfg.RsaPriv, addr.Hex(), cfg.Chain, signer),
		extURL:     cfg.Url,
		signer:     signer,
		journal:    journal,
	}
}

// WalletAddr returns the account address.
func (a *Account) WalletAddr() common.Address {
	if a.useExtSign {
		return a.extAddress
	}
	return crypto.PubkeyToAddress(a.privKey.PublicKey)
}

// Sign signs tx with this account's key (or external signing address).
func (a *Account) Sign(tx *ethtypes.Transaction) (*ethtypes.Transaction, error) {
	if a.useExtSign {
		signedTx, err := a.extSigner.RequestSign(a.extURL, tx)
		if err != nil {
			return nil, fmt.Errorf("externalsign sign tx error:%v", err)
		}
		return signedTx, nil
	}
	signedTx, err := ethtypes.SignTx(tx, a.signer, a.privKey)
	if err != nil {
		return nil, fmt.Errorf("sign tx error:%v", err)
	}
	return signedTx, nil
}

// getNextNonce returns the next nonce of this account, taking its own pending
// pool into account.
func (a *Account) getNextNonce(l1 iface.Client) (uint64, error) {
	nonce, err := l1.PendingNonceAt(context.Background(), a.WalletAddr())
	if err != nil {
		return 0, fmt.Errorf("failed to get pending nonce: %w", err)
	}
	if a.pendingTxs.Len() == 0 {
		return nonce, nil
	}
	pn := a.pendingTxs.GetNonce()
	if pn+1 > nonce {
		return pn + 1, nil
	}
	return nonce, nil
}

// initPending opens the journal and recovers this account's pending txs.
func (a *Account) initPending(rollupAbi *abi.ABI) error {
	if err := a.journal.Init(); err != nil {
		return fmt.Errorf("journal file init failed: %w", err)
	}
	a.pendingTxs = NewPendingTxs(a.journal)
	txs, err := a.journal.ParseAllTxs()
	if err != nil {
		return fmt.Errorf("parse journal error: %w", err)
	}
	if err := a.pendingTxs.Recover(txs, rollupAbi); err != nil {
		return fmt.Errorf("failed to recover pending transactions: %w", err)
	}
	return nil
}
