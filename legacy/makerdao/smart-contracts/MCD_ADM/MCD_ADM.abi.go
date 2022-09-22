// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCD_ADM

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

// MCDADMABI is the input ABI used to generate the binding from.
const MCDADMABI = "[{\"inputs\":[{\"internalType\":\"contractDSToken\",\"name\":\"GOV\",\"type\":\"address\"},{\"internalType\":\"contractDSToken\",\"name\":\"IOU\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"MAX_YAYS\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"slate\",\"type\":\"bytes32\"}],\"name\":\"Etch\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"GOV\",\"outputs\":[{\"internalType\":\"contractDSToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"IOU\",\"outputs\":[{\"internalType\":\"contractDSToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_YAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"yays\",\"type\":\"address[]\"}],\"name\":\"etch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"slate\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"free\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"getCapabilityRoles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"getUserRoles\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"hasUserRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"isCapabilityPublic\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isUserRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"launch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"lift\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setPublicCapability\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"code\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setRoleCapability\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setRootUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setUserRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slate\",\"type\":\"bytes32\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"yays\",\"type\":\"address[]\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDADM is an auto generated Go binding around an Ethereum contract.
type MCDADM struct {
	MCDADMCaller     // Read-only binding to the contract
	MCDADMTransactor // Write-only binding to the contract
	MCDADMFilterer   // Log filterer for contract events
}

// MCDADMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDADMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDADMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDADMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDADMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDADMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDADMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDADMSession struct {
	Contract     *MCDADM           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDADMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDADMCallerSession struct {
	Contract *MCDADMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MCDADMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDADMTransactorSession struct {
	Contract     *MCDADMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDADMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDADMRaw struct {
	Contract *MCDADM // Generic contract binding to access the raw methods on
}

// MCDADMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDADMCallerRaw struct {
	Contract *MCDADMCaller // Generic read-only contract binding to access the raw methods on
}

// MCDADMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDADMTransactorRaw struct {
	Contract *MCDADMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDADM creates a new instance of MCDADM, bound to a specific deployed contract.
func NewMCDADM(address common.Address, backend bind.ContractBackend) (*MCDADM, error) {
	contract, err := bindMCDADM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDADM{MCDADMCaller: MCDADMCaller{contract: contract}, MCDADMTransactor: MCDADMTransactor{contract: contract}, MCDADMFilterer: MCDADMFilterer{contract: contract}}, nil
}

// NewMCDADMCaller creates a new read-only instance of MCDADM, bound to a specific deployed contract.
func NewMCDADMCaller(address common.Address, caller bind.ContractCaller) (*MCDADMCaller, error) {
	contract, err := bindMCDADM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDADMCaller{contract: contract}, nil
}

// NewMCDADMTransactor creates a new write-only instance of MCDADM, bound to a specific deployed contract.
func NewMCDADMTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDADMTransactor, error) {
	contract, err := bindMCDADM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDADMTransactor{contract: contract}, nil
}

// NewMCDADMFilterer creates a new log filterer instance of MCDADM, bound to a specific deployed contract.
func NewMCDADMFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDADMFilterer, error) {
	contract, err := bindMCDADM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDADMFilterer{contract: contract}, nil
}

// bindMCDADM binds a generic wrapper to an already deployed contract.
func bindMCDADM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDADMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDADM *MCDADMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDADM.Contract.MCDADMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDADM *MCDADMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDADM.Contract.MCDADMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDADM *MCDADMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDADM.Contract.MCDADMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDADM *MCDADMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDADM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDADM *MCDADMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDADM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDADM *MCDADMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDADM.Contract.contract.Transact(opts, method, params...)
}

// GOV is a free data retrieval call binding the contract method 0x180cb47f.
//
// Solidity: function GOV() view returns(address)
func (_MCDADM *MCDADMCaller) GOV(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "GOV")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOV is a free data retrieval call binding the contract method 0x180cb47f.
//
// Solidity: function GOV() view returns(address)
func (_MCDADM *MCDADMSession) GOV() (common.Address, error) {
	return _MCDADM.Contract.GOV(&_MCDADM.CallOpts)
}

// GOV is a free data retrieval call binding the contract method 0x180cb47f.
//
// Solidity: function GOV() view returns(address)
func (_MCDADM *MCDADMCallerSession) GOV() (common.Address, error) {
	return _MCDADM.Contract.GOV(&_MCDADM.CallOpts)
}

// IOU is a free data retrieval call binding the contract method 0x046c472f.
//
// Solidity: function IOU() view returns(address)
func (_MCDADM *MCDADMCaller) IOU(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "IOU")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IOU is a free data retrieval call binding the contract method 0x046c472f.
//
// Solidity: function IOU() view returns(address)
func (_MCDADM *MCDADMSession) IOU() (common.Address, error) {
	return _MCDADM.Contract.IOU(&_MCDADM.CallOpts)
}

// IOU is a free data retrieval call binding the contract method 0x046c472f.
//
// Solidity: function IOU() view returns(address)
func (_MCDADM *MCDADMCallerSession) IOU() (common.Address, error) {
	return _MCDADM.Contract.IOU(&_MCDADM.CallOpts)
}

// MAXYAYS is a free data retrieval call binding the contract method 0x362344b8.
//
// Solidity: function MAX_YAYS() view returns(uint256)
func (_MCDADM *MCDADMCaller) MAXYAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "MAX_YAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXYAYS is a free data retrieval call binding the contract method 0x362344b8.
//
// Solidity: function MAX_YAYS() view returns(uint256)
func (_MCDADM *MCDADMSession) MAXYAYS() (*big.Int, error) {
	return _MCDADM.Contract.MAXYAYS(&_MCDADM.CallOpts)
}

// MAXYAYS is a free data retrieval call binding the contract method 0x362344b8.
//
// Solidity: function MAX_YAYS() view returns(uint256)
func (_MCDADM *MCDADMCallerSession) MAXYAYS() (*big.Int, error) {
	return _MCDADM.Contract.MAXYAYS(&_MCDADM.CallOpts)
}

// Approvals is a free data retrieval call binding the contract method 0x5d0341ba.
//
// Solidity: function approvals(address ) view returns(uint256)
func (_MCDADM *MCDADMCaller) Approvals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "approvals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Approvals is a free data retrieval call binding the contract method 0x5d0341ba.
//
// Solidity: function approvals(address ) view returns(uint256)
func (_MCDADM *MCDADMSession) Approvals(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Approvals(&_MCDADM.CallOpts, arg0)
}

// Approvals is a free data retrieval call binding the contract method 0x5d0341ba.
//
// Solidity: function approvals(address ) view returns(uint256)
func (_MCDADM *MCDADMCallerSession) Approvals(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Approvals(&_MCDADM.CallOpts, arg0)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDADM *MCDADMCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDADM *MCDADMSession) Authority() (common.Address, error) {
	return _MCDADM.Contract.Authority(&_MCDADM.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDADM *MCDADMCallerSession) Authority() (common.Address, error) {
	return _MCDADM.Contract.Authority(&_MCDADM.CallOpts)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address caller, address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMCaller) CanCall(opts *bind.CallOpts, caller common.Address, code common.Address, sig [4]byte) (bool, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "canCall", caller, code, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address caller, address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMSession) CanCall(caller common.Address, code common.Address, sig [4]byte) (bool, error) {
	return _MCDADM.Contract.CanCall(&_MCDADM.CallOpts, caller, code, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(address caller, address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMCallerSession) CanCall(caller common.Address, code common.Address, sig [4]byte) (bool, error) {
	return _MCDADM.Contract.CanCall(&_MCDADM.CallOpts, caller, code, sig)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_MCDADM *MCDADMCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_MCDADM *MCDADMSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Deposits(&_MCDADM.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256)
func (_MCDADM *MCDADMCallerSession) Deposits(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Deposits(&_MCDADM.CallOpts, arg0)
}

// GetCapabilityRoles is a free data retrieval call binding the contract method 0x27538e90.
//
// Solidity: function getCapabilityRoles(address code, bytes4 sig) view returns(bytes32)
func (_MCDADM *MCDADMCaller) GetCapabilityRoles(opts *bind.CallOpts, code common.Address, sig [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "getCapabilityRoles", code, sig)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCapabilityRoles is a free data retrieval call binding the contract method 0x27538e90.
//
// Solidity: function getCapabilityRoles(address code, bytes4 sig) view returns(bytes32)
func (_MCDADM *MCDADMSession) GetCapabilityRoles(code common.Address, sig [4]byte) ([32]byte, error) {
	return _MCDADM.Contract.GetCapabilityRoles(&_MCDADM.CallOpts, code, sig)
}

// GetCapabilityRoles is a free data retrieval call binding the contract method 0x27538e90.
//
// Solidity: function getCapabilityRoles(address code, bytes4 sig) view returns(bytes32)
func (_MCDADM *MCDADMCallerSession) GetCapabilityRoles(code common.Address, sig [4]byte) ([32]byte, error) {
	return _MCDADM.Contract.GetCapabilityRoles(&_MCDADM.CallOpts, code, sig)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(bytes32)
func (_MCDADM *MCDADMCaller) GetUserRoles(opts *bind.CallOpts, who common.Address) ([32]byte, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "getUserRoles", who)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(bytes32)
func (_MCDADM *MCDADMSession) GetUserRoles(who common.Address) ([32]byte, error) {
	return _MCDADM.Contract.GetUserRoles(&_MCDADM.CallOpts, who)
}

// GetUserRoles is a free data retrieval call binding the contract method 0x06a36aee.
//
// Solidity: function getUserRoles(address who) view returns(bytes32)
func (_MCDADM *MCDADMCallerSession) GetUserRoles(who common.Address) ([32]byte, error) {
	return _MCDADM.Contract.GetUserRoles(&_MCDADM.CallOpts, who)
}

// HasUserRole is a free data retrieval call binding the contract method 0xa078f737.
//
// Solidity: function hasUserRole(address who, uint8 role) view returns(bool)
func (_MCDADM *MCDADMCaller) HasUserRole(opts *bind.CallOpts, who common.Address, role uint8) (bool, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "hasUserRole", who, role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserRole is a free data retrieval call binding the contract method 0xa078f737.
//
// Solidity: function hasUserRole(address who, uint8 role) view returns(bool)
func (_MCDADM *MCDADMSession) HasUserRole(who common.Address, role uint8) (bool, error) {
	return _MCDADM.Contract.HasUserRole(&_MCDADM.CallOpts, who, role)
}

// HasUserRole is a free data retrieval call binding the contract method 0xa078f737.
//
// Solidity: function hasUserRole(address who, uint8 role) view returns(bool)
func (_MCDADM *MCDADMCallerSession) HasUserRole(who common.Address, role uint8) (bool, error) {
	return _MCDADM.Contract.HasUserRole(&_MCDADM.CallOpts, who, role)
}

// Hat is a free data retrieval call binding the contract method 0xfe95a5ce.
//
// Solidity: function hat() view returns(address)
func (_MCDADM *MCDADMCaller) Hat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "hat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hat is a free data retrieval call binding the contract method 0xfe95a5ce.
//
// Solidity: function hat() view returns(address)
func (_MCDADM *MCDADMSession) Hat() (common.Address, error) {
	return _MCDADM.Contract.Hat(&_MCDADM.CallOpts)
}

// Hat is a free data retrieval call binding the contract method 0xfe95a5ce.
//
// Solidity: function hat() view returns(address)
func (_MCDADM *MCDADMCallerSession) Hat() (common.Address, error) {
	return _MCDADM.Contract.Hat(&_MCDADM.CallOpts)
}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMCaller) IsCapabilityPublic(opts *bind.CallOpts, code common.Address, sig [4]byte) (bool, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "isCapabilityPublic", code, sig)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMSession) IsCapabilityPublic(code common.Address, sig [4]byte) (bool, error) {
	return _MCDADM.Contract.IsCapabilityPublic(&_MCDADM.CallOpts, code, sig)
}

// IsCapabilityPublic is a free data retrieval call binding the contract method 0x2f47571f.
//
// Solidity: function isCapabilityPublic(address code, bytes4 sig) view returns(bool)
func (_MCDADM *MCDADMCallerSession) IsCapabilityPublic(code common.Address, sig [4]byte) (bool, error) {
	return _MCDADM.Contract.IsCapabilityPublic(&_MCDADM.CallOpts, code, sig)
}

// IsUserRoot is a free data retrieval call binding the contract method 0xfbf80773.
//
// Solidity: function isUserRoot(address who) view returns(bool)
func (_MCDADM *MCDADMCaller) IsUserRoot(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "isUserRoot", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserRoot is a free data retrieval call binding the contract method 0xfbf80773.
//
// Solidity: function isUserRoot(address who) view returns(bool)
func (_MCDADM *MCDADMSession) IsUserRoot(who common.Address) (bool, error) {
	return _MCDADM.Contract.IsUserRoot(&_MCDADM.CallOpts, who)
}

// IsUserRoot is a free data retrieval call binding the contract method 0xfbf80773.
//
// Solidity: function isUserRoot(address who) view returns(bool)
func (_MCDADM *MCDADMCallerSession) IsUserRoot(who common.Address) (bool, error) {
	return _MCDADM.Contract.IsUserRoot(&_MCDADM.CallOpts, who)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_MCDADM *MCDADMCaller) Last(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "last", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_MCDADM *MCDADMSession) Last(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Last(&_MCDADM.CallOpts, arg0)
}

// Last is a free data retrieval call binding the contract method 0x9a816f7d.
//
// Solidity: function last(address ) view returns(uint256)
func (_MCDADM *MCDADMCallerSession) Last(arg0 common.Address) (*big.Int, error) {
	return _MCDADM.Contract.Last(&_MCDADM.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_MCDADM *MCDADMCaller) Live(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_MCDADM *MCDADMSession) Live() (bool, error) {
	return _MCDADM.Contract.Live(&_MCDADM.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(bool)
func (_MCDADM *MCDADMCallerSession) Live() (bool, error) {
	return _MCDADM.Contract.Live(&_MCDADM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDADM *MCDADMCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDADM *MCDADMSession) Owner() (common.Address, error) {
	return _MCDADM.Contract.Owner(&_MCDADM.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDADM *MCDADMCallerSession) Owner() (common.Address, error) {
	return _MCDADM.Contract.Owner(&_MCDADM.CallOpts)
}

// Slates is a free data retrieval call binding the contract method 0xc2ffc7bb.
//
// Solidity: function slates(bytes32 , uint256 ) view returns(address)
func (_MCDADM *MCDADMCaller) Slates(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "slates", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slates is a free data retrieval call binding the contract method 0xc2ffc7bb.
//
// Solidity: function slates(bytes32 , uint256 ) view returns(address)
func (_MCDADM *MCDADMSession) Slates(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _MCDADM.Contract.Slates(&_MCDADM.CallOpts, arg0, arg1)
}

// Slates is a free data retrieval call binding the contract method 0xc2ffc7bb.
//
// Solidity: function slates(bytes32 , uint256 ) view returns(address)
func (_MCDADM *MCDADMCallerSession) Slates(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _MCDADM.Contract.Slates(&_MCDADM.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bytes32)
func (_MCDADM *MCDADMCaller) Votes(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _MCDADM.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bytes32)
func (_MCDADM *MCDADMSession) Votes(arg0 common.Address) ([32]byte, error) {
	return _MCDADM.Contract.Votes(&_MCDADM.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xd8bff5a5.
//
// Solidity: function votes(address ) view returns(bytes32)
func (_MCDADM *MCDADMCallerSession) Votes(arg0 common.Address) ([32]byte, error) {
	return _MCDADM.Contract.Votes(&_MCDADM.CallOpts, arg0)
}

// Etch is a paid mutator transaction binding the contract method 0x5123e1fa.
//
// Solidity: function etch(address[] yays) returns(bytes32 slate)
func (_MCDADM *MCDADMTransactor) Etch(opts *bind.TransactOpts, yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "etch", yays)
}

// Etch is a paid mutator transaction binding the contract method 0x5123e1fa.
//
// Solidity: function etch(address[] yays) returns(bytes32 slate)
func (_MCDADM *MCDADMSession) Etch(yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Etch(&_MCDADM.TransactOpts, yays)
}

// Etch is a paid mutator transaction binding the contract method 0x5123e1fa.
//
// Solidity: function etch(address[] yays) returns(bytes32 slate)
func (_MCDADM *MCDADMTransactorSession) Etch(yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Etch(&_MCDADM.TransactOpts, yays)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_MCDADM *MCDADMTransactor) Free(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "free", wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_MCDADM *MCDADMSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.Contract.Free(&_MCDADM.TransactOpts, wad)
}

// Free is a paid mutator transaction binding the contract method 0xd8ccd0f3.
//
// Solidity: function free(uint256 wad) returns()
func (_MCDADM *MCDADMTransactorSession) Free(wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.Contract.Free(&_MCDADM.TransactOpts, wad)
}

// Launch is a paid mutator transaction binding the contract method 0x01339c21.
//
// Solidity: function launch() returns()
func (_MCDADM *MCDADMTransactor) Launch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "launch")
}

// Launch is a paid mutator transaction binding the contract method 0x01339c21.
//
// Solidity: function launch() returns()
func (_MCDADM *MCDADMSession) Launch() (*types.Transaction, error) {
	return _MCDADM.Contract.Launch(&_MCDADM.TransactOpts)
}

// Launch is a paid mutator transaction binding the contract method 0x01339c21.
//
// Solidity: function launch() returns()
func (_MCDADM *MCDADMTransactorSession) Launch() (*types.Transaction, error) {
	return _MCDADM.Contract.Launch(&_MCDADM.TransactOpts)
}

// Lift is a paid mutator transaction binding the contract method 0x3c278bd5.
//
// Solidity: function lift(address whom) returns()
func (_MCDADM *MCDADMTransactor) Lift(opts *bind.TransactOpts, whom common.Address) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "lift", whom)
}

// Lift is a paid mutator transaction binding the contract method 0x3c278bd5.
//
// Solidity: function lift(address whom) returns()
func (_MCDADM *MCDADMSession) Lift(whom common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Lift(&_MCDADM.TransactOpts, whom)
}

// Lift is a paid mutator transaction binding the contract method 0x3c278bd5.
//
// Solidity: function lift(address whom) returns()
func (_MCDADM *MCDADMTransactorSession) Lift(whom common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Lift(&_MCDADM.TransactOpts, whom)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_MCDADM *MCDADMTransactor) Lock(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "lock", wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_MCDADM *MCDADMSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.Contract.Lock(&_MCDADM.TransactOpts, wad)
}

// Lock is a paid mutator transaction binding the contract method 0xdd467064.
//
// Solidity: function lock(uint256 wad) returns()
func (_MCDADM *MCDADMTransactorSession) Lock(wad *big.Int) (*types.Transaction, error) {
	return _MCDADM.Contract.Lock(&_MCDADM.TransactOpts, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDADM *MCDADMTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDADM *MCDADMSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.SetAuthority(&_MCDADM.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDADM *MCDADMTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.SetAuthority(&_MCDADM.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDADM *MCDADMTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDADM *MCDADMSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.SetOwner(&_MCDADM.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDADM *MCDADMTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.SetOwner(&_MCDADM.TransactOpts, owner_)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMTransactor) SetPublicCapability(opts *bind.TransactOpts, code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setPublicCapability", code, sig, enabled)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMSession) SetPublicCapability(code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetPublicCapability(&_MCDADM.TransactOpts, code, sig, enabled)
}

// SetPublicCapability is a paid mutator transaction binding the contract method 0xc6b0263e.
//
// Solidity: function setPublicCapability(address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMTransactorSession) SetPublicCapability(code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetPublicCapability(&_MCDADM.TransactOpts, code, sig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMTransactor) SetRoleCapability(opts *bind.TransactOpts, role uint8, code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setRoleCapability", role, code, sig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMSession) SetRoleCapability(role uint8, code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetRoleCapability(&_MCDADM.TransactOpts, role, code, sig, enabled)
}

// SetRoleCapability is a paid mutator transaction binding the contract method 0x7d40583d.
//
// Solidity: function setRoleCapability(uint8 role, address code, bytes4 sig, bool enabled) returns()
func (_MCDADM *MCDADMTransactorSession) SetRoleCapability(role uint8, code common.Address, sig [4]byte, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetRoleCapability(&_MCDADM.TransactOpts, role, code, sig, enabled)
}

// SetRootUser is a paid mutator transaction binding the contract method 0xd381ba7c.
//
// Solidity: function setRootUser(address who, bool enabled) returns()
func (_MCDADM *MCDADMTransactor) SetRootUser(opts *bind.TransactOpts, who common.Address, enabled bool) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setRootUser", who, enabled)
}

// SetRootUser is a paid mutator transaction binding the contract method 0xd381ba7c.
//
// Solidity: function setRootUser(address who, bool enabled) returns()
func (_MCDADM *MCDADMSession) SetRootUser(who common.Address, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetRootUser(&_MCDADM.TransactOpts, who, enabled)
}

// SetRootUser is a paid mutator transaction binding the contract method 0xd381ba7c.
//
// Solidity: function setRootUser(address who, bool enabled) returns()
func (_MCDADM *MCDADMTransactorSession) SetRootUser(who common.Address, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetRootUser(&_MCDADM.TransactOpts, who, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address who, uint8 role, bool enabled) returns()
func (_MCDADM *MCDADMTransactor) SetUserRole(opts *bind.TransactOpts, who common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "setUserRole", who, role, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address who, uint8 role, bool enabled) returns()
func (_MCDADM *MCDADMSession) SetUserRole(who common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetUserRole(&_MCDADM.TransactOpts, who, role, enabled)
}

// SetUserRole is a paid mutator transaction binding the contract method 0x67aff484.
//
// Solidity: function setUserRole(address who, uint8 role, bool enabled) returns()
func (_MCDADM *MCDADMTransactorSession) SetUserRole(who common.Address, role uint8, enabled bool) (*types.Transaction, error) {
	return _MCDADM.Contract.SetUserRole(&_MCDADM.TransactOpts, who, role, enabled)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 slate) returns()
func (_MCDADM *MCDADMTransactor) Vote(opts *bind.TransactOpts, slate [32]byte) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "vote", slate)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 slate) returns()
func (_MCDADM *MCDADMSession) Vote(slate [32]byte) (*types.Transaction, error) {
	return _MCDADM.Contract.Vote(&_MCDADM.TransactOpts, slate)
}

// Vote is a paid mutator transaction binding the contract method 0xa69beaba.
//
// Solidity: function vote(bytes32 slate) returns()
func (_MCDADM *MCDADMTransactorSession) Vote(slate [32]byte) (*types.Transaction, error) {
	return _MCDADM.Contract.Vote(&_MCDADM.TransactOpts, slate)
}

// Vote0 is a paid mutator transaction binding the contract method 0xed081329.
//
// Solidity: function vote(address[] yays) returns(bytes32)
func (_MCDADM *MCDADMTransactor) Vote0(opts *bind.TransactOpts, yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.contract.Transact(opts, "vote0", yays)
}

// Vote0 is a paid mutator transaction binding the contract method 0xed081329.
//
// Solidity: function vote(address[] yays) returns(bytes32)
func (_MCDADM *MCDADMSession) Vote0(yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Vote0(&_MCDADM.TransactOpts, yays)
}

// Vote0 is a paid mutator transaction binding the contract method 0xed081329.
//
// Solidity: function vote(address[] yays) returns(bytes32)
func (_MCDADM *MCDADMTransactorSession) Vote0(yays []common.Address) (*types.Transaction, error) {
	return _MCDADM.Contract.Vote0(&_MCDADM.TransactOpts, yays)
}

// MCDADMEtchIterator is returned from FilterEtch and is used to iterate over the raw logs and unpacked data for Etch events raised by the MCDADM contract.
type MCDADMEtchIterator struct {
	Event *MCDADMEtch // Event containing the contract specifics and raw log

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
func (it *MCDADMEtchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDADMEtch)
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
		it.Event = new(MCDADMEtch)
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
func (it *MCDADMEtchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDADMEtchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDADMEtch represents a Etch event raised by the MCDADM contract.
type MCDADMEtch struct {
	Slate [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEtch is a free log retrieval operation binding the contract event 0x4f0892983790f53eea39a7a496f6cb40e8811b313871337b6a761efc6c67bb1f.
//
// Solidity: event Etch(bytes32 indexed slate)
func (_MCDADM *MCDADMFilterer) FilterEtch(opts *bind.FilterOpts, slate [][32]byte) (*MCDADMEtchIterator, error) {

	var slateRule []interface{}
	for _, slateItem := range slate {
		slateRule = append(slateRule, slateItem)
	}

	logs, sub, err := _MCDADM.contract.FilterLogs(opts, "Etch", slateRule)
	if err != nil {
		return nil, err
	}
	return &MCDADMEtchIterator{contract: _MCDADM.contract, event: "Etch", logs: logs, sub: sub}, nil
}

// WatchEtch is a free log subscription operation binding the contract event 0x4f0892983790f53eea39a7a496f6cb40e8811b313871337b6a761efc6c67bb1f.
//
// Solidity: event Etch(bytes32 indexed slate)
func (_MCDADM *MCDADMFilterer) WatchEtch(opts *bind.WatchOpts, sink chan<- *MCDADMEtch, slate [][32]byte) (event.Subscription, error) {

	var slateRule []interface{}
	for _, slateItem := range slate {
		slateRule = append(slateRule, slateItem)
	}

	logs, sub, err := _MCDADM.contract.WatchLogs(opts, "Etch", slateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDADMEtch)
				if err := _MCDADM.contract.UnpackLog(event, "Etch", log); err != nil {
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

// ParseEtch is a log parse operation binding the contract event 0x4f0892983790f53eea39a7a496f6cb40e8811b313871337b6a761efc6c67bb1f.
//
// Solidity: event Etch(bytes32 indexed slate)
func (_MCDADM *MCDADMFilterer) ParseEtch(log types.Log) (*MCDADMEtch, error) {
	event := new(MCDADMEtch)
	if err := _MCDADM.contract.UnpackLog(event, "Etch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDADMLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the MCDADM contract.
type MCDADMLogSetAuthorityIterator struct {
	Event *MCDADMLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *MCDADMLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDADMLogSetAuthority)
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
		it.Event = new(MCDADMLogSetAuthority)
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
func (it *MCDADMLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDADMLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDADMLogSetAuthority represents a LogSetAuthority event raised by the MCDADM contract.
type MCDADMLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDADM *MCDADMFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MCDADMLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDADM.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MCDADMLogSetAuthorityIterator{contract: _MCDADM.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDADM *MCDADMFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MCDADMLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDADM.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDADMLogSetAuthority)
				if err := _MCDADM.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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

// ParseLogSetAuthority is a log parse operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDADM *MCDADMFilterer) ParseLogSetAuthority(log types.Log) (*MCDADMLogSetAuthority, error) {
	event := new(MCDADMLogSetAuthority)
	if err := _MCDADM.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDADMLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the MCDADM contract.
type MCDADMLogSetOwnerIterator struct {
	Event *MCDADMLogSetOwner // Event containing the contract specifics and raw log

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
func (it *MCDADMLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDADMLogSetOwner)
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
		it.Event = new(MCDADMLogSetOwner)
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
func (it *MCDADMLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDADMLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDADMLogSetOwner represents a LogSetOwner event raised by the MCDADM contract.
type MCDADMLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDADM *MCDADMFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MCDADMLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDADM.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDADMLogSetOwnerIterator{contract: _MCDADM.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDADM *MCDADMFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MCDADMLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDADM.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDADMLogSetOwner)
				if err := _MCDADM.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDADM *MCDADMFilterer) ParseLogSetOwner(log types.Log) (*MCDADMLogSetOwner, error) {
	event := new(MCDADMLogSetOwner)
	if err := _MCDADM.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
