import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, assertContractVariableWithSigner, awaitCondition, storge } from "../src/deploy-utils";
import { ethers } from 'ethers'
import { predeploys } from '../src/constants'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"


export const adminTransferByProxyStorageName = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string,
): Promise<string> => {
    const EmptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract)
    const ProxyAdminImplAddr = getContractAddressByName(path, ImplStorageName.ProxyAdmin)
    const ProxyAddr = getContractAddressByName(path, storageName)
    const deployerAddr = (await deployer.getAddress()).toLocaleLowerCase()

    const IProxyContract = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, ProxyAddr, deployer)
    {
        const implAddr = (await IProxyContract.implementation()).toLocaleLowerCase()
        const admin = (await IProxyContract.admin()).toLocaleLowerCase()
        if (implAddr === EmptyContractImplAddr.toLocaleLowerCase()) {
            return `Proxy implementation address ${implAddr} should not be empty contract address ${EmptyContractImplAddr}`
        }
        if (admin !== deployerAddr) {
            return `Proxy admin address ${admin} should deployer address ${deployerAddr}`

        }
    }
    console.log("change admin")
    await IProxyContract.changeAdmin(ProxyAdminImplAddr)
    await assertContractVariable(
        IProxyContract,
        'admin',
        ProxyAdminImplAddr,
        ProxyAdminImplAddr // caller
    )
    console.log(`${storageName} admin transfer from ${deployerAddr} to ProxyAdmin ${ProxyAdminImplAddr} `)
    return ''
}
export const AdminTransfer = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
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
    // const EnforcedTxGatewayProxyStroageName = ProxyStorageName.EnforcedTxGatewayProxyStroageName
    
    // ************************ messenger contracts admin change ************************
    // L1CrossDomainMessengerProxy admin change 
    let err = await adminTransferByProxyStorageName(hre, path, deployer, L1CrossDomainMessengerStroageName)
    if (err != '') {
        return err
    }

    // L1MessageQueueWithGasPriceOracleProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ staking contracts admin change ************************
    // StakingProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, StakingProxyStroageName)
    if (err != '') {
        return err
    }

    // L1SequencerProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1SequencerProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts admin change ************************
    // RollupProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, RollupProxyStroageName)
    if (err != '') {
        return err
    }

    // ************************ gateway contracts admin change ************************
    // L1GatewayRouterProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1GatewayRouterProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ETHGatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ERC721GatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ERC721GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // L1ERC1155GatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ERC1155GatewayProxyStroageName)
    if (err != '') {
        return err
    }

    // EnforcedTxGatewayProxy admin change
    // err = await adminTransferByProxyStorageName(hre, path, deployer, EnforcedTxGatewayProxyStroageName)
    // if (err != '') {
    //     return err
    // }
    // return nil
    return ''
}

export default AdminTransfer
