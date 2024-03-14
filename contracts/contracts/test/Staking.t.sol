// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {L1MessageBaseTest} from "./base/L1MessageBase.t.sol";
import {Types} from "../libraries/common/Types.sol";

contract StakingRegisterTest is L1MessageBaseTest {
    function testSetUpCheck() external {
        assertEq(l1Sequencer.stakingContract(), address(staking));
        assertEq(l1Sequencer.rollupContract(), address(rollup));
        assertEq(address(l1Sequencer.OTHER_SEQUENCER()), address(l2Sequencer));
        assertEq(
            address(l1Sequencer.MESSENGER()),
            address(l1CrossDomainMessenger)
        );
        assertEq(staking.sequencerContract(), address(l1Sequencer));
        assertEq(staking.sequencersSize(), SEQUENCER_SIZE);
        assertEq(staking.limit(), MIN_DEPOSIT);
    }

    function testRegisterAlice() external {
        bytes[] memory sequencers = new bytes[](1);
        hevm.deal(alice, 5 * MIN_DEPOSIT);
        hevm.startPrank(alice);

        Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
            alice
        );
        sequencers[0] = sequencerInfo.blsKey;

        address[] memory add = new address[](1);
        address[] memory remove;
        add[0] = alice;
        staking.updateWhitelist(add, remove);

        hevm.expectEmit(true, true, true, true);
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
            defaultGasLimit,
            0
        );

        (
            address addrCheck,
            bytes32 tmKeyCheck,
            bytes memory blsKeyCheck,
            uint256 balanceCheck
        ) = staking.stakings(alice);
        assertEq(addrCheck, alice);
        assertEq(tmKeyCheck, sequencerInfo.tmKey);
        assertBytesEq(blsKeyCheck, sequencerInfo.blsKey);
        assertEq(balanceCheck, 2 * MIN_DEPOSIT);
        hevm.stopPrank();
    }

    function testRegisterExceedSequencerSize() external {
        uint256 beginSeq = 10;
        for (uint256 i = 0; i < SEQUENCER_SIZE + 1; i++) {
            address user = address(uint160(beginSeq + i));
            hevm.deal(user, 3 * MIN_DEPOSIT);
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            if (i < SEQUENCER_SIZE) {
                sequencerAddrs.push(sequencerInfo.addr);
                sequencerBLSKeys.push(sequencerInfo.blsKey);
            }

            hevm.prank(alice);
            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = user;
            staking.updateWhitelist(add, remove);

            hevm.prank(user);
            hevm.expectEmit(true, true, true, true);
            emit Registered(
                user,
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                2 * MIN_DEPOSIT
            );
            staking.register{value: 2 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defaultGasLimit,
                0
            );
            (
                address addrCheck,
                bytes32 tmKeyCheck,
                bytes memory blsKeyCheck,
                uint256 balanceCheck
            ) = staking.stakings(user);
            assertEq(addrCheck, user);
            assertEq(tmKeyCheck, sequencerInfo.tmKey);
            assertBytesEq(blsKeyCheck, sequencerInfo.blsKey);
            assertEq(balanceCheck, 2 * MIN_DEPOSIT);
        }
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address sequencerAddr = staking.sequencersAddr(i);
            bytes memory sequencerBLS = staking.sequencersBLS(i);
            address user = address(uint160(beginSeq + i));
            (address addrCheck, , bytes memory blsKeyCheck, ) = staking
                .stakings(user);
            assertEq(sequencerAddr, addrCheck);
            assertBytesEq(sequencerBLS, blsKeyCheck);
            assertEq(sequencerAddr, sequencerAddrs[i]);
            assertBytesEq(sequencerBLS, sequencerBLSKeys[i]);
        }

        {
            hevm.deal(bob, 5 * MIN_DEPOSIT);
            // bob staking 3 eth
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                bob
            );
            for (uint256 i = SEQUENCER_SIZE - 1; i > 0; i--) {
                sequencerAddrs[i] = sequencerAddrs[i - 1];
                sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
            }
            sequencerAddrs[0] = sequencerInfo.addr;
            sequencerBLSKeys[0] = sequencerInfo.blsKey;

            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = bob;
            hevm.prank(alice);
            staking.updateWhitelist(add, remove);

            hevm.expectEmit(true, true, true, true);
            version++;
            emit SequencerUpdated(sequencerAddrs, sequencerBLSKeys, version);

            hevm.prank(bob);
            staking.register{value: 3 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defaultGasLimit,
                0
            );
            address sequencerAddr = staking.sequencersAddr(0);
            bytes memory sequencerBLS = staking.sequencersBLS(0);
            assertEq(sequencerAddr, sequencerInfo.addr);
            assertBytesEq(sequencerBLS, sequencerInfo.blsKey);
        }
    }
}

contract StakingStakeAndWithdrawTest is L1MessageBaseTest {
    event Staked(address addr, uint256 balance);
    event Withdrawed(address addr, uint256 balance);
    event Claimed(address addr, uint256 balance);

    function setUp() public virtual override {
        super.setUp();
        // staking 10 -> 13
        for (uint256 i = 0; i < SEQUENCER_SIZE * 2; i++) {
            address user = address(uint160(beginSeq + i));
            hevm.deal(user, 3 * MIN_DEPOSIT);
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            address[] memory add = new address[](1);
            address[] memory remove;
            add[0] = user;
            hevm.prank(alice);
            staking.updateWhitelist(add, remove);

            hevm.expectEmit(true, true, true, true);
            emit Registered(
                user,
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                2 * MIN_DEPOSIT
            );

            hevm.prank(user);
            staking.register{value: 2 * MIN_DEPOSIT}(
                sequencerInfo.tmKey,
                sequencerInfo.blsKey,
                defaultGasLimit,
                0
            );
            (
                address addrCheck,
                bytes32 tmKeyCheck,
                bytes memory blsKeyCheck,
                uint256 balanceCheck
            ) = staking.stakings(user);
            assertEq(addrCheck, user);
            assertEq(tmKeyCheck, sequencerInfo.tmKey);
            assertBytesEq(blsKeyCheck, sequencerInfo.blsKey);
            assertEq(balanceCheck, 2 * MIN_DEPOSIT);
            if (i < SEQUENCER_SIZE) {
                sequencerAddrs.push(sequencerInfo.addr);
                sequencerBLSKeys.push(sequencerInfo.blsKey);
            }
        }
        version++;
        checkSeuqnsers();
    }

    function test_stakeETH_sequencer_stake_more() external {
        bytes memory blsKeyCheck;
        uint256 balancesPre;
        bytes memory sequencerBLS = staking.sequencersBLS(0);
        address user = address(uint160(beginSeq));
        (, , blsKeyCheck, ) = staking.stakings(user);
        assertBytesEq(blsKeyCheck, sequencerBLS);

        user = address(uint160(beginSeq + 1));
        hevm.deal(user, 3 * MIN_DEPOSIT);
        (, , blsKeyCheck, balancesPre) = staking.stakings(user);
        hevm.prank(user);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defaultGasLimit, 0);
        checkSeuqnsers();
        assertEq(user, staking.stakers(0));
    }

    function test_stakeETH_staker_stake_more() external {
        address addrCheck;
        bytes memory blsKeyCheck;
        uint256 balancesPre;
        address sequencerAddr = staking.sequencersAddr(0);
        bytes memory sequencerBLS = staking.sequencersBLS(0);
        address user = address(uint160(beginSeq));
        (addrCheck, , blsKeyCheck, ) = staking.stakings(user);
        assertEq(addrCheck, sequencerAddr);
        assertBytesEq(blsKeyCheck, sequencerBLS);

        // user(begin + 3) stake 3eth
        user = address(uint160(beginSeq + SEQUENCER_SIZE));
        hevm.deal(user, 3 * MIN_DEPOSIT);
        (addrCheck, , blsKeyCheck, balancesPre) = staking.stakings(user);
        for (uint256 i = SEQUENCER_SIZE - 1; i > 0; i--) {
            sequencerAddrs[i] = sequencerAddrs[i - 1];
            sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
        }
        sequencerAddrs[0] = addrCheck;
        sequencerBLSKeys[0] = blsKeyCheck;
        version++;
        hevm.prank(user);
        hevm.expectEmit(true, true, true, true);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defaultGasLimit, 0);
        sequencerAddr = staking.sequencersAddr(0);
        sequencerBLS = staking.sequencersBLS(0);
        checkSeuqnsers();
        assertEq(addrCheck, sequencerAddr);
        assertBytesEq(blsKeyCheck, sequencerBLS);

        // user(begin + 2) stake 3eth
        user = address(uint160(beginSeq + SEQUENCER_SIZE - 1));
        hevm.deal(user, 3 * MIN_DEPOSIT);
        (addrCheck, , blsKeyCheck, balancesPre) = staking.stakings(user);
        for (uint256 i = SEQUENCER_SIZE - 1; i > 1; i--) {
            sequencerAddrs[i] = sequencerAddrs[i - 1];
            sequencerBLSKeys[i] = sequencerBLSKeys[i - 1];
        }
        sequencerAddrs[1] = addrCheck;
        sequencerBLSKeys[1] = blsKeyCheck;
        version++;
        hevm.prank(user);
        hevm.expectEmit(true, true, true, true);
        emit Staked(user, balancesPre + MIN_DEPOSIT);
        staking.stakeETH{value: MIN_DEPOSIT}(defaultGasLimit, 0);
        sequencerAddr = staking.sequencersAddr(1);
        sequencerBLS = staking.sequencersBLS(1);
        checkSeuqnsers();
        assertEq(addrCheck, sequencerAddr);
        assertBytesEq(blsKeyCheck, sequencerBLS);
    }

    function test_staker_withdraw_claim() external {
        uint256 preStakerNum = staking.stakersNumber();

        // user(begin + 3) withdraw
        address user = address(uint160(beginSeq + SEQUENCER_SIZE));
        (, , , uint256 balancesPre) = staking.stakings(user);
        hevm.expectEmit(true, true, true, true);
        emit Withdrawed(user, balancesPre);
        hevm.prank(user);
        staking.withdrawETH(defaultGasLimit, 0);
        (, , , uint256 balancesStaking) = staking.stakings(user);
        uint256 stakerNum = staking.stakersNumber();
        (uint256 lockBalances, uint256 lockToNum, bool exit) = staking
            .withdrawals(user);
        assertEq(balancesStaking, 0);
        assertEq(stakerNum, preStakerNum - 1);
        assertEq(lockBalances, balancesPre);
        assertEq(lockToNum, block.number + LOCK);
        assertTrue(exit);

        // roll to block.number + LOCK
        uint256 preBalance = user.balance;
        hevm.roll(block.number + LOCK + 1);
        hevm.prank(user);
        hevm.expectEmit(true, true, true, true);
        emit Claimed(user, lockBalances);
        staking.claimETH();
        assertEq(user.balance, preBalance + lockBalances);
        (lockBalances, lockToNum, exit) = staking.withdrawals(user);
        assertEq(lockBalances, 0);
        assertEq(lockToNum, 0);
        assertFalse(exit);
    }

    function test_sequencer_withdraw_claim() external {
        uint256 preStakerNum = staking.stakersNumber();

        // user(begin) withdraw
        address inUser = address(uint160(beginSeq + SEQUENCER_SIZE));
        (address inAddr, , bytes memory inBlsKey, ) = staking.stakings(inUser);

        address user = address(uint160(beginSeq));
        (, , , uint256 balancesPre) = staking.stakings(user);
        // update sequencers array
        for (uint256 i = 0; i < SEQUENCER_SIZE - 1; i++) {
            sequencerAddrs[i] = sequencerAddrs[i + 1];
            sequencerBLSKeys[i] = sequencerBLSKeys[i + 1];
        }
        sequencerAddrs.pop();
        sequencerAddrs.push(inAddr);
        sequencerBLSKeys.pop();
        sequencerBLSKeys.push(inBlsKey);
        version++;
        // expect emit events
        emit Withdrawed(user, balancesPre);
        emit SequencerUpdated(sequencerAddrs, sequencerBLSKeys, version);
        // withdraw
        hevm.prank(user);
        staking.withdrawETH(defaultGasLimit, 0);
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
        assertTrue(exit);

        // roll to block.number + LOCK
        uint256 preBalance = user.balance;
        hevm.roll(block.number + LOCK + 1);
        hevm.prank(user);
        hevm.expectEmit(true, true, true, true);
        emit Claimed(user, lockBalances);
        staking.claimETH();
        assertEq(user.balance, preBalance + lockBalances);
        (lockBalances, lockToNum, exit) = staking.withdrawals(user);
        assertEq(lockBalances, 0);
        assertEq(lockToNum, 0);
        assertFalse(exit);
    }

    function checkSeuqnsers() internal {
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address sequencerAddr = staking.sequencersAddr(i);
            bytes memory sequencerBlsKey = staking.sequencersBLS(i);
            assertEq(sequencerAddrs[i], sequencerAddr);
            assertBytesEq(sequencerBLSKeys[i], sequencerBlsKey);
        }
    }
}
