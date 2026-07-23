// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {VmSafe} from "forge-std/Vm.sol";

/// @notice Key-value address registry.  Writes a flat JSON object to disk.
///
/// Builds JSON manually via abi.encodePacked because serializeAddress
/// chaining is broken in forge 1.5.1 (each call overwrites the previous).
library Storage {
    function get(VmSafe vm, string memory path, string memory name) internal returns (address) {
        string memory json = vm.readFile(path);
        return vm.parseJsonAddress(json, string(abi.encodePacked("$.", name)));
    }

    function getOrFail(VmSafe vm, string memory path, string memory name) internal returns (address addr) {
        addr = get(vm, path, name);
        require(addr != address(0), string(abi.encodePacked("Storage: missing ", name)));
    }

    function setMany(VmSafe vm, string memory path, string[] memory names, address[] memory addrs) internal {
        require(names.length == addrs.length, "Storage: length mismatch");
        vm.writeFile(path, _buildFlatJson(vm, names, addrs));
    }

    function mergeMany(VmSafe vm, string memory path, string[] memory names, address[] memory addrs) internal {
        setMany(vm, path, names, addrs);
    }

    /// @dev Build "{\"k1\":\"v1\",\"k2\":\"v2\"}" using abi.encodePacked.
    function _buildFlatJson(VmSafe v, string[] memory ks, address[] memory vs) private returns (string memory) {
        bytes memory b = bytes("{");
        for (uint256 i = 0; i < ks.length; i++) {
            if (i > 0) {
                b = abi.encodePacked(b, bytes(","));
            }
            b = abi.encodePacked(
                b,
                bytes('"'), bytes(ks[i]), bytes('":"'),
                bytes(v.toString(vs[i])),
                bytes('"')
            );
        }
        b = abi.encodePacked(b, bytes("}"));
        return string(b);
    }
}
