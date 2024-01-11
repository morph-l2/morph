// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Types} from "./Types.sol";
import {Hashing} from "./Hashing.sol";
import {RLPWriter} from "./rlp/RLPWriter.sol";
import {RLPReader} from "./rlp/RLPReader.sol";

/**
 * @title Encoding
 * @notice Encoding handles Morph's various different encoding schemes.
 */
library Encoding {
    /**
     * @notice RLP encodes the L2 transaction that would be generated when a given deposit is sent
     *         to the L2 system. Useful for searching for a deposit in the L2 system. The
     *         transaction is prefixed with 0x7e to identify its EIP-2718 type.
     *
     * @param _tx User deposit transaction to encode.
     *
     * @return RLP encoded L2 deposit transaction.
     */
    function encodeL1MessageTx(
        Types.L1MessageTx memory _tx
    ) internal pure returns (bytes memory) {
        bytes[] memory raw = new bytes[](6);
        raw[0] = RLPWriter.writeUint(uint256(_tx.queueIndex));
        raw[1] = RLPWriter.writeUint(uint256(_tx.gas));
        raw[2] = RLPWriter.writeAddress(_tx.to);
        raw[3] = RLPWriter.writeUint(_tx.value);
        raw[4] = RLPWriter.writeBytes(_tx.data);
        raw[5] = RLPWriter.writeAddress(_tx.sender);

        return abi.encodePacked(uint8(0x7e), RLPWriter.writeList(raw));
    }

    /**
     * @notice Encodes the cross domain message based on the version that is encoded into the
     *         message nonce.
     *
     * @param _nonce    Message nonce with version encoded into the first two bytes.
     * @param _sender   Address of the sender of the message.
     * @param _target   Address of the target of the message.
     * @param _value    ETH value to send to the target.
     * @param _gasLimit Gas limit to use for the message.
     * @param _data     Data to send with the message.
     *
     * @return Encoded cross domain message.
     */
    function encodeCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) internal pure returns (bytes memory) {
        (, uint16 version) = decodeVersionedNonce(_nonce);
        if (version == 0) {
            return encodeCrossDomainMessageV0(_target, _sender, _data, _nonce);
        } else if (version == 1) {
            return
                encodeCrossDomainMessageV1(
                    _nonce,
                    _sender,
                    _target,
                    _value,
                    _gasLimit,
                    _data
                );
        } else {
            revert("Encoding: unknown cross domain message version");
        }
    }

    /**
     * @notice Encodes a cross domain message based on the V0 (legacy) encoding.
     *
     * @param _target Address of the target of the message.
     * @param _sender Address of the sender of the message.
     * @param _data   Data to send with the message.
     * @param _nonce  Message nonce.
     *
     * @return Encoded cross domain message.
     */
    function encodeCrossDomainMessageV0(
        address _target,
        address _sender,
        bytes memory _data,
        uint256 _nonce
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "relayMessage(address,address,bytes,uint256)",
                _target,
                _sender,
                _data,
                _nonce
            );
    }

    /**
     * @notice Encodes a cross domain message based on the V1 (current) encoding.
     *
     * @param _nonce    Message nonce.
     * @param _sender   Address of the sender of the message.
     * @param _target   Address of the target of the message.
     * @param _value    ETH value to send to the target.
     * @param _gasLimit Gas limit to use for the message.
     * @param _data     Data to send with the message.
     *
     * @return Encoded cross domain message.
     */
    function encodeCrossDomainMessageV1(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) internal pure returns (bytes memory) {
        return
            abi.encodeWithSignature(
                "relayMessage(uint256,address,address,uint256,uint256,bytes)",
                _nonce,
                _sender,
                _target,
                _value,
                _gasLimit,
                _data
            );
    }

    /**
     * @notice Adds a version number into the first two bytes of a message nonce.
     *
     * @param _nonce   Message nonce to encode into.
     * @param _version Version number to encode into the message nonce.
     *
     * @return Message nonce with version encoded into the first two bytes.
     */
    function encodeVersionedNonce(
        uint240 _nonce,
        uint16 _version
    ) internal pure returns (uint256) {
        uint256 nonce;
        assembly {
            nonce := or(shl(240, _version), _nonce)
        }
        return nonce;
    }

    /**
     * @notice Pulls the version out of a version-encoded nonce.
     *
     * @param _nonce Message nonce with version encoded into the first two bytes.
     *
     * @return Nonce without encoded version.
     * @return Version of the message.
     */
    function decodeVersionedNonce(
        uint256 _nonce
    ) internal pure returns (uint240, uint16) {
        uint240 nonce;
        uint16 version;
        assembly {
            nonce := and(
                _nonce,
                0x0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
            )
            version := shr(240, _nonce)
        }
        return (nonce, version);
    }

    /**
     * @notice Encodes sequencers information with length.
     *
     * @param infos Sequencer informations.
     *
     * @return Encoded Sequencer informations with length.
     */
    function encodeSequencerInfos(
        Types.SequencerInfo[] memory infos
    ) internal pure returns (bytes memory) {
        uint256 len = infos.length;
        bytes[] memory raw = new bytes[](len + 1);
        raw[0] = RLPWriter.writeUint(len);
        for (uint256 i = 0; i < len; i++) {
            bytes memory data = abi.encode(infos[i]);
            raw[i + 1] = RLPWriter.writeBytes(data);
        }
        return RLPWriter.writeList(raw);
    }

    /**
     * @notice Decodes sequencers information.
     *
     * @return Decode Sequencer informations array.
     */
    function decodeSequencerInfos(
        bytes memory infosBytes
    ) internal pure returns (Types.SequencerInfo[] memory) {
        RLPReader.RLPItem memory rlpitem = RLPReader.toRLPItem(infosBytes);
        RLPReader.RLPItem[] memory list = RLPReader.readList(rlpitem);
        bytes memory lenBytes = RLPReader.readBytes(list[0]);
        uint256 len = abi.decode(padWithZero(lenBytes), (uint256));
        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            len
        );
        for (uint256 i = 0; i < len; i++) {
            sequencerInfos[i] = abi.decode(
                RLPReader.readBytes(list[i + 1]),
                (Types.SequencerInfo)
            );
        }
        return sequencerInfos;
    }

    function padWithZero(
        bytes memory data
    ) internal pure returns (bytes memory) {
        require(data.length <= 32, "Data length exceeds 32 bytes");

        uint256 bytesLen = 32;
        bytes memory paddedData = new bytes(bytesLen);
        for (uint256 i = 0; i < data.length; i++) {
            paddedData[bytesLen - data.length + i] = data[i];
        }
        return paddedData;
    }
}
