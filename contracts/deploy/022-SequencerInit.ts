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
        // The first sequencer is registered separately via setFirstSequencer(address).
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

// SetFirstSequencer must run *after* the proxy admin has been handed over to
// ProxyAdmin. While the deployer is still the proxy admin, any call that falls
// through to the implementation reverts with
// "TransparentUpgradeableProxy: admin cannot fallback to proxy target".
export const SetFirstSequencer = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStorageName)
    const L1SequencerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Sequencer)

    const L1Sequencer = new ethers.Contract(
        L1SequencerProxyAddress,
        L1SequencerFactory.interface,
        deployer,
    )

    // deployConfig is a Proxy that throws on unknown keys, so probe with `in`
    // before reading.
    const firstSequencer: string = process.env.firstSequencerAddress ||
        ('firstSequencerAddress' in configTmp ? configTmp.firstSequencerAddress : '')
    if (!firstSequencer) {
        // Networks that register the first sequencer out-of-band (devnet does it
        // from ops/devnet-morph) leave this unset.
        console.log('firstSequencerAddress not configured, skipping setFirstSequencer')
        return ''
    }
    if (!ethers.utils.isAddress(firstSequencer)) {
        return `invalid firstSequencerAddress: ${firstSequencer}`
    }

    if ((await L1Sequencer.getSequencerHistoryLength()).gt(0)) {
        console.log('First sequencer already set:', await L1Sequencer.getSequencer())
        return ''
    }

    console.log('Setting first sequencer:', firstSequencer)
    const tx = await L1Sequencer.setFirstSequencer(firstSequencer)
    await tx.wait()

    await assertContractVariable(
        L1Sequencer,
        'getSequencer',
        firstSequencer,
    )

    console.log('setFirstSequencer success')
    return ''
}

export default SequencerInit
