// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";
import {IRecord} from "./IRecord.sol";

contract Record is IRecord, OwnableUpgradeable {
    // sequencer contract address
    address public immutable SEQUENCER_CONTRACT;
    // gov contract address
    address public immutable GOV_CONTRACT;
    // oracle address
    address public ORACLE;

    // If the sequencer set or rollup epoch changed, reset the rotation

    // mapping(batch_index => batch_submission)
    mapping(uint256 => BatchSubmission) public batchSubmissions;
    // mapping(epoch_index => rollup_epoch_info)
    mapping(uint256 => RollupEpochInfo) public rollupEpochs;
    // latest batch submission index
    uint256 latestBatchSubmissionIndex;
    // latest rollup epoch index
    uint256 latestRollupEpochIndex;

    /*********************** modifiers **************************/

    /// @notice Only prover allowed.
    modifier onlyOracle() {
        require(msg.sender == ORACLE, "only oracle allowed");
        _;
    }

    /*********************** Constructor **************************/

    /**
     * @notice constructor
     */
    constructor() {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        GOV_CONTRACT = Predeploys.GOV;
    }

    /*********************** Init **************************/

    /**
     * @notice Initializer.
     * @param _admin    params admin
     * @param _oracle   oracle address
     */
    function initialize(address _admin, address _oracle) public initializer {
        require(_oracle != address(0), "invalid oracle address");
        ORACLE = _oracle;

        // transfer owner to admin
        _transferOwnership(_admin);
    }

    /*********************** External Functions **************************/

    /**
     * @notice set oracle address
     * @param _oracle     oracle address
     */
    function setOracleAddress(address _oracle) external onlyOwner {
        require(_oracle != address(0), "invalid oracle address");
        ORACLE = _oracle;
    }

    /**
     * @notice record batch submissions
     */
    function recordBatchSubmissions(
        BatchSubmission[] calldata _batchSubmissions
    ) external onlyOracle {
        // TODO
    }

    /**
     * @notice record epochs
     */
    function recordEpochs(
        RollupEpochInfo[] calldata _rollupEpochs
    ) external onlyOracle {
        // TODO
    }

    /*********************** External View Functions **************************/

    /**
     * @notice return epoch index start time and end time
     * @param index     epoch index
     */
    function epochInfo(uint256 index) external returns (uint256, uint256) {
        // TODO
    }

    /**
     * @notice sequencer indicates the proportion of the epoch index in this epoch index
     * This scale is provisionally of the type uint256, which can be divided by 100 in subsequent operations
     * @param epochIndex    epoch index
     * @param sequencer     sequencer address
     */
    function sequencerEpochRatio(
        uint256 epochIndex,
        address sequencer
    ) external returns (uint256) {
        // TODO
    }

    /*********************** Internal Functions **************************/
}
