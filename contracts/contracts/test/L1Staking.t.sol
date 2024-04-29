// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Types} from "../libraries/common/Types.sol";
import {IL1Staking} from "../L1/staking/IL1Staking.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";

contract StakingRegisterTest is L1MessageBaseTest {
    function testSetUpCheck() external {
        assertEq(rollup.l1StakingContract(), address(l1Staking));
        assertEq(
            address(l1Staking.OTHER_STAKING()),
            address(Predeploys.L2_STAKING)
        );
        assertEq(
            address(l1Staking.MESSENGER()),
            address(l1CrossDomainMessenger)
        );
        assertEq(l1Staking.rollupContract(), address(rollup));
        assertEq(l1Staking.stakingValue(), STAKING_VALUE);
        assertEq(l1Staking.withdrawalLockBlocks(), LOCK_BLOCKS);
        assertEq(l1Staking.rewardPercentage(), rewardPercentage);
        assertEq(l1Staking.gasLimitAddStaker(), 1000000);
        assertEq(l1Staking.gasLimitRemoveStakers(), 10000000);
    }

    function testRegisterAlice() external {
        hevm.deal(alice, 5 * STAKING_VALUE);
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(alice);
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        hevm.stopPrank();
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Registered(alice, stakerInfo.tmKey, stakerInfo.blsKey);
        l1Staking.register{value: STAKING_VALUE}(
            stakerInfo.tmKey,
            stakerInfo.blsKey
        );
        (
            address addrCheck,
            bytes32 tmKeyCheck,
            bytes memory blsKeyCheck
        ) = l1Staking.stakers(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, stakerInfo.tmKey);
        assertBytesEq(blsKeyCheck, stakerInfo.blsKey);
        assertTrue(l1Staking.blsKeys(stakerInfo.blsKey));
        assertTrue(l1Staking.tmKeys(stakerInfo.tmKey));
        hevm.stopPrank();
    }

    function testStakerWithdrawAndClaim() external {
        // register

        hevm.deal(bob, 5 * STAKING_VALUE);
        Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(bob);
        address[] memory add = new address[](1);
        add[0] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        hevm.stopPrank();
        hevm.startPrank(bob);
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Registered(bob, stakerInfo.tmKey, stakerInfo.blsKey);
        l1Staking.register{value: STAKING_VALUE}(
            stakerInfo.tmKey,
            stakerInfo.blsKey
        );
        (
            address addrCheck,
            bytes32 tmKeyCheck,
            bytes memory blsKeyCheck
        ) = l1Staking.stakers(bob);
        assertEq(addrCheck, bob);
        assertEq(tmKeyCheck, stakerInfo.tmKey);
        assertBytesEq(blsKeyCheck, stakerInfo.blsKey);
        assertTrue(l1Staking.blsKeys(stakerInfo.blsKey));
        assertTrue(l1Staking.tmKeys(stakerInfo.tmKey));
        hevm.stopPrank();

        // withdraw

        hevm.prank(bob);
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Withdrawn(bob, block.number + LOCK_BLOCKS);
        l1Staking.withdraw();
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.isStaker(bob));
        assertTrue(l1Staking.removedList(bob));
        assertEq(l1Staking.withdrawals(bob), block.number + LOCK_BLOCKS);
        hevm.stopPrank();

        // claim

        hevm.prank(bob);
        assertEq(bob.balance, 4 * STAKING_VALUE);
        hevm.roll(block.number + LOCK_BLOCKS + 1);
        l1Staking.claimWithdrawal(bob);
        assertEq(bob.balance, 5 * STAKING_VALUE);
        assertEq(l1Staking.withdrawals(bob), 0);
        hevm.stopPrank();
    }

    function testSlash() external {
        // TODO
    }

    function testUpdateParams() external {
        // TODO
    }
}
