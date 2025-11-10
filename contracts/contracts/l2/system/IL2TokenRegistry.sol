// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

/**
 * @title IL2TokenRegistry
 * @dev Interface for L2TokenRegistry contract
 * @notice Interface defining all external functions for ERC20 price oracle and token registry
 */
interface IL2TokenRegistry {
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

    /// @notice Token entry structure containing ID and address
    struct TokenEntry {
        uint16 tokenID; // Token ID
        address tokenAddress; // ERC20 token contract address
    }

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
    event TokenInfoUpdated(
        uint16 indexed tokenID,
        address indexed tokenAddress,
        bytes32 balanceSlot,
        bool isActive,
        uint8 decimals,
        uint256 scale
    );
    event TokenActivated(uint16 indexed tokenID);
    event TokenDeactivated(uint16 indexed tokenID);
    event TokenRemoved(uint16 indexed tokenID, address indexed tokenAddress);
    event PriceRatioUpdated(uint16 indexed tokenID, uint256 newPrice);
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
    error InvalidPrice();
    error InvalidPercent();
    error CallerNotAllowed();
    error InvalidArrayLength();
    error DifferentLength();
    error AlreadyInitialized();

    /*//////////////////////////////////////////////////////////////
                            Allow List Functions
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Set Allow List
     * @param user Array of user addresses
     * @param val Array of permission values
     */
    function setAllowList(address[] memory user, bool[] memory val) external;

    /**
     * @notice Set whether Allow List is enabled
     * @param _allowListEnabled Whether to enable
     */
    function setAllowListEnabled(bool _allowListEnabled) external;

    /*//////////////////////////////////////////////////////////////
                            Token Registration Functions
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
    ) external;

    /**
     * @notice Register a single token
     * @param _tokenID Token ID
     * @param _tokenAddress Token contract address
     * @param _balanceSlot Balance storage slot
     * @param _scale Scale value
     */
    function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot, uint256 _scale) external;

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
    ) external;

    /**
     * @notice Remove a token from registry
     * @param _tokenID Token ID to remove
     */
    function removeToken(uint16 _tokenID) external;

    /**
     * @notice Batch update token activation status
     * @param _tokenIDs Array of token IDs
     * @param _isActives Array of activation statuses
     */
    function batchUpdateTokenStatus(uint16[] memory _tokenIDs, bool[] memory _isActives) external;

    /*//////////////////////////////////////////////////////////////
                            Price Management Functions
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Update price ratio
     * @param _tokenID Token ID
     * @param _newPrice New price ratio (relative to ETH)
     * @dev priceRatio should follow: priceRatio = tokenScale * (tokenPrice / ethPrice) * 10^(ethDecimals - tokenDecimals)
     */
    function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) external;
    
    /**
     * @notice Batch update price ratios
     * @param _tokenIDs Array of token IDs
     * @param _prices Array of price ratios
     */
    function batchUpdatePrices(uint16[] memory _tokenIDs, uint256[] memory _prices) external;

    /**
     * @notice Get token price
     * @param _tokenID Token ID
     * @return price Price ratio
     */
    function getTokenPrice(uint16 _tokenID) external view returns (uint256);

    /**
     * @notice Calculate the corresponding token amount for a given ETH amount
     * @param _tokenID Token ID of the ERC20 token
     * @param _ethAmount ETH amount (unit: wei)
     * @return tokenAmount Corresponding token amount (unit: token's smallest unit)
     */
    function calculateTokenAmount(uint16 _tokenID, uint256 _ethAmount) external view returns (uint256 tokenAmount);

    /**
     * @notice Get token information
     * @param _tokenID Token ID
     * @return TokenInfo structure
     */
    function getTokenInfo(uint16 _tokenID) external view returns (TokenInfo memory);

    /**
     * @notice Get token ID by address
     * @param tokenAddress Token address
     * @return tokenID Token ID
     */
    function getTokenIdByAddress(address tokenAddress) external view returns (uint16);

    /*//////////////////////////////////////////////////////////////
                            Scale Management Functions
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Update token scale
     * @param _tokenID Token ID
     * @param _newScale New scale value
     */
    function updateTokenScale(uint16 _tokenID, uint256 _newScale) external;

    /**
     * @notice Get token scale
     * @param _tokenID Token ID
     * @return scale Token scale value
     */
    function getTokenScale(uint16 _tokenID) external view returns (uint256);

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice Check if token is active
     * @param _tokenID Token ID
     * @return Whether the token is active
     */
    function isTokenActive(uint16 _tokenID) external view returns (bool);

    /**
     * @notice Check if a token ID is in the supported list
     * @param _tokenID Token ID to check
     * @return Whether the token ID is registered
     */
    function isTokenSupported(uint16 _tokenID) external view returns (bool);

    /**
     * @notice Get all supported token IDs and their addresses
     * @return Array of TokenEntry containing token ID and address pairs
     */
    function getSupportedTokenList() external view returns (TokenEntry[] memory);

    /**
     * @notice Get all supported token IDs
     * @return Array of all registered token IDs
     */
    function getSupportedIDList() external view returns (uint16[] memory);

    /**
     * @notice Get the count of supported tokens
     * @return The number of registered tokens
     */
    function getSupportedTokenCount() external view returns (uint256);
}

