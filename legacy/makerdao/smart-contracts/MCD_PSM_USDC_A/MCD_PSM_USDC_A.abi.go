// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_PSM_USDC_A

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

// MCDPSMUSDCAABI is the input ABI used to generate the binding from.
const MCDPSMUSDCAABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BuyGem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SellGem\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"buyGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractDaiAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gemJoin\",\"outputs\":[{\"internalType\":\"contractAuthGemJoinAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gemAmt\",\"type\":\"uint256\"}],\"name\":\"sellGem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatAbstract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDPSMUSDCA is an auto generated Go binding around an Ethereum contract.
type MCDPSMUSDCA struct {
	MCDPSMUSDCACaller     // Read-only binding to the contract
	MCDPSMUSDCATransactor // Write-only binding to the contract
	MCDPSMUSDCAFilterer   // Log filterer for contract events
}

// MCDPSMUSDCACaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDPSMUSDCACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMUSDCATransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDPSMUSDCATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMUSDCAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDPSMUSDCAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDPSMUSDCASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDPSMUSDCASession struct {
	Contract     *MCDPSMUSDCA      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDPSMUSDCACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDPSMUSDCACallerSession struct {
	Contract *MCDPSMUSDCACaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MCDPSMUSDCATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDPSMUSDCATransactorSession struct {
	Contract     *MCDPSMUSDCATransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MCDPSMUSDCARaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDPSMUSDCARaw struct {
	Contract *MCDPSMUSDCA // Generic contract binding to access the raw methods on
}

// MCDPSMUSDCACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDPSMUSDCACallerRaw struct {
	Contract *MCDPSMUSDCACaller // Generic read-only contract binding to access the raw methods on
}

// MCDPSMUSDCATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDPSMUSDCATransactorRaw struct {
	Contract *MCDPSMUSDCATransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDPSMUSDCA creates a new instance of MCDPSMUSDCA, bound to a specific deployed contract.
func NewMCDPSMUSDCA(address common.Address, backend bind.ContractBackend) (*MCDPSMUSDCA, error) {
	contract, err := bindMCDPSMUSDCA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCA{MCDPSMUSDCACaller: MCDPSMUSDCACaller{contract: contract}, MCDPSMUSDCATransactor: MCDPSMUSDCATransactor{contract: contract}, MCDPSMUSDCAFilterer: MCDPSMUSDCAFilterer{contract: contract}}, nil
}

// NewMCDPSMUSDCACaller creates a new read-only instance of MCDPSMUSDCA, bound to a specific deployed contract.
func NewMCDPSMUSDCACaller(address common.Address, caller bind.ContractCaller) (*MCDPSMUSDCACaller, error) {
	contract, err := bindMCDPSMUSDCA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCACaller{contract: contract}, nil
}

// NewMCDPSMUSDCATransactor creates a new write-only instance of MCDPSMUSDCA, bound to a specific deployed contract.
func NewMCDPSMUSDCATransactor(address common.Address, transactor bind.ContractTransactor) (*MCDPSMUSDCATransactor, error) {
	contract, err := bindMCDPSMUSDCA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCATransactor{contract: contract}, nil
}

// NewMCDPSMUSDCAFilterer creates a new log filterer instance of MCDPSMUSDCA, bound to a specific deployed contract.
func NewMCDPSMUSDCAFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDPSMUSDCAFilterer, error) {
	contract, err := bindMCDPSMUSDCA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCAFilterer{contract: contract}, nil
}

// bindMCDPSMUSDCA binds a generic wrapper to an already deployed contract.
func bindMCDPSMUSDCA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDPSMUSDCAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMUSDCA *MCDPSMUSDCARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMUSDCA.Contract.MCDPSMUSDCACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMUSDCA *MCDPSMUSDCARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.MCDPSMUSDCATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMUSDCA *MCDPSMUSDCARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.MCDPSMUSDCATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDPSMUSDCA *MCDPSMUSDCACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDPSMUSDCA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.contract.Transact(opts, method, params...)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Dai() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Dai(&_MCDPSMUSDCA.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Dai() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Dai(&_MCDPSMUSDCA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) DaiJoin() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.DaiJoin(&_MCDPSMUSDCA.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) DaiJoin() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.DaiJoin(&_MCDPSMUSDCA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) GemJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "gemJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) GemJoin() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.GemJoin(&_MCDPSMUSDCA.CallOpts)
}

// GemJoin is a free data retrieval call binding the contract method 0x01664f66.
//
// Solidity: function gemJoin() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) GemJoin() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.GemJoin(&_MCDPSMUSDCA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Ilk() ([32]byte, error) {
	return _MCDPSMUSDCA.Contract.Ilk(&_MCDPSMUSDCA.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Ilk() ([32]byte, error) {
	return _MCDPSMUSDCA.Contract.Ilk(&_MCDPSMUSDCA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Tin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "tin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Tin() (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Tin(&_MCDPSMUSDCA.CallOpts)
}

// Tin is a free data retrieval call binding the contract method 0x568d4b6f.
//
// Solidity: function tin() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Tin() (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Tin(&_MCDPSMUSDCA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Tout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "tout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Tout() (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Tout(&_MCDPSMUSDCA.CallOpts)
}

// Tout is a free data retrieval call binding the contract method 0xfae036d5.
//
// Solidity: function tout() view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Tout() (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Tout(&_MCDPSMUSDCA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Vat() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Vat(&_MCDPSMUSDCA.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Vat() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Vat(&_MCDPSMUSDCA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Vow() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Vow(&_MCDPSMUSDCA.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Vow() (common.Address, error) {
	return _MCDPSMUSDCA.Contract.Vow(&_MCDPSMUSDCA.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDPSMUSDCA.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Wards(&_MCDPSMUSDCA.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDPSMUSDCA *MCDPSMUSDCACallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDPSMUSDCA.Contract.Wards(&_MCDPSMUSDCA.CallOpts, arg0)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) BuyGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "buyGem", usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.BuyGem(&_MCDPSMUSDCA.TransactOpts, usr, gemAmt)
}

// BuyGem is a paid mutator transaction binding the contract method 0x8d7ef9bb.
//
// Solidity: function buyGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) BuyGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.BuyGem(&_MCDPSMUSDCA.TransactOpts, usr, gemAmt)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Deny(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Deny(&_MCDPSMUSDCA.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.File(&_MCDPSMUSDCA.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.File(&_MCDPSMUSDCA.TransactOpts, what, data)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Hope(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Hope(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Nope(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Nope(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Rely(&_MCDPSMUSDCA.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.Rely(&_MCDPSMUSDCA.TransactOpts, usr)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactor) SellGem(opts *bind.TransactOpts, usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.contract.Transact(opts, "sellGem", usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCASession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.SellGem(&_MCDPSMUSDCA.TransactOpts, usr, gemAmt)
}

// SellGem is a paid mutator transaction binding the contract method 0x95991276.
//
// Solidity: function sellGem(address usr, uint256 gemAmt) returns()
func (_MCDPSMUSDCA *MCDPSMUSDCATransactorSession) SellGem(usr common.Address, gemAmt *big.Int) (*types.Transaction, error) {
	return _MCDPSMUSDCA.Contract.SellGem(&_MCDPSMUSDCA.TransactOpts, usr, gemAmt)
}

// MCDPSMUSDCABuyGemIterator is returned from FilterBuyGem and is used to iterate over the raw logs and unpacked data for BuyGem events raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCABuyGemIterator struct {
	Event *MCDPSMUSDCABuyGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMUSDCABuyGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMUSDCABuyGem)
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
		it.Event = new(MCDPSMUSDCABuyGem)
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
func (it *MCDPSMUSDCABuyGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMUSDCABuyGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMUSDCABuyGem represents a BuyGem event raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCABuyGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBuyGem is a free log retrieval operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) FilterBuyGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMUSDCABuyGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.FilterLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCABuyGemIterator{contract: _MCDPSMUSDCA.contract, event: "BuyGem", logs: logs, sub: sub}, nil
}

// WatchBuyGem is a free log subscription operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) WatchBuyGem(opts *bind.WatchOpts, sink chan<- *MCDPSMUSDCABuyGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.WatchLogs(opts, "BuyGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMUSDCABuyGem)
				if err := _MCDPSMUSDCA.contract.UnpackLog(event, "BuyGem", log); err != nil {
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

// ParseBuyGem is a log parse operation binding the contract event 0x085d06ecf4c34b237767a31c0888e121d89546a77f186f1987c6b8715e1a8caa.
//
// Solidity: event BuyGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) ParseBuyGem(log types.Log) (*MCDPSMUSDCABuyGem, error) {
	event := new(MCDPSMUSDCABuyGem)
	if err := _MCDPSMUSDCA.contract.UnpackLog(event, "BuyGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMUSDCADenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCADenyIterator struct {
	Event *MCDPSMUSDCADeny // Event containing the contract specifics and raw log

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
func (it *MCDPSMUSDCADenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMUSDCADeny)
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
		it.Event = new(MCDPSMUSDCADeny)
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
func (it *MCDPSMUSDCADenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMUSDCADenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMUSDCADeny represents a Deny event raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCADeny struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) FilterDeny(opts *bind.FilterOpts) (*MCDPSMUSDCADenyIterator, error) {

	logs, sub, err := _MCDPSMUSDCA.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCADenyIterator{contract: _MCDPSMUSDCA.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDPSMUSDCADeny) (event.Subscription, error) {

	logs, sub, err := _MCDPSMUSDCA.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMUSDCADeny)
				if err := _MCDPSMUSDCA.contract.UnpackLog(event, "Deny", log); err != nil {
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
// Solidity: event Deny(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) ParseDeny(log types.Log) (*MCDPSMUSDCADeny, error) {
	event := new(MCDPSMUSDCADeny)
	if err := _MCDPSMUSDCA.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMUSDCAFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCAFileIterator struct {
	Event *MCDPSMUSDCAFile // Event containing the contract specifics and raw log

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
func (it *MCDPSMUSDCAFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMUSDCAFile)
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
		it.Event = new(MCDPSMUSDCAFile)
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
func (it *MCDPSMUSDCAFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMUSDCAFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMUSDCAFile represents a File event raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCAFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDPSMUSDCAFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCAFileIterator{contract: _MCDPSMUSDCA.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDPSMUSDCAFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMUSDCAFile)
				if err := _MCDPSMUSDCA.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) ParseFile(log types.Log) (*MCDPSMUSDCAFile, error) {
	event := new(MCDPSMUSDCAFile)
	if err := _MCDPSMUSDCA.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMUSDCARelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCARelyIterator struct {
	Event *MCDPSMUSDCARely // Event containing the contract specifics and raw log

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
func (it *MCDPSMUSDCARelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMUSDCARely)
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
		it.Event = new(MCDPSMUSDCARely)
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
func (it *MCDPSMUSDCARelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMUSDCARelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMUSDCARely represents a Rely event raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCARely struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) FilterRely(opts *bind.FilterOpts) (*MCDPSMUSDCARelyIterator, error) {

	logs, sub, err := _MCDPSMUSDCA.contract.FilterLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCARelyIterator{contract: _MCDPSMUSDCA.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDPSMUSDCARely) (event.Subscription, error) {

	logs, sub, err := _MCDPSMUSDCA.contract.WatchLogs(opts, "Rely")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMUSDCARely)
				if err := _MCDPSMUSDCA.contract.UnpackLog(event, "Rely", log); err != nil {
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
// Solidity: event Rely(address user)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) ParseRely(log types.Log) (*MCDPSMUSDCARely, error) {
	event := new(MCDPSMUSDCARely)
	if err := _MCDPSMUSDCA.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDPSMUSDCASellGemIterator is returned from FilterSellGem and is used to iterate over the raw logs and unpacked data for SellGem events raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCASellGemIterator struct {
	Event *MCDPSMUSDCASellGem // Event containing the contract specifics and raw log

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
func (it *MCDPSMUSDCASellGemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDPSMUSDCASellGem)
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
		it.Event = new(MCDPSMUSDCASellGem)
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
func (it *MCDPSMUSDCASellGemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDPSMUSDCASellGemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDPSMUSDCASellGem represents a SellGem event raised by the MCDPSMUSDCA contract.
type MCDPSMUSDCASellGem struct {
	Owner common.Address
	Value *big.Int
	Fee   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSellGem is a free log retrieval operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) FilterSellGem(opts *bind.FilterOpts, owner []common.Address) (*MCDPSMUSDCASellGemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.FilterLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDPSMUSDCASellGemIterator{contract: _MCDPSMUSDCA.contract, event: "SellGem", logs: logs, sub: sub}, nil
}

// WatchSellGem is a free log subscription operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) WatchSellGem(opts *bind.WatchOpts, sink chan<- *MCDPSMUSDCASellGem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDPSMUSDCA.contract.WatchLogs(opts, "SellGem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDPSMUSDCASellGem)
				if err := _MCDPSMUSDCA.contract.UnpackLog(event, "SellGem", log); err != nil {
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

// ParseSellGem is a log parse operation binding the contract event 0xef75f5a47cc9a929968796ceb84f19e7541617b4577f2c228ea95200e1572081.
//
// Solidity: event SellGem(address indexed owner, uint256 value, uint256 fee)
func (_MCDPSMUSDCA *MCDPSMUSDCAFilterer) ParseSellGem(log types.Log) (*MCDPSMUSDCASellGem, error) {
	event := new(MCDPSMUSDCASellGem)
	if err := _MCDPSMUSDCA.contract.UnpackLog(event, "SellGem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
