// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PerpRewardVesting

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

// MerkleRedeemUpgradeSafeClaim is an auto generated low-level Go binding around an user-defined struct.
type MerkleRedeemUpgradeSafeClaim struct {
	Week        *big.Int
	Balance     *big.Int
	MerkleProof [][32]byte
}

// PerpRewardVestingABI is the input ABI used to generate the binding from.
const PerpRewardVestingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"candidate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"claimStatus\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_week\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimedBalance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimWeek\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"week\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMerkleRedeemUpgradeSafe.Claim[]\",\"name\":\"_claims\",\"type\":\"tuple[]\"}],\"name\":\"claimWeeks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLengthOfMerkleRoots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_vestingPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRootIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"merkleRootTimestampMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_week\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_totalAllocation\",\"type\":\"uint256\"}],\"name\":\"seedAllocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_week\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimedBalance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"weekMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PerpRewardVesting is an auto generated Go binding around an Ethereum contract.
type PerpRewardVesting struct {
	PerpRewardVestingCaller     // Read-only binding to the contract
	PerpRewardVestingTransactor // Write-only binding to the contract
	PerpRewardVestingFilterer   // Log filterer for contract events
}

// PerpRewardVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PerpRewardVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpRewardVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PerpRewardVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpRewardVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PerpRewardVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PerpRewardVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PerpRewardVestingSession struct {
	Contract     *PerpRewardVesting // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PerpRewardVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PerpRewardVestingCallerSession struct {
	Contract *PerpRewardVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PerpRewardVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PerpRewardVestingTransactorSession struct {
	Contract     *PerpRewardVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PerpRewardVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PerpRewardVestingRaw struct {
	Contract *PerpRewardVesting // Generic contract binding to access the raw methods on
}

// PerpRewardVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PerpRewardVestingCallerRaw struct {
	Contract *PerpRewardVestingCaller // Generic read-only contract binding to access the raw methods on
}

// PerpRewardVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PerpRewardVestingTransactorRaw struct {
	Contract *PerpRewardVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPerpRewardVesting creates a new instance of PerpRewardVesting, bound to a specific deployed contract.
func NewPerpRewardVesting(address common.Address, backend bind.ContractBackend) (*PerpRewardVesting, error) {
	contract, err := bindPerpRewardVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PerpRewardVesting{PerpRewardVestingCaller: PerpRewardVestingCaller{contract: contract}, PerpRewardVestingTransactor: PerpRewardVestingTransactor{contract: contract}, PerpRewardVestingFilterer: PerpRewardVestingFilterer{contract: contract}}, nil
}

// NewPerpRewardVestingCaller creates a new read-only instance of PerpRewardVesting, bound to a specific deployed contract.
func NewPerpRewardVestingCaller(address common.Address, caller bind.ContractCaller) (*PerpRewardVestingCaller, error) {
	contract, err := bindPerpRewardVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PerpRewardVestingCaller{contract: contract}, nil
}

// NewPerpRewardVestingTransactor creates a new write-only instance of PerpRewardVesting, bound to a specific deployed contract.
func NewPerpRewardVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*PerpRewardVestingTransactor, error) {
	contract, err := bindPerpRewardVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PerpRewardVestingTransactor{contract: contract}, nil
}

// NewPerpRewardVestingFilterer creates a new log filterer instance of PerpRewardVesting, bound to a specific deployed contract.
func NewPerpRewardVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*PerpRewardVestingFilterer, error) {
	contract, err := bindPerpRewardVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PerpRewardVestingFilterer{contract: contract}, nil
}

// bindPerpRewardVesting binds a generic wrapper to an already deployed contract.
func bindPerpRewardVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PerpRewardVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpRewardVesting *PerpRewardVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpRewardVesting.Contract.PerpRewardVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpRewardVesting *PerpRewardVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.PerpRewardVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpRewardVesting *PerpRewardVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.PerpRewardVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PerpRewardVesting *PerpRewardVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PerpRewardVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PerpRewardVesting *PerpRewardVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PerpRewardVesting *PerpRewardVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.contract.Transact(opts, method, params...)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCaller) Candidate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "candidate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingSession) Candidate() (common.Address, error) {
	return _PerpRewardVesting.Contract.Candidate(&_PerpRewardVesting.CallOpts)
}

// Candidate is a free data retrieval call binding the contract method 0x6c8381f8.
//
// Solidity: function candidate() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) Candidate() (common.Address, error) {
	return _PerpRewardVesting.Contract.Candidate(&_PerpRewardVesting.CallOpts)
}

// ClaimStatus is a free data retrieval call binding the contract method 0x47fb23c1.
//
// Solidity: function claimStatus(address _liquidityProvider, uint256 _begin, uint256 _end) view returns(bool[])
func (_PerpRewardVesting *PerpRewardVestingCaller) ClaimStatus(opts *bind.CallOpts, _liquidityProvider common.Address, _begin *big.Int, _end *big.Int) ([]bool, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "claimStatus", _liquidityProvider, _begin, _end)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ClaimStatus is a free data retrieval call binding the contract method 0x47fb23c1.
//
// Solidity: function claimStatus(address _liquidityProvider, uint256 _begin, uint256 _end) view returns(bool[])
func (_PerpRewardVesting *PerpRewardVestingSession) ClaimStatus(_liquidityProvider common.Address, _begin *big.Int, _end *big.Int) ([]bool, error) {
	return _PerpRewardVesting.Contract.ClaimStatus(&_PerpRewardVesting.CallOpts, _liquidityProvider, _begin, _end)
}

// ClaimStatus is a free data retrieval call binding the contract method 0x47fb23c1.
//
// Solidity: function claimStatus(address _liquidityProvider, uint256 _begin, uint256 _end) view returns(bool[])
func (_PerpRewardVesting *PerpRewardVestingCallerSession) ClaimStatus(_liquidityProvider common.Address, _begin *big.Int, _end *big.Int) ([]bool, error) {
	return _PerpRewardVesting.Contract.ClaimStatus(&_PerpRewardVesting.CallOpts, _liquidityProvider, _begin, _end)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_PerpRewardVesting *PerpRewardVestingCaller) Claimed(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "claimed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_PerpRewardVesting *PerpRewardVestingSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _PerpRewardVesting.Contract.Claimed(&_PerpRewardVesting.CallOpts, arg0, arg1)
}

// Claimed is a free data retrieval call binding the contract method 0x120aa877.
//
// Solidity: function claimed(uint256 , address ) view returns(bool)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) Claimed(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _PerpRewardVesting.Contract.Claimed(&_PerpRewardVesting.CallOpts, arg0, arg1)
}

// GetLengthOfMerkleRoots is a free data retrieval call binding the contract method 0x39144f50.
//
// Solidity: function getLengthOfMerkleRoots() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCaller) GetLengthOfMerkleRoots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "getLengthOfMerkleRoots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLengthOfMerkleRoots is a free data retrieval call binding the contract method 0x39144f50.
//
// Solidity: function getLengthOfMerkleRoots() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingSession) GetLengthOfMerkleRoots() (*big.Int, error) {
	return _PerpRewardVesting.Contract.GetLengthOfMerkleRoots(&_PerpRewardVesting.CallOpts)
}

// GetLengthOfMerkleRoots is a free data retrieval call binding the contract method 0x39144f50.
//
// Solidity: function getLengthOfMerkleRoots() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) GetLengthOfMerkleRoots() (*big.Int, error) {
	return _PerpRewardVesting.Contract.GetLengthOfMerkleRoots(&_PerpRewardVesting.CallOpts)
}

// MerkleRootIndexes is a free data retrieval call binding the contract method 0xf373579f.
//
// Solidity: function merkleRootIndexes(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCaller) MerkleRootIndexes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "merkleRootIndexes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootIndexes is a free data retrieval call binding the contract method 0xf373579f.
//
// Solidity: function merkleRootIndexes(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingSession) MerkleRootIndexes(arg0 *big.Int) (*big.Int, error) {
	return _PerpRewardVesting.Contract.MerkleRootIndexes(&_PerpRewardVesting.CallOpts, arg0)
}

// MerkleRootIndexes is a free data retrieval call binding the contract method 0xf373579f.
//
// Solidity: function merkleRootIndexes(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) MerkleRootIndexes(arg0 *big.Int) (*big.Int, error) {
	return _PerpRewardVesting.Contract.MerkleRootIndexes(&_PerpRewardVesting.CallOpts, arg0)
}

// MerkleRootTimestampMap is a free data retrieval call binding the contract method 0xe2bd3e35.
//
// Solidity: function merkleRootTimestampMap(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCaller) MerkleRootTimestampMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "merkleRootTimestampMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MerkleRootTimestampMap is a free data retrieval call binding the contract method 0xe2bd3e35.
//
// Solidity: function merkleRootTimestampMap(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingSession) MerkleRootTimestampMap(arg0 *big.Int) (*big.Int, error) {
	return _PerpRewardVesting.Contract.MerkleRootTimestampMap(&_PerpRewardVesting.CallOpts, arg0)
}

// MerkleRootTimestampMap is a free data retrieval call binding the contract method 0xe2bd3e35.
//
// Solidity: function merkleRootTimestampMap(uint256 ) view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) MerkleRootTimestampMap(arg0 *big.Int) (*big.Int, error) {
	return _PerpRewardVesting.Contract.MerkleRootTimestampMap(&_PerpRewardVesting.CallOpts, arg0)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 _begin, uint256 _end) view returns(bytes32[])
func (_PerpRewardVesting *PerpRewardVestingCaller) MerkleRoots(opts *bind.CallOpts, _begin *big.Int, _end *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "merkleRoots", _begin, _end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 _begin, uint256 _end) view returns(bytes32[])
func (_PerpRewardVesting *PerpRewardVestingSession) MerkleRoots(_begin *big.Int, _end *big.Int) ([][32]byte, error) {
	return _PerpRewardVesting.Contract.MerkleRoots(&_PerpRewardVesting.CallOpts, _begin, _end)
}

// MerkleRoots is a free data retrieval call binding the contract method 0x39436b00.
//
// Solidity: function merkleRoots(uint256 _begin, uint256 _end) view returns(bytes32[])
func (_PerpRewardVesting *PerpRewardVestingCallerSession) MerkleRoots(_begin *big.Int, _end *big.Int) ([][32]byte, error) {
	return _PerpRewardVesting.Contract.MerkleRoots(&_PerpRewardVesting.CallOpts, _begin, _end)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingSession) Owner() (common.Address, error) {
	return _PerpRewardVesting.Contract.Owner(&_PerpRewardVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) Owner() (common.Address, error) {
	return _PerpRewardVesting.Contract.Owner(&_PerpRewardVesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingSession) Token() (common.Address, error) {
	return _PerpRewardVesting.Contract.Token(&_PerpRewardVesting.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) Token() (common.Address, error) {
	return _PerpRewardVesting.Contract.Token(&_PerpRewardVesting.CallOpts)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) view returns(bool valid)
func (_PerpRewardVesting *PerpRewardVestingCaller) VerifyClaim(opts *bind.CallOpts, _liquidityProvider common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "verifyClaim", _liquidityProvider, _week, _claimedBalance, _merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) view returns(bool valid)
func (_PerpRewardVesting *PerpRewardVestingSession) VerifyClaim(_liquidityProvider common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _PerpRewardVesting.Contract.VerifyClaim(&_PerpRewardVesting.CallOpts, _liquidityProvider, _week, _claimedBalance, _merkleProof)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xeb0d07f5.
//
// Solidity: function verifyClaim(address _liquidityProvider, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) view returns(bool valid)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) VerifyClaim(_liquidityProvider common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (bool, error) {
	return _PerpRewardVesting.Contract.VerifyClaim(&_PerpRewardVesting.CallOpts, _liquidityProvider, _week, _claimedBalance, _merkleProof)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCaller) VestingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "vestingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingSession) VestingPeriod() (*big.Int, error) {
	return _PerpRewardVesting.Contract.VestingPeriod(&_PerpRewardVesting.CallOpts)
}

// VestingPeriod is a free data retrieval call binding the contract method 0x7313ee5a.
//
// Solidity: function vestingPeriod() view returns(uint256)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) VestingPeriod() (*big.Int, error) {
	return _PerpRewardVesting.Contract.VestingPeriod(&_PerpRewardVesting.CallOpts)
}

// WeekMerkleRoots is a free data retrieval call binding the contract method 0xdd8c9c9d.
//
// Solidity: function weekMerkleRoots(uint256 ) view returns(bytes32)
func (_PerpRewardVesting *PerpRewardVestingCaller) WeekMerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PerpRewardVesting.contract.Call(opts, &out, "weekMerkleRoots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WeekMerkleRoots is a free data retrieval call binding the contract method 0xdd8c9c9d.
//
// Solidity: function weekMerkleRoots(uint256 ) view returns(bytes32)
func (_PerpRewardVesting *PerpRewardVestingSession) WeekMerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _PerpRewardVesting.Contract.WeekMerkleRoots(&_PerpRewardVesting.CallOpts, arg0)
}

// WeekMerkleRoots is a free data retrieval call binding the contract method 0xdd8c9c9d.
//
// Solidity: function weekMerkleRoots(uint256 ) view returns(bytes32)
func (_PerpRewardVesting *PerpRewardVestingCallerSession) WeekMerkleRoots(arg0 *big.Int) ([32]byte, error) {
	return _PerpRewardVesting.Contract.WeekMerkleRoots(&_PerpRewardVesting.CallOpts, arg0)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _account, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) ClaimWeek(opts *bind.TransactOpts, _account common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "claimWeek", _account, _week, _claimedBalance, _merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _account, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) returns()
func (_PerpRewardVesting *PerpRewardVestingSession) ClaimWeek(_account common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.ClaimWeek(&_PerpRewardVesting.TransactOpts, _account, _week, _claimedBalance, _merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address _account, uint256 _week, uint256 _claimedBalance, bytes32[] _merkleProof) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) ClaimWeek(_account common.Address, _week *big.Int, _claimedBalance *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.ClaimWeek(&_PerpRewardVesting.TransactOpts, _account, _week, _claimedBalance, _merkleProof)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address _account, (uint256,uint256,bytes32[])[] _claims) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) ClaimWeeks(opts *bind.TransactOpts, _account common.Address, _claims []MerkleRedeemUpgradeSafeClaim) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "claimWeeks", _account, _claims)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address _account, (uint256,uint256,bytes32[])[] _claims) returns()
func (_PerpRewardVesting *PerpRewardVestingSession) ClaimWeeks(_account common.Address, _claims []MerkleRedeemUpgradeSafeClaim) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.ClaimWeeks(&_PerpRewardVesting.TransactOpts, _account, _claims)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address _account, (uint256,uint256,bytes32[])[] _claims) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) ClaimWeeks(_account common.Address, _claims []MerkleRedeemUpgradeSafeClaim) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.ClaimWeeks(&_PerpRewardVesting.TransactOpts, _account, _claims)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _token, uint256 _vestingPeriod) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _vestingPeriod *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "initialize", _token, _vestingPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _token, uint256 _vestingPeriod) returns()
func (_PerpRewardVesting *PerpRewardVestingSession) Initialize(_token common.Address, _vestingPeriod *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.Initialize(&_PerpRewardVesting.TransactOpts, _token, _vestingPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _token, uint256 _vestingPeriod) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) Initialize(_token common.Address, _vestingPeriod *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.Initialize(&_PerpRewardVesting.TransactOpts, _token, _vestingPeriod)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpRewardVesting *PerpRewardVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.RenounceOwnership(&_PerpRewardVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.RenounceOwnership(&_PerpRewardVesting.TransactOpts)
}

// SeedAllocations is a paid mutator transaction binding the contract method 0x4cd488ab.
//
// Solidity: function seedAllocations(uint256 _week, bytes32 _merkleRoot, uint256 _totalAllocation) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) SeedAllocations(opts *bind.TransactOpts, _week *big.Int, _merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "seedAllocations", _week, _merkleRoot, _totalAllocation)
}

// SeedAllocations is a paid mutator transaction binding the contract method 0x4cd488ab.
//
// Solidity: function seedAllocations(uint256 _week, bytes32 _merkleRoot, uint256 _totalAllocation) returns()
func (_PerpRewardVesting *PerpRewardVestingSession) SeedAllocations(_week *big.Int, _merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.SeedAllocations(&_PerpRewardVesting.TransactOpts, _week, _merkleRoot, _totalAllocation)
}

// SeedAllocations is a paid mutator transaction binding the contract method 0x4cd488ab.
//
// Solidity: function seedAllocations(uint256 _week, bytes32 _merkleRoot, uint256 _totalAllocation) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) SeedAllocations(_week *big.Int, _merkleRoot [32]byte, _totalAllocation *big.Int) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.SeedAllocations(&_PerpRewardVesting.TransactOpts, _week, _merkleRoot, _totalAllocation)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PerpRewardVesting *PerpRewardVestingSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.SetOwner(&_PerpRewardVesting.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.SetOwner(&_PerpRewardVesting.TransactOpts, newOwner)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_PerpRewardVesting *PerpRewardVestingTransactor) UpdateOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PerpRewardVesting.contract.Transact(opts, "updateOwner")
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_PerpRewardVesting *PerpRewardVestingSession) UpdateOwner() (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.UpdateOwner(&_PerpRewardVesting.TransactOpts)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0xbc5920ba.
//
// Solidity: function updateOwner() returns()
func (_PerpRewardVesting *PerpRewardVestingTransactorSession) UpdateOwner() (*types.Transaction, error) {
	return _PerpRewardVesting.Contract.UpdateOwner(&_PerpRewardVesting.TransactOpts)
}

// PerpRewardVestingClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the PerpRewardVesting contract.
type PerpRewardVestingClaimedIterator struct {
	Event *PerpRewardVestingClaimed // Event containing the contract specifics and raw log

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
func (it *PerpRewardVestingClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpRewardVestingClaimed)
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
		it.Event = new(PerpRewardVestingClaimed)
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
func (it *PerpRewardVestingClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpRewardVestingClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpRewardVestingClaimed represents a Claimed event raised by the PerpRewardVesting contract.
type PerpRewardVestingClaimed struct {
	Claimant common.Address
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address _claimant, uint256 _balance)
func (_PerpRewardVesting *PerpRewardVestingFilterer) FilterClaimed(opts *bind.FilterOpts) (*PerpRewardVestingClaimedIterator, error) {

	logs, sub, err := _PerpRewardVesting.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &PerpRewardVestingClaimedIterator{contract: _PerpRewardVesting.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address _claimant, uint256 _balance)
func (_PerpRewardVesting *PerpRewardVestingFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *PerpRewardVestingClaimed) (event.Subscription, error) {

	logs, sub, err := _PerpRewardVesting.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpRewardVestingClaimed)
				if err := _PerpRewardVesting.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xd8138f8a3f377c5259ca548e70e4c2de94f129f5a11036a15b69513cba2b426a.
//
// Solidity: event Claimed(address _claimant, uint256 _balance)
func (_PerpRewardVesting *PerpRewardVestingFilterer) ParseClaimed(log types.Log) (*PerpRewardVestingClaimed, error) {
	event := new(PerpRewardVestingClaimed)
	if err := _PerpRewardVesting.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PerpRewardVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PerpRewardVesting contract.
type PerpRewardVestingOwnershipTransferredIterator struct {
	Event *PerpRewardVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PerpRewardVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PerpRewardVestingOwnershipTransferred)
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
		it.Event = new(PerpRewardVestingOwnershipTransferred)
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
func (it *PerpRewardVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PerpRewardVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PerpRewardVestingOwnershipTransferred represents a OwnershipTransferred event raised by the PerpRewardVesting contract.
type PerpRewardVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpRewardVesting *PerpRewardVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PerpRewardVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PerpRewardVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PerpRewardVestingOwnershipTransferredIterator{contract: _PerpRewardVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpRewardVesting *PerpRewardVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PerpRewardVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PerpRewardVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PerpRewardVestingOwnershipTransferred)
				if err := _PerpRewardVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PerpRewardVesting *PerpRewardVestingFilterer) ParseOwnershipTransferred(log types.Log) (*PerpRewardVestingOwnershipTransferred, error) {
	event := new(PerpRewardVestingOwnershipTransferred)
	if err := _PerpRewardVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
