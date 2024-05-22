// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

contract MockTree {
    error MerkleTreeFull();

    uint256 internal constant _TREE_DEPTH = 32;

    uint256 internal constant _MAX_DEPOSIT_COUNT = 2 ** _TREE_DEPTH - 1;

    bytes32[_TREE_DEPTH] internal _branch;

    uint256 public leafNodesCount;

    function getTreeRoot() public view returns (bytes32) {
        bytes32 node;
        uint256 size = leafNodesCount;
        bytes32 currentZeroHashHeight = 0;

        for (uint256 height = 0; height < _TREE_DEPTH; height++) {
            if (((size >> height) & 1) == 1) node = keccak256(abi.encodePacked(_branch[height], node));
            else node = keccak256(abi.encodePacked(node, currentZeroHashHeight));

            currentZeroHashHeight = keccak256(abi.encodePacked(currentZeroHashHeight, currentZeroHashHeight));
        }
        return node;
    }

    function _appendMessageHash(bytes32 leafHash) internal {
        bytes32 node = leafHash;

        // Avoid overflowing the Merkle tree (and prevent edge case in computing `_branch`)
        if (leafNodesCount >= _MAX_DEPOSIT_COUNT) {
            revert MerkleTreeFull();
        }

        // Add deposit data root to Merkle tree (update a single `_branch` node)
        uint256 size = ++leafNodesCount;
        for (uint256 height = 0; height < _TREE_DEPTH; height++) {
            if (((size >> height) & 1) == 1) {
                _branch[height] = node;
                return;
            }
            node = keccak256(abi.encodePacked(_branch[height], node));
        }
        // As the loop should always end prematurely with the `return` statement,
        // this code should be unreachable. We assert `false` just to be safe.
        assert(false);
    }

    function verifyMerkleProof(
        bytes32 leafHash,
        bytes32[_TREE_DEPTH] memory smtProof,
        uint256 index,
        bytes32 root
    ) public pure returns (bool) {
        bytes32 node = leafHash;

        for (uint256 height = 0; height < _TREE_DEPTH; height++) {
            if (((index >> height) & 1) == 1) node = keccak256(abi.encodePacked(smtProof[height], node));
            else node = keccak256(abi.encodePacked(node, smtProof[height]));
        }

        return node == root;
    }
}
