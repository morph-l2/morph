// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "../../libraries/common/Types.sol";

interface IL2Sequencer {
    function updateSequencers(uint256 version, Types.SequencerInfo[] memory _sequencers) external;

    /**
     * @notice current sequencers version
     */
    function currentVersion() external view returns (uint256);

    /**
     * @notice current sequencers version height
     */
    function currentVersionHeight() external view returns (uint256);

    /**
     * @notice pre sequencers version height
     */
    function preVersionHeight() external view returns (uint256);

    /**
     * @notice get sequencer history
     */
    function getSequencerHistory(
        uint256 version
    ) external view returns (Types.SequencerHistory memory);

    /**
     * @notice get sequencers addresses
     */
    function getSequencerAddresses(bool previous) external view returns (address[] memory);

    /**
     * @notice get sequencers addresses
     */
    function getSequencerInfos(bool previous) external view returns (Types.SequencerInfo[] memory);

    /**
     * @notice get address is in sequencers set
     */
    function inSequencersSet(bool previous, address addr) external view returns (bool, uint256);

    /**
     * @notice get the index of address in sequencers set
     */
    function sequencerIndex(bool previous, address addr) external view returns (uint256, uint256);

    /**
     * @notice get the length of sequencerAddresses
     */
    function sequencersLen(bool previous) external view returns (uint256, uint256);
}
