// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDFlap

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

// MCDFLAPABI is the input ABI used to generate the binding from.
const MCDFLAPABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"beg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tic\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"end\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fill\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike_1\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"tend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike_3\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDFLAP is an auto generated Go binding around an Ethereum contract.
type MCDFLAP struct {
	MCDFLAPCaller     // Read-only binding to the contract
	MCDFLAPTransactor // Write-only binding to the contract
	MCDFLAPFilterer   // Log filterer for contract events
}

// MCDFLAPCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDFLAPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLAPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDFLAPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLAPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDFLAPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLAPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDFLAPSession struct {
	Contract     *MCDFLAP          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDFLAPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDFLAPCallerSession struct {
	Contract *MCDFLAPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MCDFLAPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDFLAPTransactorSession struct {
	Contract     *MCDFLAPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MCDFLAPRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDFLAPRaw struct {
	Contract *MCDFLAP // Generic contract binding to access the raw methods on
}

// MCDFLAPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDFLAPCallerRaw struct {
	Contract *MCDFLAPCaller // Generic read-only contract binding to access the raw methods on
}

// MCDFLAPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDFLAPTransactorRaw struct {
	Contract *MCDFLAPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDFLAP creates a new instance of MCDFLAP, bound to a specific deployed contract.
func NewMCDFLAP(address common.Address, backend bind.ContractBackend) (*MCDFLAP, error) {
	contract, err := bindMCDFLAP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDFLAP{MCDFLAPCaller: MCDFLAPCaller{contract: contract}, MCDFLAPTransactor: MCDFLAPTransactor{contract: contract}, MCDFLAPFilterer: MCDFLAPFilterer{contract: contract}}, nil
}

// NewMCDFLAPCaller creates a new read-only instance of MCDFLAP, bound to a specific deployed contract.
func NewMCDFLAPCaller(address common.Address, caller bind.ContractCaller) (*MCDFLAPCaller, error) {
	contract, err := bindMCDFLAP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLAPCaller{contract: contract}, nil
}

// NewMCDFLAPTransactor creates a new write-only instance of MCDFLAP, bound to a specific deployed contract.
func NewMCDFLAPTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDFLAPTransactor, error) {
	contract, err := bindMCDFLAP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLAPTransactor{contract: contract}, nil
}

// NewMCDFLAPFilterer creates a new log filterer instance of MCDFLAP, bound to a specific deployed contract.
func NewMCDFLAPFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDFLAPFilterer, error) {
	contract, err := bindMCDFLAP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDFLAPFilterer{contract: contract}, nil
}

// bindMCDFLAP binds a generic wrapper to an already deployed contract.
func bindMCDFLAP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDFLAPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLAP *MCDFLAPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLAP.Contract.MCDFLAPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLAP *MCDFLAPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLAP.Contract.MCDFLAPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLAP *MCDFLAPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLAP.Contract.MCDFLAPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLAP *MCDFLAPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLAP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLAP *MCDFLAPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLAP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLAP *MCDFLAPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLAP.Contract.contract.Transact(opts, method, params...)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Beg(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "beg")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Beg() (*big.Int, error) {
	return _MCDFLAP.Contract.Beg(&_MCDFLAP.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Beg() (*big.Int, error) {
	return _MCDFLAP.Contract.Beg(&_MCDFLAP.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_MCDFLAP *MCDFLAPCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "bids", arg0)

	outstruct := new(struct {
		Bid *big.Int
		Lot *big.Int
		Guy common.Address
		Tic *big.Int
		End *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bid = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Guy = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Tic = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_MCDFLAP *MCDFLAPSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _MCDFLAP.Contract.Bids(&_MCDFLAP.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_MCDFLAP *MCDFLAPCallerSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _MCDFLAP.Contract.Bids(&_MCDFLAP.CallOpts, arg0)
}

// Fill is a free data retrieval call binding the contract method 0xd9c55ce1.
//
// Solidity: function fill() view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Fill(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "fill")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fill is a free data retrieval call binding the contract method 0xd9c55ce1.
//
// Solidity: function fill() view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Fill() (*big.Int, error) {
	return _MCDFLAP.Contract.Fill(&_MCDFLAP.CallOpts)
}

// Fill is a free data retrieval call binding the contract method 0xd9c55ce1.
//
// Solidity: function fill() view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Fill() (*big.Int, error) {
	return _MCDFLAP.Contract.Fill(&_MCDFLAP.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLAP *MCDFLAPCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLAP *MCDFLAPSession) Gem() (common.Address, error) {
	return _MCDFLAP.Contract.Gem(&_MCDFLAP.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLAP *MCDFLAPCallerSession) Gem() (common.Address, error) {
	return _MCDFLAP.Contract.Gem(&_MCDFLAP.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "kicks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Kicks() (*big.Int, error) {
	return _MCDFLAP.Contract.Kicks(&_MCDFLAP.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Kicks() (*big.Int, error) {
	return _MCDFLAP.Contract.Kicks(&_MCDFLAP.CallOpts)
}

// Lid is a free data retrieval call binding the contract method 0x26d2addc.
//
// Solidity: function lid() view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Lid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "lid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lid is a free data retrieval call binding the contract method 0x26d2addc.
//
// Solidity: function lid() view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Lid() (*big.Int, error) {
	return _MCDFLAP.Contract.Lid(&_MCDFLAP.CallOpts)
}

// Lid is a free data retrieval call binding the contract method 0x26d2addc.
//
// Solidity: function lid() view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Lid() (*big.Int, error) {
	return _MCDFLAP.Contract.Lid(&_MCDFLAP.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Live() (*big.Int, error) {
	return _MCDFLAP.Contract.Live(&_MCDFLAP.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Live() (*big.Int, error) {
	return _MCDFLAP.Contract.Live(&_MCDFLAP.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLAP *MCDFLAPCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "tau")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLAP *MCDFLAPSession) Tau() (*big.Int, error) {
	return _MCDFLAP.Contract.Tau(&_MCDFLAP.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLAP *MCDFLAPCallerSession) Tau() (*big.Int, error) {
	return _MCDFLAP.Contract.Tau(&_MCDFLAP.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLAP *MCDFLAPCaller) Ttl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "ttl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLAP *MCDFLAPSession) Ttl() (*big.Int, error) {
	return _MCDFLAP.Contract.Ttl(&_MCDFLAP.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLAP *MCDFLAPCallerSession) Ttl() (*big.Int, error) {
	return _MCDFLAP.Contract.Ttl(&_MCDFLAP.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLAP *MCDFLAPCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLAP *MCDFLAPSession) Vat() (common.Address, error) {
	return _MCDFLAP.Contract.Vat(&_MCDFLAP.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLAP *MCDFLAPCallerSession) Vat() (common.Address, error) {
	return _MCDFLAP.Contract.Vat(&_MCDFLAP.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLAP *MCDFLAPCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLAP.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLAP *MCDFLAPSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLAP.Contract.Wards(&_MCDFLAP.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLAP *MCDFLAPCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLAP.Contract.Wards(&_MCDFLAP.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_MCDFLAP *MCDFLAPTransactor) Cage(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "cage", rad)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_MCDFLAP *MCDFLAPSession) Cage(rad *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Cage(&_MCDFLAP.TransactOpts, rad)
}

// Cage is a paid mutator transaction binding the contract method 0xa2f91af2.
//
// Solidity: function cage(uint256 rad) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Cage(rad *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Cage(&_MCDFLAP.TransactOpts, rad)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactor) Deal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "deal", id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLAP *MCDFLAPSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Deal(&_MCDFLAP.TransactOpts, id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Deal(&_MCDFLAP.TransactOpts, id)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLAP *MCDFLAPTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLAP *MCDFLAPSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Deny(&_MCDFLAP.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Deny(&_MCDFLAP.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLAP *MCDFLAPTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLAP *MCDFLAPSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.File(&_MCDFLAP.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.File(&_MCDFLAP.TransactOpts, what, data)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLAP *MCDFLAPTransactor) Kick(opts *bind.TransactOpts, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "kick", lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLAP *MCDFLAPSession) Kick(lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Kick(&_MCDFLAP.TransactOpts, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xca40c419.
//
// Solidity: function kick(uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLAP *MCDFLAPTransactorSession) Kick(lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Kick(&_MCDFLAP.TransactOpts, lot, bid)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLAP *MCDFLAPTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLAP *MCDFLAPSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Rely(&_MCDFLAP.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Rely(&_MCDFLAP.TransactOpts, usr)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLAP *MCDFLAPTransactor) Tend(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "tend", id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLAP *MCDFLAPSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Tend(&_MCDFLAP.TransactOpts, id, lot, bid)
}

// Tend is a paid mutator transaction binding the contract method 0x4b43ed12.
//
// Solidity: function tend(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Tend(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Tend(&_MCDFLAP.TransactOpts, id, lot, bid)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactor) Tick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "tick", id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLAP *MCDFLAPSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Tick(&_MCDFLAP.TransactOpts, id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Tick(&_MCDFLAP.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLAP *MCDFLAPSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Yank(&_MCDFLAP.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLAP *MCDFLAPTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _MCDFLAP.Contract.Yank(&_MCDFLAP.TransactOpts, id)
}

// MCDFLAPKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the MCDFLAP contract.
type MCDFLAPKickIterator struct {
	Event *MCDFLAPKick // Event containing the contract specifics and raw log

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
func (it *MCDFLAPKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLAPKick)
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
		it.Event = new(MCDFLAPKick)
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
func (it *MCDFLAPKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLAPKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLAPKick represents a Kick event raised by the MCDFLAP contract.
type MCDFLAPKick struct {
	Id  *big.Int
	Lot *big.Int
	Bid *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_MCDFLAP *MCDFLAPFilterer) FilterKick(opts *bind.FilterOpts) (*MCDFLAPKickIterator, error) {

	logs, sub, err := _MCDFLAP.contract.FilterLogs(opts, "Kick")
	if err != nil {
		return nil, err
	}
	return &MCDFLAPKickIterator{contract: _MCDFLAP.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_MCDFLAP *MCDFLAPFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *MCDFLAPKick) (event.Subscription, error) {

	logs, sub, err := _MCDFLAP.contract.WatchLogs(opts, "Kick")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLAPKick)
				if err := _MCDFLAP.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0xe6dde59cbc017becba89714a037778d234a84ce7f0a137487142a007e580d609.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid)
func (_MCDFLAP *MCDFLAPFilterer) ParseKick(log types.Log) (*MCDFLAPKick, error) {
	event := new(MCDFLAPKick)
	if err := _MCDFLAP.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
