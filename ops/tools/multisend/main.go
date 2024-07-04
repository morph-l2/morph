package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
)

var (
	chainId     *big.Int
	fundPrivKey = crypto.ToECDSAUnsafe(hexutil.MustDecode("0x3e4bde571b86929bf08e2aaad9a6a1882664cd5e65b96fff7d03e1c4e6dfa15c"))
	urls        = []string{"http://localhost:8545", "http://localhost:8645", "http://localhost:8745", "http://localhost:8845"}
	senderNum   = 20
	duration    = 120 * time.Minute
)

func main() {
	clients := make([]*ethclient.Client, 0)
	for _, url := range urls {
		client, err := ethclient.Dial(url)
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	getChainId, err := clients[0].ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	chainId = getChainId
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
			go func(s *ecdsa.PrivateKey, index int) {
				transactOpts, err := bind.NewKeyedTransactorWithChainID(s, chainId)
				if err != nil {
					panic(err)
				}
				transactOpts.Value = big.NewInt(10)
				client := clients[index%len(clients)]
				_, err = Transfer(transactOpts, to, client)
				if err != nil {
					fmt.Printf("[%d]error found when transfer: %v \n", index, err)
				} else {
					fmt.Printf("[%d]sent tx done\n", index)
				}
			}(sender, i)
		}
		count++
		time.Sleep(50 * time.Millisecond)
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
