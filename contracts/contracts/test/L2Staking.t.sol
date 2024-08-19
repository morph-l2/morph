// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {IRecord} from "../l2/staking/IRecord.sol";
import {Types} from "../libraries/common/Types.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {L2Staking} from "../l2/staking/L2Staking.sol";
import {IL2Staking} from "../l2/staking/IL2Staking.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {IDistribute} from "../l2/staking/IDistribute.sol";

// import "forge-std/console.sol";

contract L2StakingTest is L2StakingBaseTest {
    uint256 public morphBalance = 20 ether;
    address[] public stakers;
    address public firstStaker;
    address public secondStaker;
    address public thirdStaker;

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
        l2Staking.initialize(multisig, 0, 0, 0, _stakerInfos);

        // reset initialize
        hevm.store(address(l2Staking), bytes32(uint256(0)), bytes32(uint256(0)));

        hevm.expectRevert("sequencersSize must greater than 0");
        hevm.prank(multisig);
        l2Staking.initialize(multisig, 0, 0, 0, _stakerInfos);

        hevm.expectRevert("invalid undelegateLockEpochs");
        hevm.prank(multisig);
        l2Staking.initialize(multisig, 1, 0, 0, _stakerInfos);

        hevm.expectRevert("invalid reward start time");
        hevm.prank(multisig);
        l2Staking.initialize(multisig, 1, 1, 100, _stakerInfos);

        hevm.expectRevert("invalid initial stakers");
        hevm.prank(multisig);
        l2Staking.initialize(multisig, 1, 1, rewardStartTime * 2, _stakerInfos);
    }

    /**
     * @notice initialize: Test that the initialization of the L2Staking contract succeeds and emits the correct events.
     */
    function test_initialize_succeeds() public {
        // Deploy a TransparentUpgradeableProxy contract for l2StakingProxyTemp.
        TransparentUpgradeableProxy l2StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy L2Staking implementation.
        L2Staking l2StakingImplTemp = new L2Staking(payable(NON_ZERO_ADDRESS));

        Types.StakerInfo[] memory stakerInfos = new Types.StakerInfo[](SEQUENCER_SIZE);
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos[i] = stakerInfo;
            sequencerAddresses.push(stakerInfo.addr);
        }

        // Set the blockchain timestamp to half of the REWARD_EPOCH duration.
        hevm.warp(REWARD_EPOCH / 2);

        // Expect the SequencerSetMaxSizeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.SequencerSetMaxSizeUpdated(0, SEQUENCER_SIZE * 2);

        // Expect the RewardStartTimeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.RewardStartTimeUpdated(0, rewardStartTime);

        hevm.startPrank(multisig);

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2StakingProxyTemp)).upgradeToAndCall(
            address(l2StakingImplTemp),
            abi.encodeCall(
                L2Staking.initialize,
                (multisig, SEQUENCER_SIZE * 2, ROLLUP_EPOCH, rewardStartTime, stakerInfos)
            )
        );
        hevm.stopPrank();

        // Cast the proxy address to the L2Staking contract type to call its methods.
        L2Staking l2StakingTemp = L2Staking(payable(address(l2StakingProxyTemp)));

        // Verify the state varialbes are initialized succeefully.
        assertEq(l2StakingTemp.sequencerSetMaxSize(), SEQUENCER_SIZE * 2);
        assertEq(l2StakingTemp.undelegateLockEpochs(), ROLLUP_EPOCH);
        assertEq(l2StakingTemp.rewardStartTime(), rewardStartTime);
        assertEq(l2StakingTemp.latestSequencerSetSize(), stakerInfos.length);
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
     * @notice addStaker: Expect revert if not called by the other staking contract.
     */
    function test_addStaker_notOtherStaking_reverts() public {
        address staker = address(uint160(beginSeq));
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);

        // Expect revert due to unauthorized call.
        hevm.expectRevert("staking: only other staking contract allowed");
        l2Staking.addStaker(stakerInfo);
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
        assertEq(SEQUENCER_SIZE, l2Staking.getStakerAddressesLength());
        hevm.startPrank(address(l2CrossDomainMessenger));
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2 + 1; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);

            // Expect the SequencerSetMaxSizeUpdated event is emitted successfully.
            hevm.expectEmit(true, true, true, true);
            emit IL2Staking.StakerAdded(stakerInfo.addr, stakerInfo.tmKey, stakerInfo.blsKey);
            l2Staking.addStaker(stakerInfo);
        }
        assertEq(7, l2Staking.getStakerAddressesLength());
        hevm.stopPrank();
        for (uint256 i = 0; i < SEQUENCER_SIZE * 2 + 1; i++) {
            address user = address(uint160(beginSeq + i));
            (address staker, , ) = l2Staking.stakers(user);
            assertEq(user, staker);
            uint256 ranking = l2Staking.stakerRankings(user);
            assertEq(ranking, i + 1);
        }

        assertEq(sequencer.getSequencerSet2Size(), l2Staking.sequencerSetMaxSize());
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
     * @notice addStaker: Test that adding the same staker twice does not increase the staker addresses length.
     */
    function test_addStaker_addSameStaker_doesNotIncreaseLength_succeeds() public {
        // Mock the call to the L2Staking's messenger to simulate the xDomainMessageSender function call.
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );

        address staker = address(uint160(beginSeq));

        // Generate staker information using the ffi library.
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
        hevm.startPrank(address(l2CrossDomainMessenger));

        // Add the staker to the L2Staking contract.
        l2Staking.addStaker(stakerInfo);
        uint256 initialLength = l2Staking.getStakerAddressesLength();

        // Add the same staker again to the L2Staking contract.
        l2Staking.addStaker(stakerInfo);
        uint256 finalLength = l2Staking.getStakerAddressesLength();

        // Assert that the initial length and the final length are equal.
        assertEq(initialLength, finalLength);
        hevm.stopPrank();
    }

    /**
     * @notice removeStaker: Expect revert if not called by the other staking contract.
     */
    function test_removeStaker_notOtherStaking_reverts() public {
        address[] memory removed = new address[](1);
        removed[0] = address(uint160(beginSeq));

        // Expect revert due to unauthorized call.
        hevm.expectRevert("staking: only other staking contract allowed");
        l2Staking.removeStakers(removed);
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

        assertEq(sequencer.getSequencerSet2Size(), 4);
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
     * @notice removeStakers: Test that removing non-existent stakers does not change the staker addresses length.
     */
    function test_removeStakers_notExist_doesNotChangeLength_succeeds() public {
        // Mock the call to the L2Staking's messenger to simulate the xDomainMessageSender function call.
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));

        // Add a set of new stakers to the L2Staking contract.
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }

        // Create an array of addresses that do not exist in the staker set.
        address[] memory removed = new address[](2);
        removed[0] = address(uint160(1));
        removed[1] = address(uint160(2));

        uint256 initialLength = l2Staking.getStakerAddressesLength();
        l2Staking.removeStakers(removed);
        uint256 finalLength = l2Staking.getStakerAddressesLength();

        // Assert that the initial length and the final length are equal.
        assertEq(initialLength, finalLength);

        hevm.stopPrank();
    }

    /**
     * @notice removeStakers: Test that repeatedly removing the same stakers does not change the staker addresses length.
     */
    function test_removeStakers_repeated_doesNotChangeLength_succeeds() public {
        // Mock the call to the L2Staking's messenger to simulate the xDomainMessageSender function call.
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.startPrank(address(l2CrossDomainMessenger));

        // Add a set of new stakers to the L2Staking contract.
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(stakerInfo);
        }

        // Create an array of addresses to be removed.
        address[] memory removed = new address[](2);
        removed[0] = address(uint160(beginSeq + 1));
        removed[1] = address(uint160(beginSeq + 4));

        l2Staking.removeStakers(removed);
        uint256 initialLength = l2Staking.getStakerAddressesLength();
        l2Staking.removeStakers(removed);
        uint256 finalLength = l2Staking.getStakerAddressesLength();

        // Assert that the initial length and the final length are equal.
        assertEq(initialLength, finalLength);
        hevm.stopPrank();
    }

    /**
     * @notice setCommissionRate: Expect revert if not called by a staker.
     */
    function test_setCommissionRate_notStaker_reverts() public {
        // Expect revert due to only staker allowed.
        hevm.expectRevert("only staker allowed");
        hevm.startPrank(address(1));
        l2Staking.setCommissionRate(21);
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
     * @notice setCommissionRate: Test setting the commission rate before rewards start.
     */
    function test_setCommissionRate_rewardNotStarted_succeeds() public {
        // Expect the CommissionUpdated event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.CommissionUpdated(firstStaker, 18, 0);

        hevm.prank(firstStaker);
        l2Staking.setCommissionRate(18);
    }

    /**
     * @notice setCommissionRate: Test setting the commission rate after rewards start.
     */
    function test_setCommissionRate_rewardStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.prank(multisig);
        l2Staking.startReward();

        // Expect the CommissionUpdated event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.CommissionUpdated(firstStaker, 18, l2Staking.currentEpoch() + 1);

        hevm.prank(firstStaker);
        l2Staking.setCommissionRate(18);
    }

    /**
     * @notice failed delegate, staker not exists
     */
    function test_delegateStake_notStaker_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        hevm.expectRevert("not staker");
        l2Staking.delegateStake(alice, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice delegateStake: Expect revert with invalid stake amount.
     */
    function test_delegateStake_invalidStakeAmount_reverts() public {
        hevm.expectRevert("invalid stake amount");
        l2Staking.delegateStake(firstStaker, 0);
    }

    /**
     * @notice staking by delegator
     * stag0
     */
    function test_delegateStake_stakeWhenRewardNotStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        // Expect the Delegated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, bob, morphBalance, morphBalance, 0);

        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 amount = l2Staking.delegations(firstStaker, bob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice delegateStake: Test that delegating the same stake again succeeds.
     */
    function test_delegateStake_delegateTheSameStakeAgain_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        // Delegate half of Bob's stake to firstStaker.
        l2Staking.delegateStake(firstStaker, morphBalance / 2);

        // Verify that the delegation amount is correct.
        uint256 amount = l2Staking.delegations(firstStaker, bob);
        assertEq(morphBalance / 2, amount);

        // Delegate the remaining half of Bob's stake to firstStaker.
        l2Staking.delegateStake(firstStaker, morphBalance / 2);

        // Verify that the total delegation amount is correct.
        amount = l2Staking.delegations(firstStaker, bob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice delegateStake: Test that the candidate number increases when a stake is delegated.
     */
    function test_delegateStake_candidateNumberIncreases_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        uint256 initialCandidateNumber = l2Staking.candidateNumber();

        // Delegate Bob's stake to firstStaker.
        l2Staking.delegateStake(firstStaker, morphBalance);

        // Verify that the candidate number has increased by 1.
        uint256 newCandidateNumber = l2Staking.candidateNumber();
        assertEq(newCandidateNumber, initialCandidateNumber + 1);
        hevm.stopPrank();
    }

    /**
     * @notice delegateStake: Test that delegating a stake succeeds after rewards have started.
     */
    function test_delegateStake_rewardStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(secondStaker, morphBalance / 2);
        uint256 oldRanking = l2Staking.stakerRankings(secondStaker);
        hevm.stopPrank();

        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.startPrank(bob);

        // Verify the Delegated event is emitted successfully.
        uint256 beforeAmount = l2Staking.delegations(secondStaker, bob);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(
            secondStaker,
            bob,
            beforeAmount + morphBalance / 2,
            beforeAmount,
            l2Staking.currentEpoch() + 1
        );

        l2Staking.delegateStake(secondStaker, morphBalance / 2);

        // Verify rankings is updated.
        uint256 newRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(newRanking, oldRanking - 1);

        hevm.stopPrank();
    }

    /**
     * @notice delegateStake: Test that notifyDelegation is called successfully when delegating a stake.
     */
    function test_delegateStake_notifyDelegationCalled_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        // Expect call method notifyDelegation.
        hevm.expectCall(
            address(l2Staking.DISTRIBUTE_CONTRACT()),
            abi.encodeWithSelector(
                IDistribute.notifyDelegation.selector,
                firstStaker,
                bob,
                0,
                morphBalance,
                morphBalance,
                1,
                true
            )
        );

        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice normal undelegate
     */
    function test_undelegateStake_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 stakerAmount0 = l2Staking.stakerDelegations(firstStaker);
        uint256 amount0 = l2Staking.delegations(firstStaker, bob);

        // Verify the Undelegated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Undelegated(firstStaker, bob, morphBalance, 0, 0);

        l2Staking.undelegateStake(firstStaker);
        uint256 amount1 = l2Staking.delegations(firstStaker, bob);
        assertEq(amount1, 0);

        // Verify the total stake amount of firstStaker is correct
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
    function test_undelegateStake_stakingAmountIsZero_reverts() public {
        hevm.startPrank(bob);

        hevm.expectRevert("staking amount is zero");
        l2Staking.undelegateStake(firstStaker);

        hevm.stopPrank();
    }

    /**
     * @notice undelegateStake: Updates rankings and candidate number when staker is not removed.
     */
    function test_undelegateStake_updatesRankingsAndCandidateNumber() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 1 ether);
        l2Staking.delegateStake(secondStaker, 1 ether);
        l2Staking.delegateStake(thirdStaker, 1 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.startPrank(bob);
        l2Staking.delegateStake(firstStaker, morphBalance / 2);
        hevm.stopPrank();

        // Check rankings and candidate number before undelegating.
        uint256 beforeRanking = l2Staking.stakerRankings(firstStaker);
        uint256 beforeCandidateNumber = l2Staking.candidateNumber();

        hevm.startPrank(bob);
        l2Staking.undelegateStake(firstStaker);
        hevm.stopPrank();

        // Check rankings and candidate number after undelegating.
        uint256 afterRanking = l2Staking.stakerRankings(firstStaker);
        uint256 afterCandidateNumber = l2Staking.candidateNumber();

        assertTrue(afterRanking > beforeRanking);
        assertEq(afterCandidateNumber, beforeCandidateNumber - 1);
    }

    /**
     * @notice undelegateStake: Test that notifyUndelegation is called successfully when undelegating a stake.
     */
    function test_undelegateStake_notifyUndelegationCalled_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);

        // Expect call method notifyUndelegation.
        hevm.expectCall(
            address(l2Staking.DISTRIBUTE_CONTRACT()),
            abi.encodeWithSelector(IDistribute.notifyUndelegation.selector, firstStaker, bob, 0, 0, 0)
        );

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
    function test_claimUndelegation_delegatorClaimUndelegation_succeeds() public {
        hevm.startPrank(bob);
        uint256 currentBalance = morphToken.balanceOf(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.undelegateStake(firstStaker);

        uint256 time = rewardStartTime + REWARD_EPOCH * 1;
        hevm.warp(time);

        // Expect the UndelegationClaimed event is emitted successfully.
        (address delegatee, uint256 amount, uint256 unlockEpoch) = l2Staking.undelegations(bob, 0);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.UndelegationClaimed(delegatee, bob, unlockEpoch, amount);

        l2Staking.claimUndelegation();
        uint256 newBalance = morphToken.balanceOf(bob);

        // Verify the balance of bob is correct.
        assertEq(currentBalance, newBalance);

        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegation
     */
    function test_claimUndelegation_delegatorClaimUndelegations_succeeds() public {
        hevm.startPrank(bob);
        uint256 currentBalance = morphToken.balanceOf(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance / 2);
        l2Staking.delegateStake(secondStaker, morphBalance / 2);

        hevm.warp(rewardStartTime + REWARD_EPOCH * 2);

        l2Staking.undelegateStake(firstStaker);
        l2Staking.undelegateStake(secondStaker);

        hevm.warp(rewardStartTime + REWARD_EPOCH * 4);

        // Expect the UndelegationClaimed event is emitted successfully.
        uint256 amount;
        uint256 unlockEpoch;
        hevm.expectEmit(true, true, true, true);
        (, amount, unlockEpoch) = l2Staking.undelegations(bob, 0);
        emit IL2Staking.UndelegationClaimed(firstStaker, bob, unlockEpoch, amount);
        hevm.expectEmit(true, true, true, true);
        (, amount, unlockEpoch) = l2Staking.undelegations(bob, 1);
        emit IL2Staking.UndelegationClaimed(secondStaker, bob, unlockEpoch, amount);

        l2Staking.claimUndelegation();
        uint256 newBalance = morphToken.balanceOf(bob);

        // Verify the balance of bob is correct.
        assertEq(currentBalance, newBalance);

        hevm.stopPrank();
    }

    /**
     * @notice failed restaking, pre claim in lock period
     */
    function test_delegateStake_delegatorRestake_inLockPeriod_reverts() public {
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
    function test_claimUndelegation_delegatorRestakeAfterLockPeriod_succeeds() public {
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
     * @notice updateSequencerSetMaxSize: Reverts if called by a non-owner.
     */
    function test_updateSequencerSetMaxSize_notOwner_reverts() public {
        hevm.prank(address(1));

        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.updateSequencerSetMaxSize(2);
    }

    /**
     * @notice updateSequencerSetMaxSize: Reverts if set to zero.
     */
    function test_updateSequencerSetMaxSize_eqZero_reverts() public {
        hevm.prank(multisig);

        // Expect revert due to _sequencerSetMaxSize equals zero.
        hevm.expectRevert("invalid new sequencer set max size");
        l2Staking.updateSequencerSetMaxSize(0);
    }

    /**
     * @notice updateSequencerSetMaxSize: Reverts if set to current max size.
     */
    function test_updateSequencerSetMaxSize_eqCurrentMaxSize_reverts() public {
        uint256 oldSize = l2Staking.sequencerSetMaxSize();

        // Expect revert due to _sequencerSetMaxSize equals currect sequencerSetMaxSize.
        hevm.expectRevert("invalid new sequencer set max size");
        hevm.prank(multisig);
        l2Staking.updateSequencerSetMaxSize(oldSize);
    }

    /**
     * @notice update sequencerSetMaxSize
     */
    function test_updateSequencerSetMaxSize_succeeds() public {
        // Expect the SequencerSetMaxSizeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.SequencerSetMaxSizeUpdated(l2Staking.sequencerSetMaxSize(), 2);

        hevm.prank(multisig);
        l2Staking.updateSequencerSetMaxSize(2);

        assertEq(sequencer.getSequencerSet2Size(), 2);
    }

    /**
     * @notice updateSequencerSetMaxSize: No _updateSequencerSet call when increasing size.
     */
    function test_updateSequencerSetMaxSize_largerValue_noUpdateSequencerSet_succeeds() public {
        uint256 oldSequencerSet2Size = sequencer.getSequencerSet2Size();
        uint256 oldSize = l2Staking.sequencerSetMaxSize();
        uint256 newSize = oldSize + 1;

        hevm.prank(multisig);
        l2Staking.updateSequencerSetMaxSize(newSize);

        // Verify that _updateSequencerSet was not called.
        assertEq(sequencer.getSequencerSet2Size(), oldSequencerSet2Size);
    }

    /**
     * @notice updateRewardStartTime: Reverts if called by non-owner.
     */
    function test_updateRewardStartTime_onlyOwner_reverts() public {
        hevm.prank(address(1));

        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.updateRewardStartTime(block.timestamp + REWARD_EPOCH);
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward already started.
     */
    function test_updateRewardStartTime_rewardAlreadyStarted_reverts() public {
        // Expect revert due to reward already started.
        hevm.expectRevert("reward already started");
        hevm.prank(multisig);
        l2Staking.updateRewardStartTime(block.timestamp + REWARD_EPOCH * 2);
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward start time is before current block time.
     */
    function test_updateRewardStartTime_rewardStartTimeBeforeBlockTime_reverts() public {
        // Expect revert due to rewardStartTime being before block.timestamp.
        hevm.expectRevert("reward already started");
        hevm.prank(multisig);
        l2Staking.updateRewardStartTime(block.timestamp - REWARD_EPOCH);
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward already started.
     */
    function test_updateRewardStartTime_rewardStartTimeEqBlockTime_reverts() public {
        // Expect revert due to rewardStartTime equals to block.timestamp.
        hevm.expectRevert("reward already started");
        hevm.prank(multisig);
        l2Staking.updateRewardStartTime(block.timestamp);
    }

    /**
     * @notice updateRewardStartTime: Reverts if not a multiple of REWARD_EPOCH.
     */
    function test_updateRewardStartTime_notMultipleRewardEpoch_reverts() public {
        hevm.warp(REWARD_EPOCH / 2);
        hevm.prank(multisig);

        // Expect revert due to updateRewardStartTime not REWARD_EPOCH multiple.
        hevm.expectRevert("invalid reward start time");
        l2Staking.updateRewardStartTime(block.timestamp + REWARD_EPOCH / 2);
    }

    /**
     * @notice updateRewardStartTime: Updates the reward start time successfully.
     */
    function test_updateRewardStartTime_succeeds() public {
        uint256 oldTime = l2Staking.rewardStartTime();
        uint256 newTime = block.timestamp + REWARD_EPOCH * 2;
        hevm.warp(REWARD_EPOCH / 2);

        // Verify the SequencerSetMaxSizeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.RewardStartTimeUpdated(oldTime, newTime);

        hevm.prank(multisig);
        l2Staking.updateRewardStartTime(newTime);

        assertEq(l2Staking.rewardStartTime(), newTime);
    }

    /**
     * @notice startReward: Reverts if called by a non-owner.
     */
    function test_startReward_notOwner_reverts() public {
        hevm.prank(address(1));

        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.startReward();
    }

    /**
     * @notice startReward: Reverts if called before reward start time.
     */
    function test_startReward_beforeRewardStartTime_reverts() public {
        hevm.warp(REWARD_EPOCH / 2);
        hevm.startPrank(multisig);

        uint256 newRewardStartTime = REWARD_EPOCH + 4800;
        newRewardStartTime = ((newRewardStartTime / REWARD_EPOCH) + 1) * REWARD_EPOCH;

        // Update the reward start time.
        l2Staking.updateRewardStartTime(newRewardStartTime);

        // Expect revert due to "can't start before reward start time".
        hevm.expectRevert("can't start before reward start time");
        l2Staking.startReward();

        hevm.stopPrank();
    }

    /**
     * @notice startReward: Reverts if no candidates.
     */
    function test_startReward_notCandidateNumber_reverts() public {
        hevm.prank(multisig);

        // Expect revert due to none candidate.
        hevm.expectRevert("none candidate");
        l2Staking.startReward();
    }

    function _updateDistribute(uint256 epochIndex) internal returns (uint256) {
        uint256 sequencerSize = SEQUENCER_SIZE;
        uint256 blockCount = REWARD_EPOCH / 3; // 1 block per 3s
        address[] memory sequencers = sequencerAddresses;
        uint256[] memory sequencerBlocks = new uint256[](sequencerSize);
        uint256[] memory sequencerRatios = new uint256[](sequencerSize);
        uint256[] memory sequencerCommissions = new uint256[](sequencerSize);
        for (uint256 i = 0; i < sequencerSize; i++) {
            // same blocks
            sequencerBlocks[i] = blockCount / sequencerSize;
            sequencerRatios[i] = SEQUENCER_RATIO_PRECISION / sequencerSize;
            sequencerCommissions[i] = l2Staking.commissions(sequencers[i]);
        }

        IRecord.RewardEpochInfo[] memory rewardEpochInfos = new IRecord.RewardEpochInfo[](1);

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

        uint256 totalInflations = (totalSupply * 1596535874529) / INFLATION_RATIO_PRECISION;
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
        uint256 sequencerEpochReward = ((totalInflations * (SEQUENCER_RATIO_PRECISION / sequencerSize)) /
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
    function test_delegatorClaimAllRewardWhenRewardStarting_succeeds() public {
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
        _updateDistribute(2);

        (address[] memory delegetees, uint256[] memory aliceRewards) = distribute.queryAllUnclaimed(alice);
        uint256 aliceReward1 = distribute.queryUnclaimed(firstStaker, alice);
        uint256 aliceReward2 = distribute.queryUnclaimed(secondStaker, alice);
        uint256 aliceReward3 = distribute.queryUnclaimed(thirdStaker, alice);
        assertEq(delegetees[0], firstStaker);
        assertEq(delegetees[1], secondStaker);
        assertEq(delegetees[2], thirdStaker);
        assertEq(aliceRewards[0], aliceReward1);
        assertEq(aliceRewards[1], aliceReward2);
        assertEq(aliceRewards[2], aliceReward3);

        // console.logString("......................");
        // console.logUint(aliceReward1);
        // console.logUint(aliceReward2);
        // console.logUint(aliceReward3);
        // console.logString("......................");

        // *************** epoch = 4 ******************** //
        time = REWARD_EPOCH * 5;
        hevm.roll(blocksCountOfEpoch * 5);
        hevm.warp(time);
        _updateDistribute(3);

        aliceReward1 = distribute.queryUnclaimed(firstStaker, alice);
        aliceReward2 = distribute.queryUnclaimed(secondStaker, alice);
        aliceReward3 = distribute.queryUnclaimed(thirdStaker, alice);

        // console.logString("......................");
        // console.logUint(aliceReward1);
        // console.logUint(aliceReward2);
        // console.logUint(aliceReward3);
        // console.logString("......................");

        hevm.startPrank(alice);
        uint256 balanceBefore = morphToken.balanceOf(alice);
        l2Staking.claimReward(address(0), 0);
        uint256 balanceAfter = morphToken.balanceOf(alice);

        // console.logString("......................");
        // console.logUint(balanceBefore);
        // console.logUint(balanceAfter);
        // console.logString("......................");

        assertEq(balanceAfter, balanceBefore + aliceReward1 + aliceReward2 + aliceReward3);
        hevm.stopPrank();
    }

    /**
     * @notice  staking -> distribute -> claim
     */
    function test_delegatorClaimAllRewardAfterUndelegate_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(bob);
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

        // *************** epoch = 2 ******************** //
        time = REWARD_EPOCH * 3;
        hevm.roll(blocksCountOfEpoch * 3);
        hevm.warp(time);
        _updateDistribute(1);

        // *************** epoch = 3 ******************** //
        time = REWARD_EPOCH * 4;
        hevm.roll(blocksCountOfEpoch * 4);
        hevm.warp(time);
        _updateDistribute(2);

        (address[] memory delegetees, uint256[] memory aliceRewards) = distribute.queryAllUnclaimed(alice);
        uint256 aliceReward1 = distribute.queryUnclaimed(firstStaker, alice);
        uint256 aliceReward2 = distribute.queryUnclaimed(secondStaker, alice);
        assertEq(delegetees[0], firstStaker);
        assertEq(delegetees[1], secondStaker);
        assertEq(aliceRewards[0], aliceReward1);
        assertEq(aliceRewards[1], aliceReward2);

        // console.logString("......................");
        // console.logUint(aliceReward1);
        // console.logUint(aliceReward2);
        // console.logString("......................");

        // *************** epoch = 4 ******************** //
        time = REWARD_EPOCH * 5;
        hevm.roll(blocksCountOfEpoch * 5);
        hevm.warp(time);
        _updateDistribute(3);

        aliceReward1 = distribute.queryUnclaimed(firstStaker, alice);
        aliceReward2 = distribute.queryUnclaimed(secondStaker, alice);

        hevm.startPrank(alice);
        l2Staking.undelegateStake(firstStaker);
        l2Staking.undelegateStake(secondStaker);
        IL2Staking.Undelegation[] memory undelegations = l2Staking.getUndelegations(alice);
        assertEq(undelegations.length, 2);

        // console.logString("......................");
        // console.logUint(aliceReward1);
        // console.logUint(aliceReward2);
        // console.logString("......................");

        // *************** epoch = 5 ******************** //
        time = REWARD_EPOCH * 6;
        hevm.roll(blocksCountOfEpoch * 6);
        hevm.warp(time);
        _updateDistribute(4);

        aliceReward1 = distribute.queryUnclaimed(firstStaker, alice);
        aliceReward2 = distribute.queryUnclaimed(secondStaker, alice);

        // console.logString("......................");
        // console.logUint(aliceReward1);
        // console.logUint(aliceReward2);
        // console.logString("......................");

        hevm.startPrank(alice);
        uint256 balanceBefore = morphToken.balanceOf(alice);
        l2Staking.claimReward(firstStaker, 0);
        l2Staking.claimReward(secondStaker, 0);
        uint256 balanceAfter = morphToken.balanceOf(alice);

        // console.logString("......................");
        // console.logUint(balanceBefore);
        // console.logUint(balanceAfter);
        // console.logString("......................");

        assertEq(balanceAfter, balanceBefore + aliceReward1 + aliceReward2);
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
            uint256 sequencerEpochReward = ((rewardInflations[i] * (SEQUENCER_RATIO_PRECISION / SEQUENCER_SIZE)) /
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
        Types.StakerInfo[] memory stakerInfos0 = new Types.StakerInfo[](SEQUENCER_SIZE);

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos0[i] = stakerInfo;
            _sequencerAddresses[i] = stakerInfo.addr;
        }

        Types.StakerInfo[] memory stakerInfos1 = l2Staking.getStakesInfo(_sequencerAddresses);

        // check params
        assertEq(stakerInfos1.length, stakerInfos0.length);
        for (uint256 i = 0; i < stakerInfos1.length; i++) {
            assertEq(stakerInfos0[i].addr, stakerInfos1[i].addr);
            assertEq(stakerInfos0[i].tmKey, stakerInfos1[i].tmKey);
        }
    }

    /**
     * @notice get stakers
     */
    function test_getStakers_succeeds() public {
        Types.StakerInfo[] memory stakerInfos0 = new Types.StakerInfo[](SEQUENCER_SIZE);

        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            stakerInfos0[i] = stakerInfo;
        }

        Types.StakerInfo[] memory stakerInfos1 = l2Staking.getStakers();

        // check params
        assertEq(stakerInfos1.length, stakerInfos0.length);
        for (uint256 i = 0; i < stakerInfos1.length; i++) {
            assertEq(stakerInfos0[i].addr, stakerInfos1[i].addr);
            assertEq(stakerInfos0[i].tmKey, stakerInfos1[i].tmKey);
        }
    }

    /**
     * @notice getDelegators
     */
    function test_getDelegators_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        address[] memory delegator0 = l2Staking.getAllDelegators(firstStaker);
        address[] memory delegator1 = l2Staking.getAllDelegators(secondStaker);
        address[] memory delegator2 = l2Staking.getAllDelegators(thirdStaker);

        assertEq(delegator0.length, 1);
        assertEq(delegator1.length, 1);
        assertEq(delegator2.length, 1);

        assertEq(delegator0[0], alice);
        assertEq(delegator1[0], alice);
        assertEq(delegator2[0], alice);
    }

    /**
     * @notice isStakingTo
     */
    function test_isStakingTo_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice);
        assertBoolEq(l2Staking.isStakingTo(firstStaker), true);
        assertBoolEq(l2Staking.isStakingTo(secondStaker), true);
        assertBoolEq(l2Staking.isStakingTo(thirdStaker), true);
        hevm.stopPrank();
    }
}
