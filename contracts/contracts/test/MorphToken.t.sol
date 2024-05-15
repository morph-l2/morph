// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Types} from "../libraries/common/Types.sol";
import {L2StakingBaseTest} from "./base/L2StakingBase.t.sol";

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

    function test_balanceOf_succeeds() public {
        assertEq(morphToken.balanceOf(multisig), 1000000000 ether);
    }

    function test_inflationRate_succeeds() public {
        uint256 count = morphToken.inflationRatesCount();
        assertEq(morphToken.epochInflationRates(count - 1).rate, 1596535874529);
    }

    function test_inflationMintedEpochs_succeeds() public {
        assertEq(morphToken.inflationMintedEpochs(), 0);
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
}
