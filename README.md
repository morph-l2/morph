# Welcome to Morph

Welcome to the official GitHub repository for Morph: The Consumer Blockchain Revolution

Morph is an innovative force reshaping the consumer blockchain landscape for practical, everyday use. At the core of Morph is a revolutionary approach to Ethereum Layer 2 scalability, harnessing the power of advanced rollup technology. Our platform is uniquely designed to enhance the blockchain experience, making it more accessible, efficient, and user-friendly for both developers and consumers alike.

## Dive deeper into our vision and objectives here

1. [What is & Why Responsive Validity Proof?](https://docs.morphl2.io/docs/how-morph-works/optimistic-zkevm/#what-is-rvp)

2. [How Does RVP Run in Morph?](https://medium.com/@morphlayer2/how-does-rvp-run-in-morph-6025233a21cc)

## Development

### Setting Up Local Development Network

This repository supports launching a local private Morph network for development and debugging purposes.

#### Start the Development Network

```bash
make devnet-up
```

This command performs the following steps:
1. Initializes and updates the go-ethereum submodule to the specified version
2. Builds the go-ubuntu-builder Docker image (if not already present)
3. Compiles all necessary components (L1 execution layer, consensus layer, L2 node, and services)
4. Generates genesis configurations for both L1 and L2 networks
5. Deploys smart contracts to the L1 network
6. Starts a 4-node Morph network with associated services

**Note:** The initial startup may take considerable time due to extensive building operations.

#### Clean Up the Network

To completely remove the development network including Docker images:

```bash
make devnet-clean
```

To clean up only the data and build artifacts while preserving Docker images:

```bash
make devnet-clean-build
```

### Managing Dependencies

#### Update Module Dependencies

To download or update dependencies for all modules in this monorepo:

```bash
make update
```

This command synchronizes the Go workspace and updates all module dependencies according to the versions specified in the Makefile:
- `ETHEREUM_TARGET_VERSION`: Specifies the go-ethereum dependency version
- `TENDERMINT_TARGET_VERSION`: Specifies the tendermint dependency version

#### Update Submodules

To update the go-ethereum submodule to the target version:

```bash
make submodules
```

This command updates the go-ethereum submodule to the commit/tag specified by `ETHEREUM_SUBMODULE_COMMIT_OR_TAG` in the Makefile.

### Additional Commands

- `make lint`: Run linters for both Solidity and Go code
- `make fmt`: Format Solidity and Go code
- `make bindings`: Generate Go bindings from smart contracts
- `make geth`: Build the geth binary from the go-ethereum submodule

## Learn more

Website: https://www.morphl2.io/

X (formerly Twitter): https://x.com/Morphl2

Medium: [Morph – Medium](https://medium.com/@morphlayer2)

Telegram: https://t.me/morphoffical

Discord: https://discord.com/invite/MorphLayer

Gmail: info@morphl2.io
