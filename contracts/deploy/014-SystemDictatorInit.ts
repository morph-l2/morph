import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, getContractAddressByName, awaitCondition, storge } from "../src/deploy-utils";
import { ethers, BigNumber } from 'ethers'
import assert from 'assert'
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"
import { predeploys } from '../src/constants'
import { expect } from "chai";

export const SystemDictatorInit = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    configTmp: any
): Promise<string> => {
    // Load the contracts we need to interact with.
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const RollupAddress = getContractAddressByName(path, ImplStorageName.RollupStorageName)
    const RollupProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const RollupFactory = await hre.ethers.getContractFactory(ContractFactoryName.Rollup)

    const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStroageName)
    const StakingAddress = getContractAddressByName(path, ImplStorageName.StakingStorageName)
    const StakingProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.Staking)

    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStroageName)
    const L1SequencerAddress = getContractAddressByName(path, ImplStorageName.L1SequencerStorageName)
    const L1SequencerProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const L1SequencerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Sequencer)

    const ZkEvmVerifierV1Address = getContractAddressByName(path, ImplStorageName.ZkEvmVerifierV1StorageName)
    const MorphPortalProxyAddress = getContractAddressByName(path, ProxyStorageName.MorphPortalProxyStroageName)
    // const MultipleVersionRollupVerifierProxyAddress = getContractAddressByName(path, ProxyStorageName.MultipleVersionRollupVerifierStorageName)
    // const MultipleVersionRollupVerifierAddress = getContractAddressByName(path, ImplStorageName.MultipleVersionRollupVerifierStorageName)
    // const MultipleVersionRollupVerifierProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    // const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(ContractFactoryName.MultipleVersionRollupVerifier)

    const SystemDictatorProxyAddress = getContractAddressByName(path, ProxyStorageName.SystemDictatorProxyStorageName)
    const SystemDictatorAddress = getContractAddressByName(path, ImplStorageName.SystemDictatorStorageName)
    const SystemDictatorFactory = await hre.ethers.getContractFactory(ContractFactoryName.SystemDictator)
    const SystemConfigProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)

    const Rollup = new ethers.Contract(
        RollupAddress,
        RollupFactory.interface,
        deployer,
    )
    const RollupProxy = new ethers.Contract(
        RollupProxyAddress,
        RollupProxyFactory.interface,
        deployer.provider,
    )

    const RollupWithSigner = new ethers.Contract(
        RollupProxyAddress,
        RollupFactory.interface,
        deployer,
    )

    const RollupProxyWithSigner = new ethers.Contract(
        RollupProxyAddress,
        RollupProxyFactory.interface,
        deployer,
    )

    const Staking = new ethers.Contract(
        StakingAddress,
        StakingFactory.interface,
        deployer.provider,
    )
    const StakingWithSigner = new ethers.Contract(
        StakingProxyAddress,
        StakingFactory.interface,
        deployer,
    )
    const StakingProxy = new ethers.Contract(
        StakingProxyAddress,
        StakingProxyFactory.interface,
        deployer.provider,
    )
    const StakingProxyWithSigner = new ethers.Contract(
        StakingProxyAddress,
        StakingProxyFactory.interface,
        deployer,
    )

    const L1Sequencer = new ethers.Contract(
        L1SequencerAddress,
        L1SequencerFactory.interface,
        deployer.provider,
    )
    const L1SequencerProxy = new ethers.Contract(
        L1SequencerProxyAddress,
        L1SequencerProxyFactory.interface,
        deployer.provider,
    )
    const L1SequencerProxyWithSigner = new ethers.Contract(
        L1SequencerProxyAddress,
        L1SequencerProxyFactory.interface,
        deployer,
    )

    const SystemDictator = new ethers.Contract(
        SystemDictatorProxyAddress,
        SystemDictatorFactory.interface,
        deployer,
    )
    const SystemDictatorProxy = new ethers.Contract(
        SystemDictatorProxyAddress,
        SystemConfigProxyFactory.interface,
        deployer.provider,
    )
    const SystemDictatorProxyWithSigner = new ethers.Contract(
        SystemDictatorProxyAddress,
        SystemConfigProxyFactory.interface,
        deployer
    )
    const SystemDictatorImpl = new ethers.Contract(
        SystemDictatorAddress,
        SystemDictatorFactory.interface,
        deployer
    )

    // Load the dictator configuration.
    const config = {
        globalConfig: {
            proxyAdmin: getContractAddressByName(path, ImplStorageName.ProxyAdmin),
            controller: configTmp.controller,
            finalOwner: configTmp.finalSystemOwner,
            addressManager: getContractAddressByName(path, ImplStorageName.AddressManager),
        },
        proxyAddressConfig: {
            MorphPortalProxy: getContractAddressByName(
                path,
                ProxyStorageName.MorphPortalProxyStroageName
            ),
            l1CrossDomainMessengerProxy: getContractAddressByName(
                path,
                ProxyStorageName.L1CrossDomainMessengerProxyStroageName
            ),
            l1StandardBridgeProxy: getContractAddressByName(
                path,
                ProxyStorageName.L1StandardBridgeProxyStroageName
            ),
            MorphMintableERC20FactoryProxy: getContractAddressByName(
                path,
                ProxyStorageName.MorphMintableERC20FactoryProxyStroageName
            ),
            l1ERC721BridgeProxy: getContractAddressByName(
                path,
                ProxyStorageName.L1ERC721BridgeProxyStroageName
            ),
            systemConfigProxy: getContractAddressByName(path, ProxyStorageName.SystemConfigProxyStorageName),
        },
        implementationAddressConfig: {
            MorphPortalImpl: getContractAddressByName(path, ImplStorageName.MorphPortalStroageName),
            l1CrossDomainMessengerImpl: getContractAddressByName(
                path,
                ImplStorageName.L1CrossDomainMessengerStorageName
            ),
            l1StandardBridgeImpl: getContractAddressByName(path, ImplStorageName.L1StandardBridgeStroageName),
            MorphMintableERC20FactoryImpl: getContractAddressByName(
                path,
                ImplStorageName.MorphMintableERC20FactoryStroageName
            ),
            l1ERC721BridgeImpl: getContractAddressByName(path, ImplStorageName.L1ERC721BridgeStroageName),
            systemConfigImpl: getContractAddressByName(path, ImplStorageName.SystemConfigStorageName),
        },
        systemConfigConfig: {
            owner: configTmp.finalSystemOwner,
            overhead: configTmp.gasPriceOracleOverhead,
            scalar: configTmp.gasPriceOracleScalar,
            batcherHash: hre.ethers.utils.hexZeroPad(
                configTmp.batchSenderAddress,
                32
            ),
            gasLimit: configTmp.l2GenesisBlockGasLimit,
            unsafeBlockSigner: configTmp.p2pSequencerAddress,
            // The resource config is not exposed to the end user
            // to simplify deploy config. It may be introduced in the future.
            resourceConfig: {
                maxResourceLimit: 20_000_000,
                elasticityMultiplier: 10,
                baseFeeMaxChangeDenominator: 8,
                minimumBaseFee: ethers.utils.parseUnits('1', 'gwei'),
                systemTxMaxGas: 1_000_000,
                maximumBaseFee: BigNumber.from(
                    '0xffffffffffffffffffffffffffffffff'
                ).toString(),
            },
        },
    }


    // deploy and initialize MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
    const MultipleVersionRollupVerifierImplStorageName = ImplStorageName.MultipleVersionRollupVerifierStorageName
    console.log('Deploy the MultipleVersionRollupVerifier ...')
    const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(MultipleVersionRollupVerifierFactoryName)
    const MultipleVersionRollupVerifierContract = await MultipleVersionRollupVerifierFactory.deploy()
    await MultipleVersionRollupVerifierContract.deployed()
    await MultipleVersionRollupVerifierContract.initialize(ZkEvmVerifierV1Address, RollupProxyAddress)
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
        })).toLocaleLowerCase() !== Rollup.address.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Rollup proxy...')
        const finalizationPeriodSeconds: number = configTmp.finalizationPeriodSeconds
        const submitter: string = configTmp.rollupProposer
        const challenger: string = configTmp.rollupChallenger
        const minDeposit: number = configTmp.rollupMinDeposit
        const proofWindow: number = configTmp.rollupProofWindow
        const genesisStateRoot: string = configTmp.rollupGenesisStateRoot
        const withdrawRoot: string = configTmp.withdrawRoot
        const batchHeader: string = configTmp.batchHeader
        const maxNumTxInChunk: number = 100

        if (!ethers.utils.isAddress(submitter)
            || !ethers.utils.isAddress(challenger)
            || !ethers.utils.isAddress(MorphPortalProxyAddress)
            || !ethers.utils.isAddress(MultipleVersionRollupVerifierContract.address)
        ) {
            console.error('please check your address')
            return ''
        }

        // Upgrade and initialize the proxy.
        await RollupProxyWithSigner.upgradeToAndCall(
            Rollup.address,
            Rollup.interface.encodeFunctionData('initialize', [
                MorphPortalProxyAddress,
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
                    })).toLocaleLowerCase() === Rollup.address.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )
        console.log('importGenesisBatch(%s, %s, %s)', batchHeader, genesisStateRoot, withdrawRoot)
        await RollupWithSigner.importGenesisBatch(batchHeader, genesisStateRoot, withdrawRoot)
        console.log('addSequencer(%s)', submitter)
        await RollupWithSigner.addSequencer(submitter)
        console.log('addProver(%s)', submitter)
        await RollupWithSigner.addProver(submitter)
        console.log('addChallenger(%s)', challenger)
        await RollupWithSigner.addChallenger(challenger)

        // Wait for the transaction to execute properly.
        console.log('RollupProxy upgrade success')
    }

    // upgrade and initialize StakingProxy
    if (
        (await StakingProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== Staking.address.toLocaleLowerCase()
    ) {
        console.log('Upgrading the Staking proxy...')
        const stakingSequencerSize: number = configTmp.stakingSequencerSize
        const stakingLockNumber: number = configTmp.stakingLockNumber
        const stakingMinDeposit: number = configTmp.stakingMinDeposit

        // Upgrade and initialize the proxy.
        await StakingProxyWithSigner.upgradeToAndCall(
            Staking.address,
            Staking.interface.encodeFunctionData('initialize', [
                await deployer.getAddress(),
                L1SequencerProxyAddress,
                stakingSequencerSize,
                ethers.utils.parseEther(stakingMinDeposit.toString()),
                stakingLockNumber
            ])
        )
        // Wait for the transaction to execute properly.
        await awaitCondition(
            async () => {
                return (
                    (await StakingProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === Staking.address.toLocaleLowerCase()
                )
            },
            3000,
            1000
        )

        const whiteListAdd = configTmp.l2SequencerAddresses

        const contractTmp = new ethers.Contract(
            StakingProxyAddress,
            StakingFactory.interface,
            deployer,
        )
        // set sequencer to white list
        await contractTmp.updateWhitelist(whiteListAdd, [])
        for (let i = 0; i < configTmp.l2SequencerAddresses.length; i++) {
            // Wait for the transaction to execute properly.
            await awaitCondition(
                async () => {
                    return (
                        await contractTmp.whitelist(configTmp.l2SequencerAddresses[i]) === true
                    )
                },
                3000,
                1000
            )
            console.log(`address ${configTmp.l2SequencerAddresses[i]} is in white list`)
        }

        await assertContractVariable(
            contractTmp,
            'sequencerContract',
            L1SequencerProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'sequencersSize',
            stakingSequencerSize
        )
        await assertContractVariable(
            contractTmp,
            'limit',
            ethers.utils.parseEther(stakingMinDeposit.toString()),
        )
        await assertContractVariable(
            contractTmp,
            'lock',
            stakingLockNumber
        )
        console.log('StakingProxy upgrade success')
    }

    // upgrade and initialize L1SequencerProxy
    if (
        (await L1SequencerProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== L1Sequencer.address.toLocaleLowerCase()
    ) {
        console.log('Upgrading the L1Sequencer proxy...')
        // Upgrade and initialize the proxy.
        await L1SequencerProxyWithSigner.upgradeToAndCall(
            L1Sequencer.address,
            L1Sequencer.interface.encodeFunctionData('initialize', [
                StakingProxyAddress,
                RollupProxyAddress
            ])
        )
        // Wait for the transaction to execute properly.
        await awaitCondition(
            async () => {
                if ((await L1SequencerProxy.callStatic.implementation({
                    from: ethers.constants.AddressZero,
                })).toLocaleLowerCase() !== L1Sequencer.address.toLocaleLowerCase()) {
                    return false
                }
                return true
            },
            3000,
            1000
        )
        const contractTmp = new ethers.Contract(
            L1SequencerProxyAddress,
            L1SequencerFactory.interface,
            deployer,
        )
        await assertContractVariable(
            contractTmp,
            'OTHER_SEQUENCER',
            predeploys.L2Sequencer
        )
        await assertContractVariable(
            contractTmp,
            'stakingContract',
            StakingProxyAddress
        )
        await assertContractVariable(
            contractTmp,
            'rollupContract',
            RollupProxyAddress
        )
        console.log('L1SequencerProxy upgrade success')
    }

    // Update the implementation if necessary.
    if (
        (await SystemDictatorProxy.callStatic.implementation({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== SystemDictatorImpl.address.toLocaleLowerCase()
    ) {
        console.log('Upgrading the SystemDictator proxy...')

        // Upgrade and initialize the proxy.
        await SystemDictatorProxyWithSigner.upgradeToAndCall(
            SystemDictatorImpl.address,
            SystemDictatorImpl.interface.encodeFunctionData('initialize', [config])
        )

        // Wait for the transaction to execute properly.
        await awaitCondition(
            async () => {
                return (
                    (await SystemDictatorProxy.callStatic.implementation({
                        from: ethers.constants.AddressZero,
                    })).toLocaleLowerCase() === SystemDictatorImpl.address
                )
            },
            3000,
            1000
        )
        console.log('SystemDictatorProxy upgrade success')
        // Verify that the contract was initialized correctly.
        const dictatorConfig = await SystemDictator.config()
        for (const [outerConfigKey, outerConfigValue] of Object.entries(config)) {
            for (const [innerConfigKey, innerConfigValue] of Object.entries(
                outerConfigValue
            )) {
                let have = dictatorConfig[outerConfigKey][innerConfigKey]
                let want = innerConfigValue as any

                if (ethers.utils.isAddress(want)) {
                    want = want.toLowerCase()
                    have = have.toLowerCase()
                } else if (typeof want === 'number') {
                    want = ethers.BigNumber.from(want)
                    have = ethers.BigNumber.from(have)
                    assert(
                        want.eq(have),
                        `incorrect config for ${outerConfigKey}.${innerConfigKey}. Want: ${want}, have: ${have}`
                    )
                    return ''
                }

                assert(
                    want === have,
                    `incorrect config for ${outerConfigKey}.${innerConfigKey}. Want: ${want}, have: ${have}`
                )
            }
        }
    }
    console.log('Profile verification successful')

    // Update the owner if necessary.
    if (
        (await SystemDictatorProxy.callStatic.admin({
            from: ethers.constants.AddressZero,
        })).toLocaleLowerCase() !== configTmp.controller.toLocaleLowerCase()
    ) {
        console.log('Transferring ownership of the SystemDictator proxy...')

        // Transfer ownership to the controller address.
        await SystemDictatorProxyWithSigner.changeAdmin(configTmp.controller)

        // Wait for the transaction to execute properly.
        await awaitCondition(
            async () => {
                return (
                    (await SystemDictatorProxy.callStatic.admin({
                        from: ethers.constants.AddressZero,
                    })).toLowerCase() === configTmp.controller
                )
            },
            3000,
            1000
        )
    }
    return ''
}

export default SystemDictatorInit
