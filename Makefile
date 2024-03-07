.PHONY: go-work-sync format

go-work-sync: ## update the dependencies
	go work sync
.PHONY: go-work-sync

submodules:
	git submodule update --init --recursive
.PHONY: submodules

format: ## format the code
	go work sync
	goimports -local $(PWD)/bindings/ -w .
	goimports -local $(PWD)/contracts/ -w .
	goimports -local $(PWD)/node/ -w .
	goimports -local $(PWD)/tx-submitter/ -w .
.PHONY: format

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

devnet-clean: devnet-down
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=docker-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf ops/L2-genesis/.devnet
	rm -rf ops/docker/consensus/beacondata ops/docker/consensus/validatordata ops/docker/consensus/genesis.ssz
	rm -rf ops/docker/execution/geth
.PHONY: devnet-clean

devnet-clean-build:
	cd ops/docker && docker compose -f docker-compose-4nodes.yml down
	docker volume ls --filter name=docker-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf ops/L2-genesis/.devnet
	rm -rf ops/docker/consensus/beacondata ops/docker/consensus/validatordata ops/docker/consensus/genesis.ssz
	rm -rf ops/docker/execution/geth
.PHONY: devnet-clean-build

devnet-l1:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-logs:
	@(cd ops/docker && docker-compose logs -f)
	.PHONY: devnet-logs