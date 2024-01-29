
build:
	cd ops && docker compose build

base_image:
	cd ops && docker build -t morph/go-rust-builder:go-1.19-rust-nightly-2022-12-10 . -f go-rust-builder.Dockerfile

################## devnet 4 nodes ####################

devnet-up:
	python3 ./devnet-morph/main.py --polyrepo-dir=.
.PHONY: devnet-up

devnet-up-mockccc:
	python3 ./devnet-morph/main.py --polyrepo-dir=. --mockccc
.PHONY: devnet-up-deploy

devnet-up-debugccc:
	python3 ./devnet-morph/main.py --polyrepo-dir=. --debugccc
.PHONY: devnet-up-debugccc

devnet-stake:
	cast send  0x6900000000000000000000000000000000000010 "stake()"  --rpc-url http://127.0.0.1:9545  --private-key 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --value 2ether

devnet-smoketest:
	node contracts/scripts/eth_bridge.js

devnet-down:
	cd ./ops && docker compose -f docker-compose-4nodes.yml down
.PHONY: devnet-down

devnet-clean:
	cd ops && docker compose -f docker-compose-4nodes.yml down
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf L2-genesis/.devnet
	rm -rf ops/.morphdevnet
.PHONY: devnet-clean

devnet-reset:
	make devnet-clean
	make devnet-up-deploy
.PHONY: devnet-reset

devnet-l1:
	python3 ./devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-logs:
	@(cd ./ops && docker-compose logs -f)
	.PHONY: devnet-logs
