// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

contract Test {
    uint256 public immutable AMOUNT;
    uint256 public test;
    constructor(uint256 _amount,uint256 _test){
        AMOUNT = _amount;
        initialize(_test);
    }

    function initialize(uint256 _test) public{
        test=_test;
    }

    function extractBlockNumber(bytes memory bkWitness) public pure returns (uint64) {
        require(bkWitness.length >= 32, "Data length too short");

        uint64 value;
        assembly {
            value := mload(add(bkWitness, 32))
        }
        return value;
    }
}