// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package axienft

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

// AxienftABI is the input ABI used to generate the binding from.
const AxienftABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approved\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_axieId\",\"type\":\"uint256\"},{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"rebirthAxie\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whitelistSetterAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketplaceManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setRetirementManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_geneScientist\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"bool\"}],\"name\":\"setGeneScientist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setMarketplaceManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"retirementManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSetter\",\"type\":\"address\"}],\"name\":\"setWhitelistSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"spawnAxie\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_marketplace\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"bool\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_byeSayer\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"bool\"}],\"name\":\"setByeSayer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setSpawningManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setGeneManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedByeSayer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spawningManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedGeneScientist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spawner\",\"type\":\"address\"},{\"name\":\"_whitelisted\",\"type\":\"bool\"}],\"name\":\"setSpawner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedMarketplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_axieId\",\"type\":\"uint256\"}],\"name\":\"getAxie\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_axieId\",\"type\":\"uint256\"},{\"name\":\"_newGenes\",\"type\":\"uint256\"}],\"name\":\"evolveAxie\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenURIPrefix\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedSpawner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_axieId\",\"type\":\"uint256\"},{\"name\":\"_rip\",\"type\":\"bool\"}],\"name\":\"retireAxie\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenURISuffix\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_prefix\",\"type\":\"string\"},{\"name\":\"_suffix\",\"type\":\"string\"}],\"name\":\"setTokenURIAffixes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_axieId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"AxieSpawned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_axieId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"AxieRebirthed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_axieId\",\"type\":\"uint256\"}],\"name\":\"AxieRetired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_axieId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_oldGenes\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_newGenes\",\"type\":\"uint256\"}],\"name\":\"AxieEvolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// Axienft is an auto generated Go binding around an Ethereum contract.
type Axienft struct {
	AxienftCaller     // Read-only binding to the contract
	AxienftTransactor // Write-only binding to the contract
	AxienftFilterer   // Log filterer for contract events
}

// AxienftCaller is an auto generated read-only Go binding around an Ethereum contract.
type AxienftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxienftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AxienftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxienftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AxienftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AxienftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AxienftSession struct {
	Contract     *Axienft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AxienftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AxienftCallerSession struct {
	Contract *AxienftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AxienftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AxienftTransactorSession struct {
	Contract     *AxienftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AxienftRaw is an auto generated low-level Go binding around an Ethereum contract.
type AxienftRaw struct {
	Contract *Axienft // Generic contract binding to access the raw methods on
}

// AxienftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AxienftCallerRaw struct {
	Contract *AxienftCaller // Generic read-only contract binding to access the raw methods on
}

// AxienftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AxienftTransactorRaw struct {
	Contract *AxienftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAxienft creates a new instance of Axienft, bound to a specific deployed contract.
func NewAxienft(address common.Address, backend bind.ContractBackend) (*Axienft, error) {
	contract, err := bindAxienft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Axienft{AxienftCaller: AxienftCaller{contract: contract}, AxienftTransactor: AxienftTransactor{contract: contract}, AxienftFilterer: AxienftFilterer{contract: contract}}, nil
}

// NewAxienftCaller creates a new read-only instance of Axienft, bound to a specific deployed contract.
func NewAxienftCaller(address common.Address, caller bind.ContractCaller) (*AxienftCaller, error) {
	contract, err := bindAxienft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AxienftCaller{contract: contract}, nil
}

// NewAxienftTransactor creates a new write-only instance of Axienft, bound to a specific deployed contract.
func NewAxienftTransactor(address common.Address, transactor bind.ContractTransactor) (*AxienftTransactor, error) {
	contract, err := bindAxienft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AxienftTransactor{contract: contract}, nil
}

// NewAxienftFilterer creates a new log filterer instance of Axienft, bound to a specific deployed contract.
func NewAxienftFilterer(address common.Address, filterer bind.ContractFilterer) (*AxienftFilterer, error) {
	contract, err := bindAxienft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AxienftFilterer{contract: contract}, nil
}

// bindAxienft binds a generic wrapper to an already deployed contract.
func bindAxienft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AxienftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axienft *AxienftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axienft.Contract.AxienftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axienft *AxienftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axienft.Contract.AxienftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axienft *AxienftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axienft.Contract.AxienftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Axienft *AxienftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Axienft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Axienft *AxienftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axienft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Axienft *AxienftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Axienft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Axienft *AxienftCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Axienft *AxienftSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Axienft.Contract.BalanceOf(&_Axienft.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Axienft *AxienftCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Axienft.Contract.BalanceOf(&_Axienft.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Axienft *AxienftCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Axienft *AxienftSession) CeoAddress() (common.Address, error) {
	return _Axienft.Contract.CeoAddress(&_Axienft.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_Axienft *AxienftCallerSession) CeoAddress() (common.Address, error) {
	return _Axienft.Contract.CeoAddress(&_Axienft.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Axienft *AxienftCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Axienft *AxienftSession) CfoAddress() (common.Address, error) {
	return _Axienft.Contract.CfoAddress(&_Axienft.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_Axienft *AxienftCallerSession) CfoAddress() (common.Address, error) {
	return _Axienft.Contract.CfoAddress(&_Axienft.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Axienft *AxienftCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Axienft *AxienftSession) CooAddress() (common.Address, error) {
	return _Axienft.Contract.CooAddress(&_Axienft.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_Axienft *AxienftCallerSession) CooAddress() (common.Address, error) {
	return _Axienft.Contract.CooAddress(&_Axienft.CallOpts)
}

// GeneManager is a free data retrieval call binding the contract method 0x80173a19.
//
// Solidity: function geneManager() view returns(address)
func (_Axienft *AxienftCaller) GeneManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "geneManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneManager is a free data retrieval call binding the contract method 0x80173a19.
//
// Solidity: function geneManager() view returns(address)
func (_Axienft *AxienftSession) GeneManager() (common.Address, error) {
	return _Axienft.Contract.GeneManager(&_Axienft.CallOpts)
}

// GeneManager is a free data retrieval call binding the contract method 0x80173a19.
//
// Solidity: function geneManager() view returns(address)
func (_Axienft *AxienftCallerSession) GeneManager() (common.Address, error) {
	return _Axienft.Contract.GeneManager(&_Axienft.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Axienft.Contract.GetApproved(&_Axienft.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Axienft.Contract.GetApproved(&_Axienft.CallOpts, _tokenId)
}

// GetAxie is a free data retrieval call binding the contract method 0xa6472906.
//
// Solidity: function getAxie(uint256 _axieId) view returns(uint256, uint256)
func (_Axienft *AxienftCaller) GetAxie(opts *bind.CallOpts, _axieId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "getAxie", _axieId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAxie is a free data retrieval call binding the contract method 0xa6472906.
//
// Solidity: function getAxie(uint256 _axieId) view returns(uint256, uint256)
func (_Axienft *AxienftSession) GetAxie(_axieId *big.Int) (*big.Int, *big.Int, error) {
	return _Axienft.Contract.GetAxie(&_Axienft.CallOpts, _axieId)
}

// GetAxie is a free data retrieval call binding the contract method 0xa6472906.
//
// Solidity: function getAxie(uint256 _axieId) view returns(uint256, uint256)
func (_Axienft *AxienftCallerSession) GetAxie(_axieId *big.Int) (*big.Int, *big.Int, error) {
	return _Axienft.Contract.GetAxie(&_Axienft.CallOpts, _axieId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Axienft *AxienftCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Axienft *AxienftSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Axienft.Contract.IsApprovedForAll(&_Axienft.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_Axienft *AxienftCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Axienft.Contract.IsApprovedForAll(&_Axienft.CallOpts, _owner, _operator)
}

// MarketplaceManager is a free data retrieval call binding the contract method 0x1889500c.
//
// Solidity: function marketplaceManager() view returns(address)
func (_Axienft *AxienftCaller) MarketplaceManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "marketplaceManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceManager is a free data retrieval call binding the contract method 0x1889500c.
//
// Solidity: function marketplaceManager() view returns(address)
func (_Axienft *AxienftSession) MarketplaceManager() (common.Address, error) {
	return _Axienft.Contract.MarketplaceManager(&_Axienft.CallOpts)
}

// MarketplaceManager is a free data retrieval call binding the contract method 0x1889500c.
//
// Solidity: function marketplaceManager() view returns(address)
func (_Axienft *AxienftCallerSession) MarketplaceManager() (common.Address, error) {
	return _Axienft.Contract.MarketplaceManager(&_Axienft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Axienft *AxienftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Axienft *AxienftSession) Name() (string, error) {
	return _Axienft.Contract.Name(&_Axienft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Axienft *AxienftCallerSession) Name() (string, error) {
	return _Axienft.Contract.Name(&_Axienft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Axienft.Contract.OwnerOf(&_Axienft.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Axienft *AxienftCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Axienft.Contract.OwnerOf(&_Axienft.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Axienft *AxienftCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Axienft *AxienftSession) Paused() (bool, error) {
	return _Axienft.Contract.Paused(&_Axienft.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Axienft *AxienftCallerSession) Paused() (bool, error) {
	return _Axienft.Contract.Paused(&_Axienft.CallOpts)
}

// RetirementManager is a free data retrieval call binding the contract method 0x52ac882c.
//
// Solidity: function retirementManager() view returns(address)
func (_Axienft *AxienftCaller) RetirementManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "retirementManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RetirementManager is a free data retrieval call binding the contract method 0x52ac882c.
//
// Solidity: function retirementManager() view returns(address)
func (_Axienft *AxienftSession) RetirementManager() (common.Address, error) {
	return _Axienft.Contract.RetirementManager(&_Axienft.CallOpts)
}

// RetirementManager is a free data retrieval call binding the contract method 0x52ac882c.
//
// Solidity: function retirementManager() view returns(address)
func (_Axienft *AxienftCallerSession) RetirementManager() (common.Address, error) {
	return _Axienft.Contract.RetirementManager(&_Axienft.CallOpts)
}

// SpawningManager is a free data retrieval call binding the contract method 0x98f5ee5d.
//
// Solidity: function spawningManager() view returns(address)
func (_Axienft *AxienftCaller) SpawningManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "spawningManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpawningManager is a free data retrieval call binding the contract method 0x98f5ee5d.
//
// Solidity: function spawningManager() view returns(address)
func (_Axienft *AxienftSession) SpawningManager() (common.Address, error) {
	return _Axienft.Contract.SpawningManager(&_Axienft.CallOpts)
}

// SpawningManager is a free data retrieval call binding the contract method 0x98f5ee5d.
//
// Solidity: function spawningManager() view returns(address)
func (_Axienft *AxienftCallerSession) SpawningManager() (common.Address, error) {
	return _Axienft.Contract.SpawningManager(&_Axienft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Axienft *AxienftCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Axienft *AxienftSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Axienft.Contract.SupportsInterface(&_Axienft.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Axienft *AxienftCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Axienft.Contract.SupportsInterface(&_Axienft.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Axienft *AxienftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Axienft *AxienftSession) Symbol() (string, error) {
	return _Axienft.Contract.Symbol(&_Axienft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Axienft *AxienftCallerSession) Symbol() (string, error) {
	return _Axienft.Contract.Symbol(&_Axienft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_Axienft *AxienftCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "tokenByIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_Axienft *AxienftSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Axienft.Contract.TokenByIndex(&_Axienft.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 _index) view returns(uint256)
func (_Axienft *AxienftCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Axienft.Contract.TokenByIndex(&_Axienft.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Axienft *AxienftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Axienft *AxienftSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Axienft.Contract.TokenOfOwnerByIndex(&_Axienft.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_Axienft *AxienftCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Axienft.Contract.TokenOfOwnerByIndex(&_Axienft.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Axienft *AxienftCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Axienft *AxienftSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Axienft.Contract.TokenURI(&_Axienft.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Axienft *AxienftCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Axienft.Contract.TokenURI(&_Axienft.CallOpts, _tokenId)
}

// TokenURIPrefix is a free data retrieval call binding the contract method 0xc0ac9983.
//
// Solidity: function tokenURIPrefix() view returns(string)
func (_Axienft *AxienftCaller) TokenURIPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "tokenURIPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURIPrefix is a free data retrieval call binding the contract method 0xc0ac9983.
//
// Solidity: function tokenURIPrefix() view returns(string)
func (_Axienft *AxienftSession) TokenURIPrefix() (string, error) {
	return _Axienft.Contract.TokenURIPrefix(&_Axienft.CallOpts)
}

// TokenURIPrefix is a free data retrieval call binding the contract method 0xc0ac9983.
//
// Solidity: function tokenURIPrefix() view returns(string)
func (_Axienft *AxienftCallerSession) TokenURIPrefix() (string, error) {
	return _Axienft.Contract.TokenURIPrefix(&_Axienft.CallOpts)
}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Axienft *AxienftCaller) TokenURISuffix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "tokenURISuffix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Axienft *AxienftSession) TokenURISuffix() (string, error) {
	return _Axienft.Contract.TokenURISuffix(&_Axienft.CallOpts)
}

// TokenURISuffix is a free data retrieval call binding the contract method 0xdbbc853b.
//
// Solidity: function tokenURISuffix() view returns(string)
func (_Axienft *AxienftCallerSession) TokenURISuffix() (string, error) {
	return _Axienft.Contract.TokenURISuffix(&_Axienft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Axienft *AxienftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Axienft *AxienftSession) TotalSupply() (*big.Int, error) {
	return _Axienft.Contract.TotalSupply(&_Axienft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Axienft *AxienftCallerSession) TotalSupply() (*big.Int, error) {
	return _Axienft.Contract.TotalSupply(&_Axienft.CallOpts)
}

// WhitelistSetterAddress is a free data retrieval call binding the contract method 0x1412409a.
//
// Solidity: function whitelistSetterAddress() view returns(address)
func (_Axienft *AxienftCaller) WhitelistSetterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "whitelistSetterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistSetterAddress is a free data retrieval call binding the contract method 0x1412409a.
//
// Solidity: function whitelistSetterAddress() view returns(address)
func (_Axienft *AxienftSession) WhitelistSetterAddress() (common.Address, error) {
	return _Axienft.Contract.WhitelistSetterAddress(&_Axienft.CallOpts)
}

// WhitelistSetterAddress is a free data retrieval call binding the contract method 0x1412409a.
//
// Solidity: function whitelistSetterAddress() view returns(address)
func (_Axienft *AxienftCallerSession) WhitelistSetterAddress() (common.Address, error) {
	return _Axienft.Contract.WhitelistSetterAddress(&_Axienft.CallOpts)
}

// WhitelistedByeSayer is a free data retrieval call binding the contract method 0x9848146a.
//
// Solidity: function whitelistedByeSayer(address ) view returns(bool)
func (_Axienft *AxienftCaller) WhitelistedByeSayer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "whitelistedByeSayer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedByeSayer is a free data retrieval call binding the contract method 0x9848146a.
//
// Solidity: function whitelistedByeSayer(address ) view returns(bool)
func (_Axienft *AxienftSession) WhitelistedByeSayer(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedByeSayer(&_Axienft.CallOpts, arg0)
}

// WhitelistedByeSayer is a free data retrieval call binding the contract method 0x9848146a.
//
// Solidity: function whitelistedByeSayer(address ) view returns(bool)
func (_Axienft *AxienftCallerSession) WhitelistedByeSayer(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedByeSayer(&_Axienft.CallOpts, arg0)
}

// WhitelistedGeneScientist is a free data retrieval call binding the contract method 0x99165bf4.
//
// Solidity: function whitelistedGeneScientist(address ) view returns(bool)
func (_Axienft *AxienftCaller) WhitelistedGeneScientist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "whitelistedGeneScientist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedGeneScientist is a free data retrieval call binding the contract method 0x99165bf4.
//
// Solidity: function whitelistedGeneScientist(address ) view returns(bool)
func (_Axienft *AxienftSession) WhitelistedGeneScientist(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedGeneScientist(&_Axienft.CallOpts, arg0)
}

// WhitelistedGeneScientist is a free data retrieval call binding the contract method 0x99165bf4.
//
// Solidity: function whitelistedGeneScientist(address ) view returns(bool)
func (_Axienft *AxienftCallerSession) WhitelistedGeneScientist(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedGeneScientist(&_Axienft.CallOpts, arg0)
}

// WhitelistedMarketplace is a free data retrieval call binding the contract method 0xa306dfbe.
//
// Solidity: function whitelistedMarketplace(address ) view returns(bool)
func (_Axienft *AxienftCaller) WhitelistedMarketplace(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "whitelistedMarketplace", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedMarketplace is a free data retrieval call binding the contract method 0xa306dfbe.
//
// Solidity: function whitelistedMarketplace(address ) view returns(bool)
func (_Axienft *AxienftSession) WhitelistedMarketplace(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedMarketplace(&_Axienft.CallOpts, arg0)
}

// WhitelistedMarketplace is a free data retrieval call binding the contract method 0xa306dfbe.
//
// Solidity: function whitelistedMarketplace(address ) view returns(bool)
func (_Axienft *AxienftCallerSession) WhitelistedMarketplace(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedMarketplace(&_Axienft.CallOpts, arg0)
}

// WhitelistedSpawner is a free data retrieval call binding the contract method 0xc84c9e45.
//
// Solidity: function whitelistedSpawner(address ) view returns(bool)
func (_Axienft *AxienftCaller) WhitelistedSpawner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Axienft.contract.Call(opts, &out, "whitelistedSpawner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedSpawner is a free data retrieval call binding the contract method 0xc84c9e45.
//
// Solidity: function whitelistedSpawner(address ) view returns(bool)
func (_Axienft *AxienftSession) WhitelistedSpawner(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedSpawner(&_Axienft.CallOpts, arg0)
}

// WhitelistedSpawner is a free data retrieval call binding the contract method 0xc84c9e45.
//
// Solidity: function whitelistedSpawner(address ) view returns(bool)
func (_Axienft *AxienftCallerSession) WhitelistedSpawner(arg0 common.Address) (bool, error) {
	return _Axienft.Contract.WhitelistedSpawner(&_Axienft.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactor) Approve(opts *bind.TransactOpts, _approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "approve", _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_Axienft *AxienftSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.Approve(&_Axienft.TransactOpts, _approved, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _approved, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactorSession) Approve(_approved common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.Approve(&_Axienft.TransactOpts, _approved, _tokenId)
}

// EvolveAxie is a paid mutator transaction binding the contract method 0xae67b4c3.
//
// Solidity: function evolveAxie(uint256 _axieId, uint256 _newGenes) returns()
func (_Axienft *AxienftTransactor) EvolveAxie(opts *bind.TransactOpts, _axieId *big.Int, _newGenes *big.Int) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "evolveAxie", _axieId, _newGenes)
}

// EvolveAxie is a paid mutator transaction binding the contract method 0xae67b4c3.
//
// Solidity: function evolveAxie(uint256 _axieId, uint256 _newGenes) returns()
func (_Axienft *AxienftSession) EvolveAxie(_axieId *big.Int, _newGenes *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.EvolveAxie(&_Axienft.TransactOpts, _axieId, _newGenes)
}

// EvolveAxie is a paid mutator transaction binding the contract method 0xae67b4c3.
//
// Solidity: function evolveAxie(uint256 _axieId, uint256 _newGenes) returns()
func (_Axienft *AxienftTransactorSession) EvolveAxie(_axieId *big.Int, _newGenes *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.EvolveAxie(&_Axienft.TransactOpts, _axieId, _newGenes)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Axienft *AxienftTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Axienft *AxienftSession) Pause() (*types.Transaction, error) {
	return _Axienft.Contract.Pause(&_Axienft.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Axienft *AxienftTransactorSession) Pause() (*types.Transaction, error) {
	return _Axienft.Contract.Pause(&_Axienft.TransactOpts)
}

// RebirthAxie is a paid mutator transaction binding the contract method 0x0d2cc54a.
//
// Solidity: function rebirthAxie(uint256 _axieId, uint256 _genes) returns()
func (_Axienft *AxienftTransactor) RebirthAxie(opts *bind.TransactOpts, _axieId *big.Int, _genes *big.Int) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "rebirthAxie", _axieId, _genes)
}

// RebirthAxie is a paid mutator transaction binding the contract method 0x0d2cc54a.
//
// Solidity: function rebirthAxie(uint256 _axieId, uint256 _genes) returns()
func (_Axienft *AxienftSession) RebirthAxie(_axieId *big.Int, _genes *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.RebirthAxie(&_Axienft.TransactOpts, _axieId, _genes)
}

// RebirthAxie is a paid mutator transaction binding the contract method 0x0d2cc54a.
//
// Solidity: function rebirthAxie(uint256 _axieId, uint256 _genes) returns()
func (_Axienft *AxienftTransactorSession) RebirthAxie(_axieId *big.Int, _genes *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.RebirthAxie(&_Axienft.TransactOpts, _axieId, _genes)
}

// RetireAxie is a paid mutator transaction binding the contract method 0xc9644b77.
//
// Solidity: function retireAxie(uint256 _axieId, bool _rip) returns()
func (_Axienft *AxienftTransactor) RetireAxie(opts *bind.TransactOpts, _axieId *big.Int, _rip bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "retireAxie", _axieId, _rip)
}

// RetireAxie is a paid mutator transaction binding the contract method 0xc9644b77.
//
// Solidity: function retireAxie(uint256 _axieId, bool _rip) returns()
func (_Axienft *AxienftSession) RetireAxie(_axieId *big.Int, _rip bool) (*types.Transaction, error) {
	return _Axienft.Contract.RetireAxie(&_Axienft.TransactOpts, _axieId, _rip)
}

// RetireAxie is a paid mutator transaction binding the contract method 0xc9644b77.
//
// Solidity: function retireAxie(uint256 _axieId, bool _rip) returns()
func (_Axienft *AxienftTransactorSession) RetireAxie(_axieId *big.Int, _rip bool) (*types.Transaction, error) {
	return _Axienft.Contract.RetireAxie(&_Axienft.TransactOpts, _axieId, _rip)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.SafeTransferFrom(&_Axienft.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.SafeTransferFrom(&_Axienft.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) payable returns()
func (_Axienft *AxienftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) payable returns()
func (_Axienft *AxienftSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Axienft.Contract.SafeTransferFrom0(&_Axienft.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) payable returns()
func (_Axienft *AxienftTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Axienft.Contract.SafeTransferFrom0(&_Axienft.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Axienft *AxienftTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Axienft *AxienftSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetApprovalForAll(&_Axienft.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_Axienft *AxienftTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetApprovalForAll(&_Axienft.TransactOpts, _operator, _approved)
}

// SetByeSayer is a paid mutator transaction binding the contract method 0x81b2dad9.
//
// Solidity: function setByeSayer(address _byeSayer, bool _whitelisted) returns()
func (_Axienft *AxienftTransactor) SetByeSayer(opts *bind.TransactOpts, _byeSayer common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setByeSayer", _byeSayer, _whitelisted)
}

// SetByeSayer is a paid mutator transaction binding the contract method 0x81b2dad9.
//
// Solidity: function setByeSayer(address _byeSayer, bool _whitelisted) returns()
func (_Axienft *AxienftSession) SetByeSayer(_byeSayer common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetByeSayer(&_Axienft.TransactOpts, _byeSayer, _whitelisted)
}

// SetByeSayer is a paid mutator transaction binding the contract method 0x81b2dad9.
//
// Solidity: function setByeSayer(address _byeSayer, bool _whitelisted) returns()
func (_Axienft *AxienftTransactorSession) SetByeSayer(_byeSayer common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetByeSayer(&_Axienft.TransactOpts, _byeSayer, _whitelisted)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Axienft *AxienftTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Axienft *AxienftSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCEO(&_Axienft.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_Axienft *AxienftTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCEO(&_Axienft.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Axienft *AxienftTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Axienft *AxienftSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCFO(&_Axienft.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_Axienft *AxienftTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCFO(&_Axienft.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Axienft *AxienftTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Axienft *AxienftSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCOO(&_Axienft.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_Axienft *AxienftTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetCOO(&_Axienft.TransactOpts, _newCOO)
}

// SetGeneManager is a paid mutator transaction binding the contract method 0x927aaa7c.
//
// Solidity: function setGeneManager(address _manager) returns()
func (_Axienft *AxienftTransactor) SetGeneManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setGeneManager", _manager)
}

// SetGeneManager is a paid mutator transaction binding the contract method 0x927aaa7c.
//
// Solidity: function setGeneManager(address _manager) returns()
func (_Axienft *AxienftSession) SetGeneManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetGeneManager(&_Axienft.TransactOpts, _manager)
}

// SetGeneManager is a paid mutator transaction binding the contract method 0x927aaa7c.
//
// Solidity: function setGeneManager(address _manager) returns()
func (_Axienft *AxienftTransactorSession) SetGeneManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetGeneManager(&_Axienft.TransactOpts, _manager)
}

// SetGeneScientist is a paid mutator transaction binding the contract method 0x1f8df2cd.
//
// Solidity: function setGeneScientist(address _geneScientist, bool _whitelisted) returns()
func (_Axienft *AxienftTransactor) SetGeneScientist(opts *bind.TransactOpts, _geneScientist common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setGeneScientist", _geneScientist, _whitelisted)
}

// SetGeneScientist is a paid mutator transaction binding the contract method 0x1f8df2cd.
//
// Solidity: function setGeneScientist(address _geneScientist, bool _whitelisted) returns()
func (_Axienft *AxienftSession) SetGeneScientist(_geneScientist common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetGeneScientist(&_Axienft.TransactOpts, _geneScientist, _whitelisted)
}

// SetGeneScientist is a paid mutator transaction binding the contract method 0x1f8df2cd.
//
// Solidity: function setGeneScientist(address _geneScientist, bool _whitelisted) returns()
func (_Axienft *AxienftTransactorSession) SetGeneScientist(_geneScientist common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetGeneScientist(&_Axienft.TransactOpts, _geneScientist, _whitelisted)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x7419e77a.
//
// Solidity: function setMarketplace(address _marketplace, bool _whitelisted) returns()
func (_Axienft *AxienftTransactor) SetMarketplace(opts *bind.TransactOpts, _marketplace common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setMarketplace", _marketplace, _whitelisted)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x7419e77a.
//
// Solidity: function setMarketplace(address _marketplace, bool _whitelisted) returns()
func (_Axienft *AxienftSession) SetMarketplace(_marketplace common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetMarketplace(&_Axienft.TransactOpts, _marketplace, _whitelisted)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x7419e77a.
//
// Solidity: function setMarketplace(address _marketplace, bool _whitelisted) returns()
func (_Axienft *AxienftTransactorSession) SetMarketplace(_marketplace common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetMarketplace(&_Axienft.TransactOpts, _marketplace, _whitelisted)
}

// SetMarketplaceManager is a paid mutator transaction binding the contract method 0x2ffb054e.
//
// Solidity: function setMarketplaceManager(address _manager) returns()
func (_Axienft *AxienftTransactor) SetMarketplaceManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setMarketplaceManager", _manager)
}

// SetMarketplaceManager is a paid mutator transaction binding the contract method 0x2ffb054e.
//
// Solidity: function setMarketplaceManager(address _manager) returns()
func (_Axienft *AxienftSession) SetMarketplaceManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetMarketplaceManager(&_Axienft.TransactOpts, _manager)
}

// SetMarketplaceManager is a paid mutator transaction binding the contract method 0x2ffb054e.
//
// Solidity: function setMarketplaceManager(address _manager) returns()
func (_Axienft *AxienftTransactorSession) SetMarketplaceManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetMarketplaceManager(&_Axienft.TransactOpts, _manager)
}

// SetRetirementManager is a paid mutator transaction binding the contract method 0x1a68b1a1.
//
// Solidity: function setRetirementManager(address _manager) returns()
func (_Axienft *AxienftTransactor) SetRetirementManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setRetirementManager", _manager)
}

// SetRetirementManager is a paid mutator transaction binding the contract method 0x1a68b1a1.
//
// Solidity: function setRetirementManager(address _manager) returns()
func (_Axienft *AxienftSession) SetRetirementManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetRetirementManager(&_Axienft.TransactOpts, _manager)
}

// SetRetirementManager is a paid mutator transaction binding the contract method 0x1a68b1a1.
//
// Solidity: function setRetirementManager(address _manager) returns()
func (_Axienft *AxienftTransactorSession) SetRetirementManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetRetirementManager(&_Axienft.TransactOpts, _manager)
}

// SetSpawner is a paid mutator transaction binding the contract method 0x9f186edb.
//
// Solidity: function setSpawner(address _spawner, bool _whitelisted) returns()
func (_Axienft *AxienftTransactor) SetSpawner(opts *bind.TransactOpts, _spawner common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setSpawner", _spawner, _whitelisted)
}

// SetSpawner is a paid mutator transaction binding the contract method 0x9f186edb.
//
// Solidity: function setSpawner(address _spawner, bool _whitelisted) returns()
func (_Axienft *AxienftSession) SetSpawner(_spawner common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetSpawner(&_Axienft.TransactOpts, _spawner, _whitelisted)
}

// SetSpawner is a paid mutator transaction binding the contract method 0x9f186edb.
//
// Solidity: function setSpawner(address _spawner, bool _whitelisted) returns()
func (_Axienft *AxienftTransactorSession) SetSpawner(_spawner common.Address, _whitelisted bool) (*types.Transaction, error) {
	return _Axienft.Contract.SetSpawner(&_Axienft.TransactOpts, _spawner, _whitelisted)
}

// SetSpawningManager is a paid mutator transaction binding the contract method 0x8c37e31e.
//
// Solidity: function setSpawningManager(address _manager) returns()
func (_Axienft *AxienftTransactor) SetSpawningManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setSpawningManager", _manager)
}

// SetSpawningManager is a paid mutator transaction binding the contract method 0x8c37e31e.
//
// Solidity: function setSpawningManager(address _manager) returns()
func (_Axienft *AxienftSession) SetSpawningManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetSpawningManager(&_Axienft.TransactOpts, _manager)
}

// SetSpawningManager is a paid mutator transaction binding the contract method 0x8c37e31e.
//
// Solidity: function setSpawningManager(address _manager) returns()
func (_Axienft *AxienftTransactorSession) SetSpawningManager(_manager common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetSpawningManager(&_Axienft.TransactOpts, _manager)
}

// SetTokenURIAffixes is a paid mutator transaction binding the contract method 0xfaf42125.
//
// Solidity: function setTokenURIAffixes(string _prefix, string _suffix) returns()
func (_Axienft *AxienftTransactor) SetTokenURIAffixes(opts *bind.TransactOpts, _prefix string, _suffix string) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setTokenURIAffixes", _prefix, _suffix)
}

// SetTokenURIAffixes is a paid mutator transaction binding the contract method 0xfaf42125.
//
// Solidity: function setTokenURIAffixes(string _prefix, string _suffix) returns()
func (_Axienft *AxienftSession) SetTokenURIAffixes(_prefix string, _suffix string) (*types.Transaction, error) {
	return _Axienft.Contract.SetTokenURIAffixes(&_Axienft.TransactOpts, _prefix, _suffix)
}

// SetTokenURIAffixes is a paid mutator transaction binding the contract method 0xfaf42125.
//
// Solidity: function setTokenURIAffixes(string _prefix, string _suffix) returns()
func (_Axienft *AxienftTransactorSession) SetTokenURIAffixes(_prefix string, _suffix string) (*types.Transaction, error) {
	return _Axienft.Contract.SetTokenURIAffixes(&_Axienft.TransactOpts, _prefix, _suffix)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _newSetter) returns()
func (_Axienft *AxienftTransactor) SetWhitelistSetter(opts *bind.TransactOpts, _newSetter common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "setWhitelistSetter", _newSetter)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _newSetter) returns()
func (_Axienft *AxienftSession) SetWhitelistSetter(_newSetter common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetWhitelistSetter(&_Axienft.TransactOpts, _newSetter)
}

// SetWhitelistSetter is a paid mutator transaction binding the contract method 0x547a5eee.
//
// Solidity: function setWhitelistSetter(address _newSetter) returns()
func (_Axienft *AxienftTransactorSession) SetWhitelistSetter(_newSetter common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SetWhitelistSetter(&_Axienft.TransactOpts, _newSetter)
}

// SpawnAxie is a paid mutator transaction binding the contract method 0x5dcc6dbc.
//
// Solidity: function spawnAxie(uint256 _genes, address _owner) returns(uint256)
func (_Axienft *AxienftTransactor) SpawnAxie(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "spawnAxie", _genes, _owner)
}

// SpawnAxie is a paid mutator transaction binding the contract method 0x5dcc6dbc.
//
// Solidity: function spawnAxie(uint256 _genes, address _owner) returns(uint256)
func (_Axienft *AxienftSession) SpawnAxie(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SpawnAxie(&_Axienft.TransactOpts, _genes, _owner)
}

// SpawnAxie is a paid mutator transaction binding the contract method 0x5dcc6dbc.
//
// Solidity: function spawnAxie(uint256 _genes, address _owner) returns(uint256)
func (_Axienft *AxienftTransactorSession) SpawnAxie(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _Axienft.Contract.SpawnAxie(&_Axienft.TransactOpts, _genes, _owner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.TransferFrom(&_Axienft.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) payable returns()
func (_Axienft *AxienftTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Axienft.Contract.TransferFrom(&_Axienft.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Axienft *AxienftTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Axienft *AxienftSession) Unpause() (*types.Transaction, error) {
	return _Axienft.Contract.Unpause(&_Axienft.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Axienft *AxienftTransactorSession) Unpause() (*types.Transaction, error) {
	return _Axienft.Contract.Unpause(&_Axienft.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Axienft *AxienftTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Axienft.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Axienft *AxienftSession) WithdrawBalance() (*types.Transaction, error) {
	return _Axienft.Contract.WithdrawBalance(&_Axienft.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_Axienft *AxienftTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _Axienft.Contract.WithdrawBalance(&_Axienft.TransactOpts)
}

// AxienftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Axienft contract.
type AxienftApprovalIterator struct {
	Event *AxienftApproval // Event containing the contract specifics and raw log

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
func (it *AxienftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftApproval)
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
		it.Event = new(AxienftApproval)
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
func (it *AxienftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftApproval represents a Approval event raised by the Axienft contract.
type AxienftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Axienft *AxienftFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*AxienftApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &AxienftApprovalIterator{contract: _Axienft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Axienft *AxienftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AxienftApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftApproval)
				if err := _Axienft.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_Axienft *AxienftFilterer) ParseApproval(log types.Log) (*AxienftApproval, error) {
	event := new(AxienftApproval)
	if err := _Axienft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Axienft contract.
type AxienftApprovalForAllIterator struct {
	Event *AxienftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AxienftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftApprovalForAll)
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
		it.Event = new(AxienftApprovalForAll)
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
func (it *AxienftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftApprovalForAll represents a ApprovalForAll event raised by the Axienft contract.
type AxienftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Axienft *AxienftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*AxienftApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &AxienftApprovalForAllIterator{contract: _Axienft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_Axienft *AxienftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AxienftApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftApprovalForAll)
				if err := _Axienft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Axienft *AxienftFilterer) ParseApprovalForAll(log types.Log) (*AxienftApprovalForAll, error) {
	event := new(AxienftApprovalForAll)
	if err := _Axienft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftAxieEvolvedIterator is returned from FilterAxieEvolved and is used to iterate over the raw logs and unpacked data for AxieEvolved events raised by the Axienft contract.
type AxienftAxieEvolvedIterator struct {
	Event *AxienftAxieEvolved // Event containing the contract specifics and raw log

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
func (it *AxienftAxieEvolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftAxieEvolved)
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
		it.Event = new(AxienftAxieEvolved)
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
func (it *AxienftAxieEvolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftAxieEvolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftAxieEvolved represents a AxieEvolved event raised by the Axienft contract.
type AxienftAxieEvolved struct {
	AxieId   *big.Int
	OldGenes *big.Int
	NewGenes *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAxieEvolved is a free log retrieval operation binding the contract event 0x6715bb2e56d8b9cd079426816318abbb6ee679a7e8e64602b38bf5c074b58e55.
//
// Solidity: event AxieEvolved(uint256 indexed _axieId, uint256 _oldGenes, uint256 _newGenes)
func (_Axienft *AxienftFilterer) FilterAxieEvolved(opts *bind.FilterOpts, _axieId []*big.Int) (*AxienftAxieEvolvedIterator, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "AxieEvolved", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return &AxienftAxieEvolvedIterator{contract: _Axienft.contract, event: "AxieEvolved", logs: logs, sub: sub}, nil
}

// WatchAxieEvolved is a free log subscription operation binding the contract event 0x6715bb2e56d8b9cd079426816318abbb6ee679a7e8e64602b38bf5c074b58e55.
//
// Solidity: event AxieEvolved(uint256 indexed _axieId, uint256 _oldGenes, uint256 _newGenes)
func (_Axienft *AxienftFilterer) WatchAxieEvolved(opts *bind.WatchOpts, sink chan<- *AxienftAxieEvolved, _axieId []*big.Int) (event.Subscription, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "AxieEvolved", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftAxieEvolved)
				if err := _Axienft.contract.UnpackLog(event, "AxieEvolved", log); err != nil {
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

// ParseAxieEvolved is a log parse operation binding the contract event 0x6715bb2e56d8b9cd079426816318abbb6ee679a7e8e64602b38bf5c074b58e55.
//
// Solidity: event AxieEvolved(uint256 indexed _axieId, uint256 _oldGenes, uint256 _newGenes)
func (_Axienft *AxienftFilterer) ParseAxieEvolved(log types.Log) (*AxienftAxieEvolved, error) {
	event := new(AxienftAxieEvolved)
	if err := _Axienft.contract.UnpackLog(event, "AxieEvolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftAxieRebirthedIterator is returned from FilterAxieRebirthed and is used to iterate over the raw logs and unpacked data for AxieRebirthed events raised by the Axienft contract.
type AxienftAxieRebirthedIterator struct {
	Event *AxienftAxieRebirthed // Event containing the contract specifics and raw log

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
func (it *AxienftAxieRebirthedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftAxieRebirthed)
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
		it.Event = new(AxienftAxieRebirthed)
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
func (it *AxienftAxieRebirthedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftAxieRebirthedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftAxieRebirthed represents a AxieRebirthed event raised by the Axienft contract.
type AxienftAxieRebirthed struct {
	AxieId *big.Int
	Genes  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAxieRebirthed is a free log retrieval operation binding the contract event 0x90f712fe1864fe3fbc2bd97b660fbfec7c30d013d3961ac72ea12cffc09e0663.
//
// Solidity: event AxieRebirthed(uint256 indexed _axieId, uint256 _genes)
func (_Axienft *AxienftFilterer) FilterAxieRebirthed(opts *bind.FilterOpts, _axieId []*big.Int) (*AxienftAxieRebirthedIterator, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "AxieRebirthed", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return &AxienftAxieRebirthedIterator{contract: _Axienft.contract, event: "AxieRebirthed", logs: logs, sub: sub}, nil
}

// WatchAxieRebirthed is a free log subscription operation binding the contract event 0x90f712fe1864fe3fbc2bd97b660fbfec7c30d013d3961ac72ea12cffc09e0663.
//
// Solidity: event AxieRebirthed(uint256 indexed _axieId, uint256 _genes)
func (_Axienft *AxienftFilterer) WatchAxieRebirthed(opts *bind.WatchOpts, sink chan<- *AxienftAxieRebirthed, _axieId []*big.Int) (event.Subscription, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "AxieRebirthed", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftAxieRebirthed)
				if err := _Axienft.contract.UnpackLog(event, "AxieRebirthed", log); err != nil {
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

// ParseAxieRebirthed is a log parse operation binding the contract event 0x90f712fe1864fe3fbc2bd97b660fbfec7c30d013d3961ac72ea12cffc09e0663.
//
// Solidity: event AxieRebirthed(uint256 indexed _axieId, uint256 _genes)
func (_Axienft *AxienftFilterer) ParseAxieRebirthed(log types.Log) (*AxienftAxieRebirthed, error) {
	event := new(AxienftAxieRebirthed)
	if err := _Axienft.contract.UnpackLog(event, "AxieRebirthed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftAxieRetiredIterator is returned from FilterAxieRetired and is used to iterate over the raw logs and unpacked data for AxieRetired events raised by the Axienft contract.
type AxienftAxieRetiredIterator struct {
	Event *AxienftAxieRetired // Event containing the contract specifics and raw log

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
func (it *AxienftAxieRetiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftAxieRetired)
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
		it.Event = new(AxienftAxieRetired)
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
func (it *AxienftAxieRetiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftAxieRetiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftAxieRetired represents a AxieRetired event raised by the Axienft contract.
type AxienftAxieRetired struct {
	AxieId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAxieRetired is a free log retrieval operation binding the contract event 0xc7e9cc512eb27debd93c5ff5830060459d7e6777ebcb540f3c9e2e520eb2502b.
//
// Solidity: event AxieRetired(uint256 indexed _axieId)
func (_Axienft *AxienftFilterer) FilterAxieRetired(opts *bind.FilterOpts, _axieId []*big.Int) (*AxienftAxieRetiredIterator, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "AxieRetired", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return &AxienftAxieRetiredIterator{contract: _Axienft.contract, event: "AxieRetired", logs: logs, sub: sub}, nil
}

// WatchAxieRetired is a free log subscription operation binding the contract event 0xc7e9cc512eb27debd93c5ff5830060459d7e6777ebcb540f3c9e2e520eb2502b.
//
// Solidity: event AxieRetired(uint256 indexed _axieId)
func (_Axienft *AxienftFilterer) WatchAxieRetired(opts *bind.WatchOpts, sink chan<- *AxienftAxieRetired, _axieId []*big.Int) (event.Subscription, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "AxieRetired", _axieIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftAxieRetired)
				if err := _Axienft.contract.UnpackLog(event, "AxieRetired", log); err != nil {
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

// ParseAxieRetired is a log parse operation binding the contract event 0xc7e9cc512eb27debd93c5ff5830060459d7e6777ebcb540f3c9e2e520eb2502b.
//
// Solidity: event AxieRetired(uint256 indexed _axieId)
func (_Axienft *AxienftFilterer) ParseAxieRetired(log types.Log) (*AxienftAxieRetired, error) {
	event := new(AxienftAxieRetired)
	if err := _Axienft.contract.UnpackLog(event, "AxieRetired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftAxieSpawnedIterator is returned from FilterAxieSpawned and is used to iterate over the raw logs and unpacked data for AxieSpawned events raised by the Axienft contract.
type AxienftAxieSpawnedIterator struct {
	Event *AxienftAxieSpawned // Event containing the contract specifics and raw log

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
func (it *AxienftAxieSpawnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftAxieSpawned)
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
		it.Event = new(AxienftAxieSpawned)
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
func (it *AxienftAxieSpawnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftAxieSpawnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftAxieSpawned represents a AxieSpawned event raised by the Axienft contract.
type AxienftAxieSpawned struct {
	AxieId *big.Int
	Owner  common.Address
	Genes  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAxieSpawned is a free log retrieval operation binding the contract event 0x0aa75ef3ef0b2c4f092828edf4e4ded5a223d08350535710b957679883ff3a3c.
//
// Solidity: event AxieSpawned(uint256 indexed _axieId, address indexed _owner, uint256 _genes)
func (_Axienft *AxienftFilterer) FilterAxieSpawned(opts *bind.FilterOpts, _axieId []*big.Int, _owner []common.Address) (*AxienftAxieSpawnedIterator, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "AxieSpawned", _axieIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &AxienftAxieSpawnedIterator{contract: _Axienft.contract, event: "AxieSpawned", logs: logs, sub: sub}, nil
}

// WatchAxieSpawned is a free log subscription operation binding the contract event 0x0aa75ef3ef0b2c4f092828edf4e4ded5a223d08350535710b957679883ff3a3c.
//
// Solidity: event AxieSpawned(uint256 indexed _axieId, address indexed _owner, uint256 _genes)
func (_Axienft *AxienftFilterer) WatchAxieSpawned(opts *bind.WatchOpts, sink chan<- *AxienftAxieSpawned, _axieId []*big.Int, _owner []common.Address) (event.Subscription, error) {

	var _axieIdRule []interface{}
	for _, _axieIdItem := range _axieId {
		_axieIdRule = append(_axieIdRule, _axieIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "AxieSpawned", _axieIdRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftAxieSpawned)
				if err := _Axienft.contract.UnpackLog(event, "AxieSpawned", log); err != nil {
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

// ParseAxieSpawned is a log parse operation binding the contract event 0x0aa75ef3ef0b2c4f092828edf4e4ded5a223d08350535710b957679883ff3a3c.
//
// Solidity: event AxieSpawned(uint256 indexed _axieId, address indexed _owner, uint256 _genes)
func (_Axienft *AxienftFilterer) ParseAxieSpawned(log types.Log) (*AxienftAxieSpawned, error) {
	event := new(AxienftAxieSpawned)
	if err := _Axienft.contract.UnpackLog(event, "AxieSpawned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AxienftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Axienft contract.
type AxienftTransferIterator struct {
	Event *AxienftTransfer // Event containing the contract specifics and raw log

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
func (it *AxienftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AxienftTransfer)
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
		it.Event = new(AxienftTransfer)
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
func (it *AxienftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AxienftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AxienftTransfer represents a Transfer event raised by the Axienft contract.
type AxienftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Axienft *AxienftFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*AxienftTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Axienft.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &AxienftTransferIterator{contract: _Axienft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Axienft *AxienftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AxienftTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Axienft.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AxienftTransfer)
				if err := _Axienft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_Axienft *AxienftFilterer) ParseTransfer(log types.Log) (*AxienftTransfer, error) {
	event := new(AxienftTransfer)
	if err := _Axienft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
