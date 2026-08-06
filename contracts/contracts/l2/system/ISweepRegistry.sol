// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title ISweepRegistry
 * @notice Registry for recoverable EOAs whose whitelisted ERC-20 balances are
 *         swept by the Morph execution layer.
 * @dev A source authorizes one stable controller. Each controller exposes one
 *      mutable destination, so changing that pointer updates every source bound
 *      to the controller without requiring new source signatures.
 *
 *      The EL-facing surface is frozen:
 *        - `resolveSweep(address,address)` returns an exact ABI-encoded address;
 *        - zero means the pair is not currently sweepable;
 *        - `SweepRequested` creates a sweep candidate;
 *        - `Swept` and `SweepFailed` are appended by the EL, not Solidity.
 */
interface ISweepRegistry {
    struct SourceRecord {
        address controller;
        bool enabled;
    }

    struct SweepRegistration {
        address source;
        address controller;
        uint64 deadline;
        bytes sourceSignature;
    }

    event SweepOperatorSet(address indexed controller, address indexed operator, bool enabled);

    /// @notice Emitted when a controller initializes, moves, pauses or resumes
    ///         its destination pointer. A zero destination means paused.
    event SweepDestinationChanged(
        address indexed controller, address indexed oldDestination, address indexed newDestination
    );

    event SweepRegistered(
        address indexed source, address indexed controller, address indexed destination, address operator
    );

    event SweepDisabled(
        address indexed source, address indexed controller, address indexed destination, address operator
    );

    event SweepRequested(address indexed token, address indexed source);

    /// @notice EL-only settlement record. Never emitted by Solidity.
    event Swept(
        address indexed token,
        address indexed source,
        address indexed destination,
        uint256 amount,
        uint32 transferLogOffset
    );

    /// @notice EL-only failure record. Never emitted by Solidity.
    event SweepFailed(address indexed token, address indexed source, address indexed destination, bytes32 reason);

    event TokenWhitelistSet(address indexed token, bool enabled);

    error ZeroAddress();
    error NotAuthorized();
    error SourceNotEOA();
    error DestinationNotConfigured();
    error DestinationUnchanged();
    /// @notice The (token, source) pair is not currently sweepable, so there is
    ///         nothing for the EL to retry. Call `resolveSweep` to see the same
    ///         verdict without reverting, and `getSweep` / `tokenWhitelist` /
    ///         `paused` to see which condition is the blocking one.
    error NotSweepable();
    error SourceAlreadyRegistered();
    error SourceNotActive();
    error SignatureExpired();
    error InvalidSignature();

    /// @notice Enable or disable an operator that may register sources for the
    ///         caller.
    function setSweepOperator(address operator, bool enabled) external;

    /// @notice Set the caller's sole current destination.
    /// @dev Setting zero pauses every source bound to the caller. Setting a new
    ///      non-zero address initializes, moves or resumes the pointer.
    function setSweepDestination(address destination) external;

    /// @notice Register one source under a controller.
    /// @dev The caller must be the controller or one of its operators, and the
    ///      source must sign `SweepAuthorization(address source,address
    ///      controller,uint64 deadline)`.
    function registerSweep(address source, address controller, uint64 deadline, bytes calldata sourceSignature) external;

    /// @notice Register many sources in one transaction.
    /// @dev Each entry is subject to the same rules as `registerSweep`, and the
    ///      whole batch reverts if any single entry fails. Entries may name
    ///      different controllers, provided the caller is authorized for each.
    function registerSweeps(SweepRegistration[] calldata registrations) external;

    /// @notice Terminally disable a source.
    /// @dev Only the source's controller or the source itself may disable.
    ///      Operators register sources but cannot terminally disable them: a
    ///      disabled source can never re-register, so terminal power stays with
    ///      the cold key / the source owner.
    function disableSweep(address source) external;

    /// @notice Permissionlessly ask the EL to retry one currently sweepable pair.
    /// @dev Enforces exactly the `resolveSweep` rule and reverts with
    ///      {NotSweepable} otherwise, so a request is never admitted for a pair
    ///      the EL would resolve to zero — including while the registry is
    ///      paused. Those requests would consume the block's sweep preflight
    ///      budget and settle nothing.
    function pokeSweep(address token, address source) external;

    function setTokenWhitelist(address token, bool enabled) external;

    /// @notice Pause or resume the whole registry.
    /// @dev While paused `resolveSweep` returns zero for every pair, `pokeSweep`
    ///      reverts, and new registrations are rejected — so this is the
    ///      chain-wide sweep kill switch, not just an admission valve. Existing
    ///      bindings and destination pointers are preserved and resume on
    ///      unpause; controllers may still move their own pointer while paused.
    function setPause(bool status) external;

    /// @notice Read a source's binding: its controller, that controller's current
    ///         destination, and whether the source is still active.
    /// @dev Ungated — unlike `resolveSweep` this ignores the token whitelist and
    ///      the pause flag, so a non-zero destination here does not imply the
    ///      pair is currently sweepable.
    function getSweep(address source) external view returns (address controller, address destination, bool enabled);

    function destinations(address controller) external view returns (address destination);

    function resolveSweep(address token, address source) external view returns (address destination);
}
