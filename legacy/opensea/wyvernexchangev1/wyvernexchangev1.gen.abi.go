// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wyvernexchangev1

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

// Wyvernexchangev1MetaData contains all meta data concerning the Wyvernexchangev1 contract.
var Wyvernexchangev1MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenTransferProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"extradata\",\"type\":\"bytes\"}],\"name\":\"staticCall\",\"outputs\":[{\"name\":\"result\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMinimumMakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumMakerProtocolFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMinimumTakerProtocolFee\",\"type\":\"uint256\"}],\"name\":\"changeMinimumTakerProtocolFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"array\",\"type\":\"bytes\"},{\"name\":\"desired\",\"type\":\"bytes\"},{\"name\":\"mask\",\"type\":\"bytes\"}],\"name\":\"guardedArrayReplace\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumTakerProtocolFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"codename\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"testCopyAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"arrToCopy\",\"type\":\"bytes\"}],\"name\":\"testCopy\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"calculateCurrentPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"changeProtocolFeeRecipient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"buyCalldata\",\"type\":\"bytes\"},{\"name\":\"buyReplacementPattern\",\"type\":\"bytes\"},{\"name\":\"sellCalldata\",\"type\":\"bytes\"},{\"name\":\"sellReplacementPattern\",\"type\":\"bytes\"}],\"name\":\"orderCalldataCanMatch\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"validateOrder_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"basePrice\",\"type\":\"uint256\"},{\"name\":\"extra\",\"type\":\"uint256\"},{\"name\":\"listingTime\",\"type\":\"uint256\"},{\"name\":\"expirationTime\",\"type\":\"uint256\"}],\"name\":\"calculateFinalPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"hashOrder_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"ordersCanMatch_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"approveOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumMakerProtocolFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"hashToSign_\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFinalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder_\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"},{\"name\":\"vs\",\"type\":\"uint8[2]\"},{\"name\":\"rssMetadata\",\"type\":\"bytes32[5]\"}],\"name\":\"atomicMatch_\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[7]\"},{\"name\":\"uints\",\"type\":\"uint256[9]\"},{\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"name\":\"side\",\"type\":\"uint8\"},{\"name\":\"saleKind\",\"type\":\"uint8\"},{\"name\":\"howToCall\",\"type\":\"uint8\"},{\"name\":\"calldata\",\"type\":\"bytes\"},{\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"name\":\"staticExtradata\",\"type\":\"bytes\"}],\"name\":\"validateOrderParameters_\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[14]\"},{\"name\":\"uints\",\"type\":\"uint256[18]\"},{\"name\":\"feeMethodsSidesKindsHowToCalls\",\"type\":\"uint8[8]\"},{\"name\":\"calldataBuy\",\"type\":\"bytes\"},{\"name\":\"calldataSell\",\"type\":\"bytes\"},{\"name\":\"replacementPatternBuy\",\"type\":\"bytes\"},{\"name\":\"replacementPatternSell\",\"type\":\"bytes\"},{\"name\":\"staticExtradataBuy\",\"type\":\"bytes\"},{\"name\":\"staticExtradataSell\",\"type\":\"bytes\"}],\"name\":\"calculateMatchPrice_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedOrders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"tokenTransferProxyAddress\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"protocolFeeAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"exchange\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerRelayerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"makerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"takerProtocolFee\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"feeMethod\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"saleKind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"target\",\"type\":\"address\"}],\"name\":\"OrderApprovedPartOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"howToCall\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"calldata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"replacementPattern\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"staticTarget\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"staticExtradata\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"basePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"extra\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"listingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"orderbookInclusionDesired\",\"type\":\"bool\"}],\"name\":\"OrderApprovedPartTwo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"buyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
}

// Wyvernexchangev1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Wyvernexchangev1MetaData.ABI instead.
var Wyvernexchangev1ABI = Wyvernexchangev1MetaData.ABI

// Wyvernexchangev1 is an auto generated Go binding around an Ethereum contract.
type Wyvernexchangev1 struct {
	Wyvernexchangev1Caller     // Read-only binding to the contract
	Wyvernexchangev1Transactor // Write-only binding to the contract
	Wyvernexchangev1Filterer   // Log filterer for contract events
}

// Wyvernexchangev1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Wyvernexchangev1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Wyvernexchangev1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Wyvernexchangev1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Wyvernexchangev1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Wyvernexchangev1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Wyvernexchangev1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Wyvernexchangev1Session struct {
	Contract     *Wyvernexchangev1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Wyvernexchangev1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Wyvernexchangev1CallerSession struct {
	Contract *Wyvernexchangev1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Wyvernexchangev1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Wyvernexchangev1TransactorSession struct {
	Contract     *Wyvernexchangev1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Wyvernexchangev1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Wyvernexchangev1Raw struct {
	Contract *Wyvernexchangev1 // Generic contract binding to access the raw methods on
}

// Wyvernexchangev1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Wyvernexchangev1CallerRaw struct {
	Contract *Wyvernexchangev1Caller // Generic read-only contract binding to access the raw methods on
}

// Wyvernexchangev1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Wyvernexchangev1TransactorRaw struct {
	Contract *Wyvernexchangev1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWyvernexchangev1 creates a new instance of Wyvernexchangev1, bound to a specific deployed contract.
func NewWyvernexchangev1(address common.Address, backend bind.ContractBackend) (*Wyvernexchangev1, error) {
	contract, err := bindWyvernexchangev1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1{Wyvernexchangev1Caller: Wyvernexchangev1Caller{contract: contract}, Wyvernexchangev1Transactor: Wyvernexchangev1Transactor{contract: contract}, Wyvernexchangev1Filterer: Wyvernexchangev1Filterer{contract: contract}}, nil
}

// NewWyvernexchangev1Caller creates a new read-only instance of Wyvernexchangev1, bound to a specific deployed contract.
func NewWyvernexchangev1Caller(address common.Address, caller bind.ContractCaller) (*Wyvernexchangev1Caller, error) {
	contract, err := bindWyvernexchangev1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1Caller{contract: contract}, nil
}

// NewWyvernexchangev1Transactor creates a new write-only instance of Wyvernexchangev1, bound to a specific deployed contract.
func NewWyvernexchangev1Transactor(address common.Address, transactor bind.ContractTransactor) (*Wyvernexchangev1Transactor, error) {
	contract, err := bindWyvernexchangev1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1Transactor{contract: contract}, nil
}

// NewWyvernexchangev1Filterer creates a new log filterer instance of Wyvernexchangev1, bound to a specific deployed contract.
func NewWyvernexchangev1Filterer(address common.Address, filterer bind.ContractFilterer) (*Wyvernexchangev1Filterer, error) {
	contract, err := bindWyvernexchangev1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1Filterer{contract: contract}, nil
}

// bindWyvernexchangev1 binds a generic wrapper to an already deployed contract.
func bindWyvernexchangev1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Wyvernexchangev1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wyvernexchangev1 *Wyvernexchangev1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wyvernexchangev1.Contract.Wyvernexchangev1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wyvernexchangev1 *Wyvernexchangev1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.Wyvernexchangev1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wyvernexchangev1 *Wyvernexchangev1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.Wyvernexchangev1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wyvernexchangev1 *Wyvernexchangev1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wyvernexchangev1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.contract.Transact(opts, method, params...)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) INVERSEBASISPOINT() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.INVERSEBASISPOINT(&_Wyvernexchangev1.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.INVERSEBASISPOINT(&_Wyvernexchangev1.CallOpts)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) ApprovedOrders(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "approvedOrders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ApprovedOrders(&_Wyvernexchangev1.CallOpts, arg0)
}

// ApprovedOrders is a free data retrieval call binding the contract method 0xe57d4adb.
//
// Solidity: function approvedOrders(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) ApprovedOrders(arg0 [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ApprovedOrders(&_Wyvernexchangev1.CallOpts, arg0)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) CalculateCurrentPrice(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "calculateCurrentPrice_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) CalculateCurrentPrice(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateCurrentPrice(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0x3f67ee0d.
//
// Solidity: function calculateCurrentPrice_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) CalculateCurrentPrice(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateCurrentPrice(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) CalculateFinalPrice(opts *bind.CallOpts, side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "calculateFinalPrice", side, saleKind, basePrice, extra, listingTime, expirationTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateFinalPrice(&_Wyvernexchangev1.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateFinalPrice is a free data retrieval call binding the contract method 0x63d36c0b.
//
// Solidity: function calculateFinalPrice(uint8 side, uint8 saleKind, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) CalculateFinalPrice(side uint8, saleKind uint8, basePrice *big.Int, extra *big.Int, listingTime *big.Int, expirationTime *big.Int) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateFinalPrice(&_Wyvernexchangev1.CallOpts, side, saleKind, basePrice, extra, listingTime, expirationTime)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) CalculateMatchPrice(opts *bind.CallOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "calculateMatchPrice_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) CalculateMatchPrice(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateMatchPrice(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CalculateMatchPrice is a free data retrieval call binding the contract method 0xd537e131.
//
// Solidity: function calculateMatchPrice_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) CalculateMatchPrice(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (*big.Int, error) {
	return _Wyvernexchangev1.Contract.CalculateMatchPrice(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) CancelledOrFinalized(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "cancelledOrFinalized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.CancelledOrFinalized(&_Wyvernexchangev1.CallOpts, arg0)
}

// CancelledOrFinalized is a free data retrieval call binding the contract method 0x8076f005.
//
// Solidity: function cancelledOrFinalized(bytes32 ) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) CancelledOrFinalized(arg0 [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.CancelledOrFinalized(&_Wyvernexchangev1.CallOpts, arg0)
}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) Codename(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "codename")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) Codename() (string, error) {
	return _Wyvernexchangev1.Contract.Codename(&_Wyvernexchangev1.CallOpts)
}

// Codename is a free data retrieval call binding the contract method 0x31e63199.
//
// Solidity: function codename() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) Codename() (string, error) {
	return _Wyvernexchangev1.Contract.Codename(&_Wyvernexchangev1.CallOpts)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) ExchangeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "exchangeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ExchangeToken() (common.Address, error) {
	return _Wyvernexchangev1.Contract.ExchangeToken(&_Wyvernexchangev1.CallOpts)
}

// ExchangeToken is a free data retrieval call binding the contract method 0xa25eb5d9.
//
// Solidity: function exchangeToken() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) ExchangeToken() (common.Address, error) {
	return _Wyvernexchangev1.Contract.ExchangeToken(&_Wyvernexchangev1.CallOpts)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) GuardedArrayReplace(opts *bind.CallOpts, array []byte, desired []byte, mask []byte) ([]byte, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "guardedArrayReplace", array, desired, mask)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _Wyvernexchangev1.Contract.GuardedArrayReplace(&_Wyvernexchangev1.CallOpts, array, desired, mask)
}

// GuardedArrayReplace is a free data retrieval call binding the contract method 0x239e83df.
//
// Solidity: function guardedArrayReplace(bytes array, bytes desired, bytes mask) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) GuardedArrayReplace(array []byte, desired []byte, mask []byte) ([]byte, error) {
	return _Wyvernexchangev1.Contract.GuardedArrayReplace(&_Wyvernexchangev1.CallOpts, array, desired, mask)
}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) HashOrder(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "hashOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) HashOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _Wyvernexchangev1.Contract.HashOrder(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashOrder is a free data retrieval call binding the contract method 0x71d02b38.
//
// Solidity: function hashOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) HashOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _Wyvernexchangev1.Contract.HashOrder(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) HashToSign(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "hashToSign_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) HashToSign(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _Wyvernexchangev1.Contract.HashToSign(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// HashToSign is a free data retrieval call binding the contract method 0x7d766981.
//
// Solidity: function hashToSign_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) pure returns(bytes32)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) HashToSign(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) ([32]byte, error) {
	return _Wyvernexchangev1.Contract.HashToSign(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) MinimumMakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "minimumMakerProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) MinimumMakerProtocolFee() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.MinimumMakerProtocolFee(&_Wyvernexchangev1.CallOpts)
}

// MinimumMakerProtocolFee is a free data retrieval call binding the contract method 0x7ccefc52.
//
// Solidity: function minimumMakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) MinimumMakerProtocolFee() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.MinimumMakerProtocolFee(&_Wyvernexchangev1.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) MinimumTakerProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "minimumTakerProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) MinimumTakerProtocolFee() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.MinimumTakerProtocolFee(&_Wyvernexchangev1.CallOpts)
}

// MinimumTakerProtocolFee is a free data retrieval call binding the contract method 0x28a8ee68.
//
// Solidity: function minimumTakerProtocolFee() view returns(uint256)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) MinimumTakerProtocolFee() (*big.Int, error) {
	return _Wyvernexchangev1.Contract.MinimumTakerProtocolFee(&_Wyvernexchangev1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) Name() (string, error) {
	return _Wyvernexchangev1.Contract.Name(&_Wyvernexchangev1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) Name() (string, error) {
	return _Wyvernexchangev1.Contract.Name(&_Wyvernexchangev1.CallOpts)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) OrderCalldataCanMatch(opts *bind.CallOpts, buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "orderCalldataCanMatch", buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.OrderCalldataCanMatch(&_Wyvernexchangev1.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrderCalldataCanMatch is a free data retrieval call binding the contract method 0x562b2ebc.
//
// Solidity: function orderCalldataCanMatch(bytes buyCalldata, bytes buyReplacementPattern, bytes sellCalldata, bytes sellReplacementPattern) pure returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) OrderCalldataCanMatch(buyCalldata []byte, buyReplacementPattern []byte, sellCalldata []byte, sellReplacementPattern []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.OrderCalldataCanMatch(&_Wyvernexchangev1.CallOpts, buyCalldata, buyReplacementPattern, sellCalldata, sellReplacementPattern)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) OrdersCanMatch(opts *bind.CallOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "ordersCanMatch_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) OrdersCanMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.OrdersCanMatch(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// OrdersCanMatch is a free data retrieval call binding the contract method 0x72593b4c.
//
// Solidity: function ordersCanMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) OrdersCanMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.OrdersCanMatch(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) Owner() (common.Address, error) {
	return _Wyvernexchangev1.Contract.Owner(&_Wyvernexchangev1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) Owner() (common.Address, error) {
	return _Wyvernexchangev1.Contract.Owner(&_Wyvernexchangev1.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ProtocolFeeRecipient() (common.Address, error) {
	return _Wyvernexchangev1.Contract.ProtocolFeeRecipient(&_Wyvernexchangev1.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Wyvernexchangev1.Contract.ProtocolFeeRecipient(&_Wyvernexchangev1.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) Registry() (common.Address, error) {
	return _Wyvernexchangev1.Contract.Registry(&_Wyvernexchangev1.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) Registry() (common.Address, error) {
	return _Wyvernexchangev1.Contract.Registry(&_Wyvernexchangev1.CallOpts)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) StaticCall(opts *bind.CallOpts, target common.Address, calldata []byte, extradata []byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "staticCall", target, calldata, extradata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.StaticCall(&_Wyvernexchangev1.CallOpts, target, calldata, extradata)
}

// StaticCall is a free data retrieval call binding the contract method 0x10796a47.
//
// Solidity: function staticCall(address target, bytes calldata, bytes extradata) view returns(bool result)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) StaticCall(target common.Address, calldata []byte, extradata []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.StaticCall(&_Wyvernexchangev1.CallOpts, target, calldata, extradata)
}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) TestCopy(opts *bind.CallOpts, arrToCopy []byte) ([]byte, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "testCopy", arrToCopy)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) TestCopy(arrToCopy []byte) ([]byte, error) {
	return _Wyvernexchangev1.Contract.TestCopy(&_Wyvernexchangev1.CallOpts, arrToCopy)
}

// TestCopy is a free data retrieval call binding the contract method 0x3e1e292a.
//
// Solidity: function testCopy(bytes arrToCopy) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) TestCopy(arrToCopy []byte) ([]byte, error) {
	return _Wyvernexchangev1.Contract.TestCopy(&_Wyvernexchangev1.CallOpts, arrToCopy)
}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) TestCopyAddress(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "testCopyAddress", addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) TestCopyAddress(addr common.Address) ([]byte, error) {
	return _Wyvernexchangev1.Contract.TestCopyAddress(&_Wyvernexchangev1.CallOpts, addr)
}

// TestCopyAddress is a free data retrieval call binding the contract method 0x3464af6a.
//
// Solidity: function testCopyAddress(address addr) pure returns(bytes)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) TestCopyAddress(addr common.Address) ([]byte, error) {
	return _Wyvernexchangev1.Contract.TestCopyAddress(&_Wyvernexchangev1.CallOpts, addr)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) TokenTransferProxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "tokenTransferProxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) TokenTransferProxy() (common.Address, error) {
	return _Wyvernexchangev1.Contract.TokenTransferProxy(&_Wyvernexchangev1.CallOpts)
}

// TokenTransferProxy is a free data retrieval call binding the contract method 0x0eefdbad.
//
// Solidity: function tokenTransferProxy() view returns(address)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) TokenTransferProxy() (common.Address, error) {
	return _Wyvernexchangev1.Contract.TokenTransferProxy(&_Wyvernexchangev1.CallOpts)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) ValidateOrderParameters(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "validateOrderParameters_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ValidateOrderParameters(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ValidateOrderParameters(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// ValidateOrderParameters is a free data retrieval call binding the contract method 0xca595b9a.
//
// Solidity: function validateOrderParameters_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) ValidateOrderParameters(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ValidateOrderParameters(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) ValidateOrder(opts *bind.CallOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "validateOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ValidateOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ValidateOrder(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x60bef33a.
//
// Solidity: function validateOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) ValidateOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Wyvernexchangev1.Contract.ValidateOrder(&_Wyvernexchangev1.CallOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wyvernexchangev1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1Session) Version() (string, error) {
	return _Wyvernexchangev1.Contract.Version(&_Wyvernexchangev1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Wyvernexchangev1 *Wyvernexchangev1CallerSession) Version() (string, error) {
	return _Wyvernexchangev1.Contract.Version(&_Wyvernexchangev1.CallOpts)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) ApproveOrder(opts *bind.TransactOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "approveOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ApproveOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ApproveOrder(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// ApproveOrder is a paid mutator transaction binding the contract method 0x79666868.
//
// Solidity: function approveOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, bool orderbookInclusionDesired) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) ApproveOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, orderbookInclusionDesired bool) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ApproveOrder(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, orderbookInclusionDesired)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) payable returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) AtomicMatch(opts *bind.TransactOpts, addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "atomicMatch_", addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) payable returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) AtomicMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.AtomicMatch(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// AtomicMatch is a paid mutator transaction binding the contract method 0xab834bab.
//
// Solidity: function atomicMatch_(address[14] addrs, uint256[18] uints, uint8[8] feeMethodsSidesKindsHowToCalls, bytes calldataBuy, bytes calldataSell, bytes replacementPatternBuy, bytes replacementPatternSell, bytes staticExtradataBuy, bytes staticExtradataSell, uint8[2] vs, bytes32[5] rssMetadata) payable returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) AtomicMatch(addrs [14]common.Address, uints [18]*big.Int, feeMethodsSidesKindsHowToCalls [8]uint8, calldataBuy []byte, calldataSell []byte, replacementPatternBuy []byte, replacementPatternSell []byte, staticExtradataBuy []byte, staticExtradataSell []byte, vs [2]uint8, rssMetadata [5][32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.AtomicMatch(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethodsSidesKindsHowToCalls, calldataBuy, calldataSell, replacementPatternBuy, replacementPatternSell, staticExtradataBuy, staticExtradataSell, vs, rssMetadata)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) CancelOrder(opts *bind.TransactOpts, addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "cancelOrder_", addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) CancelOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.CancelOrder(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa8a41c70.
//
// Solidity: function cancelOrder_(address[7] addrs, uint256[9] uints, uint8 feeMethod, uint8 side, uint8 saleKind, uint8 howToCall, bytes calldata, bytes replacementPattern, bytes staticExtradata, uint8 v, bytes32 r, bytes32 s) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) CancelOrder(addrs [7]common.Address, uints [9]*big.Int, feeMethod uint8, side uint8, saleKind uint8, howToCall uint8, calldata []byte, replacementPattern []byte, staticExtradata []byte, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.CancelOrder(&_Wyvernexchangev1.TransactOpts, addrs, uints, feeMethod, side, saleKind, howToCall, calldata, replacementPattern, staticExtradata, v, r, s)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) ChangeMinimumMakerProtocolFee(opts *bind.TransactOpts, newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "changeMinimumMakerProtocolFee", newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeMinimumMakerProtocolFee(&_Wyvernexchangev1.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumMakerProtocolFee is a paid mutator transaction binding the contract method 0x14350c24.
//
// Solidity: function changeMinimumMakerProtocolFee(uint256 newMinimumMakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) ChangeMinimumMakerProtocolFee(newMinimumMakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeMinimumMakerProtocolFee(&_Wyvernexchangev1.TransactOpts, newMinimumMakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) ChangeMinimumTakerProtocolFee(opts *bind.TransactOpts, newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "changeMinimumTakerProtocolFee", newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeMinimumTakerProtocolFee(&_Wyvernexchangev1.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeMinimumTakerProtocolFee is a paid mutator transaction binding the contract method 0x1a6b13e2.
//
// Solidity: function changeMinimumTakerProtocolFee(uint256 newMinimumTakerProtocolFee) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) ChangeMinimumTakerProtocolFee(newMinimumTakerProtocolFee *big.Int) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeMinimumTakerProtocolFee(&_Wyvernexchangev1.TransactOpts, newMinimumTakerProtocolFee)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) ChangeProtocolFeeRecipient(opts *bind.TransactOpts, newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "changeProtocolFeeRecipient", newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeProtocolFeeRecipient(&_Wyvernexchangev1.TransactOpts, newProtocolFeeRecipient)
}

// ChangeProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x514f0330.
//
// Solidity: function changeProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) ChangeProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.ChangeProtocolFeeRecipient(&_Wyvernexchangev1.TransactOpts, newProtocolFeeRecipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) RenounceOwnership() (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.RenounceOwnership(&_Wyvernexchangev1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.RenounceOwnership(&_Wyvernexchangev1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.TransferOwnership(&_Wyvernexchangev1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wyvernexchangev1 *Wyvernexchangev1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wyvernexchangev1.Contract.TransferOwnership(&_Wyvernexchangev1.TransactOpts, newOwner)
}

// Wyvernexchangev1OrderApprovedPartOneIterator is returned from FilterOrderApprovedPartOne and is used to iterate over the raw logs and unpacked data for OrderApprovedPartOne events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderApprovedPartOneIterator struct {
	Event *Wyvernexchangev1OrderApprovedPartOne // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OrderApprovedPartOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OrderApprovedPartOne)
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
		it.Event = new(Wyvernexchangev1OrderApprovedPartOne)
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
func (it *Wyvernexchangev1OrderApprovedPartOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OrderApprovedPartOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OrderApprovedPartOne represents a OrderApprovedPartOne event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderApprovedPartOne struct {
	Hash             [32]byte
	Exchange         common.Address
	Maker            common.Address
	Taker            common.Address
	MakerRelayerFee  *big.Int
	TakerRelayerFee  *big.Int
	MakerProtocolFee *big.Int
	TakerProtocolFee *big.Int
	FeeRecipient     common.Address
	FeeMethod        uint8
	Side             uint8
	SaleKind         uint8
	Target           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartOne is a free log retrieval operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOrderApprovedPartOne(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (*Wyvernexchangev1OrderApprovedPartOneIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OrderApprovedPartOneIterator{contract: _Wyvernexchangev1.contract, event: "OrderApprovedPartOne", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartOne is a free log subscription operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOrderApprovedPartOne(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OrderApprovedPartOne, hash [][32]byte, maker []common.Address, feeRecipient []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OrderApprovedPartOne", hashRule, makerRule, feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OrderApprovedPartOne)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
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

// ParseOrderApprovedPartOne is a log parse operation binding the contract event 0x90c7f9f5b58c15f0f635bfb99f55d3d78fdbef3559e7d8abf5c81052a5276622.
//
// Solidity: event OrderApprovedPartOne(bytes32 indexed hash, address exchange, address indexed maker, address taker, uint256 makerRelayerFee, uint256 takerRelayerFee, uint256 makerProtocolFee, uint256 takerProtocolFee, address indexed feeRecipient, uint8 feeMethod, uint8 side, uint8 saleKind, address target)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOrderApprovedPartOne(log types.Log) (*Wyvernexchangev1OrderApprovedPartOne, error) {
	event := new(Wyvernexchangev1OrderApprovedPartOne)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderApprovedPartOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Wyvernexchangev1OrderApprovedPartTwoIterator is returned from FilterOrderApprovedPartTwo and is used to iterate over the raw logs and unpacked data for OrderApprovedPartTwo events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderApprovedPartTwoIterator struct {
	Event *Wyvernexchangev1OrderApprovedPartTwo // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OrderApprovedPartTwoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OrderApprovedPartTwo)
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
		it.Event = new(Wyvernexchangev1OrderApprovedPartTwo)
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
func (it *Wyvernexchangev1OrderApprovedPartTwoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OrderApprovedPartTwoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OrderApprovedPartTwo represents a OrderApprovedPartTwo event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderApprovedPartTwo struct {
	Hash                      [32]byte
	HowToCall                 uint8
	Calldata                  []byte
	ReplacementPattern        []byte
	StaticTarget              common.Address
	StaticExtradata           []byte
	PaymentToken              common.Address
	BasePrice                 *big.Int
	Extra                     *big.Int
	ListingTime               *big.Int
	ExpirationTime            *big.Int
	Salt                      *big.Int
	OrderbookInclusionDesired bool
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterOrderApprovedPartTwo is a free log retrieval operation binding the contract event 0xe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOrderApprovedPartTwo(opts *bind.FilterOpts, hash [][32]byte) (*Wyvernexchangev1OrderApprovedPartTwoIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OrderApprovedPartTwoIterator{contract: _Wyvernexchangev1.contract, event: "OrderApprovedPartTwo", logs: logs, sub: sub}, nil
}

// WatchOrderApprovedPartTwo is a free log subscription operation binding the contract event 0xe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOrderApprovedPartTwo(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OrderApprovedPartTwo, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OrderApprovedPartTwo", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OrderApprovedPartTwo)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
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

// ParseOrderApprovedPartTwo is a log parse operation binding the contract event 0xe55393c778364e440d958b39ac1debd99dcfae3775a8a04d1e79124adf6a2d08.
//
// Solidity: event OrderApprovedPartTwo(bytes32 indexed hash, uint8 howToCall, bytes calldata, bytes replacementPattern, address staticTarget, bytes staticExtradata, address paymentToken, uint256 basePrice, uint256 extra, uint256 listingTime, uint256 expirationTime, uint256 salt, bool orderbookInclusionDesired)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOrderApprovedPartTwo(log types.Log) (*Wyvernexchangev1OrderApprovedPartTwo, error) {
	event := new(Wyvernexchangev1OrderApprovedPartTwo)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderApprovedPartTwo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Wyvernexchangev1OrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderCancelledIterator struct {
	Event *Wyvernexchangev1OrderCancelled // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OrderCancelled)
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
		it.Event = new(Wyvernexchangev1OrderCancelled)
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
func (it *Wyvernexchangev1OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OrderCancelled represents a OrderCancelled event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOrderCancelled(opts *bind.FilterOpts, hash [][32]byte) (*Wyvernexchangev1OrderCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OrderCancelledIterator{contract: _Wyvernexchangev1.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OrderCancelled, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OrderCancelled", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OrderCancelled)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed hash)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOrderCancelled(log types.Log) (*Wyvernexchangev1OrderCancelled, error) {
	event := new(Wyvernexchangev1OrderCancelled)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Wyvernexchangev1OrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrdersMatchedIterator struct {
	Event *Wyvernexchangev1OrdersMatched // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OrdersMatched)
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
		it.Event = new(Wyvernexchangev1OrdersMatched)
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
func (it *Wyvernexchangev1OrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OrdersMatched represents a OrdersMatched event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OrdersMatched struct {
	BuyHash  [32]byte
	SellHash [32]byte
	Maker    common.Address
	Taker    common.Address
	Price    *big.Int
	Metadata [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price, bytes32 indexed metadata)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address, metadata [][32]byte) (*Wyvernexchangev1OrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OrdersMatchedIterator{contract: _Wyvernexchangev1.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price, bytes32 indexed metadata)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OrdersMatched, maker []common.Address, taker []common.Address, metadata [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	var metadataRule []interface{}
	for _, metadataItem := range metadata {
		metadataRule = append(metadataRule, metadataItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule, metadataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OrdersMatched)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0xc4109843e0b7d514e4c093114b863f8e7d8d9a458c372cd51bfe526b588006c9.
//
// Solidity: event OrdersMatched(bytes32 buyHash, bytes32 sellHash, address indexed maker, address indexed taker, uint256 price, bytes32 indexed metadata)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOrdersMatched(log types.Log) (*Wyvernexchangev1OrdersMatched, error) {
	event := new(Wyvernexchangev1OrdersMatched)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Wyvernexchangev1OwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OwnershipRenouncedIterator struct {
	Event *Wyvernexchangev1OwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OwnershipRenounced)
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
		it.Event = new(Wyvernexchangev1OwnershipRenounced)
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
func (it *Wyvernexchangev1OwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OwnershipRenounced represents a OwnershipRenounced event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*Wyvernexchangev1OwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OwnershipRenouncedIterator{contract: _Wyvernexchangev1.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OwnershipRenounced)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOwnershipRenounced(log types.Log) (*Wyvernexchangev1OwnershipRenounced, error) {
	event := new(Wyvernexchangev1OwnershipRenounced)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Wyvernexchangev1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OwnershipTransferredIterator struct {
	Event *Wyvernexchangev1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Wyvernexchangev1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Wyvernexchangev1OwnershipTransferred)
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
		it.Event = new(Wyvernexchangev1OwnershipTransferred)
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
func (it *Wyvernexchangev1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Wyvernexchangev1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Wyvernexchangev1OwnershipTransferred represents a OwnershipTransferred event raised by the Wyvernexchangev1 contract.
type Wyvernexchangev1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Wyvernexchangev1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Wyvernexchangev1OwnershipTransferredIterator{contract: _Wyvernexchangev1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Wyvernexchangev1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wyvernexchangev1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Wyvernexchangev1OwnershipTransferred)
				if err := _Wyvernexchangev1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Wyvernexchangev1 *Wyvernexchangev1Filterer) ParseOwnershipTransferred(log types.Log) (*Wyvernexchangev1OwnershipTransferred, error) {
	event := new(Wyvernexchangev1OwnershipTransferred)
	if err := _Wyvernexchangev1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
