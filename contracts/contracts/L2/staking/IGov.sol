// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IGov {
    struct ProposalData {
        uint256 batchBlockInterval;
        uint256 batchMaxBytes;
        uint256 batchTimeout;
        uint256 rollupEpoch;
        uint256 maxChunks;
    }

    /**
     * @notice batch block interval
     */
    function batchBlockInterval() external view returns (uint256);

    /**
     * @notice batch max bytes
     */
    function batchMaxBytes() external view returns (uint256);

    /**
     * @notice next batch index
     */
    function batchTimeout() external view returns (uint256);

    /**
     * @notice next batch index
     */
    function rollupEpoch() external view returns (uint256);

    /**
     * @notice max chunks
     */
    function maxChunks() external view returns (uint256);

    /**
     * @notice create a proposal
     */
    function propose(ProposalData memory proposal) external;

    /**
     * @notice vote a propsal
     */
    function vote(uint256 _proposalId) external;

    /**
     * @notice execute an approved proposal
     */
    function executeProposal(uint256 _proposalId) external;

    /**
     * @notice whether the proposal can be approved
     */
    function isProposalCanBeApproved(
        uint256 _proposalId
    ) external view returns (bool);

    /**
     * @notice proposal information.
     * @custom:field _proposalId
     * @return {approved, end timestamp}
     */
    function proposalInfos(
        uint256 _proposalId
    ) external view returns (uint256, bool);
}
