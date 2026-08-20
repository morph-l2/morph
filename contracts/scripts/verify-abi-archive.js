const crypto = require('crypto')
const fs = require('fs')
const path = require('path')

const archiveDir = path.resolve(__dirname, '..', 'abi', 'archive')
const manifestPath = path.join(archiveDir, 'manifest.json')

function verifyAbiArchive() {
    if (!fs.existsSync(manifestPath)) {
        throw new Error(`frozen ABI manifest is missing: ${manifestPath}`)
    }

    const manifest = JSON.parse(fs.readFileSync(manifestPath, 'utf8'))
    for (const [relativePath, expectedHash] of Object.entries(manifest.artifacts)) {
        const artifactPath = path.resolve(archiveDir, relativePath)
        if (!fs.existsSync(artifactPath)) {
            throw new Error(`frozen ABI artifact is missing: ${artifactPath}`)
        }
        const actualHash = crypto
            .createHash('sha256')
            .update(fs.readFileSync(artifactPath))
            .digest('hex')
        if (actualHash !== expectedHash) {
            throw new Error(
                `frozen ABI hash mismatch for ${relativePath}: expected ${expectedHash}, got ${actualHash}`
            )
        }
    }
}

if (require.main === module) {
    verifyAbiArchive()
    console.log('Frozen ABI archive hashes verified')
}

module.exports = { archiveDir, verifyAbiArchive }
