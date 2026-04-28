package main

import (
	"fmt"

	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
)

func main() {
	pv := privval.GenFilePV("", "")
	pvJsbz, err := tmjson.Marshal(pv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-----------------------generating content for priv_validator_key.json-----------------------\n")
	fmt.Printf("%v \n", string(pvJsbz))
	pubKey, _ := pv.GetPubKey()
	fmt.Printf("hex format public key: 0x%x \n", pubKey.Bytes())

	fmt.Println()
	fmt.Printf("--------------------------generating content for node_key.json--------------------------\n")
	privKey := ed25519.GenPrivKey()
	nodeKey := &p2p.NodeKey{
		PrivKey: privKey,
	}
	nodeKeyJsBz, err := tmjson.Marshal(nodeKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", string(nodeKeyJsBz))
	fmt.Printf("nodeID: %v \n", nodeKey.ID())

	fmt.Println()
	fmt.Printf("--------------------------generating a new ETH account--------------------------\n")
	ethPrivKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	ethAddress := crypto.PubkeyToAddress(ethPrivKey.PublicKey)
	fmt.Printf("eth account private key: 0x%x \n", crypto.FromECDSA(ethPrivKey))
	fmt.Printf("eth account address: %s", ethAddress.Hex())
}
