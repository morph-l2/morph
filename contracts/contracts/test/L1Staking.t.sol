// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Types} from "../libraries/common/Types.sol";
import {IL1Staking} from "../l1/staking/IL1Staking.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";

contract StakingTest is L1MessageBaseTest {
    function test_setUpCheck_succeeds() external {
        assertEq(rollup.l1StakingContract(), address(l1Staking));
        assertEq(address(l1Staking.OTHER_STAKING()), address(Predeploys.L2_STAKING));
        assertEq(address(l1Staking.MESSENGER()), address(l1CrossDomainMessenger));
        assertEq(l1Staking.rollupContract(), address(rollup));
        assertEq(l1Staking.stakingValue(), STAKING_VALUE);
        assertEq(l1Staking.withdrawalLockBlocks(), LOCK_BLOCKS);
        assertEq(l1Staking.rewardPercentage(), rewardPercentage);
        assertEq(l1Staking.gasLimitAddStaker(), 1000000);
        assertEq(l1Staking.gasLimitRemoveStakers(), 10000000);
    }

    function test_registerAlice_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(stakerInfo.tmKey, stakerInfo.blsKey);
        (address addrCheck, bytes32 tmKeyCheck, bytes memory blsKeyCheck) = l1Staking.stakers(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, stakerInfo.tmKey);
        assertBytesEq(blsKeyCheck, stakerInfo.blsKey);
        assertTrue(l1Staking.blsKeys(stakerInfo.blsKey));
        assertTrue(l1Staking.tmKeys(stakerInfo.tmKey));
        hevm.stopPrank();
    }

    function test_stakerWithdrawAndClaim_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(stakerInfo.tmKey, stakerInfo.blsKey);
        (address addrCheck, bytes32 tmKeyCheck, bytes memory blsKeyCheck) = l1Staking.stakers(bob);
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

    function test_slash_succeeds() external {
        // TODO
    }

    function test_updateParams_succeeds() external {
        // TODO
    }
}

contract StakingWhitelistTest is L1MessageBaseTest {
    function test_whitelist_notOwner_reverts() external {
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.updateWhitelist(add, new address[](0));
    }

    function test_whitelist_add_succeeds() external {
        // make add param
        address[] memory add = new address[](2);
        add[0] = alice;
        add[1] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(l1Staking.whitelist(bob));
    }

    function test_whitelist_remove_succeeds() external {
        address[] memory add = new address[](1);
        add[0] = alice;
        address[] memory remove = new address[](1);
        remove[0] = bob;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, remove);
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
    }

    function test_whitelist_withdraw_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // bob withdraw
        l1Staking.withdraw();
        hevm.stopPrank();

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }

    function test_whitelist_reAdd_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

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

    function test_whitelist_slash_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);
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

contract StakingRegisterTest is L1MessageBaseTest {
    function test_register_notInWhitelist_reverts() external {
        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.startPrank(alice);
        hevm.expectRevert("not in whitelist");
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);
        hevm.stopPrank();
    }

    function test_register_alreadyRegistered_reverts() external {
        // add to whitelist
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.startPrank(alice);
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        hevm.expectRevert("already registered");
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);
        hevm.stopPrank();
    }

    function test_register_notStakingValue_reverts() external {
        // add to whitelist
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.startPrank(alice);
        hevm.expectRevert("invalid staking value");
        l1Staking.register{value: STAKING_VALUE - 1}(aliceInfo.tmKey, aliceInfo.blsKey);
        hevm.stopPrank();
    }

    function test_register_notTmKey_reverts() external {
        // add to whitelist
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.startPrank(alice);
        hevm.expectRevert("invalid tendermint pubkey");
        l1Staking.register{value: STAKING_VALUE}(0, aliceInfo.blsKey);
        hevm.stopPrank();
    }

    function test_register_notBlsKey_reverts() external {
        // add to whitelist
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));

        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.startPrank(alice);
        hevm.expectRevert("invalid bls pubkey");
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, "");
        hevm.stopPrank();
    }

    function test_register_succeeds() external {
        // add to whitelist
        address[] memory add = new address[](1);
        add[0] = alice;
        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, new address[](0));

        (address addressCheck, , ) = l1Staking.stakers(alice);
        assertEq(addressCheck, address(0));
        // alice register
        Types.StakerInfo memory aliceInfo = ffi.generateStakerInfo(alice);
        assertTrue(!l1Staking.blsKeys(aliceInfo.blsKey));
        assertTrue(!l1Staking.tmKeys(aliceInfo.tmKey));
        hevm.deal(alice, 5 * STAKING_VALUE);
        hevm.prank(alice);
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        (address addrCheck, bytes32 tmKeyCheck, bytes memory blsKeyCheck) = l1Staking.stakers(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, aliceInfo.tmKey);
        assertBytesEq(blsKeyCheck, aliceInfo.blsKey);
        assertTrue(l1Staking.blsKeys(aliceInfo.blsKey));
        assertTrue(l1Staking.tmKeys(aliceInfo.tmKey));
    }
}

contract StakingWithdrawTest is L1MessageBaseTest {
    function test_withdraw_notStaker_reverts() external {
        hevm.startPrank(bob);
        // bob withdraw
        hevm.expectRevert("only staker");
        l1Staking.withdraw();
        hevm.stopPrank();
    }

    function test_withdraw_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // before withdraw
        assertEq(l1Staking.withdrawals(bob), 0);
        assertTrue(l1Staking.isStaker(bob));
        assertTrue(!l1Staking.removedList(bob));

        // bob withdraw
        l1Staking.withdraw();
        hevm.stopPrank();

        // after withdraw
        // check
        assertEq(l1Staking.withdrawals(bob), block.number + l1Staking.withdrawalLockBlocks());
        assertTrue(!l1Staking.isStaker(bob));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(l1Staking.removedList(bob));

        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }
}

contract StakingSlashTest is L1MessageBaseTest {
    function test_slash_notRollup_reverts() external {
        // make param
        address[] memory sequencers = new address[](2);
        sequencers[0] = alice;
        sequencers[1] = bob;
        hevm.startPrank(address(emptyContract));
        hevm.expectRevert("only rollup contract");
        l1Staking.slash(sequencers);
        hevm.stopPrank();
    }

    function test_slash_withdraw_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // bob withdraw
        l1Staking.withdraw();

        hevm.stopPrank();

        address[] memory sequencers = new address[](1);
        sequencers[0] = bob;
        hevm.prank(l1Staking.rollupContract());
        // bob sequencer to slash
        uint256 reward = l1Staking.slash(sequencers);
        assertEq(reward, (STAKING_VALUE * l1Staking.rewardPercentage()) / 100);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }

    function test_slash_notWithdraw_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);
        hevm.stopPrank();

        address[] memory sequencers = new address[](1);
        sequencers[0] = bob;
        hevm.prank(l1Staking.rollupContract());
        // bob sequencer to slash
        uint256 reward = l1Staking.slash(sequencers);
        assertEq(reward, (STAKING_VALUE * l1Staking.rewardPercentage()) / 100);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertTrue(!l1Staking.whitelist(bob));
        assertTrue(!l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }

    function test_slash_all_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);
        hevm.stopPrank();

        address[] memory sequencers = new address[](2);
        sequencers[0] = alice;
        sequencers[1] = bob;
        uint256 beforeBalance = l1Staking.rollupContract().balance;
        hevm.prank(l1Staking.rollupContract());
        // bob sequencer to slash
        uint256 reward = l1Staking.slash(sequencers);
        assertEq(reward, (2 * STAKING_VALUE * l1Staking.rewardPercentage()) / 100);
        uint256 afterBalance = l1Staking.rollupContract().balance;
        assertEq(reward, afterBalance - beforeBalance);
    }
}

contract StakingClaimSlashRemainingTest is L1MessageBaseTest {
    function test_claimSlashRemaining_notOwner_reverts() external {
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.claimSlashRemaining(bob);
    }

    function test_claimSlashRemaining_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);
        hevm.stopPrank();

        address[] memory sequencers = new address[](2);
        sequencers[0] = alice;
        sequencers[1] = bob;
        uint256 beforeBalance = l1Staking.rollupContract().balance;
        hevm.prank(l1Staking.rollupContract());
        // bob sequencer to slash
        uint256 reward = l1Staking.slash(sequencers);
        assertEq(reward, (2 * STAKING_VALUE * l1Staking.rewardPercentage()) / 100);
        uint256 afterBalance = l1Staking.rollupContract().balance;
        assertEq(reward, afterBalance - beforeBalance);

        uint256 beforeBalanceOfM = multisig.balance;
        hevm.prank(multisig);
        l1Staking.claimSlashRemaining(multisig);
        uint256 afterBalanceOfM = multisig.balance;
        emit log_uint(STAKING_VALUE * 2);
        emit log_uint(afterBalanceOfM - beforeBalanceOfM);
        emit log_uint(reward);
        uint256 remaining = 2 * STAKING_VALUE - reward;
        assertEq(afterBalanceOfM - beforeBalanceOfM, remaining);
    }
}

contract StakingUpdateGasLimitAddStakerTest is L1MessageBaseTest {
    function test_updateGasLimitAddStaker_notOwner_reverts() external {
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.updateGasLimitAddStaker(0);
    }

    function test_updateGasLimitAddStaker_eqZero_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid new gas limit");
        l1Staking.updateGasLimitAddStaker(0);
        // defaultGasLimitAdd
    }

    function test_updateGasLimitAddStaker_eqDefaultValue_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid new gas limit");
        l1Staking.updateGasLimitAddStaker(defaultGasLimitAdd);
    }

    function test_updateGasLimitAddStaker_succeeds() external {
        hevm.prank(multisig);
        uint256 newValue = defaultGasLimitAdd / 10;
        l1Staking.updateGasLimitAddStaker(newValue);
        assertEq(l1Staking.gasLimitAddStaker(), newValue);
    }
}

contract StakingUpdateGasLimitRemoveStakersTest is L1MessageBaseTest {
    function test_updateGasLimitRemoveStakers_notOwner_reverts() external {
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.updateGasLimitRemoveStakers(0);
    }

    function test_updateGasLimitRemoveStakers_eqZero_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid new gas limit");
        l1Staking.updateGasLimitRemoveStakers(0);
    }

    function test_updateGasLimitRemoveStakers_eqDefaultValue_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid new gas limit");
        l1Staking.updateGasLimitRemoveStakers(defaultGasLimitRemove);
    }

    function test_updateGasLimitRemoveStakers_succeeds() external {
        hevm.prank(multisig);
        uint256 newValue = defaultGasLimitRemove / 10;
        l1Staking.updateGasLimitRemoveStakers(newValue);
        assertEq(l1Staking.gasLimitRemoveStakers(), newValue);
    }
}

contract StakingUpdateRewardPercentageTest is L1MessageBaseTest {
    function test_updateRewardPercentage_notOwner_reverts() external {
        hevm.prank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        l1Staking.updateRewardPercentage(0);
    }

    function test_updateRewardPercentage_eqZero_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid reward percentage");
        l1Staking.updateRewardPercentage(0);
    }

    function test_updateRewardPercentage_ge100_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid reward percentage");
        l1Staking.updateRewardPercentage(101);
    }

    function test_updateRewardPercentage_eqDefaultValue_reverts() external {
        hevm.prank(multisig);
        hevm.expectRevert("invalid reward percentage");
        l1Staking.updateRewardPercentage(rewardPercentage);
    }

    function test_updateRewardPercentage_succeeds() external {
        hevm.prank(multisig);
        uint256 newValue = 66;
        l1Staking.updateRewardPercentage(newValue);
        assertEq(l1Staking.rewardPercentage(), newValue);
    }
}

contract StakingClaimWithdrawalTest is L1MessageBaseTest {
    function test_claimWithdrawal_notWithdraw_reverts() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        hevm.expectRevert("withdrawal not exist");
        l1Staking.claimWithdrawal(bob);
        hevm.stopPrank();
    }

    function test_claimWithdrawal_locked_reverts() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // bob withdraw
        l1Staking.withdraw();

        hevm.expectRevert("withdrawal locked");
        l1Staking.claimWithdrawal(bob);
        hevm.stopPrank();
    }

    function test_claimWithdrawal_eqLockedBlockNumber_reverts() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // bob withdraw
        l1Staking.withdraw();

        // set block number
        hevm.roll(block.number + l1Staking.withdrawalLockBlocks());
        hevm.expectRevert("withdrawal locked");
        l1Staking.claimWithdrawal(bob);
        hevm.stopPrank();
    }

    function test_claimWithdrawal_succeeds() external {
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
        l1Staking.register{value: STAKING_VALUE}(aliceInfo.tmKey, aliceInfo.blsKey);

        // bob register
        Types.StakerInfo memory bobInfo = ffi.generateStakerInfo(bob);
        hevm.deal(bob, 5 * STAKING_VALUE);
        hevm.startPrank(bob);
        l1Staking.register{value: STAKING_VALUE}(bobInfo.tmKey, bobInfo.blsKey);

        // bob withdraw
        l1Staking.withdraw();

        // set block number
        hevm.roll(block.number + l1Staking.withdrawalLockBlocks() + 1);
        uint256 beforeBalance = bob.balance;
        l1Staking.claimWithdrawal(bob);
        uint256 afterBalance = bob.balance;
        assertEq(afterBalance - beforeBalance, STAKING_VALUE);
        hevm.stopPrank();
    }
}

contract StakingVerifySignatureTest is L1MessageBaseTest {
    //    function testVerifySignature() external {
    //        l1Staking.verifySignature();
    //    }
}
