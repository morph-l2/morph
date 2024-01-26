// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {L2MessageBaseTest} from "./L2MessageBase.t.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L2ToL1MessagePasserTest is L2MessageBaseTest {
    function testAppendMessage() external {
        for (uint256 i = 0; i < 1024; i++) {
            bytes32 leafHash = bytes32(i);

            _appendMessageHash(leafHash);

            hevm.prank(Predeploys.L2_CROSS_DOMAIN_MESSENGER);
            hevm.expectEmit(true, true, true, true);
            emit AppendMessage(i, leafHash);

            l2ToL1MessagePasser.appendMessage(leafHash);

            assertEq(l2ToL1MessagePasser.messageRoot(), getTreeRoot());
        }
    }
}
