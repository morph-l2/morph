// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {L2ERC20Gateway} from "../l2/gateways/L2ERC20Gateway.sol";
import {IL2CrossDomainMessenger} from "../l2/IL2CrossDomainMessenger.sol";
import {IMorphERC20Upgradeable} from "../libraries/token/IMorphERC20Upgradeable.sol";
import {GatewayBase} from "../libraries/gateway/GatewayBase.sol";

import {LidoBridgeableTokens} from "./LidoBridgeableTokens.sol";
import {LidoGatewayManager} from "./LidoGatewayManager.sol";

/**
 * @custom:security-contact official@morphl2.io
 */
contract L2LidoGateway is L2ERC20Gateway, LidoBridgeableTokens, LidoGatewayManager {
    /**********
     * Errors *
     **********/

    /// @dev Thrown when withdraw zero amount token.
    error ErrorWithdrawZeroAmount();

    /// @dev Thrown when withdraw erc20 with calldata.
    error WithdrawAndCallIsNotAllowed();

    /*************
     * Variables *
     *************/

    /// @dev The initial version of `L2LidoGateway` use `L2CustomERC20Gateway`. We keep the storage
    /// slot for `tokenMapping` for compatibility. It should no longer be used.
    mapping(address => address) private __tokenMapping;

    /***************
     * Constructor *
     ***************/

    /// @notice Constructor for `L2LidoGateway` implementation contract.
    ///
    /// @param _l1Token The address of the bridged token in the L1 chain
    /// @param _l2Token The address of the token minted on the L2 chain when token bridged
    constructor(address _l1Token, address _l2Token) LidoBridgeableTokens(_l1Token, _l2Token) {
        if (_l1Token == address(0) || _l2Token ==address(0)){
              revert ErrorZeroAddress();
        }

        _disableInitializers();
    }

    /// @notice Initialize the storage of L2LidoGateway v1.
    ///
    /// @dev The parameters `_counterpart`, `_router` and `_messenger` are no longer used.
    ///
    /// @param _counterpart The address of `L1LidoGateway` contract in L1.
    /// @param _router The address of `L2GatewayRouter` contract in L2.
    /// @param _messenger The address of `L2CrossDomainMessenger` contract in L2.
    function initialize(address _counterpart, address _router, address _messenger) external initializer {
        GatewayBase._initialize(_counterpart, _router, _messenger);
    }

    /// @notice Initialize the storage of L2LidoGateway v2.
    /// @param _depositsEnabler The address of user who can enable deposits
    /// @param _depositsDisabler The address of user who can disable deposits
    /// @param _withdrawalsEnabler The address of user who can enable withdrawals
    /// @param _withdrawalsDisabler The address of user who can disable withdrawals
    function initializeV2(
        address _depositsEnabler,
        address _depositsDisabler,
        address _withdrawalsEnabler,
        address _withdrawalsDisabler
    ) external reinitializer(2) {
        __LidoGatewayManager_init(_depositsEnabler, _depositsDisabler, _withdrawalsEnabler, _withdrawalsDisabler);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IL2ERC20Gateway
    function getL1ERC20Address(
        address _l2Token
    ) external view override onlySupportedL2Token(_l2Token) returns (address) {
        return l1Token;
    }

    /// @inheritdoc IL2ERC20Gateway
    function getL2ERC20Address(
        address _l1Token
    ) external view override onlySupportedL1Token(_l1Token) returns (address) {
        return l2Token;
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @inheritdoc IL2ERC20Gateway
    /// @dev The length of `_data` always be zero, which guarantee by `L1LidoGateway`.
    function finalizeDepositERC20(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    )
        external
        payable
        override
        onlyCallByCounterpart
        nonReentrant
        onlySupportedL1Token(_l1Token)
        onlySupportedL2Token(_l2Token)
        whenDepositsEnabled
    {
        if (msg.value != 0) revert ErrorNonZeroMsgValue();

        IMorphERC20Upgradeable(_l2Token).mint(_to, _amount);

        emit FinalizeDepositERC20(_l1Token, _l2Token, _from, _to, _amount, _data);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @inheritdoc L2ERC20Gateway
    function _withdraw(
        address _l2Token,
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    )
        internal
        virtual
        override
        nonReentrant
        onlySupportedL2Token(_l2Token)
        onlyNonZeroAccount(_to)
        whenWithdrawalsEnabled
    {
        if (_amount == 0) revert ErrorWithdrawZeroAmount();

        // 1. Extract real sender if this call is from L2GatewayRouter.
        address _from = _msgSender();
        if (router == _from) {
            (_from, _data) = abi.decode(_data, (address, bytes));
        }
        if (_data.length != 0) revert WithdrawAndCallIsNotAllowed();

        // 2. Burn token.
        IMorphERC20Upgradeable(_l2Token).burn(_from, _amount);

        // 3. Generate message passed to L1LidoGateway.
        address _l1Token = l1Token;
        bytes memory _message = abi.encodeCall(
            IL1ERC20Gateway.finalizeWithdrawERC20,
            (_l1Token, _l2Token, _from, _to, _amount, _data)
        );

        uint256 nonce = IL2CrossDomainMessenger(messenger).messageNonce();

        // 4. send message to L2CrossDomainMessenger
        IL2CrossDomainMessenger(messenger).sendMessage{value: msg.value}(counterpart, 0, _message, _gasLimit);

        emit WithdrawERC20(_l1Token, _l2Token, _from, _to, _amount, _data, nonce);
    }
}
