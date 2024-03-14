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
    ) external returns (bool);

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        uint256[] memory sequencerIndex,
        address challenger,
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external;

    /**
     * @notice update sequencers version
     * @param _sequencerAddrs sequencer addresses
     * @param _sequencerBytes sequencer information bytes
     * @param _sequencerBLSKeys sequencer BLS keys
     */
    function updateAndSendSequencerSet(
        bytes memory _sequencerBytes,
        address[] memory _sequencerAddrs,
        bytes[] memory _sequencerBLSKeys,
        uint32 gasLimit,
        address _refundAddress
    ) external payable;

    /**
     * @notice sequencer addresses
     * @param version version
     */
    function getSequencerAddrs(
        uint256 version
    ) external view returns (address[] memory);

    /**
     * @notice whether is current sequencer
     * @param addr address
     */
    function isSequencer(address addr) external view returns (bool);

    /**
     * @notice sequencer BLS keys
     * @param version version
     */
    function getSequencerBLSKeys(
        uint256 version
    ) external view returns (bytes[] memory);

    /**
     * @notice sequencers num
     * @param version sequencer version
     */
    function sequencerNum(uint256 version) external view returns (uint256);
}
