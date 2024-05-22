// SPDX-License-Identifier: MIT

pragma solidity ^0.8.16;

interface IRollupVerifier {
    /// @notice Verify aggregate zk proof.
    /// @param _version The version of the verifier.
    /// @param _batchIndex The batch index to verify.
    /// @param _aggrProof The aggregated proof.
    /// @param _publicInputHash The public input hash.
    function verifyAggregateProof(
        uint256 _version,
        uint256 _batchIndex,
        bytes calldata _aggrProof,
        bytes32 _publicInputHash
    ) external view;
}
