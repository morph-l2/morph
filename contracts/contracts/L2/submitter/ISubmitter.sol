// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {Types} from "../../libraries/common/Types.sol";

interface ISubmitter {
    /**
     * @notice next batch index
     */
    function nextBatchIndex() external view returns (uint256);

    /**
     * @notice next batch start block
     */
    function nextBatchStartBlock() external view returns (uint256);

    /**
     * @notice get confirmed batch info
     */
    function getConfirmedBatch(
        uint256 batchIndex
    ) external view returns (Types.BatchInfo memory batchInfo);

    /**
     * @notice get epoch info
     */
    function getEpoch(
        uint256 epochIndex
    ) external view returns (Types.EpochInfo memory epochInfo);

    /**
     * @notice get the current sequencer's turn
     */
    function getTurn(
        address submitter
    ) external view returns (uint256 startTime, uint256 endTime);

    /**
     * @notice get next submitter
     */
    function getNextSubmitter()
        external
        view
        returns (address nextSubmitter, uint256 startTime, uint256 endTime);

    // ============================================================================

    /**
     * @notice set rollup acknowledge, only call by bridge
     */
    function ackRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    ) external;

    /**
     * @notice notify epoch updated
     */
    function epochUpdated(uint256 epoch) external;

    /**
     * @notice notify sequencers updated
     */
    function sequencersUpdated(address[] memory sequencers) external;
}
