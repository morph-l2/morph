// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Staking} from "../L1/staking/Staking.sol";
import {CrossDomainMessenger} from "../universal/CrossDomainMessenger.sol";
import {Sequencer} from "../universal/Sequencer.sol";
import {L2Sequencer} from "../L2/L2Sequencer.sol";
import {Submitter} from "../L2/Submitter.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {Types} from "../libraries/Types.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {Gov} from "../L2/Gov.sol";
import "forge-std/console.sol";
import "./CommonTest.t.sol";

contract Submitter_Test is Staking_Initializer {
    function test_ackRollup() external {
        uint256 batchIndex = 0;
        uint256 batchStartBlock = 0;
        uint256 batchEndBlock = 10;
        uint256 rollupTime = 1700001000;

        bytes memory data = abi.encodeWithSelector(
            Submitter.ackRollup.selector,
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );
        uint256 _nonce = L1Messenger.messageNonce();
        address _sender = AddressAliasHelper.applyL1ToL2Alias(
            address(L1Messenger)
        );

        vm.expectEmit(true, true, true, true);
        emit ACKRollup(
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );

        bytes32 versionedHash = Hashing.hashCrossDomainMessageV1(
            _nonce,
            address(l1Sequencer),
            address(l2Submitter),
            0,
            defultGasLimit,
            data
        );

        // TODO: test send message

        vm.expectEmit(true, true, true, true);
        emit RelayedMessage(versionedHash);

        vm.prank(_sender);
        L2Messenger.relayMessage(
            _nonce,
            address(l1Sequencer),
            address(l2Submitter),
            0,
            defultGasLimit,
            data
        );
        assertEq(L2Messenger.failedMessages(versionedHash), false);
        assertEq(L2Messenger.successfulMessages(versionedHash), true);

        Types.BatchInfo memory batchInfo = l2Submitter.getConfirmedBatch(
            batchIndex
        );
        assertEq(l2Submitter.nextBatchIndex(), batchIndex + 1);
        assertEq(batchInfo.submitter, address(alice));
        assertEq(batchInfo.startBlock, batchStartBlock);
        assertEq(batchInfo.endBlock, batchEndBlock);
        assertEq(batchInfo.rollupTime, rollupTime);
    }

    function test_updateEpoch() external {
        vm.expectEmit(true, true, true, true);
        emit EpochUpdated(l2Gov.rollupEpoch(), SEQUENCER_SIZE);

        l2Submitter.updateEpochExternal();
    }
}
