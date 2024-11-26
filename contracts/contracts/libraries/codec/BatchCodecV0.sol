// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

/* solhint-disable */

/// @dev Below is the encoding for `Batch`, total 60*n+1 bytes.
/// ```text
///   * Field                     Bytes   Type      Index       Comments
///   * lastBlockNumber           8       uint64    0           The last block number in this batch
///   * block[0].numL1Messages    2       uint16    8           The number of L1 messages in the first block
///   * ......
///   * block[i].numL1Messages    2       uint16    2*i+8       The number of L1 messages in the i-th block
///   * ......
///   * block[n-1].numL1Messages  2       uint16    2*(n-1)+8   The number of L1 messages in the last block
/// ```

library BatchCodecV0 {
    /// @dev Thrown when no blocks in batch.
    error ErrorNoBlockInBatch();

    /// @dev Thrown when the length of batch is incorrect.
    error ErrorIncorrectBatchLength();

    /// @dev The length of one block context.
    uint256 internal constant BLOCK_CONTEXT_LENGTH = 2;

    /// @notice Validate the length of batch.
    /// @param batchPtr The start memory offset of the batch in memory.
    /// @param _preLastBlockNumber The last block number of the previous batch.
    /// @param _length The length of the batch.
    /// @return _numBlocks The number of blocks in current batch.
    function validateBatchLength(
        uint256 batchPtr,
        uint256 _preLastBlockNumber,
        uint256 _length
    ) internal pure returns (uint256 _numBlocks) {
        _numBlocks = getLastBlockNumber(batchPtr) - _preLastBlockNumber;

        // should contain at least one block
        if (_numBlocks == 0) revert ErrorNoBlockInBatch();

        // should contain the number of the blocks and block contexts
        if (_length != 2 + _numBlocks * BLOCK_CONTEXT_LENGTH) revert ErrorIncorrectBatchLength();
    }

    /// @notice Return the number of blocks in current batch.
    /// @param batchPtr The start memory offset of the batch in memory.
    /// @return _numBlocks The number of blocks in current batch.
    function getLastBlockNumber(uint256 batchPtr) internal pure returns (uint256 _numBlocks) {
        assembly {
            _numBlocks := shr(192, mload(batchPtr))
        }
    }

    /// @notice Copy the block context to another memory.
    /// @param blockPtr The start memory offset of the first block context in memory.
    /// @param dstPtr The destination memory offset to store the block context.
    /// @param index The index of block context to copy.
    /// @return uint256 The new destination memory offset after copy.
    function copyBlockContext(uint256 blockPtr, uint256 dstPtr, uint256 index) internal pure returns (uint256) {
        // only first 2 bytes is needed.
        assembly {
            blockPtr := add(blockPtr, mul(BLOCK_CONTEXT_LENGTH, index))
            mstore(dstPtr, and(mload(blockPtr), 0xffff000000000000000000000000000000000000000000000000000000000000)) // first 2 bytes
        }

        return dstPtr;
    }

    /// @notice Return the number of L1 messages in current block.
    /// @param blockPtr The start memory offset of the block context in memory.
    /// @return _numL1Messages The number of L1 messages in current block.
    function getNumL1Messages(uint256 blockPtr) internal pure returns (uint256 _numL1Messages) {
        assembly {
            _numL1Messages := shr(240, mload(blockPtr))
        }
    }
}
