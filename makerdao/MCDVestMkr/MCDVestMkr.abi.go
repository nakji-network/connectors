// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDVestMkr

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

// MCDVESTMKRABI is the input ABI used to generate the binding from.
const MCDVESTMKRABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gem\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Restrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Unrestrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Vest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"Yank\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TWENTY_YEARS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"accrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"awards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"bgn\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"clf\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"fin\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mgr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"res\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"tot\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rxd\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"bgn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"clf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bgn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tau\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eta\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mgr\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"fin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"contractMintLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"mgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dst\",\"type\":\"address\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"res\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"restrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rxd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unpaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unrestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"usr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmt\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDVESTMKR is an auto generated Go binding around an Ethereum contract.
type MCDVESTMKR struct {
	MCDVESTMKRCaller     // Read-only binding to the contract
	MCDVESTMKRTransactor // Write-only binding to the contract
	MCDVESTMKRFilterer   // Log filterer for contract events
}

// MCDVESTMKRCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDVESTMKRCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDVESTMKRTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDVESTMKRFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTMKRSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDVESTMKRSession struct {
	Contract     *MCDVESTMKR       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDVESTMKRCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDVESTMKRCallerSession struct {
	Contract *MCDVESTMKRCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MCDVESTMKRTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDVESTMKRTransactorSession struct {
	Contract     *MCDVESTMKRTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MCDVESTMKRRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDVESTMKRRaw struct {
	Contract *MCDVESTMKR // Generic contract binding to access the raw methods on
}

// MCDVESTMKRCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDVESTMKRCallerRaw struct {
	Contract *MCDVESTMKRCaller // Generic read-only contract binding to access the raw methods on
}

// MCDVESTMKRTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDVESTMKRTransactorRaw struct {
	Contract *MCDVESTMKRTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDVESTMKR creates a new instance of MCDVESTMKR, bound to a specific deployed contract.
func NewMCDVESTMKR(address common.Address, backend bind.ContractBackend) (*MCDVESTMKR, error) {
	contract, err := bindMCDVESTMKR(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKR{MCDVESTMKRCaller: MCDVESTMKRCaller{contract: contract}, MCDVESTMKRTransactor: MCDVESTMKRTransactor{contract: contract}, MCDVESTMKRFilterer: MCDVESTMKRFilterer{contract: contract}}, nil
}

// NewMCDVESTMKRCaller creates a new read-only instance of MCDVESTMKR, bound to a specific deployed contract.
func NewMCDVESTMKRCaller(address common.Address, caller bind.ContractCaller) (*MCDVESTMKRCaller, error) {
	contract, err := bindMCDVESTMKR(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRCaller{contract: contract}, nil
}

// NewMCDVESTMKRTransactor creates a new write-only instance of MCDVESTMKR, bound to a specific deployed contract.
func NewMCDVESTMKRTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDVESTMKRTransactor, error) {
	contract, err := bindMCDVESTMKR(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRTransactor{contract: contract}, nil
}

// NewMCDVESTMKRFilterer creates a new log filterer instance of MCDVESTMKR, bound to a specific deployed contract.
func NewMCDVESTMKRFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDVESTMKRFilterer, error) {
	contract, err := bindMCDVESTMKR(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRFilterer{contract: contract}, nil
}

// bindMCDVESTMKR binds a generic wrapper to an already deployed contract.
func bindMCDVESTMKR(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDVESTMKRABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTMKR *MCDVESTMKRRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTMKR.Contract.MCDVESTMKRCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTMKR *MCDVESTMKRRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.MCDVESTMKRTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTMKR *MCDVESTMKRRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.MCDVESTMKRTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTMKR *MCDVESTMKRCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTMKR.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTMKR *MCDVESTMKRTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTMKR *MCDVESTMKRTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.contract.Transact(opts, method, params...)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) TWENTYYEARS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "TWENTY_YEARS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTMKR.Contract.TWENTYYEARS(&_MCDVESTMKR.CallOpts)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTMKR.Contract.TWENTYYEARS(&_MCDVESTMKR.CallOpts)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRCaller) Accrued(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "accrued", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Accrued(&_MCDVESTMKR.CallOpts, _id)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Accrued(&_MCDVESTMKR.CallOpts, _id)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTMKR *MCDVESTMKRCaller) Awards(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _MCDVESTMKR.contract.Call(opts, &out, "awards", arg0)

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
func (_MCDVESTMKR *MCDVESTMKRSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTMKR.Contract.Awards(&_MCDVESTMKR.CallOpts, arg0)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTMKR.Contract.Awards(&_MCDVESTMKR.CallOpts, arg0)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Bgn(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "bgn", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Bgn(&_MCDVESTMKR.CallOpts, _id)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Bgn(&_MCDVESTMKR.CallOpts, _id)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Cap() (*big.Int, error) {
	return _MCDVESTMKR.Contract.Cap(&_MCDVESTMKR.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Cap() (*big.Int, error) {
	return _MCDVESTMKR.Contract.Cap(&_MCDVESTMKR.CallOpts)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Clf(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "clf", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Clf(&_MCDVESTMKR.CallOpts, _id)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Clf(&_MCDVESTMKR.CallOpts, _id)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Fin(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "fin", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Fin(&_MCDVESTMKR.CallOpts, _id)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Fin(&_MCDVESTMKR.CallOpts, _id)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCaller) Gem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "gem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKR *MCDVESTMKRSession) Gem() (common.Address, error) {
	return _MCDVESTMKR.Contract.Gem(&_MCDVESTMKR.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x7bd2bea7.
//
// Solidity: function gem() view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Gem() (common.Address, error) {
	return _MCDVESTMKR.Contract.Gem(&_MCDVESTMKR.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Ids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "ids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Ids() (*big.Int, error) {
	return _MCDVESTMKR.Contract.Ids(&_MCDVESTMKR.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Ids() (*big.Int, error) {
	return _MCDVESTMKR.Contract.Ids(&_MCDVESTMKR.CallOpts)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCaller) Mgr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "mgr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKR.Contract.Mgr(&_MCDVESTMKR.CallOpts, _id)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKR.Contract.Mgr(&_MCDVESTMKR.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Res(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "res", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Res(&_MCDVESTMKR.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Res(&_MCDVESTMKR.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Rxd(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "rxd", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Rxd(&_MCDVESTMKR.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Rxd(&_MCDVESTMKR.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Tot(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "tot", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Tot(&_MCDVESTMKR.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Tot(&_MCDVESTMKR.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRCaller) Unpaid(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "unpaid", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Unpaid(&_MCDVESTMKR.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Unpaid(&_MCDVESTMKR.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCaller) Usr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "usr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKR.Contract.Usr(&_MCDVESTMKR.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTMKR.Contract.Usr(&_MCDVESTMKR.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKR *MCDVESTMKRCaller) Valid(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "valid", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKR *MCDVESTMKRSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTMKR.Contract.Valid(&_MCDVESTMKR.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTMKR.Contract.Valid(&_MCDVESTMKR.CallOpts, _id)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTMKR.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Wards(&_MCDVESTMKR.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTMKR *MCDVESTMKRCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTMKR.Contract.Wards(&_MCDVESTMKR.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKR *MCDVESTMKRTransactor) Create(opts *bind.TransactOpts, _usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "create", _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKR *MCDVESTMKRSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Create(&_MCDVESTMKR.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Create(&_MCDVESTMKR.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Deny(&_MCDVESTMKR.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Deny(&_MCDVESTMKR.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.File(&_MCDVESTMKR.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.File(&_MCDVESTMKR.TransactOpts, what, data)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Move(opts *bind.TransactOpts, _id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "move", _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Move(&_MCDVESTMKR.TransactOpts, _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Move(&_MCDVESTMKR.TransactOpts, _id, _dst)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Rely(&_MCDVESTMKR.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Rely(&_MCDVESTMKR.TransactOpts, usr)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Restrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "restrict", _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Restrict(&_MCDVESTMKR.TransactOpts, _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Restrict(&_MCDVESTMKR.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Unrestrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "unrestrict", _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Unrestrict(&_MCDVESTMKR.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Unrestrict(&_MCDVESTMKR.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Vest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "vest", _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Vest(&_MCDVESTMKR.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Vest(&_MCDVESTMKR.TransactOpts, _id)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Vest0(opts *bind.TransactOpts, _id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "vest0", _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Vest0(&_MCDVESTMKR.TransactOpts, _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Vest0(&_MCDVESTMKR.TransactOpts, _id, _maxAmt)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Yank(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "yank", _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Yank(&_MCDVESTMKR.TransactOpts, _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Yank(&_MCDVESTMKR.TransactOpts, _id)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactor) Yank0(opts *bind.TransactOpts, _id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.contract.Transact(opts, "yank0", _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKR *MCDVESTMKRSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Yank0(&_MCDVESTMKR.TransactOpts, _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTMKR *MCDVESTMKRTransactorSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTMKR.Contract.Yank0(&_MCDVESTMKR.TransactOpts, _id, _end)
}

// MCDVESTMKRDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDVESTMKR contract.
type MCDVESTMKRDenyIterator struct {
	Event *MCDVESTMKRDeny // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRDeny)
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
		it.Event = new(MCDVESTMKRDeny)
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
func (it *MCDVESTMKRDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRDeny represents a Deny event raised by the MCDVESTMKR contract.
type MCDVESTMKRDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTMKRDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRDenyIterator{contract: _MCDVESTMKR.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRDeny)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseDeny(log types.Log) (*MCDVESTMKRDeny, error) {
	event := new(MCDVESTMKRDeny)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDVESTMKR contract.
type MCDVESTMKRFileIterator struct {
	Event *MCDVESTMKRFile // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRFile)
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
		it.Event = new(MCDVESTMKRFile)
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
func (it *MCDVESTMKRFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRFile represents a File event raised by the MCDVESTMKR contract.
type MCDVESTMKRFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDVESTMKRFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRFileIterator{contract: _MCDVESTMKR.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRFile)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseFile(log types.Log) (*MCDVESTMKRFile, error) {
	event := new(MCDVESTMKRFile)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MCDVESTMKR contract.
type MCDVESTMKRInitIterator struct {
	Event *MCDVESTMKRInit // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRInit)
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
		it.Event = new(MCDVESTMKRInit)
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
func (it *MCDVESTMKRInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRInit represents a Init event raised by the MCDVESTMKR contract.
type MCDVESTMKRInit struct {
	Id  *big.Int
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterInit(opts *bind.FilterOpts, id []*big.Int, usr []common.Address) (*MCDVESTMKRInitIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRInitIterator{contract: _MCDVESTMKR.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRInit, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRInit)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Init", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseInit(log types.Log) (*MCDVESTMKRInit, error) {
	event := new(MCDVESTMKRInit)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the MCDVESTMKR contract.
type MCDVESTMKRMoveIterator struct {
	Event *MCDVESTMKRMove // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRMove)
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
		it.Event = new(MCDVESTMKRMove)
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
func (it *MCDVESTMKRMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRMove represents a Move event raised by the MCDVESTMKR contract.
type MCDVESTMKRMove struct {
	Id  *big.Int
	Dst common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterMove(opts *bind.FilterOpts, id []*big.Int, dst []common.Address) (*MCDVESTMKRMoveIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRMoveIterator{contract: _MCDVESTMKR.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRMove, id []*big.Int, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRMove)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Move", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseMove(log types.Log) (*MCDVESTMKRMove, error) {
	event := new(MCDVESTMKRMove)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDVESTMKR contract.
type MCDVESTMKRRelyIterator struct {
	Event *MCDVESTMKRRely // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRRely)
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
		it.Event = new(MCDVESTMKRRely)
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
func (it *MCDVESTMKRRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRRely represents a Rely event raised by the MCDVESTMKR contract.
type MCDVESTMKRRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTMKRRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRRelyIterator{contract: _MCDVESTMKR.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRRely)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseRely(log types.Log) (*MCDVESTMKRRely, error) {
	event := new(MCDVESTMKRRely)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRRestrictIterator is returned from FilterRestrict and is used to iterate over the raw logs and unpacked data for Restrict events raised by the MCDVESTMKR contract.
type MCDVESTMKRRestrictIterator struct {
	Event *MCDVESTMKRRestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRRestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRRestrict)
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
		it.Event = new(MCDVESTMKRRestrict)
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
func (it *MCDVESTMKRRestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRRestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRRestrict represents a Restrict event raised by the MCDVESTMKR contract.
type MCDVESTMKRRestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestrict is a free log retrieval operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterRestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRRestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRRestrictIterator{contract: _MCDVESTMKR.contract, event: "Restrict", logs: logs, sub: sub}, nil
}

// WatchRestrict is a free log subscription operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchRestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRRestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRRestrict)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Restrict", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseRestrict(log types.Log) (*MCDVESTMKRRestrict, error) {
	event := new(MCDVESTMKRRestrict)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Restrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRUnrestrictIterator is returned from FilterUnrestrict and is used to iterate over the raw logs and unpacked data for Unrestrict events raised by the MCDVESTMKR contract.
type MCDVESTMKRUnrestrictIterator struct {
	Event *MCDVESTMKRUnrestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRUnrestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRUnrestrict)
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
		it.Event = new(MCDVESTMKRUnrestrict)
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
func (it *MCDVESTMKRUnrestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRUnrestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRUnrestrict represents a Unrestrict event raised by the MCDVESTMKR contract.
type MCDVESTMKRUnrestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnrestrict is a free log retrieval operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterUnrestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRUnrestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRUnrestrictIterator{contract: _MCDVESTMKR.contract, event: "Unrestrict", logs: logs, sub: sub}, nil
}

// WatchUnrestrict is a free log subscription operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchUnrestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRUnrestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRUnrestrict)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Unrestrict", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseUnrestrict(log types.Log) (*MCDVESTMKRUnrestrict, error) {
	event := new(MCDVESTMKRUnrestrict)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Unrestrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRVestIterator is returned from FilterVest and is used to iterate over the raw logs and unpacked data for Vest events raised by the MCDVESTMKR contract.
type MCDVESTMKRVestIterator struct {
	Event *MCDVESTMKRVest // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRVestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRVest)
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
		it.Event = new(MCDVESTMKRVest)
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
func (it *MCDVESTMKRVestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRVestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRVest represents a Vest event raised by the MCDVESTMKR contract.
type MCDVESTMKRVest struct {
	Id  *big.Int
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVest is a free log retrieval operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterVest(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRVestIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRVestIterator{contract: _MCDVESTMKR.contract, event: "Vest", logs: logs, sub: sub}, nil
}

// WatchVest is a free log subscription operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchVest(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRVest, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRVest)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Vest", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseVest(log types.Log) (*MCDVESTMKRVest, error) {
	event := new(MCDVESTMKRVest)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Vest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTMKRYankIterator is returned from FilterYank and is used to iterate over the raw logs and unpacked data for Yank events raised by the MCDVESTMKR contract.
type MCDVESTMKRYankIterator struct {
	Event *MCDVESTMKRYank // Event containing the contract specifics and raw log

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
func (it *MCDVESTMKRYankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTMKRYank)
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
		it.Event = new(MCDVESTMKRYank)
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
func (it *MCDVESTMKRYankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTMKRYankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTMKRYank represents a Yank event raised by the MCDVESTMKR contract.
type MCDVESTMKRYank struct {
	Id  *big.Int
	End *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterYank is a free log retrieval operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTMKR *MCDVESTMKRFilterer) FilterYank(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTMKRYankIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.FilterLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTMKRYankIterator{contract: _MCDVESTMKR.contract, event: "Yank", logs: logs, sub: sub}, nil
}

// WatchYank is a free log subscription operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTMKR *MCDVESTMKRFilterer) WatchYank(opts *bind.WatchOpts, sink chan<- *MCDVESTMKRYank, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTMKR.contract.WatchLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTMKRYank)
				if err := _MCDVESTMKR.contract.UnpackLog(event, "Yank", log); err != nil {
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
func (_MCDVESTMKR *MCDVESTMKRFilterer) ParseYank(log types.Log) (*MCDVESTMKRYank, error) {
	event := new(MCDVESTMKRYank)
	if err := _MCDVESTMKR.contract.UnpackLog(event, "Yank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
