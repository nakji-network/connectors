// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CdpRegistry

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

// CDPREGISTRYABI is the input ABI used to generate the binding from.
const CDPREGISTRYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cdpManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"NewCdpRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cdpManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cdps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CDPREGISTRY is an auto generated Go binding around an Ethereum contract.
type CDPREGISTRY struct {
	CDPREGISTRYCaller     // Read-only binding to the contract
	CDPREGISTRYTransactor // Write-only binding to the contract
	CDPREGISTRYFilterer   // Log filterer for contract events
}

// CDPREGISTRYCaller is an auto generated read-only Go binding around an Ethereum contract.
type CDPREGISTRYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPREGISTRYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CDPREGISTRYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPREGISTRYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CDPREGISTRYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CDPREGISTRYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CDPREGISTRYSession struct {
	Contract     *CDPREGISTRY      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CDPREGISTRYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CDPREGISTRYCallerSession struct {
	Contract *CDPREGISTRYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CDPREGISTRYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CDPREGISTRYTransactorSession struct {
	Contract     *CDPREGISTRYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CDPREGISTRYRaw is an auto generated low-level Go binding around an Ethereum contract.
type CDPREGISTRYRaw struct {
	Contract *CDPREGISTRY // Generic contract binding to access the raw methods on
}

// CDPREGISTRYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CDPREGISTRYCallerRaw struct {
	Contract *CDPREGISTRYCaller // Generic read-only contract binding to access the raw methods on
}

// CDPREGISTRYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CDPREGISTRYTransactorRaw struct {
	Contract *CDPREGISTRYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCDPREGISTRY creates a new instance of CDPREGISTRY, bound to a specific deployed contract.
func NewCDPREGISTRY(address common.Address, backend bind.ContractBackend) (*CDPREGISTRY, error) {
	contract, err := bindCDPREGISTRY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CDPREGISTRY{CDPREGISTRYCaller: CDPREGISTRYCaller{contract: contract}, CDPREGISTRYTransactor: CDPREGISTRYTransactor{contract: contract}, CDPREGISTRYFilterer: CDPREGISTRYFilterer{contract: contract}}, nil
}

// NewCDPREGISTRYCaller creates a new read-only instance of CDPREGISTRY, bound to a specific deployed contract.
func NewCDPREGISTRYCaller(address common.Address, caller bind.ContractCaller) (*CDPREGISTRYCaller, error) {
	contract, err := bindCDPREGISTRY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CDPREGISTRYCaller{contract: contract}, nil
}

// NewCDPREGISTRYTransactor creates a new write-only instance of CDPREGISTRY, bound to a specific deployed contract.
func NewCDPREGISTRYTransactor(address common.Address, transactor bind.ContractTransactor) (*CDPREGISTRYTransactor, error) {
	contract, err := bindCDPREGISTRY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CDPREGISTRYTransactor{contract: contract}, nil
}

// NewCDPREGISTRYFilterer creates a new log filterer instance of CDPREGISTRY, bound to a specific deployed contract.
func NewCDPREGISTRYFilterer(address common.Address, filterer bind.ContractFilterer) (*CDPREGISTRYFilterer, error) {
	contract, err := bindCDPREGISTRY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CDPREGISTRYFilterer{contract: contract}, nil
}

// bindCDPREGISTRY binds a generic wrapper to an already deployed contract.
func bindCDPREGISTRY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CDPREGISTRYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CDPREGISTRY *CDPREGISTRYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CDPREGISTRY.Contract.CDPREGISTRYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CDPREGISTRY *CDPREGISTRYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.CDPREGISTRYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CDPREGISTRY *CDPREGISTRYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.CDPREGISTRYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CDPREGISTRY *CDPREGISTRYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CDPREGISTRY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CDPREGISTRY *CDPREGISTRYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CDPREGISTRY *CDPREGISTRYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.contract.Transact(opts, method, params...)
}

// CdpManager is a free data retrieval call binding the contract method 0xbb038e15.
//
// Solidity: function cdpManager() view returns(address)
func (_CDPREGISTRY *CDPREGISTRYCaller) CdpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CDPREGISTRY.contract.Call(opts, &out, "cdpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CdpManager is a free data retrieval call binding the contract method 0xbb038e15.
//
// Solidity: function cdpManager() view returns(address)
func (_CDPREGISTRY *CDPREGISTRYSession) CdpManager() (common.Address, error) {
	return _CDPREGISTRY.Contract.CdpManager(&_CDPREGISTRY.CallOpts)
}

// CdpManager is a free data retrieval call binding the contract method 0xbb038e15.
//
// Solidity: function cdpManager() view returns(address)
func (_CDPREGISTRY *CDPREGISTRYCallerSession) CdpManager() (common.Address, error) {
	return _CDPREGISTRY.Contract.CdpManager(&_CDPREGISTRY.CallOpts)
}

// Cdps is a free data retrieval call binding the contract method 0x3a6916ad.
//
// Solidity: function cdps(bytes32 , address ) view returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYCaller) Cdps(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CDPREGISTRY.contract.Call(opts, &out, "cdps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cdps is a free data retrieval call binding the contract method 0x3a6916ad.
//
// Solidity: function cdps(bytes32 , address ) view returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYSession) Cdps(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _CDPREGISTRY.Contract.Cdps(&_CDPREGISTRY.CallOpts, arg0, arg1)
}

// Cdps is a free data retrieval call binding the contract method 0x3a6916ad.
//
// Solidity: function cdps(bytes32 , address ) view returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYCallerSession) Cdps(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _CDPREGISTRY.Contract.Cdps(&_CDPREGISTRY.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPREGISTRY *CDPREGISTRYCaller) Ilks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CDPREGISTRY.contract.Call(opts, &out, "ilks", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPREGISTRY *CDPREGISTRYSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _CDPREGISTRY.Contract.Ilks(&_CDPREGISTRY.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0x2c2cb9fd.
//
// Solidity: function ilks(uint256 ) view returns(bytes32)
func (_CDPREGISTRY *CDPREGISTRYCallerSession) Ilks(arg0 *big.Int) ([32]byte, error) {
	return _CDPREGISTRY.Contract.Ilks(&_CDPREGISTRY.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPREGISTRY *CDPREGISTRYCaller) Owns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CDPREGISTRY.contract.Call(opts, &out, "owns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPREGISTRY *CDPREGISTRYSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _CDPREGISTRY.Contract.Owns(&_CDPREGISTRY.CallOpts, arg0)
}

// Owns is a free data retrieval call binding the contract method 0x8161b120.
//
// Solidity: function owns(uint256 ) view returns(address)
func (_CDPREGISTRY *CDPREGISTRYCallerSession) Owns(arg0 *big.Int) (common.Address, error) {
	return _CDPREGISTRY.Contract.Owns(&_CDPREGISTRY.CallOpts, arg0)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYTransactor) Open(opts *bind.TransactOpts, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPREGISTRY.contract.Transact(opts, "open", ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.Open(&_CDPREGISTRY.TransactOpts, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6090dec5.
//
// Solidity: function open(bytes32 ilk, address usr) returns(uint256)
func (_CDPREGISTRY *CDPREGISTRYTransactorSession) Open(ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _CDPREGISTRY.Contract.Open(&_CDPREGISTRY.TransactOpts, ilk, usr)
}

// CDPREGISTRYNewCdpRegisteredIterator is returned from FilterNewCdpRegistered and is used to iterate over the raw logs and unpacked data for NewCdpRegistered events raised by the CDPREGISTRY contract.
type CDPREGISTRYNewCdpRegisteredIterator struct {
	Event *CDPREGISTRYNewCdpRegistered // Event containing the contract specifics and raw log

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
func (it *CDPREGISTRYNewCdpRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CDPREGISTRYNewCdpRegistered)
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
		it.Event = new(CDPREGISTRYNewCdpRegistered)
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
func (it *CDPREGISTRYNewCdpRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CDPREGISTRYNewCdpRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CDPREGISTRYNewCdpRegistered represents a NewCdpRegistered event raised by the CDPREGISTRY contract.
type CDPREGISTRYNewCdpRegistered struct {
	Sender common.Address
	Owner  common.Address
	Cdp    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewCdpRegistered is a free log retrieval operation binding the contract event 0xa5f6af9041f75a8d636848197dfbc0abf5daf6a5fa6b15fb6d59916ddecce312.
//
// Solidity: event NewCdpRegistered(address indexed sender, address indexed owner, uint256 indexed cdp)
func (_CDPREGISTRY *CDPREGISTRYFilterer) FilterNewCdpRegistered(opts *bind.FilterOpts, sender []common.Address, owner []common.Address, cdp []*big.Int) (*CDPREGISTRYNewCdpRegisteredIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _CDPREGISTRY.contract.FilterLogs(opts, "NewCdpRegistered", senderRule, ownerRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return &CDPREGISTRYNewCdpRegisteredIterator{contract: _CDPREGISTRY.contract, event: "NewCdpRegistered", logs: logs, sub: sub}, nil
}

// WatchNewCdpRegistered is a free log subscription operation binding the contract event 0xa5f6af9041f75a8d636848197dfbc0abf5daf6a5fa6b15fb6d59916ddecce312.
//
// Solidity: event NewCdpRegistered(address indexed sender, address indexed owner, uint256 indexed cdp)
func (_CDPREGISTRY *CDPREGISTRYFilterer) WatchNewCdpRegistered(opts *bind.WatchOpts, sink chan<- *CDPREGISTRYNewCdpRegistered, sender []common.Address, owner []common.Address, cdp []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var cdpRule []interface{}
	for _, cdpItem := range cdp {
		cdpRule = append(cdpRule, cdpItem)
	}

	logs, sub, err := _CDPREGISTRY.contract.WatchLogs(opts, "NewCdpRegistered", senderRule, ownerRule, cdpRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CDPREGISTRYNewCdpRegistered)
				if err := _CDPREGISTRY.contract.UnpackLog(event, "NewCdpRegistered", log); err != nil {
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

// ParseNewCdpRegistered is a log parse operation binding the contract event 0xa5f6af9041f75a8d636848197dfbc0abf5daf6a5fa6b15fb6d59916ddecce312.
//
// Solidity: event NewCdpRegistered(address indexed sender, address indexed owner, uint256 indexed cdp)
func (_CDPREGISTRY *CDPREGISTRYFilterer) ParseNewCdpRegistered(log types.Log) (*CDPREGISTRYNewCdpRegistered, error) {
	event := new(CDPREGISTRYNewCdpRegistered)
	if err := _CDPREGISTRY.contract.UnpackLog(event, "NewCdpRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
