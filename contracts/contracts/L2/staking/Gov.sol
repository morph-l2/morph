// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";

contract Gov is IGov, Initializable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

    struct ProposalInfo {
        uint256 endTime;
        bool approved;
    }

    // sequencer contract address
    address public immutable SEQUENCER_CONTRACT;

    // record contract address
    address public immutable RECORD_CONTRACT;

    // batch configs
    uint256 public override batchBlockInterval = 0;
    uint256 public override batchMaxBytes = 0;
    uint256 public override batchTimeout = 0;
    uint256 public override rollupEpoch = 0;
    uint256 public override maxChunks = 0;

    // proposal duration
    uint256 public proposalInterval;
    // proposal id
    uint256 public proposalId = 0;
    // proposal data
    mapping(uint256 => ProposalData) public proposalData;
    // proposal info
    mapping(uint256 => ProposalInfo) public override proposalInfos;
    // proposal voter info
    mapping(uint256 => EnumerableSetUpgradeable.AddressSet) internal votes;

    /*********************** modifiers **************************/
    /**
     * @notice Ensures that the caller is a sequencer in the sequencer contract.
     */
    modifier onlySequencer() {
        bool _in = ISequencer(SEQUENCER_CONTRACT).isSequencer(msg.sender);
        require(_in, "only sequencer can propose");
        _;
    }

    modifier proposalCheck(uint256 _proposalId) {
        require(
            !proposalInfos[_proposalId].approved,
            "proposal already approved"
        );
        require(_proposalActive(_proposalId), "proposal out of date");
        _;
    }

    /*********************** Constructor **************************/
    /**
     * @notice constructor
     */
    constructor() {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
    }

    /*********************** Init **************************/
    /**
     * @notice Initializer.
     * @param _proposalInterval proposal interval
     * @param _batchBlockInterval batch block interval
     * @param _batchMaxBytes max batch bytes
     * @param _batchTimeout batch timeout
     * @param _rollupEpoch rollup epoch
     * @param _maxChunks max chunks
     */
    function initialize(
        uint256 _proposalInterval,
        uint256 _batchBlockInterval,
        uint256 _batchMaxBytes,
        uint256 _batchTimeout,
        uint256 _rollupEpoch,
        uint256 _maxChunks
    ) public initializer {
        require(_proposalInterval > 0, "invalid proposal interval");
        require(_rollupEpoch > 0, "invalid rollup epoch");
        require(_maxChunks > 0, "invalid max chunks");
        require(
            _batchBlockInterval != 0 ||
                _batchMaxBytes != 0 ||
                _batchTimeout != 0,
            "invalid batch params"
        );
        proposalInterval = _proposalInterval;
        batchBlockInterval = _batchBlockInterval;
        batchMaxBytes = _batchMaxBytes;
        batchTimeout = _batchTimeout;
        rollupEpoch = _rollupEpoch;
        maxChunks = _maxChunks;
    }

    /*********************** External Functions **************************/
    /**
     * @notice create a proposal
     */
    function propose(ProposalData memory proposal) external onlySequencer {
        require(proposal.rollupEpoch != 0, "invalid rollup epoch");
        require(proposal.maxChunks > 0, "invalid max chunks");
        require(
            proposal.batchBlockInterval != 0 ||
                proposal.batchMaxBytes != 0 ||
                proposal.batchTimeout != 0,
            "invalid batch params"
        );

        proposalId = proposalId + 1;
        proposalData[proposalId] = proposal;
        proposalInfos[proposalId] = ProposalInfo(
            block.timestamp + proposalInterval, // end time
            false // approved
        );
    }

    /**
     * @notice vote a propsal
     */
    function vote(
        uint256 _proposalId
    ) external onlySequencer proposalCheck(_proposalId) {
        require(
            !votes[_proposalId].contains(msg.sender),
            "sequencer already vote for this proposal"
        );

        // update votes
        votes[_proposalId].add(msg.sender);

        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getLatestSeqeuncerSet();
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (!votes[_proposalId].contains(latestSequencerSet[i])) {
                votes[_proposalId].remove(latestSequencerSet[i]);
            }
        }

        // check votes
        if (votes[_proposalId].length() > (latestSequencerSet.length * 2) / 3) {
            _executeProposal(_proposalId);
        }
    }

    /**
     * @notice execute an approved proposal
     */
    function executeProposal(
        uint256 _proposalId
    ) external proposalCheck(_proposalId) {
        bool approved = _checkProposal(_proposalId);
        if (approved) {
            _executeProposal(_proposalId);
        }
    }

    /**
     * @notice check whether proposal has been approved
     */
    function isProposalApproved(
        uint256 _proposalId
    ) external view returns (bool) {
        return _checkProposal(_proposalId);
    }

    /*********************** Internal Functions **************************/
    /**
     * @notice execute an approved proposal
     */
    function _executeProposal(uint256 _proposalId) internal {
        if (rollupEpoch != proposalData[_proposalId].rollupEpoch) {
            // ISubmitter(SUBMITTER_CONTRACT).epochUpdated(rollupEpoch);
        }
        batchBlockInterval = proposalData[_proposalId].batchBlockInterval;
        batchMaxBytes = proposalData[_proposalId].batchMaxBytes;
        batchTimeout = proposalData[_proposalId].batchTimeout;
        rollupEpoch = proposalData[_proposalId].rollupEpoch;
        maxChunks = proposalData[_proposalId].maxChunks;

        proposalInfos[_proposalId].approved = true;
    }

    /**
     * @notice check whether proposal has been approved
     */
    function _checkProposal(uint256 _proposalId) internal view returns (bool) {
        // already approved
        if (proposalInfos[_proposalId].approved) {
            return true;
        }

        // out of date
        if (!_proposalActive(_proposalId)) {
            return false;
        }

        // checking voided ballots
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getLatestSeqeuncerSet();
        uint256 validVotes = 0;
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (votes[_proposalId].contains(latestSequencerSet[i])) {
                validVotes = validVotes + 1;
            }
        }
        if (validVotes > (latestSequencerSet.length * 2) / 3) {
            return true;
        }
        return false;
    }

    function _proposalActive(uint256 _proposalId) internal view returns (bool) {
        return proposalInfos[_proposalId].endTime >= block.timestamp;
    }
}
