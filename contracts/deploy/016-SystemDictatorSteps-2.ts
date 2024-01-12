import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { HardhatRuntimeEnvironment } from 'hardhat/types';
import {
  assertContractVariable,
  getContractAddressByName,
  awaitCondition,
  doPhase,
  liveDeployer,
  isStartOfPhase,
  printJsonTransaction,
  printCastCommand,
  isStep
} from "../src/deploy-utils";
import { ethers } from 'ethers'
import assert from 'assert'
import {
  ImplStorageName,
  ProxyStorageName,
  ContractFactoryName,
} from "./types"

export const SystemDictatorSteps2 = async (
  hre: HardhatRuntimeEnvironment,
  path: string,
  deployer: any,
  configTmp: any
): Promise<string> => {
  // Load the contracts we need to interact with.
  const SystemDictatorProxyAddress = getContractAddressByName(path, ProxyStorageName.SystemDictatorProxyStorageName)
  const SystemDictatorFactory = await hre.ethers.getContractFactory(ContractFactoryName.SystemDictator)
  const SystemDictator = new ethers.Contract(
    SystemDictatorProxyAddress,
    SystemDictatorFactory.interface,
    deployer,
  )

  const ProxyAdminAddress = getContractAddressByName(path, ImplStorageName.ProxyAdmin)
  const ProxyAdminFactory = await hre.ethers.getContractFactory(ContractFactoryName.ProxyAdmin)
  const ProxyAdmin = new ethers.Contract(
    ProxyAdminAddress,
    ProxyAdminFactory.interface,
    deployer,
  )

  const AddressManagerAddress = getContractAddressByName(path, ImplStorageName.AddressManager)
  const AddressManagernFactory = await hre.ethers.getContractFactory(ContractFactoryName.AddressManager)
  const AddressManager = new ethers.Contract(
    AddressManagerAddress,
    AddressManagernFactory.interface,
    deployer,
  )

  const L1StandardBridgeProxyAddress = getContractAddressByName(path, ProxyStorageName.L1StandardBridgeProxyStroageName)
  const L1StandardBridgeProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1ChugSplashProxy)
  const L1StandardBridgeProxy = new ethers.Contract(
    L1StandardBridgeProxyAddress,
    L1StandardBridgeProxyFactory.interface,
    deployer.provider,
  )
  const L1StandardBridgeFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1StandardBridge)
  const L1StandardBridge = new ethers.Contract(
    L1StandardBridgeProxyAddress,
    L1StandardBridgeFactory.interface,
    deployer,
  )

  const L1ERC721BridgeProxyAddress = getContractAddressByName(path, ProxyStorageName.L1ERC721BridgeProxyStroageName)
  const L1ERC721BridgeProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy)
  const L1ERC721BridgeProxy = new ethers.Contract(
    L1ERC721BridgeProxyAddress,
    L1ERC721BridgeProxyFactory.interface,
    deployer.provider,
  )
  const L1ERC721BridgeFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1ERC721Bridge)
  const L1ERC721Bridge = new ethers.Contract(
    L1ERC721BridgeProxyAddress,
    L1ERC721BridgeFactory.interface,
    deployer.provider,
  )
  const L1CrossDomainMessengerProxyAddress = getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
  const L1CrossDomainMessengerFactory = await hre.ethers.getContractFactory(ContractFactoryName.L1CrossDomainMessenger)
  const L1CrossDomainMessenger = new ethers.Contract(
    L1CrossDomainMessengerProxyAddress,
    L1CrossDomainMessengerFactory.interface,
    deployer,
  )

  const MorphPortalProxyAddress = getContractAddressByName(path, ProxyStorageName.MorphPortalProxyStroageName)
  const MorphPortalFactory = await hre.ethers.getContractFactory(ContractFactoryName.MorphPortal)
  const MorphPortal = new ethers.Contract(
    MorphPortalProxyAddress,
    MorphPortalFactory.interface,
    deployer,
  )

  const MorphMintableERC20FactoryProxyAddress = getContractAddressByName(path, ProxyStorageName.MorphMintableERC20FactoryProxyStroageName)
  const MorphMintableERC20FactoryFactory = await hre.ethers.getContractFactory(ContractFactoryName.MorphMintableERC20Factory)
  const MorphMintableERC20Factory = new ethers.Contract(
    MorphMintableERC20FactoryProxyAddress,
    MorphMintableERC20FactoryFactory.interface,
    deployer,
  )

  // If we have the key for the controller then we don't need to wait for external txns.
  // Set the DISABLE_LIVE_DEPLOYER=true in the env to ensure the script will pause to simulate scenarios
  // where the controller is not the deployer.
  const isLiveDeployer = await liveDeployer({
    disabled: process.env.DISABLE_LIVE_DEPLOYER,
    controller: configTmp.controller,
    deployer: deployer,
  })


  // Make sure the dynamic system configuration has been set.
  if (
    (await isStartOfPhase(SystemDictator, 2))
    // &&!(await SystemDictator.dynamicConfigSet())
  ) {
    console.log(`
      You must now set the dynamic L2OutputOracle configuration by calling the function
      updateL2OutputOracleDynamicConfig. You will need to provide the
      l2OutputOracleStartingBlockNumber and the l2OutputOracleStartingTimestamp which can both be
      found by querying the last finalized block in the L2 node.
    `)
  }


  await doPhase({
    isLiveDeployer,
    SystemDictator,
    phase: 2,
    message: `
      Phase 2 includes the following steps:

      Step 3 will clear out some legacy state from the AddressManager. Once you execute this step,
      you WILL NOT BE ABLE TO RESTART THE SYSTEM using exit1(). You should confirm that the L2
      system is entirely operational before executing this step.

      Step 4 will transfer ownership of the AddressManager and L1StandardBridge to the ProxyAdmin.

      Step 5 will initialize all Bedrock contracts. After this step is executed, the MorphPortal
      will be open for deposits but withdrawals will be paused if deploying a production network.
      The Proposer will also be able to submit L2 outputs to the L2OutputOracle.

      Lastly the finalize step will be executed. This will transfer ownership of the ProxyAdmin to
      the final system owner as specified in the deployment configuration.
    `,
    checks: async () => {
      // Step 3 checks
      const deads = [
        'CanonicalTransactionChain',
        'L2CrossDomainMessenger',
        'DecompressionPrecompileAddress',
        'Sequencer',
        'Proposer',
        'ChainStorageContainer-CTC-batches',
        'ChainStorageContainer-CTC-queue',
        'CanonicalTransactionChain',
        'StateCommitmentChain',
        'BondManager',
        'ExecutionManager',
        'FraudVerifier',
        'StateManagerFactory',
        'StateTransitionerFactory',
        'SafetyChecker',
        'L1MultiMessageRelayer',
        'BondManager',
      ]
      for (const dead of deads) {
        const addr = await AddressManager.getAddress(dead)
        assert(addr === ethers.constants.AddressZero)
      }

      // Step 4 checks
      await assertContractVariable(AddressManager, 'owner', ProxyAdmin.address)

      assert(
        (await L1StandardBridgeProxy.callStatic.getOwner({
          from: ethers.constants.AddressZero,
        })).toLowerCase() === ProxyAdmin.address.toLowerCase()
      )

      assert(
        (await L1ERC721BridgeProxy.callStatic.admin({
          from: ProxyAdmin.address,
        })).toLowerCase() === ProxyAdmin.address.toLowerCase()
      )

      // Check MorphPortal was initialized properly.
      await assertContractVariable(
        MorphPortal,
        'l2Sender',
        '0x000000000000000000000000000000000000dEaD'
      )
      const resourceParams = await MorphPortal.params()
      assert(
        resourceParams.prevBaseFee.eq(ethers.utils.parseUnits('1', 'gwei')),
        `MorphPortal was not initialized with the correct initial base fee`
      )
      assert(
        resourceParams.prevBoughtGas.eq(0),
        `MorphPortal was not initialized with the correct initial bought gas`
      )
      assert(
        !resourceParams.prevBlockNum.eq(0),
        `MorphPortal was not initialized with the correct initial block number`
      )
      assert(
        (await hre.ethers.provider.getBalance(L1StandardBridge.address)).eq(0)
      )

      if (isLiveDeployer) {
        await assertContractVariable(MorphPortal, 'paused', false)
      } else {
        await assertContractVariable(MorphPortal, 'paused', true)
      }

      await assertContractVariable(
        MorphPortal,
        'l1Messenger',
        L1CrossDomainMessenger.address
      )
      // Check L1CrossDomainMessenger was initialized properly.
      try {
        await L1CrossDomainMessenger.xDomainMessageSender()
        assert(false, `L1CrossDomainMessenger was not initialized properly`)
      } catch (err) {
        // Expected.
      }

      // Check L1StandardBridge was initialized properly.
      await assertContractVariable(
        L1StandardBridge,
        'messenger',
        L1CrossDomainMessenger.address
      )
      assert(
        (await hre.ethers.provider.getBalance(L1StandardBridge.address)).eq(0)
      )

      // Check MorphMintableERC20Factory was initialized properly.
      await assertContractVariable(
        MorphMintableERC20Factory,
        'BRIDGE',
        L1StandardBridge.address
      )

      // Check L1ERC721Bridge was initialized properly.
      await assertContractVariable(
        L1ERC721Bridge,
        'messenger',
        L1CrossDomainMessenger.address
      )

      // finalize checks
      await assertContractVariable(
        ProxyAdmin,
        'owner',
        configTmp.finalSystemOwner
      )
    },
  })


  // Step 6 unpauses the MorphPortal.
  if (await isStep(SystemDictator, 6)) {
    console.log(`
      Unpause the MorphPortal. The GUARDIAN account should be used. In practice
      this is the multisig. In test networks, the MorphPortal is initialized
      without being paused.
    `)

    if (isLiveDeployer) {
      console.log('WARNING: MorphPortal configured to not be paused')
      console.log('This should only happen for test environments')
      await assertContractVariable(MorphPortal, 'paused', false)
    } else {
      const tx = await MorphPortal.populateTransaction.unpause()
      console.log(`Please unpause the MorphPortal...`)
      console.log(`MorphPortal address: ${MorphPortal.address}`)
      printJsonTransaction(tx)
      printCastCommand(tx)
      // await printTenderlySimulationLink(SystemDictator.provider, tx)
    }

    await awaitCondition(
      async () => {
        const paused = await MorphPortal.paused()
        return !paused
      },
      5000,
      1000
    )

    await assertContractVariable(MorphPortal, 'paused', false)

    await awaitCondition(
      async () => {
        return SystemDictator.finalized()
      },
      5000,
      1000
    )
  }
  return ''
}

export default SystemDictatorSteps2

