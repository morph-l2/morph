// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {TransferReentrantToken} from "../mock/tokens/TransferReentrantToken.sol";
import {FeeOnTransferToken} from "../mock/tokens/FeeOnTransferToken.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {L1StandardERC20Gateway} from "../l1/gateways/L1StandardERC20Gateway.sol";
import {Constants} from "../libraries/constants/Constants.sol";

contract L1StandardERC20GatewayTest is L1GatewayBaseTest {
    address public counterpartGateway;
    MockERC20 private l1Token;
    MockERC20 private l2Token;
    TransferReentrantToken private reentrantToken;
    FeeOnTransferToken private feeToken;

    function setUp() public override {
        super.setUp();
        hevm.startPrank(multisig);

        // Deploy tokens
        l1Token = new MockERC20("Mock", "M", 18);
        reentrantToken = new TransferReentrantToken("Reentrant", "R", 18);
        feeToken = new FeeOnTransferToken("Fee", "F", 18);

        counterpartGateway = l1StandardERC20Gateway.counterpart();

        l2Token = MockERC20(l1StandardERC20Gateway.getL2ERC20Address(address(l1Token)));
        l1Token.mint(address(this), type(uint128).max);
        reentrantToken.mint(address(this), type(uint128).max);
        feeToken.mint(address(this), type(uint128).max);

        hevm.stopPrank();

        l1Token.approve(address(l1StandardERC20Gateway), type(uint256).max);
        l1Token.approve(address(l1GatewayRouter), type(uint256).max);
        feeToken.approve(address(l1StandardERC20Gateway), type(uint256).max);
        feeToken.approve(address(l1GatewayRouter), type(uint256).max);
        reentrantToken.approve(address(l1StandardERC20Gateway), type(uint256).max);
        reentrantToken.approve(address(l1GatewayRouter), type(uint256).max);
    }

    function test_initialize_reverts() public {
        // verify the initialize only can be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1StandardERC20Gateway.initialize(address(1), address(1), address(1), address(1), address(1));

        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1StandardERC20Gateway.
        TransparentUpgradeableProxy l1StandardERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1StandardERC20Gateway contract.
        L1StandardERC20Gateway l1StandardERC20GatewayImplTemp = new L1StandardERC20Gateway();

        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(1), // _router
                    address(1), // _messenger
                    address(1), // _l2TokenImplementation
                    address(1) // _l2TokenFactory
                )
            )
        );

        // Expect revert due to zero router address.
        hevm.expectRevert("zero router address");
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(1), // _counterpart
                    address(0), // _router
                    address(1), // _messenger
                    address(1), // _l2TokenImplementation
                    address(1) // _l2TokenFactory
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(1), // _counterpart
                    address(1), // _router
                    address(0), // _messenger
                    address(1), // _l2TokenImplementation
                    address(1) // _l2TokenFactory
                )
            )
        );

        // Expect revert due to zero implementation address.
        hevm.expectRevert("zero implementation hash");
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(1), // _counterpart
                    address(1), // _router
                    address(1), // _messenger
                    address(0), // _l2TokenImplementation
                    address(1) // _l2TokenFactory
                )
            )
        );

        // Expect revert due to zero factory address.
        hevm.expectRevert("zero factory address");
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(1), // _counterpart
                    address(1), // _router
                    address(1), // _messenger
                    address(1), // _l2TokenImplementation
                    address(0) // _l2TokenFactory
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1StandardERC20Gateway.
        TransparentUpgradeableProxy l1StandardERC20GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1StandardERC20Gateway contract.
        L1StandardERC20Gateway l1StandardERC20GatewayImplTemp = new L1StandardERC20Gateway();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxyTemp)).upgradeToAndCall(
            address(l1StandardERC20GatewayImplTemp),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(1), // _counterpart
                    address(2), // _router
                    address(3), // _messenger
                    address(4), // _l2TokenImplementation
                    address(5) // _l2TokenFactory
                )
            )
        );
        hevm.stopPrank();

        // Cast the proxy contract address to the L1StandardERC20Gateway contract type to call its methods.
        L1StandardERC20Gateway l1StandardERC20GatewayTemp = L1StandardERC20Gateway(
            address(l1StandardERC20GatewayProxyTemp)
        );

        // Verify the counterpart, router, messenger, l2TokenImplementation, and l2TokenFactory are initialized successfully.
        assertEq(l1StandardERC20GatewayTemp.counterpart(), address(1));
        assertEq(l1StandardERC20GatewayTemp.router(), address(2));
        assertEq(l1StandardERC20GatewayTemp.messenger(), address(3));
        assertEq(l1StandardERC20GatewayTemp.l2TokenImplementation(), address(4));
        assertEq(l1StandardERC20GatewayTemp.l2TokenFactory(), address(5));
    }

    function test_getL2ERC20Address_succeeds(address l1Address) public {
        hevm.assume(l1Address != address(0));
        assertEq(
            l1StandardERC20Gateway.getL2ERC20Address(l1Address),
            factory.computeL2TokenAddress(address(counterpartGateway), l1Address)
        );
    }

    function test_depositERC20_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositERC20(false, amount, gasLimit, feePerGas);
    }

    function test_routerDepositERC20_succeeds(uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _depositERC20(true, amount, gasLimit, feePerGas);
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

    function test_routerDepositERC20WithRecipient_succeeds(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipient(true, amount, recipient, gasLimit, feePerGas);
    }

    function tes_routerDepositERC20WithRecipientAndCalldata_succeeds(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipientAndCalldata(true, amount, recipient, dataToCall, gasLimit, feePerGas);
    }

    function test_depositReentrantToken_succeeds(uint256 amount) public {
        // should revert, reentrant before transfer
        reentrantToken.setReentrantCall(
            address(l1StandardERC20Gateway),
            0,
            abi.encodeWithSignature("depositERC20(address,uint256,uint256)", address(0), 1, 0),
            true
        );
        amount = bound(amount, 1, reentrantToken.balanceOf(address(this)));
        hevm.expectRevert("ReentrancyGuard: reentrant call");
        l1StandardERC20Gateway.depositERC20(address(reentrantToken), amount, defaultGasLimit);

        // should revert, reentrant after transfer
        reentrantToken.setReentrantCall(
            address(l1StandardERC20Gateway),
            0,
            abi.encodeWithSignature("depositERC20(address,uint256,uint256)", address(0), 1, 0),
            false
        );
        amount = bound(amount, 1, reentrantToken.balanceOf(address(this)));
        hevm.expectRevert("ReentrancyGuard: reentrant call");
        l1StandardERC20Gateway.depositERC20(address(reentrantToken), amount, defaultGasLimit);
    }

    function test_feeOnTransferToken_zeroAmount_fails(uint256 amount) public {
        feeToken.setFeeRate(1e9);
        amount = bound(amount, 1, feeToken.balanceOf(address(this)));
        hevm.expectRevert("deposit zero amount");
        l1StandardERC20Gateway.depositERC20(address(feeToken), amount, defaultGasLimit);
    }

    function test_feeOnTransferToken_succeeds(uint256 amount, uint256 feeRate) public {
        feeRate = bound(feeRate, 0, 1e9 - 1);
        amount = bound(amount, 1e9, feeToken.balanceOf(address(this)));
        feeToken.setFeeRate(feeRate);

        // should succeed, for valid amount
        uint256 balanceBefore = feeToken.balanceOf(address(l1StandardERC20Gateway));
        uint256 fee = (amount * feeRate) / 1e9;
        l1StandardERC20Gateway.depositERC20(address(feeToken), amount, defaultGasLimit);
        uint256 balanceAfter = feeToken.balanceOf(address(l1StandardERC20Gateway));
        assertEq(balanceBefore + amount - fee, balanceAfter);
    }

    function test_dropMessage_succeeds(uint256 amount, address recipient, bytes memory dataToCall) public {
        amount = bound(amount, 1, l1Token.balanceOf(address(this)) / 2);
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                abi.encode(
                    true,
                    abi.encode(dataToCall, abi.encode(l1Token.symbol(), l1Token.name(), l1Token.decimals()))
                )
            )
        );
        l1StandardERC20Gateway.depositERC20AndCall(address(l1Token), recipient, amount, dataToCall, defaultGasLimit);
        l1StandardERC20Gateway.depositERC20AndCall(address(l1Token), recipient, amount, dataToCall, defaultGasLimit);

        // skip message 0 and 1
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 2, 0x3);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 2);
        hevm.stopPrank();

        // drop message 1
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC20Gateway.RefundERC20(address(l1Token), address(this), amount);

        uint256 balance = l1Token.balanceOf(address(this));
        l1CrossDomainMessenger.dropMessage(address(l1StandardERC20Gateway), address(counterpartGateway), 0, 1, message);
        assertEq(balance + amount, l1Token.balanceOf(address(this)));
    }

    function test_finalizeWithdrawERC20_beforeFinalizeWithdrawERC20_reverts() public {
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
        l1StandardERC20Gateway.finalizeWithdrawERC20{value: 1}(
            address(l1Token),
            address(l2Token),
            _from,
            recipient,
            1,
            ""
        );

        // Expect revert due to _l2Token being zero.
        hevm.expectRevert("token address cannot be 0");
        l1StandardERC20Gateway.finalizeWithdrawERC20(address(l1Token), address(0), _from, recipient, 1, "");

        // Expect revert due to token mismatch.
        hevm.expectRevert("l2 token mismatch");
        l1StandardERC20Gateway.finalizeWithdrawERC20(address(l1Token), address(l1Token), _from, recipient, 1, "");
        hevm.stopPrank();
    }

    function test_finalizeWithdrawERC20_counterErr_fails(
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));
        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        address _from = address(uint160(address(counterpartGateway)) + 1);

        // deposit some token to L1StandardERC20Gateway
        l1StandardERC20Gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), _from, recipient, amount, dataToCall)
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, address(l1StandardERC20Gateway), 0, 0, message)
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            _from,
            address(l1StandardERC20Gateway),
            0,
            0,
            message
        );
        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(_xDomainCalldataHash);

        uint256 gatewayBalance = l1Token.balanceOf(address(l1StandardERC20Gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));

        l1CrossDomainMessenger.proveAndRelayMessage(
            _from,
            address(l1StandardERC20Gateway),
            0,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(gatewayBalance, l1Token.balanceOf(address(l1StandardERC20Gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
    }

    function test_finalizeWithdrawERC20_succeeds(address sender, uint256 amount, bytes memory dataToCall) public {
        address recipient = address(2048);
        address _from = address(counterpartGateway);

        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        // deposit some token to L1StandardERC20Gateway
        l1StandardERC20Gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (address(l1Token), address(l2Token), sender, address(recipient), amount, dataToCall)
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, address(l1StandardERC20Gateway), 0, 0, message)
        );
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            _from,
            address(l1StandardERC20Gateway),
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
            emit ICrossDomainMessenger.RelayedMessage(_xDomainCalldataHash);
        }

        uint256 gatewayBalance = l1Token.balanceOf(address(l1StandardERC20Gateway));
        uint256 recipientBalance = l1Token.balanceOf(address(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));

        l1CrossDomainMessenger.proveAndRelayMessage(
            _from,
            address(l1StandardERC20Gateway),
            0,
            0,
            message,
            wdProof,
            wdRoot
        );
        assertEq(gatewayBalance - amount, l1Token.balanceOf(address(l1StandardERC20Gateway)));
        assertEq(recipientBalance + amount, l1Token.balanceOf(address(recipient)));
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash));
    }

    function test_onDropMessage_beforeDropMessage_reverts() public {
        uint256 amount = 1000;

        // Assign 10 ether to l1CrossDomainMessenger.
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Deposit some tokens to L1StandardERC20Gateway.
        l1StandardERC20Gateway.depositERC20(address(l1Token), amount, defaultGasLimit);

        // Create a message with the valid selector and set the sender as address(this).
        bytes memory data = new bytes(0);
        bytes memory message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (
                address(l1Token),
                l1StandardERC20Gateway.getL2ERC20Address(address(l1Token)),
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
        hevm.expectRevert("nonzero msg.value");
        l1StandardERC20Gateway.onDropMessage{value: 1}(message);

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
            (
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                amount,
                abi.encode(
                    true,
                    abi.encode(new bytes(0), abi.encode(l1Token.symbol(), l1Token.name(), l1Token.decimals()))
                )
            )
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            } else {
                l1StandardERC20Gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
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
                    address(l1StandardERC20Gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1StandardERC20Gateway
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

            uint256 l1StandardERC20GatewayBalance = l1Token.balanceOf(address(l1StandardERC20Gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            } else {
                l1StandardERC20Gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(address(l1Token), amount, gasLimit);
            }
            assertEq(amount + l1StandardERC20GatewayBalance, l1Token.balanceOf(address(l1StandardERC20Gateway)));
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
            (
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                abi.encode(
                    true,
                    abi.encode(new bytes(0), abi.encode(l1Token.symbol(), l1Token.name(), l1Token.decimals()))
                )
            )
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
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
                    address(l1StandardERC20Gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1StandardERC20Gateway
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

            uint256 gatewayBalance = l1Token.balanceOf(address(l1StandardERC20Gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
            }
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(l1StandardERC20Gateway)));
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
            (
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                abi.encode(
                    true,
                    abi.encode(dataToCall, abi.encode(l1Token.symbol(), l1Token.name(), l1Token.decimals()))
                )
            )
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
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
                address sender = AddressAliasHelper.applyL1ToL2Alias(address(l1CrossDomainMessenger));
                emit IL1MessageQueue.QueueTransaction(sender, address(l2Messenger), 0, 0, gasLimit, xDomainCalldata);
            }

            // emit SentMessage from L1CrossDomainMessenger
            {
                hevm.expectEmit(true, true, false, true);
                emit ICrossDomainMessenger.SentMessage(
                    address(l1StandardERC20Gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit DepositERC20 from L1StandardERC20Gateway
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

            uint256 gatewayBalance = l1Token.balanceOf(address(l1StandardERC20Gateway));
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            if (useRouter) {
                l1GatewayRouter.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20AndCall{value: feeToPay + EXTRA_VALUE}(
                    address(l1Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(l1StandardERC20Gateway)));
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }
}
