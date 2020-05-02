// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultisendABI is the input ABI used to generate the binding from.
const MultisendABI = "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MultisendBin is the compiled bytecode used for deploying new contracts.
var MultisendBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e2c806100536000396000f3fe6080604052600436106100345760003560e01c8063046ef2bf146100395780638da5cb5b146101bd5780639bb3a22d14610214575b600080fd5b6101a36004803603606081101561004f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561008c57600080fd5b82018360208201111561009e57600080fd5b803590602001918460208302840111640100000000831117156100c057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561012057600080fd5b82018360208201111561013257600080fd5b8035906020019184602083028401116401000000008311171561015457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610378565b604051808215151515815260200191505060405180910390f35b3480156101c957600080fd5b506101d26104f0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61035e6004803603604081101561022a57600080fd5b810190808035906020019064010000000081111561024757600080fd5b82018360208201111561025957600080fd5b8035906020019184602083028401116401000000008311171561027b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290803590602001906401000000008111156102db57600080fd5b8201836020820111156102ed57600080fd5b8035906020019184602083028401116401000000008311171561030f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610515565b604051808215151515815260200191505060405180910390f35b600080849050600080905060008090505b84518160ff1610156103cd576103be858260ff16815181106103a757fe5b60200260200101518361067a90919063ffffffff16565b91508080600101915050610389565b50610419307fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8473ffffffffffffffffffffffffffffffffffffffff166107029092919063ffffffff16565b60008090505b85518160ff1610156104915761048433878360ff168151811061043e57fe5b6020026020010151878460ff168151811061045557fe5b60200260200101518673ffffffffffffffffffffffffffffffffffffffff16610909909392919063ffffffff16565b808060010191505061041f565b5060003411156104e3573373ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156104e1573d6000803e3d6000fd5b505b6001925050509392505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000806000905060008090505b83518160ff16101561056657610557848260ff168151811061054057fe5b60200260200101518361067a90919063ffffffff16565b91508080600101915050610522565b508034101561057457600080fd5b60008090505b84518160ff16101561060657848160ff168151811061059557fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc858360ff16815181106105c557fe5b60200260200101519081150290604051600060405180830381858888f193505050501580156105f8573d6000803e3d6000fd5b50808060010191505061057a565b508034111561066f57600061062482346109f690919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561066c573d6000803e3d6000fd5b50505b600191505092915050565b6000808284019050838110156106f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b60008114806107fc575060008373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e30856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b1580156107bf57600080fd5b505afa1580156107d3573d6000803e3d6000fd5b505050506040513d60208110156107e957600080fd5b8101908080519060200190929190505050145b610851576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610dc16036913960400191505060405180910390fd5b6109048363095ea7b360e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610a40565b505050565b6109f0846323b872dd60e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610a40565b50505050565b6000610a3883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610c8b565b905092915050565b610a5f8273ffffffffffffffffffffffffffffffffffffffff16610d4b565b610ad1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b60208310610b205780518252602082019150602081019050602083039250610afd565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610b82576040519150601f19603f3d011682016040523d82523d6000602084013e610b87565b606091505b509150915081610bff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b600081511115610c8557808060200190516020811015610c1e57600080fd5b8101908080519060200190929190505050610c84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610d97602a913960400191505060405180910390fd5b5b50505050565b6000838311158290610d38576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610cfd578082015181840152602081019050610ce2565b50505050905090810190601f168015610d2a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f9150808214158015610d8d57506000801b8214155b9250505091905056fe5361666545524332303a204552433230206f7065726174696f6e20646964206e6f7420737563636565645361666545524332303a20617070726f76652066726f6d206e6f6e2d7a65726f20746f206e6f6e2d7a65726f20616c6c6f77616e6365a26469706673582212205df5ec4c52b2c1466408eac49bd536bde64c75e6d4d8e7f6848c89f9000183b264736f6c63430006060033"

// DeployMultisend deploys a new Ethereum contract, binding an instance of Multisend to it.
func DeployMultisend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multisend, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultisendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// Multisend is an auto generated Go binding around an Ethereum contract.
type Multisend struct {
	MultisendCaller     // Read-only binding to the contract
	MultisendTransactor // Write-only binding to the contract
	MultisendFilterer   // Log filterer for contract events
}

// MultisendCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisendSession struct {
	Contract     *Multisend        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisendCallerSession struct {
	Contract *MultisendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultisendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisendTransactorSession struct {
	Contract     *MultisendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultisendRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisendRaw struct {
	Contract *Multisend // Generic contract binding to access the raw methods on
}

// MultisendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisendCallerRaw struct {
	Contract *MultisendCaller // Generic read-only contract binding to access the raw methods on
}

// MultisendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisendTransactorRaw struct {
	Contract *MultisendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisend creates a new instance of Multisend, bound to a specific deployed contract.
func NewMultisend(address common.Address, backend bind.ContractBackend) (*Multisend, error) {
	contract, err := bindMultisend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// NewMultisendCaller creates a new read-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendCaller(address common.Address, caller bind.ContractCaller) (*MultisendCaller, error) {
	contract, err := bindMultisend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendCaller{contract: contract}, nil
}

// NewMultisendTransactor creates a new write-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisendTransactor, error) {
	contract, err := bindMultisend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendTransactor{contract: contract}, nil
}

// NewMultisendFilterer creates a new log filterer instance of Multisend, bound to a specific deployed contract.
func NewMultisendFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisendFilterer, error) {
	contract, err := bindMultisend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisendFilterer{contract: contract}, nil
}

// bindMultisend binds a generic wrapper to an already deployed contract.
func bindMultisend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.MultisendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Multisend *MultisendCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Multisend.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Multisend *MultisendSession) Owner() (common.Address, error) {
	return _Multisend.Contract.Owner(&_Multisend.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Multisend *MultisendCallerSession) Owner() (common.Address, error) {
	return _Multisend.Contract.Owner(&_Multisend.CallOpts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x9bb3a22d.
//
// Solidity: function bulkSendEth(address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactor) BulkSendEth(opts *bind.TransactOpts, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "bulkSendEth", addresses, amounts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x9bb3a22d.
//
// Solidity: function bulkSendEth(address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendSession) BulkSendEth(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendEth(&_Multisend.TransactOpts, addresses, amounts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x9bb3a22d.
//
// Solidity: function bulkSendEth(address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactorSession) BulkSendEth(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendEth(&_Multisend.TransactOpts, addresses, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address asset, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactor) BulkSendToken(opts *bind.TransactOpts, asset common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "bulkSendToken", asset, addresses, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address asset, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendSession) BulkSendToken(asset common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, asset, addresses, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address asset, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactorSession) BulkSendToken(asset common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, asset, addresses, amounts)
}
