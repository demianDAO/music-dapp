// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sm

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SongNFTTradeMusicDetails is an auto generated low-level Go binding around an user-defined struct.
type SongNFTTradeMusicDetails struct {
	TokenId  *big.Int
	TokenURI string
	Balance  *big.Int
	Price    *big.Int
}

// SongNFTTradeMetaData contains all meta data concerning the SongNFTTrade contract.
var SongNFTTradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMPToken\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractSongNFT\",\"name\":\"_nft\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SongPurchased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"CreateMusic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMusicInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structSongNFTTrade.MusicDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"musicInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractSongNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"purchaseSong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractMPToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SongNFTTradeABI is the input ABI used to generate the binding from.
// Deprecated: Use SongNFTTradeMetaData.ABI instead.
var SongNFTTradeABI = SongNFTTradeMetaData.ABI

// SongNFTTrade is an auto generated Go binding around an Ethereum contract.
type SongNFTTrade struct {
	SongNFTTradeCaller     // Read-only binding to the contract
	SongNFTTradeTransactor // Write-only binding to the contract
	SongNFTTradeFilterer   // Log filterer for contract events
}

// SongNFTTradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SongNFTTradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SongNFTTradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SongNFTTradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SongNFTTradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SongNFTTradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SongNFTTradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SongNFTTradeSession struct {
	Contract     *SongNFTTrade     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SongNFTTradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SongNFTTradeCallerSession struct {
	Contract *SongNFTTradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SongNFTTradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SongNFTTradeTransactorSession struct {
	Contract     *SongNFTTradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SongNFTTradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SongNFTTradeRaw struct {
	Contract *SongNFTTrade // Generic contract binding to access the raw methods on
}

// SongNFTTradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SongNFTTradeCallerRaw struct {
	Contract *SongNFTTradeCaller // Generic read-only contract binding to access the raw methods on
}

// SongNFTTradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SongNFTTradeTransactorRaw struct {
	Contract *SongNFTTradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSongNFTTrade creates a new instance of SongNFTTrade, bound to a specific deployed contract.
func NewSongNFTTrade(address common.Address, backend bind.ContractBackend) (*SongNFTTrade, error) {
	contract, err := bindSongNFTTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SongNFTTrade{SongNFTTradeCaller: SongNFTTradeCaller{contract: contract}, SongNFTTradeTransactor: SongNFTTradeTransactor{contract: contract}, SongNFTTradeFilterer: SongNFTTradeFilterer{contract: contract}}, nil
}

// NewSongNFTTradeCaller creates a new read-only instance of SongNFTTrade, bound to a specific deployed contract.
func NewSongNFTTradeCaller(address common.Address, caller bind.ContractCaller) (*SongNFTTradeCaller, error) {
	contract, err := bindSongNFTTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeCaller{contract: contract}, nil
}

// NewSongNFTTradeTransactor creates a new write-only instance of SongNFTTrade, bound to a specific deployed contract.
func NewSongNFTTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*SongNFTTradeTransactor, error) {
	contract, err := bindSongNFTTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeTransactor{contract: contract}, nil
}

// NewSongNFTTradeFilterer creates a new log filterer instance of SongNFTTrade, bound to a specific deployed contract.
func NewSongNFTTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*SongNFTTradeFilterer, error) {
	contract, err := bindSongNFTTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeFilterer{contract: contract}, nil
}

// bindSongNFTTrade binds a generic wrapper to an already deployed contract.
func bindSongNFTTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SongNFTTradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SongNFTTrade *SongNFTTradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SongNFTTrade.Contract.SongNFTTradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SongNFTTrade *SongNFTTradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.SongNFTTradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SongNFTTrade *SongNFTTradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.SongNFTTradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SongNFTTrade *SongNFTTradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SongNFTTrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SongNFTTrade *SongNFTTradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SongNFTTrade *SongNFTTradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.contract.Transact(opts, method, params...)
}

// GetMusicInfos is a free data retrieval call binding the contract method 0xf4e3ddc2.
//
// Solidity: function getMusicInfos() view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeCaller) GetMusicInfos(opts *bind.CallOpts) ([]SongNFTTradeMusicDetails, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "getMusicInfos")

	if err != nil {
		return *new([]SongNFTTradeMusicDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]SongNFTTradeMusicDetails)).(*[]SongNFTTradeMusicDetails)

	return out0, err

}

// GetMusicInfos is a free data retrieval call binding the contract method 0xf4e3ddc2.
//
// Solidity: function getMusicInfos() view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeSession) GetMusicInfos() ([]SongNFTTradeMusicDetails, error) {
	return _SongNFTTrade.Contract.GetMusicInfos(&_SongNFTTrade.CallOpts)
}

// GetMusicInfos is a free data retrieval call binding the contract method 0xf4e3ddc2.
//
// Solidity: function getMusicInfos() view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeCallerSession) GetMusicInfos() ([]SongNFTTradeMusicDetails, error) {
	return _SongNFTTrade.Contract.GetMusicInfos(&_SongNFTTrade.CallOpts)
}

// MusicInfos is a free data retrieval call binding the contract method 0x22d281be.
//
// Solidity: function musicInfos(address , uint256 ) view returns(uint256 price, uint256 tokenId)
func (_SongNFTTrade *SongNFTTradeCaller) MusicInfos(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Price   *big.Int
	TokenId *big.Int
}, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "musicInfos", arg0, arg1)

	outstruct := new(struct {
		Price   *big.Int
		TokenId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MusicInfos is a free data retrieval call binding the contract method 0x22d281be.
//
// Solidity: function musicInfos(address , uint256 ) view returns(uint256 price, uint256 tokenId)
func (_SongNFTTrade *SongNFTTradeSession) MusicInfos(arg0 common.Address, arg1 *big.Int) (struct {
	Price   *big.Int
	TokenId *big.Int
}, error) {
	return _SongNFTTrade.Contract.MusicInfos(&_SongNFTTrade.CallOpts, arg0, arg1)
}

// MusicInfos is a free data retrieval call binding the contract method 0x22d281be.
//
// Solidity: function musicInfos(address , uint256 ) view returns(uint256 price, uint256 tokenId)
func (_SongNFTTrade *SongNFTTradeCallerSession) MusicInfos(arg0 common.Address, arg1 *big.Int) (struct {
	Price   *big.Int
	TokenId *big.Int
}, error) {
	return _SongNFTTrade.Contract.MusicInfos(&_SongNFTTrade.CallOpts, arg0, arg1)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_SongNFTTrade *SongNFTTradeCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_SongNFTTrade *SongNFTTradeSession) Nft() (common.Address, error) {
	return _SongNFTTrade.Contract.Nft(&_SongNFTTrade.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_SongNFTTrade *SongNFTTradeCallerSession) Nft() (common.Address, error) {
	return _SongNFTTrade.Contract.Nft(&_SongNFTTrade.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SongNFTTrade *SongNFTTradeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SongNFTTrade *SongNFTTradeSession) Token() (common.Address, error) {
	return _SongNFTTrade.Contract.Token(&_SongNFTTrade.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SongNFTTrade *SongNFTTradeCallerSession) Token() (common.Address, error) {
	return _SongNFTTrade.Contract.Token(&_SongNFTTrade.CallOpts)
}

// CreateMusic is a paid mutator transaction binding the contract method 0xa907b35c.
//
// Solidity: function CreateMusic(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeTransactor) CreateMusic(opts *bind.TransactOpts, amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.contract.Transact(opts, "CreateMusic", amount, price, uri)
}

// CreateMusic is a paid mutator transaction binding the contract method 0xa907b35c.
//
// Solidity: function CreateMusic(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeSession) CreateMusic(amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.CreateMusic(&_SongNFTTrade.TransactOpts, amount, price, uri)
}

// CreateMusic is a paid mutator transaction binding the contract method 0xa907b35c.
//
// Solidity: function CreateMusic(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeTransactorSession) CreateMusic(amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.CreateMusic(&_SongNFTTrade.TransactOpts, amount, price, uri)
}

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 id, address singer) returns()
func (_SongNFTTrade *SongNFTTradeTransactor) PurchaseSong(opts *bind.TransactOpts, id *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.contract.Transact(opts, "purchaseSong", id, singer)
}

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 id, address singer) returns()
func (_SongNFTTrade *SongNFTTradeSession) PurchaseSong(id *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.PurchaseSong(&_SongNFTTrade.TransactOpts, id, singer)
}

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 id, address singer) returns()
func (_SongNFTTrade *SongNFTTradeTransactorSession) PurchaseSong(id *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.PurchaseSong(&_SongNFTTrade.TransactOpts, id, singer)
}

// SongNFTTradeSongPurchasedIterator is returned from FilterSongPurchased and is used to iterate over the raw logs and unpacked data for SongPurchased events raised by the SongNFTTrade contract.
type SongNFTTradeSongPurchasedIterator struct {
	Event *SongNFTTradeSongPurchased // Event containing the contract specifics and raw log

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
func (it *SongNFTTradeSongPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SongNFTTradeSongPurchased)
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
		it.Event = new(SongNFTTradeSongPurchased)
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
func (it *SongNFTTradeSongPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SongNFTTradeSongPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SongNFTTradeSongPurchased represents a SongPurchased event raised by the SongNFTTrade contract.
type SongNFTTradeSongPurchased struct {
	Id    *big.Int
	Buyer common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSongPurchased is a free log retrieval operation binding the contract event 0x29d26227489513ba0617affd158f71dfd9f7c32d002e58a4393a98b8608d6a13.
//
// Solidity: event SongPurchased(uint256 indexed id, address buyer, uint256 price)
func (_SongNFTTrade *SongNFTTradeFilterer) FilterSongPurchased(opts *bind.FilterOpts, id []*big.Int) (*SongNFTTradeSongPurchasedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SongNFTTrade.contract.FilterLogs(opts, "SongPurchased", idRule)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeSongPurchasedIterator{contract: _SongNFTTrade.contract, event: "SongPurchased", logs: logs, sub: sub}, nil
}

// WatchSongPurchased is a free log subscription operation binding the contract event 0x29d26227489513ba0617affd158f71dfd9f7c32d002e58a4393a98b8608d6a13.
//
// Solidity: event SongPurchased(uint256 indexed id, address buyer, uint256 price)
func (_SongNFTTrade *SongNFTTradeFilterer) WatchSongPurchased(opts *bind.WatchOpts, sink chan<- *SongNFTTradeSongPurchased, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SongNFTTrade.contract.WatchLogs(opts, "SongPurchased", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SongNFTTradeSongPurchased)
				if err := _SongNFTTrade.contract.UnpackLog(event, "SongPurchased", log); err != nil {
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

// ParseSongPurchased is a log parse operation binding the contract event 0x29d26227489513ba0617affd158f71dfd9f7c32d002e58a4393a98b8608d6a13.
//
// Solidity: event SongPurchased(uint256 indexed id, address buyer, uint256 price)
func (_SongNFTTrade *SongNFTTradeFilterer) ParseSongPurchased(log types.Log) (*SongNFTTradeSongPurchased, error) {
	event := new(SongNFTTradeSongPurchased)
	if err := _SongNFTTrade.contract.UnpackLog(event, "SongPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
