import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storage, getContractAddressByName } from "../src/deploy-utils";
import { predeploys } from '../src/constants'
import { Mutex } from 'async-mutex';

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const deployContractImplsConcurrently = async (
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
    const L1ERC721GatewayImplStorageName = ImplStorageName.L1ERC721GatewayStorageName
    const L1ERC1155GatewayImplStorageName = ImplStorageName.L1ERC1155GatewayStorageName
    const WhitelistImplStorageName = ImplStorageName.Whitelist
    const EnforcedTxGatewayStorageName = ImplStorageName.EnforcedTxGatewayStorageName

    // proxy contract address
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const EnforcedTxGatewayAddress = getContractAddressByName(path, ProxyStorageName.EnforcedTxGatewayProxyStorageName)

    console.log("start to deploy contract implementations concurrently...")
    let nonce = await hre.ethers.provider.getTransactionCount(deployer.getAddress())
    const mutex = new Mutex();
    const deployContract = async (factoryName: string, storageName: string, args: any[] = []) => {
        const release = await mutex.acquire();
        const nonceToUse = nonce
        nonce++;  // Increment nonce for each deployment
        release();  // Release the lock
        
        console.log(`Starting deployment for: ${storageName}, args: ${args}, nonce: `, nonceToUse);
        const Factory = await hre.ethers.getContractFactory(factoryName)
        const contract = await Factory.deploy(...args, {
                nonce: nonceToUse,
            })
        

        await contract.deployed()
        console.log("%s=%s ; TX_HASH: %s", storageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
        
        if (factoryName == L1StakingFactoryName) {
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
        }
        const blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        console.log(`Deployment completed for: ${storageName}`);
        const err = await storage(path, storageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
        return err
    }

    try {
        const deployPromises = []

        deployPromises.push(deployContract(WhitelistFactoryName, WhitelistImplStorageName, [config.contractAdmin]))
        deployPromises.push(deployContract(L1CrossDomainMessengerFactoryName, L1CrossDomainMessengerImplStorageName))
        deployPromises.push(deployContract(L1MessageQueueWithGasPriceOracleFactoryName, L1MessageQueueWithGasPriceOracleImplStorageName, [L1CrossDomainMessengerProxyAddress, RollupProxyAddress, EnforcedTxGatewayAddress]))
        deployPromises.push(deployContract(RollupFactoryName, RollupImplStorageName, [config.l2ChainID]))
        deployPromises.push(deployContract(L1GatewayRouterFactoryName, L1GatewayRouterImplStorageName))
        deployPromises.push(deployContract(L1StandardERC20GatewayFactoryName, L1StandardERC20GatewayImplStorageName))
        deployPromises.push(deployContract(L1CustomERC20GatewayFactoryName, L1CustomERC20GatewayImplStorageName))
        deployPromises.push(deployContract(L1WithdrawLockERC20GatewayFactoryName, L1WithdrawLockERC20GatewayImplStorageName))
        deployPromises.push(deployContract(L1ReverseCustomGatewayFactoryName, L1ReverseCustomGatewayImplStorageName))
        deployPromises.push(deployContract(L1ETHGatewayFactoryName, L1ETHGatewayImplStorageName))

        const WETHAddress = getContractAddressByName(path, ImplStorageName.WETH)
        deployPromises.push(deployContract(L1WETHGatewayFactoryName, L1WETHGatewayImplStorageName, [WETHAddress, predeploys.L2WETH]))

        deployPromises.push(deployContract(EnforcedTxGatewayFactoryName, EnforcedTxGatewayStorageName))
        deployPromises.push(deployContract(L1ERC721GatewayFactoryName, L1ERC721GatewayImplStorageName))
        deployPromises.push(deployContract(L1ERC1155GatewayFactoryName, L1ERC1155GatewayImplStorageName))

        deployPromises.push(deployContract(L1StakingFactoryName, StakingImplStorageName, [L1CrossDomainMessengerProxyAddress]))

        // L1Sequencer deploy (no constructor args)
        const L1SequencerFactoryName = ContractFactoryName.L1Sequencer
        const L1SequencerImplStorageName = ImplStorageName.L1SequencerStorageName
        deployPromises.push(deployContract(L1SequencerFactoryName, L1SequencerImplStorageName))

        const results = await Promise.all(deployPromises)

        for (const result of results) {
            if (result != '') {
                return result
            }
        }

        console.log("Complete contract implementations deployment...")
        return ''
    } catch (err) {
        console.error("Error during deployment:", err)
        return "Deployment failed"
    }
}

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

    // ************************ sequencer contracts deploy ************************
    // L1Sequencer deploy
    const L1SequencerFactoryName = ContractFactoryName.L1Sequencer
    const L1SequencerImplStorageName = ImplStorageName.L1SequencerStorageName
    Factory = await hre.ethers.getContractFactory(L1SequencerFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storage(path, L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // return
    return ''
}

export default deployContractImpls
