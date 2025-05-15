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

// IL2StakingUndelegateRequest is an auto generated low-level Go binding around an user-defined struct.
type IL2StakingUndelegateRequest struct {
	Amount      *big.Int
	UnlockEpoch *big.Int
}

// TypesStakerInfo is an auto generated low-level Go binding around an user-defined struct.
type TypesStakerInfo struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}

// L2StakingMetaData contains all meta data concerning the L2Staking contract.
var L2StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherStaking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidPageSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidSequencerSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoClaimableUndelegateRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoStakers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoUndelegateRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrOnlyMorphTokenContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrOnlySystem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRequestExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRewardNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRewardStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrStartTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroLockEpochs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroSequencerSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"CommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commissionAmount\",\"type\":\"uint256\"}],\"name\":\"Distributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegateeFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegateeTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmountFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmountTo\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RewardStartTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"SequencerSetMaxSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"StakerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakerAddresses\",\"type\":\"address[]\"}],\"name\":\"StakerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockEpoch\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegationClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_STAKING\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo\",\"name\":\"add\",\"type\":\"tuple\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"claimUndelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"claimableUndelegateRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"commissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"delegateeDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegatorDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo\",\"name\":\"add\",\"type\":\"tuple\"}],\"name\":\"emergencyAddStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"emergencyRemoveStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seequencer\",\"type\":\"address\"}],\"name\":\"epochSequencerBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochTotalBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageIndex\",\"type\":\"uint256\"}],\"name\":\"getAllDelegatorsInPagination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorsTotalNumber\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegatorsInPage\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakerAddressesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakerAddresses\",\"type\":\"address[]\"}],\"name\":\"getStakesInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sequencersMaxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_undelegateLockEpochs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardStartTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"_stakers\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStakingTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"lockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"pendingUndelegateRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"queryDelegationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"queryUnclaimedCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"recordBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegateeFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegateeTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"removeStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerSetMaxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakerRankings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ranking\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegateLockEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"undelegateRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIL2Staking.UndelegateRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"undelegateSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardStartTime\",\"type\":\"uint256\"}],\"name\":\"updateRewardStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sequencerSetMaxSize\",\"type\":\"uint256\"}],\"name\":\"updateSequencerSetMaxSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61012060405234801562000011575f80fd5b506040516200608f3803806200608f8339810160408190526200003491620000a7565b7353000000000000000000000000000000000000076080526001600160a01b031660a05273530000000000000000000000000000000000001360c05273530000000000000000000000000000000000001760e05273530000000000000000000000000000000000002161010052620000d6565b5f60208284031215620000b8575f80fd5b81516001600160a01b0381168114620000cf575f80fd5b9392505050565b60805160a05160c05160e05161010051615f0b620001845f395f818161047b015261465c01525f818161064101526148a101525f81816107ac01528181613654015281816149cc01528181614a8101528181614b5d01528181614e2001528181614ecd0152614fa901525f81816105e60152818161311a0152613aaa01525f81816104ea0152818161069d015281816130f00152818161314401528181613a800152613ad40152615f0b5ff3fe608060405234801561000f575f80fd5b506004361061034f575f3560e01c806376671808116101be578063affed0e0116100fe578063eefecafd1161009e578063fad99f9811610079578063fad99f98146107f3578063fc6facc6146107fb578063fe3488841461080e578063ff4840cd1461082d575f80fd5b8063eefecafd146107ce578063f0261bc2146107d7578063f2fde38b146107e0575f80fd5b8063bf2dca0a116100d9578063bf2dca0a14610748578063cce6cf9f14610773578063d31d83d914610786578063d5577141146107a7575f80fd5b8063affed0e0146106f2578063b5d2e0dc146106fb578063b7a587bf1461071a575f80fd5b80638e21d5fb11610169578063927ede2d11610144578063927ede2d1461069857806396ab994d146106bf5780639d51c3b9146106cc578063a61bb764146106df575f80fd5b80638e21d5fb1461063c5780639168ae721461066357806391c05b0b14610685575f80fd5b8063831cfb5811610199578063831cfb58146105e157806384d7d1d4146106085780638da5cb5b1461062b575f80fd5b806376671808146105a05780637b05afb5146105a85780637c7e8bd2146105ce575f80fd5b80633434735f11610294578063459598a2116102345780636bd8f8041161020f5780636bd8f8041461056a5780637046529b1461057d578063715018a614610590578063746c8ae114610598575f80fd5b8063459598a21461053c57806346cdc18a1461054f5780634d99dd1614610557575f80fd5b80633cb747bf1161026f5780633cb747bf146104e857806340b5c8371461050e57806343352d6114610521578063439162b514610529575f80fd5b80633434735f146104765780633b2713c5146104b55780633b802421146104df575f80fd5b806313f22527116102ff578063201018fb116102da578063201018fb146104315780632cc138be146104445780632e787be31461044d57806330158eea14610456575f80fd5b806313f22527146103d057806319fac8fd146103e35780631d5611b8146103f6575f80fd5b80630321731c1161032f5780630321731c146103a15780630eb573af146103b457806312a3e947146103c7575f80fd5b806243b758146103535780629c6f0c14610379578063026e402b1461038e575b5f80fd5b6103666103613660046155b7565b610840565b6040519081526020015b60405180910390f35b61038c6103873660046155d2565b610866565b005b61038c61039c36600461561c565b610a29565b6103666103af3660046155b7565b610fcd565b61038c6103c2366004615646565b611007565b610366609a5481565b6103666103de3660046155b7565b6110dc565b61038c6103f1366004615646565b6111a4565b61041c6104043660046155b7565b60a36020525f90815260409020805460019091015482565b60408051928352602083019190915201610370565b61036661043f366004615646565b61129d565b61036660985481565b61036660995481565b6104696104643660046156a5565b611487565b6040516103709190615745565b61049d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610370565b6103666104c33660046157eb565b60a460209081525f928352604080842090915290825290205481565b610366609c5481565b7f000000000000000000000000000000000000000000000000000000000000000061049d565b61038c61051c366004615646565b6116ae565b610469611795565b61038c610537366004615817565b6119b1565b61049d61054a366004615646565b611e72565b609e54610366565b61038c61056536600461561c565b611e9a565b61038c610578366004615885565b6125e0565b61038c61058b3660046155d2565b6130e5565b61038c613259565b61038c61326c565b6103666134fa565b61041c6105b63660046155b7565b60a16020525f90815260409020805460019091015482565b6103666105dc3660046155b7565b613558565b61049d7f000000000000000000000000000000000000000000000000000000000000000081565b61061b6106163660046155b7565b613575565b6040519015158152602001610370565b6033546001600160a01b031661049d565b61049d7f000000000000000000000000000000000000000000000000000000000000000081565b6106766106713660046155b7565b61359f565b604051610370939291906158c3565b61038c610693366004615646565b613651565b61049d7f000000000000000000000000000000000000000000000000000000000000000081565b60975461061b9060ff1681565b6103666106da3660046157eb565b6138ca565b6103666106ed36600461561c565b6138dc565b610366609d5481565b6103666107093660046155b7565b609f6020525f908152604090205481565b61072d61072836600461561c565b613a14565b60408051825181526020928301519281019290925201610370565b6103666107563660046155b7565b6001600160a01b03165f90815260a1602052604090206001015490565b61038c6107813660046158f3565b613a75565b61079961079436600461593b565b613fd7565b6040516103709291906159b0565b61049d7f000000000000000000000000000000000000000000000000000000000000000081565b61036660aa5481565b610366609b5481565b61038c6107ee3660046155b7565b61414f565b61038c6141f9565b61038c6108093660046158f3565b6142ab565b61036661081c3660046155b7565b60ab6020525f908152604090205481565b61038c61083b3660046155b7565b614659565b6001600160a01b0381165f90815260a2602052604081206108609061470e565b92915050565b61086e614717565b81609d5481146108aa576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108b58360016159fd565b609d55609f5f6108c860208501856155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f20545f0361096a57609e6108fd60208401846155b7565b81546001810183555f9283526020808420909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039390931692909217909155609e5491609f91610950908601866155b7565b6001600160a01b0316815260208101919091526040015f20555b8160a05f61097b60208401846155b7565b6001600160a01b0316815260208101919091526040015f2061099d8282615b2d565b506109ad905060208301836155b7565b6001600160a01b03167f058ecb29c230cd5df283c89e996187ed521393fe4546cd1b097921c4b2de293d60208401356109e96040860186615a10565b6040516109f893929190615c94565b60405180910390a260975460ff16158015610a175750609954609e5411155b15610a2457610a2461478b565b505050565b6001600160a01b0382165f908152609f6020526040812054839103610a7a576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8261490b565b815f03610abb576040517f608294ac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610adc336001600160a01b0385165f90815260a2602052604090209061497e565b5060975460ff16610b6e576001600160a01b0383165f90815260a46020908152604080832033845290915281208054849290610b199084906159fd565b90915550506001600160a01b0383165f90815260a3602052604081208054849290610b459084906159fd565b90915550506001600160a01b0383165f90815260a3602052604090208054600190910155610c8a565b6001600160a01b0383165f90815260a3602090815260408083206001810154905460a48452828520338652909352908320549092829003610be5576001600160a01b0386165f81815260a460209081526040808320338452825280832089905592825260a390522060018101869055859055610c86565b81610bf08487615ce7565b610bfa9190615d2b565b610c0490826159fd565b6001600160a01b0387165f81815260a46020908152604080832033845282528083209490945591815260a39091529081208054879290610c459084906159fd565b90915550829050610c568487615ce7565b610c609190615d2b565b610c6a90846159fd565b6001600160a01b0387165f90815260a360205260409020600101555b5050505b6001600160a01b0383165f90815260a36020526040902054829003610cc1576001609c5f828254610cbb91906159fd565b90915550505b6001600160a01b0383165f908152609f602052604090205460975460ff168015610ceb5750600181115b15610f20575f610cfc600183615d3e565b90505b8015610f1e5760a35f609e610d15600185615d3e565b81548110610d2557610d25615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a39290919085908110610d6557610d65615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115610f0c575f609e610d9d600184615d3e565b81548110610dad57610dad615d51565b5f91825260209091200154609e80546001600160a01b0390921692509083908110610dda57610dda615d51565b5f918252602090912001546001600160a01b0316609e610dfb600185615d3e565b81548110610e0b57610e0b615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e8381548110610e4a57610e4a615d51565b5f9182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0393909316929092179091558290609f90609e610e91600185615d3e565b81548110610ea157610ea1615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055610ed08260016159fd565b609f5f609e8581548110610ee657610ee6615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b80610f1681615d7e565b915050610cff565b505b6001600160a01b0384165f81815260a360209081526040918290205482518781529182015281513393927f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04928290030190a3610f7d333085614992565b60975460ff168015610f905750609b5481115b8015610fb557506099546001600160a01b0385165f908152609f602052604090205411155b15610fc257610fc261478b565b50610a246001606555565b6001600160a01b0381165f90815260a66020526040812054600f81810b700100000000000000000000000000000000909204900b03610860565b61100f614717565b80158061101d575060995481145b15611054576040517f383a648e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609980549082905560408051828152602081018490527f98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed910160405180910390a16097545f9060ff166110a957609e546110ad565b609c545b90505f60995482106110c1576099546110c3565b815b9050609b5481146110d6576110d661478b565b50505050565b6001600160a01b0381165f90815260a66020526040812054609754600f82810b700100000000000000000000000000000000909304900b919091039060ff166111255792915050565b5f5b8181101561119d576001600160a01b0384165f90815260a6602052604081206111509083614c1e565b5f81815260a5602090815260409182902082518084019093528054835260010154908201819052919250906111836134fa565b1015611193575090949350505050565b5050600101611127565b5092915050565b335f818152609f602052604081205490036111eb576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6014821115611226576040517f6e11528c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f81815260a1602081815260408084208054825180840184528981526001830180548287019081529789905295855251909155935190925581518681529081018390529192917f6e500db30ce535d38852e318f333e9be41a3fec6c65d234ebb06203c896db9a5910160405180910390a2505050565b5f6112a661490b565b6112ee60a65f335b6001600160a01b03166001600160a01b031681526020019081526020015f2054600f81810b700100000000000000000000000000000000909204900b0390565b5f03611326576040517f5f013ef800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81158061133d575061133a60a65f336112ae565b82115b6113475781611353565b61135360a65f336112ae565b91505f5b82156113fd57335f90815260a66020526040812061137490614cb3565b5f81815260a5602090815260409182902082518084019093528054835260010154908201526097549192509060ff1680156113b9575080602001516113b76134fa565b105b156113c55750506113fd565b335f90815260a6602052604090206113dc90614d2b565b5080516113e990846159fd565b92506113f485615d7e565b94505050611357565b805f03611436576040517f3cc5dedc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611441335b82614de6565b60405181815233907fcc3089abc79631b3c0c81414a72e237c08559073a970cf474e36ae965e382fb39060200160405180910390a290506114826001606555565b919050565b60605f8267ffffffffffffffff8111156114a3576114a3615a71565b6040519080825280602002602001820160405280156114ef57816020015b60408051606080820183525f8083526020830152918101919091528152602001906001900390816114c15790505b5090505f5b838110156116a657604051806060016040528060a05f88888681811061151c5761151c615d51565b905060200201602081019061153191906155b7565b6001600160a01b03908116825260208083019390935260409091015f90812054909116835291019060a09088888681811061156e5761156e615d51565b905060200201602081019061158391906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f2060010154815260200160a05f8888868181106115c0576115c0615d51565b90506020020160208101906115d591906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f20600201805461160290615a9e565b80601f016020809104026020016040519081016040528092919081815260200182805461162e90615a9e565b80156116795780601f1061165057610100808354040283529160200191611679565b820191905f5260205f20905b81548152906001019060200180831161165c57829003601f168201915b505050505081525082828151811061169357611693615d51565b60209081029190910101526001016114f4565b509392505050565b6116b6614717565b60975460ff16156116f3576040517fbd51da0d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b428111158061170d575061170a6201518082615d93565b15155b80611719575060985481145b15611750576040517fde16b26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609880549082905560408051828152602081018490527f91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555910160405180910390a15050565b609e546060905f9067ffffffffffffffff8111156117b5576117b5615a71565b60405190808252806020026020018201604052801561180157816020015b60408051606080820183525f8083526020830152918101919091528152602001906001900390816117d35790505b5090505f5b609e548110156119ab57604051806060016040528060a05f609e858154811061183157611831615d51565b5f9182526020808320909101546001600160a01b0390811684528382019490945260409092018120549092168352609e8054939091019260a09291908690811061187d5761187d615d51565b905f5260205f20015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f2060010154815260200160a05f609e85815481106118d6576118d6615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020600201805461190790615a9e565b80601f016020809104026020016040519081016040528092919081815260200182805461193390615a9e565b801561197e5780601f106119555761010080835404028352916020019161197e565b820191905f5260205f20905b81548152906001019060200180831161196157829003601f168201915b505050505081525082828151811061199857611998615d51565b6020908102919091010152600101611806565b50919050565b5f54610100900460ff16158080156119cf57505f54600160ff909116105b806119e85750303b1580156119e857505f5460ff166001145b611a79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611ad5575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b038716611b15576040517fee77070400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855f03611b4e576040517f2da55d0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f03611b87576040517f7d8ad8a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4284111580611ba15750611b9e6201518085615d93565b15155b15611bd8576040517fde16b26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f829003611c12576040517fbb01aad100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c1b87615063565b611c236150c1565b6099869055609a8590556098849055609b8290555f5b609b54811015611d9457838382818110611c5557611c55615d51565b9050602002810190611c679190615da6565b60a05f868685818110611c7c57611c7c615d51565b9050602002810190611c8e9190615da6565b611c9c9060208101906155b7565b6001600160a01b0316815260208101919091526040015f20611cbe8282615b2d565b905050609e848483818110611cd557611cd5615d51565b9050602002810190611ce79190615da6565b611cf59060208101906155b7565b8154600180820184555f938452602090932001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055611d3e9082906159fd565b609f5f868685818110611d5357611d53615d51565b9050602002810190611d659190615da6565b611d739060208101906155b7565b6001600160a01b0316815260208101919091526040015f2055600101611c39565b50604080515f8152602081018890527f98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed910160405180910390a1604080515f8152602081018690527f91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555910160405180910390a18015611e69575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b609e8181548110611e81575f80fd5b5f918252602090912001546001600160a01b0316905081565b611ea261490b565b5f611ead833361515f565b9050805f03611ee8576040517f857ad50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81811015611f22576040517f08c2348a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8215611f2f5782611f31565b815b6001600160a01b0385165f908152609f6020526040812054609754929350159160ff16611f5e575f611f6c565b609a54611f6c9060016159fd565b60408051808201909152848152602081018290529091505f33611f8e336151e5565b60405160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301526034820152605401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f81815260a590935291205490915015612042576040517fdeeb052700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81815260a560209081526040808320855181558583015160019182015533845260a68352818420805470010000000000000000000000000000000090819004600f0b8087528284019095529290942085905583546fffffffffffffffffffffffffffffffff90811693909101160217905560975460ff16612145576001600160a01b0388165f90815260a460209081526040808320338452909152812080548792906120f0908490615d3e565b90915550506001600160a01b0388165f90815260a360205260408120805487929061211c908490615d3e565b90915550506001600160a01b0388165f90815260a360205260409020805460019091015561221a565b6001600160a01b0388165f90815260a3602090815260408083206001810154905460a4845282852033865290935292205481612181848a615ce7565b61218b9190615d2b565b6121959082615d3e565b6001600160a01b038c165f81815260a46020908152604080832033845282528083209490945591815260a390915290812080548a92906121d6908490615d3e565b909155508290506121e7848a615ce7565b6121f19190615d2b565b6121fb9084615d3e565b6001600160a01b038c165f90815260a360205260409020600101555050505b6001600160a01b0388165f908152609f602052604090205484158015612242575060975460ff165b801561224f5750609c5481105b156124b0576001600160a01b0389165f908152609f602052604081205461227890600190615d3e565b90505b6001609c5461228a9190615d3e565b8110156124ae5760a35f609e83815481106122a7576122a7615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040018120549060a390609e6122dc8560016159fd565b815481106122ec576122ec615d51565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205411156124a6575f609e828154811061232a5761232a615d51565b5f918252602090912001546001600160a01b03169050609e61234d8360016159fd565b8154811061235d5761235d615d51565b5f91825260209091200154609e80546001600160a01b03909216918490811061238857612388615d51565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905580609e6123cb8460016159fd565b815481106123db576123db615d51565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905561241b8260016159fd565b609f5f609e858154811061243157612431615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020556124608260026159fd565b609f5f609e6124708660016159fd565b8154811061248057612480615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b60010161227b565b505b841580156124d357506001600160a01b0389165f90815260a36020526040902054155b156124f0576001609c5f8282546124ea9190615d3e565b90915550505b6001600160a01b0389165f90815260a360205260409020543360408051898152602081018490529081018790526001600160a01b03918216918c16907f92039db29d8c0a1aa1433fe109c69488c8c5e51b23c9de7d303ad80c1fef778c9060600160405180910390a385158015612569575060975460ff165b80156125775750609b548211155b80156125bd5750609b546001600160a01b038b165f908152609f602052604090205411806125bd5750609c546001600160a01b038b165f908152609f6020526040902054115b156125ca576125ca61478b565b50505050505050506125dc6001606555565b5050565b6001600160a01b0383165f908152609f6020526040812054849103612631576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0383165f908152609f6020526040812054849103612682576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61268a61490b565b5f612695863361515f565b90506126a1863361515f565b5f036126d9576040517f857ad50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b836126e4873361515f565b101561271c576040517f08c2348a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8415612729578461272b565b815b6001600160a01b0388165f908152609f6020526040812054609754929350909190159060ff166127dc576001600160a01b0389165f90815260a46020908152604080832033845290915281208054859290612787908490615d3e565b90915550506001600160a01b0389165f90815260a36020526040812080548592906127b3908490615d3e565b90915550506001600160a01b0389165f90815260a36020526040902080546001909101556128b1565b6001600160a01b0389165f90815260a3602090815260408083206001810154905460a4845282852033865290935292205481612818848c615ce7565b6128229190615d2b565b61282c9082615d3e565b6001600160a01b038d165f81815260a46020908152604080832033845282528083209490945591815260a3909152908120805488929061286d908490615d3e565b9091555082905061287e848c615ce7565b6128889190615d2b565b6128929084615d3e565b6001600160a01b038d165f90815260a360205260409020600101555050505b6001600160a01b0389165f908152609f6020526040902054811580156128d9575060975460ff165b80156128e65750609c5481105b15612b47576001600160a01b038a165f908152609f602052604081205461290f90600190615d3e565b90505b6001609c546129219190615d3e565b811015612b455760a35f609e838154811061293e5761293e615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040018120549060a390609e6129738560016159fd565b8154811061298357612983615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115612b3d575f609e82815481106129c1576129c1615d51565b5f918252602090912001546001600160a01b03169050609e6129e48360016159fd565b815481106129f4576129f4615d51565b5f91825260209091200154609e80546001600160a01b039092169184908110612a1f57612a1f615d51565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905580609e612a628460016159fd565b81548110612a7257612a72615d51565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055612ab28260016159fd565b609f5f609e8581548110612ac857612ac8615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055612af78260026159fd565b609f5f609e612b078660016159fd565b81548110612b1757612b17615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b600101612912565b505b81158015612b6a57506001600160a01b038a165f90815260a36020526040902054155b15612b87576001609c5f828254612b819190615d3e565b90915550505b81158015612b97575060975460ff165b8015612ba55750609b548111155b8015612beb5750609b546001600160a01b038b165f908152609f60205260409020541180612beb5750609c546001600160a01b038b165f908152609f6020526040902054115b15612bf557600192505b612c16336001600160a01b038b165f90815260a2602052604090209061497e565b5060975460ff16612ca8576001600160a01b0389165f90815260a46020908152604080832033845290915281208054869290612c539084906159fd565b90915550506001600160a01b0389165f90815260a3602052604081208054869290612c7f9084906159fd565b90915550506001600160a01b0389165f90815260a3602052604090208054600190910155612d7d565b6001600160a01b0389165f90815260a3602090815260408083206001810154905460a4845282852033865290935292205481612ce4848d615ce7565b612cee9190615d2b565b612cf890826159fd565b6001600160a01b038d165f81815260a46020908152604080832033845282528083209490945591815260a39091529081208054899290612d399084906159fd565b90915550829050612d4a848d615ce7565b612d549190615d2b565b612d5e90846159fd565b6001600160a01b038d165f90815260a360205260409020600101555050505b6001600160a01b0389165f90815260a36020526040902054849003612db4576001609c5f828254612dae91906159fd565b90915550505b506001600160a01b0388165f908152609f602052604090205460975460ff168015612ddf5750600181115b15613014575f612df0600183615d3e565b90505b80156130125760a35f609e612e09600185615d3e565b81548110612e1957612e19615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a39290919085908110612e5957612e59615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115613000575f609e612e91600184615d3e565b81548110612ea157612ea1615d51565b5f91825260209091200154609e80546001600160a01b0390921692509083908110612ece57612ece615d51565b5f918252602090912001546001600160a01b0316609e612eef600185615d3e565b81548110612eff57612eff615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e8381548110612f3e57612f3e615d51565b5f9182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0393909316929092179091558290609f90609e612f85600185615d3e565b81548110612f9557612f95615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055612fc48260016159fd565b609f5f609e8581548110612fda57612fda615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b8061300a81615d7e565b915050612df3565b505b60975460ff1680156130275750609b5481115b801561304c57506099546001600160a01b038a165f908152609f602052604090205411155b1561305657600192505b82156130645761306461478b565b6001600160a01b038a81165f81815260a36020908152604080832054948e16808452928190205481518a8152928301869052908201819052923392917ffdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b49060600160405180910390a4505050505050506130de6001606555565b5050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156131cd57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561319e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131c29190615de2565b6001600160a01b0316145b61086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f7374616b696e673a206f6e6c79206f74686572207374616b696e6720636f6e7460448201527f7261637420616c6c6f77656400000000000000000000000000000000000000006064820152608401611a70565b613261614717565b61326a5f615063565b565b613274614717565b6098544210156132b0576040517f080bb11a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609c545f036132eb576040517fd7d776cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091555b609e54811015613497575f5b8181101561348e5760a35f609e838154811061334457613344615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a3929091908690811061338457613384615d51565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115613486575f609e82815481106133c2576133c2615d51565b5f91825260209091200154609e80546001600160a01b03909216925090849081106133ef576133ef615d51565b5f91825260209091200154609e80546001600160a01b03909216918490811061341a5761341a615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e848154811061345957613459615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550505b600101613326565b5060010161331a565b505f5b609e548110156134f1576134af8160016159fd565b609f5f609e84815481106134c5576134c5615d51565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205560010161349a565b5061326a61478b565b5f609854421015613537576040517fd021716f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b62015180609854426135499190615d3e565b6135539190615d2b565b905090565b6001600160a01b0381165f90815260a76020526040812054610860565b6001600160a01b0381165f90815260a4602090815260408083203384529091528120541515610860565b60a06020525f90815260409020805460018201546002830180546001600160a01b039093169391926135d090615a9e565b80601f01602080910402602001604051908101604052809291908181526020018280546135fc90615a9e565b80156136475780601f1061361e57610100808354040283529160200191613647565b820191905f5260205f20905b81548152906001019060200180831161362a57829003601f168201915b5050505050905083565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146136b3576040517f4032cbb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60aa5415613852575f5b6136c760a861470e565b811015613850575f60a1816136dd60a88561520a565b6001600160a01b0316815260208101919091526040015f9081205460aa5490925060ab8261370c60a88761520a565b6001600160a01b0316815260208101919091526040015f205461372f9086615ce7565b6137399190615d2b565b90505f60646137488484615ce7565b6137529190615d2b565b90505f61375f8284615d3e565b90508160a15f61377060a88961520a565b6001600160a01b03166001600160a01b031681526020019081526020015f206001015f8282546137a091906159fd565b9091555081905060a35f6137b560a88961520a565b6001600160a01b03166001600160a01b031681526020019081526020015f205f015f8282546137e491906159fd565b909155506137f5905060a88661520a565b6001600160a01b03167f60ce3cc2d133631eac66a476f14997a9fa682bd05a60dd993cf02285822d78d88284604051613838929190918252602082015260400190565b60405180910390a25050600190920191506136bd9050565b505b5f61385d60a861470e565b90505f5b818110156138c15760ab5f61387760a88261520a565b6001600160a01b03166001600160a01b031681526020019081526020015f205f90556138b86138b05f60a861520a90919063ffffffff16565b60a890615215565b50600101613861565b50505f60aa5550565b5f6138d5838361515f565b9392505050565b6001600160a01b0382165f90815260a66020526040812054600f81810b700100000000000000000000000000000000909204900b035f0361391e57505f610860565b81158061395e57506001600160a01b0383165f90815260a66020526040902054600f81810b700100000000000000000000000000000000909204900b0382115b613968578161399e565b6001600160a01b0383165f90815260a66020526040902054600f81810b700100000000000000000000000000000000909204900b035b91505f805b838110156116a6576001600160a01b0385165f90815260a6602052604081206139cc9083614c1e565b5f81815260a56020908152604091829020825180840190935280548084526001909101549183019190915291925090613a0590856159fd565b935050508060010190506139a3565b604080518082019091525f80825260208201526001600160a01b0383165f90815260a660205260408120613a489084614c1e565b5f90815260a560209081526040918290208251808401909352805483526001015490820152949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015613b5d57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b2e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b529190615de2565b6001600160a01b0316145b613be9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f7374616b696e673a206f6e6c79206f74686572207374616b696e6720636f6e7460448201527f7261637420616c6c6f77656400000000000000000000000000000000000000006064820152608401611a70565b82609d548114613c25576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613c308460016159fd565b609d555f805b83811015613f8f57609b54609f5f878785818110613c5657613c56615d51565b9050602002016020810190613c6b91906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f205411613c9557600191505b5f609f5f878785818110613cab57613cab615d51565b9050602002016020810190613cc091906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f20541115613f11575f6001609f5f888886818110613cfe57613cfe615d51565b9050602002016020810190613d1391906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f2054613d3d9190615d3e565b90505b609e54613d4f90600190615d3e565b811015613e2157609e613d638260016159fd565b81548110613d7357613d73615d51565b5f91825260209091200154609e80546001600160a01b039092169183908110613d9e57613d9e615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001609f5f609e8481548110613de157613de1615d51565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190613e14908490615d3e565b9091555050600101613d40565b50609e805480613e3357613e33615dfd565b5f8281526020812082015f19908101805473ffffffffffffffffffffffffffffffffffffffff19169055909101909155609f90868684818110613e7857613e78615d51565b9050602002016020810190613e8d91906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f205f90555f60a35f878785818110613ec557613ec5615d51565b9050602002016020810190613eda91906155b7565b6001600160a01b0316815260208101919091526040015f20541115613f11576001609c5f828254613f0b9190615d3e565b90915550505b60a05f868684818110613f2657613f26615d51565b9050602002016020810190613f3b91906155b7565b6001600160a01b0316815260208101919091526040015f908120805473ffffffffffffffffffffffffffffffffffffffff191681556001810182905590613f85600283018261555d565b5050600101613c36565b507f3511bf213f9290ba907e91e12a43e8471251e1879580ae5509292a3514c23f618484604051613fc1929190615e2a565b60405180910390a180156130de576130de61478b565b5f6060835f03614013576040517f89076b3900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385165f90815260a2602052604090206140339061470e565b91508367ffffffffffffffff81111561404e5761404e615a71565b604051908082528060200260200182016040528015614077578160200160208202803683370190505b5090505f6140858486615ce7565b90505f600161409486826159fd565b61409e9088615ce7565b6140a89190615d3e565b90506140b5600185615d3e565b8111156140ca576140c7600185615d3e565b90505b815f5b82821161414357614101826140e181615e77565b6001600160a01b038c165f90815260a2602052604090209094509061520a565b858261410c81615e77565b93508151811061411e5761411e615d51565b60200260200101906001600160a01b031690816001600160a01b0316815250506140cd565b50505050935093915050565b614157614717565b6001600160a01b0381166141ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401611a70565b6141f681615063565b50565b61420161490b565b335f90815260a16020526040812060010154900361424b576040517f5426dfcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f81815260a16020526040812060010180549190559061426b9061143b565b60405181815233907f8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e4119060200160405180910390a25061326a6001606555565b6142b3614717565b82609d5481146142ef576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6142fa8460016159fd565b609d555f805b83811015613f8f57609b54609f5f87878581811061432057614320615d51565b905060200201602081019061433591906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f20541161435f57600191505b5f609f5f87878581811061437557614375615d51565b905060200201602081019061438a91906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f205411156145db575f6001609f5f8888868181106143c8576143c8615d51565b90506020020160208101906143dd91906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f20546144079190615d3e565b90505b609e5461441990600190615d3e565b8110156144eb57609e61442d8260016159fd565b8154811061443d5761443d615d51565b5f91825260209091200154609e80546001600160a01b03909216918390811061446857614468615d51565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001609f5f609e84815481106144ab576144ab615d51565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906144de908490615d3e565b909155505060010161440a565b50609e8054806144fd576144fd615dfd565b5f8281526020812082015f19908101805473ffffffffffffffffffffffffffffffffffffffff19169055909101909155609f9086868481811061454257614542615d51565b905060200201602081019061455791906155b7565b6001600160a01b03166001600160a01b031681526020019081526020015f205f90555f60a35f87878581811061458f5761458f615d51565b90506020020160208101906145a491906155b7565b6001600160a01b0316815260208101919091526040015f205411156145db576001609c5f8282546145d59190615d3e565b90915550505b60a05f8686848181106145f0576145f0615d51565b905060200201602081019061460591906155b7565b6001600160a01b0316815260208101919091526040015f908120805473ffffffffffffffffffffffffffffffffffffffff19168155600181018290559061464f600283018261555d565b5050600101614300565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316146146bb576040517f52d033bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6146c660a88261497e565b50600160aa5f8282546146d991906159fd565b90915550506001600160a01b0381165f90815260ab602052604081208054600192906147069084906159fd565b909155505050565b5f610860825490565b6033546001600160a01b0316331461326a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611a70565b60995460975460ff16156147af57609954609c5410156147aa5750609c545b6147c0565b609954609e5410156147c05750609e545b5f8167ffffffffffffffff8111156147da576147da615a71565b604051908082528060200260200182016040528015614803578160200160208202803683370190505b5090505f5b8281101561487057609e818154811061482357614823615d51565b905f5260205f20015f9054906101000a90046001600160a01b031682828151811061485057614850615d51565b6001600160a01b0390921660209283029190910190910152600101614808565b506040517f9b8201a40000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639b8201a4906148d6908490600401615e8f565b5f604051808303815f87803b1580156148ed575f80fd5b505af11580156148ff573d5f803e3d5ffd5b50509151609b55505050565b600260655403614977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611a70565b6002606555565b5f6138d5836001600160a01b038416615229565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614a13573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614a379190615ea1565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301528581166024830152604482018590529192507f0000000000000000000000000000000000000000000000000000000000000000909116906323b872dd906064016020604051808303815f875af1158015614ac9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614aed9190615eb8565b614b23576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614ba4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614bc89190615ea1565b9050821580614be0575082614bdd8383615d3e565b14155b156130de576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001606555565b5f80614c40614c2c84615275565b8554614c3b9190600f0b615ed7565b61532a565b84549091507001000000000000000000000000000000009004600f90810b9082900b12614c99576040517fb4120f1400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f614cda8254600f81810b700100000000000000000000000000000000909204900b131590565b15614d11576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f614d528254600f81810b700100000000000000000000000000000000909204900b131590565b15614d89576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614e67573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614e8b9190615ea1565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303815f875af1158015614f15573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614f399190615eb8565b614f6f576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614ff0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906150149190615ea1565b905082158061502c5750826150298383615d3e565b14155b156110d6576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff16615157576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611a70565b61326a6153be565b6001600160a01b038083165f90815260a460209081526040808320938516835292905290812054156151dd576001600160a01b038084165f81815260a360208181526040808420600181015460a48452828620978a168652968352908420549490935252546151ce9190615ce7565b6151d89190615d2b565b6138d5565b505f92915050565b6001600160a01b0381165f90815260a7602052604090208054600181018255906119ab565b5f6138d58383615454565b5f6138d5836001600160a01b03841661547a565b5f81815260018301602052604081205461526e57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610860565b505f610860565b5f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821115615326576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e206160448201527f6e20696e743235360000000000000000000000000000000000000000000000006064820152608401611a70565b5090565b80600f81900b8114611482576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203160448201527f32382062697473000000000000000000000000000000000000000000000000006064820152608401611a70565b5f54610100900460ff16614c17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611a70565b5f825f01828154811061546957615469615d51565b905f5260205f200154905092915050565b5f8181526001830160205260408120548015615554575f61549c600183615d3e565b85549091505f906154af90600190615d3e565b905081811461550e575f865f0182815481106154cd576154cd615d51565b905f5260205f200154905080875f0184815481106154ed576154ed615d51565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061551f5761551f615dfd565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610860565b5f915050610860565b50805461556990615a9e565b5f825580601f10615578575050565b601f0160209004905f5260205f20908101906141f691905b80821115615326575f8155600101615590565b6001600160a01b03811681146141f6575f80fd5b5f602082840312156155c7575f80fd5b81356138d5816155a3565b5f80604083850312156155e3575f80fd5b82359150602083013567ffffffffffffffff811115615600575f80fd5b830160608186031215615611575f80fd5b809150509250929050565b5f806040838503121561562d575f80fd5b8235615638816155a3565b946020939093013593505050565b5f60208284031215615656575f80fd5b5035919050565b5f8083601f84011261566d575f80fd5b50813567ffffffffffffffff811115615684575f80fd5b6020830191508360208260051b850101111561569e575f80fd5b9250929050565b5f80602083850312156156b6575f80fd5b823567ffffffffffffffff8111156156cc575f80fd5b6156d88582860161565d565b90969095509350505050565b5f81518084525f5b81811015615708576020818501810151868301820152016156ec565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156157dd578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180516001600160a01b03168452878101518885015286015160608785018190526157c9818601836156e4565b96890196945050509086019060010161576c565b509098975050505050505050565b5f80604083850312156157fc575f80fd5b8235615807816155a3565b91506020830135615611816155a3565b5f805f805f8060a0878903121561582c575f80fd5b8635615837816155a3565b9550602087013594506040870135935060608701359250608087013567ffffffffffffffff811115615867575f80fd5b61587389828a0161565d565b979a9699509497509295939492505050565b5f805f60608486031215615897575f80fd5b83356158a2816155a3565b925060208401356158b2816155a3565b929592945050506040919091013590565b6001600160a01b0384168152826020820152606060408201525f6158ea60608301846156e4565b95945050505050565b5f805f60408486031215615905575f80fd5b83359250602084013567ffffffffffffffff811115615922575f80fd5b61592e8682870161565d565b9497909650939450505050565b5f805f6060848603121561594d575f80fd5b8335615958816155a3565b95602085013595506040909401359392505050565b5f815180845260208085019450602084015f5b838110156159a55781516001600160a01b031687529582019590820190600101615980565b509495945050505050565b828152604060208201525f6159c8604083018461596d565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b80820180821115610860576108606159d0565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112615a43575f80fd5b83018035915067ffffffffffffffff821115615a5d575f80fd5b60200191503681900382131561569e575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b600181811c90821680615ab257607f821691505b6020821081036119ab577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b601f821115610a2457805f5260205f20601f840160051c81016020851015615b0e5750805b601f840160051c820191505b818110156130de575f8155600101615b1a565b8135615b38816155a3565b6001600160a01b03811673ffffffffffffffffffffffffffffffffffffffff1983541617825550600160208084013560018401556002830160408501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1863603018112615ba4575f80fd5b8501803567ffffffffffffffff811115615bbc575f80fd5b8036038483011315615bcc575f80fd5b615be081615bda8554615a9e565b85615ae9565b5f601f821160018114615c13575f8315615bfc57508382018601355b5f19600385901b1c1916600184901b178555615c89565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015615c5f57868501890135825593880193908901908801615c40565b5084821015615c7d575f1960f88660031b161c198885880101351681555b505060018360011b0185555b505050505050505050565b83815260406020820152816040820152818360608301375f818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b8082028115828204841417610860576108606159d0565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615d3957615d39615cfe565b500490565b81810381811115610860576108606159d0565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81615d8c57615d8c6159d0565b505f190190565b5f82615da157615da1615cfe565b500690565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615dd8575f80fd5b9190910192915050565b5f60208284031215615df2575f80fd5b81516138d5816155a3565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b60208082528181018390525f908460408401835b86811015615e6c578235615e51816155a3565b6001600160a01b031682529183019190830190600101615e3e565b509695505050505050565b5f5f198203615e8857615e886159d0565b5060010190565b602081525f6138d5602083018461596d565b5f60208284031215615eb1575f80fd5b5051919050565b5f60208284031215615ec8575f80fd5b815180151581146138d5575f80fd5b8082018281125f831280158216821582161715615ef657615ef66159d0565b50509291505056fea164736f6c6343000818000a",
}

// L2StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use L2StakingMetaData.ABI instead.
var L2StakingABI = L2StakingMetaData.ABI

// L2StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2StakingMetaData.Bin instead.
var L2StakingBin = L2StakingMetaData.Bin

// DeployL2Staking deploys a new Ethereum contract, binding an instance of L2Staking to it.
func DeployL2Staking(auth *bind.TransactOpts, backend bind.ContractBackend, _otherStaking common.Address) (common.Address, *types.Transaction, *L2Staking, error) {
	parsed, err := L2StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2StakingBin), backend, _otherStaking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2Staking{L2StakingCaller: L2StakingCaller{contract: contract}, L2StakingTransactor: L2StakingTransactor{contract: contract}, L2StakingFilterer: L2StakingFilterer{contract: contract}}, nil
}

// L2Staking is an auto generated Go binding around an Ethereum contract.
type L2Staking struct {
	L2StakingCaller     // Read-only binding to the contract
	L2StakingTransactor // Write-only binding to the contract
	L2StakingFilterer   // Log filterer for contract events
}

// L2StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2StakingSession struct {
	Contract     *L2Staking        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2StakingCallerSession struct {
	Contract *L2StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// L2StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2StakingTransactorSession struct {
	Contract     *L2StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2StakingRaw struct {
	Contract *L2Staking // Generic contract binding to access the raw methods on
}

// L2StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2StakingCallerRaw struct {
	Contract *L2StakingCaller // Generic read-only contract binding to access the raw methods on
}

// L2StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2StakingTransactorRaw struct {
	Contract *L2StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2Staking creates a new instance of L2Staking, bound to a specific deployed contract.
func NewL2Staking(address common.Address, backend bind.ContractBackend) (*L2Staking, error) {
	contract, err := bindL2Staking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2Staking{L2StakingCaller: L2StakingCaller{contract: contract}, L2StakingTransactor: L2StakingTransactor{contract: contract}, L2StakingFilterer: L2StakingFilterer{contract: contract}}, nil
}

// NewL2StakingCaller creates a new read-only instance of L2Staking, bound to a specific deployed contract.
func NewL2StakingCaller(address common.Address, caller bind.ContractCaller) (*L2StakingCaller, error) {
	contract, err := bindL2Staking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2StakingCaller{contract: contract}, nil
}

// NewL2StakingTransactor creates a new write-only instance of L2Staking, bound to a specific deployed contract.
func NewL2StakingTransactor(address common.Address, transactor bind.ContractTransactor) (*L2StakingTransactor, error) {
	contract, err := bindL2Staking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2StakingTransactor{contract: contract}, nil
}

// NewL2StakingFilterer creates a new log filterer instance of L2Staking, bound to a specific deployed contract.
func NewL2StakingFilterer(address common.Address, filterer bind.ContractFilterer) (*L2StakingFilterer, error) {
	contract, err := bindL2Staking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2StakingFilterer{contract: contract}, nil
}

// bindL2Staking binds a generic wrapper to an already deployed contract.
func bindL2Staking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Staking *L2StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Staking.Contract.L2StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Staking *L2StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Staking.Contract.L2StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Staking *L2StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Staking.Contract.L2StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2Staking *L2StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2Staking *L2StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2Staking *L2StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2Staking.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Staking *L2StakingCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Staking *L2StakingSession) MESSENGER() (common.Address, error) {
	return _L2Staking.Contract.MESSENGER(&_L2Staking.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_L2Staking *L2StakingCallerSession) MESSENGER() (common.Address, error) {
	return _L2Staking.Contract.MESSENGER(&_L2Staking.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_L2Staking *L2StakingCaller) MORPHTOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "MORPH_TOKEN_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_L2Staking *L2StakingSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _L2Staking.Contract.MORPHTOKENCONTRACT(&_L2Staking.CallOpts)
}

// MORPHTOKENCONTRACT is a free data retrieval call binding the contract method 0xd5577141.
//
// Solidity: function MORPH_TOKEN_CONTRACT() view returns(address)
func (_L2Staking *L2StakingCallerSession) MORPHTOKENCONTRACT() (common.Address, error) {
	return _L2Staking.Contract.MORPHTOKENCONTRACT(&_L2Staking.CallOpts)
}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L2Staking *L2StakingCaller) OTHERSTAKING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "OTHER_STAKING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L2Staking *L2StakingSession) OTHERSTAKING() (common.Address, error) {
	return _L2Staking.Contract.OTHERSTAKING(&_L2Staking.CallOpts)
}

// OTHERSTAKING is a free data retrieval call binding the contract method 0x831cfb58.
//
// Solidity: function OTHER_STAKING() view returns(address)
func (_L2Staking *L2StakingCallerSession) OTHERSTAKING() (common.Address, error) {
	return _L2Staking.Contract.OTHERSTAKING(&_L2Staking.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_L2Staking *L2StakingCaller) SEQUENCERCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "SEQUENCER_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_L2Staking *L2StakingSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _L2Staking.Contract.SEQUENCERCONTRACT(&_L2Staking.CallOpts)
}

// SEQUENCERCONTRACT is a free data retrieval call binding the contract method 0x8e21d5fb.
//
// Solidity: function SEQUENCER_CONTRACT() view returns(address)
func (_L2Staking *L2StakingCallerSession) SEQUENCERCONTRACT() (common.Address, error) {
	return _L2Staking.Contract.SEQUENCERCONTRACT(&_L2Staking.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_L2Staking *L2StakingCaller) SYSTEMADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "SYSTEM_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_L2Staking *L2StakingSession) SYSTEMADDRESS() (common.Address, error) {
	return _L2Staking.Contract.SYSTEMADDRESS(&_L2Staking.CallOpts)
}

// SYSTEMADDRESS is a free data retrieval call binding the contract method 0x3434735f.
//
// Solidity: function SYSTEM_ADDRESS() view returns(address)
func (_L2Staking *L2StakingCallerSession) SYSTEMADDRESS() (common.Address, error) {
	return _L2Staking.Contract.SYSTEMADDRESS(&_L2Staking.CallOpts)
}

// CandidateNumber is a free data retrieval call binding the contract method 0x3b802421.
//
// Solidity: function candidateNumber() view returns(uint256)
func (_L2Staking *L2StakingCaller) CandidateNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "candidateNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateNumber is a free data retrieval call binding the contract method 0x3b802421.
//
// Solidity: function candidateNumber() view returns(uint256)
func (_L2Staking *L2StakingSession) CandidateNumber() (*big.Int, error) {
	return _L2Staking.Contract.CandidateNumber(&_L2Staking.CallOpts)
}

// CandidateNumber is a free data retrieval call binding the contract method 0x3b802421.
//
// Solidity: function candidateNumber() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) CandidateNumber() (*big.Int, error) {
	return _L2Staking.Contract.CandidateNumber(&_L2Staking.CallOpts)
}

// ClaimableUndelegateRequest is a free data retrieval call binding the contract method 0x13f22527.
//
// Solidity: function claimableUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCaller) ClaimableUndelegateRequest(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "claimableUndelegateRequest", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableUndelegateRequest is a free data retrieval call binding the contract method 0x13f22527.
//
// Solidity: function claimableUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingSession) ClaimableUndelegateRequest(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.ClaimableUndelegateRequest(&_L2Staking.CallOpts, delegator)
}

// ClaimableUndelegateRequest is a free data retrieval call binding the contract method 0x13f22527.
//
// Solidity: function claimableUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) ClaimableUndelegateRequest(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.ClaimableUndelegateRequest(&_L2Staking.CallOpts, delegator)
}

// Commissions is a free data retrieval call binding the contract method 0x7b05afb5.
//
// Solidity: function commissions(address staker) view returns(uint256 rate, uint256 amount)
func (_L2Staking *L2StakingCaller) Commissions(opts *bind.CallOpts, staker common.Address) (struct {
	Rate   *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "commissions", staker)

	outstruct := new(struct {
		Rate   *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Commissions is a free data retrieval call binding the contract method 0x7b05afb5.
//
// Solidity: function commissions(address staker) view returns(uint256 rate, uint256 amount)
func (_L2Staking *L2StakingSession) Commissions(staker common.Address) (struct {
	Rate   *big.Int
	Amount *big.Int
}, error) {
	return _L2Staking.Contract.Commissions(&_L2Staking.CallOpts, staker)
}

// Commissions is a free data retrieval call binding the contract method 0x7b05afb5.
//
// Solidity: function commissions(address staker) view returns(uint256 rate, uint256 amount)
func (_L2Staking *L2StakingCallerSession) Commissions(staker common.Address) (struct {
	Rate   *big.Int
	Amount *big.Int
}, error) {
	return _L2Staking.Contract.Commissions(&_L2Staking.CallOpts, staker)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_L2Staking *L2StakingCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_L2Staking *L2StakingSession) CurrentEpoch() (*big.Int, error) {
	return _L2Staking.Contract.CurrentEpoch(&_L2Staking.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) CurrentEpoch() (*big.Int, error) {
	return _L2Staking.Contract.CurrentEpoch(&_L2Staking.CallOpts)
}

// DelegateeDelegations is a free data retrieval call binding the contract method 0x1d5611b8.
//
// Solidity: function delegateeDelegations(address staker) view returns(uint256 amount, uint256 share)
func (_L2Staking *L2StakingCaller) DelegateeDelegations(opts *bind.CallOpts, staker common.Address) (struct {
	Amount *big.Int
	Share  *big.Int
}, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "delegateeDelegations", staker)

	outstruct := new(struct {
		Amount *big.Int
		Share  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Share = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelegateeDelegations is a free data retrieval call binding the contract method 0x1d5611b8.
//
// Solidity: function delegateeDelegations(address staker) view returns(uint256 amount, uint256 share)
func (_L2Staking *L2StakingSession) DelegateeDelegations(staker common.Address) (struct {
	Amount *big.Int
	Share  *big.Int
}, error) {
	return _L2Staking.Contract.DelegateeDelegations(&_L2Staking.CallOpts, staker)
}

// DelegateeDelegations is a free data retrieval call binding the contract method 0x1d5611b8.
//
// Solidity: function delegateeDelegations(address staker) view returns(uint256 amount, uint256 share)
func (_L2Staking *L2StakingCallerSession) DelegateeDelegations(staker common.Address) (struct {
	Amount *big.Int
	Share  *big.Int
}, error) {
	return _L2Staking.Contract.DelegateeDelegations(&_L2Staking.CallOpts, staker)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x3b2713c5.
//
// Solidity: function delegatorDelegations(address staker, address delegator) view returns(uint256 share)
func (_L2Staking *L2StakingCaller) DelegatorDelegations(opts *bind.CallOpts, staker common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "delegatorDelegations", staker, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x3b2713c5.
//
// Solidity: function delegatorDelegations(address staker, address delegator) view returns(uint256 share)
func (_L2Staking *L2StakingSession) DelegatorDelegations(staker common.Address, delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.DelegatorDelegations(&_L2Staking.CallOpts, staker, delegator)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x3b2713c5.
//
// Solidity: function delegatorDelegations(address staker, address delegator) view returns(uint256 share)
func (_L2Staking *L2StakingCallerSession) DelegatorDelegations(staker common.Address, delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.DelegatorDelegations(&_L2Staking.CallOpts, staker, delegator)
}

// EpochSequencerBlocks is a free data retrieval call binding the contract method 0xfe348884.
//
// Solidity: function epochSequencerBlocks(address seequencer) view returns(uint256)
func (_L2Staking *L2StakingCaller) EpochSequencerBlocks(opts *bind.CallOpts, seequencer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "epochSequencerBlocks", seequencer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochSequencerBlocks is a free data retrieval call binding the contract method 0xfe348884.
//
// Solidity: function epochSequencerBlocks(address seequencer) view returns(uint256)
func (_L2Staking *L2StakingSession) EpochSequencerBlocks(seequencer common.Address) (*big.Int, error) {
	return _L2Staking.Contract.EpochSequencerBlocks(&_L2Staking.CallOpts, seequencer)
}

// EpochSequencerBlocks is a free data retrieval call binding the contract method 0xfe348884.
//
// Solidity: function epochSequencerBlocks(address seequencer) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) EpochSequencerBlocks(seequencer common.Address) (*big.Int, error) {
	return _L2Staking.Contract.EpochSequencerBlocks(&_L2Staking.CallOpts, seequencer)
}

// EpochTotalBlocks is a free data retrieval call binding the contract method 0xeefecafd.
//
// Solidity: function epochTotalBlocks() view returns(uint256)
func (_L2Staking *L2StakingCaller) EpochTotalBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "epochTotalBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochTotalBlocks is a free data retrieval call binding the contract method 0xeefecafd.
//
// Solidity: function epochTotalBlocks() view returns(uint256)
func (_L2Staking *L2StakingSession) EpochTotalBlocks() (*big.Int, error) {
	return _L2Staking.Contract.EpochTotalBlocks(&_L2Staking.CallOpts)
}

// EpochTotalBlocks is a free data retrieval call binding the contract method 0xeefecafd.
//
// Solidity: function epochTotalBlocks() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) EpochTotalBlocks() (*big.Int, error) {
	return _L2Staking.Contract.EpochTotalBlocks(&_L2Staking.CallOpts)
}

// GetAllDelegatorsInPagination is a free data retrieval call binding the contract method 0xd31d83d9.
//
// Solidity: function getAllDelegatorsInPagination(address staker, uint256 pageSize, uint256 pageIndex) view returns(uint256 delegatorsTotalNumber, address[] delegatorsInPage)
func (_L2Staking *L2StakingCaller) GetAllDelegatorsInPagination(opts *bind.CallOpts, staker common.Address, pageSize *big.Int, pageIndex *big.Int) (struct {
	DelegatorsTotalNumber *big.Int
	DelegatorsInPage      []common.Address
}, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "getAllDelegatorsInPagination", staker, pageSize, pageIndex)

	outstruct := new(struct {
		DelegatorsTotalNumber *big.Int
		DelegatorsInPage      []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DelegatorsTotalNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DelegatorsInPage = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetAllDelegatorsInPagination is a free data retrieval call binding the contract method 0xd31d83d9.
//
// Solidity: function getAllDelegatorsInPagination(address staker, uint256 pageSize, uint256 pageIndex) view returns(uint256 delegatorsTotalNumber, address[] delegatorsInPage)
func (_L2Staking *L2StakingSession) GetAllDelegatorsInPagination(staker common.Address, pageSize *big.Int, pageIndex *big.Int) (struct {
	DelegatorsTotalNumber *big.Int
	DelegatorsInPage      []common.Address
}, error) {
	return _L2Staking.Contract.GetAllDelegatorsInPagination(&_L2Staking.CallOpts, staker, pageSize, pageIndex)
}

// GetAllDelegatorsInPagination is a free data retrieval call binding the contract method 0xd31d83d9.
//
// Solidity: function getAllDelegatorsInPagination(address staker, uint256 pageSize, uint256 pageIndex) view returns(uint256 delegatorsTotalNumber, address[] delegatorsInPage)
func (_L2Staking *L2StakingCallerSession) GetAllDelegatorsInPagination(staker common.Address, pageSize *big.Int, pageIndex *big.Int) (struct {
	DelegatorsTotalNumber *big.Int
	DelegatorsInPage      []common.Address
}, error) {
	return _L2Staking.Contract.GetAllDelegatorsInPagination(&_L2Staking.CallOpts, staker, pageSize, pageIndex)
}

// GetDelegatorsLength is a free data retrieval call binding the contract method 0x0043b758.
//
// Solidity: function getDelegatorsLength(address staker) view returns(uint256)
func (_L2Staking *L2StakingCaller) GetDelegatorsLength(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "getDelegatorsLength", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegatorsLength is a free data retrieval call binding the contract method 0x0043b758.
//
// Solidity: function getDelegatorsLength(address staker) view returns(uint256)
func (_L2Staking *L2StakingSession) GetDelegatorsLength(staker common.Address) (*big.Int, error) {
	return _L2Staking.Contract.GetDelegatorsLength(&_L2Staking.CallOpts, staker)
}

// GetDelegatorsLength is a free data retrieval call binding the contract method 0x0043b758.
//
// Solidity: function getDelegatorsLength(address staker) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) GetDelegatorsLength(staker common.Address) (*big.Int, error) {
	return _L2Staking.Contract.GetDelegatorsLength(&_L2Staking.CallOpts, staker)
}

// GetStakerAddressesLength is a free data retrieval call binding the contract method 0x46cdc18a.
//
// Solidity: function getStakerAddressesLength() view returns(uint256)
func (_L2Staking *L2StakingCaller) GetStakerAddressesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "getStakerAddressesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakerAddressesLength is a free data retrieval call binding the contract method 0x46cdc18a.
//
// Solidity: function getStakerAddressesLength() view returns(uint256)
func (_L2Staking *L2StakingSession) GetStakerAddressesLength() (*big.Int, error) {
	return _L2Staking.Contract.GetStakerAddressesLength(&_L2Staking.CallOpts)
}

// GetStakerAddressesLength is a free data retrieval call binding the contract method 0x46cdc18a.
//
// Solidity: function getStakerAddressesLength() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) GetStakerAddressesLength() (*big.Int, error) {
	return _L2Staking.Contract.GetStakerAddressesLength(&_L2Staking.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingCaller) GetStakers(opts *bind.CallOpts) ([]TypesStakerInfo, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "getStakers")

	if err != nil {
		return *new([]TypesStakerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TypesStakerInfo)).(*[]TypesStakerInfo)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingSession) GetStakers() ([]TypesStakerInfo, error) {
	return _L2Staking.Contract.GetStakers(&_L2Staking.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingCallerSession) GetStakers() ([]TypesStakerInfo, error) {
	return _L2Staking.Contract.GetStakers(&_L2Staking.CallOpts)
}

// GetStakesInfo is a free data retrieval call binding the contract method 0x30158eea.
//
// Solidity: function getStakesInfo(address[] _stakerAddresses) view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingCaller) GetStakesInfo(opts *bind.CallOpts, _stakerAddresses []common.Address) ([]TypesStakerInfo, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "getStakesInfo", _stakerAddresses)

	if err != nil {
		return *new([]TypesStakerInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]TypesStakerInfo)).(*[]TypesStakerInfo)

	return out0, err

}

// GetStakesInfo is a free data retrieval call binding the contract method 0x30158eea.
//
// Solidity: function getStakesInfo(address[] _stakerAddresses) view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingSession) GetStakesInfo(_stakerAddresses []common.Address) ([]TypesStakerInfo, error) {
	return _L2Staking.Contract.GetStakesInfo(&_L2Staking.CallOpts, _stakerAddresses)
}

// GetStakesInfo is a free data retrieval call binding the contract method 0x30158eea.
//
// Solidity: function getStakesInfo(address[] _stakerAddresses) view returns((address,bytes32,bytes)[])
func (_L2Staking *L2StakingCallerSession) GetStakesInfo(_stakerAddresses []common.Address) ([]TypesStakerInfo, error) {
	return _L2Staking.Contract.GetStakesInfo(&_L2Staking.CallOpts, _stakerAddresses)
}

// IsStakingTo is a free data retrieval call binding the contract method 0x84d7d1d4.
//
// Solidity: function isStakingTo(address staker) view returns(bool)
func (_L2Staking *L2StakingCaller) IsStakingTo(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "isStakingTo", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingTo is a free data retrieval call binding the contract method 0x84d7d1d4.
//
// Solidity: function isStakingTo(address staker) view returns(bool)
func (_L2Staking *L2StakingSession) IsStakingTo(staker common.Address) (bool, error) {
	return _L2Staking.Contract.IsStakingTo(&_L2Staking.CallOpts, staker)
}

// IsStakingTo is a free data retrieval call binding the contract method 0x84d7d1d4.
//
// Solidity: function isStakingTo(address staker) view returns(bool)
func (_L2Staking *L2StakingCallerSession) IsStakingTo(staker common.Address) (bool, error) {
	return _L2Staking.Contract.IsStakingTo(&_L2Staking.CallOpts, staker)
}

// LatestSequencerSetSize is a free data retrieval call binding the contract method 0xf0261bc2.
//
// Solidity: function latestSequencerSetSize() view returns(uint256)
func (_L2Staking *L2StakingCaller) LatestSequencerSetSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "latestSequencerSetSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestSequencerSetSize is a free data retrieval call binding the contract method 0xf0261bc2.
//
// Solidity: function latestSequencerSetSize() view returns(uint256)
func (_L2Staking *L2StakingSession) LatestSequencerSetSize() (*big.Int, error) {
	return _L2Staking.Contract.LatestSequencerSetSize(&_L2Staking.CallOpts)
}

// LatestSequencerSetSize is a free data retrieval call binding the contract method 0xf0261bc2.
//
// Solidity: function latestSequencerSetSize() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) LatestSequencerSetSize() (*big.Int, error) {
	return _L2Staking.Contract.LatestSequencerSetSize(&_L2Staking.CallOpts)
}

// LockedAmount is a free data retrieval call binding the contract method 0xa61bb764.
//
// Solidity: function lockedAmount(address delegator, uint256 number) view returns(uint256)
func (_L2Staking *L2StakingCaller) LockedAmount(opts *bind.CallOpts, delegator common.Address, number *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "lockedAmount", delegator, number)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedAmount is a free data retrieval call binding the contract method 0xa61bb764.
//
// Solidity: function lockedAmount(address delegator, uint256 number) view returns(uint256)
func (_L2Staking *L2StakingSession) LockedAmount(delegator common.Address, number *big.Int) (*big.Int, error) {
	return _L2Staking.Contract.LockedAmount(&_L2Staking.CallOpts, delegator, number)
}

// LockedAmount is a free data retrieval call binding the contract method 0xa61bb764.
//
// Solidity: function lockedAmount(address delegator, uint256 number) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) LockedAmount(delegator common.Address, number *big.Int) (*big.Int, error) {
	return _L2Staking.Contract.LockedAmount(&_L2Staking.CallOpts, delegator, number)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Staking *L2StakingCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Staking *L2StakingSession) Messenger() (common.Address, error) {
	return _L2Staking.Contract.Messenger(&_L2Staking.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2Staking *L2StakingCallerSession) Messenger() (common.Address, error) {
	return _L2Staking.Contract.Messenger(&_L2Staking.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_L2Staking *L2StakingCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_L2Staking *L2StakingSession) Nonce() (*big.Int, error) {
	return _L2Staking.Contract.Nonce(&_L2Staking.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) Nonce() (*big.Int, error) {
	return _L2Staking.Contract.Nonce(&_L2Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2Staking *L2StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2Staking *L2StakingSession) Owner() (common.Address, error) {
	return _L2Staking.Contract.Owner(&_L2Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2Staking *L2StakingCallerSession) Owner() (common.Address, error) {
	return _L2Staking.Contract.Owner(&_L2Staking.CallOpts)
}

// PendingUndelegateRequest is a free data retrieval call binding the contract method 0x0321731c.
//
// Solidity: function pendingUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCaller) PendingUndelegateRequest(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "pendingUndelegateRequest", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingUndelegateRequest is a free data retrieval call binding the contract method 0x0321731c.
//
// Solidity: function pendingUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingSession) PendingUndelegateRequest(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.PendingUndelegateRequest(&_L2Staking.CallOpts, delegator)
}

// PendingUndelegateRequest is a free data retrieval call binding the contract method 0x0321731c.
//
// Solidity: function pendingUndelegateRequest(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) PendingUndelegateRequest(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.PendingUndelegateRequest(&_L2Staking.CallOpts, delegator)
}

// QueryDelegationAmount is a free data retrieval call binding the contract method 0x9d51c3b9.
//
// Solidity: function queryDelegationAmount(address delegatee, address delegator) view returns(uint256 amount)
func (_L2Staking *L2StakingCaller) QueryDelegationAmount(opts *bind.CallOpts, delegatee common.Address, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "queryDelegationAmount", delegatee, delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryDelegationAmount is a free data retrieval call binding the contract method 0x9d51c3b9.
//
// Solidity: function queryDelegationAmount(address delegatee, address delegator) view returns(uint256 amount)
func (_L2Staking *L2StakingSession) QueryDelegationAmount(delegatee common.Address, delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.QueryDelegationAmount(&_L2Staking.CallOpts, delegatee, delegator)
}

// QueryDelegationAmount is a free data retrieval call binding the contract method 0x9d51c3b9.
//
// Solidity: function queryDelegationAmount(address delegatee, address delegator) view returns(uint256 amount)
func (_L2Staking *L2StakingCallerSession) QueryDelegationAmount(delegatee common.Address, delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.QueryDelegationAmount(&_L2Staking.CallOpts, delegatee, delegator)
}

// QueryUnclaimedCommission is a free data retrieval call binding the contract method 0xbf2dca0a.
//
// Solidity: function queryUnclaimedCommission(address delegatee) view returns(uint256 amount)
func (_L2Staking *L2StakingCaller) QueryUnclaimedCommission(opts *bind.CallOpts, delegatee common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "queryUnclaimedCommission", delegatee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryUnclaimedCommission is a free data retrieval call binding the contract method 0xbf2dca0a.
//
// Solidity: function queryUnclaimedCommission(address delegatee) view returns(uint256 amount)
func (_L2Staking *L2StakingSession) QueryUnclaimedCommission(delegatee common.Address) (*big.Int, error) {
	return _L2Staking.Contract.QueryUnclaimedCommission(&_L2Staking.CallOpts, delegatee)
}

// QueryUnclaimedCommission is a free data retrieval call binding the contract method 0xbf2dca0a.
//
// Solidity: function queryUnclaimedCommission(address delegatee) view returns(uint256 amount)
func (_L2Staking *L2StakingCallerSession) QueryUnclaimedCommission(delegatee common.Address) (*big.Int, error) {
	return _L2Staking.Contract.QueryUnclaimedCommission(&_L2Staking.CallOpts, delegatee)
}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_L2Staking *L2StakingCaller) RewardStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "rewardStartTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_L2Staking *L2StakingSession) RewardStartTime() (*big.Int, error) {
	return _L2Staking.Contract.RewardStartTime(&_L2Staking.CallOpts)
}

// RewardStartTime is a free data retrieval call binding the contract method 0x2cc138be.
//
// Solidity: function rewardStartTime() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) RewardStartTime() (*big.Int, error) {
	return _L2Staking.Contract.RewardStartTime(&_L2Staking.CallOpts)
}

// RewardStarted is a free data retrieval call binding the contract method 0x96ab994d.
//
// Solidity: function rewardStarted() view returns(bool)
func (_L2Staking *L2StakingCaller) RewardStarted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "rewardStarted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardStarted is a free data retrieval call binding the contract method 0x96ab994d.
//
// Solidity: function rewardStarted() view returns(bool)
func (_L2Staking *L2StakingSession) RewardStarted() (bool, error) {
	return _L2Staking.Contract.RewardStarted(&_L2Staking.CallOpts)
}

// RewardStarted is a free data retrieval call binding the contract method 0x96ab994d.
//
// Solidity: function rewardStarted() view returns(bool)
func (_L2Staking *L2StakingCallerSession) RewardStarted() (bool, error) {
	return _L2Staking.Contract.RewardStarted(&_L2Staking.CallOpts)
}

// SequencerSetMaxSize is a free data retrieval call binding the contract method 0x2e787be3.
//
// Solidity: function sequencerSetMaxSize() view returns(uint256)
func (_L2Staking *L2StakingCaller) SequencerSetMaxSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "sequencerSetMaxSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerSetMaxSize is a free data retrieval call binding the contract method 0x2e787be3.
//
// Solidity: function sequencerSetMaxSize() view returns(uint256)
func (_L2Staking *L2StakingSession) SequencerSetMaxSize() (*big.Int, error) {
	return _L2Staking.Contract.SequencerSetMaxSize(&_L2Staking.CallOpts)
}

// SequencerSetMaxSize is a free data retrieval call binding the contract method 0x2e787be3.
//
// Solidity: function sequencerSetMaxSize() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) SequencerSetMaxSize() (*big.Int, error) {
	return _L2Staking.Contract.SequencerSetMaxSize(&_L2Staking.CallOpts)
}

// StakerAddresses is a free data retrieval call binding the contract method 0x459598a2.
//
// Solidity: function stakerAddresses(uint256 ) view returns(address)
func (_L2Staking *L2StakingCaller) StakerAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "stakerAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerAddresses is a free data retrieval call binding the contract method 0x459598a2.
//
// Solidity: function stakerAddresses(uint256 ) view returns(address)
func (_L2Staking *L2StakingSession) StakerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Staking.Contract.StakerAddresses(&_L2Staking.CallOpts, arg0)
}

// StakerAddresses is a free data retrieval call binding the contract method 0x459598a2.
//
// Solidity: function stakerAddresses(uint256 ) view returns(address)
func (_L2Staking *L2StakingCallerSession) StakerAddresses(arg0 *big.Int) (common.Address, error) {
	return _L2Staking.Contract.StakerAddresses(&_L2Staking.CallOpts, arg0)
}

// StakerRankings is a free data retrieval call binding the contract method 0xb5d2e0dc.
//
// Solidity: function stakerRankings(address staker) view returns(uint256 ranking)
func (_L2Staking *L2StakingCaller) StakerRankings(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "stakerRankings", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerRankings is a free data retrieval call binding the contract method 0xb5d2e0dc.
//
// Solidity: function stakerRankings(address staker) view returns(uint256 ranking)
func (_L2Staking *L2StakingSession) StakerRankings(staker common.Address) (*big.Int, error) {
	return _L2Staking.Contract.StakerRankings(&_L2Staking.CallOpts, staker)
}

// StakerRankings is a free data retrieval call binding the contract method 0xb5d2e0dc.
//
// Solidity: function stakerRankings(address staker) view returns(uint256 ranking)
func (_L2Staking *L2StakingCallerSession) StakerRankings(staker common.Address) (*big.Int, error) {
	return _L2Staking.Contract.StakerRankings(&_L2Staking.CallOpts, staker)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingCaller) Stakers(opts *bind.CallOpts, staker common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "stakers", staker)

	outstruct := new(struct {
		Addr   common.Address
		TmKey  [32]byte
		BlsKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TmKey = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.BlsKey = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingSession) Stakers(staker common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Staking.Contract.Stakers(&_L2Staking.CallOpts, staker)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address staker) view returns(address addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingCallerSession) Stakers(staker common.Address) (struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
}, error) {
	return _L2Staking.Contract.Stakers(&_L2Staking.CallOpts, staker)
}

// UndelegateLockEpochs is a free data retrieval call binding the contract method 0x12a3e947.
//
// Solidity: function undelegateLockEpochs() view returns(uint256)
func (_L2Staking *L2StakingCaller) UndelegateLockEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "undelegateLockEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegateLockEpochs is a free data retrieval call binding the contract method 0x12a3e947.
//
// Solidity: function undelegateLockEpochs() view returns(uint256)
func (_L2Staking *L2StakingSession) UndelegateLockEpochs() (*big.Int, error) {
	return _L2Staking.Contract.UndelegateLockEpochs(&_L2Staking.CallOpts)
}

// UndelegateLockEpochs is a free data retrieval call binding the contract method 0x12a3e947.
//
// Solidity: function undelegateLockEpochs() view returns(uint256)
func (_L2Staking *L2StakingCallerSession) UndelegateLockEpochs() (*big.Int, error) {
	return _L2Staking.Contract.UndelegateLockEpochs(&_L2Staking.CallOpts)
}

// UndelegateRequest is a free data retrieval call binding the contract method 0xb7a587bf.
//
// Solidity: function undelegateRequest(address delegator, uint256 _index) view returns((uint256,uint256))
func (_L2Staking *L2StakingCaller) UndelegateRequest(opts *bind.CallOpts, delegator common.Address, _index *big.Int) (IL2StakingUndelegateRequest, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "undelegateRequest", delegator, _index)

	if err != nil {
		return *new(IL2StakingUndelegateRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IL2StakingUndelegateRequest)).(*IL2StakingUndelegateRequest)

	return out0, err

}

// UndelegateRequest is a free data retrieval call binding the contract method 0xb7a587bf.
//
// Solidity: function undelegateRequest(address delegator, uint256 _index) view returns((uint256,uint256))
func (_L2Staking *L2StakingSession) UndelegateRequest(delegator common.Address, _index *big.Int) (IL2StakingUndelegateRequest, error) {
	return _L2Staking.Contract.UndelegateRequest(&_L2Staking.CallOpts, delegator, _index)
}

// UndelegateRequest is a free data retrieval call binding the contract method 0xb7a587bf.
//
// Solidity: function undelegateRequest(address delegator, uint256 _index) view returns((uint256,uint256))
func (_L2Staking *L2StakingCallerSession) UndelegateRequest(delegator common.Address, _index *big.Int) (IL2StakingUndelegateRequest, error) {
	return _L2Staking.Contract.UndelegateRequest(&_L2Staking.CallOpts, delegator, _index)
}

// UndelegateSequence is a free data retrieval call binding the contract method 0x7c7e8bd2.
//
// Solidity: function undelegateSequence(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCaller) UndelegateSequence(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2Staking.contract.Call(opts, &out, "undelegateSequence", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegateSequence is a free data retrieval call binding the contract method 0x7c7e8bd2.
//
// Solidity: function undelegateSequence(address delegator) view returns(uint256)
func (_L2Staking *L2StakingSession) UndelegateSequence(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.UndelegateSequence(&_L2Staking.CallOpts, delegator)
}

// UndelegateSequence is a free data retrieval call binding the contract method 0x7c7e8bd2.
//
// Solidity: function undelegateSequence(address delegator) view returns(uint256)
func (_L2Staking *L2StakingCallerSession) UndelegateSequence(delegator common.Address) (*big.Int, error) {
	return _L2Staking.Contract.UndelegateSequence(&_L2Staking.CallOpts, delegator)
}

// AddStaker is a paid mutator transaction binding the contract method 0x7046529b.
//
// Solidity: function addStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingTransactor) AddStaker(opts *bind.TransactOpts, _nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "addStaker", _nonce, add)
}

// AddStaker is a paid mutator transaction binding the contract method 0x7046529b.
//
// Solidity: function addStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingSession) AddStaker(_nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.AddStaker(&_L2Staking.TransactOpts, _nonce, add)
}

// AddStaker is a paid mutator transaction binding the contract method 0x7046529b.
//
// Solidity: function addStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingTransactorSession) AddStaker(_nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.AddStaker(&_L2Staking.TransactOpts, _nonce, add)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns()
func (_L2Staking *L2StakingTransactor) ClaimCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "claimCommission")
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns()
func (_L2Staking *L2StakingSession) ClaimCommission() (*types.Transaction, error) {
	return _L2Staking.Contract.ClaimCommission(&_L2Staking.TransactOpts)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0xfad99f98.
//
// Solidity: function claimCommission() returns()
func (_L2Staking *L2StakingTransactorSession) ClaimCommission() (*types.Transaction, error) {
	return _L2Staking.Contract.ClaimCommission(&_L2Staking.TransactOpts)
}

// ClaimUndelegation is a paid mutator transaction binding the contract method 0x201018fb.
//
// Solidity: function claimUndelegation(uint256 number) returns(uint256)
func (_L2Staking *L2StakingTransactor) ClaimUndelegation(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "claimUndelegation", number)
}

// ClaimUndelegation is a paid mutator transaction binding the contract method 0x201018fb.
//
// Solidity: function claimUndelegation(uint256 number) returns(uint256)
func (_L2Staking *L2StakingSession) ClaimUndelegation(number *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.ClaimUndelegation(&_L2Staking.TransactOpts, number)
}

// ClaimUndelegation is a paid mutator transaction binding the contract method 0x201018fb.
//
// Solidity: function claimUndelegation(uint256 number) returns(uint256)
func (_L2Staking *L2StakingTransactorSession) ClaimUndelegation(number *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.ClaimUndelegation(&_L2Staking.TransactOpts, number)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "delegate", delegatee, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingSession) Delegate(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Delegate(&_L2Staking.TransactOpts, delegatee, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingTransactorSession) Delegate(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Delegate(&_L2Staking.TransactOpts, delegatee, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(uint256 amount) returns()
func (_L2Staking *L2StakingTransactor) Distribute(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "distribute", amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(uint256 amount) returns()
func (_L2Staking *L2StakingSession) Distribute(amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Distribute(&_L2Staking.TransactOpts, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x91c05b0b.
//
// Solidity: function distribute(uint256 amount) returns()
func (_L2Staking *L2StakingTransactorSession) Distribute(amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Distribute(&_L2Staking.TransactOpts, amount)
}

// EmergencyAddStaker is a paid mutator transaction binding the contract method 0x009c6f0c.
//
// Solidity: function emergencyAddStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingTransactor) EmergencyAddStaker(opts *bind.TransactOpts, _nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "emergencyAddStaker", _nonce, add)
}

// EmergencyAddStaker is a paid mutator transaction binding the contract method 0x009c6f0c.
//
// Solidity: function emergencyAddStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingSession) EmergencyAddStaker(_nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.EmergencyAddStaker(&_L2Staking.TransactOpts, _nonce, add)
}

// EmergencyAddStaker is a paid mutator transaction binding the contract method 0x009c6f0c.
//
// Solidity: function emergencyAddStaker(uint256 _nonce, (address,bytes32,bytes) add) returns()
func (_L2Staking *L2StakingTransactorSession) EmergencyAddStaker(_nonce *big.Int, add TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.EmergencyAddStaker(&_L2Staking.TransactOpts, _nonce, add)
}

// EmergencyRemoveStakers is a paid mutator transaction binding the contract method 0xfc6facc6.
//
// Solidity: function emergencyRemoveStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingTransactor) EmergencyRemoveStakers(opts *bind.TransactOpts, _nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "emergencyRemoveStakers", _nonce, remove)
}

// EmergencyRemoveStakers is a paid mutator transaction binding the contract method 0xfc6facc6.
//
// Solidity: function emergencyRemoveStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingSession) EmergencyRemoveStakers(_nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.EmergencyRemoveStakers(&_L2Staking.TransactOpts, _nonce, remove)
}

// EmergencyRemoveStakers is a paid mutator transaction binding the contract method 0xfc6facc6.
//
// Solidity: function emergencyRemoveStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingTransactorSession) EmergencyRemoveStakers(_nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.EmergencyRemoveStakers(&_L2Staking.TransactOpts, _nonce, remove)
}

// Initialize is a paid mutator transaction binding the contract method 0x439162b5.
//
// Solidity: function initialize(address _owner, uint256 _sequencersMaxSize, uint256 _undelegateLockEpochs, uint256 _rewardStartTime, (address,bytes32,bytes)[] _stakers) returns()
func (_L2Staking *L2StakingTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _sequencersMaxSize *big.Int, _undelegateLockEpochs *big.Int, _rewardStartTime *big.Int, _stakers []TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "initialize", _owner, _sequencersMaxSize, _undelegateLockEpochs, _rewardStartTime, _stakers)
}

// Initialize is a paid mutator transaction binding the contract method 0x439162b5.
//
// Solidity: function initialize(address _owner, uint256 _sequencersMaxSize, uint256 _undelegateLockEpochs, uint256 _rewardStartTime, (address,bytes32,bytes)[] _stakers) returns()
func (_L2Staking *L2StakingSession) Initialize(_owner common.Address, _sequencersMaxSize *big.Int, _undelegateLockEpochs *big.Int, _rewardStartTime *big.Int, _stakers []TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.Initialize(&_L2Staking.TransactOpts, _owner, _sequencersMaxSize, _undelegateLockEpochs, _rewardStartTime, _stakers)
}

// Initialize is a paid mutator transaction binding the contract method 0x439162b5.
//
// Solidity: function initialize(address _owner, uint256 _sequencersMaxSize, uint256 _undelegateLockEpochs, uint256 _rewardStartTime, (address,bytes32,bytes)[] _stakers) returns()
func (_L2Staking *L2StakingTransactorSession) Initialize(_owner common.Address, _sequencersMaxSize *big.Int, _undelegateLockEpochs *big.Int, _rewardStartTime *big.Int, _stakers []TypesStakerInfo) (*types.Transaction, error) {
	return _L2Staking.Contract.Initialize(&_L2Staking.TransactOpts, _owner, _sequencersMaxSize, _undelegateLockEpochs, _rewardStartTime, _stakers)
}

// RecordBlocks is a paid mutator transaction binding the contract method 0xff4840cd.
//
// Solidity: function recordBlocks(address sequencerAddr) returns()
func (_L2Staking *L2StakingTransactor) RecordBlocks(opts *bind.TransactOpts, sequencerAddr common.Address) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "recordBlocks", sequencerAddr)
}

// RecordBlocks is a paid mutator transaction binding the contract method 0xff4840cd.
//
// Solidity: function recordBlocks(address sequencerAddr) returns()
func (_L2Staking *L2StakingSession) RecordBlocks(sequencerAddr common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.RecordBlocks(&_L2Staking.TransactOpts, sequencerAddr)
}

// RecordBlocks is a paid mutator transaction binding the contract method 0xff4840cd.
//
// Solidity: function recordBlocks(address sequencerAddr) returns()
func (_L2Staking *L2StakingTransactorSession) RecordBlocks(sequencerAddr common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.RecordBlocks(&_L2Staking.TransactOpts, sequencerAddr)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address delegateeFrom, address delegateeTo, uint256 amount) returns()
func (_L2Staking *L2StakingTransactor) Redelegate(opts *bind.TransactOpts, delegateeFrom common.Address, delegateeTo common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "redelegate", delegateeFrom, delegateeTo, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address delegateeFrom, address delegateeTo, uint256 amount) returns()
func (_L2Staking *L2StakingSession) Redelegate(delegateeFrom common.Address, delegateeTo common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Redelegate(&_L2Staking.TransactOpts, delegateeFrom, delegateeTo, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address delegateeFrom, address delegateeTo, uint256 amount) returns()
func (_L2Staking *L2StakingTransactorSession) Redelegate(delegateeFrom common.Address, delegateeTo common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Redelegate(&_L2Staking.TransactOpts, delegateeFrom, delegateeTo, amount)
}

// RemoveStakers is a paid mutator transaction binding the contract method 0xcce6cf9f.
//
// Solidity: function removeStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingTransactor) RemoveStakers(opts *bind.TransactOpts, _nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "removeStakers", _nonce, remove)
}

// RemoveStakers is a paid mutator transaction binding the contract method 0xcce6cf9f.
//
// Solidity: function removeStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingSession) RemoveStakers(_nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.RemoveStakers(&_L2Staking.TransactOpts, _nonce, remove)
}

// RemoveStakers is a paid mutator transaction binding the contract method 0xcce6cf9f.
//
// Solidity: function removeStakers(uint256 _nonce, address[] remove) returns()
func (_L2Staking *L2StakingTransactorSession) RemoveStakers(_nonce *big.Int, remove []common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.RemoveStakers(&_L2Staking.TransactOpts, _nonce, remove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2Staking *L2StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2Staking *L2StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2Staking.Contract.RenounceOwnership(&_L2Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2Staking *L2StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2Staking.Contract.RenounceOwnership(&_L2Staking.TransactOpts)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 rate) returns()
func (_L2Staking *L2StakingTransactor) SetCommissionRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "setCommissionRate", rate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 rate) returns()
func (_L2Staking *L2StakingSession) SetCommissionRate(rate *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.SetCommissionRate(&_L2Staking.TransactOpts, rate)
}

// SetCommissionRate is a paid mutator transaction binding the contract method 0x19fac8fd.
//
// Solidity: function setCommissionRate(uint256 rate) returns()
func (_L2Staking *L2StakingTransactorSession) SetCommissionRate(rate *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.SetCommissionRate(&_L2Staking.TransactOpts, rate)
}

// StartReward is a paid mutator transaction binding the contract method 0x746c8ae1.
//
// Solidity: function startReward() returns()
func (_L2Staking *L2StakingTransactor) StartReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "startReward")
}

// StartReward is a paid mutator transaction binding the contract method 0x746c8ae1.
//
// Solidity: function startReward() returns()
func (_L2Staking *L2StakingSession) StartReward() (*types.Transaction, error) {
	return _L2Staking.Contract.StartReward(&_L2Staking.TransactOpts)
}

// StartReward is a paid mutator transaction binding the contract method 0x746c8ae1.
//
// Solidity: function startReward() returns()
func (_L2Staking *L2StakingTransactorSession) StartReward() (*types.Transaction, error) {
	return _L2Staking.Contract.StartReward(&_L2Staking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2Staking *L2StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2Staking *L2StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.TransferOwnership(&_L2Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2Staking *L2StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2Staking.Contract.TransferOwnership(&_L2Staking.TransactOpts, newOwner)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingTransactor) Undelegate(opts *bind.TransactOpts, delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "undelegate", delegatee, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingSession) Undelegate(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Undelegate(&_L2Staking.TransactOpts, delegatee, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address delegatee, uint256 amount) returns()
func (_L2Staking *L2StakingTransactorSession) Undelegate(delegatee common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.Undelegate(&_L2Staking.TransactOpts, delegatee, amount)
}

// UpdateRewardStartTime is a paid mutator transaction binding the contract method 0x40b5c837.
//
// Solidity: function updateRewardStartTime(uint256 _rewardStartTime) returns()
func (_L2Staking *L2StakingTransactor) UpdateRewardStartTime(opts *bind.TransactOpts, _rewardStartTime *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "updateRewardStartTime", _rewardStartTime)
}

// UpdateRewardStartTime is a paid mutator transaction binding the contract method 0x40b5c837.
//
// Solidity: function updateRewardStartTime(uint256 _rewardStartTime) returns()
func (_L2Staking *L2StakingSession) UpdateRewardStartTime(_rewardStartTime *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.UpdateRewardStartTime(&_L2Staking.TransactOpts, _rewardStartTime)
}

// UpdateRewardStartTime is a paid mutator transaction binding the contract method 0x40b5c837.
//
// Solidity: function updateRewardStartTime(uint256 _rewardStartTime) returns()
func (_L2Staking *L2StakingTransactorSession) UpdateRewardStartTime(_rewardStartTime *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.UpdateRewardStartTime(&_L2Staking.TransactOpts, _rewardStartTime)
}

// UpdateSequencerSetMaxSize is a paid mutator transaction binding the contract method 0x0eb573af.
//
// Solidity: function updateSequencerSetMaxSize(uint256 _sequencerSetMaxSize) returns()
func (_L2Staking *L2StakingTransactor) UpdateSequencerSetMaxSize(opts *bind.TransactOpts, _sequencerSetMaxSize *big.Int) (*types.Transaction, error) {
	return _L2Staking.contract.Transact(opts, "updateSequencerSetMaxSize", _sequencerSetMaxSize)
}

// UpdateSequencerSetMaxSize is a paid mutator transaction binding the contract method 0x0eb573af.
//
// Solidity: function updateSequencerSetMaxSize(uint256 _sequencerSetMaxSize) returns()
func (_L2Staking *L2StakingSession) UpdateSequencerSetMaxSize(_sequencerSetMaxSize *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.UpdateSequencerSetMaxSize(&_L2Staking.TransactOpts, _sequencerSetMaxSize)
}

// UpdateSequencerSetMaxSize is a paid mutator transaction binding the contract method 0x0eb573af.
//
// Solidity: function updateSequencerSetMaxSize(uint256 _sequencerSetMaxSize) returns()
func (_L2Staking *L2StakingTransactorSession) UpdateSequencerSetMaxSize(_sequencerSetMaxSize *big.Int) (*types.Transaction, error) {
	return _L2Staking.Contract.UpdateSequencerSetMaxSize(&_L2Staking.TransactOpts, _sequencerSetMaxSize)
}

// L2StakingCommissionClaimedIterator is returned from FilterCommissionClaimed and is used to iterate over the raw logs and unpacked data for CommissionClaimed events raised by the L2Staking contract.
type L2StakingCommissionClaimedIterator struct {
	Event *L2StakingCommissionClaimed // Event containing the contract specifics and raw log

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
func (it *L2StakingCommissionClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingCommissionClaimed)
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
		it.Event = new(L2StakingCommissionClaimed)
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
func (it *L2StakingCommissionClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingCommissionClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingCommissionClaimed represents a CommissionClaimed event raised by the L2Staking contract.
type L2StakingCommissionClaimed struct {
	Delegatee common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionClaimed is a free log retrieval operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 amount)
func (_L2Staking *L2StakingFilterer) FilterCommissionClaimed(opts *bind.FilterOpts, delegatee []common.Address) (*L2StakingCommissionClaimedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "CommissionClaimed", delegateeRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingCommissionClaimedIterator{contract: _L2Staking.contract, event: "CommissionClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionClaimed is a free log subscription operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 amount)
func (_L2Staking *L2StakingFilterer) WatchCommissionClaimed(opts *bind.WatchOpts, sink chan<- *L2StakingCommissionClaimed, delegatee []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "CommissionClaimed", delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingCommissionClaimed)
				if err := _L2Staking.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
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

// ParseCommissionClaimed is a log parse operation binding the contract event 0x8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e411.
//
// Solidity: event CommissionClaimed(address indexed delegatee, uint256 amount)
func (_L2Staking *L2StakingFilterer) ParseCommissionClaimed(log types.Log) (*L2StakingCommissionClaimed, error) {
	event := new(L2StakingCommissionClaimed)
	if err := _L2Staking.contract.UnpackLog(event, "CommissionClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingCommissionUpdatedIterator is returned from FilterCommissionUpdated and is used to iterate over the raw logs and unpacked data for CommissionUpdated events raised by the L2Staking contract.
type L2StakingCommissionUpdatedIterator struct {
	Event *L2StakingCommissionUpdated // Event containing the contract specifics and raw log

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
func (it *L2StakingCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingCommissionUpdated)
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
		it.Event = new(L2StakingCommissionUpdated)
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
func (it *L2StakingCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingCommissionUpdated represents a CommissionUpdated event raised by the L2Staking contract.
type L2StakingCommissionUpdated struct {
	Staker  common.Address
	OldRate *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCommissionUpdated is a free log retrieval operation binding the contract event 0x6e500db30ce535d38852e318f333e9be41a3fec6c65d234ebb06203c896db9a5.
//
// Solidity: event CommissionUpdated(address indexed staker, uint256 oldRate, uint256 newRate)
func (_L2Staking *L2StakingFilterer) FilterCommissionUpdated(opts *bind.FilterOpts, staker []common.Address) (*L2StakingCommissionUpdatedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "CommissionUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingCommissionUpdatedIterator{contract: _L2Staking.contract, event: "CommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionUpdated is a free log subscription operation binding the contract event 0x6e500db30ce535d38852e318f333e9be41a3fec6c65d234ebb06203c896db9a5.
//
// Solidity: event CommissionUpdated(address indexed staker, uint256 oldRate, uint256 newRate)
func (_L2Staking *L2StakingFilterer) WatchCommissionUpdated(opts *bind.WatchOpts, sink chan<- *L2StakingCommissionUpdated, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "CommissionUpdated", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingCommissionUpdated)
				if err := _L2Staking.contract.UnpackLog(event, "CommissionUpdated", log); err != nil {
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

// ParseCommissionUpdated is a log parse operation binding the contract event 0x6e500db30ce535d38852e318f333e9be41a3fec6c65d234ebb06203c896db9a5.
//
// Solidity: event CommissionUpdated(address indexed staker, uint256 oldRate, uint256 newRate)
func (_L2Staking *L2StakingFilterer) ParseCommissionUpdated(log types.Log) (*L2StakingCommissionUpdated, error) {
	event := new(L2StakingCommissionUpdated)
	if err := _L2Staking.contract.UnpackLog(event, "CommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the L2Staking contract.
type L2StakingDelegatedIterator struct {
	Event *L2StakingDelegated // Event containing the contract specifics and raw log

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
func (it *L2StakingDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingDelegated)
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
		it.Event = new(L2StakingDelegated)
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
func (it *L2StakingDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingDelegated represents a Delegated event raised by the L2Staking contract.
type L2StakingDelegated struct {
	Delegatee       common.Address
	Delegator       common.Address
	Amount          *big.Int
	DelegateeAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount)
func (_L2Staking *L2StakingFilterer) FilterDelegated(opts *bind.FilterOpts, delegatee []common.Address, delegator []common.Address) (*L2StakingDelegatedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "Delegated", delegateeRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingDelegatedIterator{contract: _L2Staking.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount)
func (_L2Staking *L2StakingFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *L2StakingDelegated, delegatee []common.Address, delegator []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "Delegated", delegateeRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingDelegated)
				if err := _L2Staking.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0x24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04.
//
// Solidity: event Delegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount)
func (_L2Staking *L2StakingFilterer) ParseDelegated(log types.Log) (*L2StakingDelegated, error) {
	event := new(L2StakingDelegated)
	if err := _L2Staking.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingDistributedIterator is returned from FilterDistributed and is used to iterate over the raw logs and unpacked data for Distributed events raised by the L2Staking contract.
type L2StakingDistributedIterator struct {
	Event *L2StakingDistributed // Event containing the contract specifics and raw log

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
func (it *L2StakingDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingDistributed)
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
		it.Event = new(L2StakingDistributed)
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
func (it *L2StakingDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingDistributed represents a Distributed event raised by the L2Staking contract.
type L2StakingDistributed struct {
	Sequencer        common.Address
	DelegatorReward  *big.Int
	CommissionAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDistributed is a free log retrieval operation binding the contract event 0x60ce3cc2d133631eac66a476f14997a9fa682bd05a60dd993cf02285822d78d8.
//
// Solidity: event Distributed(address indexed sequencer, uint256 delegatorReward, uint256 commissionAmount)
func (_L2Staking *L2StakingFilterer) FilterDistributed(opts *bind.FilterOpts, sequencer []common.Address) (*L2StakingDistributedIterator, error) {

	var sequencerRule []interface{}
	for _, sequencerItem := range sequencer {
		sequencerRule = append(sequencerRule, sequencerItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "Distributed", sequencerRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingDistributedIterator{contract: _L2Staking.contract, event: "Distributed", logs: logs, sub: sub}, nil
}

// WatchDistributed is a free log subscription operation binding the contract event 0x60ce3cc2d133631eac66a476f14997a9fa682bd05a60dd993cf02285822d78d8.
//
// Solidity: event Distributed(address indexed sequencer, uint256 delegatorReward, uint256 commissionAmount)
func (_L2Staking *L2StakingFilterer) WatchDistributed(opts *bind.WatchOpts, sink chan<- *L2StakingDistributed, sequencer []common.Address) (event.Subscription, error) {

	var sequencerRule []interface{}
	for _, sequencerItem := range sequencer {
		sequencerRule = append(sequencerRule, sequencerItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "Distributed", sequencerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingDistributed)
				if err := _L2Staking.contract.UnpackLog(event, "Distributed", log); err != nil {
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

// ParseDistributed is a log parse operation binding the contract event 0x60ce3cc2d133631eac66a476f14997a9fa682bd05a60dd993cf02285822d78d8.
//
// Solidity: event Distributed(address indexed sequencer, uint256 delegatorReward, uint256 commissionAmount)
func (_L2Staking *L2StakingFilterer) ParseDistributed(log types.Log) (*L2StakingDistributed, error) {
	event := new(L2StakingDistributed)
	if err := _L2Staking.contract.UnpackLog(event, "Distributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2Staking contract.
type L2StakingInitializedIterator struct {
	Event *L2StakingInitialized // Event containing the contract specifics and raw log

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
func (it *L2StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingInitialized)
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
		it.Event = new(L2StakingInitialized)
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
func (it *L2StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingInitialized represents a Initialized event raised by the L2Staking contract.
type L2StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2Staking *L2StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2StakingInitializedIterator, error) {

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2StakingInitializedIterator{contract: _L2Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2Staking *L2StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingInitialized)
				if err := _L2Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L2Staking *L2StakingFilterer) ParseInitialized(log types.Log) (*L2StakingInitialized, error) {
	event := new(L2StakingInitialized)
	if err := _L2Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2Staking contract.
type L2StakingOwnershipTransferredIterator struct {
	Event *L2StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingOwnershipTransferred)
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
		it.Event = new(L2StakingOwnershipTransferred)
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
func (it *L2StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingOwnershipTransferred represents a OwnershipTransferred event raised by the L2Staking contract.
type L2StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2Staking *L2StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingOwnershipTransferredIterator{contract: _L2Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2Staking *L2StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingOwnershipTransferred)
				if err := _L2Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2Staking *L2StakingFilterer) ParseOwnershipTransferred(log types.Log) (*L2StakingOwnershipTransferred, error) {
	event := new(L2StakingOwnershipTransferred)
	if err := _L2Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the L2Staking contract.
type L2StakingRedelegatedIterator struct {
	Event *L2StakingRedelegated // Event containing the contract specifics and raw log

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
func (it *L2StakingRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingRedelegated)
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
		it.Event = new(L2StakingRedelegated)
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
func (it *L2StakingRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingRedelegated represents a Redelegated event raised by the L2Staking contract.
type L2StakingRedelegated struct {
	DelegateeFrom       common.Address
	DelegateeTo         common.Address
	Delegator           common.Address
	Amount              *big.Int
	DelegateeAmountFrom *big.Int
	DelegateeAmountTo   *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed delegateeFrom, address indexed delegateeTo, address indexed delegator, uint256 amount, uint256 delegateeAmountFrom, uint256 delegateeAmountTo)
func (_L2Staking *L2StakingFilterer) FilterRedelegated(opts *bind.FilterOpts, delegateeFrom []common.Address, delegateeTo []common.Address, delegator []common.Address) (*L2StakingRedelegatedIterator, error) {

	var delegateeFromRule []interface{}
	for _, delegateeFromItem := range delegateeFrom {
		delegateeFromRule = append(delegateeFromRule, delegateeFromItem)
	}
	var delegateeToRule []interface{}
	for _, delegateeToItem := range delegateeTo {
		delegateeToRule = append(delegateeToRule, delegateeToItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "Redelegated", delegateeFromRule, delegateeToRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingRedelegatedIterator{contract: _L2Staking.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed delegateeFrom, address indexed delegateeTo, address indexed delegator, uint256 amount, uint256 delegateeAmountFrom, uint256 delegateeAmountTo)
func (_L2Staking *L2StakingFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *L2StakingRedelegated, delegateeFrom []common.Address, delegateeTo []common.Address, delegator []common.Address) (event.Subscription, error) {

	var delegateeFromRule []interface{}
	for _, delegateeFromItem := range delegateeFrom {
		delegateeFromRule = append(delegateeFromRule, delegateeFromItem)
	}
	var delegateeToRule []interface{}
	for _, delegateeToItem := range delegateeTo {
		delegateeToRule = append(delegateeToRule, delegateeToItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "Redelegated", delegateeFromRule, delegateeToRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingRedelegated)
				if err := _L2Staking.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0xfdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b4.
//
// Solidity: event Redelegated(address indexed delegateeFrom, address indexed delegateeTo, address indexed delegator, uint256 amount, uint256 delegateeAmountFrom, uint256 delegateeAmountTo)
func (_L2Staking *L2StakingFilterer) ParseRedelegated(log types.Log) (*L2StakingRedelegated, error) {
	event := new(L2StakingRedelegated)
	if err := _L2Staking.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingRewardStartTimeUpdatedIterator is returned from FilterRewardStartTimeUpdated and is used to iterate over the raw logs and unpacked data for RewardStartTimeUpdated events raised by the L2Staking contract.
type L2StakingRewardStartTimeUpdatedIterator struct {
	Event *L2StakingRewardStartTimeUpdated // Event containing the contract specifics and raw log

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
func (it *L2StakingRewardStartTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingRewardStartTimeUpdated)
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
		it.Event = new(L2StakingRewardStartTimeUpdated)
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
func (it *L2StakingRewardStartTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingRewardStartTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingRewardStartTimeUpdated represents a RewardStartTimeUpdated event raised by the L2Staking contract.
type L2StakingRewardStartTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardStartTimeUpdated is a free log retrieval operation binding the contract event 0x91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555.
//
// Solidity: event RewardStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_L2Staking *L2StakingFilterer) FilterRewardStartTimeUpdated(opts *bind.FilterOpts) (*L2StakingRewardStartTimeUpdatedIterator, error) {

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "RewardStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &L2StakingRewardStartTimeUpdatedIterator{contract: _L2Staking.contract, event: "RewardStartTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardStartTimeUpdated is a free log subscription operation binding the contract event 0x91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555.
//
// Solidity: event RewardStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_L2Staking *L2StakingFilterer) WatchRewardStartTimeUpdated(opts *bind.WatchOpts, sink chan<- *L2StakingRewardStartTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "RewardStartTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingRewardStartTimeUpdated)
				if err := _L2Staking.contract.UnpackLog(event, "RewardStartTimeUpdated", log); err != nil {
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

// ParseRewardStartTimeUpdated is a log parse operation binding the contract event 0x91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555.
//
// Solidity: event RewardStartTimeUpdated(uint256 oldTime, uint256 newTime)
func (_L2Staking *L2StakingFilterer) ParseRewardStartTimeUpdated(log types.Log) (*L2StakingRewardStartTimeUpdated, error) {
	event := new(L2StakingRewardStartTimeUpdated)
	if err := _L2Staking.contract.UnpackLog(event, "RewardStartTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingSequencerSetMaxSizeUpdatedIterator is returned from FilterSequencerSetMaxSizeUpdated and is used to iterate over the raw logs and unpacked data for SequencerSetMaxSizeUpdated events raised by the L2Staking contract.
type L2StakingSequencerSetMaxSizeUpdatedIterator struct {
	Event *L2StakingSequencerSetMaxSizeUpdated // Event containing the contract specifics and raw log

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
func (it *L2StakingSequencerSetMaxSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingSequencerSetMaxSizeUpdated)
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
		it.Event = new(L2StakingSequencerSetMaxSizeUpdated)
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
func (it *L2StakingSequencerSetMaxSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingSequencerSetMaxSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingSequencerSetMaxSizeUpdated represents a SequencerSetMaxSizeUpdated event raised by the L2Staking contract.
type L2StakingSequencerSetMaxSizeUpdated struct {
	OldSize *big.Int
	NewSize *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSequencerSetMaxSizeUpdated is a free log retrieval operation binding the contract event 0x98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed.
//
// Solidity: event SequencerSetMaxSizeUpdated(uint256 oldSize, uint256 newSize)
func (_L2Staking *L2StakingFilterer) FilterSequencerSetMaxSizeUpdated(opts *bind.FilterOpts) (*L2StakingSequencerSetMaxSizeUpdatedIterator, error) {

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "SequencerSetMaxSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &L2StakingSequencerSetMaxSizeUpdatedIterator{contract: _L2Staking.contract, event: "SequencerSetMaxSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencerSetMaxSizeUpdated is a free log subscription operation binding the contract event 0x98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed.
//
// Solidity: event SequencerSetMaxSizeUpdated(uint256 oldSize, uint256 newSize)
func (_L2Staking *L2StakingFilterer) WatchSequencerSetMaxSizeUpdated(opts *bind.WatchOpts, sink chan<- *L2StakingSequencerSetMaxSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "SequencerSetMaxSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingSequencerSetMaxSizeUpdated)
				if err := _L2Staking.contract.UnpackLog(event, "SequencerSetMaxSizeUpdated", log); err != nil {
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

// ParseSequencerSetMaxSizeUpdated is a log parse operation binding the contract event 0x98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed.
//
// Solidity: event SequencerSetMaxSizeUpdated(uint256 oldSize, uint256 newSize)
func (_L2Staking *L2StakingFilterer) ParseSequencerSetMaxSizeUpdated(log types.Log) (*L2StakingSequencerSetMaxSizeUpdated, error) {
	event := new(L2StakingSequencerSetMaxSizeUpdated)
	if err := _L2Staking.contract.UnpackLog(event, "SequencerSetMaxSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingStakerAddedIterator is returned from FilterStakerAdded and is used to iterate over the raw logs and unpacked data for StakerAdded events raised by the L2Staking contract.
type L2StakingStakerAddedIterator struct {
	Event *L2StakingStakerAdded // Event containing the contract specifics and raw log

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
func (it *L2StakingStakerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingStakerAdded)
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
		it.Event = new(L2StakingStakerAdded)
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
func (it *L2StakingStakerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingStakerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingStakerAdded represents a StakerAdded event raised by the L2Staking contract.
type L2StakingStakerAdded struct {
	Addr   common.Address
	TmKey  [32]byte
	BlsKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakerAdded is a free log retrieval operation binding the contract event 0x058ecb29c230cd5df283c89e996187ed521393fe4546cd1b097921c4b2de293d.
//
// Solidity: event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingFilterer) FilterStakerAdded(opts *bind.FilterOpts, addr []common.Address) (*L2StakingStakerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "StakerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingStakerAddedIterator{contract: _L2Staking.contract, event: "StakerAdded", logs: logs, sub: sub}, nil
}

// WatchStakerAdded is a free log subscription operation binding the contract event 0x058ecb29c230cd5df283c89e996187ed521393fe4546cd1b097921c4b2de293d.
//
// Solidity: event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingFilterer) WatchStakerAdded(opts *bind.WatchOpts, sink chan<- *L2StakingStakerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "StakerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingStakerAdded)
				if err := _L2Staking.contract.UnpackLog(event, "StakerAdded", log); err != nil {
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

// ParseStakerAdded is a log parse operation binding the contract event 0x058ecb29c230cd5df283c89e996187ed521393fe4546cd1b097921c4b2de293d.
//
// Solidity: event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey)
func (_L2Staking *L2StakingFilterer) ParseStakerAdded(log types.Log) (*L2StakingStakerAdded, error) {
	event := new(L2StakingStakerAdded)
	if err := _L2Staking.contract.UnpackLog(event, "StakerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingStakerRemovedIterator is returned from FilterStakerRemoved and is used to iterate over the raw logs and unpacked data for StakerRemoved events raised by the L2Staking contract.
type L2StakingStakerRemovedIterator struct {
	Event *L2StakingStakerRemoved // Event containing the contract specifics and raw log

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
func (it *L2StakingStakerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingStakerRemoved)
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
		it.Event = new(L2StakingStakerRemoved)
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
func (it *L2StakingStakerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingStakerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingStakerRemoved represents a StakerRemoved event raised by the L2Staking contract.
type L2StakingStakerRemoved struct {
	StakerAddresses []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakerRemoved is a free log retrieval operation binding the contract event 0x3511bf213f9290ba907e91e12a43e8471251e1879580ae5509292a3514c23f61.
//
// Solidity: event StakerRemoved(address[] stakerAddresses)
func (_L2Staking *L2StakingFilterer) FilterStakerRemoved(opts *bind.FilterOpts) (*L2StakingStakerRemovedIterator, error) {

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "StakerRemoved")
	if err != nil {
		return nil, err
	}
	return &L2StakingStakerRemovedIterator{contract: _L2Staking.contract, event: "StakerRemoved", logs: logs, sub: sub}, nil
}

// WatchStakerRemoved is a free log subscription operation binding the contract event 0x3511bf213f9290ba907e91e12a43e8471251e1879580ae5509292a3514c23f61.
//
// Solidity: event StakerRemoved(address[] stakerAddresses)
func (_L2Staking *L2StakingFilterer) WatchStakerRemoved(opts *bind.WatchOpts, sink chan<- *L2StakingStakerRemoved) (event.Subscription, error) {

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "StakerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingStakerRemoved)
				if err := _L2Staking.contract.UnpackLog(event, "StakerRemoved", log); err != nil {
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

// ParseStakerRemoved is a log parse operation binding the contract event 0x3511bf213f9290ba907e91e12a43e8471251e1879580ae5509292a3514c23f61.
//
// Solidity: event StakerRemoved(address[] stakerAddresses)
func (_L2Staking *L2StakingFilterer) ParseStakerRemoved(log types.Log) (*L2StakingStakerRemoved, error) {
	event := new(L2StakingStakerRemoved)
	if err := _L2Staking.contract.UnpackLog(event, "StakerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the L2Staking contract.
type L2StakingUndelegatedIterator struct {
	Event *L2StakingUndelegated // Event containing the contract specifics and raw log

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
func (it *L2StakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingUndelegated)
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
		it.Event = new(L2StakingUndelegated)
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
func (it *L2StakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingUndelegated represents a Undelegated event raised by the L2Staking contract.
type L2StakingUndelegated struct {
	Delegatee       common.Address
	Delegator       common.Address
	Amount          *big.Int
	DelegateeAmount *big.Int
	UnlockEpoch     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x92039db29d8c0a1aa1433fe109c69488c8c5e51b23c9de7d303ad80c1fef778c.
//
// Solidity: event Undelegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount, uint256 unlockEpoch)
func (_L2Staking *L2StakingFilterer) FilterUndelegated(opts *bind.FilterOpts, delegatee []common.Address, delegator []common.Address) (*L2StakingUndelegatedIterator, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "Undelegated", delegateeRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingUndelegatedIterator{contract: _L2Staking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x92039db29d8c0a1aa1433fe109c69488c8c5e51b23c9de7d303ad80c1fef778c.
//
// Solidity: event Undelegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount, uint256 unlockEpoch)
func (_L2Staking *L2StakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *L2StakingUndelegated, delegatee []common.Address, delegator []common.Address) (event.Subscription, error) {

	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "Undelegated", delegateeRule, delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingUndelegated)
				if err := _L2Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x92039db29d8c0a1aa1433fe109c69488c8c5e51b23c9de7d303ad80c1fef778c.
//
// Solidity: event Undelegated(address indexed delegatee, address indexed delegator, uint256 amount, uint256 delegateeAmount, uint256 unlockEpoch)
func (_L2Staking *L2StakingFilterer) ParseUndelegated(log types.Log) (*L2StakingUndelegated, error) {
	event := new(L2StakingUndelegated)
	if err := _L2Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2StakingUndelegationClaimedIterator is returned from FilterUndelegationClaimed and is used to iterate over the raw logs and unpacked data for UndelegationClaimed events raised by the L2Staking contract.
type L2StakingUndelegationClaimedIterator struct {
	Event *L2StakingUndelegationClaimed // Event containing the contract specifics and raw log

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
func (it *L2StakingUndelegationClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2StakingUndelegationClaimed)
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
		it.Event = new(L2StakingUndelegationClaimed)
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
func (it *L2StakingUndelegationClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2StakingUndelegationClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2StakingUndelegationClaimed represents a UndelegationClaimed event raised by the L2Staking contract.
type L2StakingUndelegationClaimed struct {
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegationClaimed is a free log retrieval operation binding the contract event 0xcc3089abc79631b3c0c81414a72e237c08559073a970cf474e36ae965e382fb3.
//
// Solidity: event UndelegationClaimed(address indexed delegator, uint256 amount)
func (_L2Staking *L2StakingFilterer) FilterUndelegationClaimed(opts *bind.FilterOpts, delegator []common.Address) (*L2StakingUndelegationClaimedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.FilterLogs(opts, "UndelegationClaimed", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &L2StakingUndelegationClaimedIterator{contract: _L2Staking.contract, event: "UndelegationClaimed", logs: logs, sub: sub}, nil
}

// WatchUndelegationClaimed is a free log subscription operation binding the contract event 0xcc3089abc79631b3c0c81414a72e237c08559073a970cf474e36ae965e382fb3.
//
// Solidity: event UndelegationClaimed(address indexed delegator, uint256 amount)
func (_L2Staking *L2StakingFilterer) WatchUndelegationClaimed(opts *bind.WatchOpts, sink chan<- *L2StakingUndelegationClaimed, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _L2Staking.contract.WatchLogs(opts, "UndelegationClaimed", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2StakingUndelegationClaimed)
				if err := _L2Staking.contract.UnpackLog(event, "UndelegationClaimed", log); err != nil {
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

// ParseUndelegationClaimed is a log parse operation binding the contract event 0xcc3089abc79631b3c0c81414a72e237c08559073a970cf474e36ae965e382fb3.
//
// Solidity: event UndelegationClaimed(address indexed delegator, uint256 amount)
func (_L2Staking *L2StakingFilterer) ParseUndelegationClaimed(log types.Log) (*L2StakingUndelegationClaimed, error) {
	event := new(L2StakingUndelegationClaimed)
	if err := _L2Staking.contract.UnpackLog(event, "UndelegationClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
