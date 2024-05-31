// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Constants} from "../libraries/constants/Constants.sol";
import {CrossDomainMessenger} from "../libraries/CrossDomainMessenger.sol";
import {ICrossDomainMessenger} from "../libraries/CrossDomainMessenger.sol";
import {L2ToL1MessagePasser} from "./system/L2ToL1MessagePasser.sol";
import {IL2CrossDomainMessenger} from "./IL2CrossDomainMessenger.sol";

/// @title L2CrossDomainMessenger
/// @notice The L2CrossDomainMessenger is a high-level interface for message passing between L1 and
///         L2 on the L2 side. Users are generally encouraged to use this contract instead of lower
///         level message passing contracts.
contract L2CrossDomainMessenger is CrossDomainMessenger, IL2CrossDomainMessenger {
    /*************
     * Variables *
     *************/

    /// @notice Mapping from L2 message hash to the timestamp when the message is sent.
    mapping(bytes32 => uint256) public messageSendTimestamp;

    /// @notice Mapping from L1 message hash to a boolean value indicating if the message has been successfully executed.
    mapping(bytes32 => bool) public isL1MessageExecuted;

    /***************
     * Constructor *
     ***************/

    constructor() {
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    function initialize(address _counterpart) external initializer {
        if (_counterpart == address(0)) revert ErrZeroAddress();
        __Messenger_init(_counterpart, address(0));
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc ICrossDomainMessenger
    function sendMessage(
        address _to,
        uint256 _value,
        bytes memory _message,
        uint256 _gasLimit
    ) external payable override whenNotPaused {
        _sendMessage(_to, _value, _message, _gasLimit);
    }

    /// @inheritdoc ICrossDomainMessenger
    function sendMessage(
        address _to,
        uint256 _value,
        bytes calldata _message,
        uint256 _gasLimit,
        address
    ) external payable override whenNotPaused {
        _sendMessage(_to, _value, _message, _gasLimit);
    }

    /// @inheritdoc IL2CrossDomainMessenger
    function relayMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _nonce,
        bytes memory _message
    ) external override whenNotPaused {
        // It is impossible to deploy a contract with the same address, reentrance is prevented in nature.
        require(
            AddressAliasHelper.undoL1ToL2Alias(_msgSender()) == counterpart,
            "Caller is not L1CrossDomainMessenger"
        );

        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(_from, _to, _value, _nonce, _message));

        require(!isL1MessageExecuted[_xDomainCalldataHash], "Message was already successfully executed");

        _executeMessage(_from, _to, _value, _message, _xDomainCalldataHash);
    }

    function messageNonce() external view override(ICrossDomainMessenger, CrossDomainMessenger) returns (uint256) {
        return L2ToL1MessagePasser(Predeploys.L2_TO_L1_MESSAGE_PASSER).leafNodesCount();
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev Internal function to send cross domain message.
    /// @param _to The address of account who receive the message.
    /// @param _value The amount of ether passed when call target contract.
    /// @param _message The content of the message.
    /// @param _gasLimit Optional gas limit to complete the message relay on corresponding chain.
    function _sendMessage(address _to, uint256 _value, bytes memory _message, uint256 _gasLimit) internal nonReentrant {
        require(msg.value == _value, "msg.value mismatch");

        address messagePasser = Predeploys.L2_TO_L1_MESSAGE_PASSER;
        uint256 _nonce = L2ToL1MessagePasser(messagePasser).leafNodesCount();
        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(_msgSender(), _to, _value, _nonce, _message));

        // normally this won't happen, since each message has different nonce, but just in case.
        require(messageSendTimestamp[_xDomainCalldataHash] == 0, "Duplicated message");
        messageSendTimestamp[_xDomainCalldataHash] = block.timestamp;

        L2ToL1MessagePasser(messagePasser).appendMessage(_xDomainCalldataHash);

        emit SentMessage(_msgSender(), _to, _value, _nonce, _gasLimit, _message);
    }

    /// @dev Internal function to execute a L1 => L2 message.
    /// @param _from The address of the sender of the message.
    /// @param _to The address of the recipient of the message.
    /// @param _value The msg.value passed to the message call.
    /// @param _message The content of the message.
    /// @param _xDomainCalldataHash The hash of the message.
    function _executeMessage(
        address _from,
        address _to,
        uint256 _value,
        bytes memory _message,
        bytes32 _xDomainCalldataHash
    ) internal {
        // @note check more `_to` address to avoid attack in the future when we add more gateways.
        require(_to != Predeploys.L2_TO_L1_MESSAGE_PASSER, "Forbid to call l2 to l1 message passer");
        _validateTargetAddress(_to);

        // @note This usually will never happen, just in case.
        require(_from != xDomainMessageSender, "Invalid message sender");

        xDomainMessageSender = _from;
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, ) = _to.call{value: _value}(_message);
        // reset value to refund gas.
        xDomainMessageSender = Constants.DEFAULT_XDOMAIN_MESSAGE_SENDER;

        if (success) {
            isL1MessageExecuted[_xDomainCalldataHash] = true;
            emit RelayedMessage(_xDomainCalldataHash);
        } else {
            emit FailedRelayedMessage(_xDomainCalldataHash);
        }
    }
}
