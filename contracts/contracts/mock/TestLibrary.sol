// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Hashing} from "../libraries/Hashing.sol";
import {Types} from "../libraries/Types.sol";

contract TestLibrary {
    function hashL1MessageTx(
        Types.L1MessageTx memory _tx
    ) public pure returns (bytes32) {
        bytes32 hash = Hashing.hashL1MessageTx(
            Types.L1MessageTx({
                queueIndex: _tx.queueIndex,
                gas: _tx.gas,
                to: _tx.to,
                value: _tx.value,
                data: _tx.data,
                sender: _tx.sender
            })
        );
        return hash;
    }

    function hashWithdrawal(
        Types.WithdrawalTransaction memory _tx
    ) public pure returns (bytes32) {
        bytes32 withdrawalHash = Hashing.hashWithdrawal(
            Types.WithdrawalTransaction({
                nonce: _tx.nonce,
                sender: _tx.sender,
                target: _tx.target,
                value: _tx.value,
                gasLimit: _tx.gasLimit,
                data: _tx.data
            })
        );
        return withdrawalHash;
    }

    function hashCrossDomainMessageV1(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _minGasLimit,
        bytes calldata _message
    ) public pure returns (bytes32) {
        return
            Hashing.hashCrossDomainMessageV1(
                _nonce,
                _sender,
                _target,
                _value,
                _minGasLimit,
                _message
            );
    }
}
