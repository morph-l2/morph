rm -rf ./deployFile.json &&
    yarn hardhat deploy --storagepath ./deployFile.json --network local &&
    yarn hardhat initialize --storagepath ./deployFile.json --network local &&
    yarn hardhat fund --network local &&
    yarn hardhat register --storagepath ./deployFile.json --network local
