// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {IRollup} from "../l1/rollup/IRollup.sol";
import {IL1Staking} from "../l1/staking/IL1Staking.sol";
import {BatchHeaderCodecV0} from "../libraries/codec/BatchHeaderCodecV0.sol";
import {BatchHeaderCodecV1} from "../libraries/codec/BatchHeaderCodecV1.sol";

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

        // commit batch1, verison 0, one batch with one block, 1 tx, 1 L1 message, no skip
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
        bytes memory batchHeader1 = new bytes(BatchHeaderCodecV0.BATCH_HEADER_LENGTH);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
            mstore(add(batchHeader1, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 25)), 0xcf774750b0b5e45500519bc354320a44c42c1e1e4faf0766b1141a7d3e5bc2ea) // dataHash
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
        uint64 lastBlockNumber = 1;
        uint16 numL1Messages = 1;
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
        emit IRollup.CommitBatch(1, bytes32(0xc1862b08d265f073817a8ce0d7cbb426c16d58a86b93464244ab1d027318642e));
        batchDataInput = IRollup.BatchDataInput(
            0,
            batchHeader0,
            lastBlockNumber,
            numL1Messages,
            bytesData1,
            bytesData1,
            bytesData3
        );
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        assertFalse(rollup.isBatchFinalized(1));
        bytes32 batchHash1 = rollup.committedBatches(1);
        assertEq(batchHash1, bytes32(0xc1862b08d265f073817a8ce0d7cbb426c16d58a86b93464244ab1d027318642e));
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

        // commit batch2, verison 1, 4 blocks, correctly
        // 1. block0 has 3 tx, no L1 messages
        // 2. block1 has 5 tx, 3 L1 messages, no skips
        // 3. block2 has 10 tx, 5 L1 messages, even is skipped, last is not skipped
        // 4. block3 has 300 tx, 256 L1 messages, odd position is skipped, last is not skipped
        bytes memory batchHeader2 = new bytes(BatchHeaderCodecV1.BATCH_HEADER_LENGTH);
        lastBlockNumber = 5;
        numL1Messages = 264;
        assembly {
            mstore(add(batchHeader2, 0x20), shl(248, 1)) // version
            mstore(add(batchHeader2, add(0x20, 1)), shl(192, 2)) // batchIndex = 2
            mstore(add(batchHeader2, add(0x20, 9)), shl(192, 264)) // l1MessagePopped = 264
            mstore(add(batchHeader2, add(0x20, 17)), shl(192, 265)) // totalL1MessagePopped = 265
            mstore(add(batchHeader2, add(0x20, 25)), 0x688bd49eddb8f52296974c2a243b10b91b305edd68f46190df21a83c13a6e2ec) // dataHash
            mstore(add(batchHeader2, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014) // l2 tx blob versioned hash
            mstore(add(batchHeader2, add(0x20, 89)), bytesData1) // prevStateHash
            mstore(add(batchHeader2, add(0x20, 121)), bytesData1) // postStateHash
            mstore(add(batchHeader2, add(0x20, 153)), bytesData4) // withdrawRootHash
            mstore(
                add(batchHeader2, add(0x20, 185)),
                0xf1f58308e98844ec99e2990d88bfb36e1a30f0e6591e62af90ae6f8498a1b067
            ) // sequencerSetVerifyHash
            mstore(add(batchHeader2, add(0x20, 217)), batchHash1) // parentBatchHash
            mstore(add(batchHeader2, add(0x20, 249)), shl(192, lastBlockNumber)) // lastBlockNumber
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
        emit IRollup.CommitBatch(2, bytes32(0x772132c2e12f21bfc5f2792838e480830f2c1dd2be0f3207b159905a9f321038));

        batchDataInput = IRollup.BatchDataInput(
            1,
            batchHeader1,
            lastBlockNumber,
            numL1Messages,
            bytesData1,
            bytesData1,
            bytesData4
        );
        rollup.commitBatch(batchDataInput, batchSignatureInput);

        hevm.stopPrank();
        assertFalse(rollup.isBatchFinalized(2));
        bytes32 batchHash2 = rollup.committedBatches(2);
        assertEq(batchHash2, bytes32(0x772132c2e12f21bfc5f2792838e480830f2c1dd2be0f3207b159905a9f321038));
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
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 0, 0, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // invalid version, revert
        hevm.startPrank(alice);
        hevm.expectRevert("invalid version");
        batchDataInput = IRollup.BatchDataInput(2, batchHeader0, 0, 0, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // batch header length incorrect, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch header length too small");
        batchDataInput = IRollup.BatchDataInput(0, new bytes(120), 0, 0, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // incorrect batch index, revert
        assembly {
            mstore(add(batchHeader0, add(0x20, 1)), shl(192, 1)) // change batch index to 1
        }
        hevm.startPrank(alice);
        hevm.expectRevert("incorrect batch index");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 0, 0, stateRoot, stateRoot, getTreeRoot());
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
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 0, 0, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1) // change back
            mstore(add(batchHeader0, add(0x20, 57)), 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014)
        }

        // incorrect previous state root, revert
        hevm.startPrank(alice);
        hevm.expectRevert("incorrect previous state root");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 0, 0, bytes32(uint256(2)), stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // incorrect batch header length, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch header length too small");
        batchDataInput = IRollup.BatchDataInput(0, new bytes(248), 0, 0, stateRoot, stateRoot, getTreeRoot());
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();

        // commit batch with one chunk, no tx, correctly
        hevm.startPrank(alice);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 1, 0, stateRoot, stateRoot, getTreeRoot());
        hevm.deal(address(0), 10 ether);
        rollup.commitBatch(batchDataInput, batchSignatureInput);
        hevm.stopPrank();
        assertGt(uint256(rollup.committedBatches(1)), 0);

        // batch is already committed, revert
        hevm.startPrank(alice);
        hevm.expectRevert("batch already committed");
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 1, 0, stateRoot, stateRoot, getTreeRoot());
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
        hevm.startPrank(alice);
        batchDataInput = IRollup.BatchDataInput(0, batchHeader0, 1, 0, stateRoot, stateRoot, bytes32(uint256(4)));
        rollup.commitBatch(batchDataInput, batchSignatureInput); // first chunk with too many txs
        hevm.stopPrank();
        assertEq(rollup.committedBatches(1), 0x25c3e4fee90e53de960c1092746c431ab570eacf8513011902fa65f10c814541);
        bytes memory batchHeader1 = new bytes(249);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex
            mstore(add(batchHeader1, add(0x20, 9)), 0) // l1MessagePopped
            mstore(add(batchHeader1, add(0x20, 17)), 0) // totalL1MessagePopped
            mstore(add(batchHeader1, add(0x20, 25)), 0xe979da4b80d60a17ce56fa19278c6f3a7e1b43359fb8a8ea46d0264de7d653ab) // dataHash
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
        batchDataInput = IRollup.BatchDataInput(0, batchHeader1, 1, 0, stateRoot, stateRoot, bytes32(uint256(4)));

        rollup.commitBatch(batchDataInput, batchSignatureInput); // first chunk with too many txs
        hevm.stopPrank();

        hevm.startPrank(multisig);
        // count must be nonzero, revert
        hevm.expectRevert("count must be nonzero");
        rollup.revertBatch(batchHeader0, 0);

        // incorrect batch hash, revert
        hevm.expectRevert("incorrect batch hash");
        batchHeader1[1] = bytes1(uint8(1)); // change 2nd byte to 1
        rollup.revertBatch(batchHeader1, 1);
        batchHeader1[1] = bytes1(uint8(0)); // change back

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
        batchDataInput = IRollup.BatchDataInput(0, new bytes(0), 0, 0, stateRoot, stateRoot, bytes32(uint256(4)));
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
