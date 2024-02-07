package iface

import (
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
)

type IRollup interface {
	LatestL2BlockNumber(opts *bind.CallOpts) (*big.Int, error)
	//SubmitBatches(opts *bind.TransactOpts, batches []bindings.RollupBatchData) (*types.Transaction, error)
	//IsStaked(opts *bind.CallOpts, addr common.Address) (bool, error)
	//Stake(opts *bind.TransactOpts) (*types.Transaction, error)
	LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	CommitBatch(opts *bind.TransactOpts, batchData bindings.IRollupBatchData, minGasLimit uint32) (*types.Transaction, error)
	Stake(opts *bind.TransactOpts) (*types.Transaction, error)
	IsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error)
	MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error)
	Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)
	LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error)
	FinalizeBatchsByNum(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error)
	FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error)
	CommittedBatchStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
		BatchHash              [32]byte
		OriginTimestamp        *big.Int
		PrevStateRoot          [32]byte
		PostStateRoot          [32]byte
		WithdrawalRoot         [32]byte
		DataHash               [32]byte
		Sequencer              common.Address
		L1MessagePopped        *big.Int
		TotalL1MessagePopped   *big.Int
		SkippedL1MessageBitmap []byte
		BlockNumber            *big.Int
	}, error)
}

// ISubmitter is the interface for the submitter on L2
type ISubmitter interface {
	GetNextSubmitter(opts *bind.CallOpts) (common.Address, *big.Int, *big.Int, error)
}
