// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IL2Staking {
    /**
     * @notice Unstaking representing a unstaking info.
     *
     * @custom:field amount  staking amuont
     * @custom:field unlock  unlock block height
     */
    struct Unstaking {
        uint256 amount;
        uint256 unlock;
    }

    /**
     * @notice stakers information.
     * @custom:field staker
     * @return {addr, tmKey, blsKey, active}
     */
    function stakers(
        address staker
    ) external view returns (address, bytes32, bytes memory);

    /**
     * @notice staker's status
     * @param staker    staker
     */
    function stakerStatus(address staker) external view returns (bool);

    /**
     * @notice information staking by delegator to staker
     * @param staker    delegatee
     * @param delegator delegator
     */
    function stakings(
        address staker,
        address delegator
    ) external view returns (uint256);

    /**
     * @notice delegator's unstaking info
     * @param staker    delegatee
     * @param delegator delegator
     */
    function unstakings(
        address staker,
        address delegator
    ) external view returns (uint256, uint256);

    /**
     * @notice staker's all morph amount
     * @param staker stake to whom
     */
    function stakersAmount(address staker) external view returns (uint256);

    /**
     * @notice number of currently active
     */
    function sequencersSize() external view returns (uint256);

    // /**
    //  * @notice get all stakers
    //  */
    // function getStakers() external view returns (address[] memory);

    // /**
    //  * @notice get stakers info
    //  */
    // function getStakesInfo(
    //     address[] memory stakers
    // ) external view returns (Types.StakerInfo[] memory);

    // /**
    //  * @notice Get all the delegators which staked to staker
    //  * @param staker sequencers size
    //  */
    // function getDelegators(
    //     address staker
    // ) external view returns (address[] memory);

    // /**
    //  * @notice check if the delegator has staked to staker
    //  * @param staker sequencers size
    //  */
    // function isStakingTo(address staker) external view returns (bool);

    // /**
    //  * @notice delegator withdrawal
    //  * @param staker stake to whom
    //  */
    // function withdrawal(address staker) external;

    // /**
    //  * @notice delegator claim reward
    //  * @param staker stake to whom
    //  */
    // function claim(address staker) external;

    // /**
    //  * @notice delegator claim all reward
    //  */
    // function claimAll() external;

    // /**
    //  * @notice delegator unstake morph to staker
    //  * @param staker stake to whom
    //  */
    // function unDelegateStake(address staker) external;

    // /**
    //  * @notice delegator stake morph to staker
    //  * @param staker stake to whom
    //  * @param amount stake amount
    //  */
    // function delegateStake(address staker, uint256 amount) external;

    // /**
    //  * @notice update params
    //  * @param _sequencersSize sequencers size
    //  * @param _epoch epoch number
    //  */
    // function updateParams(uint256 _sequencersSize, uint256 _epoch) external;

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
