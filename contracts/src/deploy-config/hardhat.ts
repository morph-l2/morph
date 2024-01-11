const config = {
  finalSystemOwner: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
  controller: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
  portalGuardian: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
  proxyAdminOwner: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',

  l1StartingBlockTag:
    '0x126e52a0cc0ae18948f567ee9443f4a8f0db67c437706e35baee424eb314a0d0',
  l1ChainID: 900,
  l2ChainID: 53077,


  maxSequencerDrift: 300,
  sequencerWindowSize: 200,
  channelTimeout: 120,

  rollupMinDeposit: 1,
  rollupProofWindow: 100,
  rollupGenesisBlockNumber: 0,
  rollupProposer: '0x70997970C51812dc3A010C7d01b50e0d17dc79C8',
  rollupChallenger: '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
  rollupGenesisStateRoot: '0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421',
  finalizationPeriodSeconds: 2,

  
  stakingMinDeposit: 1,
  stakingSequencerSize: 3,
  stakingLockNumber: 3,

  p2pSequencerAddress: '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc',
  batchInboxAddress: '0xff00000000000000000000000000000000000010',
  batchSenderAddress: '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',

  baseFeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
  l1FeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
  sequencerFeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',

  governanceTokenName: 'Morph',
  governanceTokenSymbol: 'MORPH',
  governanceTokenOwner: '0xBcd4042DE499D14e55001CcbB24a551F3b954096',

  l2GenesisBlockGasLimit: '0x1c9c380',
  l2GenesisBlockCoinbase: '0x4200000000000000000000000000000000000011',
  l2GenesisBlockBaseFeePerGas: '0x3B9ACA00',

  gasPriceOracleOverhead: 2100,
  gasPriceOracleScalar: 1000000,
  eip1559Denominator: 8,
  eip1559Elasticity: 2,

  l2GenesisRegolithTimeOffset: '0x0',

  l2BlockTime: 2,
}

export default config 