// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package examples

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PingPongABI is the input ABI used to generate the binding from.
const PingPongABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"Ping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"funcHash\",\"type\":\"bytes4\"},{\"indexed\":false,\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"blockhash\",\"type\":\"bytes32\"}],\"name\":\"Pong\",\"type\":\"event\"}]"

// PingPongBin is the compiled bytecode used for deploying new contracts.
const PingPongBin = `0x6060604052341561000f57600080fd5b6101ba8061001e6000396000f3006060604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663ca6e822d81146100e2575b341561004b57600080fd5b7f07bbfefdb98ceca20faf6a118ca5a8f2f706f1a6d9243f7416be416ed746c3b033600035600160e060020a031916600036434060405173ffffffffffffffffffffffffffffffffffffffff86168152600160e060020a0319851660208201526060810182905260806040820181815290820184905260a082018585808284378201915050965050505050505060405180910390a1005b34156100ed57600080fd5b6100f56100f7565b005b7f07bbfefdb98ceca20faf6a118ca5a8f2f706f1a6d9243f7416be416ed746c3b033600035600160e060020a031916600036434060405173ffffffffffffffffffffffffffffffffffffffff86168152600160e060020a0319851660208201526060810182905260806040820181815290820184905260a082018585808284378201915050965050505050505060405180910390a15600a165627a7a723058202e0dae04b5e710a595b2544b155d6b01a25dd40b5d5032b4f605e44adc51e69d0029`

// DeployPingPong deploys a new Ethereum contract, binding an instance of PingPong to it.
func DeployPingPong(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PingPong, error) {
	parsed, err := abi.JSON(strings.NewReader(PingPongABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PingPongBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// PingPong is an auto generated Go binding around an Ethereum contract.
type PingPong struct {
	PingPongCaller     // Read-only binding to the contract
	PingPongTransactor // Write-only binding to the contract
	PingPongFilterer   // Log filterer for contract events
}

// PingPongCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingPongCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingPongTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingPongFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingPongSession struct {
	Contract     *PingPong         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingPongCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingPongCallerSession struct {
	Contract *PingPongCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PingPongTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingPongTransactorSession struct {
	Contract     *PingPongTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PingPongRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingPongRaw struct {
	Contract *PingPong // Generic contract binding to access the raw methods on
}

// PingPongCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingPongCallerRaw struct {
	Contract *PingPongCaller // Generic read-only contract binding to access the raw methods on
}

// PingPongTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingPongTransactorRaw struct {
	Contract *PingPongTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPingPong creates a new instance of PingPong, bound to a specific deployed contract.
func NewPingPong(address common.Address, backend bind.ContractBackend) (*PingPong, error) {
	contract, err := bindPingPong(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPong{PingPongCaller: PingPongCaller{contract: contract}, PingPongTransactor: PingPongTransactor{contract: contract}, PingPongFilterer: PingPongFilterer{contract: contract}}, nil
}

// NewPingPongCaller creates a new read-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongCaller(address common.Address, caller bind.ContractCaller) (*PingPongCaller, error) {
	contract, err := bindPingPong(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongCaller{contract: contract}, nil
}

// NewPingPongTransactor creates a new write-only instance of PingPong, bound to a specific deployed contract.
func NewPingPongTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongTransactor, error) {
	contract, err := bindPingPong(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongTransactor{contract: contract}, nil
}

// NewPingPongFilterer creates a new log filterer instance of PingPong, bound to a specific deployed contract.
func NewPingPongFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongFilterer, error) {
	contract, err := bindPingPong(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongFilterer{contract: contract}, nil
}

// bindPingPong binds a generic wrapper to an already deployed contract.
func bindPingPong(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PingPongABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.PingPongCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.PingPongTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPong *PingPongCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PingPong.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPong *PingPongTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPong *PingPongTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPong.Contract.contract.Transact(opts, method, params...)
}

// Ping is a paid mutator transaction binding the contract method 0xca6e822d.
//
// Solidity: function Ping() returns()
func (_PingPong *PingPongTransactor) Ping(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPong.contract.Transact(opts, "Ping")
}

// Ping is a paid mutator transaction binding the contract method 0xca6e822d.
//
// Solidity: function Ping() returns()
func (_PingPong *PingPongSession) Ping() (*types.Transaction, error) {
	return _PingPong.Contract.Ping(&_PingPong.TransactOpts)
}

// Ping is a paid mutator transaction binding the contract method 0xca6e822d.
//
// Solidity: function Ping() returns()
func (_PingPong *PingPongTransactorSession) Ping() (*types.Transaction, error) {
	return _PingPong.Contract.Ping(&_PingPong.TransactOpts)
}

// PingPongPongIterator is returned from FilterPong and is used to iterate over the raw logs and unpacked data for Pong events raised by the PingPong contract.
type PingPongPongIterator struct {
	Event *PingPongPong // Event containing the contract specifics and raw log

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
func (it *PingPongPongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongPong)
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
		it.Event = new(PingPongPong)
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
func (it *PingPongPongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongPongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongPong represents a Pong event raised by the PingPong contract.
type PingPongPong struct {
	Sender    common.Address
	FuncHash  [4]byte
	CallData  []byte
	Blockhash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPong is a free log retrieval operation binding the contract event 0x07bbfefdb98ceca20faf6a118ca5a8f2f706f1a6d9243f7416be416ed746c3b0.
//
// Solidity: event Pong(sender address, funcHash bytes4, callData bytes, blockhash bytes32)
func (_PingPong *PingPongFilterer) FilterPong(opts *bind.FilterOpts) (*PingPongPongIterator, error) {

	logs, sub, err := _PingPong.contract.FilterLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return &PingPongPongIterator{contract: _PingPong.contract, event: "Pong", logs: logs, sub: sub}, nil
}

// WatchPong is a free log subscription operation binding the contract event 0x07bbfefdb98ceca20faf6a118ca5a8f2f706f1a6d9243f7416be416ed746c3b0.
//
// Solidity: event Pong(sender address, funcHash bytes4, callData bytes, blockhash bytes32)
func (_PingPong *PingPongFilterer) WatchPong(opts *bind.WatchOpts, sink chan<- *PingPongPong) (event.Subscription, error) {

	logs, sub, err := _PingPong.contract.WatchLogs(opts, "Pong")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongPong)
				if err := _PingPong.contract.UnpackLog(event, "Pong", log); err != nil {
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
