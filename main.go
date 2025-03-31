package main

import (
	"crypto/ecdsa"
	"infinity/miner/internal/listener"
	"infinity/miner/internal/solver"
	"infinity/miner/internal/submitter"
	"infinity/miner/internal/utils"
	"log/slog"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func worker(problemCh <-chan *listener.Problem, solutionCh chan<- Solution) {
	var problem *listener.Problem
	for {
		select {
		case problem = <-problemCh:
		default:
			if problem == nil {
				time.Sleep(time.Second / 10)
				continue
			}

			privateKeyB, _ := solver.Solve(problem.PrivateKeyA, problem.Difficulty)
			if privateKeyB != nil {
				solutionCh <- Solution{
					PrivateKeyA: problem.PrivateKeyA,
					PrivateKeyB: *privateKeyB,
				}
				problem = nil
			}
		}
	}
}

type Solution struct {
	PrivateKeyA ecdsa.PrivateKey
	PrivateKeyB ecdsa.PrivateKey
}

func main() {
	const N = 10

	submitter := submitter.NewSubmitter()
	problems := make(chan listener.Problem)
	listener.SubscribeToProblems(problems)

	// We propogate problem for submitter and for each worker (1 + N)
	problemsCh := make([]chan *listener.Problem, 1+N)
	for i := range len(problemsCh) {
		problemsCh[i] = make(chan *listener.Problem, 1)
	}
	go func() {
		for {
			problem := <-problems
			slog.Info("Got new problem", "difficulty", common.BigToAddress(&problem.Difficulty))
			for i := range len(problemsCh) {
				problemsCh[i] <- &problem
			}
		}
	}()

	// workers stream solutions to channel, when found it
	solutionCh := make(chan Solution, N)
	for i := range N {
		go worker(problemsCh[i+1], solutionCh)
	}

	var currentProblem *listener.Problem
	for {
		select {
		case currentProblem = <-problemsCh[0]:
		case solution := <-solutionCh:
			if currentProblem == nil || solution.PrivateKeyA != currentProblem.PrivateKeyA {
				continue
			}

			privateKeyAB, _ := utils.EcAdd(solution.PrivateKeyA, solution.PrivateKeyB)
			slog.Info(
				"Submitting solution",
				"solution", crypto.PubkeyToAddress(privateKeyAB.PublicKey),
			)
			for i := range N {
				problemsCh[i] <- nil
			}

			submitter.Submit(solution.PrivateKeyB, *privateKeyAB)
			currentProblem = nil
		}
	}
}
