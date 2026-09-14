const config = {
    // Global configuration
    l1FeeVaultRecipient: '0x7161DB99e6ffA72053f0817aBfaF2fDb4ab6ccC8',
    contractAdmin: '0x7161DB99e6ffA72053f0817aBfaF2fDb4ab6ccC8',

    l1WETHAddress: "",

    // chainID config
    l1ChainID: 900,
    l2ChainID: 53077,

    // L1MessageQueue config
    l1MessageQueueMaxGasLimit: 30000000,

    // gasPriceOracle config
    l2BaseFee: 0.1,  // Gwei

    // verify contract config
    programVkey: '0x00b450ec2a1b8dfba81ade90afbcc96842055548b814c991bb13bdca34980c63',
    // rollup contract config
    // initialize config
    finalizationPeriodSeconds: 600,
    rollupProofWindow: 86400,
    proofRewardPercent: 70,
    // QA requires an explicit delay in seconds; zero is rejected before broadcasting deployment transactions.
    rollupDelayPeriod: Number(process.env.QA_ROLLUP_DELAY_PERIOD || 0),
    // challenge config
    batchSubmitterAddresses: ['0x7161aE5B767Df97BAc34e5776311Ded4Cb01691d'],
    rollupChallenger: '0x71616250f7CAaa8a5DC295dc5851D6b1E49188a7',
    // genesis config
    batchHeader: '',

    // submitter initialize config
    submitterOwner: '0x7161DB99e6ffA72053f0817aBfaF2fDb4ab6ccC8',
    submitterChallengeDeposit: 1, // 1 ether
    submitterMinimumStake: 1, // 1 ether
    submitterRewardPercentage: 50,
    // L1Sequencer initialize config
    // Set the first sequencer through the deployment override file or firstSequencerAddress environment variable.
    firstSequencerAddress: '',

}

export default config
