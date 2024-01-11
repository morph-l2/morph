import { ethers } from 'ethers'

/**
 * Core required deployment configuration.
 */
interface RequiredDeployConfig {
  /**
   * Number of confirmations to wait when deploying contracts.
   */
  numDeployConfirmations?: number

  /**
   * Address that will own the entire system on L1 when the deploy is complete.
   */
  finalSystemOwner?: string

  /**
   * Address that is deployed as the GUARDIAN in the MorphPortal. Has the
   * ability to pause withdrawals.
   */
  portalGuardian: string

  /**
   * Address that will own the entire system on L1 during the deployment process. This address will
   * not own the system after the deployment is complete, ownership will be transferred to the
   * final system owner.
   */
  controller?: string

  /**
   * The L2 genesis script uses this to fill the storage of the L1Block info predeploy. The rollup
   * config script uses this to fill the L1 genesis info for the rollup. The Output oracle deploy
   * script may use it if the L2 starting timestamp is undefined, assuming the L2 genesis is set up
   * with this.
   */
  l1StartingBlockTag: string

  /**
   * Chain ID for the L1 network.
   */
  l1ChainID: number

  /**
   * Chain ID for the L2 network.
   */
  l2ChainID: number

  /**
   * Number of seconds in between each L2 block.
   */
  l2BlockTime: number

  /**
   * Sequencer batches may not be more than maxSequencerDrift seconds after the L1 timestamp of the
   * end of the sequencing window end.
   */
  maxSequencerDrift: number

  /**
   * Number of L1 blocks per sequencing window.
   */
  sequencerWindowSize: number

  /**
   * Number of L1 blocks that a frame stays valid when included in L1.
   */
  channelTimeout: number

  /**
   * Address of the key the sequencer uses to sign blocks on the P2P layer.
   */
  p2pSequencerAddress: string

  /**
   * L1 address that batches are sent to.
   */
  batchInboxAddress: string

  /**
   * Acceptable batch-sender address, to filter transactions going into the batchInboxAddress on L1 for data.
   * Warning: this address is hardcoded now, but is intended to become governed via L1.
   */
  batchSenderAddress: string

  /**
   * Output Oracle submission interval in L2 blocks.
   */
  l2OutputOracleSubmissionInterval: number

  /**
   * Starting block number for the output oracle.
   * Must be greater than or equal to the first Bedrock block. The first L2 output will correspond
   * to this value plus the submission interval.
   */
  l2OutputOracleStartingBlockNumber?: number

  /**
   * Starting timestamp for the output oracle.
   * MUST be the same as the timestamp of the L2OO start block.
   */
  l2OutputOracleStartingTimestamp?: number

  /**
   * Address of the L2 output oracle proposer.
   */
  l2OutputOracleProposer: string

  /**
   * Address of the L2 output oracle challenger.
   */
  l2OutputOracleChallenger: string

  /**
   * ERC20 symbol used for the L2 GovernanceToken.
   */
  governanceTokenSymbol: string

  /**
   * ERC20 name used for the L2 GovernanceToken.
   */
  governanceTokenName: string

  /**
   * Owner of the L2 GovernanceToken. Has mint/burn capability.
   */
  governanceTokenOwner: string

  /**
   * Output finalization period in seconds.
   */
  finalizationPeriodSeconds: number

  /**
   * Owner of the ProxyAdmin contract.
   */
  proxyAdminOwner: string

  /**
   * L1 address which receives the base fee for the L2 network.
   */
  baseFeeVaultRecipient: string

  /**
   * L1 address which receives data fees for the L2 network.
   */
  l1FeeVaultRecipient: string

  /**
   * L1 address which receives tip fees for the L2 network.
   */
  sequencerFeeVaultRecipient: string
}

/**
 * Optional deployment configuration when spinning up an L1 network as part of the deployment.
 */
interface OptionalL1DeployConfig {
  cliqueSignerAddress: string
  l1BlockTime: number
  l1GenesisBlockNonce: string
  l1GenesisBlockGasLimit: string
  l1GenesisBlockDifficulty: string
  l1GenesisBlockMixHash: string
  l1GenesisBlockCoinbase: string
  l1GenesisBlockNumber: string
  l1GenesisBlockGasUsed: string
  l1GenesisBlockParentHash: string
  l1GenesisBlockBaseFeePerGas: string
}

/**
 * Optional deployment configuration when spinning up an L2 network as part of the deployment.
 */
interface OptionalL2DeployConfig {
  l2GenesisBlockNonce: string
  l2GenesisBlockGasLimit: string
  l2GenesisBlockDifficulty: string
  l2GenesisBlockMixHash: string
  l2GenesisBlockNumber: string
  l2GenesisBlockGasUsed: string
  l2GenesisBlockParentHash: string
  l2GenesisBlockBaseFeePerGas: string
  l2GenesisBlockCoinbase: string
  l2GenesisRegolithTimeOffset: string
  eip1559Denominator: number
  eip1559Elasticity: number
  gasPriceOracleOverhead: number
  gasPriceOracleScalar: number
}

/**
 * Full deployment configuration.
 */
export type DeployConfig = RequiredDeployConfig &
  Partial<OptionalL1DeployConfig> &
  Partial<OptionalL2DeployConfig>