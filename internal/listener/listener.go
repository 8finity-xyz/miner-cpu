package listener

import (
	"context"
	"crypto/ecdsa"
	"infinity/miner/internal"
	"infinity/miner/internal/utils"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Problem struct {
	PrivateKeyA ecdsa.PrivateKey
	Difficulty  big.Int
}

func SubscribeToProblems(ch chan<- Problem) {
	WS := os.Getenv("INFINITY_WS")
	if WS == "" {
		log.Fatal("set INFINITY_WS variable")
	}

	client, err := ethclient.Dial(WS)
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{internal.PowAddress},
		Topics:    [][]common.Hash{{internal.PowAbi.Events["NewProblem"].ID}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal("Subscription failed:", err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Println("Subscription error:", err)
			case problemLog := <-logs:
				var parsedLog struct {
					PrivateKeyA *big.Int
					Difficulty  *big.Int
				}

				err := internal.PowAbi.UnpackIntoInterface(&parsedLog, "NewProblem", problemLog.Data)
				if err != nil {
					log.Println("Parse error:", err)
					continue
				}

				privateKeyA, err := utils.ParsePrivateKey(*parsedLog.PrivateKeyA)
				if err != nil {
					log.Println("Can't parse privateKeyA", err)
					continue
				}

				ch <- Problem{
					PrivateKeyA: *privateKeyA,
					Difficulty:  *parsedLog.Difficulty,
				}
			}
		}
	}()
}
