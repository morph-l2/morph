package iface

import (
	"context"
	"errors"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth"
)

type Client interface {
	bind.ContractBackend
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	BlockNumber(ctx context.Context) (uint64, error)
}

type L1Client interface {
	Client
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

type L2Client interface {
	Client
	GetBlockTraceByNumber(ctx context.Context, number *big.Int) (*types.BlockTrace, error)
	GetRollupBatchByIndex(ctx context.Context, batchIndex uint64) (*eth.RPCRollupBatch, error)
}

type L2Clients struct {
	Clients []L2Client
}

// getFirstClient returns the first available client, or an error if no clients are available
func (c *L2Clients) getFirstClient() (L2Client, error) {
	if len(c.Clients) == 0 {
		return nil, errors.New("no L2 clients available")
	}
	return c.Clients[0], nil
}

// tryAllClients tries all clients until one succeeds, returns the last error if all fail
func (c *L2Clients) tryAllClients(fn func(L2Client) error) error {
	if len(c.Clients) == 0 {
		return errors.New("no L2 clients available")
	}
	var lastErr error
	for _, client := range c.Clients {
		if err := fn(client); err == nil {
			return nil
		} else {
			lastErr = err
		}
	}
	return lastErr
}

// CodeAt implements bind.ContractCaller
func (c *L2Clients) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var result []byte
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.CodeAt(ctx, contract, blockNumber)
		return err
	})
	return result, err
}

// CallContract implements bind.ContractCaller
func (c *L2Clients) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var result []byte
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.CallContract(ctx, call, blockNumber)
		return err
	})
	return result, err
}

// PendingCodeAt implements bind.PendingContractCaller and bind.ContractTransactor
func (c *L2Clients) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return nil, err
	}
	return client.PendingCodeAt(ctx, account)
}

// PendingNonceAt implements bind.ContractTransactor
func (c *L2Clients) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return 0, err
	}
	return client.PendingNonceAt(ctx, account)
}

// SuggestGasPrice implements bind.ContractTransactor
func (c *L2Clients) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return nil, err
	}
	return client.SuggestGasPrice(ctx)
}

// SuggestGasTipCap implements bind.ContractTransactor
func (c *L2Clients) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return nil, err
	}
	return client.SuggestGasTipCap(ctx)
}

// EstimateGas implements bind.ContractTransactor
func (c *L2Clients) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return 0, err
	}
	return client.EstimateGas(ctx, call)
}

// SendTransaction implements bind.ContractTransactor
func (c *L2Clients) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	client, err := c.getFirstClient()
	if err != nil {
		return err
	}
	return client.SendTransaction(ctx, tx)
}

// FilterLogs implements bind.ContractFilterer
func (c *L2Clients) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.FilterLogs(ctx, query)
		return err
	})
	return result, err
}

// SubscribeFilterLogs implements bind.ContractFilterer
func (c *L2Clients) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	client, err := c.getFirstClient()
	if err != nil {
		return nil, err
	}
	return client.SubscribeFilterLogs(ctx, query, ch)
}

// TransactionByHash implements Client
func (c *L2Clients) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	err = c.tryAllClients(func(client L2Client) error {
		var e error
		tx, isPending, e = client.TransactionByHash(ctx, hash)
		return e
	})
	return tx, isPending, err
}

// BlockByNumber implements Client
func (c *L2Clients) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	var result *types.Block
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.BlockByNumber(ctx, number)
		return err
	})
	return result, err
}

// NonceAt implements Client
func (c *L2Clients) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var result uint64
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.NonceAt(ctx, account, blockNumber)
		return err
	})
	return result, err
}

// TransactionReceipt implements Client
func (c *L2Clients) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var result *types.Receipt
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.TransactionReceipt(ctx, txHash)
		return err
	})
	return result, err
}

// BalanceAt implements Client
func (c *L2Clients) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var result *big.Int
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.BalanceAt(ctx, account, blockNumber)
		return err
	})
	return result, err
}

// HeaderByNumber implements Client and bind.ContractTransactor
func (c *L2Clients) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var result *types.Header
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.HeaderByNumber(ctx, number)
		return err
	})
	return result, err
}

// BlockNumber implements Client
func (c *L2Clients) BlockNumber(ctx context.Context) (uint64, error) {
	var result uint64
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.BlockNumber(ctx)
		return err
	})
	return result, err
}

// GetBlockTraceByNumber implements L2Client
func (c *L2Clients) GetBlockTraceByNumber(ctx context.Context, number *big.Int) (*types.BlockTrace, error) {
	var result *types.BlockTrace
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.GetBlockTraceByNumber(ctx, number)
		return err
	})
	return result, err
}

// GetRollupBatchByIndex implements L2Client
func (c *L2Clients) GetRollupBatchByIndex(ctx context.Context, batchIndex uint64) (*eth.RPCRollupBatch, error) {
	var result *eth.RPCRollupBatch
	err := c.tryAllClients(func(client L2Client) error {
		var err error
		result, err = client.GetRollupBatchByIndex(ctx, batchIndex)
		return err
	})
	return result, err
}
