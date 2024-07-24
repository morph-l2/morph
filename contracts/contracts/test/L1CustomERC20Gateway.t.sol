// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {L1CustomERC20Gateway} from "../l1/gateways/L1CustomERC20Gateway.sol";
import {L1GatewayRouter} from "../l1/gateways/L1GatewayRouter.sol";
import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {Constants} from "../libraries/constants/Constants.sol";

contract L1CustomERC20GatewayTest is L1GatewayBaseTest {
    address private counterpartGateway;
    L1CustomERC20Gateway private gateway;
    L1GatewayRouter private router;

    MockERC20 private l1Token;
    MockERC20 private l2Token;

    function setUp() public override {
        super.setUp();
        // Deploy tokens
        l1Token = new MockERC20("Mock L1", "ML1", 18);
        l2Token = new MockERC20("Mock L2", "ML2", 18);

        counterpartGateway = l1CustomERC20Gateway.counterpart();
        gateway = l1CustomERC20Gateway;
        router = l1GatewayRouter;

        // Prepare token balances
        l1Token.mint(address(this), type(uint128).max);
        l1Token.approve(address(gateway), type(uint256).max);
        l1Token.approve(address(router), type(uint256).max);

        hevm.prank(multisig);
        gateway.transferOwnership(address(this));
    }

    function test_initialize_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1), address(1));

        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1CustomERC20Gateway.
        TransparentUpgradeableProxy l1CustomERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1CustomERC20Gateway contract.
        L1CustomERC20Gateway l1CustomERC20GatewayImplTemp = new L1CustomERC20Gateway();

        // Expect revert due to zero router address.
        hevm.expectRevert("zero router address");
        ITransparentUpgradeableProxy(address(l1CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L1CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );

        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L1CustomERC20Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(NON_ZERO_ADDRESS), // _router
                    address(NON_ZERO_ADDRESS) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L1CustomERC20Gateway.initialize,
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

        // Deploy a proxy contract for the L1CustomERC20Gateway.
        TransparentUpgradeableProxy l1CustomERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1CustomERC20Gateway contract.
        L1CustomERC20Gateway l1CustomERC20GatewayImplTemp = new L1CustomERC20Gateway();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1CustomERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1CustomERC20GatewayImplTemp),
            abi.encodeCall(
                L1CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1GatewayRouter), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Cast the proxy contract address to the L1CustomERC20Gateway contract type to call its methods.
        L1CustomERC20Gateway l1CustomERC20GatewayTemp = L1CustomERC20Gateway((address(l1CustomERC20GatewayProxyTemp)));
        hevm.stopPrank();

        // Verify the counterpart, router and messenger are initialized successfully.
        assertEq(l1CustomERC20GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l1CustomERC20GatewayTemp.router(), address(l1GatewayRouter));
        assertEq(l1CustomERC20GatewayTemp.messenger(), address(l1CrossDomainMessenger));
    }

    function test_updateTokenMapping_onlyOwner_reverts(address token1) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token1, token1);
        hevm.stopPrank();

        // l2 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token1, address(0));
    }

    function test_updateTokenMapping_succeeds(address token1, address token2) public {
        hevm.assume(token2 != address(0));
        hevm.assume(token1 != address(l1Token));

        assertEq(gateway.getL2ERC20Address(token1), address(0));
        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.getL2ERC20Address(token1), token2);
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

    function test_dropMessage_succeeds(uint256 amount, address recipient, bytes memory dataToCall) public {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        amount = bound(amount, 1, l1Token.balanceOf(address(this)));
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), address(this), recipient, amount, dataToCall)
        );
        gateway.depositERC20AndCall(address(l1Token), recipient, amount, dataToCall, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC20Gateway.RefundERC20(address(l1Token), address(this), amount);

        uint256 balance = l1Token.balanceOf(address(this));
        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), 0, 0, message);
        assertEq(balance + amount, l1Token.balanceOf(address(this)));
    }

    function test_onDropMessage_beforeDropMessage_reverts() public {
        uint256 amount = 1000;

        // Assign 10 ether to l1CrossDomainMessenger.
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Update the token mapping from l1Token to l2Token.
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        // Deposit some tokens to L1StandardERC20Gateway.
        gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // Create a message with the valid selector and set the sender as address(this).
        bytes memory data = new bytes(0);
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), gateway.getL2ERC20Address(address(l1Token)), address(this), address(this), amount, data)
        );

        hevm.startPrank(address(l1CrossDomainMessenger));

        // Simulate the xDomainMessageSender as the DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(Constants.DROP_XDOMAIN_MESSAGE_SENDER)
        );

        // Expect revert due to nonzero msg.value.
        hevm.expectRevert("nonzero msg.value");
        gateway.onDropMessage{value: 1}(message);

        hevm.stopPrank();
    }

    function test_finalizeWithdrawERC20_counterError_fails(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));
        hevm.assume(recipient != address(gateway));

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        // deposit some token to L1StandardERC20Gateway
        gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), sender, recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );
        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
        l1CrossDomainMessenger.proveAndRelayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawERC20_succeeds(address sender, uint256 amount, bytes memory dataToCall) public {
        address recipient = address(2048);

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        // deposit some token to L1StandardERC20Gateway
        gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), sender, address(recipient), amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );

        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );

        // emit FinalizeWithdrawERC20 from L1StandardERC20Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.FinalizeWithdrawERC20(
                address(l1Token),
                address(l2Token),
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

        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(address(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
        l1CrossDomainMessenger.proveAndRelayMessage(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(gatewayBalance - amount, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance + amount, l1Token.balanceOf(address(recipient)));
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawERC20_beforeFinalizeWithdrawERC20_revert() public {
        address recipient = address(2048);
        address _from = address(counterpartGateway);

        // Assign 10 ether to the required addresses.
        hevm.deal(_from, 10 ether);
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Start simulating the calls as l1CrossDomainMessenger.
        hevm.startPrank(address(l1CrossDomainMessenger));

        // Simulate the xDomainMessageSender as the counterpartGateway
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(_from)
        );

        // Expect revert due to non-zero msg.value.
        hevm.expectRevert("nonzero msg.value");
        gateway.finalizeWithdrawERC20{value: 1}(address(l1Token), address(l2Token), _from, recipient, 1, "");

        // Expect revert due to _l2Token being zero.
        hevm.expectRevert("token address cannot be 0");
        gateway.finalizeWithdrawERC20(address(l1Token), address(0), _from, recipient, 1, "");

        // Expect revert due to token mismatch.
        hevm.expectRevert("l2 token mismatch");
        gateway.finalizeWithdrawERC20(address(l1Token), address(l1Token), _from, recipient, 1, "");
        hevm.stopPrank();
    }

    function _depositERC20(bool useRouter, uint256 amount, uint256 gasLimit, uint256 feePerGas) private {
        amount = bound(amount, 0, l1Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), address(this), address(this), amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l2 token");
        if (useRouter) {
            router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        } else {
            gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            }
        } else {
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
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            }
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(gateway)));
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
        amount = bound(amount, 0, l1Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), address(this), recipient, amount, new bytes(0))
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l2 token");
        if (useRouter) {
            router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        } else {
            gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, amount, gasLimit);
            }
        } else {
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
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, amount, gasLimit);
            } else {
                gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, amount, gasLimit);
            }
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(gateway)));
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
        amount = bound(amount, 0, l1Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (address(l1Token), address(l2Token), address(this), recipient, amount, dataToCall)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l2 token");
        if (useRouter) {
            router.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        } else {
            gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                router.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
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
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC20Gateway.DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                dataToCall,
                0
            );

            uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                router.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }
}
