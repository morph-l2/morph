// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Test} from "forge-std/Test.sol";
import {MorphToken} from "../../L2/MorphToken.sol";

contract MorphTokenTest is Test {
    MorphToken morphToken;

    function setUp() public {
        morphToken = new MorphToken();
        morphToken.initialize("Morph", "MPH", address(10), 6, 1000000000e18);
    }

    function test_name() public {
        assertEq(morphToken.name(), "Morph");
    }

    function test_symbol() public {
        assertEq(morphToken.symbol(), "MPH");
    }

    function test_decimals() public {
        assertEq(morphToken.decimals(), 18);
    }

    function test_calculateOneDayRateOfInflation() public {
        assertEq(morphToken.calculateOneDayRateOfInflation(), 290411);
    }
}