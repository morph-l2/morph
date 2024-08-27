// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import {IZkEvmVerifier} from "./libs/IZkEvmVerifier.sol";

/// @title ShadowRollup
/// @notice This contract maintains data for shadow rollup.
contract ShadowRollup is Ownable {
    /// Scalar field modulus of BLS12-381
    uint256 constant BLS_MODULUS = 52435875175126190479447740508185965837690552500527637822603658699938581184513;

    /// @dev Address of the point evaluation precompile used for EIP-4844 blob verification.
    address internal constant POINT_EVALUATION_PRECOMPILE_ADDR = address(0x0A);

    /// @notice The chain id of the corresponding layer 2 chain.
    uint64 public immutable layer2ChainId;

    /// @notice The address of zkevmVerifier.
    address public zkevm_verifier;

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
        zkevm_verifier = _verifier;
    }

    /// @notice Commit a batch of transactions on layer 1.
    ///
    /// @param _batchIndex The index of batch
    /// @param _batchData The batch data
    function commitBatch(uint64 _batchIndex, BatchStore calldata _batchData) external onlyOwner {
        committedBatchStores[_batchIndex] = _batchData;
    }

    // proveState proves a batch by submitting a proof.
    // _kzgData: [y(32) | commitment(48) | proof(48)]
    function proveState(uint64 _batchIndex, bytes calldata _aggrProof, bytes calldata _kzgDataProof) external {
        // Check validity of proof
        require(_aggrProof.length > 0, "Invalid aggregation proof");

        // Check validity of KZG data
        require(_kzgDataProof.length == 160, "Invalid KZG data proof");

        // Calls the point evaluation precompile and verifies the output
        {
            (bool success, bytes memory data) = POINT_EVALUATION_PRECOMPILE_ADDR.staticcall(
                abi.encodePacked(committedBatchStores[_batchIndex].blobVersionedHash, _kzgDataProof)
            );
            // We verify that the point evaluation precompile call was successful by testing the latter 32 bytes of the
            // response is equal to BLS_MODULUS as defined in https://eips.ethereum.org/EIPS/eip-4844#point-evaluation-precompile
            require(success, "failed to call point evaluation precompile");
            (, uint256 result) = abi.decode(data, (uint256, uint256));
            require(result == BLS_MODULUS, "precompile unexpected output");
        }

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                layer2ChainId,
                committedBatchStores[_batchIndex].prevStateRoot,
                committedBatchStores[_batchIndex].postStateRoot,
                committedBatchStores[_batchIndex].withdrawalRoot,
                committedBatchStores[_batchIndex].sequencerSetVerifyHash,
                committedBatchStores[_batchIndex].dataHash,
                _kzgDataProof[0:64],
                committedBatchStores[_batchIndex].blobVersionedHash
            )
        );

        // Verify zk-proof
        IZkEvmVerifier(zkevm_verifier).verify(_aggrProof, _publicInputHash);

        proveStatus[_batchIndex] = true;
    }

    /// @notice Update the address verifier contract.
    function updateVerifier(address _verifier) external onlyOwner {
        zkevm_verifier = _verifier;
    }

    function isProveSuccess(uint256 _batchIndex) external view returns (bool) {
        require(_batchIndex > 0, "invalid batchIndex");
        return proveStatus[_batchIndex];
    }
}
