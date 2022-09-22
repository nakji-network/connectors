// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_IAM_AUTO_LINE

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

// MCDIAMAUTOLINEABI is the input ABI used to generate the binding from.
const MCDIAMAUTOLINEABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lineNew\",\"type\":\"uint256\"}],\"name\":\"Exec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"name\":\"Setup\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ilk\",\"type\":\"bytes32\"}],\"name\":\"exec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gap\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"ttl\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"last\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"lastInc\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"remIlk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ttl\",\"type\":\"uint256\"}],\"name\":\"setIlk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDIAMAUTOLINE is an auto generated Go binding around an Ethereum contract.
type MCDIAMAUTOLINE struct {
	MCDIAMAUTOLINECaller     // Read-only binding to the contract
	MCDIAMAUTOLINETransactor // Write-only binding to the contract
	MCDIAMAUTOLINEFilterer   // Log filterer for contract events
}

// MCDIAMAUTOLINECaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDIAMAUTOLINECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDIAMAUTOLINETransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDIAMAUTOLINETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDIAMAUTOLINEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDIAMAUTOLINEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDIAMAUTOLINESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDIAMAUTOLINESession struct {
	Contract     *MCDIAMAUTOLINE   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDIAMAUTOLINECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDIAMAUTOLINECallerSession struct {
	Contract *MCDIAMAUTOLINECaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MCDIAMAUTOLINETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDIAMAUTOLINETransactorSession struct {
	Contract     *MCDIAMAUTOLINETransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MCDIAMAUTOLINERaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDIAMAUTOLINERaw struct {
	Contract *MCDIAMAUTOLINE // Generic contract binding to access the raw methods on
}

// MCDIAMAUTOLINECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDIAMAUTOLINECallerRaw struct {
	Contract *MCDIAMAUTOLINECaller // Generic read-only contract binding to access the raw methods on
}

// MCDIAMAUTOLINETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDIAMAUTOLINETransactorRaw struct {
	Contract *MCDIAMAUTOLINETransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDIAMAUTOLINE creates a new instance of MCDIAMAUTOLINE, bound to a specific deployed contract.
func NewMCDIAMAUTOLINE(address common.Address, backend bind.ContractBackend) (*MCDIAMAUTOLINE, error) {
	contract, err := bindMCDIAMAUTOLINE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINE{MCDIAMAUTOLINECaller: MCDIAMAUTOLINECaller{contract: contract}, MCDIAMAUTOLINETransactor: MCDIAMAUTOLINETransactor{contract: contract}, MCDIAMAUTOLINEFilterer: MCDIAMAUTOLINEFilterer{contract: contract}}, nil
}

// NewMCDIAMAUTOLINECaller creates a new read-only instance of MCDIAMAUTOLINE, bound to a specific deployed contract.
func NewMCDIAMAUTOLINECaller(address common.Address, caller bind.ContractCaller) (*MCDIAMAUTOLINECaller, error) {
	contract, err := bindMCDIAMAUTOLINE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINECaller{contract: contract}, nil
}

// NewMCDIAMAUTOLINETransactor creates a new write-only instance of MCDIAMAUTOLINE, bound to a specific deployed contract.
func NewMCDIAMAUTOLINETransactor(address common.Address, transactor bind.ContractTransactor) (*MCDIAMAUTOLINETransactor, error) {
	contract, err := bindMCDIAMAUTOLINE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINETransactor{contract: contract}, nil
}

// NewMCDIAMAUTOLINEFilterer creates a new log filterer instance of MCDIAMAUTOLINE, bound to a specific deployed contract.
func NewMCDIAMAUTOLINEFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDIAMAUTOLINEFilterer, error) {
	contract, err := bindMCDIAMAUTOLINE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINEFilterer{contract: contract}, nil
}

// bindMCDIAMAUTOLINE binds a generic wrapper to an already deployed contract.
func bindMCDIAMAUTOLINE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDIAMAUTOLINEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDIAMAUTOLINE.Contract.MCDIAMAUTOLINECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.MCDIAMAUTOLINETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.MCDIAMAUTOLINETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDIAMAUTOLINE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.contract.Transact(opts, method, params...)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 line, uint256 gap, uint48 ttl, uint48 last, uint48 lastInc)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Line    *big.Int
	Gap     *big.Int
	Ttl     *big.Int
	Last    *big.Int
	LastInc *big.Int
}, error) {
	var out []interface{}
	err := _MCDIAMAUTOLINE.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Line    *big.Int
		Gap     *big.Int
		Ttl     *big.Int
		Last    *big.Int
		LastInc *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Line = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Gap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ttl = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Last = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastInc = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 line, uint256 gap, uint48 ttl, uint48 last, uint48 lastInc)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Ilks(arg0 [32]byte) (struct {
	Line    *big.Int
	Gap     *big.Int
	Ttl     *big.Int
	Last    *big.Int
	LastInc *big.Int
}, error) {
	return _MCDIAMAUTOLINE.Contract.Ilks(&_MCDIAMAUTOLINE.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 line, uint256 gap, uint48 ttl, uint48 last, uint48 lastInc)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECallerSession) Ilks(arg0 [32]byte) (struct {
	Line    *big.Int
	Gap     *big.Int
	Ttl     *big.Int
	Last    *big.Int
	LastInc *big.Int
}, error) {
	return _MCDIAMAUTOLINE.Contract.Ilks(&_MCDIAMAUTOLINE.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDIAMAUTOLINE.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Vat() (common.Address, error) {
	return _MCDIAMAUTOLINE.Contract.Vat(&_MCDIAMAUTOLINE.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECallerSession) Vat() (common.Address, error) {
	return _MCDIAMAUTOLINE.Contract.Vat(&_MCDIAMAUTOLINE.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDIAMAUTOLINE.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDIAMAUTOLINE.Contract.Wards(&_MCDIAMAUTOLINE.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINECallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDIAMAUTOLINE.Contract.Wards(&_MCDIAMAUTOLINE.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Deny(&_MCDIAMAUTOLINE.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Deny(&_MCDIAMAUTOLINE.TransactOpts, usr)
}

// Exec is a paid mutator transaction binding the contract method 0xb5e98b3b.
//
// Solidity: function exec(bytes32 _ilk) returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactor) Exec(opts *bind.TransactOpts, _ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.contract.Transact(opts, "exec", _ilk)
}

// Exec is a paid mutator transaction binding the contract method 0xb5e98b3b.
//
// Solidity: function exec(bytes32 _ilk) returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Exec(_ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Exec(&_MCDIAMAUTOLINE.TransactOpts, _ilk)
}

// Exec is a paid mutator transaction binding the contract method 0xb5e98b3b.
//
// Solidity: function exec(bytes32 _ilk) returns(uint256)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorSession) Exec(_ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Exec(&_MCDIAMAUTOLINE.TransactOpts, _ilk)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Rely(&_MCDIAMAUTOLINE.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.Rely(&_MCDIAMAUTOLINE.TransactOpts, usr)
}

// RemIlk is a paid mutator transaction binding the contract method 0xc465f077.
//
// Solidity: function remIlk(bytes32 ilk) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactor) RemIlk(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.contract.Transact(opts, "remIlk", ilk)
}

// RemIlk is a paid mutator transaction binding the contract method 0xc465f077.
//
// Solidity: function remIlk(bytes32 ilk) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) RemIlk(ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.RemIlk(&_MCDIAMAUTOLINE.TransactOpts, ilk)
}

// RemIlk is a paid mutator transaction binding the contract method 0xc465f077.
//
// Solidity: function remIlk(bytes32 ilk) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorSession) RemIlk(ilk [32]byte) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.RemIlk(&_MCDIAMAUTOLINE.TransactOpts, ilk)
}

// SetIlk is a paid mutator transaction binding the contract method 0x2a48322d.
//
// Solidity: function setIlk(bytes32 ilk, uint256 line, uint256 gap, uint256 ttl) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactor) SetIlk(opts *bind.TransactOpts, ilk [32]byte, line *big.Int, gap *big.Int, ttl *big.Int) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.contract.Transact(opts, "setIlk", ilk, line, gap, ttl)
}

// SetIlk is a paid mutator transaction binding the contract method 0x2a48322d.
//
// Solidity: function setIlk(bytes32 ilk, uint256 line, uint256 gap, uint256 ttl) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINESession) SetIlk(ilk [32]byte, line *big.Int, gap *big.Int, ttl *big.Int) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.SetIlk(&_MCDIAMAUTOLINE.TransactOpts, ilk, line, gap, ttl)
}

// SetIlk is a paid mutator transaction binding the contract method 0x2a48322d.
//
// Solidity: function setIlk(bytes32 ilk, uint256 line, uint256 gap, uint256 ttl) returns()
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINETransactorSession) SetIlk(ilk [32]byte, line *big.Int, gap *big.Int, ttl *big.Int) (*types.Transaction, error) {
	return _MCDIAMAUTOLINE.Contract.SetIlk(&_MCDIAMAUTOLINE.TransactOpts, ilk, line, gap, ttl)
}

// MCDIAMAUTOLINEDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINEDenyIterator struct {
	Event *MCDIAMAUTOLINEDeny // Event containing the contract specifics and raw log

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
func (it *MCDIAMAUTOLINEDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDIAMAUTOLINEDeny)
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
		it.Event = new(MCDIAMAUTOLINEDeny)
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
func (it *MCDIAMAUTOLINEDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDIAMAUTOLINEDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDIAMAUTOLINEDeny represents a Deny event raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINEDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDIAMAUTOLINEDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINEDenyIterator{contract: _MCDIAMAUTOLINE.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDIAMAUTOLINEDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDIAMAUTOLINEDeny)
				if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) ParseDeny(log types.Log) (*MCDIAMAUTOLINEDeny, error) {
	event := new(MCDIAMAUTOLINEDeny)
	if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDIAMAUTOLINEExecIterator is returned from FilterExec and is used to iterate over the raw logs and unpacked data for Exec events raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINEExecIterator struct {
	Event *MCDIAMAUTOLINEExec // Event containing the contract specifics and raw log

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
func (it *MCDIAMAUTOLINEExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDIAMAUTOLINEExec)
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
		it.Event = new(MCDIAMAUTOLINEExec)
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
func (it *MCDIAMAUTOLINEExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDIAMAUTOLINEExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDIAMAUTOLINEExec represents a Exec event raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINEExec struct {
	Ilk     [32]byte
	Line    *big.Int
	LineNew *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExec is a free log retrieval operation binding the contract event 0x2696a4655cbecf97fdc3c9c74f3eba424f2e404790389c2b4e31d2e32129c7dc.
//
// Solidity: event Exec(bytes32 indexed ilk, uint256 line, uint256 lineNew)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) FilterExec(opts *bind.FilterOpts, ilk [][32]byte) (*MCDIAMAUTOLINEExecIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.FilterLogs(opts, "Exec", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINEExecIterator{contract: _MCDIAMAUTOLINE.contract, event: "Exec", logs: logs, sub: sub}, nil
}

// WatchExec is a free log subscription operation binding the contract event 0x2696a4655cbecf97fdc3c9c74f3eba424f2e404790389c2b4e31d2e32129c7dc.
//
// Solidity: event Exec(bytes32 indexed ilk, uint256 line, uint256 lineNew)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) WatchExec(opts *bind.WatchOpts, sink chan<- *MCDIAMAUTOLINEExec, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.WatchLogs(opts, "Exec", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDIAMAUTOLINEExec)
				if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Exec", log); err != nil {
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

// ParseExec is a log parse operation binding the contract event 0x2696a4655cbecf97fdc3c9c74f3eba424f2e404790389c2b4e31d2e32129c7dc.
//
// Solidity: event Exec(bytes32 indexed ilk, uint256 line, uint256 lineNew)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) ParseExec(log types.Log) (*MCDIAMAUTOLINEExec, error) {
	event := new(MCDIAMAUTOLINEExec)
	if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Exec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDIAMAUTOLINERelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINERelyIterator struct {
	Event *MCDIAMAUTOLINERely // Event containing the contract specifics and raw log

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
func (it *MCDIAMAUTOLINERelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDIAMAUTOLINERely)
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
		it.Event = new(MCDIAMAUTOLINERely)
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
func (it *MCDIAMAUTOLINERelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDIAMAUTOLINERelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDIAMAUTOLINERely represents a Rely event raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINERely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDIAMAUTOLINERelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINERelyIterator{contract: _MCDIAMAUTOLINE.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDIAMAUTOLINERely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDIAMAUTOLINERely)
				if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) ParseRely(log types.Log) (*MCDIAMAUTOLINERely, error) {
	event := new(MCDIAMAUTOLINERely)
	if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDIAMAUTOLINERemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINERemoveIterator struct {
	Event *MCDIAMAUTOLINERemove // Event containing the contract specifics and raw log

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
func (it *MCDIAMAUTOLINERemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDIAMAUTOLINERemove)
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
		it.Event = new(MCDIAMAUTOLINERemove)
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
func (it *MCDIAMAUTOLINERemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDIAMAUTOLINERemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDIAMAUTOLINERemove represents a Remove event raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINERemove struct {
	Ilk [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 indexed ilk)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) FilterRemove(opts *bind.FilterOpts, ilk [][32]byte) (*MCDIAMAUTOLINERemoveIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.FilterLogs(opts, "Remove", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINERemoveIterator{contract: _MCDIAMAUTOLINE.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 indexed ilk)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *MCDIAMAUTOLINERemove, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.WatchLogs(opts, "Remove", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDIAMAUTOLINERemove)
				if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Remove", log); err != nil {
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

// ParseRemove is a log parse operation binding the contract event 0xa56fb2a6d4126f324526f0668c53927c0cd8e08f41ba0fe0f2d6090a84bc75c8.
//
// Solidity: event Remove(bytes32 indexed ilk)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) ParseRemove(log types.Log) (*MCDIAMAUTOLINERemove, error) {
	event := new(MCDIAMAUTOLINERemove)
	if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDIAMAUTOLINESetupIterator is returned from FilterSetup and is used to iterate over the raw logs and unpacked data for Setup events raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINESetupIterator struct {
	Event *MCDIAMAUTOLINESetup // Event containing the contract specifics and raw log

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
func (it *MCDIAMAUTOLINESetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDIAMAUTOLINESetup)
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
		it.Event = new(MCDIAMAUTOLINESetup)
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
func (it *MCDIAMAUTOLINESetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDIAMAUTOLINESetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDIAMAUTOLINESetup represents a Setup event raised by the MCDIAMAUTOLINE contract.
type MCDIAMAUTOLINESetup struct {
	Ilk  [32]byte
	Line *big.Int
	Gap  *big.Int
	Ttl  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetup is a free log retrieval operation binding the contract event 0x1326c8f8fb452b5d26a2406d96790efbb561bf5d2686986a30a5cd847eefc673.
//
// Solidity: event Setup(bytes32 indexed ilk, uint256 line, uint256 gap, uint256 ttl)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) FilterSetup(opts *bind.FilterOpts, ilk [][32]byte) (*MCDIAMAUTOLINESetupIterator, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.FilterLogs(opts, "Setup", ilkRule)
	if err != nil {
		return nil, err
	}
	return &MCDIAMAUTOLINESetupIterator{contract: _MCDIAMAUTOLINE.contract, event: "Setup", logs: logs, sub: sub}, nil
}

// WatchSetup is a free log subscription operation binding the contract event 0x1326c8f8fb452b5d26a2406d96790efbb561bf5d2686986a30a5cd847eefc673.
//
// Solidity: event Setup(bytes32 indexed ilk, uint256 line, uint256 gap, uint256 ttl)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) WatchSetup(opts *bind.WatchOpts, sink chan<- *MCDIAMAUTOLINESetup, ilk [][32]byte) (event.Subscription, error) {

	var ilkRule []interface{}
	for _, ilkItem := range ilk {
		ilkRule = append(ilkRule, ilkItem)
	}

	logs, sub, err := _MCDIAMAUTOLINE.contract.WatchLogs(opts, "Setup", ilkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDIAMAUTOLINESetup)
				if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Setup", log); err != nil {
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

// ParseSetup is a log parse operation binding the contract event 0x1326c8f8fb452b5d26a2406d96790efbb561bf5d2686986a30a5cd847eefc673.
//
// Solidity: event Setup(bytes32 indexed ilk, uint256 line, uint256 gap, uint256 ttl)
func (_MCDIAMAUTOLINE *MCDIAMAUTOLINEFilterer) ParseSetup(log types.Log) (*MCDIAMAUTOLINESetup, error) {
	event := new(MCDIAMAUTOLINESetup)
	if err := _MCDIAMAUTOLINE.contract.UnpackLog(event, "Setup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
