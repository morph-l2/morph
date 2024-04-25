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
        uint256 batchBlockInterval,
        uint256 batchMaxBytes,
        uint256 batchTimeout,
        uint256 maxChunks,
        uint256 rollupEpoch
    );

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
    function createProposal(ProposalData memory proposal) external;

    /// @notice vote a proposal
    function vote(uint256 proposalID) external;

    /// @notice execute an approved proposal
    function executeProposal(uint256 proposalID) external;
}
