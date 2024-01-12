// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {StandardBridge} from "../universal/StandardBridge.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {SafeCall} from "../libraries/SafeCall.sol";

contract TestL1Multiple {
    CrossDomainMessenger public immutable MESSENGER;
    StandardBridge public immutable BRIDGE;

    uint256 public immutable ERC20DepositAmount;
    address public to;

    constructor(
        address payable _messenger,
        address payable _bridge,
        uint256 _erc20Amount,
        address _to
    ) {
        MESSENGER = CrossDomainMessenger(_messenger);
        BRIDGE = StandardBridge(_bridge);
        ERC20DepositAmount = _erc20Amount;
        to = _to;
    }

    function sendDeposit(
        address[] memory l1tokens,
        address[] memory l2tokens,
        uint256 ethAmount,
        uint32 minGasLimit,
        bytes calldata extraData
    ) public payable {
        require(l1tokens.length == l2tokens.length, "length not equal");
        for (uint256 i = 0; i < l1tokens.length; i++) {
            IERC20(l1tokens[i]).approve(address(BRIDGE), ERC20DepositAmount);
            // call l1 standBridge send erc20
            BRIDGE.bridgeERC20To(
                l1tokens[i],
                l2tokens[i],
                to,
                ERC20DepositAmount,
                minGasLimit,
                extraData
            );
        }
        if (ethAmount > 0) {
            // call l1 standBridge send eth
            BRIDGE.bridgeETHTo{value: ethAmount}(to, minGasLimit, extraData);
        }
        if (msg.value > 0) {
            uint256 crossETHAmount = msg.value - ethAmount;
            // call l1 cross domain messenger
            MESSENGER.sendMessage{value: crossETHAmount}(
                to,
                extraData,
                minGasLimit
            );
        }
    }
}
