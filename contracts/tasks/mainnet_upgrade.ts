import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ContractFactoryName } from "../src/types";
import { ethers } from "ethers";

task("rollup-upgrade-hc")
    .addParam("l2cid")
    .addParam("prevStateRoot")
    .setAction(async (taskArgs, hre) => {
        const chainId = taskArgs.l2cid

        const RollupFactoryName = ContractFactoryName.Rollup

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const rollupNewImpl = await RollupFactory.deploy(chainId)
        await rollupNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`Rollup new impl deploy at ${rollupNewImpl.address} and height ${blockNumber}`)

        console.log("initialize2 abi code: ", RollupFactory.interface.encodeFunctionData('initialize2', [
            taskArgs.prevStateRoot,
        ]))
    })


task("deploy-morph-placement-token")
    .setAction(async (taskArgs, hre) => {
        const MorphPlacementTokenFactory = await hre.ethers.getContractFactory("MorphPlacementToken");
        const morphPlacementToken = await MorphPlacementTokenFactory.deploy();
        await morphPlacementToken.deployed();

        let blockNumber = await hre.ethers.provider.getBlockNumber();
        console.log(`MorphPlacementToken deployed at ${morphPlacementToken.address} at block ${blockNumber}`);
    });


task("init-morph-placement-token")
    .addParam("proxyaddr")
    .addParam("owner")
    .setAction(async (taskArgs, hre) => {

        if (!hre.ethers.utils.isAddress(taskArgs.owner) || !hre.ethers.utils.isAddress(taskArgs.proxyaddr)) {
            throw new Error("address invalid");
        }

        const MorphPlacementTokenFactory = await hre.ethers.getContractFactory("MorphPlacementToken");
        const token = MorphPlacementTokenFactory.attach(taskArgs.proxyaddr)

        let res = await token.initialize(
            "Morph Placement Token",
            "MPHP",
            taskArgs.owner,
            hre.ethers.utils.parseEther("10000000000")
        )
        let rec = await res.wait()

        console.log(`MorphPlacementToken init at ${token.address} and initialized ${rec.status == 1}, at blockNum ${rec.blockNumber}`);

        let name = await token.name()
        let symbol = await token.symbol()
        let totalSupply = await token.totalSupply()
        let owner = await token.owner()
        console.log(`init name ${name}, symbol ${symbol}, totalSupply ${totalSupply}, owner ${owner}`)
    });


task("deploy-st-token")
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