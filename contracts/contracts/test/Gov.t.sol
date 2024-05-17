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

        hevm.expectRevert("invalid proposal interval");
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

    /**
     * @notice approval by more than 2/3 of the valid votes
     */
    function test_proposalCanBeApproved_succeeds() external {
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

        assertEq(gov.batchBlockInterval(), 1);
        assertEq(gov.batchMaxBytes(), 1);
        assertEq(gov.batchTimeout(), finalizationPeriodSeconds);
        assertEq(gov.maxChunks(), MAX_CHUNKS);
        assertEq(gov.rollupEpoch(), ROLLUP_EPOCH);
    }

    /**
     * @notice approval by more than 2/3 of the valid votes
     * 1. remove all old sequencer which vote a proposal
     * 2. add new sequencer and vote, more thran 2 / 3, the proposal should be approved
     */
    function test_canBeApprovedWithNewSequencers_succeeds() external {
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

        canBeApproved = gov.isProposalCanBeApproved(currentproposalID);
        assertFalse(canBeApproved);

        // invalide votes
        for (uint256 i = 0; i < removed.length - 1; i++) {
            assertTrue(gov.isVoted(currentproposalID, removed[i]));
        }

        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            user = address(uint160(beginSeq + i));
            hevm.prank(address(user));
            gov.vote(currentproposalID);
            assertTrue(gov.isVoted(currentproposalID, user));
        }

        (, approved) = gov.proposalInfos(currentproposalID);
        assertTrue(approved);
    }

    /**
     * @notice proposal is outdated
     */
    function test_vote_outOfDate_reverts() external {
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
        hevm.expectRevert("only sequencer can propose");
        hevm.prank(alice);
        gov.createProposal(proposal);
    }

    /**
     * @notice vote: only sequencer allowed
     */
    function test_vote_onlySequencer_reverts() external {
        uint256 proposalID = gov.currentProposalID();

        hevm.expectRevert("only sequencer can propose");
        hevm.prank(alice);
        gov.vote(proposalID);
    }

    /**
     * @notice setProposalInterval: check params
     */
    function test_setProposalInterval_succeeds() external {
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(alice);
        gov.setProposalInterval(0);

        hevm.expectRevert("invalid new proposal interval");
        hevm.prank(multisig);
        gov.setProposalInterval(0);

        uint256 oldProposalInterval = gov.proposalInterval();
        hevm.expectRevert("invalid new proposal interval");
        hevm.prank(multisig);
        gov.setProposalInterval(oldProposalInterval);

        uint256 newProposalInterval = 100;
        hevm.prank(multisig);
        gov.setProposalInterval(newProposalInterval);
        assertEq(newProposalInterval, gov.proposalInterval());
    }
}
