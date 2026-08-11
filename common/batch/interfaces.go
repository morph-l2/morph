package batch

import (
	"context"
	"errors"
	"math/big"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
)

// ErrKeyNotFound is returned by SealedBatchKV implementations when a key is absent.
var ErrKeyNotFound = errors.New("batch storage: key not found")

// KVPair is a key/value entry applied as part of an atomic WriteBatch.
type KVPair struct {
	Key   []byte
	Value []byte
}

// SealedBatchKV is a minimal key-value store used by BatchStorage.
type SealedBatchKV interface {
	GetBytes(key []byte) ([]byte, error)
	PutBytes(key, val []byte) error
	Delete(key []byte) error
	// WriteBatch applies all puts and deletes as a single atomic write, so that
	// batch data, batch header and the indices snapshot can never get out of sync
	// with each other on crash or partial failure.
	WriteBatch(puts []KVPair, deletes [][]byte) error
	// IteratePrefixKeys returns every key currently stored under prefix. It backs
	// the force-wipe self-heal path, which must remove orphaned sealed batch data
	// and header keys even when the indices snapshot is corrupt or unreadable.
	IteratePrefixKeys(prefix []byte) ([][]byte, error)
}

// L1HeaderClient is the L1 RPC surface required to recover batch headers from events.
type L1HeaderClient interface {
	bind.ContractCaller
	BlockNumber(ctx context.Context) (uint64, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*ethtypes.Header, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*ethtypes.Transaction, bool, error)
}

// L2MultiClient fans out read calls across multiple L2 endpoints (same role as tx-submitter iface.L2Clients).
type L2MultiClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error)
	Len() int
}

// SingleL2Client adapts a single L2 RPC backend as L2MultiClient (Len is always 1).
type SingleL2Client struct {
	C interface {
		BlockNumber(ctx context.Context) (uint64, error)
		BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error)
	}
}

func (s *SingleL2Client) BlockNumber(ctx context.Context) (uint64, error) {
	return s.C.BlockNumber(ctx)
}

func (s *SingleL2Client) BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error) {
	return s.C.BlockByNumber(ctx, number)
}

func (s *SingleL2Client) Len() int { return 1 }

// RollupBatchReader is the rollup contract view BatchCache needs (subset of generated Rollup bindings).
type RollupBatchReader interface {
	BatchBlobVersionedHashes(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}, error)
	FinalizedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	MessageQueue(opts *bind.CallOpts) (common.Address, error)
	FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupCommitBatchIterator, error)
	FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupRevertBatchIterator, error)
}

// L2GovCaller is the sole L2 contract view needed by batch assembly.
type L2GovCaller interface {
	GetTreeRoot(opts *bind.CallOpts) ([32]byte, error)
}
