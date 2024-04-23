// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voter_registry

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

// IVoterRegistrySignature is an auto generated low-level Go binding around an user-defined struct.
type IVoterRegistrySignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// VoterRegistryMetaData contains all meta data concerning the VoterRegistry contract.
var VoterRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"untilRewardEpochId\",\"type\":\"uint256\"}],\"name\":\"BeneficiaryChilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicKeyPart1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicKeyPart2\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registrationWeight\",\"type\":\"uint256\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"VoterRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_beneficiary\",\"type\":\"bytes20\"}],\"name\":\"chilledUntilRewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getNumberOfRegisteredVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"isVoterRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"newSigningPolicyInitializationStartBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicKeyRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIVoterRegistry.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VoterRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use VoterRegistryMetaData.ABI instead.
var VoterRegistryABI = VoterRegistryMetaData.ABI

// VoterRegistry is an auto generated Go binding around an Ethereum contract.
type VoterRegistry struct {
	VoterRegistryCaller     // Read-only binding to the contract
	VoterRegistryTransactor // Write-only binding to the contract
	VoterRegistryFilterer   // Log filterer for contract events
}

// VoterRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterRegistrySession struct {
	Contract     *VoterRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterRegistryCallerSession struct {
	Contract *VoterRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VoterRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterRegistryTransactorSession struct {
	Contract     *VoterRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VoterRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterRegistryRaw struct {
	Contract *VoterRegistry // Generic contract binding to access the raw methods on
}

// VoterRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterRegistryCallerRaw struct {
	Contract *VoterRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// VoterRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterRegistryTransactorRaw struct {
	Contract *VoterRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterRegistry creates a new instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistry(address common.Address, backend bind.ContractBackend) (*VoterRegistry, error) {
	contract, err := bindVoterRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterRegistry{VoterRegistryCaller: VoterRegistryCaller{contract: contract}, VoterRegistryTransactor: VoterRegistryTransactor{contract: contract}, VoterRegistryFilterer: VoterRegistryFilterer{contract: contract}}, nil
}

// NewVoterRegistryCaller creates a new read-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryCaller(address common.Address, caller bind.ContractCaller) (*VoterRegistryCaller, error) {
	contract, err := bindVoterRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryCaller{contract: contract}, nil
}

// NewVoterRegistryTransactor creates a new write-only instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterRegistryTransactor, error) {
	contract, err := bindVoterRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryTransactor{contract: contract}, nil
}

// NewVoterRegistryFilterer creates a new log filterer instance of VoterRegistry, bound to a specific deployed contract.
func NewVoterRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterRegistryFilterer, error) {
	contract, err := bindVoterRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryFilterer{contract: contract}, nil
}

// bindVoterRegistry binds a generic wrapper to an already deployed contract.
func bindVoterRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VoterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.VoterRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.VoterRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterRegistry *VoterRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VoterRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterRegistry *VoterRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterRegistry.Contract.contract.Transact(opts, method, params...)
}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x3c5cb76f.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (_VoterRegistry *VoterRegistryCaller) ChilledUntilRewardEpochId(opts *bind.CallOpts, _beneficiary [20]byte) (*big.Int, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "chilledUntilRewardEpochId", _beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x3c5cb76f.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (_VoterRegistry *VoterRegistrySession) ChilledUntilRewardEpochId(_beneficiary [20]byte) (*big.Int, error) {
	return _VoterRegistry.Contract.ChilledUntilRewardEpochId(&_VoterRegistry.CallOpts, _beneficiary)
}

// ChilledUntilRewardEpochId is a free data retrieval call binding the contract method 0x3c5cb76f.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (_VoterRegistry *VoterRegistryCallerSession) ChilledUntilRewardEpochId(_beneficiary [20]byte) (*big.Int, error) {
	return _VoterRegistry.Contract.ChilledUntilRewardEpochId(&_VoterRegistry.CallOpts, _beneficiary)
}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistryCaller) GetNumberOfRegisteredVoters(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "getNumberOfRegisteredVoters", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistrySession) GetNumberOfRegisteredVoters(_rewardEpochId *big.Int) (*big.Int, error) {
	return _VoterRegistry.Contract.GetNumberOfRegisteredVoters(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// GetNumberOfRegisteredVoters is a free data retrieval call binding the contract method 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistryCallerSession) GetNumberOfRegisteredVoters(_rewardEpochId *big.Int) (*big.Int, error) {
	return _VoterRegistry.Contract.GetNumberOfRegisteredVoters(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_VoterRegistry *VoterRegistryCaller) GetRegisteredVoters(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "getRegisteredVoters", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_VoterRegistry *VoterRegistrySession) GetRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _VoterRegistry.Contract.GetRegisteredVoters(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// GetRegisteredVoters is a free data retrieval call binding the contract method 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (_VoterRegistry *VoterRegistryCallerSession) GetRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _VoterRegistry.Contract.GetRegisteredVoters(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) IsVoterRegistered(opts *bind.CallOpts, _voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "isVoterRegistered", _voter, _rewardEpochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_VoterRegistry *VoterRegistrySession) IsVoterRegistered(_voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	return _VoterRegistry.Contract.IsVoterRegistered(&_VoterRegistry.CallOpts, _voter, _rewardEpochId)
}

// IsVoterRegistered is a free data retrieval call binding the contract method 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) IsVoterRegistered(_voter common.Address, _rewardEpochId *big.Int) (bool, error) {
	return _VoterRegistry.Contract.IsVoterRegistered(&_VoterRegistry.CallOpts, _voter, _rewardEpochId)
}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_VoterRegistry *VoterRegistryCaller) MaxVoters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "maxVoters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_VoterRegistry *VoterRegistrySession) MaxVoters() (*big.Int, error) {
	return _VoterRegistry.Contract.MaxVoters(&_VoterRegistry.CallOpts)
}

// MaxVoters is a free data retrieval call binding the contract method 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (_VoterRegistry *VoterRegistryCallerSession) MaxVoters() (*big.Int, error) {
	return _VoterRegistry.Contract.MaxVoters(&_VoterRegistry.CallOpts)
}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistryCaller) NewSigningPolicyInitializationStartBlockNumber(opts *bind.CallOpts, _rewardEpochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "newSigningPolicyInitializationStartBlockNumber", _rewardEpochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistrySession) NewSigningPolicyInitializationStartBlockNumber(_rewardEpochId *big.Int) (*big.Int, error) {
	return _VoterRegistry.Contract.NewSigningPolicyInitializationStartBlockNumber(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// NewSigningPolicyInitializationStartBlockNumber is a free data retrieval call binding the contract method 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (_VoterRegistry *VoterRegistryCallerSession) NewSigningPolicyInitializationStartBlockNumber(_rewardEpochId *big.Int) (*big.Int, error) {
	return _VoterRegistry.Contract.NewSigningPolicyInitializationStartBlockNumber(&_VoterRegistry.CallOpts, _rewardEpochId)
}

// PublicKeyRequired is a free data retrieval call binding the contract method 0x92e3e45f.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (_VoterRegistry *VoterRegistryCaller) PublicKeyRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VoterRegistry.contract.Call(opts, &out, "publicKeyRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicKeyRequired is a free data retrieval call binding the contract method 0x92e3e45f.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (_VoterRegistry *VoterRegistrySession) PublicKeyRequired() (bool, error) {
	return _VoterRegistry.Contract.PublicKeyRequired(&_VoterRegistry.CallOpts)
}

// PublicKeyRequired is a free data retrieval call binding the contract method 0x92e3e45f.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (_VoterRegistry *VoterRegistryCallerSession) PublicKeyRequired() (bool, error) {
	return _VoterRegistry.Contract.PublicKeyRequired(&_VoterRegistry.CallOpts)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_VoterRegistry *VoterRegistryTransactor) RegisterVoter(opts *bind.TransactOpts, _voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _VoterRegistry.contract.Transact(opts, "registerVoter", _voter, _signature)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_VoterRegistry *VoterRegistrySession) RegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RegisterVoter(&_VoterRegistry.TransactOpts, _voter, _signature)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x8f7d0957.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_VoterRegistry *VoterRegistryTransactorSession) RegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _VoterRegistry.Contract.RegisterVoter(&_VoterRegistry.TransactOpts, _voter, _signature)
}

// VoterRegistryBeneficiaryChilledIterator is returned from FilterBeneficiaryChilled and is used to iterate over the raw logs and unpacked data for BeneficiaryChilled events raised by the VoterRegistry contract.
type VoterRegistryBeneficiaryChilledIterator struct {
	Event *VoterRegistryBeneficiaryChilled // Event containing the contract specifics and raw log

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
func (it *VoterRegistryBeneficiaryChilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryBeneficiaryChilled)
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
		it.Event = new(VoterRegistryBeneficiaryChilled)
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
func (it *VoterRegistryBeneficiaryChilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryBeneficiaryChilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryBeneficiaryChilled represents a BeneficiaryChilled event raised by the VoterRegistry contract.
type VoterRegistryBeneficiaryChilled struct {
	Beneficiary        [20]byte
	UntilRewardEpochId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryChilled is a free log retrieval operation binding the contract event 0x0a5e087b026d8f1c57e75d9d0cb0394c2ad3535e7a15d97d553be80476274cd0.
//
// Solidity: event BeneficiaryChilled(bytes20 indexed beneficiary, uint256 untilRewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) FilterBeneficiaryChilled(opts *bind.FilterOpts, beneficiary [][20]byte) (*VoterRegistryBeneficiaryChilledIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "BeneficiaryChilled", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryBeneficiaryChilledIterator{contract: _VoterRegistry.contract, event: "BeneficiaryChilled", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryChilled is a free log subscription operation binding the contract event 0x0a5e087b026d8f1c57e75d9d0cb0394c2ad3535e7a15d97d553be80476274cd0.
//
// Solidity: event BeneficiaryChilled(bytes20 indexed beneficiary, uint256 untilRewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) WatchBeneficiaryChilled(opts *bind.WatchOpts, sink chan<- *VoterRegistryBeneficiaryChilled, beneficiary [][20]byte) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "BeneficiaryChilled", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryBeneficiaryChilled)
				if err := _VoterRegistry.contract.UnpackLog(event, "BeneficiaryChilled", log); err != nil {
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

// ParseBeneficiaryChilled is a log parse operation binding the contract event 0x0a5e087b026d8f1c57e75d9d0cb0394c2ad3535e7a15d97d553be80476274cd0.
//
// Solidity: event BeneficiaryChilled(bytes20 indexed beneficiary, uint256 untilRewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) ParseBeneficiaryChilled(log types.Log) (*VoterRegistryBeneficiaryChilled, error) {
	event := new(VoterRegistryBeneficiaryChilled)
	if err := _VoterRegistry.contract.UnpackLog(event, "BeneficiaryChilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryVoterRegisteredIterator is returned from FilterVoterRegistered and is used to iterate over the raw logs and unpacked data for VoterRegistered events raised by the VoterRegistry contract.
type VoterRegistryVoterRegisteredIterator struct {
	Event *VoterRegistryVoterRegistered // Event containing the contract specifics and raw log

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
func (it *VoterRegistryVoterRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryVoterRegistered)
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
		it.Event = new(VoterRegistryVoterRegistered)
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
func (it *VoterRegistryVoterRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryVoterRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryVoterRegistered represents a VoterRegistered event raised by the VoterRegistry contract.
type VoterRegistryVoterRegistered struct {
	Voter                   common.Address
	RewardEpochId           *big.Int
	SigningPolicyAddress    common.Address
	SubmitAddress           common.Address
	SubmitSignaturesAddress common.Address
	PublicKeyPart1          [32]byte
	PublicKeyPart2          [32]byte
	RegistrationWeight      *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistered is a free log retrieval operation binding the contract event 0x824bc2cc10bfe21ead60b8c8a90716eb325b9335aa73eaede799abf38fce062c.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address submitAddress, address submitSignaturesAddress, bytes32 publicKeyPart1, bytes32 publicKeyPart2, uint256 registrationWeight)
func (_VoterRegistry *VoterRegistryFilterer) FilterVoterRegistered(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int, signingPolicyAddress []common.Address) (*VoterRegistryVoterRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "VoterRegistered", voterRule, rewardEpochIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryVoterRegisteredIterator{contract: _VoterRegistry.contract, event: "VoterRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterRegistered is a free log subscription operation binding the contract event 0x824bc2cc10bfe21ead60b8c8a90716eb325b9335aa73eaede799abf38fce062c.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address submitAddress, address submitSignaturesAddress, bytes32 publicKeyPart1, bytes32 publicKeyPart2, uint256 registrationWeight)
func (_VoterRegistry *VoterRegistryFilterer) WatchVoterRegistered(opts *bind.WatchOpts, sink chan<- *VoterRegistryVoterRegistered, voter []common.Address, rewardEpochId []*big.Int, signingPolicyAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "VoterRegistered", voterRule, rewardEpochIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryVoterRegistered)
				if err := _VoterRegistry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
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

// ParseVoterRegistered is a log parse operation binding the contract event 0x824bc2cc10bfe21ead60b8c8a90716eb325b9335aa73eaede799abf38fce062c.
//
// Solidity: event VoterRegistered(address indexed voter, uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address submitAddress, address submitSignaturesAddress, bytes32 publicKeyPart1, bytes32 publicKeyPart2, uint256 registrationWeight)
func (_VoterRegistry *VoterRegistryFilterer) ParseVoterRegistered(log types.Log) (*VoterRegistryVoterRegistered, error) {
	event := new(VoterRegistryVoterRegistered)
	if err := _VoterRegistry.contract.UnpackLog(event, "VoterRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoterRegistryVoterRemovedIterator is returned from FilterVoterRemoved and is used to iterate over the raw logs and unpacked data for VoterRemoved events raised by the VoterRegistry contract.
type VoterRegistryVoterRemovedIterator struct {
	Event *VoterRegistryVoterRemoved // Event containing the contract specifics and raw log

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
func (it *VoterRegistryVoterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterRegistryVoterRemoved)
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
		it.Event = new(VoterRegistryVoterRemoved)
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
func (it *VoterRegistryVoterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterRegistryVoterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterRegistryVoterRemoved represents a VoterRemoved event raised by the VoterRegistry contract.
type VoterRegistryVoterRemoved struct {
	Voter         common.Address
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoterRemoved is a free log retrieval operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) FilterVoterRemoved(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int) (*VoterRegistryVoterRemovedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _VoterRegistry.contract.FilterLogs(opts, "VoterRemoved", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &VoterRegistryVoterRemovedIterator{contract: _VoterRegistry.contract, event: "VoterRemoved", logs: logs, sub: sub}, nil
}

// WatchVoterRemoved is a free log subscription operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) WatchVoterRemoved(opts *bind.WatchOpts, sink chan<- *VoterRegistryVoterRemoved, voter []common.Address, rewardEpochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _VoterRegistry.contract.WatchLogs(opts, "VoterRemoved", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterRegistryVoterRemoved)
				if err := _VoterRegistry.contract.UnpackLog(event, "VoterRemoved", log); err != nil {
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

// ParseVoterRemoved is a log parse operation binding the contract event 0x98a7f87f8e2aa2f23f43769eff67782bb12946384b142d1ce1e8e38e05d9a3e6.
//
// Solidity: event VoterRemoved(address indexed voter, uint256 indexed rewardEpochId)
func (_VoterRegistry *VoterRegistryFilterer) ParseVoterRemoved(log types.Log) (*VoterRegistryVoterRemoved, error) {
	event := new(VoterRegistryVoterRemoved)
	if err := _VoterRegistry.contract.UnpackLog(event, "VoterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
