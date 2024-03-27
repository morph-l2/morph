// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {IERC20MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol";

/**
 * @dev Interface of the MorphToken standard as defined in the EIP.
 */
interface IMorphToken is IERC20MetadataUpgradeable {

    /**
     * @dev Initialization parameter, which can only be called once.
     * @param name_ Assign the _name field to use.
     * @param symbol_ Assign the _symbol field to use.
     * @param distribute_ Assign the _distribute field to use.
     * @param initialSupply_ Initialize amount.
     */
    function initialize(string memory name_, string memory symbol_, address distribute_, uint256 rate_, uint256 initialSupply_, uint256 beginTime_) external;

    /**
     * @dev Returns the base value of the data for the current year.
     */
    function additionalBase() external view returns (uint256);

    /**
     * @dev Returns _rate field.
     */
    function rate() external view returns (uint256);

    /**
     * @dev set _postRate field.
     */
    function setPostRate(uint256 rate_) external;

    /**
     * @dev Returns Additional issue start time.
     */
    function additionalBeginTime() external view returns (uint256);

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
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool);

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
    function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool);

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     * Only mint once a day,
     * but can unify the previous days of mint after several days
     *
     * Requirements:
     *
     * - `account` Used if passed a non-zero address, otherwise the caller address.
     */
    function mint(address account) external;
}
