// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {DoubleEndedQueue} from "@openzeppelin/contracts/utils/structs/DoubleEndedQueue.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {IMorphToken} from "./IMorphToken.sol";

contract MorphToken is OwnableUpgradeable, IMorphToken {
    using DoubleEndedQueue for DoubleEndedQueue.Bytes32Deque;

    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    // begin time of each day => reward
    mapping(uint256 => uint256) private _rewardRecord;
    Rate private _rate;
    uint256 private _totalSupply;
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
     * @dev See {IMorphToken-initialize}.
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        address distribute_,
        uint256 initialSupply_,
        uint256 rate_,
        uint256 beginTime_
    ) public initializer {
        require(
            distribute_ != address(0),
            "invalid distribute contract address"
        );
        require(
            beginTime_ % 86400 == 0,
            "beginTime must be the start of the day"
        );
        require(
            beginTime_ >= block.timestamp,
            "beginTime is less than the current time"
        );

        __Ownable_init();

        _name = name_;
        _symbol = symbol_;
        _distribute = distribute_;
        _preDayAdditionalTime = beginTime_;
        _mint(msg.sender, initialSupply_);

        _rate.currentBeginTime = beginTime_;
        _rate.currentRate = rate_;

        emit SetRate(rate_, beginTime_);
    }

    /**
     * @dev See {IMorphToken-setRate}.
     */
    function setRate(uint256 rate, uint256 beginTime) public onlyOwner {
        require(
            beginTime > block.timestamp,
            "beginTime must be more than the current time"
        );

        // 1209600 = 14 * 24 * 60 * 60 (fortnight)
        if (_rate.pending.index.empty()) {
            require(
                beginTime >= _rate.currentBeginTime + 1209600,
                "beginTime must be two weeks after the current validity period"
            );
        } else {
            require(
                beginTime > uint256(_rate.pending.index.back()) + 1209600,
                "beginTime must be more than two weeks after the last exchange rate takes effect"
            );
        }

        // mint function is in days, when a new issue exchange rate effective time in a unit,
        // then the effective exchange rate of this unit is the previous exchange rate value,
        // and the effective exchange rate value can only be calculated until the next mint unit.
        //
        // In addition, since the start time of mint function is known,
        // the start and end time of each unit can be calculated,
        // and the specific effective time of rate can be calculated in advance.
        if ((beginTime - _preDayAdditionalTime) % 86400 > 0) {
            beginTime =
                ((beginTime - _preDayAdditionalTime) / 86400 + 1) *
                86400 +
                _preDayAdditionalTime;
        }

        _rate.pending.index.pushBack(bytes32(beginTime));
        _rate.pending.values[beginTime] = rate;

        emit SetRate(rate, beginTime);
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
     * @dev See {IMorphToken-reward}.
     */
    function reward(uint256 beginTimeOfOneDay) public view returns (uint256) {
        return _rewardRecord[beginTimeOfOneDay];
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
    function mint()
        external
        onlyDistribute
        returns (uint256 begin, uint256 end)
    {
        require(
            block.timestamp > _preDayAdditionalTime,
            "the mint function is not enabled"
        );
        require(
            (block.timestamp - _preDayAdditionalTime) / 86400 != 0,
            "only mint once a day"
        );

        begin = _preDayAdditionalTime;

        uint256 length = _rate.pending.index.length();
        for (uint256 i = 0; i < length; i++) {
            if (_rate.nextEffectiveTime == 0) {
                _rate.nextEffectiveTime = uint256(_rate.pending.index.front());
            }

            if (_rate.nextEffectiveTime <= block.timestamp) {
                // end time
                _rate.outmoded.index.pushBack(bytes32(_rate.nextEffectiveTime));
                _rate.outmoded.values[_rate.nextEffectiveTime] = _rate
                    .currentRate;

                _rate.currentBeginTime = _rate.nextEffectiveTime;
                _rate.currentRate = _rate.pending.values[
                    _rate.nextEffectiveTime
                ];

                bytes32 front = _rate.pending.index.popFront();
                if (uint256(front) != _rate.nextEffectiveTime) {
                    revert("internal error");
                }
                delete _rate.pending.values[uint256(front)];

                if (!_rate.pending.index.empty()) {
                    _rate.nextEffectiveTime = uint256(
                        _rate.pending.index.front()
                    );
                } else {
                    _rate.nextEffectiveTime = 0;
                }
            } else {
                break;
            }
        }

        uint256 outmodedLength = _rate.outmoded.index.length();
        for (uint256 i = 0; i < outmodedLength; i++) {
            uint256 validTime = uint256(_rate.outmoded.index.front());

            if (validTime > _preDayAdditionalTime) {
                // Current time Indicates the number of days since the last issue.
                uint256 day = (validTime - _preDayAdditionalTime) / 86400;
                for (uint256 k = 0; k < day; k++) {
                    uint256 outmodedReward = (_totalSupply *
                        _rate.outmoded.values[validTime]) / 1e16;
                    _rewardRecord[_preDayAdditionalTime] = outmodedReward;
                    _mint(msg.sender, outmodedReward);
                    // 86400 = 24 * 60 * 60 (one day)
                    _preDayAdditionalTime += 86400;
                }
            }

            bytes32 popFront = _rate.outmoded.index.popFront();
            if (uint256(popFront) != validTime) {
                revert("internal error");
            }
            delete _rate.outmoded.values[validTime];
        }

        // use current rate
        uint256 currentDays = (block.timestamp - _preDayAdditionalTime) / 86400;
        for (uint256 i = 0; i < currentDays; i++) {
            uint256 currentReward = (_totalSupply * _rate.currentRate) / 1e16;
            _rewardRecord[_preDayAdditionalTime] = currentReward;
            _mint(msg.sender, currentReward);
            _preDayAdditionalTime += 86400;
        }
        end = _preDayAdditionalTime;
    }

    function _mint(address account, uint256 amount) internal {
        // require(account != address(0), "mint to the zero address");

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
