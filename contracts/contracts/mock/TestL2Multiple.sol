// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {StandardBridge} from "../universal/StandardBridge.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {SafeCall} from "../libraries/SafeCall.sol";

contract TestL2Multiple {
    CrossDomainMessenger public immutable L2MESSENGER;
    StandardBridge public immutable L2BRIDGE;

    uint256 public immutable ERC20WithdrawalAmount;
    address public to;

    constructor(
        address payable _messenger,
        address payable _bridge,
        uint256 _erc20Amount,
        address _to
    ) {
        L2MESSENGER = CrossDomainMessenger(_messenger);
        L2BRIDGE = StandardBridge(_bridge);
        ERC20WithdrawalAmount = _erc20Amount;
        to = _to;
    }

    function sendWithdrawal(
        address[] memory l1tokens,
        address[] memory l2tokens,
        uint256 ethAmount,
        uint32 minGasLimit,
        bytes calldata extraData
    ) public payable {
        require(l1tokens.length == l2tokens.length, "length not equal");
        for (uint256 i = 0; i < l2tokens.length; i++) {
            IERC20(l2tokens[i]).approve(address(L2BRIDGE), ERC20WithdrawalAmount);
            // call l1 standBridge send erc20
            L2BRIDGE.bridgeERC20To(
                l2tokens[i],
                l1tokens[i],
                to,
                ERC20WithdrawalAmount,
                minGasLimit,
                extraData
            );
        }
        if (ethAmount > 0) {
            // call l1 standBridge send eth
            L2BRIDGE.bridgeETHTo{value: ethAmount}(to, minGasLimit, extraData);
        }
        if (msg.value > 0) {
            uint256 crossETHAmount = msg.value - ethAmount;
            // call l1 cross domain messenger
            L2MESSENGER.sendMessage{value: crossETHAmount}(
                to,
                extraData,
                minGasLimit
            );
        }
    }
}
