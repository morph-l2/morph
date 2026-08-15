import "@nomiclabs/hardhat-ethers";

import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";

import { assertContractVariable, awaitCondition, getContractAddressByName } from "../src/deploy-utils";
import { ContractFactoryName, ImplStorageName, ProxyStorageName } from "../src/types";

export const SubmitterInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    const proxyAddress = getContractAddressByName(path, ProxyStorageName.SubmitterProxyStorageName);
    const implementationAddress = getContractAddressByName(path, ImplStorageName.SubmitterStorageName);
    const rollupAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName);
    const owner: string = config.submitterOwner;
    const minimumStake: number = config.submitterMinimumStake;
    const challengeDeposit: number = config.submitterChallengeDeposit;
    const rewardPercentage: number = config.submitterRewardPercentage;

    if (
        !ethers.utils.isAddress(proxyAddress) ||
        !ethers.utils.isAddress(implementationAddress) ||
        !ethers.utils.isAddress(rollupAddress) ||
        !ethers.utils.isAddress(owner) ||
        minimumStake <= 0 ||
        challengeDeposit <= 0 ||
        rewardPercentage <= 0 ||
        rewardPercentage > 100
    ) {
        return "invalid Submitter configuration";
    }

    const proxy = await hre.ethers.getContractAt(
        ContractFactoryName.DefaultProxyInterface,
        proxyAddress,
        deployer
    );
    const factory = await hre.ethers.getContractFactory(ContractFactoryName.Submitter);
    if ((await proxy.implementation()).toLowerCase() !== implementationAddress.toLowerCase()) {
        const initializer = factory.interface.encodeFunctionData("initialize", [
            owner,
            rollupAddress,
            hre.ethers.utils.parseEther(minimumStake.toString()),
            hre.ethers.utils.parseEther(challengeDeposit.toString()),
            rewardPercentage,
        ]);
        await (await proxy.upgradeToAndCall(implementationAddress, initializer)).wait();
    }

    await awaitCondition(
        async () => (await proxy.implementation()).toLowerCase() === implementationAddress.toLowerCase(),
        3000,
        1000
    );

    const submitter = new ethers.Contract(proxyAddress, factory.interface, deployer);
    await assertContractVariable(submitter, "owner", owner);
    await assertContractVariable(submitter, "rollupContract", rollupAddress);
    await assertContractVariable(
        submitter,
        "minimumStake",
        hre.ethers.utils.parseEther(minimumStake.toString())
    );
    await assertContractVariable(
        submitter,
        "challengeDeposit",
        hre.ethers.utils.parseEther(challengeDeposit.toString())
    );
    await assertContractVariable(submitter, "rewardPercentage", rewardPercentage);

    console.log(`Submitter proxy initialized at ${proxyAddress}`);
    return "";
};

export default SubmitterInit;
