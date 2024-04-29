// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/// @title Types
/// @notice Contains various types used throughout the Morph contract system.
library Types {
    /// @notice Struct representing a staker information.
    ///
    /// @custom:field addr   Address of the sequencer.
    /// @custom:field tmKey  Tendermint key(ED25519) of the seuqencer.
    /// @custom:field blsKey BLS key of the seuqencer.
    struct StakerInfo {
        address addr;
        bytes32 tmKey;
        bytes blsKey;
    }
}
