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
}

// L1HeaderClient is the L1 RPC surface required to recover batch headers from events.
type L1HeaderClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
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
	CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		BlockNumber            *big.Int
		SignedSequencersBitmap *big.Int
	}, error)
	FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupFinalizeBatchIterator, error)
}

// L2GovCaller reads batch-related Gov / bridge / sequencer data on L2.
type L2GovCaller interface {
	BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error)
	BatchTimeout(opts *bind.CallOpts) (*big.Int, error)
	GetTreeRoot(opts *bind.CallOpts) ([32]byte, error)
	GetSequencerSetBytes(opts *bind.CallOpts) ([]byte, common.Hash, error)
}
