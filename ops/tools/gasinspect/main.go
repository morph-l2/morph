package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/rollup/fees"
)

func main() {
	// 0xb033f4a9c375ee95d7bda9ce462850d5b063bd0c66f841936c6b5b6ee5b80740
	// 0xef0119b11c06eefe6e2db2770d36fdf535e30bee45acc44cc3f67314fc96afe7
	// 0x38fa0d472a3774c180b99c3330292f54b576982c03862c51b2550a0b66a46c6b
	// 0xcf36b911809af47fb00e68e87d1b62957893edb42a8e65b7c92e05edd781b5ff
	// 0x862d1243ee2cbc5cdd8b938e161ca5853cf7304018306b7c30646b448641fdfc
	// 0x354d9cfa06a124a44e5b9723cd84169a2573ebd6011a727ca2e37818ff28ec76
	parseBatch(common.HexToHash("0x2e63e169ca6b26516c8fedfa6666de67ffba0a26772feda99771a212a181e674"))

	// 0x104377bf038501f09f8a1ae1ed17bceae5ea781e371eedf057f0d624d449f342
	// 0x2e40fd91306d2e60c569222bde27754c73ef7f9edcf1046ba23152020ea8ef36
	// 0xc4a8d18102e1cea37bbf93a69b6e7141b23a6420b7ff0017a8492192f8718969
	// 0x1b36e9727d159f13e818fa8543700d709d6e84d297240c9924a9a3106c161214
	// 0x15286e0866642596595e3bbb625ce62ef444b8de7d2f2df840e9ad58df88ce8b
	// 0x5a2de6e67f488ca43c36b38d8ae278162acdc16730c97e3ed00f9e000680d44a
	// 0x4b5e99a8b75ff4d1f3e98cab9183d36e9eee9db28637eebc0567e5a4c65e986e
	//parsePreBatch(common.HexToHash("0x4b5e99a8b75ff4d1f3e98cab9183d36e9eee9db28637eebc0567e5a4c65e986e"))
}

func parsePreBatch(txHash common.Hash) {
	//client, err := ethclient.Dial("http://l2-qa-morph-l1-geth.bitkeep.tools")
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/193cbc4d8995453898757cd575ed37f9")
	if err != nil {
		panic(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	if receipt.Status != 1 {
		panic("tx does not success")
	}
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	rollupABI, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method := rollupABI.Methods["commitBatch"]
	ret, err := method.Inputs.UnpackValues(tx.Data()[4:])
	if err != nil {
		panic(err)
	}
	batchDataInterface := ret[0]
	bd := batchDataInterface.(struct {
		Version                uint8     "json:\"version\""
		ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
		Chunks                 [][]byte  "json:\"chunks\""
		SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
		PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot          [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		Signature              struct {
			Version   *big.Int   "json:\"version\""
			Signers   []*big.Int "json:\"signers\""
			Signature []uint8    "json:\"signature\""
		} "json:\"signature\""
	})
	chunks := bd.Chunks
	fmt.Println("chunks length", len(chunks))
	calldataGasUsed := fees.CalculateL1GasUsed(tx.Data(), big.NewInt(0))
	totalBlockNum, totalL2TxNum, gasUsedForTxs, _ := parseChunks(chunks, true)

	fmt.Printf("Batch details for txHash: %x \n  ===totalBlockNum: %d \n  ===totalL2TxNum: %d \n  ===totalGasUsed: %d \n  ===calldataGasUsed: %s \n ===gasUsedForTxs: %s \n",
		txHash, totalBlockNum, totalL2TxNum, receipt.GasUsed, calldataGasUsed, gasUsedForTxs.String())

	//r := new(big.Float).Quo(big.NewFloat(float64(receipt.GasUsed+156000)), big.NewFloat(1.15))
	//r1, _ := r.Int64()

	r1 := int64(receipt.GasUsed + 156000)

	overhead := (r1 - gasUsedForTxs.Int64()) / int64(totalL2TxNum)
	fmt.Println("overhead: ", overhead)

	collectedGasFloat := new(big.Float).Mul(big.NewFloat(float64(gasUsedForTxs.Int64()+2500*int64(totalL2TxNum))), big.NewFloat(1.15))
	collectedGas, _ := collectedGasFloat.Int64()
	earnGas := collectedGas - int64(receipt.GasUsed+156000)
	fmt.Println("earnGas: ", earnGas)
}

func parseBatch(txHash common.Hash) {
	//client, err := ethclient.Dial("http://l2-qa-morph-l1-geth.bitkeep.tools")
	client, err := ethclient.Dial("http://127.0.0.1:9545")
	if err != nil {
		panic(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	if receipt.Status != 1 {
		panic("tx does not success")
	}
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	rollupABI, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method := rollupABI.Methods["commitBatch"]
	ret, err := method.Inputs.UnpackValues(tx.Data()[4:])
	if err != nil {
		panic(err)
	}
	batchDataInterface := ret[0]
	bd := batchDataInterface.(struct {
		Version                uint8     "json:\"version\""
		ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
		Chunks                 [][]byte  "json:\"chunks\""
		SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
		PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot          [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		Signature              struct {
			Version   *big.Int   "json:\"version\""
			Signers   []*big.Int "json:\"signers\""
			Signature []uint8    "json:\"signature\""
		} "json:\"signature\""
	})
	chunks := bd.Chunks
	fmt.Println("chunks length", len(chunks))
	calldataGasUsed := fees.CalculateL1GasUsed(tx.Data(), big.NewInt(0))
	totalBlockNum, totalL2TxNum, _, l2TxHashes := parseChunks(chunks, false)

	l2client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	var totalTxBytes []byte
	for _, txHash := range l2TxHashes {
		tx, _, err := l2client.TransactionByHash(context.Background(), txHash)
		if err != nil {
			panic(err)
		}
		txBytes, err := tx.MarshalBinary()
		if err != nil {
			panic(err)
		}
		fmt.Println("one tx cost: ", fees.CalculateL1GasUsed(txBytes, big.NewInt(2500)))
		fmt.Println("one tx cost2: ", fees.CalculateL1GasUsed(txBytes, big.NewInt(0)))
		fmt.Println()
		totalTxBytes = append(totalTxBytes, txBytes...)
	}
	collectedTxGasUsed := fees.CalculateL1GasUsed(totalTxBytes, big.NewInt(0))
	fmt.Printf("Batch details for txHash: %x \n  ===totalBlockNum: %d \n  ===totalL2TxNum: %d \n  ===totalGasUsed: %d \n  ===calldataGasUsed: %s \n  ===collectedTxGasUsed: %s \n \n",
		txHash, totalBlockNum, totalL2TxNum, receipt.GasUsed, calldataGasUsed, collectedTxGasUsed.String())

	//r := new(big.Float).Quo(big.NewFloat(float64(receipt.GasUsed+131072+156000)), big.NewFloat(1.15))
	//r1, _ := r.Int64()

	r1 := int64(receipt.GasUsed + 131072 + 156000)
	overhead := (r1 - collectedTxGasUsed.Int64()) / int64(totalL2TxNum)
	fmt.Println("overhead: ", overhead)
}

func parseChunks(chunks [][]byte, non4844 bool) (totalBlockNum uint, totalL2TxNum uint16, gasUsedForTxs *big.Int, l2TxHashes []common.Hash) {
	var txBytes []byte
	for i, chunk := range chunks {
		// oriChunkBytes := make([]byte, 0)
		// copy(oriChunkBytes, chunk)
		fmt.Printf("chunk %d: %s \n", i, hexutil.Encode(chunk))
		reader := bytes.NewReader(chunk)
		var blockNum uint8
		if err := binary.Read(reader, binary.BigEndian, &blockNum); err != nil {
			panic(err)
		}
		totalBlockNum += uint(blockNum)
		for i := 0; i < int(blockNum); i++ {
			bc := make([]byte, 60)
			if err := binary.Read(reader, binary.BigEndian, &bc); err != nil {
				panic(err)
			}
			totalTxNum := binary.BigEndian.Uint16(bc[56:58])
			l1TxNum := binary.BigEndian.Uint16(bc[58:])
			l2TxNum := totalTxNum - l1TxNum
			totalL2TxNum += l2TxNum
		}
		txBytes = append(txBytes, chunk[1+60*uint(blockNum):]...)
	}
	if non4844 {
		fmt.Println("txBytes length: ", len(txBytes))
		if len(txBytes) > 0 {
			gasUsedForTxs = fees.CalculateL1GasUsed(txBytes, big.NewInt(0))
		} else {
			gasUsedForTxs = big.NewInt(0)
		}
	} else {
		count := len(txBytes) / 32
		fmt.Println("txHashes count: ", len(txBytes)/32)
		reader := bytes.NewReader(txBytes)
		for i := 0; i < count; i++ {
			hashBz := make([]byte, 32)
			if err := binary.Read(reader, binary.BigEndian, &hashBz); err != nil {
				panic(err)
			}
			l2TxHashes = append(l2TxHashes, common.BytesToHash(hashBz))
		}
	}

	return
}
