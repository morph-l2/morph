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
    const L1StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StakingProxyStorageName)

    // Rollup config
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const RollupImplAddress = getContractAddressByName(path, ImplStorageName.RollupStorageName)
    const RollupFactory = await hre.ethers.getContractFactory(ContractFactoryName.Rollup)

    // deploy and initialize MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierImplStorageName = ImplStorageName.MultipleVersionRollupVerifierStorageName
    console.log('Deploy the MultipleVersionRollupVerifier ...')
    const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(MultipleVersionRollupVerifierFactoryName)
    const version = [0]
    const verifiers = [ZkEvmVerifierV1Address]
    const MultipleVersionRollupVerifierContract = await MultipleVersionRollupVerifierFactory.deploy(version, verifiers)
    await MultipleVersionRollupVerifierContract.deployed()
    await MultipleVersionRollupVerifierContract.initialize(RollupProxyAddress)
    console.log("%s=%s ; TX_HASH: %s", MultipleVersionRollupVerifierImplStorageName, MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(), MultipleVersionRollupVerifierContract.deployTransaction.hash);
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storage(path, MultipleVersionRollupVerifierImplStorageName, MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    const IRollupProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, RollupProxyAddress, deployer)
    // upgrade and initialize RollupProxy
    if (
        (await IRollupProxy.implementation()).toLocaleLowerCase() !== RollupImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Rollup proxy...')
        const finalizationPeriodSeconds: number = config.finalizationPeriodSeconds
        const proofWindow: number = config.rollupProofWindow
        const maxNumTxInChunk: number = config.rollupMaxNumTxInChunk

        if (!ethers.utils.isAddress(L1MessageQueueWithGasPriceOracleProxyAddress)
            || !ethers.utils.isAddress(MultipleVersionRollupVerifierContract.address)
            || !ethers.utils.isAddress(L1StakingProxyAddress)

        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await IRollupProxy.upgradeToAndCall(
            RollupImplAddress,
            RollupFactory.interface.encodeFunctionData('initialize', [
                L1StakingProxyAddress,
                L1MessageQueueWithGasPriceOracleProxyAddress,
                MultipleVersionRollupVerifierContract.address,
                maxNumTxInChunk,
                finalizationPeriodSeconds,
                proofWindow
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
            'l1StakingContract',
            L1StakingProxyAddress
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
            'maxNumTxInChunk',
            maxNumTxInChunk,
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
