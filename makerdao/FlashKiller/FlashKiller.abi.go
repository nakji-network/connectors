// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FlashKiller

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

// FLASHKILLERABI is the input ABI used to generate the binding from.
const FLASHKILLERABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flash_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"flash\",\"outputs\":[{\"internalType\":\"contractFlashLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FLASHKILLER is an auto generated Go binding around an Ethereum contract.
type FLASHKILLER struct {
	FLASHKILLERCaller     // Read-only binding to the contract
	FLASHKILLERTransactor // Write-only binding to the contract
	FLASHKILLERFilterer   // Log filterer for contract events
}

// FLASHKILLERCaller is an auto generated read-only Go binding around an Ethereum contract.
type FLASHKILLERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLASHKILLERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FLASHKILLERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLASHKILLERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FLASHKILLERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FLASHKILLERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FLASHKILLERSession struct {
	Contract     *FLASHKILLER      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FLASHKILLERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FLASHKILLERCallerSession struct {
	Contract *FLASHKILLERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FLASHKILLERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FLASHKILLERTransactorSession struct {
	Contract     *FLASHKILLERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FLASHKILLERRaw is an auto generated low-level Go binding around an Ethereum contract.
type FLASHKILLERRaw struct {
	Contract *FLASHKILLER // Generic contract binding to access the raw methods on
}

// FLASHKILLERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FLASHKILLERCallerRaw struct {
	Contract *FLASHKILLERCaller // Generic read-only contract binding to access the raw methods on
}

// FLASHKILLERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FLASHKILLERTransactorRaw struct {
	Contract *FLASHKILLERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFLASHKILLER creates a new instance of FLASHKILLER, bound to a specific deployed contract.
func NewFLASHKILLER(address common.Address, backend bind.ContractBackend) (*FLASHKILLER, error) {
	contract, err := bindFLASHKILLER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FLASHKILLER{FLASHKILLERCaller: FLASHKILLERCaller{contract: contract}, FLASHKILLERTransactor: FLASHKILLERTransactor{contract: contract}, FLASHKILLERFilterer: FLASHKILLERFilterer{contract: contract}}, nil
}

// NewFLASHKILLERCaller creates a new read-only instance of FLASHKILLER, bound to a specific deployed contract.
func NewFLASHKILLERCaller(address common.Address, caller bind.ContractCaller) (*FLASHKILLERCaller, error) {
	contract, err := bindFLASHKILLER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FLASHKILLERCaller{contract: contract}, nil
}

// NewFLASHKILLERTransactor creates a new write-only instance of FLASHKILLER, bound to a specific deployed contract.
func NewFLASHKILLERTransactor(address common.Address, transactor bind.ContractTransactor) (*FLASHKILLERTransactor, error) {
	contract, err := bindFLASHKILLER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FLASHKILLERTransactor{contract: contract}, nil
}

// NewFLASHKILLERFilterer creates a new log filterer instance of FLASHKILLER, bound to a specific deployed contract.
func NewFLASHKILLERFilterer(address common.Address, filterer bind.ContractFilterer) (*FLASHKILLERFilterer, error) {
	contract, err := bindFLASHKILLER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FLASHKILLERFilterer{contract: contract}, nil
}

// bindFLASHKILLER binds a generic wrapper to an already deployed contract.
func bindFLASHKILLER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FLASHKILLERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLASHKILLER *FLASHKILLERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLASHKILLER.Contract.FLASHKILLERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLASHKILLER *FLASHKILLERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLASHKILLER.Contract.FLASHKILLERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLASHKILLER *FLASHKILLERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLASHKILLER.Contract.FLASHKILLERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FLASHKILLER *FLASHKILLERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FLASHKILLER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FLASHKILLER *FLASHKILLERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLASHKILLER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FLASHKILLER *FLASHKILLERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FLASHKILLER.Contract.contract.Transact(opts, method, params...)
}

// Flash is a free data retrieval call binding the contract method 0xd336c82d.
//
// Solidity: function flash() view returns(address)
func (_FLASHKILLER *FLASHKILLERCaller) Flash(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FLASHKILLER.contract.Call(opts, &out, "flash")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flash is a free data retrieval call binding the contract method 0xd336c82d.
//
// Solidity: function flash() view returns(address)
func (_FLASHKILLER *FLASHKILLERSession) Flash() (common.Address, error) {
	return _FLASHKILLER.Contract.Flash(&_FLASHKILLER.CallOpts)
}

// Flash is a free data retrieval call binding the contract method 0xd336c82d.
//
// Solidity: function flash() view returns(address)
func (_FLASHKILLER *FLASHKILLERCallerSession) Flash() (common.Address, error) {
	return _FLASHKILLER.Contract.Flash(&_FLASHKILLER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_FLASHKILLER *FLASHKILLERCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FLASHKILLER.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_FLASHKILLER *FLASHKILLERSession) Vat() (common.Address, error) {
	return _FLASHKILLER.Contract.Vat(&_FLASHKILLER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_FLASHKILLER *FLASHKILLERCallerSession) Vat() (common.Address, error) {
	return _FLASHKILLER.Contract.Vat(&_FLASHKILLER.CallOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FLASHKILLER *FLASHKILLERTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FLASHKILLER.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FLASHKILLER *FLASHKILLERSession) Kill() (*types.Transaction, error) {
	return _FLASHKILLER.Contract.Kill(&_FLASHKILLER.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FLASHKILLER *FLASHKILLERTransactorSession) Kill() (*types.Transaction, error) {
	return _FLASHKILLER.Contract.Kill(&_FLASHKILLER.TransactOpts)
}
