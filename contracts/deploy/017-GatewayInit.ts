import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { ethers } from 'ethers'
import { predeploys } from '../src/constants'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const GatewayInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const ProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)

    // L1GatewayRouter config
    const L1GatewayRouterProxyAddress = getContractAddressByName(path, ProxyStorageName.L1GatewayRouterProxyStorageName)
    const L1GatewayRouterImplAddress = getContractAddressByName(path, ImplStorageName.L1GatewayRouterStorageName)
    const L1GatewayRouterFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1GatewayRouter)

    // L1ETHGateway config
    const L1ETHGatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1ETHGatewayProxyStorageName)
    const L1ETHGatewayImplAddress = getContractAddressByName(path, ImplStorageName.L1ETHGatewayStorageName)
    const L1ETHGatewayFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1ETHGateway)

    // L1StandardERC20Gateway config
    const L1StandardERC20GatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StandardERC20GatewayProxyStorageName)
    const L1StandardERC20GatewayImplAddress = getContractAddressByName(path, ImplStorageName.L1StandardERC20GatewayStorageName)
    const L1StandardERC20GatewayFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1StandardERC20Gateway)

    // L1ERC721Gateway config
    const L1ERC721GatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1ERC721GatewayProxyStorageName)
    const L1ERC721GatewayImplAddress = getContractAddressByName(path, ImplStorageName.L1ERC721GatewayStorageName)
    const L1ERC721GatewayFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1ERC721Gateway)

    // L1ERC1155Gateway config
    const L1ERC1155GatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1ERC1155GatewayProxyStorageName)
    const L1ERC1155GatewayImplAddress = getContractAddressByName(path, ImplStorageName.L1ERC1155GatewayStorageName)
    const L1ERC1155GatewayFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1ERC1155Gateway)

    // L1WETHGateway config
    const L1WETHGatewayProxyAddress = getContractAddressByName(path, ProxyStorageName.L1WETHGatewayProxyStorageName)
    const L1WETHGatewayImplAddress = getContractAddressByName(path, ImplStorageName.L1WETHGatewayStorageName)
    const L1WETHGatewayFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1WETHGateway)

    const WETHAddress = getContractAddressByName(path, ImplStorageName.WETH)

    // L1GatewayRouter init
    const IL1GatewayRouterProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1GatewayRouterProxyAddress, deployer)
    if (
        (await IL1GatewayRouterProxy.implementation()).toLocaleLowerCase() !== L1GatewayRouterImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1GatewayRouter proxy...')
        if (!ethers.utils.isAddress(L1ETHGatewayProxyAddress)
            || !ethers.utils.isAddress(L1StandardERC20GatewayProxyAddress)
            || !ethers.utils.isAddress(WETHAddress)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1GatewayRouterProxy.connect(deployer).upgradeToAndCall(
            L1GatewayRouterImplAddress,
            L1GatewayRouterFactory.interface.encodeFunctionData('initialize', [
                L1ETHGatewayProxyAddress,
                L1StandardERC20GatewayProxyAddress
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1GatewayRouterProxy.implementation()).toLocaleLowerCase() === L1GatewayRouterImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1GatewayRouterProxyAddress,
            L1GatewayRouterFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'ethGateway',
            L1ETHGatewayProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'defaultERC20Gateway',
            L1StandardERC20GatewayProxyAddress
        )
        console.log('L1GatewayRouter upgrade success')
    }

    // L1ETHGateway init
    const IL1ETHGatewayProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1ETHGatewayProxyAddress, deployer)
    if (
        (await IL1ETHGatewayProxy.implementation()).toLocaleLowerCase() !== L1ETHGatewayImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1ETHGateway proxy...')
        const counterpart: string = predeploys.L2ETHGateway

        if (!ethers.utils.isAddress(counterpart)
            || !ethers.utils.isAddress(L1GatewayRouterProxyAddress)
            || !ethers.utils.isAddress(L1CrossDomainMessengerProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1ETHGatewayProxy.connect(deployer).upgradeToAndCall(
            L1ETHGatewayImplAddress,
            L1ETHGatewayFactory.interface.encodeFunctionData('initialize', [
                counterpart,
                L1GatewayRouterProxyAddress,
                L1CrossDomainMessengerProxyAddress
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1ETHGatewayProxy.implementation()).toLocaleLowerCase() === L1ETHGatewayImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1ETHGatewayProxyAddress,
            L1ETHGatewayFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            counterpart
        )
        await assertContractVariable(
            contractTmp,
            'router',
            L1GatewayRouterProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'messenger',
            L1CrossDomainMessengerProxyAddress
        )
        console.log('L1ETHGatewayProxy upgrade success')
    }

    // L1StandardERC20Gateway init
    const IL1StandardERC20GatewayProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1StandardERC20GatewayProxyAddress, deployer)
    if (
        (await IL1StandardERC20GatewayProxy.implementation()).toLocaleLowerCase() !== L1StandardERC20GatewayImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1StandardERC20Gateway proxy...')
        const counterpart: string = predeploys.L2StandardERC20Gateway
        const l2TokenImplementation: string = predeploys.MorphStandardERC20
        const l2TokenFactory: string = predeploys.MorphStandardERC20Factory

        if (!ethers.utils.isAddress(counterpart)
            || !ethers.utils.isAddress(L1GatewayRouterProxyAddress)
            || !ethers.utils.isAddress(l2TokenImplementation)
            || !ethers.utils.isAddress(l2TokenFactory)
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await IL1StandardERC20GatewayProxy.connect(deployer).upgradeToAndCall(
            L1StandardERC20GatewayImplAddress,
            L1StandardERC20GatewayFactory.interface.encodeFunctionData('initialize', [
                counterpart,
                L1GatewayRouterProxyAddress,
                L1CrossDomainMessengerProxyAddress,
                l2TokenImplementation,
                l2TokenFactory
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1StandardERC20GatewayProxy.implementation()).toLocaleLowerCase() === L1StandardERC20GatewayImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1StandardERC20GatewayProxyAddress,
            L1StandardERC20GatewayFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            counterpart
        )
        await assertContractVariable(
            contractTmp,
            'router',
            L1GatewayRouterProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'messenger',
            L1CrossDomainMessengerProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'l2TokenImplementation',
            l2TokenImplementation
        )
        await assertContractVariable(
            contractTmp,
            'l2TokenFactory',
            l2TokenFactory
        )
        console.log('L1StandardERC20Gateway upgrade success')
    }


    // L1ERC721Gateway init
    const IL1ERC721GatewayProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1ERC721GatewayProxyAddress, deployer)
    if (
        (await IL1ERC721GatewayProxy.implementation()).toLocaleLowerCase() !== L1ERC721GatewayImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1ERC721Gateway proxy...')
        const counterpart: string = predeploys.L2ERC721Gateway

        if (!ethers.utils.isAddress(counterpart)
            || !ethers.utils.isAddress(L1GatewayRouterProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await IL1ERC721GatewayProxy.upgradeToAndCall(
            L1ERC721GatewayImplAddress,
            L1ERC721GatewayFactory.interface.encodeFunctionData('initialize', [
                counterpart,
                L1CrossDomainMessengerProxyAddress,
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1ERC721GatewayProxy.implementation()).toLocaleLowerCase() === L1ERC721GatewayImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1ERC721GatewayProxyAddress,
            L1ERC721GatewayFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            counterpart
        )
        await assertContractVariable(
            contractTmp,
            'messenger',
            L1CrossDomainMessengerProxyAddress
        )
        console.log('L1ERC721Gateway upgrade success')
    }

    // L1ERC1155Gateway init
    const IL1ERC1155GatewayProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1ERC1155GatewayProxyAddress, deployer)
    if (
        (await IL1ERC1155GatewayProxy.implementation()).toLocaleLowerCase() !== L1ERC1155GatewayImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1ERC1155Gateway proxy...')
        const counterpart: string = predeploys.L2ERC1155Gateway

        if (!ethers.utils.isAddress(counterpart)
            || !ethers.utils.isAddress(L1GatewayRouterProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await IL1ERC1155GatewayProxy.upgradeToAndCall(
            L1ERC1155GatewayImplAddress,
            L1ERC1155GatewayFactory.interface.encodeFunctionData('initialize', [
                counterpart,
                L1CrossDomainMessengerProxyAddress,
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1ERC1155GatewayProxy.implementation()).toLocaleLowerCase() === L1ERC1155GatewayImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1ERC1155GatewayProxyAddress,
            L1ERC1155GatewayFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            counterpart
        )
        await assertContractVariable(
            contractTmp,
            'messenger',
            L1CrossDomainMessengerProxyAddress
        )
        console.log('L1ERC1155Gateway upgrade success')
    }

    // L1WETHGateway init
    const IL1WETHGatewayProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1WETHGatewayProxyAddress, deployer)
    if (
        (await IL1WETHGatewayProxy.implementation()).toLocaleLowerCase() !== L1WETHGatewayImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1WETHGateway proxy...')
        const counterpart: string = predeploys.L2WETHGateway

        if (!ethers.utils.isAddress(counterpart)
            || !ethers.utils.isAddress(L1GatewayRouterProxyAddress)
            || !ethers.utils.isAddress(L1CrossDomainMessengerProxyAddress)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IL1WETHGatewayProxy.connect(deployer).upgradeToAndCall(
            L1WETHGatewayImplAddress,
            L1WETHGatewayFactory.interface.encodeFunctionData('initialize', [
                counterpart,
                L1GatewayRouterProxyAddress,
                L1CrossDomainMessengerProxyAddress
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await IL1WETHGatewayProxy.implementation()).toLocaleLowerCase() === L1WETHGatewayImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1WETHGatewayProxyAddress,
            L1WETHGatewayFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'counterpart',
            counterpart
        )
        await assertContractVariable(
            contractTmp,
            'router',
            L1GatewayRouterProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'messenger',
            L1CrossDomainMessengerProxyAddress
        )
        console.log('L1WETHGatewayProxy upgrade success')
    }
    return ''
}

export default GatewayInit
