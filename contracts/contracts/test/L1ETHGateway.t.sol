// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {IL2ETHGateway} from "../l2/gateways/IL2ETHGateway.sol";
import {IL1ETHGateway} from "../l1/gateways/IL1ETHGateway.sol";
import {L1ETHGateway} from "../l1/gateways/L1ETHGateway.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {Constants} from "../libraries/constants/Constants.sol";

contract L1ETHGatewayTest is L1GatewayBaseTest {
    address counterpartGateway;
    L1ETHGateway l1ETHGatewayImplTemp;

    function setUp() public virtual override {
        super.setUp();
        counterpartGateway = l1ETHGateway.counterpart();
    }

    function test_initialize_initializeAgain_reverts() public {
        // Test the initializer modifier to ensure initialize() can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1ETHGateway.initialize(address(1), address(1), address(1));
    }

    function test_initialize_zeroAddress_reverts() external {
        hevm.startPrank(multisig);

        // Deploy a transparent upgradeable proxy for the L1ETHGateway contract.
        TransparentUpgradeableProxy l1ETHGatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new instance of the L1ETHGateway contract implementation.
        l1ETHGatewayImplTemp = new L1ETHGateway();

        // Expect revert when the address of _counterpart equals the zero address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1ETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1ETHGatewayImplTemp),
            abi.encodeCall(
                L1ETHGateway.initialize,
                (
                    address(0), // _counterpart
                    address(l1ETHGatewayProxyTemp), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert when the address of _router equals the zero address.
        hevm.expectRevert("zero router address");
        ITransparentUpgradeableProxy(address(l1ETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1ETHGatewayImplTemp),
            abi.encodeCall(
                L1ETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert when the address of _messenger equals the zero address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1ETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1ETHGatewayImplTemp),
            abi.encodeCall(
                L1ETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1ETHGatewayProxyTemp), // _router
                    address(0) // _messenger
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_succeed() public {
        hevm.startPrank(multisig);

        // Deploy a transparent upgradeable proxy for the L1ETHGateway contract.
        TransparentUpgradeableProxy l1ETHGatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new instance of the L1ETHGateway contract implementation.
        l1ETHGatewayImplTemp = new L1ETHGateway();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1ETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1ETHGatewayImplTemp),
            abi.encodeCall(
                L1ETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1ETHGatewayProxyTemp), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Cast the proxy contract address to the L1ETHGateway contract type to call its methods.
        L1ETHGateway l1ETHGatewayTemp = L1ETHGateway((address(l1ETHGatewayProxyTemp)));
        hevm.stopPrank();

        // Verify the counterpart, router and messenger are initialized successfully.
        assertEq(l1ETHGatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l1ETHGatewayTemp.router(), address(l1ETHGatewayProxyTemp));
        assertEq(l1ETHGatewayTemp.messenger(), address(l1CrossDomainMessenger));
    }

    function test_depositETH_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositETH(false, amount, gasLimit, feePerGas);
    }

    function test_depositETHWithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipient(false, amount, recipient, gasLimit, feePerGas);
    }

    function test_depositETHWithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipientAndCalldata(false, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_routerDepositETH_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositETH(true, amount, gasLimit, feePerGas);
    }

    function test_routerDepositETHWithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipient(true, amount, recipient, gasLimit, feePerGas);
    }

    function test_routerDepositETHWithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipientAndCalldata(true, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_dropMessage_succeeds(uint256 amount, address recipient, bytes memory dataToCall) public {
        amount = bound(amount, 1, address(this).balance);
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (address(this), recipient, amount, dataToCall)
        );
        l1ETHGateway.depositETHAndCall{value: amount}(recipient, amount, dataToCall, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // ETH transfer failed, revert
        revertOnReceive = true;
        hevm.expectRevert("ETH transfer failed");
        l1CrossDomainMessenger.dropMessage(address(l1ETHGateway), address(counterpartGateway), amount, 0, message);

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ETHGateway.RefundETH(address(this), amount);

        revertOnReceive = false;
        uint256 balance = address(this).balance;
        l1CrossDomainMessenger.dropMessage(address(l1ETHGateway), address(counterpartGateway), amount, 0, message);
        assertEq(balance + amount, address(this).balance);
    }

    function test_finalizeWithdrawETH_counterErr_fails(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, address(this).balance / 2);

        address _from = address(uint160(address(counterpartGateway)) + 1);
        // deposit some ETH to L1CrossDomainMessenger
        l1ETHGateway.depositETH{value: amount}(amount, defaultGasLimit);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(
            IL1ETHGateway.finalizeWithdrawETH,
            (sender, recipient, amount, dataToCall)
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, address(l1ETHGateway), amount, 0, message)
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            _from,
            address(l1ETHGateway),
            amount,
            0,
            message
        );

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = recipient.balance;
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
        // counterpart is not L2ETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(_xDomainCalldataHash);
        l1CrossDomainMessenger.proveAndRelayMessage(_from, address(l1ETHGateway), amount, 0, message, wdProof, wdRoot);

        assertEq(messengerBalance, address(l1CrossDomainMessenger).balance);
        assertEq(recipientBalance, recipient.balance);
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
    }

    function test_finalizeWithdrawETH_succeeds() public {
        address sender = bob;
        uint256 amount = 1000;
        address recipient = address(2048);
        address _from = counterpartGateway;
        // deposit some ETH to L1CrossDomainMessenger
        l1ETHGateway.depositETH{value: amount}(amount, defaultGasLimit);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(IL1ETHGateway.finalizeWithdrawETH, (sender, recipient, amount, ""));
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, address(l1ETHGateway), amount, 0, message)
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            _from,
            address(l1ETHGateway),
            amount,
            0,
            message
        );

        // Verify it throws an error "msg.value mismatch" as expected when msg.value != _amount.
        hevm.expectRevert("msg.value mismatch");
        // Simulate the value of l1CrossDomainMessenger.xDomainMessageSender() as counterpartGateway.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(counterpartGateway)
        );
        // Update msg.sender as address(l1CrossDomainMessenger).
        hevm.prank(address(l1CrossDomainMessenger));
        l1ETHGateway.finalizeWithdrawETH(sender, recipient, amount, "");

        // Set recipient to the current contract's address to trigger the receive function.
        recipient = address(this);
        // Simulate the value of l1CrossDomainMessenger.xDomainMessageSender() as counterpartGateway.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(counterpartGateway)
        );

        // Update msg.sender is address(l1CrossDomainMessenger).
        hevm.prank(address(l1CrossDomainMessenger));
        // Set the revertOnReceive as true to simulate a failure in the receive function.
        revertOnReceive = true;
        // Verify it throws an error "ETH transfer failed" as expected when recipient.call() failed.
        hevm.expectRevert("ETH transfer failed");
        l1ETHGateway.finalizeWithdrawETH{value: amount}(sender, recipient, amount, "");
        recipient = address(2048);

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = recipient.balance;
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
        // counterpart is L2ETHGateway
        // emit FinalizeWithdrawETH from L1ETHGateway
        {
            hevm.expectEmit(true, false, false, true);
            emit IL1ETHGateway.FinalizeWithdrawETH(sender, address(recipient), amount, "");
        }

        l1CrossDomainMessenger.proveAndRelayMessage(_from, address(l1ETHGateway), amount, 0, message, wdProof, wdRoot);
        assertEq(messengerBalance - amount, address(l1CrossDomainMessenger).balance);
        assertEq(recipientBalance + amount, address(recipient).balance);
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
    }

    function testOnDropMessage() external {
        address sender = bob;
        uint256 amount = 1000;
        address recipient = address(2048);
        bytes memory message = "message";

        // Assign 10 ether to the current contract's address and l1CrossDomainMessenger contract's address.
        hevm.deal(address(this), 10 ether);
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Verify the onlyInDropContext modifier.
        // Verify it throws an error "only messenger can call" when msg.sender is not the messenger.
        hevm.expectRevert("only messenger can call");
        l1ETHGateway.onDropMessage(message);

        // Verify it throws an error "only called in drop context" when xDomainMessageSender is not the DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.prank(address(l1CrossDomainMessenger));
        hevm.expectRevert("only called in drop context");
        l1ETHGateway.onDropMessage(message);

        // Simulate l1CrossDomainMessenger.xDomainMessageSender() returning DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(Constants.DROP_XDOMAIN_MESSAGE_SENDER)
        );
        // Update msg.sender is address(l1CrossDomainMessenger).
        hevm.prank(address(l1CrossDomainMessenger));
        // Verify it throws an error "invalid selector" when condition is not satisfied.
        hevm.expectRevert("invalid selector");
        l1ETHGateway.onDropMessage(message);

        // Verify _message should start with 0x232e8748 (finalizeDepositETH(address,address,uint256,bytes).
        // Create a message with the valid selector and set the sender as address(this).
        message = abi.encodeWithSelector(
            IL2ETHGateway.finalizeDepositETH.selector,
            address(this),
            recipient,
            amount,
            message
        );

        // Update msg.sender is address(l1CrossDomainMessenger).
        hevm.prank(address(l1CrossDomainMessenger));
        // Set the revertOnReceive as true to simulate a failure in the receive function.
        revertOnReceive = true;

        // Verify it throws an error "ETH transfer failed" when recipient.call() fails as expected.
        hevm.expectRevert("ETH transfer failed");
        l1ETHGateway.onDropMessage{value: amount}(message);

        message = abi.encodeWithSelector(IL2ETHGateway.finalizeDepositETH.selector, sender, recipient, amount, message);

        // Verify the RefundETH event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1ETHGateway.RefundETH(sender, amount);

        // Update msg.sender is address(l1CrossDomainMessenger).
        hevm.prank(address(l1CrossDomainMessenger));
        uint256 originalBalance = address(sender).balance;
        l1ETHGateway.onDropMessage{value: amount}(message);
        // Verify that the sender's balance increased by the specified amount.
        assertEq(originalBalance + amount, address(sender).balance);
    }

    function _depositETH(bool useRouter, uint256 amount, uint256 gasLimit, uint256 feePerGas) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (address(this), address(this), amount, new bytes(0))
        );

        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1ETHGateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero eth");
            if (useRouter) {
                l1GatewayRouter.depositETH{value: amount}(amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount}(amount, gasLimit);
            }
        } else {
            // emit QueueTransaction from L1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(l1ETHGateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositETH from L1ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL1ETHGateway.DepositETH(address(this), address(this), amount, new bytes(0), 0);

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositETH{value: amount + feeToPay + EXTRA_VALUE}(amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount + feeToPay + EXTRA_VALUE}(amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l1CrossDomainMessenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _depositETHWithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (address(this), recipient, amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1ETHGateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero eth");
            if (useRouter) {
                l1GatewayRouter.depositETH{value: amount}(recipient, amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount}(recipient, amount, gasLimit);
            }
        } else {
            // emit QueueTransaction from L1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(l1ETHGateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositETH from L1ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL1ETHGateway.DepositETH(address(this), recipient, amount, new bytes(0), 0);

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositETH{value: amount + feeToPay + EXTRA_VALUE}(recipient, amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount + feeToPay + EXTRA_VALUE}(recipient, amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l1CrossDomainMessenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _depositETHWithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (address(this), recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1ETHGateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero eth");
            if (useRouter) {
                l1GatewayRouter.depositETHAndCall{value: amount}(recipient, amount, dataToCall, gasLimit);
            } else {
                l1ETHGateway.depositETHAndCall{value: amount}(recipient, amount, dataToCall, gasLimit);
            }
        } else {
            // emit QueueTransaction from L1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(l1ETHGateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositETH from L1ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL1ETHGateway.DepositETH(address(this), recipient, amount, dataToCall, 0);

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositETHAndCall{value: amount + feeToPay + EXTRA_VALUE}(
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                l1ETHGateway.depositETHAndCall{value: amount + feeToPay + EXTRA_VALUE}(
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(amount + messengerBalance, address(l1CrossDomainMessenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }
}
