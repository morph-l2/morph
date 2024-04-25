// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {IERC20MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol";

/**
 /// @dev Interface of the MorphToken standard as defined in the EIP.
 */
interface IMorphToken is IERC20MetadataUpgradeable {
    /***********
     * Structs *
     ***********/

    /// @notice DailyInflationRate representing a daily inflation rate.
    ///
    /// @custom:field rate               daily inflation ratio, precision is 1e16
    /// @custom:field effectiveDayIndex  effective day index
    struct DailyInflationRate {
        uint256 rate;
        uint256 effectiveDayIndex;
    }

    /**********
     * Events *
     **********/

    /// @notice Emitted the owner sets the next valid exchange rate.
    /// @param rate                 new rate
    /// @param effectiveDayIndex    effective day index
    event UpdateDailyInflationRate(
        uint256 indexed rate,
        uint256 indexed effectiveDayIndex
    );

    /// @notice Inflation minted
    /// @param dayIndex     minted day index
    /// @param amount       inflation amount
    event InflationMinted(uint256 indexed dayIndex, uint256 amount);

    /*************************
     * Public View Functions *
     *************************/

    /// @dev inflationRatesCount returns the total rate for all Settings.
    function inflationRatesCount() external view returns (uint256);

    /// @dev query daily inflation rates.
    /// @param index in array
    function dailyInflationRates(
        uint256 index
    ) external view returns (DailyInflationRate memory);

    /// @dev inflation returns amount of daily issues.
    /// @param dayIndex day index from start inflation.
    /// greater than or equal to zero,
    /// and less than or equal to the return value of inflationMintedDays.
    function inflation(uint256 dayIndex) external view returns (uint256);

    /// @dev inflationMintedDays
    /// returns the maximum number of days that have been mint recently.
    function inflationMintedDays() external view returns (uint256);

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @dev update rate
    /// 1.0001596535874529 is the 365 square root of 1.06.
    /// 1.0019008376772350 is the 365 square root of 2.
    ///
    /// @param newRate The value of rate must be a decimal place multiplied by 1e16.
    /// eg: When 6% is issued annually, the value of the rate is 1596535874529.
    /// eg: When 100% is issued annually, the value of the rate is 19008376772350.
    /// the exchange rate must be greater than or equal to zero and less than or equal to 19008376772350.
    /// That is, there will be no additional issuance or the maximum annual increase will be doubled.
    ///
    /// @param effectiveDayIndex effective day index
    function updateRate(uint256 newRate, uint256 effectiveDayIndex) external;

    /// @dev Atomically increases the allowance granted to `spender` by the caller.
    ///
    /// This is an alternative to {approve} that can be used as a mitigation for
    /// problems described in {IMorphToken-approve}.
    ///
    /// Requirements:
    ///
    /// - `spender` cannot be the zero address.
    function increaseAllowance(
        address spender,
        uint256 addedValue
    ) external returns (bool);

    /// @dev Atomically decreases the allowance granted to `spender` by the caller.
    ///
    /// This is an alternative to {approve} that can be used as a mitigation for
    /// problems described in {IMorphToken-approve}.
    ///
    /// Requirements:
    ///
    /// - `spender` cannot be the zero address.
    /// - `spender` must have allowance for the caller of at least `subtractedValue`.
    function decreaseAllowance(
        address spender,
        uint256 subtractedValue
    ) external returns (bool);

    /// @dev mint inflations
    /// @param upToDayIndex mint up to which day
    function mintInflations(uint256 upToDayIndex) external;
}
