// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PoW

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ECCPoint is an auto generated low-level Go binding around an user-defined struct.
type ECCPoint struct {
	X *big.Int
	Y *big.Int
}

// PoWMetaData contains all meta data concerning the PoW contract.
var PoWMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressAB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"BadSolution\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"privateKeyA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"difficulty\",\"type\":\"uint160\"}],\"name\":\"NewProblem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReward\",\"type\":\"uint256\"}],\"name\":\"RewardReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addressAB\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INFINITY\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentProblem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numSubmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"privateKeyA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"problemNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startReward\",\"type\":\"uint256\"}],\"name\":\"startMining\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structECCPoint\",\"name\":\"publicKeyB\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signatureAB\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "PoW",
}

// PoW is an auto generated Go binding around an Ethereum contract.
type PoW struct {
	abi abi.ABI
}

// NewPoW creates a new instance of PoW.
func NewPoW() *PoW {
	parsed, err := PoWMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &PoW{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *PoW) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackINFINITY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f48f44b.
//
// Solidity: function INFINITY() view returns(address)
func (poW *PoW) PackINFINITY() []byte {
	enc, err := poW.abi.Pack("INFINITY")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackINFINITY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f48f44b.
//
// Solidity: function INFINITY() view returns(address)
func (poW *PoW) UnpackINFINITY(data []byte) (common.Address, error) {
	out, err := poW.abi.Unpack("INFINITY", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (poW *PoW) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := poW.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (poW *PoW) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := poW.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// PackCurrentProblem is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88f116d0.
//
// Solidity: function currentProblem() view returns(uint256, uint256, uint160)
func (poW *PoW) PackCurrentProblem() []byte {
	enc, err := poW.abi.Pack("currentProblem")
	if err != nil {
		panic(err)
	}
	return enc
}

// CurrentProblemOutput serves as a container for the return parameters of contract
// method CurrentProblem.
type CurrentProblemOutput struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Arg2 *big.Int
}

// UnpackCurrentProblem is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88f116d0.
//
// Solidity: function currentProblem() view returns(uint256, uint256, uint160)
func (poW *PoW) UnpackCurrentProblem(data []byte) (CurrentProblemOutput, error) {
	out, err := poW.abi.Unpack("currentProblem", data)
	outstruct := new(CurrentProblemOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Arg1 = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.Arg2 = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, err

}

// PackDifficulty is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x19cae462.
//
// Solidity: function difficulty() view returns(uint160)
func (poW *PoW) PackDifficulty() []byte {
	enc, err := poW.abi.Pack("difficulty")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDifficulty is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x19cae462.
//
// Solidity: function difficulty() view returns(uint160)
func (poW *PoW) UnpackDifficulty(data []byte) (*big.Int, error) {
	out, err := poW.abi.Unpack("difficulty", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (poW *PoW) PackInitialize(owner common.Address) []byte {
	enc, err := poW.abi.Pack("initialize", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize2 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x472abf68.
//
// Solidity: function initialize2() returns()
func (poW *PoW) PackInitialize2() []byte {
	enc, err := poW.abi.Pack("initialize2")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackNumSubmissions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x77e19824.
//
// Solidity: function numSubmissions() view returns(uint256)
func (poW *PoW) PackNumSubmissions() []byte {
	enc, err := poW.abi.Pack("numSubmissions")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumSubmissions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x77e19824.
//
// Solidity: function numSubmissions() view returns(uint256)
func (poW *PoW) UnpackNumSubmissions(data []byte) (*big.Int, error) {
	out, err := poW.abi.Unpack("numSubmissions", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (poW *PoW) PackOwner() []byte {
	enc, err := poW.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (poW *PoW) UnpackOwner(data []byte) (common.Address, error) {
	out, err := poW.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (poW *PoW) PackPaused() []byte {
	enc, err := poW.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (poW *PoW) UnpackPaused(data []byte) (bool, error) {
	out, err := poW.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackPrivateKeyA is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4ffbbc8.
//
// Solidity: function privateKeyA() view returns(uint256)
func (poW *PoW) PackPrivateKeyA() []byte {
	enc, err := poW.abi.Pack("privateKeyA")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPrivateKeyA is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb4ffbbc8.
//
// Solidity: function privateKeyA() view returns(uint256)
func (poW *PoW) UnpackPrivateKeyA(data []byte) (*big.Int, error) {
	out, err := poW.abi.Unpack("privateKeyA", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProblemNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9ef1500.
//
// Solidity: function problemNonce() view returns(uint256)
func (poW *PoW) PackProblemNonce() []byte {
	enc, err := poW.abi.Pack("problemNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProblemNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb9ef1500.
//
// Solidity: function problemNonce() view returns(uint256)
func (poW *PoW) UnpackProblemNonce(data []byte) (*big.Int, error) {
	out, err := poW.abi.Unpack("problemNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (poW *PoW) PackProxiableUUID() []byte {
	enc, err := poW.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (poW *PoW) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := poW.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, err
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (poW *PoW) PackRenounceOwnership() []byte {
	enc, err := poW.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackReward is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x228cb733.
//
// Solidity: function reward() view returns(uint256)
func (poW *PoW) PackReward() []byte {
	enc, err := poW.abi.Pack("reward")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackReward is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x228cb733.
//
// Solidity: function reward() view returns(uint256)
func (poW *PoW) UnpackReward(data []byte) (*big.Int, error) {
	out, err := poW.abi.Unpack("reward", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackStartMining is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47b272c0.
//
// Solidity: function startMining(uint256 startReward) returns()
func (poW *PoW) PackStartMining(startReward *big.Int) []byte {
	enc, err := poW.abi.Pack("startMining", startReward)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSubmit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76fbe328.
//
// Solidity: function submit(address recipient, (uint256,uint256) publicKeyB, bytes signatureAB, bytes data) returns()
func (poW *PoW) PackSubmit(recipient common.Address, publicKeyB ECCPoint, signatureAB []byte, data []byte) []byte {
	enc, err := poW.abi.Pack("submit", recipient, publicKeyB, signatureAB, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (poW *PoW) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := poW.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (poW *PoW) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := poW.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PoWInitialized represents a Initialized event raised by the PoW contract.
type PoWInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const PoWInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (PoWInitialized) ContractEventName() string {
	return PoWInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (poW *PoW) UnpackInitializedEvent(log *types.Log) (*PoWInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWInitialized)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWNewProblem represents a NewProblem event raised by the PoW contract.
type PoWNewProblem struct {
	Nonce       *big.Int
	PrivateKeyA *big.Int
	Difficulty  *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const PoWNewProblemEventName = "NewProblem"

// ContractEventName returns the user-defined event name.
func (PoWNewProblem) ContractEventName() string {
	return PoWNewProblemEventName
}

// UnpackNewProblemEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewProblem(uint256 nonce, uint256 privateKeyA, uint160 difficulty)
func (poW *PoW) UnpackNewProblemEvent(log *types.Log) (*PoWNewProblem, error) {
	event := "NewProblem"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWNewProblem)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWOwnershipTransferred represents a OwnershipTransferred event raised by the PoW contract.
type PoWOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const PoWOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (PoWOwnershipTransferred) ContractEventName() string {
	return PoWOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (poW *PoW) UnpackOwnershipTransferredEvent(log *types.Log) (*PoWOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWPaused represents a Paused event raised by the PoW contract.
type PoWPaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PoWPausedEventName = "Paused"

// ContractEventName returns the user-defined event name.
func (PoWPaused) ContractEventName() string {
	return PoWPausedEventName
}

// UnpackPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Paused(address account)
func (poW *PoW) UnpackPausedEvent(log *types.Log) (*PoWPaused, error) {
	event := "Paused"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWPaused)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWRewardReduced represents a RewardReduced event raised by the PoW contract.
type PoWRewardReduced struct {
	NewReward *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const PoWRewardReducedEventName = "RewardReduced"

// ContractEventName returns the user-defined event name.
func (PoWRewardReduced) ContractEventName() string {
	return PoWRewardReducedEventName
}

// UnpackRewardReducedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardReduced(uint256 newReward)
func (poW *PoW) UnpackRewardReducedEvent(log *types.Log) (*PoWRewardReduced, error) {
	event := "RewardReduced"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWRewardReduced)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWSubmission represents a Submission event raised by the PoW contract.
type PoWSubmission struct {
	Miner     common.Address
	AddressAB common.Address
	Reward    *big.Int
	Data      []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const PoWSubmissionEventName = "Submission"

// ContractEventName returns the user-defined event name.
func (PoWSubmission) ContractEventName() string {
	return PoWSubmissionEventName
}

// UnpackSubmissionEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Submission(address indexed miner, address indexed addressAB, uint256 reward, bytes data)
func (poW *PoW) UnpackSubmissionEvent(log *types.Log) (*PoWSubmission, error) {
	event := "Submission"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWSubmission)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWUnpaused represents a Unpaused event raised by the PoW contract.
type PoWUnpaused struct {
	Account common.Address
	Raw     *types.Log // Blockchain specific contextual infos
}

const PoWUnpausedEventName = "Unpaused"

// ContractEventName returns the user-defined event name.
func (PoWUnpaused) ContractEventName() string {
	return PoWUnpausedEventName
}

// UnpackUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Unpaused(address account)
func (poW *PoW) UnpackUnpausedEvent(log *types.Log) (*PoWUnpaused, error) {
	event := "Unpaused"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWUnpaused)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// PoWUpgraded represents a Upgraded event raised by the PoW contract.
type PoWUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const PoWUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (PoWUpgraded) ContractEventName() string {
	return PoWUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (poW *PoW) UnpackUpgradedEvent(log *types.Log) (*PoWUpgraded, error) {
	event := "Upgraded"
	if log.Topics[0] != poW.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PoWUpgraded)
	if len(log.Data) > 0 {
		if err := poW.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range poW.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (poW *PoW) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], poW.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return poW.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["BadSignature"].ID.Bytes()[:4]) {
		return poW.UnpackBadSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["BadSolution"].ID.Bytes()[:4]) {
		return poW.UnpackBadSolutionError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return poW.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return poW.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return poW.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return poW.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return poW.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["EnforcedPause"].ID.Bytes()[:4]) {
		return poW.UnpackEnforcedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["ExpectedPause"].ID.Bytes()[:4]) {
		return poW.UnpackExpectedPauseError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return poW.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return poW.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return poW.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return poW.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return poW.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return poW.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], poW.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return poW.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// PoWAddressEmptyCode represents a AddressEmptyCode error raised by the PoW contract.
type PoWAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func PoWAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (poW *PoW) UnpackAddressEmptyCodeError(raw []byte) (*PoWAddressEmptyCode, error) {
	out := new(PoWAddressEmptyCode)
	if err := poW.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWBadSignature represents a BadSignature error raised by the PoW contract.
type PoWBadSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadSignature()
func PoWBadSignatureErrorID() common.Hash {
	return common.HexToHash("0x5cd5d2335541c4f2ed05fbe44f397e8b79f8e2333157122d2dab06e378ef7685")
}

// UnpackBadSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadSignature()
func (poW *PoW) UnpackBadSignatureError(raw []byte) (*PoWBadSignature, error) {
	out := new(PoWBadSignature)
	if err := poW.abi.UnpackIntoInterface(out, "BadSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWBadSolution represents a BadSolution error raised by the PoW contract.
type PoWBadSolution struct {
	AddressAB common.Address
	Target    common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadSolution(address addressAB, address target)
func PoWBadSolutionErrorID() common.Hash {
	return common.HexToHash("0x4245eb38592ebd3177c8ff2a9c12dbcf51091250327a98be62cb43252272d1f5")
}

// UnpackBadSolutionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadSolution(address addressAB, address target)
func (poW *PoW) UnpackBadSolutionError(raw []byte) (*PoWBadSolution, error) {
	out := new(PoWBadSolution)
	if err := poW.abi.UnpackIntoInterface(out, "BadSolution", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the PoW contract.
type PoWECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func PoWECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (poW *PoW) UnpackECDSAInvalidSignatureError(raw []byte) (*PoWECDSAInvalidSignature, error) {
	out := new(PoWECDSAInvalidSignature)
	if err := poW.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the PoW contract.
type PoWECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func PoWECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (poW *PoW) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*PoWECDSAInvalidSignatureLength, error) {
	out := new(PoWECDSAInvalidSignatureLength)
	if err := poW.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the PoW contract.
type PoWECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func PoWECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (poW *PoW) UnpackECDSAInvalidSignatureSError(raw []byte) (*PoWECDSAInvalidSignatureS, error) {
	out := new(PoWECDSAInvalidSignatureS)
	if err := poW.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the PoW contract.
type PoWERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func PoWERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (poW *PoW) UnpackERC1967InvalidImplementationError(raw []byte) (*PoWERC1967InvalidImplementation, error) {
	out := new(PoWERC1967InvalidImplementation)
	if err := poW.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWERC1967NonPayable represents a ERC1967NonPayable error raised by the PoW contract.
type PoWERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func PoWERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (poW *PoW) UnpackERC1967NonPayableError(raw []byte) (*PoWERC1967NonPayable, error) {
	out := new(PoWERC1967NonPayable)
	if err := poW.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWEnforcedPause represents a EnforcedPause error raised by the PoW contract.
type PoWEnforcedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EnforcedPause()
func PoWEnforcedPauseErrorID() common.Hash {
	return common.HexToHash("0xd93c0665d6c96d04a8f174024fc4ddd66c250604aff22bbec808de86dd3637e3")
}

// UnpackEnforcedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EnforcedPause()
func (poW *PoW) UnpackEnforcedPauseError(raw []byte) (*PoWEnforcedPause, error) {
	out := new(PoWEnforcedPause)
	if err := poW.abi.UnpackIntoInterface(out, "EnforcedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWExpectedPause represents a ExpectedPause error raised by the PoW contract.
type PoWExpectedPause struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExpectedPause()
func PoWExpectedPauseErrorID() common.Hash {
	return common.HexToHash("0x8dfc202bcfe9a735b559bee70674422512bc5c30f687046ae8778315fb81da44")
}

// UnpackExpectedPauseError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExpectedPause()
func (poW *PoW) UnpackExpectedPauseError(raw []byte) (*PoWExpectedPause, error) {
	out := new(PoWExpectedPause)
	if err := poW.abi.UnpackIntoInterface(out, "ExpectedPause", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWFailedCall represents a FailedCall error raised by the PoW contract.
type PoWFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func PoWFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (poW *PoW) UnpackFailedCallError(raw []byte) (*PoWFailedCall, error) {
	out := new(PoWFailedCall)
	if err := poW.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWInvalidInitialization represents a InvalidInitialization error raised by the PoW contract.
type PoWInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func PoWInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (poW *PoW) UnpackInvalidInitializationError(raw []byte) (*PoWInvalidInitialization, error) {
	out := new(PoWInvalidInitialization)
	if err := poW.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWNotInitializing represents a NotInitializing error raised by the PoW contract.
type PoWNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func PoWNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (poW *PoW) UnpackNotInitializingError(raw []byte) (*PoWNotInitializing, error) {
	out := new(PoWNotInitializing)
	if err := poW.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the PoW contract.
type PoWOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func PoWOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (poW *PoW) UnpackOwnableInvalidOwnerError(raw []byte) (*PoWOwnableInvalidOwner, error) {
	out := new(PoWOwnableInvalidOwner)
	if err := poW.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the PoW contract.
type PoWOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func PoWOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (poW *PoW) UnpackOwnableUnauthorizedAccountError(raw []byte) (*PoWOwnableUnauthorizedAccount, error) {
	out := new(PoWOwnableUnauthorizedAccount)
	if err := poW.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the PoW contract.
type PoWUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func PoWUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (poW *PoW) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*PoWUUPSUnauthorizedCallContext, error) {
	out := new(PoWUUPSUnauthorizedCallContext)
	if err := poW.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// PoWUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the PoW contract.
type PoWUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func PoWUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (poW *PoW) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*PoWUUPSUnsupportedProxiableUUID, error) {
	out := new(PoWUUPSUnsupportedProxiableUUID)
	if err := poW.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
