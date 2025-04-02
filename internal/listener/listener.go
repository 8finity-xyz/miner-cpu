package listener

import (
	"errors"
	"infinity/miner/internal"
	"infinity/miner/internal/contracts/PoW"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SubscribeToProblems() (chan PoW.PoWNewProblem, error) {
	WS := os.Getenv("INFINITY_WS")
	if WS == "" {
		return nil, errors.New("set INFINITY_WS variable")
	}

	conn, err := ethclient.Dial(WS)
	if err != nil {
		return nil, err
	}

	pow := PoW.NewPoW()
	instance := pow.Instance(conn, common.HexToAddress(internal.PoWAddress))
	logs, sub, err := instance.WatchLogs(nil, "NewProblem")
	if err != nil {
		return nil, err
	}

	problems := make(chan PoW.PoWNewProblem)

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal("Subscription error:", err)
			case newPorblemLog := <-logs:
				newProblem, err := pow.UnpackNewProblemEvent(&newPorblemLog)
				if err != nil {
					log.Println("Parse error:", err)
					continue
				}
				problems <- *newProblem
			}
		}
	}()

	return problems, nil
}
