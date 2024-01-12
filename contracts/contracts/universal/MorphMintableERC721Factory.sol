// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

import { MorphMintableERC721 } from "./MorphMintableERC721.sol";
import { Semver } from "./Semver.sol";

/**
 * @title MorphMintableERC721Factory
 * @notice Factory contract for creating MorphMintableERC721 contracts.
 */
contract MorphMintableERC721Factory is Semver {
    /**
     * @notice Address of the ERC721 bridge on this network.
     */
    address public immutable BRIDGE;

    /**
     * @notice Chain ID for the remote network.
     */
    uint256 public immutable REMOTE_CHAIN_ID;

    /**
     * @notice Tracks addresses created by this factory.
     */
    mapping(address => bool) public isMorphMintableERC721;

    /**
     * @notice Emitted whenever a new MorphMintableERC721 contract is created.
     *
     * @param localToken  Address of the token on the this domain.
     * @param remoteToken Address of the token on the remote domain.
     * @param deployer    Address of the initiator of the deployment
     */
    event MorphMintableERC721Created(
        address indexed localToken,
        address indexed remoteToken,
        address deployer
    );

    /**
     * @custom:semver 1.2.0
     * @notice The semver MUST be bumped any time that there is a change in
     *         the MorphMintableERC721 token contract since this contract
     *         is responsible for deploying MorphMintableERC721 contracts.
     *
     * @param _bridge Address of the ERC721 bridge on this network.
     * @param _remoteChainId Chain ID for the remote network.
     */
    constructor(address _bridge, uint256 _remoteChainId) Semver(1, 2, 0) {
        BRIDGE = _bridge;
        REMOTE_CHAIN_ID = _remoteChainId;
    }

    /**
     * @notice Creates an instance of the standard ERC721.
     *
     * @param _remoteToken Address of the corresponding token on the other domain.
     * @param _name        ERC721 name.
     * @param _symbol      ERC721 symbol.
     */
    function createMorphMintableERC721(
        address _remoteToken,
        string memory _name,
        string memory _symbol
    ) external returns (address) {
        require(
            _remoteToken != address(0),
            "MorphMintableERC721Factory: L1 token address cannot be address(0)"
        );

        address localToken = address(
            new MorphMintableERC721(BRIDGE, REMOTE_CHAIN_ID, _remoteToken, _name, _symbol)
        );

        isMorphMintableERC721[localToken] = true;
        emit MorphMintableERC721Created(localToken, _remoteToken, msg.sender);

        return localToken;
    }
}
