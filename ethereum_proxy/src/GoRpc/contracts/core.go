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
const BattleGroupsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoCardsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countBattleGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"groupID\",\"type\":\"uint256\"}],\"name\":\"setGroupOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_CARDS_PER_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"groupID\",\"type\":\"uint256\"}],\"name\":\"isGroupReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_cards\",\"type\":\"uint256[5]\"}],\"name\":\"createBattleGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"battleGroupID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cards\",\"type\":\"uint256[5]\"}],\"name\":\"NewBattleGroup\",\"type\":\"event\"}]"

// BattleGroupsBin is the compiled bytecode used for deploying new contracts.
const BattleGroupsBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a03161790556106ff806100396000396000f3006060604052361561005c5763ffffffff60e060020a600035041663624fe81f81146100615780637bf13d8214610090578063b42cb37d146100b5578063bfeea708146100df578063ecf9036b146100f2578063fa74efc114610108575b600080fd5b341561006c57600080fd5b610074610129565b604051600160a060020a03909116815260200160405180910390f35b341561009b57600080fd5b6100a3610138565b60405190815260200160405180910390f35b34156100c057600080fd5b6100cb60043561013f565b604051901515815260200160405180910390f35b34156100ea57600080fd5b6100a361029a565b34156100fd57600080fd5b6100cb60043561029f565b341561011357600080fd5b6100a3600160a060020a036004351660246103fd565b600154600160a060020a031681565b6000545b90565b600061014961058e565b6000610153610138565b841061015e57600080fd5b600080548590811061016c57fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116101d25750505091909252509193506000925050505b816040015150600581101561028e57600154600160a060020a031663685b7a5e6040840151836005811061022757fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561026b57600080fd5b6102c65a03f1151561027c57600080fd5b505050604051805150506001016101f7565b600192505b5050919050565b600581565b60006102a961058e565b60006102b3610138565b84106102be57600080fd5b60008054859081106102cc57fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116103325750505091909252509193506000925050505b816040015150600581101561028e57600154600160a060020a0316632109f89d6040840151836005811061038757fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103cb57600080fd5b6102c65a03f115156103dc57600080fd5b5050506040518051905015156103f55760009250610293565b600101610357565b600061040761058e565b60006060604051908101604052804267ffffffffffffffff16815260200186600160a060020a03168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161046e83826105b4565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518154600160a060020a039190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516104f790600183019060056105e5565b50505003905063ffffffff8116811461050f57600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d85828460400151604051600160a060020a038416815260208101839052604081018260a080838360005b8381101561057257808201518382015260200161055a565b50505050905001935050505060405180910390a1949350505050565b60e060405190810160409081526000808352602083015281016105af610623565b905290565b8154818355818115116105e0576006028160060283600052602060002091820191016105e0919061064a565b505050565b8260058101928215610613579160200282015b828111156106135782518255916020019190600101906105f8565b5061061f929150610696565b5090565b60a06040519081016040526005815b60008152602001906001900390816106325790505090565b61013c91905b8082111561061f5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061068d60018301826106b0565b50600601610650565b61013c91905b8082111561061f576000815560010161069c565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a723058206cc3ff584860016f2e985e2ecda1e0fb1a0ea6294c52e91101e1787baddcc7090029`

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

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleGroups *BattleGroupsCaller) CryptoCardsContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "cryptoCardsContract")
	return *ret0, err
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleGroups *BattleGroupsSession) CryptoCardsContract() (common.Address, error) {
	return _BattleGroups.Contract.CryptoCardsContract(&_BattleGroups.CallOpts)
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleGroups *BattleGroupsCallerSession) CryptoCardsContract() (common.Address, error) {
	return _BattleGroups.Contract.CryptoCardsContract(&_BattleGroups.CallOpts)
}

// IsGroupReadyForBattle is a free data retrieval call binding the contract method 0xecf9036b.
//
// Solidity: function isGroupReadyForBattle(groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsCaller) IsGroupReadyForBattle(opts *bind.CallOpts, groupID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "isGroupReadyForBattle", groupID)
	return *ret0, err
}

// IsGroupReadyForBattle is a free data retrieval call binding the contract method 0xecf9036b.
//
// Solidity: function isGroupReadyForBattle(groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsSession) IsGroupReadyForBattle(groupID *big.Int) (bool, error) {
	return _BattleGroups.Contract.IsGroupReadyForBattle(&_BattleGroups.CallOpts, groupID)
}

// IsGroupReadyForBattle is a free data retrieval call binding the contract method 0xecf9036b.
//
// Solidity: function isGroupReadyForBattle(groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsCallerSession) IsGroupReadyForBattle(groupID *big.Int) (bool, error) {
	return _BattleGroups.Contract.IsGroupReadyForBattle(&_BattleGroups.CallOpts, groupID)
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

// SetGroupOnBattleCooldown is a paid mutator transaction binding the contract method 0xb42cb37d.
//
// Solidity: function setGroupOnBattleCooldown(groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsTransactor) SetGroupOnBattleCooldown(opts *bind.TransactOpts, groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.contract.Transact(opts, "setGroupOnBattleCooldown", groupID)
}

// SetGroupOnBattleCooldown is a paid mutator transaction binding the contract method 0xb42cb37d.
//
// Solidity: function setGroupOnBattleCooldown(groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsSession) SetGroupOnBattleCooldown(groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.SetGroupOnBattleCooldown(&_BattleGroups.TransactOpts, groupID)
}

// SetGroupOnBattleCooldown is a paid mutator transaction binding the contract method 0xb42cb37d.
//
// Solidity: function setGroupOnBattleCooldown(groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsTransactorSession) SetGroupOnBattleCooldown(groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.SetGroupOnBattleCooldown(&_BattleGroups.TransactOpts, groupID)
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

// BattleQueueABI is the input ABI used to generate the binding from.
const BattleQueueABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoCardsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_battleGroupID\",\"type\":\"uint256\"}],\"name\":\"joinQueue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleGroupID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventTime\",\"type\":\"uint64\"}],\"name\":\"QueueJoined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleGroup1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"battleGroup2\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"eventTime\",\"type\":\"uint64\"}],\"name\":\"QueueMatched\",\"type\":\"event\"}]"

// BattleQueueBin is the compiled bytecode used for deploying new contracts.
const BattleQueueBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561034f806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461002e578063969494201461005d57600080fd5b341561003957600080fd5b610041610087565b604051600160a060020a03909116815260200160405180910390f35b341561006857600080fd5b610073600435610096565b604051901515815260200160405180910390f35b600154600160a060020a031681565b60006100a182610176565b15156100af57506000610171565b7f4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7824260405191825267ffffffffffffffff1660208201526040908101905180910390a16000541515610109575060008190556001610171565b7f9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a0548260005442604051928352602083019190915267ffffffffffffffff166040808301919091526060909101905180910390a161016882600054610249565b50506000805560015b919050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156101c057600080fd5b6102c65a03f115156101d157600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561022957600080fd5b6102c65a03f1151561023a57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e01196ef82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029357600080fd5b6102c65a03f115156102a457600080fd5b50505060405180519050600160a060020a031663df73c49f848460006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561030257600080fd5b6102c65a03f1151561031357600080fd5b50505060405180519493505050505600a165627a7a7230582091e420400377b31861f8fbb146d0354c712fed3dfb906e6a993c0dff417dce340029`

// DeployBattleQueue deploys a new Ethereum contract, binding an instance of BattleQueue to it.
func DeployBattleQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BattleQueue, error) {
	parsed, err := abi.JSON(strings.NewReader(BattleQueueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BattleQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BattleQueue{BattleQueueCaller: BattleQueueCaller{contract: contract}, BattleQueueTransactor: BattleQueueTransactor{contract: contract}, BattleQueueFilterer: BattleQueueFilterer{contract: contract}}, nil
}

// BattleQueue is an auto generated Go binding around an Ethereum contract.
type BattleQueue struct {
	BattleQueueCaller     // Read-only binding to the contract
	BattleQueueTransactor // Write-only binding to the contract
	BattleQueueFilterer   // Log filterer for contract events
}

// BattleQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type BattleQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BattleQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BattleQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BattleQueueSession struct {
	Contract     *BattleQueue      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BattleQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BattleQueueCallerSession struct {
	Contract *BattleQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BattleQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BattleQueueTransactorSession struct {
	Contract     *BattleQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BattleQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type BattleQueueRaw struct {
	Contract *BattleQueue // Generic contract binding to access the raw methods on
}

// BattleQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BattleQueueCallerRaw struct {
	Contract *BattleQueueCaller // Generic read-only contract binding to access the raw methods on
}

// BattleQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BattleQueueTransactorRaw struct {
	Contract *BattleQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBattleQueue creates a new instance of BattleQueue, bound to a specific deployed contract.
func NewBattleQueue(address common.Address, backend bind.ContractBackend) (*BattleQueue, error) {
	contract, err := bindBattleQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BattleQueue{BattleQueueCaller: BattleQueueCaller{contract: contract}, BattleQueueTransactor: BattleQueueTransactor{contract: contract}, BattleQueueFilterer: BattleQueueFilterer{contract: contract}}, nil
}

// NewBattleQueueCaller creates a new read-only instance of BattleQueue, bound to a specific deployed contract.
func NewBattleQueueCaller(address common.Address, caller bind.ContractCaller) (*BattleQueueCaller, error) {
	contract, err := bindBattleQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BattleQueueCaller{contract: contract}, nil
}

// NewBattleQueueTransactor creates a new write-only instance of BattleQueue, bound to a specific deployed contract.
func NewBattleQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*BattleQueueTransactor, error) {
	contract, err := bindBattleQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BattleQueueTransactor{contract: contract}, nil
}

// NewBattleQueueFilterer creates a new log filterer instance of BattleQueue, bound to a specific deployed contract.
func NewBattleQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*BattleQueueFilterer, error) {
	contract, err := bindBattleQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BattleQueueFilterer{contract: contract}, nil
}

// bindBattleQueue binds a generic wrapper to an already deployed contract.
func bindBattleQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BattleQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BattleQueue *BattleQueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BattleQueue.Contract.BattleQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BattleQueue *BattleQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BattleQueue.Contract.BattleQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BattleQueue *BattleQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BattleQueue.Contract.BattleQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BattleQueue *BattleQueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BattleQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BattleQueue *BattleQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BattleQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BattleQueue *BattleQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BattleQueue.Contract.contract.Transact(opts, method, params...)
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleQueue *BattleQueueCaller) CryptoCardsContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BattleQueue.contract.Call(opts, out, "cryptoCardsContract")
	return *ret0, err
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleQueue *BattleQueueSession) CryptoCardsContract() (common.Address, error) {
	return _BattleQueue.Contract.CryptoCardsContract(&_BattleQueue.CallOpts)
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_BattleQueue *BattleQueueCallerSession) CryptoCardsContract() (common.Address, error) {
	return _BattleQueue.Contract.CryptoCardsContract(&_BattleQueue.CallOpts)
}

// JoinQueue is a paid mutator transaction binding the contract method 0x96949420.
//
// Solidity: function joinQueue(_battleGroupID uint256) returns(bool)
func (_BattleQueue *BattleQueueTransactor) JoinQueue(opts *bind.TransactOpts, _battleGroupID *big.Int) (*types.Transaction, error) {
	return _BattleQueue.contract.Transact(opts, "joinQueue", _battleGroupID)
}

// JoinQueue is a paid mutator transaction binding the contract method 0x96949420.
//
// Solidity: function joinQueue(_battleGroupID uint256) returns(bool)
func (_BattleQueue *BattleQueueSession) JoinQueue(_battleGroupID *big.Int) (*types.Transaction, error) {
	return _BattleQueue.Contract.JoinQueue(&_BattleQueue.TransactOpts, _battleGroupID)
}

// JoinQueue is a paid mutator transaction binding the contract method 0x96949420.
//
// Solidity: function joinQueue(_battleGroupID uint256) returns(bool)
func (_BattleQueue *BattleQueueTransactorSession) JoinQueue(_battleGroupID *big.Int) (*types.Transaction, error) {
	return _BattleQueue.Contract.JoinQueue(&_BattleQueue.TransactOpts, _battleGroupID)
}

// BattleQueueQueueJoinedIterator is returned from FilterQueueJoined and is used to iterate over the raw logs and unpacked data for QueueJoined events raised by the BattleQueue contract.
type BattleQueueQueueJoinedIterator struct {
	Event *BattleQueueQueueJoined // Event containing the contract specifics and raw log

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
func (it *BattleQueueQueueJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleQueueQueueJoined)
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
		it.Event = new(BattleQueueQueueJoined)
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
func (it *BattleQueueQueueJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleQueueQueueJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleQueueQueueJoined represents a QueueJoined event raised by the BattleQueue contract.
type BattleQueueQueueJoined struct {
	BattleGroupID *big.Int
	EventTime     uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterQueueJoined is a free log retrieval operation binding the contract event 0x4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7.
//
// Solidity: event QueueJoined(battleGroupID uint256, eventTime uint64)
func (_BattleQueue *BattleQueueFilterer) FilterQueueJoined(opts *bind.FilterOpts) (*BattleQueueQueueJoinedIterator, error) {

	logs, sub, err := _BattleQueue.contract.FilterLogs(opts, "QueueJoined")
	if err != nil {
		return nil, err
	}
	return &BattleQueueQueueJoinedIterator{contract: _BattleQueue.contract, event: "QueueJoined", logs: logs, sub: sub}, nil
}

// WatchQueueJoined is a free log subscription operation binding the contract event 0x4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7.
//
// Solidity: event QueueJoined(battleGroupID uint256, eventTime uint64)
func (_BattleQueue *BattleQueueFilterer) WatchQueueJoined(opts *bind.WatchOpts, sink chan<- *BattleQueueQueueJoined) (event.Subscription, error) {

	logs, sub, err := _BattleQueue.contract.WatchLogs(opts, "QueueJoined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleQueueQueueJoined)
				if err := _BattleQueue.contract.UnpackLog(event, "QueueJoined", log); err != nil {
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

// BattleQueueQueueMatchedIterator is returned from FilterQueueMatched and is used to iterate over the raw logs and unpacked data for QueueMatched events raised by the BattleQueue contract.
type BattleQueueQueueMatchedIterator struct {
	Event *BattleQueueQueueMatched // Event containing the contract specifics and raw log

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
func (it *BattleQueueQueueMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleQueueQueueMatched)
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
		it.Event = new(BattleQueueQueueMatched)
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
func (it *BattleQueueQueueMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleQueueQueueMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleQueueQueueMatched represents a QueueMatched event raised by the BattleQueue contract.
type BattleQueueQueueMatched struct {
	BattleGroup1 *big.Int
	BattleGroup2 *big.Int
	EventTime    uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterQueueMatched is a free log retrieval operation binding the contract event 0x9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a054.
//
// Solidity: event QueueMatched(battleGroup1 uint256, battleGroup2 uint256, eventTime uint64)
func (_BattleQueue *BattleQueueFilterer) FilterQueueMatched(opts *bind.FilterOpts) (*BattleQueueQueueMatchedIterator, error) {

	logs, sub, err := _BattleQueue.contract.FilterLogs(opts, "QueueMatched")
	if err != nil {
		return nil, err
	}
	return &BattleQueueQueueMatchedIterator{contract: _BattleQueue.contract, event: "QueueMatched", logs: logs, sub: sub}, nil
}

// WatchQueueMatched is a free log subscription operation binding the contract event 0x9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a054.
//
// Solidity: event QueueMatched(battleGroup1 uint256, battleGroup2 uint256, eventTime uint64)
func (_BattleQueue *BattleQueueFilterer) WatchQueueMatched(opts *bind.WatchOpts, sink chan<- *BattleQueueQueueMatched) (event.Subscription, error) {

	logs, sub, err := _BattleQueue.contract.WatchLogs(opts, "QueueMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleQueueQueueMatched)
				if err := _BattleQueue.contract.UnpackLog(event, "QueueMatched", log); err != nil {
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
const BattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoCardsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countBattles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op1\",\"type\":\"uint256\"},{\"name\":\"_op2\",\"type\":\"uint256\"}],\"name\":\"createBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerBattleGroup\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loserBattleGroup\",\"type\":\"uint256\"}],\"name\":\"BattleResult\",\"type\":\"event\"}]"

// BattlesBin is the compiled bytecode used for deploying new contracts.
const BattlesBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561059e806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461003957806363c785aa14610068578063df73c49f1461008d57600080fd5b341561004457600080fd5b61004c6100a6565b604051600160a060020a03909116815260200160405180910390f35b341561007357600080fd5b61007b6100b5565b60405190815260200160405180910390f35b341561009857600080fd5b61007b6004356024356100bc565b600154600160a060020a031681565b6000545b90565b60006100c66104c2565b60008060006100d487610231565b15806100e657506100e486610231565b155b156100f45760009450610227565b60a0604051908101604052804267ffffffffffffffff168152602001888152602001878152602001600081526020016000815250935060016000805480600101828161014091906104fc565b600092835260209092208791600502018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151816003015560808201516004909101555003925063ffffffff831683146101ae57600080fd5b6101b784610304565b606086018290526080860181905290925090506101d38461031a565b507f17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b838560600151866080015160405180848152602001838152602001828152602001935050505060405180910390a18294505b5050505092915050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561027b57600080fd5b6102c65a03f1151561028c57600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102e457600080fd5b6102c65a03f115156102f557600080fd5b50505060405180519392505050565b6000808260400151836020015191509150915091565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561036457600080fd5b6102c65a03f1151561037557600080fd5b50505060405180519050600160a060020a031663b42cb37d836020015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103d157600080fd5b6102c65a03f115156103e257600080fd5b50505060405180515050600154600160a060020a031663e77dad5c6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561043457600080fd5b6102c65a03f1151561044557600080fd5b50505060405180519050600160a060020a031663b42cb37d836040015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156104a157600080fd5b6102c65a03f115156104b257600080fd5b5050506040518051905050919050565b60a060405190810160405280600067ffffffffffffffff168152602001600081526020016000815260200160008152602001600081525090565b81548183558181151161052857600502816005028360005260206000209182019101610528919061052d565b505050565b6100b991905b8082111561056e57805467ffffffffffffffff1916815560006001820181905560028201819055600382018190556004820155600501610533565b50905600a165627a7a7230582048d780e88e975e1b17f7be246e6312b7349527c4e1afed76aa57daccff81a6030029`

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

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_Battles *BattlesCaller) CryptoCardsContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Battles.contract.Call(opts, out, "cryptoCardsContract")
	return *ret0, err
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_Battles *BattlesSession) CryptoCardsContract() (common.Address, error) {
	return _Battles.Contract.CryptoCardsContract(&_Battles.CallOpts)
}

// CryptoCardsContract is a free data retrieval call binding the contract method 0x624fe81f.
//
// Solidity: function cryptoCardsContract() constant returns(address)
func (_Battles *BattlesCallerSession) CryptoCardsContract() (common.Address, error) {
	return _Battles.Contract.CryptoCardsContract(&_Battles.CallOpts)
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
const CardBaseBin = `0x6060604052341561000f57600080fd5b6101498061001e6000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461004757806378d0df501461008857600080fd5b341561005257600080fd5b61007673ffffffffffffffffffffffffffffffffffffffff600435166024356100c7565b60405190815260200160405180910390f35b341561009357600080fd5b61009e6004356100f5565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6002602052816000526040600020818154811015156100e257fe5b6000918252602090912001549150829050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058206ae49c1f5d8a8136d6967c21e857fe5a9289f9fe05894dd0dff721b73de41fe40029`

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

// CardBattlesABI is the input ABI used to generate the binding from.
const CardBattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardBattlesBin is the compiled bytecode used for deploying new contracts.
const CardBattlesBin = `0x6060604052341561000f57600080fd5b610ab68061001e6000396000f300606060405236156100cd5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d257806312c2debf1461015c57806318160ddd146101905780632109f89d146101a357806330259720146101cd5780635de038b1146101e55780636352211e14610207578063685b7a5e1461023957806370a082311461024f57806378d0df501461026e5780638462151c1461028457806395d89b41146102f65780639db0225714610309578063a9059cbb14610339575b600080fd5b34156100dd57600080fd5b6100e561035b565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610121578082015183820152602001610109565b50505050905090810190601f16801561014e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016757600080fd5b61017e600160a060020a0360043516602435610392565b60405190815260200160405180910390f35b341561019b57600080fd5b61017e6103c0565b34156101ae57600080fd5b6101b96004356103c7565b604051901515815260200160405180910390f35b34156101d857600080fd5b6101e360043561042a565b005b34156101f057600080fd5b61017e600160a060020a0360043516602435610440565b341561021257600080fd5b61021d600435610478565b604051600160a060020a03909116815260200160405180910390f35b341561024457600080fd5b6101b960043561049c565b341561025a57600080fd5b61017e600160a060020a03600435166104fb565b341561027957600080fd5b61021d600435610516565b341561028f57600080fd5b6102a3600160a060020a0360043516610531565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156102e25780820151838201526020016102ca565b505050509050019250505060405180910390f35b341561030157600080fd5b6100e5610612565b341561031457600080fd5b61031c610649565b60405167ffffffffffffffff909116815260200160405180910390f35b341561034457600080fd5b6101e3600160a060020a036004351660243561064f565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156103ad57fe5b6000918252602090912001549150829050565b6000545b90565b60006103d28261042a565b4267ffffffffffffffff166000838154811015156103ec57fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff16111561042157506000610425565b5060015b919050565b6104326103c0565b811061043d57600080fd5b50565b60008030600160a060020a031684600160a060020a03161415151561046457600080fd5b610470600084866106a9565b949350505050565b600081815260016020526040902054600160a060020a031680151561042557600080fd5b60006104a78261042a565b610e1042016000838154811015156104bb57fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6105396109e7565b60006105436109e7565b6000806000610551876104fb565b94508415156105815760006040518059106105695750595b90808252806020026020018201604052509550610608565b8460405180591061058f5750595b908082528060200260200182016040525093506105aa6103c0565b925060009150600190505b82811161060457600081815260016020526040902054600160a060020a03888116911614156105fc57808483815181106105eb57fe5b602090810290910101526001909101905b6001016105b5565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b600160a060020a038216151561066457600080fd5b30600160a060020a031682600160a060020a03161415151561068557600080fd5b61068f33826108fe565b151561069a57600080fd5b6106a533838361091e565b5050565b60006106b36109f9565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816107489190610a2e565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff8116811461087557600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a16108f56000858361091e565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561099257600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610a5a57600302816003028360005260206000209182019101610a5a9190610a5f565b505050565b6103c491905b80821115610a86576000808255600182018190556002820155600301610a65565b50905600a165627a7a72305820da95b6aea3f05f16ef135bcc1dfe231569f67831eb34276532ece6a2ecbcc3790029`

// DeployCardBattles deploys a new Ethereum contract, binding an instance of CardBattles to it.
func DeployCardBattles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CardBattles, error) {
	parsed, err := abi.JSON(strings.NewReader(CardBattlesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CardBattlesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CardBattles{CardBattlesCaller: CardBattlesCaller{contract: contract}, CardBattlesTransactor: CardBattlesTransactor{contract: contract}, CardBattlesFilterer: CardBattlesFilterer{contract: contract}}, nil
}

// CardBattles is an auto generated Go binding around an Ethereum contract.
type CardBattles struct {
	CardBattlesCaller     // Read-only binding to the contract
	CardBattlesTransactor // Write-only binding to the contract
	CardBattlesFilterer   // Log filterer for contract events
}

// CardBattlesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CardBattlesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBattlesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CardBattlesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBattlesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CardBattlesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CardBattlesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CardBattlesSession struct {
	Contract     *CardBattles      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CardBattlesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CardBattlesCallerSession struct {
	Contract *CardBattlesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CardBattlesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CardBattlesTransactorSession struct {
	Contract     *CardBattlesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CardBattlesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CardBattlesRaw struct {
	Contract *CardBattles // Generic contract binding to access the raw methods on
}

// CardBattlesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CardBattlesCallerRaw struct {
	Contract *CardBattlesCaller // Generic read-only contract binding to access the raw methods on
}

// CardBattlesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CardBattlesTransactorRaw struct {
	Contract *CardBattlesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCardBattles creates a new instance of CardBattles, bound to a specific deployed contract.
func NewCardBattles(address common.Address, backend bind.ContractBackend) (*CardBattles, error) {
	contract, err := bindCardBattles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CardBattles{CardBattlesCaller: CardBattlesCaller{contract: contract}, CardBattlesTransactor: CardBattlesTransactor{contract: contract}, CardBattlesFilterer: CardBattlesFilterer{contract: contract}}, nil
}

// NewCardBattlesCaller creates a new read-only instance of CardBattles, bound to a specific deployed contract.
func NewCardBattlesCaller(address common.Address, caller bind.ContractCaller) (*CardBattlesCaller, error) {
	contract, err := bindCardBattles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CardBattlesCaller{contract: contract}, nil
}

// NewCardBattlesTransactor creates a new write-only instance of CardBattles, bound to a specific deployed contract.
func NewCardBattlesTransactor(address common.Address, transactor bind.ContractTransactor) (*CardBattlesTransactor, error) {
	contract, err := bindCardBattles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CardBattlesTransactor{contract: contract}, nil
}

// NewCardBattlesFilterer creates a new log filterer instance of CardBattles, bound to a specific deployed contract.
func NewCardBattlesFilterer(address common.Address, filterer bind.ContractFilterer) (*CardBattlesFilterer, error) {
	contract, err := bindCardBattles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CardBattlesFilterer{contract: contract}, nil
}

// bindCardBattles binds a generic wrapper to an already deployed contract.
func bindCardBattles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CardBattlesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardBattles *CardBattlesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardBattles.Contract.CardBattlesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardBattles *CardBattlesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBattles.Contract.CardBattlesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardBattles *CardBattlesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardBattles.Contract.CardBattlesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CardBattles *CardBattlesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CardBattles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CardBattles *CardBattlesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBattles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CardBattles *CardBattlesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CardBattles.Contract.contract.Transact(opts, method, params...)
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardBattles *CardBattlesCaller) BATTLECOOLDOWNTIME(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "BATTLE_COOLDOWN_TIME")
	return *ret0, err
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardBattles *CardBattlesSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CardBattles.Contract.BATTLECOOLDOWNTIME(&_CardBattles.CallOpts)
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardBattles *CardBattlesCallerSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CardBattles.Contract.BATTLECOOLDOWNTIME(&_CardBattles.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardBattles *CardBattlesCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardBattles *CardBattlesSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardBattles.Contract.BalanceOf(&_CardBattles.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(count uint256)
func (_CardBattles *CardBattlesCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _CardBattles.Contract.BalanceOf(&_CardBattles.CallOpts, _owner)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBattles *CardBattlesCaller) CardIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "cardIndexToOwner", arg0)
	return *ret0, err
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBattles *CardBattlesSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardBattles.Contract.CardIndexToOwner(&_CardBattles.CallOpts, arg0)
}

// CardIndexToOwner is a free data retrieval call binding the contract method 0x78d0df50.
//
// Solidity: function cardIndexToOwner( uint256) constant returns(address)
func (_CardBattles *CardBattlesCallerSession) CardIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _CardBattles.Contract.CardIndexToOwner(&_CardBattles.CallOpts, arg0)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBattles *CardBattlesCaller) CardsHeldByOwner(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "cardsHeldByOwner", arg0, arg1)
	return *ret0, err
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBattles *CardBattlesSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardBattles.Contract.CardsHeldByOwner(&_CardBattles.CallOpts, arg0, arg1)
}

// CardsHeldByOwner is a free data retrieval call binding the contract method 0x12c2debf.
//
// Solidity: function cardsHeldByOwner( address,  uint256) constant returns(uint256)
func (_CardBattles *CardBattlesCallerSession) CardsHeldByOwner(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _CardBattles.Contract.CardsHeldByOwner(&_CardBattles.CallOpts, arg0, arg1)
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardBattles *CardBattlesCaller) IsReadyForBattle(opts *bind.CallOpts, _cardID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "isReadyForBattle", _cardID)
	return *ret0, err
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardBattles *CardBattlesSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CardBattles.Contract.IsReadyForBattle(&_CardBattles.CallOpts, _cardID)
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardBattles *CardBattlesCallerSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CardBattles.Contract.IsReadyForBattle(&_CardBattles.CallOpts, _cardID)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardBattles *CardBattlesCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardBattles *CardBattlesSession) Name() (string, error) {
	return _CardBattles.Contract.Name(&_CardBattles.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_CardBattles *CardBattlesCallerSession) Name() (string, error) {
	return _CardBattles.Contract.Name(&_CardBattles.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardBattles *CardBattlesCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardBattles *CardBattlesSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardBattles.Contract.OwnerOf(&_CardBattles.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(owner address)
func (_CardBattles *CardBattlesCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _CardBattles.Contract.OwnerOf(&_CardBattles.CallOpts, _tokenId)
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardBattles *CardBattlesCaller) RequireCardExists(opts *bind.CallOpts, cardID *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _CardBattles.contract.Call(opts, out, "requireCardExists", cardID)
	return err
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardBattles *CardBattlesSession) RequireCardExists(cardID *big.Int) error {
	return _CardBattles.Contract.RequireCardExists(&_CardBattles.CallOpts, cardID)
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardBattles *CardBattlesCallerSession) RequireCardExists(cardID *big.Int) error {
	return _CardBattles.Contract.RequireCardExists(&_CardBattles.CallOpts, cardID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardBattles *CardBattlesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardBattles *CardBattlesSession) Symbol() (string, error) {
	return _CardBattles.Contract.Symbol(&_CardBattles.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_CardBattles *CardBattlesCallerSession) Symbol() (string, error) {
	return _CardBattles.Contract.Symbol(&_CardBattles.CallOpts)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardBattles *CardBattlesCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardBattles *CardBattlesSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardBattles.Contract.TokensOfOwner(&_CardBattles.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(_owner address) constant returns(ownerTokens uint256[])
func (_CardBattles *CardBattlesCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _CardBattles.Contract.TokensOfOwner(&_CardBattles.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardBattles *CardBattlesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardBattles *CardBattlesSession) TotalSupply() (*big.Int, error) {
	return _CardBattles.Contract.TotalSupply(&_CardBattles.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CardBattles *CardBattlesCallerSession) TotalSupply() (*big.Int, error) {
	return _CardBattles.Contract.TotalSupply(&_CardBattles.CallOpts)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardBattles *CardBattlesTransactor) CreateCard(opts *bind.TransactOpts, _owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "createCard", _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardBattles *CardBattlesSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.CreateCard(&_CardBattles.TransactOpts, _owner, _attributes)
}

// CreateCard is a paid mutator transaction binding the contract method 0x5de038b1.
//
// Solidity: function createCard(_owner address, _attributes uint256) returns(uint256)
func (_CardBattles *CardBattlesTransactorSession) CreateCard(_owner common.Address, _attributes *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.CreateCard(&_CardBattles.TransactOpts, _owner, _attributes)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardBattles *CardBattlesTransactor) SetOnBattleCooldown(opts *bind.TransactOpts, _cardID *big.Int) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "setOnBattleCooldown", _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardBattles *CardBattlesSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.SetOnBattleCooldown(&_CardBattles.TransactOpts, _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardBattles *CardBattlesTransactorSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.SetOnBattleCooldown(&_CardBattles.TransactOpts, _cardID)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardBattles *CardBattlesTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "transfer", _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardBattles *CardBattlesSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.Transfer(&_CardBattles.TransactOpts, _to, _cardId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _cardId uint256) returns()
func (_CardBattles *CardBattlesTransactorSession) Transfer(_to common.Address, _cardId *big.Int) (*types.Transaction, error) {
	return _CardBattles.Contract.Transfer(&_CardBattles.TransactOpts, _to, _cardId)
}

// CardBattlesNewCardIterator is returned from FilterNewCard and is used to iterate over the raw logs and unpacked data for NewCard events raised by the CardBattles contract.
type CardBattlesNewCardIterator struct {
	Event *CardBattlesNewCard // Event containing the contract specifics and raw log

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
func (it *CardBattlesNewCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBattlesNewCard)
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
		it.Event = new(CardBattlesNewCard)
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
func (it *CardBattlesNewCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBattlesNewCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBattlesNewCard represents a NewCard event raised by the CardBattles contract.
type CardBattlesNewCard struct {
	Owner            common.Address
	CardID           *big.Int
	CreationBattleID *big.Int
	Attributes       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewCard is a free log retrieval operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardBattles *CardBattlesFilterer) FilterNewCard(opts *bind.FilterOpts) (*CardBattlesNewCardIterator, error) {

	logs, sub, err := _CardBattles.contract.FilterLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return &CardBattlesNewCardIterator{contract: _CardBattles.contract, event: "NewCard", logs: logs, sub: sub}, nil
}

// WatchNewCard is a free log subscription operation binding the contract event 0xc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4.
//
// Solidity: event NewCard(owner address, cardID uint256, creationBattleID uint128, attributes uint256)
func (_CardBattles *CardBattlesFilterer) WatchNewCard(opts *bind.WatchOpts, sink chan<- *CardBattlesNewCard) (event.Subscription, error) {

	logs, sub, err := _CardBattles.contract.WatchLogs(opts, "NewCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBattlesNewCard)
				if err := _CardBattles.contract.UnpackLog(event, "NewCard", log); err != nil {
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

// CardBattlesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CardBattles contract.
type CardBattlesTransferIterator struct {
	Event *CardBattlesTransfer // Event containing the contract specifics and raw log

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
func (it *CardBattlesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBattlesTransfer)
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
		it.Event = new(CardBattlesTransfer)
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
func (it *CardBattlesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBattlesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBattlesTransfer represents a Transfer event raised by the CardBattles contract.
type CardBattlesTransfer struct {
	From    common.Address
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardBattles *CardBattlesFilterer) FilterTransfer(opts *bind.FilterOpts) (*CardBattlesTransferIterator, error) {

	logs, sub, err := _CardBattles.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &CardBattlesTransferIterator{contract: _CardBattles.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from address, to address, tokenID uint256)
func (_CardBattles *CardBattlesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CardBattlesTransfer) (event.Subscription, error) {

	logs, sub, err := _CardBattles.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBattlesTransfer)
				if err := _CardBattles.contract.UnpackLog(event, "Transfer", log); err != nil {
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
const CardMintingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"newCardID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardMintingBin is the compiled bytecode used for deploying new contracts.
const CardMintingBin = `0x6060604052341561000f57600080fd5b610ae98061001e6000396000f300606060405236156100d85763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100dd57806312c2debf1461016757806318160ddd1461019b5780632109f89d146101ae57806330259720146101d857806345812958146101f05780635de038b1146102035780636352211e14610225578063685b7a5e1461025757806370a082311461026d57806378d0df501461028c5780638462151c146102a257806395d89b41146103145780639db0225714610327578063a9059cbb14610357575b600080fd5b34156100e857600080fd5b6100f0610379565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561012c578082015183820152602001610114565b50505050905090810190601f1680156101595780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561017257600080fd5b610189600160a060020a03600435166024356103b0565b60405190815260200160405180910390f35b34156101a657600080fd5b6101896103de565b34156101b957600080fd5b6101c46004356103e5565b604051901515815260200160405180910390f35b34156101e357600080fd5b6101ee600435610448565b005b34156101fb57600080fd5b61018961045e565b341561020e57600080fd5b610189600160a060020a0360043516602435610473565b341561023057600080fd5b61023b6004356104ab565b604051600160a060020a03909116815260200160405180910390f35b341561026257600080fd5b6101c46004356104cf565b341561027857600080fd5b610189600160a060020a036004351661052e565b341561029757600080fd5b61023b600435610549565b34156102ad57600080fd5b6102c1600160a060020a0360043516610564565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103005780820151838201526020016102e8565b505050509050019250505060405180910390f35b341561031f57600080fd5b6100f0610645565b341561033257600080fd5b61033a61067c565b60405167ffffffffffffffff909116815260200160405180910390f35b341561036257600080fd5b6101ee600160a060020a0360043516602435610682565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156103cb57fe5b6000918252602090912001549150829050565b6000545b90565b60006103f082610448565b4267ffffffffffffffff1660008381548110151561040a57fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff16111561043f57506000610443565b5060015b919050565b6104506103de565b811061045b57600080fd5b50565b600061046e60006201e240610473565b905090565b60008030600160a060020a031684600160a060020a03161415151561049757600080fd5b6104a3600084866106dc565b949350505050565b600081815260016020526040902054600160a060020a031680151561044357600080fd5b60006104da82610448565b610e1042016000838154811015156104ee57fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b61056c610a1a565b6000610576610a1a565b60008060006105848761052e565b94508415156105b457600060405180591061059c5750595b9080825280602002602001820160405250955061063b565b846040518059106105c25750595b908082528060200260200182016040525093506105dd6103de565b925060009150600190505b82811161063757600081815260016020526040902054600160a060020a038881169116141561062f578084838151811061061e57fe5b602090810290910101526001909101905b6001016105e8565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b600160a060020a038216151561069757600080fd5b30600160a060020a031682600160a060020a0316141515156106b857600080fd5b6106c23382610931565b15156106cd57600080fd5b6106d8338383610951565b5050565b60006106e6610a2c565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200186815250915060016000805480600101828161077b9190610a61565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146108a857600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161092860008583610951565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff191690911790558316156109c557600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610a8d57600302816003028360005260206000209182019101610a8d9190610a92565b505050565b6103e291905b80821115610ab9576000808255600182018190556002820155600301610a98565b50905600a165627a7a723058208a3b2d3d9be1f359e7b72e11feda3a2cc8e50426e75acc3de19289ec9b67842a0029`

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

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardMinting *CardMintingCaller) BATTLECOOLDOWNTIME(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "BATTLE_COOLDOWN_TIME")
	return *ret0, err
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardMinting *CardMintingSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CardMinting.Contract.BATTLECOOLDOWNTIME(&_CardMinting.CallOpts)
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CardMinting *CardMintingCallerSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CardMinting.Contract.BATTLECOOLDOWNTIME(&_CardMinting.CallOpts)
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

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardMinting *CardMintingCaller) IsReadyForBattle(opts *bind.CallOpts, _cardID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "isReadyForBattle", _cardID)
	return *ret0, err
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardMinting *CardMintingSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CardMinting.Contract.IsReadyForBattle(&_CardMinting.CallOpts, _cardID)
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CardMinting *CardMintingCallerSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CardMinting.Contract.IsReadyForBattle(&_CardMinting.CallOpts, _cardID)
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

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardMinting *CardMintingCaller) RequireCardExists(opts *bind.CallOpts, cardID *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _CardMinting.contract.Call(opts, out, "requireCardExists", cardID)
	return err
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardMinting *CardMintingSession) RequireCardExists(cardID *big.Int) error {
	return _CardMinting.Contract.RequireCardExists(&_CardMinting.CallOpts, cardID)
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardMinting *CardMintingCallerSession) RequireCardExists(cardID *big.Int) error {
	return _CardMinting.Contract.RequireCardExists(&_CardMinting.CallOpts, cardID)
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
// Solidity: function createGen0Card() returns(newCardID uint256)
func (_CardMinting *CardMintingTransactor) CreateGen0Card(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "createGen0Card")
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(newCardID uint256)
func (_CardMinting *CardMintingSession) CreateGen0Card() (*types.Transaction, error) {
	return _CardMinting.Contract.CreateGen0Card(&_CardMinting.TransactOpts)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(newCardID uint256)
func (_CardMinting *CardMintingTransactorSession) CreateGen0Card() (*types.Transaction, error) {
	return _CardMinting.Contract.CreateGen0Card(&_CardMinting.TransactOpts)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardMinting *CardMintingTransactor) SetOnBattleCooldown(opts *bind.TransactOpts, _cardID *big.Int) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "setOnBattleCooldown", _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardMinting *CardMintingSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.SetOnBattleCooldown(&_CardMinting.TransactOpts, _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CardMinting *CardMintingTransactorSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CardMinting.Contract.SetOnBattleCooldown(&_CardMinting.TransactOpts, _cardID)
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
const CardOwnershipABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardOwnershipBin is the compiled bytecode used for deploying new contracts.
const CardOwnershipBin = `0x6060604052341561000f57600080fd5b6109628061001e6000396000f300606060405236156100ac5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b157806312c2debf1461013b57806318160ddd1461016f57806330259720146101825780635de038b11461019a5780636352211e146101bc57806370a08231146101ee57806378d0df501461020d5780638462151c1461022357806395d89b4114610295578063a9059cbb146102a8575b600080fd5b34156100bc57600080fd5b6100c46102ca565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101005780820151838201526020016100e8565b50505050905090810190601f16801561012d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014657600080fd5b61015d600160a060020a0360043516602435610301565b60405190815260200160405180910390f35b341561017a57600080fd5b61015d61032f565b341561018d57600080fd5b610198600435610336565b005b34156101a557600080fd5b61015d600160a060020a036004351660243561034c565b34156101c757600080fd5b6101d2600435610384565b604051600160a060020a03909116815260200160405180910390f35b34156101f957600080fd5b61015d600160a060020a03600435166103ad565b341561021857600080fd5b6101d26004356103c8565b341561022e57600080fd5b610242600160a060020a03600435166103e3565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610281578082015183820152602001610269565b505050509050019250505060405180910390f35b34156102a057600080fd5b6100c46104c4565b34156102b357600080fd5b610198600160a060020a03600435166024356104fb565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60026020528160005260406000208181548110151561031c57fe5b6000918252602090912001549150829050565b6000545b90565b61033e61032f565b811061034957600080fd5b50565b60008030600160a060020a031684600160a060020a03161415151561037057600080fd5b61037c60008486610555565b949350505050565b600081815260016020526040902054600160a060020a03168015156103a857600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103eb610893565b60006103f5610893565b6000806000610403876103ad565b945084151561043357600060405180591061041b5750595b908082528060200260200182016040525095506104ba565b846040518059106104415750595b9080825280602002602001820160405250935061045c61032f565b925060009150600190505b8281116104b657600081815260016020526040902054600160a060020a03888116911614156104ae578084838151811061049d57fe5b602090810290910101526001909101905b600101610467565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a038216151561051057600080fd5b30600160a060020a031682600160a060020a03161415151561053157600080fd5b61053b33826107aa565b151561054657600080fd5b6105513383836107ca565b5050565b600061055f6108a5565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105f491906108da565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff8116811461072157600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a16107a1600085836107ca565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561083e57600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b81548183558181151161090657600302816003028360005260206000209182019101610906919061090b565b505050565b61033391905b80821115610932576000808255600182018190556002820155600301610911565b50905600a165627a7a72305820abf2cac6472d2118ffe712aeac9e7b4408f5e3a67a9a92e66f513dac4bc5d9ae0029`

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

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardOwnership *CardOwnershipCaller) RequireCardExists(opts *bind.CallOpts, cardID *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _CardOwnership.contract.Call(opts, out, "requireCardExists", cardID)
	return err
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardOwnership *CardOwnershipSession) RequireCardExists(cardID *big.Int) error {
	return _CardOwnership.Contract.RequireCardExists(&_CardOwnership.CallOpts, cardID)
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CardOwnership *CardOwnershipCallerSession) RequireCardExists(cardID *big.Int) error {
	return _CardOwnership.Contract.RequireCardExists(&_CardOwnership.CallOpts, cardID)
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
const CryptoCardsCoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"newCardID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleQueueContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleGroupContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CryptoCardsCoreBin is the compiled bytecode used for deploying new contracts.
const CryptoCardsCoreBin = `0x6060604052341561000f57600080fd5b60048054600160a060020a03191633600160a060020a03161790556100326100e9565b604051809103906000f080151561004857600080fd5b60058054600160a060020a031916600160a060020a03929092169190911790556100706100f9565b604051809103906000f080151561008657600080fd5b60068054600160a060020a031916600160a060020a03929092169190911790556100ae610109565b604051809103906000f08015156100c457600080fd5b60078054600160a060020a031916600160a060020a0392909216919091179055610119565b6040516105d780610d6d83390190565b6040516107388061134483390190565b60405161038880611a7c83390190565b610c45806101286000396000f3006060604052361561010f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461012c57806312c2debf146101b657806318160ddd146101ea5780632109f89d146101fd5780633025972014610227578063458129581461023d5780635de038b1146102505780636352211e14610272578063685b7a5e146102a45780636af04a57146102ba57806370a08231146102cd57806371587988146102ec57806378d0df501461030b5780638462151c1461032157806395d89b41146103935780639db02257146103a6578063a9059cbb146103d6578063d78d16b6146103f8578063e01196ef1461040b578063e77dad5c1461041e575b60055433600160a060020a0390811691161461012a57600080fd5b005b341561013757600080fd5b61013f610431565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561017b578082015183820152602001610163565b50505050905090810190601f1680156101a85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101c157600080fd5b6101d8600160a060020a0360043516602435610468565b60405190815260200160405180910390f35b34156101f557600080fd5b6101d8610496565b341561020857600080fd5b61021360043561049d565b604051901515815260200160405180910390f35b341561023257600080fd5b61012a600435610500565b341561024857600080fd5b6101d8610516565b341561025b57600080fd5b6101d8600160a060020a036004351660243561052b565b341561027d57600080fd5b610288600435610563565b604051600160a060020a03909116815260200160405180910390f35b34156102af57600080fd5b610213600435610587565b34156102c557600080fd5b6102886105e6565b34156102d857600080fd5b6101d8600160a060020a03600435166105f5565b34156102f757600080fd5b61012a600160a060020a0360043516610610565b341561031657600080fd5b610288600435610678565b341561032c57600080fd5b610340600160a060020a0360043516610693565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561037f578082015183820152602001610367565b505050509050019250505060405180910390f35b341561039e57600080fd5b61013f610774565b34156103b157600080fd5b6103b96107ab565b60405167ffffffffffffffff909116815260200160405180910390f35b34156103e157600080fd5b61012a600160a060020a03600435166024356107b1565b341561040357600080fd5b61028861080b565b341561041657600080fd5b61028861081a565b341561042957600080fd5b610288610829565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60026020528160005260406000208181548110151561048357fe5b6000918252602090912001549150829050565b6000545b90565b60006104a882610500565b4267ffffffffffffffff166000838154811015156104c257fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff1611156104f7575060006104fb565b5060015b919050565b610508610496565b811061051357600080fd5b50565b600061052660006201e24061052b565b905090565b60008030600160a060020a031684600160a060020a03161415151561054f57600080fd5b61055b60008486610838565b949350505050565b600081815260016020526040902054600160a060020a03168015156104fb57600080fd5b600061059282610500565b610e1042016000838154811015156105a657fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600854600160a060020a031681565b600160a060020a031660009081526003602052604090205490565b6008805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600160205260009081526040902054600160a060020a031681565b61069b610b76565b60006106a5610b76565b60008060006106b3876105f5565b94508415156106e35760006040518059106106cb5750595b9080825280602002602001820160405250955061076a565b846040518059106106f15750595b9080825280602002602001820160405250935061070c610496565b925060009150600190505b82811161076657600081815260016020526040902054600160a060020a038881169116141561075e578084838151811061074d57fe5b602090810290910101526001909101905b600101610717565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b600160a060020a03821615156107c657600080fd5b30600160a060020a031682600160a060020a0316141515156107e757600080fd5b6107f13382610a8d565b15156107fc57600080fd5b610807338383610aad565b5050565b600754600160a060020a031690565b600554600160a060020a031690565b600654600160a060020a031690565b6000610842610b88565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816108d79190610bbd565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff81168114610a0457600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a1610a8460008583610aad565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055831615610b2157600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610be957600302816003028360005260206000209182019101610be99190610bee565b505050565b61049a91905b80821115610c15576000808255600182018190556002820155600301610bf4565b50905600a165627a7a7230582081ba93efd0211dbe529e0010df650cdd0a593d9502a32c5e72d64551368231ff00296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561059e806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461003957806363c785aa14610068578063df73c49f1461008d57600080fd5b341561004457600080fd5b61004c6100a6565b604051600160a060020a03909116815260200160405180910390f35b341561007357600080fd5b61007b6100b5565b60405190815260200160405180910390f35b341561009857600080fd5b61007b6004356024356100bc565b600154600160a060020a031681565b6000545b90565b60006100c66104c2565b60008060006100d487610231565b15806100e657506100e486610231565b155b156100f45760009450610227565b60a0604051908101604052804267ffffffffffffffff168152602001888152602001878152602001600081526020016000815250935060016000805480600101828161014091906104fc565b600092835260209092208791600502018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151816003015560808201516004909101555003925063ffffffff831683146101ae57600080fd5b6101b784610304565b606086018290526080860181905290925090506101d38461031a565b507f17a28617c80116086701f1712a79248bbe6e8e30c63a1e4735995ec3b0170a8b838560600151866080015160405180848152602001838152602001828152602001935050505060405180910390a18294505b5050505092915050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561027b57600080fd5b6102c65a03f1151561028c57600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102e457600080fd5b6102c65a03f115156102f557600080fd5b50505060405180519392505050565b6000808260400151836020015191509150915091565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561036457600080fd5b6102c65a03f1151561037557600080fd5b50505060405180519050600160a060020a031663b42cb37d836020015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103d157600080fd5b6102c65a03f115156103e257600080fd5b50505060405180515050600154600160a060020a031663e77dad5c6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561043457600080fd5b6102c65a03f1151561044557600080fd5b50505060405180519050600160a060020a031663b42cb37d836040015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156104a157600080fd5b6102c65a03f115156104b257600080fd5b5050506040518051905050919050565b60a060405190810160405280600067ffffffffffffffff168152602001600081526020016000815260200160008152602001600081525090565b81548183558181151161052857600502816005028360005260206000209182019101610528919061052d565b505050565b6100b991905b8082111561056e57805467ffffffffffffffff1916815560006001820181905560028201819055600382018190556004820155600501610533565b50905600a165627a7a7230582048d780e88e975e1b17f7be246e6312b7349527c4e1afed76aa57daccff81a60300296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a03161790556106ff806100396000396000f3006060604052361561005c5763ffffffff60e060020a600035041663624fe81f81146100615780637bf13d8214610090578063b42cb37d146100b5578063bfeea708146100df578063ecf9036b146100f2578063fa74efc114610108575b600080fd5b341561006c57600080fd5b610074610129565b604051600160a060020a03909116815260200160405180910390f35b341561009b57600080fd5b6100a3610138565b60405190815260200160405180910390f35b34156100c057600080fd5b6100cb60043561013f565b604051901515815260200160405180910390f35b34156100ea57600080fd5b6100a361029a565b34156100fd57600080fd5b6100cb60043561029f565b341561011357600080fd5b6100a3600160a060020a036004351660246103fd565b600154600160a060020a031681565b6000545b90565b600061014961058e565b6000610153610138565b841061015e57600080fd5b600080548590811061016c57fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116101d25750505091909252509193506000925050505b816040015150600581101561028e57600154600160a060020a031663685b7a5e6040840151836005811061022757fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561026b57600080fd5b6102c65a03f1151561027c57600080fd5b505050604051805150506001016101f7565b600192505b5050919050565b600581565b60006102a961058e565b60006102b3610138565b84106102be57600080fd5b60008054859081106102cc57fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116103325750505091909252509193506000925050505b816040015150600581101561028e57600154600160a060020a0316632109f89d6040840151836005811061038757fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103cb57600080fd5b6102c65a03f115156103dc57600080fd5b5050506040518051905015156103f55760009250610293565b600101610357565b600061040761058e565b60006060604051908101604052804267ffffffffffffffff16815260200186600160a060020a03168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161046e83826105b4565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518154600160a060020a039190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516104f790600183019060056105e5565b50505003905063ffffffff8116811461050f57600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d85828460400151604051600160a060020a038416815260208101839052604081018260a080838360005b8381101561057257808201518382015260200161055a565b50505050905001935050505060405180910390a1949350505050565b60e060405190810160409081526000808352602083015281016105af610623565b905290565b8154818355818115116105e0576006028160060283600052602060002091820191016105e0919061064a565b505050565b8260058101928215610613579160200282015b828111156106135782518255916020019190600101906105f8565b5061061f929150610696565b5090565b60a06040519081016040526005815b60008152602001906001900390816106325790505090565b61013c91905b8082111561061f5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061068d60018301826106b0565b50600601610650565b61013c91905b8082111561061f576000815560010161069c565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a723058206cc3ff584860016f2e985e2ecda1e0fb1a0ea6294c52e91101e1787baddcc70900296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561034f806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461002e578063969494201461005d57600080fd5b341561003957600080fd5b610041610087565b604051600160a060020a03909116815260200160405180910390f35b341561006857600080fd5b610073600435610096565b604051901515815260200160405180910390f35b600154600160a060020a031681565b60006100a182610176565b15156100af57506000610171565b7f4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7824260405191825267ffffffffffffffff1660208201526040908101905180910390a16000541515610109575060008190556001610171565b7f9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a0548260005442604051928352602083019190915267ffffffffffffffff166040808301919091526060909101905180910390a161016882600054610249565b50506000805560015b919050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156101c057600080fd5b6102c65a03f115156101d157600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561022957600080fd5b6102c65a03f1151561023a57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e01196ef82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029357600080fd5b6102c65a03f115156102a457600080fd5b50505060405180519050600160a060020a031663df73c49f848460006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561030257600080fd5b6102c65a03f1151561031357600080fd5b50505060405180519493505050505600a165627a7a7230582091e420400377b31861f8fbb146d0354c712fed3dfb906e6a993c0dff417dce340029`

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

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CryptoCardsCore *CryptoCardsCoreCaller) BATTLECOOLDOWNTIME(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "BATTLE_COOLDOWN_TIME")
	return *ret0, err
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CryptoCardsCore *CryptoCardsCoreSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CryptoCardsCore.Contract.BATTLECOOLDOWNTIME(&_CryptoCardsCore.CallOpts)
}

// BATTLECOOLDOWNTIME is a free data retrieval call binding the contract method 0x9db02257.
//
// Solidity: function BATTLE_COOLDOWN_TIME() constant returns(uint64)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) BATTLECOOLDOWNTIME() (uint64, error) {
	return _CryptoCardsCore.Contract.BATTLECOOLDOWNTIME(&_CryptoCardsCore.CallOpts)
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

// BattleQueueContract is a free data retrieval call binding the contract method 0xd78d16b6.
//
// Solidity: function BattleQueueContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCaller) BattleQueueContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "BattleQueueContract")
	return *ret0, err
}

// BattleQueueContract is a free data retrieval call binding the contract method 0xd78d16b6.
//
// Solidity: function BattleQueueContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreSession) BattleQueueContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleQueueContract(&_CryptoCardsCore.CallOpts)
}

// BattleQueueContract is a free data retrieval call binding the contract method 0xd78d16b6.
//
// Solidity: function BattleQueueContract() constant returns(address)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) BattleQueueContract() (common.Address, error) {
	return _CryptoCardsCore.Contract.BattleQueueContract(&_CryptoCardsCore.CallOpts)
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

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreCaller) IsReadyForBattle(opts *bind.CallOpts, _cardID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "isReadyForBattle", _cardID)
	return *ret0, err
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CryptoCardsCore.Contract.IsReadyForBattle(&_CryptoCardsCore.CallOpts, _cardID)
}

// IsReadyForBattle is a free data retrieval call binding the contract method 0x2109f89d.
//
// Solidity: function isReadyForBattle(_cardID uint256) constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) IsReadyForBattle(_cardID *big.Int) (bool, error) {
	return _CryptoCardsCore.Contract.IsReadyForBattle(&_CryptoCardsCore.CallOpts, _cardID)
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

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CryptoCardsCore *CryptoCardsCoreCaller) RequireCardExists(opts *bind.CallOpts, cardID *big.Int) error {
	var ()
	out := &[]interface{}{}
	err := _CryptoCardsCore.contract.Call(opts, out, "requireCardExists", cardID)
	return err
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) RequireCardExists(cardID *big.Int) error {
	return _CryptoCardsCore.Contract.RequireCardExists(&_CryptoCardsCore.CallOpts, cardID)
}

// RequireCardExists is a free data retrieval call binding the contract method 0x30259720.
//
// Solidity: function requireCardExists(cardID uint256) constant returns()
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) RequireCardExists(cardID *big.Int) error {
	return _CryptoCardsCore.Contract.RequireCardExists(&_CryptoCardsCore.CallOpts, cardID)
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
// Solidity: function createGen0Card() returns(newCardID uint256)
func (_CryptoCardsCore *CryptoCardsCoreTransactor) CreateGen0Card(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "createGen0Card")
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(newCardID uint256)
func (_CryptoCardsCore *CryptoCardsCoreSession) CreateGen0Card() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.CreateGen0Card(&_CryptoCardsCore.TransactOpts)
}

// CreateGen0Card is a paid mutator transaction binding the contract method 0x45812958.
//
// Solidity: function createGen0Card() returns(newCardID uint256)
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

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreTransactor) SetOnBattleCooldown(opts *bind.TransactOpts, _cardID *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "setOnBattleCooldown", _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetOnBattleCooldown(&_CryptoCardsCore.TransactOpts, _cardID)
}

// SetOnBattleCooldown is a paid mutator transaction binding the contract method 0x685b7a5e.
//
// Solidity: function setOnBattleCooldown(_cardID uint256) returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) SetOnBattleCooldown(_cardID *big.Int) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetOnBattleCooldown(&_CryptoCardsCore.TransactOpts, _cardID)
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
