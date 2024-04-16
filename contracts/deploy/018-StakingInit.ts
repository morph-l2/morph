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
} from "./types"

export const StakingInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const ProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)

    // Sequencer config
    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStorageName)
    const L1SequencerImplAddress = getContractAddressByName(path, ImplStorageName.L1SequencerStorageName)
    const L1SequencerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Sequencer)

    // Staking config
    const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStorageName)
    const StakingImplAddress = getContractAddressByName(path, ImplStorageName.StakingStorageName)
    const StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.Staking)

    const IL1SequencerProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1SequencerProxyAddress, deployer)
    if (
        (await IL1SequencerProxy.implementation()).toLocaleLowerCase() !== L1SequencerImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1Sequencer proxy...')
        if (!ethers.utils.isAddress(RollupProxyAddress)
            || !ethers.utils.isAddress(StakingProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1SequencerProxy.upgradeToAndCall(
            L1SequencerImplAddress,
            L1SequencerFactory.interface.encodeFunctionData('initialize', [
                StakingProxyAddress,
                RollupProxyAddress
            ])
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
            'rollupContract',
            RollupProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'stakingContract',
            StakingProxyAddress
        )
        console.log('L1SequencerProxy upgrade success')
    }

    const IStakingProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, StakingProxyAddress, deployer)
    if (
        (await IStakingProxy.implementation()).toLocaleLowerCase() !== StakingImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Staking proxy...')
        const admin: string = configTmp.contractAdmin
        const sequencerSize: number = configTmp.stakingSequencerSize
        const limit: number = configTmp.stakingMinDeposit
        const lock: number = configTmp.stakingLockNumber

        if (!ethers.utils.isAddress(admin)
            || !ethers.utils.isAddress(L1SequencerProxyAddress)
            || sequencerSize == 0
            || lock == 0
            || limit == 0
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await IStakingProxy.upgradeToAndCall(
            StakingImplAddress,
            StakingFactory.interface.encodeFunctionData('initialize', [
                admin,
                L1SequencerProxyAddress,
                sequencerSize,
                hre.ethers.utils.parseEther(limit.toString()),
                hre.ethers.utils.parseEther(lock.toString()),
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IStakingProxy.implementation()).toLocaleLowerCase() === StakingImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        const contractTmp = new ethers.Contract(
            StakingProxyAddress,
            StakingFactory.interface,
            deployer,
        )

        await assertContractVariable(
            contractTmp,
            'sequencerContract',
            L1SequencerProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'sequencersSize',
            sequencerSize
        )
        await assertContractVariable(
            contractTmp,
            'limit',
            hre.ethers.utils.parseEther(limit.toString()),
        )
        await assertContractVariable(
            contractTmp,
            'lock',
            hre.ethers.utils.parseEther(lock.toString()),
        )
        await assertContractVariable(
            contractTmp,
            'owner',
            await deployer.getAddress(),
        )
        // Wait for the transaction to execute properly.
        console.log('StakingProxy upgrade success')
    }
    return ''
}

export default StakingInit
