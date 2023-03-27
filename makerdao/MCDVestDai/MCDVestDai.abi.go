// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDVestDai

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

// MCDVESTDAIABI is the input ABI used to generate the binding from.
const MCDVESTDAIABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlog\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"Move\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Restrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Unrestrict\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Vest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"Yank\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TWENTY_YEARS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"accrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"awards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"bgn\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"clf\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"fin\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mgr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"res\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"tot\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"rxd\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"bgn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainlog\",\"outputs\":[{\"internalType\":\"contractChainlogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"clf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bgn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tau\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_eta\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mgr\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoinLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"fin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"mgr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dst\",\"type\":\"address\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"res\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"restrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"rxd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unpaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"unrestrict\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"usr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmt\",\"type\":\"uint256\"}],\"name\":\"vest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MCDVESTDAI is an auto generated Go binding around an Ethereum contract.
type MCDVESTDAI struct {
	MCDVESTDAICaller     // Read-only binding to the contract
	MCDVESTDAITransactor // Write-only binding to the contract
	MCDVESTDAIFilterer   // Log filterer for contract events
}

// MCDVESTDAICaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDVESTDAICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAITransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDVESTDAITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDVESTDAIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDVESTDAISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDVESTDAISession struct {
	Contract     *MCDVESTDAI       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDVESTDAICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDVESTDAICallerSession struct {
	Contract *MCDVESTDAICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MCDVESTDAITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDVESTDAITransactorSession struct {
	Contract     *MCDVESTDAITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MCDVESTDAIRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDVESTDAIRaw struct {
	Contract *MCDVESTDAI // Generic contract binding to access the raw methods on
}

// MCDVESTDAICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDVESTDAICallerRaw struct {
	Contract *MCDVESTDAICaller // Generic read-only contract binding to access the raw methods on
}

// MCDVESTDAITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDVESTDAITransactorRaw struct {
	Contract *MCDVESTDAITransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDVESTDAI creates a new instance of MCDVESTDAI, bound to a specific deployed contract.
func NewMCDVESTDAI(address common.Address, backend bind.ContractBackend) (*MCDVESTDAI, error) {
	contract, err := bindMCDVESTDAI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAI{MCDVESTDAICaller: MCDVESTDAICaller{contract: contract}, MCDVESTDAITransactor: MCDVESTDAITransactor{contract: contract}, MCDVESTDAIFilterer: MCDVESTDAIFilterer{contract: contract}}, nil
}

// NewMCDVESTDAICaller creates a new read-only instance of MCDVESTDAI, bound to a specific deployed contract.
func NewMCDVESTDAICaller(address common.Address, caller bind.ContractCaller) (*MCDVESTDAICaller, error) {
	contract, err := bindMCDVESTDAI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAICaller{contract: contract}, nil
}

// NewMCDVESTDAITransactor creates a new write-only instance of MCDVESTDAI, bound to a specific deployed contract.
func NewMCDVESTDAITransactor(address common.Address, transactor bind.ContractTransactor) (*MCDVESTDAITransactor, error) {
	contract, err := bindMCDVESTDAI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAITransactor{contract: contract}, nil
}

// NewMCDVESTDAIFilterer creates a new log filterer instance of MCDVESTDAI, bound to a specific deployed contract.
func NewMCDVESTDAIFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDVESTDAIFilterer, error) {
	contract, err := bindMCDVESTDAI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIFilterer{contract: contract}, nil
}

// bindMCDVESTDAI binds a generic wrapper to an already deployed contract.
func bindMCDVESTDAI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDVESTDAIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTDAI *MCDVESTDAIRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTDAI.Contract.MCDVESTDAICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTDAI *MCDVESTDAIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.MCDVESTDAITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTDAI *MCDVESTDAIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.MCDVESTDAITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDVESTDAI *MCDVESTDAICallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDVESTDAI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDVESTDAI *MCDVESTDAITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDVESTDAI *MCDVESTDAITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.contract.Transact(opts, method, params...)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) TWENTYYEARS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "TWENTY_YEARS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTDAI.Contract.TWENTYYEARS(&_MCDVESTDAI.CallOpts)
}

// TWENTYYEARS is a free data retrieval call binding the contract method 0x60fb494b.
//
// Solidity: function TWENTY_YEARS() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) TWENTYYEARS() (*big.Int, error) {
	return _MCDVESTDAI.Contract.TWENTYYEARS(&_MCDVESTDAI.CallOpts)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAICaller) Accrued(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "accrued", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAISession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Accrued(&_MCDVESTDAI.CallOpts, _id)
}

// Accrued is a free data retrieval call binding the contract method 0xf52981f4.
//
// Solidity: function accrued(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Accrued(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Accrued(&_MCDVESTDAI.CallOpts, _id)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTDAI *MCDVESTDAICaller) Awards(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _MCDVESTDAI.contract.Call(opts, &out, "awards", arg0)

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
func (_MCDVESTDAI *MCDVESTDAISession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTDAI.Contract.Awards(&_MCDVESTDAI.CallOpts, arg0)
}

// Awards is a free data retrieval call binding the contract method 0xfc5a5b63.
//
// Solidity: function awards(uint256 ) view returns(address usr, uint48 bgn, uint48 clf, uint48 fin, address mgr, uint8 res, uint128 tot, uint128 rxd)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Awards(arg0 *big.Int) (struct {
	Usr common.Address
	Bgn *big.Int
	Clf *big.Int
	Fin *big.Int
	Mgr common.Address
	Res uint8
	Tot *big.Int
	Rxd *big.Int
}, error) {
	return _MCDVESTDAI.Contract.Awards(&_MCDVESTDAI.CallOpts, arg0)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Bgn(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "bgn", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Bgn(&_MCDVESTDAI.CallOpts, _id)
}

// Bgn is a free data retrieval call binding the contract method 0x21f6c0cf.
//
// Solidity: function bgn(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Bgn(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Bgn(&_MCDVESTDAI.CallOpts, _id)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Cap() (*big.Int, error) {
	return _MCDVESTDAI.Contract.Cap(&_MCDVESTDAI.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Cap() (*big.Int, error) {
	return _MCDVESTDAI.Contract.Cap(&_MCDVESTDAI.CallOpts)
}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICaller) Chainlog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "chainlog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAI *MCDVESTDAISession) Chainlog() (common.Address, error) {
	return _MCDVESTDAI.Contract.Chainlog(&_MCDVESTDAI.CallOpts)
}

// Chainlog is a free data retrieval call binding the contract method 0xce6f74aa.
//
// Solidity: function chainlog() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Chainlog() (common.Address, error) {
	return _MCDVESTDAI.Contract.Chainlog(&_MCDVESTDAI.CallOpts)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Clf(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "clf", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Clf(&_MCDVESTDAI.CallOpts, _id)
}

// Clf is a free data retrieval call binding the contract method 0xcdf43497.
//
// Solidity: function clf(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Clf(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Clf(&_MCDVESTDAI.CallOpts, _id)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAI *MCDVESTDAISession) DaiJoin() (common.Address, error) {
	return _MCDVESTDAI.Contract.DaiJoin(&_MCDVESTDAI.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICallerSession) DaiJoin() (common.Address, error) {
	return _MCDVESTDAI.Contract.DaiJoin(&_MCDVESTDAI.CallOpts)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Fin(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "fin", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Fin(&_MCDVESTDAI.CallOpts, _id)
}

// Fin is a free data retrieval call binding the contract method 0xe529780d.
//
// Solidity: function fin(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Fin(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Fin(&_MCDVESTDAI.CallOpts, _id)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Ids(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "ids")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Ids() (*big.Int, error) {
	return _MCDVESTDAI.Contract.Ids(&_MCDVESTDAI.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xe7657e15.
//
// Solidity: function ids() view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Ids() (*big.Int, error) {
	return _MCDVESTDAI.Contract.Ids(&_MCDVESTDAI.CallOpts)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAICaller) Mgr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "mgr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAISession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAI.Contract.Mgr(&_MCDVESTDAI.CallOpts, _id)
}

// Mgr is a free data retrieval call binding the contract method 0xdc2c788f.
//
// Solidity: function mgr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Mgr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAI.Contract.Mgr(&_MCDVESTDAI.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Res(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "res", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Res(&_MCDVESTDAI.CallOpts, _id)
}

// Res is a free data retrieval call binding the contract method 0xd4e8fd2e.
//
// Solidity: function res(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Res(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Res(&_MCDVESTDAI.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Rxd(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "rxd", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Rxd(&_MCDVESTDAI.CallOpts, _id)
}

// Rxd is a free data retrieval call binding the contract method 0xe054720f.
//
// Solidity: function rxd(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Rxd(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Rxd(&_MCDVESTDAI.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Tot(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "tot", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Tot(&_MCDVESTDAI.CallOpts, _id)
}

// Tot is a free data retrieval call binding the contract method 0x892de51d.
//
// Solidity: function tot(uint256 _id) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Tot(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Tot(&_MCDVESTDAI.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAICaller) Unpaid(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "unpaid", _id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAISession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Unpaid(&_MCDVESTDAI.CallOpts, _id)
}

// Unpaid is a free data retrieval call binding the contract method 0x53e8863d.
//
// Solidity: function unpaid(uint256 _id) view returns(uint256 amt)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Unpaid(_id *big.Int) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Unpaid(&_MCDVESTDAI.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAICaller) Usr(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "usr", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAISession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAI.Contract.Usr(&_MCDVESTDAI.CallOpts, _id)
}

// Usr is a free data retrieval call binding the contract method 0xc659cd45.
//
// Solidity: function usr(uint256 _id) view returns(address)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Usr(_id *big.Int) (common.Address, error) {
	return _MCDVESTDAI.Contract.Usr(&_MCDVESTDAI.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAI *MCDVESTDAICaller) Valid(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "valid", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAI *MCDVESTDAISession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTDAI.Contract.Valid(&_MCDVESTDAI.CallOpts, _id)
}

// Valid is a free data retrieval call binding the contract method 0xbf8712c5.
//
// Solidity: function valid(uint256 _id) view returns(bool isValid)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Valid(_id *big.Int) (bool, error) {
	return _MCDVESTDAI.Contract.Valid(&_MCDVESTDAI.CallOpts, _id)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAI *MCDVESTDAISession) Vat() (common.Address, error) {
	return _MCDVESTDAI.Contract.Vat(&_MCDVESTDAI.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Vat() (common.Address, error) {
	return _MCDVESTDAI.Contract.Vat(&_MCDVESTDAI.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDVESTDAI.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAISession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Wards(&_MCDVESTDAI.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_MCDVESTDAI *MCDVESTDAICallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _MCDVESTDAI.Contract.Wards(&_MCDVESTDAI.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAI *MCDVESTDAITransactor) Create(opts *bind.TransactOpts, _usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "create", _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAI *MCDVESTDAISession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Create(&_MCDVESTDAI.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Create is a paid mutator transaction binding the contract method 0xdb64ff8f.
//
// Solidity: function create(address _usr, uint256 _tot, uint256 _bgn, uint256 _tau, uint256 _eta, address _mgr) returns(uint256 id)
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Create(_usr common.Address, _tot *big.Int, _bgn *big.Int, _tau *big.Int, _eta *big.Int, _mgr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Create(&_MCDVESTDAI.TransactOpts, _usr, _tot, _bgn, _tau, _eta, _mgr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Deny(opts *bind.TransactOpts, _usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "deny", _usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Deny(_usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Deny(&_MCDVESTDAI.TransactOpts, _usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Deny(_usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Deny(&_MCDVESTDAI.TransactOpts, _usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAI *MCDVESTDAISession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.File(&_MCDVESTDAI.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.File(&_MCDVESTDAI.TransactOpts, what, data)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Move(opts *bind.TransactOpts, _id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "move", _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Move(&_MCDVESTDAI.TransactOpts, _id, _dst)
}

// Move is a paid mutator transaction binding the contract method 0xd8a8e03a.
//
// Solidity: function move(uint256 _id, address _dst) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Move(_id *big.Int, _dst common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Move(&_MCDVESTDAI.TransactOpts, _id, _dst)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Rely(opts *bind.TransactOpts, _usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "rely", _usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Rely(_usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Rely(&_MCDVESTDAI.TransactOpts, _usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address _usr) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Rely(_usr common.Address) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Rely(&_MCDVESTDAI.TransactOpts, _usr)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Restrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "restrict", _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Restrict(&_MCDVESTDAI.TransactOpts, _id)
}

// Restrict is a paid mutator transaction binding the contract method 0x3c433d5f.
//
// Solidity: function restrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Restrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Restrict(&_MCDVESTDAI.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Unrestrict(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "unrestrict", _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Unrestrict(&_MCDVESTDAI.TransactOpts, _id)
}

// Unrestrict is a paid mutator transaction binding the contract method 0x7d8d2702.
//
// Solidity: function unrestrict(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Unrestrict(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Unrestrict(&_MCDVESTDAI.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Vest(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "vest", _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Vest(&_MCDVESTDAI.TransactOpts, _id)
}

// Vest is a paid mutator transaction binding the contract method 0x6a760b80.
//
// Solidity: function vest(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Vest(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Vest(&_MCDVESTDAI.TransactOpts, _id)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Vest0(opts *bind.TransactOpts, _id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "vest0", _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Vest0(&_MCDVESTDAI.TransactOpts, _id, _maxAmt)
}

// Vest0 is a paid mutator transaction binding the contract method 0xbb7c46f3.
//
// Solidity: function vest(uint256 _id, uint256 _maxAmt) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Vest0(_id *big.Int, _maxAmt *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Vest0(&_MCDVESTDAI.TransactOpts, _id, _maxAmt)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Yank(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "yank", _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Yank(&_MCDVESTDAI.TransactOpts, _id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 _id) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Yank(_id *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Yank(&_MCDVESTDAI.TransactOpts, _id)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAI *MCDVESTDAITransactor) Yank0(opts *bind.TransactOpts, _id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.contract.Transact(opts, "yank0", _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAI *MCDVESTDAISession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Yank0(&_MCDVESTDAI.TransactOpts, _id, _end)
}

// Yank0 is a paid mutator transaction binding the contract method 0x509aaa1d.
//
// Solidity: function yank(uint256 _id, uint256 _end) returns()
func (_MCDVESTDAI *MCDVESTDAITransactorSession) Yank0(_id *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _MCDVESTDAI.Contract.Yank0(&_MCDVESTDAI.TransactOpts, _id, _end)
}

// MCDVESTDAIDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the MCDVESTDAI contract.
type MCDVESTDAIDenyIterator struct {
	Event *MCDVESTDAIDeny // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIDeny)
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
		it.Event = new(MCDVESTDAIDeny)
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
func (it *MCDVESTDAIDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIDeny represents a Deny event raised by the MCDVESTDAI contract.
type MCDVESTDAIDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTDAIDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIDenyIterator{contract: _MCDVESTDAI.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIDeny)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Deny", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseDeny(log types.Log) (*MCDVESTDAIDeny, error) {
	event := new(MCDVESTDAIDeny)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the MCDVESTDAI contract.
type MCDVESTDAIFileIterator struct {
	Event *MCDVESTDAIFile // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIFile)
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
		it.Event = new(MCDVESTDAIFile)
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
func (it *MCDVESTDAIFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIFile represents a File event raised by the MCDVESTDAI contract.
type MCDVESTDAIFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*MCDVESTDAIFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIFileIterator{contract: _MCDVESTDAI.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIFile)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "File", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseFile(log types.Log) (*MCDVESTDAIFile, error) {
	event := new(MCDVESTDAIFile)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIInitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MCDVESTDAI contract.
type MCDVESTDAIInitIterator struct {
	Event *MCDVESTDAIInit // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIInit)
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
		it.Event = new(MCDVESTDAIInit)
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
func (it *MCDVESTDAIInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIInit represents a Init event raised by the MCDVESTDAI contract.
type MCDVESTDAIInit struct {
	Id  *big.Int
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterInit(opts *bind.FilterOpts, id []*big.Int, usr []common.Address) (*MCDVESTDAIInitIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIInitIterator{contract: _MCDVESTDAI.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x2e3cc5298d3204a0f0fc2be0f6fdefcef002025f4c75caf950b23e6cfbfb78d0.
//
// Solidity: event Init(uint256 indexed id, address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIInit, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Init", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIInit)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Init", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseInit(log types.Log) (*MCDVESTDAIInit, error) {
	event := new(MCDVESTDAIInit)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIMoveIterator is returned from FilterMove and is used to iterate over the raw logs and unpacked data for Move events raised by the MCDVESTDAI contract.
type MCDVESTDAIMoveIterator struct {
	Event *MCDVESTDAIMove // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIMove)
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
		it.Event = new(MCDVESTDAIMove)
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
func (it *MCDVESTDAIMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIMove represents a Move event raised by the MCDVESTDAI contract.
type MCDVESTDAIMove struct {
	Id  *big.Int
	Dst common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMove is a free log retrieval operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterMove(opts *bind.FilterOpts, id []*big.Int, dst []common.Address) (*MCDVESTDAIMoveIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIMoveIterator{contract: _MCDVESTDAI.contract, event: "Move", logs: logs, sub: sub}, nil
}

// WatchMove is a free log subscription operation binding the contract event 0x8ceddd02f4fb8ef0d5d6212cf4c91d59d366e04b977e8b2b944168d2a6d85081.
//
// Solidity: event Move(uint256 indexed id, address indexed dst)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchMove(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIMove, id []*big.Int, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Move", idRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIMove)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Move", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseMove(log types.Log) (*MCDVESTDAIMove, error) {
	event := new(MCDVESTDAIMove)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Move", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the MCDVESTDAI contract.
type MCDVESTDAIRelyIterator struct {
	Event *MCDVESTDAIRely // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIRely)
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
		it.Event = new(MCDVESTDAIRely)
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
func (it *MCDVESTDAIRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIRely represents a Rely event raised by the MCDVESTDAI contract.
type MCDVESTDAIRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*MCDVESTDAIRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIRelyIterator{contract: _MCDVESTDAI.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIRely)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Rely", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseRely(log types.Log) (*MCDVESTDAIRely, error) {
	event := new(MCDVESTDAIRely)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIRestrictIterator is returned from FilterRestrict and is used to iterate over the raw logs and unpacked data for Restrict events raised by the MCDVESTDAI contract.
type MCDVESTDAIRestrictIterator struct {
	Event *MCDVESTDAIRestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIRestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIRestrict)
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
		it.Event = new(MCDVESTDAIRestrict)
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
func (it *MCDVESTDAIRestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIRestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIRestrict represents a Restrict event raised by the MCDVESTDAI contract.
type MCDVESTDAIRestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRestrict is a free log retrieval operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterRestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAIRestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIRestrictIterator{contract: _MCDVESTDAI.contract, event: "Restrict", logs: logs, sub: sub}, nil
}

// WatchRestrict is a free log subscription operation binding the contract event 0x9247a2bf1b75bc397d4043d99b9cebce531548a01dbb56a5d4c5f5ca26051e8d.
//
// Solidity: event Restrict(uint256 indexed id)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchRestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIRestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Restrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIRestrict)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Restrict", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseRestrict(log types.Log) (*MCDVESTDAIRestrict, error) {
	event := new(MCDVESTDAIRestrict)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Restrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIUnrestrictIterator is returned from FilterUnrestrict and is used to iterate over the raw logs and unpacked data for Unrestrict events raised by the MCDVESTDAI contract.
type MCDVESTDAIUnrestrictIterator struct {
	Event *MCDVESTDAIUnrestrict // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIUnrestrictIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIUnrestrict)
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
		it.Event = new(MCDVESTDAIUnrestrict)
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
func (it *MCDVESTDAIUnrestrictIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIUnrestrictIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIUnrestrict represents a Unrestrict event raised by the MCDVESTDAI contract.
type MCDVESTDAIUnrestrict struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnrestrict is a free log retrieval operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterUnrestrict(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAIUnrestrictIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIUnrestrictIterator{contract: _MCDVESTDAI.contract, event: "Unrestrict", logs: logs, sub: sub}, nil
}

// WatchUnrestrict is a free log subscription operation binding the contract event 0x3d1b575f06b2d660af77eec35d9b3ffcfa956b6c1fdbc840992d4b03b03e622b.
//
// Solidity: event Unrestrict(uint256 indexed id)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchUnrestrict(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIUnrestrict, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Unrestrict", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIUnrestrict)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Unrestrict", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseUnrestrict(log types.Log) (*MCDVESTDAIUnrestrict, error) {
	event := new(MCDVESTDAIUnrestrict)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Unrestrict", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIVestIterator is returned from FilterVest and is used to iterate over the raw logs and unpacked data for Vest events raised by the MCDVESTDAI contract.
type MCDVESTDAIVestIterator struct {
	Event *MCDVESTDAIVest // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIVestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIVest)
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
		it.Event = new(MCDVESTDAIVest)
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
func (it *MCDVESTDAIVestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIVestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIVest represents a Vest event raised by the MCDVESTDAI contract.
type MCDVESTDAIVest struct {
	Id  *big.Int
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVest is a free log retrieval operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterVest(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAIVestIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIVestIterator{contract: _MCDVESTDAI.contract, event: "Vest", logs: logs, sub: sub}, nil
}

// WatchVest is a free log subscription operation binding the contract event 0xa2906882572b0e9dfe893158bb064bc308eb1bd87d1da481850f9d17fc293847.
//
// Solidity: event Vest(uint256 indexed id, uint256 amt)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchVest(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIVest, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Vest", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIVest)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Vest", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseVest(log types.Log) (*MCDVESTDAIVest, error) {
	event := new(MCDVESTDAIVest)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Vest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDVESTDAIYankIterator is returned from FilterYank and is used to iterate over the raw logs and unpacked data for Yank events raised by the MCDVESTDAI contract.
type MCDVESTDAIYankIterator struct {
	Event *MCDVESTDAIYank // Event containing the contract specifics and raw log

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
func (it *MCDVESTDAIYankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDVESTDAIYank)
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
		it.Event = new(MCDVESTDAIYank)
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
func (it *MCDVESTDAIYankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDVESTDAIYankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDVESTDAIYank represents a Yank event raised by the MCDVESTDAI contract.
type MCDVESTDAIYank struct {
	Id  *big.Int
	End *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterYank is a free log retrieval operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTDAI *MCDVESTDAIFilterer) FilterYank(opts *bind.FilterOpts, id []*big.Int) (*MCDVESTDAIYankIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.FilterLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return &MCDVESTDAIYankIterator{contract: _MCDVESTDAI.contract, event: "Yank", logs: logs, sub: sub}, nil
}

// WatchYank is a free log subscription operation binding the contract event 0x6f2a3ed78a3066d89360b6c89e52bf3313f52e859401a3ea5fa0f033fd540c3c.
//
// Solidity: event Yank(uint256 indexed id, uint256 end)
func (_MCDVESTDAI *MCDVESTDAIFilterer) WatchYank(opts *bind.WatchOpts, sink chan<- *MCDVESTDAIYank, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MCDVESTDAI.contract.WatchLogs(opts, "Yank", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDVESTDAIYank)
				if err := _MCDVESTDAI.contract.UnpackLog(event, "Yank", log); err != nil {
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
func (_MCDVESTDAI *MCDVESTDAIFilterer) ParseYank(log types.Log) (*MCDVESTDAIYank, error) {
	event := new(MCDVESTDAIYank)
	if err := _MCDVESTDAI.contract.UnpackLog(event, "Yank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
