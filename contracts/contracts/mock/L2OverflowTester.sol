// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

contract L2OverflowTester {
    bytes32 private messageHash;
    uint256 private hashCount;

    function hash(string calldata _message, uint256 count) public {
        bytes memory hashed = bytes(_message);
        for (uint256 i = 0; i < count; i++) {
            messageHash = keccak256(hashed);
            hashed = abi.encodePacked(messageHash);
        }
        hashCount = count;
    }

    function getMessageHash() public view returns (bytes32) {
        return messageHash;
    }

    function getHashCount() public view returns (uint256) {
        return hashCount;
    }
}
