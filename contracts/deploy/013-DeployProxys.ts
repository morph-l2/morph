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

export const deployContractProxyByStorageName = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string,
): Promise<string> => {
    const ProxyFactoryName = ContractFactoryName.DefaultProxy
    // const ProxyAdminAddress = getContractAddressByName(path, ImplStorageName.ProxyAdmin)

    const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    const proxy = await ProxyFactory.deploy(await deployer.getAddress())
    console.log("%s=%s ; TX_HASH: %s", storageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    await assertContractVariable(
        proxy,
        'admin',
        await deployer.getAddress()
    )
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    const err = await storge(path, storageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    // return 
    return ''
}

export const deployContractProxys = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
): Promise<string> => {
    const L1CrossDomainMessengerStroageName = ProxyStorageName.L1CrossDomainMessengerProxyStroageName
    const L1MessageQueueProxyStroageName = ProxyStorageName.L1MessageQueueProxyStroageName
    const L2GasPriceOracleProxyStorageName = ProxyStorageName.L2GasPriceOracleProxyStorageName

    const RollupProxyStroageName = ProxyStorageName.RollupProxyStorageName
    const StakingProxyStroageName = ProxyStorageName.StakingProxyStroageName
    const L1SequencerProxyStroageName = ProxyStorageName.L1SequencerProxyStroageName

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessengerProxy deploy 
    let err = await deployContractProxyByStorageName(hre, path, deployer, L1CrossDomainMessengerStroageName)
    if (err != '') {
        return err
    }

    // L1MessageQueueProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1MessageQueueProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ staking contracts deploy ************************
    // StakingProxy ddeploy
    err = await deployContractProxyByStorageName(hre, path, deployer, StakingProxyStroageName)
    if (err != '') {
        return err
    }

    // L1SequencerProxy ddeploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1SequencerProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // L2GasPriceOracle deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L2GasPriceOracleProxyStorageName)
    if (err != '') {
        return err
    }

    // RollupProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, RollupProxyStroageName)
    if (err != '') {
        return err
    }

    // return nil
    return ''
}

export default deployContractProxys
