// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OSM_MOM

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

// OSMMOMABI is the input ABI used to generate the binding from.
const OSMMOMABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"osms\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"osm\",\"type\":\"address\"}],\"name\":\"setOsm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OSMMOM is an auto generated Go binding around an Ethereum contract.
type OSMMOM struct {
	OSMMOMCaller     // Read-only binding to the contract
	OSMMOMTransactor // Write-only binding to the contract
	OSMMOMFilterer   // Log filterer for contract events
}

// OSMMOMCaller is an auto generated read-only Go binding around an Ethereum contract.
type OSMMOMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OSMMOMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OSMMOMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OSMMOMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OSMMOMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OSMMOMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OSMMOMSession struct {
	Contract     *OSMMOM           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OSMMOMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OSMMOMCallerSession struct {
	Contract *OSMMOMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OSMMOMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OSMMOMTransactorSession struct {
	Contract     *OSMMOMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OSMMOMRaw is an auto generated low-level Go binding around an Ethereum contract.
type OSMMOMRaw struct {
	Contract *OSMMOM // Generic contract binding to access the raw methods on
}

// OSMMOMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OSMMOMCallerRaw struct {
	Contract *OSMMOMCaller // Generic read-only contract binding to access the raw methods on
}

// OSMMOMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OSMMOMTransactorRaw struct {
	Contract *OSMMOMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOSMMOM creates a new instance of OSMMOM, bound to a specific deployed contract.
func NewOSMMOM(address common.Address, backend bind.ContractBackend) (*OSMMOM, error) {
	contract, err := bindOSMMOM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OSMMOM{OSMMOMCaller: OSMMOMCaller{contract: contract}, OSMMOMTransactor: OSMMOMTransactor{contract: contract}, OSMMOMFilterer: OSMMOMFilterer{contract: contract}}, nil
}

// NewOSMMOMCaller creates a new read-only instance of OSMMOM, bound to a specific deployed contract.
func NewOSMMOMCaller(address common.Address, caller bind.ContractCaller) (*OSMMOMCaller, error) {
	contract, err := bindOSMMOM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OSMMOMCaller{contract: contract}, nil
}

// NewOSMMOMTransactor creates a new write-only instance of OSMMOM, bound to a specific deployed contract.
func NewOSMMOMTransactor(address common.Address, transactor bind.ContractTransactor) (*OSMMOMTransactor, error) {
	contract, err := bindOSMMOM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OSMMOMTransactor{contract: contract}, nil
}

// NewOSMMOMFilterer creates a new log filterer instance of OSMMOM, bound to a specific deployed contract.
func NewOSMMOMFilterer(address common.Address, filterer bind.ContractFilterer) (*OSMMOMFilterer, error) {
	contract, err := bindOSMMOM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OSMMOMFilterer{contract: contract}, nil
}

// bindOSMMOM binds a generic wrapper to an already deployed contract.
func bindOSMMOM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OSMMOMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OSMMOM *OSMMOMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OSMMOM.Contract.OSMMOMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OSMMOM *OSMMOMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OSMMOM.Contract.OSMMOMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OSMMOM *OSMMOMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OSMMOM.Contract.OSMMOMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OSMMOM *OSMMOMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OSMMOM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OSMMOM *OSMMOMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OSMMOM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OSMMOM *OSMMOMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OSMMOM.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OSMMOM *OSMMOMCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OSMMOM.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OSMMOM *OSMMOMSession) Authority() (common.Address, error) {
	return _OSMMOM.Contract.Authority(&_OSMMOM.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_OSMMOM *OSMMOMCallerSession) Authority() (common.Address, error) {
	return _OSMMOM.Contract.Authority(&_OSMMOM.CallOpts)
}

// Osms is a free data retrieval call binding the contract method 0x6c4ba760.
//
// Solidity: function osms(bytes32 ) view returns(address)
func (_OSMMOM *OSMMOMCaller) Osms(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _OSMMOM.contract.Call(opts, &out, "osms", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Osms is a free data retrieval call binding the contract method 0x6c4ba760.
//
// Solidity: function osms(bytes32 ) view returns(address)
func (_OSMMOM *OSMMOMSession) Osms(arg0 [32]byte) (common.Address, error) {
	return _OSMMOM.Contract.Osms(&_OSMMOM.CallOpts, arg0)
}

// Osms is a free data retrieval call binding the contract method 0x6c4ba760.
//
// Solidity: function osms(bytes32 ) view returns(address)
func (_OSMMOM *OSMMOMCallerSession) Osms(arg0 [32]byte) (common.Address, error) {
	return _OSMMOM.Contract.Osms(&_OSMMOM.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OSMMOM *OSMMOMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OSMMOM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OSMMOM *OSMMOMSession) Owner() (common.Address, error) {
	return _OSMMOM.Contract.Owner(&_OSMMOM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OSMMOM *OSMMOMCallerSession) Owner() (common.Address, error) {
	return _OSMMOM.Contract.Owner(&_OSMMOM.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_OSMMOM *OSMMOMTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_OSMMOM *OSMMOMSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetAuthority(&_OSMMOM.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_OSMMOM *OSMMOMTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetAuthority(&_OSMMOM.TransactOpts, authority_)
}

// SetOsm is a paid mutator transaction binding the contract method 0xc98cdf86.
//
// Solidity: function setOsm(bytes32 ilk, address osm) returns()
func (_OSMMOM *OSMMOMTransactor) SetOsm(opts *bind.TransactOpts, ilk [32]byte, osm common.Address) (*types.Transaction, error) {
	return _OSMMOM.contract.Transact(opts, "setOsm", ilk, osm)
}

// SetOsm is a paid mutator transaction binding the contract method 0xc98cdf86.
//
// Solidity: function setOsm(bytes32 ilk, address osm) returns()
func (_OSMMOM *OSMMOMSession) SetOsm(ilk [32]byte, osm common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetOsm(&_OSMMOM.TransactOpts, ilk, osm)
}

// SetOsm is a paid mutator transaction binding the contract method 0xc98cdf86.
//
// Solidity: function setOsm(bytes32 ilk, address osm) returns()
func (_OSMMOM *OSMMOMTransactorSession) SetOsm(ilk [32]byte, osm common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetOsm(&_OSMMOM.TransactOpts, ilk, osm)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_OSMMOM *OSMMOMTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_OSMMOM *OSMMOMSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetOwner(&_OSMMOM.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_OSMMOM *OSMMOMTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _OSMMOM.Contract.SetOwner(&_OSMMOM.TransactOpts, owner_)
}

// Stop is a paid mutator transaction binding the contract method 0x63c4f031.
//
// Solidity: function stop(bytes32 ilk) returns()
func (_OSMMOM *OSMMOMTransactor) Stop(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _OSMMOM.contract.Transact(opts, "stop", ilk)
}

// Stop is a paid mutator transaction binding the contract method 0x63c4f031.
//
// Solidity: function stop(bytes32 ilk) returns()
func (_OSMMOM *OSMMOMSession) Stop(ilk [32]byte) (*types.Transaction, error) {
	return _OSMMOM.Contract.Stop(&_OSMMOM.TransactOpts, ilk)
}

// Stop is a paid mutator transaction binding the contract method 0x63c4f031.
//
// Solidity: function stop(bytes32 ilk) returns()
func (_OSMMOM *OSMMOMTransactorSession) Stop(ilk [32]byte) (*types.Transaction, error) {
	return _OSMMOM.Contract.Stop(&_OSMMOM.TransactOpts, ilk)
}
