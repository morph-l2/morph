pragma solidity =0.8.16;

import "../L1/Rollup.sol";

contract TestGetWithdrawRoot {
    Rollup public rollup;

    constructor(Rollup _rollup) {
        rollup = _rollup;
    }

    function getPreAndCur(
        bytes32 _withdrawalRoot
    ) public view returns (uint256) {
        uint256 withdrawBatchIndex = rollup.withdrawalRoots(_withdrawalRoot);
        
        require(withdrawBatchIndex>0, "do not submit withdrawalRoot");
        return withdrawBatchIndex;
    }
}
