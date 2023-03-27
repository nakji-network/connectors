// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDSpot

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

// MCDSPOTABI is the input ABI used to generate the binding from.
const MCDSPOTABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"val\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"}],\"name\":\"Poke\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"pip_\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"contractPipLike\",\"name\":\"pip\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mat\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"par\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDSPOT is an auto generated Go binding around an Ethereum contract.
type MCDSPOT struct {
	MCDSPOTCaller     // Read-only binding to the contract
	MCDSPOTTransactor // Write-only binding to the contract
	MCDSPOTFilterer   // Log filterer for contract events
}

// MCDSPOTCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDSPOTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDSPOTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDSPOTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDSPOTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDSPOTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDSPOTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDSPOTSession struct {
	Contract     *MCDSPOT          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDSPOTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDSPOTCallerSession struct {
	Contract *MCDSPOTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MCDSPOTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDSPOTTransactorSession struct {
	Contract     *MCDSPOTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MCDSPOTRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDSPOTRaw struct {
	Contract *MCDSPOT // Generic contract binding to access the raw methods on
}

// MCDSPOTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDSPOTCallerRaw struct {
	Contract *MCDSPOTCaller // Generic read-only contract binding to access the raw methods on
}

// MCDSPOTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDSPOTTransactorRaw struct {
	Contract *MCDSPOTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDSPOT creates a new instance of MCDSPOT, bound to a specific deployed contract.
func NewMCDSPOT(address common.Address, backend bind.ContractBackend) (*MCDSPOT, error) {
	contract, err := bindMCDSPOT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDSPOT{MCDSPOTCaller: MCDSPOTCaller{contract: contract}, MCDSPOTTransactor: MCDSPOTTransactor{contract: contract}, MCDSPOTFilterer: MCDSPOTFilterer{contract: contract}}, nil
}

// NewMCDSPOTCaller creates a new read-only instance of MCDSPOT, bound to a specific deployed contract.
func NewMCDSPOTCaller(address common.Address, caller bind.ContractCaller) (*MCDSPOTCaller, error) {
	contract, err := bindMCDSPOT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDSPOTCaller{contract: contract}, nil
}

// NewMCDSPOTTransactor creates a new write-only instance of MCDSPOT, bound to a specific deployed contract.
func NewMCDSPOTTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDSPOTTransactor, error) {
	contract, err := bindMCDSPOT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDSPOTTransactor{contract: contract}, nil
}

// NewMCDSPOTFilterer creates a new log filterer instance of MCDSPOT, bound to a specific deployed contract.
func NewMCDSPOTFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDSPOTFilterer, error) {
	contract, err := bindMCDSPOT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDSPOTFilterer{contract: contract}, nil
}

// bindMCDSPOT binds a generic wrapper to an already deployed contract.
func bindMCDSPOT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDSPOTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDSPOT *MCDSPOTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDSPOT.Contract.MCDSPOTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDSPOT *MCDSPOTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDSPOT.Contract.MCDSPOTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDSPOT *MCDSPOTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDSPOT.Contract.MCDSPOTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDSPOT *MCDSPOTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDSPOT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDSPOT *MCDSPOTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDSPOT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDSPOT *MCDSPOTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDSPOT.Contract.contract.Transact(opts, method, params...)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_MCDSPOT *MCDSPOTCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	var out []interface{}
	err := _MCDSPOT.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Pip common.Address
		Mat *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Mat = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_MCDSPOT *MCDSPOTSession) Ilks(arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	return _MCDSPOT.Contract.Ilks(&_MCDSPOT.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address pip, uint256 mat)
func (_MCDSPOT *MCDSPOTCallerSession) Ilks(arg0 [32]byte) (struct {
	Pip common.Address
	Mat *big.Int
}, error) {
	return _MCDSPOT.Contract.Ilks(&_MCDSPOT.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDSPOT *MCDSPOTCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDSPOT.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDSPOT *MCDSPOTSession) Live() (*big.Int, error) {
	return _MCDSPOT.Contract.Live(&_MCDSPOT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDSPOT *MCDSPOTCallerSession) Live() (*big.Int, error) {
	return _MCDSPOT.Contract.Live(&_MCDSPOT.CallOpts)
}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_MCDSPOT *MCDSPOTCaller) Par(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDSPOT.contract.Call(opts, &out, "par")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_MCDSPOT *MCDSPOTSession) Par() (*big.Int, error) {
	return _MCDSPOT.Contract.Par(&_MCDSPOT.CallOpts)
}

// Par is a free data retrieval call binding the contract method 0x495d32cb.
//
// Solidity: function par() view returns(uint256)
func (_MCDSPOT *MCDSPOTCallerSession) Par() (*big.Int, error) {
	return _MCDSPOT.Contract.Par(&_MCDSPOT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDSPOT *MCDSPOTCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDSPOT.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDSPOT *MCDSPOTSession) Vat() (common.Address, error) {
	return _MCDSPOT.Contract.Vat(&_MCDSPOT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDSPOT *MCDSPOTCallerSession) Vat() (common.Address, error) {
	return _MCDSPOT.Contract.Vat(&_MCDSPOT.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDSPOT *MCDSPOTCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDSPOT.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDSPOT *MCDSPOTSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDSPOT.Contract.Wards(&_MCDSPOT.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDSPOT *MCDSPOTCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDSPOT.Contract.Wards(&_MCDSPOT.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDSPOT *MCDSPOTTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDSPOT *MCDSPOTSession) Cage() (*types.Transaction, error) {
	return _MCDSPOT.Contract.Cage(&_MCDSPOT.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDSPOT *MCDSPOTTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDSPOT.Contract.Cage(&_MCDSPOT.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDSPOT *MCDSPOTTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDSPOT *MCDSPOTSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Deny(&_MCDSPOT.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Deny(&_MCDSPOT.TransactOpts, guy)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File(&_MCDSPOT.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File(&_MCDSPOT.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File0(&_MCDSPOT.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File0(&_MCDSPOT.TransactOpts, what, data)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_MCDSPOT *MCDSPOTTransactor) File1(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "file1", ilk, what, pip_)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_MCDSPOT *MCDSPOTSession) File1(ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File1(&_MCDSPOT.TransactOpts, ilk, what, pip_)
}

// File1 is a paid mutator transaction binding the contract method 0xebecb39d.
//
// Solidity: function file(bytes32 ilk, bytes32 what, address pip_) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) File1(ilk [32]byte, what [32]byte, pip_ common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.File1(&_MCDSPOT.TransactOpts, ilk, what, pip_)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_MCDSPOT *MCDSPOTTransactor) Poke(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "poke", ilk)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_MCDSPOT *MCDSPOTSession) Poke(ilk [32]byte) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Poke(&_MCDSPOT.TransactOpts, ilk)
}

// Poke is a paid mutator transaction binding the contract method 0x1504460f.
//
// Solidity: function poke(bytes32 ilk) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) Poke(ilk [32]byte) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Poke(&_MCDSPOT.TransactOpts, ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDSPOT *MCDSPOTTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDSPOT *MCDSPOTSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Rely(&_MCDSPOT.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_MCDSPOT *MCDSPOTTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _MCDSPOT.Contract.Rely(&_MCDSPOT.TransactOpts, guy)
}

// MCDSPOTPokeIterator is returned from FilterPoke and is used to iterate over the raw logs and unpacked data for Poke events raised by the MCDSPOT contract.
type MCDSPOTPokeIterator struct {
	Event *MCDSPOTPoke // Event containing the contract specifics and raw log

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
func (it *MCDSPOTPokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDSPOTPoke)
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
		it.Event = new(MCDSPOTPoke)
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
func (it *MCDSPOTPokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDSPOTPokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDSPOTPoke represents a Poke event raised by the MCDSPOT contract.
type MCDSPOTPoke struct {
	Ilk  [32]byte
	Val  [32]byte
	Spot *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoke is a free log retrieval operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_MCDSPOT *MCDSPOTFilterer) FilterPoke(opts *bind.FilterOpts) (*MCDSPOTPokeIterator, error) {

	logs, sub, err := _MCDSPOT.contract.FilterLogs(opts, "Poke")
	if err != nil {
		return nil, err
	}
	return &MCDSPOTPokeIterator{contract: _MCDSPOT.contract, event: "Poke", logs: logs, sub: sub}, nil
}

// WatchPoke is a free log subscription operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_MCDSPOT *MCDSPOTFilterer) WatchPoke(opts *bind.WatchOpts, sink chan<- *MCDSPOTPoke) (event.Subscription, error) {

	logs, sub, err := _MCDSPOT.contract.WatchLogs(opts, "Poke")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDSPOTPoke)
				if err := _MCDSPOT.contract.UnpackLog(event, "Poke", log); err != nil {
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

// ParsePoke is a log parse operation binding the contract event 0xdfd7467e425a8107cfd368d159957692c25085aacbcf5228ce08f10f2146486e.
//
// Solidity: event Poke(bytes32 ilk, bytes32 val, uint256 spot)
func (_MCDSPOT *MCDSPOTFilterer) ParsePoke(log types.Log) (*MCDSPOTPoke, error) {
	event := new(MCDSPOTPoke)
	if err := _MCDSPOT.contract.UnpackLog(event, "Poke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
