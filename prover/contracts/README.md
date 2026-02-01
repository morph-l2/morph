## Contracts
Morph EvmVerifier.
These contracts are used to verify batch proofs of the EVM.

### Init

```
mkdir -p lib  
git clone https://github.com/foundry-rs/forge-std.git lib/forge-std
```

### Test

```shell
$ forge test
```

### Deploy
```
forge create --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8545 prover/contracts/src/EvmVerifier.sol:EvmVerifier --constructor-args 0x0063251fce15999146d8c71f04770e5ce66a929f0a5d0c6d9772fa39d41b6654
```