// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {BatchHeaderCodecV0} from "../../libraries/codec/BatchHeaderCodecV0.sol";
import {BatchHeaderCodecV1} from "../../libraries/codec/BatchHeaderCodecV1.sol";
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
    bytes32 internal constant ZERO_VERSIONED_HASH = 0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;

    /// @notice The BLS MODULUS. Deprecated.
    uint256 internal constant __BLS_MODULUS =
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

    /// @notice The maximum number of transactions allowed in each chunk. Deprecated.
    uint256 public __maxNumTxInChunk;

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
    mapping(uint256 batchIndex => bytes32 stateRoot) public override finalizedStateRoots;

    /// @notice Store committed batch hash.
    mapping(uint256 batchIndex => bytes32 batchHash) public override committedBatches;

    /// @notice Store committed batch base.
    mapping(uint256 batchIndex => BatchData) public batchDataStore;

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

    /// @notice percentage awarded to prover
    uint256 public proofRewardPercent;

    /// @notice prove remaining
    uint256 public proveRemaining;

    /// @notice committedStateRoots
    mapping(uint256 batchIndex => bytes32 stateRoot) public committedStateRoots;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Only active staker allowed.
    modifier onlyActiveStaker() {
        require(IL1Staking(l1StakingContract).isActiveStaker(_msgSender()), "only active staker allowed");
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
    /// @param _chainID The chain ID
    constructor(uint64 _chainID) {
        LAYER_2_CHAIN_ID = _chainID;
        _disableInitializers();
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
    /// @param _finalizationPeriodSeconds finalization period seconds
    /// @param _proofWindow               proof window
    function initialize(
        address _l1StakingContract,
        address _messageQueue,
        address _verifier,
        uint256 _finalizationPeriodSeconds,
        uint256 _proofWindow,
        uint256 _proofRewardPercent
    ) public initializer {
        if (_messageQueue == address(0) || _verifier == address(0)) {
            revert ErrZeroAddress();
        }
        require(_l1StakingContract != address(0), "invalid l1 staking contract");

        __Pausable_init();
        __Ownable_init();

        l1StakingContract = _l1StakingContract;
        messageQueue = _messageQueue;
        verifier = _verifier;
        finalizationPeriodSeconds = _finalizationPeriodSeconds;
        proofWindow = _proofWindow;
        proofRewardPercent = _proofRewardPercent;
        emit UpdateVerifier(address(0), _verifier);
        emit UpdateFinalizationPeriodSeconds(0, _finalizationPeriodSeconds);
        emit UpdateProofWindow(0, _proofWindow);
        emit UpdateProofRewardPercent(0, _proofRewardPercent);
    }

    function initialize2(bytes32 _prevStateRoot) external reinitializer(2) {
        require(_getInitializedVersion() == 2, "must have initialized!");
        require(_prevStateRoot != bytes32(0), "can not set state root with bytes32(0)!");

        if (committedStateRoots[lastCommittedBatchIndex] == bytes32(0)) {
            committedStateRoots[lastCommittedBatchIndex] = _prevStateRoot;
        }
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice Import layer 2 genesis block
    function importGenesisBatch(bytes calldata _batchHeader) external onlyOwner {
        // check whether the genesis batch is imported
        require(finalizedStateRoots[0] == bytes32(0), "genesis batch imported");

        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        // check batch index is 0
        require(_batchIndex == 0, "invalid batch index");
        bytes32 _postStateRoot = BatchHeaderCodecV0.getPostStateHash(memPtr);
        require(_postStateRoot != bytes32(0), "zero state root");
        // check all fields except `dataHash` and `lastBlockHash` are zero
        require(BatchHeaderCodecV0.getL1MessagePopped(memPtr) == 0, "l1 message popped should be 0");
        require(BatchHeaderCodecV0.getDataHash(memPtr) != bytes32(0), "zero data hash");
        require(BatchHeaderCodecV0.getBlobVersionedHash(memPtr) == ZERO_VERSIONED_HASH, "invalid versioned hash");

        committedBatches[_batchIndex] = _batchHash;
        batchDataStore[_batchIndex] = BatchData(block.timestamp, block.timestamp, 0, 0);

        committedStateRoots[_batchIndex] = _postStateRoot;
        finalizedStateRoots[_batchIndex] = _postStateRoot;
        lastCommittedBatchIndex = _batchIndex;
        lastFinalizedBatchIndex = _batchIndex;

        emit CommitBatch(_batchIndex, _batchHash);
        emit FinalizeBatch(_batchIndex, _batchHash, _postStateRoot, bytes32(0));
    }

    /// @inheritdoc IRollup
    function commitBatch(
        BatchDataInput calldata batchDataInput,
        BatchSignatureInput calldata batchSignatureInput
    ) external payable override onlyActiveStaker nonReqRevert whenNotPaused {
        require(batchDataInput.version == 0 || batchDataInput.version == 1, "invalid version");
        require(batchDataInput.prevStateRoot != bytes32(0), "previous state root is zero");
        require(batchDataInput.postStateRoot != bytes32(0), "new state root is zero");

        // The overall memory layout in this function is organized as follows
        // +---------------------+-------------------+------------------+
        // | parent batch header | batch data hashes | new batch header |
        // +---------------------+-------------------+------------------+
        // ^                     ^                   ^
        // batchPtr              dataPtr             newBatchPtr (re-use var batchPtr)
        //
        // 1. We copy the parent batch header from calldata to memory starting at batchPtr
        // 2. we store the batch data hash
        // 3. The memory starting at `newBatchPtr` is used to store the new batch header and compute
        //    the batch hash.
        // the variable `batchPtr` will be reused later for the current batch
        (uint256 _batchPtr, bytes32 _parentBatchHash) = _loadBatchHeader(batchDataInput.parentBatchHeader);
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(_batchPtr);
        require(committedBatches[_batchIndex + 1] == bytes32(0), "batch already committed");
        require(_batchIndex == lastCommittedBatchIndex, "incorrect batch index");
        require(committedBatches[_batchIndex] == _parentBatchHash, "incorrect parent batch hash");
        require(committedStateRoots[_batchIndex] == batchDataInput.prevStateRoot, "incorrect previous state root");

        uint256 _totalL1MessagesPoppedOverall = BatchHeaderCodecV0.getTotalL1MessagePopped(_batchPtr);
        // compute the data hash for batch
        bytes32 dataHash = _commitBatch(
            batchDataInput.lastBlockNumber,
            batchDataInput.numL1Messages,
            _totalL1MessagesPoppedOverall
        );

        unchecked {
            _totalL1MessagesPoppedOverall += batchDataInput.numL1Messages;
        }
        assembly {
            _batchIndex := add(_batchIndex, 1) // increase batch index
        }
        bytes32 _blobVersionedHash = (blobhash(0) == bytes32(0)) ? ZERO_VERSIONED_HASH : blobhash(0);

        {            
            uint256 _headerLength = BatchHeaderCodecV0.BATCH_HEADER_LENGTH;
            if (batchDataInput.version == 1) {
                _headerLength = BatchHeaderCodecV1.BATCH_HEADER_LENGTH;
            }
            assembly {
                _batchPtr := mload(0x40)
                mstore(0x40, add(_batchPtr, _headerLength))
            }

            // store entries, the order matters
            BatchHeaderCodecV0.storeVersion(_batchPtr, batchDataInput.version);
            BatchHeaderCodecV0.storeBatchIndex(_batchPtr, _batchIndex);
            BatchHeaderCodecV0.storeL1MessagePopped(_batchPtr, batchDataInput.numL1Messages);
            BatchHeaderCodecV0.storeTotalL1MessagePopped(_batchPtr, _totalL1MessagesPoppedOverall);
            BatchHeaderCodecV0.storeDataHash(_batchPtr, dataHash);
            BatchHeaderCodecV0.storeBlobVersionedHash(_batchPtr, _blobVersionedHash);
            BatchHeaderCodecV0.storePrevStateHash(_batchPtr, batchDataInput.prevStateRoot);
            BatchHeaderCodecV0.storePostStateHash(_batchPtr, batchDataInput.postStateRoot);
            BatchHeaderCodecV0.storeWithdrawRootHash(_batchPtr, batchDataInput.withdrawalRoot);
            BatchHeaderCodecV0.storeSequencerSetVerifyHash(_batchPtr, keccak256(batchSignatureInput.sequencerSets));
            BatchHeaderCodecV0.storeParentBatchHash(_batchPtr, _parentBatchHash);
            // store last block number if version >= 1
            if (batchDataInput.version >= 1) {
                BatchHeaderCodecV1.storeLastBlockNumber(_batchPtr, batchDataInput.lastBlockNumber);
            }
            committedBatches[_batchIndex] = BatchHeaderCodecV0.computeBatchHash(_batchPtr, _headerLength);
            committedStateRoots[_batchIndex] = batchDataInput.postStateRoot;
            uint256 proveRemainingTime = 0;
            if (inChallenge) {
                // Make the batch finalize time longer than the time required for the current challenge
                proveRemainingTime = proofWindow + challenges[batchChallenged].startTime - block.timestamp;
            }
            // storage batch data for challenge status check
            batchDataStore[_batchIndex] = BatchData(
                block.timestamp,
                block.timestamp + finalizationPeriodSeconds + proveRemainingTime,
                batchDataInput.lastBlockNumber,
                // Before BLS is implemented, the accuracy of the sequencer set uploaded by rollup cannot be guaranteed.
                // Therefore, if the batch is successfully challenged, only the submitter will be punished.
                IL1Staking(l1StakingContract).getStakerBitmap(_msgSender()) // => batchSignature.signedSequencersBitmap
            );

            lastCommittedBatchIndex = _batchIndex;
        }

        // verify bls signature
        require(
            IL1Staking(l1StakingContract).verifySignature(
                batchSignatureInput.signedSequencersBitmap,
                _getValidSequencerSet(batchSignatureInput.sequencerSets, 0),
                _getBLSMsgHash(batchDataInput),
                batchSignatureInput.signature
            ),
            "the signature verification failed"
        );
        emit CommitBatch(_batchIndex, committedBatches[_batchIndex]);
    }

    /// @inheritdoc IRollup
    /// @dev If the owner wants to revert a sequence of batches by sending multiple transactions,
    ///      make sure to revert recent batches first.
    function revertBatch(bytes calldata _batchHeader, uint256 _count) external onlyOwner {
        require(_count > 0, "count must be nonzero");

        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        // check batch hash
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        require(committedBatches[_batchIndex] == _batchHash, "incorrect batch hash");

        // make sure no gap is left when reverting from the ending to the beginning.
        require(committedBatches[_batchIndex + _count] == bytes32(0), "reverting must start from the ending");
        // check finalization
        require(_batchIndex > lastFinalizedBatchIndex, "can only revert unFinalized batch");

        lastCommittedBatchIndex = _batchIndex - 1;
        while (_count > 0) {
            emit RevertBatch(_batchIndex, _batchHash);

            committedBatches[_batchIndex] = bytes32(0);
            // if challenge exist and not finished yet, return challenge deposit to challenger
            if (batchInChallenge(_batchIndex)) {
                batchChallengeReward[challenges[_batchIndex].challenger] += challenges[_batchIndex].challengeDeposit;
                inChallenge = false;
            }
            delete challenges[_batchIndex];

            if (revertReqIndex > 0 && _batchIndex == revertReqIndex) {
                revertReqIndex = 0;
            }

            unchecked {
                _batchIndex += 1;
                _count -= 1;
            }
            _batchHash = committedBatches[_batchIndex];
            if (_batchHash == bytes32(0)) {
                break;
            }
        }
    }

    /// @dev challengeState challenges a batch by submitting a deposit.
    function challengeState(
        uint64 batchIndex,
        bytes32 _batchHash
    ) external payable onlyChallenger nonReqRevert whenNotPaused {
        require(!inChallenge, "already in challenge");
        require(lastFinalizedBatchIndex < batchIndex, "batch already finalized");
        require(committedBatches[batchIndex] == _batchHash, "incorrect batch hash");
        require(batchExist(batchIndex), "batch not exist");
        require(challenges[batchIndex].challenger == address(0), "batch already challenged");
        // check challenge window
        require(batchInsideChallengeWindow(batchIndex), "cannot challenge batch outside the challenge window");
        // check challenge amount
        require(msg.value >= IL1Staking(l1StakingContract).challengeDeposit(), "insufficient value");

        batchChallenged = batchIndex;
        challenges[batchIndex] = BatchChallenge(batchIndex, _msgSender(), msg.value, block.timestamp, false, false);
        emit ChallengeState(batchIndex, _msgSender(), msg.value);

        for (uint256 i = lastFinalizedBatchIndex + 1; i <= lastCommittedBatchIndex; i++) {
            if (i != batchIndex) {
                batchDataStore[i].finalizeTimestamp += proofWindow;
            }
        }

        inChallenge = true;
    }

    /// @notice Update proofWindow.
    /// @param _newWindow New proof window.
    function updateProofWindow(uint256 _newWindow) external onlyOwner {
        require(_newWindow > 0 && _newWindow != proofWindow, "invalid new proof window");
        uint256 _oldProofWindow = proofWindow;
        proofWindow = _newWindow;
        emit UpdateProofWindow(_oldProofWindow, proofWindow);
    }

    /// @notice Update finalizationPeriodSeconds.
    /// @param _newPeriod New finalize period seconds.
    function updateFinalizePeriodSeconds(uint256 _newPeriod) external onlyOwner {
        require(_newPeriod > 0 && _newPeriod != finalizationPeriodSeconds, "invalid new finalize period");
        uint256 _oldFinalizationPeriodSeconds = finalizationPeriodSeconds;
        finalizationPeriodSeconds = _newPeriod;
        emit UpdateFinalizationPeriodSeconds(_oldFinalizationPeriodSeconds, finalizationPeriodSeconds);
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
        require(_newVerifier != address(0) && _newVerifier != verifier, "invalid new verifier");
        address _oldVerifier = verifier;
        verifier = _newVerifier;
        emit UpdateVerifier(_oldVerifier, _newVerifier);
    }

    /// @notice Update proof reward percentage
    /// @param _newProofRewardPercent Percentage awarded to prover
    function updateRewardPercentage(uint256 _newProofRewardPercent) external onlyOwner {
        require(
            _newProofRewardPercent > 0 && _newProofRewardPercent <= 100 && _newProofRewardPercent != proofRewardPercent,
            "invalid proof reward percentage"
        );
        uint256 _oldRewardPercentage = proofRewardPercent;
        proofRewardPercent = _newProofRewardPercent;
        emit UpdateProofRewardPercent(_oldRewardPercentage, _newProofRewardPercent);
    }

    /// @notice claim prove remaining
    /// @param receiver  receiver address
    function claimProveRemaining(address receiver) external onlyOwner {
        uint256 _proveRemaining = proveRemaining;
        proveRemaining = 0;
        _transfer(receiver, _proveRemaining);
        emit ProveRemainingClaimed(receiver, _proveRemaining);
    }

    /// @notice Pause the contract
    /// @param _status The pause status to update.
    function setPause(bool _status) external onlyOwner {
        if (_status) {
            _pause();
            // if challenge exist and not finished yet, return challenge deposit to challenger
            if (inChallenge) {
                batchChallengeReward[challenges[batchChallenged].challenger] += challenges[batchChallenged]
                    .challengeDeposit;
                delete challenges[batchChallenged];
                inChallenge = false;
            }
            emit Paused(_msgSender());
        } else {
            _unpause();
            emit Unpaused(_msgSender());
        }
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @dev proveState proves a batch by submitting a proof.
    function proveState(bytes calldata _batchHeader, bytes calldata _batchProof) external nonReqRevert whenNotPaused onlyActiveStaker{
        // get batch data from batch header
        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        // check batch hash
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        require(committedBatches[_batchIndex] == _batchHash, "incorrect batch hash");

        // Ensure challenge exists and is not finished
        require(batchInChallenge(_batchIndex), "batch in challenge");

        // Mark challenge as finished
        challenges[_batchIndex].finished = true;
        inChallenge = false;

        // Check for timeout
        if (challenges[_batchIndex].startTime + proofWindow <= block.timestamp) {
            // set status
            challenges[_batchIndex].challengeSuccess = true;
            _challengerWin(_batchIndex, batchDataStore[_batchIndex].signedSequencersBitmap, "Timeout");
        } else {
            _verifyProof(memPtr, _batchProof);
            // Record defender win
            _defenderWin(_batchIndex, _msgSender(), "Proof success");
        }
    }

    /// @dev finalize batch
    function finalizeBatch(bytes calldata _batchHeader) public nonReqRevert whenNotPaused {
        // get batch data from batch header
        (uint256 memPtr, bytes32 _batchHash) = _loadBatchHeader(_batchHeader);
        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        require(committedBatches[_batchIndex] == _batchHash, "incorrect batch hash");
        require(batchExist(_batchIndex), "batch not exist");
        require(!batchInChallenge(_batchIndex), "batch in challenge");
        require(!batchChallengedSuccess(_batchIndex), "batch should be revert");
        require(!batchInsideChallengeWindow(_batchIndex), "batch in challenge window");
        // verify previous state root.
        require(
            finalizedStateRoots[_batchIndex - 1] == BatchHeaderCodecV0.getPrevStateHash(memPtr),
            "incorrect previous state root"
        );
        // avoid duplicated verification
        require(finalizedStateRoots[_batchIndex] == bytes32(0), "batch already verified");
        // check and update lastFinalizedBatchIndex
        unchecked {
            require(lastFinalizedBatchIndex + 1 == _batchIndex, "incorrect batch index");
            lastFinalizedBatchIndex = _batchIndex;
        }

        // record state root and withdraw root
        withdrawalRoots[BatchHeaderCodecV0.getWithdrawRootHash(memPtr)] = true;
        finalizedStateRoots[_batchIndex] = BatchHeaderCodecV0.getPostStateHash(memPtr);

        // Pop finalized and non-skipped message from L1MessageQueue.
        _popL1Messages(
            BatchHeaderCodecV0.getTotalL1MessagePopped(memPtr),
            BatchHeaderCodecV0.getL1MessagePopped(memPtr)
        );

        delete batchDataStore[_batchIndex - 1];
        delete committedStateRoots[_batchIndex - 1];
        delete challenges[_batchIndex - 1];

        emit FinalizeBatch(
            _batchIndex,
            committedBatches[_batchIndex],
            BatchHeaderCodecV0.getPostStateHash(memPtr),
            BatchHeaderCodecV0.getWithdrawRootHash(memPtr)
        );
    }

    /// @notice Claim challenge reward
    /// @param receiver The receiver address
    function claimReward(address receiver) external {
        uint256 amount = batchChallengeReward[_msgSender()];
        require(amount != 0, "invalid batchChallengeReward");
        delete batchChallengeReward[_msgSender()];
        _transfer(receiver, amount);
        emit ChallengeRewardClaim(receiver, amount);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IRollup
    function isBatchFinalized(uint256 _batchIndex) external view override returns (bool) {
        return _batchIndex <= lastFinalizedBatchIndex;
    }

    /// @dev Public function to checks whether the batch is in challenge.
    /// @param batchIndex The index of the batch to be checked.
    function batchInChallenge(uint256 batchIndex) public view returns (bool) {
        return challenges[batchIndex].challenger != address(0) && !challenges[batchIndex].finished;
    }

    /// @dev Retrieves the success status of a batch challenge.
    /// @param batchIndex The index of the batch to check.
    function batchChallengedSuccess(uint256 batchIndex) public view returns (bool) {
        return challenges[batchIndex].challengeSuccess;
    }

    /// @dev Public function to checks whether batch exists.
    /// @param batchIndex The index of the batch to be checked.
    function batchExist(uint256 batchIndex) public view returns (bool) {
        return batchDataStore[batchIndex].originTimestamp > 0 && committedBatches[batchIndex] != bytes32(0);
    }

    /// @dev Public function to checks whether the batch is in challengeWindow.
    /// @param batchIndex The index of the batch to be checked.
    function batchInsideChallengeWindow(uint256 batchIndex) public view returns (bool) {
        return batchDataStore[batchIndex].finalizeTimestamp > block.timestamp;
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev Internal function to pop finalized l1 messages.
    /// @param totalL1MessagePopped The total number of L1 messages popped in all batches including current batch.
    /// @param l1MessagePopped The number of L1 messages popped in current batch.
    function _popL1Messages(uint256 totalL1MessagePopped, uint256 l1MessagePopped) internal {
        if (l1MessagePopped == 0) return;
        unchecked {
            uint256 startIndex = totalL1MessagePopped - l1MessagePopped;

            for (uint256 i = 0; i < l1MessagePopped; i += 256) {
                uint256 _count = 256;
                if (l1MessagePopped - i < _count) {
                    _count = l1MessagePopped - i;
                }
                IL1MessageQueue(messageQueue).popCrossDomainMessage(startIndex, _count);
                startIndex += 256;
            }
        }
    }

    /// @dev Internal function to verify the zk proof.
    function _verifyProof(uint256 memPtr, bytes calldata _batchProof) private view {
        // Check validity of proof
        require(_batchProof.length > 0, "Invalid batch proof");

        uint256 _batchIndex = BatchHeaderCodecV0.getBatchIndex(memPtr);
        bytes32 _blobVersionedHash = BatchHeaderCodecV0.getBlobVersionedHash(memPtr);

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                LAYER_2_CHAIN_ID,
                BatchHeaderCodecV0.getPrevStateHash(memPtr),
                BatchHeaderCodecV0.getPostStateHash(memPtr),
                BatchHeaderCodecV0.getWithdrawRootHash(memPtr),
                BatchHeaderCodecV0.getSequencerSetVerifyHash(memPtr),
                BatchHeaderCodecV0.getDataHash(memPtr),
                _blobVersionedHash
            )
        );

        IRollupVerifier(verifier).verifyAggregateProof(
            BatchHeaderCodecV0.getVersion(memPtr),
            _batchIndex,
            _batchProof,
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

    /// @dev todo
    function _getValidSequencerSet(
        bytes calldata sequencerSets,
        uint256 blockHeight
    ) internal pure returns (address[] memory) {
        // TODO require submitter was in valid sequencer set after BLS was implemented
        (
            ,
            address[] memory sequencerSet0,
            uint256 blockHeight1,
            address[] memory sequencerSet1,
            uint256 blockHeight2,
            address[] memory sequencerSet2
        ) = abi.decode(sequencerSets, (uint256, address[], uint256, address[], uint256, address[]));
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
    function _defenderWin(uint256 batchIndex, address prover, string memory _type) internal {
        uint256 challengeDeposit = challenges[batchIndex].challengeDeposit;
        uint256 reward = (challengeDeposit * proofRewardPercent) / 100;
        proveRemaining += challengeDeposit - reward;
        batchChallengeReward[prover] += reward;
        emit ChallengeRes(batchIndex, prover, _type);
    }

    /// @dev Internal function executed when the challenger wins.
    /// @param batchIndex           The index of the batch indicating where the challenge occurred.
    /// @param sequencersBitmap     An array containing the sequencers to be slashed.
    /// @param _type                Description of the challenge type.
    function _challengerWin(uint256 batchIndex, uint256 sequencersBitmap, string memory _type) internal {
        revertReqIndex = batchIndex;
        address challenger = challenges[batchIndex].challenger;
        uint256 reward = IL1Staking(l1StakingContract).slash(sequencersBitmap);
        batchChallengeReward[challenges[batchIndex].challenger] += (challenges[batchIndex].challengeDeposit + reward);
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

    /// @notice Extract the version number from a batch header
    /// @param batchHeader The encoded batch header bytes
    /// @return version The version of the batch header
    function _getBatchVersion(bytes calldata batchHeader) internal pure returns (uint8 version) {
        require(batchHeader.length > 0, "Empty batch header");
        version = uint8(batchHeader[0]); // Safe extraction of the first byte
    }

    /// @dev Internal function to load batch header from calldata to memory.
    /// @param _batchHeader The batch header in calldata.
    /// @return _memPtr     The start memory offset of loaded batch header.
    /// @return _batchHash  The hash of the loaded batch header.
    function _loadBatchHeader(bytes calldata _batchHeader) internal pure returns (uint256 _memPtr, bytes32 _batchHash) {
        uint8 _version = _getBatchVersion(_batchHeader);

        // load to memory
        uint256 _length;
        if (_version == 0) {
            (_memPtr, _length) = BatchHeaderCodecV0.loadAndValidate(_batchHeader);
        } else if (_version == 1) {
             (_memPtr, _length) = BatchHeaderCodecV1.loadAndValidate(_batchHeader);
        } else {
            revert("Unsupported batch version");
        }

        // compute batch hash
        // all the versions use the same way to compute batch hash
        _batchHash = BatchHeaderCodecV0.computeBatchHash(_memPtr, _length);
    }

    /// @dev Internal function to commit a batch with version 0.
    /// @param _lastBlockNumber The last block number in this batch.
    /// @param _numL1Messages The number of L1 messages in this batch
    /// @param _totalL1MessagesPoppedOverall The total number of L1 messages popped in all batches including current batch.
    /// @return _dataHash The computed data hash for this batch.
    function _commitBatch(
        uint64 _lastBlockNumber,
        uint16 _numL1Messages,
        uint256 _totalL1MessagesPoppedOverall
    ) internal view returns (bytes32 _dataHash) {
        uint256 startDataPtr;
        uint256 dataPtr;

        assembly {
            dataPtr := mload(0x40)
            startDataPtr := dataPtr
        }

        assembly {
            mstore(dataPtr, shl(192, _lastBlockNumber)) // store lastBlockNumber
            dataPtr := add(dataPtr, 8)
            mstore(dataPtr, shl(240, _numL1Messages)) // store numL1Messages
            dataPtr := add(dataPtr, 2)
            mstore(0x40, add(dataPtr, mul(_numL1Messages, 0x20))) // reserve memory for l1 message hashes
        }

        // concatenate l1 message hashes
        dataPtr = _loadL1MessageHashes(dataPtr, _numL1Messages, _totalL1MessagesPoppedOverall);

        // compute data hash and store to memory
        assembly {
            _dataHash := keccak256(startDataPtr, sub(dataPtr, startDataPtr))
        }
    }

    /// @dev Internal function to load L1 message hashes from the message queue.
    /// @param _ptr                             The memory offset to store the transaction hash.
    /// @param _numL1Messages                   The number of L1 messages to load.
    /// @param _totalL1MessagesPoppedOverall    The total number of L1 messages popped in all batches including current batch.
    /// @return uint256                         The new memory offset after loading.
    function _loadL1MessageHashes(
        uint256 _ptr,
        uint256 _numL1Messages,
        uint256 _totalL1MessagesPoppedOverall
    ) internal view returns (uint256) {
        if (_numL1Messages == 0) {
            return _ptr;
        }
        IL1MessageQueue _messageQueue = IL1MessageQueue(messageQueue);

        unchecked {
            for (uint256 i = 0; i < _numL1Messages; i++) {
                bytes32 _hash = _messageQueue.getCrossDomainMessage(_totalL1MessagesPoppedOverall);
                assembly {
                    mstore(_ptr, _hash)
                    _ptr := add(_ptr, 0x20)
                }
                _totalL1MessagesPoppedOverall += 1;
            }
        }

        return _ptr;
    }
}
