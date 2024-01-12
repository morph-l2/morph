
import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { ethers } from 'ethers'
import {
    HardhatRuntimeEnvironment
} from 'hardhat/types';
import { assertContractVariable, storge, awaitCondition } from "../src/deploy-utils";
import {
    ImplStorageName,
    ProxyStorageName,
    ContractFactoryName
} from "./types"
import assert from 'assert'

export const deploySystemConfig = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log('\n---------------------------------- deploy  SystemConfig ----------------------------------')
    const proxyStorageName = ProxyStorageName.SystemConfigProxyStorageName
    const implStorageName = ImplStorageName.SystemConfigStorageName
    // deploy proxy
    const ProxyFactoy = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
    const proxy = await ProxyFactoy.deploy(await deployer.getAddress())
    console.log("%s=%s ; TX_HASH: %s", proxyStorageName, proxy.address.toLocaleLowerCase(), proxy.deployTransaction.hash);
    await assertContractVariable(proxy, 'admin', await deployer.getAddress())
    let blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    let err = await storge(path, proxyStorageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // SystemConfig deploy
    const batcherHash = hre.ethers.utils
        .hexZeroPad(config.batchSenderAddress, 32)
        .toLowerCase()
    const uint128Max = ethers.BigNumber.from('0xffffffffffffffffffffffffffffffff')
    const Factory = await hre.ethers.getContractFactory(ContractFactoryName.SystemConfig)
    const contract = await Factory.deploy(
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
    // check params then storge
    console.log("%s=%s ; TX_HASH: %s", implStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
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
    blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, implStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    // proxy upgradeAndCall
    console.log('Upgrading the SystemConfig proxy...')
    await proxy.upgradeToAndCall(
        contract.address,
        contract.interface.encodeFunctionData('initialize', [
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
        ])
    )
    // Wait for the transaction to execute properly.
    await awaitCondition(
        async () => {
            const temp = new ethers.Contract(
                proxy.address,
                proxy.interface,
                proxy.provider
            )
            const actual = await temp.callStatic['implementation']({
                from: ethers.constants.AddressZero,
            })
            return (
                actual.toLocaleLowerCase() === contract.address.toLocaleLowerCase()
            )
        },
        30000,
        1000
    )
    // check params
    const checkContract = new ethers.Contract(
        proxy.address,
        contract.interface,
        proxy.provider
    )
    await assertContractVariable(
        checkContract,
        'owner',
        config.finalSystemOwner,
    )
    await assertContractVariable(
        checkContract,
        'overhead',
        config.gasPriceOracleOverhead
    )
    await assertContractVariable(
        checkContract,
        'scalar',
        config.gasPriceOracleScalar
    )
    await assertContractVariable(checkContract, 'batcherHash', batcherHash)
    await assertContractVariable(
        checkContract,
        'unsafeBlockSigner',
        config.p2pSequencerAddress
    )

    configs = await checkContract.resourceConfig()
    assert(configs.maxResourceLimit === 20_000_000)
    assert(configs.elasticityMultiplier === 10)
    assert(configs.baseFeeMaxChangeDenominator === 8)
    assert(configs.systemTxMaxGas === 1_000_000)
    assert(ethers.utils.parseUnits('1', 'gwei').eq(configs.minimumBaseFee))
    assert(configs.maximumBaseFee.eq(uint128Max))
    console.log('Upgrading the SystemConfig proxy Success...')

    return ''
}

export default deploySystemConfig

