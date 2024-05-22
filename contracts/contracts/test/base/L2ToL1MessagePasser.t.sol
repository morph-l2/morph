// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {L2ToL1MessagePasser} from "../../l2/system/L2ToL1MessagePasser.sol";
import {L2MessageBaseTest} from "./L2MessageBase.t.sol";

contract L2ToL1MessagePasserTest is L2MessageBaseTest {
    function test_appendMessage_succeeds() external {
        for (uint256 i = 0; i < 1024; i++) {
            bytes32 leafHash = bytes32(i);

            _appendMessageHash(leafHash);

            hevm.prank(Predeploys.L2_CROSS_DOMAIN_MESSENGER);
            hevm.expectEmit(true, true, true, true);
            emit L2ToL1MessagePasser.AppendMessage(i, leafHash, getTreeRoot());

            l2ToL1MessagePasser.appendMessage(leafHash);

            assertEq(l2ToL1MessagePasser.messageRoot(), getTreeRoot());
        }
    }
}
