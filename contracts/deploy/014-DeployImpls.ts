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
    const L1MessageQueueFactoryName = ContractFactoryName.L1MessageQueue
    const RollupFactoryName = ContractFactoryName.Rollup
    const L2GasPriceOracleFactoryName = ContractFactoryName.L2GasPriceOracle

    // implement storage name
    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const StakingImplStorageName = ImplStorageName.StakingStorageName
    const L1SequencerImplStorageName = ImplStorageName.L1SequencerStorageName
    const L1MessageQueueImplStroageName = ImplStorageName.L1MessageQueueStroageName
    const RollupImplStorageName = ImplStorageName.RollupStorageName
    const L2GasPriceOracleImplStorageName = ImplStorageName.L2GasPriceOracleStorageName

    // proxy contract address
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)

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
    // L1MessageQueue deploy
    Factory = await hre.ethers.getContractFactory(L1MessageQueueFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1MessageQueueImplStroageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1MessageQueueImplStroageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
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

    // ************************ rollup contracts deploy ************************
    // L2GasPriceOracle deploy
    Factory = await hre.ethers.getContractFactory(L2GasPriceOracleFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L2GasPriceOracleImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L2GasPriceOracleImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
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


    // return
    return ''
}

export default deployContractImpls
