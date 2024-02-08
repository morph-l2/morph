import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";
import { predeploys } from "../src/constants";

task("check-l2")
    .setAction(async (taskArgs, hre) => {
        let ContractAddresss = []
        let keys = Object.keys(predeploys);

        keys.forEach((key) => {
            ContractAddresss.push(predeploys[key])
        });
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
            const adminAddr = await temp.callStatic['admin']({
                from: ethers.constants.AddressZero,
            })
            console.log(`implementation is equal ProxyAdmin: ${adminAddr == predeploys.ProxyAdmin}`)
        }
    });

task("check-l2-status")
    .setAction(async (taskArgs, hre) => {
        const mpFactory = await hre.ethers.getContractFactory('L2ToL1MessagePasser')
        const mpContract = mpFactory.attach(predeploys.L2ToL1MessagePasser)
        const messageRoot = await mpContract.messageRoot()
        console.log(`L2ToL1MessagePasser params check \n root set status ${messageRoot == '0x27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757'}`)

        const gwrFactory = await hre.ethers.getContractFactory('L2GatewayRouter')
        const gwrContract = gwrFactory.attach(predeploys.L2GatewayRouter)
        let owner = await gwrContract.owner()
        let ethGateway = await gwrContract.ethGateway()
        let defaultERC20Gateway = await gwrContract.defaultERC20Gateway()
        console.log(`L2GatewayRouter params check \n owner ${owner} \n ETHGateway set ${ethGateway == predeploys.L2ETHGateway} \n  ERC20Gateway set ${defaultERC20Gateway == predeploys.L2StandardERC20Gateway}`)

        const sqFactory = await hre.ethers.getContractFactory('L2Sequencer')
        const sqContract = sqFactory.attach(predeploys.L2Sequencer)
        const currVersion = await sqContract.currentVersion()
        console.log(`L2Sequencer params check \n currentVersion status ${currVersion == 1}`)

        const govFactory = await hre.ethers.getContractFactory('Gov')
        const govContract = govFactory.attach(predeploys.L2Gov)
        const proposalInterval = await govContract.proposalInterval()
        const batchBlockInterval = await govContract.batchBlockInterval()
        const batchMaxBytes = await govContract.batchMaxBytes()
        const batchTimeout = await govContract.batchTimeout()
        const rollupEpoch = await govContract.rollupEpoch()
        const maxChunks = await govContract.maxChunks()
        console.log(`Gov params check \n proposalInterval ${proposalInterval} \n batchMaxBytes ${batchMaxBytes} \n batchBlockInterval ${batchBlockInterval} \n batchTimeout ${batchTimeout} \n rollupEpoch ${rollupEpoch} \n maxChunks ${maxChunks}`)

        const submitterFactory = await hre.ethers.getContractFactory('Submitter')
        const submitterContract = submitterFactory.attach(predeploys.L2Submitter)
        const nextEpochStart = await submitterContract.nextEpochStart()
        console.log(`Submitter params check \n nextEpochStart ${nextEpochStart}`)

        const ethgwFactory = await hre.ethers.getContractFactory('L2ETHGateway')
        const ethgwContract = ethgwFactory.attach(predeploys.L2ETHGateway)
        let router = await ethgwContract.router()
        let messenger = await ethgwContract.messenger()
        let counterpart = await ethgwContract.counterpart()
        console.log(`L2ETHGateway params check \n router ${router == predeploys.L2GatewayRouter} \n messenger ${messenger == predeploys.L2CrossDomainMessenger} \n counterpart ${counterpart}`)

        const cdmFactory = await hre.ethers.getContractFactory('L2CrossDomainMessenger')
        const cdmContract = cdmFactory.attach(predeploys.L2CrossDomainMessenger)
        let paused = await cdmContract.paused()
        let xDomainMessageSender = await cdmContract.xDomainMessageSender()
        counterpart = await cdmContract.counterpart()
        let feeVault = await cdmContract.feeVault()
        console.log(`L2CrossDomainMessenger params check \n paused ${paused == false} \n xDomainMessageSender ${xDomainMessageSender} \n counterpart ${counterpart} \n feeVault ${feeVault}`)

        const segwFactory = await hre.ethers.getContractFactory('L2StandardERC20Gateway')
        const segwContract = segwFactory.attach(predeploys.L2StandardERC20Gateway)
        const tokenFactory = await segwContract.tokenFactory()
        router = await segwContract.router()
        messenger = await segwContract.messenger()
        counterpart = await segwContract.counterpart()
        console.log(`L2StandardERC20Gateway params check \n tokenFactory ${tokenFactory == predeploys.MorphStandardERC20Factory} \n router ${router == predeploys.L2GatewayRouter} \n messenger ${messenger == predeploys.L2CrossDomainMessenger} \n counterpart ${counterpart}`)

        const gw721Factory = await hre.ethers.getContractFactory('L2ERC721Gateway')
        const gw721Contract = gw721Factory.attach(predeploys.L2ERC721Gateway)
        router = await gw721Contract.router()
        messenger = await gw721Contract.messenger()
        counterpart = await gw721Contract.counterpart()
        console.log(`L2ERC721Gateway params check \n tokenFactory ${router == ethers.constants.AddressZero} \n messenger ${messenger == predeploys.L2CrossDomainMessenger} \n counterpart ${counterpart}`)

        const txFeeFactory = await hre.ethers.getContractFactory('L2TxFeeVault')
        const txFeeContract = txFeeFactory.attach(predeploys.L2TxFeeVault)
        owner = await txFeeContract.owner()
        const minWithdrawAmount = await txFeeContract.minWithdrawAmount()
        const recipient = await txFeeContract.recipient()
        console.log(`L2TxFeeVault params check \n owner ${owner} \n minWithdrawAmount ${minWithdrawAmount} \n recipient ${recipient}`)

        const paFactory = await hre.ethers.getContractFactory('ProxyAdmin')
        const paContract = paFactory.attach(predeploys.ProxyAdmin)
        owner = await paContract.owner()
        console.log(`ProxyAdmin params check \n owner ${owner}`)

        const gw1155Factory = await hre.ethers.getContractFactory('L2ERC1155Gateway')
        const gw1155Contract = gw1155Factory.attach(predeploys.L2ERC1155Gateway)
        router = await gw1155Contract.router()
        messenger = await gw1155Contract.messenger()
        counterpart = await gw1155Contract.counterpart()
        console.log(`L2ERC1155Gateway params check \n tokenFactory ${router == ethers.constants.AddressZero} \n messenger ${messenger == predeploys.L2CrossDomainMessenger} \n counterpart ${counterpart}`)

        const ms20Factory = await hre.ethers.getContractFactory('MorphStandardERC20')
        const ms20Contract = ms20Factory.attach(predeploys.MorphStandardERC20)
        const name = await ms20Contract.name()
        const symbol = await ms20Contract.symbol()
        const decimals = await ms20Contract.decimals()
        const gateway = await ms20Contract.gateway()
        counterpart = await ms20Contract.counterpart()
        console.log(`MorphStandardERC20 params check \n name ${name} \n symbol ${symbol} \n decimals ${decimals} \n gateway ${gateway} \n counterpart ${counterpart}`)

        const ms20fFactory = await hre.ethers.getContractFactory('MorphStandardERC20Factory')
        const ms20fContract = ms20fFactory.attach(predeploys.MorphStandardERC20Factory)
        owner = await ms20fContract.owner()
        const implementation = await ms20fContract.implementation()
        console.log(`MorphStandardERC20 params check \n owner ${owner == predeploys.L2StandardERC20Gateway} \n implementation ${implementation == predeploys.MorphStandardERC20}`)

        const gpoFactory = await hre.ethers.getContractFactory('GasPriceOracle')
        const gpoContract = gpoFactory.attach(predeploys.GasPriceOracle)
        owner = await gpoContract.owner()
        const overhead = await gpoContract.overhead()
        const scalar = await gpoContract.scalar()
        const l1BaseFee = await gpoContract.l1BaseFee()
        const allowListEnabled = await gpoContract.allowListEnabled()
        console.log(`GasPriceOracle params check \n owner ${owner} \n overhead ${overhead} \n scalar ${scalar} \n l1BaseFee ${l1BaseFee} \n allowListEnabled ${allowListEnabled}`)
    });

task("deposit-l1-eth")
    .setAction(async (taskArgs, hre) => {
        const routerFactory = await hre.ethers.getContractFactory('L1GatewayRouter')
        const router = routerFactory.attach('0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6')
        const res = await router["depositETH(uint256,uint256)"](hre.ethers.utils.parseEther('1'), 10000000, { value: hre.ethers.utils.parseEther('1.1') })
        const receipt = await res.wait()
        console.log(`Deposit\n from ${receipt.from}\n blockNum ${receipt.blockNumber}\n tx ${receipt.transactionHash}\n status ${receipt.status == 1}`)
    });

task("deposit-l1-gateway-eth")
    .setAction(async (taskArgs, hre) => {
        const Factory = await hre.ethers.getContractFactory('L1ETHGateway')
        const contract = Factory.attach("0x8a791620dd6260079bf849dc5567adc3f2fdc318")
        const res = await contract["depositETH(uint256,uint256)"](hre.ethers.utils.parseEther('1'), 100000, { value: hre.ethers.utils.parseEther('1.1') })
        const recipet = await res.wait()
        console.log(`Deposit status ${recipet.status == 1}`)
    });

task("getBalances")
    .addParam('address')
    .setAction(async (taskArgs, hre) => {
        console.log(`${taskArgs.address} has ${await hre.waffle.provider.getBalance(taskArgs.address)}`)
    });

task("withdraw-l2-eth")
    .setAction(async (taskArgs, hre) => {
        const Factory = await hre.ethers.getContractFactory('L2GatewayRouter')
        const contract = Factory.attach(predeploys.L2GatewayRouter)
        const res = await contract["withdrawETH(uint256,uint256)"](
            ethers.utils.parseEther('1'),
            0,
            {
                value: ethers.utils.parseEther('1'),
            }
        )
        const receipt = await res.wait()
        console.log(`Withdraw\n from ${receipt.from}\n blockNum ${receipt.blockNumber}\n tx ${receipt.transactionHash}\n status ${receipt.status == 1}`)
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
