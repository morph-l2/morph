// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Bridge_Initializer} from "./CommonTest.t.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {MorphMintableERC20} from "../universal/MorphMintableERC20.sol";
import {Predeploys} from "../libraries/Predeploys.sol";
import {console} from "forge-std/console.sol";
import {StandardBridge} from "../universal/StandardBridge.sol";
import {L2ToL1MessagePasser} from "../L2/L2ToL1MessagePasser.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {Types} from "../libraries/Types.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {MorphMintableERC20} from "../universal/MorphMintableERC20.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";

contract L2StandardBridge_Getter_Test is Bridge_Initializer {
    function test_getters_succeeds() external {
        assert(L2Bridge.l1TokenBridge() == address(L1Bridge));
        assert(L2Bridge.OTHER_BRIDGE() == L1Bridge);
        assert(L2Bridge.messenger() == L2Messenger);
        assert(L2Bridge.MESSENGER() == L2Messenger);
        assertEq(L2Bridge.version(), "1.1.0");
        // deposits This is verified in subsequent write operations
    }
}

contract L2StandardBridge_Initialize_Test is Bridge_Initializer {
    function test_initialize_succeeds() external {
        assertEq(address(L2Bridge.MESSENGER()), address(L2Messenger));
        assertEq(address(L2Bridge.messenger()), address(L2Messenger));

        assertEq(address(L2Bridge.l1TokenBridge()), address(L1Bridge));
        assert(L2Bridge.OTHER_BRIDGE() == L1Bridge);

        assertEq(L2Bridge.version(), "1.1.0");
    }
}

contract L2StandardBridge_Initialize_TestFail is Bridge_Initializer {}

contract PreBridgeETH is Bridge_Initializer {
    // depositETH
    // - emits ETHDepositInitiated
    // - emits ETHBridgeInitiated
    // - calls MorphPortal.depositTransaction
    // - only EOA
    // - ETH ends up in the MorphPortal
    bytes internal L2BridgeMessage;
    bytes internal _extraData;

    function _preBridgeETH(
        bool isLegacy,
        address _from,
        address _to,
        uint8 _type,
        uint256 _value,
        uint32 _gasLimit
    ) internal {
        assertEq(address(messagePasser).balance, 0);
        uint256 nonce = L2Messenger.messageNonce();

        if (_type != 0) {
            _extraData = hex"dead";

            if (_from == _to) {
                if (_type == 1) {
                    L2BridgeMessage = abi.encodeWithSelector(
                        L2Bridge.withdraw.selector,
                        Predeploys.LEGACY_ERC20_ETH,
                        _value,
                        _gasLimit,
                        hex"dead"
                    );
                } else {
                    L2BridgeMessage = abi.encodeWithSelector(
                        L2Bridge.bridgeETH.selector,
                        _gasLimit,
                        _extraData
                    );
                }
            } else {
                if (_type == 1) {
                    L2BridgeMessage = abi.encodeWithSelector(
                        L2Bridge.withdrawTo.selector,
                        Predeploys.LEGACY_ERC20_ETH,
                        _to,
                        _value,
                        _gasLimit,
                        hex"dead"
                    );
                } else {
                    L2BridgeMessage = abi.encodeWithSelector(
                        L2Bridge.bridgeETHTo.selector,
                        _to,
                        _gasLimit,
                        _extraData
                    );
                }
            }

            if (isLegacy) {
                vm.expectCall(address(L2Bridge), _value, L2BridgeMessage);
            } else {
                vm.expectCall(address(L2Bridge), _value, L2BridgeMessage);
            }
        }

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeETH.selector,
            _from,
            _to,
            _value,
            _extraData
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L1Bridge),
                message,
                _gasLimit
            )
        );

        uint64 baseGas = L2Messenger.baseGas(message, _gasLimit);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L2Bridge),
            address(L1Bridge),
            _value,
            _gasLimit,
            message
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(L1Messenger),
                baseGas,
                withdrawalData
            )
        );

        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: _value,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        _appendMessageHash(withdrawalHash);

        vm.expectEmit(true, true, true, true);
        emit WithdrawalInitiated(
            address(0),
            Predeploys.LEGACY_ERC20_ETH,
            _from,
            _to,
            _value,
            _extraData,
            L2Messenger.messageNonce()
        );

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeInitiated(_from, _to, _value, _extraData);

        // L2ToL1MessagePasser will emit a MessagePassed event
        vm.expectEmit(true, true, true, true, address(messagePasser));
        emit MessagePassed(
            nonce,
            address(L2Messenger),
            address(L1Messenger),
            _value,
            baseGas,
            withdrawalData,
            withdrawalHash,
            getTreeRoot()
        );

        if (_value > 0) {
            vm.expectEmit(true, true, true, true, address(messagePasser));
            emit WithdrawerBalanceBurnt(_value);
        }

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessage(
            address(L1Bridge),
            address(L2Bridge),
            message,
            nonce,
            _gasLimit
        );

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessageExtension1(address(L2Bridge), _value);

        vm.prank(_from, _from);
    }
}

contract L2StandardBridge_Receive_Test is PreBridgeETH {
    function test_receive_succeeds(uint256 _value) external {
        vm.assume(_value <= alice.balance);
        uint256 bf = alice.balance;
        uint256 nonce = L2Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: alice,
            _type: 0,
            _value: _value,
            _gasLimit: 200_000
        });

        (bool success, ) = address(L2Bridge).call{value: _value}(hex"");
        assertEq(success, true);

        assertEq(++nonce, L2Messenger.messageNonce());
        assertEq(address(messagePasser).balance, 0);
        assertEq(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
            _value
        );
        assertEq(bf - _value, alice.balance);
    }
}

contract L2StandardBridge_Receive_TestFail is Bridge_Initializer {
    function test_receive_faileds() external {
        assertEq(address(messagePasser).balance, 0);

        // turn alice into a contract
        vm.etch(alice, address(L1Token).code);
        // emit log_named_uint("alice's balance: ", alice.balance);

        // Here expectRevert asserts that error msg is invalid，todo……
        // vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice);
        address(L2Bridge).call{value: 100}(hex"");

        // emit log_named_uint("alice's balance: ", alice.balance);
        // emit log_bytes(address(alice).code);

        // assertEq(
        //     AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
        //     100
        // );
    }
}

contract L2StandardBridge_WithdrawETH_Test is PreBridgeETH {
    function test_withdraw_ether_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);

        uint256 bf = alice.balance;
        uint256 nonce = L2Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: alice,
            _type: 1,
            _value: _value,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdraw{value: _value}({
            _l2Token: Predeploys.LEGACY_ERC20_ETH,
            _amount: _value,
            _minGasLimit: _gasLimit,
            _extraData: hex"dead"
        });

        assertEq(++nonce, L2Messenger.messageNonce());
        assertEq(address(messagePasser).balance, 0);
        assertEq(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
            _value
        );
        assertEq(bf - _value, alice.balance);
    }
}

contract L2StandardBridge_WithdrawETH_TestFail is PreBridgeETH {
    function test_withdraw_insufficientValue_reverts(uint256 _value) external {
        vm.assume(_value <= alice.balance);

        vm.expectRevert(
            "StandardBridge: bridging ETH must include sufficient ETH value"
        );

        vm.prank(alice, alice);
        L2Bridge.withdraw{value: _value}(
            address(Predeploys.LEGACY_ERC20_ETH),
            _value + 1,
            1000,
            hex""
        );
    }

    // The following exceptions cannot be tested using foundry
    // function test_withdraw_OutOfFund_reverts() external {

    //     uint256 _value = alice.balance + 1;

    //     vm.expectRevert(
    //         "EvmError: OutOfFund"
    //     );

    //     vm.prank(alice, alice);
    //     L2Bridge.withdraw{value: _value}(
    //         address(Predeploys.LEGACY_ERC20_ETH),
    //         _value,
    //         1000,
    //         hex""
    //     );
    // }
}

contract L2StandardBridge_WithdrawETHTo_Test is PreBridgeETH {
    function test_withdrawTo_ether_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);

        uint256 bf = alice.balance;
        uint256 nonce = L2Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: bob,
            _type: 1,
            _value: _value,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdrawTo{value: _value}({
            _l2Token: Predeploys.LEGACY_ERC20_ETH,
            _to: bob,
            _amount: _value,
            _minGasLimit: _gasLimit,
            _extraData: hex"dead"
        });

        assertEq(++nonce, L2Messenger.messageNonce());
        assertEq(address(messagePasser).balance, 0);
        assertEq(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
            _value
        );
        assertEq(bf - _value, alice.balance);
    }
}

contract L2StandardBridge_WithdrawETHTo_TestFail is PreBridgeETH {}

contract L2StandardBridge_BridgeETH_Test is PreBridgeETH {
    function test_bridge_ether_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);

        uint256 bf = alice.balance;
        uint256 nonce = L2Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: alice,
            _type: 2,
            _value: _value,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeETH{value: _value}(_gasLimit, hex"dead");

        assertEq(++nonce, L2Messenger.messageNonce());
        assertEq(address(messagePasser).balance, 0);
        assertEq(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
            _value
        );
        assertEq(bf - _value, alice.balance);
    }
}

contract L2StandardBridge_BridgeETH_TestFail is PreBridgeETH {}

contract L2StandardBridge_BridgeETHTo_Test is PreBridgeETH {
    function test_bridgeTo_ether_succeeds(
        uint256 _value,
        uint32 _gasLimit
    ) external {
        vm.assume(_value <= alice.balance);

        uint256 bf = alice.balance;
        uint256 nonce = L2Messenger.messageNonce();

        _preBridgeETH({
            isLegacy: true,
            _from: alice,
            _to: bob,
            _type: 2,
            _value: _value,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeETHTo{value: _value}(bob, _gasLimit, hex"dead");

        assertEq(++nonce, L2Messenger.messageNonce());
        assertEq(address(messagePasser).balance, 0);
        assertEq(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)).balance,
            _value
        );
        assertEq(bf - _value, alice.balance);
    }
}

contract L2StandardBridge_BridgeETHTo_TestFail is PreBridgeETH {}

contract PreBridgeERC20 is Bridge_Initializer {
    bytes internal L2BridgeMessage;
    bytes internal _extraData;

    function _preBridgeERC20(
        bool _isLegacy,
        address _from,
        address _to,
        address _l2Token,
        uint256 _amount,
        uint32 _gasLimit
    ) internal {
        assertEq(address(messagePasser).balance, 0);
        uint256 nonce = L2Messenger.messageNonce();

        _extraData = hex"dead";

        if (_from == _to) {
            if (_isLegacy) {
                L2BridgeMessage = abi.encodeWithSelector(
                    L2Bridge.withdraw.selector,
                    _l2Token,
                    _amount,
                    _gasLimit,
                    _extraData
                );
            } else {
                L2BridgeMessage = abi.encodeWithSelector(
                    L2Bridge.bridgeERC20.selector,
                    _l2Token,
                    address(L1Token),
                    _amount,
                    _gasLimit,
                    _extraData
                );
            }
        } else {
            if (_isLegacy) {
                L2BridgeMessage = abi.encodeWithSelector(
                    L2Bridge.withdrawTo.selector,
                    _l2Token,
                    _to,
                    _amount,
                    _gasLimit,
                    _extraData
                );
            } else {
                L2BridgeMessage = abi.encodeWithSelector(
                    L2Bridge.bridgeERC20To.selector,
                    _l2Token,
                    address(L1Token),
                    _to,
                    _amount,
                    _gasLimit,
                    _extraData
                );
            }
        }

        if (_isLegacy) {
            vm.expectCall(address(L2Bridge), L2BridgeMessage);
        } else {
            vm.expectCall(address(L2Bridge), L2BridgeMessage);
        }

        vm.expectCall(
            _l2Token,
            abi.encodeWithSelector(
                MorphMintableERC20.burn.selector,
                _from,
                _amount
            )
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector,
            address(L1Token),
            _l2Token,
            _from,
            _to,
            _amount,
            _extraData
        );

        vm.expectCall(
            address(L2Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L1Bridge),
                message,
                _gasLimit
            )
        );

        uint64 baseGas = L2Messenger.baseGas(message, _gasLimit);
        bytes memory withdrawalData = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector,
            nonce,
            address(L2Bridge),
            address(L1Bridge),
            0,
            _gasLimit,
            message
        );

        vm.expectCall(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            abi.encodeWithSelector(
                L2ToL1MessagePasser.initiateWithdrawal.selector,
                address(L1Messenger),
                baseGas,
                withdrawalData
            )
        );

        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: nonce,
                sender: address(L2Messenger),
                target: address(L1Messenger),
                value: 0,
                gasLimit: baseGas,
                data: withdrawalData
            })
        );

        _appendMessageHash(withdrawalHash);

        vm.expectEmit(true, true, true, true, address(_l2Token));
        emit Burn(_from, _amount);

        vm.expectEmit(true, true, true, true);
        emit WithdrawalInitiated(
            address(L1Token),
            _l2Token,
            _from,
            _to,
            _amount,
            _extraData,
            L2Messenger.messageNonce()
        );

        vm.expectEmit(true, true, true, true);
        emit ERC20BridgeInitiated(
            _l2Token,
            address(L1Token),
            _from,
            _to,
            _amount,
            _extraData
        );

        // L2ToL1MessagePasser will emit a MessagePassed event
        vm.expectEmit(true, true, true, true, address(messagePasser));
        emit MessagePassed(
            nonce,
            address(L2Messenger),
            address(L1Messenger),
            0,
            baseGas,
            withdrawalData,
            withdrawalHash,
            getTreeRoot()
        );

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessage(
            address(L1Bridge),
            address(L2Bridge),
            message,
            nonce,
            _gasLimit
        );

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L2Messenger));
        emit SentMessageExtension1(address(L2Bridge), 0);

        vm.prank(_from, _from);
    }
}

contract L2StandardBridge_WithdrawERC20_Test is PreBridgeERC20 {
    // withdraw
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_withdraw_withdrawingERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(L2Token), alice, _amount, true);
        assertEq(L2Token.balanceOf(alice), _amount);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: true,
            _from: alice,
            _to: alice,
            _l2Token: address(L2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdraw({
            _l2Token: address(L2Token),
            _amount: _amount,
            _minGasLimit: _gasLimit,
            _extraData: hex"dead"
        });

        assertEq(L2Token.balanceOf(alice), 0);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_WithdrawERC20_TestFail is PreBridgeERC20 {
    function test_withdraw_notEOA_reverts() external {
        // This contract has 100 L2Token
        deal(address(L2Token), address(this), 100, true);

        vm.expectRevert(
            "StandardBridge: function can only be called from an EOA"
        );
        L2Bridge.withdraw(address(L2Token), 100, 1000, hex"");
    }
}

contract L2StandardBridge_WithdrawLegacyERC20_Test is PreBridgeERC20 {
    // withdraw
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_withdrawLegacyERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(LegacyL2Token), alice, _amount, true);
        assertEq(LegacyL2Token.balanceOf(alice), _amount);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: true,
            _from: alice,
            _to: alice,
            _l2Token: address(LegacyL2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdraw({
            _l2Token: address(LegacyL2Token),
            _amount: _amount,
            _minGasLimit: _gasLimit,
            _extraData: hex"dead"
        });

        assertEq(LegacyL2Token.balanceOf(alice), 0);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_WithdrawLegacyERC20_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_BridgeERC20_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(L2Token), alice, _amount, true);
        assertEq(L2Token.balanceOf(alice), _amount);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: false,
            _from: alice,
            _to: alice,
            _l2Token: address(L2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeERC20(
            address(L2Token),
            address(L1Token),
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(L2Token.balanceOf(alice), 0);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_BridgeERC20_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_BridgeLegacyERC20_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeLegacyERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(LegacyL2Token), alice, _amount, true);
        assertEq(LegacyL2Token.balanceOf(alice), _amount);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: false,
            _from: alice,
            _to: alice,
            _l2Token: address(LegacyL2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeERC20(
            address(LegacyL2Token),
            address(L1Token),
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(LegacyL2Token.balanceOf(alice), 0);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_BridgeLegacyERC20_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_WithdrawERC20To_Test is PreBridgeERC20 {
    // withdraw
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_withdrawTo_withdrawingERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(L2Token), alice, _amount, true);
        assertEq(L2Token.balanceOf(alice), _amount);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: true,
            _from: alice,
            _to: bob,
            _l2Token: address(L2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdrawTo(
            address(L2Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(L2Token.balanceOf(alice), 0);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_WithdrawERC20To_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_WithdrawLegacyERC20To_Test is PreBridgeERC20 {
    // withdraw
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_withdrawTo_withdrawLegacyERC20_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(LegacyL2Token), alice, _amount, true);
        assertEq(LegacyL2Token.balanceOf(alice), _amount);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: true,
            _from: alice,
            _to: bob,
            _l2Token: address(LegacyL2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.withdrawTo(
            address(LegacyL2Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(LegacyL2Token.balanceOf(alice), 0);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_WithdrawLegacyERC20To_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_BridgeERC20To_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeERC20To_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(L2Token), alice, _amount, true);
        assertEq(L2Token.balanceOf(alice), _amount);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: false,
            _from: alice,
            _to: bob,
            _l2Token: address(L2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeERC20To(
            address(L2Token),
            address(L1Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(L2Token.balanceOf(alice), 0);
        assertEq(L2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_BridgeERC20To_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_BridgeLegacyERC20To_Test is PreBridgeERC20 {
    // BridgeERC20
    // - token is burned
    // - emits WithdrawalInitiated
    // - calls Withdrawer.initiateWithdrawal
    function test_bridgeLegacyERC20To_succeeds(
        uint256 _amount,
        uint32 _gasLimit
    ) external {
        deal(address(LegacyL2Token), alice, _amount, true);
        assertEq(LegacyL2Token.balanceOf(alice), _amount);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);

        _preBridgeERC20({
            _isLegacy: false,
            _from: alice,
            _to: bob,
            _l2Token: address(LegacyL2Token),
            _amount: _amount,
            _gasLimit: _gasLimit
        });

        L2Bridge.bridgeERC20To(
            address(LegacyL2Token),
            address(L1Token),
            bob,
            _amount,
            _gasLimit,
            hex"dead"
        );

        assertEq(LegacyL2Token.balanceOf(alice), 0);
        assertEq(LegacyL2Token.balanceOf(address(L2Bridge)), 0);
    }
}

contract L2StandardBridge_BridgeLegacyERC20To_TestFail is PreBridgeERC20 {}

contract L2StandardBridge_Bridge_Test is Bridge_Initializer {
    // finalizeDeposit
    // - only callable by l1TokenBridge
    // - supported token pair emits DepositFinalized
    // - invalid deposit calls Withdrawer.initiateWithdrawal
    function test_finalizeDeposit_depositingERC20_succeeds() external {
        vm.mockCall(
            address(L2Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );

        vm.expectCall(
            address(L2Token),
            abi.encodeWithSelector(
                MorphMintableERC20.mint.selector,
                alice,
                100
            )
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit DepositFinalized(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit ERC20BridgeFinalized(
            address(L2Token),
            address(L1Token),
            alice,
            alice,
            100,
            hex""
        );

        vm.prank(address(L2Messenger));
        L2Bridge.finalizeDeposit(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );
    }

    function test_finalizeDeposit_depositingETH_succeeds() external {
        vm.mockCall(
            address(L2Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit DepositFinalized(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );

        vm.expectEmit(true, true, true, true, address(L2Bridge));
        emit ERC20BridgeFinalized(
            address(L2Token), // localToken
            address(L1Token), // remoteToken
            alice,
            alice,
            100,
            hex""
        );

        vm.prank(address(L2Messenger));
        L2Bridge.finalizeDeposit(
            address(L1Token),
            address(L2Token),
            alice,
            alice,
            100,
            hex""
        );
    }

    function test_finalizeBridgeETH_incorrectValue_reverts() external {
        vm.mockCall(
            address(L2Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: ETH transfer failed");
        L2Bridge.finalizeBridgeETH{value: 50}(alice, alice, 100, hex"");
    }

    function test_finalizeBridgeETH_sendToSelf_reverts() external {
        vm.mockCall(
            address(L2Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: cannot send to self");
        L2Bridge.finalizeBridgeETH{value: 100}(
            alice,
            address(L2Bridge),
            100,
            hex""
        );
    }

    function test_finalizeBridgeETH_sendToMessenger_reverts() external {
        vm.mockCall(
            address(L2Bridge.messenger()),
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(address(L2Messenger), 100);
        vm.prank(address(L2Messenger));
        vm.expectRevert("StandardBridge: cannot send to messenger");
        L2Bridge.finalizeBridgeETH{value: 100}(
            alice,
            address(L2Messenger),
            100,
            hex""
        );
    }
}

contract L2StandardBridge_FinalizeBridgeETH_Test is Bridge_Initializer {
    function test_finalizeBridgeETH_succeeds() external {
        address messenger = address(L2Bridge.messenger());
        vm.mockCall(
            messenger,
            abi.encodeWithSelector(
                CrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(L2Bridge.OTHER_BRIDGE()))
        );
        vm.deal(messenger, 100);
        vm.prank(messenger);

        vm.expectEmit(true, true, true, true);
        emit DepositFinalized(
            address(0),
            Predeploys.LEGACY_ERC20_ETH,
            alice,
            alice,
            100,
            hex""
        );

        vm.expectEmit(true, true, true, true);
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        L2Bridge.finalizeBridgeETH{value: 100}(alice, alice, 100, hex"");
    }
}
