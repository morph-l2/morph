// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Proxy} from "../../libraries/proxy/Proxy.sol";
import {L1MessageBaseTest} from "./L1MessageBase.t.sol";
import {L1GatewayRouter} from "../../L1/gateways/L1GatewayRouter.sol";
import {L1ETHGateway} from "../../L1/gateways/L1ETHGateway.sol";
import {L1ERC721Gateway} from "../../L1/gateways/L1ERC721Gateway.sol";
import {L1ERC1155Gateway} from "../../L1/gateways/L1ERC1155Gateway.sol";
import {L1StandardERC20Gateway} from "../../L1/gateways/L1StandardERC20Gateway.sol";
import {L1CustomERC20Gateway} from "../../L1/gateways/L1CustomERC20Gateway.sol";
import {L1WETHGateway} from "../../L1/gateways/L1WETHGateway.sol";
import {L2WETHGateway} from "../../L2/gateways/L2WETHGateway.sol";
import {MorphStandardERC20Factory} from "../../libraries/token/MorphStandardERC20Factory.sol";
import {MorphStandardERC20} from "../../libraries/token/MorphStandardERC20.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L1GatewayBaseTest is L1MessageBaseTest {
    // L1GatewayRouter config
    L1GatewayRouter l1GatewayRouter;

    // L1ETHGateway config
    L1ETHGateway l1ETHGateway;

    uint256 internal constant extraValue = 1e17;

    // L1StandardERC20Gateway config
    L1StandardERC20Gateway l1StandardERC20Gateway;

    // L1CustomERC20Gateway config
    L1CustomERC20Gateway l1CustomERC20Gateway;

    bool internal revertOnReceive;

    // MorphStandardERC20 config
    MorphStandardERC20 template;

    // MorphStandardERC20Factory config
    MorphStandardERC20Factory factory;

    // L1ERC721Gateway config
    L1ERC721Gateway l1ERC721Gateway;

    // L1ERC1155Gateway config
    L1ERC1155Gateway l1ERC1155Gateway;

    // L1WETHGateway config
    L1WETHGateway l1WETHGateway;
    address l2Messenger = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

    receive() external payable {
        if (revertOnReceive) {
            revert("RevertOnReceive");
        }
    }

    function setUp() public virtual override {
        super.setUp();
        hevm.startPrank(multisig);
        // deploy proxys
        Proxy l1GatewayRouterProxy = new Proxy(multisig);
        Proxy l1ETHGatewayProxy = new Proxy(multisig);
        Proxy l1StandardERC20GatewayProxy = new Proxy(multisig);
        Proxy l1CustomERC20GatewayProxy = new Proxy(multisig);

        // deploy impls
        L1GatewayRouter l1GatewayRouterImpl = new L1GatewayRouter();
        L1ETHGateway l1ETHGatewayImpl = new L1ETHGateway();
        L1StandardERC20Gateway l1StandardERC20GatewayImpl = new L1StandardERC20Gateway();
        L1CustomERC20Gateway l1CustomERC20GatewayImpl = new L1CustomERC20Gateway();

        template = new MorphStandardERC20();
        factory = new MorphStandardERC20Factory(address(template));

        // upgrade and initialize
        l1GatewayRouterProxy.upgradeToAndCall(
            address(l1GatewayRouterImpl),
            abi.encodeWithSelector(
                L1GatewayRouter.initialize.selector,
                address(l1ETHGatewayProxy),
                address(l1StandardERC20GatewayProxy)
            )
        );
        l1ETHGatewayProxy.upgradeToAndCall(
            address(l1ETHGatewayImpl),
            abi.encodeWithSelector(
                L1ETHGateway.initialize.selector,
                address(NON_ZERO_ADDRESS), // _counterpart
                address(l1GatewayRouterProxy), // _router
                address(l1CrossDomainMessenger) // _messenger
            )
        );
        l1StandardERC20GatewayProxy.upgradeToAndCall(
            address(l1StandardERC20GatewayImpl),
            abi.encodeWithSelector(
                L1StandardERC20Gateway.initialize.selector,
                address(NON_ZERO_ADDRESS), // _counterpart
                address(l1GatewayRouterProxy), // _router
                address(l1CrossDomainMessenger), // _messenger
                address(template), // _l2TokenImplementation
                address(factory) // _l2TokenFactory
            )
        );
        l1CustomERC20GatewayProxy.upgradeToAndCall(
            address(l1CustomERC20GatewayImpl),
            abi.encodeWithSelector(
                L1CustomERC20Gateway.initialize.selector,
                address(NON_ZERO_ADDRESS), // _counterpart
                address(l1GatewayRouterProxy), // _router
                address(l1CrossDomainMessenger) // _messenger
            )
        );

        l1GatewayRouter = L1GatewayRouter(address(l1GatewayRouterProxy));
        l1ETHGateway = L1ETHGateway(address(l1ETHGatewayProxy));
        l1StandardERC20Gateway = L1StandardERC20Gateway(
            address(l1StandardERC20GatewayProxy)
        );
        l1CustomERC20Gateway = L1CustomERC20Gateway(
            address(l1CustomERC20GatewayProxy)
        );

        hevm.stopPrank();
    }

    function _deployERC721() public {
        hevm.startPrank(multisig);
        Proxy l1ERC721GatewayProxy = new Proxy(multisig);
        L1ERC721Gateway l1ERC721GatewayImpl = new L1ERC721Gateway();
        l1ERC721GatewayProxy.upgradeToAndCall(
            address(l1ERC721GatewayImpl),
            abi.encodeWithSelector(
                L1ERC721Gateway.initialize.selector,
                address(NON_ZERO_ADDRESS), // _counterpart
                address(l1CrossDomainMessenger) // _messenger
            )
        );
        l1ERC721Gateway = L1ERC721Gateway(address(l1ERC721GatewayProxy));
        hevm.stopPrank();
    }

    function _deployERC1155() public {
        hevm.startPrank(multisig);
        Proxy l1ERC1155GatewayProxy = new Proxy(multisig);
        L1ERC1155Gateway l1ERC1155GatewayImpl = new L1ERC1155Gateway();
        l1ERC1155GatewayProxy.upgradeToAndCall(
            address(l1ERC1155GatewayImpl),
            abi.encodeWithSelector(
                L1ERC1155Gateway.initialize.selector,
                address(NON_ZERO_ADDRESS), // _counterpart
                address(l1CrossDomainMessenger) // _messenger
            )
        );
        l1ERC1155Gateway = L1ERC1155Gateway(address(l1ERC1155GatewayProxy));
        hevm.stopPrank();
    }

    function _deployWETH(address l1WETH, address l2WETH) public {
        hevm.startPrank(multisig);
        Proxy l1WETHGatewayProxy = new Proxy(multisig);
        L1WETHGateway l1WETHGatewayImpl = new L1WETHGateway(l1WETH, l2WETH);
        L2WETHGateway l2WETHGateway = new L2WETHGateway(
            address(l2WETH),
            address(l1WETH)
        );
        l1WETHGatewayProxy.upgradeToAndCall(
            address(l1WETHGatewayImpl),
            abi.encodeWithSelector(
                L1WETHGateway.initialize.selector,
                address(l2WETHGateway), // _counterpart
                address(l1GatewayRouter), // _router
                address(l1CrossDomainMessenger) // _messenger
            )
        );
        l1WETHGateway = L1WETHGateway(payable(address(l1WETHGatewayProxy)));
        hevm.stopPrank();
    }
}
