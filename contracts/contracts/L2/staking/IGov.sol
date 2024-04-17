// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IGov {
    struct ProposalInfo {
        uint256 endTime;
        bool approved;
    }

    struct ProposalData {
        uint256 batchBlockInterval;
        uint256 batchMaxBytes;
        uint256 batchTimeout;
        uint256 maxChunks;
        uint256 rollupEpoch;
    }

    // event of proposal executed
    event ProposalExecuted(
        uint256 batchBlockInterval,
        uint256 batchMaxBytes,
        uint256 batchTimeout,
        uint256 maxChunks,
        uint256 rollupEpoch
    );

    /**
     * @notice batch block interval
     */
    function batchBlockInterval() external view returns (uint256);

    /**
     * @notice batch max bytes
     */
    function batchMaxBytes() external view returns (uint256);

    /**
     * @notice batch timeout
     */
    function batchTimeout() external view returns (uint256);

    /**
     * @notice rollup epoch
     */
    function rollupEpoch() external view returns (uint256);

    /**
     * @notice max chunks
     */
    function maxChunks() external view returns (uint256);

    /**
     * @notice current proposal ID number
     */
    function proposalID() external view returns (uint256);

    /**
     * @notice create a proposal
     */
    function propose(ProposalData memory proposal) external;

    /**
     * @notice vote a propsal
     */
    function vote(uint256 _proposalID) external;

    /**
     * @notice execute an approved proposal
     */
    function executeProposal(uint256 _proposalID) external;

    /**
     * @notice whether the proposal can be approved
     */
    function isProposalCanBeApproved(
        uint256 _proposalID
    ) external view returns (bool);

    /**
     * @notice proposal information.
     * @custom:field _proposalID
     * @return {approved, end timestamp}
     */
    function proposalInfos(
        uint256 _proposalID
    ) external view returns (uint256, bool);

    /**
     * @custom:field _proposalID
     * @custom:field _voter
     * @return {bool}, check if an account has been voted
     */
    function isVoted(
        uint256 _proposalID,
        address _voter
    ) external view returns (bool);
}
