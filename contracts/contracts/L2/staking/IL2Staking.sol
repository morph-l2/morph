// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice Undelegation representing a undelegation info.
     *
     * @custom:field delegatee  delegatee
     * @custom:field amount     staking amuont
     * @custom:field unlock     unlock epoch index
     */
    struct Undelegation {
        address delegatee;
        uint256 amount;
        uint256 unlockEpoch;
    }

    /**
     * @notice stake info
     */
    event Delegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount
    );

    /**
     * @notice unstake info
     */
    event Undelegated(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount,
        uint256 ublockEpoch
    );

    /**
     * @notice claim info
     */
    event Claimed(address indexed delegator, uint256 amount);

    /**
     * @notice withdrawal info
     */
    event withdrawn(
        address indexed delegatee,
        address indexed delegator,
        uint256 amount
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
}
