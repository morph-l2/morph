import "@nomiclabs/hardhat-web3"
import "@nomiclabs/hardhat-ethers"
import "@nomiclabs/hardhat-waffle"

import assert from "assert"
import { task } from "hardhat/config"
import { ethers } from "ethers"
import { assertContractVariable, getContractAddressByName, awaitCondition, storage } from "../src/deploy-utils"
import { ImplStorageName, ProxyStorageName, ContractFactoryName } from "../src/types"
import { predeploys } from "../src"

task("rollup-deploy-init")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const chainId = config.l2ChainID
        const deployer = hre.ethers.provider.getSigner()

        const ProxyFactoryName = ContractFactoryName.DefaultProxy
        const IProxyFactoryName = ContractFactoryName.DefaultProxyInterface
        const RollupFactoryName = ContractFactoryName.Rollup
        const RollupImplStorageName = ImplStorageName.RollupStorageName
        const RollupProxyStorageName = ProxyStorageName.RollupProxyStorageName
        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)

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

        // ------------------ rollup import genesis batch -----------------
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
            let res = await Rollup.importGenesisBatch(batchIndex, batchHeader, genesisStateRoot)
            let rec = await res.wait()
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

task("l1mq-upgrade")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath

        const L1MessageQueueWithGasPriceOracleFactoryName = ContractFactoryName.L1MessageQueueWithGasPriceOracle
        const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName
        )
        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.L1CrossDomainMessengerProxyStorageName
        )
        const EnforcedTxGatewayAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.EnforcedTxGatewayProxyStorageName
        )
        const ProxyAdminAddress = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)

        // deploy L1MessageQueueWithGasPriceOracle impl
        {
            const NewRollupProxyAddress = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

            const Factory = await hre.ethers.getContractFactory(L1MessageQueueWithGasPriceOracleFactoryName)
            const contract = await Factory.deploy(
                L1CrossDomainMessengerProxyAddress,
                NewRollupProxyAddress,
                EnforcedTxGatewayAddress
            )
            await contract.deployed()
            console.log(
                "%s=%s ; TX_HASH: %s",
                ImplStorageName.L1MessageQueueWithGasPriceOracle,
                contract.address.toLocaleLowerCase(),
                contract.deployTransaction.hash
            )
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            console.log("BLOCK_NUMBER: %s", blockNumber)
            const err = await storage(
                newPath,
                ImplStorageName.L1MessageQueueWithGasPriceOracle,
                contract.address.toLocaleLowerCase(),
                blockNumber || 0
            )
            if (err != "") {
                return err
            }
        }

        // L1MessageQueueWithGasPriceOracle proxy upgrade
        {
            const L1MessageQueueWithGasPriceOracleNewImplAddress = getContractAddressByName(
                newPath,
                ImplStorageName.L1MessageQueueWithGasPriceOracle
            )

            const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
            const proxyAdmin = ProxyAdminFactory.attach(ProxyAdminAddress)
            const res = await proxyAdmin.upgrade(
                L1MessageQueueWithGasPriceOracleProxyAddress,
                L1MessageQueueWithGasPriceOracleNewImplAddress
            )
            const rec = await res.wait()
            console.log(`upgrade l1MessageQueueWithGasPriceOracle.rollup ${rec.status === 1}`)
        }
    })

task("l1cdm-upgrade")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath

        const L1CrossDomainMessengerFactoryName = ContractFactoryName.L1CrossDomainMessenger
        const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName
        const ProxyAdminAddress = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.L1CrossDomainMessengerProxyStorageName
        )

        // deploy l1CrossDomainMessenger impl
        {
            const Factory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
            const contract = await Factory.deploy()
            await contract.deployed()
            console.log(
                "%s=%s ; TX_HASH: %s",
                L1CrossDomainMessengerImplStorageName,
                contract.address.toLocaleLowerCase(),
                contract.deployTransaction.hash
            )
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            console.log("BLOCK_NUMBER: %s", blockNumber)
            const err = await storage(
                newPath,
                L1CrossDomainMessengerImplStorageName,
                contract.address.toLocaleLowerCase(),
                blockNumber || 0
            )
            if (err != "") {
                return err
            }
        }

        // l1CrossDomainMessenger proxy upgrade
        {
            const L1CrossDomainMessengerNewImplAddress = getContractAddressByName(
                newPath,
                L1CrossDomainMessengerImplStorageName
            )
            const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
            const proxyAdmin = ProxyAdminFactory.attach(ProxyAdminAddress)
            const res = await proxyAdmin.upgrade(
                L1CrossDomainMessengerProxyAddress,
                L1CrossDomainMessengerNewImplAddress
            )
            const rec = await res.wait()
            console.log(`upgrade l1CrossDomainMessenger ${rec.status === 1}`)
        }

        // l1CrossDomainMessenger update Rollup address
        {
            const NewRollupProxyAddress = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

            const L1CDMFactory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
            const l1CrossDomainMessenger = L1CDMFactory.attach(L1CrossDomainMessengerProxyAddress)
            const res = await l1CrossDomainMessenger.updateRollup(NewRollupProxyAddress)
            const rec = await res.wait()
            console.log(`update l1CrossDomainMessenger.rollup ${rec.status === 1}`)
        }
    })


task("l1staking-deploy-init")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const deployer = hre.ethers.provider.getSigner()

        const ProxyFactoryName = ContractFactoryName.DefaultProxy
        const L1StakingProxyStorageName = ProxyStorageName.L1StakingProxyStorageName
        const WhitelistImplAddress = getContractAddressByName(storagePath, ImplStorageName.Whitelist)
        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)
        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(storagePath, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)

        // deploy L1Staking proxy
        {
            const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
            // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
            const proxy = await ProxyFactory.deploy(EmptyContractAddr, await deployer.getAddress(), "0x")
            await proxy.deployed()
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            const err = await storage(newPath, L1StakingProxyStorageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
            if (err != "") {
                console.log(`deploy L1Staking proxy failed ${err}`)
                return err
            }
            console.log(`L1Staking proxy deploy at ${proxy.address}`)
        }

        // deploy impl
        {
            const Factory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)
            const contract = await Factory.deploy(L1CrossDomainMessengerProxyAddress)
            await contract.deployed()
            console.log("%s=%s ; TX_HASH: %s", ImplStorageName.L1StakingStorageName, contract.address.toLocaleLowerCase(), contract.deployTransaction.hash);
            await assertContractVariable(
                contract,
                'MESSENGER',
                L1CrossDomainMessengerProxyAddress
            )
            await assertContractVariable(
                contract,
                'OTHER_STAKING',
                predeploys.L2Staking.toLowerCase()
            )
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            console.log("BLOCK_NUMBER: %s", blockNumber)
            const err = await storage(newPath, ImplStorageName.L1StakingStorageName, contract.address.toLocaleLowerCase(), blockNumber || 0)
            if (err != '') {
                console.log(`deploy L1Staking implemention failed ${err}`)
                return err
            }
        }

        // upgrade
        {
            const RollupProxyAddress = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)
            // Staking config
            const L1StakingProxyAddress = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)
            const L1StakingImplAddress = getContractAddressByName(newPath, ImplStorageName.L1StakingStorageName)
            const L1StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)
            const IL1StakingProxy = await hre.ethers.getContractAt(ContractFactoryName.DefaultProxyInterface, L1StakingProxyAddress, deployer)

            console.log('Upgrading the Staking proxy...')
            const admin: string = config.contractAdmin
            const stakingChallengerRewardPercentage: number = config.stakingChallengerRewardPercentage
            const limit: number = config.stakingMinDeposit
            const lock: number = config.stakingLockNumber
            const gasLimitAdd: number = config.stakingCrossChainGaslimitAdd
            const gasLimitRemove: number = config.stakingCrossChainGaslimitRemove

            if (!ethers.utils.isAddress(admin)
                || lock <= 0
                || limit <= 0
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
                    hre.ethers.utils.parseEther(lock.toString()),
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
                hre.ethers.utils.parseEther(lock.toString())
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

        // ------------------ Admin transfer -----------------
        {
            const L1StakingProxyAddress = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)
            const deployerAddr = (await deployer.getAddress()).toLocaleLowerCase()
            const ProxyAdminImplAddr = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
            const IProxyContract = await hre.ethers.getContractAt(
                ContractFactoryName.DefaultProxyInterface,
                L1StakingProxyAddress,
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
            console.log(`change L1Staking admin transfer from ${deployerAddr} to ProxyAdmin ${ProxyAdminImplAddr}`)
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

        // set L1Staking address to gasPriceOracle whitelist
        {
            const L1StakingProxyAddress = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)
            const WhitelistCheckerImpl = await hre.ethers.getContractAt(ContractFactoryName.Whitelist, WhitelistImplAddress, deployer)
            let addList = [L1StakingProxyAddress]
            const res = await WhitelistCheckerImpl.updateWhitelistStatus(addList, true)
            await res.wait()
            for (let i = 0; i < addList.length; i++) {
                let res = await WhitelistCheckerImpl.isSenderAllowed(addList[i])
                if (res != true) {
                    console.error('whitelist check failed')
                    return ''
                }
            }
            console.log(`add ${addList} to whitelist success`)
        }

        // set staker whitelist
        {
            const L1StakingProxyAddress = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)
            const L1Staking = await hre.ethers.getContractAt(ContractFactoryName.L1Staking, L1StakingProxyAddress, deployer)
            const whiteListAdd = config.l2SequencerAddresses
            // set sequencer to white list
            await L1Staking.updateWhitelist(whiteListAdd, [])
            for (let i = 0; i < config.l2SequencerAddresses.length; i++) {
                // Wait for the transaction to execute properly.
                await awaitCondition(
                    async () => {
                        return (
                            await L1Staking.whitelist(config.l2SequencerAddresses[i]) === true
                        )
                    },
                    3000,
                    1000
                )
                console.log(`address ${config.l2SequencerAddresses[i]} is in white list`)
            }
        }
    })

task("check-params")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath

        const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName
        )
        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(
            storagePath,
            ProxyStorageName.L1CrossDomainMessengerProxyStorageName
        )
        const RollupNewProxyAddress = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

        const L1CDMFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1CrossDomainMessenger)
        const l1CrossDomainMessenger = L1CDMFactory.attach(L1CrossDomainMessengerProxyAddress)
        let res = await l1CrossDomainMessenger.rollup()
        assert(
            RollupNewProxyAddress.toLowerCase() === res.toLowerCase(),
            `l1CrossDomainMessenger assert rollup address assert failed, expect ${RollupNewProxyAddress}, actual ${res}`
        )

        const L1MQFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1MessageQueueWithGasPriceOracle)
        const l1mq = L1MQFactory.attach(L1MessageQueueWithGasPriceOracleProxyAddress)
        res = await l1mq.ROLLUP_CONTRACT()
        assert(
            RollupNewProxyAddress.toLowerCase() === res.toLowerCase(),
            `l1mq assert rollup address assert failed, expect ${RollupNewProxyAddress}, actual ${res}`
        )
        console.log("Check new rollup address success")
    })

// test command
// rm -rf ./deployFile.json && \                                                                            
// yarn hardhat deploy --storagepath ./deployFile.json --network l1 && \
// yarn hardhat initialize  --storagepath ./deployFile.json --network l1 && \
// yarn hardhat fund --network l1 && \
// yarn hardhat register --storagepath ./deployFile.json --network l1 && \
// rm -rf ./newFile.json && \
// yarn hardhat rollup-deploy-init --storagepath ./deployFile.json --newpath ./newFile.json --network l1 && \
// yarn hardhat l1mq-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network l1 && \
// yarn hardhat l1cdm-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network l1 && \
// yarn hardhat check-params --storagepath ./deployFile.json --newpath ./newFile.json --network l1 && \
// yarn hardhat l1staking-deploy-init --storagepath ./deployFile.json --newpath ./newFile.json --network l1 && \
// yarn hardhat register --storagepath ./newFile.json --network l1