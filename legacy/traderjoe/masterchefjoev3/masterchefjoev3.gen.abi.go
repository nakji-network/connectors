// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masterchefjoev3

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

// Masterchefjoev3MetaData contains all meta data concerning the Masterchefjoev3 contract.
var Masterchefjoev3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"_MASTER_CHEF_V2\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_joe\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_MASTER_PID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"JOE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_CHEF_V2\",\"outputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFromMasterChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"joePerSec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"pids\",\"type\":\"uint256[]\"}],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingJoe\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bonusTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"bonusTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pendingBonusToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"accJoePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"rewarder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIRewarder\",\"name\":\"_rewarder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"overwrite\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Masterchefjoev3ABI is the input ABI used to generate the binding from.
// Deprecated: Use Masterchefjoev3MetaData.ABI instead.
var Masterchefjoev3ABI = Masterchefjoev3MetaData.ABI

// Masterchefjoev3 is an auto generated Go binding around an Ethereum contract.
type Masterchefjoev3 struct {
	Masterchefjoev3Caller     // Read-only binding to the contract
	Masterchefjoev3Transactor // Write-only binding to the contract
	Masterchefjoev3Filterer   // Log filterer for contract events
}

// Masterchefjoev3Caller is an auto generated read-only Go binding around an Ethereum contract.
type Masterchefjoev3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Masterchefjoev3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Masterchefjoev3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Masterchefjoev3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Masterchefjoev3Session struct {
	Contract     *Masterchefjoev3  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Masterchefjoev3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Masterchefjoev3CallerSession struct {
	Contract *Masterchefjoev3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Masterchefjoev3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Masterchefjoev3TransactorSession struct {
	Contract     *Masterchefjoev3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Masterchefjoev3Raw is an auto generated low-level Go binding around an Ethereum contract.
type Masterchefjoev3Raw struct {
	Contract *Masterchefjoev3 // Generic contract binding to access the raw methods on
}

// Masterchefjoev3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Masterchefjoev3CallerRaw struct {
	Contract *Masterchefjoev3Caller // Generic read-only contract binding to access the raw methods on
}

// Masterchefjoev3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Masterchefjoev3TransactorRaw struct {
	Contract *Masterchefjoev3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterchefjoev3 creates a new instance of Masterchefjoev3, bound to a specific deployed contract.
func NewMasterchefjoev3(address common.Address, backend bind.ContractBackend) (*Masterchefjoev3, error) {
	contract, err := bindMasterchefjoev3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3{Masterchefjoev3Caller: Masterchefjoev3Caller{contract: contract}, Masterchefjoev3Transactor: Masterchefjoev3Transactor{contract: contract}, Masterchefjoev3Filterer: Masterchefjoev3Filterer{contract: contract}}, nil
}

// NewMasterchefjoev3Caller creates a new read-only instance of Masterchefjoev3, bound to a specific deployed contract.
func NewMasterchefjoev3Caller(address common.Address, caller bind.ContractCaller) (*Masterchefjoev3Caller, error) {
	contract, err := bindMasterchefjoev3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3Caller{contract: contract}, nil
}

// NewMasterchefjoev3Transactor creates a new write-only instance of Masterchefjoev3, bound to a specific deployed contract.
func NewMasterchefjoev3Transactor(address common.Address, transactor bind.ContractTransactor) (*Masterchefjoev3Transactor, error) {
	contract, err := bindMasterchefjoev3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3Transactor{contract: contract}, nil
}

// NewMasterchefjoev3Filterer creates a new log filterer instance of Masterchefjoev3, bound to a specific deployed contract.
func NewMasterchefjoev3Filterer(address common.Address, filterer bind.ContractFilterer) (*Masterchefjoev3Filterer, error) {
	contract, err := bindMasterchefjoev3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3Filterer{contract: contract}, nil
}

// bindMasterchefjoev3 binds a generic wrapper to an already deployed contract.
func bindMasterchefjoev3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Masterchefjoev3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchefjoev3 *Masterchefjoev3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchefjoev3.Contract.Masterchefjoev3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchefjoev3 *Masterchefjoev3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Masterchefjoev3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchefjoev3 *Masterchefjoev3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Masterchefjoev3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Masterchefjoev3 *Masterchefjoev3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Masterchefjoev3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Masterchefjoev3 *Masterchefjoev3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Masterchefjoev3 *Masterchefjoev3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.contract.Transact(opts, method, params...)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Caller) JOE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "JOE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Session) JOE() (common.Address, error) {
	return _Masterchefjoev3.Contract.JOE(&_Masterchefjoev3.CallOpts)
}

// JOE is a free data retrieval call binding the contract method 0xffebad30.
//
// Solidity: function JOE() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) JOE() (common.Address, error) {
	return _Masterchefjoev3.Contract.JOE(&_Masterchefjoev3.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Caller) MASTERCHEFV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "MASTER_CHEF_V2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Session) MASTERCHEFV2() (common.Address, error) {
	return _Masterchefjoev3.Contract.MASTERCHEFV2(&_Masterchefjoev3.CallOpts)
}

// MASTERCHEFV2 is a free data retrieval call binding the contract method 0x27bf88ad.
//
// Solidity: function MASTER_CHEF_V2() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) MASTERCHEFV2() (common.Address, error) {
	return _Masterchefjoev3.Contract.MASTERCHEFV2(&_Masterchefjoev3.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3Caller) MASTERPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "MASTER_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3Session) MASTERPID() (*big.Int, error) {
	return _Masterchefjoev3.Contract.MASTERPID(&_Masterchefjoev3.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) MASTERPID() (*big.Int, error) {
	return _Masterchefjoev3.Contract.MASTERPID(&_Masterchefjoev3.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Caller) JoePerSec(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "joePerSec")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Session) JoePerSec() (*big.Int, error) {
	return _Masterchefjoev3.Contract.JoePerSec(&_Masterchefjoev3.CallOpts)
}

// JoePerSec is a free data retrieval call binding the contract method 0xca418d23.
//
// Solidity: function joePerSec() view returns(uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) JoePerSec() (*big.Int, error) {
	return _Masterchefjoev3.Contract.JoePerSec(&_Masterchefjoev3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3Session) Owner() (common.Address, error) {
	return _Masterchefjoev3.Contract.Owner(&_Masterchefjoev3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) Owner() (common.Address, error) {
	return _Masterchefjoev3.Contract.Owner(&_Masterchefjoev3.CallOpts)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_Masterchefjoev3 *Masterchefjoev3Caller) PendingTokens(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "pendingTokens", _pid, _user)

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
func (_Masterchefjoev3 *Masterchefjoev3Session) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _Masterchefjoev3.Contract.PendingTokens(&_Masterchefjoev3.CallOpts, _pid, _user)
}

// PendingTokens is a free data retrieval call binding the contract method 0xffcd4263.
//
// Solidity: function pendingTokens(uint256 _pid, address _user) view returns(uint256 pendingJoe, address bonusTokenAddress, string bonusTokenSymbol, uint256 pendingBonusToken)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) PendingTokens(_pid *big.Int, _user common.Address) (struct {
	PendingJoe        *big.Int
	BonusTokenAddress common.Address
	BonusTokenSymbol  string
	PendingBonusToken *big.Int
}, error) {
	return _Masterchefjoev3.Contract.PendingTokens(&_Masterchefjoev3.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_Masterchefjoev3 *Masterchefjoev3Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken             common.Address
		AccJoePerShare      *big.Int
		LastRewardTimestamp *big.Int
		AllocPoint          *big.Int
		Rewarder            common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AccJoePerShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Rewarder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_Masterchefjoev3 *Masterchefjoev3Session) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _Masterchefjoev3.Contract.PoolInfo(&_Masterchefjoev3.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 accJoePerShare, uint256 lastRewardTimestamp, uint256 allocPoint, address rewarder)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken             common.Address
	AccJoePerShare      *big.Int
	LastRewardTimestamp *big.Int
	AllocPoint          *big.Int
	Rewarder            common.Address
}, error) {
	return _Masterchefjoev3.Contract.PoolInfo(&_Masterchefjoev3.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_Masterchefjoev3 *Masterchefjoev3Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_Masterchefjoev3 *Masterchefjoev3Session) PoolLength() (*big.Int, error) {
	return _Masterchefjoev3.Contract.PoolLength(&_Masterchefjoev3.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) PoolLength() (*big.Int, error) {
	return _Masterchefjoev3.Contract.PoolLength(&_Masterchefjoev3.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3Caller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3Session) TotalAllocPoint() (*big.Int, error) {
	return _Masterchefjoev3.Contract.TotalAllocPoint(&_Masterchefjoev3.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Masterchefjoev3.Contract.TotalAllocPoint(&_Masterchefjoev3.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchefjoev3 *Masterchefjoev3Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Masterchefjoev3.contract.Call(opts, &out, "userInfo", arg0, arg1)

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
func (_Masterchefjoev3 *Masterchefjoev3Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchefjoev3.Contract.UserInfo(&_Masterchefjoev3.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Masterchefjoev3 *Masterchefjoev3CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Masterchefjoev3.Contract.UserInfo(&_Masterchefjoev3.CallOpts, arg0, arg1)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) Add(opts *bind.TransactOpts, allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "add", allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Add(&_Masterchefjoev3.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Add is a paid mutator transaction binding the contract method 0xab7de098.
//
// Solidity: function add(uint256 allocPoint, address _lpToken, address _rewarder) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) Add(allocPoint *big.Int, _lpToken common.Address, _rewarder common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Add(&_Masterchefjoev3.TransactOpts, allocPoint, _lpToken, _rewarder)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) Deposit(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "deposit", pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Deposit(&_Masterchefjoev3.TransactOpts, pid, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) Deposit(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Deposit(&_Masterchefjoev3.TransactOpts, pid, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) EmergencyWithdraw(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "emergencyWithdraw", pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) EmergencyWithdraw(pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.EmergencyWithdraw(&_Masterchefjoev3.TransactOpts, pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) EmergencyWithdraw(pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.EmergencyWithdraw(&_Masterchefjoev3.TransactOpts, pid)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) HarvestFromMasterChef(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "harvestFromMasterChef")
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) HarvestFromMasterChef() (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.HarvestFromMasterChef(&_Masterchefjoev3.TransactOpts)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.HarvestFromMasterChef(&_Masterchefjoev3.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Init(&_Masterchefjoev3.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Init(&_Masterchefjoev3.TransactOpts, dummyToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) MassUpdatePools(opts *bind.TransactOpts, pids []*big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "massUpdatePools", pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.MassUpdatePools(&_Masterchefjoev3.TransactOpts, pids)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x57a5b58c.
//
// Solidity: function massUpdatePools(uint256[] pids) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) MassUpdatePools(pids []*big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.MassUpdatePools(&_Masterchefjoev3.TransactOpts, pids)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.RenounceOwnership(&_Masterchefjoev3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.RenounceOwnership(&_Masterchefjoev3.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "set", _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Set(&_Masterchefjoev3.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// Set is a paid mutator transaction binding the contract method 0x88bba42f.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, address _rewarder, bool overwrite) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _rewarder common.Address, overwrite bool) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Set(&_Masterchefjoev3.TransactOpts, _pid, _allocPoint, _rewarder, overwrite)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.TransferOwnership(&_Masterchefjoev3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.TransferOwnership(&_Masterchefjoev3.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) UpdatePool(opts *bind.TransactOpts, pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "updatePool", pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.UpdatePool(&_Masterchefjoev3.TransactOpts, pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 pid) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) UpdatePool(pid *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.UpdatePool(&_Masterchefjoev3.TransactOpts, pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3Transactor) Withdraw(opts *bind.TransactOpts, pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.contract.Transact(opts, "withdraw", pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3Session) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Withdraw(&_Masterchefjoev3.TransactOpts, pid, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 pid, uint256 amount) returns()
func (_Masterchefjoev3 *Masterchefjoev3TransactorSession) Withdraw(pid *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Masterchefjoev3.Contract.Withdraw(&_Masterchefjoev3.TransactOpts, pid, amount)
}

// Masterchefjoev3AddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Masterchefjoev3 contract.
type Masterchefjoev3AddIterator struct {
	Event *Masterchefjoev3Add // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3AddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Add)
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
		it.Event = new(Masterchefjoev3Add)
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
func (it *Masterchefjoev3AddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3AddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Add represents a Add event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Add struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	Rewarder   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterAdd(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (*Masterchefjoev3AddIterator, error) {

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

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3AddIterator{contract: _Masterchefjoev3.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x4b16bd2431ad24dc020ab0e1de7fcb6563dead6a24fb10089d6c23e97a70381f.
//
// Solidity: event Add(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, address indexed rewarder)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Add, pid []*big.Int, lpToken []common.Address, rewarder []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Add", pidRule, lpTokenRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Add)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Add", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseAdd(log types.Log) (*Masterchefjoev3Add, error) {
	event := new(Masterchefjoev3Add)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Masterchefjoev3 contract.
type Masterchefjoev3DepositIterator struct {
	Event *Masterchefjoev3Deposit // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Deposit)
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
		it.Event = new(Masterchefjoev3Deposit)
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
func (it *Masterchefjoev3DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Deposit represents a Deposit event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev3DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3DepositIterator{contract: _Masterchefjoev3.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Deposit)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseDeposit(log types.Log) (*Masterchefjoev3Deposit, error) {
	event := new(Masterchefjoev3Deposit)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Masterchefjoev3 contract.
type Masterchefjoev3EmergencyWithdrawIterator struct {
	Event *Masterchefjoev3EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3EmergencyWithdraw)
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
		it.Event = new(Masterchefjoev3EmergencyWithdraw)
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
func (it *Masterchefjoev3EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3EmergencyWithdraw represents a EmergencyWithdraw event raised by the Masterchefjoev3 contract.
type Masterchefjoev3EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev3EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3EmergencyWithdrawIterator{contract: _Masterchefjoev3.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3EmergencyWithdraw)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseEmergencyWithdraw(log types.Log) (*Masterchefjoev3EmergencyWithdraw, error) {
	event := new(Masterchefjoev3EmergencyWithdraw)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3HarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Masterchefjoev3 contract.
type Masterchefjoev3HarvestIterator struct {
	Event *Masterchefjoev3Harvest // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3HarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Harvest)
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
		it.Event = new(Masterchefjoev3Harvest)
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
func (it *Masterchefjoev3HarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3HarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Harvest represents a Harvest event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Harvest struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterHarvest(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev3HarvestIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3HarvestIterator{contract: _Masterchefjoev3.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0x71bab65ced2e5750775a0613be067df48ef06cf92a496ebf7663ae0660924954.
//
// Solidity: event Harvest(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Harvest, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Harvest", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Harvest)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Harvest", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseHarvest(log types.Log) (*Masterchefjoev3Harvest, error) {
	event := new(Masterchefjoev3Harvest)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3InitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the Masterchefjoev3 contract.
type Masterchefjoev3InitIterator struct {
	Event *Masterchefjoev3Init // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3InitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Init)
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
		it.Event = new(Masterchefjoev3Init)
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
func (it *Masterchefjoev3InitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3InitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Init represents a Init event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Init struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterInit(opts *bind.FilterOpts) (*Masterchefjoev3InitIterator, error) {

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3InitIterator{contract: _Masterchefjoev3.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchInit(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Init) (event.Subscription, error) {

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Init)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseInit(log types.Log) (*Masterchefjoev3Init, error) {
	event := new(Masterchefjoev3Init)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Masterchefjoev3 contract.
type Masterchefjoev3OwnershipTransferredIterator struct {
	Event *Masterchefjoev3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3OwnershipTransferred)
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
		it.Event = new(Masterchefjoev3OwnershipTransferred)
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
func (it *Masterchefjoev3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3OwnershipTransferred represents a OwnershipTransferred event raised by the Masterchefjoev3 contract.
type Masterchefjoev3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Masterchefjoev3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3OwnershipTransferredIterator{contract: _Masterchefjoev3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3OwnershipTransferred)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseOwnershipTransferred(log types.Log) (*Masterchefjoev3OwnershipTransferred, error) {
	event := new(Masterchefjoev3OwnershipTransferred)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3SetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Masterchefjoev3 contract.
type Masterchefjoev3SetIterator struct {
	Event *Masterchefjoev3Set // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Set)
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
		it.Event = new(Masterchefjoev3Set)
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
func (it *Masterchefjoev3SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Set represents a Set event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Set struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Rewarder   common.Address
	Overwrite  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterSet(opts *bind.FilterOpts, pid []*big.Int, rewarder []common.Address) (*Masterchefjoev3SetIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3SetIterator{contract: _Masterchefjoev3.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0xa54644aae5c48c5971516f334e4fe8ecbc7930e23f34877d4203c6551e67ffaa.
//
// Solidity: event Set(uint256 indexed pid, uint256 allocPoint, address indexed rewarder, bool overwrite)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchSet(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Set, pid []*big.Int, rewarder []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var rewarderRule []interface{}
	for _, rewarderItem := range rewarder {
		rewarderRule = append(rewarderRule, rewarderItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Set", pidRule, rewarderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Set)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Set", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseSet(log types.Log) (*Masterchefjoev3Set, error) {
	event := new(Masterchefjoev3Set)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3UpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the Masterchefjoev3 contract.
type Masterchefjoev3UpdatePoolIterator struct {
	Event *Masterchefjoev3UpdatePool // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3UpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3UpdatePool)
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
		it.Event = new(Masterchefjoev3UpdatePool)
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
func (it *Masterchefjoev3UpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3UpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3UpdatePool represents a UpdatePool event raised by the Masterchefjoev3 contract.
type Masterchefjoev3UpdatePool struct {
	Pid                 *big.Int
	LastRewardTimestamp *big.Int
	LpSupply            *big.Int
	AccJoePerShare      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*Masterchefjoev3UpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3UpdatePoolIterator{contract: _Masterchefjoev3.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardTimestamp, uint256 lpSupply, uint256 accJoePerShare)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3UpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3UpdatePool)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseUpdatePool(log types.Log) (*Masterchefjoev3UpdatePool, error) {
	event := new(Masterchefjoev3UpdatePool)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Masterchefjoev3WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Masterchefjoev3 contract.
type Masterchefjoev3WithdrawIterator struct {
	Event *Masterchefjoev3Withdraw // Event containing the contract specifics and raw log

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
func (it *Masterchefjoev3WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Masterchefjoev3Withdraw)
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
		it.Event = new(Masterchefjoev3Withdraw)
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
func (it *Masterchefjoev3WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Masterchefjoev3WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Masterchefjoev3Withdraw represents a Withdraw event raised by the Masterchefjoev3 contract.
type Masterchefjoev3Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*Masterchefjoev3WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &Masterchefjoev3WithdrawIterator{contract: _Masterchefjoev3.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Masterchefjoev3 *Masterchefjoev3Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Masterchefjoev3Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Masterchefjoev3.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Masterchefjoev3Withdraw)
				if err := _Masterchefjoev3.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Masterchefjoev3 *Masterchefjoev3Filterer) ParseWithdraw(log types.Log) (*Masterchefjoev3Withdraw, error) {
	event := new(Masterchefjoev3Withdraw)
	if err := _Masterchefjoev3.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
