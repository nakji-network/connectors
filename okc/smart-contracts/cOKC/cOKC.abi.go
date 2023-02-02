// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cOKC

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// XENCryptoMintInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoMintInfo struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}

// XENCryptoStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoStakeInfo struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}

// COKCMetaData contains all meta data concerning the COKC contract.
var COKCMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"MintClaimed\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"name\":\"RankClaimed\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"payable\":false,\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"DAYS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"EAA_PM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"EAA_PM_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"EAA_RANK_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"GENESIS_RANK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"MAX_PENALTY_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"MAX_TERM_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"MAX_TERM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"MIN_TERM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"SECONDS_IN_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"TERM_AMPLIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"TERM_AMPLIFIER_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"WITHDRAWAL_WINDOW_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"XEN_APY_DAYS_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"XEN_APY_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"XEN_APY_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"XEN_MIN_BURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"XEN_MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activeMinters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"genesisTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"globalRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalXenStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBurns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rankDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eaa\",\"type\":\"uint256\"}],\"name\":\"getGrossReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getUserMint\",\"outputs\":[{\"components\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"term\",\"type\":\"uint256\"},{\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"name\":\"rank\",\"type\":\"uint256\"},{\"name\":\"amplifier\",\"type\":\"uint256\"},{\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.MintInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getUserStake\",\"outputs\":[{\"components\":[{\"name\":\"term\",\"type\":\"uint256\"},{\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"apy\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getCurrentAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getCurrentEAAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getCurrentAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getCurrentMaxTerm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimRank\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimMintReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"other\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndShare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// COKCABI is the input ABI used to generate the binding from.
// Deprecated: Use COKCMetaData.ABI instead.
var COKCABI = COKCMetaData.ABI

// COKC is an auto generated Go binding around an Ethereum contract.
type COKC struct {
	COKCCaller     // Read-only binding to the contract
	COKCTransactor // Write-only binding to the contract
	COKCFilterer   // Log filterer for contract events
}

// COKCCaller is an auto generated read-only Go binding around an Ethereum contract.
type COKCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// COKCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type COKCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// COKCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type COKCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// COKCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type COKCSession struct {
	Contract     *COKC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// COKCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type COKCCallerSession struct {
	Contract *COKCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// COKCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type COKCTransactorSession struct {
	Contract     *COKCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// COKCRaw is an auto generated low-level Go binding around an Ethereum contract.
type COKCRaw struct {
	Contract *COKC // Generic contract binding to access the raw methods on
}

// COKCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type COKCCallerRaw struct {
	Contract *COKCCaller // Generic read-only contract binding to access the raw methods on
}

// COKCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type COKCTransactorRaw struct {
	Contract *COKCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCOKC creates a new instance of COKC, bound to a specific deployed contract.
func NewCOKC(address common.Address, backend bind.ContractBackend) (*COKC, error) {
	contract, err := bindCOKC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &COKC{COKCCaller: COKCCaller{contract: contract}, COKCTransactor: COKCTransactor{contract: contract}, COKCFilterer: COKCFilterer{contract: contract}}, nil
}

// NewCOKCCaller creates a new read-only instance of COKC, bound to a specific deployed contract.
func NewCOKCCaller(address common.Address, caller bind.ContractCaller) (*COKCCaller, error) {
	contract, err := bindCOKC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &COKCCaller{contract: contract}, nil
}

// NewCOKCTransactor creates a new write-only instance of COKC, bound to a specific deployed contract.
func NewCOKCTransactor(address common.Address, transactor bind.ContractTransactor) (*COKCTransactor, error) {
	contract, err := bindCOKC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &COKCTransactor{contract: contract}, nil
}

// NewCOKCFilterer creates a new log filterer instance of COKC, bound to a specific deployed contract.
func NewCOKCFilterer(address common.Address, filterer bind.ContractFilterer) (*COKCFilterer, error) {
	contract, err := bindCOKC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &COKCFilterer{contract: contract}, nil
}

// bindCOKC binds a generic wrapper to an already deployed contract.
func bindCOKC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(COKCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_COKC *COKCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _COKC.Contract.COKCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_COKC *COKCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _COKC.Contract.COKCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_COKC *COKCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _COKC.Contract.COKCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_COKC *COKCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _COKC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_COKC *COKCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _COKC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_COKC *COKCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _COKC.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_COKC *COKCCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_COKC *COKCSession) AUTHORS() (string, error) {
	return _COKC.Contract.AUTHORS(&_COKC.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_COKC *COKCCallerSession) AUTHORS() (string, error) {
	return _COKC.Contract.AUTHORS(&_COKC.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_COKC *COKCCaller) DAYSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "DAYS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_COKC *COKCSession) DAYSINYEAR() (*big.Int, error) {
	return _COKC.Contract.DAYSINYEAR(&_COKC.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_COKC *COKCCallerSession) DAYSINYEAR() (*big.Int, error) {
	return _COKC.Contract.DAYSINYEAR(&_COKC.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_COKC *COKCCaller) EAAPMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "EAA_PM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_COKC *COKCSession) EAAPMSTART() (*big.Int, error) {
	return _COKC.Contract.EAAPMSTART(&_COKC.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_COKC *COKCCallerSession) EAAPMSTART() (*big.Int, error) {
	return _COKC.Contract.EAAPMSTART(&_COKC.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_COKC *COKCCaller) EAAPMSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "EAA_PM_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_COKC *COKCSession) EAAPMSTEP() (*big.Int, error) {
	return _COKC.Contract.EAAPMSTEP(&_COKC.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_COKC *COKCCallerSession) EAAPMSTEP() (*big.Int, error) {
	return _COKC.Contract.EAAPMSTEP(&_COKC.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_COKC *COKCCaller) EAARANKSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "EAA_RANK_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_COKC *COKCSession) EAARANKSTEP() (*big.Int, error) {
	return _COKC.Contract.EAARANKSTEP(&_COKC.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_COKC *COKCCallerSession) EAARANKSTEP() (*big.Int, error) {
	return _COKC.Contract.EAARANKSTEP(&_COKC.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_COKC *COKCCaller) GENESISRANK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "GENESIS_RANK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_COKC *COKCSession) GENESISRANK() (*big.Int, error) {
	return _COKC.Contract.GENESISRANK(&_COKC.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_COKC *COKCCallerSession) GENESISRANK() (*big.Int, error) {
	return _COKC.Contract.GENESISRANK(&_COKC.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_COKC *COKCCaller) MAXPENALTYPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "MAX_PENALTY_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_COKC *COKCSession) MAXPENALTYPCT() (*big.Int, error) {
	return _COKC.Contract.MAXPENALTYPCT(&_COKC.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_COKC *COKCCallerSession) MAXPENALTYPCT() (*big.Int, error) {
	return _COKC.Contract.MAXPENALTYPCT(&_COKC.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_COKC *COKCCaller) MAXTERMEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "MAX_TERM_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_COKC *COKCSession) MAXTERMEND() (*big.Int, error) {
	return _COKC.Contract.MAXTERMEND(&_COKC.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_COKC *COKCCallerSession) MAXTERMEND() (*big.Int, error) {
	return _COKC.Contract.MAXTERMEND(&_COKC.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_COKC *COKCCaller) MAXTERMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "MAX_TERM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_COKC *COKCSession) MAXTERMSTART() (*big.Int, error) {
	return _COKC.Contract.MAXTERMSTART(&_COKC.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_COKC *COKCCallerSession) MAXTERMSTART() (*big.Int, error) {
	return _COKC.Contract.MAXTERMSTART(&_COKC.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_COKC *COKCCaller) MINTERM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "MIN_TERM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_COKC *COKCSession) MINTERM() (*big.Int, error) {
	return _COKC.Contract.MINTERM(&_COKC.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_COKC *COKCCallerSession) MINTERM() (*big.Int, error) {
	return _COKC.Contract.MINTERM(&_COKC.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_COKC *COKCCaller) REWARDAMPLIFIEREND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "REWARD_AMPLIFIER_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_COKC *COKCSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _COKC.Contract.REWARDAMPLIFIEREND(&_COKC.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_COKC *COKCCallerSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _COKC.Contract.REWARDAMPLIFIEREND(&_COKC.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_COKC *COKCCaller) REWARDAMPLIFIERSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "REWARD_AMPLIFIER_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_COKC *COKCSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _COKC.Contract.REWARDAMPLIFIERSTART(&_COKC.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_COKC *COKCCallerSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _COKC.Contract.REWARDAMPLIFIERSTART(&_COKC.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_COKC *COKCCaller) SECONDSINDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "SECONDS_IN_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_COKC *COKCSession) SECONDSINDAY() (*big.Int, error) {
	return _COKC.Contract.SECONDSINDAY(&_COKC.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_COKC *COKCCallerSession) SECONDSINDAY() (*big.Int, error) {
	return _COKC.Contract.SECONDSINDAY(&_COKC.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_COKC *COKCCaller) TERMAMPLIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "TERM_AMPLIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_COKC *COKCSession) TERMAMPLIFIER() (*big.Int, error) {
	return _COKC.Contract.TERMAMPLIFIER(&_COKC.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_COKC *COKCCallerSession) TERMAMPLIFIER() (*big.Int, error) {
	return _COKC.Contract.TERMAMPLIFIER(&_COKC.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_COKC *COKCCaller) TERMAMPLIFIERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "TERM_AMPLIFIER_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_COKC *COKCSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _COKC.Contract.TERMAMPLIFIERTHRESHOLD(&_COKC.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_COKC *COKCCallerSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _COKC.Contract.TERMAMPLIFIERTHRESHOLD(&_COKC.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_COKC *COKCCaller) WITHDRAWALWINDOWDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "WITHDRAWAL_WINDOW_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_COKC *COKCSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _COKC.Contract.WITHDRAWALWINDOWDAYS(&_COKC.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_COKC *COKCCallerSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _COKC.Contract.WITHDRAWALWINDOWDAYS(&_COKC.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_COKC *COKCCaller) XENAPYDAYSSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "XEN_APY_DAYS_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_COKC *COKCSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _COKC.Contract.XENAPYDAYSSTEP(&_COKC.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_COKC *COKCCallerSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _COKC.Contract.XENAPYDAYSSTEP(&_COKC.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_COKC *COKCCaller) XENAPYEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "XEN_APY_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_COKC *COKCSession) XENAPYEND() (*big.Int, error) {
	return _COKC.Contract.XENAPYEND(&_COKC.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_COKC *COKCCallerSession) XENAPYEND() (*big.Int, error) {
	return _COKC.Contract.XENAPYEND(&_COKC.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_COKC *COKCCaller) XENAPYSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "XEN_APY_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_COKC *COKCSession) XENAPYSTART() (*big.Int, error) {
	return _COKC.Contract.XENAPYSTART(&_COKC.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_COKC *COKCCallerSession) XENAPYSTART() (*big.Int, error) {
	return _COKC.Contract.XENAPYSTART(&_COKC.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_COKC *COKCCaller) XENMINBURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "XEN_MIN_BURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_COKC *COKCSession) XENMINBURN() (*big.Int, error) {
	return _COKC.Contract.XENMINBURN(&_COKC.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_COKC *COKCCallerSession) XENMINBURN() (*big.Int, error) {
	return _COKC.Contract.XENMINBURN(&_COKC.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_COKC *COKCCaller) XENMINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "XEN_MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_COKC *COKCSession) XENMINSTAKE() (*big.Int, error) {
	return _COKC.Contract.XENMINSTAKE(&_COKC.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_COKC *COKCCallerSession) XENMINSTAKE() (*big.Int, error) {
	return _COKC.Contract.XENMINSTAKE(&_COKC.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_COKC *COKCCaller) ActiveMinters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "activeMinters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_COKC *COKCSession) ActiveMinters() (*big.Int, error) {
	return _COKC.Contract.ActiveMinters(&_COKC.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_COKC *COKCCallerSession) ActiveMinters() (*big.Int, error) {
	return _COKC.Contract.ActiveMinters(&_COKC.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_COKC *COKCCaller) ActiveStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "activeStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_COKC *COKCSession) ActiveStakes() (*big.Int, error) {
	return _COKC.Contract.ActiveStakes(&_COKC.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_COKC *COKCCallerSession) ActiveStakes() (*big.Int, error) {
	return _COKC.Contract.ActiveStakes(&_COKC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_COKC *COKCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_COKC *COKCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _COKC.Contract.Allowance(&_COKC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_COKC *COKCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _COKC.Contract.Allowance(&_COKC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_COKC *COKCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_COKC *COKCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _COKC.Contract.BalanceOf(&_COKC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_COKC *COKCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _COKC.Contract.BalanceOf(&_COKC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_COKC *COKCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_COKC *COKCSession) Decimals() (uint8, error) {
	return _COKC.Contract.Decimals(&_COKC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_COKC *COKCCallerSession) Decimals() (uint8, error) {
	return _COKC.Contract.Decimals(&_COKC.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_COKC *COKCCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_COKC *COKCSession) GenesisTs() (*big.Int, error) {
	return _COKC.Contract.GenesisTs(&_COKC.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_COKC *COKCCallerSession) GenesisTs() (*big.Int, error) {
	return _COKC.Contract.GenesisTs(&_COKC.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_COKC *COKCCaller) GetCurrentAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getCurrentAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_COKC *COKCSession) GetCurrentAMP() (*big.Int, error) {
	return _COKC.Contract.GetCurrentAMP(&_COKC.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_COKC *COKCCallerSession) GetCurrentAMP() (*big.Int, error) {
	return _COKC.Contract.GetCurrentAMP(&_COKC.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_COKC *COKCCaller) GetCurrentAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getCurrentAPY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_COKC *COKCSession) GetCurrentAPY() (*big.Int, error) {
	return _COKC.Contract.GetCurrentAPY(&_COKC.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_COKC *COKCCallerSession) GetCurrentAPY() (*big.Int, error) {
	return _COKC.Contract.GetCurrentAPY(&_COKC.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_COKC *COKCCaller) GetCurrentEAAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getCurrentEAAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_COKC *COKCSession) GetCurrentEAAR() (*big.Int, error) {
	return _COKC.Contract.GetCurrentEAAR(&_COKC.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_COKC *COKCCallerSession) GetCurrentEAAR() (*big.Int, error) {
	return _COKC.Contract.GetCurrentEAAR(&_COKC.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_COKC *COKCCaller) GetCurrentMaxTerm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getCurrentMaxTerm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_COKC *COKCSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _COKC.Contract.GetCurrentMaxTerm(&_COKC.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_COKC *COKCCallerSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _COKC.Contract.GetCurrentMaxTerm(&_COKC.CallOpts)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_COKC *COKCCaller) GetGrossReward(opts *bind.CallOpts, rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getGrossReward", rankDelta, amplifier, term, eaa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_COKC *COKCSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _COKC.Contract.GetGrossReward(&_COKC.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_COKC *COKCCallerSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _COKC.Contract.GetGrossReward(&_COKC.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_COKC *COKCCaller) GetUserMint(opts *bind.CallOpts) (XENCryptoMintInfo, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getUserMint")

	if err != nil {
		return *new(XENCryptoMintInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoMintInfo)).(*XENCryptoMintInfo)

	return out0, err

}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_COKC *COKCSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _COKC.Contract.GetUserMint(&_COKC.CallOpts)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_COKC *COKCCallerSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _COKC.Contract.GetUserMint(&_COKC.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_COKC *COKCCaller) GetUserStake(opts *bind.CallOpts) (XENCryptoStakeInfo, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "getUserStake")

	if err != nil {
		return *new(XENCryptoStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoStakeInfo)).(*XENCryptoStakeInfo)

	return out0, err

}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_COKC *COKCSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _COKC.Contract.GetUserStake(&_COKC.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_COKC *COKCCallerSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _COKC.Contract.GetUserStake(&_COKC.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_COKC *COKCCaller) GlobalRank(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "globalRank")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_COKC *COKCSession) GlobalRank() (*big.Int, error) {
	return _COKC.Contract.GlobalRank(&_COKC.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_COKC *COKCCallerSession) GlobalRank() (*big.Int, error) {
	return _COKC.Contract.GlobalRank(&_COKC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_COKC *COKCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_COKC *COKCSession) Name() (string, error) {
	return _COKC.Contract.Name(&_COKC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_COKC *COKCCallerSession) Name() (string, error) {
	return _COKC.Contract.Name(&_COKC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_COKC *COKCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_COKC *COKCSession) Symbol() (string, error) {
	return _COKC.Contract.Symbol(&_COKC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_COKC *COKCCallerSession) Symbol() (string, error) {
	return _COKC.Contract.Symbol(&_COKC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_COKC *COKCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_COKC *COKCSession) TotalSupply() (*big.Int, error) {
	return _COKC.Contract.TotalSupply(&_COKC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_COKC *COKCCallerSession) TotalSupply() (*big.Int, error) {
	return _COKC.Contract.TotalSupply(&_COKC.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_COKC *COKCCaller) TotalXenStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "totalXenStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_COKC *COKCSession) TotalXenStaked() (*big.Int, error) {
	return _COKC.Contract.TotalXenStaked(&_COKC.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_COKC *COKCCallerSession) TotalXenStaked() (*big.Int, error) {
	return _COKC.Contract.TotalXenStaked(&_COKC.CallOpts)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_COKC *COKCCaller) UserBurns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "userBurns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_COKC *COKCSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _COKC.Contract.UserBurns(&_COKC.CallOpts, arg0)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_COKC *COKCCallerSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _COKC.Contract.UserBurns(&_COKC.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_COKC *COKCCaller) UserMints(opts *bind.CallOpts, arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "userMints", arg0)

	outstruct := new(struct {
		User       common.Address
		Term       *big.Int
		MaturityTs *big.Int
		Rank       *big.Int
		Amplifier  *big.Int
		EaaRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Term = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amplifier = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EaaRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_COKC *COKCSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _COKC.Contract.UserMints(&_COKC.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_COKC *COKCCallerSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _COKC.Contract.UserMints(&_COKC.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_COKC *COKCCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	var out []interface{}
	err := _COKC.contract.Call(opts, &out, "userStakes", arg0)

	outstruct := new(struct {
		Term       *big.Int
		MaturityTs *big.Int
		Amount     *big.Int
		Apy        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Term = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Apy = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_COKC *COKCSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _COKC.Contract.UserStakes(&_COKC.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_COKC *COKCCallerSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _COKC.Contract.UserStakes(&_COKC.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_COKC *COKCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_COKC *COKCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Approve(&_COKC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_COKC *COKCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Approve(&_COKC.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_COKC *COKCTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_COKC *COKCSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Burn(&_COKC.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_COKC *COKCTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Burn(&_COKC.TransactOpts, user, amount)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_COKC *COKCTransactor) ClaimMintReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "claimMintReward")
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_COKC *COKCSession) ClaimMintReward() (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintReward(&_COKC.TransactOpts)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_COKC *COKCTransactorSession) ClaimMintReward() (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintReward(&_COKC.TransactOpts)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_COKC *COKCTransactor) ClaimMintRewardAndShare(opts *bind.TransactOpts, other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "claimMintRewardAndShare", other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_COKC *COKCSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintRewardAndShare(&_COKC.TransactOpts, other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_COKC *COKCTransactorSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintRewardAndShare(&_COKC.TransactOpts, other, pct)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_COKC *COKCTransactor) ClaimMintRewardAndStake(opts *bind.TransactOpts, pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "claimMintRewardAndStake", pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_COKC *COKCSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintRewardAndStake(&_COKC.TransactOpts, pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_COKC *COKCTransactorSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimMintRewardAndStake(&_COKC.TransactOpts, pct, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_COKC *COKCTransactor) ClaimRank(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "claimRank", term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_COKC *COKCSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimRank(&_COKC.TransactOpts, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_COKC *COKCTransactorSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.ClaimRank(&_COKC.TransactOpts, term)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_COKC *COKCTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_COKC *COKCSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.DecreaseAllowance(&_COKC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_COKC *COKCTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.DecreaseAllowance(&_COKC.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_COKC *COKCTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_COKC *COKCSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.IncreaseAllowance(&_COKC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_COKC *COKCTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.IncreaseAllowance(&_COKC.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_COKC *COKCTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "stake", amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_COKC *COKCSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Stake(&_COKC.TransactOpts, amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_COKC *COKCTransactorSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Stake(&_COKC.TransactOpts, amount, term)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_COKC *COKCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_COKC *COKCSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Transfer(&_COKC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_COKC *COKCTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.Transfer(&_COKC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_COKC *COKCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_COKC *COKCSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.TransferFrom(&_COKC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_COKC *COKCTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _COKC.Contract.TransferFrom(&_COKC.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_COKC *COKCTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _COKC.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_COKC *COKCSession) Withdraw() (*types.Transaction, error) {
	return _COKC.Contract.Withdraw(&_COKC.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_COKC *COKCTransactorSession) Withdraw() (*types.Transaction, error) {
	return _COKC.Contract.Withdraw(&_COKC.TransactOpts)
}

// COKCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the COKC contract.
type COKCApprovalIterator struct {
	Event *COKCApproval // Event containing the contract specifics and raw log

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
func (it *COKCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCApproval)
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
		it.Event = new(COKCApproval)
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
func (it *COKCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCApproval represents a Approval event raised by the COKC contract.
type COKCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_COKC *COKCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*COKCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &COKCApprovalIterator{contract: _COKC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_COKC *COKCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *COKCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCApproval)
				if err := _COKC.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_COKC *COKCFilterer) ParseApproval(log types.Log) (*COKCApproval, error) {
	event := new(COKCApproval)
	if err := _COKC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// COKCMintClaimedIterator is returned from FilterMintClaimed and is used to iterate over the raw logs and unpacked data for MintClaimed events raised by the COKC contract.
type COKCMintClaimedIterator struct {
	Event *COKCMintClaimed // Event containing the contract specifics and raw log

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
func (it *COKCMintClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCMintClaimed)
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
		it.Event = new(COKCMintClaimed)
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
func (it *COKCMintClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCMintClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCMintClaimed represents a MintClaimed event raised by the COKC contract.
type COKCMintClaimed struct {
	User         common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintClaimed is a free log retrieval operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_COKC *COKCFilterer) FilterMintClaimed(opts *bind.FilterOpts, user []common.Address) (*COKCMintClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &COKCMintClaimedIterator{contract: _COKC.contract, event: "MintClaimed", logs: logs, sub: sub}, nil
}

// WatchMintClaimed is a free log subscription operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_COKC *COKCFilterer) WatchMintClaimed(opts *bind.WatchOpts, sink chan<- *COKCMintClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCMintClaimed)
				if err := _COKC.contract.UnpackLog(event, "MintClaimed", log); err != nil {
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

// ParseMintClaimed is a log parse operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_COKC *COKCFilterer) ParseMintClaimed(log types.Log) (*COKCMintClaimed, error) {
	event := new(COKCMintClaimed)
	if err := _COKC.contract.UnpackLog(event, "MintClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// COKCRankClaimedIterator is returned from FilterRankClaimed and is used to iterate over the raw logs and unpacked data for RankClaimed events raised by the COKC contract.
type COKCRankClaimedIterator struct {
	Event *COKCRankClaimed // Event containing the contract specifics and raw log

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
func (it *COKCRankClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCRankClaimed)
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
		it.Event = new(COKCRankClaimed)
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
func (it *COKCRankClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCRankClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCRankClaimed represents a RankClaimed event raised by the COKC contract.
type COKCRankClaimed struct {
	User common.Address
	Term *big.Int
	Rank *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRankClaimed is a free log retrieval operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_COKC *COKCFilterer) FilterRankClaimed(opts *bind.FilterOpts, user []common.Address) (*COKCRankClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &COKCRankClaimedIterator{contract: _COKC.contract, event: "RankClaimed", logs: logs, sub: sub}, nil
}

// WatchRankClaimed is a free log subscription operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_COKC *COKCFilterer) WatchRankClaimed(opts *bind.WatchOpts, sink chan<- *COKCRankClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCRankClaimed)
				if err := _COKC.contract.UnpackLog(event, "RankClaimed", log); err != nil {
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

// ParseRankClaimed is a log parse operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_COKC *COKCFilterer) ParseRankClaimed(log types.Log) (*COKCRankClaimed, error) {
	event := new(COKCRankClaimed)
	if err := _COKC.contract.UnpackLog(event, "RankClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// COKCStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the COKC contract.
type COKCStakedIterator struct {
	Event *COKCStaked // Event containing the contract specifics and raw log

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
func (it *COKCStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCStaked)
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
		it.Event = new(COKCStaked)
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
func (it *COKCStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCStaked represents a Staked event raised by the COKC contract.
type COKCStaked struct {
	User   common.Address
	Amount *big.Int
	Term   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_COKC *COKCFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*COKCStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &COKCStakedIterator{contract: _COKC.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_COKC *COKCFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *COKCStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCStaked)
				if err := _COKC.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_COKC *COKCFilterer) ParseStaked(log types.Log) (*COKCStaked, error) {
	event := new(COKCStaked)
	if err := _COKC.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// COKCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the COKC contract.
type COKCTransferIterator struct {
	Event *COKCTransfer // Event containing the contract specifics and raw log

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
func (it *COKCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCTransfer)
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
		it.Event = new(COKCTransfer)
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
func (it *COKCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCTransfer represents a Transfer event raised by the COKC contract.
type COKCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_COKC *COKCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*COKCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &COKCTransferIterator{contract: _COKC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_COKC *COKCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *COKCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCTransfer)
				if err := _COKC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_COKC *COKCFilterer) ParseTransfer(log types.Log) (*COKCTransfer, error) {
	event := new(COKCTransfer)
	if err := _COKC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// COKCWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the COKC contract.
type COKCWithdrawnIterator struct {
	Event *COKCWithdrawn // Event containing the contract specifics and raw log

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
func (it *COKCWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(COKCWithdrawn)
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
		it.Event = new(COKCWithdrawn)
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
func (it *COKCWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *COKCWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// COKCWithdrawn represents a Withdrawn event raised by the COKC contract.
type COKCWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_COKC *COKCFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*COKCWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &COKCWithdrawnIterator{contract: _COKC.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_COKC *COKCFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *COKCWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _COKC.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(COKCWithdrawn)
				if err := _COKC.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_COKC *COKCFilterer) ParseWithdrawn(log types.Log) (*COKCWithdrawn, error) {
	event := new(COKCWithdrawn)
	if err := _COKC.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
