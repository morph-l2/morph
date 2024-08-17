import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storage, getContractAddressByName } from "../src/deploy-utils";
import { predeploys } from '../src/constants'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const deployContractImpls = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    // factory name
    const L1CrossDomainMessengerFactoryName = ContractFactoryName.L1CrossDomainMessenger
    const L1StakingFactoryName = ContractFactoryName.L1Staking
    const L1MessageQueueWithGasPriceOracleFactoryName = ContractFactoryName.L1MessageQueueWithGasPriceOracle
    const RollupFactoryName = ContractFactoryName.Rollup
    const WhitelistFactoryName = ContractFactoryName.Whitelist

    const L1GatewayRouterFactoryName = ContractFactoryName.L1GatewayRouter
    const L1ETHGatewayFactoryName = ContractFactoryName.L1ETHGateway
    const L1StandardERC20GatewayFactoryName = ContractFactoryName.L1StandardERC20Gateway
    const L1CustomERC20GatewayFactoryName = ContractFactoryName.L1CustomERC20Gateway
    const L1WithdrawLockERC20GatewayFactoryName = ContractFactoryName.L1WithdrawLockERC20Gateway
    const L1ReverseCustomGatewayFactoryName = ContractFactoryName.L1ReverseCustomGateway
    const L1ERC721GatewayFactoryName = ContractFactoryName.L1ERC721Gateway
    const L1ERC1155GatewayFactoryName = ContractFactoryName.L1ERC1155Gateway
    const L1WETHGatewayFactoryName = ContractFactoryName.L1WETHGateway
    const L1USDCGatewayFactoryName = ContractFactoryName.L1USDCGateway
    const EnforcedTxGatewayFactoryName = ContractFactoryName.EnforcedTxGateway

    // implement storage name
    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const StakingImplStorageName = ImplStorageName.L1StakingStorageName
    const L1MessageQueueWithGasPriceOracleImplStorageName = ImplStorageName.L1MessageQueueWithGasPriceOracle
    const RollupImplStorageName = ImplStorageName.RollupStorageName
    const L1GatewayRouterImplStorageName = ImplStorageName.L1GatewayRouterStorageName
    const L1ETHGatewayImplStorageName = ImplStorageName.L1ETHGatewayStorageName
    const L1StandardERC20GatewayImplStorageName = ImplStorageName.L1StandardERC20GatewayStorageName
    const L1CustomERC20GatewayImplStorageName = ImplStorageName.L1CustomERC20GatewayStorageName
    const L1WithdrawLockERC20GatewayImplStorageName = ImplStorageName.L1WithdrawLockERC20GatewayStorageName
    const L1ReverseCustomGatewayImplStorageName = ImplStorageName.L1ReverseCustomGatewayStorageName
    const L1WETHGatewayImplStorageName = ImplStorageName.L1WETHGatewayStorageName
    const L1USDCGatewayImplStorageName = ImplStorageName.L1USDCGatewayStorageName
    const L1ERC721GatewayImplStorageName = ImplStorageName.L1ERC721GatewayStorageName
    const L1ERC1155GatewayImplStorageName = ImplStorageName.L1ERC1155GatewayStorageName
    const WhitelistImplStorageName = ImplStorageName.Whitelist
    const EnforcedTxGatewayStorageName = ImplStorageName.EnforcedTxGatewayStorageName

    // proxy contract address
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const EnforcedTxGatewayAddress = getContractAddressByName(path, ProxyStorageName.EnforcedTxGatewayProxyStorageName)

    // ************************ system contracts deploy ************************
    // whitelist deploy
    let Factory = await hre.ethers.getContractFactory(WhitelistFactoryName)
    let contract = await Factory.deploy(config.contractAdmin)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", WhitelistImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    let blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storage(path, WhitelistImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessenger deploy
    Factory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    // L1MessageQueueWithGasPriceOracle deploy
    Factory = await hre.ethers.getContractFactory(L1MessageQueueWithGasPriceOracleFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress, RollupProxyAddress, EnforcedTxGatewayAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1MessageQueueWithGasPriceOracleImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1MessageQueueWithGasPriceOracleImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // Rollup deploy
    const l2ChainID: string = config.l2ChainID
    Factory = await hre.ethers.getContractFactory(RollupFactoryName)
    contract = await Factory.deploy(l2ChainID)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", RollupImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    // check params
    await assertContractVariable(contract, 'LAYER_2_CHAIN_ID', l2ChainID)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, RollupImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ gateway contracts deploy ************************
    // L1GatewayRouter deploy
    Factory = await hre.ethers.getContractFactory(L1GatewayRouterFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1GatewayRouterImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1GatewayRouterImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1StandardERC20Gateway deploy
    Factory = await hre.ethers.getContractFactory(L1StandardERC20GatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1StandardERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1StandardERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1CustomERC20Gateway deploy
    Factory = await hre.ethers.getContractFactory(L1CustomERC20GatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CustomERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1CustomERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1WithdrawLockERC20Gateway deploy
    Factory = await hre.ethers.getContractFactory(L1WithdrawLockERC20GatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1WithdrawLockERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1WithdrawLockERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ReverseCustomGateway deploy
    Factory = await hre.ethers.getContractFactory(L1ReverseCustomGatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ReverseCustomGatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1ReverseCustomGatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ETHGateway deploy
    Factory = await hre.ethers.getContractFactory(L1ETHGatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ETHGatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1ETHGatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1WETHGateway deploy
    const WETHAddress = getContractAddressByName(path, ImplStorageName.WETH)
    Factory = await hre.ethers.getContractFactory(L1WETHGatewayFactoryName)
    contract = await Factory.deploy(WETHAddress, predeploys.L2WETH)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1WETHGatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1WETHGatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1USDCGateway deploy
    const L1USDCAddress = getContractAddressByName(path, ImplStorageName.USDC)
    Factory = await hre.ethers.getContractFactory(L1USDCGatewayFactoryName)
    contract = await Factory.deploy(L1USDCAddress, predeploys.L2USDC)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1USDCGatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1USDCGatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // EnforcedTxGateway deploy
    Factory = await hre.ethers.getContractFactory(EnforcedTxGatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", EnforcedTxGatewayStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, EnforcedTxGatewayStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ERC721Gateway deploy
    Factory = await hre.ethers.getContractFactory(L1ERC721GatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ERC721GatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1ERC721GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ERC1155Gateway deploy
    Factory = await hre.ethers.getContractFactory(L1ERC1155GatewayFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ERC1155GatewayImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1ERC1155GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ staking contracts deploy ************************
    // Staking deploy 
    Factory = await hre.ethers.getContractFactory(L1StakingFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", StakingImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    await assertContractVariable(
        contract,
        'OTHER_STAKING',
        predeploys.L2Staking.toLowerCase()
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, StakingImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // return
    return ''
}

export default deployContractImpls
