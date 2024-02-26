// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title Types
 * @notice Contains various types used throughout the Morph contract system.
 */
library Types {
    /**
     * @notice Struct representing a sequencer information.
     *
     * @custom:field addr   Address of the sequencer.
     * @custom:field tmKey  Tendermint key(ED25519) of the seuqencer.
     * @custom:field blsKey BLS key of the seuqencer.
     */
    struct SequencerInfo {
        address addr;
        bytes32 tmKey;
        bytes blsKey;
    }

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
     * @notice BatchInfo representing a epoch.
     *
     * @custom:field submitter     submitter
     * @custom:field startTime     epoch start time
     * @custom:field endTime       epoch end time
     */
    struct EpochInfo {
        address submitter;
        uint256 startTime;
        uint256 endTime;
    }
}
