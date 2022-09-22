// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clockauction

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

// ClockauctionABI is the input ABI used to generate the binding from.
const ClockauctionABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenMaxOccurrences\",\"type\":\"uint256\"},{\"internalType\":\"contractIExchange\",\"name\":\"_exchangeContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ownerCut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_startingPrices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_endingPrices\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"contractIERC20[]\",\"name\":\"_exchangeTokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startingTimestamps\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_exchangeTokens\",\"type\":\"address\"}],\"name\":\"TokenAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"cancelTokenAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"enumIExchange.TokenType[]\",\"name\":\"_tokenTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokenAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenNumbers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_startingPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_endingPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_exchangeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_startingPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_endingPrices\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"_exchangeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_durations\",\"type\":\"uint256[]\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeContract\",\"outputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrices\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenNumber\",\"type\":\"uint256\"}],\"name\":\"getTokenAuctions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_sellers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_listingIndexes\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenNumber\",\"type\":\"uint256\"}],\"name\":\"getTokenAuctionsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"isAuctionExisting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"revalidateAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"revalidateRelatedAuctions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newOwnerCut\",\"type\":\"uint256\"}],\"name\":\"setOwnerCut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenMaxOccurrences\",\"type\":\"uint256\"}],\"name\":\"setTokenMaxOccurrences\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_listingIndex\",\"type\":\"uint256\"}],\"name\":\"settleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenMaxOccurrences\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIExchange\",\"name\":\"_exchangeContract\",\"type\":\"address\"}],\"name\":\"updateExchangeContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Clockauction is an auto generated Go binding around an Ethereum contract.
type Clockauction struct {
	ClockauctionCaller     // Read-only binding to the contract
	ClockauctionTransactor // Write-only binding to the contract
	ClockauctionFilterer   // Log filterer for contract events
}

// ClockauctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockauctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockauctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockauctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockauctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockauctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockauctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockauctionSession struct {
	Contract     *Clockauction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockauctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockauctionCallerSession struct {
	Contract *ClockauctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClockauctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockauctionTransactorSession struct {
	Contract     *ClockauctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClockauctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockauctionRaw struct {
	Contract *Clockauction // Generic contract binding to access the raw methods on
}

// ClockauctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockauctionCallerRaw struct {
	Contract *ClockauctionCaller // Generic read-only contract binding to access the raw methods on
}

// ClockauctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockauctionTransactorRaw struct {
	Contract *ClockauctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockauction creates a new instance of Clockauction, bound to a specific deployed contract.
func NewClockauction(address common.Address, backend bind.ContractBackend) (*Clockauction, error) {
	contract, err := bindClockauction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clockauction{ClockauctionCaller: ClockauctionCaller{contract: contract}, ClockauctionTransactor: ClockauctionTransactor{contract: contract}, ClockauctionFilterer: ClockauctionFilterer{contract: contract}}, nil
}

// NewClockauctionCaller creates a new read-only instance of Clockauction, bound to a specific deployed contract.
func NewClockauctionCaller(address common.Address, caller bind.ContractCaller) (*ClockauctionCaller, error) {
	contract, err := bindClockauction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockauctionCaller{contract: contract}, nil
}

// NewClockauctionTransactor creates a new write-only instance of Clockauction, bound to a specific deployed contract.
func NewClockauctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockauctionTransactor, error) {
	contract, err := bindClockauction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockauctionTransactor{contract: contract}, nil
}

// NewClockauctionFilterer creates a new log filterer instance of Clockauction, bound to a specific deployed contract.
func NewClockauctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockauctionFilterer, error) {
	contract, err := bindClockauction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockauctionFilterer{contract: contract}, nil
}

// bindClockauction binds a generic wrapper to an already deployed contract.
func bindClockauction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockauctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clockauction *ClockauctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clockauction.Contract.ClockauctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clockauction *ClockauctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.Contract.ClockauctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clockauction *ClockauctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clockauction.Contract.ClockauctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clockauction *ClockauctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clockauction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clockauction *ClockauctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clockauction *ClockauctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clockauction.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Clockauction *ClockauctionCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Clockauction *ClockauctionSession) Admin() (common.Address, error) {
	return _Clockauction.Contract.Admin(&_Clockauction.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Clockauction *ClockauctionCallerSession) Admin() (common.Address, error) {
	return _Clockauction.Contract.Admin(&_Clockauction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address seller)
func (_Clockauction *ClockauctionCaller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "auctions", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address seller)
func (_Clockauction *ClockauctionSession) Auctions(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Clockauction.Contract.Auctions(&_Clockauction.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address seller)
func (_Clockauction *ClockauctionCallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Clockauction.Contract.Auctions(&_Clockauction.CallOpts, arg0, arg1)
}

// ExchangeContract is a free data retrieval call binding the contract method 0x3f0a0797.
//
// Solidity: function exchangeContract() view returns(address)
func (_Clockauction *ClockauctionCaller) ExchangeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "exchangeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeContract is a free data retrieval call binding the contract method 0x3f0a0797.
//
// Solidity: function exchangeContract() view returns(address)
func (_Clockauction *ClockauctionSession) ExchangeContract() (common.Address, error) {
	return _Clockauction.Contract.ExchangeContract(&_Clockauction.CallOpts)
}

// ExchangeContract is a free data retrieval call binding the contract method 0x3f0a0797.
//
// Solidity: function exchangeContract() view returns(address)
func (_Clockauction *ClockauctionCallerSession) ExchangeContract() (common.Address, error) {
	return _Clockauction.Contract.ExchangeContract(&_Clockauction.CallOpts)
}

// GetCurrentPrices is a free data retrieval call binding the contract method 0x92ba8114.
//
// Solidity: function getCurrentPrices(address _seller, uint256 _listingIndex) view returns(address[], uint256[])
func (_Clockauction *ClockauctionCaller) GetCurrentPrices(opts *bind.CallOpts, _seller common.Address, _listingIndex *big.Int) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "getCurrentPrices", _seller, _listingIndex)

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetCurrentPrices is a free data retrieval call binding the contract method 0x92ba8114.
//
// Solidity: function getCurrentPrices(address _seller, uint256 _listingIndex) view returns(address[], uint256[])
func (_Clockauction *ClockauctionSession) GetCurrentPrices(_seller common.Address, _listingIndex *big.Int) ([]common.Address, []*big.Int, error) {
	return _Clockauction.Contract.GetCurrentPrices(&_Clockauction.CallOpts, _seller, _listingIndex)
}

// GetCurrentPrices is a free data retrieval call binding the contract method 0x92ba8114.
//
// Solidity: function getCurrentPrices(address _seller, uint256 _listingIndex) view returns(address[], uint256[])
func (_Clockauction *ClockauctionCallerSession) GetCurrentPrices(_seller common.Address, _listingIndex *big.Int) ([]common.Address, []*big.Int, error) {
	return _Clockauction.Contract.GetCurrentPrices(&_Clockauction.CallOpts, _seller, _listingIndex)
}

// GetTokenAuctions is a free data retrieval call binding the contract method 0xae500b7c.
//
// Solidity: function getTokenAuctions(address _tokenAddress, uint256 _tokenNumber) view returns(address[] _sellers, uint256[] _listingIndexes)
func (_Clockauction *ClockauctionCaller) GetTokenAuctions(opts *bind.CallOpts, _tokenAddress common.Address, _tokenNumber *big.Int) (struct {
	Sellers        []common.Address
	ListingIndexes []*big.Int
}, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "getTokenAuctions", _tokenAddress, _tokenNumber)

	outstruct := new(struct {
		Sellers        []common.Address
		ListingIndexes []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sellers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.ListingIndexes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTokenAuctions is a free data retrieval call binding the contract method 0xae500b7c.
//
// Solidity: function getTokenAuctions(address _tokenAddress, uint256 _tokenNumber) view returns(address[] _sellers, uint256[] _listingIndexes)
func (_Clockauction *ClockauctionSession) GetTokenAuctions(_tokenAddress common.Address, _tokenNumber *big.Int) (struct {
	Sellers        []common.Address
	ListingIndexes []*big.Int
}, error) {
	return _Clockauction.Contract.GetTokenAuctions(&_Clockauction.CallOpts, _tokenAddress, _tokenNumber)
}

// GetTokenAuctions is a free data retrieval call binding the contract method 0xae500b7c.
//
// Solidity: function getTokenAuctions(address _tokenAddress, uint256 _tokenNumber) view returns(address[] _sellers, uint256[] _listingIndexes)
func (_Clockauction *ClockauctionCallerSession) GetTokenAuctions(_tokenAddress common.Address, _tokenNumber *big.Int) (struct {
	Sellers        []common.Address
	ListingIndexes []*big.Int
}, error) {
	return _Clockauction.Contract.GetTokenAuctions(&_Clockauction.CallOpts, _tokenAddress, _tokenNumber)
}

// GetTokenAuctionsCount is a free data retrieval call binding the contract method 0x30ffb66c.
//
// Solidity: function getTokenAuctionsCount(address _tokenAddress, uint256 _tokenNumber) view returns(uint256)
func (_Clockauction *ClockauctionCaller) GetTokenAuctionsCount(opts *bind.CallOpts, _tokenAddress common.Address, _tokenNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "getTokenAuctionsCount", _tokenAddress, _tokenNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAuctionsCount is a free data retrieval call binding the contract method 0x30ffb66c.
//
// Solidity: function getTokenAuctionsCount(address _tokenAddress, uint256 _tokenNumber) view returns(uint256)
func (_Clockauction *ClockauctionSession) GetTokenAuctionsCount(_tokenAddress common.Address, _tokenNumber *big.Int) (*big.Int, error) {
	return _Clockauction.Contract.GetTokenAuctionsCount(&_Clockauction.CallOpts, _tokenAddress, _tokenNumber)
}

// GetTokenAuctionsCount is a free data retrieval call binding the contract method 0x30ffb66c.
//
// Solidity: function getTokenAuctionsCount(address _tokenAddress, uint256 _tokenNumber) view returns(uint256)
func (_Clockauction *ClockauctionCallerSession) GetTokenAuctionsCount(_tokenAddress common.Address, _tokenNumber *big.Int) (*big.Int, error) {
	return _Clockauction.Contract.GetTokenAuctionsCount(&_Clockauction.CallOpts, _tokenAddress, _tokenNumber)
}

// IsAuctionExisting is a free data retrieval call binding the contract method 0x44e290b2.
//
// Solidity: function isAuctionExisting(address _seller, uint256 _listingIndex) view returns(bool)
func (_Clockauction *ClockauctionCaller) IsAuctionExisting(opts *bind.CallOpts, _seller common.Address, _listingIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "isAuctionExisting", _seller, _listingIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionExisting is a free data retrieval call binding the contract method 0x44e290b2.
//
// Solidity: function isAuctionExisting(address _seller, uint256 _listingIndex) view returns(bool)
func (_Clockauction *ClockauctionSession) IsAuctionExisting(_seller common.Address, _listingIndex *big.Int) (bool, error) {
	return _Clockauction.Contract.IsAuctionExisting(&_Clockauction.CallOpts, _seller, _listingIndex)
}

// IsAuctionExisting is a free data retrieval call binding the contract method 0x44e290b2.
//
// Solidity: function isAuctionExisting(address _seller, uint256 _listingIndex) view returns(bool)
func (_Clockauction *ClockauctionCallerSession) IsAuctionExisting(_seller common.Address, _listingIndex *big.Int) (bool, error) {
	return _Clockauction.Contract.IsAuctionExisting(&_Clockauction.CallOpts, _seller, _listingIndex)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Clockauction *ClockauctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Clockauction *ClockauctionSession) OwnerCut() (*big.Int, error) {
	return _Clockauction.Contract.OwnerCut(&_Clockauction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_Clockauction *ClockauctionCallerSession) OwnerCut() (*big.Int, error) {
	return _Clockauction.Contract.OwnerCut(&_Clockauction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clockauction *ClockauctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clockauction *ClockauctionSession) Paused() (bool, error) {
	return _Clockauction.Contract.Paused(&_Clockauction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Clockauction *ClockauctionCallerSession) Paused() (bool, error) {
	return _Clockauction.Contract.Paused(&_Clockauction.CallOpts)
}

// TokenMaxOccurrences is a free data retrieval call binding the contract method 0x0622a388.
//
// Solidity: function tokenMaxOccurrences() view returns(uint256)
func (_Clockauction *ClockauctionCaller) TokenMaxOccurrences(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clockauction.contract.Call(opts, &out, "tokenMaxOccurrences")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenMaxOccurrences is a free data retrieval call binding the contract method 0x0622a388.
//
// Solidity: function tokenMaxOccurrences() view returns(uint256)
func (_Clockauction *ClockauctionSession) TokenMaxOccurrences() (*big.Int, error) {
	return _Clockauction.Contract.TokenMaxOccurrences(&_Clockauction.CallOpts)
}

// TokenMaxOccurrences is a free data retrieval call binding the contract method 0x0622a388.
//
// Solidity: function tokenMaxOccurrences() view returns(uint256)
func (_Clockauction *ClockauctionCallerSession) TokenMaxOccurrences() (*big.Int, error) {
	return _Clockauction.Contract.TokenMaxOccurrences(&_Clockauction.CallOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactor) CancelAuction(opts *bind.TransactOpts, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "cancelAuction", _listingIndex)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionSession) CancelAuction(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CancelAuction(&_Clockauction.TransactOpts, _listingIndex)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactorSession) CancelAuction(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CancelAuction(&_Clockauction.TransactOpts, _listingIndex)
}

// CancelTokenAuction is a paid mutator transaction binding the contract method 0xd10c7d1e.
//
// Solidity: function cancelTokenAuction(uint256 _listingIndex, address _token) returns()
func (_Clockauction *ClockauctionTransactor) CancelTokenAuction(opts *bind.TransactOpts, _listingIndex *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "cancelTokenAuction", _listingIndex, _token)
}

// CancelTokenAuction is a paid mutator transaction binding the contract method 0xd10c7d1e.
//
// Solidity: function cancelTokenAuction(uint256 _listingIndex, address _token) returns()
func (_Clockauction *ClockauctionSession) CancelTokenAuction(_listingIndex *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.CancelTokenAuction(&_Clockauction.TransactOpts, _listingIndex, _token)
}

// CancelTokenAuction is a paid mutator transaction binding the contract method 0xd10c7d1e.
//
// Solidity: function cancelTokenAuction(uint256 _listingIndex, address _token) returns()
func (_Clockauction *ClockauctionTransactorSession) CancelTokenAuction(_listingIndex *big.Int, _token common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.CancelTokenAuction(&_Clockauction.TransactOpts, _listingIndex, _token)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Clockauction *ClockauctionTransactor) ChangeAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "changeAdmin", _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Clockauction *ClockauctionSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.ChangeAdmin(&_Clockauction.TransactOpts, _newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _newAdmin) returns()
func (_Clockauction *ClockauctionTransactorSession) ChangeAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.ChangeAdmin(&_Clockauction.TransactOpts, _newAdmin)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x0b830218.
//
// Solidity: function createAuction(uint8[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenNumbers, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenTypes []uint8, _tokenAddresses []common.Address, _tokenNumbers []*big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "createAuction", _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x0b830218.
//
// Solidity: function createAuction(uint8[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenNumbers, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionSession) CreateAuction(_tokenTypes []uint8, _tokenAddresses []common.Address, _tokenNumbers []*big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CreateAuction(&_Clockauction.TransactOpts, _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x0b830218.
//
// Solidity: function createAuction(uint8[] _tokenTypes, address[] _tokenAddresses, uint256[] _tokenNumbers, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionTransactorSession) CreateAuction(_tokenTypes []uint8, _tokenAddresses []common.Address, _tokenNumbers []*big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CreateAuction(&_Clockauction.TransactOpts, _tokenTypes, _tokenAddresses, _tokenNumbers, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction0 is a paid mutator transaction binding the contract method 0x77ae0b83.
//
// Solidity: function createAuction(uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionTransactor) CreateAuction0(opts *bind.TransactOpts, _listingIndex *big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "createAuction0", _listingIndex, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction0 is a paid mutator transaction binding the contract method 0x77ae0b83.
//
// Solidity: function createAuction(uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionSession) CreateAuction0(_listingIndex *big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CreateAuction0(&_Clockauction.TransactOpts, _listingIndex, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// CreateAuction0 is a paid mutator transaction binding the contract method 0x77ae0b83.
//
// Solidity: function createAuction(uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations) returns()
func (_Clockauction *ClockauctionTransactorSession) CreateAuction0(_listingIndex *big.Int, _startingPrices []*big.Int, _endingPrices []*big.Int, _exchangeTokens []common.Address, _durations []*big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.CreateAuction0(&_Clockauction.TransactOpts, _listingIndex, _startingPrices, _endingPrices, _exchangeTokens, _durations)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clockauction *ClockauctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clockauction *ClockauctionSession) Pause() (*types.Transaction, error) {
	return _Clockauction.Contract.Pause(&_Clockauction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Clockauction *ClockauctionTransactorSession) Pause() (*types.Transaction, error) {
	return _Clockauction.Contract.Pause(&_Clockauction.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Clockauction *ClockauctionTransactor) RemoveAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "removeAdmin")
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Clockauction *ClockauctionSession) RemoveAdmin() (*types.Transaction, error) {
	return _Clockauction.Contract.RemoveAdmin(&_Clockauction.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x9a202d47.
//
// Solidity: function removeAdmin() returns()
func (_Clockauction *ClockauctionTransactorSession) RemoveAdmin() (*types.Transaction, error) {
	return _Clockauction.Contract.RemoveAdmin(&_Clockauction.TransactOpts)
}

// RevalidateAuction is a paid mutator transaction binding the contract method 0x45496ed9.
//
// Solidity: function revalidateAuction(address _seller, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactor) RevalidateAuction(opts *bind.TransactOpts, _seller common.Address, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "revalidateAuction", _seller, _listingIndex)
}

// RevalidateAuction is a paid mutator transaction binding the contract method 0x45496ed9.
//
// Solidity: function revalidateAuction(address _seller, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionSession) RevalidateAuction(_seller common.Address, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.RevalidateAuction(&_Clockauction.TransactOpts, _seller, _listingIndex)
}

// RevalidateAuction is a paid mutator transaction binding the contract method 0x45496ed9.
//
// Solidity: function revalidateAuction(address _seller, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactorSession) RevalidateAuction(_seller common.Address, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.RevalidateAuction(&_Clockauction.TransactOpts, _seller, _listingIndex)
}

// RevalidateRelatedAuctions is a paid mutator transaction binding the contract method 0x3bd63ba2.
//
// Solidity: function revalidateRelatedAuctions(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactor) RevalidateRelatedAuctions(opts *bind.TransactOpts, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "revalidateRelatedAuctions", _listingIndex)
}

// RevalidateRelatedAuctions is a paid mutator transaction binding the contract method 0x3bd63ba2.
//
// Solidity: function revalidateRelatedAuctions(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionSession) RevalidateRelatedAuctions(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.RevalidateRelatedAuctions(&_Clockauction.TransactOpts, _listingIndex)
}

// RevalidateRelatedAuctions is a paid mutator transaction binding the contract method 0x3bd63ba2.
//
// Solidity: function revalidateRelatedAuctions(uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactorSession) RevalidateRelatedAuctions(_listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.RevalidateRelatedAuctions(&_Clockauction.TransactOpts, _listingIndex)
}

// SetOwnerCut is a paid mutator transaction binding the contract method 0x757de573.
//
// Solidity: function setOwnerCut(uint256 _newOwnerCut) returns()
func (_Clockauction *ClockauctionTransactor) SetOwnerCut(opts *bind.TransactOpts, _newOwnerCut *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "setOwnerCut", _newOwnerCut)
}

// SetOwnerCut is a paid mutator transaction binding the contract method 0x757de573.
//
// Solidity: function setOwnerCut(uint256 _newOwnerCut) returns()
func (_Clockauction *ClockauctionSession) SetOwnerCut(_newOwnerCut *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SetOwnerCut(&_Clockauction.TransactOpts, _newOwnerCut)
}

// SetOwnerCut is a paid mutator transaction binding the contract method 0x757de573.
//
// Solidity: function setOwnerCut(uint256 _newOwnerCut) returns()
func (_Clockauction *ClockauctionTransactorSession) SetOwnerCut(_newOwnerCut *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SetOwnerCut(&_Clockauction.TransactOpts, _newOwnerCut)
}

// SetTokenMaxOccurrences is a paid mutator transaction binding the contract method 0xd67eb0b5.
//
// Solidity: function setTokenMaxOccurrences(uint256 _tokenMaxOccurrences) returns()
func (_Clockauction *ClockauctionTransactor) SetTokenMaxOccurrences(opts *bind.TransactOpts, _tokenMaxOccurrences *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "setTokenMaxOccurrences", _tokenMaxOccurrences)
}

// SetTokenMaxOccurrences is a paid mutator transaction binding the contract method 0xd67eb0b5.
//
// Solidity: function setTokenMaxOccurrences(uint256 _tokenMaxOccurrences) returns()
func (_Clockauction *ClockauctionSession) SetTokenMaxOccurrences(_tokenMaxOccurrences *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SetTokenMaxOccurrences(&_Clockauction.TransactOpts, _tokenMaxOccurrences)
}

// SetTokenMaxOccurrences is a paid mutator transaction binding the contract method 0xd67eb0b5.
//
// Solidity: function setTokenMaxOccurrences(uint256 _tokenMaxOccurrences) returns()
func (_Clockauction *ClockauctionTransactorSession) SetTokenMaxOccurrences(_tokenMaxOccurrences *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SetTokenMaxOccurrences(&_Clockauction.TransactOpts, _tokenMaxOccurrences)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xc354be6d.
//
// Solidity: function settleAuction(address _seller, address _token, uint256 _bidAmount, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactor) SettleAuction(opts *bind.TransactOpts, _seller common.Address, _token common.Address, _bidAmount *big.Int, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "settleAuction", _seller, _token, _bidAmount, _listingIndex)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xc354be6d.
//
// Solidity: function settleAuction(address _seller, address _token, uint256 _bidAmount, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionSession) SettleAuction(_seller common.Address, _token common.Address, _bidAmount *big.Int, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SettleAuction(&_Clockauction.TransactOpts, _seller, _token, _bidAmount, _listingIndex)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xc354be6d.
//
// Solidity: function settleAuction(address _seller, address _token, uint256 _bidAmount, uint256 _listingIndex) returns()
func (_Clockauction *ClockauctionTransactorSession) SettleAuction(_seller common.Address, _token common.Address, _bidAmount *big.Int, _listingIndex *big.Int) (*types.Transaction, error) {
	return _Clockauction.Contract.SettleAuction(&_Clockauction.TransactOpts, _seller, _token, _bidAmount, _listingIndex)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clockauction *ClockauctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clockauction *ClockauctionSession) Unpause() (*types.Transaction, error) {
	return _Clockauction.Contract.Unpause(&_Clockauction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Clockauction *ClockauctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _Clockauction.Contract.Unpause(&_Clockauction.TransactOpts)
}

// UpdateExchangeContract is a paid mutator transaction binding the contract method 0x75640815.
//
// Solidity: function updateExchangeContract(address _exchangeContract) returns()
func (_Clockauction *ClockauctionTransactor) UpdateExchangeContract(opts *bind.TransactOpts, _exchangeContract common.Address) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "updateExchangeContract", _exchangeContract)
}

// UpdateExchangeContract is a paid mutator transaction binding the contract method 0x75640815.
//
// Solidity: function updateExchangeContract(address _exchangeContract) returns()
func (_Clockauction *ClockauctionSession) UpdateExchangeContract(_exchangeContract common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.UpdateExchangeContract(&_Clockauction.TransactOpts, _exchangeContract)
}

// UpdateExchangeContract is a paid mutator transaction binding the contract method 0x75640815.
//
// Solidity: function updateExchangeContract(address _exchangeContract) returns()
func (_Clockauction *ClockauctionTransactorSession) UpdateExchangeContract(_exchangeContract common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.UpdateExchangeContract(&_Clockauction.TransactOpts, _exchangeContract)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Clockauction *ClockauctionTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Clockauction *ClockauctionSession) WithdrawEther() (*types.Transaction, error) {
	return _Clockauction.Contract.WithdrawEther(&_Clockauction.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_Clockauction *ClockauctionTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _Clockauction.Contract.WithdrawEther(&_Clockauction.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Clockauction *ClockauctionTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Clockauction.contract.Transact(opts, "withdrawToken", _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Clockauction *ClockauctionSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.WithdrawToken(&_Clockauction.TransactOpts, _token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _token) returns()
func (_Clockauction *ClockauctionTransactorSession) WithdrawToken(_token common.Address) (*types.Transaction, error) {
	return _Clockauction.Contract.WithdrawToken(&_Clockauction.TransactOpts, _token)
}

// ClockauctionAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Clockauction contract.
type ClockauctionAdminChangedIterator struct {
	Event *ClockauctionAdminChanged // Event containing the contract specifics and raw log

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
func (it *ClockauctionAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionAdminChanged)
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
		it.Event = new(ClockauctionAdminChanged)
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
func (it *ClockauctionAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionAdminChanged represents a AdminChanged event raised by the Clockauction contract.
type ClockauctionAdminChanged struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Clockauction *ClockauctionFilterer) FilterAdminChanged(opts *bind.FilterOpts, _oldAdmin []common.Address, _newAdmin []common.Address) (*ClockauctionAdminChangedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return &ClockauctionAdminChangedIterator{contract: _Clockauction.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address indexed _oldAdmin, address indexed _newAdmin)
func (_Clockauction *ClockauctionFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ClockauctionAdminChanged, _oldAdmin []common.Address, _newAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}
	var _newAdminRule []interface{}
	for _, _newAdminItem := range _newAdmin {
		_newAdminRule = append(_newAdminRule, _newAdminItem)
	}

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "AdminChanged", _oldAdminRule, _newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionAdminChanged)
				if err := _Clockauction.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Clockauction *ClockauctionFilterer) ParseAdminChanged(log types.Log) (*ClockauctionAdminChanged, error) {
	event := new(ClockauctionAdminChanged)
	if err := _Clockauction.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the Clockauction contract.
type ClockauctionAdminRemovedIterator struct {
	Event *ClockauctionAdminRemoved // Event containing the contract specifics and raw log

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
func (it *ClockauctionAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionAdminRemoved)
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
		it.Event = new(ClockauctionAdminRemoved)
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
func (it *ClockauctionAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionAdminRemoved represents a AdminRemoved event raised by the Clockauction contract.
type ClockauctionAdminRemoved struct {
	OldAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Clockauction *ClockauctionFilterer) FilterAdminRemoved(opts *bind.FilterOpts, _oldAdmin []common.Address) (*ClockauctionAdminRemovedIterator, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return &ClockauctionAdminRemovedIterator{contract: _Clockauction.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed _oldAdmin)
func (_Clockauction *ClockauctionFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *ClockauctionAdminRemoved, _oldAdmin []common.Address) (event.Subscription, error) {

	var _oldAdminRule []interface{}
	for _, _oldAdminItem := range _oldAdmin {
		_oldAdminRule = append(_oldAdminRule, _oldAdminItem)
	}

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "AdminRemoved", _oldAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionAdminRemoved)
				if err := _Clockauction.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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
func (_Clockauction *ClockauctionFilterer) ParseAdminRemoved(log types.Log) (*ClockauctionAdminRemoved, error) {
	event := new(ClockauctionAdminRemoved)
	if err := _Clockauction.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the Clockauction contract.
type ClockauctionAuctionCancelledIterator struct {
	Event *ClockauctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockauctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionAuctionCancelled)
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
		it.Event = new(ClockauctionAuctionCancelled)
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
func (it *ClockauctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionAuctionCancelled represents a AuctionCancelled event raised by the Clockauction contract.
type ClockauctionAuctionCancelled struct {
	Seller       common.Address
	ListingIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address _seller, uint256 _listingIndex)
func (_Clockauction *ClockauctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockauctionAuctionCancelledIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockauctionAuctionCancelledIterator{contract: _Clockauction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address _seller, uint256 _listingIndex)
func (_Clockauction *ClockauctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockauctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionAuctionCancelled)
				if err := _Clockauction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x018b64b6242d32aa550e95d78985b938d71af5b3f10827b0683f55da16393048.
//
// Solidity: event AuctionCancelled(address _seller, uint256 _listingIndex)
func (_Clockauction *ClockauctionFilterer) ParseAuctionCancelled(log types.Log) (*ClockauctionAuctionCancelled, error) {
	event := new(ClockauctionAuctionCancelled)
	if err := _Clockauction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Clockauction contract.
type ClockauctionAuctionCreatedIterator struct {
	Event *ClockauctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockauctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionAuctionCreated)
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
		it.Event = new(ClockauctionAuctionCreated)
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
func (it *ClockauctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionAuctionCreated represents a AuctionCreated event raised by the Clockauction contract.
type ClockauctionAuctionCreated struct {
	Seller             common.Address
	ListingIndex       *big.Int
	StartingPrices     []*big.Int
	EndingPrices       []*big.Int
	ExchangeTokens     []common.Address
	Durations          []*big.Int
	StartingTimestamps *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xae3392a96856e8c1881402157f65e69336cb9e04ffba578babad5b29909def82.
//
// Solidity: event AuctionCreated(address _seller, uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations, uint256 _startingTimestamps)
func (_Clockauction *ClockauctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockauctionAuctionCreatedIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockauctionAuctionCreatedIterator{contract: _Clockauction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xae3392a96856e8c1881402157f65e69336cb9e04ffba578babad5b29909def82.
//
// Solidity: event AuctionCreated(address _seller, uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations, uint256 _startingTimestamps)
func (_Clockauction *ClockauctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockauctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionAuctionCreated)
				if err := _Clockauction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xae3392a96856e8c1881402157f65e69336cb9e04ffba578babad5b29909def82.
//
// Solidity: event AuctionCreated(address _seller, uint256 _listingIndex, uint256[] _startingPrices, uint256[] _endingPrices, address[] _exchangeTokens, uint256[] _durations, uint256 _startingTimestamps)
func (_Clockauction *ClockauctionFilterer) ParseAuctionCreated(log types.Log) (*ClockauctionAuctionCreated, error) {
	event := new(ClockauctionAuctionCreated)
	if err := _Clockauction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the Clockauction contract.
type ClockauctionAuctionSuccessfulIterator struct {
	Event *ClockauctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockauctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionAuctionSuccessful)
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
		it.Event = new(ClockauctionAuctionSuccessful)
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
func (it *ClockauctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionAuctionSuccessful represents a AuctionSuccessful event raised by the Clockauction contract.
type ClockauctionAuctionSuccessful struct {
	Seller       common.Address
	Buyer        common.Address
	ListingIndex *big.Int
	Token        common.Address
	TotalPrice   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x0c0258cd7f0d9474f62106c6981c027ea54bee0b323ea1991f4caa7e288a5725.
//
// Solidity: event AuctionSuccessful(address _seller, address _buyer, uint256 _listingIndex, address _token, uint256 _totalPrice)
func (_Clockauction *ClockauctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockauctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockauctionAuctionSuccessfulIterator{contract: _Clockauction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x0c0258cd7f0d9474f62106c6981c027ea54bee0b323ea1991f4caa7e288a5725.
//
// Solidity: event AuctionSuccessful(address _seller, address _buyer, uint256 _listingIndex, address _token, uint256 _totalPrice)
func (_Clockauction *ClockauctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockauctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionAuctionSuccessful)
				if err := _Clockauction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x0c0258cd7f0d9474f62106c6981c027ea54bee0b323ea1991f4caa7e288a5725.
//
// Solidity: event AuctionSuccessful(address _seller, address _buyer, uint256 _listingIndex, address _token, uint256 _totalPrice)
func (_Clockauction *ClockauctionFilterer) ParseAuctionSuccessful(log types.Log) (*ClockauctionAuctionSuccessful, error) {
	event := new(ClockauctionAuctionSuccessful)
	if err := _Clockauction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Clockauction contract.
type ClockauctionPausedIterator struct {
	Event *ClockauctionPaused // Event containing the contract specifics and raw log

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
func (it *ClockauctionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionPaused)
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
		it.Event = new(ClockauctionPaused)
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
func (it *ClockauctionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionPaused represents a Paused event raised by the Clockauction contract.
type ClockauctionPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Clockauction *ClockauctionFilterer) FilterPaused(opts *bind.FilterOpts) (*ClockauctionPausedIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ClockauctionPausedIterator{contract: _Clockauction.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_Clockauction *ClockauctionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ClockauctionPaused) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionPaused)
				if err := _Clockauction.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Clockauction *ClockauctionFilterer) ParsePaused(log types.Log) (*ClockauctionPaused, error) {
	event := new(ClockauctionPaused)
	if err := _Clockauction.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionTokenAuctionCancelledIterator is returned from FilterTokenAuctionCancelled and is used to iterate over the raw logs and unpacked data for TokenAuctionCancelled events raised by the Clockauction contract.
type ClockauctionTokenAuctionCancelledIterator struct {
	Event *ClockauctionTokenAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockauctionTokenAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionTokenAuctionCancelled)
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
		it.Event = new(ClockauctionTokenAuctionCancelled)
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
func (it *ClockauctionTokenAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionTokenAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionTokenAuctionCancelled represents a TokenAuctionCancelled event raised by the Clockauction contract.
type ClockauctionTokenAuctionCancelled struct {
	Seller         common.Address
	ListingIndex   *big.Int
	ExchangeTokens common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenAuctionCancelled is a free log retrieval operation binding the contract event 0x6399a705451a5f1ca57a7f6e1596b475ed0dc9b2d6304c5adbf075ed83283d63.
//
// Solidity: event TokenAuctionCancelled(address _seller, uint256 _listingIndex, address _exchangeTokens)
func (_Clockauction *ClockauctionFilterer) FilterTokenAuctionCancelled(opts *bind.FilterOpts) (*ClockauctionTokenAuctionCancelledIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "TokenAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockauctionTokenAuctionCancelledIterator{contract: _Clockauction.contract, event: "TokenAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchTokenAuctionCancelled is a free log subscription operation binding the contract event 0x6399a705451a5f1ca57a7f6e1596b475ed0dc9b2d6304c5adbf075ed83283d63.
//
// Solidity: event TokenAuctionCancelled(address _seller, uint256 _listingIndex, address _exchangeTokens)
func (_Clockauction *ClockauctionFilterer) WatchTokenAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockauctionTokenAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "TokenAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionTokenAuctionCancelled)
				if err := _Clockauction.contract.UnpackLog(event, "TokenAuctionCancelled", log); err != nil {
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

// ParseTokenAuctionCancelled is a log parse operation binding the contract event 0x6399a705451a5f1ca57a7f6e1596b475ed0dc9b2d6304c5adbf075ed83283d63.
//
// Solidity: event TokenAuctionCancelled(address _seller, uint256 _listingIndex, address _exchangeTokens)
func (_Clockauction *ClockauctionFilterer) ParseTokenAuctionCancelled(log types.Log) (*ClockauctionTokenAuctionCancelled, error) {
	event := new(ClockauctionTokenAuctionCancelled)
	if err := _Clockauction.contract.UnpackLog(event, "TokenAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockauctionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Clockauction contract.
type ClockauctionUnpausedIterator struct {
	Event *ClockauctionUnpaused // Event containing the contract specifics and raw log

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
func (it *ClockauctionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockauctionUnpaused)
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
		it.Event = new(ClockauctionUnpaused)
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
func (it *ClockauctionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockauctionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockauctionUnpaused represents a Unpaused event raised by the Clockauction contract.
type ClockauctionUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Clockauction *ClockauctionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ClockauctionUnpausedIterator, error) {

	logs, sub, err := _Clockauction.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ClockauctionUnpausedIterator{contract: _Clockauction.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_Clockauction *ClockauctionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ClockauctionUnpaused) (event.Subscription, error) {

	logs, sub, err := _Clockauction.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockauctionUnpaused)
				if err := _Clockauction.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Clockauction *ClockauctionFilterer) ParseUnpaused(log types.Log) (*ClockauctionUnpaused, error) {
	event := new(ClockauctionUnpaused)
	if err := _Clockauction.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
