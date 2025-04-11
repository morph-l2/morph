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
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_otherStaking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidPageSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidSequencerSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrInvalidStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoCandidate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoClaimableUndelegateRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoStakers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNoUndelegateRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrNotStaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrOnlyMorphTokenContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrOnlySystem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRequestExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRewardNotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrRewardStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrStartTimeNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroLockEpochs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroSequencerSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrZeroShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CommissionClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"CommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commissionAmount\",\"type\":\"uint256\"}],\"name\":\"Distributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegateeFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegateeTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmountFrom\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmountTo\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTime\",\"type\":\"uint256\"}],\"name\":\"RewardStartTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSize\",\"type\":\"uint256\"}],\"name\":\"SequencerSetMaxSizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"name\":\"StakerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"stakerAddresses\",\"type\":\"address[]\"}],\"name\":\"StakerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockEpoch\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegationClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPH_TOKEN_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_STAKING\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEQUENCER_CONTRACT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo\",\"name\":\"add\",\"type\":\"tuple\"}],\"name\":\"addStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidateNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"claimUndelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"claimableUndelegateRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"commissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"delegateeDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegatorDelegations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo\",\"name\":\"add\",\"type\":\"tuple\"}],\"name\":\"emergencyAddStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"emergencyRemoveStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageIndex\",\"type\":\"uint256\"}],\"name\":\"getAllDelegatorsInPagination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorsTotalNumber\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegatorsInPage\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getDelegatorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakerAddressesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakers\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakerAddresses\",\"type\":\"address[]\"}],\"name\":\"getStakesInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sequencersMaxSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_undelegateLockEpochs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardStartTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.StakerInfo[]\",\"name\":\"_stakers\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStakingTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestSequencerSetSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"lockedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"pendingUndelegateRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"queryDelegationAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"queryUnclaimedCommission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"recordBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegateeFrom\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegateeTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"remove\",\"type\":\"address[]\"}],\"name\":\"removeStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStartTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardStarted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerSetMaxSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakerRankings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ranking\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"tmKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegateLockEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"undelegateRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIL2Staking.UndelegateRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"undelegateSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardStartTime\",\"type\":\"uint256\"}],\"name\":\"updateRewardStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sequencerSetMaxSize\",\"type\":\"uint256\"}],\"name\":\"updateSequencerSetMaxSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x61012060405234801562000011575f80fd5b5060405162006057380380620060578339810160408190526200003491620000a7565b7353000000000000000000000000000000000000076080526001600160a01b031660a05273530000000000000000000000000000000000001360c05273530000000000000000000000000000000000001760e05273530000000000000000000000000000000000002161010052620000d6565b5f60208284031215620000b8575f80fd5b81516001600160a01b0381168114620000cf575f80fd5b9392505050565b60805160a05160c05160e05161010051615ed3620001845f395f8181610465015261462401525f818161062b015261486901525f8181610796015281816136390152818161499401528181614a4901528181614b2501528181614de801528181614e950152614f7101525f81816105d0015281816130ff0152613a7201525f81816104d401528181610687015281816130d50152818161312901528181613a480152613a9c0152615ed35ff3fe608060405234801561000f575f80fd5b5060043610610339575f3560e01c8063746c8ae1116101b3578063a61bb764116100f3578063d31d83d91161009e578063f2fde38b11610079578063f2fde38b146107c1578063fad99f98146107d4578063fc6facc6146107dc578063ff4840cd146107ef575f80fd5b8063d31d83d914610770578063d557714114610791578063f0261bc2146107b8575f80fd5b8063b7a587bf116100ce578063b7a587bf14610704578063bf2dca0a14610732578063cce6cf9f1461075d575f80fd5b8063a61bb764146106c9578063affed0e0146106dc578063b5d2e0dc146106e5575f80fd5b80638da5cb5b1161015e57806391c05b0b1161013957806391c05b0b1461066f578063927ede2d1461068257806396ab994d146106a95780639d51c3b9146106b6575f80fd5b80638da5cb5b146106155780638e21d5fb146106265780639168ae721461064d575f80fd5b80637c7e8bd21161018e5780637c7e8bd2146105b8578063831cfb58146105cb57806384d7d1d4146105f2575f80fd5b8063746c8ae114610582578063766718081461058a5780637b05afb514610592575f80fd5b80633434735f1161027e578063439162b5116102295780634d99dd16116102045780634d99dd16146105415780636bd8f804146105545780637046529b14610567578063715018a61461057a575f80fd5b8063439162b514610513578063459598a21461052657806346cdc18a14610539575f80fd5b80633cb747bf116102595780633cb747bf146104d257806340b5c837146104f857806343352d611461050b575f80fd5b80633434735f146104605780633b2713c51461049f5780633b802421146104c9575f80fd5b806313f22527116102e9578063201018fb116102c4578063201018fb1461041b5780632cc138be1461042e5780632e787be31461043757806330158eea14610440575f80fd5b806313f22527146103ba57806319fac8fd146103cd5780631d5611b8146103e0575f80fd5b80630321731c116103195780630321731c1461038b5780630eb573af1461039e57806312a3e947146103b1575f80fd5b806243b7581461033d5780629c6f0c14610363578063026e402b14610378575b5f80fd5b61035061034b36600461557f565b610802565b6040519081526020015b60405180910390f35b61037661037136600461559a565b610828565b005b6103766103863660046155e4565b6109eb565b61035061039936600461557f565b610f8f565b6103766103ac36600461560e565b610fc9565b610350609a5481565b6103506103c836600461557f565b61109e565b6103766103db36600461560e565b611160565b6104066103ee36600461557f565b60a36020525f90815260409020805460019091015482565b6040805192835260208301919091520161035a565b61035061042936600461560e565b611259565b61035060985481565b61035060995481565b61045361044e36600461566d565b611432565b60405161035a919061570d565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161035a565b6103506104ad3660046157b3565b60a460209081525f928352604080842090915290825290205481565b610350609c5481565b7f0000000000000000000000000000000000000000000000000000000000000000610487565b61037661050636600461560e565b611651565b610453611738565b6103766105213660046157df565b611954565b61048761053436600461560e565b611e15565b609e54610350565b61037661054f3660046155e4565b611e3d565b61037661056236600461584d565b6125af565b61037661057536600461559a565b6130ca565b61037661323e565b610376613251565b6103506134df565b6104066105a036600461557f565b60a16020525f90815260409020805460019091015482565b6103506105c636600461557f565b61353d565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b61060561060036600461557f565b61355a565b604051901515815260200161035a565b6033546001600160a01b0316610487565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b61066061065b36600461557f565b613584565b60405161035a9392919061588b565b61037661067d36600461560e565b613636565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b6097546106059060ff1681565b6103506106c43660046157b3565b613892565b6103506106d73660046155e4565b6138a4565b610350609d5481565b6103506106f336600461557f565b609f6020525f908152604090205481565b6107176107123660046155e4565b6139dc565b6040805182518152602092830151928101929092520161035a565b61035061074036600461557f565b6001600160a01b03165f90815260a1602052604090206001015490565b61037661076b3660046158bb565b613a3d565b61078361077e366004615903565b613f9f565b60405161035a929190615978565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b610350609b5481565b6103766107cf36600461557f565b614117565b6103766141c1565b6103766107ea3660046158bb565b614273565b6103766107fd36600461557f565b614621565b6001600160a01b0381165f90815260a260205260408120610822906146d6565b92915050565b6108306146df565b81609d54811461086c576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108778360016159c5565b609d55609f5f61088a602085018561557f565b6001600160a01b03166001600160a01b031681526020019081526020015f20545f0361092c57609e6108bf602084018461557f565b81546001810183555f9283526020808420909101805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039390931692909217909155609e5491609f916109129086018661557f565b6001600160a01b0316815260208101919091526040015f20555b8160a05f61093d602084018461557f565b6001600160a01b0316815260208101919091526040015f2061095f8282615af5565b5061096f9050602083018361557f565b6001600160a01b03167f058ecb29c230cd5df283c89e996187ed521393fe4546cd1b097921c4b2de293d60208401356109ab60408601866159d8565b6040516109ba93929190615c5c565b60405180910390a260975460ff161580156109d95750609954609e5411155b156109e6576109e6614753565b505050565b6001600160a01b0382165f908152609f6020526040812054839103610a3c576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a446148d3565b815f03610a7d576040517f608294ac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a9e336001600160a01b0385165f90815260a26020526040902090614946565b5060975460ff16610b30576001600160a01b0383165f90815260a46020908152604080832033845290915281208054849290610adb9084906159c5565b90915550506001600160a01b0383165f90815260a3602052604081208054849290610b079084906159c5565b90915550506001600160a01b0383165f90815260a3602052604090208054600190910155610c4c565b6001600160a01b0383165f90815260a3602090815260408083206001810154905460a48452828520338652909352908320549092829003610ba7576001600160a01b0386165f81815260a460209081526040808320338452825280832089905592825260a390522060018101869055859055610c48565b81610bb28487615caf565b610bbc9190615cf3565b610bc690826159c5565b6001600160a01b0387165f81815260a46020908152604080832033845282528083209490945591815260a39091529081208054879290610c079084906159c5565b90915550829050610c188487615caf565b610c229190615cf3565b610c2c90846159c5565b6001600160a01b0387165f90815260a360205260409020600101555b5050505b6001600160a01b0383165f90815260a36020526040902054829003610c83576001609c5f828254610c7d91906159c5565b90915550505b6001600160a01b0383165f908152609f602052604090205460975460ff168015610cad5750600181115b15610ee2575f610cbe600183615d06565b90505b8015610ee05760a35f609e610cd7600185615d06565b81548110610ce757610ce7615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a39290919085908110610d2757610d27615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115610ece575f609e610d5f600184615d06565b81548110610d6f57610d6f615d19565b5f91825260209091200154609e80546001600160a01b0390921692509083908110610d9c57610d9c615d19565b5f918252602090912001546001600160a01b0316609e610dbd600185615d06565b81548110610dcd57610dcd615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e8381548110610e0c57610e0c615d19565b5f9182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0393909316929092179091558290609f90609e610e53600185615d06565b81548110610e6357610e63615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055610e928260016159c5565b609f5f609e8581548110610ea857610ea8615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b80610ed881615d46565b915050610cc1565b505b6001600160a01b0384165f81815260a360209081526040918290205482518781529182015281513393927f24d7bda8602b916d64417f0dbfe2e2e88ec9b1157bd9f596dfdb91ba26624e04928290030190a3610f3f33308561495a565b60975460ff168015610f525750609b5481115b8015610f7757506099546001600160a01b0385165f908152609f602052604090205411155b15610f8457610f84614753565b506109e66001606555565b6001600160a01b0381165f90815260a66020526040812054600f81810b700100000000000000000000000000000000909204900b03610822565b610fd16146df565b801580610fdf575060995481145b15611016576040517f383a648e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609980549082905560408051828152602081018490527f98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed910160405180910390a16097545f9060ff1661106b57609e5461106f565b609c545b90505f609954821061108357609954611085565b815b9050609b54811461109857611098614753565b50505050565b6001600160a01b0381165f90815260a66020526040812054600f81810b700100000000000000000000000000000000909204900b0381805b82811015611158576001600160a01b0385165f90815260a6602052604081206110ff9083614be6565b5f81815260a5602090815260409182902082518084019093528054835260010154908201819052919250906111326134df565b106111475761114084615d5b565b935061114e565b5050611158565b50506001016110d6565b509392505050565b335f818152609f602052604081205490036111a7576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60148211156111e2576040517f6e11528c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f81815260a1602081815260408084208054825180840184528981526001830180548287019081529789905295855251909155935190925581518681529081018390529192917f6e500db30ce535d38852e318f333e9be41a3fec6c65d234ebb06203c896db9a5910160405180910390a2505050565b5f6112626148d3565b6112aa60a65f335b6001600160a01b03166001600160a01b031681526020019081526020015f2054600f81810b700100000000000000000000000000000000909204900b0390565b5f036112e2576040517f5f013ef800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8115806112f957506112f660a65f3361126a565b82115b611303578161130f565b61130f60a65f3361126a565b91505f5b82156113a857335f90815260a66020526040812061133090614c7b565b5f81815260a5602090815260409182902082518084019093528054835260010154908201819052919250906113636134df565b10156113705750506113a8565b335f90815260a66020526040902061138790614cf3565b50805161139490846159c5565b925061139f85615d46565b94505050611313565b805f036113e1576040517f3cc5dedc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113ec335b82614dae565b60405181815233907fcc3089abc79631b3c0c81414a72e237c08559073a970cf474e36ae965e382fb39060200160405180910390a2905061142d6001606555565b919050565b60605f8267ffffffffffffffff81111561144e5761144e615a39565b60405190808252806020026020018201604052801561149a57816020015b60408051606080820183525f80835260208301529181019190915281526020019060019003908161146c5790505b5090505f5b8381101561115857604051806060016040528060a05f8888868181106114c7576114c7615d19565b90506020020160208101906114dc919061557f565b6001600160a01b03908116825260208083019390935260409091015f90812054909116835291019060a09088888681811061151957611519615d19565b905060200201602081019061152e919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f2060010154815260200160a05f88888681811061156b5761156b615d19565b9050602002016020810190611580919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f2060020180546115ad90615a66565b80601f01602080910402602001604051908101604052809291908181526020018280546115d990615a66565b80156116245780601f106115fb57610100808354040283529160200191611624565b820191905f5260205f20905b81548152906001019060200180831161160757829003601f168201915b505050505081525082828151811061163e5761163e615d19565b602090810291909101015260010161149f565b6116596146df565b60975460ff1615611696576040517fbd51da0d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b42811115806116b057506116ad6201518082615d73565b15155b806116bc575060985481145b156116f3576040517fde16b26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609880549082905560408051828152602081018490527f91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555910160405180910390a15050565b609e546060905f9067ffffffffffffffff81111561175857611758615a39565b6040519080825280602002602001820160405280156117a457816020015b60408051606080820183525f8083526020830152918101919091528152602001906001900390816117765790505b5090505f5b609e5481101561194e57604051806060016040528060a05f609e85815481106117d4576117d4615d19565b5f9182526020808320909101546001600160a01b0390811684528382019490945260409092018120549092168352609e8054939091019260a09291908690811061182057611820615d19565b905f5260205f20015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020019081526020015f2060010154815260200160a05f609e858154811061187957611879615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902060020180546118aa90615a66565b80601f01602080910402602001604051908101604052809291908181526020018280546118d690615a66565b80156119215780601f106118f857610100808354040283529160200191611921565b820191905f5260205f20905b81548152906001019060200180831161190457829003601f168201915b505050505081525082828151811061193b5761193b615d19565b60209081029190910101526001016117a9565b50919050565b5f54610100900460ff161580801561197257505f54600160ff909116105b8061198b5750303b15801561198b57505f5460ff166001145b611a1c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015611a78575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6001600160a01b038716611ab8576040517fee77070400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855f03611af1576040517f2da55d0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845f03611b2a576040517f7d8ad8a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4284111580611b445750611b416201518085615d73565b15155b15611b7b576040517fde16b26100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f829003611bb5576040517fbb01aad100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bbe8761502b565b611bc6615089565b6099869055609a8590556098849055609b8290555f5b609b54811015611d3757838382818110611bf857611bf8615d19565b9050602002810190611c0a9190615d86565b60a05f868685818110611c1f57611c1f615d19565b9050602002810190611c319190615d86565b611c3f90602081019061557f565b6001600160a01b0316815260208101919091526040015f20611c618282615af5565b905050609e848483818110611c7857611c78615d19565b9050602002810190611c8a9190615d86565b611c9890602081019061557f565b8154600180820184555f938452602090932001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055611ce19082906159c5565b609f5f868685818110611cf657611cf6615d19565b9050602002810190611d089190615d86565b611d1690602081019061557f565b6001600160a01b0316815260208101919091526040015f2055600101611bdc565b50604080515f8152602081018890527f98b982a120d9be7d9c68d85a1aed8158d1d52e517175bfb3eb4280692f19b1ed910160405180910390a1604080515f8152602081018690527f91c38708087fb4ba51bd0e6a106cc1fbaf340479a2e81d18f2341e8c78f97555910160405180910390a18015611e0c575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b609e8181548110611e24575f80fd5b5f918252602090912001546001600160a01b0316905081565b611e456148d3565b805f03611e7e576040517f608294ac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611e888233615127565b5f03611ec0576040517f857ad50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611ecb8333615127565b1015611f03576040517f08c2348a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0382165f908152609f60205260408120546097549015919060ff16611f2f575f611f3d565b609a54611f3d9060016159c5565b60408051808201909152848152602081018290529091505f33611f5f336151ad565b60405160609290921b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001660208301526034820152605401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301205f81815260a590935291205490915015612013576040517fdeeb052700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f81815260a560209081526040808320855181558583015160019182015533845260a68352818420805470010000000000000000000000000000000090819004600f0b8087528284019095529290942085905583546fffffffffffffffffffffffffffffffff90811693909101160217905560975460ff16612116576001600160a01b0386165f90815260a460209081526040808320338452909152812080548792906120c1908490615d06565b90915550506001600160a01b0386165f90815260a36020526040812080548792906120ed908490615d06565b90915550506001600160a01b0386165f90815260a36020526040902080546001909101556121eb565b6001600160a01b0386165f90815260a3602090815260408083206001810154905460a4845282852033865290935292205481612152848a615caf565b61215c9190615cf3565b6121669082615d06565b6001600160a01b038a165f81815260a46020908152604080832033845282528083209490945591815260a390915290812080548a92906121a7908490615d06565b909155508290506121b8848a615caf565b6121c29190615cf3565b6121cc9084615d06565b6001600160a01b038a165f90815260a360205260409020600101555050505b6001600160a01b0386165f908152609f602052604090205484158015612213575060975460ff165b80156122205750609c5481105b15612481576001600160a01b0387165f908152609f602052604081205461224990600190615d06565b90505b6001609c5461225b9190615d06565b81101561247f5760a35f609e838154811061227857612278615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040018120549060a390609e6122ad8560016159c5565b815481106122bd576122bd615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115612477575f609e82815481106122fb576122fb615d19565b5f918252602090912001546001600160a01b03169050609e61231e8360016159c5565b8154811061232e5761232e615d19565b5f91825260209091200154609e80546001600160a01b03909216918490811061235957612359615d19565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905580609e61239c8460016159c5565b815481106123ac576123ac615d19565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b03929092169190911790556123ec8260016159c5565b609f5f609e858154811061240257612402615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040019020556124318260026159c5565b609f5f609e6124418660016159c5565b8154811061245157612451615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b60010161224c565b505b841580156124a457506001600160a01b0387165f90815260a36020526040902054155b156124c1576001609c5f8282546124bb9190615d06565b90915550505b6001600160a01b0387165f90815260a360205260409020543360408051898152602081018490529081018790526001600160a01b03918216918a16907f92039db29d8c0a1aa1433fe109c69488c8c5e51b23c9de7d303ad80c1fef778c9060600160405180910390a38515801561253a575060975460ff165b80156125485750609b548211155b801561258e5750609b546001600160a01b0389165f908152609f6020526040902054118061258e5750609c546001600160a01b0389165f908152609f6020526040902054115b1561259b5761259b614753565b5050505050506125ab6001606555565b5050565b6001600160a01b0383165f908152609f6020526040812054849103612600576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0383165f908152609f6020526040812054849103612651576040517f3efa0ab900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6126596148d3565b825f03612692576040517f608294ac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61269c8533615127565b5f036126d4576040517f857ad50500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826126df8633615127565b1015612717576040517f08c2348a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385165f908152609f602052604081205460975490159060ff166127c3576001600160a01b0387165f90815260a4602090815260408083203384529091528120805487929061276e908490615d06565b90915550506001600160a01b0387165f90815260a360205260408120805487929061279a908490615d06565b90915550506001600160a01b0387165f90815260a3602052604090208054600190910155612898565b6001600160a01b0387165f90815260a3602090815260408083206001810154905460a48452828520338652909352922054816127ff848a615caf565b6128099190615cf3565b6128139082615d06565b6001600160a01b038b165f81815260a46020908152604080832033845282528083209490945591815260a390915290812080548a9290612854908490615d06565b90915550829050612865848a615caf565b61286f9190615cf3565b6128799084615d06565b6001600160a01b038b165f90815260a360205260409020600101555050505b6001600160a01b0387165f908152609f6020526040902054811580156128c0575060975460ff165b80156128cd5750609c5481105b15612b2e576001600160a01b0388165f908152609f60205260408120546128f690600190615d06565b90505b6001609c546129089190615d06565b811015612b2c5760a35f609e838154811061292557612925615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040018120549060a390609e61295a8560016159c5565b8154811061296a5761296a615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115612b24575f609e82815481106129a8576129a8615d19565b5f918252602090912001546001600160a01b03169050609e6129cb8360016159c5565b815481106129db576129db615d19565b5f91825260209091200154609e80546001600160a01b039092169184908110612a0657612a06615d19565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b039290921691909117905580609e612a498460016159c5565b81548110612a5957612a59615d19565b5f918252602090912001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055612a998260016159c5565b609f5f609e8581548110612aaf57612aaf615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055612ade8260026159c5565b609f5f609e612aee8660016159c5565b81548110612afe57612afe615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b6001016128f9565b505b81158015612b5157506001600160a01b0388165f90815260a36020526040902054155b15612b6e576001609c5f828254612b689190615d06565b90915550505b81158015612b7e575060975460ff165b8015612b8c5750609b548111155b8015612bd25750609b546001600160a01b0389165f908152609f60205260409020541180612bd25750609c546001600160a01b0389165f908152609f6020526040902054115b15612bdc57600192505b612bfd336001600160a01b0389165f90815260a26020526040902090614946565b5060975460ff16612c8f576001600160a01b0387165f90815260a46020908152604080832033845290915281208054889290612c3a9084906159c5565b90915550506001600160a01b0387165f90815260a3602052604081208054889290612c669084906159c5565b90915550506001600160a01b0387165f90815260a3602052604090208054600190910155612d64565b6001600160a01b0387165f90815260a3602090815260408083206001810154905460a4845282852033865290935292205481612ccb848b615caf565b612cd59190615cf3565b612cdf90826159c5565b6001600160a01b038b165f81815260a46020908152604080832033845282528083209490945591815260a390915290812080548b9290612d209084906159c5565b90915550829050612d31848b615caf565b612d3b9190615cf3565b612d4590846159c5565b6001600160a01b038b165f90815260a360205260409020600101555050505b6001600160a01b0387165f90815260a36020526040902054869003612d9b576001609c5f828254612d9591906159c5565b90915550505b506001600160a01b0386165f908152609f602052604090205460975460ff168015612dc65750600181115b15612ffb575f612dd7600183615d06565b90505b8015612ff95760a35f609e612df0600185615d06565b81548110612e0057612e00615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a39290919085908110612e4057612e40615d19565b5f9182526020808320909101546001600160a01b031683528201929092526040019020541115612fe7575f609e612e78600184615d06565b81548110612e8857612e88615d19565b5f91825260209091200154609e80546001600160a01b0390921692509083908110612eb557612eb5615d19565b5f918252602090912001546001600160a01b0316609e612ed6600185615d06565b81548110612ee657612ee6615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e8381548110612f2557612f25615d19565b5f9182526020822001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0393909316929092179091558290609f90609e612f6c600185615d06565b81548110612f7c57612f7c615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055612fab8260016159c5565b609f5f609e8581548110612fc157612fc1615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055505b80612ff181615d46565b915050612dda565b505b60975460ff16801561300e5750609b5481115b801561303357506099546001600160a01b0388165f908152609f602052604090205411155b1561303d57600192505b821561304b5761304b614753565b6001600160a01b038881165f81815260a36020908152604080832054948c16808452928190205481518c8152928301869052908201819052923392917ffdac6e81913996d95abcc289e90f2d8bd235487ce6fe6f821e7d21002a1915b49060600160405180910390a450505050506130c36001606555565b5050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156131b257507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015613183573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131a79190615dc2565b6001600160a01b0316145b610830576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f7374616b696e673a206f6e6c79206f74686572207374616b696e6720636f6e7460448201527f7261637420616c6c6f77656400000000000000000000000000000000000000006064820152608401611a13565b6132466146df565b61324f5f61502b565b565b6132596146df565b609854421015613295576040517f080bb11a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609c545f036132d0576040517fd7d776cb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b609780547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660019081179091555b609e5481101561347c575f5b818110156134735760a35f609e838154811061332957613329615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001812054609e8054919260a3929091908690811061336957613369615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001902054111561346b575f609e82815481106133a7576133a7615d19565b5f91825260209091200154609e80546001600160a01b03909216925090849081106133d4576133d4615d19565b5f91825260209091200154609e80546001600160a01b0390921691849081106133ff576133ff615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555080609e848154811061343e5761343e615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550505b60010161330b565b506001016132ff565b505f5b609e548110156134d6576134948160016159c5565b609f5f609e84815481106134aa576134aa615d19565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205560010161347f565b5061324f614753565b5f60985442101561351c576040517fd021716f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b620151806098544261352e9190615d06565b6135389190615cf3565b905090565b6001600160a01b0381165f90815260a76020526040812054610822565b6001600160a01b0381165f90815260a4602090815260408083203384529091528120541515610822565b60a06020525f90815260409020805460018201546002830180546001600160a01b039093169391926135b590615a66565b80601f01602080910402602001604051908101604052809291908181526020018280546135e190615a66565b801561362c5780601f106136035761010080835404028352916020019161362c565b820191905f5260205f20905b81548152906001019060200180831161360f57829003601f168201915b5050505050905083565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614613698576040517f4032cbb200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60aa5415613837575f5b6136ac60a86146d6565b811015613835575f60a1816136c260a8856151d2565b6001600160a01b0316815260208101919091526040015f9081205460aa5490925060ab826136f160a8876151d2565b6001600160a01b0316815260208101919091526040015f20546137149086615caf565b61371e9190615cf3565b90505f606461372d8484615caf565b6137379190615cf3565b90505f6137448284615d06565b90508160a15f61375560a8896151d2565b6001600160a01b03166001600160a01b031681526020019081526020015f206001015f82825461378591906159c5565b9091555081905060a35f61379a60a8896151d2565b6001600160a01b03166001600160a01b031681526020019081526020015f205f015f8282546137c991906159c5565b909155506137da905060a8866151d2565b6001600160a01b03167f60ce3cc2d133631eac66a476f14997a9fa682bd05a60dd993cf02285822d78d8828460405161381d929190918252602082015260400190565b60405180910390a25050600190920191506136a29050565b505b5f5b61384360a86146d6565b8110156125ab5760ab5f61385860a8846151d2565b6001600160a01b0316815260208101919091526040015f9081205561388961388160a8836151d2565b60a8906151dd565b50600101613839565b5f61389d8383615127565b9392505050565b6001600160a01b0382165f90815260a66020526040812054600f81810b700100000000000000000000000000000000909204900b035f036138e657505f610822565b81158061392657506001600160a01b0383165f90815260a66020526040902054600f81810b700100000000000000000000000000000000909204900b0382115b6139305781613966565b6001600160a01b0383165f90815260a66020526040902054600f81810b700100000000000000000000000000000000909204900b035b91505f805b83811015611158576001600160a01b0385165f90815260a6602052604081206139949083614be6565b5f81815260a560209081526040918290208251808401909352805480845260019091015491830191909152919250906139cd90856159c5565b9350505080600101905061396b565b604080518082019091525f80825260208201526001600160a01b0383165f90815260a660205260408120613a109084614be6565b5f90815260a560209081526040918290208251808401909352805483526001015490820152949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015613b2557507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015613af6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613b1a9190615dc2565b6001600160a01b0316145b613bb1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f7374616b696e673a206f6e6c79206f74686572207374616b696e6720636f6e7460448201527f7261637420616c6c6f77656400000000000000000000000000000000000000006064820152608401611a13565b82609d548114613bed576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613bf88460016159c5565b609d555f805b83811015613f5757609b54609f5f878785818110613c1e57613c1e615d19565b9050602002016020810190613c33919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f205411613c5d57600191505b5f609f5f878785818110613c7357613c73615d19565b9050602002016020810190613c88919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f20541115613ed9575f6001609f5f888886818110613cc657613cc6615d19565b9050602002016020810190613cdb919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f2054613d059190615d06565b90505b609e54613d1790600190615d06565b811015613de957609e613d2b8260016159c5565b81548110613d3b57613d3b615d19565b5f91825260209091200154609e80546001600160a01b039092169183908110613d6657613d66615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001609f5f609e8481548110613da957613da9615d19565b5f9182526020808320909101546001600160a01b0316835282019290925260400181208054909190613ddc908490615d06565b9091555050600101613d08565b50609e805480613dfb57613dfb615ddd565b5f8281526020812082015f19908101805473ffffffffffffffffffffffffffffffffffffffff19169055909101909155609f90868684818110613e4057613e40615d19565b9050602002016020810190613e55919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f205f90555f60a35f878785818110613e8d57613e8d615d19565b9050602002016020810190613ea2919061557f565b6001600160a01b0316815260208101919091526040015f20541115613ed9576001609c5f828254613ed39190615d06565b90915550505b60a05f868684818110613eee57613eee615d19565b9050602002016020810190613f03919061557f565b6001600160a01b0316815260208101919091526040015f908120805473ffffffffffffffffffffffffffffffffffffffff191681556001810182905590613f4d6002830182615525565b5050600101613bfe565b507f3511bf213f9290ba907e91e12a43e8471251e1879580ae5509292a3514c23f618484604051613f89929190615e0a565b60405180910390a180156130c3576130c3614753565b5f6060835f03613fdb576040517f89076b3900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b0385165f90815260a260205260409020613ffb906146d6565b91508367ffffffffffffffff81111561401657614016615a39565b60405190808252806020026020018201604052801561403f578160200160208202803683370190505b5090505f61404d8486615caf565b90505f600161405c86826159c5565b6140669088615caf565b6140709190615d06565b905061407d600185615d06565b8111156140925761408f600185615d06565b90505b815f5b82821161410b576140c9826140a981615d5b565b6001600160a01b038c165f90815260a260205260409020909450906151d2565b85826140d481615d5b565b9350815181106140e6576140e6615d19565b60200260200101906001600160a01b031690816001600160a01b031681525050614095565b50505050935093915050565b61411f6146df565b6001600160a01b0381166141b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401611a13565b6141be8161502b565b50565b6141c96148d3565b335f90815260a160205260408120600101549003614213576040517f5426dfcd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b335f81815260a160205260408120600101805491905590614233906113e6565b60405181815233907f8e14daa5332205b1634040e1054e93d1f5396ec8bf0115d133b7fbaf4a52e4119060200160405180910390a25061324f6001606555565b61427b6146df565b82609d5481146142b7576040517f2f0fd70500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6142c28460016159c5565b609d555f805b83811015613f5757609b54609f5f8787858181106142e8576142e8615d19565b90506020020160208101906142fd919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f20541161432757600191505b5f609f5f87878581811061433d5761433d615d19565b9050602002016020810190614352919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f205411156145a3575f6001609f5f88888681811061439057614390615d19565b90506020020160208101906143a5919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f20546143cf9190615d06565b90505b609e546143e190600190615d06565b8110156144b357609e6143f58260016159c5565b8154811061440557614405615d19565b5f91825260209091200154609e80546001600160a01b03909216918390811061443057614430615d19565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055506001609f5f609e848154811061447357614473615d19565b5f9182526020808320909101546001600160a01b03168352820192909252604001812080549091906144a6908490615d06565b90915550506001016143d2565b50609e8054806144c5576144c5615ddd565b5f8281526020812082015f19908101805473ffffffffffffffffffffffffffffffffffffffff19169055909101909155609f9086868481811061450a5761450a615d19565b905060200201602081019061451f919061557f565b6001600160a01b03166001600160a01b031681526020019081526020015f205f90555f60a35f87878581811061455757614557615d19565b905060200201602081019061456c919061557f565b6001600160a01b0316815260208101919091526040015f205411156145a3576001609c5f82825461459d9190615d06565b90915550505b60a05f8686848181106145b8576145b8615d19565b90506020020160208101906145cd919061557f565b6001600160a01b0316815260208101919091526040015f908120805473ffffffffffffffffffffffffffffffffffffffff1916815560018101829055906146176002830182615525565b50506001016142c8565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031614614683576040517f52d033bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61468e60a882614946565b50600160aa5f8282546146a191906159c5565b90915550506001600160a01b0381165f90815260ab602052604081208054600192906146ce9084906159c5565b909155505050565b5f610822825490565b6033546001600160a01b0316331461324f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401611a13565b60995460975460ff161561477757609954609c5410156147725750609c545b614788565b609954609e5410156147885750609e545b5f8167ffffffffffffffff8111156147a2576147a2615a39565b6040519080825280602002602001820160405280156147cb578160200160208202803683370190505b5090505f5b8281101561483857609e81815481106147eb576147eb615d19565b905f5260205f20015f9054906101000a90046001600160a01b031682828151811061481857614818615d19565b6001600160a01b03909216602092830291909101909101526001016147d0565b506040517f9b8201a40000000000000000000000000000000000000000000000000000000081526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639b8201a49061489e908490600401615e57565b5f604051808303815f87803b1580156148b5575f80fd5b505af11580156148c7573d5f803e3d5ffd5b50509151609b55505050565b60026065540361493f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401611a13565b6002606555565b5f61389d836001600160a01b0384166151f1565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa1580156149db573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906149ff9190615e69565b6040517f23b872dd0000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301528581166024830152604482018590529192507f0000000000000000000000000000000000000000000000000000000000000000909116906323b872dd906064016020604051808303815f875af1158015614a91573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614ab59190615e80565b614aeb576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614b6c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614b909190615e69565b9050821580614ba8575082614ba58383615d06565b14155b156130c3576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001606555565b5f80614c08614bf48461523d565b8554614c039190600f0b615e9f565b6152f2565b84549091507001000000000000000000000000000000009004600f90810b9082900b12614c61576040517fb4120f1400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600f0b5f9081526001939093016020525050604090205490565b5f614ca28254600f81810b700100000000000000000000000000000000909204900b131590565b15614cd9576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b5f9081526001909101602052604090205490565b5f614d1a8254600f81810b700100000000000000000000000000000000909204900b131590565b15614d51576040517f3db2a12a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b508054600f0b5f818152600180840160205260408220805492905583547fffffffffffffffffffffffffffffffff000000000000000000000000000000001692016fffffffffffffffffffffffffffffffff169190911790915590565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614e2f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614e539190615e69565b6040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526001600160a01b038581166004830152602482018590529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303815f875af1158015614edd573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614f019190615e80565b614f37576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301525f917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa158015614fb8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614fdc9190615e69565b9050821580614ff4575082614ff18383615d06565b14155b15611098576040517f9a7058e100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603380546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f54610100900460ff1661511f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611a13565b61324f615386565b6001600160a01b038083165f90815260a460209081526040808320938516835292905290812054156151a5576001600160a01b038084165f81815260a360208181526040808420600181015460a48452828620978a168652968352908420549490935252546151969190615caf565b6151a09190615cf3565b61389d565b505f92915050565b6001600160a01b0381165f90815260a76020526040902080546001810182559061194e565b5f61389d838361541c565b5f61389d836001600160a01b038416615442565b5f81815260018301602052604081205461523657508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610822565b505f610822565b5f7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8211156152ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e206160448201527f6e20696e743235360000000000000000000000000000000000000000000000006064820152608401611a13565b5090565b80600f81900b811461142d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203160448201527f32382062697473000000000000000000000000000000000000000000000000006064820152608401611a13565b5f54610100900460ff16614bdf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401611a13565b5f825f01828154811061543157615431615d19565b905f5260205f200154905092915050565b5f818152600183016020526040812054801561551c575f615464600183615d06565b85549091505f9061547790600190615d06565b90508181146154d6575f865f01828154811061549557615495615d19565b905f5260205f200154905080875f0184815481106154b5576154b5615d19565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806154e7576154e7615ddd565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610822565b5f915050610822565b50805461553190615a66565b5f825580601f10615540575050565b601f0160209004905f5260205f20908101906141be91905b808211156152ee575f8155600101615558565b6001600160a01b03811681146141be575f80fd5b5f6020828403121561558f575f80fd5b813561389d8161556b565b5f80604083850312156155ab575f80fd5b82359150602083013567ffffffffffffffff8111156155c8575f80fd5b8301606081860312156155d9575f80fd5b809150509250929050565b5f80604083850312156155f5575f80fd5b82356156008161556b565b946020939093013593505050565b5f6020828403121561561e575f80fd5b5035919050565b5f8083601f840112615635575f80fd5b50813567ffffffffffffffff81111561564c575f80fd5b6020830191508360208260051b8501011115615666575f80fd5b9250929050565b5f806020838503121561567e575f80fd5b823567ffffffffffffffff811115615694575f80fd5b6156a085828601615625565b90969095509350505050565b5f81518084525f5b818110156156d0576020818501810151868301820152016156b4565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b5f60208083018184528085518083526040925060408601915060408160051b8701018488015f5b838110156157a5578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180516001600160a01b0316845287810151888501528601516060878501819052615791818601836156ac565b968901969450505090860190600101615734565b509098975050505050505050565b5f80604083850312156157c4575f80fd5b82356157cf8161556b565b915060208301356155d98161556b565b5f805f805f8060a087890312156157f4575f80fd5b86356157ff8161556b565b9550602087013594506040870135935060608701359250608087013567ffffffffffffffff81111561582f575f80fd5b61583b89828a01615625565b979a9699509497509295939492505050565b5f805f6060848603121561585f575f80fd5b833561586a8161556b565b9250602084013561587a8161556b565b929592945050506040919091013590565b6001600160a01b0384168152826020820152606060408201525f6158b260608301846156ac565b95945050505050565b5f805f604084860312156158cd575f80fd5b83359250602084013567ffffffffffffffff8111156158ea575f80fd5b6158f686828701615625565b9497909650939450505050565b5f805f60608486031215615915575f80fd5b83356159208161556b565b95602085013595506040909401359392505050565b5f815180845260208085019450602084015f5b8381101561596d5781516001600160a01b031687529582019590820190600101615948565b509495945050505050565b828152604060208201525f6159906040830184615935565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8082018082111561082257610822615998565b5f8083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112615a0b575f80fd5b83018035915067ffffffffffffffff821115615a25575f80fd5b602001915036819003821315615666575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b600181811c90821680615a7a57607f821691505b60208210810361194e577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b601f8211156109e657805f5260205f20601f840160051c81016020851015615ad65750805b601f840160051c820191505b818110156130c3575f8155600101615ae2565b8135615b008161556b565b6001600160a01b03811673ffffffffffffffffffffffffffffffffffffffff1983541617825550600160208084013560018401556002830160408501357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1863603018112615b6c575f80fd5b8501803567ffffffffffffffff811115615b84575f80fd5b8036038483011315615b94575f80fd5b615ba881615ba28554615a66565b85615ab1565b5f601f821160018114615bdb575f8315615bc457508382018601355b5f19600385901b1c1916600184901b178555615c51565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015615c2757868501890135825593880193908901908801615c08565b5084821015615c45575f1960f88660031b161c198885880101351681555b505060018360011b0185555b505050505050505050565b83815260406020820152816040820152818360608301375f818301606090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010192915050565b808202811582820484141761082257610822615998565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82615d0157615d01615cc6565b500490565b8181038181111561082257610822615998565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81615d5457615d54615998565b505f190190565b5f5f198203615d6c57615d6c615998565b5060010190565b5f82615d8157615d81615cc6565b500690565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa1833603018112615db8575f80fd5b9190910192915050565b5f60208284031215615dd2575f80fd5b815161389d8161556b565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b60208082528181018390525f908460408401835b86811015615e4c578235615e318161556b565b6001600160a01b031682529183019190830190600101615e1e565b509695505050505050565b602081525f61389d6020830184615935565b5f60208284031215615e79575f80fd5b5051919050565b5f60208284031215615e90575f80fd5b8151801515811461389d575f80fd5b8082018281125f831280158216821582161715615ebe57615ebe615998565b50509291505056fea164736f6c6343000818000a",
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
