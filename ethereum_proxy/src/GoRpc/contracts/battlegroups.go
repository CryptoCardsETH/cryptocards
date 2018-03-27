// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"
)

// BattleGroupsABI is the input ABI used to generate the binding from.
const BattleGroupsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"countBattleGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_CARDS_PER_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_cards\",\"type\":\"uint256[5]\"}],\"name\":\"createBattleGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"battleGroupID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cards\",\"type\":\"uint256[5]\"}],\"name\":\"NewBattleGroup\",\"type\":\"event\"}]"

// BattleGroupsBin is the compiled bytecode used for deploying new contracts.
const BattleGroupsBin = `0x6060604052341561000f57600080fd5b6103ed8061001e6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416637bf13d828114610052578063bfeea70814610077578063fa74efc11461008a57600080fd5b341561005d57600080fd5b6100656100b8565b60405190815260200160405180910390f35b341561008257600080fd5b6100656100bf565b341561009557600080fd5b61006573ffffffffffffffffffffffffffffffffffffffff6004351660246100c4565b6000545b90565b600581565b60006100ce61027c565b60006060604051908101604052804267ffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001856005806020026040519081016040529190828260a0808284375050509190925250506000805491935060019180830161014283826102a2565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815473ffffffffffffffffffffffffffffffffffffffff9190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516101d890600183019060056102d3565b50505003905063ffffffff811681146101f057600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d8582846040015160405173ffffffffffffffffffffffffffffffffffffffff8416815260208101839052604081018260a080838360005b83811015610260578082015183820152602001610248565b50505050905001935050505060405180910390a1949350505050565b60e0604051908101604090815260008083526020830152810161029d610311565b905290565b8154818355818115116102ce576006028160060283600052602060002091820191016102ce9190610338565b505050565b8260058101928215610301579160200282015b828111156103015782518255916020019190600101906102e6565b5061030d929150610384565b5090565b60a06040519081016040526005815b60008152602001906001900390816103205790505090565b6100bc91905b8082111561030d5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061037b600183018261039e565b5060060161033e565b6100bc91905b8082111561030d576000815560010161038a565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a72305820aa35573fce77fd4a1b828b381b33aadcb44f37da4db9695264b65f28256454be0029`

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
