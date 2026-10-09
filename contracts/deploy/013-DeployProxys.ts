import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";
import { getContractAddressByName, storage } from "../src/deploy-utils";
import { deployRecordedContract } from "../src/deployment-state";
import { ContractFactoryName, ImplStorageName, ProxyStorageName } from "../src/types";

export const deployContractProxyByStorageName = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, storageName: string
): Promise<string> => {
    await deployRecordedContract(hre, path, deployer, storageName, ContractFactoryName.DefaultProxy, [
        getContractAddressByName(path, ImplStorageName.EmptyContract), await deployer.getAddress(), "0x",
    ]);
    return "";
};

export const deployContractProxies = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, config: any
): Promise<string> => {
    if (config.l1WETHAddress === "") {
        await deployRecordedContract(hre, path, deployer, ImplStorageName.WETH, ContractFactoryName.WETH);
    } else {
        if (!ethers.utils.isAddress(config.l1WETHAddress) || await hre.ethers.provider.getCode(config.l1WETHAddress) === "0x") {
            throw new Error("l1WETHAddress must be a deployed contract address on the current network");
        }
        await storage(path, ImplStorageName.WETH, config.l1WETHAddress, await hre.ethers.provider.getBlockNumber());
    }
    for (const name of Object.values(ProxyStorageName)) {
        await deployContractProxyByStorageName(hre, path, deployer, name);
    }
    return "";
};

// Keep the existing entry point; confirm and record each transaction, and deploy only missing contracts on retries.
export const deployContractProxiesConcurrently = deployContractProxies;
export default deployContractProxies;
