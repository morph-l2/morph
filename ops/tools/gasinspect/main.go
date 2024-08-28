package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/params"
	"github.com/morph-l2/go-ethereum/rlp"

	"morph-l2/bindings/bindings"
)

var (
	l1Url  = "https://eth-sepolia.g.alchemy.com/v2/3VpooM4yCL1TYnkaedcHraMo234vhN5y"
	txHash = common.HexToHash("0x2fe371018e5c307eb1fe4b784158adcdd7741e27c18c43e723c8948b2583fde7")
)

func main() {
	client, err := ethclient.Dial(l1Url)
	if err != nil {
		panic(err)
	}

	rollupTxCostInfo(client, txHash)

	//rollupTxCostInfo(client, common.HexToHash("0x25101c10431c31dc21304b8984b2d936f78ec53e0152bb6a1b6294aa7af941a4"))
	//
	//rollupTxCostInfo(client, common.HexToHash("0xf81c57f1aaf21026184c0869a61774d49e0d421f0dcdab0c1569d0709dabbd81"))
	//
	//rollupTxCostInfo(client, common.HexToHash("0x9e001e97ffd400ff1761065e84b358125f021daaa22082fd880ec0e46225a054"))
	//
	//rollupTxCostInfo(client, common.HexToHash("0x7be03a5c2a121cb133c10e34029762c80913fd09e0ea856aa4726b6926a52065"))
	//
	//rollupTxCostInfo(client, common.HexToHash("0xdd82c92dc80f55c76471a805f26f3f740dfccbe0179db3f539c4da0f3a7ea677"))
	//
	//rollupTxCostInfo(client, common.HexToHash("0xf5ab1ddc7d5aed9e016ad943475d0baffdbe6c2d94be1b951b419dbeca9edea8"))

}

func rollupTxCostInfo(client *ethclient.Client, txHash common.Hash) {
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	abi, err := bindings.RollupMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	args, err := abi.Methods["submitBatches"].Inputs.Unpack(tx.Data()[4:])
	if err != nil {
		panic(err)
	}

	rollupBatchDatas := args[0].([]struct {
		BlockNumber    uint64    "json:\"blockNumber\""
		Transactions   []uint8   "json:\"transactions\""
		BlockWitness   []uint8   "json:\"blockWitness\""
		PreStateRoot   [32]uint8 "json:\"preStateRoot\""
		PostStateRoot  [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot [32]uint8 "json:\"withdrawalRoot\""
		Signature      struct {
			Signers   [][]uint8 "json:\"signers\""
			Signature []uint8   "json:\"signature\""
		} "json:\"signature\""
	})

	var (
		blockNum         int
		txNum            int
		flexibleCost     uint64
		witnessCost      uint64
		transactionsCost uint64
	)
	for _, batch := range rollupBatchDatas {
		thisWitnessC := dataGasCost(batch.BlockWitness)
		thisTxC := dataGasCost(batch.Transactions)
		witnessCost += thisWitnessC
		transactionsCost += thisTxC
		flexibleCost += thisWitnessC + thisTxC

		blockContexts, err := DecodeBlockContext(batch.BlockWitness)
		if err != nil {
			panic(err)
		}
		blockNum += len(blockContexts)

		txs, err := DecodeTransactions(batch.Transactions)
		if err != nil {
			panic(err)
		}
		txNum += len(txs)
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		panic(err)
	}
	dataGas := dataGasCost(tx.Data())
	blockHeader, err := client.HeaderByHash(context.Background(), receipt.BlockHash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("====>(%s)tx hash: %s \n	batchNum: %d \n	blockNum: %d \n	txNum: %d \n	gas used total: %d \n	calldata cost: %d \n	  -witnessCost: %d\n	  -transactionsCost: %d \n	  -otherFixedCost: %d \n	execution cost: %d \n	G: fixedCallData+execution cost: %d \n",
		time.Unix(int64(blockHeader.Time), 0).String(),
		txHash.Hex(),
		len(rollupBatchDatas),
		blockNum,
		txNum,
		receipt.GasUsed,
		dataGas,
		witnessCost,
		transactionsCost,
		dataGas-flexibleCost,
		receipt.GasUsed-dataGas,
		receipt.GasUsed-flexibleCost,
	)
	fmt.Println()
}

func dataGasCost(data []byte) uint64 {
	zeroes, ones := zeroesAndOnes(data)
	zeroesGas := zeroes * params.TxDataZeroGas
	onesGas := (ones) * params.TxDataNonZeroGasEIP2028
	return zeroesGas + onesGas
}

// zeroesAndOnes counts the number of 0 bytes and non 0 bytes in a byte slice
func zeroesAndOnes(data []byte) (uint64, uint64) {
	var zeroes uint64
	var ones uint64
	for _, byt := range data {
		if byt == 0 {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}

// number || timestamp || base_fee || gas_limit || num_txs || tx_hashs
type BlockInfo struct {
	Number    *big.Int
	Timestamp uint64
	BaseFee   *big.Int
	GasLimit  uint64
	NumTxs    uint64
}

func DecodeBlockContext(bs []byte) ([]*BlockInfo, error) {
	blockContexts := make([]*BlockInfo, 0)
	// [block1, block2, ..., blockN]
	reader := bytes.NewReader(bs)
	for {
		if reader.Len() == 0 {
			break
		}
		block := new(BlockInfo)
		// number || timestamp || base_fee || gas_limit || num_txs
		bsBlockNumber := make([]byte, 32)
		if _, err := reader.Read(bsBlockNumber[:]); err != nil {
			return nil, err
		}
		block.Number = new(big.Int).SetBytes(bsBlockNumber)

		if err := binary.Read(reader, binary.BigEndian, &block.Timestamp); err != nil {
			return nil, err
		}
		// [32]byte uint256
		bsBaseFee := make([]byte, 32)
		if _, err := reader.Read(bsBaseFee[:]); err != nil {
			return nil, err
		}
		block.BaseFee = new(big.Int).SetBytes(bsBaseFee)
		if err := binary.Read(reader, binary.BigEndian, &block.GasLimit); err != nil {
			return nil, err
		}
		if err := binary.Read(reader, binary.BigEndian, &block.NumTxs); err != nil {
			return nil, err
		}
		for i := 0; i < int(block.NumTxs); i++ {
			txHash := common.Hash{}
			if _, err := reader.Read(txHash[:]); err != nil {
				return nil, err
			}
			// drop txHash
		}
		blockContexts = append(blockContexts, block)
	}
	return blockContexts, nil
}

func DecodeTransactions(bs []byte) ([]*eth.Transaction, error) {
	ret := make([]*eth.Transaction, 0)
	if err := rlp.DecodeBytes(bs, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}
