// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {CommonTest} from "../CommonTest.t.sol";
import {Types} from "../../libraries/Types.sol";
import {Hashing} from "../../libraries/Hashing.sol";
import {Encoding} from "../../libraries/Encoding.sol";
import {LegacyCrossDomainUtils} from "../../libraries/LegacyCrossDomainUtils.sol";

contract Hashing_hashDepositSource_Test is CommonTest {
    /**
     * @notice Tests that hashDepositSource returns the correct hash in a simple case.
     */
    function test_hashDepositSource_succeeds() external {
        assertEq(
            Hashing.hashDepositSource(
                0xd25df7858efc1778118fb133ac561b138845361626dfb976699c5287ed0f4959,
                0x1
            ),
            0xf923fb07134d7d287cb52c770cc619e17e82606c21a875c92f4c63b65280a5cc
        );
    }
}

contract Hashing_hashCrossDomainMessage_Test is CommonTest {
    /**
     * @notice Tests that hashCrossDomainMessage returns the correct hash in a simple case.
     */
    function testDiff_hashCrossDomainMessage_succeeds(
        uint240 _nonce,
        uint16 _version,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        // Ensure the version is valid.
        uint16 version = uint16(bound(uint256(_version), 0, 1));
        uint256 nonce = Encoding.encodeVersionedNonce(_nonce, version);

        assertEq(
            Hashing.hashCrossDomainMessage(
                nonce,
                _sender,
                _target,
                _value,
                _gasLimit,
                _data
            ),
            ffi.hashCrossDomainMessage(
                nonce,
                _sender,
                _target,
                _value,
                _gasLimit,
                _data
            )
        );
    }

    /**
     * @notice Tests that hashCrossDomainMessageV0 matches the hash of the legacy encoding.
     */
    function testFuzz_hashCrossDomainMessageV0_matchesLegacy_succeeds(
        address _target,
        address _sender,
        bytes memory _message,
        uint256 _messageNonce
    ) external {
        assertEq(
            keccak256(
                LegacyCrossDomainUtils.encodeXDomainCalldata(
                    _target,
                    _sender,
                    _message,
                    _messageNonce
                )
            ),
            Hashing.hashCrossDomainMessageV0(
                _target,
                _sender,
                _message,
                _messageNonce
            )
        );
    }
}

contract Hashing_hashWithdrawal_Test is CommonTest {
    /**
     * @notice Tests that hashWithdrawal returns the correct hash in a simple case.
     */
    function testDiff_hashWithdrawal_succeeds(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        assertEq(
            Hashing.hashWithdrawal(
                Types.WithdrawalTransaction(
                    _nonce,
                    _sender,
                    _target,
                    _value,
                    _gasLimit,
                    _data
                )
            ),
            ffi.hashWithdrawal(
                _nonce,
                _sender,
                _target,
                _value,
                _gasLimit,
                _data
            )
        );
    }
}

contract Hashing_hashDepositTransaction_Test is CommonTest {
    /**
     * @notice Tests that hashDepositTransaction returns the correct hash in a simple case.
     */
    function testDiff_hashDepositTransaction_succeeds(
        uint64 _queueIndex,
        uint64 _gas,
        address _to,
        uint256 _value,
        bytes memory _data,
        address _sender
    ) external {
        assertEq(
            Hashing.hashL1MessageTx(
                Types.L1MessageTx(
                    _queueIndex,
                    _gas,
                    _to,
                    _value,
                    _data,
                    _sender
                )
            ),
            ffi.hashL1MessageTx(_queueIndex, _gas, _to, _value, _data, _sender)
        );
    }
}

contract Test_Tree is CommonTest {
    function testDiff_getProveWithdrawalTransactionInputs_succeeds(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external {
        Types.WithdrawalTransaction memory _tx = Types.WithdrawalTransaction(
            _nonce,
            _sender,
            _target,
            _value,
            _gasLimit,
            _data
        );

        bytes32 _withdrawalHash;
        bytes32[_TREE_DEPTH] memory _withdrawalProof;
        bytes32 _withdrawalRoot;

        (_withdrawalHash, _withdrawalProof, _withdrawalRoot) = ffi
            .getProveWithdrawalTransactionInputs(_tx);

        _appendMessageHash(
            Hashing.hashWithdrawal(
                Types.WithdrawalTransaction(
                    _nonce,
                    _sender,
                    _target,
                    _value,
                    _gasLimit,
                    _data
                )
            )
        );
        assertEq(_withdrawalRoot, getTreeRoot());
        assertEq(
            verifyMerkleProof(
                _withdrawalHash,
                _withdrawalProof,
                0,
                _withdrawalRoot
            ),
            true
        );
    }
}
