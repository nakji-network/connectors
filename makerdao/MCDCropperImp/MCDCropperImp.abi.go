// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDCropperImp

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

// MCDCROPPERIMPABI is the input ABI used to generate the binding from.
const MCDCROPPERIMPABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Hope\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"urp\",\"type\":\"address\"}],\"name\":\"NewProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Nope\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"flee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"getOrCreateProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"urp\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"onLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"crop\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"onVatFlux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"quit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDCROPPERIMP is an auto generated Go binding around an Ethereum contract.
type MCDCROPPERIMP struct {
	MCDCROPPERIMPCaller     // Read-only binding to the contract
	MCDCROPPERIMPTransactor // Write-only binding to the contract
	MCDCROPPERIMPFilterer   // Log filterer for contract events
}

// MCDCROPPERIMPCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDCROPPERIMPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERIMPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDCROPPERIMPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERIMPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDCROPPERIMPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDCROPPERIMPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDCROPPERIMPSession struct {
	Contract     *MCDCROPPERIMP    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDCROPPERIMPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDCROPPERIMPCallerSession struct {
	Contract *MCDCROPPERIMPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MCDCROPPERIMPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDCROPPERIMPTransactorSession struct {
	Contract     *MCDCROPPERIMPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MCDCROPPERIMPRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDCROPPERIMPRaw struct {
	Contract *MCDCROPPERIMP // Generic contract binding to access the raw methods on
}

// MCDCROPPERIMPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDCROPPERIMPCallerRaw struct {
	Contract *MCDCROPPERIMPCaller // Generic read-only contract binding to access the raw methods on
}

// MCDCROPPERIMPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDCROPPERIMPTransactorRaw struct {
	Contract *MCDCROPPERIMPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDCROPPERIMP creates a new instance of MCDCROPPERIMP, bound to a specific deployed contract.
func NewMCDCROPPERIMP(address common.Address, backend bind.ContractBackend) (*MCDCROPPERIMP, error) {
	contract, err := bindMCDCROPPERIMP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMP{MCDCROPPERIMPCaller: MCDCROPPERIMPCaller{contract: contract}, MCDCROPPERIMPTransactor: MCDCROPPERIMPTransactor{contract: contract}, MCDCROPPERIMPFilterer: MCDCROPPERIMPFilterer{contract: contract}}, nil
}

// NewMCDCROPPERIMPCaller creates a new read-only instance of MCDCROPPERIMP, bound to a specific deployed contract.
func NewMCDCROPPERIMPCaller(address common.Address, caller bind.ContractCaller) (*MCDCROPPERIMPCaller, error) {
	contract, err := bindMCDCROPPERIMP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPCaller{contract: contract}, nil
}

// NewMCDCROPPERIMPTransactor creates a new write-only instance of MCDCROPPERIMP, bound to a specific deployed contract.
func NewMCDCROPPERIMPTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDCROPPERIMPTransactor, error) {
	contract, err := bindMCDCROPPERIMP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPTransactor{contract: contract}, nil
}

// NewMCDCROPPERIMPFilterer creates a new log filterer instance of MCDCROPPERIMP, bound to a specific deployed contract.
func NewMCDCROPPERIMPFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDCROPPERIMPFilterer, error) {
	contract, err := bindMCDCROPPERIMP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPFilterer{contract: contract}, nil
}

// bindMCDCROPPERIMP binds a generic wrapper to an already deployed contract.
func bindMCDCROPPERIMP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDCROPPERIMPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCROPPERIMP *MCDCROPPERIMPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCROPPERIMP.Contract.MCDCROPPERIMPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCROPPERIMP *MCDCROPPERIMPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.MCDCROPPERIMPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCROPPERIMP *MCDCROPPERIMPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.MCDCROPPERIMPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDCROPPERIMP *MCDCROPPERIMPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDCROPPERIMP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.contract.Transact(opts, method, params...)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MCDCROPPERIMP *MCDCROPPERIMPCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDCROPPERIMP.contract.Call(opts, &out, "can", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MCDCROPPERIMP.Contract.Can(&_MCDCROPPERIMP.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_MCDCROPPERIMP *MCDCROPPERIMPCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MCDCROPPERIMP.Contract.Can(&_MCDCROPPERIMP.CallOpts, arg0, arg1)
}

// Proxy is a free data retrieval call binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address ) view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPCaller) Proxy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MCDCROPPERIMP.contract.Call(opts, &out, "proxy", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address ) view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Proxy(arg0 common.Address) (common.Address, error) {
	return _MCDCROPPERIMP.Contract.Proxy(&_MCDCROPPERIMP.CallOpts, arg0)
}

// Proxy is a free data retrieval call binding the contract method 0x06713c3e.
//
// Solidity: function proxy(address ) view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPCallerSession) Proxy(arg0 common.Address) (common.Address, error) {
	return _MCDCROPPERIMP.Contract.Proxy(&_MCDCROPPERIMP.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDCROPPERIMP.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Vat() (common.Address, error) {
	return _MCDCROPPERIMP.Contract.Vat(&_MCDCROPPERIMP.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDCROPPERIMP *MCDCROPPERIMPCallerSession) Vat() (common.Address, error) {
	return _MCDCROPPERIMP.Contract.Vat(&_MCDCROPPERIMP.CallOpts)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Exit(opts *bind.TransactOpts, crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "exit", crop, usr, val)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Exit(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Exit(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Exit is a paid mutator transaction binding the contract method 0x71006c09.
//
// Solidity: function exit(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Exit(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Exit(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Flee is a paid mutator transaction binding the contract method 0xf11eae5c.
//
// Solidity: function flee(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Flee(opts *bind.TransactOpts, crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "flee", crop, usr, val)
}

// Flee is a paid mutator transaction binding the contract method 0xf11eae5c.
//
// Solidity: function flee(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Flee(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Flee(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Flee is a paid mutator transaction binding the contract method 0xf11eae5c.
//
// Solidity: function flee(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Flee(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Flee(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Flux is a paid mutator transaction binding the contract method 0x8ff91f0e.
//
// Solidity: function flux(address crop, address src, address dst, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Flux(opts *bind.TransactOpts, crop common.Address, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "flux", crop, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x8ff91f0e.
//
// Solidity: function flux(address crop, address src, address dst, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Flux(crop common.Address, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Flux(&_MCDCROPPERIMP.TransactOpts, crop, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x8ff91f0e.
//
// Solidity: function flux(address crop, address src, address dst, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Flux(crop common.Address, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Flux(&_MCDCROPPERIMP.TransactOpts, crop, src, dst, wad)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 ilk, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Frob(opts *bind.TransactOpts, ilk [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "frob", ilk, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 ilk, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Frob(ilk [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Frob(&_MCDCROPPERIMP.TransactOpts, ilk, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 ilk, address u, address v, address w, int256 dink, int256 dart) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Frob(ilk [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Frob(&_MCDCROPPERIMP.TransactOpts, ilk, u, v, w, dink, dart)
}

// GetOrCreateProxy is a paid mutator transaction binding the contract method 0xe5c0704e.
//
// Solidity: function getOrCreateProxy(address usr) returns(address urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) GetOrCreateProxy(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "getOrCreateProxy", usr)
}

// GetOrCreateProxy is a paid mutator transaction binding the contract method 0xe5c0704e.
//
// Solidity: function getOrCreateProxy(address usr) returns(address urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) GetOrCreateProxy(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.GetOrCreateProxy(&_MCDCROPPERIMP.TransactOpts, usr)
}

// GetOrCreateProxy is a paid mutator transaction binding the contract method 0xe5c0704e.
//
// Solidity: function getOrCreateProxy(address usr) returns(address urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) GetOrCreateProxy(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.GetOrCreateProxy(&_MCDCROPPERIMP.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Hope(&_MCDCROPPERIMP.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Hope(&_MCDCROPPERIMP.TransactOpts, usr)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Join(opts *bind.TransactOpts, crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "join", crop, usr, val)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Join(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Join(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Join is a paid mutator transaction binding the contract method 0x9f6c3dbd.
//
// Solidity: function join(address crop, address usr, uint256 val) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Join(crop common.Address, usr common.Address, val *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Join(&_MCDCROPPERIMP.TransactOpts, crop, usr, val)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address u, address dst, uint256 rad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Move(opts *bind.TransactOpts, u common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "move", u, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address u, address dst, uint256 rad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Move(u common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Move(&_MCDCROPPERIMP.TransactOpts, u, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address u, address dst, uint256 rad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Move(u common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Move(&_MCDCROPPERIMP.TransactOpts, u, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Nope(&_MCDCROPPERIMP.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Nope(&_MCDCROPPERIMP.TransactOpts, usr)
}

// OnLiquidation is a paid mutator transaction binding the contract method 0xbc1e9539.
//
// Solidity: function onLiquidation(address crop, address usr, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) OnLiquidation(opts *bind.TransactOpts, crop common.Address, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "onLiquidation", crop, usr, wad)
}

// OnLiquidation is a paid mutator transaction binding the contract method 0xbc1e9539.
//
// Solidity: function onLiquidation(address crop, address usr, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) OnLiquidation(crop common.Address, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.OnLiquidation(&_MCDCROPPERIMP.TransactOpts, crop, usr, wad)
}

// OnLiquidation is a paid mutator transaction binding the contract method 0xbc1e9539.
//
// Solidity: function onLiquidation(address crop, address usr, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) OnLiquidation(crop common.Address, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.OnLiquidation(&_MCDCROPPERIMP.TransactOpts, crop, usr, wad)
}

// OnVatFlux is a paid mutator transaction binding the contract method 0xb08ac7d3.
//
// Solidity: function onVatFlux(address crop, address from, address to, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) OnVatFlux(opts *bind.TransactOpts, crop common.Address, from common.Address, to common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "onVatFlux", crop, from, to, wad)
}

// OnVatFlux is a paid mutator transaction binding the contract method 0xb08ac7d3.
//
// Solidity: function onVatFlux(address crop, address from, address to, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) OnVatFlux(crop common.Address, from common.Address, to common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.OnVatFlux(&_MCDCROPPERIMP.TransactOpts, crop, from, to, wad)
}

// OnVatFlux is a paid mutator transaction binding the contract method 0xb08ac7d3.
//
// Solidity: function onVatFlux(address crop, address from, address to, uint256 wad) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) OnVatFlux(crop common.Address, from common.Address, to common.Address, wad *big.Int) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.OnVatFlux(&_MCDCROPPERIMP.TransactOpts, crop, from, to, wad)
}

// Quit is a paid mutator transaction binding the contract method 0x9a30d031.
//
// Solidity: function quit(bytes32 ilk, address u, address dst) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactor) Quit(opts *bind.TransactOpts, ilk [32]byte, u common.Address, dst common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.contract.Transact(opts, "quit", ilk, u, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x9a30d031.
//
// Solidity: function quit(bytes32 ilk, address u, address dst) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPSession) Quit(ilk [32]byte, u common.Address, dst common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Quit(&_MCDCROPPERIMP.TransactOpts, ilk, u, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x9a30d031.
//
// Solidity: function quit(bytes32 ilk, address u, address dst) returns()
func (_MCDCROPPERIMP *MCDCROPPERIMPTransactorSession) Quit(ilk [32]byte, u common.Address, dst common.Address) (*types.Transaction, error) {
	return _MCDCROPPERIMP.Contract.Quit(&_MCDCROPPERIMP.TransactOpts, ilk, u, dst)
}

// MCDCROPPERIMPHopeIterator is returned from FilterHope and is used to iterate over the raw logs and unpacked data for Hope events raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPHopeIterator struct {
	Event *MCDCROPPERIMPHope // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERIMPHopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERIMPHope)
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
		it.Event = new(MCDCROPPERIMPHope)
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
func (it *MCDCROPPERIMPHopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERIMPHopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERIMPHope represents a Hope event raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPHope struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHope is a free log retrieval operation binding the contract event 0xaa731fc3330498a56e191236785be109218ed38365faa8c33965e6de3b78ee4c.
//
// Solidity: event Hope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) FilterHope(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MCDCROPPERIMPHopeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.FilterLogs(opts, "Hope", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPHopeIterator{contract: _MCDCROPPERIMP.contract, event: "Hope", logs: logs, sub: sub}, nil
}

// WatchHope is a free log subscription operation binding the contract event 0xaa731fc3330498a56e191236785be109218ed38365faa8c33965e6de3b78ee4c.
//
// Solidity: event Hope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) WatchHope(opts *bind.WatchOpts, sink chan<- *MCDCROPPERIMPHope, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.WatchLogs(opts, "Hope", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERIMPHope)
				if err := _MCDCROPPERIMP.contract.UnpackLog(event, "Hope", log); err != nil {
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

// ParseHope is a log parse operation binding the contract event 0xaa731fc3330498a56e191236785be109218ed38365faa8c33965e6de3b78ee4c.
//
// Solidity: event Hope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) ParseHope(log types.Log) (*MCDCROPPERIMPHope, error) {
	event := new(MCDCROPPERIMPHope)
	if err := _MCDCROPPERIMP.contract.UnpackLog(event, "Hope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDCROPPERIMPNewProxyIterator is returned from FilterNewProxy and is used to iterate over the raw logs and unpacked data for NewProxy events raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPNewProxyIterator struct {
	Event *MCDCROPPERIMPNewProxy // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERIMPNewProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERIMPNewProxy)
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
		it.Event = new(MCDCROPPERIMPNewProxy)
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
func (it *MCDCROPPERIMPNewProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERIMPNewProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERIMPNewProxy represents a NewProxy event raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPNewProxy struct {
	Usr common.Address
	Urp common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewProxy is a free log retrieval operation binding the contract event 0xa94ca626c15b5acbc6f820b88b794bce0d7a9198f2ed3e248aa0a362e2bdd60b.
//
// Solidity: event NewProxy(address indexed usr, address indexed urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) FilterNewProxy(opts *bind.FilterOpts, usr []common.Address, urp []common.Address) (*MCDCROPPERIMPNewProxyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var urpRule []interface{}
	for _, urpItem := range urp {
		urpRule = append(urpRule, urpItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.FilterLogs(opts, "NewProxy", usrRule, urpRule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPNewProxyIterator{contract: _MCDCROPPERIMP.contract, event: "NewProxy", logs: logs, sub: sub}, nil
}

// WatchNewProxy is a free log subscription operation binding the contract event 0xa94ca626c15b5acbc6f820b88b794bce0d7a9198f2ed3e248aa0a362e2bdd60b.
//
// Solidity: event NewProxy(address indexed usr, address indexed urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) WatchNewProxy(opts *bind.WatchOpts, sink chan<- *MCDCROPPERIMPNewProxy, usr []common.Address, urp []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var urpRule []interface{}
	for _, urpItem := range urp {
		urpRule = append(urpRule, urpItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.WatchLogs(opts, "NewProxy", usrRule, urpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERIMPNewProxy)
				if err := _MCDCROPPERIMP.contract.UnpackLog(event, "NewProxy", log); err != nil {
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

// ParseNewProxy is a log parse operation binding the contract event 0xa94ca626c15b5acbc6f820b88b794bce0d7a9198f2ed3e248aa0a362e2bdd60b.
//
// Solidity: event NewProxy(address indexed usr, address indexed urp)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) ParseNewProxy(log types.Log) (*MCDCROPPERIMPNewProxy, error) {
	event := new(MCDCROPPERIMPNewProxy)
	if err := _MCDCROPPERIMP.contract.UnpackLog(event, "NewProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDCROPPERIMPNopeIterator is returned from FilterNope and is used to iterate over the raw logs and unpacked data for Nope events raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPNopeIterator struct {
	Event *MCDCROPPERIMPNope // Event containing the contract specifics and raw log

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
func (it *MCDCROPPERIMPNopeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDCROPPERIMPNope)
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
		it.Event = new(MCDCROPPERIMPNope)
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
func (it *MCDCROPPERIMPNopeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDCROPPERIMPNopeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDCROPPERIMPNope represents a Nope event raised by the MCDCROPPERIMP contract.
type MCDCROPPERIMPNope struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNope is a free log retrieval operation binding the contract event 0x181131ad57ffc99f2486240a094384037710b935bcd941b626ca2856316bb2c5.
//
// Solidity: event Nope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) FilterNope(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MCDCROPPERIMPNopeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.FilterLogs(opts, "Nope", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MCDCROPPERIMPNopeIterator{contract: _MCDCROPPERIMP.contract, event: "Nope", logs: logs, sub: sub}, nil
}

// WatchNope is a free log subscription operation binding the contract event 0x181131ad57ffc99f2486240a094384037710b935bcd941b626ca2856316bb2c5.
//
// Solidity: event Nope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) WatchNope(opts *bind.WatchOpts, sink chan<- *MCDCROPPERIMPNope, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MCDCROPPERIMP.contract.WatchLogs(opts, "Nope", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDCROPPERIMPNope)
				if err := _MCDCROPPERIMP.contract.UnpackLog(event, "Nope", log); err != nil {
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

// ParseNope is a log parse operation binding the contract event 0x181131ad57ffc99f2486240a094384037710b935bcd941b626ca2856316bb2c5.
//
// Solidity: event Nope(address indexed from, address indexed to)
func (_MCDCROPPERIMP *MCDCROPPERIMPFilterer) ParseNope(log types.Log) (*MCDCROPPERIMPNope, error) {
	event := new(MCDCROPPERIMPNope)
	if err := _MCDCROPPERIMP.contract.UnpackLog(event, "Nope", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
