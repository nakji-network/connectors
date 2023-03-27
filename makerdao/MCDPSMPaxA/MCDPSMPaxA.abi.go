// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDPSMPaxA

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

// MCDPSMPAXAABI is the input ABI used to generate the binding from.
const MCDPSMPAXAABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BuyGem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SellGem\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"buyGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractDaiAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gemJoin\",\"outputs\":[{\"internalType\":\"contractAuthGemJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"sellGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDPSMPAXA is an auto generated Go binding around an Ethereum contract.
type MCDPSMPAXA struct {
	MCDPSMPAXACaller     // Read-only binding to the contract
	MCDPSMPAXATransactor // Write-only binding to the contract
	MCDPSMPAXAFilterer   // Log filterer for contract events
}

// MCDPSMPAXACaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPSMPAXACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMPAXATransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPSMPAXATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMPAXAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPSMPAXAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMPAXASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPSMPAXASession struct {
	Contract     *MCDPSMPAXA       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPSMPAXACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPSMPAXACallerSession struct {
	Contract *MCDPSMPAXACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MCDPSMPAXATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPSMPAXATransactorSession struct {
	Contract     *MCDPSMPAXATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MCDPSMPAXARaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPSMPAXARaw struct {
	Contract *MCDPSMPAXA // Generic contract binding to access the raw methods on
}

// MCDPSMPAXACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPSMPAXACallerRaw struct {
	Contract *MCDPSMPAXACaller // Generic read-only contract binding to access the raw methods on
}

// MCDPSMPAXATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPSMPAXATransactorRaw struct {
	Contract *MCDPSMPAXATransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPSMPAXA creates a new instance of MCDPSMPAXA, bound to a specific deployed contract.
func NewMCDPSMPAXA(address common.Address, backend bind.ContractBackend) (*MCDPSMPAXA, error) {
	contract, err := bindMCDPSMPAXA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXA{MCDPSMPAXACaller: MCDPSMPAXACaller{contract: contract}, MCDPSMPAXATransactor: MCDPSMPAXATransactor{contract: contract}, MCDPSMPAXAFilterer: MCDPSMPAXAFilterer{contract: contract}}, nil
}

// NewMCDPSMPAXACaller creates a new read-only instance of MCDPSMPAXA, bound to a specific deployed contract.
func NewMCDPSMPAXACaller(address common.Address, caller bind.ContractCaller) (*MCDPSMPAXACaller, error) {
	contract, err := bindMCDPSMPAXA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXACaller{contract: contract}, nil
}

// NewMCDPSMPAXATransactor creates a new write-only instance of MCDPSMPAXA, bound to a specific deployed contract.
func NewMCDPSMPAXATransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPSMPAXATransactor, error) {
	contract, err := bindMCDPSMPAXA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXATransactor{contract: contract}, nil
}

// NewMCDPSMPAXAFilterer creates a new log filterer instance of MCDPSMPAXA, bound to a specific deployed contract.
func NewMCDPSMPAXAFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPSMPAXAFilterer, error) {
	contract, err := bindMCDPSMPAXA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXAFilterer{contract: contract}, nil
}

// bindMCDPSMPAXA binds a generic wrapper to an already deployed contract.
func bindMCDPSMPAXA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPSMPAXAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMPAXA *MCDPSMPAXARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMPAXA.Contract.MCDPSMPAXACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMPAXA *MCDPSMPAXARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.MCDPSMPAXATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMPAXA *MCDPSMPAXARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.MCDPSMPAXATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMPAXA *MCDPSMPAXACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMPAXA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMPAXA *MCDPSMPAXATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMPAXA *MCDPSMPAXATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXASession) Dai() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Dai(&_MCDPSMPAXA.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Dai() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Dai(&_MCDPSMPAXA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXASession) DaiJoin() (common.Address, error) {
	return _MCDPSMPAXA.Contract.DaiJoin(&_MCDPSMPAXA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) DaiJoin() (common.Address, error) {
	return _MCDPSMPAXA.Contract.DaiJoin(&_MCDPSMPAXA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACaller) GemJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "gemJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXASession) GemJoin() (common.Address, error) {
	return _MCDPSMPAXA.Contract.GemJoin(&_MCDPSMPAXA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) GemJoin() (common.Address, error) {
	return _MCDPSMPAXA.Contract.GemJoin(&_MCDPSMPAXA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMPAXA *MCDPSMPAXASession) Ilk() ([32]byte, error) {
	return _MCDPSMPAXA.Contract.Ilk(&_MCDPSMPAXA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Ilk() ([32]byte, error) {
	return _MCDPSMPAXA.Contract.Ilk(&_MCDPSMPAXA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Tin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "tin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXASession) Tin() (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Tin(&_MCDPSMPAXA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Tin() (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Tin(&_MCDPSMPAXA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Tout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "tout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXASession) Tout() (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Tout(&_MCDPSMPAXA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Tout() (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Tout(&_MCDPSMPAXA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXASession) Vat() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Vat(&_MCDPSMPAXA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Vat() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Vat(&_MCDPSMPAXA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXASession) Vow() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Vow(&_MCDPSMPAXA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Vow() (common.Address, error) {
	return _MCDPSMPAXA.Contract.Vow(&_MCDPSMPAXA.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMPAXA.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXASession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Wards(&_MCDPSMPAXA.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMPAXA *MCDPSMPAXACallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMPAXA.Contract.Wards(&_MCDPSMPAXA.CallOpts, arg0)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) BuyGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "buyGem", usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.BuyGem(&_MCDPSMPAXA.TransactOpts, usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.BuyGem(&_MCDPSMPAXA.TransactOpts, usr, gemAmt)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Deny(&_MCDPSMPAXA.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Deny(&_MCDPSMPAXA.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.File(&_MCDPSMPAXA.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.File(&_MCDPSMPAXA.TransactOpts, what, data)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Hope(&_MCDPSMPAXA.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Hope(&_MCDPSMPAXA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Nope(&_MCDPSMPAXA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Nope(&_MCDPSMPAXA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Rely(&_MCDPSMPAXA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.Rely(&_MCDPSMPAXA.TransactOpts, usr)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactor) SellGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.contract.Transact(opts, "sellGem", usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXASession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.SellGem(&_MCDPSMPAXA.TransactOpts, usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMPAXA *MCDPSMPAXATransactorSession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMPAXA.Contract.SellGem(&_MCDPSMPAXA.TransactOpts, usr, gemAmt)
}

// MCDPSMPAXABuyGemIterator is returned from FilterBuyGem and is used to iterate over the raw logs and unpacked data for BuyGem events raised by the MCDPSMPAXA contract.
type MCDPSMPAXABuyGemIterator struct {
	Event *MCDPSMPAXABuyGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMPAXABuyGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMPAXABuyGem)
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
		it.Event = new(MCDPSMPAXABuyGem)
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
func (it *MCDPSMPAXABuyGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMPAXABuyGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMPAXABuyGem represents a BuyGem event raised by the MCDPSMPAXA contract.
type MCDPSMPAXABuyGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBuyGem is a free log retrieval operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) FilterBuyGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMPAXABuyGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.FilterLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXABuyGemIterator{contract: _MCDPSMPAXA.contract, event: "BuyGem", logs: logs, sub: sub}, nil
}

// WatchBuyGem is a free log subscription operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) WatchBuyGem(opts *bind.WatchOpts, sink chan<- *MCDPSMPAXABuyGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.WatchLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMPAXABuyGem)
				if err := _MCDPSMPAXA.contract.UnpackLog(event, "BuyGem", log); err != nil {
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

// ParseBuyGem is a log parse operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) ParseBuyGem(log types.Log) (*MCDPSMPAXABuyGem, error) {
	event := new(MCDPSMPAXABuyGem)
	if err := _MCDPSMPAXA.contract.UnpackLog(event, "BuyGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMPAXADenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDPSMPAXA contract.
type MCDPSMPAXADenyIterator struct {
	Event *MCDPSMPAXADeny // Event containing the contract specifics and raw log

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
func (it *MCDPSMPAXADenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMPAXADeny)
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
		it.Event = new(MCDPSMPAXADeny)
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
func (it *MCDPSMPAXADenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMPAXADenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMPAXADeny represents a Deny event raised by the MCDPSMPAXA contract.
type MCDPSMPAXADeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDPSMPAXADenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXADenyIterator{contract: _MCDPSMPAXA.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDPSMPAXADeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMPAXADeny)
				if err := _MCDPSMPAXA.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) ParseDeny(log types.Log) (*MCDPSMPAXADeny, error) {
	event := new(MCDPSMPAXADeny)
	if err := _MCDPSMPAXA.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMPAXAFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDPSMPAXA contract.
type MCDPSMPAXAFileIterator struct {
	Event *MCDPSMPAXAFile // Event containing the contract specifics and raw log

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
func (it *MCDPSMPAXAFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMPAXAFile)
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
		it.Event = new(MCDPSMPAXAFile)
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
func (it *MCDPSMPAXAFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMPAXAFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMPAXAFile represents a File event raised by the MCDPSMPAXA contract.
type MCDPSMPAXAFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDPSMPAXAFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXAFileIterator{contract: _MCDPSMPAXA.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDPSMPAXAFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMPAXAFile)
				if err := _MCDPSMPAXA.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) ParseFile(log types.Log) (*MCDPSMPAXAFile, error) {
	event := new(MCDPSMPAXAFile)
	if err := _MCDPSMPAXA.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMPAXARelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDPSMPAXA contract.
type MCDPSMPAXARelyIterator struct {
	Event *MCDPSMPAXARely // Event containing the contract specifics and raw log

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
func (it *MCDPSMPAXARelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMPAXARely)
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
		it.Event = new(MCDPSMPAXARely)
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
func (it *MCDPSMPAXARelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMPAXARelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMPAXARely represents a Rely event raised by the MCDPSMPAXA contract.
type MCDPSMPAXARely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDPSMPAXARelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXARelyIterator{contract: _MCDPSMPAXA.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDPSMPAXARely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMPAXARely)
				if err := _MCDPSMPAXA.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) ParseRely(log types.Log) (*MCDPSMPAXARely, error) {
	event := new(MCDPSMPAXARely)
	if err := _MCDPSMPAXA.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMPAXASellGemIterator is returned from FilterSellGem and is used to iterate over the raw logs and unpacked data for SellGem events raised by the MCDPSMPAXA contract.
type MCDPSMPAXASellGemIterator struct {
	Event *MCDPSMPAXASellGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMPAXASellGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMPAXASellGem)
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
		it.Event = new(MCDPSMPAXASellGem)
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
func (it *MCDPSMPAXASellGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMPAXASellGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMPAXASellGem represents a SellGem event raised by the MCDPSMPAXA contract.
type MCDPSMPAXASellGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSellGem is a free log retrieval operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) FilterSellGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMPAXASellGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.FilterLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMPAXASellGemIterator{contract: _MCDPSMPAXA.contract, event: "SellGem", logs: logs, sub: sub}, nil
}

// WatchSellGem is a free log subscription operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) WatchSellGem(opts *bind.WatchOpts, sink chan<- *MCDPSMPAXASellGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMPAXA.contract.WatchLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMPAXASellGem)
				if err := _MCDPSMPAXA.contract.UnpackLog(event, "SellGem", log); err != nil {
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

// ParseSellGem is a log parse operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMPAXA *MCDPSMPAXAFilterer) ParseSellGem(log types.Log) (*MCDPSMPAXASellGem, error) {
	event := new(MCDPSMPAXASellGem)
	if err := _MCDPSMPAXA.contract.UnpackLog(event, "SellGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
