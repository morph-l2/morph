// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IL1Sequencer {
    function pause() external;

    /**
     * @notice newest sequencers version
     */
    function newestVersion() external view returns (uint256);

    /**
     * @notice current sequencers version
     */
    function currentVersion() external view returns (uint256);

    /**
     * @notice verify BLS signature
     * @param version sequencer set version
     * @param indexs sequencer index
     * @param signature batch signature
     */
    function verifySignature(
        uint256 version,
        uint256[] memory indexs,
        bytes memory signature
    ) external;

    /**
     * @notice update sequencers version
     * @param _sequencerBytes sequencer information bytes
     * @param _sequencerBLSKeys sequencer BLS keys
     */
    function updateAndSendSequencerSet(
        bytes memory _sequencerBytes,
        bytes[] memory _sequencerBLSKeys,
        uint32 gasLimit,
        address _refundAddress
    ) external payable;
}
