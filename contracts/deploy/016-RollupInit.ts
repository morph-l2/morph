import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition, storge } from "../src/deploy-utils";
import { ethers } from 'ethers'

import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const RollupInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const ProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const ZkEvmVerifierV1Address = getContractAddressByName(path, ImplStorageName.ZkEvmVerifierV1StorageName)
    const L1MessageQueueProxyAddress = getContractAddressByName(path, ProxyStorageName.L1MessageQueueProxyStroageName)

    // Rollup config
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const RollupImplAddress = getContractAddressByName(path, ImplStorageName.RollupStorageName)
    const RollupFactory = await hre.ethers.getContractFactory(ContractFactoryName.Rollup)

    const RollupProxy = new ethers.Contract(
        RollupProxyAddress,
        ProxyFactory.interface,
        deployer.provider,
    )
    const Rollup = new ethers.Contract(
        RollupProxyAddress,
        RollupFactory.interface,
        deployer.provider,
    )
    // deploy and initialize MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierImplStorageName = ImplStorageName.MultipleVersionRollupVerifierStorageName
    console.log('Deploy the MultipleVersionRollupVerifier ...')
    const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(MultipleVersionRollupVerifierFactoryName)
    const MultipleVersionRollupVerifierContract = await MultipleVersionRollupVerifierFactory.deploy(ZkEvmVerifierV1Address)
    await MultipleVersionRollupVerifierContract.deployed()
    await MultipleVersionRollupVerifierContract.initialize(RollupProxyAddress)
    console.log("%s=%s ; TX_HASH: %s", MultipleVersionRollupVerifierImplStorageName, MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(), MultipleVersionRollupVerifierContract.deployTransaction.hash);
    const blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, MultipleVersionRollupVerifierImplStorageName, MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // upgrade and initialize RollupProxy
    if (
        (await RollupProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== RollupImplAddress.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Rollup proxy...')
        const finalizationPeriodSeconds: number = configTmp.finalizationPeriodSeconds
        const minDeposit: number = configTmp.rollupMinDeposit
        const proofWindow: number = configTmp.rollupProofWindow
        const maxNumTxInChunk: number = configTmp.rollupMaxNumTxInChunk

        // submitter and challenger
        const submitter: string = configTmp.rollupProposer
        const challenger: string = configTmp.rollupChallenger

        // import genesis batch 
        const genesisStateRoot: string = configTmp.rollupGenesisStateRoot
        const withdrawRoot: string = configTmp.withdrawRoot
        const batchHeader: string = configTmp.batchHeader

        if (!ethers.utils.isAddress(submitter)
            || !ethers.utils.isAddress(challenger)
            || !ethers.utils.isAddress(L1MessageQueueProxyAddress)
            || !ethers.utils.isAddress(MultipleVersionRollupVerifierContract.address)
        ) {
            console.error('please check your address')
            return ''
        }
        // Upgrade and initialize the proxy.
        await RollupProxy.connect(deployer).upgradeToAndCall(
            RollupImplAddress,
            RollupFactory.interface.encodeFunctionData('initialize', [
                L1MessageQueueProxyAddress,
                MultipleVersionRollupVerifierContract.address,
                maxNumTxInChunk,
                ethers.utils.parseEther(minDeposit.toString()),
                finalizationPeriodSeconds,
                proofWindow
            ])
        )
        await awaitCondition(
            async () => {
                return (
                    (await RollupProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === RollupImplAddress.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        console.log('importGenesisBatch(%s, %s, %s)', batchHeader, genesisStateRoot, withdrawRoot)
        await Rollup.connect(deployer).importGenesisBatch(batchHeader, genesisStateRoot, withdrawRoot)
        console.log('addSequencer(%s)', submitter)
        await Rollup.connect(deployer).addSequencer(submitter)
        console.log('addProver(%s)', submitter)
        await Rollup.connect(deployer).addProver(submitter)
        console.log('addChallenger(%s)', challenger)
        await Rollup.connect(deployer).addChallenger(challenger)

        // params check
        const contractTmp = new ethers.Contract(
            RollupProxyAddress,
            RollupFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'messageQueue',
            L1MessageQueueProxyAddress
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
            'FINALIZATION_PERIOD_SECONDS',
            finalizationPeriodSeconds,
        )
        await assertContractVariable(
            contractTmp,
            'PROOF_WINDOW',
            proofWindow,
        )
        await assertContractVariable(
            contractTmp,
            'MIN_DEPOSIT',
            ethers.utils.parseEther(minDeposit.toString()),
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
