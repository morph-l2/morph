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
    }

    /*//////////////////////////////////////////////////////////////
                               Storage
    //////////////////////////////////////////////////////////////*/

    /// @notice Mapping from tokenID to TokenInfo
    mapping(uint16 => TokenInfo) public tokenRegistry;

    /// @notice Mapping from token address to tokenID
    mapping(address => uint16) public tokenRegistration;

    /// @notice Mapping from tokenID to price ratio (relative to ETH)
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
        uint8 decimals
    );
    event TokensRegistered(uint16[] tokenIDs, address[] tokenAddresses);
    event TokenInfoUpdated(
        uint16 indexed tokenID,
        address indexed tokenAddress,
        bytes32 balanceSlot,
        bool isActive,
        uint8 decimals
    );
    event TokenDeactivated(uint16 indexed tokenID);
    event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice);
    event FeeDiscountPercentUpdated(uint16 indexed tokenID, uint256 newPercent);
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
     */
    function registerTokens(
        uint16[] memory _tokenIDs,
        address[] memory _tokenAddresses,
        bytes32[] memory _balanceSlots
    ) external onlyOwner {
        if (_tokenIDs.length != _tokenAddresses.length || _tokenIDs.length != _balanceSlots.length) {
            revert InvalidArrayLength();
        }

        for (uint256 i = 0; i < _tokenIDs.length; i++) {
            _registerSingleToken(_tokenIDs[i], _tokenAddresses[i], _balanceSlots[i]);
        }

        emit TokensRegistered(_tokenIDs, _tokenAddresses);
    }

    /**
     * @notice Register a single token
     * @param _tokenID Token ID
     * @param _tokenAddress Token contract address
     * @param _balanceSlot Balance storage slot
     */
    function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot) external onlyOwner {
        _registerSingleToken(_tokenID, _tokenAddress, _balanceSlot);

        TokenInfo memory info = tokenRegistry[_tokenID];
        emit TokenRegistered(_tokenID, _tokenAddress, _balanceSlot, info.isActive, info.decimals);
    }

    /**
     * @notice Internal function: Register a single token
     */
    function _registerSingleToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot) internal {
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
            decimals: decimals
        });
        tokenRegistration[_tokenAddress] = _tokenID;
    }

    /**
     * @notice Update token information
     * @param _tokenID Token ID
     * @param _tokenAddress New token contract address
     * @param _balanceSlot New balance storage slot
     * @param _isActive Whether to activate
     */
    function updateTokenInfo(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _isActive
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
            decimals: decimals
        });

        // Update address mapping
        if (oldAddress != _tokenAddress) {
            delete tokenRegistration[oldAddress];
            tokenRegistration[_tokenAddress] = _tokenID;
        }

        emit TokenInfoUpdated(_tokenID, _tokenAddress, _balanceSlot, _isActive, decimals);
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
     * @notice Calculate the gas price for a specified ERC20 token as gas fee
     * @dev Calculation formula: tokenGasPrice = (ethGasPrice * 10^decimals) / priceRatio
     * @param _tokenID Token ID of the ERC20 token
     * @param _ethGasPrice ETH gas price (unit: wei)
     * @return tokenGasPrice Corresponding ERC20 token gas price (unit: token's smallest unit)
     * - First scale ethGasPrice by 10^decimals to compensate for token precision
     * - Then divide by the token's current priceRatio
     * - Will revert if token is not registered or priceRatio is not set
     */
    function calculateTokenGasPrice(
        uint16 _tokenID,
        uint256 _ethGasPrice
    ) external view returns (uint256 tokenGasPrice) {
        // Validate: token must be registered
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Get token's ETH price ratio (priceRatio) and precision (decimals)
        uint256 ratio = priceRatio[_tokenID];
        if (ratio == 0) revert InvalidPrice();

        uint8 decimals = tokenRegistry[_tokenID].decimals;

        // Scale precision: ethGasPrice * 10^decimals
        uint256 scaledPrice = _ethGasPrice * (10 ** decimals);

        // Convert to token price
        tokenGasPrice = scaledPrice / ratio;

        return tokenGasPrice;
    }

    /**
     * @notice Calculate corresponding ETH gas price from ERC20 token gas price
     * @param _tokenID ERC20 token ID
     * @param _tokenGasPrice ERC20 token gas price (token unit)
     * @return ethGasPrice ETH gas price (wei unit)
     * @dev Price calculation formula:
     *      - ethGasPrice = (tokenGasPrice * priceRatio) / 10^decimals
     */
    function calculateEthGasPrice(uint16 _tokenID, uint256 _tokenGasPrice) external view returns (uint256 ethGasPrice) {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Get priceRatio and decimals
        uint256 ratio = priceRatio[_tokenID];
        if (ratio == 0) revert InvalidPrice();

        uint8 decimals = tokenRegistry[_tokenID].decimals;

        // Calculate: eth gas price = (token gas price * priceRatio) / 10^decimals
        uint256 scaledPrice = _tokenGasPrice * ratio;
        ethGasPrice = scaledPrice / (10 ** decimals);

        return ethGasPrice;
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
