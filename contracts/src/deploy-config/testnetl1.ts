const config = {
    finalSystemOwner: '0xD77c708607e72f520c2841E359cF54fca75d9C04',
    controller: '0xD77c708607e72f520c2841E359cF54fca75d9C04',
    portalGuardian: '0xD77c708607e72f520c2841E359cF54fca75d9C04',
    proxyAdminOwner: '0xD77c708607e72f520c2841E359cF54fca75d9C04',

    l1StartingBlockTag:
        '0x126e52a0cc0ae18948f567ee9443f4a8f0db67c437706e35baee424eb314a0d0',
    l1ChainID: 900,
    l2ChainID: 2710,

    maxSequencerDrift: 300,
    sequencerWindowSize: 200,
    channelTimeout: 120,

    rollupMinDeposit: 1,
    rollupProofWindow: 100,
    rollupGenesisBlockNumber: 0,
    rollupProposer: '0xE70a4102e4caA3d9B8968FFb142E6A2ceFd22Ec3',
    rollupChallenger: '0xF730477971E88b3162ed7FD950Be63474975bdC3',
    rollupGenesisStateRoot: '',
    finalizationPeriodSeconds: 2,

    stakingMinDeposit: 1,
    stakingSequencerSize: 3,
    stakingLockNumber: 3,
    /**
     *  Only used for SystemConfig
     *  --------from--------
    */
    p2pSequencerAddress: '0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc',
    batchInboxAddress: '0xff00000000000000000000000000000000000010',
    batchSenderAddress: '0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC',

    baseFeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
    l1FeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
    sequencerFeeVaultRecipient: '0x90F79bf6EB2c4f870365E785982E1f101E93b906',
    /**
     * -----------to--------
     */
    governanceTokenName: 'Morph',
    governanceTokenSymbol: 'MORPH',
    governanceTokenOwner: '0xD77c708607e72f520c2841E359cF54fca75d9C04',

    l2GenesisBlockGasLimit: '0x1c9c380',
    l2GenesisBlockCoinbase: '0x4200000000000000000000000000000000000011',
    l2GenesisBlockBaseFeePerGas: '0x3B9ACA00',

    gasPriceOracleOverhead: 2100,
    gasPriceOracleScalar: 1000000,
    eip1559Denominator: 8,
    eip1559Elasticity: 2,

    l2GenesisRegolithTimeOffset: '0x0',
    l2BlockTime: 1,
}

export default config 
