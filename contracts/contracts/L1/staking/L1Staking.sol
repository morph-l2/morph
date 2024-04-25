// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

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
    /*************
     * Variables *
     *************/

    /// @notice rollup Contract
    address public rollupContract;

    /// @notice staking value, immutable
    uint256 public stakingValue;

    /// @notice exit lock blocks
    uint256 public withdrawalLockBlocks;

    /// @notice percentage awarded to challenger
    uint256 public rewardPercentage;

    /// @notice default cross-chain gas limit
    uint256 public defaultGasLimit;

    /// @notice slash remaining
    uint256 public slashRemaining;

    /// @notice staker whitelist
    mapping(address => bool) public whitelist;

    /// @notice all stakers
    address[] public stakerList;

    /// @notice all stakers info
    mapping(address => Types.StakerInfo) public stakers;

    /// @notice bls key map
    mapping(bytes => bool) public blsKeys;

    /// @notice tendermint key map
    mapping(bytes32 => bool) public tmKeys;

    /// @notice withdraw unlock block height
    mapping(address => uint256) public withdrawals;

    /**********************
     * Function Modifiers *
     **********************/

    /// @notice whether in whitelist
    modifier inWhitelist(address addr) {
        require(whitelist[addr], "not in whitelist");
        _;
    }

    /// @notice only rollup contract
    modifier onlyRollupContract() {
        require(msg.sender == rollupContract, "only rollup contract");
        _;
    }

    /***************
     * Constructor *
     ***************/

    /// @param _messenger   Address of CrossDomainMessenger on this network.
    constructor(
        address payable _messenger
    ) Staking(_messenger, payable(Predeploys.L2_STAKING)) {}

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    /// @param _admin             params admin
    /// @param _rollupContract    rollup contract address
    /// @param _rewardPercentage  percentage awarded to challenger
    /// @param _stakingValue      smallest staking value
    /// @param _lockBlocks        withdraw lock blocks
    /// @param _gasLimit          default cross-chain gas limit
    function initialize(
        address _admin,
        address _rollupContract,
        uint256 _rewardPercentage,
        uint256 _stakingValue,
        uint256 _lockBlocks,
        uint256 _gasLimit
    ) public initializer {
        require(_admin != address(0), "invalid admin");
        require(_rollupContract != address(0), "invalid rollup contract");
        require(_stakingValue > 0, "staking limit must greater than 0");
        require(_lockBlocks > 0, "staking limit must greater than 0");
        require(_gasLimit > 0, "gas limit must greater than 0");
        require(
            _rewardPercentage > 0 && _rewardPercentage <= 100,
            "invalid reward percentage"
        );

        __Ownable_init();
        __ReentrancyGuard_init();

        rollupContract = _rollupContract;
        rewardPercentage = _rewardPercentage;
        stakingValue = _stakingValue;
        withdrawalLockBlocks = _lockBlocks;
        defaultGasLimit = _gasLimit;

        // TODO event
    }

    /************************
     * Restricted Functions *
     ************************/

    /// @notice update whitelist
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

    /// @notice register staker
    /// @param tmKey     tendermint pubkey
    /// @param blsKey    bls pubkey
    function register(
        bytes32 tmKey,
        bytes memory blsKey
    ) external payable inWhitelist(_msgSender()) {
        require(stakers[_msgSender()].addr == address(0), "already registered");
        require(tmKey != 0 && !tmKeys[tmKey], "invalid tendermint pubkey");
        require(blsKey.length == 256 && !blsKeys[blsKey], "invalid bls pubkey");
        require(msg.value == stakingValue, "invalid staking value");

        stakers[_msgSender()] = Types.StakerInfo(_msgSender(), tmKey, blsKey);
        stakerList.push(_msgSender());
        blsKeys[blsKey] = true;
        tmKeys[tmKey] = true;
        emit Registered(_msgSender(), tmKey, blsKey);

        // send message to add staker on l2
        _addStaker(stakers[_msgSender()]);
    }

    /// @notice withdraw staking
    function withdraw() external {
        (bool exist, uint256 index) = _getStakerIndex(msg.sender);
        require(exist, "only staker");
        require(withdrawals[msg.sender] == 0, "withdrawing");

        withdrawals[msg.sender] = block.number + withdrawalLockBlocks;
        stakerList[index] = stakerList[stakerList.length - 1];
        stakerList.pop();
        emit Withdrawn(msg.sender, withdrawals[msg.sender]);

        // send message to remove staker on l2
        address[] memory remove = new address[](1);
        remove[0] = msg.sender;
        _removeStakers(remove);
    }

    /// @notice challenger win, slash sequencers
    function slash(
        address[] memory sequencers
    ) external onlyRollupContract nonReentrant returns (uint256) {
        uint256 valueSum;
        for (uint256 i = 0; i < sequencers.length; i++) {
            if (withdrawals[sequencers[i]] > 0) {
                delete withdrawals[sequencers[i]];
                delete stakers[sequencers[i]];
                valueSum += stakingValue;
            } else {
                (bool exist, uint256 index) = _getStakerIndex(sequencers[i]);
                if (exist) {
                    stakerList[index] = stakerList[stakerList.length - 1];
                    stakerList.pop();
                    delete stakers[sequencers[i]];
                    valueSum += stakingValue;
                }
                // TBD: handle the case "exist == false"
            }
            // remove from whitelist
            delete whitelist[sequencers[i]];
        }

        uint256 reward = (valueSum * rewardPercentage) / 100;
        slashRemaining += valueSum - reward;
        _transfer(rollupContract, reward);

        emit Slashed(sequencers);

        // send message to remove stakers on l2
        _removeStakers(sequencers);

        return reward;
    }

    /// @notice claim slash remaining
    /// @param receiver  receiver address
    function claimSlashRemaining(
        address receiver
    ) external onlyOwner nonReentrant {
        _transfer(receiver, slashRemaining);
        slashRemaining = 0;
    }

    /// @notice claim slash remaining
    /// @param gasLimit  gas limit
    function updateParams(uint256 gasLimit) external onlyOwner {
        require(gasLimit > 0, "gas limit must greater than 0");
        defaultGasLimit = gasLimit;
        emit ParamsUpdated(gasLimit);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice claim withdrawal
    /// @param receiver  receiver address
    function claimWithdrawal(address receiver) external nonReentrant {
        require(withdrawals[msg.sender] > 0, "withdrawal not exist");
        require(withdrawals[msg.sender] < block.number, "withdrawal locked");

        delete stakers[msg.sender];
        delete withdrawals[msg.sender];
        emit Claimed(msg.sender, receiver);

        _transfer(receiver, stakingValue);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice verify BLS signature
    /// @param signedSequencers  signed sequencers
    /// @param sequencerSet      sequencer set
    /// @param msgHash           bls message hash
    /// @param signature         batch signature
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

    /// @notice whether address is staker
    /// @param addr  address to check
    function isStaker(address addr) external view returns (bool) {
        for (uint256 i = 0; i < stakerList.length; i++) {
            if (addr == stakerList[i]) {
                return true;
            }
        }
        return false;
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice transfer ETH
    /// @param _to      The address to transfer ETH to.
    /// @param _amount  The amount of ETH to transfer.
    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}("0x");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /// @notice get staker index
    /// @param staker    staker address
    function _getStakerIndex(
        address staker
    ) internal view returns (bool exist, uint256 index) {
        for (uint256 i = 0; i < stakerList.length; i++) {
            if (stakerList[i] == staker) {
                return (true, i);
            }
        }
        return (false, 0);
    }

    /// @notice add staker
    /// @param add       staker to add
    function _addStaker(Types.StakerInfo memory add) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.addStaker, (add)),
            defaultGasLimit
        );
    }

    /// @notice remove stakers
    /// @param remove    stakers to remove
    function _removeStakers(address[] memory remove) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.removeStakers, (remove)),
            defaultGasLimit
        );
    }
}
