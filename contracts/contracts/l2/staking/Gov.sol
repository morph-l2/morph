// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";

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
    uint256 public override batchBlockInterval;

    /// @notice deprecated, to delete
    uint256 private batchMaxBytes;

    /// @notice batch timeout
    uint256 public override batchTimeout;

    /// @notice deprecated, to delete
    uint256 private maxChunks;

    /// @notice proposal voting duration
    uint256 public votingDuration;

    /// @notice current proposal ID
    uint256 public override currentProposalID;

    /// @notice the start index of undeleted proposals
    uint256 public undeletedProposalStart;

    /// @notice proposal data
    mapping(uint256 proposalID => ProposalData) public proposalData;

    /// @notice proposal info
    mapping(uint256 proposalID => ProposalInfo) public override proposalInfos;

    /// @notice proposal voter info
    mapping(uint256 proposalID => EnumerableSetUpgradeable.AddressSet) internal votes;

    /// @notice latest executed proposal ID
    uint256 public latestExecutedProposalID;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Ensures that the caller is a sequencer in the sequencer contract.
    modifier onlySequencer() {
        bool _in = ISequencer(SEQUENCER_CONTRACT).isSequencer(_msgSender());
        require(_in, "only sequencer allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    /// @notice Initializer
    /// @param _owner owner
    /// @param _votingDuration proposal interval
    /// @param _batchBlockInterval batch block interval
    /// @param _batchTimeout batch timeout
    function initialize(
        address _owner,
        uint256 _votingDuration,
        uint256 _batchBlockInterval,
        uint256 _batchTimeout
    ) public initializer {
        require(_owner != address(0), "invalid owner address");
        require(_votingDuration > 0, "invalid proposal voting duration");
        require(_batchBlockInterval != 0 || _batchTimeout != 0, "invalid batch params");

        _transferOwnership(_owner);

        votingDuration = _votingDuration;
        batchBlockInterval = _batchBlockInterval;
        batchTimeout = _batchTimeout;

        emit VotingDurationUpdated(0, _votingDuration);
        emit BatchBlockIntervalUpdated(0, _batchBlockInterval);
        emit BatchTimeoutUpdated(0, _batchTimeout);
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice create a proposal
    function createProposal(ProposalData calldata proposal) external onlySequencer returns (uint256) {
        require(proposal.batchBlockInterval != 0 || proposal.batchTimeout != 0, "invalid batch params");

        currentProposalID++;
        proposalData[currentProposalID] = proposal;
        proposalInfos[currentProposalID] = ProposalInfo(block.timestamp + votingDuration, false);

        emit ProposalCreated(currentProposalID, _msgSender(), proposal.batchBlockInterval, proposal.batchTimeout);

        return (currentProposalID);
    }

    /// @notice vote a proposal
    function vote(uint256 proposalID) external onlySequencer {
        require(proposalID <= currentProposalID, "invalid proposalID");
        require(proposalID > latestExecutedProposalID, "expired proposalID");
        require(proposalID >= undeletedProposalStart, "proposal pruned");
        uint256 expirationTime = proposalInfos[proposalID].expirationTime;
        require(
            !(proposalInfos[proposalID].executed || expirationTime == 0 || expirationTime < block.timestamp),
            "voting has ended"
        );
        require(!votes[proposalID].contains(_msgSender()), "sequencer already voted for this proposal");

        votes[proposalID].add(_msgSender());
        if (_checkPassed(proposalID)) {
            _executeProposal(proposalID);
        }
    }

    /// @notice set voting duration
    function setVotingDuration(uint256 _votingDuration) external onlyOwner {
        require(_votingDuration > 0 && _votingDuration != votingDuration, "invalid new proposal voting duration");
        uint256 _oldVotingDuration = votingDuration;
        votingDuration = _votingDuration;
        emit VotingDurationUpdated(_oldVotingDuration, _votingDuration);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice execute a passed proposal
    function executeProposal(uint256 proposalID) external {
        (bool finished, bool passed, ) = proposalStatus(proposalID);
        require(!finished, "voting has ended");
        require(passed, "proposal has not been passed yet");

        _executeProposal(proposalID);
    }

    /// @notice execute a passed proposal
    /// @param deleteTo      last proposal ID to delete
    function cleanUpExpiredProposals(uint256 deleteTo) external {
        require(deleteTo < latestExecutedProposalID, "only allow to delete the proposal befor latest passed proposal");
        // when a proposal is passed, the previous proposals will be invalidated and deleted
        for (uint256 i = undeletedProposalStart; i <= deleteTo; i++) {
            delete proposalData[i];
            delete proposalInfos[i];
            delete votes[i];
        }
        undeletedProposalStart = deleteTo + 1;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return proposal status. {finished, passed, executed}
    function proposalStatus(uint256 proposalID) public view returns (bool, bool, bool) {
        require(proposalID <= currentProposalID, "invalid proposalID");
        require(proposalID >= latestExecutedProposalID, "expired proposal");
        require(proposalID >= undeletedProposalStart, "proposal pruned");

        if (proposalID == latestExecutedProposalID) {
            return (true, true, true);
        }

        bool executed = proposalInfos[proposalID].executed;
        uint256 expirationTime = proposalInfos[proposalID].expirationTime;
        return (
            executed || expirationTime == 0 || expirationTime < block.timestamp,
            _checkPassed(proposalID),
            executed
        );
    }

    /// @notice return whether the address has voted
    /// @param proposalID  proposal ID
    /// @param voter       voter
    function isVoted(uint256 proposalID, address voter) external view returns (bool) {
        return votes[proposalID].contains(voter);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice execute a passed proposal
    function _executeProposal(uint256 proposalID) internal {
        latestExecutedProposalID = proposalID;

        if (batchBlockInterval != proposalData[proposalID].batchBlockInterval) {
            uint256 _oldValue = batchBlockInterval;
            batchBlockInterval = proposalData[proposalID].batchBlockInterval;
            emit BatchBlockIntervalUpdated(_oldValue, proposalData[proposalID].batchBlockInterval);
        }
        if (batchTimeout != proposalData[proposalID].batchTimeout) {
            uint256 _oldValue = batchTimeout;
            batchTimeout = proposalData[proposalID].batchTimeout;
            emit BatchTimeoutUpdated(_oldValue, proposalData[proposalID].batchTimeout);
        }
        proposalInfos[proposalID].executed = true;

        emit ProposalExecuted(proposalID, batchBlockInterval, batchTimeout);
    }

    /// @notice check whether the proposal has been passed
    function _checkPassed(uint256 proposalID) internal view returns (bool) {
        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT).getSequencerSet2();
        uint256 validVotes = 0;
        for (uint256 i = 0; i < latestSequencerSet.length; i++) {
            if (votes[proposalID].contains(latestSequencerSet[i])) {
                validVotes = validVotes + 1;
            }
        }
        return validVotes > (latestSequencerSet.length * 2) / 3;
    }
}
