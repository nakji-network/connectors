// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ARBITRUM_L2_GOV_RELAY

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

// ARBITRUML2GOVRELAYABI is the input ABI used to generate the binding from.
const ARBITRUML2GOVRELAYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2DAITokenBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_escrow\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2DAITokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ARBITRUML2GOVRELAY is an auto generated Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAY struct {
	ARBITRUML2GOVRELAYCaller     // Read-only binding to the contract
	ARBITRUML2GOVRELAYTransactor // Write-only binding to the contract
	ARBITRUML2GOVRELAYFilterer   // Log filterer for contract events
}

// ARBITRUML2GOVRELAYCaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2GOVRELAYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2GOVRELAYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUML2GOVRELAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUML2GOVRELAYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUML2GOVRELAYSession struct {
	Contract     *ARBITRUML2GOVRELAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ARBITRUML2GOVRELAYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUML2GOVRELAYCallerSession struct {
	Contract *ARBITRUML2GOVRELAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ARBITRUML2GOVRELAYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUML2GOVRELAYTransactorSession struct {
	Contract     *ARBITRUML2GOVRELAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ARBITRUML2GOVRELAYRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAYRaw struct {
	Contract *ARBITRUML2GOVRELAY // Generic contract binding to access the raw methods on
}

// ARBITRUML2GOVRELAYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAYCallerRaw struct {
	Contract *ARBITRUML2GOVRELAYCaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUML2GOVRELAYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUML2GOVRELAYTransactorRaw struct {
	Contract *ARBITRUML2GOVRELAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUML2GOVRELAY creates a new instance of ARBITRUML2GOVRELAY, bound to a specific deployed contract.
func NewARBITRUML2GOVRELAY(address common.Address, backend bind.ContractBackend) (*ARBITRUML2GOVRELAY, error) {
	contract, err := bindARBITRUML2GOVRELAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAY{ARBITRUML2GOVRELAYCaller: ARBITRUML2GOVRELAYCaller{contract: contract}, ARBITRUML2GOVRELAYTransactor: ARBITRUML2GOVRELAYTransactor{contract: contract}, ARBITRUML2GOVRELAYFilterer: ARBITRUML2GOVRELAYFilterer{contract: contract}}, nil
}

// NewARBITRUML2GOVRELAYCaller creates a new read-only instance of ARBITRUML2GOVRELAY, bound to a specific deployed contract.
func NewARBITRUML2GOVRELAYCaller(address common.Address, caller bind.ContractCaller) (*ARBITRUML2GOVRELAYCaller, error) {
	contract, err := bindARBITRUML2GOVRELAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYCaller{contract: contract}, nil
}

// NewARBITRUML2GOVRELAYTransactor creates a new write-only instance of ARBITRUML2GOVRELAY, bound to a specific deployed contract.
func NewARBITRUML2GOVRELAYTransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUML2GOVRELAYTransactor, error) {
	contract, err := bindARBITRUML2GOVRELAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYTransactor{contract: contract}, nil
}

// NewARBITRUML2GOVRELAYFilterer creates a new log filterer instance of ARBITRUML2GOVRELAY, bound to a specific deployed contract.
func NewARBITRUML2GOVRELAYFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUML2GOVRELAYFilterer, error) {
	contract, err := bindARBITRUML2GOVRELAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYFilterer{contract: contract}, nil
}

// bindARBITRUML2GOVRELAY binds a generic wrapper to an already deployed contract.
func bindARBITRUML2GOVRELAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUML2GOVRELAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2GOVRELAY.Contract.ARBITRUML2GOVRELAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.ARBITRUML2GOVRELAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.ARBITRUML2GOVRELAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUML2GOVRELAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.contract.Transact(opts, method, params...)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Escrow() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.Escrow(&_ARBITRUML2GOVRELAY.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) Escrow() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.Escrow(&_ARBITRUML2GOVRELAY.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) IsOpen() (*big.Int, error) {
	return _ARBITRUML2GOVRELAY.Contract.IsOpen(&_ARBITRUML2GOVRELAY.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) IsOpen() (*big.Int, error) {
	return _ARBITRUML2GOVRELAY.Contract.IsOpen(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) L1Token() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L1Token(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) L1Token() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L1Token(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) L2DAITokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "l2DAITokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) L2DAITokenBridge() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L2DAITokenBridge(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) L2DAITokenBridge() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L2DAITokenBridge(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) L2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "l2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) L2Token() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L2Token(&_ARBITRUML2GOVRELAY.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) L2Token() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.L2Token(&_ARBITRUML2GOVRELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Messenger() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.Messenger(&_ARBITRUML2GOVRELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) Messenger() (common.Address, error) {
	return _ARBITRUML2GOVRELAY.Contract.Messenger(&_ARBITRUML2GOVRELAY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUML2GOVRELAY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2GOVRELAY.Contract.Wards(&_ARBITRUML2GOVRELAY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUML2GOVRELAY.Contract.Wards(&_ARBITRUML2GOVRELAY.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Close() (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Close(&_ARBITRUML2GOVRELAY.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) Close() (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Close(&_ARBITRUML2GOVRELAY.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Deny(&_ARBITRUML2GOVRELAY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Deny(&_ARBITRUML2GOVRELAY.TransactOpts, usr)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.DepositERC20(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.DepositERC20(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.DepositERC20To(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.DepositERC20To(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.FinalizeERC20Withdrawal(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.FinalizeERC20Withdrawal(&_ARBITRUML2GOVRELAY.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Rely(&_ARBITRUML2GOVRELAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUML2GOVRELAY.Contract.Rely(&_ARBITRUML2GOVRELAY.TransactOpts, usr)
}

// ARBITRUML2GOVRELAYClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYClosedIterator struct {
	Event *ARBITRUML2GOVRELAYClosed // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2GOVRELAYClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2GOVRELAYClosed)
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
		it.Event = new(ARBITRUML2GOVRELAYClosed)
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
func (it *ARBITRUML2GOVRELAYClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2GOVRELAYClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2GOVRELAYClosed represents a Closed event raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) FilterClosed(opts *bind.FilterOpts) (*ARBITRUML2GOVRELAYClosedIterator, error) {

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYClosedIterator{contract: _ARBITRUML2GOVRELAY.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *ARBITRUML2GOVRELAYClosed) (event.Subscription, error) {

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2GOVRELAYClosed)
				if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) ParseClosed(log types.Log) (*ARBITRUML2GOVRELAYClosed, error) {
	event := new(ARBITRUML2GOVRELAYClosed)
	if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2GOVRELAYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYDenyIterator struct {
	Event *ARBITRUML2GOVRELAYDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2GOVRELAYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2GOVRELAYDeny)
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
		it.Event = new(ARBITRUML2GOVRELAYDeny)
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
func (it *ARBITRUML2GOVRELAYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2GOVRELAYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2GOVRELAYDeny represents a Deny event raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2GOVRELAYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYDenyIterator{contract: _ARBITRUML2GOVRELAY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUML2GOVRELAYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2GOVRELAYDeny)
				if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) ParseDeny(log types.Log) (*ARBITRUML2GOVRELAYDeny, error) {
	event := new(ARBITRUML2GOVRELAYDeny)
	if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2GOVRELAYERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYERC20DepositInitiatedIterator struct {
	Event *ARBITRUML2GOVRELAYERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2GOVRELAYERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2GOVRELAYERC20DepositInitiated)
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
		it.Event = new(ARBITRUML2GOVRELAYERC20DepositInitiated)
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
func (it *ARBITRUML2GOVRELAYERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2GOVRELAYERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2GOVRELAYERC20DepositInitiated represents a ERC20DepositInitiated event raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYERC20DepositInitiated struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*ARBITRUML2GOVRELAYERC20DepositInitiatedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYERC20DepositInitiatedIterator{contract: _ARBITRUML2GOVRELAY.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *ARBITRUML2GOVRELAYERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2GOVRELAYERC20DepositInitiated)
				if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) ParseERC20DepositInitiated(log types.Log) (*ARBITRUML2GOVRELAYERC20DepositInitiated, error) {
	event := new(ARBITRUML2GOVRELAYERC20DepositInitiated)
	if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator struct {
	Event *ARBITRUML2GOVRELAYERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2GOVRELAYERC20WithdrawalFinalized)
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
		it.Event = new(ARBITRUML2GOVRELAYERC20WithdrawalFinalized)
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
func (it *ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2GOVRELAYERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYERC20WithdrawalFinalized struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYERC20WithdrawalFinalizedIterator{contract: _ARBITRUML2GOVRELAY.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *ARBITRUML2GOVRELAYERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

	var _l1TokenRule []interface{}
	for _, _l1TokenItem := range _l1Token {
		_l1TokenRule = append(_l1TokenRule, _l1TokenItem)
	}
	var _l2TokenRule []interface{}
	for _, _l2TokenItem := range _l2Token {
		_l2TokenRule = append(_l2TokenRule, _l2TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2GOVRELAYERC20WithdrawalFinalized)
				if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*ARBITRUML2GOVRELAYERC20WithdrawalFinalized, error) {
	event := new(ARBITRUML2GOVRELAYERC20WithdrawalFinalized)
	if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUML2GOVRELAYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYRelyIterator struct {
	Event *ARBITRUML2GOVRELAYRely // Event containing the contract specifics and raw log

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
func (it *ARBITRUML2GOVRELAYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUML2GOVRELAYRely)
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
		it.Event = new(ARBITRUML2GOVRELAYRely)
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
func (it *ARBITRUML2GOVRELAYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUML2GOVRELAYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUML2GOVRELAYRely represents a Rely event raised by the ARBITRUML2GOVRELAY contract.
type ARBITRUML2GOVRELAYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUML2GOVRELAYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUML2GOVRELAYRelyIterator{contract: _ARBITRUML2GOVRELAY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUML2GOVRELAYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUML2GOVRELAY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUML2GOVRELAYRely)
				if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_ARBITRUML2GOVRELAY *ARBITRUML2GOVRELAYFilterer) ParseRely(log types.Log) (*ARBITRUML2GOVRELAYRely, error) {
	event := new(ARBITRUML2GOVRELAYRely)
	if err := _ARBITRUML2GOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
