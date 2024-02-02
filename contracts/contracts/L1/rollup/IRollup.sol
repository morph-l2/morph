// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IRollup {
    /**********
     * Events *
     **********/

    /// @notice Emitted when a new batch is committed.
    /// @param batchIndex The index of the batch.
    /// @param batchHash The hash of the batch.
    event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);

    /// @notice revert a pending batch.
    /// @param batchIndex The index of the batch.
    /// @param batchHash The hash of the batch
    event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);

    /// @notice Emitted when a batch is finalized.
    /// @param batchIndex The index of the batch.
    /// @param batchHash The hash of the batch
    /// @param stateRoot The state root on layer 2 after this batch.
    /// @param withdrawRoot The merkle root on layer2 after this batch.
    event FinalizeBatch(
        uint256 indexed batchIndex,
        bytes32 indexed batchHash,
        bytes32 stateRoot,
        bytes32 withdrawRoot
    );

    /// @notice Emitted when owner updates the status of sequencer.
    /// @param account The address of account updated.
    /// @param status The status of the account updated.
    event UpdateSequencer(address indexed account, bool status);

    /// @notice Emitted when owner updates the status of prover.
    /// @param account The address of account updated.
    /// @param status The status of the account updated.
    event UpdateProver(address indexed account, bool status);

    /// @notice Emitted when owner updates the status of prover.
    /// @param account The address of account updated.
    /// @param status The status of the account updated.
    event UpdateChallenger(address indexed account, bool status);

    /// @notice Emitted when the address of rollup verifier is updated.
    /// @param oldVerifier The address of old rollup verifier.
    /// @param newVerifier The address of new rollup verifier.
    event UpdateVerifier(
        address indexed oldVerifier,
        address indexed newVerifier
    );

    /// @notice Emitted when the value of `maxNumTxInChunk` is updated.
    /// @param oldMaxNumTxInChunk The old value of `maxNumTxInChunk`.
    /// @param newMaxNumTxInChunk The new value of `maxNumTxInChunk`.
    event UpdateMaxNumTxInChunk(
        uint256 oldMaxNumTxInChunk,
        uint256 newMaxNumTxInChunk
    );

    /// @notice Emitted when the state of Chanllenge is updated.
    /// @param batchIndex The index of the batch.
    /// @param challenger The address of challenger.
    /// @param challengeDeposit The deposit of challenger.
    event ChallengeState(
        uint64 indexed batchIndex,
        address challenger,
        uint256 challengeDeposit
    );

    /// @notice Emitted when the result of Chanllenge is updated.
    /// @param batchIndex The index of the batch.
    /// @param winner  The address of winner.
    /// @param res The result of challenge.
    event ChallengeRes(uint64 indexed batchIndex, address winner, string res);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice The latest finalized batch index.
    function lastFinalizedBatchIndex() external view returns (uint256);

    /// @notice The latest finalized batch index.
    function lastCommittedBatchIndex() external view returns (uint256);

    /// @notice Return the batch hash of a committed batch.
    /// @param batchIndex The index of the batch.
    function committedBatches(
        uint256 batchIndex
    ) external view returns (bytes32);

    /// @notice Return the state root of a committed batch.
    /// @param batchIndex The index of the batch.
    function finalizedStateRoots(
        uint256 batchIndex
    ) external view returns (bytes32);

    /// @notice Return the the committed batch of withdrawalRoot.
    /// @param withdrawalRoot The withdrawal root.
    function withdrawalRoots(
        bytes32 withdrawalRoot
    ) external view returns (uint256);

    /// @notice Return whether the batch is finalized by batch index.
    /// @param batchIndex The index of the batch.
    function isBatchFinalized(uint256 batchIndex) external view returns (bool);

    /// @notice Return the rollup config of FINALIZATION_PERIOD_SECONDS.
    function FINALIZATION_PERIOD_SECONDS() external view returns (uint256);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Commit a batch of transactions on layer 1.
    ///
    /// @param batchData The BatchData struct
    /// @param minGasLimit The min gas limit.
    function commitBatch(
        BatchData calldata batchData,
        uint32 minGasLimit
    ) external payable;

    /// @notice Revert a pending batch.
    /// @dev one can only revert unfinalized batches.
    /// @param batchHeader The header of current batch, see the encoding in comments of `commitBatch`.
    /// @param count The number of subsequent batches to revert, including current batch.
    function revertBatch(bytes calldata batchHeader, uint256 count) external;

    /// @param version The version of current batch.
    /// @param parentBatchHeader The header of parent batch, see the comments of `BatchHeaderV0Codec`.
    /// @param chunks The list of encoded chunks, see the comments of `ChunkCodec`.
    /// @param skippedL1MessageBitmap The bitmap indicates whether each L1 message is skipped or not.
    /// @param prevStateRoot The state root of parent batch.
    /// @param postStateRoot The state root of current batch.
    /// @param withdrawalRoot The withdraw trie root of current batch.
    /// @param signature The signature of current batch.
    struct BatchData {
        uint8 version;
        bytes parentBatchHeader;
        bytes[] chunks;
        bytes skippedL1MessageBitmap;
        bytes32 prevStateRoot;
        bytes32 postStateRoot;
        bytes32 withdrawalRoot;
        BatchSignature signature;
    }

    /// @param version the version of staking contract
    /// @param signers The index list of signers of current batch.
    /// @param signature The bls signature.
    struct BatchSignature {
        uint256 version;
        uint256[] signers;
        bytes signature;
    }
}
