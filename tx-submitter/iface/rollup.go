package iface

import (
	"math/big"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
)

type IRollup interface {
	BatchBlobVersionedHashes(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	CommitBatch(opts *bind.TransactOpts, batchDataInput bindings.IRollupBatchDataInput) (*types.Transaction, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	MessageQueue(opts *bind.CallOpts) (common.Address, error)
	FinalizeBatch(*bind.TransactOpts, []byte) (*types.Transaction, error)
	BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}, error)

	FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupCommitBatchIterator, error)
	FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupFinalizeBatchIterator, error)
	FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupRevertBatchIterator, error)
}

// ISubmitter is the minimum qualification surface required by the official
// submitter. Selection/rotation is deliberately not part of this interface.
type ISubmitter interface {
	IsActive(opts *bind.CallOpts, addr common.Address) (bool, error)
}

type IL2MessagePasser interface {
	GetTreeRoot(opts *bind.CallOpts) ([32]byte, error)
}
