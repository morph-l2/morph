// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice Undelegation representing a undelegation info.
     *
     * @custom:field delegatee  delegatee
     * @custom:field amount     staking amount
     * @custom:field unlock     unlock epoch index
     */
    struct Undelegation {
        address delegatee;
        uint256 amount;
        uint256 unlockEpoch;
    }

    /**
     * @notice delegated stake
     *
     * @custom:field delegatee          delegatee
     * @custom:field delegator          unlock epoch index
     * @custom:field amount             new delegation amount, not increment
     * @custom:field effectiveEpoch     effective epoch
     */
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch
    );

    /**
     * @notice undelegated stake
     *
     * @custom:field delegatee          delegatee
     * @custom:field delegator          unlock epoch index
     * @custom:field amount             delegation amount
     * @custom:field effectiveEpoch     effective epoch
     * @custom:field ublockEpoch        effective epoch
     */
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 effectiveEpoch,
        uint256 ublockEpoch
    );

    /**
     * @notice claim info
     */
    event UndelegationClaimed(address indexed delegator, uint256 amount);

    /**
     * @notice commission updated
     */
    event CommissionUpdated(
        address indexed staker,
        uint256 percentage,
        uint256 epochEffective
    );

    /**
     * @notice staker added
     */
    event StakerAdded(address indexed addr, bytes32 tmKey, bytes blsKey);

    /**
     * @notice Staker removed
     */
    event StakerRemoved(address[] stakerAddresses);

    /**
     * @notice params updated
     */
    event ParamsUpdated(uint256 sequencersSize);

    /**
     * @notice reward start time updated
     */
    event RewardStartTimeUpdated(uint256 rewardStartTime);

    /**
     * @notice reward epoch
     */
    function REWARD_START_TIME() external view returns (uint256);

    /**
     * @notice reward epoch
     */
    function REWARD_EPOCH() external view returns (uint256);

    /**
     * @notice max size of sequencer set
     */
    function SEQUENCER_MAX_SIZE() external view returns (uint256);

    /**
     * @notice undelegate lock epochs
     */
    function UNDELEGATE_LOCK_EPOCHS() external view returns (uint256);

    /**
     * @notice initializer
     * @param _admin                params admin
     * @param _sequencersMaxSize    max size of sequencer set
     * @param _undelegateLockEpochs undelegate lock epochs
     * @param _rewardStartTime      reward start time
     * @param _stakers              initial stakers, must be same as initial sequencer set in sequencer contract
     **/
    function initialize(
        address _admin,
        uint256 _sequencersMaxSize,
        uint256 _undelegateLockEpochs,
        uint256 _rewardStartTime,
        Types.StakerInfo[] calldata _stakers
    ) external;

    /**
     * @notice add staker, sync from L1
     * @param add       staker to add. {addr, tmKey, blsKey}
     */
    function addStaker(Types.StakerInfo memory add) external;

    /**
     * @notice remove stakers, sync from L1
     * @param remove    staker to remove
     */
    function removeStakers(address[] memory remove) external;

    /**
     * @notice setCommissionRate set delegate commission percentage
     * @param commission    commission percentage
     */
    function setCommissionRate(uint256 commission) external;

    /**
     * @notice delegator stake morph to staker
     * @param staker    stake to whom
     * @param amount    stake amount
     */
    function delegateStake(address staker, uint256 amount) external;

    /**
     * @notice delegator unstake morph
     * @param delegatee delegatee address
     */
    function undelegateStake(address delegatee) external;

    /**
     * @notice delegator cliam delegate staking value
     */
    function claimUndelegation() external;

    /**
     * @notice delegator claim reward
     * @param delegatee         delegatee address, claim all if empty
     * @param targetEpochIndex  up to the epoch index that the delegator wants to claim
     */
    function claimReward(address delegatee, uint256 targetEpochIndex) external;

    /**
     * @notice claimCommission claim commission reward
     * @param targetEpochIndex   up to the epoch index that the staker wants to claim
     */
    function claimCommission(uint256 targetEpochIndex) external;

    /**
     * @notice update params
     * @param _sequencersMaxSize   max size of sequencer set
     */
    function updateParams(uint256 _sequencersMaxSize) external;

    /**
     * @notice advance layer2 stage
     * @param _rewardStartTime   reward start time
     */
    function updateRewardStartTime(uint256 _rewardStartTime) external;

    /**
     * @notice start reward
     */
    function startReward() external;

    /**
     * @notice return current reward epoch index
     */
    function currentEpoch() external view returns (uint256);

    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function isStakingTo(address staker) external view returns (bool);

    /**
     * @notice Get all the delegators which staked to staker
     * @param staker sequencers size
     */
    function getDelegators(
        address staker
    ) external view returns (address[] memory);

    /**
     * @notice get stakers info
     */
    function getStakesInfo(
        address[] calldata _stakerAddresses
    ) external view returns (Types.StakerInfo[] memory);
}
