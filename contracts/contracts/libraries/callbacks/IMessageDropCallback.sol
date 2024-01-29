// SPDX-License-Identifier: MIT

pragma solidity ^0.8.9;

interface IMessageDropCallback {
    function onDropMessage(bytes memory message) external payable;
}
