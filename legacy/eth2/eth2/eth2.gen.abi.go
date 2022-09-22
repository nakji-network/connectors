// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth2

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

// Eth2MetaData contains all meta data concerning the Eth2 contract.
var Eth2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"amount\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"index\",\"type\":\"bytes\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"deposit_data_root\",\"type\":\"bytes32\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_count\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_deposit_root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// Eth2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Eth2MetaData.ABI instead.
var Eth2ABI = Eth2MetaData.ABI

// Eth2 is an auto generated Go binding around an Ethereum contract.
type Eth2 struct {
	Eth2Caller     // Read-only binding to the contract
	Eth2Transactor // Write-only binding to the contract
	Eth2Filterer   // Log filterer for contract events
}

// Eth2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Eth2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Eth2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Eth2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eth2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Eth2Session struct {
	Contract     *Eth2             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eth2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Eth2CallerSession struct {
	Contract *Eth2Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Eth2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Eth2TransactorSession struct {
	Contract     *Eth2Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eth2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Eth2Raw struct {
	Contract *Eth2 // Generic contract binding to access the raw methods on
}

// Eth2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Eth2CallerRaw struct {
	Contract *Eth2Caller // Generic read-only contract binding to access the raw methods on
}

// Eth2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Eth2TransactorRaw struct {
	Contract *Eth2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEth2 creates a new instance of Eth2, bound to a specific deployed contract.
func NewEth2(address common.Address, backend bind.ContractBackend) (*Eth2, error) {
	contract, err := bindEth2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth2{Eth2Caller: Eth2Caller{contract: contract}, Eth2Transactor: Eth2Transactor{contract: contract}, Eth2Filterer: Eth2Filterer{contract: contract}}, nil
}

// NewEth2Caller creates a new read-only instance of Eth2, bound to a specific deployed contract.
func NewEth2Caller(address common.Address, caller bind.ContractCaller) (*Eth2Caller, error) {
	contract, err := bindEth2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Eth2Caller{contract: contract}, nil
}

// NewEth2Transactor creates a new write-only instance of Eth2, bound to a specific deployed contract.
func NewEth2Transactor(address common.Address, transactor bind.ContractTransactor) (*Eth2Transactor, error) {
	contract, err := bindEth2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Eth2Transactor{contract: contract}, nil
}

// NewEth2Filterer creates a new log filterer instance of Eth2, bound to a specific deployed contract.
func NewEth2Filterer(address common.Address, filterer bind.ContractFilterer) (*Eth2Filterer, error) {
	contract, err := bindEth2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Eth2Filterer{contract: contract}, nil
}

// bindEth2 binds a generic wrapper to an already deployed contract.
func bindEth2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Eth2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth2 *Eth2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth2.Contract.Eth2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth2 *Eth2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth2.Contract.Eth2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth2 *Eth2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth2.Contract.Eth2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth2 *Eth2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eth2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth2 *Eth2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth2 *Eth2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth2.Contract.contract.Transact(opts, method, params...)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Eth2 *Eth2Caller) GetDepositCount(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Eth2.contract.Call(opts, &out, "get_deposit_count")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Eth2 *Eth2Session) GetDepositCount() ([]byte, error) {
	return _Eth2.Contract.GetDepositCount(&_Eth2.CallOpts)
}

// GetDepositCount is a free data retrieval call binding the contract method 0x621fd130.
//
// Solidity: function get_deposit_count() view returns(bytes)
func (_Eth2 *Eth2CallerSession) GetDepositCount() ([]byte, error) {
	return _Eth2.Contract.GetDepositCount(&_Eth2.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Eth2 *Eth2Caller) GetDepositRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Eth2.contract.Call(opts, &out, "get_deposit_root")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Eth2 *Eth2Session) GetDepositRoot() ([32]byte, error) {
	return _Eth2.Contract.GetDepositRoot(&_Eth2.CallOpts)
}

// GetDepositRoot is a free data retrieval call binding the contract method 0xc5f2892f.
//
// Solidity: function get_deposit_root() view returns(bytes32)
func (_Eth2 *Eth2CallerSession) GetDepositRoot() ([32]byte, error) {
	return _Eth2.Contract.GetDepositRoot(&_Eth2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Eth2 *Eth2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Eth2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Eth2 *Eth2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Eth2.Contract.SupportsInterface(&_Eth2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Eth2 *Eth2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Eth2.Contract.SupportsInterface(&_Eth2.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Eth2 *Eth2Transactor) Deposit(opts *bind.TransactOpts, pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2.contract.Transact(opts, "deposit", pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Eth2 *Eth2Session) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2.Contract.Deposit(&_Eth2.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Deposit is a paid mutator transaction binding the contract method 0x22895118.
//
// Solidity: function deposit(bytes pubkey, bytes withdrawal_credentials, bytes signature, bytes32 deposit_data_root) payable returns()
func (_Eth2 *Eth2TransactorSession) Deposit(pubkey []byte, withdrawal_credentials []byte, signature []byte, deposit_data_root [32]byte) (*types.Transaction, error) {
	return _Eth2.Contract.Deposit(&_Eth2.TransactOpts, pubkey, withdrawal_credentials, signature, deposit_data_root)
}

// Eth2DepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Eth2 contract.
type Eth2DepositEventIterator struct {
	Event *Eth2DepositEvent // Event containing the contract specifics and raw log

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
func (it *Eth2DepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Eth2DepositEvent)
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
		it.Event = new(Eth2DepositEvent)
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
func (it *Eth2DepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Eth2DepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Eth2DepositEvent represents a DepositEvent event raised by the Eth2 contract.
type Eth2DepositEvent struct {
	Pubkey                []byte
	WithdrawalCredentials []byte
	Amount                []byte
	Signature             []byte
	Index                 []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2 *Eth2Filterer) FilterDepositEvent(opts *bind.FilterOpts) (*Eth2DepositEventIterator, error) {

	logs, sub, err := _Eth2.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &Eth2DepositEventIterator{contract: _Eth2.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2 *Eth2Filterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *Eth2DepositEvent) (event.Subscription, error) {

	logs, sub, err := _Eth2.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Eth2DepositEvent)
				if err := _Eth2.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x649bbc62d0e31342afea4e5cd82d4049e7e1ee912fc0889aa790803be39038c5.
//
// Solidity: event DepositEvent(bytes pubkey, bytes withdrawal_credentials, bytes amount, bytes signature, bytes index)
func (_Eth2 *Eth2Filterer) ParseDepositEvent(log types.Log) (*Eth2DepositEvent, error) {
	event := new(Eth2DepositEvent)
	if err := _Eth2.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
