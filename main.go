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

	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

	var currentProblemNonce *big.Int
	for {
		select {
		case problem := <-problems:
			currentProblemNonce = problem.Nonce
			slog.Info("Got new problem", "difficulty", common.BigToAddress(problem.Difficulty))
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
			slog.Info(
				"Submitting solution",
				"solution", crypto.PubkeyToAddress(privateKeyAB.PublicKey),
			)
			isNextProblem := submitter.Submit(solution.PrivateKeyB, *privateKeyAB)
			if isNextProblem {
				currentProblemNonce = big.NewInt(-1)
			}
		}
	}
}
