import { createHash } from "crypto";
import fs from "fs";
import path from "path";

import { BigNumber, Contract, Wallet, ethers } from "ethers";

const ADMIN_SLOT = "0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103";
const IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";
const ZERO_ADDRESS = ethers.constants.AddressZero;

export type RetirementMode = "plan" | "dry-run" | "execute" | "verify";

export interface ReplayEvent {
    name: string;
    args: Record<string, unknown>;
    blockNumber: number;
    blockHash: string;
    transactionHash: string;
    logIndex: number;
}

export interface CandidateState {
    address: string;
    firstEvent?: string;
    lastEvent?: string;
    everRegistered: boolean;
    withdrawn: boolean;
    claimed: boolean;
    slashed: boolean;
    removed: boolean;
    currentWhitelist?: boolean;
    currentActive?: boolean;
    withdrawalUnlockBlock?: string;
}

export interface RetirementOperation {
    kind: "freeze" | "withdraw" | "claim" | "claim-slash" | "renounce-owner" | "burn-admin";
    actor: string;
    target: string;
    value: string;
    data: string;
    description: string;
    nonce?: number;
    expectedState: string[];
    txHash?: string;
    receiptBlockHash?: string;
}

export interface RelayEvidence {
    messageHash: string;
    addresses?: string[];
    transactionHash: string;
    blockHash: string;
    success: boolean;
    canonical: boolean;
}

export interface RetirementManifest {
    version: number;
    legacy: {
        proxy: string;
        implementation?: string;
        deployBlock: number;
        deployBlockHash?: string;
        proxyAdmin?: string;
        proxyAdminOwner?: string;
        proxyAdminIsDedicated: boolean;
        burnAdmin: string;
    };
    retirement: {
        confirmations: number;
        stakingValue?: string;
        treasury: string;
    };
    snapshot?: {
        chainId: number;
        safeHead: number;
        safeHeadHash: string;
        stakingValue: string;
        contractBalance: string;
        expectedLiabilities: string;
        unexplainedBalance: string;
        slashRemaining: string;
        owner: string;
        activeStakers: string[];
    };
    candidates: CandidateState[];
    canonicalEventIds: string[];
    checkpoints: Array<{ blockNumber: number; blockHash: string }>;
    relays: RelayEvidence[];
    operations: RetirementOperation[];
    verifiedAt?: string;
}

function normalizeAddress(value: string): string {
    return ethers.utils.getAddress(value);
}

function eventId(event: ReplayEvent): string {
    return `${event.blockHash}:${event.transactionHash}:${event.logIndex}`;
}

function eventAddresses(value: unknown): string[] {
    if (!Array.isArray(value)) return [];
    return value.filter((item): item is string => typeof item === "string").map(normalizeAddress);
}

/**
 * Rebuilds state from scratch against canonical block hashes. Rebuilding instead of appending
 * means orphaned tip events disappear automatically after an L1 reorg.
 */
export function rebuildLegacyState(
    events: ReplayEvent[],
    canonicalBlockHashes: Record<number, string>
): { candidates: CandidateState[]; canonicalEventIds: string[] } {
    const candidates = new Map<string, CandidateState>();
    const canonicalEventIds: string[] = [];

    const get = (address: string, identity: string): CandidateState => {
        const normalized = normalizeAddress(address);
        const key = normalized.toLowerCase();
        let state = candidates.get(key);
        if (!state) {
            state = {
                address: normalized,
                firstEvent: identity,
                lastEvent: identity,
                everRegistered: false,
                withdrawn: false,
                claimed: false,
                slashed: false,
                removed: false,
            };
            candidates.set(key, state);
        }
        state.lastEvent = identity;
        return state;
    };

    for (const event of [...events].sort((a, b) => a.blockNumber - b.blockNumber || a.logIndex - b.logIndex)) {
        const canonicalHash = canonicalBlockHashes[event.blockNumber];
        if (!canonicalHash || canonicalHash.toLowerCase() !== event.blockHash.toLowerCase()) continue;
        const identity = eventId(event);
        canonicalEventIds.push(identity);

        if (event.name === "WhitelistUpdated") {
            for (const address of eventAddresses(event.args.add)) get(address, identity);
            for (const address of eventAddresses(event.args.remove)) get(address, identity);
        } else if (event.name === "Registered") {
            const state = get(String(event.args.addr), identity);
            state.everRegistered = true;
            state.withdrawn = false;
            state.claimed = false;
            state.slashed = false;
            state.removed = false;
        } else if (event.name === "Withdrawn") {
            get(String(event.args.addr), identity).withdrawn = true;
        } else if (event.name === "Claimed") {
            const state = get(String(event.args.staker), identity);
            state.claimed = true;
            state.withdrawn = false;
        } else if (event.name === "Slashed") {
            for (const address of eventAddresses(event.args.stakers)) {
                const state = get(address, identity);
                state.slashed = true;
                state.removed = true;
            }
        } else if (event.name === "StakersRemoved") {
            for (const address of eventAddresses(event.args.stakers)) get(address, identity).removed = true;
        }
    }

    return {
        candidates: [...candidates.values()].sort((a, b) => a.address.localeCompare(b.address)),
        canonicalEventIds,
    };
}

export function verifyExecutePreconditions(
    manifest: RetirementManifest,
    current: { stakingValue: BigNumber | string; deployBlockHash?: string }
): void {
    if (!manifest.retirement.stakingValue) throw new Error("retirement stakingValue snapshot missing");
    if (!BigNumber.from(current.stakingValue).eq(manifest.retirement.stakingValue)) {
        throw new Error("stakingValue changed during retirement");
    }
    if (
        manifest.legacy.deployBlockHash &&
        current.deployBlockHash &&
        manifest.legacy.deployBlockHash.toLowerCase() !== current.deployBlockHash.toLowerCase()
    ) {
        throw new Error("legacy deployment block is no longer canonical");
    }
}

export function retirementBlockers(manifest: RetirementManifest): string[] {
    const blockers: string[] = [];
    if (!manifest.snapshot) return ["snapshot missing"];
    for (const candidate of manifest.candidates) {
        if (candidate.currentWhitelist) blockers.push(`${candidate.address}: whitelist still enabled`);
        if (candidate.currentActive) blockers.push(`${candidate.address}: staker still active`);
        if (BigNumber.from(candidate.withdrawalUnlockBlock || 0).gt(0)) {
            blockers.push(`${candidate.address}: withdrawal still pending`);
        }
    }
    if (!BigNumber.from(manifest.snapshot.slashRemaining).isZero()) blockers.push("slashRemaining is nonzero");
    if (!BigNumber.from(manifest.snapshot.contractBalance).isZero()) blockers.push("legacy contract balance is nonzero");
    for (const relay of manifest.relays) {
        if (!relay.success || !relay.canonical) blockers.push(`${relay.messageHash}: removeStakers relay incomplete`);
    }
    for (const candidate of manifest.candidates.filter(
        (item) => item.removed || item.withdrawn || item.slashed
    )) {
        const covered = manifest.relays.some(
            (relay) =>
                relay.success &&
                relay.canonical &&
                (relay.addresses || []).some(
                    (address) => address.toLowerCase() === candidate.address.toLowerCase()
                )
        );
        if (!covered) blockers.push(`${candidate.address}: canonical removeStakers relay evidence missing`);
    }
    if (!manifest.legacy.proxyAdminIsDedicated) blockers.push("legacy proxy admin is not confirmed dedicated");
    return blockers;
}

function archivedAbiPath(): string {
    return path.resolve(__dirname, "../abi/archive/legacy-l1-staking/L1Staking.json");
}

function loadFrozenAbi(): unknown[] {
    const abiPath = archivedAbiPath();
    const archiveManifestPath = path.resolve(__dirname, "../abi/archive/manifest.json");
    if (!fs.existsSync(abiPath) || !fs.existsSync(archiveManifestPath)) {
        throw new Error("frozen legacy L1Staking ABI archive is missing");
    }
    const raw = fs.readFileSync(abiPath);
    const archiveManifest = JSON.parse(fs.readFileSync(archiveManifestPath, "utf8"));
    const expected = archiveManifest.artifacts?.["legacy-l1-staking/L1Staking.json"];
    const actual = createHash("sha256").update(raw).digest("hex");
    if (!expected || actual !== expected) throw new Error("frozen legacy L1Staking ABI hash mismatch");
    return JSON.parse(raw.toString("utf8"));
}

function addressFromSlot(raw: string): string {
    return normalizeAddress(ethers.utils.hexDataSlice(raw, 12));
}

async function collectCanonicalEvents(
    provider: ethers.providers.Provider,
    contractAddress: string,
    deployBlock: number,
    safeHead: number,
    abi: unknown[]
): Promise<{ events: ReplayEvent[]; hashes: Record<number, string>; checkpoints: RetirementManifest["checkpoints"] }> {
    const iface = new ethers.utils.Interface(abi);
    const logs = await provider.getLogs({ address: contractAddress, fromBlock: deployBlock, toBlock: safeHead });
    const blockNumbers = [...new Set([deployBlock, safeHead, ...logs.map((log) => log.blockNumber)])];
    const blocks = await Promise.all(blockNumbers.map((number) => provider.getBlock(number)));
    const hashes: Record<number, string> = {};
    const checkpoints: RetirementManifest["checkpoints"] = [];
    for (const block of blocks) {
        if (!block) throw new Error("canonical block lookup failed");
        hashes[block.number] = block.hash;
        checkpoints.push({ blockNumber: block.number, blockHash: block.hash });
    }

    const events: ReplayEvent[] = [];
    for (const log of logs) {
        try {
            const parsed = iface.parseLog(log);
            if (!["WhitelistUpdated", "Registered", "Withdrawn", "Claimed", "Slashed", "StakersRemoved"].includes(parsed.name)) {
                continue;
            }
            const args: Record<string, unknown> = {};
            for (const [key, value] of Object.entries(parsed.args)) {
                if (!/^\d+$/.test(key)) args[key] = value;
            }
            events.push({
                name: parsed.name,
                args,
                blockNumber: log.blockNumber,
                blockHash: log.blockHash,
                transactionHash: log.transactionHash,
                logIndex: log.logIndex,
            });
        } catch {
            // The proxy may emit unrelated proxy-admin events; the frozen business ABI does not parse them.
        }
    }
    return { events, hashes, checkpoints };
}

async function refreshManifest(
    provider: ethers.providers.JsonRpcProvider,
    manifest: RetirementManifest,
    abi: unknown[]
): Promise<RetirementManifest> {
    const latest = await provider.getBlockNumber();
    const safeHead = latest - manifest.retirement.confirmations;
    if (safeHead < manifest.legacy.deployBlock) throw new Error("safe head precedes legacy deployment");
    const chain = await provider.getNetwork();
    const contract = new Contract(manifest.legacy.proxy, abi, provider);
    const replay = await collectCanonicalEvents(provider, manifest.legacy.proxy, manifest.legacy.deployBlock, safeHead, abi);
    const rebuilt = rebuildLegacyState(replay.events, replay.hashes);
    const activeStakers: string[] = (await contract.getActiveStakers({ blockTag: safeHead })).map(normalizeAddress);
    const candidateMap = new Map(rebuilt.candidates.map((candidate) => [candidate.address.toLowerCase(), candidate]));
    for (const address of activeStakers) {
        if (!candidateMap.has(address.toLowerCase())) {
            candidateMap.set(address.toLowerCase(), {
                address,
                everRegistered: true,
                withdrawn: false,
                claimed: false,
                slashed: false,
                removed: false,
            });
        }
    }
    const candidates = [...candidateMap.values()].sort((a, b) => a.address.localeCompare(b.address));
    for (const candidate of candidates) {
        const [whitelist, active, withdrawal] = await Promise.all([
            contract.whitelist(candidate.address, { blockTag: safeHead }),
            contract.isActiveStaker(candidate.address, { blockTag: safeHead }),
            contract.withdrawals(candidate.address, { blockTag: safeHead }),
        ]);
        candidate.currentWhitelist = whitelist;
        candidate.currentActive = active;
        candidate.withdrawalUnlockBlock = withdrawal.toString();
    }

    const [stakingValue, slashRemaining, owner, balance] = await Promise.all([
        contract.stakingValue({ blockTag: safeHead }),
        contract.slashRemaining({ blockTag: safeHead }),
        contract.owner({ blockTag: safeHead }),
        provider.getBalance(manifest.legacy.proxy, safeHead),
    ]);
    const liabilityCount = candidates.filter(
        (candidate) => candidate.currentActive || BigNumber.from(candidate.withdrawalUnlockBlock || 0).gt(0)
    ).length;
    const expectedLiabilities = stakingValue.mul(liabilityCount).add(slashRemaining);
    const unexplainedBalance = balance.gt(expectedLiabilities) ? balance.sub(expectedLiabilities) : BigNumber.from(0);
    const safeBlock = await provider.getBlock(safeHead);
    const deployBlock = await provider.getBlock(manifest.legacy.deployBlock);
    if (!safeBlock || !deployBlock) throw new Error("snapshot block lookup failed");

    const rawAdmin = await provider.getStorageAt(manifest.legacy.proxy, ADMIN_SLOT, safeHead);
    const rawImplementation = await provider.getStorageAt(manifest.legacy.proxy, IMPLEMENTATION_SLOT, safeHead);
    manifest.legacy.proxyAdmin = addressFromSlot(rawAdmin);
    manifest.legacy.implementation = addressFromSlot(rawImplementation);
    manifest.legacy.deployBlockHash = deployBlock.hash;
    if (!manifest.retirement.stakingValue) manifest.retirement.stakingValue = stakingValue.toString();
    verifyExecutePreconditions(manifest, { stakingValue, deployBlockHash: deployBlock.hash });

    manifest.snapshot = {
        chainId: chain.chainId,
        safeHead,
        safeHeadHash: safeBlock.hash,
        stakingValue: stakingValue.toString(),
        contractBalance: balance.toString(),
        expectedLiabilities: expectedLiabilities.toString(),
        unexplainedBalance: unexplainedBalance.toString(),
        slashRemaining: slashRemaining.toString(),
        owner: normalizeAddress(owner),
        activeStakers,
    };
    manifest.candidates = candidates;
    manifest.canonicalEventIds = rebuilt.canonicalEventIds;
    manifest.checkpoints = replay.checkpoints;
    manifest.operations = await buildOperations(provider, contract, manifest, abi);
    return manifest;
}

async function buildOperations(
    provider: ethers.providers.Provider,
    contract: Contract,
    manifest: RetirementManifest,
    abi: unknown[]
): Promise<RetirementOperation[]> {
    if (!manifest.snapshot) throw new Error("snapshot missing");
    const iface = new ethers.utils.Interface(abi);
    const owner = manifest.snapshot.owner;
    const operations: RetirementOperation[] = [];
    const whitelisted = manifest.candidates.filter((candidate) => candidate.currentWhitelist).map((candidate) => candidate.address);
    if (whitelisted.length > 0) {
        operations.push({
            kind: "freeze",
            actor: owner,
            target: manifest.legacy.proxy,
            value: "0",
            data: iface.encodeFunctionData("updateWhitelist", [[], whitelisted]),
            description: `remove ${whitelisted.length} canonical ever-whitelisted address(es) from whitelist`,
            expectedState: whitelisted.map((address) => `whitelist(${address}) == false`),
        });
    }
    for (const candidate of manifest.candidates) {
        if (candidate.currentActive) {
            operations.push({
                kind: "withdraw",
                actor: candidate.address,
                target: manifest.legacy.proxy,
                value: "0",
                data: iface.encodeFunctionData("withdraw"),
                description: `withdraw legacy stake for ${candidate.address}`,
                expectedState: [
                    `isActiveStaker(${candidate.address}) == false`,
                    `withdrawals(${candidate.address}) > 0`,
                ],
            });
        } else if (
            BigNumber.from(candidate.withdrawalUnlockBlock || 0).gt(0) &&
            BigNumber.from(candidate.withdrawalUnlockBlock || 0).lte(manifest.snapshot.safeHead)
        ) {
            operations.push({
                kind: "claim",
                actor: candidate.address,
                target: manifest.legacy.proxy,
                value: "0",
                data: iface.encodeFunctionData("claimWithdrawal", [candidate.address]),
                description: `claim matured legacy withdrawal for ${candidate.address}`,
                expectedState: [`withdrawals(${candidate.address}) == 0`],
            });
        }
    }
    if (!BigNumber.from(manifest.snapshot.slashRemaining).isZero()) {
        operations.push({
            kind: "claim-slash",
            actor: owner,
            target: manifest.legacy.proxy,
            value: "0",
            data: iface.encodeFunctionData("claimSlashRemaining", [manifest.retirement.treasury]),
            description: "reconcile legacy slashRemaining to retirement treasury",
            expectedState: ["slashRemaining() == 0"],
        });
    }

    if (retirementBlockers(manifest).length === 0 && owner !== ZERO_ADDRESS) {
        operations.push({
            kind: "renounce-owner",
            actor: owner,
            target: manifest.legacy.proxy,
            value: "0",
            data: iface.encodeFunctionData("renounceOwnership"),
            description: "renounce legacy L1Staking business ownership",
            expectedState: [`owner() == ${ZERO_ADDRESS}`],
        });
        const admin = normalizeAddress(manifest.legacy.proxyAdmin || ZERO_ADDRESS);
        if (admin === ZERO_ADDRESS || admin.toLowerCase() === manifest.legacy.burnAdmin.toLowerCase()) {
            return assignOperationNonces(provider, operations);
        }
        const adminCode = await provider.getCode(admin);
        if (adminCode === "0x") {
            const proxyIface = new ethers.utils.Interface(["function changeAdmin(address newAdmin)"]);
            operations.push({
                kind: "burn-admin",
                actor: admin,
                target: manifest.legacy.proxy,
                value: "0",
                data: proxyIface.encodeFunctionData("changeAdmin", [manifest.legacy.burnAdmin]),
                description: "move the dedicated legacy proxy admin to the configured burn address",
                expectedState: [`proxy admin == ${manifest.legacy.burnAdmin}`],
            });
        } else {
            const proxyAdmin = new Contract(admin, ["function owner() view returns (address)"], provider);
            manifest.legacy.proxyAdminOwner = normalizeAddress(
                await proxyAdmin.owner({ blockTag: manifest.snapshot.safeHead })
            );
            const proxyAdminIface = new ethers.utils.Interface([
                "function changeProxyAdmin(address proxy,address newAdmin)",
            ]);
            operations.push({
                kind: "burn-admin",
                actor: manifest.legacy.proxyAdminOwner,
                target: admin,
                value: "0",
                data: proxyAdminIface.encodeFunctionData("changeProxyAdmin", [
                    manifest.legacy.proxy,
                    manifest.legacy.burnAdmin,
                ]),
                description: "detach only the legacy proxy from its confirmed-dedicated ProxyAdmin",
                expectedState: [`proxy admin == ${manifest.legacy.burnAdmin}`],
            });
        }
    }
    return assignOperationNonces(provider, operations);
}

async function assignOperationNonces(
    provider: ethers.providers.Provider,
    operations: RetirementOperation[]
): Promise<RetirementOperation[]> {
    const nextNonce = new Map<string, number>();
    for (const operation of operations) {
        const actor = operation.actor.toLowerCase();
        let nonce = nextNonce.get(actor);
        if (nonce === undefined) nonce = await provider.getTransactionCount(operation.actor, "pending");
        operation.nonce = nonce;
        nextNonce.set(actor, nonce + 1);
    }
    return operations;
}

function signerMap(provider: ethers.providers.JsonRpcProvider): Map<string, Wallet> {
    const keys: string[] = JSON.parse(process.env.RETIRE_SIGNER_KEYS || "[]");
    const result = new Map<string, Wallet>();
    for (const key of keys) {
        const wallet = new Wallet(key, provider);
        result.set(wallet.address.toLowerCase(), wallet);
    }
    return result;
}

async function executeOperations(
    provider: ethers.providers.JsonRpcProvider,
    manifest: RetirementManifest,
    manifestPath: string,
    abi: unknown[]
): Promise<void> {
    if (process.env.RETIRE_CONFIRM_EXECUTE !== "YES") {
        throw new Error("execute mode requires RETIRE_CONFIRM_EXECUTE=YES");
    }
    if (!manifest.snapshot) throw new Error("snapshot missing");
    const contract = new Contract(manifest.legacy.proxy, abi, provider);
    const deployBlock = await provider.getBlock(manifest.legacy.deployBlock);
    verifyExecutePreconditions(manifest, {
        stakingValue: await contract.stakingValue(),
        deployBlockHash: deployBlock?.hash,
    });
    if (manifest.operations.some((operation) => operation.kind === "renounce-owner" || operation.kind === "burn-admin")) {
        const [active, balance, slashRemaining] = await Promise.all([
            contract.getActiveStakers(),
            provider.getBalance(manifest.legacy.proxy),
            contract.slashRemaining(),
        ]);
        if (active.length > 0 || !balance.isZero() || !slashRemaining.isZero()) {
            throw new Error("irreversible retirement gate changed after planning");
        }
        for (const candidate of manifest.candidates) {
            const [whitelisted, withdrawal] = await Promise.all([
                contract.whitelist(candidate.address),
                contract.withdrawals(candidate.address),
            ]);
            if (whitelisted || !withdrawal.isZero()) {
                throw new Error(`irreversible retirement gate changed for ${candidate.address}`);
            }
        }
    }
    const signers = signerMap(provider);
    for (const operation of manifest.operations) {
        if (!signers.has(operation.actor.toLowerCase())) {
            throw new Error(`missing signer for ${operation.kind} actor ${operation.actor}`);
        }
        if (!Number.isSafeInteger(operation.nonce) || operation.nonce! < 0) {
            throw new Error(`missing deterministic nonce for ${operation.kind} actor ${operation.actor}`);
        }
    }
    // Resolve every actor and nonce before broadcasting so execution never partially applies a stale plan.
    const firstNonceByActor = new Map<string, number>();
    for (const operation of manifest.operations) {
        const actor = operation.actor.toLowerCase();
        if (!firstNonceByActor.has(actor)) firstNonceByActor.set(actor, operation.nonce!);
    }
    for (const [actor, nonce] of firstNonceByActor) {
        const current = await provider.getTransactionCount(actor, "pending");
        if (current !== nonce) throw new Error(`planned nonce changed for ${actor}: expected ${nonce}, got ${current}`);
    }
    for (const operation of manifest.operations) {
        const signer = signers.get(operation.actor.toLowerCase())!;
        const response = await signer.sendTransaction({
            to: operation.target,
            data: operation.data,
            value: BigNumber.from(operation.value),
            nonce: operation.nonce,
        });
        const receipt = await response.wait();
        if (receipt.status !== 1) throw new Error(`${operation.kind} transaction failed`);
        operation.txHash = receipt.transactionHash;
        operation.receiptBlockHash = receipt.blockHash;
        fs.writeFileSync(manifestPath, JSON.stringify(manifest, null, 2));
    }
}

async function dryRunOperations(
    provider: ethers.providers.JsonRpcProvider,
    manifest: RetirementManifest
): Promise<void> {
    if (!manifest.snapshot) throw new Error("snapshot missing");
    for (const operation of manifest.operations) {
        try {
            await provider.call(
                {
                    from: operation.actor,
                    to: operation.target,
                    data: operation.data,
                    value: BigNumber.from(operation.value),
                },
                manifest.snapshot.safeHead
            );
        } catch (error) {
            throw new Error(
                `dry-run simulation failed for ${operation.kind} (${operation.description}): ${
                    error instanceof Error ? error.message : String(error)
                }`
            );
        }
    }
}

async function verifyRelayEvidence(
    provider: ethers.providers.Provider,
    relays: RelayEvidence[]
): Promise<void> {
    for (const relay of relays) {
        if (!relay.transactionHash) continue;
        const receipt = await provider.getTransactionReceipt(relay.transactionHash);
        relay.success = Boolean(receipt && receipt.status === 1);
        relay.canonical = Boolean(receipt && receipt.blockHash.toLowerCase() === relay.blockHash.toLowerCase());
    }
}

function parseArgs(): { mode: RetirementMode; manifestPath: string } {
    const mode = (process.env.RETIRE_MODE || process.argv[2] || "plan") as RetirementMode;
    if (!["plan", "dry-run", "execute", "verify"].includes(mode)) throw new Error(`invalid mode ${mode}`);
    const manifestPath = path.resolve(process.env.RETIRE_MANIFEST || process.argv[3] || "legacy-l1-staking-retirement.json");
    return { mode, manifestPath };
}

async function main(): Promise<void> {
    const { mode, manifestPath } = parseArgs();
    if (!process.env.L1_RPC_URL) throw new Error("L1_RPC_URL is required");
    if (!fs.existsSync(manifestPath)) throw new Error(`retirement manifest not found: ${manifestPath}`);
    const provider = new ethers.providers.JsonRpcProvider(process.env.L1_RPC_URL);
    const manifest: RetirementManifest = JSON.parse(fs.readFileSync(manifestPath, "utf8"));
    if (
        !ethers.utils.isAddress(manifest.legacy.proxy) ||
        manifest.legacy.proxy === ZERO_ADDRESS ||
        !ethers.utils.isAddress(manifest.legacy.burnAdmin) ||
        manifest.legacy.burnAdmin === ZERO_ADDRESS ||
        !ethers.utils.isAddress(manifest.retirement.treasury)
    ) {
        throw new Error("retirement manifest contains invalid addresses");
    }
    const abi = loadFrozenAbi();
    await verifyRelayEvidence(provider, manifest.relays);
    await refreshManifest(provider, manifest, abi);

    console.log(JSON.stringify({ mode, snapshot: manifest.snapshot, operations: manifest.operations }, null, 2));
    if (mode === "dry-run") await dryRunOperations(provider, manifest);
    if (mode === "execute") await executeOperations(provider, manifest, manifestPath, abi);
    if (mode === "verify") {
        const blockers = retirementBlockers(manifest);
        if (blockers.length > 0) throw new Error(`retirement verification blocked:\n${blockers.join("\n")}`);
        manifest.verifiedAt = new Date().toISOString();
    }
    fs.writeFileSync(manifestPath, JSON.stringify(manifest, null, 2));
}

if (require.main === module) {
    main().catch((error) => {
        console.error(error);
        process.exitCode = 1;
    });
}
