// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IZkEvmVerifier} from "./IZkEvmVerifier.sol";
import {SP1Verifier} from "./SP1VerifierPlonk.sol";

// solhint-disable no-inline-assembly
contract ZkEvmVerifierV1 is IZkEvmVerifier, SP1Verifier {

    /// @notice The verification key for the morph executor program.
    bytes32 public programVkey;

    constructor(bytes32 _programVkey) {
        programVkey = _programVkey;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IZkEvmVerifier
    function verifyBatch(bytes calldata proof, bytes32 publicInputHash) external view override {
        this.verifyPlonk(proof, bytes.concat(publicInputHash));
    }

    function verifyPlonk(bytes calldata proof, bytes calldata publicInputHash) external view {
        this.verifyProof(programVkey, publicInputHash, proof);
    }
}
