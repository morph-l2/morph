// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import "forge-std/console2.sol";
import {ERC20PresetFixedSupply} from "@openzeppelin/contracts/token/ERC20/presets/ERC20PresetFixedSupply.sol";
import {L2Staking} from "../L2/staking/L2Staking.sol";
import {Types} from "../libraries/common/Types.sol";

contract L2StakingTest is DSTestPlus {
    address public alice = address(1);
    address public blob = address(2);

    ERC20PresetFixedSupply morphToken;
    L2Staking iStaking;

    uint256 totalSupply = 100000000000000000000 ether;

    uint256 epoch = 10;
    uint256 sequencersSize = 7;
    uint256 limit = 1000 ether;

    uint256 morphBalance = 20 ether;

    Types.StakerInfo tom =
        Types.StakerInfo(
            address(5),
            0x0000000000000000000000000000000000000000000000000000000000000001,
            hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001"
        );

    Types.StakerInfo jerry =
        Types.StakerInfo(
            address(6),
            0x0000000000000000000000000000000000000000000000000000000000000002,
            hex"00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002"
        );

    function setUp() public {
        morphToken = new ERC20PresetFixedSupply("Morph", "MPH", totalSupply, alice);

        iStaking = new L2Staking();
        iStaking.initialize(alice, address(1), address(morphToken), sequencersSize, limit, epoch);

        Types.StakerInfo[] memory sequencers = new Types.StakerInfo[](2);
        sequencers[0] = tom;
        sequencers[1] = jerry;

        iStaking.updateStakers(sequencers, true);

        hevm.startPrank(alice);
        morphToken.transfer(blob, morphBalance);
        morphToken.transfer(tom.addr, limit + morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice using the standard erc20 token
     */
    function testMorph() public {
        assertEq(
            totalSupply,
            morphToken.balanceOf(blob) +
                morphToken.balanceOf(alice) +
                morphToken.balanceOf(tom.addr)
        );
    }

    /**
     * @notice normal staking by delegator
     */
    function testDelegatorStaking() public {
        address[] memory stakers = iStaking.getStakers();
        assertEq(tom.addr, stakers[0]);

        hevm.startPrank(blob);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom.addr, morphBalance);

        uint256 amount = iStaking.stakingInfo(tom.addr, blob);
        assertEq(morphBalance, amount);
        hevm.stopPrank();
    }

    /**
     * @notice failed staking, staker not exists
     */
    function testDelegatorStakingToNotExistsStaker() public {
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
        iStaking.unDelegateStake(tom.addr);

        hevm.stopPrank();
    }

    /**
     * @notice normal unstaking
     */
    function testDelegatorUnstaking() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom.addr, morphBalance);

        uint256 stakerAmount0 = iStaking.stakersAmount(tom.addr);

        uint256 amount0 = iStaking.stakingInfo(tom.addr, blob);
        iStaking.unDelegateStake(tom.addr);
        uint256 amount1 = iStaking.stakingInfo(tom.addr, blob);
        assertEq(amount1, 0);

        uint256 stakerAmount1 = iStaking.stakersAmount(tom.addr);

        assertEq(stakerAmount1, stakerAmount0 - amount0);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, no record of unstaking
     */
    function testDelegatorInvalidclaim() public {
        hevm.startPrank(blob);

        hevm.expectRevert("no information on unstaking");
        iStaking.claim(tom.addr);

        hevm.stopPrank();
    }

    /**
     * @notice failed claim, amount in lock period
     */
    function testDelegatorclaimInLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom.addr, morphBalance);
        iStaking.unDelegateStake(tom.addr);

        (uint256 amount, uint256 unlock) = iStaking.unstakingInfo(tom.addr, blob);

        hevm.expectRevert("claim cannot be made during the lock-up period");
        iStaking.claim(tom.addr);

        hevm.stopPrank();
    }

    /**
     * @notice normal claim
     */
    function testDelegatorclaim() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom.addr, morphBalance);
        iStaking.unDelegateStake(tom.addr);

        hevm.roll(epoch);

        iStaking.claim(tom.addr);

        hevm.stopPrank();
    }

    /**
     * @notice failed restaking, pre claim in lock period
     */
    function testDelegatorRestakeInLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom.addr, morphBalance);
        iStaking.unDelegateStake(tom.addr);
        hevm.expectRevert("re-staking cannot be made during the lock-up period");
        iStaking.delegateStake(tom.addr, morphBalance);
        hevm.stopPrank();
    }

    /**
     * @notice normal restaking
     */
    function testDelegatorRestakeAfterLockPeriod() public {
        hevm.startPrank(blob);

        morphToken.approve(address(iStaking), type(uint256).max);
        iStaking.delegateStake(tom.addr, morphBalance);
        iStaking.unDelegateStake(tom.addr);

        hevm.roll(epoch);

        iStaking.claim(tom.addr);

        iStaking.delegateStake(tom.addr, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice failed staking, staker own staking needs to meet the limit amount
     */
    function testStakerAmountLessThanLimit() public {
        hevm.startPrank(tom.addr);
        morphToken.approve(address(iStaking), type(uint256).max);

        hevm.expectRevert("staking amount is not enough");
        iStaking.delegateStake(tom.addr, morphBalance);

        hevm.stopPrank();
    }

    /**
     * @notice normal staking, staker own staking meet the limit amount
     */
    function testStakerAmountEqualLimit() public {
        hevm.startPrank(tom.addr);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom.addr, limit);

        uint256 amount = iStaking.stakingInfo(tom.addr, tom.addr);
        assertEq(limit, amount);

        hevm.stopPrank();
    }

    /**
     * @notice normal staking, staker own staking greater than limit amount
     */
    function testStakerAmountGreaterThanLimit() public {
        hevm.startPrank(tom.addr);
        morphToken.approve(address(iStaking), type(uint256).max);

        iStaking.delegateStake(tom.addr, limit);

        uint256 amount0 = iStaking.stakingInfo(tom.addr, tom.addr);
        assertEq(limit, amount0);

        iStaking.delegateStake(tom.addr, morphBalance);
        uint256 amount1 = iStaking.stakingInfo(tom.addr, tom.addr);
        assertEq(limit + morphBalance, amount1);

        hevm.stopPrank();
    }
}
