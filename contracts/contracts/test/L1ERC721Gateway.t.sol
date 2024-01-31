// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {MockERC721} from "@rari-capital/solmate/src/test/utils/mocks/MockERC721.sol";
import {ERC721TokenReceiver} from "@rari-capital/solmate/src/tokens/ERC721.sol";

import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {L2ERC721Gateway} from "../L2/gateways/L2ERC721Gateway.sol";
import {IL2ERC721Gateway} from "../L2/gateways/IL2ERC721Gateway.sol";
import {IL1ERC721Gateway} from "../L1/gateways/IL1ERC721Gateway.sol";
import {L1ERC721Gateway} from "../L1/gateways/L1ERC721Gateway.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";

contract L1ERC721GatewayTest is L1GatewayBaseTest, ERC721TokenReceiver {
    // from L1ERC721Gateway
    event FinalizeWithdrawERC721(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _tokenId
    );
    event FinalizeBatchWithdrawERC721(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256[] _tokenIds
    );
    event DepositERC721(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _tokenId
    );
    event BatchDepositERC721(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256[] _tokenIds
    );
    event RefundERC721(
        address indexed token,
        address indexed recipient,
        uint256 tokenId
    );
    event BatchRefundERC721(
        address indexed token,
        address indexed recipient,
        uint256[] tokenIds
    );

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

    function testUpdateTokenMappingFailed(address token1) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token1, token1);
        hevm.stopPrank();

        // l2 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token1, address(0));
    }

    function testUpdateTokenMappingSuccess(
        address token1,
        address token2
    ) public {
        hevm.assume(token2 != address(0));

        assertEq(gateway.tokenMapping(token1), address(0));
        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.tokenMapping(token1), token2);
    }

    function testDepositERC721(
        uint256 tokenId,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testDepositERC721(tokenId, gasLimit, feePerGas);
    }

    function testDepositERC721WithRecipient(
        uint256 tokenId,
        address to,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testDepositERC721WithRecipient(tokenId, to, gasLimit, feePerGas);
    }

    function testBatchDepositERC721WithGatewaySuccess(
        uint256 tokenCount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC721(tokenCount, gasLimit, feePerGas);
    }

    /// @dev batch deposit erc721 with recipient
    function testBatchDepositERC721WithGatewaySuccess(
        uint256 tokenCount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _testBatchDepositERC721WithRecipient(
            tokenCount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testDropMessage(uint256 tokenId) public {
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            tokenId
        );
        gateway.depositERC721(address(l1Token), tokenId, defaultGasLimit);

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueue.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueue.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit RefundERC721(address(l1Token), address(this), tokenId);

        assertEq(l1Token.ownerOf(tokenId), address(gateway));
        l1CrossDomainMessenger.dropMessage(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );
        assertEq(l1Token.ownerOf(tokenId), address(this));
    }

    function testDropMessageBatch(uint256 tokenCount) public {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        gateway.updateTokenMapping(address(l1Token), address(l2Token));

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeBatchDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            _tokenIds
        );
        gateway.batchDepositERC721(
            address(l1Token),
            _tokenIds,
            defaultGasLimit
        );

        // skip message 0
        hevm.startPrank(address(rollup));
        l1MessageQueue.popCrossDomainMessage(0, 1, 0x1);
        assertEq(l1MessageQueue.pendingQueueIndex(), 1);
        hevm.stopPrank();

        // drop message 0
        hevm.expectEmit(true, true, false, true);
        emit BatchRefundERC721(address(l1Token), address(this), _tokenIds);
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(_tokenIds[i]), address(gateway));
        }

        l1CrossDomainMessenger.dropMessage(
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(_tokenIds[i]), address(this));
        }
    }

    function testFinalizeWithdrawERC721Failed(
        address sender,
        address recipient,
        uint256 tokenId
    ) public {
        hevm.assume(recipient != address(0));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);

        gateway.updateTokenMapping(address(l1Token), address(l2Token));
        gateway.depositERC721(address(l1Token), tokenId, defaultGasLimit);

        // do finalize withdraw token
        bytes memory message = abi.encodeWithSelector(
            IL1ERC721Gateway.finalizeWithdrawERC721.selector,
            address(l1Token),
            address(l2Token),
            sender,
            recipient,
            tokenId
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        messageProve(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );
        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(keccak256(xDomainCalldata));

        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
        l1CrossDomainMessenger.relayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
    }

    function testFinalizeBatchWithdrawERC721Failed(
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
        gateway.batchDepositERC721(
            address(l1Token),
            _tokenIds,
            defaultGasLimit
        );

        // do finalize withdraw token
        bytes memory message = abi.encodeWithSelector(
            IL1ERC721Gateway.finalizeBatchWithdrawERC721.selector,
            address(l1Token),
            address(l2Token),
            sender,
            recipient,
            _tokenIds
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        messageProve(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(keccak256(xDomainCalldata));

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
        l1CrossDomainMessenger.relayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
        assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
    }

    function testFinalizeBatchWithdrawERC721(
        address sender,
        address recipient,
        uint256 tokenCount
    ) public {
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
        gateway.batchDepositERC721(
            address(l1Token),
            _tokenIds,
            defaultGasLimit
        );

        // do finalize withdraw token
        bytes memory message = abi.encodeWithSelector(
            IL1ERC721Gateway.finalizeBatchWithdrawERC721.selector,
            address(l1Token),
            address(l2Token),
            sender,
            recipient,
            _tokenIds
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );
        messageProve(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );

        // emit FinalizeBatchWithdrawERC721 from L1ERC721Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit FinalizeBatchWithdrawERC721(
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
            emit RelayedMessage(keccak256(xDomainCalldata));
        }

        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(address(gateway), l1Token.ownerOf(_tokenIds[i]));
        }
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
        l1CrossDomainMessenger.relayMessage(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(recipient, l1Token.ownerOf(_tokenIds[i]));
        }
        assertEq(
            gatewayBalance - tokenCount,
            l1Token.balanceOf(address(gateway))
        );
        assertEq(recipientBalance + tokenCount, l1Token.balanceOf(recipient));
        assertBoolEq(
            true,
            l1CrossDomainMessenger.finalizedWithdrawals(
                keccak256(xDomainCalldata)
            )
        );
    }

    function _testDepositERC721(
        uint256 tokenId,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l2GasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        hevm.expectRevert("no corresponding l2 token");
        gateway.depositERC721(address(l1Token), tokenId, gasLimit);

        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            tokenId
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
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
        emit DepositERC721(
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            tokenId
        );

        assertEq(l1Token.ownerOf(tokenId), address(this));
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
        gateway.depositERC721{value: feeToPay + extraValue}(
            address(l1Token),
            tokenId,
            gasLimit
        );
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(1 + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        hevm.expectRevert("no corresponding l2 token");
        gateway.depositERC721(address(l1Token), tokenId, gasLimit);

        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            tokenId
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
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
        emit DepositERC721(
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            tokenId
        );

        assertEq(l1Token.ownerOf(tokenId), address(this));
        uint256 gatewayBalance = l1Token.balanceOf(address(gateway));
        uint256 feeVaultBalance = address(l1FeeVault).balance;
        assertEq(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
        gateway.depositERC721{value: feeToPay + extraValue}(
            address(l1Token),
            recipient,
            tokenId,
            gasLimit
        );
        assertEq(address(gateway), l1Token.ownerOf(tokenId));
        assertEq(1 + gatewayBalance, l1Token.balanceOf(address(gateway)));
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
    }

    function _testBatchDepositERC721(
        uint256 tokenCount,
        uint256 gasLimit,
        uint256 feePerGas
    ) internal {
        tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l2GasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC721(
            address(l1Token),
            new uint256[](0),
            gasLimit
        );

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC721(address(l1Token), _tokenIds, gasLimit);

        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeBatchDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            _tokenIds
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
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
        emit BatchDepositERC721(
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
        assertEq(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
        gateway.batchDepositERC721{value: feeToPay + extraValue}(
            address(l1Token),
            _tokenIds,
            gasLimit
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(gateway));
        }
        assertEq(
            tokenCount + gatewayBalance,
            l1Token.balanceOf(address(gateway))
        );
        assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
        assertGt(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);
        uint256 feeToPay = feePerGas * gasLimit;

        uint256[] memory _tokenIds = new uint256[](tokenCount);
        for (uint256 i = 0; i < tokenCount; i++) {
            _tokenIds[i] = i;
        }

        hevm.expectRevert("no token to deposit");
        gateway.batchDepositERC721(
            address(l1Token),
            recipient,
            new uint256[](0),
            gasLimit
        );

        hevm.expectRevert("no corresponding l2 token");
        gateway.batchDepositERC721(
            address(l1Token),
            recipient,
            _tokenIds,
            gasLimit
        );

        bytes memory message = abi.encodeWithSelector(
            IL2ERC721Gateway.finalizeBatchDepositERC721.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            _tokenIds
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
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
        emit BatchDepositERC721(
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
        assertEq(
            l1CrossDomainMessenger.messageSendTimestamp(
                keccak256(xDomainCalldata)
            ),
            0
        );
        gateway.batchDepositERC721{value: feeToPay + extraValue}(
            address(l1Token),
            recipient,
            _tokenIds,
            gasLimit
        );
        for (uint256 i = 0; i < tokenCount; i++) {
            assertEq(l1Token.ownerOf(i), address(gateway));
        }
        assertEq(
            tokenCount + gatewayBalance,
            l1Token.balanceOf(address(gateway))
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
