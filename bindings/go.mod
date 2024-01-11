module github.com/morph-l2/bindings

go 1.19

require github.com/scroll-tech/go-ethereum v1.11.4

require (
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/btcsuite/btcd v0.23.3 // indirect
	github.com/deckarep/golang-set v0.0.0-20180603214616-504e848d77ea // indirect
	github.com/go-ole/go-ole v1.2.1 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/google/uuid v1.1.5 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/iden3/go-iden3-crypto v0.0.12 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/scroll-tech/zktrie v0.6.0 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
)

replace (
	github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.20.1-beta
	github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20231206022513-be65ab941093
)
