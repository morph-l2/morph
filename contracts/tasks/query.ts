import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";

task("getLastFinalizedBatchIndex")
    .setAction(async (taskArgs, hre) => {
        const Factort = await hre.ethers.getContractFactory('Rollup')
        const rollup = Factort.attach('0xb7f8bc63bbcad18155201308c8f3540b07f84f5e')
        const latestFinalizeIndex = await rollup.lastFinalizedBatchIndex()
        console.log(latestFinalizeIndex)
    });
