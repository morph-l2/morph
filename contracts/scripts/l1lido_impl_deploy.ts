const hre = require("hardhat");
import 'dotenv/config';
const { ethers } = hre;

const L1WstETHAddr = process.env.L1WSTETH;
const L2WstETHAddr = process.env.L2WSTETH;

async function main() {
    if (!ethers.utils.isAddress(L1WstETHAddr) || !ethers.utils.isAddress(L2WstETHAddr)) {
        throw new Error("token address invalid");
    }

    const L1LidoGatewayFactory = await ethers.getContractFactory("L1LidoGateway");

    const l1LidoGatewayImpl = await L1LidoGatewayFactory.deploy(L1WstETHAddr, L2WstETHAddr)
    await l1LidoGatewayImpl.deployed()
    console.log(`impl deployed l1LidoGatewayImpl ${l1LidoGatewayImpl.address}`)
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});