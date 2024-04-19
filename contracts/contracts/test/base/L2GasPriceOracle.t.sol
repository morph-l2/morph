// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/* Testing utilities */
import {CommonTest} from "./CommonTest.t.sol";
import {L1MessageBaseTest} from "./L1MessageBase.t.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../L1/rollup/L1MessageQueueWithGasPriceOracle.sol";

contract L2GasPriceOracleTest is L1MessageBaseTest {
    event UpdateL2BaseFee(uint256 oldL2BaseFee, uint256 newL2BaseFee);
    event UpdateWhitelistChecker(
        address indexed _oldWhitelistChecker,
        address indexed _newWhitelistChecker
    );
    event WhitelistStatusChanged(address indexed _account, bool _status);

    /// @notice The intrinsic gas for transaction.
    uint256 INTRINSIC_GAS_TX = 21000;
    /// @notice The appropriate intrinsic gas for each byte.
    uint256 APPROPRIATE_INTRINSIC_GAS_PER_BYTE = 16;

    L1MessageQueueWithGasPriceOracle l2GasPriceOracle;

    function setUp() public virtual override {
        super.setUp();
        l2GasPriceOracle = l1MessageQueueWithGasPriceOracle;
    }

    function testEstimateCrossDomainMessageFee() external {
        hevm.startPrank(multisig);
        uint256 gasLimit = 100;
        l2GasPriceOracle.setL2BaseFee(1);

        // estimate fee without sender in whitelist
        uint256 fee = l2GasPriceOracle.estimateCrossDomainMessageFee(
            multisig,
            gasLimit
        );
        assertEq(fee, gasLimit * 1);

        // add address this to whitelist
        address[] memory addList = new address[](1);
        addList[0] = address(multisig);
        hevm.expectEmit(true, true, true, true);
        emit WhitelistStatusChanged(address(multisig), true);
        whitelistChecker.updateWhitelistStatus(addList, true);
        assertTrue(whitelistChecker.isSenderAllowed(address(multisig)));

        // estimate fee with sender in whitelist
        fee = l2GasPriceOracle.estimateCrossDomainMessageFee(
            multisig,
            gasLimit
        );
        assertEq(fee, 0);

        hevm.stopPrank();
    }

    function testUpdateWhitelistChecker() external {
        address testChecker = address(123);
        hevm.startPrank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit UpdateWhitelistChecker(address(whitelistChecker), testChecker);
        l2GasPriceOracle.updateWhitelistChecker(testChecker);
        assertEq(l2GasPriceOracle.whitelistChecker(), testChecker);
        hevm.stopPrank();
    }

    function testUpdateL2BaseFee() external {
        hevm.startPrank(multisig);
        hevm.expectEmit(true, true, true, true);
        emit UpdateL2BaseFee(0, 1);
        l2GasPriceOracle.setL2BaseFee(1);
        assertEq(l2GasPriceOracle.l2BaseFee(), 1);
        hevm.stopPrank();
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
