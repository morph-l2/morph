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
        
        // Get initial sequencer address from config (first sequencer address)
        // Note: l2SequencerAddresses is defined in contracts/src/deploy-config/l1.ts
        const initialSequencer = (configTmp.l2SequencerAddresses && configTmp.l2SequencerAddresses.length > 0)
            ? configTmp.l2SequencerAddresses[0]
            : ethers.constants.AddressZero
        
        console.log('Initial sequencer address:', initialSequencer)

        // Upgrade and initialize the proxy with owner and initial sequencer
        // Note: We set sequencer in initialize() to avoid TransparentUpgradeableProxy admin restriction
        await IL1SequencerProxy.upgradeToAndCall(
            L1SequencerImplAddress,
            L1SequencerFactory.interface.encodeFunctionData('initialize', [owner, initialSequencer])
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

        if (initialSequencer !== ethers.constants.AddressZero) {
            await assertContractVariable(
                contractTmp,
                'sequencer',
                initialSequencer,
            )
            console.log('L1SequencerProxy upgrade success, initial sequencer set:', initialSequencer)
        } else {
            console.log('L1SequencerProxy upgrade success (no initial sequencer set)')
        }
    }
    return ''
}

export default SequencerInit
