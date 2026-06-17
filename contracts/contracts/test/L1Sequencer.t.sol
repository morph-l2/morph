// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L1SequencerBaseTest} from "./base/L1SequencerBase.t.sol";
import {L1Sequencer} from "../l1/L1Sequencer.sol";

contract L1SequencerTest is L1SequencerBaseTest {
    // ============ initialize ============

    function test_initialize_setsOwner() public {
        assertEq(l1Sequencer.owner(), owner);
    }

    function test_initialize_revertOnReinit() public {
        vm.expectRevert("Initializable: contract is already initialized");
        l1Sequencer.initialize(owner);
    }

    function test_initialize_revertOnZeroOwner() public {
        L1Sequencer impl = new L1Sequencer();
        vm.expectRevert("invalid owner");
        impl.initialize(address(0));
    }

    // ============ setFirstSequencer ============
    // The first sequencer is always registered at L2 block 0 (identity from
    // genesis); the upgrade height is no longer stored on-chain.

    function test_setFirstSequencer_success() public {
        _setFirstSequencer(sequencerA);

        assertEq(l1Sequencer.getSequencerHistoryLength(), 1);
        assertEq(l1Sequencer.getSequencer(), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(0), sequencerA);

        L1Sequencer.HistoryRecord[] memory history = l1Sequencer.getSequencerHistory();
        assertEq(history[0].startL2Block, 0);
        assertEq(history[0].sequencerAddr, sequencerA);
    }

    function test_setFirstSequencer_emitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit L1Sequencer.SequencerUpdated(address(0), sequencerA, 0);

        vm.prank(owner);
        l1Sequencer.setFirstSequencer(sequencerA);
    }

    function test_setFirstSequencer_revertOnSecondCall() public {
        _setFirstSequencer(sequencerA);

        vm.expectRevert("already initialized");
        vm.prank(owner);
        l1Sequencer.setFirstSequencer(sequencerB);
    }

    function test_setFirstSequencer_revertOnZeroAddress() public {
        vm.expectRevert("invalid address");
        vm.prank(owner);
        l1Sequencer.setFirstSequencer(address(0));
    }

    function test_setFirstSequencer_revertNonOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(nonOwner);
        l1Sequencer.setFirstSequencer(sequencerA);
    }

    // ============ updateSequencer ============

    function test_updateSequencer_success() public {
        _setFirstSequencer(sequencerA);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT);

        assertEq(l1Sequencer.getSequencerHistoryLength(), 2);
        assertEq(l1Sequencer.getSequencer(), sequencerB);
    }

    function test_updateSequencer_emitsEvent() public {
        _setFirstSequencer(sequencerA);

        vm.expectEmit(true, true, false, true);
        emit L1Sequencer.SequencerUpdated(sequencerA, sequencerB, UPGRADE_HEIGHT);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT);
    }

    function test_updateSequencer_revertNotInitialized() public {
        vm.expectRevert("not initialized");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);
    }

    function test_updateSequencer_revertZeroAddress() public {
        _setFirstSequencer(sequencerA);

        vm.expectRevert("invalid address");
        vm.prank(owner);
        l1Sequencer.updateSequencer(address(0), UPGRADE_HEIGHT);
    }

    function test_updateSequencer_revertStartBlockNotGreater() public {
        _setFirstSequencer(sequencerA); // first record at block 0

        vm.expectRevert("startL2Block must be greater than last record");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 0); // equal to last (0), not greater
    }

    function test_updateSequencer_revertStartBlockLessThanLast() public {
        _setFirstSequencer(sequencerA);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT); // last = 100

        vm.expectRevert("startL2Block must be greater than last record");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerC, UPGRADE_HEIGHT - 1); // 99 < 100
    }

    function test_updateSequencer_revertNonOwner() public {
        _setFirstSequencer(sequencerA);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(nonOwner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT);
    }

    // ============ getSequencerAt (binary search) ============

    function test_getSequencerAt_singleRecord_atZero() public {
        _setFirstSequencer(sequencerA);
        assertEq(l1Sequencer.getSequencerAt(0), sequencerA);
    }

    function test_getSequencerAt_singleRecord_aboveZero() public {
        _setFirstSequencer(sequencerA);
        assertEq(l1Sequencer.getSequencerAt(9999), sequencerA);
    }

    function test_getSequencerAt_multipleRecords() public {
        _setFirstSequencer(sequencerA); // start 0

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerC, 300);

        // First record covers from block 0
        assertEq(l1Sequencer.getSequencerAt(0), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(199), sequencerA);

        // Exact boundaries
        assertEq(l1Sequencer.getSequencerAt(200), sequencerB);
        assertEq(l1Sequencer.getSequencerAt(300), sequencerC);

        // Between records
        assertEq(l1Sequencer.getSequencerAt(250), sequencerB);
        assertEq(l1Sequencer.getSequencerAt(299), sequencerB);

        // After last record
        assertEq(l1Sequencer.getSequencerAt(1000), sequencerC);
    }

    function test_getSequencerAt_twoRecords_boundary() public {
        _setFirstSequencer(sequencerA); // start 0

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 101);

        assertEq(l1Sequencer.getSequencerAt(0), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(100), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(101), sequencerB);
    }

    function test_getSequencerAt_manyRecords_binarySearchStress() public {
        _setFirstSequencer(sequencerA); // start 0

        // Add 9 more records (10 total) at 100, 200, ... 900
        for (uint64 i = 1; i < 10; i++) {
            address seq = address(uint160(0xA000 + i));
            vm.prank(owner);
            l1Sequencer.updateSequencer(seq, i * 100);
        }

        assertEq(l1Sequencer.getSequencerHistoryLength(), 10);

        // Query each boundary
        assertEq(l1Sequencer.getSequencerAt(0), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(99), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(100), address(uint160(0xA001)));
        assertEq(l1Sequencer.getSequencerAt(900), address(uint160(0xA009)));
        assertEq(l1Sequencer.getSequencerAt(99999), address(uint160(0xA009)));
    }

    function test_getSequencerAt_revertEmptyHistory() public {
        vm.expectRevert("no sequencer configured");
        l1Sequencer.getSequencerAt(100);
    }

    // ============ getSequencer ============

    function test_getSequencer_revertEmpty() public {
        vm.expectRevert("no sequencer configured");
        l1Sequencer.getSequencer();
    }

    function test_getSequencer_returnsLatest() public {
        _setFirstSequencer(sequencerA);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT);

        assertEq(l1Sequencer.getSequencer(), sequencerB);
    }

    // ============ getSequencerHistory ============

    function test_getSequencerHistory_returnsAll() public {
        _setFirstSequencer(sequencerA); // start 0

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);

        L1Sequencer.HistoryRecord[] memory history = l1Sequencer.getSequencerHistory();
        assertEq(history.length, 2);
        assertEq(history[0].startL2Block, 0);
        assertEq(history[0].sequencerAddr, sequencerA);
        assertEq(history[1].startL2Block, 200);
        assertEq(history[1].sequencerAddr, sequencerB);
    }

    // ============ ownership ============

    function test_transferOwnership() public {
        vm.prank(owner);
        l1Sequencer.transferOwnership(nonOwner);
        assertEq(l1Sequencer.owner(), nonOwner);

        // New owner can now call admin functions
        vm.prank(nonOwner);
        l1Sequencer.setFirstSequencer(sequencerA);
        assertEq(l1Sequencer.getSequencerHistoryLength(), 1);
    }

    function test_renounceOwnership() public {
        vm.prank(owner);
        l1Sequencer.renounceOwnership();
        assertEq(l1Sequencer.owner(), address(0));

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(owner);
        l1Sequencer.setFirstSequencer(sequencerA);
    }
}
