// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MorphToken} from "../../L2/MorphToken.sol";
import {DSTestPlus} from "@rari-capital/solmate/src/test/utils/DSTestPlus.sol";
import {Test} from "forge-std/Test.sol";

contract MorphTokenTest is Test {
    MorphToken morphToken;

    address public alice = address(1);
    address public distribute = address(10);
    function setUp() public {
        morphToken = new MorphToken();
        morphToken.initialize("Morph", "MPH", distribute, 1000000000e18, 6, block.timestamp);
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


//contract MorphTokenTest_mint is DSTestPlus {
//    MorphToken morphToken;
//
//    uint256 beginTime = block.timestamp;
//    uint256 rate = 6;
//    uint256 totalSupply = 1000000000e18;
//
//    address public alice = address(1);
//    function setUp() public {
//        morphToken = new MorphToken();
//        morphToken.initialize("Morph", "MPH", address(this), rate, totalSupply, beginTime);
//        emit log_uint(block.timestamp);
//    }
//
//    function test_mint_withOtherAddress() public {
//        hevm.prank(alice);
//        hevm.expectRevert("only distribute contract can call");
//        morphToken.mint(address(0));
//    }
//
//
//    function test_mint_lessThanOneDay() public {
//        hevm.warp(beginTime + 86300);
//        emit log_uint(block.timestamp);
//        hevm.expectRevert("only mint once a day");
//        morphToken.mint(address(0));
//    }
//
//    function test_mint_moreThanOneDay() public {
//        hevm.warp(beginTime + 86500);
//        emit log_uint(block.timestamp);
//        morphToken.mint(address(0));
//        // _totalSupply * 1596535874529 / 1e16
//        totalSupply += totalSupply * 1596535874529 / 1e16;
//        emit log_uint(totalSupply);
//        assertEq(morphToken.totalSupply(), totalSupply);
//    }
//
//    function test_mint_moreThanTenDay() public {
//        hevm.warp(beginTime + 86400 * 10 + 100);
//        emit log_uint(block.timestamp);
//        //hevm.expectRevert("only mint once a day");
//        morphToken.mint(address(0));
//
//        for (uint256 i = 0; i < 10; i++) {
//            totalSupply += totalSupply * 1596535874529 / 1e16;
//        }
//        emit log_uint(totalSupply);
//        assertEq(morphToken.totalSupply(), totalSupply);
//    }
////
//    function test_mint_equalToOneYear() public {
//        hevm.warp(beginTime + 86400 * 365);
//        emit log_uint(block.timestamp);
//        //hevm.expectRevert("only mint once a day");
//        morphToken.mint(address(0));
//
//        for (uint256 i = 0; i < 365; i++) {
//            totalSupply += totalSupply * 1596535874529 / 1e16;
//        }
//        emit log_uint(totalSupply);
//        assertEq(morphToken.totalSupply(), totalSupply);
//    }
////
////    function test_mint_moreThanOneYear() public {
////        hevm.warp(beginTime + 31536000 + 86400 * 10);
////        emit log_uint(block.timestamp);
////        //hevm.expectRevert("only mint once a day");
////        morphToken.mint(address(0));
////        uint256 add = totalSupply * rate / 100 / 365;
////        emit log_uint(add);
////        uint256 nextBase = totalSupply + add * 365;
////        emit log_uint(nextBase);
////        uint256 nextAdd = nextBase * rate / 100 / 365;
////        emit log_uint(nextAdd);
////        assertEq(morphToken.totalSupply(), totalSupply + add * 365 + nextAdd * 10);
////    }
////
////    function test_mint_toOther() public {
////        uint256 oneDayIncrement = totalSupply * rate / 100 / 365;
////        hevm.warp(beginTime + 86500);
////        morphToken.mint(alice);
////        uint256 oneDayTotal = totalSupply + oneDayIncrement;
////        assertEq(morphToken.balanceOf(alice), oneDayIncrement);
////        assertEq(morphToken.totalSupply(), oneDayTotal);
////    }
////
////    function test_mint() public {
////        uint256 oneDayIncrement = totalSupply * rate / 100 / 365;
////        hevm.warp(beginTime + 86500);
////        morphToken.mint(address(0));
////        uint256 oneDayTotal = totalSupply + oneDayIncrement;
////        assertEq(morphToken.totalSupply(), oneDayTotal);
////
////        hevm.warp(beginTime + 86500 + 17 * 86500);
////        morphToken.mint(address(0));
////        uint256 seventeenDayTotal = oneDayTotal + 17 * oneDayIncrement;
////        assertEq(morphToken.totalSupply(), seventeenDayTotal);
////
////        hevm.warp(beginTime + 86500 + 17 * 86500 + 365 * 86500);
////        morphToken.mint(address(0));
////
////        uint256 base = seventeenDayTotal + 347 * oneDayIncrement;
////        uint256 total = base + (base * rate / 100 / 365) * 18;
////        assertEq(morphToken.totalSupply(), total);
////    }
////
////
////    function test_mint_setRate() public {
////        uint256 oneDayIncrement = totalSupply * rate / 100 / 365;
////        hevm.warp(beginTime + 86500);
////        morphToken.mint(address(0));
////        uint256 oneDayTotal = totalSupply + oneDayIncrement;
////        assertEq(morphToken.totalSupply(), oneDayTotal);
////
////        assertEq(morphToken.rate(), rate);
////        morphToken.setPostRate(7);
////
////        hevm.warp(beginTime + 86500 + 86500);
////        morphToken.mint(address(0));
////        uint256 twoDayTotal = oneDayTotal + oneDayIncrement;
////        assertEq(morphToken.totalSupply(), twoDayTotal);
////
////        assertEq(morphToken.rate(), 7);
////
////        hevm.warp(beginTime + 86500 + 86500 + 86500);
////        morphToken.mint(address(0));
////        uint256 newOneDayIncrement = twoDayTotal * 7 / 100 / 365;
////        uint256 threeDayTotal = twoDayTotal + newOneDayIncrement;
////        assertEq(morphToken.totalSupply(), threeDayTotal);
////    }
//}