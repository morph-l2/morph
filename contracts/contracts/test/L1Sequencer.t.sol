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

    // ============ initializeHistory ============

    function test_initializeHistory_success() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        assertEq(l1Sequencer.activeHeight(), UPGRADE_HEIGHT);
        assertEq(l1Sequencer.getSequencerHistoryLength(), 1);
        assertEq(l1Sequencer.getSequencer(), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(UPGRADE_HEIGHT), sequencerA);
    }

    function test_initializeHistory_emitsEvent() public {
        vm.expectEmit(true, true, false, true);
        emit L1Sequencer.SequencerUpdated(address(0), sequencerA, UPGRADE_HEIGHT);

        vm.prank(owner);
        l1Sequencer.initializeHistory(sequencerA, UPGRADE_HEIGHT);
    }

    function test_initializeHistory_revertOnSecondCall() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("already initialized");
        vm.prank(owner);
        l1Sequencer.initializeHistory(sequencerB, UPGRADE_HEIGHT + 100);
    }

    function test_initializeHistory_revertOnZeroAddress() public {
        vm.expectRevert("invalid address");
        vm.prank(owner);
        l1Sequencer.initializeHistory(address(0), UPGRADE_HEIGHT);
    }

    function test_initializeHistory_revertNonOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(nonOwner);
        l1Sequencer.initializeHistory(sequencerA, UPGRADE_HEIGHT);
    }

    // ============ updateSequencer ============

    function test_updateSequencer_success() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT + 100);

        assertEq(l1Sequencer.getSequencerHistoryLength(), 2);
        assertEq(l1Sequencer.getSequencer(), sequencerB);
    }

    function test_updateSequencer_emitsEvent() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectEmit(true, true, false, true);
        emit L1Sequencer.SequencerUpdated(sequencerA, sequencerB, UPGRADE_HEIGHT + 100);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT + 100);
    }

    function test_updateSequencer_revertNotInitialized() public {
        vm.expectRevert("not initialized");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);
    }

    function test_updateSequencer_revertZeroAddress() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("invalid address");
        vm.prank(owner);
        l1Sequencer.updateSequencer(address(0), UPGRADE_HEIGHT + 100);
    }

    function test_updateSequencer_revertStartBlockNotGreater() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("startL2Block must be greater than last record");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT); // equal, not greater
    }

    function test_updateSequencer_revertStartBlockLessThanLast() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("startL2Block must be greater than last record");
        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT - 1);
    }

    function test_updateSequencer_revertNonOwner() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(nonOwner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT + 100);
    }

    // ============ getSequencerAt (binary search) ============

    function test_getSequencerAt_singleRecord_exactHeight() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);
        assertEq(l1Sequencer.getSequencerAt(UPGRADE_HEIGHT), sequencerA);
    }

    function test_getSequencerAt_singleRecord_aboveHeight() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);
        assertEq(l1Sequencer.getSequencerAt(UPGRADE_HEIGHT + 9999), sequencerA);
    }

    function test_getSequencerAt_singleRecord_revertBelowHeight() public {
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.expectRevert("no sequencer at height");
        l1Sequencer.getSequencerAt(UPGRADE_HEIGHT - 1);
    }

    function test_getSequencerAt_multipleRecords() public {
        _initHistory(sequencerA, 100);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerC, 300);

        // Before first record
        vm.expectRevert("no sequencer at height");
        l1Sequencer.getSequencerAt(99);

        // Exact boundaries
        assertEq(l1Sequencer.getSequencerAt(100), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(200), sequencerB);
        assertEq(l1Sequencer.getSequencerAt(300), sequencerC);

        // Between records
        assertEq(l1Sequencer.getSequencerAt(150), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(199), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(250), sequencerB);
        assertEq(l1Sequencer.getSequencerAt(299), sequencerB);

        // After last record
        assertEq(l1Sequencer.getSequencerAt(1000), sequencerC);
    }

    function test_getSequencerAt_twoRecords_boundary() public {
        _initHistory(sequencerA, 100);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 101);

        assertEq(l1Sequencer.getSequencerAt(100), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(101), sequencerB);
    }

    function test_getSequencerAt_manyRecords_binarySearchStress() public {
        _initHistory(sequencerA, 10);

        // Add 9 more records (10 total)
        for (uint64 i = 1; i < 10; i++) {
            address seq = address(uint160(0xA000 + i));
            vm.prank(owner);
            l1Sequencer.updateSequencer(seq, 10 + i * 100);
        }

        assertEq(l1Sequencer.getSequencerHistoryLength(), 10);

        // Query each boundary
        assertEq(l1Sequencer.getSequencerAt(10), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(99), sequencerA);
        assertEq(l1Sequencer.getSequencerAt(110), address(uint160(0xA001)));
        assertEq(l1Sequencer.getSequencerAt(910), address(uint160(0xA009)));
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
        _initHistory(sequencerA, UPGRADE_HEIGHT);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, UPGRADE_HEIGHT + 100);

        assertEq(l1Sequencer.getSequencer(), sequencerB);
    }

    // ============ getSequencerHistory ============

    function test_getSequencerHistory_returnsAll() public {
        _initHistory(sequencerA, 100);

        vm.prank(owner);
        l1Sequencer.updateSequencer(sequencerB, 200);

        L1Sequencer.SequencerRecord[] memory history = l1Sequencer.getSequencerHistory();
        assertEq(history.length, 2);
        assertEq(history[0].startL2Block, 100);
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
        l1Sequencer.initializeHistory(sequencerA, UPGRADE_HEIGHT);
        assertEq(l1Sequencer.getSequencerHistoryLength(), 1);
    }

    function test_renounceOwnership() public {
        vm.prank(owner);
        l1Sequencer.renounceOwnership();
        assertEq(l1Sequencer.owner(), address(0));

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(owner);
        l1Sequencer.initializeHistory(sequencerA, UPGRADE_HEIGHT);
    }
}
