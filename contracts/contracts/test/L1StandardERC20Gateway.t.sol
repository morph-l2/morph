// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";

import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {IL2ERC20Gateway} from "../L2/gateways/IL2ERC20Gateway.sol";
import {IL1ERC20Gateway} from "../L1/gateways/IL1ERC20Gateway.sol";
import {TransferReentrantToken} from "../mock/tokens/TransferReentrantToken.sol";
import {FeeOnTransferToken} from "../mock/tokens/FeeOnTransferToken.sol";

contract L1StandardERC20GatewayTest is L1GatewayBaseTest {
    // from L1StandardERC20Gateway
    event FinalizeWithdrawERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data
    );
    event DepositERC20(
        address indexed _l1Token,
        address indexed _l2Token,
        address indexed _from,
        address _to,
        uint256 _amount,
        bytes _data
    );
    event RefundERC20(
        address indexed token,
        address indexed recipient,
        uint256 amount
    );

    address counterpartGateway;

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

        l2Token = MockERC20(
            l1StandardERC20Gateway.getL2ERC20Address(address(l1Token))
        );
        l1Token.mint(address(this), type(uint128).max);
        reentrantToken.mint(address(this), type(uint128).max);
        feeToken.mint(address(this), type(uint128).max);

        hevm.stopPrank();

        l1Token.approve(address(l1StandardERC20Gateway), type(uint256).max);
        l1Token.approve(address(l1GatewayRouter), type(uint256).max);
        feeToken.approve(address(l1StandardERC20Gateway), type(uint256).max);
        feeToken.approve(address(l1GatewayRouter), type(uint256).max);
        reentrantToken.approve(
            address(l1StandardERC20Gateway),
            type(uint256).max
        );
        reentrantToken.approve(address(l1GatewayRouter), type(uint256).max);
    }

    function testGetL2ERC20Address(address l1Address) public {
        hevm.assume(l1Address != address(0));
        assertEq(
            l1StandardERC20Gateway.getL2ERC20Address(l1Address),
            factory.computeL2TokenAddress(
                address(counterpartGateway),
                l1Address
            )
        );
    }

    function testDepositERC20(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20(false, amount, gasLimit, feePerGas);
    }

    function testRouterDepositERC20(
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20(true, amount, gasLimit, feePerGas);
    }

    function testDepositERC20WithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipient(
            false,
            amount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testDepositERC20WithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipientAndCalldata(
            false,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testRouterDepositERC20WithRecipient(
        uint256 amount,
        address recipient,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipient(
            true,
            amount,
            recipient,
            gasLimit,
            feePerGas
        );
    }

    function testRouterDepositERC20WithRecipientAndCalldata(
        uint256 amount,
        address recipient,
        bytes memory dataToCall,
        uint256 gasLimit,
        uint256 feePerGas
    ) public {
        _depositERC20WithRecipientAndCalldata(
            true,
            amount,
            recipient,
            dataToCall,
            gasLimit,
            feePerGas
        );
    }

    function testDepositReentrantToken(uint256 amount) public {
        // should revert, reentrant before transfer
        reentrantToken.setReentrantCall(
            address(l1StandardERC20Gateway),
            0,
            abi.encodeWithSignature(
                "depositERC20(address,uint256,uint256)",
                address(0),
                1,
                0
            ),
            true
        );
        amount = bound(amount, 1, reentrantToken.balanceOf(address(this)));
        hevm.expectRevert("ReentrancyGuard: reentrant call");
        l1StandardERC20Gateway.depositERC20(
            address(reentrantToken),
            amount,
            defaultGasLimit
        );

        // should revert, reentrant after transfer
        reentrantToken.setReentrantCall(
            address(l1StandardERC20Gateway),
            0,
            abi.encodeWithSignature(
                "depositERC20(address,uint256,uint256)",
                address(0),
                1,
                0
            ),
            false
        );
        amount = bound(amount, 1, reentrantToken.balanceOf(address(this)));
        hevm.expectRevert("ReentrancyGuard: reentrant call");
        l1StandardERC20Gateway.depositERC20(
            address(reentrantToken),
            amount,
            defaultGasLimit
        );
    }

    function testFeeOnTransferTokenFailed(uint256 amount) public {
        feeToken.setFeeRate(1e9);
        amount = bound(amount, 1, feeToken.balanceOf(address(this)));
        hevm.expectRevert("deposit zero amount");
        l1StandardERC20Gateway.depositERC20(
            address(feeToken),
            amount,
            defaultGasLimit
        );
    }

    function testFeeOnTransferTokenSucceed(
        uint256 amount,
        uint256 feeRate
    ) public {
        feeRate = bound(feeRate, 0, 1e9 - 1);
        amount = bound(amount, 1e9, feeToken.balanceOf(address(this)));
        feeToken.setFeeRate(feeRate);

        // should succeed, for valid amount
        uint256 balanceBefore = feeToken.balanceOf(
            address(l1StandardERC20Gateway)
        );
        uint256 fee = (amount * feeRate) / 1e9;
        l1StandardERC20Gateway.depositERC20(
            address(feeToken),
            amount,
            defaultGasLimit
        );
        uint256 balanceAfter = feeToken.balanceOf(
            address(l1StandardERC20Gateway)
        );
        assertEq(balanceBefore + amount - fee, balanceAfter);
    }

    function testDropMessage(
        uint256 amount,
        address recipient,
        bytes memory dataToCall
    ) public {
        amount = bound(amount, 1, l1Token.balanceOf(address(this)) / 2);
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            amount,
            abi.encode(
                true,
                abi.encode(
                    dataToCall,
                    abi.encode(
                        l1Token.symbol(),
                        l1Token.name(),
                        l1Token.decimals()
                    )
                )
            )
        );
        l1StandardERC20Gateway.depositERC20AndCall(
            address(l1Token),
            recipient,
            amount,
            dataToCall,
            defaultGasLimit
        );
        l1StandardERC20Gateway.depositERC20AndCall(
            address(l1Token),
            recipient,
            amount,
            dataToCall,
            defaultGasLimit
        );

        // skip message 0 and 1
        hevm.startPrank(address(rollup));
        l1MessageQueue.popCrossDomainMessage(0, 2, 0x3);
        assertEq(l1MessageQueue.pendingQueueIndex(), 2);
        hevm.stopPrank();

        // drop message 1
        hevm.expectEmit(true, true, false, true);
        emit RefundERC20(address(l1Token), address(this), amount);

        uint256 balance = l1Token.balanceOf(address(this));
        l1CrossDomainMessenger.dropMessage(
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            1,
            message
        );
        assertEq(balance + amount, l1Token.balanceOf(address(this)));
    }

    function testFinalizeWithdrawERC20Failed(
        address recipient,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        // blacklist some addresses
        hevm.assume(recipient != address(0));
        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        address _from = address(uint160(address(counterpartGateway)) + 1);

        // deposit some token to L1StandardERC20Gateway
        l1StandardERC20Gateway.depositERC20(
            address(l1Token),
            amount,
            defaultGasLimit
        );

        // do finalize withdraw token
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1Token),
            address(l2Token),
            _from,
            recipient,
            amount,
            dataToCall
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(
                _from,
                address(l1StandardERC20Gateway),
                0,
                0,
                message
            )
        );

        messageProve(_from, address(l1StandardERC20Gateway), 0, 0, message);

        // counterpart is not L2WETHGateway
        // emit FailedRelayedMessage from L1CrossDomainMessenger
        hevm.expectEmit(true, false, false, true);
        emit FailedRelayedMessage(_xDomainCalldataHash);

        uint256 gatewayBalance = l1Token.balanceOf(
            address(l1StandardERC20Gateway)
        );
        uint256 recipientBalance = l1Token.balanceOf(recipient);
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );

        l1CrossDomainMessenger.relayMessage(
            _from,
            address(l1StandardERC20Gateway),
            0,
            0,
            message
        );

        assertEq(
            gatewayBalance,
            l1Token.balanceOf(address(l1StandardERC20Gateway))
        );
        assertEq(recipientBalance, l1Token.balanceOf(recipient));
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
    }

    function testFinalizeWithdrawERC20(
        address sender,
        uint256 amount,
        bytes memory dataToCall
    ) public {
        address recipient = address(2048);
        address _from = address(counterpartGateway);

        amount = bound(amount, 1, l1Token.balanceOf(address(this)));

        // deposit some token to L1StandardERC20Gateway
        l1StandardERC20Gateway.depositERC20(
            address(l1Token),
            amount,
            defaultGasLimit
        );

        // do finalize withdraw token
        bytes memory message = abi.encodeWithSelector(
            IL1ERC20Gateway.finalizeWithdrawERC20.selector,
            address(l1Token),
            address(l2Token),
            sender,
            address(recipient),
            amount,
            dataToCall
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(
                _from,
                address(l1StandardERC20Gateway),
                0,
                0,
                message
            )
        );

        messageProve(_from, address(l1StandardERC20Gateway), 0, 0, message);

        // emit FinalizeWithdrawERC20 from L1StandardERC20Gateway
        {
            hevm.expectEmit(true, true, true, true);
            emit FinalizeWithdrawERC20(
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
            emit RelayedMessage(_xDomainCalldataHash);
        }

        uint256 gatewayBalance = l1Token.balanceOf(
            address(l1StandardERC20Gateway)
        );
        uint256 recipientBalance = l1Token.balanceOf(address(recipient));
        assertBoolEq(
            false,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );

        l1CrossDomainMessenger.relayMessage(
            _from,
            address(l1StandardERC20Gateway),
            0,
            0,
            message
        );

        assertEq(
            gatewayBalance - amount,
            l1Token.balanceOf(address(l1StandardERC20Gateway))
        );
        assertEq(
            recipientBalance + amount,
            l1Token.balanceOf(address(recipient))
        );
        assertBoolEq(
            true,
            l1CrossDomainMessenger.finalizedWithdrawals(_xDomainCalldataHash)
        );
    }

    function _depositERC20(
        bool useRouter,
        uint256 amount,
        uint256 gasLimit,
        uint256 feePerGas
    ) private {
        amount = bound(amount, 0, l1Token.balanceOf(address(this)));
        gasLimit = bound(gasLimit, defaultGasLimit / 2, defaultGasLimit);
        feePerGas = bound(feePerGas, 0, 1000);

        hevm.prank(multisig);
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            address(this),
            amount,
            abi.encode(
                true,
                abi.encode(
                    new bytes(0),
                    abi.encode(
                        l1Token.symbol(),
                        l1Token.name(),
                        l1Token.decimals()
                    )
                )
            )
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + extraValue}(
                    address(l1Token),
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{
                    value: feeToPay + extraValue
                }(address(l1Token), amount, gasLimit);
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
            emit DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                address(this),
                amount,
                new bytes(0)
            );

            uint256 l1StandardERC20GatewayBalance = l1Token.balanceOf(
                address(l1StandardERC20Gateway)
            );
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + extraValue}(
                    address(l1Token),
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{
                    value: feeToPay + extraValue
                }(address(l1Token), amount, gasLimit);
            }
            assertEq(
                amount + l1StandardERC20GatewayBalance,
                l1Token.balanceOf(address(l1StandardERC20Gateway))
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            amount,
            abi.encode(
                true,
                abi.encode(
                    new bytes(0),
                    abi.encode(
                        l1Token.symbol(),
                        l1Token.name(),
                        l1Token.decimals()
                    )
                )
            )
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + extraValue}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, gasLimit);
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
            emit DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                new bytes(0)
            );

            uint256 gatewayBalance = l1Token.balanceOf(
                address(l1StandardERC20Gateway)
            );
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositERC20{value: feeToPay + extraValue}(
                    address(l1Token),
                    recipient,
                    amount,
                    gasLimit
                );
            } else {
                l1StandardERC20Gateway.depositERC20{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, gasLimit);
            }
            assertEq(
                amount + gatewayBalance,
                l1Token.balanceOf(address(l1StandardERC20Gateway))
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
        l2GasPriceOracle.setL2BaseFee(feePerGas);

        uint256 feeToPay = feePerGas * gasLimit;
        bytes memory message = abi.encodeWithSelector(
            IL2ERC20Gateway.finalizeDepositERC20.selector,
            address(l1Token),
            address(l2Token),
            address(this),
            recipient,
            amount,
            abi.encode(
                true,
                abi.encode(
                    dataToCall,
                    abi.encode(
                        l1Token.symbol(),
                        l1Token.name(),
                        l1Token.decimals()
                    )
                )
            )
        );
        bytes memory xDomainCalldata = abi.encodeWithSignature(
            "relayMessage(address,address,uint256,uint256,bytes)",
            address(l1StandardERC20Gateway),
            address(counterpartGateway),
            0,
            0,
            message
        );

        if (amount == 0) {
            hevm.expectRevert("deposit zero amount");
            if (useRouter) {
                l1GatewayRouter.depositERC20AndCall{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, dataToCall, gasLimit);
            } else {
                l1StandardERC20Gateway.depositERC20AndCall{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, dataToCall, gasLimit);
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
            emit DepositERC20(
                address(l1Token),
                address(l2Token),
                address(this),
                recipient,
                amount,
                dataToCall
            );

            uint256 gatewayBalance = l1Token.balanceOf(
                address(l1StandardERC20Gateway)
            );
            uint256 feeVaultBalance = address(l1FeeVault).balance;
            assertEq(
                l1CrossDomainMessenger.messageSendTimestamp(
                    keccak256(xDomainCalldata)
                ),
                0
            );
            if (useRouter) {
                l1GatewayRouter.depositERC20AndCall{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, dataToCall, gasLimit);
            } else {
                l1StandardERC20Gateway.depositERC20AndCall{
                    value: feeToPay + extraValue
                }(address(l1Token), recipient, amount, dataToCall, gasLimit);
            }
            assertEq(
                amount + gatewayBalance,
                l1Token.balanceOf(address(l1StandardERC20Gateway))
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
