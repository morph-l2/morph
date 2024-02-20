// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L1MessageQueueWithGasPriceOracle} from "../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {IRollup} from "../L1/rollup/IRollup.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";

contract L1CrossDomainMessengerTest is L1MessageBaseTest {
    event UpdateMaxReplayTimes(
        uint256 oldMaxReplayTimes,
        uint256 newMaxReplayTimes
    );

    uint256 L1CrossDomainMessenger_provenWithdrawals_slot = 251;
    address refundAddress = address(2048);

    function testProveWithdrawalTransaction_relayMessage() external {
        // tx msg set
        address from = address(alice);
        address to = address(bob);
        uint256 value = 10;
        hevm.deal(address(l1CrossDomainMessenger), value);

        uint256 nonce = 0;
        bytes memory message = "";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        // get proof from ffi
        _appendMessageHash(_xDomainCalldataHash);
        (bytes32 wdHashRes, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(_xDomainCalldataHash);
        assertEq(_xDomainCalldataHash, wdHashRes);
        assertEq(getTreeRoot(), wdRoot);

        // prove without rollup
        hevm.expectRevert("Messenger: do not submit withdrawalRoot");
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // mock call rollup withdrawal root submitted success
        uint256 withdrawalBatchIndex = 1;
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(IRollup.withdrawalRoots.selector, wdRoot),
            abi.encode(withdrawalBatchIndex)
        );
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        // prove again
        hevm.expectRevert("Messenger: withdrawal hash has already been proven");
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            wdProof,
            wdRoot
        );

        hevm.expectRevert(
            "Messenger: proven withdrawal finalization period has not elapsed"
        );
        l1CrossDomainMessenger.relayMessage(from, to, value, nonce, message);

        // warp finalization period
        (, uint256 provenTime, ) = l1CrossDomainMessenger.provenWithdrawals(
            _xDomainCalldataHash
        );
        hevm.warp(provenTime + FINALIZATION_PERIOD_SECONDS + 1);

        hevm.expectRevert("Messenger: batch not verified");
        l1CrossDomainMessenger.relayMessage(from, to, value, nonce, message);

        // finalize batch
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(
                IRollup.finalizedStateRoots.selector,
                withdrawalBatchIndex
            ),
            abi.encode(bytes32(uint256(1)))
        );
        uint256 balanceBefore = address(bob).balance;
        l1CrossDomainMessenger.relayMessage(from, to, value, nonce, message);
        assertEq(balanceBefore + value, address(bob).balance);
    }

    function testReplayMessage(uint256 exceedValue) external {
        hevm.deal(address(this), 1 ether);
        // tx msg set
        exceedValue = bound(exceedValue, 1, address(this).balance / 2);
        address from = address(this);
        address to = address(bob);
        uint256 value = 0;
        bytes memory message = "";
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            from,
            to,
            value,
            nonce,
            message
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );

        // updateMaxReplayTimes to 0
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(0);

        // append a message
        l1CrossDomainMessenger.sendMessage{value: 100}(
            to,
            value,
            message,
            gas,
            refundAddress
        );

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
        l1CrossDomainMessenger.replayMessage(
            from,
            to,
            value,
            nonce,
            message,
            defaultGasLimit,
            refundAddress
        );

        hevm.prank(multisig);
        uint256 _fee = l1MessageQueueWithGasPriceOracle.l2BaseFee() * defaultGasLimit;

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

        // test replay list
        // 1. send a message with nonce 2
        // 2. replay 3 times
        hevm.startPrank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(0);
        l1CrossDomainMessenger.updateMaxReplayTimes(100);
        hevm.stopPrank();
        l1CrossDomainMessenger.sendMessage{value: 100}(
            address(0),
            100,
            new bytes(0),
            defaultGasLimit,
            refundAddress
        );
        bytes32 hash = keccak256(
            abi.encodeWithSignature(
                "relayMessage(address,address,uint256,uint256,bytes)",
                address(this),
                address(0),
                100,
                2,
                new bytes(0)
            )
        );
        (uint256 _replayTimes, uint256 _lastIndex) = l1CrossDomainMessenger
            .replayStates(hash);
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
            (_replayTimes, _lastIndex) = l1CrossDomainMessenger.replayStates(
                hash
            );
            assertEq(_replayTimes, i + 1);
            assertEq(_lastIndex, i + 3);
            assertEq(l1CrossDomainMessenger.prevReplayIndex(i + 3), i + 2 + 1);
            for (uint256 j = 0; j <= i; j++) {
                assertEq(
                    l1CrossDomainMessenger.prevReplayIndex(i + 3 - j),
                    i + 2 - j + 1
                );
            }
        }
    }

    function testForbidCallMessageQueueFromL2() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1MessageQueueWithGasPriceOracle);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(
                IRollup.withdrawalRoots.selector,
                withdrawalRoot
            ),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call message queue");
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            withdrawalProof,
            withdrawalRoot
        );
    }

    function testForbidCallSelfFromL2() external {
        // withdrawal tx
        address from = address(alice);
        address to = address(l1CrossDomainMessenger);
        uint256 value = 0;
        uint256 nonce = 0;
        bytes memory message = "send message";
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(from, to, value, nonce, message)
        );
        _appendMessageHash(_xDomainCalldataHash);
        bytes32[32] memory withdrawalProof;
        bytes32 withdrawalRoot = getTreeRoot();
        uint256 withdrawalBatchIndex = 1;
        // mock call withdrawalRoots
        hevm.mockCall(
            address(l1CrossDomainMessenger.rollup()),
            abi.encodeWithSelector(
                IRollup.withdrawalRoots.selector,
                withdrawalRoot
            ),
            abi.encode(withdrawalBatchIndex)
        );

        hevm.expectRevert("Messenger: Forbid to call self");
        l1CrossDomainMessenger.proveMessage(
            from,
            to,
            value,
            nonce,
            message,
            withdrawalProof,
            withdrawalRoot
        );
    }

    function test_sendMessage() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        hevm.expectEmit(true, true, true, true);
        emit SentMessage(sender, to, value, nonce, gas, data);

        hevm.expectCall(
            address(l1MessageQueueWithGasPriceOracle),
            abi.encodeWithSelector(
                l1MessageQueueWithGasPriceOracle
                    .appendCrossDomainMessage
                    .selector,
                Predeploys.L2_CROSS_DOMAIN_MESSENGER,
                gas,
                _xDomainCalldata
            )
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);
    }

    function test_sendMessage_value() external {
        address sender = address(this);
        address to = address(bob);
        bytes memory data = "send message";
        hevm.deal(sender, 10 ether);

        // send value zero
        uint256 value = 0;
        uint256 nonce = l1MessageQueueWithGasPriceOracle
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        uint256 gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        l1CrossDomainMessenger.sendMessage(to, value, data, gas);

        // send value not zero
        // set base fee to 100
        hevm.prank(multisig);
        l1MessageQueueWithGasPriceOracle.setL2BaseFee(100);

        value = 1 ether;
        nonce = l1MessageQueueWithGasPriceOracle.nextCrossDomainMessageIndex();
        _xDomainCalldata = _encodeXDomainCalldata(
            sender,
            to,
            value,
            nonce,
            data
        );
        gas = l1MessageQueueWithGasPriceOracle.calculateIntrinsicGasFee(
            _xDomainCalldata
        );
        hevm.expectRevert("Insufficient msg.value");
        l1CrossDomainMessenger.sendMessage{value: 1 ether}(
            to,
            value,
            data,
            gas
        );
        // give enought value
        uint256 fee = l1MessageQueueWithGasPriceOracle
            .estimateCrossDomainMessageFee(gas);
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee}(
            to,
            value,
            data,
            gas
        );

        // send more value with refund address
        fee = l1MessageQueueWithGasPriceOracle.estimateCrossDomainMessageFee(
            gas
        );
        l1CrossDomainMessenger.sendMessage{value: 1 ether + fee * 2}(
            to,
            value,
            data,
            gas,
            refundAddress
        );
        assertEq(address(refundAddress).balance, fee);
    }

    function testUpdateMaxReplayTimes(uint256 _maxReplayTimes) external {
        // not owner, revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        hevm.stopPrank();

        hevm.expectEmit(false, false, false, true);
        emit UpdateMaxReplayTimes(3, _maxReplayTimes);

        assertEq(l1CrossDomainMessenger.maxReplayTimes(), 3);
        hevm.prank(multisig);
        l1CrossDomainMessenger.updateMaxReplayTimes(_maxReplayTimes);
        assertEq(l1CrossDomainMessenger.maxReplayTimes(), _maxReplayTimes);
    }

    function testSetPause() external {
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
        l1CrossDomainMessenger.sendMessage(
            address(0),
            0,
            new bytes(0),
            defaultGasLimit
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.sendMessage(
            address(0),
            0,
            new bytes(0),
            defaultGasLimit,
            address(0)
        );

        (, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi
            .getProveWithdrawalTransactionInputs(bytes32(uint256(1)));
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.proveMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0),
            wdProof,
            wdRoot
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.relayMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0)
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.replayMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0),
            0,
            address(0)
        );
        hevm.expectRevert("Pausable: paused");
        l1CrossDomainMessenger.dropMessage(
            address(0),
            address(0),
            0,
            0,
            new bytes(0)
        );

        // unpause
        hevm.prank(multisig);
        l1CrossDomainMessenger.setPause(false);
        assertBoolEq(false, l1CrossDomainMessenger.paused());
    }
}
