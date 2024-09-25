// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {EvmVerifier} from "../src/EvmVerifier.sol";

contract EvmScript is Script {
    EvmVerifier public evm;

    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        evm = new EvmVerifier("");

        vm.stopBroadcast();
    }
}
