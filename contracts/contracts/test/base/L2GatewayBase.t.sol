// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L2MessageBaseTest} from "./L2MessageBase.t.sol";
import {L2ERC721Gateway} from "../../L2/gateways/L2ERC721Gateway.sol";
import {L2ETHGateway} from "../../L2/gateways/L2ETHGateway.sol";
import {L2ERC20Gateway} from "../../L2/gateways/L2ERC20Gateway.sol";
import {L2StandardERC20Gateway} from "../../L2/gateways/L2StandardERC20Gateway.sol";
import {L2ERC721Gateway} from "../../L2/gateways/L2ERC721Gateway.sol";
import {L2ERC1155Gateway} from "../../L2/gateways/L2ERC1155Gateway.sol";
import {L2GatewayRouter} from "../../L2/gateways/L2GatewayRouter.sol";
import {L2CustomERC20Gateway} from "../../L2/gateways/L2CustomERC20Gateway.sol";
import {L2WETHGateway} from "../../L2/gateways/L2WETHGateway.sol";
import {L1WETHGateway} from "../../L1/gateways/L1WETHGateway.sol";
import {MorphStandardERC20} from "../../libraries/token/MorphStandardERC20.sol";
import {MorphStandardERC20Factory} from "../../libraries/token/MorphStandardERC20Factory.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";

contract L2GatewayBaseTest is L2MessageBaseTest {
    // L2GatewayRouter config
    L2GatewayRouter l2GatewayRouter;

    // L2ETHGateway config
    L2ETHGateway l2ETHGateway;

    // L2ERC20Gateway config
    L2ERC20Gateway l2ERC20Gateway;

    // L2CustomERC20Gateway config
    L2CustomERC20Gateway l2CustomERC20Gateway;

    // L2StandardERC20Gateway config
    L2StandardERC20Gateway l2StandardERC20Gateway;

    // L2ERC721Gateway config
    L2ERC721Gateway l2ERC721Gateway;

    // L2ERC1155Gateway config
    L2ERC1155Gateway l2ERC1155Gateway;

    // MorphStandardERC20 config
    MorphStandardERC20 template;

    // MorphStandardERC20Factory config
    MorphStandardERC20Factory factory;

    // l2WETHGateway config
    L2WETHGateway l2WETHGateway;

    bool internal revertOnReceive;

    address l2FeeVault = address(3033);
    uint256 eth_erc20_messenger_slot = 153;
    uint256 erc721_messenger_slot = 203;
    uint256 erc1155_messenger_slot = 303;

    receive() external payable {
        if (revertOnReceive) {
            revert("RevertOnReceive");
        }
    }

    function setUp() public virtual override {
        super.setUp();

        // deploy proxy
        hevm.etch(
            Predeploys.L2_Gateway_Router,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_ETH_Gateway,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_Standard_ERC20_Gateway,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_ERC721_Gateway,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );

        TransparentUpgradeableProxy l2GatewayRouterProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_Gateway_Router)
            );
        TransparentUpgradeableProxy l2ETHGatewayProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_ETH_Gateway)
            );
        TransparentUpgradeableProxy l2StandardERC20GatewayProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_Standard_ERC20_Gateway)
            );
        TransparentUpgradeableProxy l2ERC721GatewayProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_ERC721_Gateway)
            );
        hevm.store(
            address(l2GatewayRouterProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2ETHGatewayProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2StandardERC20GatewayProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2ERC721GatewayProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );

        // deploy impl
        hevm.startPrank(multisig);
        L2GatewayRouter l2GatewayRouterImpl = new L2GatewayRouter();
        L2ETHGateway l2ETHGatewayImpl = new L2ETHGateway();
        L2StandardERC20Gateway l2StandardERC20GatewayImpl = new L2StandardERC20Gateway();
        L2ERC721Gateway l2ERC721GatewayImpl = new L2ERC721Gateway();

        template = new MorphStandardERC20();
        factory = new MorphStandardERC20Factory(address(template));

        // upgrade and initialize
        ITransparentUpgradeableProxy(address(l2GatewayRouterProxy))
            .upgradeToAndCall(
                address(l2GatewayRouterImpl),
                abi.encodeCall(
                    L2GatewayRouter.initialize,(
                    address(l2ETHGatewayProxy), // eth gateway
                    address(l2StandardERC20GatewayProxy) // erc20 gateway
                ))
            );
        ITransparentUpgradeableProxy(address(l2ETHGatewayProxy))
            .upgradeToAndCall(
                address(l2ETHGatewayImpl),
                abi.encodeCall(
                    L2ETHGateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2GatewayRouterProxy), // _router
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        ITransparentUpgradeableProxy(address(l2StandardERC20GatewayProxy))
            .upgradeToAndCall(
                address(l2StandardERC20GatewayImpl),
                abi.encodeCall(
                    L2StandardERC20Gateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2GatewayRouterProxy), // _router
                    address(l2CrossDomainMessenger), // _messenger
                    address(factory) // _tokenFactory
                ))
            );

        ITransparentUpgradeableProxy(address(l2ERC721GatewayProxy))
            .upgradeToAndCall(
                address(l2ERC721GatewayImpl),
                abi.encodeCall(
                    L2ERC721Gateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        l2GatewayRouter = L2GatewayRouter(
            payable(address(l2GatewayRouterProxy))
        );
        l2ETHGateway = L2ETHGateway(payable(address(l2ETHGatewayProxy)));
        l2StandardERC20Gateway = L2StandardERC20Gateway(
            address(l2StandardERC20GatewayProxy)
        );
        l2ERC721Gateway = L2ERC721Gateway(address(l2ERC721GatewayProxy));

        _changeAdmin(address(l2GatewayRouter));
        _changeAdmin(address(l2ETHGateway));
        _changeAdmin(address(l2StandardERC20Gateway));
        _changeAdmin(address(l2ERC721Gateway));
        hevm.stopPrank();
    }

    function _deployCustomERC20() public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l2CustomERC20GatewayProxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        L2CustomERC20Gateway l2CustomERC20GatewayImpl = new L2CustomERC20Gateway();
        ITransparentUpgradeableProxy(address(l2CustomERC20GatewayProxy))
            .upgradeToAndCall(
                address(l2CustomERC20GatewayImpl),
                abi.encodeCall(
                    L2CustomERC20Gateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2GatewayRouter), // _router
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        l2CustomERC20Gateway = L2CustomERC20Gateway(
            address(l2CustomERC20GatewayProxy)
        );
        _changeAdmin(address(l2CustomERC20Gateway));
        hevm.stopPrank();
    }

    function _deployERC721() public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l2ERC721GatewayProxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        L2ERC721Gateway l2ERC721GatewayImpl = new L2ERC721Gateway();
        ITransparentUpgradeableProxy(address(l2ERC721GatewayProxy))
            .upgradeToAndCall(
                address(l2ERC721GatewayImpl),
                abi.encodeCall(
                    L2ERC721Gateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        l2ERC721Gateway = L2ERC721Gateway(address(l2ERC721GatewayProxy));
        _changeAdmin(address(l2ERC721Gateway));

        hevm.stopPrank();
    }

    function _deployERC1155() public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l2ERC1155GatewayProxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        L2ERC1155Gateway l2ERC1155GatewayImpl = new L2ERC1155Gateway();
        ITransparentUpgradeableProxy(address(l2ERC1155GatewayProxy))
            .upgradeToAndCall(
                address(l2ERC1155GatewayImpl),
                abi.encodeCall(
                    L2ERC1155Gateway.initialize,(
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        l2ERC1155Gateway = L2ERC1155Gateway(address(l2ERC1155GatewayProxy));
        _changeAdmin(address(l2ERC1155Gateway));

        hevm.stopPrank();
    }

    function _deployWETH(address l1WETH, address l2WETH) public {
        hevm.startPrank(multisig);
        TransparentUpgradeableProxy l2WETHGatewayProxy = new TransparentUpgradeableProxy(
                address(emptyContract),
                address(multisig),
                new bytes(0)
            );
        L1WETHGateway l1WETHGatewayImpl = new L1WETHGateway(l1WETH, l2WETH);
        L2WETHGateway l2WETHGatewayImpl = new L2WETHGateway(
            address(l2WETH),
            address(l1WETH)
        );
        ITransparentUpgradeableProxy(address(l2WETHGatewayProxy))
            .upgradeToAndCall(
                address(l2WETHGatewayImpl),
                abi.encodeCall(
                    L2WETHGateway.initialize,(
                    address(l1WETHGatewayImpl), // _counterpart
                    address(l2GatewayRouter), // _router
                    address(l2CrossDomainMessenger) // _messenger
                ))
            );

        l2WETHGateway = L2WETHGateway(payable(address(l2WETHGatewayProxy)));
        _changeAdmin(address(l2WETHGateway));

        hevm.stopPrank();
    }
}
