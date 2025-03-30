// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IERC20PermitUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20PermitUpgradeable.sol";
import {ERC20PermitUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PermitUpgradeable.sol";
import {SignatureCheckerUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/SignatureCheckerUpgradeable.sol";

import {MorphStandardERC20} from "../libraries/token/MorphStandardERC20.sol";

/**
 * @custom:security-contact official@morphl2.io
 */
contract L2WstETHToken is MorphStandardERC20 {
    /**********
     * Errors *
     **********/
    
    /// @dev Thrown when the deadline is expired.
    error ErrorExpiredDeadline();
    
    /// @dev Thrown when the given signature is invalid.
    error ErrorInvalidSignature();
    
    /*************
     * Constants *
     *************/

    /// @dev See {ERC20PermitUpgradeable-_PERMIT_TYPEHASH}
    bytes32 private constant _PERMIT_TYPEHASH =
        keccak256("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)");

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc IERC20PermitUpgradeable
    ///
    /// @dev The code is copied from `ERC20PermitUpgradeable` with modifications to support ERC-1271.
    function permit(
        address owner,
        address spender,
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public virtual override(ERC20PermitUpgradeable, IERC20PermitUpgradeable) {
        if (block.timestamp > deadline) {
            revert ErrorExpiredDeadline();
        }
        bytes32 structHash = keccak256(abi.encode(_PERMIT_TYPEHASH, owner, spender, value, _useNonce(owner), deadline));

        bytes32 hash = _hashTypedDataV4(structHash);

        if (!SignatureCheckerUpgradeable.isValidSignatureNow(owner, hash, abi.encodePacked(r, s, v))){
            revert ErrorInvalidSignature();
        }

        _approve(owner, spender, value);
    }
}
