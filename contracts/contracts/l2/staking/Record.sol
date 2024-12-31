// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IRecord} from "./IRecord.sol";

contract Record is IRecord, OwnableUpgradeable {
    /*************
     * Variables *
     *************/

    /// @notice oracle address
    address public oracle;

    /// @notice If the sequencer set or rollup epoch changed, reset the submitter round
    mapping(uint256 batchIndex => BatchSubmission) public batchSubmissions;

    /// @notice next batch submission index
    uint256 public override nextBatchSubmissionIndex;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Only prover allowed.
    modifier onlyOracle() {
        require(_msgSender() == oracle, "only oracle allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    /// @notice Initializer.
    /// @param _owner                       owner
    /// @param _oracle                      oracle address
    /// @param _nextBatchSubmissionIndex    next batch submission index
    function initialize(address _owner, address _oracle, uint256 _nextBatchSubmissionIndex) public initializer {
        require(_owner != address(0), "invalid owner address");
        require(_nextBatchSubmissionIndex != 0, "invalid next batch submission index");
        require(_oracle != address(0), "invalid oracle address");

        _transferOwnership(_owner);

        oracle = _oracle;
        nextBatchSubmissionIndex = _nextBatchSubmissionIndex;
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice set oracle address
    /// @param _oracle     oracle address
    function setOracleAddress(address _oracle) external onlyOwner {
        require(_oracle != address(0), "invalid oracle address");
        oracle = _oracle;
    }

    /// @notice record batch submissions
    function recordFinalizedBatchSubmissions(BatchSubmission[] calldata _batchSubmissions) external onlyOracle {
        require(_batchSubmissions.length > 0, "empty batch submissions");
        for (uint256 i = 0; i < _batchSubmissions.length; i++) {
            require(_batchSubmissions[i].index == nextBatchSubmissionIndex + i, "invalid index");
            batchSubmissions[_batchSubmissions[i].index] = BatchSubmission(
                _batchSubmissions[i].index,
                _batchSubmissions[i].submitter,
                _batchSubmissions[i].startBlock,
                _batchSubmissions[i].endBlock,
                _batchSubmissions[i].rollupTime,
                _batchSubmissions[i].rollupBlock
            );
        }
        emit BatchSubmissionsUploaded(nextBatchSubmissionIndex, _batchSubmissions.length);
        nextBatchSubmissionIndex += _batchSubmissions.length;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice getBatchSubmissions
    /// @param start start index
    /// @param end   end index
    function getBatchSubmissions(uint256 start, uint256 end) external view returns (BatchSubmission[] memory res) {
        require(end >= start, "invalid index");
        res = new BatchSubmission[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = batchSubmissions[i];
        }
    }
}
