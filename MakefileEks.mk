# New release system: MakefileEks.mk only holds service start / stop commands.
# Build artifacts are produced by MakefileBuild.mk.
#
# Target naming convention:
#   start-${env_key}-${app_name}
#   stop-${env_key}-${app_name}   (define when graceful shutdown is needed)

# gas-oracle
# mainnet
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	/data/secret-manager-wrapper ./app

# qanet
start-bk-test-morph-test-qanet-to-morph-gas-price-oracle-qanet:
	/data/secret-manager-wrapper ./app

# hoodi
start-bk-prod-morph-prod-testnet-to-morph-gas-price-oracle-hoodi:
	/data/secret-manager-wrapper ./app

# prover
# mainnet
start-bk-prod-morph-prod-mainnet-to-morph-prover:
	/data/secret-manager-wrapper ./prover-server

# testnet
start-bk-prod-morph-prod-testnet-to-morph-prover-hoodi:
	/data/secret-manager-wrapper ./prover-server

# qanet
start-bk-test-morph-test-qanet-to-morph-prover:
	/data/secret-manager-wrapper ./prover-server

# challenge-handler
# mainnet
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler

# testnet
start-bk-prod-morph-prod-testnet-to-morph-challenge-handler-hoodi:
	/data/secret-manager-wrapper ./challenge-handler

# qanet
start-bk-test-morph-test-qanet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler

# shadow-proving
# mainnet
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	/data/secret-manager-wrapper  ./shadow-proving

# testnet
start-bk-prod-morph-prod-testnet-to-morph-shadow-proving-hoodi:
	/data/secret-manager-wrapper ./shadow-proving

# qanet
start-bk-test-morph-test-qanet-to-morph-shadow-proving:
	/data/secret-manager-wrapper ./shadow-proving

# token-price-oracle
# qanet
start-bk-test-morph-test-qanet-to-morph-token-price-oracle:
	/data/secret-manager-wrapper  ./token-price-oracle

# hoodi
start-bk-prod-morph-prod-testnet-to-morph-token-price-oracle-hoodi:
	/data/secret-manager-wrapper ./token-price-oracle

# mainnet
start-bk-prod-morph-prod-mainnet-to-morph-token-price-oracle:
	/data/secret-manager-wrapper ./token-price-oracle
