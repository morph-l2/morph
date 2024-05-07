package oracle

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/morph-l2/morph/oracle/config"
	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	tmhttp "github.com/tendermint/tendermint/rpc/client/http"
	"golang.org/x/crypto/sha3"
)

func testNewOracleClient(t *testing.T) *Oracle {
	l1Client, err := ethclient.Dial("http://localhost:9545")
	require.NoError(t, err)
	var secret [32]byte
	jwtSecret := common.FromHex(strings.TrimSpace("688f5d737bad920bdfb2fc2f488d6b6209eebda1dae949a8de91398d932c517a"))
	require.True(t, len(jwtSecret) == 32)
	copy(secret[:], jwtSecret)
	l2Client, err := ethclient.Dial("http://localhost:8545")
	require.NoError(t, err)
	l2Staking, err := bindings.NewL2Staking(predeploys.L2StakingAddr, l2Client)
	require.NoError(t, err)
	tmClient, err := tmhttp.New("http://localhost:26657", "http://localhost:26656")
	require.NoError(t, err)
	record, err := bindings.NewRecord(predeploys.RecordAddr, l2Client)
	require.NoError(t, err)
	cfg := &config.Config{
		StartBlock: 1,
	}
	rollup, err := bindings.NewRollup(common.HexToAddress("0x0165878a594ca255338adfa4d48449f69242eb8f"), l1Client)
	require.NoError(t, err)
	hex := strings.TrimPrefix("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	return &Oracle{
		l1Client:    l1Client,
		l2Client:    l2Client,
		rollup:      rollup,
		l2Staking:   l2Staking,
		record:      record,
		cfg:         cfg,
		rewardEpoch: defaultRewardEpoch,
		ctx:         context.Background(),
		TmClient:    tmClient,
		privKey:     privKey,
	}

}

func TestFindStartBlock(t *testing.T) {
	o := testNewOracleClient(t)
	bn, err := o.findStartBlock(1, 3476, 1714435200)
	require.NoError(t, err)
	fmt.Println("bn:", bn)
}

func TestSetStartBlock(t *testing.T) {
	o := testNewOracleClient(t)
	rewardStartTime, err := o.l2Staking.RewardStartTime(nil)
	require.NoError(t, err)
	fmt.Println("rewardStartTime:", rewardStartTime.Uint64())
	hex := strings.TrimPrefix("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	o.privKey = privKey

	o.setStartBlock()
}

func TestGetStakers(t *testing.T) {
	o := testNewOracleClient(t)
	stakers, err := o.l2Staking.GetStakers(nil)
	require.NoError(t, err)
	for _, staker := range stakers {
		fmt.Println("staker addr:", staker.Addr)
		commission, err := o.l2Staking.Commissions(nil, staker.Addr)
		require.NoError(t, err)
		fmt.Println("commission:", commission)
	}
}

func TestQueryRewardStartTime(t *testing.T) {
	o := testNewOracleClient(t)
	rewardStartTime, err := o.l2Staking.RewardStartTime(nil)
	require.NoError(t, err)
	fmt.Println("rewardStartTime:", rewardStartTime.Uint64())
	fmt.Println(time.Now().Unix())
	fmt.Println(rewardStartTime.Uint64())

}

func TestSetRewardStart(t *testing.T) {
	o := testNewOracleClient(t)
	chainID, err := o.l2Client.ChainID(o.ctx)
	if err != nil {

	}
	hex := strings.TrimPrefix("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	o.privKey = privKey
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		return
	}
	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))
	recordOwner, err := o.record.Owner(nil)
	require.NoError(t, err)
	fmt.Println(recordOwner.Hex())
	owner, err := o.l2Staking.Owner(nil)
	require.NoError(t, err)
	fmt.Println(owner.Hex())
	candidateNumber, err := o.l2Staking.CandidateNumber(nil)
	require.NoError(t, err)
	fmt.Println("candidateNumber:", candidateNumber)
	// 1714435200
	require.NoError(t, err)
	tx, err := o.l2Staking.StartReward(opts)
	require.NoError(t, err)
	fmt.Println("ops from:", opts.From)
	signedTx, err := opts.Signer(opts.From, tx)
	if err != nil {
		return
	}
	fmt.Println(signedTx)

	err = o.l2Client.SendTransaction(o.ctx, signedTx)
	if err != nil {

	}
	receipt, err := o.l2Client.TransactionReceipt(o.ctx, tx.Hash())
	require.NoError(t, err)
	if receipt.Status != types.ReceiptStatusSuccessful {
		fmt.Println("set stark block failed")
	}
	require.NoError(t, err)
	fmt.Println("Tx:", tx.Hash())
}

func TestUpdateRewardTime(t *testing.T) {
	o := testNewOracleClient(t)
	chainID, err := o.l2Client.ChainID(o.ctx)
	if err != nil {

	}
	hex := strings.TrimPrefix("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	TransferMorphToken(o.l2Client, privKey, "0x783698dCDEBdc96785c5c60ED96113612bA09c2b", 0)

	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		return
	}
	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))

	tx1, err := o.l2Staking.UpdateRewardStartTime(opts, big.NewInt(1714435200))
	require.NoError(t, err)
	signedTx1, err := opts.Signer(opts.From, tx1)
	require.NoError(t, err)
	err = o.l2Client.SendTransaction(o.ctx, signedTx1)
	require.NoError(t, err)
	fmt.Println(signedTx1)
	fmt.Println("Tx:", tx1.Hash())
}

func TestDelegateStake(t *testing.T) {
	o := testNewOracleClient(t)
	chainID, err := o.l2Client.ChainID(o.ctx)
	require.NoError(t, err)
	hex := strings.TrimPrefix("0xd99870855d97327d20c666abc78588f1449b1fac76ed0c86c1afb9ce2db85f32", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		return
	}
	opts.Nonce = big.NewInt(int64(nonce))
	balance, err := o.l2Client.BalanceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	require.NoError(t, err)
	fmt.Printf("address:%v,balance:%v\n", crypto.PubkeyToAddress(privKey.PublicKey), balance)
	//stakerRankings
	stakerBanking, err := o.l2Staking.StakerRankings(nil, crypto.PubkeyToAddress(privKey.PublicKey))
	fmt.Println("stakerBanking:", stakerBanking)
	tx, err := o.l2Staking.DelegateStake(opts, crypto.PubkeyToAddress(privKey.PublicKey), big.NewInt(20))
	require.NoError(t, err)
	fmt.Println("tx hash:", tx.Hash())

}

func TestTransfer(t *testing.T) {
	o := testNewOracleClient(t)
	chainID, err := o.l2Client.ChainID(o.ctx)
	if err != nil {

	}
	hex := strings.TrimPrefix("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "0x")
	privKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		panic(err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(privKey.PublicKey), nil)
	if err != nil {
		return
	}
	opts.NoSend = true
	opts.Nonce = big.NewInt(int64(nonce))
	gasPrice, err := o.l2Client.SuggestGasPrice(context.Background())
	require.NoError(t, err)

	rawTx := types.NewTransaction(nonce, common.HexToAddress("0x783698dCDEBdc96785c5c60ED96113612bA09c2b"), big.NewInt(9000000000000000000), 21000, gasPrice, []byte{})

	signedTx, err := types.SignTx(rawTx, types.NewEIP155Signer(chainID), privKey)
	require.NoError(t, err)
	err = o.l2Client.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)

}

func TestMorphTokenBalance(t *testing.T) {
	o := testNewOracleClient(t)
	morphToken, err := bindings.NewMorphToken(predeploys.MorphTokenAddr, o.l2Client)
	require.NoError(t, err)
	total, err := morphToken.TotalSupply(nil)
	require.NoError(t, err)
	fmt.Println("total:", total)
	morphTokenOwner, err := morphToken.Owner(nil)
	require.NoError(t, err)
	fmt.Println("morphTokenOwner:", morphTokenOwner.Hex())
	balance, err := morphToken.BalanceOf(nil, common.HexToAddress("0x783698dCDEBdc96785c5c60ED96113612bA09c2b"))
	require.NoError(t, err)
	fmt.Println("balance:", balance)
}

func TransferMorphToken(client *ethclient.Client, privateKey *ecdsa.PrivateKey, toAddr string, nonceQ uint64) {
	// ferc1
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	n, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	if nonceQ == 0 {
		nonceQ = n
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	gasPrice = gasPrice.Div(gasPrice.Mul(gasPrice, big.NewInt(10)), big.NewInt(10))
	toAddress := common.HexToAddress(toAddr)
	amount := new(big.Int)
	// d 6
	amount.SetString("20", 10) // sets the value to 1000 tokens, in the token denomination
	data := GenerateTransferErc20Data(toAddress, amount)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &predeploys.MorphTokenAddr,
		Data: data,
	})
	if err != nil {
		fmt.Println("EstimateGas error:", err)
		log.Fatal(err)
	}
	fmt.Println(gasLimit) // 23256
	tx := types.NewTransaction(nonceQ, predeploys.MorphTokenAddr, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("SignTx error:", err)
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("SendTransaction error:", err)
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}

func GenerateTransferErc20Data(toAddress common.Address, amount *big.Int) []byte {
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return data
}

func TestOracle_GetBatchSubmission(t *testing.T) {
	o := testNewOracleClient(t)
	startRewardEpochIndex, err := o.record.NextRewardEpochIndex(nil)
	require.NoError(t, err)
	startHeight, err := o.getNextHeight()
	recordRewardEpochInfo, err := o.getRewardEpochs(startRewardEpochIndex, startHeight)
	require.NoError(t, err)
	fmt.Println(recordRewardEpochInfo)
}

func TestGetStartBlock(t *testing.T) {
	o := testNewOracleClient(t)
	LatestBlock, err := o.record.LatestRewardEpochBlock(nil)
	require.NoError(t, err)
	fmt.Println(LatestBlock)
}

func TestOracle_SyncRewardEpoch(t *testing.T) {
	o := testNewOracleClient(t)
	err := o.syncRewardEpoch()
	require.NoError(t, err)
}
