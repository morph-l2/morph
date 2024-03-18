// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISubmitter} from "./ISubmitter.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";

contract Gov is Initializable, IGov {
    struct ProposalData {
        uint256 batchBlockInterval;
        uint256 batchMaxBytes;
        uint256 batchTimeout;
        uint256 rollupEpoch;
        uint256 maxChunks;
    }

    struct ProposalInfo {
        bool active;
        uint256 endTime;
        uint256 seqsVersion;
        uint256 votes;
    }

    /*************
     * Constants *
     *************/

    // sequencers infos
    address public immutable SEQUENCER_CONTRACT;

    // submitter contract address
    address public immutable SUBMITTER_CONTRACT;

    uint256 public sequencersVersion = 0;

    // batch configs
    uint256 public override batchBlockInterval = 0;
    uint256 public override batchMaxBytes = 0;
    uint256 public override batchTimeout = 0;
    uint256 public override rollupEpoch = 0;
    uint256 public override maxChunks = 0;

    // proposals infos
    uint256 public proposalNumbers = 0;
    mapping(uint256 => mapping(uint256 => bool)) public votes;
    mapping(uint256 => ProposalData) public proposalData;
    mapping(uint256 => ProposalInfo) public proposalInfos;
    uint256 public proposalInterval;

    /**
     * @notice Ensures that the caller is a sequencer in the sequencer contract.
     */
    modifier onlySequencer() {
        (bool _in, ) = ISequencer(SEQUENCER_CONTRACT).inSequencersSet(
            false,
            msg.sender
        );
        require(_in, "only sequencer can propose");
        _;
    }

    /**
     * @notice constructor
     */
    constructor() {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        SUBMITTER_CONTRACT = Predeploys.SUBMITTER;
    }

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

    /**
     * @notice submit a proposal
     */
    function propose(ProposalData memory proposal) external onlySequencer {
        require(
            (proposal.batchBlockInterval != 0 ||
                proposal.batchMaxBytes != 0 ||
                proposal.batchTimeout != 0) &&
                proposal.rollupEpoch != 0 &&
                proposal.maxChunks != 0,
            "invalid batch params"
        );
        uint256 version = ISequencer(SEQUENCER_CONTRACT).currentVersion();
        if (version > sequencersVersion) {
            // update version and sequencers
            sequencersVersion = version;
        }

        proposalNumbers++;
        proposalData[proposalNumbers] = proposal;
        proposalInfos[proposalNumbers] = ProposalInfo(
            true, // active
            block.timestamp + proposalInterval, // end time
            version,
            0 // votes
        );
    }

    /**
     * @notice vote a propsal
     */
    function vote(uint256 propID) external onlySequencer {
        require(proposalInfos[propID].active, "proposal inactive");
        (uint256 index, uint256 version) = ISequencer(SEQUENCER_CONTRACT)
            .sequencerIndex(false, msg.sender);
        require(
            !votes[propID][index],
            "sequencer already vote for this proposal"
        );

        // check proposal version and end time
        require(
            proposalInfos[propID].seqsVersion == version,
            "version mismatch"
        );
        require(proposalInfos[propID].endTime >= block.timestamp, "time end");

        // vote
        votes[propID][index] = true;
        proposalInfos[propID].votes += 1;

        (uint256 _length, ) = ISequencer(SEQUENCER_CONTRACT).sequencersLen(
            false
        );
        // check votes
        if (proposalInfos[propID].votes > (_length * 2) / 3) {
            if (rollupEpoch != proposalData[propID].rollupEpoch) {
                ISubmitter(SUBMITTER_CONTRACT).epochUpdated(rollupEpoch);
            }
            batchBlockInterval = proposalData[propID].batchBlockInterval;
            batchMaxBytes = proposalData[propID].batchMaxBytes;
            batchTimeout = proposalData[propID].batchTimeout;
            rollupEpoch = proposalData[propID].rollupEpoch;
            maxChunks = proposalData[propID].maxChunks;
            proposalInfos[propID].active = false;
        }
    }

    function l2Sequencer() public view returns (address) {
        return SEQUENCER_CONTRACT;
    }
}
