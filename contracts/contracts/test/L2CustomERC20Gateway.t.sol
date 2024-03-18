// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {L2CustomERC20Gateway} from "../L2/gateways/L2CustomERC20Gateway.sol";
import {L2GatewayRouter} from "../L2/gateways/L2GatewayRouter.sol";
import {IL2ERC20Gateway} from "../L2/gateways/IL2ERC20Gateway.sol";
import {L2CrossDomainMessenger} from "../L2/L2CrossDomainMessenger.sol";
import {IL1ERC20Gateway} from "../L1/gateways/IL1ERC20Gateway.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";

contract L2CustomERC20GatewayTest is L2GatewayBaseTest {
    event WithdrawERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data,
        uint256 nonce
    );
    event FinalizeDepositERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data
    );

    L2CustomERC20Gateway private gateway;
    L2GatewayRouter private router;
    L2CrossDomainMessenger private l2Messenger;

    address private counterpartGateway;
    address private feeVault;
    address private l1Messenger;

    MockERC20 private l1Token;
    MockERC20 private l2Token;

    function setUp() public override {
        super.setUp();
        _deployCustomERC20();
        gateway = l2CustomERC20Gateway;
        router = l2GatewayRouter;
        counterpartGateway = gateway.counterpart();
        l2Messenger = l2CrossDomainMessenger;
        feeVault = l2FeeVault;
        l1Messenger = address(NON_ZERO_ADDRESS);

        // transfer ownership
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        // Deploy tokens
        l1Token = new MockERC20("Mock L1", "ML1", 18);
        l2Token = new MockERC20("Mock L2", "ML2", 18);

        // Prepare token balances
        l2Token.mint(address(this), type(uint128).max);
        l2Token.approve(address(gateway), type(uint256).max);
    }

    function testUpdateTokenMappingFailed(address token2) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token2, token2);
        hevm.stopPrank();

        // l1 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token2, address(0));
    }

    function testUpdateTokenMappingSuccess(
        address token1,
        address token2
    ) public {
        hevm.assume(token1 != address(0));

        assertEq(gateway.getL1ERC20Address(token2), address(0));
        gateway.updateTokenMapping(token2, token1);
        assertEq(gateway.getL1ERC20Address(token2), token1);
    }

    function testWithdrawERC20(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20(false, amount, gasLimit, feePerGas);
    }

    function testWithdrawERC20WithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipient(
            false,
            amount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testWithdrawERC20WithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _withdrawERC20WithRecipientAndCalldata(
            false,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testFinalizeDepositERC20FailedMocking(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, 100000);

        // revert when caller is not messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositERC20(
            address(l1Token),
            address(l2Token),
            sender,
            recipient,
            amount,
            dataToCall
        );

        MockCrossDomainMessenger mockMessenger = new MockCrossDomainMessenger();
        hevm.store(
            address(gateway),
            bytes32(eth_erc20_messenger_slot),
            bytes32(abi.encode(address(mockMessenger)))
        );

        // only call by counterpart
        hevm.expectRevert("only call by counterpart");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );

        mockMessenger.setXDomainMessageSender(address(counterpartGateway));

        // msg.value mismatch
        hevm.expectRevert("nonzero msg.value");
        mockMessenger.callTarget{value: 1}(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );

        // l1 token mismatch
        hevm.expectRevert("l1 token mismatch");
        mockMessenger.callTarget(
            address(gateway),
            abi.encodeWithSelector(
                gateway.finalizeDepositERC20.selector,
                address(l1Token),
                address(l2Token),
                sender,
                recipient,
                amount,
                dataToCall
            )
        );
    }

    function testFinalizeDepositERC20Failed(
        address sender,
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));

        gateway.updateTokenMapping(address(l2Token), address(l1Token));

        amount = bound(amount, 1, l2Token.balanceOf(address(this)));

        // do finalize deposit token
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            sender,
            recipient,
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );

        // counterpart is not L1CustomERC20Gateway
        // emit FailedRelayedMessage from L2CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(keccak256(xDomainCalldata));

        uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
        uint256 recipientBalance = l2Token.balanceOf(recipient);
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
        hevm.startPrank(
            AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger))
        );
        l2Messenger.relayMessage(
            address(uint160(address(counterpartGateway)) + 1),
            address(gateway),
            0,
            0,
            message
        );
        hevm.stopPrank();
        assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
        assertEq(recipientBalance, l2Token.balanceOf(recipient));
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
    }

    function testFinalizeDepositERC20(
        address sender,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        address recipient = address(2048);

        gateway.updateTokenMapping(address(l2Token), address(l1Token));

        amount = bound(amount, 1, l2Token.balanceOf(address(this)));

        // do finalize deposit token
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            sender,
            address(recipient),
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );

        // emit FinalizeDepositERC20 from L2CustomERC20Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit FinalizeDepositERC20(
                address(l1Token),
                address(l2Token),
                sender,
                address(recipient),
                amount,
                dataToCall
            );
        }

        // emit RelayedMessage from L2CrossDomainMessenger
        {
            hevm.expectEmit(true, false, false, true);
            emit RelayedMessage(keccak256(xDomainCalldata));
        }

        uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
        uint256 recipientBalance = l2Token.balanceOf(address(recipient));
        assertBoolEq(
            false,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
        hevm.startPrank(
            AddressAliasHelper.applyL1ToL2Alias(address(l1Messenger))
        );
        l2Messenger.relayMessage(
            address(counterpartGateway),
            address(gateway),
            0,
            0,
            message
        );
        hevm.stopPrank();
        assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
        assertEq(
            recipientBalance + amount,
            l2Token.balanceOf(address(recipient))
        );
        assertBoolEq(
            true,
            l2Messenger.isL1MessageExecuted(keccak256(xDomainCalldata))
        );
    }

    function _withdrawERC20(
        bool useRouter,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            amount,
            new bytes(0)
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        } else {
            gateway.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    amount,
                    gasLimit
                );
            }
        } else {
            // emit AppendMessage from L2MessageQueue
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
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

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    amount,
                    gasLimit
                );
            }
            assertEq(gatewayBalance, l1Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }

    function _withdrawERC20WithRecipient(
        bool useRouter,
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            amount,
            new bytes(0)
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        } else {
            gateway.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    gasLimit
                );
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
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

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                new bytes(0),
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    gasLimit
                );
            }
            assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }

    function _withdrawERC20WithRecipientAndCalldata(
        bool useRouter,
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l2Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, 21000, 1000000);
        feePerGas = 0;

        setL1BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            amount,
            dataToCall
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        hevm.expectRevert("no corresponding l1 token");
        if (useRouter) {
            router.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        } else {
            gateway.withdrawERC20{value: feeToPay}(
                address(l2Token),
                amount,
                gasLimit
            );
        }

        gateway.updateTokenMapping(address(l2Token), address(l1Token));
        if (amount == 0) {
            hevm.expectRevert("withdraw zero amount");
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
        } else {
            _appendMessageHash(keccak256(xDomainCalldata));
            bytes32 rootHash = getTreeRoot();
            // emit AppendMessage from L2MessageQueue
            {
                hevm.expectEmit(false, false, false, true);
                emit AppendMessage(0, keccak256(xDomainCalldata), rootHash);
            }

            // emit SentMessage from L2CrossDomainMessenger
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

            // emit WithdrawERC20 from L2CustomERC20Gateway
            hevm.expectEmit(true, true, true, true);
            emit WithdrawERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                dataToCall,
                0
            );

            uint256 gatewayBalance = l2Token.balanceOf(address(gateway));
            uint256 feeVaultBalance = address(feeVault).balance;
            assertEq(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
            if (useRouter) {
                router.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            } else {
                gateway.withdrawERC20AndCall{value: feeToPay}(
                    address(l2Token),
                    recipient,
                    amount,
                    dataToCall,
                    gasLimit
                );
            }
            assertEq(gatewayBalance, l2Token.balanceOf(address(gateway)));
            assertEq(feeToPay + feeVaultBalance, address(feeVault).balance);
            assertGt(
                l2Messenger.messageSendTimestamp(keccak256(xDomainCalldata)),
                0
            );
        }
    }
}
