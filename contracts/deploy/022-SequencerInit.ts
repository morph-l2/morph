import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { ethers } from 'ethers'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const SequencerInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // L1Sequencer addresses
    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStorageName)
    const L1SequencerImplAddress = getContractAddressByName(path, ImplStorageName.L1SequencerStorageName)
    const L1SequencerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Sequencer)

    const IL1SequencerProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1SequencerProxyAddress, deployer)
    
    if (
        (await IL1SequencerProxy.implementation()).toLocaleLowerCase() !== L1SequencerImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1Sequencer proxy...')
        
        // Owner is the deployer (will be transferred to multisig in production)
        const owner = await deployer.getAddress()

        // Upgrade and initialize the proxy with owner only.
        // Sequencer history is initialized separately via initializeHistory().
        await IL1SequencerProxy.upgradeToAndCall(
            L1SequencerImplAddress,
            L1SequencerFactory.interface.encodeFunctionData('initialize', [owner])
        )

        await awaitCondition(
            async () => {
                return (
                    (await IL1SequencerProxy.implementation()).toLocaleLowerCase() === L1SequencerImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        const contractTmp = new ethers.Contract(
            L1SequencerProxyAddress,
            L1SequencerFactory.interface,
            deployer,
        )

        await assertContractVariable(
            contractTmp,
            'owner',
            owner,
        )

        console.log('L1SequencerProxy upgrade success')
    }
    return ''
}

export default SequencerInit
