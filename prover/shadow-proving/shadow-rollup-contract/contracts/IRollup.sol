// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

interface IRollup {

    function layer2ChainId() external view returns (uint64 chainId);
}
