// SPDX-License-Identifier: MIT
pragma solidity =0.8.23;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Types} from "../../libraries/common/Types.sol";
import {IL1Sequencer} from "./IL1Sequencer.sol";
import {IL2Sequencer} from "../../L2/staking/IL2Sequencer.sol";

contract Staking is Initializable, OwnableUpgradeable {
    // Staker info
    struct StakingInfo {
        address addr;
        bytes32 tmKey;
        bytes blsKey;
        uint256 balance;
    }
    // withdrawal info
    struct Withdrawal {
        uint256 balance;
        uint256 unlock;
        bool exit;
    }

    // sequencer contract
    address public sequencerContract;

    // Staking limit
    uint256 public limit;
    // Exit lock blocks
    uint256 public lock;

    // staker whitelist
    mapping(address => bool) public whitelist;

    // all stakers, sort by staking value
    address[] public stakers;
    // all staker infos
    mapping(address => StakingInfo) public stakings;
    // withdrawl infos
    mapping(address => Withdrawal) public withdrawals;

    // stakers size reached sequencersSize first time
    bool initialized = false;
    // total number of sequencers
    uint256 public sequencersSize = 0;

    // current sequencer bls keys, without sort
    bytes[] public sequencers;

    // TODO: ETH LP token supported
    // mapping(address => bool) ethLPTokens;

    /**
     * @notice staker registered
     */
    event Registered(
        address addr,
        bytes32 tmKey,
        bytes blsKey,
        uint256 balance
    );

    /**
     * @notice staker staked
     */
    event Staked(address addr, uint256 balance);

    /**
     * @notice withdrawed
     */
    event Withdrawed(address addr, uint256 balance);

    /**
     * @notice staker claimed
     */
    event Claimed(address addr, uint256 balance);

    /**
     * @notice sequencer updated
     */
    event SequencerUpdated(bytes[] sequencers, uint256 version);

    /**
     * @notice only staker
     */
    modifier onlyStaker() {
        bool isStaker = false;
        for (uint256 i = 0; i < stakers.length; i++) {
            if (stakers[i] == msg.sender) {
                isStaker = true;
                break;
            }
        }
        require(isStaker, "staker not exist");
        _;
    }

    /**
     * @notice xxx
     */
    modifier inWhitelist() {
        require(whitelist[msg.sender], "not in whitelist");
        _;
    }

    /**
     * @notice xxx
     */
    modifier noStaker() {
        bool isStaker = false;
        for (uint256 i = 0; i < stakers.length; i++) {
            if (stakers[i] == msg.sender) {
                isStaker = true;
                break;
            }
        }
        require(!isStaker, "already registered");
        _;
    }

    /**
     * @notice xxx
     */
    modifier noExit() {
        require(!withdrawals[msg.sender].exit, "staker is exited");
        _;
    }

    /**
     * @notice initializer
     * @param _admin params admin
     * @param _sequencerContract sequencer contract address
     * @param _sequencersSize size of sequencer set
     * @param _limit smallest staking value
     */
    function initialize(
        address _admin,
        address _sequencerContract,
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _lock
    ) public initializer {
        require(_sequencerContract != address(0), "invalid sequencer contract");
        require(_sequencersSize > 0, "sequencersSize must greater than 0");
        require(_limit > 0, "staking limit must greater than 0");
        sequencerContract = _sequencerContract;
        sequencersSize = _sequencersSize;
        limit = _limit;
        lock = _lock;
        __Ownable_init();
        _transferOwnership(_admin);
    }

    /**
     * @notice register staker
     * @param tmKey tendermint pubkey
     * @param blsKey bls pubkey
     * @param _minGasLimit Minimum amount of gas that the bridge can be relayed with.
     */
    function register(
        bytes32 tmKey,
        bytes memory blsKey,
        uint32 _minGasLimit
    ) external payable inWhitelist noStaker noExit {
        require(sequencersSize > 0, "sequencersSize must greater than 0");
        require(tmKey != 0, "invalid tendermint pubkey");
        require(blsKey.length == 256, "invalid bls pubkey");
        require(msg.value > limit, "staking value is not enough");

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

        stakings[msg.sender] = StakingInfo(
            msg.sender,
            tmKey,
            blsKey,
            msg.value
        );
        stakers.push(msg.sender);

        emit Registered(msg.sender, tmKey, blsKey, msg.value);

        // sort sequencers
        uint256 i = stakers.length - 1;
        while (i > 0) {
            if (
                stakings[stakers[i]].balance > stakings[stakers[i - 1]].balance
            ) {
                address tmp = stakers[i - 1];
                stakers[i - 1] = stakers[i];
                stakers[i] = tmp;
            } else {
                break;
            }
            i--;
        }

        // stakers size reached sequencersSize first time
        if (!initialized && stakers.length == sequencersSize) {
            initialized = true;
            updateSequencers(_minGasLimit);
            return;
        }

        if (
            initialized &&
            (stakers.length <= sequencersSize || i < sequencersSize)
        ) {
            updateSequencers(_minGasLimit);
        }
    }

    /**
     * @notice stake ETH
     */
    function stakeETH(
        uint32 _minGasLimit
    ) external payable inWhitelist onlyStaker {
        require(
            msg.value > 0 && stakings[msg.sender].balance + msg.value > limit,
            "staking value not enough"
        );
        stakings[msg.sender].balance += msg.value;

        emit Staked(msg.sender, stakings[msg.sender].balance);

        uint256 indexBeforeSort = getStakerIndex(msg.sender);

        for (uint256 i = stakers.length - 1; i > 0; i--) {
            if (
                stakings[stakers[i]].balance > stakings[stakers[i - 1]].balance
            ) {
                address tmp = stakers[i - 1];
                stakers[i - 1] = stakers[i];
                stakers[i] = tmp;
            }
        }

        uint256 indexAfterSort = getStakerIndex(msg.sender);

        if (
            initialized &&
            indexBeforeSort >= sequencersSize &&
            indexAfterSort < sequencersSize
        ) {
            updateSequencers(_minGasLimit);
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
    function withdrawETH(uint32 _minGasLimit) external payable noExit {
        uint256 index = getStakerIndex(msg.sender);

        withdrawals[msg.sender] = Withdrawal(
            stakings[msg.sender].balance,
            block.number + lock,
            true
        );
        emit Withdrawed(msg.sender, withdrawals[msg.sender].balance);

        for (uint256 i = index; i < stakers.length - 1; i++) {
            stakers[i] = stakers[i + 1];
        }
        stakers.pop();
        delete stakings[msg.sender];

        if (stakers.length == 0) {
            IL1Sequencer(sequencerContract).pause();
            return;
        }

        if (index < sequencersSize) {
            updateSequencers(_minGasLimit);
        }
    }

    /**
     * @notice update params
     */
    function updateParams(
        uint256 _sequencersSize,
        uint32 _minGasLimit
    ) external onlyOwner {
        require(
            _sequencersSize != sequencersSize &&
                _sequencersSize >= stakers.length &&
                _sequencersSize > 0,
            "invalid new sequencers size"
        );

        if (sequencersSize < stakers.length) {
            sequencersSize = _sequencersSize;
            updateSequencers(_minGasLimit);
            return;
        }
        sequencersSize = _sequencersSize;
    }

    /**
     * @notice update whitelist
     */
    function updateWhitelist(
        address[] calldata add,
        address[] calldata remove
    ) external onlyOwner {
        for (uint i = 0; i < add.length; i++) {
            whitelist[add[i]] = true;
        }
        for (uint i = 0; i < remove.length; i++) {
            whitelist[remove[i]] = false;
        }
    }

    /**
     * @notice claim ETH
     */
    function claimETH() external {
        require(
            withdrawals[msg.sender].exit &&
                withdrawals[msg.sender].balance > 0 &&
                block.number > withdrawals[msg.sender].unlock,
            "invalid withdrawal"
        );
        payable(msg.sender).transfer(withdrawals[msg.sender].balance);
        emit Claimed(msg.sender, withdrawals[msg.sender].balance);
        delete withdrawals[msg.sender];
    }

    /**
     * @notice update sequencer set
     */
    function updateSequencers(uint32 _gasLimit) internal {
        delete sequencers;

        uint256 sequencersCount = sequencersSize;
        if (stakers.length < sequencersSize) {
            sequencersCount = stakers.length;
        }

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            sequencersCount
        );
        for (uint256 i = 0; i < sequencersCount; i++) {
            sequencers.push(stakings[stakers[i]].blsKey);
            sequencerInfos[i] = Types.SequencerInfo(
                stakings[stakers[i]].addr, // addr;
                stakings[stakers[i]].tmKey, // tmKey;
                stakings[stakers[i]].blsKey // blsKey;
            );
        }

        // abi encode updateSequencers data
        bytes memory data = abi.encodeWithSelector(
            IL2Sequencer.updateSequencers.selector,
            // Because this call will be executed on the remote chain, we reverse the order of
            // the remote and local token addresses relative to their order in the
            // updateSequencers function.
            IL1Sequencer(sequencerContract).newestVersion() + 1,
            sequencerInfos
        );
        IL1Sequencer(sequencerContract).updateAndSendSequencerSet(
            data,
            sequencers,
            _gasLimit,
            _msgSender()
        );

        emit SequencerUpdated(
            sequencers,
            IL1Sequencer(sequencerContract).newestVersion()
        );
    }

    function stakersNumber() public view returns (uint256) {
        return stakers.length;
    }
}
