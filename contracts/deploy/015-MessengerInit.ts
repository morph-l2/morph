import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { BigNumber, ethers } from 'ethers'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"
import { predeploys } from "../src";

export const MessengerInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const ProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)

    // L1MessageQueue config
    const L1MessageQueueProxyAddress = getContractAddressByName(path, ProxyStorageName.L1MessageQueueProxyStroageName)
    const L1MessageQueueImplAddress = getContractAddressByName(path, ImplStorageName.L1MessageQueueStroageName)
    const L1MessageQueueFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1MessageQueue)

    // L1CrossDomainMessenge config
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
    const L1CrossDomainMessengerImplAddress = getContractAddressByName(path, ImplStorageName.L1CrossDomainMessengerStorageName)
    const L1CrossDomainMessengerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1CrossDomainMessenger)

    // L2GasPriceOracle config
    const L2GasPriceOracleProxyAddress = getContractAddressByName(path, ProxyStorageName.L2GasPriceOracleProxyStorageName)
    const L2GasPriceOracleImplAddress = getContractAddressByName(path, ImplStorageName.L2GasPriceOracleStorageName)
    const L2GasPriceOracleFactory = await hre.ethers.getContractFactory(ContractFactoryName.L2GasPriceOracle)

    const L1CrossDomainMessengerProxy = new ethers.Contract(
        L1CrossDomainMessengerProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )

    // upgrade and initialize L1CrossDomainMessengerProxy
    if (
        (await L1CrossDomainMessengerProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== L1CrossDomainMessengerImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1CrossDomainMessenger proxy...')
        const l1FeeVaultRecipient: string = configTmp.l1FeeVaultRecipient

        if (!ethers.utils.isAddress(l1FeeVaultRecipient)
            || !ethers.utils.isAddress(RollupProxyAddress)
            || !ethers.utils.isAddress(L1MessageQueueProxyAddress)
        ) {
            console.error('upgrade l1CrossDomainMessenger failed !!! please check your params')
            return ''
        }
        // Upgrade and initialize the proxy.
        await L1CrossDomainMessengerProxy.connect(deployer).upgradeToAndCall(
            L1CrossDomainMessengerImplAddress,
            L1CrossDomainMessengerFactory.interface.encodeFunctionData('initialize', [
                l1FeeVaultRecipient, RollupProxyAddress, L1MessageQueueProxyAddress
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await L1CrossDomainMessengerProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === L1CrossDomainMessengerImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        // params check
        const contractTmp = new ethers.Contract(
            L1CrossDomainMessengerProxyAddress,
            L1CrossDomainMessengerFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'rollup',
            RollupProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'messageQueue',
            L1MessageQueueProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            predeploys.L2CrossDomainMessenger
        )
        await assertContractVariable(
            contractTmp,
            'feeVault',
            l1FeeVaultRecipient
        )
        await assertContractVariable(
            contractTmp,
            'xDomainMessageSender',
            '0x000000000000000000000000000000000000dEaD'
        )
        await assertContractVariable(
            contractTmp,
            'maxReplayTimes',
            3
        )
        // Wait for the transaction to execute properly.
        console.log('L1CrossDomainMessengerProxy upgrade success')
    }

    const L2GasPriceOracleProxy = new ethers.Contract(
        L2GasPriceOracleProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )
    if (
        (await L2GasPriceOracleProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== L2GasPriceOracleImplAddress.toLocaleLowerCase()
    ) {
        const txGas: number = configTmp.gasPriceOracleTxGas
        const txGasContractCreation: number = configTmp.gasPriceOracleTxGasContractCreation
        const zeroGas: number = configTmp.gasPriceOracleZeroGas
        const nonZeroGas: number = configTmp.gasPriceOracleNonZeroGas
        console.log('Upgrading the L2GasPriceOracle proxy...')
        if (txGas == 0
            || txGasContractCreation == 0
            || zeroGas == 0
            || nonZeroGas == 0
            || txGasContractCreation <= txGas) {
            console.error('upgrade L2GasPriceOracle failed !!! please check your params')
            return ''
        }

        // Upgrade and initialize the proxy.
        await L2GasPriceOracleProxy.connect(deployer).upgradeToAndCall(
            L2GasPriceOracleImplAddress,
            L2GasPriceOracleFactory.interface.encodeFunctionData('initialize', [
                txGas, txGasContractCreation, zeroGas, nonZeroGas
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await L2GasPriceOracleProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === L2GasPriceOracleImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        // Wait for the transaction to execute properly.
        console.log('L2GasPriceOracleProxy upgrade success')
    }

    const L1MessageQueueProxy = new ethers.Contract(
        L1MessageQueueProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )

    if (
        (await L1MessageQueueProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== L1MessageQueueImplAddress.toLocaleLowerCase()
    ) {
        const maxGasLimit: number = configTmp.l1MessageQueueMaxGasLimit

        console.log('Upgrading the L1MessageQueue proxy...')
        if (maxGasLimit == 0) {
            console.error('upgrade L1MessageQueue failed !!! please check your params')
            return ''
        }
        // Upgrade and initialize the proxy.
        await L1MessageQueueProxy.connect(deployer).upgradeToAndCall(
            L1MessageQueueImplAddress,
            L1MessageQueueFactory.interface.encodeFunctionData('initialize', [
                L1CrossDomainMessengerProxyAddress,
                RollupProxyAddress,
                hre.ethers.constants.AddressZero,
                L2GasPriceOracleProxyAddress,
                maxGasLimit
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await L1MessageQueueProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === L1MessageQueueImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        // Wait for the transaction to execute properly.
        console.log('L1MessageQueueProxy upgrade success')
    }
    return ''
}

export default MessengerInit
