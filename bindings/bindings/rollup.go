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

// IRollupBatchSignatureInput is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchSignatureInput struct {
	SignedSequencersBitmap *big.Int
	SequencerSets          []byte
	Signature              []byte
}

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTiming\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"name\":\"ChallengeRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChallengeRewardClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"}],\"name\":\"ChallengeState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProveRemainingClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"RollupDelayPeriodUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateChallenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateFinalizationPeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"UpdateProofRewardPercent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"UpdateProofWindow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LAYER_2_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__maxNumTxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"batchChallengeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchChallenged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchChallengedSuccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchDataStore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signedSequencersBitmap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInsideChallengeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_batchHash\",\"type\":\"bytes32\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"challengeSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimProveRemaining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"lastBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numL1Messages\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIRollup.BatchDataInput\",\"name\":\"batchDataInput\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signedSequencersBitmap\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerSets\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRollup.BatchSignatureInput\",\"name\":\"batchSignatureInput\",\"type\":\"tuple\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"lastBlockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"numL1Messages\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIRollup.BatchDataInput\",\"name\":\"batchDataInput\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"signedSequencersBitmap\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sequencerSets\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRollup.BatchSignatureInput\",\"name\":\"batchSignatureInput\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_batchProof\",\"type\":\"bytes\"}],\"name\":\"commitBatchWithProof\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizationPeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"}],\"name\":\"finalizeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1StakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofWindow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofRewardPercent\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_prevStateRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rollupDelayPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"}],\"name\":\"isChallenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isChallenger\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1StakingContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCommittedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofRewardPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proveRemaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_batchProof\",\"type\":\"bytes\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertReqIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupDelayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateFinalizePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newWindow\",\"type\":\"uint256\"}],\"name\":\"updateProofWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newProofRewardPercent\",\"type\":\"uint256\"}],\"name\":\"updateRewardPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateRollupDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a060405234801562000010575f80fd5b5060405162004fa438038062004fa483398101604081905262000033916200010f565b6001600160401b0381166080526200004a62000051565b506200013e565b5f54610100900460ff1615620000bd5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff908116146200010d575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f6020828403121562000120575f80fd5b81516001600160401b038116811462000137575f80fd5b9392505050565b608051614e466200015e5f395f818161062b0152613c150152614e465ff3fe608060405260043610610332575f3560e01c806388b1ea09116101a7578063bedb86fb116100e7578063d8dc99d211610092578063dff7827e1161006d578063dff7827e14610a2e578063e3fff1dd14610a43578063f2fde38b14610a62578063fb1e8b0414610a81575f80fd5b8063d8dc99d2146109db578063ddd8a3dc146109f0578063de8b303514610a0f575f80fd5b8063ce5db8d6116100c2578063ce5db8d614610988578063cf9a67451461099d578063d279c191146109bc575f80fd5b8063bedb86fb14610937578063c555389214610956578063cd4edc6914610975575f80fd5b8063a479265d11610152578063b31a77d31161012d578063b31a77d3146108c5578063b3484425146108da578063b35dac4e146108f9578063b3e0a50914610918575f80fd5b8063a479265d14610866578063a4f209b01461087b578063abc8d68d1461089a575f80fd5b8063910129d411610182578063910129d4146107e857806397fc007c14610819578063a415d8dc14610838575f80fd5b806388b1ea09146107025780638da5cb5b1461071b5780638f1d377614610738575f80fd5b80633b70c18a116102725780635ef7a94a1161021d57806368589dfa116101f857806368589dfa146106855780636c578c1d146106b0578063715018a6146106cf578063728cdbca146106e3575f80fd5b80635ef7a94a146105ba5780635f77cf1d1461061a5780636126729014610666575f80fd5b80634e8f1d671161024d5780634e8f1d671461057157806357e0af6c146105845780635c975abb146105a3575f80fd5b80633b70c18a1461052a5780633e001b6614610549578063428868b51461055e575f80fd5b806313361101116102dd57806321e2f9e0116102b857806321e2f9e01461047e5780632362f03e1461049d5780632571098d146104c85780632b7ac3f3146104f3575f80fd5b8063133611011461041a57806318463fb01461043957806318af3b2b1461044e575f80fd5b806310d445831161030d57806310d44583146103c4578063116a1f42146103e3578063121dcd5014610405575f80fd5b806304d772151461033d578063059def61146103805780630ceb6780146103a3575f80fd5b3661033957005b5f80fd5b348015610348575f80fd5b5061036b61035736600461461f565b60a36020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561038b575f80fd5b50610395609d5481565b604051908152602001610377565b3480156103ae575f80fd5b506103c26103bd366004614651565b610a96565b005b3480156103cf575f80fd5b506103c26103de3660046146af565b610b68565b3480156103ee575f80fd5b5061036b6103fd36600461461f565b609d54101590565b348015610410575f80fd5b50610395609e5481565b348015610425575f80fd5b506103c26104343660046146f7565b610ea4565b348015610444575f80fd5b5061039560a75481565b348015610459575f80fd5b5061036b61046836600461461f565b5f90815260a26020526040902060010154421090565b348015610489575f80fd5b5061036b61049836600461461f565b61137e565b3480156104a8575f80fd5b506103956104b736600461461f565b60a16020525f908152604090205481565b3480156104d3575f80fd5b506103956104e236600461461f565b60a06020525f908152604090205481565b3480156104fe575f80fd5b50609c54610512906001600160a01b031681565b6040516001600160a01b039091168152602001610377565b348015610535575f80fd5b50609b54610512906001600160a01b031681565b348015610554575f80fd5b5061039560aa5481565b6103c261056c36600461475c565b6113ac565b6103c261057f3660046147bc565b6115e8565b34801561058f575f80fd5b506103c261059e36600461461f565b6118b6565b3480156105ae575f80fd5b5060655460ff1661036b565b3480156105c5575f80fd5b506105fa6105d436600461461f565b60a26020525f908152604090208054600182015460028301546003909301549192909184565b604080519485526020850193909352918301526060820152608001610377565b348015610625575f80fd5b5061064d7f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff9091168152602001610377565b348015610671575f80fd5b506103c261068036600461461f565b611961565b348015610690575f80fd5b5061039561069f36600461461f565b60ab6020525f908152604090205481565b3480156106bb575f80fd5b506103c26106ca366004614651565b611b3c565b3480156106da575f80fd5b506103c2611bf9565b3480156106ee575f80fd5b506103c26106fd36600461486d565b611c0c565b34801561070d575f80fd5b5060a65461036b9060ff1681565b348015610726575f80fd5b506033546001600160a01b0316610512565b348015610743575f80fd5b506107a461075236600461461f565b60a46020525f9081526040902080546001820154600283015460039093015467ffffffffffffffff831693680100000000000000009093046001600160a01b0316929060ff8082169161010090041686565b6040805167ffffffffffffffff90971687526001600160a01b03909516602087015293850192909252606084015215156080830152151560a082015260c001610377565b3480156107f3575f80fd5b5061036b61080236600461461f565b5f90815260a4602052604090206003015460ff1690565b348015610824575f80fd5b506103c2610833366004614651565b611f15565b348015610843575f80fd5b5061036b610852366004614651565b609f6020525f908152604090205460ff1681565b348015610871575f80fd5b5061039560995481565b348015610886575f80fd5b506103c261089536600461461f565b611ff8565b3480156108a5575f80fd5b506103956108b4366004614651565b60a56020525f908152604090205481565b3480156108d0575f80fd5b5061039560a85481565b3480156108e5575f80fd5b506103c26108f43660046148c8565b6120a8565b348015610904575f80fd5b506103c2610913366004614651565b6123ac565b348015610923575f80fd5b506103c26109323660046146f7565b612405565b348015610942575f80fd5b506103c261095136600461493c565b612755565b348015610961575f80fd5b506103c261097036600461461f565b612893565b6103c261098336600461496e565b612936565b348015610993575f80fd5b5061039560985481565b3480156109a8575f80fd5b506103c26109b736600461461f565b612ec8565b3480156109c7575f80fd5b506103c26109d6366004614651565b613037565b3480156109e6575f80fd5b5061039560ac5481565b3480156109fb575f80fd5b50609754610512906001600160a01b031681565b348015610a1a575f80fd5b5061036b610a2936600461461f565b6130f6565b348015610a39575f80fd5b50610395609a5481565b348015610a4e575f80fd5b506103c2610a5d36600461461f565b613140565b348015610a6d575f80fd5b506103c2610a7c366004614651565b6131e3565b348015610a8c575f80fd5b5061039560a95481565b610a9e613270565b6001600160a01b0381165f908152609f602052604090205460ff1615610b0b5760405162461bcd60e51b815260206004820152601f60248201527f6163636f756e7420697320616c72656164792061206368616c6c656e6765720060448201526064015b60405180910390fd5b6001600160a01b0381165f818152609f6020908152604091829020805460ff1916600190811790915591519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc00991015b60405180910390a250565b610b70613270565b5f8111610bbf5760405162461bcd60e51b815260206004820152601560248201527f636f756e74206d757374206265206e6f6e7a65726f00000000000000000000006044820152606401610b02565b5f80610bcb85856132ca565b915091505f610bde836001015160c01c90565b5f81815260a160205260409020549091508214610c3d5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610b02565b5f60a181610c4b87856149c3565b81526020019081526020015f205414610ccb5760405162461bcd60e51b8152602060048201526024808201527f726576657274696e67206d7573742073746172742066726f6d2074686520656e60448201527f64696e67000000000000000000000000000000000000000000000000000000006064820152608401610b02565b609d548111610d425760405162461bcd60e51b815260206004820152602160248201527f63616e206f6e6c792072657665727420756e46696e616c697a6564206261746360448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610b02565b610d4d6001826149d6565b609e555b8315610e9c57604051829082907ecae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146905f90a35f81815260a16020526040812055610d99816130f6565b15610df4575f81815260a460209081526040808320600181015490546801000000000000000090046001600160a01b0316845260a59092528220805491929091610de49084906149c3565b909155505060a6805460ff191690555b5f81815260a46020526040812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600181018290556002810191909155600301805461ffff1916905560a85415801590610e54575060a85481145b15610e5e575f60a8555b6001015f81815260a160205260409020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90940193915081610d51575b505050505050565b60a85415610ef45760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610b02565b610efc613363565b5f80610f0884846132ca565b915091505f610f1b836001015160c01c90565b5f81815260a160205260409020549091508214610f7a5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610b02565b610f838161137e565b610fcf5760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610b02565b610fd8816130f6565b156110255760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610b02565b5f81815260a4602052604090206003015460ff16156110865760405162461bcd60e51b815260206004820152601660248201527f62617463682073686f756c6420626520726576657274000000000000000000006044820152606401610b02565b5f81815260a260205260409020600101544210156110e65760405162461bcd60e51b815260206004820152601960248201527f626174636820696e206368616c6c656e67652077696e646f77000000000000006044820152606401610b02565b605983015160a05f6110f96001856149d6565b81526020019081526020015f2054146111545760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610b02565b5f81815260a06020526040902054156111af5760405162461bcd60e51b815260206004820152601660248201527f626174636820616c7265616479207665726966696564000000000000000000006044820152606401610b02565b80609d54600101146112035760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610b02565b609d819055600160a35f611218866099015190565b815260208101919091526040015f20805460ff191691151591909117905560798301515f82815260a0602052604090205561126761125a846011015160c01c90565b600985015160c01c6133b6565b60a25f6112756001846149d6565b815260208101919091526040015f908120818155600180820183905560028201839055600390910182905560ab91906112ae90846149d6565b81526020019081526020015f205f905560a45f6001836112ce91906149d6565b815260208082019290925260409081015f90812080547fffffffff000000000000000000000000000000000000000000000000000000001681556001810182905560028101829055600301805461ffff1916905583815260a1909252902054817f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d61135a866079015190565b60998701516040805192835260208301919091520160405180910390a35050505050565b5f81815260a26020526040812054158015906113a657505f82815260a1602052604090205415155b92915050565b6097546001600160a01b03166368015791336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611418573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061143c91906149e9565b6114885760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c7920616374697665207374616b657220616c6c6f7765640000000000006044820152606401610b02565b60a854156114d85760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610b02565b6114e0613363565b60ac54609b54604080517fb59b1a7800000000000000000000000000000000000000000000000000000000815290514293926001600160a01b03169163b59b1a789160048083019260209291908290030181865afa158015611544573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115689190614a04565b61157291906149c3565b10156115da575f6115896080840160608501614a1b565b61ffff16116115da5760405162461bcd60e51b815260206004820152600b60248201527f6c316d73672064656c61790000000000000000000000000000000000000000006044820152606401610b02565b6115e48282613473565b5050565b60a854156116385760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610b02565b611640613363565b60ac54609e545f90815260a260205260408120549091429161166291906149c3565b1090505f4260ac54609b5f9054906101000a90046001600160a01b03166001600160a01b031663b59b1a786040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116de9190614a04565b6116e891906149c3565b109050811580156116f65750805b1561175d575f61170c60808a0160608b01614a1b565b61ffff161161175d5760405162461bcd60e51b815260206004820152600b60248201527f6c316d73672064656c61790000000000000000000000000000000000000000006044820152606401610b02565b81806117665750805b6117b25760405162461bcd60e51b815260206004820152600e60248201527f696e76616c69642074696d696e670000000000000000000000000000000000006044820152606401610b02565b6117bc8888613473565b5f806117c888886132ca565b915091505f6117db836001015160c01c90565b905080609e541461182e5760405162461bcd60e51b815260206004820152601660248201527f696e636f727265637420626174636820686561646572000000000000000000006044820152606401610b02565b5f81815260a16020526040902054821461188a5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610b02565b5f81815260a260205260409020426001909101556118a9838888613ba6565b5050505050505050505050565b6118be613270565b5f811180156118cf57506099548114155b61191b5760405162461bcd60e51b815260206004820152601860248201527f696e76616c6964206e65772070726f6f662077696e646f7700000000000000006044820152606401610b02565b609980549082905560408051828152602081018490527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a6191015b60405180910390a15050565b5f54600290610100900460ff1615801561198157505f5460ff8083169116105b6119f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b02565b5f805461ffff191660ff831617610100179055611a115f5460ff1690565b60ff16600214611a635760405162461bcd60e51b815260206004820152601660248201527f6d757374206861766520696e697469616c697a656421000000000000000000006044820152606401610b02565b81611ad65760405162461bcd60e51b815260206004820152602760248201527f63616e206e6f742073657420737461746520726f6f742077697468206279746560448201527f73333228302921000000000000000000000000000000000000000000000000006064820152608401610b02565b609e545f90815260ab6020526040902054611aff57609e545f90815260ab602052604090208290555b5f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611955565b611b44613270565b6001600160a01b0381165f908152609f602052604090205460ff16611bab5760405162461bcd60e51b815260206004820152601b60248201527f6163636f756e74206973206e6f742061206368616c6c656e67657200000000006044820152606401610b02565b6001600160a01b0381165f818152609f60209081526040808320805460ff19169055519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc0099101610b5d565b611c01613270565b611c0a5f613d59565b565b5f54610100900460ff1615808015611c2a57505f54600160ff909116105b80611c435750303b158015611c4357505f5460ff166001145b611cb55760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b02565b5f805460ff191660011790558015611cd6575f805461ff0019166101001790555b6001600160a01b0386161580611cf357506001600160a01b038516155b15611d2a576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b038716611d805760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206c31207374616b696e6720636f6e747261637400000000006044820152606401610b02565b611d88613dc2565b611d90613e46565b609780546001600160a01b03808a167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609b8054898416908316179055609c805492881692909116821790556098859055609984905560a98390556040515f907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96908290a3604080515f8152602081018690527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437910160405180910390a1604080515f8152602081018590527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61910160405180910390a1604080515f8152602081018490527ffb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b0223910160405180910390a18015611f0c575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b611f1d613270565b6001600160a01b03811615801590611f435750609c546001600160a01b03828116911614155b611f8f5760405162461bcd60e51b815260206004820152601460248201527f696e76616c6964206e65772076657269666965720000000000000000000000006044820152606401610b02565b609c80546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96905f90a35050565b612000613270565b5f81118015612010575060648111155b801561201e575060a9548114155b61206a5760405162461bcd60e51b815260206004820152601f60248201527f696e76616c69642070726f6f66207265776172642070657263656e74616765006044820152606401610b02565b60a980549082905560408051828152602081018490527ffb81bce17f015797e11949d3c332e2bf9453faf68f728447426803138f2b02239101611955565b60a854156120f85760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610b02565b612100613363565b6097546001600160a01b03166368015791336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa15801561216c573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061219091906149e9565b6121dc5760405162461bcd60e51b815260206004820152601a60248201527f6f6e6c7920616374697665207374616b657220616c6c6f7765640000000000006044820152606401610b02565b5f806121e886866132ca565b915091505f6121fb836001015160c01c90565b5f81815260a16020526040902054909150821461225a5760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610b02565b612263816130f6565b6122af5760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610b02565b5f81815260a46020526040902060038101805461ff00191661010017905560a6805460ff1916905560995460029091015442916122eb916149c3565b11612361575f81815260a4602090815260408083206003908101805460ff1916600117905560a2835292819020909201548251808401909352600783527f54696d656f7574000000000000000000000000000000000000000000000000009183019190915261235c91839190613eca565b611f0c565b61236c838686613ba6565b611f0c81336040518060400160405280600d81526020017f50726f6f66207375636365737300000000000000000000000000000000000000815250614026565b6123b4613270565b60aa80545f9091556123c682826140f2565b604080516001600160a01b0384168152602081018390527fb1b2058a6969e2d25e47bcaebe8ae21c29a23b2752429315b75e2f4f285f3d879101611955565b61240d613270565b5f805260a06020527fb84a74ec6ef4d0e83b6006dfaa014ab4026f9f3b97d186e604d29998a4e808ea54156124845760405162461bcd60e51b815260206004820152601660248201527f67656e6573697320626174636820696d706f72746564000000000000000000006044820152606401610b02565b5f8061249084846132ca565b915091505f6124a3836001015160c01c90565b905080156124f35760405162461bcd60e51b815260206004820152601360248201527f696e76616c696420626174636820696e646578000000000000000000000000006044820152606401610b02565b5f6124ff846079015190565b90508061254e5760405162461bcd60e51b815260206004820152600f60248201527f7a65726f20737461746520726f6f7400000000000000000000000000000000006044820152606401610b02565b600984015160c01c156125a35760405162461bcd60e51b815260206004820152601d60248201527f6c31206d65737361676520706f707065642073686f756c6420626520300000006044820152606401610b02565b5f6125af856019015190565b036125fc5760405162461bcd60e51b815260206004820152600e60248201527f7a65726f206461746120686173680000000000000000000000000000000000006044820152606401610b02565b7f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014612628856039015190565b146126755760405162461bcd60e51b815260206004820152601660248201527f696e76616c69642076657273696f6e65642068617368000000000000000000006044820152606401610b02565b5f82815260a1602090815260408083208690558051608081018252428082528184019081528183018581526060830186815288875260a2865284872093518455915160018401555160028301555160039091015560ab825280832084905560a0909152808220839055609e849055609d84905551849184917f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f9190a3604080518281525f6020820152849184917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a3505050505050565b61275d613270565b80156128615761276b61419c565b60a65460ff16156128205760a7545f90815260a460209081526040808320600181015490546801000000000000000090046001600160a01b0316845260a590925282208054919290916127bf9084906149c3565b909155505060a7545f90815260a46020526040812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600181018290556002810191909155600301805461ffff1916905560a6805460ff191690555b7f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258335b6040516001600160a01b03909116815260200160405180910390a150565b6128696141f6565b7f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612843565b50565b61289b613270565b5f811180156128ac575060ac548114155b6128f85760405162461bcd60e51b815260206004820152601f60248201527f696e76616c6964206e657720726f6c6c75702064656c617920706572696f64006044820152606401610b02565b60ac80549082905560408051828152602081018490527f2392c64c3c2ac54ae8093f1b546601e09b9c3ee6086d2f2595db2d3c54f3f56b9101611955565b335f908152609f602052604090205460ff166129945760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206368616c6c656e67657220616c6c6f776564000000000000006044820152606401610b02565b60a854156129e45760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610b02565b6129ec613363565b60a65460ff1615612a3f5760405162461bcd60e51b815260206004820152601460248201527f616c726561647920696e206368616c6c656e67650000000000000000000000006044820152606401610b02565b8167ffffffffffffffff16609d5410612a9a5760405162461bcd60e51b815260206004820152601760248201527f626174636820616c72656164792066696e616c697a65640000000000000000006044820152606401610b02565b67ffffffffffffffff82165f90815260a160205260409020548114612b015760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610b02565b612b148267ffffffffffffffff1661137e565b612b605760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610b02565b67ffffffffffffffff82165f90815260a460205260409020546801000000000000000090046001600160a01b031615612bdb5760405162461bcd60e51b815260206004820152601860248201527f626174636820616c7265616479206368616c6c656e67656400000000000000006044820152606401610b02565b67ffffffffffffffff82165f90815260a260205260409020600101544210612c6b5760405162461bcd60e51b815260206004820152603360248201527f63616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f77000000000000000000000000006064820152608401610b02565b60975f9054906101000a90046001600160a01b03166001600160a01b0316630d13fd7b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cbb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cdf9190614a04565b341015612d2e5760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742076616c756500000000000000000000000000006044820152606401610b02565b67ffffffffffffffff82811660a78190556040805160c0810182528281523360208083018281523484860190815242606086019081525f6080870181815260a0880182815299825260a4909552969096209451855492516001600160a01b031668010000000000000000027fffffffff000000000000000000000000000000000000000000000000000000009093169816979097171783559451600183015591516002820155925160039093018054925115156101000261ff00199415159490941661ffff19909316929092179290921790556001600160a01b03168267ffffffffffffffff167f3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b997634604051612e4691815260200190565b60405180910390a35f609d546001612e5e91906149c3565b90505b609e548111612eb6578267ffffffffffffffff168114612ea4576099545f82815260a2602052604081206001018054909190612e9e9084906149c3565b90915550505b80612eae81614a3c565b915050612e61565b505060a6805460ff1916600117905550565b5f54600390610100900460ff16158015612ee857505f5460ff8083169116105b612f5a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610b02565b5f805461ffff191660ff831617610100178155829003612fbc5760405162461bcd60e51b815260206004820152601b60248201527f696e76616c696420726f6c6c75702064656c617920706572696f6400000000006044820152606401610b02565b60ac829055604080515f8152602081018490527f2392c64c3c2ac54ae8093f1b546601e09b9c3ee6086d2f2595db2d3c54f3f56b910160405180910390a15f805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001611955565b335f90815260a56020526040812054908190036130965760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642062617463684368616c6c656e6765526577617264000000006044820152606401610b02565b335f90815260a560205260408120556130af82826140f2565b816001600160a01b03167f9c25fa83f414ed363c8d39c98fb3e17567b3431cede71eb062c49d2a63ce247a826040516130ea91815260200190565b60405180910390a25050565b5f81815260a460205260408120546801000000000000000090046001600160a01b0316158015906113a65750505f90815260a46020526040902060030154610100900460ff161590565b613148613270565b5f8111801561315957506098548114155b6131a55760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206e65772066696e616c697a6520706572696f6400000000006044820152606401610b02565b609880549082905560408051828152602081018490527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a4379101611955565b6131eb613270565b6001600160a01b0381166132675760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610b02565b61289081613d59565b6033546001600160a01b03163314611c0a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610b02565b5f805f6132d7858561422f565b90505f8160ff165f036132f8576132ee868661429f565b9094509050613355565b8160ff1660010361330d576132ee8686614308565b60405162461bcd60e51b815260206004820152601960248201527f556e737570706f727465642062617463682076657273696f6e000000000000006044820152606401610b02565b808420925050509250929050565b60655460ff1615611c0a5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610b02565b805f036133c1575050565b8082035f5b8281101561346d576101008184038111156133e057508083035b609b546040517f3c7f528300000000000000000000000000000000000000000000000000000000815260048101859052602481018390526001600160a01b0390911690633c7f5283906044015f604051808303815f87803b158015613443575f80fd5b505af1158015613455573d5f803e3d5ffd5b505050506101008301925050610100810190506133c6565b50505050565b6134806020830183614a73565b60ff16158061349e57506134976020830183614a73565b60ff166001145b6134ea5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c69642076657273696f6e00000000000000000000000000000000006044820152606401610b02565b608082013561353b5760405162461bcd60e51b815260206004820152601b60248201527f70726576696f757320737461746520726f6f74206973207a65726f00000000006044820152606401610b02565b60a082013561358c5760405162461bcd60e51b815260206004820152601660248201527f6e657720737461746520726f6f74206973207a65726f000000000000000000006044820152606401610b02565b5f806135a361359e6020860186614a93565b6132ca565b915091505f6135b6836001015160c01c90565b90505f60a1816135c78460016149c3565b81526020019081526020015f2054146136225760405162461bcd60e51b815260206004820152601760248201527f626174636820616c726561647920636f6d6d69747465640000000000000000006044820152606401610b02565b609e5481146136735760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610b02565b5f81815260a1602052604090205482146136cf5760405162461bcd60e51b815260206004820152601b60248201527f696e636f727265637420706172656e74206261746368206861736800000000006044820152606401610b02565b5f81815260ab602052604090205460808601351461372f5760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610b02565b5f61373e846011015160c01c90565b90505f61376a6137546060890160408a01614af4565b61376460808a0160608b01614a1b565b8461435b565b905061377c6080880160608901614a1b565b6001939093019261ffff1691909101905f80491561379b575f496137bd565b7f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440145b905060f96137ce60208a018a614a73565b60ff166001036137dd57506101015b604080518281019091529650613802876137fa60208c018c614a73565b60ff1661439f565b60c085901b60018801526138308761382060808c0160608d01614a1b565b61ffff1660c01b60099190910152565b60c084811b6011890152601988018490526039880183905260808a0135605989015260a08a0135607989015289013560998801526138948761387560208b018b614a93565b604051613883929190614b0d565b604051809103902060b99190910152565b60d9870186905260016138aa60208b018b614a73565b60ff16106138dc576138dc876138c660608c0160408d01614af4565b67ffffffffffffffff1660c01b60f99190910152565b8087205f86815260a1602090815260408083209390935560ab905290812060a08b0135905560a65460ff161561393b5760a7545f90815260a46020526040902060020154609954429161392e916149c3565b61393891906149d6565b90505b6040518060800160405280428152602001826098544261395b91906149c3565b61396591906149c3565b815260200161397a60608d0160408e01614af4565b67ffffffffffffffff1681526097546020909101906001600160a01b031663d096c3c6336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156139f8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a1c9190614a04565b90525f87815260a2602090815260409182902083518155838201516001820155918301516002830155606090920151600390910155609e8790556097546001600160a01b031692506374fe27b79150893590613a8490613a7e908c018c614a93565b5f6143a6565b5f613a9260408d018d614a93565b6040518663ffffffff1660e01b8152600401613ab2959493929190614b63565b602060405180830381865afa158015613acd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613af191906149e9565b613b635760405162461bcd60e51b815260206004820152602160248201527f746865207369676e617475726520766572696669636174696f6e206661696c6560448201527f64000000000000000000000000000000000000000000000000000000000000006064820152608401610b02565b5f84815260a16020526040808220549051909186917f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f9190a35050505050505050565b80613bf35760405162461bcd60e51b815260206004820152601360248201527f496e76616c69642062617463682070726f6f66000000000000000000000000006044820152606401610b02565b5f613c02846001015160c01c90565b90505f613c10856039015190565b90505f7f0000000000000000000000000000000000000000000000000000000000000000613c3f876059015190565b6079880151609989015160b98a015160198b015160405160c09690961b7fffffffffffffffff000000000000000000000000000000000000000000000000166020870152602886019490945260488501929092526068840152608883015260a882015260c8810183905260e801604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120609c549091506001600160a01b0316632c09a848613d01885160f81c90565b858888866040518663ffffffff1660e01b8152600401613d25959493929190614bd6565b5f6040518083038186803b158015613d3b575f80fd5b505afa158015613d4d573d5f803e3d5ffd5b50505050505050505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16613e3e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b02565b611c0a6143fa565b5f54610100900460ff16613ec25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b02565b611c0a614482565b60a88390555f83815260a460205260408082205460975491517f45bc4d1000000000000000000000000000000000000000000000000000000000815260048101869052680100000000000000009091046001600160a01b03908116939216906345bc4d10906024016020604051808303815f875af1158015613f4e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f729190614a04565b5f86815260a46020526040902060010154909150613f919082906149c3565b5f86815260a460209081526040808320546801000000000000000090046001600160a01b0316835260a590915281208054909190613fd09084906149c3565b9091555050604051613fe3908490614c07565b604051908190038120906001600160a01b0384169087907fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a905f90a45050505050565b5f83815260a4602052604081206001015460a95490919060649061404a9084614c33565b6140549190614c4a565b905061406081836149d6565b60aa5f82825461407091906149c3565b90915550506001600160a01b0384165f90815260a560205260408120805483929061409c9084906149c3565b90915550506040516140af908490614c07565b604051908190038120906001600160a01b0386169087907fe70d3820e244d5f71d1a6395db24f3460e8dca966edc1fd3625b6292880a877a905f90a45050505050565b80156115e4575f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114614141576040519150601f19603f3d011682016040523d82523d5f602084013e614146565b606091505b50509050806141975760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610b02565b505050565b6141a4613363565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586141d93390565b6040516001600160a01b03909116815260200160405180910390a1565b6141fe614507565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336141d9565b5f8161427d5760405162461bcd60e51b815260206004820152601260248201527f456d7074792062617463682068656164657200000000000000000000000000006044820152606401610b02565b82825f81811061428f5761428f614c82565b919091013560f81c949350505050565b5f8160f98110156142f25760405162461bcd60e51b815260206004820152601d60248201527f626174636820686561646572206c656e67746820746f6f20736d616c6c0000006044820152606401610b02565b6040519150808483378082016040529250929050565b5f8161010181146142f25760405162461bcd60e51b815260206004820181905260248201527f626174636820686561646572206c656e67746820697320696e636f72726563746044820152606401610b02565b6040805160c085901b815260f084901b6008820152600a60208502820181019092525f9181016143908161ffff871686614559565b82900390912095945050505050565b8082535050565b60605f808080806143b9888a018a614d9c565b95509550955095509550508187106143d75794506143f39350505050565b8387106143eb5782955050505050506143f3565b509293505050505b9392505050565b5f54610100900460ff166144765760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b02565b6065805460ff19169055565b5f54610100900460ff166144fe5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610b02565b611c0a33613d59565b60655460ff16611c0a5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b02565b5f825f036145685750826143f3565b609b546001600160a01b03165f5b84811015614615576040517fae453cd5000000000000000000000000000000000000000000000000000000008152600481018590525f906001600160a01b0384169063ae453cd590602401602060405180830381865afa1580156145dc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906146009190614a04565b87525060209095019460019384019301614576565b5093949350505050565b5f6020828403121561462f575f80fd5b5035919050565b80356001600160a01b038116811461464c575f80fd5b919050565b5f60208284031215614661575f80fd5b6143f382614636565b5f8083601f84011261467a575f80fd5b50813567ffffffffffffffff811115614691575f80fd5b6020830191508360208285010111156146a8575f80fd5b9250929050565b5f805f604084860312156146c1575f80fd5b833567ffffffffffffffff8111156146d7575f80fd5b6146e38682870161466a565b909790965060209590950135949350505050565b5f8060208385031215614708575f80fd5b823567ffffffffffffffff81111561471e575f80fd5b61472a8582860161466a565b90969095509350505050565b5f60e08284031215614746575f80fd5b50919050565b5f60608284031215614746575f80fd5b5f806040838503121561476d575f80fd5b823567ffffffffffffffff80821115614784575f80fd5b61479086838701614736565b935060208501359150808211156147a5575f80fd5b506147b28582860161474c565b9150509250929050565b5f805f805f80608087890312156147d1575f80fd5b863567ffffffffffffffff808211156147e8575f80fd5b6147f48a838b01614736565b97506020890135915080821115614809575f80fd5b6148158a838b0161474c565b9650604089013591508082111561482a575f80fd5b6148368a838b0161466a565b9096509450606089013591508082111561484e575f80fd5b5061485b89828a0161466a565b979a9699509497509295939492505050565b5f805f805f8060c08789031215614882575f80fd5b61488b87614636565b955061489960208801614636565b94506148a760408801614636565b9350606087013592506080870135915060a087013590509295509295509295565b5f805f80604085870312156148db575f80fd5b843567ffffffffffffffff808211156148f2575f80fd5b6148fe8883890161466a565b90965094506020870135915080821115614916575f80fd5b506149238782880161466a565b95989497509550505050565b8015158114612890575f80fd5b5f6020828403121561494c575f80fd5b81356143f38161492f565b803567ffffffffffffffff8116811461464c575f80fd5b5f806040838503121561497f575f80fd5b61498883614957565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156113a6576113a6614996565b818103818111156113a6576113a6614996565b5f602082840312156149f9575f80fd5b81516143f38161492f565b5f60208284031215614a14575f80fd5b5051919050565b5f60208284031215614a2b575f80fd5b813561ffff811681146143f3575f80fd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614a6c57614a6c614996565b5060010190565b5f60208284031215614a83575f80fd5b813560ff811681146143f3575f80fd5b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112614ac6575f80fd5b83018035915067ffffffffffffffff821115614ae0575f80fd5b6020019150368190038213156146a8575f80fd5b5f60208284031215614b04575f80fd5b6143f382614957565b818382375f9101908152919050565b81835281816020850137505f602082840101525f60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b5f6080820187835260206080602085015281885180845260a08601915060208a0193505f5b81811015614bad5784516001600160a01b031683529383019391830191600101614b88565b50508760408601528481036060860152614bc8818789614b1c565b9a9950505050505050505050565b858152846020820152608060408201525f614bf5608083018587614b1c565b90508260608301529695505050505050565b5f82515f5b81811015614c265760208186018101518583015201614c0c565b505f920191825250919050565b80820281158282048414176113a6576113a6614996565b5f82614c7d577f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b500490565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112614ceb575f80fd5b8135602067ffffffffffffffff80831115614d0857614d08614caf565b8260051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f83011681018181108482111715614d4b57614d4b614caf565b6040529384526020818701810194908101925087851115614d6a575f80fd5b6020870191505b84821015614d9157614d8282614636565b83529183019190830190614d71565b979650505050505050565b5f805f805f8060c08789031215614db1575f80fd5b86359550602087013567ffffffffffffffff80821115614dcf575f80fd5b614ddb8a838b01614cdc565b9650604089013595506060890135915080821115614df7575f80fd5b614e038a838b01614cdc565b94506080890135935060a0890135915080821115614e1f575f80fd5b50614e2c89828a01614cdc565b915050929550929550929556fea164736f6c6343000818000a",
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
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, uint256 signedSequencersBitmap)
func (_Rollup *RollupCaller) BatchDataStore(opts *bind.CallOpts, batchIndex *big.Int) (struct {
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	BlockNumber            *big.Int
	SignedSequencersBitmap *big.Int
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchDataStore", batchIndex)

	outstruct := new(struct {
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		BlockNumber            *big.Int
		SignedSequencersBitmap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FinalizeTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SignedSequencersBitmap = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BatchDataStore is a free data retrieval call binding the contract method 0x5ef7a94a.
//
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, uint256 signedSequencersBitmap)
func (_Rollup *RollupSession) BatchDataStore(batchIndex *big.Int) (struct {
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	BlockNumber            *big.Int
	SignedSequencersBitmap *big.Int
}, error) {
	return _Rollup.Contract.BatchDataStore(&_Rollup.CallOpts, batchIndex)
}

// BatchDataStore is a free data retrieval call binding the contract method 0x5ef7a94a.
//
// Solidity: function batchDataStore(uint256 batchIndex) view returns(uint256 originTimestamp, uint256 finalizeTimestamp, uint256 blockNumber, uint256 signedSequencersBitmap)
func (_Rollup *RollupCallerSession) BatchDataStore(batchIndex *big.Int) (struct {
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	BlockNumber            *big.Int
	SignedSequencersBitmap *big.Int
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

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupCaller) L1StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "l1StakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupSession) L1StakingContract() (common.Address, error) {
	return _Rollup.Contract.L1StakingContract(&_Rollup.CallOpts)
}

// L1StakingContract is a free data retrieval call binding the contract method 0xddd8a3dc.
//
// Solidity: function l1StakingContract() view returns(address)
func (_Rollup *RollupCallerSession) L1StakingContract() (common.Address, error) {
	return _Rollup.Contract.L1StakingContract(&_Rollup.CallOpts)
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

// CommitBatch is a paid mutator transaction binding the contract method 0x428868b5.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput) payable returns()
func (_Rollup *RollupTransactor) CommitBatch(opts *bind.TransactOpts, batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatch", batchDataInput, batchSignatureInput)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x428868b5.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput) payable returns()
func (_Rollup *RollupSession) CommitBatch(batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchDataInput, batchSignatureInput)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x428868b5.
//
// Solidity: function commitBatch((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput) payable returns()
func (_Rollup *RollupTransactorSession) CommitBatch(batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchDataInput, batchSignatureInput)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x4e8f1d67.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput, bytes _batchHeader, bytes _batchProof) payable returns()
func (_Rollup *RollupTransactor) CommitBatchWithProof(opts *bind.TransactOpts, batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatchWithProof", batchDataInput, batchSignatureInput, _batchHeader, _batchProof)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x4e8f1d67.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput, bytes _batchHeader, bytes _batchProof) payable returns()
func (_Rollup *RollupSession) CommitBatchWithProof(batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatchWithProof(&_Rollup.TransactOpts, batchDataInput, batchSignatureInput, _batchHeader, _batchProof)
}

// CommitBatchWithProof is a paid mutator transaction binding the contract method 0x4e8f1d67.
//
// Solidity: function commitBatchWithProof((uint8,bytes,uint64,uint16,bytes32,bytes32,bytes32) batchDataInput, (uint256,bytes,bytes) batchSignatureInput, bytes _batchHeader, bytes _batchProof) payable returns()
func (_Rollup *RollupTransactorSession) CommitBatchWithProof(batchDataInput IRollupBatchDataInput, batchSignatureInput IRollupBatchSignatureInput, _batchHeader []byte, _batchProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatchWithProof(&_Rollup.TransactOpts, batchDataInput, batchSignatureInput, _batchHeader, _batchProof)
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
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _l1StakingContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupSession) Initialize(_l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1StakingContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _finalizationPeriodSeconds, uint256 _proofWindow, uint256 _proofRewardPercent) returns()
func (_Rollup *RollupTransactorSession) Initialize(_l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int, _proofRewardPercent *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1StakingContract, _messageQueue, _verifier, _finalizationPeriodSeconds, _proofWindow, _proofRewardPercent)
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

// RollupRollupDelayPeriodUpdateIterator is returned from FilterRollupDelayPeriodUpdate and is used to iterate over the raw logs and unpacked data for RollupDelayPeriodUpdate events raised by the Rollup contract.
type RollupRollupDelayPeriodUpdateIterator struct {
	Event *RollupRollupDelayPeriodUpdate // Event containing the contract specifics and raw log

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
func (it *RollupRollupDelayPeriodUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRollupDelayPeriodUpdate)
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
		it.Event = new(RollupRollupDelayPeriodUpdate)
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
func (it *RollupRollupDelayPeriodUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRollupDelayPeriodUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRollupDelayPeriodUpdate represents a RollupDelayPeriodUpdate event raised by the Rollup contract.
type RollupRollupDelayPeriodUpdate struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRollupDelayPeriodUpdate is a free log retrieval operation binding the contract event 0x2392c64c3c2ac54ae8093f1b546601e09b9c3ee6086d2f2595db2d3c54f3f56b.
//
// Solidity: event RollupDelayPeriodUpdate(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) FilterRollupDelayPeriodUpdate(opts *bind.FilterOpts) (*RollupRollupDelayPeriodUpdateIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "RollupDelayPeriodUpdate")
	if err != nil {
		return nil, err
	}
	return &RollupRollupDelayPeriodUpdateIterator{contract: _Rollup.contract, event: "RollupDelayPeriodUpdate", logs: logs, sub: sub}, nil
}

// WatchRollupDelayPeriodUpdate is a free log subscription operation binding the contract event 0x2392c64c3c2ac54ae8093f1b546601e09b9c3ee6086d2f2595db2d3c54f3f56b.
//
// Solidity: event RollupDelayPeriodUpdate(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) WatchRollupDelayPeriodUpdate(opts *bind.WatchOpts, sink chan<- *RollupRollupDelayPeriodUpdate) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "RollupDelayPeriodUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRollupDelayPeriodUpdate)
				if err := _Rollup.contract.UnpackLog(event, "RollupDelayPeriodUpdate", log); err != nil {
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

// ParseRollupDelayPeriodUpdate is a log parse operation binding the contract event 0x2392c64c3c2ac54ae8093f1b546601e09b9c3ee6086d2f2595db2d3c54f3f56b.
//
// Solidity: event RollupDelayPeriodUpdate(uint256 oldPeriod, uint256 newPeriod)
func (_Rollup *RollupFilterer) ParseRollupDelayPeriodUpdate(log types.Log) (*RollupRollupDelayPeriodUpdate, error) {
	event := new(RollupRollupDelayPeriodUpdate)
	if err := _Rollup.contract.UnpackLog(event, "RollupDelayPeriodUpdate", log); err != nil {
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
