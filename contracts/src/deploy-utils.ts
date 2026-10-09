import assert from 'assert'

import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'

const fs = require("fs")

/**
 * Helper function for asserting that a contract variable is set to the expected value.
 *
 * @param contract Contract object to query.
 * @param variable Name of the variable to query.
 * @param expected Expected value of the variable.
 */
export const assertContractVariable = async (
    contract: ethers.Contract,
    variable: string,
    expected: any,
    caller?: string
) => {
    // Need to make a copy that doesn't have a signer or we get the error that contracts with
    // signers cannot override the from address.
    const temp = new ethers.Contract(
        contract.address,
        contract.interface,
        contract.provider
    )
    if (caller === null || !ethers.utils.isAddress(caller)) {
        caller = ethers.constants.AddressZero
    }
    const actual = await temp.callStatic[variable]({
        from: caller,
    })

    if (ethers.utils.isAddress(expected)) {
        assert(
            actual.toLowerCase() === expected.toLowerCase(),
            `[FATAL] ${variable} is ${actual} but should be ${expected}`
        )
        return
    }

    assert(
        actual === expected || (actual.eq && actual.eq(expected)),
        `[FATAL] ${variable} is ${actual} but should be ${expected}`
    )
}


export const assertContractVariableWithSigner = async (
    contract: ethers.Contract,
    variable: string,
    expected: any,
) => {
    // Need to make a copy that doesn't have a signer or we get the error that contracts with
    // signers cannot override the from address.
    const temp = new ethers.Contract(
        contract.address,
        contract.interface,
        contract.signer,
    )

    const actual = await temp.callStatic[variable]()

    if (ethers.utils.isAddress(expected)) {
        assert(
            actual.toLowerCase() === expected.toLowerCase(),
            `[FATAL] ${variable} is ${actual} but should be ${expected}`
        )
        return
    }

    assert(
        actual === expected || (actual.eq && actual.eq(expected)),
        `[FATAL] ${variable} is ${actual} but should be ${expected}`
    )
}

export const readDeploymentRecords = (path: string): any[] => {
    if (!fs.existsSync(path)) return [];
    const records = JSON.parse(fs.readFileSync(path, "utf8"));
    if (!Array.isArray(records)) throw new Error(`${path} must contain a deployment record array`);
    const names = new Map<string, string>();
    for (const record of records) {
        if (!record || typeof record.name !== "string" || !ethers.utils.isAddress(record.address)) {
            throw new Error(`${path} contains an invalid contract deployment record`);
        }
        if (names.has(record.name) && names.get(record.name) !== record.address.toLowerCase()) {
            throw new Error(`${path} maps ${record.name} to multiple addresses; check on-chain state and resolve the conflict`);
        }
        names.set(record.name, record.address.toLowerCase());
    }
    return records;
};

export const contractExistCheck = (path: string, contractsName: string): Boolean =>
    readDeploymentRecords(path).some(record => record.name === contractsName);

export const getContractAddressByName = (path: string, contractsName: string): string =>
    readDeploymentRecords(path).find(record => record.name === contractsName)?.address || "";

export const storage = async (
    path: string, contractsName: string, contractAddress: string, deployedBlockNumber: number,
    transactionHash?: string, pending = false, creation?: { deployer: string; nonce: number }
): Promise<string> => {
    if (!ethers.utils.isAddress(contractAddress)) throw new Error(`${contractsName} has an invalid address`);
    const records = readDeploymentRecords(path);
    const existing = records.find(record => record.name === contractsName);
    if (existing) {
        if (existing.address.toLowerCase() !== contractAddress.toLowerCase()) {
            throw new Error(`${contractsName} is already recorded as ${existing.address}; cannot replace it with ${contractAddress}`);
        }
        if (!existing.pending) return "";
        Object.assign(existing, { number: deployedBlockNumber, pending, transactionHash: transactionHash || existing.transactionHash });
    } else {
        records.push({ name: contractsName, address: contractAddress.toLowerCase(), time: new Date().toISOString(), number: deployedBlockNumber,
            ...((transactionHash || pending) ? { transactionHash, pending, ...creation } : {}) });
    }
    // Write a temporary file in the same directory and atomically replace the record file to preserve it on interruption.
    const temporary = `${path}.${process.pid}.tmp`;
    fs.writeFileSync(temporary, JSON.stringify(records, null, 2) + "\n", "utf8");
    fs.renameSync(temporary, path);
    return "";
};

export const awaitCondition = async (
    cond: () => Promise<boolean>,
    rate = 1000,
    attempts = 10
) => {
    for (let i = 0; i < attempts; i++) {
        const ok = await cond()
        if (ok) {
            return
        }
        await sleep(rate)
    }

    throw new Error('Timed out.')
}

/**
 * Basic timeout-based async sleep function.
 *
 * @param ms Number of milliseconds to sleep.
 */
export const sleep = async (ms: number): Promise<void> => {
    return new Promise<void>((resolve) => {
        setTimeout(() => {
            resolve();
        }, ms);
    });
};
