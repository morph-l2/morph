// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IL2Staking} from "../staking/IL2Staking.sol";
import {IMorphToken} from "./IMorphToken.sol";

contract MorphToken is IMorphToken, OwnableUpgradeable {
    // day seconds
    uint256 private constant DAY_SECONDS = 86400;
    // daily inflation ratio precision
    uint256 private constant PRECISION = 1e16;

    // l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;
    // distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;
    // record contract address
    address public immutable RECORD_CONTRACT;

    string private _name;
    string private _symbol;
    uint256 private _totalSupply;
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;

    // daily inflation rate
    DailyInflationRate[] private _dailyInflationRates; // TODO: precision confirm
    // mapping(day_index => inflation_amount)
    mapping(uint256 => uint256) private _inflations;
    // inflation minted days
    uint256 private _inflationMintedDays;

    /**
     * @notice Ensures that the caller message from record contract.
     */
    modifier onlyRecordContract() {
        require(msg.sender == RECORD_CONTRACT, "only record contract allowed");
        _;
    }

    /**
     * @notice constructor
     */
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
        RECORD_CONTRACT = Predeploys.RECORD;
    }

    /**
     * @dev See {IMorphToken-initialize}.
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        uint256 initialSupply_,
        uint256 dailyInflationRate_
    ) public initializer {
        __Ownable_init();

        _name = name_;
        _symbol = symbol_;
        _mint(msg.sender, initialSupply_);

        _dailyInflationRates.push(DailyInflationRate(dailyInflationRate_, 0));

        emit UpdateDailyInflationRate(dailyInflationRate_, 0);
    }

    /**
     * @dev See {IMorphToken-setRate}.
     */
    function updateRate(
        uint256 newRate,
        uint256 effectiveDaysAfterLatestUpdate
    ) public onlyOwner {
        require(
            effectiveDaysAfterLatestUpdate > 0,
            "effective days after must be greater than 0"
        );

        uint256 effectiveDay = _dailyInflationRates[
            _dailyInflationRates.length - 1
        ].effectiveDayIndex + effectiveDaysAfterLatestUpdate;
        _dailyInflationRates.push(DailyInflationRate(newRate, effectiveDay));

        emit UpdateDailyInflationRate(newRate, effectiveDay);
    }

    /**
     * @dev Returns the name of the token.
     */
    function name() public view returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token, usually a shorter version of the
     * name.
     */
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the number of decimals used to get its user representation.
     * For example, if `decimals` equals `2`, a balance of `505` tokens should
     * be displayed to a user as `5.05` (`505 / 10 ** 2`).
     *
     * Tokens usually opt for a value of 18, imitating the relationship between
     * Ether and Wei. This is the default value returned by this function, unless
     * it's overridden.
     *
     * NOTE: This information is only used for _display_ purposes: it in
     * no way affects any of the arithmetic of the contract, including
     * {IMorphToken-balanceOf} and {IMorphToken-transfer}.
     */
    function decimals() public pure returns (uint8) {
        return 18;
    }

    /**
     * @dev See {IMorphToken-totalSupply}.
     */
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev See {IMorphToken-balanceOf}.
     */
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev See {IMorphToken-inflationRatesCount}.
     */
    function inflationRatesCount() public view returns (uint256) {
        return _dailyInflationRates.length;
    }

    /**
     * @dev See {IMorphToken-dailyInflationRates}.
     */
    function dailyInflationRates(
        uint256 index
    ) public view returns (DailyInflationRate memory) {
        return _dailyInflationRates[index];
    }

    /**
     * @dev See {IMorphToken-inflation}.
     */
    function inflation(uint256 dayIndex) public view returns (uint256) {
        return _inflations[dayIndex];
    }

    /**
     * @dev See {IMorphToken-inflationMintedDays}.
     */
    function inflationMintedDays() public view returns (uint256) {
        return _inflationMintedDays;
    }

    /**
     * @dev See {IMorphToken-transfer}.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(address to, uint256 amount) public returns (bool) {
        address owner = msg.sender;
        _transfer(owner, to, amount);
        return true;
    }

    /**
     * @dev See {IMorphToken-allowance}.
     */
    function allowance(
        address owner,
        address spender
    ) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev See {IMorphToken-approve}.
     *
     * NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
     * `transferFrom`. This is semantically equivalent to an infinite approval.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function approve(address spender, uint256 amount) public returns (bool) {
        address owner = msg.sender;
        _approve(owner, spender, amount);
        return true;
    }

    /**
     * @dev See {IMorphToken-transferFrom}.
     *
     * Emits an {Approval} event indicating the updated allowance. This is not
     * required by the EIP. See the note at the beginning of {MorphToken}.
     *
     * NOTE: Does not update the allowance if the current allowance
     * is the maximum `uint256`.
     *
     * Requirements:
     *
     * - `from` and `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     * - the caller must have allowance for ``from``'s tokens of at least
     * `amount`.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public override returns (bool) {
        address spender = msg.sender;
        _spendAllowance(from, spender, amount);
        _transfer(from, to, amount);
        return true;
    }

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IMorphToken-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(
        address spender,
        uint256 addedValue
    ) public virtual returns (bool) {
        address owner = msg.sender;
        _approve(owner, spender, allowance(owner, spender) + addedValue);
        return true;
    }

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IMorphToken-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
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
    ) public virtual returns (bool) {
        address owner = msg.sender;
        uint256 currentAllowance = allowance(owner, spender);
        require(
            currentAllowance >= subtractedValue,
            "decreased allowance below zero"
        );
        unchecked {
            _approve(owner, spender, currentAllowance - subtractedValue);
        }

        return true;
    }

    /**
     * @dev Moves `amount` of tokens from `from` to `to`.
     *
     * This internal function is equivalent to {transfer}, and can be used to
     * e.g. implement automatic token fees, slashing mechanisms, etc.
     *
     * Emits a {Transfer} event.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     */
    function _transfer(address from, address to, uint256 amount) internal {
        require(from != address(0), "transfer from the zero address");
        require(to != address(0), "transfer to the zero address");

        uint256 fromBalance = _balances[from];
        require(fromBalance >= amount, "transfer amount exceeds balance");
        unchecked {
            _balances[from] = fromBalance - amount;
            // Overflow not possible: the sum of all balances is capped by totalSupply, and the sum is preserved by
            // decrementing then incrementing.
            _balances[to] += amount;
        }

        emit Transfer(from, to, amount);
    }

    /**
     * @dev mint inflations
     * @param upToDayIndex mint up to which day
     */
    function mintInflations(uint256 upToDayIndex) external onlyRecordContract {
        uint256 currentDayIndex = (block.timestamp -
            IL2Staking(L2_STAKING_CONTRACT).REWARD_START_TIME()) /
            DAY_SECONDS +
            1;
        require(
            currentDayIndex > upToDayIndex,
            "the specified time has not yet been reached"
        );
        require(
            currentDayIndex > _inflationMintedDays,
            "all inflations minted"
        );

        uint256 _inflationMintedDaysSubstitute = _inflationMintedDays;
        for (uint256 i = _inflationMintedDays; i <= upToDayIndex; i++) {
            uint256 rate = _dailyInflationRates[0].rate;
            // find inflation rate of the day
            for (uint256 j = _dailyInflationRates.length - 1; j > 0; j--) {
                if (_dailyInflationRates[j].effectiveDayIndex <= i) {
                    rate = _dailyInflationRates[j].rate;
                }
            }
            _inflations[i] = (_totalSupply * rate) / PRECISION;
            _mint(DISTRIBUTE_CONTRACT, _inflations[i]);
            _inflationMintedDaysSubstitute++;
        }
        _inflationMintedDays = _inflationMintedDaysSubstitute;
    }

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     */
    function _mint(address account, uint256 amount) internal {
        require(account != address(0), "mint to the zero address");
        _totalSupply += amount;
        unchecked {
            // Overflow not possible: balance + amount is at most totalSupply + amount, which is checked above.
            _balances[account] += amount;
        }
        emit Transfer(address(0), account, amount);
    }

    /**
     * @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
     *
     * This internal function is equivalent to `approve`, and can be used to
     * e.g. set automatic allowances for certain subsystems, etc.
     *
     * Emits an {Approval} event.
     *
     * Requirements:
     *
     * - `owner` cannot be the zero address.
     * - `spender` cannot be the zero address.
     */
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "approve from the zero address");
        require(spender != address(0), "approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Updates `owner` s allowance for `spender` based on spent `amount`.
     *
     * Does not update the allowance amount in case of infinite allowance.
     * Revert if not enough allowance is available.
     *
     * Might emit an {Approval} event.
     */
    function _spendAllowance(
        address owner,
        address spender,
        uint256 amount
    ) internal {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= amount, "insufficient allowance");
            unchecked {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }

    /**
     * @dev This empty reserved space is put in place to allow future versions to add new
     * variables without shifting down storage in the inheritance chain.
     * See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
     */
    uint256[38] private __gap;
}
