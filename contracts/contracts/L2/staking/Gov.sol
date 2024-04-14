// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";
import {IL2Staking} from "./IL2Staking.sol";

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
    // staking contract address
    address public immutable L2_STAKING_CONTRACT;

    // batch configs
    uint256 public override batchBlockInterval = 0;
    uint256 public override batchMaxBytes = 0;
    uint256 public override batchTimeout = 0;
    uint256 public override rollupEpoch = 0;
    uint256 public override rewardEpoch = 0;
    uint256 public override maxChunks = 0;

    // proposal duration
    uint256 public proposalInterval;
    // proposal id
    uint256 public override proposalID = 0;
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

    modifier proposalCheck(uint256 _proposalID) {
        require(
            !proposalInfos[_proposalID].approved,
            "proposal already approved"
        );
        require(_proposalActive(_proposalID), "proposal out of date");
        _;
    }

    /*********************** Constructor **************************/
    /**
     * @notice constructor
     */
    constructor() {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
    }

    /*********************** Init **************************/
    /**
     * @notice Initializer.
     * @param _proposalInterval proposal interval
     * @param _batchBlockInterval batch block interval
     * @param _batchMaxBytes max batch bytes
     * @param _batchTimeout batch timeout
     * @param _rollupEpoch rollup epoch
     * @param _rewardEpoch rollup epoch
     * @param _maxChunks max chunks
     */
    function initialize(
        uint256 _proposalInterval,
        uint256 _batchBlockInterval,
        uint256 _batchMaxBytes,
        uint256 _batchTimeout,
        uint256 _rollupEpoch,
        uint256 _rewardEpoch,
        uint256 _maxChunks
    ) public initializer {
        require(_proposalInterval > 0, "invalid proposal interval");
        require(_rollupEpoch > 0, "invalid rollup epoch");
        require(_rewardEpoch > 0, "invalid reward epoch");
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
        rewardEpoch = _rewardEpoch;
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

        proposalID = proposalID + 1;
        proposalData[proposalID] = proposal;
        proposalInfos[proposalID] = ProposalInfo(
            block.timestamp + proposalInterval, // end time
            false // approved
        );
    }

    /**
     * @notice vote a propsal
     */
    function vote(
        uint256 _proposalID
    ) external onlySequencer proposalCheck(_proposalID) {
        require(
            !votes[_proposalID].contains(msg.sender),
            "sequencer already vote for this proposal"
        );

        // update votes
        votes[_proposalID].add(msg.sender);

        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getLatestSeqeuncerSet();
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (!votes[_proposalID].contains(latestSequencerSet[i])) {
                votes[_proposalID].remove(latestSequencerSet[i]);
            }
        }

        // check votes
        if (votes[_proposalID].length() > (latestSequencerSet.length * 2) / 3) {
            _executeProposal(_proposalID);
        }
    }

    /**
     * @notice execute an approved proposal
     */
    function executeProposal(
        uint256 _proposalID
    ) external proposalCheck(_proposalID) {
        bool approved = _checkProposal(_proposalID);
        if (approved) {
            _executeProposal(_proposalID);
        }
    }

    /**
     * @notice whether the proposal can be approved
     */
    function isProposalCanBeApproved(
        uint256 _proposalID
    ) external view returns (bool) {
        // already approved
        if (proposalInfos[_proposalID].approved) {
            return false;
        }

        // out of date
        if (!_proposalActive(_proposalID)) {
            return false;
        }
        return _checkProposal(_proposalID);
    }

    /**
     * @custom:field _proposalID
     * @custom:field _voter
     * @return {bool}, check if an account has been voted
     */
    function isVoted(
        uint256 _proposalID,
        address _voter
    ) external view returns (bool) {
        return votes[_proposalID].contains(_voter);
    }

    /*********************** Internal Functions **************************/
    /**
     * @notice execute an approved proposal
     */
    function _executeProposal(uint256 _proposalID) internal {
        if (rollupEpoch != proposalData[_proposalID].rollupEpoch) {
            // IL2Staking(L2_STAKING_CONTRACT).updateParams(_sequenceSize, rollupEpoch);
        }
        batchBlockInterval = proposalData[_proposalID].batchBlockInterval;
        batchMaxBytes = proposalData[_proposalID].batchMaxBytes;
        batchTimeout = proposalData[_proposalID].batchTimeout;
        rollupEpoch = proposalData[_proposalID].rollupEpoch;
        rewardEpoch = proposalData[_proposalID].rewardEpoch;
        maxChunks = proposalData[_proposalID].maxChunks;

        proposalInfos[_proposalID].approved = true;
    }

    /**
     * @notice check whether proposal has been approved
     */
    function _checkProposal(uint256 _proposalID) internal view returns (bool) {
        // checking invalidate votes
        address[] memory latestSequencerSet = ISequencer(SEQUENCER_CONTRACT)
            .getLatestSeqeuncerSet();
        uint256 validVotes = 0;
        for (uint i = 0; i < latestSequencerSet.length; i++) {
            if (votes[_proposalID].contains(latestSequencerSet[i])) {
                validVotes = validVotes + 1;
            }
        }
        if (validVotes > (latestSequencerSet.length * 2) / 3) {
            return true;
        }
        return false;
    }

    function _proposalActive(uint256 _proposalID) internal view returns (bool) {
        return proposalInfos[_proposalID].endTime >= block.timestamp;
    }
}
