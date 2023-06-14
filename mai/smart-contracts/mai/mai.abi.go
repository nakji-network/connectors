// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mai

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MaiMetaData contains all meta data concerning the Mai contract.
var MaiMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenPeg\",\"type\":\"uint256\"}],\"name\":\"setTokenPeg\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stabilityPool\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethPriceSourceAddress\",\"type\":\"address\"}],\"name\":\"changeEthPriceSource\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"borrowToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setClosingFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethPriceSource\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createVault\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_treasury\",\"type\":\"uint256\"}],\"name\":\"setTreasury\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"openingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"payBackToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"}],\"name\":\"destroyVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setOpeningFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDebtCeiling\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthPriceSource\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\"}],\"name\":\"setStabilityPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultExistence\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getClosingFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOpeningFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setDebtCeiling\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"}],\"name\":\"depositCollateral\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenPriceSource\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenPeg\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultID\",\"type\":\"uint256\"}],\"name\":\"buyRiskyVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultDebt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaultCollateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"debtCeiling\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"ethPriceSourceAddress\",\"type\":\"address\"},{\"name\":\"minimumCollateralPercentage\",\"type\":\"uint256\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"vaultAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreateVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"}],\"name\":\"DestroyVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"TransferVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BorrowToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"closingFee\",\"type\":\"uint256\"}],\"name\":\"PayBackToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vaultID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountPaid\",\"type\":\"uint256\"}],\"name\":\"BuyRiskyVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// MaiABI is the input ABI used to generate the binding from.
// Deprecated: Use MaiMetaData.ABI instead.
var MaiABI = MaiMetaData.ABI

// Mai is an auto generated Go binding around an Ethereum contract.
type Mai struct {
	MaiCaller     // Read-only binding to the contract
	MaiTransactor // Write-only binding to the contract
	MaiFilterer   // Log filterer for contract events
}

// MaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type MaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MaiSession struct {
	Contract     *Mai              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MaiCallerSession struct {
	Contract *MaiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MaiTransactorSession struct {
	Contract     *MaiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type MaiRaw struct {
	Contract *Mai // Generic contract binding to access the raw methods on
}

// MaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MaiCallerRaw struct {
	Contract *MaiCaller // Generic read-only contract binding to access the raw methods on
}

// MaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MaiTransactorRaw struct {
	Contract *MaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMai creates a new instance of Mai, bound to a specific deployed contract.
func NewMai(address common.Address, backend bind.ContractBackend) (*Mai, error) {
	contract, err := bindMai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mai{MaiCaller: MaiCaller{contract: contract}, MaiTransactor: MaiTransactor{contract: contract}, MaiFilterer: MaiFilterer{contract: contract}}, nil
}

// NewMaiCaller creates a new read-only instance of Mai, bound to a specific deployed contract.
func NewMaiCaller(address common.Address, caller bind.ContractCaller) (*MaiCaller, error) {
	contract, err := bindMai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MaiCaller{contract: contract}, nil
}

// NewMaiTransactor creates a new write-only instance of Mai, bound to a specific deployed contract.
func NewMaiTransactor(address common.Address, transactor bind.ContractTransactor) (*MaiTransactor, error) {
	contract, err := bindMai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MaiTransactor{contract: contract}, nil
}

// NewMaiFilterer creates a new log filterer instance of Mai, bound to a specific deployed contract.
func NewMaiFilterer(address common.Address, filterer bind.ContractFilterer) (*MaiFilterer, error) {
	contract, err := bindMai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MaiFilterer{contract: contract}, nil
}

// bindMai binds a generic wrapper to an already deployed contract.
func bindMai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MaiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mai *MaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mai.Contract.MaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mai *MaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mai.Contract.MaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mai *MaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mai.Contract.MaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mai *MaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mai *MaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mai *MaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mai.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mai *MaiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mai *MaiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mai.Contract.Allowance(&_Mai.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Mai *MaiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Mai.Contract.Allowance(&_Mai.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mai *MaiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mai *MaiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mai.Contract.BalanceOf(&_Mai.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Mai *MaiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Mai.Contract.BalanceOf(&_Mai.CallOpts, account)
}

// ClosingFee is a free data retrieval call binding the contract method 0x1c883e7b.
//
// Solidity: function closingFee() view returns(uint256)
func (_Mai *MaiCaller) ClosingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "closingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingFee is a free data retrieval call binding the contract method 0x1c883e7b.
//
// Solidity: function closingFee() view returns(uint256)
func (_Mai *MaiSession) ClosingFee() (*big.Int, error) {
	return _Mai.Contract.ClosingFee(&_Mai.CallOpts)
}

// ClosingFee is a free data retrieval call binding the contract method 0x1c883e7b.
//
// Solidity: function closingFee() view returns(uint256)
func (_Mai *MaiCallerSession) ClosingFee() (*big.Int, error) {
	return _Mai.Contract.ClosingFee(&_Mai.CallOpts)
}

// DebtCeiling is a free data retrieval call binding the contract method 0xe1c84ea4.
//
// Solidity: function debtCeiling() view returns(uint256)
func (_Mai *MaiCaller) DebtCeiling(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "debtCeiling")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtCeiling is a free data retrieval call binding the contract method 0xe1c84ea4.
//
// Solidity: function debtCeiling() view returns(uint256)
func (_Mai *MaiSession) DebtCeiling() (*big.Int, error) {
	return _Mai.Contract.DebtCeiling(&_Mai.CallOpts)
}

// DebtCeiling is a free data retrieval call binding the contract method 0xe1c84ea4.
//
// Solidity: function debtCeiling() view returns(uint256)
func (_Mai *MaiCallerSession) DebtCeiling() (*big.Int, error) {
	return _Mai.Contract.DebtCeiling(&_Mai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mai *MaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mai *MaiSession) Decimals() (uint8, error) {
	return _Mai.Contract.Decimals(&_Mai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Mai *MaiCallerSession) Decimals() (uint8, error) {
	return _Mai.Contract.Decimals(&_Mai.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_Mai *MaiCaller) Erc721(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "erc721")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_Mai *MaiSession) Erc721() (common.Address, error) {
	return _Mai.Contract.Erc721(&_Mai.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_Mai *MaiCallerSession) Erc721() (common.Address, error) {
	return _Mai.Contract.Erc721(&_Mai.CallOpts)
}

// EthPriceSource is a free data retrieval call binding the contract method 0x42f371c6.
//
// Solidity: function ethPriceSource() view returns(address)
func (_Mai *MaiCaller) EthPriceSource(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "ethPriceSource")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPriceSource is a free data retrieval call binding the contract method 0x42f371c6.
//
// Solidity: function ethPriceSource() view returns(address)
func (_Mai *MaiSession) EthPriceSource() (common.Address, error) {
	return _Mai.Contract.EthPriceSource(&_Mai.CallOpts)
}

// EthPriceSource is a free data retrieval call binding the contract method 0x42f371c6.
//
// Solidity: function ethPriceSource() view returns(address)
func (_Mai *MaiCallerSession) EthPriceSource() (common.Address, error) {
	return _Mai.Contract.EthPriceSource(&_Mai.CallOpts)
}

// GetClosingFee is a free data retrieval call binding the contract method 0xa5e98837.
//
// Solidity: function getClosingFee() view returns(uint256)
func (_Mai *MaiCaller) GetClosingFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "getClosingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClosingFee is a free data retrieval call binding the contract method 0xa5e98837.
//
// Solidity: function getClosingFee() view returns(uint256)
func (_Mai *MaiSession) GetClosingFee() (*big.Int, error) {
	return _Mai.Contract.GetClosingFee(&_Mai.CallOpts)
}

// GetClosingFee is a free data retrieval call binding the contract method 0xa5e98837.
//
// Solidity: function getClosingFee() view returns(uint256)
func (_Mai *MaiCallerSession) GetClosingFee() (*big.Int, error) {
	return _Mai.Contract.GetClosingFee(&_Mai.CallOpts)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x94cd4ba7.
//
// Solidity: function getDebtCeiling() view returns(uint256)
func (_Mai *MaiCaller) GetDebtCeiling(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "getDebtCeiling")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x94cd4ba7.
//
// Solidity: function getDebtCeiling() view returns(uint256)
func (_Mai *MaiSession) GetDebtCeiling() (*big.Int, error) {
	return _Mai.Contract.GetDebtCeiling(&_Mai.CallOpts)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x94cd4ba7.
//
// Solidity: function getDebtCeiling() view returns(uint256)
func (_Mai *MaiCallerSession) GetDebtCeiling() (*big.Int, error) {
	return _Mai.Contract.GetDebtCeiling(&_Mai.CallOpts)
}

// GetEthPriceSource is a free data retrieval call binding the contract method 0x98c3f2db.
//
// Solidity: function getEthPriceSource() view returns(uint256)
func (_Mai *MaiCaller) GetEthPriceSource(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "getEthPriceSource")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthPriceSource is a free data retrieval call binding the contract method 0x98c3f2db.
//
// Solidity: function getEthPriceSource() view returns(uint256)
func (_Mai *MaiSession) GetEthPriceSource() (*big.Int, error) {
	return _Mai.Contract.GetEthPriceSource(&_Mai.CallOpts)
}

// GetEthPriceSource is a free data retrieval call binding the contract method 0x98c3f2db.
//
// Solidity: function getEthPriceSource() view returns(uint256)
func (_Mai *MaiCallerSession) GetEthPriceSource() (*big.Int, error) {
	return _Mai.Contract.GetEthPriceSource(&_Mai.CallOpts)
}

// GetOpeningFee is a free data retrieval call binding the contract method 0xab806f15.
//
// Solidity: function getOpeningFee() view returns(uint256)
func (_Mai *MaiCaller) GetOpeningFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "getOpeningFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOpeningFee is a free data retrieval call binding the contract method 0xab806f15.
//
// Solidity: function getOpeningFee() view returns(uint256)
func (_Mai *MaiSession) GetOpeningFee() (*big.Int, error) {
	return _Mai.Contract.GetOpeningFee(&_Mai.CallOpts)
}

// GetOpeningFee is a free data retrieval call binding the contract method 0xab806f15.
//
// Solidity: function getOpeningFee() view returns(uint256)
func (_Mai *MaiCallerSession) GetOpeningFee() (*big.Int, error) {
	return _Mai.Contract.GetOpeningFee(&_Mai.CallOpts)
}

// GetTokenPriceSource is a free data retrieval call binding the contract method 0xcd44db1b.
//
// Solidity: function getTokenPriceSource() view returns(uint256)
func (_Mai *MaiCaller) GetTokenPriceSource(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "getTokenPriceSource")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPriceSource is a free data retrieval call binding the contract method 0xcd44db1b.
//
// Solidity: function getTokenPriceSource() view returns(uint256)
func (_Mai *MaiSession) GetTokenPriceSource() (*big.Int, error) {
	return _Mai.Contract.GetTokenPriceSource(&_Mai.CallOpts)
}

// GetTokenPriceSource is a free data retrieval call binding the contract method 0xcd44db1b.
//
// Solidity: function getTokenPriceSource() view returns(uint256)
func (_Mai *MaiCallerSession) GetTokenPriceSource() (*big.Int, error) {
	return _Mai.Contract.GetTokenPriceSource(&_Mai.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Mai *MaiCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Mai *MaiSession) IsOwner() (bool, error) {
	return _Mai.Contract.IsOwner(&_Mai.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Mai *MaiCallerSession) IsOwner() (bool, error) {
	return _Mai.Contract.IsOwner(&_Mai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mai *MaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mai *MaiSession) Name() (string, error) {
	return _Mai.Contract.Name(&_Mai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mai *MaiCallerSession) Name() (string, error) {
	return _Mai.Contract.Name(&_Mai.CallOpts)
}

// OpeningFee is a free data retrieval call binding the contract method 0x728bbbb5.
//
// Solidity: function openingFee() view returns(uint256)
func (_Mai *MaiCaller) OpeningFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "openingFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningFee is a free data retrieval call binding the contract method 0x728bbbb5.
//
// Solidity: function openingFee() view returns(uint256)
func (_Mai *MaiSession) OpeningFee() (*big.Int, error) {
	return _Mai.Contract.OpeningFee(&_Mai.CallOpts)
}

// OpeningFee is a free data retrieval call binding the contract method 0x728bbbb5.
//
// Solidity: function openingFee() view returns(uint256)
func (_Mai *MaiCallerSession) OpeningFee() (*big.Int, error) {
	return _Mai.Contract.OpeningFee(&_Mai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mai *MaiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mai *MaiSession) Owner() (common.Address, error) {
	return _Mai.Contract.Owner(&_Mai.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mai *MaiCallerSession) Owner() (common.Address, error) {
	return _Mai.Contract.Owner(&_Mai.CallOpts)
}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_Mai *MaiCaller) StabilityPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "stabilityPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_Mai *MaiSession) StabilityPool() (common.Address, error) {
	return _Mai.Contract.StabilityPool(&_Mai.CallOpts)
}

// StabilityPool is a free data retrieval call binding the contract method 0x048c661d.
//
// Solidity: function stabilityPool() view returns(address)
func (_Mai *MaiCallerSession) StabilityPool() (common.Address, error) {
	return _Mai.Contract.StabilityPool(&_Mai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mai *MaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mai *MaiSession) Symbol() (string, error) {
	return _Mai.Contract.Symbol(&_Mai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mai *MaiCallerSession) Symbol() (string, error) {
	return _Mai.Contract.Symbol(&_Mai.CallOpts)
}

// TokenPeg is a free data retrieval call binding the contract method 0xcdfedd63.
//
// Solidity: function tokenPeg() view returns(uint256)
func (_Mai *MaiCaller) TokenPeg(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "tokenPeg")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPeg is a free data retrieval call binding the contract method 0xcdfedd63.
//
// Solidity: function tokenPeg() view returns(uint256)
func (_Mai *MaiSession) TokenPeg() (*big.Int, error) {
	return _Mai.Contract.TokenPeg(&_Mai.CallOpts)
}

// TokenPeg is a free data retrieval call binding the contract method 0xcdfedd63.
//
// Solidity: function tokenPeg() view returns(uint256)
func (_Mai *MaiCallerSession) TokenPeg() (*big.Int, error) {
	return _Mai.Contract.TokenPeg(&_Mai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mai *MaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mai *MaiSession) TotalSupply() (*big.Int, error) {
	return _Mai.Contract.TotalSupply(&_Mai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mai *MaiCallerSession) TotalSupply() (*big.Int, error) {
	return _Mai.Contract.TotalSupply(&_Mai.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Mai *MaiCaller) Treasury(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Mai *MaiSession) Treasury() (*big.Int, error) {
	return _Mai.Contract.Treasury(&_Mai.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(uint256)
func (_Mai *MaiCallerSession) Treasury() (*big.Int, error) {
	return _Mai.Contract.Treasury(&_Mai.CallOpts)
}

// VaultCollateral is a free data retrieval call binding the contract method 0xd4a9b2c5.
//
// Solidity: function vaultCollateral(uint256 ) view returns(uint256)
func (_Mai *MaiCaller) VaultCollateral(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "vaultCollateral", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCollateral is a free data retrieval call binding the contract method 0xd4a9b2c5.
//
// Solidity: function vaultCollateral(uint256 ) view returns(uint256)
func (_Mai *MaiSession) VaultCollateral(arg0 *big.Int) (*big.Int, error) {
	return _Mai.Contract.VaultCollateral(&_Mai.CallOpts, arg0)
}

// VaultCollateral is a free data retrieval call binding the contract method 0xd4a9b2c5.
//
// Solidity: function vaultCollateral(uint256 ) view returns(uint256)
func (_Mai *MaiCallerSession) VaultCollateral(arg0 *big.Int) (*big.Int, error) {
	return _Mai.Contract.VaultCollateral(&_Mai.CallOpts, arg0)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Mai *MaiCaller) VaultCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "vaultCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Mai *MaiSession) VaultCount() (*big.Int, error) {
	return _Mai.Contract.VaultCount(&_Mai.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() view returns(uint256)
func (_Mai *MaiCallerSession) VaultCount() (*big.Int, error) {
	return _Mai.Contract.VaultCount(&_Mai.CallOpts)
}

// VaultDebt is a free data retrieval call binding the contract method 0xd310f49b.
//
// Solidity: function vaultDebt(uint256 ) view returns(uint256)
func (_Mai *MaiCaller) VaultDebt(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "vaultDebt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VaultDebt is a free data retrieval call binding the contract method 0xd310f49b.
//
// Solidity: function vaultDebt(uint256 ) view returns(uint256)
func (_Mai *MaiSession) VaultDebt(arg0 *big.Int) (*big.Int, error) {
	return _Mai.Contract.VaultDebt(&_Mai.CallOpts, arg0)
}

// VaultDebt is a free data retrieval call binding the contract method 0xd310f49b.
//
// Solidity: function vaultDebt(uint256 ) view returns(uint256)
func (_Mai *MaiCallerSession) VaultDebt(arg0 *big.Int) (*big.Int, error) {
	return _Mai.Contract.VaultDebt(&_Mai.CallOpts, arg0)
}

// VaultExistence is a free data retrieval call binding the contract method 0xa525323d.
//
// Solidity: function vaultExistence(uint256 ) view returns(bool)
func (_Mai *MaiCaller) VaultExistence(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "vaultExistence", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultExistence is a free data retrieval call binding the contract method 0xa525323d.
//
// Solidity: function vaultExistence(uint256 ) view returns(bool)
func (_Mai *MaiSession) VaultExistence(arg0 *big.Int) (bool, error) {
	return _Mai.Contract.VaultExistence(&_Mai.CallOpts, arg0)
}

// VaultExistence is a free data retrieval call binding the contract method 0xa525323d.
//
// Solidity: function vaultExistence(uint256 ) view returns(bool)
func (_Mai *MaiCallerSession) VaultExistence(arg0 *big.Int) (bool, error) {
	return _Mai.Contract.VaultExistence(&_Mai.CallOpts, arg0)
}

// VaultOwner is a free data retrieval call binding the contract method 0x3c2ecfe5.
//
// Solidity: function vaultOwner(uint256 ) view returns(address)
func (_Mai *MaiCaller) VaultOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mai.contract.Call(opts, &out, "vaultOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultOwner is a free data retrieval call binding the contract method 0x3c2ecfe5.
//
// Solidity: function vaultOwner(uint256 ) view returns(address)
func (_Mai *MaiSession) VaultOwner(arg0 *big.Int) (common.Address, error) {
	return _Mai.Contract.VaultOwner(&_Mai.CallOpts, arg0)
}

// VaultOwner is a free data retrieval call binding the contract method 0x3c2ecfe5.
//
// Solidity: function vaultOwner(uint256 ) view returns(address)
func (_Mai *MaiCallerSession) VaultOwner(arg0 *big.Int) (common.Address, error) {
	return _Mai.Contract.VaultOwner(&_Mai.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mai *MaiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mai *MaiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Approve(&_Mai.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Mai *MaiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Approve(&_Mai.TransactOpts, spender, amount)
}

// BorrowToken is a paid mutator transaction binding the contract method 0x08ec5927.
//
// Solidity: function borrowToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactor) BorrowToken(opts *bind.TransactOpts, vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "borrowToken", vaultID, amount)
}

// BorrowToken is a paid mutator transaction binding the contract method 0x08ec5927.
//
// Solidity: function borrowToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiSession) BorrowToken(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.BorrowToken(&_Mai.TransactOpts, vaultID, amount)
}

// BorrowToken is a paid mutator transaction binding the contract method 0x08ec5927.
//
// Solidity: function borrowToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactorSession) BorrowToken(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.BorrowToken(&_Mai.TransactOpts, vaultID, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Mai *MaiTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Mai *MaiSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Burn(&_Mai.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Mai *MaiTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Burn(&_Mai.TransactOpts, account, amount)
}

// BuyRiskyVault is a paid mutator transaction binding the contract method 0xce77f243.
//
// Solidity: function buyRiskyVault(uint256 vaultID) returns()
func (_Mai *MaiTransactor) BuyRiskyVault(opts *bind.TransactOpts, vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "buyRiskyVault", vaultID)
}

// BuyRiskyVault is a paid mutator transaction binding the contract method 0xce77f243.
//
// Solidity: function buyRiskyVault(uint256 vaultID) returns()
func (_Mai *MaiSession) BuyRiskyVault(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.BuyRiskyVault(&_Mai.TransactOpts, vaultID)
}

// BuyRiskyVault is a paid mutator transaction binding the contract method 0xce77f243.
//
// Solidity: function buyRiskyVault(uint256 vaultID) returns()
func (_Mai *MaiTransactorSession) BuyRiskyVault(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.BuyRiskyVault(&_Mai.TransactOpts, vaultID)
}

// ChangeEthPriceSource is a paid mutator transaction binding the contract method 0x07960532.
//
// Solidity: function changeEthPriceSource(address ethPriceSourceAddress) returns()
func (_Mai *MaiTransactor) ChangeEthPriceSource(opts *bind.TransactOpts, ethPriceSourceAddress common.Address) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "changeEthPriceSource", ethPriceSourceAddress)
}

// ChangeEthPriceSource is a paid mutator transaction binding the contract method 0x07960532.
//
// Solidity: function changeEthPriceSource(address ethPriceSourceAddress) returns()
func (_Mai *MaiSession) ChangeEthPriceSource(ethPriceSourceAddress common.Address) (*types.Transaction, error) {
	return _Mai.Contract.ChangeEthPriceSource(&_Mai.TransactOpts, ethPriceSourceAddress)
}

// ChangeEthPriceSource is a paid mutator transaction binding the contract method 0x07960532.
//
// Solidity: function changeEthPriceSource(address ethPriceSourceAddress) returns()
func (_Mai *MaiTransactorSession) ChangeEthPriceSource(ethPriceSourceAddress common.Address) (*types.Transaction, error) {
	return _Mai.Contract.ChangeEthPriceSource(&_Mai.TransactOpts, ethPriceSourceAddress)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns(uint256)
func (_Mai *MaiTransactor) CreateVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "createVault")
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns(uint256)
func (_Mai *MaiSession) CreateVault() (*types.Transaction, error) {
	return _Mai.Contract.CreateVault(&_Mai.TransactOpts)
}

// CreateVault is a paid mutator transaction binding the contract method 0x5d12928b.
//
// Solidity: function createVault() returns(uint256)
func (_Mai *MaiTransactorSession) CreateVault() (*types.Transaction, error) {
	return _Mai.Contract.CreateVault(&_Mai.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mai *MaiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mai *MaiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DecreaseAllowance(&_Mai.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Mai *MaiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DecreaseAllowance(&_Mai.TransactOpts, spender, subtractedValue)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 vaultID) payable returns()
func (_Mai *MaiTransactor) DepositCollateral(opts *bind.TransactOpts, vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "depositCollateral", vaultID)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 vaultID) payable returns()
func (_Mai *MaiSession) DepositCollateral(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DepositCollateral(&_Mai.TransactOpts, vaultID)
}

// DepositCollateral is a paid mutator transaction binding the contract method 0xbad4a01f.
//
// Solidity: function depositCollateral(uint256 vaultID) payable returns()
func (_Mai *MaiTransactorSession) DepositCollateral(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DepositCollateral(&_Mai.TransactOpts, vaultID)
}

// DestroyVault is a paid mutator transaction binding the contract method 0x85e290a3.
//
// Solidity: function destroyVault(uint256 vaultID) returns()
func (_Mai *MaiTransactor) DestroyVault(opts *bind.TransactOpts, vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "destroyVault", vaultID)
}

// DestroyVault is a paid mutator transaction binding the contract method 0x85e290a3.
//
// Solidity: function destroyVault(uint256 vaultID) returns()
func (_Mai *MaiSession) DestroyVault(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DestroyVault(&_Mai.TransactOpts, vaultID)
}

// DestroyVault is a paid mutator transaction binding the contract method 0x85e290a3.
//
// Solidity: function destroyVault(uint256 vaultID) returns()
func (_Mai *MaiTransactorSession) DestroyVault(vaultID *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.DestroyVault(&_Mai.TransactOpts, vaultID)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mai *MaiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mai *MaiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.IncreaseAllowance(&_Mai.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Mai *MaiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.IncreaseAllowance(&_Mai.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Mai *MaiTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Mai *MaiSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Mint(&_Mai.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Mai *MaiTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Mint(&_Mai.TransactOpts, account, amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0x85af3c16.
//
// Solidity: function payBackToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactor) PayBackToken(opts *bind.TransactOpts, vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "payBackToken", vaultID, amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0x85af3c16.
//
// Solidity: function payBackToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiSession) PayBackToken(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.PayBackToken(&_Mai.TransactOpts, vaultID, amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0x85af3c16.
//
// Solidity: function payBackToken(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactorSession) PayBackToken(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.PayBackToken(&_Mai.TransactOpts, vaultID, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mai *MaiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mai *MaiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mai.Contract.RenounceOwnership(&_Mai.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mai *MaiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mai.Contract.RenounceOwnership(&_Mai.TransactOpts)
}

// SetClosingFee is a paid mutator transaction binding the contract method 0x3db99177.
//
// Solidity: function setClosingFee(uint256 amount) returns()
func (_Mai *MaiTransactor) SetClosingFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setClosingFee", amount)
}

// SetClosingFee is a paid mutator transaction binding the contract method 0x3db99177.
//
// Solidity: function setClosingFee(uint256 amount) returns()
func (_Mai *MaiSession) SetClosingFee(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetClosingFee(&_Mai.TransactOpts, amount)
}

// SetClosingFee is a paid mutator transaction binding the contract method 0x3db99177.
//
// Solidity: function setClosingFee(uint256 amount) returns()
func (_Mai *MaiTransactorSession) SetClosingFee(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetClosingFee(&_Mai.TransactOpts, amount)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb1511cc9.
//
// Solidity: function setDebtCeiling(uint256 amount) returns()
func (_Mai *MaiTransactor) SetDebtCeiling(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setDebtCeiling", amount)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb1511cc9.
//
// Solidity: function setDebtCeiling(uint256 amount) returns()
func (_Mai *MaiSession) SetDebtCeiling(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetDebtCeiling(&_Mai.TransactOpts, amount)
}

// SetDebtCeiling is a paid mutator transaction binding the contract method 0xb1511cc9.
//
// Solidity: function setDebtCeiling(uint256 amount) returns()
func (_Mai *MaiTransactorSession) SetDebtCeiling(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetDebtCeiling(&_Mai.TransactOpts, amount)
}

// SetOpeningFee is a paid mutator transaction binding the contract method 0x86375994.
//
// Solidity: function setOpeningFee(uint256 amount) returns()
func (_Mai *MaiTransactor) SetOpeningFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setOpeningFee", amount)
}

// SetOpeningFee is a paid mutator transaction binding the contract method 0x86375994.
//
// Solidity: function setOpeningFee(uint256 amount) returns()
func (_Mai *MaiSession) SetOpeningFee(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetOpeningFee(&_Mai.TransactOpts, amount)
}

// SetOpeningFee is a paid mutator transaction binding the contract method 0x86375994.
//
// Solidity: function setOpeningFee(uint256 amount) returns()
func (_Mai *MaiTransactorSession) SetOpeningFee(amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetOpeningFee(&_Mai.TransactOpts, amount)
}

// SetStabilityPool is a paid mutator transaction binding the contract method 0x98d721e0.
//
// Solidity: function setStabilityPool(address _pool) returns()
func (_Mai *MaiTransactor) SetStabilityPool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setStabilityPool", _pool)
}

// SetStabilityPool is a paid mutator transaction binding the contract method 0x98d721e0.
//
// Solidity: function setStabilityPool(address _pool) returns()
func (_Mai *MaiSession) SetStabilityPool(_pool common.Address) (*types.Transaction, error) {
	return _Mai.Contract.SetStabilityPool(&_Mai.TransactOpts, _pool)
}

// SetStabilityPool is a paid mutator transaction binding the contract method 0x98d721e0.
//
// Solidity: function setStabilityPool(address _pool) returns()
func (_Mai *MaiTransactorSession) SetStabilityPool(_pool common.Address) (*types.Transaction, error) {
	return _Mai.Contract.SetStabilityPool(&_Mai.TransactOpts, _pool)
}

// SetTokenPeg is a paid mutator transaction binding the contract method 0x01e49d0a.
//
// Solidity: function setTokenPeg(uint256 _tokenPeg) returns()
func (_Mai *MaiTransactor) SetTokenPeg(opts *bind.TransactOpts, _tokenPeg *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setTokenPeg", _tokenPeg)
}

// SetTokenPeg is a paid mutator transaction binding the contract method 0x01e49d0a.
//
// Solidity: function setTokenPeg(uint256 _tokenPeg) returns()
func (_Mai *MaiSession) SetTokenPeg(_tokenPeg *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetTokenPeg(&_Mai.TransactOpts, _tokenPeg)
}

// SetTokenPeg is a paid mutator transaction binding the contract method 0x01e49d0a.
//
// Solidity: function setTokenPeg(uint256 _tokenPeg) returns()
func (_Mai *MaiTransactorSession) SetTokenPeg(_tokenPeg *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetTokenPeg(&_Mai.TransactOpts, _tokenPeg)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x7139c929.
//
// Solidity: function setTreasury(uint256 _treasury) returns()
func (_Mai *MaiTransactor) SetTreasury(opts *bind.TransactOpts, _treasury *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "setTreasury", _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x7139c929.
//
// Solidity: function setTreasury(uint256 _treasury) returns()
func (_Mai *MaiSession) SetTreasury(_treasury *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetTreasury(&_Mai.TransactOpts, _treasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0x7139c929.
//
// Solidity: function setTreasury(uint256 _treasury) returns()
func (_Mai *MaiTransactorSession) SetTreasury(_treasury *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.SetTreasury(&_Mai.TransactOpts, _treasury)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mai *MaiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mai *MaiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Transfer(&_Mai.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Mai *MaiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.Transfer(&_Mai.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mai *MaiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mai *MaiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.TransferFrom(&_Mai.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Mai *MaiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.TransferFrom(&_Mai.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mai *MaiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mai *MaiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mai.Contract.TransferOwnership(&_Mai.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mai *MaiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mai.Contract.TransferOwnership(&_Mai.TransactOpts, newOwner)
}

// TransferVault is a paid mutator transaction binding the contract method 0x3e61facd.
//
// Solidity: function transferVault(uint256 vaultID, address to) returns()
func (_Mai *MaiTransactor) TransferVault(opts *bind.TransactOpts, vaultID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "transferVault", vaultID, to)
}

// TransferVault is a paid mutator transaction binding the contract method 0x3e61facd.
//
// Solidity: function transferVault(uint256 vaultID, address to) returns()
func (_Mai *MaiSession) TransferVault(vaultID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mai.Contract.TransferVault(&_Mai.TransactOpts, vaultID, to)
}

// TransferVault is a paid mutator transaction binding the contract method 0x3e61facd.
//
// Solidity: function transferVault(uint256 vaultID, address to) returns()
func (_Mai *MaiTransactorSession) TransferVault(vaultID *big.Int, to common.Address) (*types.Transaction, error) {
	return _Mai.Contract.TransferVault(&_Mai.TransactOpts, vaultID, to)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x767a7b05.
//
// Solidity: function withdrawCollateral(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactor) WithdrawCollateral(opts *bind.TransactOpts, vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.contract.Transact(opts, "withdrawCollateral", vaultID, amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x767a7b05.
//
// Solidity: function withdrawCollateral(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiSession) WithdrawCollateral(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.WithdrawCollateral(&_Mai.TransactOpts, vaultID, amount)
}

// WithdrawCollateral is a paid mutator transaction binding the contract method 0x767a7b05.
//
// Solidity: function withdrawCollateral(uint256 vaultID, uint256 amount) returns()
func (_Mai *MaiTransactorSession) WithdrawCollateral(vaultID *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Mai.Contract.WithdrawCollateral(&_Mai.TransactOpts, vaultID, amount)
}

// MaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mai contract.
type MaiApprovalIterator struct {
	Event *MaiApproval // Event containing the contract specifics and raw log

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
func (it *MaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiApproval)
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
		it.Event = new(MaiApproval)
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
func (it *MaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiApproval represents a Approval event raised by the Mai contract.
type MaiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mai *MaiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MaiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mai.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MaiApprovalIterator{contract: _Mai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mai *MaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MaiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Mai.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiApproval)
				if err := _Mai.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Mai *MaiFilterer) ParseApproval(log types.Log) (*MaiApproval, error) {
	event := new(MaiApproval)
	if err := _Mai.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiBorrowTokenIterator is returned from FilterBorrowToken and is used to iterate over the raw logs and unpacked data for BorrowToken events raised by the Mai contract.
type MaiBorrowTokenIterator struct {
	Event *MaiBorrowToken // Event containing the contract specifics and raw log

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
func (it *MaiBorrowTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiBorrowToken)
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
		it.Event = new(MaiBorrowToken)
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
func (it *MaiBorrowTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiBorrowTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiBorrowToken represents a BorrowToken event raised by the Mai contract.
type MaiBorrowToken struct {
	VaultID *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrowToken is a free log retrieval operation binding the contract event 0x3e08df88d8e28f37df9bf227d3142ac506a364403445661a60891a49ed6792ca.
//
// Solidity: event BorrowToken(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) FilterBorrowToken(opts *bind.FilterOpts) (*MaiBorrowTokenIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "BorrowToken")
	if err != nil {
		return nil, err
	}
	return &MaiBorrowTokenIterator{contract: _Mai.contract, event: "BorrowToken", logs: logs, sub: sub}, nil
}

// WatchBorrowToken is a free log subscription operation binding the contract event 0x3e08df88d8e28f37df9bf227d3142ac506a364403445661a60891a49ed6792ca.
//
// Solidity: event BorrowToken(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) WatchBorrowToken(opts *bind.WatchOpts, sink chan<- *MaiBorrowToken) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "BorrowToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiBorrowToken)
				if err := _Mai.contract.UnpackLog(event, "BorrowToken", log); err != nil {
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

// ParseBorrowToken is a log parse operation binding the contract event 0x3e08df88d8e28f37df9bf227d3142ac506a364403445661a60891a49ed6792ca.
//
// Solidity: event BorrowToken(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) ParseBorrowToken(log types.Log) (*MaiBorrowToken, error) {
	event := new(MaiBorrowToken)
	if err := _Mai.contract.UnpackLog(event, "BorrowToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiBuyRiskyVaultIterator is returned from FilterBuyRiskyVault and is used to iterate over the raw logs and unpacked data for BuyRiskyVault events raised by the Mai contract.
type MaiBuyRiskyVaultIterator struct {
	Event *MaiBuyRiskyVault // Event containing the contract specifics and raw log

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
func (it *MaiBuyRiskyVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiBuyRiskyVault)
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
		it.Event = new(MaiBuyRiskyVault)
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
func (it *MaiBuyRiskyVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiBuyRiskyVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiBuyRiskyVault represents a BuyRiskyVault event raised by the Mai contract.
type MaiBuyRiskyVault struct {
	VaultID    *big.Int
	Owner      common.Address
	Buyer      common.Address
	AmountPaid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBuyRiskyVault is a free log retrieval operation binding the contract event 0xd5f2917d6c1a1057183be88f3f1c867cf597030f2cd1197a2fdfe1f51376d0cb.
//
// Solidity: event BuyRiskyVault(uint256 vaultID, address owner, address buyer, uint256 amountPaid)
func (_Mai *MaiFilterer) FilterBuyRiskyVault(opts *bind.FilterOpts) (*MaiBuyRiskyVaultIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "BuyRiskyVault")
	if err != nil {
		return nil, err
	}
	return &MaiBuyRiskyVaultIterator{contract: _Mai.contract, event: "BuyRiskyVault", logs: logs, sub: sub}, nil
}

// WatchBuyRiskyVault is a free log subscription operation binding the contract event 0xd5f2917d6c1a1057183be88f3f1c867cf597030f2cd1197a2fdfe1f51376d0cb.
//
// Solidity: event BuyRiskyVault(uint256 vaultID, address owner, address buyer, uint256 amountPaid)
func (_Mai *MaiFilterer) WatchBuyRiskyVault(opts *bind.WatchOpts, sink chan<- *MaiBuyRiskyVault) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "BuyRiskyVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiBuyRiskyVault)
				if err := _Mai.contract.UnpackLog(event, "BuyRiskyVault", log); err != nil {
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

// ParseBuyRiskyVault is a log parse operation binding the contract event 0xd5f2917d6c1a1057183be88f3f1c867cf597030f2cd1197a2fdfe1f51376d0cb.
//
// Solidity: event BuyRiskyVault(uint256 vaultID, address owner, address buyer, uint256 amountPaid)
func (_Mai *MaiFilterer) ParseBuyRiskyVault(log types.Log) (*MaiBuyRiskyVault, error) {
	event := new(MaiBuyRiskyVault)
	if err := _Mai.contract.UnpackLog(event, "BuyRiskyVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiCreateVaultIterator is returned from FilterCreateVault and is used to iterate over the raw logs and unpacked data for CreateVault events raised by the Mai contract.
type MaiCreateVaultIterator struct {
	Event *MaiCreateVault // Event containing the contract specifics and raw log

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
func (it *MaiCreateVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiCreateVault)
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
		it.Event = new(MaiCreateVault)
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
func (it *MaiCreateVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiCreateVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiCreateVault represents a CreateVault event raised by the Mai contract.
type MaiCreateVault struct {
	VaultID *big.Int
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateVault is a free log retrieval operation binding the contract event 0x8b6c1d05c678fa59695e26465a85918ce0fc63a88f74af53d1daef8f0a9c7804.
//
// Solidity: event CreateVault(uint256 vaultID, address creator)
func (_Mai *MaiFilterer) FilterCreateVault(opts *bind.FilterOpts) (*MaiCreateVaultIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "CreateVault")
	if err != nil {
		return nil, err
	}
	return &MaiCreateVaultIterator{contract: _Mai.contract, event: "CreateVault", logs: logs, sub: sub}, nil
}

// WatchCreateVault is a free log subscription operation binding the contract event 0x8b6c1d05c678fa59695e26465a85918ce0fc63a88f74af53d1daef8f0a9c7804.
//
// Solidity: event CreateVault(uint256 vaultID, address creator)
func (_Mai *MaiFilterer) WatchCreateVault(opts *bind.WatchOpts, sink chan<- *MaiCreateVault) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "CreateVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiCreateVault)
				if err := _Mai.contract.UnpackLog(event, "CreateVault", log); err != nil {
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

// ParseCreateVault is a log parse operation binding the contract event 0x8b6c1d05c678fa59695e26465a85918ce0fc63a88f74af53d1daef8f0a9c7804.
//
// Solidity: event CreateVault(uint256 vaultID, address creator)
func (_Mai *MaiFilterer) ParseCreateVault(log types.Log) (*MaiCreateVault, error) {
	event := new(MaiCreateVault)
	if err := _Mai.contract.UnpackLog(event, "CreateVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiDepositCollateralIterator is returned from FilterDepositCollateral and is used to iterate over the raw logs and unpacked data for DepositCollateral events raised by the Mai contract.
type MaiDepositCollateralIterator struct {
	Event *MaiDepositCollateral // Event containing the contract specifics and raw log

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
func (it *MaiDepositCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiDepositCollateral)
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
		it.Event = new(MaiDepositCollateral)
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
func (it *MaiDepositCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiDepositCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiDepositCollateral represents a DepositCollateral event raised by the Mai contract.
type MaiDepositCollateral struct {
	VaultID *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositCollateral is a free log retrieval operation binding the contract event 0x52c4e7127ec34e8fc95f09ce2d06b4f00acca12ccbcdfb246ef67ee6aefe068d.
//
// Solidity: event DepositCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) FilterDepositCollateral(opts *bind.FilterOpts) (*MaiDepositCollateralIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return &MaiDepositCollateralIterator{contract: _Mai.contract, event: "DepositCollateral", logs: logs, sub: sub}, nil
}

// WatchDepositCollateral is a free log subscription operation binding the contract event 0x52c4e7127ec34e8fc95f09ce2d06b4f00acca12ccbcdfb246ef67ee6aefe068d.
//
// Solidity: event DepositCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) WatchDepositCollateral(opts *bind.WatchOpts, sink chan<- *MaiDepositCollateral) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "DepositCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiDepositCollateral)
				if err := _Mai.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
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

// ParseDepositCollateral is a log parse operation binding the contract event 0x52c4e7127ec34e8fc95f09ce2d06b4f00acca12ccbcdfb246ef67ee6aefe068d.
//
// Solidity: event DepositCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) ParseDepositCollateral(log types.Log) (*MaiDepositCollateral, error) {
	event := new(MaiDepositCollateral)
	if err := _Mai.contract.UnpackLog(event, "DepositCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiDestroyVaultIterator is returned from FilterDestroyVault and is used to iterate over the raw logs and unpacked data for DestroyVault events raised by the Mai contract.
type MaiDestroyVaultIterator struct {
	Event *MaiDestroyVault // Event containing the contract specifics and raw log

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
func (it *MaiDestroyVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiDestroyVault)
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
		it.Event = new(MaiDestroyVault)
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
func (it *MaiDestroyVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiDestroyVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiDestroyVault represents a DestroyVault event raised by the Mai contract.
type MaiDestroyVault struct {
	VaultID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDestroyVault is a free log retrieval operation binding the contract event 0x4fe08624ee65b341c38ab9693d216b909d4ddee1bc8d3fe0fea14026c361b465.
//
// Solidity: event DestroyVault(uint256 vaultID)
func (_Mai *MaiFilterer) FilterDestroyVault(opts *bind.FilterOpts) (*MaiDestroyVaultIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "DestroyVault")
	if err != nil {
		return nil, err
	}
	return &MaiDestroyVaultIterator{contract: _Mai.contract, event: "DestroyVault", logs: logs, sub: sub}, nil
}

// WatchDestroyVault is a free log subscription operation binding the contract event 0x4fe08624ee65b341c38ab9693d216b909d4ddee1bc8d3fe0fea14026c361b465.
//
// Solidity: event DestroyVault(uint256 vaultID)
func (_Mai *MaiFilterer) WatchDestroyVault(opts *bind.WatchOpts, sink chan<- *MaiDestroyVault) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "DestroyVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiDestroyVault)
				if err := _Mai.contract.UnpackLog(event, "DestroyVault", log); err != nil {
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

// ParseDestroyVault is a log parse operation binding the contract event 0x4fe08624ee65b341c38ab9693d216b909d4ddee1bc8d3fe0fea14026c361b465.
//
// Solidity: event DestroyVault(uint256 vaultID)
func (_Mai *MaiFilterer) ParseDestroyVault(log types.Log) (*MaiDestroyVault, error) {
	event := new(MaiDestroyVault)
	if err := _Mai.contract.UnpackLog(event, "DestroyVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mai contract.
type MaiOwnershipTransferredIterator struct {
	Event *MaiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MaiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiOwnershipTransferred)
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
		it.Event = new(MaiOwnershipTransferred)
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
func (it *MaiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiOwnershipTransferred represents a OwnershipTransferred event raised by the Mai contract.
type MaiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mai *MaiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MaiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mai.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MaiOwnershipTransferredIterator{contract: _Mai.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mai *MaiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MaiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mai.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiOwnershipTransferred)
				if err := _Mai.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mai *MaiFilterer) ParseOwnershipTransferred(log types.Log) (*MaiOwnershipTransferred, error) {
	event := new(MaiOwnershipTransferred)
	if err := _Mai.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiPayBackTokenIterator is returned from FilterPayBackToken and is used to iterate over the raw logs and unpacked data for PayBackToken events raised by the Mai contract.
type MaiPayBackTokenIterator struct {
	Event *MaiPayBackToken // Event containing the contract specifics and raw log

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
func (it *MaiPayBackTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiPayBackToken)
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
		it.Event = new(MaiPayBackToken)
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
func (it *MaiPayBackTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiPayBackTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiPayBackToken represents a PayBackToken event raised by the Mai contract.
type MaiPayBackToken struct {
	VaultID    *big.Int
	Amount     *big.Int
	ClosingFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPayBackToken is a free log retrieval operation binding the contract event 0x31f96762af4051f367185773cc2f55bfb112a6c114b3407ded1f321a9eb199ac.
//
// Solidity: event PayBackToken(uint256 vaultID, uint256 amount, uint256 closingFee)
func (_Mai *MaiFilterer) FilterPayBackToken(opts *bind.FilterOpts) (*MaiPayBackTokenIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "PayBackToken")
	if err != nil {
		return nil, err
	}
	return &MaiPayBackTokenIterator{contract: _Mai.contract, event: "PayBackToken", logs: logs, sub: sub}, nil
}

// WatchPayBackToken is a free log subscription operation binding the contract event 0x31f96762af4051f367185773cc2f55bfb112a6c114b3407ded1f321a9eb199ac.
//
// Solidity: event PayBackToken(uint256 vaultID, uint256 amount, uint256 closingFee)
func (_Mai *MaiFilterer) WatchPayBackToken(opts *bind.WatchOpts, sink chan<- *MaiPayBackToken) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "PayBackToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiPayBackToken)
				if err := _Mai.contract.UnpackLog(event, "PayBackToken", log); err != nil {
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

// ParsePayBackToken is a log parse operation binding the contract event 0x31f96762af4051f367185773cc2f55bfb112a6c114b3407ded1f321a9eb199ac.
//
// Solidity: event PayBackToken(uint256 vaultID, uint256 amount, uint256 closingFee)
func (_Mai *MaiFilterer) ParsePayBackToken(log types.Log) (*MaiPayBackToken, error) {
	event := new(MaiPayBackToken)
	if err := _Mai.contract.UnpackLog(event, "PayBackToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mai contract.
type MaiTransferIterator struct {
	Event *MaiTransfer // Event containing the contract specifics and raw log

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
func (it *MaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiTransfer)
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
		it.Event = new(MaiTransfer)
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
func (it *MaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiTransfer represents a Transfer event raised by the Mai contract.
type MaiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mai *MaiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MaiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mai.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MaiTransferIterator{contract: _Mai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mai *MaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MaiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Mai.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiTransfer)
				if err := _Mai.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Mai *MaiFilterer) ParseTransfer(log types.Log) (*MaiTransfer, error) {
	event := new(MaiTransfer)
	if err := _Mai.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiTransferVaultIterator is returned from FilterTransferVault and is used to iterate over the raw logs and unpacked data for TransferVault events raised by the Mai contract.
type MaiTransferVaultIterator struct {
	Event *MaiTransferVault // Event containing the contract specifics and raw log

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
func (it *MaiTransferVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiTransferVault)
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
		it.Event = new(MaiTransferVault)
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
func (it *MaiTransferVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiTransferVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiTransferVault represents a TransferVault event raised by the Mai contract.
type MaiTransferVault struct {
	VaultID *big.Int
	From    common.Address
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferVault is a free log retrieval operation binding the contract event 0xa8159d062909288cfe1bc71a9cb71a800664f2658fc2588d52676a844f1b0f13.
//
// Solidity: event TransferVault(uint256 vaultID, address from, address to)
func (_Mai *MaiFilterer) FilterTransferVault(opts *bind.FilterOpts) (*MaiTransferVaultIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "TransferVault")
	if err != nil {
		return nil, err
	}
	return &MaiTransferVaultIterator{contract: _Mai.contract, event: "TransferVault", logs: logs, sub: sub}, nil
}

// WatchTransferVault is a free log subscription operation binding the contract event 0xa8159d062909288cfe1bc71a9cb71a800664f2658fc2588d52676a844f1b0f13.
//
// Solidity: event TransferVault(uint256 vaultID, address from, address to)
func (_Mai *MaiFilterer) WatchTransferVault(opts *bind.WatchOpts, sink chan<- *MaiTransferVault) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "TransferVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiTransferVault)
				if err := _Mai.contract.UnpackLog(event, "TransferVault", log); err != nil {
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

// ParseTransferVault is a log parse operation binding the contract event 0xa8159d062909288cfe1bc71a9cb71a800664f2658fc2588d52676a844f1b0f13.
//
// Solidity: event TransferVault(uint256 vaultID, address from, address to)
func (_Mai *MaiFilterer) ParseTransferVault(log types.Log) (*MaiTransferVault, error) {
	event := new(MaiTransferVault)
	if err := _Mai.contract.UnpackLog(event, "TransferVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MaiWithdrawCollateralIterator is returned from FilterWithdrawCollateral and is used to iterate over the raw logs and unpacked data for WithdrawCollateral events raised by the Mai contract.
type MaiWithdrawCollateralIterator struct {
	Event *MaiWithdrawCollateral // Event containing the contract specifics and raw log

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
func (it *MaiWithdrawCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MaiWithdrawCollateral)
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
		it.Event = new(MaiWithdrawCollateral)
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
func (it *MaiWithdrawCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MaiWithdrawCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MaiWithdrawCollateral represents a WithdrawCollateral event raised by the Mai contract.
type MaiWithdrawCollateral struct {
	VaultID *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawCollateral is a free log retrieval operation binding the contract event 0x6c0ea3bea9dd66afa8f9d39d6eb93d833466190330813b42835efc650dca4cb9.
//
// Solidity: event WithdrawCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) FilterWithdrawCollateral(opts *bind.FilterOpts) (*MaiWithdrawCollateralIterator, error) {

	logs, sub, err := _Mai.contract.FilterLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return &MaiWithdrawCollateralIterator{contract: _Mai.contract, event: "WithdrawCollateral", logs: logs, sub: sub}, nil
}

// WatchWithdrawCollateral is a free log subscription operation binding the contract event 0x6c0ea3bea9dd66afa8f9d39d6eb93d833466190330813b42835efc650dca4cb9.
//
// Solidity: event WithdrawCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) WatchWithdrawCollateral(opts *bind.WatchOpts, sink chan<- *MaiWithdrawCollateral) (event.Subscription, error) {

	logs, sub, err := _Mai.contract.WatchLogs(opts, "WithdrawCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MaiWithdrawCollateral)
				if err := _Mai.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
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

// ParseWithdrawCollateral is a log parse operation binding the contract event 0x6c0ea3bea9dd66afa8f9d39d6eb93d833466190330813b42835efc650dca4cb9.
//
// Solidity: event WithdrawCollateral(uint256 vaultID, uint256 amount)
func (_Mai *MaiFilterer) ParseWithdrawCollateral(log types.Log) (*MaiWithdrawCollateral, error) {
	event := new(MaiWithdrawCollateral)
	if err := _Mai.contract.UnpackLog(event, "WithdrawCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
