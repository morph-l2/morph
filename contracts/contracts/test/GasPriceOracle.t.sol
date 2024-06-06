// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";

import {CommonTest} from "./base/CommonTest.t.sol";
import {GasPriceOracle} from "../l2/system/GasPriceOracle.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";

import {TransparentUpgradeableProxy, ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract GasPriceOracleTest is Test {
    GasPriceOracle internal gasPriceOracle;

    address internal multisig = 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d;
    address internal alice = address(500);
    address internal bob = address(200);

    function setUp() public virtual {
        gasPriceOracle = new GasPriceOracle(multisig);
    }

    /**
     * @notice upgrade to new implementation
     */
    function test_upgradeToNewImplementation_succeeds() public {
        string memory MORPH_HOLESKY_RPC_URL = vm.envString("HOLESKY_RPC_URL");
        bytes memory testUrl = bytes(MORPH_HOLESKY_RPC_URL);
        if (testUrl.length != 0) {
            uint256 morphTestnetFork = vm.createFork(MORPH_HOLESKY_RPC_URL);
            vm.selectFork(morphTestnetFork);

            // owner
            (, bytes memory data00) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0x8da5cb5b));
            address owner00 = abi.decode(data00, (address));
            assertEq(owner00, multisig);

            // overhead
            (, bytes memory data01) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0x0c18c162));
            uint256 overhead01 = abi.decode(data01, (uint256));

            // scalar
            (, bytes memory data02) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0xf45e65d8));
            uint256 scalar02 = abi.decode(data02, (uint256));

            address[] memory allowList = new address[](1);
            allowList[0] = bob;

            bool[] memory allowed = new bool[](1);
            allowed[0] = true;

            bytes memory setAllowedData = abi.encodeWithSignature("setAllowList(address[],bool[])", allowList, allowed);
            vm.prank(multisig);
            (bool succeed, ) = Predeploys.GAS_PRICE_ORACLE.call(setAllowedData);
            assertTrue(succeed);

            _upgradeToNewImplementation(owner00, overhead01, scalar02);
        }
    }

    function _upgradeToNewImplementation(address owner00, uint256 overhead01, uint256 scalar02) internal {
        ITransparentUpgradeableProxy gasPriceOracleProxy = ITransparentUpgradeableProxy(Predeploys.GAS_PRICE_ORACLE);
        // upgrade
        GasPriceOracle newGasPriceOracleImpl = new GasPriceOracle(address(1));
        vm.prank(Predeploys.PROXY_ADMIN);
        gasPriceOracleProxy.upgradeTo(address(newGasPriceOracleImpl));

        // owner
        (, bytes memory data10) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0x8da5cb5b));
        address owner10 = abi.decode(data10, (address));
        assertEq(owner10, multisig);

        // overhead
        (, bytes memory data11) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0x0c18c162));
        uint256 overhead11 = abi.decode(data11, (uint256));

        // scalar
        (, bytes memory data12) = Predeploys.GAS_PRICE_ORACLE.call(abi.encodeWithSelector(0xf45e65d8));
        uint256 scalar12 = abi.decode(data12, (uint256));

        assertEq(owner00, owner10);
        assertEq(overhead01, overhead11);
        assertEq(scalar02, scalar12);

        GasPriceOracle _gasPriceOracle = GasPriceOracle(Predeploys.GAS_PRICE_ORACLE);
        assertEq(_gasPriceOracle.owner(), owner00);

        vm.prank(multisig);
        _gasPriceOracle.setCommitScalar(1000);
        assertEq(_gasPriceOracle.commitScalar(), 1000);

        address newOp = address(300);
        address[] memory allowList = new address[](1);
        allowList[0] = newOp;

        bool[] memory allowed = new bool[](1);
        allowed[0] = true;
        vm.prank(multisig);
        _gasPriceOracle.setAllowList(allowList, allowed);

        assertTrue(_gasPriceOracle.isAllowed(bob));
        assertTrue(_gasPriceOracle.isAllowed(newOp));
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

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(alice);
        gasPriceOracle.setAllowList(allowed, vals);

        vm.expectRevert(bytes4(keccak256("ErrDifferentLength()")));
        vm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);

        allowed = new address[](1);
        vals = new bool[](1);

        allowed[0] = bob;
        vals[0] = true;

        vm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);

        assertTrue(gasPriceOracle.isAllowed(bob));

        vals[0] = false;
        vm.prank(multisig);
        gasPriceOracle.setAllowList(allowed, vals);
        assertFalse(gasPriceOracle.isAllowed(bob));
    }

    /**
     * @notice setAllowListEnabled
     */
    function test_setAllowListEnabled_works() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(alice);
        gasPriceOracle.setAllowListEnabled(false);

        bool val = gasPriceOracle.allowListEnabled();
        vm.expectRevert(bytes4(keccak256("ErrSettintSameValue()")));
        vm.prank(multisig);
        gasPriceOracle.setAllowListEnabled(val);

        val = !gasPriceOracle.allowListEnabled();
        vm.prank(multisig);
        gasPriceOracle.setAllowListEnabled(val);
        assertEq(val, gasPriceOracle.allowListEnabled());
    }

    /**
     * @notice setL1BaseFee
     */
    function test_setL1BaseFee_works() public {
        vm.expectRevert(bytes4(keccak256("ErrCallerNotAllowed()")));
        vm.prank(alice);
        gasPriceOracle.setL1BaseFee(100);

        vm.prank(multisig);
        gasPriceOracle.setL1BaseFee(100);
        assertEq(100, gasPriceOracle.l1BaseFee());
    }

    /**
     * @notice setOverhead
     */
    function test_setOverhead_works() public {
        vm.expectRevert(bytes4(keccak256("ErrCallerNotAllowed()")));
        vm.prank(alice);
        gasPriceOracle.setOverhead(100);

        vm.prank(multisig);
        vm.expectRevert(bytes4(keccak256("ErrExceedMaxOverhead()")));
        gasPriceOracle.setOverhead(30000000 / 16 + 1);

        vm.prank(multisig);
        gasPriceOracle.setOverhead(1);
        assertEq(1, gasPriceOracle.overhead());
    }

    /**
     * @notice setScalar
     */
    function test_setScalar_works() public {
        vm.expectRevert(bytes4(keccak256("ErrCallerNotAllowed()")));
        vm.prank(alice);
        gasPriceOracle.setScalar(100);

        vm.prank(multisig);
        vm.expectRevert(bytes4(keccak256("ErrExceedMaxScalar()")));
        gasPriceOracle.setScalar(1000 * 1e9 + 1);

        vm.prank(multisig);
        gasPriceOracle.setScalar(1);
        assertEq(1, gasPriceOracle.scalar());
    }

    // /**
    //  * @notice getL1GasUsed
    //  */
    // function test_getL1GasUsed_works() public {
    //     uint256 overhead = gasPriceOracle.overhead();

    //     bytes memory data = hex"0000";
    //     uint256 expected = overhead + 4 * 2 + (16 * 4);
    //     uint256 gasUsed = gasPriceOracle.getL1GasUsed(data);
    //     assertEq(gasUsed, expected);

    //     data = hex"0001";
    //     expected = overhead + 4 + 16 + (16 * 4);
    //     gasUsed = gasPriceOracle.getL1GasUsed(data);
    //     assertEq(gasUsed, expected);

    //     data = hex"0101";
    //     expected = overhead + 16 + 16 + (16 * 4);
    //     gasUsed = gasPriceOracle.getL1GasUsed(data);
    //     assertEq(gasUsed, expected);
    // }

    // /**
    //  * @notice getL1Fee
    //  */
    // function test_getL1Fee_works() public {
    //     bytes memory data = hex"0101";
    //     uint256 expected = gasPriceOracle.overhead() + 16 + 16 + (16 * 4);
    //     uint256 gasUsed = gasPriceOracle.getL1GasUsed(data);
    //     assertEq(gasUsed, expected);

    //     vm.prank(multisig);
    //     gasPriceOracle.setL1BaseFee(100);
    //     assertEq(100, gasPriceOracle.l1BaseFee());

    //     vm.prank(multisig);
    //     gasPriceOracle.setScalar(5e9);
    //     assertEq(5e9, gasPriceOracle.scalar());

    //     uint256 l1Fee = gasPriceOracle.getL1Fee(data);
    //     expected = (gasPriceOracle.l1BaseFee() * gasUsed * gasPriceOracle.scalar()) / 1e9;
    //     assertEq(l1Fee, expected);
    // }
}
