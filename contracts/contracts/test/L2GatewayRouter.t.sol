// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";

import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {L2GatewayRouter} from "../l2/gateways/L2GatewayRouter.sol";
import {IL2GatewayRouter} from "../l2/gateways/IL2GatewayRouter.sol";
import {L2CrossDomainMessenger} from "../l2/L2CrossDomainMessenger.sol";

contract L2GatewayRouterTest is L2GatewayBaseTest {
    L2GatewayRouter private router;
    L2CrossDomainMessenger private l2Messenger;

    address private feeVault;
    address private l1Messenger;

    MockERC20 private l1Token;

    function setUp() public override {
        super.setUp();
        router = l2GatewayRouter;
        l2Messenger = l2CrossDomainMessenger;
        feeVault = l2FeeVault;
        l1Messenger = address(NON_ZERO_ADDRESS);

        hevm.prank(multisig);
        router.transferOwnership(address(this));
        // Deploy tokens
        l1Token = new MockERC20("Mock", "M", 18);
    }

    function test_ownership_succeeds() public {
        assertEq(address(this), router.owner());
    }

    function test_initialized_reInit_reverts() public {
        assertEq(address(l2ETHGateway), router.ethGateway());
        assertEq(address(l2StandardERC20Gateway), router.defaultERC20Gateway());
        assertEq(address(l2StandardERC20Gateway), router.getERC20Gateway(address(l1Token)));

        hevm.expectRevert("Initializable: contract is already initialized");
        router.initialize(address(l2ETHGateway), address(l2StandardERC20Gateway));
    }

    function test_setDefaultERC20Gateway_succeeds() public {
        router.setDefaultERC20Gateway(address(0));

        // set by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        router.setDefaultERC20Gateway(address(l2StandardERC20Gateway));
        hevm.stopPrank();

        // set by owner, should succeed
        hevm.expectEmit(true, true, false, true);
        emit IL2GatewayRouter.SetDefaultERC20Gateway(address(0), address(l2StandardERC20Gateway));

        assertEq(address(0), router.getERC20Gateway(address(l1Token)));
        assertEq(address(0), router.defaultERC20Gateway());
        router.setDefaultERC20Gateway(address(l2StandardERC20Gateway));
        assertEq(address(l2StandardERC20Gateway), router.getERC20Gateway(address(l1Token)));
        assertEq(address(l2StandardERC20Gateway), router.defaultERC20Gateway());
    }

    function test_setERC20Gateway_succeeds() public {
        router.setDefaultERC20Gateway(address(0));

        // length mismatch, should revert
        address[] memory empty = new address[](0);
        address[] memory single = new address[](1);
        hevm.expectRevert("length mismatch");
        router.setERC20Gateway(empty, single);
        hevm.expectRevert("length mismatch");
        router.setERC20Gateway(single, empty);

        // set by owner, should succeed
        address[] memory _tokens = new address[](1);
        address[] memory _gateways = new address[](1);
        _tokens[0] = address(l1Token);
        _gateways[0] = address(l2StandardERC20Gateway);

        hevm.expectEmit(true, true, true, true);
        emit IL2GatewayRouter.SetERC20Gateway(address(l1Token), address(0), address(l2StandardERC20Gateway));

        assertEq(address(0), router.getERC20Gateway(address(l1Token)));
        router.setERC20Gateway(_tokens, _gateways);
        assertEq(address(l2StandardERC20Gateway), router.getERC20Gateway(address(l1Token)));
    }

    function test_finalizeDepositERC20_neverCalled_reverts() public {
        hevm.expectRevert("should never be called");
        router.finalizeDepositERC20(address(0), address(0), address(0), address(0), 0, "");
    }

    function test_finalizeDepositETH_neverCalled_reverts() public {
        hevm.expectRevert("should never be called");
        router.finalizeDepositETH(address(0), address(0), 0, "");
    }
}
