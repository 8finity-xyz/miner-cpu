package submitter

import (
	"context"
	"crypto/ecdsa"
	"infinity/miner/internal"
	"log"
	"log/slog"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Submitter struct {
	client     ethclient.Client
	chainId    *big.Int
	privateKey ecdsa.PrivateKey
	address    common.Address
	nonce      uint64
}

const gasLimit = uint64(1_000_000)

func NewSubmitter() *Submitter {
	RPC := os.Getenv("INFINITY_RPC")
	if RPC == "" {
		log.Fatal("set INFINITY_RPC variable")
	}

	client, err := ethclient.Dial(RPC)
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	PRIVATE_KEY := os.Getenv("INFINITY_PRIVATE_KEY")
	if PRIVATE_KEY == "" {
		log.Fatal("set INFINITY_PRIVATE_KEY variable")
	}

	privateKey, err := crypto.HexToECDSA(PRIVATE_KEY)
	if err != nil {
		log.Fatal(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	return &Submitter{client: *client, chainId: chainId, privateKey: *privateKey, address: address, nonce: nonce}
}

func (s *Submitter) Submit(privateKeyB ecdsa.PrivateKey, privateKeyAB ecdsa.PrivateKey) {
	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		slog.Error("Can't get gasprice", "error", err)
		return
	}

	// signature
	data := []byte{}
	var packed []byte
	packed = append(packed, s.address.Bytes()...)
	packed = append(packed, data...)
	digest := crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), crypto.Keccak256(packed))
	signature, err := crypto.Sign(digest, &privateKeyAB)
	if err != nil {
		slog.Error("Can't sign message", "error", err)
		return
	}
	signature[crypto.RecoveryIDOffset] += 27

	call, err := internal.PowAbi.Pack("submit", s.address, privateKeyB.PublicKey, signature, data)
	if err != nil {
		slog.Error("Can't prepare contract call", "error", err)
		return
	}

	tx := types.NewTransaction(s.nonce, internal.PowAddress, nil, gasLimit, gasPrice, call)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(s.chainId), &s.privateKey)
	if err != nil {
		slog.Error("Can't sign submit tx", "error", err)
		return
	}

	err = s.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		slog.Error("Can't send tx", "error", err)
		return
	}
	slog.Info("Submission transaction sended", "tx", signedTx.Hash())

	s.nonce++
}
