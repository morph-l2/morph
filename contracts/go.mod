module morph-l2/contract

go 1.22

replace (
	github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20240429050506-03fd4c3e771d
	github.com/tendermint/tendermint => github.com/morph-l2/tendermint v0.2.0-beta.0.20240513090937-03bf2a578b48
)

require (
	github.com/iden3/go-iden3-crypto v0.0.16
	github.com/scroll-tech/go-ethereum v1.11.4
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
