module github.com/morph-l2/bindings

go 1.20

require github.com/scroll-tech/go-ethereum v1.11.4

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/bits-and-blooms/bitset v1.7.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.1 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/ethereum/c-kzg-4844/bindings/go v0.0.0-20230126171313-363c7d7593b4 // indirect
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/iden3/go-iden3-crypto v0.0.16 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/scroll-tech/zktrie v0.6.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	github.com/supranational/blst v0.3.11 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.20.1-beta
	github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20240429050506-03fd4c3e771d
	github.com/tendermint/tendermint => github.com/morph-l2/tendermint v0.2.0-beta
)
