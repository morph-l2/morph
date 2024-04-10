// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {IERC20MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol";

/**
 * @dev Interface of the MorphToken standard as defined in the EIP.
 */
interface IMorphToken is IERC20MetadataUpgradeable {
    /**
     * @dev Emitted the owner sets the next valid exchange rate.
     */
    event SetRate(uint256 indexed rate, uint256 indexed beginTime);

    /**
     * @dev Initialization parameter, which can only be called once.
     * @param name_ Assign the _name field to use.
     * @param symbol_ Assign the _symbol field to use.
     * @param distribute_ Assign the _distribute field to use.
     * @param initialSupply_ Initialize amount.
     * @param rate_ Annual increment parameter.
     * @param beginTime_ Mint begin time.
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        address distribute_,
        uint256 initialSupply_,
        uint256 rate_,
        uint256 beginTime_
    ) external;

    /**
     * @dev query reward
     * @param beginTimeOfOneDay begin time of each day
     */
    function reward(uint256 beginTimeOfOneDay) external returns (uint256);

    /**
     * @dev set rate
     * 1.0001596535874529 is the 365 square root of 1.06.
     * 1.0019008376772350 is the 365 square root of 2.
     *
     * @param rate The value of rate must be a decimal place multiplied by 1e16.
     * eg: When 6% is issued annually, the value of the rate is 1596535874529.
     * eg: When 100% is issued annually, the value of the rate is 19008376772350.
     * the exchange rate must be greater than or equal to zero and less than or equal to 19008376772350.
     * That is, there will be no additional issuance or the maximum annual increase will be doubled.
     *
     * @param beginTime The effective time of the exchange rate is incremental,
     * and the effective time interval of every two exchange rates must be more than two weeks.
     *
     */
    function setRate(uint256 rate, uint256 beginTime) external;

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IMorphToken-approve}.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(
        address spender,
        uint256 addedValue
    ) external returns (bool);

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IMorphToken-approve}.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `spender` must have allowance for the caller of at least
     * `subtractedValue`.
     */
    function decreaseAllowance(
        address spender,
        uint256 subtractedValue
    ) external returns (bool);

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     * Only mint once a day,
     * but can unify the previous days of mint after several days
     *
     * Requirements:
     *
     * - `account` Used if passed a non-zero address, otherwise the caller address.
     */
    function mint() external returns (uint256, uint256);
}
