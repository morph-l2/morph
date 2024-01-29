#/bin/bash
set -eu

CONTRACTS_PATH="../contracts/"


if [ "$#" -ne 2 ]; then
	echo "This script takes 2 arguments - CONTRACT_NAME PACKAGE"
	exit 1
fi

need_cmd() {
  if ! command -v "$1" > /dev/null 2>&1; then
    echo "need '$1' (command not found)"
    exit 1
  fi
}

need_cmd forge
need_cmd abigen

NAME=$1
# This can handle both fully qualified syntax or just
# the name of the contract.
# Fully qualified: path-to-contract-file:contract-name
TYPE=$(echo "$NAME" | cut -d ':' -f2)
CATPATH=$(echo "$NAME" | cut -d ':' -f1)
CATPATH=$(echo "artifacts/$CATPATH/$TYPE.json" )

PACKAGE=$2

# Convert to lower case to respect golang package naming conventions
TYPE_LOWER=$(echo ${TYPE} | tr '[:upper:]' '[:lower:]')
FILENAME="${TYPE_LOWER}_deployed.go"


mkdir -p bin
TEMP=$(mktemp -d)

CWD=$(pwd)
# Build contracts
cd ${CONTRACTS_PATH}
cat ${CATPATH} | jq ".abi" > ${TEMP}/${TYPE}.abi
cat ${CATPATH} | jq ".bytecode" | sed 's/\"//g' > ${TEMP}/${TYPE}.bin
cat ${CATPATH} | jq ".deployedBytecode" | sed 's/\"//g' > ${CWD}/bin/${TYPE_LOWER}_deployed.hex

# Run ABIGEN
cd ${CWD}
abigen \
	--abi ${TEMP}/${TYPE}.abi \
	--bin ${TEMP}/${TYPE}.bin \
	--pkg ${PACKAGE} \
	--type ${TYPE} \
	--out ./${PACKAGE}/${TYPE_LOWER}.go

# Place import "github.com/ethereum/go-ethereum" to "github.com/scroll-tech/go-ethereum"
TO_REPLACED_STRING="github.com/ethereum/go-ethereum"
NEW_STRING="github.com/scroll-tech/go-ethereum"
sed -i "" -e "s#${TO_REPLACED_STRING}#${NEW_STRING}#g" ./${PACKAGE}/${TYPE_LOWER}.go