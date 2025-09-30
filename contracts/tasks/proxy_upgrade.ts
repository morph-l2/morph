import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";

task("upgradeProxyWithPorxyAdmin")
    .addParam("proxyadminaddr")
    .addParam("proxyaddr")
    .addParam("newimpladdr")
    .setAction(async (taskArgs, hre) => {
        if (!hre.ethers.utils.isAddress(taskArgs.proxyadminaddr) || !hre.ethers.utils.isAddress(taskArgs.proxyaddr)
            || !hre.ethers.utils.isAddress(taskArgs.newimpladdr)) {
            throw new Error("token address invalid");
        }

        const proxyAdmin = await hre.ethers.getContractAt('ProxyAdmin', taskArgs.proxyadminaddr)

        const proxy = await hre.ethers.getContractAt('ITransparentUpgradeableProxy', taskArgs.proxyaddr)
        console.log("before upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: taskArgs.proxyadminaddr,
        }))

        const res = await proxyAdmin.upgrade(taskArgs.proxyaddr, taskArgs.newimpladdr)
        const receipt = await res.wait()
        console.log(`receipt status : ${receipt.status}`)
        console.log("after upgrade the impl contract is :", await proxy.connect(hre.waffle.provider).callStatic.implementation({
            from: taskArgs.proxyadminaddr,
        }))
    });