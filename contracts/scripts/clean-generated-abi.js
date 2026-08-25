const fs = require('fs')
const path = require('path')

const { archiveDir, verifyAbiArchive } = require('./verify-abi-archive')

verifyAbiArchive()

const abiDir = path.dirname(archiveDir)
for (const entry of fs.readdirSync(abiDir, { withFileTypes: true })) {
    if (entry.name === path.basename(archiveDir)) {
        continue
    }
    fs.rmSync(path.join(abiDir, entry.name), { recursive: true, force: true })
}

verifyAbiArchive()
