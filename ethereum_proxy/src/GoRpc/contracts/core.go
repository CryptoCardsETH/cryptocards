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
const BattleGroupsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoCardsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_groupID\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countBattleGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_groupID\",\"type\":\"uint256\"}],\"name\":\"setGroupOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_CARDS_PER_GROUP\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_groupID\",\"type\":\"uint256\"}],\"name\":\"isGroupReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_cards\",\"type\":\"uint256[5]\"}],\"name\":\"createBattleGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"battleGroupID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cards\",\"type\":\"uint256[5]\"}],\"name\":\"NewBattleGroup\",\"type\":\"event\"}]"

// BattleGroupsBin is the compiled bytecode used for deploying new contracts.
const BattleGroupsBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561077c806100396000396000f300606060405236156100675763ffffffff60e060020a600035041663624fe81f811461006c5780636352211e1461009b5780637bf13d82146100b1578063b42cb37d146100d6578063bfeea70814610100578063ecf9036b14610113578063fa74efc114610129575b600080fd5b341561007757600080fd5b61007f61014a565b604051600160a060020a03909116815260200160405180910390f35b34156100a657600080fd5b61007f600435610159565b34156100bc57600080fd5b6100c46101b5565b60405190815260200160405180910390f35b34156100e157600080fd5b6100ec6004356101bc565b604051901515815260200160405180910390f35b341561010b57600080fd5b6100c4610317565b341561011e57600080fd5b6100ec60043561031c565b341561013457600080fd5b6100c4600160a060020a0360043516602461047a565b600154600160a060020a031681565b60006101636101b5565b821061016e57600080fd5b600080548390811061017c57fe5b6000918252602090912060069091020154680100000000000000009004600160a060020a031690508015156101b057600080fd5b919050565b6000545b90565b60006101c661060b565b60006101d06101b5565b84106101db57600080fd5b60008054859081106101e957fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b81548152602001906001019080831161024f5750505091909252509193506000925050505b816040015150600581101561030b57600154600160a060020a031663685b7a5e604084015183600581106102a457fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102e857600080fd5b6102c65a03f115156102f957600080fd5b50505060405180515050600101610274565b600192505b5050919050565b600581565b600061032661060b565b60006103306101b5565b841061033b57600080fd5b600080548590811061034957fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116103af5750505091909252509193506000925050505b816040015150600581101561030b57600154600160a060020a0316632109f89d6040840151836005811061040457fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561044857600080fd5b6102c65a03f1151561045957600080fd5b5050506040518051905015156104725760009250610310565b6001016103d4565b600061048461060b565b60006060604051908101604052804267ffffffffffffffff16815260200186600160a060020a03168152602001856005806020026040519081016040529190828260a080828437505050919092525050600080549193506001918083016104eb8382610631565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518154600160a060020a039190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516105749060018301906005610662565b50505003905063ffffffff8116811461058c57600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d85828460400151604051600160a060020a038416815260208101839052604081018260a080838360005b838110156105ef5780820151838201526020016105d7565b50505050905001935050505060405180910390a1949350505050565b60e0604051908101604090815260008083526020830152810161062c6106a0565b905290565b81548183558181151161065d5760060281600602836000526020600020918201910161065d91906106c7565b505050565b8260058101928215610690579160200282015b82811115610690578251825591602001919060010190610675565b5061069c929150610713565b5090565b60a06040519081016040526005815b60008152602001906001900390816106af5790505090565b6101b991905b8082111561069c5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061070a600183018261072d565b506006016106cd565b6101b991905b8082111561069c5760008155600101610719565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a7230582004abbec0c641e4cf752f6cfe4fb5293843bee15a71e9ff11f738d37f4cbe16dd0029`

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
// Solidity: function isGroupReadyForBattle(_groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsCaller) IsGroupReadyForBattle(opts *bind.CallOpts, _groupID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "isGroupReadyForBattle", _groupID)
	return *ret0, err
}

// IsGroupReadyForBattle is a free data retrieval call binding the contract method 0xecf9036b.
//
// Solidity: function isGroupReadyForBattle(_groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsSession) IsGroupReadyForBattle(_groupID *big.Int) (bool, error) {
	return _BattleGroups.Contract.IsGroupReadyForBattle(&_BattleGroups.CallOpts, _groupID)
}

// IsGroupReadyForBattle is a free data retrieval call binding the contract method 0xecf9036b.
//
// Solidity: function isGroupReadyForBattle(_groupID uint256) constant returns(bool)
func (_BattleGroups *BattleGroupsCallerSession) IsGroupReadyForBattle(_groupID *big.Int) (bool, error) {
	return _BattleGroups.Contract.IsGroupReadyForBattle(&_BattleGroups.CallOpts, _groupID)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_groupID uint256) constant returns(owner address)
func (_BattleGroups *BattleGroupsCaller) OwnerOf(opts *bind.CallOpts, _groupID *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BattleGroups.contract.Call(opts, out, "ownerOf", _groupID)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_groupID uint256) constant returns(owner address)
func (_BattleGroups *BattleGroupsSession) OwnerOf(_groupID *big.Int) (common.Address, error) {
	return _BattleGroups.Contract.OwnerOf(&_BattleGroups.CallOpts, _groupID)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_groupID uint256) constant returns(owner address)
func (_BattleGroups *BattleGroupsCallerSession) OwnerOf(_groupID *big.Int) (common.Address, error) {
	return _BattleGroups.Contract.OwnerOf(&_BattleGroups.CallOpts, _groupID)
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
// Solidity: function setGroupOnBattleCooldown(_groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsTransactor) SetGroupOnBattleCooldown(opts *bind.TransactOpts, _groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.contract.Transact(opts, "setGroupOnBattleCooldown", _groupID)
}

// SetGroupOnBattleCooldown is a paid mutator transaction binding the contract method 0xb42cb37d.
//
// Solidity: function setGroupOnBattleCooldown(_groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsSession) SetGroupOnBattleCooldown(_groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.SetGroupOnBattleCooldown(&_BattleGroups.TransactOpts, _groupID)
}

// SetGroupOnBattleCooldown is a paid mutator transaction binding the contract method 0xb42cb37d.
//
// Solidity: function setGroupOnBattleCooldown(_groupID uint256) returns(bool)
func (_BattleGroups *BattleGroupsTransactorSession) SetGroupOnBattleCooldown(_groupID *big.Int) (*types.Transaction, error) {
	return _BattleGroups.Contract.SetGroupOnBattleCooldown(&_BattleGroups.TransactOpts, _groupID)
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
const BattleQueueBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561034f806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461002e578063969494201461005d57600080fd5b341561003957600080fd5b610041610087565b604051600160a060020a03909116815260200160405180910390f35b341561006857600080fd5b610073600435610096565b604051901515815260200160405180910390f35b600154600160a060020a031681565b60006100a182610176565b15156100af57506000610171565b7f4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7824260405191825267ffffffffffffffff1660208201526040908101905180910390a16000541515610109575060008190556001610171565b7f9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a0548260005442604051928352602083019190915267ffffffffffffffff166040808301919091526060909101905180910390a161016882600054610249565b50506000805560015b919050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156101c057600080fd5b6102c65a03f115156101d157600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561022957600080fd5b6102c65a03f1151561023a57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e01196ef82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029357600080fd5b6102c65a03f115156102a457600080fd5b50505060405180519050600160a060020a031663df73c49f848460006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561030257600080fd5b6102c65a03f1151561031357600080fd5b50505060405180519493505050505600a165627a7a72305820559e2597a25686984f19672b4062de93ef47b200e6fd5ac6a21e9722458032160029`

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
const BattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cryptoCardsContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"countBattles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op1\",\"type\":\"uint256\"},{\"name\":\"_op2\",\"type\":\"uint256\"}],\"name\":\"createBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"battleID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winnerBattleGroup\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loserBattleGroup\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"rewardCardID\",\"type\":\"uint256\"}],\"name\":\"BattleResult\",\"type\":\"event\"}]"

// BattlesBin is the compiled bytecode used for deploying new contracts.
const BattlesBin = `0x6060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a0316179055610746806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461003957806363c785aa14610068578063df73c49f1461008d57600080fd5b341561004457600080fd5b61004c6100a6565b604051600160a060020a03909116815260200160405180910390f35b341561007357600080fd5b61007b6100b5565b60405190815260200160405180910390f35b341561009857600080fd5b61007b6004356024356100bc565b600154600160a060020a031681565b6000545b90565b60006100c661066a565b6000806000806100d588610248565b15806100e757506100e587610248565b155b156100f5576000955061023d565b60a0604051908101604052804267ffffffffffffffff168152602001898152602001888152602001600081526020016000815250945060016000805480600101828161014191906106a4565b600092835260209092208891600502018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151816003015560808201516004909101555003935063ffffffff841684146101af57600080fd5b6101b88561031b565b506101c384866104c3565b606087018281526080880182905291945092506101e09051610514565b90507f78d72c60b0f38a9a8dac0805b805c5f038ede4b314e696b2caefc2161058ffa58486606001518760800151846040518085815260200184815260200183815260200182815260200194505050505060405180910390a18395505b505050505092915050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029257600080fd5b6102c65a03f115156102a357600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102fb57600080fd5b6102c65a03f1151561030c57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561036557600080fd5b6102c65a03f1151561037657600080fd5b50505060405180519050600160a060020a031663b42cb37d836020015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103d257600080fd5b6102c65a03f115156103e357600080fd5b50505060405180515050600154600160a060020a031663e77dad5c6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561043557600080fd5b6102c65a03f1151561044657600080fd5b50505060405180519050600160a060020a031663b42cb37d836040015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156104a257600080fd5b6102c65a03f115156104b357600080fd5b5050506040518051905050919050565b60008060006002846020015185604001518701018115156104e057fe5b0690508015156104fd57836020015184604001519250925061050c565b83604001518460200151925092505b509250929050565b6001546000908190600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561056057600080fd5b6102c65a03f1151561057157600080fd5b50505060405180519050600160a060020a0316636352211e8460006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156105c957600080fd5b6102c65a03f115156105da57600080fd5b5050506040518051600154909250600160a060020a03169050635de038b1826201b20760006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561064957600080fd5b6102c65a03f1151561065a57600080fd5b5050506040518051949350505050565b60a060405190810160405280600067ffffffffffffffff168152602001600081526020016000815260200160008152602001600081525090565b8154818355818115116106d0576005028160050283600052602060002091820191016106d091906106d5565b505050565b6100b991905b8082111561071657805467ffffffffffffffff19168155600060018201819055600282018190556003820181905560048201556005016106db565b50905600a165627a7a72305820386a2b1be06d6b92619bba00bb6e6b79859934336616994788f50c4d2935e9930029`

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
	RewardCardID      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBattleResult is a free log retrieval operation binding the contract event 0x78d72c60b0f38a9a8dac0805b805c5f038ede4b314e696b2caefc2161058ffa5.
//
// Solidity: event BattleResult(battleID uint256, winnerBattleGroup uint256, loserBattleGroup uint256, rewardCardID uint256)
func (_Battles *BattlesFilterer) FilterBattleResult(opts *bind.FilterOpts) (*BattlesBattleResultIterator, error) {

	logs, sub, err := _Battles.contract.FilterLogs(opts, "BattleResult")
	if err != nil {
		return nil, err
	}
	return &BattlesBattleResultIterator{contract: _Battles.contract, event: "BattleResult", logs: logs, sub: sub}, nil
}

// WatchBattleResult is a free log subscription operation binding the contract event 0x78d72c60b0f38a9a8dac0805b805c5f038ede4b314e696b2caefc2161058ffa5.
//
// Solidity: event BattleResult(battleID uint256, winnerBattleGroup uint256, loserBattleGroup uint256, rewardCardID uint256)
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
const CardBaseABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CardBaseBin is the compiled bytecode used for deploying new contracts.
const CardBaseBin = `0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b6103fc8061002e6000396000f300606060405236156100965763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461009b57806313af4035146100cf5780633f4ba83a146100f05780635c975abb146101035780636af04a571461012a578063715879881461015957806378d0df50146101785780638456cb591461018e5780638f84aa09146101a1575b600080fd5b34156100a657600080fd5b6100bd600160a060020a03600435166024356101b4565b60405190815260200160405180910390f35b34156100da57600080fd5b6100ee600160a060020a03600435166101e2565b005b34156100fb57600080fd5b6100ee610241565b341561010e57600080fd5b610116610294565b604051901515815260200160405180910390f35b341561013557600080fd5b61013d6102a4565b604051600160a060020a03909116815260200160405180910390f35b341561016457600080fd5b6100ee600160a060020a03600435166102b3565b341561018357600080fd5b61013d60043561034e565b341561019957600080fd5b6100ee610369565b34156101ac57600080fd5b61013d6103c1565b6004602052816000526040600020818154811015156101cf57fe5b6000918252602090912001549150829050565b60005433600160a060020a039081169116146101fd57600080fd5b600160a060020a038116151561021257600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461025c57600080fd5b60005460a060020a900460ff16151561027457600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b600154600160a060020a031681565b60005433600160a060020a039081169116146102ce57600080fd5b60005460a060020a900460ff1615156102e657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600360205260009081526040902054600160a060020a031681565b60005433600160a060020a0390811691161461038457600080fd5b60005460a060020a900460ff161561039b57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b600054600160a060020a0316815600a165627a7a723058200c6d42a8417c7503a9fea7d433c462d6834ae2ee1f8ba83449a06003731040330029`

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

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBase *CardBaseCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBase.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBase *CardBaseSession) NewContractAddress() (common.Address, error) {
	return _CardBase.Contract.NewContractAddress(&_CardBase.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBase *CardBaseCallerSession) NewContractAddress() (common.Address, error) {
	return _CardBase.Contract.NewContractAddress(&_CardBase.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBase *CardBaseCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBase.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBase *CardBaseSession) OwnerAddress() (common.Address, error) {
	return _CardBase.Contract.OwnerAddress(&_CardBase.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBase *CardBaseCallerSession) OwnerAddress() (common.Address, error) {
	return _CardBase.Contract.OwnerAddress(&_CardBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBase *CardBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardBase.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBase *CardBaseSession) Paused() (bool, error) {
	return _CardBase.Contract.Paused(&_CardBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBase *CardBaseCallerSession) Paused() (bool, error) {
	return _CardBase.Contract.Paused(&_CardBase.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBase *CardBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBase *CardBaseSession) Pause() (*types.Transaction, error) {
	return _CardBase.Contract.Pause(&_CardBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBase *CardBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _CardBase.Contract.Pause(&_CardBase.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBase *CardBaseTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _CardBase.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBase *CardBaseSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardBase.Contract.SetNewAddress(&_CardBase.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBase *CardBaseTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardBase.Contract.SetNewAddress(&_CardBase.TransactOpts, _newAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBase *CardBaseTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CardBase.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBase *CardBaseSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardBase.Contract.SetOwner(&_CardBase.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBase *CardBaseTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardBase.Contract.SetOwner(&_CardBase.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBase *CardBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBase *CardBaseSession) Unpause() (*types.Transaction, error) {
	return _CardBase.Contract.Unpause(&_CardBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBase *CardBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _CardBase.Contract.Unpause(&_CardBase.TransactOpts)
}

// CardBaseContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CardBase contract.
type CardBaseContractUpgradeIterator struct {
	Event *CardBaseContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CardBaseContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBaseContractUpgrade)
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
		it.Event = new(CardBaseContractUpgrade)
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
func (it *CardBaseContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBaseContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBaseContractUpgrade represents a ContractUpgrade event raised by the CardBase contract.
type CardBaseContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardBase *CardBaseFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CardBaseContractUpgradeIterator, error) {

	logs, sub, err := _CardBase.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CardBaseContractUpgradeIterator{contract: _CardBase.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardBase *CardBaseFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CardBaseContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CardBase.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBaseContractUpgrade)
				if err := _CardBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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
const CardBattlesABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CardBattlesBin is the compiled bytecode used for deploying new contracts.
const CardBattlesBin = `0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b610db68061002e6000396000f3006060604052361561011a5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461011f57806312c2debf146101a957806313af4035146101dd57806318160ddd146101fe5780632109f89d14610211578063302597201461023b5780633f4ba83a146102515780635c975abb146102645780635de038b1146102775780636352211e14610299578063685b7a5e146102cb5780636af04a57146102e157806370a08231146102f4578063715879881461031357806378d0df50146103325780638456cb59146103485780638462151c1461035b5780638f84aa09146103cd57806395d89b41146103e05780639db02257146103f3578063a9059cbb14610423575b600080fd5b341561012a57600080fd5b610132610445565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561016e578082015183820152602001610156565b50505050905090810190601f16801561019b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b457600080fd5b6101cb600160a060020a036004351660243561047c565b60405190815260200160405180910390f35b34156101e857600080fd5b6101fc600160a060020a03600435166104aa565b005b341561020957600080fd5b6101cb610509565b341561021c57600080fd5b610227600435610510565b604051901515815260200160405180910390f35b341561024657600080fd5b6101fc600435610589565b341561025c57600080fd5b6101fc61059f565b341561026f57600080fd5b6102276105f2565b341561028257600080fd5b6101cb600160a060020a0360043516602435610602565b34156102a457600080fd5b6102af600435610651565b604051600160a060020a03909116815260200160405180910390f35b34156102d657600080fd5b610227600435610675565b34156102ec57600080fd5b6102af6106d4565b34156102ff57600080fd5b6101cb600160a060020a03600435166106e3565b341561031e57600080fd5b6101fc600160a060020a03600435166106fe565b341561033d57600080fd5b6102af600435610799565b341561035357600080fd5b6101fc6107b4565b341561036657600080fd5b61037a600160a060020a036004351661080c565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103b95780820151838201526020016103a1565b505050509050019250505060405180910390f35b34156103d857600080fd5b6102af6108ed565b34156103eb57600080fd5b6101326108fc565b34156103fe57600080fd5b610406610933565b60405167ffffffffffffffff909116815260200160405180910390f35b341561042e57600080fd5b6101fc600160a060020a0360043516602435610939565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60046020528160005260406000208181548110151561049757fe5b6000918252602090912001549150829050565b60005433600160a060020a039081169116146104c557600080fd5b600160a060020a03811615156104da57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002545b90565b6000805460a060020a900460ff161561052857600080fd5b61053182610589565b4267ffffffffffffffff1660028381548110151561054b57fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff16111561058057506000610584565b5060015b919050565b610591610509565b811061059c57600080fd5b50565b60005433600160a060020a039081169116146105ba57600080fd5b60005460a060020a900460ff1615156105d257600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff161561061c57600080fd5b30600160a060020a031684600160a060020a03161415151561063d57600080fd5b610649600084866109aa565b949350505050565b600081815260036020526040902054600160a060020a031680151561058457600080fd5b600061068082610589565b610e10420160028381548110151561069457fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600154600160a060020a031681565b600160a060020a031660009081526005602052604090205490565b60005433600160a060020a0390811691161461071957600080fd5b60005460a060020a900460ff16151561073157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600360205260009081526040902054600160a060020a031681565b60005433600160a060020a039081169116146107cf57600080fd5b60005460a060020a900460ff16156107e657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b610814610ce7565b600061081e610ce7565b600080600061082c876106e3565b945084151561085c5760006040518059106108445750595b908082528060200260200182016040525095506108e3565b8460405180591061086a5750595b90808252806020026020018201604052509350610885610509565b925060009150600190505b8281116108df57600081815260036020526040902054600160a060020a03888116911614156108d757808483815181106108c657fe5b602090810290910101526001909101905b600101610890565b8395505b5050505050919050565b600054600160a060020a031681565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b60005460a060020a900460ff161561095057600080fd5b600160a060020a038216151561096557600080fd5b30600160a060020a031682600160a060020a03161415151561098657600080fd5b6109903382610bff565b151561099b57600080fd5b6109a6338383610c1f565b5050565b60006109b4610cf9565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152602001868152509150600160028054806001018281610a499190610d2e565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff81168114610b7657600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a1610bf660008583610c1f565b95945050505050565b600090815260036020526040902054600160a060020a0391821691161490565b600160a060020a0380831660008181526005602090815260408083208054600101905585835260039091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055831615610c9257600160a060020a038316600090815260056020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610d5a57600302816003028360005260206000209182019101610d5a9190610d5f565b505050565b61050d91905b80821115610d86576000808255600182018190556002820155600301610d65565b50905600a165627a7a7230582081a0bcddd07e5e05263f196b845a586d5909ce96c035916f9fb7fc76b854adf10029`

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

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBattles *CardBattlesCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBattles *CardBattlesSession) NewContractAddress() (common.Address, error) {
	return _CardBattles.Contract.NewContractAddress(&_CardBattles.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardBattles *CardBattlesCallerSession) NewContractAddress() (common.Address, error) {
	return _CardBattles.Contract.NewContractAddress(&_CardBattles.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBattles *CardBattlesCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBattles *CardBattlesSession) OwnerAddress() (common.Address, error) {
	return _CardBattles.Contract.OwnerAddress(&_CardBattles.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardBattles *CardBattlesCallerSession) OwnerAddress() (common.Address, error) {
	return _CardBattles.Contract.OwnerAddress(&_CardBattles.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBattles *CardBattlesCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardBattles.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBattles *CardBattlesSession) Paused() (bool, error) {
	return _CardBattles.Contract.Paused(&_CardBattles.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardBattles *CardBattlesCallerSession) Paused() (bool, error) {
	return _CardBattles.Contract.Paused(&_CardBattles.CallOpts)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBattles *CardBattlesTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBattles *CardBattlesSession) Pause() (*types.Transaction, error) {
	return _CardBattles.Contract.Pause(&_CardBattles.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardBattles *CardBattlesTransactorSession) Pause() (*types.Transaction, error) {
	return _CardBattles.Contract.Pause(&_CardBattles.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBattles *CardBattlesTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBattles *CardBattlesSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardBattles.Contract.SetNewAddress(&_CardBattles.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardBattles *CardBattlesTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardBattles.Contract.SetNewAddress(&_CardBattles.TransactOpts, _newAddress)
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

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBattles *CardBattlesTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBattles *CardBattlesSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardBattles.Contract.SetOwner(&_CardBattles.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardBattles *CardBattlesTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardBattles.Contract.SetOwner(&_CardBattles.TransactOpts, _newOwner)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBattles *CardBattlesTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardBattles.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBattles *CardBattlesSession) Unpause() (*types.Transaction, error) {
	return _CardBattles.Contract.Unpause(&_CardBattles.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardBattles *CardBattlesTransactorSession) Unpause() (*types.Transaction, error) {
	return _CardBattles.Contract.Unpause(&_CardBattles.TransactOpts)
}

// CardBattlesContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CardBattles contract.
type CardBattlesContractUpgradeIterator struct {
	Event *CardBattlesContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CardBattlesContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardBattlesContractUpgrade)
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
		it.Event = new(CardBattlesContractUpgrade)
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
func (it *CardBattlesContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardBattlesContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardBattlesContractUpgrade represents a ContractUpgrade event raised by the CardBattles contract.
type CardBattlesContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardBattles *CardBattlesFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CardBattlesContractUpgradeIterator, error) {

	logs, sub, err := _CardBattles.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CardBattlesContractUpgradeIterator{contract: _CardBattles.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardBattles *CardBattlesFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CardBattlesContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CardBattles.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardBattlesContractUpgrade)
				if err := _CardBattles.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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
const CardMintingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"newCardID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CardMintingBin is the compiled bytecode used for deploying new contracts.
const CardMintingBin = `0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b610dff8061002e6000396000f300606060405236156101255763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461012a57806312c2debf146101b457806313af4035146101e857806318160ddd146102095780632109f89d1461021c57806330259720146102465780633f4ba83a1461025c578063458129581461026f5780635c975abb146102825780635de038b1146102955780636352211e146102b7578063685b7a5e146102e95780636af04a57146102ff57806370a0823114610312578063715879881461033157806378d0df50146103505780638456cb59146103665780638462151c146103795780638f84aa09146103eb57806395d89b41146103fe5780639db0225714610411578063a9059cbb14610441575b600080fd5b341561013557600080fd5b61013d610463565b60405160208082528190810183818151815260200191508051906020019080838360005b83811015610179578082015183820152602001610161565b50505050905090810190601f1680156101a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101bf57600080fd5b6101d6600160a060020a036004351660243561049a565b60405190815260200160405180910390f35b34156101f357600080fd5b610207600160a060020a03600435166104c8565b005b341561021457600080fd5b6101d6610527565b341561022757600080fd5b61023260043561052e565b604051901515815260200160405180910390f35b341561025157600080fd5b6102076004356105a7565b341561026757600080fd5b6102076105bd565b341561027a57600080fd5b6101d6610610565b341561028d57600080fd5b61023261063b565b34156102a057600080fd5b6101d6600160a060020a036004351660243561064b565b34156102c257600080fd5b6102cd60043561069a565b604051600160a060020a03909116815260200160405180910390f35b34156102f457600080fd5b6102326004356106be565b341561030a57600080fd5b6102cd61071d565b341561031d57600080fd5b6101d6600160a060020a036004351661072c565b341561033c57600080fd5b610207600160a060020a0360043516610747565b341561035b57600080fd5b6102cd6004356107e2565b341561037157600080fd5b6102076107fd565b341561038457600080fd5b610398600160a060020a0360043516610855565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103d75780820151838201526020016103bf565b505050509050019250505060405180910390f35b34156103f657600080fd5b6102cd610936565b341561040957600080fd5b61013d610945565b341561041c57600080fd5b61042461097c565b60405167ffffffffffffffff909116815260200160405180910390f35b341561044c57600080fd5b610207600160a060020a0360043516602435610982565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6004602052816000526040600020818154811015156104b557fe5b6000918252602090912001549150829050565b60005433600160a060020a039081169116146104e357600080fd5b600160a060020a03811615156104f857600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002545b90565b6000805460a060020a900460ff161561054657600080fd5b61054f826105a7565b4267ffffffffffffffff1660028381548110151561056957fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff16111561059e575060006105a2565b5060015b919050565b6105af610527565b81106105ba57600080fd5b50565b60005433600160a060020a039081169116146105d857600080fd5b60005460a060020a900460ff1615156105f057600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b6000805460a060020a900460ff161561062857600080fd5b61063660006201e24061064b565b905090565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff161561066557600080fd5b30600160a060020a031684600160a060020a03161415151561068657600080fd5b610692600084866109f3565b949350505050565b600081815260036020526040902054600160a060020a03168015156105a257600080fd5b60006106c9826105a7565b610e1042016002838154811015156106dd57fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600154600160a060020a031681565b600160a060020a031660009081526005602052604090205490565b60005433600160a060020a0390811691161461076257600080fd5b60005460a060020a900460ff16151561077a57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600360205260009081526040902054600160a060020a031681565b60005433600160a060020a0390811691161461081857600080fd5b60005460a060020a900460ff161561082f57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b61085d610d30565b6000610867610d30565b60008060006108758761072c565b94508415156108a557600060405180591061088d5750595b9080825280602002602001820160405250955061092c565b846040518059106108b35750595b908082528060200260200182016040525093506108ce610527565b925060009150600190505b82811161092857600081815260036020526040902054600160a060020a0388811691161415610920578084838151811061090f57fe5b602090810290910101526001909101905b6001016108d9565b8395505b5050505050919050565b600054600160a060020a031681565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b60005460a060020a900460ff161561099957600080fd5b600160a060020a03821615156109ae57600080fd5b30600160a060020a031682600160a060020a0316141515156109cf57600080fd5b6109d93382610c48565b15156109e457600080fd5b6109ef338383610c68565b5050565b60006109fd610d42565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152602001868152509150600160028054806001018281610a929190610d77565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff81168114610bbf57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a1610c3f60008583610c68565b95945050505050565b600090815260036020526040902054600160a060020a0391821691161490565b600160a060020a0380831660008181526005602090815260408083208054600101905585835260039091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055831615610cdb57600160a060020a038316600090815260056020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610da357600302816003028360005260206000209182019101610da39190610da8565b505050565b61052b91905b80821115610dcf576000808255600182018190556002820155600301610dae565b50905600a165627a7a72305820e80cc544571a4c9fb19b62e124a029a9adfa1494a4e506a6ec93d38df4b58eff0029`

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

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardMinting *CardMintingCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardMinting *CardMintingSession) NewContractAddress() (common.Address, error) {
	return _CardMinting.Contract.NewContractAddress(&_CardMinting.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardMinting *CardMintingCallerSession) NewContractAddress() (common.Address, error) {
	return _CardMinting.Contract.NewContractAddress(&_CardMinting.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardMinting *CardMintingCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardMinting *CardMintingSession) OwnerAddress() (common.Address, error) {
	return _CardMinting.Contract.OwnerAddress(&_CardMinting.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardMinting *CardMintingCallerSession) OwnerAddress() (common.Address, error) {
	return _CardMinting.Contract.OwnerAddress(&_CardMinting.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardMinting *CardMintingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardMinting.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardMinting *CardMintingSession) Paused() (bool, error) {
	return _CardMinting.Contract.Paused(&_CardMinting.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardMinting *CardMintingCallerSession) Paused() (bool, error) {
	return _CardMinting.Contract.Paused(&_CardMinting.CallOpts)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardMinting *CardMintingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardMinting *CardMintingSession) Pause() (*types.Transaction, error) {
	return _CardMinting.Contract.Pause(&_CardMinting.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardMinting *CardMintingTransactorSession) Pause() (*types.Transaction, error) {
	return _CardMinting.Contract.Pause(&_CardMinting.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardMinting *CardMintingTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardMinting *CardMintingSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardMinting.Contract.SetNewAddress(&_CardMinting.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardMinting *CardMintingTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardMinting.Contract.SetNewAddress(&_CardMinting.TransactOpts, _newAddress)
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

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardMinting *CardMintingTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardMinting *CardMintingSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardMinting.Contract.SetOwner(&_CardMinting.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardMinting *CardMintingTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardMinting.Contract.SetOwner(&_CardMinting.TransactOpts, _newOwner)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardMinting *CardMintingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardMinting.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardMinting *CardMintingSession) Unpause() (*types.Transaction, error) {
	return _CardMinting.Contract.Unpause(&_CardMinting.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardMinting *CardMintingTransactorSession) Unpause() (*types.Transaction, error) {
	return _CardMinting.Contract.Unpause(&_CardMinting.TransactOpts)
}

// CardMintingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CardMinting contract.
type CardMintingContractUpgradeIterator struct {
	Event *CardMintingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CardMintingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardMintingContractUpgrade)
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
		it.Event = new(CardMintingContractUpgrade)
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
func (it *CardMintingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardMintingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardMintingContractUpgrade represents a ContractUpgrade event raised by the CardMinting contract.
type CardMintingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardMinting *CardMintingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CardMintingContractUpgradeIterator, error) {

	logs, sub, err := _CardMinting.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CardMintingContractUpgradeIterator{contract: _CardMinting.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardMinting *CardMintingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CardMintingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CardMinting.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardMintingContractUpgrade)
				if err := _CardMinting.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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
const CardOwnershipABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CardOwnershipBin is the compiled bytecode used for deploying new contracts.
const CardOwnershipBin = `0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b610c608061002e6000396000f300606060405236156100f95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100fe57806312c2debf1461018857806313af4035146101bc57806318160ddd146101dd57806330259720146101f05780633f4ba83a146102065780635c975abb146102195780635de038b1146102405780636352211e146102625780636af04a571461029457806370a08231146102a757806371587988146102c657806378d0df50146102e55780638456cb59146102fb5780638462151c1461030e5780638f84aa091461038057806395d89b4114610393578063a9059cbb146103a6575b600080fd5b341561010957600080fd5b6101116103c8565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561014d578082015183820152602001610135565b50505050905090810190601f16801561017a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561019357600080fd5b6101aa600160a060020a03600435166024356103ff565b60405190815260200160405180910390f35b34156101c757600080fd5b6101db600160a060020a036004351661042d565b005b34156101e857600080fd5b6101aa61048c565b34156101fb57600080fd5b6101db600435610493565b341561021157600080fd5b6101db6104a9565b341561022457600080fd5b61022c6104fc565b604051901515815260200160405180910390f35b341561024b57600080fd5b6101aa600160a060020a036004351660243561050c565b341561026d57600080fd5b61027860043561055b565b604051600160a060020a03909116815260200160405180910390f35b341561029f57600080fd5b610278610584565b34156102b257600080fd5b6101aa600160a060020a0360043516610593565b34156102d157600080fd5b6101db600160a060020a03600435166105ae565b34156102f057600080fd5b610278600435610649565b341561030657600080fd5b6101db610664565b341561031957600080fd5b61032d600160a060020a03600435166106bc565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561036c578082015183820152602001610354565b505050509050019250505060405180910390f35b341561038b57600080fd5b61027861079d565b341561039e57600080fd5b6101116107ac565b34156103b157600080fd5b6101db600160a060020a03600435166024356107e3565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60046020528160005260406000208181548110151561041a57fe5b6000918252602090912001549150829050565b60005433600160a060020a0390811691161461044857600080fd5b600160a060020a038116151561045d57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002545b90565b61049b61048c565b81106104a657600080fd5b50565b60005433600160a060020a039081169116146104c457600080fd5b60005460a060020a900460ff1615156104dc57600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff161561052657600080fd5b30600160a060020a031684600160a060020a03161415151561054757600080fd5b61055360008486610854565b949350505050565b600081815260036020526040902054600160a060020a031680151561057f57600080fd5b919050565b600154600160a060020a031681565b600160a060020a031660009081526005602052604090205490565b60005433600160a060020a039081169116146105c957600080fd5b60005460a060020a900460ff1615156105e157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600360205260009081526040902054600160a060020a031681565b60005433600160a060020a0390811691161461067f57600080fd5b60005460a060020a900460ff161561069657600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b6106c4610b91565b60006106ce610b91565b60008060006106dc87610593565b945084151561070c5760006040518059106106f45750595b90808252806020026020018201604052509550610793565b8460405180591061071a5750595b9080825280602002602001820160405250935061073561048c565b925060009150600190505b82811161078f57600081815260036020526040902054600160a060020a0388811691161415610787578084838151811061077657fe5b602090810290910101526001909101905b600101610740565b8395505b5050505050919050565b600054600160a060020a031681565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b60005460a060020a900460ff16156107fa57600080fd5b600160a060020a038216151561080f57600080fd5b30600160a060020a031682600160a060020a03161415151561083057600080fd5b61083a3382610aa9565b151561084557600080fd5b610850338383610ac9565b5050565b600061085e610ba3565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600280548060010182816108f39190610bd8565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff81168114610a2057600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a1610aa060008583610ac9565b95945050505050565b600090815260036020526040902054600160a060020a0391821691161490565b600160a060020a0380831660008181526005602090815260408083208054600101905585835260039091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055831615610b3c57600160a060020a038316600090815260056020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610c0457600302816003028360005260206000209182019101610c049190610c09565b505050565b61049091905b80821115610c30576000808255600182018190556002820155600301610c0f565b50905600a165627a7a723058206ace7a487d4e2f36b4f9f5582513be3137cae7ca43491d9545949237fd7163ab0029`

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

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardOwnership *CardOwnershipCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardOwnership *CardOwnershipSession) NewContractAddress() (common.Address, error) {
	return _CardOwnership.Contract.NewContractAddress(&_CardOwnership.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_CardOwnership *CardOwnershipCallerSession) NewContractAddress() (common.Address, error) {
	return _CardOwnership.Contract.NewContractAddress(&_CardOwnership.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardOwnership *CardOwnershipCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardOwnership *CardOwnershipSession) OwnerAddress() (common.Address, error) {
	return _CardOwnership.Contract.OwnerAddress(&_CardOwnership.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_CardOwnership *CardOwnershipCallerSession) OwnerAddress() (common.Address, error) {
	return _CardOwnership.Contract.OwnerAddress(&_CardOwnership.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardOwnership *CardOwnershipCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CardOwnership.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardOwnership *CardOwnershipSession) Paused() (bool, error) {
	return _CardOwnership.Contract.Paused(&_CardOwnership.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CardOwnership *CardOwnershipCallerSession) Paused() (bool, error) {
	return _CardOwnership.Contract.Paused(&_CardOwnership.CallOpts)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardOwnership *CardOwnershipTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardOwnership *CardOwnershipSession) Pause() (*types.Transaction, error) {
	return _CardOwnership.Contract.Pause(&_CardOwnership.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CardOwnership *CardOwnershipTransactorSession) Pause() (*types.Transaction, error) {
	return _CardOwnership.Contract.Pause(&_CardOwnership.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardOwnership *CardOwnershipTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardOwnership *CardOwnershipSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardOwnership.Contract.SetNewAddress(&_CardOwnership.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_CardOwnership *CardOwnershipTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _CardOwnership.Contract.SetNewAddress(&_CardOwnership.TransactOpts, _newAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardOwnership *CardOwnershipTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardOwnership *CardOwnershipSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardOwnership.Contract.SetOwner(&_CardOwnership.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CardOwnership *CardOwnershipTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CardOwnership.Contract.SetOwner(&_CardOwnership.TransactOpts, _newOwner)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardOwnership *CardOwnershipTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CardOwnership.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardOwnership *CardOwnershipSession) Unpause() (*types.Transaction, error) {
	return _CardOwnership.Contract.Unpause(&_CardOwnership.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CardOwnership *CardOwnershipTransactorSession) Unpause() (*types.Transaction, error) {
	return _CardOwnership.Contract.Unpause(&_CardOwnership.TransactOpts)
}

// CardOwnershipContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the CardOwnership contract.
type CardOwnershipContractUpgradeIterator struct {
	Event *CardOwnershipContractUpgrade // Event containing the contract specifics and raw log

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
func (it *CardOwnershipContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CardOwnershipContractUpgrade)
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
		it.Event = new(CardOwnershipContractUpgrade)
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
func (it *CardOwnershipContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CardOwnershipContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CardOwnershipContractUpgrade represents a ContractUpgrade event raised by the CardOwnership contract.
type CardOwnershipContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardOwnership *CardOwnershipFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*CardOwnershipContractUpgradeIterator, error) {

	logs, sub, err := _CardOwnership.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &CardOwnershipContractUpgradeIterator{contract: _CardOwnership.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_CardOwnership *CardOwnershipFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *CardOwnershipContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _CardOwnership.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CardOwnershipContractUpgrade)
				if err := _CardOwnership.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ContractAccessControlABI is the input ABI used to generate the binding from.
const ContractAccessControlABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// ContractAccessControlBin is the compiled bytecode used for deploying new contracts.
const ContractAccessControlBin = `0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b6103538061002e6000396000f300606060405236156100805763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166313af403581146100855780633f4ba83a146100a65780635c975abb146100b95780636af04a57146100e0578063715879881461010f5780638456cb591461012e5780638f84aa0914610141575b600080fd5b341561009057600080fd5b6100a4600160a060020a0360043516610154565b005b34156100b157600080fd5b6100a46101b3565b34156100c457600080fd5b6100cc610206565b604051901515815260200160405180910390f35b34156100eb57600080fd5b6100f3610216565b604051600160a060020a03909116815260200160405180910390f35b341561011a57600080fd5b6100a4600160a060020a0360043516610225565b341561013957600080fd5b6100a46102c0565b341561014c57600080fd5b6100f3610318565b60005433600160a060020a0390811691161461016f57600080fd5b600160a060020a038116151561018457600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a039081169116146101ce57600080fd5b60005460a060020a900460ff1615156101e657600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b60005460a060020a900460ff1681565b600154600160a060020a031681565b60005433600160a060020a0390811691161461024057600080fd5b60005460a060020a900460ff16151561025857600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b60005433600160a060020a039081169116146102db57600080fd5b60005460a060020a900460ff16156102f257600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b600054600160a060020a0316815600a165627a7a72305820a174c923fd837be5b88aad67b54ff795ef6355b67ccdeb597df782593938b0100029`

// DeployContractAccessControl deploys a new Ethereum contract, binding an instance of ContractAccessControl to it.
func DeployContractAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractAccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAccessControl{ContractAccessControlCaller: ContractAccessControlCaller{contract: contract}, ContractAccessControlTransactor: ContractAccessControlTransactor{contract: contract}, ContractAccessControlFilterer: ContractAccessControlFilterer{contract: contract}}, nil
}

// ContractAccessControl is an auto generated Go binding around an Ethereum contract.
type ContractAccessControl struct {
	ContractAccessControlCaller     // Read-only binding to the contract
	ContractAccessControlTransactor // Write-only binding to the contract
	ContractAccessControlFilterer   // Log filterer for contract events
}

// ContractAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAccessControlSession struct {
	Contract     *ContractAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAccessControlCallerSession struct {
	Contract *ContractAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAccessControlTransactorSession struct {
	Contract     *ContractAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAccessControlRaw struct {
	Contract *ContractAccessControl // Generic contract binding to access the raw methods on
}

// ContractAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAccessControlCallerRaw struct {
	Contract *ContractAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAccessControlTransactorRaw struct {
	Contract *ContractAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAccessControl creates a new instance of ContractAccessControl, bound to a specific deployed contract.
func NewContractAccessControl(address common.Address, backend bind.ContractBackend) (*ContractAccessControl, error) {
	contract, err := bindContractAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAccessControl{ContractAccessControlCaller: ContractAccessControlCaller{contract: contract}, ContractAccessControlTransactor: ContractAccessControlTransactor{contract: contract}, ContractAccessControlFilterer: ContractAccessControlFilterer{contract: contract}}, nil
}

// NewContractAccessControlCaller creates a new read-only instance of ContractAccessControl, bound to a specific deployed contract.
func NewContractAccessControlCaller(address common.Address, caller bind.ContractCaller) (*ContractAccessControlCaller, error) {
	contract, err := bindContractAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAccessControlCaller{contract: contract}, nil
}

// NewContractAccessControlTransactor creates a new write-only instance of ContractAccessControl, bound to a specific deployed contract.
func NewContractAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAccessControlTransactor, error) {
	contract, err := bindContractAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAccessControlTransactor{contract: contract}, nil
}

// NewContractAccessControlFilterer creates a new log filterer instance of ContractAccessControl, bound to a specific deployed contract.
func NewContractAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAccessControlFilterer, error) {
	contract, err := bindContractAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAccessControlFilterer{contract: contract}, nil
}

// bindContractAccessControl binds a generic wrapper to an already deployed contract.
func bindContractAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAccessControl *ContractAccessControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAccessControl.Contract.ContractAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAccessControl *ContractAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.ContractAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAccessControl *ContractAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.ContractAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAccessControl *ContractAccessControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAccessControl *ContractAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAccessControl *ContractAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.contract.Transact(opts, method, params...)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractAccessControl.contract.Call(opts, out, "newContractAddress")
	return *ret0, err
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlSession) NewContractAddress() (common.Address, error) {
	return _ContractAccessControl.Contract.NewContractAddress(&_ContractAccessControl.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlCallerSession) NewContractAddress() (common.Address, error) {
	return _ContractAccessControl.Contract.NewContractAddress(&_ContractAccessControl.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlCaller) OwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractAccessControl.contract.Call(opts, out, "ownerAddress")
	return *ret0, err
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlSession) OwnerAddress() (common.Address, error) {
	return _ContractAccessControl.Contract.OwnerAddress(&_ContractAccessControl.CallOpts)
}

// OwnerAddress is a free data retrieval call binding the contract method 0x8f84aa09.
//
// Solidity: function ownerAddress() constant returns(address)
func (_ContractAccessControl *ContractAccessControlCallerSession) OwnerAddress() (common.Address, error) {
	return _ContractAccessControl.Contract.OwnerAddress(&_ContractAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ContractAccessControl *ContractAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ContractAccessControl.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ContractAccessControl *ContractAccessControlSession) Paused() (bool, error) {
	return _ContractAccessControl.Contract.Paused(&_ContractAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ContractAccessControl *ContractAccessControlCallerSession) Paused() (bool, error) {
	return _ContractAccessControl.Contract.Paused(&_ContractAccessControl.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAccessControl *ContractAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAccessControl *ContractAccessControlSession) Pause() (*types.Transaction, error) {
	return _ContractAccessControl.Contract.Pause(&_ContractAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAccessControl *ContractAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _ContractAccessControl.Contract.Pause(&_ContractAccessControl.TransactOpts)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_ContractAccessControl *ContractAccessControlTransactor) SetNewAddress(opts *bind.TransactOpts, _newAddress common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.contract.Transact(opts, "setNewAddress", _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_ContractAccessControl *ContractAccessControlSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.SetNewAddress(&_ContractAccessControl.TransactOpts, _newAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(_newAddress address) returns()
func (_ContractAccessControl *ContractAccessControlTransactorSession) SetNewAddress(_newAddress common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.SetNewAddress(&_ContractAccessControl.TransactOpts, _newAddress)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_ContractAccessControl *ContractAccessControlTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_ContractAccessControl *ContractAccessControlSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.SetOwner(&_ContractAccessControl.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_ContractAccessControl *ContractAccessControlTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _ContractAccessControl.Contract.SetOwner(&_ContractAccessControl.TransactOpts, _newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAccessControl *ContractAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAccessControl *ContractAccessControlSession) Unpause() (*types.Transaction, error) {
	return _ContractAccessControl.Contract.Unpause(&_ContractAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAccessControl *ContractAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _ContractAccessControl.Contract.Unpause(&_ContractAccessControl.TransactOpts)
}

// ContractAccessControlContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the ContractAccessControl contract.
type ContractAccessControlContractUpgradeIterator struct {
	Event *ContractAccessControlContractUpgrade // Event containing the contract specifics and raw log

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
func (it *ContractAccessControlContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAccessControlContractUpgrade)
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
		it.Event = new(ContractAccessControlContractUpgrade)
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
func (it *ContractAccessControlContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAccessControlContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAccessControlContractUpgrade represents a ContractUpgrade event raised by the ContractAccessControl contract.
type ContractAccessControlContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_ContractAccessControl *ContractAccessControlFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*ContractAccessControlContractUpgradeIterator, error) {

	logs, sub, err := _ContractAccessControl.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &ContractAccessControlContractUpgradeIterator{contract: _ContractAccessControl.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(newContract address)
func (_ContractAccessControl *ContractAccessControlFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *ContractAccessControlContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _ContractAccessControl.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAccessControlContractUpgrade)
				if err := _ContractAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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
const CryptoCardsCoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"isReadyForBattle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cardID\",\"type\":\"uint256\"}],\"name\":\"requireCardExists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createGen0Card\",\"outputs\":[{\"name\":\"newCardID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cardID\",\"type\":\"uint256\"}],\"name\":\"setOnBattleCooldown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAddress\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BATTLE_COOLDOWN_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleQueueContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BattleGroupContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// CryptoCardsCoreBin is the compiled bytecode used for deploying new contracts.
const CryptoCardsCoreBin = `0x60606040526000805460a060020a60ff021916905534156200002057600080fd5b60008054600160a060020a03191633600160a060020a03161781556200004562000553565b604051809103906000f08015156200005c57600080fd5b60068054600160a060020a031916600160a060020a03929092169190911790556200008662000564565b604051809103906000f08015156200009d57600080fd5b60078054600160a060020a031916600160a060020a0392909216919091179055620000c762000575565b604051809103906000f0801515620000de57600080fd5b60088054600160a060020a031916600160a060020a03929092169190911790556200011b60006186816401000000006200020d8102620006bb1704565b600754909150600160a060020a031663fa74efc13060a060405190810160405280858152602001858152602001858152602001858152602001858152506000604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a03831660048201908152906024018260a080838360005b83811015620001c4578082015183820152602001620001aa565b5050505090500192505050602060405180830381600087803b1515620001e957600080fd5b6102c65a03f11515620001fb57600080fd5b50505060405180519050505062000620565b60008054819074010000000000000000000000000000000000000000900460ff16156200023957600080fd5b30600160a060020a031684600160a060020a0316141515156200025b57600080fd5b620002776000848664010000000062000a906200027f82021704565b949350505050565b60006200028b62000586565b600060c060405190810160405280426001604060020a0316815260200160006001604060020a03168152602001876001608060020a0316815260200160006001608060020a0316815260200160006001608060020a03168152602001868152509150600160028054806001018281620003059190620005bb565b600092835260209092208591600302018151815467ffffffffffffffff19166001604060020a0391909116178155602082015181546001604060020a03919091166801000000000000000002604060020a608060020a0319909116178155604082015181546001608060020a0391821670010000000000000000000000000000000002911617815560608201516001820180546001608060020a0319166001608060020a039290921691909117905560808201516001820180546001608060020a0392831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146200040757600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526001608060020a031660408084019190915260608301919091526080909101905180910390a16200048e6000858364010000000062000d056200049782021704565b95945050505050565b600160a060020a03808316600081815260056020908152604080832080546001019055858352600390915290208054600160a060020a0319169091179055831615620004fe57600160a060020a038316600090815260056020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60405161077f80620014cc83390190565b6040516107b58062001c4b83390190565b604051610388806200240083390190565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511620005ea57600302816003028360005260206000209182019101620005ea9190620005ef565b505050565b6200061d91905b8082111562000619576000808255600182018190556002820155600301620005f6565b5090565b90565b610e9c80620006306000396000f300606060405236156101465763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461016357806312c2debf146101ed57806313af40351461022157806318160ddd146102405780632109f89d14610253578063302597201461027d5780633f4ba83a1461029357806345812958146102a65780635c975abb146102b95780635de038b1146102cc5780636352211e146102ee578063685b7a5e146103205780636af04a571461033657806370a0823114610349578063715879881461036857806378d0df50146103875780638456cb591461039d5780638462151c146103b05780638f84aa091461042257806395d89b41146104355780639db0225714610448578063a9059cbb14610478578063d78d16b61461049a578063e01196ef146104ad578063e77dad5c146104c0575b60065433600160a060020a0390811691161461016157600080fd5b005b341561016e57600080fd5b6101766104d3565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101b257808201518382015260200161019a565b50505050905090810190601f1680156101df5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101f857600080fd5b61020f600160a060020a036004351660243561050a565b60405190815260200160405180910390f35b341561022c57600080fd5b610161600160a060020a0360043516610538565b341561024b57600080fd5b61020f610597565b341561025e57600080fd5b61026960043561059e565b604051901515815260200160405180910390f35b341561028857600080fd5b610161600435610617565b341561029e57600080fd5b61016161062d565b34156102b157600080fd5b61020f610680565b34156102c457600080fd5b6102696106ab565b34156102d757600080fd5b61020f600160a060020a03600435166024356106bb565b34156102f957600080fd5b61030460043561070a565b604051600160a060020a03909116815260200160405180910390f35b341561032b57600080fd5b61026960043561072e565b341561034157600080fd5b61030461078d565b341561035457600080fd5b61020f600160a060020a036004351661079c565b341561037357600080fd5b610161600160a060020a03600435166107b7565b341561039257600080fd5b610304600435610852565b34156103a857600080fd5b61016161086d565b34156103bb57600080fd5b6103cf600160a060020a03600435166108c5565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561040e5780820151838201526020016103f6565b505050509050019250505060405180910390f35b341561042d57600080fd5b6103046109a6565b341561044057600080fd5b6101766109b5565b341561045357600080fd5b61045b6109ec565b60405167ffffffffffffffff909116815260200160405180910390f35b341561048357600080fd5b610161600160a060020a03600435166024356109f2565b34156104a557600080fd5b610304610a63565b34156104b857600080fd5b610304610a72565b34156104cb57600080fd5b610304610a81565b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b60046020528160005260406000208181548110151561052557fe5b6000918252602090912001549150829050565b60005433600160a060020a0390811691161461055357600080fd5b600160a060020a038116151561056857600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002545b90565b6000805460a060020a900460ff16156105b657600080fd5b6105bf82610617565b4267ffffffffffffffff166002838154811015156105d957fe5b600091825260209091206003909102015468010000000000000000900467ffffffffffffffff16111561060e57506000610612565b5060015b919050565b61061f610597565b811061062a57600080fd5b50565b60005433600160a060020a0390811691161461064857600080fd5b60005460a060020a900460ff16151561066057600080fd5b6000805474ff000000000000000000000000000000000000000019169055565b6000805460a060020a900460ff161561069857600080fd5b6106a660006201e2406106bb565b905090565b60005460a060020a900460ff1681565b60008054819060a060020a900460ff16156106d557600080fd5b30600160a060020a031684600160a060020a0316141515156106f657600080fd5b61070260008486610a90565b949350505050565b600081815260036020526040902054600160a060020a031680151561061257600080fd5b600061073982610617565b610e10420160028381548110151561074d57fe5b906000526020600020906003020160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060019050919050565b600154600160a060020a031681565b600160a060020a031660009081526005602052604090205490565b60005433600160a060020a039081169116146107d257600080fd5b60005460a060020a900460ff1615156107ea57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b600360205260009081526040902054600160a060020a031681565b60005433600160a060020a0390811691161461088857600080fd5b60005460a060020a900460ff161561089f57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a179055565b6108cd610dcd565b60006108d7610dcd565b60008060006108e58761079c565b94508415156109155760006040518059106108fd5750595b9080825280602002602001820160405250955061099c565b846040518059106109235750595b9080825280602002602001820160405250935061093e610597565b925060009150600190505b82811161099857600081815260036020526040902054600160a060020a0388811691161415610990578084838151811061097f57fe5b602090810290910101526001909101905b600101610949565b8395505b5050505050919050565b600054600160a060020a031681565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b610e1081565b60005460a060020a900460ff1615610a0957600080fd5b600160a060020a0382161515610a1e57600080fd5b30600160a060020a031682600160a060020a031614151515610a3f57600080fd5b610a493382610ce5565b1515610a5457600080fd5b610a5f338383610d05565b5050565b600854600160a060020a031690565b600654600160a060020a031690565b600754600160a060020a031690565b6000610a9a610ddf565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff168152602001868152509150600160028054806001018281610b2f9190610e14565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff81168114610c5c57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a1610cdc60008583610d05565b95945050505050565b600090815260036020526040902054600160a060020a0391821691161490565b600160a060020a0380831660008181526005602090815260408083208054600101905585835260039091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055831615610d7857600160a060020a038316600090815260056020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b815481835581811511610e4057600302816003028360005260206000209182019101610e409190610e45565b505050565b61059b91905b80821115610e6c576000808255600182018190556002820155600301610e4b565b50905600a165627a7a723058209d8655ba6bf837a8beef7b4fe6571df26641c7991d3aaca3f74c85a306da09e000296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a0316179055610746806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461003957806363c785aa14610068578063df73c49f1461008d57600080fd5b341561004457600080fd5b61004c6100a6565b604051600160a060020a03909116815260200160405180910390f35b341561007357600080fd5b61007b6100b5565b60405190815260200160405180910390f35b341561009857600080fd5b61007b6004356024356100bc565b600154600160a060020a031681565b6000545b90565b60006100c661066a565b6000806000806100d588610248565b15806100e757506100e587610248565b155b156100f5576000955061023d565b60a0604051908101604052804267ffffffffffffffff168152602001898152602001888152602001600081526020016000815250945060016000805480600101828161014191906106a4565b600092835260209092208891600502018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518160010155604082015181600201556060820151816003015560808201516004909101555003935063ffffffff841684146101af57600080fd5b6101b88561031b565b506101c384866104c3565b606087018281526080880182905291945092506101e09051610514565b90507f78d72c60b0f38a9a8dac0805b805c5f038ede4b314e696b2caefc2161058ffa58486606001518760800151846040518085815260200184815260200183815260200182815260200194505050505060405180910390a18395505b505050505092915050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029257600080fd5b6102c65a03f115156102a357600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102fb57600080fd5b6102c65a03f1151561030c57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561036557600080fd5b6102c65a03f1151561037657600080fd5b50505060405180519050600160a060020a031663b42cb37d836020015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156103d257600080fd5b6102c65a03f115156103e357600080fd5b50505060405180515050600154600160a060020a031663e77dad5c6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561043557600080fd5b6102c65a03f1151561044657600080fd5b50505060405180519050600160a060020a031663b42cb37d836040015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156104a257600080fd5b6102c65a03f115156104b357600080fd5b5050506040518051905050919050565b60008060006002846020015185604001518701018115156104e057fe5b0690508015156104fd57836020015184604001519250925061050c565b83604001518460200151925092505b509250929050565b6001546000908190600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561056057600080fd5b6102c65a03f1151561057157600080fd5b50505060405180519050600160a060020a0316636352211e8460006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156105c957600080fd5b6102c65a03f115156105da57600080fd5b5050506040518051600154909250600160a060020a03169050635de038b1826201b20760006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561064957600080fd5b6102c65a03f1151561065a57600080fd5b5050506040518051949350505050565b60a060405190810160405280600067ffffffffffffffff168152602001600081526020016000815260200160008152602001600081525090565b8154818355818115116106d0576005028160050283600052602060002091820191016106d091906106d5565b505050565b6100b991905b8082111561071657805467ffffffffffffffff19168155600060018201819055600282018190556003820181905560048201556005016106db565b50905600a165627a7a72305820386a2b1be06d6b92619bba00bb6e6b79859934336616994788f50c4d2935e99300296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561077c806100396000396000f300606060405236156100675763ffffffff60e060020a600035041663624fe81f811461006c5780636352211e1461009b5780637bf13d82146100b1578063b42cb37d146100d6578063bfeea70814610100578063ecf9036b14610113578063fa74efc114610129575b600080fd5b341561007757600080fd5b61007f61014a565b604051600160a060020a03909116815260200160405180910390f35b34156100a657600080fd5b61007f600435610159565b34156100bc57600080fd5b6100c46101b5565b60405190815260200160405180910390f35b34156100e157600080fd5b6100ec6004356101bc565b604051901515815260200160405180910390f35b341561010b57600080fd5b6100c4610317565b341561011e57600080fd5b6100ec60043561031c565b341561013457600080fd5b6100c4600160a060020a0360043516602461047a565b600154600160a060020a031681565b60006101636101b5565b821061016e57600080fd5b600080548390811061017c57fe5b6000918252602090912060069091020154680100000000000000009004600160a060020a031690508015156101b057600080fd5b919050565b6000545b90565b60006101c661060b565b60006101d06101b5565b84106101db57600080fd5b60008054859081106101e957fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b81548152602001906001019080831161024f5750505091909252509193506000925050505b816040015150600581101561030b57600154600160a060020a031663685b7a5e604084015183600581106102a457fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b15156102e857600080fd5b6102c65a03f115156102f957600080fd5b50505060405180515050600101610274565b600192505b5050919050565b600581565b600061032661060b565b60006103306101b5565b841061033b57600080fd5b600080548590811061034957fe5b906000526020600020906006020160606040519081016040908152825467ffffffffffffffff81168352680100000000000000009004600160a060020a0316602083015290919080830190600183019060059060a09051908101604052919060a0830182845b8154815260200190600101908083116103af5750505091909252509193506000925050505b816040015150600581101561030b57600154600160a060020a0316632109f89d6040840151836005811061040457fe5b602002015160006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561044857600080fd5b6102c65a03f1151561045957600080fd5b5050506040518051905015156104725760009250610310565b6001016103d4565b600061048461060b565b60006060604051908101604052804267ffffffffffffffff16815260200186600160a060020a03168152602001856005806020026040519081016040529190828260a080828437505050919092525050600080549193506001918083016104eb8382610631565b600092835260209092208591600602018151815467ffffffffffffffff191667ffffffffffffffff9190911617815560208201518154600160a060020a039190911668010000000000000000027fffffffff0000000000000000000000000000000000000000ffffffffffffffff90911617815560408201516105749060018301906005610662565b50505003905063ffffffff8116811461058c57600080fd5b7ff7841da7904048ca49ded1df3a41ff46907a4db7880f86ebde13572a7154f59d85828460400151604051600160a060020a038416815260208101839052604081018260a080838360005b838110156105ef5780820151838201526020016105d7565b50505050905001935050505060405180910390a1949350505050565b60e0604051908101604090815260008083526020830152810161062c6106a0565b905290565b81548183558181151161065d5760060281600602836000526020600020918201910161065d91906106c7565b505050565b8260058101928215610690579160200282015b82811115610690578251825591602001919060010190610675565b5061069c929150610713565b5090565b60a06040519081016040526005815b60008152602001906001900390816106af5790505090565b6101b991905b8082111561069c5780547fffffffff00000000000000000000000000000000000000000000000000000000168155600061070a600183018261072d565b506006016106cd565b6101b991905b8082111561069c5760008155600101610719565b5060008155600101600081556001016000815560010160008155600101600090555600a165627a7a7230582004abbec0c641e4cf752f6cfe4fb5293843bee15a71e9ff11f738d37f4cbe16dd00296060604052341561000f57600080fd5b60018054600160a060020a03191633600160a060020a031617905561034f806100396000396000f300606060405263ffffffff60e060020a600035041663624fe81f811461002e578063969494201461005d57600080fd5b341561003957600080fd5b610041610087565b604051600160a060020a03909116815260200160405180910390f35b341561006857600080fd5b610073600435610096565b604051901515815260200160405180910390f35b600154600160a060020a031681565b60006100a182610176565b15156100af57506000610171565b7f4459da58fdcdad9b33fef29d7f321bc7dae89ee030117887627a115627703ef7824260405191825267ffffffffffffffff1660208201526040908101905180910390a16000541515610109575060008190556001610171565b7f9ede9f883c5c5099d0d01328ba10287c2ef97ebb8c272e30cd3d40cef113a0548260005442604051928352602083019190915267ffffffffffffffff166040808301919091526060909101905180910390a161016882600054610249565b50506000805560015b919050565b600154600090600160a060020a031663e77dad5c82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156101c057600080fd5b6102c65a03f115156101d157600080fd5b50505060405180519050600160a060020a031663ecf9036b8360006040516020015260405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561022957600080fd5b6102c65a03f1151561023a57600080fd5b50505060405180519392505050565b600154600090600160a060020a031663e01196ef82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561029357600080fd5b6102c65a03f115156102a457600080fd5b50505060405180519050600160a060020a031663df73c49f848460006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b151561030257600080fd5b6102c65a03f1151561031357600080fd5b50505060405180519493505050505600a165627a7a72305820559e2597a25686984f19672b4062de93ef47b200e6fd5ac6a21e9722458032160029`

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

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CryptoCardsCore.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreSession) Paused() (bool, error) {
	return _CryptoCardsCore.Contract.Paused(&_CryptoCardsCore.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_CryptoCardsCore *CryptoCardsCoreCallerSession) Paused() (bool, error) {
	return _CryptoCardsCore.Contract.Paused(&_CryptoCardsCore.CallOpts)
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

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) Pause() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Pause(&_CryptoCardsCore.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) Pause() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Pause(&_CryptoCardsCore.TransactOpts)
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

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetOwner(&_CryptoCardsCore.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(_newOwner address) returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.SetOwner(&_CryptoCardsCore.TransactOpts, _newOwner)
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

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoCardsCore.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CryptoCardsCore *CryptoCardsCoreSession) Unpause() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Unpause(&_CryptoCardsCore.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CryptoCardsCore *CryptoCardsCoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _CryptoCardsCore.Contract.Unpause(&_CryptoCardsCore.TransactOpts)
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
