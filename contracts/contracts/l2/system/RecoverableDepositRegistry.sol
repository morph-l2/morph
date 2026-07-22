// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ECDSAUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {IRecoverableDepositRegistry} from "./IRecoverableDepositRegistry.sol";

/**
 * @title RecoverableDepositRegistry
 * @notice Registry of exchange-controlled "recoverable deposit" EOAs whose
 *         whitelisted ERC-20 inflows are auto-swept to a master wallet by the
 *         Morph execution layer at transaction end (Onyx hardfork).
 * @dev Consensus surface: the EL calls {resolveSweep} through a frozen ABI and
 *      lifts {RecoverableSweepRequested} logs into sweep candidates. This
 *      contract only records authorizations and emits requests — it never moves
 *      tokens itself. See the recoverable-deposit spec S0Z9 §5–§6.
 *
 *      Double authorization (§5.2): a deposit's private key signs an EIP-712
 *      {RecoverableDepositAuthorization} (proves consent, prevents hijacking a
 *      third party's address) AND a master/operator submits it (prevents an
 *      outside user from attaching addresses to the exchange's book).
 *
 *      v1 lifecycle is one-way: register -> (poke to drain residuals) -> disable.
 *      Disabling is terminal; re-enable, master swaps and deposit-side revoke are
 *      intentionally NOT supported in v1.
 */
contract RecoverableDepositRegistry is
    IRecoverableDepositRegistry,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    EIP712Upgradeable
{
    /*//////////////////////////////////////////////////////////////
                               Constants
    //////////////////////////////////////////////////////////////*/

    /// @dev EIP-712 type hash for the deposit authorization (spec §5.3). The
    ///      field order/types are frozen and must match signer tooling exactly.
    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _AUTHORIZATION_TYPEHASH =
        keccak256(
            "RecoverableDepositAuthorization(address deposit,address master,address registry,uint256 chainId,uint256 nonce,uint64 deadline,bytes32 mode,bytes32 sweepScope)"
        );

    /// @notice Authorization mode tag. Marks the signature as a recoverable-deposit
    ///         authorization, not a generic approval or arbitrary-call grant.
    bytes32 public constant MODE = keccak256("MORPH_RECOVERABLE_DEPOSIT_V1");

    /// @notice Sweep-scope tag. Restricts the grant to `transfer(master, balance)`
    ///         of whitelisted ERC-20s only (spec §5.3).
    bytes32 public constant SWEEP_SCOPE = keccak256("WHITELISTED_ERC20_TO_MASTER_ONLY");

    /// @dev High 144 bits of the reserved system address segment (0x5300…0000).
    ///      Any address sharing this prefix (0x5300…0000 – 0x5300…FFFF) is a
    ///      system/predeploy slot and may not be a deposit or master.
    /// @dev TODO(onyx): keep this range aligned with the go-ethereum system
    ///      address definition when the geth sweep implementation lands.
    uint160 private constant _SYSTEM_SEGMENT_PREFIX = uint160(0x5300000000000000000000000000000000000000) >> 16;

    /*//////////////////////////////////////////////////////////////
                                Storage
    //////////////////////////////////////////////////////////////*/

    /// @notice Deposit address => registration record.
    mapping(address deposit => DepositRecord record) public deposits;

    /// @notice ERC-20 token => whether it is sweepable.
    mapping(address token => bool enabled) public tokenWhitelist;

    /// @notice Master => operator => whether the operator may act for the master.
    mapping(address master => mapping(address operator => bool enabled)) public operators;

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
        __EIP712_init("RecoverableDepositRegistry", "1");

        _transferOwnership(owner_);
    }

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IRecoverableDepositRegistry
    /// @dev Not gated by pause: pausing halts new registrations/pokes but must not
    ///      silently strand already-authorized deposits mid-reconciliation. Sweep
    ///      of a specific token is stopped by de-whitelisting it instead.
    function resolveSweep(address token, address deposit) external view returns (address master) {
        DepositRecord storage rec = deposits[deposit];
        if (tokenWhitelist[token] && rec.enabled && rec.master != address(0)) {
            return rec.master;
        }
        return address(0);
    }

    /// @inheritdoc IRecoverableDepositRegistry
    function getRecoverableDeposit(
        address deposit
    ) external view returns (address master, bool enabled, uint256 nonce) {
        DepositRecord storage rec = deposits[deposit];
        return (rec.master, rec.enabled, rec.nonce);
    }

    /// @notice EIP-712 domain separator for building deposit authorizations.
    // solhint-disable-next-line func-name-mixedcase
    function DOMAIN_SEPARATOR() external view returns (bytes32) {
        return _domainSeparatorV4();
    }

    /*//////////////////////////////////////////////////////////////
                          Mutating Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IRecoverableDepositRegistry
    function setRecoverableOperator(address operator, bool enabled) external {
        if (operator == address(0)) revert ZeroAddress();
        operators[msg.sender][operator] = enabled;
        emit RecoverableOperatorSet(msg.sender, operator, enabled);
    }

    /// @inheritdoc IRecoverableDepositRegistry
    function registerRecoverableDeposit(
        address deposit,
        address master,
        uint256 nonce,
        uint64 deadline,
        bytes calldata depositSignature
    ) external whenNotPaused {
        _register(deposit, master, nonce, deadline, depositSignature);
    }

    /// @inheritdoc IRecoverableDepositRegistry
    function registerRecoverableDeposits(
        RecoverableDepositRegistration[] calldata registrations
    ) external whenNotPaused nonReentrant {
        for (uint256 i = 0; i < registrations.length; i++) {
            RecoverableDepositRegistration calldata r = registrations[i];
            _register(r.deposit, r.master, r.nonce, r.deadline, r.depositSignature);
        }
    }

    /// @inheritdoc IRecoverableDepositRegistry
    function disableRecoverableDeposit(address deposit) external whenNotPaused {
        DepositRecord storage rec = deposits[deposit];
        if (msg.sender != rec.master && !operators[rec.master][msg.sender]) revert NotAuthorized();
        if (!rec.enabled) revert DepositNotActive();

        rec.enabled = false;
        emit RecoverableDepositDisabled(deposit, rec.master, msg.sender);
    }

    /// @inheritdoc IRecoverableDepositRegistry
    function pokeRecoverableSweep(address token, address deposit) external whenNotPaused {
        // Permissionless, but must not bypass the whitelist / active checks the
        // resolver enforces (§6.2). The EL still re-runs the full sweep path
        // (resolver + EOA/code recheck) after this request log is emitted.
        if (!tokenWhitelist[token]) revert TokenNotWhitelisted();
        if (!deposits[deposit].enabled) revert DepositNotActive();

        emit RecoverableSweepRequested(token, deposit);
    }

    /*//////////////////////////////////////////////////////////////
                          Restricted Functions
    //////////////////////////////////////////////////////////////*/

    /// @inheritdoc IRecoverableDepositRegistry
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

    /// @dev Shared registration path enforcing the spec §5.5 validation rules.
    function _register(
        address deposit,
        address master,
        uint256 nonce,
        uint64 deadline,
        bytes calldata depositSignature
    ) internal {
        // Master must be a valid, non-zero destination before the auth check so a
        // zero master is reported as such rather than masked by NotAuthorized.
        if (master == address(0)) revert ZeroAddress();

        // Caller must be the master or one of its authorized operators (§5.5).
        if (msg.sender != master && !operators[master][msg.sender]) revert NotAuthorized();

        // Master: not a system address, not itself an active deposit (avoids
        // recursive collection and master-wallet ambiguity, §5.5).
        if (_isSystemSegment(master)) revert SystemAddressNotAllowed();
        if (deposits[master].enabled) revert MasterIsRecoverableDeposit();

        // Deposit: non-zero, not a system address, plain EOA. A non-empty code
        // length also rejects EIP-7702 delegations, keeping error-chain recovery
        // and EOA semantics intact (§5.5). The EL rechecks this at sweep time.
        if (deposit == address(0)) revert ZeroAddress();
        if (_isSystemSegment(deposit)) revert SystemAddressNotAllowed();
        if (deposit.code.length != 0) revert DepositNotEOA();

        DepositRecord storage rec = deposits[deposit];
        // v1: a deposit may be registered exactly once. A non-zero master means it
        // is currently active or already disabled (terminal) — neither may be
        // re-registered (no re-enable / master swap in v1).
        if (rec.master != address(0)) revert DepositAlreadyRegistered();

        // Replay protection: nonce must match and the signature must be unexpired.
        if (nonce != rec.nonce) revert InvalidNonce();
        // solhint-disable-next-line not-rely-on-time
        if (block.timestamp > deadline) revert SignatureExpired();

        // EIP-712 authorization must recover to the deposit key (§5.3). `registry`
        // and `chainId` in the struct bind the signature to this contract/chain.
        bytes32 structHash = keccak256(
            abi.encode(
                _AUTHORIZATION_TYPEHASH,
                deposit,
                master,
                address(this),
                block.chainid,
                nonce,
                deadline,
                MODE,
                SWEEP_SCOPE
            )
        );
        address signer = ECDSAUpgradeable.recover(_hashTypedDataV4(structHash), depositSignature);
        if (signer != deposit) revert InvalidSignature();

        rec.master = master;
        rec.enabled = true;
        unchecked {
            rec.nonce = nonce + 1;
        }
        emit RecoverableDepositRegistered(deposit, master, msg.sender);
    }

    /// @dev True if `addr` falls in the reserved system address segment.
    function _isSystemSegment(address addr) internal pure returns (bool) {
        return (uint160(addr) >> 16) == _SYSTEM_SEGMENT_PREFIX;
    }

    /*//////////////////////////////////////////////////////////////
        v2 TODO (intentionally NOT in v1): master staking / commercial
        onboarding gate before a master may register (§5.6/§6.x), deposit
        re-enable, master swap, and deposit-side revoke. Adding any of these
        must preserve the frozen resolveSweep ABI and event topics above.
    //////////////////////////////////////////////////////////////*/

    /// @dev Reserved storage gap for future upgrades.
    uint256[50] private __gap;
}
