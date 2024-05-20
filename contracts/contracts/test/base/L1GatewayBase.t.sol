// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L1MessageBaseTest} from "./L1MessageBase.t.sol";
import {L1GatewayRouter} from "../../l1/gateways/L1GatewayRouter.sol";
import {L1ETHGateway} from "../../l1/gateways/L1ETHGateway.sol";
import {L1ERC721Gateway} from "../../l1/gateways/L1ERC721Gateway.sol";
import {L1ERC1155Gateway} from "../../l1/gateways/L1ERC1155Gateway.sol";
import {L1StandardERC20Gateway} from "../../l1/gateways/L1StandardERC20Gateway.sol";
import {L1CustomERC20Gateway} from "../../l1/gateways/L1CustomERC20Gateway.sol";
import {L1WETHGateway} from "../../l1/gateways/L1WETHGateway.sol";
import {L2WETHGateway} from "../../l2/gateways/L2WETHGateway.sol";
import {MorphStandardERC20Factory} from "../../libraries/token/MorphStandardERC20Factory.sol";
import {MorphStandardERC20} from "../../libraries/token/MorphStandardERC20.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L1GatewayBaseTest is L1MessageBaseTest {
    uint256 internal constant EXTRA_VALUE = 1e17;

    // L1GatewayRouter config
    L1GatewayRouter public l1GatewayRouter;

    // L1ETHGateway config
    L1ETHGateway public l1ETHGateway;

    // L1StandardERC20Gateway config
    L1StandardERC20Gateway public l1StandardERC20Gateway;

    // L1CustomERC20Gateway config
    L1CustomERC20Gateway public l1CustomERC20Gateway;

    bool internal revertOnReceive;

    // MorphStandardERC20 config
    MorphStandardERC20 public template;

    // MorphStandardERC20Factory config
    MorphStandardERC20Factory public factory;

    // L1ERC721Gateway config
    L1ERC721Gateway public l1ERC721Gateway;

    // L1ERC1155Gateway config
    L1ERC1155Gateway public l1ERC1155Gateway;

    // L1WETHGateway config
    L1WETHGateway public l1WETHGateway;
    address public l2Messenger = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

    receive() external payable {
        if (revertOnReceive) {
            revert("RevertOnReceive");
        }
    }

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // deploy proxys
        TransparentUpgradeableProxy l1GatewayRouterProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1ETHGatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1StandardERC20GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        TransparentUpgradeableProxy l1CustomERC20GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );

        // deploy impls
        L1GatewayRouter l1GatewayRouterImpl = new L1GatewayRouter();
        L1ETHGateway l1ETHGatewayImpl = new L1ETHGateway();
        L1StandardERC20Gateway l1StandardERC20GatewayImpl = new L1StandardERC20Gateway();
        L1CustomERC20Gateway l1CustomERC20GatewayImpl = new L1CustomERC20Gateway();

        template = new MorphStandardERC20();
        factory = new MorphStandardERC20Factory(address(template));

        // upgrade and initialize
        ITransparentUpgradeableProxy(address(l1GatewayRouterProxy)).upgradeToAndCall(
            address(l1GatewayRouterImpl),
            abi.encodeCall(
                L1GatewayRouter.initialize,
                (address(l1ETHGatewayProxy), address(l1StandardERC20GatewayProxy))
            )
        );

        ITransparentUpgradeableProxy(address(l1ETHGatewayProxy)).upgradeToAndCall(
            address(l1ETHGatewayImpl),
            abi.encodeCall(
                L1ETHGateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1GatewayRouterProxy), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );
        ITransparentUpgradeableProxy(address(l1StandardERC20GatewayProxy)).upgradeToAndCall(
            address(l1StandardERC20GatewayImpl),
            abi.encodeCall(
                L1StandardERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1GatewayRouterProxy), // _router
                    address(l1CrossDomainMessenger), // _messenger
                    address(template), // _l2TokenImplementation
                    address(factory) // _l2TokenFactory
                )
            )
        );
        ITransparentUpgradeableProxy(address(l1CustomERC20GatewayProxy)).upgradeToAndCall(
            address(l1CustomERC20GatewayImpl),
            abi.encodeCall(
                L1CustomERC20Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1GatewayRouterProxy), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        l1GatewayRouter = L1GatewayRouter(address(l1GatewayRouterProxy));
        l1ETHGateway = L1ETHGateway(address(l1ETHGatewayProxy));
        l1StandardERC20Gateway = L1StandardERC20Gateway(address(l1StandardERC20GatewayProxy));
        l1CustomERC20Gateway = L1CustomERC20Gateway(address(l1CustomERC20GatewayProxy));
        _changeAdmin(address(l1GatewayRouter));
        _changeAdmin(address(l1ETHGateway));
        _changeAdmin(address(l1StandardERC20Gateway));
        _changeAdmin(address(l1CustomERC20Gateway));
        hevm.stopPrank();
    }

    function _deployERC721() public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l1ERC721GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        L1ERC721Gateway l1ERC721GatewayImpl = new L1ERC721Gateway();
        ITransparentUpgradeableProxy(address(l1ERC721GatewayProxy)).upgradeToAndCall(
            address(l1ERC721GatewayImpl),
            abi.encodeCall(
                L1ERC721Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );
        l1ERC721Gateway = L1ERC721Gateway(address(l1ERC721GatewayProxy));
        _changeAdmin(address(l1ERC721Gateway));

        hevm.stopPrank();
    }

    function _deployERC1155() public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l1ERC1155GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        L1ERC1155Gateway l1ERC1155GatewayImpl = new L1ERC1155Gateway();
        ITransparentUpgradeableProxy(address(l1ERC1155GatewayProxy)).upgradeToAndCall(
            address(l1ERC1155GatewayImpl),
            abi.encodeCall(
                L1ERC1155Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );

        l1ERC1155Gateway = L1ERC1155Gateway(address(l1ERC1155GatewayProxy));
        _changeAdmin(address(l1ERC1155Gateway));
        hevm.stopPrank();
    }

    function _deployWETH(address l1WETH, address l2WETH) public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l1WETHGatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        L1WETHGateway l1WETHGatewayImpl = new L1WETHGateway(l1WETH, l2WETH);
        L2WETHGateway l2WETHGateway = new L2WETHGateway(address(l2WETH), address(l1WETH));
        ITransparentUpgradeableProxy(address(l1WETHGatewayProxy)).upgradeToAndCall(
            address(l1WETHGatewayImpl),
            abi.encodeCall(
                L1WETHGateway.initialize,
                (
                    address(l2WETHGateway), // _counterpart
                    address(l1GatewayRouter), // _router
                    address(l1CrossDomainMessenger) // _messenger
                )
            )
        );
        l1WETHGateway = L1WETHGateway(payable(address(l1WETHGatewayProxy)));
        _changeAdmin(address(l1WETHGateway));

        hevm.stopPrank();
    }
}
