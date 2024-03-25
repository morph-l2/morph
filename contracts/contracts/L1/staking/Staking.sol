// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Types} from "../../libraries/common/Types.sol";
import {IL1Sequencer} from "./IL1Sequencer.sol";
import {IStaking} from "./IStaking.sol";
import {IL2Sequencer} from "../../L2/staking/IL2Sequencer.sol";

contract Staking is IStaking, OwnableUpgradeable {
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
    // rollup Contract
    address public rollupContract;

    // Staking limit
    uint256 public override limit;
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

    // current sequencer addresses, without sort
    address[] public sequencersAddr;
    // current sequencer bls keys, without sort
    bytes[] public sequencersBLS;

    // enanble slash
    bool enableSlash;

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
     * @notice params updated
     */
    event ParamsUpdated(uint256 _sequencersSize, uint256 _limit, uint256 _lock);

    /**
     * @notice whitelist updated
     */
    event WhitelistUpdated(address[] add, address[] remove);

    /**
     * @notice only rollup contract
     */
    modifier onlyRollupContract() {
        require(msg.sender == rollupContract, "only rollup contract");
        _;
    }

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
     * @notice whether in whitelist
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
     * @param _rollupContract rollup contract address
     * @param _sequencersSize size of sequencer set
     * @param _limit smallest staking value
     * @param _lock withdraw lock time
     */
    function initialize(
        address _admin,
        address _sequencerContract,
        address _rollupContract,
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _lock
    ) public initializer {
        require(_rollupContract != address(0), "invalid rollup contract");
        require(_sequencerContract != address(0), "invalid sequencer contract");
        require(_sequencersSize > 0, "sequencersSize must greater than 0");
        require(_limit > 0, "staking limit must greater than 0");
        sequencerContract = _sequencerContract;
        rollupContract = _rollupContract;
        sequencersSize = _sequencersSize;
        limit = _limit;
        lock = _lock;
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
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external payable inWhitelist noStaker noExit {
        require(sequencersSize > 0, "sequencersSize must greater than 0");
        require(tmKey != 0, "invalid tendermint pubkey");
        require(blsKey.length == 256, "invalid bls pubkey");
        require(
            limit > 0 && msg.value >= _gasFee + limit,
            "staking value is not enough"
        );

        uint256 stakingAmount = msg.value - _gasFee;

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
            stakingAmount
        );
        stakers.push(msg.sender);

        emit Registered(msg.sender, tmKey, blsKey, stakingAmount);

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
            updateSequencers(_minGasLimit, _gasFee);
            return;
        }

        if (
            initialized &&
            (stakers.length <= sequencersSize || i < sequencersSize)
        ) {
            updateSequencers(_minGasLimit, _gasFee);
        }
    }

    /**
     * @notice stake ETH
     */
    function stakeETH(
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external payable inWhitelist onlyStaker {
        require(
            limit > 0 &&
                msg.value > 0 &&
                stakings[msg.sender].balance + msg.value > limit,
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
            updateSequencers(_minGasLimit, _gasFee);
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
    function withdrawETH(
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external payable noExit {
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
            updateSequencers(_minGasLimit, _gasFee);
            IL1Sequencer(sequencerContract).pause();
            return;
        }

        if (index < sequencersSize) {
            updateSequencers(_minGasLimit, _gasFee);
        }
    }

    /**
     * @notice challenger win, slash sequencers
     */
    function toggleSlash(bool enanble) external onlyOwner {
        enableSlash = enanble;
    }

    function unpauseSequencer() external onlyOwner {
        IL1Sequencer(sequencerContract).unpause();
    }

    /**
     * @notice challenger win, slash sequencers
     */
    function slash(
        address[] memory sequencers,
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external onlyRollupContract returns (uint256) {
        if (!enableSlash) {
            return 0;
        }

        // do slash & update sequencer set
        uint256 valueSum;
        for (uint256 i = 0; i < sequencers.length; i++) {
            address sequencer = sequencers[i];
            valueSum += stakings[sequencer].balance;
            uint256 index = getStakerIndex(sequencer);
            for (uint256 j = index; j < stakers.length - 1; j++) {
                stakers[j] = stakers[j + 1];
            }
            stakers.pop();
            delete stakings[sequencer];
        }
        updateSequencers(_minGasLimit, _gasFee);
        if (stakers.length == 0) {
            IL1Sequencer(sequencerContract).pause();
        }
        _transfer(rollupContract, valueSum);
        return valueSum;
    }

    function _transfer(address _to, uint256 _amount) internal {
        if (_amount > 0) {
            (bool success, ) = _to.call{value: _amount}(hex"");
            require(success, "Rollup: ETH transfer failed");
        }
    }

    /**
     * @notice update params
     * @param _limit smallest staking value
     * @param _lock withdraw lock time
     * @param _sequencersSize sequencers size
     * @param _minGasLimit min gas limit
     * @param _gasFee gas fee
     */
    function updateParams(
        uint256 _sequencersSize,
        uint256 _limit,
        uint256 _lock,
        uint32 _minGasLimit,
        uint256 _gasFee
    ) external onlyOwner {
        require(
            _sequencersSize != sequencersSize &&
                _sequencersSize >= stakers.length &&
                _sequencersSize > 0,
            "invalid new sequencers size"
        );

        if (_limit > 0) {
            limit = _limit;
        }
        if (_lock > 0) {
            lock = _lock;
        }

        if (sequencersSize < stakers.length) {
            sequencersSize = _sequencersSize;
            updateSequencers(_minGasLimit, _gasFee);
            return;
        }
        sequencersSize = _sequencersSize;

        emit ParamsUpdated(sequencersSize, limit, lock);
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

        emit WhitelistUpdated(add, remove);
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
    function updateSequencers(uint32 _gasLimit, uint256 _gasFee) internal {
        delete sequencersAddr;
        delete sequencersBLS;

        uint256 sequencersCount = sequencersSize;
        if (stakers.length < sequencersSize) {
            sequencersCount = stakers.length;
        }

        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            sequencersCount
        );
        for (uint256 i = 0; i < sequencersCount; i++) {
            sequencersAddr.push(stakings[stakers[i]].addr);
            sequencersBLS.push(stakings[stakers[i]].blsKey);
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
        IL1Sequencer(sequencerContract).updateAndSendSequencerSet{
            value: _gasFee
        }(data, sequencersAddr, sequencersBLS, _gasLimit, _msgSender());
    }

    function stakersNumber() public view returns (uint256) {
        return stakers.length;
    }
}
