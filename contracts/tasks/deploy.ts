import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "dotenv/config";
import "@types/node";

import { task } from "hardhat/config";
import { HardhatRuntimeEnvironment } from "hardhat/types";
import { Signer } from "@ethersproject/abstract-signer";
import { ethers } from "hardhat";

import {
    deployProxyAdmin,
    deployEmptyContract,
    deployZkEvmVerifierV1,
    deployContractProxies,
    deployContractProxiesConcurrently,
    deployContractImpls,
    deployContractImplsConcurrently,
    MessengerInit,
    RollupInit,
    GatewayInit,
    StakingInit,
    AdminTransfer,
    AdminTransferConcurrently,
    ContractInit,
    StakingRegister,
} from '../deploy/index'

interface DeployError extends Error {
    step?: string;
    details?: any;
}

async function handleDeployError(error: DeployError, step: string): Promise<string> {
    console.error(`Error during ${step}:`, {
        message: error.message,
        details: error.details,
        stack: error.stack
    });
    return `${step} failed: ${error.message}`;
}

task("deploy")
    .addParam('storagepath')
    .addOptionalParam('concurrent', 'Use concurrent deployment', 'false')
    .setAction(async (taskArgs, hre: HardhatRuntimeEnvironment) => {
        try {
            const storagePath = taskArgs.storagepath;
            const concurrent = taskArgs.concurrent === 'true';
            const deployer: Signer = await hre.ethers.provider.getSigner();
            const config = hre.deployConfig;

            console.log('################################## Deployment Parameters ##################################');
            console.log('Deployer:', await deployer.getAddress());

            // Deploy ProxyAdmin
            console.log('\n---------------------------------- Deploying ProxyAdmin ----------------------------------');
            let err = await deployProxyAdmin(hre, storagePath, deployer);
            if (err) throw { message: err, step: 'ProxyAdmin deployment' };

            // Deploy EmptyContract
            console.log('\n---------------------------------- Deploying EmptyContract ----------------------------------');
            err = await deployEmptyContract(hre, storagePath, deployer);
            if (err) throw { message: err, step: 'EmptyContract deployment' };

            // Deploy Proxies
            console.log('\n---------------------------------- Deploying Proxies ----------------------------------');
            err = concurrent 
                ? await deployContractProxiesConcurrently(hre, storagePath, deployer, config)
                : await deployContractProxies(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Proxies deployment' };

            // Deploy ZkEvmVerifierV1
            console.log('\n---------------------------------- Deploying ZkEvmVerifierV1 ----------------------------------');
            err = await deployZkEvmVerifierV1(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'ZkEvmVerifierV1 deployment' };

            console.log('\nDeployment completed successfully!');
        } catch (error) {
            const deployError = error as DeployError;
            console.error('Deployment failed:', {
                step: deployError.step || 'unknown step',
                error: deployError.message,
                details: deployError.details
            });
            process.exit(1);
        }
    });

task("initialize")
    .addParam('storagepath')
    .addOptionalParam('concurrent', 'Use concurrent deployment', 'false')
    .setAction(async (taskArgs, hre: HardhatRuntimeEnvironment) => {
        try {
            const storagePath = taskArgs.storagepath;
            const concurrent = taskArgs.concurrent === 'true';
            const config = hre.deployConfig;
            const deployer: Signer = await hre.ethers.provider.getSigner();

            console.log('################################## Initialization Parameters ##################################');
            console.log('Deployer:', await deployer.getAddress());

            // Deploy Implementations
            console.log('\n---------------------------------- Deploying Implementations ----------------------------------');
            let err = concurrent
                ? await deployContractImplsConcurrently(hre, storagePath, deployer, config)
                : await deployContractImpls(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Contract implementations deployment' };

            // Initialize Messenger
            console.log('\n---------------------------------- Initializing Messenger ----------------------------------');
            err = await MessengerInit(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Messenger initialization' };

            // Initialize Rollup
            console.log('\n---------------------------------- Initializing Rollup ----------------------------------');
            err = await RollupInit(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Rollup initialization' };

            // Initialize Gateway
            console.log('\n---------------------------------- Initializing Gateway ----------------------------------');
            err = await GatewayInit(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Gateway initialization' };

            // Initialize Staking
            console.log('\n---------------------------------- Initializing Staking ----------------------------------');
            err = await StakingInit(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Staking initialization' };

            // Transfer Admin
            console.log('\n---------------------------------- Transferring Admin ----------------------------------');
            err = concurrent
                ? await AdminTransferConcurrently(hre, storagePath, deployer, config)
                : await AdminTransfer(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Admin transfer' };

            // Initialize Contract
            console.log('\n---------------------------------- Initializing Contract ----------------------------------');
            err = await ContractInit(hre, storagePath, deployer, config);
            if (err) throw { message: err, step: 'Contract initialization' };

            console.log('\nInitialization completed successfully!');
        } catch (error) {
            const initError = error as DeployError;
            console.error('Initialization failed:', {
                step: initError.step || 'unknown step',
                error: initError.message,
                details: initError.details
            });
            process.exit(1);
        }
    });

task("fund")
    .setAction(async (taskArgs, hre) => {
        console.log('\n---------------------------------- Fund Staking ----------------------------------')
        const signer = await hre.ethers.getSigners()
        console.log(process.env.l2SequencerPks)
        let l2SequencerPkList = JSON.parse(process.env.l2SequencerPks);
        console.log(l2SequencerPkList)
        for (let i = 0; i < l2SequencerPkList.length; i++) {
            let sequencer = new ethers.Wallet(l2SequencerPkList[i], hre.ethers.provider)
            const tx = {
                to: sequencer.address,
                value: ethers.utils.parseEther("100")
            }
            let balance = (await sequencer.getBalance()).toString()

            if (balance.length < 20) {
                let receipt = await signer[0].sendTransaction(tx)
                await receipt.wait()
            }
            balance = (await sequencer.getBalance()).toString()
            console.log(`${sequencer.address} has balance: ${balance}`)
        }
    })


task("register")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig
        let l2SequencerPkList = JSON.parse(process.env.l2SequencerPks);
        for (let i = 0; i < l2SequencerPkList.length; i++) {
            let sequencer = new ethers.Wallet(l2SequencerPkList[i], hre.ethers.provider)
            console.log(`\n---------------------------------- register  sequencer-${i} ----------------------------------`)
            console.log(`sequencer-${i}:` + await sequencer.getAddress() + ', Balance: ' + await sequencer.getBalance())
            let err = await StakingRegister(hre, storagePath, sequencer, config.l2SequencerTmKeys[i], config.l2SequencerBlsKeys[i])
            if (err != '') {
                console.log(`Deploy Staking Sequencer-${i} failed, err: `, err)
                return
            }
        }
    });
