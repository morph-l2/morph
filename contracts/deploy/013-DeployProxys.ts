import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName, assertContractVariableWithSigner } from "../src/deploy-utils";
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
    const emptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract)
    const ProxyFactoryName = ContractFactoryName.DefaultProxy

    const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    // TransparentUpgradeableProxy deploy with empthContract as impl, deployer as admin
    const proxy = await ProxyFactory.deploy(emptyContractImplAddr, await deployer.getAddress(), "0x")
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", storageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    // check params
    const IProxyContract = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, proxy.address)

    console.log(await IProxyContract.admin())
    await assertContractVariableWithSigner(
        IProxyContract,
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
    const L1MessageQueueWithGasPriceOracleProxyStroageName = ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStroageName

    const RollupProxyStroageName = ProxyStorageName.RollupProxyStorageName
    const StakingProxyStroageName = ProxyStorageName.StakingProxyStroageName
    const L1SequencerProxyStroageName = ProxyStorageName.L1SequencerProxyStroageName

    const L1GatewayRouterProxyStroageName = ProxyStorageName.L1GatewayRouterProxyStroageName
    const L1ETHGatewayProxyStroageName = ProxyStorageName.L1ETHGatewayProxyStroageName
    const L1StandardERC20GatewayProxyStroageName = ProxyStorageName.L1StandardERC20GatewayProxyStroageName
    const L1ERC721GatewayProxyStroageName = ProxyStorageName.L1ERC721GatewayProxyStroageName
    const L1ERC1155GatewayProxyStroageName = ProxyStorageName.L1ERC1155GatewayProxyStroageName
    const EnforcedTxGatewayProxyStroageName = ProxyStorageName.EnforcedTxGatewayProxyStroageName

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessengerProxy deploy 
    let err = await deployContractProxyByStorageName(hre, path, deployer, L1CrossDomainMessengerStroageName)
    if (err != '') {
        return err
    }

    // L1MessageQueueWithGasPriceOracleProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ staking contracts deploy ************************
    // StakingProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, StakingProxyStroageName)
    if (err != '') {
        return err
    }

    // L1SequencerProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1SequencerProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // RollupProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, RollupProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ gateway contracts deploy ************************
    // L1GatewayRouterProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1GatewayRouterProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ETHGatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ERC721GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC721GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ERC1155GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC1155GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // EnforcedTxGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, EnforcedTxGatewayProxyStroageName)
    if (err != '') {
        return err
    }
    // return nil
    return ''
}

export default deployContractProxys