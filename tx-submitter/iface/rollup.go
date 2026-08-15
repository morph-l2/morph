package iface

import (
	"math/big"

	"morph-l2/bindings/bindings"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
)

type IRollup interface {
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizeBatch(*bind.TransactOpts, []byte) (*types.Transaction, error)
	BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	BatchBlobVersionedHashes(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	}, error)

	FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupCommitBatchIterator, error)
	FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupFinalizeBatchIterator, error)
}

type ISubmitter interface {
	IsActive(opts *bind.CallOpts, addr common.Address) (bool, error)
}
