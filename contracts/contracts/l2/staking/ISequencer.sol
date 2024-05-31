// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface ISequencer {
    /**********
     * Events *
     **********/

    /// @notice event of sequencer update
    event SequencerSetUpdated(address[] sequencerSet, uint256 blockHeight);

    /*************************
     * Public View Functions *
     *************************/

    /// @notice get current sequencer set
    function getCurrentSequencerSet() external view returns (address[] memory);

    /// @notice get current sequencer set size
    function getCurrentSequencerSetSize() external view returns (uint256);

    /// @notice get sequencer set 0
    function getSequencerSet0() external view returns (address[] memory);

    /// @notice get size of sequencer set 0
    function getSequencerSet0Size() external view returns (uint256);

    /// @notice get sequencer set 1
    function getSequencerSet1() external view returns (address[] memory);

    /// @notice get size of sequencer set 1
    function getSequencerSet1Size() external view returns (uint256);

    /// @notice get sequencer set 2
    function getSequencerSet2() external view returns (address[] memory);

    /// @notice get size of sequencer set 2
    function getSequencerSet2Size() external view returns (uint256);

    /// @notice whether the address is a latest sequencer
    function isSequencer(address addr) external view returns (bool);

    /// @notice whether the address is a current sequencer
    function isCurrentSequencer(address addr) external view returns (bool);

    /// @notice get the encoded sequencer set bytes
    function getSequencerSetBytes() external view returns (bytes memory);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice update sequencer set
    function updateSequencerSet(address[] calldata newSequencerSet) external;
}
