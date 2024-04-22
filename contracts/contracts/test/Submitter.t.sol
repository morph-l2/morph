// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2SequencerTest} from "./L2Sequencer.t.sol";
import {Types} from "../libraries/common/Types.sol";
import {Submitter} from "../L2/submitter/Submitter.sol";
import {AddressAliasHelper} from "../libraries/common/AddressAliasHelper.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract SubmitterTest is L2SequencerTest {
    event ACKRollup(
        uint256 indexed batchIndex,
        address indexed submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    );

    function setUp() public virtual override {
        super.setUp();

        hevm.mockCall(
            address(l2Sequencer.messenger()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Sequencer.OTHER_SEQUENCER()))
        );

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerBLSKeys.push(sequencerInfo.blsKey);
            sequencerInfos[i] = sequencerInfo;
        }
        version++;
        hevm.prank(address(l2CrossDomainMessenger));
        // updateSequencers
        l2Sequencer.updateSequencers(version, sequencerInfos);
        assertEq(l2Sequencer.currentVersion(), version);
    }

    function testAckRollup() external {
        uint256 batchIndex = 0;
        uint256 batchStartBlock = 0;
        uint256 batchEndBlock = 10;
        uint256 rollupTime = 1700001000;

        hevm.prank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit ACKRollup(
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );

        l2Submitter.ackRollup(
            batchIndex,
            alice,
            batchStartBlock,
            batchEndBlock,
            rollupTime
        );

        Types.BatchInfo memory batchInfo = l2Submitter.getConfirmedBatch(
            batchIndex
        );
        assertEq(l2Submitter.nextBatchIndex(), batchIndex + 1);
        assertEq(batchInfo.submitter, address(alice));
        assertEq(batchInfo.startBlock, batchStartBlock);
        assertEq(batchInfo.endBlock, batchEndBlock);
        assertEq(batchInfo.rollupTime, rollupTime);
        hevm.stopPrank();
    }

    function test_getTurn() external {
        uint256 start;
        uint256 end;

        hevm.warp(1841071100);
        emit log_uint(block.timestamp);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 0)));
        emit log_address(address(uint160(beginSeq + 0)));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 1)));
        emit log_address(address(uint160(beginSeq + 1)));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 2)));
        emit log_address(address(uint160(beginSeq + 2)));
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841074440);
        emit log_uint(block.timestamp);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 0)));
        emit log_address(address(uint160(beginSeq + 0)));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 1)));
        emit log_address(address(uint160(beginSeq + 1)));
        emit log_uint(start);
        emit log_uint(end);
        (start, end) = l2Submitter.getTurn(address(uint160(beginSeq + 2)));
        emit log_address(address(uint160(beginSeq + 2)));
        emit log_uint(start);
        emit log_uint(end);
        emit log("");
    }

    function test_getCurrentSubmitter() external {
        address submitter;
        uint256 start;
        uint256 end;

        hevm.warp(1841071100);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841072220);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841073330);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841074440);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841075550);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");

        hevm.warp(1841076660);
        emit log_uint(block.timestamp);
        (submitter, start, end) = l2Submitter.getCurrentSubmitter();
        emit log_address(submitter);
        emit log_uint(start);
        emit log_uint(end);
        emit log("");
    }
}
