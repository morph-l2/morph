// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

/* Testing utilities */
import {CommonTest} from "./CommonTest.t.sol";
import {L1MessageBaseTest} from "./L1MessageBase.t.sol";

contract L2GasPriceOracleTest is L1MessageBaseTest {
    function testCalculateIntrinsicGasFee() external {
        hevm.startPrank(multisig);
        uint256 intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(
            hex"00"
        );
        assertEq(intrinsicGasFee, 2);
        uint64 _zeroGas = 5;
        uint64 _nonZeroGas = 10;
        l2GasPriceOracle.setIntrinsicParams(
            20000,
            50000,
            _zeroGas,
            _nonZeroGas
        );

        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(
            hex"001122"
        );
        // 20000 + 1 zero bytes * 5 + 2 nonzero byte * 10 = 20025
        assertEq(intrinsicGasFee, 20025);

        _zeroGas = 50;
        _nonZeroGas = 100;
        l2GasPriceOracle.setIntrinsicParams(
            10000,
            20000,
            _zeroGas,
            _nonZeroGas
        );

        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(
            hex"0011220033"
        );
        // 10000 + 3 nonzero byte * 100 + 2 zero bytes * 50 = 10000 + 300 + 100 = 10400
        assertEq(intrinsicGasFee, 10400);
        hevm.stopPrank();
    }

    function testSetIntrinsicParamsAccess() external {
        hevm.startPrank(address(4));
        hevm.expectRevert("Ownable: caller is not the owner");
        l2GasPriceOracle.setIntrinsicParams(1, 0, 0, 1);
        hevm.stopPrank();
    }
}
