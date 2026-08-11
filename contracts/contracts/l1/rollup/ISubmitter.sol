// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/// @notice Runtime interface used by Rollup after the centralized submitter cutover.
interface ISubmitter {
    function rollupContract() external view returns (address);

    function isActive(address submitter) external view returns (bool);

    function challengeDeposit() external view returns (uint256);

    function slash(address submitter) external returns (uint256 reward);
}

/// @notice Bounded liability view used by Submitter before releasing stake.
interface IRollupSubmitterView {
    function pendingBatchCount(address submitter) external view returns (uint256);
}
