// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {EnumerableSetUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol";
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
    using EnumerableSetUpgradeable for EnumerableSetUpgradeable.AddressSet;

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

    /// @notice cross-chain gas limit add staker
    uint256 public gasLimitAddStaker;

    /// @notice cross-chain gas limit remove stakers
    uint256 public gasLimitRemoveStakers;

    /// @notice slash remaining
    uint256 public slashRemaining;

    /// @notice staker whitelist
    mapping(address => bool) public whitelist;

    /// @notice all stakers
    EnumerableSetUpgradeable.AddressSet internal stakerSet;

    /// @notice all stakers info
    mapping(address => Types.StakerInfo) public stakers;

    /// @notice bls key map
    mapping(bytes blsPubkey => bool) public blsKeys;

    /// @notice tendermint key map
    mapping(bytes32 tmPubkey => bool) public tmKeys;

    /// @notice withdraw unlock block height
    mapping(address staker => uint256) public withdrawals;

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
        require(_msgSender() == rollupContract, "only rollup contract");
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
    /// @param _rollupContract    rollup contract address
    /// @param _stakingValue      smallest staking value
    /// @param _lockBlocks        withdraw lock blocks
    /// @param _rewardPercentage  percentage awarded to challenger
    /// @param _gasLimitAdd       cross-chain gas limit add staker
    /// @param _gasLimitRemove    cross-chain gas limit remove stakers
    function initialize(
        address _rollupContract,
        uint256 _stakingValue,
        uint256 _lockBlocks,
        uint256 _rewardPercentage,
        uint256 _gasLimitAdd,
        uint256 _gasLimitRemove
    ) public initializer {
        require(_rollupContract != address(0), "invalid rollup contract");
        require(_stakingValue > 0, "invalid staking value");
        require(_lockBlocks > 0, "invalid withdrawal lock blocks");
        require(_gasLimitAdd > 0, "invalid gas limit add staker");
        require(_gasLimitRemove > 0, "invalid gas limit remove stakers");
        require(
            _rewardPercentage > 0 && _rewardPercentage <= 100,
            "invalid challenger reward percentage"
        );

        __Ownable_init();
        __ReentrancyGuard_init();

        rollupContract = _rollupContract;
        rewardPercentage = _rewardPercentage;
        stakingValue = _stakingValue;
        withdrawalLockBlocks = _lockBlocks;
        gasLimitAddStaker = _gasLimitAdd;
        gasLimitRemoveStakers = _gasLimitRemove;

        emit GasLimitAddStakerUpdated(0, _gasLimitAdd);
        emit GasLimitRemoveStakersUpdated(0, _gasLimitRemove);
        emit RewardPercentageUpdated(0, _rewardPercentage);
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
        emit WhitelistUpdated(add, remove);
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
        stakerSet.add(_msgSender());
        blsKeys[blsKey] = true;
        tmKeys[tmKey] = true;
        emit Registered(_msgSender(), tmKey, blsKey);

        // send message to add staker on l2
        _addStaker(stakers[_msgSender()]);
    }

    /// @notice withdraw staking
    function withdraw() external {
        require(stakerSet.contains(_msgSender()), "only staker");
        require(withdrawals[_msgSender()] == 0, "withdrawing");

        withdrawals[_msgSender()] = block.number + withdrawalLockBlocks;
        stakerSet.remove(_msgSender());
        emit Withdrawn(_msgSender(), withdrawals[_msgSender()]);

        // send message to remove staker on l2
        address[] memory remove = new address[](1);
        remove[0] = _msgSender();
        emit StakersRemoved(remove);

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
                if (stakerSet.contains(sequencers[i])) {
                    valueSum += stakingValue;
                }
                stakerSet.remove(sequencers[i]);
                delete stakers[sequencers[i]];
            }
            // remove from whitelist
            delete whitelist[sequencers[i]];
        }

        uint256 reward = (valueSum * rewardPercentage) / 100;
        slashRemaining += valueSum - reward;
        _transfer(rollupContract, reward);

        emit Slashed(sequencers);
        emit StakersRemoved(sequencers);

        // send message to remove stakers on l2
        _removeStakers(sequencers);

        return reward;
    }

    /// @notice claim slash remaining
    /// @param receiver  receiver address
    function claimSlashRemaining(
        address receiver
    ) external onlyOwner nonReentrant {
        uint256 _slashRemaining = slashRemaining;
        _transfer(receiver, slashRemaining);
        slashRemaining = 0;
        emit SlashRemainingClaimed(receiver, _slashRemaining);
    }

    /// @notice update gas limit of add staker
    /// @param _gasLimitAdd       cross-chain gas limit add staker
    function updateGasLimitAddStaker(uint256 _gasLimitAdd) external onlyOwner {
        require(
            _gasLimitAdd > 0 && _gasLimitAdd != gasLimitAddStaker,
            "invalid new gas limit"
        );
        uint256 _oldGasLimitAddStaker = gasLimitAddStaker;
        gasLimitAddStaker = _gasLimitAdd;
        emit GasLimitAddStakerUpdated(_oldGasLimitAddStaker, _gasLimitAdd);
    }

    /// @notice update gas limit of remove stakers
    /// @param _gasLimitRemove    cross-chain gas limit remove stakers
    function updateGasLimitRemoveStakers(
        uint256 _gasLimitRemove
    ) external onlyOwner {
        require(
            _gasLimitRemove > 0 && _gasLimitRemove != gasLimitRemoveStakers,
            "invalid new gas limit"
        );
        uint256 _oldGasLimitRemove = gasLimitRemoveStakers;
        gasLimitRemoveStakers = _gasLimitRemove;
        emit GasLimitRemoveStakersUpdated(_oldGasLimitRemove, _gasLimitRemove);
    }

    /// @notice update reward percentage
    /// @param _rewardPercentage       percentage awarded to challenger
    function updateRewardPercentage(
        uint256 _rewardPercentage
    ) external onlyOwner {
        require(
            _rewardPercentage > 0 &&
                _rewardPercentage <= 100 &&
                _rewardPercentage != rewardPercentage,
            "invalid reward percentage"
        );
        uint256 _oldRewardPercentage = rewardPercentage;
        emit RewardPercentageUpdated(_oldRewardPercentage, _rewardPercentage);
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice claim withdrawal
    /// @param receiver  receiver address
    function claimWithdrawal(address receiver) external nonReentrant {
        require(withdrawals[_msgSender()] > 0, "withdrawal not exist");
        require(withdrawals[_msgSender()] < block.number, "withdrawal locked");

        delete stakers[_msgSender()];
        delete withdrawals[_msgSender()];
        emit Claimed(_msgSender(), receiver);

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
        address[] calldata signedSequencers,
        address[] calldata sequencerSet,
        bytes32 msgHash,
        bytes calldata signature
    ) external pure returns (bool) {
        // TODO verify BLS signature
        signedSequencers = signedSequencers;
        sequencerSet = sequencerSet;
        msgHash = msgHash;
        signature = signature;

        return true;
    }

    /// @notice staking value
    function getStakers() external view returns (address[] memory) {
        return stakerSet.values();
    }

    /// @notice whether address is staker
    /// @param addr  address to check
    function isStaker(address addr) external view returns (bool) {
        return stakerSet.contains(addr);
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

    /// @notice add staker
    /// @param add       staker to add
    function _addStaker(Types.StakerInfo memory add) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.addStaker, (add)),
            gasLimitAddStaker
        );
    }

    /// @notice remove stakers
    /// @param remove    stakers to remove
    function _removeStakers(address[] memory remove) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.removeStakers, (remove)),
            gasLimitRemoveStakers
        );
    }
}
