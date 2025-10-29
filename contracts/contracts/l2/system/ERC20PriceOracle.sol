// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

interface IERC20Infos {
    function decimals() external view returns (uint8);
}

/**
 * @title ERC20PriceOracle
 * @dev TokenRegistry 合约 - 用于注册 tokenID、管理代币信息和价格
 * @notice 在以 ERC20 作为 gas 费用支付的交易场景中，用于存储价格和 token 注册功能
 */
contract ERC20PriceOracle is OwnableUpgradeable {
    /*//////////////////////////////////////////////////////////////
                               Structs
    //////////////////////////////////////////////////////////////*/

    /// @notice Token 信息结构
    struct TokenInfo {
        address tokenAddress; // ERC20 代币合约地址
        bytes32 balanceSlot; // 代币余额存储槽 bytes32(0) -> nil
        bool isActive; // 代币是否激活
        uint8 decimals; // 代币精度
    }

    /*//////////////////////////////////////////////////////////////
                               Storage
    //////////////////////////////////////////////////////////////*/

    /// @notice tokenID 到 TokenInfo 的映射
    mapping(uint16 => TokenInfo) public tokenRegistry;

    /// @notice token 地址到 tokenID 的映射
    mapping(address => uint16) public tokenRegistration;

    /// @notice tokenID 到价格比率的映射（相对于 ETH）
    mapping(uint16 => uint256) public priceRatio;

    /// @notice tokenID 到手续费减免百分比的映射
    mapping(uint16 => uint256) public feeDiscountPercent;

    /// @notice Allow List 白名单
    mapping(address => bool) public allowList;

    /// @notice 是否启用白名单
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
     * @notice 初始化函数，用于代理部署
     * @param owner_ 合约所有者地址
     */
    function initialize(address owner_) external initializer {
        _transferOwnership(owner_);
        allowListEnabled = true;
    }

    /*//////////////////////////////////////////////////////////////
                            Allow List
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice 设置 Allow List
     * @param user 用户地址数组
     * @param val 权限值数组
     */
    function setAllowList(address[] memory user, bool[] memory val) external onlyOwner {
        if (user.length != val.length) revert DifferentLength();

        for (uint256 i = 0; i < user.length; i++) {
            allowList[user[i]] = val[i];
            emit AllowListSet(user[i], val[i]);
        }
    }

    /**
     * @notice 设置是否启用 Allow List
     * @param _allowListEnabled 是否启用
     */
    function setAllowListEnabled(bool _allowListEnabled) external onlyOwner {
        allowListEnabled = _allowListEnabled;
        emit AllowListEnabledUpdated(_allowListEnabled);
    }

    /**
     * @notice 检查是否在 Allow List 中
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
     * @notice 批量注册 token
     * @param _tokenIDs token ID 数组
     * @param _tokenAddresses token 地址数组
     * @param _balanceSlots 余额存储槽数组
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
     * @notice 注册单个 token
     * @param _tokenID token ID
     * @param _tokenAddress token 合约地址
     * @param _balanceSlot 余额存储槽
     */
    function registerToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot) external onlyOwner {
        _registerSingleToken(_tokenID, _tokenAddress, _balanceSlot);

        TokenInfo memory info = tokenRegistry[_tokenID];
        emit TokenRegistered(_tokenID, _tokenAddress, _balanceSlot, info.isActive, info.decimals);
    }

    /**
     * @notice 内部函数：注册单个 token
     */
    function _registerSingleToken(uint16 _tokenID, address _tokenAddress, bytes32 _balanceSlot) internal {
        // 检查 token 地址
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // 检查是否已注册
        if (tokenRegistry[_tokenID].tokenAddress == address(0) && tokenRegistration[_tokenAddress] != 0) {
            revert TokenAlreadyRegistered();
        }

        // 从合约获取 decimals
        uint8 decimals = 18; // 默认值
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            if (v > 18) revert InvalidDecimals();
            decimals = v;
        } catch {
            // 如果调用失败，使用默认值 18
        }

        // 注册 token（isActive 默认为 false）
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _balanceSlot,
            isActive: false,
            decimals: decimals
        });
        tokenRegistration[_tokenAddress] = _tokenID;
    }

    /**
     * @notice 更新 token 信息
     * @param _tokenID token ID
     * @param _tokenAddress 新的 token 合约地址
     * @param _balanceSlot 新的余额存储槽
     * @param _isActive 是否激活
     */
    function updateTokenInfo(
        uint16 _tokenID,
        address _tokenAddress,
        bytes32 _balanceSlot,
        bool _isActive
    ) external onlyOwner {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // 检查新信息
        if (_tokenAddress == address(0)) revert InvalidTokenAddress();

        // 从合约获取 decimals
        uint8 decimals = 18; // 默认值
        try IERC20Infos(_tokenAddress).decimals() returns (uint8 v) {
            if (v > 18) revert InvalidDecimals();
            decimals = v;
        } catch {
            // 如果调用失败，使用默认值 18
        }

        // 更新注册信息
        address oldAddress = tokenRegistry[_tokenID].tokenAddress;
        tokenRegistry[_tokenID] = TokenInfo({
            tokenAddress: _tokenAddress,
            balanceSlot: _balanceSlot,
            isActive: _isActive,
            decimals: decimals
        });

        // 更新地址映射
        if (oldAddress != _tokenAddress) {
            delete tokenRegistration[oldAddress];
            tokenRegistration[_tokenAddress] = _tokenID;
        }

        emit TokenInfoUpdated(_tokenID, _tokenAddress, _balanceSlot, _isActive, decimals);
    }

    /**
     * @notice 停用 token
     * @param _tokenID token ID
     */
    function deactivateToken(uint16 _tokenID) external onlyOwner {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // 停用 token
        tokenRegistry[_tokenID].isActive = false;

        emit TokenDeactivated(_tokenID);
    }

    /*//////////////////////////////////////////////////////////////
                            Price Management
    //////////////////////////////////////////////////////////////*/

    /**
     * @notice 更新价格比率
     * @param _tokenID token ID
     * @param _newPrice 新的价格比率（相对于 ETH）
     */
    function updatePriceRatio(uint16 _tokenID, uint256 _newPrice) external onlyAllowed {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        if (_newPrice == 0) revert InvalidPrice();

        priceRatio[_tokenID] = _newPrice;

        emit PriceRatioUpdated(_tokenID, _newPrice);
    }

    /**
     * @notice 批量更新价格比率
     * @param _tokenIDs token ID 数组
     * @param _prices 价格比率数组
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
     * @notice 获取 token 价格
     * @param _tokenID token ID
     * @return price 价格比率
     */
    function getTokenPrice(uint16 _tokenID) external view returns (uint256) {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        return priceRatio[_tokenID];
    }

    /**
     * @notice 计算指定 ERC20 Token 作为 gas 费用的价格
     * @dev 计算公式：tokenGasPrice = (ethGasPrice * 10^decimals) / priceRatio
     * @param _tokenID ERC20 代币的 token ID
     * @param _ethGasPrice ETH 的 gas price（单位: wei）
     * @return tokenGasPrice 对应的 ERC20 token gas price（单位: token 的最小单位）
     * - 先将 ethGasPrice 扩大 10^decimals 倍以补偿代币精度
     * - 然后除以该 token 当前设定的 priceRatio
     * - 若 token 未注册或 priceRatio 未设置，将抛出相应异常
     */
    function calculateTokenGasPrice(
        uint16 _tokenID,
        uint256 _ethGasPrice
    ) external view returns (uint256 tokenGasPrice) {
        // 校验：token 必须已注册
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // 获取代币对应的 ETH 价格比率（priceRatio）及精度（decimals）
        uint256 ratio = priceRatio[_tokenID];
        if (ratio == 0) revert InvalidPrice();

        uint8 decimals = tokenRegistry[_tokenID].decimals;

        // 扩大精度：ethGasPrice * 10^decimals
        uint256 scaledPrice = _ethGasPrice * (10 ** decimals);

        // 转换为 token 价格
        tokenGasPrice = scaledPrice / ratio;

        return tokenGasPrice;
    }

    /**
     * @notice 根据 ERC20 token gas 价格计算对应的 ETH gas 价格
     * @param _tokenID ERC20 token ID
     * @param _tokenGasPrice ERC20 token gas 价格（token 单位）
     * @return ethGasPrice ETH gas 价格（wei 单位）
     * @dev 价格计算公式：
     *      - ethGasPrice = (tokenGasPrice * priceRatio) / 10^decimals
     */
    function calculateEthGasPrice(uint16 _tokenID, uint256 _tokenGasPrice) external view returns (uint256 ethGasPrice) {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // 获取 priceRatio 和 decimals
        uint256 ratio = priceRatio[_tokenID];
        if (ratio == 0) revert InvalidPrice();

        uint8 decimals = tokenRegistry[_tokenID].decimals;

        // 计算：eth gas price = (token gas price * priceRatio) / 10^decimals
        uint256 scaledPrice = _tokenGasPrice * ratio;
        ethGasPrice = scaledPrice / (10 ** decimals);

        return ethGasPrice;
    }

    /**
     * @notice 获取 token 信息
     * @param _tokenID token ID
     * @return TokenInfo 结构
     */
    function getTokenInfo(uint16 _tokenID) external view returns (TokenInfo memory) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();
        return tokenRegistry[_tokenID];
    }

    /**
     * @notice 通过地址获取 token ID
     * @param tokenAddress token 地址
     * @return tokenID token ID
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
     * @notice 更新手续费减免百分比
     * @param _tokenID token ID
     * @param _newPercent 新的手续费减免百分比（基点，10000 = 100%）
     */
    function updateFeeDiscountPercent(uint16 _tokenID, uint256 _newPercent) external onlyAllowed {
        // 检查 token 是否存在
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();

        // 检查百分比范围（0-100%）
        if (_newPercent > 10000) revert InvalidPercent();

        feeDiscountPercent[_tokenID] = _newPercent;

        emit FeeDiscountPercentUpdated(_tokenID, _newPercent);
    }

    /**
     * @notice 获取手续费减免百分比
     * @param _tokenID token ID
     * @return percent 手续费减免百分比
     */
    function getFeeDiscountPercent(uint16 _tokenID) external view returns (uint256) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) revert TokenNotFound();
        return feeDiscountPercent[_tokenID];
    }

    /*//////////////////////////////////////////////////////////////
                            View Functions
    //////////////////////////////////////////////////////////////*/
    /**
     * @notice 检查 token 是否激活
     * @param _tokenID token ID
     * @return 是否激活
     */
    function isTokenActive(uint16 _tokenID) external view returns (bool) {
        if (tokenRegistry[_tokenID].tokenAddress == address(0)) return false;
        return tokenRegistry[_tokenID].isActive;
    }
}
