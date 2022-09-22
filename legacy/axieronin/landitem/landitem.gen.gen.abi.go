// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package landitem

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

// LanditemABI is the input ABI used to generate the binding from.
const LanditemABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"NonceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderUnwhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"SpenderWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addedMinters\",\"type\":\"address[]\"}],\"name\":\"addMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"name\":\"addTokenType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseTokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenTypes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"}],\"name\":\"deconstructItemId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"name\":\"editTokenMetadata\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getItemId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenTypeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenType\",\"type\":\"uint256\"}],\"name\":\"mintNew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removedMinters\",\"type\":\"address[]\"}],\"name\":\"removeMinters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseTokenURI\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_supported\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"unwhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Landitem is an auto generated Go binding around an Ethereum contract.
type Landitem struct {
	LanditemCaller     // Read-only binding to the contract
	LanditemTransactor // Write-only binding to the contract
	LanditemFilterer   // Log filterer for contract events
}

// LanditemCaller is an auto generated read-only Go binding around an Ethereum contract.
type LanditemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LanditemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LanditemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LanditemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LanditemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LanditemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LanditemSession struct {
	Contract     *Landitem         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LanditemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LanditemCallerSession struct {
	Contract *LanditemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LanditemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LanditemTransactorSession struct {
	Contract     *LanditemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LanditemRaw is an auto generated low-level Go binding around an Ethereum contract.
type LanditemRaw struct {
	Contract *Landitem // Generic contract binding to access the raw methods on
}

// LanditemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LanditemCallerRaw struct {
	Contract *LanditemCaller // Generic read-only contract binding to access the raw methods on
}

// LanditemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LanditemTransactorRaw struct {
	Contract *LanditemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLanditem creates a new instance of Landitem, bound to a specific deployed contract.
func NewLanditem(address common.Address, backend bind.ContractBackend) (*Landitem, error) {
	contract, err := bindLanditem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Landitem{LanditemCaller: LanditemCaller{contract: contract}, LanditemTransactor: LanditemTransactor{contract: contract}, LanditemFilterer: LanditemFilterer{contract: contract}}, nil
}

// NewLanditemCaller creates a new read-only instance of Landitem, bound to a specific deployed contract.
func NewLanditemCaller(address common.Address, caller bind.ContractCaller) (*LanditemCaller, error) {
	contract, err := bindLanditem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LanditemCaller{contract: contract}, nil
}

// NewLanditemTransactor creates a new write-only instance of Landitem, bound to a specific deployed contract.
func NewLanditemTransactor(address common.Address, transactor bind.ContractTransactor) (*LanditemTransactor, error) {
	contract, err := bindLanditem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LanditemTransactor{contract: contract}, nil
}

// NewLanditemFilterer creates a new log filterer instance of Landitem, bound to a specific deployed contract.
func NewLanditemFilterer(address common.Address, filterer bind.ContractFilterer) (*LanditemFilterer, error) {
	contract, err := bindLanditem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LanditemFilterer{contract: contract}, nil
}

// bindLanditem binds a generic wrapper to an already deployed contract.
func bindLanditem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LanditemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Landitem *LanditemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Landitem.Contract.LanditemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Landitem *LanditemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Landitem.Contract.LanditemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Landitem *LanditemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Landitem.Contract.LanditemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Landitem *LanditemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Landitem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Landitem *LanditemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Landitem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Landitem *LanditemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Landitem.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Landitem *LanditemCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Landitem *LanditemSession) Admin() (common.Address, error) {
	return _Landitem.Contract.Admin(&_Landitem.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Landitem *LanditemCallerSession) Admin() (common.Address, error) {
	return _Landitem.Contract.Admin(&_Landitem.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Landitem *LanditemCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Landitem *LanditemSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Landitem.Contract.BalanceOf(&_Landitem.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_Landitem *LanditemCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Landitem.Contract.BalanceOf(&_Landitem.CallOpts, _owner)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Landitem *LanditemCaller) BaseTokenURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "baseTokenURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Landitem *LanditemSession) BaseTokenURI() (string, error) {
	return _Landitem.Contract.BaseTokenURI(&_Landitem.CallOpts)
}

// BaseTokenURI is a free data retrieval call binding the contract method 0xd547cfb7.
//
// Solidity: function baseTokenURI() view returns(string)
func (_Landitem *LanditemCallerSession) BaseTokenURI() (string, error) {
	return _Landitem.Contract.BaseTokenURI(&_Landitem.CallOpts)
}

// DeconstructItemId is a free data retrieval call binding the contract method 0x229566b0.
//
// Solidity: function deconstructItemId(uint256 _itemId) pure returns(uint256 _tokenType, uint256 _tokenId)
func (_Landitem *LanditemCaller) DeconstructItemId(opts *bind.CallOpts, _itemId *big.Int) (struct {
	TokenType *big.Int
	TokenId   *big.Int
}, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "deconstructItemId", _itemId)

	outstruct := new(struct {
		TokenType *big.Int
		TokenId   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DeconstructItemId is a free data retrieval call binding the contract method 0x229566b0.
//
// Solidity: function deconstructItemId(uint256 _itemId) pure returns(uint256 _tokenType, uint256 _tokenId)
func (_Landitem *LanditemSession) DeconstructItemId(_itemId *big.Int) (struct {
	TokenType *big.Int
	TokenId   *big.Int
}, error) {
	return _Landitem.Contract.DeconstructItemId(&_Landitem.CallOpts, _itemId)
}

// DeconstructItemId is a free data retrieval call binding the contract method 0x229566b0.
//
// Solidity: function deconstructItemId(uint256 _itemId) pure returns(uint256 _tokenType, uint256 _tokenId)
func (_Landitem *LanditemCallerSession) DeconstructItemId(_itemId *big.Int) (struct {
	TokenType *big.Int
	TokenId   *big.Int
}, error) {
	return _Landitem.Contract.DeconstructItemId(&_Landitem.CallOpts, _itemId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Landitem *LanditemCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Landitem *LanditemSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Landitem.Contract.GetApproved(&_Landitem.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _operator)
func (_Landitem *LanditemCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Landitem.Contract.GetApproved(&_Landitem.CallOpts, _tokenId)
}

// GetItemId is a free data retrieval call binding the contract method 0x69bfb9de.
//
// Solidity: function getItemId(uint256 _tokenType, uint256 _tokenId) view returns(uint256)
func (_Landitem *LanditemCaller) GetItemId(opts *bind.CallOpts, _tokenType *big.Int, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "getItemId", _tokenType, _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetItemId is a free data retrieval call binding the contract method 0x69bfb9de.
//
// Solidity: function getItemId(uint256 _tokenType, uint256 _tokenId) view returns(uint256)
func (_Landitem *LanditemSession) GetItemId(_tokenType *big.Int, _tokenId *big.Int) (*big.Int, error) {
	return _Landitem.Contract.GetItemId(&_Landitem.CallOpts, _tokenType, _tokenId)
}

// GetItemId is a free data retrieval call binding the contract method 0x69bfb9de.
//
// Solidity: function getItemId(uint256 _tokenType, uint256 _tokenId) view returns(uint256)
func (_Landitem *LanditemCallerSession) GetItemId(_tokenType *big.Int, _tokenId *big.Int) (*big.Int, error) {
	return _Landitem.Contract.GetItemId(&_Landitem.CallOpts, _tokenType, _tokenId)
}

// GetTokenTypeCount is a free data retrieval call binding the contract method 0x4b1264de.
//
// Solidity: function getTokenTypeCount() view returns(uint256)
func (_Landitem *LanditemCaller) GetTokenTypeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "getTokenTypeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenTypeCount is a free data retrieval call binding the contract method 0x4b1264de.
//
// Solidity: function getTokenTypeCount() view returns(uint256)
func (_Landitem *LanditemSession) GetTokenTypeCount() (*big.Int, error) {
	return _Landitem.Contract.GetTokenTypeCount(&_Landitem.CallOpts)
}

// GetTokenTypeCount is a free data retrieval call binding the contract method 0x4b1264de.
//
// Solidity: function getTokenTypeCount() view returns(uint256)
func (_Landitem *LanditemCallerSession) GetTokenTypeCount() (*big.Int, error) {
	return _Landitem.Contract.GetTokenTypeCount(&_Landitem.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Landitem *LanditemCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Landitem *LanditemSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Landitem.Contract.IsApprovedForAll(&_Landitem.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool _approved)
func (_Landitem *LanditemCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Landitem.Contract.IsApprovedForAll(&_Landitem.CallOpts, _owner, _operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Landitem *LanditemCaller) IsMinter(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "isMinter", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Landitem *LanditemSession) IsMinter(_addr common.Address) (bool, error) {
	return _Landitem.Contract.IsMinter(&_Landitem.CallOpts, _addr)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address _addr) view returns(bool)
func (_Landitem *LanditemCallerSession) IsMinter(_addr common.Address) (bool, error) {
	return _Landitem.Contract.IsMinter(&_Landitem.CallOpts, _addr)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Landitem *LanditemCaller) Minter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "minter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Landitem *LanditemSession) Minter(arg0 common.Address) (bool, error) {
	return _Landitem.Contract.Minter(&_Landitem.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x3dd08c38.
//
// Solidity: function minter(address ) view returns(bool)
func (_Landitem *LanditemCallerSession) Minter(arg0 common.Address) (bool, error) {
	return _Landitem.Contract.Minter(&_Landitem.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Landitem *LanditemCaller) Minters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Landitem *LanditemSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _Landitem.Contract.Minters(&_Landitem.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0x8623ec7b.
//
// Solidity: function minters(uint256 ) view returns(address)
func (_Landitem *LanditemCallerSession) Minters(arg0 *big.Int) (common.Address, error) {
	return _Landitem.Contract.Minters(&_Landitem.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Landitem *LanditemCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Landitem *LanditemSession) Name() (string, error) {
	return _Landitem.Contract.Name(&_Landitem.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Landitem *LanditemCallerSession) Name() (string, error) {
	return _Landitem.Contract.Name(&_Landitem.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Landitem *LanditemCaller) Nonces(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Landitem *LanditemSession) Nonces(arg0 *big.Int) (*big.Int, error) {
	return _Landitem.Contract.Nonces(&_Landitem.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x141a468c.
//
// Solidity: function nonces(uint256 ) view returns(uint256)
func (_Landitem *LanditemCallerSession) Nonces(arg0 *big.Int) (*big.Int, error) {
	return _Landitem.Contract.Nonces(&_Landitem.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Landitem *LanditemCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Landitem *LanditemSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Landitem.Contract.OwnerOf(&_Landitem.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_Landitem *LanditemCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Landitem.Contract.OwnerOf(&_Landitem.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Landitem *LanditemCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Landitem *LanditemSession) Paused() (bool, error) {
	return _Landitem.Contract.Paused(&_Landitem.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Landitem *LanditemCallerSession) Paused() (bool, error) {
	return _Landitem.Contract.Paused(&_Landitem.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Landitem *LanditemCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Landitem *LanditemSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Landitem.Contract.SupportsInterface(&_Landitem.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool _supported)
func (_Landitem *LanditemCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Landitem.Contract.SupportsInterface(&_Landitem.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Landitem *LanditemCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Landitem *LanditemSession) Symbol() (string, error) {
	return _Landitem.Contract.Symbol(&_Landitem.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Landitem *LanditemCallerSession) Symbol() (string, error) {
	return _Landitem.Contract.Symbol(&_Landitem.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0x49afc6e5.
//
// Solidity: function tokenBalance(uint256 ) view returns(uint256)
func (_Landitem *LanditemCaller) TokenBalance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "tokenBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalance is a free data retrieval call binding the contract method 0x49afc6e5.
//
// Solidity: function tokenBalance(uint256 ) view returns(uint256)
func (_Landitem *LanditemSession) TokenBalance(arg0 *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenBalance(&_Landitem.CallOpts, arg0)
}

// TokenBalance is a free data retrieval call binding the contract method 0x49afc6e5.
//
// Solidity: function tokenBalance(uint256 ) view returns(uint256)
func (_Landitem *LanditemCallerSession) TokenBalance(arg0 *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenBalance(&_Landitem.CallOpts, arg0)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "tokenByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenByIndex(&_Landitem.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenByIndex(&_Landitem.CallOpts, _index)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 ) view returns(string name, string symbol, string baseTokenURI)
func (_Landitem *LanditemCaller) TokenMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name         string
	Symbol       string
	BaseTokenURI string
}, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "tokenMetadata", arg0)

	outstruct := new(struct {
		Name         string
		Symbol       string
		BaseTokenURI string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BaseTokenURI = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 ) view returns(string name, string symbol, string baseTokenURI)
func (_Landitem *LanditemSession) TokenMetadata(arg0 *big.Int) (struct {
	Name         string
	Symbol       string
	BaseTokenURI string
}, error) {
	return _Landitem.Contract.TokenMetadata(&_Landitem.CallOpts, arg0)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 ) view returns(string name, string symbol, string baseTokenURI)
func (_Landitem *LanditemCallerSession) TokenMetadata(arg0 *big.Int) (struct {
	Name         string
	Symbol       string
	BaseTokenURI string
}, error) {
	return _Landitem.Contract.TokenMetadata(&_Landitem.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenOfOwnerByIndex(&_Landitem.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Landitem *LanditemCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Landitem.Contract.TokenOfOwnerByIndex(&_Landitem.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Landitem *LanditemCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Landitem *LanditemSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Landitem.Contract.TokenURI(&_Landitem.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string _uri)
func (_Landitem *LanditemCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Landitem.Contract.TokenURI(&_Landitem.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Landitem *LanditemCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Landitem *LanditemSession) TotalSupply() (*big.Int, error) {
	return _Landitem.Contract.TotalSupply(&_Landitem.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _supply)
func (_Landitem *LanditemCallerSession) TotalSupply() (*big.Int, error) {
	return _Landitem.Contract.TotalSupply(&_Landitem.CallOpts)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Landitem *LanditemCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Landitem.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Landitem *LanditemSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Landitem.Contract.Whitelisted(&_Landitem.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_Landitem *LanditemCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _Landitem.Contract.Whitelisted(&_Landitem.CallOpts, arg0)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Landitem *LanditemTransactor) AddMinters(opts *bind.TransactOpts, _addedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "addMinters", _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Landitem *LanditemSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.AddMinters(&_Landitem.TransactOpts, _addedMinters)
}

// AddMinters is a paid mutator transaction binding the contract method 0x71e2a657.
//
// Solidity: function addMinters(address[] _addedMinters) returns()
func (_Landitem *LanditemTransactorSession) AddMinters(_addedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.AddMinters(&_Landitem.TransactOpts, _addedMinters)
}

// AddTokenType is a paid mutator transaction binding the contract method 0xae81cf41.
//
// Solidity: function addTokenType(string _name, string _symbol, string _baseTokenURI) returns(uint256 _tokenType)
func (_Landitem *LanditemTransactor) AddTokenType(opts *bind.TransactOpts, _name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "addTokenType", _name, _symbol, _baseTokenURI)
}

// AddTokenType is a paid mutator transaction binding the contract method 0xae81cf41.
//
// Solidity: function addTokenType(string _name, string _symbol, string _baseTokenURI) returns(uint256 _tokenType)
func (_Landitem *LanditemSession) AddTokenType(_name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.AddTokenType(&_Landitem.TransactOpts, _name, _symbol, _baseTokenURI)
}

// AddTokenType is a paid mutator transaction binding the contract method 0xae81cf41.
//
// Solidity: function addTokenType(string _name, string _symbol, string _baseTokenURI) returns(uint256 _tokenType)
func (_Landitem *LanditemTransactorSession) AddTokenType(_name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.AddTokenType(&_Landitem.TransactOpts, _name, _symbol, _baseTokenURI)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.Approve(&_Landitem.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.Approve(&_Landitem.TransactOpts, _to, _tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd559f05b.
//
// Solidity: function batchMint(address[] _recipients, uint256[] _tokenTypes, uint256[] _tokenIds) returns()
func (_Landitem *LanditemTransactor) BatchMint(opts *bind.TransactOpts, _recipients []common.Address, _tokenTypes []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "batchMint", _recipients, _tokenTypes, _tokenIds)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd559f05b.
//
// Solidity: function batchMint(address[] _recipients, uint256[] _tokenTypes, uint256[] _tokenIds) returns()
func (_Landitem *LanditemSession) BatchMint(_recipients []common.Address, _tokenTypes []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.BatchMint(&_Landitem.TransactOpts, _recipients, _tokenTypes, _tokenIds)
}

// BatchMint is a paid mutator transaction binding the contract method 0xd559f05b.
//
// Solidity: function batchMint(address[] _recipients, uint256[] _tokenTypes, uint256[] _tokenIds) returns()
func (_Landitem *LanditemTransactorSession) BatchMint(_recipients []common.Address, _tokenTypes []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.BatchMint(&_Landitem.TransactOpts, _recipients, _tokenTypes, _tokenIds)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Landitem *LanditemTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Landitem *LanditemSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.ChangeAdmin(&_Landitem.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Landitem *LanditemTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.ChangeAdmin(&_Landitem.TransactOpts, _newAdmin)
}

// EditTokenMetadata is a paid mutator transaction binding the contract method 0xc4f6aef5.
//
// Solidity: function editTokenMetadata(uint256 _tokenType, string _name, string _symbol, string _baseTokenURI) returns()
func (_Landitem *LanditemTransactor) EditTokenMetadata(opts *bind.TransactOpts, _tokenType *big.Int, _name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "editTokenMetadata", _tokenType, _name, _symbol, _baseTokenURI)
}

// EditTokenMetadata is a paid mutator transaction binding the contract method 0xc4f6aef5.
//
// Solidity: function editTokenMetadata(uint256 _tokenType, string _name, string _symbol, string _baseTokenURI) returns()
func (_Landitem *LanditemSession) EditTokenMetadata(_tokenType *big.Int, _name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.EditTokenMetadata(&_Landitem.TransactOpts, _tokenType, _name, _symbol, _baseTokenURI)
}

// EditTokenMetadata is a paid mutator transaction binding the contract method 0xc4f6aef5.
//
// Solidity: function editTokenMetadata(uint256 _tokenType, string _name, string _symbol, string _baseTokenURI) returns()
func (_Landitem *LanditemTransactorSession) EditTokenMetadata(_tokenType *big.Int, _name string, _symbol string, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.EditTokenMetadata(&_Landitem.TransactOpts, _tokenType, _name, _symbol, _baseTokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _tokenType, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _tokenType *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "mint", _to, _tokenType, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _tokenType, uint256 _tokenId) returns()
func (_Landitem *LanditemSession) Mint(_to common.Address, _tokenType *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.Mint(&_Landitem.TransactOpts, _to, _tokenType, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _to, uint256 _tokenType, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactorSession) Mint(_to common.Address, _tokenType *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.Mint(&_Landitem.TransactOpts, _to, _tokenType, _tokenId)
}

// MintNew is a paid mutator transaction binding the contract method 0x97fa91c1.
//
// Solidity: function mintNew(address _to, uint256 _tokenType) returns()
func (_Landitem *LanditemTransactor) MintNew(opts *bind.TransactOpts, _to common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "mintNew", _to, _tokenType)
}

// MintNew is a paid mutator transaction binding the contract method 0x97fa91c1.
//
// Solidity: function mintNew(address _to, uint256 _tokenType) returns()
func (_Landitem *LanditemSession) MintNew(_to common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.MintNew(&_Landitem.TransactOpts, _to, _tokenType)
}

// MintNew is a paid mutator transaction binding the contract method 0x97fa91c1.
//
// Solidity: function mintNew(address _to, uint256 _tokenType) returns()
func (_Landitem *LanditemTransactorSession) MintNew(_to common.Address, _tokenType *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.MintNew(&_Landitem.TransactOpts, _to, _tokenType)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Landitem *LanditemTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Landitem *LanditemSession) Pause() (*types.Transaction, error) {
	return _Landitem.Contract.Pause(&_Landitem.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Landitem *LanditemTransactorSession) Pause() (*types.Transaction, error) {
	return _Landitem.Contract.Pause(&_Landitem.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Landitem *LanditemTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Landitem *LanditemSession) RemoveAdmin() (*types.Transaction, error) {
	return _Landitem.Contract.RemoveAdmin(&_Landitem.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Landitem *LanditemTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Landitem.Contract.RemoveAdmin(&_Landitem.TransactOpts)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Landitem *LanditemTransactor) RemoveMinters(opts *bind.TransactOpts, _removedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "removeMinters", _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Landitem *LanditemSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.RemoveMinters(&_Landitem.TransactOpts, _removedMinters)
}

// RemoveMinters is a paid mutator transaction binding the contract method 0x5fc1964f.
//
// Solidity: function removeMinters(address[] _removedMinters) returns()
func (_Landitem *LanditemTransactorSession) RemoveMinters(_removedMinters []common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.RemoveMinters(&_Landitem.TransactOpts, _removedMinters)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.SafeTransferFrom(&_Landitem.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.SafeTransferFrom(&_Landitem.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Landitem *LanditemTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Landitem *LanditemSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Landitem.Contract.SafeTransferFrom0(&_Landitem.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_Landitem *LanditemTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Landitem.Contract.SafeTransferFrom0(&_Landitem.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Landitem *LanditemTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Landitem *LanditemSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Landitem.Contract.SetApprovalForAll(&_Landitem.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Landitem *LanditemTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Landitem.Contract.SetApprovalForAll(&_Landitem.TransactOpts, _operator, _approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Landitem *LanditemTransactor) SetBaseTokenURI(opts *bind.TransactOpts, _baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "setBaseTokenURI", _baseTokenURI)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Landitem *LanditemSession) SetBaseTokenURI(_baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.SetBaseTokenURI(&_Landitem.TransactOpts, _baseTokenURI)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string _baseTokenURI) returns()
func (_Landitem *LanditemTransactorSession) SetBaseTokenURI(_baseTokenURI string) (*types.Transaction, error) {
	return _Landitem.Contract.SetBaseTokenURI(&_Landitem.TransactOpts, _baseTokenURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.TransferFrom(&_Landitem.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_Landitem *LanditemTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Landitem.Contract.TransferFrom(&_Landitem.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Landitem *LanditemTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Landitem *LanditemSession) Unpause() (*types.Transaction, error) {
	return _Landitem.Contract.Unpause(&_Landitem.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Landitem *LanditemTransactorSession) Unpause() (*types.Transaction, error) {
	return _Landitem.Contract.Unpause(&_Landitem.TransactOpts)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Landitem *LanditemTransactor) Unwhitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "unwhitelist", _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Landitem *LanditemSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.Unwhitelist(&_Landitem.TransactOpts, _spender)
}

// Unwhitelist is a paid mutator transaction binding the contract method 0x9a590427.
//
// Solidity: function unwhitelist(address _spender) returns()
func (_Landitem *LanditemTransactorSession) Unwhitelist(_spender common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.Unwhitelist(&_Landitem.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Landitem *LanditemTransactor) Whitelist(opts *bind.TransactOpts, _spender common.Address) (*types.Transaction, error) {
	return _Landitem.contract.Transact(opts, "whitelist", _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Landitem *LanditemSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.Whitelist(&_Landitem.TransactOpts, _spender)
}

// Whitelist is a paid mutator transaction binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address _spender) returns()
func (_Landitem *LanditemTransactorSession) Whitelist(_spender common.Address) (*types.Transaction, error) {
	return _Landitem.Contract.Whitelist(&_Landitem.TransactOpts, _spender)
}

// LanditemAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Landitem contract.
type LanditemAdminChangedIterator struct {
	Event *LanditemAdminChanged // Event containing the contract specifics and raw log

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
func (it *LanditemAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemAdminChanged)
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
		it.Event = new(LanditemAdminChanged)
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
func (it *LanditemAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemAdminChanged represents a AdminChanged event raised by the Landitem contract.
type LanditemAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Landitem *LanditemFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*LanditemAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &LanditemAdminChangedIterator{contract: _Landitem.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Landitem *LanditemFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LanditemAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemAdminChanged)
				if err := _Landitem.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseAdminChanged(log types.Log) (*LanditemAdminChanged, error) {
	event := new(LanditemAdminChanged)
	if err := _Landitem.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Landitem contract.
type LanditemAdminRemovedIterator struct {
	Event *LanditemAdminRemoved // Event containing the contract specifics and raw log

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
func (it *LanditemAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemAdminRemoved)
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
		it.Event = new(LanditemAdminRemoved)
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
func (it *LanditemAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemAdminRemoved represents a AdminRemoved event raised by the Landitem contract.
type LanditemAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Landitem *LanditemFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*LanditemAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &LanditemAdminRemovedIterator{contract: _Landitem.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Landitem *LanditemFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *LanditemAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemAdminRemoved)
				if err := _Landitem.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseAdminRemoved(log types.Log) (*LanditemAdminRemoved, error) {
	event := new(LanditemAdminRemoved)
	if err := _Landitem.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Landitem contract.
type LanditemApprovalIterator struct {
	Event *LanditemApproval // Event containing the contract specifics and raw log

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
func (it *LanditemApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemApproval)
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
		it.Event = new(LanditemApproval)
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
func (it *LanditemApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemApproval represents a Approval event raised by the Landitem contract.
type LanditemApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Landitem *LanditemFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*LanditemApprovalIterator, error) {

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

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LanditemApprovalIterator{contract: _Landitem.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_Landitem *LanditemFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LanditemApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemApproval)
				if err := _Landitem.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseApproval(log types.Log) (*LanditemApproval, error) {
	event := new(LanditemApproval)
	if err := _Landitem.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Landitem contract.
type LanditemApprovalForAllIterator struct {
	Event *LanditemApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LanditemApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemApprovalForAll)
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
		it.Event = new(LanditemApprovalForAll)
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
func (it *LanditemApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemApprovalForAll represents a ApprovalForAll event raised by the Landitem contract.
type LanditemApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Landitem *LanditemFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*LanditemApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &LanditemApprovalForAllIterator{contract: _Landitem.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Landitem *LanditemFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LanditemApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemApprovalForAll)
				if err := _Landitem.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseApprovalForAll(log types.Log) (*LanditemApprovalForAll, error) {
	event := new(LanditemApprovalForAll)
	if err := _Landitem.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the Landitem contract.
type LanditemMinterAddedIterator struct {
	Event *LanditemMinterAdded // Event containing the contract specifics and raw log

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
func (it *LanditemMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemMinterAdded)
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
		it.Event = new(LanditemMinterAdded)
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
func (it *LanditemMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemMinterAdded represents a MinterAdded event raised by the Landitem contract.
type LanditemMinterAdded struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_Landitem *LanditemFilterer) FilterMinterAdded(opts *bind.FilterOpts, _minter []common.Address) (*LanditemMinterAddedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return &LanditemMinterAddedIterator{contract: _Landitem.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed _minter)
func (_Landitem *LanditemFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *LanditemMinterAdded, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "MinterAdded", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemMinterAdded)
				if err := _Landitem.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseMinterAdded(log types.Log) (*LanditemMinterAdded, error) {
	event := new(LanditemMinterAdded)
	if err := _Landitem.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Landitem contract.
type LanditemMinterRemovedIterator struct {
	Event *LanditemMinterRemoved // Event containing the contract specifics and raw log

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
func (it *LanditemMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemMinterRemoved)
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
		it.Event = new(LanditemMinterRemoved)
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
func (it *LanditemMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemMinterRemoved represents a MinterRemoved event raised by the Landitem contract.
type LanditemMinterRemoved struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_Landitem *LanditemFilterer) FilterMinterRemoved(opts *bind.FilterOpts, _minter []common.Address) (*LanditemMinterRemovedIterator, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return &LanditemMinterRemovedIterator{contract: _Landitem.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed _minter)
func (_Landitem *LanditemFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *LanditemMinterRemoved, _minter []common.Address) (event.Subscription, error) {

	var _minterRule []interface{}
	for _, _minterItem := range _minter {
		_minterRule = append(_minterRule, _minterItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "MinterRemoved", _minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemMinterRemoved)
				if err := _Landitem.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseMinterRemoved(log types.Log) (*LanditemMinterRemoved, error) {
	event := new(LanditemMinterRemoved)
	if err := _Landitem.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemNonceUpdatedIterator is returned from FilterNonceUpdated and is used to iterate over the raw logs and unpacked data for NonceUpdated events raised by the Landitem contract.
type LanditemNonceUpdatedIterator struct {
	Event *LanditemNonceUpdated // Event containing the contract specifics and raw log

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
func (it *LanditemNonceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemNonceUpdated)
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
		it.Event = new(LanditemNonceUpdated)
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
func (it *LanditemNonceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemNonceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemNonceUpdated represents a NonceUpdated event raised by the Landitem contract.
type LanditemNonceUpdated struct {
	TokenId *big.Int
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNonceUpdated is a free log retrieval operation binding the contract event 0xcc2c68164f9f7f0c063ba98bcf89498c0f3f5e3acc32bf4ab46195ecb489c13b.
//
// Solidity: event NonceUpdated(uint256 indexed _tokenId, uint256 indexed _nonce)
func (_Landitem *LanditemFilterer) FilterNonceUpdated(opts *bind.FilterOpts, _tokenId []*big.Int, _nonce []*big.Int) (*LanditemNonceUpdatedIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "NonceUpdated", _tokenIdRule, _nonceRule)
	if err != nil {
		return nil, err
	}
	return &LanditemNonceUpdatedIterator{contract: _Landitem.contract, event: "NonceUpdated", logs: logs, sub: sub}, nil
}

// WatchNonceUpdated is a free log subscription operation binding the contract event 0xcc2c68164f9f7f0c063ba98bcf89498c0f3f5e3acc32bf4ab46195ecb489c13b.
//
// Solidity: event NonceUpdated(uint256 indexed _tokenId, uint256 indexed _nonce)
func (_Landitem *LanditemFilterer) WatchNonceUpdated(opts *bind.WatchOpts, sink chan<- *LanditemNonceUpdated, _tokenId []*big.Int, _nonce []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}
	var _nonceRule []interface{}
	for _, _nonceItem := range _nonce {
		_nonceRule = append(_nonceRule, _nonceItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "NonceUpdated", _tokenIdRule, _nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemNonceUpdated)
				if err := _Landitem.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseNonceUpdated(log types.Log) (*LanditemNonceUpdated, error) {
	event := new(LanditemNonceUpdated)
	if err := _Landitem.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Landitem contract.
type LanditemPausedIterator struct {
	Event *LanditemPaused // Event containing the contract specifics and raw log

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
func (it *LanditemPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemPaused)
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
		it.Event = new(LanditemPaused)
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
func (it *LanditemPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemPaused represents a Paused event raised by the Landitem contract.
type LanditemPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Landitem *LanditemFilterer) FilterPaused(opts *bind.FilterOpts) (*LanditemPausedIterator, error) {

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LanditemPausedIterator{contract: _Landitem.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Landitem *LanditemFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LanditemPaused) (event.Subscription, error) {

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemPaused)
				if err := _Landitem.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParsePaused(log types.Log) (*LanditemPaused, error) {
	event := new(LanditemPaused)
	if err := _Landitem.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemSpenderUnwhitelistedIterator is returned from FilterSpenderUnwhitelisted and is used to iterate over the raw logs and unpacked data for SpenderUnwhitelisted events raised by the Landitem contract.
type LanditemSpenderUnwhitelistedIterator struct {
	Event *LanditemSpenderUnwhitelisted // Event containing the contract specifics and raw log

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
func (it *LanditemSpenderUnwhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemSpenderUnwhitelisted)
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
		it.Event = new(LanditemSpenderUnwhitelisted)
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
func (it *LanditemSpenderUnwhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemSpenderUnwhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemSpenderUnwhitelisted represents a SpenderUnwhitelisted event raised by the Landitem contract.
type LanditemSpenderUnwhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderUnwhitelisted is a free log retrieval operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_Landitem *LanditemFilterer) FilterSpenderUnwhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*LanditemSpenderUnwhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &LanditemSpenderUnwhitelistedIterator{contract: _Landitem.contract, event: "SpenderUnwhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderUnwhitelisted is a free log subscription operation binding the contract event 0x1b3787b2ded581f9fe4b6277fdd8e10f36368c08f211dbcb089a1c8940582dea.
//
// Solidity: event SpenderUnwhitelisted(address indexed _spender)
func (_Landitem *LanditemFilterer) WatchSpenderUnwhitelisted(opts *bind.WatchOpts, sink chan<- *LanditemSpenderUnwhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "SpenderUnwhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemSpenderUnwhitelisted)
				if err := _Landitem.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseSpenderUnwhitelisted(log types.Log) (*LanditemSpenderUnwhitelisted, error) {
	event := new(LanditemSpenderUnwhitelisted)
	if err := _Landitem.contract.UnpackLog(event, "SpenderUnwhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemSpenderWhitelistedIterator is returned from FilterSpenderWhitelisted and is used to iterate over the raw logs and unpacked data for SpenderWhitelisted events raised by the Landitem contract.
type LanditemSpenderWhitelistedIterator struct {
	Event *LanditemSpenderWhitelisted // Event containing the contract specifics and raw log

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
func (it *LanditemSpenderWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemSpenderWhitelisted)
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
		it.Event = new(LanditemSpenderWhitelisted)
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
func (it *LanditemSpenderWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemSpenderWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemSpenderWhitelisted represents a SpenderWhitelisted event raised by the Landitem contract.
type LanditemSpenderWhitelisted struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderWhitelisted is a free log retrieval operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_Landitem *LanditemFilterer) FilterSpenderWhitelisted(opts *bind.FilterOpts, _spender []common.Address) (*LanditemSpenderWhitelistedIterator, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return &LanditemSpenderWhitelistedIterator{contract: _Landitem.contract, event: "SpenderWhitelisted", logs: logs, sub: sub}, nil
}

// WatchSpenderWhitelisted is a free log subscription operation binding the contract event 0x8fa80cb4996c90aaf328c499cae29ba648e2686924355d501d9e4755fc82360a.
//
// Solidity: event SpenderWhitelisted(address indexed _spender)
func (_Landitem *LanditemFilterer) WatchSpenderWhitelisted(opts *bind.WatchOpts, sink chan<- *LanditemSpenderWhitelisted, _spender []common.Address) (event.Subscription, error) {

	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "SpenderWhitelisted", _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemSpenderWhitelisted)
				if err := _Landitem.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseSpenderWhitelisted(log types.Log) (*LanditemSpenderWhitelisted, error) {
	event := new(LanditemSpenderWhitelisted)
	if err := _Landitem.contract.UnpackLog(event, "SpenderWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Landitem contract.
type LanditemTransferIterator struct {
	Event *LanditemTransfer // Event containing the contract specifics and raw log

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
func (it *LanditemTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemTransfer)
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
		it.Event = new(LanditemTransfer)
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
func (it *LanditemTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemTransfer represents a Transfer event raised by the Landitem contract.
type LanditemTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Landitem *LanditemFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*LanditemTransferIterator, error) {

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

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LanditemTransferIterator{contract: _Landitem.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_Landitem *LanditemFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LanditemTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemTransfer)
				if err := _Landitem.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseTransfer(log types.Log) (*LanditemTransfer, error) {
	event := new(LanditemTransfer)
	if err := _Landitem.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LanditemUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Landitem contract.
type LanditemUnpausedIterator struct {
	Event *LanditemUnpaused // Event containing the contract specifics and raw log

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
func (it *LanditemUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LanditemUnpaused)
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
		it.Event = new(LanditemUnpaused)
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
func (it *LanditemUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LanditemUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LanditemUnpaused represents a Unpaused event raised by the Landitem contract.
type LanditemUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Landitem *LanditemFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LanditemUnpausedIterator, error) {

	logs, sub, err := _Landitem.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LanditemUnpausedIterator{contract: _Landitem.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Landitem *LanditemFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LanditemUnpaused) (event.Subscription, error) {

	logs, sub, err := _Landitem.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LanditemUnpaused)
				if err := _Landitem.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Landitem *LanditemFilterer) ParseUnpaused(log types.Log) (*LanditemUnpaused, error) {
	event := new(LanditemUnpaused)
	if err := _Landitem.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
