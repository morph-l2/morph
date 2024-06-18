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

    // L1MessageQueueWithGasPriceOracle config
    const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(path, ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName)
    const L1MessageQueueWithGasPriceOracleImplAddress = getContractAddressByName(path, ImplStorageName.L1MessageQueueWithGasPriceOracle)
    const L1MessageQueueWithGasPriceOracleFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1MessageQueueWithGasPriceOracle)

    // L1CrossDomainMessenger config
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)
    const L1CrossDomainMessengerImplAddress = getContractAddressByName(path, ImplStorageName.L1CrossDomainMessengerStorageName)
    const L1CrossDomainMessengerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1CrossDomainMessenger)

    const IL1CrossDomainMessengerProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1CrossDomainMessengerProxyAddress, deployer)
    // upgrade and initialize L1CrossDomainMessengerProxy
    if (
        ((await IL1CrossDomainMessengerProxy.implementation()).toLocaleLowerCase() !== L1CrossDomainMessengerImplAddress.toLocaleLowerCase()
        )) {
        console.log('Upgrading the L1CrossDomainMessenger proxy...')
        const l1FeeVaultRecipient: string = configTmp.l1FeeVaultRecipient

        if (!ethers.utils.isAddress(l1FeeVaultRecipient)
            || !ethers.utils.isAddress(RollupProxyAddress)
            || !ethers.utils.isAddress(L1MessageQueueWithGasPriceOracleProxyAddress)
        ) {
            console.error('upgrade l1CrossDomainMessenger failed !!! please check your params')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1CrossDomainMessengerProxy.upgradeToAndCall(
            L1CrossDomainMessengerImplAddress,
            L1CrossDomainMessengerFactory.interface.encodeFunctionData('initialize', [
                l1FeeVaultRecipient, RollupProxyAddress, L1MessageQueueWithGasPriceOracleProxyAddress
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await IL1CrossDomainMessengerProxy.implementation()).toLocaleLowerCase() === L1CrossDomainMessengerImplAddress.toLocaleLowerCase()
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
            L1MessageQueueWithGasPriceOracleProxyAddress
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

    const IL1MessageQueueWithGasPriceOracleProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1MessageQueueWithGasPriceOracleProxyAddress, deployer)

    if (
        (await IL1MessageQueueWithGasPriceOracleProxy.implementation()).toLocaleLowerCase() !== L1MessageQueueWithGasPriceOracleImplAddress.toLocaleLowerCase()
    ) {
        const maxGasLimit: number = configTmp.l1MessageQueueMaxGasLimit
        const whitelistAddress = getContractAddressByName(path, ImplStorageName.Whitelist)
        console.log('Upgrading the L1MessageQueueWithGasPriceOracle proxy...')
        if (
            maxGasLimit == 0 || !ethers.utils.isAddress(whitelistAddress)
        ) {
            console.error('upgrade L1MessageQueueWithGasPriceOracle failed !!! please check your params')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1MessageQueueWithGasPriceOracleProxy.upgradeToAndCall(
            L1MessageQueueWithGasPriceOracleImplAddress,
            L1MessageQueueWithGasPriceOracleFactory.interface.encodeFunctionData('initialize', [
                maxGasLimit,
                whitelistAddress
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await IL1MessageQueueWithGasPriceOracleProxy.implementation()).toLocaleLowerCase() === L1MessageQueueWithGasPriceOracleImplAddress.toLocaleLowerCase()
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
