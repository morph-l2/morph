// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "forge-std/Test.sol";

import {L2TokenRegistry} from "../l2/system/L2TokenRegistry.sol";
import {IL2TokenRegistry} from "../l2/system/IL2TokenRegistry.sol";
import {MockERC20} from "@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {ProxyAdmin} from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract L2TokenRegistryTest is Test {
    L2TokenRegistry internal priceOracle;
    L2TokenRegistry internal priceOracleImpl;
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

    uint256 constant SCALE_USDC = 1e6; // 10^6
    uint256 constant SCALE_USDT = 1e6; // 10^6
    uint256 constant SCALE_DAI = 1e18; // 10^18

    function setUp() public {
        // Deploy proxy admin
        vm.prank(multisig);
        proxyAdmin = new ProxyAdmin();

        // Deploy implementation contract
        priceOracleImpl = new L2TokenRegistry();

        // Deploy proxy and initialize
        vm.prank(multisig);
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(priceOracleImpl),
            address(proxyAdmin),
            abi.encodeWithSelector(L2TokenRegistry.initialize.selector, owner)
        );

        priceOracle = L2TokenRegistry(payable(address(proxy)));

        // Deploy Mock ERC20 tokens
        usdc = new MockERC20("USD Coin", "USDC", 6);
        usdt = new MockERC20("Tether USD", "USDT", 6);
        dai = new MockERC20("Dai Stablecoin", "DAI", 18);

        vm.label(address(usdc), "USDC");
        vm.label(address(usdt), "USDT");
        vm.label(address(dai), "DAI");
        vm.label(address(priceOracle), "L2TokenRegistry");
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
        L2TokenRegistry impl = new L2TokenRegistry();
        vm.expectRevert();
        impl.initialize(owner);
    }

    /*//////////////////////////////////////////////////////////////
                            Token Registration Tests
    //////////////////////////////////////////////////////////////*/

    function test_registerToken_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.tokenAddress, address(usdc));
        assertEq(info.balanceSlot, BALANCE_SLOT_USDC);
        assertEq(info.isActive, false);
        assertEq(info.decimals, 6);
    }

    function test_registerToken_reverts_when_tokenID_is_zero() public {
        vm.expectRevert(bytes4(keccak256("InvalidTokenID()")));
        vm.prank(owner);
        priceOracle.registerToken(0, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
    }

    function test_registerToken_reverts_when_tokenID_already_registered() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectRevert(bytes4(keccak256("TokenIDAlreadyRegistered()")));
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);
    }

    function test_registerToken_reverts_when_address_already_registered() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectRevert(bytes4(keccak256("TokenAddressAlreadyRegistered()")));
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdc), BALANCE_SLOT_USDT, true, SCALE_USDT);
    }

    function test_registerToken_autoFetchesDecimals() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.decimals, 6); // USDC has 6 decimals

        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI, true, SCALE_DAI);

        info = priceOracle.getTokenInfo(TOKEN_ID_DAI);
        assertEq(info.decimals, 18); // DAI has 18 decimals
    }

    function test_registerToken_setsIsActiveToFalse() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertFalse(info.isActive);
    }

    function test_registerToken_reverts_when_not_owner() public {
        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(alice);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
    }

    function test_registerToken_reverts_when_tokenAddress_zero() public {
        vm.expectRevert(bytes4(keccak256("InvalidTokenAddress()")));
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(0), BALANCE_SLOT_USDC, true, SCALE_USDC);
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

        bool[] memory needBalanceSlots = new bool[](3);
        needBalanceSlots[0] = true;
        needBalanceSlots[1] = true;
        needBalanceSlots[2] = true;

        uint256[] memory scales = new uint256[](3);
        scales[0] = SCALE_USDC;
        scales[1] = SCALE_USDT;
        scales[2] = SCALE_DAI;

        vm.prank(owner);
        priceOracle.registerTokens(tokenIDs, tokenAddresses, balanceSlots, needBalanceSlots, scales);

        assertEq(priceOracle.getTokenInfo(TOKEN_ID_USDC).tokenAddress, address(usdc));
        assertEq(priceOracle.getTokenInfo(TOKEN_ID_USDT).tokenAddress, address(usdt));
        assertEq(priceOracle.getTokenInfo(TOKEN_ID_DAI).tokenAddress, address(dai));
    }

    function test_registerTokens_reverts_when_arrayLength_mismatch() public {
        uint16[] memory tokenIDs = new uint16[](2);
        address[] memory tokenAddresses = new address[](3);
        bytes32[] memory balanceSlots = new bytes32[](2);
        bool[] memory needBalanceSlots = new bool[](2);
        uint256[] memory scales = new uint256[](2);

        vm.expectRevert(bytes4(keccak256("InvalidArrayLength()")));
        vm.prank(owner);
        priceOracle.registerTokens(tokenIDs, tokenAddresses, balanceSlots, needBalanceSlots, scales);
    }

    function test_getTokenIdByAddress_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        uint16 tokenID = priceOracle.getTokenIdByAddress(address(usdc));
        assertEq(tokenID, TOKEN_ID_USDC);
    }

    function test_getTokenIdByAddress_reverts_when_not_registered() public {
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        priceOracle.getTokenIdByAddress(address(usdc));
    }

    /*//////////////////////////////////////////////////////////////
                            BalanceSlot Storage Tests
    //////////////////////////////////////////////////////////////*/

    function test_balanceSlot_storage_query_with_minus_one() public {
        // Register token with balanceSlot = 9
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Get balanceSlot through getTokenInfo (should return actual value = 9)
        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.balanceSlot, BALANCE_SLOT_USDC);

        // Read balanceSlot directly from storage
        // tokenRegistry is at slot 151
        // TokenInfo struct layout:
        // - slot 0: tokenAddress (20 bytes)
        // - slot 1: balanceSlot (32 bytes)
        // - slot 2: isActive (1 byte) + decimals (1 byte) + scale (32 bytes packed)
        uint256 mappingSlot = 151;
        
        // Calculate storage location: keccak256(tokenID || mappingSlot)
        bytes32 key = keccak256(abi.encode(TOKEN_ID_USDC, mappingSlot));
        
        // balanceSlot is stored in key + 1
        bytes32 balanceSlotStorageLocation = bytes32(uint256(key) + 1);
        
        // Read stored value from storage
        bytes32 storedBalanceSlot = vm.load(address(priceOracle), balanceSlotStorageLocation);
        
        // Stored value should be actualSlot + 1 = 9 + 1 = 10
        assertEq(uint256(storedBalanceSlot), uint256(BALANCE_SLOT_USDC) + 1);
        
        // Apply -1 to get actual value
        bytes32 actualBalanceSlot = bytes32(uint256(storedBalanceSlot) - 1);
        
        // Verify that manual -1 gives us the same value as getTokenInfo
        assertEq(actualBalanceSlot, BALANCE_SLOT_USDC);
        assertEq(actualBalanceSlot, info.balanceSlot);
    }

    function test_balanceSlot_storage_query_with_slot_zero() public {
        // Test with balanceSlot = 0 (edge case)
        bytes32 balanceSlot0 = bytes32(uint256(0));
        
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), balanceSlot0, true, SCALE_USDT);

        // Get balanceSlot through getTokenInfo (should return actual value = 0)
        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDT);
        assertEq(info.balanceSlot, balanceSlot0);

        // Read balanceSlot directly from storage
        uint256 mappingSlot = 151;
        bytes32 key = keccak256(abi.encode(TOKEN_ID_USDT, mappingSlot));
        bytes32 balanceSlotStorageLocation = bytes32(uint256(key) + 1);
        bytes32 storedBalanceSlot = vm.load(address(priceOracle), balanceSlotStorageLocation);
        
        // Stored value should be actualSlot + 1 = 0 + 1 = 1
        assertEq(uint256(storedBalanceSlot), 1);
        
        // Apply -1 to get actual value
        bytes32 actualBalanceSlot = bytes32(uint256(storedBalanceSlot) - 1);
        
        // Verify that manual -1 gives us 0
        assertEq(actualBalanceSlot, balanceSlot0);
        assertEq(uint256(actualBalanceSlot), 0);
        assertEq(actualBalanceSlot, info.balanceSlot);
    }

    function test_balanceSlot_storage_query_multiple_tokens() public {
        // Register multiple tokens with different balanceSlots
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), bytes32(uint256(9)), true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), bytes32(uint256(10)), true, SCALE_USDT);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), bytes32(uint256(11)), true, SCALE_DAI);

        uint256 mappingSlot = 151;

        // Verify USDC: stored=10, actual=9
        bytes32 key = keccak256(abi.encode(TOKEN_ID_USDC, mappingSlot));
        bytes32 storedValue = vm.load(address(priceOracle), bytes32(uint256(key) + 1));
        assertEq(uint256(storedValue), 10);
        assertEq(bytes32(uint256(storedValue) - 1), priceOracle.getTokenInfo(TOKEN_ID_USDC).balanceSlot);

        // Verify USDT: stored=11, actual=10
        key = keccak256(abi.encode(TOKEN_ID_USDT, mappingSlot));
        storedValue = vm.load(address(priceOracle), bytes32(uint256(key) + 1));
        assertEq(uint256(storedValue), 11);
        assertEq(bytes32(uint256(storedValue) - 1), priceOracle.getTokenInfo(TOKEN_ID_USDT).balanceSlot);

        // Verify DAI: stored=12, actual=11
        key = keccak256(abi.encode(TOKEN_ID_DAI, mappingSlot));
        storedValue = vm.load(address(priceOracle), bytes32(uint256(key) + 1));
        assertEq(uint256(storedValue), 12);
        assertEq(bytes32(uint256(storedValue) - 1), priceOracle.getTokenInfo(TOKEN_ID_DAI).balanceSlot);
    }

    function test_balanceSlot_storage_query_needBalanceSlot_false() public {
        // Test with needBalanceSlot = false (token doesn't need balanceSlot)
        bytes32 anySlot = bytes32(uint256(999));  // Value doesn't matter when needBalanceSlot = false
        uint16 tokenID = 100;
        
        vm.prank(owner);
        priceOracle.registerToken(tokenID, address(usdc), anySlot, false, SCALE_USDC);

        // Get balanceSlot through getTokenInfo (should return 0 because needBalanceSlot was false)
        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(tokenID);
        assertEq(info.balanceSlot, bytes32(0));

        // Read balanceSlot directly from storage
        uint256 mappingSlot = 151;
        bytes32 key = keccak256(abi.encode(tokenID, mappingSlot));
        bytes32 balanceSlotStorageLocation = bytes32(uint256(key) + 1);
        bytes32 storedBalanceSlot = vm.load(address(priceOracle), balanceSlotStorageLocation);
        
        // When needBalanceSlot = false, stored value should be 0 (no +1)
        assertEq(uint256(storedBalanceSlot), 0);
        
        // getTokenInfo should return 0 (no -1 conversion needed)
        assertEq(info.balanceSlot, bytes32(0));
    }

    function test_balanceSlot_reverts_when_max_uint256() public {
        // Test that registering with max uint256 as balanceSlot reverts
        bytes32 maxSlot = bytes32(type(uint256).max);
        uint16 tokenID = 101;
        
        vm.expectRevert(bytes4(keccak256("InvalidBalanceSlot()")));
        vm.prank(owner);
        priceOracle.registerToken(tokenID, address(usdc), maxSlot, true, SCALE_USDC);
    }

    function test_registerToken_reverts_when_scale_is_zero() public {
        // Test that registering with scale = 0 reverts
        bytes32 balanceSlot = bytes32(uint256(9));
        uint16 tokenID = 102;
        
        vm.expectRevert(bytes4(keccak256("InvalidScale()")));
        vm.prank(owner);
        priceOracle.registerToken(tokenID, address(usdc), balanceSlot, true, 0);
    }

    /*//////////////////////////////////////////////////////////////
                            Token Update Tests
    //////////////////////////////////////////////////////////////*/

    function test_updateTokenInfo_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        bytes32 newBalanceSlot = bytes32(uint256(99));
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), newBalanceSlot, true, true, SCALE_USDC);

        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.balanceSlot, newBalanceSlot);
        assertTrue(info.isActive);
    }

    function test_updateTokenInfo_reverts_when_address_collision() public {
        // Register two tokens
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);

        // Try to update USDT to use USDC's address - should revert
        vm.expectRevert(bytes4(keccak256("TokenAddressAlreadyRegistered()")));
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDT, address(usdc), BALANCE_SLOT_USDT, true, true, SCALE_USDT);
    }
    
    function test_updateTokenInfo_autoFetchesDecimals() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Update to DAI address
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(dai), BALANCE_SLOT_USDC, true, true, SCALE_DAI);

        L2TokenRegistry.TokenInfo memory info = priceOracle.getTokenInfo(TOKEN_ID_USDC);
        assertEq(info.tokenAddress, address(dai));
        assertEq(info.decimals, 18); // Should fetch DAI's decimals
    }

    function test_deactivateToken_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, true, SCALE_USDC);

        assertTrue(priceOracle.getTokenInfo(TOKEN_ID_USDC).isActive);

        // Use batchUpdateTokenStatus to deactivate token
        uint16[] memory tokenIDs = new uint16[](1);
        bool[] memory isActives = new bool[](1);
        tokenIDs[0] = TOKEN_ID_USDC;
        isActives[0] = false;

        vm.prank(owner);
        priceOracle.batchUpdateTokenStatus(tokenIDs, isActives);

        assertFalse(priceOracle.getTokenInfo(TOKEN_ID_USDC).isActive);
    }

    /*//////////////////////////////////////////////////////////////
                            Price Management Tests
    //////////////////////////////////////////////////////////////*/

    function test_updatePriceRatio_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Set price: 1 USDC = 0.000001 ETH = 1e12 wei
        uint256 priceRatio = 1e12;

        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, priceRatio);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), priceRatio);
    }

    function test_updatePriceRatio_reverts_when_not_allowed() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectRevert(bytes4(keccak256("CallerNotAllowed()")));
        vm.prank(alice);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);
    }

    function test_updatePriceRatio_succeeds_when_allowListDisabled() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.prank(owner);
        priceOracle.setAllowListEnabled(false);

        vm.prank(alice);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        assertEq(priceOracle.getTokenPrice(TOKEN_ID_USDC), 1e12);
    }

    function test_updatePriceRatio_succeeds_when_in_allowList() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

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
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectRevert(bytes4(keccak256("InvalidPrice()")));
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 0);
    }

    function test_batchUpdatePrices_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);

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
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Set price: 1 USDC = 0.000001 ETH = 1e12 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        // ETH gas price = 1 gwei = 1e9 wei
        uint256 ethGasPrice = 1 gwei;
        uint256 expectedTokenAmount = (ethGasPrice * SCALE_USDC) / 1e12; // (1e9 * 1e6) / 1e12 = 1e3

        uint256 tokenGasAmount = priceOracle.calculateTokenAmount(TOKEN_ID_USDC, ethGasPrice);
        assertEq(tokenGasAmount, expectedTokenAmount);
    }

    function test_calculateEthGasPrice_succeeds() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Set price: 1 USDC = 0.000001 ETH = 1e12 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        // Token gas price = 1000 USDC
        uint256 tokenGasPrice = 1000;
        uint256 expectedEthGasPrice = (tokenGasPrice * 1e12) / SCALE_USDC; // (1000 * 1e12) / 1e6 = 1e9

        // Inverse using on-chain values
        uint256 ratio = priceOracle.getTokenPrice(TOKEN_ID_USDC);
        uint256 scale = priceOracle.getTokenInfo(TOKEN_ID_USDC).scale;
        uint256 ethGasPrice = (tokenGasPrice * ratio) / scale;
        assertEq(ethGasPrice, expectedEthGasPrice);
    }

    function test_calculateTokenGasPrice_withDAI() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI, true, SCALE_DAI);

        // Set price: 1 DAI = 0.001 ETH = 1e15 wei
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_DAI, 1e15);

        // ETH gas price = 1 gwei = 1e9 wei
        uint256 ethGasPrice = 1 gwei;
        uint256 expectedTokenGasPrice = (ethGasPrice * SCALE_DAI) / 1e15; // (1e9 * 1e18) / 1e15 = 1e12

        uint256 tokenGasPrice = priceOracle.calculateTokenAmount(TOKEN_ID_DAI, ethGasPrice);
        assertEq(tokenGasPrice, expectedTokenGasPrice);
    }

    function test_calculateTokenAmount_ceiling_division() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Set price ratio that will result in a remainder
        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 7e11); // Custom ratio for testing

        // Test case 1: Division with remainder
        // numerator = 10 * 1e6 = 1e7
        // ratio = 7e11
        // Floor division: 1e7 / 7e11 = 0 (rounds down)
        // Ceiling division: ceil(1e7 / 7e11) = 1 (rounds up)
        uint256 ethAmount1 = 10;
        uint256 tokenAmount1 = priceOracle.calculateTokenAmount(TOKEN_ID_USDC, ethAmount1);
        uint256 numerator1 = ethAmount1 * SCALE_USDC;
        uint256 expectedCeiling1 = (numerator1 + 7e11 - 1) / 7e11;
        assertEq(tokenAmount1, expectedCeiling1);
        assertGt(tokenAmount1, numerator1 / 7e11); // Should be greater than floor division

        // Test case 2: Exact division (no remainder)
        // numerator = 7e11 * 1e6 = 7e17
        // ratio = 7e11
        // Division: 7e17 / 7e11 = 1e6
        uint256 ethAmount2 = 7e11;
        uint256 tokenAmount2 = priceOracle.calculateTokenAmount(TOKEN_ID_USDC, ethAmount2);
        assertEq(tokenAmount2, 1e6);

        // Test case 3: Large amount with small remainder
        // numerator = 1e15 * 1e6 = 1e21
        // ratio = 7e11
        // Floor: 1e21 / 7e11 = 1428571428571 (approximately 1.43e12)
        // Ceiling: ceil(1e21 / 7e11) = 1428571428572
        uint256 ethAmount3 = 1e15;
        uint256 tokenAmount3 = priceOracle.calculateTokenAmount(TOKEN_ID_USDC, ethAmount3);
        uint256 numerator3 = ethAmount3 * SCALE_USDC;
        uint256 expectedCeiling3 = (numerator3 + 7e11 - 1) / 7e11;
        assertEq(tokenAmount3, expectedCeiling3);
        
        // Verify ceiling behavior: result should be greater than floor when there's a remainder
        if (numerator3 % 7e11 > 0) {
            assertGt(tokenAmount3, numerator3 / 7e11);
        }
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
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        assertFalse(priceOracle.isTokenActive(TOKEN_ID_USDC));

        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, true, SCALE_USDC);

        assertTrue(priceOracle.isTokenActive(TOKEN_ID_USDC));
    }

    function test_isTokenActive_returns_false_for_nonexistent_token() public {
        assertFalse(priceOracle.isTokenActive(TOKEN_ID_USDC));
    }

    /*//////////////////////////////////////////////////////////////
                            Supported Token List Tests
    //////////////////////////////////////////////////////////////*/

    function test_isTokenSupported_returns_false_when_not_registered() public {
        assertFalse(priceOracle.isTokenSupported(TOKEN_ID_USDC));
    }

    function test_isTokenSupported_returns_true_when_registered() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));
    }

    function test_getSupportedTokenCount_returns_zero_initially() public {
        assertEq(priceOracle.getSupportedTokenCount(), 0);
    }

    function test_getSupportedTokenCount_increments_on_register() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        assertEq(priceOracle.getSupportedTokenCount(), 1);

        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);
        assertEq(priceOracle.getSupportedTokenCount(), 2);

        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI, true, SCALE_DAI);
        assertEq(priceOracle.getSupportedTokenCount(), 3);
    }

    function test_getSupportedIDList_returns_empty_when_no_tokens() public {
        uint16[] memory tokenIDs = priceOracle.getSupportedIDList();
        assertEq(tokenIDs.length, 0);
    }

    function test_getSupportedIDList_returns_all_registered_tokenIDs() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_DAI, address(dai), BALANCE_SLOT_DAI, true, SCALE_DAI);

        uint16[] memory tokenIDs = priceOracle.getSupportedIDList();
        assertEq(tokenIDs.length, 3);
        
        // Check that all token IDs are present (order may vary)
        bool foundUSDC = false;
        bool foundUSDT = false;
        bool foundDAI = false;
        
        for (uint256 i = 0; i < tokenIDs.length; ++i) {
            if (tokenIDs[i] == TOKEN_ID_USDC) foundUSDC = true;
            if (tokenIDs[i] == TOKEN_ID_USDT) foundUSDT = true;
            if (tokenIDs[i] == TOKEN_ID_DAI) foundDAI = true;
        }
        
        assertTrue(foundUSDC);
        assertTrue(foundUSDT);
        assertTrue(foundDAI);
    }

    function test_getSupportedTokenList_returns_empty_when_no_tokens() public {
        L2TokenRegistry.TokenEntry[] memory tokenList = priceOracle.getSupportedTokenList();
        assertEq(tokenList.length, 0);
    }

    function test_getSupportedTokenList_returns_all_registered_tokens() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);

        L2TokenRegistry.TokenEntry[] memory tokenList = priceOracle.getSupportedTokenList();
        assertEq(tokenList.length, 2);

        // Check that all tokens are present with correct addresses
        bool foundUSDC = false;
        bool foundUSDT = false;

        for (uint256 i = 0; i < tokenList.length; ++i) {
            if (tokenList[i].tokenID == TOKEN_ID_USDC) {
                assertEq(tokenList[i].tokenAddress, address(usdc));
                foundUSDC = true;
            }
            if (tokenList[i].tokenID == TOKEN_ID_USDT) {
                assertEq(tokenList[i].tokenAddress, address(usdt));
                foundUSDT = true;
            }
        }

        assertTrue(foundUSDC);
        assertTrue(foundUSDT);
    }

    function test_getSupportedTokenList_includes_correct_tokenAddress() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        L2TokenRegistry.TokenEntry[] memory tokenList = priceOracle.getSupportedTokenList();
        assertEq(tokenList.length, 1);
        assertEq(tokenList[0].tokenID, TOKEN_ID_USDC);
        assertEq(tokenList[0].tokenAddress, address(usdc));
    }

    function test_registerToken_adds_to_supported_list() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));
        assertEq(priceOracle.getSupportedTokenCount(), 1);
    }

    function test_registerTokens_adds_all_to_supported_list() public {
        uint16[] memory tokenIDs = new uint16[](3);
        address[] memory tokenAddresses = new address[](3);
        bytes32[] memory balanceSlots = new bytes32[](3);
        uint256[] memory scales = new uint256[](3);

        tokenIDs[0] = TOKEN_ID_USDC;
        tokenIDs[1] = TOKEN_ID_USDT;
        tokenIDs[2] = TOKEN_ID_DAI;

        tokenAddresses[0] = address(usdc);
        tokenAddresses[1] = address(usdt);
        tokenAddresses[2] = address(dai);

        balanceSlots[0] = BALANCE_SLOT_USDC;
        balanceSlots[1] = BALANCE_SLOT_USDT;
        balanceSlots[2] = BALANCE_SLOT_DAI;

        bool[] memory needBalanceSlots = new bool[](3);
        needBalanceSlots[0] = true;
        needBalanceSlots[1] = true;
        needBalanceSlots[2] = true;

        scales[0] = SCALE_USDC;
        scales[1] = SCALE_USDT;
        scales[2] = SCALE_DAI;

        vm.prank(owner);
        priceOracle.registerTokens(tokenIDs, tokenAddresses, balanceSlots, needBalanceSlots, scales);

        assertEq(priceOracle.getSupportedTokenCount(), 3);
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDT));
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_DAI));
    }

    function test_removeToken_removes_from_supported_list() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);

        assertEq(priceOracle.getSupportedTokenCount(), 2);
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));

        vm.prank(owner);
        priceOracle.removeToken(TOKEN_ID_USDC);

        assertEq(priceOracle.getSupportedTokenCount(), 1);
        assertFalse(priceOracle.isTokenSupported(TOKEN_ID_USDC));
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDT));
    }

    function test_removeToken_removes_from_tokenList() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDT, address(usdt), BALANCE_SLOT_USDT, true, SCALE_USDT);

        vm.prank(owner);
        priceOracle.removeToken(TOKEN_ID_USDC);

        uint16[] memory tokenIDs = priceOracle.getSupportedIDList();
        assertEq(tokenIDs.length, 1);
        assertEq(tokenIDs[0], TOKEN_ID_USDT);

        L2TokenRegistry.TokenEntry[] memory tokenList = priceOracle.getSupportedTokenList();
        assertEq(tokenList.length, 1);
        assertEq(tokenList[0].tokenID, TOKEN_ID_USDT);
        assertEq(tokenList[0].tokenAddress, address(usdt));
    }

    function test_removeToken_cleans_up_all_mappings() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.prank(owner);
        priceOracle.updatePriceRatio(TOKEN_ID_USDC, 1e12);

        vm.prank(owner);
        priceOracle.removeToken(TOKEN_ID_USDC);

        // Token should be removed from registry
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        priceOracle.getTokenInfo(TOKEN_ID_USDC);

        // Token address mapping should be cleared
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        priceOracle.getTokenIdByAddress(address(usdc));

        // Price should be cleared
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        priceOracle.getTokenPrice(TOKEN_ID_USDC);
    }

    function test_removeToken_reverts_when_not_owner() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectRevert("Ownable: caller is not the owner");
        vm.prank(alice);
        priceOracle.removeToken(TOKEN_ID_USDC);
    }

    function test_removeToken_reverts_when_token_not_found() public {
        vm.expectRevert(bytes4(keccak256("TokenNotFound()")));
        vm.prank(owner);
        priceOracle.removeToken(TOKEN_ID_USDC);
    }

    function test_updateTokenScale_reverts_when_scale_is_zero() public {
        // First register a token
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        // Try to update scale to 0
        vm.expectRevert(bytes4(keccak256("InvalidScale()")));
        vm.prank(owner);
        priceOracle.updateTokenScale(TOKEN_ID_USDC, 0);
    }

    function test_removeToken_emits_TokenRemoved_event() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        vm.expectEmit(true, true, false, false);
        emit IL2TokenRegistry.TokenRemoved(TOKEN_ID_USDC, address(usdc));

        vm.prank(owner);
        priceOracle.removeToken(TOKEN_ID_USDC);
    }

    function test_updateTokenInfo_keeps_token_in_supported_list() public {
        vm.prank(owner);
        priceOracle.registerToken(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, SCALE_USDC);

        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));
        assertEq(priceOracle.getSupportedTokenCount(), 1);

        // Update token info
        vm.prank(owner);
        priceOracle.updateTokenInfo(TOKEN_ID_USDC, address(usdc), BALANCE_SLOT_USDC, true, true, SCALE_USDC);

        // Token should still be in supported list
        assertTrue(priceOracle.isTokenSupported(TOKEN_ID_USDC));
        assertEq(priceOracle.getSupportedTokenCount(), 1);
    }
}
