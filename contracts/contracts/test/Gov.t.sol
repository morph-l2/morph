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
        gov.initialize(multisig, 0, 0, 0, 0, 0, 0);

        // reset initialize
        hevm.store(address(gov), bytes32(uint256(0)), bytes32(uint256(0)));

        hevm.expectRevert("invalid proposal voting duration");
        hevm.prank(multisig);
        gov.initialize(multisig, 0, 0, 0, 0, 0, 0);

        hevm.expectRevert("invalid max chunks");
        hevm.prank(multisig);
        gov.initialize(multisig, 1, 0, 0, 0, 0, 0);

        hevm.expectRevert("invalid rollup epoch");
        hevm.prank(multisig);
        gov.initialize(multisig, 1, 0, 0, 0, 1, 0);

        // _batchBlockInterval
        hevm.expectRevert("invalid batch params");
        hevm.prank(multisig);
        gov.initialize(multisig, 1, 0, 0, 0, 1, 1);
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
        uint256 proposalID = gov.createProposal(proposal);
        uint256 currentProposalID = gov.currentProposalID();
        assertEq(proposalID, nextProposalID);
        assertEq(proposalID, currentProposalID);
        (
            uint256 batchBlockInterval_,
            uint256 batchMaxBytes_,
            uint256 batchTimeout_,
            uint256 maxChunks_,
            uint256 rollupEpoch_
        ) = gov.proposalData(proposalID);
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

        (expirationTime, executed) = gov.proposalInfos(proposalID);
        assertFalse(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
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
        uint256 proposalID = gov.createProposal(proposal);
        hevm.stopPrank();

        uint256 voteCnt = 0;
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));

            voteCnt++;
            if (voteCnt > ((SEQUENCER_SIZE * 2) / 3)) {
                hevm.expectEmit(true, true, true, true);
                emit IGov.ProposalExecuted(proposalID, 0, 0, finalizationPeriodSeconds, MAX_CHUNKS, ROLLUP_EPOCH);
            }
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        uint256 expirationTime;
        bool finished;
        bool passed;
        bool executed;

        (expirationTime, executed) = gov.proposalInfos(proposalID);
        assertTrue(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
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
        uint256 proposalID = gov.createProposal(proposal);
        hevm.stopPrank();

        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        bool finished;
        bool passed;
        bool executed;

        (, executed) = gov.proposalInfos(proposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);

        hevm.prank(address(multisig));
        // decrease sequencer size
        l2Staking.updateSequencerSetMaxSize(SEQUENCER_SIZE - 1);

        (, executed) = gov.proposalInfos(proposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertFalse(finished);
        assertTrue(passed);
        assertFalse(executed);

        gov.executeProposal(proposalID);
        hevm.stopPrank();

        (, executed) = gov.proposalInfos(proposalID);
        assertTrue(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);

        hevm.expectRevert("voting has ended");
        hevm.prank(address(user));
        gov.vote(proposalID);
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
        uint256 proposalID = gov.createProposal(proposal);
        hevm.stopPrank();

        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        bool finished;
        bool passed;
        bool executed;

        (, executed) = gov.proposalInfos(proposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
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

        (, executed) = gov.proposalInfos(proposalID);
        assertFalse(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertFalse(finished);
        assertFalse(passed);
        assertFalse(executed);

        // invalide votes
        for (uint256 i = 0; i < removed.length - 1; i++) {
            assertTrue(gov.isVoted(proposalID, removed[i]));
        }

        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        (, executed) = gov.proposalInfos(proposalID);
        assertTrue(executed);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
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
        uint256 proposalID = gov.createProposal(proposal);
        hevm.stopPrank();

        hevm.warp(block.timestamp + VOTING_DURATION + 1);
        hevm.expectRevert("voting has ended");
        user = address(uint160(beginSeq + 1));
        hevm.prank(address(user));
        gov.vote(proposalID);
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
        uint256 proposalID = gov.createProposal(proposal);
        hevm.stopPrank();

        user = address(uint160(beginSeq + 1));
        bool voted = gov.isVoted(proposalID, user);
        assertFalse(voted);
        hevm.prank(address(user));
        gov.vote(proposalID);
        hevm.stopPrank();
        voted = gov.isVoted(proposalID, user);
        assertTrue(voted);

        hevm.expectRevert("sequencer already voted for this proposal");
        hevm.prank(address(user));
        gov.vote(proposalID);
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
     * @notice createProposal: Reverts if rollup epoch is zero.
     */
    function test_createProposal_zeroRollupEpoch_reverts() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            0 // rollupEpoch
        );

        // Expect revert due to zero rollup epoch.
        hevm.expectRevert("invalid rollup epoch");
        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));
        gov.createProposal(proposal);
        hevm.stopPrank();
    }

    /**
     * @notice createProposal: Reverts if max chunks is zero.
     */
    function test_createProposal_zeroMaxChunks_reverts() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            0, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // Expect revert due to zero max chunks.
        hevm.expectRevert("invalid max chunks");
        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));
        gov.createProposal(proposal);
        hevm.stopPrank();
    }

    /**
     * @notice createProposal: Reverts if batch parameters are zero.
     */
    function test_createProposal_zeroBatchParams_reverts() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            0, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // Expect revert due to zero batch params.
        hevm.expectRevert("invalid batch params");
        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));
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
     * @notice vote: Reverts if proposal ID is invalid.
     */
    function test_vote_invalidProposalID_reverts() external {
        uint256 proposalID = gov.currentProposalID();

        hevm.expectRevert("invalid proposalID");
        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));
        gov.vote(proposalID + 10);
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

    /**
     * @notice delete old data
     */
    function test_deleteOldData_succeeds() external {
        IGov.ProposalData memory proposal0 = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        address user = address(uint160(beginSeq));
        hevm.prank(address(user));
        uint256 proposalID = gov.createProposal(proposal0);
        assertEq(proposalID, 1);
        hevm.stopPrank();

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        uint256 expirationTime;
        bool finished;
        bool passed;
        bool executed;

        (expirationTime, executed) = gov.proposalInfos(proposalID);
        assertTrue(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);

        assertEq(gov.batchBlockInterval(), 0);
        assertEq(gov.batchMaxBytes(), 0);
        assertEq(gov.batchTimeout(), finalizationPeriodSeconds);
        assertEq(gov.maxChunks(), MAX_CHUNKS);
        assertEq(gov.rollupEpoch(), ROLLUP_EPOCH);

        // undeletedProposalStart slot value is 109
        bytes32 undeletedProposalStartBytes32 = hevm.load(address(gov), bytes32(uint256(109)));
        uint256 undeletedProposalStart = uint256(undeletedProposalStartBytes32);

        assertEq(undeletedProposalStart, gov.currentProposalID());

        IGov.ProposalData memory proposal1 = IGov.ProposalData(
            100, // batchBlockInterval
            200, // batchMaxBytes
            finalizationPeriodSeconds, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // create proposal
        hevm.prank(address(user));
        proposalID = gov.createProposal(proposal1);
        assertEq(proposalID, 2);
        hevm.stopPrank();

        // vote
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(proposalID);
            hevm.stopPrank();
            assertTrue(gov.isVoted(proposalID, user));
        }

        (expirationTime, executed) = gov.proposalInfos(proposalID);
        assertTrue(executed);
        assertEq(block.timestamp + VOTING_DURATION, expirationTime);
        (finished, passed, executed) = gov.proposalStatus(proposalID);
        assertTrue(finished);
        assertTrue(passed);
        assertTrue(executed);

        // undeletedProposalStart slot value is 109
        undeletedProposalStartBytes32 = hevm.load(address(gov), bytes32(uint256(109)));
        undeletedProposalStart = uint256(undeletedProposalStartBytes32);

        assertEq(undeletedProposalStart, gov.currentProposalID());

        // test old data
        uint256 preProposalID = 1;
        (expirationTime, executed) = gov.proposalInfos(preProposalID);
        assertEq(expirationTime, 0);
        assertFalse(executed);
        (
            uint256 batchBlockInterval,
            uint256 batchMaxBytes,
            uint256 batchTimeout,
            uint256 maxChunks,
            uint256 rollupEpoch
        ) = gov.proposalData(preProposalID);
        assertEq(batchBlockInterval, 0);
        assertEq(batchMaxBytes, 0);
        assertEq(batchTimeout, 0);
        assertEq(maxChunks, 0);
        assertEq(rollupEpoch, 0);

        assertEq(gov.batchBlockInterval(), 100);
        assertEq(gov.batchMaxBytes(), 200);
        assertEq(gov.batchTimeout(), finalizationPeriodSeconds);
        assertEq(gov.maxChunks(), MAX_CHUNKS);
        assertEq(gov.rollupEpoch(), ROLLUP_EPOCH);
    }
}
