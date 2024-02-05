import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition, storge } from "../src/deploy-utils";
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
    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStroageName)
    const L1SequencerImplAddress = getContractAddressByName(path, ImplStorageName.L1SequencerStorageName)
    const L1SequencerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Sequencer)

    // Staking config
    const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStroageName)
    const StakingImplAddress = getContractAddressByName(path, ImplStorageName.StakingStorageName)
    const StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.Staking)

    const L1SequencerProxy = new ethers.Contract(
        L1SequencerProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )
    if (
        (await L1SequencerProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== L1SequencerImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1Sequencer proxy...')
        if (!ethers.utils.isAddress(RollupProxyAddress)
            || !ethers.utils.isAddress(StakingProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await L1SequencerProxy.connect(deployer).upgradeToAndCall(
            L1SequencerImplAddress,
            L1SequencerFactory.interface.encodeFunctionData('initialize', [
                StakingProxyAddress,
                RollupProxyAddress
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await L1SequencerProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === L1SequencerImplAddress.toLocaleLowerCase()
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


    const StakingProxy = new ethers.Contract(
        StakingProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )
    if (
        (await StakingProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== StakingImplAddress.toLocaleLowerCase()
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
        await StakingProxy.connect(deployer).upgradeToAndCall(
            StakingImplAddress,
            StakingFactory.interface.encodeFunctionData('initialize', [
                admin,
                L1SequencerProxyAddress,
                sequencerSize,
                limit,
                hre.ethers.utils.parseEther(lock.toString()),
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await StakingProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === StakingImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        const whiteListAdd = configTmp.l2SequencerAddresses

        const contractTmp = new ethers.Contract(
            StakingProxyAddress,
            StakingFactory.interface,
            deployer,
        )
        // set sequencer to white list
        await contractTmp.updateWhitelist(whiteListAdd, [])
        for (let i = 0; i < configTmp.l2SequencerAddresses.length; i++) {
            // Wait for the transaction to execute properly.
            await awaitCondition(
                async () => {
                    return (
                        await contractTmp.whitelist(configTmp.l2SequencerAddresses[i]) === true
                    )
                },
                3000,
                1000
            )
            console.log(`address ${configTmp.l2SequencerAddresses[i]} is in white list`)
        }

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
            limit
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
