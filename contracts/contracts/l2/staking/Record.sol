// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IMorphToken} from "../system/IMorphToken.sol";
import {IL2Staking} from "./IL2Staking.sol";
import {IDistribute} from "./IDistribute.sol";
import {IRecord} from "./IRecord.sol";

contract Record is IRecord, OwnableUpgradeable {
    /*************
     * Constants *
     *************/

    /// @notice inflation rate precision
    uint256 private constant PRECISION = 1e8;

    /// @notice MorphToken contract address
    address public immutable MORPH_TOKEN_CONTRACT;

    /// @notice l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// @notice sequencer contract address
    address public immutable SEQUENCER_CONTRACT;

    /// @notice distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;

    /// @notice gov contract address
    address public immutable GOV_CONTRACT;

    /*************
     * Variables *
     *************/

    /// @notice oracle address
    address public oracle;

    /// @notice If the sequencer set or rollup epoch changed, reset the submitter round
    mapping(uint256 batchIndex => BatchSubmission) public batchSubmissions;

    /// @notice rollup epoch info
    mapping(uint256 rollupEpochIndex => RollupEpochInfo) public rollupEpochs;

    /// @notice reward epoch info
    mapping(uint256 rewardEpochIndex => RewardEpochInfo) public rewardEpochs;

    /// @notice next batch submission index
    uint256 public override nextBatchSubmissionIndex;

    /// @notice next rollup epoch index
    uint256 public override nextRollupEpochIndex;

    /// @notice next reward epoch index
    uint256 public override nextRewardEpochIndex;

    /// @notice latest reward epoch block
    uint256 public override latestRewardEpochBlock;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Only prover allowed.
    modifier onlyOracle() {
        require(_msgSender() == oracle, "only oracle allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        MORPH_TOKEN_CONTRACT = Predeploys.MORPH_TOKEN;
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
        GOV_CONTRACT = Predeploys.GOV;
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    /// @notice Initializer.
    /// @param _owner                       owner
    /// @param _oracle                      oracle address
    /// @param _nextBatchSubmissionIndex    next batch submission index
    function initialize(address _owner, address _oracle, uint256 _nextBatchSubmissionIndex) public initializer {
        require(_owner != address(0), "invalid owner address");
        require(_nextBatchSubmissionIndex != 0, "invalid next batch submission index");
        require(_oracle != address(0), "invalid oracle address");

        _transferOwnership(_owner);

        oracle = _oracle;
        nextBatchSubmissionIndex = _nextBatchSubmissionIndex;
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice set oracle address
    /// @param _oracle     oracle address
    function setOracleAddress(address _oracle) external onlyOwner {
        require(_oracle != address(0), "invalid oracle address");
        oracle = _oracle;
    }

    /// @notice set latest block
    /// @param _latestBlock   latest block
    function setLatestRewardEpochBlock(uint256 _latestBlock) external onlyOracle {
        require(latestRewardEpochBlock == 0, "already set");
        require(_latestBlock > 0, "invalid latest block");
        latestRewardEpochBlock = _latestBlock;
    }

    /// @notice record batch submissions
    function recordFinalizedBatchSubmissions(BatchSubmission[] calldata _batchSubmissions) external onlyOracle {
        require(_batchSubmissions.length > 0, "empty batch submissions");
        for (uint256 i = 0; i < _batchSubmissions.length; i++) {
            require(_batchSubmissions[i].index == nextBatchSubmissionIndex + i, "invalid index");
            // TODO: check more
            batchSubmissions[_batchSubmissions[i].index] = BatchSubmission(
                _batchSubmissions[i].index,
                _batchSubmissions[i].submitter,
                _batchSubmissions[i].startBlock,
                _batchSubmissions[i].endBlock,
                _batchSubmissions[i].rollupTime,
                _batchSubmissions[i].rollupBlock
            );
        }
        emit BatchSubmissionsUploaded(nextBatchSubmissionIndex, _batchSubmissions.length);
        nextBatchSubmissionIndex += _batchSubmissions.length;
    }

    /// @notice record epochs
    function recordRollupEpochs(RollupEpochInfo[] calldata _rollupEpochs) external onlyOracle {
        require(_rollupEpochs.length > 0, "empty rollup epochs");
        for (uint256 i = 0; i < _rollupEpochs.length; i++) {
            require(_rollupEpochs[i].index == nextRollupEpochIndex + i, "invalid index");
            // TODO: check more
            rollupEpochs[_rollupEpochs[i].index] = RollupEpochInfo(
                _rollupEpochs[i].index,
                _rollupEpochs[i].submitter,
                _rollupEpochs[i].startTime,
                _rollupEpochs[i].endTime,
                _rollupEpochs[i].endBlock
            );
        }
        emit RollupEpochsUploaded(nextRollupEpochIndex, _rollupEpochs.length);
        nextRollupEpochIndex += _rollupEpochs.length;
    }

    /// @notice record epochs
    function recordRewardEpochs(RewardEpochInfo[] calldata _rewardEpochs) external onlyOracle {
        require(_rewardEpochs.length > 0, "empty reward epochs");
        require(latestRewardEpochBlock > 0, "start block should be set");
        require(
            nextRewardEpochIndex + _rewardEpochs.length - 1 < IL2Staking(L2_STAKING_CONTRACT).currentEpoch(),
            "unfinished epochs cannot be uploaded"
        );
        IMorphToken(MORPH_TOKEN_CONTRACT).mintInflations(nextRewardEpochIndex + _rewardEpochs.length - 1);

        uint256 totalBlocks;
        for (uint256 i = 0; i < _rewardEpochs.length; i++) {
            uint256 dataLen = _rewardEpochs[i].sequencers.length;
            uint256 index = _rewardEpochs[i].index;
            require(index == nextRewardEpochIndex + i, "invalid epoch index");
            require(
                _rewardEpochs[i].sequencerBlocks.length == dataLen &&
                    _rewardEpochs[i].sequencerRatios.length == dataLen &&
                    _rewardEpochs[i].sequencerCommissions.length == dataLen,
                "invalid data length"
            );

            totalBlocks += _rewardEpochs[i].blockCount;
            rewardEpochs[index] = RewardEpochInfo(
                index,
                _rewardEpochs[i].blockCount,
                _rewardEpochs[i].sequencers,
                _rewardEpochs[i].sequencerBlocks,
                _rewardEpochs[i].sequencerRatios,
                _rewardEpochs[i].sequencerCommissions
            );

            uint256 inflationAmount = IMorphToken(MORPH_TOKEN_CONTRACT).inflation(index);
            uint256 blockCount;
            uint256 ratioSum;
            uint256[] memory delegatorRewards = new uint256[](dataLen);
            uint256[] memory commissions = new uint256[](dataLen);
            for (uint256 j = 0; j < dataLen; j++) {
                require(_rewardEpochs[i].sequencerCommissions[j] <= 20, "invalid sequencers commission");
                ratioSum += _rewardEpochs[i].sequencerRatios[j];
                blockCount += _rewardEpochs[i].sequencerBlocks[j];

                // compute rewards per sequencer
                uint256 reward = (inflationAmount * _rewardEpochs[i].sequencerRatios[j]) / PRECISION;
                commissions[j] = (reward * _rewardEpochs[i].sequencerCommissions[j]) / 100;
                delegatorRewards[j] = reward - commissions[j];
            }
            require(blockCount == _rewardEpochs[i].blockCount, "invalid sequencers blocks");
            require(ratioSum <= PRECISION, "invalid sequencers ratios");

            // update sequencers reward data
            IDistribute(DISTRIBUTE_CONTRACT).updateEpochReward(
                index,
                _rewardEpochs[i].sequencers,
                delegatorRewards,
                commissions
            );
        }

        emit RewardEpochsUploaded(nextRewardEpochIndex, _rewardEpochs.length);
        latestRewardEpochBlock += totalBlocks;
        nextRewardEpochIndex += _rewardEpochs.length;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice getBatchSubmissions
    /// @param start start index
    /// @param end   end index
    function getBatchSubmissions(uint256 start, uint256 end) external view returns (BatchSubmission[] memory res) {
        require(end >= start, "invalid index");
        res = new BatchSubmission[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = batchSubmissions[i];
        }
    }

    /// @notice get rollup epochs
    /// @param start start index
    /// @param end   end index
    function getRollupEpochs(uint256 start, uint256 end) external view returns (RollupEpochInfo[] memory res) {
        require(end >= start, "invalid index");
        res = new RollupEpochInfo[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = rollupEpochs[i];
        }
    }

    /// @notice get reward epochs
    /// @param start start index
    /// @param end   end index
    function getRewardEpochs(uint256 start, uint256 end) external view returns (RewardEpochInfo[] memory res) {
        require(end >= start, "invalid index");
        res = new RewardEpochInfo[](end - start + 1);
        for (uint256 i = start; i <= end; i++) {
            res[i] = rewardEpochs[i];
        }
    }
}
