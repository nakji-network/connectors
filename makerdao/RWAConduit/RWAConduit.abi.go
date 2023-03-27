// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RWAConduit

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

// RWACONDUITABI is the input ABI used to generate the binding from.
const RWACONDUITABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dai_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"drop_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tranche_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"Cull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Draw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Join\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recovered\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payBack\",\"type\":\"uint256\"}],\"name\":\"Recover\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Tell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payBack\",\"type\":\"uint256\"}],\"name\":\"Unwind\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Wipe\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"internalType\":\"contractEndLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractGemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"glad\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liq\",\"outputs\":[{\"internalType\":\"contractMIP21LiquidationLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractRedeemLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"quit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"}],\"name\":\"recover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"safe\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tab\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"tell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tranche\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"endEpoch\",\"type\":\"uint256\"}],\"name\":\"unwind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"urn\",\"outputs\":[{\"internalType\":\"contractMIP21UrnLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RWACONDUIT is an auto generated Go binding around an Ethereum contract.
type RWACONDUIT struct {
	RWACONDUITCaller     // Read-only binding to the contract
	RWACONDUITTransactor // Write-only binding to the contract
	RWACONDUITFilterer   // Log filterer for contract events
}

// RWACONDUITCaller is an auto generated read-only Go binding around an Ethereum contract.
type RWACONDUITCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWACONDUITTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RWACONDUITTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWACONDUITFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RWACONDUITFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWACONDUITSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RWACONDUITSession struct {
	Contract     *RWACONDUIT       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWACONDUITCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RWACONDUITCallerSession struct {
	Contract *RWACONDUITCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RWACONDUITTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RWACONDUITTransactorSession struct {
	Contract     *RWACONDUITTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RWACONDUITRaw is an auto generated low-level Go binding around an Ethereum contract.
type RWACONDUITRaw struct {
	Contract *RWACONDUIT // Generic contract binding to access the raw methods on
}

// RWACONDUITCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RWACONDUITCallerRaw struct {
	Contract *RWACONDUITCaller // Generic read-only contract binding to access the raw methods on
}

// RWACONDUITTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RWACONDUITTransactorRaw struct {
	Contract *RWACONDUITTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRWACONDUIT creates a new instance of RWACONDUIT, bound to a specific deployed contract.
func NewRWACONDUIT(address common.Address, backend bind.ContractBackend) (*RWACONDUIT, error) {
	contract, err := bindRWACONDUIT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RWACONDUIT{RWACONDUITCaller: RWACONDUITCaller{contract: contract}, RWACONDUITTransactor: RWACONDUITTransactor{contract: contract}, RWACONDUITFilterer: RWACONDUITFilterer{contract: contract}}, nil
}

// NewRWACONDUITCaller creates a new read-only instance of RWACONDUIT, bound to a specific deployed contract.
func NewRWACONDUITCaller(address common.Address, caller bind.ContractCaller) (*RWACONDUITCaller, error) {
	contract, err := bindRWACONDUIT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITCaller{contract: contract}, nil
}

// NewRWACONDUITTransactor creates a new write-only instance of RWACONDUIT, bound to a specific deployed contract.
func NewRWACONDUITTransactor(address common.Address, transactor bind.ContractTransactor) (*RWACONDUITTransactor, error) {
	contract, err := bindRWACONDUIT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITTransactor{contract: contract}, nil
}

// NewRWACONDUITFilterer creates a new log filterer instance of RWACONDUIT, bound to a specific deployed contract.
func NewRWACONDUITFilterer(address common.Address, filterer bind.ContractFilterer) (*RWACONDUITFilterer, error) {
	contract, err := bindRWACONDUIT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITFilterer{contract: contract}, nil
}

// bindRWACONDUIT binds a generic wrapper to an already deployed contract.
func bindRWACONDUIT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RWACONDUITABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWACONDUIT *RWACONDUITRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWACONDUIT.Contract.RWACONDUITCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWACONDUIT *RWACONDUITRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.RWACONDUITTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWACONDUIT *RWACONDUITRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.RWACONDUITTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWACONDUIT *RWACONDUITCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWACONDUIT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWACONDUIT *RWACONDUITTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWACONDUIT *RWACONDUITTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Dai() (common.Address, error) {
	return _RWACONDUIT.Contract.Dai(&_RWACONDUIT.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Dai() (common.Address, error) {
	return _RWACONDUIT.Contract.Dai(&_RWACONDUIT.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) DaiJoin() (common.Address, error) {
	return _RWACONDUIT.Contract.DaiJoin(&_RWACONDUIT.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) DaiJoin() (common.Address, error) {
	return _RWACONDUIT.Contract.DaiJoin(&_RWACONDUIT.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) End(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "end")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) End() (common.Address, error) {
	return _RWACONDUIT.Contract.End(&_RWACONDUIT.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) End() (common.Address, error) {
	return _RWACONDUIT.Contract.End(&_RWACONDUIT.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Gem() (common.Address, error) {
	return _RWACONDUIT.Contract.Gem(&_RWACONDUIT.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Gem() (common.Address, error) {
	return _RWACONDUIT.Contract.Gem(&_RWACONDUIT.CallOpts)
}

// Glad is a free data retrieval call binding the contract method 0x968e7953.
//
// Solidity: function glad() view returns(bool)
func (_RWACONDUIT *RWACONDUITCaller) Glad(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "glad")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Glad is a free data retrieval call binding the contract method 0x968e7953.
//
// Solidity: function glad() view returns(bool)
func (_RWACONDUIT *RWACONDUITSession) Glad() (bool, error) {
	return _RWACONDUIT.Contract.Glad(&_RWACONDUIT.CallOpts)
}

// Glad is a free data retrieval call binding the contract method 0x968e7953.
//
// Solidity: function glad() view returns(bool)
func (_RWACONDUIT *RWACONDUITCallerSession) Glad() (bool, error) {
	return _RWACONDUIT.Contract.Glad(&_RWACONDUIT.CallOpts)
}

// Liq is a free data retrieval call binding the contract method 0xfaba8c3b.
//
// Solidity: function liq() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Liq(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "liq")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Liq is a free data retrieval call binding the contract method 0xfaba8c3b.
//
// Solidity: function liq() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Liq() (common.Address, error) {
	return _RWACONDUIT.Contract.Liq(&_RWACONDUIT.CallOpts)
}

// Liq is a free data retrieval call binding the contract method 0xfaba8c3b.
//
// Solidity: function liq() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Liq() (common.Address, error) {
	return _RWACONDUIT.Contract.Liq(&_RWACONDUIT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_RWACONDUIT *RWACONDUITCaller) Live(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_RWACONDUIT *RWACONDUITSession) Live() (bool, error) {
	return _RWACONDUIT.Contract.Live(&_RWACONDUIT.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_RWACONDUIT *RWACONDUITCallerSession) Live() (bool, error) {
	return _RWACONDUIT.Contract.Live(&_RWACONDUIT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Owner() (common.Address, error) {
	return _RWACONDUIT.Contract.Owner(&_RWACONDUIT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Owner() (common.Address, error) {
	return _RWACONDUIT.Contract.Owner(&_RWACONDUIT.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Pool() (common.Address, error) {
	return _RWACONDUIT.Contract.Pool(&_RWACONDUIT.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Pool() (common.Address, error) {
	return _RWACONDUIT.Contract.Pool(&_RWACONDUIT.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(bool)
func (_RWACONDUIT *RWACONDUITCaller) Safe(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "safe")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(bool)
func (_RWACONDUIT *RWACONDUITSession) Safe() (bool, error) {
	return _RWACONDUIT.Contract.Safe(&_RWACONDUIT.CallOpts)
}

// Safe is a free data retrieval call binding the contract method 0x186f0354.
//
// Solidity: function safe() view returns(bool)
func (_RWACONDUIT *RWACONDUITCallerSession) Safe() (bool, error) {
	return _RWACONDUIT.Contract.Safe(&_RWACONDUIT.CallOpts)
}

// Tab is a free data retrieval call binding the contract method 0xdb9cd8d3.
//
// Solidity: function tab() view returns(uint256)
func (_RWACONDUIT *RWACONDUITCaller) Tab(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "tab")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tab is a free data retrieval call binding the contract method 0xdb9cd8d3.
//
// Solidity: function tab() view returns(uint256)
func (_RWACONDUIT *RWACONDUITSession) Tab() (*big.Int, error) {
	return _RWACONDUIT.Contract.Tab(&_RWACONDUIT.CallOpts)
}

// Tab is a free data retrieval call binding the contract method 0xdb9cd8d3.
//
// Solidity: function tab() view returns(uint256)
func (_RWACONDUIT *RWACONDUITCallerSession) Tab() (*big.Int, error) {
	return _RWACONDUIT.Contract.Tab(&_RWACONDUIT.CallOpts)
}

// Tranche is a free data retrieval call binding the contract method 0x6ebc0af1.
//
// Solidity: function tranche() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Tranche(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "tranche")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tranche is a free data retrieval call binding the contract method 0x6ebc0af1.
//
// Solidity: function tranche() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Tranche() (common.Address, error) {
	return _RWACONDUIT.Contract.Tranche(&_RWACONDUIT.CallOpts)
}

// Tranche is a free data retrieval call binding the contract method 0x6ebc0af1.
//
// Solidity: function tranche() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Tranche() (common.Address, error) {
	return _RWACONDUIT.Contract.Tranche(&_RWACONDUIT.CallOpts)
}

// Urn is a free data retrieval call binding the contract method 0x2f8080de.
//
// Solidity: function urn() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Urn(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "urn")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Urn is a free data retrieval call binding the contract method 0x2f8080de.
//
// Solidity: function urn() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Urn() (common.Address, error) {
	return _RWACONDUIT.Contract.Urn(&_RWACONDUIT.CallOpts)
}

// Urn is a free data retrieval call binding the contract method 0x2f8080de.
//
// Solidity: function urn() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Urn() (common.Address, error) {
	return _RWACONDUIT.Contract.Urn(&_RWACONDUIT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Vat() (common.Address, error) {
	return _RWACONDUIT.Contract.Vat(&_RWACONDUIT.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Vat() (common.Address, error) {
	return _RWACONDUIT.Contract.Vat(&_RWACONDUIT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_RWACONDUIT *RWACONDUITCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_RWACONDUIT *RWACONDUITSession) Vow() (common.Address, error) {
	return _RWACONDUIT.Contract.Vow(&_RWACONDUIT.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_RWACONDUIT *RWACONDUITCallerSession) Vow() (common.Address, error) {
	return _RWACONDUIT.Contract.Vow(&_RWACONDUIT.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWACONDUIT *RWACONDUITCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RWACONDUIT.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWACONDUIT *RWACONDUITSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _RWACONDUIT.Contract.Wards(&_RWACONDUIT.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_RWACONDUIT *RWACONDUITCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _RWACONDUIT.Contract.Wards(&_RWACONDUIT.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_RWACONDUIT *RWACONDUITTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_RWACONDUIT *RWACONDUITSession) Cage() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Cage(&_RWACONDUIT.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Cage() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Cage(&_RWACONDUIT.TransactOpts)
}

// Cull is a paid mutator transaction binding the contract method 0x22121f58.
//
// Solidity: function cull() returns()
func (_RWACONDUIT *RWACONDUITTransactor) Cull(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "cull")
}

// Cull is a paid mutator transaction binding the contract method 0x22121f58.
//
// Solidity: function cull() returns()
func (_RWACONDUIT *RWACONDUITSession) Cull() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Cull(&_RWACONDUIT.TransactOpts)
}

// Cull is a paid mutator transaction binding the contract method 0x22121f58.
//
// Solidity: function cull() returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Cull() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Cull(&_RWACONDUIT.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWACONDUIT *RWACONDUITSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Deny(&_RWACONDUIT.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Deny(&_RWACONDUIT.TransactOpts, usr)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Draw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "draw", wad)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Draw(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Draw(&_RWACONDUIT.TransactOpts, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x3b304147.
//
// Solidity: function draw(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Draw(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Draw(&_RWACONDUIT.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Exit(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "exit", wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Exit(&_RWACONDUIT.TransactOpts, wad)
}

// Exit is a paid mutator transaction binding the contract method 0x7f8661a1.
//
// Solidity: function exit(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Exit(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Exit(&_RWACONDUIT.TransactOpts, wad)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWACONDUIT *RWACONDUITTransactor) File(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWACONDUIT *RWACONDUITSession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.File(&_RWACONDUIT.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) File(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.File(&_RWACONDUIT.TransactOpts, what, data)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Free(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "free", wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Free(&_RWACONDUIT.TransactOpts, wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Free(&_RWACONDUIT.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Join(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "join", wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Join(&_RWACONDUIT.TransactOpts, wad)
}

// Join is a paid mutator transaction binding the contract method 0x049878f3.
//
// Solidity: function join(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Join(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Join(&_RWACONDUIT.TransactOpts, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Lock(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "lock", wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Lock(&_RWACONDUIT.TransactOpts, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Lock(&_RWACONDUIT.TransactOpts, wad)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address dst) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Migrate(opts *bind.TransactOpts, dst common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "migrate", dst)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address dst) returns()
func (_RWACONDUIT *RWACONDUITSession) Migrate(dst common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Migrate(&_RWACONDUIT.TransactOpts, dst)
}

// Migrate is a paid mutator transaction binding the contract method 0xce5494bb.
//
// Solidity: function migrate(address dst) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Migrate(dst common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Migrate(&_RWACONDUIT.TransactOpts, dst)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWACONDUIT *RWACONDUITTransactor) Quit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "quit")
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWACONDUIT *RWACONDUITSession) Quit() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Quit(&_RWACONDUIT.TransactOpts)
}

// Quit is a paid mutator transaction binding the contract method 0xfc2b8cc3.
//
// Solidity: function quit() returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Quit() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Quit(&_RWACONDUIT.TransactOpts)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Recover(opts *bind.TransactOpts, endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "recover", endEpoch)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITSession) Recover(endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Recover(&_RWACONDUIT.TransactOpts, endEpoch)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Recover(endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Recover(&_RWACONDUIT.TransactOpts, endEpoch)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWACONDUIT *RWACONDUITSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Rely(&_RWACONDUIT.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Rely(&_RWACONDUIT.TransactOpts, usr)
}

// Tell is a paid mutator transaction binding the contract method 0x53d700e5.
//
// Solidity: function tell() returns()
func (_RWACONDUIT *RWACONDUITTransactor) Tell(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "tell")
}

// Tell is a paid mutator transaction binding the contract method 0x53d700e5.
//
// Solidity: function tell() returns()
func (_RWACONDUIT *RWACONDUITSession) Tell() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Tell(&_RWACONDUIT.TransactOpts)
}

// Tell is a paid mutator transaction binding the contract method 0x53d700e5.
//
// Solidity: function tell() returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Tell() (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Tell(&_RWACONDUIT.TransactOpts)
}

// Unwind is a paid mutator transaction binding the contract method 0xfad898c2.
//
// Solidity: function unwind(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Unwind(opts *bind.TransactOpts, endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "unwind", endEpoch)
}

// Unwind is a paid mutator transaction binding the contract method 0xfad898c2.
//
// Solidity: function unwind(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITSession) Unwind(endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Unwind(&_RWACONDUIT.TransactOpts, endEpoch)
}

// Unwind is a paid mutator transaction binding the contract method 0xfad898c2.
//
// Solidity: function unwind(uint256 endEpoch) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Unwind(endEpoch *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Unwind(&_RWACONDUIT.TransactOpts, endEpoch)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactor) Wipe(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.contract.Transact(opts, "wipe", wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITSession) Wipe(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Wipe(&_RWACONDUIT.TransactOpts, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xb38a1620.
//
// Solidity: function wipe(uint256 wad) returns()
func (_RWACONDUIT *RWACONDUITTransactorSession) Wipe(wad *big.Int) (*types.Transaction, error) {
	return _RWACONDUIT.Contract.Wipe(&_RWACONDUIT.TransactOpts, wad)
}

// RWACONDUITCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the RWACONDUIT contract.
type RWACONDUITCageIterator struct {
	Event *RWACONDUITCage // Event containing the contract specifics and raw log

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
func (it *RWACONDUITCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITCage)
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
		it.Event = new(RWACONDUITCage)
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
func (it *RWACONDUITCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITCage represents a Cage event raised by the RWACONDUIT contract.
type RWACONDUITCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_RWACONDUIT *RWACONDUITFilterer) FilterCage(opts *bind.FilterOpts) (*RWACONDUITCageIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITCageIterator{contract: _RWACONDUIT.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_RWACONDUIT *RWACONDUITFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *RWACONDUITCage) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITCage)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Cage", log); err != nil {
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

// ParseCage is a log parse operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_RWACONDUIT *RWACONDUITFilterer) ParseCage(log types.Log) (*RWACONDUITCage, error) {
	event := new(RWACONDUITCage)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITCullIterator is returned from FilterCull and is used to iterate over the raw logs and unpacked data for Cull events raised by the RWACONDUIT contract.
type RWACONDUITCullIterator struct {
	Event *RWACONDUITCull // Event containing the contract specifics and raw log

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
func (it *RWACONDUITCullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITCull)
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
		it.Event = new(RWACONDUITCull)
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
func (it *RWACONDUITCullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITCullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITCull represents a Cull event raised by the RWACONDUIT contract.
type RWACONDUITCull struct {
	Tab *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCull is a free log retrieval operation binding the contract event 0xf94111cbc0e0e4fd91b7b59cc1bfb186c7fd822567ee6963fa41582df5d0c8ff.
//
// Solidity: event Cull(uint256 tab)
func (_RWACONDUIT *RWACONDUITFilterer) FilterCull(opts *bind.FilterOpts) (*RWACONDUITCullIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Cull")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITCullIterator{contract: _RWACONDUIT.contract, event: "Cull", logs: logs, sub: sub}, nil
}

// WatchCull is a free log subscription operation binding the contract event 0xf94111cbc0e0e4fd91b7b59cc1bfb186c7fd822567ee6963fa41582df5d0c8ff.
//
// Solidity: event Cull(uint256 tab)
func (_RWACONDUIT *RWACONDUITFilterer) WatchCull(opts *bind.WatchOpts, sink chan<- *RWACONDUITCull) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Cull")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITCull)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Cull", log); err != nil {
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

// ParseCull is a log parse operation binding the contract event 0xf94111cbc0e0e4fd91b7b59cc1bfb186c7fd822567ee6963fa41582df5d0c8ff.
//
// Solidity: event Cull(uint256 tab)
func (_RWACONDUIT *RWACONDUITFilterer) ParseCull(log types.Log) (*RWACONDUITCull, error) {
	event := new(RWACONDUITCull)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Cull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the RWACONDUIT contract.
type RWACONDUITDenyIterator struct {
	Event *RWACONDUITDeny // Event containing the contract specifics and raw log

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
func (it *RWACONDUITDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITDeny)
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
		it.Event = new(RWACONDUITDeny)
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
func (it *RWACONDUITDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITDeny represents a Deny event raised by the RWACONDUIT contract.
type RWACONDUITDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_RWACONDUIT *RWACONDUITFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*RWACONDUITDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITDenyIterator{contract: _RWACONDUIT.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_RWACONDUIT *RWACONDUITFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *RWACONDUITDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITDeny)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_RWACONDUIT *RWACONDUITFilterer) ParseDeny(log types.Log) (*RWACONDUITDeny, error) {
	event := new(RWACONDUITDeny)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITDrawIterator is returned from FilterDraw and is used to iterate over the raw logs and unpacked data for Draw events raised by the RWACONDUIT contract.
type RWACONDUITDrawIterator struct {
	Event *RWACONDUITDraw // Event containing the contract specifics and raw log

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
func (it *RWACONDUITDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITDraw)
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
		it.Event = new(RWACONDUITDraw)
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
func (it *RWACONDUITDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITDraw represents a Draw event raised by the RWACONDUIT contract.
type RWACONDUITDraw struct {
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDraw is a free log retrieval operation binding the contract event 0x3fd10e0b772395f38ab47fc59126002084ce1e04d2b65432f5c90bbd9e233999.
//
// Solidity: event Draw(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) FilterDraw(opts *bind.FilterOpts) (*RWACONDUITDrawIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Draw")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITDrawIterator{contract: _RWACONDUIT.contract, event: "Draw", logs: logs, sub: sub}, nil
}

// WatchDraw is a free log subscription operation binding the contract event 0x3fd10e0b772395f38ab47fc59126002084ce1e04d2b65432f5c90bbd9e233999.
//
// Solidity: event Draw(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) WatchDraw(opts *bind.WatchOpts, sink chan<- *RWACONDUITDraw) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Draw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITDraw)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Draw", log); err != nil {
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

// ParseDraw is a log parse operation binding the contract event 0x3fd10e0b772395f38ab47fc59126002084ce1e04d2b65432f5c90bbd9e233999.
//
// Solidity: event Draw(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) ParseDraw(log types.Log) (*RWACONDUITDraw, error) {
	event := new(RWACONDUITDraw)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Draw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the RWACONDUIT contract.
type RWACONDUITExitIterator struct {
	Event *RWACONDUITExit // Event containing the contract specifics and raw log

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
func (it *RWACONDUITExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITExit)
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
		it.Event = new(RWACONDUITExit)
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
func (it *RWACONDUITExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITExit represents a Exit event raised by the RWACONDUIT contract.
type RWACONDUITExit struct {
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0xb20101a10c7cc8d4a9b5accf3d34c34f89d53ec195fce51620af16429526c755.
//
// Solidity: event Exit(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) FilterExit(opts *bind.FilterOpts) (*RWACONDUITExitIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITExitIterator{contract: _RWACONDUIT.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0xb20101a10c7cc8d4a9b5accf3d34c34f89d53ec195fce51620af16429526c755.
//
// Solidity: event Exit(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *RWACONDUITExit) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITExit)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0xb20101a10c7cc8d4a9b5accf3d34c34f89d53ec195fce51620af16429526c755.
//
// Solidity: event Exit(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) ParseExit(log types.Log) (*RWACONDUITExit, error) {
	event := new(RWACONDUITExit)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the RWACONDUIT contract.
type RWACONDUITFileIterator struct {
	Event *RWACONDUITFile // Event containing the contract specifics and raw log

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
func (it *RWACONDUITFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITFile)
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
		it.Event = new(RWACONDUITFile)
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
func (it *RWACONDUITFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITFile represents a File event raised by the RWACONDUIT contract.
type RWACONDUITFile struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address indexed data)
func (_RWACONDUIT *RWACONDUITFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte, data []common.Address) (*RWACONDUITFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "File", whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITFileIterator{contract: _RWACONDUIT.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address indexed data)
func (_RWACONDUIT *RWACONDUITFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *RWACONDUITFile, what [][32]byte, data []common.Address) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "File", whatRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITFile)
				if err := _RWACONDUIT.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address indexed data)
func (_RWACONDUIT *RWACONDUITFilterer) ParseFile(log types.Log) (*RWACONDUITFile, error) {
	event := new(RWACONDUITFile)
	if err := _RWACONDUIT.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITJoinIterator is returned from FilterJoin and is used to iterate over the raw logs and unpacked data for Join events raised by the RWACONDUIT contract.
type RWACONDUITJoinIterator struct {
	Event *RWACONDUITJoin // Event containing the contract specifics and raw log

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
func (it *RWACONDUITJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITJoin)
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
		it.Event = new(RWACONDUITJoin)
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
func (it *RWACONDUITJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITJoin represents a Join event raised by the RWACONDUIT contract.
type RWACONDUITJoin struct {
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJoin is a free log retrieval operation binding the contract event 0x858d2e17a8121c939a8c52f6821c748d2592cc8ecd8e6afcda3fc4c84248002f.
//
// Solidity: event Join(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) FilterJoin(opts *bind.FilterOpts) (*RWACONDUITJoinIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Join")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITJoinIterator{contract: _RWACONDUIT.contract, event: "Join", logs: logs, sub: sub}, nil
}

// WatchJoin is a free log subscription operation binding the contract event 0x858d2e17a8121c939a8c52f6821c748d2592cc8ecd8e6afcda3fc4c84248002f.
//
// Solidity: event Join(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) WatchJoin(opts *bind.WatchOpts, sink chan<- *RWACONDUITJoin) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Join")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITJoin)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Join", log); err != nil {
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

// ParseJoin is a log parse operation binding the contract event 0x858d2e17a8121c939a8c52f6821c748d2592cc8ecd8e6afcda3fc4c84248002f.
//
// Solidity: event Join(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) ParseJoin(log types.Log) (*RWACONDUITJoin, error) {
	event := new(RWACONDUITJoin)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Join", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITMigrateIterator is returned from FilterMigrate and is used to iterate over the raw logs and unpacked data for Migrate events raised by the RWACONDUIT contract.
type RWACONDUITMigrateIterator struct {
	Event *RWACONDUITMigrate // Event containing the contract specifics and raw log

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
func (it *RWACONDUITMigrateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITMigrate)
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
		it.Event = new(RWACONDUITMigrate)
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
func (it *RWACONDUITMigrateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITMigrateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITMigrate represents a Migrate event raised by the RWACONDUIT contract.
type RWACONDUITMigrate struct {
	Dst common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMigrate is a free log retrieval operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address indexed dst)
func (_RWACONDUIT *RWACONDUITFilterer) FilterMigrate(opts *bind.FilterOpts, dst []common.Address) (*RWACONDUITMigrateIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Migrate", dstRule)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITMigrateIterator{contract: _RWACONDUIT.contract, event: "Migrate", logs: logs, sub: sub}, nil
}

// WatchMigrate is a free log subscription operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address indexed dst)
func (_RWACONDUIT *RWACONDUITFilterer) WatchMigrate(opts *bind.WatchOpts, sink chan<- *RWACONDUITMigrate, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Migrate", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITMigrate)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Migrate", log); err != nil {
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

// ParseMigrate is a log parse operation binding the contract event 0xd58a618a39de682696ea37dd9a6bf9c793afa426fa1438e75c3966e3b541e45a.
//
// Solidity: event Migrate(address indexed dst)
func (_RWACONDUIT *RWACONDUITFilterer) ParseMigrate(log types.Log) (*RWACONDUITMigrate, error) {
	event := new(RWACONDUITMigrate)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Migrate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITRecoverIterator is returned from FilterRecover and is used to iterate over the raw logs and unpacked data for Recover events raised by the RWACONDUIT contract.
type RWACONDUITRecoverIterator struct {
	Event *RWACONDUITRecover // Event containing the contract specifics and raw log

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
func (it *RWACONDUITRecoverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITRecover)
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
		it.Event = new(RWACONDUITRecover)
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
func (it *RWACONDUITRecoverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITRecoverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITRecover represents a Recover event raised by the RWACONDUIT contract.
type RWACONDUITRecover struct {
	Recovered *big.Int
	PayBack   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecover is a free log retrieval operation binding the contract event 0x945a6f50c51e4c794dfbca655185828caa3416c52206ad59fb46940963c0c4a5.
//
// Solidity: event Recover(uint256 recovered, uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) FilterRecover(opts *bind.FilterOpts) (*RWACONDUITRecoverIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Recover")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITRecoverIterator{contract: _RWACONDUIT.contract, event: "Recover", logs: logs, sub: sub}, nil
}

// WatchRecover is a free log subscription operation binding the contract event 0x945a6f50c51e4c794dfbca655185828caa3416c52206ad59fb46940963c0c4a5.
//
// Solidity: event Recover(uint256 recovered, uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) WatchRecover(opts *bind.WatchOpts, sink chan<- *RWACONDUITRecover) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Recover")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITRecover)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Recover", log); err != nil {
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

// ParseRecover is a log parse operation binding the contract event 0x945a6f50c51e4c794dfbca655185828caa3416c52206ad59fb46940963c0c4a5.
//
// Solidity: event Recover(uint256 recovered, uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) ParseRecover(log types.Log) (*RWACONDUITRecover, error) {
	event := new(RWACONDUITRecover)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Recover", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the RWACONDUIT contract.
type RWACONDUITRelyIterator struct {
	Event *RWACONDUITRely // Event containing the contract specifics and raw log

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
func (it *RWACONDUITRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITRely)
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
		it.Event = new(RWACONDUITRely)
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
func (it *RWACONDUITRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITRely represents a Rely event raised by the RWACONDUIT contract.
type RWACONDUITRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_RWACONDUIT *RWACONDUITFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*RWACONDUITRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &RWACONDUITRelyIterator{contract: _RWACONDUIT.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_RWACONDUIT *RWACONDUITFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *RWACONDUITRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITRely)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_RWACONDUIT *RWACONDUITFilterer) ParseRely(log types.Log) (*RWACONDUITRely, error) {
	event := new(RWACONDUITRely)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITTellIterator is returned from FilterTell and is used to iterate over the raw logs and unpacked data for Tell events raised by the RWACONDUIT contract.
type RWACONDUITTellIterator struct {
	Event *RWACONDUITTell // Event containing the contract specifics and raw log

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
func (it *RWACONDUITTellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITTell)
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
		it.Event = new(RWACONDUITTell)
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
func (it *RWACONDUITTellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITTellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITTell represents a Tell event raised by the RWACONDUIT contract.
type RWACONDUITTell struct {
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTell is a free log retrieval operation binding the contract event 0x740179f722c3e7a82062471e1e7cc2c5bf62b0bfafb3d95ae5e97f8605876e4b.
//
// Solidity: event Tell(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) FilterTell(opts *bind.FilterOpts) (*RWACONDUITTellIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Tell")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITTellIterator{contract: _RWACONDUIT.contract, event: "Tell", logs: logs, sub: sub}, nil
}

// WatchTell is a free log subscription operation binding the contract event 0x740179f722c3e7a82062471e1e7cc2c5bf62b0bfafb3d95ae5e97f8605876e4b.
//
// Solidity: event Tell(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) WatchTell(opts *bind.WatchOpts, sink chan<- *RWACONDUITTell) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Tell")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITTell)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Tell", log); err != nil {
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

// ParseTell is a log parse operation binding the contract event 0x740179f722c3e7a82062471e1e7cc2c5bf62b0bfafb3d95ae5e97f8605876e4b.
//
// Solidity: event Tell(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) ParseTell(log types.Log) (*RWACONDUITTell, error) {
	event := new(RWACONDUITTell)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Tell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITUnwindIterator is returned from FilterUnwind and is used to iterate over the raw logs and unpacked data for Unwind events raised by the RWACONDUIT contract.
type RWACONDUITUnwindIterator struct {
	Event *RWACONDUITUnwind // Event containing the contract specifics and raw log

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
func (it *RWACONDUITUnwindIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITUnwind)
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
		it.Event = new(RWACONDUITUnwind)
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
func (it *RWACONDUITUnwindIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITUnwindIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITUnwind represents a Unwind event raised by the RWACONDUIT contract.
type RWACONDUITUnwind struct {
	PayBack *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnwind is a free log retrieval operation binding the contract event 0xb71844fcaa5d37721eafde73962d7ab2f9857c0795b063ae0ddf2233d371b698.
//
// Solidity: event Unwind(uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) FilterUnwind(opts *bind.FilterOpts) (*RWACONDUITUnwindIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Unwind")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITUnwindIterator{contract: _RWACONDUIT.contract, event: "Unwind", logs: logs, sub: sub}, nil
}

// WatchUnwind is a free log subscription operation binding the contract event 0xb71844fcaa5d37721eafde73962d7ab2f9857c0795b063ae0ddf2233d371b698.
//
// Solidity: event Unwind(uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) WatchUnwind(opts *bind.WatchOpts, sink chan<- *RWACONDUITUnwind) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Unwind")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITUnwind)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Unwind", log); err != nil {
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

// ParseUnwind is a log parse operation binding the contract event 0xb71844fcaa5d37721eafde73962d7ab2f9857c0795b063ae0ddf2233d371b698.
//
// Solidity: event Unwind(uint256 payBack)
func (_RWACONDUIT *RWACONDUITFilterer) ParseUnwind(log types.Log) (*RWACONDUITUnwind, error) {
	event := new(RWACONDUITUnwind)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Unwind", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWACONDUITWipeIterator is returned from FilterWipe and is used to iterate over the raw logs and unpacked data for Wipe events raised by the RWACONDUIT contract.
type RWACONDUITWipeIterator struct {
	Event *RWACONDUITWipe // Event containing the contract specifics and raw log

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
func (it *RWACONDUITWipeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWACONDUITWipe)
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
		it.Event = new(RWACONDUITWipe)
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
func (it *RWACONDUITWipeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWACONDUITWipeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWACONDUITWipe represents a Wipe event raised by the RWACONDUIT contract.
type RWACONDUITWipe struct {
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWipe is a free log retrieval operation binding the contract event 0xa73c45139e83076a25bae26e55528d4ccd00789363e6ac938f6750a608cf1285.
//
// Solidity: event Wipe(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) FilterWipe(opts *bind.FilterOpts) (*RWACONDUITWipeIterator, error) {

	logs, sub, err := _RWACONDUIT.contract.FilterLogs(opts, "Wipe")
	if err != nil {
		return nil, err
	}
	return &RWACONDUITWipeIterator{contract: _RWACONDUIT.contract, event: "Wipe", logs: logs, sub: sub}, nil
}

// WatchWipe is a free log subscription operation binding the contract event 0xa73c45139e83076a25bae26e55528d4ccd00789363e6ac938f6750a608cf1285.
//
// Solidity: event Wipe(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) WatchWipe(opts *bind.WatchOpts, sink chan<- *RWACONDUITWipe) (event.Subscription, error) {

	logs, sub, err := _RWACONDUIT.contract.WatchLogs(opts, "Wipe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWACONDUITWipe)
				if err := _RWACONDUIT.contract.UnpackLog(event, "Wipe", log); err != nil {
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

// ParseWipe is a log parse operation binding the contract event 0xa73c45139e83076a25bae26e55528d4ccd00789363e6ac938f6750a608cf1285.
//
// Solidity: event Wipe(uint256 wad)
func (_RWACONDUIT *RWACONDUITFilterer) ParseWipe(log types.Log) (*RWACONDUITWipe, error) {
	event := new(RWACONDUITWipe)
	if err := _RWACONDUIT.contract.UnpackLog(event, "Wipe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
