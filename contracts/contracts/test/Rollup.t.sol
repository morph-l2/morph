// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import "./CommonTest.t.sol";
import "forge-std/console.sol";
import {Rollup} from "../L1/Rollup.sol";
import {IRollup} from "../L1/IRollup.sol";
import {ChunkCodec} from "../libraries/codec/ChunkCodec.sol";

contract Rollup_Test is Portal_Initializer {
    address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
    bytes32 public l1MessagerIndex = bytes32(uint256(55));

    function test_stake_withdraw() external {
        assertEq(true, rollup.isSequencer(alice));
        assertEq(true, rollup.isProver(alice));
        assertEq(false, rollup.isChallenger(alice));

        assertEq(false, rollup.isSequencer(bob));
        assertEq(true, rollup.isProver(bob));
        assertEq(true, rollup.isChallenger(bob));

        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        rollup.stake{value: MIN_DEPOSIT}();
        assertEq(MIN_DEPOSIT, rollup.deposits(alice));

        rollup.withdraw(MIN_DEPOSIT);
        assertEq(0, rollup.deposits(alice));
    }

    function testAddAndRemoveSequencer(address _sequencer) public {
        // set by non-owner, should revert
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.addSequencer(_sequencer);
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.removeSequencer(_sequencer);
        vm.stopPrank();

        // change to random EOA operator
        vm.expectEmit(true, false, false, true);
        emit UpdateSequencer(_sequencer, true);

        assertEq(rollup.isSequencer(_sequencer), false);
        vm.prank(multisig);
        rollup.addSequencer(_sequencer);
        assertEq(rollup.isSequencer(_sequencer), true);

        vm.expectEmit(true, false, false, true);
        emit UpdateSequencer(_sequencer, false);
        vm.prank(multisig);
        rollup.removeSequencer(_sequencer);
        assertEq(rollup.isSequencer(_sequencer), false);
    }

    function test_stake_revertCallerNotSequencer() external {
        // bob submit batch: revert with caller not sequencer
        assertEq(true, rollup.isSequencer(alice));
        assertEq(false, rollup.isSequencer(bob));
        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.startPrank(bob);
        vm.expectRevert("caller not sequencer");
        rollup.stake{value: MIN_DEPOSIT}();
        vm.stopPrank();
    }

    function test_comitBatchs() external {
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();

        bytes memory batchHeader0 = new bytes(89);
        bytes32 stateRoot = bytes32(uint256(1));
        // import 10 L1 messages
        for (uint256 i = 0; i < 10; i++) {
            L1Messenger.sendMessage(address(this), new bytes(0), 1000000);
        }

        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );
        // caller not sequencer, revert
        vm.expectRevert("caller not sequencer");
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            new bytes[](0),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);

        // invalid version, revert
        vm.startPrank(address(0));
        vm.expectRevert("invalid version");
        batchData = IRollup.BatchData(
            1,
            batchHeader0,
            new bytes[](0),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // batch is empty, revert
        vm.startPrank(address(0));
        vm.expectRevert("batch is empty");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            new bytes[](0),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // batch header length too small, revert
        vm.startPrank(address(0));
        vm.expectRevert("batch header length too small");
        batchData = IRollup.BatchData(
            0,
            new bytes(88),
            new bytes[](1),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // wrong bitmap length, revert
        vm.startPrank(address(0));
        vm.expectRevert("wrong bitmap length");
        batchData = IRollup.BatchData(
            0,
            new bytes(90),
            new bytes[](1),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // incorrect parent batch hash, revert
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 2) // change data hash for batch0
        }
        vm.startPrank(address(0));
        vm.expectRevert("incorrect parent batch hash");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            new bytes[](1),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1) // change back
        }

        bytes[] memory chunks = new bytes[](1);
        bytes memory chunk0;
        // no block in chunk, revert
        chunk0 = new bytes(1);
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        vm.expectRevert("no block in chunk");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // invalid chunk length, revert
        chunk0 = new bytes(1);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        vm.expectRevert("invalid chunk length");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // cannot skip last L1 message, revert
        chunk0 = new bytes(1 + 60);
        bytes memory bitmap = new bytes(32);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunk0[58] = bytes1(uint8(1)); // numTransactions = 1
        chunk0[60] = bytes1(uint8(1)); // numL1Messages = 1
        bitmap[31] = bytes1(uint8(1));
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        vm.expectRevert("cannot skip last L1 message");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // num txs less than num L1 msgs, revert
        chunk0 = new bytes(1 + 60);
        bitmap = new bytes(32);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunk0[58] = bytes1(uint8(1)); // numTransactions = 1
        chunk0[60] = bytes1(uint8(3)); // numL1Messages = 3
        bitmap[31] = bytes1(uint8(3));
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        vm.expectRevert("num txs less than num L1 msgs");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // incomplete l2 transaction data, revert
        chunk0 = new bytes(1 + 60 + 1);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        vm.expectRevert("incomplete l2 transaction data");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // commit batch with one chunk, no tx, correctly
        chunk0 = new bytes(1 + 60);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();
        assertGt(uint256(rollup.committedBatches(1)), 0);

        // batch is already committed, revert
        vm.startPrank(address(0));
        vm.expectRevert("batch already committed");
        batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();
    }

    function testFinalizeBatchs() public {
        rollup.finalizeBatchs();
    }

    function testCommitAndFinalizeWithL1Messages() public {
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();

        // update portal l1Messager to caller
        vm.store(
            address(portal),
            bytes32(l1MessagerIndex),
            bytes32(abi.encode(caller))
        );
        vm.deal(caller, 5 * MIN_DEPOSIT);
        vm.startPrank(caller);

        bytes memory batchHeader0 = new bytes(89);
        bytes32 stateRoot = bytes32(uint256(1));

        // import 300 L1 messages
        for (uint256 i = 0; i < 300; i++) {
            portal.depositTransaction(caller, 0, 1000000, false, new bytes(0));
            vm.roll(block.number + 1);
        }
        vm.stopPrank();
        
        vm.store(
            address(portal),
            bytes32(l1MessagerIndex),
            bytes32(abi.encode(address(L1Messenger)))
        );

        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );
        bytes32 batchHash0 = rollup.committedBatches(0);

        bytes memory bitmap;
        bytes[] memory chunks;
        bytes memory chunk0;
        bytes memory chunk1;

        // commit batch1, one chunk with one block, 1 tx, 1 L1 message, no skip
        // => payload for data hash of chunk0
        //   0000000000000000
        //   0000000000000000
        //   0000000000000000000000000000000000000000000000000000000000000000
        //   0000000000000000
        //   0001
        //   a2277fd30bbbe74323309023b56035b376d7768ad237ae4fc46ead7dc9591ae1
        // => data hash for chunk0
        //   9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110
        // => data hash for all chunks
        //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
        // => payload for batch header
        //   00
        //   0000000000000001
        //   0000000000000001
        //   0000000000000001
        //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
        //   119b828c2a2798d2c957228ebeaff7e10bb099ae0d4e224f3eeb779ff61cba61
        //   0000000000000000000000000000000000000000000000000000000000000000
        // => hash for batch header
        //   00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
        bytes memory batchHeader1 = new bytes(89 + 32);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
            mstore(add(batchHeader1, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
            mstore(
                add(batchHeader1, add(0x20, 25)),
                0xd9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
            ) // dataHash
            mstore(add(batchHeader1, add(0x20, 57)), batchHash0) // parentBatchHash
            mstore(add(batchHeader1, add(0x20, 89)), 0) // bitmap0
        }
        chunk0 = new bytes(1 + 60);
        assembly {
            mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
            mstore(add(chunk0, add(0x21, 56)), shl(240, 1)) // numTransactions = 1
            mstore(add(chunk0, add(0x21, 58)), shl(240, 1)) // numL1Messages = 1
        }
        chunks = new bytes[](1);
        chunks[0] = chunk0;
        bitmap = new bytes(32);
        vm.startPrank(address(0));
        vm.expectEmit(true, true, false, true);
        emit CommitBatch(
            1,
            bytes32(
                0x00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
            )
        );
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            bytes32(uint256(3)),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        assertEq(rollup.isBatchFinalized(1), false);
        bytes32 batchHash1 = rollup.committedBatches(1);
        assertEq(
            batchHash1,
            bytes32(
                0x00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
            )
        );

        // finalize batch1
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        assertEq(rollup.isBatchFinalized(1), true);
        assertEq(rollup.finalizedStateRoots(1), stateRoot);
        assertEq(rollup.withdrawalRoots(bytes32(uint256(3))), 1);
        assertEq(rollup.lastFinalizedBatchIndex(), 1);
        assertEq(portal.isMessageSkipped(0), false);
        assertEq(portal.pendingQueueIndex(), 1);

        // commit batch2 with two chunks, correctly
        // 1. chunk0 has one block, 3 tx, no L1 messages
        //   => payload for chunk0
        //    0000000000000000
        //    0000000000000000
        //    0000000000000000000000000000000000000000000000000000000000000000
        //    0000000000000000
        //    0003
        //    ... (some tx hashes)
        //   => data hash for chunk0
        //    2ac1dad3f3696e5581dfc10f2c7a7a8fc5b344285f7d332c7895a8825fca609a
        // 2. chunk1 has three blocks
        //   2.1 block0 has 5 tx, 3 L1 messages, no skips
        //   2.2 block1 has 10 tx, 5 L1 messages, even is skipped, last is not skipped
        //   2.2 block1 has 300 tx, 256 L1 messages, odd position is skipped, last is not skipped
        //   => payload for chunk1
        //    0000000000000000
        //    0000000000000000
        //    0000000000000000000000000000000000000000000000000000000000000000
        //    0000000000000000
        //    0005
        //    0000000000000000
        //    0000000000000000
        //    0000000000000000000000000000000000000000000000000000000000000000
        //    0000000000000000
        //    000a
        //    0000000000000000
        //    0000000000000000
        //    0000000000000000000000000000000000000000000000000000000000000000
        //    0000000000000000
        //    012c
        //    ... (some tx hashes)
        //   => data hash for chunk2
        //    e1276f58354ab2372050bde30d8c970ccc3728c76e97f37deebeee83ecbf5705
        // => data hash for all chunks
        //   3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
        // => payload for batch header
        //  00
        //  0000000000000002
        //  0000000000000108
        //  0000000000000109
        //  3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
        //  cef70bf80683c4d9b8b2813e90c314e8c56648e231300b8cfed9d666b0caf14e
        //  aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa800000000000000000000000000000000000000000000000000000000000000aa
        // => hash for batch header
        //  03a9cdcb9d582251acf60937db006ec99f3505fd4751b7c1f92c9a8ef413e873
        bytes memory batchHeader2 = new bytes(89 + 32 + 32);
        assembly {
            mstore(add(batchHeader2, 0x20), 0) // version
            mstore(add(batchHeader2, add(0x20, 1)), shl(192, 2)) // batchIndex = 2
            mstore(add(batchHeader2, add(0x20, 9)), shl(192, 264)) // l1MessagePopped = 264
            mstore(add(batchHeader2, add(0x20, 17)), shl(192, 265)) // totalL1MessagePopped = 265
            mstore(
                add(batchHeader2, add(0x20, 25)),
                0x3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
            ) // dataHash
            mstore(add(batchHeader2, add(0x20, 57)), batchHash1) // parentBatchHash
            mstore(
                add(batchHeader2, add(0x20, 89)),
                77194726158210796949047323339125271902179989777093709359638389338608753093160
            ) // bitmap0
            mstore(add(batchHeader2, add(0x20, 121)), 42) // bitmap1
        }
        chunk0 = new bytes(1 + 60 + 3 * 5);
        assembly {
            mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
            mstore(add(chunk0, add(0x21, 56)), shl(240, 3)) // numTransactions = 3
            mstore(add(chunk0, add(0x21, 58)), shl(240, 0)) // numL1Messages = 0
        }
        for (uint256 i = 0; i < 3; i++) {
            assembly {
                mstore(add(chunk0, add(93, mul(i, 5))), shl(224, 1)) // tx = "0x00"
            }
        }
        chunk1 = new bytes(1 + 60 * 3 + 51 * 5);
        assembly {
            mstore(add(chunk1, 0x20), shl(248, 3)) // numBlocks = 3
            mstore(add(chunk1, add(33, 56)), shl(240, 5)) // block0.numTransactions = 5
            mstore(add(chunk1, add(33, 58)), shl(240, 3)) // block0.numL1Messages = 3
            mstore(add(chunk1, add(93, 56)), shl(240, 10)) // block1.numTransactions = 10
            mstore(add(chunk1, add(93, 58)), shl(240, 5)) // block1.numL1Messages = 5
            mstore(add(chunk1, add(153, 56)), shl(240, 300)) // block1.numTransactions = 300
            mstore(add(chunk1, add(153, 58)), shl(240, 256)) // block1.numL1Messages = 256
        }
        for (uint256 i = 0; i < 51; i++) {
            assembly {
                mstore(add(chunk1, add(213, mul(i, 5))), shl(224, 1)) // tx = "0x00"
            }
        }
        chunks = new bytes[](2);
        chunks[0] = chunk0;
        chunks[1] = chunk1;
        bitmap = new bytes(64);
        assembly {
            mstore(
                add(bitmap, add(0x20, 0)),
                77194726158210796949047323339125271902179989777093709359638389338608753093160
            ) // bitmap0
            mstore(add(bitmap, add(0x20, 32)), 42) // bitmap1
        }

        // too many txs in one chunk, revert
        vm.prank(multisig);
        rollup.updateMaxNumTxInChunk(2);
        vm.startPrank(address(0));
        vm.expectRevert("too many txs in one chunk");
        batchData = IRollup.BatchData(
            0,
            batchHeader1,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            bytes32(uint256(4)),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit); // first chunk with too many txs

        vm.stopPrank();
        vm.prank(multisig);
        rollup.updateMaxNumTxInChunk(10);
        vm.startPrank(address(0));
        vm.expectRevert("too many txs in one chunk");
        batchData = IRollup.BatchData(
            0,
            batchHeader1,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            bytes32(uint256(4)),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit); // second chunk with too many txs
        vm.stopPrank();

        vm.prank(multisig);
        rollup.updateMaxNumTxInChunk(186);
        vm.startPrank(address(0));
        vm.expectEmit(true, true, false, true);
        emit CommitBatch(
            2,
            bytes32(
                0x03a9cdcb9d582251acf60937db006ec99f3505fd4751b7c1f92c9a8ef413e873
            )
        );
        batchData = IRollup.BatchData(
            0,
            batchHeader1,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            bytes32(uint256(5)),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();
        assertEq(rollup.isBatchFinalized(2), false);
        bytes32 batchHash2 = rollup.committedBatches(2);
        assertEq(
            batchHash2,
            bytes32(
                0x03a9cdcb9d582251acf60937db006ec99f3505fd4751b7c1f92c9a8ef413e873
            )
        );

        // verify committed batch correctly
        vm.startPrank(address(0));
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS());
        rollup.finalizeBatchs();
        vm.stopPrank();

        assertEq(rollup.isBatchFinalized(2), true);
        assertEq(rollup.finalizedStateRoots(2), stateRoot);
        assertEq(rollup.withdrawalRoots(bytes32(uint256(5))), 2);
        assertEq(rollup.lastFinalizedBatchIndex(), 2);
        assertEq(portal.pendingQueueIndex(), 265);
        // 1 ~ 4, zero
        for (uint256 i = 1; i < 4; i++) {
            assertEq(portal.isMessageSkipped(i), false);
        }
        // 4 ~ 9, even is nonzero, odd is zero
        for (uint256 i = 4; i < 9; i++) {
            if (i % 2 == 1 || i == 8) {
                assertEq(portal.isMessageSkipped(i), false);
            } else {
                assertEq(portal.isMessageSkipped(i), true);
            }
        }
        // 9 ~ 265, even is nonzero, odd is zero
        for (uint256 i = 9; i < 265; i++) {
            if (i % 2 == 1 || i == 264) {
                assertEq(portal.isMessageSkipped(i), false);
            } else {
                assertEq(portal.isMessageSkipped(i), true);
            }
        }
    }
}

contract Rollup_Rollup_Test is Portal_Initializer {
    uint64 public preSubmitNum = 5;
    uint64 public challengeBatchIndex = 3;

    function setUp() public virtual override {
        super.setUp();
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();
        bytes memory batchHeader0 = new bytes(89);
        bytes32 stateRoot = bytes32(uint256(1));
        // import 10 L1 messages
        for (uint256 i = 0; i < 10; i++) {
            L1Messenger.sendMessage(address(this), new bytes(0), 1000000);
        }
        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );

        bytes[] memory chunks = new bytes[](1);
        bytes memory chunk0;
        // commit batch with one chunk, no tx, correctly
        chunk0 = new bytes(1 + 60 + 3 * 5);
        assembly {
            mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
            mstore(add(chunk0, add(0x21, 0)), shl(192, 123)) // blockNumber = 123
            mstore(add(chunk0, add(0x21, 56)), shl(240, 3)) // numTransactions = 3
            mstore(add(chunk0, add(0x21, 58)), shl(240, 0)) // numL1Messages = 0
        }
        for (uint256 i = 0; i < 3; i++) {
            assembly {
                mstore(add(chunk0, add(93, mul(i, 5))), shl(224, 1)) // tx = "0x00"
            }
        }
        chunks[0] = chunk0;
        vm.startPrank(address(0));
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();
        assertGt(uint256(rollup.committedBatches(1)), 0);
        assertEq(rollup.latestL2BlockNumber(), 123);
    }

    function test_challenge_revert() external {
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);
        rollup.stake{value: MIN_DEPOSIT}();
        // rollup.submitBatches(batches);
        vm.stopPrank();
        uint64 wrongIndex = 2;
        uint64 challengeIndex = 1;
        // bob challenge batch 1 : revert with batch not exist
        vm.prank(bob);
        vm.deal(bob, 5 * MIN_DEPOSIT);
        vm.expectRevert("batch not exist");
        rollup.challengeState{value: MIN_DEPOSIT}(wrongIndex);
        // chloe challenge batch 0 : revert with caller not challenger
        address chloe = address(1024);
        vm.deal(chloe, 5 * MIN_DEPOSIT);
        vm.expectRevert("caller not challenger");
        vm.prank(chloe);
        rollup.challengeState{value: MIN_DEPOSIT}(wrongIndex);
        // bob challenge btach 0 twice: revert with already has challenge
        vm.prank(bob);
        rollup.challengeState{value: MIN_DEPOSIT}(challengeIndex);
        vm.expectRevert("already has challenge");
        // bob challenge state batch 1
        vm.prank(bob);
        rollup.challengeState{value: MIN_DEPOSIT}(challengeIndex);

        // prove time out
        vm.prank(alice);
        (, , , uint256 startTime_, ) = rollup.challenges(challengeIndex);
        vm.warp(startTime_ + PROOF_WINDOW + 1);
        vm.prank(bob);
        vm.expectEmit(true, true, true, true);
        emit ChallengeRes(challengeIndex, bob, "timeout");
        rollup.proveState(challengeIndex, "");
    }

    function testAddAndRemoveProver() public {
        address _prover = address(1234567);
        // set by non-owner, should revert
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.addProver(_prover);
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.removeProver(_prover);
        vm.stopPrank();

        // change to random EOA operator
        vm.expectEmit(true, false, false, true);
        emit UpdateProver(_prover, true);

        assertEq(rollup.isProver(_prover), false);
        vm.prank(multisig);
        rollup.addProver(_prover);
        assertEq(rollup.isProver(_prover), true);

        vm.expectEmit(true, false, false, true);
        emit UpdateProver(_prover, false);
        vm.prank(multisig);
        rollup.removeProver(_prover);
        assertEq(rollup.isProver(_prover), false);
    }

    function testSetPause() external {
        vm.startPrank(multisig);
        rollup.addSequencer(address(0));
        rollup.addProver(address(0));
        vm.stopPrank();

        // not owner, revert
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.setPause(false);
        vm.stopPrank();

        // pause
        vm.prank(multisig);
        rollup.setPause(true);
        assertEq(true, rollup.paused());

        vm.startPrank(address(0));
        vm.expectRevert("Pausable: paused");
        bytes memory batchHeader0 = new bytes(89);
        bytes32 stateRoot = bytes32(uint256(1));
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            new bytes[](0),
            new bytes(0),
            stateRoot,
            stateRoot,
            getTreeRoot(),
            nilBatchSig
        );
        rollup.commitBatch(batchData, minGasLimit);
        vm.expectRevert("Pausable: paused");
        rollup.finalizeBatchs();
        vm.stopPrank();

        // unpause
        vm.prank(multisig);
        rollup.setPause(false);
        assertEq(false, rollup.paused());
    }

    function testUpdateVerifier(address _newVerifier) public {
        // set by non-owner, should revert
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.updateVerifier(_newVerifier);
        vm.stopPrank();

        // change to random operator
        vm.expectEmit(true, true, false, true);
        emit UpdateVerifier(address(verifier), _newVerifier);

        assertEq(rollup.verifier(), address(verifier));
        vm.prank(multisig);
        rollup.updateVerifier(_newVerifier);
        assertEq(rollup.verifier(), _newVerifier);
    }

    function testUpdateMaxNumTxInChunk(uint256 _maxNumTxInChunk) public {
        // set by non-owner, should revert
        vm.startPrank(address(1));
        vm.expectRevert("Ownable: caller is not the owner");
        rollup.updateMaxNumTxInChunk(_maxNumTxInChunk);
        vm.stopPrank();

        // change to random operator
        vm.expectEmit(false, false, false, true);
        emit UpdateMaxNumTxInChunk(rollup.maxNumTxInChunk(), _maxNumTxInChunk);

        vm.prank(multisig);
        rollup.updateMaxNumTxInChunk(_maxNumTxInChunk);
        assertEq(rollup.maxNumTxInChunk(), _maxNumTxInChunk);
    }
}
