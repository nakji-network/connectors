// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ArbitrumGovRelay

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

// ARBITRUMGOVRELAYABI is the input ABI used to generate the binding from.
const ARBITRUMGOVRELAYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2GovernanceRelay\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"seqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TxToL2\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"contractIInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2GovernanceRelay\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"targetData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"l1CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ARBITRUMGOVRELAY is an auto generated Go binding around an Ethereum contract.
type ARBITRUMGOVRELAY struct {
	ARBITRUMGOVRELAYCaller     // Read-only binding to the contract
	ARBITRUMGOVRELAYTransactor // Write-only binding to the contract
	ARBITRUMGOVRELAYFilterer   // Log filterer for contract events
}

// ARBITRUMGOVRELAYCaller is an auto generated read-only Go binding around an Ethereum contract.
type ARBITRUMGOVRELAYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMGOVRELAYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ARBITRUMGOVRELAYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMGOVRELAYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ARBITRUMGOVRELAYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ARBITRUMGOVRELAYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ARBITRUMGOVRELAYSession struct {
	Contract     *ARBITRUMGOVRELAY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ARBITRUMGOVRELAYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ARBITRUMGOVRELAYCallerSession struct {
	Contract *ARBITRUMGOVRELAYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ARBITRUMGOVRELAYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ARBITRUMGOVRELAYTransactorSession struct {
	Contract     *ARBITRUMGOVRELAYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ARBITRUMGOVRELAYRaw is an auto generated low-level Go binding around an Ethereum contract.
type ARBITRUMGOVRELAYRaw struct {
	Contract *ARBITRUMGOVRELAY // Generic contract binding to access the raw methods on
}

// ARBITRUMGOVRELAYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ARBITRUMGOVRELAYCallerRaw struct {
	Contract *ARBITRUMGOVRELAYCaller // Generic read-only contract binding to access the raw methods on
}

// ARBITRUMGOVRELAYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ARBITRUMGOVRELAYTransactorRaw struct {
	Contract *ARBITRUMGOVRELAYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewARBITRUMGOVRELAY creates a new instance of ARBITRUMGOVRELAY, bound to a specific deployed contract.
func NewARBITRUMGOVRELAY(address common.Address, backend bind.ContractBackend) (*ARBITRUMGOVRELAY, error) {
	contract, err := bindARBITRUMGOVRELAY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAY{ARBITRUMGOVRELAYCaller: ARBITRUMGOVRELAYCaller{contract: contract}, ARBITRUMGOVRELAYTransactor: ARBITRUMGOVRELAYTransactor{contract: contract}, ARBITRUMGOVRELAYFilterer: ARBITRUMGOVRELAYFilterer{contract: contract}}, nil
}

// NewARBITRUMGOVRELAYCaller creates a new read-only instance of ARBITRUMGOVRELAY, bound to a specific deployed contract.
func NewARBITRUMGOVRELAYCaller(address common.Address, caller bind.ContractCaller) (*ARBITRUMGOVRELAYCaller, error) {
	contract, err := bindARBITRUMGOVRELAY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYCaller{contract: contract}, nil
}

// NewARBITRUMGOVRELAYTransactor creates a new write-only instance of ARBITRUMGOVRELAY, bound to a specific deployed contract.
func NewARBITRUMGOVRELAYTransactor(address common.Address, transactor bind.ContractTransactor) (*ARBITRUMGOVRELAYTransactor, error) {
	contract, err := bindARBITRUMGOVRELAY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYTransactor{contract: contract}, nil
}

// NewARBITRUMGOVRELAYFilterer creates a new log filterer instance of ARBITRUMGOVRELAY, bound to a specific deployed contract.
func NewARBITRUMGOVRELAYFilterer(address common.Address, filterer bind.ContractFilterer) (*ARBITRUMGOVRELAYFilterer, error) {
	contract, err := bindARBITRUMGOVRELAY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYFilterer{contract: contract}, nil
}

// bindARBITRUMGOVRELAY binds a generic wrapper to an already deployed contract.
func bindARBITRUMGOVRELAY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ARBITRUMGOVRELAYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMGOVRELAY.Contract.ARBITRUMGOVRELAYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.ARBITRUMGOVRELAYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.ARBITRUMGOVRELAYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ARBITRUMGOVRELAY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.contract.Transact(opts, method, params...)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMGOVRELAY.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Inbox() (common.Address, error) {
	return _ARBITRUMGOVRELAY.Contract.Inbox(&_ARBITRUMGOVRELAY.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCallerSession) Inbox() (common.Address, error) {
	return _ARBITRUMGOVRELAY.Contract.Inbox(&_ARBITRUMGOVRELAY.CallOpts)
}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCaller) L2GovernanceRelay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ARBITRUMGOVRELAY.contract.Call(opts, &out, "l2GovernanceRelay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) L2GovernanceRelay() (common.Address, error) {
	return _ARBITRUMGOVRELAY.Contract.L2GovernanceRelay(&_ARBITRUMGOVRELAY.CallOpts)
}

// L2GovernanceRelay is a free data retrieval call binding the contract method 0x862a98a1.
//
// Solidity: function l2GovernanceRelay() view returns(address)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCallerSession) L2GovernanceRelay() (common.Address, error) {
	return _ARBITRUMGOVRELAY.Contract.L2GovernanceRelay(&_ARBITRUMGOVRELAY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ARBITRUMGOVRELAY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMGOVRELAY.Contract.Wards(&_ARBITRUMGOVRELAY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _ARBITRUMGOVRELAY.Contract.Wards(&_ARBITRUMGOVRELAY.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Deny(&_ARBITRUMGOVRELAY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Deny(&_ARBITRUMGOVRELAY.TransactOpts, usr)
}

// Reclaim is a paid mutator transaction binding the contract method 0x8bd317eb.
//
// Solidity: function reclaim(address receiver, uint256 amount) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactor) Reclaim(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.contract.Transact(opts, "reclaim", receiver, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0x8bd317eb.
//
// Solidity: function reclaim(address receiver, uint256 amount) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Reclaim(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Reclaim(&_ARBITRUMGOVRELAY.TransactOpts, receiver, amount)
}

// Reclaim is a paid mutator transaction binding the contract method 0x8bd317eb.
//
// Solidity: function reclaim(address receiver, uint256 amount) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorSession) Reclaim(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Reclaim(&_ARBITRUMGOVRELAY.TransactOpts, receiver, amount)
}

// Relay is a paid mutator transaction binding the contract method 0xe668cdb0.
//
// Solidity: function relay(address target, bytes targetData, uint256 l1CallValue, uint256 maxGas, uint256 gasPriceBid, uint256 maxSubmissionCost) payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactor) Relay(opts *bind.TransactOpts, target common.Address, targetData []byte, l1CallValue *big.Int, maxGas *big.Int, gasPriceBid *big.Int, maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.contract.Transact(opts, "relay", target, targetData, l1CallValue, maxGas, gasPriceBid, maxSubmissionCost)
}

// Relay is a paid mutator transaction binding the contract method 0xe668cdb0.
//
// Solidity: function relay(address target, bytes targetData, uint256 l1CallValue, uint256 maxGas, uint256 gasPriceBid, uint256 maxSubmissionCost) payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Relay(target common.Address, targetData []byte, l1CallValue *big.Int, maxGas *big.Int, gasPriceBid *big.Int, maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Relay(&_ARBITRUMGOVRELAY.TransactOpts, target, targetData, l1CallValue, maxGas, gasPriceBid, maxSubmissionCost)
}

// Relay is a paid mutator transaction binding the contract method 0xe668cdb0.
//
// Solidity: function relay(address target, bytes targetData, uint256 l1CallValue, uint256 maxGas, uint256 gasPriceBid, uint256 maxSubmissionCost) payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorSession) Relay(target common.Address, targetData []byte, l1CallValue *big.Int, maxGas *big.Int, gasPriceBid *big.Int, maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Relay(&_ARBITRUMGOVRELAY.TransactOpts, target, targetData, l1CallValue, maxGas, gasPriceBid, maxSubmissionCost)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Rely(&_ARBITRUMGOVRELAY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Rely(&_ARBITRUMGOVRELAY.TransactOpts, usr)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYSession) Receive() (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Receive(&_ARBITRUMGOVRELAY.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYTransactorSession) Receive() (*types.Transaction, error) {
	return _ARBITRUMGOVRELAY.Contract.Receive(&_ARBITRUMGOVRELAY.TransactOpts)
}

// ARBITRUMGOVRELAYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYDenyIterator struct {
	Event *ARBITRUMGOVRELAYDeny // Event containing the contract specifics and raw log

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
func (it *ARBITRUMGOVRELAYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMGOVRELAYDeny)
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
		it.Event = new(ARBITRUMGOVRELAYDeny)
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
func (it *ARBITRUMGOVRELAYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMGOVRELAYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMGOVRELAYDeny represents a Deny event raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMGOVRELAYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYDenyIterator{contract: _ARBITRUMGOVRELAY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ARBITRUMGOVRELAYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMGOVRELAYDeny)
				if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) ParseDeny(log types.Log) (*ARBITRUMGOVRELAYDeny, error) {
	event := new(ARBITRUMGOVRELAYDeny)
	if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMGOVRELAYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYRelyIterator struct {
	Event *ARBITRUMGOVRELAYRely // Event containing the contract specifics and raw log

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
func (it *ARBITRUMGOVRELAYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMGOVRELAYRely)
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
		it.Event = new(ARBITRUMGOVRELAYRely)
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
func (it *ARBITRUMGOVRELAYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMGOVRELAYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMGOVRELAYRely represents a Rely event raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ARBITRUMGOVRELAYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYRelyIterator{contract: _ARBITRUMGOVRELAY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ARBITRUMGOVRELAYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMGOVRELAYRely)
				if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) ParseRely(log types.Log) (*ARBITRUMGOVRELAYRely, error) {
	event := new(ARBITRUMGOVRELAYRely)
	if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ARBITRUMGOVRELAYTxToL2Iterator is returned from FilterTxToL2 and is used to iterate over the raw logs and unpacked data for TxToL2 events raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYTxToL2Iterator struct {
	Event *ARBITRUMGOVRELAYTxToL2 // Event containing the contract specifics and raw log

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
func (it *ARBITRUMGOVRELAYTxToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ARBITRUMGOVRELAYTxToL2)
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
		it.Event = new(ARBITRUMGOVRELAYTxToL2)
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
func (it *ARBITRUMGOVRELAYTxToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ARBITRUMGOVRELAYTxToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ARBITRUMGOVRELAYTxToL2 represents a TxToL2 event raised by the ARBITRUMGOVRELAY contract.
type ARBITRUMGOVRELAYTxToL2 struct {
	From   common.Address
	To     common.Address
	SeqNum *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxToL2 is a free log retrieval operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) FilterTxToL2(opts *bind.FilterOpts, from []common.Address, to []common.Address, seqNum []*big.Int) (*ARBITRUMGOVRELAYTxToL2Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.FilterLogs(opts, "TxToL2", fromRule, toRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return &ARBITRUMGOVRELAYTxToL2Iterator{contract: _ARBITRUMGOVRELAY.contract, event: "TxToL2", logs: logs, sub: sub}, nil
}

// WatchTxToL2 is a free log subscription operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) WatchTxToL2(opts *bind.WatchOpts, sink chan<- *ARBITRUMGOVRELAYTxToL2, from []common.Address, to []common.Address, seqNum []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var seqNumRule []interface{}
	for _, seqNumItem := range seqNum {
		seqNumRule = append(seqNumRule, seqNumItem)
	}

	logs, sub, err := _ARBITRUMGOVRELAY.contract.WatchLogs(opts, "TxToL2", fromRule, toRule, seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ARBITRUMGOVRELAYTxToL2)
				if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "TxToL2", log); err != nil {
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

// ParseTxToL2 is a log parse operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed from, address indexed to, uint256 indexed seqNum, bytes data)
func (_ARBITRUMGOVRELAY *ARBITRUMGOVRELAYFilterer) ParseTxToL2(log types.Log) (*ARBITRUMGOVRELAYTxToL2, error) {
	event := new(ARBITRUMGOVRELAYTxToL2)
	if err := _ARBITRUMGOVRELAY.contract.UnpackLog(event, "TxToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
