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

// MusicNFTMgrTrackListing is an auto generated low-level Go binding around an user-defined struct.
type MusicNFTMgrTrackListing struct {
	TrackId     *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Artist      common.Address
	Owner       common.Address
}

// MusicNFTMgrMetaData contains all meta data concerning the MusicNFTMgr contract.
var MusicNFTMgrMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trackId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"artist\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TrackListingCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"createTrackListing\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchTrackListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"trackId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"artist\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structMusicNFTMgr.TrackListing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumTrackListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTrackListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trackListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MusicNFTMgrABI is the input ABI used to generate the binding from.
// Deprecated: Use MusicNFTMgrMetaData.ABI instead.
var MusicNFTMgrABI = MusicNFTMgrMetaData.ABI

// MusicNFTMgr is an auto generated Go binding around an Ethereum contract.
type MusicNFTMgr struct {
	MusicNFTMgrCaller     // Read-only binding to the contract
	MusicNFTMgrTransactor // Write-only binding to the contract
	MusicNFTMgrFilterer   // Log filterer for contract events
}

// MusicNFTMgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type MusicNFTMgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusicNFTMgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MusicNFTMgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusicNFTMgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MusicNFTMgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MusicNFTMgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MusicNFTMgrSession struct {
	Contract     *MusicNFTMgr      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MusicNFTMgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MusicNFTMgrCallerSession struct {
	Contract *MusicNFTMgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MusicNFTMgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MusicNFTMgrTransactorSession struct {
	Contract     *MusicNFTMgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MusicNFTMgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type MusicNFTMgrRaw struct {
	Contract *MusicNFTMgr // Generic contract binding to access the raw methods on
}

// MusicNFTMgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MusicNFTMgrCallerRaw struct {
	Contract *MusicNFTMgrCaller // Generic read-only contract binding to access the raw methods on
}

// MusicNFTMgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MusicNFTMgrTransactorRaw struct {
	Contract *MusicNFTMgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMusicNFTMgr creates a new instance of MusicNFTMgr, bound to a specific deployed contract.
func NewMusicNFTMgr(address common.Address, backend bind.ContractBackend) (*MusicNFTMgr, error) {
	contract, err := bindMusicNFTMgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MusicNFTMgr{MusicNFTMgrCaller: MusicNFTMgrCaller{contract: contract}, MusicNFTMgrTransactor: MusicNFTMgrTransactor{contract: contract}, MusicNFTMgrFilterer: MusicNFTMgrFilterer{contract: contract}}, nil
}

// NewMusicNFTMgrCaller creates a new read-only instance of MusicNFTMgr, bound to a specific deployed contract.
func NewMusicNFTMgrCaller(address common.Address, caller bind.ContractCaller) (*MusicNFTMgrCaller, error) {
	contract, err := bindMusicNFTMgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MusicNFTMgrCaller{contract: contract}, nil
}

// NewMusicNFTMgrTransactor creates a new write-only instance of MusicNFTMgr, bound to a specific deployed contract.
func NewMusicNFTMgrTransactor(address common.Address, transactor bind.ContractTransactor) (*MusicNFTMgrTransactor, error) {
	contract, err := bindMusicNFTMgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MusicNFTMgrTransactor{contract: contract}, nil
}

// NewMusicNFTMgrFilterer creates a new log filterer instance of MusicNFTMgr, bound to a specific deployed contract.
func NewMusicNFTMgrFilterer(address common.Address, filterer bind.ContractFilterer) (*MusicNFTMgrFilterer, error) {
	contract, err := bindMusicNFTMgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MusicNFTMgrFilterer{contract: contract}, nil
}

// bindMusicNFTMgr binds a generic wrapper to an already deployed contract.
func bindMusicNFTMgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MusicNFTMgrMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusicNFTMgr *MusicNFTMgrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MusicNFTMgr.Contract.MusicNFTMgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusicNFTMgr *MusicNFTMgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.MusicNFTMgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusicNFTMgr *MusicNFTMgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.MusicNFTMgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MusicNFTMgr *MusicNFTMgrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MusicNFTMgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MusicNFTMgr *MusicNFTMgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MusicNFTMgr *MusicNFTMgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.contract.Transact(opts, method, params...)
}

// FetchTrackListings is a free data retrieval call binding the contract method 0x0e543f4a.
//
// Solidity: function fetchTrackListings() view returns((uint256,address,uint256,address,address)[])
func (_MusicNFTMgr *MusicNFTMgrCaller) FetchTrackListings(opts *bind.CallOpts) ([]MusicNFTMgrTrackListing, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "fetchTrackListings")

	if err != nil {
		return *new([]MusicNFTMgrTrackListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MusicNFTMgrTrackListing)).(*[]MusicNFTMgrTrackListing)

	return out0, err

}

// FetchTrackListings is a free data retrieval call binding the contract method 0x0e543f4a.
//
// Solidity: function fetchTrackListings() view returns((uint256,address,uint256,address,address)[])
func (_MusicNFTMgr *MusicNFTMgrSession) FetchTrackListings() ([]MusicNFTMgrTrackListing, error) {
	return _MusicNFTMgr.Contract.FetchTrackListings(&_MusicNFTMgr.CallOpts)
}

// FetchTrackListings is a free data retrieval call binding the contract method 0x0e543f4a.
//
// Solidity: function fetchTrackListings() view returns((uint256,address,uint256,address,address)[])
func (_MusicNFTMgr *MusicNFTMgrCallerSession) FetchTrackListings() ([]MusicNFTMgrTrackListing, error) {
	return _MusicNFTMgr.Contract.FetchTrackListings(&_MusicNFTMgr.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 balance)
func (_MusicNFTMgr *MusicNFTMgrCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 balance)
func (_MusicNFTMgr *MusicNFTMgrSession) GetContractBalance() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetContractBalance(&_MusicNFTMgr.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256 balance)
func (_MusicNFTMgr *MusicNFTMgrCallerSession) GetContractBalance() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetContractBalance(&_MusicNFTMgr.CallOpts)
}

// GetNumTrackListings is a free data retrieval call binding the contract method 0x4e7cc4f3.
//
// Solidity: function getNumTrackListings() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCaller) GetNumTrackListings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "getNumTrackListings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTrackListings is a free data retrieval call binding the contract method 0x4e7cc4f3.
//
// Solidity: function getNumTrackListings() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrSession) GetNumTrackListings() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetNumTrackListings(&_MusicNFTMgr.CallOpts)
}

// GetNumTrackListings is a free data retrieval call binding the contract method 0x4e7cc4f3.
//
// Solidity: function getNumTrackListings() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCallerSession) GetNumTrackListings() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetNumTrackListings(&_MusicNFTMgr.CallOpts)
}

// GetTrackListingPrice is a free data retrieval call binding the contract method 0x9020bc15.
//
// Solidity: function getTrackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCaller) GetTrackListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "getTrackListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTrackListingPrice is a free data retrieval call binding the contract method 0x9020bc15.
//
// Solidity: function getTrackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrSession) GetTrackListingPrice() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetTrackListingPrice(&_MusicNFTMgr.CallOpts)
}

// GetTrackListingPrice is a free data retrieval call binding the contract method 0x9020bc15.
//
// Solidity: function getTrackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCallerSession) GetTrackListingPrice() (*big.Int, error) {
	return _MusicNFTMgr.Contract.GetTrackListingPrice(&_MusicNFTMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MusicNFTMgr *MusicNFTMgrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MusicNFTMgr *MusicNFTMgrSession) Owner() (common.Address, error) {
	return _MusicNFTMgr.Contract.Owner(&_MusicNFTMgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MusicNFTMgr *MusicNFTMgrCallerSession) Owner() (common.Address, error) {
	return _MusicNFTMgr.Contract.Owner(&_MusicNFTMgr.CallOpts)
}

// TrackListingPrice is a free data retrieval call binding the contract method 0x3b6f4318.
//
// Solidity: function trackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCaller) TrackListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MusicNFTMgr.contract.Call(opts, &out, "trackListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TrackListingPrice is a free data retrieval call binding the contract method 0x3b6f4318.
//
// Solidity: function trackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrSession) TrackListingPrice() (*big.Int, error) {
	return _MusicNFTMgr.Contract.TrackListingPrice(&_MusicNFTMgr.CallOpts)
}

// TrackListingPrice is a free data retrieval call binding the contract method 0x3b6f4318.
//
// Solidity: function trackListingPrice() view returns(uint256)
func (_MusicNFTMgr *MusicNFTMgrCallerSession) TrackListingPrice() (*big.Int, error) {
	return _MusicNFTMgr.Contract.TrackListingPrice(&_MusicNFTMgr.CallOpts)
}

// CreateTrackListing is a paid mutator transaction binding the contract method 0x1fb1a137.
//
// Solidity: function createTrackListing(address nftContract, uint256 tokenId) payable returns()
func (_MusicNFTMgr *MusicNFTMgrTransactor) CreateTrackListing(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MusicNFTMgr.contract.Transact(opts, "createTrackListing", nftContract, tokenId)
}

// CreateTrackListing is a paid mutator transaction binding the contract method 0x1fb1a137.
//
// Solidity: function createTrackListing(address nftContract, uint256 tokenId) payable returns()
func (_MusicNFTMgr *MusicNFTMgrSession) CreateTrackListing(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.CreateTrackListing(&_MusicNFTMgr.TransactOpts, nftContract, tokenId)
}

// CreateTrackListing is a paid mutator transaction binding the contract method 0x1fb1a137.
//
// Solidity: function createTrackListing(address nftContract, uint256 tokenId) payable returns()
func (_MusicNFTMgr *MusicNFTMgrTransactorSession) CreateTrackListing(nftContract common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MusicNFTMgr.Contract.CreateTrackListing(&_MusicNFTMgr.TransactOpts, nftContract, tokenId)
}

// MusicNFTMgrTrackListingCreatedIterator is returned from FilterTrackListingCreated and is used to iterate over the raw logs and unpacked data for TrackListingCreated events raised by the MusicNFTMgr contract.
type MusicNFTMgrTrackListingCreatedIterator struct {
	Event *MusicNFTMgrTrackListingCreated // Event containing the contract specifics and raw log

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
func (it *MusicNFTMgrTrackListingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MusicNFTMgrTrackListingCreated)
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
		it.Event = new(MusicNFTMgrTrackListingCreated)
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
func (it *MusicNFTMgrTrackListingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MusicNFTMgrTrackListingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MusicNFTMgrTrackListingCreated represents a TrackListingCreated event raised by the MusicNFTMgr contract.
type MusicNFTMgrTrackListingCreated struct {
	TrackId     *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Artist      common.Address
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTrackListingCreated is a free log retrieval operation binding the contract event 0x52b8746476a7ddd0aa159ae8c78368ffecfa6c31192acfbe470ef2c961d1d678.
//
// Solidity: event TrackListingCreated(uint256 indexed trackId, address indexed nftContract, uint256 indexed tokenId, address artist, address owner)
func (_MusicNFTMgr *MusicNFTMgrFilterer) FilterTrackListingCreated(opts *bind.FilterOpts, trackId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*MusicNFTMgrTrackListingCreatedIterator, error) {

	var trackIdRule []interface{}
	for _, trackIdItem := range trackId {
		trackIdRule = append(trackIdRule, trackIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MusicNFTMgr.contract.FilterLogs(opts, "TrackListingCreated", trackIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MusicNFTMgrTrackListingCreatedIterator{contract: _MusicNFTMgr.contract, event: "TrackListingCreated", logs: logs, sub: sub}, nil
}

// WatchTrackListingCreated is a free log subscription operation binding the contract event 0x52b8746476a7ddd0aa159ae8c78368ffecfa6c31192acfbe470ef2c961d1d678.
//
// Solidity: event TrackListingCreated(uint256 indexed trackId, address indexed nftContract, uint256 indexed tokenId, address artist, address owner)
func (_MusicNFTMgr *MusicNFTMgrFilterer) WatchTrackListingCreated(opts *bind.WatchOpts, sink chan<- *MusicNFTMgrTrackListingCreated, trackId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var trackIdRule []interface{}
	for _, trackIdItem := range trackId {
		trackIdRule = append(trackIdRule, trackIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MusicNFTMgr.contract.WatchLogs(opts, "TrackListingCreated", trackIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MusicNFTMgrTrackListingCreated)
				if err := _MusicNFTMgr.contract.UnpackLog(event, "TrackListingCreated", log); err != nil {
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

// ParseTrackListingCreated is a log parse operation binding the contract event 0x52b8746476a7ddd0aa159ae8c78368ffecfa6c31192acfbe470ef2c961d1d678.
//
// Solidity: event TrackListingCreated(uint256 indexed trackId, address indexed nftContract, uint256 indexed tokenId, address artist, address owner)
func (_MusicNFTMgr *MusicNFTMgrFilterer) ParseTrackListingCreated(log types.Log) (*MusicNFTMgrTrackListingCreated, error) {
	event := new(MusicNFTMgrTrackListingCreated)
	if err := _MusicNFTMgr.contract.UnpackLog(event, "TrackListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
