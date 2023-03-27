// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CdpManager

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

// CDPMANAGERABI is the input ABI used to generate the binding from.
const CDPMANAGERABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"own\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"NewCdp\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"cdpAllow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cdpCan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cdpi\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"first\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"give\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"prev\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"quit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cdpSrc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cdpDst\",\"type\":\"uint256\"}],\"name\":\"shift\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"urnAllow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urnCan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CDPMANAGER is an auto generated Go binding around an Ethereum contract.
type CDPMANAGER struct {
	CDPMANAGERCaller     // Read-only binding to the contract
	CDPMANAGERTransactor // Write-only binding to the contract
	CDPMANAGERFilterer   // Log filterer for contract events
}

// CDPMANAGERCaller is an auto generated read-only Go binding around an Ethereum contract.
type CDPMANAGERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPMANAGERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CDPMANAGERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPMANAGERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CDPMANAGERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPMANAGERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CDPMANAGERSession struct {
	Contract     *CDPMANAGER       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CDPMANAGERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CDPMANAGERCallerSession struct {
	Contract *CDPMANAGERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CDPMANAGERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CDPMANAGERTransactorSession struct {
	Contract     *CDPMANAGERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CDPMANAGERRaw is an auto generated low-level Go binding around an Ethereum contract.
type CDPMANAGERRaw struct {
	Contract *CDPMANAGER // Generic contract binding to access the raw methods on
}

// CDPMANAGERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CDPMANAGERCallerRaw struct {
	Contract *CDPMANAGERCaller // Generic read-only contract binding to access the raw methods on
}

// CDPMANAGERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CDPMANAGERTransactorRaw struct {
	Contract *CDPMANAGERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCDPMANAGER creates a new instance of CDPMANAGER, bound to a specific deployed contract.
func NewCDPMANAGER(address common.Address, backend bind.ContractBackend) (*CDPMANAGER, error) {
	contract, err := bindCDPMANAGER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CDPMANAGER{CDPMANAGERCaller: CDPMANAGERCaller{contract: contract}, CDPMANAGERTransactor: CDPMANAGERTransactor{contract: contract}, CDPMANAGERFilterer: CDPMANAGERFilterer{contract: contract}}, nil
}

// NewCDPMANAGERCaller creates a new read-only instance of CDPMANAGER, bound to a specific deployed contract.
func NewCDPMANAGERCaller(address common.Address, caller bind.ContractCaller) (*CDPMANAGERCaller, error) {
	contract, err := bindCDPMANAGER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CDPMANAGERCaller{contract: contract}, nil
}

// NewCDPMANAGERTransactor creates a new write-only instance of CDPMANAGER, bound to a specific deployed contract.
func NewCDPMANAGERTransactor(address common.Address, transactor bind.ContractTransactor) (*CDPMANAGERTransactor, error) {
	contract, err := bindCDPMANAGER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CDPMANAGERTransactor{contract: contract}, nil
}

// NewCDPMANAGERFilterer creates a new log filterer instance of CDPMANAGER, bound to a specific deployed contract.
func NewCDPMANAGERFilterer(address common.Address, filterer bind.ContractFilterer) (*CDPMANAGERFilterer, error) {
	contract, err := bindCDPMANAGER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CDPMANAGERFilterer{contract: contract}, nil
}

// bindCDPMANAGER binds a generic wrapper to an already deployed contract.
func bindCDPMANAGER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CDPMANAGERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CDPMANAGER *CDPMANAGERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CDPMANAGER.Contract.CDPMANAGERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CDPMANAGER *CDPMANAGERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.CDPMANAGERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CDPMANAGER *CDPMANAGERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.CDPMANAGERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CDPMANAGER *CDPMANAGERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CDPMANAGER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CDPMANAGER *CDPMANAGERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CDPMANAGER *CDPMANAGERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.contract.Transact(opts, method, params...)
}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) CdpCan(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "cdpCan", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) CdpCan(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.CdpCan(&_CDPMANAGER.CallOpts, arg0, arg1, arg2)
}

// CdpCan is a free data retrieval call binding the contract method 0x5aebb460.
//
// Solidity: function cdpCan(address , uint256 , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) CdpCan(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.CdpCan(&_CDPMANAGER.CallOpts, arg0, arg1, arg2)
}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) Cdpi(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "cdpi")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) Cdpi() (*big.Int, error) {
	return _CDPMANAGER.Contract.Cdpi(&_CDPMANAGER.CallOpts)
}

// Cdpi is a free data retrieval call binding the contract method 0xb3d178f2.
//
// Solidity: function cdpi() view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) Cdpi() (*big.Int, error) {
	return _CDPMANAGER.Contract.Cdpi(&_CDPMANAGER.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) Count(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "count", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) Count(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.Count(&_CDPMANAGER.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x05d85eda.
//
// Solidity: function count(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) Count(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.Count(&_CDPMANAGER.CallOpts, arg0)
}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) First(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "first", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) First(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.First(&_CDPMANAGER.CallOpts, arg0)
}

// First is a free data retrieval call binding the contract method 0xfc73d771.
//
// Solidity: function first(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) First(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.First(&_CDPMANAGER.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPMANAGER *CDPMANAGERCaller) Ilks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "ilks", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPMANAGER *CDPMANAGERSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _CDPMANAGER.Contract.Ilks(&_CDPMANAGER.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPMANAGER *CDPMANAGERCallerSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _CDPMANAGER.Contract.Ilks(&_CDPMANAGER.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) Last(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "last", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) Last(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.Last(&_CDPMANAGER.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) Last(arg0 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.Last(&_CDPMANAGER.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_CDPMANAGER *CDPMANAGERCaller) List(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "list", arg0)

	outstruct := new(struct {
		Prev *big.Int
		Next *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Prev = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Next = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_CDPMANAGER *CDPMANAGERSession) List(arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	return _CDPMANAGER.Contract.List(&_CDPMANAGER.CallOpts, arg0)
}

// List is a free data retrieval call binding the contract method 0x80c9419e.
//
// Solidity: function list(uint256 ) view returns(uint256 prev, uint256 next)
func (_CDPMANAGER *CDPMANAGERCallerSession) List(arg0 *big.Int) (struct {
	Prev *big.Int
	Next *big.Int
}, error) {
	return _CDPMANAGER.Contract.List(&_CDPMANAGER.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERCaller) Owns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "owns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _CDPMANAGER.Contract.Owns(&_CDPMANAGER.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERCallerSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _CDPMANAGER.Contract.Owns(&_CDPMANAGER.CallOpts, arg0)
}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCaller) UrnCan(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "urnCan", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) UrnCan(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.UrnCan(&_CDPMANAGER.CallOpts, arg0, arg1)
}

// UrnCan is a free data retrieval call binding the contract method 0xb2b192e6.
//
// Solidity: function urnCan(address , address ) view returns(uint256)
func (_CDPMANAGER *CDPMANAGERCallerSession) UrnCan(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _CDPMANAGER.Contract.UrnCan(&_CDPMANAGER.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERCaller) Urns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "urns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERSession) Urns(arg0 *big.Int) (common.Address, error) {
	return _CDPMANAGER.Contract.Urns(&_CDPMANAGER.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2726b073.
//
// Solidity: function urns(uint256 ) view returns(address)
func (_CDPMANAGER *CDPMANAGERCallerSession) Urns(arg0 *big.Int) (common.Address, error) {
	return _CDPMANAGER.Contract.Urns(&_CDPMANAGER.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_CDPMANAGER *CDPMANAGERCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CDPMANAGER.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_CDPMANAGER *CDPMANAGERSession) Vat() (common.Address, error) {
	return _CDPMANAGER.Contract.Vat(&_CDPMANAGER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_CDPMANAGER *CDPMANAGERCallerSession) Vat() (common.Address, error) {
	return _CDPMANAGER.Contract.Vat(&_CDPMANAGER.CallOpts)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) CdpAllow(opts *bind.TransactOpts, cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "cdpAllow", cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERSession) CdpAllow(cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.CdpAllow(&_CDPMANAGER.TransactOpts, cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0x0b63fb62.
//
// Solidity: function cdpAllow(uint256 cdp, address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) CdpAllow(cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.CdpAllow(&_CDPMANAGER.TransactOpts, cdp, usr, ok)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Enter(opts *bind.TransactOpts, src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "enter", src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_CDPMANAGER *CDPMANAGERSession) Enter(src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Enter(&_CDPMANAGER.TransactOpts, src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0x7e348b7d.
//
// Solidity: function enter(address src, uint256 cdp) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Enter(src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Enter(&_CDPMANAGER.TransactOpts, src, cdp)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "flux", ilk, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERSession) Flux(ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Flux(&_CDPMANAGER.TransactOpts, ilk, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x18af4d60.
//
// Solidity: function flux(bytes32 ilk, uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Flux(ilk [32]byte, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Flux(&_CDPMANAGER.TransactOpts, ilk, cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Flux0(opts *bind.TransactOpts, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "flux0", cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERSession) Flux0(cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Flux0(&_CDPMANAGER.TransactOpts, cdp, dst, wad)
}

// Flux0 is a paid mutator transaction binding the contract method 0x9bb8f838.
//
// Solidity: function flux(uint256 cdp, address dst, uint256 wad) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Flux0(cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Flux0(&_CDPMANAGER.TransactOpts, cdp, dst, wad)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Frob(opts *bind.TransactOpts, cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "frob", cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_CDPMANAGER *CDPMANAGERSession) Frob(cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Frob(&_CDPMANAGER.TransactOpts, cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x45e6bdcd.
//
// Solidity: function frob(uint256 cdp, int256 dink, int256 dart) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Frob(cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Frob(&_CDPMANAGER.TransactOpts, cdp, dink, dart)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Give(opts *bind.TransactOpts, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "give", cdp, dst)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERSession) Give(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Give(&_CDPMANAGER.TransactOpts, cdp, dst)
}

// Give is a paid mutator transaction binding the contract method 0xfcafcc68.
//
// Solidity: function give(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Give(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Give(&_CDPMANAGER.TransactOpts, cdp, dst)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Move(opts *bind.TransactOpts, cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "move", cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_CDPMANAGER *CDPMANAGERSession) Move(cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Move(&_CDPMANAGER.TransactOpts, cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xf9f30db6.
//
// Solidity: function move(uint256 cdp, address dst, uint256 rad) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Move(cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Move(&_CDPMANAGER.TransactOpts, cdp, dst, rad)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPMANAGER *CDPMANAGERTransactor) Open(opts *bind.TransactOpts, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "open", ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPMANAGER *CDPMANAGERSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Open(&_CDPMANAGER.TransactOpts, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPMANAGER *CDPMANAGERTransactorSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Open(&_CDPMANAGER.TransactOpts, ilk, usr)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Quit(opts *bind.TransactOpts, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "quit", cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERSession) Quit(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Quit(&_CDPMANAGER.TransactOpts, cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x1b0dbf72.
//
// Solidity: function quit(uint256 cdp, address dst) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Quit(cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Quit(&_CDPMANAGER.TransactOpts, cdp, dst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) Shift(opts *bind.TransactOpts, cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "shift", cdpSrc, cdpDst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_CDPMANAGER *CDPMANAGERSession) Shift(cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Shift(&_CDPMANAGER.TransactOpts, cdpSrc, cdpDst)
}

// Shift is a paid mutator transaction binding the contract method 0xe50322a2.
//
// Solidity: function shift(uint256 cdpSrc, uint256 cdpDst) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) Shift(cdpSrc *big.Int, cdpDst *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.Shift(&_CDPMANAGER.TransactOpts, cdpSrc, cdpDst)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERTransactor) UrnAllow(opts *bind.TransactOpts, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.contract.Transact(opts, "urnAllow", usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERSession) UrnAllow(usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.UrnAllow(&_CDPMANAGER.TransactOpts, usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0xb68f4004.
//
// Solidity: function urnAllow(address usr, uint256 ok) returns()
func (_CDPMANAGER *CDPMANAGERTransactorSession) UrnAllow(usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _CDPMANAGER.Contract.UrnAllow(&_CDPMANAGER.TransactOpts, usr, ok)
}

// CDPMANAGERNewCdpIterator is returned from FilterNewCdp and is used to iterate over the raw logs and unpacked data for NewCdp events raised by the CDPMANAGER contract.
type CDPMANAGERNewCdpIterator struct {
	Event *CDPMANAGERNewCdp // Event containing the contract specifics and raw log

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
func (it *CDPMANAGERNewCdpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CDPMANAGERNewCdp)
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
		it.Event = new(CDPMANAGERNewCdp)
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
func (it *CDPMANAGERNewCdpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CDPMANAGERNewCdpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CDPMANAGERNewCdp represents a NewCdp event raised by the CDPMANAGER contract.
type CDPMANAGERNewCdp struct {
	Usr common.Address
	Own common.Address
	Cdp *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewCdp is a free log retrieval operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_CDPMANAGER *CDPMANAGERFilterer) FilterNewCdp(opts *bind.FilterOpts, usr []common.Address, own []common.Address, cdp []*big.Int) (*CDPMANAGERNewCdpIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var ownRule []interface{}
	for _, ownItem := range own {
		ownRule = append(ownRule, ownItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _CDPMANAGER.contract.FilterLogs(opts, "NewCdp", usrRule, ownRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return &CDPMANAGERNewCdpIterator{contract: _CDPMANAGER.contract, event: "NewCdp", logs: logs, sub: sub}, nil
}

// WatchNewCdp is a free log subscription operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_CDPMANAGER *CDPMANAGERFilterer) WatchNewCdp(opts *bind.WatchOpts, sink chan<- *CDPMANAGERNewCdp, usr []common.Address, own []common.Address, cdp []*big.Int) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var ownRule []interface{}
	for _, ownItem := range own {
		ownRule = append(ownRule, ownItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _CDPMANAGER.contract.WatchLogs(opts, "NewCdp", usrRule, ownRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CDPMANAGERNewCdp)
				if err := _CDPMANAGER.contract.UnpackLog(event, "NewCdp", log); err != nil {
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

// ParseNewCdp is a log parse operation binding the contract event 0xd6be0bc178658a382ff4f91c8c68b542aa6b71685b8fe427966b87745c3ea7a2.
//
// Solidity: event NewCdp(address indexed usr, address indexed own, uint256 indexed cdp)
func (_CDPMANAGER *CDPMANAGERFilterer) ParseNewCdp(log types.Log) (*CDPMANAGERNewCdp, error) {
	event := new(CDPMANAGERNewCdp)
	if err := _CDPMANAGER.contract.UnpackLog(event, "NewCdp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
