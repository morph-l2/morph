// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";
import {IL2Staking} from "./IL2Staking.sol";

contract Gov is IGov, OwnableUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    /*************
     * Constants *
     *************/

    /// @notice staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// @notice sequencer contract address
    address public immutable SEQUENCER_CONTRACT;

    /*************
     * Variables *
     *************/

    /// @notice batch block interval
    uint256 public override batchBlockInterval = 0;

    /// @notice batch max bytes
    uint256 public override batchMaxBytes = 0;

    /// @notice batch timeout
    uint256 public override batchTimeout = 0;

    /// @notice max chunks
    uint256 public override maxChunks = 0;

    /// @notice rollup epoch
    uint256 public override rollupEpoch = 0;

    /// @notice proposal duration
    uint256 public proposalInterval;

    /// @notice proposal id
    uint256 public override currentProposalID = 0;

    /// @notice proposal data
    mapping(uint256 proposalID => ProposalData) public proposalData;

    /// @notice proposal info
    mapping(uint256 proposalID => ProposalInfo) public override proposalInfos;

    /// @notice proposal voter info
    mapping(uint256 proposalID => EnumerableSetUpgradeable.AddressSet)
        internal votes;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Ensures that the caller is a sequencer in the sequencer contract.
    modifier onlySequencer() {
        bool _in = ISequencer(SEQUENCER_CONTRACT).isSequencer(msg.sender);
        require(_in, "only sequencer can propose");
        _;
    }

    modifier proposalCheck(uint256 proposalID) {
        require(
            !proposalInfos[proposalID].approved,
            "proposal already approved"
        );
        require(_proposalActive(proposalID), "proposal out of date");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
    }

    /***************
     * Initializer *
     ***************/

    /// @notice Initializer
    /// @param _proposalInterval proposal interval
    /// @param _batchBlockInterval batch block interval
    /// @param _batchMaxBytes max batch bytes
    /// @param _batchTimeout batch timeout
    /// @param _maxChunks max chunks
    /// @param _rollupEpoch rollup epoch
    function initialize(
        uint256 _proposalInterval,
        uint256 _batchBlockInterval,
        uint256 _batchMaxBytes,
        uint256 _batchTimeout,
        uint256 _maxChunks,
        uint256 _rollupEpoch
    ) public initializer {
        require(_proposalInterval > 0, "invalid proposal interval");
        require(_maxChunks > 0, "invalid max chunks");
        require(_rollupEpoch > 0, "invalid rollup epoch");
        require(
            _batchBlockInterval != 0 ||
                _batchMaxBytes != 0 ||
                _batchTimeout != 0,
            "invalid batch params"
        );

        __Ownable_init();

        proposalInterval = _proposalInterval;
        batchBlockInterval = _batchBlockInterval;
        batchMaxBytes = _batchMaxBytes;
        batchTimeout = _batchTimeout;
        maxChunks = _maxChunks;
        rollupEpoch = _rollupEpoch;

        emit ProposalExecuted(
            batchBlockInterval,
            batchMaxBytes,
            batchTimeout,
            maxChunks,
            rollupEpoch
        );
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice create a proposal
    function createProposal(
        ProposalData memory proposal
    ) external onlySequencer {
        require(proposal.rollupEpoch != 0, "invalid rollup epoch");
        require(proposal.maxChunks > 0, "invalid max chunks");
        require(
            proposal.batchBlockInterval != 0 ||
                proposal.batchMaxBytes != 0 ||
                proposal.batchTimeout != 0,
            "invalid batch params"
        );

        currentProposalID++;
        proposalData[currentProposalID] = proposal;
        proposalInfos[currentProposalID] = ProposalInfo(
            block.timestamp + proposalInterval, // end time
            false // approved
        );
    }

    /// @notice vote a proposal
    function vote(
        uint256 proposalID
    ) external onlySequencer proposalCheck(proposalID) {
        require(
            !votes[proposalID].contains(msg.sender),
            "sequencer already vote for this proposal"
        );

        // update votes
        votes[proposalID].add(msg.sender);

        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getSequencerSet2();
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (!votes[proposalID].contains(latestSequencerSet[i])) {
                votes[proposalID].remove(latestSequencerSet[i]);
            }
        }

        // check votes
        if (votes[proposalID].length() > (latestSequencerSet.length * 2) / 3) {
            _executeProposal(proposalID);
        }
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice execute an approved proposal
    function executeProposal(
        uint256 proposalID
    ) external proposalCheck(proposalID) {
        bool approved = _checkProposal(proposalID);
        if (approved) {
            _executeProposal(proposalID);
        }
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice whether the proposal can be approved
    function isProposalCanBeApproved(
        uint256 proposalID
    ) external view returns (bool) {
        // already approved
        if (proposalInfos[proposalID].approved) {
            return false;
        }

        // out of date
        if (!_proposalActive(proposalID)) {
            return false;
        }
        return _checkProposal(proposalID);
    }

    /// @notice return is voted
    /// @param proposalID  proposal ID
    /// @param voter       voter
    function isVoted(
        uint256 proposalID,
        address voter
    ) external view returns (bool) {
        return votes[proposalID].contains(voter);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice execute an approved proposal
    function _executeProposal(uint256 proposalID) internal {
        batchBlockInterval = proposalData[proposalID].batchBlockInterval;
        batchMaxBytes = proposalData[proposalID].batchMaxBytes;
        batchTimeout = proposalData[proposalID].batchTimeout;
        maxChunks = proposalData[proposalID].maxChunks;
        rollupEpoch = proposalData[proposalID].rollupEpoch;

        proposalInfos[proposalID].approved = true;

        emit ProposalExecuted(
            batchBlockInterval,
            batchMaxBytes,
            batchTimeout,
            maxChunks,
            rollupEpoch
        );
    }

    /// @notice check whether proposal has been approved
    function _checkProposal(uint256 proposalID) internal view returns (bool) {
        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getSequencerSet2();
        uint256 validVotes = 0;
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (votes[proposalID].contains(latestSequencerSet[i])) {
                validVotes = validVotes + 1;
            }
        }
        return validVotes > (latestSequencerSet.length * 2) / 3;
    }

    /// @notice check active
    function _proposalActive(uint256 proposalID) internal view returns (bool) {
        return proposalInfos[proposalID].endTime >= block.timestamp;
    }
}
