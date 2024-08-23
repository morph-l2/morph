SHELL := /bin/bash

pkg := bindings

all: version mkdir bindings more

bindings: \
	l1-staking-bindings \
	l1-cross-domain-messenger-bindings \
	l1-standard-erc20-gateway-bindings \
	l1-eth-gateway-bindings \
	l1-erc20-gateway-bindings \
	l1-gateway-router-bindings \
	l1-weth-gateway-bindings \
	l1-usdc-gateway-bindings \
	l1-message-queue-with-gas-price-oracle-bindings \
	l1-erc721-gateway-bindings \
	l1-erc1155-gateway-bindings \
	l1-custom-erc20-gateway-bindings \
	l1-reverse-custom-gateway-bindings \
	enforced-tx-gateway-bindings \
    morph-rollup-bindings \
    multiple-version-rollup-verifier-bindings \
	l2-cross-domain-messenger-bindings \
    l2-gas-price-oracle-bindings \
    l2-to-l1-message-passer-bindings \
    l2-tx-fee-vault-bindings \
	l2-gov-bindings \
	l2-sequencer-bindings \
	l2-staking-bindings \
	l2-distribute-bindings \
	l2-record-bindings \
	l2-eth-gateway-bindings \
	l2-standard-erc20-gateway-bindings \
	l2-gateway-router-bindings \
	l2-weth-gateway-bindings \
	l2-usdc-gateway-bindings \
	l2-erc20-gateway-bindings \
	l2-erc721-gateway-bindings \
	l2-erc1155-gateway-bindings \
	l2-custom-erc20-gateway-bindings \
	l2-withdrawlock-erc20-gateway-bindings \
	l2-reverse-custom-gateway-bindings \
	white-list-bindings \
	transparent-upgradeable-proxy-bindings \
	proxy-admin-bindings \
	morph-token-bindings \
	morph-standard-erc20-bindings \
	morph-standard-erc20-factory-bindings \
	morph-wrapped-ether-bindings \
	erc20-bindings
version:
	forge --version
	abigen --version

compile:
	./compile.sh

compile-forge:
	./compile.sh forge

## L1
# Staking
l1-staking-bindings: compile
	./gen_bindings.sh contracts/l1/staking/L1Staking.sol:L1Staking $(pkg)

# CrossDomainMessenger
l1-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/l1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger $(pkg)

# Gateways
l1-standard-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1StandardERC20Gateway.sol:L1StandardERC20Gateway $(pkg)
l1-eth-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1ETHGateway.sol:L1ETHGateway $(pkg)
l1-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway $(pkg)
l1-gateway-router-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1GatewayRouter.sol:L1GatewayRouter $(pkg)
l1-weth-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1WETHGateway.sol:L1WETHGateway $(pkg)
l1-usdc-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/usdc/L1USDCGateway.sol:L1USDCGateway $(pkg)
l1-erc721-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1ERC721Gateway.sol:L1ERC721Gateway $(pkg)
l1-erc1155-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1ERC1155Gateway.sol:L1ERC1155Gateway $(pkg)
l1-custom-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1CustomERC20Gateway.sol:L1CustomERC20Gateway $(pkg)
l1-reverse-custom-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/L1ReverseCustomGateway.sol:L1ReverseCustomGateway $(pkg)
enforced-tx-gateway-bindings: compile
	./gen_bindings.sh contracts/l1/gateways/EnforcedTxGateway.sol:EnforcedTxGateway $(pkg)

# Rollup
l1-message-queue-with-gas-price-oracle-bindings: compile
	./gen_bindings.sh contracts/l1/rollup/L1MessageQueueWithGasPriceOracle.sol:L1MessageQueueWithGasPriceOracle $(pkg)
morph-rollup-bindings: compile
	./gen_bindings.sh contracts/l1/rollup/Rollup.sol:Rollup $(pkg)
multiple-version-rollup-verifier-bindings: compile
	./gen_bindings.sh contracts/l1/rollup/MultipleVersionRollupVerifier.sol:MultipleVersionRollupVerifier $(pkg)

## L2
# CrossDomainMessenger
l2-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/l2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger $(pkg)

# System
l2-gas-price-oracle-bindings: compile
	./gen_bindings.sh contracts/l2/system/GasPriceOracle.sol:GasPriceOracle $(pkg)
l2-to-l1-message-passer-bindings: compile
	./gen_bindings.sh contracts/l2/system/L2ToL1MessagePasser.sol:L2ToL1MessagePasser $(pkg)
l2-tx-fee-vault-bindings: compile
	./gen_bindings.sh contracts/l2/system/L2TxFeeVault.sol:L2TxFeeVault $(pkg)

# Staking
l2-gov-bindings: compile
	./gen_bindings.sh contracts/l2/staking/Gov.sol:Gov $(pkg)
l2-sequencer-bindings: compile
	./gen_bindings.sh contracts/l2/staking/Sequencer.sol:Sequencer $(pkg)
l2-staking-bindings: compile
	./gen_bindings.sh contracts/l2/staking/L2Staking.sol:L2Staking $(pkg)
l2-distribute-bindings: compile
	./gen_bindings.sh contracts/l2/staking/Distribute.sol:Distribute $(pkg)
l2-record-bindings: compile
	./gen_bindings.sh contracts/l2/staking/Record.sol:Record $(pkg)

# Gateways
l2-eth-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2ETHGateway.sol:L2ETHGateway $(pkg)
l2-standard-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway $(pkg)
l2-gateway-router-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2GatewayRouter.sol:L2GatewayRouter $(pkg)
l2-weth-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2WETHGateway.sol:L2WETHGateway $(pkg)
l2-usdc-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/usdc/L2USDCGateway.sol:L2USDCGateway $(pkg)
l2-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway $(pkg)
l2-erc721-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2ERC721Gateway.sol:L2ERC721Gateway $(pkg)
l2-erc1155-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway $(pkg)
l2-custom-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2CustomERC20Gateway.sol:L2CustomERC20Gateway $(pkg)
l2-withdrawlock-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2WithdrawLockERC20Gateway.sol:L2WithdrawLockERC20Gateway $(pkg)
l2-reverse-custom-gateway-bindings: compile
	./gen_bindings.sh contracts/l2/gateways/L2ReverseCustomGateway.sol:L2ReverseCustomGateway $(pkg)

# Common
white-list-bindings: compile
	./gen_bindings.sh contracts/libraries/common/Whitelist.sol:Whitelist $(pkg)

# Proxy
transparent-upgradeable-proxy-bindings: compile
	./gen_bindings.sh node_modules/@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy $(pkg)
proxy-admin-bindings: compile
	./gen_bindings.sh node_modules/@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol:ProxyAdmin $(pkg)

# Token
morph-token-bindings: compile
	./gen_bindings.sh contracts/l2/system/MorphToken.sol:MorphToken $(pkg)
morph-standard-erc20-bindings: compile
	./gen_bindings.sh contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20 $(pkg)
morph-standard-erc20-factory-bindings: compile
	./gen_bindings.sh contracts/libraries/token/MorphStandardERC20Factory.sol:MorphStandardERC20Factory $(pkg)
morph-wrapped-ether-bindings: compile
	./gen_bindings.sh contracts/l2/system/WrappedEther.sol:WrappedEther $(pkg)
erc20-bindings: compile
	./gen_bindings.sh node_modules/@rari-capital/solmate/src/tokens/ERC20.sol:ERC20 $(pkg)
#fiat-v1-bindings: compile
#	./gen_bindings.sh contracts/libraries/token/FiatTokenV1.sol:FiatTokenV1 $(pkg)

more: compile-forge
	go run ./gen/main.go \
		-artifacts ../contracts/artifacts \
		-out ./bindings \
		-contracts ProxyAdmin,TransparentUpgradeableProxy,L1MessageQueueWithGasPriceOracle,L1USDCGateway,L1Staking,L1CrossDomainMessenger,L1StandardERC20Gateway,L1ETHGateway,L1ERC20Gateway,L1GatewayRouter,L1WETHGateway,Rollup,MultipleVersionRollupVerifier,L2CrossDomainMessenger,GasPriceOracle,L2ToL1MessagePasser,L2TxFeeVault,Sequencer,Gov,Distribute,L2Staking,Record,L2ETHGateway,L2StandardERC20Gateway,L2GatewayRouter,L2WETHGateway,L2ERC20Gateway,L2ERC721Gateway,L2ERC1155Gateway,MorphToken,MorphStandardERC20,MorphStandardERC20Factory,WrappedEther,L2USDCGateway,EnforcedTxGateway,L1ERC721Gateway,L1ERC1155Gateway,L1ReverseCustomGateway,L2WithdrawLockERC20Gateway,L2ReverseCustomGateway,Whitelist \
		-package bindings

mkdir:
	mkdir -p bin $(pkg)

clean:
	rm -rf bin $(pkg)
.PHONY: clean

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...
.PHONY: test

lint:
	GOBIN=$(PWD)/build/bin go run ../build/lint.go
.PHONY: lint

