// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {L2MessageBaseTest} from "./base/L2MessageBase.t.sol";
import {L2CrossDomainMessenger} from "../l2/L2CrossDomainMessenger.sol";

contract L2CrossDomainMessengerTest is L2MessageBaseTest {
    function test_initialize_initializeAgain_revert() external {
        // verify the initialize only can be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l2CrossDomainMessenger.initialize(address(1));
    }

    function test_initialize_succeeds() public {
        // Deploy a transparent upgradeable proxy for the L2CrossDomainMessenger contract.
        TransparentUpgradeableProxy l2CrossDomainMessengerProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new instance of the L2CrossDomainMessenger contract implementation.
        L2CrossDomainMessenger l2CrossDomainMessengerImplTemp = new L2CrossDomainMessenger();
        hevm.startPrank(multisig);

        // Expect revert with zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);

        // Initialize the proxy with the new implementation with zero address.
        ITransparentUpgradeableProxy(address(l2CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l2CrossDomainMessengerImplTemp),
            abi.encodeCall(L2CrossDomainMessenger.initialize, address(0))
        );

        // Initialize the proxy with the new implementation with non zero address.
        ITransparentUpgradeableProxy(address(l2CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l2CrossDomainMessengerImplTemp),
            abi.encodeCall(L2CrossDomainMessenger.initialize, (NON_ZERO_ADDRESS))
        );

        L2CrossDomainMessenger l2CrossDomainMessengerTemp = L2CrossDomainMessenger(
            payable(address(l2CrossDomainMessengerProxyTemp))
        );
        hevm.stopPrank();

        // Verify the counterpart is initialized successfully.
        assertEq(l2CrossDomainMessengerTemp.counterpart(), address(NON_ZERO_ADDRESS));
    }

    function test_sendMessage_succeeds() external {
        // message config
        address to = address(bob);
        uint256 _value = 100;
        bytes memory message = "";
        uint256 gasLimit = 0;
        uint256 nonce = l2ToL1MessagePasser.leafNodesCount();

        // append message to tree
        bytes32 msgHash = keccak256(_encodeXDomainCalldata(address(this), to, _value, nonce, message));
        _appendMessageHash(msgHash);

        // revert with msg.value mismatch
        hevm.expectRevert("msg.value mismatch");
        l2CrossDomainMessenger.sendMessage(to, _value, message, gasLimit);

        // send message
        hevm.deal(address(this), _value);
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.SentMessage(address(this), to, _value, nonce, gasLimit, message);
        l2CrossDomainMessenger.sendMessage{value: _value}(to, _value, message, gasLimit);

        assertEq(getTreeRoot(), l2ToL1MessagePasser.messageRoot());
        assertEq(address(l2CrossDomainMessenger).balance, _value);
        // revert with Duplicated message
        hevm.deal(address(this), _value);
        hevm.store(address(l2ToL1MessagePasser), bytes32(l2ToL1MessagePasserLeafNodesCount), bytes32(abi.encode(0)));
        hevm.expectRevert("Duplicated message");
        l2CrossDomainMessenger.sendMessage{value: _value}(to, _value, message, gasLimit);

        // Verify the timestamp is recorded correctly.
        uint256 recordedTimestamp = l2CrossDomainMessenger.messageSendTimestamp(msgHash);
        assertTrue(recordedTimestamp != 0);
    }

    function test_relayMessage_succeeds() external {
        // send 100 to L2CrossDomainMessenger contract
        hevm.deal(address(l2CrossDomainMessenger), 100);
        // message config
        address from = address(alice);
        address to = address(Predeploys.L2_TO_L1_MESSAGE_PASSER);
        uint256 value = 100;
        uint256 nonce = 0;
        bytes memory message = "";

        // revert with Caller is not L1CrossDomainMessenger
        hevm.expectRevert("Caller is not L1CrossDomainMessenger");
        l2CrossDomainMessenger.relayMessage(from, to, value, nonce, message);

        // revert wit Forbid to call l2 to l1 message passer
        hevm.startPrank(AddressAliasHelper.applyL1ToL2Alias(l2CrossDomainMessenger.counterpart()));
        hevm.expectRevert("Forbid to call l2 to l1 message passer");
        l2CrossDomainMessenger.relayMessage(from, to, value, nonce, message);

        // relay message by L1CrossDomainMessenger alias
        to = address(bob);
        bytes32 msgHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.RelayedMessage(msgHash);
        l2CrossDomainMessenger.relayMessage(from, to, value, nonce, message);
        assertEq(address(l2CrossDomainMessenger).balance, 0);
        assertEq(address(bob).balance, (1 << 16) + 100);

        // revert wit Message was already successfully executed
        hevm.expectRevert("Message was already successfully executed");
        l2CrossDomainMessenger.relayMessage(from, to, value, nonce, message);
    }
}
