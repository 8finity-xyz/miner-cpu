package main

import (
	"infinity/miner/internal"
	"infinity/miner/internal/contracts/PoW"
	"infinity/miner/internal/listener"
	"infinity/miner/internal/solver"
	"infinity/miner/internal/submitter"
	"infinity/miner/internal/utils"
	"log"
	"log/slog"
	"math/big"
	"os"
	"runtime"
	"time"

	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
)

func main() {
	N := runtime.NumCPU()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	RPC := os.Getenv("INFINITY_RPC")
	if RPC == "" {
		log.Fatal("set INFINITY_RPC variable")
	}

	conn, err := ethclient.Dial(RPC)
	if err != nil {
		log.Fatal(err)
	}

	pow := PoW.NewPoW()
	instance := pow.Instance(
		conn,
		common.HexToAddress(internal.PoWAddress),
	)

	submitter := submitter.NewSubmitter(conn)
	problems, err := listener.SubscribeToProblems()
	if err != nil {
		log.Fatal("Cant subscribe for problems", err)
	}

	submitterBalance, err := submitter.GetBalance()
	if err != nil {
		log.Fatal("Cant get submitter balance")
	}
	log.Printf("∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞")
	log.Printf("Submitter address: %s", submitter.Address)
	log.Printf("Submitter balance: %f $S", new(big.Float).Quo(new(big.Float).SetInt(submitterBalance), big.NewFloat(params.Ether)))
	log.Printf("Ensure, that it have enough funds")
	log.Printf("∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞∞")

	solutionCh := make(chan solver.Solution, N)
	solvers := make([]*solver.Solver, N)
	for i := range len(solvers) {
		solvers[i] = solver.NewSolver()
		go solvers[i].Solve(solutionCh)
	}

	go func() {
		currentProblem, _ := bind2.Call(instance, nil, pow.PackCurrentProblem(), pow.UnpackCurrentProblem)
		problems <- PoW.PoWNewProblem{
			Nonce:       currentProblem.Arg0,
			PrivateKeyA: currentProblem.Arg1,
			Difficulty:  currentProblem.Arg2,
		}
	}()

	startTime := time.Now()
	ticker := time.NewTicker(time.Minute / 10)
	defer ticker.Stop()

	totalProblems := uint64(0)
	totalSubmits := uint64(0)

	var currentProblemNonce *big.Int
	for {
		select {
		case problem := <-problems:
			totalProblems += 1
			currentProblemNonce = problem.Nonce
			log.Printf("Got new problem: %s", common.BigToAddress(problem.Difficulty))
			for _, solver := range solvers {
				go func() {
					solver.ProblemCh <- problem
				}()
			}
		case solution := <-solutionCh:
			if solution.Nonce.Cmp(currentProblemNonce) != 0 {
				continue
			}

			privateKeyAB, _ := utils.EcAdd(solution.PrivateKeyA, solution.PrivateKeyB)
			isNextProblem, err := submitter.Submit(solution.PrivateKeyB, *privateKeyAB)
			if err == nil {
				totalSubmits += 1
			} else {
				slog.Debug("Submission failed", "err", err)
			}
			if isNextProblem {
				currentProblemNonce = big.NewInt(-1)
			}
		case <-ticker.C:
			totalTries := uint64(0)
			totalSolutions := uint64(0)
			for _, solver := range solvers {
				totalTries += solver.NumTries
				totalSolutions += solver.NumSolutions
			}

			log.Printf(
				"num problems: %d, num solutions: %d, hashrate: %s",
				totalProblems,
				totalSolutions,
				utils.FormatHashrate(totalTries, startTime),
			)
		}

	}
}
