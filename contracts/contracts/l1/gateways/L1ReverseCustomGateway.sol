// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IL2ERC20Gateway} from "../../l2/gateways/IL2ERC20Gateway.sol";
import {IL1CrossDomainMessenger} from "../IL1CrossDomainMessenger.sol";
import {IL1ERC20Gateway} from "./IL1ERC20Gateway.sol";
import {IMorphERC20Upgradeable} from "../../libraries/token/IMorphERC20Upgradeable.sol";

import {GatewayBase} from "../../libraries/gateway/GatewayBase.sol";
import {L1ERC20Gateway} from "./L1ERC20Gateway.sol";

contract L1ReverseCustomGateway is L1ERC20Gateway {
    /**********
     * Events *
     **********/

    /// @notice Emitted when token mapping for ERC20 token is updated.
    /// @param l1Token The address of ERC20 token in layer 1.
    /// @param oldL2Token The address of the old corresponding ERC20 token in layer 2.
    /// @param newL2Token The address of the new corresponding ERC20 token in layer 2.
    event UpdateTokenMapping(address indexed l1Token, address indexed oldL2Token, address indexed newL2Token);

    /*************
     * Variables *
     *************/

    /// @notice Mapping from l1 token address to l2 token address for ERC20 token.
    mapping(address => address) public tokenMapping;

    /***************
     * Constructor *
     ***************/

    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize the storage of L1CustomERC20Gateway.
    /// @param _counterpart The address of L2CustomERC20Gateway in L2.
    /// @param _router The address of L1GatewayRouter.
    /// @param _messenger The address of L1CrossDomainMessenger.
    function initialize(address _counterpart, address _router, address _messenger) external initializer {
        require(_router != address(0), "zero router address");

        GatewayBase._initialize(_counterpart, _router, _messenger);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @inheritdoc IL1ERC20Gateway
    function getL2ERC20Address(address _l1Token) public view override returns (address) {
        return tokenMapping[_l1Token];
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice Update layer 1 to layer 2 token mapping.
    /// @param _l1Token The address of ERC20 token on layer 1.
    /// @param _l2Token The address of corresponding ERC20 token on layer 2.
    function updateTokenMapping(address _l1Token, address _l2Token) external onlyOwner {
        require(_l2Token != address(0), "token address cannot be 0");

        address _oldL2Token = tokenMapping[_l1Token];
        tokenMapping[_l1Token] = _l2Token;

        emit UpdateTokenMapping(_l1Token, _oldL2Token, _l2Token);
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @inheritdoc L1ERC20Gateway
    function _beforeFinalizeWithdrawERC20(
        address _l1Token,
        address _l2Token,
        address,
        address,
        uint256,
        bytes calldata
    ) internal virtual override {
        require(msg.value == 0, "nonzero msg.value");
        require(_l2Token != address(0), "token address cannot be 0");
        require(_l2Token == tokenMapping[_l1Token], "l2 token mismatch");
    }

    /// @inheritdoc L1ERC20Gateway
    function _beforeDropMessage(address, address, uint256) internal virtual override {
        require(msg.value == 0, "nonzero msg.value");
    }

    /// @inheritdoc L1ERC20Gateway
    function finalizeWithdrawERC20(
        address _l1Token,
        address _l2Token,
        address _from,
        address _to,
        uint256 _amount,
        bytes calldata _data
    ) external payable virtual override onlyCallByCounterpart nonReentrant {
        _beforeFinalizeWithdrawERC20(_l1Token, _l2Token, _from, _to, _amount, _data);

        IMorphERC20Upgradeable(_l1Token).mint(_to, _amount);

        _doCallback(_to, _data);

        emit FinalizeWithdrawERC20(_l1Token, _l2Token, _from, _to, _amount, _data);
    }

    /// @inheritdoc L1ERC20Gateway
    function _deposit(
        address _token,
        address _to,
        uint256 _amount,
        bytes memory _data,
        uint256 _gasLimit
    ) internal virtual override nonReentrant {
        address _l2Token = tokenMapping[_token];
        require(_l2Token != address(0), "no corresponding l2 token");

        // 1. Transfer token into this contract.
        address _from = _msgSender();
        if (router == _from) {
            (_from, _data) = abi.decode(_data, (address, bytes));
        }

        // 2. Burn token.
        IMorphERC20Upgradeable(_token).burn(_from, _amount);

        // 2. Generate message passed to L2ReverseCustomGateway.
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
