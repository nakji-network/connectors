// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proxyTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_old\",\"type\":\"address\"}],\"name\":\"ProxyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_depositId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sidechainAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_standard\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenNumber\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_withdrawId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenNumber\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sidechainAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"standard\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"tokenNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyTo\",\"type\":\"address\"}],\"name\":\"updateProxyTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeSession) Admin() (common.Address, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Bridge *BridgeCallerSession) Admin() (common.Address, error) {
	return _Bridge.Contract.Admin(&_Bridge.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridge *BridgeCaller) DepositCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridge *BridgeSession) DepositCount() (*big.Int, error) {
	return _Bridge.Contract.DepositCount(&_Bridge.CallOpts)
}

// DepositCount is a free data retrieval call binding the contract method 0x2dfdf0b5.
//
// Solidity: function depositCount() view returns(uint256)
func (_Bridge *BridgeCallerSession) DepositCount() (*big.Int, error) {
	return _Bridge.Contract.DepositCount(&_Bridge.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, address tokenAddress, address sidechainAddress, uint32 standard, uint256 tokenNumber)
func (_Bridge *BridgeCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner            common.Address
	TokenAddress     common.Address
	SidechainAddress common.Address
	Standard         uint32
	TokenNumber      *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Owner            common.Address
		TokenAddress     common.Address
		SidechainAddress common.Address
		Standard         uint32
		TokenNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SidechainAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Standard = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.TokenNumber = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, address tokenAddress, address sidechainAddress, uint32 standard, uint256 tokenNumber)
func (_Bridge *BridgeSession) Deposits(arg0 *big.Int) (struct {
	Owner            common.Address
	TokenAddress     common.Address
	SidechainAddress common.Address
	Standard         uint32
	TokenNumber      *big.Int
}, error) {
	return _Bridge.Contract.Deposits(&_Bridge.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(address owner, address tokenAddress, address sidechainAddress, uint32 standard, uint256 tokenNumber)
func (_Bridge *BridgeCallerSession) Deposits(arg0 *big.Int) (struct {
	Owner            common.Address
	TokenAddress     common.Address
	SidechainAddress common.Address
	Standard         uint32
	TokenNumber      *big.Int
}, error) {
	return _Bridge.Contract.Deposits(&_Bridge.CallOpts, arg0)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Bridge *BridgeCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Bridge *BridgeSession) Implementation() (common.Address, error) {
	return _Bridge.Contract.Implementation(&_Bridge.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Bridge *BridgeCallerSession) Implementation() (common.Address, error) {
	return _Bridge.Contract.Implementation(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256)
func (_Bridge *BridgeCaller) ProxyType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proxyType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256)
func (_Bridge *BridgeSession) ProxyType() (*big.Int, error) {
	return _Bridge.Contract.ProxyType(&_Bridge.CallOpts)
}

// ProxyType is a free data retrieval call binding the contract method 0x4555d5c9.
//
// Solidity: function proxyType() pure returns(uint256)
func (_Bridge *BridgeCallerSession) ProxyType() (*big.Int, error) {
	return _Bridge.Contract.ProxyType(&_Bridge.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bridge *BridgeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bridge *BridgeSession) Registry() (common.Address, error) {
	return _Bridge.Contract.Registry(&_Bridge.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Bridge *BridgeCallerSession) Registry() (common.Address, error) {
	return _Bridge.Contract.Registry(&_Bridge.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 ) view returns(address owner, address tokenAddress, uint256 tokenNumber)
func (_Bridge *BridgeCaller) Withdrawals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner        common.Address
	TokenAddress common.Address
	TokenNumber  *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "withdrawals", arg0)

	outstruct := new(struct {
		Owner        common.Address
		TokenAddress common.Address
		TokenNumber  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 ) view returns(address owner, address tokenAddress, uint256 tokenNumber)
func (_Bridge *BridgeSession) Withdrawals(arg0 *big.Int) (struct {
	Owner        common.Address
	TokenAddress common.Address
	TokenNumber  *big.Int
}, error) {
	return _Bridge.Contract.Withdrawals(&_Bridge.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 ) view returns(address owner, address tokenAddress, uint256 tokenNumber)
func (_Bridge *BridgeCallerSession) Withdrawals(arg0 *big.Int) (struct {
	Owner        common.Address
	TokenAddress common.Address
	TokenNumber  *big.Int
}, error) {
	return _Bridge.Contract.Withdrawals(&_Bridge.CallOpts, arg0)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Bridge *BridgeTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Bridge *BridgeSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdmin(&_Bridge.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Bridge *BridgeTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeAdmin(&_Bridge.TransactOpts, _newAdmin)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Bridge *BridgeTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Bridge *BridgeSession) RemoveAdmin() (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAdmin(&_Bridge.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Bridge *BridgeTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Bridge.Contract.RemoveAdmin(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// UpdateProxyTo is a paid mutator transaction binding the contract method 0xfd840de2.
//
// Solidity: function updateProxyTo(address _newProxyTo) returns()
func (_Bridge *BridgeTransactor) UpdateProxyTo(opts *bind.TransactOpts, _newProxyTo common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateProxyTo", _newProxyTo)
}

// UpdateProxyTo is a paid mutator transaction binding the contract method 0xfd840de2.
//
// Solidity: function updateProxyTo(address _newProxyTo) returns()
func (_Bridge *BridgeSession) UpdateProxyTo(_newProxyTo common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateProxyTo(&_Bridge.TransactOpts, _newProxyTo)
}

// UpdateProxyTo is a paid mutator transaction binding the contract method 0xfd840de2.
//
// Solidity: function updateProxyTo(address _newProxyTo) returns()
func (_Bridge *BridgeTransactorSession) UpdateProxyTo(_newProxyTo common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateProxyTo(&_Bridge.TransactOpts, _newProxyTo)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Bridge *BridgeTransactor) UpdateRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "updateRegistry", _registry)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Bridge *BridgeSession) UpdateRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateRegistry(&_Bridge.TransactOpts, _registry)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x1a5da6c8.
//
// Solidity: function updateRegistry(address _registry) returns()
func (_Bridge *BridgeTransactorSession) UpdateRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.UpdateRegistry(&_Bridge.TransactOpts, _registry)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Bridge *BridgeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Fallback(&_Bridge.TransactOpts, calldata)
}

// BridgeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Bridge contract.
type BridgeAdminChangedIterator struct {
	Event *BridgeAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminChanged)
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
		it.Event = new(BridgeAdminChanged)
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
func (it *BridgeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminChanged represents a AdminChanged event raised by the Bridge contract.
type BridgeAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Bridge *BridgeFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*BridgeAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeAdminChangedIterator{contract: _Bridge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Bridge *BridgeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Bridge *BridgeFilterer) ParseAdminChanged(log types.Log) (*BridgeAdminChanged, error) {
	event := new(BridgeAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Bridge contract.
type BridgeAdminRemovedIterator struct {
	Event *BridgeAdminRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAdminRemoved)
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
		it.Event = new(BridgeAdminRemoved)
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
func (it *BridgeAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAdminRemoved represents a AdminRemoved event raised by the Bridge contract.
type BridgeAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Bridge *BridgeFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*BridgeAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &BridgeAdminRemovedIterator{contract: _Bridge.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Bridge *BridgeFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *BridgeAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAdminRemoved)
				if err := _Bridge.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Bridge *BridgeFilterer) ParseAdminRemoved(log types.Log) (*BridgeAdminRemoved, error) {
	event := new(BridgeAdminRemoved)
	if err := _Bridge.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProxyUpdatedIterator is returned from FilterProxyUpdated and is used to iterate over the raw logs and unpacked data for ProxyUpdated events raised by the Bridge contract.
type BridgeProxyUpdatedIterator struct {
	Event *BridgeProxyUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeProxyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProxyUpdated)
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
		it.Event = new(BridgeProxyUpdated)
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
func (it *BridgeProxyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProxyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProxyUpdated represents a ProxyUpdated event raised by the Bridge contract.
type BridgeProxyUpdated struct {
	New common.Address
	Old common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProxyUpdated is a free log retrieval operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_Bridge *BridgeFilterer) FilterProxyUpdated(opts *bind.FilterOpts, _new []common.Address, _old []common.Address) (*BridgeProxyUpdatedIterator, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProxyUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProxyUpdatedIterator{contract: _Bridge.contract, event: "ProxyUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyUpdated is a free log subscription operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_Bridge *BridgeFilterer) WatchProxyUpdated(opts *bind.WatchOpts, sink chan<- *BridgeProxyUpdated, _new []common.Address, _old []common.Address) (event.Subscription, error) {

	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}
	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProxyUpdated", _newRule, _oldRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProxyUpdated)
				if err := _Bridge.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
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

// ParseProxyUpdated is a log parse operation binding the contract event 0xd32d24edea94f55e932d9a008afc425a8561462d1b1f57bc6e508e9a6b9509e1.
//
// Solidity: event ProxyUpdated(address indexed _new, address indexed _old)
func (_Bridge *BridgeFilterer) ParseProxyUpdated(log types.Log) (*BridgeProxyUpdated, error) {
	event := new(BridgeProxyUpdated)
	if err := _Bridge.contract.UnpackLog(event, "ProxyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the Bridge contract.
type BridgeTokenDepositedIterator struct {
	Event *BridgeTokenDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenDeposited)
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
		it.Event = new(BridgeTokenDeposited)
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
func (it *BridgeTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenDeposited represents a TokenDeposited event raised by the Bridge contract.
type BridgeTokenDeposited struct {
	DepositId        *big.Int
	Owner            common.Address
	TokenAddress     common.Address
	SidechainAddress common.Address
	Standard         uint32
	TokenNumber      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0x72848855a2461abf0dd243723dfcc9163eec2ea5215469d101c0d9c9ef58940d.
//
// Solidity: event TokenDeposited(uint256 indexed _depositId, address indexed _owner, address indexed _tokenAddress, address _sidechainAddress, uint32 _standard, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) FilterTokenDeposited(opts *bind.FilterOpts, _depositId []*big.Int, _owner []common.Address, _tokenAddress []common.Address) (*BridgeTokenDepositedIterator, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenDeposited", _depositIdRule, _ownerRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenDepositedIterator{contract: _Bridge.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0x72848855a2461abf0dd243723dfcc9163eec2ea5215469d101c0d9c9ef58940d.
//
// Solidity: event TokenDeposited(uint256 indexed _depositId, address indexed _owner, address indexed _tokenAddress, address _sidechainAddress, uint32 _standard, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *BridgeTokenDeposited, _depositId []*big.Int, _owner []common.Address, _tokenAddress []common.Address) (event.Subscription, error) {

	var _depositIdRule []interface{}
	for _, _depositIdItem := range _depositId {
		_depositIdRule = append(_depositIdRule, _depositIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenDeposited", _depositIdRule, _ownerRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenDeposited)
				if err := _Bridge.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
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

// ParseTokenDeposited is a log parse operation binding the contract event 0x72848855a2461abf0dd243723dfcc9163eec2ea5215469d101c0d9c9ef58940d.
//
// Solidity: event TokenDeposited(uint256 indexed _depositId, address indexed _owner, address indexed _tokenAddress, address _sidechainAddress, uint32 _standard, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) ParseTokenDeposited(log types.Log) (*BridgeTokenDeposited, error) {
	event := new(BridgeTokenDeposited)
	if err := _Bridge.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenWithdrewIterator is returned from FilterTokenWithdrew and is used to iterate over the raw logs and unpacked data for TokenWithdrew events raised by the Bridge contract.
type BridgeTokenWithdrewIterator struct {
	Event *BridgeTokenWithdrew // Event containing the contract specifics and raw log

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
func (it *BridgeTokenWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenWithdrew)
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
		it.Event = new(BridgeTokenWithdrew)
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
func (it *BridgeTokenWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenWithdrew represents a TokenWithdrew event raised by the Bridge contract.
type BridgeTokenWithdrew struct {
	WithdrawId   *big.Int
	Owner        common.Address
	TokenAddress common.Address
	TokenNumber  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrew is a free log retrieval operation binding the contract event 0x86174ea401f083b9bb1bdebca3068f27fb023c7091365ed2a8a02b8d75cf0e52.
//
// Solidity: event TokenWithdrew(uint256 indexed _withdrawId, address indexed _owner, address indexed _tokenAddress, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) FilterTokenWithdrew(opts *bind.FilterOpts, _withdrawId []*big.Int, _owner []common.Address, _tokenAddress []common.Address) (*BridgeTokenWithdrewIterator, error) {

	var _withdrawIdRule []interface{}
	for _, _withdrawIdItem := range _withdrawId {
		_withdrawIdRule = append(_withdrawIdRule, _withdrawIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenWithdrew", _withdrawIdRule, _ownerRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenWithdrewIterator{contract: _Bridge.contract, event: "TokenWithdrew", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrew is a free log subscription operation binding the contract event 0x86174ea401f083b9bb1bdebca3068f27fb023c7091365ed2a8a02b8d75cf0e52.
//
// Solidity: event TokenWithdrew(uint256 indexed _withdrawId, address indexed _owner, address indexed _tokenAddress, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) WatchTokenWithdrew(opts *bind.WatchOpts, sink chan<- *BridgeTokenWithdrew, _withdrawId []*big.Int, _owner []common.Address, _tokenAddress []common.Address) (event.Subscription, error) {

	var _withdrawIdRule []interface{}
	for _, _withdrawIdItem := range _withdrawId {
		_withdrawIdRule = append(_withdrawIdRule, _withdrawIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenWithdrew", _withdrawIdRule, _ownerRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenWithdrew)
				if err := _Bridge.contract.UnpackLog(event, "TokenWithdrew", log); err != nil {
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

// ParseTokenWithdrew is a log parse operation binding the contract event 0x86174ea401f083b9bb1bdebca3068f27fb023c7091365ed2a8a02b8d75cf0e52.
//
// Solidity: event TokenWithdrew(uint256 indexed _withdrawId, address indexed _owner, address indexed _tokenAddress, uint256 _tokenNumber)
func (_Bridge *BridgeFilterer) ParseTokenWithdrew(log types.Log) (*BridgeTokenWithdrew, error) {
	event := new(BridgeTokenWithdrew)
	if err := _Bridge.contract.UnpackLog(event, "TokenWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
