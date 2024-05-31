// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {ITransparentUpgradeableProxy, TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {CommonTest} from "./CommonTest.t.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {L2CrossDomainMessenger} from "../../l2/L2CrossDomainMessenger.sol";
import {L2ToL1MessagePasser} from "../../l2/system/L2ToL1MessagePasser.sol";
import {GasPriceOracle} from "../../l2/system/GasPriceOracle.sol";

contract L2MessageBaseTest is CommonTest {
    // L2ToL1MessagePasser config
    L2ToL1MessagePasser public l2ToL1MessagePasser;
    L2ToL1MessagePasser public l2ToL1MessagePasserImpl;

    bytes32 public l2ToL1MessagePasserLeafNodesCount = bytes32(uint256(32));

    // L2CrossDomainMessenger config
    L2CrossDomainMessenger public l2CrossDomainMessenger;
    L2CrossDomainMessenger public l2CrossDomainMessengerImpl;

    // GasPriceOracle config
    GasPriceOracle public gasPriceOracle;

    uint256 public gasOracle_owner_slot = 0;

    function setUp() public virtual override {
        super.setUp();

        // Set the proxy at the correct address
        hevm.etch(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            address(new TransparentUpgradeableProxy(address(emptyContract), address(multisig), new bytes(0))).code
        );
        hevm.etch(Predeploys.GAS_PRICE_ORACLE, address(new GasPriceOracle(multisig)).code);

        TransparentUpgradeableProxy l2ToL1MessagePasserProxy = TransparentUpgradeableProxy(
            payable(Predeploys.L2_TO_L1_MESSAGE_PASSER)
        );
        TransparentUpgradeableProxy l2CrossDomainMessengerProxy = TransparentUpgradeableProxy(
            payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER)
        );
        gasPriceOracle = GasPriceOracle(Predeploys.GAS_PRICE_ORACLE);

        hevm.store(address(l2ToL1MessagePasserProxy), bytes32(PROXY_OWNER_KEY), bytes32(abi.encode(address(multisig))));
        hevm.store(
            address(l2CrossDomainMessengerProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(address(gasPriceOracle), bytes32(gasOracle_owner_slot), bytes32(abi.encode(address(multisig))));
        // deploy impl contracts and upgrade
        hevm.startPrank(multisig);

        l2ToL1MessagePasserImpl = new L2ToL1MessagePasser();
        ITransparentUpgradeableProxy(address(l2ToL1MessagePasserProxy)).upgradeTo(address(l2ToL1MessagePasserImpl));
        l2ToL1MessagePasser = L2ToL1MessagePasser(address(l2ToL1MessagePasserProxy));

        l2CrossDomainMessengerImpl = new L2CrossDomainMessenger();
        ITransparentUpgradeableProxy(address(l2CrossDomainMessengerProxy)).upgradeToAndCall(
            address(l2CrossDomainMessengerImpl),
            abi.encodeCall(L2CrossDomainMessenger.initialize, (NON_ZERO_ADDRESS))
        );
        l2CrossDomainMessenger = L2CrossDomainMessenger(payable(address(l2CrossDomainMessengerProxy)));
        _changeAdmin(address(l2ToL1MessagePasser));
        _changeAdmin(address(l2CrossDomainMessenger));
        hevm.stopPrank();
    }

    function setL1BaseFee(uint256 baseFee) internal {
        hevm.prank(multisig);
        gasPriceOracle.setL1BaseFee(baseFee);
    }
}
