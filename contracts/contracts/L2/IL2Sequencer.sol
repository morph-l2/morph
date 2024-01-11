// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import "../libraries/Types.sol";

interface IL2Sequencer {
    /**
     * @notice current sequencers version
     */
    function currentVersion() external view returns (uint256);

    /**
     * @notice current sequencers version height
     */
    function currentVersionHeight() external view returns (uint256);

    /**
     * @notice pre sequencers version
     */
    function preVersion() external view returns (uint256);

    /**
     * @notice pre sequencers version height
     */
    function preVersionHeight() external view returns (uint256);

    /**
     * @notice get sequencers addresses
     */
    function sequencerAddresses(uint256 index) external view returns (address);

    /**
     * @notice get pre sequencers addresses
     */
    function preSequencerAddresses(
        uint256 index
    ) external view returns (address);

    /**
     * @notice get sequencers addresses
     */
    function getSequencerAddresses(
        bool previous
    ) external view returns (address[] memory);

    /**
     * @notice get sequencers addresses
     */
    function getSequencerInfos(
        bool previous
    ) external view returns (Types.SequencerInfo[] memory);

    /**
     * @notice get address is in sequencers set
     */
    function inSequencersSet(
        bool previous,
        address addr
    ) external view returns (bool, uint256);

    /**
     * @notice get the index of address in sequencers set
     */
    function sequencerIndex(
        bool previous,
        address addr
    ) external view returns (uint256, uint256);

    /**
     * @notice get the length of sequencerAddresses
     */
    function sequencersLen(
        bool previous
    ) external view returns (uint256, uint256);
}
