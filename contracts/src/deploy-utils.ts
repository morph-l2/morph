import assert from 'assert'

import '@nomiclabs/hardhat-ethers'
import { ethers } from 'ethers'

const fs = require("fs")

// Public RPC endpoints sit behind several nodes: a read issued right after a mined write can reach a node that has not
// seen the block yet and fail (empty return data) or return stale state. Retrying the read is opt-in: with
// MORPH_READ_RETRIES unset (the default) behaviour is exactly as before, so local devnets are unaffected.
const READ_RETRIES = Number(process.env.MORPH_READ_RETRIES || 0)
const READ_RETRY_DELAY_MS = Number(process.env.MORPH_READ_RETRY_DELAY_MS || 2000)

const matchesExpected = (actual: any, expected: any): boolean => {
    if (ethers.utils.isAddress(expected)) {
        return typeof actual === 'string' && actual.toLowerCase() === expected.toLowerCase()
    }
    return actual === expected || !!(actual && actual.eq && actual.eq(expected))
}

// Calls `read` until it succeeds and (if `expected` is given) matches, at most READ_RETRIES extra times.
// The final attempt is returned/thrown as-is so the caller's own assertion reports the real problem.
const readWithRetry = async (read: () => Promise<any>, expected?: any): Promise<any> => {
    for (let attempt = 0; ; attempt++) {
        const last = attempt >= READ_RETRIES
        try {
            const actual = await read()
            if (last || expected === undefined || matchesExpected(actual, expected)) {
                return actual
            }
        } catch (e) {
            if (last) throw e
        }
        await sleep(READ_RETRY_DELAY_MS)
    }
}

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
    const actual = await readWithRetry(() => temp.callStatic[variable]({
        from: caller,
    }), expected)

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

    const actual = await readWithRetry(() => temp.callStatic[variable](), expected)

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

export const contractExistCheck = (
    path: string,
    contractsName: string
): Boolean => {
    if (fs.existsSync(path)) {
        let data = fs.readFileSync(path)
        let array = JSON.parse(data)
        array.forEach(element => {
            if (element.name = contractsName && ethers.utils.isAddress(element.address)) {
                return true
            }
        });
    }
    return false
}

export const getContractAddressByName = (
    path: string,
    contractsName: string
): string => {
    if (fs.existsSync(path)) {
        let data = fs.readFileSync(path)
        let array = JSON.parse(data)
        for (let element of array) {
            if (element.name == contractsName && ethers.utils.isAddress(element.address)) {
                return element.address
            }
        }
        return `do not find ${contractsName} in path ${path}`
    }
    return ''
}

export const storage = async (
    path: string,
    contractsName: string,
    contractAddress: string,
    deployedBlockNumber: number,
): Promise<string> => {
    let Contract = {
        name: contractsName,
        address: contractAddress,
        time: new Date().toISOString(),
        number: deployedBlockNumber
    }
    if (fs.existsSync(path)) {
        let data = fs.readFileSync(path)
        let array = JSON.parse(data)
        array.push(Contract)
        const box = JSON.stringify(array, null, 2)
        fs.writeFileSync(path, box, 'utf8', (err) => {
            console.log(err)
            return err
        })
    } else {
        var Contracts = new Array();
        Contracts[0] = Contract
        const box = JSON.stringify(Contracts, null, 2)
        fs.writeFileSync(path, box, 'utf8', (err) => {
            console.log(err)
            return err
        })
    }
    return ''
}

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

