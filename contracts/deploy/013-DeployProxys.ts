import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { updateStorage, getContractAddressByName, assertContractVariableWithSigner } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const deployContractProxyByStorageName = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string,
): Promise<string> => {
    const emptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract)
    const ProxyFactoryName = ContractFactoryName.DefaultProxy

    const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
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
    const err = await updateStorage(path, storageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    // return 
    return ''
}

export const deployContractProxies = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    component: string,
): Promise<string> => {
    const L1CrossDomainMessengerStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStorageName
    const L1MessageQueueWithGasPriceOracleProxyStorageName = ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName

    const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
    const L1StakingProxyStorageName = ProxyStorageName.L1StakingProxyStorageName

    const L1GatewayRouterProxyStorageName = ProxyStorageName.L1GatewayRouterProxyStorageName
    const L1ETHGatewayProxyStorageName = ProxyStorageName.L1ETHGatewayProxyStorageName
    const L1StandardERC20GatewayProxyStorageName = ProxyStorageName.L1StandardERC20GatewayProxyStorageName
    const L1CustomERC20GatewayProxyStorageName = ProxyStorageName.L1CustomERC20GatewayProxyStorageName
    const L1ERC721GatewayProxyStorageName = ProxyStorageName.L1ERC721GatewayProxyStorageName
    const L1ERC1155GatewayProxyStorageName = ProxyStorageName.L1ERC1155GatewayProxyStorageName
    const EnforcedTxGatewayProxyStorageName = ProxyStorageName.EnforcedTxGatewayProxyStorageName
    const L1WETHGatewayProxyStorageName = ProxyStorageName.L1WETHGatewayProxyStorageName

    const WETHFactoryName = ContractFactoryName.WETH
    const WETHImplStorageName = ImplStorageName.WETH

    let err = '';

    // ************************ token contracts deploy ************************
    // L1WETH deploy
    if (component.includes('c')) {
        let Factory = await hre.ethers.getContractFactory(WETHFactoryName)
        let contract = await Factory.deploy()
        await contract.deployed()
        console.log("%s=%s ; TX_HASH: %s", WETHImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        let err = await updateStorage(path, WETHImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
        if (err != '') {
            return err
        }
    }

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessengerProxy deploy 
    if (component.includes('d')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1CrossDomainMessengerStorageName)
        if (err != '') {
            return err
        }
    }

    // L1MessageQueueWithGasPriceOracleProxy deploy
    if (component.includes('e')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ staking contracts deploy ************************
    // StakingProxy deploy
    if (component.includes('f')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1StakingProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ rollup contracts deploy ************************
    // RollupProxy deploy
    if (component.includes('g')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, RollupProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // ************************ gateway contracts deploy ************************
    // L1GatewayRouterProxy deploy
    if (component.includes('h')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1GatewayRouterProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ETHGatewayProxy deploy
    if (component.includes('i')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1ETHGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1StandardERC20GatewayProxy deploy
    if (component.includes('j')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1CustomERC20GatewayProxy deploy
    if (component.includes('k')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1CustomERC20GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ERC721GatewayProxy deploy
    if (component.includes('l')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC721GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1ERC1155GatewayProxy deploy
    if (component.includes('m')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC1155GatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // EnforcedTxGatewayProxy deploy
    if (component.includes('n')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, EnforcedTxGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // L1WETHGatewayProxy deploy
    if (component.includes('o')) {
        err = await deployContractProxyByStorageName(hre, path, deployer, L1WETHGatewayProxyStorageName)
        if (err != '') {
            return err
        }
    }

    // return nil
    return ''
}

export default deployContractProxies