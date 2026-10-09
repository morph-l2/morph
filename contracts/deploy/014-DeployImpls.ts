import { HardhatRuntimeEnvironment } from "hardhat/types";
import { assertContractVariable, getContractAddressByName } from "../src/deploy-utils";
import { deployRecordedContract } from "../src/deployment-state";
import { predeploys } from "../src/constants";
import { ContractFactoryName as F, ImplStorageName as I, ProxyStorageName as P } from "../src/types";

export const deployContractImpls = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, config: any
): Promise<string> => {
    const address = (name: string) => getContractAddressByName(path, name);
    const contracts: [string, string, any[]][] = [
        [I.Whitelist, F.Whitelist, [config.contractAdmin]],
        [I.L1CrossDomainMessengerStorageName, F.L1CrossDomainMessenger, []],
        [I.L1MessageQueueWithGasPriceOracle, F.L1MessageQueueWithGasPriceOracle, [address(P.L1CrossDomainMessengerProxyStorageName), address(P.RollupProxyStorageName), address(P.EnforcedTxGatewayProxyStorageName)]],
        [I.RollupStorageName, F.Rollup, [config.l2ChainID]],
        [I.L1GatewayRouterStorageName, F.L1GatewayRouter, []],
        [I.L1StandardERC20GatewayStorageName, F.L1StandardERC20Gateway, []],
        [I.L1CustomERC20GatewayStorageName, F.L1CustomERC20Gateway, []],
        [I.L1WithdrawLockERC20GatewayStorageName, F.L1WithdrawLockERC20Gateway, []],
        [I.L1ReverseCustomGatewayStorageName, F.L1ReverseCustomGateway, []],
        [I.L1ETHGatewayStorageName, F.L1ETHGateway, []],
        [I.L1WETHGatewayStorageName, F.L1WETHGateway, [address(I.WETH), predeploys.L2WETH]],
        [I.EnforcedTxGatewayStorageName, F.EnforcedTxGateway, []],
        [I.L1ERC721GatewayStorageName, F.L1ERC721Gateway, []],
        [I.L1ERC1155GatewayStorageName, F.L1ERC1155Gateway, []],
        [I.SubmitterStorageName, F.Submitter, []],
        [I.L1SequencerStorageName, F.L1Sequencer, []],
    ];
    for (const [name, factory, args] of contracts) {
        const contract = await deployRecordedContract(hre, path, deployer, name, factory, args);
        if (name === I.RollupStorageName) await assertContractVariable(contract, "LAYER_2_CHAIN_ID", config.l2ChainID);
        if (name === I.Whitelist) await assertContractVariable(contract, "owner", config.contractAdmin);
    }
    return "";
};

// Keep the concurrent entry point, confirming and recording transactions sequentially for resumable deployment.
export const deployContractImplsConcurrently = deployContractImpls;
export default deployContractImpls;
