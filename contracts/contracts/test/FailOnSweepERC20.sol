// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title FailOnSweepERC20
 * @notice Devnet-only ERC-20 whose `transfer` returns `false` for exactly one
 *         caller, so the execution layer's sweep hits the `transfer_false`
 *         classification and appends a `SweepFailed` log.
 * @dev The EL issues the sweep as `transfer(destination, balance)` with the
 *      SOURCE as caller. Returning false only for that caller lets a test fund
 *      the source normally and still force a reportable sweep failure — the one
 *      path that is otherwise hard to reach on a live devnet.
 *
 *      Deliberately minimal: no allowances, no supply accounting. Not for any
 *      network other than a local devnet.
 */
contract FailOnSweepERC20 {
    mapping(address account => uint256 balance) public balanceOf;

    /// @notice The address whose outbound transfers always fail.
    address public immutable sweepSource;

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor(address sweepSource_) {
        sweepSource = sweepSource_;
    }

    /// @notice Credits `to`. The emitted `Transfer` is itself a sweep trigger.
    function mint(address to, uint256 amount) external {
        balanceOf[to] += amount;
        emit Transfer(address(0), to, amount);
    }

    function transfer(address to, uint256 amount) external returns (bool) {
        // The EL's sweep call: report failure without moving anything.
        if (msg.sender == sweepSource) {
            return false;
        }
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }
}
