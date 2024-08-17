// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Types} from "../libraries/common/Types.sol";
import {IL1Staking} from "../l1/staking/IL1Staking.sol";
import {L1Staking} from "../l1/staking/L1Staking.sol";
import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";

contract StakingInitializeTest is L1MessageBaseTest {
    function test_initialize_initializeAgain_revert() external {
        // verify the initialize only can be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1Staking.initialize(
            address(1),
            STAKING_VALUE,
            CHALLENGE_DEPOSIT,
            LOCK_BLOCKS,
            rewardPercentage,
            defaultGasLimitAdd,
            defaultGasLimitRemove
        );
    }

    function test_initialize_rollupZeroAddress_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid rollup contract.
        hevm.expectRevert("invalid rollup contract");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(0),
                    STAKING_VALUE,
                    CHALLENGE_DEPOSIT,
                    LOCK_BLOCKS,
                    rewardPercentage,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_stakingValueEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid staking value.
        hevm.expectRevert("invalid staking value");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(1),
                    0,
                    CHALLENGE_DEPOSIT,
                    LOCK_BLOCKS,
                    rewardPercentage,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_challengeDepositEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid staking value.
        hevm.expectRevert("invalid challenge deposit value");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (address(1), STAKING_VALUE, 0, LOCK_BLOCKS, rewardPercentage, defaultGasLimitAdd, defaultGasLimitRemove)
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_lockBlocksEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid withdrawal lock blocks.
        hevm.expectRevert("invalid withdrawal lock blocks");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(1),
                    STAKING_VALUE,
                    CHALLENGE_DEPOSIT,
                    0,
                    rewardPercentage,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_gasLimitAddEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid gas limit add staker.
        hevm.expectRevert("invalid gas limit add staker");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (address(1), STAKING_VALUE, CHALLENGE_DEPOSIT, LOCK_BLOCKS, rewardPercentage, 0, defaultGasLimitRemove)
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_gasLimitRemoveEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid gas limit remove stakers.
        hevm.expectRevert("invalid gas limit remove stakers");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (address(1), STAKING_VALUE, CHALLENGE_DEPOSIT, LOCK_BLOCKS, rewardPercentage, defaultGasLimitAdd, 0)
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_rewardPercentageEqZero_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid challenger reward percentage.
        hevm.expectRevert("invalid challenger reward percentage");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(1),
                    STAKING_VALUE,
                    CHALLENGE_DEPOSIT,
                    LOCK_BLOCKS,
                    0,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_rewardPercentageEq101_revert() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Expect revert due to invalid challenger reward percentage.
        hevm.expectRevert("invalid challenger reward percentage");
        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (
                    address(1),
                    STAKING_VALUE,
                    CHALLENGE_DEPOSIT,
                    LOCK_BLOCKS,
                    101,
                    defaultGasLimitAdd,
                    defaultGasLimitRemove
                )
            )
        );
        hevm.stopPrank();
    }

    function test_initialize_succeeds() external {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for l1StakingProxyTemp.
        TransparentUpgradeableProxy l1StakingProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Verify the GasLimitAddStakerUpdated event is emitted successful.
        hevm.expectEmit(false, false, false, true);
        emit IL1Staking.GasLimitAddStakerUpdated(0, 22);

        // Verify the GasLimitRemoveStakersUpdated event is emitted successful.
        hevm.expectEmit(false, false, false, true);
        emit IL1Staking.GasLimitRemoveStakersUpdated(0, 33);

        // Verify the RewardPercentageUpdated event is emitted successful.
        hevm.expectEmit(false, false, false, true);
        emit IL1Staking.RewardPercentageUpdated(0, 11);

        ITransparentUpgradeableProxy(address(l1StakingProxyTemp)).upgradeToAndCall(
            address(l1StakingImpl),
            abi.encodeCall(
                L1Staking.initialize,
                (address(1), STAKING_VALUE, CHALLENGE_DEPOSIT, LOCK_BLOCKS, 11, 22, 33)
            )
        );
        hevm.stopPrank();
    }
}

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

    function test_invalidStaker_reverts() external {
        // make param
        address[] memory sequencers = new address[](2);
        sequencers[0] = alice;
        sequencers[1] = bob;
        hevm.startPrank(address(emptyContract));
        hevm.expectRevert("invalid staker");
        l1Staking.getStakersBitmap(sequencers);
        hevm.stopPrank();
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
        assertTrue(l1Staking.isStaker(bob));
        assertTrue(l1Staking.isStakerInDeleteList(bob));
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

        // Verify the WhitelistUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.WhitelistUpdated(add, new address[](0));

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

        // Verify the WhitelistUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.WhitelistUpdated(add, remove);

        hevm.prank(multisig);
        l1Staking.updateWhitelist(add, remove);
        assertTrue(l1Staking.whitelist(alice));
        assertFalse(l1Staking.whitelist(bob));
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
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }

    function test_whitelist_reAdd_inRemovedList_reverts() external {
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
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.removedList(alice));
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

        // bob sequencer to slash
        uint256 bitmap = l1Staking.getStakerBitmap(bob);
        assertEq(bitmap, 4);
        address[] memory stakers = l1Staking.getStakersFromBitmap(bitmap);
        assertEq(stakers.length, 1);
        assertEq(stakers[0], bob);
        hevm.prank(l1Staking.rollupContract());
        l1Staking.slash(bitmap);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.removedList(alice));
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
        assertFalse(l1Staking.blsKeys(aliceInfo.blsKey));
        assertFalse(l1Staking.tmKeys(aliceInfo.tmKey));
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
        // bob withdraw
        hevm.expectRevert("only active staker");
        hevm.startPrank(bob);
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
        assertFalse(l1Staking.removedList(bob));

        // Verify the StakersRemoved event is emitted successfully.
        address[] memory remove = new address[](1);
        remove[0] = bob;
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.StakersRemoved(remove);

        // bob withdraw
        l1Staking.withdraw();
        hevm.stopPrank();

        // after withdraw
        // check
        assertEq(l1Staking.withdrawals(bob), block.number + l1Staking.withdrawalLockBlocks());
        assertTrue(l1Staking.isStaker(bob));
        assertTrue(l1Staking.isStakerInDeleteList(bob));

        assertTrue(l1Staking.whitelist(alice));
        assertFalse(l1Staking.whitelist(bob));

        assertFalse(l1Staking.removedList(alice));
        assertTrue(l1Staking.removedList(bob));
    }
}

contract StakingRemoveStakerTest is L1MessageBaseTest {
    // TODO
}

contract StakingSlashTest is L1MessageBaseTest {
    function test_slash_notRollup_reverts() external {
        hevm.startPrank(address(emptyContract));
        hevm.expectRevert("only rollup contract");
        l1Staking.slash(0);
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

        // bob sequencer to slash
        uint256 bitmap = l1Staking.getStakerBitmap(bob);
        hevm.prank(l1Staking.rollupContract());
        uint256 reward = l1Staking.slash(bitmap);
        assertEq(reward, (STAKING_VALUE * l1Staking.rewardPercentage()) / 100);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.removedList(alice));
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

        // bob sequencer to slash
        uint256 bitmap = l1Staking.getStakerBitmap(bob);
        hevm.prank(l1Staking.rollupContract());
        uint256 reward = l1Staking.slash(bitmap);
        assertEq(reward, (STAKING_VALUE * l1Staking.rewardPercentage()) / 100);

        // check
        assertTrue(l1Staking.whitelist(alice));
        assertFalse(l1Staking.whitelist(bob));
        assertFalse(l1Staking.removedList(alice));
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

        // Verify the Slashed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Slashed(sequencers);

        // Verify the StakersRemoved event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.StakersRemoved(sequencers);

        // sequencer to slash
        uint256 bitmap = l1Staking.getStakersBitmap(sequencers);
        hevm.prank(l1Staking.rollupContract());
        uint256 reward = l1Staking.slash(bitmap);
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

        // sequencer to slash
        uint256 bitmap = l1Staking.getStakersBitmap(sequencers);
        hevm.prank(l1Staking.rollupContract());
        uint256 reward = l1Staking.slash(bitmap);
        assertEq(reward, (2 * STAKING_VALUE * l1Staking.rewardPercentage()) / 100);
        uint256 afterBalance = l1Staking.rollupContract().balance;
        assertEq(reward, afterBalance - beforeBalance);

        // Verify the SlashRemainingClaimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.SlashRemainingClaimed(multisig, 2 * STAKING_VALUE - reward);

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
        uint256 newValue = defaultGasLimitAdd / 10;
        uint256 oldValue = l1Staking.gasLimitAddStaker();

        // Verify the GasLimitAddStakerUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.GasLimitAddStakerUpdated(oldValue, newValue);

        hevm.prank(multisig);
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
        uint256 newValue = defaultGasLimitRemove / 10;
        uint256 oldValue = l1Staking.gasLimitRemoveStakers();

        // Verify the GasLimitRemoveStakersUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.GasLimitRemoveStakersUpdated(oldValue, newValue);

        hevm.prank(multisig);
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
        uint256 newValue = 66;
        uint256 oldValue = l1Staking.rewardPercentage();

        // Verify the GasLimitRemoveStakersUpdated event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.RewardPercentageUpdated(oldValue, newValue);

        hevm.prank(multisig);
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

        // Verify the Claimed event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL1Staking.Claimed(bob, bob);

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
