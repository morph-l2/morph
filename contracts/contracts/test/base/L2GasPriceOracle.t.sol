// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/* Testing utilities */
import {IL1MessageQueueWithGasPriceOracle} from "../../l1/rollup/IL1MessageQueueWithGasPriceOracle.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {IWhitelist} from "../../libraries/common/IWhitelist.sol";
import {L1MessageBaseTest} from "./L1MessageBase.t.sol";

contract L2GasPriceOracleTest is L1MessageBaseTest {
    /// @notice The intrinsic gas for transaction.
    uint256 public constant INTRINSIC_GAS_TX = 21000;
    /// @notice The appropriate intrinsic gas for each byte.
    uint256 public constant APPROPRIATE_INTRINSIC_GAS_PER_BYTE = 16;

    L1MessageQueueWithGasPriceOracle public l2GasPriceOracle;

    function setUp() public virtual override {
        super.setUp();
        l2GasPriceOracle = l1MessageQueueWithGasPriceOracle;
    }

    function test_estimateCrossDomainMessageFee_succeeds() external {
        hevm.startPrank(multisig);
        uint256 gasLimit = 100;
        l2GasPriceOracle.setL2BaseFee(1);

        // estimate fee without sender in whitelist
        uint256 fee = l2GasPriceOracle.estimateCrossDomainMessageFee(multisig, gasLimit);
        assertEq(fee, gasLimit * 1);

        // add address this to whitelist
        address[] memory addList = new address[](1);
        addList[0] = address(multisig);
        hevm.expectEmit(true, true, true, true);
        emit IWhitelist.WhitelistStatusChanged(address(multisig), true);
        whitelistChecker.updateWhitelistStatus(addList, true);
        assertTrue(whitelistChecker.isSenderAllowed(address(multisig)));

        // estimate fee with sender in whitelist
        fee = l2GasPriceOracle.estimateCrossDomainMessageFee(multisig, gasLimit);
        assertEq(fee, 0);

        hevm.stopPrank();
    }

    function test_updateWhitelistChecker_succeeds() external {
        address testChecker = address(123);
        hevm.startPrank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueueWithGasPriceOracle.UpdateWhitelistChecker(address(whitelistChecker), testChecker);
        l2GasPriceOracle.updateWhitelistChecker(testChecker);
        assertEq(l2GasPriceOracle.whitelistChecker(), testChecker);
        hevm.stopPrank();
    }

    function test_updateL2BaseFee_succeeds() external {
        hevm.startPrank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit IL1MessageQueueWithGasPriceOracle.UpdateL2BaseFee(0, 1);
        l2GasPriceOracle.setL2BaseFee(1);
        assertEq(l2GasPriceOracle.l2BaseFee(), 1);
        hevm.stopPrank();
    }

    function test_calculateIntrinsicGasFee_succeeds() external {
        hevm.startPrank(multisig);
        bytes memory _calldata = "0x00";
        uint256 intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(_calldata);
        assertEq(intrinsicGasFee, INTRINSIC_GAS_TX + _calldata.length * APPROPRIATE_INTRINSIC_GAS_PER_BYTE);
        _calldata = "0x001122";
        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(_calldata);
        assertEq(intrinsicGasFee, INTRINSIC_GAS_TX + _calldata.length * APPROPRIATE_INTRINSIC_GAS_PER_BYTE);

        _calldata = "0x0011220033";
        intrinsicGasFee = l2GasPriceOracle.calculateIntrinsicGasFee(_calldata);
        assertEq(intrinsicGasFee, INTRINSIC_GAS_TX + _calldata.length * APPROPRIATE_INTRINSIC_GAS_PER_BYTE);
        hevm.stopPrank();
    }
}
