// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IRollupBatchDataInput is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchDataInput struct {
	Version           uint8
	ParentBatchHeader []byte
	LastBlockNumber   uint64
	NumL1Messages     uint16
	PrevStateRoot     [32]byte
	PostStateRoot     [32]byte
	WithdrawalRoot    [32]byte
}

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"name\":\"ChallengeRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChallengeRewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"}],\"name\":\"ChallengeState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProveRemainingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddr\",\"type\":\"address\"}],\"name\":\"SubmitterContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateChallenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateFinalizationPeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"UpdateProofRewardPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"UpdateProofWindow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateRollupDelayPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LAYER_2_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__maxNumTxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchBlobVersionedHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blobVersionedHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"batchChallengeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchChallenged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchChallengedSuccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchDataStore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInsideChallengeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_batchHash\",\"type\":\"bytes32\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"challengeSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimProveRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"lastBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numL1Messages\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIRollup.BatchDataInput\",\"name\":\"batchDataInput\",\"type\":\"tuple\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"lastBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numL1Messages\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIRollup.BatchDataInput\",\"name\":\"batchDataInput\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_batchProof\",\"type\":\"bytes\"}],\"name\":\"commitBatchWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"lastBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numL1Messages\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIRollup.BatchDataInput\",\"name\":\"batchDataInput\",\"type\":\"tuple\"}],\"name\":\"commitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationPeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"}],\"name\":\"finalizeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitterContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofRewardPercent\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_prevStateRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rollupDelayPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSubmitter\",\"type\":\"address\"}],\"name\":\"initialize4\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"isChallenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isChallenger\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCommittedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofRewardPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_batchProof\",\"type\":\"bytes\"}],\"name\":\"proveCommittedBatchState\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proveRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_batchProof\",\"type\":\"bytes\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertReqIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupDelayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitterContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateFinalizePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newWindow\",\"type\":\"uint256\"}],\"name\":\"updateProofWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newProofRewardPercent\",\"type\":\"uint256\"}],\"name\":\"updateRewardPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateRollupDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b50604051620055d4380380620055d483398101604081905262000033916200010f565b6001600160401b0381166080526200004a62000051565b506200013e565b5f54610100900460ff1615620000bd5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146200010d575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f6020828403121562000120575f80fd5b81516001600160401b038116811462000137575f80fd5b9392505050565b6080516154766200015e5f395f81816106a0015261451901526154765ff3fe60806040526004361061035e575f3560e01c806388b1ea09116101bd578063b8d0a1b0116100f2578063d8dc99d211610092578063e3fff1dd1161006d578063e3fff1dd14610af6578063f2fde38b14610b15578063f8fa010f14610b34578063fb1e8b0414610b53575f80fd5b8063d8dc99d214610aad578063de8b303514610ac2578063dff7827e14610ae1575f80fd5b8063cd4edc69116100cd578063cd4edc6914610a47578063ce5db8d614610a5a578063cf9a674514610a6f578063d279c19114610a8e575f80fd5b8063b8d0a1b0146109ea578063bedb86fb14610a09578063c555389214610a28575f80fd5b8063a479265d1161015d578063b31a77d311610138578063b31a77d314610978578063b34844251461098d578063b35dac4e146109ac578063b3e0a509146109cb575f80fd5b8063a479265d14610919578063a4f209b01461092e578063abc8d68d1461094d575f80fd5b8063910129d411610198578063910129d41461087c57806397fc007c146108ad578063a3c9e94c146108cc578063a415d8dc146108eb575f80fd5b806388b1ea09146107965780638da5cb5b146107af5780638f1d3776146107cc575f80fd5b80632b7ac3f3116102935780635f77cf1d1161023357806368589dfa1161020e57806368589dfa146107195780636c578c1d14610744578063715018a614610763578063728cdbca14610777575f80fd5b80635f77cf1d1461068f57806361267290146106db57806367caa37a146106fa575f80fd5b806341f756da1161026e57806341f756da146105d457806357e0af6c146105e75780635c975abb146106065780635ef7a94a1461061d575f80fd5b80632b7ac3f3146105695780633b70c18a146105a05780633e001b66146105bf575f80fd5b80631544ba3a116102fe57806321e2f9e0116102d957806321e2f9e0146104c95780632362f03e146104e85780632571098d146105135780632a213ba11461053e575f80fd5b80631544ba3a1461046557806318463fb01461048457806318af3b2b14610499575f80fd5b806310d445831161033957806310d44583146103f0578063116a1f421461040f578063121dcd50146104315780631336110114610446575f80fd5b806304d7721514610369578063059def61146103ac5780630ceb6780146103cf575f80fd5b3661036557005b5f80fd5b348015610374575f80fd5b50610397610383366004614ed4565b60a36020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156103b7575f80fd5b506103c1609d5481565b6040519081526020016103a3565b3480156103da575f80fd5b506103ee6103e9366004614f01565b610b68565b005b3480156103fb575f80fd5b506103ee61040a366004614f5f565b610c3a565b34801561041a575f80fd5b50610397610429366004614ed4565b609d54101590565b34801561043c575f80fd5b506103c1609e5481565b348015610451575f80fd5b506103ee610460366004614fa7565b610f76565b348015610470575f80fd5b506103ee61047f366004614ff6565b611493565b34801561048f575f80fd5b506103c160a75481565b3480156104a4575f80fd5b506103976104b3366004614ed4565b5f90815260a26020526040902060010154421090565b3480156104d4575f80fd5b506103976104e3366004614ed4565b611852565b3480156104f3575f80fd5b506103c1610502366004614ed4565b60a16020525f908152604090205481565b34801561051e575f80fd5b506103c161052d366004614ed4565b60a06020525f908152604090205481565b348015610549575f80fd5b506103c1610558366004614ed4565b60ad6020525f908152604090205481565b348015610574575f80fd5b50609c54610588906001600160a01b031681565b6040516001600160a01b0390911681526020016103a3565b3480156105ab575f80fd5b50609b54610588906001600160a01b031681565b3480156105ca575f80fd5b506103c160aa5481565b6103ee6105e2366004615084565b611880565b3480156105f2575f80fd5b506103ee610601366004614ed4565b611b84565b348015610611575f80fd5b5060655460ff16610397565b348015610628575f80fd5b50610666610637366004614ed4565b60a26020525f90815260409020805460018201546002830154600390930154919290916001600160a01b031684565b604080519485526020850193909352918301526001600160a01b031660608201526080016103a3565b34801561069a575f80fd5b506106c27f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff90911681526020016103a3565b3480156106e6575f80fd5b506103ee6106f5366004614ed4565b611c2f565b348015610705575f80fd5b506103ee610714366004615084565b611e0a565b348015610724575f80fd5b506103c1610733366004614ed4565b60ab6020525f908152604090205481565b34801561074f575f80fd5b506103ee61075e366004614f01565b612155565b34801561076e575f80fd5b506103ee612212565b348015610782575f80fd5b506103ee6107913660046150be565b612225565b3480156107a1575f80fd5b5060a6546103979060ff1681565b3480156107ba575f80fd5b506033546001600160a01b0316610588565b3480156107d7575f80fd5b506108386107e6366004614ed4565b60a46020525f9081526040902080546001820154600283015460039093015467ffffffffffffffff831693680100000000000000009093046001600160a01b0316929060ff8082169161010090041686565b6040805167ffffffffffffffff90971687526001600160a01b03909516602087015293850192909252606084015215156080830152151560a082015260c0016103a3565b348015610887575f80fd5b50610397610896366004614ed4565b5f90815260a4602052604090206003015460ff1690565b3480156108b8575f80fd5b506103ee6108c7366004614f01565b61252e565b3480156108d7575f80fd5b50609754610588906001600160a01b031681565b3480156108f6575f80fd5b50610397610905366004614f01565b609f6020525f908152604090205460ff1681565b348015610924575f80fd5b506103c160995481565b348015610939575f80fd5b506103ee610948366004614ed4565b612611565b348015610958575f80fd5b506103c1610967366004614f01565b60a56020525f908152604090205481565b348015610983575f80fd5b506103c160a85481565b348015610998575f80fd5b506103ee6109a7366004615119565b6126c1565b3480156109b7575f80fd5b506103ee6109c6366004614f01565b6129ce565b3480156109d6575f80fd5b506103ee6109e5366004614fa7565b612a27565b3480156109f5575f80fd5b506103ee610a04366004615119565b612db8565b348015610a14575f80fd5b506103ee610a2336600461518d565b612e41565b348015610a33575f80fd5b506103ee610a42366004614ed4565b612f7f565b6103ee610a553660046151bf565b613022565b348015610a65575f80fd5b506103c160985481565b348015610a7a575f80fd5b506103ee610a89366004614ed4565b6135b4565b348015610a99575f80fd5b506103ee610aa8366004614f01565b613723565b348015610ab8575f80fd5b506103c160ac5481565b348015610acd575f80fd5b50610397610adc366004614ed4565b6137e2565b348015610aec575f80fd5b506103c1609a5481565b348015610b01575f80fd5b506103ee610b10366004614ed4565b61382c565b348015610b20575f80fd5b506103ee610b2f366004614f01565b6138cf565b348015610b3f575f80fd5b506103ee610b4e366004614f01565b61395c565b348015610b5e575f80fd5b506103c160a95481565b610b70613c0d565b6001600160a01b0381165f908152609f602052604090205460ff1615610bdd5760405162461bcd60e51b815260206004820152601f60248201527f6163636f756e7420697320616c72656164792061206368616c6c656e6765720060448201526064015b60405180910390fd5b6001600160a01b0381165f818152609f6020908152604091829020805460ff1916600190811790915591519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc00991015b60405180910390a250565b610c42613c0d565b5f8111610c915760405162461bcd60e51b815260206004820152601560248201527f636f756e74206d757374206265206e6f6e7a65726f00000000000000000000006044820152606401610bd4565b5f80610c9d8585613c67565b915091505f610cb0836001015160c01c90565b5f81815260a160205260409020549091508214610d0f5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b5f60a181610d1d8785615214565b81526020019081526020015f205414610d9d5760405162461bcd60e51b8152602060048201526024808201527f726576657274696e67206d7573742073746172742066726f6d2074686520656e60448201527f64696e67000000000000000000000000000000000000000000000000000000006064820152608401610bd4565b609d548111610e145760405162461bcd60e51b815260206004820152602160248201527f63616e206f6e6c792072657665727420756e46696e616c697a6564206261746360448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610bd4565b610e1f600182615227565b609e555b8315610f6e57604051829082907ecae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146905f90a35f81815260a16020526040812055610e6b816137e2565b15610ec6575f81815260a460209081526040808320600181015490546801000000000000000090046001600160a01b0316845260a59092528220805491929091610eb6908490615214565b909155505060a6805460ff191690555b5f81815260a46020526040812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600181018290556002810191909155600301805461ffff1916905560a85415801590610f26575060a85481145b15610f30575f60a8555b6001015f81815260a160205260409020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90940193915081610e23575b505050505050565b60a85415610fc65760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b610fce613d0f565b5f80610fda8484613c67565b915091505f610fed836001015160c01c90565b5f81815260a16020526040902054909150821461104c5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b61105581611852565b6110a15760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610bd4565b6110aa816137e2565b156110f75760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610bd4565b5f81815260a4602052604090206003015460ff16156111585760405162461bcd60e51b815260206004820152601660248201527f62617463682073686f756c6420626520726576657274000000000000000000006044820152606401610bd4565b5f81815260a260205260409020600101544210156111b85760405162461bcd60e51b815260206004820152601960248201527f626174636820696e206368616c6c656e67652077696e646f77000000000000006044820152606401610bd4565b605983015160a05f6111cb600185615227565b81526020019081526020015f2054146112265760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610bd4565b5f81815260a06020526040902054156112815760405162461bcd60e51b815260206004820152601660248201527f626174636820616c7265616479207665726966696564000000000000000000006044820152606401610bd4565b80609d54600101146112d55760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610bd4565b609d819055600160a35f6112ea866099015190565b815260208101919091526040015f20805460ff191691151591909117905560798301515f82815260a0602052604090205561133961132c846011015160c01c90565b600985015160c01c613d62565b60a25f611347600184615227565b815260208101919091526040015f908120818155600180820183905560028201839055600390910180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560ab91906113a39084615227565b81526020019081526020015f205f905560ad5f6001836113c39190615227565b81526020019081526020015f205f905560a45f6001836113e39190615227565b815260208082019290925260409081015f90812080547fffffffff000000000000000000000000000000000000000000000000000000001681556001810182905560028101829055600301805461ffff1916905583815260a1909252902054817f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d61146f866079015190565b60998701516040805192835260208301919091520160405180910390a35050505050565b60a854156114e35760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b6114eb613d0f565b60ac54609e545f90815260a260205260408120549091429161150d9190615214565b1090505f4260ac54609b5f9054906101000a90046001600160a01b03166001600160a01b031663b59b1a786040518163ffffffff1660e01b8152600401602060405180830381865afa158015611565573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611589919061523a565b6115939190615214565b109050811580156115a15750805b15611608575f6115b76080890160608a01615251565b61ffff16116116085760405162461bcd60e51b815260206004820152600b60248201527f6c316d73672064656c61790000000000000000000000000000000000000000006044820152606401610bd4565b81806116115750805b61165d5760405162461bcd60e51b815260206004820152600e60248201527f696e76616c69642074696d696e670000000000000000000000000000000000006044820152606401610bd4565b5f61167361166e60208a018a615272565b613c67565b5090505f611685826001015160c01c90565b611690906001615214565b5f81815260ad60205260408120549192509015611730575f491561171c5760405162461bcd60e51b815260206004820152602f60248201527f6d757374206e6f7420636172727920626c6f62207768656e207573696e67207360448201527f746f72656420626c6f62206861736800000000000000000000000000000000006064820152608401610bd4565b505f81815260ad602052604090205461174b565b61174861174060208c018c6152d3565b60ff16613e19565b90505b6117568a5f83613f54565b5f806117628b8b613c67565b915091505f611775836001015160c01c90565b905080609e54146117c85760405162461bcd60e51b815260206004820152601660248201527f696e636f727265637420626174636820686561646572000000000000000000006044820152606401610bd4565b5f81815260a1602052604090205482146118245760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b5f81815260a26020526040902042600190910155611843838b8b6144aa565b50505050505050505050505050565b5f81815260a260205260408120541580159061187a57505f82815260a1602052604090205415155b92915050565b6097546001600160a01b0316639f8a13d7336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156118ec573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061191091906152f3565b61195c5760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c7920616374697665207375626d697474657220616c6c6f7765640000006044820152606401610bd4565b60a854156119ac5760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b6119b4613d0f565b5f6119c561166e6020840184615272565b5090505f6119d7826001015160c01c90565b6119e2906001615214565b5f81815260ad602052604090205490915015611a665760405162461bcd60e51b815260206004820152602860248201527f636f6d6d69744261746368207265717569726573206e6f2073746f726564206260448201527f6c6f6220686173680000000000000000000000000000000000000000000000006064820152608401610bd4565b60ac54609b54604080517fb59b1a7800000000000000000000000000000000000000000000000000000000815290514293926001600160a01b03169163b59b1a789160048083019260209291908290030181865afa158015611aca573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611aee919061523a565b611af89190615214565b1015611b60575f611b0f6080850160608601615251565b61ffff1611611b605760405162461bcd60e51b815260206004820152600b60248201527f6c316d73672064656c61790000000000000000000000000000000000000000006044820152606401610bd4565b5f611b7161174060208601866152d3565b9050611b7e843383613f54565b50505050565b611b8c613c0d565b5f81118015611b9d57506099548114155b611be95760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964206e65772070726f6f662077696e646f7700000000000000006044820152606401610bd4565b609980549082905560408051828152602081018490527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a6191015b60405180910390a15050565b5f54600290610100900460ff16158015611c4f57505f5460ff8083169116105b611cc15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610bd4565b5f805461ffff191660ff831617610100179055611cdf5f5460ff1690565b60ff16600214611d315760405162461bcd60e51b815260206004820152601660248201527f6d757374206861766520696e697469616c697a656421000000000000000000006044820152606401610bd4565b81611da45760405162461bcd60e51b815260206004820152602760248201527f63616e206e6f742073657420737461746520726f6f742077697468206279746560448201527f73333228302921000000000000000000000000000000000000000000000000006064820152608401610bd4565b609e545f90815260ab6020526040902054611dcd57609e545f90815260ab602052604090208290555b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611c23565b6097546001600160a01b0316639f8a13d7336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611e76573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e9a91906152f3565b611ee65760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c7920616374697665207375626d697474657220616c6c6f7765640000006044820152606401610bd4565b60a85415611f365760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b611f3e613d0f565b5f4915611f8d5760405162461bcd60e51b815260206004820152601f60248201527f636f6d6d69745374617465206d757374206e6f7420636172727920626c6f62006044820152606401610bd4565b5f611f9e61166e6020840184615272565b5090505f611fb0826001015160c01c90565b611fbb906001615214565b5f81815260ad602052604090205490915061203e5760405162461bcd60e51b815260206004820152602260248201527f6e6f2073746f72656420626c6f62206861736820666f7220746869732062617460448201527f63680000000000000000000000000000000000000000000000000000000000006064820152608401610bd4565b60ac54609b54604080517fb59b1a7800000000000000000000000000000000000000000000000000000000815290514293926001600160a01b03169163b59b1a789160048083019260209291908290030181865afa1580156120a2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120c6919061523a565b6120d09190615214565b1015612138575f6120e76080850160608601615251565b61ffff16116121385760405162461bcd60e51b815260206004820152600b60248201527f6c316d73672064656c61790000000000000000000000000000000000000000006044820152606401610bd4565b61215083335f84815260ad6020526040902054613f54565b505050565b61215d613c0d565b6001600160a01b0381165f908152609f602052604090205460ff166121c45760405162461bcd60e51b815260206004820152601b60248201527f6163636f756e74206973206e6f742061206368616c6c656e67657200000000006044820152606401610bd4565b6001600160a01b0381165f818152609f60209081526040808320805460ff19169055519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc0099101610c2f565b61221a613c0d565b6122235f61465d565b565b5f54610100900460ff161580801561224357505f54600160ff909116105b8061225c5750303b15801561225c57505f5460ff166001145b6122ce5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610bd4565b5f805460ff1916600117905580156122ef575f805461ff0019166101001790555b6001600160a01b038616158061230c57506001600160a01b038516155b15612343576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0387166123995760405162461bcd60e51b815260206004820152601a60248201527f696e76616c6964207375626d697474657220636f6e74726163740000000000006044820152606401610bd4565b6123a16146c6565b6123a961474a565b609780546001600160a01b03808a167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609b8054898416908316179055609c805492881692909116821790556098859055609984905560a98390556040515f907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96908290a3604080515f8152602081018690527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437910160405180910390a1604080515f8152602081018590527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61910160405180910390a1604080515f8152602081018490527ffb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b0223910160405180910390a18015612525575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b612536613c0d565b6001600160a01b0381161580159061255c5750609c546001600160a01b03828116911614155b6125a85760405162461bcd60e51b815260206004820152601460248201527f696e76616c6964206e65772076657269666965720000000000000000000000006044820152606401610bd4565b609c80546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96905f90a35050565b612619613c0d565b5f81118015612629575060648111155b8015612637575060a9548114155b6126835760405162461bcd60e51b815260206004820152601f60248201527f696e76616c69642070726f6f66207265776172642070657263656e74616765006044820152606401610bd4565b60a980549082905560408051828152602081018490527ffb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b02239101611c23565b60a854156127115760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b612719613d0f565b6097546001600160a01b0316639f8a13d7336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015612785573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127a991906152f3565b6127f55760405162461bcd60e51b815260206004820152601d60248201527f6f6e6c7920616374697665207375626d697474657220616c6c6f7765640000006044820152606401610bd4565b5f806128018686613c67565b915091505f612814836001015160c01c90565b5f81815260a1602052604090205490915082146128735760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b61287c816137e2565b6128c85760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610bd4565b5f81815260a46020526040902060038101805461ff00191661010017905560a6805460ff19169055609954600290910154429161290491615214565b11612983575f81815260a4602090815260408083206003908101805460ff1916600117905560a2835292819020909201548251808401909352600783527f54696d656f7574000000000000000000000000000000000000000000000000009183019190915261297e9183916001600160a01b0316906147ce565b612525565b61298e8386866144aa565b61252581336040518060400160405280600d81526020017f50726f6f6620737563636573730000000000000000000000000000000000000081525061492d565b6129d6613c0d565b60aa80545f9091556129e882826149f9565b604080516001600160a01b0384168152602081018390527fb1b2058a6969e2d25e47bcaebe8ae21c29a23b2752429315b75e2f4f285f3d879101611c23565b612a2f613c0d565b5f805260a06020527fb84a74ec6ef4d0e83b6006dfaa014ab4026f9f3b97d186e604d29998a4e808ea5415612aa65760405162461bcd60e51b815260206004820152601660248201527f67656e6573697320626174636820696d706f72746564000000000000000000006044820152606401610bd4565b5f80612ab28484613c67565b915091505f612ac5836001015160c01c90565b90508015612b155760405162461bcd60e51b815260206004820152601360248201527f696e76616c696420626174636820696e646578000000000000000000000000006044820152606401610bd4565b5f612b21846079015190565b905080612b705760405162461bcd60e51b815260206004820152600f60248201527f7a65726f20737461746520726f6f7400000000000000000000000000000000006044820152606401610bd4565b600984015160c01c15612bc55760405162461bcd60e51b815260206004820152601d60248201527f6c31206d65737361676520706f707065642073686f756c6420626520300000006044820152606401610bd4565b5f612bd1856019015190565b03612c1e5760405162461bcd60e51b815260206004820152600e60248201527f7a65726f206461746120686173680000000000000000000000000000000000006044820152606401610bd4565b7f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014612c4a856039015190565b14612c975760405162461bcd60e51b815260206004820152601660248201527f696e76616c69642076657273696f6e65642068617368000000000000000000006044820152606401610bd4565b5f82815260a1602090815260408083208690558051608081018252428082528184019081528183018581526060830186815288875260a28652848720935184559151600184015551600283015551600390910180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909216919091179055603987015160ad83528184205560ab825280832084905560a0909152808220839055609e849055609d84905551849184917f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f9190a3604080518281525f6020820152849184917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a3505050505050565b5f80612dc48686613c67565b915091505f612dd7836001015160c01c90565b5f81815260a160205260409020549091508214612e365760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b6125258386866144aa565b612e49613c0d565b8015612f4d57612e57614aa2565b60a65460ff1615612f0c5760a7545f90815260a460209081526040808320600181015490546801000000000000000090046001600160a01b0316845260a59092528220805491929091612eab908490615214565b909155505060a7545f90815260a46020526040812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600181018290556002810191909155600301805461ffff1916905560a6805460ff191690555b7f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b612f55614afc565b7f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612f2f565b50565b612f87613c0d565b5f81118015612f98575060ac548114155b612fe45760405162461bcd60e51b815260206004820152601f60248201527f696e76616c6964206e657720726f6c6c75702064656c617920706572696f64006044820152606401610bd4565b60ac80549082905560408051828152602081018490527f624e47dc9fb8f8cfeaf4ead4710277cc1757136cfa885e465514cf6d510f0ad19101611c23565b335f908152609f602052604090205460ff166130805760405162461bcd60e51b815260206004820152601760248201527f6f6e6c79206368616c6c656e67657220616c6c6f7765640000000000000000006044820152606401610bd4565b60a854156130d05760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b6130d8613d0f565b60a65460ff161561312b5760405162461bcd60e51b815260206004820152601460248201527f616c726561647920696e206368616c6c656e67650000000000000000000000006044820152606401610bd4565b8167ffffffffffffffff16609d54106131865760405162461bcd60e51b815260206004820152601760248201527f626174636820616c72656164792066696e616c697a65640000000000000000006044820152606401610bd4565b67ffffffffffffffff82165f90815260a1602052604090205481146131ed5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610bd4565b6132008267ffffffffffffffff16611852565b61324c5760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610bd4565b67ffffffffffffffff82165f90815260a460205260409020546801000000000000000090046001600160a01b0316156132c75760405162461bcd60e51b815260206004820152601860248201527f626174636820616c7265616479206368616c6c656e67656400000000000000006044820152606401610bd4565b67ffffffffffffffff82165f90815260a2602052604090206001015442106133575760405162461bcd60e51b815260206004820152603360248201527f63616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f77000000000000000000000000006064820152608401610bd4565b60975f9054906101000a90046001600160a01b03166001600160a01b0316630d13fd7b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156133a7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133cb919061523a565b34101561341a5760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742076616c756500000000000000000000000000006044820152606401610bd4565b67ffffffffffffffff82811660a78190556040805160c0810182528281523360208083018281523484860190815242606086019081525f6080870181815260a0880182815299825260a4909552969096209451855492516001600160a01b031668010000000000000000027fffffffff000000000000000000000000000000000000000000000000000000009093169816979097171783559451600183015591516002820155925160039093018054925115156101000261ff00199415159490941661ffff19909316929092179290921790556001600160a01b03168267ffffffffffffffff167f3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b99763460405161353291815260200190565b60405180910390a35f609d54600161354a9190615214565b90505b609e5481116135a2578267ffffffffffffffff168114613590576099545f82815260a260205260408120600101805490919061358a908490615214565b90915550505b8061359a8161530e565b91505061354d565b505060a6805460ff1916600117905550565b5f54600390610100900460ff161580156135d457505f5460ff8083169116105b6136465760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610bd4565b5f805461ffff191660ff8316176101001781558290036136a85760405162461bcd60e51b815260206004820152601b60248201527f696e76616c696420726f6c6c75702064656c617920706572696f6400000000006044820152606401610bd4565b60ac829055604080515f8152602081018490527f624e47dc9fb8f8cfeaf4ead4710277cc1757136cfa885e465514cf6d510f0ad1910160405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611c23565b335f90815260a56020526040812054908190036137825760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642062617463684368616c6c656e6765526577617264000000006044820152606401610bd4565b335f90815260a5602052604081205561379b82826149f9565b816001600160a01b03167f9c25fa83f414ed363c8d39c98fb3e17567b3431cede71eb062c49d2a63ce247a826040516137d691815260200190565b60405180910390a25050565b5f81815260a460205260408120546801000000000000000090046001600160a01b03161580159061187a5750505f90815260a46020526040902060030154610100900460ff161590565b613834613c0d565b5f8111801561384557506098548114155b6138915760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206e65772066696e616c697a6520706572696f6400000000006044820152606401610bd4565b609880549082905560408051828152602081018490527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a4379101611c23565b6138d7613c0d565b6001600160a01b0381166139535760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610bd4565b612f7c8161465d565b5f54600490610100900460ff1615801561397c57505f5460ff8083169116105b6139ee5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610bd4565b5f805461ffff191660ff8316176101001790556001600160a01b03821615801590613a2257505f826001600160a01b03163b115b613a6e5760405162461bcd60e51b815260206004820152601160248201527f696e76616c6964207375626d69747465720000000000000000000000000000006044820152606401610bd4565b609d54609e5414613ac15760405162461bcd60e51b815260206004820152600f60248201527f70656e64696e67206261746368657300000000000000000000000000000000006044820152606401610bd4565b60a65460ff1615613b145760405162461bcd60e51b815260206004820152601560248201527f6368616c6c656e676520696e2070726f677265737300000000000000000000006044820152606401610bd4565b60a85415613b645760405162461bcd60e51b815260206004820152601660248201527f70656e64696e67207265766572742072657175657374000000000000000000006044820152606401610bd4565b6097546040516001600160a01b038085169216907fde3ada76ae1aa387cc4c586c70304b9fae54b251958786c828aaa0ce0290f939905f90a3609780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384161790555f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611c23565b6033546001600160a01b031633146122235760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610bd4565b5f805f613c748585614b35565b90505f8160ff165f03613c9557613c8b8686614ba5565b9094509050613d01565b8160ff1660011480613caa57508160ff166002145b15613cb957613c8b8686614c0e565b60405162461bcd60e51b815260206004820152601960248201527f556e737570706f727465642062617463682076657273696f6e000000000000006044820152606401610bd4565b808420925050509250929050565b60655460ff16156122235760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610bd4565b805f03613d6d575050565b8082035f5b82811015611b7e57610100818403811115613d8c57508083035b609b546040517f3c7f528300000000000000000000000000000000000000000000000000000000815260048101859052602481018390526001600160a01b0390911690633c7f5283906044015f604051808303815f87803b158015613def575f80fd5b505af1158015613e01573d5f803e3d5ffd5b50505050610100830192505061010081019050613d72565b5f81600203613ea8575f6040515f5b804980613e355750613e44565b60208202830152600101613e28565b602081028083209201604052909250905080613ea25760405162461bcd60e51b815260206004820152601b60248201527f5632207265717569726573206174206c65617374203120626c6f6200000000006044820152606401610bd4565b50919050565b60014915613f1e5760405162461bcd60e51b815260206004820152602560248201527f6c6567616379206261746368657320737570706f72742065786163746c79203160448201527f20626c6f620000000000000000000000000000000000000000000000000000006064820152608401610bd4565b5f4915613f2c575f4961187a565b507f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440145b919050565b6002613f6360208501856152d3565b60ff161115613fb45760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642076657273696f6e00000000000000000000000000000000006044820152606401610bd4565b60808301356140055760405162461bcd60e51b815260206004820152601b60248201527f70726576696f757320737461746520726f6f74206973207a65726f00000000006044820152606401610bd4565b60a08301356140565760405162461bcd60e51b815260206004820152601660248201527f6e657720737461746520726f6f74206973207a65726f000000000000000000006044820152606401610bd4565b5f8061406861166e6020870187615272565b915091505f61407b836001015160c01c90565b90505f60a18161408c846001615214565b81526020019081526020015f2054146140e75760405162461bcd60e51b815260206004820152601760248201527f626174636820616c726561647920636f6d6d69747465640000000000000000006044820152606401610bd4565b609e5481146141385760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610bd4565b5f81815260a1602052604090205482146141945760405162461bcd60e51b815260206004820152601b60248201527f696e636f727265637420706172656e74206261746368206861736800000000006044820152606401610bd4565b5f81815260ab60205260409020546080870135146141f45760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610bd4565b5f614203846011015160c01c90565b90505f61422f61421960608a0160408b01615345565b61422960808b0160608c01615251565b84614c61565b90506142416080890160608a01615251565b60019384019361ffff91909116929092019160f99061426360208b018b6152d3565b60ff161061427057506101015b6040805182810190915295506142958661428d60208c018c6152d3565b60ff16614ca5565b60c084901b60018701526142c3866142b360808c0160608d01615251565b61ffff1660c01b60099190910152565b60c083811b6011880152601987018390526039870188905260808a0135605988015260a08a0135607988015289013560998701525f60b987015260d98601859052600161431360208b018b6152d3565b60ff1610614345576143458661432f60608c0160408d01615345565b67ffffffffffffffff1660c01b60f99190910152565b8086205f85815260a1602090815260408083209390935560ab815282822060a08d0135905560ad905290812088905560a65460ff16156143ae5760a7545f90815260a4602052604090206002015460995442916143a191615214565b6143ab9190615227565b90505b604051806080016040528042815260200182609854426143ce9190615214565b6143d89190615214565b81526020016143ed60608d0160408e01615345565b67ffffffffffffffff1681526001600160a01b038b81166020928301525f88815260a28352604080822085518155858501516001820155858201516002820155606090950151600390950180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169590931694909417909155609e88905560a19091528181205491519193508692507f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f91a35050505050505050565b806144f75760405162461bcd60e51b815260206004820152601360248201527f496e76616c69642062617463682070726f6f66000000000000000000000000006044820152606401610bd4565b5f614506846001015160c01c90565b90505f614514856039015190565b90505f7f0000000000000000000000000000000000000000000000000000000000000000614543876059015190565b6079880151609989015160b98a015160198b015160405160c09690961b7fffffffffffffffff000000000000000000000000000000000000000000000000166020870152602886019490945260488501929092526068840152608883015260a882015260c8810183905260e801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120609c549091506001600160a01b0316632c09a848614605885160f81c90565b858888866040518663ffffffff1660e01b815260040161462995949392919061535e565b5f6040518083038186803b15801561463f575f80fd5b505afa158015614651573d5f803e3d5ffd5b50505050505050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff166147425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bd4565b612223614cac565b5f54610100900460ff166147c65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bd4565b612223614d34565b60a88390555f83815260a460205260408082205460975491517fc96be4cb0000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301526801000000000000000090920482169392919091169063c96be4cb906024016020604051808303815f875af1158015614855573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614879919061523a565b5f86815260a46020526040902060010154909150614898908290615214565b5f86815260a460209081526040808320546801000000000000000090046001600160a01b0316835260a5909152812080549091906148d7908490615214565b90915550506040516148ea9084906153c1565b604051908190038120906001600160a01b0384169087907fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a905f90a45050505050565b5f83815260a4602052604081206001015460a95490919060649061495190846153ed565b61495b9190615404565b90506149678183615227565b60aa5f8282546149779190615214565b90915550506001600160a01b0384165f90815260a56020526040812080548392906149a3908490615214565b90915550506040516149b69084906153c1565b604051908190038120906001600160a01b0386169087907fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a905f90a45050505050565b8015614a9e575f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114614a48576040519150601f19603f3d011682016040523d82523d5f602084013e614a4d565b606091505b50509050806121505760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610bd4565b5050565b614aaa613d0f565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258614adf3390565b6040516001600160a01b03909116815260200160405180910390a1565b614b04614db9565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33614adf565b5f81614b835760405162461bcd60e51b815260206004820152601260248201527f456d7074792062617463682068656164657200000000000000000000000000006044820152606401610bd4565b82825f818110614b9557614b9561543c565b919091013560f81c949350505050565b5f8160f9811015614bf85760405162461bcd60e51b815260206004820152601d60248201527f626174636820686561646572206c656e67746820746f6f20736d616c6c0000006044820152606401610bd4565b6040519150808483378082016040529250929050565b5f816101018114614bf85760405162461bcd60e51b815260206004820181905260248201527f626174636820686561646572206c656e67746820697320696e636f72726563746044820152606401610bd4565b6040805160c085901b815260f084901b6008820152600a60208502820181019092525f918101614c968161ffff871686614e0b565b82900390912095945050505050565b8082535050565b5f54610100900460ff16614d285760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bd4565b6065805460ff19169055565b5f54610100900460ff16614db05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610bd4565b6122233361465d565b60655460ff166122235760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610bd4565b5f825f03614e1a575082614ecd565b609b546001600160a01b03165f5b84811015614ec7576040517fae453cd5000000000000000000000000000000000000000000000000000000008152600481018590525f906001600160a01b0384169063ae453cd590602401602060405180830381865afa158015614e8e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614eb2919061523a565b87525060209095019460019384019301614e28565b50849150505b9392505050565b5f60208284031215614ee4575f80fd5b5035919050565b80356001600160a01b0381168114613f4f575f80fd5b5f60208284031215614f11575f80fd5b614ecd82614eeb565b5f8083601f840112614f2a575f80fd5b50813567ffffffffffffffff811115614f41575f80fd5b602083019150836020828501011115614f58575f80fd5b9250929050565b5f805f60408486031215614f71575f80fd5b833567ffffffffffffffff811115614f87575f80fd5b614f9386828701614f1a565b909790965060209590950135949350505050565b5f8060208385031215614fb8575f80fd5b823567ffffffffffffffff811115614fce575f80fd5b614fda85828601614f1a565b90969095509350505050565b5f60e08284031215613ea2575f80fd5b5f805f805f6060868803121561500a575f80fd5b853567ffffffffffffffff80821115615021575f80fd5b61502d89838a01614fe6565b96506020880135915080821115615042575f80fd5b61504e89838a01614f1a565b90965094506040880135915080821115615066575f80fd5b5061507388828901614f1a565b969995985093965092949392505050565b5f60208284031215615094575f80fd5b813567ffffffffffffffff8111156150aa575f80fd5b6150b684828501614fe6565b949350505050565b5f805f805f8060c087890312156150d3575f80fd5b6150dc87614eeb565b95506150ea60208801614eeb565b94506150f860408801614eeb565b9350606087013592506080870135915060a087013590509295509295509295565b5f805f806040858703121561512c575f80fd5b843567ffffffffffffffff80821115615143575f80fd5b61514f88838901614f1a565b90965094506020870135915080821115615167575f80fd5b5061517487828801614f1a565b95989497509550505050565b8015158114612f7c575f80fd5b5f6020828403121561519d575f80fd5b8135614ecd81615180565b803567ffffffffffffffff81168114613f4f575f80fd5b5f80604083850312156151d0575f80fd5b6151d9836151a8565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561187a5761187a6151e7565b8181038181111561187a5761187a6151e7565b5f6020828403121561524a575f80fd5b5051919050565b5f60208284031215615261575f80fd5b813561ffff81168114614ecd575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126152a5575f80fd5b83018035915067ffffffffffffffff8211156152bf575f80fd5b602001915036819003821315614f58575f80fd5b5f602082840312156152e3575f80fd5b813560ff81168114614ecd575f80fd5b5f60208284031215615303575f80fd5b8151614ecd81615180565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361533e5761533e6151e7565b5060010190565b5f60208284031215615355575f80fd5b614ecd826151a8565b85815284602082015260806040820152826080820152828460a08301375f60a084830101525f60a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f86011683010190508260608301529695505050505050565b5f82515f5b818110156153e057602081860181015185830152016153c6565b505f920191825250919050565b808202811582820484141761187a5761187a6151e7565b5f82615437577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea164736f6c6343000818000a",
}

// RollupABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMetaData.ABI instead.
var RollupABI = RollupMetaData.ABI

// RollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMetaData.Bin instead.
var RollupBin = RollupMetaData.Bin

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint64) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupBin), backend, _chainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// LAYER2CHAINID is a free data retrieval call binding the contract method 0x5f77cf1d.
//
// Solidity: function LAYER_2_CHAIN_ID() view returns(uint64)
func (_Rollup *RollupCaller) LAYER2CHAINID(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "LAYER_2_CHAIN_ID")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LAYER2CHAINID is a free data retrieval call binding the contract method 0x5f77cf1d.
//
// Solidity: function LAYER_2_CHAIN_ID() view returns(uint64)
func (_Rollup *RollupSession) LAYER2CHAINID() (uint64, error) {
	return _Rollup.Contract.LAYER2CHAINID(&_Rollup.CallOpts)
}

// LAYER2CHAINID is a free data retrieval call binding the contract method 0x5f77cf1d.
//
// Solidity: function LAYER_2_CHAIN_ID() view returns(uint64)
func (_Rollup *RollupCallerSession) LAYER2CHAINID() (uint64, error) {
	return _Rollup.Contract.LAYER2CHAINID(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xdff7827e.
//
// Solidity: function __maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCaller) MaxNumTxInChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "__maxNumTxInChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xdff7827e.
//
// Solidity: function __maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xdff7827e.
//
// Solidity: function __maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCallerSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
}

// BatchBlobVersionedHashes is a free data retrieval call binding the contract method 0x2a213ba1.
//
// Solidity: function batchBlobVersionedHashes(uint256 batchIndex) view returns(bytes32 blobVersionedHash)
func (_Rollup *RollupCaller) BatchBlobVersionedHashes(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchBlobVersionedHashes", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchBlobVersionedHashes is a free data retrieval call binding the contract method 0x2a213ba1.
//
// Solidity: function batchBlobVersionedHashes(uint256 batchIndex) view returns(bytes32 blobVersionedHash)
func (_Rollup *RollupSession) BatchBlobVersionedHashes(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.BatchBlobVersionedHashes(&_Rollup.CallOpts, batchIndex)
}

// BatchBlobVersionedHashes is a free data retrieval call binding the contract method 0x2a213ba1.
//
// Solidity: function batchBlobVersionedHashes(uint256 batchIndex) view returns(bytes32 blobVersionedHash)
func (_Rollup *RollupCallerSession) BatchBlobVersionedHashes(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.BatchBlobVersionedHashes(&_Rollup.CallOpts, batchIndex)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address owner) view returns(uint256 amount)
func (_Rollup *RollupCaller) BatchChallengeReward(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallengeReward", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address owner) view returns(uint256 amount)
func (_Rollup *RollupSession) BatchChallengeReward(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, owner)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address owner) view returns(uint256 amount)
func (_Rollup *RollupCallerSession) BatchChallengeReward(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, owner)
}

// BatchChallenged is a free data retrieval call binding the contract method 0x18463fb0.
//
// Solidity: function batchChallenged() view returns(uint256)
func (_Rollup *RollupCaller) BatchChallenged(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallenged")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchChallenged is a free data retrieval call binding the contract method 0x18463fb0.
//
// Solidity: function batchChallenged() view returns(uint256)
func (_Rollup *RollupSession) BatchChallenged() (*big.Int, error) {
	return _Rollup.Contract.BatchChallenged(&_Rollup.CallOpts)
}

// BatchChallenged is a free data retrieval call binding the contract method 0x18463fb0.
//
// Solidity: function batchChallenged() view returns(uint256)
func (_Rollup *RollupCallerSession) BatchChallenged() (*big.Int, error) {
	return _Rollup.Contract.BatchChallenged(&_Rollup.CallOpts)
}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchChallengedSuccess(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallengedSuccess", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchChallengedSuccess(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchChallengedSuccess(&_Rollup.CallOpts, batchIndex)
}

// BatchChallengedSuccess is a free data retrieval call binding the contract method 0x910129d4.
//
// Solidity: function batchChallengedSuccess(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchChallengedSuccess(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchChallengedSuccess(&_Rollup.CallOpts, batchIndex)
}

// BatchDataStore is a free data retrieval call binding the contract method 0x5ef7a94a.
//
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, address submitter)
func (_Rollup *RollupCaller) BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchDataStore", batchIndex)

	outstruct := new(struct {
		OriginTimestamp   *big.Int
		FinalizeTimestamp *big.Int
		BlockNumber       *big.Int
		Submitter         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FinalizeTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Submitter = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BatchDataStore is a free data retrieval call binding the contract method 0x5ef7a94a.
//
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, address submitter)
func (_Rollup *RollupSession) BatchDataStore(batchIndex *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	return _Rollup.Contract.BatchDataStore(&_Rollup.CallOpts, batchIndex)
}

// BatchDataStore is a free data retrieval call binding the contract method 0x5ef7a94a.
//
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, address submitter)
func (_Rollup *RollupCallerSession) BatchDataStore(batchIndex *big.Int) (struct {
	OriginTimestamp   *big.Int
	FinalizeTimestamp *big.Int
	BlockNumber       *big.Int
	Submitter         common.Address
}, error) {
	return _Rollup.Contract.BatchDataStore(&_Rollup.CallOpts, batchIndex)
}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchExist(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchExist", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchExist(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchExist(&_Rollup.CallOpts, batchIndex)
}

// BatchExist is a free data retrieval call binding the contract method 0x21e2f9e0.
//
// Solidity: function batchExist(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchExist(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchExist(&_Rollup.CallOpts, batchIndex)
}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchInChallenge(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchInChallenge", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchInChallenge(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInChallenge(&_Rollup.CallOpts, batchIndex)
}

// BatchInChallenge is a free data retrieval call binding the contract method 0xde8b3035.
//
// Solidity: function batchInChallenge(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchInChallenge(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInChallenge(&_Rollup.CallOpts, batchIndex)
}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCaller) BatchInsideChallengeWindow(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchInsideChallengeWindow", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupSession) BatchInsideChallengeWindow(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInsideChallengeWindow(&_Rollup.CallOpts, batchIndex)
}

// BatchInsideChallengeWindow is a free data retrieval call binding the contract method 0x18af3b2b.
//
// Solidity: function batchInsideChallengeWindow(uint256 batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) BatchInsideChallengeWindow(batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.BatchInsideChallengeWindow(&_Rollup.CallOpts, batchIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 batchIndex) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCaller) Challenges(opts *bind.CallOpts, batchIndex *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challenges", batchIndex)

	outstruct := new(struct {
		BatchIndex       uint64
		Challenger       common.Address
		ChallengeDeposit *big.Int
		StartTime        *big.Int
		ChallengeSuccess bool
		Finished         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchIndex = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Challenger = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ChallengeDeposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ChallengeSuccess = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Finished = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 batchIndex) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupSession) Challenges(batchIndex *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, batchIndex)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 batchIndex) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCallerSession) Challenges(batchIndex *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32 batchHash)
func (_Rollup *RollupCaller) CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "committedBatches", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32 batchHash)
func (_Rollup *RollupSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32 batchHash)
func (_Rollup *RollupCallerSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// CommittedStateRoots is a free data retrieval call binding the contract method 0x68589dfa.
//
// Solidity: function committedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupCaller) CommittedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "committedStateRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedStateRoots is a free data retrieval call binding the contract method 0x68589dfa.
//
// Solidity: function committedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupSession) CommittedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedStateRoots(&_Rollup.CallOpts, batchIndex)
}

// CommittedStateRoots is a free data retrieval call binding the contract method 0x68589dfa.
//
// Solidity: function committedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupCallerSession) CommittedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedStateRoots(&_Rollup.CallOpts, batchIndex)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_Rollup *RollupCaller) FinalizationPeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "finalizationPeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_Rollup *RollupSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _Rollup.Contract.FinalizationPeriodSeconds(&_Rollup.CallOpts)
}

// FinalizationPeriodSeconds is a free data retrieval call binding the contract method 0xce5db8d6.
//
// Solidity: function finalizationPeriodSeconds() view returns(uint256)
func (_Rollup *RollupCallerSession) FinalizationPeriodSeconds() (*big.Int, error) {
	return _Rollup.Contract.FinalizationPeriodSeconds(&_Rollup.CallOpts)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupCaller) FinalizedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "finalizedStateRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32 stateRoot)
func (_Rollup *RollupCallerSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, batchIndex)
}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupCaller) InChallenge(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "inChallenge")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupSession) InChallenge() (bool, error) {
	return _Rollup.Contract.InChallenge(&_Rollup.CallOpts)
}

// InChallenge is a free data retrieval call binding the contract method 0x88b1ea09.
//
// Solidity: function inChallenge() view returns(bool)
func (_Rollup *RollupCallerSession) InChallenge() (bool, error) {
	return _Rollup.Contract.InChallenge(&_Rollup.CallOpts)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupCaller) IsBatchFinalized(opts *bind.CallOpts, _batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isBatchFinalized", _batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.IsBatchFinalized(&_Rollup.CallOpts, _batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 _batchIndex) view returns(bool)
func (_Rollup *RollupCallerSession) IsBatchFinalized(_batchIndex *big.Int) (bool, error) {
	return _Rollup.Contract.IsBatchFinalized(&_Rollup.CallOpts, _batchIndex)
}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address challengerAddress) view returns(bool isChallenger)
func (_Rollup *RollupCaller) IsChallenger(opts *bind.CallOpts, challengerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isChallenger", challengerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address challengerAddress) view returns(bool isChallenger)
func (_Rollup *RollupSession) IsChallenger(challengerAddress common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, challengerAddress)
}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address challengerAddress) view returns(bool isChallenger)
func (_Rollup *RollupCallerSession) IsChallenger(challengerAddress common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, challengerAddress)
}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupCaller) LastCommittedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastCommittedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupSession) LastCommittedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastCommittedBatchIndex(&_Rollup.CallOpts)
}

// LastCommittedBatchIndex is a free data retrieval call binding the contract method 0x121dcd50.
//
// Solidity: function lastCommittedBatchIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) LastCommittedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastCommittedBatchIndex(&_Rollup.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupCaller) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "lastFinalizedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastFinalizedBatchIndex(&_Rollup.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _Rollup.Contract.LastFinalizedBatchIndex(&_Rollup.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupSession) MessageQueue() (common.Address, error) {
	return _Rollup.Contract.MessageQueue(&_Rollup.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_Rollup *RollupCallerSession) MessageQueue() (common.Address, error) {
	return _Rollup.Contract.MessageQueue(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Rollup *RollupCallerSession) Paused() (bool, error) {
	return _Rollup.Contract.Paused(&_Rollup.CallOpts)
}

// ProofRewardPercent is a free data retrieval call binding the contract method 0xfb1e8b04.
//
// Solidity: function proofRewardPercent() view returns(uint256)
func (_Rollup *RollupCaller) ProofRewardPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "proofRewardPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofRewardPercent is a free data retrieval call binding the contract method 0xfb1e8b04.
//
// Solidity: function proofRewardPercent() view returns(uint256)
func (_Rollup *RollupSession) ProofRewardPercent() (*big.Int, error) {
	return _Rollup.Contract.ProofRewardPercent(&_Rollup.CallOpts)
}

// ProofRewardPercent is a free data retrieval call binding the contract method 0xfb1e8b04.
//
// Solidity: function proofRewardPercent() view returns(uint256)
func (_Rollup *RollupCallerSession) ProofRewardPercent() (*big.Int, error) {
	return _Rollup.Contract.ProofRewardPercent(&_Rollup.CallOpts)
}

// ProofWindow is a free data retrieval call binding the contract method 0xa479265d.
//
// Solidity: function proofWindow() view returns(uint256)
func (_Rollup *RollupCaller) ProofWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "proofWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofWindow is a free data retrieval call binding the contract method 0xa479265d.
//
// Solidity: function proofWindow() view returns(uint256)
func (_Rollup *RollupSession) ProofWindow() (*big.Int, error) {
	return _Rollup.Contract.ProofWindow(&_Rollup.CallOpts)
}

// ProofWindow is a free data retrieval call binding the contract method 0xa479265d.
//
// Solidity: function proofWindow() view returns(uint256)
func (_Rollup *RollupCallerSession) ProofWindow() (*big.Int, error) {
	return _Rollup.Contract.ProofWindow(&_Rollup.CallOpts)
}

// ProveCommittedBatchState is a free data retrieval call binding the contract method 0xb8d0a1b0.
//
// Solidity: function proveCommittedBatchState(bytes _batchHeader, bytes _batchProof) view returns()
func (_Rollup *RollupCaller) ProveCommittedBatchState(opts *bind.CallOpts, _batchHeader []byte, _batchProof []byte) error {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "proveCommittedBatchState", _batchHeader, _batchProof)

	if err != nil {
		return err
	}

	return err

}

// ProveCommittedBatchState is a free data retrieval call binding the contract method 0xb8d0a1b0.
//
// Solidity: function proveCommittedBatchState(bytes _batchHeader, bytes _batchProof) view returns()
func (_Rollup *RollupSession) ProveCommittedBatchState(_batchHeader []byte, _batchProof []byte) error {
	return _Rollup.Contract.ProveCommittedBatchState(&_Rollup.CallOpts, _batchHeader, _batchProof)
}

// ProveCommittedBatchState is a free data retrieval call binding the contract method 0xb8d0a1b0.
//
// Solidity: function proveCommittedBatchState(bytes _batchHeader, bytes _batchProof) view returns()
func (_Rollup *RollupCallerSession) ProveCommittedBatchState(_batchHeader []byte, _batchProof []byte) error {
	return _Rollup.Contract.ProveCommittedBatchState(&_Rollup.CallOpts, _batchHeader, _batchProof)
}

// ProveRemaining is a free data retrieval call binding the contract method 0x3e001b66.
//
// Solidity: function proveRemaining() view returns(uint256)
func (_Rollup *RollupCaller) ProveRemaining(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "proveRemaining")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProveRemaining is a free data retrieval call binding the contract method 0x3e001b66.
//
// Solidity: function proveRemaining() view returns(uint256)
func (_Rollup *RollupSession) ProveRemaining() (*big.Int, error) {
	return _Rollup.Contract.ProveRemaining(&_Rollup.CallOpts)
}

// ProveRemaining is a free data retrieval call binding the contract method 0x3e001b66.
//
// Solidity: function proveRemaining() view returns(uint256)
func (_Rollup *RollupCallerSession) ProveRemaining() (*big.Int, error) {
	return _Rollup.Contract.ProveRemaining(&_Rollup.CallOpts)
}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupCaller) RevertReqIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "revertReqIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupSession) RevertReqIndex() (*big.Int, error) {
	return _Rollup.Contract.RevertReqIndex(&_Rollup.CallOpts)
}

// RevertReqIndex is a free data retrieval call binding the contract method 0xb31a77d3.
//
// Solidity: function revertReqIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) RevertReqIndex() (*big.Int, error) {
	return _Rollup.Contract.RevertReqIndex(&_Rollup.CallOpts)
}

// RollupDelayPeriod is a free data retrieval call binding the contract method 0xd8dc99d2.
//
// Solidity: function rollupDelayPeriod() view returns(uint256)
func (_Rollup *RollupCaller) RollupDelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "rollupDelayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupDelayPeriod is a free data retrieval call binding the contract method 0xd8dc99d2.
//
// Solidity: function rollupDelayPeriod() view returns(uint256)
func (_Rollup *RollupSession) RollupDelayPeriod() (*big.Int, error) {
	return _Rollup.Contract.RollupDelayPeriod(&_Rollup.CallOpts)
}

// RollupDelayPeriod is a free data retrieval call binding the contract method 0xd8dc99d2.
//
// Solidity: function rollupDelayPeriod() view returns(uint256)
func (_Rollup *RollupCallerSession) RollupDelayPeriod() (*big.Int, error) {
	return _Rollup.Contract.RollupDelayPeriod(&_Rollup.CallOpts)
}

// SubmitterContract is a free data retrieval call binding the contract method 0xa3c9e94c.
//
// Solidity: function submitterContract() view returns(address)
func (_Rollup *RollupCaller) SubmitterContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "submitterContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubmitterContract is a free data retrieval call binding the contract method 0xa3c9e94c.
//
// Solidity: function submitterContract() view returns(address)
func (_Rollup *RollupSession) SubmitterContract() (common.Address, error) {
	return _Rollup.Contract.SubmitterContract(&_Rollup.CallOpts)
}

// SubmitterContract is a free data retrieval call binding the contract method 0xa3c9e94c.
//
// Solidity: function submitterContract() view returns(address)
func (_Rollup *RollupCallerSession) SubmitterContract() (common.Address, error) {
	return _Rollup.Contract.SubmitterContract(&_Rollup.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Rollup *RollupCallerSession) Verifier() (common.Address, error) {
	return _Rollup.Contract.Verifier(&_Rollup.CallOpts)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 withdrawalRoot) view returns(bool exist)
func (_Rollup *RollupCaller) WithdrawalRoots(opts *bind.CallOpts, withdrawalRoot [32]byte) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawalRoots", withdrawalRoot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 withdrawalRoot) view returns(bool exist)
func (_Rollup *RollupSession) WithdrawalRoots(withdrawalRoot [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, withdrawalRoot)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 withdrawalRoot) view returns(bool exist)
func (_Rollup *RollupCallerSession) WithdrawalRoots(withdrawalRoot [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, withdrawalRoot)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupTransactor) AddChallenger(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addChallenger", _account)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupSession) AddChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddChallenger(&_Rollup.TransactOpts, _account)
}

// AddChallenger is a paid mutator transaction binding the contract method 0x0ceb6780.
//
// Solidity: function addChallenger(address _account) returns()
func (_Rollup *RollupTransactorSession) AddChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddChallenger(&_Rollup.TransactOpts, _account)
}

// ChallengeState is a paid mutator transaction binding the contract method 0xcd4edc69.
//
// Solidity: function challengeState(uint64 batchIndex, bytes32 _batchHash) payable returns()
func (_Rollup *RollupTransactor) ChallengeState(opts *bind.TransactOpts, batchIndex uint64, _batchHash [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "challengeState", batchIndex, _batchHash)
}

// ChallengeState is a paid mutator transaction binding the contract method 0xcd4edc69.
//
// Solidity: function challengeState(uint64 batchIndex, bytes32 _batchHash) payable returns()
func (_Rollup *RollupSession) ChallengeState(batchIndex uint64, _batchHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex, _batchHash)
}

// ChallengeState is a paid mutator transaction binding the contract method 0xcd4edc69.
//
// Solidity: function challengeState(uint64 batchIndex, bytes32 _batchHash) payable returns()
func (_Rollup *RollupTransactorSession) ChallengeState(batchIndex uint64, _batchHash [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex, _batchHash)
}

// ClaimProveRemaining is a paid mutator transaction binding the contract method 0xb35dac4e.
//
// Solidity: function claimProveRemaining(address receiver) returns()
func (_Rollup *RollupTransactor) ClaimProveRemaining(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "claimProveRemaining", receiver)
}

// ClaimProveRemaining is a paid mutator transaction binding the contract method 0xb35dac4e.
//
// Solidity: function claimProveRemaining(address receiver) returns()
func (_Rollup *RollupSession) ClaimProveRemaining(receiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ClaimProveRemaining(&_Rollup.TransactOpts, receiver)
}

// ClaimProveRemaining is a paid mutator transaction binding the contract method 0xb35dac4e.
//
// Solidity: function claimProveRemaining(address receiver) returns()
func (_Rollup *RollupTransactorSession) ClaimProveRemaining(receiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ClaimProveRemaining(&_Rollup.TransactOpts, receiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns()
func (_Rollup *RollupTransactor) ClaimReward(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "claimReward", receiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns()
func (_Rollup *RollupSession) ClaimReward(receiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ClaimReward(&_Rollup.TransactOpts, receiver)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xd279c191.
//
// Solidity: function claimReward(address receiver) returns()
func (_Rollup *RollupTransactorSession) ClaimReward(receiver common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.ClaimReward(&_Rollup.TransactOpts, receiver)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x41f756da.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) payable returns()
func (_Rollup *RollupTransactor) CommitBatch(opts *bind.TransactOpts, batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatch", batchDataInput)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x41f756da.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) payable returns()
func (_Rollup *RollupSession) CommitBatch(batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchDataInput)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x41f756da.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) payable returns()
func (_Rollup *RollupTransactorSession) CommitBatch(batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchDataInput)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x1544ba3a.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupTransactor) CommitBatchWithProof(opts *bind.TransactOpts, batchDataInput IRollupBatchDataInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatchWithProof", batchDataInput, _batchHeader, _batchProof)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x1544ba3a.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupSession) CommitBatchWithProof(batchDataInput IRollupBatchDataInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatchWithProof(&_Rollup.TransactOpts, batchDataInput, _batchHeader, _batchProof)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x1544ba3a.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupTransactorSession) CommitBatchWithProof(batchDataInput IRollupBatchDataInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatchWithProof(&_Rollup.TransactOpts, batchDataInput, _batchHeader, _batchProof)
}

// CommitState is a paid mutator transaction binding the contract method 0x67caa37a.
//
// Solidity: function commitState((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) returns()
func (_Rollup *RollupTransactor) CommitState(opts *bind.TransactOpts, batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitState", batchDataInput)
}

// CommitState is a paid mutator transaction binding the contract method 0x67caa37a.
//
// Solidity: function commitState((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) returns()
func (_Rollup *RollupSession) CommitState(batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitState(&_Rollup.TransactOpts, batchDataInput)
}

// CommitState is a paid mutator transaction binding the contract method 0x67caa37a.
//
// Solidity: function commitState((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput) returns()
func (_Rollup *RollupTransactorSession) CommitState(batchDataInput IRollupBatchDataInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitState(&_Rollup.TransactOpts, batchDataInput)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x13361101.
//
// Solidity: function finalizeBatch(bytes _batchHeader) returns()
func (_Rollup *RollupTransactor) FinalizeBatch(opts *bind.TransactOpts, _batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finalizeBatch", _batchHeader)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x13361101.
//
// Solidity: function finalizeBatch(bytes _batchHeader) returns()
func (_Rollup *RollupSession) FinalizeBatch(_batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchHeader)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0x13361101.
//
// Solidity: function finalizeBatch(bytes _batchHeader) returns()
func (_Rollup *RollupTransactorSession) FinalizeBatch(_batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchHeader)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0xb3e0a509.
//
// Solidity: function importGenesisBatch(bytes _batchHeader) returns()
func (_Rollup *RollupTransactor) ImportGenesisBatch(opts *bind.TransactOpts, _batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "importGenesisBatch", _batchHeader)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0xb3e0a509.
//
// Solidity: function importGenesisBatch(bytes _batchHeader) returns()
func (_Rollup *RollupSession) ImportGenesisBatch(_batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0xb3e0a509.
//
// Solidity: function importGenesisBatch(bytes _batchHeader) returns()
func (_Rollup *RollupTransactorSession) ImportGenesisBatch(_batchHeader []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _submitterContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _submitterContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _submitterContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _submitterContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupSession) Initialize(_submitterContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _submitterContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _submitterContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupTransactorSession) Initialize(_submitterContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _submitterContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
}

// Initialize2 is a paid mutator transaction binding the contract method 0x61267290.
//
// Solidity: function initialize2(bytes32 _prevStateRoot) returns()
func (_Rollup *RollupTransactor) Initialize2(opts *bind.TransactOpts, _prevStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize2", _prevStateRoot)
}

// Initialize2 is a paid mutator transaction binding the contract method 0x61267290.
//
// Solidity: function initialize2(bytes32 _prevStateRoot) returns()
func (_Rollup *RollupSession) Initialize2(_prevStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize2(&_Rollup.TransactOpts, _prevStateRoot)
}

// Initialize2 is a paid mutator transaction binding the contract method 0x61267290.
//
// Solidity: function initialize2(bytes32 _prevStateRoot) returns()
func (_Rollup *RollupTransactorSession) Initialize2(_prevStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize2(&_Rollup.TransactOpts, _prevStateRoot)
}

// Initialize3 is a paid mutator transaction binding the contract method 0xcf9a6745.
//
// Solidity: function initialize3(uint256 _rollupDelayPeriod) returns()
func (_Rollup *RollupTransactor) Initialize3(opts *bind.TransactOpts, _rollupDelayPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize3", _rollupDelayPeriod)
}

// Initialize3 is a paid mutator transaction binding the contract method 0xcf9a6745.
//
// Solidity: function initialize3(uint256 _rollupDelayPeriod) returns()
func (_Rollup *RollupSession) Initialize3(_rollupDelayPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize3(&_Rollup.TransactOpts, _rollupDelayPeriod)
}

// Initialize3 is a paid mutator transaction binding the contract method 0xcf9a6745.
//
// Solidity: function initialize3(uint256 _rollupDelayPeriod) returns()
func (_Rollup *RollupTransactorSession) Initialize3(_rollupDelayPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize3(&_Rollup.TransactOpts, _rollupDelayPeriod)
}

// Initialize4 is a paid mutator transaction binding the contract method 0xf8fa010f.
//
// Solidity: function initialize4(address _newSubmitter) returns()
func (_Rollup *RollupTransactor) Initialize4(opts *bind.TransactOpts, _newSubmitter common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize4", _newSubmitter)
}

// Initialize4 is a paid mutator transaction binding the contract method 0xf8fa010f.
//
// Solidity: function initialize4(address _newSubmitter) returns()
func (_Rollup *RollupSession) Initialize4(_newSubmitter common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize4(&_Rollup.TransactOpts, _newSubmitter)
}

// Initialize4 is a paid mutator transaction binding the contract method 0xf8fa010f.
//
// Solidity: function initialize4(address _newSubmitter) returns()
func (_Rollup *RollupTransactorSession) Initialize4(_newSubmitter common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize4(&_Rollup.TransactOpts, _newSubmitter)
}

// ProveState is a paid mutator transaction binding the contract method 0xb3484425.
//
// Solidity: function proveState(bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupTransactor) ProveState(opts *bind.TransactOpts, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "proveState", _batchHeader, _batchProof)
}

// ProveState is a paid mutator transaction binding the contract method 0xb3484425.
//
// Solidity: function proveState(bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupSession) ProveState(_batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchHeader, _batchProof)
}

// ProveState is a paid mutator transaction binding the contract method 0xb3484425.
//
// Solidity: function proveState(bytes _batchHeader, bytes _batchProof) returns()
func (_Rollup *RollupTransactorSession) ProveState(_batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchHeader, _batchProof)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupTransactor) RemoveChallenger(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeChallenger", _account)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupSession) RemoveChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveChallenger(&_Rollup.TransactOpts, _account)
}

// RemoveChallenger is a paid mutator transaction binding the contract method 0x6c578c1d.
//
// Solidity: function removeChallenger(address _account) returns()
func (_Rollup *RollupTransactorSession) RemoveChallenger(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveChallenger(&_Rollup.TransactOpts, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rollup.Contract.RenounceOwnership(&_Rollup.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rollup *RollupTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rollup.Contract.RenounceOwnership(&_Rollup.TransactOpts)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupTransactor) RevertBatch(opts *bind.TransactOpts, _batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "revertBatch", _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RevertBatch(&_Rollup.TransactOpts, _batchHeader, _count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes _batchHeader, uint256 _count) returns()
func (_Rollup *RollupTransactorSession) RevertBatch(_batchHeader []byte, _count *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.RevertBatch(&_Rollup.TransactOpts, _batchHeader, _count)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupSession) SetPause(_status bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetPause(&_Rollup.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_Rollup *RollupTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _Rollup.Contract.SetPause(&_Rollup.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.TransferOwnership(&_Rollup.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollup *RollupTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.TransferOwnership(&_Rollup.TransactOpts, newOwner)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactor) UpdateFinalizePeriodSeconds(opts *bind.TransactOpts, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateFinalizePeriodSeconds", _newPeriod)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupSession) UpdateFinalizePeriodSeconds(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateFinalizePeriodSeconds(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateFinalizePeriodSeconds is a paid mutator transaction binding the contract method 0xe3fff1dd.
//
// Solidity: function updateFinalizePeriodSeconds(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactorSession) UpdateFinalizePeriodSeconds(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateFinalizePeriodSeconds(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupTransactor) UpdateProofWindow(opts *bind.TransactOpts, _newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateProofWindow", _newWindow)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupSession) UpdateProofWindow(_newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateProofWindow(&_Rollup.TransactOpts, _newWindow)
}

// UpdateProofWindow is a paid mutator transaction binding the contract method 0x57e0af6c.
//
// Solidity: function updateProofWindow(uint256 _newWindow) returns()
func (_Rollup *RollupTransactorSession) UpdateProofWindow(_newWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateProofWindow(&_Rollup.TransactOpts, _newWindow)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 _newProofRewardPercent) returns()
func (_Rollup *RollupTransactor) UpdateRewardPercentage(opts *bind.TransactOpts, _newProofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateRewardPercentage", _newProofRewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 _newProofRewardPercent) returns()
func (_Rollup *RollupSession) UpdateRewardPercentage(_newProofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateRewardPercentage(&_Rollup.TransactOpts, _newProofRewardPercent)
}

// UpdateRewardPercentage is a paid mutator transaction binding the contract method 0xa4f209b0.
//
// Solidity: function updateRewardPercentage(uint256 _newProofRewardPercent) returns()
func (_Rollup *RollupTransactorSession) UpdateRewardPercentage(_newProofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateRewardPercentage(&_Rollup.TransactOpts, _newProofRewardPercent)
}

// UpdateRollupDelayPeriod is a paid mutator transaction binding the contract method 0xc5553892.
//
// Solidity: function updateRollupDelayPeriod(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactor) UpdateRollupDelayPeriod(opts *bind.TransactOpts, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateRollupDelayPeriod", _newPeriod)
}

// UpdateRollupDelayPeriod is a paid mutator transaction binding the contract method 0xc5553892.
//
// Solidity: function updateRollupDelayPeriod(uint256 _newPeriod) returns()
func (_Rollup *RollupSession) UpdateRollupDelayPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateRollupDelayPeriod(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateRollupDelayPeriod is a paid mutator transaction binding the contract method 0xc5553892.
//
// Solidity: function updateRollupDelayPeriod(uint256 _newPeriod) returns()
func (_Rollup *RollupTransactorSession) UpdateRollupDelayPeriod(_newPeriod *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateRollupDelayPeriod(&_Rollup.TransactOpts, _newPeriod)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupTransactor) UpdateVerifier(opts *bind.TransactOpts, _newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateVerifier", _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateVerifier(&_Rollup.TransactOpts, _newVerifier)
}

// UpdateVerifier is a paid mutator transaction binding the contract method 0x97fc007c.
//
// Solidity: function updateVerifier(address _newVerifier) returns()
func (_Rollup *RollupTransactorSession) UpdateVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateVerifier(&_Rollup.TransactOpts, _newVerifier)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupSession) Receive() (*types.Transaction, error) {
	return _Rollup.Contract.Receive(&_Rollup.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Rollup *RollupTransactorSession) Receive() (*types.Transaction, error) {
	return _Rollup.Contract.Receive(&_Rollup.TransactOpts)
}

// RollupChallengeResIterator is returned from FilterChallengeRes and is used to iterate over the raw logs and unpacked data for ChallengeRes events raised by the Rollup contract.
type RollupChallengeResIterator struct {
	Event *RollupChallengeRes // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupChallengeResIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupChallengeRes)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupChallengeRes)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupChallengeResIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupChallengeResIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupChallengeRes represents a ChallengeRes event raised by the Rollup contract.
type RollupChallengeRes struct {
	BatchIndex *big.Int
	Winner     common.Address
	Res        common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeRes is a free log retrieval operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) FilterChallengeRes(opts *bind.FilterOpts, batchIndex []*big.Int, winner []common.Address, res []string) (*RollupChallengeResIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var resRule []interface{}
	for _, resItem := range res {
		resRule = append(resRule, resItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ChallengeRes", batchIndexRule, winnerRule, resRule)
	if err != nil {
		return nil, err
	}
	return &RollupChallengeResIterator{contract: _Rollup.contract, event: "ChallengeRes", logs: logs, sub: sub}, nil
}

// WatchChallengeRes is a free log subscription operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) WatchChallengeRes(opts *bind.WatchOpts, sink chan<- *RollupChallengeRes, batchIndex []*big.Int, winner []common.Address, res []string) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var resRule []interface{}
	for _, resItem := range res {
		resRule = append(resRule, resItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ChallengeRes", batchIndexRule, winnerRule, resRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupChallengeRes)
				if err := _Rollup.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeRes is a log parse operation binding the contract event 0xe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a.
//
// Solidity: event ChallengeRes(uint256 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) ParseChallengeRes(log types.Log) (*RollupChallengeRes, error) {
	event := new(RollupChallengeRes)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupChallengeRewardClaimIterator is returned from FilterChallengeRewardClaim and is used to iterate over the raw logs and unpacked data for ChallengeRewardClaim events raised by the Rollup contract.
type RollupChallengeRewardClaimIterator struct {
	Event *RollupChallengeRewardClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupChallengeRewardClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupChallengeRewardClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupChallengeRewardClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupChallengeRewardClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupChallengeRewardClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupChallengeRewardClaim represents a ChallengeRewardClaim event raised by the Rollup contract.
type RollupChallengeRewardClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChallengeRewardClaim is a free log retrieval operation binding the contract event 0x9c25fa83f414ed363c8d39c98fb3e17567b3431cede71eb062c49d2a63ce247a.
//
// Solidity: event ChallengeRewardClaim(address indexed receiver, uint256 amount)
func (_Rollup *RollupFilterer) FilterChallengeRewardClaim(opts *bind.FilterOpts, receiver []common.Address) (*RollupChallengeRewardClaimIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ChallengeRewardClaim", receiverRule)
	if err != nil {
		return nil, err
	}
	return &RollupChallengeRewardClaimIterator{contract: _Rollup.contract, event: "ChallengeRewardClaim", logs: logs, sub: sub}, nil
}

// WatchChallengeRewardClaim is a free log subscription operation binding the contract event 0x9c25fa83f414ed363c8d39c98fb3e17567b3431cede71eb062c49d2a63ce247a.
//
// Solidity: event ChallengeRewardClaim(address indexed receiver, uint256 amount)
func (_Rollup *RollupFilterer) WatchChallengeRewardClaim(opts *bind.WatchOpts, sink chan<- *RollupChallengeRewardClaim, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ChallengeRewardClaim", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupChallengeRewardClaim)
				if err := _Rollup.contract.UnpackLog(event, "ChallengeRewardClaim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeRewardClaim is a log parse operation binding the contract event 0x9c25fa83f414ed363c8d39c98fb3e17567b3431cede71eb062c49d2a63ce247a.
//
// Solidity: event ChallengeRewardClaim(address indexed receiver, uint256 amount)
func (_Rollup *RollupFilterer) ParseChallengeRewardClaim(log types.Log) (*RollupChallengeRewardClaim, error) {
	event := new(RollupChallengeRewardClaim)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeRewardClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupChallengeStateIterator is returned from FilterChallengeState and is used to iterate over the raw logs and unpacked data for ChallengeState events raised by the Rollup contract.
type RollupChallengeStateIterator struct {
	Event *RollupChallengeState // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupChallengeStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupChallengeState)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupChallengeState)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupChallengeStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupChallengeStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupChallengeState represents a ChallengeState event raised by the Rollup contract.
type RollupChallengeState struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterChallengeState is a free log retrieval operation binding the contract event 0x3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b9976.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address indexed challenger, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) FilterChallengeState(opts *bind.FilterOpts, batchIndex []uint64, challenger []common.Address) (*RollupChallengeStateIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ChallengeState", batchIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &RollupChallengeStateIterator{contract: _Rollup.contract, event: "ChallengeState", logs: logs, sub: sub}, nil
}

// WatchChallengeState is a free log subscription operation binding the contract event 0x3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b9976.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address indexed challenger, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) WatchChallengeState(opts *bind.WatchOpts, sink chan<- *RollupChallengeState, batchIndex []uint64, challenger []common.Address) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ChallengeState", batchIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupChallengeState)
				if err := _Rollup.contract.UnpackLog(event, "ChallengeState", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChallengeState is a log parse operation binding the contract event 0x3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b9976.
//
// Solidity: event ChallengeState(uint64 indexed batchIndex, address indexed challenger, uint256 challengeDeposit)
func (_Rollup *RollupFilterer) ParseChallengeState(log types.Log) (*RollupChallengeState, error) {
	event := new(RollupChallengeState)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCommitBatchIterator is returned from FilterCommitBatch and is used to iterate over the raw logs and unpacked data for CommitBatch events raised by the Rollup contract.
type RollupCommitBatchIterator struct {
	Event *RollupCommitBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupCommitBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCommitBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupCommitBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupCommitBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCommitBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCommitBatch represents a CommitBatch event raised by the Rollup contract.
type RollupCommitBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitBatch is a free log retrieval operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupCommitBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupCommitBatchIterator{contract: _Rollup.contract, event: "CommitBatch", logs: logs, sub: sub}, nil
}

// WatchCommitBatch is a free log subscription operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) WatchCommitBatch(opts *bind.WatchOpts, sink chan<- *RollupCommitBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCommitBatch)
				if err := _Rollup.contract.UnpackLog(event, "CommitBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitBatch is a log parse operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) ParseCommitBatch(log types.Log) (*RollupCommitBatch, error) {
	event := new(RollupCommitBatch)
	if err := _Rollup.contract.UnpackLog(event, "CommitBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupFinalizeBatchIterator is returned from FilterFinalizeBatch and is used to iterate over the raw logs and unpacked data for FinalizeBatch events raised by the Rollup contract.
type RollupFinalizeBatchIterator struct {
	Event *RollupFinalizeBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupFinalizeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupFinalizeBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupFinalizeBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupFinalizeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupFinalizeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupFinalizeBatch represents a FinalizeBatch event raised by the Rollup contract.
type RollupFinalizeBatch struct {
	BatchIndex   *big.Int
	BatchHash    [32]byte
	StateRoot    [32]byte
	WithdrawRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatch is a free log retrieval operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupFinalizeBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupFinalizeBatchIterator{contract: _Rollup.contract, event: "FinalizeBatch", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatch is a free log subscription operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) WatchFinalizeBatch(opts *bind.WatchOpts, sink chan<- *RollupFinalizeBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupFinalizeBatch)
				if err := _Rollup.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFinalizeBatch is a log parse operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_Rollup *RollupFilterer) ParseFinalizeBatch(log types.Log) (*RollupFinalizeBatch, error) {
	event := new(RollupFinalizeBatch)
	if err := _Rollup.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Rollup contract.
type RollupInitializedIterator struct {
	Event *RollupInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupInitialized represents a Initialized event raised by the Rollup contract.
type RollupInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) FilterInitialized(opts *bind.FilterOpts) (*RollupInitializedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RollupInitializedIterator{contract: _Rollup.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RollupInitialized) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupInitialized)
				if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Rollup *RollupFilterer) ParseInitialized(log types.Log) (*RollupInitialized, error) {
	event := new(RollupInitialized)
	if err := _Rollup.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rollup contract.
type RollupOwnershipTransferredIterator struct {
	Event *RollupOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupOwnershipTransferred represents a OwnershipTransferred event raised by the Rollup contract.
type RollupOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupOwnershipTransferredIterator{contract: _Rollup.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupOwnershipTransferred)
				if err := _Rollup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollup *RollupFilterer) ParseOwnershipTransferred(log types.Log) (*RollupOwnershipTransferred, error) {
	event := new(RollupOwnershipTransferred)
	if err := _Rollup.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Rollup contract.
type RollupPausedIterator struct {
	Event *RollupPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupPaused represents a Paused event raised by the Rollup contract.
type RollupPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupPausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupPausedIterator{contract: _Rollup.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupPaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupPaused)
				if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Rollup *RollupFilterer) ParsePaused(log types.Log) (*RollupPaused, error) {
	event := new(RollupPaused)
	if err := _Rollup.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupProveRemainingClaimedIterator is returned from FilterProveRemainingClaimed and is used to iterate over the raw logs and unpacked data for ProveRemainingClaimed events raised by the Rollup contract.
type RollupProveRemainingClaimedIterator struct {
	Event *RollupProveRemainingClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupProveRemainingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupProveRemainingClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupProveRemainingClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupProveRemainingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupProveRemainingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupProveRemainingClaimed represents a ProveRemainingClaimed event raised by the Rollup contract.
type RollupProveRemainingClaimed struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProveRemainingClaimed is a free log retrieval operation binding the contract event 0xb1b2058a6969e2d25e47bcaebe8ae21c29a23b2752429315b75e2f4f285f3d87.
//
// Solidity: event ProveRemainingClaimed(address receiver, uint256 amount)
func (_Rollup *RollupFilterer) FilterProveRemainingClaimed(opts *bind.FilterOpts) (*RollupProveRemainingClaimedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "ProveRemainingClaimed")
	if err != nil {
		return nil, err
	}
	return &RollupProveRemainingClaimedIterator{contract: _Rollup.contract, event: "ProveRemainingClaimed", logs: logs, sub: sub}, nil
}

// WatchProveRemainingClaimed is a free log subscription operation binding the contract event 0xb1b2058a6969e2d25e47bcaebe8ae21c29a23b2752429315b75e2f4f285f3d87.
//
// Solidity: event ProveRemainingClaimed(address receiver, uint256 amount)
func (_Rollup *RollupFilterer) WatchProveRemainingClaimed(opts *bind.WatchOpts, sink chan<- *RollupProveRemainingClaimed) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "ProveRemainingClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupProveRemainingClaimed)
				if err := _Rollup.contract.UnpackLog(event, "ProveRemainingClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProveRemainingClaimed is a log parse operation binding the contract event 0xb1b2058a6969e2d25e47bcaebe8ae21c29a23b2752429315b75e2f4f285f3d87.
//
// Solidity: event ProveRemainingClaimed(address receiver, uint256 amount)
func (_Rollup *RollupFilterer) ParseProveRemainingClaimed(log types.Log) (*RollupProveRemainingClaimed, error) {
	event := new(RollupProveRemainingClaimed)
	if err := _Rollup.contract.UnpackLog(event, "ProveRemainingClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRevertBatchIterator is returned from FilterRevertBatch and is used to iterate over the raw logs and unpacked data for RevertBatch events raised by the Rollup contract.
type RollupRevertBatchIterator struct {
	Event *RollupRevertBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupRevertBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRevertBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupRevertBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupRevertBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRevertBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRevertBatch represents a RevertBatch event raised by the Rollup contract.
type RollupRevertBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevertBatch is a free log retrieval operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*RollupRevertBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupRevertBatchIterator{contract: _Rollup.contract, event: "RevertBatch", logs: logs, sub: sub}, nil
}

// WatchRevertBatch is a free log subscription operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) WatchRevertBatch(opts *bind.WatchOpts, sink chan<- *RollupRevertBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRevertBatch)
				if err := _Rollup.contract.UnpackLog(event, "RevertBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRevertBatch is a log parse operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_Rollup *RollupFilterer) ParseRevertBatch(log types.Log) (*RollupRevertBatch, error) {
	event := new(RollupRevertBatch)
	if err := _Rollup.contract.UnpackLog(event, "RevertBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupSubmitterContractUpdatedIterator is returned from FilterSubmitterContractUpdated and is used to iterate over the raw logs and unpacked data for SubmitterContractUpdated events raised by the Rollup contract.
type RollupSubmitterContractUpdatedIterator struct {
	Event *RollupSubmitterContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupSubmitterContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupSubmitterContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupSubmitterContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupSubmitterContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupSubmitterContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupSubmitterContractUpdated represents a SubmitterContractUpdated event raised by the Rollup contract.
type RollupSubmitterContractUpdated struct {
	OldAddr common.Address
	NewAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitterContractUpdated is a free log retrieval operation binding the contract event 0xde3ada76ae1aa387cc4c586c70304b9fae54b251958786c828aaa0ce0290f939.
//
// Solidity: event SubmitterContractUpdated(address indexed oldAddr, address indexed newAddr)
func (_Rollup *RollupFilterer) FilterSubmitterContractUpdated(opts *bind.FilterOpts, oldAddr []common.Address, newAddr []common.Address) (*RollupSubmitterContractUpdatedIterator, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "SubmitterContractUpdated", oldAddrRule, newAddrRule)
	if err != nil {
		return nil, err
	}
	return &RollupSubmitterContractUpdatedIterator{contract: _Rollup.contract, event: "SubmitterContractUpdated", logs: logs, sub: sub}, nil
}

// WatchSubmitterContractUpdated is a free log subscription operation binding the contract event 0xde3ada76ae1aa387cc4c586c70304b9fae54b251958786c828aaa0ce0290f939.
//
// Solidity: event SubmitterContractUpdated(address indexed oldAddr, address indexed newAddr)
func (_Rollup *RollupFilterer) WatchSubmitterContractUpdated(opts *bind.WatchOpts, sink chan<- *RollupSubmitterContractUpdated, oldAddr []common.Address, newAddr []common.Address) (event.Subscription, error) {

	var oldAddrRule []interface{}
	for _, oldAddrItem := range oldAddr {
		oldAddrRule = append(oldAddrRule, oldAddrItem)
	}
	var newAddrRule []interface{}
	for _, newAddrItem := range newAddr {
		newAddrRule = append(newAddrRule, newAddrItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "SubmitterContractUpdated", oldAddrRule, newAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupSubmitterContractUpdated)
				if err := _Rollup.contract.UnpackLog(event, "SubmitterContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitterContractUpdated is a log parse operation binding the contract event 0xde3ada76ae1aa387cc4c586c70304b9fae54b251958786c828aaa0ce0290f939.
//
// Solidity: event SubmitterContractUpdated(address indexed oldAddr, address indexed newAddr)
func (_Rollup *RollupFilterer) ParseSubmitterContractUpdated(log types.Log) (*RollupSubmitterContractUpdated, error) {
	event := new(RollupSubmitterContractUpdated)
	if err := _Rollup.contract.UnpackLog(event, "SubmitterContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Rollup contract.
type RollupUnpausedIterator struct {
	Event *RollupUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUnpaused represents a Unpaused event raised by the Rollup contract.
type RollupUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupUnpausedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupUnpausedIterator{contract: _Rollup.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupUnpaused) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUnpaused)
				if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Rollup *RollupFilterer) ParseUnpaused(log types.Log) (*RollupUnpaused, error) {
	event := new(RollupUnpaused)
	if err := _Rollup.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateChallengerIterator is returned from FilterUpdateChallenger and is used to iterate over the raw logs and unpacked data for UpdateChallenger events raised by the Rollup contract.
type RollupUpdateChallengerIterator struct {
	Event *RollupUpdateChallenger // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateChallengerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateChallenger)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateChallenger)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateChallengerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateChallengerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateChallenger represents a UpdateChallenger event raised by the Rollup contract.
type RollupUpdateChallenger struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateChallenger is a free log retrieval operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) FilterUpdateChallenger(opts *bind.FilterOpts, account []common.Address) (*RollupUpdateChallengerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateChallenger", accountRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateChallengerIterator{contract: _Rollup.contract, event: "UpdateChallenger", logs: logs, sub: sub}, nil
}

// WatchUpdateChallenger is a free log subscription operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) WatchUpdateChallenger(opts *bind.WatchOpts, sink chan<- *RollupUpdateChallenger, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateChallenger", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateChallenger)
				if err := _Rollup.contract.UnpackLog(event, "UpdateChallenger", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateChallenger is a log parse operation binding the contract event 0x7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc009.
//
// Solidity: event UpdateChallenger(address indexed account, bool status)
func (_Rollup *RollupFilterer) ParseUpdateChallenger(log types.Log) (*RollupUpdateChallenger, error) {
	event := new(RollupUpdateChallenger)
	if err := _Rollup.contract.UnpackLog(event, "UpdateChallenger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateFinalizationPeriodSecondsIterator is returned from FilterUpdateFinalizationPeriodSeconds and is used to iterate over the raw logs and unpacked data for UpdateFinalizationPeriodSeconds events raised by the Rollup contract.
type RollupUpdateFinalizationPeriodSecondsIterator struct {
	Event *RollupUpdateFinalizationPeriodSeconds // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateFinalizationPeriodSeconds)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateFinalizationPeriodSeconds)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateFinalizationPeriodSecondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateFinalizationPeriodSeconds represents a UpdateFinalizationPeriodSeconds event raised by the Rollup contract.
type RollupUpdateFinalizationPeriodSeconds struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateFinalizationPeriodSeconds is a free log retrieval operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) FilterUpdateFinalizationPeriodSeconds(opts *bind.FilterOpts) (*RollupUpdateFinalizationPeriodSecondsIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateFinalizationPeriodSeconds")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateFinalizationPeriodSecondsIterator{contract: _Rollup.contract, event: "UpdateFinalizationPeriodSeconds", logs: logs, sub: sub}, nil
}

// WatchUpdateFinalizationPeriodSeconds is a free log subscription operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) WatchUpdateFinalizationPeriodSeconds(opts *bind.WatchOpts, sink chan<- *RollupUpdateFinalizationPeriodSeconds) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateFinalizationPeriodSeconds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateFinalizationPeriodSeconds)
				if err := _Rollup.contract.UnpackLog(event, "UpdateFinalizationPeriodSeconds", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateFinalizationPeriodSeconds is a log parse operation binding the contract event 0xa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437.
//
// Solidity: event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) ParseUpdateFinalizationPeriodSeconds(log types.Log) (*RollupUpdateFinalizationPeriodSeconds, error) {
	event := new(RollupUpdateFinalizationPeriodSeconds)
	if err := _Rollup.contract.UnpackLog(event, "UpdateFinalizationPeriodSeconds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateProofRewardPercentIterator is returned from FilterUpdateProofRewardPercent and is used to iterate over the raw logs and unpacked data for UpdateProofRewardPercent events raised by the Rollup contract.
type RollupUpdateProofRewardPercentIterator struct {
	Event *RollupUpdateProofRewardPercent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateProofRewardPercentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateProofRewardPercent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateProofRewardPercent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateProofRewardPercentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateProofRewardPercentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateProofRewardPercent represents a UpdateProofRewardPercent event raised by the Rollup contract.
type RollupUpdateProofRewardPercent struct {
	OldPercent *big.Int
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateProofRewardPercent is a free log retrieval operation binding the contract event 0xfb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b0223.
//
// Solidity: event UpdateProofRewardPercent(uint256 oldPercent, uint256 newPercent)
func (_Rollup *RollupFilterer) FilterUpdateProofRewardPercent(opts *bind.FilterOpts) (*RollupUpdateProofRewardPercentIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateProofRewardPercent")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateProofRewardPercentIterator{contract: _Rollup.contract, event: "UpdateProofRewardPercent", logs: logs, sub: sub}, nil
}

// WatchUpdateProofRewardPercent is a free log subscription operation binding the contract event 0xfb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b0223.
//
// Solidity: event UpdateProofRewardPercent(uint256 oldPercent, uint256 newPercent)
func (_Rollup *RollupFilterer) WatchUpdateProofRewardPercent(opts *bind.WatchOpts, sink chan<- *RollupUpdateProofRewardPercent) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateProofRewardPercent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateProofRewardPercent)
				if err := _Rollup.contract.UnpackLog(event, "UpdateProofRewardPercent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateProofRewardPercent is a log parse operation binding the contract event 0xfb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b0223.
//
// Solidity: event UpdateProofRewardPercent(uint256 oldPercent, uint256 newPercent)
func (_Rollup *RollupFilterer) ParseUpdateProofRewardPercent(log types.Log) (*RollupUpdateProofRewardPercent, error) {
	event := new(RollupUpdateProofRewardPercent)
	if err := _Rollup.contract.UnpackLog(event, "UpdateProofRewardPercent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateProofWindowIterator is returned from FilterUpdateProofWindow and is used to iterate over the raw logs and unpacked data for UpdateProofWindow events raised by the Rollup contract.
type RollupUpdateProofWindowIterator struct {
	Event *RollupUpdateProofWindow // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateProofWindowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateProofWindow)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateProofWindow)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateProofWindowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateProofWindowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateProofWindow represents a UpdateProofWindow event raised by the Rollup contract.
type RollupUpdateProofWindow struct {
	OldWindow *big.Int
	NewWindow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateProofWindow is a free log retrieval operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) FilterUpdateProofWindow(opts *bind.FilterOpts) (*RollupUpdateProofWindowIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateProofWindow")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateProofWindowIterator{contract: _Rollup.contract, event: "UpdateProofWindow", logs: logs, sub: sub}, nil
}

// WatchUpdateProofWindow is a free log subscription operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) WatchUpdateProofWindow(opts *bind.WatchOpts, sink chan<- *RollupUpdateProofWindow) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateProofWindow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateProofWindow)
				if err := _Rollup.contract.UnpackLog(event, "UpdateProofWindow", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateProofWindow is a log parse operation binding the contract event 0x1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61.
//
// Solidity: event UpdateProofWindow(uint256 oldWindow, uint256 newWindow)
func (_Rollup *RollupFilterer) ParseUpdateProofWindow(log types.Log) (*RollupUpdateProofWindow, error) {
	event := new(RollupUpdateProofWindow)
	if err := _Rollup.contract.UnpackLog(event, "UpdateProofWindow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateRollupDelayPeriodIterator is returned from FilterUpdateRollupDelayPeriod and is used to iterate over the raw logs and unpacked data for UpdateRollupDelayPeriod events raised by the Rollup contract.
type RollupUpdateRollupDelayPeriodIterator struct {
	Event *RollupUpdateRollupDelayPeriod // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateRollupDelayPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateRollupDelayPeriod)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateRollupDelayPeriod)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateRollupDelayPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateRollupDelayPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateRollupDelayPeriod represents a UpdateRollupDelayPeriod event raised by the Rollup contract.
type RollupUpdateRollupDelayPeriod struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateRollupDelayPeriod is a free log retrieval operation binding the contract event 0x624e47dc9fb8f8cfeaf4ead4710277cc1757136cfa885e465514cf6d510f0ad1.
//
// Solidity: event UpdateRollupDelayPeriod(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) FilterUpdateRollupDelayPeriod(opts *bind.FilterOpts) (*RollupUpdateRollupDelayPeriodIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateRollupDelayPeriod")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateRollupDelayPeriodIterator{contract: _Rollup.contract, event: "UpdateRollupDelayPeriod", logs: logs, sub: sub}, nil
}

// WatchUpdateRollupDelayPeriod is a free log subscription operation binding the contract event 0x624e47dc9fb8f8cfeaf4ead4710277cc1757136cfa885e465514cf6d510f0ad1.
//
// Solidity: event UpdateRollupDelayPeriod(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) WatchUpdateRollupDelayPeriod(opts *bind.WatchOpts, sink chan<- *RollupUpdateRollupDelayPeriod) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateRollupDelayPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateRollupDelayPeriod)
				if err := _Rollup.contract.UnpackLog(event, "UpdateRollupDelayPeriod", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateRollupDelayPeriod is a log parse operation binding the contract event 0x624e47dc9fb8f8cfeaf4ead4710277cc1757136cfa885e465514cf6d510f0ad1.
//
// Solidity: event UpdateRollupDelayPeriod(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) ParseUpdateRollupDelayPeriod(log types.Log) (*RollupUpdateRollupDelayPeriod, error) {
	event := new(RollupUpdateRollupDelayPeriod)
	if err := _Rollup.contract.UnpackLog(event, "UpdateRollupDelayPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupUpdateVerifierIterator is returned from FilterUpdateVerifier and is used to iterate over the raw logs and unpacked data for UpdateVerifier events raised by the Rollup contract.
type RollupUpdateVerifierIterator struct {
	Event *RollupUpdateVerifier // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RollupUpdateVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateVerifier)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RollupUpdateVerifier)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RollupUpdateVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateVerifier represents a UpdateVerifier event raised by the Rollup contract.
type RollupUpdateVerifier struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateVerifier is a free log retrieval operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) FilterUpdateVerifier(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*RollupUpdateVerifierIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateVerifierIterator{contract: _Rollup.contract, event: "UpdateVerifier", logs: logs, sub: sub}, nil
}

// WatchUpdateVerifier is a free log subscription operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) WatchUpdateVerifier(opts *bind.WatchOpts, sink chan<- *RollupUpdateVerifier, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateVerifier)
				if err := _Rollup.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateVerifier is a log parse operation binding the contract event 0x728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96.
//
// Solidity: event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_Rollup *RollupFilterer) ParseUpdateVerifier(log types.Log) (*RollupUpdateVerifier, error) {
	event := new(RollupUpdateVerifier)
	if err := _Rollup.contract.UnpackLog(event, "UpdateVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
