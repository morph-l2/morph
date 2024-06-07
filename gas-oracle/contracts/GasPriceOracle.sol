// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title GasPriceOracle contract
 * @custom:predeploy 0x530000000000000000000000000000000000000f
 * @dev Entrance to the update method for L1 gas prices
 **/
contract GasPriceOracle is Ownable {
    /*//////////////////////////////////////////////////////////////
                               STORAGE
    //////////////////////////////////////////////////////////////*/
    // Current L1 base fee
    uint256 public l1BaseFee;
    // Amortized cost of batch submission per transaction
    uint256 public overhead;
    // Value to scale the fee up by
    uint256 public scalar;

    uint256 public l1BlobBaseFee;

    uint256 public commitScalar;

    uint256 public blobScalar;
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
    function setAllowList(address[] memory user, bool[] memory val) external onlyOwner {
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
        require(owner() == msg.sender || !allowListEnabled || isAllowed[msg.sender], "not allowed");
        _;
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

    function setL1BaseFeeAndBlobBaseFee(
        uint256 _l1BaseFee,
        uint256 _l1BlobBaseFee
    ) external onlyAllowed {
        l1BaseFee = _l1BaseFee;
        l1BlobBaseFee = _l1BlobBaseFee;
    }

    /// Allows the owner to modify the commit scalar.
    /// @param _scalar New scalar
    function setCommitScalar(uint256 _scalar) external onlyAllowed {
        commitScalar = _scalar;
    }

    /// Allows the owner to modify the blob scalar.
    /// @param _scalar New scalar
    function setBlobScalar(uint256 _scalar) external onlyAllowed {
        blobScalar = _scalar;
    }
}
