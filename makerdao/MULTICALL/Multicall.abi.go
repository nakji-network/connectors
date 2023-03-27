// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Multicall

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MulticallCall is an auto generated low-level Go binding around an user-defined struct.
type MulticallCall struct {
	Target   common.Address
	CallData []byte
}

// MULTICALLABI is the input ABI used to generate the binding from.
const MULTICALLABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MULTICALL is an auto generated Go binding around an Ethereum contract.
type MULTICALL struct {
	MULTICALLCaller     // Read-only binding to the contract
	MULTICALLTransactor // Write-only binding to the contract
	MULTICALLFilterer   // Log filterer for contract events
}

// MULTICALLCaller is an auto generated read-only Go binding around an Ethereum contract.
type MULTICALLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MULTICALLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MULTICALLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MULTICALLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MULTICALLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MULTICALLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MULTICALLSession struct {
	Contract     *MULTICALL        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MULTICALLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MULTICALLCallerSession struct {
	Contract *MULTICALLCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MULTICALLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MULTICALLTransactorSession struct {
	Contract     *MULTICALLTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MULTICALLRaw is an auto generated low-level Go binding around an Ethereum contract.
type MULTICALLRaw struct {
	Contract *MULTICALL // Generic contract binding to access the raw methods on
}

// MULTICALLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MULTICALLCallerRaw struct {
	Contract *MULTICALLCaller // Generic read-only contract binding to access the raw methods on
}

// MULTICALLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MULTICALLTransactorRaw struct {
	Contract *MULTICALLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMULTICALL creates a new instance of MULTICALL, bound to a specific deployed contract.
func NewMULTICALL(address common.Address, backend bind.ContractBackend) (*MULTICALL, error) {
	contract, err := bindMULTICALL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MULTICALL{MULTICALLCaller: MULTICALLCaller{contract: contract}, MULTICALLTransactor: MULTICALLTransactor{contract: contract}, MULTICALLFilterer: MULTICALLFilterer{contract: contract}}, nil
}

// NewMULTICALLCaller creates a new read-only instance of MULTICALL, bound to a specific deployed contract.
func NewMULTICALLCaller(address common.Address, caller bind.ContractCaller) (*MULTICALLCaller, error) {
	contract, err := bindMULTICALL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MULTICALLCaller{contract: contract}, nil
}

// NewMULTICALLTransactor creates a new write-only instance of MULTICALL, bound to a specific deployed contract.
func NewMULTICALLTransactor(address common.Address, transactor bind.ContractTransactor) (*MULTICALLTransactor, error) {
	contract, err := bindMULTICALL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MULTICALLTransactor{contract: contract}, nil
}

// NewMULTICALLFilterer creates a new log filterer instance of MULTICALL, bound to a specific deployed contract.
func NewMULTICALLFilterer(address common.Address, filterer bind.ContractFilterer) (*MULTICALLFilterer, error) {
	contract, err := bindMULTICALL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MULTICALLFilterer{contract: contract}, nil
}

// bindMULTICALL binds a generic wrapper to an already deployed contract.
func bindMULTICALL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MULTICALLABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MULTICALL *MULTICALLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MULTICALL.Contract.MULTICALLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MULTICALL *MULTICALLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MULTICALL.Contract.MULTICALLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MULTICALL *MULTICALLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MULTICALL.Contract.MULTICALLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MULTICALL *MULTICALLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MULTICALL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MULTICALL *MULTICALLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MULTICALL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MULTICALL *MULTICALLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MULTICALL.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MULTICALL.Contract.GetBlockHash(&_MULTICALL.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MULTICALL.Contract.GetBlockHash(&_MULTICALL.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MULTICALL *MULTICALLCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MULTICALL *MULTICALLSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MULTICALL.Contract.GetCurrentBlockCoinbase(&_MULTICALL.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MULTICALL *MULTICALLCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MULTICALL.Contract.GetCurrentBlockCoinbase(&_MULTICALL.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MULTICALL *MULTICALLCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MULTICALL *MULTICALLSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockDifficulty(&_MULTICALL.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MULTICALL *MULTICALLCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockDifficulty(&_MULTICALL.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MULTICALL *MULTICALLCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MULTICALL *MULTICALLSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockGasLimit(&_MULTICALL.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MULTICALL *MULTICALLCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockGasLimit(&_MULTICALL.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MULTICALL *MULTICALLCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MULTICALL *MULTICALLSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockTimestamp(&_MULTICALL.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MULTICALL *MULTICALLCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MULTICALL.Contract.GetCurrentBlockTimestamp(&_MULTICALL.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MULTICALL *MULTICALLCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MULTICALL *MULTICALLSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MULTICALL.Contract.GetEthBalance(&_MULTICALL.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MULTICALL *MULTICALLCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MULTICALL.Contract.GetEthBalance(&_MULTICALL.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MULTICALL.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLSession) GetLastBlockHash() ([32]byte, error) {
	return _MULTICALL.Contract.GetLastBlockHash(&_MULTICALL.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MULTICALL *MULTICALLCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MULTICALL.Contract.GetLastBlockHash(&_MULTICALL.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MULTICALL *MULTICALLTransactor) Aggregate(opts *bind.TransactOpts, calls []MulticallCall) (*types.Transaction, error) {
	return _MULTICALL.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MULTICALL *MULTICALLSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _MULTICALL.Contract.Aggregate(&_MULTICALL.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MULTICALL *MULTICALLTransactorSession) Aggregate(calls []MulticallCall) (*types.Transaction, error) {
	return _MULTICALL.Contract.Aggregate(&_MULTICALL.TransactOpts, calls)
}
