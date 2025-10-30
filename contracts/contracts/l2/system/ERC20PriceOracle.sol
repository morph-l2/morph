// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IERC20Infos {
    function decimals() external view returns (uint8);
}

/**
 * @title ERC20PriceOracle
 * @dev TokenRegistry contract - Used for registering tokenID and managing token information and prices
 * @notice In the transaction scenario where ERC20 is used as gas fee payment, used for storing prices and token registration functionality
 */
contract ERC20PriceOracle is OwnableUpgradeable {
    /*//////////////////////////////////////////////////////////////
                               Structs
    //////////////////////////////////////////////////////////////*/

    /// @notice Token information structure
    struct TokenInfo {
        address tokenAddress; // ERC20 token contract address
        bytes32 balanceSlot; // Token balance storage slot, bytes32(0) -> nil
        bool isActive; // Whether the token is active
        uint8 decimals; // Token decimals
        uint256 scale; // Core convention: rateScaled = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
    }

    /*//////////////////////////////////////////////////////////////
                               Storage
    //////////////////////////////////////////////////////////////*/

    /// @notice Mapping from tokenID to TokenInfo
    mapping(uint16 => TokenInfo) public tokenRegistry;

    /// @notice Mapping from token address to tokenID
    mapping(address => uint16) public tokenRegistration;

    /// @notice Mapping from tokenID to price ratio (relative to ETH)
    /// @dev priceRatio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
    mapping(uint16 => uint256) public priceRatio;

    /// @notice Mapping from tokenID to fee discount percentage
    mapping(uint16 => uint256) public feeDiscountPercent;

    /// @notice Allow List whitelist
    mapping(address => bool) public allowList;

    /// @notice Whether whitelist is enabled
    bool public allowListEnabled = true;

    /*//////////////////////////////////////////////////////////////
                               Events
    //////////////////////////////////////////////////////////////*/

    event TokenRegistered(
        uint16 indexed tokenID,
        address indexed tokenAddress,
        bytes32 balanceSlot,
        bool isActive,
        uint8 decimals,
        uint256 scale
    );
    event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses);
    event TokenInfoUpdated(
        uint16 indexed tokenID,
        address indexed tokenAddress,
        bytes32 balanceSlot,
        bool isActive,
        uint8 decimals,
        uint256 scale
    );
    event TokenDeactivated(uint16 indexed tokenID);
    event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice);
    event FeeDiscountPercentUpdated(uint16 indexed tokenID, uint256 newPercent);
    event TokenScaleUpdated(uint16 indexed tokenID, uint256 newScale);
    event AllowListSet(address indexed user, bool val);
    event AllowListEnabledUpdated(bool isEnabled);

    /*//////////////////////////////////////////////////////////////
                               Errors
    //////////////////////////////////////////////////////////////*/

    error TokenAlreadyRegistered();
    error TokenNotFound();
    error InvalidTokenID();
    error InvalidTokenAddress();
    error InvalidDecimals();
    error InvalidPrice();
    error InvalidPercent();
    error CallerNotAllowed();
    error InvalidArrayLength();
    error DifferentLength();
    error AlreadyInitialized();

    /*//////////////////////////////////////////////////////////////
                             Initializer
    //////////////////////////////////////////////////////////////*/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param owner_ Contract owner address
     */
    function initialize(address owner_) external initializer {
        _transferOwnership(owner_);
        allowListEnabled = true;
    }

    /*//////////////////////////////////////////////////////////////
                            Allow List
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Set Allow List
     * @param user Array of user addresses
     * @param val Array of permission values
     */
    function setAllowList(address[] memory user, bool[] memory val) external onlyOwner {
        if (user.length != val.length) revert DifferentLength();

        for (uint256 i = 0; i < user.length; i++) {
            allowList[user[i]] = val[i];
            emit AllowListSet(user[i], val[i]);
        }
    }

    /**
     * @notice Set whether Allow List is enabled
     * @param _allowListEnabled Whether to enable
     */
    function setAllowListEnabled(bool _allowListEnabled) external onlyOwner {
        allowListEnabled = _allowListEnabled;
        emit AllowListEnabledUpdated(_allowListEnabled);
    }

    /**
     * @notice Check if caller is in Allow List
     */
    modifier onlyAllowed() {
        if (allowListEnabled && !allowList[msg.sender] && msg.sender != owner()) {
            revert CallerNotAllowed();
        }
        _;
    }

    /*//////////////////////////////////////////////////////////////
                            Token Registration
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Batch register tokens
     * @param _tokenIDs Array of token IDs
     * @param _tokenAddresses Array of token addresses
     * @param _balanceSlots Array of balance storage slots
     * @param _scales Array of scale values
     */
    function registerTokens(
        uint16[] memory _tokenIDs,
        address[] memory _tokenAddresses,
        bytes32[] memory _balanceSlots,
        uint256[] memory _scales
    ) external onlyOwner {
        if (
            _tokenIDs.length != _tokenAddresses.length ||
            _tokenIDs.length != _balanceSlots.length ||
            _tokenIDs.length != _scales.length
        ) {
            revert InvalidArrayLength();
        }

        for (uint256 i = 0; i < _tokenIDs.length; i++) {
            _registerSingleToken(_tokenIDs[i], _tokenAddresses[i], _balanceSlots[i], _scales[i]);
        }

        emit TokensRegistered(_tokenIDs, _tokenAddresses);
    }

    /**
     * @notice Register a single token
     * @param _tokenID Token ID
     * @param _tokenAddress Token contract address
     * @param _balanceSlot Balance storage slot
     * @param _scale Scale value
     */
    function registerToken(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        uint256 _scale
    ) external onlyOwner {
        _registerSingleToken(_tokenID, _tokenAddress, _balanceSlot, _scale);

        TokenInfo memory info = tokenRegistry[_tokenID];
        emit TokenRegistered(_tokenID, _tokenAddress, _balanceSlot, info.isActive, info.decimals, _scale);
    }

    /**
     * @notice Internal function: Register a single token
     */
    function _registerSingleToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) internal {
        // Check token address
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // Check if already registered
        if (tokenRegistry[_tokenID].tokenAddress == address(0) && tokenRegistration[_tokenAddress] != 0) {
            revert TokenAlreadyRegistered();
        }

        // Get decimals from contract
        uint8 decimals = 18; // Default value
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            if (v > 18) revert InvalidDecimals();
            decimals = v;
        } catch {
            // If call fails, use default value 18
        }

        // Register token (isActive defaults to false)
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _balanceSlot,
            isActive: false,
            decimals: decimals,
            scale: _scale
        });
        tokenRegistration[_tokenAddress] = _tokenID;
    }

    /**
     * @notice Update token information
     * @param _tokenID Token ID
     * @param _tokenAddress New token contract address
     * @param _balanceSlot New balance storage slot
     * @param _isActive Whether to activate
     * @param _scale Scale value
     */
    function updateTokenInfo(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _isActive,
        uint256 _scale
    ) external onlyOwner {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Check new information
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // Get decimals from contract
        uint8 decimals = 18; // Default value
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            if (v > 18) revert InvalidDecimals();
            decimals = v;
        } catch {
            // If call fails, use default value 18
        }

        // Update registration information
        address oldAddress = tokenRegistry[_tokenID].tokenAddress;
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _balanceSlot,
            isActive: _isActive,
            decimals: decimals,
            scale: _scale
        });

        // Update address mapping
        if (oldAddress != _tokenAddress) {
            delete tokenRegistration[oldAddress];
            tokenRegistration[_tokenAddress] = _tokenID;
        }

        emit TokenInfoUpdated(_tokenID, _tokenAddress, _balanceSlot, _isActive, decimals, _scale);
    }

    /**
     * @notice Deactivate token
     * @param _tokenID Token ID
     */
    function deactivateToken(uint16 _tokenID) external onlyOwner {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Deactivate token
        tokenRegistry[_tokenID].isActive = false;

        emit TokenDeactivated(_tokenID);
    }

    /*//////////////////////////////////////////////////////////////
                            Price Management
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Update price ratio
     * @param _tokenID Token ID
     * @param _newPrice New price ratio (relative to ETH)
     * @dev priceRatio should follow: priceRatio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
     */
    function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) external onlyAllowed {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        if (_newPrice == 0) revert InvalidPrice();

        priceRatio[_tokenID] = _newPrice;

        emit PriceRatioUpdated(_tokenID, _newPrice);
    }

    /**
     * @notice Batch update price ratios
     * @param _tokenIDs Array of token IDs
     * @param _prices Array of price ratios
     */
    function batchUpdatePrices(uint16[] memory _tokenIDs, uint256[] memory _prices) external onlyAllowed {
        if (_tokenIDs.length != _prices.length) revert InvalidArrayLength();

        for (uint256 i = 0; i < _tokenIDs.length; i++) {
            if (tokenRegistry[_tokenIDs[i]].tokenAddress == address(0)) continue;
            if (_prices[i] == 0) continue;

            priceRatio[_tokenIDs[i]] = _prices[i];
            emit PriceRatioUpdated(_tokenIDs[i], _prices[i]);
        }
    }

    /**
     * @notice Get token price
     * @param _tokenID Token ID
     * @return price Price ratio
     */
    function getTokenPrice(uint16 _tokenID) external view returns (uint256) {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        return priceRatio[_tokenID];
    }

    /**
     * @notice Calculate the corresponding token amount for a given ETH amount
     * @dev Calculation formula:
     *      - ratio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
     *      - tokenAmount = (ethAmount * 10^tokenDecimals) / ratio
     *      - Substituting ratio: tokenAmount = (ethAmount * 10^tokenDecimals) / (tokenScale * (tokenPrice / ethPrice) * 10^(18 - tokenDecimals))
     *      - Simplified: tokenAmount = (ethAmount * 10^tokenDecimals * 10^tokenDecimals) / (tokenScale * tokenPrice * 10^18 / ethPrice)
     *      - Final: tokenAmount = (ethAmount * ethPrice * 10^tokenDecimals) / (tokenScale * tokenPrice * 10^18)
     * @param _tokenID Token ID of the ERC20 token
     * @param _ethAmount ETH amount (unit: wei)
     * @return tokenAmount Corresponding token amount (unit: token's smallest unit)
     * - ratio follows: ratio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
     * - Will revert if token is not registered or priceRatio is not set
     */
    function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) external view returns (uint256 tokenAmount) {
        // Validate: token must be registered
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Get token information
        TokenInfo memory info = tokenRegistry[_tokenID];

        // Get priceRatio which follows:
        // ratio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
        uint256 ratio = priceRatio[_tokenID];
        if (ratio == 0) revert InvalidPrice();

        // Calculate token amount:
        // tokenAmount = (ethAmount * tokenScale) / ratio
        // where ratio already contains tokenScale and decimals adjustment to eth (wei) and token smallest unit.
        tokenAmount = (_ethAmount * uint256(info.scale)) / ratio;

        return tokenAmount;
    }

    /**
     * @notice Get token information
     * @param _tokenID Token ID
     * @return TokenInfo structure
     */
    function getTokenInfo(uint16 _tokenID) external view returns (TokenInfo memory) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();
        return tokenRegistry[_tokenID];
    }

    /**
     * @notice Get token ID by address
     * @param tokenAddress Token address
     * @return tokenID Token ID
     */
    function getTokenIdByAddress(address tokenAddress) external view returns (uint16) {
        uint16 tokenID = tokenRegistration[tokenAddress];
        if (tokenID == 0 && tokenAddress != address(0)) revert TokenNotFound();
        return tokenID;
    }

    /*//////////////////////////////////////////////////////////////
                            Fee Discount Management
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Update fee discount percentage
     * @param _tokenID Token ID
     * @param _newPercent New fee discount percentage (basis points, 10000 = 100%)
     */
    function updateFeeDiscountPercent(uint16 _tokenID, uint256 _newPercent) external onlyAllowed {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Check percentage range (0-100%)
        if (_newPercent > 10000) revert InvalidPercent();

        feeDiscountPercent[_tokenID] = _newPercent;

        emit FeeDiscountPercentUpdated(_tokenID, _newPercent);
    }

    /**
     * @notice Get fee discount percentage
     * @param _tokenID Token ID
     * @return percent Fee discount percentage
     */
    function getFeeDiscountPercent(uint16 _tokenID) external view returns (uint256) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();
        return feeDiscountPercent[_tokenID];
    }

    /*//////////////////////////////////////////////////////////////
                            Scale Management
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Update token scale
     * @param _tokenID Token ID
     * @param _newScale New scale value
     * @dev Core convention: rateScaled = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
     */
    function updateTokenScale(uint16 _tokenID, uint256 _newScale) external onlyAllowed {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        tokenRegistry[_tokenID].scale = _newScale;

        emit TokenScaleUpdated(_tokenID, _newScale);
    }

    /**
     * @notice Get token scale
     * @param _tokenID Token ID
     * @return scale Token scale value
     */
    function getTokenScale(uint16 _tokenID) external view returns (uint256) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();
        return tokenRegistry[_tokenID].scale;
    }

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/
    /**
     * @notice Check if token is active
     * @param _tokenID Token ID
     * @return Whether the token is active
     */
    function isTokenActive(uint16 _tokenID) external view returns (bool) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) return false;
        return tokenRegistry[_tokenID].isActive;
    }
}
