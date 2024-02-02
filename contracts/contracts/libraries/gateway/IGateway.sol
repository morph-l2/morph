// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

interface IGateway {
    /// @notice The address of corresponding L1/L2 Gateway contract.
    function counterpart() external view returns (address);

    /// @notice The address of L1GatewayRouter/L2GatewayRouter contract.
    function router() external view returns (address);

    /// @notice The address of corresponding L1CrossDomainMessenger/L2CrossDomainMessenger contract.
    function messenger() external view returns (address);
}
