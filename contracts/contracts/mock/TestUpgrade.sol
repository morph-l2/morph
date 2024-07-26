// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;
interface TestUpgrade {
    function va() external view returns (uint256);
    function vb() external view returns (uint256);
    function vc() external view returns (uint256);
    function vd() external view returns (uint256);
    function version() external view returns (uint256);
    function setVersion(uint256) external;
}

contract TestUpgradeV1 is TestUpgrade {
    uint256 public immutable va;
    uint256 public immutable vb;
    uint256 public immutable vc;

    uint256 public version = 1;

    constructor() {
        va = 1;
        vb = 2;
        vc = 3;
    }

    function vd() public pure returns (uint256) {
        return 0;
    }

    function setVersion(uint256 _version) public {
        version = _version;
    }
}

contract TestUpgradeV2 is TestUpgrade {
    // change vb and va slot order
    uint256 public immutable vd;
    uint256 public immutable vb;
    uint256 public immutable va;
    uint256 public immutable vc;

    uint256 public otherVersion = 2;
    uint256 public version = 2;

    constructor() {
        vb = 2;
        va = 3;
        vc = 4;
    }

    function setOtherVersion(uint256 _version) public {
        version = _version;
    }
    function setVersion(uint256 _version) public {
        version = _version;
    }
}
