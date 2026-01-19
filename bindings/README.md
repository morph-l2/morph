# morph-bindings

This package contains built go bindings of the smart contracts. It must be
updated after any changes to the smart contracts to ensure that the bindings are
up to date.

The bindings include the bytecode for each contract so that go based tests
can deploy the contracts. There are also `more` files that include the deployed
bytecode as well as the storage layout. These are used to dynamically set
bytecode and storage slots in state.

## Dependencies

- `abigen` version 1.10.25
- `make`

To check the version of `abigen`, run the command `abigen --version`.

## abigen

The `abigen` tool is part of `go-ethereum` and can be used to build go bindings
for smart contracts. It can be installed with go using the commands:

```bash
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make devtools
```

The geth docs for `abigen` can be found [here](https://geth.ethereum.org/docs/dapp/native-bindings).

## Installation and Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/morph-bindings.git
   cd morph-bindings
   ```
2. **Install Dependencies**
Ensure you have all required dependencies installed, including:
- `jq`
- `abigen`
- `solc`
- `make`
- `hardhat`

Install these as needed, following the official documentation for each tool.

3. **Compile Contracts**
Use `make` to compile the contracts and generate bindings. Run:
```bash
make all
 ```
## Structure and Contents

### Key Files

- **types.go**: Defines the main data structures for compiler input, settings, output, and storage layout, allowing smooth integration between the Go application and Ethereum contracts.
- **Makefile**: Automates tasks including compilation, generating bindings, and cleaning up build files.
- **compile.sh**: Script to compile contracts using `hardhat` or `forge`.
- **gen_bindings.sh**: Script to generate contract bindings using `abigen` and customize dependencies.

### Contract Bindings

The bindings provide access to the following types of contracts:
- **L1 Contracts**: Staking, gateways, and cross-domain messaging between L1 and L2.
- **L2 Contracts**: Fee management, staking, and distribution, along with specific gateways for assets like ERC-20, ERC-721, and ERC-1155 tokens.

### Modules

- **Storage Layouts**: Stores layout details for each contract, used for setting up state storage dynamically during testing.
- **Bytecode Management**: Includes both deployed and undeployed bytecode for each contract, allowing for flexible deployment.

