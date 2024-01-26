// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";

import {L1GatewayBaseTest} from "./base/L1GatewayBase.t.sol";
import {TransferReentrantToken} from "../mock/tokens/TransferReentrantToken.sol";

contract L1GatewayRouterTest is L1GatewayBaseTest {
    event SetETHGateway(
        address indexed oldETHGateway,
        address indexed newEthGateway
    );
    event SetDefaultERC20Gateway(
        address indexed oldDefaultERC20Gateway,
        address indexed newDefaultERC20Gateway
    );
    event SetERC20Gateway(
        address indexed token,
        address indexed oldGateway,
        address indexed newGateway
    );

    MockERC20 private l1Token;

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // Deploy tokens
        l1Token = new MockERC20("Mock", "M", 18);
        l1GatewayRouter.transferOwnership(address(this));
        hevm.stopPrank();
    }

    function testOwnership() public {
        assertEq(address(this), l1GatewayRouter.owner());
    }

    function testSetDefaultERC20Gateway() public {
        l1GatewayRouter.setDefaultERC20Gateway(address(0));

        // set by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        l1GatewayRouter.setDefaultERC20Gateway(address(l1StandardERC20Gateway));
        hevm.stopPrank();

        // set by owner, should succeed
        hevm.expectEmit(true, true, false, true);
        emit SetDefaultERC20Gateway(
            address(0),
            address(l1StandardERC20Gateway)
        );

        assertEq(address(0), l1GatewayRouter.getERC20Gateway(address(l1Token)));
        assertEq(address(0), l1GatewayRouter.defaultERC20Gateway());
        l1GatewayRouter.setDefaultERC20Gateway(address(l1StandardERC20Gateway));
        assertEq(
            address(l1StandardERC20Gateway),
            l1GatewayRouter.getERC20Gateway(address(l1Token))
        );
        assertEq(
            address(l1StandardERC20Gateway),
            l1GatewayRouter.defaultERC20Gateway()
        );
    }

    function testSetERC20Gateway() public {
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
        emit SetERC20Gateway(
            address(l1Token),
            address(0),
            address(l1StandardERC20Gateway)
        );

        assertEq(address(0), l1GatewayRouter.getERC20Gateway(address(l1Token)));
        l1GatewayRouter.setERC20Gateway(_tokens, _gateways);
        assertEq(
            address(l1StandardERC20Gateway),
            l1GatewayRouter.getERC20Gateway(address(l1Token))
        );
    }

    function testFinalizeWithdrawERC20() public {
        hevm.expectRevert("should never be called");
        l1GatewayRouter.finalizeWithdrawERC20(
            address(0),
            address(0),
            address(0),
            address(0),
            0,
            ""
        );
    }

    function testFinalizeWithdrawETH() public {
        hevm.expectRevert("should never be called");
        l1GatewayRouter.finalizeWithdrawETH(address(0), address(0), 0, "");
    }

    function testRequestERC20(
        address _sender,
        address _token,
        uint256 _amount
    ) public {
        hevm.expectRevert("Only in deposit context");
        l1GatewayRouter.requestERC20(_sender, _token, _amount);
    }

    function testReentrant() public {
        TransferReentrantToken reentrantToken = new TransferReentrantToken(
            "Reentrant",
            "R",
            18
        );
        reentrantToken.mint(address(this), type(uint128).max);
        reentrantToken.approve(address(l1GatewayRouter), type(uint256).max);

        reentrantToken.setReentrantCall(
            address(l1GatewayRouter),
            0,
            abi.encodeWithSelector(
                l1GatewayRouter.depositERC20AndCall.selector,
                address(reentrantToken),
                address(this),
                0,
                new bytes(0),
                0
            ),
            true
        );
        hevm.expectRevert("Only not in context");
        l1GatewayRouter.depositERC20(address(reentrantToken), 1, 0);

        reentrantToken.setReentrantCall(
            address(l1GatewayRouter),
            0,
            abi.encodeWithSelector(
                l1GatewayRouter.depositERC20AndCall.selector,
                address(reentrantToken),
                address(this),
                0,
                new bytes(0),
                0
            ),
            false
        );
        hevm.expectRevert("Only not in context");
        l1GatewayRouter.depositERC20(address(reentrantToken), 1, 0);
    }
}
