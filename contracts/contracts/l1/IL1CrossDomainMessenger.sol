// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

interface IL1CrossDomainMessenger is ICrossDomainMessenger {
    /**********
     * Events *
     **********/

    /// @dev Emitted when the rollup contract address is updated.
    /// @param oldRollup The address of the old rollup contract.
    /// @param newRollup The address of the new rollup contract.
    event UpdateRollup(address oldRollup, address newRollup);

    /// @notice Emitted when the maximum number of times each message can be replayed is updated.
    /// @param oldMaxReplayTimes The old maximum number of times each message can be replayed.
    /// @param newMaxReplayTimes The new maximum number of times each message can be replayed.
    event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Prove a L2 => L1 message with message proof and relay a L2 => L1 message.
    /// @param _from The address of the sender of the message.
    /// @param _to The address of the recipient of the message.
    /// @param _value The msg.value passed to the message call.
    /// @param _nonce The nonce of the message to avoid replay attack.
    /// @param _message The content of the message.
    /// @param _withdrawalProof Merkle tree proof of the message.
    /// @param _withdrawalRoot Merkle tree root of the proof.
    function proveAndRelayMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _nonce,
        bytes memory _message,
        bytes32[32] calldata _withdrawalProof,
        bytes32 _withdrawalRoot
    ) external;

    /// @notice Replay an existing message.
    /// @param from The address of the sender of the message.
    /// @param to The address of the recipient of the message.
    /// @param value The msg.value passed to the message call.
    /// @param messageNonce The nonce for the message to replay.
    /// @param message The content of the message.
    /// @param newGasLimit New gas limit to be used for this message.
    /// @param refundAddress The address of account who will receive the refunded fee.
    function replayMessage(
        address from,
        address to,
        uint256 value,
        uint256 messageNonce,
        bytes memory message,
        uint32 newGasLimit,
        address refundAddress
    ) external payable;

    /// @notice Drop a skipped message.
    /// @param from The address of the sender of the message.
    /// @param to The address of the recipient of the message.
    /// @param value The msg.value passed to the message call.
    /// @param messageNonce The nonce for the message to drop.
    /// @param message The content of the message.
    function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes memory message) external;
}
