// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_FLOP

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

// MCDFLOPABI is the input ABI used to generate the binding from.
const MCDFLOPABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gem_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"beg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tic\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"end\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"dent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gal\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pad\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tau\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tick\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ttl\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDFLOP is an auto generated Go binding around an Ethereum contract.
type MCDFLOP struct {
	MCDFLOPCaller     // Read-only binding to the contract
	MCDFLOPTransactor // Write-only binding to the contract
	MCDFLOPFilterer   // Log filterer for contract events
}

// MCDFLOPCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDFLOPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLOPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDFLOPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLOPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDFLOPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLOPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDFLOPSession struct {
	Contract     *MCDFLOP          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDFLOPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDFLOPCallerSession struct {
	Contract *MCDFLOPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MCDFLOPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDFLOPTransactorSession struct {
	Contract     *MCDFLOPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MCDFLOPRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDFLOPRaw struct {
	Contract *MCDFLOP // Generic contract binding to access the raw methods on
}

// MCDFLOPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDFLOPCallerRaw struct {
	Contract *MCDFLOPCaller // Generic read-only contract binding to access the raw methods on
}

// MCDFLOPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDFLOPTransactorRaw struct {
	Contract *MCDFLOPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDFLOP creates a new instance of MCDFLOP, bound to a specific deployed contract.
func NewMCDFLOP(address common.Address, backend bind.ContractBackend) (*MCDFLOP, error) {
	contract, err := bindMCDFLOP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDFLOP{MCDFLOPCaller: MCDFLOPCaller{contract: contract}, MCDFLOPTransactor: MCDFLOPTransactor{contract: contract}, MCDFLOPFilterer: MCDFLOPFilterer{contract: contract}}, nil
}

// NewMCDFLOPCaller creates a new read-only instance of MCDFLOP, bound to a specific deployed contract.
func NewMCDFLOPCaller(address common.Address, caller bind.ContractCaller) (*MCDFLOPCaller, error) {
	contract, err := bindMCDFLOP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLOPCaller{contract: contract}, nil
}

// NewMCDFLOPTransactor creates a new write-only instance of MCDFLOP, bound to a specific deployed contract.
func NewMCDFLOPTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDFLOPTransactor, error) {
	contract, err := bindMCDFLOP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLOPTransactor{contract: contract}, nil
}

// NewMCDFLOPFilterer creates a new log filterer instance of MCDFLOP, bound to a specific deployed contract.
func NewMCDFLOPFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDFLOPFilterer, error) {
	contract, err := bindMCDFLOP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDFLOPFilterer{contract: contract}, nil
}

// bindMCDFLOP binds a generic wrapper to an already deployed contract.
func bindMCDFLOP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDFLOPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLOP *MCDFLOPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLOP.Contract.MCDFLOPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLOP *MCDFLOPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLOP.Contract.MCDFLOPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLOP *MCDFLOPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLOP.Contract.MCDFLOPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLOP *MCDFLOPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLOP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLOP *MCDFLOPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLOP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLOP *MCDFLOPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLOP.Contract.contract.Transact(opts, method, params...)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLOP *MCDFLOPCaller) Beg(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "beg")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLOP *MCDFLOPSession) Beg() (*big.Int, error) {
	return _MCDFLOP.Contract.Beg(&_MCDFLOP.CallOpts)
}

// Beg is a free data retrieval call binding the contract method 0x7d780d82.
//
// Solidity: function beg() view returns(uint256)
func (_MCDFLOP *MCDFLOPCallerSession) Beg() (*big.Int, error) {
	return _MCDFLOP.Contract.Beg(&_MCDFLOP.CallOpts)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_MCDFLOP *MCDFLOPCaller) Bids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "bids", arg0)

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
func (_MCDFLOP *MCDFLOPSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _MCDFLOP.Contract.Bids(&_MCDFLOP.CallOpts, arg0)
}

// Bids is a free data retrieval call binding the contract method 0x4423c5f1.
//
// Solidity: function bids(uint256 ) view returns(uint256 bid, uint256 lot, address guy, uint48 tic, uint48 end)
func (_MCDFLOP *MCDFLOPCallerSession) Bids(arg0 *big.Int) (struct {
	Bid *big.Int
	Lot *big.Int
	Guy common.Address
	Tic *big.Int
	End *big.Int
}, error) {
	return _MCDFLOP.Contract.Bids(&_MCDFLOP.CallOpts, arg0)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLOP *MCDFLOPCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLOP *MCDFLOPSession) Gem() (common.Address, error) {
	return _MCDFLOP.Contract.Gem(&_MCDFLOP.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDFLOP *MCDFLOPCallerSession) Gem() (common.Address, error) {
	return _MCDFLOP.Contract.Gem(&_MCDFLOP.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLOP *MCDFLOPCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "kicks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLOP *MCDFLOPSession) Kicks() (*big.Int, error) {
	return _MCDFLOP.Contract.Kicks(&_MCDFLOP.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_MCDFLOP *MCDFLOPCallerSession) Kicks() (*big.Int, error) {
	return _MCDFLOP.Contract.Kicks(&_MCDFLOP.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLOP *MCDFLOPCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLOP *MCDFLOPSession) Live() (*big.Int, error) {
	return _MCDFLOP.Contract.Live(&_MCDFLOP.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDFLOP *MCDFLOPCallerSession) Live() (*big.Int, error) {
	return _MCDFLOP.Contract.Live(&_MCDFLOP.CallOpts)
}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() view returns(uint256)
func (_MCDFLOP *MCDFLOPCaller) Pad(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "pad")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() view returns(uint256)
func (_MCDFLOP *MCDFLOPSession) Pad() (*big.Int, error) {
	return _MCDFLOP.Contract.Pad(&_MCDFLOP.CallOpts)
}

// Pad is a free data retrieval call binding the contract method 0x9361266c.
//
// Solidity: function pad() view returns(uint256)
func (_MCDFLOP *MCDFLOPCallerSession) Pad() (*big.Int, error) {
	return _MCDFLOP.Contract.Pad(&_MCDFLOP.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLOP *MCDFLOPCaller) Tau(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "tau")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLOP *MCDFLOPSession) Tau() (*big.Int, error) {
	return _MCDFLOP.Contract.Tau(&_MCDFLOP.CallOpts)
}

// Tau is a free data retrieval call binding the contract method 0xcfc4af55.
//
// Solidity: function tau() view returns(uint48)
func (_MCDFLOP *MCDFLOPCallerSession) Tau() (*big.Int, error) {
	return _MCDFLOP.Contract.Tau(&_MCDFLOP.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLOP *MCDFLOPCaller) Ttl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "ttl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLOP *MCDFLOPSession) Ttl() (*big.Int, error) {
	return _MCDFLOP.Contract.Ttl(&_MCDFLOP.CallOpts)
}

// Ttl is a free data retrieval call binding the contract method 0x4e8b1dd5.
//
// Solidity: function ttl() view returns(uint48)
func (_MCDFLOP *MCDFLOPCallerSession) Ttl() (*big.Int, error) {
	return _MCDFLOP.Contract.Ttl(&_MCDFLOP.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLOP *MCDFLOPCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLOP *MCDFLOPSession) Vat() (common.Address, error) {
	return _MCDFLOP.Contract.Vat(&_MCDFLOP.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLOP *MCDFLOPCallerSession) Vat() (common.Address, error) {
	return _MCDFLOP.Contract.Vat(&_MCDFLOP.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLOP *MCDFLOPCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLOP *MCDFLOPSession) Vow() (common.Address, error) {
	return _MCDFLOP.Contract.Vow(&_MCDFLOP.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLOP *MCDFLOPCallerSession) Vow() (common.Address, error) {
	return _MCDFLOP.Contract.Vow(&_MCDFLOP.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLOP *MCDFLOPCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLOP.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLOP *MCDFLOPSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLOP.Contract.Wards(&_MCDFLOP.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLOP *MCDFLOPCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLOP.Contract.Wards(&_MCDFLOP.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDFLOP *MCDFLOPTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDFLOP *MCDFLOPSession) Cage() (*types.Transaction, error) {
	return _MCDFLOP.Contract.Cage(&_MCDFLOP.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDFLOP.Contract.Cage(&_MCDFLOP.TransactOpts)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactor) Deal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "deal", id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLOP *MCDFLOPSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Deal(&_MCDFLOP.TransactOpts, id)
}

// Deal is a paid mutator transaction binding the contract method 0xc959c42b.
//
// Solidity: function deal(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Deal(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Deal(&_MCDFLOP.TransactOpts, id)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLOP *MCDFLOPTransactor) Dent(opts *bind.TransactOpts, id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "dent", id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLOP *MCDFLOPSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Dent(&_MCDFLOP.TransactOpts, id, lot, bid)
}

// Dent is a paid mutator transaction binding the contract method 0x5ff3a382.
//
// Solidity: function dent(uint256 id, uint256 lot, uint256 bid) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Dent(id *big.Int, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Dent(&_MCDFLOP.TransactOpts, id, lot, bid)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLOP *MCDFLOPTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLOP *MCDFLOPSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Deny(&_MCDFLOP.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Deny(&_MCDFLOP.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLOP *MCDFLOPTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLOP *MCDFLOPSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.File(&_MCDFLOP.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.File(&_MCDFLOP.TransactOpts, what, data)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLOP *MCDFLOPTransactor) Kick(opts *bind.TransactOpts, gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "kick", gal, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLOP *MCDFLOPSession) Kick(gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Kick(&_MCDFLOP.TransactOpts, gal, lot, bid)
}

// Kick is a paid mutator transaction binding the contract method 0xb7e9cd24.
//
// Solidity: function kick(address gal, uint256 lot, uint256 bid) returns(uint256 id)
func (_MCDFLOP *MCDFLOPTransactorSession) Kick(gal common.Address, lot *big.Int, bid *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Kick(&_MCDFLOP.TransactOpts, gal, lot, bid)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLOP *MCDFLOPTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLOP *MCDFLOPSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Rely(&_MCDFLOP.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Rely(&_MCDFLOP.TransactOpts, usr)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactor) Tick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "tick", id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLOP *MCDFLOPSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Tick(&_MCDFLOP.TransactOpts, id)
}

// Tick is a paid mutator transaction binding the contract method 0xfc7b6aee.
//
// Solidity: function tick(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Tick(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Tick(&_MCDFLOP.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLOP *MCDFLOPSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Yank(&_MCDFLOP.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_MCDFLOP *MCDFLOPTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _MCDFLOP.Contract.Yank(&_MCDFLOP.TransactOpts, id)
}

// MCDFLOPKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the MCDFLOP contract.
type MCDFLOPKickIterator struct {
	Event *MCDFLOPKick // Event containing the contract specifics and raw log

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
func (it *MCDFLOPKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLOPKick)
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
		it.Event = new(MCDFLOPKick)
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
func (it *MCDFLOPKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLOPKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLOPKick represents a Kick event raised by the MCDFLOP contract.
type MCDFLOPKick struct {
	Id  *big.Int
	Lot *big.Int
	Bid *big.Int
	Gal common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_MCDFLOP *MCDFLOPFilterer) FilterKick(opts *bind.FilterOpts, gal []common.Address) (*MCDFLOPKickIterator, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _MCDFLOP.contract.FilterLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLOPKickIterator{contract: _MCDFLOP.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_MCDFLOP *MCDFLOPFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *MCDFLOPKick, gal []common.Address) (event.Subscription, error) {

	var galRule []interface{}
	for _, galItem := range gal {
		galRule = append(galRule, galItem)
	}

	logs, sub, err := _MCDFLOP.contract.WatchLogs(opts, "Kick", galRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLOPKick)
				if err := _MCDFLOP.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0x7e8881001566f9f89aedb9c5dc3d856a2b81e5235a8196413ed484be91cc0df6.
//
// Solidity: event Kick(uint256 id, uint256 lot, uint256 bid, address indexed gal)
func (_MCDFLOP *MCDFLOPFilterer) ParseKick(log types.Log) (*MCDFLOPKick, error) {
	event := new(MCDFLOPKick)
	if err := _MCDFLOP.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
