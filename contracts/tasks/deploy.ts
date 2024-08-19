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
    deployContractImpls,
    MessengerInit,
    RollupInit,
    GatewayInit,
    StakingInit,
    AdminTransfer,
    ContractInit,
    StakingRegister,
} from '../deploy/index'
import { ethers } from "ethers";

task("deploy")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
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
        err = await deployContractProxies(hre, storagePath, deployer,config)
        if (err != '') {
            console.log('Deploy Proxys failed, err: ', err)
            return
        }

        console.log('\n---------------------------------- deploy  ZkEvmVerifierV1 ----------------------------------')
        err = await deployZkEvmVerifierV1(hre, storagePath, deployer)
        if (err != '') {
            console.log('Deploy deploy ZkEvmVerifierV1 failed, err: ', err)
            return
        }
    });

task("initialize")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- deploy  Impls ----------------------------------')
        let err = await deployContractImpls(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Deploy deploy Impls failed, err: ', err)
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
        console.log('\n---------------------------------- Staking init ----------------------------------')
        err = await StakingInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Staking init failed, err: ', err)
            return
        }
        console.log('\n---------------------------------- Admin Transfer ----------------------------------')
        err = await AdminTransfer(hre, storagePath, deployer, config)
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
