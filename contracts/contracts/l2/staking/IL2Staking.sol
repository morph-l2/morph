// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /***********
     * Structs *
     ***********/

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

    /**********
     * Events *
     **********/

    /// @notice Emitted delegated stake
    /// @param delegatee          delegatee
    /// @param delegator          delegator
    /// @param amount             new delegation amount, not increment
    /// @param stakeAmount        stake amount
    /// @param effectiveEpoch     effective epoch
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 stakeAmount,
        uint256 effectiveEpoch
    );

    /// @notice Emitted undelegated stake
    /// @param delegatee          delegatee
    /// @param delegator          delegator
    /// @param amount             undelegation amount
    /// @param effectiveEpoch     effective epoch
    /// @param unlockEpoch        unlock epoch index
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
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

    /// @notice Emitted commission updated
    /// @param staker           staker address
    /// @param percentage       commission percentage
    /// @param epochEffective   epoch effective
    event CommissionUpdated(address indexed staker, uint256 percentage, uint256 epochEffective);

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

    /// @notice Get all the delegators which staked to staker
    /// @param staker staker address
    function getAllDelegators(address staker) external view returns (address[] memory);

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

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice add staker, sync from L1
    /// @param add       staker to add. {addr, tmKey, blsKey}
    function addStaker(Types.StakerInfo calldata add) external;

    /// @notice remove stakers, sync from L1
    /// @param remove    staker to remove
    function removeStakers(address[] calldata remove) external;

    /// @notice setCommissionRate set delegate commission percentage
    /// @param commission    commission percentage, denominator is 100
    function setCommissionRate(uint256 commission) external;

    /// @notice delegator stake morph to delegatee
    /// @param delegatee    stake to whom
    /// @param amount       stake amount
    function delegateStake(address delegatee, uint256 amount) external;

    /// @notice delegator unstake morph
    /// @param delegatee delegatee address
    function undelegateStake(address delegatee) external;

    /// @notice delegator cliam delegate staking value
    function claimUndelegation() external;

    /// @notice delegator claim reward
    /// @param delegatee         delegatee address, claim all if address(0)
    /// @param targetEpochIndex  up to the epoch index that the delegator wants to claim
    function claimReward(address delegatee, uint256 targetEpochIndex) external;

    /// @notice claimCommission claim commission reward
    /// @param targetEpochIndex   up to the epoch index that the staker wants to claim
    function claimCommission(uint256 targetEpochIndex) external;
}
