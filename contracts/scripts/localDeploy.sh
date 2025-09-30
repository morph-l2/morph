rm -rf ./deployFile.json &&
    yarn hardhat deploy --storagepath ./deployFile.json --network l1 &&
    yarn hardhat initialize --storagepath ./deployFile.json --network l1 &&
    yarn hardhat fund --network l1 &&
    yarn hardhat register --storagepath ./deployFile.json --network l1
