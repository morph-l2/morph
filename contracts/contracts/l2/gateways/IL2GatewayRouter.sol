// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL2ETHGateway} from "./IL2ETHGateway.sol";
import {IL2ERC20Gateway} from "./IL2ERC20Gateway.sol";

interface IL2GatewayRouter is IL2ETHGateway, IL2ERC20Gateway {
    /**********
     * Events *
     **********/

    /// @notice Emitted when the address of ETH Gateway is updated.
    /// @param oldETHGateway The address of the old ETH Gateway.
    /// @param newEthGateway The address of the new ETH Gateway.
    event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway);

    /// @notice Emitted when the address of default ERC20 Gateway is updated.
    /// @param oldDefaultERC20Gateway The address of the old default ERC20 Gateway.
    /// @param newDefaultERC20Gateway The address of the new default ERC20 Gateway.
    event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway);

    /// @notice Emitted when the `gateway` for `token` is updated.
    /// @param token The address of token updated.
    /// @param oldGateway The corresponding address of the old gateway.
    /// @param newGateway The corresponding address of the new gateway.
    event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice Return the corresponding gateway address for given token address.
    /// @param _token The address of token to query.
    function getERC20Gateway(address _token) external view returns (address);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Request ERC20 token transfer from users to gateways.
    /// @param sender The address of sender to request fund.
    /// @param token The address of token to request.
    /// @param amount The amount of token to request.
    function requestERC20(address sender, address token, uint256 amount) external returns (uint256);
}
