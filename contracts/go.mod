module github.com/morph-l2/contract

go 1.20

require (
	github.com/iden3/go-iden3-crypto v0.0.15
	github.com/scroll-tech/go-ethereum v1.11.4
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.20.1-beta
	github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20240429050506-03fd4c3e771d
	github.com/tendermint/tendermint => github.com/morph-l2/tendermint v0.2.0-beta
)

require (
	github.com/btcsuite/btcd v0.20.1-beta // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
)
