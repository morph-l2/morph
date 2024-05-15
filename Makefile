################## update dependencies ####################

ETHEREUM_TAG=v1.10.14-0.20240429050506-03fd4c3e771d
TENDERMINT_TAG=v0.2.0-beta.0.20240513090937-03bf2a578b48
BTCD_TAG=v0.20.1-beta

update:
	go work sync

	cd $(PWD)/bindings/ && \
		sed -i '' '6s/.*/	github.com\/btcsuite\/btcd => github.com\/btcsuite\/btcd ${BTCD_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '8s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
	cd $(PWD)/contracts/ && \
		sed -i '' '6s/.*/	github.com\/btcsuite\/btcd => github.com\/btcsuite\/btcd ${BTCD_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '8s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
	cd $(PWD)/node/ && \
		sed -i '' '6s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
	cd $(PWD)/ops/l2-genesis/ && \
		sed -i '' '6s/.*/	github.com\/btcsuite\/btcd => github.com\/btcsuite\/btcd ${BTCD_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '8s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
	cd $(PWD)/ops/tools/ && \
		sed -i '' '6s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
	cd $(PWD)/oracle/ && \
		sed -i '' '6s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
	cd $(PWD)/tx-submitter/ && \
		sed -i '' '6s/.*/	github.com\/scroll-tech\/go-ethereum => github.com\/morph-l2\/go-ethereum ${ETHEREUM_TAG}/' go.mod && \
		sed -i '' '7s/.*/	github.com\/tendermint\/tendermint => github.com\/morph-l2\/tendermint ${TENDERMINT_TAG}/' go.mod && \
		go mod tidy
.PHONY: update

submodules:
	git submodule update --init
	git submodule update --remote 
.PHONY: submodules

################## bindings ####################

bindings:
	make -C bindings all
.PHONY: bindings

################## lint code ####################

lint: lint-sol lint-go
.PHONY: lint

# npm install --global --save-dev solhint
lint-sol:
	solhint $$(find contracts -name '*.sol' -not -path '**/@openzeppelin/**')
.PHONY: lint-sol

lint-go:
	make -C bindings lint
	make -C contracts lint
	make -C node lint
	make -C ops/l2-genesis lint
	make -C ops/tools lint
	make -C oracle lint
	make -C tx-submitter lint
.PHONY: lint-go

################## format code ####################

fmt: fmt-sol fmt-go
.PHONY: fmt

# npm install --global --save-dev prettier-plugin-solidity
fmt-sol:
	find ./contracts/ -name '*.sol' -type f -not -path "**/node_modules*" | xargs misspell -w
	cd $(PWD)/contracts/ && yarn prettier --write --plugin=prettier-plugin-solidity './contracts/**/*.sol'
.PHONY: fmt-sol

fmt-go:
	go work sync
	cd $(PWD)/bindings/ && go mod tidy
	cd $(PWD)/contracts/ && go mod tidy
	cd $(PWD)/node/ && go mod tidy
	cd $(PWD)/ops/l2-genesis/ && go mod tidy
	cd $(PWD)/ops/tools/ && go mod tidy
	cd $(PWD)/oracle/ && go mod tidy
	cd $(PWD)/tx-submitter/ && go mod tidy
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs goimports -w -local $(PWD)/
.PHONY: fmt-go

################## docker build ####################

docker-build:
	cd ops/docker && docker compose build
.PHONY: docker-build

base-image:
	cd ops/docker && docker build -t morph/go-rust-builder:go-1.19-rust-nightly-2022-12-10 . -f go-rust-builder.Dockerfile
.PHONY: base-image

################## devnet 4 nodes ####################

devnet-up: submodules
	python3 ops/devnet-morph/main.py --polyrepo-dir=.
.PHONY: devnet-up

devnet-up-mockccc:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --mockccc
.PHONY: devnet-up-mockccc

devnet-up-debugccc:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --debugccc
.PHONY: devnet-up-debugccc

devnet-down:
	cd ops/docker && docker compose -f docker-compose-4nodes.yml down
.PHONY: devnet-down

devnet-clean-build: devnet-down
	docker volume ls --filter name=docker-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf ops/l2-genesis/.devnet
	rm -rf ops/docker/consensus/beacondata ops/docker/consensus/validatordata ops/docker/consensus/genesis.ssz
	rm -rf ops/docker/execution/geth
.PHONY: devnet-clean-build

devnet-clean: devnet-clean-build
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
.PHONY: devnet-clean

devnet-l1:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-logs:
	@(cd ops/docker && docker-compose logs -f)
.PHONY: devnet-logs

# tx-submitter
SUBMITTERS := $(shell grep -o 'tx-submitter-[0-9]*[^:]' ops/docker/docker-compose-4nodes.yml | sort | uniq)
rebuild-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker/docker-compose-4nodes.yml up -d --build $$submitter --no-deps; \
	done
stop-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker-compose-4nodes.yml stop $$submitter; \
	done
start-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker-compose-4nodes.yml start $$submitter; \
	done

# build geth
nccc_geth: submodules
	cd go-ethereum && env GO111MODULE=on GOWORK=off go run build/ci.go install ./cmd/geth
