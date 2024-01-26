#/bin/bash

CONTRACTS_PATH="../contracts/"
cd $CONTRACTS_PATH || exit


hardhat_compile(){
  yarn clean
  yarn
  yarn hardhat compile
}

forge_compile (){
CONTRACTS_PATH="../contracts/"
cd $CONTRACTS_PATH || exit
yarn clean
sed -i '.bak' '10a\
import "@foundry-rs/hardhat-forge";' hardhat.config.ts
yarn

directory=$(pwd)"/contracts"
# shellcheck disable=SC2034
main_dirname=$(basename "$directory")

backup_extension=".bak"

find "$directory" -mindepth 1 -maxdepth 1 -type d | while read dir; do
    dirname=$(basename "$dir")
    # shellcheck disable=SC2116
    subdirs=$(echo "'contracts/$dirname'")
    echo "$subdirs"
    sed -i "$backup_extension" "2s|.*|src = $subdirs|" foundry.toml
    yarn hardhat compile
done

sed -i "$backup_extension" "2s|.*|src = 'contracts'|" foundry.toml
sed -i "$backup_extension" '11s/.*//' hardhat.config.ts
rm -rf foundry.toml.bak
rm -rf hardhat.config.ts.bak
}

#!/bin/bash
if [ ! -n "$1" ] ;then
    echo "hardhat compile"
    hardhat_compile
else
    echo "forge compile"
    forge_compile
fi

#go run ./gen/main.go -artifacts ../contracts/artifacts -out ./bindings -contracts Staking -package bindings