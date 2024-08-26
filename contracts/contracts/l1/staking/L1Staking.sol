// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {Types} from "../../libraries/common/Types.sol";
import {Staking} from "../../libraries/staking/Staking.sol";
import {IL1Staking} from "./IL1Staking.sol";
import {IL2Staking} from "../../l2/staking/IL2Staking.sol";

contract L1Staking is IL1Staking, Staking, OwnableUpgradeable, ReentrancyGuardUpgradeable {
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
    mapping(address stakerAddr => bool inWhitelist) public whitelist;

    /// @notice staker removed list
    mapping(address stakerAddr => bool removed) public removedList;

    /// @notice all stakers (0-254)
    address[255] public stakerSet;

    /// @notice all stakers indexes (1-255). '0' means not exist. stakerIndexes[1] releated to stakerSet[0]
    mapping(address stakerAddr => uint8 index) public stakerIndexes;

    /// @notice stakers to delete
    address[] public deleteList;

    /// @notice staker deleteable height
    mapping(address stakerAddr => uint256 height) public deleteableHeight;

    /// @notice all stakers info
    mapping(address stakerAddr => Types.StakerInfo) public stakers;

    /// @notice bls key map
    mapping(bytes blsPubkey => bool exist) public blsKeys;

    /// @notice tendermint key map
    mapping(bytes32 tmPubkey => bool exist) public tmKeys;

    /// @notice withdraw unlock block height
    mapping(address staker => uint256 amount) public withdrawals;

    /// @notice challenge deposit value
    uint256 public challengeDeposit;

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
    constructor(address payable _messenger) Staking(_messenger, payable(Predeploys.L2_STAKING)) {
        _disableInitializers();
    }

    /***************
     * Initializer *
     ***************/

    /// @notice initializer
    /// @param _rollupContract    rollup contract address
    /// @param _stakingValue      staking value
    /// @param _challengeDeposit  challenge deposit value
    /// @param _lockBlocks        withdraw lock blocks
    /// @param _rewardPercentage  percentage awarded to challenger
    /// @param _gasLimitAdd       cross-chain gas limit add staker
    /// @param _gasLimitRemove    cross-chain gas limit remove stakers
    function initialize(
        address _rollupContract,
        uint256 _stakingValue,
        uint256 _challengeDeposit,
        uint256 _lockBlocks,
        uint256 _rewardPercentage,
        uint256 _gasLimitAdd,
        uint256 _gasLimitRemove
    ) public initializer {
        require(_rollupContract != address(0), "invalid rollup contract");
        require(_stakingValue > 0, "invalid staking value");
        require(_challengeDeposit > 0, "invalid challenge deposit value");
        require(_lockBlocks > 0, "invalid withdrawal lock blocks");
        require(_gasLimitAdd > 0, "invalid gas limit add staker");
        require(_gasLimitRemove > 0, "invalid gas limit remove stakers");
        require(_rewardPercentage > 0 && _rewardPercentage <= 100, "invalid challenger reward percentage");

        __Ownable_init();
        __ReentrancyGuard_init();

        rollupContract = _rollupContract;
        rewardPercentage = _rewardPercentage;
        stakingValue = _stakingValue;
        challengeDeposit = _challengeDeposit;
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
    function updateWhitelist(address[] calldata add, address[] calldata remove) external onlyOwner {
        for (uint256 i = 0; i < add.length; i++) {
            require(!removedList[add[i]], "in removed list");
            whitelist[add[i]] = true;
        }
        for (uint256 i = 0; i < remove.length; i++) {
            delete whitelist[remove[i]];
        }
        emit WhitelistUpdated(add, remove);
    }

    /// @notice register staker
    /// @param tmKey     tendermint pubkey
    /// @param blsKey    bls pubkey
    function register(bytes32 tmKey, bytes memory blsKey) external payable inWhitelist(_msgSender()) {
        require(stakers[_msgSender()].addr == address(0), "already registered");
        require(tmKey != 0 && !tmKeys[tmKey], "invalid tendermint pubkey");
        require(blsKey.length == 256 && !blsKeys[blsKey], "invalid bls pubkey");
        require(msg.value == stakingValue, "invalid staking value");

        stakers[_msgSender()] = Types.StakerInfo(_msgSender(), tmKey, blsKey);
        _addStaker(_msgSender());
        blsKeys[blsKey] = true;
        tmKeys[tmKey] = true;
        emit Registered(_msgSender(), tmKey, blsKey);

        // send message to add staker on l2
        _msgAddStaker(stakers[_msgSender()]);
    }

    /// @notice remove staker
    function removeStaker(address[] memory _stakers) external onlyOwner {
        for (uint256 i = 0; i < _stakers.length; i++) {
            require(isActiveStaker(_stakers[i]), "only active staker can be removed");
            require(withdrawals[_stakers[i]] == 0, "withdrawing");

            withdrawals[_stakers[i]] = block.number + withdrawalLockBlocks;
            _removeStaker(_stakers[i]);
            emit Withdrawn(_stakers[i], withdrawals[_stakers[i]]);

            delete whitelist[_stakers[i]];
            removedList[_stakers[i]] = true;
        }
        emit StakersRemoved(_stakers);

        // send message to remove stakers on l2
        _msgRemoveStakers(_stakers);
    }

    /// @notice withdraw staking
    function withdraw() external {
        require(isActiveStaker(_msgSender()), "only active staker");
        require(withdrawals[_msgSender()] == 0, "withdrawing");

        withdrawals[_msgSender()] = block.number + withdrawalLockBlocks;
        _removeStaker(_msgSender());
        emit Withdrawn(_msgSender(), withdrawals[_msgSender()]);

        delete whitelist[_msgSender()];
        removedList[_msgSender()] = true;

        address[] memory remove = new address[](1);
        remove[0] = _msgSender();
        emit StakersRemoved(remove);

        // send message to remove staker on l2
        _msgRemoveStakers(remove);
    }

    /// @notice challenger win, slash sequencers
    function slash(uint256 sequencersBitmap) external onlyRollupContract nonReentrant returns (uint256) {
        address[] memory sequencers = getStakersFromBitmap(sequencersBitmap);

        uint256 valueSum;
        for (uint256 i = 0; i < sequencers.length; i++) {
            if (withdrawals[sequencers[i]] > 0) {
                delete withdrawals[sequencers[i]];
                valueSum += stakingValue;
            } else if (!isStakerInDeleteList(sequencers[i])) {
                // If it is the first time to be slashed
                valueSum += stakingValue;
                _removeStaker(sequencers[i]);
                // remove from whitelist
                delete whitelist[sequencers[i]];
                removedList[sequencers[i]] = true;
            }
        }

        uint256 reward = (valueSum * rewardPercentage) / 100;
        slashRemaining += valueSum - reward;
        _transfer(rollupContract, reward);

        emit Slashed(sequencers);
        emit StakersRemoved(sequencers);

        // send message to remove stakers on l2
        _msgRemoveStakers(sequencers);

        return reward;
    }

    /// @notice claim slash remaining
    /// @param receiver  receiver address
    function claimSlashRemaining(address receiver) external onlyOwner nonReentrant {
        uint256 _slashRemaining = slashRemaining;
        _transfer(receiver, slashRemaining);
        slashRemaining = 0;
        emit SlashRemainingClaimed(receiver, _slashRemaining);
    }

    /// @notice update gas limit of add staker
    /// @param _gasLimitAdd       cross-chain gas limit add staker
    function updateGasLimitAddStaker(uint256 _gasLimitAdd) external onlyOwner {
        require(_gasLimitAdd > 0 && _gasLimitAdd != gasLimitAddStaker, "invalid new gas limit");
        uint256 _oldGasLimitAddStaker = gasLimitAddStaker;
        gasLimitAddStaker = _gasLimitAdd;
        emit GasLimitAddStakerUpdated(_oldGasLimitAddStaker, _gasLimitAdd);
    }

    /// @notice update gas limit of remove stakers
    /// @param _gasLimitRemove    cross-chain gas limit remove stakers
    function updateGasLimitRemoveStakers(uint256 _gasLimitRemove) external onlyOwner {
        require(_gasLimitRemove > 0 && _gasLimitRemove != gasLimitRemoveStakers, "invalid new gas limit");
        uint256 _oldGasLimitRemove = gasLimitRemoveStakers;
        gasLimitRemoveStakers = _gasLimitRemove;
        emit GasLimitRemoveStakersUpdated(_oldGasLimitRemove, _gasLimitRemove);
    }

    /// @notice update challenge deposit
    /// @param _challengeDeposit       challenge deposit value
    function updateChallengeDeposit(uint256 _challengeDeposit) external onlyOwner {
        require(_challengeDeposit > 0 && _challengeDeposit != challengeDeposit, "invalid challenge deposit value");
        uint256 _oldChallengeDeposit = challengeDeposit;
        challengeDeposit = _challengeDeposit;
        emit ChallengeDepositUpdated(_oldChallengeDeposit, _challengeDeposit);
    }

    /// @notice update reward percentage
    /// @param _rewardPercentage       percentage awarded to challenger
    function updateRewardPercentage(uint256 _rewardPercentage) external onlyOwner {
        require(
            _rewardPercentage > 0 && _rewardPercentage <= 100 && _rewardPercentage != rewardPercentage,
            "invalid reward percentage"
        );
        uint256 _oldRewardPercentage = rewardPercentage;
        rewardPercentage = _rewardPercentage;
        emit RewardPercentageUpdated(_oldRewardPercentage, _rewardPercentage);
    }

    /// @notice clean staker store
    function cleanStakerStore() external onlyOwner {
        _cleanStakerStore();
    }

    /*****************************
     * Public Mutating Functions *
     *****************************/

    /// @notice claim withdrawal
    /// @param receiver  receiver address
    function claimWithdrawal(address receiver) external nonReentrant {
        require(withdrawals[_msgSender()] > 0, "withdrawal not exist");
        require(withdrawals[_msgSender()] < block.number, "withdrawal locked");

        delete withdrawals[_msgSender()];
        _cleanStakerStore();

        emit Claimed(_msgSender(), receiver);

        _transfer(receiver, stakingValue);
    }

    /*************************
     * Public View Functions *
     *************************/

    /// @notice verify BLS signature
    function verifySignature(
        uint256, // signedSequencersBitmap
        address[] calldata, // sequencerSet
        bytes32, // msgHash
        bytes calldata // signature
    ) external pure returns (bool) {
        // TODO verify BLS signature
        return true;
    }

    /// @notice return all stakers
    function getStakers() external view returns (address[255] memory) {
        return stakerSet;
    }

    /// @notice return active stakers
    function getActiveStakers() external view returns (address[] memory) {
        uint256 activeStakersNumber;
        bool[] memory tags = new bool[](255);
        for (uint256 i = 0; i < 255; i++) {
            // valid address and not in delete list
            if (stakerSet[i] != address(0) && deleteableHeight[stakerSet[i]] == 0) {
                activeStakersNumber++;
                tags[i] = true;
            }
        }
        address[] memory activeStakers = new address[](activeStakersNumber);
        uint256 index;
        for (uint256 i = 0; i < 255; i++) {
            if (tags[i]) {
                activeStakers[index] = stakerSet[i];
                index++;
            }
        }
        return activeStakers;
    }

    /// @notice whether address is staker
    /// @param addr  address to check
    function isStaker(address addr) public view returns (bool) {
        if (stakerIndexes[addr] == 0) {
            return false;
        }
        return stakerSet[stakerIndexes[addr] - 1] == addr;
    }

    /// @notice whether address is active staker
    /// @param addr  address to check
    function isActiveStaker(address addr) public view returns (bool) {
        if (stakerIndexes[addr] == 0) {
            return false;
        }
        return (stakerSet[stakerIndexes[addr] - 1] == addr) && (deleteableHeight[addr] == 0);
    }

    /// @notice whether address in delete list
    /// @param addr  address to check
    function isStakerInDeleteList(address addr) public view returns (bool) {
        return deleteableHeight[addr] > 0;
    }

    /// @notice get staker bitmap
    /// @param _staker  the staker address
    function getStakerBitmap(address _staker) external view returns (uint256 bitmap) {
        require(isStaker(_staker), "invalid staker");
        bitmap = 1 << stakerIndexes[_staker];
        return bitmap;
    }

    /// @notice get stakers bitmap
    /// @param _stakers  the staker address array
    function getStakersBitmap(address[] calldata _stakers) external view returns (uint256 bitmap) {
        require(_stakers.length <= 255, "stakers length out of bounds");
        for (uint256 i = 0; i < _stakers.length; i++) {
            require(isStaker(_stakers[i]), "invalid staker");
            bitmap = bitmap | (1 << stakerIndexes[_stakers[i]]);
        }
        return bitmap;
    }

    /// @notice get stakers from bitmap
    /// @param bitmap  the stakers bitmap
    function getStakersFromBitmap(uint256 bitmap) public view returns (address[] memory stakerAddrs) {
        // skip first bit
        uint256 _bitmap = bitmap >> 1;
        uint256 stakersLength = 0;
        while (_bitmap > 0) {
            stakersLength = stakersLength + 1;
            _bitmap = _bitmap & (_bitmap - 1);
        }

        stakerAddrs = new address[](stakersLength);
        uint256 index = 0;
        for (uint8 i = 1; i < 255; i++) {
            if ((bitmap & (1 << i)) > 0) {
                stakerAddrs[index] = stakerSet[i - 1];
                index = index + 1;
                if (index >= stakersLength) {
                    break;
                }
            }
        }
    }

    /**********************
     * Internal Functions *
     **********************/

    /// @notice add stater
    /// @param addr  staker address
    function _addStaker(address addr) internal {
        for (uint8 i = 0; i < 255; i++) {
            if (stakerSet[i] == address(0)) {
                stakerSet[i] = addr;
                stakerIndexes[addr] = i + 1;
                return;
            }
        }
        require(false, "slot full");
    }

    /// @notice Add staker to deleteList, it will not be actually deleted until cleanStakerStore is executed
    /// @param addr  staker address
    function _removeStaker(address addr) internal {
        require(deleteableHeight[addr] == 0, "already in deleteList");
        deleteList.push(addr);
        deleteableHeight[addr] = block.number + withdrawalLockBlocks;
    }

    /// @notice transfer ETH
    /// @param _to      The address to transfer ETH to.
    /// @param _amount  The amount of ETH to transfer.
    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}("");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /// @notice add staker
    /// @param add       staker to add
    function _msgAddStaker(Types.StakerInfo memory add) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.addStaker, (add)),
            gasLimitAddStaker
        );
    }

    /// @notice remove stakers
    /// @param remove    stakers to remove
    function _msgRemoveStakers(address[] memory remove) internal {
        MESSENGER.sendMessage(
            address(OTHER_STAKING),
            0,
            abi.encodeCall(IL2Staking.removeStakers, (remove)),
            gasLimitRemoveStakers
        );
    }

    /// @notice clean staker store
    function _cleanStakerStore() internal {
        uint256 i = 0;
        while (i < deleteList.length) {
            if (deleteableHeight[deleteList[i]] <= block.number) {
                // clean stakerSet
                delete stakerSet[stakerIndexes[deleteList[i]] - 1];
                delete stakerIndexes[deleteList[i]];

                // clean staker info
                delete stakers[deleteList[i]];

                // clean deleteList
                delete deleteableHeight[deleteList[i]];
                deleteList[i] = deleteList[deleteList.length - 1];
                deleteList.pop();
            } else {
                i++;
            }
        }
    }
}
