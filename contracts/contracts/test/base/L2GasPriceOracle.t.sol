// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/* Testing utilities */
import {CommonTest} from "./CommonTest.t.sol";
import {L1MessageBaseTest} from "./L1MessageBase.t.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";

contract L2GasPriceOracleTest is L1MessageBaseTest {
    /// @notice The intrinsic gas for transaction.
    uint256 INTRINSIC_GAS_TX = 21000;
    /// @notice The appropriate intrinsic gas for each byte.
    uint256 APPROPRIATE_INTRINSIC_GAS_PER_BYTE = 16;

    L1MessageQueueWithGasPriceOracle l2GasPriceOracle;

    function setUp() public virtual override {
        super.setUp();
        l2GasPriceOracle = l1MessageQueueWithGasPriceOracle;
    }

    function testCalculateIntrinsicGasFee() external {
        hevm.startPrank(multisig);
        bytes memory _calldata = hex"00";
        uint256 intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(
            _calldata
        );
        assertEq(
            intrinsicGasFee,
            INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE
        );
        _calldata = hex"001122";
        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(_calldata);
        assertEq(
            intrinsicGasFee,
            INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE
        );

        _calldata = hex"0011220033";
        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(_calldata);
        assertEq(
            intrinsicGasFee,
            INTRINSIC_GAS_TX +
                _calldata.length *
                APPROPRIATE_INTRINSIC_GAS_PER_BYTE
        );
        hevm.stopPrank();
    }
}
