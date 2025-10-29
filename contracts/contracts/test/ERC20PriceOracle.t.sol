// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";

import {ERC20PriceOracle} from "../l2/system/ERC20PriceOracle.sol";
import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ERC20PriceOracleTest is Test {
    ERC20PriceOracle internal priceOracle;
    ERC20PriceOracle internal priceOracleImpl;
    ProxyAdmin internal proxyAdmin;

    address internal multisig = address(512);
    address internal owner = address(64);
    address internal alice = address(128);
    address internal bob = address(256);

    MockERC20 internal usdc;
    MockERC20 internal usdt;
    MockERC20 internal dai;

    uint16 constant TOKEN_ID_USDC = 1;
    uint16 constant TOKEN_ID_USDT = 2;
    uint16 constant TOKEN_ID_DAI = 3;

    bytes32 constant BALANCE_SLOT_USDC = bytes32(uint256(9));
    bytes32 constant BALANCE_SLOT_USDT = bytes32(uint256(10));
    bytes32 constant BALANCE_SLOT_DAI = bytes32(uint256(11));

    function setUp() public {
        // 部署代理管理
        vm.prank(multisig);
        proxyAdmin = new ProxyAdmin();

        // 部署实现合约
        priceOracleImpl = new ERC20PriceOracle();

        // 部署代理并初始化
        vm.prank(multisig);
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(priceOracleImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(ERC20PriceOracle.initialize.selector, owner)
        );

        priceOracle = ERC20PriceOracle(payable(address(proxy)));

        // 部署 Mock ERC20 tokens
        usdc = new MockERC20("USD Coin", "USDC", 6);
        usdt = new MockERC20("Tether USD", "USDT", 6);
        dai = new MockERC20("Dai Stablecoin", "DAI", 18);

        vm.label(address(usdc), "USDC");
        vm.label(address(usdt), "USDT");
        vm.label(address(dai), "DAI");
        vm.label(address(priceOracle), "ERC20PriceOracle");
        vm.label(multisig, "multisig");
        vm.label(alice, "alice");
        vm.label(bob, "bob");
    }

    /*//////////////////////////////////////////////////////////////
                            Initialization Tests
    //////////////////////////////////////////////////////////////*/

    function test_initialize_succeeds() public {
        assertEq(priceOracle.owner(), owner);
        assertTrue(priceOracle.allowListEnabled());
    }

    function test_initialize_reverts_when_not_called_via_proxy() public {
        ERC20PriceOracle impl = new ERC20PriceOracle();
        vm.expectRevert();
        impl.initialize(owner);
    }

    /*//////////////////////////////////////////////////////////////
                            Token Registration Tests
    //////////////////////////////////////////////////////////////*/

    function test_registerToken_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        ERC20PriceOracle.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.tokenAddress, address(usdc));
        assertEq(info.balanceSlot, BALANCE_SLOT_USDC);
        assertEq(info.isActive, false);
        assertEq(info.decimals, 6);
    }

    function test_registerToken_autoFetchesDecimals() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        ERC20PriceOracle.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.decimals, 6); // USDC has 6 decimals

        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI);

        info = priceOracle.getTokenInfo(TOKEN_ID_DAI);
        assertEq(info.decimals, 18); // DAI has 18 decimals
    }

    function test_registerToken_setsIsActiveToFalse() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        ERC20PriceOracle.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertFalse(info.isActive);
    }

    function test_registerToken_reverts_when_not_owner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(alice);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);
    }

    function test_registerToken_reverts_when_tokenAddress_zero() public {
        vm.expectRevert(bytes4(keccak256("InvalidTokenAddress()")));
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(0), BALANCE_SLOT_USDC);
    }

    function test_registerTokens_succeeds() public {
        uint16[] memory tokenIDs = new uint16[](3);
        address[] memory tokenAddresses = new address[](3);
        bytes32[] memory balanceSlots = new bytes32[](3);

        tokenIDs[0] = TOKEN_ID_USDC;
        tokenIDs[1] = TOKEN_ID_USDT;
        tokenIDs[2] = TOKEN_ID_DAI;

        tokenAddresses[0] = address(usdc);
        tokenAddresses[1] = address(usdt);
        tokenAddresses[2] = address(dai);

        balanceSlots[0] = BALANCE_SLOT_USDC;
        balanceSlots[1] = BALANCE_SLOT_USDT;
        balanceSlots[2] = BALANCE_SLOT_DAI;

        vm.prank(owner);
        priceOracle.registerTokens(tokenIDs, tokenAddresses, balanceSlots);

        assertEq(priceOracle.getTokenInfo(TOKEN_ID_USDC).tokenAddress, address(usdc));
        assertEq(priceOracle.getTokenInfo(TOKEN_ID_USDT).tokenAddress, address(usdt));
        assertEq(priceOracle.getTokenInfo(TOKEN_ID_DAI).tokenAddress, address(dai));
    }

    function test_registerTokens_reverts_when_arrayLength_mismatch() public {
        uint16[] memory tokenIDs = new uint16[](2);
        address[] memory tokenAddresses = new address[](3);
        bytes32[] memory balanceSlots = new bytes32[](2);

        vm.expectRevert(bytes4(keccak256("InvalidArrayLength()")));
        vm.prank(owner);
        priceOracle.registerTokens(tokenIDs, tokenAddresses, balanceSlots);
    }

    function test_getTokenIdByAddress_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        uint16 tokenID = priceOracle.getTokenIdByAddress(address(usdc));
        assertEq(tokenID, TOKEN_ID_USDC);
    }

    function test_getTokenIdByAddress_reverts_when_not_registered() public {
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        priceOracle.getTokenIdByAddress(address(usdc));
    }

    /*//////////////////////////////////////////////////////////////
                            Token Update Tests
    //////////////////////////////////////////////////////////////*/

    function test_updateTokenInfo_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        bytes32 newBalanceSlot = bytes32(uint256(99));
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), newBalanceSlot, true);

        ERC20PriceOracle.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.balanceSlot, newBalanceSlot);
        assertTrue(info.isActive);
    }

    function test_updateTokenInfo_autoFetchesDecimals() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        // 更新为 DAI 地址
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(dai), BALANCE_SLOT_USDC, true);

        ERC20PriceOracle.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.tokenAddress, address(dai));
        assertEq(info.decimals, 18); // Should fetch DAI's decimals
    }

    function test_deactivateToken_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true);

        assertTrue(priceOracle.getTokenInfo(TOKEN_ID_USDC).isActive);

        vm.prank(owner);
        priceOracle.deactivateToken(TOKEN_ID_USDC);

        assertFalse(priceOracle.getTokenInfo(TOKEN_ID_USDC).isActive);
    }

    /*//////////////////////////////////////////////////////////////
                            Price Management Tests
    //////////////////////////////////////////////////////////////*/

    function test_updatePriceRatio_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        // 设置价格：1 USDC = 0.000001 ETH = 1e12 wei
        uint256 priceRatio = 1e12;

        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, priceRatio);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), priceRatio);
    }

    function test_updatePriceRatio_reverts_when_not_allowed() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.expectRevert(bytes4(keccak256("CallerNotAllowed()")));
        vm.prank(alice);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);
    }

    function test_updatePriceRatio_succeeds_when_allowListDisabled() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.prank(owner);
        priceOracle.setAllowListEnabled(false);

        vm.prank(alice);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), 1e12);
    }

    function test_updatePriceRatio_succeeds_when_in_allowList() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        address[] memory users = new address[](1);
        bool[] memory allowed = new bool[](1);
        users[0] = alice;
        allowed[0] = true;

        vm.prank(owner);
        priceOracle.setAllowList(users, allowed);

        vm.prank(alice);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), 1e12);
    }

    function test_updatePriceRatio_reverts_when_invalid_price() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.expectRevert(bytes4(keccak256("InvalidPrice()")));
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 0);
    }

    function test_batchUpdatePrices_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT);

        uint16[] memory tokenIDs = new uint16[](2);
        uint256[] memory prices = new uint256[](2);

        tokenIDs[0] = TOKEN_ID_USDC;
        tokenIDs[1] = TOKEN_ID_USDT;
        prices[0] = 1e12;
        prices[1] = 1e12;

        vm.prank(owner);
        priceOracle.batchUpdatePrices(tokenIDs, prices);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), 1e12);
        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDT), 1e12);
    }

    /*//////////////////////////////////////////////////////////////
                            Gas Price Calculation Tests
    //////////////////////////////////////////////////////////////*/

    function test_calculateTokenGasPrice_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        // 设置价格：1 USDC = 0.000001 ETH = 1e12 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        // ETH gas price = 1 gwei = 1e9 wei
        uint256 ethGasPrice = 1 gwei;
        uint256 expectedTokenGasPrice = (ethGasPrice * 1e6) / 1e12; // (1e9 * 1e6) / 1e12 = 1e3

        uint256 tokenGasPrice = priceOracle.calculateTokenGasPrice(TOKEN_ID_USDC, ethGasPrice);
        assertEq(tokenGasPrice, expectedTokenGasPrice);
    }

    function test_calculateEthGasPrice_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        // 设置价格：1 USDC = 0.000001 ETH = 1e12 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        // Token gas price = 1000 USDC
        uint256 tokenGasPrice = 1000;
        uint256 expectedEthGasPrice = (tokenGasPrice * 1e12) / 1e6; // (1000 * 1e12) / 1e6 = 1e9

        uint256 ethGasPrice = priceOracle.calculateEthGasPrice(TOKEN_ID_USDC, tokenGasPrice);
        assertEq(ethGasPrice, expectedEthGasPrice);
    }

    function test_calculateTokenGasPrice_withDAI() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI);

        // 设置价格：1 DAI = 0.001 ETH = 1e15 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_DAI, 1e15);

        // ETH gas price = 1 gwei = 1e9 wei
        uint256 ethGasPrice = 1 gwei;
        uint256 expectedTokenGasPrice = (ethGasPrice * 1e18) / 1e15; // (1e9 * 1e18) / 1e15 = 1e12

        uint256 tokenGasPrice = priceOracle.calculateTokenGasPrice(TOKEN_ID_DAI, ethGasPrice);
        assertEq(tokenGasPrice, expectedTokenGasPrice);
    }

    /*//////////////////////////////////////////////////////////////
                            Fee Discount Tests
    //////////////////////////////////////////////////////////////*/

    function test_updateFeeDiscountPercent_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        uint256 discountPercent = 500; // 5%

        vm.prank(owner);
        priceOracle.updateFeeDiscountPercent(TOKEN_ID_USDC, discountPercent);

        assertEq(priceOracle.getFeeDiscountPercent(TOKEN_ID_USDC), discountPercent);
    }

    function test_updateFeeDiscountPercent_reverts_when_exceeds_100_percent() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.expectRevert(bytes4(keccak256("InvalidPercent()")));
        vm.prank(owner);
        priceOracle.updateFeeDiscountPercent(TOKEN_ID_USDC, 10001); // > 100%
    }

    function test_updateFeeDiscountPercent_reverts_when_not_allowed() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        vm.expectRevert(bytes4(keccak256("CallerNotAllowed()")));
        vm.prank(alice);
        priceOracle.updateFeeDiscountPercent(TOKEN_ID_USDC, 500);
    }

    /*//////////////////////////////////////////////////////////////
                            Allow List Tests
    //////////////////////////////////////////////////////////////*/

    function test_setAllowList_succeeds() public {
        address[] memory users = new address[](2);
        bool[] memory allowed = new bool[](2);

        users[0] = alice;
        users[1] = bob;
        allowed[0] = true;
        allowed[1] = false;

        vm.prank(owner);
        priceOracle.setAllowList(users, allowed);

        assertTrue(priceOracle.allowList(alice));
        assertFalse(priceOracle.allowList(bob));
    }

    function test_setAllowList_reverts_when_different_length() public {
        address[] memory users = new address[](2);
        bool[] memory allowed = new bool[](1);

        vm.expectRevert(bytes4(keccak256("DifferentLength()")));
        vm.prank(owner);
        priceOracle.setAllowList(users, allowed);
    }

    function test_setAllowListEnabled_succeeds() public {
        vm.prank(owner);
        priceOracle.setAllowListEnabled(false);

        assertFalse(priceOracle.allowListEnabled());

        vm.prank(owner);
        priceOracle.setAllowListEnabled(true);

        assertTrue(priceOracle.allowListEnabled());
    }

    /*//////////////////////////////////////////////////////////////
                            View Functions Tests
    //////////////////////////////////////////////////////////////*/

    function test_isTokenActive_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC);

        assertFalse(priceOracle.isTokenActive(TOKEN_ID_USDC));

        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true);

        assertTrue(priceOracle.isTokenActive(TOKEN_ID_USDC));
    }

    function test_isTokenActive_returns_false_for_nonexistent_token() public {
        assertFalse(priceOracle.isTokenActive(TOKEN_ID_USDC));
    }
}
