package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	rpcURLs := flag.String("rpc", os.Getenv("L2_RPC_URL"), "Comma-separated L2 RPC URLs (or L2_RPC_URL)")
	expectedChainID := flag.Uint64("chain-id", 0, "Expected chain ID; every endpoint must match")
	accountsFile := flag.String("accounts", "", "JSON file retaining generated sender private keys; created with mode 0600")
	senderNum := flag.Int("senders", 10, "Number of sender accounts")
	duration := flag.Duration("duration", time.Minute, "Maximum transfer duration after funding")
	interval := flag.Duration("interval", 800*time.Millisecond, "Delay between confirmed transfer rounds")
	timeout := flag.Duration("timeout", 2*time.Minute, "Maximum RPC and receipt wait per transaction")
	fundValue := flag.String("fund-value", "10000000000000000000", "Minimum balance funded to each sender, in wei")
	flag.Parse()
	if flag.NArg() != 0 || *rpcURLs == "" || *expectedChainID == 0 || *accountsFile == "" || *senderNum <= 0 || *duration <= 0 || *interval <= 0 || *timeout <= 0 {
		return errors.New("require -rpc, -chain-id, -accounts, and positive -senders, -duration, -interval, -timeout")
	}
	fundingKey, err := crypto.HexToECDSA(strings.TrimPrefix(os.Getenv("FUNDING_PRIVATE_KEY"), "0x"))
	if err != nil {
		return errors.New("set FUNDING_PRIVATE_KEY to a valid funding account private key")
	}
	value, ok := new(big.Int).SetString(*fundValue, 10)
	if !ok || value.Sign() <= 0 {
		return errors.New("-fund-value must be a positive integer in wei")
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	chainID := new(big.Int).SetUint64(*expectedChainID)
	var clients []*ethclient.Client
	var genesisHash common.Hash
	for _, url := range strings.Split(*rpcURLs, ",") {
		request, done := context.WithTimeout(ctx, *timeout)
		client, err := ethclient.DialContext(request, strings.TrimSpace(url))
		if err != nil {
			done()
			return err
		}
		defer client.Close()
		actual, err := client.ChainID(request)
		if err != nil {
			done()
			return err
		}
		if actual.Cmp(chainID) != 0 {
			done()
			return fmt.Errorf("RPC chain ID %s differs from expected %s", actual, chainID)
		}
		genesis, err := client.HeaderByNumber(request, big.NewInt(0))
		done()
		if err != nil {
			return fmt.Errorf("read RPC genesis: %w", err)
		}
		if len(clients) == 0 {
			genesisHash = genesis.Hash()
		} else if genesis.Hash() != genesisHash {
			return errors.New("RPC endpoints have different genesis hashes; no funding transactions were sent")
		}
		clients = append(clients, client)
	}
	senders, err := loadSenders(*accountsFile, *senderNum)
	if err != nil {
		return err
	}
	fundingAddress := crypto.PubkeyToAddress(fundingKey.PublicKey)
	for i, key := range append([]*ecdsa.PrivateKey{fundingKey}, senders...) {
		client := clients[0]
		if i > 0 {
			client = clients[(i-1)%len(clients)]
		}
		request, done := context.WithTimeout(ctx, *timeout)
		address := crypto.PubkeyToAddress(key.PublicKey)
		pending, err := client.PendingNonceAt(request, address)
		if err == nil {
			var confirmed uint64
			confirmed, err = client.NonceAt(request, address, nil)
			if err == nil && pending != confirmed {
				err = fmt.Errorf("account %s has pending transactions; resolve them before restarting", address)
			}
		}
		done()
		if err != nil {
			return err
		}
	}
	for _, key := range senders {
		address := crypto.PubkeyToAddress(key.PublicKey)
		request, done := context.WithTimeout(ctx, *timeout)
		balance, err := clients[0].BalanceAt(request, address, nil)
		done()
		if err != nil {
			return err
		}
		if balance.Cmp(value) < 0 {
			if err := sendConfirmed(ctx, *timeout, clients[0], chainID, fundingKey, address, new(big.Int).Sub(value, balance)); err != nil {
				return fmt.Errorf("fund sender %s: %w", address, err)
			}
		}
	}
	transferCtx, stop := context.WithTimeout(ctx, *duration)
	defer stop()
	for transferCtx.Err() == nil {
		var wg sync.WaitGroup
		failures := make(chan error, len(senders))
		for i, key := range senders {
			wg.Add(1)
			go func(client *ethclient.Client, key *ecdsa.PrivateKey) {
				defer wg.Done()
				if err := sendConfirmed(transferCtx, *timeout, client, chainID, key, fundingAddress, big.NewInt(10)); err != nil {
					failures <- err
				}
			}(clients[i%len(clients)], key)
		}
		wg.Wait()
		close(failures)
		for err := range failures {
			return err
		}
		select {
		case <-transferCtx.Done():
		case <-time.After(*interval):
		}
	}
	fmt.Println("completed; sender keys remain in", *accountsFile)
	return nil
}

func loadSenders(path string, count int) ([]*ecdsa.PrivateKey, error) {
	info, err := os.Lstat(path)
	if err == nil {
		if !info.Mode().IsRegular() || info.Mode().Perm()&0077 != 0 {
			return nil, errors.New("accounts file must be a regular file with permissions 0600; symbolic links are not accepted")
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	data, err := os.ReadFile(path)
	var encoded []string
	if errors.Is(err, os.ErrNotExist) {
		for i := 0; i < count; i++ {
			key, err := crypto.GenerateKey()
			if err != nil {
				return nil, err
			}
			encoded = append(encoded, common.Bytes2Hex(crypto.FromECDSA(key)))
		}
		data, err = json.MarshalIndent(encoded, "", "  ")
		if err != nil {
			return nil, err
		}
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
		if err != nil {
			return nil, err
		}
		_, writeErr := file.Write(data)
		if writeErr == nil {
			writeErr = file.Sync()
		}
		closeErr := file.Close()
		if writeErr != nil {
			return nil, writeErr
		}
		if closeErr != nil {
			return nil, closeErr
		}
	} else if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &encoded); err != nil {
		return nil, err
	}
	if len(encoded) != count {
		return nil, fmt.Errorf("accounts file contains %d keys, expected %d", len(encoded), count)
	}
	keys := make([]*ecdsa.PrivateKey, count)
	seen := make(map[common.Address]bool)
	for i, value := range encoded {
		key, err := crypto.HexToECDSA(strings.TrimPrefix(value, "0x"))
		if err != nil {
			return nil, fmt.Errorf("invalid sender key at index %d", i)
		}
		address := crypto.PubkeyToAddress(key.PublicKey)
		if seen[address] {
			return nil, fmt.Errorf("duplicate sender at index %d", i)
		}
		seen[address] = true
		keys[i] = key
	}
	return keys, nil
}

func sendConfirmed(ctx context.Context, timeout time.Duration, client *ethclient.Client, chainID *big.Int, key *ecdsa.PrivateKey, to common.Address, value *big.Int) error {
	request, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	opts, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return err
	}
	opts.Context, opts.Value = request, value
	tx, err := Transfer(opts, to, client)
	if err != nil {
		return fmt.Errorf("send from %s failed; check pending transactions before retrying: %w", opts.From, err)
	}
	fmt.Printf("submitted sender=%s nonce=%d tx=%s\n", opts.From, tx.Nonce(), tx.Hash())
	receipt, err := bind.WaitMined(request, client, tx)
	if err != nil {
		return fmt.Errorf("receipt for %s unavailable; check this transaction before restarting: %w", tx.Hash(), err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction %s failed", tx.Hash())
	}
	return nil
}

func Transfer(opts *bind.TransactOpts, to common.Address, client *ethclient.Client) (*types.Transaction, error) {
	bc := bind.NewBoundContract(to, abi.ABI{}, client, client, nil)
	opts.GasLimit = 21000
	return bc.Transfer(opts)
}
