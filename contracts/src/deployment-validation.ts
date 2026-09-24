import { ethers } from "ethers";
import { HardhatRuntimeEnvironment } from "hardhat/types";

const value = (config: any, key: string) => key in config ? config[key] : undefined;
const address = (input: any, key: string) => {
    if (!ethers.utils.isAddress(input) || input === ethers.constants.AddressZero) {
        throw new Error(`${key} must be a nonzero address`);
    }
};

export const firstSequencerAddress = (config: any) =>
    process.env.firstSequencerAddress || value(config, "firstSequencerAddress") || "";

export const validateGenesisHeader = (header: string) => {
    if (!ethers.utils.isHexString(header)) throw new Error("batchHeader must be a hexadecimal genesis batch header");
    const bytes = ethers.utils.arrayify(header);
    if (!((bytes[0] === 0 && bytes.length === 249) || ((bytes[0] === 1 || bytes[0] === 2) && bytes.length === 257))) {
        throw new Error("batchHeader has an invalid genesis header version or length");
    }
    const field = (start: number, end: number) => ethers.utils.hexlify(bytes.slice(start, end));
    if (!ethers.BigNumber.from(field(1, 9)).isZero() || !ethers.BigNumber.from(field(9, 17)).isZero()) {
        throw new Error("batchHeader batchIndex and l1MessagePopped must be zero");
    }
    if (field(25, 57) === ethers.constants.HashZero || field(121, 153) === ethers.constants.HashZero) {
        throw new Error("batchHeader dataHash and postStateRoot must be nonzero");
    }
    if (field(57, 89) !== "0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014") {
        throw new Error("batchHeader blobVersionedHash does not match the Rollup genesis requirement");
    }
};

export const validateDeploymentConfig = async (
    hre: HardhatRuntimeEnvironment, config: any, phase: "deploy" | "initialize" | "register"
) => {
    if (hre.network.name === "qanetl1" && !process.env.QA_RPC_URL) {
        throw new Error("qanetl1 requires an explicit QA_RPC_URL");
    }
    const chainID = value(config, "l1ChainID");
    if ((await hre.ethers.provider.getNetwork()).chainId !== chainID) {
        throw new Error("l1ChainID does not match the chainId returned by the current RPC endpoint");
    }
    registrationAddresses(config);
    if (phase === "register") return;
    const weth = value(config, "l1WETHAddress");
    if (weth !== "") address(weth, "l1WETHAddress (use an empty string to deploy WETH)");
    if (weth && await hre.ethers.provider.getCode(weth) === "0x") {
        throw new Error("l1WETHAddress has no contract code on the current network");
    }
    for (const key of ["contractAdmin", "l1FeeVaultRecipient", "submitterOwner", "rollupChallenger"]) {
        address(value(config, key), key);
    }
    if (process.env.SUBMITTER_OWNER_PRIVATE_KEY) {
        let owner: string;
        try { owner = new ethers.Wallet(process.env.SUBMITTER_OWNER_PRIVATE_KEY).address; }
        catch (_) { throw new Error("SUBMITTER_OWNER_PRIVATE_KEY is not a valid private key"); }
        if (owner.toLowerCase() !== config.submitterOwner.toLowerCase()) {
            throw new Error("SUBMITTER_OWNER_PRIVATE_KEY does not match submitterOwner");
        }
    }
    for (const key of ["l2ChainID", "l1MessageQueueMaxGasLimit", "finalizationPeriodSeconds", "rollupProofWindow", "rollupDelayPeriod"]) {
        const number = value(config, key);
        if (!Number.isSafeInteger(number) || number <= 0) throw new Error(`${key} must be a positive integer`);
    }
    for (const key of ["submitterMinimumStake", "submitterChallengeDeposit"]) {
        const amount = value(config, key);
        if (typeof amount !== "number" || !Number.isFinite(amount) || amount <= 0) throw new Error(`${key} must be greater than zero`);
        ethers.utils.parseEther(amount.toString());
    }
    for (const key of ["submitterRewardPercentage", "proofRewardPercent"]) {
        const percent = value(config, key);
        if (!Number.isInteger(percent) || percent <= 0 || percent > 100) throw new Error(`${key} must be an integer from 1 to 100`);
    }
    const fee = value(config, "l2BaseFee");
    if (typeof fee !== "number" || !Number.isFinite(fee) || fee < 0) throw new Error("l2BaseFee must be nonnegative");
    ethers.utils.parseUnits(fee.toString(), "gwei");
    if (!ethers.utils.isHexString(value(config, "programVkey"), 32)) throw new Error("programVkey must be bytes32");
    const first = firstSequencerAddress(config);
    if (first) address(first, "firstSequencerAddress");
    if (hre.network.name === "qanetl1" && !first) throw new Error("qanetl1 requires firstSequencerAddress");
    if (phase === "initialize") {
        validateGenesisHeader(value(config, "batchHeader"));
    }
};

export const registrationAddresses = (config: any): string[] => {
    const privateKeys = JSON.parse(process.env.batchSubmitterPks || "[]");
    if (!Array.isArray(privateKeys)) throw new Error("batchSubmitterPks must be a JSON array");
    const addresses = privateKeys.length > 0
        ? privateKeys.map((key, index) => {
            try { return new ethers.Wallet(key).address; }
            catch (_) { throw new Error(`batchSubmitterPks private key at index ${index} is invalid`); }
        })
        : value(config, "batchSubmitterAddresses");
    if (!Array.isArray(addresses) || addresses.length === 0) throw new Error("batchSubmitterAddresses must contain at least one address");
    for (const input of addresses) address(input, "batchSubmitterAddresses");
    if (new Set(addresses.map(input => input.toLowerCase())).size !== addresses.length) {
        throw new Error("batchSubmitterAddresses must not contain duplicate addresses");
    }
    return addresses;
};
