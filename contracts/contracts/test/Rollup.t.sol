// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {BatchCodecV0} from "../libraries/codec/BatchCodecV0.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {IRollup} from "../l1/rollup/IRollup.sol";
import {IL1Staking} from "../l1/staking/IL1Staking.sol";

contract RollupCommitBatchTest is L1MessageBaseTest {
    address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
    IRollup.BatchDataInput public batchDataInput;
    IRollup.BatchSignatureInput public batchSignatureInput;

    function setUp() public virtual override {
        super.setUp();

        batchSignatureInput = IRollup.BatchSignatureInput(
            uint256(0),
            abi.encode(uint256(0), new address[](0), uint256(0), new address[](0), uint256(0), new address[](0)),
            bytes("0x")
        );
    }

    function test_commitAndFinalizeWithL1Messages_succeeds() public {
        bytes32 bytesData0 = bytes32(uint256(0));
        bytes32 bytesData1 = bytes32(uint256(1));
        bytes32 bytesData3 = bytes32(uint256(3));
        bytes32 bytesData4 = bytes32(uint256(4));

        upgradeStorage(address(caller), address(rollup), address(alice));
        hevm.deal(caller, 5 * STAKING_VALUE);
        bytes memory batchHeader0 = new bytes(249);

        hevm.startPrank(caller);
        // import 300 L1 messages
        for (uint256 i = 0; i < 300; i++) {
            l1MessageQueueWithGasPriceOracle.appendCrossDomainMessage(address(caller), 1000000, new bytes(0));
        }
        assertEq(
            l1MessageQueueWithGasPriceOracle.getCrossDomainMessage(0),
            bytes32(0xa2277fd30bbbe74323309023b56035b376d7768ad237ae4fc46ead7dc9591ae1)
        );
        hevm.stopPrank();
        upgradeStorage(address(l1CrossDomainMessenger), address(rollup), address(alice));

        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014)
            mstore(add(batchHeader0, add(0x20, 89)), bytesData0)
            mstore(add(batchHeader0, add(0x20, 121)), bytesData1)
        }

        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader0);
        bytes32 batchHash0 = rollup.committedBatches(0);

        bytes memory batch;

        // commit batch1, one batch with one block, 1 tx, 1 L1 message, no skip
        // => l1 data hash for batch
        //   0x9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110
        // => payload for batch header
        //   00
        //   0000000000000001
        //   0000000000000001
        //   0000000000000001
        //   9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110
        //   010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014
        //   0000000000000000000000000000000000000000000000000000000000000001
        //   0000000000000000000000000000000000000000000000000000000000000002
        //   0000000000000000000000000000000000000000000000000000000000000003
        // => sequencer set verify hash
        //   f1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
        // => hash for parent batch header
        //   00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
        bytes memory batchHeader1 = new bytes(249);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
            mstore(add(batchHeader1, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 25)), 0x9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110) // dataHash
            mstore(add(batchHeader1, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // l2 tx blob versioned hash
            mstore(add(batchHeader1, add(0x20, 89)), bytesData1) // prevStateHash
            mstore(add(batchHeader1, add(0x20, 121)), bytesData1) // postStateHash
            mstore(add(batchHeader1, add(0x20, 153)), bytesData3) // withdrawRootHash
            mstore(
                add(batchHeader1, add(0x20, 185)),
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
            ) // sequencerSetVerifyHash
            mstore(add(batchHeader1, add(0x20, 217)), batchHash0) // parentBatchHash
        }
        batch = new bytes(2 + 60);
        assembly {
            mstore(add(batch, 0x20), shl(240, 1)) // numBlocks = 1
            mstore(add(batch, add(0x22, 56)), shl(240, 1)) // numTransactions = 1
            mstore(add(batch, add(0x22, 58)), shl(240, 1)) // numL1Messages = 1
        }
        hevm.mockCall(
            address(rollup.l1StakingContract()),
            abi.encodeCall(IL1Staking.isActiveStaker, (address(0))),
            abi.encode(true)
        );
        hevm.mockCall(
            address(rollup.l1StakingContract()),
            abi.encodeCall(IL1Staking.getStakerBitmap, (address(0))),
            abi.encode(2)
        );
        hevm.startPrank(address(0));
        hevm.expectEmit(true, true, false, true);
        emit IRollup.CommitBatch(1, bytes32(0x7764891041982c1f3abd964616bf5c777e752e327779f980a0881ef2bc994324));
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, bytesData1, bytesData1, bytesData3);
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        assertFalse(rollup.isBatchFinalized(1));
        bytes32 batchHash1 = rollup.committedBatches(1);
        assertEq(batchHash1, bytes32(0x7764891041982c1f3abd964616bf5c777e752e327779f980a0881ef2bc994324));
        bytes32 stateRoot1 = rollup.committedStateRoots(1);
        assertEq(stateRoot1, bytesData1);

        emit log_bytes32(batchHash0);
        // finalize batch1
        hevm.warp(block.timestamp + rollup.finalizationPeriodSeconds() + 1);
        rollup.finalizeBatch(batchHeader1);
        assertTrue(rollup.isBatchFinalized(1));
        assertEq(rollup.finalizedStateRoots(1), bytesData1);
        assertTrue(rollup.withdrawalRoots(bytes32(uint256(3))));
        assertEq(rollup.lastFinalizedBatchIndex(), 1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        // check deleted values
        assertFalse(rollup.batchExist(0));
        assertEq(rollup.committedStateRoots(0), 0);

        // commit batch2 with 4 blocks, correctly
        // 1. block0 has 3 tx, no L1 messages
        // 2. block1 has 5 tx, 3 L1 messages, no skips
        // 3. block2 has 10 tx, 5 L1 messages, even is skipped, last is not skipped
        // 4. block3 has 300 tx, 256 L1 messages, odd position is skipped, last is not skipped
        bytes memory batchHeader2 = new bytes(249);
        assembly {
            mstore(add(batchHeader2, 0x20), 0) // version
            mstore(add(batchHeader2, add(0x20, 1)), shl(192, 2)) // batchIndex = 2
            mstore(add(batchHeader2, add(0x20, 9)), shl(192, 264)) // l1MessagePopped = 264
            mstore(add(batchHeader2, add(0x20, 17)), shl(192, 265)) // totalL1MessagePopped = 265
            mstore(add(batchHeader2, add(0x20, 25)), 0x81e7e4ffa69f7496f9377a1e4140fbb16fcc81d8086b44a1316fc77cc2b9a63b) // dataHash
            mstore(add(batchHeader2, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // l2 tx blob versioned hash
            mstore(add(batchHeader2, add(0x20, 89)), bytesData1) // prevStateHash
            mstore(add(batchHeader2, add(0x20, 121)), bytesData1) // postStateHash
            mstore(add(batchHeader2, add(0x20, 153)), bytesData4) // withdrawRootHash
            mstore(
                add(batchHeader2, add(0x20, 185)),
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
            ) // sequencerSetVerifyHash
            mstore(add(batchHeader2, add(0x20, 217)), batchHash1) // parentBatchHash
        }
        batch = new bytes(2 + 60 * 4);
        assembly {
            mstore(add(batch, 0x20), shl(240, 4)) // numBlocks = 4
            mstore(add(batch, add(34, 56)), shl(240, 3)) // block0.numTransactions = 3
            mstore(add(batch, add(34, 58)), shl(240, 0)) // block0.numL1Messages = 0
            mstore(add(batch, add(94, 56)), shl(240, 5)) // block1.numTransactions = 5
            mstore(add(batch, add(94, 58)), shl(240, 3)) // block1.numL1Messages = 3
            mstore(add(batch, add(154, 56)), shl(240, 10)) // block2.numTransactions = 10
            mstore(add(batch, add(154, 58)), shl(240, 5)) // block2.numL1Messages = 5
            mstore(add(batch, add(214, 56)), shl(240, 300)) // block3.numTransactions = 300
            mstore(add(batch, add(214, 58)), shl(240, 256)) // block3.numL1Messages = 256
        }

        hevm.prank(multisig);
        hevm.mockCall(
            address(rollup.l1StakingContract()),
            abi.encodeCall(IL1Staking.isActiveStaker, (address(0))),
            abi.encode(true)
        );
        hevm.mockCall(
            address(rollup.l1StakingContract()),
            abi.encodeCall(IL1Staking.getStakerBitmap, (address(0))),
            abi.encode(2)
        );
        hevm.startPrank(address(0));
        hevm.expectEmit(true, true, false, true);
        emit IRollup.CommitBatch(2, bytes32(0x0389812d59b1230ba183f27cbe53955d98a1bce20512048a9b9e062b71403c94));

        batchDataInput = IRollup.BatchDataInput(0, batchHeader1, batch, bytesData1, bytesData1, bytesData4);
        rollup.commitBatch(batchDataInput, batchSignatureInput);

        hevm.stopPrank();
        assertFalse(rollup.isBatchFinalized(2));
        bytes32 batchHash2 = rollup.committedBatches(2);
        assertEq(batchHash2, bytes32(0x0389812d59b1230ba183f27cbe53955d98a1bce20512048a9b9e062b71403c94));
        bytes32 stateRoot2 = rollup.committedStateRoots(2);
        assertEq(stateRoot2, bytesData1);

        // verify committed batch correctly
        hevm.startPrank(address(0));
        hevm.warp(block.timestamp + rollup.finalizationPeriodSeconds());
        rollup.finalizeBatch(batchHeader2);
        hevm.stopPrank();

        // finalize batch2
        assertTrue(rollup.isBatchFinalized(2));
        assertEq(rollup.finalizedStateRoots(2), bytesData1);
        assertTrue(rollup.withdrawalRoots(bytesData4));
        assertEq(rollup.lastFinalizedBatchIndex(), 2);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 265);
        // check deleted values
        assertFalse(rollup.batchExist(1));
        assertEq(rollup.committedStateRoots(1), 0);
    }
}

contract RollupTest is L1MessageBaseTest {
    address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
    bytes32 public stateRoot = bytes32(uint256(1));
    IRollup.BatchDataInput public batchDataInput;
    IRollup.BatchSignatureInput public batchSignatureInput;

    function setUp() public virtual override {
        super.setUp();
        batchSignatureInput = IRollup.BatchSignatureInput(
            uint256(0),
            abi.encode(uint256(0), new address[](0), uint256(0), new address[](0), uint256(0), new address[](0)),
            bytes("0x")
        );

        // add staker

        hevm.deal(alice, 5 * STAKING_VALUE);
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(alice);
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        hevm.stopPrank();
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Registered(alice, stakerInfo.tmKey, stakerInfo.blsKey);
        l1Staking.register{value: STAKING_VALUE}(stakerInfo.tmKey, stakerInfo.blsKey);
        (address addrCheck, bytes32 tmKeyCheck, bytes memory blsKeyCheck) = l1Staking.stakers(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, stakerInfo.tmKey);
        assertBytesEq(blsKeyCheck, stakerInfo.blsKey);
        assertTrue(l1Staking.blsKeys(stakerInfo.blsKey));
        assertTrue(l1Staking.tmKeys(stakerInfo.tmKey));
        hevm.stopPrank();
    }

    function test_commitBatches_succeeds() external {
        bytes memory batchHeader0 = new bytes(249);
        bytes32 bytesData1 = bytes32(uint256(1));
        bytes32 bytesData0 = bytes32(uint256(0));

        // import 10 L1 messages
        for (uint256 i = 0; i < 10; i++) {
            l1CrossDomainMessenger.sendMessage(address(this), 0, new bytes(0), 1000000);
        }

        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014)
            mstore(add(batchHeader0, add(0x20, 121)), bytesData1) // stateRootHsash
            mstore(add(batchHeader0, add(0x20, 217)), bytesData0) // parentBatchHash
        }
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader0);
        // only active staker allowed, revert
        hevm.startPrank(address(0));
        hevm.expectRevert("only active staker allowed");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, new bytes(0), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // invalid version, revert
        hevm.startPrank(alice);
        hevm.expectRevert("invalid version");
        batchDataInput = IRollup.BatchDataInput(1, batchHeader0, new bytes(0), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // batch is empty, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch is empty");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, new bytes(0), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // batch header length incorrect, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch header length too small");
        batchDataInput = IRollup.BatchDataInput(0, new bytes(120), new bytes(1), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // incorrect batch index, revert
        assembly {
            mstore(add(batchHeader0, add(0x20, 1)), shl(192, 1)) // change batch index to 1
        }
        hevm.startPrank(alice);
        hevm.expectRevert("incorrect batch index");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, new bytes(1), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
        assembly {
            mstore(add(batchHeader0, add(0x20, 1)), 0) // change back
        }

        // incorrect parent batch hash, revert
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 2) // change data hash for batch0
        }
        hevm.startPrank(alice);
        hevm.expectRevert("incorrect parent batch hash");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, new bytes(1), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1) // change back
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014)
        }

        // incorrect previous state root, revert
        hevm.startPrank(alice);
        hevm.expectRevert("incorrect previous state root");
        batchDataInput = IRollup.BatchDataInput(
            0,
            batchHeader0,
            new bytes(1),
            bytes32(uint256(2)),
            stateRoot,
            getTreeRoot()
        );
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // incorrect batch header length, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch header length too small");
        batchDataInput = IRollup.BatchDataInput(0, new bytes(248), new bytes(1), stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // no block in batch, revert
        bytes memory batch = new bytes(2);
        hevm.startPrank(alice);
        hevm.expectRevert(BatchCodecV0.ErrorNoBlockInBatch.selector);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // invalid batch length, revert
        batch = new bytes(3);
        batch[1] = bytes1(uint8(1)); // one block in this batch
        hevm.startPrank(alice);
        hevm.expectRevert(BatchCodecV0.ErrorIncorrectBatchLength.selector);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // num txs less than num L1 msgs, revert
        batch = new bytes(2 + 60);
        batch[1] = bytes1(uint8(1)); // one block in this batch
        batch[59] = bytes1(uint8(1)); // numTransactions = 1
        batch[61] = bytes1(uint8(3)); // numL1Messages = 3
        hevm.startPrank(alice);
        hevm.expectRevert("num txs less than num L1 msgs");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // invalid batch length, revert
        batch = new bytes(2 + 60 + 1);
        batch[1] = bytes1(uint8(1)); // one block in this batch
        hevm.startPrank(alice);
        hevm.expectRevert(BatchCodecV0.ErrorIncorrectBatchLength.selector);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // commit batch with one chunk, no tx, correctly
        batch = new bytes(2 + 60);
        batch[1] = bytes1(uint8(1)); // one block in this batch
        hevm.startPrank(alice);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        hevm.deal(address(0), 10 ether);
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
        assertGt(uint256(rollup.committedBatches(1)), 0);

        // batch is already committed, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch already committed");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
    }

    function test_revertBatch_succeeds() public {
        bytes32 bytesData1 = bytes32(uint256(1));
        bytes32 bytesData4 = bytes32(uint256(4));
        // caller not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        rollup.revertBatch(new bytes(89), 1);
        hevm.stopPrank();

        bytes memory batchHeader0 = new bytes(249);

        // import genesis batch
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014)
            mstore(add(batchHeader0, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader0);
        bytes32 batchHash0 = rollup.committedBatches(0);

        // commit one batch
        bytes memory batch = new bytes(2 + 60);
        batch[1] = bytes1(uint8(1)); // one block in this batch
        hevm.startPrank(alice);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, batch, stateRoot, stateRoot, bytes32(uint256(4)));
        rollup.commitBatch(batchDataInput, batchSignatureInput); // first chunk with too many txs
        hevm.stopPrank();
        assertEq(rollup.committedBatches(1), 0xb7cb76cf9e9f5878136c1d14e095f5d5b435fe8252cad6eb100e51110033b6ed);
        bytes memory batchHeader1 = new bytes(249);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex
            mstore(add(batchHeader1, add(0x20, 9)), 0) // l1MessagePopped
            mstore(add(batchHeader1, add(0x20, 17)), 0) // totalL1MessagePopped
            mstore(add(batchHeader1, add(0x20, 25)), 0x7cdb9d7f02ea58dfeb797ed6b4f7ea68846e4f2b0e30ed1535fc98b60c4ec809) // dataHash
            mstore(add(batchHeader1, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // l2 tx blob versioned hash
            mstore(add(batchHeader1, add(0x20, 89)), bytesData1) // prevStateHash
            mstore(add(batchHeader1, add(0x20, 121)), bytesData1) // postStateHash
            mstore(add(batchHeader1, add(0x20, 153)), bytesData4) // withdrawRootHash
            mstore(
                add(batchHeader1, add(0x20, 185)),
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
            ) // sequencerSetVerifyHash
            mstore(add(batchHeader1, add(0x20, 217)), batchHash0) // parentBatchHash
            mstore(add(batchHeader1, add(0x20, 249)), 0) // bitmap0
        }

        // commit another batch
        hevm.startPrank(alice);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader1, batch, stateRoot, stateRoot, bytes32(uint256(4)));

        rollup.commitBatch(batchDataInput, batchSignatureInput); // first chunk with too many txs
        hevm.stopPrank();

        hevm.startPrank(multisig);
        // count must be nonzero, revert
        hevm.expectRevert("count must be nonzero");
        rollup.revertBatch(batchHeader0, 0);

        // incorrect batch hash, revert
        hevm.expectRevert("incorrect batch hash");
        batchHeader1[0] = bytes1(uint8(1)); // change version to 1
        rollup.revertBatch(batchHeader1, 1);
        batchHeader1[0] = bytes1(uint8(0)); // change back

        // revert middle batch, revert
        hevm.expectRevert("reverting must start from the ending");
        rollup.revertBatch(batchHeader1, 1);

        // can only revert unFinalized batch, revert
        hevm.expectRevert("can only revert unFinalized batch");
        rollup.revertBatch(batchHeader0, 3);

        // succeed to revert next two pending batches.
        hevm.expectEmit(true, true, false, true);
        emit IRollup.RevertBatch(1, rollup.committedBatches(1));
        hevm.expectEmit(true, true, false, true);
        emit IRollup.RevertBatch(2, rollup.committedBatches(2));

        assertGt(uint256(rollup.committedBatches(1)), 0);
        assertGt(uint256(rollup.committedBatches(2)), 0);
        assertEq(uint256(rollup.lastCommittedBatchIndex()), 2);
        rollup.revertBatch(batchHeader1, 2);
        assertEq(uint256(rollup.committedBatches(1)), 0);
        assertEq(uint256(rollup.committedBatches(2)), 0);
        assertEq(uint256(rollup.lastCommittedBatchIndex()), 0);
        hevm.stopPrank();
    }

    function test_setPause_onlyOwner_reverts() external {
        hevm.prank(multisig);
        rollup.transferOwnership(address(this));

        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        rollup.setPause(false);
        hevm.stopPrank();

        // pause
        rollup.setPause(true);
        assertBoolEq(true, rollup.paused());

        hevm.startPrank(alice);
        hevm.expectRevert("Pausable: paused");
        batchDataInput = IRollup.BatchDataInput(
            0,
            new bytes(0),
            new bytes(0),
            stateRoot,
            stateRoot,
            bytes32(uint256(4))
        );
        rollup.commitBatch(batchDataInput, batchSignatureInput);

        hevm.expectRevert("Pausable: paused");
        rollup.finalizeBatch(new bytes(0));
        hevm.stopPrank();

        // unpause
        rollup.setPause(false);
        assertBoolEq(false, rollup.paused());
    }

    function test_updateVerifier_succeeds(address _newVerifier) public {
        hevm.assume(_newVerifier != address(0));
        hevm.assume(_newVerifier != rollup.verifier());
        hevm.prank(multisig);
        rollup.transferOwnership(address(this));
        // set by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        rollup.updateVerifier(_newVerifier);
        hevm.stopPrank();

        // change to random operator
        hevm.expectEmit(true, true, false, true);
        emit IRollup.UpdateVerifier(address(verifier), _newVerifier);

        assertEq(rollup.verifier(), address(verifier));
        rollup.updateVerifier(_newVerifier);
        assertEq(rollup.verifier(), _newVerifier);
    }

    function test_importGenesisBlock_succeeds() public {
        bytes memory batchHeader;
        bytes32 bytesData1 = bytes32(uint256(1));
        // invalid batch index, revert
        batchHeader = new bytes(249);
        assembly {
            mstore(add(batchHeader, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
        }
        hevm.expectRevert("invalid batch index");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // zero state root, revert
        batchHeader = new bytes(249);
        hevm.expectRevert("zero state root");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // batch header length incorrect, revert
        batchHeader = new bytes(248);
        assembly {
            mstore(add(batchHeader, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        hevm.expectRevert("batch header length too small");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // not all fields are zero, revert
        batchHeader = new bytes(249);
        assembly {
            mstore(add(batchHeader, add(0x20, 9)), shl(192, 1)) // l1MessagePopped not zero
            mstore(add(batchHeader, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        hevm.expectRevert("l1 message popped should be 0");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // zero data hash, revert
        batchHeader = new bytes(249);
        assembly {
            mstore(add(batchHeader, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        hevm.expectRevert("zero data hash");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // invalid versioned hash, revert
        batchHeader = new bytes(249);
        assembly {
            mstore(add(batchHeader, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        batchHeader[25] = bytes1(uint8(1)); // dataHash not zero
        hevm.expectRevert("invalid versioned hash");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);

        // import correctly
        batchHeader = new bytes(249);
        batchHeader[25] = bytes1(uint8(1)); // dataHash not zero
        assembly {
            mstore(add(batchHeader, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // ZERO_VERSIONED_HASH
            mstore(add(batchHeader, add(0x20, 121)), bytesData1) // stateRootHsash
        }
        assertEq(rollup.finalizedStateRoots(0), bytes32(0));
        assertFalse(rollup.withdrawalRoots(0));
        assertEq(rollup.committedBatches(0), bytes32(0));
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);
        assertEq(rollup.finalizedStateRoots(0), bytes32(uint256(1)));
        assertFalse(rollup.withdrawalRoots(0));
        assertGt(uint256(rollup.committedBatches(0)), 0);

        // Genesis batch imported, revert
        hevm.expectRevert("genesis batch imported");
        hevm.prank(multisig);
        rollup.importGenesisBatch(batchHeader);
    }
}
