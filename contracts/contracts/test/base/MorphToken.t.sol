// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MorphToken} from "../../L2/MorphToken.sol";
import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import {Test} from "forge-std/Test.sol";

contract MorphTokenTest is Test {
    MorphToken private morphToken;

    address public alice = address(1);
    address public distribute = address(10);
    function setUp() public {
        morphToken = new MorphToken();
        morphToken.initialize("Morph", "MPH", distribute, 1000000000e18, 1596535874529, 86400);
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

    function test_totalSupply() public {
        assertEq(morphToken.totalSupply(), 1000000000e18);
    }

    function test_balanceOf() public {
        assertEq(morphToken.balanceOf(address(this)), 1000000000e18);
    }

    function test_transfer() public {
        bool s = morphToken.transfer(alice, 10000000e18);
        assertEq(s, true);
        assertEq(morphToken.balanceOf(alice), 10000000e18);
    }
}


contract MorphTokenTest_mint is DSTestPlus {
    MorphToken private morphToken;

    uint256 private beginTime = 86400;
    uint256 private rate = 1596535874529;
    uint256 private totalSupply = 1000000000e18;

    address public alice = address(1);
    function setUp() public {
        morphToken = new MorphToken();
        hevm.expectRevert("beginTime must be the start of the day");
        morphToken.initialize("Morph", "MPH", address(this), totalSupply, rate, 86300);
        morphToken.initialize("Morph", "MPH", address(this), totalSupply, rate, beginTime);
        emit log_uint(block.timestamp);
    }

    function test_mint_withOtherAddress() public {
        hevm.prank(alice);
        hevm.expectRevert("only distribute contract can call");
        morphToken.mint();
    }

    function test_mint_notBegin() public {
        hevm.expectRevert("the mint function is not enabled");
        morphToken.mint();
    }

    function test_mint_lessThanOneDay() public {
        hevm.warp(beginTime + 86300);
        emit log_uint(block.timestamp);
        hevm.expectRevert("only mint once a day");
        morphToken.mint();
    }

    function test_mint_moreThanOneDay() public {
        hevm.warp(beginTime + 86500);
        emit log_uint(block.timestamp);
        morphToken.mint();
        totalSupply += totalSupply * rate / 1e16;
        emit log_uint(totalSupply);
        assertEq(morphToken.totalSupply(), totalSupply);
    }

    function test_mint_moreThanTenDay() public {
        hevm.warp(beginTime + 86400 * 10 + 100);
        emit log_uint(block.timestamp);
        morphToken.mint();

        for (uint256 i = 0; i < 10; i++) {
            totalSupply += totalSupply * rate / 1e16;
        }
        emit log_uint(totalSupply);
        assertEq(morphToken.totalSupply(), totalSupply);
    }

    function test_setRate_pastTime() public {
        uint256 rate1 = 1696535874529;
        uint256 beginTime1 = block.timestamp - 1;
        hevm.expectRevert("beginTime must be more than the current time");
        morphToken.setRate(rate1, beginTime1);
    }

    function test_setRate_notMoreThanCurrentBeginTime() public {
        uint256 rate1 = 1696535874529;
        uint256 beginTime1 = beginTime + 86400;
        hevm.expectRevert("beginTime must be two weeks after the current validity period");
        morphToken.setRate(rate1, beginTime1);
    }

    function test_setRate_notMoreThanPendingBeginTime() public {
        uint256 rate1 = 1696535874529;
        uint256 beginTime1 = beginTime + 1209600;
        //hevm.expectRevert("beginTime must be two weeks after the current validity period");
        morphToken.setRate(rate1, beginTime1);

        uint256 rate2 = 1796535874529;
        uint256 beginTime2 = beginTime1 + 1209500;
        hevm.expectRevert("beginTime must be more than two weeks after the last exchange rate takes effect");
        morphToken.setRate(rate2, beginTime2);
    }

    function test_mint_setRate() public {

        uint256 rate1 = 1696535874529;
        uint256 beginTime1 = beginTime + 86400 * 20 + 86300;
        morphToken.setRate(rate1, beginTime1);

        uint256 rate2 = 1796535874529;
        uint256 beginTime2 = beginTime + 86400 * 7 * 6;
        morphToken.setRate(rate2, beginTime2);

        uint256 rate3 = 1896535874529;
        uint256 beginTime3 = beginTime + 86400 * 7 * 9;
        morphToken.setRate(rate3, beginTime3);

        hevm.warp(beginTime + 86400 * 43);
        emit log_uint(block.timestamp);

        for (uint256 i = 0; i < 21; i++) {
            totalSupply += totalSupply * rate / 1e16;
        }

        for (uint256 i = 0; i < 21; i++) {
            totalSupply += totalSupply * rate1 / 1e16;
        }

        for (uint256 i = 0; i < 1; i++) {
            totalSupply += totalSupply * rate1 / 1e16;
        }

        morphToken.mint();
        assertEq(morphToken.totalSupply(), totalSupply);
    }
}