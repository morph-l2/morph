import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { storage } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const deployEmptyContract = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
): Promise<string> => {
    const EmptyContractFactoryName = ContractFactoryName.EmptyContract
    const EmptyContractImplStorageName = ImplStorageName.EmptyContract

    const Factory = await hre.ethers.getContractFactory(EmptyContractFactoryName)
    const contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", EmptyContractImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storage(path, EmptyContractImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    return ''
}

export default deployEmptyContract
