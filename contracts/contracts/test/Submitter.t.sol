// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

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

    function test_getTurn() external {
        uint256 start;
        uint256 end;

        hevm.warp(1841071100);
        emit log_uint(block.timestamp);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000A)
        );
        emit log_address(address(0x000000000000000000000000000000000000000A));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000b)
        );
        emit log_address(address(0x000000000000000000000000000000000000000b));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000C)
        );
        emit log_address(address(0x000000000000000000000000000000000000000C));
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841074440);
        emit log_uint(block.timestamp);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000A)
        );
        emit log_address(address(0x000000000000000000000000000000000000000A));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000b)
        );
        emit log_address(address(0x000000000000000000000000000000000000000b));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(
            address(0x000000000000000000000000000000000000000C)
        );
        emit log_address(address(0x000000000000000000000000000000000000000C));
        emit log_uint(start);
        emit log_uint(end);
        emit log("");
    }

    function test_getNextSubmitter() external {
        address submitter;
        uint256 start;
        uint256 end;

        hevm.warp(1841071100);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841072220);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841073330);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841074440);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841075550);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841076660);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getNextSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");
    }
}
