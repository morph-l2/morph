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
forge create --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8545 prover/contracts/src/EvmVerifier.sol:EvmVerifier --constructor-args 0x00c684101523a06415be6d72d011eebaf465b4c1da4e3b18ab052a29d7ac6e88
```
