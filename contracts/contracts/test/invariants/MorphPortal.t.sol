pragma solidity =0.8.16;

import {Portal_Initializer} from "../CommonTest.t.sol";
import {Types} from "../../libraries/Types.sol";
import {Rollup} from "../../L1/Rollup.sol";
import {IRollup} from "../../L1/IRollup.sol";

contract MorphPortal_Invariant_Harness is Portal_Initializer {
    // Reusable default values for a test withdrawal
    Types.WithdrawalTransaction _defaultTx;
    uint256 _proposedOutputIndex;
    uint256 _proposedBlockNumber;
    bytes32 _withdrawalHash;
    bytes32[_TREE_DEPTH] _withdrawalProof;
    bytes32 _withdrawalRoot;

    function setUp() public virtual override {
        super.setUp();
        vm.prank(multisig);
        rollup.addSequencer(address(0));
        vm.deal(address(0), 5 * MIN_DEPOSIT);
        vm.prank(address(0));
        rollup.stake{value: MIN_DEPOSIT}();

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

        bytes memory batchHeader0 = new bytes(89);
        bytes32 stateRoot = bytes32(uint256(1));

        // import genesis batch first
        assembly {
            mstore(add(batchHeader0, add(0x20, 25)), 1)
        }
        rollup.importGenesisBatch(
            batchHeader0,
            bytes32(uint256(1)),
            getTreeRoot()
        );

        bytes[] memory chunks = new bytes[](1);
        bytes memory chunk0;
        chunk0 = new bytes(1 + 60);
        chunk0[0] = bytes1(uint8(1)); // one block in this chunk
        chunks[0] = chunk0;
        IRollup.BatchData memory batchData = IRollup.BatchData(
            0,
            batchHeader0,
            chunks,
            new bytes(0),
            stateRoot,
            stateRoot,
            _withdrawalRoot,
            nilBatchSig
        );
        // update portal l1Messager to caller
        vm.startPrank(address(0));
        rollup.commitBatch(batchData, minGasLimit);
        vm.stopPrank();

        // Warp beyond the finalization period for the block we've proposed.
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);

        // Fund the portal so that we can withdraw ETH.
        vm.deal(address(portal), 0xFFFFFFFF);
    }
}

contract MorphPortal_CannotTimeTravel is MorphPortal_Invariant_Harness {
    function setUp() public override {
        super.setUp();
        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    /**
     * @custom:invariant `finalizeWithdrawalTransaction` should revert if the finalization
     * period has not elapsed.
     *
     * A withdrawal that has been proven should not be able to be finalized until after
     * the finalization period has elapsed.
     */
    function invariant_cannotFinalizeBeforePeriodHasPassed() external {
        vm.expectRevert(
            "MorphPortal: proven withdrawal finalization period has not elapsed"
        );
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }
}

contract MorphPortal_CannotFinalizeTwice is MorphPortal_Invariant_Harness {
    function setUp() public override {
        super.setUp();
        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
        // Warp past the finalization period.
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        // Finalize the withdrawal transaction.
        portal.finalizeWithdrawalTransaction(_defaultTx);
        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    /**
     * @custom:invariant `finalizeWithdrawalTransaction` should revert if the withdrawal
     * has already been finalized.
     *
     * Ensures that there is no chain of calls that can be made that allows a withdrawal
     * to be finalized twice.
     */
    function invariant_cannotFinalizeTwice() external {
        vm.expectRevert("MorphPortal: withdrawal has already been finalized");
        portal.finalizeWithdrawalTransaction(_defaultTx);
    }
}

contract MorphPortal_CanAlwaysFinalizeAfterWindow is
    MorphPortal_Invariant_Harness
{
    function setUp() public override {
        super.setUp();
        // Prove the withdrawal transaction
        portal.proveWithdrawalTransaction(
            _defaultTx,
            _withdrawalProof,
            _withdrawalRoot
        );
        vm.warp(block.timestamp + rollup.FINALIZATION_PERIOD_SECONDS() + 1);
        rollup.finalizeBatchs();
        // Set the target contract to the portal proxy
        targetContract(address(portal));
        // Exclude the proxy multisig from the senders so that the proxy cannot be upgraded
        excludeSender(address(multisig));
    }

    function invariant_canAlwaysFinalize() external {
        uint256 bobBalanceBefore = address(bob).balance;
        portal.finalizeWithdrawalTransaction(_defaultTx);
        assertEq(address(bob).balance, bobBalanceBefore + _defaultTx.value);
    }
}
