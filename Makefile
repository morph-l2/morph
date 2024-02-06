.PHONY: update format

init:
	git submodule update --init

update: ## update the dependencies
	git submodule update --remote --merge
	go work sync

format: ## format the code
	go work sync
	goimports -local $(PWD)/bindings/ -w .
	goimports -local $(PWD)/contracts/ -w .
	goimports -local $(PWD)/node/ -w .
	goimports -local $(PWD)/tx-submitter/ -w .


build:
	cd ops/docker && docker compose build

base_image:
	cd opsdocker && docker build -t morph/go-rust-builder:go-1.19-rust-nightly-2022-12-10 . -f go-rust-builder.Dockerfile

################## devnet 4 nodes ####################

devnet-up:
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

devnet-clean:
	cd ops/docker && docker compose -f docker-compose-4nodes.yml down
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=docker-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf ops/L2-genesis/.devnet
	rm -rf ops/docker/consensus/beacondata ops/docker/consensus/validatordata ops/docker/consensus/genesis.ssz
	rm -rf ops/docker/execution/geth
.PHONY: devnet-clean

devnet-l1:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-logs:
	@(cd ops/docker && docker-compose logs -f)
	.PHONY: devnet-logs

submodules:
	git submodule update --init --recursive
.PHONY: submodules