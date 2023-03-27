// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OptimismL2DaiTokenBridge

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

// OPTIMISML2DAITOKENBRIDGEABI is the input ABI used to generate the binding from.
const OPTIMISML2DAITOKENBRIDGEABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OPTIMISML2DAITOKENBRIDGE is an auto generated Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGE struct {
	OPTIMISML2DAITOKENBRIDGECaller     // Read-only binding to the contract
	OPTIMISML2DAITOKENBRIDGETransactor // Write-only binding to the contract
	OPTIMISML2DAITOKENBRIDGEFilterer   // Log filterer for contract events
}

// OPTIMISML2DAITOKENBRIDGECaller is an auto generated read-only Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAITOKENBRIDGETransactor is an auto generated write-only Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAITOKENBRIDGEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OPTIMISML2DAITOKENBRIDGEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OPTIMISML2DAITOKENBRIDGESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OPTIMISML2DAITOKENBRIDGESession struct {
	Contract     *OPTIMISML2DAITOKENBRIDGE // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OPTIMISML2DAITOKENBRIDGECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OPTIMISML2DAITOKENBRIDGECallerSession struct {
	Contract *OPTIMISML2DAITOKENBRIDGECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// OPTIMISML2DAITOKENBRIDGETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OPTIMISML2DAITOKENBRIDGETransactorSession struct {
	Contract     *OPTIMISML2DAITOKENBRIDGETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// OPTIMISML2DAITOKENBRIDGERaw is an auto generated low-level Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGERaw struct {
	Contract *OPTIMISML2DAITOKENBRIDGE // Generic contract binding to access the raw methods on
}

// OPTIMISML2DAITOKENBRIDGECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGECallerRaw struct {
	Contract *OPTIMISML2DAITOKENBRIDGECaller // Generic read-only contract binding to access the raw methods on
}

// OPTIMISML2DAITOKENBRIDGETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OPTIMISML2DAITOKENBRIDGETransactorRaw struct {
	Contract *OPTIMISML2DAITOKENBRIDGETransactor // Generic write-only contract binding to access the raw methods on
}

// NewOPTIMISML2DAITOKENBRIDGE creates a new instance of OPTIMISML2DAITOKENBRIDGE, bound to a specific deployed contract.
func NewOPTIMISML2DAITOKENBRIDGE(address common.Address, backend bind.ContractBackend) (*OPTIMISML2DAITOKENBRIDGE, error) {
	contract, err := bindOPTIMISML2DAITOKENBRIDGE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGE{OPTIMISML2DAITOKENBRIDGECaller: OPTIMISML2DAITOKENBRIDGECaller{contract: contract}, OPTIMISML2DAITOKENBRIDGETransactor: OPTIMISML2DAITOKENBRIDGETransactor{contract: contract}, OPTIMISML2DAITOKENBRIDGEFilterer: OPTIMISML2DAITOKENBRIDGEFilterer{contract: contract}}, nil
}

// NewOPTIMISML2DAITOKENBRIDGECaller creates a new read-only instance of OPTIMISML2DAITOKENBRIDGE, bound to a specific deployed contract.
func NewOPTIMISML2DAITOKENBRIDGECaller(address common.Address, caller bind.ContractCaller) (*OPTIMISML2DAITOKENBRIDGECaller, error) {
	contract, err := bindOPTIMISML2DAITOKENBRIDGE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGECaller{contract: contract}, nil
}

// NewOPTIMISML2DAITOKENBRIDGETransactor creates a new write-only instance of OPTIMISML2DAITOKENBRIDGE, bound to a specific deployed contract.
func NewOPTIMISML2DAITOKENBRIDGETransactor(address common.Address, transactor bind.ContractTransactor) (*OPTIMISML2DAITOKENBRIDGETransactor, error) {
	contract, err := bindOPTIMISML2DAITOKENBRIDGE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGETransactor{contract: contract}, nil
}

// NewOPTIMISML2DAITOKENBRIDGEFilterer creates a new log filterer instance of OPTIMISML2DAITOKENBRIDGE, bound to a specific deployed contract.
func NewOPTIMISML2DAITOKENBRIDGEFilterer(address common.Address, filterer bind.ContractFilterer) (*OPTIMISML2DAITOKENBRIDGEFilterer, error) {
	contract, err := bindOPTIMISML2DAITOKENBRIDGE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGEFilterer{contract: contract}, nil
}

// bindOPTIMISML2DAITOKENBRIDGE binds a generic wrapper to an already deployed contract.
func bindOPTIMISML2DAITOKENBRIDGE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OPTIMISML2DAITOKENBRIDGEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.OPTIMISML2DAITOKENBRIDGECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.OPTIMISML2DAITOKENBRIDGETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.OPTIMISML2DAITOKENBRIDGETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.contract.Transact(opts, method, params...)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGECaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OPTIMISML2DAITOKENBRIDGE.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGESession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Wards(&_OPTIMISML2DAITOKENBRIDGE.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGECallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Wards(&_OPTIMISML2DAITOKENBRIDGE.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactor) Approve(opts *bind.TransactOpts, token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.contract.Transact(opts, "approve", token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGESession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Approve(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, token, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address spender, uint256 value) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactorSession) Approve(token common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Approve(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, token, spender, value)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGESession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Deny(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Deny(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGESession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Rely(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGETransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _OPTIMISML2DAITOKENBRIDGE.Contract.Rely(&_OPTIMISML2DAITOKENBRIDGE.TransactOpts, usr)
}

// OPTIMISML2DAITOKENBRIDGEApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGEApproveIterator struct {
	Event *OPTIMISML2DAITOKENBRIDGEApprove // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAITOKENBRIDGEApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAITOKENBRIDGEApprove)
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
		it.Event = new(OPTIMISML2DAITOKENBRIDGEApprove)
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
func (it *OPTIMISML2DAITOKENBRIDGEApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAITOKENBRIDGEApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAITOKENBRIDGEApprove represents a Approve event raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGEApprove struct {
	Token   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) FilterApprove(opts *bind.FilterOpts, token []common.Address, spender []common.Address) (*OPTIMISML2DAITOKENBRIDGEApproveIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.FilterLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGEApproveIterator{contract: _OPTIMISML2DAITOKENBRIDGE.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAITOKENBRIDGEApprove, token []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.WatchLogs(opts, "Approve", tokenRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAITOKENBRIDGEApprove)
				if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Approve", log); err != nil {
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

// ParseApprove is a log parse operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed token, address indexed spender, uint256 value)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) ParseApprove(log types.Log) (*OPTIMISML2DAITOKENBRIDGEApprove, error) {
	event := new(OPTIMISML2DAITOKENBRIDGEApprove)
	if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Approve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2DAITOKENBRIDGEDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGEDenyIterator struct {
	Event *OPTIMISML2DAITOKENBRIDGEDeny // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAITOKENBRIDGEDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAITOKENBRIDGEDeny)
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
		it.Event = new(OPTIMISML2DAITOKENBRIDGEDeny)
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
func (it *OPTIMISML2DAITOKENBRIDGEDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAITOKENBRIDGEDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAITOKENBRIDGEDeny represents a Deny event raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGEDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2DAITOKENBRIDGEDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGEDenyIterator{contract: _OPTIMISML2DAITOKENBRIDGE.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAITOKENBRIDGEDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAITOKENBRIDGEDeny)
				if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) ParseDeny(log types.Log) (*OPTIMISML2DAITOKENBRIDGEDeny, error) {
	event := new(OPTIMISML2DAITOKENBRIDGEDeny)
	if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OPTIMISML2DAITOKENBRIDGERelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGERelyIterator struct {
	Event *OPTIMISML2DAITOKENBRIDGERely // Event containing the contract specifics and raw log

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
func (it *OPTIMISML2DAITOKENBRIDGERelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OPTIMISML2DAITOKENBRIDGERely)
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
		it.Event = new(OPTIMISML2DAITOKENBRIDGERely)
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
func (it *OPTIMISML2DAITOKENBRIDGERelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OPTIMISML2DAITOKENBRIDGERelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OPTIMISML2DAITOKENBRIDGERely represents a Rely event raised by the OPTIMISML2DAITOKENBRIDGE contract.
type OPTIMISML2DAITOKENBRIDGERely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*OPTIMISML2DAITOKENBRIDGERelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &OPTIMISML2DAITOKENBRIDGERelyIterator{contract: _OPTIMISML2DAITOKENBRIDGE.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *OPTIMISML2DAITOKENBRIDGERely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _OPTIMISML2DAITOKENBRIDGE.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OPTIMISML2DAITOKENBRIDGERely)
				if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_OPTIMISML2DAITOKENBRIDGE *OPTIMISML2DAITOKENBRIDGEFilterer) ParseRely(log types.Log) (*OPTIMISML2DAITOKENBRIDGERely, error) {
	event := new(OPTIMISML2DAITOKENBRIDGERely)
	if err := _OPTIMISML2DAITOKENBRIDGE.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
