// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AmmReader

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

// Decimaldecimal is an auto generated low-level Go binding around an user-defined struct.
type Decimaldecimal struct {
	D *big.Int
}

// AmmReaderABI is the input ABI used to generate the binding from.
const AmmReaderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"candidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.decimal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTokenPoolDispatcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractStakedPerpToken\",\"name\":\"_stakedPerpToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeTokenPoolDispatcher\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.decimal\",\"name\":\"_reward\",\"type\":\"tuple\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"notifyStakeChanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"internalType\":\"structDecimal.decimal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedPerpToken\",\"outputs\":[{\"internalType\":\"contractStakedPerpToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"d\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AmmReader is an auto generated Go binding around an Ethereum contract.
type AmmReader struct {
	AmmReaderCaller     // Read-only binding to the contract
	AmmReaderTransactor // Write-only binding to the contract
	AmmReaderFilterer   // Log filterer for contract events
}

// AmmReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmmReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmmReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmmReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmmReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmmReaderSession struct {
	Contract     *AmmReader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AmmReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmmReaderCallerSession struct {
	Contract *AmmReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AmmReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmmReaderTransactorSession struct {
	Contract     *AmmReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AmmReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmmReaderRaw struct {
	Contract *AmmReader // Generic contract binding to access the raw methods on
}

// AmmReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmmReaderCallerRaw struct {
	Contract *AmmReaderCaller // Generic read-only contract binding to access the raw methods on
}

// AmmReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmmReaderTransactorRaw struct {
	Contract *AmmReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmmReader creates a new instance of AmmReader, bound to a specific deployed contract.
func NewAmmReader(address common.Address, backend bind.ContractBackend) (*AmmReader, error) {
	contract, err := bindAmmReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmmReader{AmmReaderCaller: AmmReaderCaller{contract: contract}, AmmReaderTransactor: AmmReaderTransactor{contract: contract}, AmmReaderFilterer: AmmReaderFilterer{contract: contract}}, nil
}

// NewAmmReaderCaller creates a new read-only instance of AmmReader, bound to a specific deployed contract.
func NewAmmReaderCaller(address common.Address, caller bind.ContractCaller) (*AmmReaderCaller, error) {
	contract, err := bindAmmReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmmReaderCaller{contract: contract}, nil
}

// NewAmmReaderTransactor creates a new write-only instance of AmmReader, bound to a specific deployed contract.
func NewAmmReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*AmmReaderTransactor, error) {
	contract, err := bindAmmReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmmReaderTransactor{contract: contract}, nil
}

// NewAmmReaderFilterer creates a new log filterer instance of AmmReader, bound to a specific deployed contract.
func NewAmmReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*AmmReaderFilterer, error) {
	contract, err := bindAmmReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmmReaderFilterer{contract: contract}, nil
}

// bindAmmReader binds a generic wrapper to an already deployed contract.
func bindAmmReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmmReaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmmReader *AmmReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmmReader.Contract.AmmReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmmReader *AmmReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmReader.Contract.AmmReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmmReader *AmmReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmmReader.Contract.AmmReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmmReader *AmmReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmmReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmmReader *AmmReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmmReader *AmmReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmmReader.Contract.contract.Transact(opts, method, params...)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_AmmReader *AmmReaderCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_AmmReader *AmmReaderSession) DURATION() (*big.Int, error) {
	return _AmmReader.Contract.DURATION(&_AmmReader.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_AmmReader *AmmReaderCallerSession) DURATION() (*big.Int, error) {
	return _AmmReader.Contract.DURATION(&_AmmReader.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_AmmReader *AmmReaderCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_AmmReader *AmmReaderSession) Candidate() (common.Address, error) {
	return _AmmReader.Contract.Candidate(&_AmmReader.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_AmmReader *AmmReaderCallerSession) Candidate() (common.Address, error) {
	return _AmmReader.Contract.Candidate(&_AmmReader.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _staker) view returns((uint256))
func (_AmmReader *AmmReaderCaller) Earned(opts *bind.CallOpts, _staker common.Address) (Decimaldecimal, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "earned", _staker)

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _staker) view returns((uint256))
func (_AmmReader *AmmReaderSession) Earned(_staker common.Address) (Decimaldecimal, error) {
	return _AmmReader.Contract.Earned(&_AmmReader.CallOpts, _staker)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _staker) view returns((uint256))
func (_AmmReader *AmmReaderCallerSession) Earned(_staker common.Address) (Decimaldecimal, error) {
	return _AmmReader.Contract.Earned(&_AmmReader.CallOpts, _staker)
}

// FeeTokenPoolDispatcher is a free data retrieval call binding the contract method 0xf6aa49e7.
//
// Solidity: function feeTokenPoolDispatcher() view returns(address)
func (_AmmReader *AmmReaderCaller) FeeTokenPoolDispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "feeTokenPoolDispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTokenPoolDispatcher is a free data retrieval call binding the contract method 0xf6aa49e7.
//
// Solidity: function feeTokenPoolDispatcher() view returns(address)
func (_AmmReader *AmmReaderSession) FeeTokenPoolDispatcher() (common.Address, error) {
	return _AmmReader.Contract.FeeTokenPoolDispatcher(&_AmmReader.CallOpts)
}

// FeeTokenPoolDispatcher is a free data retrieval call binding the contract method 0xf6aa49e7.
//
// Solidity: function feeTokenPoolDispatcher() view returns(address)
func (_AmmReader *AmmReaderCallerSession) FeeTokenPoolDispatcher() (common.Address, error) {
	return _AmmReader.Contract.FeeTokenPoolDispatcher(&_AmmReader.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_AmmReader *AmmReaderCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_AmmReader *AmmReaderSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _AmmReader.Contract.LastTimeRewardApplicable(&_AmmReader.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_AmmReader *AmmReaderCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _AmmReader.Contract.LastTimeRewardApplicable(&_AmmReader.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_AmmReader *AmmReaderCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_AmmReader *AmmReaderSession) LastUpdateTime() (*big.Int, error) {
	return _AmmReader.Contract.LastUpdateTime(&_AmmReader.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_AmmReader *AmmReaderCallerSession) LastUpdateTime() (*big.Int, error) {
	return _AmmReader.Contract.LastUpdateTime(&_AmmReader.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AmmReader *AmmReaderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AmmReader *AmmReaderSession) Owner() (common.Address, error) {
	return _AmmReader.Contract.Owner(&_AmmReader.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AmmReader *AmmReaderCallerSession) Owner() (common.Address, error) {
	return _AmmReader.Contract.Owner(&_AmmReader.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_AmmReader *AmmReaderCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_AmmReader *AmmReaderSession) PeriodFinish() (*big.Int, error) {
	return _AmmReader.Contract.PeriodFinish(&_AmmReader.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_AmmReader *AmmReaderCallerSession) PeriodFinish() (*big.Int, error) {
	return _AmmReader.Contract.PeriodFinish(&_AmmReader.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns((uint256))
func (_AmmReader *AmmReaderCaller) RewardPerToken(opts *bind.CallOpts) (Decimaldecimal, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(Decimaldecimal), err
	}

	out0 := *abi.ConvertType(out[0], new(Decimaldecimal)).(*Decimaldecimal)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns((uint256))
func (_AmmReader *AmmReaderSession) RewardPerToken() (Decimaldecimal, error) {
	return _AmmReader.Contract.RewardPerToken(&_AmmReader.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns((uint256))
func (_AmmReader *AmmReaderCallerSession) RewardPerToken() (Decimaldecimal, error) {
	return _AmmReader.Contract.RewardPerToken(&_AmmReader.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256 d)
func (_AmmReader *AmmReaderCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256 d)
func (_AmmReader *AmmReaderSession) RewardPerTokenStored() (*big.Int, error) {
	return _AmmReader.Contract.RewardPerTokenStored(&_AmmReader.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256 d)
func (_AmmReader *AmmReaderCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _AmmReader.Contract.RewardPerTokenStored(&_AmmReader.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256 d)
func (_AmmReader *AmmReaderCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256 d)
func (_AmmReader *AmmReaderSession) RewardRate() (*big.Int, error) {
	return _AmmReader.Contract.RewardRate(&_AmmReader.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256 d)
func (_AmmReader *AmmReaderCallerSession) RewardRate() (*big.Int, error) {
	return _AmmReader.Contract.RewardRate(&_AmmReader.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AmmReader.Contract.Rewards(&_AmmReader.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _AmmReader.Contract.Rewards(&_AmmReader.CallOpts, arg0)
}

// StakedPerpToken is a free data retrieval call binding the contract method 0x4cd7ee95.
//
// Solidity: function stakedPerpToken() view returns(address)
func (_AmmReader *AmmReaderCaller) StakedPerpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "stakedPerpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedPerpToken is a free data retrieval call binding the contract method 0x4cd7ee95.
//
// Solidity: function stakedPerpToken() view returns(address)
func (_AmmReader *AmmReaderSession) StakedPerpToken() (common.Address, error) {
	return _AmmReader.Contract.StakedPerpToken(&_AmmReader.CallOpts)
}

// StakedPerpToken is a free data retrieval call binding the contract method 0x4cd7ee95.
//
// Solidity: function stakedPerpToken() view returns(address)
func (_AmmReader *AmmReaderCallerSession) StakedPerpToken() (common.Address, error) {
	return _AmmReader.Contract.StakedPerpToken(&_AmmReader.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AmmReader *AmmReaderCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AmmReader *AmmReaderSession) Token() (common.Address, error) {
	return _AmmReader.Contract.Token(&_AmmReader.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AmmReader *AmmReaderCallerSession) Token() (common.Address, error) {
	return _AmmReader.Contract.Token(&_AmmReader.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AmmReader.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _AmmReader.Contract.UserRewardPerTokenPaid(&_AmmReader.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256 d)
func (_AmmReader *AmmReaderCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _AmmReader.Contract.UserRewardPerTokenPaid(&_AmmReader.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _stakedPerpToken, address _feeTokenPoolDispatcher) returns()
func (_AmmReader *AmmReaderTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _stakedPerpToken common.Address, _feeTokenPoolDispatcher common.Address) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "initialize", _token, _stakedPerpToken, _feeTokenPoolDispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _stakedPerpToken, address _feeTokenPoolDispatcher) returns()
func (_AmmReader *AmmReaderSession) Initialize(_token common.Address, _stakedPerpToken common.Address, _feeTokenPoolDispatcher common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.Initialize(&_AmmReader.TransactOpts, _token, _stakedPerpToken, _feeTokenPoolDispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _token, address _stakedPerpToken, address _feeTokenPoolDispatcher) returns()
func (_AmmReader *AmmReaderTransactorSession) Initialize(_token common.Address, _stakedPerpToken common.Address, _feeTokenPoolDispatcher common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.Initialize(&_AmmReader.TransactOpts, _token, _stakedPerpToken, _feeTokenPoolDispatcher)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xff59f86d.
//
// Solidity: function notifyRewardAmount((uint256) _reward) returns()
func (_AmmReader *AmmReaderTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _reward Decimaldecimal) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "notifyRewardAmount", _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xff59f86d.
//
// Solidity: function notifyRewardAmount((uint256) _reward) returns()
func (_AmmReader *AmmReaderSession) NotifyRewardAmount(_reward Decimaldecimal) (*types.Transaction, error) {
	return _AmmReader.Contract.NotifyRewardAmount(&_AmmReader.TransactOpts, _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xff59f86d.
//
// Solidity: function notifyRewardAmount((uint256) _reward) returns()
func (_AmmReader *AmmReaderTransactorSession) NotifyRewardAmount(_reward Decimaldecimal) (*types.Transaction, error) {
	return _AmmReader.Contract.NotifyRewardAmount(&_AmmReader.TransactOpts, _reward)
}

// NotifyStakeChanged is a paid mutator transaction binding the contract method 0xcd8ae3dd.
//
// Solidity: function notifyStakeChanged(address _staker) returns()
func (_AmmReader *AmmReaderTransactor) NotifyStakeChanged(opts *bind.TransactOpts, _staker common.Address) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "notifyStakeChanged", _staker)
}

// NotifyStakeChanged is a paid mutator transaction binding the contract method 0xcd8ae3dd.
//
// Solidity: function notifyStakeChanged(address _staker) returns()
func (_AmmReader *AmmReaderSession) NotifyStakeChanged(_staker common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.NotifyStakeChanged(&_AmmReader.TransactOpts, _staker)
}

// NotifyStakeChanged is a paid mutator transaction binding the contract method 0xcd8ae3dd.
//
// Solidity: function notifyStakeChanged(address _staker) returns()
func (_AmmReader *AmmReaderTransactorSession) NotifyStakeChanged(_staker common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.NotifyStakeChanged(&_AmmReader.TransactOpts, _staker)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AmmReader *AmmReaderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AmmReader *AmmReaderSession) RenounceOwnership() (*types.Transaction, error) {
	return _AmmReader.Contract.RenounceOwnership(&_AmmReader.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AmmReader *AmmReaderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AmmReader.Contract.RenounceOwnership(&_AmmReader.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_AmmReader *AmmReaderTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_AmmReader *AmmReaderSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.SetOwner(&_AmmReader.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_AmmReader *AmmReaderTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _AmmReader.Contract.SetOwner(&_AmmReader.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_AmmReader *AmmReaderTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_AmmReader *AmmReaderSession) UpdateOwner() (*types.Transaction, error) {
	return _AmmReader.Contract.UpdateOwner(&_AmmReader.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_AmmReader *AmmReaderTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _AmmReader.Contract.UpdateOwner(&_AmmReader.TransactOpts)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0xc885bc58.
//
// Solidity: function withdrawReward() returns()
func (_AmmReader *AmmReaderTransactor) WithdrawReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmmReader.contract.Transact(opts, "withdrawReward")
}

// WithdrawReward is a paid mutator transaction binding the contract method 0xc885bc58.
//
// Solidity: function withdrawReward() returns()
func (_AmmReader *AmmReaderSession) WithdrawReward() (*types.Transaction, error) {
	return _AmmReader.Contract.WithdrawReward(&_AmmReader.TransactOpts)
}

// WithdrawReward is a paid mutator transaction binding the contract method 0xc885bc58.
//
// Solidity: function withdrawReward() returns()
func (_AmmReader *AmmReaderTransactorSession) WithdrawReward() (*types.Transaction, error) {
	return _AmmReader.Contract.WithdrawReward(&_AmmReader.TransactOpts)
}

// AmmReaderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AmmReader contract.
type AmmReaderOwnershipTransferredIterator struct {
	Event *AmmReaderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AmmReaderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmReaderOwnershipTransferred)
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
		it.Event = new(AmmReaderOwnershipTransferred)
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
func (it *AmmReaderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmReaderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmReaderOwnershipTransferred represents a OwnershipTransferred event raised by the AmmReader contract.
type AmmReaderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AmmReader *AmmReaderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AmmReaderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AmmReader.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AmmReaderOwnershipTransferredIterator{contract: _AmmReader.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AmmReader *AmmReaderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AmmReaderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AmmReader.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmReaderOwnershipTransferred)
				if err := _AmmReader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AmmReader *AmmReaderFilterer) ParseOwnershipTransferred(log types.Log) (*AmmReaderOwnershipTransferred, error) {
	event := new(AmmReaderOwnershipTransferred)
	if err := _AmmReader.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmReaderRewardTransferredIterator is returned from FilterRewardTransferred and is used to iterate over the raw logs and unpacked data for RewardTransferred events raised by the AmmReader contract.
type AmmReaderRewardTransferredIterator struct {
	Event *AmmReaderRewardTransferred // Event containing the contract specifics and raw log

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
func (it *AmmReaderRewardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmReaderRewardTransferred)
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
		it.Event = new(AmmReaderRewardTransferred)
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
func (it *AmmReaderRewardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmReaderRewardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmReaderRewardTransferred represents a RewardTransferred event raised by the AmmReader contract.
type AmmReaderRewardTransferred struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTransferred is a free log retrieval operation binding the contract event 0xe7b1adc654f4988d82260a00603bcc1f508dc3aa94b2b9cd5bc9ab42cfc591dc.
//
// Solidity: event RewardTransferred(uint256 amount)
func (_AmmReader *AmmReaderFilterer) FilterRewardTransferred(opts *bind.FilterOpts) (*AmmReaderRewardTransferredIterator, error) {

	logs, sub, err := _AmmReader.contract.FilterLogs(opts, "RewardTransferred")
	if err != nil {
		return nil, err
	}
	return &AmmReaderRewardTransferredIterator{contract: _AmmReader.contract, event: "RewardTransferred", logs: logs, sub: sub}, nil
}

// WatchRewardTransferred is a free log subscription operation binding the contract event 0xe7b1adc654f4988d82260a00603bcc1f508dc3aa94b2b9cd5bc9ab42cfc591dc.
//
// Solidity: event RewardTransferred(uint256 amount)
func (_AmmReader *AmmReaderFilterer) WatchRewardTransferred(opts *bind.WatchOpts, sink chan<- *AmmReaderRewardTransferred) (event.Subscription, error) {

	logs, sub, err := _AmmReader.contract.WatchLogs(opts, "RewardTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmReaderRewardTransferred)
				if err := _AmmReader.contract.UnpackLog(event, "RewardTransferred", log); err != nil {
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

// ParseRewardTransferred is a log parse operation binding the contract event 0xe7b1adc654f4988d82260a00603bcc1f508dc3aa94b2b9cd5bc9ab42cfc591dc.
//
// Solidity: event RewardTransferred(uint256 amount)
func (_AmmReader *AmmReaderFilterer) ParseRewardTransferred(log types.Log) (*AmmReaderRewardTransferred, error) {
	event := new(AmmReaderRewardTransferred)
	if err := _AmmReader.contract.UnpackLog(event, "RewardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmmReaderRewardWithdrawnIterator is returned from FilterRewardWithdrawn and is used to iterate over the raw logs and unpacked data for RewardWithdrawn events raised by the AmmReader contract.
type AmmReaderRewardWithdrawnIterator struct {
	Event *AmmReaderRewardWithdrawn // Event containing the contract specifics and raw log

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
func (it *AmmReaderRewardWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmmReaderRewardWithdrawn)
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
		it.Event = new(AmmReaderRewardWithdrawn)
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
func (it *AmmReaderRewardWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmmReaderRewardWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmmReaderRewardWithdrawn represents a RewardWithdrawn event raised by the AmmReader contract.
type AmmReaderRewardWithdrawn struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardWithdrawn is a free log retrieval operation binding the contract event 0x1d3eee4ca001cff39eec6ec7615aacf2f2bd61791273830728ba00ccbd6e1337.
//
// Solidity: event RewardWithdrawn(address indexed staker, uint256 amount)
func (_AmmReader *AmmReaderFilterer) FilterRewardWithdrawn(opts *bind.FilterOpts, staker []common.Address) (*AmmReaderRewardWithdrawnIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AmmReader.contract.FilterLogs(opts, "RewardWithdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return &AmmReaderRewardWithdrawnIterator{contract: _AmmReader.contract, event: "RewardWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardWithdrawn is a free log subscription operation binding the contract event 0x1d3eee4ca001cff39eec6ec7615aacf2f2bd61791273830728ba00ccbd6e1337.
//
// Solidity: event RewardWithdrawn(address indexed staker, uint256 amount)
func (_AmmReader *AmmReaderFilterer) WatchRewardWithdrawn(opts *bind.WatchOpts, sink chan<- *AmmReaderRewardWithdrawn, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _AmmReader.contract.WatchLogs(opts, "RewardWithdrawn", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmmReaderRewardWithdrawn)
				if err := _AmmReader.contract.UnpackLog(event, "RewardWithdrawn", log); err != nil {
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

// ParseRewardWithdrawn is a log parse operation binding the contract event 0x1d3eee4ca001cff39eec6ec7615aacf2f2bd61791273830728ba00ccbd6e1337.
//
// Solidity: event RewardWithdrawn(address indexed staker, uint256 amount)
func (_AmmReader *AmmReaderFilterer) ParseRewardWithdrawn(log types.Log) (*AmmReaderRewardWithdrawn, error) {
	event := new(AmmReaderRewardWithdrawn)
	if err := _AmmReader.contract.UnpackLog(event, "RewardWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
