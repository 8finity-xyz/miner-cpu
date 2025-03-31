package solver

import (
	"crypto/ecdsa"
	"infinity/miner/internal/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	magic = new(big.Int).SetBytes(common.FromHex("8888888888888888888888888888888888888888"))
)

func Solve(privateKeyA ecdsa.PrivateKey, difficulty big.Int) (*ecdsa.PrivateKey, error) {
	privateKeyB, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	privateKeyAB, err := utils.EcAdd(privateKeyA, *privateKeyB)
	if err != nil {
		return nil, err
	}
	addressAB := crypto.PubkeyToAddress(privateKeyAB.PublicKey)

	result := new(big.Int)
	result.Xor(magic, addressAB.Big())
	if result.Cmp(&difficulty) < 0 {
		return privateKeyB, nil
	}

	return nil, nil
}
