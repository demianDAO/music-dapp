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

// SongNFTTradeSongDetails is an auto generated low-level Go binding around an user-defined struct.
type SongNFTTradeSongDetails struct {
	TokenId  *big.Int
	TokenURI string
	Balance  *big.Int
	Price    *big.Int
}

// SongNFTTradeMetaData contains all meta data concerning the SongNFTTrade contract.
var SongNFTTradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReleasedSong\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"SongPurchased\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getSongInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structSongNFTTrade.SongDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractSongNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"singer\",\"type\":\"address\"}],\"name\":\"purchaseSong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"releasedSong\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"songPricesByAddr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractMPToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetSongInfos is a free data retrieval call binding the contract method 0xc89e774e.
//
// Solidity: function getSongInfos(address user) view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeCaller) GetSongInfos(opts *bind.CallOpts, user common.Address) ([]SongNFTTradeSongDetails, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "getSongInfos", user)

	if err != nil {
		return *new([]SongNFTTradeSongDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]SongNFTTradeSongDetails)).(*[]SongNFTTradeSongDetails)

	return out0, err

}

// GetSongInfos is a free data retrieval call binding the contract method 0xc89e774e.
//
// Solidity: function getSongInfos(address user) view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeSession) GetSongInfos(user common.Address) ([]SongNFTTradeSongDetails, error) {
	return _SongNFTTrade.Contract.GetSongInfos(&_SongNFTTrade.CallOpts, user)
}

// GetSongInfos is a free data retrieval call binding the contract method 0xc89e774e.
//
// Solidity: function getSongInfos(address user) view returns((uint256,string,uint256,uint256)[])
func (_SongNFTTrade *SongNFTTradeCallerSession) GetSongInfos(user common.Address) ([]SongNFTTradeSongDetails, error) {
	return _SongNFTTrade.Contract.GetSongInfos(&_SongNFTTrade.CallOpts, user)
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

// SongPricesByAddr is a free data retrieval call binding the contract method 0xd6fc9f89.
//
// Solidity: function songPricesByAddr(address , uint256 ) view returns(uint256)
func (_SongNFTTrade *SongNFTTradeCaller) SongPricesByAddr(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SongNFTTrade.contract.Call(opts, &out, "songPricesByAddr", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SongPricesByAddr is a free data retrieval call binding the contract method 0xd6fc9f89.
//
// Solidity: function songPricesByAddr(address , uint256 ) view returns(uint256)
func (_SongNFTTrade *SongNFTTradeSession) SongPricesByAddr(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SongNFTTrade.Contract.SongPricesByAddr(&_SongNFTTrade.CallOpts, arg0, arg1)
}

// SongPricesByAddr is a free data retrieval call binding the contract method 0xd6fc9f89.
//
// Solidity: function songPricesByAddr(address , uint256 ) view returns(uint256)
func (_SongNFTTrade *SongNFTTradeCallerSession) SongPricesByAddr(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SongNFTTrade.Contract.SongPricesByAddr(&_SongNFTTrade.CallOpts, arg0, arg1)
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

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 tokenId, address singer) returns()
func (_SongNFTTrade *SongNFTTradeTransactor) PurchaseSong(opts *bind.TransactOpts, tokenId *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.contract.Transact(opts, "purchaseSong", tokenId, singer)
}

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 tokenId, address singer) returns()
func (_SongNFTTrade *SongNFTTradeSession) PurchaseSong(tokenId *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.PurchaseSong(&_SongNFTTrade.TransactOpts, tokenId, singer)
}

// PurchaseSong is a paid mutator transaction binding the contract method 0xbaa1c491.
//
// Solidity: function purchaseSong(uint256 tokenId, address singer) returns()
func (_SongNFTTrade *SongNFTTradeTransactorSession) PurchaseSong(tokenId *big.Int, singer common.Address) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.PurchaseSong(&_SongNFTTrade.TransactOpts, tokenId, singer)
}

// ReleasedSong is a paid mutator transaction binding the contract method 0x5aba30e3.
//
// Solidity: function releasedSong(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeTransactor) ReleasedSong(opts *bind.TransactOpts, amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.contract.Transact(opts, "releasedSong", amount, price, uri)
}

// ReleasedSong is a paid mutator transaction binding the contract method 0x5aba30e3.
//
// Solidity: function releasedSong(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeSession) ReleasedSong(amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.ReleasedSong(&_SongNFTTrade.TransactOpts, amount, price, uri)
}

// ReleasedSong is a paid mutator transaction binding the contract method 0x5aba30e3.
//
// Solidity: function releasedSong(uint256 amount, uint256 price, string uri) returns()
func (_SongNFTTrade *SongNFTTradeTransactorSession) ReleasedSong(amount *big.Int, price *big.Int, uri string) (*types.Transaction, error) {
	return _SongNFTTrade.Contract.ReleasedSong(&_SongNFTTrade.TransactOpts, amount, price, uri)
}

// SongNFTTradeReleasedSongIterator is returned from FilterReleasedSong and is used to iterate over the raw logs and unpacked data for ReleasedSong events raised by the SongNFTTrade contract.
type SongNFTTradeReleasedSongIterator struct {
	Event *SongNFTTradeReleasedSong // Event containing the contract specifics and raw log

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
func (it *SongNFTTradeReleasedSongIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SongNFTTradeReleasedSong)
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
		it.Event = new(SongNFTTradeReleasedSong)
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
func (it *SongNFTTradeReleasedSongIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SongNFTTradeReleasedSongIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SongNFTTradeReleasedSong represents a ReleasedSong event raised by the SongNFTTrade contract.
type SongNFTTradeReleasedSong struct {
	Singer   common.Address
	Price    *big.Int
	TokenURI string
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReleasedSong is a free log retrieval operation binding the contract event 0xd454868b32bff89b28c358765c82e681ed1012c2bccbd81e2e7d406d3bae687a.
//
// Solidity: event ReleasedSong(address indexed singer, uint256 price, string tokenURI, uint256 tokenId, uint256 amount)
func (_SongNFTTrade *SongNFTTradeFilterer) FilterReleasedSong(opts *bind.FilterOpts, singer []common.Address) (*SongNFTTradeReleasedSongIterator, error) {

	var singerRule []interface{}
	for _, singerItem := range singer {
		singerRule = append(singerRule, singerItem)
	}

	logs, sub, err := _SongNFTTrade.contract.FilterLogs(opts, "ReleasedSong", singerRule)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeReleasedSongIterator{contract: _SongNFTTrade.contract, event: "ReleasedSong", logs: logs, sub: sub}, nil
}

// WatchReleasedSong is a free log subscription operation binding the contract event 0xd454868b32bff89b28c358765c82e681ed1012c2bccbd81e2e7d406d3bae687a.
//
// Solidity: event ReleasedSong(address indexed singer, uint256 price, string tokenURI, uint256 tokenId, uint256 amount)
func (_SongNFTTrade *SongNFTTradeFilterer) WatchReleasedSong(opts *bind.WatchOpts, sink chan<- *SongNFTTradeReleasedSong, singer []common.Address) (event.Subscription, error) {

	var singerRule []interface{}
	for _, singerItem := range singer {
		singerRule = append(singerRule, singerItem)
	}

	logs, sub, err := _SongNFTTrade.contract.WatchLogs(opts, "ReleasedSong", singerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SongNFTTradeReleasedSong)
				if err := _SongNFTTrade.contract.UnpackLog(event, "ReleasedSong", log); err != nil {
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

// ParseReleasedSong is a log parse operation binding the contract event 0xd454868b32bff89b28c358765c82e681ed1012c2bccbd81e2e7d406d3bae687a.
//
// Solidity: event ReleasedSong(address indexed singer, uint256 price, string tokenURI, uint256 tokenId, uint256 amount)
func (_SongNFTTrade *SongNFTTradeFilterer) ParseReleasedSong(log types.Log) (*SongNFTTradeReleasedSong, error) {
	event := new(SongNFTTradeReleasedSong)
	if err := _SongNFTTrade.contract.UnpackLog(event, "ReleasedSong", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	TokenId *big.Int
	Buyer   common.Address
	Price   *big.Int
	Singer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSongPurchased is a free log retrieval operation binding the contract event 0x4c6b442b91275055375949d16e86a32fe7c913ef1e584669570e5074022ba63e.
//
// Solidity: event SongPurchased(uint256 indexed tokenId, address buyer, uint256 price, address singer)
func (_SongNFTTrade *SongNFTTradeFilterer) FilterSongPurchased(opts *bind.FilterOpts, tokenId []*big.Int) (*SongNFTTradeSongPurchasedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SongNFTTrade.contract.FilterLogs(opts, "SongPurchased", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SongNFTTradeSongPurchasedIterator{contract: _SongNFTTrade.contract, event: "SongPurchased", logs: logs, sub: sub}, nil
}

// WatchSongPurchased is a free log subscription operation binding the contract event 0x4c6b442b91275055375949d16e86a32fe7c913ef1e584669570e5074022ba63e.
//
// Solidity: event SongPurchased(uint256 indexed tokenId, address buyer, uint256 price, address singer)
func (_SongNFTTrade *SongNFTTradeFilterer) WatchSongPurchased(opts *bind.WatchOpts, sink chan<- *SongNFTTradeSongPurchased, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SongNFTTrade.contract.WatchLogs(opts, "SongPurchased", tokenIdRule)
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

// ParseSongPurchased is a log parse operation binding the contract event 0x4c6b442b91275055375949d16e86a32fe7c913ef1e584669570e5074022ba63e.
//
// Solidity: event SongPurchased(uint256 indexed tokenId, address buyer, uint256 price, address singer)
func (_SongNFTTrade *SongNFTTradeFilterer) ParseSongPurchased(log types.Log) (*SongNFTTradeSongPurchased, error) {
	event := new(SongNFTTradeSongPurchased)
	if err := _SongNFTTrade.contract.UnpackLog(event, "SongPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
