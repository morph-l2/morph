import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition } from "../src/deploy-utils";
import { ethers } from 'ethers'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const StakingInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)

    // Staking config
    const L1StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StakingProxyStorageName)
    const L1StakingImplAddress = getContractAddressByName(path, ImplStorageName.L1StakingStorageName)
    const L1StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)

    const IL1StakingProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1StakingProxyAddress, deployer)
    if (
        (await IL1StakingProxy.implementation()).toLocaleLowerCase() !== L1StakingImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Staking proxy...')
        const admin: string = configTmp.contractAdmin
        const stakingChallengerRewardPercentage: number = configTmp.stakingChallengerRewardPercentage
        const limit: number = configTmp.stakingMinDeposit
        const challengeDeposit: number = configTmp.stakingChallengeDeposit
        const lock: number = configTmp.stakingLockNumber
        const gasLimitAdd: number = configTmp.stakingCrossChainGaslimitAdd
        const gasLimitRemove: number = configTmp.stakingCrossChainGaslimitRemove

        if (!ethers.utils.isAddress(admin)
            || lock <= 0
            || limit <= 0
            || challengeDeposit <= 0
            || gasLimitAdd <= 0
            || gasLimitRemove <= 0
            || stakingChallengerRewardPercentage > 100
            || stakingChallengerRewardPercentage <= 0
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await IL1StakingProxy.upgradeToAndCall(
            L1StakingImplAddress,
            L1StakingFactory.interface.encodeFunctionData('initialize', [
                RollupProxyAddress,
                hre.ethers.utils.parseEther(limit.toString()),
                hre.ethers.utils.parseEther(challengeDeposit.toString()),
                lock.toString(),
                stakingChallengerRewardPercentage,
                gasLimitAdd,
                gasLimitRemove,
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await IL1StakingProxy.implementation()).toLocaleLowerCase() === L1StakingImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        const contractTmp = new ethers.Contract(
            L1StakingProxyAddress,
            L1StakingFactory.interface,
            deployer,
        )

        await assertContractVariable(
            contractTmp,
            'rollupContract',
            RollupProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'rewardPercentage',
            stakingChallengerRewardPercentage
        )
        await assertContractVariable(
            contractTmp,
            'stakingValue',
            hre.ethers.utils.parseEther(limit.toString())
        )
        await assertContractVariable(
            contractTmp,
            'withdrawalLockBlocks',
            lock.toString()
        )
        await assertContractVariable(
            contractTmp,
            'gasLimitAddStaker',
            gasLimitAdd
        )
        await assertContractVariable(
            contractTmp,
            'gasLimitRemoveStakers',
            gasLimitRemove
        )
        await assertContractVariable(
            contractTmp,
            'owner',
            await deployer.getAddress(),
        )
        // Wait for the transaction to execute properly.
        console.log('StakingProxy upgrade success')
    }
    return ''
}

export default StakingInit
