// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Types} from "../libraries/common/Types.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {L2Staking} from "../l2/staking/L2Staking.sol";
import {IL2Staking} from "../l2/staking/IL2Staking.sol";
import {IMorphToken} from "../l2/system/IMorphToken.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {console} from "forge-std/Test.sol";

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
        morphToken.transfer(alice, morphBalance);
        morphToken.transfer(alice1, morphBalance);
        morphToken.transfer(alice2, morphBalance);
        morphToken.transfer(alice3, morphBalance);
        morphToken.transfer(alice4, morphBalance);
        morphToken.transfer(alice5, morphBalance);

        morphToken.transfer(bob, morphBalance);
        morphToken.transfer(bob1, morphBalance);
        morphToken.transfer(bob2, morphBalance);
        morphToken.transfer(bob3, morphBalance);
        morphToken.transfer(bob4, morphBalance);
        morphToken.transfer(bob5, morphBalance);
        hevm.stopPrank();

        hevm.warp(rewardStartTime);
    }

    /**
     * @notice initialize: re-initialize
     */
    function test_initialize_paramsCheck_reverts() public {
        Types.StakerInfo[] memory _stakerInfos = new Types.StakerInfo[](0);

        hevm.startPrank(multisig);
        hevm.expectRevert("Initializable: contract is already initialized");
        l2Staking.initialize(multisig, 0, 0, 0, _stakerInfos);
        hevm.stopPrank();

        // reset initialize
        hevm.store(address(l2Staking), bytes32(uint256(0)), bytes32(uint256(0)));

        hevm.startPrank(multisig);
        hevm.expectRevert(IL2Staking.ErrZeroSequencerSize.selector);
        l2Staking.initialize(multisig, 0, 0, 0, _stakerInfos);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.expectRevert(IL2Staking.ErrZeroLockEpochs.selector);
        l2Staking.initialize(multisig, 1, 0, 0, _stakerInfos);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.expectRevert(IL2Staking.ErrInvalidStartTime.selector);
        l2Staking.initialize(multisig, 1, 1, 100, _stakerInfos);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.expectRevert(IL2Staking.ErrNoStakers.selector);
        l2Staking.initialize(multisig, 1, 1, rewardStartTime * 2, _stakerInfos);
        hevm.stopPrank();
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
                (multisig, SEQUENCER_SIZE * 2, UNDELEGATE_LOCKED_EPOCHS, rewardStartTime, stakerInfos)
            )
        );
        hevm.stopPrank();

        // Cast the proxy address to the L2Staking contract type to call its methods.
        L2Staking l2StakingTemp = L2Staking(payable(address(l2StakingProxyTemp)));

        // Verify the state varialbes are initialized succeefully.
        assertEq(l2StakingTemp.sequencerSetMaxSize(), SEQUENCER_SIZE * 2);
        assertEq(l2StakingTemp.undelegateLockEpochs(), UNDELEGATE_LOCKED_EPOCHS);
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
        l2Staking.addStaker(0, stakerInfo);
    }

    /**
     * @notice test add staker: Reverts if invalid nonce
     */
    function test_addStakers_invalidNonce_reverts() public {
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        assertEq(SEQUENCER_SIZE, l2Staking.getStakerAddressesLength());

        hevm.startPrank(address(l2CrossDomainMessenger));
        address staker = address(uint160(beginSeq));
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
        uint256 nonce = l2Staking.nonce();
        hevm.expectRevert(IL2Staking.ErrInvalidNonce.selector);
        l2Staking.addStaker(nonce + 1, stakerInfo);
        hevm.stopPrank();
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

        uint256 nonce = 0;
        hevm.startPrank(address(l2CrossDomainMessenger));
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2 + 1; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);

            // Expect the SequencerSetMaxSizeUpdated event is emitted successfully.
            hevm.expectEmit(true, true, true, true);
            emit IL2Staking.StakerAdded(stakerInfo.addr, stakerInfo.tmKey, stakerInfo.blsKey);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
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
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        l2Staking.delegate(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );
        hevm.stopPrank();

        hevm.startPrank(address(l2CrossDomainMessenger));
        uint256 nonce = 0;
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
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
        l2Staking.addStaker(0, stakerInfo);
        uint256 initialLength = l2Staking.getStakerAddressesLength();

        // Add the same staker again to the L2Staking contract.
        l2Staking.addStaker(1, stakerInfo);
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
        l2Staking.removeStakers(0, removed);
    }

    /**
     * @notice test removed staker: Reverts if invalid nonce
     */
    function test_removeStakers_invalidNonce_reverts() public {
        hevm.mockCall(
            address(l2Staking.MESSENGER()),
            abi.encodeCall(ICrossDomainMessenger.xDomainMessageSender, ()),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );

        hevm.startPrank(address(l2CrossDomainMessenger));
        uint256 nonce = 0;
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
        }
        address[] memory removed = new address[](2);
        removed[0] = address(uint160(beginSeq + 1));
        removed[1] = address(uint160(beginSeq + 4));
        hevm.expectRevert(IL2Staking.ErrInvalidNonce.selector);
        l2Staking.removeStakers(nonce + 1, removed);
        hevm.expectRevert(IL2Staking.ErrInvalidNonce.selector);
        l2Staking.removeStakers(nonce - 1, removed);
        hevm.stopPrank();
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
        uint256 nonce = 0;
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
        }

        address[] memory removed = new address[](2);
        removed[0] = address(uint160(beginSeq + 1));
        removed[1] = address(uint160(beginSeq + 4));

        l2Staking.removeStakers(nonce, removed);

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
        uint256 nonce = 0;
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
        }

        // Create an array of addresses that do not exist in the staker set.
        address[] memory removed = new address[](2);
        removed[0] = address(uint160(1));
        removed[1] = address(uint160(2));

        uint256 initialLength = l2Staking.getStakerAddressesLength();
        l2Staking.removeStakers(nonce, removed);
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
        uint256 nonce = 0;
        for (uint256 i = SEQUENCER_SIZE; i < SEQUENCER_SIZE * 2; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(staker);
            l2Staking.addStaker(nonce, stakerInfo);
            nonce++;
        }

        // Create an array of addresses to be removed.
        address[] memory removed = new address[](2);
        removed[0] = address(uint160(beginSeq + 1));
        removed[1] = address(uint160(beginSeq + 4));

        l2Staking.removeStakers(nonce, removed);
        nonce++;
        uint256 initialLength = l2Staking.getStakerAddressesLength();
        l2Staking.removeStakers(nonce, removed);
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
        hevm.startPrank(address(1));
        hevm.expectRevert(IL2Staking.ErrNotStaker.selector);
        l2Staking.setCommissionRate(21);
        hevm.stopPrank();
    }

    /**
     * @notice test set commission rate
     */
    function test_setCommissionRate_invalidCommission_reverts() public {
        hevm.startPrank(firstStaker);
        hevm.expectRevert(IL2Staking.ErrInvalidCommissionRate.selector);
        l2Staking.setCommissionRate(21);
        hevm.stopPrank();
    }

    /**
     * @notice setCommissionRate: Test setting the commission rate before rewards start.
     */
    function test_setCommissionRate_rewardNotStarted_succeeds() public {
        // Expect the CommissionUpdated event to be emitted successfully.
        hevm.startPrank(firstStaker);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.CommissionUpdated(firstStaker, 18, 0);
        l2Staking.setCommissionRate(18);
        hevm.stopPrank();
    }

    /**
     * @notice setCommissionRate: Test setting the commission rate after rewards start.
     */
    function test_setCommissionRate_rewardStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(firstStaker);
        // Expect the CommissionUpdated event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.CommissionUpdated(firstStaker, 18, 0);
        l2Staking.setCommissionRate(18);
        hevm.stopPrank();
    }

    /**
     * @notice failed delegate, staker not exists
     */
    function test_delegate_notStaker_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        hevm.expectRevert(IL2Staking.ErrNotStaker.selector);
        l2Staking.delegate(alice, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice delegate: Expect revert with invalid stake amount.
     */
    function test_delegate_invalidStakeAmount_reverts() public {
        hevm.expectRevert(IL2Staking.ErrZeroAmount.selector);
        l2Staking.delegate(firstStaker, 0);
    }

    /**
     * @notice staking by delegator
     * stag0
     */
    function test_delegate_stakeWhenRewardNotStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        // Expect the Delegated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, bob, morphBalance, morphBalance);
        l2Staking.delegate(firstStaker, morphBalance);
        uint256 amount = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(morphBalance, amount);
        assertEq(l2Staking.candidateNumber(), 1);
        hevm.stopPrank();
    }

    /**
     * @notice delegate: Test that delegating the same stake again succeeds.
     */
    function test_delegate_delegateTheSameStakeAgain_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        // Delegate half of Bob's stake to firstStaker.
        l2Staking.delegate(firstStaker, morphBalance / 2);
        // Verify that the delegation amount is correct.
        uint256 amount = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(morphBalance / 2, amount);
        // Delegate the remaining half of Bob's stake to firstStaker.
        l2Staking.delegate(firstStaker, morphBalance / 2);
        // Verify that the total delegation amount is correct.
        amount = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice delegate: Test that the candidate number increases when a stake is delegated.
     */
    function test_delegate_candidateNumberIncreases_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        uint256 initialCandidateNumber = l2Staking.candidateNumber();
        // Delegate Bob's stake to firstStaker.
        l2Staking.delegate(firstStaker, morphBalance);
        // Verify that the candidate number has increased by 1.
        uint256 newCandidateNumber = l2Staking.candidateNumber();
        assertEq(newCandidateNumber, initialCandidateNumber + 1);
        hevm.stopPrank();
    }

    /**
     * @notice delegate: Test that delegating a stake succeeds after rewards have started.
     */
    function test_delegate_rewardStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(secondStaker, morphBalance / 2);
        uint256 oldRanking = l2Staking.stakerRankings(secondStaker);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        // Verify the Delegated event is emitted successfully.
        uint256 beforeAmount = l2Staking.queryDelegationAmount(secondStaker, bob);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(secondStaker, bob, morphBalance / 2, beforeAmount + morphBalance / 2);
        l2Staking.delegate(secondStaker, morphBalance / 2);
        // Verify rankings is updated.
        uint256 newRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(newRanking, oldRanking - 1);
        hevm.stopPrank();
    }

    /**
     * @notice normal undelegate
     */
    function test_undelegate_rewardNotStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        (uint256 stakerAmount0, ) = l2Staking.delegateeDelegations(firstStaker);
        uint256 amount0 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(amount0, morphBalance);
        assertEq(stakerAmount0, morphBalance);

        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Undelegated(firstStaker, bob, morphBalance, 0, 0);
        l2Staking.undelegate(firstStaker, amount0);
        uint256 amount1 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(amount1, 0);

        // Verify the total stake amount of firstStaker is correct
        (uint256 stakerAmount1, ) = l2Staking.delegateeDelegations(firstStaker);
        assertEq(stakerAmount1, stakerAmount0 - amount0);
        hevm.stopPrank();
    }

    /**
     * @notice normal undelegate
     */
    function test_undelegate_rewardStarted_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        (uint256 stakerAmount0, ) = l2Staking.delegateeDelegations(firstStaker);
        uint256 amount0 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(stakerAmount0, amount0);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        // Verify the Undelegated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Undelegated(firstStaker, bob, morphBalance, 0, UNDELEGATE_LOCKED_EPOCHS + 1);
        l2Staking.undelegate(firstStaker, amount0);
        uint256 amount1 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(amount1, 0);
        // Verify the total stake amount of firstStaker is correct
        (uint256 stakerAmount1, ) = l2Staking.delegateeDelegations(firstStaker);
        assertEq(stakerAmount1, 0);
        hevm.stopPrank();
    }

    /**
     * @notice undelegate all when undelegate amount is zero
     */
    function test_undelegate_zeroAmount_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        (uint256 stakerAmount0, ) = l2Staking.delegateeDelegations(firstStaker);
        uint256 amount0 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(stakerAmount0, amount0);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        // Verify the Undelegated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Undelegated(firstStaker, bob, morphBalance, 0, UNDELEGATE_LOCKED_EPOCHS + 1);
        l2Staking.undelegate(firstStaker, 0);
        uint256 amount1 = l2Staking.queryDelegationAmount(firstStaker, bob);
        assertEq(amount1, 0);
        // Verify the total stake amount of firstStaker is correct
        (uint256 stakerAmount1, ) = l2Staking.delegateeDelegations(firstStaker);
        assertEq(stakerAmount1, 0);
        hevm.stopPrank();
    }

    /**
     * @notice undelegate, staker removed
     */
    function test_undelegateWhenStakerRemoved_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        l2Staking.delegate(thirdStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.warp(rewardStartTime);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
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
        l2Staking.removeStakers(0, removed);
        hevm.stopPrank();

        // sequenser size decrease
        assertEq(sequencer.getSequencerSet2Size(), SEQUENCER_SIZE - 1);
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 1);
        // staker ranking is 0, removed
        assertTrue(l2Staking.stakerRankings(firstStaker) == 0);

        hevm.startPrank(bob);
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 1);
        hevm.stopPrank();

        hevm.startPrank(alice);
        l2Staking.undelegate(secondStaker, l2Staking.queryDelegationAmount(secondStaker, alice));
        assertEq(l2Staking.candidateNumber(), SEQUENCER_SIZE - 2);
        hevm.stopPrank();
    }

    /**
     * @notice failed unstaking, when share is zero
     */
    function test_undelegate_zeroShare_reverts() public {
        hevm.startPrank(bob);
        hevm.expectRevert(IL2Staking.ErrZeroShares.selector);
        l2Staking.undelegate(firstStaker, 1 ether);
        hevm.stopPrank();
    }

    /**
     * @notice failed unstaking, when share is insufficient
     */
    function test_undelegate_insufficientShare_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5);
        hevm.expectRevert(IL2Staking.ErrInsufficientBalance.selector);
        l2Staking.undelegate(firstStaker, 10 ether);
        hevm.stopPrank();
    }

    /**
     * @notice undelegate: Updates rankings and candidate number when staker is not removed.
     */
    function test_undelegate_updatesRankingsAndCandidateNumber() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 1 ether);
        l2Staking.delegate(secondStaker, 1 ether);
        l2Staking.delegate(thirdStaker, 1 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.warp(rewardStartTime);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        l2Staking.delegate(firstStaker, morphBalance / 2);
        hevm.stopPrank();

        hevm.startPrank(bob);
        // Check rankings and candidate number before undelegating.
        uint256 beforeRanking = l2Staking.stakerRankings(firstStaker);
        uint256 beforeCandidateNumber = l2Staking.candidateNumber();
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));
        // Check rankings and candidate number after undelegating.
        uint256 afterRanking = l2Staking.stakerRankings(firstStaker);
        uint256 afterCandidateNumber = l2Staking.candidateNumber();
        assertTrue(afterRanking > beforeRanking);
        assertEq(afterCandidateNumber, beforeCandidateNumber - 1);
        hevm.stopPrank();
    }

    /**
     * @notice failed claim, no pending undelegate requests
     */
    function test_delegatorClaim_noPendingundelegateRequests_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        hevm.expectRevert(IL2Staking.ErrNoUndelegateRequest.selector);
        l2Staking.claimUndelegation(0);
        hevm.stopPrank();
    }

    /**
     * @notice failed claim, amount in lock period
     */
    function test_delegatorClaim_inLockPeriod_reverts() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));
        hevm.expectRevert(IL2Staking.ErrNoClaimableUndelegateRequest.selector);
        l2Staking.claimUndelegation(0);
        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegation when reward not started
     */
    function test_claimUndelegation_delegatorClaimUndelegationRewardNotStarted_succeeds() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(bob);
        uint256 balanceBefore = morphToken.balanceOf(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));

        // Expect the UndelegationClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.UndelegationClaimed(bob, l2Staking.lockedAmount(bob, 0));
        l2Staking.claimUndelegation(0);
        uint256 balanceAfter = morphToken.balanceOf(bob);
        // Verify the balance of bob is correct.
        assertEq(balanceBefore, balanceAfter);
        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegations when reward not started
     */
    function test_claimUndelegation_delegatorClaimUndelegationsRewardNotStarted_succeeds() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(bob);
        uint256 balanceBefore = morphToken.balanceOf(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance / 2);
        l2Staking.delegate(secondStaker, morphBalance / 2);

        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));
        l2Staking.undelegate(secondStaker, l2Staking.queryDelegationAmount(secondStaker, bob));

        // Expect the UndelegationClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.UndelegationClaimed(bob, l2Staking.lockedAmount(bob, 0));
        l2Staking.claimUndelegation(0);
        uint256 balanceAfter = morphToken.balanceOf(bob);
        // Verify the balance of bob is correct.
        assertEq(balanceBefore, balanceAfter);
        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegation when reward started
     */
    function test_claimUndelegation_delegatorClaimUndelegationRewardStarted_succeeds() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(bob);
        uint256 balanceBefore = morphToken.balanceOf(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));

        // Unlock time not reached
        hevm.expectRevert(IL2Staking.ErrNoClaimableUndelegateRequest.selector);
        l2Staking.claimUndelegation(0);

        // Unlock time reached
        hevm.warp(rewardStartTime + REWARD_EPOCH * (UNDELEGATE_LOCKED_EPOCHS + 2));
        // Expect the UndelegationClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.UndelegationClaimed(bob, l2Staking.lockedAmount(bob, 0));
        l2Staking.claimUndelegation(0);
        uint256 balanceAfter = morphToken.balanceOf(bob);
        // Verify the balance of bob is correct.
        assertEq(balanceBefore, balanceAfter);
        hevm.stopPrank();
    }

    /**
     * @notice normal claim undelegations when reward started
     */
    function test_claimUndelegation_delegatorClaimUndelegationsRewardStarted_succeeds() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(bob);
        uint256 balanceBefore = morphToken.balanceOf(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, morphBalance / 2);
        l2Staking.delegate(secondStaker, morphBalance / 2);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(bob);
        l2Staking.undelegate(firstStaker, l2Staking.queryDelegationAmount(firstStaker, bob));
        l2Staking.undelegate(secondStaker, l2Staking.queryDelegationAmount(secondStaker, bob));

        // Unlock time not reached
        hevm.expectRevert(IL2Staking.ErrNoClaimableUndelegateRequest.selector);
        l2Staking.claimUndelegation(0);

        // Unlock time reached
        hevm.warp(rewardStartTime + REWARD_EPOCH * (UNDELEGATE_LOCKED_EPOCHS + 2));
        // Expect the UndelegationClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.UndelegationClaimed(bob, l2Staking.lockedAmount(bob, 0));
        l2Staking.claimUndelegation(0);
        uint256 balanceAfter = morphToken.balanceOf(bob);
        // Verify the balance of bob is correct.
        assertEq(balanceBefore, balanceAfter);
        hevm.stopPrank();
    }

    /**
     * @notice delegate again after undelegate
     */
    function test_delegate_delegateAfterUndelegate_succeeds() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        assertEq(morphToken.balanceOf(bob), morphBalance);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, bob, morphBalance / 2, morphBalance / 2);
        l2Staking.delegate(firstStaker, morphBalance / 2);
        assertEq(morphToken.balanceOf(bob), morphBalance / 2);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), morphBalance / 2);

        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Undelegated(firstStaker, bob, morphBalance / 2, 0, 0);
        l2Staking.undelegate(firstStaker, morphBalance / 2);
        assertEq(morphToken.balanceOf(bob), morphBalance / 2);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), 0);

        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, bob, morphBalance / 2, morphBalance / 2);
        l2Staking.delegate(firstStaker, morphBalance / 2);
        assertEq(morphToken.balanceOf(bob), 0);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), morphBalance / 2);
        assertEq(l2Staking.lockedAmount(bob, 0), morphBalance / 2);
        assertEq(l2Staking.undelegateSequence(bob), 1);

        hevm.stopPrank();
    }

    /**
     * @notice test ranking, reward_start = false
     */
    function test_rankWhenRewardNotStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(secondStaker, morphBalance);
        uint256 secondRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(secondRanking, 1 + 1);
        uint256 firstRanking = l2Staking.stakerRankings(firstStaker);
        assertEq(firstRanking, 0 + 1);
        hevm.stopPrank();
    }

    /**
     * @notice test ranking, reward_start = true
     */
    function test_rankWhenRewardStarting_succeeds() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(secondStaker, 10 ether);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        uint256 secondRanking = l2Staking.stakerRankings(secondStaker);
        assertEq(secondRanking, 0 + 1);
        uint256 firstRanking = l2Staking.stakerRankings(firstStaker);
        assertEq(firstRanking, 1 + 1);
        address[] memory sequencerSet = sequencer.getSequencerSet2();
        assertEq(secondStaker, sequencerSet[0]);
        assertEq(firstStaker, sequencerSet[1]);
        hevm.stopPrank();
    }

    /**
     * @notice updateSequencerSetMaxSize: Reverts if called by a non-owner.
     */
    function test_updateSequencerSetMaxSize_notOwner_reverts() public {
        hevm.startPrank(address(1));
        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.updateSequencerSetMaxSize(2);
        hevm.stopPrank();
    }

    /**
     * @notice updateSequencerSetMaxSize: Reverts if set to zero.
     */
    function test_updateSequencerSetMaxSize_eqZero_reverts() public {
        hevm.startPrank(multisig);
        // Expect revert due to _sequencerSetMaxSize equals zero.
        hevm.expectRevert(IL2Staking.ErrInvalidSequencerSize.selector);
        l2Staking.updateSequencerSetMaxSize(0);
        hevm.stopPrank();
    }

    /**
     * @notice updateSequencerSetMaxSize: Reverts if set to current max size.
     */
    function test_updateSequencerSetMaxSize_eqCurrentMaxSize_reverts() public {
        hevm.startPrank(multisig);
        uint256 oldSize = l2Staking.sequencerSetMaxSize();
        // Expect revert due to _sequencerSetMaxSize equals currect sequencerSetMaxSize.
        hevm.expectRevert(IL2Staking.ErrInvalidSequencerSize.selector);
        l2Staking.updateSequencerSetMaxSize(oldSize);
        hevm.stopPrank();
    }

    /**
     * @notice update sequencerSetMaxSize
     */
    function test_updateSequencerSetMaxSize_succeeds() public {
        hevm.startPrank(multisig);
        // Expect the SequencerSetMaxSizeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.SequencerSetMaxSizeUpdated(l2Staking.sequencerSetMaxSize(), 2);
        l2Staking.updateSequencerSetMaxSize(2);
        assertEq(sequencer.getSequencerSet2Size(), 2);
        hevm.stopPrank();
    }

    /**
     * @notice updateSequencerSetMaxSize: No _updateSequencerSet call when increasing size.
     */
    function test_updateSequencerSetMaxSize_largerValue_noUpdateSequencerSet_succeeds() public {
        hevm.startPrank(multisig);
        uint256 oldSequencerSet2Size = sequencer.getSequencerSet2Size();
        uint256 oldSize = l2Staking.sequencerSetMaxSize();
        uint256 newSize = oldSize + 1;
        l2Staking.updateSequencerSetMaxSize(newSize);
        // Verify that _updateSequencerSet was not called.
        assertEq(sequencer.getSequencerSet2Size(), oldSequencerSet2Size);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Reverts if called by non-owner.
     */
    function test_updateRewardStartTime_onlyOwner_reverts() public {
        hevm.startPrank(address(1));
        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.updateRewardStartTime(block.timestamp + REWARD_EPOCH);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward start time is before current block time.
     */
    function test_updateRewardStartTime_NoChange_reverts() public {
        hevm.warp(rewardStartTime);
        uint256 oldTime = l2Staking.rewardStartTime();
        uint256 newTime = block.timestamp + REWARD_EPOCH * 2;

        hevm.startPrank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.RewardStartTimeUpdated(oldTime, newTime);
        l2Staking.updateRewardStartTime(newTime);

        hevm.expectRevert(IL2Staking.ErrInvalidStartTime.selector);
        l2Staking.updateRewardStartTime(newTime);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward start time is before current block time.
     */
    function test_updateRewardStartTime_PastTime_reverts() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(multisig);
        // Expect revert due to rewardStartTime being before block.timestamp.
        hevm.expectRevert(IL2Staking.ErrInvalidStartTime.selector);
        l2Staking.updateRewardStartTime(block.timestamp);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Reverts if reward start time is before current block time.
     */
    function test_updateRewardStartTime_ZeroTime_reverts() public {
        hevm.warp(rewardStartTime);

        hevm.startPrank(multisig);
        // Expect revert due to rewardStartTime being before block.timestamp.
        hevm.expectRevert(IL2Staking.ErrInvalidStartTime.selector);
        l2Staking.updateRewardStartTime(0);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Reverts if not a multiple of REWARD_EPOCH.
     */
    function test_updateRewardStartTime_notMultipleRewardEpoch_reverts() public {
        hevm.warp(rewardStartTime / 2);

        hevm.startPrank(multisig);
        // Expect revert due to updateRewardStartTime not REWARD_EPOCH multiple.
        hevm.expectRevert(IL2Staking.ErrInvalidStartTime.selector);
        l2Staking.updateRewardStartTime(block.timestamp + REWARD_EPOCH / 2);
        hevm.stopPrank();
    }

    /**
     * @notice updateRewardStartTime: Updates the reward start time successfully.
     */
    function test_updateRewardStartTime_succeeds() public {
        hevm.warp(rewardStartTime);
        uint256 oldTime = l2Staking.rewardStartTime();
        uint256 newTime = block.timestamp + REWARD_EPOCH * 2;

        hevm.startPrank(multisig);
        // Verify the SequencerSetMaxSizeUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.RewardStartTimeUpdated(oldTime, newTime);
        l2Staking.updateRewardStartTime(newTime);
        assertEq(l2Staking.rewardStartTime(), newTime);
        hevm.stopPrank();
    }

    /**
     * @notice startReward: Reverts if called by a non-owner.
     */
    function test_startReward_notOwner_reverts() public {
        hevm.startPrank(address(1));
        // Expect revert due to caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        l2Staking.startReward();
        hevm.stopPrank();
    }

    /**
     * @notice startReward: Reverts if called before reward start time.
     */
    function test_startReward_beforeRewardStartTime_reverts() public {
        hevm.warp(rewardStartTime / 2);

        hevm.startPrank(multisig);
        uint256 newRewardStartTime = REWARD_EPOCH + 4800;
        newRewardStartTime = ((newRewardStartTime / REWARD_EPOCH) + 1) * REWARD_EPOCH;
        // Update the reward start time.
        l2Staking.updateRewardStartTime(newRewardStartTime);

        // Expect revert due to "can't start before reward start time".
        hevm.expectRevert(IL2Staking.ErrStartTimeNotReached.selector);
        l2Staking.startReward();
        hevm.stopPrank();
    }

    /**
     * @notice startReward: Reverts if no candidates.
     */
    function test_startReward_notCandidateNumber_reverts() public {
        hevm.startPrank(multisig);
        // Expect revert due to none candidate.
        hevm.expectRevert(IL2Staking.ErrNoCandidate.selector);
        l2Staking.startReward();
        hevm.stopPrank();
    }

    /**
     * @notice update to next epoch with distribute
     */
    function _updateToNextEpochWithDistribute(uint256 sequencerSize) internal returns (uint256) {
        uint256 totalSupplyBefore = morphToken.totalSupply();
        uint256 inflation = (totalSupplyBefore * 1596535874529) / INFLATION_RATIO_PRECISION;

        hevm.startPrank(system);
        for (uint256 i = 0; i < sequencerSize; i++) {
            // per sequencer gnerate 2 blocks
            l2Staking.recordBlocks(address(uint160(beginSeq + i)));
            l2Staking.recordBlocks(address(uint160(beginSeq + i)));
        }
        // console.log("........................................");
        // console.log("current epoch: %s", l2Staking.currentEpoch());
        // console.log("total blocks: %s", l2Staking.epochTotalBlocks());
        // console.log("first staker blocks: %s", l2Staking.epochSequencerBlocks(firstStaker));
        // console.log("second staker blocks: %s", l2Staking.epochSequencerBlocks(secondStaker));
        assertEq(l2Staking.epochTotalBlocks(), 4, "total-blocks");
        assertEq(l2Staking.epochSequencerBlocks(firstStaker), 2, "blocks-firstStaker");
        assertEq(l2Staking.epochSequencerBlocks(secondStaker), 2, "blocks-secondStaker");
        // console.log("........................................");

        hevm.warp(block.timestamp + REWARD_EPOCH);
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.InflationMinted(l2Staking.currentEpoch() - 1, inflation);
        morphToken.mintInflations();
        hevm.stopPrank();

        uint256 totalSupplyAfter = morphToken.totalSupply();
        assertEq(totalSupplyAfter, totalSupplyBefore + inflation);
        // console.log("........................................");
        // console.log(totalSupplyBefore);
        // console.log(inflation);
        // console.log(totalSupplyAfter);
        // console.log("........................................");

        return inflation;
    }

    /**
     * @notice delegate: test delegate when reward started in multi epochs
     */
    function test_delegateWhenRewardStarted_mulitEpochs_succeeds() public {
        hevm.warp(rewardStartTime);
        assertEq(l2Staking.currentEpoch(), 0);

        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, alice, 5 ether, 5 ether);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(secondStaker, alice, 5 ether, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, alice), 5 ether);
        assertEq(l2Staking.queryDelegationAmount(secondStaker, alice), 5 ether);
        assertEq(morphToken.balanceOf(alice), 10 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        l2Staking.startReward();
        hevm.stopPrank();

        uint256 inflation0 = _updateToNextEpochWithDistribute(2);
        assertEq(inflation0, 159653587452900000000000, "inflation-0");
        assertEq(l2Staking.currentEpoch(), 1, "current-e1");
        assertEq(l2Staking.epochTotalBlocks(), 0, "total-blocks-e1");
        assertEq(l2Staking.epochSequencerBlocks(firstStaker), 0, "blocks-firstStaker-e1");
        assertEq(l2Staking.epochSequencerBlocks(secondStaker), 0, "blocks-secondStaker-e1");
        assertEq(l2Staking.queryDelegationAmount(firstStaker, alice), 79831793726450000000000, "alice-first-e1");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, alice), 79831793726450000000000, "alice-second-e1");
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), 0 ether, "bob-first-e1");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, bob), 0 ether, "bob-second-e1");

        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, bob, 5 ether, 79836793726450000000000);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(secondStaker, bob, 5 ether, 79836793726450000000000);
        l2Staking.delegate(secondStaker, 5 ether);
        // TODO: Should the loss of precision be handled?
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), 4999999999999988121);
        assertEq(l2Staking.queryDelegationAmount(secondStaker, bob), 4999999999999988121);
        assertEq(morphToken.balanceOf(bob), 10 ether);
        hevm.stopPrank();

        uint256 inflation1 = _updateToNextEpochWithDistribute(2);
        assertEq(inflation1, 159679076720886580788309, "inflation-1");
        assertEq(l2Staking.currentEpoch(), 2, "current-e2");
        assertEq(l2Staking.epochTotalBlocks(), 0, "total-blocks-e2");
        assertEq(l2Staking.epochSequencerBlocks(firstStaker), 0, "blocks-firstStaker-e2");
        assertEq(l2Staking.epochSequencerBlocks(secondStaker), 0, "blocks-secondStaker-e2");
        assertEq(l2Staking.queryDelegationAmount(firstStaker, alice), 159666331915002996157441, "alice-first-e2");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, alice), 159666331915002996157441, "alice-second-e2");
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), 10000171890294236712, "bob-first-e2");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, bob), 10000171890294236712, "bob-second-e2");

        hevm.startPrank(alice);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(firstStaker, alice, 5 ether, 159681332086893290394154);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.expectEmit(true, true, true, true);
        emit IL2Staking.Delegated(secondStaker, alice, 5 ether, 159681332086893290394154);
        l2Staking.delegate(secondStaker, 5 ether);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, alice), 159671331915002996157439);
        assertEq(l2Staking.queryDelegationAmount(secondStaker, alice), 159671331915002996157439);
        assertEq(morphToken.balanceOf(alice), 0 ether);
        hevm.stopPrank();

        uint256 inflation2 = _updateToNextEpochWithDistribute(2);
        assertEq(inflation2, 159704570058326237182599, "inflation-2");
        assertEq(l2Staking.currentEpoch(), 3, "current-e3");
        assertEq(l2Staking.epochTotalBlocks(), 0, "total-blocks-e3");
        assertEq(morphToken.balanceOf(alice), 0 ether);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, alice), 239518616130572639318717, "alice-first-e3");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, alice), 239518616130572639318717, "alice-second-e3");
        assertEq(morphToken.balanceOf(bob), 10 ether);
        assertEq(l2Staking.queryDelegationAmount(firstStaker, bob), 15000985483769666735, "bob-first-e3");
        assertEq(l2Staking.queryDelegationAmount(secondStaker, bob), 15000985483769666735, "bob-second-e3");
    }

    /**
     * @notice redelegate: test redelegate when reward started in multi epochs
     */
    function test_redelegateWhenRewardStarted_mulitEpochs_succeeds() public {
        // TODO multi delegator & multi delegatee
    }

    /**
     * @notice undelegate: test undelegate all when reward started in multi epochs
     */
    function test_undelegateWhenRewardStarted_undelegateAll_succeeds() public {
        // TODO multi delegator & multi delegatee
    }

    /**
     * @notice undelegate: test undelegate part when reward started in multi epochs
     */
    function test_undelegateWhenRewardStarted_undelegatePart_succeeds() public {
        // TODO multi delegator & multi delegatee
    }

    /**
     * @notice sequencer update: add one sequencer
     */
    function test_sequencerSetUpdate_addOneSequencer_succeeds() public {
        // TODO
    }

    /**
     * @notice sequencer update: remove one sequencer
     */
    function test_sequencerSetUpdate_removeOneSequencer_succeeds() public {
        // TODO
    }

    /**
     * @notice sequencer update: replace one sequencer
     */
    function test_sequencerSetUpdate_replaceOneSequencer_succeeds() public {
        // TODO
    }

    /**
     * @notice commission: test commission in multi epochs
     */
    function test_commission_claimCommissionMultiEpochs_succeeds() public {
        // TODO multi delegator & multi delegatee
    }

    /**
     * @notice commission: test commission in multi epochs with changed commission rate
     */
    function test_commission_claimCommissionMultiEpochsWithChanged_succeeds() public {
        // TODO multi delegator & multi delegatee
    }

    /**
     * @notice distribute: multi epoch no blocks
     */
    function test_distribute_multiEpochNoBlocks_succeeds() public {
        // TODO
    }

    /**
     * @notice currentEpoch
     */
    function test_currentEpoch_succeeds() public {
        hevm.warp(rewardStartTime);
        uint256 currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 0);

        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 1);

        hevm.warp(rewardStartTime + REWARD_EPOCH * 2);
        currentEpoch = l2Staking.currentEpoch();
        assertEq(currentEpoch, 2);
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
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        l2Staking.delegate(thirdStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice1);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice2);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice3);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice4);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        hevm.stopPrank();

        // check firstStaker
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(firstStaker, 10, 0);
            assertEq(total, 5);
            assertEq(delegators.length, 10);
            assertEq(delegators[0], alice);
            assertEq(delegators[1], alice1);
            assertEq(delegators[2], alice2);
            assertEq(delegators[3], alice3);
            assertEq(delegators[4], alice4);
            assertEq(delegators[5], address(0));
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(firstStaker, 1, 0);
            assertEq(total, 5);
            assertEq(delegators.length, 1);
            assertEq(delegators[0], alice);
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(firstStaker, 1, 1);
            assertEq(total, 5);
            assertEq(delegators.length, 1);
            assertEq(delegators[0], alice1);
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(firstStaker, 2, 2);
            assertEq(total, 5);
            assertEq(delegators.length, 2);
            assertEq(delegators[0], alice4);
            assertEq(delegators[1], address(0));
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(firstStaker, 10, 3);
            assertEq(total, 5);
            assertEq(delegators.length, 10);
            assertEq(delegators[0], address(0));
            assertEq(delegators[1], address(0));
        }
        // check secondStaker
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(secondStaker, 10, 0);
            assertEq(total, 3);
            assertEq(delegators.length, 10);
            assertEq(delegators[0], alice);
            assertEq(delegators[1], alice1);
            assertEq(delegators[2], alice2);
            assertEq(delegators[3], address(0));
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(secondStaker, 2, 1);
            assertEq(total, 3);
            assertEq(delegators.length, 2);
            assertEq(delegators[0], alice2);
            assertEq(delegators[1], address(0));
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(secondStaker, 2, 2);
            assertEq(total, 3);
            assertEq(delegators.length, 2);
            assertEq(delegators[0], address(0));
            assertEq(delegators[1], address(0));
        }
        // check thirdStaker
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(thirdStaker, 10, 0);
            assertEq(total, 1);
            assertEq(delegators.length, 10);
            assertEq(delegators[0], alice);
            assertEq(delegators[1], address(0));
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(thirdStaker, 1, 0);
            assertEq(total, 1);
            assertEq(delegators.length, 1);
            assertEq(delegators[0], alice);
        }
        {
            (uint256 total, address[] memory delegators) = l2Staking.getAllDelegatorsInPagination(thirdStaker, 1, 1);
            assertEq(total, 1);
            assertEq(delegators.length, 1);
            assertEq(delegators[0], address(0));
        }
    }

    /**
     * @notice isStakingTo
     */
    function test_isStakingTo_succeeds() public {
        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        l2Staking.delegate(thirdStaker, 5 ether);
        hevm.stopPrank();

        hevm.startPrank(alice);
        assertBoolEq(l2Staking.isStakingTo(firstStaker), true);
        assertBoolEq(l2Staking.isStakingTo(secondStaker), true);
        assertBoolEq(l2Staking.isStakingTo(thirdStaker), true);
        hevm.stopPrank();
    }
}
