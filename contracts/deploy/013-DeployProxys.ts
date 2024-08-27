import "@nomiclabs/hardhat-web3"
import "@nomiclabs/hardhat-ethers"
import "@nomiclabs/hardhat-waffle"

import { HardhatRuntimeEnvironment } from "hardhat/types"
import { storage, getContractAddressByName, assertContractVariableWithSigner } from "../src/deploy-utils"
import { ImplStorageName, ProxyStorageName, ContractFactoryName } from "../src/types"

export const deployContractProxyByStorageName = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    storageName: string
): Promise<string> => {
    const emptyContractImplAddr = getContractAddressByName(path, ImplStorageName.EmptyContract)
    const ProxyFactoryName = ContractFactoryName.DefaultProxy

    const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
    // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
    const proxy = await ProxyFactory.deploy(emptyContractImplAddr, await deployer.getAddress(), "0x")
    await proxy.deployed()
    console.log("%s=%s ; TX_HASH: %s", storageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash)
    // check params
    const IProxyContract = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, proxy.address)

    console.log(await IProxyContract.admin())
    await assertContractVariableWithSigner(IProxyContract, "admin", await deployer.getAddress())
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    const err = await storage(path, storageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != "") {
        return err
    }
    // return
    return ""
}

export const deployContractProxies = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    const L1CrossDomainMessengerStorageName = ProxyStorageName.L1CrossDomainMessengerProxyStorageName
    const L1MessageQueueWithGasPriceOracleProxyStorageName =
        ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName

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

    const WETHFactoryName = ContractFactoryName.WETH
    const WETHImplStorageName = ImplStorageName.WETH

    const USDCFactoryName = ContractFactoryName.USDC
    const USDCImplStorageName = ImplStorageName.USDC
    let err = ""

    // ************************ token contracts deploy ************************
    if (config.l1WETHAddress == "") {
        // L1WETH deploy
        let Factory = await hre.ethers.getContractFactory(WETHFactoryName)
        let contract = await Factory.deploy()
        await contract.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            WETHImplStorageName,
            contract.address.toLocaleLowerCase(),
            contract.deployTransaction.hash
        )
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        err = await storage(path, WETHImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
    } else {
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        err = await storage(path, WETHImplStorageName, config.l1WETHAddress.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
    }

    if (config.l1USDCAddress == "") {
        // L1WETH deploy
        let Factory = await hre.ethers.getContractFactory(USDCFactoryName)
        let contract = await Factory.deploy()
        await contract.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            USDCImplStorageName,
            contract.address.toLocaleLowerCase(),
            contract.deployTransaction.hash
        )
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        err = await storage(path, USDCImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
        await contract.initialize(
            "USDC",
            "USDC",
            "USD",
            18,
            config.contractAdmin,
            config.contractAdmin,
            config.contractAdmin,
            config.contractAdmin
        )
    } else {
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        err = await storage(path, USDCImplStorageName, config.l1USDCAddress.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
    }

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessengerProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1CrossDomainMessengerStorageName)
    if (err != "") {
        return err
    }

    // L1MessageQueueWithGasPriceOracleProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1MessageQueueWithGasPriceOracleProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ staking contracts deploy ************************
    // StakingProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1StakingProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // RollupProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, RollupProxyStorageName)
    if (err != "") {
        return err
    }

    // ************************ gateway contracts deploy ************************
    // L1GatewayRouterProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1GatewayRouterProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ETHGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1StandardERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1StandardERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1CustomERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1CustomERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1WithdrawLockERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1WithdrawLockERC20GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ReverseCustomERC20GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ReverseCustomGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ERC721GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC721GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1ERC1155GatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1ERC1155GatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // EnforcedTxGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, EnforcedTxGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1WETHGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1WETHGatewayProxyStorageName)
    if (err != "") {
        return err
    }

    // L1USDCGatewayProxy deploy
    err = await deployContractProxyByStorageName(hre, path, deployer, L1USDCGatewayProxyStorageName)
    if (err != "") {
        return err
    }
    // return nil
    return ""
}

export default deployContractProxies
