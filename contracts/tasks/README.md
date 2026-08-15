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
the generic proxy upgrade. It checks that committed batches are fully finalized,
that no challenge or revert request is pending, and then calls
`ProxyAdmin.upgradeAndCall(initialize4)` atomically:

`yarn hardhat upgradeRollupProxy --proxyadminaddr <proxy-admin> \
--rollupproxyaddr <rollup-proxy> --submitterproxyaddr <submitter-proxy> \
--chainid <l2-chain-id> --network <l1-network>`

For a fresh deployment, `batchSubmitterPks` must be a JSON array whose order
matches `batchSubmitterAddresses` in the selected deploy config. The `register`
task owner-authorizes and stakes every configured work or backup submitter.
