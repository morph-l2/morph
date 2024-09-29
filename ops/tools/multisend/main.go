package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
)

var (
	chainId     *big.Int
	fundPrivKey = crypto.ToECDSAUnsafe(hexutil.MustDecode("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"))
	urls        = []string{"http://localhost:8945", "http://localhost:8545", "http://localhost:8645", "http://localhost:8745", "http://localhost:8845"}
	senderNum   = 50
	duration    = 120 * time.Minute
)

func main() {
	var err error
	clients := make([]*ethclient.Client, len(urls))
	for i, url := range urls {
		client, err := ethclient.Dial(url)
		if err != nil {
			panic(err)
		}
		clients[i] = client
	}
	chainId, err = clients[0].ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	senderPks := distribute(clients[0])
	receiver, _ := crypto.GenerateKey()
	to := crypto.PubkeyToAddress(receiver.PublicKey)
	start := time.Now()
	var count int
	for {
		if count == 10000000 || time.Now().Sub(start) > duration {
			fmt.Println("completed")
			break
		}
		for i, sender := range senderPks {
			go func(s *ecdsa.PrivateKey, senderIndex int) {
				transactOpts, err := bind.NewKeyedTransactorWithChainID(s, chainId)
				if err != nil {
					panic(err)
				}
				transactOpts.Value = big.NewInt(10)
				index := senderIndex % len(clients)
				client := clients[index]
				_, err = Transfer(transactOpts, to, client)
				if err != nil {
					fmt.Printf("error found when transfer: %v \n", err)
				}
				fmt.Println("sent tx done")
			}(sender, i)
		}
		count++
		time.Sleep(300 * time.Millisecond)
		fmt.Println()
	}
}

func distribute(client *ethclient.Client) []*ecdsa.PrivateKey {
	senderPks := make([]*ecdsa.PrivateKey, senderNum)
	distributeTxHashes := make([]common.Hash, senderNum)
	for i := 0; i < senderNum; i++ {
		senderPriKey, _ := crypto.GenerateKey()
		to := crypto.PubkeyToAddress(senderPriKey.PublicKey)
		transactOpts, err := bind.NewKeyedTransactorWithChainID(fundPrivKey, chainId)
		if err != nil {
			panic(err)
		}
		value, _ := new(big.Int).SetString("10000000000000000000", 10)
		transactOpts.Value = value
		tx, err := Transfer(transactOpts, to, client)
		if err != nil {
			panic(err)
		}
		senderPks[i] = senderPriKey
		distributeTxHashes[i] = tx.Hash()
	}

	time.Sleep(2 * time.Second)

	for _, txHash := range distributeTxHashes {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err != nil {
			if err == ethereum.NotFound {
				time.Sleep(time.Second)
				receipt, err = client.TransactionReceipt(context.Background(), txHash)
				if err != nil {
					panic(fmt.Errorf("still cannot find receipt after retry: %v", err))
				}
			} else {
				panic(err)
			}
		}
		if receipt.Status != 1 {
			panic(errors.New("tx failed when distribution"))
		}
	}
	fmt.Println("finished distribution")
	return senderPks
}

func Transfer(opts *bind.TransactOpts, to common.Address, client *ethclient.Client) (*types.Transaction, error) {
	bc := bind.NewBoundContract(to, abi.ABI{}, client, client, nil)
	opts.GasLimit = 21000
	return bc.Transfer(opts)
}
