// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

contract Verify {
    function verifyMerkleProof(
        bytes32 leafHash,
        bytes32[32] calldata smtProof,
        uint256 index,
        bytes32 root
    ) public pure returns (bool) {
        bytes32 node = leafHash;

        for (uint256 height = 0; height < 32; height++) {
            if (((index >> height) & 1) == 1) node = keccak256(abi.encodePacked(smtProof[height], node));
            else node = keccak256(abi.encodePacked(node, smtProof[height]));
        }

        return node == root;
    }
}
