// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IRecord} from "../L2/staking/IRecord.sol";

contract RecordTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    /**
     * @notice setOracleAddress: check params
     */
    function testSetOracleAddress() public {
        hevm.expectRevert("invalid oracle address");
        hevm.prank(multisig);
        record.setOracleAddress(address(0));
    }

    /**
     * @notice setOracleAddress: check owner
     */
    function testOwner() public {
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(alice);
        record.setOracleAddress(address(0));
    }

    /**
     * @notice setLatestRewardEpochBlock: check params
     */
    function testSetLatestRewardEpochBlock() public {
        hevm.expectRevert("only oracle allowed");
        hevm.prank(multisig);
        record.setLatestRewardEpochBlock(0);

        hevm.expectRevert("invalid latest block");
        hevm.prank(oracleAddress);
        record.setLatestRewardEpochBlock(0);

        hevm.prank(oracleAddress);
        record.setLatestRewardEpochBlock(100);

        hevm.expectRevert("already set");
        hevm.prank(oracleAddress);
        record.setLatestRewardEpochBlock(100);
    }

    /**
     * @notice recordFinalizedBatchSubmissions: check owner
     */
    function testRecordFinalizedBatchSubmissions() public {
        IRecord.BatchSubmission memory submission = IRecord.BatchSubmission(
            0,
            address(0),
            0,
            0,
            0,
            0
        );

        IRecord.BatchSubmission[]
            memory submissions = new IRecord.BatchSubmission[](1);
        submissions[0] = submission;

        hevm.expectRevert("only oracle allowed");
        hevm.prank(multisig);
        record.recordFinalizedBatchSubmissions(submissions);
    }

    /**
     * @notice recordRollupEpochs: check owner
     */
    function testRecordRollupEpochs() public {
        IRecord.RollupEpochInfo memory epochInfo = IRecord.RollupEpochInfo(
            0,
            address(0),
            0,
            0
        );

        IRecord.RollupEpochInfo[]
            memory epochInfos = new IRecord.RollupEpochInfo[](1);
        epochInfos[0] = epochInfo;

        hevm.expectRevert("only oracle allowed");
        hevm.prank(multisig);
        record.recordRollupEpochs(epochInfos);
    }

    /**
     * @notice recordRewardEpochs: check owner
     */
    function testRecordRewardEpochs() public {
        uint256 sequencerSize = sequencer.getSequencerSet2Size();
        address[] memory sequencers = sequencer.getSequencerSet2();
        uint256[] memory sequencerBlocks = new uint256[](sequencerSize);
        uint256[] memory sequencerRatios = new uint256[](sequencerSize);
        uint256[] memory sequencerCommissions = new uint256[](sequencerSize);

        for (uint i = 0; i < sequencerSize; i++) {
            sequencerBlocks[i] = 0;
            sequencerRatios[i] = SEQUENCER_RATIO_PRECISION / sequencerSize;
            sequencerCommissions[i] = 1;
        }

        IRecord.RewardEpochInfo[]
            memory rewardEpochInfos = new IRecord.RewardEpochInfo[](1);

        rewardEpochInfos[0] = IRecord.RewardEpochInfo(
            0,
            1,
            sequencers,
            sequencerBlocks,
            sequencerRatios,
            sequencerCommissions
        );

        hevm.expectRevert("only oracle allowed");
        hevm.prank(multisig);
        record.recordRewardEpochs(rewardEpochInfos);
    }

    /**
     * @notice getBatchSubmissions: check params
     */
    function testGetBatchSubmissions() public {
        hevm.expectRevert("invalid index");
        hevm.prank(oracleAddress);
        record.getBatchSubmissions(2, 1);
    }

    /**
     * @notice getRollupEpochs: check params
     */
    function testGetRollupEpochs() public {
        hevm.expectRevert("invalid index");
        hevm.prank(oracleAddress);
        record.getRollupEpochs(2, 1);
    }

    /**
     * @notice getRewardEpochs: check params
     */
    function testGetRewardEpochs() public {
        hevm.expectRevert("invalid index");
        hevm.prank(oracleAddress);
        record.getRewardEpochs(2, 1);
    }
}
