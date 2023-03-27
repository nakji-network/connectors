// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDFlash

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

// MCDFLASHABI is the input ABI used to generate the binding from.
const MCDFLASHABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vow_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"VatDaiFlashLoan\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CALLBACK_SUCCESS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALLBACK_SUCCESS_VAT_DAI\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accrue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractDaiLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toll\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike_4\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIVatDaiFlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"vatDaiFlashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDFLASH is an auto generated Go binding around an Ethereum contract.
type MCDFLASH struct {
	MCDFLASHCaller     // Read-only binding to the contract
	MCDFLASHTransactor // Write-only binding to the contract
	MCDFLASHFilterer   // Log filterer for contract events
}

// MCDFLASHCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDFLASHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLASHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDFLASHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLASHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDFLASHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDFLASHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDFLASHSession struct {
	Contract     *MCDFLASH         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDFLASHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDFLASHCallerSession struct {
	Contract *MCDFLASHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MCDFLASHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDFLASHTransactorSession struct {
	Contract     *MCDFLASHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MCDFLASHRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDFLASHRaw struct {
	Contract *MCDFLASH // Generic contract binding to access the raw methods on
}

// MCDFLASHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDFLASHCallerRaw struct {
	Contract *MCDFLASHCaller // Generic read-only contract binding to access the raw methods on
}

// MCDFLASHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDFLASHTransactorRaw struct {
	Contract *MCDFLASHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDFLASH creates a new instance of MCDFLASH, bound to a specific deployed contract.
func NewMCDFLASH(address common.Address, backend bind.ContractBackend) (*MCDFLASH, error) {
	contract, err := bindMCDFLASH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDFLASH{MCDFLASHCaller: MCDFLASHCaller{contract: contract}, MCDFLASHTransactor: MCDFLASHTransactor{contract: contract}, MCDFLASHFilterer: MCDFLASHFilterer{contract: contract}}, nil
}

// NewMCDFLASHCaller creates a new read-only instance of MCDFLASH, bound to a specific deployed contract.
func NewMCDFLASHCaller(address common.Address, caller bind.ContractCaller) (*MCDFLASHCaller, error) {
	contract, err := bindMCDFLASH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHCaller{contract: contract}, nil
}

// NewMCDFLASHTransactor creates a new write-only instance of MCDFLASH, bound to a specific deployed contract.
func NewMCDFLASHTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDFLASHTransactor, error) {
	contract, err := bindMCDFLASH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHTransactor{contract: contract}, nil
}

// NewMCDFLASHFilterer creates a new log filterer instance of MCDFLASH, bound to a specific deployed contract.
func NewMCDFLASHFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDFLASHFilterer, error) {
	contract, err := bindMCDFLASH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHFilterer{contract: contract}, nil
}

// bindMCDFLASH binds a generic wrapper to an already deployed contract.
func bindMCDFLASH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDFLASHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLASH *MCDFLASHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLASH.Contract.MCDFLASHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLASH *MCDFLASHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLASH.Contract.MCDFLASHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLASH *MCDFLASHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLASH.Contract.MCDFLASHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDFLASH *MCDFLASHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDFLASH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDFLASH *MCDFLASHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLASH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDFLASH *MCDFLASHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDFLASH.Contract.contract.Transact(opts, method, params...)
}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_MCDFLASH *MCDFLASHCaller) CALLBACKSUCCESS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "CALLBACK_SUCCESS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_MCDFLASH *MCDFLASHSession) CALLBACKSUCCESS() ([32]byte, error) {
	return _MCDFLASH.Contract.CALLBACKSUCCESS(&_MCDFLASH.CallOpts)
}

// CALLBACKSUCCESS is a free data retrieval call binding the contract method 0x8237e538.
//
// Solidity: function CALLBACK_SUCCESS() view returns(bytes32)
func (_MCDFLASH *MCDFLASHCallerSession) CALLBACKSUCCESS() ([32]byte, error) {
	return _MCDFLASH.Contract.CALLBACKSUCCESS(&_MCDFLASH.CallOpts)
}

// CALLBACKSUCCESSVATDAI is a free data retrieval call binding the contract method 0x8878e8c7.
//
// Solidity: function CALLBACK_SUCCESS_VAT_DAI() view returns(bytes32)
func (_MCDFLASH *MCDFLASHCaller) CALLBACKSUCCESSVATDAI(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "CALLBACK_SUCCESS_VAT_DAI")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CALLBACKSUCCESSVATDAI is a free data retrieval call binding the contract method 0x8878e8c7.
//
// Solidity: function CALLBACK_SUCCESS_VAT_DAI() view returns(bytes32)
func (_MCDFLASH *MCDFLASHSession) CALLBACKSUCCESSVATDAI() ([32]byte, error) {
	return _MCDFLASH.Contract.CALLBACKSUCCESSVATDAI(&_MCDFLASH.CallOpts)
}

// CALLBACKSUCCESSVATDAI is a free data retrieval call binding the contract method 0x8878e8c7.
//
// Solidity: function CALLBACK_SUCCESS_VAT_DAI() view returns(bytes32)
func (_MCDFLASH *MCDFLASHCallerSession) CALLBACKSUCCESSVATDAI() ([32]byte, error) {
	return _MCDFLASH.Contract.CALLBACKSUCCESSVATDAI(&_MCDFLASH.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDFLASH *MCDFLASHCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDFLASH *MCDFLASHSession) Dai() (common.Address, error) {
	return _MCDFLASH.Contract.Dai(&_MCDFLASH.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDFLASH *MCDFLASHCallerSession) Dai() (common.Address, error) {
	return _MCDFLASH.Contract.Dai(&_MCDFLASH.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDFLASH *MCDFLASHCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDFLASH *MCDFLASHSession) DaiJoin() (common.Address, error) {
	return _MCDFLASH.Contract.DaiJoin(&_MCDFLASH.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDFLASH *MCDFLASHCallerSession) DaiJoin() (common.Address, error) {
	return _MCDFLASH.Contract.DaiJoin(&_MCDFLASH.CallOpts)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_MCDFLASH *MCDFLASHCaller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_MCDFLASH *MCDFLASHSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _MCDFLASH.Contract.FlashFee(&_MCDFLASH.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_MCDFLASH *MCDFLASHCallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _MCDFLASH.Contract.FlashFee(&_MCDFLASH.CallOpts, token, amount)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_MCDFLASH *MCDFLASHCaller) Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_MCDFLASH *MCDFLASHSession) Max() (*big.Int, error) {
	return _MCDFLASH.Contract.Max(&_MCDFLASH.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_MCDFLASH *MCDFLASHCallerSession) Max() (*big.Int, error) {
	return _MCDFLASH.Contract.Max(&_MCDFLASH.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_MCDFLASH *MCDFLASHCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_MCDFLASH *MCDFLASHSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _MCDFLASH.Contract.MaxFlashLoan(&_MCDFLASH.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_MCDFLASH *MCDFLASHCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _MCDFLASH.Contract.MaxFlashLoan(&_MCDFLASH.CallOpts, token)
}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_MCDFLASH *MCDFLASHCaller) Toll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "toll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_MCDFLASH *MCDFLASHSession) Toll() (*big.Int, error) {
	return _MCDFLASH.Contract.Toll(&_MCDFLASH.CallOpts)
}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_MCDFLASH *MCDFLASHCallerSession) Toll() (*big.Int, error) {
	return _MCDFLASH.Contract.Toll(&_MCDFLASH.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLASH *MCDFLASHCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLASH *MCDFLASHSession) Vat() (common.Address, error) {
	return _MCDFLASH.Contract.Vat(&_MCDFLASH.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDFLASH *MCDFLASHCallerSession) Vat() (common.Address, error) {
	return _MCDFLASH.Contract.Vat(&_MCDFLASH.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLASH *MCDFLASHCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLASH *MCDFLASHSession) Vow() (common.Address, error) {
	return _MCDFLASH.Contract.Vow(&_MCDFLASH.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDFLASH *MCDFLASHCallerSession) Vow() (common.Address, error) {
	return _MCDFLASH.Contract.Vow(&_MCDFLASH.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLASH *MCDFLASHCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDFLASH.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLASH *MCDFLASHSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLASH.Contract.Wards(&_MCDFLASH.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDFLASH *MCDFLASHCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDFLASH.Contract.Wards(&_MCDFLASH.CallOpts, arg0)
}

// Accrue is a paid mutator transaction binding the contract method 0xf8ba4cff.
//
// Solidity: function accrue() returns()
func (_MCDFLASH *MCDFLASHTransactor) Accrue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "accrue")
}

// Accrue is a paid mutator transaction binding the contract method 0xf8ba4cff.
//
// Solidity: function accrue() returns()
func (_MCDFLASH *MCDFLASHSession) Accrue() (*types.Transaction, error) {
	return _MCDFLASH.Contract.Accrue(&_MCDFLASH.TransactOpts)
}

// Accrue is a paid mutator transaction binding the contract method 0xf8ba4cff.
//
// Solidity: function accrue() returns()
func (_MCDFLASH *MCDFLASHTransactorSession) Accrue() (*types.Transaction, error) {
	return _MCDFLASH.Contract.Accrue(&_MCDFLASH.TransactOpts)
}

// Convert is a paid mutator transaction binding the contract method 0x91bbdcc7.
//
// Solidity: function convert() returns()
func (_MCDFLASH *MCDFLASHTransactor) Convert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "convert")
}

// Convert is a paid mutator transaction binding the contract method 0x91bbdcc7.
//
// Solidity: function convert() returns()
func (_MCDFLASH *MCDFLASHSession) Convert() (*types.Transaction, error) {
	return _MCDFLASH.Contract.Convert(&_MCDFLASH.TransactOpts)
}

// Convert is a paid mutator transaction binding the contract method 0x91bbdcc7.
//
// Solidity: function convert() returns()
func (_MCDFLASH *MCDFLASHTransactorSession) Convert() (*types.Transaction, error) {
	return _MCDFLASH.Contract.Convert(&_MCDFLASH.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLASH *MCDFLASHTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLASH *MCDFLASHSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.Contract.Deny(&_MCDFLASH.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDFLASH *MCDFLASHTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.Contract.Deny(&_MCDFLASH.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLASH *MCDFLASHTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLASH *MCDFLASHSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLASH.Contract.File(&_MCDFLASH.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDFLASH *MCDFLASHTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDFLASH.Contract.File(&_MCDFLASH.TransactOpts, what, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.Contract.FlashLoan(&_MCDFLASH.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.Contract.FlashLoan(&_MCDFLASH.TransactOpts, receiver, token, amount, data)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLASH *MCDFLASHTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLASH *MCDFLASHSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.Contract.Rely(&_MCDFLASH.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDFLASH *MCDFLASHTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDFLASH.Contract.Rely(&_MCDFLASH.TransactOpts, usr)
}

// VatDaiFlashLoan is a paid mutator transaction binding the contract method 0x3f03653f.
//
// Solidity: function vatDaiFlashLoan(address receiver, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHTransactor) VatDaiFlashLoan(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.contract.Transact(opts, "vatDaiFlashLoan", receiver, amount, data)
}

// VatDaiFlashLoan is a paid mutator transaction binding the contract method 0x3f03653f.
//
// Solidity: function vatDaiFlashLoan(address receiver, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHSession) VatDaiFlashLoan(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.Contract.VatDaiFlashLoan(&_MCDFLASH.TransactOpts, receiver, amount, data)
}

// VatDaiFlashLoan is a paid mutator transaction binding the contract method 0x3f03653f.
//
// Solidity: function vatDaiFlashLoan(address receiver, uint256 amount, bytes data) returns(bool)
func (_MCDFLASH *MCDFLASHTransactorSession) VatDaiFlashLoan(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _MCDFLASH.Contract.VatDaiFlashLoan(&_MCDFLASH.TransactOpts, receiver, amount, data)
}

// MCDFLASHDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDFLASH contract.
type MCDFLASHDenyIterator struct {
	Event *MCDFLASHDeny // Event containing the contract specifics and raw log

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
func (it *MCDFLASHDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLASHDeny)
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
		it.Event = new(MCDFLASHDeny)
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
func (it *MCDFLASHDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLASHDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLASHDeny represents a Deny event raised by the MCDFLASH contract.
type MCDFLASHDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDFLASH *MCDFLASHFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDFLASHDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDFLASH.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHDenyIterator{contract: _MCDFLASH.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDFLASH *MCDFLASHFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDFLASHDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDFLASH.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLASHDeny)
				if err := _MCDFLASH.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDFLASH *MCDFLASHFilterer) ParseDeny(log types.Log) (*MCDFLASHDeny, error) {
	event := new(MCDFLASHDeny)
	if err := _MCDFLASH.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDFLASHFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDFLASH contract.
type MCDFLASHFileIterator struct {
	Event *MCDFLASHFile // Event containing the contract specifics and raw log

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
func (it *MCDFLASHFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLASHFile)
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
		it.Event = new(MCDFLASHFile)
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
func (it *MCDFLASHFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLASHFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLASHFile represents a File event raised by the MCDFLASH contract.
type MCDFLASHFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDFLASH *MCDFLASHFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDFLASHFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDFLASH.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHFileIterator{contract: _MCDFLASH.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDFLASH *MCDFLASHFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDFLASHFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDFLASH.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLASHFile)
				if err := _MCDFLASH.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDFLASH *MCDFLASHFilterer) ParseFile(log types.Log) (*MCDFLASHFile, error) {
	event := new(MCDFLASHFile)
	if err := _MCDFLASH.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDFLASHFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the MCDFLASH contract.
type MCDFLASHFlashLoanIterator struct {
	Event *MCDFLASHFlashLoan // Event containing the contract specifics and raw log

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
func (it *MCDFLASHFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLASHFlashLoan)
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
		it.Event = new(MCDFLASHFlashLoan)
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
func (it *MCDFLASHFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLASHFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLASHFlashLoan represents a FlashLoan event raised by the MCDFLASH contract.
type MCDFLASHFlashLoan struct {
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address token, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) FilterFlashLoan(opts *bind.FilterOpts, receiver []common.Address) (*MCDFLASHFlashLoanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MCDFLASH.contract.FilterLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHFlashLoanIterator{contract: _MCDFLASH.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address token, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *MCDFLASHFlashLoan, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MCDFLASH.contract.WatchLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLASHFlashLoan)
				if err := _MCDFLASH.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed receiver, address token, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) ParseFlashLoan(log types.Log) (*MCDFLASHFlashLoan, error) {
	event := new(MCDFLASHFlashLoan)
	if err := _MCDFLASH.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDFLASHRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDFLASH contract.
type MCDFLASHRelyIterator struct {
	Event *MCDFLASHRely // Event containing the contract specifics and raw log

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
func (it *MCDFLASHRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLASHRely)
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
		it.Event = new(MCDFLASHRely)
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
func (it *MCDFLASHRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLASHRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLASHRely represents a Rely event raised by the MCDFLASH contract.
type MCDFLASHRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDFLASH *MCDFLASHFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDFLASHRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDFLASH.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHRelyIterator{contract: _MCDFLASH.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDFLASH *MCDFLASHFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDFLASHRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDFLASH.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLASHRely)
				if err := _MCDFLASH.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDFLASH *MCDFLASHFilterer) ParseRely(log types.Log) (*MCDFLASHRely, error) {
	event := new(MCDFLASHRely)
	if err := _MCDFLASH.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDFLASHVatDaiFlashLoanIterator is returned from FilterVatDaiFlashLoan and is used to iterate over the raw logs and unpacked data for VatDaiFlashLoan events raised by the MCDFLASH contract.
type MCDFLASHVatDaiFlashLoanIterator struct {
	Event *MCDFLASHVatDaiFlashLoan // Event containing the contract specifics and raw log

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
func (it *MCDFLASHVatDaiFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDFLASHVatDaiFlashLoan)
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
		it.Event = new(MCDFLASHVatDaiFlashLoan)
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
func (it *MCDFLASHVatDaiFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDFLASHVatDaiFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDFLASHVatDaiFlashLoan represents a VatDaiFlashLoan event raised by the MCDFLASH contract.
type MCDFLASHVatDaiFlashLoan struct {
	Receiver common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVatDaiFlashLoan is a free log retrieval operation binding the contract event 0xbca56acc64a74a4c131755895cf7f72fc3f9e39af64241f7ad0f77e86f41ada9.
//
// Solidity: event VatDaiFlashLoan(address indexed receiver, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) FilterVatDaiFlashLoan(opts *bind.FilterOpts, receiver []common.Address) (*MCDFLASHVatDaiFlashLoanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MCDFLASH.contract.FilterLogs(opts, "VatDaiFlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MCDFLASHVatDaiFlashLoanIterator{contract: _MCDFLASH.contract, event: "VatDaiFlashLoan", logs: logs, sub: sub}, nil
}

// WatchVatDaiFlashLoan is a free log subscription operation binding the contract event 0xbca56acc64a74a4c131755895cf7f72fc3f9e39af64241f7ad0f77e86f41ada9.
//
// Solidity: event VatDaiFlashLoan(address indexed receiver, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) WatchVatDaiFlashLoan(opts *bind.WatchOpts, sink chan<- *MCDFLASHVatDaiFlashLoan, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MCDFLASH.contract.WatchLogs(opts, "VatDaiFlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDFLASHVatDaiFlashLoan)
				if err := _MCDFLASH.contract.UnpackLog(event, "VatDaiFlashLoan", log); err != nil {
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

// ParseVatDaiFlashLoan is a log parse operation binding the contract event 0xbca56acc64a74a4c131755895cf7f72fc3f9e39af64241f7ad0f77e86f41ada9.
//
// Solidity: event VatDaiFlashLoan(address indexed receiver, uint256 amount, uint256 fee)
func (_MCDFLASH *MCDFLASHFilterer) ParseVatDaiFlashLoan(log types.Log) (*MCDFLASHVatDaiFlashLoan, error) {
	event := new(MCDFLASHVatDaiFlashLoan)
	if err := _MCDFLASH.contract.UnpackLog(event, "VatDaiFlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
