package client

import (
	"context"
	"errors"
	"math/big"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
)

// L2Client wraps L2 chain client
type L2Client struct {
	client  *ethclient.Client
	chainID *big.Int
	opts    *bind.TransactOpts
}

// NewL2Client creates new L2 client
func NewL2Client(rpcURL string, privateKey string) (*L2Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	// Get chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	// Parse private key (remove 0x prefix if present)
	privateKeyHex := privateKey
	if len(privateKey) > 2 && privateKey[:2] == "0x" {
		privateKeyHex = privateKey[2:]
	}
	key, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.New("invalid hex character 'x' in private key")
	}

	// Create transaction options
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}

	return &L2Client{
		client:  client,
		chainID: chainID,
		opts:    opts,
	}, nil
}

// Close closes client connection
func (c *L2Client) Close() {
	c.client.Close()
}

// GetClient returns the underlying ethclient
func (c *L2Client) GetClient() *ethclient.Client {
	return c.client
}

// GetOpts returns transaction options
func (c *L2Client) GetOpts() *bind.TransactOpts {
	return c.opts
}

// GetBalance returns account balance
func (c *L2Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	return c.client.BalanceAt(ctx, address, nil)
}

// WalletAddress returns wallet address
func (c *L2Client) WalletAddress() common.Address {
	return c.opts.From
}
