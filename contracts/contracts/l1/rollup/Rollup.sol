// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {BatchHeaderCodecV0} from "../../libraries/codec/BatchHeaderCodecV0.sol";
import {ChunkCodecV0} from "../../libraries/codec/ChunkCodecV0.sol";
import {IRollupVerifier} from "../../libraries/verifier/IRollupVerifier.sol";
import {IL1MessageQueue} from "./IL1MessageQueue.sol";
import {IRollup} from "./IRollup.sol";
import {IL1Staking} from "../staking/IL1Staking.sol";

// solhint-disable no-inline-assembly
// solhint-disable reason-string

/// @title Rollup
/// @notice This contract maintains data for the Morph rollup.
contract Rollup is IRollup, OwnableUpgradeable, PausableUpgradeable {
    /*************
     * Constants *
     *************/

    /// @notice The zero versioned hash.
    bytes32 internal constant ZERO_VERSIONED_HASH =
        0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;

    /// @notice The BLS MODULUS
    uint256 internal constant BLS_MODULUS =
        52435875175126190479447740508185965837690552500527637822603658699938581184513;

    /// @dev Address of the point evaluation precompile used for EIP-4844 blob verification.
    address internal constant POINT_EVALUATION_PRECOMPILE_ADDR = address(0x0A);

    /// @notice The chain id of the corresponding layer 2 chain.
    uint64 public immutable LAYER_2_CHAIN_ID;

    /*************
     * Variables *
     *************/

    /// @notice L1 staking contract
    address public l1StakingContract;

    /// @notice Batch challenge time.
    uint256 public finalizationPeriodSeconds;

    /// @notice The time when zkProof was generated and executed.
    uint256 public proofWindow;

    /// @notice The maximum number of transactions allowed in each chunk.
    uint256 public maxNumTxInChunk;

    /// @notice The address of L1MessageQueue.
    address public messageQueue;

    /// @notice The address of RollupVerifier.
    address public verifier;

    /// @inheritdoc IRollup
    uint256 public override lastFinalizedBatchIndex;

    /// @inheritdoc IRollup
    uint256 public override lastCommittedBatchIndex;

    /// @notice Whether an account is a challenger.
    mapping(address challengerAddress => bool isChallenger) public isChallenger;

    /// @inheritdoc IRollup
    mapping(uint256 batchIndex => bytes32 stateRoot)
        public
        override finalizedStateRoots;

    /// @notice Store committed batch base.
    mapping(uint256 batchIndex => BatchBase) public batchBaseStore;

    /// @notice Store committed batch data.
    mapping(uint256 batchIndex => BatchData) public batchDataStore;

    /// @notice Store committed batch signature.
    mapping(uint256 batchIndex => BatchSignature) public batchSignatureStore;

    /// @notice Store the withdrawalRoot.
    mapping(bytes32 withdrawalRoot => bool exist) public withdrawalRoots;

    /// @notice Store Challenge Information.
    mapping(uint256 batchIndex => BatchChallenge) public challenges;

    /// @notice Store Challenge reward information.
    mapping(address owner => uint256 amount) public batchChallengeReward;

    /// @notice Whether in challenge
    bool public inChallenge;

    /// @notice The batch being challenged
    uint256 public batchChallenged;

    /// @notice The index of the revert request.
    uint256 public revertReqIndex;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Only staker allowed.
    modifier OnlyStaker() {
        require(
            IL1Staking(l1StakingContract).isStaker(_msgSender()),
            "only staker allowed"
        );
        _;
    }

    /// @notice Only challenger allowed.
    modifier onlyChallenger() {
        require(isChallenger[_msgSender()], "caller challenger allowed");
        _;
    }

    /// @notice Modifier to ensure that there is no pending revert request.
    modifier nonReqRevert() {
        require(revertReqIndex == 0, "need revert");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    /// @param _chainID     The chain ID
    constructor(uint64 _chainID) {
        LAYER_2_CHAIN_ID = _chainID;
    }

    /// @notice Allow the contract to receive ETH.
    receive() external payable {}

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    /// @param _l1StakingContract         l1 staking contract
    /// @param _messageQueue              message queue
    /// @param _verifier                  verifier
    /// @param _maxNumTxInChunk           max num tx in chunk
    /// @param _finalizationPeriodSeconds finalization period seconds
    /// @param _proofWindow               proof window
    function initialize(
        address _l1StakingContract,
        address _messageQueue,
        address _verifier,
        uint256 _maxNumTxInChunk,
        uint256 _finalizationPeriodSeconds,
        uint256 _proofWindow
    ) public initializer {
        if (_messageQueue == address(0) || _verifier == address(0)) {
            revert ErrZeroAddress();
        }
        require(
            _l1StakingContract != address(0),
            "invalid l1 staking contract"
        );

        __Pausable_init();
        __Ownable_init();

        l1StakingContract = _l1StakingContract;
        messageQueue = _messageQueue;
        verifier = _verifier;
        maxNumTxInChunk = _maxNumTxInChunk;
        finalizationPeriodSeconds = _finalizationPeriodSeconds;
        proofWindow = _proofWindow;

        emit UpdateVerifier(address(0), _verifier);
        emit UpdateMaxNumTxInChunk(0, _maxNumTxInChunk);
        emit UpdateFinalizationPeriodSeconds(0, _finalizationPeriodSeconds);
        emit UpdateProofWindow(0, _proofWindow);
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice Import layer 2 genesis block
    function importGenesisBatch(
        bytes calldata _batchHeader,
        bytes32 _postStateRoot
    ) external onlyOwner {
        // check genesis batch header length
        require(_postStateRoot != bytes32(0), "zero state root");

        // check whether the genesis batch is imported
        require(finalizedStateRoots[0] == bytes32(0), "genesis batch imported");

        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        bytes32 _l1DataHash = BatchHeaderCodecV0.getDataHash(memPtr);
        uint256 _batchVersion = BatchHeaderCodecV0.getVersion(memPtr);

        // check all fields except `l1DataHash` and `lastBlockHash` are zero
        unchecked {
            uint256 sum = BatchHeaderCodecV0.getVersion(memPtr) +
                BatchHeaderCodecV0.getBatchIndex(memPtr) +
                BatchHeaderCodecV0.getL1MessagePopped(memPtr) +
                BatchHeaderCodecV0.getTotalL1MessagePopped(memPtr);
            require(sum == 0, "not all fields are zero");
        }
        require(
            BatchHeaderCodecV0.getDataHash(memPtr) != bytes32(0),
            "zero data hash"
        );
        require(
            BatchHeaderCodecV0.getParentBatchHash(memPtr) == bytes32(0),
            "nonzero parent batch hash"
        );

        require(
            BatchHeaderCodecV0.getBlobVersionedHash(memPtr) ==
                ZERO_VERSIONED_HASH,
            "invalid versioned hash"
        );

        batchBaseStore[0] = BatchBase(
            _batchHash,
            _batchVersion,
            block.timestamp,
            block.timestamp,
            0
        );
        batchDataStore[0] = BatchData(
            ZERO_VERSIONED_HASH,
            _l1DataHash,
            bytes32(0),
            _postStateRoot,
            bytes32(0),
            0,
            0,
            bytes("0x")
        );
        batchSignatureStore[0] = BatchSignature(bytes32(0), bytes32(0), "0x");
        finalizedStateRoots[0] = _postStateRoot;

        emit CommitBatch(0, _batchHash);
        emit FinalizeBatch(0, _batchHash, _postStateRoot, bytes32(0));
    }

    /// @inheritdoc IRollup
    function commitBatch(
        BatchDataInput calldata batchDataInput,
        BatchSignatureInput calldata batchSignatureInput
    ) external payable override OnlyStaker nonReqRevert whenNotPaused {
        require(batchDataInput.version == 0, "invalid version");
        // check whether the batch is empty
        uint256 _chunksLength = batchDataInput.chunks.length;
        require(_chunksLength > 0, "batch is empty");
        require(
            batchDataInput.prevStateRoot != bytes32(0),
            "previous state root is zero"
        );
        require(
            batchDataInput.postStateRoot != bytes32(0),
            "new state root is zero"
        );

        // The overall memory layout in this function is organized as follows
        // +---------------------+-------------------+------------------+
        // | parent batch header | chunk data hashes | new batch header |
        // +---------------------+-------------------+------------------+
        // ^                     ^                   ^
        // batchPtr              dataPtr             newBatchPtr (re-use var batchPtr)
        //
        // 1. We copy the parent batch header from calldata to memory starting at batchPtr
        // 2. We store `_chunksLength` number of Keccak hashes starting at `dataPtr`. Each Keccak
        //    hash corresponds to the data hash of a chunk. So we reserve the memory region from
        //    `dataPtr` to `dataPtr + _chunkLength * 32` for the chunk data hashes.
        // 3. The memory starting at `newBatchPtr` is used to store the new batch header and compute
        //    the batch hash.
        // the variable `batchPtr` will be reused later for the current batch
        (uint256 _batchPtr, bytes32 _parentBatchHash) = _loadBatchHeader(
            batchDataInput.parentBatchHeader
        );
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(_batchPtr);

        require(
            batchBaseStore[_batchIndex].batchHash == _parentBatchHash,
            "incorrect parent batch hash"
        );
        require(
            batchBaseStore[_batchIndex + 1].batchHash == bytes32(0),
            "batch already committed"
        );
        require(
            _batchIndex == lastCommittedBatchIndex,
            "incorrect batch index"
        );
        require(
            batchDataStore[_batchIndex].postStateRoot ==
                batchDataInput.prevStateRoot,
            "incorrect previous state root"
        );

        uint256 _totalL1MessagesPoppedOverall = BatchHeaderCodecV0
            .getTotalL1MessagePopped(_batchPtr);

        // load `dataPtr` and reserve the memory region for chunk data hashes
        uint256 dataPtr;
        assembly {
            dataPtr := mload(0x40)
            mstore(0x40, add(dataPtr, mul(_chunksLength, 32)))
        }
        // compute the data hash for each chunk
        uint256 _totalL1MessagesPoppedInBatch;
        for (uint256 i = 0; i < _chunksLength; i++) {
            uint256 _totalNumL1MessagesInChunk = _commitChunk(
                dataPtr,
                batchDataInput.chunks[i],
                _totalL1MessagesPoppedInBatch,
                _totalL1MessagesPoppedOverall,
                batchDataInput.skippedL1MessageBitmap
            );
            unchecked {
                _totalL1MessagesPoppedInBatch += _totalNumL1MessagesInChunk;
                _totalL1MessagesPoppedOverall += _totalNumL1MessagesInChunk;
                dataPtr += 32;
            }
        }
        // check the length of bitmap
        unchecked {
            require(
                ((_totalL1MessagesPoppedInBatch + 255) / 256) * 32 ==
                    batchDataInput.skippedL1MessageBitmap.length,
                "wrong bitmap length"
            );
        }
        // compute the data hash for current batch
        bytes32 _l1DataHash;
        assembly {
            let dataLen := mul(_chunksLength, 0x20)
            _l1DataHash := keccak256(sub(dataPtr, dataLen), dataLen)
            _batchPtr := mload(0x40) // reset batchPtr
            _batchIndex := add(_batchIndex, 1) // increase batch index
        }
        // store entries, the order matters
        BatchHeaderCodecV0.storeVersion(_batchPtr, batchDataInput.version);
        BatchHeaderCodecV0.storeBatchIndex(_batchPtr, _batchIndex);
        BatchHeaderCodecV0.storeL1MessagePopped(
            _batchPtr,
            _totalL1MessagesPoppedInBatch
        );
        BatchHeaderCodecV0.storeTotalL1MessagePopped(
            _batchPtr,
            _totalL1MessagesPoppedOverall
        );
        BatchHeaderCodecV0.storeDataHash(_batchPtr, _l1DataHash);
        BatchHeaderCodecV0.storeParentBatchHash(_batchPtr, _parentBatchHash);
        BatchHeaderCodecV0.storeSkippedBitmap(
            _batchPtr,
            batchDataInput.skippedL1MessageBitmap
        );

        bytes32 _blobVersionedHash = (blobhash(0) == bytes32(0))
            ? ZERO_VERSIONED_HASH
            : blobhash(0);

        BatchHeaderCodecV0.storeBlobVersionedHash(
            _batchPtr,
            _blobVersionedHash
        );

        batchBaseStore[_batchIndex] = BatchBase(
            BatchHeaderCodecV0.computeBatchHash(
                _batchPtr,
                BatchHeaderCodecV0.BATCH_HEADER_FIXED_LENGTH +
                    batchDataInput.skippedL1MessageBitmap.length
            ),
            batchDataInput.version,
            block.timestamp,
            block.timestamp + finalizationPeriodSeconds,
            _loadL2BlockNumber(batchDataInput.chunks[_chunksLength - 1])
        );
        batchDataStore[_batchIndex] = BatchData(
            _blobVersionedHash,
            _l1DataHash,
            batchDataInput.prevStateRoot,
            batchDataInput.postStateRoot,
            batchDataInput.withdrawalRoot,
            _totalL1MessagesPoppedInBatch,
            _totalL1MessagesPoppedOverall,
            batchDataInput.skippedL1MessageBitmap
        );
        address[] memory submitter = new address[](1);
        submitter[0] = _msgSender();
        batchSignatureStore[_batchIndex] = BatchSignature(
            _getBLSMsgHash(batchDataInput),
            keccak256(batchSignatureInput.sequencerSets),
            // Before BLS is implemented, the accuracy of the sequencer set uploaded by rollup cannot be guaranteed.
            // Therefore, if the batch is successfully challenged, only the submitter will be punished.
            abi.encode(submitter) // => batchSignature.signedSequencers
        );

        lastCommittedBatchIndex = _batchIndex;

        require(
            _checkSequencerSetVerifyHash(
                batchDataInput,
                batchSignatureStore[_batchIndex].sequencerSetVerifyHash
            ),
            "the sequencer set verification failed"
        );

        // verify bls signature
        require(
            IL1Staking(l1StakingContract).verifySignature(
                abi.decode(batchSignatureInput.signedSequencers, (address[])),
                _getValidSequencerSet(batchSignatureInput.sequencerSets, 0),
                batchSignatureStore[_batchIndex].blsMsgHash,
                batchSignatureInput.signature
            ),
            "the signature verification failed"
        );

        emit CommitBatch(_batchIndex, batchBaseStore[_batchIndex].batchHash);
    }

    /// @inheritdoc IRollup
    /// @dev If the owner wants to revert a sequence of batches by sending multiple transactions,
    ///      make sure to revert recent batches first.
    function revertBatch(
        bytes calldata _batchHeader,
        uint256 _count
    ) external onlyOwner {
        require(_count > 0, "count must be nonzero");

        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        // check batch hash
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);

        require(
            batchBaseStore[_batchIndex].batchHash == _batchHash,
            "incorrect batch hash"
        );
        // make sure no gap is left when reverting from the ending to the beginning.
        require(
            batchBaseStore[_batchIndex + _count].batchHash == bytes32(0),
            "reverting must start from the ending"
        );
        // check finalization
        require(
            _batchIndex > lastFinalizedBatchIndex,
            "can only revert unFinalized batch"
        );

        lastCommittedBatchIndex = _batchIndex - 1;
        while (_count > 0) {
            batchBaseStore[_batchIndex].batchHash = bytes32(0);
            if (revertReqIndex > 0 && _batchIndex == revertReqIndex) {
                // if challenge exist and not finished yet, return challenge deposit to challenger
                if (!challenges[_batchIndex].finished) {
                    batchChallengeReward[
                        challenges[_batchIndex].challenger
                    ] += challenges[_batchIndex].challengeDeposit;
                }
                delete challenges[_batchIndex];
                revertReqIndex = 0;
            }
            emit RevertBatch(_batchIndex, _batchHash);

            unchecked {
                _batchIndex += 1;
                _count -= 1;
            }
            _batchHash = batchBaseStore[_batchIndex].batchHash;
            if (_batchHash == bytes32(0)) {
                break;
            }
        }
    }

    /// @dev challengeState challenges a batch by submitting a deposit.
    function challengeState(
        uint64 batchIndex
    ) external payable onlyChallenger nonReqRevert whenNotPaused {
        require(!inChallenge, "already in challenge");
        require(
            lastFinalizedBatchIndex < batchIndex,
            "batch already finalized"
        );
        require(batchBaseStore[batchIndex].batchHash != 0, "batch not exist");
        require(
            challenges[batchIndex].challenger == address(0),
            "batch already challenged"
        );
        // check challenge window
        require(
            batchInsideChallengeWindow(batchIndex),
            "cannot challenge batch outside the challenge window"
        );
        // check challenge amount
        require(
            msg.value >= IL1Staking(l1StakingContract).stakingValue(),
            "insufficient value"
        );

        batchChallenged = batchIndex;
        challenges[batchIndex] = BatchChallenge(
            batchIndex,
            _msgSender(),
            msg.value,
            block.timestamp,
            false,
            false
        );
        emit ChallengeState(batchIndex, _msgSender(), msg.value);

        for (
            uint256 i = lastFinalizedBatchIndex + 1;
            i <= lastCommittedBatchIndex;
            i++
        ) {
            if (i != batchIndex) {
                batchBaseStore[i].finalizeTimestamp += proofWindow;
            }
        }

        inChallenge = true;
    }

    /// @notice Update proofWindow.
    /// @param _newWindow New proof window.
    function updateProofWindow(uint256 _newWindow) external onlyOwner {
        require(
            _newWindow > 0 && _newWindow != proofWindow,
            "invalid new proof window"
        );
        uint256 _oldProofWindow = proofWindow;
        proofWindow = _newWindow;
        emit UpdateProofWindow(_oldProofWindow, proofWindow);
    }

    /// @notice Update finalizationPeriodSeconds.
    /// @param _newPeriod New finalize period seconds.
    function updateFinalizePeriodSeconds(
        uint256 _newPeriod
    ) external onlyOwner {
        require(
            _newPeriod > 0 && _newPeriod != finalizationPeriodSeconds,
            "invalid new finalize period"
        );
        uint256 _oldFinalizationPeriodSeconds = finalizationPeriodSeconds;
        finalizationPeriodSeconds = _newPeriod;
        emit UpdateFinalizationPeriodSeconds(
            _oldFinalizationPeriodSeconds,
            finalizationPeriodSeconds
        );
    }

    /// @notice Add an account to the challenger list.
    /// @param _account The address of account to add.
    function addChallenger(address _account) external onlyOwner {
        require(!isChallenger[_account], "account is already a challenger");
        isChallenger[_account] = true;
        emit UpdateChallenger(_account, true);
    }

    /// @notice Remove an account from the challenger list.
    /// @param _account The address of account to remove.
    function removeChallenger(address _account) external onlyOwner {
        require(isChallenger[_account], "account is not a challenger");
        isChallenger[_account] = false;
        emit UpdateChallenger(_account, false);
    }

    /// @notice Update the address verifier contract.
    /// @param _newVerifier The address of new verifier contract.
    function updateVerifier(address _newVerifier) external onlyOwner {
        require(
            _newVerifier != address(0) && _newVerifier != verifier,
            "invalid new verifier"
        );
        address _oldVerifier = verifier;
        verifier = _newVerifier;
        emit UpdateVerifier(_oldVerifier, _newVerifier);
    }

    /// @notice Update the value of `maxNumTxInChunk`.
    /// @param _maxNumTxInChunk The new value of `maxNumTxInChunk`.
    function updateMaxNumTxInChunk(
        uint256 _maxNumTxInChunk
    ) external onlyOwner {
        require(
            _maxNumTxInChunk > 0 && _maxNumTxInChunk != maxNumTxInChunk,
            "invalid new maxNumTxInChunk"
        );
        uint256 _oldMaxNumTxInChunk = maxNumTxInChunk;
        maxNumTxInChunk = _maxNumTxInChunk;
        emit UpdateMaxNumTxInChunk(_oldMaxNumTxInChunk, _maxNumTxInChunk);
    }

    /// @notice Pause the contract
    /// @param _status The pause status to update.
    function setPause(bool _status) external onlyOwner {
        if (_status) {
            _pause();
            // if challenge exist and not finished yet, return challenge deposit to challenger
            if (inChallenge && !challenges[batchChallenged].finished) {
                batchChallengeReward[
                    challenges[batchChallenged].challenger
                ] += challenges[batchChallenged].challengeDeposit;
                delete challenges[batchChallenged];
            }
            inChallenge = false;
        } else {
            _unpause();
        }
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @dev proveState proves a batch by submitting a proof.
    /// _kzgData: [y(32) | commitment(48) | proof(48)]
    function proveState(
        uint64 _batchIndex,
        bytes calldata _aggrProof,
        bytes calldata _kzgDataProof
    ) external nonReqRevert whenNotPaused {
        // Ensure challenge exists and is not finished
        require(batchInChallenge(_batchIndex), "batch in challenge");

        // Mark challenge as finished
        challenges[_batchIndex].finished = true;
        inChallenge = false;

        // Check for timeout
        if (
            challenges[_batchIndex].startTime + proofWindow <= block.timestamp
        ) {
            // set status
            challenges[_batchIndex].challengeSuccess = true;
            _challengerWin(
                _batchIndex,
                batchSignatureStore[_batchIndex].signedSequencers,
                "Timeout"
            );
        } else {
            _verifyProof(_batchIndex, _aggrProof, _kzgDataProof);
            // Record defender win
            _defenderWin(_batchIndex, _msgSender(), "Proof success");
        }
    }

    /// @dev finalize batch
    function finalizeBatch(
        uint256 _batchIndex
    ) public nonReqRevert whenNotPaused {
        require(batchExist(_batchIndex), "batch not exist");
        require(!batchInChallenge(_batchIndex), "batch in challenge");
        require(!batchChallengedSuccess(_batchIndex), "batch should be revert");
        require(
            !batchInsideChallengeWindow(_batchIndex),
            "batch in challenge window"
        );
        // verify previous state root.
        require(
            finalizedStateRoots[_batchIndex - 1] ==
                batchDataStore[_batchIndex].prevStateRoot,
            "incorrect previous state root"
        );
        // avoid duplicated verification
        require(
            finalizedStateRoots[_batchIndex] == bytes32(0),
            "batch already verified"
        );
        // check and update lastFinalizedBatchIndex
        unchecked {
            require(
                lastFinalizedBatchIndex + 1 == _batchIndex,
                "incorrect batch index"
            );
            lastFinalizedBatchIndex = _batchIndex;
        }
        withdrawalRoots[batchDataStore[_batchIndex].withdrawalRoot] = true;
        // record state root and withdrawal root
        finalizedStateRoots[_batchIndex] = batchDataStore[_batchIndex]
            .postStateRoot;

        // Pop finalized and non-skipped message from L1MessageQueue.
        uint256 _l1MessagePopped = batchDataStore[_batchIndex].l1MessagePopped;
        if (_l1MessagePopped > 0) {
            IL1MessageQueue _queue = IL1MessageQueue(messageQueue);
            bytes memory _skippedL1MessageBitmap = batchDataStore[_batchIndex]
                .skippedL1MessageBitmap;
            uint256 bitmapPtr;
            assembly {
                bitmapPtr := add(
                    _skippedL1MessageBitmap,
                    /*BYTES_HEADER_SIZE*/ 32
                )
            }

            unchecked {
                uint256 _startIndex = batchDataStore[_batchIndex]
                    .totalL1MessagePopped - _l1MessagePopped;
                for (uint256 i = 0; i < _l1MessagePopped; i += 256) {
                    uint256 _count = 256;
                    if (_l1MessagePopped - i < _count) {
                        _count = _l1MessagePopped - i;
                    }

                    uint256 _skippedBitmap;
                    uint256 _index = i / 256;
                    assembly {
                        _skippedBitmap := mload(add(bitmapPtr, mul(_index, 32)))
                    }

                    _queue.popCrossDomainMessage(
                        _startIndex,
                        _count,
                        _skippedBitmap
                    );

                    _startIndex += 256;
                }
            }
        }
        delete batchBaseStore[_batchIndex - 1];
        delete batchDataStore[_batchIndex - 1];
        delete batchSignatureStore[_batchIndex - 1];
        delete challenges[_batchIndex - 1];

        emit FinalizeBatch(
            _batchIndex,
            batchBaseStore[_batchIndex].batchHash,
            batchDataStore[_batchIndex].postStateRoot,
            batchDataStore[_batchIndex].withdrawalRoot
        );
    }

    /// @notice Claim challenge reward
    /// @param receiver The receiver address
    function claimReward(address receiver) external {
        uint256 amount = batchChallengeReward[_msgSender()];
        require(amount != 0, "invalid batchChallengeReward");
        delete batchChallengeReward[_msgSender()];
        _transfer(receiver, amount);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IRollup
    function isBatchFinalized(
        uint256 _batchIndex
    ) external view override returns (bool) {
        return _batchIndex <= lastFinalizedBatchIndex;
    }

    /// @inheritdoc IRollup
    function committedBatches(
        uint256 batchIndex
    ) external view override returns (bytes32) {
        return batchBaseStore[batchIndex].batchHash;
    }

    /// @dev Public function to checks whether the batch is in challenge.
    /// @param batchIndex The index of the batch to be checked.
    function batchInChallenge(uint256 batchIndex) public view returns (bool) {
        return
            challenges[batchIndex].challenger != address(0) &&
            !challenges[batchIndex].finished;
    }

    /// @dev Retrieves the success status of a batch challenge.
    /// @param batchIndex The index of the batch to check.
    function batchChallengedSuccess(
        uint256 batchIndex
    ) public view returns (bool) {
        return challenges[batchIndex].challengeSuccess;
    }

    /// @dev Public function to checks whether batch exists.
    /// @param batchIndex The index of the batch to be checked.
    function batchExist(uint256 batchIndex) public view returns (bool) {
        return
            batchBaseStore[batchIndex].originTimestamp > 0 &&
            batchBaseStore[batchIndex].batchHash != bytes32(0);
    }

    /// @dev Public function to checks whether the batch is in challengeWindow.
    /// @param batchIndex The index of the batch to be checked.
    function batchInsideChallengeWindow(
        uint256 batchIndex
    ) public view returns (bool) {
        return batchBaseStore[batchIndex].finalizeTimestamp > block.timestamp;
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev todo
    function _verifyProof(
        uint64 _batchIndex,
        bytes calldata _aggrProof,
        bytes calldata _kzgDataProof
    ) private view {
        // Check validity of proof
        require(_aggrProof.length > 0, "Invalid aggregation proof");

        // Check validity of KZG data
        require(_kzgDataProof.length == 160, "Invalid KZG data proof");

        // Calls the point evaluation precompile and verifies the output
        {
            (bool success, bytes memory data) = POINT_EVALUATION_PRECOMPILE_ADDR
                .staticcall(
                    abi.encodePacked(
                        batchDataStore[_batchIndex].blobVersionedHash,
                        _kzgDataProof
                    )
                );
            // We verify that the point evaluation precompile call was successful by testing the latter 32 bytes of the
            // response is equal to BLS_MODULUS as defined in https://eips.ethereum.org/EIPS/eip-4844#point-evaluation-precompile
            require(success, "failed to call point evaluation precompile");
            (, uint256 result) = abi.decode(data, (uint256, uint256));
            require(result == BLS_MODULUS, "precompile unexpected output");
        }

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                LAYER_2_CHAIN_ID,
                batchDataStore[_batchIndex].prevStateRoot,
                batchDataStore[_batchIndex].postStateRoot,
                batchDataStore[_batchIndex].withdrawalRoot,
                batchDataStore[_batchIndex].l1DataHash,
                _kzgDataProof[0:64],
                batchDataStore[_batchIndex].blobVersionedHash
            )
        );

        IRollupVerifier(verifier).verifyAggregateProof(
            batchBaseStore[_batchIndex].batchVersion,
            _batchIndex,
            _aggrProof,
            _publicInputHash
        );
    }

    /// @dev Internal function to compute BLS msg hash
    function _getBLSMsgHash(
        BatchDataInput calldata // batchDataInput
    ) internal pure returns (bytes32) {
        // TODO compute bls message hash
        return bytes32(0);
    }

    /// @dev Internal function to compute BLS msg hash
    function _checkSequencerSetVerifyHash(
        BatchDataInput calldata, // batchDataInput
        bytes32 // sequencerSetVerifyHash
    ) internal pure returns (bool) {
        // TODO check sequencerSetVerifyHash in batch
        return true;
    }

    /// @dev todo
    function _getValidSequencerSet(
        bytes calldata sequencerSets,
        uint256 blockHeight
    ) internal pure returns (address[] memory) {
        // TODO require submitter was in valid sequencer set after BLS was implementated
        (
            ,
            address[] memory sequencerSet0,
            uint256 blockHeight1,
            address[] memory sequencerSet1,
            uint256 blockHeight2,
            address[] memory sequencerSet2
        ) = abi.decode(
                sequencerSets,
                (uint256, address[], uint256, address[], uint256, address[])
            );
        if (blockHeight >= blockHeight2) {
            return sequencerSet2;
        }
        if (blockHeight >= blockHeight1) {
            return sequencerSet1;
        }
        return sequencerSet0;
    }

    /// @dev Internal function executed when the defender wins.
    /// @param batchIndex   The index of the batch indicating where the challenge occurred.
    /// @param prover       The zkProof prover address.
    /// @param _type        Description of the challenge type.
    function _defenderWin(
        uint64 batchIndex,
        address prover,
        string memory _type
    ) internal {
        uint256 challengeDeposit = challenges[batchIndex].challengeDeposit;
        batchChallengeReward[prover] += challengeDeposit;
        emit ChallengeRes(batchIndex, prover, _type);
    }

    /// @dev Internal function executed when the challenger wins.
    /// @param batchIndex   The index of the batch indicating where the challenge occurred.
    /// @param sequencers   An array containing the sequencers to be slashed.
    /// @param _type        Description of the challenge type.
    function _challengerWin(
        uint64 batchIndex,
        bytes memory sequencers,
        string memory _type
    ) internal {
        revertReqIndex = batchIndex;
        address challenger = challenges[batchIndex].challenger;
        uint256 reward = IL1Staking(l1StakingContract).slash(
            abi.decode(sequencers, (address[]))
        );
        batchChallengeReward[challenges[batchIndex].challenger] += (challenges[
            batchIndex
        ].challengeDeposit + reward);
        emit ChallengeRes(batchIndex, challenger, _type);
    }

    /// @dev Internal function to transfer ETH to a specified address.
    /// @param _to      The address to transfer ETH to.
    /// @param _amount  The amount of ETH to transfer.
    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}("");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /// @dev Internal function to load batch header from calldata to memory.
    /// @param _batchHeader The batch header in calldata.
    /// @return _memPtr     The start memory offset of loaded batch header.
    /// @return _batchHash  The hash of the loaded batch header.
    function _loadBatchHeader(
        bytes calldata _batchHeader
    ) internal pure returns (uint256 _memPtr, bytes32 _batchHash) {
        // load to memory
        uint256 _length;
        (_memPtr, _length) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);

        // compute batch hash
        _batchHash = BatchHeaderCodecV0.computeBatchHash(_memPtr, _length);
    }

    /// @dev Internal function to load the latestL2BlockNumber.
    /// @param _chunk The batch chunk in memory.
    function _loadL2BlockNumber(
        bytes memory _chunk
    ) internal pure returns (uint256) {
        uint256 blockPtr;
        uint256 chunkPtr;
        assembly {
            chunkPtr := add(_chunk, 0x20)
            blockPtr := add(chunkPtr, 1)
        }
        uint256 _numBlocks = ChunkCodecV0.validateChunkLength(
            chunkPtr,
            _chunk.length
        );
        for (uint256 i = 0; i < _numBlocks - 1; i++) {
            unchecked {
                blockPtr += ChunkCodecV0.BLOCK_CONTEXT_LENGTH;
            }
        }
        uint256 l2BlockNumber = ChunkCodecV0.getBlockNumber(blockPtr);
        return l2BlockNumber;
    }

    /// @dev Internal function to commit a chunk.
    /// @param _memPtr                          The start memory offset to store list of `dataHash`.
    /// @param _chunk                           The encoded chunk to commit.
    /// @param _totalL1MessagesPoppedInBatch    The total number of L1 messages popped in current batch.
    /// @param _totalL1MessagesPoppedOverall    The total number of L1 messages popped in all batches including current batch.
    /// @param _skippedL1MessageBitmap          The bitmap indicates whether each L1 message is skipped or not.
    /// @return _totalNumL1MessagesInChunk      The total number of L1 message popped in current chunk
    function _commitChunk(
        uint256 _memPtr,
        bytes memory _chunk,
        uint256 _totalL1MessagesPoppedInBatch,
        uint256 _totalL1MessagesPoppedOverall,
        bytes calldata _skippedL1MessageBitmap
    ) internal view returns (uint256 _totalNumL1MessagesInChunk) {
        uint256 chunkPtr;
        uint256 startDataPtr;
        uint256 dataPtr;
        uint256 blockPtr;

        assembly {
            dataPtr := mload(0x40)
            startDataPtr := dataPtr
            chunkPtr := add(_chunk, 0x20) // skip chunkLength
            blockPtr := add(chunkPtr, 1) // skip numBlocks
        }

        uint256 _numBlocks = ChunkCodecV0.validateChunkLength(
            chunkPtr,
            _chunk.length
        );

        // concatenate block contexts, use scope to avoid stack too deep
        {
            uint256 _totalTransactionsInChunk;
            for (uint256 i = 0; i < _numBlocks; i++) {
                dataPtr = ChunkCodecV0.copyBlockContext(chunkPtr, dataPtr, i);
                uint256 _numTransactionsInBlock = ChunkCodecV0
                    .getNumTransactions(blockPtr);
                unchecked {
                    _totalTransactionsInChunk += _numTransactionsInBlock;
                    blockPtr += ChunkCodecV0.BLOCK_CONTEXT_LENGTH;
                }
            }
            assembly {
                mstore(0x40, add(dataPtr, mul(_totalTransactionsInChunk, 0x20))) // reserve memory for tx hashes
            }
        }

        // It is used to compute the actual number of transactions in chunk.
        uint256 txHashStartDataPtr;
        assembly {
            txHashStartDataPtr := dataPtr
            blockPtr := add(chunkPtr, 1) // reset block ptr
        }

        // concatenate tx hashes
        while (_numBlocks > 0) {
            // concatenate l1 message hashes
            uint256 _numL1MessagesInBlock = ChunkCodecV0.getNumL1Messages(
                blockPtr
            );
            dataPtr = _loadL1MessageHashes(
                dataPtr,
                _numL1MessagesInBlock,
                _totalL1MessagesPoppedInBatch,
                _totalL1MessagesPoppedOverall,
                _skippedL1MessageBitmap
            );

            // concatenate l2 transaction hashes
            require(
                ChunkCodecV0.getNumTransactions(blockPtr) >=
                    _numL1MessagesInBlock,
                "num txs less than num L1 msgs"
            );

            unchecked {
                _totalNumL1MessagesInChunk += _numL1MessagesInBlock;
                _totalL1MessagesPoppedInBatch += _numL1MessagesInBlock;
                _totalL1MessagesPoppedOverall += _numL1MessagesInBlock;

                _numBlocks -= 1;
                blockPtr += ChunkCodecV0.BLOCK_CONTEXT_LENGTH;
            }
        }

        // check the actual number of transactions in the chunk
        require(
            (dataPtr - txHashStartDataPtr) / 32 <= maxNumTxInChunk,
            "too many txs in one chunk"
        );

        // compute data hash and store to memory
        assembly {
            let dataHash := keccak256(startDataPtr, sub(dataPtr, startDataPtr))
            mstore(_memPtr, dataHash)
        }

        return _totalNumL1MessagesInChunk;
    }

    /// @dev Internal function to load L1 message hashes from the message queue.
    /// @param _ptr                             The memory offset to store the transaction hash.
    /// @param _numL1Messages                   The number of L1 messages to load.
    /// @param _totalL1MessagesPoppedInBatch    The total number of L1 messages popped in current batch.
    /// @param _totalL1MessagesPoppedOverall    The total number of L1 messages popped in all batches including current batch.
    /// @param _skippedL1MessageBitmap          The bitmap indicates whether each L1 message is skipped or not.
    /// @return uint256                         The new memory offset after loading.
    function _loadL1MessageHashes(
        uint256 _ptr,
        uint256 _numL1Messages,
        uint256 _totalL1MessagesPoppedInBatch,
        uint256 _totalL1MessagesPoppedOverall,
        bytes calldata _skippedL1MessageBitmap
    ) internal view returns (uint256) {
        if (_numL1Messages == 0) {
            return _ptr;
        }
        IL1MessageQueue _messageQueue = IL1MessageQueue(messageQueue);

        unchecked {
            uint256 _bitmap;
            uint256 rem;
            for (uint256 i = 0; i < _numL1Messages; i++) {
                uint256 quo = _totalL1MessagesPoppedInBatch >> 8;
                rem = _totalL1MessagesPoppedInBatch & 0xff;

                // load bitmap every 256 bits
                if (i == 0 || rem == 0) {
                    assembly {
                        _bitmap := calldataload(
                            add(_skippedL1MessageBitmap.offset, mul(0x20, quo))
                        )
                    }
                }
                if (((_bitmap >> rem) & 1) == 0) {
                    // message not skipped
                    bytes32 _hash = _messageQueue.getCrossDomainMessage(
                        _totalL1MessagesPoppedOverall
                    );
                    assembly {
                        mstore(_ptr, _hash)
                        _ptr := add(_ptr, 0x20)
                    }
                }

                _totalL1MessagesPoppedInBatch += 1;
                _totalL1MessagesPoppedOverall += 1;
            }

            // check last L1 message is not skipped, _totalL1MessagesPoppedInBatch must > 0
            rem = (_totalL1MessagesPoppedInBatch - 1) & 0xff;
            require(((_bitmap >> rem) & 1) == 0, "cannot skip last L1 message");
        }

        return _ptr;
    }
}
