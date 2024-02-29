// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {ICrossDomainMessenger} from "../../libraries/ICrossDomainMessenger.sol";

contract CorssMessage {
    address public immutable messenger;

    constructor(address _messenger) {
        messenger = _messenger;
    }

    // function sendMessage(
    //     address counterpart,
    //     bytes memory _msg,
    //     uint256 _gasLimit
    // ) public payable {
    //        bytes memory _message = abi.encodeCall(
    //         IL2ETHGateway.finalizeDepositETH,
    //         (_from, _to, _amount, _data)
    //     );

    //     ICrossDomainMessenger(messenger).sendMessage(
    //         counterpart,
    //         0,
    //         _msg,
    //         _gasLimit
    //     );
    // }

    // function relayMessage(
    //     bytes memory _msg,
    //     uint256 _gasLimit
    // ) public payable {
    //     ICrossDomainMessenger(messenger).sendMessage(
    //         counterpart,
    //         0,
    //         _msg,
    //         _gasLimit
    //     );
    // }
}
