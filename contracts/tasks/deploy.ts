import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "dotenv/config";

import { task } from "hardhat/config";

import {
    deployProxyAdmin,
    deployEmptyContract,
    deployZkEvmVerifierV1,
    deployContractProxies,
    deployContractProxiesConcurrently,
    deployContractImpls,
    deployContractImplsConcurrently,
    SubmitterInit,
    MessengerInit,
    RollupInit,
    GatewayInit,
    AdminTransfer,
    AdminTransferConcurrently,
    ContractInit,
    SubmitterRegister,
    SequencerInit,
} from '../deploy/index'
import { ethers } from "ethers";

task("deploy")
    .addParam('storagepath')
    .addOptionalParam('concurrent', 'Use concurrent deployment', 'false')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const concurrent = taskArgs.concurrent
        const deployer = await hre.ethers.provider.getSigner();
        const config = hre.deployConfig

        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- deploy  ProxyAdmin ----------------------------------')
        // ProxyAdmin
        let err = await deployProxyAdmin(hre, storagePath, deployer)
        if (err != '') {
            console.log('Deploy deployProxyAdmin failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- deploy  EmptyContract ----------------------------------')
        // EmptyContract
        err = await deployEmptyContract(hre, storagePath, deployer)
        if (err != '') {
            console.log('Deploy address manager failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- deploy  Proxys ----------------------------------')
        if (concurrent === 'true') {
            err = await deployContractProxiesConcurrently(hre, storagePath, deployer,config)
        } else {
            err = await deployContractProxies(hre, storagePath, deployer,config)
        }
        if (err != '') {
            console.log('Deploy Proxys failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- deploy  ZkEvmVerifierV1 ----------------------------------')
        err = await deployZkEvmVerifierV1(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Deploy deploy ZkEvmVerifierV1 failed, err: ', err)
            return
        }
    });

task("initialize")
    .addParam('storagepath')
    .addOptionalParam('concurrent', 'Use concurrent deployment', 'false')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const concurrent = taskArgs.concurrent
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- deploy  Impls ----------------------------------')
        let err: any
        if (concurrent === 'true') {
            err = await deployContractImplsConcurrently(hre, storagePath, deployer, config)
        } else {
            err = await deployContractImpls(hre, storagePath, deployer, config)
        }
        if (err != '') {
            console.log('Deploy deploy Impls failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- Submitter init ----------------------------------')
        err = await SubmitterInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Submitter init failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- Messenger init ----------------------------------')
        err = await MessengerInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Messenger init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Rollup init ----------------------------------')
        err = await RollupInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Rollup init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Gateway init ----------------------------------')
        err = await GatewayInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Rollup init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Sequencer init ----------------------------------')
        err = await SequencerInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Sequencer init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Admin Transfer ----------------------------------')
        if (concurrent === 'true') {
            err = await AdminTransferConcurrently(hre, storagePath, deployer, config)
        } else {
            err = await AdminTransfer(hre, storagePath, deployer, config)
        }
        
        if (err != '') {
            console.log('OwnerTransfer failed, err: ', err)
            return
        }



        console.log('\n---------------------------------- Contract Init ----------------------------------')
        err = await ContractInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('ContractInit failed, err: ', err)
            return
        }
        // todo transfer contract owner 
    });

task("fund")
    .setAction(async (taskArgs, hre) => {
        console.log('\n---------------------------------- Fund Submitters ----------------------------------')
        const signer = await hre.ethers.getSigners()
        const batchSubmitterPkList: string[] = JSON.parse(process.env.batchSubmitterPks || "[]");
        for (let i = 0; i < batchSubmitterPkList.length; i++) {
            const submitter = new ethers.Wallet(batchSubmitterPkList[i], hre.ethers.provider)
            const tx = {
                to: submitter.address,
                value: ethers.utils.parseEther("100")
            }
            let balance = (await submitter.getBalance()).toString()

            if (balance.length < 20) {
                let receipt = await signer[0].sendTransaction(tx)
                await receipt.wait()
            }
            balance = (await submitter.getBalance()).toString()
            console.log(`${submitter.address} has balance: ${balance}`)
        }
    })


task("register")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig
        const owner = await hre.ethers.provider.getSigner();
        const batchSubmitterPkList: string[] = JSON.parse(process.env.batchSubmitterPks || "[]");
        const configuredAddresses: string[] = config.batchSubmitterAddresses;
        if (batchSubmitterPkList.length !== configuredAddresses.length || batchSubmitterPkList.length === 0) {
            throw new Error("batchSubmitterPks must contain one key for every configured work or backup submitter")
        }
        for (let i = 0; i < batchSubmitterPkList.length; i++) {
            const submitter = new ethers.Wallet(batchSubmitterPkList[i], hre.ethers.provider)
            if (configuredAddresses[i].toLowerCase() !== submitter.address.toLowerCase()) {
                throw new Error(`batch submitter key ${i} does not match configured address`)
            }
            console.log(`\n---------------------------------- register submitter-${i} ----------------------------------`)
            console.log(`submitter-${i}:` + submitter.address + ', Balance: ' + await submitter.getBalance())
            const err = await SubmitterRegister(hre, storagePath, owner, submitter)
            if (err != '') {
                console.log(`Register Submitter-${i} failed, err: `, err)
                return
            }
        }
    });
