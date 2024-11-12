// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import {SafeERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";

contract MigrationUSDC is OwnableUpgradeable, PausableUpgradeable, ReentrancyGuardUpgradeable {
    using SafeERC20Upgradeable for IERC20Upgradeable;

    error ErrZeroAddress();

    /// @dev Thrown the token balance is zero.
    error ErrorTokenBalanceZero();

    event Migrate(address indexed user, uint256 amount);
    event Transfer(address indexed token, address indexed to, uint256 amount);
    event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient);

    /// @notice The address of old USDC address.
    address public immutable OLD_USDC;

    /// @notice The address of new USDC address.
    address public immutable NEW_USDC;

    /// @notice Wallet that will receive the tokens on L2.
    address public recipient;

    /***************
     * Constructor *
     ***************/
    /// @notice Constructor for `MigrationUSDC` implementation contract.
    ///
    /// @param _oldUSDC The address of old USDC in L2.
    /// @param _newUSDC The address of new USDC in L2.
    constructor(address _oldUSDC, address _newUSDC) {
        _disableInitializers();

        OLD_USDC = _oldUSDC;
        NEW_USDC = _newUSDC;
    }

    // initialize contract status
    function initialize(address _recipient) external initializer {
        if (_recipient == address(0)) {
            revert ErrZeroAddress();
        }
        recipient = _recipient;

        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        emit UpdateRecipient(address(0), recipient);
    }

    // Transfer all old USDC to this contract and then transfer new USDC token to msg sender.
    function migrate() external nonReentrant whenNotPaused {
        // Get old USDC balance.
        uint256 balance = IERC20Upgradeable(OLD_USDC).balanceOf(_msgSender());
        if (balance == 0) {
            revert ErrorTokenBalanceZero();
        }
        // Transfer token into this contract.
        IERC20Upgradeable(OLD_USDC).safeTransferFrom(_msgSender(), address(this), balance);
        // Transfer new USDC token to msg sender.
        IERC20Upgradeable(NEW_USDC).transfer(_msgSender(), balance);
        emit Migrate(_msgSender(), balance);
    }

    // Transfer token to other address.
    function transferToken(address _token, uint256 _amount) external onlyOwner {
        if (recipient == address(0)) {
            revert ErrZeroAddress();
        }
        uint256 balance = IERC20Upgradeable(_token).balanceOf(address(this));
        if (balance == 0) {
            revert ErrorTokenBalanceZero();
        }
        // transfer all token
        if (balance < _amount) {
            _amount = balance;
        }
        // Transfer token.
        IERC20Upgradeable(_token).transfer(recipient, _amount);
        emit Transfer(_token, recipient, _amount);
    }

    // Update the address of recipient.
    function updateRecipient(address _newRecipient) external onlyOwner {
        address _oldRecipient = recipient;
        recipient = _newRecipient;

        emit UpdateRecipient(_oldRecipient, _newRecipient);
    }

    // Change pause status
    function setPause(bool status) external onlyOwner {
        if (status) {
            _requireNotPaused();
            _pause();
        } else {
            _requirePaused();
            _unpause();
        }
    }
}
