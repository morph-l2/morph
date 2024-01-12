import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { ethers } from 'ethers'
import { predeploys } from '../src/constants'
import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName ,awaitCondition} from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName
} from "./types"

export const deployL1ERC721Bridge = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log('\n---------------------------------- deploy  L1ERC721Bridge ----------------------------------')
    const proxyStorageName = ProxyStorageName.L1ERC721BridgeProxyStroageName
    const implStorageName = ImplStorageName.L1ERC721BridgeStroageName
    const contractFactoryName = ContractFactoryName.L1ERC721Bridge
    // deploy proxy
    const ProxyFactoy = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const proxy = await ProxyFactoy.deploy(await deployer.getAddress())
    console.log("%s=%s ; TX_HASH: %s", proxyStorageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    let blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, proxyStorageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    // get init params
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)

    // deploy
    const Factory = await hre.ethers.getContractFactory(contractFactoryName)
    const contract = await Factory.deploy(
        L1CrossDomainMessengerProxyAddress,
        predeploys.L2ERC721Bridge
    )
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", implStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params then storge
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, implStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // proxy upgradeAndCall
    console.log('Upgrading the L1StandardBridge proxy...')
    await proxy.upgradeToAndCall(
        contract.address,
        contract.interface.encodeFunctionData('initialize', [])
    )
    // Wait for the transaction to execute properly.
    await awaitCondition(
        async () => {
            const temp = new ethers.Contract(
                proxy.address,
                proxy.interface,
                proxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: ethers.constants.AddressZero,
            })
            return (
                actual.toLocaleLowerCase() === contract.address.toLocaleLowerCase()
            )
        },
        30000,
        1000
    )

    // check params
    const checkContract = new ethers.Contract(
        proxy.address,
        contract.interface,
        proxy.provider
    )
    await assertContractVariable(
        checkContract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    console.log('Upgrading the L1StandardBridge proxy success...')

    return ''
}

export default deployL1ERC721Bridge
