/**
 * Governance deployment script for the Onyx `SweepRegistry`.
 *
 * This Registry is NOT a genesis-injected predeploy and is intentionally kept
 * out of the L1 `deploy/010-022` numbered orchestration (that flow deploys the
 * L1 system contracts). Per spec S0Z9 §5/§181 the Registry is deployed by a
 * governance transaction BEFORE the Onyx hardfork activates, and its proxy
 * address is then fixed in the execution-layer chain config
 * (`sweepRegistryAddress`) that both morph-reth and go-ethereum read.
 *
 * Usage:
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
 * OPEN COORDINATION POINT (see plan): if the EL/geth teams require a specific
 * deterministic address (e.g. 0x53..0023) rather than a plain CREATE address,
 * deploy the proxy via CREATE2 with a fixed salt, or inject it at genesis, and
 * update this script accordingly. The contract itself does not depend on its own
 * address (EIP-712 uses `address(this)`), so either path is safe.
 */
import { ethers } from "hardhat"

// L2 ProxyAdmin predeploy (Predeploys.PROXY_ADMIN).
const DEFAULT_PROXY_ADMIN = "0x530000000000000000000000000000000000000b"

// OpenZeppelin transparent proxy, fully qualified to avoid name ambiguity.
const PROXY_FQN =
    "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy"

async function main() {
    const [deployer] = await ethers.getSigners()
    const owner = ethers.utils.getAddress(process.env.REGISTRY_OWNER ?? deployer.address)
    const proxyAdmin = ethers.utils.getAddress(process.env.PROXY_ADMIN ?? DEFAULT_PROXY_ADMIN)
    const whitelistTokens = (process.env.WHITELIST_TOKENS ?? "")
        .split(",")
        .map((t) => t.trim())
        .filter((t) => t.length > 0)
        .map((t) => ethers.utils.getAddress(t))

    console.log("Deployer:    ", deployer.address)
    console.log("Owner:       ", owner)
    console.log("ProxyAdmin:  ", proxyAdmin)

    // 1. Deploy the implementation (constructor disables initializers).
    const Impl = await ethers.getContractFactory("SweepRegistry")
    const impl = await Impl.deploy()
    await impl.deployed()
    console.log("Implementation:", impl.address)

    // 2. Deploy the transparent proxy and initialize it in the same tx.
    const initData = Impl.interface.encodeFunctionData("initialize", [owner])
    const Proxy = await ethers.getContractFactory(PROXY_FQN)
    const proxy = await Proxy.deploy(impl.address, proxyAdmin, initData)
    await proxy.deployed()
    console.log("Proxy:         ", proxy.address)

    // 3. Optionally whitelist the initial sweep tokens. Only possible here when
    //    the deployer is the owner; otherwise the owner must call setTokenWhitelist.
    if (whitelistTokens.length > 0) {
        if (owner.toLowerCase() !== deployer.address.toLowerCase()) {
            console.log(
                "Skipping token whitelist: deployer is not the owner. Have the owner call setTokenWhitelist for:",
                whitelistTokens
            )
        } else {
            const registry = Impl.attach(proxy.address)
            for (const token of whitelistTokens) {
                const tx = await registry.setTokenWhitelist(token, true)
                await tx.wait()
                console.log("Whitelisted token:", token)
            }
        }
    }

    console.log("")
    console.log("Done. Next steps:")
    console.log("  - Set sweepRegistryAddress =", proxy.address, "in the")
    console.log("    morph-reth and go-ethereum chain config for the Onyx activation.")
    console.log("  - Whitelist production sweep tokens via setTokenWhitelist (owner).")
}

main().catch((err) => {
    console.error(err)
    process.exitCode = 1
})
