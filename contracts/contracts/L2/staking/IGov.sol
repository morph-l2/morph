// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IGov {
    /***********
     * Structs *
     ***********/

    /// @custom:field endTime
    /// @custom:field approved
    struct ProposalInfo {
        uint256 endTime;
        bool approved;
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

    /// @notice event of proposal executed
    event ProposalExecuted(
        uint256 indexed proposalID,
        uint256 batchBlockInterval,
        uint256 batchMaxBytes,
        uint256 batchTimeout,
        uint256 maxChunks,
        uint256 rollupEpoch
    );

    /// @notice proposal interval updated
    /// @param oldProposalInterval    old proposal interval
    /// @param newProposalInterval    new proposal interval
    event ProposalIntervalUpdated(
        uint256 oldProposalInterval,
        uint256 newProposalInterval
    );

    /// @notice batch block interval updated
    /// @param oldBatchBlockInterval    old batch block interval
    /// @param newBatchBlockInterval    new batch block interval
    event BatchBlockIntervalUpdated(
        uint256 oldBatchBlockInterval,
        uint256 newBatchBlockInterval
    );

    /// @notice batch max bytes updated
    /// @param oldBatchMaxBytes     old batch max bytes
    /// @param newBatchMaxBytes     new batch max bytes
    event BatchMaxBytesUpdated(
        uint256 oldBatchMaxBytes,
        uint256 newBatchMaxBytes
    );

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

    /// @notice whether the proposal can be approved
    function isProposalCanBeApproved(
        uint256 proposalID
    ) external view returns (bool);

    /// @notice proposal information.
    /// @param proposalID  proposal ID
    function proposalInfos(
        uint256 proposalID
    ) external view returns (uint256 endTimestamp, bool approved);

    /// @notice return is voted
    /// @param proposalID  proposal ID
    /// @param voter       voter
    function isVoted(
        uint256 proposalID,
        address voter
    ) external view returns (bool voted);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice create a proposal
    function createProposal(ProposalData calldata proposal) external;

    /// @notice vote a proposal
    function vote(uint256 proposalID) external;

    /// @notice execute an approved proposal
    function executeProposal(uint256 proposalID) external;
}
