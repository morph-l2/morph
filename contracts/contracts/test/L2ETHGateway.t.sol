// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2ToL1MessagePasser} from "../l2/system/L2ToL1MessagePasser.sol";
import {IL2ETHGateway} from "../l2/gateways/IL2ETHGateway.sol";
import {L2ETHGateway} from "../l2/gateways/L2ETHGateway.sol";
import {L2GatewayRouter} from "../l2/gateways/L2GatewayRouter.sol";
import {L2CrossDomainMessenger} from "../l2/L2CrossDomainMessenger.sol";
import {IL1ETHGateway} from "../l1/gateways/IL1ETHGateway.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";
import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";

contract L2ETHGatewayTest is L2GatewayBaseTest {
    L2ETHGateway private gateway;
    L2GatewayRouter private router;
    L2CrossDomainMessenger private l2Messenger;

    address private counterpartGateway;
    address private feeVault;
    address private l1Messenger;

    function setUp() public override {
        super.setUp();
        gateway = l2ETHGateway;
        router = l2GatewayRouter;
        counterpartGateway = gateway.counterpart();
        l2Messenger = l2CrossDomainMessenger;
        feeVault = l2FeeVault;
        l1Messenger = address(NON_ZERO_ADDRESS);
    }

    function test_withdrawETH_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _withdrawETH(false, amount, gasLimit, feePerGas);
    }

    function test_withdrawETHWithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawETHWithRecipient(false, amount, recipient, gasLimit, feePerGas);
    }

    function test_withdrawETHWithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawETHWithRecipientAndCalldata(false, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_routerWithdrawETH_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _withdrawETH(true, amount, gasLimit, feePerGas);
    }

    function test_routerWithdrawETHWithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawETHWithRecipient(true, amount, recipient, gasLimit, feePerGas);
    }

    function test_routerWithdrawETHWithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawETHWithRecipientAndCalldata(true, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_finalizeDepositETHFailedMocking_succeeds(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, address(this).balance / 2);

        // revert when caller is not messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositETH(sender, recipient, amount, dataToCall);

        MockCrossDomainMessenger mockMessenger = new MockCrossDomainMessenger();
        hevm.store(address(gateway), bytes32(eth_erc20_messenger_slot), bytes32(abi.encode(address(mockMessenger))));

        // only call by counterpart
        hevm.expectRevert("only call by counterpart");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeCall(gateway.finalizeDepositETH, (sender, recipient, amount, dataToCall))
        );

        mockMessenger.setXDomainMessageSender(address(counterpartGateway));

        // msg.value mismatch
        hevm.expectRevert("msg.value mismatch");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeCall(gateway.finalizeDepositETH, (sender, recipient, amount, dataToCall))
        );
    }

    function test_finalizeWithdrawETHFailed_succeeds(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, address(this).balance / 2);

        // send some ETH to L2CrossDomainMessenger
        gateway.withdrawETH{value: amount}(amount, 21000);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (sender, recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );

        // counterpart is not L1ETHGateway
        // emit FailedRelayedMessage from L2CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 messengerBalance = address(l2Messenger).balance;
        uint256 recipientBalance = recipient.balance;
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
        hevm.startPrank(AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger)));
        l2Messenger.relayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );
        hevm.stopPrank();
        assertEq(messengerBalance, address(l2Messenger).balance);
        assertEq(recipientBalance, recipient.balance);
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawETH_succeeds(address sender, uint256 amount, bytes memory dataToCall) public {
        address recipient = address(2048);

        amount = bound(amount, 1, address(this).balance / 2);

        // send some ETH to L2CrossDomainMessenger
        gateway.withdrawETH{value: amount}(amount, 21000);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(
            IL2ETHGateway.finalizeDepositETH,
            (sender, address(recipient), amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message
        );

        // emit FinalizeDepositETH from L2ETHGateway
        {
            hevm.expectEmit(true, true, false, true);
            emit IL2ETHGateway.FinalizeDepositETH(sender, address(recipient), amount, dataToCall);
        }

        // emit RelayedMessage from L2CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 messengerBalance = address(l2Messenger).balance;
        uint256 recipientBalance = address(recipient).balance;
        assertBoolEq(false, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
        hevm.startPrank(AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger)));
        l2Messenger.relayMessage(address(counterpartGateway), address(gateway), amount, 0, message);
        hevm.stopPrank();
        assertEq(messengerBalance - amount, address(l2Messenger).balance);
        assertEq(recipientBalance + amount, address(recipient).balance);
        assertBoolEq(true, l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata)));
    }

    function _withdrawETH(bool useRouter, uint256 amount, uint256 gasLimit, uint256 feePerGas) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ETHGateway.finalizeWithdrawETH,
            (address(this), address(this), amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero eth");
            if (useRouter) {
                router.withdrawETH{value: amount}(amount, gasLimit);
            } else {
                gateway.withdrawETH{value: amount}(amount, gasLimit);
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawETH from L2ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL2ETHGateway.WithdrawETH(address(this), address(this), amount, new bytes(0), 0);

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawETH{value: amount + feeToPay}(amount, gasLimit);
            } else {
                gateway.withdrawETH{value: amount + feeToPay}(amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _withdrawETHWithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ETHGateway.finalizeWithdrawETH,
            (address(this), recipient, amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero eth");
            if (useRouter) {
                router.withdrawETH{value: amount}(recipient, amount, gasLimit);
            } else {
                gateway.withdrawETH{value: amount}(recipient, amount, gasLimit);
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawETH from L2ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL2ETHGateway.WithdrawETH(address(this), recipient, amount, new bytes(0), 0);

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawETH{value: amount + feeToPay}(recipient, amount, gasLimit);
            } else {
                gateway.withdrawETH{value: amount + feeToPay}(recipient, amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _withdrawETHWithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL1ETHGateway.finalizeWithdrawETH,
            (address(this), recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("withdraw zero eth");
            if (useRouter) {
                router.withdrawETHAndCall{value: amount}(recipient, amount, dataToCall, gasLimit);
            } else {
                gateway.withdrawETHAndCall{value: amount}(recipient, amount, dataToCall, gasLimit);
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit L2ToL1MessagePasser.AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(gateway),
                    address(counterpartGateway),
                    amount,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit WithdrawETH from L2ETHGateway
            hevm.expectEmit(true, true, false, true);
            emit IL2ETHGateway.WithdrawETH(address(this), recipient, amount, dataToCall, 0);

            uint256 messengerBalance = address(l2Messenger).balance;
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.withdrawETHAndCall{value: amount + feeToPay}(recipient, amount, dataToCall, gasLimit);
            } else {
                gateway.withdrawETHAndCall{value: amount + feeToPay}(recipient, amount, dataToCall, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l2Messenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }
}
