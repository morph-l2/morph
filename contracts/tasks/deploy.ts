import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "dotenv/config";

import { task } from "hardhat/config";

import {
    deployProxyAdmin,
    deployLibAddressManager,
    deployContractProxys,
    deployContractImpls,
    SystemDictatorInit,
    SystemDictatorSteps1,
    SystemDictatorSteps2,

    deployZkEvmVerifierV1,
    deployRollup,
    deploySystemConfig,
    deployPortalAndMessengerImpl,
    deployPortalProxyAndMessengerProxy,
    deployL1StandardBridge,
    deployL1ERC721Bridge
} from '../deploy/index'
import { ethers } from "ethers";
import StakingInit from "../deploy/017-Staking";

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
            console.log('Deploy deploy Proxys failed, err: ', err)
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

        console.log('\n---------------------------------- SystemDictator init ----------------------------------')
        err = await SystemDictatorInit(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Deploy SystemDictator init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- SystemDictator Steps1 ----------------------------------')
        err = await SystemDictatorSteps1(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Deploy SystemDictator steps1 failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- SystemDictator Steps2 ----------------------------------')
        err = await SystemDictatorSteps2(hre, stroagePath, deployer, config)
        if (err != '') {
            console.log('Deploy SystemDictator steps2 failed, err: ', err)
            return
        }
    });

const sequencerNum = 4

task("staking")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const stroagePath = taskArgs.storagepath
        const config = hre.deployConfig
        console.log('################################## console parameters ##################################')
        const l2SequencerPkList = JSON.parse(process.env.l2SequencerPks);
        for (let i = 0; i < sequencerNum; i++) {
            let sequencer = new ethers.Wallet(l2SequencerPkList[i], hre.ethers.provider)
            console.log(`sequencer-${i}:` + await sequencer.getAddress() + ', Balance: ' + await sequencer.getBalance())

            console.log(`\n---------------------------------- staking  sequencer-${i} ----------------------------------`)
            let err = await StakingInit(hre, stroagePath, sequencer, config.l2SequencerTmKeys[i], config.l2SequencerBlsKeys[i])
            if (err != '') {
                console.log(`Deploy Staking Sequencer-${i} failed, err: `, err)
                return
            }
        }
    });

task("deployonebyone")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const stroagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())
        await deployRollup(hre, stroagePath, deployer, config)
        await deploySystemConfig(hre, stroagePath, deployer, config)
        await deployPortalProxyAndMessengerProxy(hre, stroagePath, deployer, config)
        await deployPortalAndMessengerImpl(hre, stroagePath, deployer, config)
        await deployL1StandardBridge(hre, stroagePath, deployer, config)
        await deployL1ERC721Bridge(hre, stroagePath, deployer, config)
    });
