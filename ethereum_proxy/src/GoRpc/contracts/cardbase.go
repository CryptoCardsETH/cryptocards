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

// CardBaseABI is the input ABI used to generate the binding from.
const CardBaseABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardBaseBin is the compiled bytecode used for deploying new contracts.
const CardBaseBin = `0x6060604052341561000f57600080fd5b6101528061001e6000396000f30060606040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312c2debf811461005057806378d0df5014610091575b600080fd5b341561005b57600080fd5b61007f73ffffffffffffffffffffffffffffffffffffffff600435166024356100d0565b60405190815260200160405180910390f35b341561009c57600080fd5b6100a76004356100fe565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6002602052816000526040600020818154811015156100eb57fe5b6000918252602090912001549150829050565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058201c4c4d8369bd337e235fc64a7d901218fa7356965e3d9d2af66442482c66b6df0029`

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

// CardOwnershipABI is the input ABI used to generate the binding from.
const CardOwnershipABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardsHeldByOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_attributes\",\"type\":\"uint256\"}],\"name\":\"createCard\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cardIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_cardId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cardID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"creationBattleID\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"attributes\",\"type\":\"uint256\"}],\"name\":\"NewCard\",\"type\":\"event\"}]"

// CardOwnershipBin is the compiled bytecode used for deploying new contracts.
const CardOwnershipBin = `0x6060604052341561000f57600080fd5b61092d8061001e6000396000f3006060604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100a857806312c2debf1461013257806318160ddd146101665780635de038b1146101795780636352211e1461019b57806370a08231146101cd57806378d0df50146101ec5780638462151c1461020257806395d89b4114610274578063a9059cbb14610287575b600080fd5b34156100b357600080fd5b6100bb6102ab565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f75780820151838201526020016100df565b50505050905090810190601f1680156101245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561013d57600080fd5b610154600160a060020a03600435166024356102e2565b60405190815260200160405180910390f35b341561017157600080fd5b610154610310565b341561018457600080fd5b610154600160a060020a0360043516602435610317565b34156101a657600080fd5b6101b160043561034f565b604051600160a060020a03909116815260200160405180910390f35b34156101d857600080fd5b610154600160a060020a0360043516610378565b34156101f757600080fd5b6101b1600435610393565b341561020d57600080fd5b610221600160a060020a03600435166103ae565b60405160208082528190810183818151815260200191508051906020019060200280838360005b83811015610260578082015183820152602001610248565b505050509050019250505060405180910390f35b341561027f57600080fd5b6100bb61048f565b341561029257600080fd5b6102a9600160a060020a03600435166024356104c6565b005b60408051908101604052600b81527f43727970746f4361726473000000000000000000000000000000000000000000602082015281565b6002602052816000526040600020818154811015156102fd57fe5b6000918252602090912001549150829050565b6000545b90565b60008030600160a060020a031684600160a060020a03161415151561033b57600080fd5b61034760008486610520565b949350505050565b600081815260016020526040902054600160a060020a031680151561037357600080fd5b919050565b600160a060020a031660009081526003602052604090205490565b600160205260009081526040902054600160a060020a031681565b6103b661085e565b60006103c061085e565b60008060006103ce87610378565b94508415156103fe5760006040518059106103e65750595b90808252806020026020018201604052509550610485565b8460405180591061040c5750595b90808252806020026020018201604052509350610427610310565b925060009150600190505b82811161048157600081815260016020526040902054600160a060020a0388811691161415610479578084838151811061046857fe5b602090810290910101526001909101905b600101610432565b8395505b5050505050919050565b60408051908101604052600381527f4343420000000000000000000000000000000000000000000000000000000000602082015281565b600160a060020a03821615156104db57600080fd5b30600160a060020a031682600160a060020a0316141515156104fc57600080fd5b6105063382610775565b151561051157600080fd5b61051c338383610795565b5050565b600061052a610870565b600060c0604051908101604052804267ffffffffffffffff168152602001600067ffffffffffffffff168152602001876fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff16815260200160006fffffffffffffffffffffffffffffffff1681526020018681525091506001600080548060010182816105bf91906108a5565b600092835260209092208591600302018151815467ffffffffffffffff191667ffffffffffffffff919091161781556020820151815467ffffffffffffffff9190911668010000000000000000026fffffffffffffffff000000000000000019909116178155604082015181546fffffffffffffffffffffffffffffffff91821670010000000000000000000000000000000002911617815560608201516001820180546fffffffffffffffffffffffffffffffff19166fffffffffffffffffffffffffffffffff9290921691909117905560808201516001820180546fffffffffffffffffffffffffffffffff92831670010000000000000000000000000000000002921691909117905560a08201516002909101555003905063ffffffff811681146106ec57600080fd5b7fc56570397f0bcf235fc72dc30f5c4d9c77206bfce1443759a8ee3a19dcd196e4848284604001518560a00151604051600160a060020a03909416845260208401929092526fffffffffffffffffffffffffffffffff1660408084019190915260608301919091526080909101905180910390a161076c60008583610795565b95945050505050565b600090815260016020526040902054600160a060020a0391821691161490565b600160a060020a038083166000818152600360209081526040808320805460019081019091558684529091529020805473ffffffffffffffffffffffffffffffffffffffff1916909117905583161561080957600160a060020a038316600090815260036020526040902080546000190190555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60206040519081016040526000815290565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116108d1576003028160030283600052602060002091820191016108d191906108d6565b505050565b61031491905b808211156108fd5760008082556001820181905560028201556003016108dc565b50905600a165627a7a72305820dfcf4962f4af27ace33a9c196ac5bb9934714dc82a52daceed6f84274037e73f0029`

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
