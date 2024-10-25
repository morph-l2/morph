import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";
import {predeploys} from "../src";

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

        let res = await gateway.initialize(taskArgs.counterpart,taskArgs.router,taskArgs.messenger)
        let rec = await res.wait()
        console.log(`initialize ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        res = await gateway.initializeV2(taskArgs.opter,taskArgs.opter,taskArgs.opter,taskArgs.opter)
        rec = await res.wait()
        console.log(`initializeV2 ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        let counterpart = await gateway.counterpart()
        let router = await gateway.router()
        let messenger = await gateway.messenger()
        let depositEnabled = await gateway.isDepositsEnabled()
        let withdrawalEnabled = await gateway.isWithdrawalsEnabled()
        console.log(`counterpart ${counterpart}, router ${router}, messenger ${messenger}, depositEnabled ${depositEnabled}, withdrawalEnabled ${withdrawalEnabled}`)

        const DEPOSITS_ENABLER_ROLE =await gateway.DEPOSITS_ENABLER_ROLE()
        const DEPOSITS_DISABLER_ROLE =await gateway.DEPOSITS_DISABLER_ROLE()
        const WITHDRAWALS_ENABLER_ROLE =await gateway.WITHDRAWALS_ENABLER_ROLE()
        const WITHDRAWALS_DISABLER_ROLE =await gateway.WITHDRAWALS_DISABLER_ROLE()
        const depositE = await gateway.hasRole(DEPOSITS_ENABLER_ROLE,taskArgs.opter)
        const depositD = await gateway.hasRole(DEPOSITS_DISABLER_ROLE,taskArgs.opter)
        const withdrawE = await gateway.hasRole(WITHDRAWALS_ENABLER_ROLE,taskArgs.opter)
        const withdrawD = await gateway.hasRole(WITHDRAWALS_DISABLER_ROLE,taskArgs.opter)
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

        let res = await gateway.initialize(taskArgs.counterpart,taskArgs.router,taskArgs.messenger)
        let rec = await res.wait()
        console.log(`initialize ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        res = await gateway.initializeV2(taskArgs.opter,taskArgs.opter,taskArgs.opter,taskArgs.opter)
        rec = await res.wait()
        console.log(`initializeV2 ${rec.status == 1}, txHash: ${rec.transactionHash}`)

        let counterpart = await gateway.counterpart()
        let router = await gateway.router()
        let messenger = await gateway.messenger()
        let depositEnabled = await gateway.isDepositsEnabled()
        let withdrawalEnabled = await gateway.isWithdrawalsEnabled()
        console.log(`counterpart ${counterpart}, router ${router}, messenger ${messenger}, depositEnabled ${depositEnabled}, withdrawalEnabled ${withdrawalEnabled}`)

        const DEPOSITS_ENABLER_ROLE =await gateway.DEPOSITS_ENABLER_ROLE()
        const DEPOSITS_DISABLER_ROLE =await gateway.DEPOSITS_DISABLER_ROLE()
        const WITHDRAWALS_ENABLER_ROLE =await gateway.WITHDRAWALS_ENABLER_ROLE()
        const WITHDRAWALS_DISABLER_ROLE =await gateway.WITHDRAWALS_DISABLER_ROLE()
        const depositE = await gateway.hasRole(DEPOSITS_ENABLER_ROLE,taskArgs.opter)
        const depositD = await gateway.hasRole(DEPOSITS_DISABLER_ROLE,taskArgs.opter)
        const withdrawE = await gateway.hasRole(WITHDRAWALS_ENABLER_ROLE,taskArgs.opter)
        const withdrawD = await gateway.hasRole(WITHDRAWALS_DISABLER_ROLE,taskArgs.opter)
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

            let res = await proxyAdmin.upgradeAndCall(
                taskArgs.l1usdcgatewayproxy,
                gateway.address,
                GatewayFactory.interface.encodeFunctionData("initialize", [
                            taskArgs.counterpart,
                            taskArgs.router,
                            taskArgs.messenger
                    ]
                )
            )
            let rec =await res.wait()
            console.log(`upgrade gateway ${rec.status == 1}`)

            const gatewayProxy = GatewayFactory.attach(taskArgs.l1usdcgatewayproxy)
            const counterpart = await gatewayProxy.counterpart()
            const router = await gatewayProxy.router()
            const l1USDC = await gatewayProxy.l1USDC()
            const l2USDC = await gatewayProxy.l2USDC()
            console.log(`gatewayProxy ${gatewayProxy.address}, counterpart ${counterpart}, router ${router}, l1USDC ${l1USDC}, l2USDC ${l2USDC}`)
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

            let res = await proxyAdmin.upgradeAndCall(
                taskArgs.l2usdcgatewayproxy,
                gateway.address,
                GatewayFactory.interface.encodeFunctionData("initialize", [
                            taskArgs.counterpart,
                            taskArgs.router,
                            taskArgs.messenger
                    ]
                )
            )
            let rec =await res.wait()
            console.log(`upgrade gateway ${rec.status == 1}`)

            const gatewayProxy = GatewayFactory.attach(taskArgs.l2usdcgatewayproxy)
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
            new ethers.utils.Interface(V2_1ABI).encodeFunctionData("initializeV2",[
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

        const l2usdc = new ethers.Contract(taskArgs.l2usdcproxy,new ethers.utils.Interface(V2_1ABI))

        res = await l2usdc.connect(deployer).initializeV2("USD Coin")
        rec = await res.wait()
        console.log(`l2 usdc initializeV2 ${rec.status == 1}`)

        res = await l2usdc.connect(deployer).initializeV2_1(taskArgs.lostfund)
        rec = await res.wait()
        console.log(`l2 usdc initializeV2_1 ${rec.status == 1}`)
    })

task("deploy-l1-customgateway")
    .addParam("proxyadmin")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L1CustomERC20Gateway")
        const gateway = await GatewayFactory.deploy()
        await gateway.deployed()
        console.log(`L1CustomERC20Gateway impl deployed at ${gateway.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            gateway.address, //logic
            taskArgs.proxyadmin, //admin
            gateway.interface.encodeFunctionData("initialize",[
                taskArgs.counterpart,
                taskArgs.router,
                taskArgs.messenger
            ])
        )
        await proxy.deployed()
        console.log(`L1CustomERC20Gateway proxy deployed at ${proxy.address}`)
    })

task("deploy-l2-customgateway")
    .addParam("proxyadmin")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2CustomERC20Gateway")
        const gateway = await GatewayFactory.deploy()
        await gateway.deployed()
        console.log(`L2CustomERC20Gateway impl deployed at ${gateway.address}`)

        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgradeAndCall(
            predeploys.L2CustomERC20Gateway,
            gateway.address,
            GatewayFactory.interface.encodeFunctionData("initialize", [
                    taskArgs.counterpart,
                    taskArgs.router,
                    taskArgs.messenger
                ]
            )
        )
        let rec = await res.wait()
        console.log(`proxy admin upgrade ${rec.status == 1}`)
    })

task("deploy-l2-withdrawlockgateway")
    .addParam("proxyadmin")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2WithdrawLockERC20Gateway")
        const gateway = await GatewayFactory.deploy()
        await gateway.deployed()
        console.log(`L2WithdrawLockERC20Gateway impl deployed at ${gateway.address}`)

        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgradeAndCall(
            predeploys.L2WithdrawLockERC20Gateway,
            gateway.address,
            GatewayFactory.interface.encodeFunctionData("initialize", [
                    taskArgs.counterpart,
                    taskArgs.router,
                    taskArgs.messenger
                ]
            )
        )
        let rec = await res.wait()
        console.log(`proxy admin upgrade ${rec.status == 1}`)
    })

task("deploy-l1-reversegateway")
    .addParam("proxyadmin")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L1ReverseCustomGateway")
        const gateway = await GatewayFactory.deploy()
        await gateway.deployed()
        console.log(`L1ReverseCustomGateway impl deployed at ${gateway.address}`)

        // deploy proxy with initialize
        const TransparentProxyFactory = await hre.ethers.getContractFactory("TransparentUpgradeableProxy")
        const proxy = await TransparentProxyFactory.deploy(
            gateway.address, //logic
            taskArgs.proxyadmin, //admin
            gateway.interface.encodeFunctionData("initialize",[
                taskArgs.counterpart,
                taskArgs.router,
                taskArgs.messenger
            ])
        )
        await proxy.deployed()
        console.log(`L1ReverseCustomGateway proxy deployed at ${proxy.address}`)
    })

task("deploy-l2-reversegateway")
    .addParam("proxyadmin")
    .addParam("counterpart")
    .addParam("router")
    .addParam("messenger")
    .setAction(async (taskArgs, hre) => {
        // params check
        if (!ethers.utils.isAddress(taskArgs.proxyadmin) ||
            !ethers.utils.isAddress(taskArgs.counterpart) ||
            !ethers.utils.isAddress(taskArgs.router) ||
            !ethers.utils.isAddress(taskArgs.messenger)
        ) {
            console.error(`address params check failed`)
            return
        }

        // deploy gateway impl
        const GatewayFactory = await hre.ethers.getContractFactory("L2ReverseCustomGateway")
        const gateway = await GatewayFactory.deploy()
        await gateway.deployed()
        console.log(`L2ReverseCustomGateway impl deployed at ${gateway.address}`)

        // upgrade proxy with initialize
        const ProxyAdminFactory = await hre.ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await ProxyAdminFactory.attach(taskArgs.proxyadmin)
        console.log(`proxy admin at ${proxyAdmin.address}`)

        let res = await proxyAdmin.upgradeAndCall(
            predeploys.L2ReverseERC20Gateway,
            gateway.address,
            GatewayFactory.interface.encodeFunctionData("initialize", [
                    taskArgs.counterpart,
                    taskArgs.router,
                    taskArgs.messenger
                ]
            )
        )
        let rec = await res.wait()
        console.log(`proxy admin upgrade ${rec.status == 1}`)
    })
