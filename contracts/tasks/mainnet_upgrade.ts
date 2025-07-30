import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ContractFactoryName } from "../src/types";
import { predeploys } from "../src"
import { assertContractVariable } from "../src/deploy-utils";

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
    .setAction(async (taskArgs, hre) => {
        const MorphPlacementTokenFactory = await hre.ethers.getContractFactory("MorphPlacementToken");
        const token = MorphPlacementTokenFactory.attach(predeploys.MorphToken)

        let res = await token.initialize(
            "morph placement token",
            "mphp",
            "0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E",
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
