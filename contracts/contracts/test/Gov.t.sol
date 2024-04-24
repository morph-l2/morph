// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {Gov} from "../L2/staking/Gov.sol";
import {IGov} from "../L2/staking/IGov.sol";

contract govTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();

        Types.StakerInfo[] memory sequencers = new Types.StakerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakingInfo(
                staker
            );
            sequencers[i] = Types.StakerInfo(
                stakerInfo.addr,
                stakerInfo.tmKey,
                stakerInfo.blsKey
            );
        }

        // hevm.mockCall(
        //     address(l2Staking.messenger()),
        //     abi.encodeWithSelector(
        //         ICrossDomainMessenger.xDomainMessageSender.selector
        //     ),
        //     abi.encode(address(l2Staking.OTHER_STAKING()))
        // );

        // hevm.startPrank(address(l2CrossDomainMessenger));
        // l2Staking.updateStakers(sequencers, true);
        // hevm.stopPrank();
    }

    function testProposal() external {
        IGov.ProposalData memory proposal = IGov.ProposalData(
            0, // batchBlockInterval
            0, // batchMaxBytes
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        address user = address(uint160(beginSeq));
        hevm.startPrank(address(user));

        uint256 nextproposalID = gov.proposalID() + 1;
        gov.propose(proposal);
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
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // proposal
        address user = address(uint160(beginSeq));

        hevm.prank(address(user));
        gov.propose(proposal);

        uint256 currentproposalID = gov.proposalID();
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
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // proposal
        address user = address(uint160(beginSeq));

        hevm.prank(address(user));
        gov.propose(proposal);

        uint256 currentproposalID = gov.proposalID();
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
        l2Staking.updateParams(SEQUENCER_SIZE - 1);

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
            FINALIZATION_PERIOD_SECONDS, // batchTimeout
            MAX_CHUNKS, // maxChunks
            ROLLUP_EPOCH // rollupEpoch
        );

        // proposal
        address user = address(uint160(beginSeq));

        hevm.prank(address(user));
        gov.propose(proposal);

        uint256 currentproposalID = gov.proposalID();

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
