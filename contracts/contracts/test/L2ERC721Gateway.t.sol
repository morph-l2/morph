// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {MockERC721} from "@rari-capital/solmate/src/test/utils/mocks/MockERC721.sol";

import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {L2ERC721Gateway} from "../l2/gateways/L2ERC721Gateway.sol";
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

        gateway.withdrawERC721(address(token), tokenId, 0);
        hevm.expectRevert("NOT_MINTED");
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - 1);

        // @todo check event
    }

    /// @dev withdraw erc721 with recipient
    function test_withdrawERC721WithGateway_succeeds(uint256 tokenId, address to) public {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        gateway.updateTokenMapping(address(token), address(token));

        gateway.withdrawERC721(address(token), to, tokenId, 0);
        hevm.expectRevert("NOT_MINTED");
        token.ownerOf(tokenId);
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - 1);

        // @todo check event
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

        gateway.batchWithdrawERC721(address(token), _tokenIds, 0);
        for (uint256 i = 0; i < count; i++) {
            hevm.expectRevert("NOT_MINTED");
            token.ownerOf(i);
        }
        assertEq(token.balanceOf(address(this)), TOKEN_COUNT - count);

        // @todo check event
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

        // then withdraw
        messenger.setXDomainMessageSender(address(counterpart));
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
