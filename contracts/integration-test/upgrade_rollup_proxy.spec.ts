import { expect } from "chai"
import { ethers, network, run } from "hardhat"

describe("upgradeRollupProxy", () => {
    const layer2ChainID = 900
    const lastFinalizedBatchIndexSlot = 157
    const lastCommittedBatchIndexSlot = 158
    const minimumStake = ethers.utils.parseEther("1")
    const challengeDeposit = ethers.utils.parseEther("2")
    const rewardPercentage = 25

    async function setUintStorage(address: string, slot: number, value: number): Promise<void> {
        await network.provider.send("hardhat_setStorageAt", [
            address,
            ethers.utils.hexValue(slot),
            ethers.utils.hexZeroPad(ethers.utils.hexlify(value), 32),
        ])
        await network.provider.send("evm_mine")
    }

    async function deployPendingRollup(): Promise<{
        proxyAdminAddress: string
        rollupProxyAddress: string
        submitterProxyAddress: string
        submitterImplementationAddress: string
        activeOperatorAddress: string
        inactiveOperatorAddress: string
    }> {
        const [owner, messageQueue, verifier, activeOperator, inactiveOperator] = await ethers.getSigners()
        const rollupFactory = await ethers.getContractFactory("Rollup")
        const rollupImplementation = await rollupFactory.deploy(layer2ChainID)
        await rollupImplementation.deployed()

        const proxyAdminFactory = await ethers.getContractFactory("ProxyAdmin")
        const proxyAdmin = await proxyAdminFactory.deploy()
        await proxyAdmin.deployed()

        const initializer = rollupFactory.interface.encodeFunctionData("initialize", [
            owner.address,
            messageQueue.address,
            verifier.address,
            2 * 24 * 60 * 60,
            60 * 60,
            50,
        ])
        const proxyFactory = await ethers.getContractFactory("TransparentUpgradeableProxy")
        const rollupProxy = await proxyFactory.deploy(rollupImplementation.address, proxyAdmin.address, initializer)
        await rollupProxy.deployed()

        const submitterFactory = await ethers.getContractFactory("Submitter")
        const submitterImplementation = await submitterFactory.deploy()
        await submitterImplementation.deployed()
        const submitterInitializer = submitterFactory.interface.encodeFunctionData("initialize", [
            owner.address,
            rollupProxy.address,
            minimumStake,
            challengeDeposit,
            rewardPercentage,
        ])
        const submitterProxy = await proxyFactory.deploy(
            submitterImplementation.address,
            proxyAdmin.address,
            submitterInitializer
        )
        await submitterProxy.deployed()
        const submitter = await ethers.getContractAt("Submitter", submitterProxy.address)
        await (await submitter.addSubmitter(activeOperator.address)).wait()
        await (await submitter.stake(activeOperator.address, { value: minimumStake })).wait()

        await setUintStorage(rollupProxy.address, lastFinalizedBatchIndexSlot, 5)
        await setUintStorage(rollupProxy.address, lastCommittedBatchIndexSlot, 7)

        return {
            proxyAdminAddress: proxyAdmin.address,
            rollupProxyAddress: rollupProxy.address,
            submitterProxyAddress: submitterProxy.address,
            submitterImplementationAddress: submitterImplementation.address,
            activeOperatorAddress: activeOperator.address,
            inactiveOperatorAddress: inactiveOperator.address,
        }
    }

    it("upgrades with pending batches and persists the operator-provided cutover", async () => {
        const deployment = await deployPendingRollup()

        await run("upgradeRollupProxy", {
            proxyadminaddr: deployment.proxyAdminAddress,
            rollupproxyaddr: deployment.rollupProxyAddress,
            submitterproxyaddr: deployment.submitterProxyAddress,
            submitteroperatoraddr: deployment.activeOperatorAddress,
            expectedcutoverbatchindex: "7",
        })

        const rollup = await ethers.getContractAt("Rollup", deployment.rollupProxyAddress)
        expect(await rollup.lastFinalizedBatchIndex()).to.equal(5)
        expect(await rollup.lastCommittedBatchIndex()).to.equal(7)
        expect(await rollup.legacyCutoverBatchIndex()).to.equal(7)
        expect(await rollup.submitterContract()).to.equal(deployment.submitterProxyAddress)
        expect(await rollup.LAYER_2_CHAIN_ID()).to.equal(layer2ChainID)
    })

    it("rejects a stale expected cutover before deploying the upgrade", async () => {
        const deployment = await deployPendingRollup()
        let failure: Error | undefined

        try {
            await run("upgradeRollupProxy", {
                proxyadminaddr: deployment.proxyAdminAddress,
                rollupproxyaddr: deployment.rollupProxyAddress,
                submitterproxyaddr: deployment.submitterProxyAddress,
                submitteroperatoraddr: deployment.activeOperatorAddress,
                expectedcutoverbatchindex: "6",
            })
        } catch (error) {
            failure = error as Error
        }

        expect(failure).not.to.equal(undefined)
        expect(failure?.message).to.contain("Rollup cutover batch mismatch: expected=6 committed=7")
    })

    it("rejects a bare Submitter implementation", async () => {
        const deployment = await deployPendingRollup()
        let failure: Error | undefined

        try {
            await run("upgradeRollupProxy", {
                proxyadminaddr: deployment.proxyAdminAddress,
                rollupproxyaddr: deployment.rollupProxyAddress,
                submitterproxyaddr: deployment.submitterImplementationAddress,
                submitteroperatoraddr: deployment.activeOperatorAddress,
                expectedcutoverbatchindex: "7",
            })
        } catch (error) {
            failure = error as Error
        }

        expect(failure).not.to.equal(undefined)
        expect(failure?.message).to.contain("Submitter rollup mismatch")
    })

    it("rejects a Submitter proxy when the supplied operator is not active", async () => {
        const deployment = await deployPendingRollup()
        let failure: Error | undefined

        try {
            await run("upgradeRollupProxy", {
                proxyadminaddr: deployment.proxyAdminAddress,
                rollupproxyaddr: deployment.rollupProxyAddress,
                submitterproxyaddr: deployment.submitterProxyAddress,
                submitteroperatoraddr: deployment.inactiveOperatorAddress,
                expectedcutoverbatchindex: "7",
            })
        } catch (error) {
            failure = error as Error
        }

        expect(failure).not.to.equal(undefined)
        expect(failure?.message).to.contain("Submitter operator is not active")
    })
})
