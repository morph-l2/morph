# Shadow Proving
The contract required for running shadow-proving mode for morph-l2.

### Init
```
mkdir -p lib  
git clone https://github.com/foundry-rs/forge-std.git lib/forge-std
```

### Build
```
forge build  
```

## deploy

EvmVerifier:
```
forge create --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8545 src/libs/EvmVerifier.sol:EvmVerifier --constructor-args 0x0063251fce15999146d8c71f04770e5ce66a929f0a5d0c6d9772fa39d41b6654
```

ShadowRollup
```
forge create --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8545 src/ShadowRollup.sol:ShadowRollup --constructor-args 53077 0x5FbDB2315678afecb367f032d93F642f64180aa3
```


### Test
```
forge test
```
