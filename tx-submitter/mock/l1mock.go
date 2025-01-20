package mock

import (
	"context"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
)

type L1ClientWrapper struct {
	Block  *types.Block
	TipCap *big.Int
}

func NewL1ClientWrapper() *L1ClientWrapper {
	return &L1ClientWrapper{}
}

func (c *L1ClientWrapper) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	return nil, nil
}

func (c *L1ClientWrapper) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (c *L1ClientWrapper) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	return nil, false, nil
}

func (c *L1ClientWrapper) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return 0, nil
}

func (c *L1ClientWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, nil
}

func (c *L1ClientWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return c.Block.Header(), nil
}

func (c *L1ClientWrapper) BlockNumber(ctx context.Context) (uint64, error) {
	return c.Block.NumberU64(), nil
}

func (c *L1ClientWrapper) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error) {
	ch <- c.Block.Header()
	return nil, nil
}

func (c *L1ClientWrapper) SetBlock(block *types.Block) {
	c.Block = block
}

func (c *L1ClientWrapper) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (c *L1ClientWrapper) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

func (c *L1ClientWrapper) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}

func (c *L1ClientWrapper) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}

func (c *L1ClientWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (c *L1ClientWrapper) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return c.TipCap, nil
}

func (c *L1ClientWrapper) EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error) {
	return 0, nil
}

func (c *L1ClientWrapper) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

func (c *L1ClientWrapper) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

func (c *L1ClientWrapper) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}
