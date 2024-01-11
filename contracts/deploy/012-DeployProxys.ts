import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"


export const deployContractProxys = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
): Promise<string> => {
    const ChugSplashProxyFactoryName = ContractFactoryName.L1ChugSplashProxy
    const ProxyFactoryName = ContractFactoryName.DefaultProxy
    const ResolvedDelegateProxyFactoryName = ContractFactoryName.ResolvedDelegateProxy

    const L1StandardBridgeProxyStroageName = ProxyStorageName.L1StandardBridgeProxyStroageName
    const L1CrossDomainMessengerStroageName = ProxyStorageName.L1CrossDomainMessengerProxyStroageName
    const MorphPortalProxyStroageName = ProxyStorageName.MorphPortalProxyStroageName
    const MorphMintableERC20FactoryProxyStroageName = ProxyStorageName.MorphMintableERC20FactoryProxyStroageName
    const L1ERC721BridgeProxyStroageName = ProxyStorageName.L1ERC721BridgeProxyStroageName
    const SystemConfigProxyStroageName = ProxyStorageName.SystemConfigProxyStorageName
    const SystemDictatorProxyStroageName = ProxyStorageName.SystemDictatorProxyStorageName
    const RollupProxyStroageName = ProxyStorageName.RollupProxyStorageName
    const StakingProxyStroageName = ProxyStorageName.StakingProxyStroageName
    const L1SequencerProxyStroageName = ProxyStorageName.L1SequencerProxyStroageName
    // const MultipleVersionRollupVerifierProxyStorageName = ProxyStorageName.MultipleVersionRollupVerifierStorageName

    const proxyAdmin = getContractAddressByName(path, ImplStorageName.ProxyAdmin)
    const addressManager = getContractAddressByName(path, ImplStorageName.AddressManager)

    // Proxy__L1StandardBridge deploy
    let ProxyFactory = await hre.ethers.getContractFactory(ChugSplashProxyFactoryName)
    let proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1StandardBridgeProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'getOwner', await deployer.getAddress())
    let blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, L1StandardBridgeProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1CrossDomainMessengerProxy deploy 
    ProxyFactory = await hre.ethers.getContractFactory(ResolvedDelegateProxyFactoryName)
    proxy = await ProxyFactory.deploy(addressManager, 'L1CrossDomainMessenger')
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1CrossDomainMessengerStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // MorphPortalProxy deploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(proxyAdmin)
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", MorphPortalProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', proxyAdmin)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, MorphPortalProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }


    // StakingProxy ddeploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", StakingProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, StakingProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1SequencerProxy ddeploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1SequencerProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1SequencerProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // MorphMintableERC20FactoryProxy deploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(proxyAdmin)
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", MorphMintableERC20FactoryProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', proxyAdmin)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, MorphMintableERC20FactoryProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ERC721BridgeProxy ddeploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ERC721BridgeProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1ERC721BridgeProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // SystemConfigProxy deploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(proxyAdmin)
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", SystemConfigProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', proxyAdmin)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, SystemConfigProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // SystemDictatorProxy deploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", SystemDictatorProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, SystemDictatorProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // RollupProxy deploy
    ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    proxy = await ProxyFactory.deploy(await deployer.getAddress())
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", RollupProxyStroageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash)
    // check paramss
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, RollupProxyStroageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    return ''
}

export default deployContractProxys
