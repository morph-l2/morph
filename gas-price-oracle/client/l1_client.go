package client

import (
	"context"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
)

// L1Client wraps L1 chain client
type L1Client struct {
	client *ethclient.Client
}

// NewL1Client creates L1 client
func NewL1Client(rpcURL string) (*L1Client, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &L1Client{client: client}, nil
}

// Close closes client connection
func (c *L1Client) Close() {
	c.client.Close()
}

// GetClient returns the underlying ethclient
func (c *L1Client) GetClient() *ethclient.Client {
	return c.client
}

// GetLatestBlock gets the latest block
func (c *L1Client) GetLatestBlock(ctx context.Context) (*types.Block, error) {
	header, err := c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return c.client.BlockByNumber(ctx, header.Number)
}

// GetBlockByNumber gets block by number
func (c *L1Client) GetBlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return c.client.BlockByNumber(ctx, number)
}

// GetTransaction gets transaction details
func (c *L1Client) GetTransaction(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return c.client.TransactionByHash(ctx, hash)
}

// GetTransactionReceipt gets transaction receipt
func (c *L1Client) GetTransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	return c.client.TransactionReceipt(ctx, hash)
}

// GetBlockNumber gets current block number
func (c *L1Client) GetBlockNumber(ctx context.Context) (uint64, error) {
	return c.client.BlockNumber(ctx)
}

// GetBaseFee gets L1 base fee and blob base fee
func (c *L1Client) GetBaseFee(ctx context.Context) (*big.Int, *big.Int, error) {
	block, err := c.GetLatestBlock(ctx)
	if err != nil {
		return nil, nil, err
	}

	baseFee := block.BaseFee()
	if baseFee == nil {
		baseFee = big.NewInt(0)
	}

	// Calculate blob base fee
	// Try to get ExcessBlobGas from block header if available
	var excessBlobGas uint64
	if header := block.Header(); header.ExcessBlobGas != nil {
		excessBlobGas = *header.ExcessBlobGas
	}
	blobBaseFee := calcBlobBaseFee(excessBlobGas)

	return baseFee, blobBaseFee, nil
}

// GetGasPrice gets L1 gas price
func (c *L1Client) GetGasPrice(ctx context.Context) (*big.Int, error) {
	return c.client.SuggestGasPrice(ctx)
}

// calcBlobBaseFee calculates blob base fee (EIP-4844)
func calcBlobBaseFee(excessBlobGas uint64) *big.Int {
	return fakeExponential(
		big.NewInt(minBlobGasPrice),
		big.NewInt(int64(excessBlobGas)),
		big.NewInt(blobGasPriceUpdateFraction),
	)
}

const (
	minBlobGasPrice            = 1
	blobGasPriceUpdateFraction = 3338477
)

// fakeExponential approximates exponential function for blob gas price calculation
// Implements the fake_exponential algorithm from EIP-4844
func fakeExponential(factor, numerator, denominator *big.Int) *big.Int {
	i := big.NewInt(1)
	output := big.NewInt(0)
	numeratorAccum := new(big.Int).Mul(factor, denominator)

	for numeratorAccum.Sign() > 0 {
		output.Add(output, numeratorAccum)

		// numerator_accum = (numerator_accum * numerator) / (denominator * i)
		numeratorAccum.Mul(numeratorAccum, numerator)
		numeratorAccum.Div(numeratorAccum, new(big.Int).Mul(denominator, i))

		i.Add(i, big.NewInt(1))
	}

	return output.Div(output, denominator)
}
