package iface

import (
	"context"
	"math/big"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/eth"
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
	PendingTransactionCount(ctx context.Context) (uint, error)
}

type L2Client interface {
	Client
	GetBlockTraceByNumber(ctx context.Context, number *big.Int) (*types.BlockTrace, error)
	GetRollupBatchByIndex(ctx context.Context, batchIndex uint64) (*eth.RPCRollupBatch, error)
}
