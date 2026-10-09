import { HardhatRuntimeEnvironment } from "hardhat/types";
import { ethers } from "ethers";
import { assertContractVariable, getContractAddressByName } from "./deploy-utils";
import { readProxyAddress, validateDeploymentRecords } from "./deployment-state";
import { firstSequencerAddress, registrationAddresses, validateDeploymentConfig } from "./deployment-validation";
import { ContractFactoryName as F, ImplStorageName as I, ProxyStorageName as P } from "./types";

export const requireDeploymentAddresses = (path: string, names: string[]) => {
    for (const name of names) {
        if (!getContractAddressByName(path, name)) throw new Error(`Deployment file is missing ${name}`);
    }
};

export const validateVerifierVersions = async (hre: HardhatRuntimeEnvironment, path: string) => {
    const activation = process.env.TX_SUBMITTER_BATCH_V2_UPGRADE_TIME || "0";
    if (!/^\d+$/.test(activation)) throw new Error("TX_SUBMITTER_BATCH_V2_UPGRADE_TIME must be a nonnegative integer");
    if (ethers.BigNumber.from(activation).isZero()) return;
    const address = getContractAddressByName(path, I.MultipleVersionRollupVerifierStorageName);
    if (!address) throw new Error("A version 2 proof verifier is required before enabling V2 batches; this deployment flow configures version 1 only");
    const rollupAddress = getContractAddressByName(path, P.RollupProxyStorageName);
    const rollup = await hre.ethers.getContractAt(F.Rollup, rollupAddress);
    const next = (await rollup.lastCommittedBatchIndex()).add(1);
    const verifier = await hre.ethers.getContractAt(F.MultipleVersionRollupVerifier, address);
    const version2 = await verifier.getVerifier(2, next);
    if (version2 === ethers.constants.AddressZero || await hre.ethers.provider.getCode(version2) === "0x") {
        throw new Error(`No valid version 2 proof verifier is configured for batch ${next}; V2 batches cannot be enabled`);
    }
};

// Check initialized contracts while allowing missing deployments and initializations to resume.
export const verifyRecordedConfiguration = async (hre: HardhatRuntimeEnvironment, path: string, config: any) => {
    const address = (name: string) => getContractAddressByName(path, name);
    const empty = address(I.EmptyContract).toLowerCase();
    if (config.l1WETHAddress && address(I.WETH) && config.l1WETHAddress.toLowerCase() !== address(I.WETH).toLowerCase()) {
        throw new Error("l1WETHAddress does not match the recorded WETH address");
    }
    const entries: [string, string, [string, any][]][] = [
        [I.ZkEvmVerifierV1StorageName, F.ZkEvmVerifierV1, [["programVkey", config.programVkey]]],
        [I.Whitelist, F.Whitelist, [["owner", config.contractAdmin]]],
        [P.RollupProxyStorageName, F.Rollup, [
            ["LAYER_2_CHAIN_ID", config.l2ChainID], ["finalizationPeriodSeconds", config.finalizationPeriodSeconds],
            ["proofWindow", config.rollupProofWindow], ["proofRewardPercent", config.proofRewardPercent],
        ]],
        [P.L1MessageQueueWithGasPriceOracleProxyStorageName, F.L1MessageQueueWithGasPriceOracle, [
            ["maxGasLimit", config.l1MessageQueueMaxGasLimit], ["whitelistChecker", address(I.Whitelist)],
            ["MESSENGER", address(P.L1CrossDomainMessengerProxyStorageName)],
            ["ROLLUP_CONTRACT", address(P.RollupProxyStorageName)],
            ["ENFORCED_TX_GATEWAAY", address(P.EnforcedTxGatewayProxyStorageName)],
        ]],
        [P.L1CrossDomainMessengerProxyStorageName, F.L1CrossDomainMessenger, [["feeVault", config.l1FeeVaultRecipient]]],
        [P.EnforcedTxGatewayProxyStorageName, F.EnforcedTxGateway, [["feeVault", config.l1FeeVaultRecipient]]],
        [P.SubmitterProxyStorageName, F.Submitter, [
            ["owner", config.submitterOwner], ["minimumStake", ethers.utils.parseEther(config.submitterMinimumStake.toString())],
            ["challengeDeposit", ethers.utils.parseEther(config.submitterChallengeDeposit.toString())],
            ["rewardPercentage", config.submitterRewardPercentage],
        ]],
    ];
    for (const [name, factory, fields] of entries) {
        const target = address(name);
        if (!target) continue;
        if (name.startsWith("Proxy__") && (await readProxyAddress(hre, target, "implementation")).toLowerCase() === empty) continue;
        const contract = await hre.ethers.getContractAt(factory, target);
        for (const [field, expected] of fields) await assertContractVariable(contract, field, expected);
    }
};

export const verifyDeployment = async (
    hre: HardhatRuntimeEnvironment, path: string, config: any, requireRegistration = true, runtime = false
) => {
    await validateDeploymentConfig(hre, config, "initialize");
    await validateDeploymentRecords(hre, path);
    await validateVerifierVersions(hre, path);
    await verifyRecordedConfiguration(hre, path, config);
    requireDeploymentAddresses(path, [...Object.values(P), I.ProxyAdmin, I.EmptyContract, I.WETH]);
    const address = (name: string) => getContractAddressByName(path, name);
    const manager = address(I.ProxyAdmin).toLowerCase();
    for (const name of Object.values(P)) {
        if ((await readProxyAddress(hre, address(name), "admin")).toLowerCase() !== manager) {
            throw new Error(`${name} admin is not the recorded ProxyAdmin`);
        }
        // The USDC proxy reserves an address; this deployment flow does not configure its token or initialize an implementation.
        if (name === P.L1USDCGatewayProxyStorageName) continue;
        const expected = address(name.replace("Proxy__", "Impl__"));
        if (!expected || (await readProxyAddress(hre, address(name), "implementation")).toLowerCase() !== expected.toLowerCase()) {
            throw new Error(`${name} implementation does not match deployment records`);
        }
    }
    const rollup = await hre.ethers.getContractAt(F.Rollup, address(P.RollupProxyStorageName));
    if (await rollup.committedBatches(0) !== ethers.utils.keccak256(config.batchHeader)) {
        throw new Error("Rollup.committedBatches(0) does not match batchHeader");
    }
    const root = ethers.utils.hexDataSlice(config.batchHeader, 121, 153);
    if (await rollup.finalizedStateRoots(0) !== root || await rollup.committedStateRoots(0) !== root) {
        throw new Error("Rollup genesis state roots do not match batchHeader");
    }
    await assertContractVariable(rollup, "submitterContract", address(P.SubmitterProxyStorageName));
    await assertContractVariable(rollup, "messageQueue", address(P.L1MessageQueueWithGasPriceOracleProxyStorageName));
    await assertContractVariable(rollup, "verifier", address(I.MultipleVersionRollupVerifierStorageName));
    await assertContractVariable(rollup, "LAYER_2_CHAIN_ID", config.l2ChainID);
    await assertContractVariable(rollup, "rollupDelayPeriod", config.rollupDelayPeriod);
    await assertContractVariable(rollup, "finalizationPeriodSeconds", config.finalizationPeriodSeconds);
    await assertContractVariable(rollup, "proofWindow", config.rollupProofWindow);
    if (!(await rollup.isChallenger(config.rollupChallenger))) throw new Error("rollupChallenger is not authorized to challenge");

    const submitter = await hre.ethers.getContractAt(F.Submitter, address(P.SubmitterProxyStorageName));
    await assertContractVariable(submitter, "owner", config.submitterOwner);
    await assertContractVariable(submitter, "rollupContract", rollup.address);
    await assertContractVariable(submitter, "minimumStake", ethers.utils.parseEther(config.submitterMinimumStake.toString()));
    await assertContractVariable(submitter, "challengeDeposit", ethers.utils.parseEther(config.submitterChallengeDeposit.toString()));
    await assertContractVariable(submitter, "rewardPercentage", config.submitterRewardPercentage);
    if (requireRegistration) {
        for (const account of registrationAddresses(config)) {
            if (!(await submitter.isActive(account))) throw new Error(`Submitter ${account} has not completed registration and staking`);
        }
    }
    const first = firstSequencerAddress(config);
    if (first) {
        const sequencer = await hre.ethers.getContractAt(F.L1Sequencer, address(P.L1SequencerProxyStorageName));
        if ((await sequencer.getSequencerAt(0)).toLowerCase() !== first.toLowerCase()) {
            throw new Error("L1Sequencer sequencer at block 0 does not match firstSequencerAddress");
        }
    }
    const queue = await hre.ethers.getContractAt(F.L1MessageQueueWithGasPriceOracle, address(P.L1MessageQueueWithGasPriceOracleProxyStorageName));
    if (!runtime) await assertContractVariable(queue, "l2BaseFee", ethers.utils.parseUnits(config.l2BaseFee.toString(), "gwei"));
    const router = await hre.ethers.getContractAt(F.L1GatewayRouter, address(P.L1GatewayRouterProxyStorageName));
    if ((await router.getERC20Gateway(address(I.WETH))).toLowerCase() !== address(P.L1WETHGatewayProxyStorageName).toLowerCase()) {
        throw new Error("L1GatewayRouter has no matching WETH gateway configuration");
    }
    console.log("L1 deployment verification passed");
};
