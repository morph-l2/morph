import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ContractFactoryName } from "../src/types";

task("rollup-upgrade-hc")
    .addParam("l1pa")
    .addParam("l2cid")
    .addParam("rollup")
    .addParam("prevStateRoot")
    .setAction(async (taskArgs, hre) => {
        // const ProxyAdminImplAddr = taskArgs.l1pa
        const chainId = taskArgs.l2cid
        // const RollupProxyAddr = taskArgs.rollup
        // const deployer = hre.ethers.provider.getSigner()

        const RollupFactoryName = ContractFactoryName.Rollup

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const rollupNewImpl = await RollupFactory.deploy(chainId)
        await rollupNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`Rollup new impl deploy at ${rollupNewImpl.address} and height ${blockNumber}`)

        // const ProxyAdminFactoryName = ContractFactoryName.ProxyAdmin
        // const ProxyAdmin = await hre.ethers.getContractAt(ProxyAdminFactoryName, ProxyAdminImplAddr, deployer)

        console.log("initialize2 abi code: ", RollupFactory.interface.encodeFunctionData('initialize2', [
            taskArgs.prevStateRoot,
        ]))

        // const res = await ProxyAdmin.upgradeAndCall(
        //     RollupProxyAddr,
        //     rollupNewImpl.address,
        //     RollupFactory.interface.encodeFunctionData('initialize2', [
        //         taskArgs.prevStateRoot,
        //     ])
        // )
        // const rec = await res.wait()

        // console.log(`upgrade rollup ${rec.status == 1 ? "success" : "failed"}`)
    })
