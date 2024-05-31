import "@nomiclabs/hardhat-web3"
import "@nomiclabs/hardhat-ethers"
import "@nomiclabs/hardhat-waffle"

import assert from "assert"
import { task } from "hardhat/config"
import { ethers } from "ethers"
import { assertContractVariable, getContractAddressByName, awaitCondition, storage } from "../src/deploy-utils"
import { ImplStorageName, ProxyStorageName, ContractFactoryName } from "../src/types"

task("rollup-deploy-init")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const ProxyFactoryName = ContractFactoryName.DefaultProxy
        const IProxyFactoryName = ContractFactoryName.DefaultProxyInterface
        const RollupFactoryName = ContractFactoryName.Rollup

        const RollupImplStorageName = ImplStorageName.RollupStorageName
        const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const chainId = 53077

        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)
        const deployer = hre.ethers.provider.getSigner()

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const rollupNewImpl = await RollupFactory.deploy(chainId)
        await rollupNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        let err = await storage(
            newPath,
            RollupImplStorageName,
            rollupNewImpl.address.toLocaleLowerCase(),
            blockNumber || 0
        )
        if (err != "") {
            return err
        }
        console.log(`Rollup new impl deploy at ${rollupNewImpl.address}`)

        // Proxy deploy
        const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
        // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
        const proxy = await ProxyFactory.deploy(EmptyContractAddr, await deployer.getAddress(), "0x")
        await proxy.deployed()
        blockNumber = await hre.ethers.provider.getBlockNumber()
        err = await storage(newPath, RollupProxyStorageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
        if (err != "") {
            return err
        }
        console.log(`Rollup proxy deploy at ${proxy.address}`)

        // ------------------ deploy MultipleVersionRollupVerifier -----------------
        {
            const ZkEvmVerifierV1Address = getContractAddressByName(
                storagePath,
                ImplStorageName.ZkEvmVerifierV1StorageName
            )
            const MultipleVersionRollupVerifierFactoryName = ContractFactoryName.MultipleVersionRollupVerifier
            const MultipleVersionRollupVerifierImplStorageName =
                ImplStorageName.MultipleVersionRollupVerifierStorageName
            console.log("Deploy the MultipleVersionRollupVerifier ...")
            const MultipleVersionRollupVerifierFactory = await hre.ethers.getContractFactory(
                MultipleVersionRollupVerifierFactoryName
            )
            const version = [0]
            const verifiers = [ZkEvmVerifierV1Address]
            const MultipleVersionRollupVerifierContract = await MultipleVersionRollupVerifierFactory.deploy(
                version,
                verifiers
            )
            await MultipleVersionRollupVerifierContract.deployed()
            await MultipleVersionRollupVerifierContract.initialize(proxy.address)
            console.log(
                "%s=%s ; TX_HASH: %s",
                MultipleVersionRollupVerifierImplStorageName,
                MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(),
                MultipleVersionRollupVerifierContract.deployTransaction.hash
            )
            blockNumber = await hre.ethers.provider.getBlockNumber()
            console.log("BLOCK_NUMBER: %s", blockNumber)
            err = await storage(
                newPath,
                MultipleVersionRollupVerifierImplStorageName,
                MultipleVersionRollupVerifierContract.address.toLocaleLowerCase(),
                blockNumber || 0
            )
            if (err != "") {
                return err
            }
        }

        // ------------------ rollup initialize -----------------
        {
            const IRollupProxy = await hre.ethers.getContractAt(IProxyFactoryName, proxy.address, deployer)
            console.log("Upgrading the Rollup proxy...")
            const finalizationPeriodSeconds: number = config.finalizationPeriodSeconds
            const proofWindow: number = config.rollupProofWindow
            const maxNumTxInChunk: number = config.rollupMaxNumTxInChunk

            const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(
                storagePath,
                ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName
            )
            const L1StakingProxyAddress = getContractAddressByName(
                storagePath,
                ProxyStorageName.L1StakingProxyStorageName
            )
            const MultipleVersionRollupVerifierContractAddress = getContractAddressByName(
                newPath,
                ImplStorageName.MultipleVersionRollupVerifierStorageName
            )
            if (
                !ethers.utils.isAddress(L1MessageQueueWithGasPriceOracleProxyAddress) ||
                !ethers.utils.isAddress(MultipleVersionRollupVerifierContractAddress) ||
                !ethers.utils.isAddress(L1StakingProxyAddress)
            ) {
                console.error("please check your address")
                return ""
            }

            // Upgrade and initialize the proxy.
            await IRollupProxy.upgradeToAndCall(
                rollupNewImpl.address,
                RollupFactory.interface.encodeFunctionData("initialize", [
                    L1StakingProxyAddress,
                    L1MessageQueueWithGasPriceOracleProxyAddress,
                    MultipleVersionRollupVerifierContractAddress,
                    maxNumTxInChunk,
                    finalizationPeriodSeconds,
                    proofWindow,
                ])
            )

            await awaitCondition(
                async () => {
                    return (
                        (await IRollupProxy.implementation()).toLocaleLowerCase() ===
                        rollupNewImpl.address.toLocaleLowerCase()
                    )
                },
                3000,
                1000
            )
            // params check
            const contractTmp = new ethers.Contract(proxy.address, RollupFactory.interface, deployer)
            await assertContractVariable(contractTmp, "l1StakingContract", L1StakingProxyAddress)
            await assertContractVariable(contractTmp, "messageQueue", L1MessageQueueWithGasPriceOracleProxyAddress)
            await assertContractVariable(contractTmp, "verifier", MultipleVersionRollupVerifierContractAddress)
            await assertContractVariable(contractTmp, "maxNumTxInChunk", maxNumTxInChunk)
            await assertContractVariable(contractTmp, "finalizationPeriodSeconds", finalizationPeriodSeconds)
            await assertContractVariable(contractTmp, "proofWindow", proofWindow)
            await assertContractVariable(contractTmp, "owner", await deployer.getAddress())

            // Wait for the transaction to execute properly.
            console.log(`RollupProxy upgrade success, Rollup address at ${proxy.address}`)
        }

        // ------------------ Admin transfer -----------------
        {
            const deployerAddr = (await deployer.getAddress()).toLocaleLowerCase()
            const ProxyAdminImplAddr = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
            const IProxyContract = await hre.ethers.getContractAt(
                ContractFactoryName.DefaultProxyInterface,
                proxy.address,
                deployer
            )
            {
                const implAddr = (await IProxyContract.implementation()).toLocaleLowerCase()
                const admin = (await IProxyContract.admin()).toLocaleLowerCase()
                if (implAddr === EmptyContractAddr.toLocaleLowerCase()) {
                    return `Proxy implementation address ${implAddr} should not be empty contract address ${EmptyContractAddr}`
                }
                if (admin !== deployerAddr) {
                    return `Proxy admin address ${admin} should deployer address ${deployerAddr}`
                }
            }
            console.log(`change rollup admin transfer from ${deployerAddr} to ProxyAdmin ${ProxyAdminImplAddr}`)
            const res = await IProxyContract.changeAdmin(ProxyAdminImplAddr)
            await res.wait()
            await assertContractVariable(
                IProxyContract,
                "admin",
                ProxyAdminImplAddr,
                ProxyAdminImplAddr // caller
            )
            console.log(`admin transfer successful`)
        }

        // ------------------ rollup init -----------------
        {
            const Rollup = await hre.ethers.getContractAt(ContractFactoryName.Rollup, proxy.address, deployer)
            // import genesis batch
            const genesisStateRoot: string = config.rollupGenesisStateRoot
            const batchHeader: string = config.batchHeader
            const batchIndex: string = config.rollupBatchIndex

            // submitter and challenger
            const submitter: string = config.rollupProposer
            const challenger: string = config.rollupChallenger
            if (!ethers.utils.isAddress(submitter) || !ethers.utils.isAddress(challenger)) {
                console.error("please check your address")
                return ""
            }
            const res = await Rollup.importGenesisBatch(batchIndex,batchHeader, genesisStateRoot)
            const rec = await res.wait()
            console.log(
                `importGenesisBatch(%s, %s) ${rec.status == 1 ? "success" : "failed"}`,
                batchHeader,
                genesisStateRoot
            )
            res = await Rollup.addChallenger(challenger)
            rec = await res.wait()
            console.log(`addChallenger(%s) ${rec.status == 1 ? "success" : "failed"}`, challenger)
            const batch = await Rollup.batchBaseStore(0)
            assert(batch.batchHash.toLowerCase() != "", `[FATAL] import genesis batch failed`)
        }
    })