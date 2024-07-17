// SPDX-License-Identifier: MIT

pragma solidity ^0.8.24;

/* Testing utilities */
import "forge-std/Test.sol";
import "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import {BatchHeaderCodecV0} from "../../libraries/codec/BatchHeaderCodecV0.sol";
import {BatchHeaderCodecTest} from "../../mock/BatchHeaderCodecTest.sol";

contract BatchHeaderCodeV0Test is DSTestPlus {
    bytes32 public constant ZERO_VERSIONED_HASH = 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;
    BatchHeaderCodecTest public codecTest;

    function setUp() public virtual {
        codecTest = new BatchHeaderCodecTest();
    }

    function test_btach_decode() public {
        bytes memory batchHeader0 = new bytes(249 + 32);
        assembly {
            mstore(add(batchHeader0, 0x20), shl(248, 1)) // version
            mstore(add(batchHeader0, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
            mstore(add(batchHeader0, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
            mstore(add(batchHeader0, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
            mstore(add(batchHeader0, add(0x20, 25)), ZERO_VERSIONED_HASH) // l1dataHash
            mstore(add(batchHeader0, add(0x20, 57)), ZERO_VERSIONED_HASH) // l2 tx blob versioned hash
            mstore(add(batchHeader0, add(0x20, 89)), ZERO_VERSIONED_HASH) // prevStateHash
            mstore(add(batchHeader0, add(0x20, 121)), ZERO_VERSIONED_HASH) // postStateHash
            mstore(add(batchHeader0, add(0x20, 153)), ZERO_VERSIONED_HASH) // withdrawRootHash
            mstore(add(batchHeader0, add(0x20, 185)), ZERO_VERSIONED_HASH) // sequencerSetVerifyHash
            mstore(add(batchHeader0, add(0x20, 217)), ZERO_VERSIONED_HASH) // parentBatchHash
            mstore(add(batchHeader0, add(0x20, 249)), 0) // bitmap0
        }

        uint256 version = codecTest.getVersion(batchHeader0);
        uint256 index = codecTest.getBatchIndex(batchHeader0);
        uint256 l1MessagePopped = codecTest.getL1MessagePopped(batchHeader0);
        bytes32 l1dataHash = codecTest.getL1DataHash(batchHeader0);
        bytes32 blobVersionedHash = codecTest.getBlobVersionedHash(batchHeader0);
        bytes32 prevStateHash = codecTest.getPrevStateHash(batchHeader0);
        bytes32 postStateHash = codecTest.getPostStateHash(batchHeader0);
        bytes32 withdrawRootHash = codecTest.getWithdrawRootHash(batchHeader0);
        bytes32 sequencerSetVerifyHash = codecTest.getSequencerSetVerifyHash(batchHeader0);
        bytes32 parentBatchHash = codecTest.getParentBatchHash(batchHeader0);

        assertEq(version, 1);
        assertEq(index, 1);
        assertEq(l1MessagePopped, 1);
        assertEq(l1dataHash, ZERO_VERSIONED_HASH);
        assertEq(blobVersionedHash, ZERO_VERSIONED_HASH);
        assertEq(prevStateHash, ZERO_VERSIONED_HASH);
        assertEq(postStateHash, ZERO_VERSIONED_HASH);
        assertEq(withdrawRootHash, ZERO_VERSIONED_HASH);
        assertEq(sequencerSetVerifyHash, ZERO_VERSIONED_HASH);
        assertEq(parentBatchHash, ZERO_VERSIONED_HASH);
    }

    function test_btach_decode1() public {
        bytes memory batchHeader0 = new bytes(249);
        bytes32 bytesData1 = bytes32(uint256(1));
        bytes32 bytesData4 = bytes32(uint256(4));
        assembly {
            mstore(add(batchHeader0, 0x20), 0) // version
            mstore(add(batchHeader0, add(0x20, 1)), shl(192, 1)) // batchIndex
            mstore(add(batchHeader0, add(0x20, 9)), 0) // l1MessagePopped
            mstore(add(batchHeader0, add(0x20, 17)), 0) // totalL1MessagePopped
            mstore(add(batchHeader0, add(0x20, 25)), 0x246394445f4fe64ed5598554d55d1682d6fb3fe04bf58eb54ef81d1189fafb51) // l1dataHash
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // l2 tx blob versioned hash
            mstore(add(batchHeader0, add(0x20, 89)), bytesData1) // prevStateHash
            mstore(add(batchHeader0, add(0x20, 121)), bytesData1) // postStateHash
            mstore(add(batchHeader0, add(0x20, 153)), bytesData4) // withdrawRootHash
            mstore(
                add(batchHeader0, add(0x20, 185)),
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
            ) // sequencerSetVerifyHash
            mstore(
                add(batchHeader0, add(0x20, 217)),
                0x5db199130d0c9334530358520a66d09c98c62325011a3665cf3efcc93e77623c
            ) // parentBatchHash
            mstore(add(batchHeader0, add(0x20, 249)), 0) // bitmap0
        }
        {
            uint256 version = codecTest.getVersion(batchHeader0);
            uint256 index = codecTest.getBatchIndex(batchHeader0);
            uint256 l1MessagePopped = codecTest.getL1MessagePopped(batchHeader0);
            bytes32 l1dataHash = codecTest.getL1DataHash(batchHeader0);
            bytes32 blobVersionedHash = codecTest.getBlobVersionedHash(batchHeader0);
            bytes32 prevStateHash = codecTest.getPrevStateHash(batchHeader0);
            bytes32 postStateHash = codecTest.getPostStateHash(batchHeader0);
            bytes32 withdrawRootHash = codecTest.getWithdrawRootHash(batchHeader0);
            bytes32 sequencerSetVerifyHash = codecTest.getSequencerSetVerifyHash(batchHeader0);
            bytes32 parentBatchHash = codecTest.getParentBatchHash(batchHeader0);
            bytes32 batchHash = codecTest.computeBatchHash(batchHeader0);

            assertEq(batchHash, 0xf14fe204d280f7483ae29385d880e7ba7a8a715c3450004da41cf82bfdd8415d);
            assertEq(version, 0);
            assertEq(index, 1);
            assertEq(l1MessagePopped, 0);
            assertEq(l1dataHash, 0x246394445f4fe64ed5598554d55d1682d6fb3fe04bf58eb54ef81d1189fafb51);
            assertEq(blobVersionedHash, 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014);
            assertEq(prevStateHash, bytesData1);
            assertEq(postStateHash, bytesData1);
            assertEq(withdrawRootHash, bytesData4);
            assertEq(sequencerSetVerifyHash, 0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067);
            assertEq(parentBatchHash, 0x5db199130d0c9334530358520a66d09c98c62325011a3665cf3efcc93e77623c);
        }
        {
            BatchHeaderCodecTest.BatchHeaderData memory data = BatchHeaderCodecTest.BatchHeaderData(
                0,
                1,
                0,
                0,
                0x246394445f4fe64ed5598554d55d1682d6fb3fe04bf58eb54ef81d1189fafb51,
                0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014,
                bytesData1,
                bytesData1,
                bytesData4,
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067,
                0x5db199130d0c9334530358520a66d09c98c62325011a3665cf3efcc93e77623c,
                "0x"
            );
            bytes32 batchHash2 = codecTest.computeBatchHashWithData(data);
            assertEq(batchHash2, 0xf14fe204d280f7483ae29385d880e7ba7a8a715c3450004da41cf82bfdd8415d);
        }
    }
}
