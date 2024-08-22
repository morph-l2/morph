package main

import (
	"bytes"
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/core/vm"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/blssignatures"
)

var n = 100

// the expected process on contracts:
// A. get the left side value of the equation e(`G1 point(message)`, `G2 point(agg public key)`)
// 1. map signed message to a point on G1, returns `128` bytes of G1 point encoding
//   - convert signed message to the required 64bytes input: append([32]bytes{}, keccak256(signedMessage))
//   - call bls12381MapG1 to map the input to a point on G1
//
// 2. aggregate bls public keys, produced the point on G2, returns `256` bytes of G2 point encoding
//   - call bls12381G2Add recursively to aggregate public keys
//   - the bls public key is supposed to be 256 bytes(encoded style)
//
// 3. paring the above G1(message) and G2(aggregated public keys) points, returns the `32` bytes.
//   - combine G1 and G2 bytes, get the 384 size bytes
//   - call the bls12381Pairing with this 384 input
//
// B. get the right side value of the equation e(`G1 point(agg signature)`, `G2 point(1)`)
// 1. acquire the aggregated signature from the calldata. the agg signature should be the `128` bytes of G1 point encoding.
//
// 2. acquire the value `one` which is encoded as a G2 point with 256 bytes.
//   - this value can be defined as the immutable value on the contracts: `0x00000000000000000000000000000000024aa2b2f08f0a91260805272dc51051c6e47ad4fa403b02b4510b647ae3d1770bac0326a805bbefd48056c8c121bdb80000000000000000000000000000000013e02b6052719f607dacd3a088274f65596bd0d09920b61ab5da61bbdc7f5049334cf11213945d57e5ac7d055d042b7e000000000000000000000000000000000ce5d527727d6e118cc9cdc6da2e351aadfd9baa8cbdd3a76d429a695160d12c923ac9cc3baca289e193548608b82801000000000000000000000000000000000606c4a02ea734cc32acd2b02bc28b99cb3e287e85a763af267492ab572e99ab3f370d275cec1da1aaa9075ff05f79be`
//
// 3. paring the above G1(agg signature) and G2(one) points, returns the `32` bytes.
//   - combine G1 and G2 bytes, get the 384 size bytes
//   - call the bls12381Pairing with this 384 input
//
// final step: compares the left side value and the right side value, true if they are equal
func main() {
	message := []byte("messageneedtobesigned")
	fmt.Println("message length: ", len(message))

	messageHash := crypto.Keccak256(message)
	var err error
	pubKeys := make([]blssignatures.PublicKey, n)
	privKeys := make([]blssignatures.PrivateKey, n)
	messageSigs := make([]blssignatures.Signature, n)
	for i := 0; i < n; i++ {
		pubKeys[i], privKeys[i], err = blssignatures.GenerateKeys()
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < n; i++ {
		messageSigs[i], err = blssignatures.SignMessage(privKeys[i], messageHash)
		if err != nil {
			panic(err)
		}
	}

	aggPubKeys := blssignatures.AggregatePublicKeys(pubKeys)
	aggSig := blssignatures.AggregateSignatures(messageSigs)

	verified, err := blssignatures.VerifySignature(aggSig, messageHash, aggPubKeys)
	if err != nil {
		panic(err)
	}
	fmt.Println("verified: ", verified)

	var prefix [32]byte
	messageInput := append(prefix[:], crypto.Keccak256(messageHash)...)

	var costGas uint64

	// hashToG1Curve
	bls12381MapG1 := vm.PrecompiledContractsBLS[common.BytesToAddress([]byte{17})]
	messagePointOnCurveBz, err := bls12381MapG1.Run(messageInput)
	if err != nil {
		panic(err)
	}
	costGas += bls12381MapG1.RequiredGas(messageInput)

	bls12381Pairing := vm.PrecompiledContractsBLS[common.BytesToAddress([]byte{16})]
	// left side add pair, message(point on g1) & public key(point on g2)
	byteRaw := blssignatures.PublicKeyToBytes(aggPubKeys)
	fmt.Println("aggPubKeys length: ", len(byteRaw))

	//the below behavior acts like: bls12381.NewG2().EncodePoint(aggPubKeys.key)
	byteRaw = byteRaw[1:]
	out := make([]byte, 256)
	// encode x
	copy(out[16:16+48], byteRaw[48:96])
	copy(out[80:80+48], byteRaw[:48])
	// encode y
	copy(out[144:144+48], byteRaw[144:])
	copy(out[208:208+48], byteRaw[96:144])

	leftBls12381PairingInput := append(messagePointOnCurveBz, out...)
	leftPairingResult, err := bls12381Pairing.Run(leftBls12381PairingInput)
	if err != nil {
		panic(err)
	}
	costGas += bls12381Pairing.RequiredGas(leftBls12381PairingInput)

	sigRawBytes := blssignatures.SignatureToBytes(aggSig)
	fmt.Println("sigRawBytes length: ", len(sigRawBytes))
	sigOut := bls12381.NewG1().EncodePoint(aggSig)

	oneEncoded := bls12381.NewG2().EncodePoint(bls12381.NewG2().One())
	fmt.Println("oneEncoded: ", hexutil.Encode(oneEncoded))
	rightBls12381PairingInput := append(sigOut, oneEncoded...)
	rightPairingResult, err := bls12381Pairing.Run(rightBls12381PairingInput)
	if err != nil {
		panic(err)
	}
	costGas += bls12381Pairing.RequiredGas(rightBls12381PairingInput)

	// e[message(point on g1) , agg public key(point on g2)] = e[agg sig(point on g1), one(point on g2)]
	fmt.Println("verified via precompiled bytes, result: ", bytes.Equal(leftPairingResult, rightPairingResult))
	fmt.Println("total gas cost for verification: ", costGas)
}

//func verifySignature2(sig Signature, message []byte, publicKey PublicKey, keyValidationMode bool) (bool, error) {
//	pointOnCurve, err := hashToG1Curve(message, keyValidationMode)
//	if err != nil {
//		return false, err
//	}
//
//	engine := bls12381.NewPairingEngine()
//	engine.Reset()
//	engine.AddPair(pointOnCurve, publicKey.key)
//	leftSide := engine.Result()
//	engine.AddPair(sig, engine.G2.One())
//	rightSide := engine.Result()
//	return leftSide.Equal(rightSide), nil
//}
