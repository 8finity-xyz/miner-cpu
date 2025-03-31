package utils

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func ParsePrivateKey(rawPrivateKey big.Int) (*ecdsa.PrivateKey, error) {
	bytesPrivateKey := make([]byte, 32)
	rawPrivateKey.FillBytes(bytesPrivateKey)

	privateKey, err := crypto.ToECDSA(bytesPrivateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
