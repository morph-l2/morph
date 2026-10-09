import { expect } from "chai";
import hre, { ethers, network, run } from "hardhat";
import fs from "fs";
import os from "os";
import path from "path";
import { getContractAddressByName, readDeploymentRecords, storage } from "../src/deploy-utils";
import { deployRecordedContract, readProxyAddress, validateDeploymentRecords, validateDeploymentSigner } from "../src/deployment-state";
import { validateDeploymentConfig } from "../src/deployment-validation";
import { ImplStorageName as I, ProxyStorageName as P } from "../src/types";

const genesisHeader = () => {
    const bytes = Buffer.alloc(257);
    bytes[0] = 2;
    bytes[25] = 1;
    Buffer.from("010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014", "hex").copy(bytes, 57);
    Buffer.alloc(32, 17).copy(bytes, 121);
    return `0x${bytes.toString("hex")}`;
};

async function rejects(action: () => Promise<any>, message: string) {
    let failure: Error | undefined;
    try { await action(); } catch (error) { failure = error as Error; }
    expect(failure, `expected an error containing ${message}`).not.to.equal(undefined);
    expect(failure?.message).to.contain(message);
}

describe("L1 deployment validation and recovery", function () {
    this.timeout(120000);
    let directory: string;
    let storagepath: string;
    let previousConfig: any;
    let previousKeys: string | undefined;
    let previousFirst: string | undefined;
    let previousV2: string | undefined;
    let previousSubmitterOwnerKey: string | undefined;
    let config: any;

    beforeEach(async () => {
        await network.provider.send("hardhat_reset");
        directory = fs.mkdtempSync(path.join(os.tmpdir(), "morph-deployment-"));
        storagepath = path.join(directory, "deployments.json");
        previousConfig = hre.deployConfig;
        previousKeys = process.env.batchSubmitterPks;
        previousFirst = process.env.firstSequencerAddress;
        previousV2 = process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME;
        previousSubmitterOwnerKey = process.env.SUBMITTER_OWNER_PRIVATE_KEY;
        delete process.env.batchSubmitterPks;
        delete process.env.firstSequencerAddress;
        delete process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME;
        delete process.env.SUBMITTER_OWNER_PRIVATE_KEY;
        const [owner, operator, challenger, sequencer] = await ethers.getSigners();
        config = {
            l1ChainID: 900, l2ChainID: 53077, l1WETHAddress: "",
            contractAdmin: owner.address, l1FeeVaultRecipient: owner.address,
            l1MessageQueueMaxGasLimit: 30000000, l2BaseFee: 0.1,
            programVkey: ethers.utils.hexZeroPad("0x01", 32),
            finalizationPeriodSeconds: 10, rollupProofWindow: 86400,
            proofRewardPercent: 70, rollupDelayPeriod: 600,
            batchSubmitterAddresses: [operator.address], rollupChallenger: challenger.address,
            batchHeader: genesisHeader(), submitterOwner: owner.address,
            submitterMinimumStake: 1, submitterChallengeDeposit: 1,
            submitterRewardPercentage: 50, firstSequencerAddress: sequencer.address,
        };
        hre.deployConfig = config;
    });

    afterEach(() => {
        hre.deployConfig = previousConfig;
        if (previousKeys === undefined) delete process.env.batchSubmitterPks;
        else process.env.batchSubmitterPks = previousKeys;
        if (previousFirst === undefined) delete process.env.firstSequencerAddress;
        else process.env.firstSequencerAddress = previousFirst;
        if (previousV2 === undefined) delete process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME;
        else process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME = previousV2;
        if (previousSubmitterOwnerKey === undefined) delete process.env.SUBMITTER_OWNER_PRIVATE_KEY;
        else process.env.SUBMITTER_OWNER_PRIVATE_KEY = previousSubmitterOwnerKey;
        for (const name of fs.readdirSync(directory)) fs.unlinkSync(path.join(directory, name));
        fs.rmdirSync(directory);
    });

    it("rejects missing configuration and invalid genesis headers without broadcasting transactions", async () => {
        const [owner] = await ethers.getSigners();
        delete config.l1WETHAddress;
        await rejects(() => run("deploy", { storagepath }), "l1WETHAddress");
        expect(await owner.getTransactionCount()).to.equal(0);
        config.l1WETHAddress = "";
        config.rollupDelayPeriod = 0;
        await rejects(() => run("deploy", { storagepath }), "rollupDelayPeriod");
        expect(await owner.getTransactionCount()).to.equal(0);
        config.rollupDelayPeriod = 600;
        config.batchHeader = "";
        await rejects(() => run("initialize", { storagepath }), "batchHeader");
        expect(await owner.getTransactionCount()).to.equal(0);
        expect(fs.existsSync(storagepath)).to.equal(false);
    });

    it("requires an explicit QA RPC endpoint and first sequencer", async () => {
        const previousRPC = process.env.QA_RPC_URL;
        const qa = { ...hre, network: { ...hre.network, name: "qanetl1" } } as any;
        try {
            delete process.env.QA_RPC_URL;
            await rejects(() => validateDeploymentConfig(qa, config, "deploy"), "QA_RPC_URL");
            process.env.QA_RPC_URL = "http://127.0.0.1:1";
            config.firstSequencerAddress = "";
            await rejects(() => validateDeploymentConfig(qa, config, "initialize"), "firstSequencerAddress");
        } finally {
            if (previousRPC === undefined) delete process.env.QA_RPC_URL;
            else process.env.QA_RPC_URL = previousRPC;
        }
    });

    it("rejects deployment when the signer has pending transactions", async () => {
        const [owner, recipient] = await ethers.getSigners();
        await network.provider.send("evm_setAutomine", [false]);
        try {
            await owner.sendTransaction({ to: recipient.address, value: 1 });
            await rejects(() => run("deploy", { storagepath }), "Deployer has pending transactions");
            expect(fs.existsSync(storagepath)).to.equal(false);
            expect(await owner.getTransactionCount("pending")).to.equal(1);
        } finally {
            await network.provider.send("evm_setAutomine", [true]);
            await network.provider.send("evm_mine");
        }
    });

    it("rejects an explicit registration owner key mismatch before deployment", async () => {
        const [owner] = await ethers.getSigners();
        process.env.SUBMITTER_OWNER_PRIVATE_KEY = ethers.Wallet.createRandom().privateKey;
        await rejects(() => run("deploy", { storagepath }), "does not match submitterOwner");
        expect(await owner.getTransactionCount()).to.equal(0);
        expect(fs.existsSync(storagepath)).to.equal(false);
        process.env.SUBMITTER_OWNER_PRIVATE_KEY = "invalid";
        await rejects(() => run("deploy", { storagepath }), "is not a valid private key");
        expect(await owner.getTransactionCount()).to.equal(0);
    });

    it("rejects deployment when V2 batches are enabled without a V2 verifier", async () => {
        process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME = "1";
        await rejects(() => run("deploy", { storagepath }), "version 2 proof verifier");
        const [owner] = await ethers.getSigners();
        expect(await owner.getTransactionCount()).to.equal(0);
    });

    it("checks managed proxy administrators while preserving independently administered legacy records", async () => {
        const [owner, legacyAdmin] = await ethers.getSigners();
        await deployRecordedContract(hre, storagepath, owner, I.ProxyAdmin, "ProxyAdmin");
        const empty = await deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract");
        const factory = await ethers.getContractFactory("TransparentUpgradeableProxy", owner);
        const legacy = await factory.deploy(empty.address, legacyAdmin.address, "0x");
        const receipt = await legacy.deployTransaction.wait();
        await storage(storagepath, "Proxy__L1Staking", legacy.address, receipt.blockNumber);
        const managed = await deployRecordedContract(hre, storagepath, owner, P.SubmitterProxyStorageName,
            "TransparentUpgradeableProxy", [empty.address, owner.address, "0x"]);
        const before = await owner.getTransactionCount();
        await validateDeploymentRecords(hre, storagepath);
        await validateDeploymentSigner(hre, storagepath, owner);
        expect(await owner.getTransactionCount()).to.equal(before);
        expect(await readProxyAddress(hre, legacy.address, "admin")).to.equal(legacyAdmin.address);

        const managedProxy = await ethers.getContractAt("ITransparentUpgradeableProxy", managed.address, owner);
        await (await managedProxy.changeAdmin(legacyAdmin.address)).wait();
        const afterTransfer = await owner.getTransactionCount();
        await rejects(() => validateDeploymentSigner(hre, storagepath, owner),
            "Proxy__Submitter proxy admin is neither the deployer nor the recorded ProxyAdmin");
        expect(await owner.getTransactionCount()).to.equal(afterTransfer);
        const records = readDeploymentRecords(storagepath);
        records.find(record => record.name === "Proxy__L1Staking").address = legacyAdmin.address;
        fs.writeFileSync(storagepath, JSON.stringify(records));
        await rejects(() => validateDeploymentRecords(hre, storagepath), "Proxy__L1Staking: recorded address");
        expect(await owner.getTransactionCount()).to.equal(afterTransfer);
    });

    it("resumes after admin transfer and genesis import and repeats without new transactions", async () => {
        await run("deploy", { storagepath, concurrent: "true" });
        const [owner] = await ethers.getSigners();
        const deployCount = await owner.getTransactionCount();
        const originalRecords = readDeploymentRecords(storagepath).length;
        await run("deploy", { storagepath });
        expect(await owner.getTransactionCount()).to.equal(deployCount);
        expect(readDeploymentRecords(storagepath).length).to.equal(originalRecords);
        const originalGetContractAt = hre.ethers.getContractAt;
        hre.ethers.getContractAt = (async (...args: any[]) => {
            const contract = await (originalGetContractAt as any)(...args);
            if (args[0] === "Rollup") return { ...contract, initialize3: async () => { throw new Error("simulated failure before initialization broadcast"); } };
            return contract;
        }) as any;
        try {
            await rejects(() => run("initialize", { storagepath }), "simulated failure before initialization broadcast");
        } finally { hre.ethers.getContractAt = originalGetContractAt; }
        const rollupAddress = getContractAddressByName(storagepath, P.RollupProxyStorageName);
        expect(await readProxyAddress(hre, rollupAddress, "admin")).to.equal(
            ethers.utils.getAddress(getContractAddressByName(storagepath, I.ProxyAdmin))
        );
        const rollup = await ethers.getContractAt("Rollup", rollupAddress);
        expect(await rollup.committedBatches(0)).to.equal(ethers.utils.keccak256(config.batchHeader));
        expect(await rollup.rollupDelayPeriod()).to.equal(0);
        await run("initialize", { storagepath });
        await run("register", { storagepath });
        await run("verify-deployment", { storagepath });
        const count = await owner.getTransactionCount();
        const records = fs.readFileSync(storagepath, "utf8");
        await run("initialize", { storagepath });
        await run("register", { storagepath });
        expect(await owner.getTransactionCount()).to.equal(count);
        expect(fs.readFileSync(storagepath, "utf8")).to.equal(records);
    });

    it("rejects an unauthorized registration signer and detects inactive submitters", async () => {
        const [, otherOwner] = await ethers.getSigners();
        config.submitterOwner = otherOwner.address;
        await run("deploy", { storagepath });
        await run("initialize", { storagepath });
        const [deployer] = await ethers.getSigners();
        const before = await deployer.getTransactionCount();
        await rejects(() => run("register", { storagepath }), "registration signer is not Submitter owner");
        expect(await deployer.getTransactionCount()).to.equal(before);
        await rejects(() => run("verify-deployment", { storagepath }), "has not completed registration and staking");
    });

    it("restores initialization through ProxyAdmin while preserving the Rollup owner", async () => {
        await run("deploy", { storagepath });
        const [owner] = await ethers.getSigners();
        const address = getContractAddressByName(storagepath, P.RollupProxyStorageName);
        const manager = getContractAddressByName(storagepath, I.ProxyAdmin);
        const proxy = await ethers.getContractAt("ITransparentUpgradeableProxy", address, owner);
        await (await proxy.changeAdmin(manager)).wait();
        await run("initialize", { storagepath });
        const rollup = await ethers.getContractAt("Rollup", address);
        expect(await rollup.owner()).to.equal(owner.address);
        expect(await readProxyAddress(hre, address, "admin")).to.equal(ethers.utils.getAddress(manager));
    });

    it("rejects configuration mismatches during retries and verification without sending transactions", async () => {
        await run("deploy", { storagepath });
        await run("initialize", { storagepath });
        await run("register", { storagepath });
        const [owner, another] = await ethers.getSigners();
        const before = await owner.getTransactionCount();
        const cases: [string, any, string][] = [
            ["proofRewardPercent", 60, "proofRewardPercent"],
            ["l1MessageQueueMaxGasLimit", 20000000, "maxGasLimit"],
            ["l1FeeVaultRecipient", another.address, "feeVault"],
            ["programVkey", ethers.utils.hexZeroPad("0x02", 32), "programVkey"],
        ];
        for (const [key, changed, field] of cases) {
            const previous = config[key];
            config[key] = changed;
            try {
                await rejects(() => run("initialize", { storagepath }), field);
                await rejects(() => run("verify-deployment", { storagepath }), field);
                expect(await owner.getTransactionCount()).to.equal(before);
            } finally { config[key] = previous; }
        }
    });

    it("allows runtime fee changes only with runtime verification", async () => {
        await run("deploy", { storagepath });
        await run("initialize", { storagepath });
        await run("register", { storagepath });
        const queue = await ethers.getContractAt("L1MessageQueueWithGasPriceOracle",
            getContractAddressByName(storagepath, P.L1MessageQueueWithGasPriceOracleProxyStorageName));
        await (await queue.setL2BaseFee(ethers.utils.parseUnits("0.2", "gwei"))).wait();
        await rejects(() => run("verify-deployment", { storagepath }), "l2BaseFee");
        await run("verify-deployment", { storagepath, runtime: true });
    });

    it("leaves no pending records when balance or constructor checks fail before broadcasting", async () => {
        const [owner] = await ethers.getSigners();
        await network.provider.send("hardhat_setBalance", [owner.address, "0x0"]);
        await rejects(() => deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract"), "insufficient deployer balance");
        expect(fs.existsSync(storagepath)).to.equal(false);
        await network.provider.send("hardhat_setBalance", [owner.address, ethers.utils.parseEther("10").toHexString()]);
        await rejects(() => deployRecordedContract(hre, storagepath, owner, I.MultipleVersionRollupVerifierStorageName, "MultipleVersionRollupVerifier", [[1], [ethers.constants.AddressZero]]), "reverted");
        expect(fs.existsSync(storagepath)).to.equal(false);
        await deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract");
        expect(readDeploymentRecords(storagepath)[0].pending).to.equal(false);
    });

    it("records transaction hashes before waiting and reuses confirmed deployments on retries", async () => {
        const [owner] = await ethers.getSigners();
        const original = hre.ethers.getContractFactory;
        hre.ethers.getContractFactory = (async (...args: any[]) => {
            const factory = await (original as any)(...args);
            return { getDeployTransaction: (...args: any[]) => factory.getDeployTransaction(...args), deploy: async (...constructorArgs: any[]) => {
                const deployed = await factory.deploy(...constructorArgs);
                return { ...deployed, deployTransaction: { ...deployed.deployTransaction,
                    wait: async () => { throw new Error("simulated interruption while waiting for receipt"); },
                } };
            } };
        }) as any;
        try {
            await rejects(() => deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract"), "simulated interruption while waiting for receipt");
        } finally { hre.ethers.getContractFactory = original; }
        const pending = readDeploymentRecords(storagepath)[0];
        expect(pending.pending).to.equal(true);
        expect(pending.transactionHash).to.match(/^0x[0-9a-f]{64}$/);
        const before = await owner.getTransactionCount();
        const resumed = await deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract");
        expect(resumed.address.toLowerCase()).to.equal(pending.address);
        expect(await owner.getTransactionCount()).to.equal(before);
        expect(readDeploymentRecords(storagepath)[0].pending).to.equal(false);
    });

    it("preserves expected addresses and prevents duplicate deployments after an unknown broadcast result", async () => {
        const [owner] = await ethers.getSigners();
        const original = hre.ethers.getContractFactory;
        let transactionHash = "";
        hre.ethers.getContractFactory = (async (...args: any[]) => {
            const factory = await (original as any)(...args);
            return { getDeployTransaction: (...args: any[]) => factory.getDeployTransaction(...args), deploy: async (...constructorArgs: any[]) => {
                const deployed = await factory.deploy(...constructorArgs);
                transactionHash = deployed.deployTransaction.hash;
                throw new Error("simulated loss of RPC broadcast response");
            } };
        }) as any;
        try {
            await rejects(() => deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract"), "simulated loss of RPC broadcast response");
        } finally { hre.ethers.getContractFactory = original; }
        const record = readDeploymentRecords(storagepath)[0];
        expect(record.nonce).to.equal(0);
        expect(record.deployer).to.equal(owner.address);
        expect(record.transactionHash).to.equal(undefined);
        const before = await owner.getTransactionCount();
        await rejects(() => deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract"), "broadcast result was not recorded");
        expect(await owner.getTransactionCount()).to.equal(before);
        // Simulate an operator checking the original on-chain transaction and recording its hash before retrying.
        fs.writeFileSync(storagepath, JSON.stringify([{ ...record, transactionHash }]));
        const contract = await deployRecordedContract(hre, storagepath, owner, I.EmptyContract, "EmptyContract");
        expect(contract.address.toLowerCase()).to.equal(record.address);
        expect(await owner.getTransactionCount()).to.equal(before);
    });
});
