// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title MockSweepRegistryEL
 * @notice Minimal SweepRegistry test double for the morph-reth execution-layer tests.
 * @dev The EL resolves sweep candidates through the production
 *      {resolveSweep(address,address)} ABI. The fixture retains the production
 *      storage layout as an additional compatibility check, but the EL does not
 *      depend on those slot numbers.
 *
 *      Unlike production this exposes an unauthenticated {setSweep} so EL tests can
 *      register without building EIP-712 signatures. Everything the EL actually
 *      calls is byte-for-byte faithful; only the write path is relaxed.
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

    /// @dev Occupies slots 0..252 so the mappings below retain the production layout.
    uint256[253] private __layoutPad;

    /// @notice Source address => registration record.
    mapping(address source => SourceRecord record) public sources;

    /// @notice ERC-20 token => whether it is sweepable.
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
