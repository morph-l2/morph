// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {Submitter} from "../L2/submitter/Submitter.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";

contract SubmitterTest is L2StakingBaseTest {
    event ACKRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    );

    event EpochUpdated(uint256 interval, uint256 sequencersLen);

    function testAckRollup() external {
        uint256 batchIndex = 0;
        uint256 batchStartBlock = 0;
        uint256 batchEndBlock = 10;
        uint256 rollupTime = 1700001000;

        bytes memory _message = abi.encodeWithSelector(
            Submitter.ackRollup.selector,
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );
        uint256 _nonce = 0;
        address _sender = AddressAliasHelper.applyL1ToL2Alias(
            l2CrossDomainMessenger.counterpart()
        );

        hevm.expectEmit(true, true, true, true);
        emit ACKRollup(
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(
                NON_ZERO_ADDRESS,
                address(l2Submitter),
                0,
                _nonce,
                _message
            )
        );

        hevm.expectEmit(true, true, true, true);
        emit RelayedMessage(_xDomainCalldataHash);

        hevm.prank(_sender);
        l2CrossDomainMessenger.relayMessage(
            NON_ZERO_ADDRESS,
            address(l2Submitter),
            0,
            _nonce,
            _message
        );
        assertTrue(
            l2CrossDomainMessenger.isL1MessageExecuted(_xDomainCalldataHash)
        );

        Types.BatchInfo memory batchInfo = l2Submitter.getConfirmedBatch(
            batchIndex
        );
        assertEq(l2Submitter.nextBatchIndex(), batchIndex + 1);
        assertEq(batchInfo.submitter, address(alice));
        assertEq(batchInfo.startBlock, batchStartBlock);
        assertEq(batchInfo.endBlock, batchEndBlock);
        assertEq(batchInfo.rollupTime, rollupTime);
    }

    function testUpdateEpoch() external {
        hevm.expectEmit(true, true, true, true);
        emit EpochUpdated(l2Gov.rollupEpoch(), SEQUENCER_SIZE);

        l2Submitter.updateEpochExternal();
    }
}
