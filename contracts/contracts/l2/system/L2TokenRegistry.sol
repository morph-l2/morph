// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
import {IL2TokenRegistry} from "./IL2TokenRegistry.sol";

interface IERC20Infos {
    function decimals() external view returns (uint8);
}

/**
 * @title L2TokenRegistry
 * @dev TokenRegistry contract - Used for registering tokenID and managing token information and prices
 * @notice In the transaction scenario where ERC20 is used as gas fee payment, used for storing prices and token registration functionality
 */
contract L2TokenRegistry is IL2TokenRegistry, OwnableUpgradeable, ReentrancyGuardUpgradeable {
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.UintSet;

    /// @notice Mapping from tokenID to TokenInfo
    mapping(uint16 tokenID => TokenInfo tokenInfo) public tokenRegistry;

    /// @notice Mapping from token address to tokenID
    mapping(address tokenAddress => uint16 tokenID) public tokenRegistration;

    /// @notice Mapping from tokenID to price ratio (relative to ETH)
    /// @dev priceRatio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
    mapping(uint16 tokenID => uint256 priceRatio) public priceRatio;

    /// @notice Allow List whitelist
    mapping(address user => bool allowed) public allowList;

    /// @notice Whether whitelist is enabled
    bool public allowListEnabled;

    /// @notice Set of supported token IDs
    EnumerableSetUpgradeable.UintSet private supportedTokenSet;

    /*//////////////////////////////////////////////////////////////
                           Modifier
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Check if caller is in Allow List
     */
    modifier onlyAllowed() {
        bool isOwner = msg.sender == owner();
        bool isAllowedByList = allowListEnabled && allowList[msg.sender];

        // Owner always has access
        // When allowList is enabled, allowList users can access
        // When allowList is disabled, only owner can access
        if (!isOwner && !isAllowedByList) {
            revert CallerNotAllowed();
        }
        _;
    }

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
        __Ownable_init();
        __ReentrancyGuard_init();

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

    /*//////////////////////////////////////////////////////////////
                            Token Registration
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Batch register tokens
     * @param _tokenIDs Array of token IDs
     * @param _tokenAddresses Array of token addresses
     * @param _balanceSlots Array of balance storage slots
     * @param _needBalanceSlots Array of boolean flags indicating whether balanceSlot is needed
     * @param _scales Array of scale values
     */
    function registerTokens(
        uint16[] memory _tokenIDs,
        address[] memory _tokenAddresses,
        bytes32[] memory _balanceSlots,
        bool[] memory _needBalanceSlots,
        uint256[] memory _scales
    ) external onlyOwner nonReentrant {
        if (
            _tokenIDs.length != _tokenAddresses.length ||
            _tokenIDs.length != _balanceSlots.length ||
            _tokenIDs.length != _needBalanceSlots.length ||
            _tokenIDs.length != _scales.length
        ) {
            revert InvalidArrayLength();
        }

        for (uint256 i = 0; i < _tokenIDs.length; i++) {
            _registerSingleToken(_tokenIDs[i], _tokenAddresses[i], _balanceSlots[i], _needBalanceSlots[i], _scales[i]);
        }
    }

    /**
     * @notice Register a single token
     * @param _tokenID Token ID
     * @param _tokenAddress Token contract address
     * @param _balanceSlot Balance storage slot
     * @param _needBalanceSlot Whether balanceSlot is needed (if false, stores 0; if true, stores balanceSlot+1)
     * @param _scale Scale value
     */
    function registerToken(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _needBalanceSlot,
        uint256 _scale
    ) external onlyOwner nonReentrant {
        _registerSingleToken(_tokenID, _tokenAddress, _balanceSlot, _needBalanceSlot, _scale);
    }

    /**
     * @notice Internal function: Add token ID to supported list
     * @param _tokenID Token ID to add
     */
    function _addTokenToList(uint16 _tokenID) internal {
        supportedTokenSet.add(uint256(_tokenID));
    }

    /**
     * @notice Internal function: Remove token ID from supported list
     * @param _tokenID Token ID to remove
     */
    function _removeTokenFromList(uint16 _tokenID) internal {
        supportedTokenSet.remove(uint256(_tokenID));
    }

    /**
     * @notice Internal function: Convert actual balanceSlot to stored value (adds 1 if needed)
     * @param _actualSlot The actual balance slot value
     * @param _needBalanceSlot Whether balanceSlot is needed
     * @return The stored balance slot value (actualSlot + 1 if needed, otherwise 0)
     */
    function _toStoredBalanceSlot(bytes32 _actualSlot, bool _needBalanceSlot) internal pure returns (bytes32) {
        if (!_needBalanceSlot) {
            return bytes32(0); // Don't store balanceSlot
        }
        if (_actualSlot == bytes32(type(uint256).max)) revert InvalidBalanceSlot();
        bytes32 storedSlot;
        assembly {
            storedSlot := add(_actualSlot, 1)
        }
        return storedSlot;
    }

    /**
     * @notice Internal function: Convert stored balanceSlot to actual value (subtracts 1 if non-zero)
     * @param _storedSlot The stored balance slot value
     * @return The actual balance slot value (storedSlot - 1 if non-zero, otherwise 0)
     */
    function _toActualBalanceSlot(bytes32 _storedSlot) internal pure returns (bytes32) {
        if (_storedSlot == bytes32(0)) {
            return bytes32(0); // No balanceSlot stored
        }
        bytes32 actualSlot;
        assembly {
            actualSlot := sub(_storedSlot, 1)
        }
        return actualSlot;
    }

    /**
     * @notice Internal function: Register a single token
     */
    function _registerSingleToken(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _needBalanceSlot,
        uint256 _scale
    ) internal {
        // Check token address
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // Forbid zero ID and enforce uniqueness for both ID and address
        if (_tokenID == 0) revert InvalidTokenID();
        if (tokenRegistry[_tokenID].tokenAddress != address(0)) revert TokenAlreadyRegistered();
        if (tokenRegistration[_tokenAddress] != 0) revert TokenAlreadyRegistered();

        // Validate scale is non-zero
        if (_scale == 0) revert InvalidScale();

        // Get decimals from contract
        uint8 decimals = 18; // Default value
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            decimals = v;
        } catch {
            // If call fails, use default value 18
        }
        
        // Register token (isActive defaults to false)
        // Note: balanceSlot is stored as actualSlot + 1 if needBalanceSlot is true, otherwise 0
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _toStoredBalanceSlot(_balanceSlot, _needBalanceSlot),
            isActive: false,
            decimals: decimals,
            scale: _scale
        });
        tokenRegistration[_tokenAddress] = _tokenID;
        _addTokenToList(_tokenID);
        emit TokenRegistered(_tokenID, _tokenAddress, _balanceSlot, false, decimals, _scale);
    }

    /**
     * @notice Update token information
     * @param _tokenID Token ID
     * @param _tokenAddress New token contract address
     * @param _balanceSlot New balance storage slot
     * @param _needBalanceSlot Whether balanceSlot is needed (if false, stores 0; if true, stores balanceSlot+1)
     * @param _isActive Whether to activate
     * @param _scale Scale value
     */
    function updateTokenInfo(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _needBalanceSlot,
        bool _isActive,
        uint256 _scale
    ) external onlyOwner nonReentrant {
        // Check if token exists
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // Check new information
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // Check new scale
        if (_scale == 0) revert InvalidScale();

        // Prevent address being shared across different tokenIDs
        uint16 existing = tokenRegistration[_tokenAddress];
        if (existing != 0 && existing != _tokenID) revert TokenAlreadyRegistered();

        // Get decimals from contract
        uint8 decimals = 18; // Default value
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            decimals = v;
        } catch {
            // If call fails, use default value 18
        }
        // Update registration information
        // Note: balanceSlot is stored as actualSlot + 1 if needBalanceSlot is true, otherwise 0
        address oldAddress = tokenRegistry[_tokenID].tokenAddress;
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _toStoredBalanceSlot(_balanceSlot, _needBalanceSlot),
            isActive: _isActive,
            decimals: decimals,
            scale: _scale
        });

        // Update address mapping
        if (oldAddress != _tokenAddress) {
            delete tokenRegistration[oldAddress];
            tokenRegistration[_tokenAddress] = _tokenID;
        }

        // Reset priceRatio to 0 to ensure consistency
        // priceRatio depends on scale and decimals, so it must be recalculated after tokenInfo changes
        priceRatio[_tokenID] = 0;

        // Note: tokenID should already be in supportedTokenSet from registration
        // No need to add again as EnumerableSet.add() is idempotent but wastes gas

        emit TokenInfoUpdated(_tokenID, _tokenAddress, _balanceSlot, _isActive, decimals, _scale);
    }

    /**
     * @notice Remove a token from registry
     * @param _tokenID Token ID to remove
     */
    function removeToken(uint16 _tokenID) external onlyOwner {
        // Check if token exists
        address tokenAddress = tokenRegistry[_tokenID].tokenAddress;
        if (tokenAddress == address(0)) revert TokenNotFound();

        // Check if token is in supported list
        if (!supportedTokenSet.contains(uint256(_tokenID))) revert TokenNotFound();

        // Remove from mappings
        delete tokenRegistry[_tokenID];
        delete tokenRegistration[tokenAddress];
        delete priceRatio[_tokenID];

        // Remove from supported list
        _removeTokenFromList(_tokenID);

        emit TokenRemoved(_tokenID, tokenAddress);
    }

    /**
     * @notice Batch update token activation status
     * @param _tokenIDs Array of token IDs
     * @param _isActives Array of activation statuses
     */
    function batchUpdateTokenStatus(uint16[] calldata _tokenIDs, bool[] calldata _isActives) external onlyOwner {
        if (_tokenIDs.length != _isActives.length) revert InvalidArrayLength();

        for (uint256 i = 0; i < _tokenIDs.length; i++) {
            uint16 tokenId = _tokenIDs[i];
            bool newStatus = _isActives[i];

            if (tokenRegistry[tokenId].tokenAddress == address(0)) continue;
            bool oldStatus = tokenRegistry[tokenId].isActive;

            if (oldStatus != newStatus) {
                tokenRegistry[tokenId].isActive = newStatus;
                if (newStatus) {
                    emit TokenActivated(tokenId);
                } else {
                    emit TokenDeactivated(tokenId);
                }
            }
        }
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
     *      - tokenAmount = ⌈(ethAmount × tokenScale) / tokenRate⌉
     *      - Note: Uses ceiling division to ensure users receive fair token amounts
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

        // Calculate token amount with ceiling division:
        // tokenAmount = ceil((ethAmount * tokenScale) / ratio)
        // Using formula: ceil(a/b) = (a + b - 1) / b
        uint256 numerator = _ethAmount * uint256(info.scale);
        tokenAmount = (numerator + ratio - 1) / ratio;
        
        if (tokenAmount == 0) revert ZeroTokenAmount();

        return tokenAmount;
    }

    /**
     * @notice Get token information
     * @param _tokenID Token ID
     * @return TokenInfo structure with actual balanceSlot (automatically -1 from stored value)
     */
    function getTokenInfo(uint16 _tokenID) external view returns (TokenInfo memory) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        TokenInfo memory info = tokenRegistry[_tokenID];
        // Convert stored balanceSlot to actual value
        info.balanceSlot = _toActualBalanceSlot(info.balanceSlot);

        return info;
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

        // Validate scale is non-zero
        if (_newScale == 0) revert InvalidScale();
        tokenRegistry[_tokenID].scale = _newScale;

        // Reset priceRatio to 0 to ensure consistency
        // priceRatio depends on scale, so it must be recalculated after scale changes
        priceRatio[_tokenID] = 0;

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

    /**
     * @notice Check if a token ID is in the supported list
     * @param _tokenID Token ID to check
     * @return Whether the token ID is registered
     */
    function isTokenSupported(uint16 _tokenID) external view returns (bool) {
        return supportedTokenSet.contains(uint256(_tokenID));
    }

    /**
     * @notice Get all supported token IDs and their addresses
     * @return Array of TokenEntry containing token ID and address pairs
     */
    function getSupportedTokenList() external view returns (TokenEntry[] memory) {
        uint256[] memory values = supportedTokenSet.values();
        TokenEntry[] memory tokenList = new TokenEntry[](values.length);

        for (uint256 i = 0; i < values.length; ++i) {
            uint16 tokenID = uint16(values[i]);
            address tokenAddress = tokenRegistry[tokenID].tokenAddress;
            tokenList[i] = TokenEntry({tokenID: tokenID, tokenAddress: tokenAddress});
        }

        return tokenList;
    }

    /**
     * @notice Get all supported token IDs
     * @return Array of all registered token IDs
     */
    function getSupportedIDList() external view returns (uint16[] memory) {
        uint256[] memory values = supportedTokenSet.values();
        uint16[] memory tokenIDs = new uint16[](values.length);

        for (uint256 i = 0; i < values.length; ++i) {
            tokenIDs[i] = uint16(values[i]);
        }

        return tokenIDs;
    }

    /**
     * @notice Get the count of supported tokens
     * @return The number of registered tokens
     */
    function getSupportedTokenCount() external view returns (uint256) {
        return supportedTokenSet.length();
    }

    // Reserve storage space to allow future layout changes
    uint256[50] private __gap;
}
