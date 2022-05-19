// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GET_CDPS

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

// GETCDPSABI is the input ABI used to generate the binding from.
const GETCDPSABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsAsc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"getCdpsDesc\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"urns\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ilks\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GETCDPS is an auto generated Go binding around an Ethereum contract.
type GETCDPS struct {
	GETCDPSCaller     // Read-only binding to the contract
	GETCDPSTransactor // Write-only binding to the contract
	GETCDPSFilterer   // Log filterer for contract events
}

// GETCDPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type GETCDPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GETCDPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GETCDPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GETCDPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GETCDPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GETCDPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GETCDPSSession struct {
	Contract     *GETCDPS          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GETCDPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GETCDPSCallerSession struct {
	Contract *GETCDPSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GETCDPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GETCDPSTransactorSession struct {
	Contract     *GETCDPSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GETCDPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type GETCDPSRaw struct {
	Contract *GETCDPS // Generic contract binding to access the raw methods on
}

// GETCDPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GETCDPSCallerRaw struct {
	Contract *GETCDPSCaller // Generic read-only contract binding to access the raw methods on
}

// GETCDPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GETCDPSTransactorRaw struct {
	Contract *GETCDPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGETCDPS creates a new instance of GETCDPS, bound to a specific deployed contract.
func NewGETCDPS(address common.Address, backend bind.ContractBackend) (*GETCDPS, error) {
	contract, err := bindGETCDPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GETCDPS{GETCDPSCaller: GETCDPSCaller{contract: contract}, GETCDPSTransactor: GETCDPSTransactor{contract: contract}, GETCDPSFilterer: GETCDPSFilterer{contract: contract}}, nil
}

// NewGETCDPSCaller creates a new read-only instance of GETCDPS, bound to a specific deployed contract.
func NewGETCDPSCaller(address common.Address, caller bind.ContractCaller) (*GETCDPSCaller, error) {
	contract, err := bindGETCDPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GETCDPSCaller{contract: contract}, nil
}

// NewGETCDPSTransactor creates a new write-only instance of GETCDPS, bound to a specific deployed contract.
func NewGETCDPSTransactor(address common.Address, transactor bind.ContractTransactor) (*GETCDPSTransactor, error) {
	contract, err := bindGETCDPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GETCDPSTransactor{contract: contract}, nil
}

// NewGETCDPSFilterer creates a new log filterer instance of GETCDPS, bound to a specific deployed contract.
func NewGETCDPSFilterer(address common.Address, filterer bind.ContractFilterer) (*GETCDPSFilterer, error) {
	contract, err := bindGETCDPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GETCDPSFilterer{contract: contract}, nil
}

// bindGETCDPS binds a generic wrapper to an already deployed contract.
func bindGETCDPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GETCDPSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GETCDPS *GETCDPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GETCDPS.Contract.GETCDPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GETCDPS *GETCDPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GETCDPS.Contract.GETCDPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GETCDPS *GETCDPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GETCDPS.Contract.GETCDPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GETCDPS *GETCDPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GETCDPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GETCDPS *GETCDPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GETCDPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GETCDPS *GETCDPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GETCDPS.Contract.contract.Transact(opts, method, params...)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSCaller) GetCdpsAsc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _GETCDPS.contract.Call(opts, &out, "getCdpsAsc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _GETCDPS.Contract.GetCdpsAsc(&_GETCDPS.CallOpts, manager, guy)
}

// GetCdpsAsc is a free data retrieval call binding the contract method 0x1ce03f38.
//
// Solidity: function getCdpsAsc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSCallerSession) GetCdpsAsc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _GETCDPS.Contract.GetCdpsAsc(&_GETCDPS.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSCaller) GetCdpsDesc(opts *bind.CallOpts, manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	var out []interface{}
	err := _GETCDPS.contract.Call(opts, &out, "getCdpsDesc", manager, guy)

	outstruct := new(struct {
		Ids  []*big.Int
		Urns []common.Address
		Ilks [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ids = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Urns = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Ilks = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _GETCDPS.Contract.GetCdpsDesc(&_GETCDPS.CallOpts, manager, guy)
}

// GetCdpsDesc is a free data retrieval call binding the contract method 0x38f7acb4.
//
// Solidity: function getCdpsDesc(address manager, address guy) view returns(uint256[] ids, address[] urns, bytes32[] ilks)
func (_GETCDPS *GETCDPSCallerSession) GetCdpsDesc(manager common.Address, guy common.Address) (struct {
	Ids  []*big.Int
	Urns []common.Address
	Ilks [][32]byte
}, error) {
	return _GETCDPS.Contract.GetCdpsDesc(&_GETCDPS.CallOpts, manager, guy)
}
