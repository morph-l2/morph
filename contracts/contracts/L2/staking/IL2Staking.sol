// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice information staking by delegator to staker
     * @param staker stake to
     * @param delegator who
     */
    function stakingInfo(address delegator, address staker) external view returns (uint256);

    /**
     * @notice staker's morph amount
     * @param staker stake to
     */
    function stakersAmount(address staker) external view returns (uint256);

    /**
     * @notice delegator's unstaking info
     * @param staker stake to
     * @param delegator who
     */
    function unstakingInfo(
        address delegator,
        address staker
    ) external view returns (uint256, uint256, bool);

    /**
     * @notice smallest staking value
     */
    function limit() external view returns (uint256);

    /**
     * @notice lock block number after withdrawal
     */
    function lock() external view returns (uint256);

    /**
     * @notice number of currently active
     */
    function sequencersSize() external view returns (uint256);

    /**
     * @notice get all stakers
     */
    function getStakers() external view returns (address[] memory);

    /**
     * @notice Get all the delegators which staked to staker
     * @param staker sequencers size
     */
    function getDelegators(address staker) external view returns (address[] memory);

    /**
     * @notice check if the user has staked to staker
     * @param staker sequencers size
     */
    function isStakingTo(address staker) external view returns (bool);

    /**
     * @notice user withdrawal
     * @param staker stake to
     */
    function withdrawal(address staker) external;

    /**
     * @notice user unstake morph to staker
     * @param staker stake to
     */
    function unstake(address staker) external;

    /**
     * @notice user stake morph to staker
     * @param staker stake to
     * @param amount stake amount
     */
    function stake(address staker, uint256 amount) external;

    /**
     * @notice update params
     * @param _limit smallest staking value
     * @param _lock withdraw lock time
     * @param _sequencersSize sequencers size
     */
    function updateParams(uint256 _sequencersSize, uint256 _limit, uint256 _lock) external;

    /**
     * @notice update stakers
     */
    function updateStakers(
        Types.StakerInfo[] memory add,
        Types.StakerInfo[] memory remove
    ) external;

    /**
     * @notice test stakers
     */
    function testStakers(address[] memory _stakers) external;
}
