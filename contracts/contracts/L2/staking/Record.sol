// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {ISequencer} from "./ISequencer.sol";
import {IGov} from "./IGov.sol";
import {IRecord} from "./IRecord.sol";

contract Record is IRecord, Initializable {
    // sequencer contract address
    address public immutable SEQUENCER_CONTRACT;
    // gov contract address
    address public immutable GOV_CONTRACT;
    // oracle address
    address public ORACLE;

    /**
     * @notice constructor
     */
    constructor() {
        SEQUENCER_CONTRACT = Predeploys.SEQUENCER;
        GOV_CONTRACT = Predeploys.GOV;
    }

    /**
     * @notice Initializer.
     * @param _oracle oracle address
     */
    function initialize(address _oracle) public initializer {
        require(_oracle != address(0), "invalid oracle address");
        ORACLE = _oracle;
    }
}
