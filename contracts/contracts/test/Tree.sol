// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {CommonTest} from "./base/CommonTest.t.sol";

contract TestTree is CommonTest {
    function setUp() public virtual override {
        super.setUp();
        for (uint256 i = 0; i < 1025; i++) {
            _appendMessageHash(bytes32(i));
        }
    }

    function test_tree_succeeds() public {
        for (uint64 i = 1; i < 1025; i = i * i + i) {
            (bytes32 leafHash, bytes32[32] memory wdProof, ) = ffi.getProveWithdrawalCheckProof(i);
            bytes32 root = getTreeRoot();
            verifyMerkleProof(leafHash, wdProof, i, root);
        }
    }
}
