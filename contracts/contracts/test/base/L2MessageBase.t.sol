// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {CommonTest} from "./CommonTest.t.sol";

import {Proxy} from "../../libraries/proxy/Proxy.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {L2CrossDomainMessenger} from "../../L2/L2CrossDomainMessenger.sol";
import {L2ToL1MessagePasser} from "../../L2/system/L2ToL1MessagePasser.sol";
import {GasPriceOracle} from "../../L2/system/GasPriceOracle.sol";

contract L2MessageBaseTest is CommonTest {
    // L2ToL1MessagePasser config
    L2ToL1MessagePasser l2ToL1MessagePasser;
    L2ToL1MessagePasser l2ToL1MessagePasserImpl;

    bytes32 l2ToL1MessagePasser_leafNodesCount = bytes32(uint256(32));
    event AppendMessage(uint256 index, bytes32 messageHash);
    event SentMessage(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 messageNonce,
        uint256 gasLimit,
        bytes message
    );
    event RelayedMessage(bytes32 indexed messageHash);
    event FailedRelayedMessage(bytes32 indexed messageHash);

    // L2CrossDomainMessenger config
    L2CrossDomainMessenger l2CrossDomainMessenger;
    L2CrossDomainMessenger l2CrossDomainMessengerImpl;

    // GasPriceOracle config
    GasPriceOracle gasPriceOracle;

    uint256 gasOracle_owner_slot = 0;

    function setUp() public virtual override {
        super.setUp();

        // Set the proxy at the correct address
        hevm.etch(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            address(new Proxy(multisig)).code
        );
        hevm.etch(
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            address(new Proxy(multisig)).code
        );
        hevm.etch(
            Predeploys.GAS_PRICE_ORACLE,
            address(new GasPriceOracle(multisig)).code
        );
        Proxy l2ToL1MessagePasserProxy = Proxy(
            payable(Predeploys.L2_TO_L1_MESSAGE_PASSER)
        );
        Proxy l2CrossDomainMessengerProxy = Proxy(
            payable(Predeploys.L2_CROSS_DOMAIN_MESSENGER)
        );
        gasPriceOracle = GasPriceOracle(Predeploys.GAS_PRICE_ORACLE);

        hevm.store(
            address(l2ToL1MessagePasserProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2CrossDomainMessengerProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(gasPriceOracle),
            bytes32(gasOracle_owner_slot),
            bytes32(abi.encode(address(multisig)))
        );
        // deploy impl contracts and upgrade
        hevm.startPrank(multisig);

        l2ToL1MessagePasserImpl = new L2ToL1MessagePasser();
        l2ToL1MessagePasserProxy.upgradeTo(address(l2ToL1MessagePasserImpl));
        l2ToL1MessagePasser = L2ToL1MessagePasser(
            address(l2ToL1MessagePasserProxy)
        );

        l2CrossDomainMessengerImpl = new L2CrossDomainMessenger();
        l2CrossDomainMessengerProxy.upgradeToAndCall(
            address(l2CrossDomainMessengerImpl),
            abi.encodeWithSelector(
                L2CrossDomainMessenger.initialize.selector,
                NON_ZERO_ADDRESS
            )
        );
        l2CrossDomainMessenger = L2CrossDomainMessenger(
            payable(address(l2CrossDomainMessengerProxy))
        );

        hevm.stopPrank();
    }

    function setL1BaseFee(uint256 baseFee) internal {
        hevm.prank(multisig);
        gasPriceOracle.setL1BaseFee(baseFee);
    }
}
