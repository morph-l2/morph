// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "./Ownable.sol";
import {EvmVerifier} from "./libs/EvmVerifier.sol";

/// @title ShadowRollup
/// @notice This contract maintains data for shadow rollup.
contract ShadowRollup is Ownable {

    /// @notice The chain id of the corresponding layer 2 chain.
    uint64 public immutable layer2ChainId;

    /// @notice The address of zkvm verifier.
    address public verifier;

    struct BatchStore {
        bytes32 prevStateRoot;
        bytes32 postStateRoot;
        bytes32 withdrawalRoot;
        bytes32 dataHash;
        bytes32 blobVersionedHash;
        bytes32 sequencerSetVerifyHash;
    }

    mapping(uint256 => BatchStore) public committedBatchStores;

    // Prove result storage
    mapping(uint256 => bool) public proveStatus;

    /***************
     * Constructor *
     ***************/

    constructor(uint64 _chainId, address _verifier) {
        layer2ChainId = _chainId;
        verifier = _verifier;
    }

    /// @notice Commit a batch of transactions on layer 1.
    ///
    /// @param _batchIndex The index of batch
    /// @param _batchData The batch data
    function commitBatch(
        uint64 _batchIndex,
        BatchStore calldata _batchData
    ) external onlyOwner {
        committedBatchStores[_batchIndex] = _batchData;
    }

    // proveState proves a batch by submitting a proof.
    function proveState(uint64 _batchIndex, bytes calldata _proof) external {
        // Check validity of proof
        require(_proof.length > 0, "Invalid zk proof");

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                layer2ChainId,
                committedBatchStores[_batchIndex].prevStateRoot,
                committedBatchStores[_batchIndex].postStateRoot,
                committedBatchStores[_batchIndex].withdrawalRoot,
                committedBatchStores[_batchIndex].sequencerSetVerifyHash,
                committedBatchStores[_batchIndex].dataHash,
                committedBatchStores[_batchIndex].blobVersionedHash
            )
        );

        // Verify zk-proof
        EvmVerifier(verifier).verifyBatchProof(
            _proof,
            abi.encodePacked(_publicInputHash)
        );

        proveStatus[_batchIndex] = true;
    }

    /// @notice Update the address verifier contract.
    function updateVerifier(address _verifier) external onlyOwner {
        verifier = _verifier;
    }

    function isProveSuccess(uint256 _batchIndex) external view returns (bool) {
        require(_batchIndex > 0, "invalid batchIndex");
        return proveStatus[_batchIndex];
    }
}
