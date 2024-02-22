import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, assertContractVariableWithSigner, awaitCondition, storge } from "../src/deploy-utils";
import { ethers } from 'ethers'
import { predeploys } from '../src/constants'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const ContractInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log("ContractInit")
    // ------------------ rollup init -----------------
    {
        const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
        const Rollup = await hre.ethers.getContractAt(ContractFactoryName.Rollup, RollupProxyAddress, deployer)
        // import genesis batch 
        const genesisStateRoot: string = config.rollupGenesisStateRoot
        const withdrawRoot: string = config.withdrawRoot
        const batchHeader: string = config.batchHeader
        // submitter and challenger
        const submitter: string = config.rollupProposer
        const challenger: string = config.rollupChallenger
        if (!ethers.utils.isAddress(submitter)
            || !ethers.utils.isAddress(challenger)
        ) {
            console.error('please check your address')
            return ''
        }
        console.log('importGenesisBatch(%s, %s, %s)', batchHeader, genesisStateRoot, withdrawRoot)
        await Rollup.importGenesisBatch(batchHeader, genesisStateRoot, withdrawRoot)
        console.log('addProver(%s)', submitter)
        await Rollup.addProver(submitter)
        console.log('addChallenger(%s)', challenger)
        await Rollup.addChallenger(challenger)
    }
    // ------------------ staking init -----------------
    {
        const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStroageName)
        const Staking = await hre.ethers.getContractAt(ContractFactoryName.Staking, StakingProxyAddress, deployer)
        const whiteListAdd = config.l2SequencerAddresses
        // set sequencer to white list
        await Staking.updateWhitelist(whiteListAdd, [])
        for (let i = 0; i < config.l2SequencerAddresses.length; i++) {
            // Wait for the transaction to execute properly.
            await awaitCondition(
                async () => {
                    return (
                        await Staking.whitelist(config.l2SequencerAddresses[i]) === true
                    )
                },
                3000,
                1000
            )
            console.log(`address ${config.l2SequencerAddresses[i]} is in white list`)
        }
    }
    // return nil
    return ''
}

export default ContractInit
