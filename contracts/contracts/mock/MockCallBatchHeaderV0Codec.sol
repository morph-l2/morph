// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

import {BatchHeaderV0Codec} from "../libraries/codec/BatchHeaderV0Codec.sol";

contract MockBatchHeaderV0CodecCall {
    function loadAndValidate(
        bytes calldata _batchHeader
    ) public pure returns (uint256 batchPtr, uint256 length) {
        return BatchHeaderV0Codec.loadAndValidate(_batchHeader);
    }

    function mockGetIndex(
        bytes calldata _batchHeader
    ) public pure returns (uint256 index) {
        (uint256 memPtr, ) = BatchHeaderV0Codec.loadAndValidate(_batchHeader);
        uint256 _batchIndex = BatchHeaderV0Codec.batchIndex(memPtr);
        return _batchIndex;
    }

    function mockDataHash(
        bytes calldata _batchHeader
    ) public pure returns (bytes32 _dataHash) {
        (uint256 memPtr, ) = BatchHeaderV0Codec.loadAndValidate(_batchHeader);
        bytes32 dataHash = BatchHeaderV0Codec.dataHash(memPtr);
        return dataHash;
    }

    function mockBatchHash(
        bytes calldata _batchHeader
    ) public pure returns (bytes32 _batchHash) {
        uint256 _length;
        uint256 memPtr;

        (memPtr, _length) = BatchHeaderV0Codec.loadAndValidate(_batchHeader);
        // compute batch hash
        _batchHash = BatchHeaderV0Codec.computeBatchHash(memPtr, _length);
        return _batchHash;
    }
}
