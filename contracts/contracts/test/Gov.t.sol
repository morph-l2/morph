// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Staking} from "../L1/staking/Staking.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {Sequencer} from "../universal/Sequencer.sol";
import {L2Sequencer} from "../L2/L2Sequencer.sol";
import {Types} from "../libraries/Types.sol";
import {Gov} from "../L2/Gov.sol";
import "forge-std/console.sol";
import "./CommonTest.t.sol";

contract L2Gov_Test is Staking_Initializer {
    function setUp() public virtual override {
        super.setUp();

        // set to L2Sequencer
        vm.mockCall(
            address(l2Sequencer.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(l2Sequencer.OTHER_SEQUENCER()))
        );

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerBLSKeys.push(sequencerInfo.blsKey);
            sequencerInfos[i] = sequencerInfo;
        }
        version++;
        vm.prank(address(L2Messenger));
        // updateSequencers
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(l2Sequencer.sequencerAddresses(i), sequencerInfos[i].addr);

            (address user, bytes32 tmKey, bytes memory blsKey) = l2Sequencer
                .sequencerInfos(i);
            assertEq(user, sequencerInfos[i].addr);
            assertEq(tmKey, sequencerInfos[i].tmKey);
            assertEq(blsKey, sequencerInfos[i].blsKey);
        }
    }

    function test_proposal() external {
        Gov.ProposalData memory proposal = Gov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            ROLLUP_EPOCH, // rollupEpoch
            MAX_CHUNKS // maxChunks
        );

        address user = address(uint160(beginSeq));
        vm.prank(address(user));
        l2Gov.propose(proposal);
        (
            uint256 batchBlockInterval_,
            uint256 batchMaxBytes_,
            uint256 batchTimeout_,
            uint256 rollupEpoch_,
            uint256 maxChunks_
        ) = l2Gov.proposalData(version);
        assertEq(batchBlockInterval_, proposal.batchBlockInterval);
        assertEq(batchMaxBytes_, proposal.batchMaxBytes);
        assertEq(batchTimeout_, proposal.batchTimeout);
        assertEq(rollupEpoch_, proposal.rollupEpoch);
        assertEq(maxChunks_, proposal.maxChunks);
        (
            bool active_,
            uint256 endTime_,
            uint256 seqsVersion_,
            uint256 votes_
        ) = l2Gov.proposalInfos(version);
        assertEq(true, active_);
        assertEq(block.timestamp + PROPOSAL_INTERVAL, endTime_);
        assertEq(version, seqsVersion_);
        assertEq(0, votes_);
    }

    function test_vote() external {
        Gov.ProposalData memory proposal = Gov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            ROLLUP_EPOCH, // rollupEpoch
            MAX_CHUNKS // maxChunks
        );

        // proposal
        address user = address(uint160(beginSeq));
        vm.prank(address(user));
        l2Gov.propose(proposal);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            vm.prank(address(user));
            l2Gov.vote(version);
            (, , , uint256 votes_) = l2Gov.proposalInfos(version);
            assertEq(true, l2Gov.votes(version, i));
            assertEq(i + 1, votes_);
        }
    }
}

contract L2Gov_vote_Test is Staking_Initializer {
    Gov.ProposalData public proposal;

    function setUp() public virtual override {
        super.setUp();

        // set to L2Sequencer
        vm.mockCall(
            address(l2Sequencer.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(l2Sequencer.OTHER_SEQUENCER()))
        );

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerBLSKeys.push(sequencerInfo.blsKey);
            sequencerInfos[i] = sequencerInfo;
        }
        version++;
        vm.prank(address(L2Messenger));
        // updateSequencers
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            assertEq(l2Sequencer.sequencerAddresses(i), sequencerInfos[i].addr);

            (address user, bytes32 tmKey, bytes memory blsKey) = l2Sequencer
                .sequencerInfos(i);
            assertEq(user, sequencerInfos[i].addr);
            assertEq(tmKey, sequencerInfos[i].tmKey);
            assertEq(blsKey, sequencerInfos[i].blsKey);
        }
        // proposal version 1
        proposal = Gov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            ROLLUP_EPOCH, // rollupEpoch
            MAX_CHUNKS // maxChunks
        );
        address userBegin = address(uint160(beginSeq));
        vm.prank(address(userBegin));
        l2Gov.propose(proposal);

        // proposal version 2
        version++;
        vm.prank(address(L2Messenger));
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
        userBegin = address(uint160(beginSeq));
        vm.prank(address(userBegin));
        l2Gov.propose(proposal);

        // proposal version 3
        version++;
        vm.prank(address(L2Messenger));
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
    }

    function test_vote_version_endTime() external {
        // check proposal
        uint256 checkVersion = 1;
        uint256 secVersion = 2;

        (
            uint256 batchBlockInterval_,
            uint256 batchMaxBytes_,
            uint256 batchTimeout_,
            uint256 rollupEpoch_,
            uint256 maxChunks_
        ) = l2Gov.proposalData(checkVersion);
        assertEq(batchBlockInterval_, proposal.batchBlockInterval);
        assertEq(batchMaxBytes_, proposal.batchMaxBytes);
        assertEq(batchTimeout_, proposal.batchTimeout);
        assertEq(rollupEpoch_, proposal.rollupEpoch);
        assertEq(maxChunks_, proposal.maxChunks);
        (
            bool active_,
            uint256 endTime_,
            uint256 seqsVersion_,
            uint256 votes_
        ) = l2Gov.proposalInfos(checkVersion);
        assertEq(true, active_);
        assertEq(block.timestamp + PROPOSAL_INTERVAL, endTime_);
        assertEq(checkVersion, seqsVersion_);
        assertEq(0, votes_);

        // revert with proposal inactive
        vm.expectRevert("proposal inactive");
        address user = address(uint160(beginSeq));
        vm.prank(address(user));
        l2Gov.vote(version);

        // revert with version mismatch
        vm.expectRevert("version mismatch");
        vm.prank(address(user));
        l2Gov.vote(secVersion);

        // revert with time end
        user = address(uint160(beginSeq));
        vm.prank(address(user));
        l2Gov.propose(proposal);
        vm.expectRevert("time end");
        vm.warp(block.timestamp + PROPOSAL_INTERVAL + 100);
        vm.prank(address(user));
        l2Gov.vote(version);

        // revert with sequencer already vote for this proposal
        vm.warp(block.timestamp - PROPOSAL_INTERVAL);
        vm.prank(address(user));
        l2Gov.vote(version);
        vm.expectRevert("sequencer already vote for this proposal");
        vm.prank(address(user));
        l2Gov.vote(version);
    }
}
