SHELL := /bin/bash

pkg := bindings

all: version mkdir bindings more

bindings: l1-staking-bindings \
	l1-sequencer-bindings \
    submitter-bindings \
	l1-cross-domain-messenger-bindings \
	l1-standard-erc20-gateway-bindings \
	l1-eth-gateway-bindings \
	l1-erc20-gateway-bindings \
	l1-gateway-router-bindings \
	l1-weth-gateway-bindings \
	l1-message-queue-bindings \
    l1-gas-price-oracle-bindings \
    morph-rollup-bindings \
    multiple-version-rollup-verifier-bindings \
	address-manager-bindings \
	l2-cross-domain-messenger-bindings \
    l2-submitter-bindings \
    l2-gas-price-oracle-bindings \
    l2-to-l1-message-passer-bindings \
    l2-tx-fee-vault-bindings \
	l2-gov-bindings \
	l2-sequencer-bindings \
	l2-eth-gateway-bindings \
	l2-standard-erc20-gateway-bindings \
	l2-gateway-router-bindings \
	l2-weth-gateway-bindings \
	l2-erc20-gateway-bindings \
	l2-erc721-gateway-bindings \
	l2-erc1155-gateway-bindings \
	proxy-bindings \
	proxy-admin-bindings \
	governance-token-bindings \
	mint-manager-bindings \
	morph-standard-erc20-bindings \
	morph-standard-erc20-factory-bindings \
	morph-wrapped-ether-bindings

version:
	forge --version
	abigen --version

compile:
	./compile.sh

## L1
# Staking
l1-staking-bindings: compile
	./gen_bindings.sh contracts/L1/staking/Staking.sol:Staking $(pkg)
l1-sequencer-bindings: compile
	./gen_bindings.sh contracts/L1/staking/L1Sequencer.sol:L1Sequencer $(pkg)

# Submitter
submitter-bindings: compile
	./gen_bindings.sh contracts/L2/submitter/Submitter.sol:Submitter $(pkg)

# CrossDomainMessenger
l1-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger $(pkg)

# Gateways
l1-standard-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/L1/gateways/L1StandardERC20Gateway.sol:L1StandardERC20Gateway $(pkg)
l1-eth-gateway-bindings: compile
	./gen_bindings.sh contracts/L1/gateways/L1ETHGateway.sol:L1ETHGateway $(pkg)
l1-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/L1/gateways/L1ERC20Gateway.sol:L1ERC20Gateway $(pkg)
l1-gateway-router-bindings: compile
	./gen_bindings.sh contracts/L1/gateways/L1GatewayRouter.sol:L1GatewayRouter $(pkg)
l1-weth-gateway-bindings: compile
	./gen_bindings.sh contracts/L1/gateways/L1WETHGateway.sol:L1WETHGateway $(pkg)

# Rollup
l1-message-queue-bindings: compile
	./gen_bindings.sh contracts/L1/rollup/L1MessageQueue.sol:L1MessageQueue $(pkg)
l1-gas-price-oracle-bindings: compile
	./gen_bindings.sh contracts/L1/rollup/L2GasPriceOracle.sol:L2GasPriceOracle $(pkg)
morph-rollup-bindings: compile
	./gen_bindings.sh contracts/L1/rollup/Rollup.sol:Rollup $(pkg)
multiple-version-rollup-verifier-bindings: compile
	./gen_bindings.sh contracts/L1/rollup/MultipleVersionRollupVerifier.sol:MultipleVersionRollupVerifier $(pkg)

# Deployment
address-manager-bindings: compile
	./gen_bindings.sh contracts/deployment/AddressManager.sol:AddressManager $(pkg)

## L2
# CrossDomainMessenger
l2-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger $(pkg)

# Submitter
l2-submitter-bindings: compile
	./gen_bindings.sh contracts/L2/submitter/Submitter.sol:Submitter $(pkg)
l2-gas-price-oracle-bindings: compile
	./gen_bindings.sh contracts/L2/system/GasPriceOracle.sol:GasPriceOracle $(pkg)
l2-to-l1-message-passer-bindings: compile
	./gen_bindings.sh contracts/L2/system/L2ToL1MessagePasser.sol:L2ToL1MessagePasser $(pkg)
l2-tx-fee-vault-bindings: compile
	./gen_bindings.sh contracts/L2/system/L2TxFeeVault.sol:L2TxFeeVault $(pkg)

# Staking
l2-gov-bindings: compile
	./gen_bindings.sh contracts/L2/staking/Gov.sol:Gov $(pkg)
l2-sequencer-bindings: compile
	./gen_bindings.sh contracts/L2/staking/L2Sequencer.sol:L2Sequencer $(pkg)

# Gateways
l2-eth-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2ETHGateway.sol:L2ETHGateway $(pkg)
l2-standard-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2StandardERC20Gateway.sol:L2StandardERC20Gateway $(pkg)
l2-gateway-router-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2GatewayRouter.sol:L2GatewayRouter $(pkg)
l2-weth-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2WETHGateway.sol:L2WETHGateway $(pkg)
l2-erc20-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2ERC20Gateway.sol:L2ERC20Gateway $(pkg)
l2-erc721-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2ERC721Gateway.sol:L2ERC721Gateway $(pkg)
l2-erc1155-gateway-bindings: compile
	./gen_bindings.sh contracts/L2/gateways/L2ERC1155Gateway.sol:L2ERC1155Gateway $(pkg)

## Common
# Proxy
proxy-bindings: compile
	./gen_bindings.sh contracts/libraries/proxy/Proxy.sol:Proxy $(pkg)
proxy-admin-bindings: compile
	./gen_bindings.sh contracts/libraries/proxy/ProxyAdmin.sol:ProxyAdmin $(pkg)

# Governance
governance-token-bindings: compile
	./gen_bindings.sh contracts/governance/GovernanceToken.sol:GovernanceToken $(pkg)
mint-manager-bindings: compile
	./gen_bindings.sh contracts/governance/MintManager.sol:MintManager $(pkg)

# Token
morph-standard-erc20-bindings: compile
	./gen_bindings.sh contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20 $(pkg)
morph-standard-erc20-factory-bindings: compile
	./gen_bindings.sh contracts/libraries/token/MorphStandardERC20Factory.sol:MorphStandardERC20Factory $(pkg)
morph-wrapped-ether-bindings: compile
	./gen_bindings.sh contracts/L2/system/WrappedEther.sol:WrappedEther $(pkg)

more:
	go run ./gen/main.go \
		-artifacts ../contracts/artifacts \
		-out ./bindings \
		-contracts Staking,L1Sequencer,L2Sequencer,Gov,Submitter,L1CrossDomainMessenger,L1StandardERC20Gateway,L1ETHGateway,L1ERC20Gateway,L1GatewayRouter,L1WETHGateway,L1MessageQueue,L2GasPriceOracle,Rollup,MultipleVersionRollupVerifier,AddressManager,L2CrossDomainMessenger,Submitter,GasPriceOracle,L2ToL1MessagePasser,L2TxFeeVault,Gov,L2Sequencer,L2ETHGateway,L2StandardERC20Gateway,L2GatewayRouter,L2WETHGateway,L2ERC20Gateway,L2ERC721Gateway,L2ERC1155Gateway,Proxy,ProxyAdmin,GovernanceToken,MintManager,MorphStandardERC20,MorphStandardERC20Factory,WrappedEther \
		-package bindings

mkdir:
	mkdir -p bin $(pkg)

clean:
	rm -rf bin $(pkg)

test:
	go test ./...
