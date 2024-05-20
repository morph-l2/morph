// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Tree} from "../../libraries/common/Tree.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

/**
 * @title L2ToL1MessagePasser
 * @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
 *         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
 *         of the L2 output to reduce the cost of proving the existence of sent messages.
 */
contract L2ToL1MessagePasser is Tree {
    /**********
     * Events *
     **********/

    /// @notice Emitted when a new message is added to the merkle tree.
    /// @param index The index of the corresponding message.
    /// @param messageHash The hash of the corresponding message.
    /// @param rootHash The hash of the tree root after append the message.
    event AppendMessage(uint256 indexed index, bytes32 indexed messageHash, bytes32 indexed rootHash);

    /*************
     * Variables *
     *************/

    /// @notice The merkle root of the current merkle tree.
    /// @dev This is actual equal to `branches[n]`.
    bytes32 public messageRoot;

    /***************
     * Constructor *
     ***************/

    constructor() {
        messageRoot = getTreeRoot();
    }

    /// @notice Sends a message from L2 to L1.
    /// @param _messageHash  the message hash to append in tree.
    function appendMessage(bytes32 _messageHash) external returns (bytes32) {
        require(msg.sender == Predeploys.L2_CROSS_DOMAIN_MESSENGER, "only messenger");

        _appendMessageHash(_messageHash);
        messageRoot = getTreeRoot();
        // We can use the event to compute the merkle tree locally.
        emit AppendMessage(leafNodesCount - 1, _messageHash, messageRoot);
        return messageRoot;
    }
}
