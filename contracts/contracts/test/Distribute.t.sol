// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";

contract DistributeTest is L2StakingBaseTest {
    address public firstStaker;
    uint256 public mockReward;
    uint256 public mockCommission;

    function setUp() public virtual override {
        super.setUp();

        firstStaker = address(uint160(beginSeq));
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
        distribute.initialize();
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
     * 1. normal clain
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
}
