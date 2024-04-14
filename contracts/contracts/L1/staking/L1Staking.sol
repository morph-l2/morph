// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL1Staking} from "./IL1Staking.sol";
import {IL2Staking} from "../../L2/staking/IL2Staking.sol";

contract L1Staking is
    IL1Staking,
    Staking,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    // rollup Contract
    address public ROLLUP_CONTRACT;
    // staking value, immutable
    uint256 public override STAKING_VALUE;
    // exit lock blocks
    uint256 public WITHDRAWAL_LOCK_BLOCKS;
    // percentage awarded to challenger
    uint256 public REWARD_PERCENTAGE;
    // default crosschain gas limit
    uint256 public DEFAULT_GAS_LIMIT;

    // slash remaining
    uint256 slashRemaining;
    // staker whitelist
    mapping(address => bool) public whitelist;
    // all stakers
    address[] public stakerList;
    // all stakers info
    mapping(address => Types.StakerInfo) public stakers;
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
    event Claimed(address staker, address receiver);

    /**
     * @notice stakers were slashed
     */
    event Slashed(address[] stakers);

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
     * @param admin             params admin
     * @param rollupContract    rollup contract address
     * @param rewardPercentage  percentage awarded to challenger
     * @param stakingValue      smallest staking value
     * @param lockBlocks        withdraw lock blocks
     * @param gasLimit          default crosschain gas limit
     */
    function initialize(
        address admin,
        address rollupContract,
        uint256 rewardPercentage,
        uint256 stakingValue,
        uint256 lockBlocks,
        uint256 gasLimit
    ) public initializer {
        require(admin != address(0), "invalid admin");
        require(rollupContract != address(0), "invalid rollup contract");
        require(stakingValue > 0, "staking limit must greater than 0");
        require(lockBlocks > 0, "staking limit must greater than 0"); // TODO TBD
        require(gasLimit > 0, "gas limit must greater than 0");
        require(
            rewardPercentage > 0 && rewardPercentage <= 100,
            "invalid reward percentage"
        );
        ROLLUP_CONTRACT = rollupContract;
        REWARD_PERCENTAGE = rewardPercentage;
        STAKING_VALUE = stakingValue;
        WITHDRAWAL_LOCK_BLOCKS = lockBlocks;
        DEFAULT_GAS_LIMIT = gasLimit;
        _transferOwnership(admin);
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
        require(addr != address(0), "invalid address");
        require(stakers[addr].addr == address(0), "already registered");
        require(tmKey != 0, "invalid tendermint pubkey");
        require(blsKey.length == 256, "invalid bls pubkey");
        require(msg.value == STAKING_VALUE, "invalid staking value");

        // check for duplicates
        for (uint256 index = 0; index < stakerList.length; index++) {
            require(
                stakers[stakerList[index]].tmKey != tmKey,
                "tmKey already registered"
            );
            require(
                keccak256(stakers[stakerList[index]].blsKey) !=
                    keccak256(blsKey),
                "blsKey already registered"
            );
        }

        stakers[addr] = Types.StakerInfo(addr, tmKey, blsKey);
        stakerList.push(addr);
        emit Registered(addr, tmKey, blsKey);

        // send message to add staker on l2
        updateStakers(stakers[addr], new address[](0));
    }

    /**
     * @notice withdraw staking
     */
    function withdraw() external {
        (bool exist, uint256 index) = getStakerIndex(msg.sender);
        require(exist, "only staker");
        require(withdrawals[msg.sender] == 0, "withdrawing");

        withdrawals[msg.sender] = block.number + WITHDRAWAL_LOCK_BLOCKS;
        stakerList[index] = stakerList[stakerList.length - 1];
        stakerList.pop();
        emit Withdrawed(msg.sender, withdrawals[msg.sender]);

        // send message to remove staker on l2
        address[] memory remove = new address[](1);
        remove[0] = msg.sender;
        updateStakers(
            Types.StakerInfo(address(0), bytes32(0), bytes("")),
            remove
        );
    }

    /**
     * @notice claim withdrawal
     * @param receiver  receiver address
     */
    function claimWithdrawal(address receiver) external nonReentrant {
        require(withdrawals[msg.sender] > 0, "withdrawal not exist");
        require(withdrawals[msg.sender] < block.number, "withdrawal locked");

        delete stakers[msg.sender];
        delete withdrawals[msg.sender];
        emit Claimed(msg.sender, receiver);

        _transfer(receiver, STAKING_VALUE);
    }

    /**
     * @notice transfer ETH
     * @param _to      The address to transfer ETH to.
     * @param _amount  The amount of ETH to transfer.
     */
    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}(hex"");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        address[] memory sequencers
    ) external onlyRollupContract nonReentrant returns (uint256) {
        uint256 valueSum;
        for (uint256 i = 0; i < sequencers.length; i++) {
            if (withdrawals[sequencers[i]] > 0) {
                delete withdrawals[sequencers[i]];
                delete stakers[sequencers[i]];
                valueSum += STAKING_VALUE;
            } else {
                (bool exist, uint256 index) = getStakerIndex(sequencers[i]);
                if (exist) {
                    stakerList[index] = stakerList[stakerList.length - 1];
                    stakerList.pop();
                    delete stakers[sequencers[i]];
                    valueSum += STAKING_VALUE;
                }
                // TBD: handle the case "exist == false"
            }
            // remove from whitelist
            delete whitelist[sequencers[i]];
        }

        uint256 reward = (valueSum * REWARD_PERCENTAGE) / 100;
        slashRemaining += valueSum - reward;
        _transfer(ROLLUP_CONTRACT, reward);

        emit Slashed(sequencers);

        // send message to remove stakers on l2
        updateStakers(
            Types.StakerInfo(address(0), bytes32(0), bytes("")),
            sequencers
        );

        return reward;
    }

    /**
     * @notice get staker index
     * @param staker    staker address
     */
    function getStakerIndex(
        address staker
    ) internal view returns (bool exist, uint256 index) {
        for (uint256 i = 0; i < stakerList.length; i++) {
            if (stakerList[i] == staker) {
                return (true, i);
            }
        }
        return (false, 0);
    }

    /**
     * @notice update stakers
     * @param add       staker to add
     * @param remove    stakers to remove
     */
    function updateStakers(
        Types.StakerInfo memory add,
        address[] memory remove
    ) internal {
        MESSENGER.sendMessage{value: msg.value}(
            address(OTHER_STAKING),
            0,
            abi.encodeWithSelector(
                IL2Staking.updateStakers.selector,
                add,
                remove
            ),
            DEFAULT_GAS_LIMIT
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
    ) external pure returns (bool) {
        // TODO verify BLS signature
        signedSequencers = signedSequencers;
        sequencerSet = sequencerSet;
        msgHash = msgHash;
        signature = signature;

        return true;
    }

    /**
     * @notice whether address is staker
     * @param addr  address to check
     */
    function isStaker(address addr) external view returns (bool) {
        for (uint256 i = 0; i < stakerList.length; i++) {
            if (addr == stakerList[i]) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice claim slash remaining
     * @param receiver  receiver address
     */
    function claimSlashRemaining(
        address receiver
    ) external onlyOwner nonReentrant {
        _transfer(receiver, slashRemaining);
        slashRemaining = 0;
    }
}
