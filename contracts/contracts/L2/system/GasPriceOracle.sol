// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title GasPriceOracle contract
 * @dev Entrance to the update method for L1 gas prices
 **/
contract GasPriceOracle is Ownable {
    /*//////////////////////////////////////////////////////////////
                               Constants
    //////////////////////////////////////////////////////////////*/
    /// @dev The precision used in the scalar.
    uint256 private constant PRECISION = 1e9;

    /*//////////////////////////////////////////////////////////////
                               STORAGE
    //////////////////////////////////////////////////////////////*/
    // Current L1 base fee
    uint256 public l1BaseFee;
    // Amortized cost of batch submission per transaction
    uint256 public overhead;
    // Value to scale the fee up by
    uint256 public scalar;
    // Enable permission list
    bool public allowListEnabled;
    // Address list with permission to update price oracle.
    mapping(address => bool) public isAllowed;

    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/
    event AllowListAddressSet(address indexed user, bool val);
    event AllowListEnabledUpdated(bool isEnabled);
    event L1BaseFeeUpdated(uint256);
    event OverheadUpdated(uint256);
    event ScalarUpdated(uint256);

    /**
     * @param owner_ Address that will initially own this contract.
     */
    constructor(address owner_) Ownable() {
        transferOwnership(owner_);
        allowListEnabled = true;
    }

    /*//////////////////////////////////////////////////////////////
                             ALLOWLIST
    //////////////////////////////////////////////////////////////*/
    function setAllowList(
        address[] memory user,
        bool[] memory val
    ) external onlyOwner {
        require(user.length == val.length, "INVALID_INPUT");

        for (uint256 i = 0; i < user.length; i++) {
            isAllowed[user[i]] = val[i];
            emit AllowListAddressSet(user[i], val[i]);
        }
    }

    function setAllowListEnabled(bool _allowListEnabled) external onlyOwner {
        require(_allowListEnabled != allowListEnabled, "ALREADY_SET");
        allowListEnabled = _allowListEnabled;
        emit AllowListEnabledUpdated(_allowListEnabled);
    }

    modifier onlyAllowed() {
        // solhint-disable-next-line avoid-tx-origin
        require(
            owner() == msg.sender || !allowListEnabled || isAllowed[msg.sender],
            "not allowed"
        );
        _;
    }

    function getL1Fee(bytes memory _data) external view returns (uint256) {
        uint256 _l1GasUsed = getL1GasUsed(_data);
        uint256 _l1Fee = _l1GasUsed * l1BaseFee;
        return (_l1Fee * scalar) / PRECISION;
    }

    /// @dev The `_data` is the RLP-encoded transaction with signature. And we also reserve additional
    ///      4 bytes in the non-zero bytes to store the number of bytes in the RLP-encoded transaction.
    function getL1GasUsed(bytes memory _data) public view returns (uint256) {
        uint256 _total = 0;
        uint256 _length = _data.length;
        unchecked {
            for (uint256 i = 0; i < _length; i++) {
                if (_data[i] == 0) {
                    _total += 4;
                } else {
                    _total += 16;
                }
            }
            return _total + overhead + (4 * 16);
        }
    }

    /**
     * Allows the owner to modify the l1 base fee.
     * @param _baseFee New l1 base fee
     */
    // slither-disable-next-line external-function
    function setL1BaseFee(uint256 _baseFee) external onlyAllowed {
        l1BaseFee = _baseFee;
        emit L1BaseFeeUpdated(_baseFee);
    }

    /**
     * Allows the owner to modify the overhead.
     * @param _overhead New overhead
     */
    // slither-disable-next-line external-function
    function setOverhead(uint256 _overhead) external onlyAllowed {
        overhead = _overhead;
        emit OverheadUpdated(_overhead);
    }

    /**
     * Allows the owner to modify the scalar.
     * @param _scalar New scalar
     */
    // slither-disable-next-line external-function
    function setScalar(uint256 _scalar) external onlyAllowed {
        scalar = _scalar;
        emit ScalarUpdated(_scalar);
    }
}
