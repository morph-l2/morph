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

contract StakingWhitelistTest is L1MessageBaseTest {
    function testWhitelist_notOwner() external {
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.updateWhitelist(add, new address[](0));
    }
    function testWhitelist_add() external {
        // make add param
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(l1Staking.whitelist(bob));
    }

    function testWhitelist_remove() external {
        address[] memory add = new address[](1);
        add[0] = alice;
        address[] memory remove = new address[](1);
        remove[0] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, remove);
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
    }

    function testWhitelist_withdraw() external {
        // add to whitelist
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(l1Staking.whitelist(bob));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.prank(alice);
        l1Staking.register{value: STAKING_VALUE}(
            aliceInfo.tmKey,
            aliceInfo.blsKey
        );

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(
            bobInfo.tmKey,
            bobInfo.blsKey
        );

        // bob withdraw
        l1Staking.withdraw();
        hevm.stopPrank();

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }

    function testWhitelist_reAdd() external {
        // add to whitelist
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(l1Staking.whitelist(bob));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.prank(alice);
        l1Staking.register{value: STAKING_VALUE}(
            aliceInfo.tmKey,
            aliceInfo.blsKey
        );

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(
            bobInfo.tmKey,
            bobInfo.blsKey
        );

        // bob withdraw
        l1Staking.withdraw();
        hevm.stopPrank();

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));

        hevm.startPrank(multisig);
        address[] memory reAdd = new address[](1);
        reAdd[0] = bob;
        hevm.expectRevert("in removed list");
        l1Staking.updateWhitelist(reAdd, new address[](0));
        hevm.stopPrank();
    }

    function testWhitelist_slash() external {
        // add to whitelist
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(l1Staking.whitelist(bob));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.prank(alice);
        l1Staking.register{value: STAKING_VALUE}(
            aliceInfo.tmKey,
            aliceInfo.blsKey
        );

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(
            bobInfo.tmKey,
            bobInfo.blsKey
        );
        hevm.stopPrank();

        address[] memory sequencers = new address[](1);
        sequencers[0] = bob;
        hevm.prank(l1Staking.rollupContract());
        // bob sequencer to slash
        l1Staking.slash(sequencers);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }
}