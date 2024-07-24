// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC721} from "@rari-capital/solmate/src/test/utils/mocks/MockERC721.sol";
import {ERC721TokenReceiver} from "@rari-capital/solmate/src/tokens/ERC721.sol";
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {IL2ERC721Gateway} from "../l2/gateways/IL2ERC721Gateway.sol";
import {IL1ERC721Gateway} from "../l1/gateways/IL1ERC721Gateway.sol";
import {L1ERC721Gateway} from "../l1/gateways/L1ERC721Gateway.sol";
import {IL1MessageQueue} from "../l1/rollup/IL1MessageQueue.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";

contract L1ERC721GatewayTest is L1GatewayBaseTest, ERC721TokenReceiver {
    uint256 private constant TOKEN_COUNT = 100;
    L1ERC721Gateway private gateway;
    address private counterpartGateway;
    MockERC721 private l1Token;
    MockERC721 private l2Token;

    function setUp() public override {
        super.setUp();
        _deployERC721();

        // Deploy tokens
        l1Token = new MockERC721("Mock L1", "ML1");
        l2Token = new MockERC721("Mock L2", "ML1");

        counterpartGateway = l1ERC721Gateway.counterpart();
        gateway = l1ERC721Gateway;
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));
        // Prepare token balances
        for (uint256 i = 0; i < TOKEN_COUNT; i++) {
            l1Token.mint(address(this), i);
        }
        l1Token.setApprovalForAll(address(gateway), true);
    }

    function test_initialize_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1));

        hevm.startPrank(multisig);

        // Deploy a proxy contract for L1ERC721Gateway.
        TransparentUpgradeableProxy l1ERC721GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1ERC721Gateway contract.
        L1ERC721Gateway l1ERC721GatewayImpl = new L1ERC721Gateway();

        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l1ERC721GatewayProxy)).upgradeToAndCall(
            address(l1ERC721GatewayImpl),
            abi.encodeCall(
                L1ERC721Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l1ERC721GatewayProxy)).upgradeToAndCall(
            address(l1ERC721GatewayImpl),
            abi.encodeCall(
                L1ERC721Gateway.initialize,
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

        // Deploy a proxy contract for the L1ERC721Gateway.
        TransparentUpgradeableProxy l1ERC721GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1ERC721Gateway contract.
        L1ERC721Gateway l1ERC721GatewayImplTemp = new L1ERC721Gateway();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1ERC721GatewayProxyTemp)).upgradeToAndCall(
            address(l1ERC721GatewayImplTemp),
            abi.encodeCall(
                L1ERC721Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        // Cast the proxy contract address to the L1ERC721Gateway contract type to call its methods.
        L1ERC721Gateway l1ERC721GatewayTemp = L1ERC721Gateway(address(l1ERC721GatewayProxyTemp));
        hevm.stopPrank();

        // Verify the counterpart and messenger are initialized successfully.
        assertEq(l1ERC721GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l1ERC721GatewayTemp.messenger(), address(l1CrossDomainMessenger));
    }

    function test_updateTokenMapping_onlyOwner_fails(address token1) public {
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

        assertEq(gateway.tokenMapping(token1), address(0));

        // Expect UpdateTokenMapping event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit L1ERC721Gateway.UpdateTokenMapping(token1, address(0), token2);

        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.tokenMapping(token1), token2);
    }

    function test_depositERC721_succeeds(uint256 tokenId, uint256 gasLimit, uint256 feePerGas) public {
        _testDepositERC721(tokenId, gasLimit, feePerGas);
    }

    function test_depositERC721WithRecipient_succeeds(
        uint256 tokenId,
        address to,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testDepositERC721WithRecipient(tokenId, to, gasLimit, feePerGas);
    }

    function test_batchDepositERC721WithGateway_succeeds(
        uint256 tokenCount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC721(tokenCount, gasLimit, feePerGas);
    }

    /// @dev batch deposit erc721 with recipient
    function test_batchDepositERC721WithGateway_succeeds(
        uint256 tokenCount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC721WithRecipient(tokenCount, recipient, gasLimit, feePerGas);
    }

    function test_dropMessage_succeeds(uint256 tokenId) public {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeDepositERC721,
            (address(l1Token), address(l2Token), address(this), address(this), tokenId)
        );
        gateway.depositERC721(address(l1Token), tokenId, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC721Gateway.RefundERC721(address(l1Token), address(this), tokenId);

        assertEq(l1Token.ownerOf(tokenId), address(gateway));
        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), 0, 0, message);
        assertEq(l1Token.ownerOf(tokenId), address(this));
    }

    function test_dropMessageBatch_succeeds(uint256 tokenCount) public {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeBatchDepositERC721,
            (address(l1Token), address(l2Token), address(this), address(this), _tokenIds)
        );
        gateway.batchDepositERC721(address(l1Token), _tokenIds, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit IL1ERC721Gateway.BatchRefundERC721(address(l1Token), address(this), _tokenIds);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(_tokenIds[i]), address(gateway));
        }

        l1CrossDomainMessenger.dropMessage(address(gateway), address(counterpartGateway), 0, 0, message);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(_tokenIds[i]), address(this));
        }
    }

    function test_finalizeWithdrawERC721_counterError_fails(address sender, address recipient, uint256 tokenId) public {
        hevm.assume(recipient != address(0));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.depositERC721(address(l1Token), tokenId, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC721Gateway.finalizeWithdrawERC721,
            (address(l1Token), address(l2Token), sender, recipient, tokenId)
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

        assertEq(address(gateway), l1Token.ownerOf(tokenId));
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
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeBatchWithdrawERC721_counterError_fails(
        address sender,
        address recipient,
        uint256 tokenCount
    ) public {
        hevm.assume(recipient != address(0));
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.batchDepositERC721(address(l1Token), _tokenIds, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC721Gateway.finalizeBatchWithdrawERC721,
            (address(l1Token), address(l2Token), sender, recipient, _tokenIds)
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

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
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
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function test_finalizeBatchWithdrawERC721_succeeds(address sender, address recipient, uint256 tokenCount) public {
        uint256 size;
        assembly {
            size := extcodesize(recipient)
        }
        hevm.assume(size == 0);
        hevm.assume(recipient != address(0));

        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.batchDepositERC721(address(l1Token), _tokenIds, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeCall(
            IL1ERC721Gateway.finalizeBatchWithdrawERC721,
            (address(l1Token), address(l2Token), sender, recipient, _tokenIds)
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
        // emit FinalizeBatchWithdrawERC721 from L1ERC721Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit IL1ERC721Gateway.FinalizeBatchWithdrawERC721(
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                _tokenIds
            );
        }

        // emit RelayedMessage from L1CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit ICrossDomainMessenger.RelayedMessage(keccak256(xDomainCalldata));
        }

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
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
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(recipient, l1Token.ownerOf(_tokenIds[i]));
        }
        assertEq(gatewayBalance - tokenCount, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance + tokenCount, l1Token.balanceOf(recipient));
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(keccak256(xDomainCalldata)));
    }

    function _testDepositERC721(uint256 tokenId, uint256 gasLimit, uint256 feePerGas) internal {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        hevm.expectRevert("no corresponding l2 token");
        gateway.depositERC721(address(l1Token), tokenId, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeDepositERC721,
            (address(l1Token), address(l2Token), address(this), address(this), tokenId)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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

        // emit FinalizeWithdrawERC721 from L1ERC721Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC721Gateway.DepositERC721(address(l1Token), address(l2Token), address(this), address(this), tokenId);

        assertEq(l1Token.ownerOf(tokenId), address(this));
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.depositERC721{value: feeToPay + EXTRA_VALUE}(address(l1Token), tokenId, gasLimit);
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(1 + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }

    function _testDepositERC721WithRecipient(
        uint256 tokenId,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        hevm.expectRevert("no corresponding l2 token");
        gateway.depositERC721(address(l1Token), tokenId, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeDepositERC721,
            (address(l1Token), address(l2Token), address(this), recipient, tokenId)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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

        // emit FinalizeWithdrawERC721 from L1ERC721Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC721Gateway.DepositERC721(address(l1Token), address(l2Token), address(this), recipient, tokenId);

        assertEq(l1Token.ownerOf(tokenId), address(this));
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.depositERC721{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, tokenId, gasLimit);
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(1 + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }

    function _testBatchDepositERC721(uint256 tokenCount, uint256 gasLimit, uint256 feePerGas) internal {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC721(address(l1Token), new uint256[](0), gasLimit);

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC721(address(l1Token), _tokenIds, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeBatchDepositERC721,
            (address(l1Token), address(l2Token), address(this), address(this), _tokenIds)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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

        // emit FinalizeWithdrawERC721 from L1ERC721Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC721Gateway.BatchDepositERC721(
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            _tokenIds
        );

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(this));
        }
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.batchDepositERC721{value: feeToPay + EXTRA_VALUE}(address(l1Token), _tokenIds, gasLimit);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(gateway));
        }
        assertEq(tokenCount + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }

    function _testBatchDepositERC721WithRecipient(
        uint256 tokenCount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC721(address(l1Token), recipient, new uint256[](0), gasLimit);

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC721(address(l1Token), recipient, _tokenIds, gasLimit);

        bytes memory message = abi.encodeCall(
            IL2ERC721Gateway.finalizeBatchDepositERC721,
            (address(l1Token), address(l2Token), address(this), recipient, _tokenIds)
        );
        bytes memory xDomainCalldata = _encodeXDomainCalldata(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        gateway.updateTokenMapping(address(l1Token), address(l2Token));

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

        // emit FinalizeWithdrawERC721 from L1ERC721Gateway
        hevm.expectEmit(true, true, true, true);
        emit IL1ERC721Gateway.BatchDepositERC721(
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            _tokenIds
        );

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(this));
        }
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
        gateway.batchDepositERC721{value: feeToPay + EXTRA_VALUE}(address(l1Token), recipient, _tokenIds, gasLimit);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(gateway));
        }
        assertEq(tokenCount + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(l1CrossDomainMessenger.messageSendTimestamp(keccak256(xDomainCalldata)), 0);
    }
}
