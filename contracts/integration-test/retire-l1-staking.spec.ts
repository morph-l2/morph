import { expect } from "chai";
import { BigNumber, ethers } from "ethers";

import {
    ReplayEvent,
    RetirementManifest,
    rebuildLegacyState,
    retirementBlockers,
    verifyExecutePreconditions,
} from "../scripts/retire-l1-staking";

const alice = "0x00000000000000000000000000000000000000a1";
const bob = "0x00000000000000000000000000000000000000b2";

function event(
    blockNumber: number,
    blockHash: string,
    name: string,
    args: Record<string, unknown>,
    logIndex = 0
): ReplayEvent {
    return {
        blockNumber,
        blockHash,
        transactionHash: `0x${blockNumber.toString(16).padStart(64, "0")}`,
        logIndex,
        name,
        args,
    };
}

function manifest(): RetirementManifest {
    return {
        version: 1,
        legacy: {
            proxy: "0x0000000000000000000000000000000000000100",
            deployBlock: 1,
            deployBlockHash: "0x01",
            proxyAdminIsDedicated: false,
            burnAdmin: "0x000000000000000000000000000000000000dEaD",
        },
        retirement: { confirmations: 1, stakingValue: "100", treasury: alice },
        snapshot: {
            chainId: 1,
            safeHead: 10,
            safeHeadHash: "0x10",
            stakingValue: "100",
            contractBalance: "0",
            expectedLiabilities: "0",
            unexplainedBalance: "0",
            slashRemaining: "0",
            owner: alice,
            activeStakers: [],
        },
        candidates: [],
        canonicalEventIds: [],
        checkpoints: [],
        relays: [],
        operations: [],
    };
}

describe("legacy L1Staking retirement planning", () => {
    it("aborts when stakingValue changes during the withdrawal window", () => {
        const plan = manifest();
        expect(() => verifyExecutePreconditions(plan, { stakingValue: BigNumber.from(101) })).to.throw(
            "stakingValue changed during retirement"
        );
    });

    it("rebuilds ever-whitelisted state from the canonical branch after a short reorg", () => {
        const orphanHash = `0x${"aa".repeat(32)}`;
        const canonicalHash = `0x${"bb".repeat(32)}`;
        const events = [
            event(10, orphanHash, "WhitelistUpdated", { add: [alice], remove: [] }),
            event(10, canonicalHash, "WhitelistUpdated", { add: [bob], remove: [] }),
            event(10, canonicalHash, "Registered", { addr: bob }, 1),
            event(10, canonicalHash, "Withdrawn", { addr: bob }, 2),
            event(10, canonicalHash, "StakersRemoved", { stakers: [bob] }, 3),
        ];

        const rebuilt = rebuildLegacyState(events, { 10: canonicalHash });
        expect(rebuilt.candidates.map((candidate) => candidate.address)).to.deep.equal([ethers.utils.getAddress(bob)]);
        expect(rebuilt.candidates[0].everRegistered).to.equal(true);
        expect(rebuilt.candidates[0].withdrawn).to.equal(true);
        expect(rebuilt.candidates[0].removed).to.equal(true);
        expect(rebuilt.canonicalEventIds).to.have.length(4);
    });

    it("blocks owner/admin retirement until whitelist, withdrawals, relays, and admin scope are clean", () => {
        const plan = manifest();
        plan.candidates = [
            {
                address: alice,
                everRegistered: true,
                withdrawn: true,
                claimed: false,
                slashed: false,
                removed: true,
                currentWhitelist: true,
                currentActive: false,
                withdrawalUnlockBlock: "12",
            },
        ];
        plan.relays = [
            {
                messageHash: "0x01",
                addresses: [alice],
                transactionHash: "0x02",
                blockHash: "0x03",
                success: false,
                canonical: false,
            },
        ];
        expect(retirementBlockers(plan)).to.include.members([
            `${alice}: whitelist still enabled`,
            `${alice}: withdrawal still pending`,
            "0x01: removeStakers relay incomplete",
            "legacy proxy admin is not confirmed dedicated",
        ]);
    });
});
