import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task } from "hardhat/config";
import { ethers } from "ethers";

// yarn hardhat upgradeL2StakingProxy --l1staking 0x5dA9472d947eE1824Db10dEb8d69Af703964c695 --network qanetl2
task("upgradeL2StakingProxy")
    .addParam("l1staking")
    .setAction(async (taskArgs, hre) => {
        const l1stakingAddr = taskArgs.l1staking
        const l2stakingproxyaddr = "0x5300000000000000000000000000000000000015"

        const ProxyAdminFactory = await hre.ethers.getContractFactory('ProxyAdmin')
        const proxyAdmin = ProxyAdminFactory.attach("0x530000000000000000000000000000000000000B")

        // upgrade
        const L2StakingFactory = await hre.ethers.getContractFactory("L2Staking");
        const l2StakingNewImpl = await L2StakingFactory.deploy(l1stakingAddr);
        await l2StakingNewImpl.deployed()
        console.log("new l2StakingNewImpl contract address: ", l2StakingNewImpl.address)

        if (!hre.ethers.utils.isAddress(l2stakingproxyaddr) || !hre.ethers.utils.isAddress(l2StakingNewImpl.address)) {
            console.log(`not address ${l2stakingproxyaddr} ${l2StakingNewImpl.address}`)
            return
        }

        const res = await proxyAdmin.upgrade(l2stakingproxyaddr, l2StakingNewImpl.address)
        const rec = await res.wait()

        console.log(`upgrade l2Staking ${rec.status === 1}`)
    })