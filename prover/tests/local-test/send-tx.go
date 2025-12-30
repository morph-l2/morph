// send-tx.go - Send a transaction to the local node
// Usage: go run send-tx.go [rpc_url] [to_address] [value_in_ether]
//
// Uses Hardhat test account #1 which has funds in genesis:
// Address: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8
// Private Key: 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/params"
)

func main() {
	// Default values
	rpcURL := "http://127.0.0.1:8545"
	toAddress := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	valueEther := big.NewFloat(1.0) // 1 ether

	// Parse arguments
	if len(os.Args) > 1 {
		rpcURL = os.Args[1]
	}
	if len(os.Args) > 2 {
		toAddress = common.HexToAddress(os.Args[2])
	}
	if len(os.Args) > 3 {
		valueEther, _ = new(big.Float).SetString(os.Args[3])
	}

	// Hardhat test account #1
	privateKeyHex := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

	fmt.Println("=== Send Transaction ===")
	fmt.Printf("RPC: %s\n", rpcURL)

	// Connect to node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		os.Exit(1)
	}

	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Printf("Failed to parse private key: %v\n", err)
		os.Exit(1)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Failed to get public key")
		os.Exit(1)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("From: %s\n", fromAddress.Hex())
	fmt.Printf("To: %s\n", toAddress.Hex())
	fmt.Printf("Value: %s ETH\n", valueEther.String())

	// Get balance
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		fmt.Printf("Failed to get balance: %v\n", err)
		os.Exit(1)
	}
	balanceEther := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(params.Ether))
	fmt.Printf("Balance: %s ETH\n", balanceEther.String())

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Printf("Failed to get nonce: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Nonce: %d\n", nonce)

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("Failed to get gas price: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Gas Price: %s\n", gasPrice.String())

	// Get chain ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("Failed to get chain ID: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Chain ID: %s\n", chainID.String())

	// Convert ether to wei
	valueWei := new(big.Int)
	valueFloat := new(big.Float).Mul(valueEther, big.NewFloat(params.Ether))
	valueFloat.Int(valueWei)

	// Create transaction
	gasLimit := uint64(21000)
	tx := types.NewTransaction(nonce, toAddress, valueWei, gasLimit, gasPrice, nil)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Printf("Failed to sign tx: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("Sending transaction...")

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Printf("Failed to send tx: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("✅ Transaction sent!")
	fmt.Printf("Hash: %s\n", signedTx.Hash().Hex())
	fmt.Println("")
	fmt.Println("Now use dual-node-make-block.sh to include it in a block.")
}
