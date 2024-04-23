// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mock

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

// FlareSystemMockPolicy is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemMockPolicy struct {
	Pk1    [32]byte
	Pk2    [32]byte
	Weight uint16
}

// IFtsoFeedPublisherFeed is an auto generated low-level Go binding around an user-defined struct.
type IFtsoFeedPublisherFeed struct {
	VotingRoundId uint32
	Id            [21]byte
	Value         int32
	TurnoutBIPS   uint16
	Decimals      int8
}

// MockMetaData contains all meta data concerning the Mock contract.
var MockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_randomSeed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_epochLen\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"epochLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21\",\"name\":\"\",\"type\":\"bytes21\"}],\"name\":\"getCurrentFeed\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"votingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"bytes21\",\"name\":\"id\",\"type\":\"bytes21\"},{\"internalType\":\"int32\",\"name\":\"value\",\"type\":\"int32\"},{\"internalType\":\"uint16\",\"name\":\"turnoutBIPS\",\"type\":\"uint16\"},{\"internalType\":\"int8\",\"name\":\"decimals\",\"type\":\"int8\"}],\"internalType\":\"structIFtsoFeedPublisher.Feed\",\"name\":\"_feed\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKeyAndNormalisedWeight\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_publicKeyPart1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_publicKeyPart2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeight\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_normalisedWeightsSumOfVotersWithPublicKeys\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRandom\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"policies\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"pk1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pk2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"pk1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"pk2\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"weight\",\"type\":\"uint16\"}],\"internalType\":\"structFlareSystemMock.Policy\",\"name\":\"_policy\",\"type\":\"tuple\"}],\"name\":\"registerAsVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalWeights\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106de3803806106de83398101604081905261002f9161003d565b600091909155600155610061565b6000806040838503121561005057600080fd5b505080516020909101519092909150565b61066e806100706000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806381bd4bf41161006657806381bd4bf4146101ae5780639b97b24c146101c3578063b62e1efc1461028a578063d2b3996f146102c1578063e0d4ea37146102ca57600080fd5b80630b747d911461009857806344b571d9146100b457806358fb1039146101395780637056269714610192575b600080fd5b6100a160005481565b6040519081526020015b60405180910390f35b61010f6100c2366004610477565b60008281526002602081815260408084206001600160a01b039095168452938152838320805460018201549190930154958452600390915292909120549093919261ffff90811692911690565b60408051948552602085019390935261ffff918216928401929092521660608201526080016100ab565b610173610147366004610477565b6002602081815260009384526040808520909152918352912080546001820154919092015461ffff1683565b60408051938452602084019290925261ffff16908201526060016100ab565b61019a6102dd565b60405162ffffff90911681526020016100ab565b6101c16101bc3660046104a3565b6102f2565b005b61022f6101d13660046104eb565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915250506040805160a08101825260008082526020820152620186a09181019190915261271060608201526002608082015290565b6040516100ab9190815163ffffffff1681526020808301516affffffffffffffffffffff19169082015260408083015160030b9082015260608083015161ffff169082015260809182015160000b9181019190915260a00190565b6102ae610298366004610520565b60036020526000908152604090205461ffff1681565b60405161ffff90911681526020016100ab565b6100a160015481565b6100a16102d8366004610520565b6103d4565b6000600154436102ed9190610539565b905090565b610302606082016040830161056e565b61ffff166000036103525760405162461bcd60e51b8152602060048201526016602482015275576569676874206d757374206265206e6f6e7a65726f60501b604482015260640160405180910390fd5b60008381526002602090815260408083206001600160a01b038616845290915290208190610380828261058b565b506103939050606082016040830161056e565b600084815260036020526040812080549091906103b590849061ffff166105c0565b92506101000a81548161ffff021916908361ffff160217905550505050565b60006002600154436103e69190610539565b60005460408051602081019390935282015260600160408051601f1981840301815290829052610415916105f0565b602060405180830381855afa158015610432573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250810190610455919061061f565b92915050565b80356001600160a01b038116811461047257600080fd5b919050565b6000806040838503121561048a57600080fd5b8235915061049a6020840161045b565b90509250929050565b600080600083850360a08112156104b957600080fd5b843593506104c96020860161045b565b92506060603f19820112156104dd57600080fd5b506040840190509250925092565b6000602082840312156104fd57600080fd5b81356affffffffffffffffffffff198116811461051957600080fd5b9392505050565b60006020828403121561053257600080fd5b5035919050565b60008261055657634e487b7160e01b600052601260045260246000fd5b500490565b61ffff8116811461056b57600080fd5b50565b60006020828403121561058057600080fd5b81356105198161055b565b81358155602082013560018201556002810160408301356105ab8161055b565b815461ffff191661ffff919091161790555050565b61ffff8181168382160190808211156105e957634e487b7160e01b600052601160045260246000fd5b5092915050565b6000825160005b8181101561061157602081860181015185830152016105f7565b506000920191825250919050565b60006020828403121561063157600080fd5b505191905056fea2646970667358221220988d211aac67deeaf4b0a8b724eb30649d4bd70118ebe1b94df98912381aefab64736f6c63430008140033",
}

// MockABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMetaData.ABI instead.
var MockABI = MockMetaData.ABI

// MockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockMetaData.Bin instead.
var MockBin = MockMetaData.Bin

// DeployMock deploys a new Ethereum contract, binding an instance of Mock to it.
func DeployMock(auth *bind.TransactOpts, backend bind.ContractBackend, _randomSeed *big.Int, _epochLen *big.Int) (common.Address, *types.Transaction, *Mock, error) {
	parsed, err := MockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockBin), backend, _randomSeed, _epochLen)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// Mock is an auto generated Go binding around an Ethereum contract.
type Mock struct {
	MockCaller     // Read-only binding to the contract
	MockTransactor // Write-only binding to the contract
	MockFilterer   // Log filterer for contract events
}

// MockCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockSession struct {
	Contract     *Mock             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockCallerSession struct {
	Contract *MockCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockTransactorSession struct {
	Contract     *MockTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockRaw struct {
	Contract *Mock // Generic contract binding to access the raw methods on
}

// MockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockCallerRaw struct {
	Contract *MockCaller // Generic read-only contract binding to access the raw methods on
}

// MockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockTransactorRaw struct {
	Contract *MockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMock creates a new instance of Mock, bound to a specific deployed contract.
func NewMock(address common.Address, backend bind.ContractBackend) (*Mock, error) {
	contract, err := bindMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mock{MockCaller: MockCaller{contract: contract}, MockTransactor: MockTransactor{contract: contract}, MockFilterer: MockFilterer{contract: contract}}, nil
}

// NewMockCaller creates a new read-only instance of Mock, bound to a specific deployed contract.
func NewMockCaller(address common.Address, caller bind.ContractCaller) (*MockCaller, error) {
	contract, err := bindMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockCaller{contract: contract}, nil
}

// NewMockTransactor creates a new write-only instance of Mock, bound to a specific deployed contract.
func NewMockTransactor(address common.Address, transactor bind.ContractTransactor) (*MockTransactor, error) {
	contract, err := bindMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockTransactor{contract: contract}, nil
}

// NewMockFilterer creates a new log filterer instance of Mock, bound to a specific deployed contract.
func NewMockFilterer(address common.Address, filterer bind.ContractFilterer) (*MockFilterer, error) {
	contract, err := bindMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockFilterer{contract: contract}, nil
}

// bindMock binds a generic wrapper to an already deployed contract.
func bindMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.MockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.MockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mock *MockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mock *MockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mock *MockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mock.Contract.contract.Transact(opts, method, params...)
}

// EpochLen is a free data retrieval call binding the contract method 0xd2b3996f.
//
// Solidity: function epochLen() view returns(uint256)
func (_Mock *MockCaller) EpochLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "epochLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLen is a free data retrieval call binding the contract method 0xd2b3996f.
//
// Solidity: function epochLen() view returns(uint256)
func (_Mock *MockSession) EpochLen() (*big.Int, error) {
	return _Mock.Contract.EpochLen(&_Mock.CallOpts)
}

// EpochLen is a free data retrieval call binding the contract method 0xd2b3996f.
//
// Solidity: function epochLen() view returns(uint256)
func (_Mock *MockCallerSession) EpochLen() (*big.Int, error) {
	return _Mock.Contract.EpochLen(&_Mock.CallOpts)
}

// GetCurrentFeed is a free data retrieval call binding the contract method 0x9b97b24c.
//
// Solidity: function getCurrentFeed(bytes21 ) pure returns((uint32,bytes21,int32,uint16,int8) _feed)
func (_Mock *MockCaller) GetCurrentFeed(opts *bind.CallOpts, arg0 [21]byte) (IFtsoFeedPublisherFeed, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "getCurrentFeed", arg0)

	if err != nil {
		return *new(IFtsoFeedPublisherFeed), err
	}

	out0 := *abi.ConvertType(out[0], new(IFtsoFeedPublisherFeed)).(*IFtsoFeedPublisherFeed)

	return out0, err

}

// GetCurrentFeed is a free data retrieval call binding the contract method 0x9b97b24c.
//
// Solidity: function getCurrentFeed(bytes21 ) pure returns((uint32,bytes21,int32,uint16,int8) _feed)
func (_Mock *MockSession) GetCurrentFeed(arg0 [21]byte) (IFtsoFeedPublisherFeed, error) {
	return _Mock.Contract.GetCurrentFeed(&_Mock.CallOpts, arg0)
}

// GetCurrentFeed is a free data retrieval call binding the contract method 0x9b97b24c.
//
// Solidity: function getCurrentFeed(bytes21 ) pure returns((uint32,bytes21,int32,uint16,int8) _feed)
func (_Mock *MockCallerSession) GetCurrentFeed(arg0 [21]byte) (IFtsoFeedPublisherFeed, error) {
	return _Mock.Contract.GetCurrentFeed(&_Mock.CallOpts, arg0)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
func (_Mock *MockCaller) GetCurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "getCurrentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
func (_Mock *MockSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _Mock.Contract.GetCurrentRewardEpochId(&_Mock.CallOpts)
}

// GetCurrentRewardEpochId is a free data retrieval call binding the contract method 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24 _currentRewardEpochId)
func (_Mock *MockCallerSession) GetCurrentRewardEpochId() (*big.Int, error) {
	return _Mock.Contract.GetCurrentRewardEpochId(&_Mock.CallOpts)
}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Mock *MockCaller) GetPublicKeyAndNormalisedWeight(opts *bind.CallOpts, _rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "getPublicKeyAndNormalisedWeight", _rewardEpochId, _signingPolicyAddress)

	outstruct := new(struct {
		PublicKeyPart1                             [32]byte
		PublicKeyPart2                             [32]byte
		NormalisedWeight                           uint16
		NormalisedWeightsSumOfVotersWithPublicKeys uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKeyPart1 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.PublicKeyPart2 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.NormalisedWeight = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.NormalisedWeightsSumOfVotersWithPublicKeys = *abi.ConvertType(out[3], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Mock *MockSession) GetPublicKeyAndNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Mock.Contract.GetPublicKeyAndNormalisedWeight(&_Mock.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetPublicKeyAndNormalisedWeight is a free data retrieval call binding the contract method 0x44b571d9.
//
// Solidity: function getPublicKeyAndNormalisedWeight(uint256 _rewardEpochId, address _signingPolicyAddress) view returns(bytes32 _publicKeyPart1, bytes32 _publicKeyPart2, uint16 _normalisedWeight, uint16 _normalisedWeightsSumOfVotersWithPublicKeys)
func (_Mock *MockCallerSession) GetPublicKeyAndNormalisedWeight(_rewardEpochId *big.Int, _signingPolicyAddress common.Address) (struct {
	PublicKeyPart1                             [32]byte
	PublicKeyPart2                             [32]byte
	NormalisedWeight                           uint16
	NormalisedWeightsSumOfVotersWithPublicKeys uint16
}, error) {
	return _Mock.Contract.GetPublicKeyAndNormalisedWeight(&_Mock.CallOpts, _rewardEpochId, _signingPolicyAddress)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 ) view returns(uint256 _currentRandom)
func (_Mock *MockCaller) GetSeed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "getSeed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 ) view returns(uint256 _currentRandom)
func (_Mock *MockSession) GetSeed(arg0 *big.Int) (*big.Int, error) {
	return _Mock.Contract.GetSeed(&_Mock.CallOpts, arg0)
}

// GetSeed is a free data retrieval call binding the contract method 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 ) view returns(uint256 _currentRandom)
func (_Mock *MockCallerSession) GetSeed(arg0 *big.Int) (*big.Int, error) {
	return _Mock.Contract.GetSeed(&_Mock.CallOpts, arg0)
}

// Policies is a free data retrieval call binding the contract method 0x58fb1039.
//
// Solidity: function policies(uint256 , address ) view returns(bytes32 pk1, bytes32 pk2, uint16 weight)
func (_Mock *MockCaller) Policies(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Pk1    [32]byte
	Pk2    [32]byte
	Weight uint16
}, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "policies", arg0, arg1)

	outstruct := new(struct {
		Pk1    [32]byte
		Pk2    [32]byte
		Weight uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pk1 = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Pk2 = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Weight = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// Policies is a free data retrieval call binding the contract method 0x58fb1039.
//
// Solidity: function policies(uint256 , address ) view returns(bytes32 pk1, bytes32 pk2, uint16 weight)
func (_Mock *MockSession) Policies(arg0 *big.Int, arg1 common.Address) (struct {
	Pk1    [32]byte
	Pk2    [32]byte
	Weight uint16
}, error) {
	return _Mock.Contract.Policies(&_Mock.CallOpts, arg0, arg1)
}

// Policies is a free data retrieval call binding the contract method 0x58fb1039.
//
// Solidity: function policies(uint256 , address ) view returns(bytes32 pk1, bytes32 pk2, uint16 weight)
func (_Mock *MockCallerSession) Policies(arg0 *big.Int, arg1 common.Address) (struct {
	Pk1    [32]byte
	Pk2    [32]byte
	Weight uint16
}, error) {
	return _Mock.Contract.Policies(&_Mock.CallOpts, arg0, arg1)
}

// RandomSeed is a free data retrieval call binding the contract method 0x0b747d91.
//
// Solidity: function randomSeed() view returns(uint256)
func (_Mock *MockCaller) RandomSeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "randomSeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomSeed is a free data retrieval call binding the contract method 0x0b747d91.
//
// Solidity: function randomSeed() view returns(uint256)
func (_Mock *MockSession) RandomSeed() (*big.Int, error) {
	return _Mock.Contract.RandomSeed(&_Mock.CallOpts)
}

// RandomSeed is a free data retrieval call binding the contract method 0x0b747d91.
//
// Solidity: function randomSeed() view returns(uint256)
func (_Mock *MockCallerSession) RandomSeed() (*big.Int, error) {
	return _Mock.Contract.RandomSeed(&_Mock.CallOpts)
}

// TotalWeights is a free data retrieval call binding the contract method 0xb62e1efc.
//
// Solidity: function totalWeights(uint256 ) view returns(uint16)
func (_Mock *MockCaller) TotalWeights(opts *bind.CallOpts, arg0 *big.Int) (uint16, error) {
	var out []interface{}
	err := _Mock.contract.Call(opts, &out, "totalWeights", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TotalWeights is a free data retrieval call binding the contract method 0xb62e1efc.
//
// Solidity: function totalWeights(uint256 ) view returns(uint16)
func (_Mock *MockSession) TotalWeights(arg0 *big.Int) (uint16, error) {
	return _Mock.Contract.TotalWeights(&_Mock.CallOpts, arg0)
}

// TotalWeights is a free data retrieval call binding the contract method 0xb62e1efc.
//
// Solidity: function totalWeights(uint256 ) view returns(uint16)
func (_Mock *MockCallerSession) TotalWeights(arg0 *big.Int) (uint16, error) {
	return _Mock.Contract.TotalWeights(&_Mock.CallOpts, arg0)
}

// RegisterAsVoter is a paid mutator transaction binding the contract method 0x81bd4bf4.
//
// Solidity: function registerAsVoter(uint256 _epoch, address _sender, (bytes32,bytes32,uint16) _policy) returns()
func (_Mock *MockTransactor) RegisterAsVoter(opts *bind.TransactOpts, _epoch *big.Int, _sender common.Address, _policy FlareSystemMockPolicy) (*types.Transaction, error) {
	return _Mock.contract.Transact(opts, "registerAsVoter", _epoch, _sender, _policy)
}

// RegisterAsVoter is a paid mutator transaction binding the contract method 0x81bd4bf4.
//
// Solidity: function registerAsVoter(uint256 _epoch, address _sender, (bytes32,bytes32,uint16) _policy) returns()
func (_Mock *MockSession) RegisterAsVoter(_epoch *big.Int, _sender common.Address, _policy FlareSystemMockPolicy) (*types.Transaction, error) {
	return _Mock.Contract.RegisterAsVoter(&_Mock.TransactOpts, _epoch, _sender, _policy)
}

// RegisterAsVoter is a paid mutator transaction binding the contract method 0x81bd4bf4.
//
// Solidity: function registerAsVoter(uint256 _epoch, address _sender, (bytes32,bytes32,uint16) _policy) returns()
func (_Mock *MockTransactorSession) RegisterAsVoter(_epoch *big.Int, _sender common.Address, _policy FlareSystemMockPolicy) (*types.Transaction, error) {
	return _Mock.Contract.RegisterAsVoter(&_Mock.TransactOpts, _epoch, _sender, _policy)
}
