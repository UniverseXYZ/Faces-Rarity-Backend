// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"premintedTokensCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_polymorphPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"BulkBuyLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPolymorphPrice\",\"type\":\"uint256\"}],\"name\":\"PolymorphPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumPolymorph.PolymorphEventType\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"TokenMorphed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"arweaveAssetsJSONChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arweaveAssetsJSON\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bulkBuyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"polymorphPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bulkBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPolymorphPrice\",\"type\":\"uint256\"}],\"name\":\"setPolymorphPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"setBulkBuyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_arweaveAssetsJSON\",\"type\":\"string\"}],\"name\":\"setArweaveAssetsJSON\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Store.Contract.DEFAULTADMINROLE(&_Store.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Store.Contract.DEFAULTADMINROLE(&_Store.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreSession) MINTERROLE() ([32]byte, error) {
	return _Store.Contract.MINTERROLE(&_Store.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) MINTERROLE() ([32]byte, error) {
	return _Store.Contract.MINTERROLE(&_Store.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Store *StoreCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Store *StoreSession) PAUSERROLE() ([32]byte, error) {
	return _Store.Contract.PAUSERROLE(&_Store.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Store.Contract.PAUSERROLE(&_Store.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_Store *StoreCaller) ArweaveAssetsJSON(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "arweaveAssetsJSON")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_Store *StoreSession) ArweaveAssetsJSON() (string, error) {
	return _Store.Contract.ArweaveAssetsJSON(&_Store.CallOpts)
}

// ArweaveAssetsJSON is a free data retrieval call binding the contract method 0xf528a627.
//
// Solidity: function arweaveAssetsJSON() view returns(string)
func (_Store *StoreCallerSession) ArweaveAssetsJSON() (string, error) {
	return _Store.Contract.ArweaveAssetsJSON(&_Store.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreSession) BaseURI() (string, error) {
	return _Store.Contract.BaseURI(&_Store.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreCallerSession) BaseURI() (string, error) {
	return _Store.Contract.BaseURI(&_Store.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Store *StoreCaller) BulkBuyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "bulkBuyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Store *StoreSession) BulkBuyLimit() (*big.Int, error) {
	return _Store.Contract.BulkBuyLimit(&_Store.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_Store *StoreCallerSession) BulkBuyLimit() (*big.Int, error) {
	return _Store.Contract.BulkBuyLimit(&_Store.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Store *StoreCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Store *StoreSession) DaoAddress() (common.Address, error) {
	return _Store.Contract.DaoAddress(&_Store.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_Store *StoreCallerSession) DaoAddress() (common.Address, error) {
	return _Store.Contract.DaoAddress(&_Store.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Store *StoreCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Store *StoreSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Store.Contract.GeneOf(&_Store.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_Store *StoreCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _Store.Contract.GeneOf(&_Store.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Store.Contract.GetRoleAdmin(&_Store.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Store.Contract.GetRoleAdmin(&_Store.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Store *StoreCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Store *StoreSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Store.Contract.GetRoleMember(&_Store.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Store *StoreCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Store.Contract.GetRoleMember(&_Store.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Store *StoreCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Store *StoreSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Store.Contract.GetRoleMemberCount(&_Store.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Store *StoreCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Store.Contract.GetRoleMemberCount(&_Store.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Store.Contract.HasRole(&_Store.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Store.Contract.HasRole(&_Store.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Store *StoreCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Store *StoreSession) LastTokenId() (*big.Int, error) {
	return _Store.Contract.LastTokenId(&_Store.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_Store *StoreCallerSession) LastTokenId() (*big.Int, error) {
	return _Store.Contract.LastTokenId(&_Store.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Store *StoreCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Store *StoreSession) MaxSupply() (*big.Int, error) {
	return _Store.Contract.MaxSupply(&_Store.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Store *StoreCallerSession) MaxSupply() (*big.Int, error) {
	return _Store.Contract.MaxSupply(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreSession) Paused() (bool, error) {
	return _Store.Contract.Paused(&_Store.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreCallerSession) Paused() (bool, error) {
	return _Store.Contract.Paused(&_Store.CallOpts)
}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_Store *StoreCaller) PolymorphPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "polymorphPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_Store *StoreSession) PolymorphPrice() (*big.Int, error) {
	return _Store.Contract.PolymorphPrice(&_Store.CallOpts)
}

// PolymorphPrice is a free data retrieval call binding the contract method 0x704ec036.
//
// Solidity: function polymorphPrice() view returns(uint256)
func (_Store *StoreCallerSession) PolymorphPrice() (*big.Int, error) {
	return _Store.Contract.PolymorphPrice(&_Store.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenByIndex(&_Store.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenByIndex(&_Store.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenOfOwnerByIndex(&_Store.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenOfOwnerByIndex(&_Store.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_Store *StoreTransactor) BulkBuy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "bulkBuy", amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_Store *StoreSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BulkBuy(&_Store.TransactOpts, amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_Store *StoreTransactorSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BulkBuy(&_Store.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Store *StoreTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Store *StoreSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Store *StoreTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GrantRole(&_Store.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GrantRole(&_Store.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Store *StoreTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Store *StoreSession) Mint() (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_Store *StoreTransactorSession) Mint() (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Store *StoreTransactor) Mint0(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mint0", to)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Store *StoreSession) Mint0(to common.Address) (*types.Transaction, error) {
	return _Store.Contract.Mint0(&_Store.TransactOpts, to)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_Store *StoreTransactorSession) Mint0(to common.Address) (*types.Transaction, error) {
	return _Store.Contract.Mint0(&_Store.TransactOpts, to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreSession) Pause() (*types.Transaction, error) {
	return _Store.Contract.Pause(&_Store.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreTransactorSession) Pause() (*types.Transaction, error) {
	return _Store.Contract.Pause(&_Store.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RenounceRole(&_Store.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RenounceRole(&_Store.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RevokeRole(&_Store.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RevokeRole(&_Store.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_Store *StoreTransactor) SetArweaveAssetsJSON(opts *bind.TransactOpts, _arweaveAssetsJSON string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setArweaveAssetsJSON", _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_Store *StoreSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _Store.Contract.SetArweaveAssetsJSON(&_Store.TransactOpts, _arweaveAssetsJSON)
}

// SetArweaveAssetsJSON is a paid mutator transaction binding the contract method 0x56b1b300.
//
// Solidity: function setArweaveAssetsJSON(string _arweaveAssetsJSON) returns()
func (_Store *StoreTransactorSession) SetArweaveAssetsJSON(_arweaveAssetsJSON string) (*types.Transaction, error) {
	return _Store.Contract.SetArweaveAssetsJSON(&_Store.TransactOpts, _arweaveAssetsJSON)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Store *StoreTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Store *StoreSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Store.Contract.SetBaseURI(&_Store.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_Store *StoreTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _Store.Contract.SetBaseURI(&_Store.TransactOpts, _baseURI)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_Store *StoreTransactor) SetBulkBuyLimit(opts *bind.TransactOpts, _bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setBulkBuyLimit", _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_Store *StoreSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetBulkBuyLimit(&_Store.TransactOpts, _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_Store *StoreTransactorSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetBulkBuyLimit(&_Store.TransactOpts, _bulkBuyLimit)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Store *StoreTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Store *StoreSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetMaxSupply(&_Store.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_Store *StoreTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetMaxSupply(&_Store.TransactOpts, _maxSupply)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_Store *StoreTransactor) SetPolymorphPrice(opts *bind.TransactOpts, newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setPolymorphPrice", newPolymorphPrice)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_Store *StoreSession) SetPolymorphPrice(newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetPolymorphPrice(&_Store.TransactOpts, newPolymorphPrice)
}

// SetPolymorphPrice is a paid mutator transaction binding the contract method 0x017f1e34.
//
// Solidity: function setPolymorphPrice(uint256 newPolymorphPrice) returns()
func (_Store *StoreTransactorSession) SetPolymorphPrice(newPolymorphPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SetPolymorphPrice(&_Store.TransactOpts, newPolymorphPrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreSession) Unpause() (*types.Transaction, error) {
	return _Store.Contract.Unpause(&_Store.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _Store.Contract.Unpause(&_Store.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreSession) Receive() (*types.Transaction, error) {
	return _Store.Contract.Receive(&_Store.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreTransactorSession) Receive() (*types.Transaction, error) {
	return _Store.Contract.Receive(&_Store.TransactOpts)
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

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
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
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
		it.Event = new(StoreApproval)
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
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Store contract.
type StoreApprovalForAllIterator struct {
	Event *StoreApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApprovalForAll)
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
		it.Event = new(StoreApprovalForAll)
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
func (it *StoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApprovalForAll represents a ApprovalForAll event raised by the Store contract.
type StoreApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StoreApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalForAllIterator{contract: _Store.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StoreApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApprovalForAll)
				if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) ParseApprovalForAll(log types.Log) (*StoreApprovalForAll, error) {
	event := new(StoreApprovalForAll)
	if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the Store contract.
type StoreBaseURIChangedIterator struct {
	Event *StoreBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *StoreBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBaseURIChanged)
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
		it.Event = new(StoreBaseURIChanged)
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
func (it *StoreBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBaseURIChanged represents a BaseURIChanged event raised by the Store contract.
type StoreBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Store *StoreFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*StoreBaseURIChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &StoreBaseURIChangedIterator{contract: _Store.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Store *StoreFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *StoreBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBaseURIChanged)
				if err := _Store.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_Store *StoreFilterer) ParseBaseURIChanged(log types.Log) (*StoreBaseURIChanged, error) {
	event := new(StoreBaseURIChanged)
	if err := _Store.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreBulkBuyLimitChangedIterator is returned from FilterBulkBuyLimitChanged and is used to iterate over the raw logs and unpacked data for BulkBuyLimitChanged events raised by the Store contract.
type StoreBulkBuyLimitChangedIterator struct {
	Event *StoreBulkBuyLimitChanged // Event containing the contract specifics and raw log

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
func (it *StoreBulkBuyLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBulkBuyLimitChanged)
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
		it.Event = new(StoreBulkBuyLimitChanged)
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
func (it *StoreBulkBuyLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBulkBuyLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBulkBuyLimitChanged represents a BulkBuyLimitChanged event raised by the Store contract.
type StoreBulkBuyLimitChanged struct {
	NewBulkBuyLimit *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBulkBuyLimitChanged is a free log retrieval operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_Store *StoreFilterer) FilterBulkBuyLimitChanged(opts *bind.FilterOpts) (*StoreBulkBuyLimitChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return &StoreBulkBuyLimitChangedIterator{contract: _Store.contract, event: "BulkBuyLimitChanged", logs: logs, sub: sub}, nil
}

// WatchBulkBuyLimitChanged is a free log subscription operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_Store *StoreFilterer) WatchBulkBuyLimitChanged(opts *bind.WatchOpts, sink chan<- *StoreBulkBuyLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBulkBuyLimitChanged)
				if err := _Store.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
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

// ParseBulkBuyLimitChanged is a log parse operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_Store *StoreFilterer) ParseBulkBuyLimitChanged(log types.Log) (*StoreBulkBuyLimitChanged, error) {
	event := new(StoreBulkBuyLimitChanged)
	if err := _Store.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreMaxSupplyChangedIterator is returned from FilterMaxSupplyChanged and is used to iterate over the raw logs and unpacked data for MaxSupplyChanged events raised by the Store contract.
type StoreMaxSupplyChangedIterator struct {
	Event *StoreMaxSupplyChanged // Event containing the contract specifics and raw log

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
func (it *StoreMaxSupplyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMaxSupplyChanged)
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
		it.Event = new(StoreMaxSupplyChanged)
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
func (it *StoreMaxSupplyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMaxSupplyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMaxSupplyChanged represents a MaxSupplyChanged event raised by the Store contract.
type StoreMaxSupplyChanged struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyChanged is a free log retrieval operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Store *StoreFilterer) FilterMaxSupplyChanged(opts *bind.FilterOpts) (*StoreMaxSupplyChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return &StoreMaxSupplyChangedIterator{contract: _Store.contract, event: "MaxSupplyChanged", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyChanged is a free log subscription operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Store *StoreFilterer) WatchMaxSupplyChanged(opts *bind.WatchOpts, sink chan<- *StoreMaxSupplyChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMaxSupplyChanged)
				if err := _Store.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
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

// ParseMaxSupplyChanged is a log parse operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_Store *StoreFilterer) ParseMaxSupplyChanged(log types.Log) (*StoreMaxSupplyChanged, error) {
	event := new(StoreMaxSupplyChanged)
	if err := _Store.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Store contract.
type StorePausedIterator struct {
	Event *StorePaused // Event containing the contract specifics and raw log

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
func (it *StorePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePaused)
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
		it.Event = new(StorePaused)
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
func (it *StorePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePaused represents a Paused event raised by the Store contract.
type StorePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Store *StoreFilterer) FilterPaused(opts *bind.FilterOpts) (*StorePausedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StorePausedIterator{contract: _Store.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Store *StoreFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StorePaused) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePaused)
				if err := _Store.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Store *StoreFilterer) ParsePaused(log types.Log) (*StorePaused, error) {
	event := new(StorePaused)
	if err := _Store.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePolymorphPriceChangedIterator is returned from FilterPolymorphPriceChanged and is used to iterate over the raw logs and unpacked data for PolymorphPriceChanged events raised by the Store contract.
type StorePolymorphPriceChangedIterator struct {
	Event *StorePolymorphPriceChanged // Event containing the contract specifics and raw log

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
func (it *StorePolymorphPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePolymorphPriceChanged)
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
		it.Event = new(StorePolymorphPriceChanged)
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
func (it *StorePolymorphPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePolymorphPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePolymorphPriceChanged represents a PolymorphPriceChanged event raised by the Store contract.
type StorePolymorphPriceChanged struct {
	NewPolymorphPrice *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPolymorphPriceChanged is a free log retrieval operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_Store *StoreFilterer) FilterPolymorphPriceChanged(opts *bind.FilterOpts) (*StorePolymorphPriceChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "PolymorphPriceChanged")
	if err != nil {
		return nil, err
	}
	return &StorePolymorphPriceChangedIterator{contract: _Store.contract, event: "PolymorphPriceChanged", logs: logs, sub: sub}, nil
}

// WatchPolymorphPriceChanged is a free log subscription operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_Store *StoreFilterer) WatchPolymorphPriceChanged(opts *bind.WatchOpts, sink chan<- *StorePolymorphPriceChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "PolymorphPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePolymorphPriceChanged)
				if err := _Store.contract.UnpackLog(event, "PolymorphPriceChanged", log); err != nil {
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

// ParsePolymorphPriceChanged is a log parse operation binding the contract event 0x6a08b3bba14e54ee218389c7c7444e619f3897465dc06757938cfd01a6957f6c.
//
// Solidity: event PolymorphPriceChanged(uint256 newPolymorphPrice)
func (_Store *StoreFilterer) ParsePolymorphPriceChanged(log types.Log) (*StorePolymorphPriceChanged, error) {
	event := new(StorePolymorphPriceChanged)
	if err := _Store.contract.UnpackLog(event, "PolymorphPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Store contract.
type StoreRoleAdminChangedIterator struct {
	Event *StoreRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StoreRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleAdminChanged)
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
		it.Event = new(StoreRoleAdminChanged)
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
func (it *StoreRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleAdminChanged represents a RoleAdminChanged event raised by the Store contract.
type StoreRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StoreRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleAdminChangedIterator{contract: _Store.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StoreRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleAdminChanged)
				if err := _Store.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) ParseRoleAdminChanged(log types.Log) (*StoreRoleAdminChanged, error) {
	event := new(StoreRoleAdminChanged)
	if err := _Store.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Store contract.
type StoreRoleGrantedIterator struct {
	Event *StoreRoleGranted // Event containing the contract specifics and raw log

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
func (it *StoreRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleGranted)
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
		it.Event = new(StoreRoleGranted)
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
func (it *StoreRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleGranted represents a RoleGranted event raised by the Store contract.
type StoreRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoreRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleGrantedIterator{contract: _Store.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StoreRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleGranted)
				if err := _Store.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) ParseRoleGranted(log types.Log) (*StoreRoleGranted, error) {
	event := new(StoreRoleGranted)
	if err := _Store.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Store contract.
type StoreRoleRevokedIterator struct {
	Event *StoreRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StoreRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleRevoked)
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
		it.Event = new(StoreRoleRevoked)
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
func (it *StoreRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleRevoked represents a RoleRevoked event raised by the Store contract.
type StoreRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoreRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleRevokedIterator{contract: _Store.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StoreRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleRevoked)
				if err := _Store.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) ParseRoleRevoked(log types.Log) (*StoreRoleRevoked, error) {
	event := new(StoreRoleRevoked)
	if err := _Store.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the Store contract.
type StoreTokenMintedIterator struct {
	Event *StoreTokenMinted // Event containing the contract specifics and raw log

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
func (it *StoreTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTokenMinted)
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
		it.Event = new(StoreTokenMinted)
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
func (it *StoreTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTokenMinted represents a TokenMinted event raised by the Store contract.
type StoreTokenMinted struct {
	TokenId *big.Int
	NewGene *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_Store *StoreFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*StoreTokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreTokenMintedIterator{contract: _Store.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_Store *StoreFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *StoreTokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTokenMinted)
				if err := _Store.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_Store *StoreFilterer) ParseTokenMinted(log types.Log) (*StoreTokenMinted, error) {
	event := new(StoreTokenMinted)
	if err := _Store.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTokenMorphedIterator is returned from FilterTokenMorphed and is used to iterate over the raw logs and unpacked data for TokenMorphed events raised by the Store contract.
type StoreTokenMorphedIterator struct {
	Event *StoreTokenMorphed // Event containing the contract specifics and raw log

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
func (it *StoreTokenMorphedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTokenMorphed)
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
		it.Event = new(StoreTokenMorphed)
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
func (it *StoreTokenMorphedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTokenMorphedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTokenMorphed represents a TokenMorphed event raised by the Store contract.
type StoreTokenMorphed struct {
	TokenId   *big.Int
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMorphed is a free log retrieval operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_Store *StoreFilterer) FilterTokenMorphed(opts *bind.FilterOpts, tokenId []*big.Int) (*StoreTokenMorphedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreTokenMorphedIterator{contract: _Store.contract, event: "TokenMorphed", logs: logs, sub: sub}, nil
}

// WatchTokenMorphed is a free log subscription operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_Store *StoreFilterer) WatchTokenMorphed(opts *bind.WatchOpts, sink chan<- *StoreTokenMorphed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTokenMorphed)
				if err := _Store.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
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

// ParseTokenMorphed is a log parse operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_Store *StoreFilterer) ParseTokenMorphed(log types.Log) (*StoreTokenMorphed, error) {
	event := new(StoreTokenMorphed)
	if err := _Store.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Store contract.
type StoreUnpausedIterator struct {
	Event *StoreUnpaused // Event containing the contract specifics and raw log

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
func (it *StoreUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUnpaused)
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
		it.Event = new(StoreUnpaused)
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
func (it *StoreUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUnpaused represents a Unpaused event raised by the Store contract.
type StoreUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Store *StoreFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StoreUnpausedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StoreUnpausedIterator{contract: _Store.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Store *StoreFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StoreUnpaused) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUnpaused)
				if err := _Store.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Store *StoreFilterer) ParseUnpaused(log types.Log) (*StoreUnpaused, error) {
	event := new(StoreUnpaused)
	if err := _Store.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreArweaveAssetsJSONChangedIterator is returned from FilterArweaveAssetsJSONChanged and is used to iterate over the raw logs and unpacked data for ArweaveAssetsJSONChanged events raised by the Store contract.
type StoreArweaveAssetsJSONChangedIterator struct {
	Event *StoreArweaveAssetsJSONChanged // Event containing the contract specifics and raw log

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
func (it *StoreArweaveAssetsJSONChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreArweaveAssetsJSONChanged)
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
		it.Event = new(StoreArweaveAssetsJSONChanged)
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
func (it *StoreArweaveAssetsJSONChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreArweaveAssetsJSONChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreArweaveAssetsJSONChanged represents a ArweaveAssetsJSONChanged event raised by the Store contract.
type StoreArweaveAssetsJSONChanged struct {
	ArweaveAssetsJSON string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterArweaveAssetsJSONChanged is a free log retrieval operation binding the contract event 0xe1ea6e62a0b360acc613f2021ab3c5d36492a25bce6d0d9940a6b497c87363df.
//
// Solidity: event arweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_Store *StoreFilterer) FilterArweaveAssetsJSONChanged(opts *bind.FilterOpts) (*StoreArweaveAssetsJSONChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "arweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return &StoreArweaveAssetsJSONChangedIterator{contract: _Store.contract, event: "arweaveAssetsJSONChanged", logs: logs, sub: sub}, nil
}

// WatchArweaveAssetsJSONChanged is a free log subscription operation binding the contract event 0xe1ea6e62a0b360acc613f2021ab3c5d36492a25bce6d0d9940a6b497c87363df.
//
// Solidity: event arweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_Store *StoreFilterer) WatchArweaveAssetsJSONChanged(opts *bind.WatchOpts, sink chan<- *StoreArweaveAssetsJSONChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "arweaveAssetsJSONChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreArweaveAssetsJSONChanged)
				if err := _Store.contract.UnpackLog(event, "arweaveAssetsJSONChanged", log); err != nil {
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

// ParseArweaveAssetsJSONChanged is a log parse operation binding the contract event 0xe1ea6e62a0b360acc613f2021ab3c5d36492a25bce6d0d9940a6b497c87363df.
//
// Solidity: event arweaveAssetsJSONChanged(string arweaveAssetsJSON)
func (_Store *StoreFilterer) ParseArweaveAssetsJSONChanged(log types.Log) (*StoreArweaveAssetsJSONChanged, error) {
	event := new(StoreArweaveAssetsJSONChanged)
	if err := _Store.contract.UnpackLog(event, "arweaveAssetsJSONChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
