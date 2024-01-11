// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {SafeCall} from "../libraries/SafeCall.sol";

contract MockCall {
    function callContract(address addr, bytes memory _message) public {
        bool success = SafeCall.call(addr, gasleft(), 0, _message);
        require(success, "call failed");
    }
}
