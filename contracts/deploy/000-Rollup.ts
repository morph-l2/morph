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

export const deployRollup = async (
    hre: HardhatRuntimeEnvironment,
    path: string,
    deployer: any,
    config: any
): Promise<string> => {
    console.log('\n---------------------------------- deploy  Rollup ----------------------------------')
    const proxyStorageName = ProxyStorageName.RollupProxyStorageName
    const implStorageName = ImplStorageName.RollupStorageName
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
    // get init params
    const rollupGenesisBlockNumber: number = config.rollupGenesisBlockNumber
    const finalizationPeriodSeconds: number = config.finalizationPeriodSeconds
    const submitter: string = config.rollupProposer
    const challenger: string = config.rollupChallenger
    const minDeposit: number = config.rollupMinDeposit
    const proofWindow: number = config.rollupProofWindow
    const genesisStateRoot: string = config.rollupGenesisStateRoot
    if (!ethers.utils.isAddress(submitter) || !ethers.utils.isAddress(challenger)) {
        console.error('please check your address')
        return ''
    }
    // rollup deploy
    const Factory = await hre.ethers.getContractFactory(ContractFactoryName.Rollup)
    const contract = await Factory.deploy(
        submitter,
        challenger,
        ethers.utils.parseEther(minDeposit.toString()), // min deposit
        finalizationPeriodSeconds,
        proofWindow, //_proofWindow
        rollupGenesisBlockNumber,
        genesisStateRoot
    )
    await contract.deployed()
    console.log("%s=%s ; TX_HASH: %s", implStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
    // check params then storge
    await assertContractVariable(contract, 'owner', await deployer.getAddress())
    blockNumber =  await hre.ethers.provider.getBlockNumber()
    console.log("BLOCK_NUMBER: %s", blockNumber)
    err = await storge(path, implStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
    if (err != '') {
        return err
    }

    console.log('Upgrading the Rollup proxy...')
    // Upgrade and initialize the proxy.
    await proxy.upgradeToAndCall(
        contract.address,
        contract.interface.encodeFunctionData('initialize', [
            submitter,
            challenger,
            ethers.utils.parseEther(minDeposit.toString()), // min deposit
            finalizationPeriodSeconds,
            proofWindow, //_proofWindow
            rollupGenesisBlockNumber,
            genesisStateRoot
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
    await assertContractVariable(checkContract, 'owner', await deployer.getAddress())
    await assertContractVariable(checkContract, 'submitter', submitter)
    await assertContractVariable(checkContract, 'challenger', challenger)
    await assertContractVariable(checkContract, 'startingBlockNumber', rollupGenesisBlockNumber)
    await assertContractVariable(checkContract, 'FINALIZATION_PERIOD_SECONDS', finalizationPeriodSeconds)
    await assertContractVariable(checkContract, 'MIN_DEPOSIT', ethers.utils.parseEther(minDeposit.toString()))
    await assertContractVariable(checkContract, 'PROOF_WINDOW', proofWindow)
    console.log('RollupProxy upgrade success')
    return ''
}

export default deployRollup