// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Types} from "../../libraries/common/Types.sol";
import {Semver} from "../../libraries/common/Semver.sol";
import {Tree} from "../../libraries/common/Tree.sol";
import {AddressAliasHelper} from "../../libraries/common/AddressAliasHelper.sol";
import {SafeCall} from "../../libraries/common/SafeCall.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x5300000000000000000000000000000000000001
 * @title L2ToL1MessagePasser
 * @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
 *         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
 *         of the L2 output to reduce the cost of proving the existence of sent messages.
 */
contract L2ToL1MessagePasser is Semver, Tree {
    /**********
     * Events *
     **********/

    /// @notice Emitted when a new message is added to the merkle tree.
    /// @param index The index of the corresponding message.
    /// @param messageHash The hash of the corresponding message.
    event AppendMessage(uint256 index, bytes32 messageHash);

    /*************
     * Variables *
     *************/

    /// @notice The merkle root of the current merkle tree.
    /// @dev This is actual equal to `branches[n]`.
    bytes32 public messageRoot;

    /***************
     * Constructor *
     ***************/
    /**
     * @custom:semver 1.0.0
     */
    constructor() Semver(1, 0, 0) {
        messageRoot = getTreeRoot();
    }

    /**
     * @notice Sends a message from L2 to L1.
     *
     * @param _messageHash  the message hash to append in tree.
     */
    function appendMessage(bytes32 _messageHash) external returns (bytes32) {
        require(
            msg.sender == Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            "only messenger"
        );
        // We can use the event to compute the merkle tree locally.
        emit AppendMessage(leafNodesCount, _messageHash);

        _appendMessageHash(_messageHash);
        messageRoot = getTreeRoot();

        return messageRoot;
    }
}
