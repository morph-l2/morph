// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {CommonTest} from "./CommonTest.t.sol";
import {Types} from "../../libraries/common/Types.sol";

contract FFITest is CommonTest {
    function test_get_succeeds() external {
        uint64 num = ffi.getTest();
        assertEq(num, 12345);
    }

    function test_getProveWithdrawalTransactionInputs_succeeds() external {
        bytes32 wdHash = bytes32(uint256(2));
        (bytes32 wdHashRes, bytes32[32] memory wdProof, bytes32 wdRoot) = ffi.getProveWithdrawalTransactionInputs(
            wdHash
        );
        assertEq(wdHash, wdHashRes);
        assertTrue(verifyMerkleProof(wdHashRes, wdProof, 0, wdRoot));
    }

    function test_generateStakerInfo_succeeds() external {
        for (uint256 i = 0; i < 10; i++) {
            address user = address(uint160(10 + i));
            Types.StakerInfo memory stakerInfo = ffi.generateStakerInfo(user);
            assertEq(stakerInfo.addr, user);
        }
    }
}
