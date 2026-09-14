import { expect } from 'chai'
import * as fs from 'fs'
import * as os from 'os'
import * as path from 'path'
import { applyDeployConfigOverride } from '../src/plugin'

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
