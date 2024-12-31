// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IRecord {
    /***********
     * Structs *
     ***********/

    /// @notice BatchSubmission representing a batch submission
    ///
    /// @custom:field index          batch index
    /// @custom:field submitter      batch submitter
    /// @custom:field startBlock     batch start block
    /// @custom:field endBlock       batch end block
    /// @custom:field rollupTime     batch rollup time
    /// @custom:field rollupBlock    batch rollup block number
    struct BatchSubmission {
        uint256 index;
        address submitter;
        uint256 startBlock;
        uint256 endBlock;
        uint256 rollupTime;
        uint256 rollupBlock;
    }

    /***********
     * Errors *
     ***********/

    /// @notice error xxx
    error ErrXXX();

    /**********
     * Events *
     **********/

    /// @notice Emitted batch submissions uploaded
    /// @param startIndex   The data start index
    /// @param dataLength   The data length
    event BatchSubmissionsUploaded(uint256 indexed startIndex, uint256 dataLength);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return next rollup epoch index
    function nextBatchSubmissionIndex() external view returns (uint256);

    /// @notice getBatchSubmissions
    /// @param start start index
    /// @param end   end index
    function getBatchSubmissions(uint256 start, uint256 end) external view returns (BatchSubmission[] memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice set oracle address
    /// @param _oracle   oracle address
    function setOracleAddress(address _oracle) external;

    /// @notice record batch submissions
    function recordFinalizedBatchSubmissions(BatchSubmission[] calldata _batchSubmissions) external;
}
