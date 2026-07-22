// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/// @title L1Sequencer
/// @notice L1 contract for managing sequencer address with history tracking.
///         Supports querying which sequencer was active at any given L2 block height.
contract L1Sequencer is OwnableUpgradeable {
    // ============ Types ============

    struct HistoryRecord {
        uint64 startL2Block;
        address sequencerAddr;
    }

    // ============ Storage ============

    /// @notice Ordered array of sequencer records (by startL2Block ascending).
    ///         sequencerHistory[0] is the first sequencer, active from L2 block 0.
    ///         This contract only answers "who is the sequencer at height N" (identity);
    ///         the PBFT → single-sequencer upgrade boundary is decided off-chain by block
    ///         timestamp, not stored here.
    HistoryRecord[] public sequencerHistory;

    // ============ Events ============

    event SequencerUpdated(
        address indexed oldSequencer,
        address indexed newSequencer,
        uint64 startL2Block
    );

    // ============ Initializer ============

    /// @notice Initialize the contract
    /// @param _owner Contract owner (multisig recommended)
    function initialize(address _owner) external initializer {
        require(_owner != address(0), "invalid owner");
        __Ownable_init();
        _transferOwnership(_owner);
    }

    // ============ Admin Functions ============

    /// @notice Set the first sequencer (called once). The sequencer is active from L2 block 0;
    ///         the upgrade height is no longer specified here — the PBFT → single-sequencer
    ///         switch is triggered off-chain by block timestamp.
    /// @param firstSequencer  The first sequencer address.
    function setFirstSequencer(
        address firstSequencer
    ) external onlyOwner {
        require(sequencerHistory.length == 0, "already initialized");
        require(firstSequencer != address(0), "invalid address");

        sequencerHistory.push(HistoryRecord({
            startL2Block: 0,
            sequencerAddr: firstSequencer
        }));

        emit SequencerUpdated(address(0), firstSequencer, 0);
    }

    /// @notice Register a sequencer change at a future L2 block height.
    ///         The new sequencer is NOT active until startL2Block is reached.
    /// @param newSequencer   New sequencer address.
    /// @param startL2Block   L2 block height when the new sequencer takes over.
    ///                       Must be strictly greater than the last record.
    function updateSequencer(
        address newSequencer,
        uint64 startL2Block
    ) external onlyOwner {
        require(newSequencer != address(0), "invalid address");
        require(sequencerHistory.length > 0, "not initialized");
        require(
            startL2Block > sequencerHistory[sequencerHistory.length - 1].startL2Block,
            "startL2Block must be greater than last record"
        );

        address oldSequencer = sequencerHistory[sequencerHistory.length - 1].sequencerAddr;

        sequencerHistory.push(HistoryRecord({
            startL2Block: startL2Block,
            sequencerAddr: newSequencer
        }));

        emit SequencerUpdated(oldSequencer, newSequencer, startL2Block);
    }

    // ============ View Functions ============

    /// @notice Get the sequencer that was active at a given L2 block height.
    /// @dev    Binary search: O(log n).
    function getSequencerAt(uint64 l2Height) external view returns (address) {
        uint256 len = sequencerHistory.length;
        require(len > 0, "no sequencer configured");

        uint256 low = 0;
        uint256 high = len - 1;
        uint256 result = 0;

        while (low <= high) {
            uint256 mid = (low + high) / 2;
            if (sequencerHistory[mid].startL2Block <= l2Height) {
                result = mid;
                if (mid == high) break;
                low = mid + 1;
            } else {
                if (mid == 0) break;
                high = mid - 1;
            }
        }

        require(sequencerHistory[result].startL2Block <= l2Height, "no sequencer at height");
        return sequencerHistory[result].sequencerAddr;
    }

    /// @notice Get the latest registered sequencer address (backward compat).
    /// @dev    If the latest record's startL2Block hasn't been reached yet,
    ///         this address is scheduled but not yet active.
    function getSequencer() external view returns (address) {
        require(sequencerHistory.length > 0, "no sequencer configured");
        return sequencerHistory[sequencerHistory.length - 1].sequencerAddr;
    }

    /// @notice Get the full sequencer history (for L2 node bulk sync at startup).
    function getSequencerHistory() external view returns (HistoryRecord[] memory) {
        return sequencerHistory;
    }

    /// @notice Get the number of sequencer history records.
    function getSequencerHistoryLength() external view returns (uint256) {
        return sequencerHistory.length;
    }
}
