// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {stdError} from "forge-std/Test.sol";
import {Portal_Initializer, CommonTest, NextImpl} from "./CommonTest.t.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {MorphPortal} from "../L1/MorphPortal.sol";
import {Types} from "../libraries/Types.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {Proxy} from "../universal/Proxy.sol";
import {ResourceMetering} from "../L1/ResourceMetering.sol";
import {Rollup} from "../L1/Rollup.sol";
import {IRollup} from "../L1/IRollup.sol";
import "forge-std/console.sol";

contract MorphPortal_Test is Portal_Initializer {
    event Paused(address);
    event Unpaused(address);

    function test_constructor_succeeds() external {
        assertEq(address(portal.ROLLUP()), address(rollup));
        assertEq(portal.l2Sender(), 0x000000000000000000000000000000000000dEaD);
        assertEq(portal.paused(), false);
    }

    /**
     * @notice The MorphPortal can be paused by the GUARDIAN
     */
    function test_pause_succeeds() external {
        address guardian = portal.GUARDIAN();

        assertEq(portal.paused(), false);

        vm.expectEmit(true, true, true, true, address(portal));
        emit Paused(guardian);

        vm.prank(guardian);
        portal.pause();

        assertEq(portal.paused(), true);
    }

    /**
     * @notice The MorphPortal reverts when an account that is not the
     *         GUARDIAN calls `pause()`
     */
    function test_pause_onlyGuardian_reverts() external {
        assertEq(portal.paused(), false);

        assertTrue(portal.GUARDIAN() != alice);
        vm.expectRevert("MorphPortal: only guardian can pause");
        vm.prank(alice);
        portal.pause();

        assertEq(portal.paused(), false);
    }

    /**
     * @notice The MorphPortal can be unpaused by the GUARDIAN
     */
    function test_unpause_succeeds() external {
        address guardian = portal.GUARDIAN();

        vm.prank(guardian);
        portal.pause();
        assertEq(portal.paused(), true);

        vm.expectEmit(true, true, true, true, address(portal));
        emit Unpaused(guardian);
        vm.prank(guardian);
        portal.unpause();

        assertEq(portal.paused(), false);
    }

    /**
     * @notice The MorphPortal reverts when an account that is not
     *         the GUARDIAN calls `unpause()`
     */
    function test_unpause_onlyGuardian_reverts() external {
        address guardian = portal.GUARDIAN();

        vm.prank(guardian);
        portal.pause();
        assertEq(portal.paused(), true);

        assertTrue(portal.GUARDIAN() != alice);
        vm.expectRevert("MorphPortal: only guardian can unpause");
        vm.prank(alice);
        portal.unpause();

        assertEq(portal.paused(), true);
    }

    // Test: depositTransaction fails when contract creation has a non-zero destination address
    function test_depositTransaction_contractCreation_reverts() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // contract creation must have a target of address(0)
        vm.expectRevert(
            "MorphPortal: must send to address(0) when creating a contract"
        );
        portal.depositTransaction(address(1), 1, 0, true, hex"");
    }

    /**
     * @notice Prevent deposits from being too large to have a sane upper bound
     *         on unsafe blocks sent over the p2p network.
     */
    function test_depositTransaction_largeData_reverts() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        uint256 size = 120_001;
        uint64 gasLimit = portal.minimumGasLimit(uint64(size));
        vm.expectRevert("MorphPortal: data too large");
        portal.depositTransaction({
            _to: address(0),
            _value: 0,
            _gasLimit: gasLimit,
            _isCreation: false,
            _data: new bytes(size)
        });
    }

    /**
     * @notice Prevent gasless deposits from being force processed in L2 by
     *         ensuring that they have a large enough gas limit set.
     */
    function test_depositTransaction_smallGasLimit_reverts() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.expectRevert("MorphPortal: gas limit too small");
        portal.depositTransaction({
            _to: address(1),
            _value: 0,
            _gasLimit: 0,
            _isCreation: false,
            _data: hex""
        });
    }

    /**
     * @notice Fuzz for too small of gas limits
     */
    function testFuzz_depositTransaction_smallGasLimit_succeeds(
        bytes memory _data,
        bool _shouldFail
    ) external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.assume(_data.length <= type(uint64).max);

        uint64 gasLimit = portal.minimumGasLimit(uint64(_data.length));
        if (_shouldFail) {
            gasLimit = uint64(bound(gasLimit, 0, gasLimit - 1));
            vm.expectRevert("MorphPortal: gas limit too small");
        }
        portal.depositTransaction({
            _to: address(0x40),
            _value: 0,
            _gasLimit: gasLimit,
            _isCreation: false,
            _data: _data
        });
    }

    /**
     * @notice Ensure that the 0 calldata case is covered and there is a linearly
     *         increasing gas limit for larger calldata sizes.
     */
    function test_minimumGasLimit_succeeds() external {
        assertEq(portal.minimumGasLimit(0), 21_000);
        assertTrue(portal.minimumGasLimit(2) > portal.minimumGasLimit(1));
        assertTrue(portal.minimumGasLimit(3) > portal.minimumGasLimit(2));
    }

    // Test: depositTransaction should emit the correct log when an EOA deposits a tx with 0 value
    function test_depositTransaction_noValueEOA_succeeds() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // EOA emulation
        vm.prank(address(this), address(this));
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should emit the correct log when a contract deposits a tx with 0 value
    function test_depositTransaction_noValueContract_succeeds() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should emit the correct log when an EOA deposits a contract creation with 0 value
    function test_depositTransaction_createWithZeroValueForEOA_succeeds()
        external
    {
        bytes32 l1MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l1MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );

        portal.depositTransaction(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should emit the correct log when a contract deposits a contract creation with 0 value
    function test_depositTransaction_createWithZeroValueForContract_succeeds()
        external
    {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );
        portal.depositTransaction(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should increase its eth balance when an EOA deposits a transaction with ETH
    function test_depositTransaction_withEthValueFromEOA_succeeds() external {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );

        portal.depositTransaction{value: NON_ZERO_VALUE}(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }

    // Test: depositTransaction should increase its eth balance when a contract deposits a transaction with ETH
    function test_depositTransaction_withEthValueFromContract_succeeds()
        external
    {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            NON_ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
        // vm.deal(address(L1Messenger), 5 ether);
        // vm.prank(address(L1Messenger));
        portal.depositTransaction{value: NON_ZERO_VALUE}(
            NON_ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            false,
            NON_ZERO_DATA
        );
    }

    // Test: depositTransaction should increase its eth balance when an EOA deposits a contract creation with ETH
    function test_depositTransaction_withEthValueAndEOAContractCreation_succeeds()
        external
    {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // EOA emulation
        vm.prank(address(this), address(this));

        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            hex""
        );

        portal.depositTransaction{value: NON_ZERO_VALUE}(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            hex""
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }

    // Test: depositTransaction should increase its eth balance when a contract deposits a contract creation with ETH
    function test_depositTransaction_withEthValueAndContractContractCreation_succeeds()
        external
    {
        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        vm.expectEmit(true, true, false, true);
        emitTransactionDeposited(
            AddressAliasHelper.applyL1ToL2Alias(address(this)),
            ZERO_ADDRESS,
            NON_ZERO_VALUE,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );

        portal.depositTransaction{value: NON_ZERO_VALUE}(
            ZERO_ADDRESS,
            ZERO_VALUE,
            NON_ZERO_GASLIMIT,
            true,
            NON_ZERO_DATA
        );
        assertEq(address(portal).balance, NON_ZERO_VALUE);
    }
}

contract MorphPortal_FinalizeWithdrawal_Test is Portal_Initializer {
    // Reusable default values for a test withdrawal
    Types.WithdrawalTransaction _defaultTx;

    bytes32 _withdrawalHash;
    bytes32[_TREE_DEPTH] _withdrawalProof;
    bytes32 _withdrawalRoot;
    bytes32 stateRoot;
    bytes batchHeader1Store;

    address public caller = address(0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84);
    bytes32 public l1MessagerIndex = bytes32(uint256(55));

    // Use a constructor to set the storage vars above, so as to minimize the number of ffi calls.
    constructor() {
        super.setUp();
        _defaultTx = Types.WithdrawalTransaction({
            nonce: 0,
            sender: alice,
            target: bob,
            value: 100,
            gasLimit: 100_000,
            data: hex""
        });
        // Get withdrawal proof data we can use for testing.
        (_withdrawalHash, _withdrawalProof, _withdrawalRoot) = ffi
            .getProveWithdrawalTransactionInputs(_defaultTx);
    }

    // Get the system into a nice ready-to-use state.
    function setUp() public virtual override {
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();

        // update portal l1Messager to caller
        vm.store(
            address(portal),
            bytes32(l1MessagerIndex),
            bytes32(abi.encode(caller))
        );
        vm.deal(caller, 5 * MIN_DEPOSIT);
        vm.startPrank(caller);

        bytes memory batchHeader0 = new bytes(89);
        stateRoot = bytes32(uint256(1));

        // import 10 L1 messages
        for (uint256 i = 0; i < 10; i++) {
            portal.depositTransaction(caller, 0, 1000000, false, new bytes(0));
        }
        vm.stopPrank();
        vm.store(
            address(portal),
            bytes32(l1MessagerIndex),
            bytes32(abi.encode(address(L1Messenger)))
        );
        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );

        bytes32 batchHash0 = rollup.committedBatches(0);
        bytes memory bitmap;
        bytes[] memory chunks;
        bytes memory chunk0;
        bytes memory chunk1;

        // commit batch1, one chunk with one block, 1 tx, 1 L1 message, no skip
        // => payload for data hash of chunk0
        //   0000000000000000
        //   0000000000000000
        //   0000000000000000000000000000000000000000000000000000000000000000
        //   0000000000000000
        //   0001
        //   a2277fd30bbbe74323309023b56035b376d7768ad237ae4fc46ead7dc9591ae1
        // => data hash for chunk0
        //   9ef1e5694bdb014a1eea42be756a8f63bfd8781d6332e9ef3b5126d90c62f110
        // => data hash for all chunks
        //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
        // => payload for batch header
        //   00
        //   0000000000000001
        //   0000000000000001
        //   0000000000000001
        //   d9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
        //   119b828c2a2798d2c957228ebeaff7e10bb099ae0d4e224f3eeb779ff61cba61
        //   0000000000000000000000000000000000000000000000000000000000000000
        // => hash for batch header
        //   00847173b29b238cf319cde79512b7c213e5a8b4138daa7051914c4592b6dfc7
        bytes memory batchHeader1 = new bytes(89 + 32);
        assembly {
            mstore(add(batchHeader1, 0x20), 0) // version
            mstore(add(batchHeader1, add(0x20, 1)), shl(192, 1)) // batchIndex = 1
            mstore(add(batchHeader1, add(0x20, 9)), shl(192, 1)) // l1MessagePopped = 1
            mstore(add(batchHeader1, add(0x20, 17)), shl(192, 1)) // totalL1MessagePopped = 1
            mstore(
                add(batchHeader1, add(0x20, 25)),
                0xd9cb6bf9264006fcea490d5c261f7453ab95b1b26033a3805996791b8e3a62f3
            ) // dataHash
            mstore(add(batchHeader1, add(0x20, 57)), batchHash0) // parentBatchHash
            mstore(add(batchHeader1, add(0x20, 89)), 0) // bitmap0
        }
        chunk0 = new bytes(1 + 60);
        assembly {
            mstore(add(chunk0, 0x20), shl(248, 1)) // numBlocks = 1
            mstore(add(chunk0, add(0x21, 56)), shl(240, 1)) // numTransactions = 1
            mstore(add(chunk0, add(0x21, 58)), shl(240, 1)) // numL1Messages = 1
        }
        chunks = new bytes[](1);
        chunks[0] = chunk0;
        bitmap = new bytes(32);
        batchHeader1Store = batchHeader1;
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            bitmap,
            stateRoot,
            stateRoot,
            _withdrawalRoot,
            nilBatchSig
        );
        // update portal l1Messager to caller
        vm.startPrank(address(0));
        rollup.commitBatch(batchData, defultGasLimit);
        vm.stopPrank();
        // Warp beyond the finalization period for the block we've proposed.
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        // Fund the portal so that we can withdraw ETH.
        vm.deal(address(portal), 0xFFFFFFFF);
    }

    // Utility function used in the subsequent test. This is necessary to assert that the
    // reentrant call will revert.
    function callPortalAndExpectRevert() external payable {
        vm.expectRevert(
            "MorphPortal: can only trigger one withdrawal per transaction"
        );
        // Arguments here don't matter, as the require check is the first thing that happens.
        // We assume that this has already been proven.
        portal.finalizeWithdrawalTransaction(_defaultTx);
        // Assert that the withdrawal was not finalized.
        assertFalse(
            portal.finalizedWithdrawals(Hashing.hashWithdrawal(_defaultTx))
        );
    }

    /**
     * @notice Proving withdrawal transactions should revert when paused
     */
    function test_proveWithdrawalTransaction_paused_reverts() external {
        vm.prank(portal.GUARDIAN());
        portal.pause();

        vm.expectRevert("MorphPortal: paused");
        portal.proveWithdrawalTransaction({
            _tx: _defaultTx,
            _withdrawalProof: _withdrawalProof,
            _withdrawalRoot: _withdrawalRoot
        });
    }

    // Test: proveWithdrawalTransaction cannot prove a withdrawal with itself (the MorphPortal) as the target.
    function test_proveWithdrawalTransaction_onSelfCall_reverts() external {
        _defaultTx.target = address(portal);
        vm.expectRevert(
            "MorphPortal: you cannot send messages to the portal contract"
        );
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }

    // Test: proveWithdrawalTransaction reverts if the outputRootProof does not match the output root
    function test_proveWithdrawalTransaction_onInvalidOutputRootProof_reverts()
        external
    {
        _withdrawalProof[1] = bytes32(0);
        // Modify the version to invalidate the withdrawal proof.
        vm.expectRevert("MorphPortal: invalid withdrawal inclusion proof");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }

    // Test: proveWithdrawalTransaction reverts if the proof is invalid due to non-existence of
    // the withdrawal.
    function test_proveWithdrawalTransaction_onInvalidWithdrawalProof_reverts()
        external
    {
        // modify the default test values to invalidate the proof.
        _defaultTx.data = hex"abcd";
        vm.expectRevert("MorphPortal: invalid withdrawal inclusion proof");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }

    // Test: proveWithdrawalTransaction reverts if the passed transaction's withdrawalHash has
    // already been proven.
    function test_proveWithdrawalTransaction_replayProve_reverts() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );

        vm.expectRevert("MorphPortal: withdrawal hash has already been proven");
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }

    // Test: proveWithdrawalTransaction succeeds and emits the WithdrawalProven event.
    function test_proveWithdrawalTransaction_validWithdrawalProof_succeeds()
        external
    {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
    }

    // Test: finalizeWithdrawalTransaction succeeds and emits the WithdrawalFinalized event.
    function test_finalizeWithdrawalTransaction_provenWithdrawalHash_succeeds()
        external
    {
        uint256 bobBalanceBefore = address(bob).balance;

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );

        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        vm.expectEmit(true, true, false, true);
        emit WithdrawalFinalized(_withdrawalHash, true);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore + 100);
    }

    /**
     * @notice Finalizing withdrawal transactions should revert when paused
     */
    function test_finalizeWithdrawalTransaction_paused_reverts() external {
        vm.prank(portal.GUARDIAN());
        portal.pause();

        vm.expectRevert("MorphPortal: paused");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if the withdrawal has not been proven.
    function test_finalizeWithdrawalTransaction_ifWithdrawalNotProven_reverts()
        external
    {
        uint256 bobBalanceBefore = address(bob).balance;

        vm.expectRevert("MorphPortal: withdrawal has not been proven yet");
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore);
    }

    // Test: finalizeWithdrawalTransaction fails because the target reverts,
    // and emits the WithdrawalFinalized event with success=false.
    function test_finalizeWithdrawalTransaction_targetFails_fails() external {
        uint256 bobBalanceBefore = address(bob).balance;
        vm.etch(bob, hex"fe"); // Contract with just the invalid opcode.

        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );

        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        // rollup.confirmBatchs();
        vm.expectEmit(true, true, true, true);
        emit WithdrawalFinalized(_withdrawalHash, false);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        assert(address(bob).balance == bobBalanceBefore);
    }

    // Test: finalizeWithdrawalTransaction reverts if the finalization period has not yet passed.
    function test_finalizeWithdrawalTransaction_onRecentWithdrawal_reverts()
        external
    {
        // Setup the Oracle to return an output with a recent timestamp
        uint256 recentTimestamp = 0;
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
        vm.warp(recentTimestamp);
        vm.expectRevert(
            "MorphPortal: proven withdrawal finalization period has not elapsed"
        );
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if the withdrawal has already been finalized.
    function test_finalizeWithdrawalTransaction_onReplay_reverts() external {
        vm.expectEmit(true, true, true, true);
        emit WithdrawalProven(_withdrawalHash, alice, bob);
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );

        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        vm.expectEmit(true, true, true, true);
        emit WithdrawalFinalized(_withdrawalHash, true);
        portal.finalizeWithdrawalTransaction(_defaultTx);

        vm.expectRevert("MorphPortal: withdrawal has already been finalized");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }

    // Test: finalizeWithdrawalTransaction reverts if insufficient gas is supplied.
    function test_finalizeWithdrawalTransaction_onInsufficientGas_reverts()
        external
    {
        // This number was identified through trial and error.
        uint256 gasLimit = 150_000;
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );

        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        vm.expectRevert("SafeCall: Not enough gas");
        portal.finalizeWithdrawalTransaction{gas: gasLimit}(_defaultTx);
    }

    function testDiff_finalizeWithdrawalTransaction_succeeds(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        vm.assume(
            _target != address(portal) && // Cannot call the Morph portal or a contract
                _target.code.length == 0 && // No accounts with code
                _target != CONSOLE && // The console has no code but behaves like a contract
                uint160(_target) > 9 // No precompiles (or zero address)
        );

        // Total ETH supply is currently about 120M ETH.
        uint256 value = bound(_value, 0, 200_000_000 ether);
        vm.deal(address(portal), value);

        uint256 gasLimit = bound(_gasLimit, 0, 50_000_000);
        uint256 nonce = messagePasser.messageNonce();

        // Get a withdrawal transaction and mock proof from the differential testing script.
        Types.WithdrawalTransaction memory _tx = Types.WithdrawalTransaction({
            nonce: nonce,
            sender: _sender,
            target: _target,
            value: value,
            gasLimit: gasLimit,
            data: _data
        });
        (
            bytes32 withdrawalHash,
            bytes32[_TREE_DEPTH] memory withdrawalProof,
            bytes32 withdrawalRoot
        ) = ffi.getProveWithdrawalTransactionInputs(_tx);

        bytes[] memory chunks = new bytes[](1);
        bytes memory chunk0;
        chunk0 = new bytes(1 + 60);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunks[0] = chunk0;
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader1Store,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            withdrawalRoot,
            nilBatchSig
        );
        vm.startPrank(address(0));
        rollup.commitBatch(batchData, defultGasLimit);
        vm.stopPrank();
        // Ensure the values returned from ffi are correct
        assertEq(withdrawalHash, Hashing.hashWithdrawal(_tx));

        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(_tx, withdrawalProof, withdrawalRoot);
        (bytes32 _root, , ) = portal.provenWithdrawals(withdrawalHash);
        assertTrue(_root != bytes32(0));

        // Warp past the finalization period
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        // Finalize the withdrawal transaction
        vm.expectCallMinGas(
            _tx.target,
            _tx.value,
            uint64(_tx.gasLimit),
            _tx.data
        );
        portal.finalizeWithdrawalTransaction(_tx);
        assertTrue(portal.finalizedWithdrawals(withdrawalHash));
    }
}

contract MorphPortalUpgradeable_Test is Portal_Initializer {
    Proxy internal proxy;
    uint64 initialBlockNum;

    function setUp() public override {
        super.setUp();
        initialBlockNum = uint64(block.number);
        proxy = Proxy(payable(address(portal)));
    }

    function test_params_initValuesOnProxy_succeeds() external {
        MorphPortal p = MorphPortal(payable(address(proxy)));

        (uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum) = p
            .params();

        ResourceMetering.ResourceConfig memory rcfg = systemConfig
            .resourceConfig();
        assertEq(prevBaseFee, rcfg.minimumBaseFee);
        assertEq(prevBoughtGas, 0);
        assertEq(prevBlockNum, initialBlockNum);
    }

    function test_initialize_cannotInitProxy_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        MorphPortal(payable(proxy)).initialize(false, address(L1Messenger));
    }

    function test_initialize_cannotInitImpl_reverts() external {
        vm.expectRevert("Initializable: contract is already initialized");
        MorphPortal(portalImpl).initialize(false, address(L1Messenger));
    }

    function test_upgradeToAndCall_upgrading_succeeds() external {
        // Check an unused slot before upgrading.
        bytes32 slot21Before = vm.load(address(portal), bytes32(uint256(21)));
        assertEq(bytes32(0), slot21Before);

        NextImpl nextImpl = new NextImpl();
        vm.startPrank(multisig);
        proxy.upgradeToAndCall(
            address(nextImpl),
            abi.encodeWithSelector(NextImpl.initialize.selector)
        );
        assertEq(proxy.implementation(), address(nextImpl));

        // Verify that the NextImpl contract initialized its values according as expected
        bytes32 slot21After = vm.load(address(portal), bytes32(uint256(21)));
        bytes32 slot21Expected = NextImpl(address(portal)).slot21Init();
        assertEq(slot21Expected, slot21After);
    }
}

/**
 * @title MorphPortalResourceFuzz_Test
 * @dev Test various values of the resource metering config to ensure that deposits cannot be
 *         broken by changing the config.
 */
contract MorphPortalResourceFuzz_Test is Portal_Initializer {
    /**
     * @dev The max gas limit observed throughout this test. Setting this too high can cause
     *      the test to take too long to run.
     */
    uint256 constant MAX_GAS_LIMIT = 30_000_000;

    /**
     * @dev Test that various values of the resource metering config will not break deposits.
     */
    function testFuzz_systemConfigDeposit_succeeds(
        uint32 _maxResourceLimit,
        uint8 _elasticityMultiplier,
        uint8 _baseFeeMaxChangeDenominator,
        uint32 _minimumBaseFee,
        uint32 _systemTxMaxGas,
        uint128 _maximumBaseFee,
        uint64 _gasLimit,
        uint64 _prevBoughtGas,
        uint128 _prevBaseFee,
        uint8 _blockDiff
    ) external {
        // Get the set system gas limit
        uint64 gasLimit = systemConfig.gasLimit();
        // Bound resource config
        _maxResourceLimit = uint32(
            bound(_maxResourceLimit, 21000, MAX_GAS_LIMIT / 8)
        );
        _gasLimit = uint64(bound(_gasLimit, 21000, _maxResourceLimit));
        _prevBaseFee = uint128(bound(_prevBaseFee, 0, 3 gwei));
        // Prevent values that would cause reverts
        vm.assume(gasLimit >= _gasLimit);
        vm.assume(_minimumBaseFee < _maximumBaseFee);
        vm.assume(_baseFeeMaxChangeDenominator > 1);
        vm.assume(
            uint256(_maxResourceLimit) + uint256(_systemTxMaxGas) <= gasLimit
        );
        vm.assume(_elasticityMultiplier > 0);
        vm.assume(
            ((_maxResourceLimit / _elasticityMultiplier) *
                _elasticityMultiplier) == _maxResourceLimit
        );
        _prevBoughtGas = uint64(
            bound(_prevBoughtGas, 0, _maxResourceLimit - _gasLimit)
        );
        _blockDiff = uint8(bound(_blockDiff, 0, 1));

        // Create a resource config to mock the call to the system config with
        ResourceMetering.ResourceConfig memory rcfg = ResourceMetering
            .ResourceConfig({
                maxResourceLimit: _maxResourceLimit,
                elasticityMultiplier: _elasticityMultiplier,
                baseFeeMaxChangeDenominator: _baseFeeMaxChangeDenominator,
                minimumBaseFee: _minimumBaseFee,
                systemTxMaxGas: _systemTxMaxGas,
                maximumBaseFee: _maximumBaseFee
            });
        vm.mockCall(
            address(systemConfig),
            abi.encodeWithSelector(systemConfig.resourceConfig.selector),
            abi.encode(rcfg)
        );

        // Set the resource params
        uint256 _prevBlockNum = block.number - _blockDiff;
        vm.store(
            address(portal),
            bytes32(uint256(1)),
            bytes32(
                (_prevBlockNum << 192) |
                    (uint256(_prevBoughtGas) << 128) |
                    _prevBaseFee
            )
        );
        // Ensure that the storage setting is correct
        (
            uint128 prevBaseFee,
            uint64 prevBoughtGas,
            uint64 prevBlockNum
        ) = portal.params();
        assertEq(prevBaseFee, _prevBaseFee);
        assertEq(prevBoughtGas, _prevBoughtGas);
        assertEq(prevBlockNum, _prevBlockNum);

        bytes32 l2MessagerIndex = bytes32(uint256(55));
        vm.store(
            address(portal),
            bytes32(l2MessagerIndex),
            bytes32(abi.encode(address(this)))
        );
        // Do a deposit, should not revert
        portal.depositTransaction{gas: MAX_GAS_LIMIT}({
            _to: address(0x20),
            _value: 0x40,
            _gasLimit: _gasLimit,
            _isCreation: false,
            _data: hex""
        });
    }
}
