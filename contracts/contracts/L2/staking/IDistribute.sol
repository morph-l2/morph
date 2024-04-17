// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";

/**
 * @dev Interface of the Distribute.
 */
interface IDistribute {
    struct Set {
        EnumerableSet.AddressSet index;
        mapping(address => uint256) value;
    }

    struct TimeOrderedSet {
        DoubleEndedQueue.Bytes32Deque index;
        mapping(uint256 => uint256) value;
    }

    struct Distribution {
        uint256 totalAmount;
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

    // event of claimAll
    event ClaimAll(address indexed from, address indexed to, uint256 amount);

    // event of claim
    event Claim(address indexed from, address indexed to, uint256 amount);

    event NotifyDelegate(
        address indexed sequencer,
        uint256 indexed epochIndex,
        address indexed account,
        uint256 amount,
        uint256 blockNumber);

    event NotifyUnDelegate(
        address indexed sequencer,
        address indexed account,
        uint256 deadlineClaimEpochIndex
    );

    /**
     * @dev Initialization parameter, which can only be called once.
     * @param morphToken_ address
     * @param record_ address
     * @param stake_ address
     */
    function initialize(
        address morphToken_,
        address record_,
        address stake_
    ) external;

    function notify(uint256 blockTime, uint256 blockNumber) external;

    function notifyUnDelegate(
        address sequencer,
        address account,
        uint256 deadlineClaimEpochIndex
    ) external;

    function notifyDelegate(
        address sequencer,
        uint256 epochIndex,
        address account,
        uint256 amount,
        uint256 blockNumber
    ) external;

    function mint() external;

    /**
     * @dev latestMintedEpochIndex the maximum value of the epoch index after mint.
     */
    function latestMintedEpochIndex() external returns (uint256);

    /**
     * @dev claimedEpochIndex query the latest claimed epoch index.
     * @param sequencer, the address of the sequencer to query.
     * @param account, the address of the delegator to query.
     */
    function claimedEpochIndex(address sequencer, address account) external returns (uint256);

    /**
     * @dev claimAll claim all user delegate to all sequencer rewards.
     * @param account, claim account address.
     * @param targetEpochIndex, the epoch index that the user wants to claim,
     * If zero, the latest mint epoch index is claimed,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claimAll(address account, uint256 targetEpochIndex) external;

    /**
     * @dev claim user delegate to sequencer rewards.
     * @param sequencer, sequencer address.
     * @param account, claim account address.
     * @param targetEpochIndex, the epoch index that the user wants to claim,
     * If zero, the latest mint epoch index is claimed,
     * otherwise it must be greater than the last claimed epoch index.
     */
    function claim(address sequencer, address account, uint256 targetEpochIndex) external;
}
