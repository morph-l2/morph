// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

interface IGov {
    /**
     * @notice batch block interval
     */
    function batchBlockInterval() external view returns (uint256);

    /**
     * @notice batch max bytes
     */
    function batchMaxBytes() external view returns (uint256);

    /**
     * @notice next batch index
     */
    function batchTimeout() external view returns (uint256);

    /**
     * @notice next batch index
     */
    function rollupEpoch() external view returns (uint256);

    /**
     * @notice max chunks
     */
    function maxChunks() external view returns (uint256);
}
