import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";
import * as fs from "fs";
import * as path from "path";

const V2_1ABI = `[
    {
      "inputs": [
        {
          "internalType": "string",
          "name": "newName",
          "type": "string"
        }
      ],
      "name": "initializeV2",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "lostAndFound",
          "type": "address"
        }
      ],
      "name": "initializeV2_1",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`

task("deploy-mock-token")
    .addParam("name")
    .addParam("symbol")
    .addParam("decimals")
    .setAction(async (taskArgs, hre) => {
        if (taskArgs.name == "" || taskArgs.symbol == "" || taskArgs.decimals == "") {
            console.error(`params check failed,${taskArgs.name}, ${taskArgs.symbol}, ${taskArgs.decimals}`)
            return
        }
        const TokenFactory = await hre.ethers.getContractFactory("MockERC20")
        const token = await TokenFactory.deploy(
            taskArgs.name, // name
            taskArgs.symbol, // symbol
            taskArgs.decimals, // decimals
        )
        await token.deployed()
        console.log(`token deployed at ${token.address}`)
    })

task("deploy-l2-token")
    .addParam("proxyadmin")
    .addParam("name")
    .addParam("symbol")
    .addParam("decimals")
    .addParam("gateway")
    .addParam("counterpart")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.gateway) ||
            !ethers.utils.isAddress(taskArgs.counterpart)
        ) {
            console.error(`address params check failed,${taskArgs.proxyadmin}, ${taskArgs.gateway}, ${taskArgs.counterpart}`)
            return
        }

        if (taskArgs.name == "" || taskArgs.symbol == "" || taskArgs.decimals == "") {
            console.error(`params check failed,${taskArgs.name}, ${taskArgs.symbol}, ${taskArgs.decimals}`)
            return
        }

        // deploy token impl
        const TokenFactory = await hre.ethers.getContractFactory("MorphStandardERC20")
        const token = await TokenFactory.deploy()
        await token.deployed()
        console.log(`token deployed at ${token.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            token.address, //logic
            taskArgs.proxyadmin, //admin
            TokenFactory.interface.encodeFunctionData('initialize', [
                taskArgs.name, // name
                taskArgs.symbol, // symbol
                taskArgs.decimals, // decimals
                taskArgs.gateway, // gateway
                taskArgs.counterpart // counterpart
            ]) // data
        )
        await proxy.deployed()
        console.log(`proxy deployed at ${proxy.address}`)
    })

task("deploy-l2-wstETH")
    .addParam("proxyadmin")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin)) {
            console.error(`address params check failed,${taskArgs.proxyadmin}`)
            return
        }

        // deploy token impl
        const TokenFactory = await hre.ethers.getContractFactory("L2WstETHToken")
        const token = await TokenFactory.deploy()
        await token.deployed()
        console.log(`token deployed at ${token.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            token.address, //logic
            taskArgs.proxyadmin, //admin
            "0x"
        )
        await proxy.deployed()
        console.log(`proxy deployed at ${proxy.address}`)
    })

task("deploy-l1-lidogateway")
    .addParam("proxyadmin")
    .addParam("l1token")
    .addParam("l2token")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l1token) ||
            !ethers.utils.isAddress(taskArgs.l2token)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L1LidoGateway")
        const gateway = await GatewayFactory.deploy(taskArgs.l1token, taskArgs.l2token)
        await gateway.deployed()
        console.log(`gateway impl deployed at ${gateway.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            gateway.address, //logic
            taskArgs.proxyadmin, //admin
            '0x'
        )
        await proxy.deployed()
        console.log(`gateway proxy deployed at ${proxy.address}`)
    })

task("deploy-l2-lidogateway")
    .addParam("proxyadmin")
    .addParam("l1token")
    .addParam("l2token")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l1token) ||
            !ethers.utils.isAddress(taskArgs.l2token)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2LidoGateway")
        const gateway = await GatewayFactory.deploy(taskArgs.l1token, taskArgs.l2token)
        await gateway.deployed()
        console.log(`gateway impl deployed at ${gateway.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            gateway.address, //logic
            taskArgs.proxyadmin, //admin
            "0x"
        )
        await proxy.deployed()
        console.log(`gateway proxy deployed at ${proxy.address}`)
    })

task("init-l2-wstETH")
    .addParam("l2wsteth")
    .addParam("name")
    .addParam("symbol")
    .addParam("decimals")
    .addParam("gateway")
    .addParam("counterpart")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.l2wsteth) ||
            !ethers.utils.isAddress(taskArgs.gateway) ||
            !ethers.utils.isAddress(taskArgs.counterpart)
        ) {
            console.error(`address params check failed`)
            return
        }

        if (taskArgs.name == "" || taskArgs.symbol == "" || taskArgs.decimals == "") {
            console.error(`params check failed,${taskArgs.name}, ${taskArgs.symbol}, ${taskArgs.decimals}`)
            return
        }

        // deploy token impl
        const TokenFactory = await hre.ethers.getContractFactory("L2WstETHToken")
        const token = await TokenFactory.attach(taskArgs.l2wsteth)
        console.log(`token at ${token.address}`)

        let res = await token.initialize(
            taskArgs.name,
            taskArgs.symbol,
            taskArgs.decimals,
            taskArgs.gateway,
            taskArgs.counterpart,
        )
        let rec = await res.wait()
        console.log(`init ${rec.status == 1} , txHash ${rec.transactionHash}`)

        const name = await token.name()
        const symbol = await token.symbol()
        const decimals = await token.decimals()
        const gateway = await token.gateway()
        const counterpart = await token.counterpart()
        console.log(`name ${name}, symbol ${symbol}, decimals ${decimals}, gateway ${gateway}, counterpart ${counterpart}`)
    })

task("init-l1-lidogateway")
    .addParam("l1lidogateway")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .addParam("opter")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.l1lidogateway) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger) ||
            !ethers.utils.isAddress(taskArgs.opter)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L1LidoGateway")
        const gateway = await GatewayFactory.attach(taskArgs.l1lidogateway)
        console.log(`gateway at ${gateway.address}`)

        let res = await gateway.initialize(taskArgs.counterpart, taskArgs.router, taskArgs.messenger)
        let rec = await res.wait()
        console.log(`initialize ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        res = await gateway.initializeV2(taskArgs.opter, taskArgs.opter, taskArgs.opter, taskArgs.opter)
        rec = await res.wait()
        console.log(`initializeV2 ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        let counterpart = await gateway.counterpart()
        let router = await gateway.router()
        let messenger = await gateway.messenger()
        let depositEnabled = await gateway.isDepositsEnabled()
        let withdrawalEnabled = await gateway.isWithdrawalsEnabled()
        console.log(`counterpart ${counterpart}, router ${router}, messenger ${messenger}, depositEnabled ${depositEnabled}, withdrawalEnabled ${withdrawalEnabled}`)

        const DEPOSITS_ENABLER_ROLE = await gateway.DEPOSITS_ENABLER_ROLE()
        const DEPOSITS_DISABLER_ROLE = await gateway.DEPOSITS_DISABLER_ROLE()
        const WITHDRAWALS_ENABLER_ROLE = await gateway.WITHDRAWALS_ENABLER_ROLE()
        const WITHDRAWALS_DISABLER_ROLE = await gateway.WITHDRAWALS_DISABLER_ROLE()
        const depositE = await gateway.hasRole(DEPOSITS_ENABLER_ROLE, taskArgs.opter)
        const depositD = await gateway.hasRole(DEPOSITS_DISABLER_ROLE, taskArgs.opter)
        const withdrawE = await gateway.hasRole(WITHDRAWALS_ENABLER_ROLE, taskArgs.opter)
        const withdrawD = await gateway.hasRole(WITHDRAWALS_DISABLER_ROLE, taskArgs.opter)
        console.log(`Role: depositE ${depositE}, depositD ${depositD}, withdrawE ${withdrawE}, withdrawD ${withdrawD}`)
    })

task("init-l2-lidogateway")
    .addParam("l2lidogateway")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .addParam("opter")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.l2lidogateway) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger) ||
            !ethers.utils.isAddress(taskArgs.opter)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2LidoGateway")
        const gateway = await GatewayFactory.attach(taskArgs.l2lidogateway)
        console.log(`gateway at ${gateway.address}`)

        let res = await gateway.initialize(taskArgs.counterpart, taskArgs.router, taskArgs.messenger)
        let rec = await res.wait()
        console.log(`initialize ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        res = await gateway.initializeV2(taskArgs.opter, taskArgs.opter, taskArgs.opter, taskArgs.opter)
        rec = await res.wait()
        console.log(`initializeV2 ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        let counterpart = await gateway.counterpart()
        let router = await gateway.router()
        let messenger = await gateway.messenger()
        let depositEnabled = await gateway.isDepositsEnabled()
        let withdrawalEnabled = await gateway.isWithdrawalsEnabled()
        console.log(`counterpart ${counterpart}, router ${router}, messenger ${messenger}, depositEnabled ${depositEnabled}, withdrawalEnabled ${withdrawalEnabled}`)

        const DEPOSITS_ENABLER_ROLE = await gateway.DEPOSITS_ENABLER_ROLE()
        const DEPOSITS_DISABLER_ROLE = await gateway.DEPOSITS_DISABLER_ROLE()
        const WITHDRAWALS_ENABLER_ROLE = await gateway.WITHDRAWALS_ENABLER_ROLE()
        const WITHDRAWALS_DISABLER_ROLE = await gateway.WITHDRAWALS_DISABLER_ROLE()
        const depositE = await gateway.hasRole(DEPOSITS_ENABLER_ROLE, taskArgs.opter)
        const depositD = await gateway.hasRole(DEPOSITS_DISABLER_ROLE, taskArgs.opter)
        const withdrawE = await gateway.hasRole(WITHDRAWALS_ENABLER_ROLE, taskArgs.opter)
        const withdrawD = await gateway.hasRole(WITHDRAWALS_DISABLER_ROLE, taskArgs.opter)
        console.log(`Role: depositE ${depositE}, depositD ${depositD}, withdrawE ${withdrawE}, withdrawD ${withdrawD}`)
    })

task("upgrade-l1-usdcgateway")
    .addParam("proxyadmin")
    .addParam("l1usdcgatewayproxy")
    .addParam("l1token")
    .addParam("l2token")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l1usdcgatewayproxy) ||
            !ethers.utils.isAddress(taskArgs.l1token) ||
            !ethers.utils.isAddress(taskArgs.l2token) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L1USDCGateway")
        const gateway = await GatewayFactory.deploy(taskArgs.l1token, taskArgs.l2token)
        await gateway.deployed()
        console.log(`gateway impl deployed at ${gateway.address}`)

        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgrade(
            taskArgs.l1usdcgatewayproxy,
            gateway.address
        )
        let rec = await res.wait()
        console.log(`upgrade gateway ${rec.status == 1}`)


        const gatewayProxy = GatewayFactory.attach(taskArgs.l1usdcgatewayproxy)
        res = await gatewayProxy.initialize(taskArgs.counterpart, taskArgs.router, taskArgs.messenger)
        rec = await res.wait()
        console.log(`init gateway ${rec.status == 1}`)

        const owner = await gatewayProxy.owner()
        const counterpart = await gatewayProxy.counterpart()
        const router = await gatewayProxy.router()
        const messenger = await gatewayProxy.messenger()
        const l1USDC = await gatewayProxy.l1USDC()
        const l2USDC = await gatewayProxy.l2USDC()
        console.log(`owner ${owner}, gatewayProxy ${gatewayProxy.address}, counterpart ${counterpart}, router ${router}, messenger ${messenger}, l1USDC ${l1USDC}, l2USDC ${l2USDC}`)
    })

task("upgrade-l2-usdcgateway")
    .addParam("proxyadmin")
    .addParam("l2usdcgatewayproxy")
    .addParam("l1token")
    .addParam("l2token")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l2usdcgatewayproxy) ||
            !ethers.utils.isAddress(taskArgs.l1token) ||
            !ethers.utils.isAddress(taskArgs.l2token) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2USDCGateway")
        const gateway = await GatewayFactory.deploy(taskArgs.l1token, taskArgs.l2token)
        await gateway.deployed()
        console.log(`gateway impl deployed at ${gateway.address}`)

        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgrade(
            taskArgs.l2usdcgatewayproxy,
            gateway.address
        )
        let rec = await res.wait()
        console.log(`upgrade gateway ${rec.status == 1}`)

        const gatewayProxy = GatewayFactory.attach(taskArgs.l2usdcgatewayproxy)
        res = await gatewayProxy.initialize(taskArgs.counterpart, taskArgs.router, taskArgs.messenger)
        rec = await res.wait()
        console.log(`init gateway ${rec.status == 1}`)
        const counterpart = await gatewayProxy.counterpart()
        const router = await gatewayProxy.router()
        const l1USDC = await gatewayProxy.l1USDC()
        const l2USDC = await gatewayProxy.l2USDC()
        console.log(`gatewayProxy ${gatewayProxy.address}, counterpart ${counterpart}, router ${router}, l1USDC ${l1USDC}, l2USDC ${l2USDC}`)
    })

task("upgrade-l2-usdc-v2")
    .addParam("proxyadmin")
    .addParam("l2usdcproxy")
    .addParam("newimpl")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l2usdcproxy) ||
            !ethers.utils.isAddress(taskArgs.newimpl)
        ) {
            console.error(`address params check failed`)
            return
        }
        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgradeAndCall(
            taskArgs.l2usdcproxy,
            taskArgs.newimpl,
            new ethers.utils.Interface(V2_1ABI).encodeFunctionData("initializeV2", [
                "USD Coin"
            ])
        )
        let rec = await res.wait()
        console.log(`upgrade usdc v2 ${rec.status == 1}`)
    })

task("upgrade-l2-usdc-v2-1")
    .addParam("proxyadmin")
    .addParam("l2usdcproxy")
    .addParam("newimpl")
    .addParam("lostfund")
    .setAction(async (taskArgs, hre) => {
        const deployer = await hre.ethers.provider.getSigner();

        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l2usdcproxy) ||
            !ethers.utils.isAddress(taskArgs.newimpl) ||
            !ethers.utils.isAddress(taskArgs.lostfund)
        ) {
            console.error(`address params check failed`)
            return
        }
        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgrade(
            taskArgs.l2usdcproxy,
            taskArgs.newimpl
        )
        let rec = await res.wait()
        console.log(`upgrade usdc v2 ${rec.status == 1}`)

        const l2usdc = new ethers.Contract(taskArgs.l2usdcproxy, new ethers.utils.Interface(V2_1ABI))

        res = await l2usdc.connect(deployer).initializeV2("USD Coin")
        rec = await res.wait()
        console.log(`l2 usdc initializeV2 ${rec.status == 1}`)

        res = await l2usdc.connect(deployer).initializeV2_1(taskArgs.lostfund)
        rec = await res.wait()
        console.log(`l2 usdc initializeV2_1 ${rec.status == 1}`)
    })

task("deploy-l1-proxy")
    .addParam("proxyadmin")
    .addParam("empty")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.empty)
        ) {
            console.error(`address params check failed`)
            return
        }

        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            taskArgs.empty, //logic
            taskArgs.proxyadmin, //admin
            "0x"
        )
        await proxy.deployed()
        console.log(`proxy deployed at ${proxy.address}`)
    })


task("deploy-l2-usdcgateway")
    .addParam("proxyadmin")
    .addParam("l1token")
    .addParam("l2token")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.l1token) ||
            !ethers.utils.isAddress(taskArgs.l2token) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2USDCGateway")
        const gateway = await GatewayFactory.deploy(taskArgs.l1token, taskArgs.l2token)
        await gateway.deployed()
        console.log(`gateway impl deployed at ${gateway.address}`)

        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            gateway.address, //logic
            taskArgs.proxyadmin, //admin
            GatewayFactory.interface.encodeFunctionData("initialize", [
                taskArgs.counterpart,
                taskArgs.router,
                taskArgs.messenger
            ]
            )
        )
        await proxy.deployed()
        console.log(`gateway proxy deployed at ${proxy.address}`)

        const gatewayProxy = GatewayFactory.attach(proxy.address)
        const owner = await gatewayProxy.owner()
        const counterpart = await gatewayProxy.counterpart()
        const router = await gatewayProxy.router()
        const messenger = await gatewayProxy.messenger()
        const l1USDC = await gatewayProxy.l1USDC()
        const l2USDC = await gatewayProxy.l2USDC()
        console.log(`owner ${owner}, gatewayProxy ${gatewayProxy.address}, counterpart ${counterpart}, router ${router}, messenger ${messenger}, l1USDC ${l1USDC}, l2USDC ${l2USDC}`)
    })

task("deploy-l2-MigrationUSDC")
    .addParam("proxyadmin")
    .addParam("oldtoken")
    .addParam("newtoken")
    .addParam("recipient")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.oldtoken) ||
            !ethers.utils.isAddress(taskArgs.recipient) ||
            !ethers.utils.isAddress(taskArgs.newtoken)
        ) {
            console.error(`address params check failed`)
            return
        }
        // deploy gateway impl
        const ContractFactory = await hre.ethers.getContractFactory("MigrationUSDC")
        const contract = await ContractFactory.deploy(taskArgs.oldtoken, taskArgs.newtoken)
        await contract.deployed()
        console.log(`MigrationUSDC impl deployed at ${contract.address}`)

        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            contract.address, //logic
            taskArgs.proxyadmin, //admin
            ContractFactory.interface.encodeFunctionData("initialize", [
                taskArgs.recipient
            ]
            )
        )
        await proxy.deployed()
        console.log(`MigrationUSDC proxy deployed at ${proxy.address}`)

        const migrationProxy = ContractFactory.attach(proxy.address)
        const owner = await migrationProxy.owner()
        const oldtoken = await migrationProxy.OLD_USDC()
        const newtoken = await migrationProxy.NEW_USDC()
        console.log(`owner ${owner}, oldtoken ${oldtoken}, newtoken ${newtoken}`)
    })

task("deploy-l2-token-registry")
    .addParam("proxyadmin","Proxy admin address","0x530000000000000000000000000000000000000B")
    .addOptionalParam("proxy", "Existing proxy address (if upgrading)","0x5300000000000000000000000000000000000021")
    .addParam("owner")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.owner)
        ) {
            console.error(`address params check failed, proxyadmin: ${taskArgs.proxyadmin}, owner: ${taskArgs.owner}`)
            return
        }

        // deploy L2TokenRegistry impl
        const TokenRegistryFactory = await hre.ethers.getContractFactory("L2TokenRegistry")
        const tokenRegistry = await TokenRegistryFactory.deploy()
        await tokenRegistry.deployed()
        console.log(`L2TokenRegistry impl deployed at ${tokenRegistry.address}`)

        let proxyAddress;

        // Check if proxy parameter exists for upgrade
        if (taskArgs.proxy && ethers.utils.isAddress(taskArgs.proxy)) {
            console.log(`\nUpgrading existing proxy at ${taskArgs.proxy}`)
            
            // Get ProxyAdmin contract
            const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
            const proxyAdmin = ProxyAdminFactory.attach(taskArgs.proxyadmin)
            
            // Upgrade the proxy to new implementation
            const upgradeTx = await proxyAdmin.upgradeAndCall(
                taskArgs.proxy,
                tokenRegistry.address,
                TokenRegistryFactory.interface.encodeFunctionData('initialize', [
                    taskArgs.owner // owner
                ]) // data
            )
            await upgradeTx.wait()
            console.log(`Proxy upgraded to new implementation: ${tokenRegistry.address}`)
            
            proxyAddress = taskArgs.proxy
        } else {
            console.log(`\nDeploying new proxy`)
            
            // deploy proxy with initialize
            const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
            const proxy = await TransparentProxyFactory.deploy(
                tokenRegistry.address, //logic
                taskArgs.proxyadmin, //admin
                TokenRegistryFactory.interface.encodeFunctionData('initialize', [
                    taskArgs.owner // owner
                ]) // data
            )
            await proxy.deployed()
            console.log(`L2TokenRegistry proxy deployed at ${proxy.address}`)
            
            proxyAddress = proxy.address
        }

        // Verify deployment
        const tokenRegistryProxy = TokenRegistryFactory.attach(proxyAddress)
        const registryOwner = await tokenRegistryProxy.owner()
        const allowListEnabled = await tokenRegistryProxy.allowListEnabled()
        console.log(`\nL2TokenRegistry proxy address: ${proxyAddress}`)
        console.log(`L2TokenRegistry proxy owner: ${registryOwner}`)
        console.log(`L2TokenRegistry allowListEnabled: ${allowListEnabled}`)
    })

task("deploy-test-tokens-and-register")
    .addParam("tokenregistry","Contract address","0x5300000000000000000000000000000000000021")
    .addOptionalParam("count", "Number of test tokens to deploy", "10")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.tokenregistry)) {
            console.error(`tokenregistry address check failed: ${taskArgs.tokenregistry}`)
            return
        }

        const tokenCount = parseInt(taskArgs.count || "10")
        if (tokenCount < 1) {
            console.error(`token count should be at least 1, got: ${tokenCount}`)
            return
        }

        // Load token configurations from JSON file
        const tokensFilePath = path.join(__dirname, "../src/tokens/tokens.json")
        if (!fs.existsSync(tokensFilePath)) {
            console.error(`Tokens file not found: ${tokensFilePath}`)
            return
        }

        const allTokensData = JSON.parse(fs.readFileSync(tokensFilePath, "utf8"))
        if (!Array.isArray(allTokensData) || allTokensData.length === 0) {
            console.error(`Invalid tokens file format or empty tokens array`)
            return
        }

        console.log(`\n========================================`)
        console.log(`Connecting to L2TokenRegistry...`)
        console.log(`========================================\n`)

        // Connect to L2TokenRegistry
        const TokenRegistryFactory = await hre.ethers.getContractFactory("L2TokenRegistry")
        const tokenRegistry = TokenRegistryFactory.attach(taskArgs.tokenregistry)
        
        // Verify registry
        try {
            const registryOwner = await tokenRegistry.owner()
            console.log(`L2TokenRegistry address: ${taskArgs.tokenregistry}`)
            console.log(`L2TokenRegistry owner: ${registryOwner}`)
        } catch (error) {
            console.error(`Failed to connect to L2TokenRegistry: ${error}`)
            return
        }

        // Check which tokenIDs are already registered
        console.log(`\n========================================`)
        console.log(`Checking registered tokenIDs...`)
        console.log(`========================================\n`)

        const registeredTokenIDs = new Set<number>()
        const maxTokenID = Math.min(100, allTokensData.length)

        for (let tokenID = 1; tokenID <= maxTokenID; tokenID++) {
            try {
                const tokenInfo = await tokenRegistry.tokenRegistry(tokenID)
                if (tokenInfo.tokenAddress !== ethers.constants.AddressZero) {
                    registeredTokenIDs.add(tokenID)
                }
            } catch (error) {
                // If tokenID is not registered, tokenAddress will be zero address
            }
        }

        console.log(`Found ${registeredTokenIDs.size} already registered tokenIDs: ${Array.from(registeredTokenIDs).sort((a, b) => a - b).join(", ")}`)

        // Find the next available tokenID to start from
        let startTokenID = 1
        for (let i = 1; i <= maxTokenID; i++) {
            if (!registeredTokenIDs.has(i)) {
                startTokenID = i
                break
            }
        }

        if (registeredTokenIDs.size >= maxTokenID) {
            console.log(`\nAll ${maxTokenID} token slots are already registered. Nothing to deploy.`)
            return
        }

        // Check if total registered tokens would exceed 100
        if (registeredTokenIDs.size >= 100) {
            console.log(`\n⚠ Warning: Already have ${registeredTokenIDs.size} registered tokens. Cannot register more than 100 tokens.`)
            return
        }

        // Calculate how many tokens we can deploy (max 100 total)
        const availableSlots = Math.min(100 - registeredTokenIDs.size, maxTokenID - registeredTokenIDs.size)
        const tokensToDeployCount = Math.min(tokenCount, availableSlots)

        if (tokensToDeployCount === 0) {
            console.log(`No available token slots to deploy.`)
            return
        }

        console.log(`\nWill deploy ${tokensToDeployCount} tokens starting from tokenID ${startTokenID}`)
        console.log(`Total registered after deployment: ${registeredTokenIDs.size + tokensToDeployCount}/100`)

        // Get tokens to deploy
        const tokensToDeploy = []
        let currentTokenID = startTokenID
        let deployedCount = 0

        while (deployedCount < tokensToDeployCount && currentTokenID <= maxTokenID && registeredTokenIDs.size + deployedCount < 100) {
            if (!registeredTokenIDs.has(currentTokenID)) {
                const tokenData = allTokensData.find((t: any) => t.tokenID === currentTokenID)
                if (tokenData) {
                    tokensToDeploy.push({
                        ...tokenData,
                        tokenID: currentTokenID,
                        scale: ethers.BigNumber.from(tokenData.scale),
                        priceRatio: ethers.BigNumber.from(tokenData.priceRatio)
                    })
                    deployedCount++
                }
            }
            currentTokenID++
        }

        if (tokensToDeploy.length === 0) {
            console.log(`No tokens available to deploy.`)
            return
        }

        // Final check: ensure we don't exceed 100 tokens
        if (registeredTokenIDs.size + tokensToDeploy.length > 100) {
            const maxCanDeploy = 100 - registeredTokenIDs.size
            console.log(`\n⚠ Warning: Can only deploy ${maxCanDeploy} tokens to stay within 100 token limit.`)
            console.log(`Requested: ${tokensToDeploy.length}, Will deploy: ${maxCanDeploy}`)
            tokensToDeploy.splice(maxCanDeploy)
        }

        console.log(`\n========================================`)
        console.log(`Deploying ${tokensToDeploy.length} test tokens...`)
        console.log(`========================================\n`)

        const deployedTokens = []
        const TokenFactory = await hre.ethers.getContractFactory("MockERC20")

        // Deploy tokens
        for (let i = 0; i < tokensToDeploy.length; i++) {
            const config = tokensToDeploy[i]
            console.log(`[${i + 1}/${tokensToDeploy.length}] Deploying ${config.name} (${config.symbol}) - TokenID: ${config.tokenID}...`)
            
            const token = await TokenFactory.deploy(
                config.name,
                config.symbol,
                config.decimals
            )
            await token.deployed()
            
            console.log(`  ✓ Token deployed at: ${token.address}`)
            console.log(`  - Name: ${config.name}`)
            console.log(`  - Symbol: ${config.symbol}`)
            console.log(`  - Decimals: ${config.decimals}`)
            
            deployedTokens.push({
                ...config,
                address: token.address,
                contract: token
            })
        }

        console.log(`\n========================================`)
        console.log(`Registering tokens to L2TokenRegistry...`)
        console.log(`========================================\n`)

        // Prepare arrays for batch registration
        const tokenIDs: number[] = []
        const tokenAddresses: string[] = []
        const balanceSlots: string[] = []
        const scales: string[] = []

        for (const token of deployedTokens) {
            tokenIDs.push(token.tokenID)
            tokenAddresses.push(token.address)
            // Calculate balance slot for mapping(address => uint256) at slot 0
            // For MockERC20, balance mapping is typically at slot 0
            // The actual slot for a user's balance is keccak256(abi.encode(userAddress, slot))
            // Here we use slot 0 as the base slot
            balanceSlots.push(ethers.utils.hexZeroPad(ethers.BigNumber.from(token.balanceSlot).toHexString(), 32))
            scales.push(token.scale.toString())
        }

        console.log(`Registering ${tokenIDs.length} tokens in batch...`)
        console.log(`Token IDs: ${tokenIDs.join(", ")}`)
        console.log(`Token Addresses: ${tokenAddresses.join(", ")}`)

        try {
            // Batch register tokens
            const tx = await tokenRegistry.registerTokens(
                tokenIDs,
                tokenAddresses,
                balanceSlots,
                scales
            )
            console.log(`\n  ✓ Registration transaction sent: ${tx.hash}`)
            
            const receipt = await tx.wait()
            console.log(`  ✓ Transaction confirmed in block: ${receipt.blockNumber}`)
            console.log(`  ✓ Gas used: ${receipt.gasUsed.toString()}`)

            // Set prices for registered tokens
            console.log(`\n========================================`)
            console.log(`Setting prices for registered tokens...`)
            console.log(`========================================\n`)

            try {
                const priceTokenIDs: number[] = []
                const prices: string[] = []

                for (const token of deployedTokens) {
                    priceTokenIDs.push(token.tokenID)
                    prices.push(token.priceRatio.toString())
                }

                console.log(`Setting prices for ${priceTokenIDs.length} tokens...`)
                const priceTx = await tokenRegistry.batchUpdatePrices(priceTokenIDs, prices)
                console.log(`  ✓ Price update transaction sent: ${priceTx.hash}`)
                
                const priceReceipt = await priceTx.wait()
                console.log(`  ✓ Prices confirmed in block: ${priceReceipt.blockNumber}`)
                console.log(`  ✓ Gas used: ${priceReceipt.gasUsed.toString()}`)

                // Display price information
                console.log(`\nPrice information:`)
                for (const token of deployedTokens) {
                    const price = await tokenRegistry.priceRatio(token.tokenID)
                    console.log(`  ${token.symbol} (ID: ${token.tokenID}): ${price.toString()}`)
                }
            } catch (priceError) {
                console.error(`\n⚠ Failed to set prices: ${priceError}`)
                console.log(`Attempting individual price updates...\n`)
                
                // Fallback to individual price updates
                for (const token of deployedTokens) {
                    try {
                        const priceTx = await tokenRegistry.updatePriceRatio(token.tokenID, token.priceRatio)
                        const priceReceipt = await priceTx.wait()
                        console.log(`  ✓ ${token.symbol} price set in block: ${priceReceipt.blockNumber}`)
                    } catch (err) {
                        console.error(`  ✗ Failed to set price for ${token.symbol}: ${err}`)
                    }
                }
            }

            // Verify registration
            console.log(`\n========================================`)
            console.log(`Verifying token registrations...`)
            console.log(`========================================\n`)

            for (const token of deployedTokens) {
                try {
                    const tokenInfo = await tokenRegistry.tokenRegistry(token.tokenID)
                    const registeredTokenID = await tokenRegistry.tokenRegistration(token.address)
                    const priceRatio = await tokenRegistry.priceRatio(token.tokenID)
                    
                    console.log(`Token ID ${token.tokenID} (${token.symbol}):`)
                    console.log(`  - Address: ${tokenInfo.tokenAddress}`)
                    console.log(`  - Balance Slot: ${tokenInfo.balanceSlot}`)
                    console.log(`  - Is Active: ${tokenInfo.isActive}`)
                    console.log(`  - Decimals: ${tokenInfo.decimals}`)
                    console.log(`  - Scale: ${tokenInfo.scale.toString()}`)
                    console.log(`  - Price Ratio: ${priceRatio.toString()}`)
                    console.log(`  - Registered TokenID: ${registeredTokenID}`)
                    console.log(`  ✓ Registration verified\n`)
                } catch (error) {
                    console.error(`  ✗ Failed to verify token ${token.tokenID}: ${error}\n`)
                }
            }

            console.log(`\n========================================`)
            console.log(`Summary:`)
            console.log(`========================================`)
            console.log(`Total tokens deployed: ${deployedTokens.length}`)
            console.log(`Total tokens registered: ${tokenIDs.length}`)
            console.log(`\nToken addresses:`)
            for (const token of deployedTokens) {
                console.log(`  ${token.symbol} (ID: ${token.tokenID}): ${token.address}`)
            }
            console.log(`\nL2TokenRegistry: ${taskArgs.tokenregistry}`)
            console.log(`========================================\n`)

        } catch (error) {
            console.error(`\n✗ Failed to register tokens: ${error}`)
            console.log(`\nTrying individual registration instead...\n`)
            
            // Fallback to individual registration
            for (const token of deployedTokens) {
                try {
                    console.log(`Registering ${token.symbol} (ID: ${token.tokenID}) individually...`)
                    const balanceSlot = ethers.utils.hexZeroPad(ethers.BigNumber.from(token.balanceSlot).toHexString(), 32)
                    const tx = await tokenRegistry.registerToken(
                        token.tokenID,
                        token.address,
                        balanceSlot,
                        token.scale
                    )
                    const receipt = await tx.wait()
                    console.log(`  ✓ ${token.symbol} registered in block: ${receipt.blockNumber}`)
                    
                    // Set price after registration
                    try {
                        const priceTx = await tokenRegistry.updatePriceRatio(token.tokenID, token.priceRatio)
                        const priceReceipt = await priceTx.wait()
                        console.log(`  ✓ ${token.symbol} price set in block: ${priceReceipt.blockNumber}\n`)
                    } catch (priceErr) {
                        console.error(`  ⚠ Failed to set price for ${token.symbol}: ${priceErr}\n`)
                    }
                } catch (err) {
                    console.error(`  ✗ Failed to register ${token.symbol}: ${err}\n`)
                }
            }
        }
    })