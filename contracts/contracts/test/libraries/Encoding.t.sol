// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CommonTest} from "../CommonTest.t.sol";
import {Types} from "../../libraries/Types.sol";
import {Encoding} from "../../libraries/Encoding.sol";
import {LegacyCrossDomainUtils} from "../../libraries/LegacyCrossDomainUtils.sol";

contract Encoding_Test is CommonTest {
    function testFuzz_nonceVersioning_succeeds(
        uint240 _nonce,
        uint16 _version
    ) external {
        (uint240 nonce, uint16 version) = Encoding.decodeVersionedNonce(
            Encoding.encodeVersionedNonce(_nonce, _version)
        );
        assertEq(version, _version);
        assertEq(nonce, _nonce);
    }

    function testDiff_decodeVersionedNonce_succeeds(
        uint240 _nonce,
        uint16 _version
    ) external {
        uint256 nonce = uint256(
            Encoding.encodeVersionedNonce(_nonce, _version)
        );
        (uint256 decodedNonce, uint256 decodedVersion) = ffi
            .decodeVersionedNonce(nonce);

        assertEq(_version, uint16(decodedVersion));

        assertEq(_nonce, uint240(decodedNonce));
    }

    function testDiff_encodeCrossDomainMessage_succeeds(
        uint240 _nonce,
        uint8 _version,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        uint8 version = _version % 2;
        uint256 nonce = Encoding.encodeVersionedNonce(_nonce, version);

        bytes memory encoding = Encoding.encodeCrossDomainMessage(
            nonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );

        bytes memory _encoding = ffi.encodeCrossDomainMessage(
            nonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );

        assertEq(encoding, _encoding);
    }

    function testFuzz_encodeCrossDomainMessageV0_matchesLegacy_succeeds(
        uint240 _nonce,
        address _sender,
        address _target,
        bytes memory _data
    ) external {
        uint8 version = 0;
        uint256 nonce = Encoding.encodeVersionedNonce(_nonce, version);

        bytes memory legacyEncoding = LegacyCrossDomainUtils
            .encodeXDomainCalldata(_target, _sender, _data, nonce);

        bytes memory crossMessageEncoding = Encoding.encodeCrossDomainMessageV0(
            _target,
            _sender,
            _data,
            nonce
        );

        assertEq(legacyEncoding, crossMessageEncoding);
    }

    function testDiff_encodeL1MessageTx_succeeds(
        uint64 _queueIndex,
        uint64 _gas,
        address _to,
        uint256 _value,
        bytes memory _data,
        address _sender
    ) external {
        Types.L1MessageTx memory _tx = Types.L1MessageTx(
            _queueIndex,
            _gas,
            _to,
            _value,
            _data,
            _sender
        );

        bytes memory txn = Encoding.encodeL1MessageTx(_tx);
        bytes memory _txn = ffi.encodeL1MessageTx(_tx);

        assertEq(txn, _txn);
    }
}
