// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IGov {
    /***********
     * Structs *
     ***********/

    /// @custom:field voting expiration time
    /// @custom:field executed
    struct ProposalInfo {
        uint256 expirationTime;
        bool executed;
    }

    /// @custom:field batchBlockInterval
    /// @custom:field batchMaxBytes
    /// @custom:field batchTimeout
    /// @custom:field maxChunks
    /// @custom:field rollupEpoch
    struct ProposalData {
        uint256 batchBlockInterval;
        uint256 batchMaxBytes;
        uint256 batchTimeout;
        uint256 maxChunks;
        uint256 rollupEpoch;
    }

    /**********
     * Events *
     **********/

    /// @notice event of create a proposal
    event ProposalCreated(
        uint256 indexed proposalID,
        address indexed creator,
        uint256 batchBlockInterval,
        uint256 batchMaxBytes,
        uint256 batchTimeout,
        uint256 maxChunks,
        uint256 rollupEpoch
    );

    /// @notice event of proposal executed
    event ProposalExecuted(
        uint256 indexed proposalID,
        uint256 batchBlockInterval,
        uint256 batchMaxBytes,
        uint256 batchTimeout,
        uint256 maxChunks,
        uint256 rollupEpoch
    );

    /// @notice proposal voting duration updated
    /// @param oldProposalVotingDuration    old proposal voting duration
    /// @param newProposalVotingDuration    new proposal voting duration
    event VotingDurationUpdated(uint256 oldProposalVotingDuration, uint256 newProposalVotingDuration);

    /// @notice batch block interval updated
    /// @param oldBatchBlockInterval    old batch block interval
    /// @param newBatchBlockInterval    new batch block interval
    event BatchBlockIntervalUpdated(uint256 oldBatchBlockInterval, uint256 newBatchBlockInterval);

    /// @notice batch max bytes updated
    /// @param oldBatchMaxBytes     old batch max bytes
    /// @param newBatchMaxBytes     new batch max bytes
    event BatchMaxBytesUpdated(uint256 oldBatchMaxBytes, uint256 newBatchMaxBytes);

    /// @notice batch timeout updated
    /// @param oldBatchTimeout  old batch timeout
    /// @param newBatchTimeout  new batch timeout
    event BatchTimeoutUpdated(uint256 oldBatchTimeout, uint256 newBatchTimeout);

    /// @notice max chunks updated
    /// @param oldMaxChunks     old max chunks
    /// @param newMaxChunks     new max chunks
    event MaxChunksUpdated(uint256 oldMaxChunks, uint256 newMaxChunks);

    /// @notice rollup epoch updated
    /// @param odlRollupEpoch   old rollup epoch
    /// @param newRollupEpoch   new rollup epoch
    event RollupEpochUpdated(uint256 odlRollupEpoch, uint256 newRollupEpoch);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice batch block interval
    function batchBlockInterval() external view returns (uint256);

    /// @notice batch max bytes
    function batchMaxBytes() external view returns (uint256);

    /// @notice batch timeout
    function batchTimeout() external view returns (uint256);

    /// @notice rollup epoch
    function rollupEpoch() external view returns (uint256);

    /// @notice max chunks
    function maxChunks() external view returns (uint256);

    /// @notice current proposal ID number
    function currentProposalID() external view returns (uint256);

    /// @notice return proposal status
    function proposalStatus(uint256 proposalID) external view returns (bool finished, bool passed, bool executed);

    /// @notice proposal information.
    /// @param proposalID  proposal ID
    function proposalInfos(uint256 proposalID) external view returns (uint256 endTimestamp, bool passed);

    /// @notice return whether the address has voted
    /// @param proposalID  proposal ID
    /// @param voter       voter
    function isVoted(uint256 proposalID, address voter) external view returns (bool);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice vote a proposal
    function vote(uint256 proposalID) external;

    /// @notice execute a passed proposal
    function executeProposal(uint256 proposalID) external;
}
