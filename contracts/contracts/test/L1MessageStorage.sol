// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/console.sol";

import {TestL1MessageStorage} from "../mock/TestL1MessageStorage.sol";
import {L1MessageStorage} from "../L1/L1MessageStorage.sol";

contract L1MessageStorage_Test is Test {
    TestL1MessageStorage public storageTest;
    address user = address(1);
    uint256 value = 1 ether;
    uint256 gasLimit = 100000;

    /// @dev The bitmap for skipped messages, where `skippedMessageBitmap[i]` keeps the bits from `[i*256, (i+1)*256)`.
    mapping(uint256 => uint256) private skippedMessageBitmap;

    function setUp() public virtual {
        storageTest = new TestL1MessageStorage();
    }

    function testQueueTransaction() external {
        storageTest.queueTransaction(user, user, value, gasLimit, "0x");
        assertEq(storageTest.nextCrossDomainMessageIndex(), 1);
    }

    function testPopCrossDomainMessage() external {
        // queue txs
        for (uint256 i = 0; i < 1024; i++) {
            storageTest.queueTransaction(user, user, value, gasLimit, "0x");
        }
        // pop with skip
        storageTest.popCrossDomainMessage(0, 2, 0x3);
    }
}
