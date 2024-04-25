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

    /// @notice delegated stake
    /// @param delegatee          delegatee
    /// @param delegator          unlock epoch index
    /// @param amount             new delegation amount, not increment
    /// @param effectiveEpoch     effective epoch
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch
    );

    /// @notice undelegated stake
    /// @param delegatee          delegatee
    /// @param delegator          unlock epoch index
    /// @param amount             delegation amount
    /// @param effectiveEpoch     effective epoch
    /// @param ublockEpoch        effective epoch
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch,
        uint256 ublockEpoch
    );

    /// @notice claim info
    event UndelegationClaimed(address indexed delegator, uint256 amount);

    /// @notice commission updated
    event CommissionUpdated(
        address indexed staker,
        uint256 percentage,
        uint256 epochEffective
    );

    /// @notice staker added
    event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey);

    /// @notice Staker removed
    event StakerRemoved(address[] stakerAddresses);

    /// @notice params updated
    event ParamsUpdated(uint256 sequencersSize);

    /// @notice reward start time updated
    event RewardStartTimeUpdated(uint256 rewardStartTime);

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

    /// @notice return current reward epoch index
    function currentEpoch() external view returns (uint256);

    /// @notice check if the user has staked to staker
    /// @param staker sequencers size
    function isStakingTo(address staker) external view returns (bool);

    /// @notice Get all the delegators which staked to staker
    /// @param staker sequencers size
    function getDelegators(
        address staker
    ) external view returns (address[] memory);

    /// @notice get stakers info
    function getStakesInfo(
        address[] calldata _stakerAddresses
    ) external view returns (Types.StakerInfo[] memory);

    /// @notice get stakers
    function getStakers() external view returns (Types.StakerInfo[] memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice add staker, sync from L1
    /// @param add       staker to add. {addr, tmKey, blsKey}
    function addStaker(Types.StakerInfo memory add) external;

    /// @notice remove stakers, sync from L1
    /// @param remove    staker to remove
    function removeStakers(address[] memory remove) external;

    /// @notice setCommissionRate set delegate commission percentage
    /// @param commission    commission percentage
    function setCommissionRate(uint256 commission) external;

    /// @notice delegator stake morph to staker
    /// @param staker    stake to whom
    /// @param amount    stake amount
    function delegateStake(address staker, uint256 amount) external;

    /// @notice delegator unstake morph
    /// @param delegatee delegatee address
    function undelegateStake(address delegatee) external;

    /// @notice delegator cliam delegate staking value
    function claimUndelegation() external;

    /// @notice delegator claim reward
    /// @param delegatee         delegatee address, claim all if empty
    /// @param targetEpochIndex  up to the epoch index that the delegator wants to claim
    function claimReward(address delegatee, uint256 targetEpochIndex) external;

    /// @notice claimCommission claim commission reward
    /// @param targetEpochIndex   up to the epoch index that the staker wants to claim
    function claimCommission(uint256 targetEpochIndex) external;

    /// @notice update params
    /// @param sequencersMaxSize   max size of sequencer set
    function updateParams(uint256 sequencersMaxSize) external;

    /// @notice advance layer2 stage
    /// @param rewardStartTime   reward start time
    function updateRewardStartTime(uint256 rewardStartTime) external;
}
