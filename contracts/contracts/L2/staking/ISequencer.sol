// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "../../libraries/common/Types.sol";

interface ISequencer {
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
     * @notice whether the address is a sequencer
     */
    function isSequencer(address addr) external view returns (bool);
}
