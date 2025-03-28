// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

interface IMessageDropCallback {
    function onDropMessage(bytes memory message) external payable;
}
