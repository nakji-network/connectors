// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDPSMGUSDA

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

// MCDPSMGUSDAABI is the input ABI used to generate the binding from.
const MCDPSMGUSDAABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BuyGem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SellGem\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"buyGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractDaiAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gemJoin\",\"outputs\":[{\"internalType\":\"contractAuthGemJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"sellGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDPSMGUSDA is an auto generated Go binding around an Ethereum contract.
type MCDPSMGUSDA struct {
	MCDPSMGUSDACaller     // Read-only binding to the contract
	MCDPSMGUSDATransactor // Write-only binding to the contract
	MCDPSMGUSDAFilterer   // Log filterer for contract events
}

// MCDPSMGUSDACaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPSMGUSDACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMGUSDATransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPSMGUSDATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMGUSDAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPSMGUSDAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMGUSDASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPSMGUSDASession struct {
	Contract     *MCDPSMGUSDA      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPSMGUSDACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPSMGUSDACallerSession struct {
	Contract *MCDPSMGUSDACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MCDPSMGUSDATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPSMGUSDATransactorSession struct {
	Contract     *MCDPSMGUSDATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MCDPSMGUSDARaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPSMGUSDARaw struct {
	Contract *MCDPSMGUSDA // Generic contract binding to access the raw methods on
}

// MCDPSMGUSDACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPSMGUSDACallerRaw struct {
	Contract *MCDPSMGUSDACaller // Generic read-only contract binding to access the raw methods on
}

// MCDPSMGUSDATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPSMGUSDATransactorRaw struct {
	Contract *MCDPSMGUSDATransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPSMGUSDA creates a new instance of MCDPSMGUSDA, bound to a specific deployed contract.
func NewMCDPSMGUSDA(address common.Address, backend bind.ContractBackend) (*MCDPSMGUSDA, error) {
	contract, err := bindMCDPSMGUSDA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDA{MCDPSMGUSDACaller: MCDPSMGUSDACaller{contract: contract}, MCDPSMGUSDATransactor: MCDPSMGUSDATransactor{contract: contract}, MCDPSMGUSDAFilterer: MCDPSMGUSDAFilterer{contract: contract}}, nil
}

// NewMCDPSMGUSDACaller creates a new read-only instance of MCDPSMGUSDA, bound to a specific deployed contract.
func NewMCDPSMGUSDACaller(address common.Address, caller bind.ContractCaller) (*MCDPSMGUSDACaller, error) {
	contract, err := bindMCDPSMGUSDA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDACaller{contract: contract}, nil
}

// NewMCDPSMGUSDATransactor creates a new write-only instance of MCDPSMGUSDA, bound to a specific deployed contract.
func NewMCDPSMGUSDATransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPSMGUSDATransactor, error) {
	contract, err := bindMCDPSMGUSDA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDATransactor{contract: contract}, nil
}

// NewMCDPSMGUSDAFilterer creates a new log filterer instance of MCDPSMGUSDA, bound to a specific deployed contract.
func NewMCDPSMGUSDAFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPSMGUSDAFilterer, error) {
	contract, err := bindMCDPSMGUSDA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDAFilterer{contract: contract}, nil
}

// bindMCDPSMGUSDA binds a generic wrapper to an already deployed contract.
func bindMCDPSMGUSDA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPSMGUSDAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMGUSDA *MCDPSMGUSDARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMGUSDA.Contract.MCDPSMGUSDACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMGUSDA *MCDPSMGUSDARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.MCDPSMGUSDATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMGUSDA *MCDPSMGUSDARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.MCDPSMGUSDATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMGUSDA *MCDPSMGUSDACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMGUSDA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Dai() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Dai(&_MCDPSMGUSDA.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Dai() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Dai(&_MCDPSMGUSDA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) DaiJoin() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.DaiJoin(&_MCDPSMGUSDA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) DaiJoin() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.DaiJoin(&_MCDPSMGUSDA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) GemJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "gemJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) GemJoin() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.GemJoin(&_MCDPSMGUSDA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) GemJoin() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.GemJoin(&_MCDPSMGUSDA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Ilk() ([32]byte, error) {
	return _MCDPSMGUSDA.Contract.Ilk(&_MCDPSMGUSDA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Ilk() ([32]byte, error) {
	return _MCDPSMGUSDA.Contract.Ilk(&_MCDPSMGUSDA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Tin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "tin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Tin() (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Tin(&_MCDPSMGUSDA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Tin() (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Tin(&_MCDPSMGUSDA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Tout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "tout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Tout() (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Tout(&_MCDPSMGUSDA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Tout() (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Tout(&_MCDPSMGUSDA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Vat() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Vat(&_MCDPSMGUSDA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Vat() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Vat(&_MCDPSMGUSDA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Vow() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Vow(&_MCDPSMGUSDA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Vow() (common.Address, error) {
	return _MCDPSMGUSDA.Contract.Vow(&_MCDPSMGUSDA.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMGUSDA.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Wards(&_MCDPSMGUSDA.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMGUSDA *MCDPSMGUSDACallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMGUSDA.Contract.Wards(&_MCDPSMGUSDA.CallOpts, arg0)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) BuyGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "buyGem", usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.BuyGem(&_MCDPSMGUSDA.TransactOpts, usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.BuyGem(&_MCDPSMGUSDA.TransactOpts, usr, gemAmt)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Deny(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Deny(&_MCDPSMGUSDA.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.File(&_MCDPSMGUSDA.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.File(&_MCDPSMGUSDA.TransactOpts, what, data)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Hope(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Hope(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Nope(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Nope(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Rely(&_MCDPSMGUSDA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.Rely(&_MCDPSMGUSDA.TransactOpts, usr)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactor) SellGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.contract.Transact(opts, "sellGem", usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDASession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.SellGem(&_MCDPSMGUSDA.TransactOpts, usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMGUSDA *MCDPSMGUSDATransactorSession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMGUSDA.Contract.SellGem(&_MCDPSMGUSDA.TransactOpts, usr, gemAmt)
}

// MCDPSMGUSDABuyGemIterator is returned from FilterBuyGem and is used to iterate over the raw logs and unpacked data for BuyGem events raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDABuyGemIterator struct {
	Event *MCDPSMGUSDABuyGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMGUSDABuyGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMGUSDABuyGem)
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
		it.Event = new(MCDPSMGUSDABuyGem)
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
func (it *MCDPSMGUSDABuyGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMGUSDABuyGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMGUSDABuyGem represents a BuyGem event raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDABuyGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBuyGem is a free log retrieval operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) FilterBuyGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMGUSDABuyGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.FilterLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDABuyGemIterator{contract: _MCDPSMGUSDA.contract, event: "BuyGem", logs: logs, sub: sub}, nil
}

// WatchBuyGem is a free log subscription operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) WatchBuyGem(opts *bind.WatchOpts, sink chan<- *MCDPSMGUSDABuyGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.WatchLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMGUSDABuyGem)
				if err := _MCDPSMGUSDA.contract.UnpackLog(event, "BuyGem", log); err != nil {
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
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) ParseBuyGem(log types.Log) (*MCDPSMGUSDABuyGem, error) {
	event := new(MCDPSMGUSDABuyGem)
	if err := _MCDPSMGUSDA.contract.UnpackLog(event, "BuyGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMGUSDADenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDADenyIterator struct {
	Event *MCDPSMGUSDADeny // Event containing the contract specifics and raw log

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
func (it *MCDPSMGUSDADenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMGUSDADeny)
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
		it.Event = new(MCDPSMGUSDADeny)
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
func (it *MCDPSMGUSDADenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMGUSDADenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMGUSDADeny represents a Deny event raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDADeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDPSMGUSDADenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDADenyIterator{contract: _MCDPSMGUSDA.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDPSMGUSDADeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMGUSDADeny)
				if err := _MCDPSMGUSDA.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) ParseDeny(log types.Log) (*MCDPSMGUSDADeny, error) {
	event := new(MCDPSMGUSDADeny)
	if err := _MCDPSMGUSDA.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMGUSDAFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDAFileIterator struct {
	Event *MCDPSMGUSDAFile // Event containing the contract specifics and raw log

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
func (it *MCDPSMGUSDAFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMGUSDAFile)
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
		it.Event = new(MCDPSMGUSDAFile)
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
func (it *MCDPSMGUSDAFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMGUSDAFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMGUSDAFile represents a File event raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDAFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDPSMGUSDAFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDAFileIterator{contract: _MCDPSMGUSDA.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDPSMGUSDAFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMGUSDAFile)
				if err := _MCDPSMGUSDA.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) ParseFile(log types.Log) (*MCDPSMGUSDAFile, error) {
	event := new(MCDPSMGUSDAFile)
	if err := _MCDPSMGUSDA.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMGUSDARelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDARelyIterator struct {
	Event *MCDPSMGUSDARely // Event containing the contract specifics and raw log

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
func (it *MCDPSMGUSDARelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMGUSDARely)
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
		it.Event = new(MCDPSMGUSDARely)
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
func (it *MCDPSMGUSDARelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMGUSDARelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMGUSDARely represents a Rely event raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDARely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDPSMGUSDARelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDARelyIterator{contract: _MCDPSMGUSDA.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDPSMGUSDARely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMGUSDARely)
				if err := _MCDPSMGUSDA.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) ParseRely(log types.Log) (*MCDPSMGUSDARely, error) {
	event := new(MCDPSMGUSDARely)
	if err := _MCDPSMGUSDA.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMGUSDASellGemIterator is returned from FilterSellGem and is used to iterate over the raw logs and unpacked data for SellGem events raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDASellGemIterator struct {
	Event *MCDPSMGUSDASellGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMGUSDASellGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMGUSDASellGem)
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
		it.Event = new(MCDPSMGUSDASellGem)
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
func (it *MCDPSMGUSDASellGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMGUSDASellGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMGUSDASellGem represents a SellGem event raised by the MCDPSMGUSDA contract.
type MCDPSMGUSDASellGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSellGem is a free log retrieval operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) FilterSellGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMGUSDASellGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.FilterLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMGUSDASellGemIterator{contract: _MCDPSMGUSDA.contract, event: "SellGem", logs: logs, sub: sub}, nil
}

// WatchSellGem is a free log subscription operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) WatchSellGem(opts *bind.WatchOpts, sink chan<- *MCDPSMGUSDASellGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMGUSDA.contract.WatchLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMGUSDASellGem)
				if err := _MCDPSMGUSDA.contract.UnpackLog(event, "SellGem", log); err != nil {
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
func (_MCDPSMGUSDA *MCDPSMGUSDAFilterer) ParseSellGem(log types.Log) (*MCDPSMGUSDASellGem, error) {
	event := new(MCDPSMGUSDASellGem)
	if err := _MCDPSMGUSDA.contract.UnpackLog(event, "SellGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
