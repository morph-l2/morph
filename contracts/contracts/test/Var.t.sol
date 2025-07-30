// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/* Testing utilities */
import "forge-std/Test.sol";
import "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";

contract VarTest is DSTestPlus {
    uint256 public tmpData = 3;

    function test_updateVar() public {
        tmpData = 5;
    }

    function test_resetVar() public {
        tmpData = 5;
        tmpData = 3;
    }
}
