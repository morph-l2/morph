// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/// @notice Interface used by Rollup to authorize and slash batch submitters.
interface ISubmitter {
    function isActive(address submitter) external view returns (bool);

    function challengeDeposit() external view returns (uint256);

    function slash(address submitter) external returns (uint256 reward);
}

/// @notice Existing Rollup state used to wait for batches that predate a submitter's exit.
interface IRollupSubmitterView {
    function lastCommittedBatchIndex() external view returns (uint256);

    function lastFinalizedBatchIndex() external view returns (uint256);
}
