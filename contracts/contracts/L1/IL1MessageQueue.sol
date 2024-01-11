// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface IL1MessageQueue {
    /// @notice Return the message of in `queueIndex`.
    /// @param queueIndex The index to query.
    function getCrossDomainMessage(uint256 queueIndex) external view returns (bytes32);
    
    /// @notice Pop finalized messages from queue.
    ///
    /// @dev We can pop at most 256 messages each time. And if the message is not skipped,
    ///      the corresponding entry will be cleared.
    ///
    /// @param startIndex The start index to pop.
    /// @param count The number of messages to pop.
    /// @param skippedBitmap A bitmap indicates whether a message is skipped.
    function popCrossDomainMessage(
        uint256 startIndex,
        uint256 count,
        uint256 skippedBitmap
    ) external;

}
