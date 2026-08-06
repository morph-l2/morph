// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Test} from "forge-std/Test.sol";

import {SweepRegistry} from "../l2/system/SweepRegistry.sol";
import {ISweepRegistry} from "../l2/system/ISweepRegistry.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract SweepRegistryTest is Test {
    SweepRegistry internal registry;

    /// The registry is authorization-only and never calls the token, so a bare
    /// address is a complete stand-in. A mock with a `balanceOf` would suggest
    /// otherwise.
    address internal token = address(0x7A0);

    address internal multisig = address(512);
    address internal owner = address(64);
    address internal controller = address(0xC011);
    address internal otherController = address(0xC012);
    address internal operator = address(0xACE);
    address internal stranger = address(0xBAD);
    address internal destination = address(0x11115);
    address internal newDestination = address(0x22225);

    uint256 internal sourcePk = 0xA11CE;
    address internal source;

    bytes32 internal constant AUTH_TYPEHASH =
        keccak256("SweepAuthorization(address source,address controller,uint64 deadline)");

    function setUp() public {
        source = vm.addr(sourcePk);

        vm.startPrank(multisig);
        ProxyAdmin proxyAdmin = new ProxyAdmin();
        SweepRegistry implementation = new SweepRegistry();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(implementation),
            address(proxyAdmin),
            abi.encodeWithSelector(SweepRegistry.initialize.selector, owner)
        );
        vm.stopPrank();
        registry = SweepRegistry(address(proxy));

        vm.prank(owner);
        registry.setTokenWhitelist(token, true);

        vm.prank(controller);
        registry.setSweepDestination(destination);
    }

    function _deadline() internal view returns (uint64) {
        return uint64(block.timestamp + 1 hours);
    }

    function _sign(uint256 pk, address source_, address controller_, uint64 deadline_)
        internal
        view
        returns (bytes memory)
    {
        bytes32 structHash = keccak256(abi.encode(AUTH_TYPEHASH, source_, controller_, deadline_));
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", registry.DOMAIN_SEPARATOR(), structHash));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(pk, digest);
        return abi.encodePacked(r, s, v);
    }

    function _register(address caller) internal {
        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, controller, deadline);
        vm.prank(caller);
        registry.registerSweep(source, controller, deadline, signature);
    }

    /// A fresh source EOA plus its authorization for `controller`.
    function _newSource(uint256 pk, uint64 deadline) internal view returns (address, bytes memory) {
        address newSource = vm.addr(pk);
        return (newSource, _sign(pk, newSource, controller, deadline));
    }

    function _authorizeOperator() internal {
        vm.prank(controller);
        registry.setSweepOperator(operator, true);
    }

    function test_controller_destination_is_a_mutable_pointer_for_all_sources() public {
        _register(controller);
        assertEq(registry.resolveSweep(token, source), destination);

        vm.expectEmit(true, true, true, false, address(registry));
        emit ISweepRegistry.SweepDestinationChanged(controller, destination, newDestination);
        vm.prank(controller);
        registry.setSweepDestination(newDestination);

        assertEq(registry.resolveSweep(token, source), newDestination);
        (address boundController, address resolvedDestination, bool enabled) = registry.getSweep(source);
        assertEq(boundController, controller);
        assertEq(resolvedDestination, newDestination);
        assertTrue(enabled);
    }

    function test_zero_destination_stops_and_a_new_destination_restarts_all_sources() public {
        _register(controller);

        vm.prank(controller);
        registry.setSweepDestination(address(0));
        assertEq(registry.resolveSweep(token, source), address(0));

        vm.prank(controller);
        registry.setSweepDestination(newDestination);
        assertEq(registry.resolveSweep(token, source), newDestination);
    }

    function test_each_controller_can_only_change_its_own_destination() public {
        _authorizeOperator();

        vm.prank(operator);
        registry.setSweepDestination(newDestination);
        vm.prank(destination);
        registry.setSweepDestination(newDestination);

        assertEq(registry.destinations(controller), destination);
        assertEq(registry.destinations(operator), newDestination);
        assertEq(registry.destinations(destination), newDestination);
    }

    /// Two controllers may share a destination, and an address that once served
    /// as a destination is NOT permanently blocked from becoming a source: role
    /// overlap is allowed because the EL fixes the sweep candidate set before
    /// execution, so a destination-turned-source cannot recurse through
    /// sweep-generated logs (bounded per-tx/block by the EL's budget).
    function test_shared_destination_does_not_block_future_source_registration() public {
        // Two controllers may share a destination.
        vm.prank(otherController);
        registry.setSweepDestination(destination);

        // Flag a fresh EOA as a destination, then have every controller move away.
        uint256 stalePk = 0x5EED;
        address stale = vm.addr(stalePk);
        vm.prank(controller);
        registry.setSweepDestination(stale);
        vm.prank(controller);
        registry.setSweepDestination(address(0));
        vm.prank(otherController);
        registry.setSweepDestination(address(0));

        // Point the controller somewhere live again so registration reaches the
        // signature check instead of failing on a cleared pointer.
        vm.prank(controller);
        registry.setSweepDestination(destination);

        // `stale` was a destination but nobody points at it now; it may register
        // as a source.
        uint64 deadline = _deadline();
        bytes memory signature = _sign(stalePk, stale, controller, deadline);
        vm.prank(controller);
        registry.registerSweep(stale, controller, deadline, signature);

        assertEq(registry.resolveSweep(token, stale), destination);
    }

    /// An active source may be configured as another controller's destination,
    /// and a current destination may register as a source. Registration allows
    /// the overlap, but the resolver dynamically refuses to sweep INTO an active
    /// source — which breaks self-references and multi-address cycles (A→B,
    /// B→A) without reintroducing a permanent address ban. The EL freezes the
    /// candidate set per transaction, so this dynamic check is what stops a
    /// permissionless `pokeSweep` from driving a cross-transaction ping-pong.
    function test_roles_may_overlap_but_resolver_blocks_sweep_cycles() public {
        // A = source, B = payout. Build a two-address cycle:
        //   A (under controller) → destination B
        //   B (under otherController) → destination A
        uint256 payoutPk = 0xF00D;
        address payout = vm.addr(payoutPk);

        vm.prank(controller);
        registry.setSweepDestination(payout);
        vm.prank(otherController);
        registry.setSweepDestination(source);

        uint64 deadline = _deadline();
        bytes memory sigA = _sign(sourcePk, source, controller, deadline);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, sigA);

        bytes memory sigB = _sign(payoutPk, payout, otherController, deadline);
        vm.prank(otherController);
        registry.registerSweep(payout, otherController, deadline, sigB);

        // Both arms of the cycle resolve to zero: sweeping into an active source
        // is refused from either direction.
        assertEq(registry.resolveSweep(token, source), address(0), "A->B blocked: B is an active source");
        assertEq(registry.resolveSweep(token, payout), address(0), "B->A blocked: A is an active source");
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, payout);

        // Disabling A (source) breaks the cycle: B becomes sweepable again, into
        // the now-disabled A which can receive the sweep.
        vm.prank(controller);
        registry.disableSweep(source);
        assertEq(registry.resolveSweep(token, payout), source, "cycle broken: B sweepable into disabled A");
        registry.pokeSweep(token, payout);
    }

    /// Self-reference must be refused by the resolver even though registration
    /// allows it: an active source that points at itself would otherwise be a
    /// single-address cycle every poke could re-drive.
    function test_resolver_blocks_self_reference_dynamically() public {
        vm.prank(controller);
        registry.setSweepDestination(source);

        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, controller, deadline);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, signature);

        assertEq(registry.resolveSweep(token, source), address(0), "self-reference not sweepable");
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);
    }

    function test_operator_registers_but_cannot_redirect_funds_or_disable() public {
        _authorizeOperator();
        _register(operator);
        assertEq(registry.resolveSweep(token, source), destination);

        // The operator cannot move the controller's destination pointer (only the
        // controller's own), so the source still resolves to `destination`.
        vm.prank(operator);
        registry.setSweepDestination(newDestination);
        assertEq(registry.resolveSweep(token, source), destination);

        // The operator cannot terminally disable a source either: disable is
        // terminal and a leaked operator hot key must not permanently retire the
        // controller's sources.
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(operator);
        registry.disableSweep(source);

        // The controller still can.
        vm.prank(controller);
        registry.disableSweep(source);
        assertEq(registry.resolveSweep(token, source), address(0));
    }

    /// The source owner may terminally disable its own source, which is how a
    /// source revokes its earlier authorization after the fact.
    function test_source_can_disable_itself() public {
        _register(controller);
        assertEq(registry.resolveSweep(token, source), destination);

        vm.expectEmit(true, true, true, true, address(registry));
        emit ISweepRegistry.SweepDisabled(source, controller, destination, source);
        vm.prank(source);
        registry.disableSweep(source);

        assertEq(registry.resolveSweep(token, source), address(0));
        (,, bool enabled) = registry.getSweep(source);
        assertFalse(enabled);
    }

    function test_registration_binds_the_source_to_the_controller() public {
        vm.expectEmit(true, true, true, true, address(registry));
        emit ISweepRegistry.SweepRegistered(source, controller, destination, controller);
        _register(controller);

        (address boundController, address resolvedDestination, bool enabled) = registry.getSweep(source);
        assertEq(boundController, controller);
        assertEq(resolvedDestination, destination);
        assertTrue(enabled);
    }

    function test_registration_rejects_unknown_controller_and_unauthorized_submitter() public {
        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, otherController, deadline);

        vm.expectRevert(ISweepRegistry.DestinationNotConfigured.selector);
        registry.registerSweep(source, otherController, deadline, signature);

        signature = _sign(sourcePk, source, controller, deadline);
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(stranger);
        registry.registerSweep(source, controller, deadline, signature);
    }

    function test_registration_rejects_bad_signature_expiry_and_replay() public {
        uint64 deadline = _deadline();
        bytes memory badSignature = _sign(0xB0B, source, controller, deadline);
        vm.expectRevert(ISweepRegistry.InvalidSignature.selector);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, badSignature);

        uint64 expiredDeadline = uint64(block.timestamp);
        bytes memory expiredSignature = _sign(sourcePk, source, controller, expiredDeadline);
        vm.warp(block.timestamp + 1);
        vm.expectRevert(ISweepRegistry.SignatureExpired.selector);
        vm.prank(controller);
        registry.registerSweep(source, controller, expiredDeadline, expiredSignature);

        deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, controller, deadline);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, signature);

        vm.expectRevert(ISweepRegistry.SourceAlreadyRegistered.selector);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, signature);
    }

    function test_registration_rejects_code_source() public {
        address codeSource = address(0xC0DE);
        vm.etch(codeSource, hex"600060005360016000f3");
        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, codeSource, controller, deadline);

        vm.expectRevert(ISweepRegistry.SourceNotEOA.selector);
        vm.prank(controller);
        registry.registerSweep(codeSource, controller, deadline, signature);
    }

    function test_batch_registration_binds_every_entry() public {
        uint64 deadline = _deadline();
        (address second, bytes memory secondSignature) = _newSource(0xB0B0, deadline);

        ISweepRegistry.SweepRegistration[] memory registrations = new ISweepRegistry.SweepRegistration[](2);
        registrations[0] = ISweepRegistry.SweepRegistration({
            source: source,
            controller: controller,
            deadline: deadline,
            sourceSignature: _sign(sourcePk, source, controller, deadline)
        });
        registrations[1] = ISweepRegistry.SweepRegistration({
            source: second, controller: controller, deadline: deadline, sourceSignature: secondSignature
        });

        vm.prank(controller);
        registry.registerSweeps(registrations);

        assertEq(registry.resolveSweep(token, source), destination);
        assertEq(registry.resolveSweep(token, second), destination);
    }

    function test_disable_is_terminal_and_poke_is_permissionless() public {
        _register(controller);

        vm.expectEmit(true, true, true, true, address(registry));
        emit ISweepRegistry.SweepRequested(token, source);
        vm.prank(stranger);
        registry.pokeSweep(token, source);

        vm.expectEmit(true, true, true, true, address(registry));
        emit ISweepRegistry.SweepDisabled(source, controller, destination, controller);
        vm.prank(controller);
        registry.disableSweep(source);

        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);

        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, controller, deadline);
        vm.expectRevert(ISweepRegistry.SourceAlreadyRegistered.selector);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, signature);
    }

    /// disableSweep keeps the source's binding (it is terminal), so a disabled
    /// source may become another controller's destination, but re-registering it
    /// must still report the one-shot binding (`SourceAlreadyRegistered`).
    function test_disabled_source_used_as_destination_reports_already_registered() public {
        _register(controller);
        vm.prank(controller);
        registry.disableSweep(source);
        (address boundController,,) = registry.getSweep(source);
        assertEq(boundController, controller, "disable keeps the binding");

        vm.prank(otherController);
        registry.setSweepDestination(source);
        assertEq(registry.destinations(otherController), source, "disabled source may become a destination");

        uint64 deadline = _deadline();
        bytes memory signature = _sign(sourcePk, source, controller, deadline);
        vm.expectRevert(ISweepRegistry.SourceAlreadyRegistered.selector);
        vm.prank(controller);
        registry.registerSweep(source, controller, deadline, signature);
    }

    function test_pause_blocks_registration_but_not_destination_changes() public {
        vm.prank(owner);
        registry.setPause(true);

        uint64 deadline = _deadline();
        (address second, bytes memory signature) = _newSource(0xB0B0, deadline);
        vm.expectRevert("Pausable: paused");
        vm.prank(controller);
        registry.registerSweep(second, controller, deadline, signature);

        // A controller can still move or clear its own pointer while paused, so a
        // governance pause never traps a controller's funds at a stale address.
        vm.prank(controller);
        registry.setSweepDestination(newDestination);
        assertEq(registry.destinations(controller), newDestination);
    }

    /// `pokeSweep` and `resolveSweep` are one rule, so every condition that makes
    /// the resolver return zero must also make a poke revert. Otherwise a poke
    /// admits a request the EL lifts into a candidate, charges the block's sweep
    /// preflight budget for, and then resolves to zero — permissionlessly.
    function test_poke_refuses_exactly_what_the_resolver_refuses() public {
        _register(controller);

        // Governance pause. This is the condition the two used to disagree on.
        vm.prank(owner);
        registry.setPause(true);
        assertEq(registry.resolveSweep(token, source), address(0), "paused: resolver refuses");
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);
        vm.prank(owner);
        registry.setPause(false);

        // Token de-whitelisted.
        vm.prank(owner);
        registry.setTokenWhitelist(token, false);
        assertEq(registry.resolveSweep(token, source), address(0));
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);
        vm.prank(owner);
        registry.setTokenWhitelist(token, true);

        // Controller pauses itself by clearing its destination.
        vm.prank(controller);
        registry.setSweepDestination(address(0));
        assertEq(registry.resolveSweep(token, source), address(0));
        vm.expectRevert(ISweepRegistry.NotSweepable.selector);
        registry.pokeSweep(token, source);
        vm.prank(controller);
        registry.setSweepDestination(destination);

        // Back to sweepable: both agree again.
        assertEq(registry.resolveSweep(token, source), destination);
        registry.pokeSweep(token, source);
    }

    /// The EL hardcodes these bytes (`morph-reth/crates/revm/src/sweep.rs`), so a
    /// rename, a swapped parameter, or a reordered event field compiles clean,
    /// passes every behavioural test, and forks the chain. This is the only place
    /// a Solidity-side change to the frozen surface becomes visible.
    function test_frozen_el_surface_bytes_are_unchanged() public {
        assertEq(
            bytes32(ISweepRegistry.resolveSweep.selector),
            bytes32(bytes4(0x9faa2f2f)),
            "resolveSweep(address,address) selector"
        );
        assertEq(
            ISweepRegistry.SweepRequested.selector,
            bytes32(0x24e3f180db341974dcd99a5e223d9d944422e303230ddde6659302f8620bbcff),
            "SweepRequested topic"
        );
        assertEq(
            ISweepRegistry.Swept.selector,
            bytes32(0x035b37215a69e14a80883933d6aa84f0919a67af9410a4a73e8a23baeca011f0),
            "Swept topic"
        );
        assertEq(
            ISweepRegistry.SweepFailed.selector,
            bytes32(0x0f64fa58e4261d8832b5ea6c262c691ef36e73cb21998c4fb01a83997940797c),
            "SweepFailed topic"
        );
    }

    /// The EL classifies a reverting resolver call as a resolver failure, so
    /// `resolveSweep` must answer with a zero word — never a revert — for every
    /// input it can be handed. Walk every refusal condition and every degenerate
    /// input through a raw 50k-gas staticcall, the way the EL calls it, rather
    /// than through the typed binding which would hide both the cap and the
    /// revert/zero distinction.
    function test_resolver_never_reverts_and_always_returns_one_word() public {
        // What the EL passes for essentially every transfer on the chain.
        _assertResolves(token, source, address(0), "never-registered source");
        _assertResolves(address(0), address(0), address(0), "zero token and source");
        _assertResolves(address(registry), source, address(0), "registry as token");

        _register(controller);
        _assertResolves(token, source, destination, "sweepable");

        vm.prank(owner);
        registry.setPause(true);
        _assertResolves(token, source, address(0), "paused");
        vm.prank(owner);
        registry.setPause(false);

        vm.prank(owner);
        registry.setTokenWhitelist(token, false);
        _assertResolves(token, source, address(0), "token de-whitelisted");
        vm.prank(owner);
        registry.setTokenWhitelist(token, true);

        vm.prank(controller);
        registry.setSweepDestination(address(0));
        _assertResolves(token, source, address(0), "controller destination cleared");
        vm.prank(controller);
        registry.setSweepDestination(destination);

        // Destination is itself an active source: refused (zero) without revert.
        // Register `cycleDest` as an active source, then point `controller` (the
        // owner of `source`) at it — so `source`'s destination is an active source.
        uint256 cyclePk = 0xC1CE;
        address cycleDest = vm.addr(cyclePk);
        vm.prank(otherController);
        registry.setSweepDestination(cycleDest);
        uint64 deadline = _deadline();
        bytes memory cycleSig = _sign(cyclePk, cycleDest, controller, deadline);
        vm.prank(controller);
        registry.registerSweep(cycleDest, controller, deadline, cycleSig);
        vm.prank(controller);
        registry.setSweepDestination(cycleDest);
        _assertResolves(token, source, address(0), "destination is an active source");
        // A self-referencing source is likewise refused.
        vm.prank(controller);
        registry.setSweepDestination(source);
        _assertResolves(token, source, address(0), "self-reference");

        vm.prank(controller);
        registry.disableSweep(source);
        _assertResolves(token, source, address(0), "source disabled");
    }

    function _assertResolves(address token_, address source_, address expected, string memory label) internal {
        (bool success, bytes memory output) =
            address(registry).staticcall{gas: 50_000}(abi.encodeCall(ISweepRegistry.resolveSweep, (token_, source_)));

        assertTrue(success, string.concat(label, ": resolver reverted or ran out of gas"));
        assertEq(output.length, 32, string.concat(label, ": output is not exactly one word"));
        assertEq(abi.decode(output, (address)), expected, string.concat(label, ": wrong destination"));
    }

    function test_admin_setters_are_owner_only_and_no_op_changes_revert() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(stranger);
        registry.setTokenWhitelist(token, false);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(stranger);
        registry.setPause(true);

        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(owner);
        registry.setTokenWhitelist(address(0), true);

        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(controller);
        registry.setSweepOperator(address(0), true);

        vm.expectRevert(ISweepRegistry.DestinationUnchanged.selector);
        vm.prank(controller);
        registry.setSweepDestination(destination);

        vm.expectRevert(ISweepRegistry.SourceNotActive.selector);
        vm.prank(controller);
        registry.disableSweep(source);
    }
}
