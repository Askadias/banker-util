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
const MultisendABI = "[{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendEth\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractToken\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"bulkSendToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MultisendBin is the compiled bytecode used for deploying new contracts.
var MultisendBin = "0x6080604052336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610d70806100536000396000f3fe60806040526004361061003e5760003560e01c8062f55d9d14610043578063046ef2bf146100945780638da5cb5b1461020d5780639bb3a22d14610264575b600080fd5b34801561004f57600080fd5b506100926004803603602081101561006657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103b0565b005b3480156100a057600080fd5b5061020b600480360360608110156100b757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100f457600080fd5b82018360208201111561010657600080fd5b8035906020019184602083028401116401000000008311171561012857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561018857600080fd5b82018360208201111561019a57600080fd5b803590602001918460208302840111640100000000831117156101bc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050509192919290505050610422565b005b34801561021957600080fd5b50610222610949565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6103ae6004803603604081101561027a57600080fd5b810190808035906020019064010000000081111561029757600080fd5b8201836020820111156102a957600080fd5b803590602001918460208302840111640100000000831117156102cb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561032b57600080fd5b82018360208201111561033d57600080fd5b8035906020019184602083028401116401000000008311171561035f57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061096e565b005b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461040957600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16ff5b61012c8251111561047e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610d156027913960400191505060405180910390fd5b80518251146104f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f706172616d6574657273206e6f74206d6174636800000000000000000000000081525060200191505060405180910390fd5b600080905060008090505b83518160ff16101561054457610535838260ff168151811061051e57fe5b602002602001015183610cd590919063ffffffff16565b91508080600101915050610500565b508373ffffffffffffffffffffffffffffffffffffffff1663095ea7b330836040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156105cc57600080fd5b505af11580156105e0573d6000803e3d6000fd5b505050506040513d60208110156105f657600080fd5b8101908080519060200190929190505050508373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060206040518083038186803b1580156106b957600080fd5b505afa1580156106cd573d6000803e3d6000fd5b505050506040513d60208110156106e357600080fd5b8101908080519060200190929190505050811115610769576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f6e6f7420656e6f75676820746f6b656e2062616c616e6365000000000000000081525060200191505060405180910390fd5b60008090505b83518160ff161015610942578473ffffffffffffffffffffffffffffffffffffffff166323b872dd33868460ff16815181106107a757fe5b6020026020010151868560ff16815181106107be57fe5b60200260200101516040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15801561086257600080fd5b505af1158015610876573d6000803e3d6000fd5b505050506040513d602081101561088c57600080fd5b8101908080519060200190929190505050507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2848260ff16815181106108ce57fe5b6020026020010151848360ff16815181106108e557fe5b6020026020010151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1808060010191505061076f565b5050505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61012c825111156109ca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526027815260200180610d156027913960400191505060405180910390fd5b8051825114610a41576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f706172616d6574657273206e6f74206d6174636800000000000000000000000081525060200191505060405180910390fd5b600080905060008090505b83518160ff161015610a9057610a81838260ff1681518110610a6a57fe5b602002602001015183610cd590919063ffffffff16565b91508080600101915050610a4c565b5080341015610b07576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f6e6f7420656e6f75676820746f6b656e0000000000000000000000000000000081525060200191505060405180910390fd5b60008090505b83518160ff161015610c3057838160ff1681518110610b2857fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff166108fc848360ff1681518110610b5857fe5b60200260200101519081150290604051600060405180830381858888f19350505050158015610b8b573d6000803e3d6000fd5b507f69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2848260ff1681518110610bbc57fe5b6020026020010151848360ff1681518110610bd357fe5b6020026020010151604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a18080600101915050610b0d565b5080341115610cd0576000610c4e8234610cf490919063ffffffff16565b90503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610c96573d6000803e3d6000fd5b507f2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97816040518082815260200191505060405180910390a1505b505050565b600080828401905083811015610cea57600080fd5b8091505092915050565b600082821115610d0357600080fd5b60008284039050809150509291505056fe6e756d626572206f6620726563697069656e7473206973206c6172676572207468616e20333030a265627a7a7231582009c730c67b1f8c0f9b8ec0d2daae2dc1bc406b36dcf75122c95e64be0a7ed7fc64736f6c63430005110032"

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
// Solidity: function bulkSendEth(address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendTransactor) BulkSendEth(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "bulkSendEth", recipients, amounts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x9bb3a22d.
//
// Solidity: function bulkSendEth(address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendSession) BulkSendEth(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendEth(&_Multisend.TransactOpts, recipients, amounts)
}

// BulkSendEth is a paid mutator transaction binding the contract method 0x9bb3a22d.
//
// Solidity: function bulkSendEth(address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendTransactorSession) BulkSendEth(recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendEth(&_Multisend.TransactOpts, recipients, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address tokenAddr, address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendTransactor) BulkSendToken(opts *bind.TransactOpts, tokenAddr common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "bulkSendToken", tokenAddr, recipients, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address tokenAddr, address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendSession) BulkSendToken(tokenAddr common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, tokenAddr, recipients, amounts)
}

// BulkSendToken is a paid mutator transaction binding the contract method 0x046ef2bf.
//
// Solidity: function bulkSendToken(address tokenAddr, address[] recipients, uint256[] amounts) returns()
func (_Multisend *MultisendTransactorSession) BulkSendToken(tokenAddr common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.BulkSendToken(&_Multisend.TransactOpts, tokenAddr, recipients, amounts)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Multisend *MultisendTransactor) Destroy(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "destroy", _to)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Multisend *MultisendSession) Destroy(_to common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.Destroy(&_Multisend.TransactOpts, _to)
}

// Destroy is a paid mutator transaction binding the contract method 0x00f55d9d.
//
// Solidity: function destroy(address _to) returns()
func (_Multisend *MultisendTransactorSession) Destroy(_to common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.Destroy(&_Multisend.TransactOpts, _to)
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
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
func (_Multisend *MultisendFilterer) FilterRefund(opts *bind.FilterOpts) (*MultisendRefundIterator, error) {

	logs, sub, err := _Multisend.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &MultisendRefundIterator{contract: _Multisend.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x2e1897b0591d764356194f7a795238a87c1987c7a877568e50d829d547c92b97.
//
// Solidity: event Refund(uint256 refund)
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
// Solidity: event Refund(uint256 refund)
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
