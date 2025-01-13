import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ContractFactoryName } from "../src/types";

// yarn hardhat upgradeVerifier --rollupversion 0 --startbatchindex 0 --multipleversionrollupverifier 0x0165878a594ca255338adfa4d48449f69242eb8f --network l1
task("upgradeVerifier")
    .addParam("rollupversion")
    .addParam("startbatchindex")
    .addParam("multipleversionrollupverifier")
    .setAction(async (taskArgs, hre) => {
        const config = hre.deployConfig

        // deploy ZkEvmVerifierV1
        const ZkEvmVerifierV1ContractFactoryName = ContractFactoryName.ZkEvmVerifierV1
        const Factory = await hre.ethers.getContractFactory(ZkEvmVerifierV1ContractFactoryName)
        const contract = await Factory.deploy(config.programVkey)
        await contract.deployed()
        console.log("ZkEvmVerifierV1Contract: %s ; TX_HASH: %s", contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)

        // add verifier to MultipleVersionRollupVerifier
        const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
        const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(MultipleVersionRollupVerifierFactoryName)
        const MultipleVersionRollupVerifier = MultipleVersionRollupVerifierFactory.attach(taskArgs.multipleversionrollupverifier)

        const res = await MultipleVersionRollupVerifier.updateVerifier(taskArgs.rollupversion, taskArgs.startbatchindex, contract.address.toLocaleLowerCase())

        const receipt = await res.wait()
        console.log(`receipt status : ${receipt.status}`)
        console.log("upgrade verifier successfully, verifier: %s, rollupVersion: %s, startBatchIndex: %s", contract.address.toLocaleLowerCase(), taskArgs.rollupVersion, taskArgs.startBatchIndex)
    });