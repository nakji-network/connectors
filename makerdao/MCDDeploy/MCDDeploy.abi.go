// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MCDDeploy

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

// MCDDEPLOYABI is the input ABI used to generate the binding from.
const MCDDEPLOYABI = "[{\"inputs\":[{\"internalType\":\"contractVatFab\",\"name\":\"vatFab_\",\"type\":\"address\"},{\"internalType\":\"contractJugFab\",\"name\":\"jugFab_\",\"type\":\"address\"},{\"internalType\":\"contractVowFab\",\"name\":\"vowFab_\",\"type\":\"address\"},{\"internalType\":\"contractCatFab\",\"name\":\"catFab_\",\"type\":\"address\"},{\"internalType\":\"contractDaiFab\",\"name\":\"daiFab_\",\"type\":\"address\"},{\"internalType\":\"contractDaiJoinFab\",\"name\":\"daiJoinFab_\",\"type\":\"address\"},{\"internalType\":\"contractFlapFab\",\"name\":\"flapFab_\",\"type\":\"address\"},{\"internalType\":\"contractFlopFab\",\"name\":\"flopFab_\",\"type\":\"address\"},{\"internalType\":\"contractFlipFab\",\"name\":\"flipFab_\",\"type\":\"address\"},{\"internalType\":\"contractSpotFab\",\"name\":\"spotFab_\",\"type\":\"address\"},{\"internalType\":\"contractPotFab\",\"name\":\"potFab_\",\"type\":\"address\"},{\"internalType\":\"contractEndFab\",\"name\":\"endFab_\",\"type\":\"address\"},{\"internalType\":\"contractESMFab\",\"name\":\"esmFab_\",\"type\":\"address\"},{\"internalType\":\"contractPauseFab\",\"name\":\"pauseFab_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cat\",\"outputs\":[{\"internalType\":\"contractCat\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"catFab\",\"outputs\":[{\"internalType\":\"contractCatFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dai\",\"outputs\":[{\"internalType\":\"contractDai\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiFab\",\"outputs\":[{\"internalType\":\"contractDaiFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoin\",\"outputs\":[{\"internalType\":\"contractDaiJoin\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"daiJoinFab\",\"outputs\":[{\"internalType\":\"contractDaiJoinFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"deployAuctions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pip\",\"type\":\"address\"}],\"name\":\"deployCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"deployDai\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deployLiquidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"contractDSAuthority\",\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"deployPause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pit\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"}],\"name\":\"deployShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deployTaxation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deployVat\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"internalType\":\"contractEnd\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endFab\",\"outputs\":[{\"internalType\":\"contractEndFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"esm\",\"outputs\":[{\"internalType\":\"contractESM\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"esmFab\",\"outputs\":[{\"internalType\":\"contractESMFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flap\",\"outputs\":[{\"internalType\":\"contractFlapper\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flapFab\",\"outputs\":[{\"internalType\":\"contractFlapFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flipFab\",\"outputs\":[{\"internalType\":\"contractFlipFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flop\",\"outputs\":[{\"internalType\":\"contractFlopper\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"flopFab\",\"outputs\":[{\"internalType\":\"contractFlopFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"contractFlipper\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"join\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"jug\",\"outputs\":[{\"internalType\":\"contractJug\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"jugFab\",\"outputs\":[{\"internalType\":\"contractJugFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"contractDSPause\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseFab\",\"outputs\":[{\"internalType\":\"contractPauseFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pot\",\"outputs\":[{\"internalType\":\"contractPot\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"potFab\",\"outputs\":[{\"internalType\":\"contractPotFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"releaseAuth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"releaseAuthFlip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractDSAuthority\",\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spotFab\",\"outputs\":[{\"internalType\":\"contractSpotFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spotter\",\"outputs\":[{\"internalType\":\"contractSpotter\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"step\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVat\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vatFab\",\"outputs\":[{\"internalType\":\"contractVatFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"contractVow\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vowFab\",\"outputs\":[{\"internalType\":\"contractVowFab\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MCDDEPLOY is an auto generated Go binding around an Ethereum contract.
type MCDDEPLOY struct {
	MCDDEPLOYCaller     // Read-only binding to the contract
	MCDDEPLOYTransactor // Write-only binding to the contract
	MCDDEPLOYFilterer   // Log filterer for contract events
}

// MCDDEPLOYCaller is an auto generated read-only Go binding around an Ethereum contract.
type MCDDEPLOYCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDEPLOYTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MCDDEPLOYTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDEPLOYFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MCDDEPLOYFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MCDDEPLOYSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MCDDEPLOYSession struct {
	Contract     *MCDDEPLOY        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MCDDEPLOYCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MCDDEPLOYCallerSession struct {
	Contract *MCDDEPLOYCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MCDDEPLOYTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MCDDEPLOYTransactorSession struct {
	Contract     *MCDDEPLOYTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MCDDEPLOYRaw is an auto generated low-level Go binding around an Ethereum contract.
type MCDDEPLOYRaw struct {
	Contract *MCDDEPLOY // Generic contract binding to access the raw methods on
}

// MCDDEPLOYCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MCDDEPLOYCallerRaw struct {
	Contract *MCDDEPLOYCaller // Generic read-only contract binding to access the raw methods on
}

// MCDDEPLOYTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MCDDEPLOYTransactorRaw struct {
	Contract *MCDDEPLOYTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMCDDEPLOY creates a new instance of MCDDEPLOY, bound to a specific deployed contract.
func NewMCDDEPLOY(address common.Address, backend bind.ContractBackend) (*MCDDEPLOY, error) {
	contract, err := bindMCDDEPLOY(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOY{MCDDEPLOYCaller: MCDDEPLOYCaller{contract: contract}, MCDDEPLOYTransactor: MCDDEPLOYTransactor{contract: contract}, MCDDEPLOYFilterer: MCDDEPLOYFilterer{contract: contract}}, nil
}

// NewMCDDEPLOYCaller creates a new read-only instance of MCDDEPLOY, bound to a specific deployed contract.
func NewMCDDEPLOYCaller(address common.Address, caller bind.ContractCaller) (*MCDDEPLOYCaller, error) {
	contract, err := bindMCDDEPLOY(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOYCaller{contract: contract}, nil
}

// NewMCDDEPLOYTransactor creates a new write-only instance of MCDDEPLOY, bound to a specific deployed contract.
func NewMCDDEPLOYTransactor(address common.Address, transactor bind.ContractTransactor) (*MCDDEPLOYTransactor, error) {
	contract, err := bindMCDDEPLOY(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOYTransactor{contract: contract}, nil
}

// NewMCDDEPLOYFilterer creates a new log filterer instance of MCDDEPLOY, bound to a specific deployed contract.
func NewMCDDEPLOYFilterer(address common.Address, filterer bind.ContractFilterer) (*MCDDEPLOYFilterer, error) {
	contract, err := bindMCDDEPLOY(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOYFilterer{contract: contract}, nil
}

// bindMCDDEPLOY binds a generic wrapper to an already deployed contract.
func bindMCDDEPLOY(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MCDDEPLOYABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDDEPLOY *MCDDEPLOYRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDDEPLOY.Contract.MCDDEPLOYCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDDEPLOY *MCDDEPLOYRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.MCDDEPLOYTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDDEPLOY *MCDDEPLOYRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.MCDDEPLOYTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MCDDEPLOY *MCDDEPLOYCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MCDDEPLOY.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MCDDEPLOY *MCDDEPLOYTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MCDDEPLOY *MCDDEPLOYTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Authority() (common.Address, error) {
	return _MCDDEPLOY.Contract.Authority(&_MCDDEPLOY.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Authority() (common.Address, error) {
	return _MCDDEPLOY.Contract.Authority(&_MCDDEPLOY.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Cat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "cat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Cat() (common.Address, error) {
	return _MCDDEPLOY.Contract.Cat(&_MCDDEPLOY.CallOpts)
}

// Cat is a free data retrieval call binding the contract method 0xe4881813.
//
// Solidity: function cat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Cat() (common.Address, error) {
	return _MCDDEPLOY.Contract.Cat(&_MCDDEPLOY.CallOpts)
}

// CatFab is a free data retrieval call binding the contract method 0x199bead4.
//
// Solidity: function catFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) CatFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "catFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CatFab is a free data retrieval call binding the contract method 0x199bead4.
//
// Solidity: function catFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) CatFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.CatFab(&_MCDDEPLOY.CallOpts)
}

// CatFab is a free data retrieval call binding the contract method 0x199bead4.
//
// Solidity: function catFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) CatFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.CatFab(&_MCDDEPLOY.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Dai(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "dai")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Dai() (common.Address, error) {
	return _MCDDEPLOY.Contract.Dai(&_MCDDEPLOY.CallOpts)
}

// Dai is a free data retrieval call binding the contract method 0xf4b9fa75.
//
// Solidity: function dai() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Dai() (common.Address, error) {
	return _MCDDEPLOY.Contract.Dai(&_MCDDEPLOY.CallOpts)
}

// DaiFab is a free data retrieval call binding the contract method 0x39ca1def.
//
// Solidity: function daiFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) DaiFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "daiFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiFab is a free data retrieval call binding the contract method 0x39ca1def.
//
// Solidity: function daiFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) DaiFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiFab(&_MCDDEPLOY.CallOpts)
}

// DaiFab is a free data retrieval call binding the contract method 0x39ca1def.
//
// Solidity: function daiFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) DaiFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiFab(&_MCDDEPLOY.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) DaiJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "daiJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) DaiJoin() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiJoin(&_MCDDEPLOY.CallOpts)
}

// DaiJoin is a free data retrieval call binding the contract method 0xc11645bc.
//
// Solidity: function daiJoin() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) DaiJoin() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiJoin(&_MCDDEPLOY.CallOpts)
}

// DaiJoinFab is a free data retrieval call binding the contract method 0x29ced7a0.
//
// Solidity: function daiJoinFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) DaiJoinFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "daiJoinFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaiJoinFab is a free data retrieval call binding the contract method 0x29ced7a0.
//
// Solidity: function daiJoinFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) DaiJoinFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiJoinFab(&_MCDDEPLOY.CallOpts)
}

// DaiJoinFab is a free data retrieval call binding the contract method 0x29ced7a0.
//
// Solidity: function daiJoinFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) DaiJoinFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.DaiJoinFab(&_MCDDEPLOY.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) End(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "end")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) End() (common.Address, error) {
	return _MCDDEPLOY.Contract.End(&_MCDDEPLOY.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) End() (common.Address, error) {
	return _MCDDEPLOY.Contract.End(&_MCDDEPLOY.CallOpts)
}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) EndFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "endFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) EndFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.EndFab(&_MCDDEPLOY.CallOpts)
}

// EndFab is a free data retrieval call binding the contract method 0xef0add5b.
//
// Solidity: function endFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) EndFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.EndFab(&_MCDDEPLOY.CallOpts)
}

// Esm is a free data retrieval call binding the contract method 0xaf41b248.
//
// Solidity: function esm() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Esm(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "esm")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Esm is a free data retrieval call binding the contract method 0xaf41b248.
//
// Solidity: function esm() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Esm() (common.Address, error) {
	return _MCDDEPLOY.Contract.Esm(&_MCDDEPLOY.CallOpts)
}

// Esm is a free data retrieval call binding the contract method 0xaf41b248.
//
// Solidity: function esm() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Esm() (common.Address, error) {
	return _MCDDEPLOY.Contract.Esm(&_MCDDEPLOY.CallOpts)
}

// EsmFab is a free data retrieval call binding the contract method 0x64336cd0.
//
// Solidity: function esmFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) EsmFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "esmFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsmFab is a free data retrieval call binding the contract method 0x64336cd0.
//
// Solidity: function esmFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) EsmFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.EsmFab(&_MCDDEPLOY.CallOpts)
}

// EsmFab is a free data retrieval call binding the contract method 0x64336cd0.
//
// Solidity: function esmFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) EsmFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.EsmFab(&_MCDDEPLOY.CallOpts)
}

// Flap is a free data retrieval call binding the contract method 0x0e01198b.
//
// Solidity: function flap() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Flap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "flap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flap is a free data retrieval call binding the contract method 0x0e01198b.
//
// Solidity: function flap() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Flap() (common.Address, error) {
	return _MCDDEPLOY.Contract.Flap(&_MCDDEPLOY.CallOpts)
}

// Flap is a free data retrieval call binding the contract method 0x0e01198b.
//
// Solidity: function flap() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Flap() (common.Address, error) {
	return _MCDDEPLOY.Contract.Flap(&_MCDDEPLOY.CallOpts)
}

// FlapFab is a free data retrieval call binding the contract method 0x39c98420.
//
// Solidity: function flapFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) FlapFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "flapFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlapFab is a free data retrieval call binding the contract method 0x39c98420.
//
// Solidity: function flapFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) FlapFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlapFab(&_MCDDEPLOY.CallOpts)
}

// FlapFab is a free data retrieval call binding the contract method 0x39c98420.
//
// Solidity: function flapFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) FlapFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlapFab(&_MCDDEPLOY.CallOpts)
}

// FlipFab is a free data retrieval call binding the contract method 0x7c91e177.
//
// Solidity: function flipFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) FlipFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "flipFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlipFab is a free data retrieval call binding the contract method 0x7c91e177.
//
// Solidity: function flipFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) FlipFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlipFab(&_MCDDEPLOY.CallOpts)
}

// FlipFab is a free data retrieval call binding the contract method 0x7c91e177.
//
// Solidity: function flipFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) FlipFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlipFab(&_MCDDEPLOY.CallOpts)
}

// Flop is a free data retrieval call binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Flop(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "flop")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flop is a free data retrieval call binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Flop() (common.Address, error) {
	return _MCDDEPLOY.Contract.Flop(&_MCDDEPLOY.CallOpts)
}

// Flop is a free data retrieval call binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Flop() (common.Address, error) {
	return _MCDDEPLOY.Contract.Flop(&_MCDDEPLOY.CallOpts)
}

// FlopFab is a free data retrieval call binding the contract method 0x783e10fc.
//
// Solidity: function flopFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) FlopFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "flopFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlopFab is a free data retrieval call binding the contract method 0x783e10fc.
//
// Solidity: function flopFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) FlopFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlopFab(&_MCDDEPLOY.CallOpts)
}

// FlopFab is a free data retrieval call binding the contract method 0x783e10fc.
//
// Solidity: function flopFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) FlopFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.FlopFab(&_MCDDEPLOY.CallOpts)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, address join)
func (_MCDDEPLOY *MCDDEPLOYCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Flip common.Address
	Join common.Address
}, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Flip common.Address
		Join common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Flip = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Join = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, address join)
func (_MCDDEPLOY *MCDDEPLOYSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Join common.Address
}, error) {
	return _MCDDEPLOY.Contract.Ilks(&_MCDDEPLOY.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(address flip, address join)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Ilks(arg0 [32]byte) (struct {
	Flip common.Address
	Join common.Address
}, error) {
	return _MCDDEPLOY.Contract.Ilks(&_MCDDEPLOY.CallOpts, arg0)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Jug(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "jug")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Jug() (common.Address, error) {
	return _MCDDEPLOY.Contract.Jug(&_MCDDEPLOY.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Jug() (common.Address, error) {
	return _MCDDEPLOY.Contract.Jug(&_MCDDEPLOY.CallOpts)
}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) JugFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "jugFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) JugFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.JugFab(&_MCDDEPLOY.CallOpts)
}

// JugFab is a free data retrieval call binding the contract method 0x5f35a49d.
//
// Solidity: function jugFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) JugFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.JugFab(&_MCDDEPLOY.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Owner() (common.Address, error) {
	return _MCDDEPLOY.Contract.Owner(&_MCDDEPLOY.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Owner() (common.Address, error) {
	return _MCDDEPLOY.Contract.Owner(&_MCDDEPLOY.CallOpts)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Pause(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "pause")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Pause() (common.Address, error) {
	return _MCDDEPLOY.Contract.Pause(&_MCDDEPLOY.CallOpts)
}

// Pause is a free data retrieval call binding the contract method 0x8456cb59.
//
// Solidity: function pause() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Pause() (common.Address, error) {
	return _MCDDEPLOY.Contract.Pause(&_MCDDEPLOY.CallOpts)
}

// PauseFab is a free data retrieval call binding the contract method 0xc8e334a9.
//
// Solidity: function pauseFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) PauseFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "pauseFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseFab is a free data retrieval call binding the contract method 0xc8e334a9.
//
// Solidity: function pauseFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) PauseFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.PauseFab(&_MCDDEPLOY.CallOpts)
}

// PauseFab is a free data retrieval call binding the contract method 0xc8e334a9.
//
// Solidity: function pauseFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) PauseFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.PauseFab(&_MCDDEPLOY.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Pot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "pot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Pot() (common.Address, error) {
	return _MCDDEPLOY.Contract.Pot(&_MCDDEPLOY.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4ba2363a.
//
// Solidity: function pot() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Pot() (common.Address, error) {
	return _MCDDEPLOY.Contract.Pot(&_MCDDEPLOY.CallOpts)
}

// PotFab is a free data retrieval call binding the contract method 0xeae131a9.
//
// Solidity: function potFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) PotFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "potFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PotFab is a free data retrieval call binding the contract method 0xeae131a9.
//
// Solidity: function potFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) PotFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.PotFab(&_MCDDEPLOY.CallOpts)
}

// PotFab is a free data retrieval call binding the contract method 0xeae131a9.
//
// Solidity: function potFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) PotFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.PotFab(&_MCDDEPLOY.CallOpts)
}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) SpotFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "spotFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) SpotFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.SpotFab(&_MCDDEPLOY.CallOpts)
}

// SpotFab is a free data retrieval call binding the contract method 0x92ed5460.
//
// Solidity: function spotFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) SpotFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.SpotFab(&_MCDDEPLOY.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Spotter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "spotter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Spotter() (common.Address, error) {
	return _MCDDEPLOY.Contract.Spotter(&_MCDDEPLOY.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Spotter() (common.Address, error) {
	return _MCDDEPLOY.Contract.Spotter(&_MCDDEPLOY.CallOpts)
}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_MCDDEPLOY *MCDDEPLOYCaller) Step(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "step")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_MCDDEPLOY *MCDDEPLOYSession) Step() (uint8, error) {
	return _MCDDEPLOY.Contract.Step(&_MCDDEPLOY.CallOpts)
}

// Step is a free data retrieval call binding the contract method 0xe25fe175.
//
// Solidity: function step() view returns(uint8)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Step() (uint8, error) {
	return _MCDDEPLOY.Contract.Step(&_MCDDEPLOY.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Vat() (common.Address, error) {
	return _MCDDEPLOY.Contract.Vat(&_MCDDEPLOY.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Vat() (common.Address, error) {
	return _MCDDEPLOY.Contract.Vat(&_MCDDEPLOY.CallOpts)
}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) VatFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "vatFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) VatFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.VatFab(&_MCDDEPLOY.CallOpts)
}

// VatFab is a free data retrieval call binding the contract method 0x4a847390.
//
// Solidity: function vatFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) VatFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.VatFab(&_MCDDEPLOY.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) Vow() (common.Address, error) {
	return _MCDDEPLOY.Contract.Vow(&_MCDDEPLOY.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) Vow() (common.Address, error) {
	return _MCDDEPLOY.Contract.Vow(&_MCDDEPLOY.CallOpts)
}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCaller) VowFab(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MCDDEPLOY.contract.Call(opts, &out, "vowFab")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYSession) VowFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.VowFab(&_MCDDEPLOY.CallOpts)
}

// VowFab is a free data retrieval call binding the contract method 0xb94fdb8e.
//
// Solidity: function vowFab() view returns(address)
func (_MCDDEPLOY *MCDDEPLOYCallerSession) VowFab() (common.Address, error) {
	return _MCDDEPLOY.Contract.VowFab(&_MCDDEPLOY.CallOpts)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployAuctions(opts *bind.TransactOpts, gov common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployAuctions", gov)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployAuctions(gov common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployAuctions(&_MCDDEPLOY.TransactOpts, gov)
}

// DeployAuctions is a paid mutator transaction binding the contract method 0xcc3f5ec1.
//
// Solidity: function deployAuctions(address gov) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployAuctions(gov common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployAuctions(&_MCDDEPLOY.TransactOpts, gov)
}

// DeployCollateral is a paid mutator transaction binding the contract method 0x7d3fdfa1.
//
// Solidity: function deployCollateral(bytes32 ilk, address join, address pip) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployCollateral(opts *bind.TransactOpts, ilk [32]byte, join common.Address, pip common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployCollateral", ilk, join, pip)
}

// DeployCollateral is a paid mutator transaction binding the contract method 0x7d3fdfa1.
//
// Solidity: function deployCollateral(bytes32 ilk, address join, address pip) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployCollateral(ilk [32]byte, join common.Address, pip common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployCollateral(&_MCDDEPLOY.TransactOpts, ilk, join, pip)
}

// DeployCollateral is a paid mutator transaction binding the contract method 0x7d3fdfa1.
//
// Solidity: function deployCollateral(bytes32 ilk, address join, address pip) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployCollateral(ilk [32]byte, join common.Address, pip common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployCollateral(&_MCDDEPLOY.TransactOpts, ilk, join, pip)
}

// DeployDai is a paid mutator transaction binding the contract method 0x64b07f21.
//
// Solidity: function deployDai(uint256 chainId) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployDai(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployDai", chainId)
}

// DeployDai is a paid mutator transaction binding the contract method 0x64b07f21.
//
// Solidity: function deployDai(uint256 chainId) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployDai(chainId *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployDai(&_MCDDEPLOY.TransactOpts, chainId)
}

// DeployDai is a paid mutator transaction binding the contract method 0x64b07f21.
//
// Solidity: function deployDai(uint256 chainId) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployDai(chainId *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployDai(&_MCDDEPLOY.TransactOpts, chainId)
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployLiquidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployLiquidator")
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployLiquidator() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployLiquidator(&_MCDDEPLOY.TransactOpts)
}

// DeployLiquidator is a paid mutator transaction binding the contract method 0xdcc91c9b.
//
// Solidity: function deployLiquidator() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployLiquidator() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployLiquidator(&_MCDDEPLOY.TransactOpts)
}

// DeployPause is a paid mutator transaction binding the contract method 0xce96193b.
//
// Solidity: function deployPause(uint256 delay, address authority) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployPause(opts *bind.TransactOpts, delay *big.Int, authority common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployPause", delay, authority)
}

// DeployPause is a paid mutator transaction binding the contract method 0xce96193b.
//
// Solidity: function deployPause(uint256 delay, address authority) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployPause(delay *big.Int, authority common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployPause(&_MCDDEPLOY.TransactOpts, delay, authority)
}

// DeployPause is a paid mutator transaction binding the contract method 0xce96193b.
//
// Solidity: function deployPause(uint256 delay, address authority) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployPause(delay *big.Int, authority common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployPause(&_MCDDEPLOY.TransactOpts, delay, authority)
}

// DeployShutdown is a paid mutator transaction binding the contract method 0xb0a842a8.
//
// Solidity: function deployShutdown(address gov, address pit, uint256 min) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployShutdown(opts *bind.TransactOpts, gov common.Address, pit common.Address, min *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployShutdown", gov, pit, min)
}

// DeployShutdown is a paid mutator transaction binding the contract method 0xb0a842a8.
//
// Solidity: function deployShutdown(address gov, address pit, uint256 min) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployShutdown(gov common.Address, pit common.Address, min *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployShutdown(&_MCDDEPLOY.TransactOpts, gov, pit, min)
}

// DeployShutdown is a paid mutator transaction binding the contract method 0xb0a842a8.
//
// Solidity: function deployShutdown(address gov, address pit, uint256 min) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployShutdown(gov common.Address, pit common.Address, min *big.Int) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployShutdown(&_MCDDEPLOY.TransactOpts, gov, pit, min)
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployTaxation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployTaxation")
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployTaxation() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployTaxation(&_MCDDEPLOY.TransactOpts)
}

// DeployTaxation is a paid mutator transaction binding the contract method 0x80e7cd89.
//
// Solidity: function deployTaxation() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployTaxation() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployTaxation(&_MCDDEPLOY.TransactOpts)
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) DeployVat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "deployVat")
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_MCDDEPLOY *MCDDEPLOYSession) DeployVat() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployVat(&_MCDDEPLOY.TransactOpts)
}

// DeployVat is a paid mutator transaction binding the contract method 0x2800a568.
//
// Solidity: function deployVat() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) DeployVat() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.DeployVat(&_MCDDEPLOY.TransactOpts)
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) ReleaseAuth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "releaseAuth")
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_MCDDEPLOY *MCDDEPLOYSession) ReleaseAuth() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.ReleaseAuth(&_MCDDEPLOY.TransactOpts)
}

// ReleaseAuth is a paid mutator transaction binding the contract method 0x70402bb9.
//
// Solidity: function releaseAuth() returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) ReleaseAuth() (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.ReleaseAuth(&_MCDDEPLOY.TransactOpts)
}

// ReleaseAuthFlip is a paid mutator transaction binding the contract method 0xb3731c78.
//
// Solidity: function releaseAuthFlip(bytes32 ilk) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) ReleaseAuthFlip(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "releaseAuthFlip", ilk)
}

// ReleaseAuthFlip is a paid mutator transaction binding the contract method 0xb3731c78.
//
// Solidity: function releaseAuthFlip(bytes32 ilk) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) ReleaseAuthFlip(ilk [32]byte) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.ReleaseAuthFlip(&_MCDDEPLOY.TransactOpts, ilk)
}

// ReleaseAuthFlip is a paid mutator transaction binding the contract method 0xb3731c78.
//
// Solidity: function releaseAuthFlip(bytes32 ilk) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) ReleaseAuthFlip(ilk [32]byte) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.ReleaseAuthFlip(&_MCDDEPLOY.TransactOpts, ilk)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.SetAuthority(&_MCDDEPLOY.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address authority_) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.SetAuthority(&_MCDDEPLOY.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDDEPLOY *MCDDEPLOYSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.SetOwner(&_MCDDEPLOY.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_MCDDEPLOY *MCDDEPLOYTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _MCDDEPLOY.Contract.SetOwner(&_MCDDEPLOY.TransactOpts, owner_)
}

// MCDDEPLOYLogSetAuthorityIterator is returned from FilterLogSetAuthority and is used to iterate over the raw logs and unpacked data for LogSetAuthority events raised by the MCDDEPLOY contract.
type MCDDEPLOYLogSetAuthorityIterator struct {
	Event *MCDDEPLOYLogSetAuthority // Event containing the contract specifics and raw log

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
func (it *MCDDEPLOYLogSetAuthorityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDEPLOYLogSetAuthority)
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
		it.Event = new(MCDDEPLOYLogSetAuthority)
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
func (it *MCDDEPLOYLogSetAuthorityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDEPLOYLogSetAuthorityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDEPLOYLogSetAuthority represents a LogSetAuthority event raised by the MCDDEPLOY contract.
type MCDDEPLOYLogSetAuthority struct {
	Authority common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogSetAuthority is a free log retrieval operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDDEPLOY *MCDDEPLOYFilterer) FilterLogSetAuthority(opts *bind.FilterOpts, authority []common.Address) (*MCDDEPLOYLogSetAuthorityIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDDEPLOY.contract.FilterLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOYLogSetAuthorityIterator{contract: _MCDDEPLOY.contract, event: "LogSetAuthority", logs: logs, sub: sub}, nil
}

// WatchLogSetAuthority is a free log subscription operation binding the contract event 0x1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada4.
//
// Solidity: event LogSetAuthority(address indexed authority)
func (_MCDDEPLOY *MCDDEPLOYFilterer) WatchLogSetAuthority(opts *bind.WatchOpts, sink chan<- *MCDDEPLOYLogSetAuthority, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _MCDDEPLOY.contract.WatchLogs(opts, "LogSetAuthority", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDEPLOYLogSetAuthority)
				if err := _MCDDEPLOY.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
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
func (_MCDDEPLOY *MCDDEPLOYFilterer) ParseLogSetAuthority(log types.Log) (*MCDDEPLOYLogSetAuthority, error) {
	event := new(MCDDEPLOYLogSetAuthority)
	if err := _MCDDEPLOY.contract.UnpackLog(event, "LogSetAuthority", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MCDDEPLOYLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the MCDDEPLOY contract.
type MCDDEPLOYLogSetOwnerIterator struct {
	Event *MCDDEPLOYLogSetOwner // Event containing the contract specifics and raw log

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
func (it *MCDDEPLOYLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MCDDEPLOYLogSetOwner)
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
		it.Event = new(MCDDEPLOYLogSetOwner)
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
func (it *MCDDEPLOYLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MCDDEPLOYLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MCDDEPLOYLogSetOwner represents a LogSetOwner event raised by the MCDDEPLOY contract.
type MCDDEPLOYLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDDEPLOY *MCDDEPLOYFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*MCDDEPLOYLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDDEPLOY.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MCDDEPLOYLogSetOwnerIterator{contract: _MCDDEPLOY.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_MCDDEPLOY *MCDDEPLOYFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *MCDDEPLOYLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MCDDEPLOY.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MCDDEPLOYLogSetOwner)
				if err := _MCDDEPLOY.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
func (_MCDDEPLOY *MCDDEPLOYFilterer) ParseLogSetOwner(log types.Log) (*MCDDEPLOYLogSetOwner, error) {
	event := new(MCDDEPLOYLogSetOwner)
	if err := _MCDDEPLOY.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
