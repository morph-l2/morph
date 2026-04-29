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
	CommitBatch(opts *bind.TransactOpts, batchDataInput bindings.IRollupBatchDataInput, batchSignatureInput bindings.IRollupBatchSignatureInput) (*types.Transaction, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizeBatch(*bind.TransactOpts, []byte) (*types.Transaction, error)
	BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
	CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error)
	BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		BlockNumber            *big.Int
		SignedSequencersBitmap *big.Int
	}, error)

	FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupCommitBatchIterator, error)
	FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*bindings.RollupFinalizeBatchIterator, error)
}

// IL2Sequencer is the interface for the sequencer on L2
type IL2Sequencer interface {
	SequencerSetVerifyHash(opts *bind.CallOpts) ([32]byte, error)
}

type IL2Gov interface {
	RollupEpoch(opts *bind.CallOpts) (*big.Int, error)
	BatchBlockInterval(opts *bind.CallOpts) (*big.Int, error)
	BatchTimeout(opts *bind.CallOpts) (*big.Int, error)
}

type IL1Staking interface {
	IsStaker(opts *bind.CallOpts, addr common.Address) (bool, error)
	GetStakersBitmap(opts *bind.CallOpts, _stakers []common.Address) (*big.Int, error)
	GetActiveStakers(opts *bind.CallOpts) ([]common.Address, error)
	GetStakers(opts *bind.CallOpts) ([255]common.Address, error)
}

type IL2MessagePasser interface {
	GetTreeRoot(opts *bind.CallOpts) ([32]byte, error)
}
