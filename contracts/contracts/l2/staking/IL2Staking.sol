// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /***********
     * Structs *
     ***********/

    /// @notice Commission representing a delegatee's commission info.
    ///
    /// @custom:field checkpoint        The epoch when the commission percentage was last changed
    /// @custom:field percentage        commission percentage
    /// @custom:field prePercentage     pre commission percentage
    /// @custom:field amount            unclaimed commission amount
    struct Commission {
        uint256 checkpoint;
        uint256 prePercentage;
        uint256 percentage;
        uint256 amount;
    }

    /// @notice DelegateeDelegation representing a delegatee's delegation info.
    ///
    /// @custom:field checkpoint    The epoch when the share was last changed
    /// @custom:field preAmount     Total delegations of a delegatee
    /// @custom:field amount        Total delegations of a delegatee
    /// @custom:field preShare      Total share of a delegatee at the start of an epoch
    /// @custom:field share         Total share of a delegatee at the end of an epoch
    struct DelegateeDelegation {
        uint256 checkpoint;
        uint256 preAmount;
        uint256 amount;
        uint256 preShare;
        uint256 share;
    }

    /// @notice DelegatorDelegation representing a delegator's delegation info.
    ///
    /// @custom:field checkpoint    The epoch when the share was last changed
    /// @custom:field preShare      share of a delegator at the start of an epoch
    /// @custom:field share         share of a delegator at the end of an epoch
    struct DelegatorDelegation {
        uint256 checkpoint;
        uint256 preShare;
        uint256 share;
    }

    /// @notice Undelegation representing a undelegation info.
    ///
    /// @custom:field delegatee  delegatee
    /// @custom:field amount     staking amount
    /// @custom:field unlock     unlock epoch index
    struct Undelegation {
        address delegatee;
        uint256 amount;
        uint256 unlockEpoch;
    }

    /***********
     * Errors *
     ***********/

    /// @notice error not staker
    error ErrNotStaker();

    /// @notice error invalid nonce
    error ErrInvalidNonce();

    /**********
     * Events *
     **********/

    /// @notice Emitted delegated stake
    /// @param delegatee          delegatee
    /// @param delegator          delegator
    /// @param stakeAmount        stake amount
    /// @param delegateeAmount    new delegatee total amount
    /// @param delegateeShare     new delegatee total sare
    /// @param delegatorShare     new delegator share
    /// @param effectiveEpoch     effective epoch
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 stakeAmount,
        uint256 delegateeAmount,
        uint256 delegateeShare,
        uint256 delegatorShare,
        uint256 effectiveEpoch
    );

    /// @notice Emitted undelegated stake
    /// @param delegatee          delegatee
    /// @param delegator          delegator
    /// @param unstakeAmount        stake amount
    /// @param delegateeAmount    new delegatee total amount
    /// @param delegateeShare     new delegatee total sare
    /// @param delegatorShare     new delegator share
    /// @param effectiveEpoch     effective epoch
    /// @param unlockEpoch        unlock epoch index
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 unstakeAmount,
        uint256 delegateeAmount,
        uint256 delegateeShare,
        uint256 delegatorShare,
        uint256 effectiveEpoch,
        uint256 unlockEpoch
    );

    /// @notice Emitted claim info
    /// @param delegator   delegator
    /// @param unlockEpoch unlock epoch index
    /// @param amount      staking amount
    event UndelegationClaimed(
        address indexed delegatee,
        address indexed delegator,
        uint256 unlockEpoch,
        uint256 amount
    );

    /// @notice Emitted staker added
    /// @param addr     staker address
    /// @param tmKey    staker tendermint pubkey
    /// @param blsKey   staker BLS pubkey
    event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey);

    /// @notice Emitted stakers removed
    /// @param stakerAddresses  stakers removed
    event StakerRemoved(address[] stakerAddresses);

    /// @notice Emitted reward start time updated
    /// @param oldTime    The old reward start time
    /// @param newTime    The new reward start time
    event RewardStartTimeUpdated(uint256 oldTime, uint256 newTime);

    /// @notice Emitted sequencer set max size updated
    /// @param oldSize    The old sequencer set max size
    /// @param newSize    The new sequencer set max size
    event SequencerSetMaxSizeUpdated(uint256 oldSize, uint256 newSize);

    /// @notice Emitted reward epochs uploaded
    /// @param sequencer            The sequencer address
    /// @param delegatorReward      The delegator reward amount
    /// @param commissionAmount     The commission amount
    event Distributed(address indexed sequencer, uint256 delegatorReward, uint256 commissionAmount);

    /// @notice Emitted commission updated
    /// @param staker           staker address
    /// @param oldPercentage    old commission percentage
    /// @param newPercentage    new commission percentage
    /// @param epochEffective   epoch effective
    event CommissionUpdated(
        address indexed staker,
        uint256 oldPercentage,
        uint256 newPercentage,
        uint256 epochEffective
    );

    /// @notice commission claimed
    /// @param delegatee    delegatee
    /// @param amount       amount
    event CommissionClaimed(address indexed delegatee, uint256 amount);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice reward epoch
    function rewardStartTime() external view returns (uint256);

    /// @notice max size of sequencer set
    function sequencerSetMaxSize() external view returns (uint256);

    /// @notice undelegate lock epochs
    function undelegateLockEpochs() external view returns (uint256);

    /// @notice start reward
    function startReward() external;

    /// @notice return current reward epoch index. Revert if not start reward
    function currentEpoch() external view returns (uint256);

    /// @notice check if the user has staked to staker
    /// @param staker   staker address
    function isStakingTo(address staker) external view returns (bool);

    /// @notice Get the delegators length which staked to staker
    /// @param staker staker address
    function getDelegatorsLength(address staker) external view returns (uint256);

    /// @notice Get the delegators which staked to staker in pagination
    /// @param staker       staker address
    /// @param pageSize     page size
    /// @param pageIndex    page index
    function getAllDelegatorsInPagination(
        address staker,
        uint256 pageSize,
        uint256 pageIndex
    ) external view returns (uint256 delegatorsTotalNumber, address[] memory delegatorsInPage);

    /// @notice get stakers info
    /// @param _stakerAddresses    staker's addresses
    function getStakesInfo(address[] calldata _stakerAddresses) external view returns (Types.StakerInfo[] memory);

    /// @notice get stakers
    function getStakers() external view returns (Types.StakerInfo[] memory);

    /// @notice get staker addresses length
    function getStakerAddressesLength() external view returns (uint256);

    /// @notice query all unclaimed commission of a staker
    /// @param delegatee     delegatee address
    function queryUnclaimedCommission(address delegatee) external view returns (uint256 amount);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice add staker, sync from L1
    /// @param nonce    msg nonce
    /// @param add      staker to add. {addr, tmKey, blsKey}
    function addStaker(uint256 nonce, Types.StakerInfo calldata add) external;

    /// @notice remove stakers, sync from L1
    /// @param nonce    msg nonce
    /// @param remove   staker to remove
    function removeStakers(uint256 nonce, address[] calldata remove) external;

    /// @notice setCommissionPercentage set delegate commission percentage
    /// @param percentage    commission percentage, denominator is 100
    function setCommissionPercentage(uint256 percentage) external;

    /// @notice delegator stake morph to delegatee
    /// @param delegatee    stake to whom
    /// @param amount       stake amount
    function delegateStake(address delegatee, uint256 amount) external;

    /// @notice delegator redelegate stake morph token to new delegatee
    /// @param delegateeFrom    old delegatee
    /// @param delegateeTo      new delegatee
    /// @param amount           amount
    function redelegateStake(address delegateeFrom, address delegateeTo, uint256 amount) external;

    /// @notice delegator undelegate stake morph token
    /// @param delegatee    delegatee address
    /// @param amount       undelegate stake amount, undelegate all if set 0
    function undelegateStake(address delegatee, uint256 amount) external;

    /// @notice delegator cliam delegate staking value
    function claimUndelegation() external;

    /// @notice claimCommission claim unclaimed commission reward of a staker
    function claimCommission() external;

    /// @dev distribute inflation by system at end of the epoch
    /// @param epochIndex         epoch index
    /// @param sequencers         sequencers
    /// @param rewards            total rewards
    function distributeInflation(
        uint256 epochIndex,
        address[] calldata sequencers,
        uint256[] calldata rewards
    ) external;
}
