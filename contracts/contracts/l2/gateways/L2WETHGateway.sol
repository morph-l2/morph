// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import {SafeERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";

import {IL2ERC20Gateway, L2ERC20Gateway} from "./L2ERC20Gateway.sol";
import {IL2CrossDomainMessenger} from "../IL2CrossDomainMessenger.sol";
import {IWETH} from "../../interfaces/IWETH.sol";
import {IL1ERC20Gateway} from "../../l1/gateways/IL1ERC20Gateway.sol";
import {GatewayBase} from "../../libraries/gateway/GatewayBase.sol";

/// @title L2WETHGateway
/// @notice The `L2WETHGateway` contract is used to withdraw `WETH` token on layer 2 and
/// finalize deposit `WETH` from layer 1.
/// @dev The WETH tokens are not held in the gateway. It will first be unwrapped as Ether and
/// then the Ether will be sent to the `L2CrossDomainMessenger` contract.
/// On finalizing deposit, the Ether will be transferred from `L2CrossDomainMessenger`, then
/// wrapped as WETH and finally transfer to recipient.
contract L2WETHGateway is L2ERC20Gateway {
    using SafeERC20Upgradeable for IERC20Upgradeable;

    /*************
     * Constants *
     *************/

    /// @notice The address of L1 WETH address.
    address public immutable L1_WETH;

    /// @notice The address of L2 WETH address.
    // solhint-disable-next-line var-name-mixedcase
    address public immutable L2_WETH;

    /***************
     * Constructor *
     ***************/

    constructor(address _l2WETH, address _l1WETH) {
        _disableInitializers();

        L2_WETH = _l2WETH;
        L1_WETH = _l1WETH;
    }

    function initialize(address _counterpart, address _router, address _messenger) external initializer {
        require(_router != address(0), "zero router address");
        GatewayBase._initialize(_counterpart, _router, _messenger);
    }

    receive() external payable {
        require(_msgSender() == L2_WETH, "only WETH");
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IL2ERC20Gateway
    function getL1ERC20Address(address) external view override returns (address) {
        return L1_WETH;
    }

    /// @inheritdoc IL2ERC20Gateway
    function getL2ERC20Address(address) public view override returns (address) {
        return L2_WETH;
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc IL2ERC20Gateway
    function finalizeDepositERC20(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external payable override onlyCallByCounterpart nonReentrant {
        require(_l1Token == L1_WETH, "l1 token not WETH");
        require(_l2Token == L2_WETH, "l2 token not WETH");
        require(_amount == msg.value, "msg.value mismatch");

        IWETH(_l2Token).deposit{value: _amount}();
        IERC20Upgradeable(_l2Token).safeTransfer(_to, _amount);

        _doCallback(_to, _data);

        emit FinalizeDepositERC20(_l1Token, _l2Token, _from, _to, _amount, _data);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @inheritdoc L2ERC20Gateway
    function _withdraw(
        address _token,
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    ) internal virtual override nonReentrant {
        require(_amount > 0, "withdraw zero amount");
        require(_token == L2_WETH, "only WETH is allowed");

        // 1. Extract real sender if this call is from L1GatewayRouter.
        address _from = _msgSender();
        if (router == _from) {
            (_from, _data) = abi.decode(_data, (address, bytes));
        }

        // 2. Transfer token into this contract.
        IERC20Upgradeable(_token).safeTransferFrom(_from, address(this), _amount);
        IWETH(_token).withdraw(_amount);

        // 3. Generate message passed to L2StandardERC20Gateway.
        address _l1WETH = L1_WETH;
        bytes memory _message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (_l1WETH, _token, _from, _to, _amount, _data)
        );

        uint256 nonce = IL2CrossDomainMessenger(messenger).messageNonce();
        // 4. Send message to L1CrossDomainMessenger.
        IL2CrossDomainMessenger(messenger).sendMessage{value: _amount + msg.value}(
            counterpart,
            _amount,
            _message,
            _gasLimit
        );

        emit WithdrawERC20(_l1WETH, _token, _from, _to, _amount, _data, nonce);
    }
}
