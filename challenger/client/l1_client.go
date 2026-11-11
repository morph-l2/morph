package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	remotesigner "github.com/morph-l2/remote-signer-client/go/signer"

	"morph-l2/challenger/config"
)

// L1Client wraps the L1 RPC client and the signer that pays for challengeState.
type L1Client struct {
	client       *ethclient.Client
	chainID      *big.Int
	opts         *bind.TransactOpts
	signer       *Signer
	externalSign bool
	gasFeeCap    *big.Int
	gasTipCap    *big.Int
}

// NewL1Client dials L1 and prepares either a remote signer or a local key.
// Remote signing is implemented in this service. Contract bindings build an
// unsigned transaction, and the remote signer produces the signature that is broadcast.
// The flow was referenced from token-price-oracle, not imported from it.
func NewL1Client(cfg *config.Config) (*L1Client, error) {
	client, err := ethclient.Dial(cfg.L1RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L1 RPC: %w", err)
	}
	defer func() {
		if err != nil {
			client.Close()
		}
	}()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	l1Client := &L1Client{
		client:       client,
		chainID:      chainID,
		externalSign: cfg.ExternalSign,
	}
	if cfg.GasFeeCap != nil {
		l1Client.gasFeeCap = new(big.Int).SetUint64(*cfg.GasFeeCap)
		log.Info("Using gas fee cap limit", "maxGasFeeCap", *cfg.GasFeeCap)
	}
	if cfg.GasTipCap != nil {
		l1Client.gasTipCap = new(big.Int).SetUint64(*cfg.GasTipCap)
		log.Info("Using gas tip cap limit", "maxGasTipCap", *cfg.GasTipCap)
	}

	if cfg.ExternalSign {
		rsaPriv, parseErr := remotesigner.ParseRsaPrivateKey(cfg.ExternalSignRsaPriv)
		if parseErr != nil {
			err = fmt.Errorf("failed to parse RSA private key: %w", parseErr)
			return nil, err
		}
		l1Client.signer, err = NewSigner(
			cfg.ExternalSignAppid,
			rsaPriv,
			cfg.ExternalSignAddress,
			cfg.ExternalSignChain,
			cfg.ExternalSignUrl,
			chainID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create signer: %w", err)
		}

		fromAddr := common.HexToAddress(cfg.ExternalSignAddress)
		ethSigner := types.NewLondonSigner(chainID)
		// Placeholder signer lets the binding package build a transaction so we can
		// read To, calldata, and value. NoSend stays true; the placeholder is never broadcast.
		l1Client.opts = &bind.TransactOpts{
			From:   fromAddr,
			NoSend: true,
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return tx.WithSignature(ethSigner, make([]byte, 65))
			},
		}
		log.Info("L1 client initialized with external signing", "address", cfg.ExternalSignAddress, "chainID", chainID)
	} else {
		privateKeyHex := cfg.PrivateKey
		if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
			privateKeyHex = privateKeyHex[2:]
		}
		key, keyErr := crypto.HexToECDSA(privateKeyHex)
		if keyErr != nil {
			err = fmt.Errorf("failed to parse private key: %w", keyErr)
			return nil, err
		}
		opts, optsErr := bind.NewKeyedTransactorWithChainID(key, chainID)
		if optsErr != nil {
			err = fmt.Errorf("failed to create transactor: %w", optsErr)
			return nil, err
		}
		l1Client.opts = opts
		log.Info("L1 client initialized with local signing", "address", opts.From.Hex(), "chainID", chainID)
	}

	return l1Client, nil
}

// Close closes the underlying RPC client.
func (c *L1Client) Close() {
	c.client.Close()
}

// GetClient returns the underlying ethclient.
func (c *L1Client) GetClient() *ethclient.Client {
	return c.client
}

// GetOpts returns a copy of the transaction options.
func (c *L1Client) GetOpts() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:      c.opts.From,
		Nonce:     c.opts.Nonce,
		Signer:    c.opts.Signer,
		Value:     c.opts.Value,
		GasPrice:  c.opts.GasPrice,
		GasFeeCap: c.opts.GasFeeCap,
		GasTipCap: c.opts.GasTipCap,
		GasLimit:  c.opts.GasLimit,
		Context:   c.opts.Context,
		NoSend:    c.opts.NoSend,
	}
}

// GetBalance returns the account balance at the latest block.
func (c *L1Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return c.client.BalanceAt(ctx, address, nil)
}

// WalletAddress returns the address that signs challenge transactions.
func (c *L1Client) WalletAddress() common.Address {
	return c.opts.From
}

// IsExternalSign reports whether the remote signer is enabled.
func (c *L1Client) IsExternalSign() bool {
	return c.externalSign
}

// GetSigner returns the remote signer. It is nil when local signing is used.
func (c *L1Client) GetSigner() *Signer {
	return c.signer
}

// GetChainID returns the L1 chain ID.
func (c *L1Client) GetChainID() *big.Int {
	return c.chainID
}

// GetMaxGasFeeCap returns the configured fee cap, or nil when unset.
func (c *L1Client) GetMaxGasFeeCap() *big.Int {
	return c.gasFeeCap
}

// GetMaxGasTipCap returns the configured tip cap, or nil when unset.
func (c *L1Client) GetMaxGasTipCap() *big.Int {
	return c.gasTipCap
}
