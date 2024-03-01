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
        const ProxyFactoryName = 'ITransparentUpgradeableProxy'
        for (let i = 0; i < ContractAddresss.length; i++) {
            if (ContractAddresss[i] === predeploys.MorphStandardERC20) {
                continue
            }
            const proxy = await hre.ethers.getContractAt(ProxyFactoryName, ContractAddresss[i])
            const temp = new ethers.Contract(
                proxy.address,
                proxy.interface,
                proxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: predeploys.ProxyAdmin,
            })
            console.log(`implementation is: ${actual}`)
            const adminAddr = await temp.callStatic['admin']({
                from: predeploys.ProxyAdmin,
            })

            console.log(`implementation is equal ProxyAdmin: ${adminAddr.toLowerCase() == predeploys.ProxyAdmin.toLowerCase()}`)
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
        console.log(`MorphStandardERC20Factory params check \n owner ${owner == predeploys.L2StandardERC20Gateway} \n implementation ${implementation.toLowerCase() == predeploys.MorphStandardERC20.toLowerCase()}`)

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
        const router = routerFactory.attach('0xa513e6e4b8f2a923d98304ec87f64353c4d5c853')
        const res = await router["depositETH(uint256,uint256)"](hre.ethers.utils.parseEther('1'), 110000, { value: hre.ethers.utils.parseEther('1.1') })
        const receipt = await res.wait()
        console.log(`Deposit\n from ${receipt.from}\n blockNum ${receipt.blockNumber}\n tx ${receipt.transactionHash}\n status ${receipt.status == 1}`)
    });

task("withdraw-l2-eth")
    .setAction(async (taskArgs, hre) => {
        const routerFactory = await hre.ethers.getContractFactory('L2GatewayRouter')
        const router = routerFactory.attach(predeploys.L2GatewayRouter)
        const res = await router["withdrawETH(uint256,uint256)"](1000, 110000, { value: hre.ethers.utils.parseEther('1.1') })
        const receipt = await res.wait()
        console.log(`Deposit\n from ${receipt.from}\n blockNum ${receipt.blockNumber}\n tx ${receipt.transactionHash}\n status ${receipt.status == 1}`)
    });

task("deposit-l1-gateway-eth")
    .setAction(async (taskArgs, hre) => {
        const Factory = await hre.ethers.getContractFactory('L1ETHGateway')
        const contract = Factory.attach("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6")
        const res = await contract["depositETH(uint256,uint256)"](hre.ethers.utils.parseEther('1'), 110000, { value: hre.ethers.utils.parseEther('1.1') })
        const recipet = await res.wait()
        console.log(`Deposit status ${recipet.status == 1}`)
    });

task("deploy-l1-token")
    .setAction(async (taskArgs, hre) => {
        console.log("Deploy L1 ERC20 tokken")
        const Factory = await hre.ethers.getContractFactory('MockERC20')
        const token = await Factory.deploy("L1 Token", "l1token", 18)
        const rec = await token.deployed()
        console.log(`Token deployed at L1 ${token.address}, deploy txHash: ${rec.deployTransaction.hash}`)
    });

task("deposit-erc20-token")
    .addParam('l1token')
    .addParam('balance')
    .setAction(async (taskArgs, hre) => {
        const signers = await hre.ethers.getSigners()
        console.log(`signer ${signers[0].address}`)
        const l1RouterAddr = '0xa513e6e4b8f2a923d98304ec87f64353c4d5c853'

        const routerFactory = await hre.ethers.getContractFactory('L1GatewayRouter')
        const router = routerFactory.attach(l1RouterAddr)
        const tokenFactory = await hre.ethers.getContractFactory('MockERC20')
        const l1token = tokenFactory.attach(taskArgs.l1token)

        // mint and approve
        const l2TokenAddr = await router.getL2ERC20Address(l1token.address)
        console.log(`tokenPair : l1Token ${l1token.address}, l2Token ${l2TokenAddr}`)

        let res = await l1token.mint(signers[0].address, taskArgs.balance)
        let rec = await res.wait()
        const balance = await l1token.balanceOf(signers[0].address)
        console.log(`mint ${rec.status == 1}: signer ${signers[0].address} has balance ${balance}`)

        // approve
        res = await l1token.approve(router.address, taskArgs.balance)
        rec = await res.wait()
        const allowance = await l1token.allowance(signers[0].address, router.address)
        console.log(`approve ${rec.status == 1}: router ${router.address} has allowance ${allowance}`)

        // first deposit require gasLimit > 410000
        res = await router["depositERC20(address,uint256,uint256)"](l1token.address, allowance, 420000, { value: hre.ethers.utils.parseEther('1') })
        rec = await res.wait()
        console.log(`Deposit\n from ${rec.from}\n blockNum ${rec.blockNumber}\n tx ${rec.transactionHash}\n status ${rec.status == 1}`)
    });

task("erc20Balances")
    .addParam('token')
    .addParam('address')
    .setAction(async (taskArgs, hre) => {
        const tokenFactory = await hre.ethers.getContractFactory('MockERC20')
        const token = tokenFactory.attach(taskArgs.token)
        const balance = await token.balanceOf(taskArgs.address)
        console.log(`${taskArgs.address} has ${balance}`)
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

task("getCurrentSubmitter")
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

        const res = await contract.getCurrentSubmitter()
        console.log(`res : ${res}`)
    });


task("l2factory-deploy")
    .setAction(async (taskArgs, hre) => {
        const factoryF = await hre.ethers.getContractFactory('MorphStandardERC20Factory')
        const fc = await factoryF.deploy(predeploys.MorphStandardERC20)
        await fc.deployed()
        console.log(`MorphStandardERC20Factory deployed at ${fc.address}`)
    });

task("l2token-deploy-check")
    .addParam("factory")
    .addParam("l1token")
    .setAction(async (taskArgs, hre) => {
        const factoryF = await hre.ethers.getContractFactory('MorphStandardERC20Factory')
        const fc = factoryF.attach(taskArgs.factory)
        let l2Addr = await fc.computeL2TokenAddress(predeploys.L2StandardERC20Gateway, taskArgs.l1token)
        let code = await hre.ethers.provider.getCode(l2Addr)
        console.log(`L2TokenAddress : ${l2Addr}, code ${code}`)

        let res = await fc.deployL2Token(predeploys.L2StandardERC20Gateway, taskArgs.l1token)
        let rec = await res.wait()
        console.log(`rec : ${rec.status == 1}`)
        let codeAfter = await hre.ethers.provider.getCode(l2Addr)
        console.log(`L2TokenAddress : ${l2Addr}, code ${codeAfter}`)
    });