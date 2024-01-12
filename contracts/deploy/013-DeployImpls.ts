import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, getContractAddressByName } from "../src/deploy-utils";
import { predeploys } from '../src/constants'
import { ethers } from 'ethers'
import assert from 'assert'
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName,
} from "./types"

export const deployContractImpls = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    const L1CrossDomainMessengerFactoryName = ContractFactoryName.L1CrossDomainMessenger
    const L1StandardBridgeFactoryName = ContractFactoryName.L1StandardBridge
    const MorphPortalFactoryName = ContractFactoryName.MorphPortal
    const MorphMintableERC20FactoryFactoryName = ContractFactoryName.MorphMintableERC20Factory
    const L1ERC721BridgeFactoryName = ContractFactoryName.L1ERC721Bridge
    const SystemConfigFactoryName = ContractFactoryName.SystemConfig
    const SystemDictatorFactoryName = ContractFactoryName.SystemDictator
    const RollupFactoryName = ContractFactoryName.Rollup
    const StakingFactoryName = ContractFactoryName.Staking
    const L1SequencerFactoryName = ContractFactoryName.L1Sequencer
    const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier

    const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
    const L1StandardBridgeImplStorageName = ImplStorageName.L1StandardBridgeStroageName
    const MorphPortalImplStorageName = ImplStorageName.MorphPortalStroageName
    const MorphMintableERC20FactoryImplStorageName = ImplStorageName.MorphMintableERC20FactoryStroageName
    const L1ERC721BridgeImplStorageName = ImplStorageName.L1ERC721BridgeStroageName
    const SystemConfigImplStorageName = ImplStorageName.SystemConfigStorageName
    const SystemDictatorImplStorageName = ImplStorageName.SystemDictatorStorageName
    const RollupImplStorageName = ImplStorageName.RollupStorageName
    const StakingImplStorageName = ImplStorageName.StakingStorageName
    const L1SequencerImplStorageName = ImplStorageName.L1SequencerStorageName
    const MultipleVersionRollupVerifierImplStorageName = ImplStorageName.MultipleVersionRollupVerifierStorageName

    const MorphPortalProxyAddress = getContractAddressByName(path, ProxyStorageName.MorphPortalProxyStroageName)
    const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
    const Artifact__SystemConfigProxyAddress = getContractAddressByName(path, ProxyStorageName.SystemConfigProxyStorageName)
    const L1StandardBridgeProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StandardBridgeProxyStroageName)
    const RollupProxyAddress = getContractAddressByName(path, ProxyStorageName.RollupProxyStorageName)
    const StakingProxyAddress = getContractAddressByName(path, ProxyStorageName.StakingProxyStroageName)
    const L1SequencerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1SequencerProxyStroageName)
    const MultipleVersionRollupVerifierProxyAddress = getContractAddressByName(path, ProxyStorageName.MultipleVersionRollupVerifierStorageName)

    // L1CrossDomainMessenger deploy
    let Factory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
    let contract = await Factory.deploy(MorphPortalProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'PORTAL',
        MorphPortalProxyAddress
    )
    let err = await storge(path, L1CrossDomainMessengerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.blockNumber || 0)
    if (err != '') {
        return err
    }

    // Rollup deploy
    const l2ChainID: string = config.l2ChainID
    Factory = await hre.ethers.getContractFactory(RollupFactoryName)
    contract = await Factory.deploy(l2ChainID, L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", RollupImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash)
    // check params
    await assertContractVariable(contract, 'layer2ChainId', l2ChainID)
    let blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, RollupImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // Staking deploy 
    Factory = await hre.ethers.getContractFactory(StakingFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", StakingImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, StakingImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1Sequencer deploy 
    Factory = await hre.ethers.getContractFactory(L1SequencerFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    await assertContractVariable(
        contract,
        'OTHER_SEQUENCER',
        predeploys.L2Sequencer
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1SequencerImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1StandardBridge deploy 
    Factory = await hre.ethers.getContractFactory(L1StandardBridgeFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1StandardBridgeImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    await assertContractVariable(
        contract,
        'OTHER_BRIDGE',
        predeploys.L2StandardBridge
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1StandardBridgeImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // MorphPortal deploy 
    const portalGuardian = config.portalGuardian
    const portalGuardianCode = await hre.ethers.provider.getCode(portalGuardian)
    if (portalGuardianCode === '0x') {
        console.log(
            `WARNING: setting MorphPortal.GUARDIAN to ${portalGuardian} and it has no code`
        )
    }
    Factory = await hre.ethers.getContractFactory(MorphPortalFactoryName)
    contract = await Factory.deploy(
        portalGuardian,
        true, // paused
        Artifact__SystemConfigProxyAddress,
        RollupProxyAddress,
        L1CrossDomainMessengerProxyAddress
    )
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", MorphPortalImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'GUARDIAN',
        portalGuardian
    )
    await assertContractVariable(
        contract,
        'SYSTEM_CONFIG',
        Artifact__SystemConfigProxyAddress
    )
    await assertContractVariable(
        contract,
        'l1Messenger',
        L1CrossDomainMessengerProxyAddress
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, MorphPortalImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // MorphMintableERC20Factory deploy
    Factory = await hre.ethers.getContractFactory(MorphMintableERC20FactoryFactoryName)
    contract = await Factory.deploy(L1StandardBridgeProxyAddress)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", MorphMintableERC20FactoryImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'BRIDGE',
        L1StandardBridgeProxyAddress
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, MorphMintableERC20FactoryImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // L1ERC721Bridge deploy
    Factory = await hre.ethers.getContractFactory(L1ERC721BridgeFactoryName)
    contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress, predeploys.L2ERC721Bridge)
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", L1ERC721BridgeImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'MESSENGER',
        L1CrossDomainMessengerProxyAddress
    )
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, L1ERC721BridgeImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // SystemConfig deploy
    const batcherHash = hre.ethers.utils
        .hexZeroPad(config.batchSenderAddress, 32)
        .toLowerCase()
    const uint128Max = ethers.BigNumber.from('0xffffffffffffffffffffffffffffffff')
    Factory = await hre.ethers.getContractFactory(SystemConfigFactoryName)
    contract = await Factory.deploy(
        config.finalSystemOwner,
        config.gasPriceOracleOverhead,
        config.gasPriceOracleScalar,
        batcherHash,
        config.l2GenesisBlockGasLimit,
        config.p2pSequencerAddress,
        {
            maxResourceLimit: 20_000_000,
            elasticityMultiplier: 10,
            baseFeeMaxChangeDenominator: 8,
            systemTxMaxGas: 1_000_000,
            minimumBaseFee: ethers.utils.parseUnits('1', 'gwei'),
            maximumBaseFee: uint128Max,
        },
    )
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", SystemConfigImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    await assertContractVariable(
        contract,
        'owner',
        config.finalSystemOwner,
    )
    await assertContractVariable(
        contract,
        'overhead',
        config.gasPriceOracleOverhead
    )
    await assertContractVariable(
        contract,
        'scalar',
        config.gasPriceOracleScalar
    )
    await assertContractVariable(contract, 'batcherHash', batcherHash)
    await assertContractVariable(
        contract,
        'unsafeBlockSigner',
        config.p2pSequencerAddress
    )

    let configs = await contract.resourceConfig()
    assert(configs.maxResourceLimit === 20_000_000)
    assert(configs.elasticityMultiplier === 10)
    assert(configs.baseFeeMaxChangeDenominator === 8)
    assert(configs.systemTxMaxGas === 1_000_000)
    assert(ethers.utils.parseUnits('1', 'gwei').eq(configs.minimumBaseFee))
    assert(configs.maximumBaseFee.eq(uint128Max))
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, SystemConfigImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // SystemDictator deploy
    Factory = await hre.ethers.getContractFactory(SystemDictatorFactoryName)
    contract = await Factory.deploy()
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", SystemDictatorImplStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params
    blockNumber = await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, SystemDictatorImplStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }
    return ''
}

export default deployContractImpls
