// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
 * @title TestERC20
 */
contract TestERC20 is ERC20 {
    /**
     * @notice Emitted when the token is minted by the bridge.
     */
    event Mint(address indexed _account, uint256 _amount);

    /**
     * @param _name ERC20 name.
     * @param _symbol ERC20 symbol.
     */
    constructor(
        string memory _name,
        string memory _symbol
    ) ERC20(_name, _symbol) {
    }

    /**
     * @notice Only the bridge can mint tokens.
     * @param _to     The account receiving tokens.
     * @param _amount The amount of tokens to receive.
     */
    function mint(address _to, uint256 _amount) public {
        _mint(_to, _amount);
        emit Mint(_to, _amount);
    }
}
