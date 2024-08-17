// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IDistribute} from "../l2/staking/IDistribute.sol";

contract DistributeTest is L2StakingBaseTest {
    address public firstStaker;
    address public secondStaker;
    uint256 public mockReward;
    uint256 public mockCommission;

    function setUp() public virtual override {
        super.setUp();

        firstStaker = address(uint160(beginSeq));
        secondStaker = address(uint160(beginSeq + 1));
        mockReward = 10 ether;
        mockCommission = 1 ether;
    }

    /**
     * @notice update epoch reward
     */
    function _update_epoch_reward(uint256 epochIndex) internal {
        uint256 updateEpochNum = sequencer.getSequencerSet2Size();
        address[] memory sequencers = new address[](updateEpochNum);
        uint256[] memory delegatorRewards = new uint256[](updateEpochNum);
        uint256[] memory commissions = new uint256[](updateEpochNum);

        for (uint256 i = 0; i < updateEpochNum; i++) {
            sequencers[i] = address(uint160(beginSeq + i));
            delegatorRewards[i] = mockReward;
            commissions[i] = mockCommission;
        }

        hevm.prank(address(record));
        distribute.updateEpochReward(epochIndex, sequencers, delegatorRewards, commissions);
    }

    /**
     * @notice initialize: re-initialize
     */
    function test_initialize_onlyOnce_reverts() public {
        hevm.expectRevert("Initializable: contract is already initialized");
        hevm.prank(multisig);
        distribute.initialize(multisig);
    }

    /**
     * @notice notifyDelegation: only l2 staking allowed
     */
    function test_notifyDelegation_onlyCaller_reverts() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.notifyDelegation(address(0), address(0), 0, 0, 0, 0, false);
    }

    /**
     * @notice notifyUndelegation: only l2 staking allowed
     */
    function test_notifyUndelegation_onlyCaller_reverts() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.notifyUndelegation(address(0), address(0), 0, 0, 0);
    }

    /**
     * @notice claim: check params
     * 1. only l2 staking allowed
     * 2. not minted yet
     * 3. no remaining reward
     */
    function test_claim_paramsCheck_reverts() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claim(address(0), address(0), 0);

        hevm.expectRevert("not minted yet");
        hevm.prank(address(l2Staking));
        distribute.claim(address(0), address(0), 0);

        _update_epoch_reward(0);
        hevm.expectRevert("no remaining reward");
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, 0);
    }

    /**
     * @notice claim
     * 1. normal claim
     * 2. all reward claimed
     */
    function test_claim_succeeds() public {
        hevm.prank(address(l2Staking));
        distribute.notifyDelegation(firstStaker, alice, 0, 10 ether, 10 ether, 1, true);

        _update_epoch_reward(0);

        // mock inflation
        hevm.prank(multisig);
        morphToken.transfer(address(distribute), 100 ether);

        uint256 reward = distribute.queryUnclaimed(firstStaker, alice);
        assertEq(reward, mockReward);

        // Verify the RewardClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IDistribute.RewardClaimed(alice, firstStaker, 0, mockReward);

        // delegator claim
        uint256 balanceBefore = morphToken.balanceOf(alice);
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, 0);
        uint256 balanceAfter = morphToken.balanceOf(alice);
        assertEq(balanceAfter - balanceBefore, mockReward);

        uint256 unclaimedCommissionEpoch = distribute.nextEpochToClaimCommission(firstStaker);
        assertEq(unclaimedCommissionEpoch, 0);

        // delegatee claimCommission
        balanceBefore = morphToken.balanceOf(firstStaker);
        hevm.prank(address(l2Staking));
        distribute.claimCommission(firstStaker, 0);
        balanceAfter = morphToken.balanceOf(firstStaker);
        assertEq(balanceAfter - balanceBefore, mockCommission);

        hevm.expectRevert("all reward claimed");
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, 0);
    }

    /**
     * @notice claim
     * 1. normal claim
     * 2. all reward claimed
     * 3. target epoch index > minted_count - 1
     */
    function test_claim_works() public {
        hevm.prank(address(l2Staking));
        distribute.notifyDelegation(firstStaker, alice, 0, 10 ether, 10 ether, 1, true);

        _update_epoch_reward(0);

        // mock inflation
        hevm.prank(multisig);
        morphToken.transfer(address(distribute), 100 ether);

        uint256 reward = distribute.queryUnclaimed(firstStaker, alice);
        assertEq(reward, mockReward);

        uint256 target_epoch_index = 4;

        // delegator claim
        uint256 balanceBefore = morphToken.balanceOf(alice);
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, target_epoch_index);
        uint256 balanceAfter = morphToken.balanceOf(alice);
        assertEq(balanceAfter - balanceBefore, mockReward);

        uint256 unclaimedCommissionEpoch = distribute.nextEpochToClaimCommission(firstStaker);
        assertEq(unclaimedCommissionEpoch, 0);

        // delegatee claimCommission
        balanceBefore = morphToken.balanceOf(firstStaker);
        hevm.prank(address(l2Staking));
        distribute.claimCommission(firstStaker, target_epoch_index);
        balanceAfter = morphToken.balanceOf(firstStaker);
        assertEq(balanceAfter - balanceBefore, mockCommission);

        hevm.expectRevert("all reward claimed");
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, target_epoch_index);
    }

    /**
     * @notice claimAll: only l2 staking allowed
     */
    function test_claimAll_paramsCheck_reverts() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claimAll(address(0), 0);

        hevm.expectRevert("not minted yet");
        hevm.prank(address(l2Staking));
        distribute.claim(address(0), address(0), 0);
    }

    /**
     * @notice claimAll: Test claiming rewards from multiple delegatees.
     */
    function test_claimAll_multipleDelegatees_succeeds() public {
        // Notify delegation from two stakers: 10 ether and 5 ether to Alice
        hevm.startPrank(address(l2Staking));
        distribute.notifyDelegation(firstStaker, alice, 0, 10 ether, 10 ether, 1, true);
        distribute.notifyDelegation(secondStaker, alice, 0, 5 ether, 5 ether, 1, true);
        hevm.stopPrank();

        // Update the epoch reward for epoch 0.
        _update_epoch_reward(0);

        // Transfer 100 ether to the distribute contract from multisig
        hevm.prank(multisig);
        morphToken.transfer(address(distribute), 100 ether);

        uint256 rewardBefore = morphToken.balanceOf(alice);

        // Claim all rewards for Alice for epoch 0
        hevm.prank(address(l2Staking));
        distribute.claimAll(alice, 0);
        uint256 rewardAfter = morphToken.balanceOf(alice);

        // Verify Alice claimed the expected amount of rewards (mockReward * 2).
        assertEq(rewardAfter, rewardBefore + mockReward * 2);
    }

    /**
     * @notice updateEpochReward: only record contract allowed
     */
    function test_updateEpochReward_paramsCheck_reverts() public {
        uint256 updateEpochNum = 1;
        address[] memory sequencers = new address[](updateEpochNum);
        uint256[] memory delegatorRewards = new uint256[](updateEpochNum);
        uint256[] memory commissions = new uint256[](updateEpochNum);

        hevm.expectRevert("only record contract allowed");
        hevm.prank(alice);
        distribute.updateEpochReward(0, sequencers, delegatorRewards, commissions);

        hevm.expectRevert("invalid epoch index");
        hevm.prank(address(record));
        distribute.updateEpochReward(1, sequencers, delegatorRewards, commissions);

        delegatorRewards = new uint256[](2);
        hevm.expectRevert("invalid data length");
        hevm.prank(address(record));
        distribute.updateEpochReward(0, sequencers, delegatorRewards, commissions);
    }

    /**
     * @notice claimCommission: only l2 staking allowed
     */
    function test_claimCommission_onlyCaller_reverts() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claimCommission(address(0), 0);
    }

    /**
     * @notice claimCommission: expect revert if tokens are not minted yet
     */
    function test_claimCommission_notMinted_reverts() public {
        hevm.expectRevert("not minted yet");
        hevm.prank(address(l2Staking));
        distribute.claimCommission(address(0), 0);
    }

    /**
     * @notice claimCommission: claim commission and update nextEpochToClaimCommission
     */
    function test_claimCommission_succeeds() public {
        // Simulate l2Staking address to notify delegation.
        hevm.prank(address(l2Staking));
        distribute.notifyDelegation(firstStaker, alice, 0, 10 ether, 10 ether, 1, true);

        // Update the epoch reward for epoch 0.
        _update_epoch_reward(0);

        // Transfer 10 ether to the distribute contract from multisig.
        hevm.prank(multisig);
        morphToken.transfer(address(distribute), 10 ether);

        // Expect the CommissionClaimed event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IDistribute.CommissionClaimed(firstStaker, 0, 1 ether);

        uint256 beforeReward = morphToken.balanceOf(firstStaker);

        // Simulate l2Staking address to claim the commission for the first staker for epoch 0
        hevm.prank(address(l2Staking));
        distribute.claimCommission(firstStaker, 0);
        uint256 afterReward = morphToken.balanceOf(firstStaker);

        // Verify the reward after claiming is the reward before plus the mock commission.
        assertEq(afterReward, beforeReward + mockCommission);

        // Verify the next epoch to claim commission for firstStaker is updated to 1.
        assertEq(distribute.nextEpochToClaimCommission(firstStaker), 1);
    }

    /**
     * @notice claimCommission: expect revert "all commission claimed"
     */
    function test_claimCommission_allCommissionClaimed_reverts() public {
        hevm.prank(address(l2Staking));
        distribute.notifyDelegation(firstStaker, alice, 0, 10 ether, 10 ether, 1, true);
        _update_epoch_reward(0);

        hevm.prank(multisig);
        morphToken.transfer(address(distribute), 10 ether);

        hevm.prank(address(l2Staking));
        distribute.claimCommission(firstStaker, 0);

        hevm.expectRevert("all commission claimed");
        hevm.prank(address(l2Staking));
        distribute.claimCommission(firstStaker, 0);
    }
}
