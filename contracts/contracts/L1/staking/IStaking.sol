// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IStaking {
    /**
     * @notice staking limit
     */
    function limit() external view returns (uint256);

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        address[] memory sequencers,
        uint32 _minGasLimit
    ) external returns (uint256);
}
