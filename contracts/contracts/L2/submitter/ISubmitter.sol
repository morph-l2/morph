// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

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
     * @notice get the current sequencer's turn
     */
    function getTurn(
        address submitter
    ) external view returns (uint256 startTime, uint256 endTime);

    /**
     * @notice get current submitter
     */
    function getCurrentSubmitter()
        external
        view
        returns (address currentSubmitter, uint256 startTime, uint256 endTime);

    // ============================================================================

    /**
     * @notice set rollup acknowledge
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
}
