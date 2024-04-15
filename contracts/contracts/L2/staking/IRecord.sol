// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Types} from "../../libraries/common/Types.sol";

interface IRecord {
    /**
     * @notice BatchInfo representing a batch.
     *
     * @custom:field submitter      batch submitter
     * @custom:field startBlock     batch start block
     * @custom:field endBlock       batch end block
     * @custom:field rollupTime     batch rollup time
     * @custom:field maxChunks      max chunks
     */
    struct BatchInfo {
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
     * @notice RollupEpochHistory representing a rollup epoch change.
     *
     * @custom:field epoch      epoch blocks
     * @custom:field timestamp  epoch change time
     */
    struct RollupEpochHistory {
        uint256 blocks;
        uint256 timestamp;
    }
}
