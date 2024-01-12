import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const deployProxyAdmin = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
): Promise<string> => {
    const ProxyAdminContractFactoryName = ContractFactoryName.ProxyAdmin
    const ImplStroageName = ImplStorageName.ProxyAdmin

    const Factory = await hre.ethers.getContractFactory(ProxyAdminContractFactoryName)
    const contract = await Factory.deploy(await deployer.getAddress())
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", ImplStroageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params then storge
    await assertContractVariable(contract, 'owner', await deployer.getAddress())
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, ImplStroageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    return ''
}

export default deployProxyAdmin
