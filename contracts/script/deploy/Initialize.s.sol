// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Script} from "forge-std/Script.sol";
import {DeployConfig} from "../helpers/DeployConfig.sol";
import {Storage} from "../helpers/Storage.sol";
import {Types} from "../helpers/Types.sol";
import {ITransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

/// @title Initialize — upgrade proxies, call initializers, transfer admin
contract Initialize is Script {
    using DeployConfig for DeployConfig.DeployConfigCtx;

    /// @notice Run all three init stages: core → gateways → admin transfer.
    function run(string memory configPath, string memory spath) public {
        DeployConfig.DeployConfigCtx memory cfg = DeployConfig.load(vm, configPath);
        _initCore(cfg, spath);
        _initGateways(cfg, spath);
        _transferAdmins(spath);
    }

    // ---------------------------------------------------------------
    // Stage 1: initialize core contracts (Messenger, MessageQueue, Rollup, Staking, Sequencer, Router, ETH gateway)
    function _initCore(DeployConfig.DeployConfigCtx memory cfg, string memory spath) private {
        address messengerProxy = _g(spath, Types.PROXY_L1_CROSS_DOMAIN_MESSENGER);
        address msgQueueProxy  = _g(spath, Types.PROXY_L1_MESSAGE_QUEUE);
        address rollupProxy    = _g(spath, Types.PROXY_ROLLUP);
        address stakingProxy   = _g(spath, Types.PROXY_L1_STAKING);

        vm.startBroadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

        _doUpgrade(messengerProxy, _g(spath, Types.IMPL_L1_CROSS_DOMAIN_MESSENGER),
            abi.encodeWithSignature("initialize(address,address,address)",
                cfg.l1FeeVaultRecipient(), rollupProxy, msgQueueProxy));

        _doUpgrade(msgQueueProxy, _g(spath, Types.IMPL_L1_MESSAGE_QUEUE),
            abi.encodeWithSignature("initialize(uint256,address)",
                cfg.l1MessageQueueMaxGasLimit(), _g(spath, "Impl__Whitelist")));

        _initRollup(cfg, spath, rollupProxy, stakingProxy, msgQueueProxy);
        _initStaking(cfg, spath, stakingProxy, rollupProxy);
        _initSequencerRouter(cfg, spath, messengerProxy, msgQueueProxy, rollupProxy);

        vm.stopBroadcast();
    }

    // Stage 2: initialize remaining gateways
    function _initGateways(DeployConfig.DeployConfigCtx memory cfg, string memory spath) private {
        address messengerProxy = _g(spath, Types.PROXY_L1_CROSS_DOMAIN_MESSENGER);
        vm.startBroadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);

        _gwInit(spath, Types.PROXY_L1_STANDARD_ERC20_GATEWAY,     Types.IMPL_L1_STANDARD_ERC20_GATEWAY,     0x5300000000000000000000000000000000000008);
        _gwInit(spath, Types.PROXY_L1_CUSTOM_ERC20_GATEWAY,       Types.IMPL_L1_CUSTOM_ERC20_GATEWAY,       0x5300000000000000000000000000000000000016);
        _gwInit(spath, Types.PROXY_L1_WITHDRAW_LOCK_ERC20_GATEWAY,Types.IMPL_L1_WITHDRAW_LOCK_ERC20_GATEWAY,0x5300000000000000000000000000000000000019);
        _gwInit(spath, Types.PROXY_L1_REVERSE_CUSTOM_GATEWAY,     Types.IMPL_L1_REVERSE_CUSTOM_GATEWAY,     0x5300000000000000000000000000000000000018);
        _gwInit(spath, Types.PROXY_L1_WETH_GATEWAY,               Types.IMPL_L1_WETH_GATEWAY,               0x5300000000000000000000000000000000000010);

_doUpgrade(_g(spath, Types.PROXY_L1_ERC721_GATEWAY), _g(spath, Types.IMPL_L1_ERC721_GATEWAY),
            abi.encodeWithSignature("initialize(address,address)", 0x5300000000000000000000000000000000000009, messengerProxy));
_doUpgrade(_g(spath, Types.PROXY_L1_ERC1155_GATEWAY), _g(spath, Types.IMPL_L1_ERC1155_GATEWAY),
            abi.encodeWithSignature("initialize(address,address)", 0x530000000000000000000000000000000000000c, messengerProxy));
_doUpgrade(_g(spath, Types.PROXY_ENFORCED_TX_GATEWAY), _g(spath, Types.IMPL_ENFORCED_TX_GATEWAY),
            abi.encodeWithSignature("initialize(address,address)",
                _g(spath, Types.PROXY_L1_MESSAGE_QUEUE), cfg.l1FeeVaultRecipient()));

        vm.stopBroadcast();
    }

    // Stage 3: transfer all proxy admin ownerships to ProxyAdmin
    function _transferAdmins(string memory spath) private {
        vm.startBroadcast(0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80);
        address padm = _g(spath, Types.IMPL_PROXY_ADMIN);

        _adm(_g(spath, Types.PROXY_L1_MESSAGE_QUEUE), padm);
        _adm(_g(spath, Types.PROXY_ROLLUP), padm);
        _adm(_g(spath, Types.PROXY_L1_STAKING), padm);
        _adm(_g(spath, Types.PROXY_L1_SEQUENCER), padm);
        _adm(_g(spath, Types.PROXY_L1_GATEWAY_ROUTER), padm);
        _adm(_g(spath, Types.PROXY_L1_ETH_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_STANDARD_ERC20_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_CUSTOM_ERC20_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_WITHDRAW_LOCK_ERC20_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_REVERSE_CUSTOM_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_WETH_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_ERC721_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_ERC1155_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_ENFORCED_TX_GATEWAY), padm);
        _adm(_g(spath, Types.PROXY_L1_CROSS_DOMAIN_MESSENGER), padm);

        vm.stopBroadcast();
    }

    // --- helpers ---
    function _g(string memory spath, string memory k) private returns (address a) {
        return Storage.getOrFail(vm, spath, k);
    }
    function _doUpgrade(address proxy, address impl, bytes memory data) private {
        ITransparentUpgradeableProxy(proxy).upgradeToAndCall(impl, data);
    }
    function _initRollup(DeployConfig.DeployConfigCtx memory c, string memory spath,
        address rp, address sp, address mqp) private {
        _doUpgrade(rp, _g(spath, Types.IMPL_ROLLUP),
            abi.encodeWithSignature("initialize(address,address,address,uint256,uint256,uint256)",
                sp, mqp, _g(spath, Types.IMPL_MULTIPLE_VERSION_ROLLUP_VERIFIER),
                c.finalizationPeriodSeconds(), c.rollupProofWindow(), c.proofRewardPercent()));
    }
    function _initStaking(DeployConfig.DeployConfigCtx memory c, string memory spath,
        address sp, address rp) private {
        _doUpgrade(sp, _g(spath, Types.IMPL_L1_STAKING),
            abi.encodeWithSignature("initialize(address,uint256,uint256,uint256,uint256,uint256,uint256)",
                rp,
                c.stakingMinDeposit() * 1 ether,
                c.stakingChallengeDeposit() * 1 ether,
                c.stakingLockNumber(),
                c.stakingChallengerRewardPercentage(),
                c.stakingCrossChainGaslimitAdd(),
                c.stakingCrossChainGaslimitRemove()));
    }
    function _initSequencerRouter(DeployConfig.DeployConfigCtx memory c, string memory spath,
        address mPxy, address mqPxy, address rPxy) private {
        _doUpgrade(_g(spath, Types.PROXY_L1_SEQUENCER), _g(spath, Types.IMPL_L1_SEQUENCER),
            abi.encodeWithSignature("initialize(address)", c.contractAdmin()));
        address router = _g(spath, Types.PROXY_L1_GATEWAY_ROUTER);
        address ethGw  = _g(spath, Types.PROXY_L1_ETH_GATEWAY);
        _doUpgrade(router, _g(spath, Types.IMPL_L1_GATEWAY_ROUTER),
            abi.encodeWithSignature("initialize(address,address)", ethGw, c.contractAdmin()));
        _doUpgrade(ethGw, _g(spath, Types.IMPL_L1_ETH_GATEWAY),
            abi.encodeWithSignature("initialize(address,address,address)",
                0x5300000000000000000000000000000000000006, router, mPxy));
    }
    function _gwInit(string memory spath, string memory proxyKey, string memory implKey, address counterpart) private {
        _doUpgrade(_g(spath, proxyKey), _g(spath, implKey),
            abi.encodeWithSignature("initialize(address,address,address)",
                counterpart, _g(spath, Types.PROXY_L1_GATEWAY_ROUTER), _g(spath, Types.PROXY_L1_CROSS_DOMAIN_MESSENGER)));
    }
    function _adm(address proxy, address newAdmin) private {
        ITransparentUpgradeableProxy(proxy).changeAdmin(newAdmin);
    }
}
