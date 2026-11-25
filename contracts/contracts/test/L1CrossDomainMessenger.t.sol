// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {Vm} from "forge-std/Vm.sol";

import {L1MessageQueueWithGasPriceOracle} from "../l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {IL1CrossDomainMessenger} from "../l1/L1CrossDomainMessenger.sol";
import {L1CrossDomainMessenger} from "../l1/L1CrossDomainMessenger.sol";
import {IRollup} from "../l1/rollup/IRollup.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {RateLimiter} from "../libraries/common/RateLimiter.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {IL2ETHGateway} from "../l2/gateways/IL2ETHGateway.sol";
import {ReceiveRevert} from "../mock/ReceiveRevert.sol";

contract L1CrossDomainMessengerTest is L1GatewayBaseTest {
    uint256 L1CrossDomainMessenger_provenWithdrawals_slot = 251;
    uint256 L1CrossDomainMessenger_FeeVault_slot = 203;
    address refundAddress = address(2048);
    address counterpartGateway;

    function setUp() public virtual override {
        super.setUp();
        counterpartGateway = l1ETHGateway.counterpart();
    }

    function test_initialize_initializeAgain_reverts() external {
        // Verify that the initialize function can only be called once.
        // Since the initialize function is already called in L1MessageBaseTest,
        // calling it again should trigger an error message.
        // Use expectRevert to catch and assert the error message as expected.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1CrossDomainMessenger.initialize(address(1), address(1), address(1));
    }

    function test_initialize_zeroAddress_reverts() external {
        address feeVault = address(1);
        address rollup = address(1);
        address messageQueue = address(1);

        // Deploy a TransparentUpgradeableProxy contract for l1CrossDomainMessengerProxyTemp.
        TransparentUpgradeableProxy l1CrossDomainMessengerProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy L1CrossDomainMessenger implementation.
        L1CrossDomainMessenger l1CrossDomainMessengerImplTemp = new L1CrossDomainMessenger();

        hevm.startPrank(multisig);

        // Expect revert due to feeVault being zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l1CrossDomainMessengerImplTemp),
            abi.encodeCall(L1CrossDomainMessenger.initialize, (address(0), rollup, messageQueue))
        );

        // Expect revert due to rollup being zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l1CrossDomainMessengerImplTemp),
            abi.encodeCall(L1CrossDomainMessenger.initialize, (feeVault, address(0), messageQueue))
        );

        // Expect revert due to messageQueue being zero address.
        hevm.expectRevert(ICrossDomainMessenger.ErrZeroAddress.selector);
        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l1CrossDomainMessengerImplTemp),
            abi.encodeCall(L1CrossDomainMessenger.initialize, (feeVault, rollup, address(0)))
        );

        hevm.stopPrank();
    }

    function test_initialize_succeeds() external {
        hevm.startPrank(multisig);

        // Deploy a TransparentUpgradeableProxy contract for l1CrossDomainMessengerProxyTemp.
        TransparentUpgradeableProxy l1CrossDomainMessengerProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy L1CrossDomainMessenger implementation.
        L1CrossDomainMessenger l1CrossDomainMessengerImplTemp = new L1CrossDomainMessenger();

        // Verify that the UpdateMaxReplayTimes event is emitted successfully.
        hevm.expectEmit(false, false, false, true);
        emit IL1CrossDomainMessenger.UpdateMaxReplayTimes(0, 3);

        ITransparentUpgradeableProxy(address(l1CrossDomainMessengerProxyTemp)).upgradeToAndCall(
            address(l1CrossDomainMessengerImplTemp),
            abi.encodeCall(
                L1CrossDomainMessenger.initialize,
                (l1FeeVault, address(l1CrossDomainMessengerProxyTemp), address(l1CrossDomainMessengerProxyTemp))
            )
        );

        L1CrossDomainMessenger l1CrossDomainMessengerTemp = L1CrossDomainMessenger(
            payable(address(l1CrossDomainMessengerProxyTemp))
        );
        hevm.stopPrank();

        // Verify the state variable maxReplayTimes is initialized successfully.
        assertEq(l1CrossDomainMessengerTemp.maxReplayTimes(), 3);
    }

    function test_proveWithdrawalTransaction_relayMessage_succeeds() external {
        // tx msg set
        address from = address(alice);
        address to = address(bob);
        uint256 value = 10;
        hevm.deal(address(l1CrossDomainMessenger), value);

        uint256 nonce = 0;
        bytes memory message = "";
        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));
        // get proof from ffi
        _appendMessageHash(_xDomainCalldataHash);
        (bytes32 wdHashRes, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi.getProveWithdrawalTransactionInputs(
            _xDomainCalldataHash
        );
        assertEq(_xDomainCalldataHash, wdHashRes);
        assertEq(getTreeRoot(), wdRoot);

        // prove without rollup
        hevm.expectRevert("Messenger: withdrawalRoot not finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);

        // Verify the RelayedMessage event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.RelayedMessage(_xDomainCalldataHash);

        // Mock the rollup contract to simulate successful submission of the withdrawal root.
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (wdRoot)),
            abi.encode(true)
        );
        uint256 balanceBefore = address(bob).balance;
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);
        assertEq(balanceBefore + value, address(bob).balance);

        // prove again
        hevm.expectRevert("Messenger: withdrawal has already been finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);

        message = "FailedRelayedMessage";
        _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));
        // get proof from ffi
        _appendMessageHash(_xDomainCalldataHash);
        (wdHashRes, wdProof, wdRoot) = ffi.getProveWithdrawalTransactionInputs(_xDomainCalldataHash);
        // mock call rollup withdrawal root submitted success
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (wdRoot)),
            abi.encode(true)
        );
        // verify the event FailedRelayedMessage can be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(_xDomainCalldataHash);
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);

        // verify it throws the error "Messenger: Forbid to call message queue" when the address of to is equal to messageQueue.
        to = address(l1CrossDomainMessenger.messageQueue());
        hevm.expectRevert("Messenger: Forbid to call message queue");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);

        // verify it throws the error "Messenger: Forbid to call message queue" when the address of from is equal to xDomainMessageSender.
        to = address(bob);
        from = address(l1CrossDomainMessenger.xDomainMessageSender());
        hevm.expectRevert("Messenger: Invalid message sender");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);

        // verify it throws the error "Messenger: invalid withdrawal inclusion" when verifyMerkleProof returns false.
        from = address(alice);
        message = "invalid";
        hevm.expectRevert("Messenger: invalid withdrawal inclusion proof");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, wdProof, wdRoot);
    }

    function test_replayMessage_succeeds(uint256 exceedValue) external {
        hevm.deal(address(this), 1 ether);
        // tx msg set
        exceedValue = bound(exceedValue, 1, address(this).balance / 2);
        address from = address(this);
        address to = address(bob);
        uint256 value = 0;
        bytes memory message = "";
        uint256 nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(from, to, value, nonce, message);
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_xDomainCalldata);

        // updateMaxReplayTimes to 0
        hevm.expectRevert("replay times must be greater than 0");
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(0);

        // append a message
        l1CrossDomainMessenger.sendMessage{value: 100}(to, value, message, gas, refundAddress);

        // Provided message has not been enqueued
        hevm.expectRevert("Provided message has not been enqueued");
        l1CrossDomainMessenger.replayMessage(
            address(this),
            address(0),
            101,
            0,
            new bytes(0),
            defaultGasLimit,
            refundAddress
        );

        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(1);
        // Insufficient msg.value
        hevm.expectRevert("Insufficient msg.value for fee");
        l1CrossDomainMessenger.replayMessage(from, to, value, nonce, message, defaultGasLimit, refundAddress);

        hevm.prank(multisig);
        uint256 _fee = l1MessageQueueWithGasPriceOracle.l2BaseFee() * defaultGasLimit;

        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(1);

        // refund exceed fee
        uint256 balanceBefore = refundAddress.balance;
        uint256 feeVaultBefore = l1FeeVault.balance;
        l1CrossDomainMessenger.replayMessage{value: _fee + exceedValue}(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );
        assertEq(balanceBefore + exceedValue, refundAddress.balance);
        assertEq(feeVaultBefore + _fee, l1FeeVault.balance);

        // Exceed maximum replay times
        hevm.expectRevert("Exceed maximum replay times");
        l1CrossDomainMessenger.replayMessage{value: _fee}(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );

        // test replay list
        // 1. send a message with nonce 2
        // 2. replay 3 times
        hevm.startPrank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(0);
        l1CrossDomainMessenger.updateMaxReplayTimes(100);
        hevm.stopPrank();
        l1CrossDomainMessenger.sendMessage{value: 100}(address(0), 100, new bytes(0), defaultGasLimit, refundAddress);
        bytes32 hash = keccak256(_encodeXDomainCalldata(address(this), address(0), 100, 2, new bytes(0)));
        (uint256 _replayTimes, uint256 _lastIndex) = l1CrossDomainMessenger.replayStates(hash);
        assertEq(_replayTimes, 0);
        assertEq(_lastIndex, 0);
        for (uint256 i = 0; i < 3; i++) {
            l1CrossDomainMessenger.replayMessage(
                address(this),
                address(0),
                100,
                2,
                new bytes(0),
                defaultGasLimit,
                refundAddress
            );
            (_replayTimes, _lastIndex) = l1CrossDomainMessenger.replayStates(hash);
            assertEq(_replayTimes, i + 1);
            assertEq(_lastIndex, i + 3);
            assertEq(l1CrossDomainMessenger.prevReplayIndex(i + 3), i + 2 + 1);
            for (uint256 j = 0; j <= i; j++) {
                assertEq(l1CrossDomainMessenger.prevReplayIndex(i + 3 - j), i + 2 - j + 1);
            }
        }
    }

    function test_forbidCallMessageQueueFromL2_succeeds() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1MessageQueueWithGasPriceOracle);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (withdrawalRoot)),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call message queue");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, withdrawalProof, withdrawalRoot);
    }

    function test_forbidCallSelfFromL2_succeeds() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1CrossDomainMessenger);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(_encodeXDomainCalldata(from, to, value, nonce, message));
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (withdrawalRoot)),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call self");
        l1CrossDomainMessenger.proveAndRelayMessage(from, to, value, nonce, message, withdrawalProof, withdrawalRoot);
    }

    function test_sendMessage_succeeds() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(sender, to, value, nonce, data);
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_xDomainCalldata);
        hevm.expectEmit(true, true, true, true);
        emit ICrossDomainMessenger.SentMessage(sender, to, value, nonce, gas, data);

        hevm.expectCall(
            address(l1MessageQueueWithGasPriceOracle),
            abi.encodeCall(
                l1MessageQueueWithGasPriceOracle.appendCrossDomainMessage,
                (Predeploys.L2_CROSS_DOMAIN_MESSENGER, gas, _xDomainCalldata)
            )
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
    }

    function test_sendMessage_value_succeeds() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(sender, to, value, nonce, data);
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_xDomainCalldata);
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);

        // send value not zero
        // set base fee to 100
        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(100);

        value = 1 ether;
        nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        _xDomainCalldata = _encodeXDomainCalldata(sender, to, value, nonce, data);
        gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_xDomainCalldata);
        hevm.expectRevert("Insufficient msg.value");
        l1CrossDomainMessenger.sendMessage{value: 1 ether}(to, value, data, gas);
        // give enough value
        uint256 fee = l1MessageQueueWithGasPriceOracle.estimateCrossDomainMessageFee(sender, gas);
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee}(to, value, data, gas);

        // send more value with refund address
        fee = l1MessageQueueWithGasPriceOracle.estimateCrossDomainMessageFee(sender, gas);
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee * 2}(to, value, data, gas, refundAddress);
        assertEq(address(refundAddress).balance, fee);

        // verify refundAddress.call() failed, trigger the error message "Failed to refund the fee" as expected.
        ReceiveRevert receiveRevert = new ReceiveRevert();
        hevm.expectRevert("Failed to refund the fee");
        l1CrossDomainMessenger.sendMessage{value: 2 ether}(to, value, data, gas, address(receiveRevert));

        // verify the call is executed as expected, and the fee is added into the balance of the feeVault.
        uint256 initialFeeVaultBalance = l1FeeVault.balance;
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee}(to, value, data, gas);
        assertEq(address(l1FeeVault).balance, initialFeeVaultBalance + fee);

        // verify it throws a "Failed to refund the fee" error when the call fails.
        hevm.store(
            address(l1CrossDomainMessenger),
            bytes32(L1CrossDomainMessenger_FeeVault_slot),
            bytes32(abi.encode(address(receiveRevert)))
        );
        hevm.expectRevert("Failed to deduct the fee");
        l1CrossDomainMessenger.sendMessage{value: 2 ether}(to, value, data, gas, refundAddress);
    }

    function test_sendMessage_twice_succeeds() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1CrossDomainMessenger.messageNonce();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(sender, to, value, nonce, data);
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(_xDomainCalldata);

        // Tests that the sendMessage function is able to send the same message twice.
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
        assertEq(nonce + 2, l1CrossDomainMessenger.messageNonce());
    }

    function test_updateMaxReplayTimes_succeeds(uint256 _maxReplayTimes) external {
        hevm.assume(_maxReplayTimes > 0);
        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        hevm.stopPrank();

        hevm.expectEmit(false, false, false, true);
        emit IL1CrossDomainMessenger.UpdateMaxReplayTimes(3, _maxReplayTimes);

        assertEq(l1CrossDomainMessenger.maxReplayTimes(), 3);
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        assertEq(l1CrossDomainMessenger.maxReplayTimes(), _maxReplayTimes);
    }

    function test_setPause_succeeds() external {
        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1CrossDomainMessenger.setPause(false);
        hevm.stopPrank();

        // pause
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(true);
        assertBoolEq(true, l1CrossDomainMessenger.paused());

        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.sendMessage(address(0), 0, new bytes(0), defaultGasLimit);
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.sendMessage(address(0), 0, new bytes(0), defaultGasLimit, address(0));

        (, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi.getProveWithdrawalTransactionInputs(bytes32(uint256(1)));
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.proveAndRelayMessage(address(0), address(0), 0, 0, new bytes(0), wdProof, wdRoot);
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.replayMessage(address(0), address(0), 0, 0, new bytes(0), 0, address(0));

        // unpause
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(false);
        assertBoolEq(false, l1CrossDomainMessenger.paused());
    }

    function test_proveAndRelayMessage_ethRateLimit_succeeds() external {
        hevm.deal(address(l1CrossDomainMessenger), 1000 ether);
        _setETHRateLimiterConfig(true, 100 ether, 1 ether);

        // released within the limit
        bytes32 _hash = keccak256(_encodeXDomainCalldata(alice, bob, 60 ether, 0, ""));
        (bytes32[32] memory _proof, bytes32 _root) = messageProveAndRelayPrepare(alice, bob, 60 ether, 0, "");
        hevm.expectEmit(true, false, false, true);
        emit IL1CrossDomainMessenger.ETHRateLimitConsumed(_hash, 60 ether, 40 ether);
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 60 ether, 0, "", _proof, _root);
        _assertETHRateLimiterState(40 ether, true, 100 ether, 1 ether);

        // reverted when the limit is reached, and released after the wait time
        _hash = keccak256(_encodeXDomainCalldata(alice, bob, 50 ether, 0, ""));
        (_proof, _root) = messageProveAndRelayPrepare(alice, bob, 50 ether, 0, "");
        hevm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitReached.selector, 10, 40 ether));
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 50 ether, 0, "", _proof, _root);
        assertBoolEq(false, l1CrossDomainMessenger.finalizedWithdrawals(_hash));

        hevm.warp(block.timestamp + 10);
        uint256 _balance = bob.balance;
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 50 ether, 0, "", _proof, _root);
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(_hash));
        assertEq(bob.balance, _balance + 50 ether);
        _assertETHRateLimiterState(0, true, 100 ether, 1 ether);

        // reverted when exceeding the capacity
        (_proof, _root) = messageProveAndRelayPrepare(alice, bob, 101 ether, 0, "");
        hevm.expectRevert(
            abi.encodeWithSelector(RateLimiter.ErrRateLimitCapacityExceeded.selector, 100 ether, 101 ether)
        );
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 101 ether, 0, "", _proof, _root);
    }

    function test_proveAndRelayMessage_ethRateLimitRefund_succeeds() external {
        hevm.deal(address(l1CrossDomainMessenger), 1000 ether);
        _setETHRateLimiterConfig(true, 100 ether, 1 ether);

        address _to = address(this);
        bytes32 _hash = keccak256(_encodeXDomainCalldata(alice, _to, 60 ether, 0, ""));
        (bytes32[32] memory _proof, bytes32 _root) = messageProveAndRelayPrepare(alice, _to, 60 ether, 0, "");

        // the tokens are given back when the relay fails
        revertOnReceive = true;
        hevm.expectEmit(true, false, false, true);
        emit ICrossDomainMessenger.FailedRelayedMessage(_hash);
        l1CrossDomainMessenger.proveAndRelayMessage(alice, _to, 60 ether, 0, "", _proof, _root);
        _assertETHRateLimiterState(100 ether, true, 100 ether, 1 ether);

        // and consumed once the relay succeeds
        revertOnReceive = false;
        hevm.expectEmit(true, false, false, true);
        emit IL1CrossDomainMessenger.ETHRateLimitConsumed(_hash, 60 ether, 40 ether);
        l1CrossDomainMessenger.proveAndRelayMessage(alice, _to, 60 ether, 0, "", _proof, _root);
        _assertETHRateLimiterState(40 ether, true, 100 ether, 1 ether);
    }

    function test_proveAndRelayMessage_zeroValueThrottled_succeeds() external {
        hevm.deal(address(l1CrossDomainMessenger), 1 ether);
        _setETHRateLimiterConfig(true, 0, 0);

        (bytes32[32] memory _proof, bytes32 _root) = messageProveAndRelayPrepare(alice, bob, 1, 0, "");
        hevm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitCapacityExceeded.selector, 0, 1));
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1, 0, "", _proof, _root);

        // messages without ETH, such as ERC20 withdrawals, are not limited
        bytes32 _hash = keccak256(_encodeXDomainCalldata(alice, bob, 0, 0, "message"));
        (_proof, _root) = messageProveAndRelayPrepare(alice, bob, 0, 0, "message");
        Vm(HEVM_ADDRESS).recordLogs();
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 0, 0, "message", _proof, _root);
        assertBoolEq(true, l1CrossDomainMessenger.finalizedWithdrawals(_hash));
        _assertNoETHRateLimitConsumed();
    }

    function test_proveAndRelayMessage_ethRateLimitDisabled_succeeds() external {
        hevm.deal(address(l1CrossDomainMessenger), 1000 ether);

        (bytes32[32] memory _proof, bytes32 _root) = messageProveAndRelayPrepare(alice, bob, 1000 ether, 0, "");
        Vm(HEVM_ADDRESS).recordLogs();
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1000 ether, 0, "", _proof, _root);
        _assertNoETHRateLimitConsumed();
        _assertETHRateLimiterState(0, false, 0, 0);
    }

    function test_proveAndRelayMessage_ethRateLimitCheckedLast_reverts() external {
        hevm.deal(address(l1CrossDomainMessenger), 1 ether);
        _setETHRateLimiterConfig(true, 0, 0);

        bytes32 _hash = keccak256(_encodeXDomainCalldata(alice, bob, 1 ether, 0, ""));
        (, bytes32[32] memory _proof, bytes32 _root) = ffi.getProveWithdrawalTransactionInputs(_hash);
        hevm.expectRevert("Messenger: withdrawalRoot not finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "", _proof, _root);

        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeCall(IRollup.withdrawalRoots, (_root)),
            abi.encode(true)
        );
        hevm.expectRevert("Messenger: invalid withdrawal inclusion proof");
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "invalid", _proof, _root);

        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(true);
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "", _proof, _root);
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(false);

        hevm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrRateLimitCapacityExceeded.selector, 0, 1 ether));
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "", _proof, _root);

        // relayed with the limit disabled, then replayed with the limit throttled
        _setETHRateLimiterConfig(false, 0, 0);
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "", _proof, _root);
        _setETHRateLimiterConfig(true, 0, 0);
        hevm.expectRevert("Messenger: withdrawal has already been finalized");
        l1CrossDomainMessenger.proveAndRelayMessage(alice, bob, 1 ether, 0, "", _proof, _root);
    }

    function test_setETHRateLimiterConfig_owner_succeeds() external {
        // enabled with a full bucket
        hevm.expectEmit(false, false, false, true);
        emit IL1CrossDomainMessenger.UpdateETHRateLimiterConfig(true, 100 ether, 1 ether);
        _setETHRateLimiterConfig(true, 100 ether, 1 ether);
        _assertETHRateLimiterState(100 ether, true, 100 ether, 1 ether);

        // the owner can loosen and disable the limit
        _setETHRateLimiterConfig(true, 200 ether, 2 ether);
        _assertETHRateLimiterState(100 ether, true, 200 ether, 2 ether);
        _setETHRateLimiterConfig(false, 0, 0);
        _assertETHRateLimiterState(0, false, 0, 0);

        hevm.expectRevert(abi.encodeWithSelector(RateLimiter.ErrInvalidRateLimitConfig.selector, true, 1, 2));
        _setETHRateLimiterConfig(true, 1, 2);
    }

    function test_setETHRateLimiterConfig_rateLimitAdmin_succeeds() external {
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateRateLimitAdmin(alice);

        hevm.startPrank(alice);
        // the admin can enable the limit, and lower the capacity and the rate
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 100 ether, 1 ether));
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 50 ether, 1 ether));
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 50 ether, 0.5 ether));
        _assertETHRateLimiterState(50 ether, true, 50 ether, 0.5 ether);

        // but cannot raise or disable it
        hevm.expectRevert(IL1CrossDomainMessenger.ErrRateLimitAdminCanOnlyTighten.selector);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 51 ether, 0.5 ether));
        hevm.expectRevert(IL1CrossDomainMessenger.ErrRateLimitAdminCanOnlyTighten.selector);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 50 ether, 0.6 ether));
        hevm.expectRevert(IL1CrossDomainMessenger.ErrRateLimitAdminCanOnlyTighten.selector);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(false, 0, 0));

        // the admin can throttle all ETH withdrawals
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 0, 0));
        _assertETHRateLimiterState(0, true, 0, 0);
        hevm.stopPrank();

        // the role is revoked
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateRateLimitAdmin(address(0));
        hevm.expectRevert(IL1CrossDomainMessenger.ErrNotOwnerOrRateLimitAdmin.selector);
        hevm.prank(alice);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 0, 0));
    }

    function test_setETHRateLimiterConfig_notOwnerOrRateLimitAdmin_reverts() external {
        hevm.expectRevert(IL1CrossDomainMessenger.ErrNotOwnerOrRateLimitAdmin.selector);
        hevm.prank(bob);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(true, 0, 0));
    }

    function test_updateRateLimitAdmin_succeeds(address _rateLimitAdmin) external {
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(bob);
        l1CrossDomainMessenger.updateRateLimitAdmin(_rateLimitAdmin);

        hevm.expectEmit(true, true, false, true);
        emit IL1CrossDomainMessenger.UpdateRateLimitAdmin(address(0), _rateLimitAdmin);
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateRateLimitAdmin(_rateLimitAdmin);
        assertEq(l1CrossDomainMessenger.rateLimitAdmin(), _rateLimitAdmin);
    }

    function _setETHRateLimiterConfig(bool _isEnabled, uint128 _capacity, uint128 _rate) internal {
        hevm.prank(multisig);
        l1CrossDomainMessenger.setETHRateLimiterConfig(RateLimiter.Config(_isEnabled, _capacity, _rate));
    }

    function _assertETHRateLimiterState(uint256 _tokens, bool _isEnabled, uint256 _capacity, uint256 _rate) internal {
        RateLimiter.TokenBucket memory _state = l1CrossDomainMessenger.currentETHRateLimiterState();
        assertEq(_state.tokens, _tokens);
        assertBoolEq(_state.isEnabled, _isEnabled);
        assertEq(_state.capacity, _capacity);
        assertEq(_state.rate, _rate);
    }

    function _assertNoETHRateLimitConsumed() internal {
        Vm.Log[] memory _logs = Vm(HEVM_ADDRESS).getRecordedLogs();
        for (uint256 i = 0; i < _logs.length; i++) {
            if (_logs[i].topics.length > 0) {
                assertTrue(_logs[i].topics[0] != IL1CrossDomainMessenger.ETHRateLimitConsumed.selector);
            }
        }
    }
}
