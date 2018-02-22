// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BattlesABI is the input ABI used to generate the binding from.
const BattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countBattles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op1\",\"type\":\"uint256\"},{\"name\":\"_op2\",\"type\":\"uint256\"}],\"name\":\"createBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"op1BattleGroup\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"op2BattleGroup\",\"type\":\"uint256\"}],\"name\":\"NewBattle\",\"type\":\"event\"}]"

// BattlesBin is the compiled bytecode used for deploying new contracts.
const BattlesBin = `0x6060604052341561000f57600080fd5b6102838061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166363c785aa8114610050578063df73c49f14610075575b600080fd5b341561005b57600080fd5b61006361008e565b60405190815260200160405180910390f35b341561008057600080fd5b610063600435602435610095565b6000545b90565b600061009f6101b9565b60006080604051908101604052804267ffffffffffffffff168152602001868152602001858152602001600063ffffffff1681525091506001600080548060010182816100ec91906101e0565b600092835260209092208591600402018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151600391909101805463ffffffff191663ffffffff928316179055929091039250508116811461016257600080fd5b7fea1f5e30f8edb20fc59251726f685effaf003f76c4a416c179f62bda57477116818360200151846040015160405180848152602001838152602001828152602001935050505060405180910390a1949350505050565b60806040519081016040908152600080835260208301819052908201819052606082015290565b81548183558181151161020c5760040281600402836000526020600020918201910161020c9190610211565b505050565b61009291905b8082111561025357805467ffffffffffffffff19168155600060018201819055600282015560038101805463ffffffff19169055600401610217565b50905600a165627a7a723058207c20e0a92581d3bf0bac0dc6a92a0080c8cb8154e66cf0226d42bc7cef6aa5520029`

// DeployBattles deploys a new Ethereum contract, binding an instance of Battles to it.
func DeployBattles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Battles, error) {
	parsed, err := abi.JSON(strings.NewReader(BattlesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BattlesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Battles{BattlesCaller: BattlesCaller{contract: contract}, BattlesTransactor: BattlesTransactor{contract: contract}, BattlesFilterer: BattlesFilterer{contract: contract}}, nil
}

// Battles is an auto generated Go binding around an Ethereum contract.
type Battles struct {
	BattlesCaller     // Read-only binding to the contract
	BattlesTransactor // Write-only binding to the contract
	BattlesFilterer   // Log filterer for contract events
}

// BattlesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BattlesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattlesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BattlesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattlesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BattlesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattlesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BattlesSession struct {
	Contract     *Battles          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BattlesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BattlesCallerSession struct {
	Contract *BattlesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BattlesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BattlesTransactorSession struct {
	Contract     *BattlesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BattlesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BattlesRaw struct {
	Contract *Battles // Generic contract binding to access the raw methods on
}

// BattlesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BattlesCallerRaw struct {
	Contract *BattlesCaller // Generic read-only contract binding to access the raw methods on
}

// BattlesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BattlesTransactorRaw struct {
	Contract *BattlesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBattles creates a new instance of Battles, bound to a specific deployed contract.
func NewBattles(address common.Address, backend bind.ContractBackend) (*Battles, error) {
	contract, err := bindBattles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Battles{BattlesCaller: BattlesCaller{contract: contract}, BattlesTransactor: BattlesTransactor{contract: contract}, BattlesFilterer: BattlesFilterer{contract: contract}}, nil
}

// NewBattlesCaller creates a new read-only instance of Battles, bound to a specific deployed contract.
func NewBattlesCaller(address common.Address, caller bind.ContractCaller) (*BattlesCaller, error) {
	contract, err := bindBattles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BattlesCaller{contract: contract}, nil
}

// NewBattlesTransactor creates a new write-only instance of Battles, bound to a specific deployed contract.
func NewBattlesTransactor(address common.Address, transactor bind.ContractTransactor) (*BattlesTransactor, error) {
	contract, err := bindBattles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BattlesTransactor{contract: contract}, nil
}

// NewBattlesFilterer creates a new log filterer instance of Battles, bound to a specific deployed contract.
func NewBattlesFilterer(address common.Address, filterer bind.ContractFilterer) (*BattlesFilterer, error) {
	contract, err := bindBattles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BattlesFilterer{contract: contract}, nil
}

// bindBattles binds a generic wrapper to an already deployed contract.
func bindBattles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BattlesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Battles *BattlesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Battles.Contract.BattlesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Battles *BattlesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battles.Contract.BattlesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Battles *BattlesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Battles.Contract.BattlesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Battles *BattlesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Battles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Battles *BattlesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Battles *BattlesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Battles.Contract.contract.Transact(opts, method, params...)
}

// CountBattles is a free data retrieval call binding the contract method 0x63c785aa.
//
// Solidity: function countBattles() constant returns(uint256)
func (_Battles *BattlesCaller) CountBattles(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Battles.contract.Call(opts, out, "countBattles")
	return *ret0, err
}

// CountBattles is a free data retrieval call binding the contract method 0x63c785aa.
//
// Solidity: function countBattles() constant returns(uint256)
func (_Battles *BattlesSession) CountBattles() (*big.Int, error) {
	return _Battles.Contract.CountBattles(&_Battles.CallOpts)
}

// CountBattles is a free data retrieval call binding the contract method 0x63c785aa.
//
// Solidity: function countBattles() constant returns(uint256)
func (_Battles *BattlesCallerSession) CountBattles() (*big.Int, error) {
	return _Battles.Contract.CountBattles(&_Battles.CallOpts)
}

// CreateBattle is a paid mutator transaction binding the contract method 0xdf73c49f.
//
// Solidity: function createBattle(_op1 uint256, _op2 uint256) returns(uint256)
func (_Battles *BattlesTransactor) CreateBattle(opts *bind.TransactOpts, _op1 *big.Int, _op2 *big.Int) (*types.Transaction, error) {
	return _Battles.contract.Transact(opts, "createBattle", _op1, _op2)
}

// CreateBattle is a paid mutator transaction binding the contract method 0xdf73c49f.
//
// Solidity: function createBattle(_op1 uint256, _op2 uint256) returns(uint256)
func (_Battles *BattlesSession) CreateBattle(_op1 *big.Int, _op2 *big.Int) (*types.Transaction, error) {
	return _Battles.Contract.CreateBattle(&_Battles.TransactOpts, _op1, _op2)
}

// CreateBattle is a paid mutator transaction binding the contract method 0xdf73c49f.
//
// Solidity: function createBattle(_op1 uint256, _op2 uint256) returns(uint256)
func (_Battles *BattlesTransactorSession) CreateBattle(_op1 *big.Int, _op2 *big.Int) (*types.Transaction, error) {
	return _Battles.Contract.CreateBattle(&_Battles.TransactOpts, _op1, _op2)
}

// BattlesNewBattleIterator is returned from FilterNewBattle and is used to iterate over the raw logs and unpacked data for NewBattle events raised by the Battles contract.
type BattlesNewBattleIterator struct {
	Event *BattlesNewBattle // Event containing the contract specifics and raw log

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
func (it *BattlesNewBattleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattlesNewBattle)
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
		it.Event = new(BattlesNewBattle)
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
func (it *BattlesNewBattleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattlesNewBattleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattlesNewBattle represents a NewBattle event raised by the Battles contract.
type BattlesNewBattle struct {
	BattleID       *big.Int
	Op1BattleGroup *big.Int
	Op2BattleGroup *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewBattle is a free log retrieval operation binding the contract event 0xea1f5e30f8edb20fc59251726f685effaf003f76c4a416c179f62bda57477116.
//
// Solidity: event NewBattle(battleID uint256, op1BattleGroup uint256, op2BattleGroup uint256)
func (_Battles *BattlesFilterer) FilterNewBattle(opts *bind.FilterOpts) (*BattlesNewBattleIterator, error) {

	logs, sub, err := _Battles.contract.FilterLogs(opts, "NewBattle")
	if err != nil {
		return nil, err
	}
	return &BattlesNewBattleIterator{contract: _Battles.contract, event: "NewBattle", logs: logs, sub: sub}, nil
}

// WatchNewBattle is a free log subscription operation binding the contract event 0xea1f5e30f8edb20fc59251726f685effaf003f76c4a416c179f62bda57477116.
//
// Solidity: event NewBattle(battleID uint256, op1BattleGroup uint256, op2BattleGroup uint256)
func (_Battles *BattlesFilterer) WatchNewBattle(opts *bind.WatchOpts, sink chan<- *BattlesNewBattle) (event.Subscription, error) {

	logs, sub, err := _Battles.contract.WatchLogs(opts, "NewBattle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattlesNewBattle)
				if err := _Battles.contract.UnpackLog(event, "NewBattle", log); err != nil {
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
