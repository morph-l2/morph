// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Bridge_Initializer} from "./CommonTest.t.sol";
import {StandardBridge} from "../universal/StandardBridge.sol";
import {MorphPortal} from "../L1/MorphPortal.sol";
import {L2StandardBridge} from "../L2/L2StandardBridge.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {Predeploys} from "../libraries/Predeploys.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";

contract L1StandardBridge_Getter_Test is Bridge_Initializer {
    function test_getters_succeeds() external {
        assert(L1Bridge.l2TokenBridge() == address(L2Bridge));
        assert(L1Bridge.OTHER_BRIDGE() == L2Bridge);
        assert(L1Bridge.messenger() == L1Messenger);
        assert(L1Bridge.MESSENGER() == L1Messenger);
        assertEq(L1Bridge.version(), "1.1.0");
        // deposits This is verified in subsequent write operations
    }
}

contract L1StandardBridge_Initialize_Test is Bridge_Initializer {
    function test_initialize_succeeds() external {
        assertEq(address(L1Bridge.MESSENGER()), address(L1Messenger));
        assertEq(address(L1Bridge.messenger()), address(L1Messenger));

        assertEq(
            address(L1Bridge.OTHER_BRIDGE()),
            Predeploys.L2_STANDARD_BRIDGE
        );
        assertEq(
            address(L1Bridge.l2TokenBridge()),
            Predeploys.L2_STANDARD_BRIDGE
        );
        assertEq(address(L2Bridge), Predeploys.L2_STANDARD_BRIDGE);

        assertEq(L1Bridge.version(), "1.1.0");
    }
}

contract L1StandardBridge_Initialize_TestFail is Bridge_Initializer {}

contract L1StandardBridge_Receive_Test is Bridge_Initializer {
    // receive
    // - can accept ETH
    function test_receive_succeeds(uint256 _value) external {
        vm.assume(_value <= alice.balance);

        assertEq(address(portal).balance, 0);

        // The legacy event must be emitted for backwards compatibility
        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHDepositInitiated(
            alice,
            alice,
            _value,
            hex"",
            L1Messenger.messageNonce()
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeInitiated(alice, alice, _value, hex"");

        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L2Bridge),
                abi.encodeWithSelector(
                    StandardBridge.finalizeBridgeETH.selector,
                    alice,
                    alice,
                    _value,
                    hex""
                ),
                200_000
            )
        );

        vm.prank(alice, alice);
        (bool success, ) = address(L1Bridge).call{value: _value}(hex"");
        assertEq(success, true);
        assertEq(address(portal).balance, _value);
    }
}

contract L1StandardBridge_Receive_TestFail is Bridge_Initializer {
    function test_receive_faileds() external {
        // turn alice into a contract
        vm.etch(alice, address(L1Token).code);
        // emit log_named_uint("alice's balance: ", alice.balance);

        // Here expectRevert asserts that error msg is invalid，todo……
        // vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice);
        address(L1Bridge).call{value: 100}(hex"");

        // emit log_named_uint("alice's balance: ", alice.balance);
        // emit log_bytes(address(alice).code);
    }
}

contract PreBridgeETH is Bridge_Initializer {
    // depositETH
    // - emits ETHDepositInitiated
    // - emits ETHBridgeInitiated
    // - calls MorphPortal.depositTransaction
    // - only EOA
    // - ETH ends up in the MorphPortal
    bytes public dE;

    function _preBridgeETH(
        bool isLegacy,
        address _from,
        address _to,
        uint8 _type,
        uint256 _value,
        uint32 _gasLimit
    ) internal {
        assertEq(address(portal).balance, 0);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the MorphPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(
            address(L1Messenger)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeETH.selector,
            _from,
            _to,
            _value,
            hex"dead"
        );

        if (_from == _to) {
            if (_type == 1) {
                dE = abi.encodeWithSelector(
                    L1Bridge.depositETH.selector,
                    _gasLimit,
                    hex"dead"
                );
            } else {
                dE = abi.encodeWithSelector(
                    L1Bridge.bridgeETH.selector,
                    _gasLimit,
                    hex"dead"
                );
            }
        } else {
            if (_type == 1) {
                dE = abi.encodeWithSelector(
                    L1Bridge.depositETHTo.selector,
                    _to,
                    _gasLimit,
                    hex"dead"
                );
            } else {
                dE = abi.encodeWithSelector(
                    L1Bridge.bridgeETHTo.selector,
                    _to,
                    _gasLimit,
                    hex"dead"
                );
            }
        }

        if (isLegacy) {
            vm.expectCall(address(L1Bridge), _value, dE);
        } else {
            vm.expectCall(address(L1Bridge), _value, dE);
        }

        vm.expectCall(
            address(L1Messenger),
            _value,
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L2Bridge),
                message,
                _gasLimit
            )
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L1Bridge),
            address(L2Bridge),
            _value,
            _gasLimit,
            message
        );

        uint64 baseGas = L1Messenger.baseGas(message, _gasLimit);
        vm.expectCall(
            address(portal),
            _value,
            abi.encodeWithSelector(
                MorphPortal.depositTransaction.selector,
                address(L2Messenger),
                _value,
                baseGas,
                false,
                innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(
            _value,
            _value,
            baseGas,
            false,
            innerMessage
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHDepositInitiated(
            _from,
            _to,
            _value,
            hex"dead",
            L1Messenger.messageNonce()
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeInitiated(_from, _to, _value, hex"dead");

        // MorphPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(portal));
        emit TransactionDeposited(
            l1MessengerAliased,
            address(L2Messenger),
            version,
            opaqueData
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(
            address(L2Bridge),
            address(L1Bridge),
            message,
            nonce,
            _gasLimit
        );

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), _value);

        vm.prank(_from, _from);
    }
}

// Why is the following spelling wrong
// contract L1StandardBridge_R_Test is PreBridgeETH {
//     // receive
//     // - can accept ETH
//     function test_receive_succeeds() external {
//         _preBridgeETH({ isLegacy: true, _from: alice, _to: alice, _type: 1});
//         // L1Bridge.depositETH{ value: 500 }(50000, hex"dead");
//         (bool success, ) = address(L1Bridge).call{ value: 500 }(hex"");
//         assertEq(success, true);
//         assertEq(address(portal).balance, 500);
//     }
// }

contract L1StandardBridge_DepositETH_Test is PreBridgeETH {
    function test_depositETH_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);
        vm.assume(_gasLimit <= 10_000_000);

        uint256 nonce = L1Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: alice,
            _type: 1,
            _value: _value,
            _gasLimit: _gasLimit
        });
        L1Bridge.depositETH{value: _value}(_gasLimit, hex"dead");

        assertEq(++nonce, L1Messenger.messageNonce());
        assertEq(address(portal).balance, _value);
    }
}

contract L1StandardBridge_DepositETH_TestFail is Bridge_Initializer {
    function test_depositETH_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, address(L1Token).code);

        vm.expectRevert(
            "StandardBridge: function can only be called from an EOA"
        );
        vm.prank(alice);
        L1Bridge.depositETH{value: 1}(300, hex"");
    }
}

contract L1StandardBridge_DepositETHTo_Test is PreBridgeETH {
    function test_depositETH_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);
        vm.assume(_gasLimit <= 10_000_000);

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: bob,
            _type: 1,
            _value: _value,
            _gasLimit: _gasLimit
        });
        L1Bridge.depositETHTo{value: _value}(bob, _gasLimit, hex"dead");
        assertEq(address(portal).balance, _value);
    }
}

contract L1StandardBridge_DepositETHTo_TestFail is Bridge_Initializer {}

contract L1StandardBridge_BridgeETH_Test is PreBridgeETH {
    function test_bridgeETH_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);
        vm.assume(_gasLimit <= 10_000_000);

        _preBridgeETH({
            isLegacy: false,
            _from: alice,
            _to: alice,
            _type: 0,
            _value: _value,
            _gasLimit: _gasLimit
        });
        L1Bridge.bridgeETH{value: _value}(_gasLimit, hex"dead");
        assertEq(address(portal).balance, _value);
    }
}

contract L1StandardBridge_BridgeETH_TestFail is Bridge_Initializer {
    function test_depositETH_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, address(L1Token).code);

        vm.expectRevert(
            "StandardBridge: function can only be called from an EOA"
        );
        vm.prank(alice);
        L1Bridge.depositETH{value: 1}(300, hex"");
    }
}

contract L1StandardBridge_BridgeETHTo_Test is PreBridgeETH {
    function test_bridgeETHTo_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);
        vm.assume(_gasLimit <= 10_000_000);

        _preBridgeETH({
            isLegacy: false,
            _from: alice,
            _to: bob,
            _type: 0,
            _value: _value,
            _gasLimit: _gasLimit
        });
        L1Bridge.bridgeETHTo{value: _value}(bob, _gasLimit, hex"dead");
        assertEq(address(portal).balance, _value);
    }
}

contract PreBridgeERC20 is Bridge_Initializer {
    using stdStorage for StdStorage;

    // depositERC20
    // - updates bridge.deposits
    // - emits ERC20DepositInitiated
    // - calls MorphPortal.depositTransaction
    // - only callable by EOA

    bytes public dE;

    function _preBridgeERC20(
        bool isLegacy,
        address _from,
        address _to,
        uint8 _type,
        uint256 _amount,
        uint32 _gasLimit
    ) internal {
        assertEq(address(portal).balance, 0);

        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the MorphPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(
            address(L1Messenger)
        );

        // Deal Alice's ERC20 State
        deal(address(L1Token), _from, _amount, true);
        vm.prank(_from);
        // L1Token.approve(address(L1Bridge), type(uint256).max);
        L1Token.approve(address(L1Bridge), _amount);

        // The L1Bridge should transfer alice's tokens to itself
        vm.expectCall(
            address(L1Token),
            abi.encodeWithSelector(
                ERC20.transferFrom.selector,
                _from,
                address(L1Bridge),
                _amount
            )
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector,
            address(L2Token),
            address(L1Token),
            _from,
            _to,
            _amount,
            hex"dead"
        );

        if (_from == _to) {
            if (_type == 1) {
                dE = abi.encodeWithSelector(
                    L1Bridge.depositERC20.selector,
                    address(L1Token),
                    address(L2Token),
                    _amount,
                    _gasLimit,
                    hex"dead"
                );
            } else {
                dE = abi.encodeWithSelector(
                    L1Bridge.bridgeERC20.selector,
                    address(L1Token),
                    address(L2Token),
                    _amount,
                    _gasLimit,
                    hex"dead"
                );
            }
        } else {
            if (_type == 1) {
                dE = abi.encodeWithSelector(
                    L1Bridge.depositERC20To.selector,
                    address(L1Token),
                    address(L2Token),
                    _to,
                    _amount,
                    _gasLimit,
                    hex"dead"
                );
            } else {
                dE = abi.encodeWithSelector(
                    L1Bridge.bridgeERC20To.selector,
                    address(L1Token),
                    address(L2Token),
                    _to,
                    _amount,
                    _gasLimit,
                    hex"dead"
                );
            }
        }

        if (isLegacy) {
            vm.expectCall(address(L1Bridge), 0, dE);
        } else {
            vm.expectCall(address(L1Bridge), 0, dE);
        }

        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L2Bridge),
                message,
                _gasLimit
            )
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L1Bridge),
            address(L2Bridge),
            0,
            _gasLimit,
            message
        );

        uint64 baseGas = L1Messenger.baseGas(message, _gasLimit);
        vm.expectCall(
            address(portal),
            abi.encodeWithSelector(
                MorphPortal.depositTransaction.selector,
                address(L2Messenger),
                0,
                baseGas,
                false,
                innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(
            uint256(0),
            uint256(0),
            baseGas,
            false,
            innerMessage
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20DepositInitiated(
            address(L1Token),
            address(L2Token),
            _from,
            _to,
            _amount,
            hex"dead",
            L1Messenger.messageNonce()
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeInitiated(
            address(L1Token),
            address(L2Token),
            _from,
            _to,
            _amount,
            hex"dead"
        );

        // MorphPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(portal));
        emit TransactionDeposited(
            l1MessengerAliased,
            address(L2Messenger),
            version,
            opaqueData
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(
            address(L2Bridge),
            address(L1Bridge),
            message,
            nonce,
            _gasLimit
        );

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 0);

        vm.prank(_from, _from);
    }
}

contract L1StandardBridge_DepositERC20_Test is PreBridgeERC20 {
    function test_depositERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        vm.assume(_gasLimit <= 10_000_000);

        uint256 nonce = L1Messenger.messageNonce();

        _preBridgeERC20({
            isLegacy: false,
            _from: alice,
            _to: alice,
            _type: 1,
            _amount: _amount,
            _gasLimit: _gasLimit
        });
        L1Bridge.depositERC20(
            address(L1Token),
            address(L2Token),
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(++nonce, L1Messenger.messageNonce());
        assertEq(
            L1Bridge.deposits(address(L1Token), address(L2Token)),
            _amount
        );
    }
}

contract L1StandardBridge_DepositERC20_TestFail is Bridge_Initializer {
    function test_depositERC20_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert(
            "StandardBridge: function can only be called from an EOA"
        );
        vm.prank(alice, alice);
        L1Bridge.depositERC20(address(0), address(0), 100, 100, hex"");
    }
}

contract L1StandardBridge_DepositERC20To_Test is PreBridgeERC20 {
    function test_depositERC20To_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        vm.assume(_gasLimit <= 10_000_000);

        uint256 nonce = L1Messenger.messageNonce();

        _preBridgeERC20({
            isLegacy: false,
            _from: alice,
            _to: bob,
            _type: 1,
            _amount: _amount,
            _gasLimit: _gasLimit
        });
        L1Bridge.depositERC20To(
            address(L1Token),
            address(L2Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(++nonce, L1Messenger.messageNonce());
        assertEq(
            L1Bridge.deposits(address(L1Token), address(L2Token)),
            _amount
        );
    }
}

contract L1StandardBridge_DepositERC20To_TestFail is Bridge_Initializer {}

contract L1StandardBridge_BridgeERC20_Test is PreBridgeERC20 {
    function test_depositERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        vm.assume(_gasLimit <= 10_000_000);

        uint256 nonce = L1Messenger.messageNonce();

        _preBridgeERC20({
            isLegacy: false,
            _from: alice,
            _to: alice,
            _type: 0,
            _amount: _amount,
            _gasLimit: _gasLimit
        });
        L1Bridge.bridgeERC20(
            address(L1Token),
            address(L2Token),
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(++nonce, L1Messenger.messageNonce());
        assertEq(
            L1Bridge.deposits(address(L1Token), address(L2Token)),
            _amount
        );
    }
}

contract L1StandardBridge_BridgeERC20_TestFail is Bridge_Initializer {
    function test_depositERC20_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert(
            "StandardBridge: function can only be called from an EOA"
        );
        vm.prank(alice, alice);
        L1Bridge.depositERC20(address(0), address(0), 100, 100, hex"");
    }
}

contract L1StandardBridge_BridgeRC20To_Test is PreBridgeERC20 {
    function test_depositERC20To_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        vm.assume(_gasLimit <= 10_000_000);

        uint256 nonce = L1Messenger.messageNonce();

        _preBridgeERC20({
            isLegacy: false,
            _from: alice,
            _to: bob,
            _type: 0,
            _amount: _amount,
            _gasLimit: _gasLimit
        });
        L1Bridge.bridgeERC20To(
            address(L1Token),
            address(L2Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(++nonce, L1Messenger.messageNonce());
        assertEq(
            L1Bridge.deposits(address(L1Token), address(L2Token)),
            _amount
        );
    }
}

contract L1StandardBridge_BridgeERC20To_TestFail is Bridge_Initializer {}

contract L1StandardBridge_FinalizeETHWithdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // finalizeETHWithdrawal
    // - emits ETHWithdrawalFinalized
    // - only callable by L2 bridge
    function test_finalizeETHWithdrawal_succeeds() external {
        uint256 aliceBalance = alice.balance;

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHWithdrawalFinalized(alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        vm.expectCall(alice, hex"");

        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        // ensure that the messenger has ETH to call with
        vm.deal(address(L1Bridge.messenger()), 100);
        vm.prank(address(L1Bridge.messenger()));
        L1Bridge.finalizeETHWithdrawal{value: 100}(alice, alice, 100, hex"");

        assertEq(address(L1Bridge.messenger()).balance, 0);
        assertEq(aliceBalance + 100, alice.balance);
    }
}

contract L1StandardBridge_FinalizeETHWithdrawal_TestFail is
    Bridge_Initializer
{}

contract L1StandardBridge_FinalizeERC20Withdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // finalizeERC20Withdrawal
    // - updates bridge.deposits
    // - emits ERC20WithdrawalFinalized
    // - only callable by L2 bridge
    function test_finalizeERC20Withdrawal_succeeds() external {
        deal(address(L1Token), address(L1Bridge), 100, true);

        uint256 slot = stdstore
            .target(address(L1Bridge))
            .sig("deposits(address,address)")
            .with_key(address(L1Token))
            .with_key(address(L2Token))
            .find();

        // Give the L1 bridge some ERC20 tokens
        vm.store(address(L1Bridge), bytes32(slot), bytes32(uint256(100)));
        assertEq(L1Bridge.deposits(address(L1Token), address(L2Token)), 100);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20WithdrawalFinalized(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeFinalized(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );

        vm.expectCall(
            address(L1Token),
            abi.encodeWithSelector(ERC20.transfer.selector, alice, 100)
        );

        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(address(L1Bridge.messenger()));
        L1Bridge.finalizeERC20Withdrawal(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );

        assertEq(L1Token.balanceOf(address(L1Bridge)), 0);
        assertEq(L1Token.balanceOf(address(alice)), 100);
    }
}

contract L1StandardBridge_FinalizeERC20Withdrawal_TestFail is
    Bridge_Initializer
{
    function test_finalizeERC20Withdrawal_notMessenger_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(address(28));
        vm.expectRevert(
            "StandardBridge: function can only be called from the other bridge"
        );
        L1Bridge.finalizeERC20Withdrawal(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );
    }

    function test_finalizeERC20Withdrawal_notOtherBridge_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(address(0)))
        );
        vm.prank(address(L1Bridge.messenger()));
        vm.expectRevert(
            "StandardBridge: function can only be called from the other bridge"
        );
        L1Bridge.finalizeERC20Withdrawal(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );
    }
}

contract L1StandardBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(L1Bridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        L1Bridge.finalizeBridgeETH{value: 100}(alice, alice, 100, hex"");
    }
}

contract L1StandardBridge_FinalizeBridgeETH_TestFail is Bridge_Initializer {
    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        address messenger = address(L1Bridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: ETH transfer failed");
        L1Bridge.finalizeBridgeETH{value: 50}(alice, alice, 100, hex"");
    }

    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        address messenger = address(L1Bridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to self");
        L1Bridge.finalizeBridgeETH{value: 100}(
            alice,
            address(L1Bridge),
            100,
            hex""
        );
    }

    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        address messenger = address(L1Bridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to messenger");
        L1Bridge.finalizeBridgeETH{value: 100}(alice, messenger, 100, hex"");
    }
}
