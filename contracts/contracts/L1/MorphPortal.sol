// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {SafeCall} from "../libraries/SafeCall.sol";
import {SystemConfig} from "./SystemConfig.sol";
import {Constants} from "../libraries/Constants.sol";
import {Types} from "../libraries/Types.sol";
import {Hashing} from "../libraries/Hashing.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {ResourceMetering} from "./ResourceMetering.sol";
import {Semver} from "../universal/Semver.sol";
import {Verify} from "../universal/Tree.sol";
import {Rollup} from "./Rollup.sol";
import {L1MessageStorage} from "./L1MessageStorage.sol";
import {IL1MessageQueue} from "./IL1MessageQueue.sol";

/**
 * @custom:proxied
 * @title MorphPortal
 * @notice The MorphPortal is a low-level contract responsible for passing messages between L1
 *         and L2. Messages sent directly to the MorphPortal have no form of replayability.
 *         Users are encouraged to use the L1CrossDomainMessenger for a higher-level interface.
 */
contract MorphPortal is
    Initializable,
    ResourceMetering,
    Semver,
    Verify,
    L1MessageStorage,
    IL1MessageQueue
{
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

    /**
     * @notice Version of the deposit event.
     */
    uint256 internal constant DEPOSIT_VERSION = 0;

    /**
     * @notice The L2 gas limit set when eth is deposited using the receive() function.
     */
    uint64 internal constant RECEIVE_DEFAULT_GAS_LIMIT = 100_000;

    /**
     * @notice Address of the SystemConfig contract.
     */
    SystemConfig public immutable SYSTEM_CONFIG;

    /**
     * @notice Address that has the ability to pause and unpause withdrawals.
     */
    address public immutable GUARDIAN;

    /**
     * @notice Address of the Rollup contract.
     */
    Rollup public immutable ROLLUP;

    /**
     * @notice Address of the L2 account which initiated a withdrawal in this transaction. If the
     *         of this variable is the default L2 sender address, then we are NOT inside of a call
     *         to finalizeWithdrawalTransaction.
     */
    address public l2Sender;

    /**
     * @notice Address of the l1CrossDomainMessenger contract.
     */
    address public l1Messenger;

    /**
     * @notice A list of withdrawal hashes which have been successfully finalized.
     */
    mapping(bytes32 => bool) public finalizedWithdrawals;

    /**
     * @notice A mapping of withdrawal hashes to `ProvenWithdrawal` data.
     */
    mapping(bytes32 => ProvenWithdrawal) public provenWithdrawals;

    /**
     * @notice Determines if cross domain messaging is paused. When set to true,
     *         withdrawals are paused. This may be removed in the future.
     */
    bool public paused;

    /**
     * @notice Emitted when a transaction is deposited from L1 to L2. The parameters of this event
     *         are read by the rollup node and used to derive deposit transactions on L2.
     *
     * @param from       Address that triggered the deposit transaction.
     * @param to         Address that the deposit transaction is directed to.
     * @param version    Version of this deposit transaction event.
     * @param opaqueData ABI encoded deposit data to be parsed off-chain.
     */
    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 indexed version,
        bytes opaqueData
    );

    /**
     * @notice Emitted when a withdrawal transaction is proven.
     *
     * @param withdrawalHash Hash of the withdrawal transaction.
     */
    event WithdrawalProven(
        bytes32 indexed withdrawalHash,
        address indexed from,
        address indexed to
    );

    /**
     * @notice Emitted when a withdrawal transaction is finalized.
     *
     * @param withdrawalHash Hash of the withdrawal transaction.
     * @param success        Whether the withdrawal transaction was successful.
     */
    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);

    /**
     * @notice Emitted when the pause is triggered.
     *
     * @param account Address of the account triggering the pause.
     */
    event Paused(address account);

    /**
     * @notice Emitted when the pause is lifted.
     *
     * @param account Address of the account triggering the unpause.
     */
    event Unpaused(address account);

    /**
     * @notice Reverts when paused.
     */
    modifier whenNotPaused() {
        require(paused == false, "MorphPortal: paused");
        _;
    }

    /**
     * @notice Reverts when caller not messager contract.
     */
    modifier onlyMessenger() {
        require(
            msg.sender == l1Messenger,
            "messenger contract unauthenticated"
        );
        _;
    }

    /**
     * @notice Reverts when caller not rollup contract.
     */
    modifier onlyRollup() {
        require(
            msg.sender == address(ROLLUP),
            "ROLLUP contract unauthenticated"
        );
        _;
    }

    /**
     * todo: change version
     * @custom:semver 1.6.0
     *
     * @param _guardian                  Address that can pause deposits and withdrawals.
     * @param _paused                    Sets the contract's pausability state.
     * @param _config                    Address of the SystemConfig contract.
     */
    constructor(
        address _guardian,
        bool _paused,
        SystemConfig _config,
        Rollup _rollup,
        address _l1Messenger
    ) Semver(1, 6, 0) {
        GUARDIAN = _guardian;
        SYSTEM_CONFIG = _config;
        ROLLUP = _rollup;
        initialize(_paused, _l1Messenger);
    }

    /**
     * @notice Initializer.
     */
    function initialize(bool _paused, address _l1Messenger) public initializer {
        l2Sender = Constants.DEFAULT_L2_SENDER;
        paused = _paused;
        l1Messenger = _l1Messenger;
        __ResourceMetering_init();
    }

    /**
     * @notice Pause deposits and withdrawals.
     */
    function pause() external {
        require(
            msg.sender == GUARDIAN,
            "MorphPortal: only guardian can pause"
        );
        paused = true;
        emit Paused(msg.sender);
    }

    /**
     * @notice Unpause deposits and withdrawals.
     */
    function unpause() external {
        require(
            msg.sender == GUARDIAN,
            "MorphPortal: only guardian can unpause"
        );
        paused = false;
        emit Unpaused(msg.sender);
    }

    /**
     * @notice Computes the minimum gas limit for a deposit. The minimum gas limit
     *         linearly increases based on the size of the calldata. This is to prevent
     *         users from creating L2 resource usage without paying for it. This function
     *         can be used when interacting with the portal to ensure forwards compatibility.
     *
     */
    function minimumGasLimit(uint64 _byteCount) public pure returns (uint64) {
        return _byteCount * 16 + 21000;
    }

    /**
     * @notice Getter for the resource config. Used internally by the ResourceMetering
     *         contract. The SystemConfig is the source of truth for the resource config.
     *
     * @return ResourceMetering.ResourceConfig
     */
    function _resourceConfig()
        internal
        view
        override
        returns (ResourceMetering.ResourceConfig memory)
    {
        return SYSTEM_CONFIG.resourceConfig();
    }

    /**
     * @notice Proves a withdrawal transaction.
     *
     * @param _tx              Withdrawal transaction to finalize.
     * @param _withdrawalProof Inclusion proof of the withdrawal in L2ToL1MessagePasser contract.
     * @param _withdrawalRoot  Inclusion root of the withdrawal in L2ToL1MessagePasser contract.
     */
    function proveWithdrawalTransaction(
        Types.WithdrawalTransaction memory _tx,
        bytes32[32] calldata _withdrawalProof,
        bytes32 _withdrawalRoot
    ) external whenNotPaused {
        // Prevent users from creating a deposit transaction where this address is the message
        // sender on L2. Because this is checked here, we do not need to check again in
        // `finalizeWithdrawalTransaction`.
        require(
            _tx.target != address(this),
            "MorphPortal: you cannot send messages to the portal contract"
        );

        // Load the ProvenWithdrawal into memory, using the withdrawal hash as a unique identifier.
        bytes32 withdrawalHash = Hashing.hashWithdrawal(_tx);
        ProvenWithdrawal memory provenWithdrawal = provenWithdrawals[
            withdrawalHash
        ];

        // We generally want to prevent users from proving the same withdrawal multiple times
        // because each successive proof will update the timestamp. A malicious user can take
        // advantage of this to prevent other users from finalizing their withdrawal. However,
        // since withdrawals are proven before an output root is finalized, we need to allow users
        // to re-prove their withdrawal only in the case that the output root for their specified
        // output index has been updated.
        require(
            provenWithdrawal.timestamp == 0,
            "MorphPortal: withdrawal hash has already been proven"
        );

        // withdrawalRoot for withdraw proof verify
        uint256 withdrawalBatchIndex = ROLLUP.withdrawalRoots(_withdrawalRoot);
        require(
            (withdrawalBatchIndex > 0),
            "MorphPortal: do not submit withdrawalRoot"
        );

        // Verify that the hash of this withdrawal was stored in the L2toL1MessagePasser contract
        // on L2. If this is true, under the assumption that the Tree does not have
        // bugs, then we know that this withdrawal was actually triggered on L2 and can therefore
        // be relayed on L1.
        require(
            verifyMerkleProof(
                withdrawalHash,
                _withdrawalProof,
                _tx.nonce,
                _withdrawalRoot
            ),
            "MorphPortal: invalid withdrawal inclusion proof"
        );

        // Designate the withdrawalHash as proven by storing the `withdrawalRoot`, `timestamp`, and
        // `withdrawalIndex` in the `provenWithdrawals` mapping. A `withdrawalHash` can only be
        // proven once unless it is submitted again with a different withdrawalRoot.
        provenWithdrawals[withdrawalHash] = ProvenWithdrawal({
            withdrawalRoot: _withdrawalRoot,
            timestamp: block.timestamp,
            withdrawalIndex: _tx.nonce
        });

        // Emit a `WithdrawalProven` event.
        emit WithdrawalProven(withdrawalHash, _tx.sender, _tx.target);
    }

    /**
     * @notice Finalizes a withdrawal transaction.
     *
     * @param _tx Withdrawal transaction to finalize.
     */
    function finalizeWithdrawalTransaction(
        Types.WithdrawalTransaction memory _tx
    ) external whenNotPaused {
        // Make sure that the l2Sender has not yet been set. The l2Sender is set to a value other
        // than the default value when a withdrawal transaction is being finalized. This check is
        // a defacto reentrancy guard.
        require(
            l2Sender == Constants.DEFAULT_L2_SENDER,
            "MorphPortal: can only trigger one withdrawal per transaction"
        );

        // Grab the proven withdrawal from the `provenWithdrawals` map.
        bytes32 withdrawalHash = Hashing.hashWithdrawal(_tx);
        ProvenWithdrawal memory provenWithdrawal = provenWithdrawals[
            withdrawalHash
        ];

        // A withdrawal can only be finalized if it has been proven. We know that a withdrawal has
        // been proven at least once when its timestamp is non-zero. Unproven withdrawals will have
        // a timestamp of zero.
        require(
            provenWithdrawal.timestamp != 0,
            "MorphPortal: withdrawal has not been proven yet"
        );

        // A proven withdrawal must wait at least the finalization period before it can be
        // finalized. This waiting period can elapse in parallel with the waiting period for the
        // output the withdrawal was proven against. In effect, this means that the minimum
        // withdrawal time is proposal submission time + finalization period.
        require(
            _isFinalizationPeriodElapsed(provenWithdrawal.timestamp),
            "MorphPortal: proven withdrawal finalization period has not elapsed"
        );

        // withdrawalRoot for withdraw proof verify
        uint256 withdrawalBatchIndex = ROLLUP.withdrawalRoots(
            provenWithdrawal.withdrawalRoot
        );
        require(
            (withdrawalBatchIndex > 0),
            "MorphPortal: do not submit withdrawalRoot"
        );
        bytes32 finStateRoots = ROLLUP.finalizedStateRoots(
            withdrawalBatchIndex
        );
        require(finStateRoots != bytes32(0), "batch not verified");

        // Check that this withdrawal has not already been finalized, this is replay protection.
        require(
            finalizedWithdrawals[withdrawalHash] == false,
            "MorphPortal: withdrawal has already been finalized"
        );

        // Mark the withdrawal as finalized so it can't be replayed.
        finalizedWithdrawals[withdrawalHash] = true;

        // Set the l2Sender so contracts know who triggered this withdrawal on L2.
        l2Sender = _tx.sender;

        // Trigger the call to the target contract. We use a custom low level method
        // SafeCall.callWithMinGas to ensure two key properties
        //   1. Target contracts cannot force this call to run out of gas by returning a very large
        //      amount of data (and this is OK because we don't care about the returndata here).
        //   2. The amount of gas provided to the execution context of the target is at least the
        //      gas limit specified by the user. If there is not enough gas in the current context
        //      to accomplish this, `callWithMinGas` will revert.
        bool success = SafeCall.callWithMinGas(
            _tx.target,
            _tx.gasLimit,
            _tx.value,
            _tx.data
        );

        // Reset the l2Sender back to the default value.
        l2Sender = Constants.DEFAULT_L2_SENDER;

        // All withdrawals are immediately finalized. Replayability can
        // be achieved through contracts built on top of this contract
        emit WithdrawalFinalized(withdrawalHash, success);

        // Reverting here is useful for determining the exact gas cost to successfully execute the
        // sub call to the target contract if the minimum gas limit specified by the user would not
        // be sufficient to execute the sub call.
        if (success == false && tx.origin == Constants.ESTIMATION_ADDRESS) {
            revert("MorphPortal: withdrawal failed");
        }
    }

    /**
     * @notice Accepts deposits of ETH and data, and emits a TransactionDeposited event for use in
     *         deriving deposit transactions. Note that if a deposit is made by a contract, its
     *         address will be aliased when retrieved using `tx.origin` or `msg.sender`. Consider
     *         using the CrossDomainMessenger contracts for a simpler developer experience.
     *
     * @param _to         Target address on L2.
     * @param _value      ETH value to send to the recipient.
     * @param _gasLimit   Minimum L2 gas limit (can be greater than or equal to this value).
     * @param _isCreation Whether or not the transaction is a contract creation.
     * @param _data       Data to trigger the recipient with.
     */
    function depositTransaction(
        address _to,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes calldata _data
    ) public payable metered(_gasLimit) onlyMessenger {
        // Just to be safe, make sure that people specify address(0) as the target when doing
        // contract creations.
        if (_isCreation) {
            require(
                _to == address(0),
                "MorphPortal: must send to address(0) when creating a contract"
            );
        }

        // Prevent depositing transactions that have too small of a gas limit. Users should pay
        // more for more resource usage.
        require(
            _gasLimit >= minimumGasLimit(uint64(_data.length)),
            "MorphPortal: gas limit too small"
        );

        // Prevent the creation of deposit transactions that have too much calldata. This gives an
        // upper limit on the size of unsafe blocks over the p2p network. 120kb is chosen to ensure
        // that the transaction can fit into the p2p network policy of 128kb even though deposit
        // transactions are not gossipped over the p2p network.
        require(_data.length <= 120_000, "MorphPortal: data too large");

        address from = AddressAliasHelper.applyL1ToL2Alias(msg.sender);
        _queueTransaction(from, _to, _value, _gasLimit, _data);

        // Compute the opaque data that will be emitted as part of the TransactionDeposited event.
        // We use opaque data so that we can update the TransactionDeposited event in the future
        // without breaking the current interface.
        bytes memory opaqueData = abi.encodePacked(
            msg.value,
            _value,
            _gasLimit,
            _isCreation,
            _data
        );

        // Emit a TransactionDeposited event so that the rollup node can derive a deposit
        // transaction for this deposit.
        emit TransactionDeposited(from, _to, DEPOSIT_VERSION, opaqueData);
    }

    function popCrossDomainMessage(
        uint256 _startIndex,
        uint256 _count,
        uint256 _skippedBitmap
    ) external onlyRollup {
        _popCrossDomainMessage(_startIndex, _count, _skippedBitmap);
    }

    function appendEnforcedTransaction(
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes calldata _data
    ) external onlyMessenger {
        require(_sender.code.length == 0, "only EOA");
        _queueTransaction(_sender, _target, _value, _gasLimit, _data);
    }

    function dropCrossDomainMessage(uint256 _index) external onlyMessenger {
        _dropCrossDomainMessage(_index);
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
            block.timestamp > _timestamp + ROLLUP.FINALIZATION_PERIOD_SECONDS();
    }

    // @inheritdoc IL1MessageQueue
    function getCrossDomainMessage(
        uint256 _queueIndex
    ) external view returns (bytes32) {
        return messageQueue[_queueIndex];
    }
}
