// SPDX-License-Identifier: MIT
pragma solidity =0.8.16;

/* Testing utilities */
import {Test, StdUtils} from "forge-std/Test.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {L2ToL1MessagePasser} from "../L2/L2ToL1MessagePasser.sol";
import {L1StandardBridge} from "../L1/L1StandardBridge.sol";
import {L2StandardBridge} from "../L2/L2StandardBridge.sol";
import {L1ERC721Bridge} from "../L1/L1ERC721Bridge.sol";
import {L2ERC721Bridge} from "../L2/L2ERC721Bridge.sol";
import {MorphMintableERC20Factory} from "../universal/MorphMintableERC20Factory.sol";
import {MorphMintableERC721Factory} from "../universal/MorphMintableERC721Factory.sol";
import {MorphMintableERC20} from "../universal/MorphMintableERC20.sol";
import {MorphPortal} from "../L1/MorphPortal.sol";
import {L1CrossDomainMessenger} from "../L1/L1CrossDomainMessenger.sol";
import {L2CrossDomainMessenger} from "../L2/L2CrossDomainMessenger.sol";
import {AddressAliasHelper} from "../vendor/AddressAliasHelper.sol";
import {LegacyERC20ETH} from "../legacy/LegacyERC20ETH.sol";
import {Predeploys} from "../libraries/Predeploys.sol";
import {Types} from "../libraries/Types.sol";
import {Proxy} from "../universal/Proxy.sol";
import {ResolvedDelegateProxy} from "../legacy/ResolvedDelegateProxy.sol";
import {AddressManager} from "../legacy/AddressManager.sol";
import {L1ChugSplashProxy} from "../legacy/L1ChugSplashProxy.sol";
import {IL1ChugSplashDeployer} from "../legacy/L1ChugSplashProxy.sol";
import {LegacyMintableERC20} from "../legacy/LegacyMintableERC20.sol";
import {SystemConfig} from "../L1/SystemConfig.sol";
import {ResourceMetering} from "../L1/ResourceMetering.sol";
import {Constants} from "../libraries/Constants.sol";
import {Rollup} from "../L1/Rollup.sol";
import {IRollup} from "../L1/IRollup.sol";
import {Tree} from "../universal/Tree.sol";
import {Sequencer} from "../universal/Sequencer.sol";
import {L1Sequencer} from "../L1/staking/L1Sequencer.sol";
import {L2Sequencer} from "../L2/L2Sequencer.sol";
import {Gov} from "../L2/Gov.sol";
import {Submitter} from "../L2/Submitter.sol";
import {Staking} from "../L1/staking/Staking.sol";
import {MockZkEvmVerifier} from "../mock/MockZkEvmVerifier.sol";

contract CommonTest is Test, Tree {
    address alice = address(128);
    address bob = address(256);
    address multisig = address(512);

    address immutable ZERO_ADDRESS = address(0);
    address immutable NON_ZERO_ADDRESS = address(1);
    uint256 immutable NON_ZERO_VALUE = 100;
    uint256 immutable ZERO_VALUE = 0;
    uint64 immutable NON_ZERO_GASLIMIT = 50000;
    uint32 public defultGasLimit = 1000000;
    bytes32 nonZeroHash = keccak256(abi.encode("NON_ZERO"));
    bytes NON_ZERO_DATA =
        hex"0000111122223333444455556666777788889999aaaabbbbccccddddeeeeffff0000";

    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 indexed version,
        bytes opaqueData
    );

    FFIInterface ffi;

    function setUp() public virtual {
        // Give alice and bob some ETH
        vm.deal(alice, 1 << 16);
        vm.deal(bob, 1 << 16);
        vm.deal(multisig, 1 << 16);

        vm.label(alice, "alice");
        vm.label(bob, "bob");
        vm.label(multisig, "multisig");

        // Make sure we have a non-zero base fee
        vm.fee(1000000000);

        ffi = new FFIInterface();
    }

    function emitTransactionDeposited(
        address _from,
        address _to,
        uint256 _mint,
        uint256 _value,
        uint64 _gasLimit,
        bool _isCreation,
        bytes memory _data
    ) internal {
        emit TransactionDeposited(
            _from,
            _to,
            0,
            abi.encodePacked(_mint, _value, _gasLimit, _isCreation, _data)
        );
    }
}

contract Messenger_Initializer is CommonTest  {
    Proxy public portalProxy;

    L2ToL1MessagePasser messagePasser =
        L2ToL1MessagePasser(payable(Predeploys.L2_TO_L1_MESSAGE_PASSER));

    AddressManager public addressManager;
    L1CrossDomainMessenger public L1Messenger;
    L2CrossDomainMessenger public L2Messenger =
        L2CrossDomainMessenger(Predeploys.L2_CROSS_DOMAIN_MESSENGER);

    IRollup.BatchSignature public nilBatchSig;

    event Mint(address indexed account, uint256 amount);
    event Burn(address indexed account, uint256 amount);

    event SentMessage(
        address indexed target,
        address sender,
        bytes message,
        uint256 messageNonce,
        uint256 gasLimit
    );

    event SentMessageExtension1(address indexed sender, uint256 value);

    event MessagePassed(
        uint256 indexed nonce,
        address indexed sender,
        address indexed target,
        uint256 value,
        uint256 gasLimit,
        bytes data,
        bytes32 withdrawalHash,
        bytes32 rootHash
    );

    event WithdrawerBalanceBurnt(uint256 indexed amount);

    event RelayedMessage(bytes32 indexed msgHash);
    event FailedRelayedMessage(bytes32 indexed msgHash);

    event TransactionDeposited(
        address indexed from,
        address indexed to,
        uint256 mint,
        uint256 value,
        uint64 gasLimit,
        bool isCreation,
        bytes data
    );

    event WhatHappened(bool success, bytes returndata);

    function setUp() public virtual override {
        super.setUp();

        // Deploy the address manager
        vm.prank(multisig);
        addressManager = new AddressManager();
        vm.prank(multisig);
        portalProxy = new Proxy(multisig);

        // Setup implementation
        L1CrossDomainMessenger L1MessengerImpl = new L1CrossDomainMessenger(
            MorphPortal(payable(address(portalProxy)))
        );

        // Setup the address manager and proxy
        vm.prank(multisig);
        addressManager.setAddress(
            "L1CrossDomainMessenger",
            address(L1MessengerImpl)
        );
        ResolvedDelegateProxy proxy = new ResolvedDelegateProxy(
            addressManager,
            "L1CrossDomainMessenger"
        );
        L1Messenger = L1CrossDomainMessenger(address(proxy));
        L1Messenger.initialize();

        vm.etch(
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            address(new L2CrossDomainMessenger(address(L1Messenger))).code
        );

        L2Messenger.initialize();

        // Label addresses
        vm.label(address(addressManager), "AddressManager");
        vm.label(address(L1MessengerImpl), "L1CrossDomainMessenger_Impl");
        vm.label(address(L1Messenger), "L1CrossDomainMessenger_Proxy");
        vm.label(Predeploys.LEGACY_ERC20_ETH, "LegacyERC20ETH");
        vm.label(
            Predeploys.L2_CROSS_DOMAIN_MESSENGER,
            "L2CrossDomainMessenger"
        );

        vm.label(
            AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger)),
            "L1CrossDomainMessenger_aliased"
        );

        // Set the L2ToL1MessagePasser at the correct address
        vm.etch(
            Predeploys.L2_TO_L1_MESSAGE_PASSER,
            address(new L2ToL1MessagePasser(address(L1Messenger))).code
        );
        vm.label(Predeploys.L2_TO_L1_MESSAGE_PASSER, "L2ToL1MessagePasser");

        nilBatchSig = IRollup.BatchSignature({
            version: 0,
            signers: new uint256[](0),
            signature: hex"123456"
        });
    }
}

contract Rollup_Initializer is Messenger_Initializer {
    Rollup public rollupImpl;
    Rollup public rollup;
    MockZkEvmVerifier public verifier;
    // Constructor arguments
    uint256 public PROOF_WINDOW = 100;
    uint256 public FINALIZATION_PERIOD_SECONDS = 2;
    uint256 public MIN_DEPOSIT = 1000000000000000000; // 1 eth
    uint256 public maxNumTxInChunk = 10;
    uint64 public layer2ChainId = 53077;
    uint32 public minGasLimit = 10000;
    address guardian;

    // Test data
    uint256 initL1Time;

    bytes32 genesisRoot = bytes32(0);
    uint64 genesisBlockNumber = 0;

    event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash);

    event ChallengeState(
        uint64 indexed batchIndex,
        address challenger,
        uint256 challengeDeposit
    );
    event SubmitBatches(uint64 indexed numBatch, uint64 l2Num);
    event ChallengeRes(uint64 indexed batchIndex, address winner, string res);
    event FinalizeBatch(
        uint256 indexed batchIndex,
        bytes32 indexed batchHash,
        bytes32 stateRoot,
        bytes32 withdrawRoot
    );
    event UpdateMaxNumTxInChunk(
        uint256 oldMaxNumTxInChunk,
        uint256 newMaxNumTxInChunk
    );
    event UpdateVerifier(
        address indexed oldVerifier,
        address indexed newVerifier
    );
    event UpdateProver(address indexed account, bool status);
    event UpdateSequencer(address indexed account, bool status);
    event UpdateChallenger(address indexed account, bool status);

    function setUp() public virtual override {
        super.setUp();
        vm.startPrank(multisig);
        // deploy mock verifier
        verifier = new MockZkEvmVerifier();
        // Deploy rollup impl
        rollupImpl = new Rollup(layer2ChainId, payable(address(L1Messenger)));
        Proxy proxy = new Proxy(multisig);
        proxy.upgradeToAndCall(
            address(rollupImpl),
            abi.encodeWithSelector(
                Rollup.initialize.selector,
                address(portalProxy), // _messageQueue
                address(verifier), // _verifier
                maxNumTxInChunk, // _maxNumTxInChunk
                MIN_DEPOSIT, // _minDeposit
                FINALIZATION_PERIOD_SECONDS, // _finalizationPeriodSeconds
                PROOF_WINDOW // _proofWindow
            )
        );
        rollup = Rollup(payable(address(proxy)));
        vm.label(address(rollup), "Rollup");
        rollup.addSequencer(alice);
        rollup.addProver(alice);
        rollup.addProver(bob);
        rollup.addChallenger(bob);
        vm.stopPrank();
    }
}

contract Portal_Initializer is Rollup_Initializer {
    // Test target
    MorphPortal internal portalImpl;
    MorphPortal internal portal;
    SystemConfig systemConfig;

    event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success);
    event WithdrawalProven(
        bytes32 indexed withdrawalHash,
        address indexed from,
        address indexed to
    );

    function setUp() public virtual override {
        super.setUp();
        ResourceMetering.ResourceConfig memory config = Constants
            .DEFAULT_RESOURCE_CONFIG();

        systemConfig = new SystemConfig({
            _owner: address(1),
            _overhead: 0,
            _scalar: 10000,
            _batcherHash: bytes32(0),
            _gasLimit: 30_000_000,
            _unsafeBlockSigner: address(0),
            _config: config
        });

        portalImpl = new MorphPortal({
            _guardian: guardian,
            _paused: true,
            _config: systemConfig,
            _rollup: rollup,
            _l1Messenger: address(L1Messenger)
        });

        vm.prank(multisig);
        portalProxy.upgradeToAndCall(
            address(portalImpl),
            abi.encodeWithSelector(
                MorphPortal.initialize.selector,
                false,
                address(L1Messenger)
            )
        );
        portal = MorphPortal(payable(address(portalProxy)));
        vm.label(address(portal), "MorphPortal");
    }
}

contract Staking_Initializer is Portal_Initializer {
    event Registered(
        address addr,
        bytes32 tmKey,
        bytes blsKey,
        uint256 balance
    );
    event ACKRollup(
        uint256 batchIndex,
        address submitter,
        uint256 batchStartBlock,
        uint256 batchEndBlock,
        uint256 rollupTime
    );
    event EpochUpdated(uint256 interval, uint256 sequencersLen);
    event Staked(address addr, uint256 balance);
    event Withdrawed(address addr, uint256 balance);
    event Claimed(address addr, uint256 balance);
    event SequencerUpdated(bytes[] sequencers, uint256 version);

    uint256 public beginSeq = 10;
    uint256 public version = 0;
    Staking public staking;
    L1Sequencer public l1Sequencer;
    L2Sequencer public l2Sequencer;
    Submitter public l2Submitter;
    Gov public l2Gov;
    uint256 public NEXT_EPOCH_START = 1700000000;
    uint256 public constant SEQUENCER_SIZE = 3;
    uint256 public LOCK = 3;
    bytes[] public sequencerBLSKeys;
    uint256 public PROPOSAL_INTERVAL = 1000;
    uint256 public ROLLUP_EPOCH = 1000;
    uint256 public MAX_CHUNKS = 1000000000;

    function setUp() public virtual override {
        super.setUp();
        vm.startPrank(multisig);
        Proxy stakingProxy = new Proxy(multisig);
        Proxy l1SequencerProxy = new Proxy(multisig);
        Proxy l2SequencerProxy = new Proxy(multisig);
        Proxy l2SubmitterProxy = new Proxy(multisig);
        Proxy l2GovProxy = new Proxy(multisig);

        // set Submitter
        Submitter l2SubmitterImpl = new Submitter(
            payable(address(l1SequencerProxy))
        );

        vm.etch(Predeploys.L2_SUBMITTER, address(l2SubmitterProxy).code);
        l2SubmitterProxy = Proxy(payable(address(Predeploys.L2_SUBMITTER)));
        vm.store(
            address(l2SubmitterProxy),
            bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1),
            bytes32(abi.encode(multisig))
        );

        l2SubmitterProxy.upgradeToAndCall(
            address(l2SubmitterImpl),
            abi.encodeWithSelector(
                Submitter.initialize.selector,
                NEXT_EPOCH_START
            )
        );

        l2Submitter = Submitter(payable(address(Predeploys.L2_SUBMITTER)));
        vm.label(Predeploys.L2_SUBMITTER, "L2SUBMITTER");

        // set L2Sequencer
        L2Sequencer l2SequencerImpl = new L2Sequencer(
            payable(address(l1SequencerProxy))
        );

        // Set the L2Sequencer at the correct address
        vm.etch(Predeploys.L2_SEQUENCER, address(l2SequencerProxy).code);
        l2SequencerProxy = Proxy(payable(address(Predeploys.L2_SEQUENCER)));
        vm.store(
            address(l2SequencerProxy),
            bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1),
            bytes32(abi.encode(multisig))
        );
        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerInfos[i] = sequencerInfo;
        }
        l2SequencerProxy.upgradeToAndCall(
            address(l2SequencerImpl),
            abi.encodeWithSelector(
                L2Sequencer.initialize.selector,
                sequencerInfos
            )
        );

        l2Sequencer = L2Sequencer(payable(address(Predeploys.L2_SEQUENCER)));
        vm.label(Predeploys.L2_SEQUENCER, "L2Sequencer");

        // set Staking
        Staking stakingImpl = new Staking();
        stakingProxy.upgradeToAndCall(
            address(stakingImpl),
            abi.encodeWithSelector(
                Staking.initialize.selector,
                alice, // admin
                address(l1SequencerProxy), // l1Sequencer address
                SEQUENCER_SIZE, // sequencersSize
                MIN_DEPOSIT, // limit
                LOCK // lock
            )
        );
        staking = Staking(payable(address(stakingProxy)));

        // set L1Sequencer
        L1Sequencer l1SequencerImpl = new L1Sequencer(
            payable(address(L1Messenger))
        );
        l1SequencerProxy.upgradeToAndCall(
            address(l1SequencerImpl),
            abi.encodeWithSelector(
                L1Sequencer.initialize.selector,
                address(staking),
                address(rollup)
            )
        );
        l1Sequencer = L1Sequencer(payable(address(l1SequencerProxy)));

        // set Gov
        Gov l2GovImpl = new Gov();
        // Set the L2Gov at the correct address
        vm.etch(Predeploys.L2_GOV, address(l2GovProxy).code);
        l2GovProxy = Proxy(payable(address(Predeploys.L2_GOV)));
        vm.store(
            address(l2GovProxy),
            bytes32(uint256(keccak256("eip1967.proxy.admin")) - 1),
            bytes32(abi.encode(multisig))
        );

        l2GovProxy.upgradeToAndCall(
            address(l2GovImpl),
            abi.encodeWithSelector(
                Gov.initialize.selector,
                PROPOSAL_INTERVAL, // _proposalInterval
                0, // _batchBlockInterval
                0, // _batchMaxBytes
                FINALIZATION_PERIOD_SECONDS, // _batchTimeout
                ROLLUP_EPOCH, // rollupEpoch
                MAX_CHUNKS // maxChunks
            )
        );

        l2Gov = Gov(payable(address(Predeploys.L2_GOV)));
        vm.label(Predeploys.L2_GOV, "L2Gov");

        vm.stopPrank();
    }
}

contract Bridge_Initializer is Portal_Initializer {
    L1StandardBridge L1Bridge;
    L2StandardBridge L2Bridge;
    MorphMintableERC20Factory L2TokenFactory;
    MorphMintableERC20Factory L1TokenFactory;
    ERC20 L1Token;
    ERC20 BadL1Token;
    MorphMintableERC20 L2Token;
    LegacyMintableERC20 LegacyL2Token;
    ERC20 NativeL2Token;
    ERC20 BadL2Token;
    MorphMintableERC20 RemoteL1Token;

    event ETHDepositInitiated(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data,
        uint256 messageNonce
    );

    event ETHWithdrawalFinalized(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data
    );

    event ERC20DepositInitiated(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data,
        uint256 messageNonce
    );

    event ERC20WithdrawalFinalized(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event WithdrawalInitiated(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data,
        uint256 messageNonce
    );

    event DepositFinalized(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event DepositFailed(
        address indexed l1Token,
        address indexed l2Token,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event ETHBridgeInitiated(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data
    );

    event ETHBridgeFinalized(
        address indexed from,
        address indexed to,
        uint256 amount,
        bytes data
    );

    event ERC20BridgeInitiated(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    event ERC20BridgeFinalized(
        address indexed localToken,
        address indexed remoteToken,
        address indexed from,
        address to,
        uint256 amount,
        bytes data
    );

    function setUp() public virtual override {
        super.setUp();

        vm.label(Predeploys.L2_STANDARD_BRIDGE, "L2StandardBridge");
        vm.label(
            Predeploys.Morph_MINTABLE_ERC20_FACTORY,
            "MorphMintableERC20Factory"
        );

        // Deploy the L1 bridge and initialize it with the address of the
        // L1CrossDomainMessenger
        L1ChugSplashProxy proxy = new L1ChugSplashProxy(multisig);
        vm.mockCall(
            multisig,
            abi.encodeWithSelector(IL1ChugSplashDeployer.isUpgrading.selector),
            abi.encode(true)
        );
        vm.startPrank(multisig);
        proxy.setCode(
            address(new L1StandardBridge(payable(address(L1Messenger)))).code
        );
        vm.clearMockedCalls();
        address L1Bridge_Impl = proxy.getImplementation();
        vm.stopPrank();

        L1Bridge = L1StandardBridge(payable(address(proxy)));

        vm.label(address(proxy), "L1StandardBridge_Proxy");
        vm.label(address(L1Bridge_Impl), "L1StandardBridge_Impl");

        // Deploy the L2StandardBridge, move it to the correct predeploy
        // address and then initialize it
        L2StandardBridge l2B = new L2StandardBridge(payable(proxy));
        vm.etch(Predeploys.L2_STANDARD_BRIDGE, address(l2B).code);
        L2Bridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE));

        // Set up the L2 mintable token factory
        MorphMintableERC20Factory factory = new MorphMintableERC20Factory(
            Predeploys.L2_STANDARD_BRIDGE
        );
        vm.etch(Predeploys.Morph_MINTABLE_ERC20_FACTORY, address(factory).code);
        L2TokenFactory = MorphMintableERC20Factory(
            Predeploys.Morph_MINTABLE_ERC20_FACTORY
        );

        vm.etch(
            Predeploys.LEGACY_ERC20_ETH,
            address(new LegacyERC20ETH()).code
        );

        L1Token = new ERC20("Native L1 Token", "L1T");

        LegacyL2Token = new LegacyMintableERC20({
            _l2Bridge: address(L2Bridge),
            _l1Token: address(L1Token),
            _name: string.concat("LegacyL2-", L1Token.name()),
            _symbol: string.concat("LegacyL2-", L1Token.symbol())
        });
        vm.label(address(LegacyL2Token), "LegacyMintableERC20");

        // Deploy the L2 ERC20 now
        L2Token = MorphMintableERC20(
            L2TokenFactory.createStandardL2Token(
                address(L1Token),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        BadL2Token = MorphMintableERC20(
            L2TokenFactory.createStandardL2Token(
                address(1),
                string(abi.encodePacked("L2-", L1Token.name())),
                string(abi.encodePacked("L2-", L1Token.symbol()))
            )
        );

        NativeL2Token = new ERC20("Native L2 Token", "L2T");
        L1TokenFactory = new MorphMintableERC20Factory(address(L1Bridge));

        RemoteL1Token = MorphMintableERC20(
            L1TokenFactory.createStandardL2Token(
                address(NativeL2Token),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );

        BadL1Token = MorphMintableERC20(
            L1TokenFactory.createStandardL2Token(
                address(1),
                string(abi.encodePacked("L1-", NativeL2Token.name())),
                string(abi.encodePacked("L1-", NativeL2Token.symbol()))
            )
        );
    }
}

contract ERC721Bridge_Initializer is Portal_Initializer {
    L1ERC721Bridge L1Bridge;
    L2ERC721Bridge L2Bridge;

    function setUp() public virtual override {
        super.setUp();

        // Deploy the L1ERC721Bridge.
        L1Bridge = new L1ERC721Bridge(
            address(L1Messenger),
            Predeploys.L2_ERC721_BRIDGE
        );

        // Deploy the implementation for the L2ERC721Bridge and etch it into the predeploy address.
        vm.etch(
            Predeploys.L2_ERC721_BRIDGE,
            address(
                new L2ERC721Bridge(
                    Predeploys.L2_CROSS_DOMAIN_MESSENGER,
                    address(L1Bridge)
                )
            ).code
        );

        // Set up a reference to the L2ERC721Bridge.
        L2Bridge = L2ERC721Bridge(Predeploys.L2_ERC721_BRIDGE);

        // Label the L1 and L2 bridges.
        vm.label(address(L1Bridge), "L1ERC721Bridge");
        vm.label(address(L2Bridge), "L2ERC721Bridge");
    }
}

contract FFIInterface is Test {
    function getProveWithdrawalTransactionInputs(
        Types.WithdrawalTransaction memory _tx
    ) external returns (bytes32, bytes32[32] memory, bytes32) {
        string[] memory cmds = new string[](8);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "getProveWithdrawalTransactionInputs";
        cmds[2] = vm.toString(_tx.nonce);
        cmds[3] = vm.toString(_tx.sender);
        cmds[4] = vm.toString(_tx.target);
        cmds[5] = vm.toString(_tx.value);
        cmds[6] = vm.toString(_tx.gasLimit);
        cmds[7] = vm.toString(_tx.data);

        bytes memory result = vm.ffi(cmds);
        (
            bytes32 withdrawalHash,
            bytes32[32] memory withdrawalProof,
            bytes32 withdrawalRoot
        ) = abi.decode(result, (bytes32, bytes32[32], bytes32));

        return (withdrawalHash, withdrawalProof, withdrawalRoot);
    }

    function hashCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes32) {
        string[] memory cmds = new string[](8);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "hashCrossDomainMessage";
        cmds[2] = vm.toString(_nonce);
        cmds[3] = vm.toString(_sender);
        cmds[4] = vm.toString(_target);
        cmds[5] = vm.toString(_value);
        cmds[6] = vm.toString(_gasLimit);
        cmds[7] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashWithdrawal(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes32) {
        string[] memory cmds = new string[](8);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "hashWithdrawal";
        cmds[2] = vm.toString(_nonce);
        cmds[3] = vm.toString(_sender);
        cmds[4] = vm.toString(_target);
        cmds[5] = vm.toString(_value);
        cmds[6] = vm.toString(_gasLimit);
        cmds[7] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function hashL1MessageTx(
        uint64 _queueIndex,
        uint64 _gas,
        address _to,
        uint256 _value,
        bytes memory _data,
        address _sender
    ) external returns (bytes32) {
        string[] memory cmds = new string[](10);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "hashL1MessageTx";
        cmds[2] = vm.toString(_queueIndex);
        cmds[3] = vm.toString(_gas);
        cmds[4] = vm.toString(_to);
        cmds[5] = vm.toString(_value);
        cmds[6] = vm.toString(_data);
        cmds[7] = vm.toString(_sender);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes32));
    }

    function encodeL1MessageTx(
        Types.L1MessageTx calldata txn
    ) external returns (bytes memory) {
        string[] memory cmds = new string[](11);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "encodeL1MessageTx";
        cmds[2] = vm.toString(txn.queueIndex);
        cmds[3] = vm.toString(txn.gas);
        cmds[4] = vm.toString(txn.to);
        cmds[5] = vm.toString(txn.value);
        cmds[6] = vm.toString(txn.data);
        cmds[7] = vm.toString(txn.sender);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function encodeCrossDomainMessage(
        uint256 _nonce,
        address _sender,
        address _target,
        uint256 _value,
        uint256 _gasLimit,
        bytes memory _data
    ) external returns (bytes memory) {
        string[] memory cmds = new string[](8);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "encodeCrossDomainMessage";
        cmds[2] = vm.toString(_nonce);
        cmds[3] = vm.toString(_sender);
        cmds[4] = vm.toString(_target);
        cmds[5] = vm.toString(_value);
        cmds[6] = vm.toString(_gasLimit);
        cmds[7] = vm.toString(_data);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (bytes));
    }

    function decodeVersionedNonce(
        uint256 nonce
    ) external returns (uint256, uint256) {
        string[] memory cmds = new string[](3);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "decodeVersionedNonce";
        cmds[2] = vm.toString(nonce);

        bytes memory result = vm.ffi(cmds);
        return abi.decode(result, (uint256, uint256));
    }

    function generateBatchData(
        uint64 _blockNum,
        bytes32 _preStateRoot,
        bytes32 _withdrawalRoot
    ) external returns (IRollup.BatchData memory) {
        string[] memory cmds = new string[](5);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "generateBatchData";
        cmds[2] = vm.toString(_blockNum);
        cmds[3] = vm.toString(_preStateRoot);
        cmds[4] = vm.toString(_withdrawalRoot);
        bytes memory result = vm.ffi(cmds);
        IRollup.BatchData memory batchData = abi.decode(
            result,
            (IRollup.BatchData)
        );
        return batchData;
    }

    function generateStakingInfo(
        address _staker
    ) external returns (Types.SequencerInfo memory) {
        string[] memory cmds = new string[](3);
        cmds[0] = "scripts/differential-testing/differential-testing";
        cmds[1] = "generateStakingInfo";
        cmds[2] = vm.toString(_staker);

        bytes memory result = vm.ffi(cmds);
        Types.SequencerInfo memory sequencerInfo = abi.decode(
            result,
            (Types.SequencerInfo)
        );
        return sequencerInfo;
    }

    function getMerkleTrieFuzzCase(
        string memory variant
    ) external returns (bytes32, bytes memory, bytes memory, bytes[] memory) {
        string[] memory cmds = new string[](5);
        cmds[0] = "./test-case-generator/fuzz";
        cmds[1] = "-m";
        cmds[2] = "trie";
        cmds[3] = "-v";
        cmds[4] = variant;

        return abi.decode(vm.ffi(cmds), (bytes32, bytes, bytes, bytes[]));
    }
}

// Used for testing a future upgrade beyond the current implementations.
// We include some variables so that we can sanity check accessing storage values after an upgrade.
contract NextImpl is Initializable {
    // Initializable occupies the zero-th slot.
    bytes32 slot1;
    bytes32[19] __gap;
    bytes32 slot21;
    bytes32 public constant slot21Init = bytes32(hex"1337");

    function initialize() public reinitializer(2) {
        // Slot21 is unused by an of our upgradeable contracts.
        // This is used to verify that we can access this value after an upgrade.
        slot21 = slot21Init;
    }
}

contract Reverter {
    fallback() external {
        revert();
    }
}

// Useful for testing reentrancy guards
contract CallerCaller {
    event WhatHappened(bool success, bytes returndata);

    fallback() external {
        (bool success, bytes memory returndata) = msg.sender.call(msg.data);
        emit WhatHappened(success, returndata);
        assembly {
            switch success
            case 0 {
                revert(add(returndata, 0x20), mload(returndata))
            }
            default {
                return(add(returndata, 0x20), mload(returndata))
            }
        }
    }
}

// Used for testing the `CrossDomainMessenger`'s per-message reentrancy guard.
contract ConfigurableCaller {
    bool doRevert = true;
    address target;
    bytes payload;

    event WhatHappened(bool success, bytes returndata);

    /**
     * @notice Call the configured target with the configured payload OR revert.
     */
    function call() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        } else {
            (bool success, bytes memory returndata) = address(target).call(
                payload
            );
            emit WhatHappened(success, returndata);
            assembly {
                switch success
                case 0 {
                    revert(add(returndata, 0x20), mload(returndata))
                }
                default {
                    return(add(returndata, 0x20), mload(returndata))
                }
            }
        }
    }

    /**
     * @notice Set whether or not to have `call` revert.
     */
    function setDoRevert(bool _doRevert) external {
        doRevert = _doRevert;
    }

    /**
     * @notice Set the target for the call made in `call`.
     */
    function setTarget(address _target) external {
        target = _target;
    }

    /**
     * @notice Set the payload for the call made in `call`.
     */
    function setPayload(bytes calldata _payload) external {
        payload = _payload;
    }

    /**
     * @notice Fallback function that reverts if `doRevert` is true.
     *         Otherwise, it does nothing.
     */
    fallback() external {
        if (doRevert) {
            revert("ConfigurableCaller: revert");
        }
    }
}
