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
    event StakerRemoved(address[] stakerAddrs);

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
     * @notice return current reward epoch index
     */
    function currentEpoch() external view returns (uint256);
}
