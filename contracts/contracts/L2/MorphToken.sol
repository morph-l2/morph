// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./IMorphToken.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract MorphToken is Initializable, IMorphToken {
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;
    // The basis value of the annual issue.
    uint256 private _additionalBase;
    // Annual increase ratio
    // It's an integer, and it's treated with _rate/100.
    uint256 private _rate;
    // The next exchange rate used.
    uint256 private _postRate;
    // Additional issue start time.
    uint256 private _additionalBeginTime;
    // Additional issuance time of the previous year.
    uint256 private _preYearAdditionalTime;
    // Additional issuance time of the previous day.
    uint256 private _preDayAdditionalTime;
    string private _name;
    string private _symbol;
    address public _distribute;


    /**
     * @notice Ensures that the caller message from distribute contract.
     */
    modifier onlyDistribute() {
        require(msg.sender == _distribute, "only distribute contract can call");
        _;
    }

    /**
     * @dev Sets the values for {name} and {symbol}.
     *
     * All two of these values are immutable: they can only be set once during
     * construction.
     */
    function initialize(string memory name_, string memory symbol_, address distribute_, uint256 rate_, uint256 initialSupply_, uint256 beginTime_) public initializer {
        require(distribute_ != address(0), "invalid distribute contract address");
        _name = name_;
        _symbol = symbol_;
        _distribute = distribute_;
        _rate = rate_;
        _postRate = rate_;
        _additionalBeginTime = beginTime_;
        _preDayAdditionalTime = beginTime_;
        _preYearAdditionalTime = beginTime_;
        _mint(msg.sender, initialSupply_);
        _additionalBase = initialSupply_;
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
     * @dev See {IMorphToken-additionalBase}.
     */
    function additionalBase() public view returns (uint256) {
        return _additionalBase;
    }

    /**
     * @dev See {IMorphToken-rate}.
     */
    function rate() public view returns (uint256) {
        return _rate;
    }

    /**
     * @dev See {IMorphToken-setPostRate}.
     */
    function setPostRate(uint256 rate_) public {
        _postRate = rate_;
    }

    /**
     * @dev See {IMorphToken-additionalBeginTime}.
     */
    function additionalBeginTime() public view returns (uint256) {
        return _additionalBeginTime;
    }

    /**
     * @dev See {IMorphToken-balanceOf}.
     */
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
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
    function allowance(address owner, address spender) public view returns (uint256) {
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
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
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
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
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
    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        address owner = msg.sender;
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= subtractedValue, "decreased allowance below zero");
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

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     * Only mint once a day,
     * but can unify the previous days of mint after several days
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` Used if passed a non-zero address, otherwise the caller address.
     */
    function mint(address account) external onlyDistribute {
        require(block.timestamp > _additionalBeginTime, "mint feature is not yet available");

        if (account == address(0)) {
            account = msg.sender;
        }

        // Current time Indicates the number of years since the last issue
        // 31536000 = 365 * 24 * 60 * 60 (one year)
        uint256 intervalYears = (block.timestamp - _preYearAdditionalTime) / 31536000;
        if (intervalYears == 0) {
            // Current time Indicates the number of days since the last issue
            // 86400 = 24 * 60 * 60 (one day)
            uint256 interval = (block.timestamp - _preDayAdditionalTime) / 86400;
            if (interval == 0) {
                revert("only mint once a day");
            }else {
                // The daily increment value is calculated based on the basic increment value
                // and increment rate
                // eg: (_rate / 100) * _additionalBase / 365
                uint256 increment = _additionalBase * _rate / 100 / 365;
                _mint(account, interval * increment);
                _preDayAdditionalTime = _preDayAdditionalTime + interval * 86400;
            }
        } else {
            // Cross-year cycle
            for (uint256 i = 0; i < intervalYears; i++) {
                uint256 incrementOfOneDay = _additionalBase * _rate / 100 / 365;
                uint256 intervalDays = (_preYearAdditionalTime + 31536000 - _preDayAdditionalTime) / 86400;
                _mint(account, intervalDays * incrementOfOneDay);
                // Ignore decimal places and round them
                _additionalBase += _additionalBase * _rate / 100 / 365 * 365;
                _preDayAdditionalTime = _preDayAdditionalTime + intervalDays * 86400;
                _preYearAdditionalTime += 31536000;
            }
            uint256 intervals = (block.timestamp - _preDayAdditionalTime) / 86400;
            uint256 incrementOfOneDays = _additionalBase * _rate / 100 / 365;
            _mint(account, intervals * incrementOfOneDays);
            _preDayAdditionalTime = _preDayAdditionalTime + intervals * 86400;
        }

        if (_rate != _postRate) {
            _rate = _postRate;
            _additionalBase = _totalSupply;
            _preDayAdditionalTime = block.timestamp;
            _preYearAdditionalTime = block.timestamp;
        }
    }

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
    function _spendAllowance(address owner, address spender, uint256 amount) internal {
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
