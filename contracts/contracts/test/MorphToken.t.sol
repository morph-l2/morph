// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IL2Staking} from "../l2/staking/IL2Staking.sol";
import {L2Staking} from "../l2/staking/L2Staking.sol";
import {IMorphToken} from "../l2/system/IMorphToken.sol";
import {MorphToken} from "../l2/system/MorphToken.sol";
import {console} from "forge-std/Test.sol";

contract MorphTokenTest is L2StakingBaseTest {
    function setUp() public virtual override {
        super.setUp();
    }

    function test_initialize_initializeAgain_reverts() public {
        // Test the initializer modifier to ensure initialize() can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        morphToken.initialize("Morph", "MPH", multisig, 1000000000 ether, 1596535874529);
    }

    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for MorphToken.
        TransparentUpgradeableProxy morphTokenProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new MorphToken contract.
        MorphToken morphTokenImplTemp = new MorphToken();

        // Verify that the UpdateEpochInflationRate event is emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.UpdateEpochInflationRate(1596535874529, 0);

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(morphTokenProxyTemp)).upgradeToAndCall(
            address(morphTokenImplTemp),
            abi.encodeCall(MorphToken.initialize, ("MorphTemp", "MPHT", multisig, 1000000000 ether, 1596535874529))
        );
        hevm.stopPrank();

        // // Cast the proxy contract address to the MorphToken contract type to call its methods.
        MorphToken morphTokenTemp = MorphToken(address(morphTokenProxyTemp));

        // Verify the name, symbol, owner, totalSupply, and epochInflationRates are initialized successfully.
        assertEq(morphTokenTemp.name(), "MorphTemp");
        assertEq(morphTokenTemp.symbol(), "MPHT");
        assertEq(morphTokenTemp.owner(), multisig);
        assertEq(morphTokenTemp.totalSupply(), 1000000000 ether);
        assertEq(morphTokenTemp.epochInflationRates(0).rate, 1596535874529);
    }

    function test_l2staking_contract_succeeds() public {
        assertEq(morphToken.L2_STAKING_CONTRACT(), Predeploys.L2_STAKING);
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
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.UpdateEpochInflationRate(newRate, newEpochIndex);
        hevm.startPrank(multisig);
        morphToken.updateRate(newRate, newEpochIndex);
        hevm.stopPrank();

        uint256 newCount = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(newCount - 1).rate, newRate);
        assertEq(morphToken.epochInflationRates(newCount - 1).effectiveEpochIndex, newEpochIndex);
    }

    function test_updateRate_zeroRate_succeeds() public {
        uint256 count = morphToken.inflationRatesCount();
        IMorphToken.EpochInflationRate memory info = morphToken.epochInflationRates(count - 1);
        assertEq(info.effectiveEpochIndex, 0);
        assertEq(info.rate, 1596535874529);

        uint256 newEpochIndex = info.effectiveEpochIndex + 1;
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.UpdateEpochInflationRate(0, newEpochIndex);
        hevm.startPrank(multisig);
        morphToken.updateRate(0, newEpochIndex);
        hevm.stopPrank();

        uint256 newCount = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(newCount - 1).rate, 0);
        assertEq(morphToken.epochInflationRates(newCount - 1).effectiveEpochIndex, newEpochIndex);
    }

    function test_updateRate_futureEpochIndex_succeeds() public {
        uint256 count = morphToken.inflationRatesCount();
        IMorphToken.EpochInflationRate memory info = morphToken.epochInflationRates(count - 1);
        assertEq(info.effectiveEpochIndex, 0);
        assertEq(info.rate, 1596535874529);

        uint256 newRate = info.rate + 2;
        uint256 newEpochIndex = info.effectiveEpochIndex + 100;
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.UpdateEpochInflationRate(newRate, newEpochIndex);
        hevm.startPrank(multisig);
        morphToken.updateRate(newRate, newEpochIndex);
        hevm.stopPrank();

        uint256 newCount = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(newCount - 1).rate, newRate);
        assertEq(morphToken.epochInflationRates(newCount - 1).effectiveEpochIndex, newEpochIndex);
    }

    function test_mintInflations_notSystem_reverts() public {
        hevm.startPrank(Predeploys.MORPH_TOKEN);
        hevm.expectRevert("only system address allowed");
        morphToken.mintInflations();
        hevm.stopPrank();
    }

    function test_mintInflations_notStart_reverts() public {
        hevm.startPrank(Predeploys.SYSTEM);
        hevm.warp(rewardStartTime - 1);
        hevm.expectRevert(IL2Staking.ErrRewardNotStarted.selector);
        morphToken.mintInflations();
        hevm.stopPrank();
    }

    function test_mintInflations_check_reverts() public {
        hevm.startPrank(Predeploys.SYSTEM);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        morphToken.mintInflations();
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(morphToken.inflationMintedEpochs(), 1);
        hevm.expectRevert("all inflations minted");
        morphToken.mintInflations();
        hevm.stopPrank();
    }

    function test_mintInflations_succeeds() public {
        hevm.startPrank(Predeploys.SYSTEM);
        uint256 beforeTotal = morphToken.totalSupply();
        uint256 dbb = morphToken.balanceOf(Predeploys.L2_STAKING);
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        morphToken.mintInflations();
        uint256 afterTotal = morphToken.totalSupply();
        uint256 dab = morphToken.balanceOf(Predeploys.L2_STAKING);
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(dab - dbb, morphToken.inflation(0));
        hevm.stopPrank();
    }

    function test_mintInflations_multi_epochs_succeeds() public {
        hevm.startPrank(Predeploys.SYSTEM);
        uint256 beforeTotal = morphToken.totalSupply();
        uint256 dbb = morphToken.balanceOf(Predeploys.L2_STAKING);
        hevm.warp(rewardStartTime + REWARD_EPOCH * 3);
        morphToken.mintInflations();
        uint256 afterTotal = morphToken.totalSupply();
        uint256 dab = morphToken.balanceOf(Predeploys.L2_STAKING);
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0) + morphToken.inflation(1) + morphToken.inflation(2));
        assertEq(dab - dbb, morphToken.inflation(0) + morphToken.inflation(1) + morphToken.inflation(2));
        hevm.stopPrank();
    }

    function test_mintInflations_differentRate_succeeds() public {
        hevm.startPrank(multisig);
        morphToken.updateRate(1596535874529 + 100, 1);
        hevm.stopPrank();

        hevm.startPrank(Predeploys.SYSTEM);
        uint256 oldTotal = morphToken.totalSupply();
        uint256 incrementAmount = (oldTotal * (morphToken.epochInflationRates(0).rate)) / INFLATION_RATIO_PRECISION;
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.InflationMinted(0, incrementAmount);
        morphToken.mintInflations();
        uint256 newTotal = morphToken.totalSupply();
        hevm.stopPrank();

        assertEq(incrementAmount, newTotal - oldTotal);
        assertEq(incrementAmount, morphToken.inflation(0));

        hevm.startPrank(Predeploys.SYSTEM);
        oldTotal = morphToken.totalSupply();
        incrementAmount = (oldTotal * (morphToken.epochInflationRates(1).rate)) / INFLATION_RATIO_PRECISION;
        hevm.warp(rewardStartTime + REWARD_EPOCH * 2);
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.InflationMinted(1, incrementAmount);
        morphToken.mintInflations();
        newTotal = morphToken.totalSupply();
        hevm.stopPrank();

        assertEq(incrementAmount, newTotal - oldTotal);
        assertEq(incrementAmount, morphToken.inflation(1));
    }

    function test_burn_notOwner_reverts() public {
        hevm.startPrank(alice);
        hevm.expectRevert("Ownable: caller is not the owner");
        morphToken.burn(1);
        hevm.stopPrank();
    }

    function test_burn_zeroAmount_reverts() public {
        hevm.startPrank(multisig);
        hevm.expectRevert("amount to burn is zero");
        morphToken.burn(0);
        hevm.stopPrank();
    }

    function test_burn_amountExceedsBalance_reverts() public {
        hevm.startPrank(multisig);
        uint256 oldBalance = morphToken.balanceOf(multisig);
        hevm.expectRevert("ERC20: burn amount exceeds balance");
        morphToken.burn(oldBalance + 1);
        hevm.stopPrank();
    }

    function test_burn_succeeds() public {
        hevm.startPrank(multisig);
        uint256 oldBalance = morphToken.balanceOf(multisig);
        uint256 oldTotalSupply = morphToken.totalSupply();

        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Transfer(multisig, address(0), 100);

        morphToken.burn(100);
        assertEq(morphToken.balanceOf(multisig), oldBalance - 100);
        assertEq(morphToken.totalSupply(), oldTotalSupply - 100);
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
        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Approval(multisig, alice, 100 ether);

        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);
        uint256 value = morphToken.allowance(multisig, alice);
        assertEq(value, 100 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_check_reverts() public {
        hevm.startPrank(multisig);
        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);
        hevm.stopPrank();

        hevm.startPrank(alice);
        hevm.expectRevert("insufficient allowance");
        morphToken.transferFrom(multisig, bob, 200 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_succeeds() public {
        hevm.startPrank(multisig);
        bool isApprove = morphToken.approve(alice, 100 ether);
        assert(isApprove);
        hevm.stopPrank();

        uint256 beforeBalance = morphToken.balanceOf(multisig);
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Transfer(multisig, bob, 5 ether);

        assert(morphToken.transferFrom(multisig, bob, 5 ether));
        uint256 afterBalance = morphToken.balanceOf(multisig);
        assertEq(beforeBalance - afterBalance, 5 ether);
        hevm.stopPrank();
    }

    function test_transferFrom_zeroAmount_succeeds() public {
        uint256 beforeBalance = morphToken.balanceOf(multisig);
        hevm.startPrank(alice);
        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Transfer(multisig, bob, 0);
        assert(morphToken.transferFrom(multisig, bob, 0));
        uint256 afterBalance = morphToken.balanceOf(multisig);
        assertEq(beforeBalance, afterBalance);
        hevm.stopPrank();
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

    function test_increaseAllowance_zeroAmount_succeeds() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(bob, 20 ether));
        assert(morphToken.increaseAllowance(bob, 0));
        assertEq(morphToken.allowance(multisig, bob), 20 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_check_reverts() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        hevm.expectRevert("decreased allowance below zero");
        morphToken.decreaseAllowance(alice, 300 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_zeroAddress_reverts() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        hevm.expectRevert("decreased allowance below zero");
        morphToken.decreaseAllowance(address(0), 100 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_succeeds() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        morphToken.decreaseAllowance(alice, 50 ether);
        assertEq(morphToken.allowance(multisig, alice), 150 ether);
        hevm.stopPrank();
    }

    function test_decreaseAllowance_zeroSubtractedValue_succeeds() public {
        hevm.startPrank(multisig);
        assert(morphToken.approve(alice, 200 ether));
        assert(morphToken.decreaseAllowance(alice, 0));
        assertEq(morphToken.allowance(multisig, alice), 200 ether);
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
        hevm.startPrank(Predeploys.SYSTEM);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        morphToken.mintInflations();
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        hevm.stopPrank();
    }

    function test_inflationMintedEpochs_succeeds() public {
        hevm.startPrank(Predeploys.SYSTEM);
        uint256 beforeTotal = morphToken.totalSupply();
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        assertEq(morphToken.inflationMintedEpochs(), 0);
        morphToken.mintInflations();
        uint256 afterTotal = morphToken.totalSupply();
        assertEq(afterTotal - beforeTotal, morphToken.inflation(0));
        assertEq(morphToken.inflationMintedEpochs(), 1);
        hevm.stopPrank();
    }

    function test_transfer_check_reverts() public {
        hevm.startPrank(address(0));
        hevm.expectRevert("transfer from the zero address");
        morphToken.transfer(alice, 10 ether);
        hevm.stopPrank();

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
        hevm.startPrank(multisig);
        bool success = morphToken.transfer(alice, 100 ether);
        assert(success);
        hevm.stopPrank();

        hevm.startPrank(alice);
        assert(morphToken.approve(bob, 20 ether));
        assert(morphToken.increaseAllowance(bob, 10 ether));
        assert(morphToken.decreaseAllowance(bob, 5 ether));
        assertEq(morphToken.allowance(alice, bob), 25 ether);
        hevm.stopPrank();

        hevm.startPrank(bob);
        assert(morphToken.transferFrom(alice, multisig, 10 ether));
        assertEq(morphToken.balanceOf(alice), 90 ether);
        hevm.stopPrank();
    }

    function test_mintInflations_errEpochIndex_reverts() public {
        hevm.warp(rewardStartTime - 100);

        address firstStaker = address(uint160(beginSeq));
        address secondStaker = address(uint160(beginSeq + 1));
        address thirdStaker = address(uint160(beginSeq + 2));

        hevm.startPrank(multisig);
        morphToken.transfer(bob, 20 ether);
        morphToken.transfer(alice, 20 ether);
        hevm.stopPrank();

        hevm.startPrank(alice);
        hevm.expectRevert("only system address allowed");
        morphToken.mintInflations();
        hevm.stopPrank();

        hevm.startPrank(address(system));
        hevm.expectRevert(IL2Staking.ErrRewardNotStarted.selector);
        morphToken.mintInflations();
        hevm.stopPrank();

        hevm.startPrank(alice);
        assertEq(l2Staking.candidateNumber(), 0);
        morphToken.approve(address(l2Staking), type(uint256).max);
        // console.log("........................................");
        // console.log("before delegate");
        // (uint256 amount1, ) = l2Staking.delegateeDelegations(firstStaker);
        // (uint256 amount2, ) = l2Staking.delegateeDelegations(secondStaker);
        // (uint256 amount3, ) = l2Staking.delegateeDelegations(thirdStaker);
        // console.log(amount1);
        // console.log(amount2);
        // console.log(amount3);
        // console.log("........................................");
        l2Staking.delegate(firstStaker, 5 ether);
        l2Staking.delegate(secondStaker, 5 ether);
        l2Staking.delegate(thirdStaker, 5 ether);
        // console.log("........................................");
        // console.log("after delegated");
        // (amount1, ) = l2Staking.delegateeDelegations(firstStaker);
        // (amount2, ) = l2Staking.delegateeDelegations(secondStaker);
        // (amount3, ) = l2Staking.delegateeDelegations(thirdStaker);
        // console.log(amount1);
        // console.log(amount2);
        // console.log(amount3);
        // console.log("........................................");
        assertEq(l2Staking.candidateNumber(), 3);
        hevm.stopPrank();

        // start reward before start time
        hevm.startPrank(multisig);
        hevm.expectRevert(IL2Staking.ErrStartTimeNotReached.selector);
        hevm.warp(rewardStartTime - 1);
        l2Staking.startReward();
        hevm.stopPrank();

        // start reward
        hevm.startPrank(multisig);
        hevm.warp(rewardStartTime);
        l2Staking.startReward();
        hevm.stopPrank();

        hevm.startPrank(address(system));
        hevm.expectRevert("no inflations yet");
        morphToken.mintInflations();
        hevm.stopPrank();

        // epoch update
        assertEq(l2Staking.currentEpoch(), 0);
        hevm.warp(rewardStartTime + REWARD_EPOCH * 1);
        assertEq(l2Staking.currentEpoch(), 1);
        hevm.warp(rewardStartTime + REWARD_EPOCH * 2);
        assertEq(l2Staking.currentEpoch(), 2);

        hevm.startPrank(address(system));
        hevm.expectEmit(true, true, false, true);
        emit IMorphToken.InflationMinted(0, 159653587452900000000000);
        hevm.expectEmit(true, true, false, true);
        emit IMorphToken.InflationMinted(1, 159679076720886580788309);
        morphToken.mintInflations();
        hevm.stopPrank();
    }
}
