// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

// import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

// import {ChunkCodecV0} from "../libraries/codec/ChunkCodecV0.sol";
// import {L1MessageQueueWithGasPriceOracle} from "../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";
// import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
// import {Types} from "../libraries/common/Types.sol";
// import {Rollup} from "../L1/rollup/Rollup.sol";
// import {IRollup} from "../L1/rollup/IRollup.sol";
// import {IL1Staking} from "../L1/staking/IL1Staking.sol";

// contract RollupCommitBatchTest is L1MessageBaseTest {
//     function setUp() public virtual override {
//         super.setUp();
//         nilBatchSig = IRollup.BatchSignatureData({
//             signedSequencers: new uint256[](0),
//             sequencerSets: "0x123456",
//             signature: "0x123456"
//         });
//     }

//     function testCommitAndFinalizeWithL1Messages() public {
//         hevm.mockCall(
//             address(rollup.l1SequencerContract()),
//             abi.encodeCall(IL1Staking.verifySignature, ()),
//             abi.encode(true)
//         );
//         upgradeStorage(address(caller), address(rollup), address(alice));
//         hevm.deal(caller, 5 * STAKING_VALUE);
//         bytes memory batchHeader0 = new bytes(121);

//         hevm.startPrank(caller);
//         // import 300 L1 messages
//         for (uint256 i = 0; i < 300; i++) {
//             l1MessageQueueWithGasPriceOracle.appendCrossDomainMessage(
//                 address(caller),
//                 1000000,
//                 new bytes(0)
//             );
//         }
//         assertEq(
//             l1MessageQueueWithGasPriceOracle.getCrossDomainMessage(0),
//             bytes32(
//                 0xa2277fd30bbbe74323309023b56035b376d7768ad237ae4fc46ead7dc9591ae1
//             )
//         );
//         hevm.stopPrank();

//         upgradeStorage(
//             address(l1CrossDomainMessenger),
//             address(rollup),
//             address(alice)
//         );

//         // import genesis batch first
//         assembly {
//             mstore(add(batchHeader0, add(0x20, 25)), 1)
//             mstore(
//                 add(batchHeader0, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             )
//         }
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader0, bytes32(uint256(1)));
//         bytes32 batchHash0 = rollup.committedBatches(0);

//         bytes memory bitmap;
//         bytes[] memory chunks;
//         bytes memory chunk0;
//         bytes memory chunk1;

//         // commit batch1, one chunk with one block, 1 tx, 1 L1 message, no skip
//         // => payload for data hash of chunk0
//         //   0000000000000000
//         //   0000000000000000
//         //   0000000000000000000000000000000000000000000000000000000000000000
//         //   0000000000000000
//         //   0001
//         //   a2277fd30bbbe74323309023b56035b376d7768ad237ae4fc46ead7dc9591ae1
//         // => data hash for chunk0
//         //   9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110
//         // => data hash for all chunks
//         //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
//         // => payload for batch header
//         //   00
//         //   0000000000000001
//         //   0000000000000001
//         //   0000000000000001
//         //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
//         //   119b828c2a2798d2c957228ebeaff7e10bb099ae0d4e224f3eeb779ff61cba61
//         //   0000000000000000000000000000000000000000000000000000000000000000
//         // => hash for batch header
//         //   00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
//         bytes memory batchHeader1 = new bytes(121 + 32);
//         assembly {
//             mstore(add(batchHeader1, 0x20), 0) // version
//             mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
//             mstore(add(batchHeader1, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
//             mstore(add(batchHeader1, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
//             mstore(
//                 add(batchHeader1, add(0x20, 25)),
//                 0xd9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
//             ) // l1dataHash
//             mstore(
//                 add(batchHeader1, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             ) // l2 tx blob versioned hash
//             mstore(add(batchHeader1, add(0x20, 89)), batchHash0) // parentBatchHash
//             mstore(add(batchHeader1, add(0x20, 121)), 0) // bitmap0
//         }
//         chunk0 = new bytes(1 + 60);
//         assembly {
//             mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
//             mstore(add(chunk0, add(0x21, 56)), shl(240, 1)) // numTransactions = 1
//             mstore(add(chunk0, add(0x21, 58)), shl(240, 1)) // numL1Messages = 1
//         }
//         chunks = new bytes[](1);
//         chunks[0] = chunk0;
//         bitmap = new bytes(32);
//         hevm.mockCall(
//             address(rollup.l1SequencerContract()),
//             abi.encodeCall(IL1Staking.isStaker, ()),
//             abi.encode(true)
//         );
//         hevm.startPrank(address(0));
//         hevm.expectEmit(true, true, false, true);
//         emit CommitBatch(
//             1,
//             bytes32(
//                 0xf8e8e938cb644f2207a1775542ad965c7d779d7ff70c459a6f54aab9c8057b08
//             )
//         );
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(3)),
//             nilBatchSig
//         );
//         rollup.commitBatch(batchData, sequencerSigned, signature);
//         hevm.stopPrank();

//         assertFalse(rollup.isBatchFinalized(1));
//         bytes32 batchHash1 = rollup.committedBatches(1);
//         assertEq(
//             batchHash1,
//             bytes32(
//                 0xf8e8e938cb644f2207a1775542ad965c7d779d7ff70c459a6f54aab9c8057b08
//             )
//         );

//         // finalize batch1
//         hevm.warp(block.timestamp + rollup.finalizationPeriodSeconds() + 1);
//         rollup.finalizeBatch(1);
//         assertTrue(rollup.isBatchFinalized(1));
//         assertEq(rollup.finalizedStateRoots(1), stateRoot);
//         assertTrue(rollup.withdrawalRoots(bytes32(uint256(3))));
//         assertEq(rollup.lastFinalizedBatchIndex(), 1);
//         assertFalse(l1MessageQueueWithGasPriceOracle.isMessageSkipped(0));
//         assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);

//         // commit batch2 with two chunks, correctly
//         // 1. chunk0 has one block, 3 tx, no L1 messages
//         //   => payload for chunk0
//         //    0000000000000000
//         //    0000000000000000
//         //    0000000000000000000000000000000000000000000000000000000000000000
//         //    0000000000000000
//         //    0003
//         //    ... (some tx hashes)
//         //   => data hash for chunk0
//         //    2ac1dad3f3696e5581dfc10f2c7a7a8fc5b344285f7d332c7895a8825fca609a
//         // 2. chunk1 has three blocks
//         //   2.1 block0 has 5 tx, 3 L1 messages, no skips
//         //   2.2 block1 has 10 tx, 5 L1 messages, even is skipped, last is not skipped
//         //   2.2 block1 has 300 tx, 256 L1 messages, odd position is skipped, last is not skipped
//         //   => payload for chunk1
//         //    0000000000000000
//         //    0000000000000000
//         //    0000000000000000000000000000000000000000000000000000000000000000
//         //    0000000000000000
//         //    0005
//         //    0000000000000000
//         //    0000000000000000
//         //    0000000000000000000000000000000000000000000000000000000000000000
//         //    0000000000000000
//         //    000a
//         //    0000000000000000
//         //    0000000000000000
//         //    0000000000000000000000000000000000000000000000000000000000000000
//         //    0000000000000000
//         //    012c
//         //    ... (some tx hashes)
//         //   => data hash for chunk2
//         //    e1276f58354ab2372050bde30d8c970ccc3728c76e97f37deebeee83ecbf5705
//         // => data hash for all chunks
//         //   3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
//         // => payload for batch header
//         //  00
//         //  0000000000000002
//         //  0000000000000108
//         //  0000000000000109
//         //  3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
//         //  cef70bf80683c4d9b8b2813e90c314e8c56648e231300b8cfed9d666b0caf14e
//         //  aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa800000000000000000000000000000000000000000000000000000000000000aa
//         // => hash for batch header
//         //  03a9cdcb9d582251acf60937db006ec99f3505fd4751b7c1f92c9a8ef413e873
//         bytes memory batchHeader2 = new bytes(121 + 32 + 32);
//         assembly {
//             mstore(add(batchHeader2, 0x20), 0) // version
//             mstore(add(batchHeader2, add(0x20, 1)), shl(192, 2)) // batchIndex = 2
//             mstore(add(batchHeader2, add(0x20, 9)), shl(192, 264)) // l1MessagePopped = 264
//             mstore(add(batchHeader2, add(0x20, 17)), shl(192, 265)) // totalL1MessagePopped = 265
//             mstore(
//                 add(batchHeader2, add(0x20, 25)),
//                 0x3c71d155351642d15f1542a1543ce423abeca1f8939100a0a34cdc3127b95f69
//             ) // l1dataHash
//             mstore(
//                 add(batchHeader2, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             ) // l2 tx blob versioned hash
//             mstore(add(batchHeader2, add(0x20, 89)), batchHash1) // parentBatchHash
//             mstore(
//                 add(batchHeader2, add(0x20, 121)),
//                 77194726158210796949047323339125271902179989777093709359638389338608753093160
//             ) // bitmap0
//             mstore(add(batchHeader2, add(0x20, 153)), 42) // bitmap1
//         }
//         chunk0 = new bytes(1 + 60);
//         assembly {
//             mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
//             mstore(add(chunk0, add(0x21, 56)), shl(240, 3)) // numTransactions = 3
//             mstore(add(chunk0, add(0x21, 58)), shl(240, 0)) // numL1Messages = 0
//         }

//         chunk1 = new bytes(1 + 60 * 3);
//         assembly {
//             mstore(add(chunk1, 0x20), shl(248, 3)) // numBlocks = 3
//             mstore(add(chunk1, add(33, 56)), shl(240, 5)) // block0.numTransactions = 5
//             mstore(add(chunk1, add(33, 58)), shl(240, 3)) // block0.numL1Messages = 3
//             mstore(add(chunk1, add(93, 56)), shl(240, 10)) // block1.numTransactions = 10
//             mstore(add(chunk1, add(93, 58)), shl(240, 5)) // block1.numL1Messages = 5
//             mstore(add(chunk1, add(153, 56)), shl(240, 300)) // block1.numTransactions = 300
//             mstore(add(chunk1, add(153, 58)), shl(240, 256)) // block1.numL1Messages = 256
//         }

//         chunks = new bytes[](2);
//         chunks[0] = chunk0;
//         chunks[1] = chunk1;
//         bitmap = new bytes(64);
//         assembly {
//             mstore(
//                 add(bitmap, add(0x20, 0)),
//                 77194726158210796949047323339125271902179989777093709359638389338608753093160
//             ) // bitmap0
//             mstore(add(bitmap, add(0x20, 32)), 42) // bitmap1
//         }

//         // too many txs in one chunk, revert
//         hevm.prank(multisig);
//         rollup.updateMaxNumTxInChunk(2);
//         hevm.mockCall(
//             address(rollup.l1SequencerContract()),
//             abi.encodeCall(IL1Staking.isStaker, ()),
//             abi.encode(true)
//         );
//         hevm.startPrank(address(0));
//         hevm.expectRevert("too many txs in one chunk");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader1,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(4)),
//             nilBatchSig
//         );
//         rollup.commitBatch(batchData, sequencerSigned, signature); // first chunk with too many txs

//         hevm.stopPrank();
//         hevm.prank(multisig);
//         rollup.updateMaxNumTxInChunk(10);
//         hevm.mockCall(
//             address(rollup.l1SequencerContract()),
//             abi.encodeCall(IL1Staking.isStaker, ()),
//             abi.encode(true)
//         );
//         hevm.startPrank(address(0));
//         hevm.expectRevert("too many txs in one chunk");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader1,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(4)),
//             nilBatchSig
//         );
//         rollup.commitBatch(batchData, sequencerSigned, signature); // second chunk with too many txs
//         hevm.stopPrank();

//         hevm.prank(multisig);
//         rollup.updateMaxNumTxInChunk(186);
//         hevm.mockCall(
//             address(rollup.l1SequencerContract()),
//             abi.encodeCall(IL1Staking.isStaker, ()),
//             abi.encode(true)
//         );
//         hevm.startPrank(address(0));
//         hevm.expectEmit(true, true, false, true);
//         emit CommitBatch(
//             2,
//             bytes32(
//                 0xbdc7c412b462dd39875e206312b74dbab37b2e8f3270e7212b36b709b563779d
//             )
//         );

//         // 0x01000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000030000bc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98abc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98abc36789e7a1e281436464229828f817d6612f7b477d66591ff96a9e064bcc98a
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader1,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(5)),
//             nilBatchSig
//         );
//         rollup.commitBatch(batchData, sequencerSigned, signature);
//         hevm.stopPrank();
//         assertFalse(rollup.isBatchFinalized(2));
//         bytes32 batchHash2 = rollup.committedBatches(2);
//         assertEq(
//             batchHash2,
//             bytes32(
//                 0xbdc7c412b462dd39875e206312b74dbab37b2e8f3270e7212b36b709b563779d
//             )
//         );

//         // verify committed batch correctly
//         hevm.startPrank(address(0));
//         hevm.warp(block.timestamp + rollup.finalizationPeriodSeconds());
//         rollup.finalizeBatch(2);
//         hevm.stopPrank();

//         assertTrue(rollup.isBatchFinalized(2));
//         assertEq(rollup.finalizedStateRoots(2), stateRoot);
//         assertTrue(rollup.withdrawalRoots(bytes32(uint256(5))));
//         assertEq(rollup.lastFinalizedBatchIndex(), 2);
//         assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 265);
//         // 1 ~ 4, zero
//         for (uint256 i = 1; i < 4; i++) {
//             assertFalse(l1MessageQueueWithGasPriceOracle.isMessageSkipped(i));
//         }
//         // 4 ~ 9, even is nonzero, odd is zero
//         for (uint256 i = 4; i < 9; i++) {
//             if (i % 2 == 1 || i == 8) {
//                 assertFalse(
//                     l1MessageQueueWithGasPriceOracle.isMessageSkipped(i)
//                 );
//             } else {
//                 assertTrue(
//                     l1MessageQueueWithGasPriceOracle.isMessageSkipped(i)
//                 );
//             }
//         }
//         // 9 ~ 265, even is nonzero, odd is zero
//         for (uint256 i = 9; i < 265; i++) {
//             if (i % 2 == 1 || i == 264) {
//                 assertFalse(
//                     l1MessageQueueWithGasPriceOracle.isMessageSkipped(i)
//                 );
//             } else {
//                 assertTrue(
//                     l1MessageQueueWithGasPriceOracle.isMessageSkipped(i)
//                 );
//             }
//         }
//     }
// }

// contract RollupTest is L1MessageBaseTest {
//     function setUp() public virtual override {
//         super.setUp();
//         nilBatchSig = IRollup.BatchSignature({
//             version: 0,
//             signers: new uint256[](0),
//             signature: "0x123456"
//         });

//         Types.StakerInfo[] memory stakerInfos = new Types.StakerInfo[](
//             SEQUENCER_SIZE
//         );
//         for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
//             address user = address(uint160(beginSeq + i));
//             Types.StakerInfo memory sequencerInfo = ffi.generateStakingInfo(
//                 user
//             );
//             sequencerAddresses.push(sequencerInfo.addr);
//             sequencerBLSKeys.push(sequencerInfo.blsKey);
//             stakerInfos[i] = sequencerInfo;
//         }
//         hevm.startPrank(address(staking));
//         l1Sequencer.updateAndSendSequencerSet(
//             abi.encodeCall(
//                 IL2Staking.updateSequencers,
//                 (l1Sequencer.newestVersion() + 1, stakerInfos)
//             ),
//             sequencerAddresses,
//             sequencerBLSKeys,
//             defaultGasLimit
//         );

//         hevm.stopPrank();
//         assertEq(l1Sequencer.currentVersion(), sequencerVersion);
//         assertEq(
//             l1Sequencer
//                 .getSequencerAddresses(l1Sequencer.currentVersion())
//                 .length,
//             SEQUENCER_SIZE
//         );
//     }

//     function testCommitBatches() external {
//         bytes memory batchHeader0 = new bytes(121);

//         // import 10 L1 messages
//         for (uint256 i = 0; i < 10; i++) {
//             l1CrossDomainMessenger.sendMessage(
//                 address(this),
//                 0,
//                 new bytes(0),
//                 1000000
//             );
//         }

//         // import genesis batch first
//         assembly {
//             mstore(add(batchHeader0, add(0x20, 25)), 1)
//             mstore(
//                 add(batchHeader0, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             )
//         }
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader0, bytes32(uint256(1)));
//         // caller not sequencer, revert
//         hevm.startPrank(address(0));
//         hevm.expectRevert("caller not sequencer");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             new bytes[](0),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // invalid version, revert
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("invalid version");
//         batchData = IRollup.BatchData(
//             1,
//             batchHeader0,
//             new bytes[](0),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // batch is empty, revert
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("batch is empty");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             new bytes[](0),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // batch header length too small, revert
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("batch header length too small");
//         batchData = IRollup.BatchData(
//             0,
//             new bytes(120),
//             new bytes[](1),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // wrong bitmap length, revert
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("wrong bitmap length");
//         batchData = IRollup.BatchData(
//             0,
//             new bytes(122),
//             new bytes[](1),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // incorrect parent batch hash, revert
//         assembly {
//             mstore(add(batchHeader0, add(0x20, 25)), 2) // change data hash for batch0
//         }
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("incorrect parent batch hash");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             new bytes[](1),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();
//         assembly {
//             mstore(add(batchHeader0, add(0x20, 25)), 1) // change back
//             mstore(
//                 add(batchHeader0, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             )
//         }

//         bytes[] memory chunks = new bytes[](1);
//         bytes memory chunk0;
//         // no block in chunk, revert
//         chunk0 = new bytes(1);
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert(ChunkCodecV0.ErrorNoBlockInChunk.selector);
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // invalid chunk length, revert
//         chunk0 = new bytes(1);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert(ChunkCodecV0.ErrorIncorrectChunkLength.selector);
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // cannot skip last L1 message, revert
//         chunk0 = new bytes(1 + 60);
//         bytes memory bitmap = new bytes(32);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunk0[58] = bytes1(uint8(1)); // numTransactions = 1
//         chunk0[60] = bytes1(uint8(1)); // numL1Messages = 1
//         bitmap[31] = bytes1(uint8(1));
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("cannot skip last L1 message");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // num txs less than num L1 msgs, revert
//         chunk0 = new bytes(1 + 60);
//         bitmap = new bytes(32);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunk0[58] = bytes1(uint8(1)); // numTransactions = 1
//         chunk0[60] = bytes1(uint8(3)); // numL1Messages = 3
//         bitmap[31] = bytes1(uint8(3));
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("num txs less than num L1 msgs");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             bitmap,
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // invalid chunk length, revert
//         chunk0 = new bytes(1 + 60 + 1);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert(ChunkCodecV0.ErrorIncorrectChunkLength.selector);
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();

//         // commit batch with one chunk, no tx, correctly
//         chunk0 = new bytes(1 + 60);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         hevm.deal(address(0), 10 ether);
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();
//         assertGt(uint256(rollup.committedBatches(1)), 0);

//         // batch is already committed, revert
//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("batch already committed");
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             getTreeRoot(),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         );
//         hevm.stopPrank();
//     }

//     function testRevertBatch() public {
//         // caller not owner, revert
//         hevm.startPrank(address(1));
//         hevm.expectRevert("Ownable: caller is not the owner");
//         rollup.revertBatch(new bytes(89), 1);
//         hevm.stopPrank();

//         bytes memory batchHeader0 = new bytes(121);

//         // import genesis batch
//         assembly {
//             mstore(add(batchHeader0, add(0x20, 25)), 1)
//             mstore(
//                 add(batchHeader0, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             )
//         }
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader0, bytes32(uint256(1)));
//         bytes32 batchHash0 = rollup.committedBatches(0);

//         bytes[] memory chunks = new bytes[](1);
//         bytes memory chunk0;

//         // commit one batch
//         chunk0 = new bytes(1 + 60);
//         chunk0[0] = bytes1(uint8(1)); // one block in this chunk
//         chunks[0] = chunk0;
//         hevm.startPrank(sequencerAddr);
//         // rollup.stake{value: MIN_DEPOSIT}();
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader0,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(4)),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         ); // first chunk with too many txs
//         hevm.stopPrank();

//         bytes memory batchHeader1 = new bytes(121);
//         assembly {
//             mstore(add(batchHeader1, 0x20), 0) // version
//             mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex
//             mstore(add(batchHeader1, add(0x20, 9)), 0) // l1MessagePopped
//             mstore(add(batchHeader1, add(0x20, 17)), 0) // totalL1MessagePopped
//             mstore(
//                 add(batchHeader1, add(0x20, 25)),
//                 0x246394445f4fe64ed5598554d55d1682d6fb3fe04bf58eb54ef81d1189fafb51
//             ) // l1dataHash
//             mstore(
//                 add(batchHeader1, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             ) // l2 tx blob versioned hash
//             mstore(add(batchHeader1, add(0x20, 89)), batchHash0) // parentBatchHash
//         }

//         // commit another batch
//         hevm.startPrank(sequencerAddr);
//         batchData = IRollup.BatchData(
//             0,
//             batchHeader1,
//             chunks,
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(4)),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         ); // first chunk with too many txs
//         hevm.stopPrank();

//         hevm.startPrank(multisig);
//         // count must be nonzero, revert
//         hevm.expectRevert("count must be nonzero");
//         rollup.revertBatch(batchHeader0, 0);

//         // incorrect batch hash, revert
//         hevm.expectRevert("incorrect batch hash");
//         batchHeader1[0] = bytes1(uint8(1)); // change version to 1
//         rollup.revertBatch(batchHeader1, 1);
//         batchHeader1[0] = bytes1(uint8(0)); // change back

//         // revert middle batch, revert
//         hevm.expectRevert("reverting must start from the ending");
//         rollup.revertBatch(batchHeader1, 1);

//         // can only revert unFinalized batch, revert
//         hevm.expectRevert("can only revert unFinalized batch");
//         rollup.revertBatch(batchHeader0, 3);

//         // succeed to revert next two pending batches.
//         hevm.expectEmit(true, true, false, true);
//         emit RevertBatch(1, rollup.committedBatches(1));
//         hevm.expectEmit(true, true, false, true);
//         emit RevertBatch(2, rollup.committedBatches(2));

//         assertGt(uint256(rollup.committedBatches(1)), 0);
//         assertGt(uint256(rollup.committedBatches(2)), 0);
//         assertEq(uint256(rollup.lastCommittedBatchIndex()), 2);
//         rollup.revertBatch(batchHeader1, 2);
//         assertEq(uint256(rollup.committedBatches(1)), 0);
//         assertEq(uint256(rollup.committedBatches(2)), 0);
//         assertEq(uint256(rollup.lastCommittedBatchIndex()), 0);
//         hevm.stopPrank();
//     }

//     function testSetPause() external {
//         hevm.prank(multisig);
//         rollup.transferOwnership(address(this));

//         // not owner, revert
//         hevm.startPrank(address(1));
//         hevm.expectRevert("Ownable: caller is not the owner");
//         rollup.setPause(false);
//         hevm.stopPrank();

//         // pause
//         rollup.setPause(true);
//         assertBoolEq(true, rollup.paused());

//         hevm.startPrank(sequencerAddr);
//         hevm.expectRevert("Pausable: paused");
//         batchData = IRollup.BatchData(
//             0,
//             new bytes(0),
//             new bytes[](0),
//             new bytes(0),
//             stateRoot,
//             stateRoot,
//             bytes32(uint256(4)),
//             nilBatchSig
//         );
//         rollup.commitBatch(
//             batchData,
//             sequencerVersion,
//             sequencerSigned,
//             signature
//         ); // first chunk with too many txs

//         hevm.expectRevert("Pausable: paused");
//         rollup.finalizeBatch(0);
//         hevm.stopPrank();

//         // unpause
//         rollup.setPause(false);
//         assertBoolEq(false, rollup.paused());
//     }

//     function testUpdateVerifier(address _newVerifier) public {
//         hevm.assume(_newVerifier != address(0));
//         hevm.assume(_newVerifier != rollup.verifier());
//         hevm.prank(multisig);
//         rollup.transferOwnership(address(this));
//         // set by non-owner, should revert
//         hevm.startPrank(address(1));
//         hevm.expectRevert("Ownable: caller is not the owner");
//         rollup.updateVerifier(_newVerifier);
//         hevm.stopPrank();

//         // change to random operator
//         hevm.expectEmit(true, true, false, true);
//         emit UpdateVerifier(address(verifier), _newVerifier);

//         assertEq(rollup.verifier(), address(verifier));
//         rollup.updateVerifier(_newVerifier);
//         assertEq(rollup.verifier(), _newVerifier);
//     }

//     function testUpdateMaxNumTxInChunk(uint256 _maxNumTxInChunk) public {
//         hevm.assume(_maxNumTxInChunk > 0);
//         hevm.assume(_maxNumTxInChunk != 10);

//         hevm.prank(multisig);
//         rollup.transferOwnership(address(this));
//         // set by non-owner, should revert
//         hevm.startPrank(address(1));
//         hevm.expectRevert("Ownable: caller is not the owner");
//         rollup.updateMaxNumTxInChunk(_maxNumTxInChunk);
//         hevm.stopPrank();

//         // change to random operator
//         hevm.expectEmit(false, false, false, true);
//         emit UpdateMaxNumTxInChunk(10, _maxNumTxInChunk);

//         assertEq(rollup.maxNumTxInChunk(), 10);
//         rollup.updateMaxNumTxInChunk(_maxNumTxInChunk);
//         assertEq(rollup.maxNumTxInChunk(), _maxNumTxInChunk);
//     }

//     function testImportGenesisBlock() public {
//         bytes memory batchHeader;

//         // zero state root, revert
//         batchHeader = new bytes(89);
//         hevm.expectRevert("zero state root");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(0));

//         // batch header length too small, revert
//         batchHeader = new bytes(120);
//         hevm.expectRevert("batch header length too small");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // wrong bitmap length, revert
//         batchHeader = new bytes(122);
//         hevm.expectRevert("wrong bitmap length");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // not all fields are zero, revert
//         batchHeader = new bytes(121);
//         batchHeader[0] = bytes1(uint8(1)); // version not zero
//         hevm.expectRevert("not all fields are zero");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         batchHeader = new bytes(121);
//         batchHeader[1] = bytes1(uint8(1)); // batchIndex not zero
//         hevm.expectRevert("not all fields are zero");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         batchHeader = new bytes(121 + 32);
//         assembly {
//             mstore(add(batchHeader, add(0x20, 9)), shl(192, 1)) // l1MessagePopped not zero
//         }
//         hevm.expectRevert("not all fields are zero");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         batchHeader = new bytes(121);
//         batchHeader[17] = bytes1(uint8(1)); // totalL1MessagePopped not zero
//         hevm.expectRevert("not all fields are zero");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // zero data hash, revert
//         batchHeader = new bytes(121);
//         hevm.expectRevert("zero data hash");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // nonzero parent batch hash, revert
//         batchHeader = new bytes(121);
//         batchHeader[25] = bytes1(uint8(1)); // dataHash not zero
//         batchHeader[89] = bytes1(uint8(1)); // parentBatchHash not zero
//         hevm.expectRevert("nonzero parent batch hash");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // invalid versioned hash, revert
//         batchHeader = new bytes(121);
//         batchHeader[25] = bytes1(uint8(1)); // dataHash not zero
//         hevm.expectRevert("invalid versioned hash");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));

//         // import correctly
//         batchHeader = new bytes(121);
//         batchHeader[25] = bytes1(uint8(1)); // dataHash not zero
//         assembly {
//             mstore(
//                 add(batchHeader, add(0x20, 57)),
//                 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
//             ) // ZERO_VERSIONED_HASH
//         }
//         assertEq(rollup.finalizedStateRoots(0), bytes32(0));
//         assertFalse(rollup.withdrawalRoots(0));
//         assertEq(rollup.committedBatches(0), bytes32(0));
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));
//         assertEq(rollup.finalizedStateRoots(0), bytes32(uint256(1)));
//         assertFalse(rollup.withdrawalRoots(0));
//         assertGt(uint256(rollup.committedBatches(0)), 0);

//         // Genesis batch imported, revert
//         hevm.expectRevert("genesis batch imported");
//         hevm.prank(multisig);
//         rollup.importGenesisBatch(batchHeader, bytes32(uint256(1)));
//     }
// }
