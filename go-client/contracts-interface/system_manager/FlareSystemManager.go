// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system_manager

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

// IFlareSystemsManagerNumberOfWeightBasedClaims is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemsManagerNumberOfWeightBasedClaims struct {
	RewardManagerId       *big.Int
	NoOfWeightBasedClaims *big.Int
}

// IFlareSystemsManagerSignature is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemsManagerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SystemManagerMetaData contains all meta data concerning the SystemManager contract.
var SystemManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomAcquisitionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RewardEpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIFlareSystemsManager.NumberOfWeightBasedClaims[]\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"RewardsSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SignUptimeVoteEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"SigningPolicySigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uptimeVoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"UptimeVoteSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"UptimeVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePowerBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VotePowerBlockSelected\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"firstRewardEpochStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVotingEpochId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVotePowerBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_votePowerBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVoterRegistrationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVoterRegistrationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_newSigningPolicyHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signNewSigningPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"}],\"internalType\":\"structIFlareSystemsManager.NumberOfWeightBasedClaims[]\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_uptimeVoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"submitUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SystemManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemManagerMetaData.ABI instead.
var SystemManagerABI = SystemManagerMetaData.ABI

// SystemManager is an auto generated Go binding around an Ethereum contract.
type SystemManager struct {
	SystemManagerCaller     // Read-only binding to the contract
	SystemManagerTransactor // Write-only binding to the contract
	SystemManagerFilterer   // Log filterer for contract events
}

// SystemManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemManagerSession struct {
	Contract     *SystemManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemManagerCallerSession struct {
	Contract *SystemManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SystemManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemManagerTransactorSession struct {
	Contract     *SystemManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SystemManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemManagerRaw struct {
	Contract *SystemManager // Generic contract binding to access the raw methods on
}

// SystemManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemManagerCallerRaw struct {
	Contract *SystemManagerCaller // Generic read-only contract binding to access the raw methods on
}

// SystemManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemManagerTransactorRaw struct {
	Contract *SystemManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemManager creates a new instance of SystemManager, bound to a specific deployed contract.
func NewSystemManager(address common.Address, backend bind.ContractBackend) (*SystemManager, error) {
	contract, err := bindSystemManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemManager{SystemManagerCaller: SystemManagerCaller{contract: contract}, SystemManagerTransactor: SystemManagerTransactor{contract: contract}, SystemManagerFilterer: SystemManagerFilterer{contract: contract}}, nil
}

// NewSystemManagerCaller creates a new read-only instance of SystemManager, bound to a specific deployed contract.
func NewSystemManagerCaller(address common.Address, caller bind.ContractCaller) (*SystemManagerCaller, error) {
	contract, err := bindSystemManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemManagerCaller{contract: contract}, nil
}

// NewSystemManagerTransactor creates a new write-only instance of SystemManager, bound to a specific deployed contract.
func NewSystemManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemManagerTransactor, error) {
	contract, err := bindSystemManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemManagerTransactor{contract: contract}, nil
}

// NewSystemManagerFilterer creates a new log filterer instance of SystemManager, bound to a specific deployed contract.
func NewSystemManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemManagerFilterer, error) {
	contract, err := bindSystemManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemManagerFilterer{contract: contract}, nil
}

// bindSystemManager binds a generic wrapper to an already deployed contract.
func bindSystemManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SystemManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemManager *SystemManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemManager.Contract.SystemManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemManager *SystemManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemManager.Contract.SystemManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemManager *SystemManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemManager.Contract.SystemManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemManager *SystemManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemManager *SystemManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemManager *SystemManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemManager.Contract.contract.Transact(opts, method, params...)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_SystemManager *SystemManagerCaller) FirstRewardEpochStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "firstRewardEpochStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_SystemManager *SystemManagerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _SystemManager.Contract.FirstRewardEpochStartTs(&_SystemManager.CallOpts)
}

// FirstRewardEpochStartTs is a free data retrieval call binding the contract method 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (_SystemManager *SystemManagerCallerSession) FirstRewardEpochStartTs() (uint64, error) {
	return _SystemManager.Contract.FirstRewardEpochStartTs(&_SystemManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_SystemManager *SystemManagerCaller) FirstVotingRoundStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "firstVotingRoundStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_SystemManager *SystemManagerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _SystemManager.Contract.FirstVotingRoundStartTs(&_SystemManager.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_SystemManager *SystemManagerCallerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _SystemManager.Contract.FirstVotingRoundStartTs(&_SystemManager.CallOpts)
}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_SystemManager *SystemManagerCaller) GetCurrentRewardEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getCurrentRewardEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_SystemManager *SystemManagerSession) GetCurrentRewardEpoch() (*big.Int, error) {
	return _SystemManager.Contract.GetCurrentRewardEpoch(&_SystemManager.CallOpts)
}

// GetCurrentRewardEpoch is a free data retrieval call binding the contract method 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (_SystemManager *SystemManagerCallerSession) GetCurrentRewardEpoch() (*big.Int, error) {
	return _SystemManager.Contract.GetCurrentRewardEpoch(&_SystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_SystemManager *SystemManagerCaller) GetCurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getCurrentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_SystemManager *SystemManagerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _SystemManager.Contract.GetCurrentRewardEpochId(&_SystemManager.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (_SystemManager *SystemManagerCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _SystemManager.Contract.GetCurrentRewardEpochId(&_SystemManager.CallOpts)
}

// GetCurrentVotingEpochId is a free data retrieval call binding the contract method 0x4134520b.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (_SystemManager *SystemManagerCaller) GetCurrentVotingEpochId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getCurrentVotingEpochId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCurrentVotingEpochId is a free data retrieval call binding the contract method 0x4134520b.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (_SystemManager *SystemManagerSession) GetCurrentVotingEpochId() (uint32, error) {
	return _SystemManager.Contract.GetCurrentVotingEpochId(&_SystemManager.CallOpts)
}

// GetCurrentVotingEpochId is a free data retrieval call binding the contract method 0x4134520b.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (_SystemManager *SystemManagerCallerSession) GetCurrentVotingEpochId() (uint32, error) {
	return _SystemManager.Contract.GetCurrentVotingEpochId(&_SystemManager.CallOpts)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_SystemManager *SystemManagerCaller) GetSeed(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getSeed", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_SystemManager *SystemManagerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _SystemManager.Contract.GetSeed(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (_SystemManager *SystemManagerCallerSession) GetSeed(_rewardEpochId *big.Int) (*big.Int, error) {
	return _SystemManager.Contract.GetSeed(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_SystemManager *SystemManagerCaller) GetStartVotingRoundId(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint32, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getStartVotingRoundId", _rewardEpochId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_SystemManager *SystemManagerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _SystemManager.Contract.GetStartVotingRoundId(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetStartVotingRoundId is a free data retrieval call binding the contract method 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (_SystemManager *SystemManagerCallerSession) GetStartVotingRoundId(_rewardEpochId *big.Int) (uint32, error) {
	return _SystemManager.Contract.GetStartVotingRoundId(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_SystemManager *SystemManagerCaller) GetThreshold(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint16, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getThreshold", _rewardEpochId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_SystemManager *SystemManagerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _SystemManager.Contract.GetThreshold(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetThreshold is a free data retrieval call binding the contract method 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (_SystemManager *SystemManagerCallerSession) GetThreshold(_rewardEpochId *big.Int) (uint16, error) {
	return _SystemManager.Contract.GetThreshold(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_SystemManager *SystemManagerCaller) GetVotePowerBlock(opts *bind.CallOpts, _rewardEpochId *big.Int) (uint64, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getVotePowerBlock", _rewardEpochId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_SystemManager *SystemManagerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _SystemManager.Contract.GetVotePowerBlock(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetVotePowerBlock is a free data retrieval call binding the contract method 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (_SystemManager *SystemManagerCallerSession) GetVotePowerBlock(_rewardEpochId *big.Int) (uint64, error) {
	return _SystemManager.Contract.GetVotePowerBlock(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_SystemManager *SystemManagerCaller) GetVoterRegistrationData(opts *bind.CallOpts, _rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "getVoterRegistrationData", _rewardEpochId)

	outstruct := new(struct {
		VotePowerBlock *big.Int
		Enabled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VotePowerBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_SystemManager *SystemManagerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _SystemManager.Contract.GetVoterRegistrationData(&_SystemManager.CallOpts, _rewardEpochId)
}

// GetVoterRegistrationData is a free data retrieval call binding the contract method 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (_SystemManager *SystemManagerCallerSession) GetVoterRegistrationData(_rewardEpochId *big.Int) (struct {
	VotePowerBlock *big.Int
	Enabled        bool
}, error) {
	return _SystemManager.Contract.GetVoterRegistrationData(&_SystemManager.CallOpts, _rewardEpochId)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_SystemManager *SystemManagerCaller) IsVoterRegistrationEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "isVoterRegistrationEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_SystemManager *SystemManagerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _SystemManager.Contract.IsVoterRegistrationEnabled(&_SystemManager.CallOpts)
}

// IsVoterRegistrationEnabled is a free data retrieval call binding the contract method 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (_SystemManager *SystemManagerCallerSession) IsVoterRegistrationEnabled() (bool, error) {
	return _SystemManager.Contract.IsVoterRegistrationEnabled(&_SystemManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerCaller) RewardEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "rewardEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _SystemManager.Contract.RewardEpochDurationSeconds(&_SystemManager.CallOpts)
}

// RewardEpochDurationSeconds is a free data retrieval call binding the contract method 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerCallerSession) RewardEpochDurationSeconds() (uint64, error) {
	return _SystemManager.Contract.RewardEpochDurationSeconds(&_SystemManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerCaller) VotingEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemManager.contract.Call(opts, &out, "votingEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _SystemManager.Contract.VotingEpochDurationSeconds(&_SystemManager.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_SystemManager *SystemManagerCallerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _SystemManager.Contract.VotingEpochDurationSeconds(&_SystemManager.CallOpts)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactor) SignNewSigningPolicy(opts *bind.TransactOpts, _rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.contract.Transact(opts, "signNewSigningPolicy", _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignNewSigningPolicy(&_SystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignNewSigningPolicy is a paid mutator transaction binding the contract method 0x6b4c7bd6.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactorSession) SignNewSigningPolicy(_rewardEpochId *big.Int, _newSigningPolicyHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignNewSigningPolicy(&_SystemManager.TransactOpts, _rewardEpochId, _newSigningPolicyHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xc00a1a97.
//
// Solidity: function signRewards(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactor) SignRewards(opts *bind.TransactOpts, _rewardEpochId *big.Int, _noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.contract.Transact(opts, "signRewards", _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xc00a1a97.
//
// Solidity: function signRewards(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignRewards(&_SystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignRewards is a paid mutator transaction binding the contract method 0xc00a1a97.
//
// Solidity: function signRewards(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactorSession) SignRewards(_rewardEpochId *big.Int, _noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, _rewardsHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignRewards(&_SystemManager.TransactOpts, _rewardEpochId, _noOfWeightBasedClaims, _rewardsHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactor) SignUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.contract.Transact(opts, "signUptimeVote", _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignUptimeVote(&_SystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SignUptimeVote is a paid mutator transaction binding the contract method 0xdc5a4225.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactorSession) SignUptimeVote(_rewardEpochId *big.Int, _uptimeVoteHash [32]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SignUptimeVote(&_SystemManager.TransactOpts, _rewardEpochId, _uptimeVoteHash, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactor) SubmitUptimeVote(opts *bind.TransactOpts, _rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.contract.Transact(opts, "submitUptimeVote", _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SubmitUptimeVote(&_SystemManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SubmitUptimeVote is a paid mutator transaction binding the contract method 0x9dd6850f.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (_SystemManager *SystemManagerTransactorSession) SubmitUptimeVote(_rewardEpochId *big.Int, _nodeIds [][20]byte, _signature IFlareSystemsManagerSignature) (*types.Transaction, error) {
	return _SystemManager.Contract.SubmitUptimeVote(&_SystemManager.TransactOpts, _rewardEpochId, _nodeIds, _signature)
}

// SystemManagerRandomAcquisitionStartedIterator is returned from FilterRandomAcquisitionStarted and is used to iterate over the raw logs and unpacked data for RandomAcquisitionStarted events raised by the SystemManager contract.
type SystemManagerRandomAcquisitionStartedIterator struct {
	Event *SystemManagerRandomAcquisitionStarted // Event containing the contract specifics and raw log

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
func (it *SystemManagerRandomAcquisitionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerRandomAcquisitionStarted)
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
		it.Event = new(SystemManagerRandomAcquisitionStarted)
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
func (it *SystemManagerRandomAcquisitionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerRandomAcquisitionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerRandomAcquisitionStarted represents a RandomAcquisitionStarted event raised by the SystemManager contract.
type SystemManagerRandomAcquisitionStarted struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomAcquisitionStarted is a free log retrieval operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) FilterRandomAcquisitionStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*SystemManagerRandomAcquisitionStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerRandomAcquisitionStartedIterator{contract: _SystemManager.contract, event: "RandomAcquisitionStarted", logs: logs, sub: sub}, nil
}

// WatchRandomAcquisitionStarted is a free log subscription operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) WatchRandomAcquisitionStarted(opts *bind.WatchOpts, sink chan<- *SystemManagerRandomAcquisitionStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "RandomAcquisitionStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerRandomAcquisitionStarted)
				if err := _SystemManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
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

// ParseRandomAcquisitionStarted is a log parse operation binding the contract event 0xf9991783e5e480e42d9a54d3f35f4321857f8f0ebeb3742d326dce28b1126708.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) ParseRandomAcquisitionStarted(log types.Log) (*SystemManagerRandomAcquisitionStarted, error) {
	event := new(SystemManagerRandomAcquisitionStarted)
	if err := _SystemManager.contract.UnpackLog(event, "RandomAcquisitionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerRewardEpochStartedIterator is returned from FilterRewardEpochStarted and is used to iterate over the raw logs and unpacked data for RewardEpochStarted events raised by the SystemManager contract.
type SystemManagerRewardEpochStartedIterator struct {
	Event *SystemManagerRewardEpochStarted // Event containing the contract specifics and raw log

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
func (it *SystemManagerRewardEpochStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerRewardEpochStarted)
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
		it.Event = new(SystemManagerRewardEpochStarted)
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
func (it *SystemManagerRewardEpochStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerRewardEpochStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerRewardEpochStarted represents a RewardEpochStarted event raised by the SystemManager contract.
type SystemManagerRewardEpochStarted struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Timestamp          uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardEpochStarted is a free log retrieval operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) FilterRewardEpochStarted(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*SystemManagerRewardEpochStartedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerRewardEpochStartedIterator{contract: _SystemManager.contract, event: "RewardEpochStarted", logs: logs, sub: sub}, nil
}

// WatchRewardEpochStarted is a free log subscription operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) WatchRewardEpochStarted(opts *bind.WatchOpts, sink chan<- *SystemManagerRewardEpochStarted, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "RewardEpochStarted", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerRewardEpochStarted)
				if err := _SystemManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
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

// ParseRewardEpochStarted is a log parse operation binding the contract event 0x4abb62ab1e4c42a11b90e4e45b92af1274f74cc634b759518e8c99e000d8be6d.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) ParseRewardEpochStarted(log types.Log) (*SystemManagerRewardEpochStarted, error) {
	event := new(SystemManagerRewardEpochStarted)
	if err := _SystemManager.contract.UnpackLog(event, "RewardEpochStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerRewardsSignedIterator is returned from FilterRewardsSigned and is used to iterate over the raw logs and unpacked data for RewardsSigned events raised by the SystemManager contract.
type SystemManagerRewardsSignedIterator struct {
	Event *SystemManagerRewardsSigned // Event containing the contract specifics and raw log

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
func (it *SystemManagerRewardsSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerRewardsSigned)
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
		it.Event = new(SystemManagerRewardsSigned)
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
func (it *SystemManagerRewardsSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerRewardsSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerRewardsSigned represents a RewardsSigned event raised by the SystemManager contract.
type SystemManagerRewardsSigned struct {
	RewardEpochId         *big.Int
	SigningPolicyAddress  common.Address
	Voter                 common.Address
	RewardsHash           [32]byte
	NoOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims
	Timestamp             uint64
	ThresholdReached      bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRewardsSigned is a free log retrieval operation binding the contract event 0x81b5504045130d3b82498ff414ad58271e85bbde420cc85aa66d91eff9af30fb.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, (uint256,uint256)[] noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) FilterRewardsSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*SystemManagerRewardsSignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerRewardsSignedIterator{contract: _SystemManager.contract, event: "RewardsSigned", logs: logs, sub: sub}, nil
}

// WatchRewardsSigned is a free log subscription operation binding the contract event 0x81b5504045130d3b82498ff414ad58271e85bbde420cc85aa66d91eff9af30fb.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, (uint256,uint256)[] noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) WatchRewardsSigned(opts *bind.WatchOpts, sink chan<- *SystemManagerRewardsSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "RewardsSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerRewardsSigned)
				if err := _SystemManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
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

// ParseRewardsSigned is a log parse operation binding the contract event 0x81b5504045130d3b82498ff414ad58271e85bbde420cc85aa66d91eff9af30fb.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, (uint256,uint256)[] noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) ParseRewardsSigned(log types.Log) (*SystemManagerRewardsSigned, error) {
	event := new(SystemManagerRewardsSigned)
	if err := _SystemManager.contract.UnpackLog(event, "RewardsSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerSignUptimeVoteEnabledIterator is returned from FilterSignUptimeVoteEnabled and is used to iterate over the raw logs and unpacked data for SignUptimeVoteEnabled events raised by the SystemManager contract.
type SystemManagerSignUptimeVoteEnabledIterator struct {
	Event *SystemManagerSignUptimeVoteEnabled // Event containing the contract specifics and raw log

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
func (it *SystemManagerSignUptimeVoteEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerSignUptimeVoteEnabled)
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
		it.Event = new(SystemManagerSignUptimeVoteEnabled)
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
func (it *SystemManagerSignUptimeVoteEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerSignUptimeVoteEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerSignUptimeVoteEnabled represents a SignUptimeVoteEnabled event raised by the SystemManager contract.
type SystemManagerSignUptimeVoteEnabled struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignUptimeVoteEnabled is a free log retrieval operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) FilterSignUptimeVoteEnabled(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*SystemManagerSignUptimeVoteEnabledIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "SignUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerSignUptimeVoteEnabledIterator{contract: _SystemManager.contract, event: "SignUptimeVoteEnabled", logs: logs, sub: sub}, nil
}

// WatchSignUptimeVoteEnabled is a free log subscription operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) WatchSignUptimeVoteEnabled(opts *bind.WatchOpts, sink chan<- *SystemManagerSignUptimeVoteEnabled, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "SignUptimeVoteEnabled", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerSignUptimeVoteEnabled)
				if err := _SystemManager.contract.UnpackLog(event, "SignUptimeVoteEnabled", log); err != nil {
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

// ParseSignUptimeVoteEnabled is a log parse operation binding the contract event 0x235cef7d085c1e59545613282d239e56eb0cd056135aa46b8c658cf54a078561.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) ParseSignUptimeVoteEnabled(log types.Log) (*SystemManagerSignUptimeVoteEnabled, error) {
	event := new(SystemManagerSignUptimeVoteEnabled)
	if err := _SystemManager.contract.UnpackLog(event, "SignUptimeVoteEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerSigningPolicySignedIterator is returned from FilterSigningPolicySigned and is used to iterate over the raw logs and unpacked data for SigningPolicySigned events raised by the SystemManager contract.
type SystemManagerSigningPolicySignedIterator struct {
	Event *SystemManagerSigningPolicySigned // Event containing the contract specifics and raw log

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
func (it *SystemManagerSigningPolicySignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerSigningPolicySigned)
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
		it.Event = new(SystemManagerSigningPolicySigned)
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
func (it *SystemManagerSigningPolicySignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerSigningPolicySignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerSigningPolicySigned represents a SigningPolicySigned event raised by the SystemManager contract.
type SystemManagerSigningPolicySigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicySigned is a free log retrieval operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) FilterSigningPolicySigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*SystemManagerSigningPolicySignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerSigningPolicySignedIterator{contract: _SystemManager.contract, event: "SigningPolicySigned", logs: logs, sub: sub}, nil
}

// WatchSigningPolicySigned is a free log subscription operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) WatchSigningPolicySigned(opts *bind.WatchOpts, sink chan<- *SystemManagerSigningPolicySigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "SigningPolicySigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerSigningPolicySigned)
				if err := _SystemManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
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

// ParseSigningPolicySigned is a log parse operation binding the contract event 0x154b0214ae62d8a5548c1eac25fabd87c38b04932a217732e1022f3118da67f3.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) ParseSigningPolicySigned(log types.Log) (*SystemManagerSigningPolicySigned, error) {
	event := new(SystemManagerSigningPolicySigned)
	if err := _SystemManager.contract.UnpackLog(event, "SigningPolicySigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerUptimeVoteSignedIterator is returned from FilterUptimeVoteSigned and is used to iterate over the raw logs and unpacked data for UptimeVoteSigned events raised by the SystemManager contract.
type SystemManagerUptimeVoteSignedIterator struct {
	Event *SystemManagerUptimeVoteSigned // Event containing the contract specifics and raw log

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
func (it *SystemManagerUptimeVoteSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerUptimeVoteSigned)
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
		it.Event = new(SystemManagerUptimeVoteSigned)
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
func (it *SystemManagerUptimeVoteSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerUptimeVoteSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerUptimeVoteSigned represents a UptimeVoteSigned event raised by the SystemManager contract.
type SystemManagerUptimeVoteSigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	UptimeVoteHash       [32]byte
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUptimeVoteSigned is a free log retrieval operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) FilterUptimeVoteSigned(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*SystemManagerUptimeVoteSignedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerUptimeVoteSignedIterator{contract: _SystemManager.contract, event: "UptimeVoteSigned", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSigned is a free log subscription operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) WatchUptimeVoteSigned(opts *bind.WatchOpts, sink chan<- *SystemManagerUptimeVoteSigned, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "UptimeVoteSigned", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerUptimeVoteSigned)
				if err := _SystemManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
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

// ParseUptimeVoteSigned is a log parse operation binding the contract event 0x5506337d1266599f8b64675a1c8321701657ca2f2f70be0e0c58302b6c22e797.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (_SystemManager *SystemManagerFilterer) ParseUptimeVoteSigned(log types.Log) (*SystemManagerUptimeVoteSigned, error) {
	event := new(SystemManagerUptimeVoteSigned)
	if err := _SystemManager.contract.UnpackLog(event, "UptimeVoteSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerUptimeVoteSubmittedIterator is returned from FilterUptimeVoteSubmitted and is used to iterate over the raw logs and unpacked data for UptimeVoteSubmitted events raised by the SystemManager contract.
type SystemManagerUptimeVoteSubmittedIterator struct {
	Event *SystemManagerUptimeVoteSubmitted // Event containing the contract specifics and raw log

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
func (it *SystemManagerUptimeVoteSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerUptimeVoteSubmitted)
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
		it.Event = new(SystemManagerUptimeVoteSubmitted)
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
func (it *SystemManagerUptimeVoteSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerUptimeVoteSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerUptimeVoteSubmitted represents a UptimeVoteSubmitted event raised by the SystemManager contract.
type SystemManagerUptimeVoteSubmitted struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	NodeIds              [][20]byte
	Timestamp            uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUptimeVoteSubmitted is a free log retrieval operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) FilterUptimeVoteSubmitted(opts *bind.FilterOpts, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (*SystemManagerUptimeVoteSubmittedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerUptimeVoteSubmittedIterator{contract: _SystemManager.contract, event: "UptimeVoteSubmitted", logs: logs, sub: sub}, nil
}

// WatchUptimeVoteSubmitted is a free log subscription operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) WatchUptimeVoteSubmitted(opts *bind.WatchOpts, sink chan<- *SystemManagerUptimeVoteSubmitted, rewardEpochId []*big.Int, signingPolicyAddress []common.Address, voter []common.Address) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "UptimeVoteSubmitted", rewardEpochIdRule, signingPolicyAddressRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerUptimeVoteSubmitted)
				if err := _SystemManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
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

// ParseUptimeVoteSubmitted is a log parse operation binding the contract event 0xed370d61eb315e1d46d979894585530b99f94dab64c0d40366685aebe39e3db0.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) ParseUptimeVoteSubmitted(log types.Log) (*SystemManagerUptimeVoteSubmitted, error) {
	event := new(SystemManagerUptimeVoteSubmitted)
	if err := _SystemManager.contract.UnpackLog(event, "UptimeVoteSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemManagerVotePowerBlockSelectedIterator is returned from FilterVotePowerBlockSelected and is used to iterate over the raw logs and unpacked data for VotePowerBlockSelected events raised by the SystemManager contract.
type SystemManagerVotePowerBlockSelectedIterator struct {
	Event *SystemManagerVotePowerBlockSelected // Event containing the contract specifics and raw log

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
func (it *SystemManagerVotePowerBlockSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemManagerVotePowerBlockSelected)
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
		it.Event = new(SystemManagerVotePowerBlockSelected)
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
func (it *SystemManagerVotePowerBlockSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemManagerVotePowerBlockSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemManagerVotePowerBlockSelected represents a VotePowerBlockSelected event raised by the SystemManager contract.
type SystemManagerVotePowerBlockSelected struct {
	RewardEpochId  *big.Int
	VotePowerBlock uint64
	Timestamp      uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotePowerBlockSelected is a free log retrieval operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) FilterVotePowerBlockSelected(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*SystemManagerVotePowerBlockSelectedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.FilterLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &SystemManagerVotePowerBlockSelectedIterator{contract: _SystemManager.contract, event: "VotePowerBlockSelected", logs: logs, sub: sub}, nil
}

// WatchVotePowerBlockSelected is a free log subscription operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) WatchVotePowerBlockSelected(opts *bind.WatchOpts, sink chan<- *SystemManagerVotePowerBlockSelected, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _SystemManager.contract.WatchLogs(opts, "VotePowerBlockSelected", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemManagerVotePowerBlockSelected)
				if err := _SystemManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
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

// ParseVotePowerBlockSelected is a log parse operation binding the contract event 0xf21722dbe044a7cea0f6d81c871cae750971e36c9dd10999e46f2b46f26ac7ff.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (_SystemManager *SystemManagerFilterer) ParseVotePowerBlockSelected(log types.Log) (*SystemManagerVotePowerBlockSelected, error) {
	event := new(SystemManagerVotePowerBlockSelected)
	if err := _SystemManager.contract.UnpackLog(event, "VotePowerBlockSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
