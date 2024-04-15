// contracts/GLDToken.sol
// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;
import "@openzeppelin/contracts/access/Ownable.sol";
import {IRollup} from "./IRollup.sol";
import {IZkEvmVerifier} from "./libs/IZkEvmVerifier.sol";

/// @title ShadowRollup
/// @notice This contract maintains data for shadow rollup.
contract ShadowRollup is Ownable {
    /// Scalar field modulus of BLS12-381
    uint256 constant BLS_MODULUS =
        52435875175126190479447740508185965837690552500527637822603658699938581184513;

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
    }

    mapping(uint256 => BatchStore) public committedBatchStores;

    struct BatchChallenge {
        uint64 batchIndex;
        address challenger;
        uint256 challengeDeposit;
        uint256 startTime;
        bool finished;
    }

    /**
     * @notice Store Challenge Information.(batchIndex => BatchChallenge)
     */
    mapping(uint256 => BatchChallenge) public challenges;

    /// @notice Emitted when the state of Chanllenge is updated.
    /// @param batchIndex The index of the batch.
    /// @param challenger The address of challenger.
    /// @param challengeDeposit The deposit of challenger.
    event ChallengeState(
        uint64 indexed batchIndex,
        address challenger,
        uint256 challengeDeposit
    );

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
    function commitBatch(
        uint64 _batchIndex,
        BatchStore calldata _batchData
    ) external {
        committedBatchStores[_batchIndex] = _batchData;
    }

    // challengeState challenges a batch by submitting a deposit.
    function challengeState(uint64 batchIndex) external payable onlyOwner {
        challenges[batchIndex] = BatchChallenge(
            batchIndex,
            _msgSender(),
            msg.value,
            block.timestamp,
            false
        );
        emit ChallengeState(batchIndex, _msgSender(), msg.value);
    }

    // proveState proves a batch by submitting a proof.
    // _kzgData: [y(32) | commitment(48) | proof(48)]
    function proveState(
        uint64 _batchIndex,
        bytes calldata _aggrProof,
        bytes calldata _kzgData
    ) external {
        // Check validity of proof
        require(_aggrProof.length > 0, "Invalid proof");

        // Check validity of KZG data
        require(_kzgData.length == 160, "Invalid KZG data");

        // Extract commitment
        // bytes memory _commitment = _kzgData[32:80];

        // Compute xBytes
        // bytes memory _xBytes = computeXBytes(_batchIndex, _commitment);

        // // Create input for verification
        // bytes memory _input = abi.encodePacked(
        //     committedBatchStores[_batchIndex].blobVersionedHash,
        //     _xBytes,
        //     _kzgData
        // );

        // Verify 4844-proof
        // (bool success, bytes memory data) = address(0x0A).staticcall(_input);
        // require(success, "failed to call point evaluation precompile");
        // (, uint256 result) = abi.decode(data, (uint256, uint256));
        // require(result == BLS_MODULUS, "UNKNOWN_BLS_MODULUS");

        bytes32 _publicInputHash = keccak256(
            abi.encodePacked(
                layer2ChainId,
                committedBatchStores[_batchIndex].prevStateRoot,
                committedBatchStores[_batchIndex].postStateRoot,
                committedBatchStores[_batchIndex].withdrawalRoot,
                committedBatchStores[_batchIndex].dataHash,
                _kzgData[0:64],
                committedBatchStores[_batchIndex].blobVersionedHash
            )
        );

        // Verify zk-proof
        IZkEvmVerifier(zkevm_verifier).verify(_aggrProof, _publicInputHash);

        // Record defender win
        challenges[_batchIndex].finished = true;
    }

    function computeXBytes(
        uint64 _batchIndex,
        bytes memory commitment
    ) private view returns (bytes memory) {
        bytes memory xBytes = abi.encode(
            keccak256(
                abi.encodePacked(
                    commitment,
                    committedBatchStores[_batchIndex].dataHash
                )
            )
        );
        xBytes[0] = 0x0; // make sure x < BLS_MODULUS
        return xBytes;
    }

    function computePublicInputHash(
        uint64 _batchIndex,
        bytes memory _xBytes,
        bytes memory _yBytes
    ) private view returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    layer2ChainId,
                    committedBatchStores[_batchIndex].prevStateRoot,
                    committedBatchStores[_batchIndex].postStateRoot,
                    committedBatchStores[_batchIndex].withdrawalRoot,
                    committedBatchStores[_batchIndex].dataHash,
                    splitUint256(_xBytes),
                    splitUint256(_yBytes)
                )
            );
    }

    function splitUint256(
        bytes memory _combined
    ) public pure returns (bytes memory) {
        require(_combined.length == 32, "Input length must be 32 bytes");
        uint256 combinedUint;
        assembly {
            combinedUint := mload(add(_combined, 0x20))
        }

        uint256 part1;
        uint256 part2;
        uint256 part3;

        // Extract the three parts
        part1 = (combinedUint & ((1 << 88) - 1)); // Mask the lowest 88 bits and reverse bytes
        part2 = ((combinedUint >> 88) & ((1 << 88) - 1)); // Shift right by 88 bits, mask the next 88 bits, and reverse bytes
        part3 = ((combinedUint >> 176) & ((1 << 87) - 1)); // Shift right by 176 bits, mask the next 87 bits, and reverse bytes

        bytes memory result = new bytes(96);
        assembly {
            // Store the parts in the result bytes
            mstore(add(result, 0x20), part1)
            mstore(add(result, 0x40), part2)
            mstore(add(result, 0x60), part3)
        }

        return result;
    }

    /// @notice Update the address verifier contract.
    function updateVerifier(address _verifier) external onlyOwner {
        zkevm_verifier = _verifier;
    }

    function batchInChallenge(uint256 batchIndex) public view returns (bool) {
        return
            challenges[batchIndex].challenger != address(0) &&
            !challenges[batchIndex].finished;
    }

    function isBatchFinalized(
        uint256 _batchIndex
    ) external pure returns (bool) {
        require(_batchIndex > 0, "invalid batchIndex");
        return false;
    }
}
