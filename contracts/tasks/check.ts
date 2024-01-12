import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";

task("check-l2")
    .setAction(async (taskArgs, hre) => {
        const ContractAddresss = [
            '0x5300000000000000000000000000000000000001',
            '0x530000000000000000000000000000000000000F',
            '0x4200000000000000000000000000000000000002',
            '0x4200000000000000000000000000000000000010',
            '0x4200000000000000000000000000000000000007',
            '0x4200000000000000000000000000000000000012',
            '0x4200000000000000000000000000000000000013',
            '0x4200000000000000000000000000000000000015',
            '0x4200000000000000000000000000000000000000',
            '0x4200000000000000000000000000000000000014',
            '0x4200000000000000000000000000000000000017',
            '0x4200000000000000000000000000000000000018',
        ]
        const ProxyFactoryName = 'Proxy'
        const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
        for (let i = 0; i < ContractAddresss.length; i++) {
            const proxy = ProxyFactory.attach(ContractAddresss[i])
            const temp = new ethers.Contract(
                proxy.address,
                proxy.interface,
                proxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: ethers.constants.AddressZero,
            })
            console.log(`implementation is: ${actual}`)
        }
    });

task("deposit-l1")
    .setAction(async (taskArgs, hre) => {
        const L1StandBridgeFactory = await hre.ethers.getContractFactory('L1StandardBridge')
        const l1StandBridge = L1StandBridgeFactory.attach('0x6900000000000000000000000000000000000003')
        const res = await l1StandBridge.depositETH(20000, '0x', { value: hre.ethers.utils.parseEther('1') })
        const recipet = await res.wait()
        console.log(`Deposit status ${recipet.status == 1}`)
    });

task("getBalances")
    .addParam('address')
    .setAction(async (taskArgs, hre) => {
        console.log(`${taskArgs.address} has ${await hre.waffle.provider.getBalance(taskArgs.address)}`)
    });

task("withdraw-l2")
    .setAction(async (taskArgs, hre) => {
        const L2StandBridgeFactory = await hre.ethers.getContractFactory('L2StandardBridge')
        const l2StandBridge = L2StandBridgeFactory.attach('0x4200000000000000000000000000000000000010')
        const res = await l2StandBridge.withdraw(
            '0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000',
            ethers.utils.parseEther('1'),
            0,
            '0x',
            {
                value: ethers.utils.parseEther('1'),
            }
        )
        console.log(await res.wait())
    });

task("getSequencerAddresses")
    .setAction(async (taskArgs, hre) => {
        const factory = await hre.ethers.getContractFactory('L2Sequencer')
        const contract = factory.attach('0x5300000000000000000000000000000000000003')
        const res = await contract.getSequencerAddresses(false)

        console.log(`getSequencerAddresses : ${res}`)
    });

task("rollupEpoch")
    .setAction(async (taskArgs, hre) => {
        const factory = await hre.ethers.getContractFactory('Gov')
        const contract = factory.attach('0x5300000000000000000000000000000000000004')
        const res = await contract.rollupEpoch()

        console.log(`rollupEpoch : ${res}`)
    });

task("getNextSubmitter")
    .setAction(async (taskArgs, hre) => {
        const factory = await hre.ethers.getContractFactory('Submitter')
        const contract = factory.attach('0x5300000000000000000000000000000000000005')

        const govFactory = await hre.ethers.getContractFactory('Gov')
        const gov = govFactory.attach('0x5300000000000000000000000000000000000004')

        const L2SequencerFactory = await hre.ethers.getContractFactory('L2Sequencer')
        const sequencer = L2SequencerFactory.attach('0x5300000000000000000000000000000000000003')

        let sequencers = await sequencer.getSequencerAddresses(false)
        const rollupEpoch = (await gov.rollupEpoch()).toNumber()
        const sequencersLen = sequencers.length
        console.log(`sequencersLen: ${sequencersLen} , rollupEpoch: ${rollupEpoch}`)

        let nextEpochStart = (await contract.nextEpochStart()).toNumber()
        console.log(`nextEpochStart : ${nextEpochStart}`)
        let nextSubmitterIndex = await contract.nextSubmitterIndex()
        console.log(`nextSubmitterIndex : ${nextSubmitterIndex}`)
        const block = await hre.ethers.provider.getBlock('latest')
        console.log(block.timestamp)

        const res = await contract.getNextSubmitter()
        console.log(`res : ${res}`)
    });
