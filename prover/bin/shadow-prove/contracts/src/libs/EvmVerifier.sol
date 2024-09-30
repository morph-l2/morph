// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {SP1Verifier} from "./SP1VerifierPlonk.sol";

/// @title EvmVerifier.
/// @author MorphL2.
/// @notice This contract is used to verify the batch proof of evm.
contract EvmVerifier is SP1Verifier {
    /// @notice The verification key for the morph executor program.
    bytes32 public programVkey;

    constructor(bytes32 _programVkey) {
        programVkey = _programVkey;
    }

    /// @notice The entrypoint for verifying the proof of a morph batch.
    /// @param proof The encoded proof.
    /// @param publicValues The encoded public values.
    function verifyBatchProof(bytes calldata proof, bytes calldata publicValues) public view {
        this.verifyProof(programVkey, publicValues, proof);
    }
}
