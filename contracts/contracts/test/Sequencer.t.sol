// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {ISequencer} from "../l2/staking/ISequencer.sol";

contract SequencerTest is L2StakingBaseTest {
    address public firstStaker;
    address public secondStaker;
    address public thirdStaker;

    function setUp() public virtual override {
        super.setUp();

        firstStaker = address(uint160(beginSeq));
        secondStaker = address(uint160(beginSeq + 1));
        thirdStaker = address(uint160(beginSeq + 2));
    }

    /**
     * @notice initialize: re-initialize
     */
    function test_initialize_paramsCheck_reverts() public {
        address[] memory _sequencersAddresses = new address[](0);

        hevm.expectRevert("Initializable: contract is already initialized");
        hevm.prank(multisig);
        sequencer.initialize(multisig, _sequencersAddresses);

        // reset initialize
        hevm.store(address(sequencer), bytes32(uint256(0)), bytes32(uint256(0)));

        hevm.expectRevert("invalid sequencer set");
        hevm.prank(multisig);
        sequencer.initialize(multisig, _sequencersAddresses);
    }

    /**
     * @notice initialize: Emits event successfully.
     */
    function test_initialize_event_succeeds() public {
        Types.StakerInfo[] memory stakerInfos = new Types.StakerInfo[](SEQUENCER_SIZE);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos[i] = stakerInfo;
            sequencerAddresses.push(stakerInfo.addr);
        }

        // reset initialize
        hevm.store(address(sequencer), bytes32(uint256(0)), bytes32(uint256(0)));

        // expect the SequencerSetUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit ISequencer.SequencerSetUpdated(sequencerAddresses, 0);

        hevm.prank(multisig);
        sequencer.initialize(multisig, sequencerAddresses);
    }

    /**
     * @notice update sequencer
     */
    function test_updateSequencers_succeeds() external {
        address[] memory newSequencers = new address[](SEQUENCER_SIZE);
        beginSeq = 100;
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            sequencerBLSKeys.push(stakerInfo.blsKey);

            newSequencers[i] = stakerInfo.addr;
        }

        // only l2staking contract
        hevm.expectRevert("only L2Staking contract");
        hevm.prank(alice);
        sequencer.updateSequencerSet(newSequencers);

        // expect the SequencerSetUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit ISequencer.SequencerSetUpdated(newSequencers, 3);

        hevm.prank(address(Predeploys.L2_STAKING));
        // updateSequencers
        sequencer.updateSequencerSet(newSequencers);

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(sequencer.getCurrentSequencerSet()[i], sequencerAddresses[i]);
            assertTrue(sequencer.isCurrentSequencer(sequencerAddresses[i]));
        }
        hevm.roll(2);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(sequencer.getCurrentSequencerSet()[i], sequencerAddresses[i]);
            assertTrue(sequencer.isCurrentSequencer(sequencerAddresses[i]));
        }

        hevm.roll(3);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(sequencer.getCurrentSequencerSet()[i], newSequencers[i]);
            assertTrue(sequencer.isCurrentSequencer(newSequencers[i]));
        }
    }

    /**
     * @notice remove sequencer
     */
    function test_sequencerSetAfterRemove_succeeds() external {
        hevm.prank(address(multisig));
        l2Staking.updateSequencerSetMaxSize(SEQUENCER_SIZE - 1);

        hevm.roll(1);
        assertEq(sequencer.getCurrentSequencerSetSize(), SEQUENCER_SIZE);
        assertEq(sequencer.getSequencerSet0Size(), SEQUENCER_SIZE);
        assertEq(sequencer.getSequencerSet1Size(), SEQUENCER_SIZE);
        assertEq(sequencer.getSequencerSet2Size(), SEQUENCER_SIZE - 1);

        assertTrue(sequencer.isSequencer(firstStaker));
        assertTrue(sequencer.isSequencer(secondStaker));
        assertFalse(sequencer.isSequencer(thirdStaker));

        hevm.roll(2);
        assertEq(sequencer.getCurrentSequencerSetSize(), SEQUENCER_SIZE);
        hevm.roll(3);
        assertEq(sequencer.getCurrentSequencerSetSize(), SEQUENCER_SIZE - 1);

        assertEq(sequencer.getSequencerSet0Size(), SEQUENCER_SIZE);
        assertEq(sequencer.getSequencerSet1Size(), SEQUENCER_SIZE);
        assertEq(sequencer.getSequencerSet2Size(), SEQUENCER_SIZE - 1);
    }
}
