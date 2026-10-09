const config = {
    // Global configuration
    contractAdmin: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    l1FeeVaultRecipient: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',

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
    finalizationPeriodSeconds: 10,
    rollupProofWindow: 86400,
    proofRewardPercent: 70,
    rollupDelayPeriod: 86400,

    // challenge config
    batchSubmitterAddresses: ['0x70997970C51812dc3A010C7d01b50e0d17dc79C8'],
    rollupChallenger: '0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65',
    // genesis config
    // Initialize with deployment-config.json generated for this genesis.
    batchHeader: '',

    // submitter initialize config
    submitterOwner: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    submitterChallengeDeposit: 1, // 1 ether
    submitterMinimumStake: 1, // 1 ether
    submitterRewardPercentage: 100,
    firstSequencerAddress: '',
}

export default config
