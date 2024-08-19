// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC1155} from "@rari-capital/solmate/src/test/utils/mocks/MockERC1155.sol";
import {ERC1155TokenReceiver} from "@rari-capital/solmate/src/tokens/ERC1155.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {L1ERC1155Gateway} from "../l1/gateways/L1ERC1155Gateway.sol";
import {IL1ERC1155Gateway} from "../l1/gateways/IL1ERC1155Gateway.sol";
import {IL2ERC1155Gateway} from "../l2/gateways/IL2ERC1155Gateway.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {Constants} from "../libraries/constants/Constants.sol";

contract L1ERC1155GatewayTest is L1GatewayBaseTest, ERC1155TokenReceiver {
    uint256 private constant TOKEN_COUNT = 100;
    uint256 private constant MAX_TOKEN_BALANCE = 1000000000;

    L1ERC1155Gateway private gateway;

    address private counterpartGateway;

    MockERC1155 private l1Token;
    MockERC1155 private l2Token;

    function setUp() public override {
        super.setUp();
        _deployERC1155();

        // Deploy tokens
        l1Token = new MockERC1155();
        l2Token = new MockERC1155();

        counterpartGateway = l1ERC1155Gateway.counterpart();
        gateway = l1ERC1155Gateway;
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        // Prepare token balances
        for (uint256 i = 0; i < TOKEN_COUNT; i++) {
            l1Token.mint(address(this), i, MAX_TOKEN_BALANCE, "");
        }
        l1Token.setApprovalForAll(address(gateway), true);
    }

    function test_initialize_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1));

        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1ERC1155Gateway.
        TransparentUpgradeableProxy l1ERC1155GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1ERC1155Gateway contract.
        L1ERC1155Gateway l1ERC1155GatewayImpl = new L1ERC1155Gateway();

        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1ERC1155GatewayProxy)).upgradeToAndCall(
            address(l1ERC1155GatewayImpl),
            abi.encodeCall(
                L1ERC1155Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1ERC1155GatewayProxy)).upgradeToAndCall(
            address(l1ERC1155GatewayImpl),
            abi.encodeCall(
                L1ERC1155Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0) // _messenger
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for the L1ERC1155Gateway.
        TransparentUpgradeableProxy l1ERC1155GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1ERC1155Gateway contract.
        L1ERC1155Gateway l1ERC1155GatewayImplTemp = new L1ERC1155Gateway();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1ERC1155GatewayProxyTemp)).upgradeToAndCall(
            address(l1ERC1155GatewayImplTemp),
            abi.encodeCall(
                L1ERC1155Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Cast the proxy contract address to the L1ERC1155Gateway contract type to call its methods.
        L1ERC1155Gateway l1ERC1155GatewayTemp = L1ERC1155Gateway(address(l1ERC1155GatewayProxyTemp));
        hevm.stopPrank();

        // Verify the counterpart and messenger are initialized successfully.
        assertEq(l1ERC1155GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l1ERC1155GatewayTemp.messenger(), address(l1CrossDomainMessenger));
    }

    function test_depositERC1155_succeeds(uint256 tokenId, uint256 amount, uint256 gasLimit, uint256 feePerGas) public {
        _testDepositERC1155(tokenId, amount, gasLimit, feePerGas);
    }

    function test_depositERC1155WithRecipient_succeeds(
        uint256 tokenId,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testDepositERC1155WithRecipient(tokenId, amount, recipient, gasLimit, feePerGas);
    }

    function test_batchDepositERC1155_succeeds(
        uint256 tokenCount,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC1155(tokenCount, amount, gasLimit, feePerGas);
    }

    function test_batchDepositERC1155WithRecipient_succeeds(
        uint256 tokenCount,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC1155WithRecipient(tokenCount, amount, recipient, gasLimit, feePerGas);
    }

    function test_dropMessage_succeeds(uint256 tokenId, uint256 amount) public {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeDepositERC1155,
            (address(l1Token), address(l2Token), address(this), address(this), tokenId, amount)
        );
        gateway.depositERC1155(address(l1Token), tokenId, amount, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC1155Gateway.RefundERC1155(address(l1Token), address(this), tokenId, amount);

        uint256 balance = l1Token.balanceOf(address(this), tokenId);
        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), 0, 0, message);
        assertEq(balance + amount, l1Token.balanceOf(address(this), tokenId));
    }

    function test_dropMessageBatch_succeeds(uint256 tokenCount, uint256 amount) public {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        uint256[] memory _amounts = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeBatchDepositERC1155,
            (address(l1Token), address(l2Token), address(this), address(this), _tokenIds, _amounts)
        );
        gateway.batchDepositERC1155(address(l1Token), _tokenIds, _amounts, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC1155Gateway.BatchRefundERC1155(address(l1Token), address(this), _tokenIds, _amounts);

        uint256[] memory balances = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            balances[i] = l1Token.balanceOf(address(this), _tokenIds[i]);
        }
        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), 0, 0, message);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(balances[i] + _amounts[i], l1Token.balanceOf(address(this), _tokenIds[i]));
        }
    }

    function test_onDropMessage_revert(uint256 tokenId, uint256 amount) external {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        bytes memory message = "message";

        // Assign 10 ether to the contract and l1CrossDomainMessenger.
        hevm.deal(address(this), 10 ether);
        hevm.deal(address(l1CrossDomainMessenger), 10 ether);

        // Verify onlyInDropContext modifier. Expect revert when msg.sender is not the messenger.
        hevm.expectRevert("only messenger can call");
        gateway.onDropMessage(message);

        // Expect revert when xDomainMessageSender is not DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.prank(address(l1CrossDomainMessenger));
        hevm.expectRevert("only called in drop context");
        gateway.onDropMessage(message);

        // Simulate xDomainMessageSender returning DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(Constants.DROP_XDOMAIN_MESSAGE_SENDER)
        );

        // Update msg.sender to l1CrossDomainMessenger.
        hevm.startPrank(address(l1CrossDomainMessenger));

        // Expect revert when msg.value != 0.
        hevm.expectRevert("nonzero msg.value");
        gateway.onDropMessage{value: amount}(message);

        // Expect revert when selector is invalid.
        hevm.expectRevert("invalid selector");
        gateway.onDropMessage(message);

        hevm.stopPrank();
    }

    function test_onDropMessage_succeeds(uint256 tokenId, uint256 amount) external {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);

        gateway.depositERC1155(address(l1Token), tokenId, amount, defaultGasLimit);

        // Create a message with the valid selector and set the sender as address(this).
        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeDepositERC1155,
            (address(l1Token), address(l2Token), address(this), address(this), tokenId, amount)
        );

        // Update msg.sender to l1CrossDomainMessenger.
        hevm.startPrank(address(l1CrossDomainMessenger));

        // Simulate xDomainMessageSender returning DROP_XDOMAIN_MESSAGE_SENDER.
        hevm.mockCall(
            address(l1CrossDomainMessenger),
            abi.encodeWithSignature("xDomainMessageSender()"),
            abi.encode(Constants.DROP_XDOMAIN_MESSAGE_SENDER)
        );

        // Expect RefundERC1155 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC1155Gateway.RefundERC1155(address(l1Token), address(this), tokenId, amount);

        gateway.onDropMessage(message);
        hevm.stopPrank();
    }

    function test_updateTokenMapping_revert(address token1) public {
        // Simulate a call by a non-owner address and expect a revert.
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(address(1));
        gateway.updateTokenMapping(token1, token1);

        // Expect a revert when the L2 token address is zero.
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token1, address(0));
    }

    function test_updateTokenMapping_succeeds(address token1, address token2) public {
        // Assume token2 is not a zero address.
        hevm.assume(token2 != address(0));

        // Verify the initial mapping for token1 is zero.
        assertEq(gateway.tokenMapping(token1), address(0));

        // Expect the UpdateTokenMapping event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit L1ERC1155Gateway.UpdateTokenMapping(token1, address(0), token2);

        // Update the token mapping from token1 to token2 and verify the mapping.
        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.tokenMapping(token1), token2);
    }

    function test_finalizeWithdrawERC1155_counterErr_fails(
        address sender,
        address recipient,
        uint256 tokenId,
        uint256 amount
    ) public {
        hevm.assume(recipient != address(0));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.depositERC1155(address(l1Token), tokenId, amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC1155Gateway.finalizeWithdrawERC1155,
            (address(l1Token), address(l2Token), sender, recipient, tokenId, amount)
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

        uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
        uint256 recipientBalance = l1Token.balanceOf(recipient, tokenId);
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
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway), tokenId));
        assertEq(recipientBalance, l1Token.balanceOf(recipient, tokenId));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeWithdrawERC1155_succeeds(
        address sender,
        address recipient,
        uint256 tokenId,
        uint256 amount
    ) public {
        uint256 size;
        assembly {
            size := extcodesize(recipient)
        }
        hevm.assume(size == 0);
        hevm.assume(recipient != address(0));

        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.depositERC1155(address(l1Token), tokenId, amount, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC1155Gateway.finalizeWithdrawERC1155,
            (address(l1Token), address(l2Token), sender, recipient, tokenId, amount)
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
        // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC1155Gateway.FinalizeWithdrawERC1155(
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                tokenId,
                amount
            );
        }

        // emit RelayedMessage from L1CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
        uint256 recipientBalance = l1Token.balanceOf(recipient, tokenId);
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
        assertEq(gatewayBalance - amount, l1Token.balanceOf(address(gateway), tokenId));
        assertEq(recipientBalance + amount, l1Token.balanceOf(recipient, tokenId));
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeBatchWithdrawERC1155_counterErr_fails(
        address sender,
        address recipient,
        uint256 tokenCount,
        uint256 amount
    ) public {
        hevm.assume(recipient != address(0));
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        uint256[] memory _tokenIds = new uint256[](tokenCount);
        uint256[] memory _amounts = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.batchDepositERC1155(address(l1Token), _tokenIds, _amounts, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC1155Gateway.finalizeBatchWithdrawERC1155,
            (address(l1Token), address(l2Token), sender, recipient, _tokenIds, _amounts)
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

        uint256[] memory gatewayBalances = new uint256[](tokenCount);
        uint256[] memory recipientBalances = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
            recipientBalances[i] = l1Token.balanceOf(recipient, i);
        }
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
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(gatewayBalances[i], l1Token.balanceOf(address(gateway), i));
            assertEq(recipientBalances[i], l1Token.balanceOf(recipient, i));
        }
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeBatchWithdrawERC1155_succeeds(
        address sender,
        address recipient,
        uint256 tokenCount,
        uint256 amount
    ) public {
        uint256 size;
        assembly {
            size := extcodesize(recipient)
        }
        hevm.assume(size == 0);
        hevm.assume(recipient != address(0));

        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        uint256[] memory _tokenIds = new uint256[](tokenCount);
        uint256[] memory _amounts = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.batchDepositERC1155(address(l1Token), _tokenIds, _amounts, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC1155Gateway.finalizeBatchWithdrawERC1155,
            (address(l1Token), address(l2Token), sender, recipient, _tokenIds, _amounts)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );
        _msgRelay(xDomainCalldata, message, sender, recipient, _tokenIds, _amounts, tokenCount);
    }

    function _msgRelay(
        bytes memory xDomainCalldata,
        bytes memory message,
        address sender,
        address recipient,
        uint256[] memory _tokenIds,
        uint256[] memory _amounts,
        uint256 tokenCount
    ) internal {
        (bytes32[32] memory wdProof, bytes32 wdRoot) = messageProveAndRelayPrepare(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );
        // emit FinalizeBatchWithdrawERC1155 from L1ERC1155Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC1155Gateway.FinalizeBatchWithdrawERC1155(
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                _tokenIds,
                _amounts
            );
        }

        // emit RelayedMessage from L1CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256[] memory gatewayBalances = new uint256[](tokenCount);
        uint256[] memory recipientBalances = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
            recipientBalances[i] = l1Token.balanceOf(recipient, i);
        }
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

        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
        _tokenCheck(tokenCount, gatewayBalances, recipientBalances, _amounts, recipient);
    }

    function _tokenCheck(
        uint256 tokenCount,
        uint256[] memory gatewayBalances,
        uint256[] memory recipientBalances,
        uint256[] memory _amounts,
        address recipient
    ) internal {
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(gatewayBalances[i] - _amounts[i], l1Token.balanceOf(address(gateway), i));
            assertEq(recipientBalances[i] + _amounts[i], l1Token.balanceOf(recipient, i));
        }
    }

    function _testDepositERC1155(uint256 tokenId, uint256 amount, uint256 gasLimit, uint256 feePerGas) internal {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 0, MAX_TOKEN_BALANCE);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeDepositERC1155,
            (address(l1Token), address(l2Token), address(this), address(this), tokenId, amount)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            gateway.depositERC1155{value: feeToPay + EXTRA_VALUE}(address(l1Token), tokenId, amount, gasLimit);
        } else {
            hevm.expectRevert("no corresponding l2 token");
            gateway.depositERC1155(address(l1Token), tokenId, amount, gasLimit);

            gateway.updateTokenMapping(address(l1Token), address(l2Token));
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
                    address(gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC1155Gateway.DepositERC1155(
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                tokenId,
                amount
            );

            uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            gateway.depositERC1155{value: feeToPay + EXTRA_VALUE}(address(l1Token), tokenId, amount, gasLimit);
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(gateway), tokenId));
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _testDepositERC1155WithRecipient(
        uint256 tokenId,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 0, MAX_TOKEN_BALANCE);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeDepositERC1155,
            (address(l1Token), address(l2Token), address(this), recipient, tokenId, amount)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            gateway.depositERC1155{value: feeToPay + EXTRA_VALUE}(
                address(l1Token),
                recipient,
                tokenId,
                amount,
                gasLimit
            );
        } else {
            hevm.expectRevert("no corresponding l2 token");
            gateway.depositERC1155(address(l1Token), tokenId, amount, gasLimit);

            gateway.updateTokenMapping(address(l1Token), address(l2Token));
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
                    address(gateway),
                    address(counterpartGateway),
                    0,
                    0,
                    gasLimit,
                    message
                );
            }

            // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC1155Gateway.DepositERC1155(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                tokenId,
                amount
            );

            uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
            gateway.depositERC1155{value: feeToPay + EXTRA_VALUE}(
                address(l1Token),
                recipient,
                tokenId,
                amount,
                gasLimit
            );
            assertEq(amount + gatewayBalance, l1Token.balanceOf(address(gateway), tokenId));
            assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
            assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        }
    }

    function _testBatchDepositERC1155(
        uint256 tokenCount,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        uint256[] memory _amounts = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC1155(address(l1Token), new uint256[](0), new uint256[](0), gasLimit);

        hevm.expectRevert("length mismatch");
        gateway.batchDepositERC1155(address(l1Token), new uint256[](1), new uint256[](0), gasLimit);

        hevm.expectRevert("deposit zero amount");
        gateway.batchDepositERC1155(address(l1Token), _tokenIds, new uint256[](tokenCount), gasLimit);

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC1155(address(l1Token), _tokenIds, _amounts, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeBatchDepositERC1155,
            (address(l1Token), address(l2Token), address(this), address(this), _tokenIds, _amounts)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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
                address(gateway),
                address(counterpartGateway),
                0,
                0,
                gasLimit,
                message
            );
        }

        // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC1155Gateway.BatchDepositERC1155(
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            _tokenIds,
            _amounts
        );

        uint256[] memory gatewayBalances = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
        }
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.batchDepositERC1155{value: feeToPay + EXTRA_VALUE}(address(l1Token), _tokenIds, _amounts, gasLimit);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(gatewayBalances[i] + amount, l1Token.balanceOf(address(gateway), i));
        }
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }

    function _testBatchDepositERC1155WithRecipient(
        uint256 tokenCount,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        amount = bound(amount, 1, MAX_TOKEN_BALANCE);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        uint256[] memory _amounts = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC1155(address(l1Token), recipient, new uint256[](0), new uint256[](0), gasLimit);

        hevm.expectRevert("length mismatch");
        gateway.batchDepositERC1155(address(l1Token), recipient, new uint256[](1), new uint256[](0), gasLimit);

        hevm.expectRevert("deposit zero amount");
        gateway.batchDepositERC1155(address(l1Token), recipient, _tokenIds, new uint256[](tokenCount), gasLimit);

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC1155(address(l1Token), recipient, _tokenIds, _amounts, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC1155Gateway.finalizeBatchDepositERC1155,
            (address(l1Token), address(l2Token), address(this), recipient, _tokenIds, _amounts)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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
                address(gateway),
                address(counterpartGateway),
                0,
                0,
                gasLimit,
                message
            );
        }

        // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC1155Gateway.BatchDepositERC1155(
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            _tokenIds,
            _amounts
        );

        uint256[] memory gatewayBalances = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
        }
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.batchDepositERC1155{value: feeToPay + EXTRA_VALUE}(
            address(l1Token),
            recipient,
            _tokenIds,
            _amounts,
            gasLimit
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(gatewayBalances[i] + amount, l1Token.balanceOf(address(gateway), i));
        }
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }
}
