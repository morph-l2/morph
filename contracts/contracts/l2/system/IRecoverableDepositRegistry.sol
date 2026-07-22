// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title IRecoverableDepositRegistry
 * @notice Interface for the Onyx recoverable-deposit auto-sweep Registry.
 * @dev The execution layer (EL) consumes this contract through a frozen ABI; the
 *      selectors and event topics below MUST match the EL byte-for-byte (see the
 *      Morph recoverable-deposit spec S0Z9 §5.4 and §6):
 *        - the EL performs an internal `STATICCALL resolveSweep(token, deposit)`
 *          for every sweep candidate; a zero return means "not sweepable";
 *        - `RecoverableSweepRequested` logs are lifted into candidates by the EL
 *          at transaction end and replayed through the same sweep path;
 *        - `RecoverableSweep` is appended BY THE EL (with the Registry as emitter)
 *          and is NEVER emitted from Solidity.
 */
interface IRecoverableDepositRegistry {
    /*//////////////////////////////////////////////////////////////
                               Structs
    //////////////////////////////////////////////////////////////*/

    /// @notice Per-deposit registration record.
    struct DepositRecord {
        address master; // Collection destination resolved by `resolveSweep`.
        bool enabled; // Whether the deposit is currently sweepable (active).
        uint256 nonce; // Monotonic nonce consumed by EIP-712 authorizations.
    }

    /// @notice One entry for a batched `registerRecoverableDeposits` call.
    struct RecoverableDepositRegistration {
        address deposit;
        address master;
        uint256 nonce;
        uint64 deadline;
        bytes depositSignature;
    }

    /*//////////////////////////////////////////////////////////////
                                Events
    //////////////////////////////////////////////////////////////*/

    /// @notice Emitted when a master enables or disables an operator.
    event RecoverableOperatorSet(address indexed master, address indexed operator, bool enabled);

    /// @notice Emitted when a deposit becomes an active recoverable deposit.
    event RecoverableDepositRegistered(address indexed deposit, address indexed master, address indexed operator);

    /// @notice Emitted when a deposit is disabled and stops being swept.
    event RecoverableDepositDisabled(address indexed deposit, address indexed master, address indexed operator);

    /// @notice Permissionless re-scan request. The EL lifts this into a sweep
    ///         candidate at transaction end; this contract never moves tokens.
    event RecoverableSweepRequested(address indexed token, address indexed deposit);

    /// @notice EL-ONLY protocol settlement record. Appended by the execution layer
    ///         with this Registry as the emitter after a successful internal
    ///         `transfer` drains the deposit balance. NEVER emitted from Solidity.
    /// @dev `transferLogOffset` is the zero-based index of the paired
    ///      `Transfer(deposit, master, amount)` within the same receipt's log
    ///      array — not a block-global RPC logIndex (see spec §6.3).
    event RecoverableSweep(
        address indexed token,
        address indexed deposit,
        address indexed master,
        uint256 amount,
        uint32 transferLogOffset
    );

    /// @notice Emitted when governance whitelists or de-whitelists a sweep token.
    event TokenWhitelistSet(address indexed token, bool enabled);

    /*//////////////////////////////////////////////////////////////
                                Errors
    //////////////////////////////////////////////////////////////*/

    error ZeroAddress();
    error NotAuthorized();
    error DepositNotEOA();
    error SystemAddressNotAllowed();
    error MasterIsRecoverableDeposit();
    error TokenNotWhitelisted();
    error DepositAlreadyRegistered();
    error DepositNotActive();
    error SignatureExpired();
    error InvalidNonce();
    error InvalidSignature();

    /*//////////////////////////////////////////////////////////////
                          Mutating Functions
    //////////////////////////////////////////////////////////////*/

    /// @notice Enable or disable an operator that may submit registrations and
    ///         disables on behalf of the caller (acting as a master).
    /// @param operator The operator address.
    /// @param enabled  Whether the operator is authorized.
    function setRecoverableOperator(address operator, bool enabled) external;

    /// @notice Register a single deposit as recoverable, collecting to `master`.
    /// @dev Caller must be `master` or one of its authorized operators. The
    ///      `depositSignature` must be an EIP-712 `RecoverableDepositAuthorization`
    ///      recoverable to `deposit` (see spec §5.3).
    /// @param deposit          The plain-EOA deposit address (no code/delegation).
    /// @param master           The collection destination.
    /// @param nonce            Must equal the deposit's current nonce.
    /// @param deadline         Signature expiry (unix seconds).
    /// @param depositSignature The deposit's EIP-712 authorization signature.
    function registerRecoverableDeposit(
        address deposit,
        address master,
        uint256 nonce,
        uint64 deadline,
        bytes calldata depositSignature
    ) external;

    /// @notice Batch variant of {registerRecoverableDeposit}.
    /// @param registrations The registrations to process, in order.
    function registerRecoverableDeposits(RecoverableDepositRegistration[] calldata registrations) external;

    /// @notice Disable an active recoverable deposit. Caller must be the deposit's
    ///         master or one of its authorized operators.
    /// @dev v1: a disabled deposit cannot be re-enabled and cannot be poked; a
    ///      master/operator must poke to drain any residual balance BEFORE disabling.
    /// @param deposit The deposit to disable.
    function disableRecoverableDeposit(address deposit) external;

    /// @notice Permissionless request to re-scan `deposit` for `token`. Emits
    ///         {RecoverableSweepRequested} so the EL retries the sweep at tx end.
    /// @dev Does NOT bypass whitelist/active checks and does NOT move tokens
    ///      itself (see spec §6.2). Reverts if the pair is not currently sweepable.
    /// @param token   The whitelisted ERC-20 token.
    /// @param deposit The active recoverable deposit.
    function pokeRecoverableSweep(address token, address deposit) external;

    /// @notice Governance: add or remove `token` from the sweep whitelist.
    /// @param token   The ERC-20 token address.
    /// @param enabled Whether the token is sweepable.
    function setTokenWhitelist(address token, bool enabled) external;

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/

    /// @notice Read a deposit's registration record.
    /// @param deposit The deposit address.
    /// @return master  The resolved collection destination (zero if unregistered).
    /// @return enabled Whether the deposit is currently active.
    /// @return nonce   The deposit's current authorization nonce.
    function getRecoverableDeposit(address deposit) external view returns (address master, bool enabled, uint256 nonce);

    /// @notice EL consensus entry point. Returns the collection master for a
    ///         (token, deposit) pair, or the zero address if not sweepable.
    /// @dev MUST return zero unless ALL hold: `token` is whitelisted, `deposit`
    ///      is active, and its master is non-zero (spec §5.5 / §7.5.2). Pure view
    ///      with no side effects — invoked via internal STATICCALL by the EL.
    /// @param token   The ERC-20 token being swept.
    /// @param deposit The deposit address receiving the inflow.
    /// @return master The collection destination, or `address(0)` if not sweepable.
    function resolveSweep(address token, address deposit) external view returns (address master);
}
