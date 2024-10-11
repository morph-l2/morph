// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

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

    function test_initialize_works() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for the L2GatewayRouter.
        TransparentUpgradeableProxy l2GatewayRouterProxyTempA = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
        );
        // Deploy a new L2GatewayRouter contract.
        L2GatewayRouter l2GatewayRouterImplTempA = new L2GatewayRouter();
        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2GatewayRouterProxyTempA))
            .upgradeToAndCall(
                address(l2GatewayRouterImplTempA),
                abi.encodeCall(
                    L2GatewayRouter.initialize,
                    (address(0), address(0))
                )
        );
        // Cast the proxy address to the L2GatewayRouter contract type to call its methods.
        L2GatewayRouter l2GatewayRouterTempA = L2GatewayRouter(address(l2GatewayRouterProxyTempA));
        hevm.stopPrank();
        // Verify that the ethGateway and defaultERC20Gateway are zero addresses when initialized with zero addresses.
        assertEq(l2GatewayRouterTempA.ethGateway(), address(0));
        assertEq(l2GatewayRouterTempA.defaultERC20Gateway(), address(0));
        hevm.startPrank(multisig);
        // Deploy another proxy contract for the L2GatewayRouter.
        TransparentUpgradeableProxy l2GatewayRouterProxyTempB = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
        );
        // Deploy a new L2GatewayRouter contract.
        L2GatewayRouter l2GatewayRouterImplTempB = new L2GatewayRouter();
        
        // Expect the SetDefaultERC20Gateway event to be emitted successfully.
        hevm.expectEmit(true, true, false, true);
        emit IL2GatewayRouter.SetDefaultERC20Gateway(address(0), address(2));
        
        // Expect the SetETHGateway event to be emitted successfully.
        hevm.expectEmit(true, true, false, true);
        emit IL2GatewayRouter.SetETHGateway(address(0), address(1));
        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2GatewayRouterProxyTempB))
            .upgradeToAndCall(
                address(l2GatewayRouterImplTempB),
                abi.encodeCall(
                    L2GatewayRouter.initialize,
                    (address(1), address(2))
                )
        );
        // Cast the proxy address to the L2GatewayRouter contract type to call its methods.
        L2GatewayRouter l2GatewayRouterTempB = L2GatewayRouter(address(l2GatewayRouterProxyTempB));
        hevm.stopPrank();
        
        // Verify that the ethGateway and defaultERC20Gateway are initialized correctly.
        assertEq(l2GatewayRouterTempB.ethGateway(), address(1));
        assertEq(l2GatewayRouterTempB.defaultERC20Gateway(), address(2));
    }
    function test_setETHGateway_onlyOwner_reverts() public {
        hevm.prank(address(1));
        // Expect revert when the caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        router.setETHGateway(address(1));
    }
    function test_setETHGateway_works() public {
        // Expect the SetETHGateway event to be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2GatewayRouter.SetETHGateway(router.ethGateway(), address(1));
        router.setETHGateway(address(1));
        // Verify that the ethGateway is set successfully.
        assertEq(router.ethGateway(), address(1));
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

    function test_setERC20Gateway_onlyOwner_reverts() public {
        address[] memory single = new address[](1);
        hevm.startPrank(address(1));
        // Expect revert when the caller is not the owner.
        hevm.expectRevert("Ownable: caller is not the owner");
        router.setERC20Gateway(single, single);
        hevm.stopPrank();
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

    function test_requestERC20_onlyInContext_reverts() public {
        // Expect revert when _msgSender() != gatewayInContext.
        hevm.expectRevert("Only in deposit context");
        router.requestERC20(address(1), address(1), 10);
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
