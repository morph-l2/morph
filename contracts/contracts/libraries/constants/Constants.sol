// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title Constants
 * @notice Constants is a library for storing constants. Simple! Don't put everything in here, just
 *         the stuff used in multiple contracts. Constants that only apply to a single contract
 *         should be defined in that contract instead.
 */
library Constants {
    /**
     * @notice Value used for the L2 sender storage slot in both the MorphPortal and the
     *         CrossDomainMessenger contracts before an actual sender is set. This value is
     *         non-zero to reduce the gas cost of message passing transactions.
     */
    address internal constant DEFAULT_XDOMAIN_MESSAGE_SENDER = 0x000000000000000000000000000000000000dEaD;

    /// @notice The address for dropping message.
    /// @dev The first 20 bytes of keccak("drop")
    address internal constant DROP_XDOMAIN_MESSAGE_SENDER = 0x6f297C61B5C92eF107fFD30CD56AFFE5A273e841;
}
