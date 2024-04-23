// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {IMorphToken} from "../system/IMorphToken.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {ISequencer} from "./ISequencer.sol";
import {IDistribute} from "./IDistribute.sol";
import {IGov} from "./IGov.sol";
import {IRecord} from "./IRecord.sol";

contract Record is IRecord, OwnableUpgradeable {
    // MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;
    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;
    // sequencer contract address
    address public immutable SEQUENCER_CONTRACT;
    // distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;
    // gov contract address
    address public immutable GOV_CONTRACT;
    // oracle address
    address public ORACLE;

    // sequencers reward ratio precision
    uint256 private immutable RATIO_PRECISION = 10000;

    // If the sequencer set or rollup epoch changed, reset the submitter round
    // mapping(batch_index => batch_submission)
    mapping(uint256 => BatchSubmission) public batchSubmissions;
    // mapping(rollup_epoch_index => rollup_epoch_info)
    mapping(uint256 => RollupEpochInfo) public rollupEpochs;
    // mapping(reward_epoch_index => reward_epoch_info)
    mapping(uint256 => RewardEpochInfo) public rewardpEpochs;
    // next batch submission index
    uint256 public override nextBatchSubmissionIndex;
    // next rollup epoch index
    uint256 public override nextRollupEpochIndex;
    // next reward epoch index
    uint256 public override nextRewardEpochIndex;

    /*********************** modifiers **************************/

    /// @notice Only prover allowed.
    modifier onlyOracle() {
        require(msg.sender == ORACLE, "only oracle allowed");
        _;
    }

    /*********************** Constructor **************************/

    /**
     * @notice constructor
     */
    constructor() {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
        GOV_CONTRACT = Predeploys.GOV;
    }

    /*********************** Init **************************/

    /**
     * @notice Initializer.
     * @param _admin    params admin
     * @param _oracle   oracle address
     */
    function initialize(address _admin, address _oracle) public initializer {
        require(_oracle != address(0), "invalid oracle address");
        ORACLE = _oracle;

        // transfer owner to admin
        _transferOwnership(_admin);
    }

    /*********************** External Functions **************************/

    /**
     * @notice set oracle address
     * @param _oracle     oracle address
     */
    function setOracleAddress(address _oracle) external onlyOwner {
        require(_oracle != address(0), "invalid oracle address");
        ORACLE = _oracle;
    }

    /**
     * @notice record batch submissions
     */
    function recordFinalizedBatchSubmissions(
        BatchSubmission[] calldata _batchSubmissions
    ) external onlyOracle {
        for (uint256 i = 0; i < _batchSubmissions.length; i++) {
            require(
                _batchSubmissions[i].index == nextBatchSubmissionIndex + i,
                "invalid index"
            );
            // TODO: check more
            batchSubmissions[_batchSubmissions[i].index] = BatchSubmission(
                _batchSubmissions[i].index,
                _batchSubmissions[i].submitter,
                _batchSubmissions[i].startBlock,
                _batchSubmissions[i].endBlock,
                _batchSubmissions[i].rollupTime
            );
        }
    }

    /**
     * @notice record epochs
     */
    function recordRollupEpochs(
        RollupEpochInfo[] calldata _rollupEpochs
    ) external onlyOracle {
        for (uint256 i = 0; i < _rollupEpochs.length; i++) {
            require(
                _rollupEpochs[i].index == nextRollupEpochIndex + i,
                "invalid index"
            );
            // TODO: check more
            rollupEpochs[_rollupEpochs[i].index] = RollupEpochInfo(
                _rollupEpochs[i].index,
                _rollupEpochs[i].submitter,
                _rollupEpochs[i].startTime,
                _rollupEpochs[i].endTime
            );
        }
    }

    /**
     * @notice record epochs
     */
    function recordRewardEpochs(
        RewardEpochInfo[] calldata _rewardEpochs
    ) external onlyOracle {
        require(
            nextRewardEpochIndex + _rewardEpochs.length - 1 <
                IL2Staking(L2_STAKING_CONTRACT).currentEpoch(),
            "future data cannot be uploaded"
        );
        IMorphToken(MORPH_TOKEN_CONTRACT).mintInflations(
            nextRewardEpochIndex + _rewardEpochs.length - 1
        );

        for (uint256 i = 0; i < _rewardEpochs.length; i++) {
            uint256 dataLen = _rewardEpochs[i].sequencers.length;
            uint256 index = _rewardEpochs[i].index;
            require(index == nextRewardEpochIndex + i, "invalid epoch index");
            require(
                _rewardEpochs[i].sequencerBlocks.length == dataLen &&
                    _rewardEpochs[i].sequencerRatios.length == dataLen &&
                    _rewardEpochs[i].sequencerComissions.length == dataLen,
                "invalid data length"
            );

            rewardpEpochs[index] = RewardEpochInfo(
                index,
                _rewardEpochs[i].blockCount,
                new address[](dataLen),
                new uint256[](dataLen),
                new uint256[](dataLen),
                new uint256[](dataLen)
            );

            uint256 inflationAmount = IMorphToken(MORPH_TOKEN_CONTRACT)
                .inflation(index);
            uint256 blockCount = 0;
            uint256 ratioSum = 0;
            uint256[] memory delegatorRewards = new uint256[](dataLen);
            uint256[] memory commissions = new uint256[](dataLen);
            for (uint256 j = 0; j < dataLen; j++) {
                require(
                    _rewardEpochs[i].sequencerComissions[j] <= 20,
                    "invalid sequencers comission"
                );
                ratioSum += _rewardEpochs[i].sequencerRatios[j];
                blockCount += _rewardEpochs[i].sequencerBlocks[j];

                // compute rewards per sequencer
                uint256 reward = (inflationAmount *
                    _rewardEpochs[i].sequencerRatios[j]) / RATIO_PRECISION;
                commissions[i] =
                    (reward * _rewardEpochs[i].sequencerComissions[j]) /
                    100;
                delegatorRewards[i] = reward - commissions[i];
            }
            require(
                blockCount == _rewardEpochs[i].blockCount,
                "invalid sequencers blocks"
            );
            require(ratioSum <= RATIO_PRECISION, "invalid sequencers ratios");

            // update sequecers reward data
            IDistribute(DISTRIBUTE_CONTRACT).updateEpochReward(
                index,
                _rewardEpochs[i].sequencers,
                delegatorRewards,
                commissions
            );
        }

        nextRewardEpochIndex += _rewardEpochs.length;
    }

    /*********************** External View Functions **************************/

    /**
     * @notice getBatchSubmissions
     * @param start start index
     * @param end   end index
     */
    function getBatchSubmissions(
        uint256 start,
        uint256 end
    ) external view returns (BatchSubmission[] memory res) {
        require(end >= start, "invalid index");
        res = new BatchSubmission[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = batchSubmissions[i];
        }
    }

    /**
     * @notice get rollup epochs
     * @param start start index
     * @param end   end index
     */
    function getRollupEpochs(
        uint256 start,
        uint256 end
    ) external view returns (RollupEpochInfo[] memory res) {
        require(end >= start, "invalid index");
        res = new RollupEpochInfo[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = rollupEpochs[i];
        }
    }

    /**
     * @notice get reward epochs
     * @param start start index
     * @param end   end index
     */
    function getRewardEpochs(
        uint256 start,
        uint256 end
    ) external view returns (RewardEpochInfo[] memory res) {
        require(end >= start, "invalid index");
        res = new RewardEpochInfo[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = rewardpEpochs[i];
        }
    }
}
