#!/bin/bash  
  
PACKAGE_NAME="@morph-l2/contracts"  
  
FILE_PATH="artifacts/contracts/L1/rollup/Rollup.sol/Rollup.json"  
  
TEMP_DIR=$(mktemp -d) 
  
npm pack $PACKAGE_NAME --pack-destination $TEMP_DIR  
  
tar -zxvf $TEMP_DIR/*.tgz -C $TEMP_DIR  
  
cp $TEMP_DIR/package/${FILE_PATH} .  
  
rm -rf $TEMP_DIR