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


// 1、归集的多签钱包给 deployer 转入手续费
// 这部分手续费包含二层需要消化的部分，由deployer deposit到l2上在分发
// 归集地址需要还需要给下述地址转入操作手续费：7个sequencers

// 2、deploy  ProxyAdmin
// 3、deploy  EmptyContract
// 4、deploy  Proxys
// 5、deploy  ZkEvmVerifierV1

task("deployStep1")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const deployer = await hre.ethers.provider.getSigner();
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
        err = await deployContractProxies(hre, storagePath, deployer)
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

// 6、deploy  Impls

task("deployStep2")
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
    });

// 7、Messenger init

task("deployStep3")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Messenger init ----------------------------------')
        let err = await MessengerInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Messenger init failed, err: ', err)
            return
        }
    });

// 8、Rollup init 
// 部署 Impl__MultipleVersionRollupVerifier 的步骤可以省去

task("deployStep4")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Rollup init ----------------------------------')
        let err = await RollupInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Rollup init failed, err: ', err)
            return
        }
    });

// 9、Gateway init

task("deployStep5")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Gateway init ----------------------------------')
        let err = await GatewayInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Rollup init failed, err: ', err)
            return
        }
    });

// 10、Staking init

task("deployStep6")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Staking init ----------------------------------')
        let err = await StakingInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('Staking init failed, err: ', err)
            return
        }
    });

// 11、Admin Transfer

task("deployStep7")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Admin Transfer ----------------------------------')
        let err = await AdminTransfer(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('OwnerTransfer failed, err: ', err)
            return
        }
    });

// 12、Contract Init

task("deployStep8")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('\n---------------------------------- Contract Init ----------------------------------')
        let err = await ContractInit(hre, storagePath, deployer, config)
        if (err != '') {
            console.log('ContractInit failed, err: ', err)
            return
        }
    });

// 13、Do Staking Sequencer…
task("deployStep9")
    .addParam('storagepath')
    .setAction(async (taskArgs, hre) => {
        // Initialization parameters
        const storagePath = taskArgs.storagepath
        const config = hre.deployConfig
        // console.log("====",process.env.l2SequencerPks)

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





// yarn hardhat deployStep1 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep2 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep3 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep4 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep5 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep6 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep7 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep8 --storagepath distributedDeploy.json --network l1
// yarn hardhat deployStep9 --storagepath distributedDeploy.json --network l1