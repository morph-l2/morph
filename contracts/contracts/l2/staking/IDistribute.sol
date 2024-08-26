// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    /***********
     * Structs *
     ***********/

    /// @notice Distribution representing a distribution
    ///
    /// @custom:field delegatorRewardAmount  total reward amount minus commission
    /// @custom:field commissionAmount       delegate commission amount
    /// @custom:field delegationAmount       total delegation amount
    /// @custom:field remainsNumber          unclaimed delegator number
    /// @custom:field delegators             delegator set
    /// @custom:field amounts                delegators delegation amount
    struct Distribution {
        uint256 delegatorRewardAmount;
        uint256 commissionAmount;
        uint256 delegationAmount;
        uint256 remainsNumber;
        EnumerableSetUpgradeable.AddressSet delegators;
        mapping(address delegator => uint256 amount) amounts;
    }

    /// @notice Unclaimed representing a unclaimd info of a delegator
    ///
    /// @custom:field delegatees         all delegatees if this delegator, remove delegatee after all reward claimed
    /// @custom:field undelegated        whether is undelegated
    /// @custom:field unclaimedStart     unclaimed start epoch index
    /// @custom:field unclaimedEnd       unclaimed end epoch index, set 0 if undelegated is false or all claimed
    struct Unclaimed {
        EnumerableSetUpgradeable.AddressSet delegatees;
        mapping(address delegator => bool undelegated) undelegated;
        mapping(address delegator => uint256 startEpoch) unclaimedStart;
        mapping(address delegator => uint256 endEpoch) unclaimedEnd;
    }

    /**********
     * Events *
     **********/

    /// @notice reward claimed
    /// @param delegator    delegator
    /// @param delegatee    delegatee
    /// @param upToEpoch    up to epoch index
    /// @param amount       amount
    event RewardClaimed(address indexed delegator, address indexed delegatee, uint256 upToEpoch, uint256 amount);

    /// @notice commission claimed
    /// @param delegatee    delegatee
    /// @param upToEpoch    up to epoch index
    /// @param amount       amount
    event CommissionClaimed(address indexed delegatee, uint256 upToEpoch, uint256 amount);

    /*************************
     * Public View Functions *
     *************************/

    /// @dev return the next epoch to claim commission for a delegatee
    /// @param delegatee    delegatee
    function nextEpochToClaimCommission(address delegatee) external view returns (uint256);

    /// @notice query unclaimed morph reward on a delegatee
    /// @param delegatee     delegatee address
    /// @param delegator     delegatee address
    function queryUnclaimed(address delegatee, address delegator) external view returns (uint256);

    /// @notice query all unclaimed morph reward
    /// @param delegator     delegatee address
    function queryAllUnclaimed(
        address delegator
    ) external view returns (address[] memory delegatee, uint256[] memory reward);

    /// @notice query all unclaimed morph reward epochs info
    /// @param delegator     delegatee address
    function queryAllUnclaimedEpochs(
        address delegator
    ) external view returns (address[] memory, bool[] memory, uint256[] memory, uint256[] memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @dev notify delegation
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param effectiveEpoch    delegation effective epoch
    /// @param amount            delegator amount
    /// @param totalAmount       delegatee total amount
    /// @param remainsNumber     delegator number
    /// @param newDelegation     first delegate or additional delegate
    function notifyDelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 amount,
        uint256 totalAmount,
        uint256 remainsNumber,
        bool newDelegation
    ) external;

    /// @dev notify unDelegation
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param effectiveEpoch    delegation effective epoch
    /// @param totalAmount       delegatee total amount
    /// @param remainsNumber     delegator number
    function notifyUndelegation(
        address delegatee,
        address delegator,
        uint256 effectiveEpoch,
        uint256 totalAmount,
        uint256 remainsNumber
    ) external;

    /// @dev update epoch reward
    /// @param epochIndex        epoch index
    /// @param sequencers        sequencers
    /// @param delegatorRewards  sequencer's delegatorRewardAmount
    /// @param commissions       sequencers commission
    function updateEpochReward(
        uint256 epochIndex,
        address[] calldata sequencers,
        uint256[] calldata delegatorRewards,
        uint256[] calldata commissions
    ) external;

    /// @dev claim delegation reward of all sequencers.
    /// @param delegator         delegator address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    ///
    /// If targetEpochIndex is zero, claim up to latest mint epoch,
    /// otherwise it must be greater than the last claimed epoch index.
    function claimAll(address delegator, uint256 targetEpochIndex) external;

    /// @dev claim delegation reward of a delegatee.
    /// @param delegatee         delegatee address
    /// @param delegator         delegator address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    ///
    /// If targetEpochIndex is zero, claim up to latest mint epoch,
    /// otherwise it must be greater than the last claimed epoch index.
    function claim(address delegatee, address delegator, uint256 targetEpochIndex) external;

    /// @dev claim commission reward
    /// @param delegatee         delegatee address
    /// @param targetEpochIndex  the epoch index that the user wants to claim up to
    function claimCommission(address delegatee, uint256 targetEpochIndex) external;
}
