// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IMorphToken} from "../l2/system/IMorphToken.sol";

contract MorphTokenTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    function test_l2staking_contract_succeeds() public {
        assertEq(morphToken.L2_STAKING_CONTRACT(), Predeploys.L2_STAKING);
    }

    function test_distribute_contract_succeeds() public {
        assertEq(morphToken.DISTRIBUTE_CONTRACT(), Predeploys.DISTRIBUTE);
    }

    function test_record_contract_succeeds() public {
        assertEq(morphToken.RECORD_CONTRACT(), Predeploys.RECORD);
    }

    function test_name_succeeds() public {
        assertEq(morphToken.name(), "Morph");
    }

    function test_symbol_succeeds() public {
        assertEq(morphToken.symbol(), "MPH");
    }

    function test_decimals_succeeds() public {
        assertEq(morphToken.decimals(), 18);
    }

    function test_totalSupply_succeeds() public {
        assertEq(morphToken.totalSupply(), 1000000000 ether);
    }

    function test_updateRate_notOwner_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        morphToken.updateRate(1596535874529, 0);
        hevm.stopPrank();
    }

    function test_updateRate_sameAsLatestRate_reverts() public {
        uint256 count = morphToken.inflationRatesCount();
        IMorphToken.EpochInflationRate memory info = morphToken.epochInflationRates(count - 1);
        hevm.startPrank(multisig);
        hevm.expectRevert("new rate is the same as the latest rate");
        morphToken.updateRate(info.rate, 0);
        hevm.stopPrank();
    }

    function test_updateRate_epochIndex_reverts() public {
        uint256 count = morphToken.inflationRatesCount();
        IMorphToken.EpochInflationRate memory info = morphToken.epochInflationRates(count - 1);
        hevm.startPrank(multisig);
        hevm.expectRevert("effective epochs after must be greater than before");
        morphToken.updateRate(info.rate + 1, info.effectiveEpochIndex);
        hevm.stopPrank();
    }

    function test_updateRate_succeeds() public {
        uint256 count = morphToken.inflationRatesCount();
        IMorphToken.EpochInflationRate memory info = morphToken.epochInflationRates(count - 1);
        assertEq(info.effectiveEpochIndex, 0);
        assertEq(info.rate, 1596535874529);

        uint256 newRate = info.rate + 1;
        uint256 newEpochIndex = info.effectiveEpochIndex + 1;
        hevm.prank(multisig);
        morphToken.updateRate(newRate, newEpochIndex);

        uint256 newCount = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(newCount - 1).rate, newRate);
        assertEq(morphToken.epochInflationRates(newCount - 1).effectiveEpochIndex, newEpochIndex);
    }

    function test_mintInflations_notRecord_reverts() public {
        hevm.startPrank(Predeploys.DISTRIBUTE);
        hevm.expectRevert("only record contract allowed");
        morphToken.mintInflations(0);
        hevm.stopPrank();
    }

    function test_mintInflations_notStart_reverts() public {
        hevm.startPrank(Predeploys.RECORD);
        hevm.warp(block.timestamp + rewardStartTime - 2);
        hevm.expectRevert("reward is not started yet");
        morphToken.mintInflations(0);
        hevm.stopPrank();
    }

    function test_mintInflations_invalidEpoch_reverts() public {
        hevm.startPrank(Predeploys.RECORD);
        hevm.warp(block.timestamp + rewardStartTime);
        hevm.expectRevert("the specified time has not yet been reached");
        morphToken.mintInflations(0);
        hevm.stopPrank();
    }

    function test_mintInflations_check_reverts() public {
        hevm.startPrank(Predeploys.RECORD);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(block.timestamp + rewardStartTime * 2);
        morphToken.mintInflations(0);
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(morphToken.inflationMintedEpochs(), 1);
        hevm.expectRevert("all inflations minted");
        morphToken.mintInflations(0);
        hevm.stopPrank();
    }

    function test_mintInflations_succeeds() public {
        hevm.startPrank(Predeploys.RECORD);
        uint256 beforeTotal = morphToken.totalSupply();
        uint256 dbb = morphToken.balanceOf(Predeploys.DISTRIBUTE);
        hevm.warp(block.timestamp + rewardStartTime * 2);
        morphToken.mintInflations(0);
        uint256 afterTotal = morphToken.totalSupply();
        uint256 dab = morphToken.balanceOf(Predeploys.DISTRIBUTE);
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(dab - dbb, morphToken.inflation(0));
        hevm.stopPrank();
    }

    function test_approve_check_reverts() public {
        hevm.startPrank(address(0));
        hevm.expectRevert("approve from the zero address");
        morphToken.approve(alice, 100 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.expectRevert("approve to the zero address");
        morphToken.approve(address(0), 100 ether);
        hevm.stopPrank();
    }

    function test_approve_succeeds() public {
        hevm.startPrank(multisig);
        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);
        uint256 value = morphToken.allowance(multisig, alice);
        assertEq(value, 100 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_check_reverts() public {
        hevm.prank(multisig);
        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);

        hevm.startPrank(alice);
        hevm.expectRevert("insufficient allowance");
        morphToken.transferFrom(multisig, bob, 200 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_succeeds() public {
        hevm.prank(multisig);
        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);

        uint256 beforeBalance = morphToken.balanceOf(multisig);
        hevm.prank(alice);
        assert(morphToken.transferFrom(multisig, bob, 5 ether));
        uint256 afterBalance = morphToken.balanceOf(multisig);
        assertEq(beforeBalance - afterBalance, 5 ether);
    }

    function test_increaseAllowance_check_reverts() public {
        hevm.startPrank(address(0));
        hevm.expectRevert("approve from the zero address");
        morphToken.increaseAllowance(alice, 100 ether);
        hevm.stopPrank();

        hevm.startPrank(multisig);
        hevm.expectRevert("approve to the zero address");
        morphToken.increaseAllowance(address(0), 100 ether);
        hevm.stopPrank();
    }

    function test_increaseAllowance_succeeds() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(bob, 20 ether));
        assert(morphToken.increaseAllowance(bob, 10 ether));
        assertEq(morphToken.allowance(multisig, bob), 30 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_check_reverts() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        hevm.expectRevert("decreased allowance below zero");
        morphToken.decreaseAllowance(alice, 300 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_succeeds() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        morphToken.decreaseAllowance(alice, 50 ether);
        assertEq(morphToken.allowance(multisig, alice), 150 ether);
        hevm.stopPrank();
    }

    function test_balanceOf_succeeds() public {
        assertEq(morphToken.balanceOf(multisig), 1000000000 ether);
    }

    function test_inflationRatesCount_succeeds() public {
        assertEq(morphToken.inflationRatesCount(), 1);
    }

    function test_epochInflationRates_succeeds() public {
        uint256 count = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(count - 1).rate, 1596535874529);
    }

    function test_inflation_succeeds() public {
        hevm.startPrank(Predeploys.RECORD);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(block.timestamp + rewardStartTime * 2);
        morphToken.mintInflations(0);
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        hevm.stopPrank();
    }

    function test_inflationMintedEpochs_succeeds() public {
        hevm.startPrank(Predeploys.RECORD);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(block.timestamp + rewardStartTime * 2);
        assertEq(morphToken.inflationMintedEpochs(), 0);
        morphToken.mintInflations(0);
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(morphToken.inflationMintedEpochs(), 1);
        hevm.stopPrank();
    }

    function test_transfer_check_reverts() public {
        hevm.prank(address(0));
        hevm.expectRevert("transfer from the zero address");
        morphToken.transfer(alice, 10 ether);

        hevm.startPrank(multisig);
        hevm.expectRevert("transfer to the zero address");
        morphToken.transfer(address(0), 10 ether);

        hevm.expectRevert("transfer amount exceeds balance");
        morphToken.transfer(alice, type(uint256).max);
        hevm.stopPrank();
    }

    function test_transfer_succeeds() public {
        hevm.startPrank(multisig);
        bool success = morphToken.transfer(alice, 10000000 ether);
        assert(success);
        assertEq(morphToken.balanceOf(alice), 10000000 ether);
        hevm.stopPrank();
    }

    function test_allowance_succeeds() public {
        hevm.prank(multisig);
        bool success = morphToken.transfer(alice, 100 ether);
        assert(success);

        hevm.startPrank(alice);
        assert(morphToken.approve(bob, 20 ether));
        assert(morphToken.increaseAllowance(bob, 10 ether));
        assert(morphToken.decreaseAllowance(bob, 5 ether));
        assertEq(morphToken.allowance(alice, bob), 25 ether);
        hevm.stopPrank();

        hevm.prank(bob);
        assert(morphToken.transferFrom(alice, multisig, 10 ether));

        assertEq(morphToken.balanceOf(alice), 90 ether);
    }

    function test_mintInflations_errEpochIndex_reverts() public {
        uint256 morphBalance = 20 ether;

        address firstStaker = address(uint160(beginSeq));
        address secondStaker = address(uint160(beginSeq + 1));
        address thirdStaker = address(uint160(beginSeq + 2));

        hevm.startPrank(multisig);
        morphToken.transfer(bob, morphBalance);
        morphToken.transfer(alice, morphBalance);
        hevm.stopPrank();

        hevm.expectRevert("only record contract allowed");
        hevm.prank(alice);
        morphToken.mintInflations(0);

        hevm.expectRevert("reward is not started yet");
        hevm.prank(address(record));
        morphToken.mintInflations(0);

        hevm.startPrank(alice);
        morphToken.approve(address(l2Staking), type(uint256).max);
        l2Staking.delegateStake(firstStaker, 5 ether);
        l2Staking.delegateStake(secondStaker, 5 ether);
        l2Staking.delegateStake(thirdStaker, 5 ether);
        hevm.stopPrank();

        uint256 time = REWARD_EPOCH;
        hevm.warp(time);

        // start reward
        hevm.prank(multisig);
        l2Staking.startReward();

        hevm.expectRevert("the specified time has not yet been reached");
        hevm.prank(address(record));
        morphToken.mintInflations(0);

        // epoch update
        hevm.warp(time * 2);
        assertEq(l2Staking.currentEpoch(), 1);

        hevm.prank(address(record));
        hevm.expectEmit(true, true, false, true);
        emit IMorphToken.InflationMinted(0, 159653587452900000000000);
        morphToken.mintInflations(0);
    }
}
