// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2Staking} from "../l2/staking/L2Staking.sol";
import {IRecord} from "../l2/staking/IRecord.sol";
import {Types} from "../libraries/common/Types.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract L2StakingTest is L2StakingBaseTest {
    uint256 SEQUENCER_RATIO_PRECISION = 1e8;
    uint256 INFLATION_RATIO_PRECISION = 1e16;

    uint256 morphBalance = 20 ether;

    address[] stakers;

    address firstStaker;
    address secondStaker;
    address thirdStaker;

    function setUp() public virtual override {
        super.setUp();

        firstStaker = address(uint160(beginSeq));
        secondStaker = address(uint160(beginSeq + 1));
        thirdStaker = address(uint160(beginSeq + 2));

        hevm.startPrank(multisig);
        morphToken.transfer(bob, morphBalance);
        morphToken.transfer(alice, morphBalance);
        hevm.stopPrank();

        hevm.warp(rewardStartTime);
    }

    /**
     * @notice initialize: re-initialize
     */
    function test_initialize_paramsCheck_reverts() public {
        Types.StakerInfo[] memory _stakerInfos = new Types.StakerInfo[](0);

        hevm.expectRevert("Initializable: contract is already initialized");
        hevm.prank(multisig);
        l2Staking.initialize(0, 0, 0, _stakerInfos);

        // reset initialize
        hevm.store(
            address(l2Staking),
            bytes32(uint256(0)),
            bytes32(uint256(0))
        );

        hevm.expectRevert("sequencersSize must greater than 0");
        hevm.prank(multisig);
        l2Staking.initialize(0, 0, 0, _stakerInfos);

        hevm.expectRevert("invalid undelegateLockEpochs");
        hevm.prank(multisig);
        l2Staking.initialize(1, 0, 0, _stakerInfos);

        hevm.expectRevert("invalid reward start time");
        hevm.prank(multisig);
        l2Staking.initialize(1, 1, 100, _stakerInfos);

        hevm.expectRevert("invalid initial stakers");
        hevm.prank(multisig);
        l2Staking.initialize(1, 1, rewardStartTime * 2, _stakerInfos);
    }

    /**
     * @notice test init staker info & ranking
     */
    function test_init_stakersInfo_succeeds() public {
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            (address staker, , ) = l2Staking.stakers(user);
            assertEq(user, staker);

            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, i + 1);
        }
    }

    /**
     * @notice test add staker
     */
    function test_addStakers_succeeds() public {
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2 + 1; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }
        hevm.stopPrank();
        for (uint256 i = 0; i < SEQUENCER_SIZE * 2 + 1; i++) {
            address user = address(uint160(beginSeq + i));
            (address staker, , ) = l2Staking.stakers(user);
            assertEq(user, staker);
            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, i + 1);
        }

        assertEq(
            sequencer.getSequencerSet2Size(),
            l2Staking.sequencerSetMaxSize()
        );
    }

    /**
     * @notice test add staker, reward starting
     */
    function test_addStakerWhenRewardStarting_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }
        hevm.stopPrank();
        for (uint256 i = 0; i < SEQUENCER_SIZE * 2; i++) {
            address user = address(uint160(beginSeq + i));
            (address staker, , ) = l2Staking.stakers(user);
            assertEq(user, staker);
            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, i + 1);
        }

        // sequencer did not update
        // update by staking amount
        assertEq(sequencer.getSequencerSet2Size(), SEQUENCER_SIZE);
    }

    /**
     * @notice test removed staker
     */
    function test_removeStakers_succeeds() public {
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));

        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }

        address[] memory removed = new address[](2);
        removed[0] = address(uint160(beginSeq + 1));
        removed[1] = address(uint160(beginSeq + 4));

        l2Staking.removeStakers(removed);
        hevm.stopPrank();

        for (uint256 i = 0; i < 4; i++) {
            address user = l2Staking.stakerAddresses(i);
            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, i + 1);
        }

        for (uint256 i = 0; i < removed.length; i++) {
            address user = removed[i];
            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, 0);
        }
    }

    /**
     * @notice test set commission rate
     */
    function test_setCommissionRate_invalidCommission_reverts() public {
        hevm.startPrank(firstStaker);

        // set commission rate
        hevm.expectRevert("invalid commission");
        l2Staking.setCommissionRate(21);
    }

    /**
     * @notice failed delegate, staker not exists
     */
    function test_stake_notStaker_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        hevm.expectRevert("not staker");
        l2Staking.delegateStake(alice, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice staking by delegator
     * stag0
     */
    function test_stakeWhenRewardNotStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 amount = l2Staking.delegations(firstStaker, bob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice normal undelegate
     */
    function test_undelegate_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 stakerAmount0 = l2Staking.stakerDelegations(firstStaker);

        uint256 amount0 = l2Staking.delegations(firstStaker, bob);
        l2Staking.undelegateStake(firstStaker);
        uint256 amount1 = l2Staking.delegations(firstStaker, bob);
        assertEq(amount1, 0);

        uint256 stakerAmount1 = l2Staking.stakerDelegations(firstStaker);

        assertEq(stakerAmount1, stakerAmount0 - amount0);

        hevm.stopPrank();
    }

    /**
     * @notice undelegate, staker removed
     */
    function test_undelegateWhenStakerRemoved_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();

        // remove staker
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));
        address[] memory removed = new address[](1);
        removed[0] = firstStaker;
        l2Staking.removeStakers(removed);
        hevm.stopPrank();

        // sequenser size decrease
        assertEq(sequencer.getSequencerSet2Size(), SEQUENCER_SIZE - 1);
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 1);

        // staker ranking is 0, removed
        assertTrue(l2Staking.stakerRankings(firstStaker) == 0);

        hevm.startPrank(bob);
        l2Staking.undelegateStake(firstStaker);
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 1);
        hevm.stopPrank();

        hevm.startPrank(alice);
        l2Staking.undelegateStake(secondStaker);
        hevm.stopPrank();

        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 2);
    }

    /**
     * @notice failed unstaking, when staking amount is zero
     */
    function test_delegatorUnstaking_stakingAmountIsZero_reverts() public {
        hevm.startPrank(bob);

        hevm.expectRevert("staking amount is zero");
        l2Staking.undelegateStake(firstStaker);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, amount in lock period
     */
    function test_delegatorClaim_inLockPeriod_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.startPrank(bob);
        l2Staking.undelegateStake(firstStaker);
        hevm.expectRevert("no Morph token to claim");
        l2Staking.claimUndelegation();
        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegation
     */
    function test_delegatorClaimUndelegation_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.undelegateStake(firstStaker);

        uint256 time = rewardStartTime + REWARD_EPOCH * (ROLLUP_EPOCH + 1);

        hevm.warp(time);
        l2Staking.claimUndelegation();

        hevm.stopPrank();
    }

    /**
     * @notice failed restaking, pre claim in lock period
     */
    function test_delegatorRestake_inLockPeriod_fails() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.undelegateStake(firstStaker);
        hevm.expectRevert("undelegation unclaimed");
        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice normal restaking
     */
    function test_delegatorRestakeAfterLockPeriod_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.undelegateStake(firstStaker);

        hevm.roll(ROLLUP_EPOCH);

        uint256 time = rewardStartTime + REWARD_EPOCH * (ROLLUP_EPOCH + 1);

        hevm.warp(time);
        l2Staking.claimUndelegation();

        l2Staking.delegateStake(firstStaker, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice test ranking, reward_start = false
     */
    function test_rankWhenRewardNotStarting_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(secondStaker, morphBalance);
        hevm.stopPrank();

        uint256 secondRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(secondRanking, 1 + 1);

        uint256 firstRanking = l2Staking.stakerRankings(firstStaker);
        assertEq(firstRanking, 0 + 1);
    }

    /**
     * @notice test ranking, reward_start = true
     */
    function test_rankWhenRewardStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(secondStaker, 10 ether);
        l2Staking.delegateStake(firstStaker, 5 ether);
        hevm.stopPrank();

        hevm.prank(multisig);
        l2Staking.startReward();

        uint256 secondRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(secondRanking, 0 + 1);

        uint256 firstRanking = l2Staking.stakerRankings(firstStaker);
        assertEq(firstRanking, 1 + 1);

        address[] memory sequencerSet = sequencer.getSequencerSet2();
        assertEq(secondStaker, sequencerSet[0]);
        assertEq(firstStaker, sequencerSet[1]);
    }

    /**
     * @notice update sequencerSetMaxSize
     */
    function test_updateSequencerSetMaxSize_succeeds() public {
        hevm.prank(multisig);
        l2Staking.updateSequencerSetMaxSize(2);

        assertEq(sequencer.getSequencerSet2Size(), 2);
    }

    function _updateDistribute(uint256 epochIndex) internal returns (uint256) {
        uint256 sequencerSize = SEQUENCER_SIZE;
        uint256 blockCount = REWARD_EPOCH / 3; // 1 block per 3s
        address[] memory sequencers = sequencerAddresses;
        uint256[] memory sequencerBlocks = new uint256[](sequencerSize);
        uint256[] memory sequencerRatios = new uint256[](sequencerSize);
        uint256[] memory sequencerCommissions = new uint256[](sequencerSize);
        for (uint i = 0; i < sequencerSize; i++) {
            // same blocks
            sequencerBlocks[i] = blockCount / sequencerSize;
            sequencerRatios[i] = SEQUENCER_RATIO_PRECISION / sequencerSize;
            sequencerCommissions[i] = l2Staking.commissions(sequencers[i]);
        }

        IRecord.RewardEpochInfo[]
            memory rewardEpochInfos = new IRecord.RewardEpochInfo[](1);

        rewardEpochInfos[0] = IRecord.RewardEpochInfo(
            epochIndex,
            blockCount,
            sequencers,
            sequencerBlocks,
            sequencerRatios,
            sequencerCommissions
        );

        uint256 totalSupply = morphToken.totalSupply();
        hevm.startPrank(oracleAddress);
        record.recordRewardEpochs(rewardEpochInfos);
        hevm.stopPrank();

        uint256 totalInflations = (totalSupply * 1596535874529) /
            INFLATION_RATIO_PRECISION;
        uint256 inflationAmount = morphToken.inflation(epochIndex);
        assertEq(totalInflations, inflationAmount);

        return totalInflations;
    }

    /**
     * @notice  staking -> distribute -> claim
     */
    function test_delegatorClaimRewardWhenRewardStarting_succeeds() public {
        uint256 sequencerSize = sequencer.getSequencerSet2Size();

        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        // reward starting
        // rewardStartTime = 86400
        // block.timeStamp >= rewardStartTime
        // candidateNumber > 0
        hevm.prank(multisig);
        l2Staking.startReward();

        // staker set commission
        hevm.prank(firstStaker);
        l2Staking.setCommissionRate(1);
        hevm.prank(secondStaker);
        l2Staking.setCommissionRate(1);
        hevm.prank(thirdStaker);
        l2Staking.setCommissionRate(1);

        // *************** epoch = 1 ******************** //
        time = REWARD_EPOCH * 2;
        hevm.warp(time);

        uint256 blocksCountOfEpoch = REWARD_EPOCH / 3;
        hevm.roll(blocksCountOfEpoch * 2);
        hevm.prank(oracleAddress);
        record.setLatestRewardEpochBlock(blocksCountOfEpoch);
        _updateDistribute(0);

        // effectiveEpoch = 2
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(secondStaker, morphBalance - 5 ether);
        hevm.stopPrank();

        // ranking changed by delegate amount
        uint256 secondRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(secondRanking, 0 + 1);

        // *************** epoch = 2 ******************** //
        time = REWARD_EPOCH * 3;
        hevm.roll(blocksCountOfEpoch * 3);
        hevm.warp(time);
        _updateDistribute(1);

        // *************** epoch = 3 ******************** //
        time = REWARD_EPOCH * 4;
        hevm.roll(blocksCountOfEpoch * 4);
        hevm.warp(time);
        uint256 totalInflations = _updateDistribute(2);

        /**
         * 1. reward = 0 no remaining reward
         * 2. reward > 0
         */
        hevm.startPrank(bob);
        uint256 balanceBefore = morphToken.balanceOf(bob);
        l2Staking.claimReward(secondStaker, 2);
        uint256 balanceAfter = morphToken.balanceOf(bob);

        // sequncer size = 3
        // proposal same blocks in epoch 2
        // commission = 1
        // alice delegate 5 ether morph token
        // bob delefate 15 ether morph token
        // total delegate amount = (5 + 15) ether
        // check the reward

        uint256 commissionRate = l2Staking.commissions(secondStaker);
        uint256 sequencerEpochReward = ((totalInflations *
            (SEQUENCER_RATIO_PRECISION / sequencerSize)) /
            SEQUENCER_RATIO_PRECISION);
        uint256 commission = (sequencerEpochReward * commissionRate) / 100;
        uint256 delegatorReward = sequencerEpochReward - commission;

        uint256 bobReward = (delegatorReward * 15 ether) / (20 ether);

        assertEq(balanceAfter, balanceBefore + bobReward);
        hevm.stopPrank();
    }

    /**
     * @notice  staking -> distribute -> claim
     */
    function test_delegatorUndelegateWhenRewardStarting_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        // reward starting
        // rewardStartTime = 86400
        // block.timeStamp >= rewardStartTime
        // candidateNumber > 0
        hevm.prank(multisig);
        l2Staking.startReward();

        // staker set commission
        hevm.prank(firstStaker);
        l2Staking.setCommissionRate(1);
        hevm.prank(secondStaker);
        l2Staking.setCommissionRate(1);
        hevm.prank(thirdStaker);
        l2Staking.setCommissionRate(1);

        // *************** epoch = 1 ******************** //
        time = REWARD_EPOCH * 2;
        hevm.warp(time);

        uint256 blocksCountOfEpoch = REWARD_EPOCH / 3;
        hevm.roll(blocksCountOfEpoch * 2);
        hevm.prank(oracleAddress);
        record.setLatestRewardEpochBlock(blocksCountOfEpoch);
        uint256 totalInflations0 = _updateDistribute(0);

        // effectiveEpoch = 2
        hevm.startPrank(alice);
        l2Staking.undelegateStake(firstStaker);
        hevm.stopPrank();

        // candidateNumber decrease
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 1);
        // ranking changed by undelegate action
        assertEq(l2Staking.stakerRankings(secondStaker), 0 + 1);
        assertEq(l2Staking.stakerRankings(thirdStaker), 0 + 2);
        assertEq(l2Staking.stakerRankings(firstStaker), 0 + 3);

        hevm.startPrank(alice);
        hevm.expectRevert("no Morph token to claim");
        l2Staking.claimUndelegation();
        hevm.stopPrank();

        // *************** epoch = 2 ******************** //
        time = REWARD_EPOCH * 3 + 1;
        hevm.roll(blocksCountOfEpoch * 3);
        hevm.warp(time);
        uint256 totalInflations1 = _updateDistribute(1);

        // *************** epoch = 3 ******************** //
        time = REWARD_EPOCH * 4;
        hevm.roll(blocksCountOfEpoch * 4);
        hevm.warp(time);
        _updateDistribute(2);

        // *************** at unlock epoch ******************** //
        time = rewardStartTime + REWARD_EPOCH * (ROLLUP_EPOCH + 2);
        hevm.warp(time);
        hevm.prank(alice);
        l2Staking.claimUndelegation();

        /**
         * 1. reward = 0 no remaining reward
         * 2. reward > 0
         */
        hevm.startPrank(alice);
        uint256 balanceBefore = morphToken.balanceOf(alice);
        // total 20 ether
        // stake 5 ether to second staker
        // stake 5 ether to third staker
        // undelegate 5 ether
        assertEq(balanceBefore, 10 ether);
        l2Staking.claimReward(firstStaker, 0);
        uint256 balanceAfter = morphToken.balanceOf(alice);

        // sequncer size = 3
        // proposal same blocks in every epoch
        // commission = 1
        // alice delegate 5 ether morph token in epoch 0 - 1, undeletegate at epoch 1. valid reward epoch is 0, 1
        // check the reward
        uint256 validEpoch = 2;
        uint256[] memory rewardInflations = new uint256[](validEpoch);
        rewardInflations[0] = totalInflations0;
        rewardInflations[1] = totalInflations1;

        uint256 totalReward = 0;
        for (uint256 i = 0; i < validEpoch; i++) {
            uint256 commissionRate = l2Staking.commissions(secondStaker);
            uint256 sequencerEpochReward = ((rewardInflations[i] *
                (SEQUENCER_RATIO_PRECISION / SEQUENCER_SIZE)) /
                SEQUENCER_RATIO_PRECISION);
            uint256 commission = (sequencerEpochReward * commissionRate) / 100;
            uint256 delegatorReward = sequencerEpochReward - commission;

            totalReward += (delegatorReward * 5 ether) / (5 ether);
        }

        assertEq(balanceAfter, balanceBefore + totalReward);
        hevm.stopPrank();
    }

    /**
     * @notice currentEpoch
     */

    function test_currentEpoch_succeeds() public {
        uint256 currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 0);

        hevm.warp(rewardStartTime);
        currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 0);

        hevm.warp(rewardStartTime * 2);
        currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 1);
    }

    /**
     * @notice getStakesInfo
     */
    function test_getStakesInfo_succeeds() public {
        address[] memory _sequencerAddresses = new address[](SEQUENCER_SIZE);
        Types.StakerInfo[] memory stakerInfos0 = new Types.StakerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos0[i] = stakerInfo;
            _sequencerAddresses[i] = stakerInfo.addr;
        }

        Types.StakerInfo[] memory stakerInfos1 = l2Staking.getStakesInfo(
            _sequencerAddresses
        );

        // check params
        assertEq(stakerInfos1.length, stakerInfos0.length);
        for (uint i = 0; i < stakerInfos1.length; i++) {
            assertEq(stakerInfos0[i].addr, stakerInfos1[i].addr);
            assertEq(stakerInfos0[i].tmKey, stakerInfos1[i].tmKey);
        }
    }

    /**
     * @notice get stakers
     */
    function test_getStakers_succeeds() public {
        Types.StakerInfo[] memory stakerInfos0 = new Types.StakerInfo[](
            SEQUENCER_SIZE
        );

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos0[i] = stakerInfo;
        }

        Types.StakerInfo[] memory stakerInfos1 = l2Staking.getStakers();

        // check params
        assertEq(stakerInfos1.length, stakerInfos0.length);
        for (uint i = 0; i < stakerInfos1.length; i++) {
            assertEq(stakerInfos0[i].addr, stakerInfos1[i].addr);
            assertEq(stakerInfos0[i].tmKey, stakerInfos1[i].tmKey);
        }
    }
}
