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

export const deployPortalProxyAndMessengerProxy = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log('\n---------------------------------- deploy MorphPortal proxy ----------------------------------')
    const MorphPortalProxyStorageName = ProxyStorageName.MorphPortalProxyStroageName
    const MorphPortalImplStorageName = ImplStorageName.MorphPortalStroageName
    // deploy proxy
    const MorphPortalProxyFactoy = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const MorphPortalProxy = await MorphPortalProxyFactoy.deploy(await deployer.getAddress())
    console.log("%s=%s ; TX_HASH: %s", MorphPortalProxyStorageName, MorphPortalProxy.address.toLocaleLowerCase(), MorphPortalProxy.deployTransaction.hash);
    await assertContractVariable(MorphPortalProxy, 'admin', await deployer.getAddress())
    let blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, MorphPortalProxyStorageName, MorphPortalProxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    console.log('\n---------------------------------- deploy L1CrossDomainMessenger Proxy ----------------------------------')
    const L1CrossDomainMessengerProxyStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStroageName
    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const L1CrossDomainMessengerContractFactoryName = ContractFactoryName.L1CrossDomainMessenger
    // deploy proxy
    const L1CrossDomainMessengerProxyFactoy = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const L1CrossDomainMessengerProxy = await L1CrossDomainMessengerProxyFactoy.deploy(await deployer.getAddress())
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerProxyStorageName, L1CrossDomainMessengerProxy.address.toLocaleLowerCase(), L1CrossDomainMessengerProxy.deployTransaction.hash);
    await assertContractVariable(L1CrossDomainMessengerProxy, 'admin', await deployer.getAddress())
    blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1CrossDomainMessengerProxyStorageName, L1CrossDomainMessengerProxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    return ''
}

export default deployPortalProxyAndMessengerProxy
