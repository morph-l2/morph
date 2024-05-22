// SPDX-License-Identifier: MIT

pragma solidity =0.8.24;

import {IERC1155} from "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

import {IMorphERC1155Extension} from "./IMorphERC1155Extension.sol";

// The recommended ERC1155 implementation for bridge token.
// deployed in L2 when original token is on L1
// deployed in L1 when original token is on L2
interface IMorphERC1155 is IERC1155, IMorphERC1155Extension {}
