// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {CommonTest} from "./base/CommonTest.t.sol";
import {GasPriceOracle} from "../l2/system/GasPriceOracle.sol";

contract GasPriceOracleTest is CommonTest {
    GasPriceOracle internal gasPriceOracle;

    function setUp() public virtual override {
        super.setUp();

        gasPriceOracle = new GasPriceOracle(multisig);
    }

    /**
     * @notice owner
     */
    function test_owner_succeeds() public {
        assertEq(gasPriceOracle.owner(), multisig);
    }

    /**
     * @notice setAllowList
     */
    function test_allowList_works() public {
        address[] memory allowed = new address[](2);
        bool[] memory vals = new bool[](3);

        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(alice);
        gasPriceOracle.setAllowList(allowed, vals);

        hevm.expectRevert("INVALID_INPUT");
        hevm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);

        allowed = new address[](1);
        vals = new bool[](1);

        allowed[0] = bob;
        vals[0] = true;

        hevm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);

        assertTrue(gasPriceOracle.isAllowed(bob));

        vals[0] = false;
        hevm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);
        assertFalse(gasPriceOracle.isAllowed(bob));
    }

    /**
     * @notice setAllowListEnabled
     */
    function test_setAllowListEnabled_works() public {
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(alice);
        gasPriceOracle.setAllowListEnabled(false);

        bool val = gasPriceOracle.allowListEnabled();
        hevm.expectRevert("ALREADY_SET");
        hevm.prank(multisig);
        gasPriceOracle.setAllowListEnabled(val);

        val = !gasPriceOracle.allowListEnabled();
        hevm.prank(multisig);
        gasPriceOracle.setAllowListEnabled(val);
        assertBoolEq(val, gasPriceOracle.allowListEnabled());
    }

    /**
     * @notice setL1BaseFee
     */
    function test_setL1BaseFee_works() public {
        hevm.expectRevert("not allowed");
        hevm.prank(alice);
        gasPriceOracle.setL1BaseFee(100);

        hevm.prank(multisig);
        gasPriceOracle.setL1BaseFee(100);
        assertEq(100, gasPriceOracle.l1BaseFee());
    }

    /**
     * @notice setOverhead
     */
    function test_setOverhead_works() public {
        hevm.expectRevert("not allowed");
        hevm.prank(alice);
        gasPriceOracle.setOverhead(100);

        hevm.prank(multisig);
        hevm.expectRevert("exceed maximum overhead");
        gasPriceOracle.setOverhead(30000000 / 16 + 1);

        hevm.prank(multisig);
        gasPriceOracle.setOverhead(1);
        assertEq(1, gasPriceOracle.overhead());
    }

    /**
     * @notice setScalar
     */
    function test_setScalar_works() public {
        hevm.expectRevert("not allowed");
        hevm.prank(alice);
        gasPriceOracle.setScalar(100);

        hevm.prank(multisig);
        hevm.expectRevert("exceed maximum scale");
        gasPriceOracle.setScalar(1000 * 1e9 + 1);

        hevm.prank(multisig);
        gasPriceOracle.setScalar(1);
        assertEq(1, gasPriceOracle.scalar());
    }

    /**
     * @notice getL1GasUsed
     */
    function test_getL1GasUsed_works() public {
        uint256 overhead = gasPriceOracle.overhead();

        bytes memory data = hex"0000";
        uint256 expected = overhead + 4 * 2 + (16 * 4);
        uint256 gasUsed = gasPriceOracle.getL1GasUsed(data);
        assertEq(gasUsed, expected);

        data = hex"0001";
        expected = overhead + 4 + 16 + (16 * 4);
        gasUsed = gasPriceOracle.getL1GasUsed(data);
        assertEq(gasUsed, expected);

        data = hex"0101";
        expected = overhead + 16 + 16 + (16 * 4);
        gasUsed = gasPriceOracle.getL1GasUsed(data);
        assertEq(gasUsed, expected);
    }

    /**
     * @notice getL1Fee
     */
    function test_getL1Fee_works() public {
        bytes memory data = hex"0101";
        uint256 expected = gasPriceOracle.overhead() + 16 + 16 + (16 * 4);
        uint256 gasUsed = gasPriceOracle.getL1GasUsed(data);
        assertEq(gasUsed, expected);

        hevm.prank(multisig);
        gasPriceOracle.setL1BaseFee(100);
        assertEq(100, gasPriceOracle.l1BaseFee());

        hevm.prank(multisig);
        gasPriceOracle.setScalar(5e9);
        assertEq(5e9, gasPriceOracle.scalar());

        uint256 l1Fee = gasPriceOracle.getL1Fee(data);
        expected = (gasPriceOracle.l1BaseFee() * gasUsed * gasPriceOracle.scalar()) / 1e9;
        assertEq(l1Fee, expected);
    }
}
