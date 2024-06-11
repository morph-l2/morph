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
): Promise<string> => {

    const ZkEvmVerifierV1ContractFactoryName = ContractFactoryName.ZkEvmVerifierV1
    const implStorageName = ImplStorageName.ZkEvmVerifierV1StorageName
    const network = hre.network.name
    const bytecode = hexlify(fs.readFileSync(`./contracts/libraries/verifier/plonk-verifier/${network}/plonk_verifier_0.10.3.bin`));
    const tx = await deployer.sendTransaction({ data: bytecode });
    const receipt = await tx.wait();
    console.log("%s=%s ; TX_HASH: %s", "plonk_verifier.bin", receipt.contractAddress.toLocaleLowerCase(), tx.hash);

    const Factory = await hre.ethers.getContractFactory(ZkEvmVerifierV1ContractFactoryName)
    const contract = await Factory.deploy(receipt.contractAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", implStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(contract, 'PLONK_VERIFIER', receipt.contractAddress)
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storage(path, implStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    return ''
}

export default deployZkEvmVerifierV1

