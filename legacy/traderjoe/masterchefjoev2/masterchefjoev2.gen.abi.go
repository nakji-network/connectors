// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masterchefjoev2

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
)

// Masterchefjoev2MetaData contains all meta data concerning the Masterchefjoev2 contract.
var Masterchefjoev2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractJoeToken\",\"name\":\"_joe\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_devAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_investorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_devPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_investorPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"SetDevAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"}],\"name\":\"UpdateEmissionRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_devAddr\",\"type\":\"address\"}],\"name\":\"dev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investorAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investorPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joe\",\"outputs\":[{\"internalType\":\"contractJoeToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joePerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingJoe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"rewarderBonusTokenInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDevPercent\",\"type\":\"uint256\"}],\"name\":\"setDevPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_investorAddr\",\"type\":\"address\"}],\"name\":\"setInvestorAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newInvestorPercent\",\"type\":\"uint256\"}],\"name\":\"setInvestorPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryAddr\",\"type\":\"address\"}],\"name\":\"setTreasuryAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newTreasuryPercent\",\"type\":\"uint256\"}],\"name\":\"setTreasuryPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_joePerSec\",\"type\":\"uint256\"}],\"name\":\"updateEmissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Masterchefjoev2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Masterchefjoev2MetaData.ABI instead.
var Masterchefjoev2ABI = Masterchefjoev2MetaData.ABI

// Masterchefjoev2 is an auto generated Go binding around an Ethereum contract.
type Masterchefjoev2 struct {
	Masterchefjoev2Caller     // Read-only binding to the contract
	Masterchefjoev2Transactor // Write-only binding to the contract
	Masterchefjoev2Filterer   // Log filterer for contract events
}

// Masterchefjoev2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Masterchefjoev2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Masterchefjoev2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Masterchefjoev2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Masterchefjoev2Session struct {
	Contract     *Masterchefjoev2  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Masterchefjoev2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Masterchefjoev2CallerSession struct {
	Contract *Masterchefjoev2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Masterchefjoev2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Masterchefjoev2TransactorSession struct {
	Contract     *Masterchefjoev2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Masterchefjoev2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Masterchefjoev2Raw struct {
	Contract *Masterchefjoev2 // Generic contract binding to access the raw methods on
}

// Masterchefjoev2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Masterchefjoev2CallerRaw struct {
	Contract *Masterchefjoev2Caller // Generic read-only contract binding to access the raw methods on
}

// Masterchefjoev2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Masterchefjoev2TransactorRaw struct {
	Contract *Masterchefjoev2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterchefjoev2 creates a new instance of Masterchefjoev2, bound to a specific deployed contract.
func NewMasterchefjoev2(address common.Address, backend bind.ContractBackend) (*Masterchefjoev2, error) {
	contract, err := bindMasterchefjoev2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2{Masterchefjoev2Caller: Masterchefjoev2Caller{contract: contract}, Masterchefjoev2Transactor: Masterchefjoev2Transactor{contract: contract}, Masterchefjoev2Filterer: Masterchefjoev2Filterer{contract: contract}}, nil
}

// NewMasterchefjoev2Caller creates a new read-only instance of Masterchefjoev2, bound to a specific deployed contract.
func NewMasterchefjoev2Caller(address common.Address, caller bind.ContractCaller) (*Masterchefjoev2Caller, error) {
	contract, err := bindMasterchefjoev2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2Caller{contract: contract}, nil
}

// NewMasterchefjoev2Transactor creates a new write-only instance of Masterchefjoev2, bound to a specific deployed contract.
func NewMasterchefjoev2Transactor(address common.Address, transactor bind.ContractTransactor) (*Masterchefjoev2Transactor, error) {
	contract, err := bindMasterchefjoev2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2Transactor{contract: contract}, nil
}

// NewMasterchefjoev2Filterer creates a new log filterer instance of Masterchefjoev2, bound to a specific deployed contract.
func NewMasterchefjoev2Filterer(address common.Address, filterer bind.ContractFilterer) (*Masterchefjoev2Filterer, error) {
	contract, err := bindMasterchefjoev2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2Filterer{contract: contract}, nil
}

// bindMasterchefjoev2 binds a generic wrapper to an already deployed contract.
func bindMasterchefjoev2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Masterchefjoev2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchefjoev2 *Masterchefjoev2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchefjoev2.Contract.Masterchefjoev2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchefjoev2 *Masterchefjoev2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Masterchefjoev2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchefjoev2 *Masterchefjoev2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Masterchefjoev2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchefjoev2 *Masterchefjoev2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchefjoev2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchefjoev2 *Masterchefjoev2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchefjoev2 *Masterchefjoev2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.contract.Transact(opts, method, params...)
}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Caller) DevAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "devAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Session) DevAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.DevAddr(&_Masterchefjoev2.CallOpts)
}

// DevAddr is a free data retrieval call binding the contract method 0xda09c72c.
//
// Solidity: function devAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) DevAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.DevAddr(&_Masterchefjoev2.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) DevPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "devPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) DevPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.DevPercent(&_Masterchefjoev2.CallOpts)
}

// DevPercent is a free data retrieval call binding the contract method 0xfc3c28af.
//
// Solidity: function devPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) DevPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.DevPercent(&_Masterchefjoev2.CallOpts)
}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Caller) InvestorAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "investorAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Session) InvestorAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.InvestorAddr(&_Masterchefjoev2.CallOpts)
}

// InvestorAddr is a free data retrieval call binding the contract method 0xacc4cc50.
//
// Solidity: function investorAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) InvestorAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.InvestorAddr(&_Masterchefjoev2.CallOpts)
}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) InvestorPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "investorPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) InvestorPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.InvestorPercent(&_Masterchefjoev2.CallOpts)
}

// InvestorPercent is a free data retrieval call binding the contract method 0x0735b208.
//
// Solidity: function investorPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) InvestorPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.InvestorPercent(&_Masterchefjoev2.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Caller) Joe(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "joe")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Session) Joe() (common.Address, error) {
	return _Masterchefjoev2.Contract.Joe(&_Masterchefjoev2.CallOpts)
}

// Joe is a free data retrieval call binding the contract method 0xb985a3a0.
//
// Solidity: function joe() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) Joe() (common.Address, error) {
	return _Masterchefjoev2.Contract.Joe(&_Masterchefjoev2.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) JoePerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "joePerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) JoePerSec() (*big.Int, error) {
	return _Masterchefjoev2.Contract.JoePerSec(&_Masterchefjoev2.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) JoePerSec() (*big.Int, error) {
	return _Masterchefjoev2.Contract.JoePerSec(&_Masterchefjoev2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Session) Owner() (common.Address, error) {
	return _Masterchefjoev2.Contract.Owner(&_Masterchefjoev2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) Owner() (common.Address, error) {
	return _Masterchefjoev2.Contract.Owner(&_Masterchefjoev2.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_Masterchefjoev2 *Masterchefjoev2Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "pendingTokens", _pid, _user)

	outstruct := new(struct {
		PendingJoe        *big.Int
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
		PendingBonusToken *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PendingJoe = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BonusTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PendingBonusToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_Masterchefjoev2 *Masterchefjoev2Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _Masterchefjoev2.Contract.PendingTokens(&_Masterchefjoev2.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _Masterchefjoev2.Contract.PendingTokens(&_Masterchefjoev2.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_Masterchefjoev2 *Masterchefjoev2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AllocPoint          *big.Int
		LastRewardTimestamp *big.Int
		AccJoePerShare      *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccJoePerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_Masterchefjoev2 *Masterchefjoev2Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _Masterchefjoev2.Contract.PoolInfo(&_Masterchefjoev2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardTimestamp, uint256 accJoePerShare, address rewarder)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AllocPoint          *big.Int
	LastRewardTimestamp *big.Int
	AccJoePerShare      *big.Int
	Rewarder            common.Address
}, error) {
	return _Masterchefjoev2.Contract.PoolInfo(&_Masterchefjoev2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) PoolLength() (*big.Int, error) {
	return _Masterchefjoev2.Contract.PoolLength(&_Masterchefjoev2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) PoolLength() (*big.Int, error) {
	return _Masterchefjoev2.Contract.PoolLength(&_Masterchefjoev2.CallOpts)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_Masterchefjoev2 *Masterchefjoev2Caller) RewarderBonusTokenInfo(opts *bind.CallOpts, _pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "rewarderBonusTokenInfo", _pid)

	outstruct := new(struct {
		BonusTokenAddress common.Address
		BonusTokenSymbol  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BonusTokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BonusTokenSymbol = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_Masterchefjoev2 *Masterchefjoev2Session) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _Masterchefjoev2.Contract.RewarderBonusTokenInfo(&_Masterchefjoev2.CallOpts, _pid)
}

// RewarderBonusTokenInfo is a free data retrieval call binding the contract method 0xbc70fdbc.
//
// Solidity: function rewarderBonusTokenInfo(uint256 _pid) view returns(address bonusTokenAddress, string bonusTokenSymbol)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) RewarderBonusTokenInfo(_pid *big.Int) (struct {
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
}, error) {
	return _Masterchefjoev2.Contract.RewarderBonusTokenInfo(&_Masterchefjoev2.CallOpts, _pid)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) StartTimestamp() (*big.Int, error) {
	return _Masterchefjoev2.Contract.StartTimestamp(&_Masterchefjoev2.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) StartTimestamp() (*big.Int, error) {
	return _Masterchefjoev2.Contract.StartTimestamp(&_Masterchefjoev2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) TotalAllocPoint() (*big.Int, error) {
	return _Masterchefjoev2.Contract.TotalAllocPoint(&_Masterchefjoev2.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Masterchefjoev2.Contract.TotalAllocPoint(&_Masterchefjoev2.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Caller) TreasuryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "treasuryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2Session) TreasuryAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.TreasuryAddr(&_Masterchefjoev2.CallOpts)
}

// TreasuryAddr is a free data retrieval call binding the contract method 0x30d9a62a.
//
// Solidity: function treasuryAddr() view returns(address)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) TreasuryAddr() (common.Address, error) {
	return _Masterchefjoev2.Contract.TreasuryAddr(&_Masterchefjoev2.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Caller) TreasuryPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "treasuryPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2Session) TreasuryPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.TreasuryPercent(&_Masterchefjoev2.CallOpts)
}

// TreasuryPercent is a free data retrieval call binding the contract method 0x04ef9d58.
//
// Solidity: function treasuryPercent() view returns(uint256)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) TreasuryPercent() (*big.Int, error) {
	return _Masterchefjoev2.Contract.TreasuryPercent(&_Masterchefjoev2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchefjoev2 *Masterchefjoev2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Masterchefjoev2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchefjoev2 *Masterchefjoev2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchefjoev2.Contract.UserInfo(&_Masterchefjoev2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchefjoev2 *Masterchefjoev2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchefjoev2.Contract.UserInfo(&_Masterchefjoev2.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "add", _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Add(&_Masterchefjoev2.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Add(&_Masterchefjoev2.TransactOpts, _allocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Deposit(&_Masterchefjoev2.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Deposit(&_Masterchefjoev2.TransactOpts, _pid, _amount)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) Dev(opts *bind.TransactOpts, _devAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "dev", _devAddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) Dev(_devAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Dev(&_Masterchefjoev2.TransactOpts, _devAddr)
}

// Dev is a paid mutator transaction binding the contract method 0x8d88a90e.
//
// Solidity: function dev(address _devAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) Dev(_devAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Dev(&_Masterchefjoev2.TransactOpts, _devAddr)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.EmergencyWithdraw(&_Masterchefjoev2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.EmergencyWithdraw(&_Masterchefjoev2.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) MassUpdatePools() (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.MassUpdatePools(&_Masterchefjoev2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.MassUpdatePools(&_Masterchefjoev2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.RenounceOwnership(&_Masterchefjoev2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.RenounceOwnership(&_Masterchefjoev2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Set(&_Masterchefjoev2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Set(&_Masterchefjoev2.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) SetDevPercent(opts *bind.TransactOpts, _newDevPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "setDevPercent", _newDevPercent)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) SetDevPercent(_newDevPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetDevPercent(&_Masterchefjoev2.TransactOpts, _newDevPercent)
}

// SetDevPercent is a paid mutator transaction binding the contract method 0x6eaddad2.
//
// Solidity: function setDevPercent(uint256 _newDevPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) SetDevPercent(_newDevPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetDevPercent(&_Masterchefjoev2.TransactOpts, _newDevPercent)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) SetInvestorAddr(opts *bind.TransactOpts, _investorAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "setInvestorAddr", _investorAddr)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) SetInvestorAddr(_investorAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetInvestorAddr(&_Masterchefjoev2.TransactOpts, _investorAddr)
}

// SetInvestorAddr is a paid mutator transaction binding the contract method 0x0f51f8ff.
//
// Solidity: function setInvestorAddr(address _investorAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) SetInvestorAddr(_investorAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetInvestorAddr(&_Masterchefjoev2.TransactOpts, _investorAddr)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) SetInvestorPercent(opts *bind.TransactOpts, _newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "setInvestorPercent", _newInvestorPercent)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) SetInvestorPercent(_newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetInvestorPercent(&_Masterchefjoev2.TransactOpts, _newInvestorPercent)
}

// SetInvestorPercent is a paid mutator transaction binding the contract method 0x876d3c9c.
//
// Solidity: function setInvestorPercent(uint256 _newInvestorPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) SetInvestorPercent(_newInvestorPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetInvestorPercent(&_Masterchefjoev2.TransactOpts, _newInvestorPercent)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) SetTreasuryAddr(opts *bind.TransactOpts, _treasuryAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "setTreasuryAddr", _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetTreasuryAddr(&_Masterchefjoev2.TransactOpts, _treasuryAddr)
}

// SetTreasuryAddr is a paid mutator transaction binding the contract method 0xa7e05b9c.
//
// Solidity: function setTreasuryAddr(address _treasuryAddr) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) SetTreasuryAddr(_treasuryAddr common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetTreasuryAddr(&_Masterchefjoev2.TransactOpts, _treasuryAddr)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) SetTreasuryPercent(opts *bind.TransactOpts, _newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "setTreasuryPercent", _newTreasuryPercent)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) SetTreasuryPercent(_newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetTreasuryPercent(&_Masterchefjoev2.TransactOpts, _newTreasuryPercent)
}

// SetTreasuryPercent is a paid mutator transaction binding the contract method 0x89a2bc25.
//
// Solidity: function setTreasuryPercent(uint256 _newTreasuryPercent) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) SetTreasuryPercent(_newTreasuryPercent *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.SetTreasuryPercent(&_Masterchefjoev2.TransactOpts, _newTreasuryPercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.TransferOwnership(&_Masterchefjoev2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.TransferOwnership(&_Masterchefjoev2.TransactOpts, newOwner)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) UpdateEmissionRate(opts *bind.TransactOpts, _joePerSec *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "updateEmissionRate", _joePerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) UpdateEmissionRate(_joePerSec *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.UpdateEmissionRate(&_Masterchefjoev2.TransactOpts, _joePerSec)
}

// UpdateEmissionRate is a paid mutator transaction binding the contract method 0x0ba84cd2.
//
// Solidity: function updateEmissionRate(uint256 _joePerSec) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) UpdateEmissionRate(_joePerSec *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.UpdateEmissionRate(&_Masterchefjoev2.TransactOpts, _joePerSec)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.UpdatePool(&_Masterchefjoev2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.UpdatePool(&_Masterchefjoev2.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Withdraw(&_Masterchefjoev2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Masterchefjoev2 *Masterchefjoev2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev2.Contract.Withdraw(&_Masterchefjoev2.TransactOpts, _pid, _amount)
}

// Masterchefjoev2AddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Masterchefjoev2 contract.
type Masterchefjoev2AddIterator struct {
	Event *Masterchefjoev2Add // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2AddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2Add)
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
		it.Event = new(Masterchefjoev2Add)
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
func (it *Masterchefjoev2AddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2AddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2Add represents a Add event raised by the Masterchefjoev2 contract.
type Masterchefjoev2Add struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterAdd(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*Masterchefjoev2AddIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2AddIterator{contract: _Masterchefjoev2.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2Add, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}
	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2Add)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseAdd(log types.Log) (*Masterchefjoev2Add, error) {
	event := new(Masterchefjoev2Add)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Masterchefjoev2 contract.
type Masterchefjoev2DepositIterator struct {
	Event *Masterchefjoev2Deposit // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2Deposit)
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
		it.Event = new(Masterchefjoev2Deposit)
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
func (it *Masterchefjoev2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2Deposit represents a Deposit event raised by the Masterchefjoev2 contract.
type Masterchefjoev2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2DepositIterator{contract: _Masterchefjoev2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2Deposit)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseDeposit(log types.Log) (*Masterchefjoev2Deposit, error) {
	event := new(Masterchefjoev2Deposit)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Masterchefjoev2 contract.
type Masterchefjoev2EmergencyWithdrawIterator struct {
	Event *Masterchefjoev2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2EmergencyWithdraw)
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
		it.Event = new(Masterchefjoev2EmergencyWithdraw)
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
func (it *Masterchefjoev2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2EmergencyWithdraw represents a EmergencyWithdraw event raised by the Masterchefjoev2 contract.
type Masterchefjoev2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2EmergencyWithdrawIterator{contract: _Masterchefjoev2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2EmergencyWithdraw)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseEmergencyWithdraw(log types.Log) (*Masterchefjoev2EmergencyWithdraw, error) {
	event := new(Masterchefjoev2EmergencyWithdraw)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2HarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Masterchefjoev2 contract.
type Masterchefjoev2HarvestIterator struct {
	Event *Masterchefjoev2Harvest // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2HarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2Harvest)
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
		it.Event = new(Masterchefjoev2Harvest)
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
func (it *Masterchefjoev2HarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2HarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2Harvest represents a Harvest event raised by the Masterchefjoev2 contract.
type Masterchefjoev2Harvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev2HarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2HarvestIterator{contract: _Masterchefjoev2.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2Harvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2Harvest)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseHarvest(log types.Log) (*Masterchefjoev2Harvest, error) {
	event := new(Masterchefjoev2Harvest)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Masterchefjoev2 contract.
type Masterchefjoev2OwnershipTransferredIterator struct {
	Event *Masterchefjoev2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2OwnershipTransferred)
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
		it.Event = new(Masterchefjoev2OwnershipTransferred)
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
func (it *Masterchefjoev2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2OwnershipTransferred represents a OwnershipTransferred event raised by the Masterchefjoev2 contract.
type Masterchefjoev2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Masterchefjoev2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2OwnershipTransferredIterator{contract: _Masterchefjoev2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2OwnershipTransferred)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseOwnershipTransferred(log types.Log) (*Masterchefjoev2OwnershipTransferred, error) {
	event := new(Masterchefjoev2OwnershipTransferred)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2SetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Masterchefjoev2 contract.
type Masterchefjoev2SetIterator struct {
	Event *Masterchefjoev2Set // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2Set)
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
		it.Event = new(Masterchefjoev2Set)
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
func (it *Masterchefjoev2SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2Set represents a Set event raised by the Masterchefjoev2 contract.
type Masterchefjoev2Set struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterSet(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*Masterchefjoev2SetIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2SetIterator{contract: _Masterchefjoev2.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchSet(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2Set, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2Set)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseSet(log types.Log) (*Masterchefjoev2Set, error) {
	event := new(Masterchefjoev2Set)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2SetDevAddressIterator is returned from FilterSetDevAddress and is used to iterate over the raw logs and unpacked data for SetDevAddress events raised by the Masterchefjoev2 contract.
type Masterchefjoev2SetDevAddressIterator struct {
	Event *Masterchefjoev2SetDevAddress // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2SetDevAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2SetDevAddress)
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
		it.Event = new(Masterchefjoev2SetDevAddress)
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
func (it *Masterchefjoev2SetDevAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2SetDevAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2SetDevAddress represents a SetDevAddress event raised by the Masterchefjoev2 contract.
type Masterchefjoev2SetDevAddress struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDevAddress is a free log retrieval operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterSetDevAddress(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*Masterchefjoev2SetDevAddressIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "SetDevAddress", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2SetDevAddressIterator{contract: _Masterchefjoev2.contract, event: "SetDevAddress", logs: logs, sub: sub}, nil
}

// WatchSetDevAddress is a free log subscription operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchSetDevAddress(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2SetDevAddress, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "SetDevAddress", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2SetDevAddress)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "SetDevAddress", log); err != nil {
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

// ParseSetDevAddress is a log parse operation binding the contract event 0x618c54559e94f1499a808aad71ee8729f8e74e8c48e979616328ce493a1a52e7.
//
// Solidity: event SetDevAddress(address indexed oldAddress, address indexed newAddress)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseSetDevAddress(log types.Log) (*Masterchefjoev2SetDevAddress, error) {
	event := new(Masterchefjoev2SetDevAddress)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "SetDevAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2UpdateEmissionRateIterator is returned from FilterUpdateEmissionRate and is used to iterate over the raw logs and unpacked data for UpdateEmissionRate events raised by the Masterchefjoev2 contract.
type Masterchefjoev2UpdateEmissionRateIterator struct {
	Event *Masterchefjoev2UpdateEmissionRate // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2UpdateEmissionRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2UpdateEmissionRate)
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
		it.Event = new(Masterchefjoev2UpdateEmissionRate)
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
func (it *Masterchefjoev2UpdateEmissionRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2UpdateEmissionRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2UpdateEmissionRate represents a UpdateEmissionRate event raised by the Masterchefjoev2 contract.
type Masterchefjoev2UpdateEmissionRate struct {
	User      common.Address
	JoePerSec *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateEmissionRate is a free log retrieval operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterUpdateEmissionRate(opts *bind.FilterOpts, user []common.Address) (*Masterchefjoev2UpdateEmissionRateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2UpdateEmissionRateIterator{contract: _Masterchefjoev2.contract, event: "UpdateEmissionRate", logs: logs, sub: sub}, nil
}

// WatchUpdateEmissionRate is a free log subscription operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchUpdateEmissionRate(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2UpdateEmissionRate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "UpdateEmissionRate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2UpdateEmissionRate)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
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

// ParseUpdateEmissionRate is a log parse operation binding the contract event 0xe2492e003bbe8afa53088b406f0c1cb5d9e280370fc72a74cf116ffd343c4053.
//
// Solidity: event UpdateEmissionRate(address indexed user, uint256 _joePerSec)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseUpdateEmissionRate(log types.Log) (*Masterchefjoev2UpdateEmissionRate, error) {
	event := new(Masterchefjoev2UpdateEmissionRate)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "UpdateEmissionRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2UpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the Masterchefjoev2 contract.
type Masterchefjoev2UpdatePoolIterator struct {
	Event *Masterchefjoev2UpdatePool // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2UpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2UpdatePool)
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
		it.Event = new(Masterchefjoev2UpdatePool)
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
func (it *Masterchefjoev2UpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2UpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2UpdatePool represents a UpdatePool event raised by the Masterchefjoev2 contract.
type Masterchefjoev2UpdatePool struct {
	Pid                 *big.Int
	LastRewardTimestamp *big.Int
	LpSupply            *big.Int
	AccJoePerShare      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*Masterchefjoev2UpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2UpdatePoolIterator{contract: _Masterchefjoev2.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2UpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2UpdatePool)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseUpdatePool(log types.Log) (*Masterchefjoev2UpdatePool, error) {
	event := new(Masterchefjoev2UpdatePool)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Masterchefjoev2 contract.
type Masterchefjoev2WithdrawIterator struct {
	Event *Masterchefjoev2Withdraw // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev2Withdraw)
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
		it.Event = new(Masterchefjoev2Withdraw)
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
func (it *Masterchefjoev2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev2Withdraw represents a Withdraw event raised by the Masterchefjoev2 contract.
type Masterchefjoev2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev2WithdrawIterator{contract: _Masterchefjoev2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Masterchefjoev2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev2Withdraw)
				if err := _Masterchefjoev2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev2 *Masterchefjoev2Filterer) ParseWithdraw(log types.Log) (*Masterchefjoev2Withdraw, error) {
	event := new(Masterchefjoev2Withdraw)
	if err := _Masterchefjoev2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
