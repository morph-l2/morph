// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

/* solhint-disable */

/// @dev Below is the encoding for `Batch`, total 60*n+1 bytes.
/// ```text
///   * Field           Bytes       Type            Index       Comments
///   * numBlocks       2           uint16          0           The number of blocks in this batch
///   * block[0]        60          BlockContext    1           The first block in this batch
///   * ......
///   * block[i]        60          BlockContext    60*i+1      The (i+1)'th block in this batch
///   * ......
///   * block[n-1]      60          BlockContext    60*n-59     The last block in this batch
/// ```
///
/// @dev Below is the encoding for `BlockContext`, total 60 bytes.
/// ```text
///   * Field                   Bytes      Type         Index  Comments
///   * blockNumber             8          uint64       0      The height of this block.
///   * timestamp               8          uint64       8      The timestamp of this block.
///   * baseFee                 32         uint256      16     The base fee of this block.
///   * gasLimit                8          uint64       48     The gas limit of this block.
///   * numTransactions         2          uint16       56     The number of transactions in this block, both L1 & L2 txs.
///   * numL1Messages           2          uint16       58     The number of l1 messages in this block.
/// ```
library BatchCodecV0 {
    /// @dev Thrown when no blocks in batch.
    error ErrorNoBlockInBatch();

    /// @dev Thrown when the length of batch is incorrect.
    error ErrorIncorrectBatchLength();

    /// @dev The length of one block context.
    uint256 internal constant BLOCK_CONTEXT_LENGTH = 60;

    /// @notice Validate the length of batch.
    /// @param batchPtr The start memory offset of the batch in memory.
    /// @param _length The length of the batch.
    /// @return _numBlocks The number of blocks in current batch.
    function validateBatchLength(uint256 batchPtr, uint256 _length) internal pure returns (uint256 _numBlocks) {
        _numBlocks = getNumBlocks(batchPtr);

        // should contain at least one block
        if (_numBlocks == 0) revert ErrorNoBlockInBatch();

        // should contain the number of the blocks and block contexts
        if (_length != 2 + _numBlocks * BLOCK_CONTEXT_LENGTH) revert ErrorIncorrectBatchLength();
    }

    /// @notice Return the number of blocks in current batch.
    /// @param batchPtr The start memory offset of the batch in memory.
    /// @return _numBlocks The number of blocks in current batch.
    function getNumBlocks(uint256 batchPtr) internal pure returns (uint256 _numBlocks) {
        assembly {
            _numBlocks := shr(240, mload(batchPtr))
        }
    }

    /// @notice Copy the block context to another memory.
    /// @param blockPtr The start memory offset of the first block context in memory.
    /// @param dstPtr The destination memory offset to store the block context.
    /// @param index The index of block context to copy.
    /// @return uint256 The new destination memory offset after copy.
    function copyBlockContext(uint256 blockPtr, uint256 dstPtr, uint256 index) internal pure returns (uint256) {
        // only first 58 bytes is needed.
        assembly {
            blockPtr := add(blockPtr, mul(BLOCK_CONTEXT_LENGTH, index))
            mstore(dstPtr, mload(blockPtr)) // first 32 bytes
            mstore(
                add(dstPtr, 0x20),
                and(mload(add(blockPtr, 0x20)), 0xffffffffffffffffffffffffffffffffffffffffffffffffffff000000000000)
            ) // next 26 bytes

            dstPtr := add(dstPtr, 58)
        }

        return dstPtr;
    }

    /// @notice Return the number of transactions in current block.
    /// @param blockPtr The start memory offset of the block context in memory.
    /// @return _numTransactions The number of transactions in current block.
    function getNumTransactions(uint256 blockPtr) internal pure returns (uint256 _numTransactions) {
        assembly {
            _numTransactions := shr(240, mload(add(blockPtr, 56)))
        }
    }

    /// @notice Return the number of L1 messages in current block.
    /// @param blockPtr The start memory offset of the block context in memory.
    /// @return _numL1Messages The number of L1 messages in current block.
    function getNumL1Messages(uint256 blockPtr) internal pure returns (uint256 _numL1Messages) {
        assembly {
            _numL1Messages := shr(240, mload(add(blockPtr, 58)))
        }
    }

    /// @notice Return the number of the block.
    /// @param blockPtr The start memory offset of the block context in memory.
    /// @return _blockNumber The block number of blockPtr in current block.
    function getBlockNumber(uint256 blockPtr) internal pure returns (uint256 _blockNumber) {
        assembly {
            _blockNumber := shr(192, mload(blockPtr))
        }
    }
}