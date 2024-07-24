// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";

import {IL1GatewayRouter} from "../l1/gateways/IL1GatewayRouter.sol";
import {L1GatewayRouter} from "../l1/gateways/L1GatewayRouter.sol";
import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {TransferReentrantToken} from "../mock/tokens/TransferReentrantToken.sol";

contract L1GatewayRouterTest is L1GatewayBaseTest {
    MockERC20 private l1Token;

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // Deploy tokens
        l1Token = new MockERC20("Mock", "M", 18);
        l1GatewayRouter.transferOwnership(address(this));
        hevm.stopPrank();
    }

    function test_initialize_reverts() external {
        // Verify that initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        l1GatewayRouter.initialize(address(1), address(1));
    }

    function test_initialize_works() public {
        hevm.startPrank(multisig);

        // Deploy a proxy contract for the L1GatewayRouter.
        TransparentUpgradeableProxy l1GatewayRouterProxyTempA = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1GatewayRouter contract.
        L1GatewayRouter l1GatewayRouterImplTempA = new L1GatewayRouter();

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1GatewayRouterProxyTempA)).upgradeToAndCall(
            address(l1GatewayRouterImplTempA),
            abi.encodeCall(L1GatewayRouter.initialize, (address(0), address(0)))
        );

        // Cast the proxy address to the L1GatewayRouter contract type to call its methods.
        L1GatewayRouter l1GatewayRouterTempA = L1GatewayRouter(address(l1GatewayRouterProxyTempA));
        hevm.stopPrank();

        // Verify that the ethGateway and defaultERC20Gateway are zero addresses when initialized with zero addresses.
        assertEq(l1GatewayRouterTempA.ethGateway(), address(0));
        assertEq(l1GatewayRouterTempA.defaultERC20Gateway(), address(0));

        hevm.startPrank(multisig);

        // Deploy another proxy contract for the L1GatewayRouter.
        TransparentUpgradeableProxy l1GatewayRouterProxyTempB = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // Deploy a new L1GatewayRouter contract.
        L1GatewayRouter l1GatewayRouterImplTempB = new L1GatewayRouter();

        // Verify the SetDefaultERC20Gateway event is emitted successfully.
        hevm.expectEmit(true, true, false, true);
        emit IL1GatewayRouter.SetDefaultERC20Gateway(address(0), address(2));

        // Verify the SetETHGateway event is emitted successfully.
        hevm.expectEmit(true, true, false, true);
        emit IL1GatewayRouter.SetETHGateway(address(0), address(1));

        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l1GatewayRouterProxyTempB)).upgradeToAndCall(
            address(l1GatewayRouterImplTempB),
            abi.encodeCall(L1GatewayRouter.initialize, (address(1), address(2)))
        );

        // Cast the proxy address to the L1GatewayRouter contract type to call its methods.
        L1GatewayRouter l1GatewayRouterTempB = L1GatewayRouter(address(l1GatewayRouterProxyTempB));
        hevm.stopPrank();

        // Verify that the ethGateway and defaultERC20Gateway are correctly initialized.
        assertEq(l1GatewayRouterTempB.ethGateway(), address(1));
        assertEq(l1GatewayRouterTempB.defaultERC20Gateway(), address(2));
    }

    function test_getL2ERC20Address_works() public {
        // Verify that getL2ERC20Address returns zero address when setDefaultERC20Gateway is set to zero address.
        l1GatewayRouter.setDefaultERC20Gateway(address(0));
        address l2Address = l1GatewayRouter.getL2ERC20Address(address(0));
        assertEq(l2Address, address(0));
    }

    function test_setETHGateway_works() public {
        // Verify that only the owner can set the ETH gateway.
        hevm.expectRevert("Ownable: caller is not the owner");
        hevm.prank(address(1));
        l1GatewayRouter.setETHGateway(address(l1StandardERC20Gateway));

        // Expect SetETHGateway event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL1GatewayRouter.SetETHGateway(l1GatewayRouter.ethGateway(), address(1));
        l1GatewayRouter.setETHGateway(address(1));
    }

    function test_ownership_succeeds() public {
        assertEq(address(this), l1GatewayRouter.owner());
    }

    function test_setDefaultERC20Gateway_works() public {
        l1GatewayRouter.setDefaultERC20Gateway(address(0));

        // set by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1GatewayRouter.setDefaultERC20Gateway(address(l1StandardERC20Gateway));
        hevm.stopPrank();

        // set by owner, should succeed
        hevm.expectEmit(true, true, false, true);
        emit IL1GatewayRouter.SetDefaultERC20Gateway(address(0), address(l1StandardERC20Gateway));

        assertEq(address(0), l1GatewayRouter.getERC20Gateway(address(l1Token)));
        assertEq(address(0), l1GatewayRouter.defaultERC20Gateway());
        l1GatewayRouter.setDefaultERC20Gateway(address(l1StandardERC20Gateway));
        assertEq(address(l1StandardERC20Gateway), l1GatewayRouter.getERC20Gateway(address(l1Token)));
        assertEq(address(l1StandardERC20Gateway), l1GatewayRouter.defaultERC20Gateway());
    }

    function test_setERC20Gateway_works() public {
        l1GatewayRouter.setDefaultERC20Gateway(address(0));

        // length mismatch, should revert
        address[] memory empty = new address[](0);
        address[] memory single = new address[](1);
        hevm.expectRevert("length mismatch");
        l1GatewayRouter.setERC20Gateway(empty, single);
        hevm.expectRevert("length mismatch");
        l1GatewayRouter.setERC20Gateway(single, empty);

        // set by owner, should succeed
        address[] memory _tokens = new address[](1);
        address[] memory _gateways = new address[](1);
        _tokens[0] = address(l1Token);
        _gateways[0] = address(l1StandardERC20Gateway);

        hevm.expectEmit(true, true, true, true);
        emit IL1GatewayRouter.SetERC20Gateway(address(l1Token), address(0), address(l1StandardERC20Gateway));

        assertEq(address(0), l1GatewayRouter.getERC20Gateway(address(l1Token)));
        l1GatewayRouter.setERC20Gateway(_tokens, _gateways);
        assertEq(address(l1StandardERC20Gateway), l1GatewayRouter.getERC20Gateway(address(l1Token)));
    }

    function test_finalizeWithdrawERC20_neverCalled_reverts() public {
        hevm.expectRevert("should never be called");
        l1GatewayRouter.finalizeWithdrawERC20(address(0), address(0), address(0), address(0), 0, "");
    }

    function test_finalizeWithdrawETH_neverCalled_reverts() public {
        hevm.expectRevert("should never be called");
        l1GatewayRouter.finalizeWithdrawETH(address(0), address(0), 0, "");
    }

    function test_requestERC20_context_reverts(address _sender, address _token, uint256 _amount) public {
        hevm.expectRevert("Only in deposit context");
        l1GatewayRouter.requestERC20(_sender, _token, _amount);
    }

    function test_reentrant_context_reverts() public {
        TransferReentrantToken reentrantToken = new TransferReentrantToken("Reentrant", "R", 18);
        reentrantToken.mint(address(this), type(uint128).max);
        reentrantToken.approve(address(l1GatewayRouter), type(uint256).max);

        reentrantToken.setReentrantCall(
            address(l1GatewayRouter),
            0,
            abi.encodeCall(
                l1GatewayRouter.depositERC20AndCall,
                (address(reentrantToken), address(this), 0, new bytes(0), 0)
            ),
            true
        );
        hevm.expectRevert("Only not in context");
        l1GatewayRouter.depositERC20(address(reentrantToken), 1, 0);

        reentrantToken.setReentrantCall(
            address(l1GatewayRouter),
            0,
            abi.encodeCall(
                l1GatewayRouter.depositERC20AndCall,
                (address(reentrantToken), address(this), 0, new bytes(0), 0)
            ),
            false
        );
        hevm.expectRevert("Only not in context");
        l1GatewayRouter.depositERC20(address(reentrantToken), 1, 0);
    }
}
