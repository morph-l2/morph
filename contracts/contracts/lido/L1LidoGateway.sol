// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL1ERC20Gateway} from "../l1/gateways/IL1ERC20Gateway.sol";
import {L1ERC20Gateway} from "../l1/gateways/L1ERC20Gateway.sol";
import {IL1CrossDomainMessenger} from "../l1/IL1CrossDomainMessenger.sol";
import {IL2ERC20Gateway} from "../l2/gateways/IL2ERC20Gateway.sol";
import {GatewayBase} from "../libraries/gateway/GatewayBase.sol";

import {LidoBridgeableTokens} from "./LidoBridgeableTokens.sol";
import {LidoGatewayManager} from "./LidoGatewayManager.sol";

/**
 * @custom:security-contact official@morphl2.io
 */
contract L1LidoGateway is L1ERC20Gateway, LidoBridgeableTokens, LidoGatewayManager {
    /**********
     * Errors *
     **********/

    /// @dev Thrown when deposit zero amount token.
    error ErrorDepositZeroAmount();

    /// @dev Thrown when deposit erc20 with calldata.
    error DepositAndCallIsNotAllowed();

    /*************
     * Variables *
     *************/

    /// @dev The initial version of `L1LidoGateway` use `L1CustomERC20Gateway`. We keep the storage
    /// slot for `tokenMapping` for compatibility. It should no longer be used.
    mapping(address => address) private __tokenMapping;

    /***************
     * Constructor *
     ***************/

    /// @notice Constructor for `L1LidoGateway` implementation contract.
    ///
    /// @param _l1Token The address of the bridged token in the L1 chain
    /// @param _l2Token The address of the token minted on the L2 chain when token bridged
    constructor(address _l1Token, address _l2Token) LidoBridgeableTokens(_l1Token, _l2Token) {
        if (_l1Token == address(0) || _l2Token ==address(0)){
              revert ErrorZeroAddress();
        }

        _disableInitializers();
    }

    /// @notice Initialize the storage of L1LidoGateway v1.
    ///
    /// @param _counterpart The address of `L2LidoGateway` contract in L2.
    /// @param _router The address of `L1GatewayRouter` contract.
    /// @param _messenger The address of `L1CrossDomainMessenger` contract.
    function initialize(address _counterpart, address _router, address _messenger) external initializer {
        GatewayBase._initialize(_counterpart, _router, _messenger);
    }

    /// @notice Initialize the storage of L1LidoGateway v2.
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

    /// @inheritdoc IL1ERC20Gateway
    function getL2ERC20Address(
        address _l1Token
    ) external view override onlySupportedL1Token(_l1Token) returns (address) {
        return l2Token;
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @inheritdoc L1ERC20Gateway
    /// @dev The length of `_data` always be zero, which guarantee by `L2LidoGateway`.
    function _beforeFinalizeWithdrawERC20(
        address _l1Token,
        address _l2Token,
        address,
        address,
        uint256,
        bytes calldata
    ) internal virtual override onlySupportedL1Token(_l1Token) onlySupportedL2Token(_l2Token) whenWithdrawalsEnabled {
        if (msg.value != 0) revert ErrorNonZeroMsgValue();
    }

    /// @inheritdoc L1ERC20Gateway
    function _beforeDropMessage(
        address _token,
        address,
        uint256
    ) internal virtual override onlySupportedL1Token(_token) {
        if (msg.value != 0) revert ErrorNonZeroMsgValue();
    }

    /// @inheritdoc L1ERC20Gateway
    function _deposit(
        address _token,
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    ) internal virtual override nonReentrant onlySupportedL1Token(_token) onlyNonZeroAccount(_to) whenDepositsEnabled {
        if (_amount == 0) revert ErrorDepositZeroAmount();

        // 1. Transfer token into this contract.
        address _from;
        (_from, _amount, _data) = _transferERC20In(_token, _amount, _data);
        if (_data.length != 0) revert DepositAndCallIsNotAllowed();

        // 2. Generate message passed to L2LidoGateway.
        address _l2Token = l2Token;
        bytes memory _message = abi.encodeCall(
            IL2ERC20Gateway.finalizeDepositERC20,
            (_token, _l2Token, _from, _to, _amount, _data)
        );

        uint256 nonce = IL1CrossDomainMessenger(messenger).messageNonce();

        // 3. Send message to L1CrossDomainMessenger.
        IL1CrossDomainMessenger(messenger).sendMessage{value: msg.value}(counterpart, 0, _message, _gasLimit, _from);

        emit DepositERC20(_token, _l2Token, _from, _to, _amount, _data, nonce);
    }
}
