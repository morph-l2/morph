// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Types} from "../libraries/Types.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {Encoding} from "../libraries/Encoding.sol";
import {Semver} from "../universal/Semver.sol";
import {Tree} from "../universal/Tree.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {SafeCall} from "../libraries/SafeCall.sol";

/**
 * @custom:proxied
 * @custom:predeploy 0x5300000000000000000000000000000000000001
 * @title L2ToL1MessagePasser
 * @notice The L2ToL1MessagePasser is a dedicated contract where messages that are being sent from
 *         L2 to L1 can be stored. The storage root of this contract is pulled up to the top level
 *         of the L2 output to reduce the cost of proving the existence of sent messages.
 */
contract L2ToL1MessagePasser is Semver, Tree {
    /**
     * @notice The L1 gas limit set when eth is withdrawn using the receive() function.
     */
    uint256 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /**
     * @notice Current message version identifier.
     */
    uint16 public constant MESSAGE_VERSION = 1;

    /**
     * @notice Includes the message hashes for all withdrawals
     */
    mapping(bytes32 => bool) public sentMessages;

    /**
     * @notice A unique value hashed with each withdrawal.
     */
    uint240 internal msgNonce;

    /// @notice The merkle root of the current merkle tree.
    /// @dev This is actual equal to `branches[n]`.
    bytes32 public messageRoot;

    /**
     * @notice Lock eth at this address when withdrawal.
     */
    address public immutable recAddr;

    /**
     * @notice Emitted any time a withdrawal is initiated.
     *
     * @param nonce          Unique value corresponding to each withdrawal.
     * @param sender         The L2 account address which initiated the withdrawal.
     * @param target         The L1 account address the call will be send to.
     * @param value          The ETH value submitted for withdrawal, to be forwarded to the target.
     * @param gasLimit       The minimum amount of gas that must be provided when withdrawing.
     * @param data           The data to be forwarded to the target on L1.
     * @param withdrawalHash The hash of the withdrawal.
     * @param rootHash       The root hash of the merkle tree after the withdrawal is included.
     */
    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash,
        bytes32 rootHash
    );

    /**
     * @notice Emitted when the balance of this contract is burned.
     *
     * @param amount Amount of ETh that was burned.
     */
    event WithdrawerBalanceBurnt(uint256 indexed amount);

    /**
     * @custom:semver 1.0.0
     */
    constructor(address _l1CrossDomainMessenger) Semver(1, 0, 0) {
        recAddr = AddressAliasHelper.applyL1ToL2Alias(_l1CrossDomainMessenger);
        messageRoot = getTreeRoot();
    }

    /**
     * @notice Allows users to withdraw ETH by sending directly to this contract.
     */
    receive() external payable {
        initiateWithdrawal(msg.sender, RECEIVE_DEFAULT_GAS_LIMIT, bytes(""));
    }

    /**
     * @notice Sends a message from L2 to L1.
     *
     * @param _target   Address to call on L1 execution.
     * @param _gasLimit Minimum gas limit for executing the message on L1.
     * @param _data     Data to forward to L1 target.
     */
    function initiateWithdrawal(
        address _target,
        uint256 _gasLimit,
        bytes memory _data
    ) public payable {
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: messageNonce(),
                sender: msg.sender,
                target: _target,
                value: msg.value,
                gasLimit: _gasLimit,
                data: _data
            })
        );

        _appendMessageHash(withdrawalHash);
        messageRoot = getTreeRoot();

        emit MessagePassed(
            messageNonce(),
            msg.sender,
            _target,
            msg.value,
            _gasLimit,
            _data,
            withdrawalHash,
            messageRoot
        );
        unchecked {
            msgNonce++;
        }

        if (msg.value > 0) {
            bool success = SafeCall.call(recAddr, gasleft(), msg.value, hex"");
            require(success, "L2ToL1MessagePasser: ETH transfer failed");
            emit WithdrawerBalanceBurnt(msg.value);
        }
    }

    /**
     * @notice Retrieves the next message nonce. Message version will be added to the upper two
     *         bytes of the message nonce. Message version allows us to treat messages as having
     *         different structures.
     *
     * @return Nonce of the next message to be sent, with added message version.
     */
    function messageNonce() public view returns (uint256) {
        return Encoding.encodeVersionedNonce(msgNonce, MESSAGE_VERSION);
    }
}
