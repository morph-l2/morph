// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title MockSweepRegistryEL
 * @notice Minimal SweepRegistry test double for the morph-reth execution-layer tests.
 * @dev The EL resolves sweep candidates by reading Registry storage DIRECTLY rather
 *      than calling {resolveSweep} (spec §14 item 10 ③). A test double is therefore
 *      only useful if its storage LAYOUT matches production: the padding below puts
 *      `sources` on slot 253 and `tokenWhitelist` on slot 254, and `SourceRecord`
 *      packs identically to the real one (`destination` in the low 20 bytes of
 *      `base + 0`, `enabled` at byte offset 20, `nonce` at `base + 1`).
 *
 *      Unlike production this exposes an unauthenticated {setSweep} so EL tests can
 *      register without building EIP-712 signatures. Everything the EL actually
 *      reads is byte-for-byte faithful; only the write path is relaxed.
 *
 *      Regenerate the runtime hex embedded in
 *      `morph-reth/crates/node/tests/it/sweep.rs` as `TEST_REGISTRY_RUNTIME` with:
 *
 *        forge build
 *        jq -r '.deployedBytecode.object' \
 *          forge-artifacts/MockSweepRegistryEL.sol/MockSweepRegistryEL.json
 */
contract MockSweepRegistryEL {
    struct SourceRecord {
        address destination;
        bool enabled;
        uint256 nonce;
    }

    /// @dev Occupies slots 0..252 so the mappings below land on the production slots.
    ///      Keep in sync with `SweepRegistry`; the storage-layout invariant test in
    ///      `SweepRegistry.t.sol` pins the values this must match.
    uint256[253] private __layoutPad;

    /// @notice Source address => registration record. MUST be slot 253.
    mapping(address source => SourceRecord record) public sources;

    /// @notice ERC-20 token => whether it is sweepable. MUST be slot 254.
    mapping(address token => bool enabled) public tokenWhitelist;

    event SweepRequested(address indexed token, address indexed source);

    /// @notice Test-only: whitelist `token` and point `source` at `destination`.
    /// @dev Passing `destination == address(0)` deregisters the source, which is how
    ///      tests exercise mid-transaction registration changes.
    function setSweep(address token, address source, address destination) external {
        tokenWhitelist[token] = true;
        SourceRecord storage rec = sources[source];
        rec.destination = destination;
        rec.enabled = destination != address(0);
    }

    /// @notice Emits the request log the EL lifts into a sweep candidate.
    function pokeSweep(address token, address source) external {
        emit SweepRequested(token, source);
    }

    /// @notice Mirrors the production resolver so tests can assert both agree.
    function resolveSweep(address token, address source) external view returns (address) {
        SourceRecord storage rec = sources[source];
        if (tokenWhitelist[token] && rec.enabled && rec.destination != address(0)) {
            return rec.destination;
        }
        return address(0);
    }
}
