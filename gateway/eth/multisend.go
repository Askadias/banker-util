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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultisendABI is the input ABI used to generate the binding from.
const MultisendABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MultisendBin is the compiled bytecode used for deploying new contracts.
var MultisendBin = "0x608060405234801561001057600080fd5b506112bd806100206000396000f3fe60806040526004361061003f5760003560e01c8063046ef2bf146100445780639bb3a22d146101c8578063b2bdfa7b1461032c578063c4d66de814610383575b600080fd5b6101ae6004803603606081101561005a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561009757600080fd5b8201836020820111156100a957600080fd5b803590602001918460208302840111640100000000831117156100cb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561012b57600080fd5b82018360208201111561013d57600080fd5b8035906020019184602083028401116401000000008311171561015f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506103d4565b604051808215151515815260200191505060405180910390f35b610312600480360360408110156101de57600080fd5b81019080803590602001906401000000008111156101fb57600080fd5b82018360208201111561020d57600080fd5b8035906020019184602083028401116401000000008311171561022f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561028f57600080fd5b8201836020820111156102a157600080fd5b803590602001918460208302840111640100000000831117156102c357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192905050506107ee565b604051808215151515815260200191505060405180910390f35b34801561033857600080fd5b50610341610b5d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561038f57600080fd5b506103d2600480360360208110156103a657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b83565b005b600061012c83511115610432576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061120a6027913960400191505060405180910390fd5b81518351146104a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f706172616d6574657273206e6f74206d6174636800000000000000000000000081525060200191505060405180910390fd5b600080905060008090505b83518160ff1610156104f8576104e9848260ff16815181106104d257fe5b602002602001015183610cc490919063ffffffff16565b915080806001019150506104b4565b50808573ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b1580156105ab57600080fd5b505afa1580156105bf573d6000803e3d6000fd5b505050506040513d60208110156105d557600080fd5b8101908080519060200190929190505050101561065a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f6e6f7420656e6f75676820555344542062616c616e636500000000000000000081525060200191505060405180910390fd5b60008090505b84518160ff161015610759576000858260ff168151811061067d57fe5b602002602001015190506106ca3382878560ff168151811061069b57fe5b60200260200101518a73ffffffffffffffffffffffffffffffffffffffff16610d4c909392919063ffffffff16565b7f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de281868460ff16815181106106fb57fe5b6020026020010151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1508080600101915050610660565b5060003411156107e2573373ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156107a9573d6000803e3d6000fd5b507f2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97346040518082815260200191505060405180910390a15b60019150509392505050565b600061012c8351111561084c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602781526020018061120a6027913960400191505060405180910390fd5b81518351146108c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f706172616d6574657273206e6f74206d6174636800000000000000000000000081525060200191505060405180910390fd5b600080905060008090505b83518160ff16101561091257610903848260ff16815181106108ec57fe5b602002602001015183610cc490919063ffffffff16565b915080806001019150506108ce565b5080341015610989576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6e6f7420656e6f756768204554482062616c616e63650000000000000000000081525060200191505060405180910390fd5b60008090505b84518160ff161015610ab257848160ff16815181106109aa57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc858360ff16815181106109da57fe5b60200260200101519081150290604051600060405180830381858888f19350505050158015610a0d573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2858260ff1681518110610a3e57fe5b6020026020010151858360ff1681518110610a5557fe5b6020026020010151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1808060010191505061098f565b5080341115610b52576000610ad08234610e5290919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610b18573d6000803e3d6000fd5b507f2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97816040518082815260200191505060405180910390a1505b600191505092915050565b603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900460ff1680610ba25750610ba1610e9c565b5b80610bb957506000809054906101000a900460ff16155b610c0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e815260200180611231602e913960400191505060405180910390fd5b60008060019054906101000a900460ff161590508015610c5e576001600060016101000a81548160ff02191690831515021790555060016000806101000a81548160ff0219169083151502179055505b81603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610cc05760008060016101000a81548160ff0219169083151502179055505b5050565b600080828401905083811015610d42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b610e4c848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610eb3565b50505050565b6000610e9483836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506110fe565b905092915050565b6000803090506000813b9050600081149250505090565b610ed28273ffffffffffffffffffffffffffffffffffffffff166111be565b610f44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b60208310610f935780518252602082019150602081019050602083039250610f70565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610ff5576040519150601f19603f3d011682016040523d82523d6000602084013e610ffa565b606091505b509150915081611072576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b6000815111156110f85780806020019051602081101561109157600080fd5b81019080805190602001909291905050506110f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061125f602a913960400191505060405180910390fd5b5b50505050565b60008383111582906111ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611170578082015181840152602081019050611155565b50505050905090810190601f16801561119d5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561120057506000801b8214155b9250505091905056fe6e756d626572206f6620726563697069656e7473206973206c6172676572207468616e20333030436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a65645361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a72315820a9b0197533ea5d47ea35655ce32289df984fa15974498499edecbe3393d25e5c64736f6c63430005110032"

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

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_Multisend *MultisendCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Multisend.contract.Call(opts, out, "_owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
func (_Multisend *MultisendSession) Owner() (common.Address, error) {
	return _Multisend.Contract.Owner(&_Multisend.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() constant returns(address)
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
// Solidity: function bulkSendToken(address token, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactor) BulkSendToken(opts *bind.TransactOpts, token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "bulkSendToken", token, addresses, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address token, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendSession) BulkSendToken(token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, token, addresses, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address token, address[] addresses, uint256[] amounts) returns(bool success)
func (_Multisend *MultisendTransactorSession) BulkSendToken(token common.Address, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, token, addresses, amounts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Multisend *MultisendTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Multisend *MultisendSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.Initialize(&_Multisend.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_Multisend *MultisendTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.Initialize(&_Multisend.TransactOpts, _admin)
}

// MultisendRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Multisend contract.
type MultisendRefundIterator struct {
	Event *MultisendRefund // Event containing the contract specifics and raw log

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
func (it *MultisendRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendRefund)
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
		it.Event = new(MultisendRefund)
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
func (it *MultisendRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendRefund represents a Refund event raised by the Multisend contract.
type MultisendRefund struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 amount)
func (_Multisend *MultisendFilterer) FilterRefund(opts *bind.FilterOpts) (*MultisendRefundIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &MultisendRefundIterator{contract: _Multisend.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 amount)
func (_Multisend *MultisendFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *MultisendRefund) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendRefund)
				if err := _Multisend.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 amount)
func (_Multisend *MultisendFilterer) ParseRefund(log types.Log) (*MultisendRefund, error) {
	event := new(MultisendRefund)
	if err := _Multisend.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MultisendTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Multisend contract.
type MultisendTransferIterator struct {
	Event *MultisendTransfer // Event containing the contract specifics and raw log

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
func (it *MultisendTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisendTransfer)
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
		it.Event = new(MultisendTransfer)
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
func (it *MultisendTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisendTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisendTransfer represents a Transfer event raised by the Multisend contract.
type MultisendTransfer struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) FilterTransfer(opts *bind.FilterOpts) (*MultisendTransferIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &MultisendTransferIterator{contract: _Multisend.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MultisendTransfer) (event.Subscription, error) {

	logs, sub, err := _Multisend.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisendTransfer)
				if err := _Multisend.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address recipient, uint256 amount)
func (_Multisend *MultisendFilterer) ParseTransfer(log types.Log) (*MultisendTransfer, error) {
	event := new(MultisendTransfer)
	if err := _Multisend.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
