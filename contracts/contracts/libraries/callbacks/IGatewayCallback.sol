// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

interface IGatewayCallback {
    function onGatewayCallback(bytes memory data) external;
}
