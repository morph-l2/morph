// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import "forge-std/console2.sol";
import {ERC20PresetFixedSupply} from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import {L2Staking} from "../L2/staking/L2Staking.sol";

// import {Types} from "../libraries/common/Types.sol";

contract L2StakingTest is DSTestPlus {
    address public alice = address(1);
    address public blob = address(2);

    ERC20PresetFixedSupply morphToken;
    L2Staking iStaking;

    uint256 totalSupply = 100000000000000000000 ether;

    uint256 lock = 10;
    uint256 sequencersSize = 7;
    uint256 limit = 1000 ether;

    uint256 morphBalance = 20 ether;

    address public tom = address(5);
    address public jerry = address(6);

    address[] sequencers = [tom, jerry];

    function setUp() public {
        morphToken = new ERC20PresetFixedSupply("Morph", "MPH", totalSupply, alice);

        iStaking = new L2Staking();
        iStaking.initialize(alice, address(1), address(morphToken), sequencersSize, limit, lock);
        iStaking.testStakers(sequencers);

        hevm.startPrank(alice);
        morphToken.transfer(blob, morphBalance);
        morphToken.transfer(tom, limit + morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice using the standard erc20 token
     */
    function testMorph() public {
        assertEq(
            totalSupply,
            morphToken.balanceOf(blob) + morphToken.balanceOf(alice) + morphToken.balanceOf(tom)
        );
    }

    /**
     * @notice normal staking by delegator
     */
    function testDelegatorStaking() public {
        address[] memory stakers = iStaking.getStakers();
        assertEq(tom, stakers[0]);

        hevm.startPrank(blob);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom, morphBalance);

        uint256 amount = iStaking.stakingInfo(blob, tom);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice failed staking, staker not exists
     */
    function testDelegatorStakingToNotExistsStaker() public {
        address[] memory stakers = iStaking.getStakers();
        assertEq(tom, stakers[0]);

        hevm.startPrank(blob);
        morphToken.approve(address(iStaking), type(uint256).max);

        hevm.expectRevert("staker not exist");
        iStaking.delegateStake(alice, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice failed unstaking, when staking amount zero
     */
    function testDelegatorUnstakingIfStakingAmountZero() public {
        hevm.startPrank(blob);

        hevm.expectRevert("staking amount is zero");
        iStaking.unDelegateStake(tom);

        hevm.stopPrank();
    }

    /**
     * @notice normal unstaking
     */
    function testDelegatorUnstaking() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom, morphBalance);

        uint256 stakerAmount0 = iStaking.stakersAmount(tom);

        uint256 amount0 = iStaking.stakingInfo(blob, tom);
        iStaking.unDelegateStake(tom);
        uint256 amount1 = iStaking.stakingInfo(blob, tom);
        assertEq(amount0, amount1);

        uint256 stakerAmount1 = iStaking.stakersAmount(tom);
        assertEq(stakerAmount0, stakerAmount1);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, no record of unstaking
     */
    function testDelegatorInvalidclaim() public {
        hevm.startPrank(blob);

        hevm.expectRevert("invalid claim");
        iStaking.claim(tom);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, amount in lock period
     */
    function testDelegatorclaimInLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom, morphBalance);
        iStaking.unDelegateStake(tom);

        hevm.expectRevert("invalid claim");
        iStaking.claim(tom);

        hevm.stopPrank();
    }

    /**
     * @notice normal claim
     */
    function testDelegatorclaim() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom, morphBalance);
        iStaking.unDelegateStake(tom);

        hevm.roll(lock + 2);

        iStaking.claim(tom);

        hevm.stopPrank();
    }

    /**
     * @notice failed restaking, pre claim in lock period
     */
    function testDelegatorRestakeInLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom, morphBalance);
        iStaking.unDelegateStake(tom);
        hevm.expectRevert("not allowed");
        iStaking.delegateStake(tom, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice normal restaking
     */
    function testDelegatorRestakeAfterLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom, morphBalance);
        iStaking.unDelegateStake(tom);

        hevm.roll(lock + 2);

        iStaking.claim(tom);

        iStaking.delegateStake(tom, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice failed staking, staker own staking needs to meet the limit amount
     */
    function testStakerAmountLessThanLimit() public {
        hevm.startPrank(tom);
        morphToken.approve(address(iStaking), type(uint256).max);

        hevm.expectRevert("staking amount is not enough");
        iStaking.delegateStake(tom, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice normal staking, staker own staking meet the limit amount
     */
    function testStakerAmountEqualLimit() public {
        hevm.startPrank(tom);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom, limit);

        uint256 amount = iStaking.stakingInfo(tom, tom);
        assertEq(limit, amount);

        hevm.stopPrank();
    }

    /**
     * @notice normal staking, staker own staking greater than limit amount
     */
    function testStakerAmountGreaterThanLimit() public {
        hevm.startPrank(tom);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom, limit);

        uint256 amount0 = iStaking.stakingInfo(tom, tom);
        assertEq(limit, amount0);

        iStaking.delegateStake(tom, morphBalance);
        uint256 amount1 = iStaking.stakingInfo(tom, tom);
        assertEq(limit + morphBalance, amount1);

        hevm.stopPrank();
    }
}
