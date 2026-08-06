// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ECDSAUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {ISweepRegistry} from "./ISweepRegistry.sol";

/**
 * @title SweepRegistry
 * @notice Binds recoverable source EOAs to stable controllers whose single
 *         destination pointer may be changed without new source signatures.
 * @dev The contract records authorization only. The Morph execution layer calls
 *      {resolveSweep} after transactions and performs the actual token transfer.
 *
 *      A controller may delegate source registration and disabling to operators,
 *      but only the controller can move its own destination pointer. Setting the
 *      destination to zero pauses every source bound to that controller.
 */
contract SweepRegistry is ISweepRegistry, OwnableUpgradeable, PausableUpgradeable, EIP712Upgradeable {
    bytes32 private constant _AUTHORIZATION_TYPEHASH =
        keccak256("SweepAuthorization(address source,address controller,uint64 deadline)");

    mapping(address source => SourceRecord record) public sources;
    mapping(address token => bool enabled) public tokenWhitelist;
    mapping(address controller => mapping(address operator => bool enabled)) public operators;
    mapping(address controller => address destination) public override destinations;

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_) external initializer {
        if (owner_ == address(0)) revert ZeroAddress();

        __Ownable_init();
        __Pausable_init();
        __EIP712_init("SweepRegistry", "1");
        _transferOwnership(owner_);
    }

    /// @inheritdoc ISweepRegistry
    function resolveSweep(address token, address source) external view returns (address destination) {
        return _sweepDestination(token, source);
    }

    /// @inheritdoc ISweepRegistry
    function getSweep(address source) external view returns (address controller, address destination, bool enabled) {
        SourceRecord storage rec = sources[source];
        return (rec.controller, destinations[rec.controller], rec.enabled);
    }

    /// @notice EIP-712 domain separator used by source signing tools.
    // solhint-disable-next-line func-name-mixedcase
    function DOMAIN_SEPARATOR() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /// @inheritdoc ISweepRegistry
    function setSweepOperator(address operator, bool enabled) external {
        if (operator == address(0)) revert ZeroAddress();
        operators[msg.sender][operator] = enabled;
        emit SweepOperatorSet(msg.sender, operator, enabled);
    }

    /// @inheritdoc ISweepRegistry
    function setSweepDestination(address destination) external {
        address oldDestination = destinations[msg.sender];
        if (destination == oldDestination) revert DestinationUnchanged();
        destinations[msg.sender] = destination;

        emit SweepDestinationChanged(msg.sender, oldDestination, destination);
    }

    /// @inheritdoc ISweepRegistry
    function registerSweep(address source, address controller, uint64 deadline, bytes calldata sourceSignature)
        external
        whenNotPaused
    {
        _register(source, controller, deadline, sourceSignature);
    }

    /// @inheritdoc ISweepRegistry
    function registerSweeps(SweepRegistration[] calldata registrations) external whenNotPaused {
        for (uint256 i = 0; i < registrations.length; i++) {
            SweepRegistration calldata registration = registrations[i];
            _register(registration.source, registration.controller, registration.deadline, registration.sourceSignature);
        }
    }

    /// @inheritdoc ISweepRegistry
    function disableSweep(address source) external {
        SourceRecord storage rec = sources[source];
        if (!rec.enabled) revert SourceNotActive();
        // Only the source's controller (cold key) or the source itself may
        // terminally disable. Operators register sources but must NOT be able to
        // retire them: disable is terminal (a disabled source can never
        // re-register), so a leaked operator hot key must never permanently stop
        // every source under a controller.
        if (msg.sender != rec.controller && msg.sender != source) revert NotAuthorized();

        rec.enabled = false;
        emit SweepDisabled(source, rec.controller, destinations[rec.controller], msg.sender);
    }

    /// @inheritdoc ISweepRegistry
    function pokeSweep(address token, address source) external {
        // Exactly the resolver's rule, so a request can never be admitted for a
        // pair the EL will then resolve to zero. Sharing {_sweepDestination} is
        // what keeps the two from drifting: they differ only in how they report a
        // refusal — the resolver must return zero, a poke must revert.
        if (_sweepDestination(token, source) == address(0)) revert NotSweepable();

        emit SweepRequested(token, source);
    }

    /// @inheritdoc ISweepRegistry
    function setTokenWhitelist(address token, bool enabled) external onlyOwner {
        if (token == address(0)) revert ZeroAddress();
        tokenWhitelist[token] = enabled;
        emit TokenWhitelistSet(token, enabled);
    }

    /// @inheritdoc ISweepRegistry
    function setPause(bool status) external onlyOwner {
        if (status) {
            _pause();
        } else {
            _unpause();
        }
    }

    function _register(address source, address controller, uint64 deadline, bytes calldata sourceSignature) internal {
        address destination = destinations[controller];
        if (destination == address(0)) revert DestinationNotConfigured();
        _requireControllerOrOperator(controller);

        if (source == address(0)) revert ZeroAddress();
        if (source.code.length != 0) revert SourceNotEOA();

        SourceRecord storage rec = sources[source];
        // A source is bound exactly once; disableSweep keeps the binding (it is a
        // terminal state).
        if (rec.controller != address(0)) revert SourceAlreadyRegistered();

        // solhint-disable-next-line not-rely-on-time
        if (block.timestamp > deadline) revert SignatureExpired();

        bytes32 structHash = keccak256(abi.encode(_AUTHORIZATION_TYPEHASH, source, controller, deadline));
        address signer = ECDSAUpgradeable.recover(_hashTypedDataV4(structHash), sourceSignature);
        if (signer != source) revert InvalidSignature();

        rec.controller = controller;
        rec.enabled = true;
        emit SweepRegistered(source, controller, destination, msg.sender);
    }

    /// @dev The single sweepability rule, shared by {resolveSweep} and
    ///      {pokeSweep}. Returns the destination, or zero when the pair is not
    ///      currently sweepable.
    ///
    ///      Must never revert: the EL classifies a reverting resolver call as a
    ///      resolver failure, so every refusal is expressed as a zero return.
    ///
    ///      The check ORDER is consensus-facing — it sets what each candidate
    ///      costs the EL, which calls this for every ERC-20 `Transfer` on the
    ///      chain. All five terms are one conjunction, so the order is free to
    ///      pick but not free to change: it moves the per-candidate gas and
    ///      invalidates the cross-client vectors.
    ///
    ///      Cheapest-and-most-selective first: a non-whitelisted token costs one
    ///      read and never touches a per-source slot. `paused()` is deliberately
    ///      NOT near the front — it is a chain-global flag that is almost always
    ///      false, so reading it before `rec.enabled` would charge a cold SLOAD
    ///      (~2.1k) to every whitelisted-token transfer sent by an ordinary
    ///      account. Behind `rec.enabled` only real sources pay for it.
    ///
    ///      The final `sources[destination].enabled` check blocks self-references
    ///      and multi-address sweep cycles (A→B, B→A) dynamically — without a
    ///      sticky marker — by refusing to sweep INTO an address that is itself
    ///      an active source. The candidate set is frozen per transaction, so
    ///      this is what stops a permissionless `pokeSweep` from driving a
    ///      cross-transaction ping-pong that would otherwise burn subsidized
    ///      sweep gas indefinitely. Because it is dynamic, a destination that is
    ///      later disabled is immediately sweepable again.
    function _sweepDestination(address token, address source) internal view returns (address) {
        if (!tokenWhitelist[token]) return address(0);

        SourceRecord storage rec = sources[source];
        if (!rec.enabled) return address(0);
        if (paused()) return address(0);

        // Zero when the controller has no destination set, which is also how a
        // controller pauses every source bound to it.
        address destination = destinations[rec.controller];
        if (destination == address(0)) return address(0);

        // Never sweep INTO another active source: breaks self-references and
        // sweep cycles without reintroducing a permanent address ban.
        if (sources[destination].enabled) return address(0);

        return destination;
    }

    /// @dev The delegation rule: a controller always acts for itself, and so does
    ///      any operator it has enabled. Registering sources goes through this;
    ///      moving the destination pointer and disabling sources deliberately do
    ///      not, which keeps an operator key away from the funds and away from
    ///      terminal actions.
    function _requireControllerOrOperator(address controller) internal view {
        if (msg.sender != controller && !operators[controller][msg.sender]) revert NotAuthorized();
    }

    uint256[45] private __gap;
}
