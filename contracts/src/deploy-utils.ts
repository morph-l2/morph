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

