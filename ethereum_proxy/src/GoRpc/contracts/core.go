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

// BattleGroupsABI is the input ABI used to generate the binding from.
const BattleGroupsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countBattleGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_CARDS_PER_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_cards\",\"type\":\"uint256[5]\"}],\"name\":\"createBattleGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"battleGroupID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cards\",\"type\":\"uint256[5]\"}],\"name\":\"NewBattleGroup\",\"type\":\"event\"}]"

// BattleGroupsBin is the compiled bytecode used for deploying new contracts.
const BattleGroupsBin = `0x6060604052341561000f57600080fd5b6103ed8061001e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637bf13d828114610052578063bfeea70814610077578063fa74efc11461008a57600080fd5b341561005d57600080fd5b6100656100b8565b60405190815260200160405180910390f35b341561008257600080fd5b6100656100bf565b341561009557600080fd5b61006573ffffffffffffffffffffffffffffffffffffffff6004351660246100c4565b6000545b90565b600581565b60006100ce61027c565b60006060604051908101604052804267ffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161014283826102a2565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815473ffffffffffffffffffffffffffffffffffffffff9190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516101d890600183019060056102d3565b50505003905063ffffffff811681146101f057600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d8582846040015160405173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052604081018260a080838360005b83811015610260578082015183820152602001610248565b50505050905001935050505060405180910390a1949350505050565b60e0604051908101604090815260008083526020830152810161029d610311565b905290565b8154818355818115116102ce576006028160060283600052602060002091820191016102ce9190610338565b505050565b8260058101928215610301579160200282015b828111156103015782518255916020019190600101906102e6565b5061030d929150610384565b5090565b60a06040519081016040526005815b60008152602001906001900390816103205790505090565b6100bc91905b8082111561030d5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061037b600183018261039e565b5060060161033e565b6100bc91905b8082111561030d576000815560010161038a565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a72305820eef3b01b12bd3f266c2f504ccc7120c574c8718b546256a3989165c9605814120029`

// DeployBattleGroups deploys a new Ethereum contract, binding an instance of BattleGroups to it.
func DeployBattleGroups(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BattleGroups, error) {
	parsed, err := abi.JSON(strings.NewReader(BattleGroupsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BattleGroupsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BattleGroups{BattleGroupsCaller: BattleGroupsCaller{contract: contract}, BattleGroupsTransactor: BattleGroupsTransactor{contract: contract}, BattleGroupsFilterer: BattleGroupsFilterer{contract: contract}}, nil
}

// BattleGroups is an auto generated Go binding around an Ethereum contract.
type BattleGroups struct {
	BattleGroupsCaller     // Read-only binding to the contract
	BattleGroupsTransactor // Write-only binding to the contract
	BattleGroupsFilterer   // Log filterer for contract events
}

// BattleGroupsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BattleGroupsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleGroupsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BattleGroupsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleGroupsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BattleGroupsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleGroupsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BattleGroupsSession struct {
	Contract     *BattleGroups     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BattleGroupsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BattleGroupsCallerSession struct {
	Contract *BattleGroupsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BattleGroupsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BattleGroupsTransactorSession struct {
	Contract     *BattleGroupsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BattleGroupsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BattleGroupsRaw struct {
	Contract *BattleGroups // Generic contract binding to access the raw methods on
}

// BattleGroupsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BattleGroupsCallerRaw struct {
	Contract *BattleGroupsCaller // Generic read-only contract binding to access the raw methods on
}

// BattleGroupsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BattleGroupsTransactorRaw struct {
	Contract *BattleGroupsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBattleGroups creates a new instance of BattleGroups, bound to a specific deployed contract.
func NewBattleGroups(address common.Address, backend bind.ContractBackend) (*BattleGroups, error) {
	contract, err := bindBattleGroups(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BattleGroups{BattleGroupsCaller: BattleGroupsCaller{contract: contract}, BattleGroupsTransactor: BattleGroupsTransactor{contract: contract}, BattleGroupsFilterer: BattleGroupsFilterer{contract: contract}}, nil
}

// NewBattleGroupsCaller creates a new read-only instance of BattleGroups, bound to a specific deployed contract.
func NewBattleGroupsCaller(address common.Address, caller bind.ContractCaller) (*BattleGroupsCaller, error) {
	contract, err := bindBattleGroups(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BattleGroupsCaller{contract: contract}, nil
}

// NewBattleGroupsTransactor creates a new write-only instance of BattleGroups, bound to a specific deployed contract.
func NewBattleGroupsTransactor(address common.Address, transactor bind.ContractTransactor) (*BattleGroupsTransactor, error) {
	contract, err := bindBattleGroups(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BattleGroupsTransactor{contract: contract}, nil
}

// NewBattleGroupsFilterer creates a new log filterer instance of BattleGroups, bound to a specific deployed contract.
func NewBattleGroupsFilterer(address common.Address, filterer bind.ContractFilterer) (*BattleGroupsFilterer, error) {
	contract, err := bindBattleGroups(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BattleGroupsFilterer{contract: contract}, nil
}

// bindBattleGroups binds a generic wrapper to an already deployed contract.
func bindBattleGroups(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BattleGroupsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BattleGroups *BattleGroupsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BattleGroups.Contract.BattleGroupsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BattleGroups *BattleGroupsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BattleGroups.Contract.BattleGroupsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BattleGroups *BattleGroupsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BattleGroups.Contract.BattleGroupsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BattleGroups *BattleGroupsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BattleGroups.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BattleGroups *BattleGroupsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BattleGroups.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BattleGroups *BattleGroupsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BattleGroups.Contract.contract.Transact(opts, method, params...)
}

// MAXCARDSPERGROUP is a free data retrieval call binding the contract method 0xbfeea708.
//
// Solidity: function MAX_CARDS_PER_GROUP() constant returns(uint256)
func (_BattleGroups *BattleGroupsCaller) MAXCARDSPERGROUP(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "MAX_CARDS_PER_GROUP")
	return *ret0, err
}

// MAXCARDSPERGROUP is a free data retrieval call binding the contract method 0xbfeea708.
//
// Solidity: function MAX_CARDS_PER_GROUP() constant returns(uint256)
func (_BattleGroups *BattleGroupsSession) MAXCARDSPERGROUP() (*big.Int, error) {
	return _BattleGroups.Contract.MAXCARDSPERGROUP(&_BattleGroups.CallOpts)
}

// MAXCARDSPERGROUP is a free data retrieval call binding the contract method 0xbfeea708.
//
// Solidity: function MAX_CARDS_PER_GROUP() constant returns(uint256)
func (_BattleGroups *BattleGroupsCallerSession) MAXCARDSPERGROUP() (*big.Int, error) {
	return _BattleGroups.Contract.MAXCARDSPERGROUP(&_BattleGroups.CallOpts)
}

// CountBattleGroups is a free data retrieval call binding the contract method 0x7bf13d82.
//
// Solidity: function countBattleGroups() constant returns(uint256)
func (_BattleGroups *BattleGroupsCaller) CountBattleGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "countBattleGroups")
	return *ret0, err
}

// CountBattleGroups is a free data retrieval call binding the contract method 0x7bf13d82.
//
// Solidity: function countBattleGroups() constant returns(uint256)
func (_BattleGroups *BattleGroupsSession) CountBattleGroups() (*big.Int, error) {
	return _BattleGroups.Contract.CountBattleGroups(&_BattleGroups.CallOpts)
}

// CountBattleGroups is a free data retrieval call binding the contract method 0x7bf13d82.
//
// Solidity: function countBattleGroups() constant returns(uint256)
func (_BattleGroups *BattleGroupsCallerSession) CountBattleGroups() (*big.Int, error) {
	return _BattleGroups.Contract.CountBattleGroups(&_BattleGroups.CallOpts)
}

// CreateBattleGroup is a paid mutator transaction binding the contract method 0xfa74efc1.
//
// Solidity: function createBattleGroup(_owner address, _cards uint256[5]) returns(uint256)
func (_BattleGroups *BattleGroupsTransactor) CreateBattleGroup(opts *bind.TransactOpts, _owner common.Address, _cards [5]*big.Int) (*types.Transaction, error) {
	return _BattleGroups.contract.Transact(opts, "createBattleGroup", _owner, _cards)
}

// CreateBattleGroup is a paid mutator transaction binding the contract method 0xfa74efc1.
//
// Solidity: function createBattleGroup(_owner address, _cards uint256[5]) returns(uint256)
func (_BattleGroups *BattleGroupsSession) CreateBattleGroup(_owner common.Address, _cards [5]*big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.CreateBattleGroup(&_BattleGroups.TransactOpts, _owner, _cards)
}

// CreateBattleGroup is a paid mutator transaction binding the contract method 0xfa74efc1.
//
// Solidity: function createBattleGroup(_owner address, _cards uint256[5]) returns(uint256)
func (_BattleGroups *BattleGroupsTransactorSession) CreateBattleGroup(_owner common.Address, _cards [5]*big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.CreateBattleGroup(&_BattleGroups.TransactOpts, _owner, _cards)
}

// BattleGroupsNewBattleGroupIterator is returned from FilterNewBattleGroup and is used to iterate over the raw logs and unpacked data for NewBattleGroup events raised by the BattleGroups contract.
type BattleGroupsNewBattleGroupIterator struct {
	Event *BattleGroupsNewBattleGroup // Event containing the contract specifics and raw log

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
func (it *BattleGroupsNewBattleGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleGroupsNewBattleGroup)
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
		it.Event = new(BattleGroupsNewBattleGroup)
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
func (it *BattleGroupsNewBattleGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleGroupsNewBattleGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleGroupsNewBattleGroup represents a NewBattleGroup event raised by the BattleGroups contract.
type BattleGroupsNewBattleGroup struct {
	Owner         common.Address
	BattleGroupID *big.Int
	Cards         [5]*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBattleGroup is a free log retrieval operation binding the contract event 0xf7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d.
//
// Solidity: event NewBattleGroup(owner address, battleGroupID uint256, cards uint256[5])
func (_BattleGroups *BattleGroupsFilterer) FilterNewBattleGroup(opts *bind.FilterOpts) (*BattleGroupsNewBattleGroupIterator, error) {

	logs, sub, err := _BattleGroups.contract.FilterLogs(opts, "NewBattleGroup")
	if err != nil {
		return nil, err
	}
	return &BattleGroupsNewBattleGroupIterator{contract: _BattleGroups.contract, event: "NewBattleGroup", logs: logs, sub: sub}, nil
}

// WatchNewBattleGroup is a free log subscription operation binding the contract event 0xf7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d.
//
// Solidity: event NewBattleGroup(owner address, battleGroupID uint256, cards uint256[5])
func (_BattleGroups *BattleGroupsFilterer) WatchNewBattleGroup(opts *bind.WatchOpts, sink chan<- *BattleGroupsNewBattleGroup) (event.Subscription, error) {

	logs, sub, err := _BattleGroups.contract.WatchLogs(opts, "NewBattleGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleGroupsNewBattleGroup)
				if err := _BattleGroups.contract.UnpackLog(event, "NewBattleGroup", log); err != nil {
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

// BattlesABI is the input ABI used to generate the binding from.
const BattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countBattles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op1\",\"type\":\"uint256\"},{\"name\":\"_op2\",\"type\":\"uint256\"}],\"name\":\"createBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerBattleGroup\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loserBattleGroup\",\"type\":\"uint256\"}],\"name\":\"BattleResult\",\"type\":\"event\"}]"

// BattlesBin is the compiled bytecode used for deploying new contracts.
const BattlesBin = `0x6060604052341561000f57600080fd5b61027a8061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166363c785aa8114610047578063df73c49f1461006c57600080fd5b341561005257600080fd5b61005a610085565b60405190815260200160405180910390f35b341561007757600080fd5b61005a60043560243561008c565b6000545b90565b60006100966101b0565b60006080604051908101604052804267ffffffffffffffff168152602001868152602001858152602001600063ffffffff1681525091506001600080548060010182816100e391906101d7565b600092835260209092208591600402018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151600391909101805463ffffffff191663ffffffff928316179055929091039250508116811461015957600080fd5b7f17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b818360200151846040015160405180848152602001838152602001828152602001935050505060405180910390a1949350505050565b60806040519081016040908152600080835260208301819052908201819052606082015290565b815481835581811511610203576004028160040283600052602060002091820191016102039190610208565b505050565b61008991905b8082111561024a57805467ffffffffffffffff19168155600060018201819055600282015560038101805463ffffffff1916905560040161020e565b50905600a165627a7a72305820cce188ef7fa62d34e1e0776202dbee41aa398af8f94daa9713794cd801d44ea50029`

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

// BattlesBattleResultIterator is returned from FilterBattleResult and is used to iterate over the raw logs and unpacked data for BattleResult events raised by the Battles contract.
type BattlesBattleResultIterator struct {
	Event *BattlesBattleResult // Event containing the contract specifics and raw log

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
func (it *BattlesBattleResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattlesBattleResult)
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
		it.Event = new(BattlesBattleResult)
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
func (it *BattlesBattleResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattlesBattleResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattlesBattleResult represents a BattleResult event raised by the Battles contract.
type BattlesBattleResult struct {
	BattleID          *big.Int
	WinnerBattleGroup *big.Int
	LoserBattleGroup  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBattleResult is a free log retrieval operation binding the contract event 0x17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b.
//
// Solidity: event BattleResult(battleID uint256, winnerBattleGroup uint256, loserBattleGroup uint256)
func (_Battles *BattlesFilterer) FilterBattleResult(opts *bind.FilterOpts) (*BattlesBattleResultIterator, error) {

	logs, sub, err := _Battles.contract.FilterLogs(opts, "BattleResult")
	if err != nil {
		return nil, err
	}
	return &BattlesBattleResultIterator{contract: _Battles.contract, event: "BattleResult", logs: logs, sub: sub}, nil
}

// WatchBattleResult is a free log subscription operation binding the contract event 0x17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b.
//
// Solidity: event BattleResult(battleID uint256, winnerBattleGroup uint256, loserBattleGroup uint256)
func (_Battles *BattlesFilterer) WatchBattleResult(opts *bind.WatchOpts, sink chan<- *BattlesBattleResult) (event.Subscription, error) {

	logs, sub, err := _Battles.contract.WatchLogs(opts, "BattleResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattlesBattleResult)
				if err := _Battles.contract.UnpackLog(event, "BattleResult", log); err != nil {
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

// CardBaseABI is the input ABI used to generate the binding from.
const CardBaseABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardBaseBin is the compiled bytecode used for deploying new contracts.
const CardBaseBin = `0x6060604052341561000f57600080fd5b6101498061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461004757806378d0df501461008857600080fd5b341561005257600080fd5b61007673ffffffffffffffffffffffffffffffffffffffff600435166024356100c7565b60405190815260200160405180910390f35b341561009357600080fd5b61009e6004356100f5565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6002602052816000526040600020818154811015156100e257fe5b6000918252602090912001549150829050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820da9ee75630dd5dab3be6ce7cd7ecdf9c6a743c4f329476d73630c00f1b67d4800029`

// DeployCardBase deploys a new Ethereum contract, binding an instance of CardBase to it.
func DeployCardBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CardBase, error) {
	parsed, err := abi.JSON(strings.NewReader(CardBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CardBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CardBase{CardBaseCaller: CardBaseCaller{contract: contract}, CardBaseTransactor: CardBaseTransactor{contract: contract}, CardBaseFilterer: CardBaseFilterer{contract: contract}}, nil
}

// CardBase is an auto generated Go binding around an Ethereum contract.
type CardBase struct {
	CardBaseCaller     // Read-only binding to the contract
	CardBaseTransactor // Write-only binding to the contract
	CardBaseFilterer   // Log filterer for contract events
}

// CardBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardBaseSession struct {
	Contract     *CardBase         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardBaseCallerSession struct {
	Contract *CardBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CardBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardBaseTransactorSession struct {
	Contract     *CardBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CardBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardBaseRaw struct {
	Contract *CardBase // Generic contract binding to access the raw methods on
}

// CardBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardBaseCallerRaw struct {
	Contract *CardBaseCaller // Generic read-only contract binding to access the raw methods on
}

// CardBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardBaseTransactorRaw struct {
	Contract *CardBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardBase creates a new instance of CardBase, bound to a specific deployed contract.
func NewCardBase(address common.Address, backend bind.ContractBackend) (*CardBase, error) {
	contract, err := bindCardBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardBase{CardBaseCaller: CardBaseCaller{contract: contract}, CardBaseTransactor: CardBaseTransactor{contract: contract}, CardBaseFilterer: CardBaseFilterer{contract: contract}}, nil
}

// NewCardBaseCaller creates a new read-only instance of CardBase, bound to a specific deployed contract.
func NewCardBaseCaller(address common.Address, caller bind.ContractCaller) (*CardBaseCaller, error) {
	contract, err := bindCardBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardBaseCaller{contract: contract}, nil
}

// NewCardBaseTransactor creates a new write-only instance of CardBase, bound to a specific deployed contract.
func NewCardBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*CardBaseTransactor, error) {
	contract, err := bindCardBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardBaseTransactor{contract: contract}, nil
}

// NewCardBaseFilterer creates a new log filterer instance of CardBase, bound to a specific deployed contract.
func NewCardBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*CardBaseFilterer, error) {
	contract, err := bindCardBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardBaseFilterer{contract: contract}, nil
}

// bindCardBase binds a generic wrapper to an already deployed contract.
func bindCardBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardBase *CardBaseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardBase.Contract.CardBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardBase *CardBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBase.Contract.CardBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardBase *CardBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardBase.Contract.CardBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardBase *CardBaseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardBase *CardBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardBase *CardBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardBase.Contract.contract.Transact(opts, method, params...)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBase *CardBaseCaller) CardIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBase.contract.Call(opts, out, "cardIndexToOwner", arg0)
	return *ret0, err
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBase *CardBaseSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardBase.Contract.CardIndexToOwner(&_CardBase.CallOpts, arg0)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBase *CardBaseCallerSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardBase.Contract.CardIndexToOwner(&_CardBase.CallOpts, arg0)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBase *CardBaseCaller) CardsHeldByOwner(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardBase.contract.Call(opts, out, "cardsHeldByOwner", arg0, arg1)
	return *ret0, err
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBase *CardBaseSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardBase.Contract.CardsHeldByOwner(&_CardBase.CallOpts, arg0, arg1)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBase *CardBaseCallerSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardBase.Contract.CardsHeldByOwner(&_CardBase.CallOpts, arg0, arg1)
}

// CardBaseNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the CardBase contract.
type CardBaseNewCardIterator struct {
	Event *CardBaseNewCard // Event containing the contract specifics and raw log

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
func (it *CardBaseNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBaseNewCard)
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
		it.Event = new(CardBaseNewCard)
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
func (it *CardBaseNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBaseNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBaseNewCard represents a NewCard event raised by the CardBase contract.
type CardBaseNewCard struct {
	Owner            common.Address
	CardID           *big.Int
	CreationBattleID *big.Int
	Attributes       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardBase *CardBaseFilterer) FilterNewCard(opts *bind.FilterOpts) (*CardBaseNewCardIterator, error) {

	logs, sub, err := _CardBase.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &CardBaseNewCardIterator{contract: _CardBase.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardBase *CardBaseFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *CardBaseNewCard) (event.Subscription, error) {

	logs, sub, err := _CardBase.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBaseNewCard)
				if err := _CardBase.contract.UnpackLog(event, "NewCard", log); err != nil {
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

// CardBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CardBase contract.
type CardBaseTransferIterator struct {
	Event *CardBaseTransfer // Event containing the contract specifics and raw log

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
func (it *CardBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBaseTransfer)
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
		it.Event = new(CardBaseTransfer)
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
func (it *CardBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBaseTransfer represents a Transfer event raised by the CardBase contract.
type CardBaseTransfer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardBase *CardBaseFilterer) FilterTransfer(opts *bind.FilterOpts) (*CardBaseTransferIterator, error) {

	logs, sub, err := _CardBase.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CardBaseTransferIterator{contract: _CardBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardBase *CardBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardBaseTransfer) (event.Subscription, error) {

	logs, sub, err := _CardBase.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBaseTransfer)
				if err := _CardBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// CardMintingABI is the input ABI used to generate the binding from.
const CardMintingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"isSuccess\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardMintingBin is the compiled bytecode used for deploying new contracts.
const CardMintingBin = `0x6060604052341561000f57600080fd5b6109618061001e6000396000f300606060405236156100ac5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b157806312c2debf1461013b57806318160ddd1461016f57806345812958146101825780635de038b1146101955780636352211e146101b757806370a08231146101e957806378d0df50146102085780638462151c1461021e57806395d89b4114610290578063a9059cbb146102a3575b600080fd5b34156100bc57600080fd5b6100c46102c7565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101005780820151838201526020016100e8565b50505050905090810190601f16801561012d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014657600080fd5b61015d600160a060020a03600435166024356102fe565b60405190815260200160405180910390f35b341561017a57600080fd5b61015d61032c565b341561018d57600080fd5b61015d610333565b34156101a057600080fd5b61015d600160a060020a036004351660243561034b565b34156101c257600080fd5b6101cd600435610383565b604051600160a060020a03909116815260200160405180910390f35b34156101f457600080fd5b61015d600160a060020a03600435166103ac565b341561021357600080fd5b6101cd6004356103c7565b341561022957600080fd5b61023d600160a060020a03600435166103e2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561027c578082015183820152602001610264565b505050509050019250505060405180910390f35b341561029b57600080fd5b6100c46104c3565b34156102ae57600080fd5b6102c5600160a060020a03600435166024356104fa565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60026020528160005260406000208181548110151561031957fe5b6000918252602090912001549150829050565b6000545b90565b600061034360006201e24061034b565b506001905090565b60008030600160a060020a031684600160a060020a03161415151561036f57600080fd5b61037b60008486610554565b949350505050565b600081815260016020526040902054600160a060020a03168015156103a757600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103ea610892565b60006103f4610892565b6000806000610402876103ac565b945084151561043257600060405180591061041a5750595b908082528060200260200182016040525095506104b9565b846040518059106104405750595b9080825280602002602001820160405250935061045b61032c565b925060009150600190505b8281116104b557600081815260016020526040902054600160a060020a03888116911614156104ad578084838151811061049c57fe5b602090810290910101526001909101905b600101610466565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a038216151561050f57600080fd5b30600160a060020a031682600160a060020a03161415151561053057600080fd5b61053a33826107a9565b151561054557600080fd5b6105503383836107c9565b5050565b600061055e6108a4565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105f391906108d9565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff8116811461072057600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a16107a0600085836107c9565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561083d57600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b81548183558181151161090557600302816003028360005260206000209182019101610905919061090a565b505050565b61033091905b80821115610931576000808255600182018190556002820155600301610910565b50905600a165627a7a72305820a8bc5e2bce638f8f275aee6e6e1be73ff27729cbf75deab89794309eb87390800029`

// DeployCardMinting deploys a new Ethereum contract, binding an instance of CardMinting to it.
func DeployCardMinting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CardMinting, error) {
	parsed, err := abi.JSON(strings.NewReader(CardMintingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CardMintingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CardMinting{CardMintingCaller: CardMintingCaller{contract: contract}, CardMintingTransactor: CardMintingTransactor{contract: contract}, CardMintingFilterer: CardMintingFilterer{contract: contract}}, nil
}

// CardMinting is an auto generated Go binding around an Ethereum contract.
type CardMinting struct {
	CardMintingCaller     // Read-only binding to the contract
	CardMintingTransactor // Write-only binding to the contract
	CardMintingFilterer   // Log filterer for contract events
}

// CardMintingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardMintingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardMintingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardMintingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardMintingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardMintingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardMintingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardMintingSession struct {
	Contract     *CardMinting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardMintingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardMintingCallerSession struct {
	Contract *CardMintingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CardMintingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardMintingTransactorSession struct {
	Contract     *CardMintingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CardMintingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardMintingRaw struct {
	Contract *CardMinting // Generic contract binding to access the raw methods on
}

// CardMintingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardMintingCallerRaw struct {
	Contract *CardMintingCaller // Generic read-only contract binding to access the raw methods on
}

// CardMintingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardMintingTransactorRaw struct {
	Contract *CardMintingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardMinting creates a new instance of CardMinting, bound to a specific deployed contract.
func NewCardMinting(address common.Address, backend bind.ContractBackend) (*CardMinting, error) {
	contract, err := bindCardMinting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardMinting{CardMintingCaller: CardMintingCaller{contract: contract}, CardMintingTransactor: CardMintingTransactor{contract: contract}, CardMintingFilterer: CardMintingFilterer{contract: contract}}, nil
}

// NewCardMintingCaller creates a new read-only instance of CardMinting, bound to a specific deployed contract.
func NewCardMintingCaller(address common.Address, caller bind.ContractCaller) (*CardMintingCaller, error) {
	contract, err := bindCardMinting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardMintingCaller{contract: contract}, nil
}

// NewCardMintingTransactor creates a new write-only instance of CardMinting, bound to a specific deployed contract.
func NewCardMintingTransactor(address common.Address, transactor bind.ContractTransactor) (*CardMintingTransactor, error) {
	contract, err := bindCardMinting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardMintingTransactor{contract: contract}, nil
}

// NewCardMintingFilterer creates a new log filterer instance of CardMinting, bound to a specific deployed contract.
func NewCardMintingFilterer(address common.Address, filterer bind.ContractFilterer) (*CardMintingFilterer, error) {
	contract, err := bindCardMinting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardMintingFilterer{contract: contract}, nil
}

// bindCardMinting binds a generic wrapper to an already deployed contract.
func bindCardMinting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardMintingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardMinting *CardMintingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardMinting.Contract.CardMintingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardMinting *CardMintingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.Contract.CardMintingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardMinting *CardMintingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardMinting.Contract.CardMintingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardMinting *CardMintingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardMinting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardMinting *CardMintingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardMinting *CardMintingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardMinting.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardMinting *CardMintingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardMinting *CardMintingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardMinting.Contract.BalanceOf(&_CardMinting.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardMinting *CardMintingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardMinting.Contract.BalanceOf(&_CardMinting.CallOpts, _owner)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardMinting *CardMintingCaller) CardIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "cardIndexToOwner", arg0)
	return *ret0, err
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardMinting *CardMintingSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardMinting.Contract.CardIndexToOwner(&_CardMinting.CallOpts, arg0)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardMinting *CardMintingCallerSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardMinting.Contract.CardIndexToOwner(&_CardMinting.CallOpts, arg0)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardMinting *CardMintingCaller) CardsHeldByOwner(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "cardsHeldByOwner", arg0, arg1)
	return *ret0, err
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardMinting *CardMintingSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardMinting.Contract.CardsHeldByOwner(&_CardMinting.CallOpts, arg0, arg1)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardMinting *CardMintingCallerSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardMinting.Contract.CardsHeldByOwner(&_CardMinting.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardMinting *CardMintingCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardMinting *CardMintingSession) Name() (string, error) {
	return _CardMinting.Contract.Name(&_CardMinting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardMinting *CardMintingCallerSession) Name() (string, error) {
	return _CardMinting.Contract.Name(&_CardMinting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardMinting *CardMintingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardMinting *CardMintingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardMinting.Contract.OwnerOf(&_CardMinting.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardMinting *CardMintingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardMinting.Contract.OwnerOf(&_CardMinting.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardMinting *CardMintingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardMinting *CardMintingSession) Symbol() (string, error) {
	return _CardMinting.Contract.Symbol(&_CardMinting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardMinting *CardMintingCallerSession) Symbol() (string, error) {
	return _CardMinting.Contract.Symbol(&_CardMinting.CallOpts)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardMinting *CardMintingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardMinting *CardMintingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardMinting.Contract.TokensOfOwner(&_CardMinting.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardMinting *CardMintingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardMinting.Contract.TokensOfOwner(&_CardMinting.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardMinting *CardMintingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardMinting *CardMintingSession) TotalSupply() (*big.Int, error) {
	return _CardMinting.Contract.TotalSupply(&_CardMinting.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardMinting *CardMintingCallerSession) TotalSupply() (*big.Int, error) {
	return _CardMinting.Contract.TotalSupply(&_CardMinting.CallOpts)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardMinting *CardMintingTransactor) CreateCard(opts *bind.TransactOpts, _owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "createCard", _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardMinting *CardMintingSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.CreateCard(&_CardMinting.TransactOpts, _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardMinting *CardMintingTransactorSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.CreateCard(&_CardMinting.TransactOpts, _owner, _attributes)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CardMinting *CardMintingTransactor) CreateGen0Card(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "createGen0Card")
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CardMinting *CardMintingSession) CreateGen0Card() (*types.Transaction, error) {
	return _CardMinting.Contract.CreateGen0Card(&_CardMinting.TransactOpts)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CardMinting *CardMintingTransactorSession) CreateGen0Card() (*types.Transaction, error) {
	return _CardMinting.Contract.CreateGen0Card(&_CardMinting.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardMinting *CardMintingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "transfer", _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardMinting *CardMintingSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.Transfer(&_CardMinting.TransactOpts, _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardMinting *CardMintingTransactorSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.Transfer(&_CardMinting.TransactOpts, _to, _cardId)
}

// CardMintingNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the CardMinting contract.
type CardMintingNewCardIterator struct {
	Event *CardMintingNewCard // Event containing the contract specifics and raw log

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
func (it *CardMintingNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardMintingNewCard)
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
		it.Event = new(CardMintingNewCard)
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
func (it *CardMintingNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardMintingNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardMintingNewCard represents a NewCard event raised by the CardMinting contract.
type CardMintingNewCard struct {
	Owner            common.Address
	CardID           *big.Int
	CreationBattleID *big.Int
	Attributes       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardMinting *CardMintingFilterer) FilterNewCard(opts *bind.FilterOpts) (*CardMintingNewCardIterator, error) {

	logs, sub, err := _CardMinting.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &CardMintingNewCardIterator{contract: _CardMinting.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardMinting *CardMintingFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *CardMintingNewCard) (event.Subscription, error) {

	logs, sub, err := _CardMinting.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardMintingNewCard)
				if err := _CardMinting.contract.UnpackLog(event, "NewCard", log); err != nil {
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

// CardMintingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CardMinting contract.
type CardMintingTransferIterator struct {
	Event *CardMintingTransfer // Event containing the contract specifics and raw log

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
func (it *CardMintingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardMintingTransfer)
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
		it.Event = new(CardMintingTransfer)
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
func (it *CardMintingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardMintingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardMintingTransfer represents a Transfer event raised by the CardMinting contract.
type CardMintingTransfer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardMinting *CardMintingFilterer) FilterTransfer(opts *bind.FilterOpts) (*CardMintingTransferIterator, error) {

	logs, sub, err := _CardMinting.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CardMintingTransferIterator{contract: _CardMinting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardMinting *CardMintingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardMintingTransfer) (event.Subscription, error) {

	logs, sub, err := _CardMinting.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardMintingTransfer)
				if err := _CardMinting.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// CardOwnershipABI is the input ABI used to generate the binding from.
const CardOwnershipABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardOwnershipBin is the compiled bytecode used for deploying new contracts.
const CardOwnershipBin = `0x6060604052341561000f57600080fd5b61092b8061001e6000396000f300606060405236156100a15763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a657806312c2debf1461013057806318160ddd146101645780635de038b1146101775780636352211e1461019957806370a08231146101cb57806378d0df50146101ea5780638462151c1461020057806395d89b4114610272578063a9059cbb14610285575b600080fd5b34156100b157600080fd5b6100b96102a9565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f55780820151838201526020016100dd565b50505050905090810190601f1680156101225780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013b57600080fd5b610152600160a060020a03600435166024356102e0565b60405190815260200160405180910390f35b341561016f57600080fd5b61015261030e565b341561018257600080fd5b610152600160a060020a0360043516602435610315565b34156101a457600080fd5b6101af60043561034d565b604051600160a060020a03909116815260200160405180910390f35b34156101d657600080fd5b610152600160a060020a0360043516610376565b34156101f557600080fd5b6101af600435610391565b341561020b57600080fd5b61021f600160a060020a03600435166103ac565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561025e578082015183820152602001610246565b505050509050019250505060405180910390f35b341561027d57600080fd5b6100b961048d565b341561029057600080fd5b6102a7600160a060020a03600435166024356104c4565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156102fb57fe5b6000918252602090912001549150829050565b6000545b90565b60008030600160a060020a031684600160a060020a03161415151561033957600080fd5b6103456000848661051e565b949350505050565b600081815260016020526040902054600160a060020a031680151561037157600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103b461085c565b60006103be61085c565b60008060006103cc87610376565b94508415156103fc5760006040518059106103e45750595b90808252806020026020018201604052509550610483565b8460405180591061040a5750595b9080825280602002602001820160405250935061042561030e565b925060009150600190505b82811161047f57600081815260016020526040902054600160a060020a0388811691161415610477578084838151811061046657fe5b602090810290910101526001909101905b600101610430565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03821615156104d957600080fd5b30600160a060020a031682600160a060020a0316141515156104fa57600080fd5b6105043382610773565b151561050f57600080fd5b61051a338383610793565b5050565b600061052861086e565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105bd91906108a3565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146106ea57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161076a60008583610793565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561080757600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116108cf576003028160030283600052602060002091820191016108cf91906108d4565b505050565b61031291905b808211156108fb5760008082556001820181905560028201556003016108da565b50905600a165627a7a72305820afca59efa1db36e1e9739ac0783588731ab5b86cc174d50082129b7b7502ae190029`

// DeployCardOwnership deploys a new Ethereum contract, binding an instance of CardOwnership to it.
func DeployCardOwnership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CardOwnership, error) {
	parsed, err := abi.JSON(strings.NewReader(CardOwnershipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CardOwnershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CardOwnership{CardOwnershipCaller: CardOwnershipCaller{contract: contract}, CardOwnershipTransactor: CardOwnershipTransactor{contract: contract}, CardOwnershipFilterer: CardOwnershipFilterer{contract: contract}}, nil
}

// CardOwnership is an auto generated Go binding around an Ethereum contract.
type CardOwnership struct {
	CardOwnershipCaller     // Read-only binding to the contract
	CardOwnershipTransactor // Write-only binding to the contract
	CardOwnershipFilterer   // Log filterer for contract events
}

// CardOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardOwnershipSession struct {
	Contract     *CardOwnership    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardOwnershipCallerSession struct {
	Contract *CardOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CardOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardOwnershipTransactorSession struct {
	Contract     *CardOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CardOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardOwnershipRaw struct {
	Contract *CardOwnership // Generic contract binding to access the raw methods on
}

// CardOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardOwnershipCallerRaw struct {
	Contract *CardOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// CardOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardOwnershipTransactorRaw struct {
	Contract *CardOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardOwnership creates a new instance of CardOwnership, bound to a specific deployed contract.
func NewCardOwnership(address common.Address, backend bind.ContractBackend) (*CardOwnership, error) {
	contract, err := bindCardOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardOwnership{CardOwnershipCaller: CardOwnershipCaller{contract: contract}, CardOwnershipTransactor: CardOwnershipTransactor{contract: contract}, CardOwnershipFilterer: CardOwnershipFilterer{contract: contract}}, nil
}

// NewCardOwnershipCaller creates a new read-only instance of CardOwnership, bound to a specific deployed contract.
func NewCardOwnershipCaller(address common.Address, caller bind.ContractCaller) (*CardOwnershipCaller, error) {
	contract, err := bindCardOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardOwnershipCaller{contract: contract}, nil
}

// NewCardOwnershipTransactor creates a new write-only instance of CardOwnership, bound to a specific deployed contract.
func NewCardOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*CardOwnershipTransactor, error) {
	contract, err := bindCardOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardOwnershipTransactor{contract: contract}, nil
}

// NewCardOwnershipFilterer creates a new log filterer instance of CardOwnership, bound to a specific deployed contract.
func NewCardOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*CardOwnershipFilterer, error) {
	contract, err := bindCardOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardOwnershipFilterer{contract: contract}, nil
}

// bindCardOwnership binds a generic wrapper to an already deployed contract.
func bindCardOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardOwnership *CardOwnershipRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardOwnership.Contract.CardOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardOwnership *CardOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardOwnership.Contract.CardOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardOwnership *CardOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardOwnership.Contract.CardOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardOwnership *CardOwnershipCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardOwnership *CardOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardOwnership *CardOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardOwnership.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardOwnership *CardOwnershipCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardOwnership *CardOwnershipSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardOwnership.Contract.BalanceOf(&_CardOwnership.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardOwnership *CardOwnershipCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardOwnership.Contract.BalanceOf(&_CardOwnership.CallOpts, _owner)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardOwnership *CardOwnershipCaller) CardIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "cardIndexToOwner", arg0)
	return *ret0, err
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardOwnership *CardOwnershipSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardOwnership.Contract.CardIndexToOwner(&_CardOwnership.CallOpts, arg0)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardOwnership *CardOwnershipCallerSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardOwnership.Contract.CardIndexToOwner(&_CardOwnership.CallOpts, arg0)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardOwnership *CardOwnershipCaller) CardsHeldByOwner(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "cardsHeldByOwner", arg0, arg1)
	return *ret0, err
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardOwnership *CardOwnershipSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardOwnership.Contract.CardsHeldByOwner(&_CardOwnership.CallOpts, arg0, arg1)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardOwnership *CardOwnershipCallerSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardOwnership.Contract.CardsHeldByOwner(&_CardOwnership.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardOwnership *CardOwnershipCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardOwnership *CardOwnershipSession) Name() (string, error) {
	return _CardOwnership.Contract.Name(&_CardOwnership.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardOwnership *CardOwnershipCallerSession) Name() (string, error) {
	return _CardOwnership.Contract.Name(&_CardOwnership.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardOwnership *CardOwnershipCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardOwnership *CardOwnershipSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardOwnership.Contract.OwnerOf(&_CardOwnership.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardOwnership *CardOwnershipCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardOwnership.Contract.OwnerOf(&_CardOwnership.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardOwnership *CardOwnershipCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardOwnership *CardOwnershipSession) Symbol() (string, error) {
	return _CardOwnership.Contract.Symbol(&_CardOwnership.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardOwnership *CardOwnershipCallerSession) Symbol() (string, error) {
	return _CardOwnership.Contract.Symbol(&_CardOwnership.CallOpts)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardOwnership *CardOwnershipCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardOwnership *CardOwnershipSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardOwnership.Contract.TokensOfOwner(&_CardOwnership.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardOwnership *CardOwnershipCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardOwnership.Contract.TokensOfOwner(&_CardOwnership.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardOwnership *CardOwnershipCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardOwnership *CardOwnershipSession) TotalSupply() (*big.Int, error) {
	return _CardOwnership.Contract.TotalSupply(&_CardOwnership.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardOwnership *CardOwnershipCallerSession) TotalSupply() (*big.Int, error) {
	return _CardOwnership.Contract.TotalSupply(&_CardOwnership.CallOpts)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardOwnership *CardOwnershipTransactor) CreateCard(opts *bind.TransactOpts, _owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "createCard", _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardOwnership *CardOwnershipSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardOwnership.Contract.CreateCard(&_CardOwnership.TransactOpts, _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardOwnership *CardOwnershipTransactorSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardOwnership.Contract.CreateCard(&_CardOwnership.TransactOpts, _owner, _attributes)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardOwnership *CardOwnershipTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "transfer", _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardOwnership *CardOwnershipSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardOwnership.Contract.Transfer(&_CardOwnership.TransactOpts, _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardOwnership *CardOwnershipTransactorSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardOwnership.Contract.Transfer(&_CardOwnership.TransactOpts, _to, _cardId)
}

// CardOwnershipNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the CardOwnership contract.
type CardOwnershipNewCardIterator struct {
	Event *CardOwnershipNewCard // Event containing the contract specifics and raw log

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
func (it *CardOwnershipNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardOwnershipNewCard)
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
		it.Event = new(CardOwnershipNewCard)
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
func (it *CardOwnershipNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardOwnershipNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardOwnershipNewCard represents a NewCard event raised by the CardOwnership contract.
type CardOwnershipNewCard struct {
	Owner            common.Address
	CardID           *big.Int
	CreationBattleID *big.Int
	Attributes       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardOwnership *CardOwnershipFilterer) FilterNewCard(opts *bind.FilterOpts) (*CardOwnershipNewCardIterator, error) {

	logs, sub, err := _CardOwnership.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &CardOwnershipNewCardIterator{contract: _CardOwnership.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardOwnership *CardOwnershipFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *CardOwnershipNewCard) (event.Subscription, error) {

	logs, sub, err := _CardOwnership.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardOwnershipNewCard)
				if err := _CardOwnership.contract.UnpackLog(event, "NewCard", log); err != nil {
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

// CardOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CardOwnership contract.
type CardOwnershipTransferIterator struct {
	Event *CardOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *CardOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardOwnershipTransfer)
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
		it.Event = new(CardOwnershipTransfer)
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
func (it *CardOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardOwnershipTransfer represents a Transfer event raised by the CardOwnership contract.
type CardOwnershipTransfer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardOwnership *CardOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*CardOwnershipTransferIterator, error) {

	logs, sub, err := _CardOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CardOwnershipTransferIterator{contract: _CardOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardOwnership *CardOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _CardOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardOwnershipTransfer)
				if err := _CardOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// CryptoCardsCoreABI is the input ABI used to generate the binding from.
const CryptoCardsCoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"isSuccess\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleGroupContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CryptoCardsCoreBin is the compiled bytecode used for deploying new contracts.
const CryptoCardsCoreBin = `0x6060604052341561000f57600080fd5b60048054600160a060020a03191633600160a060020a03161790556100326100ab565b604051809103906000f080151561004857600080fd5b60058054600160a060020a031916600160a060020a03929092169190911790556100706100bb565b604051809103906000f080151561008657600080fd5b60068054600160a060020a031916600160a060020a03929092169190911790556100cb565b60405161029880610b9783390190565b60405161040b80610e2f83390190565b610abd806100da6000396000f300606060405236156100e35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461010057806312c2debf1461018a57806318160ddd146101be57806345812958146101d15780635de038b1146101e45780636352211e146102065780636af04a571461023857806370a082311461024b578063715879881461026a57806378d0df50146102895780638462151c1461029f5780638f84aa091461031157806395d89b4114610324578063a9059cbb14610337578063e01196ef14610359578063e77dad5c1461036c575b60055433600160a060020a039081169116146100fe57600080fd5b005b341561010b57600080fd5b61011361037f565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561014f578082015183820152602001610137565b50505050905090810190601f16801561017c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561019557600080fd5b6101ac600160a060020a03600435166024356103b6565b60405190815260200160405180910390f35b34156101c957600080fd5b6101ac6103e4565b34156101dc57600080fd5b6101ac6103eb565b34156101ef57600080fd5b6101ac600160a060020a0360043516602435610403565b341561021157600080fd5b61021c60043561043b565b604051600160a060020a03909116815260200160405180910390f35b341561024357600080fd5b61021c610464565b341561025657600080fd5b6101ac600160a060020a0360043516610473565b341561027557600080fd5b6100fe600160a060020a036004351661048e565b341561029457600080fd5b61021c6004356104f6565b34156102aa57600080fd5b6102be600160a060020a0360043516610511565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102fd5780820151838201526020016102e5565b505050509050019250505060405180910390f35b341561031c57600080fd5b61021c6105f2565b341561032f57600080fd5b610113610601565b341561034257600080fd5b6100fe600160a060020a0360043516602435610638565b341561036457600080fd5b61021c610692565b341561037757600080fd5b61021c6106a1565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156103d157fe5b6000918252602090912001549150829050565b6000545b90565b60006103fb60006201e240610403565b506001905090565b60008030600160a060020a031684600160a060020a03161415151561042757600080fd5b610433600084866106b0565b949350505050565b600081815260016020526040902054600160a060020a031680151561045f57600080fd5b919050565b600754600160a060020a031681565b600160a060020a031660009081526003602052604090205490565b6007805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600160205260009081526040902054600160a060020a031681565b6105196109ee565b60006105236109ee565b600080600061053187610473565b94508415156105615760006040518059106105495750595b908082528060200260200182016040525095506105e8565b8460405180591061056f5750595b9080825280602002602001820160405250935061058a6103e4565b925060009150600190505b8281116105e457600081815260016020526040902054600160a060020a03888116911614156105dc57808483815181106105cb57fe5b602090810290910101526001909101905b600101610595565b8395505b5050505050919050565b600454600160a060020a031681565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a038216151561064d57600080fd5b30600160a060020a031682600160a060020a03161415151561066e57600080fd5b6106783382610905565b151561068357600080fd5b61068e338383610925565b5050565b600554600160a060020a031690565b600654600160a060020a031690565b60006106ba610a00565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200186815250915060016000805480600101828161074f9190610a35565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff8116811461087c57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a16108fc60008583610925565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561099957600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610a6157600302816003028360005260206000209182019101610a619190610a66565b505050565b6103e891905b80821115610a8d576000808255600182018190556002820155600301610a6c565b50905600a165627a7a72305820293eac7a63b802ba165a2d1df6c47880080c65a0d217e432ff22f7141949283a00296060604052341561000f57600080fd5b61027a8061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166363c785aa8114610047578063df73c49f1461006c57600080fd5b341561005257600080fd5b61005a610085565b60405190815260200160405180910390f35b341561007757600080fd5b61005a60043560243561008c565b6000545b90565b60006100966101b0565b60006080604051908101604052804267ffffffffffffffff168152602001868152602001858152602001600063ffffffff1681525091506001600080548060010182816100e391906101d7565b600092835260209092208591600402018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151600391909101805463ffffffff191663ffffffff928316179055929091039250508116811461015957600080fd5b7f17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b818360200151846040015160405180848152602001838152602001828152602001935050505060405180910390a1949350505050565b60806040519081016040908152600080835260208301819052908201819052606082015290565b815481835581811511610203576004028160040283600052602060002091820191016102039190610208565b505050565b61008991905b8082111561024a57805467ffffffffffffffff19168155600060018201819055600282015560038101805463ffffffff1916905560040161020e565b50905600a165627a7a72305820cce188ef7fa62d34e1e0776202dbee41aa398af8f94daa9713794cd801d44ea500296060604052341561000f57600080fd5b6103ed8061001e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637bf13d828114610052578063bfeea70814610077578063fa74efc11461008a57600080fd5b341561005d57600080fd5b6100656100b8565b60405190815260200160405180910390f35b341561008257600080fd5b6100656100bf565b341561009557600080fd5b61006573ffffffffffffffffffffffffffffffffffffffff6004351660246100c4565b6000545b90565b600581565b60006100ce61027c565b60006060604051908101604052804267ffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161014283826102a2565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815473ffffffffffffffffffffffffffffffffffffffff9190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516101d890600183019060056102d3565b50505003905063ffffffff811681146101f057600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d8582846040015160405173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052604081018260a080838360005b83811015610260578082015183820152602001610248565b50505050905001935050505060405180910390a1949350505050565b60e0604051908101604090815260008083526020830152810161029d610311565b905290565b8154818355818115116102ce576006028160060283600052602060002091820191016102ce9190610338565b505050565b8260058101928215610301579160200282015b828111156103015782518255916020019190600101906102e6565b5061030d929150610384565b5090565b60a06040519081016040526005815b60008152602001906001900390816103205790505090565b6100bc91905b8082111561030d5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061037b600183018261039e565b5060060161033e565b6100bc91905b8082111561030d576000815560010161038a565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a72305820eef3b01b12bd3f266c2f504ccc7120c574c8718b546256a3989165c9605814120029`

// DeployCryptoCardsCore deploys a new Ethereum contract, binding an instance of CryptoCardsCore to it.
func DeployCryptoCardsCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptoCardsCore, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoCardsCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoCardsCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoCardsCore{CryptoCardsCoreCaller: CryptoCardsCoreCaller{contract: contract}, CryptoCardsCoreTransactor: CryptoCardsCoreTransactor{contract: contract}, CryptoCardsCoreFilterer: CryptoCardsCoreFilterer{contract: contract}}, nil
}

// CryptoCardsCore is an auto generated Go binding around an Ethereum contract.
type CryptoCardsCore struct {
	CryptoCardsCoreCaller     // Read-only binding to the contract
	CryptoCardsCoreTransactor // Write-only binding to the contract
	CryptoCardsCoreFilterer   // Log filterer for contract events
}

// CryptoCardsCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoCardsCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoCardsCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoCardsCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoCardsCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoCardsCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoCardsCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoCardsCoreSession struct {
	Contract     *CryptoCardsCore  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CryptoCardsCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoCardsCoreCallerSession struct {
	Contract *CryptoCardsCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CryptoCardsCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoCardsCoreTransactorSession struct {
	Contract     *CryptoCardsCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CryptoCardsCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoCardsCoreRaw struct {
	Contract *CryptoCardsCore // Generic contract binding to access the raw methods on
}

// CryptoCardsCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoCardsCoreCallerRaw struct {
	Contract *CryptoCardsCoreCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoCardsCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoCardsCoreTransactorRaw struct {
	Contract *CryptoCardsCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoCardsCore creates a new instance of CryptoCardsCore, bound to a specific deployed contract.
func NewCryptoCardsCore(address common.Address, backend bind.ContractBackend) (*CryptoCardsCore, error) {
	contract, err := bindCryptoCardsCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCore{CryptoCardsCoreCaller: CryptoCardsCoreCaller{contract: contract}, CryptoCardsCoreTransactor: CryptoCardsCoreTransactor{contract: contract}, CryptoCardsCoreFilterer: CryptoCardsCoreFilterer{contract: contract}}, nil
}

// NewCryptoCardsCoreCaller creates a new read-only instance of CryptoCardsCore, bound to a specific deployed contract.
func NewCryptoCardsCoreCaller(address common.Address, caller bind.ContractCaller) (*CryptoCardsCoreCaller, error) {
	contract, err := bindCryptoCardsCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreCaller{contract: contract}, nil
}

// NewCryptoCardsCoreTransactor creates a new write-only instance of CryptoCardsCore, bound to a specific deployed contract.
func NewCryptoCardsCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoCardsCoreTransactor, error) {
	contract, err := bindCryptoCardsCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreTransactor{contract: contract}, nil
}

// NewCryptoCardsCoreFilterer creates a new log filterer instance of CryptoCardsCore, bound to a specific deployed contract.
func NewCryptoCardsCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoCardsCoreFilterer, error) {
	contract, err := bindCryptoCardsCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreFilterer{contract: contract}, nil
}

// bindCryptoCardsCore binds a generic wrapper to an already deployed contract.
func bindCryptoCardsCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoCardsCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoCardsCore *CryptoCardsCoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoCardsCore.Contract.CryptoCardsCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoCardsCore *CryptoCardsCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CryptoCardsCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoCardsCore *CryptoCardsCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CryptoCardsCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoCardsCore *CryptoCardsCoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CryptoCardsCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoCardsCore *CryptoCardsCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoCardsCore *CryptoCardsCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.contract.Transact(opts, method, params...)
}

// BattleContract is a free data retrieval call binding the contract method 0xe01196ef.
//
// Solidity: function BattleContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) BattleContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "BattleContract")
	return *ret0, err
}

// BattleContract is a free data retrieval call binding the contract method 0xe01196ef.
//
// Solidity: function BattleContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) BattleContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleContract(&_CryptoCardsCore.CallOpts)
}

// BattleContract is a free data retrieval call binding the contract method 0xe01196ef.
//
// Solidity: function BattleContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) BattleContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleContract(&_CryptoCardsCore.CallOpts)
}

// BattleGroupContract is a free data retrieval call binding the contract method 0xe77dad5c.
//
// Solidity: function BattleGroupContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) BattleGroupContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "BattleGroupContract")
	return *ret0, err
}

// BattleGroupContract is a free data retrieval call binding the contract method 0xe77dad5c.
//
// Solidity: function BattleGroupContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) BattleGroupContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleGroupContract(&_CryptoCardsCore.CallOpts)
}

// BattleGroupContract is a free data retrieval call binding the contract method 0xe77dad5c.
//
// Solidity: function BattleGroupContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) BattleGroupContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleGroupContract(&_CryptoCardsCore.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CryptoCardsCore *CryptoCardsCoreCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CryptoCardsCore.Contract.BalanceOf(&_CryptoCardsCore.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CryptoCardsCore.Contract.BalanceOf(&_CryptoCardsCore.CallOpts, _owner)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) CardIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "cardIndexToOwner", arg0)
	return *ret0, err
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CryptoCardsCore.Contract.CardIndexToOwner(&_CryptoCardsCore.CallOpts, arg0)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CryptoCardsCore.Contract.CardIndexToOwner(&_CryptoCardsCore.CallOpts, arg0)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreCaller) CardsHeldByOwner(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "cardsHeldByOwner", arg0, arg1)
	return *ret0, err
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CryptoCardsCore.Contract.CardsHeldByOwner(&_CryptoCardsCore.CallOpts, arg0, arg1)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CryptoCardsCore.Contract.CardsHeldByOwner(&_CryptoCardsCore.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreSession) Name() (string, error) {
	return _CryptoCardsCore.Contract.Name(&_CryptoCardsCore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) Name() (string, error) {
	return _CryptoCardsCore.Contract.Name(&_CryptoCardsCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) NewContractAddress() (common.Address, error) {
	return _CryptoCardsCore.Contract.NewContractAddress(&_CryptoCardsCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) NewContractAddress() (common.Address, error) {
	return _CryptoCardsCore.Contract.NewContractAddress(&_CryptoCardsCore.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) OwnerAddress() (common.Address, error) {
	return _CryptoCardsCore.Contract.OwnerAddress(&_CryptoCardsCore.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) OwnerAddress() (common.Address, error) {
	return _CryptoCardsCore.Contract.OwnerAddress(&_CryptoCardsCore.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CryptoCardsCore *CryptoCardsCoreSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CryptoCardsCore.Contract.OwnerOf(&_CryptoCardsCore.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CryptoCardsCore.Contract.OwnerOf(&_CryptoCardsCore.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreSession) Symbol() (string, error) {
	return _CryptoCardsCore.Contract.Symbol(&_CryptoCardsCore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) Symbol() (string, error) {
	return _CryptoCardsCore.Contract.Symbol(&_CryptoCardsCore.CallOpts)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CryptoCardsCore *CryptoCardsCoreCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CryptoCardsCore *CryptoCardsCoreSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CryptoCardsCore.Contract.TokensOfOwner(&_CryptoCardsCore.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CryptoCardsCore.Contract.TokensOfOwner(&_CryptoCardsCore.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) TotalSupply() (*big.Int, error) {
	return _CryptoCardsCore.Contract.TotalSupply(&_CryptoCardsCore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) TotalSupply() (*big.Int, error) {
	return _CryptoCardsCore.Contract.TotalSupply(&_CryptoCardsCore.CallOpts)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreTransactor) CreateCard(opts *bind.TransactOpts, _owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "createCard", _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CreateCard(&_CryptoCardsCore.TransactOpts, _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CreateCard(&_CryptoCardsCore.TransactOpts, _owner, _attributes)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CryptoCardsCore *CryptoCardsCoreTransactor) CreateGen0Card(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "createGen0Card")
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) CreateGen0Card() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CreateGen0Card(&_CryptoCardsCore.TransactOpts)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(isSuccess uint256)
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) CreateGen0Card() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CreateGen0Card(&_CryptoCardsCore.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetNewAddress(&_CryptoCardsCore.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetNewAddress(&_CryptoCardsCore.TransactOpts, _newAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "transfer", _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Transfer(&_CryptoCardsCore.TransactOpts, _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Transfer(&_CryptoCardsCore.TransactOpts, _to, _cardId)
}

// CryptoCardsCoreContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CryptoCardsCore contract.
type CryptoCardsCoreContractUpgradeIterator struct {
	Event *CryptoCardsCoreContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CryptoCardsCoreContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoCardsCoreContractUpgrade)
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
		it.Event = new(CryptoCardsCoreContractUpgrade)
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
func (it *CryptoCardsCoreContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoCardsCoreContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoCardsCoreContractUpgrade represents a ContractUpgrade event raised by the CryptoCardsCore contract.
type CryptoCardsCoreContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CryptoCardsCoreContractUpgradeIterator, error) {

	logs, sub, err := _CryptoCardsCore.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreContractUpgradeIterator{contract: _CryptoCardsCore.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CryptoCardsCoreContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CryptoCardsCore.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoCardsCoreContractUpgrade)
				if err := _CryptoCardsCore.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// CryptoCardsCoreNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the CryptoCardsCore contract.
type CryptoCardsCoreNewCardIterator struct {
	Event *CryptoCardsCoreNewCard // Event containing the contract specifics and raw log

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
func (it *CryptoCardsCoreNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoCardsCoreNewCard)
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
		it.Event = new(CryptoCardsCoreNewCard)
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
func (it *CryptoCardsCoreNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoCardsCoreNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoCardsCoreNewCard represents a NewCard event raised by the CryptoCardsCore contract.
type CryptoCardsCoreNewCard struct {
	Owner            common.Address
	CardID           *big.Int
	CreationBattleID *big.Int
	Attributes       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) FilterNewCard(opts *bind.FilterOpts) (*CryptoCardsCoreNewCardIterator, error) {

	logs, sub, err := _CryptoCardsCore.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreNewCardIterator{contract: _CryptoCardsCore.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *CryptoCardsCoreNewCard) (event.Subscription, error) {

	logs, sub, err := _CryptoCardsCore.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoCardsCoreNewCard)
				if err := _CryptoCardsCore.contract.UnpackLog(event, "NewCard", log); err != nil {
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

// CryptoCardsCoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CryptoCardsCore contract.
type CryptoCardsCoreTransferIterator struct {
	Event *CryptoCardsCoreTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoCardsCoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoCardsCoreTransfer)
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
		it.Event = new(CryptoCardsCoreTransfer)
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
func (it *CryptoCardsCoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoCardsCoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoCardsCoreTransfer represents a Transfer event raised by the CryptoCardsCore contract.
type CryptoCardsCoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) FilterTransfer(opts *bind.FilterOpts) (*CryptoCardsCoreTransferIterator, error) {

	logs, sub, err := _CryptoCardsCore.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CryptoCardsCoreTransferIterator{contract: _CryptoCardsCore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CryptoCardsCore *CryptoCardsCoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CryptoCardsCoreTransfer) (event.Subscription, error) {

	logs, sub, err := _CryptoCardsCore.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoCardsCoreTransfer)
				if err := _CryptoCardsCore.contract.UnpackLog(event, "Transfer", log); err != nil {
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
