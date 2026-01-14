// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// VolcanoERC20FactoryMetaData contains all meta data concerning the VolcanoERC20Factory contract.
var VolcanoERC20FactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_erc20MintFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_erc20MintTokenFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_routerAddressV3Fee\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_stakingAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocksFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingAmount\",\"type\":\"uint256\"}],\"name\":\"createTokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20MintFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20MintTokenFeePerc\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"}],\"name\":\"mintTokenBlocks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformContractFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddressV3Fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAddress\",\"outputs\":[{\"internalType\":\"contractVolcanoERC20StakingInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_platformContractFee\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"_erc20MintFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_erc20MintTokenFeePerc\",\"type\":\"uint96\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_fee\",\"type\":\"uint24\"}],\"name\":\"updateRouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingAddress\",\"type\":\"address\"}],\"name\":\"updateStakingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC20FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC20FactoryMetaData.ABI instead.
var VolcanoERC20FactoryABI = VolcanoERC20FactoryMetaData.ABI

// VolcanoERC20Factory is an auto generated Go binding around an Ethereum contract.
type VolcanoERC20Factory struct {
	VolcanoERC20FactoryCaller     // Read-only binding to the contract
	VolcanoERC20FactoryTransactor // Write-only binding to the contract
	VolcanoERC20FactoryFilterer   // Log filterer for contract events
}

// VolcanoERC20FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC20FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC20FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC20FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC20FactorySession struct {
	Contract     *VolcanoERC20Factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VolcanoERC20FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC20FactoryCallerSession struct {
	Contract *VolcanoERC20FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VolcanoERC20FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC20FactoryTransactorSession struct {
	Contract     *VolcanoERC20FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VolcanoERC20FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC20FactoryRaw struct {
	Contract *VolcanoERC20Factory // Generic contract binding to access the raw methods on
}

// VolcanoERC20FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC20FactoryCallerRaw struct {
	Contract *VolcanoERC20FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC20FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC20FactoryTransactorRaw struct {
	Contract *VolcanoERC20FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC20Factory creates a new instance of VolcanoERC20Factory, bound to a specific deployed contract.
func NewVolcanoERC20Factory(address common.Address, backend bind.ContractBackend) (*VolcanoERC20Factory, error) {
	contract, err := bindVolcanoERC20Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20Factory{VolcanoERC20FactoryCaller: VolcanoERC20FactoryCaller{contract: contract}, VolcanoERC20FactoryTransactor: VolcanoERC20FactoryTransactor{contract: contract}, VolcanoERC20FactoryFilterer: VolcanoERC20FactoryFilterer{contract: contract}}, nil
}

// NewVolcanoERC20FactoryCaller creates a new read-only instance of VolcanoERC20Factory, bound to a specific deployed contract.
func NewVolcanoERC20FactoryCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC20FactoryCaller, error) {
	contract, err := bindVolcanoERC20Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20FactoryCaller{contract: contract}, nil
}

// NewVolcanoERC20FactoryTransactor creates a new write-only instance of VolcanoERC20Factory, bound to a specific deployed contract.
func NewVolcanoERC20FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC20FactoryTransactor, error) {
	contract, err := bindVolcanoERC20Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20FactoryTransactor{contract: contract}, nil
}

// NewVolcanoERC20FactoryFilterer creates a new log filterer instance of VolcanoERC20Factory, bound to a specific deployed contract.
func NewVolcanoERC20FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC20FactoryFilterer, error) {
	contract, err := bindVolcanoERC20Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20FactoryFilterer{contract: contract}, nil
}

// bindVolcanoERC20Factory binds a generic wrapper to an already deployed contract.
func bindVolcanoERC20Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC20FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Factory *VolcanoERC20FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Factory.Contract.VolcanoERC20FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Factory *VolcanoERC20FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.VolcanoERC20FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Factory *VolcanoERC20FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.VolcanoERC20FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.contract.Transact(opts, method, params...)
}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) Erc20MintFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "erc20MintFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) Erc20MintFeePerc() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.Erc20MintFeePerc(&_VolcanoERC20Factory.CallOpts)
}

// Erc20MintFeePerc is a free data retrieval call binding the contract method 0xf4143c6d.
//
// Solidity: function erc20MintFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) Erc20MintFeePerc() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.Erc20MintFeePerc(&_VolcanoERC20Factory.CallOpts)
}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) Erc20MintTokenFeePerc(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "erc20MintTokenFeePerc")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) Erc20MintTokenFeePerc() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.Erc20MintTokenFeePerc(&_VolcanoERC20Factory.CallOpts)
}

// Erc20MintTokenFeePerc is a free data retrieval call binding the contract method 0x02b333f2.
//
// Solidity: function erc20MintTokenFeePerc() view returns(uint96)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) Erc20MintTokenFeePerc() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.Erc20MintTokenFeePerc(&_VolcanoERC20Factory.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC20Factory.Contract.Exists(&_VolcanoERC20Factory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _VolcanoERC20Factory.Contract.Exists(&_VolcanoERC20Factory.CallOpts, arg0)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.FeeRecipient(&_VolcanoERC20Factory.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) FeeRecipient() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.FeeRecipient(&_VolcanoERC20Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) Owner() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.Owner(&_VolcanoERC20Factory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.Owner(&_VolcanoERC20Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) PlatformContractFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "platformContractFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.PlatformContractFee(&_VolcanoERC20Factory.CallOpts)
}

// PlatformContractFee is a free data retrieval call binding the contract method 0xb69d4b98.
//
// Solidity: function platformContractFee() view returns(uint256)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) PlatformContractFee() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.PlatformContractFee(&_VolcanoERC20Factory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) RouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "routerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) RouterAddress() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.RouterAddress(&_VolcanoERC20Factory.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.RouterAddress(&_VolcanoERC20Factory.CallOpts)
}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) RouterAddressV3Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "routerAddressV3Fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) RouterAddressV3Fee() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.RouterAddressV3Fee(&_VolcanoERC20Factory.CallOpts)
}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) RouterAddressV3Fee() (*big.Int, error) {
	return _VolcanoERC20Factory.Contract.RouterAddressV3Fee(&_VolcanoERC20Factory.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCaller) StakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Factory.contract.Call(opts, &out, "stakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) StakingAddress() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.StakingAddress(&_VolcanoERC20Factory.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryCallerSession) StakingAddress() (common.Address, error) {
	return _VolcanoERC20Factory.Contract.StakingAddress(&_VolcanoERC20Factory.CallOpts)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xe67ee88f.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee, uint256 _stakingAmount) payable returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) CreateTokenContract(opts *bind.TransactOpts, _name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int, _stakingAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "createTokenContract", _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee, _stakingAmount)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xe67ee88f.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee, uint256 _stakingAmount) payable returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) CreateTokenContract(_name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int, _stakingAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.CreateTokenContract(&_VolcanoERC20Factory.TransactOpts, _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee, _stakingAmount)
}

// CreateTokenContract is a paid mutator transaction binding the contract method 0xe67ee88f.
//
// Solidity: function createTokenContract(string _name, string _symbol, string _uri, address _initialReceiver, uint256 _initialAmount, uint256 _capAmount, uint256 _mintBlocks, uint256 _mintBlocksFee, uint256 _stakingAmount) payable returns(address)
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) CreateTokenContract(_name string, _symbol string, _uri string, _initialReceiver common.Address, _initialAmount *big.Int, _capAmount *big.Int, _mintBlocks *big.Int, _mintBlocksFee *big.Int, _stakingAmount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.CreateTokenContract(&_VolcanoERC20Factory.TransactOpts, _name, _symbol, _uri, _initialReceiver, _initialAmount, _capAmount, _mintBlocks, _mintBlocksFee, _stakingAmount)
}

// MintTokenBlocks is a paid mutator transaction binding the contract method 0xad97a8fc.
//
// Solidity: function mintTokenBlocks(address tokenContractAddress, address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) MintTokenBlocks(opts *bind.TransactOpts, tokenContractAddress common.Address, to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "mintTokenBlocks", tokenContractAddress, to, count, refund)
}

// MintTokenBlocks is a paid mutator transaction binding the contract method 0xad97a8fc.
//
// Solidity: function mintTokenBlocks(address tokenContractAddress, address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) MintTokenBlocks(tokenContractAddress common.Address, to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.MintTokenBlocks(&_VolcanoERC20Factory.TransactOpts, tokenContractAddress, to, count, refund)
}

// MintTokenBlocks is a paid mutator transaction binding the contract method 0xad97a8fc.
//
// Solidity: function mintTokenBlocks(address tokenContractAddress, address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) MintTokenBlocks(tokenContractAddress common.Address, to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.MintTokenBlocks(&_VolcanoERC20Factory.TransactOpts, tokenContractAddress, to, count, refund)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.RenounceOwnership(&_VolcanoERC20Factory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.RenounceOwnership(&_VolcanoERC20Factory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.TransferOwnership(&_VolcanoERC20Factory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.TransferOwnership(&_VolcanoERC20Factory.TransactOpts, newOwner)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x712615c9.
//
// Solidity: function updatePlatformFees(uint256 _platformContractFee, uint96 _erc20MintFeePerc, uint96 _erc20MintTokenFeePerc, address _feeRecipient) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) UpdatePlatformFees(opts *bind.TransactOpts, _platformContractFee *big.Int, _erc20MintFeePerc *big.Int, _erc20MintTokenFeePerc *big.Int, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "updatePlatformFees", _platformContractFee, _erc20MintFeePerc, _erc20MintTokenFeePerc, _feeRecipient)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x712615c9.
//
// Solidity: function updatePlatformFees(uint256 _platformContractFee, uint96 _erc20MintFeePerc, uint96 _erc20MintTokenFeePerc, address _feeRecipient) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) UpdatePlatformFees(_platformContractFee *big.Int, _erc20MintFeePerc *big.Int, _erc20MintTokenFeePerc *big.Int, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdatePlatformFees(&_VolcanoERC20Factory.TransactOpts, _platformContractFee, _erc20MintFeePerc, _erc20MintTokenFeePerc, _feeRecipient)
}

// UpdatePlatformFees is a paid mutator transaction binding the contract method 0x712615c9.
//
// Solidity: function updatePlatformFees(uint256 _platformContractFee, uint96 _erc20MintFeePerc, uint96 _erc20MintTokenFeePerc, address _feeRecipient) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) UpdatePlatformFees(_platformContractFee *big.Int, _erc20MintFeePerc *big.Int, _erc20MintTokenFeePerc *big.Int, _feeRecipient common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdatePlatformFees(&_VolcanoERC20Factory.TransactOpts, _platformContractFee, _erc20MintFeePerc, _erc20MintTokenFeePerc, _feeRecipient)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0xc3c44417.
//
// Solidity: function updateRouterAddress(address _routerAddress, uint24 _fee) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) UpdateRouterAddress(opts *bind.TransactOpts, _routerAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "updateRouterAddress", _routerAddress, _fee)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0xc3c44417.
//
// Solidity: function updateRouterAddress(address _routerAddress, uint24 _fee) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) UpdateRouterAddress(_routerAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdateRouterAddress(&_VolcanoERC20Factory.TransactOpts, _routerAddress, _fee)
}

// UpdateRouterAddress is a paid mutator transaction binding the contract method 0xc3c44417.
//
// Solidity: function updateRouterAddress(address _routerAddress, uint24 _fee) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) UpdateRouterAddress(_routerAddress common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdateRouterAddress(&_VolcanoERC20Factory.TransactOpts, _routerAddress, _fee)
}

// UpdateStakingAddress is a paid mutator transaction binding the contract method 0x5b06a1d1.
//
// Solidity: function updateStakingAddress(address _stakingAddress) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactor) UpdateStakingAddress(opts *bind.TransactOpts, _stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.contract.Transact(opts, "updateStakingAddress", _stakingAddress)
}

// UpdateStakingAddress is a paid mutator transaction binding the contract method 0x5b06a1d1.
//
// Solidity: function updateStakingAddress(address _stakingAddress) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactorySession) UpdateStakingAddress(_stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdateStakingAddress(&_VolcanoERC20Factory.TransactOpts, _stakingAddress)
}

// UpdateStakingAddress is a paid mutator transaction binding the contract method 0x5b06a1d1.
//
// Solidity: function updateStakingAddress(address _stakingAddress) returns()
func (_VolcanoERC20Factory *VolcanoERC20FactoryTransactorSession) UpdateStakingAddress(_stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Factory.Contract.UpdateStakingAddress(&_VolcanoERC20Factory.TransactOpts, _stakingAddress)
}

// VolcanoERC20FactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC20Factory contract.
type VolcanoERC20FactoryOwnershipTransferredIterator struct {
	Event *VolcanoERC20FactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20FactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20FactoryOwnershipTransferred)
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
		it.Event = new(VolcanoERC20FactoryOwnershipTransferred)
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
func (it *VolcanoERC20FactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20FactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20FactoryOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC20Factory contract.
type VolcanoERC20FactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC20FactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC20Factory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20FactoryOwnershipTransferredIterator{contract: _VolcanoERC20Factory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC20FactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC20Factory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20FactoryOwnershipTransferred)
				if err := _VolcanoERC20Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC20FactoryOwnershipTransferred, error) {
	event := new(VolcanoERC20FactoryOwnershipTransferred)
	if err := _VolcanoERC20Factory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20FactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the VolcanoERC20Factory contract.
type VolcanoERC20FactoryTokenCreatedIterator struct {
	Event *VolcanoERC20FactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20FactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20FactoryTokenCreated)
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
		it.Event = new(VolcanoERC20FactoryTokenCreated)
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
func (it *VolcanoERC20FactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20FactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20FactoryTokenCreated represents a TokenCreated event raised by the VolcanoERC20Factory contract.
type VolcanoERC20FactoryTokenCreated struct {
	Caller common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts) (*VolcanoERC20FactoryTokenCreatedIterator, error) {

	logs, sub, err := _VolcanoERC20Factory.contract.FilterLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20FactoryTokenCreatedIterator{contract: _VolcanoERC20Factory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *VolcanoERC20FactoryTokenCreated) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC20Factory.contract.WatchLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20FactoryTokenCreated)
				if err := _VolcanoERC20Factory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// ParseTokenCreated is a log parse operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: event TokenCreated(address caller, address token)
func (_VolcanoERC20Factory *VolcanoERC20FactoryFilterer) ParseTokenCreated(log types.Log) (*VolcanoERC20FactoryTokenCreated, error) {
	event := new(VolcanoERC20FactoryTokenCreated)
	if err := _VolcanoERC20Factory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
