package utils

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func EcAdd(privateKeyA ecdsa.PrivateKey, privateKeyB ecdsa.PrivateKey) (*ecdsa.PrivateKey, error) {
	rawPrivateKeyAB := new(big.Int)
	rawPrivateKeyAB.Add(privateKeyA.D, privateKeyB.D)
	rawPrivateKeyAB.Mod(rawPrivateKeyAB, crypto.S256().Params().N)
	return ParsePrivateKey(*rawPrivateKeyAB)
}
