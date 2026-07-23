// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Script} from "forge-std/Script.sol";
import {VmSafe} from "forge-std/Vm.sol";
import {DeployConfig} from "../helpers/DeployConfig.sol";
import {Storage} from "../helpers/Storage.sol";
import {Types} from "../helpers/Types.sol";

// Contracts
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {EmptyContract} from "../../contracts/misc/EmptyContract.sol";
import {ZkEvmVerifierV1} from "../../contracts/libraries/verifier/ZkEvmVerifierV1.sol";
import {WrappedEther} from "../../contracts/l2/system/WrappedEther.sol";

/// @title Deploy — Phase 1 deployment script
/// @notice Deploys ProxyAdmin, EmptyContract, ZkEvmVerifierV1, and all
///         TransparentUpgradeableProxy contracts.  Writes addresses to the
///         storage-path JSON file so Phase 2 (Initialize) can read them.
///
/// Usage:
///   forge script script/deploy/Deploy.s.sol --broadcast \
///     --rpc-url http://localhost:9545 \
///     --sig "run(string,string)" \
///     script/config/l1.json /path/to/devnetL1.json
contract Deploy is Script {
    using DeployConfig for DeployConfig.DeployConfigCtx;

    function run(string memory configPath, string memory storagePath) public {
        DeployConfig.DeployConfigCtx memory cfg = DeployConfig.load(vm, configPath);

        vm.startBroadcast();

        // -----------------------------------------------------------
        // 1. Deploy ProxyAdmin
        address proxyAdminAddr = address(new ProxyAdmin());

        // 2. Deploy EmptyContract (used as initial implementation for proxies)
        address emptyContractAddr = address(new EmptyContract());

        // 3. Deploy ZkEvmVerifierV1
        address verifierAddr = address(new ZkEvmVerifierV1(cfg.programVkey()));

        // 4. Deploy all TransparentUpgradeableProxy contracts
        address[] memory addrs = new address[](16);
        string[] memory names = new string[](16);
        uint256 idx;

        // 4a - WETH (optional)
        string memory wethAddr = cfg.l1WETHAddress();
        address wethProxy = address(0);
        if (bytes(wethAddr).length == 0) {
            // Deploy WETH implementation and proxy
            address wethImpl = address(new WrappedEther());
            // Proxy for WETH: initialize with empty data (placeholder)
            wethProxy = address(new TransparentUpgradeableProxy(wethImpl, cfg.contractAdmin(), ""));
            names[idx] = Types.IMPL_WETH;
            addrs[idx] = wethImpl;
            idx++;
        }

        // 4b - L1CrossDomainMessenger
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_CROSS_DOMAIN_MESSENGER;
        idx++;
        // 4c - L1MessageQueueWithGasPriceOracle
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_MESSAGE_QUEUE;
        idx++;
        // 4d - Rollup
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_ROLLUP;
        idx++;
        // 4e - L1Staking
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_STAKING;
        idx++;
        // 4f - L1Sequencer
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_SEQUENCER;
        idx++;
        // 4g - L1GatewayRouter
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_GATEWAY_ROUTER;
        idx++;
        // 4h - L1ETHGateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_ETH_GATEWAY;
        idx++;
        // 4i - L1StandardERC20Gateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_STANDARD_ERC20_GATEWAY;
        idx++;
        // 4j - L1CustomERC20Gateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_CUSTOM_ERC20_GATEWAY;
        idx++;
        // 4k - L1WithdrawLockERC20Gateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_WITHDRAW_LOCK_ERC20_GATEWAY;
        idx++;
        // 4l - L1ReverseCustomGateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_REVERSE_CUSTOM_GATEWAY;
        idx++;
        // 4m - L1ERC721Gateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_ERC721_GATEWAY;
        idx++;
        // 4n - L1ERC1155Gateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_ERC1155_GATEWAY;
        idx++;
        // 4o - EnforcedTxGateway
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_ENFORCED_TX_GATEWAY;
        idx++;
        // 4p - L1WETHGateway (always deploy proxy, init later)
        addrs[idx] = address(new TransparentUpgradeableProxy(emptyContractAddr, cfg.contractAdmin(), ""));
        names[idx] = Types.PROXY_L1_WETH_GATEWAY;
        idx++;

        vm.stopBroadcast();

        // Collect all storage entries and write
        string[] memory allNames = new string[](idx + 3); // +3 for ProxyAdmin, EmptyContract, Verifier
        address[] memory allAddrs = new address[](idx + 3);
        allNames[0] = Types.IMPL_PROXY_ADMIN;
        allAddrs[0] = proxyAdminAddr;
        allNames[1] = Types.IMPL_EMPTY_CONTRACT;
        allAddrs[1] = emptyContractAddr;
        allNames[2] = Types.IMPL_ZKEVM_VERIFIER_V1;
        allAddrs[2] = verifierAddr;
        for (uint256 i = 0; i < idx; i++) {
            allNames[i + 3] = names[i];
            allAddrs[i + 3] = addrs[i];
        }

        Storage.setMany(vm, storagePath, allNames, allAddrs);
    }

    function _s(string memory a, string memory b) private pure returns (string[] memory) {
        string[] memory arr = new string[](2);
        arr[0] = a;
        arr[1] = b;
        return arr;
    }

    function _a(address a, address b) private pure returns (address[] memory) {
        address[] memory arr = new address[](2);
        arr[0] = a;
        arr[1] = b;
        return arr;
    }
}
