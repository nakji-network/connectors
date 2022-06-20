// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_VEST_DAI_LEGACY

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

// MCDVESTDAILEGACYABI is the input ABI used to generate the binding from.
const MCDVESTDAILEGACYABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlog\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Restrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Unrestrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Vest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"Yank\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TWENTY_YEARS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"accrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"awards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"bgn\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"clf\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"fin\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mgr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"res\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"tot\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rxd\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"bgn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainlog\",\"outputs\":[{\"internalType\":\"contractChainlogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"clf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bgn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tau\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eta\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mgr\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"fin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"mgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dst\",\"type\":\"address\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"res\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"restrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rxd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unpaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unrestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"usr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmt\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDVESTDAILEGACY is an auto generated Go binding around an Ethereum contract.
type MCDVESTDAILEGACY struct {
	MCDVESTDAILEGACYCaller     // Read-only binding to the contract
	MCDVESTDAILEGACYTransactor // Write-only binding to the contract
	MCDVESTDAILEGACYFilterer   // Log filterer for contract events
}

// MCDVESTDAILEGACYCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDVESTDAILEGACYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAILEGACYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDVESTDAILEGACYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAILEGACYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDVESTDAILEGACYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAILEGACYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDVESTDAILEGACYSession struct {
	Contract     *MCDVESTDAILEGACY // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDVESTDAILEGACYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDVESTDAILEGACYCallerSession struct {
	Contract *MCDVESTDAILEGACYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MCDVESTDAILEGACYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDVESTDAILEGACYTransactorSession struct {
	Contract     *MCDVESTDAILEGACYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MCDVESTDAILEGACYRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDVESTDAILEGACYRaw struct {
	Contract *MCDVESTDAILEGACY // Generic contract binding to access the raw methods on
}

// MCDVESTDAILEGACYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDVESTDAILEGACYCallerRaw struct {
	Contract *MCDVESTDAILEGACYCaller // Generic read-only contract binding to access the raw methods on
}

// MCDVESTDAILEGACYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDVESTDAILEGACYTransactorRaw struct {
	Contract *MCDVESTDAILEGACYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDVESTDAILEGACY creates a new instance of MCDVESTDAILEGACY, bound to a specific deployed contract.
func NewMCDVESTDAILEGACY(address common.Address, backend bind.ContractBackend) (*MCDVESTDAILEGACY, error) {
	contract, err := bindMCDVESTDAILEGACY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACY{MCDVESTDAILEGACYCaller: MCDVESTDAILEGACYCaller{contract: contract}, MCDVESTDAILEGACYTransactor: MCDVESTDAILEGACYTransactor{contract: contract}, MCDVESTDAILEGACYFilterer: MCDVESTDAILEGACYFilterer{contract: contract}}, nil
}

// NewMCDVESTDAILEGACYCaller creates a new read-only instance of MCDVESTDAILEGACY, bound to a specific deployed contract.
func NewMCDVESTDAILEGACYCaller(address common.Address, caller bind.ContractCaller) (*MCDVESTDAILEGACYCaller, error) {
	contract, err := bindMCDVESTDAILEGACY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYCaller{contract: contract}, nil
}

// NewMCDVESTDAILEGACYTransactor creates a new write-only instance of MCDVESTDAILEGACY, bound to a specific deployed contract.
func NewMCDVESTDAILEGACYTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDVESTDAILEGACYTransactor, error) {
	contract, err := bindMCDVESTDAILEGACY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYTransactor{contract: contract}, nil
}

// NewMCDVESTDAILEGACYFilterer creates a new log filterer instance of MCDVESTDAILEGACY, bound to a specific deployed contract.
func NewMCDVESTDAILEGACYFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDVESTDAILEGACYFilterer, error) {
	contract, err := bindMCDVESTDAILEGACY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYFilterer{contract: contract}, nil
}

// bindMCDVESTDAILEGACY binds a generic wrapper to an already deployed contract.
func bindMCDVESTDAILEGACY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDVESTDAILEGACYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTDAILEGACY.Contract.MCDVESTDAILEGACYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.MCDVESTDAILEGACYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.MCDVESTDAILEGACYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTDAILEGACY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.contract.Transact(opts, method, params...)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) TWENTYYEARS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "TWENTY_YEARS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.TWENTYYEARS(&_MCDVESTDAILEGACY.CallOpts)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.TWENTYYEARS(&_MCDVESTDAILEGACY.CallOpts)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Accrued(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "accrued", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Accrued(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Accrued(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Awards(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "awards", arg0)

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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTDAILEGACY.Contract.Awards(&_MCDVESTDAILEGACY.CallOpts, arg0)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTDAILEGACY.Contract.Awards(&_MCDVESTDAILEGACY.CallOpts, arg0)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Bgn(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "bgn", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Bgn(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Bgn(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Cap() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Cap(&_MCDVESTDAILEGACY.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Cap() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Cap(&_MCDVESTDAILEGACY.CallOpts)
}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Chainlog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "chainlog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Chainlog() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Chainlog(&_MCDVESTDAILEGACY.CallOpts)
}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Chainlog() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Chainlog(&_MCDVESTDAILEGACY.CallOpts)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Clf(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "clf", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Clf(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Clf(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) DaiJoin() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.DaiJoin(&_MCDVESTDAILEGACY.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) DaiJoin() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.DaiJoin(&_MCDVESTDAILEGACY.CallOpts)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Fin(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "fin", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Fin(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Fin(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Ids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "ids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Ids() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Ids(&_MCDVESTDAILEGACY.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Ids() (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Ids(&_MCDVESTDAILEGACY.CallOpts)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Mgr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "mgr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Mgr(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Mgr(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Res(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "res", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Res(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Res(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Rxd(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "rxd", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Rxd(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Rxd(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Tot(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "tot", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Tot(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Tot(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Unpaid(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "unpaid", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Unpaid(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Unpaid(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Usr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "usr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Usr(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Usr(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Valid(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "valid", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTDAILEGACY.Contract.Valid(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTDAILEGACY.Contract.Valid(&_MCDVESTDAILEGACY.CallOpts, _id)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Vat() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Vat(&_MCDVESTDAILEGACY.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Vat() (common.Address, error) {
	return _MCDVESTDAILEGACY.Contract.Vat(&_MCDVESTDAILEGACY.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAILEGACY.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Wards(&_MCDVESTDAILEGACY.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTDAILEGACY.Contract.Wards(&_MCDVESTDAILEGACY.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Create(opts *bind.TransactOpts, _usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "create", _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Create(&_MCDVESTDAILEGACY.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Create(&_MCDVESTDAILEGACY.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Deny(&_MCDVESTDAILEGACY.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Deny(&_MCDVESTDAILEGACY.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.File(&_MCDVESTDAILEGACY.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.File(&_MCDVESTDAILEGACY.TransactOpts, what, data)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Move(opts *bind.TransactOpts, _id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "move", _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Move(&_MCDVESTDAILEGACY.TransactOpts, _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Move(&_MCDVESTDAILEGACY.TransactOpts, _id, _dst)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Rely(&_MCDVESTDAILEGACY.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Rely(&_MCDVESTDAILEGACY.TransactOpts, usr)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Restrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "restrict", _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Restrict(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Restrict(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Unrestrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "unrestrict", _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Unrestrict(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Unrestrict(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Vest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "vest", _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Vest(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Vest(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Vest0(opts *bind.TransactOpts, _id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "vest0", _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Vest0(&_MCDVESTDAILEGACY.TransactOpts, _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Vest0(&_MCDVESTDAILEGACY.TransactOpts, _id, _maxAmt)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Yank(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "yank", _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Yank(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Yank(&_MCDVESTDAILEGACY.TransactOpts, _id)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactor) Yank0(opts *bind.TransactOpts, _id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.contract.Transact(opts, "yank0", _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Yank0(&_MCDVESTDAILEGACY.TransactOpts, _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYTransactorSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAILEGACY.Contract.Yank0(&_MCDVESTDAILEGACY.TransactOpts, _id, _end)
}

// MCDVESTDAILEGACYDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYDenyIterator struct {
	Event *MCDVESTDAILEGACYDeny // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYDeny)
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
		it.Event = new(MCDVESTDAILEGACYDeny)
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
func (it *MCDVESTDAILEGACYDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYDeny represents a Deny event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTDAILEGACYDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYDenyIterator{contract: _MCDVESTDAILEGACY.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYDeny)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseDeny(log types.Log) (*MCDVESTDAILEGACYDeny, error) {
	event := new(MCDVESTDAILEGACYDeny)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYFileIterator struct {
	Event *MCDVESTDAILEGACYFile // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYFile)
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
		it.Event = new(MCDVESTDAILEGACYFile)
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
func (it *MCDVESTDAILEGACYFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYFile represents a File event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDVESTDAILEGACYFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYFileIterator{contract: _MCDVESTDAILEGACY.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYFile)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseFile(log types.Log) (*MCDVESTDAILEGACYFile, error) {
	event := new(MCDVESTDAILEGACYFile)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYInitIterator struct {
	Event *MCDVESTDAILEGACYInit // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYInit)
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
		it.Event = new(MCDVESTDAILEGACYInit)
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
func (it *MCDVESTDAILEGACYInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYInit represents a Init event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYInit struct {
	Id  *big.Int
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterInit(opts *bind.FilterOpts, id []*big.Int, usr []common.Address) (*MCDVESTDAILEGACYInitIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYInitIterator{contract: _MCDVESTDAILEGACY.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYInit, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYInit)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Init", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseInit(log types.Log) (*MCDVESTDAILEGACYInit, error) {
	event := new(MCDVESTDAILEGACYInit)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYMoveIterator struct {
	Event *MCDVESTDAILEGACYMove // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYMove)
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
		it.Event = new(MCDVESTDAILEGACYMove)
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
func (it *MCDVESTDAILEGACYMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYMove represents a Move event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYMove struct {
	Id  *big.Int
	Dst common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterMove(opts *bind.FilterOpts, id []*big.Int, dst []common.Address) (*MCDVESTDAILEGACYMoveIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYMoveIterator{contract: _MCDVESTDAILEGACY.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYMove, id []*big.Int, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYMove)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Move", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseMove(log types.Log) (*MCDVESTDAILEGACYMove, error) {
	event := new(MCDVESTDAILEGACYMove)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYRelyIterator struct {
	Event *MCDVESTDAILEGACYRely // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYRely)
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
		it.Event = new(MCDVESTDAILEGACYRely)
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
func (it *MCDVESTDAILEGACYRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYRely represents a Rely event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTDAILEGACYRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYRelyIterator{contract: _MCDVESTDAILEGACY.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYRely)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseRely(log types.Log) (*MCDVESTDAILEGACYRely, error) {
	event := new(MCDVESTDAILEGACYRely)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYRestrictIterator is returned from FilterRestrict and is used to iterate over the raw logs and unpacked data for Restrict events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYRestrictIterator struct {
	Event *MCDVESTDAILEGACYRestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYRestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYRestrict)
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
		it.Event = new(MCDVESTDAILEGACYRestrict)
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
func (it *MCDVESTDAILEGACYRestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYRestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYRestrict represents a Restrict event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYRestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestrict is a free log retrieval operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterRestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAILEGACYRestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYRestrictIterator{contract: _MCDVESTDAILEGACY.contract, event: "Restrict", logs: logs, sub: sub}, nil
}

// WatchRestrict is a free log subscription operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchRestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYRestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYRestrict)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Restrict", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseRestrict(log types.Log) (*MCDVESTDAILEGACYRestrict, error) {
	event := new(MCDVESTDAILEGACYRestrict)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Restrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYUnrestrictIterator is returned from FilterUnrestrict and is used to iterate over the raw logs and unpacked data for Unrestrict events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYUnrestrictIterator struct {
	Event *MCDVESTDAILEGACYUnrestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYUnrestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYUnrestrict)
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
		it.Event = new(MCDVESTDAILEGACYUnrestrict)
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
func (it *MCDVESTDAILEGACYUnrestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYUnrestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYUnrestrict represents a Unrestrict event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYUnrestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnrestrict is a free log retrieval operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterUnrestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAILEGACYUnrestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYUnrestrictIterator{contract: _MCDVESTDAILEGACY.contract, event: "Unrestrict", logs: logs, sub: sub}, nil
}

// WatchUnrestrict is a free log subscription operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchUnrestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYUnrestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYUnrestrict)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Unrestrict", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseUnrestrict(log types.Log) (*MCDVESTDAILEGACYUnrestrict, error) {
	event := new(MCDVESTDAILEGACYUnrestrict)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Unrestrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYVestIterator is returned from FilterVest and is used to iterate over the raw logs and unpacked data for Vest events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYVestIterator struct {
	Event *MCDVESTDAILEGACYVest // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYVestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYVest)
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
		it.Event = new(MCDVESTDAILEGACYVest)
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
func (it *MCDVESTDAILEGACYVestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYVestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYVest represents a Vest event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYVest struct {
	Id  *big.Int
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVest is a free log retrieval operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterVest(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAILEGACYVestIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYVestIterator{contract: _MCDVESTDAILEGACY.contract, event: "Vest", logs: logs, sub: sub}, nil
}

// WatchVest is a free log subscription operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchVest(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYVest, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYVest)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Vest", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseVest(log types.Log) (*MCDVESTDAILEGACYVest, error) {
	event := new(MCDVESTDAILEGACYVest)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Vest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAILEGACYYankIterator is returned from FilterYank and is used to iterate over the raw logs and unpacked data for Yank events raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYYankIterator struct {
	Event *MCDVESTDAILEGACYYank // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAILEGACYYankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAILEGACYYank)
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
		it.Event = new(MCDVESTDAILEGACYYank)
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
func (it *MCDVESTDAILEGACYYankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAILEGACYYankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAILEGACYYank represents a Yank event raised by the MCDVESTDAILEGACY contract.
type MCDVESTDAILEGACYYank struct {
	Id  *big.Int
	End *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterYank is a free log retrieval operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) FilterYank(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAILEGACYYankIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.FilterLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAILEGACYYankIterator{contract: _MCDVESTDAILEGACY.contract, event: "Yank", logs: logs, sub: sub}, nil
}

// WatchYank is a free log subscription operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) WatchYank(opts *bind.WatchOpts, sink chan<- *MCDVESTDAILEGACYYank, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAILEGACY.contract.WatchLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAILEGACYYank)
				if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Yank", log); err != nil {
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
func (_MCDVESTDAILEGACY *MCDVESTDAILEGACYFilterer) ParseYank(log types.Log) (*MCDVESTDAILEGACYYank, error) {
	event := new(MCDVESTDAILEGACYYank)
	if err := _MCDVESTDAILEGACY.contract.UnpackLog(event, "Yank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
