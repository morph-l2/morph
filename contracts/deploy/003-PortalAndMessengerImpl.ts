import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { ethers } from 'ethers'
import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName
} from "./types"

export const deployPortalAndMessengerImpl = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    // config get
    const MorphPortalProxyStorageName = ProxyStorageName.MorphPortalProxyStroageName
    const MorphPortalImplStorageName = ImplStorageName.MorphPortalStroageName
    const L1CrossDomainMessengerProxyStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStroageName
    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const L1CrossDomainMessengerContractFactoryName = ContractFactoryName.L1CrossDomainMessenger
    const DefaultProxyProxyFactoy = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)

    // get init params
    const MorphPortalProxyAddress = getContractAddressByName(path, MorphPortalProxyStorageName)
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, L1CrossDomainMessengerProxyStorageName)

    const MorphPortalProxy = new ethers.Contract(
        MorphPortalProxyAddress,
        DefaultProxyProxyFactoy.interface,
        deployer
    )
    const L1CrossDomainMessengerProxy = new ethers.Contract(
        L1CrossDomainMessengerProxyAddress,
        DefaultProxyProxyFactoy.interface,
        deployer
    )
    const portalGuardian = config.portalGuardian
    const portalGuardianCode = await hre.ethers.provider.getCode(portalGuardian)
    if (portalGuardianCode === '0x') {
        console.log(
            `WARNING: setting MorphPortal.GUARDIAN to ${portalGuardian} and it has no code`
        )
    }
    const Artifact__SystemConfigProxyAddress = getContractAddressByName(path, ProxyStorageName.SystemConfigProxyStorageName)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)

    console.log('\n---------------------------------- deploy MorphPortal impl ----------------------------------')
    // contract deploy
    const MorphPortalFactory = await hre.ethers.getContractFactory(ContractFactoryName.MorphPortal)
    const morphPortal = await MorphPortalFactory.deploy(
        portalGuardian,
        true, // paused
        Artifact__SystemConfigProxyAddress,
        RollupProxyAddress,
        L1CrossDomainMessengerProxyAddress
    )
    await morphPortal.deployed()
    console.log("%s=%s ; TX_HASH: %s", MorphPortalImplStorageName, morphPortal.address.toLocaleLowerCase(), morphPortal.deployTransaction.hash);

    // check params then storge
    await assertContractVariable(
        morphPortal,
        'GUARDIAN',
        portalGuardian
    )
    await assertContractVariable(
        morphPortal,
        'SYSTEM_CONFIG',
        Artifact__SystemConfigProxyAddress
    )
    await assertContractVariable(
        morphPortal,
        'l1Messenger',
        L1CrossDomainMessengerProxyAddress
    )
    let blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, MorphPortalImplStorageName, morphPortal.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // proxy upgradeAndCall
    console.log('Upgrading the MorphPortal proxy...')
    await MorphPortalProxy.upgradeToAndCall(
        morphPortal.address,
        morphPortal.interface.encodeFunctionData('initialize', [
            false, // not paused
            L1CrossDomainMessengerProxyAddress
        ])
    )
    // Wait for the transaction to execute properly.
    await awaitCondition(
        async () => {
            const temp = new ethers.Contract(
                MorphPortalProxy.address,
                MorphPortalProxy.interface,
                MorphPortalProxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: ethers.constants.AddressZero,
            })
            return (
                actual.toLocaleLowerCase() === morphPortal.address.toLocaleLowerCase()
            )
        },
        30000,
        1000
    )
    // check params
    let checkContract = new ethers.Contract(
        MorphPortalProxy.address,
        morphPortal.interface,
        MorphPortalProxy.provider
    )
    await assertContractVariable(
        checkContract,
        'GUARDIAN',
        portalGuardian
    )
    await assertContractVariable(
        checkContract,
        'SYSTEM_CONFIG',
        Artifact__SystemConfigProxyAddress
    )
    await assertContractVariable(
        checkContract,
        'l1Messenger',
        L1CrossDomainMessengerProxyAddress
    )
    console.log('Upgrading the MorphPortal proxy success...')

    console.log('\n---------------------------------- deploy L1CrossDomainMessenger impl ----------------------------------')
    // contract deploy
    const L1CrossDomainMessengerFactory = await hre.ethers.getContractFactory(L1CrossDomainMessengerContractFactoryName)
    const l1CrossDomainMessenger = await L1CrossDomainMessengerFactory.deploy(
        MorphPortalProxyAddress
    )
    await l1CrossDomainMessenger.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerImplStorageName, l1CrossDomainMessenger.address.toLocaleLowerCase(), l1CrossDomainMessenger.deployTransaction.hash);
    // check params then storge
    await assertContractVariable(
        l1CrossDomainMessenger,
        'PORTAL',
        MorphPortalProxyAddress
    )
    blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1CrossDomainMessengerImplStorageName, l1CrossDomainMessenger.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // proxy upgradeAndCall
    console.log('Upgrading the L1CrossDomainMessenger proxy...')
    await L1CrossDomainMessengerProxy.upgradeToAndCall(
        l1CrossDomainMessenger.address,
        l1CrossDomainMessenger.interface.encodeFunctionData('initialize', [])
    )
    // Wait for the transaction to execute properly.
    await awaitCondition(
        async () => {
            const temp = new ethers.Contract(
                L1CrossDomainMessengerProxy.address,
                L1CrossDomainMessengerProxy.interface,
                L1CrossDomainMessengerProxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: ethers.constants.AddressZero,
            })
            return (
                actual.toLocaleLowerCase() === l1CrossDomainMessenger.address.toLocaleLowerCase()
            )
        },
        30000,
        1000
    )
    // check params
    checkContract = new ethers.Contract(
        L1CrossDomainMessengerProxy.address,
        l1CrossDomainMessenger.interface,
        L1CrossDomainMessengerProxy.provider
    )
    await assertContractVariable(
        checkContract,
        'PORTAL',
        MorphPortalProxyAddress
    )
    console.log('Upgrading the L1CrossDomainMessenger proxy success...')
    return ''
}

export default deployPortalAndMessengerImpl
