// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IRecord {
    /**
     * @notice BatchSubmission representing a batch submission.
     *
     * @custom:field submitter      batch submitter
     * @custom:field startBlock     batch start block
     * @custom:field endBlock       batch end block
     * @custom:field rollupTime     batch rollup time
     * @custom:field maxChunks      max chunks
     */
    struct BatchSubmission {
        address submitter;
        uint256 startBlock;
        uint256 endBlock;
        uint256 rollupTime;
    }

    /**
     * @notice RollupEpochInfo representing a rollup epoch.
     *
     * @custom:field submitter     submitter
     * @custom:field startTime     epoch start time
     * @custom:field endTime       epoch end time
     */
    struct RollupEpochInfo {
        address submitter;
        uint256 startTime;
        uint256 endTime;
    }

    /**
     * @notice return epoch index start time and end time
     * @param index     epoch index
     */
    function epochInfo(
        uint256 index
    ) external returns (uint256 startTime, uint256 endTime);

    /**
     * @notice sequencer indicates the proportion of the epoch index in this epoch index
     *          This scale is provisionally of the type uint256, which can be divided by 100 in subsequent operations
     * @param epochIndex    epoch index
     * @param sequencer     sequencer address
     */
    function sequencerEpochRatio(
        uint256 epochIndex,
        address sequencer
    ) external returns (uint256);
}
