import { deployRecordedContract, getDeploymentProxy } from "../src/deployment-state";
import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition, storage } from "../src/deploy-utils";
import { ethers } from 'ethers'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "../src/types"

export const RollupInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const ZkEvmVerifierV1Address = getContractAddressByName(path, ImplStorageName.ZkEvmVerifierV1StorageName)
    const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(path, ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName)
    const SubmitterProxyAddress = getContractAddressByName(path, ProxyStorageName.SubmitterProxyStorageName)

    // Rollup config
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const RollupImplAddress = getContractAddressByName(path, ImplStorageName.RollupStorageName)
    const RollupFactory = await hre.ethers.getContractFactory(ContractFactoryName.Rollup)

    // deploy and initialize MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierImplStorageName = ImplStorageName.MultipleVersionRollupVerifierStorageName
    console.log('Deploy the MultipleVersionRollupVerifier ...')
    const MultipleVersionRollupVerifierContract = await deployRecordedContract(
        hre, path, deployer, MultipleVersionRollupVerifierImplStorageName,
        MultipleVersionRollupVerifierFactoryName, [[1], [ZkEvmVerifierV1Address]]
    );
    const configuredRollup = await MultipleVersionRollupVerifierContract.rollup();
    if (configuredRollup === ethers.constants.AddressZero) {
        await (await MultipleVersionRollupVerifierContract.initialize(RollupProxyAddress)).wait();
    } else if (configuredRollup.toLowerCase() !== RollupProxyAddress.toLowerCase()) {
        throw new Error("MultipleVersionRollupVerifier.rollup does not match deployment records");
    }

    const IRollupProxy = await getDeploymentProxy(hre, path, RollupProxyAddress, deployer)
    // upgrade and initialize RollupProxy
    if (
        (await IRollupProxy.implementation()).toLocaleLowerCase() !== RollupImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Rollup proxy...')
        const finalizationPeriodSeconds: number = config.finalizationPeriodSeconds
        const proofWindow: number = config.rollupProofWindow
        const proofRewardPercent: number = config.proofRewardPercent

        if (!ethers.utils.isAddress(L1MessageQueueWithGasPriceOracleProxyAddress)
            || !ethers.utils.isAddress(MultipleVersionRollupVerifierContract.address)
            || !ethers.utils.isAddress(SubmitterProxyAddress)

        ) {
            throw new Error('please check your address')
        }
        // Upgrade and initialize the proxy.
        await IRollupProxy.upgradeToAndCall(
            RollupImplAddress,
            RollupFactory.interface.encodeFunctionData('initialize', [
                SubmitterProxyAddress,
                L1MessageQueueWithGasPriceOracleProxyAddress,
                MultipleVersionRollupVerifierContract.address,
                finalizationPeriodSeconds,
                proofWindow,
                proofRewardPercent
            ])
        )

        await awaitCondition(
            async () => {
                return (
                    (await IRollupProxy.implementation()).toLocaleLowerCase() === RollupImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        // params check
        const contractTmp = new ethers.Contract(
            RollupProxyAddress,
            RollupFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'submitterContract',
            SubmitterProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'messageQueue',
            L1MessageQueueWithGasPriceOracleProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'verifier',
            MultipleVersionRollupVerifierContract.address,
        )
        await assertContractVariable(
            contractTmp,
            'finalizationPeriodSeconds',
            finalizationPeriodSeconds,
        )
        await assertContractVariable(
            contractTmp,
            'proofWindow',
            proofWindow,
        )
        await assertContractVariable(
            contractTmp,
            'owner',
            await deployer.getAddress(),
        )

        // Wait for the transaction to execute properly.
        console.log('RollupProxy upgrade success')
    }

    return ''
}

export default RollupInit
