/**
 * Deterministic CREATE2 deployment script for the Onyx `SweepRegistry`.
 *
 * Both the implementation and the transparent proxy are compiled with the
 * FOUNDRY toolchain (`forge build`, foundry.toml optimizer_runs=999999) and
 * deployed via the Arachnid deterministic-deployment-proxy
 * (`0x4e59b44847b379578588920cA78FbF26c0B4956C`), which already exists on Morph
 * Mainnet and Hoodi. Using fixed salts makes the proxy address predictable and
 * identical across networks, so it matches the hardcoded consensus constant
 * every EL client pins (morph-reth `SWEEP_REGISTRY_ADDRESS`; go-ethereum
 * likewise per the Onyx spec). The address is NOT read from chain config —
 * this script asserts the deployed proxy equals that constant.
 *
 * WHY FORGE, NOT HARDHAT: the hardhat optimizer runs setting is deliberately
 * left at the repo default (10_000) so the legacy L1 contracts (e.g. Rollup,
 * ~25.8k runtime bytes) stay under the EIP-170 24,576-byte deploy limit. The
 * new Onyx contracts need optimizer_runs=999999 to byte-for-byte match the
 * bytecode the EL clients and the devnet scripts (onyx-sweep-common.sh) lock
 * onto for the deterministic `SWEEP_REGISTRY_ADDRESS`. The two toolchains must
 * never be mixed for the same contract — a different optimizer setting would
 * change the CREATE2 address and break the hardcoded consensus constant.
 *
 * Proxy initialization is a separate transaction — the proxy is deployed
 * with empty init-data so that the owner (which may differ between
 * mainnet and testnet) does not affect the deterministic address.
 *
 * Usage:
 *   forge build
 *   REGISTRY_OWNER=0x..  PROXY_ADMIN=0x..  WHITELIST_TOKENS=0x..,0x.. \
 *     npx hardhat run scripts/deploy-sweep-registry.ts --network <net>
 *
 * Environment:
 *   REGISTRY_OWNER    Governance owner (whitelist + pause). Defaults to deployer.
 *   PROXY_ADMIN       ProxyAdmin address that will own the transparent proxy.
 *                     Defaults to the L2 ProxyAdmin predeploy 0x53..000b.
 *   WHITELIST_TOKENS  Optional comma-separated ERC-20s to whitelist immediately
 *                     (only applied when the deployer is the owner).
 *
 * Precomputing the address before deployment with cast:
 *   FACTORY=0x4e59b44847b379578588920cA78FbF26c0B4956C
 *   SALT_IMPL=$(cast keccak "morph.sweep-registry.impl.v1")
 *   SALT_PROXY=$(cast keccak "morph.sweep-registry.proxy.v1")
 *   IMPL_INITCODE=$(cat forge-artifacts/SweepRegistry.sol/SweepRegistry.json | jq -r .bytecode.object)
 *   IMPL=$(cast create2 --deployer $FACTORY --salt $SALT_IMPL --init-code $IMPL_INITCODE)
 *   PROXY_ARGS=$(cast abi-encode 'x(address,address,bytes)' "$IMPL" "$PROXY_ADMIN" "")
 *   PROXY_BYTECODE=$(cat forge-artifacts/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json | jq -r .bytecode.object)
 *   PROXY_INITCODE=${PROXY_BYTECODE}${PROXY_ARGS#0x}
 *   cast create2 --deployer $FACTORY --salt $SALT_PROXY --init-code $PROXY_INITCODE
 */
import fs from "fs"
import path from "path"
import { ethers } from "hardhat"

// L2 ProxyAdmin predeploy (Predeploys.PROXY_ADMIN).
const DEFAULT_PROXY_ADMIN = "0x530000000000000000000000000000000000000b"

// Arachnid deterministic-deployment-proxy. This address is identical on every
// EVM chain where anyone has broadcast the one-time self-bootstrapping
// transaction. Confirmed present on Morph Mainnet (2818) and Hoodi (2910).
const CREATE2_FACTORY = "0x4e59b44847b379578588920cA78FbF26c0B4956C"

// The factory is raw EVM bytecode, NOT a Solidity contract — it has no ABI and no
// deploy() function. It reads calldata[0:32] as the CREATE2 salt and CREATE2s
// calldata[32:] as the initcode, which is exactly what getCreate2Address() above
// predicts. Encoding a "deploy(bytes,bytes32)" call instead prepends a 4-byte selector
// plus ABI headers and shifts the payload, so the factory reads a garbage salt and an
// initcode beginning with 0x00 (STOP) — deploying an EMPTY contract at an address
// unrelated to EXPECTED_REGISTRY.
const factoryCalldata = (salt: string, initcode: string): string =>
    ethers.utils.hexConcat([salt, initcode])

// Fixed, versioned salts so the address is predictable and stable forever.
const SALT_IMPL   = ethers.utils.keccak256(
    ethers.utils.toUtf8Bytes("morph.sweep-registry.impl.v1"),
)
const SALT_PROXY  = ethers.utils.keccak256(
    ethers.utils.toUtf8Bytes("morph.sweep-registry.proxy.v1"),
)

// Deterministic proxy address when proxyAdmin == DEFAULT_PROXY_ADMIN. This MUST
// equal the morph-reth SWEEP_REGISTRY_ADDRESS constant (crates/chainspec/src/
// constants.rs); otherwise the execution layer will never find the registry.
// Only enforced for the default admin — a custom PROXY_ADMIN legitimately changes it.
// Re-derived for the controller model with the sticky everDestination flag:
// any SweepRegistry bytecode change invalidates this address (impl → proxy
// CREATE2 chain) and requires re-syncing constants.rs, onyx-sweep-common.sh and
// the morph-reth EL test assets (Onyx spec §3.2).
const EXPECTED_REGISTRY = "0x0fF2Ea62eBca29E70aE2b0551a54eFFa4ea7DeEa"

/**
 * Reads a contract's creation initcode from its FORGE artifact.
 *
 * The SweepRegistry and its transparent proxy are compiled and deployed with the
 * foundry toolchain (`forge build`, foundry.toml optimizer_runs=999999), NOT with
 * hardhat's optimizer. This is deliberate: the hardhat optimizer runs setting is
 * left at the repo default (10_000) so the legacy L1 contracts stay deployable,
 * while the new Onyx contracts keep the runs=999999 bytecode that the EL clients
 * and the devnet scripts (onyx-sweep-common.sh) lock onto for the deterministic
 * `SWEEP_REGISTRY_ADDRESS`. Mixing the two toolchains would produce a different
 * CREATE2 address and break the hardcoded consensus constant.
 */
function readForgeInitcode(artifactPath: string): string {
    const json = JSON.parse(fs.readFileSync(artifactPath, "utf8"))
    const initcode = json?.bytecode?.object
    if (typeof initcode !== "string" || !initcode.startsWith("0x")) {
        throw new Error(
            `Forge artifact ${artifactPath} has no creation bytecode; run 'forge build' first.`,
        )
    }
    return initcode
}

const IMPL_ARTIFACT = path.join(__dirname, "../forge-artifacts/SweepRegistry.sol/SweepRegistry.json")
const PROXY_ARTIFACT =
    path.join(__dirname, "../forge-artifacts/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json")

async function main() {
    const [deployer] = await ethers.getSigners()
    const owner = ethers.utils.getAddress(
        process.env.REGISTRY_OWNER ?? deployer.address,
    )
    const proxyAdmin = ethers.utils.getAddress(
        process.env.PROXY_ADMIN ?? DEFAULT_PROXY_ADMIN,
    )
    const whitelistTokens = (process.env.WHITELIST_TOKENS ?? "")
        .split(",")
        .map((t) => t.trim())
        .filter((t) => t.length > 0)
        .map((t) => ethers.utils.getAddress(t))

    console.log("Deployer:        ", deployer.address)
    console.log("Owner:           ", owner)
    console.log("ProxyAdmin:      ", proxyAdmin)
    console.log("CREATE2 factory: ", CREATE2_FACTORY)
    console.log("")

    // ---- precompute deterministic addresses ---------------------------------
    type AddrPair = { impl: string; proxy: string }

    const precomputeAsync = async (): Promise<AddrPair> => {
        const implInitcode = readForgeInitcode(IMPL_ARTIFACT)
        const proxyInitcode = readForgeInitcode(PROXY_ARTIFACT)

        const implAddr = ethers.utils.getCreate2Address(
            CREATE2_FACTORY,
            SALT_IMPL,
            ethers.utils.keccak256(implInitcode),
        )

        const proxyConstructorArgs = ethers.utils.defaultAbiCoder.encode(
            ["address", "address", "bytes"],
            [implAddr, proxyAdmin, "0x"],
        )
        const proxyFullInitcode = ethers.utils.solidityPack(
            ["bytes", "bytes"],
            [proxyInitcode, proxyConstructorArgs],
        )
        const proxyAddr = ethers.utils.getCreate2Address(
            CREATE2_FACTORY,
            SALT_PROXY,
            ethers.utils.keccak256(proxyFullInitcode),
        )
        return { impl: implAddr, proxy: proxyAddr }
    }

    const addrs = await precomputeAsync()
    console.log("Impl (predicted):  ", addrs.impl)
    console.log("Proxy (predicted): ", addrs.proxy)
    console.log("saltImpl:          ", SALT_IMPL)
    console.log("saltProxy:         ", SALT_PROXY)
    console.log("")

    // Guard cross-network consistency: with the default ProxyAdmin the proxy must
    // land on the address morph-reth hardcodes, or the EL will never find it.
    if (
        proxyAdmin.toLowerCase() === DEFAULT_PROXY_ADMIN.toLowerCase() &&
        addrs.proxy.toLowerCase() !== EXPECTED_REGISTRY.toLowerCase()
    ) {
        throw new Error(
            `Predicted proxy ${addrs.proxy} != expected ${EXPECTED_REGISTRY}. ` +
            "morph-reth hardcodes the expected address (SWEEP_REGISTRY_ADDRESS); a " +
            "mismatch means the EL will never find the registry. Check the SweepRegistry " +
            "bytecode / solc version / OZ version / salts against the locked deployment.",
        )
    }

    // ---- verify factory exists ---------------------------------------------
    const factoryCode = await ethers.provider.getCode(CREATE2_FACTORY)
    if (factoryCode === "0x") {
        throw new Error(
            `CREATE2 factory ${CREATE2_FACTORY} does not exist on this network. ` +
            "Deploy it first via the deterministic-deployment-proxy one-shot transaction.",
        )
    }
    console.log(`Factory code: ${factoryCode.length - 2} bytes (confirmed present)`)

    // ---- 1. Deploy the implementation via CREATE2 ---------------------------
    {
        const onChainCode = await ethers.provider.getCode(addrs.impl)
        if (onChainCode !== "0x") {
            console.log(`Impl already deployed at ${addrs.impl} — skipping`)
        } else {
            const implInitcode = readForgeInitcode(IMPL_ARTIFACT)
            console.log(`Deploying impl via CREATE2 …`)
            const tx = await deployer.sendTransaction({
                to: CREATE2_FACTORY,
                data: factoryCalldata(SALT_IMPL, implInitcode),
            })
            console.log(`  tx: ${tx.hash}`)
            await tx.wait()
            console.log(`  deployed: ${addrs.impl}`)
        }
    }

    // ---- 2. Deploy the transparent proxy via CREATE2 (no init data) --------
    {
        const onChainCode = await ethers.provider.getCode(addrs.proxy)
        if (onChainCode !== "0x") {
            console.log(`Proxy already deployed at ${addrs.proxy} — skipping`)
        } else {
            const proxyInitcode = readForgeInitcode(PROXY_ARTIFACT)
            const proxyConstructorArgs = ethers.utils.defaultAbiCoder.encode(
                ["address", "address", "bytes"],
                [addrs.impl, proxyAdmin, "0x"],
            )
            const proxyFullInitcode = ethers.utils.solidityPack(
                ["bytes", "bytes"],
                [proxyInitcode, proxyConstructorArgs],
            )
            console.log(`Deploying proxy via CREATE2 …`)
            const tx = await deployer.sendTransaction({
                to: CREATE2_FACTORY,
                data: factoryCalldata(SALT_PROXY, proxyFullInitcode),
            })
            console.log(`  tx: ${tx.hash}`)
            await tx.wait()
            console.log(`  deployed: ${addrs.proxy}`)
        }
    }

    // ---- 3. Initialize the proxy (separate tx, after deployment) -----------
    const Impl = await ethers.getContractFactory("SweepRegistry")
    const registry = Impl.attach(addrs.proxy)

    // Check if already initialized (owner() call will revert if not).
    let needsInit = true
    try {
        const currentOwner = await registry.owner()
        needsInit = currentOwner === ethers.constants.AddressZero
        console.log(`Proxy already initialized, owner = ${currentOwner}`)
    } catch {
        // owner() reverted — not yet initialized
    }

    if (needsInit) {
        console.log(`Initializing proxy (owner = ${owner}) …`)
        const tx = await registry.initialize(owner)
        console.log(`  tx: ${tx.hash}`)
        await tx.wait()
        console.log(`  initialized`)
    }

    // ---- 4. Optionally whitelist initial tokens ----------------------------
    if (whitelistTokens.length > 0) {
        if (owner.toLowerCase() !== deployer.address.toLowerCase()) {
            console.log(
                "Skipping token whitelist: deployer is not the owner. " +
                "Have the owner call setTokenWhitelist for:",
                whitelistTokens,
            )
        } else {
            for (const token of whitelistTokens) {
                const tx = await registry.setTokenWhitelist(token, true)
                await tx.wait()
                console.log("Whitelisted token:", token)
            }
        }
    }

    console.log("")
    console.log("Done. Next steps:")
    console.log("  - Registry address", addrs.proxy, "matches the hardcoded")
    console.log("    consensus constant on every EL client (not set in chain config).")
    console.log("  - Proxy admin is", proxyAdmin, "(predeploy ProxyAdmin).")
    console.log("  - SweepRegistry owner is", owner, "(governs whitelist / pause).")
    console.log("  - The proxy address is deterministic — " +
                "it will be identical on every network that shares the same impl bytecode.")
}

main().catch((err) => {
    console.error(err)
    process.exitCode = 1
})
