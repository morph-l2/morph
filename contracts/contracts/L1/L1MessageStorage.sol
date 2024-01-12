// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {BitMapsUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/BitMapsUpgradeable.sol";

contract L1MessageStorage {
    using BitMapsUpgradeable for BitMapsUpgradeable.BitMap;

    /// @notice The list of queued cross domain messages.
    bytes32[] public messageQueue;

    uint256 public pendingQueueIndex;

    /// @dev The bitmap for dropped messages, where `droppedMessageBitmap[i]` keeps the bits from `[i*256, (i+1)*256)`.
    BitMapsUpgradeable.BitMap private droppedMessageBitmap;

    /// @dev The bitmap for skipped messages, where `skippedMessageBitmap[i]` keeps the bits from `[i*256, (i+1)*256)`.
    mapping(uint256 => uint256) private skippedMessageBitmap;

    /// @notice Emitted when a new L1 => L2 transaction is appended to the queue.
    /// @param sender The address of account who initiates the transaction.
    /// @param target The address of account who will receive the transaction.
    /// @param value The value passed with the transaction.
    /// @param queueIndex The index of this transaction in the queue.
    /// @param gasLimit Gas limit required to complete the message relay on L2.
    /// @param data The calldata of the transaction.
    event QueueTransaction(
        address indexed sender,
        address indexed target,
        uint256 value,
        uint64 queueIndex,
        uint256 gasLimit,
        bytes data
    );

    /// @notice Emitted when some L1 => L2 transactions are included in L1.
    /// @param startIndex The start index of messages popped.
    /// @param count The number of messages popped.
    /// @param skippedBitmap A bitmap indicates whether a message is skipped.
    event DequeueTransaction(
        uint256 startIndex,
        uint256 count,
        uint256 skippedBitmap
    );

    /// @notice Emitted when a message is dropped from L1.
    /// @param index The index of message dropped.
    event DropTransaction(uint256 index);

    constructor() {}

    function nextCrossDomainMessageIndex() external view returns (uint256) {
        return messageQueue.length;
    }

    function computeTransactionHash(
        address _sender,
        uint256 _queueIndex,
        uint256 _value,
        address _target,
        uint256 _gasLimit,
        bytes calldata _data
    ) public pure returns (bytes32) {
        // We use EIP-2718 to encode the L1 message, and the encoding of the message is
        //      `TransactionType || TransactionPayload`
        // where
        //  1. `TransactionType` is 0x7E
        //  2. `TransactionPayload` is `rlp([queueIndex, gasLimit, to, value, data, sender])`
        //
        // The spec of rlp: https://ethereum.org/en/developers/docs/data-structures-and-encoding/rlp/
        uint256 transactionType = 0x7E;
        bytes32 hash;
        assembly {
            function get_uint_bytes(v) -> len {
                if eq(v, 0) {
                    len := 1
                    leave
                }
                for {

                } gt(v, 0) {

                } {
                    len := add(len, 1)
                    v := shr(8, v)
                }
            }

            // This is used for both store uint and single byte.
            // Integer zero is special handled by geth to encode as `0x80`
            function store_uint_or_byte(_ptr, v, is_uint) -> ptr {
                ptr := _ptr
                switch lt(v, 128)
                case 1 {
                    switch and(iszero(v), is_uint)
                    case 1 {
                        // integer 0
                        mstore8(ptr, 0x80)
                    }
                    default {
                        // single byte in the [0x00, 0x7f]
                        mstore8(ptr, v)
                    }
                    ptr := add(ptr, 1)
                }
                default {
                    // 1-32 bytes long
                    let len := get_uint_bytes(v)
                    mstore8(ptr, add(len, 0x80))
                    ptr := add(ptr, 1)
                    mstore(ptr, shl(mul(8, sub(32, len)), v))
                    ptr := add(ptr, len)
                }
            }

            function store_address(_ptr, v) -> ptr {
                ptr := _ptr
                // 20 bytes long
                mstore8(ptr, 0x94) // 0x80 + 0x14
                ptr := add(ptr, 1)
                mstore(ptr, shl(96, v))
                ptr := add(ptr, 0x14)
            }

            // 1 byte for TransactionType
            // 4 byte for list payload length
            let start_ptr := add(mload(0x40), 5)
            let ptr := start_ptr
            ptr := store_uint_or_byte(ptr, _queueIndex, 1)
            ptr := store_uint_or_byte(ptr, _gasLimit, 1)
            ptr := store_address(ptr, _target)
            ptr := store_uint_or_byte(ptr, _value, 1)

            switch eq(_data.length, 1)
            case 1 {
                // single byte
                ptr := store_uint_or_byte(
                    ptr,
                    byte(0, calldataload(_data.offset)),
                    0
                )
            }
            default {
                switch lt(_data.length, 56)
                case 1 {
                    // a string is 0-55 bytes long
                    mstore8(ptr, add(0x80, _data.length))
                    ptr := add(ptr, 1)
                    calldatacopy(ptr, _data.offset, _data.length)
                    ptr := add(ptr, _data.length)
                }
                default {
                    // a string is more than 55 bytes long
                    let len_bytes := get_uint_bytes(_data.length)
                    mstore8(ptr, add(0xb7, len_bytes))
                    ptr := add(ptr, 1)
                    mstore(ptr, shl(mul(8, sub(32, len_bytes)), _data.length))
                    ptr := add(ptr, len_bytes)
                    calldatacopy(ptr, _data.offset, _data.length)
                    ptr := add(ptr, _data.length)
                }
            }
            ptr := store_address(ptr, _sender)

            let payload_len := sub(ptr, start_ptr)
            let value
            let value_bytes
            switch lt(payload_len, 56)
            case 1 {
                // the total payload of a list is 0-55 bytes long
                value := add(0xc0, payload_len)
                value_bytes := 1
            }
            default {
                // If the total payload of a list is more than 55 bytes long
                let len_bytes := get_uint_bytes(payload_len)
                value_bytes := add(len_bytes, 1)
                value := add(0xf7, len_bytes)
                value := shl(mul(len_bytes, 8), value)
                value := or(value, payload_len)
            }
            value := or(value, shl(mul(8, value_bytes), transactionType))
            value_bytes := add(value_bytes, 1)
            let value_bits := mul(8, value_bytes)
            value := or(
                shl(sub(256, value_bits), value),
                shr(value_bits, mload(start_ptr))
            )
            start_ptr := sub(start_ptr, value_bytes)
            mstore(start_ptr, value)
            hash := keccak256(start_ptr, sub(ptr, start_ptr))
        }
        return hash;
    }

    function isMessageSkipped(
        uint256 _queueIndex
    ) external view returns (bool) {
        if (_queueIndex >= pendingQueueIndex) return false;

        return _isMessageSkipped(_queueIndex);
    }

    function isMessageDropped(
        uint256 _queueIndex
    ) external view returns (bool) {
        // it should be a skipped message first.
        return
            _isMessageSkipped(_queueIndex) &&
            droppedMessageBitmap.get(_queueIndex);
    }

    function _popCrossDomainMessage(
        uint256 _startIndex,
        uint256 _count,
        uint256 _skippedBitmap
    ) internal {
        require(_count <= 256, "pop too many messages");
        require(pendingQueueIndex == _startIndex, "start index mismatch");

        unchecked {
            // clear extra bits in `_skippedBitmap`, and if _count = 256, it's designed to overflow.
            uint256 mask = (1 << _count) - 1;
            _skippedBitmap &= mask;

            uint256 bucket = _startIndex >> 8;
            uint256 offset = _startIndex & 0xff;

            skippedMessageBitmap[bucket] |= _skippedBitmap << offset;
            if (offset + _count > 256) {
                skippedMessageBitmap[bucket + 1] =
                    _skippedBitmap >>
                    (256 - offset);
            }

            pendingQueueIndex = _startIndex + _count;
        }

        emit DequeueTransaction(_startIndex, _count, _skippedBitmap);
    }

    function _dropCrossDomainMessage(uint256 _index) internal {
        require(_index < pendingQueueIndex, "cannot drop pending message");

        require(_isMessageSkipped(_index), "drop non-skipped message");
        require(!droppedMessageBitmap.get(_index), "message already dropped");
        droppedMessageBitmap.set(_index);

        emit DropTransaction(_index);
    }

    /// @dev Internal function to queue a L1 transaction.
    /// @param _sender The address of sender who will initiate this transaction in L2.
    /// @param _target The address of target contract to call in L2.
    /// @param _value The value passed
    /// @param _gasLimit The maximum gas should be used for this transaction in L2.
    /// @param _data The calldata passed to target contract.
    function _queueTransaction(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes calldata _data
    ) internal {
        // compute transaction hash
        uint256 _queueIndex = messageQueue.length;
        bytes32 _hash = computeTransactionHash(
            _sender,
            _queueIndex,
            _value,
            _target,
            _gasLimit,
            _data
        );
        messageQueue.push(_hash);

        // emit event
        emit QueueTransaction(
            _sender,
            _target,
            _value,
            uint64(_queueIndex),
            _gasLimit,
            _data
        );
    }

    /// @dev Returns whether the bit at `index` is set.
    function _isMessageSkipped(uint256 index) internal view returns (bool) {
        uint256 bucket = index >> 8;
        uint256 mask = 1 << (index & 0xff);
        return skippedMessageBitmap[bucket] & mask != 0;
    }
}
