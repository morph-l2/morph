// SPDX-License-Identifier: MIT

pragma solidity ^0.8.24;

// solhint-disable no-inline-assembly

/// @dev Below is the encoding for `BatchHeader` V1, total 249
///   * Field                   Bytes       Type        Index   Comments
///   * version                 1           uint8       0       The batch version
///   * batchIndex              8           uint64      1       The index of the batch
///   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch
///   * totalL1MessagePopped    8           uint64      17      Number of total L1 messages popped after the batch
///   * dataHash                32          bytes32     25      The data hash of the batch
///   * blobVersionedHash       32          bytes32     57      The versioned hash of the blob with this batch’s data
///   * prevStateHash           32          bytes32     89      Preview state root
///   * postStateHash           32          bytes32     121     Post state root
///   * withdrawRootHash        32          bytes32     153     L2 withdrawal tree root hash
///   * sequencerSetVerifyHash  32          bytes32     185     L2 sequencers set verify hash
///   * parentBatchHash         32          bytes32     217     The parent batch hash
/// 
/// @dev Below is the feilds for `BatchHeader` V1
///   * lastBlockNumber         8           uint64      249     The last block number in this batch
/// ```

library BatchHeaderCodecV1 {
    /// @dev The length of fixed parts of the batch header.
    uint256 internal constant BATCH_HEADER_LENGTH = 257;

    /// @notice Load batch header in calldata to memory.
    /// @param _batchHeader The encoded batch header bytes in calldata.
    /// @return batchPtr The start memory offset of the batch header in memory.
    /// @return length The length in bytes of the batch header.
    function loadAndValidate(bytes calldata _batchHeader) internal pure returns (uint256 batchPtr, uint256 length) {
        length = _batchHeader.length;
        require(length == BATCH_HEADER_LENGTH, "batch header length is incorrect");
        // copy batch header to memory.
        assembly {
            batchPtr := mload(0x40)
            calldatacopy(batchPtr, _batchHeader.offset, length)
            mstore(0x40, add(batchPtr, length))
        }
    }

    /// @notice Get the last block number of the batch.
    function getLastBlockNumber(uint256 batchPtr) internal pure returns (uint256 _lastBlockNumber) {
        assembly {
            _lastBlockNumber := shr(192, mload(add(batchPtr, 249)))
        }
    }

    /// @notice Store the last block number of the batch.
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @param _lastBlockNumber The last block number to store.
    function storeLastBlockNumber(uint256 batchPtr, uint256 _lastBlockNumber) internal pure {
        assembly {
            mstore(add(batchPtr, 249), shl(192, _lastBlockNumber))
        }
    }


    


}