// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {IL2Staking} from "../staking/IL2Staking.sol";
import {IMorphToken} from "./IMorphToken.sol";

contract MorphToken is IMorphToken, OwnableUpgradeable {
    /*************
     * Constants *
     *************/

    /// @notice daily inflation ratio precision
    uint256 private constant PRECISION = 1e16;

    /// @notice l2 staking contract address
    address public immutable L2_STAKING_CONTRACT;

    /// @notice distribute contract address
    address public immutable DISTRIBUTE_CONTRACT;

    /// @notice record contract address
    address public immutable RECORD_CONTRACT;

    /*************
     * Variables *
     *************/

    /// @notice name
    string private _name;

    /// @notice symbol
    string private _symbol;

    /// @notice total supply
    uint256 private _totalSupply;

    /// @notice balances
    mapping(address owner => uint256 amount) private _balances;

    /// @notice allowances
    mapping(address owner => mapping(address spender => uint256 amount)) private _allowances;

    /// @notice per epoch inflation rate
    EpochInflationRate[] private _epochInflationRates;

    /// @notice inflations
    mapping(uint256 epochIndex => uint256 inflationAmount) private _inflations;

    /// @notice inflation minted epochs count
    uint256 private _inflationMintedEpochs;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice Ensures that the caller message from record contract.
    modifier onlyRecordContract() {
        require(_msgSender() == RECORD_CONTRACT, "only record contract allowed");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @notice constructor
    constructor() {
        L2_STAKING_CONTRACT = Predeploys.L2_STAKING;
        DISTRIBUTE_CONTRACT = Predeploys.DISTRIBUTE;
        RECORD_CONTRACT = Predeploys.RECORD;
    }

    /**************
     * Initialize *
     **************/

    /// @dev See {IMorphToken-initialize}.
    function initialize(
        string memory name_,
        string memory symbol_,
        address _owner,
        uint256 initialSupply_,
        uint256 dailyInflationRate_
    ) public initializer {
        _name = name_;
        _symbol = symbol_;
        _mint(_owner, initialSupply_);
        _epochInflationRates.push(EpochInflationRate(dailyInflationRate_, 0));
        _transferOwnership(_owner);

        emit UpdateEpochInflationRate(dailyInflationRate_, 0);
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @dev See {IMorphToken-setRate}.
    function updateRate(uint256 newRate, uint256 effectiveEpochIndex) public onlyOwner {
        require(
            _epochInflationRates[_epochInflationRates.length - 1].rate != newRate,
            "new rate is the same as the latest rate"
        );
        require(
            effectiveEpochIndex > _epochInflationRates[_epochInflationRates.length - 1].effectiveEpochIndex,
            "effective epochs after must be greater than before"
        );

        _epochInflationRates.push(EpochInflationRate(newRate, effectiveEpochIndex));

        emit UpdateEpochInflationRate(newRate, effectiveEpochIndex);
    }

    /// @dev mint inflations
    /// @param upToEpochIndex mint up to which epoch
    function mintInflations(uint256 upToEpochIndex) external onlyRecordContract {
        // inflations can only be minted for epochs that have ended.
        require(
            IL2Staking(L2_STAKING_CONTRACT).currentEpoch() > upToEpochIndex,
            "the specified time has not yet been reached"
        );
        require(upToEpochIndex >= _inflationMintedEpochs, "all inflations minted");

        // the index of next epoch to mint is equal to the count of minted epochs
        for (uint256 i = _inflationMintedEpochs; i <= upToEpochIndex; i++) {
            uint256 rate = _epochInflationRates[0].rate;
            // find inflation rate of the epoch
            for (uint256 j = _epochInflationRates.length - 1; j > 0; j--) {
                if (_epochInflationRates[j].effectiveEpochIndex <= i) {
                    rate = _epochInflationRates[j].rate;
                    break;
                }
            }
            uint256 increment = (_totalSupply * rate) / PRECISION;
            _inflations[i] = increment;
            _mint(DISTRIBUTE_CONTRACT, increment);
            emit InflationMinted(i, increment);
        }

        _inflationMintedEpochs = upToEpochIndex + 1;
    }

    /// @dev Destroys `amount` tokens from `account`, reducing the total supply.
    /// @param amount amount to destroy
    function burn(uint256 amount) external onlyOwner {
        require(amount > 0, "amount to burn is zero");
        _burn(_msgSender(), amount);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @dev See {IMorphToken-approve}.
    ///
    /// NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
    /// `transferFrom`. This is semantically equivalent to an infinite approval.
    ///
    /// Requirements:
    ///
    /// - `spender` cannot be the zero address.
    function approve(address spender, uint256 amount) public returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, amount);
        return true;
    }

    /// @dev See {IMorphToken-transferFrom}.
    ///
    /// Emits an {Approval} event indicating the updated allowance. This is not
    /// required by the EIP. See the note at the beginning of {MorphToken}.
    ///
    /// NOTE: Does not update the allowance if the current allowance
    /// is the maximum `uint256`.
    ///
    /// Requirements:
    ///
    /// - `from` and `to` cannot be the zero address.
    /// - `from` must have a balance of at least `amount`.
    /// - the caller must have allowance for ``from``'s tokens of at least `amount`.
    function transferFrom(address from, address to, uint256 amount) public override returns (bool) {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transfer(from, to, amount);
        return true;
    }

    /// @dev Atomically increases the allowance granted to `spender` by the caller.
    ///
    /// This is an alternative to {approve} that can be used as a mitigation for
    /// problems described in {IMorphToken-approve}.
    ///
    /// Emits an {Approval} event indicating the updated allowance.
    ///
    /// Requirements:
    ///
    /// - `spender` cannot be the zero address.
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, allowance(owner, spender) + addedValue);
        return true;
    }

    /// @dev Atomically decreases the allowance granted to `spender` by the caller.
    ///
    /// This is an alternative to {approve} that can be used as a mitigation for
    /// problems described in {IMorphToken-approve}.
    ///
    /// Emits an {Approval} event indicating the updated allowance.
    ///
    /// Requirements:
    ///
    /// - `spender` cannot be the zero address.
    /// - `spender` must have allowance for the caller of at least `subtractedValue`.
    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        address owner = _msgSender();
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= subtractedValue, "decreased allowance below zero");
        unchecked {
            _approve(owner, spender, currentAllowance - subtractedValue);
        }
        return true;
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @dev Returns the name of the token.
    function name() public view returns (string memory) {
        return _name;
    }

    /// @dev Returns the symbol of the token, usually a shorter version of the name.
    function symbol() public view returns (string memory) {
        return _symbol;
    }

    /// @dev Returns the number of decimals used to get its user representation.
    /// For example, if `decimals` equals `2`, a balance of `505` tokens should
    /// be displayed to a user as `5.05` (`505 / 10 ** 2`).
    ///
    /// Tokens usually opt for a value of 18, imitating the relationship between
    /// Ether and Wei. This is the default value returned by this function, unless
    /// it's overridden.
    ///
    /// NOTE: This information is only used for _display_ purposes: it in
    /// no way affects any of the arithmetic of the contract, including
    /// {IMorphToken-balanceOf} and {IMorphToken-transfer}.
    function decimals() public pure returns (uint8) {
        return 18;
    }

    /// @dev See {IMorphToken-totalSupply}.
    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    /// @dev See {IMorphToken-balanceOf}.
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    /// @dev See {IMorphToken-inflationRatesCount}.
    function inflationRatesCount() public view returns (uint256) {
        return _epochInflationRates.length;
    }

    /// @dev See {IMorphToken-epochInflationRates}.
    function epochInflationRates(uint256 index) public view returns (EpochInflationRate memory) {
        return _epochInflationRates[index];
    }

    /// @dev See {IMorphToken-inflation}.
    function inflation(uint256 epochIndex) public view returns (uint256) {
        return _inflations[epochIndex];
    }

    /// @dev See {IMorphToken-inflationMintedEpochs}.
    function inflationMintedEpochs() public view returns (uint256) {
        return _inflationMintedEpochs;
    }

    /// @dev See {IMorphToken-transfer}.
    ///
    /// Requirements:
    ///
    /// - `to` cannot be the zero address.
    /// - the caller must have a balance of at least `amount`.
    function transfer(address to, uint256 amount) public returns (bool) {
        address owner = _msgSender();
        _transfer(owner, to, amount);
        return true;
    }

    /// @dev See {IMorphToken-allowance}.
    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @dev Moves `amount` of tokens from `from` to `to`.
    ///
    /// This internal function is equivalent to {transfer}, and can be used to
    /// e.g. implement automatic token fees, slashing mechanisms, etc.
    ///
    /// Emits a {Transfer} event.
    ///
    /// Requirements:
    ///
    /// - `from` cannot be the zero address.
    /// - `to` cannot be the zero address.
    /// - `from` must have a balance of at least `amount`.
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

    /// @dev Creates `amount` tokens and assigns them to `account`, increasing
    /// the total supply.
    ///
    /// Emits a {Transfer} event with `from` set to the zero address.
    ///
    /// Requirements:
    ///
    /// - `account` cannot be the zero address.
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
     * @dev Destroys `amount` tokens from `account`, reducing the total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function _burn(address account, uint256 amount) internal {
        require(account != address(0), "ERC20: burn from the zero address");

        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
        unchecked {
            _balances[account] = accountBalance - amount;
            // Overflow not possible: amount <= accountBalance <= totalSupply.
            _totalSupply -= amount;
        }

        emit Transfer(account, address(0), amount);
    }

    /// @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
    ///
    /// This internal function is equivalent to `approve`, and can be used to
    /// e.g. set automatic allowances for certain subsystems, etc.
    ///
    /// Emits an {Approval} event.
    ///
    /// Requirements:
    ///
    /// - `owner` cannot be the zero address.
    /// - `spender` cannot be the zero address.
    function _approve(address owner, address spender, uint256 amount) internal {
        require(owner != address(0), "approve from the zero address");
        require(spender != address(0), "approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /// @dev Updates `owner` s allowance for `spender` based on spent `amount`.
    ///
    /// Does not update the allowance amount in case of infinite allowance.
    /// Revert if not enough allowance is available.
    ///
    /// Might emit an {Approval} event.
    function _spendAllowance(address owner, address spender, uint256 amount) internal {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= amount, "insufficient allowance");
            unchecked {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }

    /// @dev This empty reserved space is put in place to allow future versions to add new
    /// variables without shifting down storage in the inheritance chain.
    /// See https://docs.openzeppelin.com/contracts/4.x/upgradeable#storage_gaps
    uint256[38] private __gap;
}
