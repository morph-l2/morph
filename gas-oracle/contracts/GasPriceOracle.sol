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

    /// @dev The maximum possible l1 fee overhead.
    ///      Computed based on current l1 block gas limit.
    uint256 private constant MAX_OVERHEAD = 30000000 / 16;

    /// @dev The maximum possible l1 fee scale.
    ///      x1000 should be enough.
    uint256 private constant MAX_SCALE = 1000 * PRECISION;

    /// @dev The maximum possible l1 commit fee scalar.
    /// We derive the commit scalar by
    /// ```
    /// commit_scalar = commit_gas_per_tx * fluctuation_multiplier * 1e9
    /// ```
    /// So, the value should not exceed 10^9 * 1e9 normally.
    uint256 private constant MAX_COMMIT_SCALAR = 10 ** 9 * PRECISION;

    /// @dev The maximum possible l1 blob fee scalar.
    /// We derive the blob scalar by
    /// ```
    /// blob_scalar = fluctuation_multiplier / compression_ratio / blob_util_ratio * 1e9
    /// ```
    /// So, the value should not exceed 10^9 * 1e9 normally.
    uint256 private constant MAX_BLOB_SCALAR = 10 ** 9 * PRECISION;

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
    /// l1 blob base fee
    uint256 public l1BlobBaseFee;
    /// commit scalar
    uint256 public commitScalar;
    /// blob scalar
    uint256 public blobScalar;

    /*//////////////////////////////////////////////////////////////
                                EVENTS
    //////////////////////////////////////////////////////////////*/
    event AllowListAddressSet(address indexed user, bool val);
    event AllowListEnabledUpdated(bool isEnabled);
    event L1BaseFeeUpdated(uint256);
    event OverheadUpdated(uint256);
    event ScalarUpdated(uint256);
    event L1BlobBaseFeeUpdated(uint256 l1BlobBaseFee);
    event CommitScalarUpdated(uint256 scalar);
    event BlobScalarUpdated(uint256 scalar);

    /*//////////////////////////////////////////////////////////////
                               Errors 
    //////////////////////////////////////////////////////////////*/
    /// @dev Thrown when the blob fee scalar exceed MAX_BLOB_SCALAR.
    error ErrExceedMaxBlobScalar();

    /// @dev Thrown when the commit fee scalar exceed MAX_COMMIT_SCALAR.
    error ErrExceedMaxCommitScalar();

    /// @dev Thrown when the l1 fee overhead exceed MAX_OVERHEAD.
    error ErrExceedMaxOverhead();

    /// @dev Thrown when the l1 fee scalar exceed MAX_SCALAR.
    error ErrExceedMaxScalar();

    /// @dev Thrown when the caller is not in allowed list.
    error ErrCallerNotAllowed();

    /// @dev Thrown when setting the same value.
    error ErrSettintSameValue();

    /// @dev Thrown when setting values to different lengths for keys.
    error ErrDifferentLength();

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
        if (user.length != val.length) revert ErrDifferentLength();

        for (uint256 i = 0; i < user.length; i++) {
            isAllowed[user[i]] = val[i];
            emit AllowListAddressSet(user[i], val[i]);
        }
    }

    function setAllowListEnabled(bool _allowListEnabled) external onlyOwner {
        if (_allowListEnabled == allowListEnabled) revert ErrSettintSameValue();

        allowListEnabled = _allowListEnabled;
        emit AllowListEnabledUpdated(_allowListEnabled);
    }

    modifier onlyAllowed() {
        // solhint-disable-next-line avoid-tx-origin
        if (!(owner() == msg.sender || (allowListEnabled && isAllowed[msg.sender]))) {
            revert ErrCallerNotAllowed();
        }
        _;
    }

    /// @dev External function to compute the L1 portion of the fee based on the size of the rlp encoded input
    ///      transaction, the current L1 base fee, and the various dynamic parameters.
    /// @param _data Signed fully RLP-encoded transaction to get the L1 fee for.
    /// @return L1 fee that should be paid for the tx
    function getL1Fee(bytes memory _data) external view returns (uint256) {
        // We have bounded the value of `commitScalar` and `blobScalar`, the whole expression won't overflow.
        return (commitScalar * l1BaseFee + blobScalar * _data.length * l1BlobBaseFee) / PRECISION;
    }

    /// @param _data The `_data` is the RLP-encoded transaction with signature. And we also reserve additional
    /// 4 bytes in the non-zero bytes to store the number of bytes in the RLP-encoded transaction.
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

    /// Allows the owner to modify the l1 base fee.
    /// @param _l1BaseFee New l1 base fee
    // slither-disable-next-line external-function
    function setL1BaseFee(uint256 _l1BaseFee) external onlyAllowed {
        l1BaseFee = _l1BaseFee;
        emit L1BaseFeeUpdated(_l1BaseFee);
    }

    /// @notice Allows the owner to modify the l1 base fee.
    /// @param _l1BaseFee New l1 base fee
    /// @param _l1BlobBaseFee New l1 blob base fee
    // slither-disable-next-line external-function
    function setL1BaseFeeAndBlobBaseFee(uint256 _l1BaseFee, uint256 _l1BlobBaseFee) external onlyAllowed {
        l1BaseFee = _l1BaseFee;
        l1BlobBaseFee = _l1BlobBaseFee;

        emit L1BaseFeeUpdated(_l1BaseFee);
        emit L1BlobBaseFeeUpdated(_l1BlobBaseFee);
    }

    /// @notice Allows the owner to modify the overhead.
    /// @param _overhead New overhead
    // slither-disable-next-line external-function
    function setOverhead(uint256 _overhead) external onlyAllowed {
        if (_overhead > MAX_OVERHEAD) revert ErrExceedMaxOverhead();

        overhead = _overhead;
        emit OverheadUpdated(_overhead);
    }

    /// @notice Allows the owner to modify the scalar.
    /// @param _scalar New scalar
    // slither-disable-next-line external-function
    function setScalar(uint256 _scalar) external onlyAllowed {
        if (_scalar > MAX_SCALE) revert ErrExceedMaxScalar();

        scalar = _scalar;
        emit ScalarUpdated(_scalar);
    }

    /// @notice Allows the owner to modify the commit scalar.
    /// @param _scalar New scalar
    function setCommitScalar(uint256 _scalar) external onlyAllowed {
        if (_scalar > MAX_COMMIT_SCALAR) revert ErrExceedMaxCommitScalar();

        commitScalar = _scalar;
        emit CommitScalarUpdated(_scalar);
    }

    /// Allows the owner to modify the blob scalar.
    /// @param _scalar New scalar
    function setBlobScalar(uint256 _scalar) external onlyAllowed {
        if (_scalar > MAX_BLOB_SCALAR) revert ErrExceedMaxBlobScalar();

        blobScalar = _scalar;
        emit BlobScalarUpdated(_scalar);
    }
}