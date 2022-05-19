// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package factory

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

// IDAOfiV1FactoryABI is the input ABI used to generate the binding from.
const IDAOfiV1FactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pairOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"slopeNumerator\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"n\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fee\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pairOwner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"slopeNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"n\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fee\",\"type\":\"uint32\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"formula\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"slopeNumerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"n\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fee\",\"type\":\"uint32\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedParams\",\"type\":\"bytes32\"}],\"name\":\"pairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IDAOfiV1Factory is an auto generated Go binding around an Ethereum contract.
type IDAOfiV1Factory struct {
	IDAOfiV1FactoryCaller     // Read-only binding to the contract
	IDAOfiV1FactoryTransactor // Write-only binding to the contract
	IDAOfiV1FactoryFilterer   // Log filterer for contract events
}

// IDAOfiV1FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDAOfiV1FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDAOfiV1FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDAOfiV1FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOfiV1FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDAOfiV1FactorySession struct {
	Contract     *IDAOfiV1Factory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAOfiV1FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDAOfiV1FactoryCallerSession struct {
	Contract *IDAOfiV1FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IDAOfiV1FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDAOfiV1FactoryTransactorSession struct {
	Contract     *IDAOfiV1FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IDAOfiV1FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDAOfiV1FactoryRaw struct {
	Contract *IDAOfiV1Factory // Generic contract binding to access the raw methods on
}

// IDAOfiV1FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDAOfiV1FactoryCallerRaw struct {
	Contract *IDAOfiV1FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IDAOfiV1FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDAOfiV1FactoryTransactorRaw struct {
	Contract *IDAOfiV1FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDAOfiV1Factory creates a new instance of IDAOfiV1Factory, bound to a specific deployed contract.
func NewIDAOfiV1Factory(address common.Address, backend bind.ContractBackend) (*IDAOfiV1Factory, error) {
	contract, err := bindIDAOfiV1Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1Factory{IDAOfiV1FactoryCaller: IDAOfiV1FactoryCaller{contract: contract}, IDAOfiV1FactoryTransactor: IDAOfiV1FactoryTransactor{contract: contract}, IDAOfiV1FactoryFilterer: IDAOfiV1FactoryFilterer{contract: contract}}, nil
}

// NewIDAOfiV1FactoryCaller creates a new read-only instance of IDAOfiV1Factory, bound to a specific deployed contract.
func NewIDAOfiV1FactoryCaller(address common.Address, caller bind.ContractCaller) (*IDAOfiV1FactoryCaller, error) {
	contract, err := bindIDAOfiV1Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1FactoryCaller{contract: contract}, nil
}

// NewIDAOfiV1FactoryTransactor creates a new write-only instance of IDAOfiV1Factory, bound to a specific deployed contract.
func NewIDAOfiV1FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IDAOfiV1FactoryTransactor, error) {
	contract, err := bindIDAOfiV1Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1FactoryTransactor{contract: contract}, nil
}

// NewIDAOfiV1FactoryFilterer creates a new log filterer instance of IDAOfiV1Factory, bound to a specific deployed contract.
func NewIDAOfiV1FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IDAOfiV1FactoryFilterer, error) {
	contract, err := bindIDAOfiV1Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1FactoryFilterer{contract: contract}, nil
}

// bindIDAOfiV1Factory binds a generic wrapper to an already deployed contract.
func bindIDAOfiV1Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDAOfiV1FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOfiV1Factory *IDAOfiV1FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOfiV1Factory.Contract.IDAOfiV1FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOfiV1Factory *IDAOfiV1FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.IDAOfiV1FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOfiV1Factory *IDAOfiV1FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.IDAOfiV1FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAOfiV1Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAOfiV1Factory *IDAOfiV1FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAOfiV1Factory *IDAOfiV1FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Factory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.AllPairs(&_IDAOfiV1Factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.AllPairs(&_IDAOfiV1Factory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IDAOfiV1Factory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) AllPairsLength() (*big.Int, error) {
	return _IDAOfiV1Factory.Contract.AllPairsLength(&_IDAOfiV1Factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _IDAOfiV1Factory.Contract.AllPairsLength(&_IDAOfiV1Factory.CallOpts)
}

// Formula is a free data retrieval call binding the contract method 0x4b75f54f.
//
// Solidity: function formula() view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCaller) Formula(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Factory.contract.Call(opts, &out, "formula")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Formula is a free data retrieval call binding the contract method 0x4b75f54f.
//
// Solidity: function formula() view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) Formula() (common.Address, error) {
	return _IDAOfiV1Factory.Contract.Formula(&_IDAOfiV1Factory.CallOpts)
}

// Formula is a free data retrieval call binding the contract method 0x4b75f54f.
//
// Solidity: function formula() view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerSession) Formula() (common.Address, error) {
	return _IDAOfiV1Factory.Contract.Formula(&_IDAOfiV1Factory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0x6b8b5e31.
//
// Solidity: function getPair(address baseToken, address quoteToken, uint32 slopeNumerator, uint32 n, uint32 fee) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCaller) GetPair(opts *bind.CallOpts, baseToken common.Address, quoteToken common.Address, slopeNumerator uint32, n uint32, fee uint32) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Factory.contract.Call(opts, &out, "getPair", baseToken, quoteToken, slopeNumerator, n, fee)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0x6b8b5e31.
//
// Solidity: function getPair(address baseToken, address quoteToken, uint32 slopeNumerator, uint32 n, uint32 fee) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) GetPair(baseToken common.Address, quoteToken common.Address, slopeNumerator uint32, n uint32, fee uint32) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.GetPair(&_IDAOfiV1Factory.CallOpts, baseToken, quoteToken, slopeNumerator, n, fee)
}

// GetPair is a free data retrieval call binding the contract method 0x6b8b5e31.
//
// Solidity: function getPair(address baseToken, address quoteToken, uint32 slopeNumerator, uint32 n, uint32 fee) view returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerSession) GetPair(baseToken common.Address, quoteToken common.Address, slopeNumerator uint32, n uint32, fee uint32) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.GetPair(&_IDAOfiV1Factory.CallOpts, baseToken, quoteToken, slopeNumerator, n, fee)
}

// Pairs is a free data retrieval call binding the contract method 0x673e0481.
//
// Solidity: function pairs(bytes32 hashedParams) view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCaller) Pairs(opts *bind.CallOpts, hashedParams [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IDAOfiV1Factory.contract.Call(opts, &out, "pairs", hashedParams)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pairs is a free data retrieval call binding the contract method 0x673e0481.
//
// Solidity: function pairs(bytes32 hashedParams) view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) Pairs(hashedParams [32]byte) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.Pairs(&_IDAOfiV1Factory.CallOpts, hashedParams)
}

// Pairs is a free data retrieval call binding the contract method 0x673e0481.
//
// Solidity: function pairs(bytes32 hashedParams) view returns(address)
func (_IDAOfiV1Factory *IDAOfiV1FactoryCallerSession) Pairs(hashedParams [32]byte) (common.Address, error) {
	return _IDAOfiV1Factory.Contract.Pairs(&_IDAOfiV1Factory.CallOpts, hashedParams)
}

// CreatePair is a paid mutator transaction binding the contract method 0x835f6ef6.
//
// Solidity: function createPair(address router, address baseToken, address quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee) returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryTransactor) CreatePair(opts *bind.TransactOpts, router common.Address, baseToken common.Address, quoteToken common.Address, pairOwner common.Address, slopeNumerator uint32, n uint32, fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Factory.contract.Transact(opts, "createPair", router, baseToken, quoteToken, pairOwner, slopeNumerator, n, fee)
}

// CreatePair is a paid mutator transaction binding the contract method 0x835f6ef6.
//
// Solidity: function createPair(address router, address baseToken, address quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee) returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactorySession) CreatePair(router common.Address, baseToken common.Address, quoteToken common.Address, pairOwner common.Address, slopeNumerator uint32, n uint32, fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.CreatePair(&_IDAOfiV1Factory.TransactOpts, router, baseToken, quoteToken, pairOwner, slopeNumerator, n, fee)
}

// CreatePair is a paid mutator transaction binding the contract method 0x835f6ef6.
//
// Solidity: function createPair(address router, address baseToken, address quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee) returns(address pair)
func (_IDAOfiV1Factory *IDAOfiV1FactoryTransactorSession) CreatePair(router common.Address, baseToken common.Address, quoteToken common.Address, pairOwner common.Address, slopeNumerator uint32, n uint32, fee uint32) (*types.Transaction, error) {
	return _IDAOfiV1Factory.Contract.CreatePair(&_IDAOfiV1Factory.TransactOpts, router, baseToken, quoteToken, pairOwner, slopeNumerator, n, fee)
}

// IDAOfiV1FactoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the IDAOfiV1Factory contract.
type IDAOfiV1FactoryPairCreatedIterator struct {
	Event *IDAOfiV1FactoryPairCreated // Event containing the contract specifics and raw log

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
func (it *IDAOfiV1FactoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOfiV1FactoryPairCreated)
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
		it.Event = new(IDAOfiV1FactoryPairCreated)
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
func (it *IDAOfiV1FactoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOfiV1FactoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOfiV1FactoryPairCreated represents a PairCreated event raised by the IDAOfiV1Factory contract.
type IDAOfiV1FactoryPairCreated struct {
	BaseToken      common.Address
	QuoteToken     common.Address
	PairOwner      common.Address
	SlopeNumerator uint32
	N              uint32
	Fee            uint32
	Pair           common.Address
	Length         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0xb3fa7f0692a1e60cd160a3ab4e54c82ddbd058ac672f4654686c6311b672a837.
//
// Solidity: event PairCreated(address indexed baseToken, address indexed quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee, address pair, uint256 length)
func (_IDAOfiV1Factory *IDAOfiV1FactoryFilterer) FilterPairCreated(opts *bind.FilterOpts, baseToken []common.Address, quoteToken []common.Address) (*IDAOfiV1FactoryPairCreatedIterator, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _IDAOfiV1Factory.contract.FilterLogs(opts, "PairCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return &IDAOfiV1FactoryPairCreatedIterator{contract: _IDAOfiV1Factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0xb3fa7f0692a1e60cd160a3ab4e54c82ddbd058ac672f4654686c6311b672a837.
//
// Solidity: event PairCreated(address indexed baseToken, address indexed quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee, address pair, uint256 length)
func (_IDAOfiV1Factory *IDAOfiV1FactoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *IDAOfiV1FactoryPairCreated, baseToken []common.Address, quoteToken []common.Address) (event.Subscription, error) {

	var baseTokenRule []interface{}
	for _, baseTokenItem := range baseToken {
		baseTokenRule = append(baseTokenRule, baseTokenItem)
	}
	var quoteTokenRule []interface{}
	for _, quoteTokenItem := range quoteToken {
		quoteTokenRule = append(quoteTokenRule, quoteTokenItem)
	}

	logs, sub, err := _IDAOfiV1Factory.contract.WatchLogs(opts, "PairCreated", baseTokenRule, quoteTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOfiV1FactoryPairCreated)
				if err := _IDAOfiV1Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0xb3fa7f0692a1e60cd160a3ab4e54c82ddbd058ac672f4654686c6311b672a837.
//
// Solidity: event PairCreated(address indexed baseToken, address indexed quoteToken, address pairOwner, uint32 slopeNumerator, uint32 n, uint32 fee, address pair, uint256 length)
func (_IDAOfiV1Factory *IDAOfiV1FactoryFilterer) ParsePairCreated(log types.Log) (*IDAOfiV1FactoryPairCreated, error) {
	event := new(IDAOfiV1FactoryPairCreated)
	if err := _IDAOfiV1Factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}
