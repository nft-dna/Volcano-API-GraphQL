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

// VolcanoERC20StakingMetaData contains all meta data concerning the VolcanoERC20Staking contract.
var VolcanoERC20StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"penaltyRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BPS_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PenaltyRecipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RegisterTokenRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"RewardPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"earlyExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyExitPenaltyBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fundRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"stakeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VolcanoERC20StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use VolcanoERC20StakingMetaData.ABI instead.
var VolcanoERC20StakingABI = VolcanoERC20StakingMetaData.ABI

// VolcanoERC20Staking is an auto generated Go binding around an Ethereum contract.
type VolcanoERC20Staking struct {
	VolcanoERC20StakingCaller     // Read-only binding to the contract
	VolcanoERC20StakingTransactor // Write-only binding to the contract
	VolcanoERC20StakingFilterer   // Log filterer for contract events
}

// VolcanoERC20StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolcanoERC20StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolcanoERC20StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolcanoERC20StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolcanoERC20StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolcanoERC20StakingSession struct {
	Contract     *VolcanoERC20Staking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VolcanoERC20StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolcanoERC20StakingCallerSession struct {
	Contract *VolcanoERC20StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VolcanoERC20StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolcanoERC20StakingTransactorSession struct {
	Contract     *VolcanoERC20StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VolcanoERC20StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolcanoERC20StakingRaw struct {
	Contract *VolcanoERC20Staking // Generic contract binding to access the raw methods on
}

// VolcanoERC20StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolcanoERC20StakingCallerRaw struct {
	Contract *VolcanoERC20StakingCaller // Generic read-only contract binding to access the raw methods on
}

// VolcanoERC20StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolcanoERC20StakingTransactorRaw struct {
	Contract *VolcanoERC20StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolcanoERC20Staking creates a new instance of VolcanoERC20Staking, bound to a specific deployed contract.
func NewVolcanoERC20Staking(address common.Address, backend bind.ContractBackend) (*VolcanoERC20Staking, error) {
	contract, err := bindVolcanoERC20Staking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20Staking{VolcanoERC20StakingCaller: VolcanoERC20StakingCaller{contract: contract}, VolcanoERC20StakingTransactor: VolcanoERC20StakingTransactor{contract: contract}, VolcanoERC20StakingFilterer: VolcanoERC20StakingFilterer{contract: contract}}, nil
}

// NewVolcanoERC20StakingCaller creates a new read-only instance of VolcanoERC20Staking, bound to a specific deployed contract.
func NewVolcanoERC20StakingCaller(address common.Address, caller bind.ContractCaller) (*VolcanoERC20StakingCaller, error) {
	contract, err := bindVolcanoERC20Staking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20StakingCaller{contract: contract}, nil
}

// NewVolcanoERC20StakingTransactor creates a new write-only instance of VolcanoERC20Staking, bound to a specific deployed contract.
func NewVolcanoERC20StakingTransactor(address common.Address, transactor bind.ContractTransactor) (*VolcanoERC20StakingTransactor, error) {
	contract, err := bindVolcanoERC20Staking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20StakingTransactor{contract: contract}, nil
}

// NewVolcanoERC20StakingFilterer creates a new log filterer instance of VolcanoERC20Staking, bound to a specific deployed contract.
func NewVolcanoERC20StakingFilterer(address common.Address, filterer bind.ContractFilterer) (*VolcanoERC20StakingFilterer, error) {
	contract, err := bindVolcanoERC20Staking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20StakingFilterer{contract: contract}, nil
}

// bindVolcanoERC20Staking binds a generic wrapper to an already deployed contract.
func bindVolcanoERC20Staking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolcanoERC20StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Staking *VolcanoERC20StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Staking.Contract.VolcanoERC20StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Staking *VolcanoERC20StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.VolcanoERC20StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Staking *VolcanoERC20StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.VolcanoERC20StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VolcanoERC20Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.contract.Transact(opts, method, params...)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) BPSDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "BPS_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) BPSDENOMINATOR() (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.BPSDENOMINATOR(&_VolcanoERC20Staking.CallOpts)
}

// BPSDENOMINATOR is a free data retrieval call binding the contract method 0xe1a45218.
//
// Solidity: function BPS_DENOMINATOR() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) BPSDENOMINATOR() (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.BPSDENOMINATOR(&_VolcanoERC20Staking.CallOpts)
}

// PenaltyRecipient is a free data retrieval call binding the contract method 0xcfb9a726.
//
// Solidity: function PenaltyRecipient() view returns(address)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) PenaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "PenaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PenaltyRecipient is a free data retrieval call binding the contract method 0xcfb9a726.
//
// Solidity: function PenaltyRecipient() view returns(address)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) PenaltyRecipient() (common.Address, error) {
	return _VolcanoERC20Staking.Contract.PenaltyRecipient(&_VolcanoERC20Staking.CallOpts)
}

// PenaltyRecipient is a free data retrieval call binding the contract method 0xcfb9a726.
//
// Solidity: function PenaltyRecipient() view returns(address)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) PenaltyRecipient() (common.Address, error) {
	return _VolcanoERC20Staking.Contract.PenaltyRecipient(&_VolcanoERC20Staking.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x4560a052.
//
// Solidity: function RewardPool(address ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) RewardPool(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "RewardPool", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x4560a052.
//
// Solidity: function RewardPool(address ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) RewardPool(arg0 common.Address) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.RewardPool(&_VolcanoERC20Staking.CallOpts, arg0)
}

// RewardPool is a free data retrieval call binding the contract method 0x4560a052.
//
// Solidity: function RewardPool(address ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) RewardPool(arg0 common.Address) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.RewardPool(&_VolcanoERC20Staking.CallOpts, arg0)
}

// EarlyExitPenaltyBps is a free data retrieval call binding the contract method 0xd827b7e8.
//
// Solidity: function earlyExitPenaltyBps() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) EarlyExitPenaltyBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "earlyExitPenaltyBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EarlyExitPenaltyBps is a free data retrieval call binding the contract method 0xd827b7e8.
//
// Solidity: function earlyExitPenaltyBps() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) EarlyExitPenaltyBps() (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.EarlyExitPenaltyBps(&_VolcanoERC20Staking.CallOpts)
}

// EarlyExitPenaltyBps is a free data retrieval call binding the contract method 0xd827b7e8.
//
// Solidity: function earlyExitPenaltyBps() view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) EarlyExitPenaltyBps() (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.EarlyExitPenaltyBps(&_VolcanoERC20Staking.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) RewardRate(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "rewardRate", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) RewardRate(arg0 *big.Int) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.RewardRate(&_VolcanoERC20Staking.CallOpts, arg0)
}

// RewardRate is a free data retrieval call binding the contract method 0xcea01962.
//
// Solidity: function rewardRate(uint256 ) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) RewardRate(arg0 *big.Int) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.RewardRate(&_VolcanoERC20Staking.CallOpts, arg0)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) StakeCount(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "stakeCount", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) StakeCount(user common.Address) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.StakeCount(&_VolcanoERC20Staking.CallOpts, user)
}

// StakeCount is a free data retrieval call binding the contract method 0x33060d90.
//
// Solidity: function stakeCount(address user) view returns(uint256)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) StakeCount(user common.Address) (*big.Int, error) {
	return _VolcanoERC20Staking.Contract.StakeCount(&_VolcanoERC20Staking.CallOpts, user)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(address token, uint256 amount, uint256 reward, uint256 unlockTime, bool claimed)
func (_VolcanoERC20Staking *VolcanoERC20StakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Token      common.Address
	Amount     *big.Int
	Reward     *big.Int
	UnlockTime *big.Int
	Claimed    bool
}, error) {
	var out []interface{}
	err := _VolcanoERC20Staking.contract.Call(opts, &out, "stakes", arg0, arg1)

	outstruct := new(struct {
		Token      common.Address
		Amount     *big.Int
		Reward     *big.Int
		UnlockTime *big.Int
		Claimed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(address token, uint256 amount, uint256 reward, uint256 unlockTime, bool claimed)
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Token      common.Address
	Amount     *big.Int
	Reward     *big.Int
	UnlockTime *big.Int
	Claimed    bool
}, error) {
	return _VolcanoERC20Staking.Contract.Stakes(&_VolcanoERC20Staking.CallOpts, arg0, arg1)
}

// Stakes is a free data retrieval call binding the contract method 0x584b62a1.
//
// Solidity: function stakes(address , uint256 ) view returns(address token, uint256 amount, uint256 reward, uint256 unlockTime, bool claimed)
func (_VolcanoERC20Staking *VolcanoERC20StakingCallerSession) Stakes(arg0 common.Address, arg1 *big.Int) (struct {
	Token      common.Address
	Amount     *big.Int
	Reward     *big.Int
	UnlockTime *big.Int
	Claimed    bool
}, error) {
	return _VolcanoERC20Staking.Contract.Stakes(&_VolcanoERC20Staking.CallOpts, arg0, arg1)
}

// RegisterTokenRewards is a paid mutator transaction binding the contract method 0x0ba1e10e.
//
// Solidity: function RegisterTokenRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) RegisterTokenRewards(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "RegisterTokenRewards", token, amount)
}

// RegisterTokenRewards is a paid mutator transaction binding the contract method 0x0ba1e10e.
//
// Solidity: function RegisterTokenRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) RegisterTokenRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.RegisterTokenRewards(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// RegisterTokenRewards is a paid mutator transaction binding the contract method 0x0ba1e10e.
//
// Solidity: function RegisterTokenRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) RegisterTokenRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.RegisterTokenRewards(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// EarlyExit is a paid mutator transaction binding the contract method 0xb8af3d3e.
//
// Solidity: function earlyExit(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) EarlyExit(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "earlyExit", index)
}

// EarlyExit is a paid mutator transaction binding the contract method 0xb8af3d3e.
//
// Solidity: function earlyExit(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) EarlyExit(index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.EarlyExit(&_VolcanoERC20Staking.TransactOpts, index)
}

// EarlyExit is a paid mutator transaction binding the contract method 0xb8af3d3e.
//
// Solidity: function earlyExit(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) EarlyExit(index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.EarlyExit(&_VolcanoERC20Staking.TransactOpts, index)
}

// FundRewards is a paid mutator transaction binding the contract method 0x928082b2.
//
// Solidity: function fundRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) FundRewards(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "fundRewards", token, amount)
}

// FundRewards is a paid mutator transaction binding the contract method 0x928082b2.
//
// Solidity: function fundRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) FundRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.FundRewards(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// FundRewards is a paid mutator transaction binding the contract method 0x928082b2.
//
// Solidity: function fundRewards(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) FundRewards(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.FundRewards(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x8cd4426d.
//
// Solidity: function rescueERC20(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) RescueERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "rescueERC20", token, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x8cd4426d.
//
// Solidity: function rescueERC20(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) RescueERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.RescueERC20(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0x8cd4426d.
//
// Solidity: function rescueERC20(address token, uint256 amount) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) RescueERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.RescueERC20(&_VolcanoERC20Staking.TransactOpts, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address token, uint256 amount, uint256 duration) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) Stake(opts *bind.TransactOpts, token common.Address, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "stake", token, amount, duration)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address token, uint256 amount, uint256 duration) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) Stake(token common.Address, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.Stake(&_VolcanoERC20Staking.TransactOpts, token, amount, duration)
}

// Stake is a paid mutator transaction binding the contract method 0x0c51b88f.
//
// Solidity: function stake(address token, uint256 amount, uint256 duration) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) Stake(token common.Address, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.Stake(&_VolcanoERC20Staking.TransactOpts, token, amount, duration)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0x11cfd210.
//
// Solidity: function stakeWithPermit(address token, uint256 amount, uint256 duration, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) StakeWithPermit(opts *bind.TransactOpts, token common.Address, amount *big.Int, duration *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "stakeWithPermit", token, amount, duration, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0x11cfd210.
//
// Solidity: function stakeWithPermit(address token, uint256 amount, uint256 duration, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) StakeWithPermit(token common.Address, amount *big.Int, duration *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.StakeWithPermit(&_VolcanoERC20Staking.TransactOpts, token, amount, duration, deadline, v, r, s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0x11cfd210.
//
// Solidity: function stakeWithPermit(address token, uint256 amount, uint256 duration, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) StakeWithPermit(token common.Address, amount *big.Int, duration *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.StakeWithPermit(&_VolcanoERC20Staking.TransactOpts, token, amount, duration, deadline, v, r, s)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactor) Unstake(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.contract.Transact(opts, "unstake", index)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingSession) Unstake(index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.Unstake(&_VolcanoERC20Staking.TransactOpts, index)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 index) returns()
func (_VolcanoERC20Staking *VolcanoERC20StakingTransactorSession) Unstake(index *big.Int) (*types.Transaction, error) {
	return _VolcanoERC20Staking.Contract.Unstake(&_VolcanoERC20Staking.TransactOpts, index)
}

// VolcanoERC20StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the VolcanoERC20Staking contract.
type VolcanoERC20StakingStakedIterator struct {
	Event *VolcanoERC20StakingStaked // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20StakingStaked)
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
		it.Event = new(VolcanoERC20StakingStaked)
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
func (it *VolcanoERC20StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20StakingStaked represents a Staked event raised by the VolcanoERC20Staking contract.
type VolcanoERC20StakingStaked struct {
	User       common.Address
	Token      common.Address
	Amount     *big.Int
	Reward     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address token, uint256 amount, uint256 reward, uint256 unlockTime)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*VolcanoERC20StakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VolcanoERC20Staking.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20StakingStakedIterator{contract: _VolcanoERC20Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address token, uint256 amount, uint256 reward, uint256 unlockTime)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *VolcanoERC20StakingStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VolcanoERC20Staking.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20StakingStaked)
				if err := _VolcanoERC20Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address token, uint256 amount, uint256 reward, uint256 unlockTime)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) ParseStaked(log types.Log) (*VolcanoERC20StakingStaked, error) {
	event := new(VolcanoERC20StakingStaked)
	if err := _VolcanoERC20Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VolcanoERC20StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the VolcanoERC20Staking contract.
type VolcanoERC20StakingUnstakedIterator struct {
	Event *VolcanoERC20StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *VolcanoERC20StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VolcanoERC20StakingUnstaked)
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
		it.Event = new(VolcanoERC20StakingUnstaked)
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
func (it *VolcanoERC20StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VolcanoERC20StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VolcanoERC20StakingUnstaked represents a Unstaked event raised by the VolcanoERC20Staking contract.
type VolcanoERC20StakingUnstaked struct {
	User    common.Address
	Token   common.Address
	Amount  *big.Int
	Reward  *big.Int
	Penalty *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x9ad65b6c7ddab48c57c60229304ab9d8a7726f680d2319c0704c705eb3f8c635.
//
// Solidity: event Unstaked(address indexed user, address token, uint256 amount, uint256 reward, uint256 penalty)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address) (*VolcanoERC20StakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VolcanoERC20Staking.contract.FilterLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return &VolcanoERC20StakingUnstakedIterator{contract: _VolcanoERC20Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x9ad65b6c7ddab48c57c60229304ab9d8a7726f680d2319c0704c705eb3f8c635.
//
// Solidity: event Unstaked(address indexed user, address token, uint256 amount, uint256 reward, uint256 penalty)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *VolcanoERC20StakingUnstaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _VolcanoERC20Staking.contract.WatchLogs(opts, "Unstaked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VolcanoERC20StakingUnstaked)
				if err := _VolcanoERC20Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x9ad65b6c7ddab48c57c60229304ab9d8a7726f680d2319c0704c705eb3f8c635.
//
// Solidity: event Unstaked(address indexed user, address token, uint256 amount, uint256 reward, uint256 penalty)
func (_VolcanoERC20Staking *VolcanoERC20StakingFilterer) ParseUnstaked(log types.Log) (*VolcanoERC20StakingUnstaked, error) {
	event := new(VolcanoERC20StakingUnstaked)
	if err := _VolcanoERC20Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
