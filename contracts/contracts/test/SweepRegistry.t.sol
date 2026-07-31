// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";

import {SweepRegistry} from "../l2/system/SweepRegistry.sol";
import {ISweepRegistry} from "../l2/system/ISweepRegistry.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract SweepRegistryTest is Test {
    SweepRegistry internal registry;
    SweepRegistry internal registryImpl;
    ProxyAdmin internal proxyAdmin;

    address internal multisig = address(512);
    address internal owner = address(64);
    address internal operator = address(0xACE);
    address internal stranger = address(0xBAD);

    // Source key pair (deterministic for signing).
    uint256 internal sourcePk = 0xA11CE;
    address internal source;

    address internal destination = address(0x11115);
    address internal token = address(0xDEADBEEF);

    // Storage slots the execution layer reads directly (spec §14 item 10 ③).
    // Asserted against the live layout in test_storage_layout_frozen_for_el.
    uint256 internal constant SOURCES_SLOT = 253;
    uint256 internal constant TOKEN_WHITELIST_SLOT = 254;

    // Must match the contract's frozen tags (asserted in a dedicated test).
    bytes32 internal constant AUTH_TYPEHASH =
        keccak256(
            "SweepAuthorization(address source,address destination,address registry,uint256 chainId,uint256 nonce,uint64 deadline,bytes32 mode,bytes32 sweepScope)"
        );
    bytes32 internal constant MODE = keccak256("MORPH_SWEEP_V1");
    bytes32 internal constant SWEEP_SCOPE = keccak256("WHITELISTED_ERC20_TO_DESTINATION_ONLY");

    function setUp() public {
        source = vm.addr(sourcePk);

        vm.prank(multisig);
        proxyAdmin = new ProxyAdmin();

        registryImpl = new SweepRegistry();

        vm.prank(multisig);
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(registryImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(SweepRegistry.initialize.selector, owner)
        );
        registry = SweepRegistry(address(proxy));

        // Whitelist the sweep token by default.
        vm.prank(owner);
        registry.setTokenWhitelist(token, true);

        vm.label(address(registry), "SweepRegistry");
        vm.label(source, "source");
        vm.label(destination, "destination");
        vm.label(operator, "operator");
    }

    /*//////////////////////////////////////////////////////////////
                                Helpers
    //////////////////////////////////////////////////////////////*/

    function _sign(
        uint256 pk,
        address source_,
        address destination_,
        uint256 nonce_,
        uint64 deadline_
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(
                AUTH_TYPEHASH,
                source_,
                destination_,
                address(registry),
                block.chainid,
                nonce_,
                deadline_,
                MODE,
                SWEEP_SCOPE
            )
        );
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", registry.DOMAIN_SEPARATOR(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    function _deadline() internal view returns (uint64) {
        return uint64(block.timestamp + 1 hours);
    }

    function _register(address caller) internal {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, destination, 0, deadline);
        vm.prank(caller);
        registry.registerSweep(source, destination, 0, deadline, sig);
    }

    /// @dev Reimplements the execution layer's resolver exactly as it will run:
    ///      two raw storage reads, no contract call. Used to prove the EL's slot
    ///      read and {ISweepRegistry.resolveSweep} cannot drift apart.
    function _resolveViaSlots(address token_, address source_) internal view returns (address) {
        bytes32 recordBase = keccak256(abi.encode(source_, SOURCES_SLOT));
        uint256 word0 = uint256(vm.load(address(registry), recordBase));
        address dest = address(uint160(word0));
        bool enabled = ((word0 >> 160) & 0xFF) != 0;

        bytes32 whitelistSlot = keccak256(abi.encode(token_, TOKEN_WHITELIST_SLOT));
        bool listed = uint256(vm.load(address(registry), whitelistSlot)) != 0;

        if (listed && enabled && dest != address(0)) return dest;
        return address(0);
    }

    /*//////////////////////////////////////////////////////////////
                            Initialization
    //////////////////////////////////////////////////////////////*/

    function test_initialize_succeeds() public {
        assertEq(registry.owner(), owner);
    }

    function test_initialize_reverts_on_impl() public {
        SweepRegistry impl = new SweepRegistry();
        vm.expectRevert();
        impl.initialize(owner);
    }

    function test_frozen_tags_match_spec() public {
        assertEq(registry.MODE(), MODE);
        assertEq(registry.SWEEP_SCOPE(), SWEEP_SCOPE);
    }

    /*//////////////////////////////////////////////////////////////
                          Whitelist / Operator
    //////////////////////////////////////////////////////////////*/

    function test_setTokenWhitelist_onlyOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(stranger);
        registry.setTokenWhitelist(token, false);
    }

    function test_setTokenWhitelist_reverts_zero() public {
        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(owner);
        registry.setTokenWhitelist(address(0), true);
    }

    function test_setSweepOperator_sets_and_emits() public {
        vm.expectEmit(true, true, false, true, address(registry));
        emit ISweepRegistry.SweepOperatorSet(destination, operator, true);
        vm.prank(destination);
        registry.setSweepOperator(operator, true);
        assertTrue(registry.operators(destination, operator));
    }

    function test_setSweepOperator_reverts_zero() public {
        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(destination);
        registry.setSweepOperator(address(0), true);
    }

    /*//////////////////////////////////////////////////////////////
                              Register
    //////////////////////////////////////////////////////////////*/

    function test_register_succeeds_and_resolves() public {
        vm.expectEmit(true, true, true, false, address(registry));
        emit ISweepRegistry.SweepRegistered(source, destination, destination);
        _register(destination);

        (address d, bool enabled, uint256 nonce) = registry.getSweep(source);
        assertEq(d, destination);
        assertTrue(enabled);
        assertEq(nonce, 1);
        assertEq(registry.resolveSweep(token, source), destination);
    }

    function test_register_via_operator() public {
        vm.prank(destination);
        registry.setSweepOperator(operator, true);

        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, destination, 0, deadline);
        vm.prank(operator);
        registry.registerSweep(source, destination, 0, deadline, sig);

        assertEq(registry.resolveSweep(token, source), destination);
    }

    function test_register_reverts_not_authorized() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, destination, 0, deadline);
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(stranger);
        registry.registerSweep(source, destination, 0, deadline, sig);
    }

    function test_register_reverts_bad_signature() public {
        uint64 deadline = _deadline();
        // Sign with the wrong key -> recovers to a different address.
        bytes memory sig = _sign(0xB0B, source, destination, 0, deadline);
        vm.expectRevert(ISweepRegistry.InvalidSignature.selector);
        vm.prank(destination);
        registry.registerSweep(source, destination, 0, deadline, sig);
    }

    function test_register_reverts_expired() public {
        uint64 deadline = uint64(block.timestamp);
        bytes memory sig = _sign(sourcePk, source, destination, 0, deadline);
        vm.warp(block.timestamp + 2);
        vm.expectRevert(ISweepRegistry.SignatureExpired.selector);
        vm.prank(destination);
        registry.registerSweep(source, destination, 0, deadline, sig);
    }

    function test_register_reverts_bad_nonce() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, destination, 1, deadline);
        vm.expectRevert(ISweepRegistry.InvalidNonce.selector);
        vm.prank(destination);
        registry.registerSweep(source, destination, 1, deadline, sig);
    }

    function test_register_reverts_source_has_code() public {
        address codeSource = address(0xC0DE);
        vm.etch(codeSource, hex"600060005360016000f3");
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, codeSource, destination, 0, deadline);
        vm.expectRevert(ISweepRegistry.SourceNotEOA.selector);
        vm.prank(destination);
        registry.registerSweep(codeSource, destination, 0, deadline, sig);
    }

    function test_register_reverts_destination_zero() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, address(0), 0, deadline);
        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(destination);
        registry.registerSweep(source, address(0), 0, deadline, sig);
    }

    function test_register_reverts_destination_in_system_segment() public {
        address sysDestination = 0x5300000000000000000000000000000000000001;
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, sysDestination, 0, deadline);
        vm.expectRevert(ISweepRegistry.SystemAddressNotAllowed.selector);
        vm.prank(sysDestination);
        registry.registerSweep(source, sysDestination, 0, deadline, sig);
    }

    function test_register_reverts_source_in_system_segment() public {
        address sysSource = 0x5300000000000000000000000000000000000009;
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, sysSource, destination, 0, deadline);
        vm.expectRevert(ISweepRegistry.SystemAddressNotAllowed.selector);
        vm.prank(destination);
        registry.registerSweep(sysSource, destination, 0, deadline, sig);
    }

    function test_register_reverts_destination_is_active_source() public {
        // `source` becomes active with destination `destination`.
        _register(destination);

        // Now try to use `source` (an active sweep source) as a destination.
        uint256 pkB = 0xB0B0;
        address sourceB = vm.addr(pkB);
        uint64 deadline = _deadline();
        bytes memory sig = _sign(pkB, sourceB, source, 0, deadline);
        vm.expectRevert(ISweepRegistry.DestinationIsActiveSource.selector);
        vm.prank(source);
        registry.registerSweep(sourceB, source, 0, deadline, sig);
    }

    /// @dev §14 item 9: self-reference is refused at registration. Sweeping an
    ///      address to itself is a guaranteed no-op that would still burn a
    ///      candidate slot on every inflow.
    function test_register_reverts_self_reference() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, source, 0, deadline);
        vm.expectRevert(ISweepRegistry.DestinationIsSource.selector);
        vm.prank(source);
        registry.registerSweep(source, source, 0, deadline, sig);
    }

    /// @dev Self-reference must be reported as such even when the signature is
    ///      invalid — the cheap structural check runs before `ecrecover`.
    function test_register_self_reference_precedes_signature_check() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(0xB0B, source, source, 0, deadline);
        vm.expectRevert(ISweepRegistry.DestinationIsSource.selector);
        vm.prank(source);
        registry.registerSweep(source, source, 0, deadline, sig);
    }

    function test_register_reverts_already_registered() public {
        _register(destination);
        // Second registration with a fresh (nonce=1) signature still rejected.
        uint64 deadline = _deadline();
        bytes memory sig = _sign(sourcePk, source, destination, 1, deadline);
        vm.expectRevert(ISweepRegistry.SourceAlreadyRegistered.selector);
        vm.prank(destination);
        registry.registerSweep(source, destination, 1, deadline, sig);
    }

    /*//////////////////////////////////////////////////////////////
                        resolveSweep three checks
    //////////////////////////////////////////////////////////////*/

    function test_resolveSweep_zero_when_unregistered() public {
        assertEq(registry.resolveSweep(token, source), address(0));
    }

    function test_resolveSweep_zero_when_token_not_whitelisted() public {
        _register(destination);
        vm.prank(owner);
        registry.setTokenWhitelist(token, false);
        assertEq(registry.resolveSweep(token, source), address(0));
    }

    function test_resolveSweep_zero_when_disabled() public {
        _register(destination);
        vm.prank(destination);
        registry.disableSweep(source);
        assertEq(registry.resolveSweep(token, source), address(0));
    }

    /*//////////////////////////////////////////////////////////////
                    EL storage-layout consensus surface
    //////////////////////////////////////////////////////////////*/

    /// @dev The execution layer resolves candidates with two raw SLOADs instead
    ///      of calling {ISweepRegistry.resolveSweep} (spec §14 item 10 ③). This
    ///      contract sits behind a proxy, so an upgrade that reorders or repacks
    ///      the storage declarations would silently change what the EL reads —
    ///      a hardfork disguised as an upgrade. This test fails if that happens.
    function test_storage_layout_frozen_for_el() public {
        _register(destination);

        bytes32 recordBase = keccak256(abi.encode(source, SOURCES_SLOT));
        uint256 word0 = uint256(vm.load(address(registry), recordBase));

        // base + 0, low 20 bytes -> destination
        assertEq(address(uint160(word0)), destination, "destination must sit at base+0 offset 0");
        // base + 0, byte offset 20 -> enabled
        assertEq((word0 >> 160) & 0xFF, 1, "enabled must sit at base+0 offset 20");
        // Nothing else may share the first word.
        assertEq(word0 >> 168, 0, "base+0 must hold only destination and enabled");
        // base + 1 -> nonce
        assertEq(
            uint256(vm.load(address(registry), bytes32(uint256(recordBase) + 1))),
            1,
            "nonce must sit at base+1"
        );

        // tokenWhitelist lives one slot after sources.
        bytes32 whitelistSlot = keccak256(abi.encode(token, TOKEN_WHITELIST_SLOT));
        assertEq(uint256(vm.load(address(registry), whitelistSlot)), 1, "whitelist flag slot moved");
    }

    /// @dev The EL's slot read and the Solidity resolver must agree in every
    ///      state, otherwise the two clients of this layout can disagree about
    ///      whether a candidate is sweepable.
    function test_el_slot_read_matches_resolveSweep_in_every_state() public {
        // Unregistered.
        assertEq(_resolveViaSlots(token, source), registry.resolveSweep(token, source));
        assertEq(_resolveViaSlots(token, source), address(0));

        // Registered and whitelisted.
        _register(destination);
        assertEq(_resolveViaSlots(token, source), registry.resolveSweep(token, source));
        assertEq(_resolveViaSlots(token, source), destination);

        // Token de-whitelisted.
        vm.prank(owner);
        registry.setTokenWhitelist(token, false);
        assertEq(_resolveViaSlots(token, source), registry.resolveSweep(token, source));
        assertEq(_resolveViaSlots(token, source), address(0));

        // Re-whitelisted, then source disabled.
        vm.prank(owner);
        registry.setTokenWhitelist(token, true);
        vm.prank(destination);
        registry.disableSweep(source);
        assertEq(_resolveViaSlots(token, source), registry.resolveSweep(token, source));
        assertEq(_resolveViaSlots(token, source), address(0));
    }

    /*//////////////////////////////////////////////////////////////
                                Disable
    //////////////////////////////////////////////////////////////*/

    function test_disable_by_destination_emits() public {
        _register(destination);
        vm.expectEmit(true, true, true, false, address(registry));
        emit ISweepRegistry.SweepDisabled(source, destination, destination);
        vm.prank(destination);
        registry.disableSweep(source);

        (, bool enabled, ) = registry.getSweep(source);
        assertFalse(enabled);
    }

    function test_disable_by_operator() public {
        _register(destination);
        vm.prank(destination);
        registry.setSweepOperator(operator, true);
        vm.prank(operator);
        registry.disableSweep(source);
        (, bool enabled, ) = registry.getSweep(source);
        assertFalse(enabled);
    }

    function test_disable_reverts_not_authorized() public {
        _register(destination);
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(stranger);
        registry.disableSweep(source);
    }

    function test_disable_reverts_not_authorized_when_unregistered() public {
        // Unregistered source has destination == address(0); a non-zero caller
        // fails the authorization check before the active check.
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(destination);
        registry.disableSweep(source);
    }

    function test_disable_reverts_not_active_on_double_disable() public {
        _register(destination);
        vm.prank(destination);
        registry.disableSweep(source);
        // Second disable: authorized (destination matches) but already inactive.
        vm.expectRevert(ISweepRegistry.SourceNotActive.selector);
        vm.prank(destination);
        registry.disableSweep(source);
    }

    /*//////////////////////////////////////////////////////////////
                                 Poke
    //////////////////////////////////////////////////////////////*/

    function test_poke_succeeds_permissionless() public {
        _register(destination);
        vm.expectEmit(true, true, false, false, address(registry));
        emit ISweepRegistry.SweepRequested(token, source);
        vm.prank(stranger); // permissionless
        registry.pokeSweep(token, source);
    }

    function test_poke_reverts_token_not_whitelisted() public {
        _register(destination);
        address other = address(0xF00D);
        vm.expectRevert(ISweepRegistry.TokenNotWhitelisted.selector);
        registry.pokeSweep(other, source);
    }

    function test_poke_reverts_source_not_active() public {
        vm.expectRevert(ISweepRegistry.SourceNotActive.selector);
        registry.pokeSweep(token, source);
    }

    function test_poke_reverts_after_disable() public {
        _register(destination);
        vm.prank(destination);
        registry.disableSweep(source);
        // v1: disabled sources cannot be poked.
        vm.expectRevert(ISweepRegistry.SourceNotActive.selector);
        registry.pokeSweep(token, source);
    }

    /*//////////////////////////////////////////////////////////////
                                 Batch
    //////////////////////////////////////////////////////////////*/

    function test_batch_register() public {
        uint256 pkB = 0xB0B0;
        address sourceB = vm.addr(pkB);
        uint64 deadline = _deadline();

        ISweepRegistry.SweepRegistration[]
            memory regs = new ISweepRegistry.SweepRegistration[](2);
        regs[0] = ISweepRegistry.SweepRegistration({
            source: source,
            destination: destination,
            nonce: 0,
            deadline: deadline,
            sourceSignature: _sign(sourcePk, source, destination, 0, deadline)
        });
        regs[1] = ISweepRegistry.SweepRegistration({
            source: sourceB,
            destination: destination,
            nonce: 0,
            deadline: deadline,
            sourceSignature: _sign(pkB, sourceB, destination, 0, deadline)
        });

        vm.prank(destination);
        registry.registerSweeps(regs);

        assertEq(registry.resolveSweep(token, source), destination);
        assertEq(registry.resolveSweep(token, sourceB), destination);
    }

    /*//////////////////////////////////////////////////////////////
                                 Pause
    //////////////////////////////////////////////////////////////*/

    function test_pause_blocks_mutations_but_not_resolve() public {
        _register(destination);
        vm.prank(owner);
        registry.setPause(true);

        // resolveSweep (view) is unaffected by pause.
        assertEq(registry.resolveSweep(token, source), destination);

        // Mutating entrypoints revert while paused.
        vm.expectRevert("Pausable: paused");
        registry.pokeSweep(token, source);

        uint64 deadline = _deadline();
        uint256 pkB = 0xB0B0;
        address sourceB = vm.addr(pkB);
        bytes memory sig = _sign(pkB, sourceB, destination, 0, deadline);
        vm.expectRevert("Pausable: paused");
        vm.prank(destination);
        registry.registerSweep(sourceB, destination, 0, deadline, sig);

        vm.expectRevert("Pausable: paused");
        vm.prank(destination);
        registry.disableSweep(source);
    }

    function test_setPause_onlyOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(stranger);
        registry.setPause(true);
    }
}
