# hoodi
version：release/0.4.x 96c7524ae8d6d9717191201a20afe45f046d8611
## L2TokenRegistry
use L2 owner's private_key
```bash
yarn hardhat deploy-l2-token-registry --proxyadmin 0x530000000000000000000000000000000000000B --proxy 0x5300000000000000000000000000000000000021 --owner 0x2910beed7B7F86f4bE9907b2e91FA0a216a67f2E --network hoodil2
```

```bash
Upgrading existing proxy at 0x5300000000000000000000000000000000000021
Proxy upgraded to new implementation: 0x75aBfc927b6e78Df695b75915ef637f3C0f23941
Proxy already initialized, skipping initialize() call.

L2TokenRegistry proxy address: 0x5300000000000000000000000000000000000021
L2TokenRegistry proxy owner: 0x2910beed7B7F86f4bE9907b2e91FA0a216a67f2E
L2TokenRegistry allowListEnabled: true
```

## ZK Verifier
use L1 owner's private_key
```bash
yarn hardhat upgradeVerifier --rollupversion 1 --startbatchindex 6860 --multipleversionrollupverifier 0x4a6d566a55b5dad41f133571ad53cf52b00a7351 --network hoodi
```

```bash
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


# mainnet
version：release/0.4.x 96c7524ae8d6d9717191201a20afe45f046d8611
## L2TokenRegistry
L2TokenRegistry impl deployed at 0x5ff102a4A4Ce2040288a797CE4CCCa85eE1E2d70

initialize：
imp、proxyadmin、owner：0xEF6ae9F469aC85954c51829EBbf54D88A2ad1809
```js
        const proxy = await TransparentProxyFactory.deploy(
            tokenRegistry.address, //logic
            taskArgs.proxyadmin, //admin
            TokenRegistryFactory.interface.encodeFunctionData('initialize', [
                taskArgs.owner // owner
            ]) // data
        )
```
multi-signature: 
https://safe.morphl2.io/transactions/tx?safe=morph:0x72C5a62D1EC964DA5D0a955fE4BC709526255A19&id=multisig_0x72C5a62D1EC964DA5D0a955fE4BC709526255A19_0x06aa7d6c5eefb9074b37288277ec28b6d43ec1eb7202c68ef228a0d11cedcd35

proxy owner now：0xEF6ae9F469aC85954c51829EBbf54D88A2ad1809

register and active token：usdt usdc bgb
```bash
  cast send 0x5300000000000000000000000000000000000021 \
    --rpc-url https://rpc.morph.network  \
    --private-key $PRIVATE_KEY \
    "registerTokens(uint16[],address[],bytes32[],bool[],uint256[])" \
    "[1,2,3]" \
    "[0xc7D67A9cBB121b3b0b9c053DD9f469523243379A,0xe34c91815d7fc18A9e2148bcD4241d0a5848b693,0x55d1f1879969bdbB9960d269974564C58DBc3238]" \
    "[0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000000000000000000000000000,0x0000000000000000000000000000000000000000000000000000000000000000]" \
    "[false,false,false]" \
    "[1000000,1000000,1000000000000000000]"


  cast send 0x5300000000000000000000000000000000000021 \
    "batchUpdateTokenStatus(uint16[],bool[])" \
    "[1,2,3]" \
    "[true,true,true]" \
    --rpc-url https://rpc.morph.network --private-key $PRIVATE_KEY
```

set tokenPriceOracle(0x83a0ac6a1c2a634939b6137adde67b74acbce203) as updater：
```bash
  cast send 0x5300000000000000000000000000000000000021 \
    "setAllowList(address[],bool[])" \
    "[0x83a0ac6a1c2a634939b6137adde67b74acbce203]" \
    "[true]" \
    --rpc-url https://rpc.morph.network --private-key $PRIVATE_KEY
```

set balanceslot
```bash
cast send 0x5300000000000000000000000000000000000021 \
  --rpc-url https://rpc.morph.network \
  --private-key $PRIVATE_KEY \
  "updateTokenInfo(uint16,address,bytes32,bool,bool,uint256)" \
  1 \
  0xc7D67A9cBB121b3b0b9c053DD9f469523243379A \
  0x0000000000000000000000000000000000000000000000000000000000000033 \
  true \
  true \
  1000000
  
cast send 0x5300000000000000000000000000000000000021 \
  --rpc-url https://rpc.morph.network \
  --private-key $PRIVATE_KEY \
  "updateTokenInfo(uint16,address,bytes32,bool,bool,uint256)" \
  3 \
  0x55d1f1879969bdbB9960d269974564C58DBc3238 \
  0x0000000000000000000000000000000000000000000000000000000000000033 \
  true \
  true \
  1000000000000000000
```


transfer owner：



## ZK Verifier
ZkEvmVerifierV1Contract: 0x5ff102a4a4ce2040288a797ce4ccca85ee1e2d70 ; TX_HASH: 0x2fd545cec0a0381b1389899e123a6a2168a5f88955f183c65906dac761139f4d
BLOCK_NUMBER: 24172926

相关合约：
rollup：https://etherscan.io/address/0x759894ced0e6af42c26668076ffa84d02e3cef60
MultipleVersionRollupVerifier：https://etherscan.io/address/0x5d1584c27b4aD233283c6da1ca1B825d6f220EC1

操作完成后，转交ZkEvmVerifierV1Contract owner地址

