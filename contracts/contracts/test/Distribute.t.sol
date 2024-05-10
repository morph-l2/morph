// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {Distribute} from "../l2/staking/Distribute.sol";
import {IDistribute} from "../l2/staking/IDistribute.sol";

contract DistributeTest is L2StakingBaseTest {
    address firstStaker;
    uint256 mockReward;
    uint256 mockCommission;

    function setUp() public virtual override {
        super.setUp();

        firstStaker = address(uint160(beginSeq));
        mockReward = 10 ether;
        mockCommission = 1 ether;
    }

    /**
     * @notice update epoch reward
     */
    function _updateEpochReward(uint256 epochIndex) internal {
        uint256 updateEpochNum = sequencer.getSequencerSet2Size();
        address[] memory sequencers = new address[](updateEpochNum);
        uint256[] memory delegatorRewards = new uint256[](updateEpochNum);
        uint256[] memory commissions = new uint256[](updateEpochNum);

        for (uint i = 0; i < updateEpochNum; i++) {
            sequencers[i] = address(uint160(beginSeq + i));
            delegatorRewards[i] = mockReward;
            commissions[i] = mockCommission;
        }

        hevm.prank(address(record));
        distribute.updateEpochReward(
            epochIndex,
            sequencers,
            delegatorRewards,
            commissions
        );
    }

    /**
     * @notice initialize: re-initialize
     */
    function testInitialize() public {
        hevm.expectRevert("Initializable: contract is already initialized");
        hevm.prank(multisig);
        distribute.initialize();
    }

    /**
     * @notice notifyDelegation: only l2 staking allowed
     */
    function testNotifyDelegation() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.notifyDelegation(address(0), address(0), 0, 0, 0, 0, false);
    }

    /**
     * @notice notifyUndelegation: only l2 staking allowed
     */
    function testNotifyUndelegation() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.notifyUndelegation(address(0), address(0), 0, 0, 0);
    }

    /**
     * @notice claim: check params
     * 1. only l2 staking allowed
     * 2. not mint yet
     * 3. no remaining reward
     */
    function testClaimParams() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claim(address(0), address(0), 0);

        hevm.expectRevert("not mint yet");
        hevm.prank(address(l2Staking));
        distribute.claim(address(0), address(0), 0);

        _updateEpochReward(0);
        hevm.expectRevert("no remaining reward");
        hevm.prank(address(l2Staking));
        distribute.claim(firstStaker, alice, 0);
    }

    /**
     * @notice claim
     * 1. normal clain
     * 2. all reward claimed
     */
    function test1Claim() public {
        hevm.prank(address(l2Staking));
        distribute.notifyDelegation(
            firstStaker,
            alice,
            0,
            10 ether,
            10 ether,
            1,
            true
        );

        _updateEpochReward(0);

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

        uint256 unclaimedCommissionEpoch = distribute.unclaimedCommission(
            firstStaker
        );
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
    function testClaimAll() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claimAll(address(0), 0);

        hevm.expectRevert("not mint yet");
        hevm.prank(address(l2Staking));
        distribute.claim(address(0), address(0), 0);
    }

    /**
     * @notice updateEpochReward: only record contract allowed
     */
    function testUpdateEpochReward() public {
        uint256 updateEpochNum = 1;
        address[] memory sequencers = new address[](updateEpochNum);
        uint256[] memory delegatorRewards = new uint256[](updateEpochNum);
        uint256[] memory commissions = new uint256[](updateEpochNum);

        hevm.expectRevert("only record contract allowed");
        hevm.prank(alice);
        distribute.updateEpochReward(
            0,
            sequencers,
            delegatorRewards,
            commissions
        );

        hevm.expectRevert("invalid epoch index");
        hevm.prank(address(record));
        distribute.updateEpochReward(
            1,
            sequencers,
            delegatorRewards,
            commissions
        );

        delegatorRewards = new uint256[](2);
        hevm.expectRevert("invalid data length");
        hevm.prank(address(record));
        distribute.updateEpochReward(
            0,
            sequencers,
            delegatorRewards,
            commissions
        );
    }

    /**
     * @notice claimCommission: only l2 staking allowed
     */
    function testClaimCommission() public {
        hevm.expectRevert("only l2 staking contract allowed");
        hevm.prank(alice);
        distribute.claimCommission(address(0), 0);
    }
}
