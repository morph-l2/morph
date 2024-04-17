// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "../../libraries/common/Types.sol";

interface ISequencer {
    // event of sequencer update
    event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight);

    /**
     * @notice update sequencer set
     */
    function updateSequencerSet(address[] memory newSequencerSet) external;

    /**
     * @notice get current sequencer set
     */
    function getCurrentSeqeuncerSet() external view returns (address[] memory);

    /**
     * @notice get current sequencer set size
     */
    function getCurrentSeqeuncerSetSize() external view returns (uint256);

    /**
     * @notice get latest sequencer set
     */
    function getLatestSeqeuncerSet() external view returns (address[] memory);

    /**
     * @notice get latest sequencer set size
     */
    function getLatestSeqeuncerSetSize() external view returns (uint256);

    /**
     * @notice whether the address is a latest sequencer
     */
    function isSequencer(address addr) external view returns (bool);

    /**
     * @notice whether the address is a current sequencer
     */
    function isCurrentSequencer(address addr) external view returns (bool);
}
