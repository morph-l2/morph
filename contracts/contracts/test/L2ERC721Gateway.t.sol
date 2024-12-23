// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {MockERC721} from "@rari-capital/solmate/src/test/utils/mocks/MockERC721.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {L2ERC721Gateway} from "../l2/gateways/L2ERC721Gateway.sol";
import {IL2ERC721Gateway} from "../l2/gateways/IL2ERC721Gateway.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";

contract L2ERC721GatewayTest is L2GatewayBaseTest {
    uint256 private constant TOKEN_COUNT = 100;
    uint256 private constant NOT_OWNED_TOKEN_ID = 233333;

    MockCrossDomainMessenger private messenger;
    address private counterpart;
    L2ERC721Gateway private gateway;

    MockERC721 private token;
    address private mockRecipient = address(3033);

    function setUp() public override {
        super.setUp();
        _deployERC721();
        messenger = new MockCrossDomainMessenger();

        gateway = l2ERC721Gateway;
        counterpart = gateway.counterpart();

        hevm.store(address(gateway), bytes32(erc721_messenger_slot), bytes32(abi.encode(address(messenger))));

        // transfer ownership
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        token = new MockERC721("Mock", "M");
        for (uint256 i = 0; i < TOKEN_COUNT; i++) {
            token.mint(address(this), i);
        }

        token.mint(address(mockRecipient), NOT_OWNED_TOKEN_ID);
    }

    function test_initialize_reInitialized_reverts() public {
        // Verify initialize can only be called once.
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(1));
    }
    function test_initialize_zeroAddress_reverts() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for L2ERC721Gateway.
        TransparentUpgradeableProxy l2ERC721GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2ERC721Gateway contract.
        L2ERC721Gateway l2ERC721GatewayImpl = new L2ERC721Gateway();
        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l2ERC721GatewayProxy)).upgradeToAndCall(
            address(l2ERC721GatewayImpl),
            abi.encodeCall(
                L2ERC721Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l2ERC721GatewayProxy)).upgradeToAndCall(
            address(l2ERC721GatewayImpl),
            abi.encodeCall(
                L2ERC721Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(0) // _messenger
                )
            )
        );
        hevm.stopPrank();
    }
    function test_initialize_succeeds() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for the L2ERC721Gateway.
        TransparentUpgradeableProxy l2ERC721GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2ERC721Gateway contract.
        L2ERC721Gateway l2ERC721GatewayImplTemp = new L2ERC721Gateway();
        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2ERC721GatewayProxyTemp)).upgradeToAndCall(
            address(l2ERC721GatewayImplTemp),
            abi.encodeCall(
                L2ERC721Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                )
            )
        );
        // Cast the proxy contract address to the L2ERC721Gateway contract type to call its methods.
        L2ERC721Gateway l2ERC721GatewayTemp = L2ERC721Gateway(address(l2ERC721GatewayProxyTemp));
        hevm.stopPrank();

        // Verify the counterpart and messenger are initialized successfully.
        assertEq(l2ERC721GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l2ERC721GatewayTemp.messenger(), address(l2CrossDomainMessenger));
    }

    function test_updateTokenMapping_onlyOwner_fails(address token1) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token1, token1);
        hevm.stopPrank();

        // l1 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token1, address(0));
    }

    function test_updateTokenMapping_succeeds(address token1, address token2) public {
        if (token2 == address(0)) token2 = address(1);

        // Expect UpdateTokenMapping event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit L2ERC721Gateway.UpdateTokenMapping(token1, address(0), token2);

        assertEq(gateway.tokenMapping(token1), address(0));
        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.tokenMapping(token1), token2);
    }

    /// @dev failed to withdraw erc721
    function test_withdrawERC721WithGateway_onlyHolder_fails(address to) public {
        // token not support
        hevm.expectRevert("no corresponding l1 token");
        if (to == address(0)) {
            gateway.withdrawERC721(address(token), 0, 0);
        } else {
            gateway.withdrawERC721(address(token), to, 0, 0);
        }

        // token not owned
        gateway.updateTokenMapping(address(token), address(token));
        hevm.expectRevert("token not owned");
        if (to == address(0)) {
            gateway.withdrawERC721(address(token), NOT_OWNED_TOKEN_ID, 0);
        } else {
            gateway.withdrawERC721(address(token), to, NOT_OWNED_TOKEN_ID, 0);
        }
    }

    /// @dev withdraw erc721 without recipient
    function test_withdrawERC721WithGateway_succeeds(uint256 tokenId) public {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gateway.updateTokenMapping(address(token), address(token));

        // Expect WithdrawERC721 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC721Gateway.WithdrawERC721(address(token), address(token), address(this), address(this), tokenId);

        gateway.withdrawERC721(address(token), tokenId, 0);
        hevm.expectRevert("NOT_MINTED");
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - 1);
    }

    /// @dev withdraw erc721 with recipient
    function test_withdrawERC721WithGateway_succeeds(uint256 tokenId, address to) public {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gateway.updateTokenMapping(address(token), address(token));

        // Expect WithdrawERC721 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC721Gateway.WithdrawERC721(address(token), address(token), address(this), to, tokenId);

        gateway.withdrawERC721(address(token), to, tokenId, 0);
        hevm.expectRevert("NOT_MINTED");
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - 1);
    }

    /// @dev failed to batch withdraw erc721
    function test_batchWithdrawERC721WithGateway_noneToken_fails(address to) public {
        // token not support
        hevm.expectRevert("no corresponding l1 token");
        if (to == address(0)) {
            gateway.batchWithdrawERC721(address(token), new uint256[](1), 0);
        } else {
            gateway.batchWithdrawERC721(address(token), to, new uint256[](1), 0);
        }

        // no token to withdraw
        hevm.expectRevert("no token to withdraw");
        if (to == address(0)) {
            gateway.batchWithdrawERC721(address(token), new uint256[](0), 0);
        } else {
            gateway.batchWithdrawERC721(address(token), to, new uint256[](0), 0);
        }

        // token not owned
        gateway.updateTokenMapping(address(token), address(token));
        uint256[] memory tokenIds = new uint256[](1);
        tokenIds[0] = NOT_OWNED_TOKEN_ID;
        hevm.expectRevert("token not owned");
        if (to == address(0)) {
            gateway.batchWithdrawERC721(address(token), tokenIds, 0);
        } else {
            gateway.batchWithdrawERC721(address(token), to, tokenIds, 0);
        }
    }

    /// @dev batch withdraw erc721 without recipient
    function test_batchWithdrawERC721WithGateway_succeeds(uint256 count) public {
        count = bound(count, 1, TOKEN_COUNT);
        gateway.updateTokenMapping(address(token), address(token));

        uint256[] memory _tokenIds = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i;
        }

        // Expect BatchWithdrawERC721 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC721Gateway.BatchWithdrawERC721(
            address(token),
            address(token),
            address(this),
            address(this),
            _tokenIds
        );

        gateway.batchWithdrawERC721(address(token), _tokenIds, 0);
        for (uint256 i = 0; i < count; i++) {
            hevm.expectRevert("NOT_MINTED");
            token.ownerOf(i);
        }
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - count);
    }

    /// @dev batch withdraw erc721 with recipient
    function test_batchWithdrawERC721WithGateway_succeeds(uint256 count, address to) public {
        count = bound(count, 1, TOKEN_COUNT);
        gateway.updateTokenMapping(address(token), address(token));

        uint256[] memory _tokenIds = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i;
        }

        gateway.batchWithdrawERC721(address(token), to, _tokenIds, 0);
        for (uint256 i = 0; i < count; i++) {
            hevm.expectRevert("NOT_MINTED");
            token.ownerOf(i);
        }
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - count);

        // @todo check event
    }

    /// @dev failed to finalize withdraw erc721
    function test_finalizeDepositERC721_counterErr_fails() public {
        // should revert, called by non-messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositERC721(address(0), address(0), address(0), address(0), 0);

        // should revert, called by messenger, xDomainMessageSender not set
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(L2ERC721Gateway.finalizeDepositERC721, (address(0), address(0), address(0), address(0), 0))
        );

        // should revert, called by messenger, xDomainMessageSender set wrong
        messenger.setXDomainMessageSender(address(2));
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(L2ERC721Gateway.finalizeDepositERC721, (address(0), address(0), address(0), address(0), 0))
        );
    }

    /// @dev finalize withdraw erc721
    function test_finalizeDepositERC721_succeeds(address from, address to, uint256 tokenId) public {
        hevm.assume(to != address(0));
        hevm.assume(to != address(mockRecipient));
        hevm.assume(to != address(this));
        hevm.assume(to.code.length == 0);

        tokenId = bound(tokenId, NOT_OWNED_TOKEN_ID + 1, type(uint256).max);
        gateway.updateTokenMapping(address(token), address(token));

        // finalize deposit
        messenger.setXDomainMessageSender(address(counterpart));
        // Expect the FinalizeDepositERC721 event can be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC721Gateway.FinalizeDepositERC721(address(token), address(token), from, to, tokenId);
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(L2ERC721Gateway.finalizeDepositERC721, (address(token), address(token), from, to, tokenId))
        );
        assertEq(token.balanceOf(to), 1);
        assertEq(token.ownerOf(tokenId), to);
    }

    /// @dev failed to finalize batch withdraw erc721
    function test_finalizeBatchDepositERC721_counterErr_fails() public {
        // should revert, called by non-messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeBatchDepositERC721(address(0), address(0), address(0), address(0), new uint256[](0));

        // should revert, called by messenger, xDomainMessageSender not set
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC721Gateway.finalizeBatchDepositERC721,
                (address(0), address(0), address(0), address(0), new uint256[](0))
            )
        );

        // should revert, called by messenger, xDomainMessageSender set wrong
        messenger.setXDomainMessageSender(address(2));
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC721Gateway.finalizeBatchDepositERC721,
                (address(0), address(0), address(0), address(0), new uint256[](0))
            )
        );
    }

    /// @dev finalize batch withdraw erc721
    function test_finalizeBatchDepositERC721_succeeds(address from, address to, uint256 count) public {
        if (to == address(0)) to = address(1);
        if (to == address(mockRecipient)) to = address(1);
        if (to == address(this)) to = address(1);

        gateway.updateTokenMapping(address(token), address(token));

        // deposit first
        count = bound(count, 1, TOKEN_COUNT);
        uint256[] memory _tokenIds = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i + NOT_OWNED_TOKEN_ID + 1;
        }

        messenger.setXDomainMessageSender(address(counterpart));
        // Expect the FinalizeDepositERC721 event can be emitted successfully
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC721Gateway.FinalizeBatchDepositERC721(address(token), address(token), from, to, _tokenIds);
        // Then finalize the batch deposit
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC721Gateway.finalizeBatchDepositERC721,
                (address(token), address(token), from, to, _tokenIds)
            )
        );
        assertEq(token.balanceOf(to), count);
        for (uint256 i = 0; i < count; i++) {
            assertEq(token.ownerOf(_tokenIds[i]), to);
        }
    }
}
