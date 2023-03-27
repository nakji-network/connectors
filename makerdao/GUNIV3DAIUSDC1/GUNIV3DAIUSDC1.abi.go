// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GUniV3DaiUSDC1

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

// GUNIV3DAIUSDC1ABI is the input ABI used to generate the binding from.
const GUNIV3DAIUSDC1ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"ProxyAdminTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"ProxyImplementationUpdated\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"proxyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"id\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferProxyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GUNIV3DAIUSDC1 is an auto generated Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1 struct {
	GUNIV3DAIUSDC1Caller     // Read-only binding to the contract
	GUNIV3DAIUSDC1Transactor // Write-only binding to the contract
	GUNIV3DAIUSDC1Filterer   // Log filterer for contract events
}

// GUNIV3DAIUSDC1Caller is an auto generated read-only Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUNIV3DAIUSDC1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUNIV3DAIUSDC1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GUNIV3DAIUSDC1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GUNIV3DAIUSDC1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GUNIV3DAIUSDC1Session struct {
	Contract     *GUNIV3DAIUSDC1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GUNIV3DAIUSDC1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GUNIV3DAIUSDC1CallerSession struct {
	Contract *GUNIV3DAIUSDC1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GUNIV3DAIUSDC1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GUNIV3DAIUSDC1TransactorSession struct {
	Contract     *GUNIV3DAIUSDC1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GUNIV3DAIUSDC1Raw is an auto generated low-level Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1Raw struct {
	Contract *GUNIV3DAIUSDC1 // Generic contract binding to access the raw methods on
}

// GUNIV3DAIUSDC1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1CallerRaw struct {
	Contract *GUNIV3DAIUSDC1Caller // Generic read-only contract binding to access the raw methods on
}

// GUNIV3DAIUSDC1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GUNIV3DAIUSDC1TransactorRaw struct {
	Contract *GUNIV3DAIUSDC1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewGUNIV3DAIUSDC1 creates a new instance of GUNIV3DAIUSDC1, bound to a specific deployed contract.
func NewGUNIV3DAIUSDC1(address common.Address, backend bind.ContractBackend) (*GUNIV3DAIUSDC1, error) {
	contract, err := bindGUNIV3DAIUSDC1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1{GUNIV3DAIUSDC1Caller: GUNIV3DAIUSDC1Caller{contract: contract}, GUNIV3DAIUSDC1Transactor: GUNIV3DAIUSDC1Transactor{contract: contract}, GUNIV3DAIUSDC1Filterer: GUNIV3DAIUSDC1Filterer{contract: contract}}, nil
}

// NewGUNIV3DAIUSDC1Caller creates a new read-only instance of GUNIV3DAIUSDC1, bound to a specific deployed contract.
func NewGUNIV3DAIUSDC1Caller(address common.Address, caller bind.ContractCaller) (*GUNIV3DAIUSDC1Caller, error) {
	contract, err := bindGUNIV3DAIUSDC1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1Caller{contract: contract}, nil
}

// NewGUNIV3DAIUSDC1Transactor creates a new write-only instance of GUNIV3DAIUSDC1, bound to a specific deployed contract.
func NewGUNIV3DAIUSDC1Transactor(address common.Address, transactor bind.ContractTransactor) (*GUNIV3DAIUSDC1Transactor, error) {
	contract, err := bindGUNIV3DAIUSDC1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1Transactor{contract: contract}, nil
}

// NewGUNIV3DAIUSDC1Filterer creates a new log filterer instance of GUNIV3DAIUSDC1, bound to a specific deployed contract.
func NewGUNIV3DAIUSDC1Filterer(address common.Address, filterer bind.ContractFilterer) (*GUNIV3DAIUSDC1Filterer, error) {
	contract, err := bindGUNIV3DAIUSDC1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1Filterer{contract: contract}, nil
}

// bindGUNIV3DAIUSDC1 binds a generic wrapper to an already deployed contract.
func bindGUNIV3DAIUSDC1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GUNIV3DAIUSDC1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GUNIV3DAIUSDC1.Contract.GUNIV3DAIUSDC1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.GUNIV3DAIUSDC1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.GUNIV3DAIUSDC1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GUNIV3DAIUSDC1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.contract.Transact(opts, method, params...)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Caller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GUNIV3DAIUSDC1.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) ProxyAdmin() (common.Address, error) {
	return _GUNIV3DAIUSDC1.Contract.ProxyAdmin(&_GUNIV3DAIUSDC1.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1CallerSession) ProxyAdmin() (common.Address, error) {
	return _GUNIV3DAIUSDC1.Contract.ProxyAdmin(&_GUNIV3DAIUSDC1.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 id) view returns(bool)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Caller) SupportsInterface(opts *bind.CallOpts, id [4]byte) (bool, error) {
	var out []interface{}
	err := _GUNIV3DAIUSDC1.contract.Call(opts, &out, "supportsInterface", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 id) view returns(bool)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) SupportsInterface(id [4]byte) (bool, error) {
	return _GUNIV3DAIUSDC1.Contract.SupportsInterface(&_GUNIV3DAIUSDC1.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 id) view returns(bool)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1CallerSession) SupportsInterface(id [4]byte) (bool, error) {
	return _GUNIV3DAIUSDC1.Contract.SupportsInterface(&_GUNIV3DAIUSDC1.CallOpts, id)
}

// TransferProxyAdmin is a paid mutator transaction binding the contract method 0x8356ca4f.
//
// Solidity: function transferProxyAdmin(address newAdmin) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Transactor) TransferProxyAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.contract.Transact(opts, "transferProxyAdmin", newAdmin)
}

// TransferProxyAdmin is a paid mutator transaction binding the contract method 0x8356ca4f.
//
// Solidity: function transferProxyAdmin(address newAdmin) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) TransferProxyAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.TransferProxyAdmin(&_GUNIV3DAIUSDC1.TransactOpts, newAdmin)
}

// TransferProxyAdmin is a paid mutator transaction binding the contract method 0x8356ca4f.
//
// Solidity: function transferProxyAdmin(address newAdmin) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorSession) TransferProxyAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.TransferProxyAdmin(&_GUNIV3DAIUSDC1.TransactOpts, newAdmin)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.UpgradeTo(&_GUNIV3DAIUSDC1.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.UpgradeTo(&_GUNIV3DAIUSDC1.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.UpgradeToAndCall(&_GUNIV3DAIUSDC1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.UpgradeToAndCall(&_GUNIV3DAIUSDC1.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.Fallback(&_GUNIV3DAIUSDC1.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.Fallback(&_GUNIV3DAIUSDC1.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Session) Receive() (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.Receive(&_GUNIV3DAIUSDC1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1TransactorSession) Receive() (*types.Transaction, error) {
	return _GUNIV3DAIUSDC1.Contract.Receive(&_GUNIV3DAIUSDC1.TransactOpts)
}

// GUNIV3DAIUSDC1ProxyAdminTransferredIterator is returned from FilterProxyAdminTransferred and is used to iterate over the raw logs and unpacked data for ProxyAdminTransferred events raised by the GUNIV3DAIUSDC1 contract.
type GUNIV3DAIUSDC1ProxyAdminTransferredIterator struct {
	Event *GUNIV3DAIUSDC1ProxyAdminTransferred // Event containing the contract specifics and raw log

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
func (it *GUNIV3DAIUSDC1ProxyAdminTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUNIV3DAIUSDC1ProxyAdminTransferred)
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
		it.Event = new(GUNIV3DAIUSDC1ProxyAdminTransferred)
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
func (it *GUNIV3DAIUSDC1ProxyAdminTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUNIV3DAIUSDC1ProxyAdminTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUNIV3DAIUSDC1ProxyAdminTransferred represents a ProxyAdminTransferred event raised by the GUNIV3DAIUSDC1 contract.
type GUNIV3DAIUSDC1ProxyAdminTransferred struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProxyAdminTransferred is a free log retrieval operation binding the contract event 0xdf435d422321da6b195902d70fc417c06a32f88379c20dd8f2a8da07088cec29.
//
// Solidity: event ProxyAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) FilterProxyAdminTransferred(opts *bind.FilterOpts, previousAdmin []common.Address, newAdmin []common.Address) (*GUNIV3DAIUSDC1ProxyAdminTransferredIterator, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _GUNIV3DAIUSDC1.contract.FilterLogs(opts, "ProxyAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1ProxyAdminTransferredIterator{contract: _GUNIV3DAIUSDC1.contract, event: "ProxyAdminTransferred", logs: logs, sub: sub}, nil
}

// WatchProxyAdminTransferred is a free log subscription operation binding the contract event 0xdf435d422321da6b195902d70fc417c06a32f88379c20dd8f2a8da07088cec29.
//
// Solidity: event ProxyAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) WatchProxyAdminTransferred(opts *bind.WatchOpts, sink chan<- *GUNIV3DAIUSDC1ProxyAdminTransferred, previousAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _GUNIV3DAIUSDC1.contract.WatchLogs(opts, "ProxyAdminTransferred", previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUNIV3DAIUSDC1ProxyAdminTransferred)
				if err := _GUNIV3DAIUSDC1.contract.UnpackLog(event, "ProxyAdminTransferred", log); err != nil {
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

// ParseProxyAdminTransferred is a log parse operation binding the contract event 0xdf435d422321da6b195902d70fc417c06a32f88379c20dd8f2a8da07088cec29.
//
// Solidity: event ProxyAdminTransferred(address indexed previousAdmin, address indexed newAdmin)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) ParseProxyAdminTransferred(log types.Log) (*GUNIV3DAIUSDC1ProxyAdminTransferred, error) {
	event := new(GUNIV3DAIUSDC1ProxyAdminTransferred)
	if err := _GUNIV3DAIUSDC1.contract.UnpackLog(event, "ProxyAdminTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator is returned from FilterProxyImplementationUpdated and is used to iterate over the raw logs and unpacked data for ProxyImplementationUpdated events raised by the GUNIV3DAIUSDC1 contract.
type GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator struct {
	Event *GUNIV3DAIUSDC1ProxyImplementationUpdated // Event containing the contract specifics and raw log

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
func (it *GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GUNIV3DAIUSDC1ProxyImplementationUpdated)
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
		it.Event = new(GUNIV3DAIUSDC1ProxyImplementationUpdated)
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
func (it *GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GUNIV3DAIUSDC1ProxyImplementationUpdated represents a ProxyImplementationUpdated event raised by the GUNIV3DAIUSDC1 contract.
type GUNIV3DAIUSDC1ProxyImplementationUpdated struct {
	PreviousImplementation common.Address
	NewImplementation      common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProxyImplementationUpdated is a free log retrieval operation binding the contract event 0x5570d70a002632a7b0b3c9304cc89efb62d8da9eca0dbd7752c83b7379068296.
//
// Solidity: event ProxyImplementationUpdated(address indexed previousImplementation, address indexed newImplementation)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) FilterProxyImplementationUpdated(opts *bind.FilterOpts, previousImplementation []common.Address, newImplementation []common.Address) (*GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator, error) {

	var previousImplementationRule []interface{}
	for _, previousImplementationItem := range previousImplementation {
		previousImplementationRule = append(previousImplementationRule, previousImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _GUNIV3DAIUSDC1.contract.FilterLogs(opts, "ProxyImplementationUpdated", previousImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &GUNIV3DAIUSDC1ProxyImplementationUpdatedIterator{contract: _GUNIV3DAIUSDC1.contract, event: "ProxyImplementationUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyImplementationUpdated is a free log subscription operation binding the contract event 0x5570d70a002632a7b0b3c9304cc89efb62d8da9eca0dbd7752c83b7379068296.
//
// Solidity: event ProxyImplementationUpdated(address indexed previousImplementation, address indexed newImplementation)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) WatchProxyImplementationUpdated(opts *bind.WatchOpts, sink chan<- *GUNIV3DAIUSDC1ProxyImplementationUpdated, previousImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var previousImplementationRule []interface{}
	for _, previousImplementationItem := range previousImplementation {
		previousImplementationRule = append(previousImplementationRule, previousImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _GUNIV3DAIUSDC1.contract.WatchLogs(opts, "ProxyImplementationUpdated", previousImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GUNIV3DAIUSDC1ProxyImplementationUpdated)
				if err := _GUNIV3DAIUSDC1.contract.UnpackLog(event, "ProxyImplementationUpdated", log); err != nil {
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

// ParseProxyImplementationUpdated is a log parse operation binding the contract event 0x5570d70a002632a7b0b3c9304cc89efb62d8da9eca0dbd7752c83b7379068296.
//
// Solidity: event ProxyImplementationUpdated(address indexed previousImplementation, address indexed newImplementation)
func (_GUNIV3DAIUSDC1 *GUNIV3DAIUSDC1Filterer) ParseProxyImplementationUpdated(log types.Log) (*GUNIV3DAIUSDC1ProxyImplementationUpdated, error) {
	event := new(GUNIV3DAIUSDC1ProxyImplementationUpdated)
	if err := _GUNIV3DAIUSDC1.contract.UnpackLog(event, "ProxyImplementationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
