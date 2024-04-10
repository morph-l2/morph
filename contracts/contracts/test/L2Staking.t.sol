// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/console2.sol";
import {ERC20PresetFixedSupplyUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/presets/ERC20PresetFixedSupplyUpgradeable.sol";
import {L2Staking} from "../L2/staking/L2Staking.sol";
import {Types} from "../libraries/common/Types.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";

contract L2StakingTest is L2StakingBaseTest {
    ERC20PresetFixedSupplyUpgradeable morphToken;
    uint256 totalSupply = 100000000000000000000 ether;

    address morphOwner = address(999);

    uint256 limit = 1000 ether;

    uint256 morphBalance = 20 ether;

    address[] stakers;

    address firstStaker;

    function setUp() public virtual override {
        super.setUp();

        morphToken = new ERC20PresetFixedSupplyUpgradeable();
        hevm.etch(Predeploys.MORPH_TOKEN, address(morphToken).code);
        morphToken = ERC20PresetFixedSupplyUpgradeable(Predeploys.MORPH_TOKEN);
        morphToken.initialize("Morph", "MPH", totalSupply, morphOwner);

        Types.StakerInfo[] memory sequencers = new Types.StakerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address staker = address(uint160(beginSeq + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakingInfo(
                staker
            );
            sequencers[i] = Types.StakerInfo(
                stakerInfo.addr,
                stakerInfo.tmKey,
                stakerInfo.blsKey
            );

            stakers.push(staker);
        }
        firstStaker = stakers[0];

        hevm.mockCall(
            address(l2Staking.messenger()),
            abi.encodeWithSelector(
                ICrossDomainMessenger.xDomainMessageSender.selector
            ),
            abi.encode(address(l2Staking.OTHER_STAKING()))
        );

        hevm.startPrank(address(l2CrossDomainMessenger));
        l2Staking.updateStakers(sequencers, true);
        hevm.stopPrank();

        hevm.startPrank(morphOwner);
        morphToken.transfer(bob, morphBalance);
        morphToken.transfer(alice, morphBalance);
        morphToken.transfer(multisig, limit + morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice using the standard erc20 token
     */
    function testMorph() public {
        assertEq(
            totalSupply,
            morphToken.balanceOf(bob) +
                morphToken.balanceOf(alice) +
                morphToken.balanceOf(multisig) +
                morphToken.balanceOf(morphOwner)
        );
    }

    /**
     * @notice normal staking by delegator
     */
    function testDelegatorStaking() public {
        address[] memory mStakers = l2Staking.getStakers();
        assertEq(firstStaker, mStakers[0]);

        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 amount = l2Staking.stakingInfo(firstStaker, bob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice failed staking, staker not exists
     */
    function testDelegatorStakingToNotExistsStaker() public {
        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        hevm.expectRevert("staker not exist");
        l2Staking.delegateStake(alice, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice failed unstaking, when staking amount zero
     */
    function testDelegatorUnstakingIfStakingAmountZero() public {
        hevm.startPrank(bob);

        hevm.expectRevert("staking amount is zero");
        l2Staking.unDelegateStake(firstStaker);

        hevm.stopPrank();
    }

    /**
     * @notice normal unstaking
     */
    function testDelegatorUnstaking() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);

        uint256 stakerAmount0 = l2Staking.stakersAmount(firstStaker);

        uint256 amount0 = l2Staking.stakingInfo(firstStaker, bob);
        l2Staking.unDelegateStake(firstStaker);
        uint256 amount1 = l2Staking.stakingInfo(firstStaker, bob);
        assertEq(amount1, 0);

        uint256 stakerAmount1 = l2Staking.stakersAmount(firstStaker);

        assertEq(stakerAmount1, stakerAmount0 - amount0);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, no record of unstaking
     */
    function testDelegatorInvalidclaim() public {
        hevm.startPrank(bob);
        hevm.expectRevert("no information on unstaking");
        l2Staking.withdrawal(firstStaker);
        hevm.stopPrank();
    }

    /**
     * @notice failed claim, amount in lock period
     */
    function testDelegatorclaimInLockPeriod() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.unDelegateStake(firstStaker);

        (uint256 amount, uint256 unlock) = l2Staking.unstakingInfo(
            firstStaker,
            bob
        );

        hevm.expectRevert(
            "withdrawal cannot be made during the lock-up period"
        );
        l2Staking.withdrawal(firstStaker);

        hevm.stopPrank();
    }

    /**
     * @notice normal claim
     */
    function testDelegatorclaim() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.unDelegateStake(firstStaker);

        hevm.roll(ROLLUP_EPOCH);

        l2Staking.withdrawal(firstStaker);

        hevm.stopPrank();
    }

    /**
     * @notice failed restaking, pre claim in lock period
     */
    function testDelegatorRestakeInLockPeriod() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.unDelegateStake(firstStaker);
        hevm.expectRevert(
            "re-staking cannot be made during the lock-up period"
        );
        l2Staking.delegateStake(firstStaker, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice normal restaking
     */
    function testDelegatorRestakeAfterLockPeriod() public {
        hevm.startPrank(bob);

        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, morphBalance);
        l2Staking.unDelegateStake(firstStaker);

        hevm.roll(ROLLUP_EPOCH);

        l2Staking.withdrawal(firstStaker);

        l2Staking.delegateStake(firstStaker, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice update sequencer set
     */
    function testUpdateSequencerSet() public {
        address[] memory mStakers = l2Staking.getStakers();

        hevm.startPrank(bob);
        morphToken.approve(address(l2Staking), type(uint256).max);

        l2Staking.delegateStake(mStakers[0], 2 ether);
        address[] memory latestSequencerSet = l2Sequencer
            .getLatestSeqeuncerSet();
        assertEq(latestSequencerSet[0], mStakers[0]);
        assertEq(latestSequencerSet[1], mStakers[1]);
        assertEq(latestSequencerSet[2], mStakers[2]);

        l2Staking.delegateStake(mStakers[1], 8 ether);
        latestSequencerSet = l2Sequencer.getLatestSeqeuncerSet();
        assertEq(latestSequencerSet[0], mStakers[1]);
        assertEq(latestSequencerSet[1], mStakers[0]);
        assertEq(latestSequencerSet[2], mStakers[2]);

        l2Staking.delegateStake(mStakers[2], 3 ether);
        latestSequencerSet = l2Sequencer.getLatestSeqeuncerSet();
        assertEq(latestSequencerSet[0], mStakers[1]);
        assertEq(latestSequencerSet[1], mStakers[2]);
        assertEq(latestSequencerSet[2], mStakers[0]);
    }

    // /**
    //  * @notice normal staking, staker own staking meet the limit amount
    //  */
    // function testStakerAmountEqualLimit() public {
    //     hevm.startPrank(firstStaker);
    //     morphToken.approve(address(l2Staking), type(uint256).max);

    //     l2Staking.delegateStake(firstStaker, limit);

    //     uint256 amount = l2Staking.stakingInfo(firstStaker, firstStaker);
    //     assertEq(limit, amount);

    //     hevm.stopPrank();
    // }
}
