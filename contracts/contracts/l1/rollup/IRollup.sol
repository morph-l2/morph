// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IRollup {
    /***********
     * Structs *
     ***********/

    /// @param version                  The version of current batch.
    /// @param parentBatchHeader        The header of parent batch, see the comments of `BatchHeaderV0Codec`.
    /// @param lastBlockNumber          The last block number in this batch
    /// @param numL1Messages            The number of L1 messages in this batch
    /// @param prevStateRoot            The state root of parent batch.
    /// @param postStateRoot            The state root of current batch.
    /// @param withdrawalRoot           The withdraw trie root of current batch.
    struct BatchDataInput {
        uint8 version;
        bytes parentBatchHeader;
        uint64 lastBlockNumber;
        uint16 numL1Messages;
        bytes32 prevStateRoot;
        bytes32 postStateRoot;
        bytes32 withdrawalRoot;
    }

    /// @param signedSequencers The bitmap of signed sequencers
    /// @param sequencerSets    The latest 3 sequencer sets
    /// @param signature        The BLS signature
    struct BatchSignatureInput {
        uint256 signedSequencersBitmap;
        bytes sequencerSets;
        bytes signature;
    }

    /// @param originTimestamp
    /// @param finalizeTimestamp
    /// @param blockNumber
    struct BatchData {
        uint256 originTimestamp;
        uint256 finalizeTimestamp;
        uint256 blockNumber;
        uint256 signedSequencersBitmap;
    }

    /// @dev Structure to store information about a batch challenge.
    /// @param batchIndex The index of the challenged batch.
    /// @param challenger The address of the challenger.
    /// @param challengeDeposit The amount of deposit put up by the challenger.
    /// @param startTime The timestamp when the challenge started.
    /// @param challengeSuccess Flag indicating whether the challenge was successful.
    /// @param finished Flag indicating whether the challenge has been resolved.
    struct BatchChallenge {
        uint64 batchIndex;
        address challenger;
        uint256 challengeDeposit;
        uint256 startTime;
        bool challengeSuccess;
        bool finished;
    }

    /// @param receiver
    /// @param amount
    struct BatchChallengeReward {
        address receiver;
        uint256 amount;
    }

    /***********
     * Errors *
     ***********/

    /// @notice error zero address
    error ErrZeroAddress();

    /**********
     * Events *
     **********/

    /// @notice Emitted when a new batch is committed.
    /// @param batchIndex   The index of the batch.
    /// @param batchHash    The hash of the batch.
    event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);

    /// @notice revert a pending batch.
    /// @param batchIndex   The index of the batch.
    /// @param batchHash    The hash of the batch
    event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);

    /// @notice Emitted when a batch is finalized.
    /// @param batchIndex   The index of the batch.
    /// @param batchHash    The hash of the batch
    /// @param stateRoot    The state root on layer 2 after this batch.
    /// @param withdrawRoot The merkle root on layer2 after this batch.
    event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot);

    /// @notice Emitted when owner updates the proofWindow parameter.
    /// @param oldWindow    The old proofWindow.
    /// @param newWindow    The new proofWindow.
    event UpdateProofWindow(uint256 oldWindow, uint256 newWindow);

    /// @notice Emitted when owner updates the finalizationPeriodSeconds parameter.
    /// @param oldPeriod    The old finalizationPeriodSeconds.
    /// @param newPeriod    The new finalizationPeriodSeconds.
    event UpdateFinalizationPeriodSeconds(uint256 oldPeriod, uint256 newPeriod);

    /// @notice Emitted when owner updates the status of challenger.
    /// @param account  The address of account updated.
    /// @param status   The status of the account updated.
    event UpdateChallenger(address indexed account, bool status);

    /// @notice Emitted when the address of rollup verifier is updated.
    /// @param oldVerifier  The address of old rollup verifier.
    /// @param newVerifier  The address of new rollup verifier.
    event UpdateVerifier(address indexed oldVerifier, address indexed newVerifier);

    /// @notice Emitted when the proof reward percent is updated.
    /// @param oldPercent  The old proofRewardPercent.
    /// @param newPercent  The new proofRewardPercent.
    event UpdateProofRewardPercent(uint256 oldPercent, uint256 newPercent);

    /// @notice Emit when prove remaining claimed.
    /// @param receiver  receiver address.
    /// @param amount    claimed amount.
    event ProveRemainingClaimed(address receiver, uint256 amount);

    /// @notice Emitted when the state of Challenge is updated.
    /// @param batchIndex       The index of the batch.
    /// @param challenger       The address of challenger.
    /// @param challengeDeposit The deposit of challenger.
    event ChallengeState(uint64 indexed batchIndex, address indexed challenger, uint256 challengeDeposit);

    /// @notice Emitted when the result of Challenge is updated.
    /// @param batchIndex   The index of the batch.
    /// @param winner       The address of winner.
    /// @param res          The result of challenge.
    event ChallengeRes(uint256 indexed batchIndex, address indexed winner, string indexed res);

    /// @notice Emitted when the challenger claim the challenge reward.
    /// @param receiver  receiver address
    /// @param amount    claimed amount
    event ChallengeRewardClaim(address indexed receiver, uint256 amount);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice The latest finalized batch index.
    function lastFinalizedBatchIndex() external view returns (uint256);

    /// @notice The latest finalized batch index.
    function lastCommittedBatchIndex() external view returns (uint256);

    /// @notice Return the batch hash of a committed batch.
    /// @param batchIndex The index of the batch.
    function committedBatches(uint256 batchIndex) external view returns (bytes32);

    /// @notice Return the state root of a committed batch.
    /// @param batchIndex The index of the batch.
    function finalizedStateRoots(uint256 batchIndex) external view returns (bytes32);

    /// @notice Return the the committed batch of withdrawalRoot.
    /// @param withdrawalRoot The withdrawal root.
    function withdrawalRoots(bytes32 withdrawalRoot) external view returns (bool);

    /// @notice Return whether the batch is finalized by batch index.
    /// @param batchIndex The index of the batch.
    function isBatchFinalized(uint256 batchIndex) external view returns (bool);

    /// @notice Return the rollup config of finalizationPeriodSeconds.
    function finalizationPeriodSeconds() external view returns (uint256);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice Commit a batch of transactions on layer 1.
    ///
    /// @param batchDataInput       The BatchDataInput struct
    /// @param batchSignatureInput  The BatchSignatureInput struct
    function commitBatch(
        BatchDataInput calldata batchDataInput,
        BatchSignatureInput calldata batchSignatureInput
    ) external payable;

    /// @notice Revert a pending batch.
    /// @dev one can only revert unfinalized batches.
    /// @param batchHeader  The header of current batch, see the encoding in comments of `commitBatch`.
    /// @param count        The number of subsequent batches to revert, including current batch.
    function revertBatch(bytes calldata batchHeader, uint256 count) external;

    /// @notice Claim challenge reward
    /// @param receiver The receiver address
    function claimReward(address receiver) external;
}
