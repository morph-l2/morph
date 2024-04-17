// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    /*********************** Structs **************************/

    struct Set {
        EnumerableSet.AddressSet index;
        mapping(address => uint256) value;
    }

    struct TimeOrderedSet {
        DoubleEndedQueue.Bytes32Deque index;
        mapping(uint256 => uint256) value;
    }

    struct Distribution {
        // total amount of all account
        uint256 totalAmount;
        // record the number of users who do not claim
        uint256 remainNumber;
        // mapping(delegator => amount)
        Set amounts;
        bool valid;
    }

    struct DelegatorEpochRecord {
        // begin delegate epoch index
        uint256 begin;
        // undelegate epoch index
        uint256 deadline;
        // claimed epoch index
        uint256 claimed;
    }

    /*********************** Events **************************/

    // event of claimAll
    event ClaimAll(address indexed from, address indexed to, uint256 amount);

    // event of claim
    event Claim(address indexed from, address indexed to, uint256 amount);

    // event of notifyDelegate
    event NotifyDelegate(
        address indexed sequencer,
        uint256 indexed epochIndex,
        address indexed account,
        uint256 amount,
        uint256 blockNumber
    );

    // event of notifyUnDelegate
    event NotifyUnDelegate(
        address indexed sequencer,
        address indexed account,
        uint256 deadlineClaimEpochIndex
    );

    /*********************** Functions **************************/

    /**
     * @dev notify the block push time corresponds to the block number message.
     * @param blockTime time of the first block of the day
     * @param blockNumber number of the first block of the day
     */
    function notify(uint256 blockTime, uint256 blockNumber) external;

    /**
     * @dev notifyUnDelegate pushes the epoch index of
     * the user's undelegate in the specified sequencer.
     *
     * @param sequencer the specified sequencer address.
     * @param account user address.
     * @param deadlineClaimEpochIndex the epoch index of the undelegate.
     */
    function notifyUnDelegate(
        address sequencer,
        address account,
        uint256 deadlineClaimEpochIndex
    ) external;

    /**
     * @dev notifyDelegate pushes account delegate message.
     * @param sequencer the specified sequencer address.
     * @param epochIndex the epoch index corresponding to the delegate.
     * @param account user address.
     * @param amount corresponding to the delegate.
     * @param blockNumber delegate Specifies the block number corresponding to the delegate.
     */
    function notifyDelegate(
        address sequencer,
        uint256 epochIndex,
        address account,
        uint256 amount,
        uint256 blockNumber
    ) external;

    /**
     * @dev mint function
     * record the amount of rewards for each epoch index.
     *
     * @notice only record contract can call.
     */
    function mint() external;

    /**
     * @dev claimedEpochIndex query the latest claimed epoch index.
     * @param sequencer, the address of the sequencer to query.
     * @param account, the address of the delegator to query.
     */
    function claimedEpochIndex(
        address sequencer,
        address account
    ) external returns (uint256);

    /**
     * @dev claimAll claim all user delegate to all sequencer rewards.
     * @param delegator         delegator address
     * @param targetEpochIndex  the epoch index that the user wants to claim
     * If zero, the latest mint epoch index is claimed,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claimAll(address delegator, uint256 targetEpochIndex) external;

    /**
     * @dev claim user delegate to sequencer rewards.
     * @param sequencer         sequencer address
     * @param delegator         delegator address
     * @param targetEpochIndex  the epoch index that the user wants to claim
     * If zero, the latest mint epoch index is claimed,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claim(
        address sequencer,
        address delegator,
        uint256 targetEpochIndex
    ) external;
}
