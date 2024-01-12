SHELL := /bin/bash

pkg := bindings

all: version mkdir bindings more

bindings: l1block-bindings \
	l1-blocknumber-bindings \
	system-config-bindings \
	l1-cross-domain-messenger-bindings \
	l1-standard-bridge-bindings \
	l2-to-l1-message-passer-bindings \
	morph-portal-bindings \
	gas-price-oracle-bindings \
	legacy-message-passer-bindings \
	address-manager-bindings \
	l2-cross-domain-messenger-bindings \
	l2-standard-bridge-bindings \
	sequencer-fee-vault-bindings \
	morph-mintable-erc20-factory-bindings \
	morph-mintable-erc20-bindings \
	proxy-bindings \
	proxy-admin-bindings \
	erc20-bindings \
	weth9-bindings \
	deployer-whitelist-bindings \
	l2-erc721-bridge-bindings \
	l1-erc721-bridge-bindings \
	morph-mintable-erc721-factory-bindings \
	l1-fee-vault-bindings \
	basefee-vault-bindings \
	morph-rollup-bindings \
	legacy-erc20-eth-bindings \
	l1-staking-bindings \
	l1-sequencer-bindings \
	l2-sequencer-bindings \
	l2-gov-bindings \
	submitter-bindings

version:
	forge --version
	abigen --version

compile:
	./compile.sh

system-config-bindings: compile
	./gen_bindings.sh contracts/L1/SystemConfig.sol:SystemConfig $(pkg)

l1-staking-bindings: compile
	./gen_bindings.sh contracts/L1/staking/Staking.sol:Staking $(pkg)

l1-sequencer-bindings: compile
	./gen_bindings.sh contracts/L1/staking/L1Sequencer.sol:L1Sequencer $(pkg)

l2-sequencer-bindings: compile
	./gen_bindings.sh contracts/L2/L2Sequencer.sol:L2Sequencer $(pkg)

l2-gov-bindings: compile
	./gen_bindings.sh contracts/L2/Gov.sol:Gov $(pkg)

submitter-bindings: compile
	./gen_bindings.sh contracts/L2/Submitter.sol:Submitter $(pkg)

l1-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger $(pkg)

l1-standard-bridge-bindings: compile
	./gen_bindings.sh contracts/L1/L1StandardBridge.sol:L1StandardBridge $(pkg)

morph-portal-bindings: compile
	./gen_bindings.sh contracts/L1/MorphPortal.sol:MorphPortal $(pkg)

morph-rollup-bindings: compile
	./gen_bindings.sh contracts/L1/Rollup.sol:Rollup $(pkg)

address-manager-bindings: compile
	./gen_bindings.sh contracts/legacy/AddressManager.sol:AddressManager $(pkg)

l1block-bindings: compile
	./gen_bindings.sh contracts/L2/L1Block.sol:L1Block $(pkg)

l2-to-l1-message-passer-bindings: compile
	./gen_bindings.sh contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser $(pkg)

gas-price-oracle-bindings: compile
	./gen_bindings.sh contracts/L2/GasPriceOracle.sol:GasPriceOracle $(pkg)

l2-cross-domain-messenger-bindings: compile
	./gen_bindings.sh contracts/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger $(pkg)

l2-standard-bridge-bindings: compile
	./gen_bindings.sh contracts/L2/L2StandardBridge.sol:L2StandardBridge $(pkg)

l2-erc721-bridge-bindings:
	./gen_bindings.sh contracts/L2/L2ERC721Bridge.sol:L2ERC721Bridge $(pkg)

l1-erc721-bridge-bindings:
	./gen_bindings.sh contracts/L1/L1ERC721Bridge.sol:L1ERC721Bridge $(pkg)

morph-mintable-erc721-factory-bindings:
	./gen_bindings.sh contracts/universal/MorphMintableERC721Factory.sol:MorphMintableERC721Factory $(pkg)

sequencer-fee-vault-bindings: compile
	./gen_bindings.sh contracts/L2/SequencerFeeVault.sol:SequencerFeeVault $(pkg)

basefee-vault-bindings: compile
	./gen_bindings.sh contracts/L2/BaseFeeVault.sol:BaseFeeVault $(pkg)

l1-fee-vault-bindings: compile
	./gen_bindings.sh contracts/L2/L1FeeVault.sol:L1FeeVault $(pkg)

morph-mintable-erc20-factory-bindings: compile
	./gen_bindings.sh contracts/universal/MorphMintableERC20Factory.sol:MorphMintableERC20Factory $(pkg)

morph-mintable-erc20-bindings: compile
	./gen_bindings.sh contracts/universal/MorphMintableERC20.sol:MorphMintableERC20 $(pkg)

legacy-erc20-eth-bindings: compile
	./gen_bindings.sh contracts/legacy/LegacyERC20ETH.sol:LegacyERC20ETH $(pkg)

proxy-bindings: compile
	./gen_bindings.sh contracts/universal/Proxy.sol:Proxy $(pkg)

proxy-admin-bindings: compile
	./gen_bindings.sh contracts/universal/ProxyAdmin.sol:ProxyAdmin $(pkg)

legacy-message-passer-bindings: compile
	./gen_bindings.sh contracts/legacy/LegacyMessagePasser.sol:LegacyMessagePasser $(pkg)

erc20-bindings: compile
	./gen_bindings.sh node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol:ERC20 $(pkg)

weth9-bindings: compile
	./gen_bindings.sh contracts/vendor/WETH9.sol:WETH9 $(pkg)

deployer-whitelist-bindings: compile
	./gen_bindings.sh contracts/legacy/DeployerWhitelist.sol:DeployerWhitelist $(pkg)

l1-blocknumber-bindings: compile
	./gen_bindings.sh contracts/legacy/L1BlockNumber.sol:L1BlockNumber $(pkg)

more:
	go run ./gen/main.go \
		-artifacts ../contracts/artifacts \
		-out ./bindings \
		-contracts GasPriceOracle,SystemConfig,MorphMintableERC20Factory,L2StandardBridge,L1BlockNumber,LegacyMessagePasser,DeployerWhitelist,Proxy,MorphPortal,L2ToL1MessagePasser,L2CrossDomainMessenger,SequencerFeeVault,L1Block,LegacyERC20ETH,WETH9,GovernanceToken,L1CrossDomainMessenger,L2ERC721Bridge,MorphMintableERC721Factory,ProxyAdmin,Rollup,Staking,L1Sequencer,L2Sequencer,Gov,Submitter \
		-package bindings

mkdir:
	mkdir -p bin $(pkg)

clean:
	rm -rf bin $(pkg)

test:
	go test ./...
