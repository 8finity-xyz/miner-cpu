package submitter

import (
	"context"
	"crypto/ecdsa"
	"infinity/miner/internal"
	"infinity/miner/internal/contracts/PoW"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Submitter struct {
	// submitter info
	Address    common.Address
	nonce      uint64
	privateKey ecdsa.PrivateKey
	conn       ethclient.Client

	// submit contract info
	chainId     *big.Int
	powAddress  common.Address
	pow         PoW.PoW
	powInstance bind.BoundContract
}

const gasLimit = uint64(1_000_000)

func NewSubmitter(conn *ethclient.Client) *Submitter {
	chainId, err := conn.NetworkID(context.Background())
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

	nonce, err := conn.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	pow := *PoW.NewPoW()
	powAddress := common.HexToAddress(internal.PoWAddress)
	powInstance := pow.Instance(conn, common.HexToAddress(internal.PoWAddress))

	return &Submitter{
		Address:     address,
		nonce:       nonce,
		privateKey:  *privateKey,
		conn:        *conn,
		chainId:     chainId,
		powAddress:  powAddress,
		pow:         pow,
		powInstance: *powInstance,
	}
}

func waitForTransactionReceipt(ctx context.Context, c ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second / 10)
	defer queryTicker.Stop()

	for {
		receipt, err := c.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (s *Submitter) GetBalance() (*big.Int, error) {
	return s.conn.PendingBalanceAt(context.Background(), s.Address)
}

func (s *Submitter) Submit(privateKeyB ecdsa.PrivateKey, privateKeyAB ecdsa.PrivateKey) (bool, error) {
	gasPrice, err := s.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return false, err
	}

	// signature
	data := common.Hex2Bytes(internal.Data)
	var packed []byte
	packed = append(packed, s.Address.Bytes()...)
	packed = append(packed, data...)
	digest := crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), crypto.Keccak256(packed))
	signature, err := crypto.Sign(digest, &privateKeyAB)
	if err != nil {
		return false, err
	}
	signature[crypto.RecoveryIDOffset] += 27

	tx, err := s.powInstance.Transact(
		&bind.TransactOpts{
			Nonce: big.NewInt(int64(s.nonce)),
			Signer: func(_ common.Address, t *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(t, types.NewEIP155Signer(s.chainId), &s.privateKey)
			},
			GasLimit: gasLimit,
			GasPrice: gasPrice,
		},
		"submit",
		s.Address,
		PoW.ECCPoint{X: privateKeyB.PublicKey.X, Y: privateKeyB.PublicKey.Y},
		signature,
		data,
	)
	if err != nil {
		return false, err
	}

	s.nonce++
	slog.Debug("Submission transaction sended", "tx", tx.Hash())

	receipt, err := waitForTransactionReceipt(context.Background(), s.conn, tx.Hash())
	if err != nil {
		return false, err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return true, err
	}

	for _, log := range receipt.Logs {
		if log.Address != s.powAddress {
			continue
		}
		newProblem, _ := s.pow.UnpackNewProblemEvent(log)
		if newProblem != nil {
			return true, nil
		}
	}
	return false, nil
}
