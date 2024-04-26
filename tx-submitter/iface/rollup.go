package iface

import (
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

type IRollup interface {
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	CommitBatch(opts *bind.TransactOpts, batchDataInput bindings.IRollupBatchDataInput, batchSignatureInput bindings.IRollupBatchSignatureInput) (*types.Transaction, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizeBatch(opts *bind.TransactOpts, _batchIndex *big.Int) (*types.Transaction, error)
	// will be used in next version
	//FinalizationPeriodSeconds(opts *bind.CallOpts) (*big.Int, error)
	BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error)
}

// IL2Sequencer is the interface for the sequencer on L2
type IL1Staking interface {
	IsStaker(opts *bind.CallOpts, addr common.Address) (bool, error)
}
