// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ECDSAUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {ISweepRegistry} from "./ISweepRegistry.sol";

/**
 * @title SweepRegistry
 * @notice Registry of destination-controlled "sweep source" EOAs whose
 *         whitelisted ERC-20 inflows are auto-swept to a destination wallet by
 *         the Morph execution layer at transaction end (Onyx hardfork).
 * @dev Consensus surface: the EL resolves candidates through a bounded
 *      `STATICCALL` to {resolveSweep} and lifts {SweepRequested} logs into sweep
 *      candidates. This contract only records authorizations and emits requests
 *      — it never moves tokens itself.
 *
 *      Double authorization: a source's private key signs an EIP-712
 *      {SweepAuthorization} (proves consent, prevents hijacking a
 *      third party's address) AND a destination/operator submits it (prevents an
 *      outside user from attaching addresses to someone else's book).
 *
 *      The lifecycle is one-way: register -> (poke to drain residuals) -> disable.
 *      Disabling is terminal; the registry does not support re-enabling a source,
 *      changing its destination, or revocation by the source.
 */
contract SweepRegistry is
    ISweepRegistry,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    EIP712Upgradeable
{
    /*//////////////////////////////////////////////////////////////
                               Constants
    //////////////////////////////////////////////////////////////*/

    /// @dev EIP-712 type hash for the source authorization. The field order and
    ///      types are frozen and must match signer tooling exactly.
    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _AUTHORIZATION_TYPEHASH =
        keccak256(
            "SweepAuthorization(address source,address destination,address registry,uint256 chainId,uint256 nonce,uint64 deadline,bytes32 mode,bytes32 sweepScope)"
        );

    /// @notice Authorization mode tag. Marks the signature as a sweep
    ///         authorization, not a generic approval or arbitrary-call grant.
    bytes32 public constant MODE = keccak256("MORPH_SWEEP_V1");

    /// @notice Sweep-scope tag. Restricts the grant to
    ///         `transfer(destination, balance)` of whitelisted ERC-20s only
    ///         (the declared sweep scope).
    /// @dev Adding a token to the whitelist retroactively widens the scope of
    ///      every authorization already signed under this tag. The grant can only
    ///      ever move tokens to the destination the source itself designated, so
    ///      it cannot be used to steal — but whitelist additions are a governance
    ///      decision that needs a published policy.
    bytes32 public constant SWEEP_SCOPE = keccak256("WHITELISTED_ERC20_TO_DESTINATION_ONLY");

    /// @dev High 144 bits of the reserved system address segment (0x5300…0000).
    ///      Any address sharing this prefix (0x5300…0000 – 0x5300…FFFF) is a
    ///      system/predeploy slot and may not be a source or destination.
    /// @dev The system address range must remain aligned with the execution-layer
    ///      system address definition.
    uint160 private constant _SYSTEM_SEGMENT_PREFIX = uint160(0x5300000000000000000000000000000000000000) >> 16;

    /*//////////////////////////////////////////////////////////////
                                Storage
    //////////////////////////////////////////////////////////////*/

    // This contract sits behind a TransparentUpgradeableProxy, so upgrades must
    // follow the ordinary OpenZeppelin storage-compatibility rules. The EL does
    // not read these slots directly; its frozen surface is the `resolveSweep`
    // ABI and semantics below.

    /// @notice Source address => registration record.
    mapping(address source => SourceRecord record) public sources;

    /// @notice ERC-20 token => whether it is sweepable.
    mapping(address token => bool enabled) public tokenWhitelist;

    /// @notice Destination => operator => whether the operator may act for it.
    mapping(address destination => mapping(address operator => bool enabled)) public operators;

    /*//////////////////////////////////////////////////////////////
                              Initializer
    //////////////////////////////////////////////////////////////*/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize the proxy.
    /// @param owner_ Governance owner (controls whitelist and pause).
    function initialize(address owner_) external initializer {
        if (owner_ == address(0)) revert ZeroAddress();

        __Ownable_init();
        __ReentrancyGuard_init();
        __Pausable_init();
        __EIP712_init("SweepRegistry", "1");

        _transferOwnership(owner_);
    }

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc ISweepRegistry
    /// @dev Not gated by pause: pausing halts new registrations/pokes but must not
    ///      silently strand already-authorized sources mid-reconciliation. Sweep
    ///      of a specific token is stopped by de-whitelisting it instead.
    ///      The return encoding and three conditions below are consensus-facing
    ///      because every EL client calls this function during sweep preflight.
    function resolveSweep(address token, address source) external view returns (address destination) {
        SourceRecord storage rec = sources[source];
        if (tokenWhitelist[token] && rec.enabled && rec.destination != address(0)) {
            return rec.destination;
        }
        return address(0);
    }

    /// @inheritdoc ISweepRegistry
    function getSweep(
        address source
    ) external view returns (address destination, bool enabled, uint256 nonce) {
        SourceRecord storage rec = sources[source];
        return (rec.destination, rec.enabled, rec.nonce);
    }

    /// @notice EIP-712 domain separator for building source authorizations.
    // solhint-disable-next-line func-name-mixedcase
    function DOMAIN_SEPARATOR() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /*//////////////////////////////////////////////////////////////
                          Mutating Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc ISweepRegistry
    function setSweepOperator(address operator, bool enabled) external {
        if (operator == address(0)) revert ZeroAddress();
        operators[msg.sender][operator] = enabled;
        emit SweepOperatorSet(msg.sender, operator, enabled);
    }

    /// @inheritdoc ISweepRegistry
    function registerSweep(
        address source,
        address destination,
        uint256 nonce,
        uint64 deadline,
        bytes calldata sourceSignature
    ) external whenNotPaused {
        _register(source, destination, nonce, deadline, sourceSignature);
    }

    /// @inheritdoc ISweepRegistry
    function registerSweeps(
        SweepRegistration[] calldata registrations
    ) external whenNotPaused nonReentrant {
        for (uint256 i = 0; i < registrations.length; i++) {
            SweepRegistration calldata r = registrations[i];
            _register(r.source, r.destination, r.nonce, r.deadline, r.sourceSignature);
        }
    }

    /// @inheritdoc ISweepRegistry
    function disableSweep(address source) external whenNotPaused {
        SourceRecord storage rec = sources[source];
        if (msg.sender != rec.destination && !operators[rec.destination][msg.sender]) revert NotAuthorized();
        if (!rec.enabled) revert SourceNotActive();

        rec.enabled = false;
        emit SweepDisabled(source, rec.destination, msg.sender);
    }

    /// @inheritdoc ISweepRegistry
    function pokeSweep(address token, address source) external whenNotPaused {
    // Permissionless, but must not bypass the whitelist / active checks the
        // resolver enforces. The EL still re-runs the full sweep path (registry
        // read + EOA/code recheck) after this request log is emitted.
        if (!tokenWhitelist[token]) revert TokenNotWhitelisted();
        if (!sources[source].enabled) revert SourceNotActive();

        emit SweepRequested(token, source);
    }

    /*//////////////////////////////////////////////////////////////
                          Restricted Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc ISweepRegistry
    function setTokenWhitelist(address token, bool enabled) external onlyOwner {
        if (token == address(0)) revert ZeroAddress();
        tokenWhitelist[token] = enabled;
        emit TokenWhitelistSet(token, enabled);
    }

    /// @notice Pause or unpause registrations, disables and pokes.
    /// @param status Pause if true, otherwise unpause.
    function setPause(bool status) external onlyOwner {
        if (status) {
            _pause();
        } else {
            _unpause();
        }
    }

    /*//////////////////////////////////////////////////////////////
                          Internal Functions
    //////////////////////////////////////////////////////////////*/

    /// @dev Shared registration path enforcing the registry validation rules.
    function _register(
        address source,
        address destination,
        uint256 nonce,
        uint64 deadline,
        bytes calldata sourceSignature
    ) internal {
        // Destination must be a valid, non-zero address before the auth check so
        // a zero destination is reported as such rather than masked by
        // NotAuthorized.
        if (destination == address(0)) revert ZeroAddress();

        // Caller must be the destination or one of its authorized operators.
        if (msg.sender != destination && !operators[destination][msg.sender]) revert NotAuthorized();

        // Destination: not a system address, not itself an active source (avoids
        // recursive collection and destination-wallet ambiguity).
        if (_isSystemSegment(destination)) revert SystemAddressNotAllowed();
        if (sources[destination].enabled) revert DestinationIsActiveSource();

        // Source: non-zero, not a system address, plain EOA. A non-empty code
        // length also rejects EIP-7702 delegations, keeping error-chain recovery
        // and EOA semantics intact. The EL rechecks this at sweep time.
        if (source == address(0)) revert ZeroAddress();
        if (_isSystemSegment(source)) revert SystemAddressNotAllowed();
        if (source.code.length != 0) revert SourceNotEOA();

        // Reject self-reference. Sweeping an address to itself is a guaranteed
        // guaranteed no-op that still consumes a candidate slot on every inflow,
        // so it is refused at registration rather than skipped at sweep time. The
        // EL keeps its own skip guard as defence in depth.
        if (source == destination) revert DestinationIsSource();

        SourceRecord storage rec = sources[source];
        // A source may be registered exactly once. A non-zero destination
        // means it is currently active or already disabled (terminal) — neither
        // may be re-registered.
        if (rec.destination != address(0)) revert SourceAlreadyRegistered();

        // Replay protection: nonce must match and the signature must be unexpired.
        if (nonce != rec.nonce) revert InvalidNonce();
        // solhint-disable-next-line not-rely-on-time
        if (block.timestamp > deadline) revert SignatureExpired();

        // EIP-712 authorization must recover to the source key. `registry`
        // and `chainId` in the struct bind the signature to this contract/chain.
        bytes32 structHash = keccak256(
            abi.encode(
                _AUTHORIZATION_TYPEHASH,
                source,
                destination,
                address(this),
                block.chainid,
                nonce,
                deadline,
                MODE,
                SWEEP_SCOPE
            )
        );
        address signer = ECDSAUpgradeable.recover(_hashTypedDataV4(structHash), sourceSignature);
        if (signer != source) revert InvalidSignature();

        rec.destination = destination;
        rec.enabled = true;
        unchecked {
            rec.nonce = nonce + 1;
        }
        emit SweepRegistered(source, destination, msg.sender);
    }

    /// @dev True if `addr` falls in the reserved system address segment.
    function _isSystemSegment(address addr) internal pure returns (bool) {
        return (uint160(addr) >> 16) == _SYSTEM_SEGMENT_PREFIX;
    }

    /// @dev Reserved storage gap preserving proxy storage compatibility.
    uint256[50] private __gap;
}
