package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
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

		// Create opts with external signer address (for read-only operations)
		l2Client.opts = &bind.TransactOpts{
			From:   common.HexToAddress(cfg.ExternalSignAddress),
			NoSend: true, // We'll handle sending manually
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
	return &bind.TransactOpts{
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
