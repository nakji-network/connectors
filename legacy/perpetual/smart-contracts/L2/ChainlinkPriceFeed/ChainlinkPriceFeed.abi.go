// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ChainlinkPriceFeed

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

// ChainlinkPriceFeedABI is the input ABI used to generate the binding from.
const ChainlinkPriceFeedABI = "[{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addAggregator\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"_aggregator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"candidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractAggregatorV3Interface\"}],\"name\":\"getAggregator\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLatestTimestamp\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPreviousPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_numOfRoundBack\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPreviousTimestamp\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_numOfRoundBack\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getTwapPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_interval\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"priceFeedDecimalMap\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"priceFeedKeys\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractAggregatorV3Interface\"}],\"name\":\"priceFeedMap\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeAggregator\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLatestData\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_roundId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateOwner\",\"inputs\":[]}]"

// ChainlinkPriceFeed is an auto generated Go binding around an Ethereum contract.
type ChainlinkPriceFeed struct {
	ChainlinkPriceFeedCaller     // Read-only binding to the contract
	ChainlinkPriceFeedTransactor // Write-only binding to the contract
	ChainlinkPriceFeedFilterer   // Log filterer for contract events
}

// ChainlinkPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainlinkPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainlinkPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainlinkPriceFeedSession struct {
	Contract     *ChainlinkPriceFeed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChainlinkPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainlinkPriceFeedCallerSession struct {
	Contract *ChainlinkPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ChainlinkPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainlinkPriceFeedTransactorSession struct {
	Contract     *ChainlinkPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ChainlinkPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainlinkPriceFeedRaw struct {
	Contract *ChainlinkPriceFeed // Generic contract binding to access the raw methods on
}

// ChainlinkPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedCallerRaw struct {
	Contract *ChainlinkPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// ChainlinkPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainlinkPriceFeedTransactorRaw struct {
	Contract *ChainlinkPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainlinkPriceFeed creates a new instance of ChainlinkPriceFeed, bound to a specific deployed contract.
func NewChainlinkPriceFeed(address common.Address, backend bind.ContractBackend) (*ChainlinkPriceFeed, error) {
	contract, err := bindChainlinkPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeed{ChainlinkPriceFeedCaller: ChainlinkPriceFeedCaller{contract: contract}, ChainlinkPriceFeedTransactor: ChainlinkPriceFeedTransactor{contract: contract}, ChainlinkPriceFeedFilterer: ChainlinkPriceFeedFilterer{contract: contract}}, nil
}

// NewChainlinkPriceFeedCaller creates a new read-only instance of ChainlinkPriceFeed, bound to a specific deployed contract.
func NewChainlinkPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*ChainlinkPriceFeedCaller, error) {
	contract, err := bindChainlinkPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedCaller{contract: contract}, nil
}

// NewChainlinkPriceFeedTransactor creates a new write-only instance of ChainlinkPriceFeed, bound to a specific deployed contract.
func NewChainlinkPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainlinkPriceFeedTransactor, error) {
	contract, err := bindChainlinkPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedTransactor{contract: contract}, nil
}

// NewChainlinkPriceFeedFilterer creates a new log filterer instance of ChainlinkPriceFeed, bound to a specific deployed contract.
func NewChainlinkPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainlinkPriceFeedFilterer, error) {
	contract, err := bindChainlinkPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedFilterer{contract: contract}, nil
}

// bindChainlinkPriceFeed binds a generic wrapper to an already deployed contract.
func bindChainlinkPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainlinkPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPriceFeed.Contract.ChainlinkPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.ChainlinkPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.ChainlinkPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainlinkPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) Candidate() (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.Candidate(&_ChainlinkPriceFeed.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) Candidate() (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.Candidate(&_ChainlinkPriceFeed.CallOpts)
}

// GetAggregator is a free data retrieval call binding the contract method 0x331b1816.
//
// Solidity: function getAggregator(bytes32 _priceFeedKey) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetAggregator(opts *bind.CallOpts, _priceFeedKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getAggregator", _priceFeedKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAggregator is a free data retrieval call binding the contract method 0x331b1816.
//
// Solidity: function getAggregator(bytes32 _priceFeedKey) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetAggregator(_priceFeedKey [32]byte) (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.GetAggregator(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetAggregator is a free data retrieval call binding the contract method 0x331b1816.
//
// Solidity: function getAggregator(bytes32 _priceFeedKey) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetAggregator(_priceFeedKey [32]byte) (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.GetAggregator(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetLatestTimestamp(opts *bind.CallOpts, _priceFeedKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getLatestTimestamp", _priceFeedKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetLatestTimestamp(_priceFeedKey [32]byte) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetLatestTimestamp(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetLatestTimestamp(_priceFeedKey [32]byte) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetLatestTimestamp(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetPreviousPrice(opts *bind.CallOpts, _priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getPreviousPrice", _priceFeedKey, _numOfRoundBack)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetPreviousPrice(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPreviousPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetPreviousPrice(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPreviousPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetPreviousTimestamp(opts *bind.CallOpts, _priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getPreviousTimestamp", _priceFeedKey, _numOfRoundBack)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetPreviousTimestamp(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPreviousTimestamp(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetPreviousTimestamp(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPreviousTimestamp(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetPrice(opts *bind.CallOpts, _priceFeedKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getPrice", _priceFeedKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetPrice(_priceFeedKey [32]byte) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetPrice(_priceFeedKey [32]byte) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) GetTwapPrice(opts *bind.CallOpts, _priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "getTwapPrice", _priceFeedKey, _interval)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) GetTwapPrice(_priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetTwapPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _interval)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) GetTwapPrice(_priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	return _ChainlinkPriceFeed.Contract.GetTwapPrice(&_ChainlinkPriceFeed.CallOpts, _priceFeedKey, _interval)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) Owner() (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.Owner(&_ChainlinkPriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) Owner() (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.Owner(&_ChainlinkPriceFeed.CallOpts)
}

// PriceFeedDecimalMap is a free data retrieval call binding the contract method 0x6d4b8585.
//
// Solidity: function priceFeedDecimalMap(bytes32 ) view returns(uint8)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) PriceFeedDecimalMap(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "priceFeedDecimalMap", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PriceFeedDecimalMap is a free data retrieval call binding the contract method 0x6d4b8585.
//
// Solidity: function priceFeedDecimalMap(bytes32 ) view returns(uint8)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) PriceFeedDecimalMap(arg0 [32]byte) (uint8, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedDecimalMap(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// PriceFeedDecimalMap is a free data retrieval call binding the contract method 0x6d4b8585.
//
// Solidity: function priceFeedDecimalMap(bytes32 ) view returns(uint8)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) PriceFeedDecimalMap(arg0 [32]byte) (uint8, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedDecimalMap(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) PriceFeedKeys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "priceFeedKeys", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) PriceFeedKeys(arg0 *big.Int) ([32]byte, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedKeys(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) PriceFeedKeys(arg0 *big.Int) ([32]byte, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedKeys(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCaller) PriceFeedMap(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ChainlinkPriceFeed.contract.Call(opts, &out, "priceFeedMap", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) PriceFeedMap(arg0 [32]byte) (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedMap(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(address)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedCallerSession) PriceFeedMap(arg0 [32]byte) (common.Address, error) {
	return _ChainlinkPriceFeed.Contract.PriceFeedMap(&_ChainlinkPriceFeed.CallOpts, arg0)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x3f0e084f.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey, address _aggregator) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) AddAggregator(opts *bind.TransactOpts, _priceFeedKey [32]byte, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "addAggregator", _priceFeedKey, _aggregator)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x3f0e084f.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey, address _aggregator) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) AddAggregator(_priceFeedKey [32]byte, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.AddAggregator(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey, _aggregator)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x3f0e084f.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey, address _aggregator) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) AddAggregator(_priceFeedKey [32]byte, _aggregator common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.AddAggregator(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey, _aggregator)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) Initialize() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.Initialize(&_ChainlinkPriceFeed.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) Initialize() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.Initialize(&_ChainlinkPriceFeed.TransactOpts)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) RemoveAggregator(opts *bind.TransactOpts, _priceFeedKey [32]byte) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "removeAggregator", _priceFeedKey)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) RemoveAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.RemoveAggregator(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) RemoveAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.RemoveAggregator(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.RenounceOwnership(&_ChainlinkPriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.RenounceOwnership(&_ChainlinkPriceFeed.TransactOpts)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) SetLatestData(opts *bind.TransactOpts, _priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "setLatestData", _priceFeedKey, _price, _timestamp, _roundId)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) SetLatestData(_priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.SetLatestData(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey, _price, _timestamp, _roundId)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) SetLatestData(_priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.SetLatestData(&_ChainlinkPriceFeed.TransactOpts, _priceFeedKey, _price, _timestamp, _roundId)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.SetOwner(&_ChainlinkPriceFeed.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.SetOwner(&_ChainlinkPriceFeed.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainlinkPriceFeed.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedSession) UpdateOwner() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.UpdateOwner(&_ChainlinkPriceFeed.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_ChainlinkPriceFeed *ChainlinkPriceFeedTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _ChainlinkPriceFeed.Contract.UpdateOwner(&_ChainlinkPriceFeed.TransactOpts)
}

// ChainlinkPriceFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChainlinkPriceFeed contract.
type ChainlinkPriceFeedOwnershipTransferredIterator struct {
	Event *ChainlinkPriceFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChainlinkPriceFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainlinkPriceFeedOwnershipTransferred)
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
		it.Event = new(ChainlinkPriceFeedOwnershipTransferred)
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
func (it *ChainlinkPriceFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainlinkPriceFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainlinkPriceFeedOwnershipTransferred represents a OwnershipTransferred event raised by the ChainlinkPriceFeed contract.
type ChainlinkPriceFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChainlinkPriceFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkPriceFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChainlinkPriceFeedOwnershipTransferredIterator{contract: _ChainlinkPriceFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ChainlinkPriceFeed *ChainlinkPriceFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChainlinkPriceFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChainlinkPriceFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainlinkPriceFeedOwnershipTransferred)
				if err := _ChainlinkPriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ChainlinkPriceFeed *ChainlinkPriceFeedFilterer) ParseOwnershipTransferred(log types.Log) (*ChainlinkPriceFeedOwnershipTransferred, error) {
	event := new(ChainlinkPriceFeedOwnershipTransferred)
	if err := _ChainlinkPriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
