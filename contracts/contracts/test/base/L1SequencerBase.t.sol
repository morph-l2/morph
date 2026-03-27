// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import {L1Sequencer} from "../../l1/L1Sequencer.sol";

contract L1SequencerBaseTest is Test {
    L1Sequencer public l1Sequencer;
    ProxyAdmin public proxyAdmin;

    address public owner = address(0x1234);
    address public nonOwner = address(0x5678);
    address public sequencerA = address(0xA001);
    address public sequencerB = address(0xA002);
    address public sequencerC = address(0xA003);

    uint64 public constant UPGRADE_HEIGHT = 100;

    function setUp() public virtual {
        vm.startPrank(owner);

        proxyAdmin = new ProxyAdmin();
        L1Sequencer impl = new L1Sequencer();

        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(impl),
            address(proxyAdmin),
            abi.encodeWithSelector(L1Sequencer.initialize.selector, owner)
        );

        l1Sequencer = L1Sequencer(address(proxy));
        vm.stopPrank();
    }

    function _initHistory(address seq, uint64 upgradeHeight) internal {
        vm.prank(owner);
        l1Sequencer.initializeHistory(seq, upgradeHeight);
    }
}
