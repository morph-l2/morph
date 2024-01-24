import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";

import {
    deployProxyAdmin,
    deployLibAddressManager,
    deployZkEvmVerifierV1,
    deployContractProxys,
    deployContractImpls,
    MessengerInit,
    RollupInit,
    StakingInit,
} from '../deploy/index'
import { ethers } from "ethers";
import StakingRegister from "../deploy/018-StakingRegister";

task("deploy")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const stroagePath = taskArgs.storagepath

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- deploy  ProxyAdmin ----------------------------------')
        // ProxyAdmin
        let err = await deployProxyAdmin(hre, stroagePath, deployer)
        if (err != '') {
            console.log('Deploy deployProxyAdmin failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- deploy  LibAddressManager ----------------------------------')
        // Lib_AddressManager
        err = await deployLibAddressManager(hre, stroagePath, deployer)
        if (err != '') {
            console.log('Deploy address manager failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- deploy  Proxys ----------------------------------')
        err = await deployContractProxys(hre, stroagePath, deployer)
        if (err != '') {
            console.log('Deploy Proxys failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- deploy  ZkEvmVerifierV1 ----------------------------------')
        err = await deployZkEvmVerifierV1(hre, stroagePath, deployer)
        if (err != '') {
            console.log('Deploy deploy ZkEvmVerifierV1 failed, err: ', err)
            return
        }
    });

task("initialize")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const stroagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- deploy  Impls ----------------------------------')
        let err = await deployContractImpls(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Deploy deploy Impls failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- Messenger init ----------------------------------')
        err = await MessengerInit(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Messenger init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Rollup init ----------------------------------')
        err = await RollupInit(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Rollup init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Staking init ----------------------------------')
        err = await StakingInit(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Staking init failed, err: ', err)
            return
        }
    });

task("fund")
    .setAction(async (taskArgs, hre) => {
        const config = hre.deployConfig
        const signer = await hre.ethers.getSigners()
        for (let i = 0; i < config.l2SequencerPks.length; i++) {
            let sequencer = new ethers.Wallet(config.l2SequencerPks[i], hre.ethers.provider)
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
        const stroagePath = taskArgs.storagepath
        const config = hre.deployConfig
        console.log('################################## console parameters ##################################')

        for (let i = 0; i < config.l2SequencerPks.length; i++) {
            let sequencer = new ethers.Wallet(config.l2SequencerPks[i], hre.ethers.provider)
            console.log(`sequencer-${i}:` + await sequencer.getAddress() + ', Balance: ' + await sequencer.getBalance())

            console.log(`\n---------------------------------- register  sequencer-${i} ----------------------------------`)
            let err = await StakingRegister(hre, stroagePath, sequencer, config.l2SequencerTmKeys[i], config.l2SequencerBlsKeys[i])
            if (err != '') {
                console.log(`Deploy Staking Sequencer-${i} failed, err: `, err)
                return
            }
        }
    });
