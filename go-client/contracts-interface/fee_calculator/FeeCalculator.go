// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee_calculator

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

// FeeCalculatorMetaData contains all meta data concerning the FeeCalculator contract.
var FeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"}],\"name\":\"CategoryFeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"CategoryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"}],\"name\":\"FeedFeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"FeedFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"}],\"name\":\"calculateFeeByIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"calculateFeeByIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdatesConfiguration\",\"outputs\":[{\"internalType\":\"contractIFastUpdatesConfiguration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_category\",\"type\":\"uint8\"}],\"name\":\"getCategoryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21\",\"name\":\"_feedId\",\"type\":\"bytes21\"}],\"name\":\"getFeedFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"_categories\",\"type\":\"uint8[]\"}],\"name\":\"removeCategoriesFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"}],\"name\":\"removeFeedsFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"_categories\",\"type\":\"uint8[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setCategoriesFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setDefaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setFeedsFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620022eb380380620022eb833981016040819052620000349162000297565b81848462000043828262000085565b506200006f9050817f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e771955565b506200007b81620001fe565b50505050620002f1565b600054600160a01b900460ff1615620000e55760405162461bcd60e51b815260206004820152601460248201527f696e697469616c6973656420213d2066616c736500000000000000000000000060448201526064015b60405180910390fd5b6001600160a01b0382166200013d5760405162461bcd60e51b815260206004820152601860248201527f676f7665726e616e63652073657474696e6773207a65726f00000000000000006044820152606401620000dc565b6001600160a01b038116620001885760405162461bcd60e51b815260206004820152601060248201526f5f676f7665726e616e6365207a65726f60801b6044820152606401620000dc565b600080546001600160a01b038481166001600160a81b031990921691909117600160a01b17909155600180549183166001600160a01b0319909216821790556040519081527f9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db9060200160405180910390a15050565b60008111620002435760405162461bcd60e51b815260206004820152601060248201526f64656661756c7420666565207a65726f60801b6044820152606401620000dc565b60068190556040518181527fe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac381279060200160405180910390a150565b6001600160a01b03811681146200029457600080fd5b50565b60008060008060808587031215620002ae57600080fd5b8451620002bb816200027e565b6020860151909450620002ce816200027e565b6040860151909350620002e1816200027e565b6060959095015193969295505050565b611fea80620003016000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c806395ec4acd116100b8578063debfda301161007c578063debfda30146102ad578063e17f212e146102d0578063e2f54db0146102e4578063e57e75d6146102f7578063ef88bf131461030a578063f5a983831461031d57600080fd5b806395ec4acd1461024e5780639f05059614610261578063b00c0b7614610274578063c10f489a14610287578063c93a6c841461029a57600080fd5b80635fee32e01161010a5780635fee32e0146101cc5780635ff27079146101df57806362354e03146101f457806367fc40291461020757806374e6310e1461021a578063755fcecd1461023b57600080fd5b806331af71a0146101475780634173680e1461016d5780635267a15d146101805780635a6c72d0146101bb5780635aa6e675146101c4575b600080fd5b61015a610155366004611879565b610325565b6040519081526020015b60405180910390f35b61015a61017b3660046118b6565b610408565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e7719545b6040516001600160a01b039091168152602001610164565b61015a60065481565b6101a36104ae565b61015a6101da3660046118e9565b61054a565b6101f26101ed366004611904565b6105c4565b005b6000546101a3906001600160a01b031681565b6101f2610215366004611904565b61084a565b61022d610228366004611904565b61092b565b60405161016492919061197e565b6101f26102493660046119f2565b6109d0565b6101f261025c366004611ab8565b610b50565b6101f261026f366004611aed565b610c2a565b6101f2610282366004611b9d565b610d86565b6003546101a3906001600160a01b031681565b6101f26102a8366004611c4c565b610e61565b6102c06102bb366004611c65565b610e99565b6040519015158152602001610164565b6000546102c090600160a81b900460ff1681565b61015a6102f2366004611c82565b610f22565b6101f2610305366004611879565b61106f565b6101f2610318366004611cb7565b611150565b6101f26112b6565b6000805b825181101561040257600083828151811061034657610346611cf0565b6020908102919091018101516001600160581b0319811660009081526005909252604090912054909150156103ac576001600160581b0319811660009081526005602052604090205461039b90600190611d1c565b6103a59084611d2f565b92506103ef565b600081811a815260046020526040902054156103df57600081811a81526004602052604090205461039b90600190611d1c565b6006546103ec9084611d2f565b92505b50806103fa81611d42565b915050610329565b50919050565b6001600160581b031981166000908152600560205260408120549081900361049d5760405162461bcd60e51b815260206004820152603b60248201527f6665656420666565206e6f74207365743b2063617465676f727920666565642060448201527f6f722064656661756c74206665652077696c6c2062652075736564000000000060648201526084015b60405180910390fd5b806104a781611d5b565b9392505050565b60008054600160a81b900460ff166104d057506001546001600160a01b031690565b60008054906101000a90046001600160a01b03166001600160a01b031663732524946040518163ffffffff1660e01b8152600401602060405180830381865afa158015610521573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105459190611d72565b905090565b60ff81166000908152600460205260408120549081900361049d5760405162461bcd60e51b815260206004820152602e60248201527f63617465676f727920666565206e6f74207365743b2064656661756c7420666560448201526d19481dda5b1b081899481d5cd95960921b6064820152608401610494565b6105cd33610e99565b6106095760405162461bcd60e51b815260206004820152600d60248201526c37b7363c9032bc32b1baba37b960991b6044820152606401610494565b6001600160e01b03198116600090815260026020526040812080549091036106735760405162461bcd60e51b815260206004820152601a60248201527f74696d656c6f636b3a20696e76616c69642073656c6563746f720000000000006044820152606401610494565b80544210156106c45760405162461bcd60e51b815260206004820152601960248201527f74696d656c6f636b3a206e6f7420616c6c6f77656420796574000000000000006044820152606401610494565b60008160010180546106d590611d8f565b80601f016020809104026020016040519081016040528092919081815260200182805461070190611d8f565b801561074e5780601f106107235761010080835404028352916020019161074e565b820191906000526020600020905b81548152906001019060200180831161073157829003601f168201915b505050506001600160e01b03198516600090815260026020526040812081815592935090506107806001830182611736565b50506000805460ff60b01b1916600160b01b17815560405130906107a5908490611dc3565b6000604051808303816000865af19150503d80600081146107e2576040519150601f19603f3d011682016040523d82523d6000602084013e6107e7565b606091505b50506000805460ff60b01b19169055604080516001600160e01b0319871681524260208201529192507fa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438910160405180910390a16108448161137c565b50505050565b610852611399565b6001600160e01b0319811660009081526002602052604081205490036108ba5760405162461bcd60e51b815260206004820152601a60248201527f74696d656c6f636b3a20696e76616c69642073656c6563746f720000000000006044820152606401610494565b604080516001600160e01b0319831681524260208201527f7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8910160405180910390a16001600160e01b031981166000908152600260205260408120818155906109266001830182611736565b505050565b6002602052600090815260409020805460018201805491929161094d90611d8f565b80601f016020809104026020016040519081016040528092919081815260200182805461097990611d8f565b80156109c65780601f1061099b576101008083540402835291602001916109c6565b820191906000526020600020905b8154815290600101906020018083116109a957829003601f168201915b5050505050905082565b600054600160b01b900460ff16806109f25750600054600160a81b900460ff16155b15610b41576109ff6113f5565b8051825114610a435760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d0e640dad2e6dac2e8c6d60831b6044820152606401610494565b60005b825181101561092657818181518110610a6157610a61611cf0565b60200260200101516001610a759190611d2f565b60056000858481518110610a8b57610a8b611cf0565b60200260200101516001600160581b0319166001600160581b031916815260200190815260200160002081905550828181518110610acb57610acb611cf0565b60200260200101516001600160581b0319167f732e2d3065b43a1e846279a2b9f63938ae3ebd6422b24c040cf2ee6a667a1b1b838381518110610b1057610b10611cf0565b6020026020010151604051610b2791815260200190565b60405180910390a280610b3981611d42565b915050610a46565b610b4c60003661142d565b5050565b600054600160b01b900460ff1680610b725750600054600160a81b900460ff16155b15610c1c57610b7f6113f5565b60005b8151811015610b4c5760046000838381518110610ba157610ba1611cf0565b602002602001015160ff1660ff16815260200190815260200160002060009055818181518110610bd357610bd3611cf0565b602002602001015160ff167f27183ed3733512b0274d16c44d6410021e474e6a8082b1ecab4b5dc8dc5763aa60405160405180910390a280610c1481611d42565b915050610b82565b610c2760003661142d565b50565b600054600160b01b900460ff1680610c4c5750600054600160a81b900460ff16155b15610b4157610c596113f5565b8051825114610c9d5760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d0e640dad2e6dac2e8c6d60831b6044820152606401610494565b60005b825181101561092657818181518110610cbb57610cbb611cf0565b60200260200101516001610ccf9190611d2f565b60046000858481518110610ce557610ce5611cf0565b602002602001015160ff1660ff16815260200190815260200160002081905550828181518110610d1757610d17611cf0565b602002602001015160ff167f7970bb47884d377ab068600af9d8ada9bf6294843fcf6ffeec86892fb44123d3838381518110610d5557610d55611cf0565b6020026020010151604051610d6c91815260200190565b60405180910390a280610d7e81611d42565b915050610ca0565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e7719546001600160a01b0316336001600160a01b031614610dff5760405162461bcd60e51b815260206004820152601460248201527337b7363c9030b2323932b9b9903ab83230ba32b960611b6044820152606401610494565b610e57610e3383836040518060400160405280600e81526020016d20b2323932b9b9aab83230ba32b960911b815250611579565b7f714f205b2abd25bef1d06a1af944e38c113fe6160375c4e1d6d5cf28848e771955565b610b4c8282611654565b600054600160b01b900460ff1680610e835750600054600160a81b900460ff16155b15610c1c57610e906113f5565b610c27816116b8565b60008054600160a01b900460ff168015610f1c5750600054604051630debfda360e41b81526001600160a01b0384811660048301529091169063debfda3090602401602060405180830381865afa158015610ef8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f1c9190611ddf565b92915050565b6000805b82518110156104025760035483516000916001600160a01b031690639310283690869085908110610f5957610f59611cf0565b60200260200101516040518263ffffffff1660e01b8152600401610f7f91815260200190565b602060405180830381865afa158015610f9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc09190611e01565b6001600160581b0319811660009081526005602052604090205490915015611019576001600160581b0319811660009081526005602052604090205461100890600190611d1c565b6110129084611d2f565b925061105c565b600081811a8152600460205260409020541561104c57600081811a81526004602052604090205461100890600190611d1c565b6006546110599084611d2f565b92505b508061106781611d42565b915050610f26565b600054600160b01b900460ff16806110915750600054600160a81b900460ff16155b15610c1c5761109e6113f5565b60005b8151811015610b4c57600560008383815181106110c0576110c0611cf0565b60200260200101516001600160581b0319166001600160581b03191681526020019081526020016000206000905581818151811061110057611100611cf0565b60200260200101516001600160581b0319167ffec81636c21a01d5af8282aefe4cd398aa5a3f2f8e85896c3373d83498317be660405160405180910390a28061114881611d42565b9150506110a1565b600054600160a01b900460ff16156111a15760405162461bcd60e51b8152602060048201526014602482015273696e697469616c6973656420213d2066616c736560601b6044820152606401610494565b6001600160a01b0382166111f75760405162461bcd60e51b815260206004820152601860248201527f676f7665726e616e63652073657474696e6773207a65726f00000000000000006044820152606401610494565b6001600160a01b0381166112405760405162461bcd60e51b815260206004820152601060248201526f5f676f7665726e616e6365207a65726f60801b6044820152606401610494565b600080546001600160a01b038481166001600160a81b031990921691909117600160a01b17909155600180549183166001600160a01b0319909216821790556040519081527f9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db9060200160405180910390a15050565b6112be611399565b600054600160a81b900460ff16156113185760405162461bcd60e51b815260206004820152601a60248201527f616c726561647920696e2070726f64756374696f6e206d6f64650000000000006044820152606401610494565b600180546001600160a01b031916905560008054600160a81b60ff60a81b198216179091556040516001600160a01b0390911681527f83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c9060200160405180910390a1565b3d604051818101604052816000823e8215611395578181f35b8181fd5b6113a16104ae565b6001600160a01b0316336001600160a01b0316146113f35760405162461bcd60e51b815260206004820152600f60248201526e6f6e6c7920676f7665726e616e636560881b6044820152606401610494565b565b600054600160b01b900460ff16156114255733301461141657611416611e1e565b6000805460ff60b01b19169055565b6113f3611399565b611435611399565b6000805460408051636221a54b60e01b81529051853593926001600160a01b031691636221a54b9160048083019260209291908290030181865afa158015611481573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a59190611e34565b905060006114b38242611d2f565b9050604051806040016040528082815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509390945250506001600160e01b031986168152600260209081526040909120835181559083015190915060018201906115319082611e9b565b509050507fed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b68382878760405161156a9493929190611f5b565b60405180910390a15050505050565b6000808260405160200161158d9190611fa1565b6040516020818303038152906040528051906020012090506000805b8651811015611605578681815181106115c4576115c4611cf0565b602002602001015183036115f3578581815181106115e4576115e4611cf0565b60200260200101519150611605565b806115fd81611d42565b9150506115a9565b506001600160a01b03811661164b5760405162461bcd60e51b815260206004820152600c60248201526b61646472657373207a65726f60a01b6044820152606401610494565b95945050505050565b61169482826040518060400160405280601881526020017f4661737455706461746573436f6e66696775726174696f6e0000000000000000815250611579565b600380546001600160a01b0319166001600160a01b03929092169190911790555050565b600081116116fb5760405162461bcd60e51b815260206004820152601060248201526f64656661756c7420666565207a65726f60801b6044820152606401610494565b60068190556040518181527fe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac381279060200160405180910390a150565b50805461174290611d8f565b6000825580601f10611752575050565b601f016020900490600052602060002090810190610c2791905b80821115611780576000815560010161176c565b5090565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156117c3576117c3611784565b604052919050565b600067ffffffffffffffff8211156117e5576117e5611784565b5060051b60200190565b6001600160581b031981168114610c2757600080fd5b600082601f83011261181657600080fd5b8135602061182b611826836117cb565b61179a565b82815260059290921b8401810191818101908684111561184a57600080fd5b8286015b8481101561186e578035611861816117ef565b835291830191830161184e565b509695505050505050565b60006020828403121561188b57600080fd5b813567ffffffffffffffff8111156118a257600080fd5b6118ae84828501611805565b949350505050565b6000602082840312156118c857600080fd5b81356104a7816117ef565b803560ff811681146118e457600080fd5b919050565b6000602082840312156118fb57600080fd5b6104a7826118d3565b60006020828403121561191657600080fd5b81356001600160e01b0319811681146104a757600080fd5b60005b83811015611949578181015183820152602001611931565b50506000910152565b6000815180845261196a81602086016020860161192e565b601f01601f19169290920160200192915050565b8281526040602082015260006118ae6040830184611952565b600082601f8301126119a857600080fd5b813560206119b8611826836117cb565b82815260059290921b840181019181810190868411156119d757600080fd5b8286015b8481101561186e57803583529183019183016119db565b60008060408385031215611a0557600080fd5b823567ffffffffffffffff80821115611a1d57600080fd5b611a2986838701611805565b93506020850135915080821115611a3f57600080fd5b50611a4c85828601611997565b9150509250929050565b600082601f830112611a6757600080fd5b81356020611a77611826836117cb565b82815260059290921b84018101918181019086841115611a9657600080fd5b8286015b8481101561186e57611aab816118d3565b8352918301918301611a9a565b600060208284031215611aca57600080fd5b813567ffffffffffffffff811115611ae157600080fd5b6118ae84828501611a56565b60008060408385031215611b0057600080fd5b823567ffffffffffffffff80821115611b1857600080fd5b611a2986838701611a56565b6001600160a01b0381168114610c2757600080fd5b600082601f830112611b4a57600080fd5b81356020611b5a611826836117cb565b82815260059290921b84018101918181019086841115611b7957600080fd5b8286015b8481101561186e578035611b9081611b24565b8352918301918301611b7d565b60008060408385031215611bb057600080fd5b823567ffffffffffffffff80821115611bc857600080fd5b818501915085601f830112611bdc57600080fd5b81356020611bec611826836117cb565b82815260059290921b84018101918181019089841115611c0b57600080fd5b948201945b83861015611c2957853582529482019490820190611c10565b96505086013592505080821115611c3f57600080fd5b50611a4c85828601611b39565b600060208284031215611c5e57600080fd5b5035919050565b600060208284031215611c7757600080fd5b81356104a781611b24565b600060208284031215611c9457600080fd5b813567ffffffffffffffff811115611cab57600080fd5b6118ae84828501611997565b60008060408385031215611cca57600080fd5b8235611cd581611b24565b91506020830135611ce581611b24565b809150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b81810381811115610f1c57610f1c611d06565b80820180821115610f1c57610f1c611d06565b600060018201611d5457611d54611d06565b5060010190565b600081611d6a57611d6a611d06565b506000190190565b600060208284031215611d8457600080fd5b81516104a781611b24565b600181811c90821680611da357607f821691505b60208210810361040257634e487b7160e01b600052602260045260246000fd5b60008251611dd581846020870161192e565b9190910192915050565b600060208284031215611df157600080fd5b815180151581146104a757600080fd5b600060208284031215611e1357600080fd5b81516104a7816117ef565b634e487b7160e01b600052600160045260246000fd5b600060208284031215611e4657600080fd5b5051919050565b601f82111561092657600081815260208120601f850160051c81016020861015611e745750805b601f850160051c820191505b81811015611e9357828155600101611e80565b505050505050565b815167ffffffffffffffff811115611eb557611eb5611784565b611ec981611ec38454611d8f565b84611e4d565b602080601f831160018114611efe5760008415611ee65750858301515b600019600386901b1c1916600185901b178555611e93565b600085815260208120601f198616915b82811015611f2d57888601518255948401946001909101908401611f0e565b5085821015611f4b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b63ffffffff60e01b8516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f909201601f191601019392505050565b6020815260006104a7602083018461195256fea2646970667358221220696b001fece26ae13408150018a1901fc5f78cfa3d2817516640eaa858170a2464736f6c63430008140033",
}

// FeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeCalculatorMetaData.ABI instead.
var FeeCalculatorABI = FeeCalculatorMetaData.ABI

// FeeCalculatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeCalculatorMetaData.Bin instead.
var FeeCalculatorBin = FeeCalculatorMetaData.Bin

// DeployFeeCalculator deploys a new Ethereum contract, binding an instance of FeeCalculator to it.
func DeployFeeCalculator(auth *bind.TransactOpts, backend bind.ContractBackend, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _defaultFee *big.Int) (common.Address, *types.Transaction, *FeeCalculator, error) {
	parsed, err := FeeCalculatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeCalculatorBin), backend, _governanceSettings, _initialGovernance, _addressUpdater, _defaultFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeCalculator{FeeCalculatorCaller: FeeCalculatorCaller{contract: contract}, FeeCalculatorTransactor: FeeCalculatorTransactor{contract: contract}, FeeCalculatorFilterer: FeeCalculatorFilterer{contract: contract}}, nil
}

// FeeCalculator is an auto generated Go binding around an Ethereum contract.
type FeeCalculator struct {
	FeeCalculatorCaller     // Read-only binding to the contract
	FeeCalculatorTransactor // Write-only binding to the contract
	FeeCalculatorFilterer   // Log filterer for contract events
}

// FeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeCalculatorSession struct {
	Contract     *FeeCalculator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeCalculatorCallerSession struct {
	Contract *FeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeCalculatorTransactorSession struct {
	Contract     *FeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeCalculatorRaw struct {
	Contract *FeeCalculator // Generic contract binding to access the raw methods on
}

// FeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeCalculatorCallerRaw struct {
	Contract *FeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// FeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeCalculatorTransactorRaw struct {
	Contract *FeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeCalculator creates a new instance of FeeCalculator, bound to a specific deployed contract.
func NewFeeCalculator(address common.Address, backend bind.ContractBackend) (*FeeCalculator, error) {
	contract, err := bindFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeCalculator{FeeCalculatorCaller: FeeCalculatorCaller{contract: contract}, FeeCalculatorTransactor: FeeCalculatorTransactor{contract: contract}, FeeCalculatorFilterer: FeeCalculatorFilterer{contract: contract}}, nil
}

// NewFeeCalculatorCaller creates a new read-only instance of FeeCalculator, bound to a specific deployed contract.
func NewFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*FeeCalculatorCaller, error) {
	contract, err := bindFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorCaller{contract: contract}, nil
}

// NewFeeCalculatorTransactor creates a new write-only instance of FeeCalculator, bound to a specific deployed contract.
func NewFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeCalculatorTransactor, error) {
	contract, err := bindFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorTransactor{contract: contract}, nil
}

// NewFeeCalculatorFilterer creates a new log filterer instance of FeeCalculator, bound to a specific deployed contract.
func NewFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeCalculatorFilterer, error) {
	contract, err := bindFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorFilterer{contract: contract}, nil
}

// bindFeeCalculator binds a generic wrapper to an already deployed contract.
func bindFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCalculator *FeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCalculator.Contract.FeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCalculator *FeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCalculator.Contract.FeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCalculator *FeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCalculator.Contract.FeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCalculator *FeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCalculator *FeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCalculator *FeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateFeeByIds is a free data retrieval call binding the contract method 0x31af71a0.
//
// Solidity: function calculateFeeByIds(bytes21[] _feedIds) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCaller) CalculateFeeByIds(opts *bind.CallOpts, _feedIds [][21]byte) (*big.Int, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "calculateFeeByIds", _feedIds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeByIds is a free data retrieval call binding the contract method 0x31af71a0.
//
// Solidity: function calculateFeeByIds(bytes21[] _feedIds) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorSession) CalculateFeeByIds(_feedIds [][21]byte) (*big.Int, error) {
	return _FeeCalculator.Contract.CalculateFeeByIds(&_FeeCalculator.CallOpts, _feedIds)
}

// CalculateFeeByIds is a free data retrieval call binding the contract method 0x31af71a0.
//
// Solidity: function calculateFeeByIds(bytes21[] _feedIds) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCallerSession) CalculateFeeByIds(_feedIds [][21]byte) (*big.Int, error) {
	return _FeeCalculator.Contract.CalculateFeeByIds(&_FeeCalculator.CallOpts, _feedIds)
}

// CalculateFeeByIndices is a free data retrieval call binding the contract method 0xe2f54db0.
//
// Solidity: function calculateFeeByIndices(uint256[] _indices) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCaller) CalculateFeeByIndices(opts *bind.CallOpts, _indices []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "calculateFeeByIndices", _indices)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeByIndices is a free data retrieval call binding the contract method 0xe2f54db0.
//
// Solidity: function calculateFeeByIndices(uint256[] _indices) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorSession) CalculateFeeByIndices(_indices []*big.Int) (*big.Int, error) {
	return _FeeCalculator.Contract.CalculateFeeByIndices(&_FeeCalculator.CallOpts, _indices)
}

// CalculateFeeByIndices is a free data retrieval call binding the contract method 0xe2f54db0.
//
// Solidity: function calculateFeeByIndices(uint256[] _indices) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCallerSession) CalculateFeeByIndices(_indices []*big.Int) (*big.Int, error) {
	return _FeeCalculator.Contract.CalculateFeeByIndices(&_FeeCalculator.CallOpts, _indices)
}

// DefaultFee is a free data retrieval call binding the contract method 0x5a6c72d0.
//
// Solidity: function defaultFee() view returns(uint256)
func (_FeeCalculator *FeeCalculatorCaller) DefaultFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "defaultFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultFee is a free data retrieval call binding the contract method 0x5a6c72d0.
//
// Solidity: function defaultFee() view returns(uint256)
func (_FeeCalculator *FeeCalculatorSession) DefaultFee() (*big.Int, error) {
	return _FeeCalculator.Contract.DefaultFee(&_FeeCalculator.CallOpts)
}

// DefaultFee is a free data retrieval call binding the contract method 0x5a6c72d0.
//
// Solidity: function defaultFee() view returns(uint256)
func (_FeeCalculator *FeeCalculatorCallerSession) DefaultFee() (*big.Int, error) {
	return _FeeCalculator.Contract.DefaultFee(&_FeeCalculator.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FeeCalculator *FeeCalculatorCaller) FastUpdatesConfiguration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "fastUpdatesConfiguration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FeeCalculator *FeeCalculatorSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FeeCalculator.Contract.FastUpdatesConfiguration(&_FeeCalculator.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FeeCalculator *FeeCalculatorCallerSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FeeCalculator.Contract.FastUpdatesConfiguration(&_FeeCalculator.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FeeCalculator *FeeCalculatorCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FeeCalculator *FeeCalculatorSession) GetAddressUpdater() (common.Address, error) {
	return _FeeCalculator.Contract.GetAddressUpdater(&_FeeCalculator.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FeeCalculator *FeeCalculatorCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FeeCalculator.Contract.GetAddressUpdater(&_FeeCalculator.CallOpts)
}

// GetCategoryFee is a free data retrieval call binding the contract method 0x5fee32e0.
//
// Solidity: function getCategoryFee(uint8 _category) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCaller) GetCategoryFee(opts *bind.CallOpts, _category uint8) (*big.Int, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "getCategoryFee", _category)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCategoryFee is a free data retrieval call binding the contract method 0x5fee32e0.
//
// Solidity: function getCategoryFee(uint8 _category) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorSession) GetCategoryFee(_category uint8) (*big.Int, error) {
	return _FeeCalculator.Contract.GetCategoryFee(&_FeeCalculator.CallOpts, _category)
}

// GetCategoryFee is a free data retrieval call binding the contract method 0x5fee32e0.
//
// Solidity: function getCategoryFee(uint8 _category) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCallerSession) GetCategoryFee(_category uint8) (*big.Int, error) {
	return _FeeCalculator.Contract.GetCategoryFee(&_FeeCalculator.CallOpts, _category)
}

// GetFeedFee is a free data retrieval call binding the contract method 0x4173680e.
//
// Solidity: function getFeedFee(bytes21 _feedId) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCaller) GetFeedFee(opts *bind.CallOpts, _feedId [21]byte) (*big.Int, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "getFeedFee", _feedId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeedFee is a free data retrieval call binding the contract method 0x4173680e.
//
// Solidity: function getFeedFee(bytes21 _feedId) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorSession) GetFeedFee(_feedId [21]byte) (*big.Int, error) {
	return _FeeCalculator.Contract.GetFeedFee(&_FeeCalculator.CallOpts, _feedId)
}

// GetFeedFee is a free data retrieval call binding the contract method 0x4173680e.
//
// Solidity: function getFeedFee(bytes21 _feedId) view returns(uint256 _fee)
func (_FeeCalculator *FeeCalculatorCallerSession) GetFeedFee(_feedId [21]byte) (*big.Int, error) {
	return _FeeCalculator.Contract.GetFeedFee(&_FeeCalculator.CallOpts, _feedId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FeeCalculator *FeeCalculatorCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FeeCalculator *FeeCalculatorSession) Governance() (common.Address, error) {
	return _FeeCalculator.Contract.Governance(&_FeeCalculator.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FeeCalculator *FeeCalculatorCallerSession) Governance() (common.Address, error) {
	return _FeeCalculator.Contract.Governance(&_FeeCalculator.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FeeCalculator *FeeCalculatorCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FeeCalculator *FeeCalculatorSession) GovernanceSettings() (common.Address, error) {
	return _FeeCalculator.Contract.GovernanceSettings(&_FeeCalculator.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FeeCalculator *FeeCalculatorCallerSession) GovernanceSettings() (common.Address, error) {
	return _FeeCalculator.Contract.GovernanceSettings(&_FeeCalculator.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FeeCalculator *FeeCalculatorCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FeeCalculator *FeeCalculatorSession) IsExecutor(_address common.Address) (bool, error) {
	return _FeeCalculator.Contract.IsExecutor(&_FeeCalculator.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FeeCalculator *FeeCalculatorCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FeeCalculator.Contract.IsExecutor(&_FeeCalculator.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FeeCalculator *FeeCalculatorCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FeeCalculator *FeeCalculatorSession) ProductionMode() (bool, error) {
	return _FeeCalculator.Contract.ProductionMode(&_FeeCalculator.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FeeCalculator *FeeCalculatorCallerSession) ProductionMode() (bool, error) {
	return _FeeCalculator.Contract.ProductionMode(&_FeeCalculator.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FeeCalculator *FeeCalculatorCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FeeCalculator.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_FeeCalculator *FeeCalculatorSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FeeCalculator.Contract.TimelockedCalls(&_FeeCalculator.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FeeCalculator *FeeCalculatorCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FeeCalculator.Contract.TimelockedCalls(&_FeeCalculator.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.CancelGovernanceCall(&_FeeCalculator.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.CancelGovernanceCall(&_FeeCalculator.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.ExecuteGovernanceCall(&_FeeCalculator.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.ExecuteGovernanceCall(&_FeeCalculator.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FeeCalculator *FeeCalculatorTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FeeCalculator *FeeCalculatorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FeeCalculator.Contract.Initialise(&_FeeCalculator.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FeeCalculator.Contract.Initialise(&_FeeCalculator.TransactOpts, _governanceSettings, _initialGovernance)
}

// RemoveCategoriesFees is a paid mutator transaction binding the contract method 0x95ec4acd.
//
// Solidity: function removeCategoriesFees(uint8[] _categories) returns()
func (_FeeCalculator *FeeCalculatorTransactor) RemoveCategoriesFees(opts *bind.TransactOpts, _categories []uint8) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "removeCategoriesFees", _categories)
}

// RemoveCategoriesFees is a paid mutator transaction binding the contract method 0x95ec4acd.
//
// Solidity: function removeCategoriesFees(uint8[] _categories) returns()
func (_FeeCalculator *FeeCalculatorSession) RemoveCategoriesFees(_categories []uint8) (*types.Transaction, error) {
	return _FeeCalculator.Contract.RemoveCategoriesFees(&_FeeCalculator.TransactOpts, _categories)
}

// RemoveCategoriesFees is a paid mutator transaction binding the contract method 0x95ec4acd.
//
// Solidity: function removeCategoriesFees(uint8[] _categories) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) RemoveCategoriesFees(_categories []uint8) (*types.Transaction, error) {
	return _FeeCalculator.Contract.RemoveCategoriesFees(&_FeeCalculator.TransactOpts, _categories)
}

// RemoveFeedsFees is a paid mutator transaction binding the contract method 0xe57e75d6.
//
// Solidity: function removeFeedsFees(bytes21[] _feedIds) returns()
func (_FeeCalculator *FeeCalculatorTransactor) RemoveFeedsFees(opts *bind.TransactOpts, _feedIds [][21]byte) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "removeFeedsFees", _feedIds)
}

// RemoveFeedsFees is a paid mutator transaction binding the contract method 0xe57e75d6.
//
// Solidity: function removeFeedsFees(bytes21[] _feedIds) returns()
func (_FeeCalculator *FeeCalculatorSession) RemoveFeedsFees(_feedIds [][21]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.RemoveFeedsFees(&_FeeCalculator.TransactOpts, _feedIds)
}

// RemoveFeedsFees is a paid mutator transaction binding the contract method 0xe57e75d6.
//
// Solidity: function removeFeedsFees(bytes21[] _feedIds) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) RemoveFeedsFees(_feedIds [][21]byte) (*types.Transaction, error) {
	return _FeeCalculator.Contract.RemoveFeedsFees(&_FeeCalculator.TransactOpts, _feedIds)
}

// SetCategoriesFees is a paid mutator transaction binding the contract method 0x9f050596.
//
// Solidity: function setCategoriesFees(uint8[] _categories, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorTransactor) SetCategoriesFees(opts *bind.TransactOpts, _categories []uint8, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "setCategoriesFees", _categories, _fees)
}

// SetCategoriesFees is a paid mutator transaction binding the contract method 0x9f050596.
//
// Solidity: function setCategoriesFees(uint8[] _categories, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorSession) SetCategoriesFees(_categories []uint8, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetCategoriesFees(&_FeeCalculator.TransactOpts, _categories, _fees)
}

// SetCategoriesFees is a paid mutator transaction binding the contract method 0x9f050596.
//
// Solidity: function setCategoriesFees(uint8[] _categories, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) SetCategoriesFees(_categories []uint8, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetCategoriesFees(&_FeeCalculator.TransactOpts, _categories, _fees)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _fee) returns()
func (_FeeCalculator *FeeCalculatorTransactor) SetDefaultFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "setDefaultFee", _fee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _fee) returns()
func (_FeeCalculator *FeeCalculatorSession) SetDefaultFee(_fee *big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetDefaultFee(&_FeeCalculator.TransactOpts, _fee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _fee) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) SetDefaultFee(_fee *big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetDefaultFee(&_FeeCalculator.TransactOpts, _fee)
}

// SetFeedsFees is a paid mutator transaction binding the contract method 0x755fcecd.
//
// Solidity: function setFeedsFees(bytes21[] _feedIds, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorTransactor) SetFeedsFees(opts *bind.TransactOpts, _feedIds [][21]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "setFeedsFees", _feedIds, _fees)
}

// SetFeedsFees is a paid mutator transaction binding the contract method 0x755fcecd.
//
// Solidity: function setFeedsFees(bytes21[] _feedIds, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorSession) SetFeedsFees(_feedIds [][21]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetFeedsFees(&_FeeCalculator.TransactOpts, _feedIds, _fees)
}

// SetFeedsFees is a paid mutator transaction binding the contract method 0x755fcecd.
//
// Solidity: function setFeedsFees(bytes21[] _feedIds, uint256[] _fees) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) SetFeedsFees(_feedIds [][21]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FeeCalculator.Contract.SetFeedsFees(&_FeeCalculator.TransactOpts, _feedIds, _fees)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FeeCalculator *FeeCalculatorTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FeeCalculator *FeeCalculatorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FeeCalculator.Contract.SwitchToProductionMode(&_FeeCalculator.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FeeCalculator.Contract.SwitchToProductionMode(&_FeeCalculator.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FeeCalculator *FeeCalculatorTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FeeCalculator.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FeeCalculator *FeeCalculatorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FeeCalculator.Contract.UpdateContractAddresses(&_FeeCalculator.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FeeCalculator *FeeCalculatorTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FeeCalculator.Contract.UpdateContractAddresses(&_FeeCalculator.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FeeCalculatorCategoryFeeRemovedIterator is returned from FilterCategoryFeeRemoved and is used to iterate over the raw logs and unpacked data for CategoryFeeRemoved events raised by the FeeCalculator contract.
type FeeCalculatorCategoryFeeRemovedIterator struct {
	Event *FeeCalculatorCategoryFeeRemoved // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorCategoryFeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorCategoryFeeRemoved)
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
		it.Event = new(FeeCalculatorCategoryFeeRemoved)
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
func (it *FeeCalculatorCategoryFeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorCategoryFeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorCategoryFeeRemoved represents a CategoryFeeRemoved event raised by the FeeCalculator contract.
type FeeCalculatorCategoryFeeRemoved struct {
	Category uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCategoryFeeRemoved is a free log retrieval operation binding the contract event 0x27183ed3733512b0274d16c44d6410021e474e6a8082b1ecab4b5dc8dc5763aa.
//
// Solidity: event CategoryFeeRemoved(uint8 indexed category)
func (_FeeCalculator *FeeCalculatorFilterer) FilterCategoryFeeRemoved(opts *bind.FilterOpts, category []uint8) (*FeeCalculatorCategoryFeeRemovedIterator, error) {

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "CategoryFeeRemoved", categoryRule)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorCategoryFeeRemovedIterator{contract: _FeeCalculator.contract, event: "CategoryFeeRemoved", logs: logs, sub: sub}, nil
}

// WatchCategoryFeeRemoved is a free log subscription operation binding the contract event 0x27183ed3733512b0274d16c44d6410021e474e6a8082b1ecab4b5dc8dc5763aa.
//
// Solidity: event CategoryFeeRemoved(uint8 indexed category)
func (_FeeCalculator *FeeCalculatorFilterer) WatchCategoryFeeRemoved(opts *bind.WatchOpts, sink chan<- *FeeCalculatorCategoryFeeRemoved, category []uint8) (event.Subscription, error) {

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "CategoryFeeRemoved", categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorCategoryFeeRemoved)
				if err := _FeeCalculator.contract.UnpackLog(event, "CategoryFeeRemoved", log); err != nil {
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

// ParseCategoryFeeRemoved is a log parse operation binding the contract event 0x27183ed3733512b0274d16c44d6410021e474e6a8082b1ecab4b5dc8dc5763aa.
//
// Solidity: event CategoryFeeRemoved(uint8 indexed category)
func (_FeeCalculator *FeeCalculatorFilterer) ParseCategoryFeeRemoved(log types.Log) (*FeeCalculatorCategoryFeeRemoved, error) {
	event := new(FeeCalculatorCategoryFeeRemoved)
	if err := _FeeCalculator.contract.UnpackLog(event, "CategoryFeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorCategoryFeeSetIterator is returned from FilterCategoryFeeSet and is used to iterate over the raw logs and unpacked data for CategoryFeeSet events raised by the FeeCalculator contract.
type FeeCalculatorCategoryFeeSetIterator struct {
	Event *FeeCalculatorCategoryFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorCategoryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorCategoryFeeSet)
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
		it.Event = new(FeeCalculatorCategoryFeeSet)
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
func (it *FeeCalculatorCategoryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorCategoryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorCategoryFeeSet represents a CategoryFeeSet event raised by the FeeCalculator contract.
type FeeCalculatorCategoryFeeSet struct {
	Category uint8
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCategoryFeeSet is a free log retrieval operation binding the contract event 0x7970bb47884d377ab068600af9d8ada9bf6294843fcf6ffeec86892fb44123d3.
//
// Solidity: event CategoryFeeSet(uint8 indexed category, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) FilterCategoryFeeSet(opts *bind.FilterOpts, category []uint8) (*FeeCalculatorCategoryFeeSetIterator, error) {

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "CategoryFeeSet", categoryRule)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorCategoryFeeSetIterator{contract: _FeeCalculator.contract, event: "CategoryFeeSet", logs: logs, sub: sub}, nil
}

// WatchCategoryFeeSet is a free log subscription operation binding the contract event 0x7970bb47884d377ab068600af9d8ada9bf6294843fcf6ffeec86892fb44123d3.
//
// Solidity: event CategoryFeeSet(uint8 indexed category, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) WatchCategoryFeeSet(opts *bind.WatchOpts, sink chan<- *FeeCalculatorCategoryFeeSet, category []uint8) (event.Subscription, error) {

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "CategoryFeeSet", categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorCategoryFeeSet)
				if err := _FeeCalculator.contract.UnpackLog(event, "CategoryFeeSet", log); err != nil {
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

// ParseCategoryFeeSet is a log parse operation binding the contract event 0x7970bb47884d377ab068600af9d8ada9bf6294843fcf6ffeec86892fb44123d3.
//
// Solidity: event CategoryFeeSet(uint8 indexed category, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) ParseCategoryFeeSet(log types.Log) (*FeeCalculatorCategoryFeeSet, error) {
	event := new(FeeCalculatorCategoryFeeSet)
	if err := _FeeCalculator.contract.UnpackLog(event, "CategoryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorDefaultFeeSetIterator is returned from FilterDefaultFeeSet and is used to iterate over the raw logs and unpacked data for DefaultFeeSet events raised by the FeeCalculator contract.
type FeeCalculatorDefaultFeeSetIterator struct {
	Event *FeeCalculatorDefaultFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorDefaultFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorDefaultFeeSet)
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
		it.Event = new(FeeCalculatorDefaultFeeSet)
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
func (it *FeeCalculatorDefaultFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorDefaultFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorDefaultFeeSet represents a DefaultFeeSet event raised by the FeeCalculator contract.
type FeeCalculatorDefaultFeeSet struct {
	Fee *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeSet is a free log retrieval operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) FilterDefaultFeeSet(opts *bind.FilterOpts) (*FeeCalculatorDefaultFeeSetIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorDefaultFeeSetIterator{contract: _FeeCalculator.contract, event: "DefaultFeeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeSet is a free log subscription operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) WatchDefaultFeeSet(opts *bind.WatchOpts, sink chan<- *FeeCalculatorDefaultFeeSet) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorDefaultFeeSet)
				if err := _FeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
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

// ParseDefaultFeeSet is a log parse operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) ParseDefaultFeeSet(log types.Log) (*FeeCalculatorDefaultFeeSet, error) {
	event := new(FeeCalculatorDefaultFeeSet)
	if err := _FeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorFeedFeeRemovedIterator is returned from FilterFeedFeeRemoved and is used to iterate over the raw logs and unpacked data for FeedFeeRemoved events raised by the FeeCalculator contract.
type FeeCalculatorFeedFeeRemovedIterator struct {
	Event *FeeCalculatorFeedFeeRemoved // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorFeedFeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorFeedFeeRemoved)
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
		it.Event = new(FeeCalculatorFeedFeeRemoved)
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
func (it *FeeCalculatorFeedFeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorFeedFeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorFeedFeeRemoved represents a FeedFeeRemoved event raised by the FeeCalculator contract.
type FeeCalculatorFeedFeeRemoved struct {
	FeedId [21]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedFeeRemoved is a free log retrieval operation binding the contract event 0xfec81636c21a01d5af8282aefe4cd398aa5a3f2f8e85896c3373d83498317be6.
//
// Solidity: event FeedFeeRemoved(bytes21 indexed feedId)
func (_FeeCalculator *FeeCalculatorFilterer) FilterFeedFeeRemoved(opts *bind.FilterOpts, feedId [][21]byte) (*FeeCalculatorFeedFeeRemovedIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "FeedFeeRemoved", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorFeedFeeRemovedIterator{contract: _FeeCalculator.contract, event: "FeedFeeRemoved", logs: logs, sub: sub}, nil
}

// WatchFeedFeeRemoved is a free log subscription operation binding the contract event 0xfec81636c21a01d5af8282aefe4cd398aa5a3f2f8e85896c3373d83498317be6.
//
// Solidity: event FeedFeeRemoved(bytes21 indexed feedId)
func (_FeeCalculator *FeeCalculatorFilterer) WatchFeedFeeRemoved(opts *bind.WatchOpts, sink chan<- *FeeCalculatorFeedFeeRemoved, feedId [][21]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "FeedFeeRemoved", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorFeedFeeRemoved)
				if err := _FeeCalculator.contract.UnpackLog(event, "FeedFeeRemoved", log); err != nil {
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

// ParseFeedFeeRemoved is a log parse operation binding the contract event 0xfec81636c21a01d5af8282aefe4cd398aa5a3f2f8e85896c3373d83498317be6.
//
// Solidity: event FeedFeeRemoved(bytes21 indexed feedId)
func (_FeeCalculator *FeeCalculatorFilterer) ParseFeedFeeRemoved(log types.Log) (*FeeCalculatorFeedFeeRemoved, error) {
	event := new(FeeCalculatorFeedFeeRemoved)
	if err := _FeeCalculator.contract.UnpackLog(event, "FeedFeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorFeedFeeSetIterator is returned from FilterFeedFeeSet and is used to iterate over the raw logs and unpacked data for FeedFeeSet events raised by the FeeCalculator contract.
type FeeCalculatorFeedFeeSetIterator struct {
	Event *FeeCalculatorFeedFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorFeedFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorFeedFeeSet)
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
		it.Event = new(FeeCalculatorFeedFeeSet)
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
func (it *FeeCalculatorFeedFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorFeedFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorFeedFeeSet represents a FeedFeeSet event raised by the FeeCalculator contract.
type FeeCalculatorFeedFeeSet struct {
	FeedId [21]byte
	Fee    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeedFeeSet is a free log retrieval operation binding the contract event 0x732e2d3065b43a1e846279a2b9f63938ae3ebd6422b24c040cf2ee6a667a1b1b.
//
// Solidity: event FeedFeeSet(bytes21 indexed feedId, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) FilterFeedFeeSet(opts *bind.FilterOpts, feedId [][21]byte) (*FeeCalculatorFeedFeeSetIterator, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "FeedFeeSet", feedIdRule)
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorFeedFeeSetIterator{contract: _FeeCalculator.contract, event: "FeedFeeSet", logs: logs, sub: sub}, nil
}

// WatchFeedFeeSet is a free log subscription operation binding the contract event 0x732e2d3065b43a1e846279a2b9f63938ae3ebd6422b24c040cf2ee6a667a1b1b.
//
// Solidity: event FeedFeeSet(bytes21 indexed feedId, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) WatchFeedFeeSet(opts *bind.WatchOpts, sink chan<- *FeeCalculatorFeedFeeSet, feedId [][21]byte) (event.Subscription, error) {

	var feedIdRule []interface{}
	for _, feedIdItem := range feedId {
		feedIdRule = append(feedIdRule, feedIdItem)
	}

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "FeedFeeSet", feedIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorFeedFeeSet)
				if err := _FeeCalculator.contract.UnpackLog(event, "FeedFeeSet", log); err != nil {
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

// ParseFeedFeeSet is a log parse operation binding the contract event 0x732e2d3065b43a1e846279a2b9f63938ae3ebd6422b24c040cf2ee6a667a1b1b.
//
// Solidity: event FeedFeeSet(bytes21 indexed feedId, uint256 fee)
func (_FeeCalculator *FeeCalculatorFilterer) ParseFeedFeeSet(log types.Log) (*FeeCalculatorFeedFeeSet, error) {
	event := new(FeeCalculatorFeedFeeSet)
	if err := _FeeCalculator.contract.UnpackLog(event, "FeedFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FeeCalculator contract.
type FeeCalculatorGovernanceCallTimelockedIterator struct {
	Event *FeeCalculatorGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorGovernanceCallTimelocked)
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
		it.Event = new(FeeCalculatorGovernanceCallTimelocked)
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
func (it *FeeCalculatorGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FeeCalculator contract.
type FeeCalculatorGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FeeCalculator *FeeCalculatorFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FeeCalculatorGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorGovernanceCallTimelockedIterator{contract: _FeeCalculator.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FeeCalculator *FeeCalculatorFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FeeCalculatorGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorGovernanceCallTimelocked)
				if err := _FeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FeeCalculator *FeeCalculatorFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FeeCalculatorGovernanceCallTimelocked, error) {
	event := new(FeeCalculatorGovernanceCallTimelocked)
	if err := _FeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FeeCalculator contract.
type FeeCalculatorGovernanceInitialisedIterator struct {
	Event *FeeCalculatorGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorGovernanceInitialised)
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
		it.Event = new(FeeCalculatorGovernanceInitialised)
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
func (it *FeeCalculatorGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorGovernanceInitialised represents a GovernanceInitialised event raised by the FeeCalculator contract.
type FeeCalculatorGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FeeCalculator *FeeCalculatorFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FeeCalculatorGovernanceInitialisedIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorGovernanceInitialisedIterator{contract: _FeeCalculator.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FeeCalculator *FeeCalculatorFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FeeCalculatorGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorGovernanceInitialised)
				if err := _FeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FeeCalculator *FeeCalculatorFilterer) ParseGovernanceInitialised(log types.Log) (*FeeCalculatorGovernanceInitialised, error) {
	event := new(FeeCalculatorGovernanceInitialised)
	if err := _FeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FeeCalculator contract.
type FeeCalculatorGovernedProductionModeEnteredIterator struct {
	Event *FeeCalculatorGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorGovernedProductionModeEntered)
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
		it.Event = new(FeeCalculatorGovernedProductionModeEntered)
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
func (it *FeeCalculatorGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FeeCalculator contract.
type FeeCalculatorGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FeeCalculator *FeeCalculatorFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FeeCalculatorGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorGovernedProductionModeEnteredIterator{contract: _FeeCalculator.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FeeCalculator *FeeCalculatorFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FeeCalculatorGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorGovernedProductionModeEntered)
				if err := _FeeCalculator.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FeeCalculator *FeeCalculatorFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FeeCalculatorGovernedProductionModeEntered, error) {
	event := new(FeeCalculatorGovernedProductionModeEntered)
	if err := _FeeCalculator.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FeeCalculator contract.
type FeeCalculatorTimelockedGovernanceCallCanceledIterator struct {
	Event *FeeCalculatorTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorTimelockedGovernanceCallCanceled)
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
		it.Event = new(FeeCalculatorTimelockedGovernanceCallCanceled)
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
func (it *FeeCalculatorTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FeeCalculator contract.
type FeeCalculatorTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FeeCalculator *FeeCalculatorFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FeeCalculatorTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorTimelockedGovernanceCallCanceledIterator{contract: _FeeCalculator.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FeeCalculator *FeeCalculatorFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FeeCalculatorTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorTimelockedGovernanceCallCanceled)
				if err := _FeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FeeCalculator *FeeCalculatorFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FeeCalculatorTimelockedGovernanceCallCanceled, error) {
	event := new(FeeCalculatorTimelockedGovernanceCallCanceled)
	if err := _FeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeCalculatorTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FeeCalculator contract.
type FeeCalculatorTimelockedGovernanceCallExecutedIterator struct {
	Event *FeeCalculatorTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FeeCalculatorTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCalculatorTimelockedGovernanceCallExecuted)
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
		it.Event = new(FeeCalculatorTimelockedGovernanceCallExecuted)
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
func (it *FeeCalculatorTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCalculatorTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCalculatorTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FeeCalculator contract.
type FeeCalculatorTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FeeCalculator *FeeCalculatorFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FeeCalculatorTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FeeCalculator.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FeeCalculatorTimelockedGovernanceCallExecutedIterator{contract: _FeeCalculator.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FeeCalculator *FeeCalculatorFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FeeCalculatorTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FeeCalculator.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCalculatorTimelockedGovernanceCallExecuted)
				if err := _FeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FeeCalculator *FeeCalculatorFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FeeCalculatorTimelockedGovernanceCallExecuted, error) {
	event := new(FeeCalculatorTimelockedGovernanceCallExecuted)
	if err := _FeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
