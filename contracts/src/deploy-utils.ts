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
    expected: any
) => {
    // Need to make a copy that doesn't have a signer or we get the error that contracts with
    // signers cannot override the from address.
    const temp = new ethers.Contract(
        contract.address,
        contract.interface,
        contract.provider
    )

    const actual = await temp.callStatic[variable]({
        from: ethers.constants.AddressZero,
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

export const storge = async (
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


/**
 * Mini helper for checking if the current step is the first step in target phase.
 *
 * @param dictator SystemDictator contract.
 * @param phase Target phase.
 * @returns True if the current step is the first step in target phase.
 */
export const isStartOfPhase = async (
    dictator: ethers.Contract,
    phase: number
): Promise<boolean> => {
    const phaseToStep = {
        1: 1,
        2: 3,
        3: 6,
    }
    return (await dictator.currentStep()) === phaseToStep[phase]
}

/**
* JSON-ifies an ethers transaction object.
*
* @param tx Ethers transaction object.
* @returns JSON-ified transaction object.
*/
export const printJsonTransaction = (tx: ethers.PopulatedTransaction): void => {
    console.log(
        'JSON transaction parameters:\n' +
        JSON.stringify(
            {
                from: tx.from,
                to: tx.to,
                data: tx.data,
                value: tx.value,
                chainId: tx.chainId,
            },
            null,
            2
        )
    )
}

/**
 * Mini helper for executing a given phase.
 *
 * @param opts Options for executing the step.
 * @param opts.isLiveDeployer True if the deployer is live.
 * @param opts.SystemDictator SystemDictator contract.
 * @param opts.step Step to execute.
 * @param opts.message Message to print before executing the step.
 * @param opts.checks Checks to perform after executing the step.
 */
export const doPhase = async (opts: {
    isLiveDeployer?: boolean
    SystemDictator: ethers.Contract
    phase: number
    message: string
    checks: () => Promise<void>
}): Promise<void> => {
    const isStart = await isStartOfPhase(opts.SystemDictator, opts.phase)
    if (!isStart) {
        console.log(`Start of phase ${opts.phase} already completed`)
        return
    }

    // Extra message to help the user understand what's going on.
    console.log(opts.message)

    // Either automatically or manually execute the step.
    if (opts.isLiveDeployer) {
        console.log(`Executing phase ${opts.phase}...`)
        await opts.SystemDictator[`phase${opts.phase}`]()
    } else {
        const tx = await opts.SystemDictator.populateTransaction[
            `phase${opts.phase}`
        ]()
        console.log(`Please execute phase ${opts.phase}...`)
        console.log(`MSD address: ${opts.SystemDictator.address}`)
        printJsonTransaction(tx)
        // await printTenderlySimulationLink(opts.SystemDictator.provider, tx)
    }

    // Wait for the step to complete.
    await awaitCondition(
        async () => {
            return isStartOfPhase(opts.SystemDictator, opts.phase + 1)
        },
        3000,
        1000
    )

    // Perform post-step checks.
    await opts.checks()
}

/**
 * Check if the script should submit the transaction or wait for the deployer to do it manually.
 *
 * @returns True if the current step is the target step.
 */
export const liveDeployer = async (opts: {
    disabled: string | undefined
    controller: string
    deployer: ethers.Signer
}): Promise<boolean> => {
    if (!!opts.disabled) {
        console.log('Live deployer manually disabled')
        return false
    }
    const ret =
        (await opts.deployer.getAddress()).toLowerCase() === opts.controller.toLowerCase()
    console.log('Setting live deployer to', ret)
    return ret
}

/**
 * Mini helper for transferring a Proxy to the MSD
 *
 * @param opts Options for executing the step.
 * @param opts.isLiveDeployer True if the deployer is live.
 * @param opts.proxy proxy contract.
 * @param opts.dictator dictator contract.
 */
export const doOwnershipTransfer = async (opts: {
    isLiveDeployer?: boolean
    proxy: ethers.Contract
    name: string
    transferFunc: string
    dictator: ethers.Contract
}): Promise<void> => {
    if (opts.isLiveDeployer) {
        console.log(`Setting ${opts.name} owner to MSD`)
        await opts.proxy[opts.transferFunc](opts.dictator.address)
    } else {
        const tx = await opts.proxy.populateTransaction[opts.transferFunc](
            opts.dictator.address
        )
        console.log(`
      Please transfer ${opts.name} (proxy) owner to MSD
        - ${opts.name} address: ${opts.proxy.address}
        - MSD address: ${opts.dictator.address}
      `)
        printJsonTransaction(tx)
        printCastCommand(tx)
        //   await printTenderlySimulationLink(opts.dictator.provider, tx)
    }
}

/**
* Prints a cast commmand for submitting a given transaction.
*
* @param tx Ethers transaction object.
*/
export const printCastCommand = (tx: ethers.PopulatedTransaction): void => {
    if (process.env.CAST_COMMANDS) {
        if (!!tx.value && tx.value.gt(0)) {
            console.log(
                `cast send ${tx.to} ${tx.data} --from ${tx.from} --value ${tx.value}`
            )
        } else {
            console.log(`cast send ${tx.to} ${tx.data} --from ${tx.from} `)
        }
    }
}

/**
 * Mini helper for checking if the current step is a target step.
 *
 * @param dictator SystemDictator contract.
 * @param step Target step.
 * @returns True if the current step is the target step.
 */
export const isStep = async (
    dictator: ethers.Contract,
    step: number
): Promise<boolean> => {
    return (await dictator.currentStep()) === step
}