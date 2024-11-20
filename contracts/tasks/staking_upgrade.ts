import "@nomiclabs/hardhat-web3"
import "@nomiclabs/hardhat-ethers"
import "@nomiclabs/hardhat-waffle"

import fs from "fs";
import assert from "assert"
import { task } from "hardhat/config"
import { ethers } from "ethers"
import { assertContractVariable, getContractAddressByName, awaitCondition, storage } from "../src/deploy-utils"
import { ImplStorageName, ProxyStorageName, ContractFactoryName } from "../src/types"
import { predeploys } from "../src"
import { hexlify } from "ethers/lib/utils";

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

// immutable upgrade test
task("impl-test")
    .setAction(async (taskArgs, hre) => {
        const deployer = hre.ethers.provider.getSigner()
        const V1Factory = await hre.ethers.getContractFactory("TestUpgradeV1")
        const V2Factory = await hre.ethers.getContractFactory("TestUpgradeV2")

        const v1Impl = await V1Factory.deploy()
        await v1Impl.deployed()
        const v2Impl = await V2Factory.deploy()
        await v2Impl.deployed()
        console.log(`V1 and V2 impl deploy success and v1: ${v1Impl.address} , v2: ${v2Impl.address}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = await ProxyAdminFactory.deploy()
        await proxyAdmin.deployed()
        const ProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
        const proxy = await ProxyFactory.deploy(v1Impl.address, await deployer.getAddress(), "0x")
        await proxy.deployed()
        const IProxyContract = await hre.ethers.getContractAt(
            ContractFactoryName.DefaultProxyInterface,
            proxy.address,
            deployer
        )

        // transfer owner to proxy admin
        {
            const res = await IProxyContract.changeAdmin(proxyAdmin.address)
            const rec = await res.wait()
            console.log(`proxy admin change to proxy admin ${rec.status === 1}`)
        }

        const consoleParams = async (factory) => {
            const contractTmp = new ethers.Contract(
                proxy.address,
                factory.interface,
                hre.ethers.provider,
            )
            let va = await contractTmp.va({ from: hre.ethers.constants.AddressZero })
            let vb = await contractTmp.vb({ from: hre.ethers.constants.AddressZero })
            let vc = await contractTmp.vc({ from: hre.ethers.constants.AddressZero })
            let version = await contractTmp.version({ from: hre.ethers.constants.AddressZero })
            console.log(`va ${va} ; vb ${vb} ; vc ${vc} ; version ${version}`)
        }
        let contract = new ethers.Contract(
            proxy.address,
            V1Factory.interface,
            deployer,
        )
        let res = await contract.setVersion(100)
        let rec = await res.wait()
        console.log(`update version to 100 ${rec.status === 1}`)
        await consoleParams(V1Factory)

        // upgrade
        {
            const res = await proxyAdmin.upgrade(proxy.address, v2Impl.address)
            const rec = await res.wait()
            console.log(`upgrade to v2 impl ${rec.status === 1}`)
        }
        contract = new ethers.Contract(
            proxy.address,
            V2Factory.interface,
            deployer,
        )
        res = await contract.setVersion(101)
        rec = await res.wait()
        console.log(`update version to 101 ${rec.status === 1}`)
        res = await contract.setOtherVersion(99)
        rec = await res.wait()
        console.log(`update otherVersion to 99 ${rec.status === 1}`)
        console.log("upgrade success")
        await consoleParams(V2Factory)
    })

// ===============================================  deploy  ===============================================

// yarn hardhat l1staking-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat l1staking-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("l1staking-deploy")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const deployer = hre.ethers.provider.getSigner()

        const ProxyFactoryName = ContractFactoryName.DefaultProxy
        const L1StakingProxyStorageName = ProxyStorageName.L1StakingProxyStorageName
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
    })

// yarn hardhat l1staking-init --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat l1staking-init --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("l1staking-init")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const deployer = hre.ethers.provider.getSigner()

        const WhitelistImplAddress = getContractAddressByName(storagePath, ImplStorageName.Whitelist)
        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)

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

// yarn hardhat rollup-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat rollup-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("rollup-deploy")
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
    })

// yarn hardhat rollup-init --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat rollup-init --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("rollup-init")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const deployer = hre.ethers.provider.getSigner()

        const IProxyFactoryName = ContractFactoryName.DefaultProxyInterface
        const RollupFactoryName = ContractFactoryName.Rollup
        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const NewRollupContractImplAddr = getContractAddressByName(newPath, ImplStorageName.RollupStorageName)
        const NewRollupContractAddr = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

        // ------------------ rollup initialize -----------------
        {
            const IRollupProxy = await hre.ethers.getContractAt(IProxyFactoryName, NewRollupContractAddr, deployer)
            console.log("Upgrading the Rollup proxy...")
            const finalizationPeriodSeconds: number = config.finalizationPeriodSeconds
            const proofWindow: number = config.rollupProofWindow
            const maxNumTxInChunk: number = config.rollupMaxNumTxInChunk

            const L1MessageQueueWithGasPriceOracleProxyAddress = getContractAddressByName(
                storagePath,
                ProxyStorageName.L1MessageQueueWithGasPriceOracleProxyStorageName
            )
            const L1StakingProxyAddress = getContractAddressByName(
                newPath,
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
                NewRollupContractImplAddr,
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
                        NewRollupContractImplAddr.toLocaleLowerCase()
                    )
                },
                3000,
                1000
            )
            // params check
            const contractTmp = new ethers.Contract(NewRollupContractAddr, RollupFactory.interface, deployer)
            await assertContractVariable(contractTmp, "l1StakingContract", L1StakingProxyAddress)
            await assertContractVariable(contractTmp, "messageQueue", L1MessageQueueWithGasPriceOracleProxyAddress)
            await assertContractVariable(contractTmp, "verifier", MultipleVersionRollupVerifierContractAddress)
            await assertContractVariable(contractTmp, "maxNumTxInChunk", maxNumTxInChunk)
            await assertContractVariable(contractTmp, "finalizationPeriodSeconds", finalizationPeriodSeconds)
            await assertContractVariable(contractTmp, "proofWindow", proofWindow)
            await assertContractVariable(contractTmp, "owner", await deployer.getAddress())

            // Wait for the transaction to execute properly.
            console.log(`RollupProxy upgrade success, Rollup address at ${NewRollupContractAddr}`)
        }

        // ------------------ Admin transfer -----------------
        {
            const deployerAddr = (await deployer.getAddress()).toLocaleLowerCase()
            const ProxyAdminImplAddr = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
            const IProxyContract = await hre.ethers.getContractAt(
                ContractFactoryName.DefaultProxyInterface,
                NewRollupContractAddr,
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
    })

// yarn hardhat rollup-import-genesis-batch --newpath ./newFile.json  --network qanetl1
// yarn hardhat rollup-import-genesis-batch --newpath ./newFile.json  --network holesky
task("rollup-import-genesis-batch")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const deployer = hre.ethers.provider.getSigner()
        // ------------------ rollup import genesis batch -----------------
        {
            const NewRollupContractAddr = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

            const Rollup = await hre.ethers.getContractAt(ContractFactoryName.Rollup, NewRollupContractAddr, deployer)
            // import genesis batch
            const batchHeader: string = config.batchHeader
            console.log("batchHeader: ", batchHeader)
            
            // submitter and challenger
            const submitter: string = config.rollupProposer
            const challenger: string = config.rollupChallenger
            if (!ethers.utils.isAddress(submitter) || !ethers.utils.isAddress(challenger)) {
                console.error("please check your address")
                return ""
            }
            let res = await Rollup.importGenesisBatch(batchHeader)
            let rec = await res.wait()
            console.log(
                `txHash: %s, importGenesisBatch(%s, %s) ${rec.status == 1 ? "success" : "failed"}`,
                res.hash,
                batchHeader
            )
            res = await Rollup.addChallenger(challenger)
            rec = await res.wait()
            console.log(`addChallenger(%s) ${rec.status == 1 ? "success" : "failed"}`, challenger)
            const batch = await Rollup.batchDataStore(0)

            console.log(batch.blockNumber)
        }
    })

// yarn hardhat l1cerc20gw-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat l1cerc20gw-deploy --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("l1cerc20gw-deploy")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const deployer = hre.ethers.provider.getSigner()

        const ProxyFactoryName = ContractFactoryName.DefaultProxy
        const L1CustomERC20GatewayProxyStorageName = ProxyStorageName.L1CustomERC20GatewayProxyStorageName
        const EmptyContractAddr = getContractAddressByName(storagePath, ImplStorageName.EmptyContract)

        const L1GatewayRouterProxyAddress = getContractAddressByName(storagePath, ProxyStorageName.L1GatewayRouterProxyStorageName)
        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(storagePath, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)

        const IProxyFactoryName = ContractFactoryName.DefaultProxyInterface

        let proxy: ethers.Contract
        let cFactory: ethers.ContractFactory
        let implContract: ethers.Contract
        let IL1CustomERC20GatewayProxy: ethers.Contract

        // deploy L1CustomERC20Gateway proxy
        {
            const ProxyFactory = await hre.ethers.getContractFactory(ProxyFactoryName)
            // TransparentUpgradeableProxy deploy with emptyContract as impl, deployer as admin
            proxy = await ProxyFactory.deploy(EmptyContractAddr, await deployer.getAddress(), "0x")
            await proxy.deployed()
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            const err = await storage(newPath, L1CustomERC20GatewayProxyStorageName, proxy.address.toLocaleLowerCase(), blockNumber || 0)
            if (err != "") {
                console.log(`deploy L1CustomERC20Gateway proxy failed ${err}`)
                return err
            }
            console.log(`L1CustomERC20Gateway proxy deploy at ${proxy.address}`)
        }

        // deploy impl
        {
            cFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1CustomERC20Gateway)
            implContract = await cFactory.deploy()
            await implContract.deployed()

            console.log("%s=%s ; TX_HASH: %s", ImplStorageName.L1CustomERC20GatewayStorageName, implContract.address.toLocaleLowerCase(), implContract.deployTransaction.hash);
            const blockNumber = await hre.ethers.provider.getBlockNumber()
            console.log("BLOCK_NUMBER: %s", blockNumber)
            const err = await storage(newPath, ImplStorageName.L1CustomERC20GatewayStorageName, implContract.address.toLocaleLowerCase(), blockNumber || 0)
            if (err != '') {
                console.log(`deploy L1Staking implemention failed ${err}`)
                return err
            }
        }

        // ------------------ l1CustomERC20Gateway initialize -----------------
        {
            IL1CustomERC20GatewayProxy = await hre.ethers.getContractAt(IProxyFactoryName, proxy.address, deployer)
            // Upgrade and initialize the proxy.
            await IL1CustomERC20GatewayProxy.upgradeToAndCall(
                implContract.address,
                cFactory.interface.encodeFunctionData("initialize", [
                    predeploys.L2CustomERC20Gateway,
                    L1GatewayRouterProxyAddress,
                    L1CrossDomainMessengerProxyAddress
                ])
            )

            // Wait for the transaction to execute properly.
            console.log(`RollupProxy upgrade success, Rollup address at ${proxy.address}`)
        }

        // ------------------ Admin transfer -----------------
        {
            const deployerAddr = (await deployer.getAddress()).toLocaleLowerCase()
            const ProxyAdminImplAddr = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)

            console.log(`change rollup admin transfer from ${deployerAddr} to ProxyAdmin ${ProxyAdminImplAddr}`)
            const res = await IL1CustomERC20GatewayProxy.changeAdmin(ProxyAdminImplAddr)
            await res.wait()
  
            console.log(`admin transfer successful`)
        }


    })

// ===============================================  l2 deploy  ===============================================

// yarn hardhat distribute-deploy --owner 0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E --network qanetl2
// yarn hardhat distribute-deploy --owner 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d --network hl2
task("distribute-deploy")
    .addParam("owner")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const DistributeFactory = await hre.ethers.getContractFactory("Distribute")
        const distribute = await DistributeFactory.deploy()
        await distribute.deployed()
        const res = await proxyAdmin.upgradeAndCall(
            predeploys.Distribute,
            distribute.address,
            DistributeFactory.interface.encodeFunctionData('initialize', [_owner])
        )
        const rec = await res.wait()
        console.log(`distribute upgrade ${rec.status === 1}, new impl ${distribute.address}`)
    })

// yarn hardhat l2-staking-deploy --owner 0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E --newpath ./newFile.json --l2config ../ops/l2-genesis/deploy-config/qanet-deploy-config.json --network qanetl2
// yarn hardhat l2-staking-deploy --owner 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d --newpath ./newFile.json --l2config ../ops/l2-genesis/deploy-config/holesky-deploy-config.json --network hl2
task("l2-staking-deploy")
    .addParam("owner")
    .addParam("newpath")
    .addParam("l2config")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner
        const newPath = taskArgs.newpath
        const data = fs.readFileSync(taskArgs.l2config);
        // @ts-ignore
        const l2Config = JSON.parse(data);
        const L1StakingProxyAddr = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const L2StakingFactory = await hre.ethers.getContractFactory("L2Staking")
        const staking = await L2StakingFactory.deploy(L1StakingProxyAddr)
        await staking.deployed()

        let infos = []
        for (let i = 0; i < l2Config.l2StakingAddresses.length; i++) {
            let info = {
                addr: l2Config.l2StakingAddresses[i],
                tmKey: l2Config.l2StakingTmKeys[i],
                blsKey: l2Config.l2StakingBlsKeys[i],
            }
            infos.push(info)
        }
        const res = await proxyAdmin.upgradeAndCall(
            predeploys.L2Staking,
            staking.address,
            L2StakingFactory.interface.encodeFunctionData('initialize', [
                _owner,
                l2Config.l2StakingSequencerMaxSize,
                l2Config.l2StakingUnDelegatedLockEpochs,
                l2Config.l2StakingRewardStartTime,
                infos,
            ])
        )
        const rec = await res.wait()
        console.log(`L2Staking upgrade ${rec.status === 1}, new impl ${staking.address}`)
    })

// yarn hardhat sequencer-deploy --owner 0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E --l2config ../ops/l2-genesis/deploy-config/qanet-deploy-config.json  --network qanetl2
// yarn hardhat sequencer-deploy --owner 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d --l2config ../ops/l2-genesis/deploy-config/holesky-deploy-config.json  --network hl2
task("sequencer-deploy")
    .addParam("owner")
    .addParam("l2config")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner
        const data = fs.readFileSync(taskArgs.l2config);
        // @ts-ignore
        const l2Config = JSON.parse(data);
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const SequencerFactory = await hre.ethers.getContractFactory("Sequencer")
        const sequencer = await SequencerFactory.deploy()
        await sequencer.deployed()

        // l2Config
        let addresses = []
        for (let i = 0; i < l2Config.l2StakingAddresses.length; i++) {
            addresses.push(l2Config.l2StakingAddresses[i])
        }
        const res = await proxyAdmin.upgradeAndCall(
            predeploys.Sequencer,
            sequencer.address,
            SequencerFactory.interface.encodeFunctionData(
                'initialize',
                [_owner, addresses])
        )
        const rec = await res.wait()
        console.log(`sequencer upgrade ${rec.status === 1}, new impl ${sequencer.address}`)
    })

// yarn hardhat morph-token-deploy --l2config ../ops/l2-genesis/deploy-config/qanet-deploy-config.json  --network qanetl2
// yarn hardhat morph-token-deploy --l2config ../ops/l2-genesis/deploy-config/holesky-deploy-config.json  --network hl2
task("morph-token-deploy")
    .addParam("l2config")
    .setAction(async (taskArgs, hre) => {
        const data = fs.readFileSync(taskArgs.l2config);
        // @ts-ignore
        const l2Config = JSON.parse(data);
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const MorphTokenFactory = await hre.ethers.getContractFactory("MorphToken")
        const morphToken = await MorphTokenFactory.deploy()
        await morphToken.deployed()

        const res = await proxyAdmin.upgradeAndCall(
            predeploys.MorphToken,
            morphToken.address,
            MorphTokenFactory.interface.encodeFunctionData('initialize', [
                l2Config.morphTokenName,
                l2Config.morphTokenSymbol,
                l2Config.morphTokenOwner,
                hre.ethers.utils.parseEther(l2Config.morphTokenInitialSupply.toString()),
                l2Config.morphTokenDailyInflationRate,
            ])
        )
        const rec = await res.wait()
        console.log(`morphToken upgrade ${rec.status === 1}, new impl ${morphToken.address}`)
    })

// ===============================================  deploy  ===============================================



// ===============================================  upgrade  ===============================================

// yarn hardhat l1cdm-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat l1cdm-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
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

task("l1crossdomainmessenger-upgrade-hc")
    .addParam("pa")
    .addParam("l1cdmpa")
    .setAction(async (taskArgs, hre) => {
        console.log(`l1CrossDomainMessenger upgrade hardcode start……`)

        const L1CrossDomainMessengerFactoryName = ContractFactoryName.L1CrossDomainMessenger
        const L1CrossDomainMessengerImplStorageName = ImplStorageName.L1CrossDomainMessengerStorageName

        // deploy l1CrossDomainMessenger impl
        const Factory = await hre.ethers.getContractFactory(L1CrossDomainMessengerFactoryName)
        const newImpl = await Factory.deploy()
        await newImpl.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            L1CrossDomainMessengerImplStorageName,
            newImpl.address.toLocaleLowerCase(),
            newImpl.deployTransaction.hash
        )


        // l1CrossDomainMessenger proxy upgrade
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.pa)
        const res = await proxyAdmin.upgrade(
            taskArgs.l1cdmpa,
            newImpl.address
        )
        const rec = await res.wait()
        console.log(`upgrade l1CrossDomainMessenger ${rec.status === 1}`)

        console.log(`l1CrossDomainMessenger upgrade hardcode finished……`)
    })

// yarn hardhat l1mq-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
// yarn hardhat l1mq-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network holesky
task("l1mq-upgrade")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        console.log(`L1MessageQueueWithGasPriceOracle upgrade start……`)

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

        console.log(`L1MessageQueueWithGasPriceOracle upgrade finished……`)
    })

// yarn hardhat l1mq-upgrade-hc --l1mq 0xa8fefb05f4fe3d06afc9a1de1f3b058d69e675b4 --l1cdm 0x4a811f1ba56a9cfd6a5bd28d50b61d154489aa72 --l1etg 0xe592bc3c98912a6574b5492511b023789b6b0548  --l1pa 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77 --ru 0xc1daf2538d8c190f45046ec7dc6e82c63041d574 --network qanetl1
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0xa8fefb05f4fe3d06afc9a1de1f3b058d69e675b4 "implementation()(address)" --from 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0xa8fefb05f4fe3d06afc9a1de1f3b058d69e675b4 "MESSENGER()(address)"
task("l1mq-upgrade-hc")
    .addParam("l1mq")
    .addParam("l1cdm")
    .addParam("l1etg")
    .addParam("l1pa")
    .addParam("ru")
    // .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        console.log(`L1MessageQueueWithGasPriceOracle upgrade hardcode start……`)
        
        // const newPath = taskArgs.newpath

        const L1MessageQueueWithGasPriceOracleFactoryName = ContractFactoryName.L1MessageQueueWithGasPriceOracle

        // deploy L1MessageQueueWithGasPriceOracle impl

        const Factory = await hre.ethers.getContractFactory(L1MessageQueueWithGasPriceOracleFactoryName)
        const NewImpl = await Factory.deploy(
            taskArgs.l1cdm,
            taskArgs.ru,
            taskArgs.l1etg
        )
        await NewImpl.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            ImplStorageName.L1MessageQueueWithGasPriceOracle,
            NewImpl.address.toLocaleLowerCase(),
            NewImpl.deployTransaction.hash
        )
        const blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)
        // const err = await storage(
        //     newPath,
        //     ImplStorageName.L1MessageQueueWithGasPriceOracle,
        //     contract.address.toLocaleLowerCase(),
        //     blockNumber || 0
        // )
        // if (err != "") {
        //     return err
        // }


        // L1MessageQueueWithGasPriceOracle proxy upgrade
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.l1pa)
        const res = await proxyAdmin.upgrade(
            taskArgs.l1mq,
            NewImpl.address
        )
        const rec = await res.wait()
        console.log(`upgrade l1MessageQueueWithGasPriceOracle.rollup ${rec.status === 1}`)
        
        console.log(`L1MessageQueueWithGasPriceOracle upgrade hardcode finished……`)
    })

// yarn hardhat l1staking-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
task("l1staking-upgrade")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        // deploy config
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath

        const L1CrossDomainMessengerProxyAddress = getContractAddressByName(storagePath, ProxyStorageName.L1CrossDomainMessengerProxyStorageName)
        const NewL1StakingProxyAddress = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)
        console.log("NewL1StakingProxyAddress: ", NewL1StakingProxyAddress)

        // deploy impl
        const l1StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)
        const l1StakingImpl = await l1StakingFactory.deploy(L1CrossDomainMessengerProxyAddress)
        await l1StakingImpl.deployed()
        console.log("%s=%s ; TX_HASH: %s", ImplStorageName.L1StakingStorageName, l1StakingImpl.address.toLocaleLowerCase(), l1StakingImpl.deployTransaction.hash);
        await assertContractVariable(
            l1StakingImpl,
            'MESSENGER',
            L1CrossDomainMessengerProxyAddress
        )
        await assertContractVariable(
            l1StakingImpl,
            'OTHER_STAKING',
            predeploys.L2Staking.toLowerCase()
        )
        const blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const L1ProxyAdminAddress = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
        console.log("L1ProxyAdminAddress: ", L1ProxyAdminAddress)

        const proxyAdmin = ProxyAdminFactory.attach(L1ProxyAdminAddress)

        let res = await proxyAdmin.upgrade(
            NewL1StakingProxyAddress, 
            l1StakingImpl.address)
        let rec = await res.wait()

        console.log(`upgrade L1Staking ${rec.status == 1 ? "success" : "failed"}`)   

    })

// yarn hardhat l1staking-upgrade-hc --l1cdmpa 0x4a811f1ba56a9cfd6a5bd28d50b61d154489aa72 --l1pa 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77 --l1spa 0xf631833debdce1b59e81e05f3aaa442fd79ff421 --network qanetl1
// yarn hardhat l1staking-upgrade-hc --l1cdmpa 0xecc966ab425f3f5bd58085ce4ebdbf81d829126f --l1pa 0x1d0846952983f836d873c7a2e1d0bc5136cc70c7 --l1spa 0xcb4496399ffd2a94c4c0132c5561b420b8c73ad8 --network holesky
task("l1staking-upgrade-hc")
    .addParam("l1cdmpa")
    .addParam("l1pa")
    .addParam("l1spa")
    .setAction(async (taskArgs, hre) => {

        // deploy impl
        const l1StakingFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1Staking)
        const l1StakingImpl = await l1StakingFactory.deploy(taskArgs.l1cdmpa)
        await l1StakingImpl.deployed()
        console.log("%s=%s ; TX_HASH: %s", ImplStorageName.L1StakingStorageName, l1StakingImpl.address.toLocaleLowerCase(), l1StakingImpl.deployTransaction.hash);
        await assertContractVariable(
            l1StakingImpl,
            'MESSENGER',
            taskArgs.l1cdmpa
        )
        await assertContractVariable(
            l1StakingImpl,
            'OTHER_STAKING',
            predeploys.L2Staking.toLowerCase()
        )
        const blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log("BLOCK_NUMBER: %s", blockNumber)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)

        console.log("L1ProxyAdminAddress: ", taskArgs.l1pa)

        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.l1pa)

        let res = await proxyAdmin.upgrade(
            taskArgs.l1spa, 
            l1StakingImpl.address)
        let rec = await res.wait()

        console.log(`upgrade L1Staking ${rec.status == 1 ? "success" : "failed"}`)   

    })

// yarn hardhat rollup-upgrade --storagepath ./deployFile.json --newpath ./newFile.json --network qanetl1
task("rollup-upgrade")
    .addParam("storagepath")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {
        const storagePath = taskArgs.storagepath
        const newPath = taskArgs.newpath
        const config = hre.deployConfig
        const deployer = hre.ethers.provider.getSigner()
        const chainId = config.l2ChainID

        const RollupFactoryName = ContractFactoryName.Rollup

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const rollupNewImpl = await RollupFactory.deploy(chainId)
        await rollupNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`Rollup new impl deploy at ${rollupNewImpl.address} and height ${blockNumber}`)

        const NewRollupContractAddr = getContractAddressByName(newPath, ProxyStorageName.RollupProxyStorageName)

        const ProxyAdminFactoryName = ContractFactoryName.ProxyAdmin
        const ProxyAdminImplAddr = getContractAddressByName(storagePath, ImplStorageName.ProxyAdmin)
        const ProxyAdmin = await hre.ethers.getContractAt(ProxyAdminFactoryName, ProxyAdminImplAddr, deployer)
        let res = await ProxyAdmin.upgrade(NewRollupContractAddr, rollupNewImpl.address)
        let rec = await res.wait()
        console.log(`upgrade rollup ${rec.status == 1 ? "success" : "failed"}`)
    })

// check admin
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0x4356d5c1fc3c56e2d3a1600bfeb14a67b85dd5e5 "admin()(address)" --from 0xbc1f4a5e934b43cb635c422a6dd02e2ac2785f5f
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0xbc1f4a5e934b43cb635c422a6dd02e2ac2785f5f "owner()(address)"
// yarn hardhat rollup-upgrade-hc --l1pa 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77 --l2cid 53077 --rollup 0xc1daf2538d8c190f45046ec7dc6e82c63041d574 --network qanetl1
// yarn hardhat rollup-upgrade-hc --l1pa 0x1d0846952983f836d873c7a2e1d0bc5136cc70c7 --l2cid 2810 --rollup 0x165b77247e71fbf53460ede5aecca4e49fbdf205 --network holesky
task("rollup-upgrade-hc")
    .addParam("l1pa")
    .addParam("l2cid")
    .addParam("rollup")
    .setAction(async (taskArgs, hre) => {
        const ProxyAdminImplAddr = taskArgs.l1pa
        const chainId = taskArgs.l2cid
        const RollupProxyAddr = taskArgs.rollup
        const deployer = hre.ethers.provider.getSigner()

        const RollupFactoryName = ContractFactoryName.Rollup

        const RollupFactory = await hre.ethers.getContractFactory(RollupFactoryName)
        const rollupNewImpl = await RollupFactory.deploy(chainId)
        await rollupNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`Rollup new impl deploy at ${rollupNewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactoryName = ContractFactoryName.ProxyAdmin
        const ProxyAdmin = await hre.ethers.getContractAt(ProxyAdminFactoryName, ProxyAdminImplAddr, deployer)
        let res = await ProxyAdmin.upgrade(RollupProxyAddr, rollupNewImpl.address)
        let rec = await res.wait()
        console.log(`upgrade rollup ${rec.status == 1 ? "success" : "failed"}`)
    })

// l1gatewayrouter
task("l1gatewayrouter-upgrade")
    .addParam("owner")
    .setAction(async (taskArgs, hre) => {
        console.log(`L1GatewayRouter upgrade start……`)

        const owner = taskArgs.owner

        
        console.log(`L1GatewayRouter upgrade finished……`)
    })

// yarn hardhat l1gatewayrouter-upgrade-hc --l1pa 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77 --l1gwr 0xf31809718ba2bff201d6783b1ea3dbf2259dcce5 --network qanetl1
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0xf31809718ba2bff201d6783b1ea3dbf2259dcce5 "implementation()(address)" --from 0xa05ccae77659adbf0c1cb371855db3d8fa56bf77
// cast call -r http://l2-qa-morph-l1-geth.bitkeep.tools 0xf31809718ba2bff201d6783b1ea3dbf2259dcce5 "ethGateway()(address)"
task("l1gatewayrouter-upgrade-hc")
    .addParam("l1pa")
    .addParam("l1gwr")
    .setAction(async (taskArgs, hre) => {
        console.log(`L1GatewayRouter upgrade hardcode start……`)

        const owner = taskArgs.owner
        const NewImplFactory = await hre.ethers.getContractFactory("L1GatewayRouter")
        const NewImpl = await NewImplFactory.deploy()
        await NewImpl.deployed()

        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`GasPriceOracle new impl deploy at ${NewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.l1pa)
        let res = await proxyAdmin.upgrade(
            taskArgs.l1gwr, 
            NewImpl.address)
        let rec = await res.wait()

        console.log(`L1GatewayRouter upgrade hardcode finished……`)
    })

// L1ETHGateway
// todo......

// L1StandardERC20Gateway
// todo......

// L1CustomERC20Gateway
// todo......

// L1WithdrawLockERC20Gateway
// todo......

// L1ReverseCustomGateway
// todo......

task("l1reversecustomgateway-upgrade-hc")
    .addParam("pa")
    .addParam("l1rcgw")
    .setAction(async (taskArgs, hre) => {
        console.log(`L1ReverseCustomGateway upgrade hardcode start……`)

        const L1ReverseCustomGatewayFactoryName = ContractFactoryName.L1ReverseCustomGateway
        const L1ReverseCustomGatewayImplStorageName = ImplStorageName.L1ReverseCustomGatewayStorageName

        // deploy L1ReverseCustomGateway impl
        const Factory = await hre.ethers.getContractFactory(L1ReverseCustomGatewayFactoryName)
        const newImpl = await Factory.deploy()
        await newImpl.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            L1ReverseCustomGatewayImplStorageName,
            newImpl.address.toLocaleLowerCase(),
            newImpl.deployTransaction.hash
        )

        // L1ReverseCustomGateway proxy upgrade
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.pa)
        const res = await proxyAdmin.upgrade(
            taskArgs.l1rcgw,
            newImpl.address
        )
        const rec = await res.wait()
        console.log(`upgrade L1ReverseCustomGateway ${rec.status === 1}`)

        console.log(`L1ReverseCustomGateway upgrade hardcode finished……`)
    })

// L1ERC721Gateway
// todo......

// L1ERC1155Gateway
// todo......

// EnforcedTxGateway
// todo......

task("enforcedtxgateway-upgrade-hc")
    .addParam("pa")
    .addParam("eftgw")
    .setAction(async (taskArgs, hre) => {
        console.log(`EnforcedTxGateway upgrade hardcode start……`)

        const EnforcedTxGatewayFactoryName = ContractFactoryName.EnforcedTxGateway
        const EnforcedTxGatewayImplStorageName = ImplStorageName.EnforcedTxGatewayStorageName

        // deploy EnforcedTxGateway impl
        const Factory = await hre.ethers.getContractFactory(EnforcedTxGatewayFactoryName)
        const newImpl = await Factory.deploy()
        await newImpl.deployed()
        console.log(
            "%s=%s ; TX_HASH: %s",
            EnforcedTxGatewayImplStorageName,
            newImpl.address.toLocaleLowerCase(),
            newImpl.deployTransaction.hash
        )

        // EnforcedTxGateway proxy upgrade
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(taskArgs.pa)
        const res = await proxyAdmin.upgrade(
            taskArgs.eftgw,
            newImpl.address
        )
        const rec = await res.wait()
        console.log(`upgrade EnforcedTxGateway ${rec.status === 1}`)

        console.log(`EnforcedTxGateway upgrade hardcode finished……`)
    })


// L1WETHGateway
// todo......

// L1USDCGateway
// todo......

// MultipleVersionRollupVerifier????
// todo......


// ===============================================  l2 upgrade  ===============================================

// yarn hardhat proxyadmin-upgrade --network l2
// 应该没写完，需要补充完全
task("proxyadmin-upgrade")
    .setAction(async (taskArgs, hre) => {
        const deployer = hre.ethers.provider.getSigner()
        const IProxyFactoryName = ContractFactoryName.DefaultProxyInterface
        const IProxyAdminProxy = await hre.ethers.getContractAt(IProxyFactoryName, predeploys.ProxyAdmin, deployer)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxya = await ProxyAdminFactory.deploy()
        await proxya.deployed()

        const res = await IProxyAdminProxy.upgradeTo(
            proxya.address
        )
        
        const rec = await res.wait()
        console.log(`ProxyAdmin upgrade ${rec.status === 1}, new impl ${proxya.address}`)
    })

// L2ToL1MessagePasser
// todo......

// L2GatewayRouter
// todo......

task("gov-upgrade")
    .addParam("owner")
    .addParam("l2config")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner
        const data = fs.readFileSync(taskArgs.l2config);
        // @ts-ignore
        const l2Config = JSON.parse(data);
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const GovFactory = await hre.ethers.getContractFactory("Gov")
        const newGovImpl = await GovFactory.deploy()
        await newGovImpl.deployed()
        const res = await proxyAdmin.upgradeAndCall(
            predeploys.Gov,
            newGovImpl.address,
            GovFactory.interface.encodeFunctionData('initialize', [
                _owner,
                l2Config.govVotingDuration,
                l2Config.govBatchBlockInterval,
                l2Config.govBatchTimeout,
                l2Config.govRollupEpoch,
            ])
        )
        const rec = await res.wait()
        console.log(`gov upgrade ${rec.status === 1}`)
    })

// L2ETHGateway
// todo......

// L2CrossDomainMessenger
// todo......

// L2StandardERC20Gateway
// todo......

//	L2ERC721Gateway
// todo......

// L2TxFeeVault
// todo......

// L2ERC1155Gateway
// todo......

// MorphStandardERC20
// todo......

// MorphStandardERC20Factory
// todo......

// GasPriceOracle
// yarn hardhat gaspriceoracle-upgrade  --owner 0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E --network qanetl2
// yarn hardhat gaspriceoracle-upgrade  --owner 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d --network hl2
task("gaspriceoracle-upgrade")
    .addParam("owner")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner

        const GPOFactory = await hre.ethers.getContractFactory("GasPriceOracle")
        const gpoNewImpl = await GPOFactory.deploy(_owner)
        await gpoNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`GasPriceOracle new impl deploy at ${gpoNewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)
        let res = await proxyAdmin.upgrade(
            predeploys.GasPriceOracle, 
            gpoNewImpl.address)
        let rec = await res.wait()

        console.log(`upgrade GasPriceOracle ${rec.status == 1 ? "success" : "failed"}`)    
    })

// L2WETHGateway
// todo......

// L2WETH
// todo......

// Record
// yarn hardhat record-upgrade  --owner 0x716173f5BBE0b4B51AaDF5A5840fA9A79D01636E --oracle 0x7161F66BDf7C980B61b426122BBEfff813c0cdF0 --nextbsi 1003 --network qanetl2
// yarn hardhat record-upgrade  --owner 0x48442fdDd92F1000861c7A26cdb5c3a73FFF294d --oracle 0x4844E1AdeF30035E1179C6c924a1B1c2b58E74B0 --nextbsi 196020 --network hl2
task("record-upgrade")
    .addParam("owner")
    .addParam("oracle")
    .addParam("nextbsi")
    .setAction(async (taskArgs, hre) => {
        const _owner = taskArgs.owner
        const _oracle = taskArgs.oracle
        const _nextBatchSubmissionIndex = taskArgs.nextbsi

        const RecordFactory = await hre.ethers.getContractFactory("Record")
        const recordNewImpl = await RecordFactory.deploy()
        await recordNewImpl.deployed()
        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`Record new impl deploy at ${recordNewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)
        let res = await proxyAdmin.upgradeAndCall(
            predeploys.Record, 
            recordNewImpl.address, 
            RecordFactory.interface.encodeFunctionData('initialize',[
                _owner,
                _oracle,
                _nextBatchSubmissionIndex
            ]))
        // console.log("res: ", res)
        let rec = await res.wait()

        console.log(`upgrade record ${rec.status == 1 ? "success" : "failed"}`)    
    })

// MorphToken
// yarn hardhat morph-upgrade --network qanetl2
// yarn hardhat morph-upgrade --network hl2
task("morph-upgrade")
    .setAction(async (taskArgs, hre) => {

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const MorphTokenFactory = await hre.ethers.getContractFactory("MorphToken")
        const morphToken = await MorphTokenFactory.deploy()
        await morphToken.deployed()

        const res = await proxyAdmin.upgrade(
            predeploys.MorphToken,
            morphToken.address
        )
        const rec = await res.wait()
        console.log(`morphToken upgrade ${rec.status === 1}, new impl ${morphToken.address}`)
    })

// Distribute
// yarn hardhat distribute-upgrade --network l2
// yarn hardhat distribute-upgrade --network qanetl2
// yarn hardhat distribute-upgrade --network hl2
task("distribute-upgrade")
    .setAction(async (taskArgs, hre) => {
        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const DistributeFactory = await hre.ethers.getContractFactory("Distribute")
        const distribute = await DistributeFactory.deploy()
        await distribute.deployed()
        const res = await proxyAdmin.upgrade(
            predeploys.Distribute,
            distribute.address
        )
        
        const rec = await res.wait()
        console.log(`distribute upgrade ${rec.status === 1}, new impl ${distribute.address}`)
    })

// L2Staking
// yarn hardhat l2staking-upgrade --newpath ./newFile.json --network qanetl2
task("l2staking-upgrade")
    .addParam("newpath")
    .setAction(async (taskArgs, hre) => {

        const newPath = taskArgs.newpath
        const L1StakingProxyAddr = getContractAddressByName(newPath, ProxyStorageName.L1StakingProxyStorageName)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const L2StakingFactory = await hre.ethers.getContractFactory("L2Staking")
        const staking = await L2StakingFactory.deploy(L1StakingProxyAddr)
        await staking.deployed()

        const res = await proxyAdmin.upgrade(
            predeploys.L2Staking,
            staking.address
        )
        const rec = await res.wait()
        console.log(`L2Staking upgrade ${rec.status === 1}, new impl ${staking.address}`)
    })

// yarn hardhat l2staking-upgrade-hc --l1staking 0xdeab3e2c1f04e543866e8f498316bf6b611bd28b --network qanetl2
// yarn hardhat l2staking-upgrade-hc --l1staking 0x78c5cb68059a7d20b1a65f7380d18a35765a1e00 --network l2
// yarn hardhat l2staking-upgrade-hc --l1staking 0xcb4496399ffd2a94c4c0132c5561b420b8c73ad8 --network hl2
task("l2staking-upgrade-hc")
    .addParam("l1staking")
    .setAction(async (taskArgs, hre) => {
        const _l1Staking = taskArgs.l1staking

        const L2StakingFactory = await hre.ethers.getContractFactory("L2Staking")
        const l2StakingNewImpl = await L2StakingFactory.deploy(_l1Staking)
        await l2StakingNewImpl.deployed()

        let blockNumber = await hre.ethers.provider.getBlockNumber()
        console.log(`L2Staking new impl deploy at ${l2StakingNewImpl.address} and height ${blockNumber}`)

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)
        let res = await proxyAdmin.upgrade(
            predeploys.L2Staking, 
            l2StakingNewImpl.address)
        let rec = await res.wait()

        console.log(`upgrade L2Staking ${rec.status == 1 ? "success" : "failed"}`)    
    })

// L2CustomERC20Gateway
// todo......

// Sequencer
// yarn hardhat sequencer-upgrade --network qanetl2
// cast send --legacy --private-key 0x2035516063ca7724d93e1bfaa01137457355f8e9e1d6cf28cff84def7a478c18 0x5300000000000000000000000000000000000017 "setSequencerSetVerifyHash()" -r http://l2-qa-morph-sentry-0.bitkeep.tools --gas-price 300000000
task("sequencer-upgrade")
    .setAction(async (taskArgs, hre) => {

        const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
        const proxyAdmin = ProxyAdminFactory.attach(predeploys.ProxyAdmin)

        const SequencerFactory = await hre.ethers.getContractFactory("Sequencer")
        const sequencer = await SequencerFactory.deploy()
        await sequencer.deployed()

        const res = await proxyAdmin.upgrade(
            predeploys.Sequencer,
            sequencer.address
        )
        const rec = await res.wait()
        console.log(`sequencer upgrade ${rec.status === 1}, new impl ${sequencer.address}`)
    })

// L2ReverseCustomGateway
// todo......

// L2WithdrawLockERC20Gateway
// todo......

// L2USDCGateway
// todo......

// L2USDC
// todo......

// ===============================================  upgrade  ===============================================