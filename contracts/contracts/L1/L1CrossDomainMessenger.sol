// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {IMessageDropCallback} from "../libraries/callbacks/IMessageDropCallback.sol";
import {Predeploys} from "../libraries/constants/Predeploys.sol";
import {Constants} from "../libraries/constants/Constants.sol";
import {CrossDomainMessenger} from "../libraries/CrossDomainMessenger.sol";
import {ICrossDomainMessenger} from "../libraries/ICrossDomainMessenger.sol";
import {Semver} from "../libraries/common/Semver.sol";
import {IL1MessageQueue} from "./rollup/IL1MessageQueue.sol";
import {IRollup} from "./rollup/IRollup.sol";
import {Verify} from "../libraries/common/Tree.sol";
import {IL1CrossDomainMessenger} from "./IL1CrossDomainMessenger.sol";

/**
 * @custom:proxied
 * @title L1CrossDomainMessenger
 * @notice The L1CrossDomainMessenger is a message passing interface between L1 and L2 responsible
 *         for sending and receiving data on the L1 side. Users are encouraged to use this
 *         interface instead of interacting with lower-level contracts directly.
 */
contract L1CrossDomainMessenger is
    Semver,
    IL1CrossDomainMessenger,
    CrossDomainMessenger,
    Verify
{
    /***********
     * Structs *
     ***********/

    /**
     * @notice Represents a proven withdrawal.
     *
     * @custom:field withdrawalRoot  Root of the L2 withdraw this was proven against.
     * @custom:field timestamp     Timestamp at which the withdrawal was proven.
     * @custom:field index         Index of the withdraw tx this was proven against.
     */
    struct ProvenWithdrawal {
        bytes32 withdrawalRoot;
        uint256 timestamp;
        uint256 withdrawalIndex;
    }

    struct ReplayState {
        // The number of replayed times.
        uint128 times;
        // The queue index of lastest replayed one. If it is zero, it means the message has not been replayed.
        uint128 lastIndex;
    }
    /*************
     * Variables *
     *************/

    /**
     * @notice A mapping of withdrawal hashes to `ProvenWithdrawal` data.
     */
    mapping(bytes32 => ProvenWithdrawal) public provenWithdrawals;

    /**
     * @notice A list of withdrawal hashes which have been successfully finalized.
     */
    mapping(bytes32 => bool) public finalizedWithdrawals;

    /// @notice Mapping from L1 message hash to the timestamp when the message is sent.
    mapping(bytes32 => uint256) public messageSendTimestamp;

    /// @notice Mapping from L1 message hash to drop status.
    mapping(bytes32 => bool) public isL1MessageDropped;

    /// @notice The address of Rollup contract.
    address public rollup;

    /// @notice The address of L1MessageQueue contract.
    address public messageQueue;

    /// @notice The maximum number of times each L1 message can be replayed.
    uint256 public maxReplayTimes;

    /// @notice Mapping from L1 message hash to replay state.
    mapping(bytes32 => ReplayState) public replayStates;

    /// @notice Mapping from queue index to previous replay queue index.
    ///
    /// @dev If a message `x` was replayed 3 times with index `q1`, `q2` and `q3`, the
    /// value of `prevReplayIndex` and `replayStates` will be `replayStates[hash(x)].lastIndex = q3`,
    /// `replayStates[hash(x)].times = 3`, `prevReplayIndex[q3] = q2`, `prevReplayIndex[q2] = q1`,
    /// `prevReplayIndex[q1] = x` and `prevReplayIndex[x]=nil`.
    ///
    /// @dev The index `x` that `prevReplayIndex[x]=nil` is used as the termination of the list.
    /// Usually we use `0` to represent `nil`, but we cannot distinguish it with the first message
    /// with index zero. So a nonzero offset `1` is added to the value of `prevReplayIndex[x]` to
    /// avoid such situation.
    mapping(uint256 => uint256) public prevReplayIndex;

    /***************
     * Constructor *
     ***************/

    constructor() Semver(1, 0, 0) {
        _disableInitializers();
    }

    /// @notice Initialize the storage of L1CrossDomainMessenger.
    /// @param _feeVault The address of fee vault, which will be used to collect relayer fee.
    /// @param _rollup The address of rollup contract.
    /// @param _messageQueue The address of L1MessageQueue contract.
    function initialize(
        address _feeVault,
        address _rollup,
        address _messageQueue
    ) public initializer {
        if (_rollup == address(0) || _messageQueue == address(0)) {
            revert ErrZeroAddress();
        }
        CrossDomainMessenger.__Messenger_init(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            _feeVault
        );

        rollup = _rollup;
        messageQueue = _messageQueue;
        counterpart = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        maxReplayTimes = 3;
        emit UpdateMaxReplayTimes(0, 3);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc ICrossDomainMessenger
    function sendMessage(
        address _to,
        uint256 _value,
        bytes memory _message,
        uint256 _gasLimit
    ) external payable override whenNotPaused {
        _sendMessage(_to, _value, _message, _gasLimit, _msgSender());
    }

    /// @inheritdoc ICrossDomainMessenger
    function sendMessage(
        address _to,
        uint256 _value,
        bytes calldata _message,
        uint256 _gasLimit,
        address _refundAddress
    ) external payable override whenNotPaused {
        _sendMessage(_to, _value, _message, _gasLimit, _refundAddress);
    }

    function proveMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _nonce,
        bytes memory _message,
        bytes32[32] calldata _withdrawalProof,
        bytes32 _withdrawalRoot
    ) external override whenNotPaused notInExecution {
        // @note check more `_to` address to avoid attack in the future when we add more gateways.
        require(_to != messageQueue, "Messenger: Forbid to call message queue");
        _validateTargetAddress(_to);

        // @note This usually will never happen, just in case.
        require(
            _from != xDomainMessageSender,
            "Messenger: Invalid message sender"
        );

        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, _to, _value, _nonce, _message)
        );
        ProvenWithdrawal memory provenWithdrawal = provenWithdrawals[
            _xDomainCalldataHash
        ];
        // We generally want to prevent users from proving the same withdrawal multiple times
        // because each successive proof will update the timestamp. A malicious user can take
        // advantage of this to prevent other users from finalizing their withdrawal. However,
        // since withdrawals are proven before an output root is finalized, we need to allow users
        // to re-prove their withdrawal only in the case that the output root for their specified
        // output index has been updated.
        require(
            provenWithdrawal.timestamp == 0,
            "Messenger: withdrawal hash has already been proven"
        );
        address _rollup = rollup;
        // withdrawalRoot for withdraw proof verify
        uint256 withdrawalBatchIndex = IRollup(_rollup).withdrawalRoots(
            _withdrawalRoot
        );
        require(
            (withdrawalBatchIndex > 0),
            "Messenger: do not submit withdrawalRoot"
        );
        // Verify that the hash of this withdrawal was stored in the L2toL1MessagePasser contract
        // on L2. If this is true, under the assumption that the Tree does not have
        // bugs, then we know that this withdrawal was actually triggered on L2 and can therefore
        // be relayed on L1.
        require(
            verifyMerkleProof(
                _xDomainCalldataHash,
                _withdrawalProof,
                _nonce,
                _withdrawalRoot
            ),
            "Messenger: invalid withdrawal inclusion proof"
        );

        // Designate the withdrawalHash as proven by storing the `withdrawalRoot`, `timestamp`, and
        // `withdrawalIndex` in the `provenWithdrawals` mapping. A `withdrawalHash` can only be
        // proven once unless it is submitted again with a different withdrawalRoot.
        provenWithdrawals[_xDomainCalldataHash] = ProvenWithdrawal({
            withdrawalRoot: _withdrawalRoot,
            timestamp: block.timestamp,
            withdrawalIndex: _nonce
        });

        // Emit a `WithdrawalProven` event.
        emit WithdrawalProven(_xDomainCalldataHash, _from, _to);
    }

    function relayMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _nonce,
        bytes memory _message
    ) external override whenNotPaused notInExecution {
        bytes32 _xDomainCalldataHash = keccak256(
            _encodeXDomainCalldata(_from, _to, _value, _nonce, _message)
        );
        // Check that this withdrawal has not already been finalized, this is replay protection.
        require(
            finalizedWithdrawals[_xDomainCalldataHash] == false,
            "Messenger: withdrawal has already been finalized"
        );

        ProvenWithdrawal memory provenWithdrawal = provenWithdrawals[
            _xDomainCalldataHash
        ];

        // A withdrawal can only be finalized if it has been proven. We know that a withdrawal has
        // been proven at least once when its timestamp is non-zero. Unproven withdrawals will have
        // a timestamp of zero.
        require(
            provenWithdrawal.timestamp != 0,
            "Messenger: withdrawal has not been proven yet"
        );

        // A proven withdrawal must wait at least the finalization period before it can be
        // finalized. This waiting period can elapse in parallel with the waiting period for the
        // output the withdrawal was proven against. In effect, this means that the minimum
        // withdrawal time is proposal submission time + finalization period.
        require(
            _isFinalizationPeriodElapsed(provenWithdrawal.timestamp),
            "Messenger: proven withdrawal finalization period has not elapsed"
        );

        {
            address _rollup = rollup;
            // withdrawalRoot for withdraw proof verify
            uint256 withdrawalBatchIndex = IRollup(_rollup).withdrawalRoots(
                provenWithdrawal.withdrawalRoot
            );
            require(
                (withdrawalBatchIndex > 0),
                "Messenger: do not submit withdrawalRoot"
            );

            bytes32 finStateRoots = IRollup(_rollup).finalizedStateRoots(
                withdrawalBatchIndex
            );
            require(
                finStateRoots != bytes32(0),
                "Messenger: batch not verified"
            );
        }

        xDomainMessageSender = _from;
        (bool success, ) = _to.call{value: _value}(_message);
        // reset value to refund gas.
        xDomainMessageSender = Constants.DEFAULT_XDOMAIN_MESSAGE_SENDER;

        if (success) {
            // Mark the withdrawal as finalized so it can't be replayed.
            finalizedWithdrawals[_xDomainCalldataHash] = true;
            emit RelayedMessage(_xDomainCalldataHash);
        } else {
            emit FailedRelayedMessage(_xDomainCalldataHash);
        }
    }

    /// @inheritdoc IL1CrossDomainMessenger
    function replayMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _messageNonce,
        bytes memory _message,
        uint32 _newGasLimit,
        address _refundAddress
    ) external payable override whenNotPaused notInExecution {
        // We will use a different `queueIndex` for the replaced message. However, the original `queueIndex` or `nonce`
        // is encoded in the `_message`. We will check the `xDomainCalldata` on layer 2 to avoid duplicated execution.
        // So, only one message will succeed on layer 2. If one of the message is executed successfully, the other one
        // will revert with "Message was already successfully executed".
        address _messageQueue = messageQueue;
        address _counterpart = counterpart;
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            _from,
            _to,
            _value,
            _messageNonce,
            _message
        );
        bytes32 _xDomainCalldataHash = keccak256(_xDomainCalldata);

        require(
            messageSendTimestamp[_xDomainCalldataHash] > 0,
            "Provided message has not been enqueued"
        );
        // cannot replay dropped message
        require(
            !isL1MessageDropped[_xDomainCalldataHash],
            "Message already dropped"
        );

        // compute and deduct the messaging fee to fee vault.
        uint256 _fee = IL1MessageQueue(_messageQueue)
            .estimateCrossDomainMessageFee(_newGasLimit);

        // charge relayer fee
        require(msg.value >= _fee, "Insufficient msg.value for fee");
        if (_fee > 0) {
            (bool _success, ) = feeVault.call{value: _fee}("");
            require(_success, "Failed to deduct the fee");
        }

        // enqueue the new transaction
        uint256 _nextQueueIndex = IL1MessageQueue(_messageQueue)
            .nextCrossDomainMessageIndex();
        IL1MessageQueue(_messageQueue).appendCrossDomainMessage(
            _counterpart,
            _newGasLimit,
            _xDomainCalldata
        );

        ReplayState memory _replayState = replayStates[_xDomainCalldataHash];
        // update the replayed message chain.
        unchecked {
            if (_replayState.lastIndex == 0) {
                // the message has not been replayed before.
                prevReplayIndex[_nextQueueIndex] = _messageNonce + 1;
            } else {
                prevReplayIndex[_nextQueueIndex] = _replayState.lastIndex + 1;
            }
        }
        _replayState.lastIndex = uint128(_nextQueueIndex);

        // update replay times
        require(
            _replayState.times < maxReplayTimes,
            "Exceed maximum replay times"
        );
        unchecked {
            _replayState.times += 1;
        }
        replayStates[_xDomainCalldataHash] = _replayState;

        // refund fee to `_refundAddress`
        unchecked {
            uint256 _refund = msg.value - _fee;
            if (_refund > 0) {
                (bool _success, ) = _refundAddress.call{value: _refund}("");
                require(_success, "Failed to refund the fee");
            }
        }
    }

    /// @inheritdoc IL1CrossDomainMessenger
    function dropMessage(
        address _from,
        address _to,
        uint256 _value,
        uint256 _messageNonce,
        bytes memory _message
    ) external override whenNotPaused notInExecution {
        // The criteria for dropping a message:
        // 1. The message is a L1 message.
        // 2. The message has not been dropped before.
        // 3. the message and all of its replacement are finalized in L1.
        // 4. the message and all of its replacement are skipped.
        //
        // Possible denial of service attack:
        // + replayMessage is called every time someone want to drop the message.
        // + replayMessage is called so many times for a skipped message, thus results a long list.
        //
        // We limit the number of `replayMessage` calls of each message, which may solve the above problem.

        address _messageQueue = messageQueue;

        // check message exists
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            _from,
            _to,
            _value,
            _messageNonce,
            _message
        );
        bytes32 _xDomainCalldataHash = keccak256(_xDomainCalldata);
        require(
            messageSendTimestamp[_xDomainCalldataHash] > 0,
            "Provided message has not been enqueued"
        );

        // check message not dropped
        require(
            !isL1MessageDropped[_xDomainCalldataHash],
            "Message already dropped"
        );

        // check message is finalized
        uint256 _lastIndex = replayStates[_xDomainCalldataHash].lastIndex;
        if (_lastIndex == 0) _lastIndex = _messageNonce;

        // check message is skipped and drop it.
        // @note If the list is very long, the message may never be dropped.
        while (true) {
            IL1MessageQueue(_messageQueue).dropCrossDomainMessage(_lastIndex);
            _lastIndex = prevReplayIndex[_lastIndex];
            if (_lastIndex == 0) break;
            unchecked {
                _lastIndex = _lastIndex - 1;
            }
        }

        isL1MessageDropped[_xDomainCalldataHash] = true;

        // set execution context
        xDomainMessageSender = Constants.DROP_XDOMAIN_MESSAGE_SENDER;
        IMessageDropCallback(_from).onDropMessage{value: _value}(_message);
        // clear execution context
        xDomainMessageSender = Constants.DEFAULT_XDOMAIN_MESSAGE_SENDER;
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice Update max replay times.
    /// @dev This function can only called by contract owner.
    /// @param _newMaxReplayTimes The new max replay times.
    function updateMaxReplayTimes(
        uint256 _newMaxReplayTimes
    ) external onlyOwner {
        uint256 _oldMaxReplayTimes = maxReplayTimes;
        maxReplayTimes = _newMaxReplayTimes;

        emit UpdateMaxReplayTimes(_oldMaxReplayTimes, _newMaxReplayTimes);
    }

    /**********************
     * Internal Functions *
     **********************/

    function _sendMessage(
        address _to,
        uint256 _value,
        bytes memory _message,
        uint256 _gasLimit,
        address _refundAddress
    ) internal nonReentrant {
        address _messageQueue = messageQueue; // gas saving
        address _counterpart = counterpart; // gas saving

        // compute the actual cross domain message calldata.
        uint256 _messageNonce = IL1MessageQueue(_messageQueue)
            .nextCrossDomainMessageIndex();
        bytes memory _xDomainCalldata = _encodeXDomainCalldata(
            _msgSender(),
            _to,
            _value,
            _messageNonce,
            _message
        );

        // compute and deduct the messaging fee to fee vault.
        uint256 _fee = IL1MessageQueue(_messageQueue)
            .estimateCrossDomainMessageFee(_gasLimit);
        require(msg.value >= _fee + _value, "Insufficient msg.value");
        if (_fee > 0) {
            (bool _success, ) = feeVault.call{value: _fee}("");
            require(_success, "Failed to deduct the fee");
        }

        // append message to L1MessageQueue
        IL1MessageQueue(_messageQueue).appendCrossDomainMessage(
            _counterpart,
            _gasLimit,
            _xDomainCalldata
        );

        // record the message hash for future use.
        bytes32 _xDomainCalldataHash = keccak256(_xDomainCalldata);

        // normally this won't happen, since each message has different nonce, but just in case.
        require(
            messageSendTimestamp[_xDomainCalldataHash] == 0,
            "Duplicated message"
        );
        messageSendTimestamp[_xDomainCalldataHash] = block.timestamp;

        emit SentMessage(
            _msgSender(),
            _to,
            _value,
            _messageNonce,
            _gasLimit,
            _message
        );

        // refund fee to `_refundAddress`
        unchecked {
            uint256 _refund = msg.value - _fee - _value;
            if (_refund > 0) {
                (bool _success, ) = _refundAddress.call{value: _refund}("");
                require(_success, "Failed to refund the fee");
            }
        }
    }

    /**
     * @notice Determines whether the finalization period has elapsed w/r/t a given timestamp.
     *
     * @param _timestamp Timestamp to check.
     *
     * @return Whether or not the finalization period has elapsed.
     */
    function _isFinalizationPeriodElapsed(
        uint256 _timestamp
    ) internal view returns (bool) {
        return
            block.timestamp >
            _timestamp + IRollup(rollup).FINALIZATION_PERIOD_SECONDS();
    }

    function messageNonce()
        external
        view
        override(ICrossDomainMessenger, CrossDomainMessenger)
        returns (uint256)
    {
        return IL1MessageQueue(messageQueue).nextCrossDomainMessageIndex();
    }
}
