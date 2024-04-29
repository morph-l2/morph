// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL1MessageQueue} from "../L1/rollup/IL1MessageQueue.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {L1MessageQueueWithGasPriceOracle} from "../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";

contract L1MessageQueueTest is L1MessageBaseTest {
    L1MessageQueueWithGasPriceOracle l1MessageQueue;

    function setUp() public virtual override {
        super.setUp();
        l1MessageQueue = l1MessageQueueWithGasPriceOracle;
    }

    function test_validateGasLimit() external {
        // store alice as messenger
        upgradeStorage(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());

        // append message
        hevm.prank(multisig);
        l1MessageQueue.updateMaxGasLimit(1);
        hevm.prank(alice);
        hevm.expectRevert("Gas limit must not exceed maxGasLimit");
        l1MessageQueue.appendCrossDomainMessage(alice, 3, "0x0");

        hevm.prank(multisig);
        l1MessageQueue.updateMaxGasLimit(100);
        hevm.prank(alice);
        hevm.expectRevert(
            "Insufficient gas limit, must be above intrinsic gas"
        );
        l1MessageQueue.appendCrossDomainMessage(alice, 3, "0x0");
    }

    function test_appendCrossDomainMessage() external {
        // store alice as messenger
        upgradeStorage(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());
        address sender = AddressAliasHelper.applyL1ToL2Alias(address(alice));
        bytes memory _calldata = "0x0";
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueue.QueueTransaction(
            sender,
            alice,
            0,
            0,
            gasLimit,
            _calldata
        );
        hevm.startPrank(alice);
        l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
        hevm.stopPrank();
    }

    function test_appendEnforcedTransaction() external {
        hevm.prank(multisig);
        assertEq(alice, l1MessageQueue.ENFORCED_TX_GATEWAAY());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());
        bytes memory _calldata = "0x0";
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");
        hevm.prank(alice);
        l1MessageQueue.appendEnforcedTransaction(
            alice,
            bob,
            0,
            gasLimit,
            _calldata
        );
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
    }

    function test_pop_dropCrossDomainMessage() external {
        // store alice as messenger and rollup
        upgradeStorage(address(alice), address(alice), address(alice));
        assertEq(alice, l1MessageQueue.MESSENGER());
        assertEq(alice, l1MessageQueue.ROLLUP_CONTRACT());
        bytes memory _calldata = "0x0";
        uint256 gasLimit = l1MessageQueue.calculateIntrinsicGasFee("0x0");
        // append 10 message
        hevm.startPrank(alice);
        for (uint64 i = 0; i < 10; i++) {
            l1MessageQueue.appendCrossDomainMessage(alice, gasLimit, _calldata);
        }
        // pop all 10 message
        l1MessageQueue.popCrossDomainMessage(0, 10, 0x3ff);
        for (uint64 i = 0; i < 10; i++) {
            assertTrue(l1MessageQueue.isMessageSkipped(i));
        }
        // drop all 10 message
        for (uint64 i = 0; i < 10; i++) {
            l1MessageQueue.dropCrossDomainMessage(i);
            assertTrue(l1MessageQueue.isMessageDropped(i));
        }
        hevm.stopPrank();
    }
}
