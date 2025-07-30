import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ContractFactoryName } from "../src/types";

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
        console.log(`MorphPlacementToken deployed at ${morphPlacementToken.address} and initialized at block ${blockNumber}`);
    });
  
