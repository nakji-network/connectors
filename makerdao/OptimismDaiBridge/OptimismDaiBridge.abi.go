// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OptimismDaiBridge

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

// OPTIMISMDAIBRIDGEABI is the input ABI used to generate the binding from.
const OPTIMISMDAIBRIDGEABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2DAITokenBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_escrow\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2DAITokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISMDAIBRIDGE is an auto generated Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGE struct {
	OPTIMISMDAIBRIDGECaller     // Read-only binding to the contract
	OPTIMISMDAIBRIDGETransactor // Write-only binding to the contract
	OPTIMISMDAIBRIDGEFilterer   // Log filterer for contract events
}

// OPTIMISMDAIBRIDGECaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMDAIBRIDGETransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMDAIBRIDGEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISMDAIBRIDGEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISMDAIBRIDGESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISMDAIBRIDGESession struct {
	Contract     *OPTIMISMDAIBRIDGE // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OPTIMISMDAIBRIDGECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISMDAIBRIDGECallerSession struct {
	Contract *OPTIMISMDAIBRIDGECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// OPTIMISMDAIBRIDGETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISMDAIBRIDGETransactorSession struct {
	Contract     *OPTIMISMDAIBRIDGETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// OPTIMISMDAIBRIDGERaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGERaw struct {
	Contract *OPTIMISMDAIBRIDGE // Generic contract binding to access the raw methods on
}

// OPTIMISMDAIBRIDGECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGECallerRaw struct {
	Contract *OPTIMISMDAIBRIDGECaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISMDAIBRIDGETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISMDAIBRIDGETransactorRaw struct {
	Contract *OPTIMISMDAIBRIDGETransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISMDAIBRIDGE creates a new instance of OPTIMISMDAIBRIDGE, bound to a specific deployed contract.
func NewOPTIMISMDAIBRIDGE(address common.Address, backend bind.ContractBackend) (*OPTIMISMDAIBRIDGE, error) {
	contract, err := bindOPTIMISMDAIBRIDGE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGE{OPTIMISMDAIBRIDGECaller: OPTIMISMDAIBRIDGECaller{contract: contract}, OPTIMISMDAIBRIDGETransactor: OPTIMISMDAIBRIDGETransactor{contract: contract}, OPTIMISMDAIBRIDGEFilterer: OPTIMISMDAIBRIDGEFilterer{contract: contract}}, nil
}

// NewOPTIMISMDAIBRIDGECaller creates a new read-only instance of OPTIMISMDAIBRIDGE, bound to a specific deployed contract.
func NewOPTIMISMDAIBRIDGECaller(address common.Address, caller bind.ContractCaller) (*OPTIMISMDAIBRIDGECaller, error) {
	contract, err := bindOPTIMISMDAIBRIDGE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGECaller{contract: contract}, nil
}

// NewOPTIMISMDAIBRIDGETransactor creates a new write-only instance of OPTIMISMDAIBRIDGE, bound to a specific deployed contract.
func NewOPTIMISMDAIBRIDGETransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISMDAIBRIDGETransactor, error) {
	contract, err := bindOPTIMISMDAIBRIDGE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGETransactor{contract: contract}, nil
}

// NewOPTIMISMDAIBRIDGEFilterer creates a new log filterer instance of OPTIMISMDAIBRIDGE, bound to a specific deployed contract.
func NewOPTIMISMDAIBRIDGEFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISMDAIBRIDGEFilterer, error) {
	contract, err := bindOPTIMISMDAIBRIDGE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGEFilterer{contract: contract}, nil
}

// bindOPTIMISMDAIBRIDGE binds a generic wrapper to an already deployed contract.
func bindOPTIMISMDAIBRIDGE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISMDAIBRIDGEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMDAIBRIDGE.Contract.OPTIMISMDAIBRIDGECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.OPTIMISMDAIBRIDGETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.OPTIMISMDAIBRIDGETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISMDAIBRIDGE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.contract.Transact(opts, method, params...)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Escrow() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Escrow(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) Escrow() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Escrow(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) IsOpen() (*big.Int, error) {
	return _OPTIMISMDAIBRIDGE.Contract.IsOpen(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) IsOpen() (*big.Int, error) {
	return _OPTIMISMDAIBRIDGE.Contract.IsOpen(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) L1Token() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L1Token(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) L1Token() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L1Token(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) L2DAITokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "l2DAITokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) L2DAITokenBridge() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L2DAITokenBridge(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) L2DAITokenBridge() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L2DAITokenBridge(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) L2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "l2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) L2Token() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L2Token(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) L2Token() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.L2Token(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Messenger() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Messenger(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) Messenger() (common.Address, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Messenger(&_OPTIMISMDAIBRIDGE.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISMDAIBRIDGE.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Wards(&_OPTIMISMDAIBRIDGE.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGECallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Wards(&_OPTIMISMDAIBRIDGE.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Close() (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Close(&_OPTIMISMDAIBRIDGE.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) Close() (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Close(&_OPTIMISMDAIBRIDGE.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Deny(&_OPTIMISMDAIBRIDGE.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Deny(&_OPTIMISMDAIBRIDGE.TransactOpts, usr)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.DepositERC20(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.DepositERC20(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.DepositERC20To(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.DepositERC20To(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.FinalizeERC20Withdrawal(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.FinalizeERC20Withdrawal(&_OPTIMISMDAIBRIDGE.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGESession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Rely(&_OPTIMISMDAIBRIDGE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGETransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISMDAIBRIDGE.Contract.Rely(&_OPTIMISMDAIBRIDGE.TransactOpts, usr)
}

// OPTIMISMDAIBRIDGEClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEClosedIterator struct {
	Event *OPTIMISMDAIBRIDGEClosed // Event containing the contract specifics and raw log

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
func (it *OPTIMISMDAIBRIDGEClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMDAIBRIDGEClosed)
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
		it.Event = new(OPTIMISMDAIBRIDGEClosed)
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
func (it *OPTIMISMDAIBRIDGEClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMDAIBRIDGEClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMDAIBRIDGEClosed represents a Closed event raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) FilterClosed(opts *bind.FilterOpts) (*OPTIMISMDAIBRIDGEClosedIterator, error) {

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGEClosedIterator{contract: _OPTIMISMDAIBRIDGE.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *OPTIMISMDAIBRIDGEClosed) (event.Subscription, error) {

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMDAIBRIDGEClosed)
				if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Closed", log); err != nil {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) ParseClosed(log types.Log) (*OPTIMISMDAIBRIDGEClosed, error) {
	event := new(OPTIMISMDAIBRIDGEClosed)
	if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMDAIBRIDGEDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEDenyIterator struct {
	Event *OPTIMISMDAIBRIDGEDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISMDAIBRIDGEDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMDAIBRIDGEDeny)
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
		it.Event = new(OPTIMISMDAIBRIDGEDeny)
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
func (it *OPTIMISMDAIBRIDGEDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMDAIBRIDGEDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMDAIBRIDGEDeny represents a Deny event raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMDAIBRIDGEDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGEDenyIterator{contract: _OPTIMISMDAIBRIDGE.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISMDAIBRIDGEDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMDAIBRIDGEDeny)
				if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) ParseDeny(log types.Log) (*OPTIMISMDAIBRIDGEDeny, error) {
	event := new(OPTIMISMDAIBRIDGEDeny)
	if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator struct {
	Event *OPTIMISMDAIBRIDGEERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMDAIBRIDGEERC20DepositInitiated)
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
		it.Event = new(OPTIMISMDAIBRIDGEERC20DepositInitiated)
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
func (it *OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMDAIBRIDGEERC20DepositInitiated represents a ERC20DepositInitiated event raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEERC20DepositInitiated struct {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator, error) {

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

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGEERC20DepositInitiatedIterator{contract: _OPTIMISMDAIBRIDGE.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *OPTIMISMDAIBRIDGEERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMDAIBRIDGEERC20DepositInitiated)
				if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) ParseERC20DepositInitiated(log types.Log) (*OPTIMISMDAIBRIDGEERC20DepositInitiated, error) {
	event := new(OPTIMISMDAIBRIDGEERC20DepositInitiated)
	if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator struct {
	Event *OPTIMISMDAIBRIDGEERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMDAIBRIDGEERC20WithdrawalFinalized)
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
		it.Event = new(OPTIMISMDAIBRIDGEERC20WithdrawalFinalized)
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
func (it *OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMDAIBRIDGEERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGEERC20WithdrawalFinalized struct {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator, error) {

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

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGEERC20WithdrawalFinalizedIterator{contract: _OPTIMISMDAIBRIDGE.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *OPTIMISMDAIBRIDGEERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMDAIBRIDGEERC20WithdrawalFinalized)
				if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*OPTIMISMDAIBRIDGEERC20WithdrawalFinalized, error) {
	event := new(OPTIMISMDAIBRIDGEERC20WithdrawalFinalized)
	if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISMDAIBRIDGERelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGERelyIterator struct {
	Event *OPTIMISMDAIBRIDGERely // Event containing the contract specifics and raw log

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
func (it *OPTIMISMDAIBRIDGERelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISMDAIBRIDGERely)
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
		it.Event = new(OPTIMISMDAIBRIDGERely)
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
func (it *OPTIMISMDAIBRIDGERelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISMDAIBRIDGERelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISMDAIBRIDGERely represents a Rely event raised by the OPTIMISMDAIBRIDGE contract.
type OPTIMISMDAIBRIDGERely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISMDAIBRIDGERelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISMDAIBRIDGERelyIterator{contract: _OPTIMISMDAIBRIDGE.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISMDAIBRIDGERely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISMDAIBRIDGE.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISMDAIBRIDGERely)
				if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISMDAIBRIDGE *OPTIMISMDAIBRIDGEFilterer) ParseRely(log types.Log) (*OPTIMISMDAIBRIDGERely, error) {
	event := new(OPTIMISMDAIBRIDGERely)
	if err := _OPTIMISMDAIBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
