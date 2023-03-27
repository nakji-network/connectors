// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDVestMkrTreasury

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

// MCDVESTMKRTREASURYABI is the input ABI used to generate the binding from.
const MCDVESTMKRTREASURYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_czar\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Restrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Unrestrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Vest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"Yank\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TWENTY_YEARS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"accrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"awards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"bgn\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"clf\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"fin\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mgr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"res\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"tot\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rxd\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"bgn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"clf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bgn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tau\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eta\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mgr\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"czar\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"fin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractTokenLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"mgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dst\",\"type\":\"address\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"res\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"restrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rxd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unpaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unrestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"usr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmt\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDVESTMKRTREASURY is an auto generated Go binding around an Ethereum contract.
type MCDVESTMKRTREASURY struct {
	MCDVESTMKRTREASURYCaller     // Read-only binding to the contract
	MCDVESTMKRTREASURYTransactor // Write-only binding to the contract
	MCDVESTMKRTREASURYFilterer   // Log filterer for contract events
}

// MCDVESTMKRTREASURYCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDVESTMKRTREASURYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRTREASURYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDVESTMKRTREASURYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRTREASURYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDVESTMKRTREASURYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRTREASURYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDVESTMKRTREASURYSession struct {
	Contract     *MCDVESTMKRTREASURY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MCDVESTMKRTREASURYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDVESTMKRTREASURYCallerSession struct {
	Contract *MCDVESTMKRTREASURYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MCDVESTMKRTREASURYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDVESTMKRTREASURYTransactorSession struct {
	Contract     *MCDVESTMKRTREASURYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MCDVESTMKRTREASURYRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDVESTMKRTREASURYRaw struct {
	Contract *MCDVESTMKRTREASURY // Generic contract binding to access the raw methods on
}

// MCDVESTMKRTREASURYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDVESTMKRTREASURYCallerRaw struct {
	Contract *MCDVESTMKRTREASURYCaller // Generic read-only contract binding to access the raw methods on
}

// MCDVESTMKRTREASURYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDVESTMKRTREASURYTransactorRaw struct {
	Contract *MCDVESTMKRTREASURYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDVESTMKRTREASURY creates a new instance of MCDVESTMKRTREASURY, bound to a specific deployed contract.
func NewMCDVESTMKRTREASURY(address common.Address, backend bind.ContractBackend) (*MCDVESTMKRTREASURY, error) {
	contract, err := bindMCDVESTMKRTREASURY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURY{MCDVESTMKRTREASURYCaller: MCDVESTMKRTREASURYCaller{contract: contract}, MCDVESTMKRTREASURYTransactor: MCDVESTMKRTREASURYTransactor{contract: contract}, MCDVESTMKRTREASURYFilterer: MCDVESTMKRTREASURYFilterer{contract: contract}}, nil
}

// NewMCDVESTMKRTREASURYCaller creates a new read-only instance of MCDVESTMKRTREASURY, bound to a specific deployed contract.
func NewMCDVESTMKRTREASURYCaller(address common.Address, caller bind.ContractCaller) (*MCDVESTMKRTREASURYCaller, error) {
	contract, err := bindMCDVESTMKRTREASURY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYCaller{contract: contract}, nil
}

// NewMCDVESTMKRTREASURYTransactor creates a new write-only instance of MCDVESTMKRTREASURY, bound to a specific deployed contract.
func NewMCDVESTMKRTREASURYTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDVESTMKRTREASURYTransactor, error) {
	contract, err := bindMCDVESTMKRTREASURY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYTransactor{contract: contract}, nil
}

// NewMCDVESTMKRTREASURYFilterer creates a new log filterer instance of MCDVESTMKRTREASURY, bound to a specific deployed contract.
func NewMCDVESTMKRTREASURYFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDVESTMKRTREASURYFilterer, error) {
	contract, err := bindMCDVESTMKRTREASURY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYFilterer{contract: contract}, nil
}

// bindMCDVESTMKRTREASURY binds a generic wrapper to an already deployed contract.
func bindMCDVESTMKRTREASURY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDVESTMKRTREASURYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTMKRTREASURY.Contract.MCDVESTMKRTREASURYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.MCDVESTMKRTREASURYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.MCDVESTMKRTREASURYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTMKRTREASURY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.contract.Transact(opts, method, params...)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) TWENTYYEARS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "TWENTY_YEARS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.TWENTYYEARS(&_MCDVESTMKRTREASURY.CallOpts)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.TWENTYYEARS(&_MCDVESTMKRTREASURY.CallOpts)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Accrued(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "accrued", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Accrued(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Accrued(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Awards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "awards", arg0)

	outstruct := new(struct {
		Usr common.Address
		Bgn *big.Int
		Clf *big.Int
		Fin *big.Int
		Mgr common.Address
		Res uint8
		Tot *big.Int
		Rxd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Usr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Bgn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Clf = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fin = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Mgr = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Res = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Tot = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Rxd = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTMKRTREASURY.Contract.Awards(&_MCDVESTMKRTREASURY.CallOpts, arg0)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTMKRTREASURY.Contract.Awards(&_MCDVESTMKRTREASURY.CallOpts, arg0)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Bgn(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "bgn", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Bgn(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Bgn(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Cap() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Cap(&_MCDVESTMKRTREASURY.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Cap() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Cap(&_MCDVESTMKRTREASURY.CallOpts)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Clf(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "clf", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Clf(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Clf(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Czar is a free data retrieval call binding the contract method 0x1a8d3a6c.
//
// Solidity: function czar() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Czar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "czar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Czar is a free data retrieval call binding the contract method 0x1a8d3a6c.
//
// Solidity: function czar() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Czar() (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Czar(&_MCDVESTMKRTREASURY.CallOpts)
}

// Czar is a free data retrieval call binding the contract method 0x1a8d3a6c.
//
// Solidity: function czar() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Czar() (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Czar(&_MCDVESTMKRTREASURY.CallOpts)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Fin(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "fin", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Fin(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Fin(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Gem() (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Gem(&_MCDVESTMKRTREASURY.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Gem() (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Gem(&_MCDVESTMKRTREASURY.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Ids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "ids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Ids() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Ids(&_MCDVESTMKRTREASURY.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Ids() (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Ids(&_MCDVESTMKRTREASURY.CallOpts)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Mgr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "mgr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Mgr(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Mgr(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Res(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "res", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Res(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Res(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Rxd(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "rxd", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Rxd(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Rxd(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Tot(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "tot", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Tot(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Tot(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Unpaid(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "unpaid", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Unpaid(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Unpaid(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Usr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "usr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Usr(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKRTREASURY.Contract.Usr(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Valid(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "valid", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTMKRTREASURY.Contract.Valid(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTMKRTREASURY.Contract.Valid(&_MCDVESTMKRTREASURY.CallOpts, _id)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKRTREASURY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Wards(&_MCDVESTMKRTREASURY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTMKRTREASURY.Contract.Wards(&_MCDVESTMKRTREASURY.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Create(opts *bind.TransactOpts, _usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "create", _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Create(&_MCDVESTMKRTREASURY.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Create(&_MCDVESTMKRTREASURY.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Deny(&_MCDVESTMKRTREASURY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Deny(&_MCDVESTMKRTREASURY.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.File(&_MCDVESTMKRTREASURY.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.File(&_MCDVESTMKRTREASURY.TransactOpts, what, data)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Move(opts *bind.TransactOpts, _id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "move", _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Move(&_MCDVESTMKRTREASURY.TransactOpts, _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Move(&_MCDVESTMKRTREASURY.TransactOpts, _id, _dst)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Rely(&_MCDVESTMKRTREASURY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Rely(&_MCDVESTMKRTREASURY.TransactOpts, usr)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Restrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "restrict", _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Restrict(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Restrict(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Unrestrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "unrestrict", _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Unrestrict(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Unrestrict(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Vest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "vest", _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Vest(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Vest(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Vest0(opts *bind.TransactOpts, _id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "vest0", _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Vest0(&_MCDVESTMKRTREASURY.TransactOpts, _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Vest0(&_MCDVESTMKRTREASURY.TransactOpts, _id, _maxAmt)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Yank(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "yank", _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Yank(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Yank(&_MCDVESTMKRTREASURY.TransactOpts, _id)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactor) Yank0(opts *bind.TransactOpts, _id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.contract.Transact(opts, "yank0", _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Yank0(&_MCDVESTMKRTREASURY.TransactOpts, _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYTransactorSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKRTREASURY.Contract.Yank0(&_MCDVESTMKRTREASURY.TransactOpts, _id, _end)
}

// MCDVESTMKRTREASURYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYDenyIterator struct {
	Event *MCDVESTMKRTREASURYDeny // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYDeny)
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
		it.Event = new(MCDVESTMKRTREASURYDeny)
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
func (it *MCDVESTMKRTREASURYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYDeny represents a Deny event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTMKRTREASURYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYDenyIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYDeny)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseDeny(log types.Log) (*MCDVESTMKRTREASURYDeny, error) {
	event := new(MCDVESTMKRTREASURYDeny)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYFileIterator struct {
	Event *MCDVESTMKRTREASURYFile // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYFile)
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
		it.Event = new(MCDVESTMKRTREASURYFile)
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
func (it *MCDVESTMKRTREASURYFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYFile represents a File event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDVESTMKRTREASURYFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYFileIterator{contract: _MCDVESTMKRTREASURY.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYFile)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseFile(log types.Log) (*MCDVESTMKRTREASURYFile, error) {
	event := new(MCDVESTMKRTREASURYFile)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYInitIterator struct {
	Event *MCDVESTMKRTREASURYInit // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYInit)
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
		it.Event = new(MCDVESTMKRTREASURYInit)
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
func (it *MCDVESTMKRTREASURYInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYInit represents a Init event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYInit struct {
	Id  *big.Int
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterInit(opts *bind.FilterOpts, id []*big.Int, usr []common.Address) (*MCDVESTMKRTREASURYInitIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYInitIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYInit, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYInit)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseInit(log types.Log) (*MCDVESTMKRTREASURYInit, error) {
	event := new(MCDVESTMKRTREASURYInit)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYMoveIterator struct {
	Event *MCDVESTMKRTREASURYMove // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYMove)
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
		it.Event = new(MCDVESTMKRTREASURYMove)
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
func (it *MCDVESTMKRTREASURYMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYMove represents a Move event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYMove struct {
	Id  *big.Int
	Dst common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterMove(opts *bind.FilterOpts, id []*big.Int, dst []common.Address) (*MCDVESTMKRTREASURYMoveIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYMoveIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYMove, id []*big.Int, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYMove)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Move", log); err != nil {
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

// ParseMove is a log parse operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseMove(log types.Log) (*MCDVESTMKRTREASURYMove, error) {
	event := new(MCDVESTMKRTREASURYMove)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYRelyIterator struct {
	Event *MCDVESTMKRTREASURYRely // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYRely)
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
		it.Event = new(MCDVESTMKRTREASURYRely)
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
func (it *MCDVESTMKRTREASURYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYRely represents a Rely event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTMKRTREASURYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYRelyIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYRely)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseRely(log types.Log) (*MCDVESTMKRTREASURYRely, error) {
	event := new(MCDVESTMKRTREASURYRely)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYRestrictIterator is returned from FilterRestrict and is used to iterate over the raw logs and unpacked data for Restrict events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYRestrictIterator struct {
	Event *MCDVESTMKRTREASURYRestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYRestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYRestrict)
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
		it.Event = new(MCDVESTMKRTREASURYRestrict)
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
func (it *MCDVESTMKRTREASURYRestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYRestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYRestrict represents a Restrict event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYRestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestrict is a free log retrieval operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterRestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRTREASURYRestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYRestrictIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Restrict", logs: logs, sub: sub}, nil
}

// WatchRestrict is a free log subscription operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchRestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYRestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYRestrict)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Restrict", log); err != nil {
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

// ParseRestrict is a log parse operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseRestrict(log types.Log) (*MCDVESTMKRTREASURYRestrict, error) {
	event := new(MCDVESTMKRTREASURYRestrict)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Restrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYUnrestrictIterator is returned from FilterUnrestrict and is used to iterate over the raw logs and unpacked data for Unrestrict events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYUnrestrictIterator struct {
	Event *MCDVESTMKRTREASURYUnrestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYUnrestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYUnrestrict)
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
		it.Event = new(MCDVESTMKRTREASURYUnrestrict)
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
func (it *MCDVESTMKRTREASURYUnrestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYUnrestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYUnrestrict represents a Unrestrict event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYUnrestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnrestrict is a free log retrieval operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterUnrestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRTREASURYUnrestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYUnrestrictIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Unrestrict", logs: logs, sub: sub}, nil
}

// WatchUnrestrict is a free log subscription operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchUnrestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYUnrestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYUnrestrict)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Unrestrict", log); err != nil {
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

// ParseUnrestrict is a log parse operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseUnrestrict(log types.Log) (*MCDVESTMKRTREASURYUnrestrict, error) {
	event := new(MCDVESTMKRTREASURYUnrestrict)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Unrestrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYVestIterator is returned from FilterVest and is used to iterate over the raw logs and unpacked data for Vest events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYVestIterator struct {
	Event *MCDVESTMKRTREASURYVest // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYVestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYVest)
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
		it.Event = new(MCDVESTMKRTREASURYVest)
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
func (it *MCDVESTMKRTREASURYVestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYVestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYVest represents a Vest event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYVest struct {
	Id  *big.Int
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVest is a free log retrieval operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterVest(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRTREASURYVestIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYVestIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Vest", logs: logs, sub: sub}, nil
}

// WatchVest is a free log subscription operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchVest(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYVest, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYVest)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Vest", log); err != nil {
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

// ParseVest is a log parse operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseVest(log types.Log) (*MCDVESTMKRTREASURYVest, error) {
	event := new(MCDVESTMKRTREASURYVest)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Vest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRTREASURYYankIterator is returned from FilterYank and is used to iterate over the raw logs and unpacked data for Yank events raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYYankIterator struct {
	Event *MCDVESTMKRTREASURYYank // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRTREASURYYankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRTREASURYYank)
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
		it.Event = new(MCDVESTMKRTREASURYYank)
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
func (it *MCDVESTMKRTREASURYYankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRTREASURYYankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRTREASURYYank represents a Yank event raised by the MCDVESTMKRTREASURY contract.
type MCDVESTMKRTREASURYYank struct {
	Id  *big.Int
	End *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterYank is a free log retrieval operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) FilterYank(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRTREASURYYankIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.FilterLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTREASURYYankIterator{contract: _MCDVESTMKRTREASURY.contract, event: "Yank", logs: logs, sub: sub}, nil
}

// WatchYank is a free log subscription operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) WatchYank(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRTREASURYYank, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKRTREASURY.contract.WatchLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRTREASURYYank)
				if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Yank", log); err != nil {
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

// ParseYank is a log parse operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTMKRTREASURY *MCDVESTMKRTREASURYFilterer) ParseYank(log types.Log) (*MCDVESTMKRTREASURYYank, error) {
	event := new(MCDVESTMKRTREASURYYank)
	if err := _MCDVESTMKRTREASURY.contract.UnpackLog(event, "Yank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
