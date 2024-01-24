import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";

task("upgradeProxy")
    .addParam("proxyaddr")
    .addParam("newimpladdr")
    .setAction(async (taskArgs, hre) => {
        const ProxyFactory = await hre.ethers.getContractFactory('Proxy')
        const proxy = ProxyFactory.attach(taskArgs.proxyaddr)
        console.log("before upgrade the impl contract is :", await proxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
        const res = await proxy.upgradeTo(taskArgs.newimpladdr)
        const receipt = await res.wait()
        console.log(`receipt status : ${receipt.status}`)
        console.log("after upgrade the impl contract is :", await proxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
    });
