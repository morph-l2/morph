// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL2ETHGateway} from "../L2/gateways/IL2ETHGateway.sol";
import {IL1ETHGateway} from "../L1/gateways/IL1ETHGateway.sol";

contract L1ETHGatewayTest is L1GatewayBaseTest {
    event DepositETH(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data
    );
    event RefundETH(address indexed recipient, uint256 amount);
    event FinalizeWithdrawETH(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data
    );

    address counterpartGateway;

    function setUp() public virtual override {
        super.setUp();
        counterpartGateway = l1ETHGateway.counterpart();
    }

    function testDepositETH(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETH(false, amount, gasLimit, feePerGas);
    }

    function testDepositETHWithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipient(false, amount, recipient, gasLimit, feePerGas);
    }

    function testDepositETHWithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipientAndCalldata(
            false,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testRouterDepositETH(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETH(true, amount, gasLimit, feePerGas);
    }

    function testRouterDepositETHWithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipient(true, amount, recipient, gasLimit, feePerGas);
    }

    function testRouterDepositETHWithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositETHWithRecipientAndCalldata(
            true,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testDropMessage(
        uint256 amount,
        address recipient,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, address(this).balance);
        bytes memory message = abi.encodeWithSelector(
            IL2ETHGateway.finalizeDepositETH.selector,
            address(this),
            recipient,
            amount,
            dataToCall
        );
        l1ETHGateway.depositETHAndCall{value: amount}(
            recipient,
            amount,
            dataToCall,
            defaultGasLimit
        );

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueue.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueue.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // ETH transfer failed, revert
        revertOnReceive = true;
        hevm.expectRevert("ETH transfer failed");
        l1CrossDomainMessenger.dropMessage(
            address(l1ETHGateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit RefundETH(address(this), amount);

        revertOnReceive = false;
        uint256 balance = address(this).balance;
        l1CrossDomainMessenger.dropMessage(
            address(l1ETHGateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );
        assertEq(balance + amount, address(this).balance);
    }

    function testFinalizeWithdrawETHFailed(
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
        bytes memory message = abi.encodeWithSelector(
            IL1ETHGateway.finalizeWithdrawETH.selector,
            sender,
            recipient,
            amount,
            dataToCall
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(
                _from,
                address(l1ETHGateway),
                amount,
                0,
                message
            )
        );

        messageProve(_from, address(l1ETHGateway), amount, 0, message);

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = recipient.balance;
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
        // counterpart is not L2ETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(_xDomainCalldataHash);
        l1CrossDomainMessenger.relayMessage(
            _from,
            address(l1ETHGateway),
            amount,
            0,
            message
        );

        assertEq(messengerBalance, address(l1CrossDomainMessenger).balance);
        assertEq(recipientBalance, recipient.balance);
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
    }

    function testFinalizeWithdrawETH() public {
        address sender = bob;
        uint256 amount = 1000;
        address recipient = address(2048);
        address _from = counterpartGateway;
        // deposit some ETH to L1CrossDomainMessenger
        l1ETHGateway.depositETH{value: amount}(amount, defaultGasLimit);

        // do finalize withdraw eth
        bytes memory message = abi.encodeWithSelector(
            IL1ETHGateway.finalizeWithdrawETH.selector,
            sender,
            recipient,
            amount,
            ""
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(
                _from,
                address(l1ETHGateway),
                amount,
                0,
                message
            )
        );
        messageProve(_from, address(l1ETHGateway), amount, 0, message);

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = recipient.balance;
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
        // counterpart is not L2ETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit FinalizeWithdrawETH(sender, address(recipient), amount, "");
        }

        l1CrossDomainMessenger.relayMessage(
            _from,
            address(l1ETHGateway),
            amount,
            0,
            message
        );

        assertEq(
            messengerBalance - amount,
            address(l1CrossDomainMessenger).balance
        );
        assertEq(recipientBalance + amount, address(recipient).balance);
        assertBoolEq(
            true,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
    }

    function _depositETH(
        bool useRouter,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, address(this).balance / 2);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ETHGateway.finalizeDepositETH.selector,
            address(this),
            address(this),
            amount,
            new bytes(0)
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
                address sender = AddressAliasHelper.applyL1ToL2Alias(
                    address(l1CrossDomainMessenger)
                );
                emit QueueTransaction(
                    sender,
                    address(l2Messenger),
                    0,
                    0,
                    gasLimit,
                    xDomainCalldata
                );
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
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
            emit DepositETH(address(this), address(this), amount, new bytes(0));

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositETH{
                    value: amount + feeToPay + extraValue
                }(amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount + feeToPay + extraValue}(
                    amount,
                    gasLimit
                );
            }
            assertEq(
                amount + messengerBalance,
                address(l1CrossDomainMessenger).balance
            );
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ETHGateway.finalizeDepositETH.selector,
            address(this),
            recipient,
            amount,
            new bytes(0)
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
                l1GatewayRouter.depositETH{value: amount}(
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                l1ETHGateway.depositETH{value: amount}(
                    recipient,
                    amount,
                    gasLimit
                );
            }
        } else {
            // emit QueueTransaction from L1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(
                    address(l1CrossDomainMessenger)
                );
                emit QueueTransaction(
                    sender,
                    address(l2Messenger),
                    0,
                    0,
                    gasLimit,
                    xDomainCalldata
                );
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
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
            emit DepositETH(address(this), recipient, amount, new bytes(0));

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositETH{
                    value: amount + feeToPay + extraValue
                }(recipient, amount, gasLimit);
            } else {
                l1ETHGateway.depositETH{value: amount + feeToPay + extraValue}(
                    recipient,
                    amount,
                    gasLimit
                );
            }
            assertEq(
                amount + messengerBalance,
                address(l1CrossDomainMessenger).balance
            );
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ETHGateway.finalizeDepositETH.selector,
            address(this),
            recipient,
            amount,
            dataToCall
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
                l1GatewayRouter.depositETHAndCall{value: amount}(
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                l1ETHGateway.depositETHAndCall{value: amount}(
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
            // emit QueueTransaction from L1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(
                    address(l1CrossDomainMessenger)
                );
                emit QueueTransaction(
                    sender,
                    address(l2Messenger),
                    0,
                    0,
                    gasLimit,
                    xDomainCalldata
                );
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit SentMessage(
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
            emit DepositETH(address(this), recipient, amount, dataToCall);

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositETHAndCall{
                    value: amount + feeToPay + extraValue
                }(recipient, amount, dataToCall, gasLimit);
            } else {
                l1ETHGateway.depositETHAndCall{
                    value: amount + feeToPay + extraValue
                }(recipient, amount, dataToCall, gasLimit);
            }
            assertEq(
                amount + messengerBalance,
                address(l1CrossDomainMessenger).balance
            );
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
        }
    }
}
