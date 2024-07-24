// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {WETH} from "@rari-capital/solmate/src/tokens/WETH.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {L1GatewayRouter} from "../l1/gateways/L1GatewayRouter.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {L1WETHGateway} from "../l1/gateways/L1WETHGateway.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {Constants} from "../libraries/constants/Constants.sol";

contract L1WETHGatewayTest is L1GatewayBaseTest {
    WETH private l1weth;
    WETH private l2weth;

    L1WETHGateway private gateway;
    L1GatewayRouter private router;

    address private counterpartGateway;

    function setUp() public override {
        super.setUp();

        // Deploy tokens
        l1weth = new WETH();
        l2weth = new WETH();
        _deployWETH(address(l1weth), address(l2weth));

        counterpartGateway = l1WETHGateway.counterpart();
        gateway = l1WETHGateway;
        router = l1GatewayRouter;

        hevm.prank(multisig);
        router.setDefaultERC20Gateway(address(gateway));

        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        l1weth.deposit{value: address(this).balance / 2}();
        l1weth.approve(address(gateway), type(uint256).max);
        l1weth.approve(address(router), type(uint256).max);
    }

    function test_initialize_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1), address(1));

        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1WETHGateway.
        TransparentUpgradeableProxy l1WETHGatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1WETHGateway contract.
        L1WETHGateway l1WETHGatewayImplTemp = new L1WETHGateway(address(l1weth), address(l2weth));

        // Expect revert due to zero router address.
        hevm.expectRevert("zero router address");
        ITransparentUpgradeableProxy(address(l1WETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1WETHGatewayImplTemp),
            abi.encodeCall(
                L1WETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );

        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1WETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1WETHGatewayImplTemp),
            abi.encodeCall(
                L1WETHGateway.initialize,
                (
                    address(0), // _counterpart
                    address(NON_ZERO_ADDRESS), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1WETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1WETHGatewayImplTemp),
            abi.encodeCall(
                L1WETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(NON_ZERO_ADDRESS), // _router
                    address(0) // _messenger
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for the L1WETHGateway.
        TransparentUpgradeableProxy l1WETHGatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1WETHGateway contract.
        L1WETHGateway l1WETHGatewayImplTemp = new L1WETHGateway(address(l1weth), address(l2weth));

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1WETHGatewayProxyTemp)).upgradeToAndCall(
            address(l1WETHGatewayImplTemp),
            abi.encodeCall(
                L1WETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1GatewayRouter), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Cast the proxy contract address to the L1WETHGateway contract type to call its methods.
        L1WETHGateway l1WETHGatewayTemp = L1WETHGateway(payable(address(l1WETHGatewayProxyTemp)));
        hevm.stopPrank();

        // Verify the counterpart, router and messenger are initialized successfully.
        assertEq(l1WETHGatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l1WETHGatewayTemp.router(), address(l1GatewayRouter));
        assertEq(l1WETHGatewayTemp.messenger(), address(l1CrossDomainMessenger));
    }

    function test_directTransferETH_onlyWeth_fails(uint256 amount) public {
        amount = bound(amount, 0, address(this).balance);
        // solhint-disable-next-line avoid-low-level-calls
        (bool success, bytes memory result) = address(gateway).call{value: amount}("");
        assertBoolEq(success, false);
        assertEq(string(result), string(abi.encodeWithSignature("Error(string)", "only WETH")));
    }

    function test_depositERC20_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositERC20(false, amount, gasLimit, feePerGas);
    }

    function test_depositERC20WithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipient(false, amount, recipient, gasLimit, feePerGas);
    }

    function test_depositERC20WithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipientAndCalldata(false, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_routerDepositERC20_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositERC20(true, amount, gasLimit, feePerGas);
    }

    function test_routerDepositERC20WithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipient(true, amount, recipient, gasLimit, feePerGas);
    }

    function test_routerDepositERC20WithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipientAndCalldata(true, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_dropMessage_succeeds(uint256 amount, address recipient, bytes memory dataToCall) public {
        amount = bound(amount, 1, l1weth.balanceOf(address(this)));
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1weth), address(l2weth), address(this), recipient, amount, dataToCall)
        );
        gateway.depositERC20AndCall(address(l1weth), recipient, amount, dataToCall, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC20Gateway.RefundERC20(address(l1weth), address(this), amount);

        uint256 balance = l1weth.balanceOf(address(this));
        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), amount, 0, message);
        assertEq(balance + amount, l1weth.balanceOf(address(this)));
    }

    function test_finalizeWithdrawERC20_counterErr_fails(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));
        hevm.assume(recipient != address(gateway));

        amount = bound(amount, 1, l1weth.balanceOf(address(this)));

        // deposit some WETH to L1CrossDomainMessenger
        gateway.depositERC20(address(l1weth), amount, defaultGasLimit);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1weth), address(l2weth), sender, recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message
        );
        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = l1weth.balanceOf(recipient);
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));

        l1CrossDomainMessenger.proveAndRelayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            amount,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(messengerBalance, address(l1CrossDomainMessenger).balance);
        assertEq(recipientBalance, l1weth.balanceOf(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawERC20_succeeds(address sender, uint256 amount, bytes memory dataToCall) public {
        address recipient = address(2048);

        amount = bound(amount, 1, l1weth.balanceOf(address(this)));

        // deposit some WETH to L1CrossDomainMessenger
        gateway.depositERC20(address(l1weth), amount, defaultGasLimit);

        // do finalize withdraw eth
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1weth), address(l2weth), sender, address(recipient), amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message
        );
        // emit FinalizeWithdrawERC20 from L1WETHGateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.FinalizeWithdrawERC20(
                address(l1weth),
                address(l2weth),
                sender,
                address(recipient),
                amount,
                dataToCall
            );
        }

        // emit RelayedMessage from L1CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
        uint256 recipientBalance = l1weth.balanceOf(address(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));

        l1CrossDomainMessenger.proveAndRelayMessage(
            address(counterpartGateway),
            address(gateway),
            amount,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(messengerBalance - amount, address(l1CrossDomainMessenger).balance);
        assertEq(recipientBalance + amount, l1weth.balanceOf(address(recipient)));
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawERC20_beforeFinalizeWithdrawERC20_reverts() public {
        address recipient = address(2048);
        address _from = address(counterpartGateway);

        // Assign 10 ether to the required address.
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Start simulating the calls as l1CrossDomainMessenger.
        hevm.startPrank(address(l1CrossDomainMessenger));

        // Simulate the xDomainMessageSender as the counterpartGateway.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(_from)
        );

        // Expect revert due to msg.value mismatch.
        hevm.expectRevert("msg.value mismatch");
        l1WETHGateway.finalizeWithdrawERC20{value: 1}(address(l1weth), address(l2weth), _from, recipient, 2, "");

        // Expect revert due to l1 token is not WETH.
        hevm.expectRevert("l1 token not WETH");
        l1WETHGateway.finalizeWithdrawERC20{value: 1}(address(1), address(l2weth), _from, recipient, 1, "");

        // Expect revert due to l2 token is not WETH.
        hevm.expectRevert("l2 token not WETH");
        l1WETHGateway.finalizeWithdrawERC20{value: 1}(address(l1weth), address(1), _from, recipient, 1, "");

        hevm.stopPrank();
    }

    function test_onDropMessage_beforeDropMessage_reverts() public {
        uint256 amount = 1000;

        // Assign 10 ether to l1CrossDomainMessenger.
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Deposit some tokens to L1WETHGateway.
        l1WETHGateway.depositERC20(address(l1weth), amount, defaultGasLimit);

        // Create a message with the valid selector and set the sender as address(this).
        bytes memory data = new bytes(0);
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (
                address(l1weth),
                l1WETHGateway.getL2ERC20Address(address(l1weth)),
                address(this),
                address(this),
                amount,
                data
            )
        );

        hevm.startPrank(address(l1CrossDomainMessenger));

        // Simulate the xDomainMessageSender as the DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(Constants.DROP_XDOMAIN_MESSAGE_SENDER)
        );

        // Expect revert due to non-zero msg.value.
        hevm.expectRevert("msg.value mismatch");
        l1WETHGateway.onDropMessage{value: 1}(message);

        // Update message to test token not being WETH.
        message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(1), address(l2weth), address(this), address(this), amount, data)
        );

        // Expect revert due to token is not WETH.
        hevm.expectRevert("token not WETH");
        l1WETHGateway.onDropMessage{value: amount}(message);

        hevm.stopPrank();
    }

    function _depositERC20(bool useRouter, uint256 amount, uint256 gasLimit, uint256 feePerGas) private {
        amount = bound(amount, 0, l1weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1weth), address(l2weth), address(this), address(this), amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), amount, gasLimit);
            }
        } else {
            // token is not l1WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.depositERC20(address(l2weth), amount, gasLimit);

            // emit QueueTransaction from L1l1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
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

            // emit DepositERC20 from L1WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                address(this),
                amount,
                new bytes(0),
                0
            );

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l1CrossDomainMessenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _depositERC20WithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l1weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1weth), address(l2weth), address(this), recipient, amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), recipient, amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), recipient, amount, gasLimit);
            }
        } else {
            // token is not l1WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.depositERC20(address(l2weth), recipient, amount, gasLimit);

            // emit QueueTransaction from L1l1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
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

            // emit DepositERC20 from L1WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                recipient,
                amount,
                new bytes(0),
                0
            );

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), recipient, amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1weth), recipient, amount, gasLimit);
            }
            assertEq(amount + messengerBalance, address(l1CrossDomainMessenger).balance);
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _depositERC20WithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l1weth.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1weth), address(l2weth), address(this), recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            amount,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
            // token is not l1WETH
            hevm.expectRevert("only WETH is allowed");
            gateway.depositERC20AndCall(address(l2weth), recipient, amount, dataToCall, gasLimit);

            // emit QueueTransaction from L1l1MessageQueue
            {
                hevm.expectEmit(true, true, false, true);
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
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

            // emit DepositERC20 from L1WETHGateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1weth),
                address(l2weth),
                address(this),
                recipient,
                amount,
                dataToCall,
                0
            );

            uint256 messengerBalance = address(l1CrossDomainMessenger).balance;
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1weth),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1weth),
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
