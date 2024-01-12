// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Staking} from "../L1/staking/Staking.sol";
import {Types} from "../libraries/Types.sol";
import "./CommonTest.t.sol";
import "forge-std/console.sol";

contract Staking_Register_Test is Staking_Initializer {
    function test_setUp_check() external {
        assertEq(l1Sequencer.stakingContract(), address(staking));
        assertEq(l1Sequencer.rollupContract(), address(rollup));
        assertEq(address(l1Sequencer.OTHER_SEQUENCER()), address(l2Sequencer));
        assertEq(address(l1Sequencer.MESSENGER()), address(L1Messenger));

        assertEq(staking.sequencerContract(), address(l1Sequencer));
        assertEq(staking.sequencersSize(), SEQUENCER_SIZE);
        assertEq(staking.limit(), MIN_DEPOSIT);

        assertEq(address(l2Sequencer.MESSENGER()), address(L2Messenger));
        assertEq(address(l2Sequencer.OTHER_SEQUENCER()), address(l1Sequencer));
    }

    function test_register_alice() external {
        bytes[] memory sequencers = new bytes[](1);
        vm.deal(alice, 5 * MIN_DEPOSIT);
        vm.startPrank(alice);

        Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
            alice
        );
        sequencers[0] = sequencerInfo.blsKey;

        address[] memory add = new address[](1);
        address[] memory remove;
        add[0] = alice;
        staking.updateWhitelist(add, remove);

        vm.expectEmit(true, true, true, true);
        emit Registered(
            alice,
            sequencerInfo.tmKey,
            sequencerInfo.blsKey,
            2 * MIN_DEPOSIT
        );
        version++;
        staking.register{value: 2 * MIN_DEPOSIT}(
            sequencerInfo.tmKey,
            sequencerInfo.blsKey,
            defultGasLimit
        );

        (
            address addrCheck,
            bytes32 tmKeyCheck,
            bytes memory blsKeyCheck,
            uint256 balanceCheck
        ) = staking.stakings(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, sequencerInfo.tmKey);
        assertEq(blsKeyCheck, sequencerInfo.blsKey);
        assertEq(balanceCheck, 2 * MIN_DEPOSIT);
        vm.stopPrank();
    }

    function test_register_exceed_sequencerSize() external {
        uint256 beginSeq = 10;
        for (uint256 i = 0; i < SEQUENCER_SIZE + 1; i++) {
            address user = address(uint160(beginSeq + i));
            vm.deal(user, 3 * MIN_DEPOSIT);
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            if (i < SEQUENCER_SIZE) {
                sequencerBLSKeys.push(sequencerInfo.blsKey);
            }

            vm.prank(alice);
            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = user;
            staking.updateWhitelist(add, remove);
            
            vm.prank(user);
            vm.expectEmit(true, true, true, true);
            emit Registered(
                user,
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                2 * MIN_DEPOSIT
            );
            staking.register{value: 2 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defultGasLimit
            );
            (
                address addrCheck,
                bytes32 tmKeyCheck,
                bytes memory blsKeyCheck,
                uint256 balanceCheck
            ) = staking.stakings(user);
            assertEq(addrCheck, user);
            assertEq(tmKeyCheck, sequencerInfo.tmKey);
            assertEq(blsKeyCheck, sequencerInfo.blsKey);
            assertEq(balanceCheck, 2 * MIN_DEPOSIT);
        }
        version++;
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            bytes memory sequencerBytes = staking.sequencers(i);
            address user = address(uint160(beginSeq + i));
            (, , bytes memory blsKeyCheck, ) = staking.stakings(user);
            assertEq(sequencerBytes, blsKeyCheck);
            assertEq(sequencerBytes, sequencerBLSKeys[i]);
        }

        {
            vm.deal(bob, 5 * MIN_DEPOSIT);
            // bob staking 3 eth
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                bob
            );
            for (uint256 i = SEQUENCER_SIZE - 1; i > 0; i--) {
                sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
            }
            sequencerBLSKeys[0] = sequencerInfo.blsKey;

            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = bob;
            vm.prank(alice);
            staking.updateWhitelist(add, remove);

            vm.expectEmit(true, true, true, true);
            version++;
            emit SequencerUpdated(sequencerBLSKeys, version);

            vm.prank(bob);
            staking.register{value: 3 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defultGasLimit
            );
            bytes memory sequencerBytes = staking.sequencers(0);
            assertEq(sequencerBytes, sequencerInfo.blsKey);
        }
    }
}

contract Staking_StakeAndWithdraw_Test is Staking_Initializer {
    function setUp() public virtual override {
        super.setUp();
        // staking 10 -> 13
        for (uint256 i = 0; i < SEQUENCER_SIZE * 2; i++) {
            address user = address(uint160(beginSeq + i));
            vm.deal(user, 3 * MIN_DEPOSIT);
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = user;
            vm.prank(alice);
            staking.updateWhitelist(add, remove);

            vm.expectEmit(true, true, true, true);
            emit Registered(
                user,
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                2 * MIN_DEPOSIT
            );

            vm.prank(user);
            staking.register{value: 2 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defultGasLimit
            );
            (
                address addrCheck,
                bytes32 tmKeyCheck,
                bytes memory blsKeyCheck,
                uint256 balanceCheck
            ) = staking.stakings(user);
            assertEq(addrCheck, user);
            assertEq(tmKeyCheck, sequencerInfo.tmKey);
            assertEq(blsKeyCheck, sequencerInfo.blsKey);
            assertEq(balanceCheck, 2 * MIN_DEPOSIT);
            if (i < SEQUENCER_SIZE) {
                sequencerBLSKeys.push(sequencerInfo.blsKey);
            }
        }
        version++;
        checkSeuqnsers();
    }

    function test_stakeETH_sequencer_stake_more() external {
        bytes memory blsKeyCheck;
        uint256 balancesPre;
        bytes memory sequencerBls = staking.sequencers(0);
        address user = address(uint160(beginSeq));
        (, , blsKeyCheck, ) = staking.stakings(user);
        assertEq(blsKeyCheck, sequencerBls);

        user = address(uint160(beginSeq + 1));
        vm.deal(user, 3 * MIN_DEPOSIT);
        (, , blsKeyCheck, balancesPre) = staking.stakings(user);
        vm.prank(user);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defultGasLimit);
        checkSeuqnsers();
        assertEq(user, staking.stakers(0));
    }

    function test_stakeETH_staker_stake_more() external {
        bytes memory blsKeyCheck;
        uint256 balancesPre;
        bytes memory sequencerBls = staking.sequencers(0);
        address user = address(uint160(beginSeq));
        (, , blsKeyCheck, ) = staking.stakings(user);
        assertEq(blsKeyCheck, sequencerBls);

        // user(begin + 3) stake 3eth
        user = address(uint160(beginSeq + SEQUENCER_SIZE));
        vm.deal(user, 3 * MIN_DEPOSIT);
        (, , blsKeyCheck, balancesPre) = staking.stakings(user);
        for (uint256 i = SEQUENCER_SIZE - 1; i > 0; i--) {
            sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
        }
        sequencerBLSKeys[0] = blsKeyCheck;
        version++;
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defultGasLimit);
        sequencerBls = staking.sequencers(0);
        checkSeuqnsers();
        assertEq(blsKeyCheck, sequencerBls);

        // user(begin + 2) stake 3eth
        user = address(uint160(beginSeq + SEQUENCER_SIZE - 1));
        vm.deal(user, 3 * MIN_DEPOSIT);
        (, , blsKeyCheck, balancesPre) = staking.stakings(user);
        for (uint256 i = SEQUENCER_SIZE - 1; i > 1; i--) {
            sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
        }
        sequencerBLSKeys[1] = blsKeyCheck;
        version++;
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defultGasLimit);
        sequencerBls = staking.sequencers(1);
        checkSeuqnsers();
        assertEq(blsKeyCheck, sequencerBls);
    }

    function test_staker_withdraw_claim() external {
        uint256 preStakerNum = staking.stakersNumber();

        // user(begin + 3) withdraw
        address user = address(uint160(beginSeq + SEQUENCER_SIZE));
        (, , , uint256 balancesPre) = staking.stakings(user);
        vm.expectEmit(true, true, true, true);
        emit Withdrawed(user, balancesPre);
        vm.prank(user);
        staking.withdrawETH(defultGasLimit);
        (, , , uint256 balancesStaking) = staking.stakings(user);
        uint256 stakerNum = staking.stakersNumber();
        (uint256 lockBalances, uint256 lockToNum, bool exit) = staking
            .withdrawals(user);
        assertEq(balancesStaking, 0);
        assertEq(stakerNum, preStakerNum - 1);
        assertEq(lockBalances, balancesPre);
        assertEq(lockToNum, block.number + LOCK);
        assertEq(exit, true);

        // roll to block.number + LOCK
        uint256 preBalance = user.balance;
        vm.roll(block.number + LOCK + 1);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit Claimed(user, lockBalances);
        staking.claimETH();
        assertEq(user.balance, preBalance + lockBalances);
        (lockBalances, lockToNum, exit) = staking.withdrawals(user);
        assertEq(lockBalances, 0);
        assertEq(lockToNum, 0);
        assertEq(exit, false);
    }

    function test_sequencer_withdraw_claim() external {
        uint256 preStakerNum = staking.stakersNumber();

        // user(begin) withdraw
        address inUser = address(uint160(beginSeq + SEQUENCER_SIZE));
        (, , bytes memory inBlsKey, ) = staking.stakings(inUser);

        address user = address(uint160(beginSeq));
        (, , , uint256 balancesPre) = staking.stakings(user);
        // update sequencers array
        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            sequencerBLSKeys[i] = sequencerBLSKeys[i + 1];
        }
        sequencerBLSKeys.pop();
        sequencerBLSKeys.push(inBlsKey);
        version++;
        // expect emit events
        emit Withdrawed(user, balancesPre);
        emit SequencerUpdated(sequencerBLSKeys, version);
        // withdraw
        vm.prank(user);
        staking.withdrawETH(defultGasLimit);
        // check params
        (, , , uint256 balancesStaking) = staking.stakings(user);
        uint256 stakerNum = staking.stakersNumber();
        (uint256 lockBalances, uint256 lockToNum, bool exit) = staking
            .withdrawals(user);
        address checkAddress = staking.stakers(SEQUENCER_SIZE - 1);
        checkSeuqnsers();
        assertEq(inUser, checkAddress);
        assertEq(balancesStaking, 0);
        assertEq(stakerNum, preStakerNum - 1);
        assertEq(lockBalances, balancesPre);
        assertEq(lockToNum, block.number + LOCK);
        assertEq(exit, true);

        // roll to block.number + LOCK
        uint256 preBalance = user.balance;
        vm.roll(block.number + LOCK + 1);
        vm.prank(user);
        vm.expectEmit(true, true, true, true);
        emit Claimed(user, lockBalances);
        staking.claimETH();
        assertEq(user.balance, preBalance + lockBalances);
        (lockBalances, lockToNum, exit) = staking.withdrawals(user);
        assertEq(lockBalances, 0);
        assertEq(lockToNum, 0);
        assertEq(exit, false);
    }

    function checkSeuqnsers() internal {
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            bytes memory sequencerBlsKey = staking.sequencers(i);
            assertEq(sequencerBLSKeys[i], sequencerBlsKey);
        }
    }
}
