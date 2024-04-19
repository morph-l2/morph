// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
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
)

// IRollupBatchData is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchData struct {
	Version                uint8
	ParentBatchHeader      []byte
	Chunks                 [][]byte
	SkippedL1MessageBitmap []byte
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	SignatureData          IRollupBatchSignatureData
}

// IRollupBatchSignature is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchSignature struct {
	BlsMsgHash             [32]byte
	SequencerSetVerifyHash [32]byte
	SignedSequencers       []common.Address
}

// IRollupBatchSignatureData is an auto generated low-level Go binding around an user-defined struct.
type IRollupBatchSignatureData struct {
	SignedSequencers []common.Address
	SequencerSets    []byte
	Signature        []byte
}

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_chainID\",\"type\":\"uint64\"},{\"internalType\":\"addresspayable\",\"name\":\"_messenger\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"}],\"name\":\"ChallengeRes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"}],\"name\":\"ChallengeState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateChallenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdateFinalizationPeriodSeconds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxNumTxInChunk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxNumTxInChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldWindow\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWindow\",\"type\":\"uint256\"}],\"name\":\"UpdateProofWindow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateProver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"UpdateVerifier\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_STAKING_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LAYER_2_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROOF_WINDOW\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"batchChallengeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchChallengedSuccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"batchInsideChallengeWindow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"}],\"name\":\"challengeState\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challengeDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"challengeSuccess\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"signedSequencers\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"sequencerSets\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRollup.BatchSignatureData\",\"name\":\"signatureData\",\"type\":\"tuple\"}],\"internalType\":\"structIRollup.BatchData\",\"name\":\"batchData\",\"type\":\"tuple\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committedBatchStores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchVersion\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"originTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalizeTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawalRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1DataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"l1MessagePopped\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalL1MessagePopped\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blobVersionedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blsMsgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sequencerSetVerifyHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"signedSequencers\",\"type\":\"address[]\"}],\"internalType\":\"structIRollup.BatchSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawalRoot\",\"type\":\"bytes32\"}],\"name\":\"importGenesisBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inChallenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1StakingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_finalizationPeriodSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofWindow\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isChallenger\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isProver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCommittedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNumTxInChunk\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_aggrProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_kzgDataProof\",\"type\":\"bytes\"}],\"name\":\"proveState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeChallenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeProver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revertReqIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPeriod\",\"type\":\"uint256\"}],\"name\":\"updateFinalizePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"updateMaxNumTxInChunk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newWindow\",\"type\":\"uint256\"}],\"name\":\"updateProofWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"updateVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawalRoots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b506040516200444838038062004448833981016040819052620000339162000053565b6001600160401b039091166080526001600160a01b031660a052620000a4565b5f806040838503121562000065575f80fd5b82516001600160401b03811681146200007c575f80fd5b60208401519092506001600160a01b038116811462000099575f80fd5b809150509250929050565b60805160a05161437b620000cd5f395f61086201525f8181610647015261329b015261437b5ff3fe6080604052600436106102fb575f3560e01c8063715018a611610191578063b31a77d3116100dc578063e3fff1dd11610087578063f2fde38b11610062578063f2fde38b146109f5578063f4daa29114610a14578063f89db0b414610a29575f80fd5b8063e3fff1dd146109ac578063eb1ec18f146109cb578063ef6602ba146109e0575f80fd5b8063d279c191116100b7578063d279c1911461094f578063de8b30351461096e578063e33491a71461098d575f80fd5b8063b31a77d3146108fc578063b571d3dd14610911578063bedb86fb14610930575f80fd5b80638f1d37761161013c57806397fc007c1161011757806397fc007c14610884578063a415d8dc146108a3578063abc8d68d146108d1575f80fd5b80638f1d377614610756578063910129d414610820578063927ede2d14610851575f80fd5b806388b1ea091161016c57806388b1ea09146107005780638d644bb7146107195780638da5cb5b1461072c575f80fd5b8063715018a6146106a1578063728cdbca146106b5578063881cbd3e146106d4575f80fd5b806321e2f9e0116102515780633e9e82ca116101fc5780635c975abb116101d75780635c975abb1461061f5780635f77cf1d146106365780636c578c1d14610682575f80fd5b80633e9e82ca146105cc5780634c4b9e4f146105e157806357e0af6c14610600575f80fd5b80632b7ac3f31161022c5780632b7ac3f31461053057806336622a30146105815780633b70c18a146105a0575f80fd5b806321e2f9e0146104b85780632362f03e146104d75780632571098d14610505575f80fd5b806310d44583116102b157806318af3b2b1161028c57806318af3b2b1461044a5780631d49e4571461047a5780631e22830214610499575f80fd5b806310d44583146103f4578063116a1f4214610413578063121dcd5014610435575f80fd5b80630a245924116102e15780630a2459241461036c5780630b79cdda1461039a5780630ceb6780146103d3575f80fd5b806304d7721514610306578063059def6114610349575f80fd5b3661030257005b5f80fd5b348015610311575f80fd5b50610334610320366004613b11565b60a26020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b348015610354575f80fd5b5061035e609d5481565b604051908152602001610340565b348015610377575f80fd5b50610334610386366004613b50565b609b6020525f908152604090205460ff1681565b3480156103a5575f80fd5b506103b96103b4366004613b11565b610a3c565b6040516103409e9d9c9b9a99989796959493929190613c09565b3480156103de575f80fd5b506103f26103ed366004613b50565b610bb4565b005b3480156103ff575f80fd5b506103f261040e366004613d15565b610c26565b34801561041e575f80fd5b5061033461042d366004613b11565b609d54101590565b348015610440575f80fd5b5061035e609e5481565b348015610455575f80fd5b50610334610464366004613b11565b5f90815260a16020526040902060030154421090565b348015610485575f80fd5b506103f2610494366004613b50565b610fa2565b3480156104a4575f80fd5b506103f26104b3366004613b11565b611072565b3480156104c3575f80fd5b506103346104d2366004613b11565b6110bf565b3480156104e2575f80fd5b5061035e6104f1366004613b11565b5f90815260a1602052604090206001015490565b348015610510575f80fd5b5061035e61051f366004613b11565b609f6020525f908152604090205481565b34801561053b575f80fd5b50609a5461055c9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610340565b34801561058c575f80fd5b506103f261059b366004613d74565b6110f3565b3480156105ab575f80fd5b5060995461055c9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156105d7575f80fd5b5061035e60a05481565b3480156105ec575f80fd5b506103f26105fb366004613def565b6113d0565b34801561060b575f80fd5b506103f261061a366004613b11565b611a01565b34801561062a575f80fd5b5060655460ff16610334565b348015610641575f80fd5b506106697f000000000000000000000000000000000000000000000000000000000000000081565b60405167ffffffffffffffff9091168152602001610340565b34801561068d575f80fd5b506103f261069c366004613b50565b611a4a565b3480156106ac575f80fd5b506103f2611aad565b3480156106c0575f80fd5b506103f26106cf366004613e3c565b611ac0565b3480156106df575f80fd5b5060975461055c9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561070b575f80fd5b5060a7546103349060ff1681565b6103f2610727366004613e97565b611dc5565b348015610737575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661055c565b348015610761575f80fd5b506107cf610770366004613b11565b60a56020525f9081526040902080546001820154600283015460039093015467ffffffffffffffff8316936801000000000000000090930473ffffffffffffffffffffffffffffffffffffffff16929060ff8082169161010090041686565b6040805167ffffffffffffffff909716875273ffffffffffffffffffffffffffffffffffffffff909516602087015293850192909252606084015215156080830152151560a082015260c001610340565b34801561082b575f80fd5b5061033461083a366004613b11565b5f90815260a5602052604090206003015460ff1690565b34801561085c575f80fd5b5061055c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561088f575f80fd5b506103f261089e366004613b50565b612395565b3480156108ae575f80fd5b506103346108bd366004613b50565b609c6020525f908152604090205460ff1681565b3480156108dc575f80fd5b5061035e6108eb366004613b50565b60a66020525f908152604090205481565b348015610907575f80fd5b5061035e60a85481565b34801561091c575f80fd5b506103f261092b366004613b50565b612413565b34801561093b575f80fd5b506103f261094a366004613ebd565b612476565b34801561095a575f80fd5b506103f2610969366004613b50565b612497565b348015610979575f80fd5b50610334610988366004613b11565b612513565b348015610998575f80fd5b506103f26109a7366004613b11565b61256a565b3480156109b7575f80fd5b506103f26109c6366004613b11565b612bbb565b3480156109d6575f80fd5b5061035e60a45481565b3480156109eb575f80fd5b5061035e60985481565b348015610a00575f80fd5b506103f2610a0f366004613b50565b612c04565b348015610a1f575f80fd5b5061035e60a35481565b6103f2610a37366004613ed8565b612c9e565b60a1602052805f5260405f205f91509050805f01549080600101549080600201549080600301549080600401549080600501549080600601549080600701549080600801549080600901549080600a018054610a9790613f10565b80601f0160208091040260200160405190810160405280929190818152602001828054610ac390613f10565b8015610b0e5780601f10610ae557610100808354040283529160200191610b0e565b820191905f5260205f20905b815481529060010190602001808311610af157829003601f168201915b50505050509080600b01549080600c01549080600d016040518060600160405290815f82015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015610ba657602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610b7b575b50505050508152505090508e565b610bbc612ded565b73ffffffffffffffffffffffffffffffffffffffff81165f818152609c6020908152604091829020805460ff1916600190811790915591519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc00991015b60405180910390a250565b610c2e612ded565b5f8111610c825760405162461bcd60e51b815260206004820152601560248201527f636f756e74206d757374206265206e6f6e7a65726f000000000000000000000060448201526064015b60405180910390fd5b5f80610c8e8585612e54565b915091505f610ca1836001015160c01c90565b5f81815260a160205260409020600101549091508214610d035760405162461bcd60e51b815260206004820152601460248201527f696e636f727265637420626174636820686173680000000000000000000000006044820152606401610c79565b5f60a181610d118785613f8e565b81526020019081526020015f206001015414610d945760405162461bcd60e51b8152602060048201526024808201527f726576657274696e67206d7573742073746172742066726f6d2074686520656e60448201527f64696e67000000000000000000000000000000000000000000000000000000006064820152608401610c79565b609d548111610e0b5760405162461bcd60e51b815260206004820152602160248201527f63616e206f6e6c792072657665727420756e46696e616c697a6564206261746360448201527f68000000000000000000000000000000000000000000000000000000000000006064820152608401610c79565b610e16600182613fa1565b609e555b8315610f9a575f81815260a1602052604081206001015560a85415801590610e43575060a85481145b15610f2c575f81815260a56020526040902060030154610100900460ff16610ebe575f81815260a5602090815260408083206001810154905468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16845260a69092528220805491929091610eb8908490613f8e565b90915550505b5f81815260a56020526040812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600181018290556002810182905560030180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905560a8555b604051829082907ecae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146905f90a360019081015f81815260a160205260409020909101547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90940193915081610e1a575b505050505050565b610faa612ded565b73ffffffffffffffffffffffffffffffffffffffff81163b1561100f5760405162461bcd60e51b815260206004820152600760248201527f6e6f7420454f41000000000000000000000000000000000000000000000000006044820152606401610c79565b73ffffffffffffffffffffffffffffffffffffffff81165f818152609b6020908152604091829020805460ff1916600190811790915591519182527f967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e9101610c1b565b61107a612ded565b609880549082905560408051828152602081018490527f6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a910160405180910390a15050565b5f81815260a16020526040812060020154158015906110ed57505f82815260a1602052604090206001015415155b92915050565b60a854156111435760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610c79565b67ffffffffffffffff85165f90815260a5602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff166111ca5760405162461bcd60e51b815260206004820152601860248201527f4368616c6c656e676520646f6573206e6f7420657869737400000000000000006044820152606401610c79565b67ffffffffffffffff85165f90815260a56020526040902060030154610100900460ff161561123b5760405162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520616c72656164792066696e69736865640000000000006044820152606401610c79565b67ffffffffffffffff85165f90815260a5602052604090206003810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1661010017905560a7805460ff1916905560a454600290910154429161129f91613f8e565b1161137c5767ffffffffffffffff85165f90815260a560209081526040808320600301805460ff1916600117905560a1825291829020600f0180548351818402810184019094528084526113779389939092919083018282801561133757602002820191905f5260205f20905b815473ffffffffffffffffffffffffffffffffffffffff16815260019091019060200180831161130c575b50505050506040518060400160405280600781526020017f54696d656f757400000000000000000000000000000000000000000000000000815250612e6f565b6113c9565b6113898585858585613021565b6113c985336040518060400160405280600d81526020017f50726f6f662073756363657373000000000000000000000000000000000000008152506133e0565b5050505050565b8161141d5760405162461bcd60e51b815260206004820152600f60248201527f7a65726f20737461746520726f6f7400000000000000000000000000000000006044820152606401610c79565b5f8052609f6020527fa705961f203609058950cfd817eb7a7627c9e270651c936aad3abdfa253727ec54156114945760405162461bcd60e51b815260206004820152601660248201527f67656e6573697320626174636820696d706f72746564000000000000000000006044820152606401610c79565b5f806114a08686612e54565b915091505f6114b0836019015190565b90505f6114be845160f81c90565b90505f6114cf856011015160c01c90565b600986015160c01c600187015160c01c875160f81c010101905080156115375760405162461bcd60e51b815260206004820152601760248201527f6e6f7420616c6c206669656c647320617265207a65726f0000000000000000006044820152606401610c79565b505f611544856019015190565b036115915760405162461bcd60e51b815260206004820152600e60248201527f7a65726f206461746120686173680000000000000000000000000000000000006044820152606401610c79565b5f61159d856059015190565b146115ea5760405162461bcd60e51b815260206004820152601960248201527f6e6f6e7a65726f20706172656e742062617463682068617368000000000000006044820152606401610c79565b7f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014611616856039015190565b146116635760405162461bcd60e51b815260206004820152601660248201527f696e76616c69642076657273696f6e65642068617368000000000000000000006044820152606401610c79565b604051806101c001604052808281526020018481526020014281526020014281526020015f801b81526020018781526020018681526020018381526020015f81526020015f815260200160405180602001604052805f81525081526020015f81526020017f010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c4440145f1b815260200160405180606001604052805f801b81526020015f801b81526020015f67ffffffffffffffff81111561172357611723613fb4565b60405190808252806020026020018201604052801561174c578160200160208202803683370190505b50905290525f805260a1602090815281517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f882908155908201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8835560408201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8845560608201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8855560808201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8865560a08201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8875560c08201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f8885560e08201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f889556101008201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88a556101208201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88b556101408201517f32ae1b88a7d4f92d7e214b63c8ea04cd13e2faaa60c50f499f2254336d98f88c906119139082614025565b50610160820151600b820155610180820151600c8201556101a08201518051600d8301908155602080830151600e8501556040830151805161195b92600f8701920190613a27565b50505f808052609f6020527fa705961f203609058950cfd817eb7a7627c9e270651c936aad3abdfa253727ec8a90556040518794509092507f2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f91508290a3604080518781525f60208201819052859290917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a35050505050505050565b611a09612ded565b60a45460408051918252602082018390527f1e3a2094feb4b696dd3d7caea38ad2f41dbdcac3fa3943c7a693aff8a64b0a61910160405180910390a160a455565b611a52612ded565b73ffffffffffffffffffffffffffffffffffffffff81165f818152609c60209081526040808320805460ff19169055519182527f7feb653c7b1f0d23daeed494225b3f28851cdc8973fcc653866d9b6e205fc0099101610c1b565b611ab5612ded565b611abe5f61349c565b565b5f54610100900460ff1615808015611ade57505f54600160ff909116105b80611af75750303b158015611af757505f5460ff166001145b611b695760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610c79565b5f805460ff191660011790558015611ba7575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b73ffffffffffffffffffffffffffffffffffffffff86161580611bde575073ffffffffffffffffffffffffffffffffffffffff8516155b15611c15576040517fecc6fdf000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c1d613512565b611c25613596565b73ffffffffffffffffffffffffffffffffffffffff8716611c885760405162461bcd60e51b815260206004820152601b60248201527f696e76616c6964206c31207374616b696e6720636f6e747261637400000000006044820152606401610c79565b6097805473ffffffffffffffffffffffffffffffffffffffff808a167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560a385905560a484905560998054898416908316179055609a8054928816929091168217905560988590556040515f907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96908290a3604080515f8152602081018690527f6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a910160405180910390a18015611dbc575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b335f908152609c602052604090205460ff16611e235760405162461bcd60e51b815260206004820152601560248201527f63616c6c6572206e6f74206368616c6c656e67657200000000000000000000006044820152606401610c79565b60a85415611e735760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610c79565b60a75460ff1615611ec65760405162461bcd60e51b815260206004820152601460248201527f616c726561647920696e206368616c6c656e67650000000000000000000000006044820152606401610c79565b8067ffffffffffffffff16609d5410611f215760405162461bcd60e51b815260206004820152601760248201527f626174636820616c72656164792066696e616c697a65640000000000000000006044820152606401610c79565b67ffffffffffffffff81165f90815260a160205260408120600101549003611f8b5760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610c79565b67ffffffffffffffff81165f90815260a5602052604090205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16156120135760405162461bcd60e51b815260206004820152601260248201527f616c7265616479206368616c6c656e67656400000000000000000000000000006044820152606401610c79565b67ffffffffffffffff81165f90815260a1602052604090206003015442106120a35760405162461bcd60e51b815260206004820152603360248201527f63616e6e6f74206368616c6c656e6765206261746368206f757473696465207460448201527f6865206368616c6c656e67652077696e646f77000000000000000000000000006064820152608401610c79565b60975f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302df7ff76040518163ffffffff1660e01b81526004016020604051808303815f875af115801561210e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612132919061413d565b3410156121815760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742076616c756500000000000000000000000000006044820152606401610c79565b6040518060c001604052808267ffffffffffffffff1681526020016121a33390565b73ffffffffffffffffffffffffffffffffffffffff908116825234602080840191909152426040808501919091525f6060808601829052608095860182905267ffffffffffffffff808916835260a5855291839020875181549589015190961668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090951695909216949094179290921782558401516001820155908301516002820155908201516003909101805460a0909301511515610100027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff921515929092167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000909316929092171790556122c23390565b73ffffffffffffffffffffffffffffffffffffffff168167ffffffffffffffff167f3a6ea19df25b49e7624e313ce7c1ab23984238e93727260db56a81735b1b99763460405161231491815260200190565b60405180910390a35f609d54600161232c9190613f8e565b90505b609e548111612384578167ffffffffffffffff1681146123725760a4545f82815260a160205260408120600301805490919061236c908490613f8e565b90915550505b8061237c81614154565b91505061232f565b505060a7805460ff19166001179055565b61239d612ded565b609a805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f728af3d16a5760405e27a082c98ab272e9f0a1d02f0085d41532a26093aedd96905f90a35050565b61241b612ded565b73ffffffffffffffffffffffffffffffffffffffff81165f818152609b60209081526040808320805460ff19169055519182527f967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e9101610c1b565b61247e612ded565b801561248f5761248c61361a565b50565b61248c613681565b335f90815260a66020526040812054908190036124f65760405162461bcd60e51b815260206004820152601c60248201527f696e76616c69642062617463684368616c6c656e6765526577617264000000006044820152606401610c79565b335f90815260a6602052604081205561250f82826136ba565b5050565b5f81815260a5602052604081205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16158015906110ed5750505f90815260a56020526040902060030154610100900460ff161590565b60a854156125ba5760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610c79565b6125c261379e565b6125cb816110bf565b6126175760405162461bcd60e51b815260206004820152600f60248201527f6261746368206e6f7420657869737400000000000000000000000000000000006044820152606401610c79565b61262081612513565b1561266d5760405162461bcd60e51b815260206004820152601260248201527f626174636820696e206368616c6c656e676500000000000000000000000000006044820152606401610c79565b5f81815260a5602052604090206003015460ff16156126ce5760405162461bcd60e51b815260206004820152601660248201527f62617463682073686f756c6420626520726576657274000000000000000000006044820152606401610c79565b5f81815260a1602052604090206003015442101561272e5760405162461bcd60e51b815260206004820152601960248201527f626174636820696e206368616c6c656e67652077696e646f77000000000000006044820152606401610c79565b5f81815260a1602052604081206004015490609f9061274e600185613fa1565b81526020019081526020015f2054146127a95760405162461bcd60e51b815260206004820152601d60248201527f696e636f72726563742070726576696f757320737461746520726f6f740000006044820152606401610c79565b5f818152609f6020526040902054156128045760405162461bcd60e51b815260206004820152601660248201527f626174636820616c7265616479207665726966696564000000000000000000006044820152606401610c79565b80609d54600101146128585760405162461bcd60e51b815260206004820152601560248201527f696e636f727265637420626174636820696e64657800000000000000000000006044820152606401610c79565b609d8190555f81815260a1602081815260408084206006810154855260a28352818520805460ff191660011790558585526005810154609f8452919094205552600801548015612a42576099545f83815260a160205260408120600a01805473ffffffffffffffffffffffffffffffffffffffff909316926128d990613f10565b80601f016020809104026020016040519081016040528092919081815260200182805461290590613f10565b80156129505780601f1061292757610100808354040283529160200191612950565b820191905f5260205f20905b81548152906001019060200180831161293357829003601f168201915b5050505f87815260a160209081526040822060090154949550850193879003925090505b85811015612a3c5761010081870381111561298e57508086035b6101008204602081028501516040517f55f613ce00000000000000000000000000000000000000000000000000000000815260048101869052602481018490526044810182905290919073ffffffffffffffffffffffffffffffffffffffff8916906355f613ce906064015f604051808303815f87803b158015612a10575f80fd5b505af1158015612a22573d5f803e3d5ffd5b505050506101008501945050505061010081019050612974565b50505050505b60a15f612a50600185613fa1565b815260208101919091526040015f90812081815560018101829055600281018290556003810182905560048101829055600581018290556006810182905560078101829055600881018290556009810182905590612ab1600a830182613aaf565b5f600b8301819055600c8301819055600d8301818155600e840182905590612adc600f850182613ae6565b5050505060a55f600184612af09190613fa1565b815260208082019290925260409081015f90812080547fffffffff00000000000000000000000000000000000000000000000000000000168155600180820183905560028201839055600390910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000016905585825260a184529082902090810154600582015460069092015483519283529382019390935284917f26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d910160405180910390a35050565b612bc3612ded565b60a35460408051918252602082018390527fa577f4223f91f74e2dad65bbb8c30807587ae95d0d34288057bb3ec0d398a437910160405180910390a160a355565b612c0c612ded565b73ffffffffffffffffffffffffffffffffffffffff8116612c955760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610c79565b61248c8161349c565b60975473ffffffffffffffffffffffffffffffffffffffff16636f1e8533336040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016020604051808303815f875af1158015612d25573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d49919061418b565b612d955760405162461bcd60e51b815260206004820152601460248201527f63616c6c6572206e6f742073657175656e6365720000000000000000000000006044820152606401610c79565b60a85415612de55760405162461bcd60e51b815260206004820152600b60248201527f6e656564207265766572740000000000000000000000000000000000000000006044820152606401610c79565b61248c61379e565b60335473ffffffffffffffffffffffffffffffffffffffff163314611abe5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c79565b5f805f612e6185856137f1565b812090969095509350505050565b67ffffffffffffffff831660a88190555f90815260a560205260408082205460975491517f8b8c24c10000000000000000000000000000000000000000000000000000000081526801000000000000000090910473ffffffffffffffffffffffffffffffffffffffff90811693921690638b8c24c190612ef39087906004016141a6565b6020604051808303815f875af1158015612f0f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f33919061413d565b67ffffffffffffffff86165f90815260a56020526040902060010154909150612f5d908290613f8e565b67ffffffffffffffff86165f90815260a5602090815260408083205468010000000000000000900473ffffffffffffffffffffffffffffffffffffffff16835260a690915281208054909190612fb4908490613f8e565b9091555050604051612fc79084906141ff565b6040519081900381209073ffffffffffffffffffffffffffffffffffffffff84169067ffffffffffffffff8816907f1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6905f90a45050505050565b8261306e5760405162461bcd60e51b815260206004820152601960248201527f496e76616c6964206167677265676174696f6e2070726f6f66000000000000006044820152606401610c79565b60a081146130be5760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964204b5a4720646174612070726f6f66000000000000000000006044820152606401610c79565b67ffffffffffffffff85165f90815260a160209081526040808320600c015490518392600a926130f4929091889188910161421a565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261312c916141ff565b5f60405180830381855afa9150503d805f8114613164576040519150601f19603f3d011682016040523d82523d5f602084013e613169565b606091505b5091509150816131e15760405162461bcd60e51b815260206004820152602a60248201527f6661696c656420746f2063616c6c20706f696e74206576616c756174696f6e2060448201527f707265636f6d70696c65000000000000000000000000000000000000000000006064820152608401610c79565b5f818060200190518101906131f69190614233565b9150507f73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff0000000181146132685760405162461bcd60e51b815260206004820152601c60248201527f707265636f6d70696c6520756e6578706563746564206f7574707574000000006044820152606401610c79565b50505067ffffffffffffffff85165f90815260a160205260408082206004810154600582015460068301546007909301547f000000000000000000000000000000000000000000000000000000000000000094929391926132cb9087898b614255565b60a15f8e67ffffffffffffffff1681526020019081526020015f20600c015460405160200161330198979695949392919061427c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181528282528051602091820120609a5467ffffffffffffffff8b165f90815260a190935292909120547f2c09a84800000000000000000000000000000000000000000000000000000000845290935073ffffffffffffffffffffffffffffffffffffffff90911691632c09a848916133ac918a908a908a9088906004016142d8565b5f6040518083038186803b1580156133c2575f80fd5b505afa1580156133d4573d5f803e3d5ffd5b50505050505050505050565b67ffffffffffffffff83165f90815260a5602090815260408083206001015473ffffffffffffffffffffffffffffffffffffffff8616845260a69092528220805491928392613430908490613f8e565b90915550506040516134439083906141ff565b6040519081900381209073ffffffffffffffffffffffffffffffffffffffff85169067ffffffffffffffff8716907f1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6905f90a450505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff1661358e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c79565b611abe6138c8565b5f54610100900460ff166136125760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c79565b611abe613950565b61362261379e565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586136573390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6136896139d5565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33613657565b801561250f575f8273ffffffffffffffffffffffffffffffffffffffff1682604051613709907f3078000000000000000000000000000000000000000000000000000000000000815260020190565b5f6040518083038185875af1925050503d805f8114613743576040519150601f19603f3d011682016040523d82523d5f602084013e613748565b606091505b50509050806137995760405162461bcd60e51b815260206004820152601b60248201527f526f6c6c75703a20455448207472616e73666572206661696c656400000000006044820152606401610c79565b505050565b60655460ff1615611abe5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610c79565b5f8160798110156138445760405162461bcd60e51b815260206004820152601d60248201527f626174636820686561646572206c656e67746820746f6f20736d616c6c0000006044820152606401610c79565b6040519150808483378082016040525f613862836009015160c01c90565b905061010060ff82010460200260790182146138c05760405162461bcd60e51b815260206004820152601360248201527f77726f6e67206269746d6170206c656e677468000000000000000000000000006044820152606401610c79565b509250929050565b5f54610100900460ff166139445760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c79565b6065805460ff19169055565b5f54610100900460ff166139cc5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c79565b611abe3361349c565b60655460ff16611abe5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c79565b828054828255905f5260205f20908101928215613a9f579160200282015b82811115613a9f57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190613a45565b50613aab929150613afd565b5090565b508054613abb90613f10565b5f825580601f10613aca575050565b601f0160209004905f5260205f209081019061248c9190613afd565b5080545f8255905f5260205f209081019061248c91905b5b80821115613aab575f8155600101613afe565b5f60208284031215613b21575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114613b4b575f80fd5b919050565b5f60208284031215613b60575f80fd5b613b6982613b28565b9392505050565b5f5b83811015613b8a578181015183820152602001613b72565b50505f910152565b5f606083018251845260208084015160208601526040840151606060408701528281518085526080880191506020830194505f92505b80831015613bfe57845173ffffffffffffffffffffffffffffffffffffffff168252938301936001929092019190830190613bc8565b509695505050505050565b8e81528d60208201528c60408201528b60608201528a60808201528960a08201528860c08201528760e082015286610100820152856101208201526101c06101408201525f8551806101c08401526101e0613c6a8282860160208b01613b70565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684019150866101608501528561018085015280848303016101a0850152613cb981830186613b92565b925050509f9e505050505050505050505050505050565b5f8083601f840112613ce0575f80fd5b50813567ffffffffffffffff811115613cf7575f80fd5b602083019150836020828501011115613d0e575f80fd5b9250929050565b5f805f60408486031215613d27575f80fd5b833567ffffffffffffffff811115613d3d575f80fd5b613d4986828701613cd0565b909790965060209590950135949350505050565b803567ffffffffffffffff81168114613b4b575f80fd5b5f805f805f60608688031215613d88575f80fd5b613d9186613d5d565b9450602086013567ffffffffffffffff80821115613dad575f80fd5b613db989838a01613cd0565b90965094506040880135915080821115613dd1575f80fd5b50613dde88828901613cd0565b969995985093965092949392505050565b5f805f8060608587031215613e02575f80fd5b843567ffffffffffffffff811115613e18575f80fd5b613e2487828801613cd0565b90989097506020870135966040013595509350505050565b5f805f805f8060c08789031215613e51575f80fd5b613e5a87613b28565b9550613e6860208801613b28565b9450613e7660408801613b28565b9350606087013592506080870135915060a087013590509295509295509295565b5f60208284031215613ea7575f80fd5b613b6982613d5d565b801515811461248c575f80fd5b5f60208284031215613ecd575f80fd5b8135613b6981613eb0565b5f60208284031215613ee8575f80fd5b813567ffffffffffffffff811115613efe575f80fd5b82016101008185031215613b69575f80fd5b600181811c90821680613f2457607f821691505b602082108103613f5b577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b808201808211156110ed576110ed613f61565b818103818111156110ed576110ed613f61565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b601f82111561379957805f5260205f20601f840160051c810160208510156140065750805b601f840160051c820191505b818110156113c9575f8155600101614012565b815167ffffffffffffffff81111561403f5761403f613fb4565b6140538161404d8454613f10565b84613fe1565b602080601f8311600181146140a5575f841561406f5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610f9a565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156140f1578886015182559484019460019091019084016140d2565b508582101561412d57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b5f6020828403121561414d575f80fd5b5051919050565b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361418457614184613f61565b5060010190565b5f6020828403121561419b575f80fd5b8151613b6981613eb0565b602080825282518282018190525f9190848201906040850190845b818110156141f357835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016141c1565b50909695505050505050565b5f8251614210818460208701613b70565b9190910192915050565b838152818360208301375f910160200190815292915050565b5f8060408385031215614244575f80fd5b505080516020909101519092909150565b5f8085851115614263575f80fd5b8386111561426f575f80fd5b5050820193919092039150565b7fffffffffffffffff0000000000000000000000000000000000000000000000008960c01b16815287600882015286602882015285604882015284606882015282846088830137608892019182015260a8019695505050505050565b85815267ffffffffffffffff8516602082015260806040820152826080820152828460a08301375f60a084830101525f60a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8601168301019050826060830152969550505050505056fea2646970667358221220b4a0dbfda1f4e8869ff7006b102b3f398550dd1b74cf2395d06180a22f11b4b064736f6c63430008180033",
}

// RollupABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMetaData.ABI instead.
var RollupABI = RollupMetaData.ABI

// RollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMetaData.Bin instead.
var RollupBin = RollupMetaData.Bin

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _chainID uint64, _messenger common.Address) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupBin), backend, _chainID, _messenger)
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
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _Rollup.Contract.FINALIZATIONPERIODSECONDS(&_Rollup.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_Rollup *RollupCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _Rollup.Contract.FINALIZATIONPERIODSECONDS(&_Rollup.CallOpts)
}

// L1STAKINGCONTRACT is a free data retrieval call binding the contract method 0x881cbd3e.
//
// Solidity: function L1_STAKING_CONTRACT() view returns(address)
func (_Rollup *RollupCaller) L1STAKINGCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "L1_STAKING_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1STAKINGCONTRACT is a free data retrieval call binding the contract method 0x881cbd3e.
//
// Solidity: function L1_STAKING_CONTRACT() view returns(address)
func (_Rollup *RollupSession) L1STAKINGCONTRACT() (common.Address, error) {
	return _Rollup.Contract.L1STAKINGCONTRACT(&_Rollup.CallOpts)
}

// L1STAKINGCONTRACT is a free data retrieval call binding the contract method 0x881cbd3e.
//
// Solidity: function L1_STAKING_CONTRACT() view returns(address)
func (_Rollup *RollupCallerSession) L1STAKINGCONTRACT() (common.Address, error) {
	return _Rollup.Contract.L1STAKINGCONTRACT(&_Rollup.CallOpts)
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

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupSession) MESSENGER() (common.Address, error) {
	return _Rollup.Contract.MESSENGER(&_Rollup.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_Rollup *RollupCallerSession) MESSENGER() (common.Address, error) {
	return _Rollup.Contract.MESSENGER(&_Rollup.CallOpts)
}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupCaller) PROOFWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "PROOF_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupSession) PROOFWINDOW() (*big.Int, error) {
	return _Rollup.Contract.PROOFWINDOW(&_Rollup.CallOpts)
}

// PROOFWINDOW is a free data retrieval call binding the contract method 0xeb1ec18f.
//
// Solidity: function PROOF_WINDOW() view returns(uint256)
func (_Rollup *RollupCallerSession) PROOFWINDOW() (*big.Int, error) {
	return _Rollup.Contract.PROOFWINDOW(&_Rollup.CallOpts)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupCaller) BatchChallengeReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "batchChallengeReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupSession) BatchChallengeReward(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, arg0)
}

// BatchChallengeReward is a free data retrieval call binding the contract method 0xabc8d68d.
//
// Solidity: function batchChallengeReward(address ) view returns(uint256)
func (_Rollup *RollupCallerSession) BatchChallengeReward(arg0 common.Address) (*big.Int, error) {
	return _Rollup.Contract.BatchChallengeReward(&_Rollup.CallOpts, arg0)
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
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCaller) Challenges(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "challenges", arg0)

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
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupSession) Challenges(arg0 *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x8f1d3776.
//
// Solidity: function challenges(uint256 ) view returns(uint64 batchIndex, address challenger, uint256 challengeDeposit, uint256 startTime, bool challengeSuccess, bool finished)
func (_Rollup *RollupCallerSession) Challenges(arg0 *big.Int) (struct {
	BatchIndex       uint64
	Challenger       common.Address
	ChallengeDeposit *big.Int
	StartTime        *big.Int
	ChallengeSuccess bool
	Finished         bool
}, error) {
	return _Rollup.Contract.Challenges(&_Rollup.CallOpts, arg0)
}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash, (bytes32,bytes32,address[]) signature)
func (_Rollup *RollupCaller) CommittedBatchStores(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
	Signature              IRollupBatchSignature
}, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "committedBatchStores", arg0)

	outstruct := new(struct {
		BatchVersion           *big.Int
		BatchHash              [32]byte
		OriginTimestamp        *big.Int
		FinalizeTimestamp      *big.Int
		PrevStateRoot          [32]byte
		PostStateRoot          [32]byte
		WithdrawalRoot         [32]byte
		L1DataHash             [32]byte
		L1MessagePopped        *big.Int
		TotalL1MessagePopped   *big.Int
		SkippedL1MessageBitmap []byte
		BlockNumber            *big.Int
		BlobVersionedHash      [32]byte
		Signature              IRollupBatchSignature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchVersion = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BatchHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.OriginTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FinalizeTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PrevStateRoot = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)
	outstruct.PostStateRoot = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.WithdrawalRoot = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.L1DataHash = *abi.ConvertType(out[7], new([32]byte)).(*[32]byte)
	outstruct.L1MessagePopped = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.TotalL1MessagePopped = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.SkippedL1MessageBitmap = *abi.ConvertType(out[10], new([]byte)).(*[]byte)
	outstruct.BlockNumber = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)
	outstruct.BlobVersionedHash = *abi.ConvertType(out[12], new([32]byte)).(*[32]byte)
	outstruct.Signature = *abi.ConvertType(out[13], new(IRollupBatchSignature)).(*IRollupBatchSignature)

	return *outstruct, err

}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash, (bytes32,bytes32,address[]) signature)
func (_Rollup *RollupSession) CommittedBatchStores(arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
	Signature              IRollupBatchSignature
}, error) {
	return _Rollup.Contract.CommittedBatchStores(&_Rollup.CallOpts, arg0)
}

// CommittedBatchStores is a free data retrieval call binding the contract method 0x0b79cdda.
//
// Solidity: function committedBatchStores(uint256 ) view returns(uint256 batchVersion, bytes32 batchHash, uint256 originTimestamp, uint256 finalizeTimestamp, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawalRoot, bytes32 l1DataHash, uint256 l1MessagePopped, uint256 totalL1MessagePopped, bytes skippedL1MessageBitmap, uint256 blockNumber, bytes32 blobVersionedHash, (bytes32,bytes32,address[]) signature)
func (_Rollup *RollupCallerSession) CommittedBatchStores(arg0 *big.Int) (struct {
	BatchVersion           *big.Int
	BatchHash              [32]byte
	OriginTimestamp        *big.Int
	FinalizeTimestamp      *big.Int
	PrevStateRoot          [32]byte
	PostStateRoot          [32]byte
	WithdrawalRoot         [32]byte
	L1DataHash             [32]byte
	L1MessagePopped        *big.Int
	TotalL1MessagePopped   *big.Int
	SkippedL1MessageBitmap []byte
	BlockNumber            *big.Int
	BlobVersionedHash      [32]byte
	Signature              IRollupBatchSignature
}, error) {
	return _Rollup.Contract.CommittedBatchStores(&_Rollup.CallOpts, arg0)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
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
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_Rollup *RollupSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_Rollup *RollupCallerSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _Rollup.Contract.CommittedBatches(&_Rollup.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupCaller) FinalizedStateRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "finalizedStateRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, arg0)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 ) view returns(bytes32)
func (_Rollup *RollupCallerSession) FinalizedStateRoots(arg0 *big.Int) ([32]byte, error) {
	return _Rollup.Contract.FinalizedStateRoots(&_Rollup.CallOpts, arg0)
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
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupCaller) IsChallenger(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isChallenger", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupSession) IsChallenger(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, arg0)
}

// IsChallenger is a free data retrieval call binding the contract method 0xa415d8dc.
//
// Solidity: function isChallenger(address ) view returns(bool)
func (_Rollup *RollupCallerSession) IsChallenger(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsChallenger(&_Rollup.CallOpts, arg0)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupCaller) IsProver(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "isProver", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupSession) IsProver(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsProver(&_Rollup.CallOpts, arg0)
}

// IsProver is a free data retrieval call binding the contract method 0x0a245924.
//
// Solidity: function isProver(address ) view returns(bool)
func (_Rollup *RollupCallerSession) IsProver(arg0 common.Address) (bool, error) {
	return _Rollup.Contract.IsProver(&_Rollup.CallOpts, arg0)
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

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupCaller) LatestL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "latestL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupSession) LatestL2BlockNumber() (*big.Int, error) {
	return _Rollup.Contract.LatestL2BlockNumber(&_Rollup.CallOpts)
}

// LatestL2BlockNumber is a free data retrieval call binding the contract method 0x3e9e82ca.
//
// Solidity: function latestL2BlockNumber() view returns(uint256)
func (_Rollup *RollupCallerSession) LatestL2BlockNumber() (*big.Int, error) {
	return _Rollup.Contract.LatestL2BlockNumber(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCaller) MaxNumTxInChunk(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "maxNumTxInChunk")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
}

// MaxNumTxInChunk is a free data retrieval call binding the contract method 0xef6602ba.
//
// Solidity: function maxNumTxInChunk() view returns(uint256)
func (_Rollup *RollupCallerSession) MaxNumTxInChunk() (*big.Int, error) {
	return _Rollup.Contract.MaxNumTxInChunk(&_Rollup.CallOpts)
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
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupCaller) WithdrawalRoots(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "withdrawalRoots", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupSession) WithdrawalRoots(arg0 [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, arg0)
}

// WithdrawalRoots is a free data retrieval call binding the contract method 0x04d77215.
//
// Solidity: function withdrawalRoots(bytes32 ) view returns(bool)
func (_Rollup *RollupCallerSession) WithdrawalRoots(arg0 [32]byte) (bool, error) {
	return _Rollup.Contract.WithdrawalRoots(&_Rollup.CallOpts, arg0)
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

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupTransactor) AddProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addProver", _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddProver(&_Rollup.TransactOpts, _account)
}

// AddProver is a paid mutator transaction binding the contract method 0x1d49e457.
//
// Solidity: function addProver(address _account) returns()
func (_Rollup *RollupTransactorSession) AddProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.AddProver(&_Rollup.TransactOpts, _account)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_Rollup *RollupTransactor) ChallengeState(opts *bind.TransactOpts, batchIndex uint64) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "challengeState", batchIndex)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_Rollup *RollupSession) ChallengeState(batchIndex uint64) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex)
}

// ChallengeState is a paid mutator transaction binding the contract method 0x8d644bb7.
//
// Solidity: function challengeState(uint64 batchIndex) payable returns()
func (_Rollup *RollupTransactorSession) ChallengeState(batchIndex uint64) (*types.Transaction, error) {
	return _Rollup.Contract.ChallengeState(&_Rollup.TransactOpts, batchIndex)
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

// CommitBatch is a paid mutator transaction binding the contract method 0xf89db0b4.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(address[],bytes,bytes)) batchData) payable returns()
func (_Rollup *RollupTransactor) CommitBatch(opts *bind.TransactOpts, batchData IRollupBatchData) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "commitBatch", batchData)
}

// CommitBatch is a paid mutator transaction binding the contract method 0xf89db0b4.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(address[],bytes,bytes)) batchData) payable returns()
func (_Rollup *RollupSession) CommitBatch(batchData IRollupBatchData) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchData)
}

// CommitBatch is a paid mutator transaction binding the contract method 0xf89db0b4.
//
// Solidity: function commitBatch((uint8,bytes,bytes[],bytes,bytes32,bytes32,bytes32,(address[],bytes,bytes)) batchData) payable returns()
func (_Rollup *RollupTransactorSession) CommitBatch(batchData IRollupBatchData) (*types.Transaction, error) {
	return _Rollup.Contract.CommitBatch(&_Rollup.TransactOpts, batchData)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupTransactor) FinalizeBatch(opts *bind.TransactOpts, _batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "finalizeBatch", _batchIndex)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupSession) FinalizeBatch(_batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchIndex)
}

// FinalizeBatch is a paid mutator transaction binding the contract method 0xe33491a7.
//
// Solidity: function finalizeBatch(uint256 _batchIndex) returns()
func (_Rollup *RollupTransactorSession) FinalizeBatch(_batchIndex *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.FinalizeBatch(&_Rollup.TransactOpts, _batchIndex)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupTransactor) ImportGenesisBatch(opts *bind.TransactOpts, _batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "importGenesisBatch", _batchHeader, _postStateRoot, _withdrawalRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupSession) ImportGenesisBatch(_batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader, _postStateRoot, _withdrawalRoot)
}

// ImportGenesisBatch is a paid mutator transaction binding the contract method 0x4c4b9e4f.
//
// Solidity: function importGenesisBatch(bytes _batchHeader, bytes32 _postStateRoot, bytes32 _withdrawalRoot) returns()
func (_Rollup *RollupTransactorSession) ImportGenesisBatch(_batchHeader []byte, _postStateRoot [32]byte, _withdrawalRoot [32]byte) (*types.Transaction, error) {
	return _Rollup.Contract.ImportGenesisBatch(&_Rollup.TransactOpts, _batchHeader, _postStateRoot, _withdrawalRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupTransactor) Initialize(opts *bind.TransactOpts, _l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "initialize", _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupSession) Initialize(_l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _l1StakingContract, address _messageQueue, address _verifier, uint256 _maxNumTxInChunk, uint256 _finalizationPeriodSeconds, uint256 _proofWindow) returns()
func (_Rollup *RollupTransactorSession) Initialize(_l1StakingContract common.Address, _messageQueue common.Address, _verifier common.Address, _maxNumTxInChunk *big.Int, _finalizationPeriodSeconds *big.Int, _proofWindow *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.Initialize(&_Rollup.TransactOpts, _l1StakingContract, _messageQueue, _verifier, _maxNumTxInChunk, _finalizationPeriodSeconds, _proofWindow)
}

// ProveState is a paid mutator transaction binding the contract method 0x36622a30.
//
// Solidity: function proveState(uint64 _batchIndex, bytes _aggrProof, bytes _kzgDataProof) returns()
func (_Rollup *RollupTransactor) ProveState(opts *bind.TransactOpts, _batchIndex uint64, _aggrProof []byte, _kzgDataProof []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "proveState", _batchIndex, _aggrProof, _kzgDataProof)
}

// ProveState is a paid mutator transaction binding the contract method 0x36622a30.
//
// Solidity: function proveState(uint64 _batchIndex, bytes _aggrProof, bytes _kzgDataProof) returns()
func (_Rollup *RollupSession) ProveState(_batchIndex uint64, _aggrProof []byte, _kzgDataProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchIndex, _aggrProof, _kzgDataProof)
}

// ProveState is a paid mutator transaction binding the contract method 0x36622a30.
//
// Solidity: function proveState(uint64 _batchIndex, bytes _aggrProof, bytes _kzgDataProof) returns()
func (_Rollup *RollupTransactorSession) ProveState(_batchIndex uint64, _aggrProof []byte, _kzgDataProof []byte) (*types.Transaction, error) {
	return _Rollup.Contract.ProveState(&_Rollup.TransactOpts, _batchIndex, _aggrProof, _kzgDataProof)
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

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupTransactor) RemoveProver(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "removeProver", _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveProver(&_Rollup.TransactOpts, _account)
}

// RemoveProver is a paid mutator transaction binding the contract method 0xb571d3dd.
//
// Solidity: function removeProver(address _account) returns()
func (_Rollup *RollupTransactorSession) RemoveProver(_account common.Address) (*types.Transaction, error) {
	return _Rollup.Contract.RemoveProver(&_Rollup.TransactOpts, _account)
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

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupTransactor) UpdateMaxNumTxInChunk(opts *bind.TransactOpts, _maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "updateMaxNumTxInChunk", _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateMaxNumTxInChunk(&_Rollup.TransactOpts, _maxNumTxInChunk)
}

// UpdateMaxNumTxInChunk is a paid mutator transaction binding the contract method 0x1e228302.
//
// Solidity: function updateMaxNumTxInChunk(uint256 _maxNumTxInChunk) returns()
func (_Rollup *RollupTransactorSession) UpdateMaxNumTxInChunk(_maxNumTxInChunk *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.UpdateMaxNumTxInChunk(&_Rollup.TransactOpts, _maxNumTxInChunk)
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
	BatchIndex uint64
	Winner     common.Address
	Res        common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeRes is a free log retrieval operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) FilterChallengeRes(opts *bind.FilterOpts, batchIndex []uint64, winner []common.Address, res []string) (*RollupChallengeResIterator, error) {

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

// WatchChallengeRes is a free log subscription operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) WatchChallengeRes(opts *bind.WatchOpts, sink chan<- *RollupChallengeRes, batchIndex []uint64, winner []common.Address, res []string) (event.Subscription, error) {

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

// ParseChallengeRes is a log parse operation binding the contract event 0x1e66d5dca70bf28588ef2f5cb3c299e65e2e7bdef2767823d3ae47a9caff95c6.
//
// Solidity: event ChallengeRes(uint64 indexed batchIndex, address indexed winner, string indexed res)
func (_Rollup *RollupFilterer) ParseChallengeRes(log types.Log) (*RollupChallengeRes, error) {
	event := new(RollupChallengeRes)
	if err := _Rollup.contract.UnpackLog(event, "ChallengeRes", log); err != nil {
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

// RollupUpdateMaxNumTxInChunkIterator is returned from FilterUpdateMaxNumTxInChunk and is used to iterate over the raw logs and unpacked data for UpdateMaxNumTxInChunk events raised by the Rollup contract.
type RollupUpdateMaxNumTxInChunkIterator struct {
	Event *RollupUpdateMaxNumTxInChunk // Event containing the contract specifics and raw log

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
func (it *RollupUpdateMaxNumTxInChunkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateMaxNumTxInChunk)
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
		it.Event = new(RollupUpdateMaxNumTxInChunk)
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
func (it *RollupUpdateMaxNumTxInChunkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateMaxNumTxInChunkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateMaxNumTxInChunk represents a UpdateMaxNumTxInChunk event raised by the Rollup contract.
type RollupUpdateMaxNumTxInChunk struct {
	OldMaxNumTxInChunk *big.Int
	NewMaxNumTxInChunk *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxNumTxInChunk is a free log retrieval operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) FilterUpdateMaxNumTxInChunk(opts *bind.FilterOpts) (*RollupUpdateMaxNumTxInChunkIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return &RollupUpdateMaxNumTxInChunkIterator{contract: _Rollup.contract, event: "UpdateMaxNumTxInChunk", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxNumTxInChunk is a free log subscription operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) WatchUpdateMaxNumTxInChunk(opts *bind.WatchOpts, sink chan<- *RollupUpdateMaxNumTxInChunk) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateMaxNumTxInChunk)
				if err := _Rollup.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
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

// ParseUpdateMaxNumTxInChunk is a log parse operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_Rollup *RollupFilterer) ParseUpdateMaxNumTxInChunk(log types.Log) (*RollupUpdateMaxNumTxInChunk, error) {
	event := new(RollupUpdateMaxNumTxInChunk)
	if err := _Rollup.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
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

// RollupUpdateProverIterator is returned from FilterUpdateProver and is used to iterate over the raw logs and unpacked data for UpdateProver events raised by the Rollup contract.
type RollupUpdateProverIterator struct {
	Event *RollupUpdateProver // Event containing the contract specifics and raw log

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
func (it *RollupUpdateProverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupUpdateProver)
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
		it.Event = new(RollupUpdateProver)
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
func (it *RollupUpdateProverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupUpdateProverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupUpdateProver represents a UpdateProver event raised by the Rollup contract.
type RollupUpdateProver struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateProver is a free log retrieval operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) FilterUpdateProver(opts *bind.FilterOpts, account []common.Address) (*RollupUpdateProverIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return &RollupUpdateProverIterator{contract: _Rollup.contract, event: "UpdateProver", logs: logs, sub: sub}, nil
}

// WatchUpdateProver is a free log subscription operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) WatchUpdateProver(opts *bind.WatchOpts, sink chan<- *RollupUpdateProver, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupUpdateProver)
				if err := _Rollup.contract.UnpackLog(event, "UpdateProver", log); err != nil {
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

// ParseUpdateProver is a log parse operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_Rollup *RollupFilterer) ParseUpdateProver(log types.Log) (*RollupUpdateProver, error) {
	event := new(RollupUpdateProver)
	if err := _Rollup.contract.UnpackLog(event, "UpdateProver", log); err != nil {
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
