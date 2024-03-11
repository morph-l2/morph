// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

// import {MockERC1155} from "@rari-capital/solmate/src/test/utils/mocks/MockERC1155.sol";
// import {ERC1155TokenReceiver} from "@rari-capital/solmate/src/tokens/ERC1155.sol";

// import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
// import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
// import {L1ERC1155Gateway} from "../L1/gateways/L1ERC1155Gateway.sol";
// import {IL1ERC1155Gateway} from "../L1/gateways/IL1ERC1155Gateway.sol";
// import {IL2ERC1155Gateway} from "../L2/gateways/IL2ERC1155Gateway.sol";

// contract L1ERC1155GatewayTest is L1GatewayBaseTest, ERC1155TokenReceiver {
//     event FinalizeWithdrawERC1155(
//         address indexed _l1Token,
//         address indexed _l2Token,
//         address indexed _from,
//         address _to,
//         uint256 _tokenId,
//         uint256 _amount
//     );
//     event FinalizeBatchWithdrawERC1155(
//         address indexed _l1Token,
//         address indexed _l2Token,
//         address indexed _from,
//         address _to,
//         uint256[] _tokenIds,
//         uint256[] _amounts
//     );
//     event DepositERC1155(
//         address indexed _l1Token,
//         address indexed _l2Token,
//         address indexed _from,
//         address _to,
//         uint256 _tokenId,
//         uint256 _amount
//     );
//     event BatchDepositERC1155(
//         address indexed _l1Token,
//         address indexed _l2Token,
//         address indexed _from,
//         address _to,
//         uint256[] _tokenIds,
//         uint256[] _amounts
//     );
//     event RefundERC1155(
//         address indexed token,
//         address indexed recipient,
//         uint256 tokenId,
//         uint256 amount
//     );
//     event BatchRefundERC1155(
//         address indexed token,
//         address indexed recipient,
//         uint256[] tokenIds,
//         uint256[] amounts
//     );

//     uint256 private constant TOKEN_COUNT = 100;
//     uint256 private constant MAX_TOKEN_BALANCE = 1000000000;

//     L1ERC1155Gateway private gateway;

//     address private counterpartGateway;

//     MockERC1155 private l1Token;
//     MockERC1155 private l2Token;

//     function setUp() public override {
//         super.setUp();
//         _deployERC1155();

//         // Deploy tokens
//         l1Token = new MockERC1155();
//         l2Token = new MockERC1155();

//         counterpartGateway = l1ERC1155Gateway.counterpart();
//         gateway = l1ERC1155Gateway;
//         hevm.prank(multisig);
//         gateway.transferOwnership(address(this));

//         // Prepare token balances
//         for (uint256 i = 0; i < TOKEN_COUNT; i++) {
//             l1Token.mint(address(this), i, MAX_TOKEN_BALANCE, "");
//         }
//         l1Token.setApprovalForAll(address(gateway), true);
//     }

//     function testDepositERC1155(
//         uint256 tokenId,
//         uint256 amount,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) public {
//         _testDepositERC1155(tokenId, amount, gasLimit, feePerGas);
//     }

//     function testDepositERC1155WithRecipient(
//         uint256 tokenId,
//         uint256 amount,
//         address recipient,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) public {
//         _testDepositERC1155WithRecipient(
//             tokenId,
//             amount,
//             recipient,
//             gasLimit,
//             feePerGas
//         );
//     }

//     function testBatchDepositERC1155(
//         uint256 tokenCount,
//         uint256 amount,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) public {
//         _testBatchDepositERC1155(tokenCount, amount, gasLimit, feePerGas);
//     }

//     function testBatchDepositERC1155WithRecipient(
//         uint256 tokenCount,
//         uint256 amount,
//         address recipient,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) public {
//         _testBatchDepositERC1155WithRecipient(
//             tokenCount,
//             amount,
//             recipient,
//             gasLimit,
//             feePerGas
//         );
//     }

//     function testDropMessage(uint256 tokenId, uint256 amount) public {
//         gateway.updateTokenMapping(address(l1Token), address(l2Token));

//         tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             address(this),
//             tokenId,
//             amount
//         );
//         gateway.depositERC1155(
//             address(l1Token),
//             tokenId,
//             amount,
//             defaultGasLimit
//         );

//         // skip message 0
//         hevm.startPrank(address(rollup));
//         l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
//         assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
//         hevm.stopPrank();

//         // drop message 0
//         hevm.expectEmit(true, true, false, true);
//         emit RefundERC1155(address(l1Token), address(this), tokenId, amount);

//         uint256 balance = l1Token.balanceOf(address(this), tokenId);
//         l1CrossDomainMessenger.dropMessage(
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );
//         assertEq(balance + amount, l1Token.balanceOf(address(this), tokenId));
//     }

//     function testDropMessageBatch(uint256 tokenCount, uint256 amount) public {
//         tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         gateway.updateTokenMapping(address(l1Token), address(l2Token));

//         uint256[] memory _tokenIds = new uint256[](tokenCount);
//         uint256[] memory _amounts = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             _tokenIds[i] = i;
//             _amounts[i] = amount;
//         }

//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeBatchDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             address(this),
//             _tokenIds,
//             _amounts
//         );
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             _tokenIds,
//             _amounts,
//             defaultGasLimit
//         );

//         // skip message 0
//         hevm.startPrank(address(rollup));
//         l1MessageQueueWithGasPriceOracle.popCrossDomainMessage(0, 1, 0x1);
//         assertEq(l1MessageQueueWithGasPriceOracle.pendingQueueIndex(), 1);
//         hevm.stopPrank();

//         // drop message 0
//         hevm.expectEmit(true, true, false, true);
//         emit BatchRefundERC1155(
//             address(l1Token),
//             address(this),
//             _tokenIds,
//             _amounts
//         );

//         uint256[] memory balances = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             balances[i] = l1Token.balanceOf(address(this), _tokenIds[i]);
//         }
//         l1CrossDomainMessenger.dropMessage(
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );
//         for (uint256 i = 0; i < tokenCount; i++) {
//             assertEq(
//                 balances[i] + _amounts[i],
//                 l1Token.balanceOf(address(this), _tokenIds[i])
//             );
//         }
//     }

//     function testFinalizeWithdrawERC1155Failed(
//         address sender,
//         address recipient,
//         uint256 tokenId,
//         uint256 amount
//     ) public {
//         hevm.assume(recipient != address(0));
//         tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));
//         gateway.depositERC1155(
//             address(l1Token),
//             tokenId,
//             amount,
//             defaultGasLimit
//         );

//         // do finalize withdraw token
//         bytes memory message = abi.encodeWithSelector(
//             IL1ERC1155Gateway.finalizeWithdrawERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             sender,
//             recipient,
//             tokenId,
//             amount
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         messageProve(
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         // counterpart is not L2WETHGateway
//         // emit FailedRelayedMessage from L1CrossDomainMessenger
//         hevm.expectEmit(true, false, false, true);
//         emit FailedRelayedMessage(keccak256(xDomainCalldata));

//         uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
//         uint256 recipientBalance = l1Token.balanceOf(recipient, tokenId);
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//         l1CrossDomainMessenger.relayMessage(
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );
//         assertEq(gatewayBalance, l1Token.balanceOf(address(gateway), tokenId));
//         assertEq(recipientBalance, l1Token.balanceOf(recipient, tokenId));
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//     }

//     function testFinalizeWithdrawERC1155(
//         address sender,
//         address recipient,
//         uint256 tokenId,
//         uint256 amount
//     ) public {
//         uint256 size;
//         assembly {
//             size := extcodesize(recipient)
//         }
//         hevm.assume(size == 0);
//         hevm.assume(recipient != address(0));

//         tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));
//         gateway.depositERC1155(
//             address(l1Token),
//             tokenId,
//             amount,
//             defaultGasLimit
//         );

//         // do finalize withdraw token
//         bytes memory message = abi.encodeWithSelector(
//             IL1ERC1155Gateway.finalizeWithdrawERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             sender,
//             recipient,
//             tokenId,
//             amount
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         messageProve(
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
//         {
//             hevm.expectEmit(true, true, true, true);
//             emit FinalizeWithdrawERC1155(
//                 address(l1Token),
//                 address(l2Token),
//                 sender,
//                 recipient,
//                 tokenId,
//                 amount
//             );
//         }

//         // emit RelayedMessage from L1CrossDomainMessenger
//         {
//             hevm.expectEmit(true, false, false, true);
//             emit RelayedMessage(keccak256(xDomainCalldata));
//         }

//         uint256 gatewayBalance = l1Token.balanceOf(address(gateway), tokenId);
//         uint256 recipientBalance = l1Token.balanceOf(recipient, tokenId);
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//         l1CrossDomainMessenger.relayMessage(
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );
//         assertEq(
//             gatewayBalance - amount,
//             l1Token.balanceOf(address(gateway), tokenId)
//         );
//         assertEq(
//             recipientBalance + amount,
//             l1Token.balanceOf(recipient, tokenId)
//         );
//         assertBoolEq(
//             true,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//     }

//     function testFinalizeBatchWithdrawERC1155Failed(
//         address sender,
//         address recipient,
//         uint256 tokenCount,
//         uint256 amount
//     ) public {
//         hevm.assume(recipient != address(0));
//         tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         uint256[] memory _tokenIds = new uint256[](tokenCount);
//         uint256[] memory _amounts = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             _tokenIds[i] = i;
//             _amounts[i] = amount;
//         }

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             _tokenIds,
//             _amounts,
//             defaultGasLimit
//         );

//         // do finalize withdraw token
//         bytes memory message = abi.encodeWithSelector(
//             IL1ERC1155Gateway.finalizeBatchWithdrawERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             sender,
//             recipient,
//             _tokenIds,
//             _amounts
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         messageProve(
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         // counterpart is not L2WETHGateway
//         // emit FailedRelayedMessage from L1CrossDomainMessenger
//         hevm.expectEmit(true, false, false, true);
//         emit FailedRelayedMessage(keccak256(xDomainCalldata));

//         uint256[] memory gatewayBalances = new uint256[](tokenCount);
//         uint256[] memory recipientBalances = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
//             recipientBalances[i] = l1Token.balanceOf(recipient, i);
//         }
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//         l1CrossDomainMessenger.relayMessage(
//             address(uint160(address(counterpartGateway)) + 1),
//             address(gateway),
//             0,
//             0,
//             message
//         );
//         for (uint256 i = 0; i < tokenCount; i++) {
//             assertEq(
//                 gatewayBalances[i],
//                 l1Token.balanceOf(address(gateway), i)
//             );
//             assertEq(recipientBalances[i], l1Token.balanceOf(recipient, i));
//         }
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//     }

//     function testFinalizeBatchWithdrawERC1155(
//         address sender,
//         address recipient,
//         uint256 tokenCount,
//         uint256 amount
//     ) public {
//         uint256 size;
//         assembly {
//             size := extcodesize(recipient)
//         }
//         hevm.assume(size == 0);
//         hevm.assume(recipient != address(0));

//         tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         uint256[] memory _tokenIds = new uint256[](tokenCount);
//         uint256[] memory _amounts = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             _tokenIds[i] = i;
//             _amounts[i] = amount;
//         }

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             _tokenIds,
//             _amounts,
//             defaultGasLimit
//         );

//         // do finalize withdraw token
//         bytes memory message = abi.encodeWithSelector(
//             IL1ERC1155Gateway.finalizeBatchWithdrawERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             sender,
//             recipient,
//             _tokenIds,
//             _amounts
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         messageProve(
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         // emit FinalizeBatchWithdrawERC1155 from L1ERC1155Gateway
//         {
//             hevm.expectEmit(true, true, true, true);
//             emit FinalizeBatchWithdrawERC1155(
//                 address(l1Token),
//                 address(l2Token),
//                 sender,
//                 recipient,
//                 _tokenIds,
//                 _amounts
//             );
//         }

//         // emit RelayedMessage from L1CrossDomainMessenger
//         {
//             hevm.expectEmit(true, false, false, true);
//             emit RelayedMessage(keccak256(xDomainCalldata));
//         }

//         uint256[] memory gatewayBalances = new uint256[](tokenCount);
//         uint256[] memory recipientBalances = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
//             recipientBalances[i] = l1Token.balanceOf(recipient, i);
//         }
//         assertBoolEq(
//             false,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//         l1CrossDomainMessenger.relayMessage(
//             address(counterpartGateway),
//             address(gateway),
//             0,
//             0,
//             message
//         );

//         for (uint256 i = 0; i < tokenCount; i++) {
//             assertEq(
//                 gatewayBalances[i] - _amounts[i],
//                 l1Token.balanceOf(address(gateway), i)
//             );
//             assertEq(
//                 recipientBalances[i] + _amounts[i],
//                 l1Token.balanceOf(recipient, i)
//             );
//         }
//         assertBoolEq(
//             true,
//             l1CrossDomainMessenger.finalizedWithdrawals(
//                 keccak256(xDomainCalldata)
//             )
//         );
//     }

//     function _testDepositERC1155(
//         uint256 tokenId,
//         uint256 amount,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) internal {
//         tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
//         amount = bound(amount, 0, MAX_TOKEN_BALANCE);
//         gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
//         feePerGas = bound(feePerGas, 0, 1000);

//         hevm.prank(multisig);
//         l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
//         uint256 feeToPay = feePerGas * gasLimit;

//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             address(this),
//             tokenId,
//             amount
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );

//         if (amount == 0) {
//             hevm.expectRevert("deposit zero amount");
//             gateway.depositERC1155{value: feeToPay + extraValue}(
//                 address(l1Token),
//                 tokenId,
//                 amount,
//                 gasLimit
//             );
//         } else {
//             hevm.expectRevert("no corresponding l2 token");
//             gateway.depositERC1155(address(l1Token), tokenId, amount, gasLimit);

//             gateway.updateTokenMapping(address(l1Token), address(l2Token));
//             // emit QueueTransaction from L1MessageQueue
//             {
//                 hevm.expectEmit(true, true, false, true);
//                 address sender = AddressAliasHelper.applyL1ToL2Alias(
//                     address(l1CrossDomainMessenger)
//                 );
//                 emit QueueTransaction(
//                     sender,
//                     address(l2Messenger),
//                     0,
//                     0,
//                     gasLimit,
//                     xDomainCalldata
//                 );
//             }

//             // emit SentMessage from L1CrossDomainMessenger
//             {
//                 hevm.expectEmit(true, true, false, true);
//                 emit SentMessage(
//                     address(gateway),
//                     address(counterpartGateway),
//                     0,
//                     0,
//                     gasLimit,
//                     message
//                 );
//             }

//             // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
//             hevm.expectEmit(true, true, true, true);
//             emit DepositERC1155(
//                 address(l1Token),
//                 address(l2Token),
//                 address(this),
//                 address(this),
//                 tokenId,
//                 amount
//             );

//             uint256 gatewayBalance = l1Token.balanceOf(
//                 address(gateway),
//                 tokenId
//             );
//             uint256 feeVaultBalance = address(l1FeeVault).balance;
//             assertEq(
//                 l1CrossDomainMessenger.messageSendTimestamp(
//                     keccak256(xDomainCalldata)
//                 ),
//                 0
//             );
//             gateway.depositERC1155{value: feeToPay + extraValue}(
//                 address(l1Token),
//                 tokenId,
//                 amount,
//                 gasLimit
//             );
//             assertEq(
//                 amount + gatewayBalance,
//                 l1Token.balanceOf(address(gateway), tokenId)
//             );
//             assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
//             assertGt(
//                 l1CrossDomainMessenger.messageSendTimestamp(
//                     keccak256(xDomainCalldata)
//                 ),
//                 0
//             );
//         }
//     }

//     function _testDepositERC1155WithRecipient(
//         uint256 tokenId,
//         uint256 amount,
//         address recipient,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) internal {
//         tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
//         amount = bound(amount, 0, MAX_TOKEN_BALANCE);
//         gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
//         feePerGas = bound(feePerGas, 0, 1000);

//         hevm.prank(multisig);
//         l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
//         uint256 feeToPay = feePerGas * gasLimit;

//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             recipient,
//             tokenId,
//             amount
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );

//         if (amount == 0) {
//             hevm.expectRevert("deposit zero amount");
//             gateway.depositERC1155{value: feeToPay + extraValue}(
//                 address(l1Token),
//                 recipient,
//                 tokenId,
//                 amount,
//                 gasLimit
//             );
//         } else {
//             hevm.expectRevert("no corresponding l2 token");
//             gateway.depositERC1155(address(l1Token), tokenId, amount, gasLimit);

//             gateway.updateTokenMapping(address(l1Token), address(l2Token));
//             // emit QueueTransaction from L1MessageQueue
//             {
//                 hevm.expectEmit(true, true, false, true);
//                 address sender = AddressAliasHelper.applyL1ToL2Alias(
//                     address(l1CrossDomainMessenger)
//                 );
//                 emit QueueTransaction(
//                     sender,
//                     address(l2Messenger),
//                     0,
//                     0,
//                     gasLimit,
//                     xDomainCalldata
//                 );
//             }

//             // emit SentMessage from L1CrossDomainMessenger
//             {
//                 hevm.expectEmit(true, true, false, true);
//                 emit SentMessage(
//                     address(gateway),
//                     address(counterpartGateway),
//                     0,
//                     0,
//                     gasLimit,
//                     message
//                 );
//             }

//             // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
//             hevm.expectEmit(true, true, true, true);
//             emit DepositERC1155(
//                 address(l1Token),
//                 address(l2Token),
//                 address(this),
//                 recipient,
//                 tokenId,
//                 amount
//             );

//             uint256 gatewayBalance = l1Token.balanceOf(
//                 address(gateway),
//                 tokenId
//             );
//             uint256 feeVaultBalance = address(l1FeeVault).balance;
//             assertEq(
//                 l1CrossDomainMessenger.messageSendTimestamp(
//                     keccak256(xDomainCalldata)
//                 ),
//                 0
//             );
//             gateway.depositERC1155{value: feeToPay + extraValue}(
//                 address(l1Token),
//                 recipient,
//                 tokenId,
//                 amount,
//                 gasLimit
//             );
//             assertEq(
//                 amount + gatewayBalance,
//                 l1Token.balanceOf(address(gateway), tokenId)
//             );
//             assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
//             assertGt(
//                 l1CrossDomainMessenger.messageSendTimestamp(
//                     keccak256(xDomainCalldata)
//                 ),
//                 0
//             );
//         }
//     }

//     function _testBatchDepositERC1155(
//         uint256 tokenCount,
//         uint256 amount,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) internal {
//         tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
//         feePerGas = bound(feePerGas, 0, 1000);

//         hevm.prank(multisig);
//         l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
//         uint256 feeToPay = feePerGas * gasLimit;

//         uint256[] memory _tokenIds = new uint256[](tokenCount);
//         uint256[] memory _amounts = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             _tokenIds[i] = i;
//             _amounts[i] = amount;
//         }

//         hevm.expectRevert("no token to deposit");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             new uint256[](0),
//             new uint256[](0),
//             gasLimit
//         );

//         hevm.expectRevert("length mismatch");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             new uint256[](1),
//             new uint256[](0),
//             gasLimit
//         );

//         hevm.expectRevert("deposit zero amount");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             _tokenIds,
//             new uint256[](tokenCount),
//             gasLimit
//         );

//         hevm.expectRevert("no corresponding l2 token");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             _tokenIds,
//             _amounts,
//             gasLimit
//         );

//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeBatchDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             address(this),
//             _tokenIds,
//             _amounts
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));

//         // emit QueueTransaction from L1MessageQueue
//         {
//             hevm.expectEmit(true, true, false, true);
//             address sender = AddressAliasHelper.applyL1ToL2Alias(
//                 address(l1CrossDomainMessenger)
//             );
//             emit QueueTransaction(
//                 sender,
//                 address(l2Messenger),
//                 0,
//                 0,
//                 gasLimit,
//                 xDomainCalldata
//             );
//         }

//         // emit SentMessage from L1CrossDomainMessenger
//         {
//             hevm.expectEmit(true, true, false, true);
//             emit SentMessage(
//                 address(gateway),
//                 address(counterpartGateway),
//                 0,
//                 0,
//                 gasLimit,
//                 message
//             );
//         }

//         // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
//         hevm.expectEmit(true, true, true, true);
//         emit BatchDepositERC1155(
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             address(this),
//             _tokenIds,
//             _amounts
//         );

//         uint256[] memory gatewayBalances = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
//         }
//         uint256 feeVaultBalance = address(l1FeeVault).balance;
//         assertEq(
//             l1CrossDomainMessenger.messageSendTimestamp(
//                 keccak256(xDomainCalldata)
//             ),
//             0
//         );
//         gateway.batchDepositERC1155{value: feeToPay + extraValue}(
//             address(l1Token),
//             _tokenIds,
//             _amounts,
//             gasLimit
//         );
//         for (uint256 i = 0; i < tokenCount; i++) {
//             assertEq(
//                 gatewayBalances[i] + amount,
//                 l1Token.balanceOf(address(gateway), i)
//             );
//         }
//         assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
//         assertGt(
//             l1CrossDomainMessenger.messageSendTimestamp(
//                 keccak256(xDomainCalldata)
//             ),
//             0
//         );
//     }

//     function _testBatchDepositERC1155WithRecipient(
//         uint256 tokenCount,
//         uint256 amount,
//         address recipient,
//         uint256 gasLimit,
//         uint256 feePerGas
//     ) internal {
//         tokenCount = bound(tokenCount, 1, TOKEN_COUNT);
//         amount = bound(amount, 1, MAX_TOKEN_BALANCE);
//         gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
//         feePerGas = bound(feePerGas, 0, 1000);

//         hevm.prank(multisig);
//         l1MessageQueueWithGasPriceOracle.setL2BaseFee(feePerGas);
//         uint256 feeToPay = feePerGas * gasLimit;

//         uint256[] memory _tokenIds = new uint256[](tokenCount);
//         uint256[] memory _amounts = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             _tokenIds[i] = i;
//             _amounts[i] = amount;
//         }

//         hevm.expectRevert("no token to deposit");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             recipient,
//             new uint256[](0),
//             new uint256[](0),
//             gasLimit
//         );

//         hevm.expectRevert("length mismatch");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             recipient,
//             new uint256[](1),
//             new uint256[](0),
//             gasLimit
//         );

//         hevm.expectRevert("deposit zero amount");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             recipient,
//             _tokenIds,
//             new uint256[](tokenCount),
//             gasLimit
//         );

//         hevm.expectRevert("no corresponding l2 token");
//         gateway.batchDepositERC1155(
//             address(l1Token),
//             recipient,
//             _tokenIds,
//             _amounts,
//             gasLimit
//         );

//         bytes memory message = abi.encodeWithSelector(
//             IL2ERC1155Gateway.finalizeBatchDepositERC1155.selector,
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             recipient,
//             _tokenIds,
//             _amounts
//         );
//         bytes memory xDomainCalldata = abi.encodeWithSignature(
//             "relayMessage(address,address,uint256,uint256,bytes)",
//             address(gateway),
//             address(counterpartGateway),
//             0,
//             0,
//             message
//         );

//         gateway.updateTokenMapping(address(l1Token), address(l2Token));

//         // emit QueueTransaction from L1MessageQueue
//         {
//             hevm.expectEmit(true, true, false, true);
//             address sender = AddressAliasHelper.applyL1ToL2Alias(
//                 address(l1CrossDomainMessenger)
//             );
//             emit QueueTransaction(
//                 sender,
//                 address(l2Messenger),
//                 0,
//                 0,
//                 gasLimit,
//                 xDomainCalldata
//             );
//         }

//         // emit SentMessage from L1CrossDomainMessenger
//         {
//             hevm.expectEmit(true, true, false, true);
//             emit SentMessage(
//                 address(gateway),
//                 address(counterpartGateway),
//                 0,
//                 0,
//                 gasLimit,
//                 message
//             );
//         }

//         // emit FinalizeWithdrawERC1155 from L1ERC1155Gateway
//         hevm.expectEmit(true, true, true, true);
//         emit BatchDepositERC1155(
//             address(l1Token),
//             address(l2Token),
//             address(this),
//             recipient,
//             _tokenIds,
//             _amounts
//         );

//         uint256[] memory gatewayBalances = new uint256[](tokenCount);
//         for (uint256 i = 0; i < tokenCount; i++) {
//             gatewayBalances[i] = l1Token.balanceOf(address(gateway), i);
//         }
//         uint256 feeVaultBalance = address(l1FeeVault).balance;
//         assertEq(
//             l1CrossDomainMessenger.messageSendTimestamp(
//                 keccak256(xDomainCalldata)
//             ),
//             0
//         );
//         gateway.batchDepositERC1155{value: feeToPay + extraValue}(
//             address(l1Token),
//             recipient,
//             _tokenIds,
//             _amounts,
//             gasLimit
//         );
//         for (uint256 i = 0; i < tokenCount; i++) {
//             assertEq(
//                 gatewayBalances[i] + amount,
//                 l1Token.balanceOf(address(gateway), i)
//             );
//         }
//         assertEq(feeToPay + feeVaultBalance, address(l1FeeVault).balance);
//         assertGt(
//             l1CrossDomainMessenger.messageSendTimestamp(
//                 keccak256(xDomainCalldata)
//             ),
//             0
//         );
//     }
// }
