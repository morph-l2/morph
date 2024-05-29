// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {IGateway} from "./IGateway.sol";
import {ICrossDomainMessenger} from "../ICrossDomainMessenger.sol";
import {IGatewayCallback} from "../callbacks/IGatewayCallback.sol";
import {Constants} from "../constants/Constants.sol";

/// @title GatewayBase
/// @notice The `GatewayBase` is a base contract for gateway contracts used in both in L1 and L2.
abstract contract GatewayBase is ReentrancyGuardUpgradeable, OwnableUpgradeable, IGateway {
    /*************
     * Variables *
     *************/

    /// @inheritdoc IGateway
    address public override counterpart;

    /// @inheritdoc IGateway
    address public override router;

    /// @inheritdoc IGateway
    address public override messenger;

    /// @dev The storage slots for future usage.
    uint256[46] private __gap;

    /**********************
     * Function Modifiers *
     **********************/

    modifier onlyCallByCounterpart() {
        address _messenger = messenger; // gas saving
        require(_msgSender() == _messenger, "only messenger can call");
        require(counterpart == ICrossDomainMessenger(_messenger).xDomainMessageSender(), "only call by counterpart");
        _;
    }

    modifier onlyInDropContext() {
        address _messenger = messenger; // gas saving
        require(_msgSender() == _messenger, "only messenger can call");
        require(
            Constants.DROP_XDOMAIN_MESSAGE_SENDER == ICrossDomainMessenger(_messenger).xDomainMessageSender(),
            "only called in drop context"
        );
        _;
    }

    /***************
     * Constructor *
     ***************/

    function _initialize(address _counterpart, address _router, address _messenger) internal {
        require(_counterpart != address(0), "zero counterpart address");
        require(_messenger != address(0), "zero messenger address");

        ReentrancyGuardUpgradeable.__ReentrancyGuard_init();
        OwnableUpgradeable.__Ownable_init();

        counterpart = _counterpart;
        messenger = _messenger;

        // @note: the address of router could be zero, if this contract is GatewayRouter.
        if (_router != address(0)) {
            router = _router;
        }
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev Internal function to forward calldata to target contract.
    /// @param _to The address of contract to call.
    /// @param _data The calldata passed to the contract.
    function _doCallback(address _to, bytes memory _data) internal {
        if (_data.length > 0 && _to.code.length > 0) {
            IGatewayCallback(_to).onGatewayCallback(_data);
        }
    }
}
