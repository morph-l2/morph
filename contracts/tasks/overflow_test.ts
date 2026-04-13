import { task } from "hardhat/config";
import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import { HardhatRuntimeEnvironment } from 'hardhat/types';
import { ethers } from "ethers";
import { waffle } from "hardhat";

export const deployL1OverflowTester = async (
    hre: HardhatRuntimeEnvironment,
    deployer: any,
    l1MessengerAddr: any,
    l2TesterAddr: any,
    gasLimit: any,
): Promise<string> => {
    const Factory = await hre.ethers.getContractFactory("L1OverflowTester", deployer)
    const contract = await Factory.deploy(l1MessengerAddr, l2TesterAddr, gasLimit)
    await contract.deployed()
    console.log("L1OverflowTester=%s ; TX_HASH: %s", contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    return ''
}

export const deployL2OverflowTester = async (
    hre: HardhatRuntimeEnvironment,
    deployer: any,
): Promise<string> => {
    const Factory = await hre.ethers.getContractFactory("L2OverflowTester", deployer)
    const contract = await Factory.deploy()
    await contract.deployed()
    console.log("L2OverflowTester=%s ; TX_HASH: %s", contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    return ''
}

task("deployL2Tester")
    .setAction(async (taskArgs, hre) => {
        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('deloying L2OverflowTester......')
        await deployL2OverflowTester(hre, deployer)
    })

task("deployL1Tester")
    .addParam("l1messenger")
    .addParam("l2tester")
    .setAction(async (taskArgs, hre) => {
        const l1Messenger = taskArgs.l1messenger
        const l2Tester = taskArgs.l2tester

        const deployer = await hre.ethers.provider.getSigner();
        console.log('################################## console parameters ##################################')
        console.log('deployer :', await deployer.getAddress())

        console.log('deloying L1OverflowTester......')
        await deployL1OverflowTester(hre, deployer, l1Messenger, l2Tester, 2000000)
    })

task("hash")
    .addParam("contractaddr")
    .addParam("message")
    .addParam("count")
    .setAction(async (taskArgs, hre) => {
        const addr = taskArgs.contractaddr
        const message = taskArgs.message
        const count = taskArgs.count

        const deployer = await hre.ethers.provider.getSigner();
        const MyContract = await hre.ethers.getContractFactory("L2OverflowTester", deployer);
        const contract = MyContract.attach(addr);

        console.log("sending tx for hash......")
        let txn = await contract.hash(message, count)
        await txn.wait();
        console.log(txn)

    })

task("crossHash")
    .addParam("contractaddr")
    .addParam("message")
    .addParam("count")
    .setAction(async (taskArgs, hre) => {
        const addr = taskArgs.contractaddr
        const message = taskArgs.message
        const count = taskArgs.count

        const deployer = await hre.ethers.provider.getSigner();
        const MyContract = await hre.ethers.getContractFactory("L1OverflowTester", deployer);
        const contract = MyContract.attach(addr);

        console.log("sending tx for crossHash......")
        //const options = {value: ethers.utils.parseEther("1.0")}
        const txn = await contract["crossHash(string,uint256)"](message, count, { value: ethers.utils.parseEther("0.002") })

        //let txn = await contract.crossHash(message, count, options)
        await txn.wait();
        console.log(txn)

    })

task("updateLimit")
    .addParam("contractaddr")
    .addParam("gaslimit")
    .setAction(async (taskArgs, hre) => {
        const addr = taskArgs.contractaddr
        const gasLimit = taskArgs.gaslimit
        const deployer = await hre.ethers.provider.getSigner();
        const MyContract = await hre.ethers.getContractFactory("L1OverflowTester", deployer);
        const contract = MyContract.attach(addr);

        console.log("sending tx for updateGasLimit......");
        let txn = await contract.updateGasLimit(gasLimit);
        await txn.wait();
        console.log(txn);
    })

task("printResult")
    .addParam("contractaddr")
    .setAction(async (taskArgs, hre) => {
        const addr = taskArgs.contractaddr
        const MyContract = await hre.ethers.getContractFactory("L2OverflowTester");
        const contract = MyContract.attach(addr);

        let messageHash = await contract.getMessageHash()
        let count = await contract.getHashCount()
        console.log("messageHash: ", messageHash)
        console.log("hashCount: ", count)
    })

task("upgradeRollup")
    .addParam("proxyaddr")
    .setAction(async (taskArgs, hre) => {
        const addr = taskArgs.proxyaddr
        const RollupProxyFactory = await hre.ethers.getContractFactory("Proxy")
        const RollupProxyWithSigner = RollupProxyFactory.attach(addr)

        const RollupFactory = await hre.ethers.getContractFactory("Rollup");
        const rollupNewImpl = await RollupFactory.deploy(53077);
        await rollupNewImpl.deployed()
        console.log("new impl contract address: ", rollupNewImpl.address)

        const RollupProxy = new ethers.Contract(
            addr,
            RollupProxyFactory.interface,
            hre.waffle.provider,
        )

        console.log("before upgrade the impl contract is :", await RollupProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))

        const upgradeTx = await RollupProxyWithSigner.upgradeTo(rollupNewImpl.address)
        console.log("upgradeTx hash: ", upgradeTx.hash)
        const recipt = await upgradeTx.wait()
        console.log("recipt status: ", recipt.status)

        console.log("after upgrade the impl contract is :", await RollupProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        }))
    })
