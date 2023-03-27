// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ProxyActionsCropper

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

// PROXYACTIONSCROPPERABI is the input ABI used to generate the binding from.
const PROXYACTIONSCROPPERABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cropper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cdpRegistry_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"cdpRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"crop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cropper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiJoin_join\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exitETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"exitGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"fleeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"fleeGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"freeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"freeGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"lockETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"lockETHAndDraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"lockGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amtC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"lockGemAndDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockETHAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amtC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockGemAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"wipeAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"}],\"name\":\"wipeAllAndFreeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amtC\",\"type\":\"uint256\"}],\"name\":\"wipeAllAndFreeGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"wipeAndFreeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amtC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"wipeAndFreeGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PROXYACTIONSCROPPER is an auto generated Go binding around an Ethereum contract.
type PROXYACTIONSCROPPER struct {
	PROXYACTIONSCROPPERCaller     // Read-only binding to the contract
	PROXYACTIONSCROPPERTransactor // Write-only binding to the contract
	PROXYACTIONSCROPPERFilterer   // Log filterer for contract events
}

// PROXYACTIONSCROPPERCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYACTIONSCROPPERCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSCROPPERTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYACTIONSCROPPERTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSCROPPERFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYACTIONSCROPPERFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSCROPPERSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYACTIONSCROPPERSession struct {
	Contract     *PROXYACTIONSCROPPER // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PROXYACTIONSCROPPERCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYACTIONSCROPPERCallerSession struct {
	Contract *PROXYACTIONSCROPPERCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PROXYACTIONSCROPPERTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYACTIONSCROPPERTransactorSession struct {
	Contract     *PROXYACTIONSCROPPERTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PROXYACTIONSCROPPERRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYACTIONSCROPPERRaw struct {
	Contract *PROXYACTIONSCROPPER // Generic contract binding to access the raw methods on
}

// PROXYACTIONSCROPPERCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYACTIONSCROPPERCallerRaw struct {
	Contract *PROXYACTIONSCROPPERCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYACTIONSCROPPERTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYACTIONSCROPPERTransactorRaw struct {
	Contract *PROXYACTIONSCROPPERTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYACTIONSCROPPER creates a new instance of PROXYACTIONSCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSCROPPER(address common.Address, backend bind.ContractBackend) (*PROXYACTIONSCROPPER, error) {
	contract, err := bindPROXYACTIONSCROPPER(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSCROPPER{PROXYACTIONSCROPPERCaller: PROXYACTIONSCROPPERCaller{contract: contract}, PROXYACTIONSCROPPERTransactor: PROXYACTIONSCROPPERTransactor{contract: contract}, PROXYACTIONSCROPPERFilterer: PROXYACTIONSCROPPERFilterer{contract: contract}}, nil
}

// NewPROXYACTIONSCROPPERCaller creates a new read-only instance of PROXYACTIONSCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSCROPPERCaller(address common.Address, caller bind.ContractCaller) (*PROXYACTIONSCROPPERCaller, error) {
	contract, err := bindPROXYACTIONSCROPPER(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSCROPPERCaller{contract: contract}, nil
}

// NewPROXYACTIONSCROPPERTransactor creates a new write-only instance of PROXYACTIONSCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSCROPPERTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYACTIONSCROPPERTransactor, error) {
	contract, err := bindPROXYACTIONSCROPPER(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSCROPPERTransactor{contract: contract}, nil
}

// NewPROXYACTIONSCROPPERFilterer creates a new log filterer instance of PROXYACTIONSCROPPER, bound to a specific deployed contract.
func NewPROXYACTIONSCROPPERFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYACTIONSCROPPERFilterer, error) {
	contract, err := bindPROXYACTIONSCROPPER(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSCROPPERFilterer{contract: contract}, nil
}

// bindPROXYACTIONSCROPPER binds a generic wrapper to an already deployed contract.
func bindPROXYACTIONSCROPPER(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYACTIONSCROPPERABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSCROPPER.Contract.PROXYACTIONSCROPPERCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.PROXYACTIONSCROPPERTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.PROXYACTIONSCROPPERTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONSCROPPER.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.contract.Transact(opts, method, params...)
}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCaller) CdpRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSCROPPER.contract.Call(opts, &out, "cdpRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) CdpRegistry() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.CdpRegistry(&_PROXYACTIONSCROPPER.CallOpts)
}

// CdpRegistry is a free data retrieval call binding the contract method 0xab2a2a82.
//
// Solidity: function cdpRegistry() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCallerSession) CdpRegistry() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.CdpRegistry(&_PROXYACTIONSCROPPER.CallOpts)
}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCaller) Cropper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSCROPPER.contract.Call(opts, &out, "cropper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Cropper() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.Cropper(&_PROXYACTIONSCROPPER.CallOpts)
}

// Cropper is a free data retrieval call binding the contract method 0xcb3facbc.
//
// Solidity: function cropper() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCallerSession) Cropper() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.Cropper(&_PROXYACTIONSCROPPER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PROXYACTIONSCROPPER.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Vat() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.Vat(&_PROXYACTIONSCROPPER.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERCallerSession) Vat() (common.Address, error) {
	return _PROXYACTIONSCROPPER.Contract.Vat(&_PROXYACTIONSCROPPER.CallOpts)
}

// Crop is a paid mutator transaction binding the contract method 0x0f2bb432.
//
// Solidity: function crop(address gemJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Crop(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "crop", gemJoin, cdp)
}

// Crop is a paid mutator transaction binding the contract method 0x0f2bb432.
//
// Solidity: function crop(address gemJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Crop(gemJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Crop(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp)
}

// Crop is a paid mutator transaction binding the contract method 0x0f2bb432.
//
// Solidity: function crop(address gemJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Crop(gemJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Crop(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) DaiJoinJoin(opts *bind.TransactOpts, daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "daiJoin_join", daiJoin, u, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) DaiJoinJoin(daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.DaiJoinJoin(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, u, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address daiJoin, address u, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) DaiJoinJoin(daiJoin common.Address, u common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.DaiJoinJoin(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, u, wad)
}

// Draw is a paid mutator transaction binding the contract method 0xf4c24276.
//
// Solidity: function draw(address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Draw(opts *bind.TransactOpts, jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "draw", jug, daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0xf4c24276.
//
// Solidity: function draw(address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Draw(jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Draw(&_PROXYACTIONSCROPPER.TransactOpts, jug, daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0xf4c24276.
//
// Solidity: function draw(address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Draw(jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Draw(&_PROXYACTIONSCROPPER.TransactOpts, jug, daiJoin, cdp, wad)
}

// ExitETH is a paid mutator transaction binding the contract method 0x153a566d.
//
// Solidity: function exitETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) ExitETH(opts *bind.TransactOpts, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "exitETH", ethJoin, cdp, wad)
}

// ExitETH is a paid mutator transaction binding the contract method 0x153a566d.
//
// Solidity: function exitETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) ExitETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.ExitETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// ExitETH is a paid mutator transaction binding the contract method 0x153a566d.
//
// Solidity: function exitETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) ExitETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.ExitETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// ExitGem is a paid mutator transaction binding the contract method 0x15661141.
//
// Solidity: function exitGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) ExitGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "exitGem", gemJoin, cdp, amt)
}

// ExitGem is a paid mutator transaction binding the contract method 0x15661141.
//
// Solidity: function exitGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) ExitGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.ExitGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// ExitGem is a paid mutator transaction binding the contract method 0x15661141.
//
// Solidity: function exitGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) ExitGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.ExitGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// FleeETH is a paid mutator transaction binding the contract method 0x8838cb63.
//
// Solidity: function fleeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) FleeETH(opts *bind.TransactOpts, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "fleeETH", ethJoin, cdp, wad)
}

// FleeETH is a paid mutator transaction binding the contract method 0x8838cb63.
//
// Solidity: function fleeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) FleeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FleeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// FleeETH is a paid mutator transaction binding the contract method 0x8838cb63.
//
// Solidity: function fleeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) FleeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FleeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// FleeGem is a paid mutator transaction binding the contract method 0x726b300a.
//
// Solidity: function fleeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) FleeGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "fleeGem", gemJoin, cdp, amt)
}

// FleeGem is a paid mutator transaction binding the contract method 0x726b300a.
//
// Solidity: function fleeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) FleeGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FleeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// FleeGem is a paid mutator transaction binding the contract method 0x726b300a.
//
// Solidity: function fleeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) FleeGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FleeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) FreeETH(opts *bind.TransactOpts, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "freeETH", ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) FreeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) FreeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) FreeGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "freeGem", gemJoin, cdp, amt)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) FreeGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) FreeGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.FreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Hope(opts *bind.TransactOpts, obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "hope", obj, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Hope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Hope(&_PROXYACTIONSCROPPER.TransactOpts, obj, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Hope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Hope(&_PROXYACTIONSCROPPER.TransactOpts, obj, usr)
}

// LockETH is a paid mutator transaction binding the contract method 0x902f1d6b.
//
// Solidity: function lockETH(address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) LockETH(opts *bind.TransactOpts, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "lockETH", ethJoin, cdp)
}

// LockETH is a paid mutator transaction binding the contract method 0x902f1d6b.
//
// Solidity: function lockETH(address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) LockETH(ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp)
}

// LockETH is a paid mutator transaction binding the contract method 0x902f1d6b.
//
// Solidity: function lockETH(address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) LockETH(ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, cdp)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0xa7467184.
//
// Solidity: function lockETHAndDraw(address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) LockETHAndDraw(opts *bind.TransactOpts, jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "lockETHAndDraw", jug, ethJoin, daiJoin, cdp, wadD)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0xa7467184.
//
// Solidity: function lockETHAndDraw(address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) LockETHAndDraw(jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockETHAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, ethJoin, daiJoin, cdp, wadD)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0xa7467184.
//
// Solidity: function lockETHAndDraw(address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) LockETHAndDraw(jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockETHAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, ethJoin, daiJoin, cdp, wadD)
}

// LockGem is a paid mutator transaction binding the contract method 0x91c64a1a.
//
// Solidity: function lockGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) LockGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "lockGem", gemJoin, cdp, amt)
}

// LockGem is a paid mutator transaction binding the contract method 0x91c64a1a.
//
// Solidity: function lockGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) LockGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// LockGem is a paid mutator transaction binding the contract method 0x91c64a1a.
//
// Solidity: function lockGem(address gemJoin, uint256 cdp, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) LockGem(gemJoin common.Address, cdp *big.Int, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, cdp, amt)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0x2460e0e4.
//
// Solidity: function lockGemAndDraw(address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) LockGemAndDraw(opts *bind.TransactOpts, jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "lockGemAndDraw", jug, gemJoin, daiJoin, cdp, amtC, wadD)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0x2460e0e4.
//
// Solidity: function lockGemAndDraw(address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) LockGemAndDraw(jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockGemAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, gemJoin, daiJoin, cdp, amtC, wadD)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0x2460e0e4.
//
// Solidity: function lockGemAndDraw(address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) LockGemAndDraw(jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.LockGemAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, gemJoin, daiJoin, cdp, amtC, wadD)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Nope(opts *bind.TransactOpts, obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "nope", obj, usr)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Nope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Nope(&_PROXYACTIONSCROPPER.TransactOpts, obj, usr)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Nope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Nope(&_PROXYACTIONSCROPPER.TransactOpts, obj, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Open(opts *bind.TransactOpts, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "open", ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Open(&_PROXYACTIONSCROPPER.TransactOpts, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Open(&_PROXYACTIONSCROPPER.TransactOpts, ilk, usr)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x0205c0db.
//
// Solidity: function openLockETHAndDraw(address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) OpenLockETHAndDraw(opts *bind.TransactOpts, jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "openLockETHAndDraw", jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x0205c0db.
//
// Solidity: function openLockETHAndDraw(address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) OpenLockETHAndDraw(jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.OpenLockETHAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x0205c0db.
//
// Solidity: function openLockETHAndDraw(address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) OpenLockETHAndDraw(jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.OpenLockETHAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xff9a89c6.
//
// Solidity: function openLockGemAndDraw(address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 amtC, uint256 wadD) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) OpenLockGemAndDraw(opts *bind.TransactOpts, jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "openLockGemAndDraw", jug, gemJoin, daiJoin, ilk, amtC, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xff9a89c6.
//
// Solidity: function openLockGemAndDraw(address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 amtC, uint256 wadD) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) OpenLockGemAndDraw(jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.OpenLockGemAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, gemJoin, daiJoin, ilk, amtC, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xff9a89c6.
//
// Solidity: function openLockGemAndDraw(address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 amtC, uint256 wadD) returns(uint256 cdp)
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) OpenLockGemAndDraw(jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.OpenLockGemAndDraw(&_PROXYACTIONSCROPPER.TransactOpts, jug, gemJoin, daiJoin, ilk, amtC, wadD)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Transfer(opts *bind.TransactOpts, gem common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "transfer", gem, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Transfer(gem common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Transfer(&_PROXYACTIONSCROPPER.TransactOpts, gem, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 amt) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Transfer(gem common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Transfer(&_PROXYACTIONSCROPPER.TransactOpts, gem, dst, amt)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) Wipe(opts *bind.TransactOpts, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipe", daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) Wipe(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Wipe(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) Wipe(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.Wipe(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, cdp, wad)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) WipeAll(opts *bind.TransactOpts, daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipeAll", daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) WipeAll(daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAll(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) WipeAll(daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAll(&_PROXYACTIONSCROPPER.TransactOpts, daiJoin, cdp)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x5e46700a.
//
// Solidity: function wipeAllAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) WipeAllAndFreeETH(opts *bind.TransactOpts, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipeAllAndFreeETH", ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x5e46700a.
//
// Solidity: function wipeAllAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) WipeAllAndFreeETH(ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAllAndFreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x5e46700a.
//
// Solidity: function wipeAllAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) WipeAllAndFreeETH(ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAllAndFreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xa271d9a0.
//
// Solidity: function wipeAllAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) WipeAllAndFreeGem(opts *bind.TransactOpts, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipeAllAndFreeGem", gemJoin, daiJoin, cdp, amtC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xa271d9a0.
//
// Solidity: function wipeAllAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) WipeAllAndFreeGem(gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAllAndFreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, daiJoin, cdp, amtC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xa271d9a0.
//
// Solidity: function wipeAllAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) WipeAllAndFreeGem(gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAllAndFreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, daiJoin, cdp, amtC)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0x362854e1.
//
// Solidity: function wipeAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) WipeAndFreeETH(opts *bind.TransactOpts, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipeAndFreeETH", ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0x362854e1.
//
// Solidity: function wipeAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) WipeAndFreeETH(ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAndFreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0x362854e1.
//
// Solidity: function wipeAndFreeETH(address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) WipeAndFreeETH(ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAndFreeETH(&_PROXYACTIONSCROPPER.TransactOpts, ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0x86dcc75f.
//
// Solidity: function wipeAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactor) WipeAndFreeGem(opts *bind.TransactOpts, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.contract.Transact(opts, "wipeAndFreeGem", gemJoin, daiJoin, cdp, amtC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0x86dcc75f.
//
// Solidity: function wipeAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERSession) WipeAndFreeGem(gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAndFreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, daiJoin, cdp, amtC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0x86dcc75f.
//
// Solidity: function wipeAndFreeGem(address gemJoin, address daiJoin, uint256 cdp, uint256 amtC, uint256 wadD) returns()
func (_PROXYACTIONSCROPPER *PROXYACTIONSCROPPERTransactorSession) WipeAndFreeGem(gemJoin common.Address, daiJoin common.Address, cdp *big.Int, amtC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONSCROPPER.Contract.WipeAndFreeGem(&_PROXYACTIONSCROPPER.TransactOpts, gemJoin, daiJoin, cdp, amtC, wadD)
}
