// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_ACTIONS_END_CROPPER

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

// PROXYACTIONSENDCROPPERABI is the input ABI used to generate the binding from.
const PROXYACTIONSENDCROPPERABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cropper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cdpRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"cashETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"cashGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cdpRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cropper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiJoin_join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"freeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"freeGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"end\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PROXYACTIONSENDCROPPER is an auto generated Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPER struct {
	PROXYACTIONSENDCROPPERCaller     // Read-only binding to the contract
	PROXYACTIONSENDCROPPERTransactor // Write-only binding to the contract
	PROXYACTIONSENDCROPPERFilterer   // Log filterer for contract events
}

// PROXYACTIONSENDCROPPERCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDCROPPERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDCROPPERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYACTIONSENDCROPPERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSENDCROPPERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYACTIONSENDCROPPERSession struct {
	Contract     *PROXYACTIONSENDCROPPER // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PROXYACTIONSENDCROPPERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYACTIONSENDCROPPERCallerSession struct {
	Contract *PROXYACTIONSENDCROPPERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// PROXYACTIONSENDCROPPERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYACTIONSENDCROPPERTransactorSession struct {
	Contract     *PROXYACTIONSENDCROPPERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// PROXYACTIONSENDCROPPERRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPERRaw struct {
	Contract *PROXYACTIONSENDCROPPER // Generic contract binding to access the raw methods on
}

// PROXYACTIONSENDCROPPERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPERCallerRaw struct {
	Contract *PROXYACTIONSENDCROPPERCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYACTIONSENDCROPPERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYACTIONSENDCROPPERTransactorRaw struct {
	Contract *PROXYACTIONSENDCROPPERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYACTIONSENDCROPPER creates a new instance of PROXYACTIONSENDCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSENDCROPPER(address common.Address, backend bind.ContractBackend) (*PROXYACTIONSENDCROPPER, error) {
	contract, err := bindPROXYACTIONSENDCROPPER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDCROPPER{PROXYACTIONSENDCROPPERCaller: PROXYACTIONSENDCROPPERCaller{contract: contract}, PROXYACTIONSENDCROPPERTransactor: PROXYACTIONSENDCROPPERTransactor{contract: contract}, PROXYACTIONSENDCROPPERFilterer: PROXYACTIONSENDCROPPERFilterer{contract: contract}}, nil
}

// NewPROXYACTIONSENDCROPPERCaller creates a new read-only instance of PROXYACTIONSENDCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSENDCROPPERCaller(address common.Address, caller bind.ContractCaller) (*PROXYACTIONSENDCROPPERCaller, error) {
	contract, err := bindPROXYACTIONSENDCROPPER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDCROPPERCaller{contract: contract}, nil
}

// NewPROXYACTIONSENDCROPPERTransactor creates a new write-only instance of PROXYACTIONSENDCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSENDCROPPERTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYACTIONSENDCROPPERTransactor, error) {
	contract, err := bindPROXYACTIONSENDCROPPER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDCROPPERTransactor{contract: contract}, nil
}

// NewPROXYACTIONSENDCROPPERFilterer creates a new log filterer instance of PROXYACTIONSENDCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSENDCROPPERFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYACTIONSENDCROPPERFilterer, error) {
	contract, err := bindPROXYACTIONSENDCROPPER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSENDCROPPERFilterer{contract: contract}, nil
}

// bindPROXYACTIONSENDCROPPER binds a generic wrapper to an already deployed contract.
func bindPROXYACTIONSENDCROPPER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYACTIONSENDCROPPERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSENDCROPPER.Contract.PROXYACTIONSENDCROPPERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.PROXYACTIONSENDCROPPERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.PROXYACTIONSENDCROPPERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSENDCROPPER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.contract.Transact(opts, method, params...)
}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCaller) CdpRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSENDCROPPER.contract.Call(opts, &out, "cdpRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) CdpRegistry() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CdpRegistry(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCallerSession) CdpRegistry() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CdpRegistry(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCaller) Cropper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSENDCROPPER.contract.Call(opts, &out, "cropper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) Cropper() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Cropper(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCallerSession) Cropper() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Cropper(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSENDCROPPER.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) Vat() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Vat(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERCallerSession) Vat() (common.Address, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Vat(&_PROXYACTIONSENDCROPPER.CallOpts)
}

// CashETH is a paid mutator transaction binding the contract method 0xc7ea039a.
//
// Solidity: function cashETH(address ethJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) CashETH(opts *bind.TransactOpts, ethJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "cashETH", ethJoin, end, wad)
}

// CashETH is a paid mutator transaction binding the contract method 0xc7ea039a.
//
// Solidity: function cashETH(address ethJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) CashETH(ethJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CashETH(&_PROXYACTIONSENDCROPPER.TransactOpts, ethJoin, end, wad)
}

// CashETH is a paid mutator transaction binding the contract method 0xc7ea039a.
//
// Solidity: function cashETH(address ethJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) CashETH(ethJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CashETH(&_PROXYACTIONSENDCROPPER.TransactOpts, ethJoin, end, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x87abf42a.
//
// Solidity: function cashGem(address gemJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) CashGem(opts *bind.TransactOpts, gemJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "cashGem", gemJoin, end, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x87abf42a.
//
// Solidity: function cashGem(address gemJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) CashGem(gemJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CashGem(&_PROXYACTIONSENDCROPPER.TransactOpts, gemJoin, end, wad)
}

// CashGem is a paid mutator transaction binding the contract method 0x87abf42a.
//
// Solidity: function cashGem(address gemJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) CashGem(gemJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.CashGem(&_PROXYACTIONSENDCROPPER.TransactOpts, gemJoin, end, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) DaiJoinJoin(opts *bind.TransactOpts, daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "daiJoin_join", daiJoin, u, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) DaiJoinJoin(daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.DaiJoinJoin(&_PROXYACTIONSENDCROPPER.TransactOpts, daiJoin, u, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) DaiJoinJoin(daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.DaiJoinJoin(&_PROXYACTIONSENDCROPPER.TransactOpts, daiJoin, u, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x573d6eb3.
//
// Solidity: function freeETH(address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) FreeETH(opts *bind.TransactOpts, ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "freeETH", ethJoin, end, cdp)
}

// FreeETH is a paid mutator transaction binding the contract method 0x573d6eb3.
//
// Solidity: function freeETH(address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) FreeETH(ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.FreeETH(&_PROXYACTIONSENDCROPPER.TransactOpts, ethJoin, end, cdp)
}

// FreeETH is a paid mutator transaction binding the contract method 0x573d6eb3.
//
// Solidity: function freeETH(address ethJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) FreeETH(ethJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.FreeETH(&_PROXYACTIONSENDCROPPER.TransactOpts, ethJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0x10d5b978.
//
// Solidity: function freeGem(address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) FreeGem(opts *bind.TransactOpts, gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "freeGem", gemJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0x10d5b978.
//
// Solidity: function freeGem(address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) FreeGem(gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.FreeGem(&_PROXYACTIONSENDCROPPER.TransactOpts, gemJoin, end, cdp)
}

// FreeGem is a paid mutator transaction binding the contract method 0x10d5b978.
//
// Solidity: function freeGem(address gemJoin, address end, uint256 cdp) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) FreeGem(gemJoin common.Address, end common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.FreeGem(&_PROXYACTIONSENDCROPPER.TransactOpts, gemJoin, end, cdp)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactor) Pack(opts *bind.TransactOpts, daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.contract.Transact(opts, "pack", daiJoin, end, wad)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERSession) Pack(daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Pack(&_PROXYACTIONSENDCROPPER.TransactOpts, daiJoin, end, wad)
}

// Pack is a paid mutator transaction binding the contract method 0x33ef33d6.
//
// Solidity: function pack(address daiJoin, address end, uint256 wad) returns()
func (_PROXYACTIONSENDCROPPER *PROXYACTIONSENDCROPPERTransactorSession) Pack(daiJoin common.Address, end common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSENDCROPPER.Contract.Pack(&_PROXYACTIONSENDCROPPER.TransactOpts, daiJoin, end, wad)
}
