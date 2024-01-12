// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Test} from "forge-std/Test.sol";
import {MockBatchHeaderV0CodecCall} from "../../mock/MockCallBatchHeaderV0Codec.sol";
import {BatchHeaderV0Codec} from "../../libraries/codec/BatchHeaderV0Codec.sol";
import "forge-std/console.sol";

contract BatchHeaderV0Codec_Test is Test {
    function test_loadAndValidate() external {
        MockBatchHeaderV0CodecCall mockCall = new MockBatchHeaderV0CodecCall();
        bytes memory batchHeader0 = new bytes(89);
        // store dataHash
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }

        bytes32 dataHash = mockCall.mockDataHash(batchHeader0);
        console.logBytes32(dataHash);

        // store batchIndex
        assembly {
            mstore(add(batchHeader0, add(0x20, 1)), shl(192, 1))
        }
        uint256 index = mockCall.mockGetIndex(batchHeader0);
        console.log(index);
    }
}
