// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    /**
     * @notice Distribution representing a distribution
     *
     * @custom:field delegatorRewardAmount  total reward amount minus commission
     * @custom:field commissionAmount       delegate commission amount
     * @custom:field delegationAmount       total delegation amount
     * @custom:field remainsNumber          unclaimed delegator number
     * @custom:field delegators             delegator set
     * @custom:field amounts                delegators delegation amount
     */
    struct Distribution {
        uint256 delegatorRewardAmount;
        uint256 commissionAmount;
        uint256 delegationAmount;
        uint256 remainsNumber;
        EnumerableSet.AddressSet delegators;
        mapping(address => uint256) amounts;
    }

    /**
     * @notice Unclaimed representing a unclaimd info of a delegator
     *
     * @custom:field undelegated        whether is undelegated
     * @custom:field delegatees         all delegatees if this delegator, remove delegatee after all reward claimed
     * @custom:field unclaimedStart     unclaimed start epoch index
     * @custom:field unclaimedEnd       unclaimed end epoch index, set 0 if undelegated is false or all claimed
     */
    struct Unclaimed {
        EnumerableSet.AddressSet delegatees;
        mapping(address => bool) undelegated;
        mapping(address => uint256) unclaimedStart;
        mapping(address => uint256) unclaimedEnd;
    }

    /**
     * @notice reward claimed
     */
    event RewardClaimed(
        address indexed delegator,
        address indexed delegatee,
        uint256 upToEpoch,
        uint256 amount
    );

    /**
     * @notice commission claimed
     */
    event CommissionClaimed(
        address indexed delegatee,
        uint256 upToEpoch,
        uint256 amount
    );

    /**
     * @dev return delegatee unclaimed epoch index
     */
    function unclaimedComission(
        address delegatee
    ) external view returns (uint256);

    /**
     * @dev notify delegation
     * @param delegatee         delegatee address
     * @param delegator         delegator address
     * @param effectiveEpoch    delegation effective epoch
     * @param amount            delegator amount
     * @param totalAmount       delegatee total amount
     * @param newDelegation     First delegate or additional delegate
     */
    function notifyDelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 amount,
        uint256 totalAmount,
        bool newDelegation
    ) external;

    /**
     * @dev notify unDelegation
     * @param delegatee         delegatee address
     * @param delegator         delegator address
     * @param effectiveEpoch    delegation effective epoch
     * @param totalAmount       delegatee total amount
     */
    function notifyUndelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 totalAmount
    ) external;

    /**
     * @dev update epoch reward
     * @param epochIndex        epoch index
     * @param sequencers        sequencers
     * @param delegatorRewards  sequencer's delegatorRewardAmount
     * @param commissions       sequencers commission
     *
     */
    function updateEpochReward(
        uint256 epochIndex,
        address[] memory sequencers,
        uint256[] memory delegatorRewards,
        uint256[] memory commissions
    ) external;

    /**
     * @dev claim delegation reward of all sequencers.
     * @param delegator         delegator address
     * @param targetEpochIndex  the epoch index that the user wants to claim up to
     *
     * If targetEpochIndex is zero, claim up to latest mint epoch,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claimAll(address delegator, uint256 targetEpochIndex) external;

    /**
     * @dev claim delegation reward of a delegatee.
     * @param delegatee         delegatee address
     * @param delegator         delegator address
     * @param targetEpochIndex  the epoch index that the user wants to claim up to
     *
     * If targetEpochIndex is zero, claim up to latest mint epoch,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claim(
        address delegatee,
        address delegator,
        uint256 targetEpochIndex
    ) external;

    /**
     * @dev claim commission reward
     * @param delegatee         delegatee address
     * @param targetEpochIndex  the epoch index that the user wants to claim up to
     */
    function claimCommission(
        address delegatee,
        uint256 targetEpochIndex
    ) external;
}
