// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL1Staking} from "./IL1Staking.sol";
import {IL2Staking} from "../../L2/staking/IL2Staking.sol";

contract L1Staking is IL1Staking, OwnableUpgradeable, Staking {
    // rollup Contract
    address public ROLLUP_CONTRACT;
    // staking value
    uint256 public override STAKING_VALUE;
    // exit lock blocks
    uint256 public LOCK_BLOCKS;

    // staker whitelist
    mapping(address => bool) public whitelist;

    // all stakers, sort by staking value
    address[] public stakers;
    // all staker infos
    mapping(address => Types.StakerInfo) public stakings;
    // withdrawl unlock time
    mapping(address => uint256) public withdrawals;
    // slash unlock time
    mapping(address => bool) public slashing;
    // slash reward unlock time
    mapping(address => uint256) public slashReward;

    // latest sequencer set
    address[] public sequencersAddr;
    // latest sequencer bls keys
    bytes[] public sequencersBLS;

    /**
     * @notice staker registered
     */
    event Registered(address addr, bytes32 tmKey, bytes blsKey);

    /**
     * @notice withdrawed
     */
    event Withdrawed(address addr, uint256 unlockTime);

    /**
     * @notice staker claimed
     */
    event Claimed(address addr);

    /**
     * @notice whether in whitelist
     */
    modifier inWhitelist(address addr) {
        require(whitelist[addr], "not in whitelist");
        _;
    }

    /**
     * @notice only rollup contract
     */
    modifier onlyRollupContract() {
        require(msg.sender == ROLLUP_CONTRACT, "only rollup contract");
        _;
    }

    /**
     * @param _messenger   Address of CrossDomainMessenger on this network.
     */
    constructor(
        address payable _messenger
    ) Staking(_messenger, payable(Predeploys.L2_STAKING)) {}

    /**
     * @notice initializer
     * @param rollupContract    rollup contract address
     * @param admin             params admin
     * @param stakingValue      smallest staking value
     * @param lockBlocks        withdraw lock blocks
     */
    function initialize(
        address rollupContract,
        address admin,
        uint256 stakingValue,
        uint256 lockBlocks
    ) public initializer {
        require(rollupContract != address(0), "invalid rollup contract");
        require(stakingValue > 0, "staking limit must greater than 0");
        ROLLUP_CONTRACT = rollupContract;
        STAKING_VALUE = stakingValue;
        LOCK_BLOCKS = lockBlocks;
        _transferOwnership(admin);
    }

    /**
     * @notice register staker
     * @param addr      staker address
     * @param tmKey     tendermint pubkey
     * @param blsKey    bls pubkey
     */
    function register(
        address addr,
        bytes32 tmKey,
        bytes memory blsKey
    ) external payable inWhitelist(addr) {
        // TODO: restaking

        require(addr != address(0), "invalid address");
        require(stakings[addr].addr == address(0), "already registered");
        require(tmKey != 0, "invalid tendermint pubkey");
        require(blsKey.length == 256, "invalid bls pubkey");
        require(msg.value == STAKING_VALUE, "invalid staking value");

        // check for duplicates
        for (uint256 index = 0; index < stakers.length; index++) {
            require(
                stakings[stakers[index]].tmKey != tmKey,
                "tmKey already registered"
            );
            require(
                keccak256(stakings[stakers[index]].blsKey) != keccak256(blsKey),
                "blsKey already registered"
            );
        }

        stakings[addr] = Types.StakerInfo(addr, tmKey, blsKey);
        stakers.push(addr);

        emit Registered(addr, tmKey, blsKey);
    }

    /**
     * @notice update whitelist
     */
    function updateWhitelist(
        address[] calldata add,
        address[] calldata remove
    ) external onlyOwner {
        for (uint256 i = 0; i < add.length; i++) {
            whitelist[add[i]] = true;
        }
        for (uint256 i = 0; i < remove.length; i++) {
            whitelist[remove[i]] = false;
        }
    }

    /**
     * @notice get staker index
     */
    function getStakerIndex(
        address staker
    ) internal view returns (uint256 index) {
        for (uint256 i = 0; i < stakers.length; i++) {
            if (stakers[i] == staker) {
                return i;
            }
        }
        revert("staker not exist");
    }

    /**
     * @notice withdraw ETH
     */
    function withdrawETH() external {
        // TODO: restaking withdraw

        require(withdrawals[msg.sender] == 0, "withdrawing");
        withdrawals[msg.sender] = block.number + LOCK_BLOCKS;
        emit Withdrawed(msg.sender, withdrawals[msg.sender]);

        // TODO: remove staker
        // updateStakers();
    }

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        address[] memory sequencers
    ) external onlyRollupContract returns (uint256) {
        // TODO record slash info (unslashed staker)
        // TODO unRestaking sequencers
        // TODO remover staker
        // updateStakers();
        // _transfer(rollupContract, valueSum);
        // return valueSum;
        return 0;
    }

    /**
     * @notice claim challenger reward
     */
    function claimChallengerReward(
        address challenge
    ) external onlyRollupContract {
        // TODO claim reward
        // TODO remove stakers
    }

    /**
     * @notice claim ETH
     */
    function claimETH() external {
        // TODO reach unlock block and not in latest sequencer set

        uint256 index = getStakerIndex(msg.sender);
        for (uint256 i = index; i < stakers.length - 1; i++) {
            stakers[i] = stakers[i + 1];
        }
        stakers.pop();
        delete stakings[msg.sender];

        // TODO transfer
    }

    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}(hex"");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /**
     * @notice stakers count
     */
    function stakersNumber() public view returns (uint256) {
        return stakers.length;
    }

    /**
     * @notice update stakers
     */
    function updateStakers(
        Types.StakerInfo[] memory add,
        Types.StakerInfo[] memory remove
    ) internal {
        MESSENGER.sendMessage{value: msg.value}(
            address(OTHER_STAKING),
            0,
            abi.encodeWithSelector(
                IL2Staking.updateStakers.selector,
                add,
                remove
            ),
            0, // TODO: gasLimit
            address(0) // TODO: refundAddress
        );
    }

    /**
     * @notice verify BLS signature
     * @param signedSequencers  signed sequencers
     * @param sequencerSet      sequencer set
     * @param msgHash           bls message hash
     * @param signature         batch signature
     */
    function verifySignature(
        address[] memory signedSequencers,
        address[] memory sequencerSet,
        bytes32 msgHash,
        bytes memory signature
    ) external returns (bool) {
        // TODO verify BLS signature
        return true;
    }

    /**
     * @notice whether address is staker
     */
    function isStaker(address addr) external returns (bool) {
        // TODO
        return true;
    }
}
