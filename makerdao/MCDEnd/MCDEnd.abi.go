// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDEnd

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

// MCDENDABI is the input ABI used to generate the binding from.
const MCDENDABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Cage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Cash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Flow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"}],\"name\":\"Free\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Pack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"name\":\"Skip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"name\":\"Snip\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Thaw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Art\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bag\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"cash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cat\",\"outputs\":[{\"internalType\":\"contractCatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fix\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"flow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"free\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"gap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"out\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pot\",\"outputs\":[{\"internalType\":\"contractPotLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"snip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spot\",\"outputs\":[{\"internalType\":\"contractSpotLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tag\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thaw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVowLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"when\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDEND is an auto generated Go binding around an Ethereum contract.
type MCDEND struct {
	MCDENDCaller     // Read-only binding to the contract
	MCDENDTransactor // Write-only binding to the contract
	MCDENDFilterer   // Log filterer for contract events
}

// MCDENDCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDENDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDENDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDENDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDENDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDENDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDENDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDENDSession struct {
	Contract     *MCDEND           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDENDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDENDCallerSession struct {
	Contract *MCDENDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDENDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDENDTransactorSession struct {
	Contract     *MCDENDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDENDRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDENDRaw struct {
	Contract *MCDEND // Generic contract binding to access the raw methods on
}

// MCDENDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDENDCallerRaw struct {
	Contract *MCDENDCaller // Generic read-only contract binding to access the raw methods on
}

// MCDENDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDENDTransactorRaw struct {
	Contract *MCDENDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDEND creates a new instance of MCDEND, bound to a specific deployed contract.
func NewMCDEND(address common.Address, backend bind.ContractBackend) (*MCDEND, error) {
	contract, err := bindMCDEND(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDEND{MCDENDCaller: MCDENDCaller{contract: contract}, MCDENDTransactor: MCDENDTransactor{contract: contract}, MCDENDFilterer: MCDENDFilterer{contract: contract}}, nil
}

// NewMCDENDCaller creates a new read-only instance of MCDEND, bound to a specific deployed contract.
func NewMCDENDCaller(address common.Address, caller bind.ContractCaller) (*MCDENDCaller, error) {
	contract, err := bindMCDEND(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDENDCaller{contract: contract}, nil
}

// NewMCDENDTransactor creates a new write-only instance of MCDEND, bound to a specific deployed contract.
func NewMCDENDTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDENDTransactor, error) {
	contract, err := bindMCDEND(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDENDTransactor{contract: contract}, nil
}

// NewMCDENDFilterer creates a new log filterer instance of MCDEND, bound to a specific deployed contract.
func NewMCDENDFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDENDFilterer, error) {
	contract, err := bindMCDEND(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDENDFilterer{contract: contract}, nil
}

// bindMCDEND binds a generic wrapper to an already deployed contract.
func bindMCDEND(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDENDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDEND *MCDENDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDEND.Contract.MCDENDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDEND *MCDENDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDEND.Contract.MCDENDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDEND *MCDENDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDEND.Contract.MCDENDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDEND *MCDENDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDEND.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDEND *MCDENDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDEND.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDEND *MCDENDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDEND.Contract.contract.Transact(opts, method, params...)
}

// Art is a free data retrieval call binding the contract method 0xe1340a3d.
//
// Solidity: function Art(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Art(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "Art", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Art is a free data retrieval call binding the contract method 0xe1340a3d.
//
// Solidity: function Art(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDSession) Art(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Art(&_MCDEND.CallOpts, arg0)
}

// Art is a free data retrieval call binding the contract method 0xe1340a3d.
//
// Solidity: function Art(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Art(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Art(&_MCDEND.CallOpts, arg0)
}

// Bag is a free data retrieval call binding the contract method 0x9255f809.
//
// Solidity: function bag(address ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Bag(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "bag", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bag is a free data retrieval call binding the contract method 0x9255f809.
//
// Solidity: function bag(address ) view returns(uint256)
func (_MCDEND *MCDENDSession) Bag(arg0 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Bag(&_MCDEND.CallOpts, arg0)
}

// Bag is a free data retrieval call binding the contract method 0x9255f809.
//
// Solidity: function bag(address ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Bag(arg0 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Bag(&_MCDEND.CallOpts, arg0)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDEND *MCDENDCaller) Cat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "cat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDEND *MCDENDSession) Cat() (common.Address, error) {
	return _MCDEND.Contract.Cat(&_MCDEND.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDEND *MCDENDCallerSession) Cat() (common.Address, error) {
	return _MCDEND.Contract.Cat(&_MCDEND.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MCDEND *MCDENDCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MCDEND *MCDENDSession) Debt() (*big.Int, error) {
	return _MCDEND.Contract.Debt(&_MCDEND.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Debt() (*big.Int, error) {
	return _MCDEND.Contract.Debt(&_MCDEND.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MCDEND *MCDENDCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MCDEND *MCDENDSession) Dog() (common.Address, error) {
	return _MCDEND.Contract.Dog(&_MCDEND.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_MCDEND *MCDENDCallerSession) Dog() (common.Address, error) {
	return _MCDEND.Contract.Dog(&_MCDEND.CallOpts)
}

// Fix is a free data retrieval call binding the contract method 0x63fad85e.
//
// Solidity: function fix(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Fix(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "fix", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fix is a free data retrieval call binding the contract method 0x63fad85e.
//
// Solidity: function fix(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDSession) Fix(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Fix(&_MCDEND.CallOpts, arg0)
}

// Fix is a free data retrieval call binding the contract method 0x63fad85e.
//
// Solidity: function fix(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Fix(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Fix(&_MCDEND.CallOpts, arg0)
}

// Gap is a free data retrieval call binding the contract method 0xe6ee62aa.
//
// Solidity: function gap(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Gap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "gap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gap is a free data retrieval call binding the contract method 0xe6ee62aa.
//
// Solidity: function gap(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDSession) Gap(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Gap(&_MCDEND.CallOpts, arg0)
}

// Gap is a free data retrieval call binding the contract method 0xe6ee62aa.
//
// Solidity: function gap(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Gap(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Gap(&_MCDEND.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDEND *MCDENDCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDEND *MCDENDSession) Live() (*big.Int, error) {
	return _MCDEND.Contract.Live(&_MCDEND.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Live() (*big.Int, error) {
	return _MCDEND.Contract.Live(&_MCDEND.CallOpts)
}

// Out is a free data retrieval call binding the contract method 0xc939ebfc.
//
// Solidity: function out(bytes32 , address ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Out(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "out", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Out is a free data retrieval call binding the contract method 0xc939ebfc.
//
// Solidity: function out(bytes32 , address ) view returns(uint256)
func (_MCDEND *MCDENDSession) Out(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Out(&_MCDEND.CallOpts, arg0, arg1)
}

// Out is a free data retrieval call binding the contract method 0xc939ebfc.
//
// Solidity: function out(bytes32 , address ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Out(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Out(&_MCDEND.CallOpts, arg0, arg1)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDEND *MCDENDCaller) Pot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "pot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDEND *MCDENDSession) Pot() (common.Address, error) {
	return _MCDEND.Contract.Pot(&_MCDEND.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDEND *MCDENDCallerSession) Pot() (common.Address, error) {
	return _MCDEND.Contract.Pot(&_MCDEND.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MCDEND *MCDENDCaller) Spot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "spot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MCDEND *MCDENDSession) Spot() (common.Address, error) {
	return _MCDEND.Contract.Spot(&_MCDEND.CallOpts)
}

// Spot is a free data retrieval call binding the contract method 0x6f265b93.
//
// Solidity: function spot() view returns(address)
func (_MCDEND *MCDENDCallerSession) Spot() (common.Address, error) {
	return _MCDEND.Contract.Spot(&_MCDEND.CallOpts)
}

// Tag is a free data retrieval call binding the contract method 0xee6447b5.
//
// Solidity: function tag(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Tag(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "tag", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tag is a free data retrieval call binding the contract method 0xee6447b5.
//
// Solidity: function tag(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDSession) Tag(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Tag(&_MCDEND.CallOpts, arg0)
}

// Tag is a free data retrieval call binding the contract method 0xee6447b5.
//
// Solidity: function tag(bytes32 ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Tag(arg0 [32]byte) (*big.Int, error) {
	return _MCDEND.Contract.Tag(&_MCDEND.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDEND *MCDENDCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDEND *MCDENDSession) Vat() (common.Address, error) {
	return _MCDEND.Contract.Vat(&_MCDEND.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDEND *MCDENDCallerSession) Vat() (common.Address, error) {
	return _MCDEND.Contract.Vat(&_MCDEND.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDEND *MCDENDCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDEND *MCDENDSession) Vow() (common.Address, error) {
	return _MCDEND.Contract.Vow(&_MCDEND.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDEND *MCDENDCallerSession) Vow() (common.Address, error) {
	return _MCDEND.Contract.Vow(&_MCDEND.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDEND *MCDENDCaller) Wait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "wait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDEND *MCDENDSession) Wait() (*big.Int, error) {
	return _MCDEND.Contract.Wait(&_MCDEND.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Wait() (*big.Int, error) {
	return _MCDEND.Contract.Wait(&_MCDEND.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDEND *MCDENDCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDEND *MCDENDSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Wards(&_MCDEND.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDEND *MCDENDCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDEND.Contract.Wards(&_MCDEND.CallOpts, arg0)
}

// When is a free data retrieval call binding the contract method 0xe2b0caef.
//
// Solidity: function when() view returns(uint256)
func (_MCDEND *MCDENDCaller) When(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDEND.contract.Call(opts, &out, "when")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// When is a free data retrieval call binding the contract method 0xe2b0caef.
//
// Solidity: function when() view returns(uint256)
func (_MCDEND *MCDENDSession) When() (*big.Int, error) {
	return _MCDEND.Contract.When(&_MCDEND.CallOpts)
}

// When is a free data retrieval call binding the contract method 0xe2b0caef.
//
// Solidity: function when() view returns(uint256)
func (_MCDEND *MCDENDCallerSession) When() (*big.Int, error) {
	return _MCDEND.Contract.When(&_MCDEND.CallOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDEND *MCDENDTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDEND *MCDENDSession) Cage() (*types.Transaction, error) {
	return _MCDEND.Contract.Cage(&_MCDEND.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_MCDEND *MCDENDTransactorSession) Cage() (*types.Transaction, error) {
	return _MCDEND.Contract.Cage(&_MCDEND.TransactOpts)
}

// Cage0 is a paid mutator transaction binding the contract method 0xe2702fdc.
//
// Solidity: function cage(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactor) Cage0(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "cage0", ilk)
}

// Cage0 is a paid mutator transaction binding the contract method 0xe2702fdc.
//
// Solidity: function cage(bytes32 ilk) returns()
func (_MCDEND *MCDENDSession) Cage0(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Cage0(&_MCDEND.TransactOpts, ilk)
}

// Cage0 is a paid mutator transaction binding the contract method 0xe2702fdc.
//
// Solidity: function cage(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactorSession) Cage0(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Cage0(&_MCDEND.TransactOpts, ilk)
}

// Cash is a paid mutator transaction binding the contract method 0xfe8507c6.
//
// Solidity: function cash(bytes32 ilk, uint256 wad) returns()
func (_MCDEND *MCDENDTransactor) Cash(opts *bind.TransactOpts, ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "cash", ilk, wad)
}

// Cash is a paid mutator transaction binding the contract method 0xfe8507c6.
//
// Solidity: function cash(bytes32 ilk, uint256 wad) returns()
func (_MCDEND *MCDENDSession) Cash(ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Cash(&_MCDEND.TransactOpts, ilk, wad)
}

// Cash is a paid mutator transaction binding the contract method 0xfe8507c6.
//
// Solidity: function cash(bytes32 ilk, uint256 wad) returns()
func (_MCDEND *MCDENDTransactorSession) Cash(ilk [32]byte, wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Cash(&_MCDEND.TransactOpts, ilk, wad)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDEND *MCDENDTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDEND *MCDENDSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Deny(&_MCDEND.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDEND *MCDENDTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Deny(&_MCDEND.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDEND *MCDENDTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDEND *MCDENDSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.File(&_MCDEND.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDEND *MCDENDTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.File(&_MCDEND.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDEND *MCDENDTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDEND *MCDENDSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.File0(&_MCDEND.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_MCDEND *MCDENDTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.File0(&_MCDEND.TransactOpts, what, data)
}

// Flow is a paid mutator transaction binding the contract method 0x4a10eaa6.
//
// Solidity: function flow(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactor) Flow(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "flow", ilk)
}

// Flow is a paid mutator transaction binding the contract method 0x4a10eaa6.
//
// Solidity: function flow(bytes32 ilk) returns()
func (_MCDEND *MCDENDSession) Flow(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Flow(&_MCDEND.TransactOpts, ilk)
}

// Flow is a paid mutator transaction binding the contract method 0x4a10eaa6.
//
// Solidity: function flow(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactorSession) Flow(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Flow(&_MCDEND.TransactOpts, ilk)
}

// Free is a paid mutator transaction binding the contract method 0xc83062c6.
//
// Solidity: function free(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactor) Free(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "free", ilk)
}

// Free is a paid mutator transaction binding the contract method 0xc83062c6.
//
// Solidity: function free(bytes32 ilk) returns()
func (_MCDEND *MCDENDSession) Free(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Free(&_MCDEND.TransactOpts, ilk)
}

// Free is a paid mutator transaction binding the contract method 0xc83062c6.
//
// Solidity: function free(bytes32 ilk) returns()
func (_MCDEND *MCDENDTransactorSession) Free(ilk [32]byte) (*types.Transaction, error) {
	return _MCDEND.Contract.Free(&_MCDEND.TransactOpts, ilk)
}

// Pack is a paid mutator transaction binding the contract method 0x6ea42555.
//
// Solidity: function pack(uint256 wad) returns()
func (_MCDEND *MCDENDTransactor) Pack(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "pack", wad)
}

// Pack is a paid mutator transaction binding the contract method 0x6ea42555.
//
// Solidity: function pack(uint256 wad) returns()
func (_MCDEND *MCDENDSession) Pack(wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Pack(&_MCDEND.TransactOpts, wad)
}

// Pack is a paid mutator transaction binding the contract method 0x6ea42555.
//
// Solidity: function pack(uint256 wad) returns()
func (_MCDEND *MCDENDTransactorSession) Pack(wad *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Pack(&_MCDEND.TransactOpts, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDEND *MCDENDTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDEND *MCDENDSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Rely(&_MCDEND.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDEND *MCDENDTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Rely(&_MCDEND.TransactOpts, usr)
}

// Skim is a paid mutator transaction binding the contract method 0x89ea45d3.
//
// Solidity: function skim(bytes32 ilk, address urn) returns()
func (_MCDEND *MCDENDTransactor) Skim(opts *bind.TransactOpts, ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "skim", ilk, urn)
}

// Skim is a paid mutator transaction binding the contract method 0x89ea45d3.
//
// Solidity: function skim(bytes32 ilk, address urn) returns()
func (_MCDEND *MCDENDSession) Skim(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Skim(&_MCDEND.TransactOpts, ilk, urn)
}

// Skim is a paid mutator transaction binding the contract method 0x89ea45d3.
//
// Solidity: function skim(bytes32 ilk, address urn) returns()
func (_MCDEND *MCDENDTransactorSession) Skim(ilk [32]byte, urn common.Address) (*types.Transaction, error) {
	return _MCDEND.Contract.Skim(&_MCDEND.TransactOpts, ilk, urn)
}

// Skip is a paid mutator transaction binding the contract method 0x503ecf06.
//
// Solidity: function skip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDTransactor) Skip(opts *bind.TransactOpts, ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "skip", ilk, id)
}

// Skip is a paid mutator transaction binding the contract method 0x503ecf06.
//
// Solidity: function skip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDSession) Skip(ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Skip(&_MCDEND.TransactOpts, ilk, id)
}

// Skip is a paid mutator transaction binding the contract method 0x503ecf06.
//
// Solidity: function skip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDTransactorSession) Skip(ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Skip(&_MCDEND.TransactOpts, ilk, id)
}

// Snip is a paid mutator transaction binding the contract method 0x38c6de40.
//
// Solidity: function snip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDTransactor) Snip(opts *bind.TransactOpts, ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "snip", ilk, id)
}

// Snip is a paid mutator transaction binding the contract method 0x38c6de40.
//
// Solidity: function snip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDSession) Snip(ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Snip(&_MCDEND.TransactOpts, ilk, id)
}

// Snip is a paid mutator transaction binding the contract method 0x38c6de40.
//
// Solidity: function snip(bytes32 ilk, uint256 id) returns()
func (_MCDEND *MCDENDTransactorSession) Snip(ilk [32]byte, id *big.Int) (*types.Transaction, error) {
	return _MCDEND.Contract.Snip(&_MCDEND.TransactOpts, ilk, id)
}

// Thaw is a paid mutator transaction binding the contract method 0x5920375c.
//
// Solidity: function thaw() returns()
func (_MCDEND *MCDENDTransactor) Thaw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDEND.contract.Transact(opts, "thaw")
}

// Thaw is a paid mutator transaction binding the contract method 0x5920375c.
//
// Solidity: function thaw() returns()
func (_MCDEND *MCDENDSession) Thaw() (*types.Transaction, error) {
	return _MCDEND.Contract.Thaw(&_MCDEND.TransactOpts)
}

// Thaw is a paid mutator transaction binding the contract method 0x5920375c.
//
// Solidity: function thaw() returns()
func (_MCDEND *MCDENDTransactorSession) Thaw() (*types.Transaction, error) {
	return _MCDEND.Contract.Thaw(&_MCDEND.TransactOpts)
}

// MCDENDCageIterator is returned from FilterCage and is used to iterate over the raw logs and unpacked data for Cage events raised by the MCDEND contract.
type MCDENDCageIterator struct {
	Event *MCDENDCage // Event containing the contract specifics and raw log

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
func (it *MCDENDCageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDCage)
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
		it.Event = new(MCDENDCage)
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
func (it *MCDENDCageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDCageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDCage represents a Cage event raised by the MCDEND contract.
type MCDENDCage struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage is a free log retrieval operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_MCDEND *MCDENDFilterer) FilterCage(opts *bind.FilterOpts) (*MCDENDCageIterator, error) {

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return &MCDENDCageIterator{contract: _MCDEND.contract, event: "Cage", logs: logs, sub: sub}, nil
}

// WatchCage is a free log subscription operation binding the contract event 0x2308ed18a14e800c39b86eb6ea43270105955ca385b603b64eca89f98ae8fbda.
//
// Solidity: event Cage()
func (_MCDEND *MCDENDFilterer) WatchCage(opts *bind.WatchOpts, sink chan<- *MCDENDCage) (event.Subscription, error) {

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Cage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDCage)
				if err := _MCDEND.contract.UnpackLog(event, "Cage", log); err != nil {
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
func (_MCDEND *MCDENDFilterer) ParseCage(log types.Log) (*MCDENDCage, error) {
	event := new(MCDENDCage)
	if err := _MCDEND.contract.UnpackLog(event, "Cage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDCage0Iterator is returned from FilterCage0 and is used to iterate over the raw logs and unpacked data for Cage0 events raised by the MCDEND contract.
type MCDENDCage0Iterator struct {
	Event *MCDENDCage0 // Event containing the contract specifics and raw log

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
func (it *MCDENDCage0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDCage0)
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
		it.Event = new(MCDENDCage0)
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
func (it *MCDENDCage0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDCage0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDCage0 represents a Cage0 event raised by the MCDEND contract.
type MCDENDCage0 struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCage0 is a free log retrieval operation binding the contract event 0x4a9efa0a0e3f548761a6924fe06ac5cb94ecdbc08b10d855bbcc04e37c4910db.
//
// Solidity: event Cage(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) FilterCage0(opts *bind.FilterOpts, ilk [][32]byte) (*MCDENDCage0Iterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Cage0", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDCage0Iterator{contract: _MCDEND.contract, event: "Cage0", logs: logs, sub: sub}, nil
}

// WatchCage0 is a free log subscription operation binding the contract event 0x4a9efa0a0e3f548761a6924fe06ac5cb94ecdbc08b10d855bbcc04e37c4910db.
//
// Solidity: event Cage(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) WatchCage0(opts *bind.WatchOpts, sink chan<- *MCDENDCage0, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Cage0", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDCage0)
				if err := _MCDEND.contract.UnpackLog(event, "Cage0", log); err != nil {
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

// ParseCage0 is a log parse operation binding the contract event 0x4a9efa0a0e3f548761a6924fe06ac5cb94ecdbc08b10d855bbcc04e37c4910db.
//
// Solidity: event Cage(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) ParseCage0(log types.Log) (*MCDENDCage0, error) {
	event := new(MCDENDCage0)
	if err := _MCDEND.contract.UnpackLog(event, "Cage0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDCashIterator is returned from FilterCash and is used to iterate over the raw logs and unpacked data for Cash events raised by the MCDEND contract.
type MCDENDCashIterator struct {
	Event *MCDENDCash // Event containing the contract specifics and raw log

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
func (it *MCDENDCashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDCash)
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
		it.Event = new(MCDENDCash)
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
func (it *MCDENDCashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDCashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDCash represents a Cash event raised by the MCDEND contract.
type MCDENDCash struct {
	Ilk [32]byte
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCash is a free log retrieval operation binding the contract event 0x888c7c01b06fd8004523e2bc9a274be1feaa9f03579ae5f568061dac078793c9.
//
// Solidity: event Cash(bytes32 indexed ilk, address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) FilterCash(opts *bind.FilterOpts, ilk [][32]byte, usr []common.Address) (*MCDENDCashIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Cash", ilkRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDCashIterator{contract: _MCDEND.contract, event: "Cash", logs: logs, sub: sub}, nil
}

// WatchCash is a free log subscription operation binding the contract event 0x888c7c01b06fd8004523e2bc9a274be1feaa9f03579ae5f568061dac078793c9.
//
// Solidity: event Cash(bytes32 indexed ilk, address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) WatchCash(opts *bind.WatchOpts, sink chan<- *MCDENDCash, ilk [][32]byte, usr []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Cash", ilkRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDCash)
				if err := _MCDEND.contract.UnpackLog(event, "Cash", log); err != nil {
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

// ParseCash is a log parse operation binding the contract event 0x888c7c01b06fd8004523e2bc9a274be1feaa9f03579ae5f568061dac078793c9.
//
// Solidity: event Cash(bytes32 indexed ilk, address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) ParseCash(log types.Log) (*MCDENDCash, error) {
	event := new(MCDENDCash)
	if err := _MCDEND.contract.UnpackLog(event, "Cash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDEND contract.
type MCDENDDenyIterator struct {
	Event *MCDENDDeny // Event containing the contract specifics and raw log

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
func (it *MCDENDDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDDeny)
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
		it.Event = new(MCDENDDeny)
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
func (it *MCDENDDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDDeny represents a Deny event raised by the MCDEND contract.
type MCDENDDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDEND *MCDENDFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDENDDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDDenyIterator{contract: _MCDEND.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDEND *MCDENDFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDENDDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDDeny)
				if err := _MCDEND.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDEND *MCDENDFilterer) ParseDeny(log types.Log) (*MCDENDDeny, error) {
	event := new(MCDENDDeny)
	if err := _MCDEND.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDEND contract.
type MCDENDFileIterator struct {
	Event *MCDENDFile // Event containing the contract specifics and raw log

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
func (it *MCDENDFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDFile)
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
		it.Event = new(MCDENDFile)
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
func (it *MCDENDFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDFile represents a File event raised by the MCDEND contract.
type MCDENDFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDEND *MCDENDFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDENDFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDFileIterator{contract: _MCDEND.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDEND *MCDENDFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDENDFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDFile)
				if err := _MCDEND.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDEND *MCDENDFilterer) ParseFile(log types.Log) (*MCDENDFile, error) {
	event := new(MCDENDFile)
	if err := _MCDEND.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the MCDEND contract.
type MCDENDFile0Iterator struct {
	Event *MCDENDFile0 // Event containing the contract specifics and raw log

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
func (it *MCDENDFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDFile0)
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
		it.Event = new(MCDENDFile0)
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
func (it *MCDENDFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDFile0 represents a File0 event raised by the MCDEND contract.
type MCDENDFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDEND *MCDENDFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*MCDENDFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDFile0Iterator{contract: _MCDEND.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDEND *MCDENDFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *MCDENDFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDFile0)
				if err := _MCDEND.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_MCDEND *MCDENDFilterer) ParseFile0(log types.Log) (*MCDENDFile0, error) {
	event := new(MCDENDFile0)
	if err := _MCDEND.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDFlowIterator is returned from FilterFlow and is used to iterate over the raw logs and unpacked data for Flow events raised by the MCDEND contract.
type MCDENDFlowIterator struct {
	Event *MCDENDFlow // Event containing the contract specifics and raw log

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
func (it *MCDENDFlowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDFlow)
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
		it.Event = new(MCDENDFlow)
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
func (it *MCDENDFlowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDFlowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDFlow represents a Flow event raised by the MCDEND contract.
type MCDENDFlow struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFlow is a free log retrieval operation binding the contract event 0x8d1d5ae676a6db1f6f14414f8a6c78941bbfb700fe3f3be6d3245f26c2f2d550.
//
// Solidity: event Flow(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) FilterFlow(opts *bind.FilterOpts, ilk [][32]byte) (*MCDENDFlowIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Flow", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDFlowIterator{contract: _MCDEND.contract, event: "Flow", logs: logs, sub: sub}, nil
}

// WatchFlow is a free log subscription operation binding the contract event 0x8d1d5ae676a6db1f6f14414f8a6c78941bbfb700fe3f3be6d3245f26c2f2d550.
//
// Solidity: event Flow(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) WatchFlow(opts *bind.WatchOpts, sink chan<- *MCDENDFlow, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Flow", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDFlow)
				if err := _MCDEND.contract.UnpackLog(event, "Flow", log); err != nil {
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

// ParseFlow is a log parse operation binding the contract event 0x8d1d5ae676a6db1f6f14414f8a6c78941bbfb700fe3f3be6d3245f26c2f2d550.
//
// Solidity: event Flow(bytes32 indexed ilk)
func (_MCDEND *MCDENDFilterer) ParseFlow(log types.Log) (*MCDENDFlow, error) {
	event := new(MCDENDFlow)
	if err := _MCDEND.contract.UnpackLog(event, "Flow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDFreeIterator is returned from FilterFree and is used to iterate over the raw logs and unpacked data for Free events raised by the MCDEND contract.
type MCDENDFreeIterator struct {
	Event *MCDENDFree // Event containing the contract specifics and raw log

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
func (it *MCDENDFreeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDFree)
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
		it.Event = new(MCDENDFree)
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
func (it *MCDENDFreeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDFreeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDFree represents a Free event raised by the MCDEND contract.
type MCDENDFree struct {
	Ilk [32]byte
	Usr common.Address
	Ink *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFree is a free log retrieval operation binding the contract event 0xf26f2b994a5e16f0960958e62541681f9e3e84d4caac2e487d25e0c75243f0d8.
//
// Solidity: event Free(bytes32 indexed ilk, address indexed usr, uint256 ink)
func (_MCDEND *MCDENDFilterer) FilterFree(opts *bind.FilterOpts, ilk [][32]byte, usr []common.Address) (*MCDENDFreeIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Free", ilkRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDFreeIterator{contract: _MCDEND.contract, event: "Free", logs: logs, sub: sub}, nil
}

// WatchFree is a free log subscription operation binding the contract event 0xf26f2b994a5e16f0960958e62541681f9e3e84d4caac2e487d25e0c75243f0d8.
//
// Solidity: event Free(bytes32 indexed ilk, address indexed usr, uint256 ink)
func (_MCDEND *MCDENDFilterer) WatchFree(opts *bind.WatchOpts, sink chan<- *MCDENDFree, ilk [][32]byte, usr []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Free", ilkRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDFree)
				if err := _MCDEND.contract.UnpackLog(event, "Free", log); err != nil {
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

// ParseFree is a log parse operation binding the contract event 0xf26f2b994a5e16f0960958e62541681f9e3e84d4caac2e487d25e0c75243f0d8.
//
// Solidity: event Free(bytes32 indexed ilk, address indexed usr, uint256 ink)
func (_MCDEND *MCDENDFilterer) ParseFree(log types.Log) (*MCDENDFree, error) {
	event := new(MCDENDFree)
	if err := _MCDEND.contract.UnpackLog(event, "Free", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDPackIterator is returned from FilterPack and is used to iterate over the raw logs and unpacked data for Pack events raised by the MCDEND contract.
type MCDENDPackIterator struct {
	Event *MCDENDPack // Event containing the contract specifics and raw log

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
func (it *MCDENDPackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDPack)
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
		it.Event = new(MCDENDPack)
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
func (it *MCDENDPackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDPackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDPack represents a Pack event raised by the MCDEND contract.
type MCDENDPack struct {
	Usr common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPack is a free log retrieval operation binding the contract event 0x47a981d8cbc0f6df64c9be4ce0a423071a088bd46c549bbd11a4d566e031fe0c.
//
// Solidity: event Pack(address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) FilterPack(opts *bind.FilterOpts, usr []common.Address) (*MCDENDPackIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Pack", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDPackIterator{contract: _MCDEND.contract, event: "Pack", logs: logs, sub: sub}, nil
}

// WatchPack is a free log subscription operation binding the contract event 0x47a981d8cbc0f6df64c9be4ce0a423071a088bd46c549bbd11a4d566e031fe0c.
//
// Solidity: event Pack(address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) WatchPack(opts *bind.WatchOpts, sink chan<- *MCDENDPack, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Pack", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDPack)
				if err := _MCDEND.contract.UnpackLog(event, "Pack", log); err != nil {
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

// ParsePack is a log parse operation binding the contract event 0x47a981d8cbc0f6df64c9be4ce0a423071a088bd46c549bbd11a4d566e031fe0c.
//
// Solidity: event Pack(address indexed usr, uint256 wad)
func (_MCDEND *MCDENDFilterer) ParsePack(log types.Log) (*MCDENDPack, error) {
	event := new(MCDENDPack)
	if err := _MCDEND.contract.UnpackLog(event, "Pack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDEND contract.
type MCDENDRelyIterator struct {
	Event *MCDENDRely // Event containing the contract specifics and raw log

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
func (it *MCDENDRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDRely)
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
		it.Event = new(MCDENDRely)
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
func (it *MCDENDRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDRely represents a Rely event raised by the MCDEND contract.
type MCDENDRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDEND *MCDENDFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDENDRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDRelyIterator{contract: _MCDEND.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDEND *MCDENDFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDENDRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDRely)
				if err := _MCDEND.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDEND *MCDENDFilterer) ParseRely(log types.Log) (*MCDENDRely, error) {
	event := new(MCDENDRely)
	if err := _MCDEND.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the MCDEND contract.
type MCDENDSkimIterator struct {
	Event *MCDENDSkim // Event containing the contract specifics and raw log

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
func (it *MCDENDSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDSkim)
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
		it.Event = new(MCDENDSkim)
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
func (it *MCDENDSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDSkim represents a Skim event raised by the MCDEND contract.
type MCDENDSkim struct {
	Ilk [32]byte
	Urn common.Address
	Wad *big.Int
	Art *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0xa05b7b56166c25efbac063da905f9ea6aa1dc5101f95b43c7a838aace979ab59.
//
// Solidity: event Skim(bytes32 indexed ilk, address indexed urn, uint256 wad, uint256 art)
func (_MCDEND *MCDENDFilterer) FilterSkim(opts *bind.FilterOpts, ilk [][32]byte, urn []common.Address) (*MCDENDSkimIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Skim", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDSkimIterator{contract: _MCDEND.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0xa05b7b56166c25efbac063da905f9ea6aa1dc5101f95b43c7a838aace979ab59.
//
// Solidity: event Skim(bytes32 indexed ilk, address indexed urn, uint256 wad, uint256 art)
func (_MCDEND *MCDENDFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *MCDENDSkim, ilk [][32]byte, urn []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var urnRule []interface{}
	for _, urnItem := range urn {
		urnRule = append(urnRule, urnItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Skim", ilkRule, urnRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDSkim)
				if err := _MCDEND.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0xa05b7b56166c25efbac063da905f9ea6aa1dc5101f95b43c7a838aace979ab59.
//
// Solidity: event Skim(bytes32 indexed ilk, address indexed urn, uint256 wad, uint256 art)
func (_MCDEND *MCDENDFilterer) ParseSkim(log types.Log) (*MCDENDSkim, error) {
	event := new(MCDENDSkim)
	if err := _MCDEND.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDSkipIterator is returned from FilterSkip and is used to iterate over the raw logs and unpacked data for Skip events raised by the MCDEND contract.
type MCDENDSkipIterator struct {
	Event *MCDENDSkip // Event containing the contract specifics and raw log

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
func (it *MCDENDSkipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDSkip)
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
		it.Event = new(MCDENDSkip)
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
func (it *MCDENDSkipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDSkipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDSkip represents a Skip event raised by the MCDEND contract.
type MCDENDSkip struct {
	Ilk [32]byte
	Id  *big.Int
	Usr common.Address
	Tab *big.Int
	Lot *big.Int
	Art *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSkip is a free log retrieval operation binding the contract event 0xbfa2310a8897203a59922debd0db38279196d8de5050df84608e2bb3e7790f69.
//
// Solidity: event Skip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) FilterSkip(opts *bind.FilterOpts, ilk [][32]byte, id []*big.Int, usr []common.Address) (*MCDENDSkipIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Skip", ilkRule, idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDSkipIterator{contract: _MCDEND.contract, event: "Skip", logs: logs, sub: sub}, nil
}

// WatchSkip is a free log subscription operation binding the contract event 0xbfa2310a8897203a59922debd0db38279196d8de5050df84608e2bb3e7790f69.
//
// Solidity: event Skip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) WatchSkip(opts *bind.WatchOpts, sink chan<- *MCDENDSkip, ilk [][32]byte, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Skip", ilkRule, idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDSkip)
				if err := _MCDEND.contract.UnpackLog(event, "Skip", log); err != nil {
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

// ParseSkip is a log parse operation binding the contract event 0xbfa2310a8897203a59922debd0db38279196d8de5050df84608e2bb3e7790f69.
//
// Solidity: event Skip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) ParseSkip(log types.Log) (*MCDENDSkip, error) {
	event := new(MCDENDSkip)
	if err := _MCDEND.contract.UnpackLog(event, "Skip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDSnipIterator is returned from FilterSnip and is used to iterate over the raw logs and unpacked data for Snip events raised by the MCDEND contract.
type MCDENDSnipIterator struct {
	Event *MCDENDSnip // Event containing the contract specifics and raw log

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
func (it *MCDENDSnipIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDSnip)
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
		it.Event = new(MCDENDSnip)
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
func (it *MCDENDSnipIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDSnipIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDSnip represents a Snip event raised by the MCDEND contract.
type MCDENDSnip struct {
	Ilk [32]byte
	Id  *big.Int
	Usr common.Address
	Tab *big.Int
	Lot *big.Int
	Art *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSnip is a free log retrieval operation binding the contract event 0xfc67e20caaffa015d51f696df8ea5c273ba269c69bdc2ec31c1334d01286eaa4.
//
// Solidity: event Snip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) FilterSnip(opts *bind.FilterOpts, ilk [][32]byte, id []*big.Int, usr []common.Address) (*MCDENDSnipIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Snip", ilkRule, idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDENDSnipIterator{contract: _MCDEND.contract, event: "Snip", logs: logs, sub: sub}, nil
}

// WatchSnip is a free log subscription operation binding the contract event 0xfc67e20caaffa015d51f696df8ea5c273ba269c69bdc2ec31c1334d01286eaa4.
//
// Solidity: event Snip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) WatchSnip(opts *bind.WatchOpts, sink chan<- *MCDENDSnip, ilk [][32]byte, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Snip", ilkRule, idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDSnip)
				if err := _MCDEND.contract.UnpackLog(event, "Snip", log); err != nil {
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

// ParseSnip is a log parse operation binding the contract event 0xfc67e20caaffa015d51f696df8ea5c273ba269c69bdc2ec31c1334d01286eaa4.
//
// Solidity: event Snip(bytes32 indexed ilk, uint256 indexed id, address indexed usr, uint256 tab, uint256 lot, uint256 art)
func (_MCDEND *MCDENDFilterer) ParseSnip(log types.Log) (*MCDENDSnip, error) {
	event := new(MCDENDSnip)
	if err := _MCDEND.contract.UnpackLog(event, "Snip", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDENDThawIterator is returned from FilterThaw and is used to iterate over the raw logs and unpacked data for Thaw events raised by the MCDEND contract.
type MCDENDThawIterator struct {
	Event *MCDENDThaw // Event containing the contract specifics and raw log

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
func (it *MCDENDThawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDENDThaw)
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
		it.Event = new(MCDENDThaw)
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
func (it *MCDENDThawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDENDThawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDENDThaw represents a Thaw event raised by the MCDEND contract.
type MCDENDThaw struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterThaw is a free log retrieval operation binding the contract event 0x4df15159e645ba7d02cadde0bc937abef5ad0134623c00de50a31750b85978b9.
//
// Solidity: event Thaw()
func (_MCDEND *MCDENDFilterer) FilterThaw(opts *bind.FilterOpts) (*MCDENDThawIterator, error) {

	logs, sub, err := _MCDEND.contract.FilterLogs(opts, "Thaw")
	if err != nil {
		return nil, err
	}
	return &MCDENDThawIterator{contract: _MCDEND.contract, event: "Thaw", logs: logs, sub: sub}, nil
}

// WatchThaw is a free log subscription operation binding the contract event 0x4df15159e645ba7d02cadde0bc937abef5ad0134623c00de50a31750b85978b9.
//
// Solidity: event Thaw()
func (_MCDEND *MCDENDFilterer) WatchThaw(opts *bind.WatchOpts, sink chan<- *MCDENDThaw) (event.Subscription, error) {

	logs, sub, err := _MCDEND.contract.WatchLogs(opts, "Thaw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDENDThaw)
				if err := _MCDEND.contract.UnpackLog(event, "Thaw", log); err != nil {
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

// ParseThaw is a log parse operation binding the contract event 0x4df15159e645ba7d02cadde0bc937abef5ad0134623c00de50a31750b85978b9.
//
// Solidity: event Thaw()
func (_MCDEND *MCDENDFilterer) ParseThaw(log types.Log) (*MCDENDThaw, error) {
	event := new(MCDENDThaw)
	if err := _MCDEND.contract.UnpackLog(event, "Thaw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
