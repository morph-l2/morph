// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";

/// @dev Mirrors Rollup._computeBlobVersionedHash legacy branch (V0/V1) for isolated testing.
contract BlobVersionedHashLegacyHarness {
    bytes32 internal constant ZERO_VERSIONED_HASH =
        0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;

    function legacyBlobVersionedHash() external view returns (bytes32) {
        require(blobhash(1) == bytes32(0), "legacy batches support exactly 1 blob");
        return (blobhash(0) == bytes32(0)) ? ZERO_VERSIONED_HASH : blobhash(0);
    }
}

contract BlobVersionedHashLegacyTest is Test {
    function test_legacyBlobVersionedHash_noBlobs_returnsZeroSentinel() public {
        BlobVersionedHashLegacyHarness h = new BlobVersionedHashLegacyHarness();
        bytes32 expected =
            0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014;
        assertEq(h.legacyBlobVersionedHash(), expected);
    }
}
