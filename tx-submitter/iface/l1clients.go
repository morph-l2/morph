package iface

import (
	"context"
	"errors"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
)

// L1Clients wraps several L1 RPC clients. Every call is tried against the
// clients in order; when a client returns an error (including ethereum.NotFound
// for transaction/receipt lookups) the next client is tried. If all clients
// fail, the error from the last client is returned, so a NotFound result is
// preserved when every client reports NotFound.
type L1Clients struct {
	Clients []L1Client
}

var _ L1Client = (*L1Clients)(nil)

// NewL1Clients builds a failover L1 client. At least one client is required.
func NewL1Clients(clients []L1Client) (*L1Clients, error) {
	if len(clients) == 0 {
		return nil, errors.New("no L1 clients available")
	}
	return &L1Clients{Clients: clients}, nil
}

// Len returns the number of configured L1 clients.
func (c *L1Clients) Len() int { return len(c.Clients) }

// tryAll runs fn on each client in order until one succeeds.
func (c *L1Clients) tryAll(fn func(L1Client) error) error {
	if len(c.Clients) == 0 {
		return errors.New("no L1 clients available")
	}
	var lastErr error
	for _, client := range c.Clients {
		err := fn(client)
		if err == nil {
			return nil
		}
		lastErr = err
	}
	return lastErr
}

func (c *L1Clients) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (res []byte, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.CodeAt(ctx, contract, blockNumber); return })
	return
}

func (c *L1Clients) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) (res []byte, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.CallContract(ctx, call, blockNumber); return })
	return
}

func (c *L1Clients) PendingCodeAt(ctx context.Context, account common.Address) (res []byte, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.PendingCodeAt(ctx, account); return })
	return
}

func (c *L1Clients) PendingNonceAt(ctx context.Context, account common.Address) (res uint64, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.PendingNonceAt(ctx, account); return })
	return
}

func (c *L1Clients) SuggestGasPrice(ctx context.Context) (res *big.Int, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.SuggestGasPrice(ctx); return })
	return
}

func (c *L1Clients) SuggestGasTipCap(ctx context.Context) (res *big.Int, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.SuggestGasTipCap(ctx); return })
	return
}

func (c *L1Clients) EstimateGas(ctx context.Context, call ethereum.CallMsg) (res uint64, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.EstimateGas(ctx, call); return })
	return
}

func (c *L1Clients) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.tryAll(func(cl L1Client) error { return cl.SendTransaction(ctx, tx) })
}

func (c *L1Clients) FilterLogs(ctx context.Context, query ethereum.FilterQuery) (res []types.Log, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.FilterLogs(ctx, query); return })
	return
}

func (c *L1Clients) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (res ethereum.Subscription, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.SubscribeFilterLogs(ctx, query, ch); return })
	return
}

func (c *L1Clients) SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (res ethereum.Subscription, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.SubscribeNewHead(ctx, ch); return })
	return
}

func (c *L1Clients) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	err = c.tryAll(func(cl L1Client) (e error) {
		tx, isPending, e = cl.TransactionByHash(ctx, hash)
		if e == nil && tx == nil {
			return ethereum.NotFound
		}
		return e
	})
	return
}

func (c *L1Clients) TransactionReceipt(ctx context.Context, txHash common.Hash) (res *types.Receipt, err error) {
	err = c.tryAll(func(cl L1Client) (e error) {
		res, e = cl.TransactionReceipt(ctx, txHash)
		if e == nil && res == nil {
			return ethereum.NotFound
		}
		return e
	})
	return
}

func (c *L1Clients) BlockByNumber(ctx context.Context, number *big.Int) (res *types.Block, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.BlockByNumber(ctx, number); return })
	return
}

func (c *L1Clients) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (res uint64, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.NonceAt(ctx, account, blockNumber); return })
	return
}

func (c *L1Clients) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (res *big.Int, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.BalanceAt(ctx, account, blockNumber); return })
	return
}

func (c *L1Clients) HeaderByNumber(ctx context.Context, number *big.Int) (res *types.Header, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.HeaderByNumber(ctx, number); return })
	return
}

func (c *L1Clients) BlockNumber(ctx context.Context) (res uint64, err error) {
	err = c.tryAll(func(cl L1Client) (e error) { res, e = cl.BlockNumber(ctx); return })
	return
}
