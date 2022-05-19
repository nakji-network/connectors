// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package land

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

// LandABI is the input ABI used to generate the binding from.
const LandABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"NonceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderUnwhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addedMinters\",\"type\":\"address[]\"}],\"name\":\"addMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"int256[]\",\"name\":\"_rows\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"_cols\",\"type\":\"int256[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"decodeTokenId\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"estateData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"}],\"name\":\"getEstateData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"}],\"name\":\"getTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"landOfOwner\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"},{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"}],\"name\":\"ownerOfLand\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedMinters\",\"type\":\"address[]\"}],\"name\":\"removeMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_row\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_col\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"setEstateData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_supported\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"unwhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Land is an auto generated Go binding around an Ethereum contract.
type Land struct {
	LandCaller     // Read-only binding to the contract
	LandTransactor // Write-only binding to the contract
	LandFilterer   // Log filterer for contract events
}

// LandCaller is an auto generated read-only Go binding around an Ethereum contract.
type LandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LandSession struct {
	Contract     *Land             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LandCallerSession struct {
	Contract *LandCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LandTransactorSession struct {
	Contract     *LandTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LandRaw is an auto generated low-level Go binding around an Ethereum contract.
type LandRaw struct {
	Contract *Land // Generic contract binding to access the raw methods on
}

// LandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LandCallerRaw struct {
	Contract *LandCaller // Generic read-only contract binding to access the raw methods on
}

// LandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LandTransactorRaw struct {
	Contract *LandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLand creates a new instance of Land, bound to a specific deployed contract.
func NewLand(address common.Address, backend bind.ContractBackend) (*Land, error) {
	contract, err := bindLand(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Land{LandCaller: LandCaller{contract: contract}, LandTransactor: LandTransactor{contract: contract}, LandFilterer: LandFilterer{contract: contract}}, nil
}

// NewLandCaller creates a new read-only instance of Land, bound to a specific deployed contract.
func NewLandCaller(address common.Address, caller bind.ContractCaller) (*LandCaller, error) {
	contract, err := bindLand(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LandCaller{contract: contract}, nil
}

// NewLandTransactor creates a new write-only instance of Land, bound to a specific deployed contract.
func NewLandTransactor(address common.Address, transactor bind.ContractTransactor) (*LandTransactor, error) {
	contract, err := bindLand(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LandTransactor{contract: contract}, nil
}

// NewLandFilterer creates a new log filterer instance of Land, bound to a specific deployed contract.
func NewLandFilterer(address common.Address, filterer bind.ContractFilterer) (*LandFilterer, error) {
	contract, err := bindLand(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LandFilterer{contract: contract}, nil
}

// bindLand binds a generic wrapper to an already deployed contract.
func bindLand(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LandABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Land.Contract.LandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.LandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Land *LandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Land.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Land *LandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Land *LandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Land.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Land *LandCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Land *LandSession) Admin() (common.Address, error) {
	return _Land.Contract.Admin(&_Land.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Land *LandCallerSession) Admin() (common.Address, error) {
	return _Land.Contract.Admin(&_Land.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Land *LandCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Land *LandSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Land.Contract.BalanceOf(&_Land.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Land *LandCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Land.Contract.BalanceOf(&_Land.CallOpts, _owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Land *LandCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Land *LandSession) BaseTokenURI() (string, error) {
	return _Land.Contract.BaseTokenURI(&_Land.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Land *LandCallerSession) BaseTokenURI() (string, error) {
	return _Land.Contract.BaseTokenURI(&_Land.CallOpts)
}

// DecodeTokenId is a free data retrieval call binding the contract method 0x7efd9112.
//
// Solidity: function decodeTokenId(uint256 _tokenId) pure returns(int256 _row, int256 _col)
func (_Land *LandCaller) DecodeTokenId(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Row *big.Int
	Col *big.Int
}, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "decodeTokenId", _tokenId)

	outstruct := new(struct {
		Row *big.Int
		Col *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Row = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Col = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DecodeTokenId is a free data retrieval call binding the contract method 0x7efd9112.
//
// Solidity: function decodeTokenId(uint256 _tokenId) pure returns(int256 _row, int256 _col)
func (_Land *LandSession) DecodeTokenId(_tokenId *big.Int) (struct {
	Row *big.Int
	Col *big.Int
}, error) {
	return _Land.Contract.DecodeTokenId(&_Land.CallOpts, _tokenId)
}

// DecodeTokenId is a free data retrieval call binding the contract method 0x7efd9112.
//
// Solidity: function decodeTokenId(uint256 _tokenId) pure returns(int256 _row, int256 _col)
func (_Land *LandCallerSession) DecodeTokenId(_tokenId *big.Int) (struct {
	Row *big.Int
	Col *big.Int
}, error) {
	return _Land.Contract.DecodeTokenId(&_Land.CallOpts, _tokenId)
}

// EstateData is a free data retrieval call binding the contract method 0x28fa17b6.
//
// Solidity: function estateData(uint256 ) view returns(string)
func (_Land *LandCaller) EstateData(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "estateData", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EstateData is a free data retrieval call binding the contract method 0x28fa17b6.
//
// Solidity: function estateData(uint256 ) view returns(string)
func (_Land *LandSession) EstateData(arg0 *big.Int) (string, error) {
	return _Land.Contract.EstateData(&_Land.CallOpts, arg0)
}

// EstateData is a free data retrieval call binding the contract method 0x28fa17b6.
//
// Solidity: function estateData(uint256 ) view returns(string)
func (_Land *LandCallerSession) EstateData(arg0 *big.Int) (string, error) {
	return _Land.Contract.EstateData(&_Land.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Land *LandCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Land *LandSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.GetApproved(&_Land.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Land *LandCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.GetApproved(&_Land.CallOpts, _tokenId)
}

// GetEstateData is a free data retrieval call binding the contract method 0x84a6afea.
//
// Solidity: function getEstateData(int256 _row, int256 _col) view returns(string)
func (_Land *LandCaller) GetEstateData(opts *bind.CallOpts, _row *big.Int, _col *big.Int) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getEstateData", _row, _col)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEstateData is a free data retrieval call binding the contract method 0x84a6afea.
//
// Solidity: function getEstateData(int256 _row, int256 _col) view returns(string)
func (_Land *LandSession) GetEstateData(_row *big.Int, _col *big.Int) (string, error) {
	return _Land.Contract.GetEstateData(&_Land.CallOpts, _row, _col)
}

// GetEstateData is a free data retrieval call binding the contract method 0x84a6afea.
//
// Solidity: function getEstateData(int256 _row, int256 _col) view returns(string)
func (_Land *LandCallerSession) GetEstateData(_row *big.Int, _col *big.Int) (string, error) {
	return _Land.Contract.GetEstateData(&_Land.CallOpts, _row, _col)
}

// GetTokenId is a free data retrieval call binding the contract method 0xaa210fb0.
//
// Solidity: function getTokenId(int256 _row, int256 _col) pure returns(uint256)
func (_Land *LandCaller) GetTokenId(opts *bind.CallOpts, _row *big.Int, _col *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "getTokenId", _row, _col)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenId is a free data retrieval call binding the contract method 0xaa210fb0.
//
// Solidity: function getTokenId(int256 _row, int256 _col) pure returns(uint256)
func (_Land *LandSession) GetTokenId(_row *big.Int, _col *big.Int) (*big.Int, error) {
	return _Land.Contract.GetTokenId(&_Land.CallOpts, _row, _col)
}

// GetTokenId is a free data retrieval call binding the contract method 0xaa210fb0.
//
// Solidity: function getTokenId(int256 _row, int256 _col) pure returns(uint256)
func (_Land *LandCallerSession) GetTokenId(_row *big.Int, _col *big.Int) (*big.Int, error) {
	return _Land.Contract.GetTokenId(&_Land.CallOpts, _row, _col)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Land *LandCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Land *LandSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Land.Contract.IsApprovedForAll(&_Land.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Land *LandCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Land.Contract.IsApprovedForAll(&_Land.CallOpts, _owner, _operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Land *LandCaller) IsMinter(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "isMinter", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Land *LandSession) IsMinter(_addr common.Address) (bool, error) {
	return _Land.Contract.IsMinter(&_Land.CallOpts, _addr)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Land *LandCallerSession) IsMinter(_addr common.Address) (bool, error) {
	return _Land.Contract.IsMinter(&_Land.CallOpts, _addr)
}

// LandOfOwner is a free data retrieval call binding the contract method 0x2f9dde8d.
//
// Solidity: function landOfOwner(address _owner) view returns(int256[], int256[])
func (_Land *LandCaller) LandOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "landOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// LandOfOwner is a free data retrieval call binding the contract method 0x2f9dde8d.
//
// Solidity: function landOfOwner(address _owner) view returns(int256[], int256[])
func (_Land *LandSession) LandOfOwner(_owner common.Address) ([]*big.Int, []*big.Int, error) {
	return _Land.Contract.LandOfOwner(&_Land.CallOpts, _owner)
}

// LandOfOwner is a free data retrieval call binding the contract method 0x2f9dde8d.
//
// Solidity: function landOfOwner(address _owner) view returns(int256[], int256[])
func (_Land *LandCallerSession) LandOfOwner(_owner common.Address) ([]*big.Int, []*big.Int, error) {
	return _Land.Contract.LandOfOwner(&_Land.CallOpts, _owner)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Land *LandCaller) Minter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "minter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Land *LandSession) Minter(arg0 common.Address) (bool, error) {
	return _Land.Contract.Minter(&_Land.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Land *LandCallerSession) Minter(arg0 common.Address) (bool, error) {
	return _Land.Contract.Minter(&_Land.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Land *LandCaller) Minters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Land *LandSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _Land.Contract.Minters(&_Land.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Land *LandCallerSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _Land.Contract.Minters(&_Land.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandSession) Name() (string, error) {
	return _Land.Contract.Name(&_Land.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Land *LandCallerSession) Name() (string, error) {
	return _Land.Contract.Name(&_Land.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Land *LandCaller) Nonces(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Land *LandSession) Nonces(arg0 *big.Int) (*big.Int, error) {
	return _Land.Contract.Nonces(&_Land.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Land *LandCallerSession) Nonces(arg0 *big.Int) (*big.Int, error) {
	return _Land.Contract.Nonces(&_Land.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Land *LandCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Land *LandSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOf(&_Land.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Land *LandCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOf(&_Land.CallOpts, _tokenId)
}

// OwnerOfLand is a free data retrieval call binding the contract method 0x1080f251.
//
// Solidity: function ownerOfLand(int256 _row, int256 _col) view returns(address)
func (_Land *LandCaller) OwnerOfLand(opts *bind.CallOpts, _row *big.Int, _col *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "ownerOfLand", _row, _col)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOfLand is a free data retrieval call binding the contract method 0x1080f251.
//
// Solidity: function ownerOfLand(int256 _row, int256 _col) view returns(address)
func (_Land *LandSession) OwnerOfLand(_row *big.Int, _col *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOfLand(&_Land.CallOpts, _row, _col)
}

// OwnerOfLand is a free data retrieval call binding the contract method 0x1080f251.
//
// Solidity: function ownerOfLand(int256 _row, int256 _col) view returns(address)
func (_Land *LandCallerSession) OwnerOfLand(_row *big.Int, _col *big.Int) (common.Address, error) {
	return _Land.Contract.OwnerOfLand(&_Land.CallOpts, _row, _col)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandSession) Paused() (bool, error) {
	return _Land.Contract.Paused(&_Land.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Land *LandCallerSession) Paused() (bool, error) {
	return _Land.Contract.Paused(&_Land.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Land *LandCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Land *LandSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Land.Contract.SupportsInterface(&_Land.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Land *LandCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Land.Contract.SupportsInterface(&_Land.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandSession) Symbol() (string, error) {
	return _Land.Contract.Symbol(&_Land.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Land *LandCallerSession) Symbol() (string, error) {
	return _Land.Contract.Symbol(&_Land.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenByIndex(&_Land.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenByIndex(&_Land.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenOfOwnerByIndex(&_Land.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Land *LandCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Land.Contract.TokenOfOwnerByIndex(&_Land.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Land *LandCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Land *LandSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Land.Contract.TokenURI(&_Land.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Land *LandCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Land.Contract.TokenURI(&_Land.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Land *LandCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Land *LandSession) TotalSupply() (*big.Int, error) {
	return _Land.Contract.TotalSupply(&_Land.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Land *LandCallerSession) TotalSupply() (*big.Int, error) {
	return _Land.Contract.TotalSupply(&_Land.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Land *LandCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Land.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Land *LandSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Land.Contract.Whitelisted(&_Land.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Land *LandCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Land.Contract.Whitelisted(&_Land.CallOpts, arg0)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Land *LandTransactor) AddMinters(opts *bind.TransactOpts, _addedMinters []common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "addMinters", _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Land *LandSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _Land.Contract.AddMinters(&_Land.TransactOpts, _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Land *LandTransactorSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _Land.Contract.AddMinters(&_Land.TransactOpts, _addedMinters)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Land *LandTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Land *LandSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Approve(&_Land.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Land *LandTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Approve(&_Land.TransactOpts, _to, _tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd9b4fcf9.
//
// Solidity: function batchMint(address[] _owners, int256[] _rows, int256[] _cols) returns()
func (_Land *LandTransactor) BatchMint(opts *bind.TransactOpts, _owners []common.Address, _rows []*big.Int, _cols []*big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "batchMint", _owners, _rows, _cols)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd9b4fcf9.
//
// Solidity: function batchMint(address[] _owners, int256[] _rows, int256[] _cols) returns()
func (_Land *LandSession) BatchMint(_owners []common.Address, _rows []*big.Int, _cols []*big.Int) (*types.Transaction, error) {
	return _Land.Contract.BatchMint(&_Land.TransactOpts, _owners, _rows, _cols)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd9b4fcf9.
//
// Solidity: function batchMint(address[] _owners, int256[] _rows, int256[] _cols) returns()
func (_Land *LandTransactorSession) BatchMint(_owners []common.Address, _rows []*big.Int, _cols []*big.Int) (*types.Transaction, error) {
	return _Land.Contract.BatchMint(&_Land.TransactOpts, _owners, _rows, _cols)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Land *LandTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Land *LandSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Land.Contract.ChangeAdmin(&_Land.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Land *LandTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Land.Contract.ChangeAdmin(&_Land.TransactOpts, _newAdmin)
}

// Mint is a paid mutator transaction binding the contract method 0xa91aa9a8.
//
// Solidity: function mint(address _to, int256 _row, int256 _col) returns(uint256 _itemId)
func (_Land *LandTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _row *big.Int, _col *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "mint", _to, _row, _col)
}

// Mint is a paid mutator transaction binding the contract method 0xa91aa9a8.
//
// Solidity: function mint(address _to, int256 _row, int256 _col) returns(uint256 _itemId)
func (_Land *LandSession) Mint(_to common.Address, _row *big.Int, _col *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Mint(&_Land.TransactOpts, _to, _row, _col)
}

// Mint is a paid mutator transaction binding the contract method 0xa91aa9a8.
//
// Solidity: function mint(address _to, int256 _row, int256 _col) returns(uint256 _itemId)
func (_Land *LandTransactorSession) Mint(_to common.Address, _row *big.Int, _col *big.Int) (*types.Transaction, error) {
	return _Land.Contract.Mint(&_Land.TransactOpts, _to, _row, _col)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandSession) Pause() (*types.Transaction, error) {
	return _Land.Contract.Pause(&_Land.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Land *LandTransactorSession) Pause() (*types.Transaction, error) {
	return _Land.Contract.Pause(&_Land.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Land *LandTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Land *LandSession) RemoveAdmin() (*types.Transaction, error) {
	return _Land.Contract.RemoveAdmin(&_Land.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Land *LandTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Land.Contract.RemoveAdmin(&_Land.TransactOpts)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Land *LandTransactor) RemoveMinters(opts *bind.TransactOpts, _removedMinters []common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "removeMinters", _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Land *LandSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _Land.Contract.RemoveMinters(&_Land.TransactOpts, _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Land *LandTransactorSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _Land.Contract.RemoveMinters(&_Land.TransactOpts, _removedMinters)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom(&_Land.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom(&_Land.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Land *LandTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Land *LandSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom0(&_Land.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Land *LandTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Land.Contract.SafeTransferFrom0(&_Land.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Land *LandTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Land *LandSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Land.Contract.SetApprovalForAll(&_Land.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Land *LandTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Land.Contract.SetApprovalForAll(&_Land.TransactOpts, _operator, _approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Land *LandTransactor) SetBaseTokenURI(opts *bind.TransactOpts, _baseTokenURI string) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "setBaseTokenURI", _baseTokenURI)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Land *LandSession) SetBaseTokenURI(_baseTokenURI string) (*types.Transaction, error) {
	return _Land.Contract.SetBaseTokenURI(&_Land.TransactOpts, _baseTokenURI)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Land *LandTransactorSession) SetBaseTokenURI(_baseTokenURI string) (*types.Transaction, error) {
	return _Land.Contract.SetBaseTokenURI(&_Land.TransactOpts, _baseTokenURI)
}

// SetEstateData is a paid mutator transaction binding the contract method 0x4f60a056.
//
// Solidity: function setEstateData(int256 _row, int256 _col, string _data) returns()
func (_Land *LandTransactor) SetEstateData(opts *bind.TransactOpts, _row *big.Int, _col *big.Int, _data string) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "setEstateData", _row, _col, _data)
}

// SetEstateData is a paid mutator transaction binding the contract method 0x4f60a056.
//
// Solidity: function setEstateData(int256 _row, int256 _col, string _data) returns()
func (_Land *LandSession) SetEstateData(_row *big.Int, _col *big.Int, _data string) (*types.Transaction, error) {
	return _Land.Contract.SetEstateData(&_Land.TransactOpts, _row, _col, _data)
}

// SetEstateData is a paid mutator transaction binding the contract method 0x4f60a056.
//
// Solidity: function setEstateData(int256 _row, int256 _col, string _data) returns()
func (_Land *LandTransactorSession) SetEstateData(_row *big.Int, _col *big.Int, _data string) (*types.Transaction, error) {
	return _Land.Contract.SetEstateData(&_Land.TransactOpts, _row, _col, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.TransferFrom(&_Land.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Land *LandTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Land.Contract.TransferFrom(&_Land.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandSession) Unpause() (*types.Transaction, error) {
	return _Land.Contract.Unpause(&_Land.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Land *LandTransactorSession) Unpause() (*types.Transaction, error) {
	return _Land.Contract.Unpause(&_Land.TransactOpts)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Land *LandTransactor) Unwhitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "unwhitelist", _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Land *LandSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _Land.Contract.Unwhitelist(&_Land.TransactOpts, _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Land *LandTransactorSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _Land.Contract.Unwhitelist(&_Land.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Land *LandTransactor) Whitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _Land.contract.Transact(opts, "whitelist", _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Land *LandSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _Land.Contract.Whitelist(&_Land.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Land *LandTransactorSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _Land.Contract.Whitelist(&_Land.TransactOpts, _spender)
}

// LandAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Land contract.
type LandAdminChangedIterator struct {
	Event *LandAdminChanged // Event containing the contract specifics and raw log

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
func (it *LandAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandAdminChanged)
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
		it.Event = new(LandAdminChanged)
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
func (it *LandAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandAdminChanged represents a AdminChanged event raised by the Land contract.
type LandAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Land *LandFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*LandAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &LandAdminChangedIterator{contract: _Land.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Land *LandFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LandAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandAdminChanged)
				if err := _Land.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Land *LandFilterer) ParseAdminChanged(log types.Log) (*LandAdminChanged, error) {
	event := new(LandAdminChanged)
	if err := _Land.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Land contract.
type LandAdminRemovedIterator struct {
	Event *LandAdminRemoved // Event containing the contract specifics and raw log

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
func (it *LandAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandAdminRemoved)
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
		it.Event = new(LandAdminRemoved)
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
func (it *LandAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandAdminRemoved represents a AdminRemoved event raised by the Land contract.
type LandAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Land *LandFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*LandAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &LandAdminRemovedIterator{contract: _Land.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Land *LandFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *LandAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandAdminRemoved)
				if err := _Land.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Land *LandFilterer) ParseAdminRemoved(log types.Log) (*LandAdminRemoved, error) {
	event := new(LandAdminRemoved)
	if err := _Land.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Land contract.
type LandApprovalIterator struct {
	Event *LandApproval // Event containing the contract specifics and raw log

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
func (it *LandApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandApproval)
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
		it.Event = new(LandApproval)
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
func (it *LandApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandApproval represents a Approval event raised by the Land contract.
type LandApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Land *LandFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*LandApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LandApprovalIterator{contract: _Land.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Land *LandFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LandApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandApproval)
				if err := _Land.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Land *LandFilterer) ParseApproval(log types.Log) (*LandApproval, error) {
	event := new(LandApproval)
	if err := _Land.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Land contract.
type LandApprovalForAllIterator struct {
	Event *LandApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LandApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandApprovalForAll)
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
		it.Event = new(LandApprovalForAll)
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
func (it *LandApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandApprovalForAll represents a ApprovalForAll event raised by the Land contract.
type LandApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Land *LandFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*LandApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &LandApprovalForAllIterator{contract: _Land.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Land *LandFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LandApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandApprovalForAll)
				if err := _Land.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Land *LandFilterer) ParseApprovalForAll(log types.Log) (*LandApprovalForAll, error) {
	event := new(LandApprovalForAll)
	if err := _Land.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the Land contract.
type LandMinterAddedIterator struct {
	Event *LandMinterAdded // Event containing the contract specifics and raw log

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
func (it *LandMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandMinterAdded)
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
		it.Event = new(LandMinterAdded)
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
func (it *LandMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandMinterAdded represents a MinterAdded event raised by the Land contract.
type LandMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_Land *LandFilterer) FilterMinterAdded(opts *bind.FilterOpts, _minter []common.Address) (*LandMinterAddedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return &LandMinterAddedIterator{contract: _Land.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_Land *LandFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *LandMinterAdded, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandMinterAdded)
				if err := _Land.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_Land *LandFilterer) ParseMinterAdded(log types.Log) (*LandMinterAdded, error) {
	event := new(LandMinterAdded)
	if err := _Land.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Land contract.
type LandMinterRemovedIterator struct {
	Event *LandMinterRemoved // Event containing the contract specifics and raw log

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
func (it *LandMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandMinterRemoved)
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
		it.Event = new(LandMinterRemoved)
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
func (it *LandMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandMinterRemoved represents a MinterRemoved event raised by the Land contract.
type LandMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_Land *LandFilterer) FilterMinterRemoved(opts *bind.FilterOpts, _minter []common.Address) (*LandMinterRemovedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return &LandMinterRemovedIterator{contract: _Land.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_Land *LandFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *LandMinterRemoved, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandMinterRemoved)
				if err := _Land.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_Land *LandFilterer) ParseMinterRemoved(log types.Log) (*LandMinterRemoved, error) {
	event := new(LandMinterRemoved)
	if err := _Land.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandNonceUpdatedIterator is returned from FilterNonceUpdated and is used to iterate over the raw logs and unpacked data for NonceUpdated events raised by the Land contract.
type LandNonceUpdatedIterator struct {
	Event *LandNonceUpdated // Event containing the contract specifics and raw log

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
func (it *LandNonceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandNonceUpdated)
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
		it.Event = new(LandNonceUpdated)
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
func (it *LandNonceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandNonceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandNonceUpdated represents a NonceUpdated event raised by the Land contract.
type LandNonceUpdated struct {
	TokenId *big.Int
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNonceUpdated is a free log retrieval operation binding the contract event 0xcc2c68164f9f7f0c063ba98bcf89498c0f3f5e3acc32bf4ab46195ecb489c13b.
//
// Solidity: event NonceUpdated(uint256 indexed _tokenId, uint256 indexed _nonce)
func (_Land *LandFilterer) FilterNonceUpdated(opts *bind.FilterOpts, _tokenId []*big.Int, _nonce []*big.Int) (*LandNonceUpdatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "NonceUpdated", _tokenIdRule, _nonceRule)
	if err != nil {
		return nil, err
	}
	return &LandNonceUpdatedIterator{contract: _Land.contract, event: "NonceUpdated", logs: logs, sub: sub}, nil
}

// WatchNonceUpdated is a free log subscription operation binding the contract event 0xcc2c68164f9f7f0c063ba98bcf89498c0f3f5e3acc32bf4ab46195ecb489c13b.
//
// Solidity: event NonceUpdated(uint256 indexed _tokenId, uint256 indexed _nonce)
func (_Land *LandFilterer) WatchNonceUpdated(opts *bind.WatchOpts, sink chan<- *LandNonceUpdated, _tokenId []*big.Int, _nonce []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "NonceUpdated", _tokenIdRule, _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandNonceUpdated)
				if err := _Land.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
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

// ParseNonceUpdated is a log parse operation binding the contract event 0xcc2c68164f9f7f0c063ba98bcf89498c0f3f5e3acc32bf4ab46195ecb489c13b.
//
// Solidity: event NonceUpdated(uint256 indexed _tokenId, uint256 indexed _nonce)
func (_Land *LandFilterer) ParseNonceUpdated(log types.Log) (*LandNonceUpdated, error) {
	event := new(LandNonceUpdated)
	if err := _Land.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Land contract.
type LandPausedIterator struct {
	Event *LandPaused // Event containing the contract specifics and raw log

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
func (it *LandPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandPaused)
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
		it.Event = new(LandPaused)
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
func (it *LandPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandPaused represents a Paused event raised by the Land contract.
type LandPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Land *LandFilterer) FilterPaused(opts *bind.FilterOpts) (*LandPausedIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LandPausedIterator{contract: _Land.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Land *LandFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LandPaused) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandPaused)
				if err := _Land.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Land *LandFilterer) ParsePaused(log types.Log) (*LandPaused, error) {
	event := new(LandPaused)
	if err := _Land.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandSpenderUnwhitelistedIterator is returned from FilterSpenderUnwhitelisted and is used to iterate over the raw logs and unpacked data for SpenderUnwhitelisted events raised by the Land contract.
type LandSpenderUnwhitelistedIterator struct {
	Event *LandSpenderUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *LandSpenderUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandSpenderUnwhitelisted)
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
		it.Event = new(LandSpenderUnwhitelisted)
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
func (it *LandSpenderUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandSpenderUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandSpenderUnwhitelisted represents a SpenderUnwhitelisted event raised by the Land contract.
type LandSpenderUnwhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderUnwhitelisted is a free log retrieval operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_Land *LandFilterer) FilterSpenderUnwhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*LandSpenderUnwhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &LandSpenderUnwhitelistedIterator{contract: _Land.contract, event: "SpenderUnwhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderUnwhitelisted is a free log subscription operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_Land *LandFilterer) WatchSpenderUnwhitelisted(opts *bind.WatchOpts, sink chan<- *LandSpenderUnwhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandSpenderUnwhitelisted)
				if err := _Land.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
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

// ParseSpenderUnwhitelisted is a log parse operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_Land *LandFilterer) ParseSpenderUnwhitelisted(log types.Log) (*LandSpenderUnwhitelisted, error) {
	event := new(LandSpenderUnwhitelisted)
	if err := _Land.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandSpenderWhitelistedIterator is returned from FilterSpenderWhitelisted and is used to iterate over the raw logs and unpacked data for SpenderWhitelisted events raised by the Land contract.
type LandSpenderWhitelistedIterator struct {
	Event *LandSpenderWhitelisted // Event containing the contract specifics and raw log

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
func (it *LandSpenderWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandSpenderWhitelisted)
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
		it.Event = new(LandSpenderWhitelisted)
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
func (it *LandSpenderWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandSpenderWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandSpenderWhitelisted represents a SpenderWhitelisted event raised by the Land contract.
type LandSpenderWhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderWhitelisted is a free log retrieval operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_Land *LandFilterer) FilterSpenderWhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*LandSpenderWhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &LandSpenderWhitelistedIterator{contract: _Land.contract, event: "SpenderWhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderWhitelisted is a free log subscription operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_Land *LandFilterer) WatchSpenderWhitelisted(opts *bind.WatchOpts, sink chan<- *LandSpenderWhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandSpenderWhitelisted)
				if err := _Land.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
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

// ParseSpenderWhitelisted is a log parse operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_Land *LandFilterer) ParseSpenderWhitelisted(log types.Log) (*LandSpenderWhitelisted, error) {
	event := new(LandSpenderWhitelisted)
	if err := _Land.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Land contract.
type LandTransferIterator struct {
	Event *LandTransfer // Event containing the contract specifics and raw log

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
func (it *LandTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandTransfer)
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
		it.Event = new(LandTransfer)
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
func (it *LandTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandTransfer represents a Transfer event raised by the Land contract.
type LandTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Land *LandFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*LandTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Land.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LandTransferIterator{contract: _Land.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Land *LandFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LandTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Land.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandTransfer)
				if err := _Land.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Land *LandFilterer) ParseTransfer(log types.Log) (*LandTransfer, error) {
	event := new(LandTransfer)
	if err := _Land.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LandUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Land contract.
type LandUnpausedIterator struct {
	Event *LandUnpaused // Event containing the contract specifics and raw log

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
func (it *LandUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LandUnpaused)
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
		it.Event = new(LandUnpaused)
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
func (it *LandUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LandUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LandUnpaused represents a Unpaused event raised by the Land contract.
type LandUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Land *LandFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LandUnpausedIterator, error) {

	logs, sub, err := _Land.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LandUnpausedIterator{contract: _Land.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Land *LandFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LandUnpaused) (event.Subscription, error) {

	logs, sub, err := _Land.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LandUnpaused)
				if err := _Land.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Land *LandFilterer) ParseUnpaused(log types.Log) (*LandUnpaused, error) {
	event := new(LandUnpaused)
	if err := _Land.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
