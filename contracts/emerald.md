# hoodi
## L2TokenRegistry
use L2 owner's private_key
```
yarn hardhat deploy-l2-token-registry --proxyadmin 0x530000000000000000000000000000000000000B --proxy 0x5300000000000000000000000000000000000021 --owner 0x2910beed7B7F86f4bE9907b2e91FA0a216a67f2E --network hoodil2
```

```
Upgrading existing proxy at 0x5300000000000000000000000000000000000021
Proxy upgraded to new implementation: 0x75aBfc927b6e78Df695b75915ef637f3C0f23941
Proxy already initialized, skipping initialize() call.

L2TokenRegistry proxy address: 0x5300000000000000000000000000000000000021
L2TokenRegistry proxy owner: 0x2910beed7B7F86f4bE9907b2e91FA0a216a67f2E
L2TokenRegistry allowListEnabled: true
```

## ZK Verifier
use L1 owner's private_key
```
yarn hardhat upgradeVerifier --rollupversion 1 --startbatchindex 6860 --multipleversionrollupverifier 0x4a6d566a55b5dad41f133571ad53cf52b00a7351 --network hoodi
```

```
ZkEvmVerifierV1Contract: 0xe253850969da0c96cfdb9a33f9b7c125674c4d8c ; TX_HASH: 0x7e4139da2e399bef8a5dc09ab30f918456dc8823243a5829aa0fcb2e4e0c30f5
BLOCK_NUMBER: 1881494
receipt status : 1
upgrade verifier successfully, verifier: 0xe253850969da0c96cfdb9a33f9b7c125674c4d8c
```

https://hoodi.etherscan.io/address/0xe253850969da0c96cfdb9a33f9b7c125674c4d8c#readContract
programVkey successfully changed to `0x00ad538a51c761c06f5075d11f3ee64d5d00c272a741ccf098e1d9f062fee13d`

matched with pr `https://github.com/morph-l2/morph/commit/96c7524ae8d6d9717191201a20afe45f046d8611`


Update Verifier Tx: 
https://hoodi.etherscan.io/tx/0x322a247619f71930a37d18e64d4b71bdfb5751de7e83b718c778f6473fb75475

batch 6860 rollup successfully