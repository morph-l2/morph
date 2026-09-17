import { expect } from 'chai'
import * as fs from 'fs'
import * as os from 'os'
import * as path from 'path'
import { applyDeployConfigOverride } from '../src/plugin'
import { ethers } from 'ethers'
import { readBatchParameters } from '../tasks/check'
import { predeploys } from '../src/constants'

describe('deployment configuration overrides', () => {
    let directory: string
    beforeEach(() => { directory = fs.mkdtempSync(path.join(os.tmpdir(), 'morph-deploy-config-')) })
    afterEach(() => {
        for (const filename of fs.readdirSync(directory)) fs.unlinkSync(path.join(directory, filename))
        fs.rmdirSync(directory)
    })

    function write(value: unknown): string {
        const filename = path.join(directory, 'config.json')
        fs.writeFileSync(filename, JSON.stringify(value))
        return filename
    }

    it('merges the generated batch header without modifying the original network configuration', () => {
        const original = { batchHeader: '', l2ChainID: 53077 }
        const result = applyDeployConfigOverride(original, write({ batchHeader: '0x1234' }))
        expect(result).to.deep.equal({ batchHeader: '0x1234', l2ChainID: 53077 })
        expect(original.batchHeader).to.equal('')
    })

    it('rejects misspelled fields and non-object inputs', () => {
        expect(() => applyDeployConfigOverride({ batchHeader: '' }, write({ batchheader: '0x1234' }))).to.throw('unknown field')
        for (const value of [null, [], '0x1234']) {
            expect(() => applyDeployConfigOverride({ batchHeader: '' }, write(value))).to.throw('JSON object')
        }
    })

    it('rejects fields that can override object methods', () => {
        const filename = path.join(directory, 'config.json')
        fs.writeFileSync(filename, '{"__proto__": {"batchHeader": "0x1234"}}')
        expect(() => applyDeployConfigOverride({ batchHeader: '' }, filename)).to.throw('unknown field')
    })
})

describe('batch parameter snapshots', () => {
    const abi = new ethers.utils.Interface([
        'function batchBlockInterval() view returns (uint256)',
        'function batchTimeout() view returns (uint256)',
    ])
    const hash = '0x' + '11'.repeat(32)

    function provider(interval: string, timeout: string, changed = false) {
        const calls: string[] = []
        let reads = 0
        return {
            calls,
            rpc: {
                getBlock: async (tag: string) => {
                    expect(tag).to.equal('0x7b')
                    return { number: 123, hash: changed && reads++ > 0 ? '0x' + '22'.repeat(32) : hash }
                },
                call: async (transaction: ethers.providers.TransactionRequest, tag: string) => {
                    expect(tag).to.equal('0x7b')
                    expect(transaction.to).to.equal(predeploys.Gov)
                    const name = abi.parseTransaction({ data: transaction.data as string }).name
                    calls.push(name)
                    return abi.encodeFunctionResult(name, [name === 'batchBlockInterval' ? interval : timeout])
                },
                getNetwork: async () => ({ chainId: 53077 }),
            } as unknown as ethers.providers.Provider,
        }
    }

    async function rejected(action: Promise<unknown>, message: string) {
        try { await action } catch (error) {
            expect((error as Error).message).to.contain(message)
            return
        }
        throw new Error('Expected snapshot rejection')
    }

    it('reads both values at the same block and records its identity without uint64 precision loss', async () => {
        const mock = provider('0', '18446744073709551615')
        expect(await readBatchParameters(mock.rpc, '123')).to.deep.equal({
            l2ChainId: 53077, l2BlockNumber: 123, l2BlockHash: hash,
            batchBlockInterval: '0', batchTimeout: '18446744073709551615',
        })
        expect(mock.calls).to.have.members(['batchBlockInterval', 'batchTimeout'])
    })

    it('rejects a changed block, out-of-range values and disabled triggers', async () => {
        await rejected(readBatchParameters(provider('1', '2', true).rpc, '123'), 'block changed')
        await rejected(readBatchParameters(provider('18446744073709551616', '2').rpc, '123'), 'exceeds uint64')
        await rejected(readBatchParameters(provider('0', '0').rpc, '123'), 'Both batch triggers')
    })

    it('requires an explicit decimal height before making RPC calls', async () => {
        for (const value of ['latest', '-1', '1.5', '0x7b', '']) {
            const mock = provider('1', '2')
            await rejected(readBatchParameters(mock.rpc, value), 'block-number')
            expect(mock.calls).to.deep.equal([])
        }
    })
})
