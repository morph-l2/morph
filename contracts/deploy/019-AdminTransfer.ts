import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName } from "../src/deploy-utils";

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
    console.log(`change ${storageName} admin transfer from ${deployerAddr} to ProxyAdmin ${ProxyAdminImplAddr} `)
    const res = await IProxyContract.changeAdmin(ProxyAdminImplAddr)
    await res.wait()
    await assertContractVariable(
        IProxyContract,
        'admin',
        ProxyAdminImplAddr,
        ProxyAdminImplAddr // caller
    )
    console.log(`admin transfer successful `)
    return ''
}
export const AdminTransfer = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    const L1CrossDomainMessengerStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStorageName
    const L1MessageQueueWithGasPriceOracleProxyStorageName = ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName

    const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
    const StakingProxyStorageName = ProxyStorageName.StakingProxyStorageName
    const L1SequencerProxyStorageName = ProxyStorageName.L1SequencerProxyStorageName

    const L1GatewayRouterProxyStorageName = ProxyStorageName.L1GatewayRouterProxyStorageName
    const L1ETHGatewayProxyStorageName = ProxyStorageName.L1ETHGatewayProxyStorageName
    const L1StandardERC20GatewayProxyStorageName = ProxyStorageName.L1StandardERC20GatewayProxyStorageName
    const L1ERC721GatewayProxyStorageName = ProxyStorageName.L1ERC721GatewayProxyStorageName
    const L1ERC1155GatewayProxyStorageName = ProxyStorageName.L1ERC1155GatewayProxyStorageName
    // const EnforcedTxGatewayProxyStorageName = ProxyStorageName.EnforcedTxGatewayProxyStorageName
    const L1WETHGatewayProxyStorageName = ProxyStorageName.L1WETHGatewayProxyStorageName

    // ************************ messenger contracts admin change ************************
    // L1CrossDomainMessengerProxy admin change 
    let err = await adminTransferByProxyStorageName(hre, path, deployer, L1CrossDomainMessengerStorageName)
    if (err != '') {
        return err
    }

    // L1MessageQueueWithGasPriceOracleProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStorageName)
    if (err != '') {
        return err
    }

    // ************************ staking contracts admin change ************************
    // StakingProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, StakingProxyStorageName)
    if (err != '') {
        return err
    }

    // L1SequencerProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1SequencerProxyStorageName)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts admin change ************************
    // RollupProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, RollupProxyStorageName)
    if (err != '') {
        return err
    }

    // ************************ gateway contracts admin change ************************
    // L1GatewayRouterProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1GatewayRouterProxyStorageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ETHGatewayProxyStorageName)
    if (err != '') {
        return err
    }

    // L1ETHGatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStorageName)
    if (err != '') {
        return err
    }

    // L1ERC721GatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ERC721GatewayProxyStorageName)
    if (err != '') {
        return err
    }

    // L1ERC1155GatewayProxy admin change
    err = await adminTransferByProxyStorageName(hre, path, deployer, L1ERC1155GatewayProxyStorageName)
    if (err != '') {
        return err
    }

    // EnforcedTxGatewayProxy admin change
    // err = await adminTransferByProxyStorageName(hre, path, deployer, EnforcedTxGatewayProxyStorageName)
    // if (err != '') {
    //     return err
    // }

    err = await adminTransferByProxyStorageName(hre, path, deployer, L1WETHGatewayProxyStorageName)
    if (err != '') {
        return err
    }

    return ''
}

export default AdminTransfer
