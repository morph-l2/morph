import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName } from "../src/deploy-utils";
import { predeploys } from '../src/constants'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const deployContractImpls = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    // factory name
    const L1CrossDomainMessengerFactoryName = ContractFactoryName.L1CrossDomainMessenger
    const StakingFactoryName = ContractFactoryName.Staking
    const L1SequencerFactoryName = ContractFactoryName.L1Sequencer
    const L1MessageQueueWithGasPriceOracleFactoryName = ContractFactoryName.L1MessageQueueWithGasPriceOracle
    const RollupFactoryName = ContractFactoryName.Rollup

    const L1GatewayRouterFactoryName = ContractFactoryName.L1GatewayRouter
    const L1ETHGatewayFactoryName = ContractFactoryName.L1ETHGateway
    const L1StandardERC20GatewayFactoryName = ContractFactoryName.L1StandardERC20Gateway
    const L1ERC721GatewayFactoryName = ContractFactoryName.L1ERC721Gateway
    const L1ERC1155GatewayFactoryName = ContractFactoryName.L1ERC1155Gateway

    // implement storage name
    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const StakingImplStorageName = ImplStorageName.StakingStorageName
    const L1SequencerImplStorageName = ImplStorageName.L1SequencerStorageName
    const L1MessageQueueWithGasPriceOracleImplStroageName = ImplStorageName.L1MessageQueueWithGasPriceOracle
    const RollupImplStorageName = ImplStorageName.RollupStorageName
    const L1GatewayRouterImplStorageName = ImplStorageName.L1GatewayRouterStorageName
    const L1ETHGatewayImplStorageName = ImplStorageName.L1ETHGatewayStorageName
    const L1StandardERC20GatewayImplStorageName = ImplStorageName.L1StandardERC20GatewayStorageName
    const L1ERC721GatewayImplStorageName = ImplStorageName.L1ERC721GatewayStorageName
    const L1ERC1155GatewayImplStorageName = ImplStorageName.L1ERC1155GatewayStorageName

    // proxy contract address
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const EnforcedTxGatewayAddress = getContractAddressByName(path, ProxyStorageName.EnforcedTxGatewayProxyStroageName)

    // ************************ messenger contracts deploy ************************
    // L1CrossDomainMessenger deploy
    let Factory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
    let contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    let blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    // L1MessageQueueWithGasPriceOracle deploy
    Factory = await hre.ethers.getContractFactory(L1MessageQueueWithGasPriceOracleFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress, RollupProxyAddress, EnforcedTxGatewayAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1MessageQueueWithGasPriceOracleImplStroageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1MessageQueueWithGasPriceOracleImplStroageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ rollup contracts deploy ************************
    // Rollup deploy
    const l2ChainID: string = config.l2ChainID
    Factory = await hre.ethers.getContractFactory(RollupFactoryName)
    contract = await Factory.deploy(l2ChainID, L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", RollupImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    // check params
    await assertContractVariable(contract, 'layer2ChainId', l2ChainID)
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, RollupImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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
    err = await storge(path, L1GatewayRouterImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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
    err = await storge(path, L1StandardERC20GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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
    err = await storge(path, L1ETHGatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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
    err = await storge(path, L1ERC721GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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
    err = await storge(path, L1ERC1155GatewayImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // ************************ staking contracts deploy ************************
    // Staking deploy 
    Factory = await hre.ethers.getContractFactory(StakingFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", StakingImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, StakingImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1Sequencer deploy 
    Factory = await hre.ethers.getContractFactory(L1SequencerFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    await assertContractVariable(
        contract,
        'OTHER_SEQUENCER',
        predeploys.L2Sequencer
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // return
    return ''
}

export default deployContractImpls
