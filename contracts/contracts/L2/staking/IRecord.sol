// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IRecord {
    /**
     * @notice BatchSubmission representing a batch submission.
     *
     * @custom:field index          batch index
     * @custom:field submitter      batch submitter
     * @custom:field startBlock     batch start block
     * @custom:field endBlock       batch end block
     * @custom:field rollupTime     batch rollup time
     */
    struct BatchSubmission {
        uint256 index;
        address submitter;
        uint256 startBlock;
        uint256 endBlock;
        uint256 rollupTime;
    }

    /**
     * @notice RollupEpochInfo representing a rollup epoch.
     *
     * @custom:field index         epoch index
     * @custom:field submitter     submitter
     * @custom:field startTime     epoch start time
     * @custom:field endTime       epoch end time
     */
    struct RollupEpochInfo {
        uint256 index;
        address submitter;
        uint256 startTime;
        uint256 endTime;
    }

    /**
     * @notice RewardEpochInfo representing a reward epoch.
     *
     * @custom:field index                  epoch index
     * @custom:field blockCount             the number of blocks included in epoch
     * @custom:field sequencers             sequencers have produced blocks
     * @custom:field sequencerBlocks        number of blocks produced by sequencer
     * @custom:field sequencerRatios        sequencers reward ratio, ten thousandths (ratio/10000)
     * @custom:field sequencerCommissions    sequencers commission percentage
     *
     * If no blocks were produced in this epoch, no sequencer will receive the reward
     */
    struct RewardEpochInfo {
        uint256 index;
        uint256 blockCount;
        address[] sequencers;
        uint256[] sequencerBlocks;
        uint256[] sequencerRatios;
        uint256[] sequencerCommissions;
    }

    /**
     * @notice return next rollup epoch index
     */
    function nextBatchSubmissionIndex() external returns (uint256);

    /**
     * @notice return next rollup epoch index
     */
    function nextRollupEpochIndex() external returns (uint256);

    /**
     * @notice return next reward epoch index
     */
    function nextRewardEpochIndex() external returns (uint256);

    /**
     * @notice return latest reward epoch block
     */
    function latestRewardEpochBlock() external returns (uint256);

    /**
     * @notice getBatchSubmissions
     * @param start start index
     * @param end   end index
     */
    function getBatchSubmissions(
        uint256 start,
        uint256 end
    ) external view returns (BatchSubmission[] memory);

    /**
     * @notice get rollup epochs
     * @param start start index
     * @param end   end index
     */
    function getRollupEpochs(
        uint256 start,
        uint256 end
    ) external view returns (RollupEpochInfo[] memory);

    /**
     * @notice get reward epochs
     * @param start start index
     * @param end   end index
     */
    function getRewardEpochs(
        uint256 start,
        uint256 end
    ) external view returns (RewardEpochInfo[] memory);
}