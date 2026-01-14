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

// VolcanoERC20TokenMetaData contains all meta data concerning the VolcanoERC20Token contract.
var VolcanoERC20TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mintBlocksFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_routerAddress\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_routerAddressV3Fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"_stakingAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BlocksMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"positionOwner\",\"type\":\"address\"}],\"name\":\"burnV3Position\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"name\":\"collectV3Position\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractVolcanoERC20FactoryInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractVolcanoERC20StakingInterface\",\"name\":\"_stakingAddress\",\"type\":\"address\"}],\"name\":\"initilizeStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"refund\",\"type\":\"bool\"}],\"name\":\"mintBlocks\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksMaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintBlocksSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddress\",\"outputs\":[{\"internalType\":\"contractUniswapRouterInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerAddressV3Fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAddress\",\"outputs\":[{\"internalType\":\"contractVolcanoERC20StakingInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"updateContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"v3positions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VolcanoERC20TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC20TokenMetaData.ABI instead.
var VolcanoERC20TokenABI = VolcanoERC20TokenMetaData.ABI

// VolcanoERC20Token is an auto generated Go binding around an Ethereum contract.
type VolcanoERC20Token struct {
	VolcanoERC20TokenCaller     // Read-only binding to the contract
	VolcanoERC20TokenTransactor // Write-only binding to the contract
	VolcanoERC20TokenFilterer   // Log filterer for contract events
}

// VolcanoERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC20TokenSession struct {
	Contract     *VolcanoERC20Token // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VolcanoERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC20TokenCallerSession struct {
	Contract *VolcanoERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// VolcanoERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC20TokenTransactorSession struct {
	Contract     *VolcanoERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// VolcanoERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC20TokenRaw struct {
	Contract *VolcanoERC20Token // Generic contract binding to access the raw methods on
}

// VolcanoERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC20TokenCallerRaw struct {
	Contract *VolcanoERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC20TokenTransactorRaw struct {
	Contract *VolcanoERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC20Token creates a new instance of VolcanoERC20Token, bound to a specific deployed contract.
func NewVolcanoERC20Token(address common.Address, backend bind.ContractBackend) (*VolcanoERC20Token, error) {
	contract, err := bindVolcanoERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20Token{VolcanoERC20TokenCaller: VolcanoERC20TokenCaller{contract: contract}, VolcanoERC20TokenTransactor: VolcanoERC20TokenTransactor{contract: contract}, VolcanoERC20TokenFilterer: VolcanoERC20TokenFilterer{contract: contract}}, nil
}

// NewVolcanoERC20TokenCaller creates a new read-only instance of VolcanoERC20Token, bound to a specific deployed contract.
func NewVolcanoERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC20TokenCaller, error) {
	contract, err := bindVolcanoERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenCaller{contract: contract}, nil
}

// NewVolcanoERC20TokenTransactor creates a new write-only instance of VolcanoERC20Token, bound to a specific deployed contract.
func NewVolcanoERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC20TokenTransactor, error) {
	contract, err := bindVolcanoERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenTransactor{contract: contract}, nil
}

// NewVolcanoERC20TokenFilterer creates a new log filterer instance of VolcanoERC20Token, bound to a specific deployed contract.
func NewVolcanoERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC20TokenFilterer, error) {
	contract, err := bindVolcanoERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenFilterer{contract: contract}, nil
}

// bindVolcanoERC20Token binds a generic wrapper to an already deployed contract.
func bindVolcanoERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC20TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Token *VolcanoERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Token.Contract.VolcanoERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Token *VolcanoERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.VolcanoERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Token *VolcanoERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.VolcanoERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Token *VolcanoERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VolcanoERC20Token.Contract.DOMAINSEPARATOR(&_VolcanoERC20Token.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _VolcanoERC20Token.Contract.DOMAINSEPARATOR(&_VolcanoERC20Token.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Allowance(&_VolcanoERC20Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Allowance(&_VolcanoERC20Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.BalanceOf(&_VolcanoERC20Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.BalanceOf(&_VolcanoERC20Token.CallOpts, account)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Cap() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Cap(&_VolcanoERC20Token.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Cap() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Cap(&_VolcanoERC20Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) ContractURI() (string, error) {
	return _VolcanoERC20Token.Contract.ContractURI(&_VolcanoERC20Token.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) ContractURI() (string, error) {
	return _VolcanoERC20Token.Contract.ContractURI(&_VolcanoERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Decimals() (uint8, error) {
	return _VolcanoERC20Token.Contract.Decimals(&_VolcanoERC20Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Decimals() (uint8, error) {
	return _VolcanoERC20Token.Contract.Decimals(&_VolcanoERC20Token.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VolcanoERC20Token.Contract.Eip712Domain(&_VolcanoERC20Token.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _VolcanoERC20Token.Contract.Eip712Domain(&_VolcanoERC20Token.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Factory() (common.Address, error) {
	return _VolcanoERC20Token.Contract.Factory(&_VolcanoERC20Token.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Factory() (common.Address, error) {
	return _VolcanoERC20Token.Contract.Factory(&_VolcanoERC20Token.CallOpts)
}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) InitialReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "initialReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) InitialReserves() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.InitialReserves(&_VolcanoERC20Token.CallOpts)
}

// InitialReserves is a free data retrieval call binding the contract method 0x87890455.
//
// Solidity: function initialReserves() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) InitialReserves() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.InitialReserves(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) MintBlocksAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "mintBlocksAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) MintBlocksAmount() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksAmount(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksAmount is a free data retrieval call binding the contract method 0x92b52b8e.
//
// Solidity: function mintBlocksAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) MintBlocksAmount() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksAmount(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) MintBlocksFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "mintBlocksFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) MintBlocksFee() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksFee(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksFee is a free data retrieval call binding the contract method 0x63441ee2.
//
// Solidity: function mintBlocksFee() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) MintBlocksFee() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksFee(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) MintBlocksMaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "mintBlocksMaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) MintBlocksMaxSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksMaxSupply(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksMaxSupply is a free data retrieval call binding the contract method 0xbbb1fd51.
//
// Solidity: function mintBlocksMaxSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) MintBlocksMaxSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksMaxSupply(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) MintBlocksSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "mintBlocksSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) MintBlocksSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksSupply(&_VolcanoERC20Token.CallOpts)
}

// MintBlocksSupply is a free data retrieval call binding the contract method 0xcf82f26d.
//
// Solidity: function mintBlocksSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) MintBlocksSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.MintBlocksSupply(&_VolcanoERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Name() (string, error) {
	return _VolcanoERC20Token.Contract.Name(&_VolcanoERC20Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Name() (string, error) {
	return _VolcanoERC20Token.Contract.Name(&_VolcanoERC20Token.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Nonces(&_VolcanoERC20Token.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.Nonces(&_VolcanoERC20Token.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Owner() (common.Address, error) {
	return _VolcanoERC20Token.Contract.Owner(&_VolcanoERC20Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Owner() (common.Address, error) {
	return _VolcanoERC20Token.Contract.Owner(&_VolcanoERC20Token.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) RouterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "routerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC20Token.Contract.RouterAddress(&_VolcanoERC20Token.CallOpts)
}

// RouterAddress is a free data retrieval call binding the contract method 0x3268cc56.
//
// Solidity: function routerAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) RouterAddress() (common.Address, error) {
	return _VolcanoERC20Token.Contract.RouterAddress(&_VolcanoERC20Token.CallOpts)
}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) RouterAddressV3Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "routerAddressV3Fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) RouterAddressV3Fee() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.RouterAddressV3Fee(&_VolcanoERC20Token.CallOpts)
}

// RouterAddressV3Fee is a free data retrieval call binding the contract method 0xa09d8698.
//
// Solidity: function routerAddressV3Fee() view returns(uint24)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) RouterAddressV3Fee() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.RouterAddressV3Fee(&_VolcanoERC20Token.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) StakingAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "stakingAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) StakingAddress() (common.Address, error) {
	return _VolcanoERC20Token.Contract.StakingAddress(&_VolcanoERC20Token.CallOpts)
}

// StakingAddress is a free data retrieval call binding the contract method 0xd7b4be24.
//
// Solidity: function stakingAddress() view returns(address)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) StakingAddress() (common.Address, error) {
	return _VolcanoERC20Token.Contract.StakingAddress(&_VolcanoERC20Token.CallOpts)
}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) StakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "stakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) StakingAmount() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.StakingAmount(&_VolcanoERC20Token.CallOpts)
}

// StakingAmount is a free data retrieval call binding the contract method 0x739a3e02.
//
// Solidity: function stakingAmount() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) StakingAmount() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.StakingAmount(&_VolcanoERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Symbol() (string, error) {
	return _VolcanoERC20Token.Contract.Symbol(&_VolcanoERC20Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) Symbol() (string, error) {
	return _VolcanoERC20Token.Contract.Symbol(&_VolcanoERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.TotalSupply(&_VolcanoERC20Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _VolcanoERC20Token.Contract.TotalSupply(&_VolcanoERC20Token.CallOpts)
}

// V3positions is a free data retrieval call binding the contract method 0x0798840c.
//
// Solidity: function v3positions(address ) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCaller) V3positions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Token.contract.Call(opts, &out, "v3positions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// V3positions is a free data retrieval call binding the contract method 0x0798840c.
//
// Solidity: function v3positions(address ) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) V3positions(arg0 common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.V3positions(&_VolcanoERC20Token.CallOpts, arg0)
}

// V3positions is a free data retrieval call binding the contract method 0x0798840c.
//
// Solidity: function v3positions(address ) view returns(uint256)
func (_VolcanoERC20Token *VolcanoERC20TokenCallerSession) V3positions(arg0 common.Address) (*big.Int, error) {
	return _VolcanoERC20Token.Contract.V3positions(&_VolcanoERC20Token.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Approve(&_VolcanoERC20Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Approve(&_VolcanoERC20Token.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Burn(&_VolcanoERC20Token.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Burn(&_VolcanoERC20Token.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.BurnFrom(&_VolcanoERC20Token.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.BurnFrom(&_VolcanoERC20Token.TransactOpts, account, amount)
}

// BurnV3Position is a paid mutator transaction binding the contract method 0x01499963.
//
// Solidity: function burnV3Position(address positionOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) BurnV3Position(opts *bind.TransactOpts, positionOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "burnV3Position", positionOwner)
}

// BurnV3Position is a paid mutator transaction binding the contract method 0x01499963.
//
// Solidity: function burnV3Position(address positionOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) BurnV3Position(positionOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.BurnV3Position(&_VolcanoERC20Token.TransactOpts, positionOwner)
}

// BurnV3Position is a paid mutator transaction binding the contract method 0x01499963.
//
// Solidity: function burnV3Position(address positionOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) BurnV3Position(positionOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.BurnV3Position(&_VolcanoERC20Token.TransactOpts, positionOwner)
}

// CollectV3Position is a paid mutator transaction binding the contract method 0xb7fe5eb6.
//
// Solidity: function collectV3Position(uint256 tokenid) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) CollectV3Position(opts *bind.TransactOpts, tokenid *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "collectV3Position", tokenid)
}

// CollectV3Position is a paid mutator transaction binding the contract method 0xb7fe5eb6.
//
// Solidity: function collectV3Position(uint256 tokenid) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) CollectV3Position(tokenid *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.CollectV3Position(&_VolcanoERC20Token.TransactOpts, tokenid)
}

// CollectV3Position is a paid mutator transaction binding the contract method 0xb7fe5eb6.
//
// Solidity: function collectV3Position(uint256 tokenid) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) CollectV3Position(tokenid *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.CollectV3Position(&_VolcanoERC20Token.TransactOpts, tokenid)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.DecreaseAllowance(&_VolcanoERC20Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.DecreaseAllowance(&_VolcanoERC20Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.IncreaseAllowance(&_VolcanoERC20Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.IncreaseAllowance(&_VolcanoERC20Token.TransactOpts, spender, addedValue)
}

// InitilizeStaking is a paid mutator transaction binding the contract method 0x86372418.
//
// Solidity: function initilizeStaking(address _stakingAddress) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) InitilizeStaking(opts *bind.TransactOpts, _stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "initilizeStaking", _stakingAddress)
}

// InitilizeStaking is a paid mutator transaction binding the contract method 0x86372418.
//
// Solidity: function initilizeStaking(address _stakingAddress) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) InitilizeStaking(_stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.InitilizeStaking(&_VolcanoERC20Token.TransactOpts, _stakingAddress)
}

// InitilizeStaking is a paid mutator transaction binding the contract method 0x86372418.
//
// Solidity: function initilizeStaking(address _stakingAddress) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) InitilizeStaking(_stakingAddress common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.InitilizeStaking(&_VolcanoERC20Token.TransactOpts, _stakingAddress)
}

// MintBlocks is a paid mutator transaction binding the contract method 0xd64d2031.
//
// Solidity: function mintBlocks(address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) MintBlocks(opts *bind.TransactOpts, to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "mintBlocks", to, count, refund)
}

// MintBlocks is a paid mutator transaction binding the contract method 0xd64d2031.
//
// Solidity: function mintBlocks(address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) MintBlocks(to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.MintBlocks(&_VolcanoERC20Token.TransactOpts, to, count, refund)
}

// MintBlocks is a paid mutator transaction binding the contract method 0xd64d2031.
//
// Solidity: function mintBlocks(address to, uint256 count, bool refund) payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) MintBlocks(to common.Address, count *big.Int, refund bool) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.MintBlocks(&_VolcanoERC20Token.TransactOpts, to, count, refund)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Permit(&_VolcanoERC20Token.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Permit(&_VolcanoERC20Token.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.RenounceOwnership(&_VolcanoERC20Token.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.RenounceOwnership(&_VolcanoERC20Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Transfer(&_VolcanoERC20Token.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Transfer(&_VolcanoERC20Token.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.TransferFrom(&_VolcanoERC20Token.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.TransferFrom(&_VolcanoERC20Token.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.TransferOwnership(&_VolcanoERC20Token.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.TransferOwnership(&_VolcanoERC20Token.TransactOpts, newOwner)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) UpdateContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "updateContractURI", _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.UpdateContractURI(&_VolcanoERC20Token.TransactOpts, _uri)
}

// UpdateContractURI is a paid mutator transaction binding the contract method 0x7e5b1e24.
//
// Solidity: function updateContractURI(string _uri) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) UpdateContractURI(_uri string) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.UpdateContractURI(&_VolcanoERC20Token.TransactOpts, _uri)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.WithdrawETH(&_VolcanoERC20Token.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.WithdrawETH(&_VolcanoERC20Token.TransactOpts, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.Transact(opts, "withdrawTokens", token, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.WithdrawTokens(&_VolcanoERC20Token.TransactOpts, token, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.WithdrawTokens(&_VolcanoERC20Token.TransactOpts, token, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Token.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenSession) Receive() (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Receive(&_VolcanoERC20Token.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_VolcanoERC20Token *VolcanoERC20TokenTransactorSession) Receive() (*types.Transaction, error) {
	return _VolcanoERC20Token.Contract.Receive(&_VolcanoERC20Token.TransactOpts)
}

// VolcanoERC20TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenApprovalIterator struct {
	Event *VolcanoERC20TokenApproval // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20TokenApproval)
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
		it.Event = new(VolcanoERC20TokenApproval)
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
func (it *VolcanoERC20TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20TokenApproval represents a Approval event raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VolcanoERC20TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenApprovalIterator{contract: _VolcanoERC20Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VolcanoERC20TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20TokenApproval)
				if err := _VolcanoERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) ParseApproval(log types.Log) (*VolcanoERC20TokenApproval, error) {
	event := new(VolcanoERC20TokenApproval)
	if err := _VolcanoERC20Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20TokenBlocksMintedIterator is returned from FilterBlocksMinted and is used to iterate over the raw logs and unpacked data for BlocksMinted events raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenBlocksMintedIterator struct {
	Event *VolcanoERC20TokenBlocksMinted // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20TokenBlocksMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20TokenBlocksMinted)
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
		it.Event = new(VolcanoERC20TokenBlocksMinted)
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
func (it *VolcanoERC20TokenBlocksMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20TokenBlocksMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20TokenBlocksMinted represents a BlocksMinted event raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenBlocksMinted struct {
	Receiver common.Address
	Count    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlocksMinted is a free log retrieval operation binding the contract event 0x50bb62406ff32b2601952fa6d2b8bf6f088dfb08752b3657f13f2f9a9acb27ca.
//
// Solidity: event BlocksMinted(address receiver, uint256 count)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) FilterBlocksMinted(opts *bind.FilterOpts) (*VolcanoERC20TokenBlocksMintedIterator, error) {

	logs, sub, err := _VolcanoERC20Token.contract.FilterLogs(opts, "BlocksMinted")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenBlocksMintedIterator{contract: _VolcanoERC20Token.contract, event: "BlocksMinted", logs: logs, sub: sub}, nil
}

// WatchBlocksMinted is a free log subscription operation binding the contract event 0x50bb62406ff32b2601952fa6d2b8bf6f088dfb08752b3657f13f2f9a9acb27ca.
//
// Solidity: event BlocksMinted(address receiver, uint256 count)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) WatchBlocksMinted(opts *bind.WatchOpts, sink chan<- *VolcanoERC20TokenBlocksMinted) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC20Token.contract.WatchLogs(opts, "BlocksMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20TokenBlocksMinted)
				if err := _VolcanoERC20Token.contract.UnpackLog(event, "BlocksMinted", log); err != nil {
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

// ParseBlocksMinted is a log parse operation binding the contract event 0x50bb62406ff32b2601952fa6d2b8bf6f088dfb08752b3657f13f2f9a9acb27ca.
//
// Solidity: event BlocksMinted(address receiver, uint256 count)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) ParseBlocksMinted(log types.Log) (*VolcanoERC20TokenBlocksMinted, error) {
	event := new(VolcanoERC20TokenBlocksMinted)
	if err := _VolcanoERC20Token.contract.UnpackLog(event, "BlocksMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20TokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenEIP712DomainChangedIterator struct {
	Event *VolcanoERC20TokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20TokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20TokenEIP712DomainChanged)
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
		it.Event = new(VolcanoERC20TokenEIP712DomainChanged)
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
func (it *VolcanoERC20TokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20TokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20TokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*VolcanoERC20TokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _VolcanoERC20Token.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenEIP712DomainChangedIterator{contract: _VolcanoERC20Token.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *VolcanoERC20TokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _VolcanoERC20Token.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20TokenEIP712DomainChanged)
				if err := _VolcanoERC20Token.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) ParseEIP712DomainChanged(log types.Log) (*VolcanoERC20TokenEIP712DomainChanged, error) {
	event := new(VolcanoERC20TokenEIP712DomainChanged)
	if err := _VolcanoERC20Token.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20TokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenOwnershipTransferredIterator struct {
	Event *VolcanoERC20TokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20TokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20TokenOwnershipTransferred)
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
		it.Event = new(VolcanoERC20TokenOwnershipTransferred)
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
func (it *VolcanoERC20TokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20TokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20TokenOwnershipTransferred represents a OwnershipTransferred event raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VolcanoERC20TokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenOwnershipTransferredIterator{contract: _VolcanoERC20Token.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VolcanoERC20TokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20TokenOwnershipTransferred)
				if err := _VolcanoERC20Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) ParseOwnershipTransferred(log types.Log) (*VolcanoERC20TokenOwnershipTransferred, error) {
	event := new(VolcanoERC20TokenOwnershipTransferred)
	if err := _VolcanoERC20Token.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenTransferIterator struct {
	Event *VolcanoERC20TokenTransfer // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20TokenTransfer)
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
		it.Event = new(VolcanoERC20TokenTransfer)
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
func (it *VolcanoERC20TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20TokenTransfer represents a Transfer event raised by the VolcanoERC20Token contract.
type VolcanoERC20TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VolcanoERC20TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20TokenTransferIterator{contract: _VolcanoERC20Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VolcanoERC20TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VolcanoERC20Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20TokenTransfer)
				if err := _VolcanoERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VolcanoERC20Token *VolcanoERC20TokenFilterer) ParseTransfer(log types.Log) (*VolcanoERC20TokenTransfer, error) {
	event := new(VolcanoERC20TokenTransfer)
	if err := _VolcanoERC20Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
