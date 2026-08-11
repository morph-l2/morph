import "@nomiclabs/hardhat-ethers";

import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";

import { getContractAddressByName, storage } from "../src/deploy-utils";
import { ContractFactoryName, ImplStorageName, ProxyStorageName } from "../src/types";

const IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";

export const SubmitterInstall = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    try {
        const proxyAddress = getContractAddressByName(path, ProxyStorageName.SubmitterProxyStorageName);
        const implementationAddress = getContractAddressByName(path, ImplStorageName.SubmitterStorageName);
        const emptyImplementation = getContractAddressByName(path, ImplStorageName.EmptyContract);
        const rollupAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName);
        const owner = config.submitterOwner || config.contractAdmin;
        const minimumStake = config.submitterMinimumStake;
        const challengeDeposit = config.submitterChallengeDeposit;
        const rewardPercentage = config.submitterRewardPercentage;
        const plannedSubmitters: string[] = config.batchSubmitterAddresses || [];

        if (
            !ethers.utils.isAddress(proxyAddress) ||
            !ethers.utils.isAddress(implementationAddress) ||
            !ethers.utils.isAddress(rollupAddress) ||
            !ethers.utils.isAddress(owner) ||
            !(minimumStake > 0) ||
            !(challengeDeposit > 0) ||
            !(rewardPercentage > 0 && rewardPercentage <= 100) ||
            !Array.isArray(plannedSubmitters) ||
            plannedSubmitters.some((address) => !ethers.utils.isAddress(address))
        ) {
            return "invalid Submitter install configuration";
        }

        const proxy = await hre.ethers.getContractAt(
            ContractFactoryName.DefaultProxyInterface,
            proxyAddress,
            deployer
        );
        const currentImplementation = await proxy.implementation();
        if (currentImplementation.toLowerCase() !== implementationAddress.toLowerCase()) {
            if (currentImplementation.toLowerCase() !== emptyImplementation.toLowerCase()) {
                return `Submitter proxy has unexpected implementation ${currentImplementation}`;
            }
            const factory = await hre.ethers.getContractFactory(ContractFactoryName.Submitter);
            const initializer = factory.interface.encodeFunctionData("initialize", [
                owner,
                rollupAddress,
                hre.ethers.utils.parseEther(minimumStake.toString()),
                hre.ethers.utils.parseEther(challengeDeposit.toString()),
                rewardPercentage,
            ]);
            const response = await proxy.upgradeToAndCall(implementationAddress, initializer);
            await response.wait();
        }

        // A transparent proxy admin cannot call business getters. Use a non-admin eth_call sender.
        const reader = hre.ethers.Wallet.createRandom().connect(hre.ethers.provider);
        const submitter = await hre.ethers.getContractAt(ContractFactoryName.Submitter, proxyAddress, reader);
        if ((await submitter.owner()).toLowerCase() !== owner.toLowerCase()) {
            return "Submitter owner verification failed";
        }
        if ((await submitter.rollupContract()).toLowerCase() !== rollupAddress.toLowerCase()) {
            return "Submitter rollup binding verification failed";
        }
        if (!(await submitter.minimumStake()).eq(hre.ethers.utils.parseEther(minimumStake.toString()))) {
            return "Submitter minimum stake verification failed";
        }
        if (!(await submitter.challengeDeposit()).eq(hre.ethers.utils.parseEther(challengeDeposit.toString()))) {
            return "Submitter challenge deposit verification failed";
        }
        if (!(await submitter.rewardPercentage()).eq(rewardPercentage)) {
            return "Submitter reward percentage verification failed";
        }
        for (const account of plannedSubmitters) {
            if ((await submitter.registered(account)) || !(await submitter.stakeOf(account)).isZero()) {
                return `Submitter ${account} must remain unregistered and unstaked before Rollup cutover`;
            }
        }

        const rawImplementation = await hre.ethers.provider.getStorageAt(proxyAddress, IMPLEMENTATION_SLOT);
        const implementationInSlot = ethers.utils.getAddress(ethers.utils.hexDataSlice(rawImplementation, 12));
        if (implementationInSlot.toLowerCase() !== implementationAddress.toLowerCase()) {
            return "Submitter EIP-1967 implementation slot verification failed";
        }
        const implementationCode = await hre.ethers.provider.getCode(implementationAddress);
        if (implementationCode === "0x") return "Submitter implementation has no code";
        const blockNumber = await hre.ethers.provider.getBlockNumber();
        return storage(path, ImplStorageName.SubmitterInstallStorageName, proxyAddress.toLowerCase(), blockNumber, {
            implementationSlot: IMPLEMENTATION_SLOT,
            implementation: implementationAddress.toLowerCase(),
            implementationCodeHash: ethers.utils.keccak256(implementationCode),
            rollupContract: rollupAddress.toLowerCase(),
            owner: owner.toLowerCase(),
        });
    } catch (error) {
        console.error("Submitter install failed", error);
        return error instanceof Error ? error.message : "Submitter install failed";
    }
};

export default SubmitterInstall;
