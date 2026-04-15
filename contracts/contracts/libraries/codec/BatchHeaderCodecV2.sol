// SPDX-License-Identifier: MIT

pragma solidity ^0.8.24;

// solhint-disable no-inline-assembly

/// @dev Below is the encoding for `BatchHeader` V2, variable length
///   The first 257 bytes are identical to V1:
///   * Field                   Bytes       Type        Index   Comments
///   * version                 1           uint8       0       The batch version (must be 2)
///   * batchIndex              8           uint64      1       The index of the batch
///   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch
///   * totalL1MessagePopped    8           uint64      17      Number of total L1 messages popped after the batch
///   * dataHash                32          bytes32     25      The data hash of the batch
///   * blobVersionedHash[0]    32          bytes32     57      The versioned hash of blob[0] (same slot as V0/V1)
///   * prevStateHash           32          bytes32     89      Preview state root
///   * postStateHash           32          bytes32     121     Post state root
///   * withdrawRootHash        32          bytes32     153     L2 withdrawal tree root hash
///   * sequencerSetVerifyHash  32          bytes32     185     L2 sequencers set verify hash
///   * parentBatchHash         32          bytes32     217     The parent batch hash
///   * lastBlockNumber         8           uint64      249     The last block number in this batch (V1 field)
///
///   V2-specific fields appended after V1:
///   * blobCount               1           uint8       257     Total number of blobs (must be >= 1)
///   * blobVersionedHash[1]    32          bytes32     258     Versioned hash of blob[1] (omitted if blobCount == 1)
///   * blobVersionedHash[2]    32          bytes32     290     Versioned hash of blob[2] (omitted if blobCount <= 2)
///   * ...
///   * blobVersionedHash[N-1]  32          bytes32     258+(N-2)*32
///
///   Total length: 258 + (blobCount - 1) * 32
///   Minimum length (blobCount=1): 258 bytes

library BatchHeaderCodecV2 {
    /// @dev Base length = 257 (V1) + 1 (blobCount field) = 258
    uint256 internal constant BASE_LENGTH = 258;

    /// @notice Load V2 batch header from calldata to memory and validate its length.
    /// @param _batchHeader The encoded batch header bytes in calldata.
    /// @return batchPtr The start memory offset of the batch header in memory.
    /// @return length The length in bytes of the batch header.
    function loadAndValidate(bytes calldata _batchHeader) internal pure returns (uint256 batchPtr, uint256 length) {
        length = _batchHeader.length;
        require(length >= BASE_LENGTH, "batch header length too small");
        uint8 blobCount = uint8(_batchHeader[257]);
        require(blobCount > 0, "blob count must be at least 1");
        require(length == BASE_LENGTH + uint256(blobCount - 1) * 32, "batch header length mismatch");
        assembly {
            batchPtr := mload(0x40)
            calldatacopy(batchPtr, _batchHeader.offset, length)
            mstore(0x40, add(batchPtr, length))
        }
    }

    /// @notice Get the number of blobs in the V2 batch header.
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @return blobCount The number of blobs.
    function getBlobCount(uint256 batchPtr) internal pure returns (uint8 blobCount) {
        assembly {
            blobCount := shr(248, mload(add(batchPtr, 257)))
        }
    }

    /// @notice Store the number of blobs into the V2 batch header.
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @param blobCount The number of blobs to store.
    function storeBlobCount(uint256 batchPtr, uint8 blobCount) internal pure {
        assembly {
            mstore8(add(batchPtr, 257), blobCount)
        }
    }

    /// @notice Get the versioned hash of blob[i] from the V2 batch header.
    /// @dev blob[0] is stored at offset 57 (same as V0/V1 for compatibility).
    ///      blob[i] for i >= 1 is stored at offset 258 + (i-1)*32.
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @param i The blob index (0-based).
    /// @return hash The versioned hash of blob[i].
    function getBlobVersionedHash(uint256 batchPtr, uint8 i) internal pure returns (bytes32 hash) {
        if (i == 0) {
            assembly {
                hash := mload(add(batchPtr, 57))
            }
        } else {
            uint256 off = 258 + uint256(i - 1) * 32;
            assembly {
                hash := mload(add(batchPtr, off))
            }
        }
    }

    /// @notice Store the versioned hash of blob[i] into the V2 batch header.
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @param i The blob index (0-based).
    /// @param hash The versioned hash to store.
    function storeBlobVersionedHash(uint256 batchPtr, uint8 i, bytes32 hash) internal pure {
        if (i == 0) {
            assembly {
                mstore(add(batchPtr, 57), hash)
            }
        } else {
            uint256 off = 258 + uint256(i - 1) * 32;
            assembly {
                mstore(add(batchPtr, off), hash)
            }
        }
    }

    /// @notice Compute the aggregate hash of all blob versioned hashes: keccak256(hash[0] || ... || hash[N-1]).
    /// @dev Used as the blobHashInput in the ZK public input hash for V2 batches.
    ///      Uses scratch memory starting at the free pointer; does not advance the free pointer
    ///      (safe for pure/view functions).
    /// @param batchPtr The start memory offset of the batch header in memory.
    /// @param n The number of blobs.
    /// @return result The keccak256 aggregate hash.
    function getBlobHashesHash(uint256 batchPtr, uint8 n) internal pure returns (bytes32 result) {
        assembly {
            let dataPtr := mload(0x40)
            // copy blob[0] from offset 57
            mstore(dataPtr, mload(add(batchPtr, 57)))
            // copy blob[1..n-1] from offset 258+(i-1)*32
            for {
                let i := 1
            } lt(i, n) {
                i := add(i, 1)
            } {
                let srcOff := add(258, mul(sub(i, 1), 32))
                mstore(add(dataPtr, mul(i, 32)), mload(add(batchPtr, srcOff)))
            }
            result := keccak256(dataPtr, mul(n, 32))
        }
    }
}