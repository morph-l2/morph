#!/usr/bin/env bash

set -euo pipefail

CONTRACTS_PATH="../contracts/"
cd $CONTRACTS_PATH || exit

hardhat_compile() {
    yarn clean
    yarn
    yarn hardhat compile
}

forge_compile() {
    CONTRACTS_PATH="../contracts/"
    cd $CONTRACTS_PATH || exit
    yarn clean

    config_backup=$(mktemp)
    foundry_backup=$(mktemp)
    cp hardhat.config.ts "$config_backup"
    cp foundry.toml "$foundry_backup"
    restore_sources() {
        cp "$config_backup" hardhat.config.ts
        cp "$foundry_backup" foundry.toml
        rm -f "$config_backup" "$foundry_backup"
    }
    trap restore_sources EXIT

    {
        echo 'import "@foundry-rs/hardhat-forge";'
        cat "$config_backup"
    } > hardhat.config.ts
    yarn

    directory=$(pwd)"/contracts"
    # shellcheck disable=SC2034
    main_dirname=$(basename "$directory")

    find "$directory" -mindepth 1 -maxdepth 1 -type d | while read dir; do
        dirname=$(basename "$dir")
        # shellcheck disable=SC2116
        subdirs=$(echo "'contracts/$dirname'")
        echo "$subdirs"
        awk -v source="$subdirs" 'NR == 2 { $0 = "src = " source } { print }' \
            "$foundry_backup" > foundry.toml
        yarn hardhat compile
    done

    restore_sources
    trap - EXIT
}

#!/bin/bash
if [ "$#" -eq 0 ]; then
    echo "hardhat compile"
    hardhat_compile
else
    echo "forge compile"
    forge_compile
fi

#go run ./gen/main.go -artifacts ../contracts/artifacts -out ./bindings -contracts Staking -package bindings
