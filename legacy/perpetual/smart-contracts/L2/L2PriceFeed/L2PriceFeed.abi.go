// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L2PriceFeed

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

// L2PriceFeedABI is the input ABI used to generate the binding from.
const L2PriceFeedABI = "[{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedDataSet\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"key\",\"internalType\":\"bytes32\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"timestamp\",\"internalType\":\"uint256\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"roundId\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addAggregator\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ambBridge\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"candidate\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLatestTimestamp\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPreviousPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_numOfRoundBack\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPreviousTimestamp\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_numOfRoundBack\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"length\",\"internalType\":\"uint256\"}],\"name\":\"getPriceFeedLength\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getTwapPrice\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_interval\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_ambBridge\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_rootBridge\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"priceFeedKeys\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"registered\",\"internalType\":\"bool\"}],\"name\":\"priceFeedMap\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeAggregator\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"rootBridge\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setLatestData\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"_priceFeedKey\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"_price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_timestamp\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_roundId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setRootBridge\",\"inputs\":[{\"type\":\"address\",\"name\":\"_rootBridge\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateOwner\",\"inputs\":[]}]"

// L2PriceFeed is an auto generated Go binding around an Ethereum contract.
type L2PriceFeed struct {
	L2PriceFeedCaller     // Read-only binding to the contract
	L2PriceFeedTransactor // Write-only binding to the contract
	L2PriceFeedFilterer   // Log filterer for contract events
}

// L2PriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2PriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2PriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2PriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2PriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2PriceFeedSession struct {
	Contract     *L2PriceFeed      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2PriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2PriceFeedCallerSession struct {
	Contract *L2PriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// L2PriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2PriceFeedTransactorSession struct {
	Contract     *L2PriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// L2PriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2PriceFeedRaw struct {
	Contract *L2PriceFeed // Generic contract binding to access the raw methods on
}

// L2PriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2PriceFeedCallerRaw struct {
	Contract *L2PriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// L2PriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2PriceFeedTransactorRaw struct {
	Contract *L2PriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2PriceFeed creates a new instance of L2PriceFeed, bound to a specific deployed contract.
func NewL2PriceFeed(address common.Address, backend bind.ContractBackend) (*L2PriceFeed, error) {
	contract, err := bindL2PriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2PriceFeed{L2PriceFeedCaller: L2PriceFeedCaller{contract: contract}, L2PriceFeedTransactor: L2PriceFeedTransactor{contract: contract}, L2PriceFeedFilterer: L2PriceFeedFilterer{contract: contract}}, nil
}

// NewL2PriceFeedCaller creates a new read-only instance of L2PriceFeed, bound to a specific deployed contract.
func NewL2PriceFeedCaller(address common.Address, caller bind.ContractCaller) (*L2PriceFeedCaller, error) {
	contract, err := bindL2PriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2PriceFeedCaller{contract: contract}, nil
}

// NewL2PriceFeedTransactor creates a new write-only instance of L2PriceFeed, bound to a specific deployed contract.
func NewL2PriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*L2PriceFeedTransactor, error) {
	contract, err := bindL2PriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2PriceFeedTransactor{contract: contract}, nil
}

// NewL2PriceFeedFilterer creates a new log filterer instance of L2PriceFeed, bound to a specific deployed contract.
func NewL2PriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*L2PriceFeedFilterer, error) {
	contract, err := bindL2PriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2PriceFeedFilterer{contract: contract}, nil
}

// bindL2PriceFeed binds a generic wrapper to an already deployed contract.
func bindL2PriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2PriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PriceFeed *L2PriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PriceFeed.Contract.L2PriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PriceFeed *L2PriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.L2PriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PriceFeed *L2PriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.L2PriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2PriceFeed *L2PriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2PriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2PriceFeed *L2PriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2PriceFeed *L2PriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.contract.Transact(opts, method, params...)
}

// AmbBridge is a free data retrieval call binding the contract method 0x0ba54df6.
//
// Solidity: function ambBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedCaller) AmbBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "ambBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmbBridge is a free data retrieval call binding the contract method 0x0ba54df6.
//
// Solidity: function ambBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedSession) AmbBridge() (common.Address, error) {
	return _L2PriceFeed.Contract.AmbBridge(&_L2PriceFeed.CallOpts)
}

// AmbBridge is a free data retrieval call binding the contract method 0x0ba54df6.
//
// Solidity: function ambBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedCallerSession) AmbBridge() (common.Address, error) {
	return _L2PriceFeed.Contract.AmbBridge(&_L2PriceFeed.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_L2PriceFeed *L2PriceFeedCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_L2PriceFeed *L2PriceFeedSession) Candidate() (common.Address, error) {
	return _L2PriceFeed.Contract.Candidate(&_L2PriceFeed.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_L2PriceFeed *L2PriceFeedCallerSession) Candidate() (common.Address, error) {
	return _L2PriceFeed.Contract.Candidate(&_L2PriceFeed.CallOpts)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCaller) GetLatestTimestamp(opts *bind.CallOpts, _priceFeedKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getLatestTimestamp", _priceFeedKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedSession) GetLatestTimestamp(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetLatestTimestamp(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetLatestTimestamp is a free data retrieval call binding the contract method 0x7ad22632.
//
// Solidity: function getLatestTimestamp(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetLatestTimestamp(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetLatestTimestamp(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCaller) GetPreviousPrice(opts *bind.CallOpts, _priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getPreviousPrice", _priceFeedKey, _numOfRoundBack)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedSession) GetPreviousPrice(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPreviousPrice(&_L2PriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousPrice is a free data retrieval call binding the contract method 0x914b14c5.
//
// Solidity: function getPreviousPrice(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetPreviousPrice(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPreviousPrice(&_L2PriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCaller) GetPreviousTimestamp(opts *bind.CallOpts, _priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getPreviousTimestamp", _priceFeedKey, _numOfRoundBack)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedSession) GetPreviousTimestamp(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPreviousTimestamp(&_L2PriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPreviousTimestamp is a free data retrieval call binding the contract method 0x59bb180e.
//
// Solidity: function getPreviousTimestamp(bytes32 _priceFeedKey, uint256 _numOfRoundBack) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetPreviousTimestamp(_priceFeedKey [32]byte, _numOfRoundBack *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPreviousTimestamp(&_L2PriceFeed.CallOpts, _priceFeedKey, _numOfRoundBack)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCaller) GetPrice(opts *bind.CallOpts, _priceFeedKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getPrice", _priceFeedKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedSession) GetPrice(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPrice(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetPrice is a free data retrieval call binding the contract method 0x31d98b3f.
//
// Solidity: function getPrice(bytes32 _priceFeedKey) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetPrice(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPrice(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetPriceFeedLength is a free data retrieval call binding the contract method 0x744aca14.
//
// Solidity: function getPriceFeedLength(bytes32 _priceFeedKey) view returns(uint256 length)
func (_L2PriceFeed *L2PriceFeedCaller) GetPriceFeedLength(opts *bind.CallOpts, _priceFeedKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getPriceFeedLength", _priceFeedKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFeedLength is a free data retrieval call binding the contract method 0x744aca14.
//
// Solidity: function getPriceFeedLength(bytes32 _priceFeedKey) view returns(uint256 length)
func (_L2PriceFeed *L2PriceFeedSession) GetPriceFeedLength(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPriceFeedLength(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetPriceFeedLength is a free data retrieval call binding the contract method 0x744aca14.
//
// Solidity: function getPriceFeedLength(bytes32 _priceFeedKey) view returns(uint256 length)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetPriceFeedLength(_priceFeedKey [32]byte) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetPriceFeedLength(&_L2PriceFeed.CallOpts, _priceFeedKey)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCaller) GetTwapPrice(opts *bind.CallOpts, _priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "getTwapPrice", _priceFeedKey, _interval)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedSession) GetTwapPrice(_priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetTwapPrice(&_L2PriceFeed.CallOpts, _priceFeedKey, _interval)
}

// GetTwapPrice is a free data retrieval call binding the contract method 0xa2173df4.
//
// Solidity: function getTwapPrice(bytes32 _priceFeedKey, uint256 _interval) view returns(uint256)
func (_L2PriceFeed *L2PriceFeedCallerSession) GetTwapPrice(_priceFeedKey [32]byte, _interval *big.Int) (*big.Int, error) {
	return _L2PriceFeed.Contract.GetTwapPrice(&_L2PriceFeed.CallOpts, _priceFeedKey, _interval)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2PriceFeed *L2PriceFeedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2PriceFeed *L2PriceFeedSession) Owner() (common.Address, error) {
	return _L2PriceFeed.Contract.Owner(&_L2PriceFeed.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2PriceFeed *L2PriceFeedCallerSession) Owner() (common.Address, error) {
	return _L2PriceFeed.Contract.Owner(&_L2PriceFeed.CallOpts)
}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_L2PriceFeed *L2PriceFeedCaller) PriceFeedKeys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "priceFeedKeys", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_L2PriceFeed *L2PriceFeedSession) PriceFeedKeys(arg0 *big.Int) ([32]byte, error) {
	return _L2PriceFeed.Contract.PriceFeedKeys(&_L2PriceFeed.CallOpts, arg0)
}

// PriceFeedKeys is a free data retrieval call binding the contract method 0x2a0ab1dd.
//
// Solidity: function priceFeedKeys(uint256 ) view returns(bytes32)
func (_L2PriceFeed *L2PriceFeedCallerSession) PriceFeedKeys(arg0 *big.Int) ([32]byte, error) {
	return _L2PriceFeed.Contract.PriceFeedKeys(&_L2PriceFeed.CallOpts, arg0)
}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(bool registered)
func (_L2PriceFeed *L2PriceFeedCaller) PriceFeedMap(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "priceFeedMap", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(bool registered)
func (_L2PriceFeed *L2PriceFeedSession) PriceFeedMap(arg0 [32]byte) (bool, error) {
	return _L2PriceFeed.Contract.PriceFeedMap(&_L2PriceFeed.CallOpts, arg0)
}

// PriceFeedMap is a free data retrieval call binding the contract method 0x250742cc.
//
// Solidity: function priceFeedMap(bytes32 ) view returns(bool registered)
func (_L2PriceFeed *L2PriceFeedCallerSession) PriceFeedMap(arg0 [32]byte) (bool, error) {
	return _L2PriceFeed.Contract.PriceFeedMap(&_L2PriceFeed.CallOpts, arg0)
}

// RootBridge is a free data retrieval call binding the contract method 0xcdea7c70.
//
// Solidity: function rootBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedCaller) RootBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2PriceFeed.contract.Call(opts, &out, "rootBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootBridge is a free data retrieval call binding the contract method 0xcdea7c70.
//
// Solidity: function rootBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedSession) RootBridge() (common.Address, error) {
	return _L2PriceFeed.Contract.RootBridge(&_L2PriceFeed.CallOpts)
}

// RootBridge is a free data retrieval call binding the contract method 0xcdea7c70.
//
// Solidity: function rootBridge() view returns(address)
func (_L2PriceFeed *L2PriceFeedCallerSession) RootBridge() (common.Address, error) {
	return _L2PriceFeed.Contract.RootBridge(&_L2PriceFeed.CallOpts)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x59d684cf.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) AddAggregator(opts *bind.TransactOpts, _priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "addAggregator", _priceFeedKey)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x59d684cf.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedSession) AddAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.AddAggregator(&_L2PriceFeed.TransactOpts, _priceFeedKey)
}

// AddAggregator is a paid mutator transaction binding the contract method 0x59d684cf.
//
// Solidity: function addAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) AddAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.AddAggregator(&_L2PriceFeed.TransactOpts, _priceFeedKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ambBridge, address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) Initialize(opts *bind.TransactOpts, _ambBridge common.Address, _rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "initialize", _ambBridge, _rootBridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ambBridge, address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedSession) Initialize(_ambBridge common.Address, _rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.Initialize(&_L2PriceFeed.TransactOpts, _ambBridge, _rootBridge)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ambBridge, address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) Initialize(_ambBridge common.Address, _rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.Initialize(&_L2PriceFeed.TransactOpts, _ambBridge, _rootBridge)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) RemoveAggregator(opts *bind.TransactOpts, _priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "removeAggregator", _priceFeedKey)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedSession) RemoveAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.RemoveAggregator(&_L2PriceFeed.TransactOpts, _priceFeedKey)
}

// RemoveAggregator is a paid mutator transaction binding the contract method 0x2bed9e0c.
//
// Solidity: function removeAggregator(bytes32 _priceFeedKey) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) RemoveAggregator(_priceFeedKey [32]byte) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.RemoveAggregator(&_L2PriceFeed.TransactOpts, _priceFeedKey)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2PriceFeed *L2PriceFeedTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2PriceFeed *L2PriceFeedSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2PriceFeed.Contract.RenounceOwnership(&_L2PriceFeed.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2PriceFeed.Contract.RenounceOwnership(&_L2PriceFeed.TransactOpts)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) SetLatestData(opts *bind.TransactOpts, _priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "setLatestData", _priceFeedKey, _price, _timestamp, _roundId)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_L2PriceFeed *L2PriceFeedSession) SetLatestData(_priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetLatestData(&_L2PriceFeed.TransactOpts, _priceFeedKey, _price, _timestamp, _roundId)
}

// SetLatestData is a paid mutator transaction binding the contract method 0x031d64bd.
//
// Solidity: function setLatestData(bytes32 _priceFeedKey, uint256 _price, uint256 _timestamp, uint256 _roundId) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) SetLatestData(_priceFeedKey [32]byte, _price *big.Int, _timestamp *big.Int, _roundId *big.Int) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetLatestData(&_L2PriceFeed.TransactOpts, _priceFeedKey, _price, _timestamp, _roundId)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2PriceFeed *L2PriceFeedSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetOwner(&_L2PriceFeed.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetOwner(&_L2PriceFeed.TransactOpts, newOwner)
}

// SetRootBridge is a paid mutator transaction binding the contract method 0xa7596ff8.
//
// Solidity: function setRootBridge(address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedTransactor) SetRootBridge(opts *bind.TransactOpts, _rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "setRootBridge", _rootBridge)
}

// SetRootBridge is a paid mutator transaction binding the contract method 0xa7596ff8.
//
// Solidity: function setRootBridge(address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedSession) SetRootBridge(_rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetRootBridge(&_L2PriceFeed.TransactOpts, _rootBridge)
}

// SetRootBridge is a paid mutator transaction binding the contract method 0xa7596ff8.
//
// Solidity: function setRootBridge(address _rootBridge) returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) SetRootBridge(_rootBridge common.Address) (*types.Transaction, error) {
	return _L2PriceFeed.Contract.SetRootBridge(&_L2PriceFeed.TransactOpts, _rootBridge)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_L2PriceFeed *L2PriceFeedTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2PriceFeed.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_L2PriceFeed *L2PriceFeedSession) UpdateOwner() (*types.Transaction, error) {
	return _L2PriceFeed.Contract.UpdateOwner(&_L2PriceFeed.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_L2PriceFeed *L2PriceFeedTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _L2PriceFeed.Contract.UpdateOwner(&_L2PriceFeed.TransactOpts)
}

// L2PriceFeedOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2PriceFeed contract.
type L2PriceFeedOwnershipTransferredIterator struct {
	Event *L2PriceFeedOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2PriceFeedOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PriceFeedOwnershipTransferred)
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
		it.Event = new(L2PriceFeedOwnershipTransferred)
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
func (it *L2PriceFeedOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PriceFeedOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PriceFeedOwnershipTransferred represents a OwnershipTransferred event raised by the L2PriceFeed contract.
type L2PriceFeedOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2PriceFeed *L2PriceFeedFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2PriceFeedOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2PriceFeed.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2PriceFeedOwnershipTransferredIterator{contract: _L2PriceFeed.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2PriceFeed *L2PriceFeedFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2PriceFeedOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2PriceFeed.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PriceFeedOwnershipTransferred)
				if err := _L2PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2PriceFeed *L2PriceFeedFilterer) ParseOwnershipTransferred(log types.Log) (*L2PriceFeedOwnershipTransferred, error) {
	event := new(L2PriceFeedOwnershipTransferred)
	if err := _L2PriceFeed.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2PriceFeedPriceFeedDataSetIterator is returned from FilterPriceFeedDataSet and is used to iterate over the raw logs and unpacked data for PriceFeedDataSet events raised by the L2PriceFeed contract.
type L2PriceFeedPriceFeedDataSetIterator struct {
	Event *L2PriceFeedPriceFeedDataSet // Event containing the contract specifics and raw log

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
func (it *L2PriceFeedPriceFeedDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2PriceFeedPriceFeedDataSet)
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
		it.Event = new(L2PriceFeedPriceFeedDataSet)
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
func (it *L2PriceFeedPriceFeedDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2PriceFeedPriceFeedDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2PriceFeedPriceFeedDataSet represents a PriceFeedDataSet event raised by the L2PriceFeed contract.
type L2PriceFeedPriceFeedDataSet struct {
	Key       [32]byte
	Price     *big.Int
	Timestamp *big.Int
	RoundId   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedDataSet is a free log retrieval operation binding the contract event 0xb3b3d4e42c62d694318e0149f90ed0f2e7be9757332da15fd55ae5a01d41722c.
//
// Solidity: event PriceFeedDataSet(bytes32 key, uint256 price, uint256 timestamp, uint256 roundId)
func (_L2PriceFeed *L2PriceFeedFilterer) FilterPriceFeedDataSet(opts *bind.FilterOpts) (*L2PriceFeedPriceFeedDataSetIterator, error) {

	logs, sub, err := _L2PriceFeed.contract.FilterLogs(opts, "PriceFeedDataSet")
	if err != nil {
		return nil, err
	}
	return &L2PriceFeedPriceFeedDataSetIterator{contract: _L2PriceFeed.contract, event: "PriceFeedDataSet", logs: logs, sub: sub}, nil
}

// WatchPriceFeedDataSet is a free log subscription operation binding the contract event 0xb3b3d4e42c62d694318e0149f90ed0f2e7be9757332da15fd55ae5a01d41722c.
//
// Solidity: event PriceFeedDataSet(bytes32 key, uint256 price, uint256 timestamp, uint256 roundId)
func (_L2PriceFeed *L2PriceFeedFilterer) WatchPriceFeedDataSet(opts *bind.WatchOpts, sink chan<- *L2PriceFeedPriceFeedDataSet) (event.Subscription, error) {

	logs, sub, err := _L2PriceFeed.contract.WatchLogs(opts, "PriceFeedDataSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2PriceFeedPriceFeedDataSet)
				if err := _L2PriceFeed.contract.UnpackLog(event, "PriceFeedDataSet", log); err != nil {
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

// ParsePriceFeedDataSet is a log parse operation binding the contract event 0xb3b3d4e42c62d694318e0149f90ed0f2e7be9757332da15fd55ae5a01d41722c.
//
// Solidity: event PriceFeedDataSet(bytes32 key, uint256 price, uint256 timestamp, uint256 roundId)
func (_L2PriceFeed *L2PriceFeedFilterer) ParsePriceFeedDataSet(log types.Log) (*L2PriceFeedPriceFeedDataSet, error) {
	event := new(L2PriceFeedPriceFeedDataSet)
	if err := _L2PriceFeed.contract.UnpackLog(event, "PriceFeedDataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
