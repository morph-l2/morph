// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL1ETHGateway} from "../../l1/gateways/IL1ETHGateway.sol";
import {IL2CrossDomainMessenger} from "../IL2CrossDomainMessenger.sol";
import {IL2ETHGateway} from "./IL2ETHGateway.sol";

import {GatewayBase} from "../../libraries/gateway/GatewayBase.sol";

/// @title L2ETHGateway
/// @notice The `L2ETHGateway` contract is used to withdraw ETH token on layer 2 and
/// finalize deposit ETH from layer 1.
/// @dev The ETH are not held in the gateway. The ETH will be sent to the `L2CrossDomainMessenger` contract.
/// On finalizing deposit, the Ether will be transferred from `L2CrossDomainMessenger`, then transfer to recipient.
contract L2ETHGateway is GatewayBase, IL2ETHGateway {
    /***************
     * Constructor *
     ***************/

    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize the storage of L2ETHGateway.
    /// @param _counterpart The address of L1ETHGateway in L2.
    /// @param _router The address of L2GatewayRouter.
    /// @param _messenger The address of L2CrossDomainMessenger.
    function initialize(address _counterpart, address _router, address _messenger) external initializer {
        require(_router != address(0), "zero router address");
        GatewayBase._initialize(_counterpart, _router, _messenger);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc IL2ETHGateway
    function withdrawETH(uint256 _amount, uint256 _gasLimit) external payable override {
        _withdraw(_msgSender(), _amount, new bytes(0), _gasLimit);
    }

    /// @inheritdoc IL2ETHGateway
    function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) public payable override {
        _withdraw(_to, _amount, new bytes(0), _gasLimit);
    }

    /// @inheritdoc IL2ETHGateway
    function withdrawETHAndCall(
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    ) public payable override {
        _withdraw(_to, _amount, _data, _gasLimit);
    }

    /// @inheritdoc IL2ETHGateway
    function finalizeDepositETH(
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external payable override onlyCallByCounterpart nonReentrant {
        require(msg.value == _amount, "msg.value mismatch");

        // solhint-disable-next-line avoid-low-level-calls
        (bool _success, ) = _to.call{value: _amount}("");
        require(_success, "ETH transfer failed");

        _doCallback(_to, _data);

        emit FinalizeDepositETH(_from, _to, _amount, _data);
    }

    /**********************
     * Internal Functions *
     **********************/

    function _withdraw(
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    ) internal virtual nonReentrant {
        require(msg.value > 0, "withdraw zero eth");

        // 1. Extract real sender if this call is from L1GatewayRouter.
        address _from = _msgSender();
        if (router == _from) {
            (_from, _data) = abi.decode(_data, (address, bytes));
        }

        // @note no rate limit here, since ETH is limited in messenger
        bytes memory _message = abi.encodeCall(IL1ETHGateway.finalizeWithdrawETH, (_from, _to, _amount, _data));

        uint256 nonce = IL2CrossDomainMessenger(messenger).messageNonce();
        IL2CrossDomainMessenger(messenger).sendMessage{value: msg.value}(counterpart, _amount, _message, _gasLimit);

        emit WithdrawETH(_from, _to, _amount, _data, nonce);
    }
}
