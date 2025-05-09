package solver

import (
	"crypto/ecdsa"
	"fmt"
	"infinity/miner/internal/contracts/PoW"
	"infinity/miner/internal/utils"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	magic = new(big.Int).SetBytes(common.FromHex("8888888888888888888888888888888888888888"))
)

type Solution struct {
	Nonce       big.Int
	PrivateKeyA ecdsa.PrivateKey
	PrivateKeyB ecdsa.PrivateKey
}

type Solver struct {
	ProblemCh    chan PoW.PoWNewProblem
	NumTries     uint64
	NumSolutions uint64
}

func NewSolver() *Solver {
	problemCh := make(chan PoW.PoWNewProblem)
	return &Solver{
		ProblemCh: problemCh,
	}
}

func logSolution(addressAB common.Address, privateKeyAB *ecdsa.PrivateKey) {
	f, err := os.OpenFile("solution.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	_, _ = f.WriteString(fmt.Sprintf("%s:%s\n", addressAB, privateKeyAB.D.Text(16)))
}

func trySolve(privateKeyA ecdsa.PrivateKey, difficulty big.Int) (*ecdsa.PrivateKey, error) {
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
		go logSolution(addressAB, privateKeyAB)
		return privateKeyB, nil
	}

	return nil, nil
}

func (s *Solver) Solve(solutionCh chan<- Solution) {
	var nonce big.Int
	var privateKeyA *ecdsa.PrivateKey
	var difficulty big.Int
	for {
		select {
		case problem := <-s.ProblemCh:
			nonce = *problem.Nonce
			privateKeyA, _ = utils.ParsePrivateKey(*problem.PrivateKeyA)
			difficulty = *problem.Difficulty
		default:
			if privateKeyA == nil {
				time.Sleep(time.Second / 10)
				continue
			}

			privateKeyB, _ := trySolve(*privateKeyA, difficulty)
			s.NumTries++
			if privateKeyB != nil {
				s.NumSolutions++
				solutionCh <- Solution{
					Nonce:       nonce,
					PrivateKeyA: *privateKeyA,
					PrivateKeyB: *privateKeyB,
				}
			}
		}
	}
}
