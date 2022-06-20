// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PROXY_ACTIONS

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

// PROXYACTIONSABI is the input ABI used to generate the binding from.
const PROXYACTIONSABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"cdpAllow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"apt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"daiJoin_join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"apt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"}],\"name\":\"ethJoin_join\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exitETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"exitGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"freeETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"freeGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"apt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"urn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferFrom\",\"type\":\"bool\"}],\"name\":\"gemJoin_join\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"give\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"giveToProxy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"lockETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"lockETHAndDraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferFrom\",\"type\":\"bool\"}],\"name\":\"lockGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferFrom\",\"type\":\"bool\"}],\"name\":\"lockGemAndDraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"}],\"name\":\"makeGemBag\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bag\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"obj\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"open\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockETHAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gntJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockGNTAndDraw\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bag\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jug\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferFrom\",\"type\":\"bool\"}],\"name\":\"openLockGemAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"quit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"safeLockETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferFrom\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"safeLockGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"safeWipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"safeWipeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdpSrc\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cdpOrg\",\"type\":\"uint256\"}],\"name\":\"shift\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gem\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ok\",\"type\":\"uint256\"}],\"name\":\"urnAllow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"wipeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"}],\"name\":\"wipeAllAndFreeETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"}],\"name\":\"wipeAllAndFreeGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"wipeAndFreeETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"wipeAndFreeGem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PROXYACTIONS is an auto generated Go binding around an Ethereum contract.
type PROXYACTIONS struct {
	PROXYACTIONSCaller     // Read-only binding to the contract
	PROXYACTIONSTransactor // Write-only binding to the contract
	PROXYACTIONSFilterer   // Log filterer for contract events
}

// PROXYACTIONSCaller is an auto generated read-only Go binding around an Ethereum contract.
type PROXYACTIONSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PROXYACTIONSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PROXYACTIONSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PROXYACTIONSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PROXYACTIONSSession struct {
	Contract     *PROXYACTIONS     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PROXYACTIONSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PROXYACTIONSCallerSession struct {
	Contract *PROXYACTIONSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PROXYACTIONSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PROXYACTIONSTransactorSession struct {
	Contract     *PROXYACTIONSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PROXYACTIONSRaw is an auto generated low-level Go binding around an Ethereum contract.
type PROXYACTIONSRaw struct {
	Contract *PROXYACTIONS // Generic contract binding to access the raw methods on
}

// PROXYACTIONSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PROXYACTIONSCallerRaw struct {
	Contract *PROXYACTIONSCaller // Generic read-only contract binding to access the raw methods on
}

// PROXYACTIONSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PROXYACTIONSTransactorRaw struct {
	Contract *PROXYACTIONSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPROXYACTIONS creates a new instance of PROXYACTIONS, bound to a specific deployed contract.
func NewPROXYACTIONS(address common.Address, backend bind.ContractBackend) (*PROXYACTIONS, error) {
	contract, err := bindPROXYACTIONS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONS{PROXYACTIONSCaller: PROXYACTIONSCaller{contract: contract}, PROXYACTIONSTransactor: PROXYACTIONSTransactor{contract: contract}, PROXYACTIONSFilterer: PROXYACTIONSFilterer{contract: contract}}, nil
}

// NewPROXYACTIONSCaller creates a new read-only instance of PROXYACTIONS, bound to a specific deployed contract.
func NewPROXYACTIONSCaller(address common.Address, caller bind.ContractCaller) (*PROXYACTIONSCaller, error) {
	contract, err := bindPROXYACTIONS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSCaller{contract: contract}, nil
}

// NewPROXYACTIONSTransactor creates a new write-only instance of PROXYACTIONS, bound to a specific deployed contract.
func NewPROXYACTIONSTransactor(address common.Address, transactor bind.ContractTransactor) (*PROXYACTIONSTransactor, error) {
	contract, err := bindPROXYACTIONS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSTransactor{contract: contract}, nil
}

// NewPROXYACTIONSFilterer creates a new log filterer instance of PROXYACTIONS, bound to a specific deployed contract.
func NewPROXYACTIONSFilterer(address common.Address, filterer bind.ContractFilterer) (*PROXYACTIONSFilterer, error) {
	contract, err := bindPROXYACTIONS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PROXYACTIONSFilterer{contract: contract}, nil
}

// bindPROXYACTIONS binds a generic wrapper to an already deployed contract.
func bindPROXYACTIONS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PROXYACTIONSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONS *PROXYACTIONSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONS.Contract.PROXYACTIONSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONS *PROXYACTIONSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.PROXYACTIONSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONS *PROXYACTIONSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.PROXYACTIONSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PROXYACTIONS *PROXYACTIONSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PROXYACTIONS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PROXYACTIONS *PROXYACTIONSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PROXYACTIONS *PROXYACTIONSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.contract.Transact(opts, method, params...)
}

// CdpAllow is a paid mutator transaction binding the contract method 0xba727a95.
//
// Solidity: function cdpAllow(address manager, uint256 cdp, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) CdpAllow(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "cdpAllow", manager, cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0xba727a95.
//
// Solidity: function cdpAllow(address manager, uint256 cdp, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) CdpAllow(manager common.Address, cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.CdpAllow(&_PROXYACTIONS.TransactOpts, manager, cdp, usr, ok)
}

// CdpAllow is a paid mutator transaction binding the contract method 0xba727a95.
//
// Solidity: function cdpAllow(address manager, uint256 cdp, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) CdpAllow(manager common.Address, cdp *big.Int, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.CdpAllow(&_PROXYACTIONS.TransactOpts, manager, cdp, usr, ok)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) DaiJoinJoin(opts *bind.TransactOpts, apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "daiJoin_join", apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.DaiJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn, wad)
}

// DaiJoinJoin is a paid mutator transaction binding the contract method 0xc56167c6.
//
// Solidity: function daiJoin_join(address apt, address urn, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) DaiJoinJoin(apt common.Address, urn common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.DaiJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x9f6f3d5b.
//
// Solidity: function draw(address manager, address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Draw(opts *bind.TransactOpts, manager common.Address, jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "draw", manager, jug, daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x9f6f3d5b.
//
// Solidity: function draw(address manager, address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Draw(manager common.Address, jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Draw(&_PROXYACTIONS.TransactOpts, manager, jug, daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0x9f6f3d5b.
//
// Solidity: function draw(address manager, address jug, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Draw(manager common.Address, jug common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Draw(&_PROXYACTIONS.TransactOpts, manager, jug, daiJoin, cdp, wad)
}

// Enter is a paid mutator transaction binding the contract method 0xeb0b9a85.
//
// Solidity: function enter(address manager, address src, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Enter(opts *bind.TransactOpts, manager common.Address, src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "enter", manager, src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0xeb0b9a85.
//
// Solidity: function enter(address manager, address src, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Enter(manager common.Address, src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Enter(&_PROXYACTIONS.TransactOpts, manager, src, cdp)
}

// Enter is a paid mutator transaction binding the contract method 0xeb0b9a85.
//
// Solidity: function enter(address manager, address src, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Enter(manager common.Address, src common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Enter(&_PROXYACTIONS.TransactOpts, manager, src, cdp)
}

// EthJoinJoin is a paid mutator transaction binding the contract method 0xa033df12.
//
// Solidity: function ethJoin_join(address apt, address urn) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) EthJoinJoin(opts *bind.TransactOpts, apt common.Address, urn common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "ethJoin_join", apt, urn)
}

// EthJoinJoin is a paid mutator transaction binding the contract method 0xa033df12.
//
// Solidity: function ethJoin_join(address apt, address urn) payable returns()
func (_PROXYACTIONS *PROXYACTIONSSession) EthJoinJoin(apt common.Address, urn common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.EthJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn)
}

// EthJoinJoin is a paid mutator transaction binding the contract method 0xa033df12.
//
// Solidity: function ethJoin_join(address apt, address urn) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) EthJoinJoin(apt common.Address, urn common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.EthJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn)
}

// ExitETH is a paid mutator transaction binding the contract method 0x08f00e34.
//
// Solidity: function exitETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) ExitETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "exitETH", manager, ethJoin, cdp, wad)
}

// ExitETH is a paid mutator transaction binding the contract method 0x08f00e34.
//
// Solidity: function exitETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) ExitETH(manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.ExitETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, wad)
}

// ExitETH is a paid mutator transaction binding the contract method 0x08f00e34.
//
// Solidity: function exitETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) ExitETH(manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.ExitETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, wad)
}

// ExitGem is a paid mutator transaction binding the contract method 0x42dd11bb.
//
// Solidity: function exitGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) ExitGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "exitGem", manager, gemJoin, cdp, wad)
}

// ExitGem is a paid mutator transaction binding the contract method 0x42dd11bb.
//
// Solidity: function exitGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) ExitGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.ExitGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad)
}

// ExitGem is a paid mutator transaction binding the contract method 0x42dd11bb.
//
// Solidity: function exitGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) ExitGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.ExitGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x1558b048.
//
// Solidity: function flux(address manager, uint256 cdp, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Flux(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "flux", manager, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x1558b048.
//
// Solidity: function flux(address manager, uint256 cdp, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Flux(manager common.Address, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Flux(&_PROXYACTIONS.TransactOpts, manager, cdp, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x1558b048.
//
// Solidity: function flux(address manager, uint256 cdp, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Flux(manager common.Address, cdp *big.Int, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Flux(&_PROXYACTIONS.TransactOpts, manager, cdp, dst, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x7b5a3b43.
//
// Solidity: function freeETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) FreeETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "freeETH", manager, ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x7b5a3b43.
//
// Solidity: function freeETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) FreeETH(manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.FreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x7b5a3b43.
//
// Solidity: function freeETH(address manager, address ethJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) FreeETH(manager common.Address, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.FreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x6ab6a491.
//
// Solidity: function freeGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) FreeGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "freeGem", manager, gemJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x6ab6a491.
//
// Solidity: function freeGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) FreeGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.FreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x6ab6a491.
//
// Solidity: function freeGem(address manager, address gemJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) FreeGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.FreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad)
}

// Frob is a paid mutator transaction binding the contract method 0x96e8d72c.
//
// Solidity: function frob(address manager, uint256 cdp, int256 dink, int256 dart) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Frob(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "frob", manager, cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x96e8d72c.
//
// Solidity: function frob(address manager, uint256 cdp, int256 dink, int256 dart) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Frob(manager common.Address, cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Frob(&_PROXYACTIONS.TransactOpts, manager, cdp, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x96e8d72c.
//
// Solidity: function frob(address manager, uint256 cdp, int256 dink, int256 dart) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Frob(manager common.Address, cdp *big.Int, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Frob(&_PROXYACTIONS.TransactOpts, manager, cdp, dink, dart)
}

// GemJoinJoin is a paid mutator transaction binding the contract method 0x7df2eb25.
//
// Solidity: function gemJoin_join(address apt, address urn, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) GemJoinJoin(opts *bind.TransactOpts, apt common.Address, urn common.Address, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "gemJoin_join", apt, urn, wad, transferFrom)
}

// GemJoinJoin is a paid mutator transaction binding the contract method 0x7df2eb25.
//
// Solidity: function gemJoin_join(address apt, address urn, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) GemJoinJoin(apt common.Address, urn common.Address, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.GemJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn, wad, transferFrom)
}

// GemJoinJoin is a paid mutator transaction binding the contract method 0x7df2eb25.
//
// Solidity: function gemJoin_join(address apt, address urn, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) GemJoinJoin(apt common.Address, urn common.Address, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.GemJoinJoin(&_PROXYACTIONS.TransactOpts, apt, urn, wad, transferFrom)
}

// Give is a paid mutator transaction binding the contract method 0x1d10f231.
//
// Solidity: function give(address manager, uint256 cdp, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Give(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "give", manager, cdp, usr)
}

// Give is a paid mutator transaction binding the contract method 0x1d10f231.
//
// Solidity: function give(address manager, uint256 cdp, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Give(manager common.Address, cdp *big.Int, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Give(&_PROXYACTIONS.TransactOpts, manager, cdp, usr)
}

// Give is a paid mutator transaction binding the contract method 0x1d10f231.
//
// Solidity: function give(address manager, uint256 cdp, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Give(manager common.Address, cdp *big.Int, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Give(&_PROXYACTIONS.TransactOpts, manager, cdp, usr)
}

// GiveToProxy is a paid mutator transaction binding the contract method 0x493c2049.
//
// Solidity: function giveToProxy(address proxyRegistry, address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) GiveToProxy(opts *bind.TransactOpts, proxyRegistry common.Address, manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "giveToProxy", proxyRegistry, manager, cdp, dst)
}

// GiveToProxy is a paid mutator transaction binding the contract method 0x493c2049.
//
// Solidity: function giveToProxy(address proxyRegistry, address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) GiveToProxy(proxyRegistry common.Address, manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.GiveToProxy(&_PROXYACTIONS.TransactOpts, proxyRegistry, manager, cdp, dst)
}

// GiveToProxy is a paid mutator transaction binding the contract method 0x493c2049.
//
// Solidity: function giveToProxy(address proxyRegistry, address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) GiveToProxy(proxyRegistry common.Address, manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.GiveToProxy(&_PROXYACTIONS.TransactOpts, proxyRegistry, manager, cdp, dst)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Hope(opts *bind.TransactOpts, obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "hope", obj, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Hope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Hope(&_PROXYACTIONS.TransactOpts, obj, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xb50a5869.
//
// Solidity: function hope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Hope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Hope(&_PROXYACTIONS.TransactOpts, obj, usr)
}

// LockETH is a paid mutator transaction binding the contract method 0xe205c108.
//
// Solidity: function lockETH(address manager, address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) LockETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "lockETH", manager, ethJoin, cdp)
}

// LockETH is a paid mutator transaction binding the contract method 0xe205c108.
//
// Solidity: function lockETH(address manager, address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONS *PROXYACTIONSSession) LockETH(manager common.Address, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp)
}

// LockETH is a paid mutator transaction binding the contract method 0xe205c108.
//
// Solidity: function lockETH(address manager, address ethJoin, uint256 cdp) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) LockETH(manager common.Address, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0x1c02d846.
//
// Solidity: function lockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) LockETHAndDraw(opts *bind.TransactOpts, manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "lockETHAndDraw", manager, jug, ethJoin, daiJoin, cdp, wadD)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0x1c02d846.
//
// Solidity: function lockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONS *PROXYACTIONSSession) LockETHAndDraw(manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockETHAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, ethJoin, daiJoin, cdp, wadD)
}

// LockETHAndDraw is a paid mutator transaction binding the contract method 0x1c02d846.
//
// Solidity: function lockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, uint256 cdp, uint256 wadD) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) LockETHAndDraw(manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockETHAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, ethJoin, daiJoin, cdp, wadD)
}

// LockGem is a paid mutator transaction binding the contract method 0x3e29e565.
//
// Solidity: function lockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) LockGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "lockGem", manager, gemJoin, cdp, wad, transferFrom)
}

// LockGem is a paid mutator transaction binding the contract method 0x3e29e565.
//
// Solidity: function lockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) LockGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad, transferFrom)
}

// LockGem is a paid mutator transaction binding the contract method 0x3e29e565.
//
// Solidity: function lockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) LockGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad, transferFrom)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0xcbd4be3f.
//
// Solidity: function lockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) LockGemAndDraw(opts *bind.TransactOpts, manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "lockGemAndDraw", manager, jug, gemJoin, daiJoin, cdp, wadC, wadD, transferFrom)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0xcbd4be3f.
//
// Solidity: function lockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) LockGemAndDraw(manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockGemAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gemJoin, daiJoin, cdp, wadC, wadD, transferFrom)
}

// LockGemAndDraw is a paid mutator transaction binding the contract method 0xcbd4be3f.
//
// Solidity: function lockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD, bool transferFrom) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) LockGemAndDraw(manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.LockGemAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gemJoin, daiJoin, cdp, wadC, wadD, transferFrom)
}

// MakeGemBag is a paid mutator transaction binding the contract method 0xfba7591d.
//
// Solidity: function makeGemBag(address gemJoin) returns(address bag)
func (_PROXYACTIONS *PROXYACTIONSTransactor) MakeGemBag(opts *bind.TransactOpts, gemJoin common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "makeGemBag", gemJoin)
}

// MakeGemBag is a paid mutator transaction binding the contract method 0xfba7591d.
//
// Solidity: function makeGemBag(address gemJoin) returns(address bag)
func (_PROXYACTIONS *PROXYACTIONSSession) MakeGemBag(gemJoin common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.MakeGemBag(&_PROXYACTIONS.TransactOpts, gemJoin)
}

// MakeGemBag is a paid mutator transaction binding the contract method 0xfba7591d.
//
// Solidity: function makeGemBag(address gemJoin) returns(address bag)
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) MakeGemBag(gemJoin common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.MakeGemBag(&_PROXYACTIONS.TransactOpts, gemJoin)
}

// Move is a paid mutator transaction binding the contract method 0x25cf37d0.
//
// Solidity: function move(address manager, uint256 cdp, address dst, uint256 rad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Move(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "move", manager, cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0x25cf37d0.
//
// Solidity: function move(address manager, uint256 cdp, address dst, uint256 rad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Move(manager common.Address, cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Move(&_PROXYACTIONS.TransactOpts, manager, cdp, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0x25cf37d0.
//
// Solidity: function move(address manager, uint256 cdp, address dst, uint256 rad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Move(manager common.Address, cdp *big.Int, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Move(&_PROXYACTIONS.TransactOpts, manager, cdp, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Nope(opts *bind.TransactOpts, obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "nope", obj, usr)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Nope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Nope(&_PROXYACTIONS.TransactOpts, obj, usr)
}

// Nope is a paid mutator transaction binding the contract method 0x9f887fde.
//
// Solidity: function nope(address obj, address usr) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Nope(obj common.Address, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Nope(&_PROXYACTIONS.TransactOpts, obj, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6aa3ee11.
//
// Solidity: function open(address manager, bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactor) Open(opts *bind.TransactOpts, manager common.Address, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "open", manager, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6aa3ee11.
//
// Solidity: function open(address manager, bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSSession) Open(manager common.Address, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Open(&_PROXYACTIONS.TransactOpts, manager, ilk, usr)
}

// Open is a paid mutator transaction binding the contract method 0x6aa3ee11.
//
// Solidity: function open(address manager, bytes32 ilk, address usr) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Open(manager common.Address, ilk [32]byte, usr common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Open(&_PROXYACTIONS.TransactOpts, manager, ilk, usr)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0xe685cc04.
//
// Solidity: function openLockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactor) OpenLockETHAndDraw(opts *bind.TransactOpts, manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "openLockETHAndDraw", manager, jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0xe685cc04.
//
// Solidity: function openLockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSSession) OpenLockETHAndDraw(manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockETHAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0xe685cc04.
//
// Solidity: function openLockETHAndDraw(address manager, address jug, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) OpenLockETHAndDraw(manager common.Address, jug common.Address, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockETHAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockGNTAndDraw is a paid mutator transaction binding the contract method 0xc0791710.
//
// Solidity: function openLockGNTAndDraw(address manager, address jug, address gntJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) returns(address bag, uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactor) OpenLockGNTAndDraw(opts *bind.TransactOpts, manager common.Address, jug common.Address, gntJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "openLockGNTAndDraw", manager, jug, gntJoin, daiJoin, ilk, wadC, wadD)
}

// OpenLockGNTAndDraw is a paid mutator transaction binding the contract method 0xc0791710.
//
// Solidity: function openLockGNTAndDraw(address manager, address jug, address gntJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) returns(address bag, uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSSession) OpenLockGNTAndDraw(manager common.Address, jug common.Address, gntJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockGNTAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gntJoin, daiJoin, ilk, wadC, wadD)
}

// OpenLockGNTAndDraw is a paid mutator transaction binding the contract method 0xc0791710.
//
// Solidity: function openLockGNTAndDraw(address manager, address jug, address gntJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) returns(address bag, uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) OpenLockGNTAndDraw(manager common.Address, jug common.Address, gntJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockGNTAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gntJoin, daiJoin, ilk, wadC, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xdb802a32.
//
// Solidity: function openLockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD, bool transferFrom) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactor) OpenLockGemAndDraw(opts *bind.TransactOpts, manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "openLockGemAndDraw", manager, jug, gemJoin, daiJoin, ilk, wadC, wadD, transferFrom)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xdb802a32.
//
// Solidity: function openLockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD, bool transferFrom) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSSession) OpenLockGemAndDraw(manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockGemAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gemJoin, daiJoin, ilk, wadC, wadD, transferFrom)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0xdb802a32.
//
// Solidity: function openLockGemAndDraw(address manager, address jug, address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD, bool transferFrom) returns(uint256 cdp)
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) OpenLockGemAndDraw(manager common.Address, jug common.Address, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int, transferFrom bool) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.OpenLockGemAndDraw(&_PROXYACTIONS.TransactOpts, manager, jug, gemJoin, daiJoin, ilk, wadC, wadD, transferFrom)
}

// Quit is a paid mutator transaction binding the contract method 0x4592aca7.
//
// Solidity: function quit(address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Quit(opts *bind.TransactOpts, manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "quit", manager, cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x4592aca7.
//
// Solidity: function quit(address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Quit(manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Quit(&_PROXYACTIONS.TransactOpts, manager, cdp, dst)
}

// Quit is a paid mutator transaction binding the contract method 0x4592aca7.
//
// Solidity: function quit(address manager, uint256 cdp, address dst) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Quit(manager common.Address, cdp *big.Int, dst common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Quit(&_PROXYACTIONS.TransactOpts, manager, cdp, dst)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xee284576.
//
// Solidity: function safeLockETH(address manager, address ethJoin, uint256 cdp, address owner) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) SafeLockETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "safeLockETH", manager, ethJoin, cdp, owner)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xee284576.
//
// Solidity: function safeLockETH(address manager, address ethJoin, uint256 cdp, address owner) payable returns()
func (_PROXYACTIONS *PROXYACTIONSSession) SafeLockETH(manager common.Address, ethJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeLockETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, owner)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xee284576.
//
// Solidity: function safeLockETH(address manager, address ethJoin, uint256 cdp, address owner) payable returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) SafeLockETH(manager common.Address, ethJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeLockETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, cdp, owner)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0xead64729.
//
// Solidity: function safeLockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) SafeLockGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "safeLockGem", manager, gemJoin, cdp, wad, transferFrom, owner)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0xead64729.
//
// Solidity: function safeLockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) SafeLockGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeLockGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad, transferFrom, owner)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0xead64729.
//
// Solidity: function safeLockGem(address manager, address gemJoin, uint256 cdp, uint256 wad, bool transferFrom, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) SafeLockGem(manager common.Address, gemJoin common.Address, cdp *big.Int, wad *big.Int, transferFrom bool, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeLockGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, cdp, wad, transferFrom, owner)
}

// SafeWipe is a paid mutator transaction binding the contract method 0x2958f8a5.
//
// Solidity: function safeWipe(address manager, address daiJoin, uint256 cdp, uint256 wad, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) SafeWipe(opts *bind.TransactOpts, manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "safeWipe", manager, daiJoin, cdp, wad, owner)
}

// SafeWipe is a paid mutator transaction binding the contract method 0x2958f8a5.
//
// Solidity: function safeWipe(address manager, address daiJoin, uint256 cdp, uint256 wad, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) SafeWipe(manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeWipe(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, wad, owner)
}

// SafeWipe is a paid mutator transaction binding the contract method 0x2958f8a5.
//
// Solidity: function safeWipe(address manager, address daiJoin, uint256 cdp, uint256 wad, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) SafeWipe(manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeWipe(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, wad, owner)
}

// SafeWipeAll is a paid mutator transaction binding the contract method 0x0aee8dec.
//
// Solidity: function safeWipeAll(address manager, address daiJoin, uint256 cdp, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) SafeWipeAll(opts *bind.TransactOpts, manager common.Address, daiJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "safeWipeAll", manager, daiJoin, cdp, owner)
}

// SafeWipeAll is a paid mutator transaction binding the contract method 0x0aee8dec.
//
// Solidity: function safeWipeAll(address manager, address daiJoin, uint256 cdp, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) SafeWipeAll(manager common.Address, daiJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeWipeAll(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, owner)
}

// SafeWipeAll is a paid mutator transaction binding the contract method 0x0aee8dec.
//
// Solidity: function safeWipeAll(address manager, address daiJoin, uint256 cdp, address owner) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) SafeWipeAll(manager common.Address, daiJoin common.Address, cdp *big.Int, owner common.Address) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.SafeWipeAll(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, owner)
}

// Shift is a paid mutator transaction binding the contract method 0x7bc3bd53.
//
// Solidity: function shift(address manager, uint256 cdpSrc, uint256 cdpOrg) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Shift(opts *bind.TransactOpts, manager common.Address, cdpSrc *big.Int, cdpOrg *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "shift", manager, cdpSrc, cdpOrg)
}

// Shift is a paid mutator transaction binding the contract method 0x7bc3bd53.
//
// Solidity: function shift(address manager, uint256 cdpSrc, uint256 cdpOrg) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Shift(manager common.Address, cdpSrc *big.Int, cdpOrg *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Shift(&_PROXYACTIONS.TransactOpts, manager, cdpSrc, cdpOrg)
}

// Shift is a paid mutator transaction binding the contract method 0x7bc3bd53.
//
// Solidity: function shift(address manager, uint256 cdpSrc, uint256 cdpOrg) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Shift(manager common.Address, cdpSrc *big.Int, cdpOrg *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Shift(&_PROXYACTIONS.TransactOpts, manager, cdpSrc, cdpOrg)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Transfer(opts *bind.TransactOpts, gem common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "transfer", gem, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Transfer(gem common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Transfer(&_PROXYACTIONS.TransactOpts, gem, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address gem, address dst, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Transfer(gem common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Transfer(&_PROXYACTIONS.TransactOpts, gem, dst, wad)
}

// UrnAllow is a paid mutator transaction binding the contract method 0x5f6ef447.
//
// Solidity: function urnAllow(address manager, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) UrnAllow(opts *bind.TransactOpts, manager common.Address, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "urnAllow", manager, usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0x5f6ef447.
//
// Solidity: function urnAllow(address manager, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) UrnAllow(manager common.Address, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.UrnAllow(&_PROXYACTIONS.TransactOpts, manager, usr, ok)
}

// UrnAllow is a paid mutator transaction binding the contract method 0x5f6ef447.
//
// Solidity: function urnAllow(address manager, address usr, uint256 ok) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) UrnAllow(manager common.Address, usr common.Address, ok *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.UrnAllow(&_PROXYACTIONS.TransactOpts, manager, usr, ok)
}

// Wipe is a paid mutator transaction binding the contract method 0x4b666199.
//
// Solidity: function wipe(address manager, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) Wipe(opts *bind.TransactOpts, manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipe", manager, daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0x4b666199.
//
// Solidity: function wipe(address manager, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) Wipe(manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Wipe(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0x4b666199.
//
// Solidity: function wipe(address manager, address daiJoin, uint256 cdp, uint256 wad) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) Wipe(manager common.Address, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.Wipe(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp, wad)
}

// WipeAll is a paid mutator transaction binding the contract method 0x036a2395.
//
// Solidity: function wipeAll(address manager, address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) WipeAll(opts *bind.TransactOpts, manager common.Address, daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipeAll", manager, daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x036a2395.
//
// Solidity: function wipeAll(address manager, address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) WipeAll(manager common.Address, daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAll(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x036a2395.
//
// Solidity: function wipeAll(address manager, address daiJoin, uint256 cdp) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) WipeAll(manager common.Address, daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAll(&_PROXYACTIONS.TransactOpts, manager, daiJoin, cdp)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x6d68b70b.
//
// Solidity: function wipeAllAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) WipeAllAndFreeETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipeAllAndFreeETH", manager, ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x6d68b70b.
//
// Solidity: function wipeAllAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) WipeAllAndFreeETH(manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAllAndFreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeETH is a paid mutator transaction binding the contract method 0x6d68b70b.
//
// Solidity: function wipeAllAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) WipeAllAndFreeETH(manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAllAndFreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xbcd6deec.
//
// Solidity: function wipeAllAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) WipeAllAndFreeGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipeAllAndFreeGem", manager, gemJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xbcd6deec.
//
// Solidity: function wipeAllAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) WipeAllAndFreeGem(manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAllAndFreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, daiJoin, cdp, wadC)
}

// WipeAllAndFreeGem is a paid mutator transaction binding the contract method 0xbcd6deec.
//
// Solidity: function wipeAllAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) WipeAllAndFreeGem(manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAllAndFreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, daiJoin, cdp, wadC)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0xbe5e6c03.
//
// Solidity: function wipeAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) WipeAndFreeETH(opts *bind.TransactOpts, manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipeAndFreeETH", manager, ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0xbe5e6c03.
//
// Solidity: function wipeAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) WipeAndFreeETH(manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAndFreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeETH is a paid mutator transaction binding the contract method 0xbe5e6c03.
//
// Solidity: function wipeAndFreeETH(address manager, address ethJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) WipeAndFreeETH(manager common.Address, ethJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAndFreeETH(&_PROXYACTIONS.TransactOpts, manager, ethJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0xa6add011.
//
// Solidity: function wipeAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactor) WipeAndFreeGem(opts *bind.TransactOpts, manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.contract.Transact(opts, "wipeAndFreeGem", manager, gemJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0xa6add011.
//
// Solidity: function wipeAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSSession) WipeAndFreeGem(manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAndFreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, daiJoin, cdp, wadC, wadD)
}

// WipeAndFreeGem is a paid mutator transaction binding the contract method 0xa6add011.
//
// Solidity: function wipeAndFreeGem(address manager, address gemJoin, address daiJoin, uint256 cdp, uint256 wadC, uint256 wadD) returns()
func (_PROXYACTIONS *PROXYACTIONSTransactorSession) WipeAndFreeGem(manager common.Address, gemJoin common.Address, daiJoin common.Address, cdp *big.Int, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _PROXYACTIONS.Contract.WipeAndFreeGem(&_PROXYACTIONS.TransactOpts, manager, gemJoin, daiJoin, cdp, wadC, wadD)
}
