// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Script} from "forge-std/Script.sol";
import {DeployConfig} from "../helpers/DeployConfig.sol";
import {Storage} from "../helpers/Storage.sol";
import {Types} from "../helpers/Types.sol";

// L1 implementation contracts
import {L1CrossDomainMessenger} from "../../contracts/l1/L1CrossDomainMessenger.sol";
import {L1MessageQueueWithGasPriceOracle} from "../../contracts/l1/rollup/L1MessageQueueWithGasPriceOracle.sol";
import {Rollup} from "../../contracts/l1/rollup/Rollup.sol";
import {MultipleVersionRollupVerifier} from "../../contracts/l1/rollup/MultipleVersionRollupVerifier.sol";
import {L1Staking} from "../../contracts/l1/staking/L1Staking.sol";
import {L1Sequencer} from "../../contracts/l1/L1Sequencer.sol";
import {L1GatewayRouter} from "../../contracts/l1/gateways/L1GatewayRouter.sol";
import {L1ETHGateway} from "../../contracts/l1/gateways/L1ETHGateway.sol";
import {L1StandardERC20Gateway} from "../../contracts/l1/gateways/L1StandardERC20Gateway.sol";
import {L1CustomERC20Gateway} from "../../contracts/l1/gateways/L1CustomERC20Gateway.sol";
import {L1WETHGateway} from "../../contracts/l1/gateways/L1WETHGateway.sol";
import {L1ERC721Gateway} from "../../contracts/l1/gateways/L1ERC721Gateway.sol";
import {L1ERC1155Gateway} from "../../contracts/l1/gateways/L1ERC1155Gateway.sol";
import {EnforcedTxGateway} from "../../contracts/l1/gateways/EnforcedTxGateway.sol";
import {Whitelist} from "../../contracts/libraries/common/Whitelist.sol";

/// @title DeployImpls — deploy all implementation contracts
///
/// Usage:
///   forge script script/deploy/InitializeImpls.s.sol --broadcast \
///     --rpc-url http://localhost:9545 \
///     --sig "run(string,string)" script/config/l1.json path/to/devnetL1.json
contract DeployImpls is Script {
    using DeployConfig for DeployConfig.DeployConfigCtx;

    address constant L2_WETH = 0x5300000000000000000000000000000000000011;

    function run(string memory configPath, string memory proxyFile, string memory outFile) public {
        DeployConfig.DeployConfigCtx memory cfg = DeployConfig.load(vm, configPath);

        address messengerProxy = Storage.getOrFail(vm, proxyFile, Types.PROXY_L1_CROSS_DOMAIN_MESSENGER);
        address rollupProxy    = Storage.getOrFail(vm, proxyFile, Types.PROXY_ROLLUP);
        address enforcedTxPxy  = Storage.getOrFail(vm, proxyFile, Types.PROXY_ENFORCED_TX_GATEWAY);
        address verifierAddr   = Storage.getOrFail(vm, proxyFile, Types.IMPL_ZKEVM_VERIFIER_V1);
        address wethImpl       = Storage.get(vm, proxyFile, Types.IMPL_WETH);

        vm.startBroadcast();

        // Deploy all 17 implementations
        address aWhitelist     = address(new Whitelist(cfg.contractAdmin()));
        address aMessenger     = address(new L1CrossDomainMessenger());
        address aMsgQueue      = address(new L1MessageQueueWithGasPriceOracle(messengerProxy, rollupProxy, enforcedTxPxy));
        address aRollup        = address(new Rollup(uint64(cfg.l2ChainID())));

        uint256[] memory vs = new uint256[](1); vs[0]=0;
        address[] memory zvs = new address[](1); zvs[0]=verifierAddr;
        address aMultiVerifier = address(new MultipleVersionRollupVerifier(vs, zvs));

        address aRouter        = address(new L1GatewayRouter());
        address aETHGateway    = address(new L1ETHGateway());
        address aStdERC20      = address(new L1StandardERC20Gateway());
        address aCustERC20     = address(new L1CustomERC20Gateway());
        address aWithdrawLock  = address(new L1CustomERC20Gateway());
        address aReverseCust   = address(new L1CustomERC20Gateway());
        address aWETHGateway   = address(new L1WETHGateway(wethImpl != address(0) ? wethImpl : address(0), L2_WETH));
        address aERC721        = address(new L1ERC721Gateway());
        address aERC1155       = address(new L1ERC1155Gateway());
        address aEnforcedTx    = address(new EnforcedTxGateway());
        address aStaking       = address(new L1Staking(payable(messengerProxy)));
        address aSequencer     = address(new L1Sequencer());

        vm.stopBroadcast();

        // Write all addresses
        string[] memory n = new string[](17);
        address[] memory a = new address[](17);
        n[0]="Impl__Whitelist";                          a[0]=aWhitelist;
        n[1]=Types.IMPL_L1_CROSS_DOMAIN_MESSENGER;       a[1]=aMessenger;
        n[2]=Types.IMPL_L1_MESSAGE_QUEUE;                 a[2]=aMsgQueue;
        n[3]=Types.IMPL_ROLLUP;                            a[3]=aRollup;
        n[4]=Types.IMPL_MULTIPLE_VERSION_ROLLUP_VERIFIER; a[4]=aMultiVerifier;
        n[5]=Types.IMPL_L1_GATEWAY_ROUTER;                a[5]=aRouter;
        n[6]=Types.IMPL_L1_ETH_GATEWAY;                    a[6]=aETHGateway;
        n[7]=Types.IMPL_L1_STANDARD_ERC20_GATEWAY;        a[7]=aStdERC20;
        n[8]=Types.IMPL_L1_CUSTOM_ERC20_GATEWAY;          a[8]=aCustERC20;
        n[9]=Types.IMPL_L1_WITHDRAW_LOCK_ERC20_GATEWAY;   a[9]=aWithdrawLock;
        n[10]=Types.IMPL_L1_REVERSE_CUSTOM_GATEWAY;       a[10]=aReverseCust;
        n[11]=Types.IMPL_L1_WETH_GATEWAY;                  a[11]=aWETHGateway;
        n[12]=Types.IMPL_L1_ERC721_GATEWAY;                a[12]=aERC721;
        n[13]=Types.IMPL_L1_ERC1155_GATEWAY;               a[13]=aERC1155;
        n[14]=Types.IMPL_ENFORCED_TX_GATEWAY;              a[14]=aEnforcedTx;
        n[15]=Types.IMPL_L1_STAKING;                        a[15]=aStaking;
        n[16]=Types.IMPL_L1_SEQUENCER;                      a[16]=aSequencer;

        Storage.setMany(vm, outFile, n, a);
    }
}
