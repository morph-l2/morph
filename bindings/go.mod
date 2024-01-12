module github.com/morph-l2/bindings

go 1.20

require github.com/scroll-tech/go-ethereum v1.11.4

require (
	github.com/VictoriaMetrics/fastcache v1.12.2 // indirect
	github.com/btcsuite/btcd v0.23.4 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/holiman/uint256 v1.2.4 // indirect
	github.com/huin/goupnp v1.3.0 // indirect
	github.com/iden3/go-iden3-crypto v0.0.15 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/prometheus/tsdb v0.10.0 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/rjeczalik/notify v0.9.3 // indirect
	github.com/scroll-tech/zktrie v0.7.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/status-im/keycard-go v0.3.2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220614013038-64ee5596c38a // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.20.1-beta
	github.com/scroll-tech/go-ethereum => ../l2geth
)
