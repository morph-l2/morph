import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import fs from "fs";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storage } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"
import { hexlify } from "ethers/lib/utils";

export const deployZkEvmVerifierV1 = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any,
): Promise<string> => {
    const ZkEvmVerifierV1ContractFactoryName = ContractFactoryName.ZkEvmVerifierV1
    const implStorageName = ImplStorageName.ZkEvmVerifierV1StorageName

    const Factory = await hre.ethers.getContractFactory(ZkEvmVerifierV1ContractFactoryName)
    const contract = await Factory.deploy(config.programVkey)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", implStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storage(path, implStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    return ''
}

export default deployZkEvmVerifierV1

