pragma solidity =0.8.16;

import {StdUtils} from "forge-std/StdUtils.sol";
import {Vm} from "forge-std/Vm.sol";
import {MorphPortal} from "../../L1/MorphPortal.sol";
import {L1CrossDomainMessenger} from "../../L1/L1CrossDomainMessenger.sol";
import {Portal_Initializer} from "../CommonTest.t.sol";
import {Types} from "../../libraries/Types.sol";
import {Predeploys} from "../../libraries/Predeploys.sol";
import {Constants} from "../../libraries/Constants.sol";
import {Encoding} from "../../libraries/Encoding.sol";
import {Hashing} from "../../libraries/Hashing.sol";

contract RelayActor is StdUtils {
    // Storage slot of the l2Sender
    uint256 constant senderSlotIndex = 54;

    uint256 public numHashes;
    bytes32[] public hashes;
    bool public reverted = false;

    MorphPortal portal;
    L1CrossDomainMessenger xdm;
    Vm vm;
    bool doFail;

    constructor(
        MorphPortal _portal,
        L1CrossDomainMessenger _xdm,
        Vm _vm,
        bool _doFail
    ) {
        portal = _portal;
        xdm = _xdm;
        vm = _vm;
        doFail = _doFail;
    }

    /**
     * Relays a message to the `L1CrossDomainMessenger` with a random `version`, and `_message`.
     */
    function relay(
        uint8 _version,
        uint8 _value,
        bytes memory _message
    ) external {
        address target = address(0x04); // ID precompile
        address sender = Predeploys.L2_CROSS_DOMAIN_MESSENGER;

        // Set the minimum gas limit to the cost of the identity precompile's execution for
        // the given message.
        // ID Precompile cost can be determined by calculating: 15 + 3 * data_word_length
        uint32 minGasLimit = uint32(15 + 3 * ((_message.length + 31) / 32));

        // set the value of portal.l2Sender() to be the L2 Cross Domain Messenger.
        vm.store(
            address(portal),
            bytes32(senderSlotIndex),
            bytes32(abi.encode(sender))
        );

        // Restrict version to the range of [0, 1]
        _version = _version % 2;

        // Restrict the value to the range of [0, 1]
        // This is just so we get variance of calls with and without value. The ID precompile
        // will not reject value being sent to it.
        _value = _value % 2;

        // If the message should succeed, supply it `baseGas`. If not, supply it an amount of
        // gas that is too low to complete the call.
        uint256 gas = doFail
            ? bound(minGasLimit, 60_000, 80_000)
            : xdm.baseGas(_message, minGasLimit);

        // Compute the cross domain message hash and store it in `hashes`.
        // The `relayMessage` function will always encode the message as a version 1
        // message after checking that the V0 hash has not already been relayed.
        bytes32 _hash = Hashing.hashCrossDomainMessageV1(
            Encoding.encodeVersionedNonce(0, _version),
            sender,
            target,
            _value,
            minGasLimit,
            _message
        );
        hashes.push(_hash);
        numHashes += 1;

        // Make sure we've got a fresh message.
        vm.assume(
            xdm.successfulMessages(_hash) == false &&
                xdm.failedMessages(_hash) == false
        );

        // Act as the Morph portal and call `relayMessage` on the `L1CrossDomainMessenger` with
        // the outer min gas limit.
        vm.startPrank(address(portal));
        if (!doFail) {
            vm.expectCallMinGas(address(0x04), _value, minGasLimit, _message);
        }
        try
            xdm.relayMessage{gas: gas, value: _value}(
                Encoding.encodeVersionedNonce(0, _version),
                sender,
                target,
                _value,
                minGasLimit,
                _message
            )
        {} catch {
            // If any of these calls revert, set `reverted` to true to fail the invariant test.
            // NOTE: This is to get around forge's invariant fuzzer ignoring reverted calls
            // to this function.
            reverted = true;
        }
        vm.stopPrank();
    }
}

contract XDM_MinGasLimits is Portal_Initializer {
    RelayActor actor;

    function init(bool doFail) public virtual {
        // Set up the `L1CrossDomainMessenger` and `MorphPortal` contracts.
        super.setUp();

        // Deploy a relay actor
        actor = new RelayActor(portal, L1Messenger, vm, doFail);

        // Give the portal some ether to send to `relayMessage`
        vm.deal(address(portal), type(uint128).max);

        // Target the `RelayActor` contract
        targetContract(address(actor));

        // Don't allow the estimation address to be the sender
        excludeSender(Constants.ESTIMATION_ADDRESS);

        // Target the actor's `relay` function
        bytes4[] memory selectors = new bytes4[](1);
        selectors[0] = actor.relay.selector;
        targetSelector(
            FuzzSelector({addr: address(actor), selectors: selectors})
        );
    }
}

contract XDM_MinGasLimits_Succeeds is XDM_MinGasLimits {
    function setUp() public override {
        // Don't fail
        super.init(false);
    }

    /**
     * @custom:invariant A call to `relayMessage` should succeed if at least the minimum gas limit
     *                   can be supplied to the target context, there is enough gas to complete
     *                   execution of `relayMessage` after the target context's execution is
     *                   finished, and the target context did not revert.
     *
     * There are two minimum gas limits here:
     *
     * - The outer min gas limit is for the call from the `MorphPortal` to the
     * `L1CrossDomainMessenger`,  and it can be retrieved by calling the xdm's `baseGas` function
     * with the `message` and inner limit.
     *
     * - The inner min gas limit is for the call from the `L1CrossDomainMessenger` to the target
     * contract.
     */
    function invariant_minGasLimits() external {
        uint256 length = actor.numHashes();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 hash = actor.hashes(i);
            // The message hash is set in the successfulMessages mapping
            assertTrue(L1Messenger.successfulMessages(hash));
            // The message hash is not set in the failedMessages mapping
            assertFalse(L1Messenger.failedMessages(hash));
        }
        assertFalse(actor.reverted());
    }
}

contract XDM_MinGasLimits_Reverts is XDM_MinGasLimits {
    function setUp() public override {
        // Do fail
        super.init(true);
    }

    /**
     * @custom:invariant A call to `relayMessage` should assign the message hash to the
     *                   `failedMessages` mapping if not enough gas is supplied to forward
     *                   `minGasLimit` to the target context or if there is not enough gas to
     *                   complete execution of `relayMessage` after the target context's execution
     *                   is finished.
     *
     * There are two minimum gas limits here:
     *
     * - The outer min gas limit is for the call from the `MorphPortal` to the
     * `L1CrossDomainMessenger`,  and it can be retrieved by calling the xdm's `baseGas` function
     * with the `message` and inner limit.
     *
     * - The inner min gas limit is for the call from the `L1CrossDomainMessenger` to the target
     * contract.
     */
    function invariant_minGasLimits() external {
        uint256 length = actor.numHashes();
        for (uint256 i = 0; i < length; ++i) {
            bytes32 hash = actor.hashes(i);
            // The message hash is not set in the successfulMessages mapping
            assertFalse(L1Messenger.successfulMessages(hash));
            // The message hash is set in the failedMessages mapping
            assertTrue(L1Messenger.failedMessages(hash));
        }
        assertFalse(actor.reverted());
    }
}
