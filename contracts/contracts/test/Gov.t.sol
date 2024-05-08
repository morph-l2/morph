// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Gov} from "../l2/staking/Gov.sol";
import {IGov} from "../l2/staking/IGov.sol";

contract GovTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    function testProposal() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));

        uint256 nextproposalID = gov.currentProposalID() + 1;
        gov.createProposal(proposal);
        (
            uint256 batchBlockInterval_,
            uint256 batchMaxBytes_,
            uint256 batchTimeout_,
            uint256 maxChunks_,
            uint256 rollupEpoch_
        ) = gov.proposalData(nextproposalID);
        assertEq(batchBlockInterval_, proposal.batchBlockInterval);
        assertEq(batchMaxBytes_, proposal.batchMaxBytes);
        assertEq(batchTimeout_, proposal.batchTimeout);
        assertEq(rollupEpoch_, proposal.rollupEpoch);
        assertEq(maxChunks_, proposal.maxChunks);
        (uint256 endTime, bool approved) = gov.proposalInfos(nextproposalID);
        assertFalse(approved);
        assertEq(block.timestamp + PROPOSAL_INTERVAL, endTime);
        hevm.stopPrank();
    }

    function testVote() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        address user = address(uint160(beginSeq));
        hevm.prank(address(user));
        gov.createProposal(proposal);

        uint256 currentproposalID = gov.currentProposalID();
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentproposalID);
            assertTrue(gov.isVoted(currentproposalID, user));
        }

        (, bool approved) = gov.proposalInfos(currentproposalID);
        assertTrue(approved);
    }

    function testCanBeApproved() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        address user = address(uint160(beginSeq));
        hevm.prank(address(user));
        gov.createProposal(proposal);

        uint256 currentproposalID = gov.currentProposalID();
        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentproposalID);
            assertTrue(gov.isVoted(currentproposalID, user));
        }

        (, bool approved) = gov.proposalInfos(currentproposalID);
        assertFalse(approved);

        bool canBeApproved = gov.isProposalCanBeApproved(currentproposalID);
        assertFalse(canBeApproved);

        hevm.prank(address(multisig));
        // decrease sequencer size
        l2Staking.updateSequencerSetMaxSize(SEQUENCER_SIZE - 1);

        canBeApproved = gov.isProposalCanBeApproved(currentproposalID);
        assertTrue(canBeApproved);

        gov.executeProposal(currentproposalID);

        (, approved) = gov.proposalInfos(currentproposalID);
        assertTrue(approved);

        hevm.expectRevert("proposal already approved");
        hevm.prank(address(user));
        gov.vote(currentproposalID);
    }

    function testVoteOutOfDate() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        address user = address(uint160(beginSeq));
        hevm.prank(address(user));
        gov.createProposal(proposal);

        uint256 currentproposalID = gov.currentProposalID();

        hevm.prank(address(user));
        gov.vote(currentproposalID);
        hevm.expectRevert("sequencer already vote for this proposal");
        hevm.prank(address(user));
        gov.vote(currentproposalID);

        user = address(uint160(beginSeq + 1));
        hevm.warp(block.timestamp + PROPOSAL_INTERVAL + 1);
        hevm.expectRevert("proposal out of date");
        hevm.prank(address(user));
        gov.vote(currentproposalID);
    }
}
