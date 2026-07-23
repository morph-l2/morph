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

    // Deposit key pair (deterministic for signing).
    uint256 internal depositPk = 0xA11CE;
    address internal deposit;

    address internal master = address(0x11115);
    address internal token = address(0xDEADBEEF);

    // Must match the contract's frozen tags (asserted in a dedicated test).
    bytes32 internal constant AUTH_TYPEHASH =
        keccak256(
            "SweepAuthorization(address deposit,address master,address registry,uint256 chainId,uint256 nonce,uint64 deadline,bytes32 mode,bytes32 sweepScope)"
        );
    bytes32 internal constant MODE = keccak256("MORPH_SWEEP_V1");
    bytes32 internal constant SWEEP_SCOPE = keccak256("WHITELISTED_ERC20_TO_MASTER_ONLY");

    function setUp() public {
        deposit = vm.addr(depositPk);

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
        vm.label(deposit, "deposit");
        vm.label(master, "master");
        vm.label(operator, "operator");
    }

    /*//////////////////////////////////////////////////////////////
                                Helpers
    //////////////////////////////////////////////////////////////*/

    function _sign(
        uint256 pk,
        address deposit_,
        address master_,
        uint256 nonce_,
        uint64 deadline_
    ) internal view returns (bytes memory) {
        bytes32 structHash = keccak256(
            abi.encode(
                AUTH_TYPEHASH,
                deposit_,
                master_,
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
        bytes memory sig = _sign(depositPk, deposit, master, 0, deadline);
        vm.prank(caller);
        registry.registerSweep(deposit, master, 0, deadline, sig);
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
        emit ISweepRegistry.SweepOperatorSet(master, operator, true);
        vm.prank(master);
        registry.setSweepOperator(operator, true);
        assertTrue(registry.operators(master, operator));
    }

    function test_setSweepOperator_reverts_zero() public {
        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(master);
        registry.setSweepOperator(address(0), true);
    }

    /*//////////////////////////////////////////////////////////////
                              Register
    //////////////////////////////////////////////////////////////*/

    function test_register_succeeds_and_resolves() public {
        vm.expectEmit(true, true, true, false, address(registry));
        emit ISweepRegistry.SweepRegistered(deposit, master, master);
        _register(master);

        (address m, bool enabled, uint256 nonce) = registry.getSweep(deposit);
        assertEq(m, master);
        assertTrue(enabled);
        assertEq(nonce, 1);
        assertEq(registry.resolveSweep(token, deposit), master);
    }

    function test_register_via_operator() public {
        vm.prank(master);
        registry.setSweepOperator(operator, true);

        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, master, 0, deadline);
        vm.prank(operator);
        registry.registerSweep(deposit, master, 0, deadline, sig);

        assertEq(registry.resolveSweep(token, deposit), master);
    }

    function test_register_reverts_not_authorized() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, master, 0, deadline);
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(stranger);
        registry.registerSweep(deposit, master, 0, deadline, sig);
    }

    function test_register_reverts_bad_signature() public {
        uint64 deadline = _deadline();
        // Sign with the wrong key -> recovers to a different address.
        bytes memory sig = _sign(0xB0B, deposit, master, 0, deadline);
        vm.expectRevert(ISweepRegistry.InvalidSignature.selector);
        vm.prank(master);
        registry.registerSweep(deposit, master, 0, deadline, sig);
    }

    function test_register_reverts_expired() public {
        uint64 deadline = uint64(block.timestamp);
        bytes memory sig = _sign(depositPk, deposit, master, 0, deadline);
        vm.warp(block.timestamp + 2);
        vm.expectRevert(ISweepRegistry.SignatureExpired.selector);
        vm.prank(master);
        registry.registerSweep(deposit, master, 0, deadline, sig);
    }

    function test_register_reverts_bad_nonce() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, master, 1, deadline);
        vm.expectRevert(ISweepRegistry.InvalidNonce.selector);
        vm.prank(master);
        registry.registerSweep(deposit, master, 1, deadline, sig);
    }

    function test_register_reverts_deposit_has_code() public {
        address codeDeposit = address(0xC0DE);
        vm.etch(codeDeposit, hex"600060005360016000f3");
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, codeDeposit, master, 0, deadline);
        vm.expectRevert(ISweepRegistry.DepositNotEOA.selector);
        vm.prank(master);
        registry.registerSweep(codeDeposit, master, 0, deadline, sig);
    }

    function test_register_reverts_master_zero() public {
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, address(0), 0, deadline);
        vm.expectRevert(ISweepRegistry.ZeroAddress.selector);
        vm.prank(master);
        registry.registerSweep(deposit, address(0), 0, deadline, sig);
    }

    function test_register_reverts_master_in_system_segment() public {
        address sysMaster = 0x5300000000000000000000000000000000000001;
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, sysMaster, 0, deadline);
        vm.expectRevert(ISweepRegistry.SystemAddressNotAllowed.selector);
        vm.prank(sysMaster);
        registry.registerSweep(deposit, sysMaster, 0, deadline, sig);
    }

    function test_register_reverts_deposit_in_system_segment() public {
        address sysDeposit = 0x5300000000000000000000000000000000000009;
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, sysDeposit, master, 0, deadline);
        vm.expectRevert(ISweepRegistry.SystemAddressNotAllowed.selector);
        vm.prank(master);
        registry.registerSweep(sysDeposit, master, 0, deadline, sig);
    }

    function test_register_reverts_master_is_active_deposit() public {
        // depositA becomes active with master `master`.
        _register(master);

        // Now try to use `deposit` (an active sweep deposit) as a master.
        uint256 pkB = 0xB0B0;
        address depositB = vm.addr(pkB);
        uint64 deadline = _deadline();
        bytes memory sig = _sign(pkB, depositB, deposit, 0, deadline);
        vm.expectRevert(ISweepRegistry.MasterIsSweepDeposit.selector);
        vm.prank(deposit);
        registry.registerSweep(depositB, deposit, 0, deadline, sig);
    }

    function test_register_reverts_already_registered() public {
        _register(master);
        // Second registration with a fresh (nonce=1) signature still rejected.
        uint64 deadline = _deadline();
        bytes memory sig = _sign(depositPk, deposit, master, 1, deadline);
        vm.expectRevert(ISweepRegistry.DepositAlreadyRegistered.selector);
        vm.prank(master);
        registry.registerSweep(deposit, master, 1, deadline, sig);
    }

    /*//////////////////////////////////////////////////////////////
                        resolveSweep three checks
    //////////////////////////////////////////////////////////////*/

    function test_resolveSweep_zero_when_unregistered() public {
        assertEq(registry.resolveSweep(token, deposit), address(0));
    }

    function test_resolveSweep_zero_when_token_not_whitelisted() public {
        _register(master);
        vm.prank(owner);
        registry.setTokenWhitelist(token, false);
        assertEq(registry.resolveSweep(token, deposit), address(0));
    }

    function test_resolveSweep_zero_when_disabled() public {
        _register(master);
        vm.prank(master);
        registry.disableSweep(deposit);
        assertEq(registry.resolveSweep(token, deposit), address(0));
    }

    /*//////////////////////////////////////////////////////////////
                                Disable
    //////////////////////////////////////////////////////////////*/

    function test_disable_by_master_emits() public {
        _register(master);
        vm.expectEmit(true, true, true, false, address(registry));
        emit ISweepRegistry.SweepDisabled(deposit, master, master);
        vm.prank(master);
        registry.disableSweep(deposit);

        (, bool enabled, ) = registry.getSweep(deposit);
        assertFalse(enabled);
    }

    function test_disable_by_operator() public {
        _register(master);
        vm.prank(master);
        registry.setSweepOperator(operator, true);
        vm.prank(operator);
        registry.disableSweep(deposit);
        (, bool enabled, ) = registry.getSweep(deposit);
        assertFalse(enabled);
    }

    function test_disable_reverts_not_authorized() public {
        _register(master);
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(stranger);
        registry.disableSweep(deposit);
    }

    function test_disable_reverts_not_authorized_when_unregistered() public {
        // Unregistered deposit has master == address(0); a non-zero caller fails
        // the authorization check before the active check.
        vm.expectRevert(ISweepRegistry.NotAuthorized.selector);
        vm.prank(master);
        registry.disableSweep(deposit);
    }

    function test_disable_reverts_not_active_on_double_disable() public {
        _register(master);
        vm.prank(master);
        registry.disableSweep(deposit);
        // Second disable: authorized (master matches) but already inactive.
        vm.expectRevert(ISweepRegistry.DepositNotActive.selector);
        vm.prank(master);
        registry.disableSweep(deposit);
    }

    /*//////////////////////////////////////////////////////////////
                                 Poke
    //////////////////////////////////////////////////////////////*/

    function test_poke_succeeds_permissionless() public {
        _register(master);
        vm.expectEmit(true, true, false, false, address(registry));
        emit ISweepRegistry.SweepRequested(token, deposit);
        vm.prank(stranger); // permissionless
        registry.pokeSweep(token, deposit);
    }

    function test_poke_reverts_token_not_whitelisted() public {
        _register(master);
        address other = address(0xF00D);
        vm.expectRevert(ISweepRegistry.TokenNotWhitelisted.selector);
        registry.pokeSweep(other, deposit);
    }

    function test_poke_reverts_deposit_not_active() public {
        vm.expectRevert(ISweepRegistry.DepositNotActive.selector);
        registry.pokeSweep(token, deposit);
    }

    function test_poke_reverts_after_disable() public {
        _register(master);
        vm.prank(master);
        registry.disableSweep(deposit);
        // v1: disabled deposits cannot be poked.
        vm.expectRevert(ISweepRegistry.DepositNotActive.selector);
        registry.pokeSweep(token, deposit);
    }

    /*//////////////////////////////////////////////////////////////
                                 Batch
    //////////////////////////////////////////////////////////////*/

    function test_batch_register() public {
        uint256 pkB = 0xB0B0;
        address depositB = vm.addr(pkB);
        uint64 deadline = _deadline();

        ISweepRegistry.SweepRegistration[]
            memory regs = new ISweepRegistry.SweepRegistration[](2);
        regs[0] = ISweepRegistry.SweepRegistration({
            deposit: deposit,
            master: master,
            nonce: 0,
            deadline: deadline,
            depositSignature: _sign(depositPk, deposit, master, 0, deadline)
        });
        regs[1] = ISweepRegistry.SweepRegistration({
            deposit: depositB,
            master: master,
            nonce: 0,
            deadline: deadline,
            depositSignature: _sign(pkB, depositB, master, 0, deadline)
        });

        vm.prank(master);
        registry.registerSweeps(regs);

        assertEq(registry.resolveSweep(token, deposit), master);
        assertEq(registry.resolveSweep(token, depositB), master);
    }

    /*//////////////////////////////////////////////////////////////
                                 Pause
    //////////////////////////////////////////////////////////////*/

    function test_pause_blocks_mutations_but_not_resolve() public {
        _register(master);
        vm.prank(owner);
        registry.setPause(true);

        // resolveSweep (view) is unaffected by pause.
        assertEq(registry.resolveSweep(token, deposit), master);

        // Mutating entrypoints revert while paused.
        vm.expectRevert("Pausable: paused");
        registry.pokeSweep(token, deposit);

        uint64 deadline = _deadline();
        uint256 pkB = 0xB0B0;
        address depositB = vm.addr(pkB);
        bytes memory sig = _sign(pkB, depositB, master, 0, deadline);
        vm.expectRevert("Pausable: paused");
        vm.prank(master);
        registry.registerSweep(depositB, master, 0, deadline, sig);

        vm.expectRevert("Pausable: paused");
        vm.prank(master);
        registry.disableSweep(deposit);
    }

    function test_setPause_onlyOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(stranger);
        registry.setPause(true);
    }
}
