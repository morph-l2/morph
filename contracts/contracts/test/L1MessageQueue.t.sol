// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/* Testing utilities */
import  "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L2GasPriceOracleTest} from "./base/L2GasPriceOracle.t.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {L1MessageQueue} from "../L1/rollup/L1MessageQueue.sol";

contract L1MessageQueueTest is L1MessageBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    function test_validateGasLimit() external {
        // store alice as messenger
        upgradeProxy(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.messenger());

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
        upgradeProxy(address(alice), address(rollup), address(alice));
        assertEq(alice, l1MessageQueue.messenger());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());
        hevm.expectEmit(true, true, true, true);
        address sender = AddressAliasHelper.applyL1ToL2Alias(address(alice));
        emit QueueTransaction(sender, alice, 0, 0, 100, "0x0");
        hevm.prank(alice);
        l1MessageQueue.appendCrossDomainMessage(alice, 100, "0x0");
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
    }

    function test_appendEnforcedTransaction() external {
        hevm.prank(multisig);
        assertEq(alice, l1MessageQueue.enforcedTxGateway());
        // append message
        assertEq(0, l1MessageQueue.nextCrossDomainMessageIndex());
        hevm.prank(alice);
        l1MessageQueue.appendEnforcedTransaction(alice, bob, 0, 100, "0x0");
        assertEq(1, l1MessageQueue.nextCrossDomainMessageIndex());
    }

    function test_pop_dropCrossDomainMessage() external {
        // store alice as messenger and rollup
        upgradeProxy(address(alice), address(alice), address(alice));
        assertEq(alice, l1MessageQueue.messenger());
        assertEq(alice, l1MessageQueue.rollup());

        // append 10 message
        hevm.startPrank(alice);
        for (uint64 i = 0; i < 10; i++) {
            l1MessageQueue.appendCrossDomainMessage(alice, 100, "0x0");
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

    function upgradeProxy(
        address _messenger,
        address _rollup,
        address _enforcedTxGateway
    ) public {
        TransparentUpgradeableProxy l1MessageQueueProxy = TransparentUpgradeableProxy(
                payable(address(l1MessageQueue))
            );
        L1MessageQueue l1MessageQueueImpl = new L1MessageQueue(
            payable(_messenger), // _messenger
            address(_rollup), // _rollup
            address(_enforcedTxGateway) // _enforcedTxGateway
        );
        assertEq(_messenger, l1MessageQueueImpl.messenger());
        assertEq(_rollup, l1MessageQueueImpl.rollup());
        assertEq(_enforcedTxGateway, l1MessageQueueImpl.enforcedTxGateway());

        hevm.prank(multisig);
        proxyAdmin.upgrade(
            ITransparentUpgradeableProxy(address(l1MessageQueueProxy)),
            address(l1MessageQueueImpl)
        );

        assertEq(_messenger, l1MessageQueue.messenger());
        assertEq(_rollup, l1MessageQueue.rollup());
        assertEq(_enforcedTxGateway, l1MessageQueue.enforcedTxGateway());
    }
}
