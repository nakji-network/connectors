// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OPTIMISM_L2_GOVERNANCE_RELAY

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

// OPTIMISML2GOVERNANCERELAYABI is the input ABI used to generate the binding from.
const OPTIMISML2GOVERNANCERELAYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2DAITokenBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1messenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_escrow\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_l2Gas\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2DAITokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISML2GOVERNANCERELAY is an auto generated Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAY struct {
	OPTIMISML2GOVERNANCERELAYCaller     // Read-only binding to the contract
	OPTIMISML2GOVERNANCERELAYTransactor // Write-only binding to the contract
	OPTIMISML2GOVERNANCERELAYFilterer   // Log filterer for contract events
}

// OPTIMISML2GOVERNANCERELAYCaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2GOVERNANCERELAYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2GOVERNANCERELAYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISML2GOVERNANCERELAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2GOVERNANCERELAYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISML2GOVERNANCERELAYSession struct {
	Contract     *OPTIMISML2GOVERNANCERELAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OPTIMISML2GOVERNANCERELAYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISML2GOVERNANCERELAYCallerSession struct {
	Contract *OPTIMISML2GOVERNANCERELAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// OPTIMISML2GOVERNANCERELAYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISML2GOVERNANCERELAYTransactorSession struct {
	Contract     *OPTIMISML2GOVERNANCERELAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// OPTIMISML2GOVERNANCERELAYRaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAYRaw struct {
	Contract *OPTIMISML2GOVERNANCERELAY // Generic contract binding to access the raw methods on
}

// OPTIMISML2GOVERNANCERELAYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAYCallerRaw struct {
	Contract *OPTIMISML2GOVERNANCERELAYCaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISML2GOVERNANCERELAYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISML2GOVERNANCERELAYTransactorRaw struct {
	Contract *OPTIMISML2GOVERNANCERELAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISML2GOVERNANCERELAY creates a new instance of OPTIMISML2GOVERNANCERELAY, bound to a specific deployed contract.
func NewOPTIMISML2GOVERNANCERELAY(address common.Address, backend bind.ContractBackend) (*OPTIMISML2GOVERNANCERELAY, error) {
	contract, err := bindOPTIMISML2GOVERNANCERELAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAY{OPTIMISML2GOVERNANCERELAYCaller: OPTIMISML2GOVERNANCERELAYCaller{contract: contract}, OPTIMISML2GOVERNANCERELAYTransactor: OPTIMISML2GOVERNANCERELAYTransactor{contract: contract}, OPTIMISML2GOVERNANCERELAYFilterer: OPTIMISML2GOVERNANCERELAYFilterer{contract: contract}}, nil
}

// NewOPTIMISML2GOVERNANCERELAYCaller creates a new read-only instance of OPTIMISML2GOVERNANCERELAY, bound to a specific deployed contract.
func NewOPTIMISML2GOVERNANCERELAYCaller(address common.Address, caller bind.ContractCaller) (*OPTIMISML2GOVERNANCERELAYCaller, error) {
	contract, err := bindOPTIMISML2GOVERNANCERELAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYCaller{contract: contract}, nil
}

// NewOPTIMISML2GOVERNANCERELAYTransactor creates a new write-only instance of OPTIMISML2GOVERNANCERELAY, bound to a specific deployed contract.
func NewOPTIMISML2GOVERNANCERELAYTransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISML2GOVERNANCERELAYTransactor, error) {
	contract, err := bindOPTIMISML2GOVERNANCERELAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYTransactor{contract: contract}, nil
}

// NewOPTIMISML2GOVERNANCERELAYFilterer creates a new log filterer instance of OPTIMISML2GOVERNANCERELAY, bound to a specific deployed contract.
func NewOPTIMISML2GOVERNANCERELAYFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISML2GOVERNANCERELAYFilterer, error) {
	contract, err := bindOPTIMISML2GOVERNANCERELAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYFilterer{contract: contract}, nil
}

// bindOPTIMISML2GOVERNANCERELAY binds a generic wrapper to an already deployed contract.
func bindOPTIMISML2GOVERNANCERELAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISML2GOVERNANCERELAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2GOVERNANCERELAY.Contract.OPTIMISML2GOVERNANCERELAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.OPTIMISML2GOVERNANCERELAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.OPTIMISML2GOVERNANCERELAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2GOVERNANCERELAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.contract.Transact(opts, method, params...)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Escrow() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Escrow(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) Escrow() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Escrow(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) IsOpen() (*big.Int, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.IsOpen(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) IsOpen() (*big.Int, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.IsOpen(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) L1Token() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L1Token(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) L1Token() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L1Token(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) L2DAITokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "l2DAITokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) L2DAITokenBridge() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L2DAITokenBridge(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L2DAITokenBridge is a free data retrieval call binding the contract method 0x2e67f7c8.
//
// Solidity: function l2DAITokenBridge() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) L2DAITokenBridge() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L2DAITokenBridge(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) L2Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "l2Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) L2Token() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L2Token(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// L2Token is a free data retrieval call binding the contract method 0x56eff267.
//
// Solidity: function l2Token() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) L2Token() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.L2Token(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Messenger() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Messenger(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) Messenger() (common.Address, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Messenger(&_OPTIMISML2GOVERNANCERELAY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2GOVERNANCERELAY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Wards(&_OPTIMISML2GOVERNANCERELAY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Wards(&_OPTIMISML2GOVERNANCERELAY.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Close() (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Close(&_OPTIMISML2GOVERNANCERELAY.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) Close() (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Close(&_OPTIMISML2GOVERNANCERELAY.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Deny(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Deny(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, usr)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.DepositERC20(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.DepositERC20(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.DepositERC20To(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _l2Gas, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _l2Gas uint32, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.DepositERC20To(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _to, _amount, _l2Gas, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.FinalizeERC20Withdrawal(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _data) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.FinalizeERC20Withdrawal(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _data)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Rely(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2GOVERNANCERELAY.Contract.Rely(&_OPTIMISML2GOVERNANCERELAY.TransactOpts, usr)
}

// OPTIMISML2GOVERNANCERELAYClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYClosedIterator struct {
	Event *OPTIMISML2GOVERNANCERELAYClosed // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2GOVERNANCERELAYClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2GOVERNANCERELAYClosed)
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
		it.Event = new(OPTIMISML2GOVERNANCERELAYClosed)
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
func (it *OPTIMISML2GOVERNANCERELAYClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2GOVERNANCERELAYClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2GOVERNANCERELAYClosed represents a Closed event raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYClosed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) FilterClosed(opts *bind.FilterOpts) (*OPTIMISML2GOVERNANCERELAYClosedIterator, error) {

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYClosedIterator{contract: _OPTIMISML2GOVERNANCERELAY.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *OPTIMISML2GOVERNANCERELAYClosed) (event.Subscription, error) {

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2GOVERNANCERELAYClosed)
				if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Closed", log); err != nil {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) ParseClosed(log types.Log) (*OPTIMISML2GOVERNANCERELAYClosed, error) {
	event := new(OPTIMISML2GOVERNANCERELAYClosed)
	if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2GOVERNANCERELAYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYDenyIterator struct {
	Event *OPTIMISML2GOVERNANCERELAYDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2GOVERNANCERELAYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2GOVERNANCERELAYDeny)
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
		it.Event = new(OPTIMISML2GOVERNANCERELAYDeny)
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
func (it *OPTIMISML2GOVERNANCERELAYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2GOVERNANCERELAYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2GOVERNANCERELAYDeny represents a Deny event raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2GOVERNANCERELAYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYDenyIterator{contract: _OPTIMISML2GOVERNANCERELAY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISML2GOVERNANCERELAYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2GOVERNANCERELAYDeny)
				if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) ParseDeny(log types.Log) (*OPTIMISML2GOVERNANCERELAYDeny, error) {
	event := new(OPTIMISML2GOVERNANCERELAYDeny)
	if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator struct {
	Event *OPTIMISML2GOVERNANCERELAYERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2GOVERNANCERELAYERC20DepositInitiated)
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
		it.Event = new(OPTIMISML2GOVERNANCERELAYERC20DepositInitiated)
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
func (it *OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2GOVERNANCERELAYERC20DepositInitiated represents a ERC20DepositInitiated event raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYERC20DepositInitiated struct {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator, error) {

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

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.FilterLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYERC20DepositInitiatedIterator{contract: _OPTIMISML2GOVERNANCERELAY.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *OPTIMISML2GOVERNANCERELAYERC20DepositInitiated, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.WatchLogs(opts, "ERC20DepositInitiated", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2GOVERNANCERELAYERC20DepositInitiated)
				if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) ParseERC20DepositInitiated(log types.Log) (*OPTIMISML2GOVERNANCERELAYERC20DepositInitiated, error) {
	event := new(OPTIMISML2GOVERNANCERELAYERC20DepositInitiated)
	if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator struct {
	Event *OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized)
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
		it.Event = new(OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized)
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
func (it *OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized struct {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (*OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator, error) {

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

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalizedIterator{contract: _OPTIMISML2GOVERNANCERELAY.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed _l1Token, address indexed _l2Token, address indexed _from, address _to, uint256 _amount, bytes _data)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized, _l1Token []common.Address, _l2Token []common.Address, _from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", _l1TokenRule, _l2TokenRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized)
				if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized, error) {
	event := new(OPTIMISML2GOVERNANCERELAYERC20WithdrawalFinalized)
	if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2GOVERNANCERELAYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYRelyIterator struct {
	Event *OPTIMISML2GOVERNANCERELAYRely // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2GOVERNANCERELAYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2GOVERNANCERELAYRely)
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
		it.Event = new(OPTIMISML2GOVERNANCERELAYRely)
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
func (it *OPTIMISML2GOVERNANCERELAYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2GOVERNANCERELAYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2GOVERNANCERELAYRely represents a Rely event raised by the OPTIMISML2GOVERNANCERELAY contract.
type OPTIMISML2GOVERNANCERELAYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2GOVERNANCERELAYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2GOVERNANCERELAYRelyIterator{contract: _OPTIMISML2GOVERNANCERELAY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISML2GOVERNANCERELAYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2GOVERNANCERELAY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2GOVERNANCERELAYRely)
				if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISML2GOVERNANCERELAY *OPTIMISML2GOVERNANCERELAYFilterer) ParseRely(log types.Log) (*OPTIMISML2GOVERNANCERELAYRely, error) {
	event := new(OPTIMISML2GOVERNANCERELAYRely)
	if err := _OPTIMISML2GOVERNANCERELAY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
