// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title ISweepRegistry
 * @notice Interface for the Onyx sweep auto-sweep Registry.
 * @dev The execution layer (EL) consumes this contract through a frozen surface;
 *      the resolver ABI and event topics below MUST match the EL byte-for-byte:
 *        - the EL resolves every sweep candidate through a bounded `STATICCALL`
 *          to {resolveSweep}; a zero destination means "not sweepable". The
 *          selector, exact 32-byte address encoding, and resolver semantics are
 *          the frozen consensus surface;
 *        - `SweepRequested` logs are lifted into candidates by the EL
 *          at transaction end and replayed through the same sweep path;
 *        - `Swept` and `SweepFailed` are appended BY THE EL (with the Registry
 *          as emitter) and are NEVER emitted from Solidity.
 */
interface ISweepRegistry {
    /*//////////////////////////////////////////////////////////////
                               Structs
    //////////////////////////////////////////////////////////////*/

    /// @notice Per-source registration record.
    struct SourceRecord {
        address destination; // Collection destination resolved by the EL.
        bool enabled; // Whether the source is currently sweepable (active).
        uint256 nonce; // Monotonic nonce consumed by EIP-712 authorizations.
    }

    /// @notice One entry for a batched `registerSweeps` call.
    struct SweepRegistration {
        address source;
        address destination;
        uint256 nonce;
        uint64 deadline;
        bytes sourceSignature;
    }

    /*//////////////////////////////////////////////////////////////
                                Events
    //////////////////////////////////////////////////////////////*/

    /// @notice Emitted when a destination enables or disables an operator.
    event SweepOperatorSet(address indexed destination, address indexed operator, bool enabled);

    /// @notice Emitted when a source becomes an active sweep source.
    event SweepRegistered(address indexed source, address indexed destination, address indexed operator);

    /// @notice Emitted when a source is disabled and stops being swept.
    event SweepDisabled(address indexed source, address indexed destination, address indexed operator);

    /// @notice Permissionless re-scan request. The EL lifts this into a sweep
    ///         candidate at transaction end; this contract never moves tokens.
    event SweepRequested(address indexed token, address indexed source);

    /// @notice EL-ONLY protocol settlement record. Appended by the execution layer
    ///         with this Registry as the emitter after a successful internal
    ///         `transfer` drains the source balance. NEVER emitted from Solidity.
    /// @dev `transferLogOffset` is the zero-based index of the paired
    ///      `Transfer(source, destination, amount)` within the same receipt's log
    ///      array — not a block-global RPC logIndex.
    event Swept(
        address indexed token,
        address indexed source,
        address indexed destination,
        uint256 amount,
        uint32 transferLogOffset
    );

    /// @notice EL-ONLY protocol failure record. Appended by the execution layer
    ///         with this Registry as the emitter when a candidate that resolved
    ///         to a non-zero destination did not complete its sweep. NEVER
    ///         emitted from Solidity.
    /// @dev Makes sweep failures an on-chain, indexable fact so reconciliation
    ///      does not depend on any single node's metrics.
    ///      The indexed layout deliberately mirrors {Swept} so an indexer can
    ///      use one filter for both outcomes.
    ///
    ///      Two classifications are deliberately NOT reported here:
    ///        - "resolver zero" (unregistered source or non-whitelisted token) —
    ///          it fires for every ERC-20 `Transfer` on the chain and would
    ///          bloat the log set without carrying information;
    ///        - "balance zero" — nothing to sweep is a no-op, not a failure.
    /// @param reason `keccak256` of the EL's stable snake_case classification
    ///        label, e.g. `keccak256("transfer_false")`. Hashing the stable
    ///        label rather than assigning ordinals keeps the encoding immune to
    ///        classifications being added or removed.
    event SweepFailed(
        address indexed token,
        address indexed source,
        address indexed destination,
        bytes32 reason
    );

    /// @notice Emitted when governance whitelists or de-whitelists a sweep token.
    event TokenWhitelistSet(address indexed token, bool enabled);

    /*//////////////////////////////////////////////////////////////
                                Errors
    //////////////////////////////////////////////////////////////*/

    error ZeroAddress();
    error NotAuthorized();
    error SourceNotEOA();
    error SystemAddressNotAllowed();
    /// @notice The requested destination is itself an active sweep source.
    error DestinationIsActiveSource();
    /// @notice `source` and `destination` are the same address (self-reference).
    error DestinationIsSource();
    error TokenNotWhitelisted();
    error SourceAlreadyRegistered();
    error SourceNotActive();
    error SignatureExpired();
    error InvalidNonce();
    error InvalidSignature();

    /*//////////////////////////////////////////////////////////////
                          Mutating Functions
    //////////////////////////////////////////////////////////////*/

    /// @notice Enable or disable an operator that may submit registrations and
    ///         disables on behalf of the caller (acting as a destination).
    /// @param operator The operator address.
    /// @param enabled  Whether the operator is authorized.
    function setSweepOperator(address operator, bool enabled) external;

    /// @notice Register a single source for sweeping, collecting to `destination`.
    /// @dev Caller must be `destination` or one of its authorized operators. The
    ///      `sourceSignature` must be an EIP-712 `SweepAuthorization`
    ///      signed by `source`.
    /// @param source          The plain-EOA source address (no code/delegation).
    /// @param destination     The collection destination.
    /// @param nonce           Must equal the source's current nonce.
    /// @param deadline        Signature expiry (unix seconds).
    /// @param sourceSignature The source's EIP-712 authorization signature.
    function registerSweep(
        address source,
        address destination,
        uint256 nonce,
        uint64 deadline,
        bytes calldata sourceSignature
    ) external;

    /// @notice Batch variant of {registerSweep}.
    /// @param registrations The registrations to process, in order.
    function registerSweeps(SweepRegistration[] calldata registrations) external;

    /// @notice Disable an active sweep source. Caller must be the source's
    ///         destination or one of its authorized operators.
    /// @dev A disabled source cannot be re-enabled and cannot be poked; a
    ///      destination/operator must poke to drain any residual balance before
    ///      disabling.
    /// @param source The source to disable.
    function disableSweep(address source) external;

    /// @notice Permissionless request to re-scan `source` for `token`. Emits
    ///         {SweepRequested} so the EL retries the sweep at tx end.
    /// @dev Does NOT bypass whitelist/active checks and does NOT move tokens
    ///      itself. Reverts if the pair is not currently sweepable.
    /// @param token  The whitelisted ERC-20 token.
    /// @param source The active sweep source.
    function pokeSweep(address token, address source) external;

    /// @notice Governance: add or remove `token` from the sweep whitelist.
    /// @param token   The ERC-20 token address.
    /// @param enabled Whether the token is sweepable.
    function setTokenWhitelist(address token, bool enabled) external;

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/

    /// @notice Read a source's registration record.
    /// @param source The source address.
    /// @return destination The resolved collection destination (zero if unregistered).
    /// @return enabled     Whether the source is currently active.
    /// @return nonce       The source's current authorization nonce.
    function getSweep(address source) external view returns (address destination, bool enabled, uint256 nonce);

    /// @notice Solidity mirror of the EL's sweepability check. Returns the
    ///         collection destination for a (token, source) pair, or the zero
    ///         address if not sweepable.
    /// @dev MUST return zero unless ALL hold: `token` is whitelisted, `source`
    ///      is active, and its destination is non-zero.
    ///      The EL calls this function with a fixed gas limit and requires an
    ///      exact 32-byte canonical address result. Reverting or malformed
    ///      output classifies the candidate as a resolver failure.
    /// @param token  The ERC-20 token being swept.
    /// @param source The source address receiving the inflow.
    /// @return destination The collection destination, or `address(0)` if not sweepable.
    function resolveSweep(address token, address source) external view returns (address destination);
}
