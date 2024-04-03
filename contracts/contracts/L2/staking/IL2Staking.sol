// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice information staking by delegator to staker
     * @param delegator who
     * @param staker stake to
     */
    function stakingInfo(address staker, address delegator) external view returns (uint256);

    /**
     * @notice staker's morph amount
     * @param staker stake to
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
     * @param staker stake to
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
     * @notice user claim
     * @param staker stake to
     */
    function claim(address staker) external;

    /**
     * @notice user claim all dstribution
     */
    function claimAll() external;

    /**
     * @notice user unstake morph to staker
     * @param staker stake to
     */
    function unDelegateStake(address staker) external;

    /**
     * @notice user stake morph to staker
     * @param staker stake to
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
    function updateStakers(Types.StakerInfo[] memory stakers, bool active) external;
}
