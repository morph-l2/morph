// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;
import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";
import {IMorphToken} from "../l2/system/IMorphToken.sol";
import {MorphToken} from "../l2/system/MorphToken.sol";

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
        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.UpdateEpochInflationRate(newRate, newEpochIndex);
        hevm.prank(multisig);
        morphToken.updateRate(newRate, newEpochIndex);

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
        hevm.prank(multisig);
        morphToken.updateRate(0, newEpochIndex);

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

    function test_mintInflations_exceedCurrentEpoch_reverts() public {
        hevm.startPrank(Predeploys.RECORD);
        hevm.warp(block.timestamp + rewardStartTime * 2);
        uint256 exceedingEpoch = l2Staking.currentEpoch() + 1;
        hevm.expectRevert("the specified time has not yet been reached");
        morphToken.mintInflations(exceedingEpoch);
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

    function test_mintInflations_differentRate_succeeds() public {
        hevm.startPrank(multisig);
        morphToken.updateRate(1596535874529 + 100, 1);
        hevm.stopPrank();

        hevm.startPrank(Predeploys.RECORD);
        hevm.warp(block.timestamp + REWARD_EPOCH * 3);
        morphToken.mintInflations(0);
        uint256 oldTotal = morphToken.totalSupply();
        uint256 incrementAmount = (oldTotal * (morphToken.epochInflationRates(1).rate)) / INFLATION_RATIO_PRECISION;

        hevm.expectEmit(true, true, true, true);
        emit IMorphToken.InflationMinted(1, incrementAmount);

        morphToken.mintInflations(1);
        uint256 newTotal = morphToken.totalSupply();

        assertEq(incrementAmount, newTotal - oldTotal);
        assertEq(incrementAmount, morphToken.inflation(1));
        hevm.stopPrank();
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

        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Transfer(multisig, bob, 5 ether);

        assert(morphToken.transferFrom(multisig, bob, 5 ether));
        uint256 afterBalance = morphToken.balanceOf(multisig);
        assertEq(beforeBalance - afterBalance, 5 ether);
    }

    function test_transferFrom_zeroAmount_succeeds() public {
        uint256 beforeBalance = morphToken.balanceOf(multisig);
        hevm.prank(alice);

        hevm.expectEmit(true, true, true, true);
        emit IERC20Upgradeable.Transfer(multisig, bob, 0);

        assert(morphToken.transferFrom(multisig, bob, 0));
        uint256 afterBalance = morphToken.balanceOf(multisig);
        assertEq(beforeBalance, afterBalance);
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
