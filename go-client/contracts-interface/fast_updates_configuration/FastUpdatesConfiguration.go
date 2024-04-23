// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fast_updates_configuration

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

// IFastUpdatesConfigurationFeedConfiguration is an auto generated low-level Go binding around an user-defined struct.
type IFastUpdatesConfigurationFeedConfiguration struct {
	FeedId          [21]byte
	RewardBandValue uint32
	InflationShare  *big.Int
}

// FastUpdatesConfigurationMetaData contains all meta data concerning the FastUpdatesConfiguration contract.
var FastUpdatesConfigurationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FeedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"}],\"internalType\":\"structIFastUpdatesConfiguration.FeedConfiguration[]\",\"name\":\"_feedConfigs\",\"type\":\"tuple[]\"}],\"name\":\"addFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdater\",\"outputs\":[{\"internalType\":\"contractIIFastUpdater\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedConfigurations\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"}],\"internalType\":\"structIFastUpdatesConfiguration.FeedConfiguration[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedConfigurationsBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_feedIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_rewardBandValues\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_inflationShares\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getFeedId\",\"outputs\":[{\"internalType\":\"bytes21\",\"name\":\"_feedId\",\"type\":\"bytes21\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeedIds\",\"outputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21\",\"name\":\"_feedId\",\"type\":\"bytes21\"}],\"name\":\"getFeedIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfFeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnusedIndices\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"}],\"name\":\"removeFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"}],\"internalType\":\"structIFastUpdatesConfiguration.FeedConfiguration[]\",\"name\":\"_feedConfigs\",\"type\":\"tuple[]\"}],\"name\":\"updateFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620027f4380380620027f483398101604081905262000034916200020b565b80838362000043828262000079565b506200006f9050817f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e771955565b505050506200025f565b600054600160a01b900460ff1615620000d95760405162461bcd60e51b815260206004820152601460248201527f696e697469616c6973656420213d2066616c736500000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b038216620001315760405162461bcd60e51b815260206004820152601860248201527f676f7665726e616e63652073657474696e6773207a65726f00000000000000006044820152606401620000d0565b6001600160a01b0381166200017c5760405162461bcd60e51b815260206004820152601060248201526f5f676f7665726e616e6365207a65726f60801b6044820152606401620000d0565b600080546001600160a01b038481166001600160a81b031990921691909117600160a01b17909155600180549183166001600160a01b0319909216821790556040519081527f9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db9060200160405180910390a15050565b6001600160a01b03811681146200020857600080fd5b50565b6000806000606084860312156200022157600080fd5b83516200022e81620001f2565b60208501519093506200024181620001f2565b60408501519092506200025481620001f2565b809150509250925092565b612585806200026f6000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806374e6310e116100b8578063c906b1b41161007c578063c906b1b4146102da578063d29a4fa9146102e2578063debfda30146102f5578063e17f212e14610318578063ef88bf131461032c578063f5a983831461033f57600080fd5b806374e6310e146102545780639310283614610275578063a087d184146102a1578063a59b2c46146102b4578063b00c0b76146102c757600080fd5b806348a924821161010a57806348a92482146101c15780635267a15d146101d85780635aa6e675146102135780635ff270791461021b57806362354e031461022e57806367fc40291461024157600080fd5b80630a9cabe7146101475780630c518dce1461016d578063247c9cf71461018257806331038aad1461019757806331864f1f146101ac575b600080fd5b61015a610155366004611c81565b610347565b6040519081526020015b60405180910390f35b61017561039b565b6040516101649190611c9e565b610195610190366004611cec565b61045b565b005b61019f6107da565b6040516101649190611d61565b6101b4610868565b6040516101649190611dcf565b6101c96108c0565b60405161016493929190611e57565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e7719545b6040516001600160a01b039091168152602001610164565b6101fb610bcf565b610195610229366004611e9a565b610c6b565b6000546101fb906001600160a01b031681565b61019561024f366004611e9a565b610ef1565b610267610262366004611e9a565b610fd2565b604051610164929190611ec4565b610288610283366004611ee5565b611077565b6040516001600160581b03199091168152602001610164565b6101956102af366004611cec565b6110e3565b6101956102c2366004611efe565b6112db565b6101956102d5366004612055565b61155b565b60045461015a565b6003546101fb906001600160a01b031681565b61030861030336600461210e565b611636565b6040519015158152602001610164565b60005461030890600160a81b900460ff1681565b61019561033a36600461212b565b6116bf565b610195611825565b6001600160581b031981166000908152600560205260408120549081900361038a5760405162461bcd60e51b815260040161038190612164565b60405180910390fd5b80610394816121a7565b9392505050565b6004546060908067ffffffffffffffff8111156103ba576103ba611f61565b6040519080825280602002602001820160405280156103e3578160200160208202803683370190505b50915060005b818110156104565760048181548110610404576104046121be565b600091825260209091200154835160589190911b9084908390811061042b5761042b6121be565b6001600160581b0319909216602092830291909101909101528061044e816121d4565b9150506103e9565b505090565b600054600160b01b900460ff168061047d5750600054600160a81b900460ff16155b156107cb5761048a6118eb565b60008167ffffffffffffffff8111156104a5576104a5611f61565b6040519080825280602002602001820160405280156104ce578160200160208202803683370190505b50905060005b828110156107635760008484838181106104f0576104f06121be565b6105069260206060909202019081019150611c81565b90506001600160581b031981166105515760405162461bcd60e51b815260206004820152600f60248201526e1a5b9d985b1a590819995959081a59608a1b6044820152606401610381565b6001600160581b03198116600090815260056020526040902054156105ae5760405162461bcd60e51b81526020600482015260136024820152726665656420616c72656164792065786973747360681b6044820152606401610381565b6006546000901561061257600680546105c9906001906121ed565b815481106105d9576105d96121be565b9060005260206000200154905060068054806105f7576105f7612200565b60019003818190600052602060002001600090559055610624565b50600480546001810182556000919091525b80848481518110610637576106376121be565b602002602001018181525050858584818110610655576106556121be565b9050606002016004828154811061066e5761066e6121be565b9060005260206000200181816106849190612239565b5061069290508160016122c0565b6001600160581b031983166000818152600560205260409020919091557f3ca318c85958cdc1745f9edcd68164b4579efa8050b27b9d634f5e0427e7e33a8787868181106106e2576106e26121be565b90506060020160200160208101906106fa91906122d3565b88888781811061070c5761070c6121be565b905060600201604001602081019061072491906122f0565b6040805163ffffffff909316835262ffffff9091166020830152810184905260600160405180910390a25050808061075b906121d4565b9150506104d4565b506003546040516363f921db60e01b81526001600160a01b03909116906363f921db90610794908490600401611dcf565b600060405180830381600087803b1580156107ae57600080fd5b505af11580156107c2573d6000803e3d6000fd5b50505050505050565b6107d6600036611925565b5050565b60606004805480602002602001604051908101604052809291908181526020016000905b8282101561085f5760008481526020908190206040805160608101825291850154605881901b6001600160581b0319168352600160a81b810463ffffffff1683850152600160c81b900462ffffff16908201528252600190920191016107fe565b50505050905090565b606060068054806020026020016040519081016040528092919081815260200182805480156108b657602002820191906000526020600020905b8154815260200190600101908083116108a2575b5050505050905090565b600454606090819081906108d581601561230d565b67ffffffffffffffff8111156108ed576108ed611f61565b6040519080825280601f01601f191660200182016040528015610917576020820181803683370190505b50935061092581600461230d565b67ffffffffffffffff81111561093d5761093d611f61565b6040519080825280601f01601f191660200182016040528015610967576020820181803683370190505b50925061097581600361230d565b67ffffffffffffffff81111561098d5761098d611f61565b6040519080825280601f01601f1916602001820160405280156109b7576020820181803683370190505b50915060005b81811015610bc8576000600482815481106109da576109da6121be565b60009182526020822001546004805460589290921b93509084908110610a0257610a026121be565b9060005260206000200160000160159054906101000a900463ffffffff1660e01b9050600060048481548110610a3a57610a3a6121be565b6000918252602082200154600160c81b900460e81b6001600160e81b03191691505b6015811015610acb57838160158110610a7757610a776121be565b1a60f81b8982610a8888601561230d565b610a9291906122c0565b81518110610aa257610aa26121be565b60200101906001600160f81b031916908160001a90535080610ac3816121d4565b915050610a5c565b5060005b6004811015610b3e57828160048110610aea57610aea6121be565b1a60f81b8882610afb88600461230d565b610b0591906122c0565b81518110610b1557610b156121be565b60200101906001600160f81b031916908160001a90535080610b36816121d4565b915050610acf565b5060005b6003811015610bb157818160038110610b5d57610b5d6121be565b1a60f81b8782610b6e88600361230d565b610b7891906122c0565b81518110610b8857610b886121be565b60200101906001600160f81b031916908160001a90535080610ba9816121d4565b915050610b42565b505050508080610bc0906121d4565b9150506109bd565b5050909192565b60008054600160a81b900460ff16610bf157506001546001600160a01b031690565b60008054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c669190612324565b905090565b610c7433611636565b610cb05760405162461bcd60e51b815260206004820152600d60248201526c37b7363c9032bc32b1baba37b960991b6044820152606401610381565b6001600160e01b0319811660009081526002602052604081208054909103610d1a5760405162461bcd60e51b815260206004820152601a60248201527f74696d656c6f636b3a20696e76616c69642073656c6563746f720000000000006044820152606401610381565b8054421015610d6b5760405162461bcd60e51b815260206004820152601960248201527f74696d656c6f636b3a206e6f7420616c6c6f77656420796574000000000000006044820152606401610381565b6000816001018054610d7c90612341565b80601f0160208091040260200160405190810160405280929190818152602001828054610da890612341565b8015610df55780601f10610dca57610100808354040283529160200191610df5565b820191906000526020600020905b815481529060010190602001808311610dd857829003601f168201915b505050506001600160e01b0319851660009081526002602052604081208181559293509050610e276001830182611c15565b50506000805460ff60b01b1916600160b01b1781556040513090610e4c90849061237b565b6000604051808303816000865af19150503d8060008114610e89576040519150601f19603f3d011682016040523d82523d6000602084013e610e8e565b606091505b50506000805460ff60b01b19169055604080516001600160e01b0319871681524260208201529192507fa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438910160405180910390a1610eeb81611a71565b50505050565b610ef9611a8e565b6001600160e01b031981166000908152600260205260408120549003610f615760405162461bcd60e51b815260206004820152601a60248201527f74696d656c6f636b3a20696e76616c69642073656c6563746f720000000000006044820152606401610381565b604080516001600160e01b0319831681524260208201527f7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8910160405180910390a16001600160e01b03198116600090815260026020526040812081815590610fcd6001830182611c15565b505050565b60026020526000908152604090208054600182018054919291610ff490612341565b80601f016020809104026020016040519081016040528092919081815260200182805461102090612341565b801561106d5780601f106110425761010080835404028352916020019161106d565b820191906000526020600020905b81548152906001019060200180831161105057829003601f168201915b5050505050905082565b60045460009082106110bb5760405162461bcd60e51b815260206004820152600d60248201526c0d2dcecc2d8d2c840d2dcc8caf609b1b6044820152606401610381565b600482815481106110ce576110ce6121be565b60009182526020909120015460581b92915050565b600054600160b01b900460ff16806111055750600054600160a81b900460ff16155b156107cb576111126118eb565b60005b81811015610fcd576000838383818110611131576111316121be565b6111479260206060909202019081019150611c81565b90506001600160581b031981166111925760405162461bcd60e51b815260206004820152600f60248201526e1a5b9d985b1a590819995959081a59608a1b6044820152606401610381565b6001600160581b03198116600090815260056020526040812054908190036111cc5760405162461bcd60e51b815260040161038190612164565b806111d6816121a7565b9150508484848181106111eb576111eb6121be565b90506060020160048281548110611204576112046121be565b90600052602060002001818161121a9190612239565b50506001600160581b031982167f317c7e03c79b6fcd22d6f17813b4e8e8a4a14104fdfb79431c6c73b550c7ca9d86868681811061125a5761125a6121be565b905060600201602001602081019061127291906122d3565b878787818110611284576112846121be565b905060600201604001602081019061129c91906122f0565b6040805163ffffffff909316835262ffffff9091166020830152810184905260600160405180910390a2505080806112d3906121d4565b915050611115565b600054600160b01b900460ff16806112fd5750600054600160a81b900460ff16155b156107cb5761130a6118eb565b60008167ffffffffffffffff81111561132557611325611f61565b60405190808252806020026020018201604052801561134e578160200160208202803683370190505b50905060005b8281101561152a57600060056000868685818110611374576113746121be565b90506020020160208101906113899190611c81565b6affffffffffffffffffffff19166affffffffffffffffffffff19168152602001908152602001600020549050806000036113d65760405162461bcd60e51b815260040161038190612164565b806113e0816121a7565b915050808383815181106113f6576113f66121be565b6020908102919091010152600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01819055600480548290811061144a5761144a6121be565b6000918252602082200180546001600160e01b0319169055600590868685818110611477576114776121be565b905060200201602081019061148c9190611c81565b6001600160581b031916815260208101919091526040016000908120558484838181106114bb576114bb6121be565b90506020020160208101906114d09190611c81565b6affffffffffffffffffffff19167fbb4bc8e9bdadd13a82544df890de25d2c6403cd23a7655410eb2ad4f542425ab8260405161150f91815260200190565b60405180910390a25080611522816121d4565b915050611354565b50600354604051630abfaf1760e41b81526001600160a01b039091169063abfaf17090610794908490600401611dcf565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e7719546001600160a01b0316336001600160a01b0316146115d45760405162461bcd60e51b815260206004820152601460248201527337b7363c9030b2323932b9b9903ab83230ba32b960611b6044820152606401610381565b61162c61160883836040518060400160405280600e81526020016d20b2323932b9b9aab83230ba32b960911b815250611ae8565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e771955565b6107d68282611bc3565b60008054600160a01b900460ff1680156116b95750600054604051630debfda360e41b81526001600160a01b0384811660048301529091169063debfda3090602401602060405180830381865afa158015611695573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b99190612397565b92915050565b600054600160a01b900460ff16156117105760405162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b6044820152606401610381565b6001600160a01b0382166117665760405162461bcd60e51b815260206004820152601860248201527f676f7665726e616e63652073657474696e6773207a65726f00000000000000006044820152606401610381565b6001600160a01b0381166117af5760405162461bcd60e51b815260206004820152601060248201526f5f676f7665726e616e6365207a65726f60801b6044820152606401610381565b600080546001600160a01b038481166001600160a81b031990921691909117600160a01b17909155600180549183166001600160a01b0319909216821790556040519081527f9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db9060200160405180910390a15050565b61182d611a8e565b600054600160a81b900460ff16156118875760405162461bcd60e51b815260206004820152601a60248201527f616c726561647920696e2070726f64756374696f6e206d6f64650000000000006044820152606401610381565b600180546001600160a01b031916905560008054600160a81b60ff60a81b198216179091556040516001600160a01b0390911681527f83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c9060200160405180910390a1565b600054600160b01b900460ff161561191b5733301461190c5761190c6123b9565b6000805460ff60b01b19169055565b611923611a8e565b565b61192d611a8e565b6000805460408051636221a54b60e01b81529051853593926001600160a01b031691636221a54b9160048083019260209291908290030181865afa158015611979573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061199d91906123cf565b905060006119ab82426122c0565b9050604051806040016040528082815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509390945250506001600160e01b03198616815260026020908152604090912083518155908301519091506001820190611a299082612436565b509050507fed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b683828787604051611a6294939291906124f6565b60405180910390a15050505050565b3d604051818101604052816000823e8215611a8a578181f35b8181fd5b611a96610bcf565b6001600160a01b0316336001600160a01b0316146119235760405162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b6044820152606401610381565b60008082604051602001611afc919061253c565b6040516020818303038152906040528051906020012090506000805b8651811015611b7457868181518110611b3357611b336121be565b60200260200101518303611b6257858181518110611b5357611b536121be565b60200260200101519150611b74565b80611b6c816121d4565b915050611b18565b506001600160a01b038116611bba5760405162461bcd60e51b815260206004820152600c60248201526b61646472657373207a65726f60a01b6044820152606401610381565b95945050505050565b611bf182826040518060400160405280600b81526020016a2330b9ba2ab83230ba32b960a91b815250611ae8565b600380546001600160a01b0319166001600160a01b03929092169190911790555050565b508054611c2190612341565b6000825580601f10611c31575050565b601f016020900490600052602060002090810190611c4f9190611c52565b50565b5b80821115611c675760008155600101611c53565b5090565b6001600160581b031981168114611c4f57600080fd5b600060208284031215611c9357600080fd5b813561039481611c6b565b6020808252825182820181905260009190848201906040850190845b81811015611ce05783516001600160581b03191683529284019291840191600101611cba565b50909695505050505050565b60008060208385031215611cff57600080fd5b823567ffffffffffffffff80821115611d1757600080fd5b818501915085601f830112611d2b57600080fd5b813581811115611d3a57600080fd5b866020606083028501011115611d4f57600080fd5b60209290920196919550909350505050565b602080825282518282018190526000919060409081850190868401855b82811015611dc257815180516001600160581b03191685528681015163ffffffff168786015285015162ffffff168585015260609093019290850190600101611d7e565b5091979650505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611ce057835183529284019291840191600101611deb565b60005b83811015611e22578181015183820152602001611e0a565b50506000910152565b60008151808452611e43816020860160208601611e07565b601f01601f19169290920160200192915050565b606081526000611e6a6060830186611e2b565b8281036020840152611e7c8186611e2b565b90508281036040840152611e908185611e2b565b9695505050505050565b600060208284031215611eac57600080fd5b81356001600160e01b03198116811461039457600080fd5b828152604060208201526000611edd6040830184611e2b565b949350505050565b600060208284031215611ef757600080fd5b5035919050565b60008060208385031215611f1157600080fd5b823567ffffffffffffffff80821115611f2957600080fd5b818501915085601f830112611f3d57600080fd5b813581811115611f4c57600080fd5b8660208260051b8501011115611d4f57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611fa057611fa0611f61565b604052919050565b600067ffffffffffffffff821115611fc257611fc2611f61565b5060051b60200190565b6001600160a01b0381168114611c4f57600080fd5b600082601f830112611ff257600080fd5b8135602061200761200283611fa8565b611f77565b82815260059290921b8401810191818101908684111561202657600080fd5b8286015b8481101561204a57803561203d81611fcc565b835291830191830161202a565b509695505050505050565b6000806040838503121561206857600080fd5b823567ffffffffffffffff8082111561208057600080fd5b818501915085601f83011261209457600080fd5b813560206120a461200283611fa8565b82815260059290921b840181019181810190898411156120c357600080fd5b948201945b838610156120e1578535825294820194908201906120c8565b965050860135925050808211156120f757600080fd5b5061210485828601611fe1565b9150509250929050565b60006020828403121561212057600080fd5b813561039481611fcc565b6000806040838503121561213e57600080fd5b823561214981611fcc565b9150602083013561215981611fcc565b809150509250929050565b6020808252601390820152721999595908191bd95cc81b9bdd08195e1a5cdd606a1b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000816121b6576121b6612191565b506000190190565b634e487b7160e01b600052603260045260246000fd5b6000600182016121e6576121e6612191565b5060010190565b818103818111156116b9576116b9612191565b634e487b7160e01b600052603160045260246000fd5b63ffffffff81168114611c4f57600080fd5b62ffffff81168114611c4f57600080fd5b813561224481611c6b565b81546001600160a81b0319811660589290921c9182178355602084013561226a81612216565b63ffffffff60a81b60a89190911b166001600160c81b03198216831781178455604085013561229881612228565b6001600160e01b0319929092169092179190911760c89190911b62ffffff60c81b1617905550565b808201808211156116b9576116b9612191565b6000602082840312156122e557600080fd5b813561039481612216565b60006020828403121561230257600080fd5b813561039481612228565b80820281158282048414176116b9576116b9612191565b60006020828403121561233657600080fd5b815161039481611fcc565b600181811c9082168061235557607f821691505b60208210810361237557634e487b7160e01b600052602260045260246000fd5b50919050565b6000825161238d818460208701611e07565b9190910192915050565b6000602082840312156123a957600080fd5b8151801515811461039457600080fd5b634e487b7160e01b600052600160045260246000fd5b6000602082840312156123e157600080fd5b5051919050565b601f821115610fcd57600081815260208120601f850160051c8101602086101561240f5750805b601f850160051c820191505b8181101561242e5782815560010161241b565b505050505050565b815167ffffffffffffffff81111561245057612450611f61565b6124648161245e8454612341565b846123e8565b602080601f83116001811461249957600084156124815750858301515b600019600386901b1c1916600185901b17855561242e565b600085815260208120601f198616915b828110156124c8578886015182559484019460019091019084016124a9565b50858210156124e65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b63ffffffff60e01b8516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f909201601f191601019392505050565b6020815260006103946020830184611e2b56fea2646970667358221220b08b7ff7fb5a3d9e52089197b82b82f21e80c777cd96ba19a88731aecbdbe75f64736f6c63430008140033",
}

// FastUpdatesConfigurationABI is the input ABI used to generate the binding from.
// Deprecated: Use FastUpdatesConfigurationMetaData.ABI instead.
var FastUpdatesConfigurationABI = FastUpdatesConfigurationMetaData.ABI

// FastUpdatesConfigurationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FastUpdatesConfigurationMetaData.Bin instead.
var FastUpdatesConfigurationBin = FastUpdatesConfigurationMetaData.Bin

// DeployFastUpdatesConfiguration deploys a new Ethereum contract, binding an instance of FastUpdatesConfiguration to it.
func DeployFastUpdatesConfiguration(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (common.Address, *types.Transaction, *FastUpdatesConfiguration, error) {
	parsed, err := FastUpdatesConfigurationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FastUpdatesConfigurationBin), backend, _governanceSettings, _initialGovernance, _addressUpdater)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FastUpdatesConfiguration{FastUpdatesConfigurationCaller: FastUpdatesConfigurationCaller{contract: contract}, FastUpdatesConfigurationTransactor: FastUpdatesConfigurationTransactor{contract: contract}, FastUpdatesConfigurationFilterer: FastUpdatesConfigurationFilterer{contract: contract}}, nil
}

// FastUpdatesConfiguration is an auto generated Go binding around an Ethereum contract.
type FastUpdatesConfiguration struct {
	FastUpdatesConfigurationCaller     // Read-only binding to the contract
	FastUpdatesConfigurationTransactor // Write-only binding to the contract
	FastUpdatesConfigurationFilterer   // Log filterer for contract events
}

// FastUpdatesConfigurationCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastUpdatesConfigurationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastUpdatesConfigurationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastUpdatesConfigurationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastUpdatesConfigurationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastUpdatesConfigurationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastUpdatesConfigurationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastUpdatesConfigurationSession struct {
	Contract     *FastUpdatesConfiguration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FastUpdatesConfigurationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastUpdatesConfigurationCallerSession struct {
	Contract *FastUpdatesConfigurationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// FastUpdatesConfigurationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastUpdatesConfigurationTransactorSession struct {
	Contract     *FastUpdatesConfigurationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// FastUpdatesConfigurationRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastUpdatesConfigurationRaw struct {
	Contract *FastUpdatesConfiguration // Generic contract binding to access the raw methods on
}

// FastUpdatesConfigurationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastUpdatesConfigurationCallerRaw struct {
	Contract *FastUpdatesConfigurationCaller // Generic read-only contract binding to access the raw methods on
}

// FastUpdatesConfigurationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastUpdatesConfigurationTransactorRaw struct {
	Contract *FastUpdatesConfigurationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastUpdatesConfiguration creates a new instance of FastUpdatesConfiguration, bound to a specific deployed contract.
func NewFastUpdatesConfiguration(address common.Address, backend bind.ContractBackend) (*FastUpdatesConfiguration, error) {
	contract, err := bindFastUpdatesConfiguration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfiguration{FastUpdatesConfigurationCaller: FastUpdatesConfigurationCaller{contract: contract}, FastUpdatesConfigurationTransactor: FastUpdatesConfigurationTransactor{contract: contract}, FastUpdatesConfigurationFilterer: FastUpdatesConfigurationFilterer{contract: contract}}, nil
}

// NewFastUpdatesConfigurationCaller creates a new read-only instance of FastUpdatesConfiguration, bound to a specific deployed contract.
func NewFastUpdatesConfigurationCaller(address common.Address, caller bind.ContractCaller) (*FastUpdatesConfigurationCaller, error) {
	contract, err := bindFastUpdatesConfiguration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationCaller{contract: contract}, nil
}

// NewFastUpdatesConfigurationTransactor creates a new write-only instance of FastUpdatesConfiguration, bound to a specific deployed contract.
func NewFastUpdatesConfigurationTransactor(address common.Address, transactor bind.ContractTransactor) (*FastUpdatesConfigurationTransactor, error) {
	contract, err := bindFastUpdatesConfiguration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationTransactor{contract: contract}, nil
}

// NewFastUpdatesConfigurationFilterer creates a new log filterer instance of FastUpdatesConfiguration, bound to a specific deployed contract.
func NewFastUpdatesConfigurationFilterer(address common.Address, filterer bind.ContractFilterer) (*FastUpdatesConfigurationFilterer, error) {
	contract, err := bindFastUpdatesConfiguration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationFilterer{contract: contract}, nil
}

// bindFastUpdatesConfiguration binds a generic wrapper to an already deployed contract.
func bindFastUpdatesConfiguration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastUpdatesConfigurationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastUpdatesConfiguration.Contract.FastUpdatesConfigurationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.FastUpdatesConfigurationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.FastUpdatesConfigurationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastUpdatesConfiguration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.contract.Transact(opts, method, params...)
}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) FastUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "fastUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) FastUpdater() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.FastUpdater(&_FastUpdatesConfiguration.CallOpts)
}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) FastUpdater() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.FastUpdater(&_FastUpdatesConfiguration.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetAddressUpdater() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.GetAddressUpdater(&_FastUpdatesConfiguration.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.GetAddressUpdater(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedConfigurations is a free data retrieval call binding the contract method 0x31038aad.
//
// Solidity: function getFeedConfigurations() view returns((bytes21,uint32,uint24)[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetFeedConfigurations(opts *bind.CallOpts) ([]IFastUpdatesConfigurationFeedConfiguration, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getFeedConfigurations")

	if err != nil {
		return *new([]IFastUpdatesConfigurationFeedConfiguration), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFastUpdatesConfigurationFeedConfiguration)).(*[]IFastUpdatesConfigurationFeedConfiguration)

	return out0, err

}

// GetFeedConfigurations is a free data retrieval call binding the contract method 0x31038aad.
//
// Solidity: function getFeedConfigurations() view returns((bytes21,uint32,uint24)[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetFeedConfigurations() ([]IFastUpdatesConfigurationFeedConfiguration, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedConfigurations(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedConfigurations is a free data retrieval call binding the contract method 0x31038aad.
//
// Solidity: function getFeedConfigurations() view returns((bytes21,uint32,uint24)[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetFeedConfigurations() ([]IFastUpdatesConfigurationFeedConfiguration, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedConfigurations(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedConfigurationsBytes is a free data retrieval call binding the contract method 0x48a92482.
//
// Solidity: function getFeedConfigurationsBytes() view returns(bytes _feedIds, bytes _rewardBandValues, bytes _inflationShares)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetFeedConfigurationsBytes(opts *bind.CallOpts) (struct {
	FeedIds          []byte
	RewardBandValues []byte
	InflationShares  []byte
}, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getFeedConfigurationsBytes")

	outstruct := new(struct {
		FeedIds          []byte
		RewardBandValues []byte
		InflationShares  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeedIds = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.RewardBandValues = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.InflationShares = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetFeedConfigurationsBytes is a free data retrieval call binding the contract method 0x48a92482.
//
// Solidity: function getFeedConfigurationsBytes() view returns(bytes _feedIds, bytes _rewardBandValues, bytes _inflationShares)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetFeedConfigurationsBytes() (struct {
	FeedIds          []byte
	RewardBandValues []byte
	InflationShares  []byte
}, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedConfigurationsBytes(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedConfigurationsBytes is a free data retrieval call binding the contract method 0x48a92482.
//
// Solidity: function getFeedConfigurationsBytes() view returns(bytes _feedIds, bytes _rewardBandValues, bytes _inflationShares)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetFeedConfigurationsBytes() (struct {
	FeedIds          []byte
	RewardBandValues []byte
	InflationShares  []byte
}, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedConfigurationsBytes(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedId is a free data retrieval call binding the contract method 0x93102836.
//
// Solidity: function getFeedId(uint256 _index) view returns(bytes21 _feedId)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetFeedId(opts *bind.CallOpts, _index *big.Int) ([21]byte, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getFeedId", _index)

	if err != nil {
		return *new([21]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([21]byte)).(*[21]byte)

	return out0, err

}

// GetFeedId is a free data retrieval call binding the contract method 0x93102836.
//
// Solidity: function getFeedId(uint256 _index) view returns(bytes21 _feedId)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetFeedId(_index *big.Int) ([21]byte, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedId(&_FastUpdatesConfiguration.CallOpts, _index)
}

// GetFeedId is a free data retrieval call binding the contract method 0x93102836.
//
// Solidity: function getFeedId(uint256 _index) view returns(bytes21 _feedId)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetFeedId(_index *big.Int) ([21]byte, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedId(&_FastUpdatesConfiguration.CallOpts, _index)
}

// GetFeedIds is a free data retrieval call binding the contract method 0x0c518dce.
//
// Solidity: function getFeedIds() view returns(bytes21[] _feedIds)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetFeedIds(opts *bind.CallOpts) ([][21]byte, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getFeedIds")

	if err != nil {
		return *new([][21]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][21]byte)).(*[][21]byte)

	return out0, err

}

// GetFeedIds is a free data retrieval call binding the contract method 0x0c518dce.
//
// Solidity: function getFeedIds() view returns(bytes21[] _feedIds)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetFeedIds() ([][21]byte, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedIds(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedIds is a free data retrieval call binding the contract method 0x0c518dce.
//
// Solidity: function getFeedIds() view returns(bytes21[] _feedIds)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetFeedIds() ([][21]byte, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedIds(&_FastUpdatesConfiguration.CallOpts)
}

// GetFeedIndex is a free data retrieval call binding the contract method 0x0a9cabe7.
//
// Solidity: function getFeedIndex(bytes21 _feedId) view returns(uint256 _index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetFeedIndex(opts *bind.CallOpts, _feedId [21]byte) (*big.Int, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getFeedIndex", _feedId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeedIndex is a free data retrieval call binding the contract method 0x0a9cabe7.
//
// Solidity: function getFeedIndex(bytes21 _feedId) view returns(uint256 _index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetFeedIndex(_feedId [21]byte) (*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedIndex(&_FastUpdatesConfiguration.CallOpts, _feedId)
}

// GetFeedIndex is a free data retrieval call binding the contract method 0x0a9cabe7.
//
// Solidity: function getFeedIndex(bytes21 _feedId) view returns(uint256 _index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetFeedIndex(_feedId [21]byte) (*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetFeedIndex(&_FastUpdatesConfiguration.CallOpts, _feedId)
}

// GetNumberOfFeeds is a free data retrieval call binding the contract method 0xc906b1b4.
//
// Solidity: function getNumberOfFeeds() view returns(uint256)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetNumberOfFeeds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getNumberOfFeeds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfFeeds is a free data retrieval call binding the contract method 0xc906b1b4.
//
// Solidity: function getNumberOfFeeds() view returns(uint256)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetNumberOfFeeds() (*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetNumberOfFeeds(&_FastUpdatesConfiguration.CallOpts)
}

// GetNumberOfFeeds is a free data retrieval call binding the contract method 0xc906b1b4.
//
// Solidity: function getNumberOfFeeds() view returns(uint256)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetNumberOfFeeds() (*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetNumberOfFeeds(&_FastUpdatesConfiguration.CallOpts)
}

// GetUnusedIndices is a free data retrieval call binding the contract method 0x31864f1f.
//
// Solidity: function getUnusedIndices() view returns(uint256[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GetUnusedIndices(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "getUnusedIndices")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUnusedIndices is a free data retrieval call binding the contract method 0x31864f1f.
//
// Solidity: function getUnusedIndices() view returns(uint256[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GetUnusedIndices() ([]*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetUnusedIndices(&_FastUpdatesConfiguration.CallOpts)
}

// GetUnusedIndices is a free data retrieval call binding the contract method 0x31864f1f.
//
// Solidity: function getUnusedIndices() view returns(uint256[])
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GetUnusedIndices() ([]*big.Int, error) {
	return _FastUpdatesConfiguration.Contract.GetUnusedIndices(&_FastUpdatesConfiguration.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) Governance() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.Governance(&_FastUpdatesConfiguration.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) Governance() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.Governance(&_FastUpdatesConfiguration.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) GovernanceSettings() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.GovernanceSettings(&_FastUpdatesConfiguration.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) GovernanceSettings() (common.Address, error) {
	return _FastUpdatesConfiguration.Contract.GovernanceSettings(&_FastUpdatesConfiguration.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) IsExecutor(_address common.Address) (bool, error) {
	return _FastUpdatesConfiguration.Contract.IsExecutor(&_FastUpdatesConfiguration.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FastUpdatesConfiguration.Contract.IsExecutor(&_FastUpdatesConfiguration.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) ProductionMode() (bool, error) {
	return _FastUpdatesConfiguration.Contract.ProductionMode(&_FastUpdatesConfiguration.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) ProductionMode() (bool, error) {
	return _FastUpdatesConfiguration.Contract.ProductionMode(&_FastUpdatesConfiguration.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FastUpdatesConfiguration.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FastUpdatesConfiguration.Contract.TimelockedCalls(&_FastUpdatesConfiguration.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FastUpdatesConfiguration.Contract.TimelockedCalls(&_FastUpdatesConfiguration.CallOpts, selector)
}

// AddFeeds is a paid mutator transaction binding the contract method 0x247c9cf7.
//
// Solidity: function addFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) AddFeeds(opts *bind.TransactOpts, _feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "addFeeds", _feedConfigs)
}

// AddFeeds is a paid mutator transaction binding the contract method 0x247c9cf7.
//
// Solidity: function addFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) AddFeeds(_feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.AddFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedConfigs)
}

// AddFeeds is a paid mutator transaction binding the contract method 0x247c9cf7.
//
// Solidity: function addFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) AddFeeds(_feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.AddFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedConfigs)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.CancelGovernanceCall(&_FastUpdatesConfiguration.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.CancelGovernanceCall(&_FastUpdatesConfiguration.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.ExecuteGovernanceCall(&_FastUpdatesConfiguration.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.ExecuteGovernanceCall(&_FastUpdatesConfiguration.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.Initialise(&_FastUpdatesConfiguration.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.Initialise(&_FastUpdatesConfiguration.TransactOpts, _governanceSettings, _initialGovernance)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xa59b2c46.
//
// Solidity: function removeFeeds(bytes21[] _feedIds) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) RemoveFeeds(opts *bind.TransactOpts, _feedIds [][21]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "removeFeeds", _feedIds)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xa59b2c46.
//
// Solidity: function removeFeeds(bytes21[] _feedIds) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) RemoveFeeds(_feedIds [][21]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.RemoveFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedIds)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xa59b2c46.
//
// Solidity: function removeFeeds(bytes21[] _feedIds) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) RemoveFeeds(_feedIds [][21]byte) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.RemoveFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedIds)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.SwitchToProductionMode(&_FastUpdatesConfiguration.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.SwitchToProductionMode(&_FastUpdatesConfiguration.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.UpdateContractAddresses(&_FastUpdatesConfiguration.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.UpdateContractAddresses(&_FastUpdatesConfiguration.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateFeeds is a paid mutator transaction binding the contract method 0xa087d184.
//
// Solidity: function updateFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactor) UpdateFeeds(opts *bind.TransactOpts, _feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.contract.Transact(opts, "updateFeeds", _feedConfigs)
}

// UpdateFeeds is a paid mutator transaction binding the contract method 0xa087d184.
//
// Solidity: function updateFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationSession) UpdateFeeds(_feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.UpdateFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedConfigs)
}

// UpdateFeeds is a paid mutator transaction binding the contract method 0xa087d184.
//
// Solidity: function updateFeeds((bytes21,uint32,uint24)[] _feedConfigs) returns()
func (_FastUpdatesConfiguration *FastUpdatesConfigurationTransactorSession) UpdateFeeds(_feedConfigs []IFastUpdatesConfigurationFeedConfiguration) (*types.Transaction, error) {
	return _FastUpdatesConfiguration.Contract.UpdateFeeds(&_FastUpdatesConfiguration.TransactOpts, _feedConfigs)
}

// FastUpdatesConfigurationFeedAddedIterator is returned from FilterFeedAdded and is used to iterate over the raw logs and unpacked data for FeedAdded events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedAddedIterator struct {
	Event *FastUpdatesConfigurationFeedAdded // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationFeedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationFeedAdded)
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
		it.Event = new(FastUpdatesConfigurationFeedAdded)
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
func (it *FastUpdatesConfigurationFeedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationFeedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationFeedAdded represents a FeedAdded event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedAdded struct {
	FeedId          [21]byte
	RewardBandValue uint32
	InflationShare  *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeedAdded is a free log retrieval operation binding the contract event 0x3ca318c85958cdc1745f9edcd68164b4579efa8050b27b9d634f5e0427e7e33a.
//
// Solidity: event FeedAdded(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterFeedAdded(opts *bind.FilterOpts, feedId [][21]byte) (*FastUpdatesConfigurationFeedAddedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "FeedAdded", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationFeedAddedIterator{contract: _FastUpdatesConfiguration.contract, event: "FeedAdded", logs: logs, sub: sub}, nil
}

// WatchFeedAdded is a free log subscription operation binding the contract event 0x3ca318c85958cdc1745f9edcd68164b4579efa8050b27b9d634f5e0427e7e33a.
//
// Solidity: event FeedAdded(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchFeedAdded(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationFeedAdded, feedId [][21]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "FeedAdded", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationFeedAdded)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedAdded", log); err != nil {
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

// ParseFeedAdded is a log parse operation binding the contract event 0x3ca318c85958cdc1745f9edcd68164b4579efa8050b27b9d634f5e0427e7e33a.
//
// Solidity: event FeedAdded(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseFeedAdded(log types.Log) (*FastUpdatesConfigurationFeedAdded, error) {
	event := new(FastUpdatesConfigurationFeedAdded)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationFeedRemovedIterator is returned from FilterFeedRemoved and is used to iterate over the raw logs and unpacked data for FeedRemoved events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedRemovedIterator struct {
	Event *FastUpdatesConfigurationFeedRemoved // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationFeedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationFeedRemoved)
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
		it.Event = new(FastUpdatesConfigurationFeedRemoved)
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
func (it *FastUpdatesConfigurationFeedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationFeedRemoved represents a FeedRemoved event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedRemoved struct {
	FeedId [21]byte
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedRemoved is a free log retrieval operation binding the contract event 0xbb4bc8e9bdadd13a82544df890de25d2c6403cd23a7655410eb2ad4f542425ab.
//
// Solidity: event FeedRemoved(bytes21 indexed feedId, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterFeedRemoved(opts *bind.FilterOpts, feedId [][21]byte) (*FastUpdatesConfigurationFeedRemovedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "FeedRemoved", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationFeedRemovedIterator{contract: _FastUpdatesConfiguration.contract, event: "FeedRemoved", logs: logs, sub: sub}, nil
}

// WatchFeedRemoved is a free log subscription operation binding the contract event 0xbb4bc8e9bdadd13a82544df890de25d2c6403cd23a7655410eb2ad4f542425ab.
//
// Solidity: event FeedRemoved(bytes21 indexed feedId, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchFeedRemoved(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationFeedRemoved, feedId [][21]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "FeedRemoved", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationFeedRemoved)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
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

// ParseFeedRemoved is a log parse operation binding the contract event 0xbb4bc8e9bdadd13a82544df890de25d2c6403cd23a7655410eb2ad4f542425ab.
//
// Solidity: event FeedRemoved(bytes21 indexed feedId, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseFeedRemoved(log types.Log) (*FastUpdatesConfigurationFeedRemoved, error) {
	event := new(FastUpdatesConfigurationFeedRemoved)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationFeedUpdatedIterator is returned from FilterFeedUpdated and is used to iterate over the raw logs and unpacked data for FeedUpdated events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedUpdatedIterator struct {
	Event *FastUpdatesConfigurationFeedUpdated // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationFeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationFeedUpdated)
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
		it.Event = new(FastUpdatesConfigurationFeedUpdated)
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
func (it *FastUpdatesConfigurationFeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationFeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationFeedUpdated represents a FeedUpdated event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationFeedUpdated struct {
	FeedId          [21]byte
	RewardBandValue uint32
	InflationShare  *big.Int
	Index           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeedUpdated is a free log retrieval operation binding the contract event 0x317c7e03c79b6fcd22d6f17813b4e8e8a4a14104fdfb79431c6c73b550c7ca9d.
//
// Solidity: event FeedUpdated(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterFeedUpdated(opts *bind.FilterOpts, feedId [][21]byte) (*FastUpdatesConfigurationFeedUpdatedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "FeedUpdated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationFeedUpdatedIterator{contract: _FastUpdatesConfiguration.contract, event: "FeedUpdated", logs: logs, sub: sub}, nil
}

// WatchFeedUpdated is a free log subscription operation binding the contract event 0x317c7e03c79b6fcd22d6f17813b4e8e8a4a14104fdfb79431c6c73b550c7ca9d.
//
// Solidity: event FeedUpdated(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchFeedUpdated(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationFeedUpdated, feedId [][21]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "FeedUpdated", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationFeedUpdated)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedUpdated", log); err != nil {
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

// ParseFeedUpdated is a log parse operation binding the contract event 0x317c7e03c79b6fcd22d6f17813b4e8e8a4a14104fdfb79431c6c73b550c7ca9d.
//
// Solidity: event FeedUpdated(bytes21 indexed feedId, uint32 rewardBandValue, uint24 inflationShare, uint256 index)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseFeedUpdated(log types.Log) (*FastUpdatesConfigurationFeedUpdated, error) {
	event := new(FastUpdatesConfigurationFeedUpdated)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "FeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernanceCallTimelockedIterator struct {
	Event *FastUpdatesConfigurationGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationGovernanceCallTimelocked)
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
		it.Event = new(FastUpdatesConfigurationGovernanceCallTimelocked)
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
func (it *FastUpdatesConfigurationGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FastUpdatesConfigurationGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationGovernanceCallTimelockedIterator{contract: _FastUpdatesConfiguration.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationGovernanceCallTimelocked)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FastUpdatesConfigurationGovernanceCallTimelocked, error) {
	event := new(FastUpdatesConfigurationGovernanceCallTimelocked)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernanceInitialisedIterator struct {
	Event *FastUpdatesConfigurationGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationGovernanceInitialised)
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
		it.Event = new(FastUpdatesConfigurationGovernanceInitialised)
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
func (it *FastUpdatesConfigurationGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationGovernanceInitialised represents a GovernanceInitialised event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FastUpdatesConfigurationGovernanceInitialisedIterator, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationGovernanceInitialisedIterator{contract: _FastUpdatesConfiguration.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationGovernanceInitialised)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseGovernanceInitialised(log types.Log) (*FastUpdatesConfigurationGovernanceInitialised, error) {
	event := new(FastUpdatesConfigurationGovernanceInitialised)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernedProductionModeEnteredIterator struct {
	Event *FastUpdatesConfigurationGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationGovernedProductionModeEntered)
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
		it.Event = new(FastUpdatesConfigurationGovernedProductionModeEntered)
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
func (it *FastUpdatesConfigurationGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FastUpdatesConfigurationGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationGovernedProductionModeEnteredIterator{contract: _FastUpdatesConfiguration.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationGovernedProductionModeEntered)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FastUpdatesConfigurationGovernedProductionModeEntered, error) {
	event := new(FastUpdatesConfigurationGovernedProductionModeEntered)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator struct {
	Event *FastUpdatesConfigurationTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationTimelockedGovernanceCallCanceled)
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
		it.Event = new(FastUpdatesConfigurationTimelockedGovernanceCallCanceled)
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
func (it *FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationTimelockedGovernanceCallCanceledIterator{contract: _FastUpdatesConfiguration.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationTimelockedGovernanceCallCanceled)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FastUpdatesConfigurationTimelockedGovernanceCallCanceled, error) {
	event := new(FastUpdatesConfigurationTimelockedGovernanceCallCanceled)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator struct {
	Event *FastUpdatesConfigurationTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastUpdatesConfigurationTimelockedGovernanceCallExecuted)
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
		it.Event = new(FastUpdatesConfigurationTimelockedGovernanceCallExecuted)
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
func (it *FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastUpdatesConfigurationTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FastUpdatesConfiguration contract.
type FastUpdatesConfigurationTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FastUpdatesConfigurationTimelockedGovernanceCallExecutedIterator{contract: _FastUpdatesConfiguration.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FastUpdatesConfigurationTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FastUpdatesConfiguration.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastUpdatesConfigurationTimelockedGovernanceCallExecuted)
				if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FastUpdatesConfiguration *FastUpdatesConfigurationFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FastUpdatesConfigurationTimelockedGovernanceCallExecuted, error) {
	event := new(FastUpdatesConfigurationTimelockedGovernanceCallExecuted)
	if err := _FastUpdatesConfiguration.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
