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
  doOwnershipTransfer
} from "../src/deploy-utils";
import { ethers } from 'ethers'
import assert from 'assert'
import {
  ImplStorageName,
  ProxyStorageName,
  ContractFactoryName,
} from "./types"

const uint128Max = ethers.BigNumber.from('0xffffffffffffffffffffffffffffffff')

export const SystemDictatorSteps1 = async (
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
  const L1StandardBridgeProxyWithSigner = new ethers.Contract(
    L1StandardBridgeProxyAddress,
    L1StandardBridgeProxyFactory.interface,
    deployer,
  )

  const L1ERC721BridgeProxyAddress = getContractAddressByName(path, ProxyStorageName.L1ERC721BridgeProxyStroageName)
  const L1ERC721BridgeProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.DefaultProxy )
  const L1ERC721BridgeProxy = new ethers.Contract(
    L1ERC721BridgeProxyAddress,
    L1ERC721BridgeProxyFactory.interface,
    deployer.provider,
  )
  const L1ERC721BridgeProxyWithSigner = new ethers.Contract(
    L1ERC721BridgeProxyAddress,
    L1ERC721BridgeProxyFactory.interface,
    deployer,
  )

  const SystemConfigProxyAddress = getContractAddressByName(path, ProxyStorageName.SystemConfigProxyStorageName)
  const SystemConfigProxyFactory = await hre.ethers.getContractFactory(ContractFactoryName.SystemConfig)
  const SystemConfigProxy = new ethers.Contract(
    SystemConfigProxyAddress,
    SystemConfigProxyFactory.interface,
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

  // Transfer ownership of the ProxyAdmin to the SystemDictator.
  if ((await ProxyAdmin.owner()).toLocaleLowerCase() !== SystemDictator.address.toLocaleLowerCase()) {
    await doOwnershipTransfer({
      isLiveDeployer,
      proxy: ProxyAdmin,
      name: 'ProxyAdmin',
      transferFunc: 'transferOwnership',
      dictator: SystemDictator,
    })
  }


  // We don't need to transfer proxy addresses if we're already beyond the proxy transfer step.
  const needsProxyTransfer =
    (await SystemDictator.currentStep()) <=
    (await SystemDictator.PROXY_TRANSFER_STEP())


  // Transfer ownership of the AddressManager to SystemDictator.
  if (
    needsProxyTransfer &&
    (await AddressManager.owner()).toLocaleLowerCase() !== SystemDictator.address.toLocaleLowerCase()
  ) {
    await doOwnershipTransfer({
      isLiveDeployer,
      proxy: AddressManager,
      name: 'AddressManager',
      transferFunc: 'transferOwnership',
      dictator: SystemDictator,
    })
  } else {
    console.log(`AddressManager already owned by the SystemDictator`)
  }

  // Transfer ownership of the L1StandardBridge (proxy) to SystemDictator.
  if (
    needsProxyTransfer &&
    (await L1StandardBridgeProxy.callStatic.getOwner({
      from: ethers.constants.AddressZero,
    })).toLocaleLowerCase() !== SystemDictator.address.toLocaleLowerCase()
  ) {
    await doOwnershipTransfer({
      isLiveDeployer,
      proxy: L1StandardBridgeProxyWithSigner,
      name: 'L1StandardBridgeProxy',
      transferFunc: 'setOwner',
      dictator: SystemDictator,
    })
  } else {
    console.log(`L1StandardBridge already owned by MSD`)
  }


  // Transfer ownership of the L1ERC721Bridge (proxy) to SystemDictator.
  if (
    needsProxyTransfer &&
    (await L1ERC721BridgeProxy.callStatic.admin({
      from: ethers.constants.AddressZero,
    })).toLocaleLowerCase() !== SystemDictator.address.toLocaleLowerCase()
  ) {
    await doOwnershipTransfer({
      isLiveDeployer,
      proxy: L1ERC721BridgeProxyWithSigner,
      name: 'L1ERC721BridgeProxy',
      transferFunc: 'changeAdmin',
      dictator: SystemDictator,
    })
  } else {
    console.log(`L1ERC721Bridge already owned by MSD`)
  }

  // Wait for the ownership transfers to complete before continuing.
  await awaitCondition(
    async (): Promise<boolean> => {
      const proxyAdminOwner = await ProxyAdmin.owner()
      const addressManagerOwner = await AddressManager.owner()
      const l1StandardBridgeOwner =
        await L1StandardBridgeProxy.callStatic.getOwner({
          from: ethers.constants.AddressZero,
        })
      const l1Erc721BridgeOwner = await L1ERC721BridgeProxy.callStatic.admin({
        from: ethers.constants.AddressZero,
      })

      return (
        proxyAdminOwner.toLowerCase() === SystemDictator.address &&
        addressManagerOwner.toLowerCase() === SystemDictator.address &&
        l1StandardBridgeOwner.toLowerCase() === SystemDictator.address &&
        l1Erc721BridgeOwner.toLowerCase() === SystemDictator.address
      )
    },
    5000,
    1000
  )

  await doPhase({
    isLiveDeployer,
    SystemDictator,
    phase: 1,
    message: `
      Phase 1 includes the following steps:

      Step 1 will configure the ProxyAdmin contract, you can safely execute this step at any time
      without impacting the functionality of the rest of the system.

      Step 2 will stop deposits and withdrawals via the L1CrossDomainMessenger and will stop the
      DTL from syncing new deposits via the CTC, effectively shutting down the legacy system. Once
      this step has been executed, you should immediately begin the L2 migration process. If you
      need to restart the system, run exit1() followed by finalize().
    `,
    checks: async () => {
      // Step 1 checks
      await assertContractVariable(
        ProxyAdmin,
        'addressManager',
        AddressManager.address
      )
      assert(
        (await ProxyAdmin.implementationName(
          getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
        )) === 'L1CrossDomainMessenger'
      )
      assert(
        (await ProxyAdmin.proxyType(
          getContractAddressByName(path, ProxyStorageName.L1CrossDomainMessengerProxyStroageName)
        )) === 2
      )
      assert(
        (await ProxyAdmin.proxyType(
          getContractAddressByName(path, ProxyStorageName.L1StandardBridgeProxyStroageName)
        )) === 1
      )

      // Check the SystemConfig was initialized properly.
      await assertContractVariable(
        SystemConfigProxy,
        'owner',
        configTmp.finalSystemOwner
      )
      await assertContractVariable(
        SystemConfigProxy,
        'overhead',
        configTmp.gasPriceOracleOverhead
      )
      await assertContractVariable(
        SystemConfigProxy,
        'scalar',
        configTmp.gasPriceOracleScalar
      )
      await assertContractVariable(
        SystemConfigProxy,
        'batcherHash',
        ethers.utils.hexZeroPad(
          configTmp.batchSenderAddress.toLowerCase(),
          32
        )
      )
      await assertContractVariable(
        SystemConfigProxy,
        'gasLimit',
        configTmp.l2GenesisBlockGasLimit
      )

      const config = await SystemConfigProxy.resourceConfig()
      assert(config.maxResourceLimit === 20_000_000)
      assert(config.elasticityMultiplier === 10)
      assert(config.baseFeeMaxChangeDenominator === 8)
      assert(config.systemTxMaxGas === 1_000_000)
      assert(ethers.utils.parseUnits('1', 'gwei').eq(config.minimumBaseFee))
      assert(config.maximumBaseFee.eq(uint128Max))

      // Step 2 checks
      const messenger = await AddressManager.getAddress(
        'L1CrossDomainMessenger'
      )
      assert(messenger === ethers.constants.AddressZero)
    },
  })

  return ''
}

export default SystemDictatorSteps1

