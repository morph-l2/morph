// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/// @title L1Sequencer
/// @notice L1 contract for managing the sequencer address.
///         The sequencer address can be updated by the owner (multisig recommended).
contract L1Sequencer is OwnableUpgradeable {
    // ============ Storage ============

    /// @notice Current sequencer address
    address public sequencer;

    // ============ Events ============

    /// @notice Emitted when sequencer is updated
    event SequencerUpdated(address indexed oldSequencer, address indexed newSequencer);

    // ============ Initializer ============

    /// @notice Initialize the contract
    /// @param _owner Contract owner (multisig recommended)
    /// @param _initialSequencer Initial sequencer address (can be address(0) to set later)
    function initialize(address _owner, address _initialSequencer) external initializer {
        require(_owner != address(0), "invalid owner");

        __Ownable_init();
        _transferOwnership(_owner);

        // Set initial sequencer if provided
        if (_initialSequencer != address(0)) {
            sequencer = _initialSequencer;
            emit SequencerUpdated(address(0), _initialSequencer);
        }
    }

    // ============ Admin Functions ============

    /// @notice Update sequencer address (takes effect immediately)
    /// @param newSequencer New sequencer address
    function updateSequencer(address newSequencer) external onlyOwner {
        require(newSequencer != address(0), "invalid sequencer");
        require(newSequencer != sequencer, "same sequencer");

        address oldSequencer = sequencer;
        sequencer = newSequencer;

        emit SequencerUpdated(oldSequencer, newSequencer);
    }

    // ============ View Functions ============

    /// @notice Get current sequencer address
    function getSequencer() external view returns (address) {
        return sequencer;
    }
}
