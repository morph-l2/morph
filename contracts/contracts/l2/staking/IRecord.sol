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

    /// @notice RollupEpochInfo representing a rollup epoch
    ///
    /// @custom:field index         epoch index
    /// @custom:field submitter     submitter
    /// @custom:field startTime     epoch start time
    /// @custom:field endTime       epoch end time
    /// @custom:field endBlock      epoch end block number
    struct RollupEpochInfo {
        uint256 index;
        address submitter;
        uint256 startTime;
        uint256 endTime;
        uint256 endBlock;
    }

    /// @notice RewardEpochInfo representing a reward epoch.
    ///
    /// @custom:field index                  epoch index
    /// @custom:field blockCount             the number of blocks included in epoch
    /// @custom:field sequencers             sequencers have produced blocks
    /// @custom:field sequencerBlocks        number of blocks produced by sequencer
    /// @custom:field sequencerRatios        sequencers reward ratio, precision is 1e8
    /// @custom:field sequencerCommissions   sequencers commission percentage
    ///
    /// If no blocks were produced in this epoch, no sequencer will receive the reward
    struct RewardEpochInfo {
        uint256 index;
        uint256 blockCount;
        address[] sequencers;
        uint256[] sequencerBlocks;
        uint256[] sequencerRatios;
        uint256[] sequencerCommissions;
    }

    /**********
     * Events *
     **********/

    /// @notice Emitted batch submissions uploaded
    /// @param startIndex   The data start index
    /// @param dataLength   The data length
    event BatchSubmissionsUploaded(uint256 indexed startIndex, uint256 dataLength);

    /// @notice Emitted rollup epochs uploaded
    /// @param startIndex   The data start index
    /// @param dataLength   The data length
    event RollupEpochsUploaded(uint256 indexed startIndex, uint256 dataLength);

    /// @notice Emitted reward epochs uploaded
    /// @param startIndex   The data start index
    /// @param dataLength   The data length
    event RewardEpochsUploaded(uint256 indexed startIndex, uint256 dataLength);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice return next rollup epoch index
    function nextBatchSubmissionIndex() external view returns (uint256);

    /// @notice return next rollup epoch index
    function nextRollupEpochIndex() external view returns (uint256);

    /// @notice return next reward epoch index
    function nextRewardEpochIndex() external view returns (uint256);

    /// @notice return latest reward epoch block
    function latestRewardEpochBlock() external view returns (uint256);

    /// @notice getBatchSubmissions
    /// @param start start index
    /// @param end   end index
    function getBatchSubmissions(uint256 start, uint256 end) external view returns (BatchSubmission[] memory);

    /// @notice get rollup epochs
    /// @param start start index
    /// @param end   end index
    function getRollupEpochs(uint256 start, uint256 end) external view returns (RollupEpochInfo[] memory);

    /// @notice get reward epochs
    /// @param start start index
    /// @param end   end index
    function getRewardEpochs(uint256 start, uint256 end) external view returns (RewardEpochInfo[] memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice set oracle address
    /// @param _oracle   oracle address
    function setOracleAddress(address _oracle) external;

    /// @notice set latest block
    /// @param _latestBlock   latest block
    function setLatestRewardEpochBlock(uint256 _latestBlock) external;

    /// @notice record batch submissions
    function recordFinalizedBatchSubmissions(BatchSubmission[] calldata _batchSubmissions) external;

    /// @notice record epochs
    function recordRollupEpochs(RollupEpochInfo[] calldata _rollupEpochs) external;

    /// @notice record epochs
    function recordRewardEpochs(RewardEpochInfo[] calldata _rewardEpochs) external;
}
