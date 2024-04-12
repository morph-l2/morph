// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice information staking by delegator to staker
     * @param delegator who
     * @param staker stake to whom
     */
    function stakingInfo(
        address staker,
        address delegator
    ) external view returns (uint256);

    /**
     * @notice staker's all morph amount
     * @param staker stake to whom
     */
    function stakersAmount(address staker) external view returns (uint256);

    /**
     * @notice staker information.
     * @custom:field staker
     * @return {addr, tmKey, blsKey, active}
     */
    function stakerInfo(
        address staker
    ) external view returns (address, bytes32, bytes memory, bool);

    /**
     * @notice delegator's unstaking info
     * @param staker stake to whom
     * @param delegator who
     */
    function unstakingInfo(
        address staker,
        address delegator
    ) external view returns (uint256, uint256);

    /**
     * @notice epoch number
     */
    function epoch() external view returns (uint256);

    /**
     * @notice number of currently active
     */
    function sequencersSize() external view returns (uint256);

    /**
     * @notice get all stakers
     */
    function getStakers() external view returns (address[] memory);

    /**
     * @notice get stakers info
     */
    function getStakesInfo(
        address[] memory stakers
    ) external view returns (Types.StakerInfo[] memory);

    /**
     * @notice Get all the delegators which staked to staker
     * @param staker sequencers size
     */
    function getDelegators(
        address staker
    ) external view returns (address[] memory);

    /**
     * @notice check if the delegator has staked to staker
     * @param staker sequencers size
     */
    function isStakingTo(address staker) external view returns (bool);

    /**
     * @notice delegator withdrawal
     * @param staker stake to whom
     */
    function withdrawal(address staker) external;

    /**
     * @notice delegator claim reward
     * @param staker stake to whom
     */
    function claim(address staker) external;

    /**
     * @notice delegator claim all reward
     */
    function claimAll() external;

    /**
     * @notice delegator unstake morph to staker
     * @param staker stake to whom
     */
    function unDelegateStake(address staker) external;

    /**
     * @notice delegator stake morph to staker
     * @param staker stake to whom
     * @param amount stake amount
     */
    function delegateStake(address staker, uint256 amount) external;

    /**
     * @notice update params
     * @param _sequencersSize sequencers size
     * @param _epoch epoch number
     */
    function updateParams(uint256 _sequencersSize, uint256 _epoch) external;

    /**
     * @notice update stakers
     * @param stakers, sync from l1, {addr, tmKey, blsKey}
     * @param active, active & inActive
     */
    function updateStakers(
        Types.StakerInfo[] memory stakers,
        bool active
    ) external;
}
