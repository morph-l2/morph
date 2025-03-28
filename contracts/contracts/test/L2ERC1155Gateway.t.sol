// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {MockERC1155} from "@rari-capital/solmate/src/test/utils/mocks/MockERC1155.sol";
import {ERC1155TokenReceiver} from "@rari-capital/solmate/src/tokens/ERC1155.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L2GatewayBaseTest} from "./base/L2GatewayBase.t.sol";
import {L2ERC1155Gateway} from "../l2/gateways/L2ERC1155Gateway.sol";
import {IL2ERC1155Gateway} from "../l2/gateways/IL2ERC1155Gateway.sol";
import {MockCrossDomainMessenger} from "../mock/MockCrossDomainMessenger.sol";

contract L2ERC1155GatewayTest is L2GatewayBaseTest, ERC1155TokenReceiver {
    uint256 private constant TOKEN_COUNT = 100;

    MockCrossDomainMessenger private messenger;
    L2ERC1155Gateway private gateway;

    MockERC1155 private token;

    address private counterpart;
    address private mockRecipient = address(3033);

    function setUp() public override {
        super.setUp();
        _deployERC1155();
        messenger = new MockCrossDomainMessenger();

        gateway = l2ERC1155Gateway;
        counterpart = gateway.counterpart();

        hevm.store(address(gateway), bytes32(erc1155_messenger_slot), bytes32(abi.encode(address(messenger))));

        // transfer ownership
        hevm.prank(multisig);
        gateway.transferOwnership(address(this));

        token = new MockERC1155();
        for (uint256 i = 0; i < TOKEN_COUNT; i++) {
            token.mint(address(this), i, type(uint256).max, "");
        }
        token.setApprovalForAll(address(gateway), true);
    }

    function test_initilize_reInit_fails() public {
        hevm.expectRevert("Initializable: contract is already initialized");
        gateway.initialize(address(1), address(messenger));
    }

    function test_initialize_zeroAddress_reverts() public {
        hevm.startPrank(multisig);
        // Deploy a proxy contract for L2ERC1155Gateway.
        TransparentUpgradeableProxy l2ERC1155GatewayProxy = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2ERC1155Gateway contract.
        L2ERC1155Gateway l2ERC1155GatewayImpl = new L2ERC1155Gateway();
        // Expect revert due to zero counterpart address.
        hevm.expectRevert("zero counterpart address");
        ITransparentUpgradeableProxy(address(l2ERC1155GatewayProxy)).upgradeToAndCall(
            address(l2ERC1155GatewayImpl),
            abi.encodeCall(
                L2ERC1155Gateway.initialize,
                (
                    address(0), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                )
            )
        );

        // Expect revert due to zero messenger address.
        hevm.expectRevert("zero messenger address");
        ITransparentUpgradeableProxy(address(l2ERC1155GatewayProxy)).upgradeToAndCall(
            address(l2ERC1155GatewayImpl),
            abi.encodeCall(
                L2ERC1155Gateway.initialize,
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
        // Deploy a proxy contract for the L2ERC1155Gateway.
        TransparentUpgradeableProxy l2ERC1155GatewayProxyTemp = new TransparentUpgradeableProxy(
            address(emptyContract),
            address(multisig),
            new bytes(0)
        );
        // Deploy a new L2ERC1155Gateway contract.
        L2ERC1155Gateway l2ERC1155GatewayImplTemp = new L2ERC1155Gateway();
        // Initialize the proxy with the new implementation.
        ITransparentUpgradeableProxy(address(l2ERC1155GatewayProxyTemp)).upgradeToAndCall(
            address(l2ERC1155GatewayImplTemp),
            abi.encodeCall(
                L2ERC1155Gateway.initialize,
                (
                    address(NON_ZERO_ADDRESS), // _counterpart
                    address(l2CrossDomainMessenger) // _messenger
                )
            )
        );
        // Cast the proxy contract address to the L2ERC1155Gateway contract type to call its methods.
        L2ERC1155Gateway l2ERC1155GatewayTemp = L2ERC1155Gateway(address(l2ERC1155GatewayProxyTemp));
        hevm.stopPrank();

        // Verify the counterpart and messenger are initialized successfully.
        assertEq(l2ERC1155GatewayTemp.counterpart(), address(NON_ZERO_ADDRESS));
        assertEq(l2ERC1155GatewayTemp.messenger(), address(l2CrossDomainMessenger));
    }

    function test_updateTokenMapping_onlyOwner_fails(address token1) public {
        // call by non-owner, should revert
        hevm.startPrank(address(1));
        hevm.expectRevert("Ownable: caller is not the owner");
        gateway.updateTokenMapping(token1, token1);
        hevm.stopPrank();

        // l2 token is zero, should revert
        hevm.expectRevert("token address cannot be 0");
        gateway.updateTokenMapping(token1, address(0));
    }

    function test_updateTokenMapping_succeeds(address token1, address token2) public {
        if (token2 == address(0)) token2 = address(1);

        // Expect the UpdateTokenMapping event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit L2ERC1155Gateway.UpdateTokenMapping(token1, address(0), token2);

        assertEq(gateway.tokenMapping(token1), address(0));
        gateway.updateTokenMapping(token1, token2);
        assertEq(gateway.tokenMapping(token1), token2);
    }

    /// @dev failed to withdraw erc1155
    function test_withdrawERC1155WithGateway_zeroAmount_fails(address to) public {
        // token not support
        hevm.expectRevert("no corresponding l1 token");
        if (to == address(0)) {
            gateway.withdrawERC1155(address(token), 0, 1, 0);
        } else {
            gateway.withdrawERC1155(address(token), to, 0, 1, 0);
        }

        // withdraw zero amount
        hevm.expectRevert("withdraw zero amount");
        if (to == address(0)) {
            gateway.withdrawERC1155(address(token), 0, 0, 0);
        } else {
            gateway.withdrawERC1155(address(token), to, 0, 0, 0);
        }
    }

    /// @dev withdraw erc1155 without recipient
    function test_withdrawERC1155WithGateway_succeeds(uint256 tokenId, uint256 amount) public {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, type(uint256).max);
        gateway.updateTokenMapping(address(token), address(token));

        // Expect WithdrawERC1155 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.WithdrawERC1155(
            address(token),
            address(token),
            address(this),
            address(this),
            tokenId,
            amount
        );

        gateway.withdrawERC1155(address(token), tokenId, amount, 0);
        assertEq(token.balanceOf(address(gateway), tokenId), 0);
        assertEq(token.balanceOf(address(this), tokenId), type(uint256).max - amount);
    }

    /// @dev withdraw erc1155 with recipient
    function test_withdrawERC1155WithGateway_succeeds(uint256 tokenId, uint256 amount, address to) public {
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, type(uint256).max);
        gateway.updateTokenMapping(address(token), address(token));

        // Expect WithdrawERC1155 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.WithdrawERC1155(address(token), address(token), address(this), to, tokenId, amount);

        gateway.withdrawERC1155(address(token), to, tokenId, amount, 0);
        assertEq(token.balanceOf(address(gateway), tokenId), 0);
        assertEq(token.balanceOf(address(this), tokenId), type(uint256).max - amount);
    }

    /// @dev failed to batch withdraw erc1155
    function test_batchWithdrawERC1155WithGateway_zeroAmount_fails(address to) public {
        // no token to withdraw
        hevm.expectRevert("no token to withdraw");
        if (to == address(0)) {
            gateway.batchWithdrawERC1155(address(token), new uint256[](0), new uint256[](0), 0);
        } else {
            gateway.batchWithdrawERC1155(address(token), to, new uint256[](0), new uint256[](0), 0);
        }

        // length mismatch
        hevm.expectRevert("length mismatch");
        if (to == address(0)) {
            gateway.batchWithdrawERC1155(address(token), new uint256[](1), new uint256[](0), 0);
        } else {
            gateway.batchWithdrawERC1155(address(token), to, new uint256[](1), new uint256[](0), 0);
        }

        uint256[] memory amounts = new uint256[](1);
        // withdraw zero amount
        hevm.expectRevert("withdraw zero amount");
        if (to == address(0)) {
            gateway.batchWithdrawERC1155(address(token), new uint256[](1), amounts, 0);
        } else {
            gateway.batchWithdrawERC1155(address(token), to, new uint256[](1), amounts, 0);
        }

        // token not support
        amounts[0] = 1;
        hevm.expectRevert("no corresponding l1 token");
        if (to == address(0)) {
            gateway.batchWithdrawERC1155(address(token), new uint256[](1), amounts, 0);
        } else {
            gateway.batchWithdrawERC1155(address(token), to, new uint256[](1), amounts, 0);
        }
    }

    /// @dev batch withdraw erc1155 without recipient
    function test_batchWithdrawERC1155WithGateway_succeeds(uint256 count, uint256 amount) public {
        count = bound(count, 1, TOKEN_COUNT);
        amount = bound(amount, 1, type(uint256).max);
        gateway.updateTokenMapping(address(token), address(token));

        uint256[] memory _tokenIds = new uint256[](count);
        uint256[] memory _amounts = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        // Expect BatchWithdrawERC1155 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.BatchWithdrawERC1155(
            address(token),
            address(token),
            address(this),
            address(this),
            _tokenIds,
            _amounts
        );

        gateway.batchWithdrawERC1155(address(token), _tokenIds, _amounts, 0);
        for (uint256 i = 0; i < count; i++) {
            assertEq(token.balanceOf(address(gateway), i), 0);
            assertEq(token.balanceOf(address(this), i), type(uint256).max - amount);
        }
    }

    /// @dev batch withdraw erc1155 with recipient
    function test_batchWithdrawERC1155WithGateway_succeeds(uint256 count, uint256 amount, address to) public {
        count = bound(count, 1, TOKEN_COUNT);
        amount = bound(amount, 1, type(uint256).max);
        gateway.updateTokenMapping(address(token), address(token));

        uint256[] memory _tokenIds = new uint256[](count);
        uint256[] memory _amounts = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        // Expect BatchWithdrawERC1155 event to be emitted.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.BatchWithdrawERC1155(
            address(token),
            address(token),
            address(this),
            to,
            _tokenIds,
            _amounts
        );

        gateway.batchWithdrawERC1155(address(token), to, _tokenIds, _amounts, 0);
        for (uint256 i = 0; i < count; i++) {
            assertEq(token.balanceOf(address(gateway), i), 0);
            assertEq(token.balanceOf(address(this), i), type(uint256).max - _amounts[i]);
        }
    }

    /// @dev failed to finalize deposit erc1155
    function test_finalizeDepositERC1155_counterErr_fails() public {
        // should revert, called by non-messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeDepositERC1155(address(0), address(0), address(0), address(0), 0, 1);

        // should revert, called by messenger, xDomainMessageSender not set
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeDepositERC1155,
                (address(0), address(0), address(0), address(0), 0, 1)
            )
        );

        // should revert, called by messenger, xDomainMessageSender set wrong
        messenger.setXDomainMessageSender(address(2));
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeDepositERC1155,
                (address(0), address(0), address(0), address(0), 0, 1)
            )
        );
    }

    /// @dev finalize deposit erc1155
    function test_finalizeDepositERC1155_succeeds(address from, address to, uint256 tokenId, uint256 amount) public {
        hevm.assume(to != address(0));
        hevm.assume(to.code.length == 0);

        gateway.updateTokenMapping(address(token), address(token));
        tokenId = bound(tokenId, 0, TOKEN_COUNT - 1);
        amount = bound(amount, 1, type(uint256).max);

        // finalize deposit
        messenger.setXDomainMessageSender(address(counterpart));

        // Expect the FinalizeDepositERC1155 event can be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.FinalizeDepositERC1155(address(token), address(token), from, to, tokenId, amount);

        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeDepositERC1155,
                (address(token), address(token), from, to, tokenId, amount)
            )
        );
        assertEq(token.balanceOf(to, tokenId), amount);
    }

    /// @dev failed to finalize batch deposit erc1155
    function test_finalizeBatchDepositERC1155_counterErr_fails() public {
        // should revert, called by non-messenger
        hevm.expectRevert("only messenger can call");
        gateway.finalizeBatchDepositERC1155(
            address(0),
            address(0),
            address(0),
            address(0),
            new uint256[](0),
            new uint256[](0)
        );

        // should revert, called by messenger, xDomainMessageSender not set
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeBatchDepositERC1155,
                (address(0), address(0), address(0), address(0), new uint256[](0), new uint256[](0))
            )
        );

        // should revert, called by messenger, xDomainMessageSender set wrong
        messenger.setXDomainMessageSender(address(2));
        hevm.expectRevert("only call by counterpart");
        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeBatchDepositERC1155,
                (address(0), address(0), address(0), address(0), new uint256[](0), new uint256[](0))
            )
        );
    }

    /// @dev finalize batch withdraw erc1155
    function test_finalizeBatchWithdrawERC1155_succeeds(
        address from,
        address to,
        uint256 count,
        uint256 amount
    ) public {
        if (to == address(0) || to.code.length > 0) to = address(1);
        gateway.updateTokenMapping(address(token), address(token));

        count = bound(count, 1, TOKEN_COUNT);
        amount = bound(amount, 1, type(uint256).max);
        uint256[] memory _tokenIds = new uint256[](count);
        uint256[] memory _amounts = new uint256[](count);
        for (uint256 i = 0; i < count; i++) {
            _tokenIds[i] = i;
            _amounts[i] = amount;
        }

        // finalize batch deposit
        messenger.setXDomainMessageSender(address(counterpart));

        // Expect the FinalizeDepositERC1155 event can be emitted successfully.
        hevm.expectEmit(true, true, true, true);
        emit IL2ERC1155Gateway.FinalizeBatchDepositERC1155(
            address(token),
            address(token),
            from,
            to,
            _tokenIds,
            _amounts
        );

        messenger.callTarget(
            address(gateway),
            abi.encodeCall(
                L2ERC1155Gateway.finalizeBatchDepositERC1155,
                (address(token), address(token), from, to, _tokenIds, _amounts)
            )
        );
        for (uint256 i = 0; i < count; i++) {
            assertEq(token.balanceOf(to, i), _amounts[i]);
        }
    }
}
