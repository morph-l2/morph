import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";
import { getContractAddressByName } from "../src/deploy-utils";
import { validateGenesisHeader } from "../src/deployment-validation";
import { ProxyStorageName as P, ContractFactoryName as F, ImplStorageName as I } from "../src/types";

export const ContractInit = async (
    hre: HardhatRuntimeEnvironment, path: string, deployer: any, config: any
): Promise<string> => {
    validateGenesisHeader(config.batchHeader);
    const address = (name: string) => getContractAddressByName(path, name);
    const rollup = await hre.ethers.getContractAt(F.Rollup, address(P.RollupProxyStorageName), deployer);
    const genesisHash = ethers.utils.keccak256(config.batchHeader);
    const existingHash = await rollup.committedBatches(0);
    if (existingHash !== ethers.constants.HashZero && existingHash !== genesisHash) {
        throw new Error("Rollup.committedBatches(0) does not match batchHeader; initialization stopped");
    }
    const delay = await rollup.rollupDelayPeriod();
    if (!delay.isZero() && !delay.eq(config.rollupDelayPeriod)) {
        throw new Error("Rollup.rollupDelayPeriod does not match configuration; initialization stopped");
    }
    if (existingHash === ethers.constants.HashZero) {
        await (await rollup.importGenesisBatch(config.batchHeader)).wait();
    }
    if (!(await rollup.isChallenger(config.rollupChallenger))) {
        await (await rollup.addChallenger(config.rollupChallenger)).wait();
    }
    // importGenesisBatch also writes committedStateRoots; the legacy initialize2 call is unnecessary.
    if (delay.isZero()) await (await rollup.initialize3(config.rollupDelayPeriod)).wait();

    const queue = await hre.ethers.getContractAt(F.L1MessageQueueWithGasPriceOracle, address(P.L1MessageQueueWithGasPriceOracleProxyStorageName), deployer);
    const baseFee = ethers.utils.parseUnits(config.l2BaseFee.toString(), "gwei");
    if (!(await queue.l2BaseFee()).eq(baseFee)) await (await queue.setL2BaseFee(baseFee)).wait();

    const weth = address(I.WETH);
    const gateway = address(P.L1WETHGatewayProxyStorageName);
    const router = await hre.ethers.getContractAt(F.L1GatewayRouter, address(P.L1GatewayRouterProxyStorageName), deployer);
    if ((await router.getERC20Gateway(weth)).toLowerCase() !== gateway.toLowerCase()) {
        await (await router.setERC20Gateway([weth], [gateway])).wait();
    }
    return "";
};

export default ContractInit;
