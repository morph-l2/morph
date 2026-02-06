package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	"morph-l2/token-price-oracle/config"
)

// L2Client wraps L2 chain client
type L2Client struct {
	client       *ethclient.Client
	chainID      *big.Int
	opts         *bind.TransactOpts
	signer       *Signer
	externalSign bool
	gasFeeCap    *big.Int // Fixed gas fee cap (nil means use dynamic)
	gasTipCap    *big.Int // Fixed gas tip cap (nil means use dynamic)
}

// NewL2Client creates new L2 client
func NewL2Client(rpcURL string, cfg *config.Config) (*L2Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to dial L2 RPC: %w", err)
	}

	// Ensure client is closed if any subsequent step fails
	defer func() {
		if err != nil {
			client.Close()
		}
	}()

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	l2Client := &L2Client{
		client:       client,
		chainID:      chainID,
		externalSign: cfg.ExternalSign,
	}

	// Set gas fee caps if configured (used as max cap, not fixed value)
	if cfg.GasFeeCap != nil {
		l2Client.gasFeeCap = new(big.Int).SetUint64(*cfg.GasFeeCap)
		log.Info("Using gas fee cap limit", "maxGasFeeCap", *cfg.GasFeeCap)
	}
	if cfg.GasTipCap != nil {
		l2Client.gasTipCap = new(big.Int).SetUint64(*cfg.GasTipCap)
		log.Info("Using gas tip cap limit", "maxGasTipCap", *cfg.GasTipCap)
	}

	if cfg.ExternalSign {
		// External sign mode
		rsaPriv, err := externalsign.ParseRsaPrivateKey(cfg.ExternalSignRsaPriv)
		if err != nil {
			return nil, fmt.Errorf("failed to parse RSA private key: %w", err)
		}

		l2Client.signer = NewSigner(
			true,
			cfg.ExternalSignAppid,
			rsaPriv,
			cfg.ExternalSignAddress,
			cfg.ExternalSignChain,
			cfg.ExternalSignUrl,
			chainID,
		)

		fromAddr := common.HexToAddress(cfg.ExternalSignAddress)
		ethSigner := types.NewLondonSigner(chainID)

		// Create opts with a placeholder signer for building transactions.
		// This allows contract bindings to construct transaction objects so we can
		// extract the calldata and target address. The placeholder signature is never
		// actually broadcast - the real signing happens via external signer.
		// SAFETY: NoSend is always true, and tx_manager.go only extracts To() and Data()
		// from the placeholder tx, then creates a new properly signed transaction.
		l2Client.opts = &bind.TransactOpts{
			From:   fromAddr,
			NoSend: true, // CRITICAL: Must always be true for external signing mode
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				// Placeholder signer - returns tx with dummy signature to satisfy bind package.
				// This tx is NEVER sent; only used to extract calldata for external signing.
				return tx.WithSignature(ethSigner, make([]byte, 65))
			},
		}

		log.Info("L2 client initialized with external signing",
			"address", cfg.ExternalSignAddress,
			"chainID", chainID)
	} else {
		// Local private key mode
		privateKeyHex := cfg.PrivateKey
		if len(cfg.PrivateKey) > 2 && cfg.PrivateKey[:2] == "0x" {
			privateKeyHex = cfg.PrivateKey[2:]
		}
		key, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			return nil, fmt.Errorf("failed to parse private key: %w", err)
		}

		// Create transaction options
		opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
		if err != nil {
			return nil, fmt.Errorf("failed to create transactor: %w", err)
		}
		l2Client.opts = opts

		log.Info("L2 client initialized with local signing",
			"address", opts.From.Hex(),
			"chainID", chainID)
	}

	return l2Client, nil
}

// Close closes client connection
func (c *L2Client) Close() {
	c.client.Close()
}

// GetClient returns the underlying ethclient
func (c *L2Client) GetClient() *ethclient.Client {
	return c.client
}

// GetOpts returns a copy of transaction options
// Returns a new instance to prevent concurrent modification
func (c *L2Client) GetOpts() *bind.TransactOpts {
	// Return a copy to prevent shared state issues
	opts := &bind.TransactOpts{
		From:     c.opts.From,
		Nonce:    c.opts.Nonce,
		Signer:   c.opts.Signer,
		Value:    c.opts.Value,
		GasPrice: c.opts.GasPrice,
		GasFeeCap: c.opts.GasFeeCap,
		GasTipCap: c.opts.GasTipCap,
		GasLimit:  c.opts.GasLimit,
		Context:   c.opts.Context,
		NoSend:    c.opts.NoSend,
	}
	// Override with fixed gas fee if configured
	if c.gasFeeCap != nil {
		opts.GasFeeCap = c.gasFeeCap
	}
	if c.gasTipCap != nil {
		opts.GasTipCap = c.gasTipCap
	}
	return opts
}

// GetBalance returns account balance
func (c *L2Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return c.client.BalanceAt(ctx, address, nil)
}

// WalletAddress returns wallet address
func (c *L2Client) WalletAddress() common.Address {
	return c.opts.From
}

// IsExternalSign returns whether external signing is enabled
func (c *L2Client) IsExternalSign() bool {
	return c.externalSign
}

// GetSigner returns the external signer (nil if using local signing)
func (c *L2Client) GetSigner() *Signer {
	return c.signer
}

// GetChainID returns the chain ID
func (c *L2Client) GetChainID() *big.Int {
	return c.chainID
}

// GetFixedGasFeeCap returns the fixed gas fee cap (nil if not configured)
func (c *L2Client) GetFixedGasFeeCap() *big.Int {
	return c.gasFeeCap
}

// GetFixedGasTipCap returns the fixed gas tip cap (nil if not configured)
func (c *L2Client) GetFixedGasTipCap() *big.Int {
	return c.gasTipCap
}
