# Morph Smart Contracts

This package contains the smart contracts that compose the on-chain component.

## upgrade Contracts steps

- update contract code

- compile updated contract
`yarn hardhat compile`

- run deploy Impl contract, e.g.
`yarn hardhat upgradeRollupImpl --l1cdmproxyaddr 0xcf7ed3acca5a467e9e704c703e8d87f634fb0fc9 --network l1`

- run upgrade's command
`yarn hardhat upgradeProxy --proxyaddr 0xb7f8bc63bbcad18155201308c8f3540b07f84f5e --newimpladdr 0x9d4454B023096f34B160D6B654540c56A1F81688 --network l1`

- check upgraded proxy contract, e.g.
`yarn hardhat upgradeProxy --proxyaddr 0xb7f8bc63bbcad18155201308c8f3540b07f84f5e --newimpladdr 0x1429859428C0aBc9C2C47C8Ee9FBaf82cFA0F20f --network l1`

or

`yarn hardhat upgradeProxyWithProxyAdmin --proxyadminaddr 0x5FbDB2315678afecb367f032d93F642f64180aa3 \
--proxyaddr 0x8a791620dd6260079bf849dc5567adc3f2fdc318 --newimpladdr 0x4ed7c70f96b99c776995fb64377f0d4ab3b0e1c1 \
--network l1`

For the one-time Rollup cutover to Submitter, use the dedicated task instead of
the generic proxy upgrade. Stop legacy submissions, challenges, and revert
requests before running it. The task verifies the expected last committed batch,
requires no active challenge or revert request, and validates that the target is
an initialized Submitter proxy configured for this Rollup with valid owner and
staking parameters. The supplied operator must already be active. The task reads
the L2 chain ID from the existing Rollup, deploys the replacement with that value,
and calls `ProxyAdmin.upgradeAndCall(initialize4)` atomically:

`yarn hardhat upgradeRollupProxy --proxyadminaddr <proxy-admin> \
--rollupproxyaddr <rollup-proxy> --submitterproxyaddr <submitter-proxy> \
--submitteroperatoraddr <active-operator> --expectedcutoverbatchindex <K> \
--network <l1-network>`

For a new deployment, the `register` task reads `batchSubmitterAddresses` from
the selected deployment configuration and does not require operator private
keys. The owner authorizes each configured work or backup submitter and calls
`stake(address)` to fund its deployment-time stake. The operator address does
not need to send the staking transaction.
