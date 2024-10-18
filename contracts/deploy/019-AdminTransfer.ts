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
} from "../src/types"

export const AdminTransferByProxyStorageName = async (
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
    configTmp: any,
    component: string
): Promise<string> => {
    const L1CrossDomainMessengerStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStorageName
    const L1MessageQueueWithGasPriceOracleProxyStorageName = ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName

    const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
    const L1StakingProxyStorageName = ProxyStorageName.L1StakingProxyStorageName

    const L1GatewayRouterProxyStorageName = ProxyStorageName.L1GatewayRouterProxyStorageName
    const L1ETHGatewayProxyStorageName = ProxyStorageName.L1ETHGatewayProxyStorageName
    const L1StandardERC20GatewayProxyStorageName = ProxyStorageName.L1StandardERC20GatewayProxyStorageName
    const L1CustomERC20GatewayProxyStorageName = ProxyStorageName.L1CustomERC20GatewayProxyStorageName
    const L1WithdrawLockERC20GatewayProxyStorageName = ProxyStorageName.L1WithdrawLockERC20GatewayProxyStorageName
    const L1ReverseCustomGatewayProxyStorageName = ProxyStorageName.L1ReverseCustomGatewayProxyStorageName
    const L1ERC721GatewayProxyStorageName = ProxyStorageName.L1ERC721GatewayProxyStorageName
    const L1ERC1155GatewayProxyStorageName = ProxyStorageName.L1ERC1155GatewayProxyStorageName
    const EnforcedTxGatewayProxyStorageName = ProxyStorageName.EnforcedTxGatewayProxyStorageName
    const L1WETHGatewayProxyStorageName = ProxyStorageName.L1WETHGatewayProxyStorageName
    const L1USDCGatewayProxyStorageName = ProxyStorageName.L1USDCGatewayProxyStorageName

    let err = ''

    // ************************ messenger contracts admin change ************************
    // L1CrossDomainMessengerProxy admin change
    if (component.includes('G')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1CrossDomainMessengerStorageName)
        if (err != '') {
            return err
        }
    }

    // L1MessageQueueWithGasPriceOracleProxy admin change
    if (component.includes('H')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ staking contracts admin change ************************
    // StakingProxy admin change
    if (component.includes('I')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1StakingProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ rollup contracts admin change ************************
    // RollupProxy admin change
    if (component.includes('J')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, RollupProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ gateway contracts admin change ************************
    // L1GatewayRouterProxy admin change
    if (component.includes('K')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1GatewayRouterProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ETHGatewayProxy admin change
    if (component.includes('L')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1ETHGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1StandardERC20GatewayProxy admin change
    if (component.includes('M')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1CustomERC20GatewayProxy admin change
    if (component.includes('N')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1CustomERC20GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1WithdrawLockERC20GatewayProxy admin change
    if (component.includes('O')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1WithdrawLockERC20GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ReverseCustomGatewayProxy admin change
    if (component.includes('P')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1ReverseCustomGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ERC721GatewayProxy admin change
    if (component.includes('Q')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1ERC721GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ERC1155GatewayProxy admin change
    if (component.includes('R')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1ERC1155GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // EnforcedTxGatewayProxy admin change
    if (component.includes('S')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, EnforcedTxGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1WETHGatewayProxy admin change
    if (component.includes('T')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1WETHGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1USDCGatewayProxy admin change
    if (component.includes('U')) {
        err = await AdminTransferByProxyStorageName(hre, path, deployer, L1USDCGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }
    
    return ''
}

module.exports = { AdminTransfer, AdminTransferByProxyStorageName }