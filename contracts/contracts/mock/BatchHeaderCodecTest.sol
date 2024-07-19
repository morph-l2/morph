// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {BatchHeaderCodecV0} from "../libraries/codec/BatchHeaderCodecV0.sol";

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

contract BatchHeaderCodecTest {
    struct BatchHeaderData {
        uint256 version;
        uint256 batchIndex;
        uint256 l1MessagePopped;
        uint256 totalL1MessagePopped;
        bytes32 dataHash;
        bytes32 blobVersionedHash;
        bytes32 prevStateHash;
        bytes32 postStateHash;
        bytes32 withdrawRootHash;
        bytes32 sequencerSetVerifyHash;
        bytes32 parentBatchHash;
        bytes skipMap;
    }

    function getVersion(bytes calldata _batchHeader) public pure returns (uint256 version) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getVersion(batchPtr);
    }

    function getBatchIndex(bytes calldata _batchHeader) public pure returns (uint256 batchIndex) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getBatchIndex(batchPtr);
    }

    function getL1MessagePopped(bytes calldata _batchHeader) public pure returns (uint256 l1MessagePopped) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getL1MessagePopped(batchPtr);
    }

    function getTotalL1MessagePopped(bytes calldata _batchHeader) public pure returns (uint256 totalL1MessagePopped) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getTotalL1MessagePopped(batchPtr);
    }

    function getL1DataHash(bytes calldata _batchHeader) public pure returns (bytes32 _dataHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getL1DataHash(batchPtr);
    }

    function getBlobVersionedHash(bytes calldata _batchHeader) public pure returns (bytes32 _blobVersionedHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getBlobVersionedHash(batchPtr);
    }

    function getPrevStateHash(bytes calldata _batchHeader) public pure returns (bytes32 _prevStateHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getPrevStateHash(batchPtr);
    }

    function getPostStateHash(bytes calldata _batchHeader) public pure returns (bytes32 _postStateHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getPostStateHash(batchPtr);
    }

    function getWithdrawRootHash(bytes calldata _batchHeader) public pure returns (bytes32 _withdrawRootHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getWithdrawRootHash(batchPtr);
    }

    function getSequencerSetVerifyHash(
        bytes calldata _batchHeader
    ) public pure returns (bytes32 _sequencerSetVerifyHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getSequencerSetVerifyHash(batchPtr);
    }

    function getParentBatchHash(bytes calldata _batchHeader) public pure returns (bytes32 _parentBatchHash) {
        (uint256 batchPtr, ) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.getParentBatchHash(batchPtr);
    }

    function computeBatchHash(bytes calldata _batchHeader) public pure returns (bytes32 batchHash) {
        (uint256 batchPtr, uint256 length) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        return BatchHeaderCodecV0.computeBatchHash(batchPtr, length);
    }

    function computeBatchHashWithData(BatchHeaderData calldata data) public pure returns (bytes32 batchHash) {
        uint256 _batchPtr;
        assembly {
            _batchPtr := mload(0x40) // reset batchPtr
        }
        BatchHeaderCodecV0.storeVersion(_batchPtr, data.version);
        BatchHeaderCodecV0.storeBatchIndex(_batchPtr, data.batchIndex);
        BatchHeaderCodecV0.storeL1MessagePopped(_batchPtr, data.l1MessagePopped);
        BatchHeaderCodecV0.storeTotalL1MessagePopped(_batchPtr, data.totalL1MessagePopped);
        BatchHeaderCodecV0.storeDataHash(_batchPtr, data.dataHash);

        BatchHeaderCodecV0.storeBlobVersionedHash(_batchPtr, data.blobVersionedHash);
        BatchHeaderCodecV0.storePrevStateHash(_batchPtr, data.prevStateHash);
        BatchHeaderCodecV0.storePostStateHash(_batchPtr, data.postStateHash);
        BatchHeaderCodecV0.storeWithdrawRootHash(_batchPtr, data.withdrawRootHash);
        BatchHeaderCodecV0.storeSequencerSetVerifyHash(_batchPtr, data.sequencerSetVerifyHash);
        BatchHeaderCodecV0.storeParentBatchHash(_batchPtr, data.parentBatchHash);
        BatchHeaderCodecV0.storeSkippedBitmap(_batchPtr, data.skipMap);
        return BatchHeaderCodecV0.computeBatchHash(_batchPtr, BatchHeaderCodecV0.BATCH_HEADER_FIXED_LENGTH);
    }
}
