// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package JOIN_FAB

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

// JOINFABABI is the input ABI used to generate the binding from.
const JOINFABABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vat\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"NewJoin\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"}],\"name\":\"newAuthGemJoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"}],\"name\":\"newGemJoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"}],\"name\":\"newGemJoin5\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// JOINFAB is an auto generated Go binding around an Ethereum contract.
type JOINFAB struct {
	JOINFABCaller     // Read-only binding to the contract
	JOINFABTransactor // Write-only binding to the contract
	JOINFABFilterer   // Log filterer for contract events
}

// JOINFABCaller is an auto generated read-only Go binding around an Ethereum contract.
type JOINFABCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JOINFABTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JOINFABTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JOINFABFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JOINFABFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JOINFABSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JOINFABSession struct {
	Contract     *JOINFAB          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JOINFABCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JOINFABCallerSession struct {
	Contract *JOINFABCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// JOINFABTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JOINFABTransactorSession struct {
	Contract     *JOINFABTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JOINFABRaw is an auto generated low-level Go binding around an Ethereum contract.
type JOINFABRaw struct {
	Contract *JOINFAB // Generic contract binding to access the raw methods on
}

// JOINFABCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JOINFABCallerRaw struct {
	Contract *JOINFABCaller // Generic read-only contract binding to access the raw methods on
}

// JOINFABTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JOINFABTransactorRaw struct {
	Contract *JOINFABTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJOINFAB creates a new instance of JOINFAB, bound to a specific deployed contract.
func NewJOINFAB(address common.Address, backend bind.ContractBackend) (*JOINFAB, error) {
	contract, err := bindJOINFAB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JOINFAB{JOINFABCaller: JOINFABCaller{contract: contract}, JOINFABTransactor: JOINFABTransactor{contract: contract}, JOINFABFilterer: JOINFABFilterer{contract: contract}}, nil
}

// NewJOINFABCaller creates a new read-only instance of JOINFAB, bound to a specific deployed contract.
func NewJOINFABCaller(address common.Address, caller bind.ContractCaller) (*JOINFABCaller, error) {
	contract, err := bindJOINFAB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JOINFABCaller{contract: contract}, nil
}

// NewJOINFABTransactor creates a new write-only instance of JOINFAB, bound to a specific deployed contract.
func NewJOINFABTransactor(address common.Address, transactor bind.ContractTransactor) (*JOINFABTransactor, error) {
	contract, err := bindJOINFAB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JOINFABTransactor{contract: contract}, nil
}

// NewJOINFABFilterer creates a new log filterer instance of JOINFAB, bound to a specific deployed contract.
func NewJOINFABFilterer(address common.Address, filterer bind.ContractFilterer) (*JOINFABFilterer, error) {
	contract, err := bindJOINFAB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JOINFABFilterer{contract: contract}, nil
}

// bindJOINFAB binds a generic wrapper to an already deployed contract.
func bindJOINFAB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JOINFABABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JOINFAB *JOINFABRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JOINFAB.Contract.JOINFABCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JOINFAB *JOINFABRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JOINFAB.Contract.JOINFABTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JOINFAB *JOINFABRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JOINFAB.Contract.JOINFABTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JOINFAB *JOINFABCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JOINFAB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JOINFAB *JOINFABTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JOINFAB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JOINFAB *JOINFABTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JOINFAB.Contract.contract.Transact(opts, method, params...)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_JOINFAB *JOINFABCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JOINFAB.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_JOINFAB *JOINFABSession) Vat() (common.Address, error) {
	return _JOINFAB.Contract.Vat(&_JOINFAB.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_JOINFAB *JOINFABCallerSession) Vat() (common.Address, error) {
	return _JOINFAB.Contract.Vat(&_JOINFAB.CallOpts)
}

// NewAuthGemJoin is a paid mutator transaction binding the contract method 0x2866be72.
//
// Solidity: function newAuthGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactor) NewAuthGemJoin(opts *bind.TransactOpts, _owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.contract.Transact(opts, "newAuthGemJoin", _owner, _ilk, _gem)
}

// NewAuthGemJoin is a paid mutator transaction binding the contract method 0x2866be72.
//
// Solidity: function newAuthGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABSession) NewAuthGemJoin(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewAuthGemJoin(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// NewAuthGemJoin is a paid mutator transaction binding the contract method 0x2866be72.
//
// Solidity: function newAuthGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactorSession) NewAuthGemJoin(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewAuthGemJoin(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// NewGemJoin is a paid mutator transaction binding the contract method 0x45949cf5.
//
// Solidity: function newGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactor) NewGemJoin(opts *bind.TransactOpts, _owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.contract.Transact(opts, "newGemJoin", _owner, _ilk, _gem)
}

// NewGemJoin is a paid mutator transaction binding the contract method 0x45949cf5.
//
// Solidity: function newGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABSession) NewGemJoin(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewGemJoin(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// NewGemJoin is a paid mutator transaction binding the contract method 0x45949cf5.
//
// Solidity: function newGemJoin(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactorSession) NewGemJoin(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewGemJoin(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// NewGemJoin5 is a paid mutator transaction binding the contract method 0x829e402f.
//
// Solidity: function newGemJoin5(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactor) NewGemJoin5(opts *bind.TransactOpts, _owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.contract.Transact(opts, "newGemJoin5", _owner, _ilk, _gem)
}

// NewGemJoin5 is a paid mutator transaction binding the contract method 0x829e402f.
//
// Solidity: function newGemJoin5(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABSession) NewGemJoin5(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewGemJoin5(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// NewGemJoin5 is a paid mutator transaction binding the contract method 0x829e402f.
//
// Solidity: function newGemJoin5(address _owner, bytes32 _ilk, address _gem) returns(address join)
func (_JOINFAB *JOINFABTransactorSession) NewGemJoin5(_owner common.Address, _ilk [32]byte, _gem common.Address) (*types.Transaction, error) {
	return _JOINFAB.Contract.NewGemJoin5(&_JOINFAB.TransactOpts, _owner, _ilk, _gem)
}

// JOINFABNewJoinIterator is returned from FilterNewJoin and is used to iterate over the raw logs and unpacked data for NewJoin events raised by the JOINFAB contract.
type JOINFABNewJoinIterator struct {
	Event *JOINFABNewJoin // Event containing the contract specifics and raw log

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
func (it *JOINFABNewJoinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JOINFABNewJoin)
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
		it.Event = new(JOINFABNewJoin)
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
func (it *JOINFABNewJoinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JOINFABNewJoinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JOINFABNewJoin represents a NewJoin event raised by the JOINFAB contract.
type JOINFABNewJoin struct {
	Join common.Address
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewJoin is a free log retrieval operation binding the contract event 0x0e3a78b6b2fd5969f1ac0cc2b65daf580334290350a8eeff1747b804e30660a8.
//
// Solidity: event NewJoin(address indexed join, bytes data)
func (_JOINFAB *JOINFABFilterer) FilterNewJoin(opts *bind.FilterOpts, join []common.Address) (*JOINFABNewJoinIterator, error) {

	var joinRule []interface{}
	for _, joinItem := range join {
		joinRule = append(joinRule, joinItem)
	}

	logs, sub, err := _JOINFAB.contract.FilterLogs(opts, "NewJoin", joinRule)
	if err != nil {
		return nil, err
	}
	return &JOINFABNewJoinIterator{contract: _JOINFAB.contract, event: "NewJoin", logs: logs, sub: sub}, nil
}

// WatchNewJoin is a free log subscription operation binding the contract event 0x0e3a78b6b2fd5969f1ac0cc2b65daf580334290350a8eeff1747b804e30660a8.
//
// Solidity: event NewJoin(address indexed join, bytes data)
func (_JOINFAB *JOINFABFilterer) WatchNewJoin(opts *bind.WatchOpts, sink chan<- *JOINFABNewJoin, join []common.Address) (event.Subscription, error) {

	var joinRule []interface{}
	for _, joinItem := range join {
		joinRule = append(joinRule, joinItem)
	}

	logs, sub, err := _JOINFAB.contract.WatchLogs(opts, "NewJoin", joinRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JOINFABNewJoin)
				if err := _JOINFAB.contract.UnpackLog(event, "NewJoin", log); err != nil {
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

// ParseNewJoin is a log parse operation binding the contract event 0x0e3a78b6b2fd5969f1ac0cc2b65daf580334290350a8eeff1747b804e30660a8.
//
// Solidity: event NewJoin(address indexed join, bytes data)
func (_JOINFAB *JOINFABFilterer) ParseNewJoin(log types.Log) (*JOINFABNewJoin, error) {
	event := new(JOINFABNewJoin)
	if err := _JOINFAB.contract.UnpackLog(event, "NewJoin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
