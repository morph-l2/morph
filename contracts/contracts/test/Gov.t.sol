// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IGov} from "../l2/staking/IGov.sol";
import {Types} from "../libraries/common/Types.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract GovTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    /**
     * @notice initialize: re-initialize
     */
    function test_initialize_paramsCheck_reverts() public {
        hevm.expectRevert("Initializable: contract is already initialized");
        hevm.prank(multisig);
        gov.initialize(0, 0, 0, 0, 0, 0);

        // reset initialize
        hevm.store(address(gov), bytes32(uint256(0)), bytes32(uint256(0)));

        hevm.expectRevert("invalid proposal voting duration");
        hevm.prank(multisig);
        gov.initialize(0, 0, 0, 0, 0, 0);

        hevm.expectRevert("invalid max chunks");
        hevm.prank(multisig);
        gov.initialize(1, 0, 0, 0, 0, 0);

        hevm.expectRevert("invalid rollup epoch");
        hevm.prank(multisig);
        gov.initialize(1, 0, 0, 0, 1, 0);

        // _batchBlockInterval
        hevm.expectRevert("invalid batch params");
        hevm.prank(multisig);
        gov.initialize(1, 0, 0, 0, 1, 1);
    }

    /**
     * @notice create a proposal
     */
    function test_createProposal_succeeds() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        address user = address(uint160(beginSeq));
        uint256 nextProposalID = gov.currentProposalID() + 1;
        hevm.expectEmit(true, true, true, true);
        emit IGov.ProposalCreated(nextProposalID, user, 0, 0, finalizationPeriodSeconds, MAX_CHUNKS, ROLLUP_EPOCH);
        hevm.startPrank(address(user));
        gov.createProposal(proposal);
        (
            uint256 batchBlockInterval_,
            uint256 batchMaxBytes_,
            uint256 batchTimeout_,
            uint256 maxChunks_,
            uint256 rollupEpoch_
        ) = gov.proposalData(nextProposalID);
        hevm.stopPrank();
        assertEq(batchBlockInterval_, proposal.batchBlockInterval);
        assertEq(batchMaxBytes_, proposal.batchMaxBytes);
        assertEq(batchTimeout_, proposal.batchTimeout);
        assertEq(rollupEpoch_, proposal.rollupEpoch);
        assertEq(maxChunks_, proposal.maxChunks);

        uint256 expirationTime;
        bool finished;
        bool passed;
        bool executed;

        (expirationTime, executed) = gov.proposalInfos(nextProposalID);
        assertFalse(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(nextProposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);
    }

    /**
     * @notice vote a proposal
     */
    function test_vote_succeeds() external {
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
        hevm.stopPrank();

        uint256 currentProposalID = gov.currentProposalID();
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentProposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(currentProposalID, user));
        }

        uint256 expirationTime;
        bool finished;
        bool passed;
        bool executed;

        (expirationTime, executed) = gov.proposalInfos(currentProposalID);
        assertTrue(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);
    }

    /**
     * @notice passed by more than 2/3 of the valid votes
     */
    function test_proposalExecute_succeeds() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            1, // batchBlockInterval
            1, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        address user = address(uint160(beginSeq));
        hevm.prank(address(user));
        gov.createProposal(proposal);
        hevm.stopPrank();

        uint256 currentProposalID = gov.currentProposalID();
        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentProposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(currentProposalID, user));
        }

        bool finished;
        bool passed;
        bool executed;

        (, executed) = gov.proposalInfos(currentProposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);

        hevm.prank(address(multisig));
        // decrease sequencer size
        l2Staking.updateSequencerSetMaxSize(SEQUENCER_SIZE - 1);

        (, executed) = gov.proposalInfos(currentProposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertFalse(finished);
        assertTrue(passed);
        assertFalse(executed);

        gov.executeProposal(currentProposalID);
        hevm.stopPrank();

        (, executed) = gov.proposalInfos(currentProposalID);
        assertTrue(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);

        hevm.expectRevert("voting has ended");
        hevm.prank(address(user));
        gov.vote(currentProposalID);
        hevm.stopPrank();

        assertEq(gov.batchBlockInterval(), 1);
        assertEq(gov.batchMaxBytes(), 1);
        assertEq(gov.batchTimeout(), finalizationPeriodSeconds);
        assertEq(gov.maxChunks(), MAX_CHUNKS);
        assertEq(gov.rollupEpoch(), ROLLUP_EPOCH);
    }

    /**
     * @notice passed by more than 2/3 of the valid votes
     * 1. remove all old sequencer which vote a proposal
     * 2. add new sequencer and vote, more thran 2 / 3, the proposal passed
     */
    function test_executeWithNewSequencers_succeeds() external {
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
        hevm.stopPrank();

        uint256 currentProposalID = gov.currentProposalID();
        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentProposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(currentProposalID, user));
        }

        bool finished;
        bool passed;
        bool executed;

        (, executed) = gov.proposalInfos(currentProposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);

        // update new sequencer
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }

        // remove old sequencer
        address[] memory removed = new address[](SEQUENCER_SIZE);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address staker = address(uint160(beginSeq + i));
            removed[i] = staker;
        }
        l2Staking.removeStakers(removed);
        hevm.stopPrank();

        (, executed) = gov.proposalInfos(currentProposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);

        // invalide votes
        for (uint256 i = 0; i < removed.length - 1; i++) {
            assertTrue(gov.isVoted(currentProposalID, removed[i]));
        }

        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentProposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(currentProposalID, user));
        }

        (, executed) = gov.proposalInfos(currentProposalID);
        assertTrue(executed);
        (finished, passed, executed) = gov.proposalStatus(currentProposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);
    }

    /**
     * @notice proposal is finished
     */
    function test_vote_expired_reverts() external {
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
        hevm.stopPrank();

        uint256 currentProposalID = gov.currentProposalID();
        hevm.warp(block.timestamp + VOTING_DURATION + 1);
        hevm.expectRevert("voting has ended");
        user = address(uint160(beginSeq + 1));
        hevm.prank(address(user));
        gov.vote(currentProposalID);
        hevm.stopPrank();
    }

    /**
     * @notice sequencer already voted for this proposal
     */
    function test_vote_repeatVoting_reverts() external {
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
        hevm.stopPrank();

        uint256 currentProposalID = gov.currentProposalID();

        user = address(uint160(beginSeq + 1));
        bool voted = gov.isVoted(currentProposalID, user);
        assertFalse(voted);
        hevm.prank(address(user));
        gov.vote(currentProposalID);
        hevm.stopPrank();
        voted = gov.isVoted(currentProposalID, user);
        assertTrue(voted);

        hevm.expectRevert("sequencer already voted for this proposal");
        hevm.prank(address(user));
        gov.vote(currentProposalID);
        hevm.stopPrank();
    }

    /**
     * @notice only sequenser is allowed to create proposal
     */
    function test_createProposal_onlySequencer_reverts() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        hevm.expectRevert("only sequencer allowed");
        hevm.prank(alice);
        gov.createProposal(proposal);
        hevm.stopPrank();
    }

    /**
     * @notice vote: only sequencer allowed
     */
    function test_vote_onlySequencer_reverts() external {
        uint256 proposalID = gov.currentProposalID();

        hevm.expectRevert("only sequencer allowed");
        hevm.prank(alice);
        gov.vote(proposalID);
        hevm.stopPrank();
    }

    /**
     * @notice setVotingDuration: check params
     */
    function test_setVotingDuration_succeeds() external {
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(alice);
        gov.setVotingDuration(0);
        hevm.stopPrank();

        hevm.expectRevert("invalid new proposal voting duration");
        hevm.prank(multisig);
        gov.setVotingDuration(0);
        hevm.stopPrank();

        uint256 oldVotingDuration = gov.votingDuration();
        hevm.expectRevert("invalid new proposal voting duration");
        hevm.prank(multisig);
        gov.setVotingDuration(oldVotingDuration);
        hevm.stopPrank();

        uint256 newVotingDuration = 100;
        hevm.prank(multisig);
        gov.setVotingDuration(newVotingDuration);
        assertEq(newVotingDuration, gov.votingDuration());
        hevm.stopPrank();
    }
}
