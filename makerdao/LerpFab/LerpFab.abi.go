// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LerpFab

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

// LERPFABABI is the input ABI used to generate the binding from.
const LERPFABABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lerp\",\"type\":\"address\"}],\"name\":\"LerpFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"NewIlkLerp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"NewLerp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"active\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lerps\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"}],\"name\":\"newIlkLerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lerp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"name_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"target_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"what_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration_\",\"type\":\"uint256\"}],\"name\":\"newLerp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lerp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LERPFAB is an auto generated Go binding around an Ethereum contract.
type LERPFAB struct {
	LERPFABCaller     // Read-only binding to the contract
	LERPFABTransactor // Write-only binding to the contract
	LERPFABFilterer   // Log filterer for contract events
}

// LERPFABCaller is an auto generated read-only Go binding around an Ethereum contract.
type LERPFABCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LERPFABTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LERPFABTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LERPFABFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LERPFABFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LERPFABSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LERPFABSession struct {
	Contract     *LERPFAB          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LERPFABCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LERPFABCallerSession struct {
	Contract *LERPFABCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LERPFABTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LERPFABTransactorSession struct {
	Contract     *LERPFABTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LERPFABRaw is an auto generated low-level Go binding around an Ethereum contract.
type LERPFABRaw struct {
	Contract *LERPFAB // Generic contract binding to access the raw methods on
}

// LERPFABCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LERPFABCallerRaw struct {
	Contract *LERPFABCaller // Generic read-only contract binding to access the raw methods on
}

// LERPFABTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LERPFABTransactorRaw struct {
	Contract *LERPFABTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLERPFAB creates a new instance of LERPFAB, bound to a specific deployed contract.
func NewLERPFAB(address common.Address, backend bind.ContractBackend) (*LERPFAB, error) {
	contract, err := bindLERPFAB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LERPFAB{LERPFABCaller: LERPFABCaller{contract: contract}, LERPFABTransactor: LERPFABTransactor{contract: contract}, LERPFABFilterer: LERPFABFilterer{contract: contract}}, nil
}

// NewLERPFABCaller creates a new read-only instance of LERPFAB, bound to a specific deployed contract.
func NewLERPFABCaller(address common.Address, caller bind.ContractCaller) (*LERPFABCaller, error) {
	contract, err := bindLERPFAB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LERPFABCaller{contract: contract}, nil
}

// NewLERPFABTransactor creates a new write-only instance of LERPFAB, bound to a specific deployed contract.
func NewLERPFABTransactor(address common.Address, transactor bind.ContractTransactor) (*LERPFABTransactor, error) {
	contract, err := bindLERPFAB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LERPFABTransactor{contract: contract}, nil
}

// NewLERPFABFilterer creates a new log filterer instance of LERPFAB, bound to a specific deployed contract.
func NewLERPFABFilterer(address common.Address, filterer bind.ContractFilterer) (*LERPFABFilterer, error) {
	contract, err := bindLERPFAB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LERPFABFilterer{contract: contract}, nil
}

// bindLERPFAB binds a generic wrapper to an already deployed contract.
func bindLERPFAB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LERPFABABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LERPFAB *LERPFABRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LERPFAB.Contract.LERPFABCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LERPFAB *LERPFABRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LERPFAB.Contract.LERPFABTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LERPFAB *LERPFABRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LERPFAB.Contract.LERPFABTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LERPFAB *LERPFABCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LERPFAB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LERPFAB *LERPFABTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LERPFAB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LERPFAB *LERPFABTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LERPFAB.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(address)
func (_LERPFAB *LERPFABCaller) Active(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LERPFAB.contract.Call(opts, &out, "active", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(address)
func (_LERPFAB *LERPFABSession) Active(arg0 *big.Int) (common.Address, error) {
	return _LERPFAB.Contract.Active(&_LERPFAB.CallOpts, arg0)
}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(address)
func (_LERPFAB *LERPFABCallerSession) Active(arg0 *big.Int) (common.Address, error) {
	return _LERPFAB.Contract.Active(&_LERPFAB.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_LERPFAB *LERPFABCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LERPFAB.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_LERPFAB *LERPFABSession) Count() (*big.Int, error) {
	return _LERPFAB.Contract.Count(&_LERPFAB.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_LERPFAB *LERPFABCallerSession) Count() (*big.Int, error) {
	return _LERPFAB.Contract.Count(&_LERPFAB.CallOpts)
}

// Lerps is a free data retrieval call binding the contract method 0x6cac18c8.
//
// Solidity: function lerps(bytes32 ) view returns(address)
func (_LERPFAB *LERPFABCaller) Lerps(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LERPFAB.contract.Call(opts, &out, "lerps", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lerps is a free data retrieval call binding the contract method 0x6cac18c8.
//
// Solidity: function lerps(bytes32 ) view returns(address)
func (_LERPFAB *LERPFABSession) Lerps(arg0 [32]byte) (common.Address, error) {
	return _LERPFAB.Contract.Lerps(&_LERPFAB.CallOpts, arg0)
}

// Lerps is a free data retrieval call binding the contract method 0x6cac18c8.
//
// Solidity: function lerps(bytes32 ) view returns(address)
func (_LERPFAB *LERPFABCallerSession) Lerps(arg0 [32]byte) (common.Address, error) {
	return _LERPFAB.Contract.Lerps(&_LERPFAB.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address[])
func (_LERPFAB *LERPFABCaller) List(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LERPFAB.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address[])
func (_LERPFAB *LERPFABSession) List() ([]common.Address, error) {
	return _LERPFAB.Contract.List(&_LERPFAB.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(address[])
func (_LERPFAB *LERPFABCallerSession) List() ([]common.Address, error) {
	return _LERPFAB.Contract.List(&_LERPFAB.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_LERPFAB *LERPFABCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LERPFAB.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_LERPFAB *LERPFABSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _LERPFAB.Contract.Wards(&_LERPFAB.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_LERPFAB *LERPFABCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _LERPFAB.Contract.Wards(&_LERPFAB.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_LERPFAB *LERPFABTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_LERPFAB *LERPFABSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.Contract.Deny(&_LERPFAB.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_LERPFAB *LERPFABTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.Contract.Deny(&_LERPFAB.TransactOpts, guy)
}

// NewIlkLerp is a paid mutator transaction binding the contract method 0xe02b0a5f.
//
// Solidity: function newIlkLerp(bytes32 name_, address target_, bytes32 ilk_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABTransactor) NewIlkLerp(opts *bind.TransactOpts, name_ [32]byte, target_ common.Address, ilk_ [32]byte, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.contract.Transact(opts, "newIlkLerp", name_, target_, ilk_, what_, startTime_, start_, end_, duration_)
}

// NewIlkLerp is a paid mutator transaction binding the contract method 0xe02b0a5f.
//
// Solidity: function newIlkLerp(bytes32 name_, address target_, bytes32 ilk_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABSession) NewIlkLerp(name_ [32]byte, target_ common.Address, ilk_ [32]byte, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.Contract.NewIlkLerp(&_LERPFAB.TransactOpts, name_, target_, ilk_, what_, startTime_, start_, end_, duration_)
}

// NewIlkLerp is a paid mutator transaction binding the contract method 0xe02b0a5f.
//
// Solidity: function newIlkLerp(bytes32 name_, address target_, bytes32 ilk_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABTransactorSession) NewIlkLerp(name_ [32]byte, target_ common.Address, ilk_ [32]byte, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.Contract.NewIlkLerp(&_LERPFAB.TransactOpts, name_, target_, ilk_, what_, startTime_, start_, end_, duration_)
}

// NewLerp is a paid mutator transaction binding the contract method 0x3f10119d.
//
// Solidity: function newLerp(bytes32 name_, address target_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABTransactor) NewLerp(opts *bind.TransactOpts, name_ [32]byte, target_ common.Address, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.contract.Transact(opts, "newLerp", name_, target_, what_, startTime_, start_, end_, duration_)
}

// NewLerp is a paid mutator transaction binding the contract method 0x3f10119d.
//
// Solidity: function newLerp(bytes32 name_, address target_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABSession) NewLerp(name_ [32]byte, target_ common.Address, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.Contract.NewLerp(&_LERPFAB.TransactOpts, name_, target_, what_, startTime_, start_, end_, duration_)
}

// NewLerp is a paid mutator transaction binding the contract method 0x3f10119d.
//
// Solidity: function newLerp(bytes32 name_, address target_, bytes32 what_, uint256 startTime_, uint256 start_, uint256 end_, uint256 duration_) returns(address lerp)
func (_LERPFAB *LERPFABTransactorSession) NewLerp(name_ [32]byte, target_ common.Address, what_ [32]byte, startTime_ *big.Int, start_ *big.Int, end_ *big.Int, duration_ *big.Int) (*types.Transaction, error) {
	return _LERPFAB.Contract.NewLerp(&_LERPFAB.TransactOpts, name_, target_, what_, startTime_, start_, end_, duration_)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_LERPFAB *LERPFABTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_LERPFAB *LERPFABSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.Contract.Rely(&_LERPFAB.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_LERPFAB *LERPFABTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _LERPFAB.Contract.Rely(&_LERPFAB.TransactOpts, guy)
}

// Tall is a paid mutator transaction binding the contract method 0x49c24b14.
//
// Solidity: function tall() returns()
func (_LERPFAB *LERPFABTransactor) Tall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LERPFAB.contract.Transact(opts, "tall")
}

// Tall is a paid mutator transaction binding the contract method 0x49c24b14.
//
// Solidity: function tall() returns()
func (_LERPFAB *LERPFABSession) Tall() (*types.Transaction, error) {
	return _LERPFAB.Contract.Tall(&_LERPFAB.TransactOpts)
}

// Tall is a paid mutator transaction binding the contract method 0x49c24b14.
//
// Solidity: function tall() returns()
func (_LERPFAB *LERPFABTransactorSession) Tall() (*types.Transaction, error) {
	return _LERPFAB.Contract.Tall(&_LERPFAB.TransactOpts)
}

// LERPFABDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the LERPFAB contract.
type LERPFABDenyIterator struct {
	Event *LERPFABDeny // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LERPFABDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LERPFABDeny)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LERPFABDeny)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LERPFABDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LERPFABDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LERPFABDeny represents a Deny event raised by the LERPFAB contract.
type LERPFABDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_LERPFAB *LERPFABFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*LERPFABDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _LERPFAB.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &LERPFABDenyIterator{contract: _LERPFAB.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_LERPFAB *LERPFABFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *LERPFABDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _LERPFAB.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LERPFABDeny)
				if err := _LERPFAB.contract.UnpackLog(event, "Deny", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_LERPFAB *LERPFABFilterer) ParseDeny(log types.Log) (*LERPFABDeny, error) {
	event := new(LERPFABDeny)
	if err := _LERPFAB.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LERPFABLerpFinishedIterator is returned from FilterLerpFinished and is used to iterate over the raw logs and unpacked data for LerpFinished events raised by the LERPFAB contract.
type LERPFABLerpFinishedIterator struct {
	Event *LERPFABLerpFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LERPFABLerpFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LERPFABLerpFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LERPFABLerpFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LERPFABLerpFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LERPFABLerpFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LERPFABLerpFinished represents a LerpFinished event raised by the LERPFAB contract.
type LERPFABLerpFinished struct {
	Lerp common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLerpFinished is a free log retrieval operation binding the contract event 0x9f8a68d4ad5fab660470a3a72f045c1832e0ea52d4ba0424ce31d9053f813692.
//
// Solidity: event LerpFinished(address indexed lerp)
func (_LERPFAB *LERPFABFilterer) FilterLerpFinished(opts *bind.FilterOpts, lerp []common.Address) (*LERPFABLerpFinishedIterator, error) {

	var lerpRule []interface{}
	for _, lerpItem := range lerp {
		lerpRule = append(lerpRule, lerpItem)
	}

	logs, sub, err := _LERPFAB.contract.FilterLogs(opts, "LerpFinished", lerpRule)
	if err != nil {
		return nil, err
	}
	return &LERPFABLerpFinishedIterator{contract: _LERPFAB.contract, event: "LerpFinished", logs: logs, sub: sub}, nil
}

// WatchLerpFinished is a free log subscription operation binding the contract event 0x9f8a68d4ad5fab660470a3a72f045c1832e0ea52d4ba0424ce31d9053f813692.
//
// Solidity: event LerpFinished(address indexed lerp)
func (_LERPFAB *LERPFABFilterer) WatchLerpFinished(opts *bind.WatchOpts, sink chan<- *LERPFABLerpFinished, lerp []common.Address) (event.Subscription, error) {

	var lerpRule []interface{}
	for _, lerpItem := range lerp {
		lerpRule = append(lerpRule, lerpItem)
	}

	logs, sub, err := _LERPFAB.contract.WatchLogs(opts, "LerpFinished", lerpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LERPFABLerpFinished)
				if err := _LERPFAB.contract.UnpackLog(event, "LerpFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLerpFinished is a log parse operation binding the contract event 0x9f8a68d4ad5fab660470a3a72f045c1832e0ea52d4ba0424ce31d9053f813692.
//
// Solidity: event LerpFinished(address indexed lerp)
func (_LERPFAB *LERPFABFilterer) ParseLerpFinished(log types.Log) (*LERPFABLerpFinished, error) {
	event := new(LERPFABLerpFinished)
	if err := _LERPFAB.contract.UnpackLog(event, "LerpFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LERPFABNewIlkLerpIterator is returned from FilterNewIlkLerp and is used to iterate over the raw logs and unpacked data for NewIlkLerp events raised by the LERPFAB contract.
type LERPFABNewIlkLerpIterator struct {
	Event *LERPFABNewIlkLerp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LERPFABNewIlkLerpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LERPFABNewIlkLerp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LERPFABNewIlkLerp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LERPFABNewIlkLerpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LERPFABNewIlkLerpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LERPFABNewIlkLerp represents a NewIlkLerp event raised by the LERPFAB contract.
type LERPFABNewIlkLerp struct {
	Name      [32]byte
	Target    common.Address
	Ilk       [32]byte
	What      [32]byte
	StartTime *big.Int
	Start     *big.Int
	End       *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewIlkLerp is a free log retrieval operation binding the contract event 0x496fde58572fe4aed9f930c3e26848b394a8180689b4f314b89f9a429c283bd9.
//
// Solidity: event NewIlkLerp(bytes32 name, address indexed target, bytes32 ilk, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) FilterNewIlkLerp(opts *bind.FilterOpts, target []common.Address) (*LERPFABNewIlkLerpIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LERPFAB.contract.FilterLogs(opts, "NewIlkLerp", targetRule)
	if err != nil {
		return nil, err
	}
	return &LERPFABNewIlkLerpIterator{contract: _LERPFAB.contract, event: "NewIlkLerp", logs: logs, sub: sub}, nil
}

// WatchNewIlkLerp is a free log subscription operation binding the contract event 0x496fde58572fe4aed9f930c3e26848b394a8180689b4f314b89f9a429c283bd9.
//
// Solidity: event NewIlkLerp(bytes32 name, address indexed target, bytes32 ilk, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) WatchNewIlkLerp(opts *bind.WatchOpts, sink chan<- *LERPFABNewIlkLerp, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LERPFAB.contract.WatchLogs(opts, "NewIlkLerp", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LERPFABNewIlkLerp)
				if err := _LERPFAB.contract.UnpackLog(event, "NewIlkLerp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewIlkLerp is a log parse operation binding the contract event 0x496fde58572fe4aed9f930c3e26848b394a8180689b4f314b89f9a429c283bd9.
//
// Solidity: event NewIlkLerp(bytes32 name, address indexed target, bytes32 ilk, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) ParseNewIlkLerp(log types.Log) (*LERPFABNewIlkLerp, error) {
	event := new(LERPFABNewIlkLerp)
	if err := _LERPFAB.contract.UnpackLog(event, "NewIlkLerp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LERPFABNewLerpIterator is returned from FilterNewLerp and is used to iterate over the raw logs and unpacked data for NewLerp events raised by the LERPFAB contract.
type LERPFABNewLerpIterator struct {
	Event *LERPFABNewLerp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LERPFABNewLerpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LERPFABNewLerp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LERPFABNewLerp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LERPFABNewLerpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LERPFABNewLerpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LERPFABNewLerp represents a NewLerp event raised by the LERPFAB contract.
type LERPFABNewLerp struct {
	Name      [32]byte
	Target    common.Address
	What      [32]byte
	StartTime *big.Int
	Start     *big.Int
	End       *big.Int
	Duration  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewLerp is a free log retrieval operation binding the contract event 0xb99aca1b7092bfed8143fec3d884486851030f969955ccb1a9e26d28211a28b8.
//
// Solidity: event NewLerp(bytes32 name, address indexed target, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) FilterNewLerp(opts *bind.FilterOpts, target []common.Address) (*LERPFABNewLerpIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LERPFAB.contract.FilterLogs(opts, "NewLerp", targetRule)
	if err != nil {
		return nil, err
	}
	return &LERPFABNewLerpIterator{contract: _LERPFAB.contract, event: "NewLerp", logs: logs, sub: sub}, nil
}

// WatchNewLerp is a free log subscription operation binding the contract event 0xb99aca1b7092bfed8143fec3d884486851030f969955ccb1a9e26d28211a28b8.
//
// Solidity: event NewLerp(bytes32 name, address indexed target, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) WatchNewLerp(opts *bind.WatchOpts, sink chan<- *LERPFABNewLerp, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _LERPFAB.contract.WatchLogs(opts, "NewLerp", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LERPFABNewLerp)
				if err := _LERPFAB.contract.UnpackLog(event, "NewLerp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLerp is a log parse operation binding the contract event 0xb99aca1b7092bfed8143fec3d884486851030f969955ccb1a9e26d28211a28b8.
//
// Solidity: event NewLerp(bytes32 name, address indexed target, bytes32 what, uint256 startTime, uint256 start, uint256 end, uint256 duration)
func (_LERPFAB *LERPFABFilterer) ParseNewLerp(log types.Log) (*LERPFABNewLerp, error) {
	event := new(LERPFABNewLerp)
	if err := _LERPFAB.contract.UnpackLog(event, "NewLerp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LERPFABRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the LERPFAB contract.
type LERPFABRelyIterator struct {
	Event *LERPFABRely // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LERPFABRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LERPFABRely)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LERPFABRely)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LERPFABRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LERPFABRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LERPFABRely represents a Rely event raised by the LERPFAB contract.
type LERPFABRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_LERPFAB *LERPFABFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*LERPFABRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _LERPFAB.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &LERPFABRelyIterator{contract: _LERPFAB.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_LERPFAB *LERPFABFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *LERPFABRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _LERPFAB.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LERPFABRely)
				if err := _LERPFAB.contract.UnpackLog(event, "Rely", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_LERPFAB *LERPFABFilterer) ParseRely(log types.Log) (*LERPFABRely, error) {
	event := new(LERPFABRely)
	if err := _LERPFAB.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
