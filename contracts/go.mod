module github.com/morph-l2/contract

go 1.19

require (
	github.com/iden3/go-iden3-crypto v0.0.15
	github.com/scroll-tech/go-ethereum v1.11.4
)

require (
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace github.com/scroll-tech/go-ethereum => github.com/morph-l2/go-ethereum v1.10.14-0.20240105030148-da6185c8d1cb
